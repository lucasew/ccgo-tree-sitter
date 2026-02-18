// Code generated for linux/amd64 by 'ccgo /tmp/tree-sitter-gen-1414839497/core_complete.c -o /tmp/tree-sitter-gen-1414839497/core.go', DO NOT EDIT.

//go:build linux && amd64

package grammar

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

type TSParser = struct {
	Flexer                           Lexer
	Fstack                           uintptr
	Ftree_pool                       SubtreePool
	Flanguage                        uintptr
	Fwasm_store                      uintptr
	Freduce_actions                  ReduceActionSet
	Ffinished_tree                   Subtree
	Ftrailing_extras                 SubtreeArray
	Ftrailing_extras2                SubtreeArray
	Fscratch_trees                   SubtreeArray
	Ftoken_cache                     TokenCache
	Freusable_node                   ReusableNode
	Fexternal_scanner_payload        uintptr
	Fdot_graph_file                  uintptr
	Faccept_count                    uint32
	Foperation_count                 uint32
	Fold_tree                        Subtree
	Fincluded_range_differences      TSRangeArray
	Fparse_options                   TSParseOptions
	Fparse_state                     TSParseState
	Fincluded_range_difference_index uint32
	Fhas_scanner_error               uint8
	Fcanceled_balancing              uint8
	Fhas_error                       uint8
}

type TSTree = struct {
	Froot                 Subtree
	Flanguage             uintptr
	Fincluded_ranges      uintptr
	Fincluded_range_count uint32
}

type TSQuery = struct {
	Fcaptures            SymbolTable
	Fpredicate_values    SymbolTable
	Fcapture_quantifiers struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fsteps struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fpattern_map struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fpredicate_steps struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fpatterns struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fstep_offsets struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fnegated_fields struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fstring_buffer struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Frepeat_symbols_with_rootless_patterns struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Flanguage                    uintptr
	Fwildcard_root_pattern_count uint16_t
}

type TSQueryCursor = struct {
	Fquery  uintptr
	Fcursor TSTreeCursor
	Fstates struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Ffinished_states struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fcapture_list_pool      CaptureListPool
	Fdepth                  uint32_t
	Fmax_start_depth        uint32_t
	Fincluded_range         TSRange
	Fcontaining_range       TSRange
	Fnext_state_id          uint32_t
	Fquery_options          uintptr
	Fquery_state            TSQueryCursorState
	Foperation_count        uint32
	Fon_visible_node        uint8
	Fascending              uint8
	Fhalted                 uint8
	Fdid_exceed_match_limit uint8
}

type TSDecodeFunction = uintptr

type DecodeFunction = uintptr

type TSInputEncoding1 = int32

type TSInputEncoding = int32

const TSInputEncodingUTF8 = 0
const TSInputEncodingUTF16LE = 1
const TSInputEncodingUTF16BE = 2
const TSInputEncodingCustom = 3

type TSSymbolType1 = int32

type TSSymbolType = int32

const TSSymbolTypeRegular = 0
const TSSymbolTypeAnonymous = 1
const TSSymbolTypeSupertype = 2
const TSSymbolTypeAuxiliary = 3

type TSPoint = struct {
	Frow    uint32_t
	Fcolumn uint32_t
}

type TSRange = struct {
	Fstart_point TSPoint
	Fend_point   TSPoint
	Fstart_byte  uint32_t
	Fend_byte    uint32_t
}

type TSInput = struct {
	Fpayload  uintptr
	Fread     uintptr
	Fencoding TSInputEncoding1
	Fdecode   TSDecodeFunction
}

type TSParseState = struct {
	Fpayload             uintptr
	Fcurrent_byte_offset uint32_t
	Fhas_error           uint8
}

type TSParseOptions = struct {
	Fpayload           uintptr
	Fprogress_callback uintptr
}

type TSLogType1 = int32

type TSLogType = int32

const TSLogTypeParse = 0
const TSLogTypeLex = 1

type TSLogger = struct {
	Fpayload uintptr
	Flog     uintptr
}

type TSInputEdit = struct {
	Fstart_byte    uint32_t
	Fold_end_byte  uint32_t
	Fnew_end_byte  uint32_t
	Fstart_point   TSPoint
	Fold_end_point TSPoint
	Fnew_end_point TSPoint
}

type TSNode = struct {
	Fcontext [4]uint32_t
	Fid      uintptr
	Ftree    uintptr
}

type TSTreeCursor = struct {
	Ftree    uintptr
	Fid      uintptr
	Fcontext [3]uint32_t
}

type TSQueryCapture = struct {
	Fnode  TSNode
	Findex uint32_t
}

type TSQuantifier1 = int32

type TSQuantifier = int32

const TSQuantifierZero = 0
const TSQuantifierZeroOrOne = 1
const TSQuantifierZeroOrMore = 2
const TSQuantifierOne = 3
const TSQuantifierOneOrMore = 4

type TSQueryMatch = struct {
	Fid            uint32_t
	Fpattern_index uint16_t
	Fcapture_count uint16_t
	Fcaptures      uintptr
}

type TSQueryPredicateStepType1 = int32

type TSQueryPredicateStepType = int32

const TSQueryPredicateStepTypeDone = 0
const TSQueryPredicateStepTypeCapture = 1
const TSQueryPredicateStepTypeString = 2

type TSQueryPredicateStep = struct {
	Ftype1    TSQueryPredicateStepType1
	Fvalue_id uint32_t
}

type TSQueryError1 = int32

type TSQueryError = int32

const TSQueryErrorNone = 0
const TSQueryErrorSyntax = 1
const TSQueryErrorNodeType = 2
const TSQueryErrorField = 3
const TSQueryErrorCapture = 4
const TSQueryErrorStructure = 5
const TSQueryErrorLanguage = 6

type TSQueryCursorState = struct {
	Fpayload             uintptr
	Fcurrent_byte_offset uint32_t
}

type TSQueryCursorOptions = struct {
	Fpayload           uintptr
	Fprogress_callback uintptr
}

type TSLanguageMetadata = struct {
	Fmajor_version uint8_t
	Fminor_version uint8_t
	Fpatch_version uint8_t
}

type TSWasmErrorKind = int32

const TSWasmErrorKindNone = 0
const TSWasmErrorKindParse = 1
const TSWasmErrorKindCompile = 2
const TSWasmErrorKindInstantiate = 3
const TSWasmErrorKindAllocate = 4

type TSWasmError = struct {
	Fkind    TSWasmErrorKind
	Fmessage uintptr
}

func ts_malloc_default(tls *libc.TLS, size size_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var result uintptr
	_ = result
	result = libc.Xmalloc(tls, size)
	if size > uint64(0) && !(result != 0) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+8, size))
		libc.Xabort(tls)
	}
	return result
}

func ts_calloc_default(tls *libc.TLS, count size_t, size size_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var result uintptr
	_ = result
	result = libc.Xcalloc(tls, count, size)
	if count > uint64(0) && !(result != 0) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+8, count*size))
		libc.Xabort(tls)
	}
	return result
}

func ts_realloc_default(tls *libc.TLS, buffer uintptr, size size_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var result uintptr
	_ = result
	result = libc.Xrealloc(tls, buffer, size)
	if size > uint64(0) && !(result != 0) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+41, libc.VaList(bp+8, size))
		libc.Xabort(tls)
	}
	return result
}

func init() {
	p := unsafe.Pointer(&ts_current_malloc)
	*(*uintptr)(unsafe.Add(p, 0)) = __ccgo_fp(ts_malloc_default)
}

func init() {
	p := unsafe.Pointer(&ts_current_calloc)
	*(*uintptr)(unsafe.Add(p, 0)) = __ccgo_fp(ts_calloc_default)
}

func init() {
	p := unsafe.Pointer(&ts_current_realloc)
	*(*uintptr)(unsafe.Add(p, 0)) = __ccgo_fp(ts_realloc_default)
}

func init() {
	p := unsafe.Pointer(&ts_current_free)
	*(*uintptr)(unsafe.Add(p, 0)) = __ccgo_fp(libc.Xfree)
}

type __ccgo_fp__Xts_set_allocator_0 = func(*libc.TLS, uint64) uintptr

type __ccgo_fp__Xts_set_allocator_1 = func(*libc.TLS, uint64, uint64) uintptr

type __ccgo_fp__Xts_set_allocator_2 = func(*libc.TLS, uintptr, uint64) uintptr

type __ccgo_fp__Xts_set_allocator_3 = func(*libc.TLS, uintptr)

func ts_set_allocator(tls *libc.TLS, __ccgo_fp_new_malloc uintptr, __ccgo_fp_new_calloc uintptr, __ccgo_fp_new_realloc uintptr, __ccgo_fp_new_free uintptr) {
	var v1 uintptr
	_ = v1
	if __ccgo_fp_new_malloc != 0 {
		v1 = __ccgo_fp_new_malloc
	} else {
		v1 = __ccgo_fp(ts_malloc_default)
	}
	ts_current_malloc = v1
	if __ccgo_fp_new_calloc != 0 {
		v1 = __ccgo_fp_new_calloc
	} else {
		v1 = __ccgo_fp(ts_calloc_default)
	}
	ts_current_calloc = v1
	if __ccgo_fp_new_realloc != 0 {
		v1 = __ccgo_fp_new_realloc
	} else {
		v1 = __ccgo_fp(ts_realloc_default)
	}
	ts_current_realloc = v1
	if __ccgo_fp_new_free != 0 {
		v1 = __ccgo_fp_new_free
	} else {
		v1 = __ccgo_fp(libc.Xfree)
	}
	ts_current_free = v1
}

func point__new(tls *libc.TLS, row uint32, column uint32) (r TSPoint) {
	var result TSPoint
	_ = result
	result = TSPoint{
		Frow:    row,
		Fcolumn: column,
	}
	return result
}

func point_add(tls *libc.TLS, a TSPoint, b TSPoint) (r TSPoint) {
	if b.Frow > uint32(0) {
		return point__new(tls, a.Frow+b.Frow, b.Fcolumn)
	} else {
		return point__new(tls, a.Frow, a.Fcolumn+b.Fcolumn)
	}
	return r
}

func point_sub(tls *libc.TLS, a TSPoint, b TSPoint) (r TSPoint) {
	var v1 uint32
	_ = v1
	if a.Frow > b.Frow {
		return point__new(tls, a.Frow-b.Frow, a.Fcolumn)
	} else {
		if a.Fcolumn >= b.Fcolumn {
			v1 = a.Fcolumn - b.Fcolumn
		} else {
			v1 = uint32(0)
		}
		return point__new(tls, uint32(0), v1)
	}
	return r
}

func point_lte(tls *libc.TLS, a TSPoint, b TSPoint) (r uint8) {
	return libc.BoolUint8(a.Frow < b.Frow || a.Frow == b.Frow && a.Fcolumn <= b.Fcolumn)
}

func point_lt(tls *libc.TLS, a TSPoint, b TSPoint) (r uint8) {
	return libc.BoolUint8(a.Frow < b.Frow || a.Frow == b.Frow && a.Fcolumn < b.Fcolumn)
}

func point_gt(tls *libc.TLS, a TSPoint, b TSPoint) (r uint8) {
	return libc.BoolUint8(a.Frow > b.Frow || a.Frow == b.Frow && a.Fcolumn > b.Fcolumn)
}

func point_gte(tls *libc.TLS, a TSPoint, b TSPoint) (r uint8) {
	return libc.BoolUint8(a.Frow > b.Frow || a.Frow == b.Frow && a.Fcolumn >= b.Fcolumn)
}

func point_eq(tls *libc.TLS, a TSPoint, b TSPoint) (r uint8) {
	return libc.BoolUint8(a.Frow == b.Frow && a.Fcolumn == b.Fcolumn)
}

type Length = struct {
	Fbytes  uint32_t
	Fextent TSPoint
}

var LENGTH_UNDEFINED = Length{
	Fextent: TSPoint{
		Fcolumn: uint32(1),
	},
}
var LENGTH_MAX = Length{
	Fbytes: libc.Uint32FromUint32(4294967295),
	Fextent: TSPoint{
		Frow:    libc.Uint32FromUint32(4294967295),
		Fcolumn: libc.Uint32FromUint32(4294967295),
	},
}

func length_is_undefined(tls *libc.TLS, length Length) (r uint8) {
	return libc.BoolUint8(length.Fbytes == uint32(0) && length.Fextent.Fcolumn != uint32(0))
}

func length_min(tls *libc.TLS, len1 Length, len2 Length) (r Length) {
	var v1 Length
	_ = v1
	if len1.Fbytes < len2.Fbytes {
		v1 = len1
	} else {
		v1 = len2
	}
	return v1
}

func length_add(tls *libc.TLS, len1 Length, len2 Length) (r Length) {
	var result Length
	_ = result
	result.Fbytes = len1.Fbytes + len2.Fbytes
	result.Fextent = point_add(tls, len1.Fextent, len2.Fextent)
	return result
}

func length_sub(tls *libc.TLS, len1 Length, len2 Length) (r Length) {
	var result Length
	var v1 uint32
	_, _ = result, v1
	if len1.Fbytes >= len2.Fbytes {
		v1 = len1.Fbytes - len2.Fbytes
	} else {
		v1 = uint32(0)
	}
	result.Fbytes = v1
	result.Fextent = point_sub(tls, len1.Fextent, len2.Fextent)
	return result
}

func length_zero(tls *libc.TLS) (r Length) {
	var result Length
	_ = result
	result = Length{}
	return result
}

func length_saturating_sub(tls *libc.TLS, len1 Length, len2 Length) (r Length) {
	if len1.Fbytes > len2.Fbytes {
		return length_sub(tls, len1, len2)
	} else {
		return length_zero(tls)
	}
	return r
}

type __locale_struct = struct {
	F__locales       [13]uintptr
	F__ctype_b       uintptr
	F__ctype_tolower uintptr
	F__ctype_toupper uintptr
	F__names         [13]uintptr
}

type __locale_t = uintptr

type locale_t = uintptr

func _array__erase(tls *libc.TLS, self_contents uintptr, size uintptr, element_size size_t, index uint32_t) {
	var contents uintptr
	_ = contents
	_ = libc.Uint64FromInt64(4)
	{
		if !(index < *(*uint32_t)(unsafe.Pointer(size))) {
			libc.X__assert_fail(tls, __ccgo_ts+84, __ccgo_ts+98, int32(199), uintptr(unsafe.Pointer(&__func__)))
		}
	}
	contents = self_contents
	libc.Xmemmove(tls, contents+uintptr(uint64(index)*element_size), contents+uintptr(uint64(index+libc.Uint32FromInt32(1))*element_size), uint64(*(*uint32_t)(unsafe.Pointer(size))-index-libc.Uint32FromInt32(1))*element_size)
	*(*uint32_t)(unsafe.Pointer(size)) = *(*uint32_t)(unsafe.Pointer(size)) - 1
}

var __func__ = [14]int8{'_', 'a', 'r', 'r', 'a', 'y', '_', '_', 'e', 'r', 'a', 's', 'e'}

func _array__reserve(tls *libc.TLS, contents uintptr, capacity uintptr, element_size size_t, new_capacity uint32_t) (r uintptr) {
	var new_contents uintptr
	_ = new_contents
	new_contents = contents
	if new_capacity > *(*uint32_t)(unsafe.Pointer(capacity)) {
		if contents != 0 {
			new_contents = (*(*func(*libc.TLS, uintptr, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_realloc})))(tls, contents, uint64(new_capacity)*element_size)
		} else {
			new_contents = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(new_capacity)*element_size)
		}
		*(*uint32_t)(unsafe.Pointer(capacity)) = new_capacity
	}
	return new_contents
}

func _array__assign(tls *libc.TLS, self_contents uintptr, self_size uintptr, self_capacity uintptr, other_contents uintptr, other_size uint32_t, element_size size_t) (r uintptr) {
	var new_contents uintptr
	_ = new_contents
	new_contents = _array__reserve(tls, self_contents, self_capacity, element_size, other_size)
	*(*uint32_t)(unsafe.Pointer(self_size)) = other_size
	libc.Xmemcpy(tls, new_contents, other_contents, uint64(*(*uint32_t)(unsafe.Pointer(self_size)))*element_size)
	return new_contents
}

func _array__swap(tls *libc.TLS, self_size uintptr, self_capacity uintptr, other_size uintptr, other_capacity uintptr) {
	var tmp_capacity, tmp_size uint32_t
	_, _ = tmp_capacity, tmp_size
	tmp_size = *(*uint32_t)(unsafe.Pointer(self_size))
	tmp_capacity = *(*uint32_t)(unsafe.Pointer(self_capacity))
	*(*uint32_t)(unsafe.Pointer(self_size)) = *(*uint32_t)(unsafe.Pointer(other_size))
	*(*uint32_t)(unsafe.Pointer(self_capacity)) = *(*uint32_t)(unsafe.Pointer(other_capacity))
	*(*uint32_t)(unsafe.Pointer(other_size)) = tmp_size
	*(*uint32_t)(unsafe.Pointer(other_capacity)) = tmp_capacity
}

func _array__grow(tls *libc.TLS, contents uintptr, size uint32_t, capacity uintptr, count uint32_t, element_size size_t) (r uintptr) {
	var new_capacity, new_size uint32_t
	var new_contents uintptr
	_, _, _ = new_capacity, new_contents, new_size
	new_contents = contents
	new_size = size + count
	if new_size > *(*uint32_t)(unsafe.Pointer(capacity)) {
		new_capacity = *(*uint32_t)(unsafe.Pointer(capacity)) * uint32(2)
		if new_capacity < uint32(8) {
			new_capacity = uint32(8)
		}
		if new_capacity < new_size {
			new_capacity = new_size
		}
		new_contents = _array__reserve(tls, contents, capacity, element_size, new_capacity)
	}
	return new_contents
}

func _array__splice(tls *libc.TLS, self_contents uintptr, size uintptr, capacity uintptr, element_size size_t, index uint32_t, old_count uint32_t, new_count uint32_t, elements uintptr) (r uintptr) {
	var contents, new_contents uintptr
	var new_end, new_size, old_end uint32_t
	_, _, _, _, _ = contents, new_contents, new_end, new_size, old_end
	new_size = *(*uint32_t)(unsafe.Pointer(size)) + new_count - old_count
	old_end = index + old_count
	new_end = index + new_count
	_ = libc.Uint64FromInt64(4)
	{
		if !(old_end <= *(*uint32_t)(unsafe.Pointer(size))) {
			libc.X__assert_fail(tls, __ccgo_ts+148, __ccgo_ts+98, int32(263), uintptr(unsafe.Pointer(&__func__1)))
		}
	}
	new_contents = _array__reserve(tls, self_contents, capacity, element_size, new_size)
	contents = new_contents
	if *(*uint32_t)(unsafe.Pointer(size)) > old_end {
		libc.Xmemmove(tls, contents+uintptr(uint64(new_end)*element_size), contents+uintptr(uint64(old_end)*element_size), uint64(*(*uint32_t)(unsafe.Pointer(size))-old_end)*element_size)
	}
	if new_count > uint32(0) {
		if elements != 0 {
			libc.Xmemcpy(tls, contents+uintptr(uint64(index)*element_size), elements, uint64(new_count)*element_size)
		} else {
			libc.Xmemset(tls, contents+uintptr(uint64(index)*element_size), 0, uint64(new_count)*element_size)
		}
	}
	*(*uint32_t)(unsafe.Pointer(size)) += new_count - old_count
	return new_contents
}

var __func__1 = [15]int8{'_', 'a', 'r', 'r', 'a', 'y', '_', '_', 's', 'p', 'l', 'i', 'c', 'e'}

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

type ExternalScannerState = struct {
	F__ccgo0_0 struct {
		Fshort_data  [0][24]int8
		Flong_data   uintptr
		F__ccgo_pad2 [16]byte
	}
	Flength uint32_t
}

type SubtreeInlineData = struct {
	F__ccgo0         uint8
	Fsymbol          uint8_t
	Fparse_state     uint16_t
	Fpadding_columns uint8_t
	F__ccgo5         uint8
	Fpadding_bytes   uint8_t
	Fsize_bytes      uint8_t
}

type SubtreeHeapData = struct {
	Fref_count       uint32_t
	Fpadding         Length
	Fsize            Length
	Flookahead_bytes uint32_t
	Ferror_cost      uint32_t
	Fchild_count     uint32_t
	Fsymbol          TSSymbol
	Fparse_state     TSStateId
	F__ccgo44        uint8
	F__ccgo45        uint8
	F__ccgo19_48     struct {
		Fexternal_scanner_state [0]ExternalScannerState
		Flookahead_char         [0]int32_t
		F__ccgo0_0              struct {
			Fvisible_child_count      uint32_t
			Fnamed_child_count        uint32_t
			Fvisible_descendant_count uint32_t
			Fdynamic_precedence       int32_t
			Frepeat_depth             uint16_t
			Fproduction_id            uint16_t
			Ffirst_leaf               struct {
				Fsymbol      TSSymbol
				Fparse_state TSStateId
			}
		}
		F__ccgo_pad3 [8]byte
	}
}

type Subtree = struct {
	Fptr  [0]uintptr
	Fdata SubtreeInlineData
}

type MutableSubtree = struct {
	Fptr  [0]uintptr
	Fdata SubtreeInlineData
}

type SubtreeArray = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type MutableSubtreeArray = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type SubtreePool = struct {
	Ffree_trees MutableSubtreeArray
	Ftree_stack MutableSubtreeArray
}

func ts_subtree_symbol(tls *libc.TLS, _self Subtree) (r TSSymbol) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = libc.Int32FromUint8((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fsymbol)
	} else {
		v1 = libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol)
	}
	return libc.Uint16FromInt32(v1)
}

func ts_subtree_visible(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x2 >> 1)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x1 >> 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_named(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x4 >> 2)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x2 >> 1)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_extra(tls *libc.TLS, _self Subtree) (r uint8) {
	// NULL check
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x8 >> 3)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x4 >> 2)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_has_changes(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x10 >> 4)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x20 >> 5)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_missing(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x20 >> 5)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 45)) & 0x2 >> 1)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_is_keyword(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = int32(*(*uint8)(unsafe.Pointer(bp + 0)) & 0x40 >> 6)
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 45)) & 0x4 >> 2)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_parse_state(tls *libc.TLS, _self Subtree) (r TSStateId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = libc.Int32FromUint16((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fparse_state)
	} else {
		v1 = libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fparse_state)
	}
	return libc.Uint16FromInt32(v1)
}

func ts_subtree_lookahead_bytes(tls *libc.TLS, _self Subtree) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 uint32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = libc.Uint32FromInt32(int32(*(*uint8)(unsafe.Pointer(bp + 5)) & 0xf0 >> 4))
	} else {
		v1 = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Flookahead_bytes
	}
	return v1
}

func ts_subtree_alloc_size(tls *libc.TLS, child_count uint32_t) (r size_t) {
	return uint64(child_count)*uint64(8) + uint64(80)
}

func ts_subtree_set_extra(tls *libc.TLS, self uintptr, is_extra uint8) {
	if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
		libc.SetBitFieldPtr8Uint8(self+0, is_extra, 3, 0x8)
	} else {
		libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(self))+44, is_extra, 2, 0x4)
	}
}

func ts_subtree_leaf_symbol(tls *libc.TLS, _self Subtree) (r TSSymbol) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		return uint16((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fsymbol)
	}
	if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count == uint32(0) {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol
	}
	return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Ffirst_leaf.Fsymbol
}

func ts_subtree_leaf_parse_state(tls *libc.TLS, _self Subtree) (r TSStateId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		return (*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fparse_state
	}
	if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count == uint32(0) {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fparse_state
	}
	return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Ffirst_leaf.Fparse_state
}

func ts_subtree_padding(tls *libc.TLS, _self Subtree) (r Length) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var result Length
	_ = result
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		result = Length{
			Fbytes: uint32((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fpadding_bytes),
			Fextent: TSPoint{
				Frow:    libc.Uint32FromInt32(int32(*(*uint8)(unsafe.Pointer(bp + 5)) & 0xf >> 0)),
				Fcolumn: uint32((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fpadding_columns),
			},
		}
		return result
	} else {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fpadding
	}
	return r
}

func ts_subtree_size(tls *libc.TLS, _self Subtree) (r Length) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var result Length
	_ = result
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		result = Length{
			Fbytes: uint32((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fsize_bytes),
			Fextent: TSPoint{
				Fcolumn: uint32((*(*SubtreeInlineData)(unsafe.Pointer(bp))).Fsize_bytes),
			},
		}
		return result
	} else {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize
	}
	return r
}

func ts_subtree_total_size(tls *libc.TLS, self Subtree) (r Length) {
	return length_add(tls, ts_subtree_padding(tls, self), ts_subtree_size(tls, self))
}

func ts_subtree_total_bytes(tls *libc.TLS, self Subtree) (r uint32_t) {
	return ts_subtree_total_size(tls, self).Fbytes
}

func ts_subtree_child_count(tls *libc.TLS, _self Subtree) (r uint32_t) {
	// NULL check - if the subtree pointer is 0, return 0
	ptr := *(*uintptr)(unsafe.Pointer(&_self))
	if ptr == 0 {
		return 0
	}
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 uint32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = uint32(0)
	} else {
		v1 = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count
	}
	return v1
}

func ts_subtree_repeat_depth(tls *libc.TLS, _self Subtree) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Frepeat_depth)
	}
	return libc.Uint32FromInt32(v1)
}

func ts_subtree_is_repetition(tls *libc.TLS, _self Subtree) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = libc.BoolInt32(!(int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x2>>1) != 0) && !(int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x1>>0) != 0) && (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count != uint32(0))
	}
	return libc.Uint32FromInt32(v1)
}

func ts_subtree_visible_descendant_count(tls *libc.TLS, _self Subtree) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 uint32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 || (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count == uint32(0) {
		v1 = uint32(0)
	} else {
		v1 = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_descendant_count
	}
	return v1
}

func ts_subtree_visible_child_count(tls *libc.TLS, _self Subtree) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count
	} else {
		return uint32(0)
	}
	return r
}

func ts_subtree_error_cost(tls *libc.TLS, _self Subtree) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 uint32
	_ = v1
	if ts_subtree_missing(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
		return libc.Uint32FromInt32(libc.Int32FromInt32(110) + libc.Int32FromInt32(500))
	} else {
		if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
			v1 = uint32(0)
		} else {
			v1 = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Ferror_cost
		}
		return v1
	}
	return r
}

func ts_subtree_dynamic_precedence(tls *libc.TLS, _self Subtree) (r int32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 || (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count == uint32(0) {
		v1 = 0
	} else {
		v1 = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fdynamic_precedence
	}
	return v1
}

func ts_subtree_production_id(tls *libc.TLS, _self Subtree) (r uint16_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id
	} else {
		return uint16(0)
	}
	return r
}

func ts_subtree_fragile_left(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x8 >> 3)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_fragile_right(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x10 >> 4)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_has_external_tokens(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x40 >> 6)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_has_external_scanner_state_change(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44)) & 0x80 >> 7)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_depends_on_column(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 45)) & 0x1 >> 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_is_fragile(tls *libc.TLS, _self Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = 0
	} else {
		v1 = libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x8>>3) != 0 || int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x10>>4) != 0)
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_subtree_is_error(tls *libc.TLS, self Subtree) (r uint8) {
	return libc.BoolUint8(libc.Int32FromUint16(ts_subtree_symbol(tls, self)) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))))
}

func ts_subtree_is_eof(tls *libc.TLS, self Subtree) (r uint8) {
	return libc.BoolUint8(libc.Int32FromUint16(ts_subtree_symbol(tls, self)) == 0)
}

func ts_subtree_from_mut(tls *libc.TLS, _self MutableSubtree) (r Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*MutableSubtree)(unsafe.Pointer(bp)) = _self
	var _ /* result at bp+8 */ Subtree
	*(*SubtreeInlineData)(unsafe.Pointer(bp + 8)) = *(*SubtreeInlineData)(unsafe.Pointer(bp))
	return *(*Subtree)(unsafe.Pointer(bp + 8))
}

func ts_subtree_to_mut_unsafe(tls *libc.TLS, _self Subtree) (r MutableSubtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var _ /* result at bp+8 */ MutableSubtree
	*(*SubtreeInlineData)(unsafe.Pointer(bp + 8)) = *(*SubtreeInlineData)(unsafe.Pointer(bp))
	return *(*MutableSubtree)(unsafe.Pointer(bp + 8))
}

type TreeCursorEntry = struct {
	Fsubtree                uintptr
	Fposition               Length
	Fchild_index            uint32_t
	Fstructural_child_index uint32_t
	Fdescendant_index       uint32_t
}

type TreeCursor = struct {
	Ftree  uintptr
	Fstack struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Froot_alias_symbol TSSymbol
}

type TreeCursorStep = int32

const TreeCursorStepNone = 0
const TreeCursorStepHidden = 1
const TreeCursorStepVisible = 2

func ts_tree_cursor_current_subtree(tls *libc.TLS, _self uintptr) (r Subtree) {
	var last_entry, self uintptr
	_, _ = last_entry, self
	self = _self
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+225, int32(42), uintptr(unsafe.Pointer(&__func__2)))
		}
	}
	last_entry = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32
	return *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fsubtree))
}

var __func__2 = [31]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'c', 'u', 'r', 'r', 'e', 'n', 't', '_', 's', 'u', 'b', 't', 'r', 'e', 'e'}

type TSRangeArray = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type TableEntry = struct {
	Factions      uintptr
	Faction_count uint32_t
	Fis_reusable  uint8
}

type LookaheadIterator = struct {
	Flanguage       uintptr
	Fdata           uintptr
	Fgroup_end      uintptr
	Fstate          TSStateId
	Ftable_value    uint16_t
	Fsection_index  uint16_t
	Fgroup_count    uint16_t
	Fis_small_state uint8
	Factions        uintptr
	Fsymbol         TSSymbol
	Fnext_state     TSStateId
	Faction_count   uint16_t
}

func ts_language_actions(tls *libc.TLS, self uintptr, state TSStateId, symbol TSSymbol, count uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* entry at bp+0 */ TableEntry
	ts_language_table_entry(tls, self, state, symbol, bp)
	*(*uint32_t)(unsafe.Pointer(count)) = (*(*TableEntry)(unsafe.Pointer(bp))).Faction_count
	return (*(*TableEntry)(unsafe.Pointer(bp))).Factions
}

func ts_language_has_reduce_action(tls *libc.TLS, self uintptr, state TSStateId, symbol TSSymbol) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* entry at bp+0 */ TableEntry
	ts_language_table_entry(tls, self, state, symbol, bp)
	return libc.BoolUint8((*(*TableEntry)(unsafe.Pointer(bp))).Faction_count > uint32(0) && libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer((*(*TableEntry)(unsafe.Pointer(bp))).Factions))))) == int32(TSParseActionTypeReduce))
}

func ts_language_lookup(tls *libc.TLS, self uintptr, state TSStateId, symbol TSSymbol) (r uint16_t) {
	var data, v1, v3, v4, v6 uintptr
	var group_count, section_value, symbol_count uint16_t
	var i, j uint32
	var index uint32_t
	_, _, _, _, _, _, _, _, _, _, _ = data, group_count, i, index, j, section_value, symbol_count, v1, v3, v4, v6
	if uint32(state) >= (*TSLanguage)(unsafe.Pointer(self)).Flarge_state_count {
		index = *(*uint32_t)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fsmall_parse_table_map + uintptr(uint32(state)-(*TSLanguage)(unsafe.Pointer(self)).Flarge_state_count)*4))
		data = (*TSLanguage)(unsafe.Pointer(self)).Fsmall_parse_table + uintptr(index)*2
		v1 = data
		data += 2
		group_count = *(*uint16_t)(unsafe.Pointer(v1))
		i = uint32(0)
		for {
			if !(i < uint32(group_count)) {
				break
			}
			v3 = data
			data += 2
			section_value = *(*uint16_t)(unsafe.Pointer(v3))
			v4 = data
			data += 2
			symbol_count = *(*uint16_t)(unsafe.Pointer(v4))
			j = uint32(0)
			for {
				if !(j < uint32(symbol_count)) {
					break
				}
				v6 = data
				data += 2
				if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(v6))) == libc.Int32FromUint16(symbol) {
					return section_value
				}
				goto _5
			_5:
				;
				j = j + 1
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		return uint16(0)
	} else {
		return *(*uint16_t)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fparse_table + uintptr(uint32(state)*(*TSLanguage)(unsafe.Pointer(self)).Fsymbol_count+uint32(symbol))*2))
	}
	return r
}

func ts_language_has_actions(tls *libc.TLS, self uintptr, state TSStateId, symbol TSSymbol) (r uint8) {
	return libc.BoolUint8(libc.Int32FromUint16(ts_language_lookup(tls, self, state, symbol)) != 0)
}

func ts_language_lookaheads(tls *libc.TLS, self uintptr, state TSStateId) (r LookaheadIterator) {
	var data, group_end uintptr
	var group_count uint16_t
	var index uint32_t
	var is_small_state uint8
	_, _, _, _, _ = data, group_count, group_end, index, is_small_state
	is_small_state = libc.BoolUint8(uint32(state) >= (*TSLanguage)(unsafe.Pointer(self)).Flarge_state_count)
	group_end = libc.UintptrFromInt32(0)
	group_count = uint16(0)
	if is_small_state != 0 {
		index = *(*uint32_t)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fsmall_parse_table_map + uintptr(uint32(state)-(*TSLanguage)(unsafe.Pointer(self)).Flarge_state_count)*4))
		data = (*TSLanguage)(unsafe.Pointer(self)).Fsmall_parse_table + uintptr(index)*2
		group_end = data + uintptr(1)*2
		group_count = *(*uint16_t)(unsafe.Pointer(data))
	} else {
		data = (*TSLanguage)(unsafe.Pointer(self)).Fparse_table + uintptr(uint32(state)*(*TSLanguage)(unsafe.Pointer(self)).Fsymbol_count)*2 - uintptr(1)*2
	}
	return LookaheadIterator{
		Flanguage:       self,
		Fdata:           data,
		Fgroup_end:      group_end,
		Fgroup_count:    group_count,
		Fis_small_state: is_small_state,
		Fsymbol:         libc.Uint16FromInt32(libc.Int32FromInt32(65535)),
	}
}

func ts_lookahead_iterator__next(tls *libc.TLS, self uintptr) (r uint8) {
	var entry, v1, v2 uintptr
	var symbol_count uint32
	_, _, _, _ = entry, symbol_count, v1, v2
	if (*LookaheadIterator)(unsafe.Pointer(self)).Fis_small_state != 0 {
		(*LookaheadIterator)(unsafe.Pointer(self)).Fdata += 2
		if (*LookaheadIterator)(unsafe.Pointer(self)).Fdata == (*LookaheadIterator)(unsafe.Pointer(self)).Fgroup_end {
			if libc.Int32FromUint16((*LookaheadIterator)(unsafe.Pointer(self)).Fgroup_count) == 0 {
				return libc.BoolUint8(0 != 0)
			}
			(*LookaheadIterator)(unsafe.Pointer(self)).Fgroup_count = (*LookaheadIterator)(unsafe.Pointer(self)).Fgroup_count - 1
			v2 = self + 8
			v1 = *(*uintptr)(unsafe.Pointer(v2))
			*(*uintptr)(unsafe.Pointer(v2)) += 2
			(*LookaheadIterator)(unsafe.Pointer(self)).Ftable_value = *(*uint16_t)(unsafe.Pointer(v1))
			v2 = self + 8
			v1 = *(*uintptr)(unsafe.Pointer(v2))
			*(*uintptr)(unsafe.Pointer(v2)) += 2
			symbol_count = uint32(*(*uint16_t)(unsafe.Pointer(v1)))
			(*LookaheadIterator)(unsafe.Pointer(self)).Fgroup_end = (*LookaheadIterator)(unsafe.Pointer(self)).Fdata + uintptr(symbol_count)*2
			(*LookaheadIterator)(unsafe.Pointer(self)).Fsymbol = *(*uint16_t)(unsafe.Pointer((*LookaheadIterator)(unsafe.Pointer(self)).Fdata))
		} else {
			(*LookaheadIterator)(unsafe.Pointer(self)).Fsymbol = *(*uint16_t)(unsafe.Pointer((*LookaheadIterator)(unsafe.Pointer(self)).Fdata))
			return libc.BoolUint8(1 != 0)
		}
	} else {
		for cond := true; cond; cond = !((*LookaheadIterator)(unsafe.Pointer(self)).Ftable_value != 0) {
			(*LookaheadIterator)(unsafe.Pointer(self)).Fdata += 2
			(*LookaheadIterator)(unsafe.Pointer(self)).Fsymbol = (*LookaheadIterator)(unsafe.Pointer(self)).Fsymbol + 1
			if uint32((*LookaheadIterator)(unsafe.Pointer(self)).Fsymbol) >= (*TSLanguage)(unsafe.Pointer((*LookaheadIterator)(unsafe.Pointer(self)).Flanguage)).Fsymbol_count {
				return libc.BoolUint8(0 != 0)
			}
			(*LookaheadIterator)(unsafe.Pointer(self)).Ftable_value = *(*uint16_t)(unsafe.Pointer((*LookaheadIterator)(unsafe.Pointer(self)).Fdata))
		}
	}
	if uint32((*LookaheadIterator)(unsafe.Pointer(self)).Fsymbol) < (*TSLanguage)(unsafe.Pointer((*LookaheadIterator)(unsafe.Pointer(self)).Flanguage)).Ftoken_count {
		entry = (*TSLanguage)(unsafe.Pointer((*LookaheadIterator)(unsafe.Pointer(self)).Flanguage)).Fparse_actions + uintptr((*LookaheadIterator)(unsafe.Pointer(self)).Ftable_value)*8
		(*LookaheadIterator)(unsafe.Pointer(self)).Faction_count = uint16((*(*struct {
			Fcount    uint8_t
			Freusable uint8
		})(unsafe.Pointer(entry))).Fcount)
		(*LookaheadIterator)(unsafe.Pointer(self)).Factions = entry + libc.UintptrFromInt32(1)*8
		(*LookaheadIterator)(unsafe.Pointer(self)).Fnext_state = uint16(0)
	} else {
		(*LookaheadIterator)(unsafe.Pointer(self)).Faction_count = uint16(0)
		(*LookaheadIterator)(unsafe.Pointer(self)).Fnext_state = (*LookaheadIterator)(unsafe.Pointer(self)).Ftable_value
	}
	return libc.BoolUint8(1 != 0)
}

func ts_language_state_is_primary(tls *libc.TLS, self uintptr, state TSStateId) (r uint8) {
	if (*TSLanguage)(unsafe.Pointer(self)).Fabi_version >= uint32(14) {
		return libc.BoolUint8(libc.Int32FromUint16(state) == libc.Int32FromUint16(*(*TSStateId)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fprimary_state_ids + uintptr(state)*2))))
	} else {
		return libc.BoolUint8(1 != 0)
	}
	return r
}

func ts_language_enabled_external_tokens(tls *libc.TLS, self uintptr, external_scanner_state uint32) (r uintptr) {
	if external_scanner_state == uint32(0) {
		return libc.UintptrFromInt32(0)
	} else {
		return (*TSLanguage)(unsafe.Pointer(self)).Fexternal_scanner.Fstates + uintptr((*TSLanguage)(unsafe.Pointer(self)).Fexternal_token_count*external_scanner_state)
	}
	return r
}

func ts_language_alias_sequence(tls *libc.TLS, self uintptr, production_id uint32_t) (r uintptr) {
	var v1 uintptr
	_ = v1
	if production_id != 0 {
		v1 = (*TSLanguage)(unsafe.Pointer(self)).Falias_sequences + uintptr(production_id*uint32((*TSLanguage)(unsafe.Pointer(self)).Fmax_alias_sequence_length))*2
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	return v1
}

func ts_language_alias_at(tls *libc.TLS, self uintptr, production_id uint32_t, child_index uint32_t) (r TSSymbol) {
	var v1 int32
	_ = v1
	if production_id != 0 {
		v1 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Falias_sequences + uintptr(production_id*uint32((*TSLanguage)(unsafe.Pointer(self)).Fmax_alias_sequence_length)+child_index)*2)))
	} else {
		v1 = 0
	}
	return libc.Uint16FromInt32(v1)
}

func ts_language_field_map(tls *libc.TLS, self uintptr, production_id uint32_t, start uintptr, end uintptr) {
	var slice TSMapSlice
	_ = slice
	if (*TSLanguage)(unsafe.Pointer(self)).Ffield_count == uint32(0) {
		*(*uintptr)(unsafe.Pointer(start)) = libc.UintptrFromInt32(0)
		*(*uintptr)(unsafe.Pointer(end)) = libc.UintptrFromInt32(0)
		return
	}
	slice = *(*TSMapSlice)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Ffield_map_slices + uintptr(production_id)*4))
	*(*uintptr)(unsafe.Pointer(start)) = (*TSLanguage)(unsafe.Pointer(self)).Ffield_map_entries + uintptr(slice.Findex)*4
	*(*uintptr)(unsafe.Pointer(end)) = (*TSLanguage)(unsafe.Pointer(self)).Ffield_map_entries + uintptr(slice.Findex)*4 + uintptr(slice.Flength)*4
}

func ts_language_aliases_for_symbol(tls *libc.TLS, self uintptr, original_symbol TSSymbol, start uintptr, end uintptr) {
	var count uint16_t
	var idx, v2 uint32
	var symbol TSSymbol
	_, _, _, _ = count, idx, symbol, v2
	*(*uintptr)(unsafe.Pointer(start)) = (*TSLanguage)(unsafe.Pointer(self)).Fpublic_symbol_map + uintptr(original_symbol)*2
	*(*uintptr)(unsafe.Pointer(end)) = *(*uintptr)(unsafe.Pointer(start)) + uintptr(1)*2
	idx = uint32(0)
	for {
		v2 = idx
		idx = idx + 1
		symbol = *(*uint16_t)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Falias_map + uintptr(v2)*2))
		if libc.Int32FromUint16(symbol) == 0 || libc.Int32FromUint16(symbol) > libc.Int32FromUint16(original_symbol) {
			break
		}
		v2 = idx
		idx = idx + 1
		count = *(*uint16_t)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Falias_map + uintptr(v2)*2))
		if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(original_symbol) {
			*(*uintptr)(unsafe.Pointer(start)) = (*TSLanguage)(unsafe.Pointer(self)).Falias_map + uintptr(idx)*2
			*(*uintptr)(unsafe.Pointer(end)) = (*TSLanguage)(unsafe.Pointer(self)).Falias_map + uintptr(idx+uint32(count))*2
			break
		}
		idx = idx + uint32(count)
		goto _1
	_1:
	}
}

func ts_language_write_symbol_as_dot_string(tls *libc.TLS, self uintptr, f uintptr, symbol TSSymbol) {
	var chr, name uintptr
	_, _ = chr, name
	name = ts_language_symbol_name(tls, self, symbol)
	chr = name
	for {
		if !(*(*int8)(unsafe.Pointer(chr)) != 0) {
			break
		}
		switch int32(*(*int8)(unsafe.Pointer(chr))) {
		case int32('"'):
			fallthrough
		case int32('\\'):
			libc.Xfputc(tls, int32('\\'), f)
			libc.Xfputc(tls, int32(*(*int8)(unsafe.Pointer(chr))), f)
		case int32('\n'):
			libc.Xfputs(tls, __ccgo_ts+277, f)
		case int32('\t'):
			libc.Xfputs(tls, __ccgo_ts+280, f)
		default:
			libc.Xfputc(tls, int32(*(*int8)(unsafe.Pointer(chr))), f)
			break
		}
		goto _1
	_1:
		;
		chr = chr + 1
	}
}

func ts_range_array_add(tls *libc.TLS, self uintptr, start Length, end Length) {
	var last_range, v2 uintptr
	var range1 TSRange
	var v1 uint32_t
	_, _, _, _ = last_range, range1, v1, v2
	if (*TSRangeArray)(unsafe.Pointer(self)).Fsize > uint32(0) {
		_ = libc.Uint64FromInt64(4)
		{
			if !((*TSRangeArray)(unsafe.Pointer(self)).Fsize-libc.Uint32FromInt32(1) < (*TSRangeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+283, __ccgo_ts+327, int32(16), uintptr(unsafe.Pointer(&__func__3)))
			}
		}
		last_range = (*TSRangeArray)(unsafe.Pointer(self)).Fcontents + uintptr((*TSRangeArray)(unsafe.Pointer(self)).Fsize-uint32(1))*24
		if start.Fbytes <= (*TSRange)(unsafe.Pointer(last_range)).Fend_byte {
			(*TSRange)(unsafe.Pointer(last_range)).Fend_byte = end.Fbytes
			(*TSRange)(unsafe.Pointer(last_range)).Fend_point = end.Fextent
			return
		}
	}
	if start.Fbytes < end.Fbytes {
		range1 = TSRange{
			Fstart_point: start.Fextent,
			Fend_point:   end.Fextent,
			Fstart_byte:  start.Fbytes,
			Fend_byte:    end.Fbytes,
		}
		(*TSRangeArray)(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*TSRangeArray)(unsafe.Pointer(self)).Fcontents, (*TSRangeArray)(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(24))
		v2 = self + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*TSRange)(unsafe.Pointer((*TSRangeArray)(unsafe.Pointer(self)).Fcontents + uintptr(v1)*24)) = range1
	}
}

var __func__3 = [19]int8{'t', 's', '_', 'r', 'a', 'n', 'g', 'e', '_', 'a', 'r', 'r', 'a', 'y', '_', 'a', 'd', 'd'}

func ts_range_array_intersects(tls *libc.TLS, self uintptr, start_index uint32, start_byte uint32_t, end_byte uint32_t) (r uint8) {
	var i uint32
	var range1 uintptr
	_, _ = i, range1
	i = start_index
	for {
		if !(i < (*TSRangeArray)(unsafe.Pointer(self)).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*TSRangeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+382, __ccgo_ts+327, int32(37), uintptr(unsafe.Pointer(&__func__4)))
			}
		}
		range1 = (*TSRangeArray)(unsafe.Pointer(self)).Fcontents + uintptr(i)*24
		if (*TSRange)(unsafe.Pointer(range1)).Fend_byte > start_byte {
			if (*TSRange)(unsafe.Pointer(range1)).Fstart_byte >= end_byte {
				break
			}
			return libc.BoolUint8(1 != 0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(0 != 0)
}

var __func__4 = [26]int8{'t', 's', '_', 'r', 'a', 'n', 'g', 'e', '_', 'a', 'r', 'r', 'a', 'y', '_', 'i', 'n', 't', 'e', 'r', 's', 'e', 'c', 't', 's'}

func ts_range_array_get_changed_ranges(tls *libc.TLS, old_ranges uintptr, old_range_count uint32, new_ranges uintptr, new_range_count uint32, differences uintptr) {
	var current_position, next_new_position, next_old_position Length
	var in_new_range, in_old_range uint8
	var new_index, old_index uint32
	var new_range, old_range uintptr
	_, _, _, _, _, _, _, _, _ = current_position, in_new_range, in_old_range, new_index, new_range, next_new_position, next_old_position, old_index, old_range
	new_index = uint32(0)
	old_index = uint32(0)
	current_position = length_zero(tls)
	in_old_range = libc.BoolUint8(0 != 0)
	in_new_range = libc.BoolUint8(0 != 0)
	for old_index < old_range_count || new_index < new_range_count {
		old_range = old_ranges + uintptr(old_index)*24
		new_range = new_ranges + uintptr(new_index)*24
		if in_old_range != 0 {
			next_old_position = Length{
				Fbytes:  (*TSRange)(unsafe.Pointer(old_range)).Fend_byte,
				Fextent: (*TSRange)(unsafe.Pointer(old_range)).Fend_point,
			}
		} else {
			if old_index < old_range_count {
				next_old_position = Length{
					Fbytes:  (*TSRange)(unsafe.Pointer(old_range)).Fstart_byte,
					Fextent: (*TSRange)(unsafe.Pointer(old_range)).Fstart_point,
				}
			} else {
				next_old_position = LENGTH_MAX
			}
		}
		if in_new_range != 0 {
			next_new_position = Length{
				Fbytes:  (*TSRange)(unsafe.Pointer(new_range)).Fend_byte,
				Fextent: (*TSRange)(unsafe.Pointer(new_range)).Fend_point,
			}
		} else {
			if new_index < new_range_count {
				next_new_position = Length{
					Fbytes:  (*TSRange)(unsafe.Pointer(new_range)).Fstart_byte,
					Fextent: (*TSRange)(unsafe.Pointer(new_range)).Fstart_point,
				}
			} else {
				next_new_position = LENGTH_MAX
			}
		}
		if next_old_position.Fbytes < next_new_position.Fbytes {
			if libc.Int32FromUint8(in_old_range) != libc.Int32FromUint8(in_new_range) {
				ts_range_array_add(tls, differences, current_position, next_old_position)
			}
			if in_old_range != 0 {
				old_index = old_index + 1
			}
			current_position = next_old_position
			in_old_range = libc.BoolUint8(!(in_old_range != 0))
		} else {
			if next_new_position.Fbytes < next_old_position.Fbytes {
				if libc.Int32FromUint8(in_old_range) != libc.Int32FromUint8(in_new_range) {
					ts_range_array_add(tls, differences, current_position, next_new_position)
				}
				if in_new_range != 0 {
					new_index = new_index + 1
				}
				current_position = next_new_position
				in_new_range = libc.BoolUint8(!(in_new_range != 0))
			} else {
				if libc.Int32FromUint8(in_old_range) != libc.Int32FromUint8(in_new_range) {
					ts_range_array_add(tls, differences, current_position, next_new_position)
				}
				if in_old_range != 0 {
					old_index = old_index + 1
				}
				if in_new_range != 0 {
					new_index = new_index + 1
				}
				in_old_range = libc.BoolUint8(!(in_old_range != 0))
				in_new_range = libc.BoolUint8(!(in_new_range != 0))
				current_position = next_new_position
			}
		}
	}
}

func ts_range_edit(tls *libc.TLS, range1 uintptr, edit uintptr) {
	if (*TSRange)(unsafe.Pointer(range1)).Fend_byte >= (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_byte {
		if (*TSRange)(unsafe.Pointer(range1)).Fend_byte != uint32(4294967295) {
			(*TSRange)(unsafe.Pointer(range1)).Fend_byte = (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_byte + ((*TSRange)(unsafe.Pointer(range1)).Fend_byte - (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_byte)
			(*TSRange)(unsafe.Pointer(range1)).Fend_point = point_add(tls, (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_point, point_sub(tls, (*TSRange)(unsafe.Pointer(range1)).Fend_point, (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_point))
			if (*TSRange)(unsafe.Pointer(range1)).Fend_byte < (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_byte {
				(*TSRange)(unsafe.Pointer(range1)).Fend_byte = libc.Uint32FromUint32(4294967295)
				(*TSRange)(unsafe.Pointer(range1)).Fend_point = TSPoint{
					Frow:    libc.Uint32FromUint32(4294967295),
					Fcolumn: libc.Uint32FromUint32(4294967295),
				}
			}
		}
	} else {
		if (*TSRange)(unsafe.Pointer(range1)).Fend_byte > (*TSInputEdit)(unsafe.Pointer(edit)).Fstart_byte {
			(*TSRange)(unsafe.Pointer(range1)).Fend_byte = (*TSInputEdit)(unsafe.Pointer(edit)).Fstart_byte
			(*TSRange)(unsafe.Pointer(range1)).Fend_point = (*TSInputEdit)(unsafe.Pointer(edit)).Fstart_point
		}
	}
	if (*TSRange)(unsafe.Pointer(range1)).Fstart_byte >= (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_byte {
		(*TSRange)(unsafe.Pointer(range1)).Fstart_byte = (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_byte + ((*TSRange)(unsafe.Pointer(range1)).Fstart_byte - (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_byte)
		(*TSRange)(unsafe.Pointer(range1)).Fstart_point = point_add(tls, (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_point, point_sub(tls, (*TSRange)(unsafe.Pointer(range1)).Fstart_point, (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_point))
		if (*TSRange)(unsafe.Pointer(range1)).Fstart_byte < (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_byte {
			(*TSRange)(unsafe.Pointer(range1)).Fstart_byte = libc.Uint32FromUint32(4294967295)
			(*TSRange)(unsafe.Pointer(range1)).Fstart_point = TSPoint{
				Frow:    libc.Uint32FromUint32(4294967295),
				Fcolumn: libc.Uint32FromUint32(4294967295),
			}
		}
	} else {
		if (*TSRange)(unsafe.Pointer(range1)).Fstart_byte > (*TSInputEdit)(unsafe.Pointer(edit)).Fstart_byte {
			(*TSRange)(unsafe.Pointer(range1)).Fstart_byte = (*TSInputEdit)(unsafe.Pointer(edit)).Fstart_byte
			(*TSRange)(unsafe.Pointer(range1)).Fstart_point = (*TSInputEdit)(unsafe.Pointer(edit)).Fstart_point
		}
	}
}

type Iterator = struct {
	Fcursor              TreeCursor
	Flanguage            uintptr
	Fvisible_depth       uint32
	Fin_padding          uint8
	Fprev_external_token Subtree
}

func iterator_new(tls *libc.TLS, cursor uintptr, tree uintptr, language uintptr) (r Iterator) {
	var v1 uint32_t
	var v2 uintptr
	_, _ = v1, v2
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor + 8)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor+8)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor+8)).Fsize, cursor+8+12, uint32(1), libc.Uint64FromInt64(32))
	v2 = cursor + 8 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor+8)).Fcontents + uintptr(v1)*32)) = TreeCursorEntry{
		Fsubtree:  tree,
		Fposition: length_zero(tls),
	}
	return Iterator{
		Fcursor:              *(*TreeCursor)(unsafe.Pointer(cursor)),
		Flanguage:            language,
		Fvisible_depth:       uint32(1),
		Fprev_external_token: Subtree{},
	}
}

func iterator_done(tls *libc.TLS, self uintptr) (r uint8) {
	return libc.BoolUint8((*Iterator)(unsafe.Pointer(self)).Fcursor.Fstack.Fsize == uint32(0))
}

func iterator_start_position(tls *libc.TLS, self uintptr) (r Length) {
	var entry TreeCursorEntry
	_ = entry
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+411, __ccgo_ts+327, int32(174), uintptr(unsafe.Pointer(&__func__5)))
		}
	}
	entry = *(*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32))
	if (*Iterator)(unsafe.Pointer(self)).Fin_padding != 0 {
		return entry.Fposition
	} else {
		return length_add(tls, entry.Fposition, ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree))))
	}
	return r
}

var __func__5 = [24]int8{'i', 't', 'e', 'r', 'a', 't', 'o', 'r', '_', 's', 't', 'a', 'r', 't', '_', 'p', 'o', 's', 'i', 't', 'i', 'o', 'n'}

func iterator_end_position(tls *libc.TLS, self uintptr) (r Length) {
	var entry TreeCursorEntry
	var result Length
	_, _ = entry, result
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+411, __ccgo_ts+327, int32(183), uintptr(unsafe.Pointer(&__func__6)))
		}
	}
	entry = *(*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32))
	result = length_add(tls, entry.Fposition, ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree))))
	if (*Iterator)(unsafe.Pointer(self)).Fin_padding != 0 {
		return result
	} else {
		return length_add(tls, result, ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree))))
	}
	return r
}

var __func__6 = [22]int8{'i', 't', 'e', 'r', 'a', 't', 'o', 'r', '_', 'e', 'n', 'd', '_', 'p', 'o', 's', 'i', 't', 'i', 'o', 'n'}

func iterator_tree_is_visible(tls *libc.TLS, self uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var entry TreeCursorEntry
	var _ /* parent at bp+0 */ Subtree
	_ = entry
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+411, __ccgo_ts+327, int32(193), uintptr(unsafe.Pointer(&__func__7)))
		}
	}
	entry = *(*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32))
	if ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree))) != 0 {
		return libc.BoolUint8(1 != 0)
	}
	if (*Iterator)(unsafe.Pointer(self)).Fcursor.Fstack.Fsize > uint32(1) {
		*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
		_ = libc.Uint64FromInt64(4)
		{
			if !((*Iterator)(unsafe.Pointer(self)).Fcursor.Fstack.Fsize-libc.Uint32FromInt32(2) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+485, __ccgo_ts+327, int32(196), uintptr(unsafe.Pointer(&__func__7)))
			}
		}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp)) = *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr((*Iterator)(unsafe.Pointer(self)).Fcursor.Fstack.Fsize-uint32(2))*32)).Fsubtree))
		return libc.BoolUint8(libc.Int32FromUint16(ts_language_alias_at(tls, (*Iterator)(unsafe.Pointer(self)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), entry.Fstructural_child_index)) != 0)
	}
	return libc.BoolUint8(0 != 0)
}

var __func__7 = [25]int8{'i', 't', 'e', 'r', 'a', 't', 'o', 'r', '_', 't', 'r', 'e', 'e', '_', 'i', 's', '_', 'v', 'i', 's', 'i', 'b', 'l', 'e'}

func iterator_get_visible_state(tls *libc.TLS, self uintptr, tree uintptr, alias_symbol uintptr, start_byte uintptr) {
	var entry TreeCursorEntry
	var i uint32_t
	var parent uintptr
	_, _, _ = entry, i, parent
	i = (*Iterator)(unsafe.Pointer(self)).Fcursor.Fstack.Fsize - uint32(1)
	if (*Iterator)(unsafe.Pointer(self)).Fin_padding != 0 {
		if i == uint32(0) {
			return
		}
		i = i - 1
	}
	for {
		if !(i+uint32(1) > uint32(0)) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+555, __ccgo_ts+327, int32(220), uintptr(unsafe.Pointer(&__func__8)))
			}
		}
		entry = *(*TreeCursorEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(i)*32))
		if i > uint32(0) {
			_ = libc.Uint64FromInt64(4)
			{
				if !(i-libc.Uint32FromInt32(1) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+599, __ccgo_ts+327, int32(223), uintptr(unsafe.Pointer(&__func__8)))
				}
			}
			parent = (*TreeCursorEntry)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents + uintptr(i-uint32(1))*32)).Fsubtree
			*(*TSSymbol)(unsafe.Pointer(alias_symbol)) = ts_language_alias_at(tls, (*Iterator)(unsafe.Pointer(self)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(parent)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), entry.Fstructural_child_index)
		}
		if ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree))) != 0 || *(*TSSymbol)(unsafe.Pointer(alias_symbol)) != 0 {
			*(*Subtree)(unsafe.Pointer(tree)) = *(*Subtree)(unsafe.Pointer(entry.Fsubtree))
			*(*uint32_t)(unsafe.Pointer(start_byte)) = entry.Fposition.Fbytes
			break
		}
		goto _1
	_1:
		;
		i = i - 1
	}
}

var __func__8 = [27]int8{'i', 't', 'e', 'r', 'a', 't', 'o', 'r', '_', 'g', 'e', 't', '_', 'v', 'i', 's', 'i', 'b', 'l', 'e', '_', 's', 't', 'a', 't', 'e'}

func iterator_ascend(tls *libc.TLS, self uintptr) {
	if iterator_done(tls, self) != 0 {
		return
	}
	if iterator_tree_is_visible(tls, self) != 0 && !((*Iterator)(unsafe.Pointer(self)).Fin_padding != 0) {
		(*Iterator)(unsafe.Pointer(self)).Fvisible_depth = (*Iterator)(unsafe.Pointer(self)).Fvisible_depth - 1
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+411, __ccgo_ts+327, int32(242), uintptr(unsafe.Pointer(&__func__9)))
		}
	}
	if (*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32)).Fchild_index > uint32(0) {
		(*Iterator)(unsafe.Pointer(self)).Fin_padding = libc.BoolUint8(0 != 0)
	}
	(*Iterator)(unsafe.Pointer(self)).Fcursor.Fstack.Fsize = (*Iterator)(unsafe.Pointer(self)).Fcursor.Fstack.Fsize - 1
}

var __func__9 = [16]int8{'i', 't', 'e', 'r', 'a', 't', 'o', 'r', '_', 'a', 's', 'c', 'e', 'n', 'd'}

func iterator_descend(tls *libc.TLS, self uintptr, goal_position uint32_t) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var child, v2 uintptr
	var child_left, child_right, position Length
	var did_descend uint8
	var entry TreeCursorEntry
	var i, n, structural_child_index, v3 uint32_t
	var _ /* last_external_token at bp+0 */ Subtree
	_, _, _, _, _, _, _, _, _, _, _ = child, child_left, child_right, did_descend, entry, i, n, position, structural_child_index, v2, v3
	if (*Iterator)(unsafe.Pointer(self)).Fin_padding != 0 {
		return libc.BoolUint8(0 != 0)
	}
	did_descend = libc.BoolUint8(0 != 0)
	for cond := true; cond; cond = did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		_ = libc.Uint64FromInt64(4)
		{
			if !((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+411, __ccgo_ts+327, int32(252), uintptr(unsafe.Pointer(&__func__10)))
			}
		}
		entry = *(*TreeCursorEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32))
		position = entry.Fposition
		structural_child_index = uint32(0)
		i = uint32(0)
		n = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree)))
		for {
			if !(i < n) {
				break
			}
			if int32(*(*uint8)(unsafe.Pointer(entry.Fsubtree + 0))&0x1>>0) != 0 {
				v2 = libc.UintptrFromInt32(0)
			} else {
				v2 = *(*uintptr)(unsafe.Pointer(entry.Fsubtree)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(entry.Fsubtree)))).Fchild_count)*8
			}
			child = v2 + uintptr(i)*8
			child_left = length_add(tls, position, ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(child))))
			child_right = length_add(tls, child_left, ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(child))))
			if child_right.Fbytes > goal_position {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
				v2 = self + 8 + 8
				v3 = *(*uint32_t)(unsafe.Pointer(v2))
				*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
				*(*TreeCursorEntry)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents + uintptr(v3)*32)) = TreeCursorEntry{
					Fsubtree:                child,
					Fposition:               position,
					Fchild_index:            i,
					Fstructural_child_index: structural_child_index,
				}
				if iterator_tree_is_visible(tls, self) != 0 {
					if child_left.Fbytes > goal_position {
						(*Iterator)(unsafe.Pointer(self)).Fin_padding = libc.BoolUint8(1 != 0)
					} else {
						(*Iterator)(unsafe.Pointer(self)).Fvisible_depth = (*Iterator)(unsafe.Pointer(self)).Fvisible_depth + 1
					}
					return libc.BoolUint8(1 != 0)
				}
				did_descend = libc.BoolUint8(1 != 0)
				break
			}
			position = child_right
			if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(child))) != 0) {
				structural_child_index = structural_child_index + 1
			}
			*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
			*(*struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			})(unsafe.Pointer(bp)) = ts_subtree_last_external_token(tls, *(*Subtree)(unsafe.Pointer(child)))
			if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
				(*Iterator)(unsafe.Pointer(self)).Fprev_external_token = *(*Subtree)(unsafe.Pointer(bp))
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	return libc.BoolUint8(0 != 0)
}

var __func__10 = [17]int8{'i', 't', 'e', 'r', 'a', 't', 'o', 'r', '_', 'd', 'e', 's', 'c', 'e', 'n', 'd'}

func iterator_advance(tls *libc.TLS, self uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var child_index, structural_child_index, v2 uint32_t
	var entry TreeCursorEntry
	var next_child, parent, v3 uintptr
	var position Length
	var _ /* last_external_token at bp+0 */ Subtree
	_, _, _, _, _, _, _, _ = child_index, entry, next_child, parent, position, structural_child_index, v2, v3
	if (*Iterator)(unsafe.Pointer(self)).Fin_padding != 0 {
		(*Iterator)(unsafe.Pointer(self)).Fin_padding = libc.BoolUint8(0 != 0)
		if iterator_tree_is_visible(tls, self) != 0 {
			(*Iterator)(unsafe.Pointer(self)).Fvisible_depth = (*Iterator)(unsafe.Pointer(self)).Fvisible_depth + 1
		} else {
			iterator_descend(tls, self, uint32(0))
		}
		return
	}
	for {
		if iterator_tree_is_visible(tls, self) != 0 {
			(*Iterator)(unsafe.Pointer(self)).Fvisible_depth = (*Iterator)(unsafe.Pointer(self)).Fvisible_depth - 1
		}
		v3 = self + 8 + 8
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) - 1
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		entry = *(*TreeCursorEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(v2)*32))
		if iterator_done(tls, self) != 0 {
			return
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+411, __ccgo_ts+327, int32(309), uintptr(unsafe.Pointer(&__func__11)))
			}
		}
		parent = (*TreeCursorEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32)).Fsubtree
		child_index = entry.Fchild_index + uint32(1)
		*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp)) = ts_subtree_last_external_token(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree)))
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			(*Iterator)(unsafe.Pointer(self)).Fprev_external_token = *(*Subtree)(unsafe.Pointer(bp))
		}
		if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(parent))) > child_index {
			position = length_add(tls, entry.Fposition, ts_subtree_total_size(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree))))
			structural_child_index = entry.Fstructural_child_index
			if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(entry.Fsubtree))) != 0) {
				structural_child_index = structural_child_index + 1
			}
			if int32(*(*uint8)(unsafe.Pointer(parent + 0))&0x1>>0) != 0 {
				v3 = libc.UintptrFromInt32(0)
			} else {
				v3 = *(*uintptr)(unsafe.Pointer(parent)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(parent)))).Fchild_count)*8
			}
			next_child = v3 + uintptr(child_index)*8
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
			v3 = self + 8 + 8
			v2 = *(*uint32_t)(unsafe.Pointer(v3))
			*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
			*(*TreeCursorEntry)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents + uintptr(v2)*32)) = TreeCursorEntry{
				Fsubtree:                next_child,
				Fposition:               position,
				Fchild_index:            child_index,
				Fstructural_child_index: structural_child_index,
			}
			if iterator_tree_is_visible(tls, self) != 0 {
				if ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(next_child))).Fbytes > uint32(0) {
					(*Iterator)(unsafe.Pointer(self)).Fin_padding = libc.BoolUint8(1 != 0)
				} else {
					(*Iterator)(unsafe.Pointer(self)).Fvisible_depth = (*Iterator)(unsafe.Pointer(self)).Fvisible_depth + 1
				}
			} else {
				iterator_descend(tls, self, uint32(0))
			}
			break
		}
		goto _1
	_1:
	}
}

var __func__11 = [17]int8{'i', 't', 'e', 'r', 'a', 't', 'o', 'r', '_', 'a', 'd', 'v', 'a', 'n', 'c', 'e'}

type IteratorComparison = int32

const IteratorDiffers = 0
const IteratorMayDiffer = 1
const IteratorMatches = 2

func iterator_compare(tls *libc.TLS, old_iter uintptr, new_iter uintptr) (r IteratorComparison) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var new_error_cost, new_size, old_error_cost, old_size uint32_t
	var new_has_external_tokens, old_has_external_tokens uint8
	var new_state, old_state TSStateId
	var new_symbol, old_symbol TSSymbol
	var _ /* new_alias_symbol at bp+26 */ TSSymbol
	var _ /* new_start at bp+20 */ uint32_t
	var _ /* new_tree at bp+8 */ Subtree
	var _ /* old_alias_symbol at bp+24 */ TSSymbol
	var _ /* old_start at bp+16 */ uint32_t
	var _ /* old_tree at bp+0 */ Subtree
	_, _, _, _, _, _, _, _, _, _ = new_error_cost, new_has_external_tokens, new_size, new_state, new_symbol, old_error_cost, old_has_external_tokens, old_size, old_state, old_symbol
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = Subtree{}
	*(*Subtree)(unsafe.Pointer(bp + 8)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp + 8)) = Subtree{}
	*(*uint32_t)(unsafe.Pointer(bp + 16)) = uint32(0)
	*(*uint32_t)(unsafe.Pointer(bp + 20)) = uint32(0)
	*(*TSSymbol)(unsafe.Pointer(bp + 24)) = uint16(0)
	*(*TSSymbol)(unsafe.Pointer(bp + 26)) = uint16(0)
	iterator_get_visible_state(tls, old_iter, bp, bp+24, bp+16)
	iterator_get_visible_state(tls, new_iter, bp+8, bp+26, bp+20)
	old_symbol = ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))
	new_symbol = ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) && !(*(*uintptr)(unsafe.Pointer(bp + 8)) != 0) {
		return int32(IteratorMatches)
	}
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) || !(*(*uintptr)(unsafe.Pointer(bp + 8)) != 0) {
		return int32(IteratorDiffers)
	}
	if libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(bp + 24))) != libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(bp + 26))) || libc.Int32FromUint16(old_symbol) != libc.Int32FromUint16(new_symbol) {
		return int32(IteratorDiffers)
	}
	old_size = ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(bp))).Fbytes
	new_size = ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(bp + 8))).Fbytes
	old_state = ts_subtree_parse_state(tls, *(*Subtree)(unsafe.Pointer(bp)))
	new_state = ts_subtree_parse_state(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
	old_has_external_tokens = ts_subtree_has_external_tokens(tls, *(*Subtree)(unsafe.Pointer(bp)))
	new_has_external_tokens = ts_subtree_has_external_tokens(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
	old_error_cost = ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp)))
	new_error_cost = ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
	if *(*uint32_t)(unsafe.Pointer(bp + 16)) != *(*uint32_t)(unsafe.Pointer(bp + 20)) || libc.Int32FromUint16(old_symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) || old_size != new_size || libc.Int32FromUint16(old_state) == libc.Int32FromInt32(0x7fff)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1) || libc.Int32FromUint16(new_state) == libc.Int32FromInt32(0x7fff)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1) || libc.BoolInt32(libc.Int32FromUint16(old_state) == 0) != libc.BoolInt32(libc.Int32FromUint16(new_state) == 0) || old_error_cost != new_error_cost || libc.Int32FromUint8(old_has_external_tokens) != libc.Int32FromUint8(new_has_external_tokens) || ts_subtree_has_changes(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 || old_has_external_tokens != 0 && !(ts_subtree_external_scanner_state_eq(tls, (*Iterator)(unsafe.Pointer(old_iter)).Fprev_external_token, (*Iterator)(unsafe.Pointer(new_iter)).Fprev_external_token) != 0) {
		return int32(IteratorMayDiffer)
	}
	return int32(IteratorMatches)
}

func ts_subtree_get_changed_ranges(tls *libc.TLS, old_tree uintptr, new_tree uintptr, cursor1 uintptr, cursor2 uintptr, language uintptr, included_range_differences uintptr, ranges uintptr) (r uint32) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var comparison IteratorComparison
	var included_range_difference_index uint32
	var is_changed uint8
	var new_size, next_position, old_size, position Length
	var range1 uintptr
	var _ /* new_iter at bp+72 */ Iterator
	var _ /* old_iter at bp+16 */ Iterator
	var _ /* results at bp+0 */ TSRangeArray
	_, _, _, _, _, _, _, _ = comparison, included_range_difference_index, is_changed, new_size, next_position, old_size, position, range1
	*(*TSRangeArray)(unsafe.Pointer(bp)) = TSRangeArray{}
	*(*Iterator)(unsafe.Pointer(bp + 16)) = Iterator{}
	*(*struct {
		Fcursor struct {
			Ftree  uintptr
			Fstack struct {
				Fcontents uintptr
				Fsize     uint32
				Fcapacity uint32
			}
			Froot_alias_symbol uint16
		}
		Flanguage            uintptr
		Fvisible_depth       uint32
		Fin_padding          uint8
		Fprev_external_token struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		}
	})(unsafe.Pointer(bp + 16)) = iterator_new(tls, cursor1, old_tree, language)
	*(*Iterator)(unsafe.Pointer(bp + 72)) = Iterator{}
	*(*struct {
		Fcursor struct {
			Ftree  uintptr
			Fstack struct {
				Fcontents uintptr
				Fsize     uint32
				Fcapacity uint32
			}
			Froot_alias_symbol uint16
		}
		Flanguage            uintptr
		Fvisible_depth       uint32
		Fin_padding          uint8
		Fprev_external_token struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		}
	})(unsafe.Pointer(bp + 72)) = iterator_new(tls, cursor2, new_tree, language)
	included_range_difference_index = uint32(0)
	position = iterator_start_position(tls, bp+16)
	next_position = iterator_start_position(tls, bp+72)
	if position.Fbytes < next_position.Fbytes {
		ts_range_array_add(tls, bp, position, next_position)
		position = next_position
	} else {
		if position.Fbytes > next_position.Fbytes {
			ts_range_array_add(tls, bp, next_position, position)
			next_position = position
		}
	}
	for cond := true; cond; cond = !(iterator_done(tls, bp+16) != 0) && !(iterator_done(tls, bp+72) != 0) {
		comparison = iterator_compare(tls, bp+16, bp+72)
		if comparison == int32(IteratorMatches) && ts_range_array_intersects(tls, included_range_differences, included_range_difference_index, position.Fbytes, iterator_end_position(tls, bp+16).Fbytes) != 0 {
			comparison = int32(IteratorMayDiffer)
		}
		is_changed = libc.BoolUint8(0 != 0)
		switch comparison {
		case int32(IteratorMatches):
			next_position = iterator_end_position(tls, bp+16)
		case int32(IteratorMayDiffer):
			if iterator_descend(tls, bp+16, position.Fbytes) != 0 {
				if !(iterator_descend(tls, bp+72, position.Fbytes) != 0) {
					is_changed = libc.BoolUint8(1 != 0)
					next_position = iterator_end_position(tls, bp+16)
				}
			} else {
				if iterator_descend(tls, bp+72, position.Fbytes) != 0 {
					is_changed = libc.BoolUint8(1 != 0)
					next_position = iterator_end_position(tls, bp+72)
				} else {
					next_position = length_min(tls, iterator_end_position(tls, bp+16), iterator_end_position(tls, bp+72))
				}
			}
		case int32(IteratorDiffers):
			is_changed = libc.BoolUint8(1 != 0)
			next_position = length_min(tls, iterator_end_position(tls, bp+16), iterator_end_position(tls, bp+72))
			break
		}
		for !(iterator_done(tls, bp+16) != 0) && iterator_end_position(tls, bp+16).Fbytes <= next_position.Fbytes {
			iterator_advance(tls, bp+16)
		}
		for !(iterator_done(tls, bp+72) != 0) && iterator_end_position(tls, bp+72).Fbytes <= next_position.Fbytes {
			iterator_advance(tls, bp+72)
		}
		for (*(*Iterator)(unsafe.Pointer(bp + 16))).Fvisible_depth > (*(*Iterator)(unsafe.Pointer(bp + 72))).Fvisible_depth {
			iterator_ascend(tls, bp+16)
		}
		for (*(*Iterator)(unsafe.Pointer(bp + 72))).Fvisible_depth > (*(*Iterator)(unsafe.Pointer(bp + 16))).Fvisible_depth {
			iterator_ascend(tls, bp+72)
		}
		if is_changed != 0 {
			ts_range_array_add(tls, bp, position, next_position)
		}
		position = next_position
		for included_range_difference_index < (*TSRangeArray)(unsafe.Pointer(included_range_differences)).Fsize {
			_ = libc.Uint64FromInt64(4)
			{
				if !(included_range_difference_index < (*TSRangeArray)(unsafe.Pointer(included_range_differences)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+647, __ccgo_ts+327, int32(534), uintptr(unsafe.Pointer(&__func__12)))
				}
			}
			range1 = (*TSRangeArray)(unsafe.Pointer(included_range_differences)).Fcontents + uintptr(included_range_difference_index)*24
			if (*TSRange)(unsafe.Pointer(range1)).Fend_byte <= position.Fbytes {
				included_range_difference_index = included_range_difference_index + 1
			} else {
				break
			}
		}
	}
	old_size = ts_subtree_total_size(tls, *(*Subtree)(unsafe.Pointer(old_tree)))
	new_size = ts_subtree_total_size(tls, *(*Subtree)(unsafe.Pointer(new_tree)))
	if old_size.Fbytes < new_size.Fbytes {
		ts_range_array_add(tls, bp, old_size, new_size)
	} else {
		if new_size.Fbytes < old_size.Fbytes {
			ts_range_array_add(tls, bp, new_size, old_size)
		}
	}
	*(*TreeCursor)(unsafe.Pointer(cursor1)) = (*(*Iterator)(unsafe.Pointer(bp + 16))).Fcursor
	*(*TreeCursor)(unsafe.Pointer(cursor2)) = (*(*Iterator)(unsafe.Pointer(bp + 72))).Fcursor
	*(*uintptr)(unsafe.Pointer(ranges)) = (*(*TSRangeArray)(unsafe.Pointer(bp))).Fcontents
	return (*(*TSRangeArray)(unsafe.Pointer(bp))).Fsize
}

var __func__12 = [30]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'g', 'e', 't', '_', 'c', 'h', 'a', 'n', 'g', 'e', 'd', '_', 'r', 'a', 'n', 'g', 'e', 's'}

func ts_language_copy(tls *libc.TLS, self uintptr) (r uintptr) {
	if self != 0 && ts_language_is_wasm(tls, self) != 0 {
		ts_wasm_language_retain(tls, self)
	}
	return self
}

func ts_language_delete(tls *libc.TLS, self uintptr) {
	if self != 0 && ts_language_is_wasm(tls, self) != 0 {
		ts_wasm_language_release(tls, self)
	}
}

func ts_language_symbol_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSLanguage)(unsafe.Pointer(self)).Fsymbol_count + (*TSLanguage)(unsafe.Pointer(self)).Falias_count
}

func ts_language_state_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSLanguage)(unsafe.Pointer(self)).Fstate_count
}

func ts_language_supertypes(tls *libc.TLS, self uintptr, length uintptr) (r uintptr) {
	if (*TSLanguage)(unsafe.Pointer(self)).Fabi_version >= uint32(15) {
		*(*uint32_t)(unsafe.Pointer(length)) = (*TSLanguage)(unsafe.Pointer(self)).Fsupertype_count
		return (*TSLanguage)(unsafe.Pointer(self)).Fsupertype_symbols
	} else {
		*(*uint32_t)(unsafe.Pointer(length)) = uint32(0)
		return libc.UintptrFromInt32(0)
	}
	return r
}

func ts_language_subtypes(tls *libc.TLS, self uintptr, supertype TSSymbol, length uintptr) (r uintptr) {
	var slice TSMapSlice
	_ = slice
	if (*TSLanguage)(unsafe.Pointer(self)).Fabi_version < uint32(15) || !(ts_language_symbol_metadata(tls, self, supertype).Fsupertype != 0) {
		*(*uint32_t)(unsafe.Pointer(length)) = uint32(0)
		return libc.UintptrFromInt32(0)
	}
	slice = *(*TSMapSlice)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fsupertype_map_slices + uintptr(supertype)*4))
	*(*uint32_t)(unsafe.Pointer(length)) = uint32(slice.Flength)
	return (*TSLanguage)(unsafe.Pointer(self)).Fsupertype_map_entries + uintptr(slice.Findex)*2
}

func ts_language_abi_version(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSLanguage)(unsafe.Pointer(self)).Fabi_version
}

func ts_language_metadata(tls *libc.TLS, self uintptr) (r uintptr) {
	var v1 uintptr
	_ = v1
	if (*TSLanguage)(unsafe.Pointer(self)).Fabi_version >= uint32(15) {
		v1 = self + 280
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	return v1
}

func ts_language_name(tls *libc.TLS, self uintptr) (r uintptr) {
	var v1 uintptr
	_ = v1
	if (*TSLanguage)(unsafe.Pointer(self)).Fabi_version >= uint32(15) {
		v1 = (*TSLanguage)(unsafe.Pointer(self)).Fname
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	return v1
}

func ts_language_field_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSLanguage)(unsafe.Pointer(self)).Ffield_count
}

func ts_language_table_entry(tls *libc.TLS, self uintptr, state TSStateId, symbol TSSymbol, result uintptr) {
	var action_index uint32_t
	var entry uintptr
	_, _ = action_index, entry
	if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) || libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
		(*TableEntry)(unsafe.Pointer(result)).Faction_count = uint32(0)
		(*TableEntry)(unsafe.Pointer(result)).Fis_reusable = libc.BoolUint8(0 != 0)
		(*TableEntry)(unsafe.Pointer(result)).Factions = libc.UintptrFromInt32(0)
	} else {
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(symbol) < (*TSLanguage)(unsafe.Pointer(self)).Ftoken_count) {
				libc.X__assert_fail(tls, __ccgo_ts+728, __ccgo_ts+755, int32(79), uintptr(unsafe.Pointer(&__func__13)))
			}
		}
		action_index = uint32(ts_language_lookup(tls, self, state, symbol))
		entry = (*TSLanguage)(unsafe.Pointer(self)).Fparse_actions + uintptr(action_index)*8
		(*TableEntry)(unsafe.Pointer(result)).Faction_count = uint32((*(*struct {
			Fcount    uint8_t
			Freusable uint8
		})(unsafe.Pointer(entry))).Fcount)
		(*TableEntry)(unsafe.Pointer(result)).Fis_reusable = (*(*struct {
			Fcount    uint8_t
			Freusable uint8
		})(unsafe.Pointer(entry))).Freusable
		(*TableEntry)(unsafe.Pointer(result)).Factions = entry + libc.UintptrFromInt32(1)*8
	}
}

var __func__13 = [24]int8{'t', 's', '_', 'l', 'a', 'n', 'g', 'u', 'a', 'g', 'e', '_', 't', 'a', 'b', 'l', 'e', '_', 'e', 'n', 't', 'r', 'y'}

func ts_language_lex_mode_for_state(tls *libc.TLS, self uintptr, state TSStateId) (r TSLexerMode) {
	var mode TSLexMode
	_ = mode
	if (*TSLanguage)(unsafe.Pointer(self)).Fabi_version < uint32(15) {
		mode = *(*TSLexMode)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Flex_modes + uintptr(state)*4))
		return TSLexerMode{
			Flex_state:          mode.Flex_state,
			Fexternal_lex_state: mode.Fexternal_lex_state,
		}
	} else {
		return *(*TSLexerMode)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Flex_modes + uintptr(state)*6))
	}
	return r
}

func ts_language_is_reserved_word(tls *libc.TLS, self uintptr, state TSStateId, symbol TSSymbol) (r uint8) {
	var end, i, start uint32
	var lex_mode TSLexerMode
	_, _, _, _ = end, i, lex_mode, start
	lex_mode = ts_language_lex_mode_for_state(tls, self, state)
	if libc.Int32FromUint16(lex_mode.Freserved_word_set_id) > 0 {
		start = libc.Uint32FromInt32(libc.Int32FromUint16(lex_mode.Freserved_word_set_id) * libc.Int32FromUint16((*TSLanguage)(unsafe.Pointer(self)).Fmax_reserved_word_set_size))
		end = start + uint32((*TSLanguage)(unsafe.Pointer(self)).Fmax_reserved_word_set_size)
		i = start
		for {
			if !(i < end) {
				break
			}
			if libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Freserved_words + uintptr(i)*2))) == libc.Int32FromUint16(symbol) {
				return libc.BoolUint8(1 != 0)
			}
			if libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Freserved_words + uintptr(i)*2))) == 0 {
				break
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	return libc.BoolUint8(0 != 0)
}

func ts_language_symbol_metadata(tls *libc.TLS, self uintptr, symbol TSSymbol) (r TSSymbolMetadata) {
	if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) {
		return TSSymbolMetadata{
			Fvisible: libc.BoolUint8(1 != 0),
			Fnamed:   libc.BoolUint8(1 != 0),
		}
	} else {
		if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
			return TSSymbolMetadata{}
		} else {
			return *(*TSSymbolMetadata)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fsymbol_metadata + uintptr(symbol)*3))
		}
	}
	return r
}

func ts_language_public_symbol(tls *libc.TLS, self uintptr, symbol TSSymbol) (r TSSymbol) {
	if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) {
		return symbol
	}
	return *(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fpublic_symbol_map + uintptr(symbol)*2))
}

func ts_language_next_state(tls *libc.TLS, self uintptr, state TSStateId, symbol TSSymbol) (r TSStateId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var action TSParseAction
	var actions uintptr
	var v1 int32
	var _ /* count at bp+0 */ uint32_t
	_, _, _ = action, actions, v1
	if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) || libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
		return uint16(0)
	} else {
		if uint32(symbol) < (*TSLanguage)(unsafe.Pointer(self)).Ftoken_count {
			actions = ts_language_actions(tls, self, state, symbol, bp)
			if *(*uint32_t)(unsafe.Pointer(bp)) > uint32(0) {
				action = *(*TSParseAction)(unsafe.Pointer(actions + uintptr(*(*uint32_t)(unsafe.Pointer(bp))-uint32(1))*8))
				if libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(&action))) == int32(TSParseActionTypeShift) {
					if action.Fshift.Fextra != 0 {
						v1 = libc.Int32FromUint16(state)
					} else {
						v1 = libc.Int32FromUint16(action.Fshift.Fstate)
					}
					return libc.Uint16FromInt32(v1)
				}
			}
			return uint16(0)
		} else {
			return ts_language_lookup(tls, self, state, symbol)
		}
	}
	return r
}

func ts_language_symbol_name(tls *libc.TLS, self uintptr, symbol TSSymbol) (r uintptr) {
	if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) {
		return __ccgo_ts + 800
	} else {
		if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
			return __ccgo_ts + 806
		} else {
			if uint32(symbol) < ts_language_symbol_count(tls, self) {
				return *(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fsymbol_names + uintptr(symbol)*8))
			} else {
				return libc.UintptrFromInt32(0)
			}
		}
	}
	return r
}

func ts_language_symbol_for_name(tls *libc.TLS, self uintptr, string1 uintptr, length uint32_t, is_named uint8) (r TSSymbol) {
	var count uint16_t
	var i TSSymbol
	var metadata TSSymbolMetadata
	var symbol_name uintptr
	_, _, _, _ = count, i, metadata, symbol_name
	if is_named != 0 && !(libc.Xstrncmp(tls, string1, __ccgo_ts+800, uint64(length)) != 0) {
		return libc.Uint16FromInt32(-libc.Int32FromInt32(1))
	}
	count = uint16(ts_language_symbol_count(tls, self))
	i = uint16(0)
	for {
		if !(libc.Int32FromUint16(i) < libc.Int32FromUint16(count)) {
			break
		}
		metadata = ts_language_symbol_metadata(tls, self, i)
		if !(metadata.Fvisible != 0) && !(metadata.Fsupertype != 0) || libc.Int32FromUint8(metadata.Fnamed) != libc.Int32FromUint8(is_named) {
			goto _1
		}
		symbol_name = *(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fsymbol_names + uintptr(i)*8))
		if !(libc.Xstrncmp(tls, symbol_name, string1, uint64(length)) != 0) && !(*(*int8)(unsafe.Pointer(symbol_name + uintptr(length))) != 0) {
			return *(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Fpublic_symbol_map + uintptr(i)*2))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return uint16(0)
}

func ts_language_symbol_type(tls *libc.TLS, self uintptr, symbol TSSymbol) (r TSSymbolType1) {
	var metadata TSSymbolMetadata
	_ = metadata
	metadata = ts_language_symbol_metadata(tls, self, symbol)
	if metadata.Fnamed != 0 && metadata.Fvisible != 0 {
		return int32(TSSymbolTypeRegular)
	} else {
		if metadata.Fvisible != 0 {
			return int32(TSSymbolTypeAnonymous)
		} else {
			if metadata.Fsupertype != 0 {
				return int32(TSSymbolTypeSupertype)
			} else {
				return int32(TSSymbolTypeAuxiliary)
			}
		}
	}
	return r
}

func ts_language_field_name_for_id(tls *libc.TLS, self uintptr, id TSFieldId) (r uintptr) {
	var count uint32_t
	_ = count
	count = ts_language_field_count(tls, self)
	if count != 0 && uint32(id) <= count {
		return *(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Ffield_names + uintptr(id)*8))
	} else {
		return libc.UintptrFromInt32(0)
	}
	return r
}

func ts_language_field_id_for_name(tls *libc.TLS, self uintptr, name uintptr, name_length uint32_t) (r TSFieldId) {
	var count uint16_t
	var i TSSymbol
	_, _ = count, i
	count = uint16(ts_language_field_count(tls, self))
	i = uint16(1)
	for {
		if !(libc.Int32FromUint16(i) < libc.Int32FromUint16(count)+int32(1)) {
			break
		}
		switch libc.Xstrncmp(tls, name, *(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Ffield_names + uintptr(i)*8)), uint64(name_length)) {
		case 0:
			if int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(self)).Ffield_names + uintptr(i)*8)) + uintptr(name_length)))) == 0 {
				return i
			}
		case -int32(1):
			return uint16(0)
		default:
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return uint16(0)
}

func ts_lookahead_iterator_new(tls *libc.TLS, self uintptr, state TSStateId) (r uintptr) {
	var iterator uintptr
	_ = iterator
	if uint32(state) >= (*TSLanguage)(unsafe.Pointer(self)).Fstate_count {
		return libc.UintptrFromInt32(0)
	}
	iterator = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(56))
	*(*LookaheadIterator)(unsafe.Pointer(iterator)) = ts_language_lookaheads(tls, self, state)
	return iterator
}

func ts_lookahead_iterator_delete(tls *libc.TLS, self uintptr) {
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, self)
}

func ts_lookahead_iterator_reset_state(tls *libc.TLS, self uintptr, state TSStateId) (r uint8) {
	var iterator uintptr
	_ = iterator
	iterator = self
	if uint32(state) >= (*TSLanguage)(unsafe.Pointer((*LookaheadIterator)(unsafe.Pointer(iterator)).Flanguage)).Fstate_count {
		return libc.BoolUint8(0 != 0)
	}
	*(*LookaheadIterator)(unsafe.Pointer(iterator)) = ts_language_lookaheads(tls, (*LookaheadIterator)(unsafe.Pointer(iterator)).Flanguage, state)
	return libc.BoolUint8(1 != 0)
}

func ts_lookahead_iterator_language(tls *libc.TLS, self uintptr) (r uintptr) {
	var iterator uintptr
	_ = iterator
	iterator = self
	return (*LookaheadIterator)(unsafe.Pointer(iterator)).Flanguage
}

func ts_lookahead_iterator_reset(tls *libc.TLS, self uintptr, language uintptr, state TSStateId) (r uint8) {
	var iterator uintptr
	_ = iterator
	if uint32(state) >= (*TSLanguage)(unsafe.Pointer(language)).Fstate_count {
		return libc.BoolUint8(0 != 0)
	}
	iterator = self
	*(*LookaheadIterator)(unsafe.Pointer(iterator)) = ts_language_lookaheads(tls, language, state)
	return libc.BoolUint8(1 != 0)
}

func ts_lookahead_iterator_next(tls *libc.TLS, self uintptr) (r uint8) {
	var iterator uintptr
	_ = iterator
	iterator = self
	return ts_lookahead_iterator__next(tls, iterator)
}

func ts_lookahead_iterator_current_symbol(tls *libc.TLS, self uintptr) (r TSSymbol) {
	var iterator uintptr
	_ = iterator
	iterator = self
	return (*LookaheadIterator)(unsafe.Pointer(iterator)).Fsymbol
}

func ts_lookahead_iterator_current_symbol_name(tls *libc.TLS, self uintptr) (r uintptr) {
	var iterator uintptr
	_ = iterator
	iterator = self
	return ts_language_symbol_name(tls, (*LookaheadIterator)(unsafe.Pointer(iterator)).Flanguage, (*LookaheadIterator)(unsafe.Pointer(iterator)).Fsymbol)
}

type ColumnData = struct {
	Fvalue uint32_t
	Fvalid uint8
}

type Lexer = struct {
	Fdata                         TSLexer
	Fcurrent_position             Length
	Ftoken_start_position         Length
	Ftoken_end_position           Length
	Fincluded_ranges              uintptr
	Fchunk                        uintptr
	Finput                        TSInput
	Flogger                       TSLogger
	Fincluded_range_count         uint32_t
	Fcurrent_included_range_index uint32_t
	Fchunk_start                  uint32_t
	Fchunk_size                   uint32_t
	Flookahead_size               uint32_t
	Fdid_get_column               uint8
	Fcolumn_data                  ColumnData
	Fdebug_buffer                 [1024]int8
}

type ptrdiff_t = int64

type max_align_t = struct {
	F__max_align_ll int64
	F__max_align_ld float64
}

type UBool = int8

type UChar = uint16

type OldUChar = uint16

type UChar32 = int32

var TS_DECODE_ERROR = -libc.Int32FromInt32(1)

func ts_decode_utf8(tls *libc.TLS, string1 uintptr, length uint32_t, code_point uintptr) (r uint32_t) {
	var __t, v12, v18, v5, v8 uint8_t
	var i, v1, v10 uint32_t
	var v11, v13, v15, v16, v17, v19, v20, v6, v9 bool
	var v2, v3 int32
	var v4, v7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = __t, i, v1, v10, v11, v12, v13, v15, v16, v17, v18, v19, v2, v20, v3, v4, v5, v6, v7, v8, v9
	i = uint32(0)
	v1 = i
	i = i + 1
	*(*int32_t)(unsafe.Pointer(code_point)) = libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(string1 + uintptr(v1))))
	if !(*(*int32_t)(unsafe.Pointer(code_point))&libc.Int32FromInt32(0x80) == libc.Int32FromInt32(0)) {
		__t = uint8(0)
		if v17 = i != length; v17 {
			if *(*int32_t)(unsafe.Pointer(code_point)) >= int32(0xe0) {
				if *(*int32_t)(unsafe.Pointer(code_point)) < int32(0xf0) {
					v4 = code_point
					*(*int32_t)(unsafe.Pointer(v4)) &= int32(0xf)
					v5 = *(*uint8_t)(unsafe.Pointer(string1 + uintptr(i)))
					__t = v5
					if v6 = int32(*(*int8)(unsafe.Pointer(__ccgo_ts + 813 + uintptr(*(*int32_t)(unsafe.Pointer(v4))))))&(int32(1)<<(libc.Int32FromUint8(v5)>>int32(5))) != 0; v6 {
						__t = libc.Uint8FromInt32(int32(__t) & libc.Int32FromInt32(0x3f))
					}
					v3 = libc.BoolInt32(v6 && libc.Bool(libc.Int32FromInt32(1) != 0))
				} else {
					v7 = code_point
					*(*int32_t)(unsafe.Pointer(v7)) -= int32(0xf0)
					if v9 = *(*int32_t)(unsafe.Pointer(v7)) <= int32(4); v9 {
						v8 = *(*uint8_t)(unsafe.Pointer(string1 + uintptr(i)))
						__t = v8
					}
					if v11 = v9 && int32(*(*int8)(unsafe.Pointer(__ccgo_ts + 830 + uintptr(libc.Int32FromUint8(v8)>>int32(4)))))&(int32(1)<<*(*int32_t)(unsafe.Pointer(code_point))) != 0; v11 {
						*(*int32_t)(unsafe.Pointer(code_point)) = *(*int32_t)(unsafe.Pointer(code_point))<<libc.Int32FromInt32(6) | libc.Int32FromUint8(__t)&int32(0x3f)
						i = i + 1
						v1 = i
					}
					if v13 = v11 && v1 != length; v13 {
						v12 = libc.Uint8FromInt32(libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(string1 + uintptr(i)))) - libc.Int32FromInt32(0x80))
						__t = v12
					}
					v3 = libc.BoolInt32(v13 && libc.Int32FromUint8(v12) <= int32(0x3f))
				}
				if v15 = v3 != 0; v15 {
					*(*int32_t)(unsafe.Pointer(code_point)) = *(*int32_t)(unsafe.Pointer(code_point))<<libc.Int32FromInt32(6) | libc.Int32FromUint8(__t)
					i = i + 1
					v10 = i
				}
				v2 = libc.BoolInt32(v15 && v10 != length)
			} else {
				if v16 = *(*int32_t)(unsafe.Pointer(code_point)) >= int32(0xc2); v16 {
					*(*int32_t)(unsafe.Pointer(code_point)) &= int32(0x1f)
				}
				v2 = libc.BoolInt32(v16 && libc.Bool(libc.Int32FromInt32(1) != 0))
			}
		}
		if v19 = v17 && v2 != 0; v19 {
			v18 = libc.Uint8FromInt32(libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(string1 + uintptr(i)))) - libc.Int32FromInt32(0x80))
			__t = v18
		}
		if v20 = v19 && libc.Int32FromUint8(v18) <= int32(0x3f); v20 {
			*(*int32_t)(unsafe.Pointer(code_point)) = *(*int32_t)(unsafe.Pointer(code_point))<<libc.Int32FromInt32(6) | libc.Int32FromUint8(__t)
			i = i + 1
		}
		if v20 && libc.Bool(libc.Int32FromInt32(1) != 0) {
		} else {
			*(*int32_t)(unsafe.Pointer(code_point)) = -libc.Int32FromInt32(1)
		}
	}
	return i
}

func ts_decode_utf16_le(tls *libc.TLS, string1 uintptr, length uint32_t, code_point uintptr) (r uint32_t) {
	var __c2, v2 uint16_t
	var i, v1 uint32_t
	var v3 bool
	_, _, _, _, _ = __c2, i, v1, v2, v3
	i = uint32(0)
	v1 = i
	i = i + 1
	*(*int32_t)(unsafe.Pointer(code_point)) = libc.Int32FromUint16(__uint16_identity(tls, *(*uint16_t)(unsafe.Pointer(string1 + uintptr(v1)*2))))
	if libc.Uint32FromInt32(*(*int32_t)(unsafe.Pointer(code_point)))&uint32(0xfffffc00) == uint32(0xd800) {
		if v3 = i != length; v3 {
			v2 = *(*uint16_t)(unsafe.Pointer(string1 + uintptr(i)*2))
			__c2 = v2
		}
		if v3 && uint32(v2)&uint32(0xfffffc00) == uint32(0xdc00) {
			i = i + 1
			*(*int32_t)(unsafe.Pointer(code_point)) = *(*int32_t)(unsafe.Pointer(code_point))<<libc.Uint64FromUint64(10) + libc.Int32FromUint16(__c2) - (libc.Int32FromInt32(0xd800)<<libc.Uint64FromUint64(10) + libc.Int32FromInt32(0xdc00) - libc.Int32FromInt32(0x10000))
		}
	}
	return i * uint32(2)
}

func ts_decode_utf16_be(tls *libc.TLS, string1 uintptr, length uint32_t, code_point uintptr) (r uint32_t) {
	var __c2, v2 uint16_t
	var i, v1 uint32_t
	var v3 bool
	_, _, _, _, _ = __c2, i, v1, v2, v3
	i = uint32(0)
	v1 = i
	i = i + 1
	*(*int32_t)(unsafe.Pointer(code_point)) = libc.Int32FromUint16(__bswap_16(tls, *(*uint16_t)(unsafe.Pointer(string1 + uintptr(v1)*2))))
	if libc.Uint32FromInt32(*(*int32_t)(unsafe.Pointer(code_point)))&uint32(0xfffffc00) == uint32(0xd800) {
		if v3 = i != length; v3 {
			v2 = *(*uint16_t)(unsafe.Pointer(string1 + uintptr(i)*2))
			__c2 = v2
		}
		if v3 && uint32(v2)&uint32(0xfffffc00) == uint32(0xdc00) {
			i = i + 1
			*(*int32_t)(unsafe.Pointer(code_point)) = *(*int32_t)(unsafe.Pointer(code_point))<<libc.Uint64FromUint64(10) + libc.Int32FromUint16(__c2) - (libc.Int32FromInt32(0xd800)<<libc.Uint64FromUint64(10) + libc.Int32FromInt32(0xdc00) - libc.Int32FromInt32(0x10000))
		}
	}
	return i * uint32(2)
}

var BYTE_ORDER_MARK = int32(0xFEFF)

var DEFAULT_RANGE = TSRange{
	Fend_point: TSPoint{
		Frow:    libc.Uint32FromUint32(4294967295),
		Fcolumn: libc.Uint32FromUint32(4294967295),
	},
	Fend_byte: libc.Uint32FromUint32(4294967295),
}

func ts_lexer__set_column_data(tls *libc.TLS, self uintptr, val uint32_t) {
	(*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalid = libc.BoolUint8(1 != 0)
	(*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalue = val
}

func ts_lexer__increment_column_data(tls *libc.TLS, self uintptr) {
	if (*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalid != 0 {
		(*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalue = (*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalue + 1
	}
}

func ts_lexer__invalidate_column_data(tls *libc.TLS, self uintptr) {
	(*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalid = libc.BoolUint8(0 != 0)
	(*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalue = uint32(0)
}

func ts_lexer__eof(tls *libc.TLS, _self uintptr) (r uint8) {
	var self uintptr
	_ = self
	self = _self
	return libc.BoolUint8((*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index == (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count)
}

func ts_lexer__clear_chunk(tls *libc.TLS, self uintptr) {
	(*Lexer)(unsafe.Pointer(self)).Fchunk = libc.UintptrFromInt32(0)
	(*Lexer)(unsafe.Pointer(self)).Fchunk_size = uint32(0)
	(*Lexer)(unsafe.Pointer(self)).Fchunk_start = uint32(0)
}

func ts_lexer__get_chunk(tls *libc.TLS, self uintptr) {
	(*Lexer)(unsafe.Pointer(self)).Fchunk_start = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes
	(*Lexer)(unsafe.Pointer(self)).Fchunk = (*(*func(*libc.TLS, uintptr, uint32_t, TSPoint, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*Lexer)(unsafe.Pointer(self)).Finput.Fread})))(tls, (*Lexer)(unsafe.Pointer(self)).Finput.Fpayload, (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes, (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fextent, self+172)
	if !((*Lexer)(unsafe.Pointer(self)).Fchunk_size != 0) {
		(*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index = (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count
		(*Lexer)(unsafe.Pointer(self)).Fchunk = libc.UintptrFromInt32(0)
	}
}

func ts_lexer__get_lookahead(tls *libc.TLS, self uintptr) {
	var chunk, v1, v2, v3 uintptr
	var decode TSDecodeFunction
	var position_in_chunk, size uint32_t
	_, _, _, _, _, _, _ = chunk, decode, position_in_chunk, size, v1, v2, v3
	position_in_chunk = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes - (*Lexer)(unsafe.Pointer(self)).Fchunk_start
	size = (*Lexer)(unsafe.Pointer(self)).Fchunk_size - position_in_chunk
	if size == uint32(0) {
		(*Lexer)(unsafe.Pointer(self)).Flookahead_size = uint32(1)
		(*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead = int32('\000')
		return
	}
	chunk = (*Lexer)(unsafe.Pointer(self)).Fchunk + uintptr(position_in_chunk)
	if (*Lexer)(unsafe.Pointer(self)).Finput.Fencoding == int32(TSInputEncodingUTF8) {
		v1 = __ccgo_fp(ts_decode_utf8)
	} else {
		if (*Lexer)(unsafe.Pointer(self)).Finput.Fencoding == int32(TSInputEncodingUTF16LE) {
			v2 = __ccgo_fp(ts_decode_utf16_le)
		} else {
			if (*Lexer)(unsafe.Pointer(self)).Finput.Fencoding == int32(TSInputEncodingUTF16BE) {
				v3 = __ccgo_fp(ts_decode_utf16_be)
			} else {
				v3 = (*Lexer)(unsafe.Pointer(self)).Finput.Fdecode
			}
			v2 = v3
		}
		v1 = v2
	}
	decode = v1
	(*Lexer)(unsafe.Pointer(self)).Flookahead_size = (*(*func(*libc.TLS, uintptr, uint32_t, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{decode})))(tls, chunk, size, self)
	if (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead == TS_DECODE_ERROR && size < uint32(4) {
		ts_lexer__get_chunk(tls, self)
		chunk = (*Lexer)(unsafe.Pointer(self)).Fchunk
		size = (*Lexer)(unsafe.Pointer(self)).Fchunk_size
		(*Lexer)(unsafe.Pointer(self)).Flookahead_size = (*(*func(*libc.TLS, uintptr, uint32_t, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{decode})))(tls, chunk, size, self)
	}
	if (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead == TS_DECODE_ERROR {
		(*Lexer)(unsafe.Pointer(self)).Flookahead_size = uint32(1)
	}
}

func ts_lexer_goto(tls *libc.TLS, self uintptr, position Length) {
	var found_included_range uint8
	var i uint32
	var included_range, last_included_range uintptr
	_, _, _, _ = found_included_range, i, included_range, last_included_range
	if position.Fbytes != (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes {
		ts_lexer__invalidate_column_data(tls, self)
	}
	(*Lexer)(unsafe.Pointer(self)).Fcurrent_position = position
	found_included_range = libc.BoolUint8(0 != 0)
	i = uint32(0)
	for {
		if !(i < (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count) {
			break
		}
		included_range = (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges + uintptr(i)*24
		if (*TSRange)(unsafe.Pointer(included_range)).Fend_byte > (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes && (*TSRange)(unsafe.Pointer(included_range)).Fend_byte > (*TSRange)(unsafe.Pointer(included_range)).Fstart_byte {
			if (*TSRange)(unsafe.Pointer(included_range)).Fstart_byte >= (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes {
				(*Lexer)(unsafe.Pointer(self)).Fcurrent_position = Length{
					Fbytes:  (*TSRange)(unsafe.Pointer(included_range)).Fstart_byte,
					Fextent: (*TSRange)(unsafe.Pointer(included_range)).Fstart_point,
				}
			}
			(*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index = i
			found_included_range = libc.BoolUint8(1 != 0)
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if found_included_range != 0 {
		if (*Lexer)(unsafe.Pointer(self)).Fchunk != 0 && ((*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes < (*Lexer)(unsafe.Pointer(self)).Fchunk_start || (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes >= (*Lexer)(unsafe.Pointer(self)).Fchunk_start+(*Lexer)(unsafe.Pointer(self)).Fchunk_size) {
			ts_lexer__clear_chunk(tls, self)
		}
		(*Lexer)(unsafe.Pointer(self)).Flookahead_size = uint32(0)
		(*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead = int32('\000')
	} else {
		(*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index = (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count
		last_included_range = (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges + uintptr((*Lexer)(unsafe.Pointer(self)).Fincluded_range_count-uint32(1))*24
		(*Lexer)(unsafe.Pointer(self)).Fcurrent_position = Length{
			Fbytes:  (*TSRange)(unsafe.Pointer(last_included_range)).Fend_byte,
			Fextent: (*TSRange)(unsafe.Pointer(last_included_range)).Fend_point,
		}
		ts_lexer__clear_chunk(tls, self)
		(*Lexer)(unsafe.Pointer(self)).Flookahead_size = uint32(1)
		(*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead = int32('\000')
	}
}

func ts_lexer__do_advance(tls *libc.TLS, self uintptr, skip uint8) {
	var current_range uintptr
	var is_bom uint8
	_, _ = current_range, is_bom
	if (*Lexer)(unsafe.Pointer(self)).Flookahead_size != 0 {
		if (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead == int32('\n') {
			(*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fextent.Frow = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fextent.Frow + 1
			(*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fextent.Fcolumn = uint32(0)
			ts_lexer__set_column_data(tls, self, uint32(0))
		} else {
			is_bom = libc.BoolUint8((*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes == uint32(0) && (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead == BYTE_ORDER_MARK)
			if !(is_bom != 0) {
				ts_lexer__increment_column_data(tls, self)
			}
			(*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fextent.Fcolumn += (*Lexer)(unsafe.Pointer(self)).Flookahead_size
		}
		(*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes += (*Lexer)(unsafe.Pointer(self)).Flookahead_size
	}
	current_range = (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges + uintptr((*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index)*24
	for (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes >= (*TSRange)(unsafe.Pointer(current_range)).Fend_byte || (*TSRange)(unsafe.Pointer(current_range)).Fend_byte == (*TSRange)(unsafe.Pointer(current_range)).Fstart_byte {
		if (*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index < (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count {
			(*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index = (*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index + 1
		}
		if (*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index < (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count {
			current_range += 24
			(*Lexer)(unsafe.Pointer(self)).Fcurrent_position = Length{
				Fbytes:  (*TSRange)(unsafe.Pointer(current_range)).Fstart_byte,
				Fextent: (*TSRange)(unsafe.Pointer(current_range)).Fstart_point,
			}
		} else {
			current_range = libc.UintptrFromInt32(0)
			break
		}
	}
	if skip != 0 {
		(*Lexer)(unsafe.Pointer(self)).Ftoken_start_position = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position
	}
	if current_range != 0 {
		if (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes < (*Lexer)(unsafe.Pointer(self)).Fchunk_start || (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes >= (*Lexer)(unsafe.Pointer(self)).Fchunk_start+(*Lexer)(unsafe.Pointer(self)).Fchunk_size {
			ts_lexer__get_chunk(tls, self)
		}
		ts_lexer__get_lookahead(tls, self)
	} else {
		ts_lexer__clear_chunk(tls, self)
		(*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead = int32('\000')
		(*Lexer)(unsafe.Pointer(self)).Flookahead_size = uint32(1)
	}
}

func ts_lexer__advance(tls *libc.TLS, _self uintptr, skip uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var self, v1 uintptr
	_, _ = self, v1
	self = _self
	if !((*Lexer)(unsafe.Pointer(self)).Fchunk != 0) {
		return
	}
	if skip != 0 {
		if (*Lexer)(unsafe.Pointer(self)).Flogger.Flog != 0 {
			if int32(32) <= (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead && (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead < int32(127) {
				v1 = __ccgo_ts + 847
			} else {
				v1 = __ccgo_ts + 867
			}
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), v1, libc.VaList(bp+8, (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead))
			(*(*func(*libc.TLS, uintptr, TSLogType1, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*Lexer)(unsafe.Pointer(self)).Flogger.Flog})))(tls, (*Lexer)(unsafe.Pointer(self)).Flogger.Fpayload, int32(TSLogTypeLex), self+192)
		}
	} else {
		if (*Lexer)(unsafe.Pointer(self)).Flogger.Flog != 0 {
			if int32(32) <= (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead && (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead < int32(127) {
				v1 = __ccgo_ts + 885
			} else {
				v1 = __ccgo_ts + 908
			}
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), v1, libc.VaList(bp+8, (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead))
			(*(*func(*libc.TLS, uintptr, TSLogType1, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*Lexer)(unsafe.Pointer(self)).Flogger.Flog})))(tls, (*Lexer)(unsafe.Pointer(self)).Flogger.Fpayload, int32(TSLogTypeLex), self+192)
		}
	}
	ts_lexer__do_advance(tls, self, skip)
}

func ts_lexer__mark_end(tls *libc.TLS, _self uintptr) {
	var current_included_range, previous_included_range, self uintptr
	_, _, _ = current_included_range, previous_included_range, self
	self = _self
	if !(ts_lexer__eof(tls, self) != 0) {
		current_included_range = (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges + uintptr((*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index)*24
		if (*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index > uint32(0) && (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes == (*TSRange)(unsafe.Pointer(current_included_range)).Fstart_byte {
			previous_included_range = current_included_range - uintptr(1)*24
			(*Lexer)(unsafe.Pointer(self)).Ftoken_end_position = Length{
				Fbytes:  (*TSRange)(unsafe.Pointer(previous_included_range)).Fend_byte,
				Fextent: (*TSRange)(unsafe.Pointer(previous_included_range)).Fend_point,
			}
			return
		}
	}
	(*Lexer)(unsafe.Pointer(self)).Ftoken_end_position = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position
}

func ts_lexer__get_column(tls *libc.TLS, _self uintptr) (r uint32_t) {
	var goal_byte uint32_t
	var self uintptr
	var start_of_col Length
	_, _, _ = goal_byte, self, start_of_col
	self = _self
	(*Lexer)(unsafe.Pointer(self)).Fdid_get_column = libc.BoolUint8(1 != 0)
	if !((*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalid != 0) {
		goal_byte = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes
		start_of_col = Length{
			Fbytes: (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes - (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fextent.Fcolumn,
			Fextent: TSPoint{
				Frow: (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fextent.Frow,
			},
		}
		ts_lexer_goto(tls, self, start_of_col)
		ts_lexer__set_column_data(tls, self, uint32(0))
		ts_lexer__get_chunk(tls, self)
		if !(ts_lexer__eof(tls, _self) != 0) {
			ts_lexer__get_lookahead(tls, self)
			for (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes < goal_byte && !(ts_lexer__eof(tls, _self) != 0) && (*Lexer)(unsafe.Pointer(self)).Fchunk != 0 {
				ts_lexer__do_advance(tls, self, libc.BoolUint8(0 != 0))
				if ts_lexer__eof(tls, _self) != 0 {
					break
				}
			}
		}
	}
	return (*Lexer)(unsafe.Pointer(self)).Fcolumn_data.Fvalue
}

func ts_lexer__is_at_included_range_start(tls *libc.TLS, _self uintptr) (r uint8) {
	var current_range, self uintptr
	_, _ = current_range, self
	self = _self
	if (*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index < (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count {
		current_range = (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges + uintptr((*Lexer)(unsafe.Pointer(self)).Fcurrent_included_range_index)*24
		return libc.BoolUint8((*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes == (*TSRange)(unsafe.Pointer(current_range)).Fstart_byte)
	} else {
		return libc.BoolUint8(0 != 0)
	}
	return r
}

func ts_lexer__log(tls *libc.TLS, _self uintptr, fmt uintptr, va uintptr) {
	var args va_list
	var self uintptr
	_, _ = args, self
	self = _self
	args = va
	if (*Lexer)(unsafe.Pointer(self)).Flogger.Flog != 0 {
		libc.X__builtin_vsnprintf(tls, self+192, uint64(1024), fmt, args)
		(*(*func(*libc.TLS, uintptr, TSLogType1, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*Lexer)(unsafe.Pointer(self)).Flogger.Flog})))(tls, (*Lexer)(unsafe.Pointer(self)).Flogger.Fpayload, int32(TSLogTypeLex), self+192)
	}
	_ = args
}

func ts_lexer_init(tls *libc.TLS, self uintptr) {
	*(*Lexer)(unsafe.Pointer(self)) = Lexer{
		Fdata: TSLexer{
			Fadvance:                    __ccgo_fp(ts_lexer__advance),
			Fmark_end:                   __ccgo_fp(ts_lexer__mark_end),
			Fget_column:                 __ccgo_fp(ts_lexer__get_column),
			Fis_at_included_range_start: __ccgo_fp(ts_lexer__is_at_included_range_start),
			Feof:                        __ccgo_fp(ts_lexer__eof),
			Flog:                        __ccgo_fp(ts_lexer__log),
		},
	}
	ts_lexer_set_included_ranges(tls, self, libc.UintptrFromInt32(0), uint32(0))
}

func ts_lexer_delete(tls *libc.TLS, self uintptr) {
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges)
}

func ts_lexer_set_input(tls *libc.TLS, self uintptr, input TSInput) {
	(*Lexer)(unsafe.Pointer(self)).Finput = input
	ts_lexer__clear_chunk(tls, self)
	ts_lexer_goto(tls, self, (*Lexer)(unsafe.Pointer(self)).Fcurrent_position)
}

func ts_lexer_reset(tls *libc.TLS, self uintptr, position Length) {
	if position.Fbytes != (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes {
		ts_lexer_goto(tls, self, position)
	}
}

func ts_lexer_start(tls *libc.TLS, self uintptr) {
	(*Lexer)(unsafe.Pointer(self)).Ftoken_start_position = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position
	(*Lexer)(unsafe.Pointer(self)).Ftoken_end_position = LENGTH_UNDEFINED
	(*Lexer)(unsafe.Pointer(self)).Fdata.Fresult_symbol = uint16(0)
	(*Lexer)(unsafe.Pointer(self)).Fdid_get_column = libc.BoolUint8(0 != 0)
	if !(ts_lexer__eof(tls, self) != 0) {
		if !((*Lexer)(unsafe.Pointer(self)).Fchunk_size != 0) {
			ts_lexer__get_chunk(tls, self)
		}
		if !((*Lexer)(unsafe.Pointer(self)).Flookahead_size != 0) {
			ts_lexer__get_lookahead(tls, self)
		}
		if (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes == uint32(0) {
			if (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead == BYTE_ORDER_MARK {
				ts_lexer__advance(tls, self, libc.BoolUint8(1 != 0))
			}
			ts_lexer__set_column_data(tls, self, uint32(0))
		}
	}
}

func ts_lexer_finish(tls *libc.TLS, self uintptr, lookahead_end_byte uintptr) {
	var current_lookahead_end_byte uint32_t
	_ = current_lookahead_end_byte
	if length_is_undefined(tls, (*Lexer)(unsafe.Pointer(self)).Ftoken_end_position) != 0 {
		ts_lexer__mark_end(tls, self)
	}
	if (*Lexer)(unsafe.Pointer(self)).Ftoken_end_position.Fbytes < (*Lexer)(unsafe.Pointer(self)).Ftoken_start_position.Fbytes {
		(*Lexer)(unsafe.Pointer(self)).Ftoken_start_position = (*Lexer)(unsafe.Pointer(self)).Ftoken_end_position
	}
	current_lookahead_end_byte = (*Lexer)(unsafe.Pointer(self)).Fcurrent_position.Fbytes + uint32(1)
	if (*Lexer)(unsafe.Pointer(self)).Fdata.Flookahead == TS_DECODE_ERROR {
		current_lookahead_end_byte = current_lookahead_end_byte + uint32(4)
	}
	if current_lookahead_end_byte > *(*uint32_t)(unsafe.Pointer(lookahead_end_byte)) {
		*(*uint32_t)(unsafe.Pointer(lookahead_end_byte)) = current_lookahead_end_byte
	}
}

func ts_lexer_mark_end(tls *libc.TLS, self uintptr) {
	ts_lexer__mark_end(tls, self)
}

func ts_lexer_set_included_ranges(tls *libc.TLS, self uintptr, ranges uintptr, count uint32_t) (r uint8) {
	var i uint32
	var previous_byte uint32_t
	var range1 uintptr
	var size size_t
	_, _, _, _ = i, previous_byte, range1, size
	if count == uint32(0) || !(ranges != 0) {
		ranges = uintptr(unsafe.Pointer(&DEFAULT_RANGE))
		count = uint32(1)
	} else {
		previous_byte = uint32(0)
		i = uint32(0)
		for {
			if !(i < count) {
				break
			}
			range1 = ranges + uintptr(i)*24
			if (*TSRange)(unsafe.Pointer(range1)).Fstart_byte < previous_byte || (*TSRange)(unsafe.Pointer(range1)).Fend_byte < (*TSRange)(unsafe.Pointer(range1)).Fstart_byte {
				return libc.BoolUint8(0 != 0)
			}
			previous_byte = (*TSRange)(unsafe.Pointer(range1)).Fend_byte
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	size = uint64(count) * uint64(24)
	(*Lexer)(unsafe.Pointer(self)).Fincluded_ranges = (*(*func(*libc.TLS, uintptr, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_realloc})))(tls, (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges, size)
	libc.Xmemcpy(tls, (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges, ranges, size)
	(*Lexer)(unsafe.Pointer(self)).Fincluded_range_count = count
	ts_lexer_goto(tls, self, (*Lexer)(unsafe.Pointer(self)).Fcurrent_position)
	return libc.BoolUint8(1 != 0)
}

func ts_lexer_included_ranges(tls *libc.TLS, self uintptr, count uintptr) (r uintptr) {
	*(*uint32_t)(unsafe.Pointer(count)) = (*Lexer)(unsafe.Pointer(self)).Fincluded_range_count
	return (*Lexer)(unsafe.Pointer(self)).Fincluded_ranges
}

type ParentCacheEntry = struct {
	Fchild        uintptr
	Fparent       uintptr
	Fposition     Length
	Falias_symbol TSSymbol
}

type NodeChildIterator = struct {
	Fparent                 Subtree
	Ftree                   uintptr
	Fposition               Length
	Fchild_index            uint32_t
	Fstructural_child_index uint32_t
	Falias_sequence         uintptr
}

func ts_node_new(tls *libc.TLS, tree uintptr, subtree uintptr, position Length, alias TSSymbol) (r TSNode) {
	return TSNode{
		Fcontext: [4]uint32_t{
			0: position.Fbytes,
			1: position.Fextent.Frow,
			2: position.Fextent.Fcolumn,
			3: uint32(alias),
		},
		Fid:   subtree,
		Ftree: tree,
	}
}

func ts_node__null(tls *libc.TLS) (r TSNode) {
	return ts_node_new(tls, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), length_zero(tls), uint16(0))
}

func ts_node_start_byte(tls *libc.TLS, _self TSNode) (r uint32_t) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	return *(*uint32_t)(unsafe.Pointer(bp))
}

func ts_node_start_point(tls *libc.TLS, _self TSNode) (r TSPoint) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	return TSPoint{
		Frow:    *(*uint32_t)(unsafe.Pointer(bp + 1*4)),
		Fcolumn: *(*uint32_t)(unsafe.Pointer(bp + 2*4)),
	}
}

func ts_node__alias(tls *libc.TLS, self uintptr) (r uint32_t) {
	return *(*uint32_t)(unsafe.Pointer(self + 3*4))
}

func ts_node__subtree(tls *libc.TLS, self TSNode) (r Subtree) {
	return *(*Subtree)(unsafe.Pointer(self.Fid))
}

func ts_node_iterate_children(tls *libc.TLS, node uintptr) (r NodeChildIterator) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var alias_sequence uintptr
	var _ /* subtree at bp+0 */ Subtree
	_ = alias_sequence
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(node)))
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) == uint32(0) {
		return NodeChildIterator{
			Fparent:   Subtree{},
			Ftree:     (*TSNode)(unsafe.Pointer(node)).Ftree,
			Fposition: length_zero(tls),
		}
	}
	alias_sequence = ts_language_alias_sequence(tls, (*TSTree)(unsafe.Pointer((*TSNode)(unsafe.Pointer(node)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id))
	return NodeChildIterator{
		Fparent: *(*Subtree)(unsafe.Pointer(bp)),
		Ftree:   (*TSNode)(unsafe.Pointer(node)).Ftree,
		Fposition: Length{
			Fbytes:  ts_node_start_byte(tls, *(*TSNode)(unsafe.Pointer(node))),
			Fextent: ts_node_start_point(tls, *(*TSNode)(unsafe.Pointer(node))),
		},
		Falias_sequence: alias_sequence,
	}
}

func ts_node_child_iterator_done(tls *libc.TLS, self uintptr) (r uint8) {
	return libc.BoolUint8((*NodeChildIterator)(unsafe.Pointer(self)).Fchild_index == (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count)
}

func ts_node_child_iterator_next(tls *libc.TLS, self uintptr, result uintptr) (r uint8) {
	var alias_symbol TSSymbol
	var child, v1 uintptr
	_, _, _ = alias_symbol, child, v1
	if !(*(*uintptr)(unsafe.Pointer(self)) != 0) || ts_node_child_iterator_done(tls, self) != 0 {
		return libc.BoolUint8(0 != 0)
	}
	if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = *(*uintptr)(unsafe.Pointer(self)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count)*8
	}
	child = v1 + uintptr((*NodeChildIterator)(unsafe.Pointer(self)).Fchild_index)*8
	alias_symbol = uint16(0)
	if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(child))) != 0) {
		if (*NodeChildIterator)(unsafe.Pointer(self)).Falias_sequence != 0 {
			alias_symbol = *(*TSSymbol)(unsafe.Pointer((*NodeChildIterator)(unsafe.Pointer(self)).Falias_sequence + uintptr((*NodeChildIterator)(unsafe.Pointer(self)).Fstructural_child_index)*2))
		}
		(*NodeChildIterator)(unsafe.Pointer(self)).Fstructural_child_index = (*NodeChildIterator)(unsafe.Pointer(self)).Fstructural_child_index + 1
	}
	if (*NodeChildIterator)(unsafe.Pointer(self)).Fchild_index > uint32(0) {
		(*NodeChildIterator)(unsafe.Pointer(self)).Fposition = length_add(tls, (*NodeChildIterator)(unsafe.Pointer(self)).Fposition, ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(child))))
	}
	*(*TSNode)(unsafe.Pointer(result)) = ts_node_new(tls, (*NodeChildIterator)(unsafe.Pointer(self)).Ftree, child, (*NodeChildIterator)(unsafe.Pointer(self)).Fposition, alias_symbol)
	(*NodeChildIterator)(unsafe.Pointer(self)).Fposition = length_add(tls, (*NodeChildIterator)(unsafe.Pointer(self)).Fposition, ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(child))))
	(*NodeChildIterator)(unsafe.Pointer(self)).Fchild_index = (*NodeChildIterator)(unsafe.Pointer(self)).Fchild_index + 1
	return libc.BoolUint8(1 != 0)
}

func ts_node__is_relevant(tls *libc.TLS, _self TSNode, include_anonymous uint8) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	var alias TSSymbol
	var tree Subtree
	_, _ = alias, tree
	tree = ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp)))
	if include_anonymous != 0 {
		return libc.BoolUint8(ts_subtree_visible(tls, tree) != 0 || ts_node__alias(tls, bp) != 0)
	} else {
		alias = uint16(ts_node__alias(tls, bp))
		if alias != 0 {
			return ts_language_symbol_metadata(tls, (*TSTree)(unsafe.Pointer((*(*TSNode)(unsafe.Pointer(bp))).Ftree)).Flanguage, alias).Fnamed
		} else {
			return libc.BoolUint8(ts_subtree_visible(tls, tree) != 0 && ts_subtree_named(tls, tree) != 0)
		}
	}
	return r
}

func ts_node__relevant_child_count(tls *libc.TLS, self TSNode, include_anonymous uint8) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* tree at bp+0 */ Subtree
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = ts_node__subtree(tls, self)
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
		if include_anonymous != 0 {
			return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count
		} else {
			return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count
		}
	} else {
		return uint32(0)
	}
	return r
}

func ts_node__child(tls *libc.TLS, self TSNode, child_index uint32_t, include_anonymous uint8) (r TSNode) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var did_descend uint8
	var grandchild_count, grandchild_index, index uint32_t
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* result at bp+0 */ TSNode
	_, _, _, _ = did_descend, grandchild_count, grandchild_index, index
	*(*TSNode)(unsafe.Pointer(bp)) = self
	did_descend = libc.BoolUint8(1 != 0)
	for did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		index = uint32(0)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), include_anonymous) != 0 {
				if index == child_index {
					return *(*TSNode)(unsafe.Pointer(bp + 32))
				}
				index = index + 1
			} else {
				grandchild_index = child_index - index
				grandchild_count = ts_node__relevant_child_count(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), include_anonymous)
				if grandchild_index < grandchild_count {
					did_descend = libc.BoolUint8(1 != 0)
					*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 32))
					child_index = grandchild_index
					break
				}
				index = index + grandchild_count
			}
		}
	}
	return ts_node__null(tls)
}

func ts_subtree_has_trailing_empty_descendant(tls *libc.TLS, _self Subtree, _other Subtree) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	*(*Subtree)(unsafe.Pointer(bp + 8)) = _other
	var i uint32
	var v2 uintptr
	var _ /* child at bp+16 */ Subtree
	_, _ = i, v2
	i = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) - uint32(1)
	for {
		if !(i+uint32(1) > uint32(0)) {
			break
		}
		*(*Subtree)(unsafe.Pointer(bp + 16)) = Subtree{}
		if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
			v2 = libc.UintptrFromInt32(0)
		} else {
			v2 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
		}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 16)) = *(*Subtree)(unsafe.Pointer(v2 + uintptr(i)*8))
		if ts_subtree_total_bytes(tls, *(*Subtree)(unsafe.Pointer(bp + 16))) > uint32(0) {
			break
		}
		if *(*uintptr)(unsafe.Pointer(bp + 16)) == *(*uintptr)(unsafe.Pointer(bp + 8)) || ts_subtree_has_trailing_empty_descendant(tls, *(*Subtree)(unsafe.Pointer(bp + 16)), *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
			return libc.BoolUint8(1 != 0)
		}
		goto _1
	_1:
		;
		i = i - 1
	}
	return libc.BoolUint8(0 != 0)
}

func ts_node__prev_sibling(tls *libc.TLS, self TSNode, include_anonymous uint8) (r TSNode) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var earlier_child, earlier_node TSNode
	var earlier_child_is_relevant, earlier_node_is_relevant, found_child_containing_target, self_is_empty uint8
	var self_subtree Subtree
	var target_end_byte uint32_t
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* node at bp+0 */ TSNode
	_, _, _, _, _, _, _, _ = earlier_child, earlier_child_is_relevant, earlier_node, earlier_node_is_relevant, found_child_containing_target, self_is_empty, self_subtree, target_end_byte
	self_subtree = ts_node__subtree(tls, self)
	self_is_empty = libc.BoolUint8(ts_subtree_total_bytes(tls, self_subtree) == uint32(0))
	target_end_byte = ts_node_end_byte(tls, self)
	*(*TSNode)(unsafe.Pointer(bp)) = ts_node_parent(tls, self)
	earlier_node = ts_node__null(tls)
	earlier_node_is_relevant = libc.BoolUint8(0 != 0)
	for !(ts_node_is_null(tls, *(*TSNode)(unsafe.Pointer(bp))) != 0) {
		earlier_child = ts_node__null(tls)
		earlier_child_is_relevant = libc.BoolUint8(0 != 0)
		found_child_containing_target = libc.BoolUint8(0 != 0)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			if (*(*TSNode)(unsafe.Pointer(bp + 32))).Fid == self.Fid {
				break
			}
			if (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fposition.Fbytes > target_end_byte {
				found_child_containing_target = libc.BoolUint8(1 != 0)
				break
			}
			if (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fposition.Fbytes == target_end_byte && (!(self_is_empty != 0) || ts_subtree_has_trailing_empty_descendant(tls, ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp + 32))), self_subtree) != 0) {
				found_child_containing_target = libc.BoolUint8(1 != 0)
				break
			}
			if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), include_anonymous) != 0 {
				earlier_child = *(*TSNode)(unsafe.Pointer(bp + 32))
				earlier_child_is_relevant = libc.BoolUint8(1 != 0)
			} else {
				if ts_node__relevant_child_count(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), include_anonymous) > uint32(0) {
					earlier_child = *(*TSNode)(unsafe.Pointer(bp + 32))
					earlier_child_is_relevant = libc.BoolUint8(0 != 0)
				}
			}
		}
		if found_child_containing_target != 0 {
			if !(ts_node_is_null(tls, earlier_child) != 0) {
				earlier_node = earlier_child
				earlier_node_is_relevant = earlier_child_is_relevant
			}
			*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 32))
		} else {
			if earlier_child_is_relevant != 0 {
				return earlier_child
			} else {
				if !(ts_node_is_null(tls, earlier_child) != 0) {
					*(*TSNode)(unsafe.Pointer(bp)) = earlier_child
				} else {
					if earlier_node_is_relevant != 0 {
						return earlier_node
					} else {
						*(*TSNode)(unsafe.Pointer(bp)) = earlier_node
						earlier_node = ts_node__null(tls)
						earlier_node_is_relevant = libc.BoolUint8(0 != 0)
					}
				}
			}
		}
	}
	return ts_node__null(tls)
}

func ts_node__next_sibling(tls *libc.TLS, self TSNode, include_anonymous uint8) (r TSNode) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var child_containing_target, later_child, later_node TSNode
	var child_start_byte, start_byte, target_end_byte uint32_t
	var contains_target, is_empty, later_child_is_relevant, later_node_is_relevant uint8
	var v1 int32
	var v2, v3 Subtree
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* node at bp+0 */ TSNode
	_, _, _, _, _, _, _, _, _, _, _, _, _ = child_containing_target, child_start_byte, contains_target, is_empty, later_child, later_child_is_relevant, later_node, later_node_is_relevant, start_byte, target_end_byte, v1, v2, v3
	target_end_byte = ts_node_end_byte(tls, self)
	*(*TSNode)(unsafe.Pointer(bp)) = ts_node_parent(tls, self)
	later_node = ts_node__null(tls)
	later_node_is_relevant = libc.BoolUint8(0 != 0)
	for !(ts_node_is_null(tls, *(*TSNode)(unsafe.Pointer(bp))) != 0) {
		later_child = ts_node__null(tls)
		later_child_is_relevant = libc.BoolUint8(0 != 0)
		child_containing_target = ts_node__null(tls)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			if (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fposition.Fbytes <= target_end_byte {
				continue
			}
			start_byte = ts_node_start_byte(tls, self)
			child_start_byte = ts_node_start_byte(tls, *(*TSNode)(unsafe.Pointer(bp + 32)))
			is_empty = libc.BoolUint8(start_byte == target_end_byte)
			if is_empty != 0 {
				v1 = libc.BoolInt32(child_start_byte < start_byte)
			} else {
				v1 = libc.BoolInt32(child_start_byte <= start_byte)
			}
			contains_target = libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
			if contains_target != 0 {
				v2 = ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp + 32)))
				v3 = ts_node__subtree(tls, self)
				if *(*uintptr)(unsafe.Pointer(&v2)) != *(*uintptr)(unsafe.Pointer(&v3)) {
					child_containing_target = *(*TSNode)(unsafe.Pointer(bp + 32))
				}
			} else {
				if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), include_anonymous) != 0 {
					later_child = *(*TSNode)(unsafe.Pointer(bp + 32))
					later_child_is_relevant = libc.BoolUint8(1 != 0)
					break
				} else {
					if ts_node__relevant_child_count(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), include_anonymous) > uint32(0) {
						later_child = *(*TSNode)(unsafe.Pointer(bp + 32))
						later_child_is_relevant = libc.BoolUint8(0 != 0)
						break
					}
				}
			}
		}
		if !(ts_node_is_null(tls, child_containing_target) != 0) {
			if !(ts_node_is_null(tls, later_child) != 0) {
				later_node = later_child
				later_node_is_relevant = later_child_is_relevant
			}
			*(*TSNode)(unsafe.Pointer(bp)) = child_containing_target
		} else {
			if later_child_is_relevant != 0 {
				return later_child
			} else {
				if !(ts_node_is_null(tls, later_child) != 0) {
					*(*TSNode)(unsafe.Pointer(bp)) = later_child
				} else {
					if later_node_is_relevant != 0 {
						return later_node
					} else {
						*(*TSNode)(unsafe.Pointer(bp)) = later_node
					}
				}
			}
		}
	}
	return ts_node__null(tls)
}

func ts_node__first_child_for_byte(tls *libc.TLS, self TSNode, goal uint32_t, include_anonymous uint8) (r TSNode) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var did_descend, has_last_iterator uint8
	var last_iterator NodeChildIterator
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* node at bp+0 */ TSNode
	_, _, _ = did_descend, has_last_iterator, last_iterator
	*(*TSNode)(unsafe.Pointer(bp)) = self
	did_descend = libc.BoolUint8(1 != 0)
	has_last_iterator = libc.BoolUint8(0 != 0)
	for did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		goto loop
	loop:
		;
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			if ts_node_end_byte(tls, *(*TSNode)(unsafe.Pointer(bp + 32))) > goal {
				if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), include_anonymous) != 0 {
					return *(*TSNode)(unsafe.Pointer(bp + 32))
				} else {
					if ts_node_child_count(tls, *(*TSNode)(unsafe.Pointer(bp + 32))) > uint32(0) {
						if (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fchild_index < ts_subtree_child_count(tls, ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp + 32)))) {
							last_iterator = *(*NodeChildIterator)(unsafe.Pointer(bp + 64))
							has_last_iterator = libc.BoolUint8(1 != 0)
						}
						did_descend = libc.BoolUint8(1 != 0)
						*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 32))
						break
					}
				}
			}
		}
		if !(did_descend != 0) && has_last_iterator != 0 {
			*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = last_iterator
			has_last_iterator = libc.BoolUint8(0 != 0)
			goto loop
		}
	}
	return ts_node__null(tls)
}

func ts_node__descendant_for_byte_range(tls *libc.TLS, self TSNode, range_start uint32_t, range_end uint32_t, include_anonymous uint8) (r TSNode) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var did_descend, is_empty uint8
	var last_visible_node TSNode
	var node_end uint32_t
	var v1 int32
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* node at bp+0 */ TSNode
	_, _, _, _, _ = did_descend, is_empty, last_visible_node, node_end, v1
	if range_start > range_end {
		return ts_node__null(tls)
	}
	*(*TSNode)(unsafe.Pointer(bp)) = self
	last_visible_node = self
	did_descend = libc.BoolUint8(1 != 0)
	for did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			node_end = (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fposition.Fbytes
			if node_end < range_end {
				continue
			}
			is_empty = libc.BoolUint8(ts_node_start_byte(tls, *(*TSNode)(unsafe.Pointer(bp + 32))) == node_end)
			if is_empty != 0 {
				v1 = libc.BoolInt32(node_end < range_start)
			} else {
				v1 = libc.BoolInt32(node_end <= range_start)
			}
			if v1 != 0 {
				continue
			}
			if range_start < ts_node_start_byte(tls, *(*TSNode)(unsafe.Pointer(bp + 32))) {
				break
			}
			*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 32))
			if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp)), include_anonymous) != 0 {
				last_visible_node = *(*TSNode)(unsafe.Pointer(bp))
			}
			did_descend = libc.BoolUint8(1 != 0)
			break
		}
	}
	return last_visible_node
}

func ts_node__descendant_for_point_range(tls *libc.TLS, self TSNode, range_start TSPoint, range_end TSPoint, include_anonymous uint8) (r TSNode) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var did_descend, is_empty uint8
	var last_visible_node TSNode
	var node_end TSPoint
	var v1 int32
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* node at bp+0 */ TSNode
	_, _, _, _, _ = did_descend, is_empty, last_visible_node, node_end, v1
	if point_gt(tls, range_start, range_end) != 0 {
		return ts_node__null(tls)
	}
	*(*TSNode)(unsafe.Pointer(bp)) = self
	last_visible_node = self
	did_descend = libc.BoolUint8(1 != 0)
	for did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			node_end = (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fposition.Fextent
			if point_lt(tls, node_end, range_end) != 0 {
				continue
			}
			is_empty = point_eq(tls, ts_node_start_point(tls, *(*TSNode)(unsafe.Pointer(bp + 32))), node_end)
			if is_empty != 0 {
				v1 = libc.Int32FromUint8(point_lt(tls, node_end, range_start))
			} else {
				v1 = libc.Int32FromUint8(point_lte(tls, node_end, range_start))
			}
			if v1 != 0 {
				continue
			}
			if point_lt(tls, range_start, ts_node_start_point(tls, *(*TSNode)(unsafe.Pointer(bp + 32)))) != 0 {
				break
			}
			*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 32))
			if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp)), include_anonymous) != 0 {
				last_visible_node = *(*TSNode)(unsafe.Pointer(bp))
			}
			did_descend = libc.BoolUint8(1 != 0)
			break
		}
	}
	return last_visible_node
}

func ts_node_end_byte(tls *libc.TLS, self TSNode) (r uint32_t) {
	return ts_node_start_byte(tls, self) + ts_subtree_size(tls, ts_node__subtree(tls, self)).Fbytes
}

func ts_node_end_point(tls *libc.TLS, self TSNode) (r TSPoint) {
	return point_add(tls, ts_node_start_point(tls, self), ts_subtree_size(tls, ts_node__subtree(tls, self)).Fextent)
}

func ts_node_symbol(tls *libc.TLS, _self TSNode) (r TSSymbol) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	var symbol TSSymbol
	_ = symbol
	symbol = uint16(ts_node__alias(tls, bp))
	if !(symbol != 0) {
		symbol = ts_subtree_symbol(tls, ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp))))
	}
	return ts_language_public_symbol(tls, (*TSTree)(unsafe.Pointer((*(*TSNode)(unsafe.Pointer(bp))).Ftree)).Flanguage, symbol)
}

func ts_node_type(tls *libc.TLS, _self TSNode) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	var symbol TSSymbol
	_ = symbol
	symbol = uint16(ts_node__alias(tls, bp))
	if !(symbol != 0) {
		symbol = ts_subtree_symbol(tls, ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp))))
	}
	return ts_language_symbol_name(tls, (*TSTree)(unsafe.Pointer((*(*TSNode)(unsafe.Pointer(bp))).Ftree)).Flanguage, symbol)
}

func ts_node_language(tls *libc.TLS, self TSNode) (r uintptr) {
	return (*TSTree)(unsafe.Pointer(self.Ftree)).Flanguage
}

func ts_node_grammar_symbol(tls *libc.TLS, self TSNode) (r TSSymbol) {
	return ts_subtree_symbol(tls, ts_node__subtree(tls, self))
}

func ts_node_grammar_type(tls *libc.TLS, self TSNode) (r uintptr) {
	var symbol TSSymbol
	_ = symbol
	symbol = ts_subtree_symbol(tls, ts_node__subtree(tls, self))
	return ts_language_symbol_name(tls, (*TSTree)(unsafe.Pointer(self.Ftree)).Flanguage, symbol)
}

func ts_node_string(tls *libc.TLS, _self TSNode) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	var alias_symbol TSSymbol
	_ = alias_symbol
	alias_symbol = uint16(ts_node__alias(tls, bp))
	return ts_subtree_string(tls, ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp))), alias_symbol, ts_language_symbol_metadata(tls, (*TSTree)(unsafe.Pointer((*(*TSNode)(unsafe.Pointer(bp))).Ftree)).Flanguage, alias_symbol).Fvisible, (*TSTree)(unsafe.Pointer((*(*TSNode)(unsafe.Pointer(bp))).Ftree)).Flanguage, libc.BoolUint8(0 != 0))
}

func ts_node_eq(tls *libc.TLS, self TSNode, other TSNode) (r uint8) {
	return libc.BoolUint8(self.Ftree == other.Ftree && self.Fid == other.Fid)
}

func ts_node_is_null(tls *libc.TLS, self TSNode) (r uint8) {
	return libc.BoolUint8(self.Fid == uintptr(0))
}

func ts_node_is_extra(tls *libc.TLS, self TSNode) (r uint8) {
	return ts_subtree_extra(tls, ts_node__subtree(tls, self))
}

func ts_node_is_named(tls *libc.TLS, _self TSNode) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	var alias TSSymbol
	var v1 int32
	_, _ = alias, v1
	alias = uint16(ts_node__alias(tls, bp))
	if alias != 0 {
		v1 = libc.Int32FromUint8(ts_language_symbol_metadata(tls, (*TSTree)(unsafe.Pointer((*(*TSNode)(unsafe.Pointer(bp))).Ftree)).Flanguage, alias).Fnamed)
	} else {
		v1 = libc.Int32FromUint8(ts_subtree_named(tls, ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp)))))
	}
	return libc.Uint8FromInt32(libc.BoolInt32(v1 != 0))
}

func ts_node_is_missing(tls *libc.TLS, self TSNode) (r uint8) {
	return ts_subtree_missing(tls, ts_node__subtree(tls, self))
}

func ts_node_has_changes(tls *libc.TLS, self TSNode) (r uint8) {
	return ts_subtree_has_changes(tls, ts_node__subtree(tls, self))
}

func ts_node_has_error(tls *libc.TLS, self TSNode) (r uint8) {
	return libc.BoolUint8(ts_subtree_error_cost(tls, ts_node__subtree(tls, self)) > uint32(0))
}

func ts_node_is_error(tls *libc.TLS, self TSNode) (r uint8) {
	var symbol TSSymbol
	_ = symbol
	symbol = ts_node_symbol(tls, self)
	return libc.BoolUint8(libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))))
}

func ts_node_descendant_count(tls *libc.TLS, self TSNode) (r uint32_t) {
	return ts_subtree_visible_descendant_count(tls, ts_node__subtree(tls, self)) + uint32(1)
}

func ts_node_parse_state(tls *libc.TLS, self TSNode) (r TSStateId) {
	return ts_subtree_parse_state(tls, ts_node__subtree(tls, self))
}

func ts_node_next_parse_state(tls *libc.TLS, self TSNode) (r TSStateId) {
	var language uintptr
	var state, symbol uint16_t
	_, _, _ = language, state, symbol
	language = (*TSTree)(unsafe.Pointer(self.Ftree)).Flanguage
	state = ts_node_parse_state(tls, self)
	if libc.Int32FromUint16(state) == libc.Int32FromInt32(0x7fff)*libc.Int32FromInt32(2)+libc.Int32FromInt32(1) {
		return libc.Uint16FromInt32(libc.Int32FromInt32(0x7fff)*libc.Int32FromInt32(2) + libc.Int32FromInt32(1))
	}
	symbol = ts_node_grammar_symbol(tls, self)
	return ts_language_next_state(tls, language, state, symbol)
}

func ts_node_parent(tls *libc.TLS, self TSNode) (r TSNode) {
	var next_node, node TSNode
	_, _ = next_node, node
	node = ts_tree_root_node(tls, self.Ftree)
	if node.Fid == self.Fid {
		return ts_node__null(tls)
	}
	for int32(1) != 0 {
		next_node = ts_node_child_with_descendant(tls, node, self)
		if next_node.Fid == self.Fid || ts_node_is_null(tls, next_node) != 0 {
			break
		}
		node = next_node
	}
	return node
}

func ts_node_child_with_descendant(tls *libc.TLS, _self TSNode, descendant TSNode) (r TSNode) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	var child, v3 TSNode
	var end_byte, start_byte uint32_t
	var is_empty uint8
	var v1 int32
	var _ /* iter at bp+32 */ NodeChildIterator
	_, _, _, _, _, _ = child, end_byte, is_empty, start_byte, v1, v3
	start_byte = ts_node_start_byte(tls, descendant)
	end_byte = ts_node_end_byte(tls, descendant)
	is_empty = libc.BoolUint8(start_byte == end_byte)
	for cond := true; cond; cond = !(ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp)), libc.BoolUint8(1 != 0)) != 0) {
		*(*NodeChildIterator)(unsafe.Pointer(bp + 32)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 32)) = ts_node_iterate_children(tls, bp)
		for {
			if !(ts_node_child_iterator_next(tls, bp+32, bp) != 0) || ts_node_start_byte(tls, *(*TSNode)(unsafe.Pointer(bp))) > start_byte {
				return ts_node__null(tls)
			}
			if (*(*TSNode)(unsafe.Pointer(bp))).Fid == descendant.Fid {
				return *(*TSNode)(unsafe.Pointer(bp))
			}
			if is_empty != 0 && (*(*NodeChildIterator)(unsafe.Pointer(bp + 32))).Fposition.Fbytes >= end_byte && ts_node_child_count(tls, *(*TSNode)(unsafe.Pointer(bp))) > uint32(0) {
				child = ts_node_child_with_descendant(tls, *(*TSNode)(unsafe.Pointer(bp)), descendant)
				if !(ts_node_is_null(tls, child) != 0) {
					if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp)), libc.BoolUint8(1 != 0)) != 0 {
						v3 = *(*TSNode)(unsafe.Pointer(bp))
					} else {
						v3 = child
					}
					return v3
				}
			}
			goto _2
		_2:
			;
			if is_empty != 0 {
				v1 = libc.BoolInt32((*(*NodeChildIterator)(unsafe.Pointer(bp + 32))).Fposition.Fbytes <= end_byte)
			} else {
				v1 = libc.BoolInt32((*(*NodeChildIterator)(unsafe.Pointer(bp + 32))).Fposition.Fbytes < end_byte)
			}
			if !(v1 != 0 || ts_node_child_count(tls, *(*TSNode)(unsafe.Pointer(bp))) == uint32(0)) {
				break
			}
		}
	}
	return *(*TSNode)(unsafe.Pointer(bp))
}

func ts_node_child(tls *libc.TLS, self TSNode, child_index uint32_t) (r TSNode) {
	return ts_node__child(tls, self, child_index, libc.BoolUint8(1 != 0))
}

func ts_node_named_child(tls *libc.TLS, self TSNode, child_index uint32_t) (r TSNode) {
	return ts_node__child(tls, self, child_index, libc.BoolUint8(0 != 0))
}

func ts_node_child_by_field_id(tls *libc.TLS, _self TSNode, field_id TSFieldId) (r TSNode) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	*(*TSNode)(unsafe.Pointer(bp)) = _self
	var index uint32_t
	var result TSNode
	var v1 Subtree
	var _ /* child at bp+48 */ TSNode
	var _ /* field_map at bp+32 */ uintptr
	var _ /* field_map_end at bp+40 */ uintptr
	var _ /* iterator at bp+80 */ NodeChildIterator
	_, _, _ = index, result, v1
	goto recur
recur:
	;
	if !(field_id != 0) || ts_node_child_count(tls, *(*TSNode)(unsafe.Pointer(bp))) == uint32(0) {
		return ts_node__null(tls)
	}
	v1 = ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp)))
	ts_language_field_map(tls, (*TSTree)(unsafe.Pointer((*(*TSNode)(unsafe.Pointer(bp))).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&v1)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), bp+32, bp+40)
	if *(*uintptr)(unsafe.Pointer(bp + 32)) == *(*uintptr)(unsafe.Pointer(bp + 40)) {
		return ts_node__null(tls)
	}
	for libc.Int32FromUint16((*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Ffield_id) < libc.Int32FromUint16(field_id) {
		*(*uintptr)(unsafe.Pointer(bp + 32)) += 4
		if *(*uintptr)(unsafe.Pointer(bp + 32)) == *(*uintptr)(unsafe.Pointer(bp + 40)) {
			return ts_node__null(tls)
		}
	}
	for libc.Int32FromUint16((*(*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)) + uintptr(-libc.Int32FromInt32(1))*4))).Ffield_id) > libc.Int32FromUint16(field_id) {
		*(*uintptr)(unsafe.Pointer(bp + 40)) -= 4
		if *(*uintptr)(unsafe.Pointer(bp + 32)) == *(*uintptr)(unsafe.Pointer(bp + 40)) {
			return ts_node__null(tls)
		}
	}
	*(*NodeChildIterator)(unsafe.Pointer(bp + 80)) = NodeChildIterator{}
	*(*struct {
		Fparent struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		}
		Ftree     uintptr
		Fposition struct {
			Fbytes  uint32
			Fextent TSPoint
		}
		Fchild_index            uint32
		Fstructural_child_index uint32
		Falias_sequence         uintptr
	})(unsafe.Pointer(bp + 80)) = ts_node_iterate_children(tls, bp)
	for ts_node_child_iterator_next(tls, bp+80, bp+48) != 0 {
		if !(ts_subtree_extra(tls, ts_node__subtree(tls, *(*TSNode)(unsafe.Pointer(bp + 48)))) != 0) {
			index = (*(*NodeChildIterator)(unsafe.Pointer(bp + 80))).Fstructural_child_index - uint32(1)
			if index < uint32((*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Fchild_index) {
				continue
			}
			if (*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Finherited != 0 {
				if *(*uintptr)(unsafe.Pointer(bp + 32))+uintptr(1)*4 == *(*uintptr)(unsafe.Pointer(bp + 40)) {
					*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 48))
					goto recur
				} else {
					result = ts_node_child_by_field_id(tls, *(*TSNode)(unsafe.Pointer(bp + 48)), field_id)
					if result.Fid != 0 {
						return result
					}
					*(*uintptr)(unsafe.Pointer(bp + 32)) += 4
					if *(*uintptr)(unsafe.Pointer(bp + 32)) == *(*uintptr)(unsafe.Pointer(bp + 40)) {
						return ts_node__null(tls)
					}
				}
			} else {
				if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp + 48)), libc.BoolUint8(1 != 0)) != 0 {
					return *(*TSNode)(unsafe.Pointer(bp + 48))
				} else {
					if ts_node_child_count(tls, *(*TSNode)(unsafe.Pointer(bp + 48))) > uint32(0) {
						return ts_node_child(tls, *(*TSNode)(unsafe.Pointer(bp + 48)), uint32(0))
					} else {
						*(*uintptr)(unsafe.Pointer(bp + 32)) += 4
						if *(*uintptr)(unsafe.Pointer(bp + 32)) == *(*uintptr)(unsafe.Pointer(bp + 40)) {
							return ts_node__null(tls)
						}
					}
				}
			}
		}
	}
	return ts_node__null(tls)
}

func ts_node__field_name_from_language(tls *libc.TLS, self TSNode, structural_child_index uint32_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 Subtree
	var _ /* field_map at bp+0 */ uintptr
	var _ /* field_map_end at bp+8 */ uintptr
	_ = v1
	v1 = ts_node__subtree(tls, self)
	ts_language_field_map(tls, (*TSTree)(unsafe.Pointer(self.Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&v1)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), bp, bp+8)
	for {
		if !(*(*uintptr)(unsafe.Pointer(bp)) != *(*uintptr)(unsafe.Pointer(bp + 8))) {
			break
		}
		if !((*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Finherited != 0) && uint32((*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_index) == structural_child_index {
			return *(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSTree)(unsafe.Pointer(self.Ftree)).Flanguage)).Ffield_names + uintptr((*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Ffield_id)*8))
		}
		goto _2
	_2:
		;
		*(*uintptr)(unsafe.Pointer(bp)) += 4
	}
	return libc.UintptrFromInt32(0)
}

func ts_node_field_name_for_child(tls *libc.TLS, self TSNode, child_index uint32_t) (r uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var did_descend uint8
	var field_name, field_name1, inherited_field_name uintptr
	var grandchild_count, grandchild_index, index uint32_t
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* result at bp+0 */ TSNode
	_, _, _, _, _, _, _ = did_descend, field_name, field_name1, grandchild_count, grandchild_index, index, inherited_field_name
	*(*TSNode)(unsafe.Pointer(bp)) = self
	did_descend = libc.BoolUint8(1 != 0)
	inherited_field_name = libc.UintptrFromInt32(0)
	for did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		index = uint32(0)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), libc.BoolUint8(1 != 0)) != 0 {
				if index == child_index {
					if ts_node_is_extra(tls, *(*TSNode)(unsafe.Pointer(bp + 32))) != 0 {
						return libc.UintptrFromInt32(0)
					}
					field_name = ts_node__field_name_from_language(tls, *(*TSNode)(unsafe.Pointer(bp)), (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fstructural_child_index-uint32(1))
					if field_name != 0 {
						return field_name
					}
					return inherited_field_name
				}
				index = index + 1
			} else {
				grandchild_index = child_index - index
				grandchild_count = ts_node__relevant_child_count(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), libc.BoolUint8(1 != 0))
				if grandchild_index < grandchild_count {
					field_name1 = ts_node__field_name_from_language(tls, *(*TSNode)(unsafe.Pointer(bp)), (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fstructural_child_index-uint32(1))
					if field_name1 != 0 {
						inherited_field_name = field_name1
					}
					did_descend = libc.BoolUint8(1 != 0)
					*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 32))
					child_index = grandchild_index
					break
				}
				index = index + grandchild_count
			}
		}
	}
	return libc.UintptrFromInt32(0)
}

func ts_node_field_name_for_named_child(tls *libc.TLS, self TSNode, named_child_index uint32_t) (r uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var did_descend uint8
	var field_name, field_name1, inherited_field_name uintptr
	var grandchild_count, index, named_grandchild_index uint32_t
	var _ /* child at bp+32 */ TSNode
	var _ /* iterator at bp+64 */ NodeChildIterator
	var _ /* result at bp+0 */ TSNode
	_, _, _, _, _, _, _ = did_descend, field_name, field_name1, grandchild_count, index, inherited_field_name, named_grandchild_index
	*(*TSNode)(unsafe.Pointer(bp)) = self
	did_descend = libc.BoolUint8(1 != 0)
	inherited_field_name = libc.UintptrFromInt32(0)
	for did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		index = uint32(0)
		*(*NodeChildIterator)(unsafe.Pointer(bp + 64)) = NodeChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 64)) = ts_node_iterate_children(tls, bp)
		for ts_node_child_iterator_next(tls, bp+64, bp+32) != 0 {
			if ts_node__is_relevant(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), libc.BoolUint8(0 != 0)) != 0 {
				if index == named_child_index {
					if ts_node_is_extra(tls, *(*TSNode)(unsafe.Pointer(bp + 32))) != 0 {
						return libc.UintptrFromInt32(0)
					}
					field_name = ts_node__field_name_from_language(tls, *(*TSNode)(unsafe.Pointer(bp)), (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fstructural_child_index-uint32(1))
					if field_name != 0 {
						return field_name
					}
					return inherited_field_name
				}
				index = index + 1
			} else {
				named_grandchild_index = named_child_index - index
				grandchild_count = ts_node__relevant_child_count(tls, *(*TSNode)(unsafe.Pointer(bp + 32)), libc.BoolUint8(0 != 0))
				if named_grandchild_index < grandchild_count {
					field_name1 = ts_node__field_name_from_language(tls, *(*TSNode)(unsafe.Pointer(bp)), (*(*NodeChildIterator)(unsafe.Pointer(bp + 64))).Fstructural_child_index-uint32(1))
					if field_name1 != 0 {
						inherited_field_name = field_name1
					}
					did_descend = libc.BoolUint8(1 != 0)
					*(*TSNode)(unsafe.Pointer(bp)) = *(*TSNode)(unsafe.Pointer(bp + 32))
					named_child_index = named_grandchild_index
					break
				}
				index = index + grandchild_count
			}
		}
	}
	return libc.UintptrFromInt32(0)
}

func ts_node_child_by_field_name(tls *libc.TLS, self TSNode, name uintptr, name_length uint32_t) (r TSNode) {
	var field_id TSFieldId
	_ = field_id
	field_id = ts_language_field_id_for_name(tls, (*TSTree)(unsafe.Pointer(self.Ftree)).Flanguage, name, name_length)
	return ts_node_child_by_field_id(tls, self, field_id)
}

func ts_node_child_count(tls *libc.TLS, self TSNode) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* tree at bp+0 */ Subtree
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = ts_node__subtree(tls, self)
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count
	} else {
		return uint32(0)
	}
	return r
}

func ts_node_named_child_count(tls *libc.TLS, self TSNode) (r uint32_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* tree at bp+0 */ Subtree
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = ts_node__subtree(tls, self)
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
		return (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count
	} else {
		return uint32(0)
	}
	return r
}

func ts_node_next_sibling(tls *libc.TLS, self TSNode) (r TSNode) {
	return ts_node__next_sibling(tls, self, libc.BoolUint8(1 != 0))
}

func ts_node_next_named_sibling(tls *libc.TLS, self TSNode) (r TSNode) {
	return ts_node__next_sibling(tls, self, libc.BoolUint8(0 != 0))
}

func ts_node_prev_sibling(tls *libc.TLS, self TSNode) (r TSNode) {
	return ts_node__prev_sibling(tls, self, libc.BoolUint8(1 != 0))
}

func ts_node_prev_named_sibling(tls *libc.TLS, self TSNode) (r TSNode) {
	return ts_node__prev_sibling(tls, self, libc.BoolUint8(0 != 0))
}

func ts_node_first_child_for_byte(tls *libc.TLS, self TSNode, byte1 uint32_t) (r TSNode) {
	return ts_node__first_child_for_byte(tls, self, byte1, libc.BoolUint8(1 != 0))
}

func ts_node_first_named_child_for_byte(tls *libc.TLS, self TSNode, byte1 uint32_t) (r TSNode) {
	return ts_node__first_child_for_byte(tls, self, byte1, libc.BoolUint8(0 != 0))
}

func ts_node_descendant_for_byte_range(tls *libc.TLS, self TSNode, start uint32_t, end uint32_t) (r TSNode) {
	return ts_node__descendant_for_byte_range(tls, self, start, end, libc.BoolUint8(1 != 0))
}

func ts_node_named_descendant_for_byte_range(tls *libc.TLS, self TSNode, start uint32_t, end uint32_t) (r TSNode) {
	return ts_node__descendant_for_byte_range(tls, self, start, end, libc.BoolUint8(0 != 0))
}

func ts_node_descendant_for_point_range(tls *libc.TLS, self TSNode, start TSPoint, end TSPoint) (r TSNode) {
	return ts_node__descendant_for_point_range(tls, self, start, end, libc.BoolUint8(1 != 0))
}

func ts_node_named_descendant_for_point_range(tls *libc.TLS, self TSNode, start TSPoint, end TSPoint) (r TSNode) {
	return ts_node__descendant_for_point_range(tls, self, start, end, libc.BoolUint8(0 != 0))
}

func ts_node_edit(tls *libc.TLS, self uintptr, edit uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* start_byte at bp+0 */ uint32_t
	var _ /* start_point at bp+8 */ TSPoint
	*(*uint32_t)(unsafe.Pointer(bp)) = ts_node_start_byte(tls, *(*TSNode)(unsafe.Pointer(self)))
	*(*TSPoint)(unsafe.Pointer(bp + 8)) = ts_node_start_point(tls, *(*TSNode)(unsafe.Pointer(self)))
	ts_point_edit(tls, bp+8, bp, edit)
	*(*uint32_t)(unsafe.Pointer(self)) = *(*uint32_t)(unsafe.Pointer(bp))
	*(*uint32_t)(unsafe.Pointer(self + 1*4)) = (*(*TSPoint)(unsafe.Pointer(bp + 8))).Frow
	*(*uint32_t)(unsafe.Pointer(self + 2*4)) = (*(*TSPoint)(unsafe.Pointer(bp + 8))).Fcolumn
}

type __gwchar_t = int32

type imaxdiv_t = struct {
	Fquot int64
	Frem  int64
}

type ReduceAction = struct {
	Fcount              uint32_t
	Fsymbol             TSSymbol
	Fdynamic_precedence int32
	Fproduction_id      uint16
}

type ReduceActionSet = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

func ts_reduce_action_set_add(tls *libc.TLS, self uintptr, new_action ReduceAction) {
	var action ReduceAction
	var i, v2 uint32_t
	var v3 uintptr
	_, _, _, _ = action, i, v2, v3
	i = uint32(0)
	for {
		if !(i < (*ReduceActionSet)(unsafe.Pointer(self)).Fsize) {
			break
		}
		action = *(*ReduceAction)(unsafe.Pointer((*ReduceActionSet)(unsafe.Pointer(self)).Fcontents + uintptr(i)*16))
		if libc.Int32FromUint16(action.Fsymbol) == libc.Int32FromUint16(new_action.Fsymbol) && action.Fcount == new_action.Fcount {
			return
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	(*ReduceActionSet)(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*ReduceActionSet)(unsafe.Pointer(self)).Fcontents, (*ReduceActionSet)(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(16))
	v3 = self + 8
	v2 = *(*uint32_t)(unsafe.Pointer(v3))
	*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
	*(*ReduceAction)(unsafe.Pointer((*ReduceActionSet)(unsafe.Pointer(self)).Fcontents + uintptr(v2)*16)) = new_action
}

type StackEntry = struct {
	Ftree        Subtree
	Fchild_index uint32_t
	Fbyte_offset uint32_t
}

type ReusableNode = struct {
	Fstack struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Flast_external_token Subtree
}

func reusable_node_new(tls *libc.TLS) (r ReusableNode) {
	return ReusableNode{
		Flast_external_token: Subtree{},
	}
}

func reusable_node_clear(tls *libc.TLS, self uintptr) {
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	(*ReusableNode)(unsafe.Pointer(self)).Flast_external_token = Subtree{}
}

func reusable_node_tree(tls *libc.TLS, self uintptr) (r Subtree) {
	var v1 Subtree
	_ = v1
	if (*ReusableNode)(unsafe.Pointer(self)).Fstack.Fsize > uint32(0) {
		v1 = (*(*StackEntry)(unsafe.Pointer((*ReusableNode)(unsafe.Pointer(self)).Fstack.Fcontents + uintptr((*ReusableNode)(unsafe.Pointer(self)).Fstack.Fsize-uint32(1))*16))).Ftree
	} else {
		v1 = Subtree{}
	}
	return v1
}

func reusable_node_byte_offset(tls *libc.TLS, self uintptr) (r uint32_t) {
	var v1 uint32
	_ = v1
	if (*ReusableNode)(unsafe.Pointer(self)).Fstack.Fsize > uint32(0) {
		v1 = (*(*StackEntry)(unsafe.Pointer((*ReusableNode)(unsafe.Pointer(self)).Fstack.Fcontents + uintptr((*ReusableNode)(unsafe.Pointer(self)).Fstack.Fsize-uint32(1))*16))).Fbyte_offset
	} else {
		v1 = uint32(4294967295)
	}
	return v1
}

func reusable_node_delete(tls *libc.TLS, self uintptr) {
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcapacity = uint32(0)
}

func reusable_node_advance(tls *libc.TLS, self uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var byte_offset, next_index, v1 uint32_t
	var last_entry, popped_entry StackEntry
	var v2, v4 uintptr
	var _ /* tree at bp+0 */ Subtree
	_, _, _, _, _, _, _ = byte_offset, last_entry, next_index, popped_entry, v1, v2, v4
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+929, int32(40), uintptr(unsafe.Pointer(&__func__14)))
		}
	}
	last_entry = *(*StackEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize-uint32(1))*16))
	byte_offset = last_entry.Fbyte_offset + ts_subtree_total_bytes(tls, last_entry.Ftree)
	if ts_subtree_has_external_tokens(tls, last_entry.Ftree) != 0 {
		(*ReusableNode)(unsafe.Pointer(self)).Flast_external_token = ts_subtree_last_external_token(tls, last_entry.Ftree)
	}
	for cond := true; cond; cond = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) <= next_index {
		v2 = self + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		popped_entry = *(*StackEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*16))
		next_index = popped_entry.Fchild_index + uint32(1)
		if (*ReusableNode)(unsafe.Pointer(self)).Fstack.Fsize == uint32(0) {
			return
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+929, int32(52), uintptr(unsafe.Pointer(&__func__14)))
			}
		}
		*(*Subtree)(unsafe.Pointer(bp)) = (*StackEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize-uint32(1))*16)).Ftree
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(16))
	v2 = self + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v4 = libc.UintptrFromInt32(0)
	} else {
		v4 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
	}
	*(*StackEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*16)) = StackEntry{
		Ftree:        *(*Subtree)(unsafe.Pointer(v4 + uintptr(next_index)*8)),
		Fchild_index: next_index,
		Fbyte_offset: byte_offset,
	}
}

var __func__14 = [22]int8{'r', 'e', 'u', 's', 'a', 'b', 'l', 'e', '_', 'n', 'o', 'd', 'e', '_', 'a', 'd', 'v', 'a', 'n', 'c', 'e'}

func reusable_node_descend(tls *libc.TLS, self uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 uint32_t
	var v2, v3 uintptr
	var _ /* last_entry at bp+0 */ StackEntry
	_, _, _ = v1, v2, v3
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+929, int32(63), uintptr(unsafe.Pointer(&__func__15)))
		}
	}
	*(*StackEntry)(unsafe.Pointer(bp)) = StackEntry{}
	*(*struct {
		Ftree struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		}
		Fchild_index uint32
		Fbyte_offset uint32
	})(unsafe.Pointer(bp)) = *(*StackEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize-uint32(1))*16))
	if ts_subtree_child_count(tls, (*(*StackEntry)(unsafe.Pointer(bp))).Ftree) > uint32(0) {
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(16))
		v2 = self + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
			v3 = libc.UintptrFromInt32(0)
		} else {
			v3 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
		}
		*(*StackEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*16)) = StackEntry{
			Ftree:        *(*Subtree)(unsafe.Pointer(v3)),
			Fbyte_offset: (*(*StackEntry)(unsafe.Pointer(bp))).Fbyte_offset,
		}
		return libc.BoolUint8(1 != 0)
	} else {
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var __func__15 = [22]int8{'r', 'e', 'u', 's', 'a', 'b', 'l', 'e', '_', 'n', 'o', 'd', 'e', '_', 'd', 'e', 's', 'c', 'e', 'n', 'd'}

func reusable_node_advance_past_leaf(tls *libc.TLS, self uintptr) {
	for reusable_node_descend(tls, self) != 0 {
	}
	reusable_node_advance(tls, self)
}

func reusable_node_reset(tls *libc.TLS, self uintptr, tree Subtree) {
	var v1 uint32_t
	var v2 uintptr
	_, _ = v1, v2
	reusable_node_clear(tls, self)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(16))
	v2 = self + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*StackEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*16)) = StackEntry{
		Ftree: tree,
	}
	if !(reusable_node_descend(tls, self) != 0) {
		reusable_node_clear(tls, self)
	}
}

type Stack = struct {
	Fheads struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fslices    StackSliceArray
	Fiterators struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fnode_pool    StackNodeArray
	Fbase_node    uintptr
	Fsubtree_pool uintptr
}

type StackVersion = uint32

type StackSlice = struct {
	Fsubtrees SubtreeArray
	Fversion  StackVersion
}

type StackSliceArray = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type StackSummaryEntry = struct {
	Fposition Length
	Fdepth    uint32
	Fstate    TSStateId
}

type StackSummary = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

var MAX_VERSION_COUNT = uint32(6)
var MAX_VERSION_COUNT_OVERFLOW = uint32(4)
var MAX_SUMMARY_DEPTH = uint32(16)
var MAX_COST_DIFFERENCE = libc.Uint32FromInt32(libc.Int32FromInt32(18) * libc.Int32FromInt32(100))
var OP_COUNT_PER_PARSER_CALLBACK_CHECK = uint32(100)

type TokenCache = struct {
	Ftoken               Subtree
	Flast_external_token Subtree
	Fbyte_index          uint32_t
}

type ErrorStatus = struct {
	Fcost               uint32
	Fnode_count         uint32
	Fdynamic_precedence int32
	Fis_in_error        uint8
}

type ErrorComparison = int32

const ErrorComparisonTakeLeft = 0
const ErrorComparisonPreferLeft = 1
const ErrorComparisonNone = 2
const ErrorComparisonPreferRight = 3
const ErrorComparisonTakeRight = 4

type TSStringInput = struct {
	Fstring1 uintptr
	Flength  uint32_t
}

func ts_string_input_read(tls *libc.TLS, _self uintptr, byte1 uint32_t, point TSPoint, length uintptr) (r uintptr) {
	var self uintptr
	_ = self
	_ = point
	self = _self
	if byte1 >= (*TSStringInput)(unsafe.Pointer(self)).Flength {
		*(*uint32_t)(unsafe.Pointer(length)) = uint32(0)
		return __ccgo_ts + 981
	} else {
		*(*uint32_t)(unsafe.Pointer(length)) = (*TSStringInput)(unsafe.Pointer(self)).Flength - byte1
		return (*TSStringInput)(unsafe.Pointer(self)).Fstring1 + uintptr(byte1)
	}
	return r
}

func ts_parser__log(tls *libc.TLS, self uintptr) {
	var chr uintptr
	_ = chr
	if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 {
		(*(*func(*libc.TLS, uintptr, TSLogType1, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog})))(tls, (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Fpayload, int32(TSLogTypeParse), self+192)
	}
	if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
		libc.Xfprintf(tls, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file, __ccgo_ts+982, 0)
		chr = self + 192
		for {
			if !(int32(*(*int8)(unsafe.Pointer(chr))) != 0) {
				break
			}
			if int32(*(*int8)(unsafe.Pointer(chr))) == int32('"') || int32(*(*int8)(unsafe.Pointer(chr))) == int32('\\') {
				libc.Xfputc(tls, int32('\\'), (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
			}
			libc.Xfputc(tls, int32(*(*int8)(unsafe.Pointer(chr))), (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
			goto _1
		_1:
			;
			chr = chr + 1
		}
		libc.Xfprintf(tls, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file, __ccgo_ts+998, 0)
	}
}

func ts_parser__breakdown_top_of_stack(tls *libc.TLS, self uintptr, version StackVersion) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var child, tree Subtree
	var did_break_down, pending uint8
	var i, j, j1, n uint32_t
	var state TSStateId
	var v3 uintptr
	var _ /* parent at bp+40 */ Subtree
	var _ /* pop at bp+0 */ StackSliceArray
	var _ /* slice at bp+16 */ StackSlice
	_, _, _, _, _, _, _, _, _, _ = child, did_break_down, i, j, j1, n, pending, state, tree, v3
	did_break_down = libc.BoolUint8(0 != 0)
	pending = libc.BoolUint8(0 != 0)
	for cond := true; cond; cond = pending != 0 {
		*(*StackSliceArray)(unsafe.Pointer(bp)) = ts_stack_pop_pending(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
		if !((*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize != 0) {
			break
		}
		did_break_down = libc.BoolUint8(1 != 0)
		pending = libc.BoolUint8(0 != 0)
		i = uint32(0)
		for {
			if !(i < (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+1004, __ccgo_ts+1033, int32(190), uintptr(unsafe.Pointer(&__func__16)))
				}
			}
			*(*StackSlice)(unsafe.Pointer(bp + 16)) = *(*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents + uintptr(i)*24))
			state = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion)
			*(*Subtree)(unsafe.Pointer(bp + 40)) = Subtree{}
			_ = libc.Uint64FromInt64(4)
			{
				if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*SubtreeArray)(unsafe.Pointer(bp+16)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+1076, __ccgo_ts+1033, int32(192), uintptr(unsafe.Pointer(&__func__16)))
				}
			}
			*(*struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			})(unsafe.Pointer(bp + 40)) = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(bp + 16)).Fcontents))
			j = uint32(0)
			n = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 40)))
			for {
				if !(j < n) {
					break
				}
				if int32(*(*uint8)(unsafe.Pointer(bp + 40 + 0))&0x1>>0) != 0 {
					v3 = libc.UintptrFromInt32(0)
				} else {
					v3 = *(*uintptr)(unsafe.Pointer(bp + 40)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)))).Fchild_count)*8
				}
				child = *(*Subtree)(unsafe.Pointer(v3 + uintptr(j)*8))
				pending = libc.BoolUint8(ts_subtree_child_count(tls, child) > uint32(0))
				if ts_subtree_is_error(tls, child) != 0 {
					state = uint16(0)
				} else {
					if !(ts_subtree_extra(tls, child) != 0) {
						state = ts_language_next_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, state, ts_subtree_symbol(tls, child))
					}
				}
				ts_subtree_retain(tls, child)
				ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion, child, pending, state)
				goto _2
			_2:
				;
				j = j + 1
			}
			j1 = uint32(1)
			for {
				if !(j1 < (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fsubtrees.Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(j1 < (*SubtreeArray)(unsafe.Pointer(bp+16)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+1116, __ccgo_ts+1033, int32(209), uintptr(unsafe.Pointer(&__func__16)))
					}
				}
				tree = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents + uintptr(j1)*8))
				ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion, tree, libc.BoolUint8(0 != 0), state)
				goto _4
			_4:
				;
				j1 = j1 + 1
			}
			ts_subtree_release(tls, self+1224, *(*Subtree)(unsafe.Pointer(bp + 40)))
			if (*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents)
			}
			(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fcontents = libc.UintptrFromInt32(0)
			(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fsize = uint32(0)
			(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fcapacity = uint32(0)
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1156, libc.VaList(bp+56, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 40))))))
				ts_parser__log(tls, self)
			}
			if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				ts_stack_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
				libc.Xfputs(tls, __ccgo_ts+1187, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	return did_break_down
}

var __func__16 = [34]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'b', 'r', 'e', 'a', 'k', 'd', 'o', 'w', 'n', '_', 't', 'o', 'p', '_', 'o', 'f', '_', 's', 't', 'a', 'c', 'k'}

func ts_parser__breakdown_lookahead(tls *libc.TLS, self uintptr, lookahead uintptr, state TSStateId, reusable_node uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var did_descend uint8
	var tree Subtree
	_, _ = did_descend, tree
	did_descend = libc.BoolUint8(0 != 0)
	tree = reusable_node_tree(tls, reusable_node)
	for ts_subtree_child_count(tls, tree) > uint32(0) && libc.Int32FromUint16(ts_subtree_parse_state(tls, tree)) != libc.Int32FromUint16(state) {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1190, libc.VaList(bp+8, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, tree))))
			ts_parser__log(tls, self)
		}
		reusable_node_descend(tls, reusable_node)
		tree = reusable_node_tree(tls, reusable_node)
		did_descend = libc.BoolUint8(1 != 0)
	}
	if did_descend != 0 {
		ts_subtree_release(tls, self+1224, *(*Subtree)(unsafe.Pointer(lookahead)))
		*(*Subtree)(unsafe.Pointer(lookahead)) = tree
		ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer(lookahead)))
	}
}

func ts_parser__compare_versions(tls *libc.TLS, self uintptr, a ErrorStatus, b ErrorStatus) (r ErrorComparison) {
	_ = self
	if !(a.Fis_in_error != 0) && b.Fis_in_error != 0 {
		if a.Fcost < b.Fcost {
			return int32(ErrorComparisonTakeLeft)
		} else {
			return int32(ErrorComparisonPreferLeft)
		}
	}
	if a.Fis_in_error != 0 && !(b.Fis_in_error != 0) {
		if b.Fcost < a.Fcost {
			return int32(ErrorComparisonTakeRight)
		} else {
			return int32(ErrorComparisonPreferRight)
		}
	}
	if a.Fcost < b.Fcost {
		if (b.Fcost-a.Fcost)*(uint32(1)+a.Fnode_count) > MAX_COST_DIFFERENCE {
			return int32(ErrorComparisonTakeLeft)
		} else {
			return int32(ErrorComparisonPreferLeft)
		}
	}
	if b.Fcost < a.Fcost {
		if (a.Fcost-b.Fcost)*(uint32(1)+b.Fnode_count) > MAX_COST_DIFFERENCE {
			return int32(ErrorComparisonTakeRight)
		} else {
			return int32(ErrorComparisonPreferRight)
		}
	}
	if a.Fdynamic_precedence > b.Fdynamic_precedence {
		return int32(ErrorComparisonPreferLeft)
	}
	if b.Fdynamic_precedence > a.Fdynamic_precedence {
		return int32(ErrorComparisonPreferRight)
	}
	return int32(ErrorComparisonNone)
}

func ts_parser__version_status(tls *libc.TLS, self uintptr, version StackVersion) (r ErrorStatus) {
	var cost uint32
	var is_paused uint8
	_, _ = cost, is_paused
	cost = ts_stack_error_cost(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	is_paused = ts_stack_is_paused(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	if is_paused != 0 {
		cost = cost + uint32(100)
	}
	return ErrorStatus{
		Fcost:               cost,
		Fnode_count:         ts_stack_node_count_since_error(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version),
		Fdynamic_precedence: ts_stack_dynamic_precedence(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version),
		Fis_in_error:        libc.BoolUint8(is_paused != 0 || libc.Int32FromUint16(ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)) == 0),
	}
}

func ts_parser__better_version_exists(tls *libc.TLS, self uintptr, version StackVersion, is_in_error uint8, cost uint32) (r uint8) {
	var i, n StackVersion
	var position Length
	var status, status_i ErrorStatus
	_, _, _, _, _ = i, n, position, status, status_i
	if *(*uintptr)(unsafe.Pointer(self + 1288)) != 0 && ts_subtree_error_cost(tls, (*TSParser)(unsafe.Pointer(self)).Ffinished_tree) <= cost {
		return libc.BoolUint8(1 != 0)
	}
	position = ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	status = ErrorStatus{
		Fcost:               cost,
		Fnode_count:         ts_stack_node_count_since_error(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version),
		Fdynamic_precedence: ts_stack_dynamic_precedence(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version),
		Fis_in_error:        is_in_error,
	}
	i = uint32(0)
	n = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	for {
		if !(i < n) {
			break
		}
		if i == version || !(ts_stack_is_active(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i) != 0) || ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i).Fbytes < position.Fbytes {
			goto _1
		}
		status_i = ts_parser__version_status(tls, self, i)
		switch ts_parser__compare_versions(tls, self, status, status_i) {
		case int32(ErrorComparisonTakeRight):
			return libc.BoolUint8(1 != 0)
		case int32(ErrorComparisonPreferRight):
			if ts_stack_can_merge(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i, version) != 0 {
				return libc.BoolUint8(1 != 0)
			}
		default:
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(0 != 0)
}

func ts_parser__call_main_lex_fn(tls *libc.TLS, self uintptr, lex_mode TSLexerMode) (r uint8) {
	if ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		return ts_wasm_store_call_lex_main(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store, lex_mode.Flex_state)
	} else {
		return (*(*func(*libc.TLS, uintptr, TSStateId) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Flex_fn})))(tls, self, lex_mode.Flex_state)
	}
	return r
}

func ts_parser__call_keyword_lex_fn(tls *libc.TLS, self uintptr) (r uint8) {
	if ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		return ts_wasm_store_call_lex_keyword(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store, uint16(0))
	} else {
		return (*(*func(*libc.TLS, uintptr, TSStateId) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fkeyword_lex_fn})))(tls, self, uint16(0))
	}
	return r
}

func ts_parser__external_scanner_create(tls *libc.TLS, self uintptr) {
	if (*TSParser)(unsafe.Pointer(self)).Flanguage != 0 && (*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fstates != 0 {
		if ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
			(*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload = uintptr(uint64(ts_wasm_store_call_scanner_create(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store)))
			if ts_wasm_store_has_error(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store) != 0 {
				(*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error = libc.BoolUint8(1 != 0)
			}
		} else {
			if (*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fcreate != 0 {
				(*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload = (*(*func(*libc.TLS) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fcreate})))(tls)
			}
		}
	}
}

func ts_parser__external_scanner_destroy(tls *libc.TLS, self uintptr) {
	if (*TSParser)(unsafe.Pointer(self)).Flanguage != 0 && (*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload != 0 && (*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fdestroy != 0 && !(ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fdestroy})))(tls, (*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload)
	}
	(*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload = libc.UintptrFromInt32(0)
}

func ts_parser__external_scanner_serialize(tls *libc.TLS, self uintptr) (r uint32) {
	var length uint32_t
	_ = length
	if ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		length = ts_wasm_store_call_scanner_serialize(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store, uint32(uint64((*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload)), self+192)
		if ts_wasm_store_has_error(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store) != 0 {
			(*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error = libc.BoolUint8(1 != 0)
		}
	} else {
		length = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fserialize})))(tls, (*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload, self+192)
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(length <= libc.Uint32FromInt32(1024)) {
			libc.X__assert_fail(tls, __ccgo_ts+1212, __ccgo_ts+1033, int32(409), uintptr(unsafe.Pointer(&__func__17)))
		}
	}
	return length
}

var __func__17 = [38]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 's', 'c', 'a', 'n', 'n', 'e', 'r', '_', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e'}

func ts_parser__external_scanner_deserialize(tls *libc.TLS, self uintptr, _external_token Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _external_token
	var data uintptr
	var length uint32_t
	_, _ = data, length
	data = libc.UintptrFromInt32(0)
	length = uint32(0)
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		data = ts_external_scanner_state_data(tls, *(*uintptr)(unsafe.Pointer(bp))+48)
		length = (*(*ExternalScannerState)(unsafe.Add(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))), 48))).Flength
	}
	if ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		ts_wasm_store_call_scanner_deserialize(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store, uint32(uint64((*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload)), data, length)
		if ts_wasm_store_has_error(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store) != 0 {
			(*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error = libc.BoolUint8(1 != 0)
		}
	} else {
		(*(*func(*libc.TLS, uintptr, uintptr, uint32))(unsafe.Pointer(&struct{ uintptr }{(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fdeserialize})))(tls, (*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload, data, length)
	}
}

func ts_parser__external_scanner_scan(tls *libc.TLS, self uintptr, external_lex_state TSStateId) (r uint8) {
	var result uint8
	var valid_external_tokens uintptr
	_, _ = result, valid_external_tokens
	if ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		result = ts_wasm_store_call_scanner_scan(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store, uint32(uint64((*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload)), uint32(external_lex_state)*(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_token_count)
		if ts_wasm_store_has_error(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store) != 0 {
			(*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error = libc.BoolUint8(1 != 0)
		}
		return result
	} else {
		valid_external_tokens = ts_language_enabled_external_tokens(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, uint32(external_lex_state))
		return (*(*func(*libc.TLS, uintptr, uintptr, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fscan})))(tls, (*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload, self, valid_external_tokens)
	}
	return r
}

func ts_parser__can_reuse_first_leaf(tls *libc.TLS, self uintptr, state TSStateId, tree Subtree, table_entry uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var leaf_state TSStateId
	var leaf_symbol TSSymbol
	var _ /* current_lex_mode at bp+0 */ TSLexerMode
	var _ /* leaf_lex_mode at bp+6 */ TSLexerMode
	_, _ = leaf_state, leaf_symbol
	leaf_symbol = ts_subtree_leaf_symbol(tls, tree)
	leaf_state = ts_subtree_leaf_parse_state(tls, tree)
	*(*TSLexerMode)(unsafe.Pointer(bp)) = ts_language_lex_mode_for_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, state)
	*(*TSLexerMode)(unsafe.Pointer(bp + 6)) = ts_language_lex_mode_for_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, leaf_state)
	if libc.Int32FromUint16((*(*TSLexerMode)(unsafe.Pointer(bp))).Flex_state) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) {
		return libc.BoolUint8(0 != 0)
	}
	if (*TableEntry)(unsafe.Pointer(table_entry)).Faction_count > uint32(0) && libc.Xmemcmp(tls, bp+6, bp, uint64(6)) == 0 && (libc.Int32FromUint16(leaf_symbol) != libc.Int32FromUint16((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fkeyword_capture_token) || !(ts_subtree_is_keyword(tls, tree) != 0) && libc.Int32FromUint16(ts_subtree_parse_state(tls, tree)) == libc.Int32FromUint16(state)) {
		return libc.BoolUint8(1 != 0)
	}
	if ts_subtree_size(tls, tree).Fbytes == uint32(0) && libc.Int32FromUint16(leaf_symbol) != 0 {
		return libc.BoolUint8(0 != 0)
	}
	return libc.BoolUint8(libc.Int32FromUint16((*(*TSLexerMode)(unsafe.Pointer(bp))).Fexternal_lex_state) == 0 && (*TableEntry)(unsafe.Pointer(table_entry)).Fis_reusable != 0)
}

func ts_parser__lex(tls *libc.TLS, self uintptr, version StackVersion, parse_state TSStateId) (r Subtree) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var buf, symbol2 uintptr
	var called_get_column, error_mode, external_scanner_state_changed, found_external_token, found_token, is_keyword, skipped_error, token_is_extra uint8
	var column_data ColumnData
	var current_position, error_end_position, error_start_position, padding, padding1, size, size1, start_position Length
	var end_byte, external_scanner_state_len, lookahead_bytes, lookahead_bytes1 uint32_t
	var external_token, result Subtree
	var first_error_character int32_t
	var i, off, v3 int32
	var lex_mode TSLexerMode
	var next_parse_state TSStateId
	var symbol, symbol1 TSSymbol
	var _ /* lookahead_end_byte at bp+0 */ uint32_t
	var _ /* mut_result at bp+8 */ MutableSubtree
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = buf, called_get_column, column_data, current_position, end_byte, error_end_position, error_mode, error_start_position, external_scanner_state_changed, external_scanner_state_len, external_token, first_error_character, found_external_token, found_token, i, is_keyword, lex_mode, lookahead_bytes, lookahead_bytes1, next_parse_state, off, padding, padding1, result, size, size1, skipped_error, start_position, symbol, symbol1, symbol2, token_is_extra, v3
	lex_mode = ts_language_lex_mode_for_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, parse_state)
	if libc.Int32FromUint16(lex_mode.Flex_state) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1227, 0)
			ts_parser__log(tls, self)
		}
		return Subtree{}
	}
	start_position = ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	external_token = ts_stack_last_external_token(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	found_external_token = libc.BoolUint8(0 != 0)
	error_mode = libc.BoolUint8(libc.Int32FromUint16(parse_state) == 0)
	skipped_error = libc.BoolUint8(0 != 0)
	called_get_column = libc.BoolUint8(0 != 0)
	first_error_character = 0
	error_start_position = length_zero(tls)
	error_end_position = length_zero(tls)
	*(*uint32_t)(unsafe.Pointer(bp)) = uint32(0)
	external_scanner_state_len = uint32(0)
	external_scanner_state_changed = libc.BoolUint8(0 != 0)
	ts_lexer_reset(tls, self, start_position)
	for {
		found_token = libc.BoolUint8(0 != 0)
		current_position = (*TSParser)(unsafe.Pointer(self)).Flexer.Fcurrent_position
		column_data = (*TSParser)(unsafe.Pointer(self)).Flexer.Fcolumn_data
		if libc.Int32FromUint16(lex_mode.Fexternal_lex_state) != 0 {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1265, libc.VaList(bp+24, libc.Int32FromUint16(lex_mode.Fexternal_lex_state), current_position.Fextent.Frow, current_position.Fextent.Fcolumn))
				ts_parser__log(tls, self)
			}
			ts_lexer_start(tls, self)
			ts_parser__external_scanner_deserialize(tls, self, external_token)
			found_token = ts_parser__external_scanner_scan(tls, self, lex_mode.Fexternal_lex_state)
			if (*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error != 0 {
				return Subtree{}
			}
			ts_lexer_finish(tls, self, bp)
			if found_token != 0 {
				external_scanner_state_len = ts_parser__external_scanner_serialize(tls, self)
				external_scanner_state_changed = libc.BoolUint8(!(ts_external_scanner_state_eq(tls, ts_subtree_external_scanner_state(tls, external_token), self+192, external_scanner_state_len) != 0))
				if (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_end_position.Fbytes <= current_position.Fbytes && !(external_scanner_state_changed != 0) {
					symbol = *(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fsymbol_map + uintptr((*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fresult_symbol)*2))
					next_parse_state = ts_language_next_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, parse_state, symbol)
					token_is_extra = libc.BoolUint8(libc.Int32FromUint16(next_parse_state) == libc.Int32FromUint16(parse_state))
					if error_mode != 0 || !(ts_stack_has_advanced_since_error(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version) != 0) || token_is_extra != 0 {
						if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
							libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1306, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fsymbol_map + uintptr((*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fresult_symbol)*2)))))
							ts_parser__log(tls, self)
						}
						found_token = libc.BoolUint8(0 != 0)
					}
				}
			}
			if found_token != 0 {
				found_external_token = libc.BoolUint8(1 != 0)
				called_get_column = (*TSParser)(unsafe.Pointer(self)).Flexer.Fdid_get_column
				break
			}
			ts_lexer_reset(tls, self, current_position)
			(*TSParser)(unsafe.Pointer(self)).Flexer.Fcolumn_data = column_data
		}
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1344, libc.VaList(bp+24, libc.Int32FromUint16(lex_mode.Flex_state), current_position.Fextent.Frow, current_position.Fextent.Fcolumn))
			ts_parser__log(tls, self)
		}
		ts_lexer_start(tls, self)
		found_token = ts_parser__call_main_lex_fn(tls, self, lex_mode)
		ts_lexer_finish(tls, self, bp)
		if found_token != 0 {
			break
		}
		if !(error_mode != 0) {
			error_mode = libc.BoolUint8(1 != 0)
			lex_mode = ts_language_lex_mode_for_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, uint16(0))
			ts_lexer_reset(tls, self, start_position)
			goto _1
		}
		if !(skipped_error != 0) {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1385, 0)
				ts_parser__log(tls, self)
			}
			skipped_error = libc.BoolUint8(1 != 0)
			error_start_position = (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_start_position
			error_end_position = (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_start_position
			first_error_character = (*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Flookahead
		}
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Fcurrent_position.Fbytes == error_end_position.Fbytes {
			if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Feof})))(tls, self) != 0 {
				(*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fresult_symbol = libc.Uint16FromInt32(-libc.Int32FromInt32(1))
				break
			}
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fadvance})))(tls, self, libc.BoolUint8(0 != 0))
		}
		error_end_position = (*TSParser)(unsafe.Pointer(self)).Flexer.Fcurrent_position
		goto _1
	_1:
	}
	if skipped_error != 0 {
		padding = length_sub(tls, error_start_position, start_position)
		size = length_sub(tls, error_end_position, error_start_position)
		lookahead_bytes = *(*uint32_t)(unsafe.Pointer(bp)) - error_end_position.Fbytes
		result = ts_subtree_new_error(tls, self+1224, first_error_character, padding, size, lookahead_bytes, parse_state, (*TSParser)(unsafe.Pointer(self)).Flanguage)
	} else {
		is_keyword = libc.BoolUint8(0 != 0)
		symbol1 = (*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fresult_symbol
		padding1 = length_sub(tls, (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_start_position, start_position)
		size1 = length_sub(tls, (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_end_position, (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_start_position)
		lookahead_bytes1 = *(*uint32_t)(unsafe.Pointer(bp)) - (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_end_position.Fbytes
		if found_external_token != 0 {
			symbol1 = *(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fexternal_scanner.Fsymbol_map + uintptr(symbol1)*2))
		} else {
			if libc.Int32FromUint16(symbol1) == libc.Int32FromUint16((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fkeyword_capture_token) && libc.Int32FromUint16(symbol1) != 0 {
				end_byte = (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_end_position.Fbytes
				ts_lexer_reset(tls, self, (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_start_position)
				ts_lexer_start(tls, self)
				is_keyword = ts_parser__call_keyword_lex_fn(tls, self)
				if is_keyword != 0 && (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_end_position.Fbytes == end_byte && (ts_language_has_actions(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, parse_state, (*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fresult_symbol) != 0 || ts_language_is_reserved_word(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, parse_state, (*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fresult_symbol) != 0) {
					symbol1 = (*TSParser)(unsafe.Pointer(self)).Flexer.Fdata.Fresult_symbol
				}
			}
		}
		result = ts_subtree_new_leaf(tls, self+1224, symbol1, padding1, size1, lookahead_bytes1, parse_state, found_external_token, called_get_column, is_keyword, (*TSParser)(unsafe.Pointer(self)).Flanguage)
		if found_external_token != 0 {
			*(*MutableSubtree)(unsafe.Pointer(bp + 8)) = MutableSubtree{}
			*(*struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			})(unsafe.Pointer(bp + 8)) = ts_subtree_to_mut_unsafe(tls, result)
			ts_external_scanner_state_init(tls, *(*uintptr)(unsafe.Pointer(bp + 8))+48, self+192, external_scanner_state_len)
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp + 8))+44, external_scanner_state_changed, 7, 0x80)
		}
	}
	if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
		buf = self + 192
		symbol2 = ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, result))
		off = libc.X__builtin_snprintf(tls, buf, uint64(1024), __ccgo_ts+1413, 0)
		i = 0
		for {
			if !(int32(*(*int8)(unsafe.Pointer(symbol2 + uintptr(i)))) != int32('\000') && off < int32(1024)) {
				break
			}
			switch int32(*(*int8)(unsafe.Pointer(symbol2 + uintptr(i)))) {
			case int32('\t'):
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('\\')
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('t')
			case int32('\n'):
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('\\')
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('n')
			case int32('\v'):
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('\\')
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('v')
			case int32('\f'):
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('\\')
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('f')
			case int32('\r'):
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('\\')
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('r')
			case int32('\\'):
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('\\')
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = int8('\\')
			default:
				v3 = off
				off = off + 1
				*(*int8)(unsafe.Pointer(buf + uintptr(v3))) = *(*int8)(unsafe.Pointer(symbol2 + uintptr(i)))
				break
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		libc.X__builtin_snprintf(tls, buf+uintptr(off), libc.Uint64FromInt32(int32(1024)-off), __ccgo_ts+1434, libc.VaList(bp+24, ts_subtree_total_size(tls, result).Fbytes))
		ts_parser__log(tls, self)
	}
	return result
}

func ts_parser__get_cached_token(tls *libc.TLS, self uintptr, state TSStateId, position size_t, last_external_token Subtree, table_entry uintptr) (r Subtree) {
	var cache uintptr
	_ = cache
	cache = self + 1344
	if *(*uintptr)(unsafe.Pointer(cache)) != 0 && uint64((*TokenCache)(unsafe.Pointer(cache)).Fbyte_index) == position && ts_subtree_external_scanner_state_eq(tls, (*TokenCache)(unsafe.Pointer(cache)).Flast_external_token, last_external_token) != 0 {
		ts_language_table_entry(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, state, ts_subtree_symbol(tls, (*TokenCache)(unsafe.Pointer(cache)).Ftoken), table_entry)
		if ts_parser__can_reuse_first_leaf(tls, self, state, (*TokenCache)(unsafe.Pointer(cache)).Ftoken, table_entry) != 0 {
			ts_subtree_retain(tls, (*TokenCache)(unsafe.Pointer(cache)).Ftoken)
			return (*TokenCache)(unsafe.Pointer(cache)).Ftoken
		}
	}
	return Subtree{}
}

func ts_parser__set_cached_token(tls *libc.TLS, self uintptr, byte_index uint32_t, _last_external_token Subtree, _token Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _last_external_token
	*(*Subtree)(unsafe.Pointer(bp + 8)) = _token
	var cache uintptr
	_ = cache
	cache = self + 1344
	if *(*uintptr)(unsafe.Pointer(bp + 8)) != 0 {
		ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
	}
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer(bp)))
	}
	if *(*uintptr)(unsafe.Pointer(cache)) != 0 {
		ts_subtree_release(tls, self+1224, (*TokenCache)(unsafe.Pointer(cache)).Ftoken)
	}
	if *(*uintptr)(unsafe.Pointer(cache + 8)) != 0 {
		ts_subtree_release(tls, self+1224, (*TokenCache)(unsafe.Pointer(cache)).Flast_external_token)
	}
	(*TokenCache)(unsafe.Pointer(cache)).Ftoken = *(*Subtree)(unsafe.Pointer(bp + 8))
	(*TokenCache)(unsafe.Pointer(cache)).Fbyte_index = byte_index
	(*TokenCache)(unsafe.Pointer(cache)).Flast_external_token = *(*Subtree)(unsafe.Pointer(bp))
}

func ts_parser__has_included_range_difference(tls *libc.TLS, self uintptr, start_position uint32_t, end_position uint32_t) (r uint8) {
	return ts_range_array_intersects(tls, self+1424, (*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index, start_position, end_position)
}

func ts_parser__reuse_node(tls *libc.TLS, self uintptr, version StackVersion, state uintptr, position uint32_t, last_external_token Subtree, table_entry uintptr) (r Subtree) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var byte_offset, end_byte_offset uint32_t
	var leaf_symbol TSSymbol
	var reason uintptr
	var _ /* result at bp+0 */ Subtree
	_, _, _, _ = byte_offset, end_byte_offset, leaf_symbol, reason
	for *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		byte_offset = reusable_node_byte_offset(tls, self+1368)
		end_byte_offset = byte_offset + ts_subtree_total_bytes(tls, *(*Subtree)(unsafe.Pointer(bp)))
		if ts_subtree_is_eof(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
			end_byte_offset = libc.Uint32FromUint32(4294967295)
		}
		if byte_offset > position {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1444, libc.VaList(bp+16, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))))
				ts_parser__log(tls, self)
			}
			break
		}
		if byte_offset < position {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1475, libc.VaList(bp+16, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))))
				ts_parser__log(tls, self)
			}
			if end_byte_offset <= position || !(reusable_node_descend(tls, self+1368) != 0) {
				reusable_node_advance(tls, self+1368)
			}
			continue
		}
		if !(ts_subtree_external_scanner_state_eq(tls, (*TSParser)(unsafe.Pointer(self)).Freusable_node.Flast_external_token, last_external_token) != 0) {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1504, libc.VaList(bp+16, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))))
				ts_parser__log(tls, self)
			}
			reusable_node_advance(tls, self+1368)
			continue
		}
		reason = libc.UintptrFromInt32(0)
		if ts_subtree_has_changes(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
			reason = __ccgo_ts + 1565
		} else {
			if ts_subtree_is_error(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
				reason = __ccgo_ts + 1577
			} else {
				if ts_subtree_missing(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
					reason = __ccgo_ts + 1586
				} else {
					if ts_subtree_is_fragile(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
						reason = __ccgo_ts + 1597
					} else {
						if ts_parser__has_included_range_difference(tls, self, byte_offset, end_byte_offset) != 0 {
							reason = __ccgo_ts + 1608
						}
					}
				}
			}
		}
		if reason != 0 {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1642, libc.VaList(bp+16, reason, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))))
				ts_parser__log(tls, self)
			}
			if !(reusable_node_descend(tls, self+1368) != 0) {
				reusable_node_advance(tls, self+1368)
				ts_parser__breakdown_top_of_stack(tls, self, version)
				*(*TSStateId)(unsafe.Pointer(state)) = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
			}
			continue
		}
		leaf_symbol = ts_subtree_leaf_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))
		ts_language_table_entry(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSStateId)(unsafe.Pointer(state)), leaf_symbol, table_entry)
		if !(ts_parser__can_reuse_first_leaf(tls, self, *(*TSStateId)(unsafe.Pointer(state)), *(*Subtree)(unsafe.Pointer(bp)), table_entry) != 0) {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1669, libc.VaList(bp+16, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, leaf_symbol)))
				ts_parser__log(tls, self)
			}
			reusable_node_advance_past_leaf(tls, self+1368)
			break
		}
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1717, libc.VaList(bp+16, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))))
			ts_parser__log(tls, self)
		}
		ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer(bp)))
		return *(*Subtree)(unsafe.Pointer(bp))
	}
	return Subtree{}
}

func ts_parser__select_tree(tls *libc.TLS, self uintptr, _left Subtree, _right Subtree) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	*(*Subtree)(unsafe.Pointer(bp)) = _left
	*(*Subtree)(unsafe.Pointer(bp + 8)) = _right
	var comparison int32
	_ = comparison
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) {
		return libc.BoolUint8(1 != 0)
	}
	if !(*(*uintptr)(unsafe.Pointer(bp + 8)) != 0) {
		return libc.BoolUint8(0 != 0)
	}
	if ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) < ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp))) {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1738, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))))
			ts_parser__log(tls, self)
		}
		return libc.BoolUint8(1 != 0)
	}
	if ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp))) < ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1738, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8))))))
			ts_parser__log(tls, self)
		}
		return libc.BoolUint8(0 != 0)
	}
	if ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) > ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp))) {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1785, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))), ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp + 8))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))), ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp)))))
			ts_parser__log(tls, self)
		}
		return libc.BoolUint8(1 != 0)
	}
	if ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp))) > ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1785, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))), ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))), ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))))
			ts_parser__log(tls, self)
		}
		return libc.BoolUint8(0 != 0)
	}
	if ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
		return libc.BoolUint8(1 != 0)
	}
	comparison = ts_subtree_compare(tls, *(*Subtree)(unsafe.Pointer(bp)), *(*Subtree)(unsafe.Pointer(bp + 8)), self+1224)
	switch comparison {
	case -int32(1):
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1860, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8))))))
			ts_parser__log(tls, self)
		}
		return libc.BoolUint8(0 != 0)
	case int32(1):
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1860, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))))
			ts_parser__log(tls, self)
		}
		return libc.BoolUint8(1 != 0)
	default:
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1901, libc.VaList(bp+24, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8))))))
			ts_parser__log(tls, self)
		}
		return libc.BoolUint8(0 != 0)
	}
	return r
}

func ts_parser__select_children(tls *libc.TLS, self uintptr, left Subtree, children uintptr) (r uint8) {
	var scratch_tree MutableSubtree
	_ = scratch_tree
	(*SubtreeArray)(unsafe.Pointer(self + 1328)).Fcontents = _array__assign(tls, (*SubtreeArray)(unsafe.Pointer(self+1328)).Fcontents, self+1328+8, self+1328+12, (*SubtreeArray)(unsafe.Pointer(children)).Fcontents, (*SubtreeArray)(unsafe.Pointer(children)).Fsize, libc.Uint64FromInt64(8))
	scratch_tree = ts_subtree_new_node(tls, ts_subtree_symbol(tls, left), self+1328, uint32(0), (*TSParser)(unsafe.Pointer(self)).Flanguage)
	return ts_parser__select_tree(tls, self, left, ts_subtree_from_mut(tls, scratch_tree))
}

func ts_parser__shift(tls *libc.TLS, self uintptr, version StackVersion, state TSStateId, lookahead Subtree, extra uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var is_leaf uint8
	var subtree_to_push Subtree
	var _ /* result at bp+0 */ MutableSubtree
	_, _ = is_leaf, subtree_to_push
	is_leaf = libc.BoolUint8(ts_subtree_child_count(tls, lookahead) == uint32(0))
	subtree_to_push = lookahead
	if libc.Int32FromUint8(extra) != libc.Int32FromUint8(ts_subtree_extra(tls, lookahead)) && is_leaf != 0 {
		*(*MutableSubtree)(unsafe.Pointer(bp)) = MutableSubtree{}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp)) = ts_subtree_make_mut(tls, self+1224, lookahead)
		ts_subtree_set_extra(tls, bp, extra)
		subtree_to_push = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp)))
	}
	ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, subtree_to_push, libc.BoolUint8(!(is_leaf != 0)), state)
	if ts_subtree_has_external_tokens(tls, subtree_to_push) != 0 {
		ts_stack_set_last_external_token(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, ts_subtree_last_external_token(tls, subtree_to_push))
	}
}

func ts_parser__reduce(tls *libc.TLS, self uintptr, version StackVersion, symbol TSSymbol, count uint32_t, dynamic_precedence int32, production_id uint16_t, is_fragile uint8, end_of_non_terminal_extra uint8) (r StackVersion) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var _array_swap_tmp uintptr
	var halted_version_count, i, initial_version_count, j, removed_version_count uint32_t
	var j1, slice_version StackVersion
	var next_state, state TSStateId
	var v4 uint32
	var _ /* children at bp+64 */ SubtreeArray
	var _ /* next_slice at bp+40 */ StackSlice
	var _ /* next_slice at bp+88 */ StackSlice
	var _ /* next_slice_children at bp+112 */ SubtreeArray
	var _ /* parent at bp+80 */ MutableSubtree
	var _ /* pop at bp+0 */ StackSliceArray
	var _ /* slice at bp+16 */ StackSlice
	_, _, _, _, _, _, _, _, _, _, _ = _array_swap_tmp, halted_version_count, i, initial_version_count, j, j1, next_state, removed_version_count, slice_version, state, v4
	initial_version_count = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	*(*StackSliceArray)(unsafe.Pointer(bp)) = ts_stack_pop_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, count)
	removed_version_count = uint32(0)
	halted_version_count = ts_stack_halted_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	i = uint32(0)
	for {
		if !(i < (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+1004, __ccgo_ts+1033, int32(952), uintptr(unsafe.Pointer(&__func__18)))
			}
		}
		*(*StackSlice)(unsafe.Pointer(bp + 16)) = *(*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents + uintptr(i)*24))
		slice_version = (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion - removed_version_count
		if slice_version > MAX_VERSION_COUNT+MAX_VERSION_COUNT_OVERFLOW+halted_version_count {
			ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, slice_version)
			ts_subtree_array_delete(tls, self+1224, bp+16)
			removed_version_count = removed_version_count + 1
			for i+uint32(1) < (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize {
				if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
					libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+1943, 0)
					ts_parser__log(tls, self)
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(i+libc.Uint32FromInt32(1) < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+1982, __ccgo_ts+1033, int32(965), uintptr(unsafe.Pointer(&__func__18)))
					}
				}
				*(*StackSlice)(unsafe.Pointer(bp + 40)) = *(*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents + uintptr(i+uint32(1))*24))
				if (*(*StackSlice)(unsafe.Pointer(bp + 40))).Fversion != (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion {
					break
				}
				ts_subtree_array_delete(tls, self+1224, bp+40)
				i = i + 1
			}
			goto _1
		}
		*(*SubtreeArray)(unsafe.Pointer(bp + 64)) = (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fsubtrees
		ts_subtree_array_remove_trailing_extras(tls, bp+64, self+1296)
		*(*MutableSubtree)(unsafe.Pointer(bp + 80)) = MutableSubtree{}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 80)) = ts_subtree_new_node(tls, symbol, bp+64, uint32(production_id), (*TSParser)(unsafe.Pointer(self)).Flanguage)
		for i+uint32(1) < (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize {
			_ = libc.Uint64FromInt64(4)
			{
				if !(i+libc.Uint32FromInt32(1) < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+1982, __ccgo_ts+1033, int32(988), uintptr(unsafe.Pointer(&__func__18)))
				}
			}
			*(*StackSlice)(unsafe.Pointer(bp + 88)) = *(*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents + uintptr(i+uint32(1))*24))
			if (*(*StackSlice)(unsafe.Pointer(bp + 88))).Fversion != (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion {
				break
			}
			i = i + 1
			*(*SubtreeArray)(unsafe.Pointer(bp + 112)) = (*(*StackSlice)(unsafe.Pointer(bp + 88))).Fsubtrees
			ts_subtree_array_remove_trailing_extras(tls, bp+112, self+1312)
			if ts_parser__select_children(tls, self, ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 80))), bp+112) != 0 {
				ts_subtree_array_clear(tls, self+1224, self+1296)
				ts_subtree_release(tls, self+1224, ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 80))))
				_array_swap_tmp = (*SubtreeArray)(unsafe.Pointer(self + 1296)).Fcontents
				(*SubtreeArray)(unsafe.Pointer(self + 1296)).Fcontents = (*SubtreeArray)(unsafe.Pointer(self + 1312)).Fcontents
				(*SubtreeArray)(unsafe.Pointer(self + 1312)).Fcontents = _array_swap_tmp
				_array__swap(tls, self+1296+8, self+1296+12, self+1312+8, self+1312+12)
				*(*MutableSubtree)(unsafe.Pointer(bp + 80)) = ts_subtree_new_node(tls, symbol, bp+112, uint32(production_id), (*TSParser)(unsafe.Pointer(self)).Flanguage)
			} else {
				(*SubtreeArray)(unsafe.Pointer(self + 1312)).Fsize = uint32(0)
				ts_subtree_array_delete(tls, self+1224, bp+88)
			}
		}
		state = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, slice_version)
		next_state = ts_language_next_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, state, symbol)
		if end_of_non_terminal_extra != 0 && libc.Int32FromUint16(next_state) == libc.Int32FromUint16(state) {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+44, libc.BoolUint8(1 != 0), 2, 0x4)
		}
		if is_fragile != 0 || (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize > uint32(1) || initial_version_count > uint32(1) {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+44, libc.BoolUint8(1 != 0), 3, 0x8)
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+44, libc.BoolUint8(1 != 0), 4, 0x10)
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fparse_state = libc.Uint16FromInt32(libc.Int32FromInt32(0x7fff)*libc.Int32FromInt32(2) + libc.Int32FromInt32(1))
		} else {
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fparse_state = state
		}
		*(*int32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)) + 60)) += dynamic_precedence
		ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, slice_version, ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 80))), libc.BoolUint8(0 != 0), next_state)
		j = uint32(0)
		for {
			if !(j < (*TSParser)(unsafe.Pointer(self)).Ftrailing_extras.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j < (*SubtreeArray)(unsafe.Pointer(self+1296)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+2015, __ccgo_ts+1033, int32(1030), uintptr(unsafe.Pointer(&__func__18)))
				}
			}
			ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, slice_version, *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self+1296)).Fcontents + uintptr(j)*8)), libc.BoolUint8(0 != 0), next_state)
			goto _2
		_2:
			;
			j = j + 1
		}
		j1 = uint32(0)
		for {
			if !(j1 < slice_version) {
				break
			}
			if j1 == version {
				goto _3
			}
			if ts_stack_merge(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, j1, slice_version) != 0 {
				removed_version_count = removed_version_count + 1
				break
			}
			goto _3
		_3:
			;
			j1 = j1 + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack) > initial_version_count {
		v4 = initial_version_count
	} else {
		v4 = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
	}
	return v4
}

var __func__18 = [18]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'r', 'e', 'd', 'u', 'c', 'e'}

func ts_parser__accept(tls *libc.TLS, self uintptr, version StackVersion, lookahead Subtree) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var child_count, i, j, k uint32_t
	var children, v3 uintptr
	var _ /* pop at bp+0 */ StackSliceArray
	var _ /* root at bp+32 */ Subtree
	var _ /* tree at bp+40 */ Subtree
	var _ /* trees at bp+16 */ SubtreeArray
	_, _, _, _, _, _ = child_count, children, i, j, k, v3
	_ = libc.Uint64FromInt64(4)
	{
		if !(ts_subtree_is_eof(tls, lookahead) != 0) {
			libc.X__assert_fail(tls, __ccgo_ts+2062, __ccgo_ts+1033, int32(1053), uintptr(unsafe.Pointer(&__func__19)))
		}
	}
	ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, lookahead, libc.BoolUint8(0 != 0), uint16(1))
	*(*StackSliceArray)(unsafe.Pointer(bp)) = ts_stack_pop_all(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	i = uint32(0)
	for {
		if !(i < (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+1004, __ccgo_ts+1033, int32(1058), uintptr(unsafe.Pointer(&__func__19)))
			}
		}
		*(*SubtreeArray)(unsafe.Pointer(bp + 16)) = (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents + uintptr(i)*24)).Fsubtrees
		*(*Subtree)(unsafe.Pointer(bp + 32)) = Subtree{}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 32)) = Subtree{}
		j = (*(*SubtreeArray)(unsafe.Pointer(bp + 16))).Fsize - uint32(1)
		for {
			if !(j+uint32(1) > uint32(0)) {
				break
			}
			*(*Subtree)(unsafe.Pointer(bp + 40)) = Subtree{}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j < (*SubtreeArray)(unsafe.Pointer(bp+16)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+2091, __ccgo_ts+1033, int32(1062), uintptr(unsafe.Pointer(&__func__19)))
				}
			}
			*(*struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			})(unsafe.Pointer(bp + 40)) = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents + uintptr(j)*8))
			if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp + 40))) != 0) {
				_ = libc.Uint64FromInt64(4)
				{
					if !!(int32(*(*uint8)(unsafe.Pointer(bp + 40 + 0))&0x1>>0) != 0) {
						libc.X__assert_fail(tls, __ccgo_ts+2122, __ccgo_ts+1033, int32(1064), uintptr(unsafe.Pointer(&__func__19)))
					}
				}
				child_count = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 40)))
				if int32(*(*uint8)(unsafe.Pointer(bp + 40 + 0))&0x1>>0) != 0 {
					v3 = libc.UintptrFromInt32(0)
				} else {
					v3 = *(*uintptr)(unsafe.Pointer(bp + 40)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)))).Fchild_count)*8
				}
				children = v3
				k = uint32(0)
				for {
					if !(k < child_count) {
						break
					}
					ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer(children + uintptr(k)*8)))
					goto _4
				_4:
					;
					k = k + 1
				}
				(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fcontents = _array__splice(tls, (*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents, bp+16+8, bp+16+12, libc.Uint64FromInt64(8), j, uint32(1), child_count, children)
				*(*Subtree)(unsafe.Pointer(bp + 32)) = ts_subtree_from_mut(tls, ts_subtree_new_node(tls, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 40))), bp+16, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), (*TSParser)(unsafe.Pointer(self)).Flanguage))
				ts_subtree_release(tls, self+1224, *(*Subtree)(unsafe.Pointer(bp + 40)))
				break
			}
			goto _2
		_2:
			;
			j = j - 1
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(*(*uintptr)(unsafe.Pointer(bp + 32)) != 0) {
				libc.X__assert_fail(tls, __ccgo_ts+2143, __ccgo_ts+1033, int32(1082), uintptr(unsafe.Pointer(&__func__19)))
			}
		}
		(*TSParser)(unsafe.Pointer(self)).Faccept_count = (*TSParser)(unsafe.Pointer(self)).Faccept_count + 1
		if *(*uintptr)(unsafe.Pointer(self + 1288)) != 0 {
			if ts_parser__select_tree(tls, self, (*TSParser)(unsafe.Pointer(self)).Ffinished_tree, *(*Subtree)(unsafe.Pointer(bp + 32))) != 0 {
				ts_subtree_release(tls, self+1224, (*TSParser)(unsafe.Pointer(self)).Ffinished_tree)
				(*TSParser)(unsafe.Pointer(self)).Ffinished_tree = *(*Subtree)(unsafe.Pointer(bp + 32))
			} else {
				ts_subtree_release(tls, self+1224, *(*Subtree)(unsafe.Pointer(bp + 32)))
			}
		} else {
			(*TSParser)(unsafe.Pointer(self)).Ffinished_tree = *(*Subtree)(unsafe.Pointer(bp + 32))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1097), uintptr(unsafe.Pointer(&__func__19)))
		}
	}
	ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents)).Fversion)
	ts_stack_halt(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
}

var __func__19 = [18]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'a', 'c', 'c', 'e', 'p', 't'}

func ts_parser__do_all_potential_reductions(tls *libc.TLS, self uintptr, starting_version StackVersion, lookahead_symbol TSSymbol) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var action1 ReduceAction
	var can_shift_lookahead_symbol, has_shift_action, merged uint8
	var end_symbol, first_symbol, symbol TSSymbol
	var i uint32
	var initial_version_count, j1, j2, version_count uint32_t
	var j, reduction_version, version StackVersion
	var state TSStateId
	var _ /* action at bp+16 */ TSParseAction
	var _ /* entry at bp+0 */ TableEntry
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = action1, can_shift_lookahead_symbol, end_symbol, first_symbol, has_shift_action, i, initial_version_count, j, j1, j2, merged, reduction_version, state, symbol, version, version_count
	initial_version_count = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	can_shift_lookahead_symbol = libc.BoolUint8(0 != 0)
	version = starting_version
	i = uint32(0)
	for {
		if !(int32(1) != 0) {
			break
		}
		version_count = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
		if version >= version_count {
			break
		}
		merged = libc.BoolUint8(0 != 0)
		j = initial_version_count
		for {
			if !(j < version) {
				break
			}
			if ts_stack_merge(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, j, version) != 0 {
				merged = libc.BoolUint8(1 != 0)
				break
			}
			goto _2
		_2:
			;
			j = j + 1
		}
		if merged != 0 {
			goto _1
		}
		state = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
		has_shift_action = libc.BoolUint8(0 != 0)
		(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fsize = uint32(0)
		if libc.Int32FromUint16(lookahead_symbol) != 0 {
			first_symbol = lookahead_symbol
			end_symbol = libc.Uint16FromInt32(libc.Int32FromUint16(lookahead_symbol) + int32(1))
		} else {
			first_symbol = uint16(1)
			end_symbol = uint16((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Ftoken_count)
		}
		symbol = first_symbol
		for {
			if !(libc.Int32FromUint16(symbol) < libc.Int32FromUint16(end_symbol)) {
				break
			}
			ts_language_table_entry(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, state, symbol, bp)
			j1 = uint32(0)
			for {
				if !(j1 < (*(*TableEntry)(unsafe.Pointer(bp))).Faction_count) {
					break
				}
				*(*TSParseAction)(unsafe.Pointer(bp + 16)) = TSParseAction{}
				*(*struct {
					Freduce [0]struct {
						Ftype1              uint8
						Fchild_count        uint8
						Fsymbol             uint16
						Fdynamic_precedence int16
						Fproduction_id      uint16
					}
					Ftype1 [0]uint8
					Fshift struct {
						Ftype1      uint8
						Fstate      uint16
						Fextra      uint8
						Frepetition uint8
					}
					F__ccgo_pad3 [2]byte
				})(unsafe.Pointer(bp + 16)) = *(*TSParseAction)(unsafe.Pointer((*(*TableEntry)(unsafe.Pointer(bp))).Factions + uintptr(j1)*8))
				switch libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(bp + 16))))) {
				case int32(TSParseActionTypeShift):
					fallthrough
				case int32(TSParseActionTypeRecover):
					if !((*(*TSParseAction)(unsafe.Pointer(bp + 16))).Fshift.Fextra != 0) && !((*(*TSParseAction)(unsafe.Pointer(bp + 16))).Fshift.Frepetition != 0) {
						has_shift_action = libc.BoolUint8(1 != 0)
					}
				case int32(TSParseActionTypeReduce):
					if libc.Int32FromUint8((*(*struct {
						Ftype1              uint8_t
						Fchild_count        uint8_t
						Fsymbol             TSSymbol
						Fdynamic_precedence int16_t
						Fproduction_id      uint16_t
					})(unsafe.Pointer(bp + 16))).Fchild_count) > 0 {
						ts_reduce_action_set_add(tls, self+1272, ReduceAction{
							Fcount: uint32((*(*struct {
								Ftype1              uint8_t
								Fchild_count        uint8_t
								Fsymbol             TSSymbol
								Fdynamic_precedence int16_t
								Fproduction_id      uint16_t
							})(unsafe.Pointer(bp + 16))).Fchild_count),
							Fsymbol: (*(*struct {
								Ftype1              uint8_t
								Fchild_count        uint8_t
								Fsymbol             TSSymbol
								Fdynamic_precedence int16_t
								Fproduction_id      uint16_t
							})(unsafe.Pointer(bp + 16))).Fsymbol,
							Fdynamic_precedence: int32((*(*struct {
								Ftype1              uint8_t
								Fchild_count        uint8_t
								Fsymbol             TSSymbol
								Fdynamic_precedence int16_t
								Fproduction_id      uint16_t
							})(unsafe.Pointer(bp + 16))).Fdynamic_precedence),
							Fproduction_id: (*(*struct {
								Ftype1              uint8_t
								Fchild_count        uint8_t
								Fsymbol             TSSymbol
								Fdynamic_precedence int16_t
								Fproduction_id      uint16_t
							})(unsafe.Pointer(bp + 16))).Fproduction_id,
						})
					}
				default:
					break
				}
				goto _4
			_4:
				;
				j1 = j1 + 1
			}
			goto _3
		_3:
			;
			symbol = symbol + 1
		}
		reduction_version = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
		j2 = uint32(0)
		for {
			if !(j2 < (*TSParser)(unsafe.Pointer(self)).Freduce_actions.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j2 < (*ReduceActionSet)(unsafe.Pointer(self+1272)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+2181, __ccgo_ts+1033, int32(1163), uintptr(unsafe.Pointer(&__func__20)))
				}
			}
			action1 = *(*ReduceAction)(unsafe.Pointer((*ReduceActionSet)(unsafe.Pointer(self+1272)).Fcontents + uintptr(j2)*16))
			reduction_version = ts_parser__reduce(tls, self, version, action1.Fsymbol, action1.Fcount, action1.Fdynamic_precedence, action1.Fproduction_id, libc.BoolUint8(1 != 0), libc.BoolUint8(0 != 0))
			goto _5
		_5:
			;
			j2 = j2 + 1
		}
		if has_shift_action != 0 {
			can_shift_lookahead_symbol = libc.BoolUint8(1 != 0)
		} else {
			if reduction_version != libc.Uint32FromInt32(-libc.Int32FromInt32(1)) && i < MAX_VERSION_COUNT {
				ts_stack_renumber_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, reduction_version, version)
				goto _1
			} else {
				if libc.Int32FromUint16(lookahead_symbol) != 0 {
					ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
				}
			}
		}
		if version == starting_version {
			version = version_count
		} else {
			version = version + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return can_shift_lookahead_symbol
}

var __func__20 = [39]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'd', 'o', '_', 'a', 'l', 'l', '_', 'p', 'o', 't', 'e', 'n', 't', 'i', 'a', 'l', '_', 'r', 'e', 'd', 'u', 'c', 't', 'i', 'o', 'n', 's'}

func ts_parser__recover_to_state(tls *libc.TLS, self uintptr, version StackVersion, depth uint32, goal_state TSStateId) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var error1, tree Subtree
	var error_child_count uint32_t
	var i, j, j1, v2 uint32
	var previous_version StackVersion
	var v4 uintptr
	var _ /* error_tree at bp+56 */ Subtree
	var _ /* error_trees at bp+40 */ SubtreeArray
	var _ /* pop at bp+0 */ StackSliceArray
	var _ /* slice at bp+16 */ StackSlice
	_, _, _, _, _, _, _, _, _ = error1, error_child_count, i, j, j1, previous_version, tree, v2, v4
	*(*StackSliceArray)(unsafe.Pointer(bp)) = ts_stack_pop_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, depth)
	previous_version = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
	i = uint32(0)
	for {
		if !(i < (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+1004, __ccgo_ts+1033, int32(1201), uintptr(unsafe.Pointer(&__func__21)))
			}
		}
		*(*StackSlice)(unsafe.Pointer(bp + 16)) = *(*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents + uintptr(i)*24))
		if (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion == previous_version {
			ts_subtree_array_delete(tls, self+1224, bp+16)
			v2 = i
			i = i - 1
			_array__erase(tls, (*StackSliceArray)(unsafe.Pointer(bp)).Fcontents, bp+8, libc.Uint64FromInt64(24), v2)
			goto _1
		}
		if libc.Int32FromUint16(ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion)) != libc.Int32FromUint16(goal_state) {
			ts_stack_halt(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion)
			ts_subtree_array_delete(tls, self+1224, bp+16)
			v2 = i
			i = i - 1
			_array__erase(tls, (*StackSliceArray)(unsafe.Pointer(bp)).Fcontents, bp+8, libc.Uint64FromInt64(24), v2)
			goto _1
		}
		*(*SubtreeArray)(unsafe.Pointer(bp + 40)) = ts_stack_pop_error(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion)
		if (*(*SubtreeArray)(unsafe.Pointer(bp + 40))).Fsize > uint32(0) {
			_ = libc.Uint64FromInt64(4)
			{
				if !((*(*SubtreeArray)(unsafe.Pointer(bp + 40))).Fsize == libc.Uint32FromInt32(1)) {
					libc.X__assert_fail(tls, __ccgo_ts+2227, __ccgo_ts+1033, int32(1218), uintptr(unsafe.Pointer(&__func__21)))
				}
			}
			*(*Subtree)(unsafe.Pointer(bp + 56)) = Subtree{}
			_ = libc.Uint64FromInt64(4)
			{
				if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*SubtreeArray)(unsafe.Pointer(bp+40)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+2249, __ccgo_ts+1033, int32(1219), uintptr(unsafe.Pointer(&__func__21)))
				}
			}
			*(*struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			})(unsafe.Pointer(bp + 56)) = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(bp + 40)).Fcontents))
			error_child_count = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 56)))
			if error_child_count > uint32(0) {
				if int32(*(*uint8)(unsafe.Pointer(bp + 56 + 0))&0x1>>0) != 0 {
					v4 = libc.UintptrFromInt32(0)
				} else {
					v4 = *(*uintptr)(unsafe.Pointer(bp + 56)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 56)))).Fchild_count)*8
				}
				(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fcontents = _array__splice(tls, (*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents, bp+16+8, bp+16+12, libc.Uint64FromInt64(8), uint32(0), uint32(0), error_child_count, v4)
				j = uint32(0)
				for {
					if !(j < error_child_count) {
						break
					}
					_ = libc.Uint64FromInt64(4)
					{
						if !(j < (*SubtreeArray)(unsafe.Pointer(bp+16)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+1116, __ccgo_ts+1033, int32(1224), uintptr(unsafe.Pointer(&__func__21)))
						}
					}
					ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents + uintptr(j)*8)))
					goto _5
				_5:
					;
					j = j + 1
				}
			}
			ts_subtree_array_delete(tls, self+1224, bp+40)
		}
		ts_subtree_array_remove_trailing_extras(tls, bp+16, self+1296)
		if (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fsubtrees.Fsize > uint32(0) {
			error1 = ts_subtree_new_error_node(tls, bp+16, libc.BoolUint8(1 != 0), (*TSParser)(unsafe.Pointer(self)).Flanguage)
			ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion, error1, libc.BoolUint8(0 != 0), goal_state)
		} else {
			if (*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*SubtreeArray)(unsafe.Pointer(bp+16)).Fcontents)
			}
			(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fcontents = libc.UintptrFromInt32(0)
			(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fsize = uint32(0)
			(*SubtreeArray)(unsafe.Pointer(bp + 16)).Fcapacity = uint32(0)
		}
		j1 = uint32(0)
		for {
			if !(j1 < (*TSParser)(unsafe.Pointer(self)).Ftrailing_extras.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j1 < (*SubtreeArray)(unsafe.Pointer(self+1296)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+2015, __ccgo_ts+1033, int32(1240), uintptr(unsafe.Pointer(&__func__21)))
				}
			}
			tree = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self+1296)).Fcontents + uintptr(j1)*8))
			ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion, tree, libc.BoolUint8(0 != 0), goal_state)
			goto _6
		_6:
			;
			j1 = j1 + 1
		}
		previous_version = (*(*StackSlice)(unsafe.Pointer(bp + 16))).Fversion
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(previous_version != libc.Uint32FromInt32(-libc.Int32FromInt32(1)))
}

var __func__21 = [28]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'r', 'e', 'c', 'o', 'v', 'e', 'r', '_', 't', 'o', '_', 's', 't', 'a', 't', 'e'}

func ts_parser__recover(tls *libc.TLS, self uintptr, version StackVersion, lookahead Subtree) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var actions, summary, v6 uintptr
	var current_error_cost, depth, i, i1, i2, i3, j, new_cost, new_cost1, node_count_since_error, previous_version_count, v4 uint32
	var did_recover, has_error, would_merge uint8
	var entry StackSummaryEntry
	var error_repeat MutableSubtree
	var parent Subtree
	var position Length
	var status ErrorStatus
	var v5 uint32_t
	var _ /* children at bp+0 */ SubtreeArray
	var _ /* children at bp+32 */ SubtreeArray
	var _ /* mutable_lookahead at bp+24 */ MutableSubtree
	var _ /* n at bp+16 */ uint32
	var _ /* pop at bp+48 */ StackSliceArray
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = actions, current_error_cost, depth, did_recover, entry, error_repeat, has_error, i, i1, i2, i3, j, new_cost, new_cost1, node_count_since_error, parent, position, previous_version_count, status, summary, would_merge, v4, v5, v6
	did_recover = libc.BoolUint8(0 != 0)
	previous_version_count = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	position = ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	summary = ts_stack_get_summary(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	node_count_since_error = ts_stack_node_count_since_error(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	current_error_cost = ts_stack_error_cost(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	if summary != 0 && !(ts_subtree_is_error(tls, lookahead) != 0) {
		i = uint32(0)
		for {
			if !(i < (*StackSummary)(unsafe.Pointer(summary)).Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*StackSummary)(unsafe.Pointer(summary)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+2286, __ccgo_ts+1033, int32(1276), uintptr(unsafe.Pointer(&__func__22)))
				}
			}
			entry = *(*StackSummaryEntry)(unsafe.Pointer((*StackSummary)(unsafe.Pointer(summary)).Fcontents + uintptr(i)*20))
			if libc.Int32FromUint16(entry.Fstate) == 0 {
				goto _1
			}
			if entry.Fposition.Fbytes == position.Fbytes {
				goto _1
			}
			depth = entry.Fdepth
			if node_count_since_error > uint32(0) {
				depth = depth + 1
			}
			would_merge = libc.BoolUint8(0 != 0)
			j = uint32(0)
			for {
				if !(j < previous_version_count) {
					break
				}
				if libc.Int32FromUint16(ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, j)) == libc.Int32FromUint16(entry.Fstate) && ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, j).Fbytes == position.Fbytes {
					would_merge = libc.BoolUint8(1 != 0)
					break
				}
				goto _2
			_2:
				;
				j = j + 1
			}
			if would_merge != 0 {
				goto _1
			}
			new_cost = current_error_cost + entry.Fdepth*uint32(100) + (position.Fbytes-entry.Fposition.Fbytes)*uint32(1) + (position.Fextent.Frow-entry.Fposition.Fextent.Frow)*uint32(30)
			if ts_parser__better_version_exists(tls, self, version, libc.BoolUint8(0 != 0), new_cost) != 0 {
				break
			}
			if ts_language_has_actions(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, entry.Fstate, ts_subtree_symbol(tls, lookahead)) != 0 {
				if ts_parser__recover_to_state(tls, self, version, depth, entry.Fstate) != 0 {
					did_recover = libc.BoolUint8(1 != 0)
					if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
						libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2318, libc.VaList(bp+72, libc.Int32FromUint16(entry.Fstate), depth))
						ts_parser__log(tls, self)
					}
					if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
						ts_stack_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
						libc.Xfputs(tls, __ccgo_ts+1187, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
					}
					break
				}
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	i1 = previous_version_count
	for {
		if !(i1 < ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)) {
			break
		}
		if !(ts_stack_is_active(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i1) != 0) {
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2357, libc.VaList(bp+72, i1))
				ts_parser__log(tls, self)
			}
			v4 = i1
			i1 = i1 - 1
			ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, v4)
			if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				ts_stack_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
				libc.Xfputs(tls, __ccgo_ts+1187, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
			}
		}
		goto _3
	_3:
		;
		i1 = i1 + 1
	}
	if ts_subtree_is_eof(tls, lookahead) != 0 {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2383, 0)
			ts_parser__log(tls, self)
		}
		*(*SubtreeArray)(unsafe.Pointer(bp)) = SubtreeArray{}
		parent = ts_subtree_new_error_node(tls, bp, libc.BoolUint8(0 != 0), (*TSParser)(unsafe.Pointer(self)).Flanguage)
		ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, parent, libc.BoolUint8(0 != 0), uint16(1))
		ts_parser__accept(tls, self, version, lookahead)
		return
	}
	if did_recover != 0 && ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack) > MAX_VERSION_COUNT {
		ts_stack_halt(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
		ts_subtree_release(tls, self+1224, lookahead)
		return
	}
	if did_recover != 0 && ts_subtree_has_external_scanner_state_change(tls, lookahead) != 0 {
		ts_stack_halt(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
		ts_subtree_release(tls, self+1224, lookahead)
		return
	}
	new_cost1 = current_error_cost + uint32(100) + ts_subtree_total_bytes(tls, lookahead)*uint32(1) + ts_subtree_total_size(tls, lookahead).Fextent.Frow*uint32(30)
	if ts_parser__better_version_exists(tls, self, version, libc.BoolUint8(0 != 0), new_cost1) != 0 {
		ts_stack_halt(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
		ts_subtree_release(tls, self+1224, lookahead)
		return
	}
	actions = ts_language_actions(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, uint16(1), ts_subtree_symbol(tls, lookahead), bp+16)
	if *(*uint32)(unsafe.Pointer(bp + 16)) > uint32(0) && libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(actions + uintptr(*(*uint32)(unsafe.Pointer(bp + 16))-uint32(1))*8))))) == int32(TSParseActionTypeShift) && (*(*TSParseAction)(unsafe.Pointer(actions + uintptr(*(*uint32)(unsafe.Pointer(bp + 16))-uint32(1))*8))).Fshift.Fextra != 0 {
		*(*MutableSubtree)(unsafe.Pointer(bp + 24)) = MutableSubtree{}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 24)) = ts_subtree_make_mut(tls, self+1224, lookahead)
		ts_subtree_set_extra(tls, bp+24, libc.BoolUint8(1 != 0))
		lookahead = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 24)))
	}
	if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
		libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2395, libc.VaList(bp+72, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, lookahead))))
		ts_parser__log(tls, self)
	}
	*(*SubtreeArray)(unsafe.Pointer(bp + 32)) = SubtreeArray{}
	(*SubtreeArray)(unsafe.Pointer(bp + 32)).Fcontents = _array__reserve(tls, (*SubtreeArray)(unsafe.Pointer(bp+32)).Fcontents, bp+32+12, libc.Uint64FromInt64(8), uint32(1))
	(*SubtreeArray)(unsafe.Pointer(bp + 32)).Fcontents = _array__grow(tls, (*SubtreeArray)(unsafe.Pointer(bp+32)).Fcontents, (*SubtreeArray)(unsafe.Pointer(bp+32)).Fsize, bp+32+12, uint32(1), libc.Uint64FromInt64(8))
	v6 = bp + 32 + 8
	v5 = *(*uint32_t)(unsafe.Pointer(v6))
	*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
	*(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(bp+32)).Fcontents + uintptr(v5)*8)) = lookahead
	error_repeat = ts_subtree_new_node(tls, libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1)), bp+32, uint32(0), (*TSParser)(unsafe.Pointer(self)).Flanguage)
	if node_count_since_error > uint32(0) {
		*(*StackSliceArray)(unsafe.Pointer(bp + 48)) = ts_stack_pop_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, uint32(1))
		if (*(*StackSliceArray)(unsafe.Pointer(bp + 48))).Fsize > uint32(1) {
			i2 = uint32(1)
			for {
				if !(i2 < (*(*StackSliceArray)(unsafe.Pointer(bp + 48))).Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(i2 < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+1004, __ccgo_ts+1033, int32(1403), uintptr(unsafe.Pointer(&__func__22)))
					}
				}
				ts_subtree_array_delete(tls, self+1224, (*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents+uintptr(i2)*24)
				goto _7
			_7:
				;
				i2 = i2 + 1
			}
			for {
				_ = libc.Uint64FromInt64(4)
				{
					if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1405), uintptr(unsafe.Pointer(&__func__22)))
					}
				}
				if !(ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack) > (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents)).Fversion+uint32(1)) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1406), uintptr(unsafe.Pointer(&__func__22)))
					}
				}
				ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents)).Fversion+uint32(1))
			}
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1410), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		ts_stack_renumber_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents)).Fversion, version)
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1411), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1411), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1411), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1411), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		(*SubtreeArray)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp + 48)).Fcontents)).Fcontents = _array__grow(tls, (*SubtreeArray)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents)).Fcontents, (*SubtreeArray)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents)).Fsize, (*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents+12, uint32(1), libc.Uint64FromInt64(8))
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1411), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1411), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		v6 = (*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents + 8
		v5 = *(*uint32_t)(unsafe.Pointer(v6))
		*(*uint32_t)(unsafe.Pointer(v6)) = *(*uint32_t)(unsafe.Pointer(v6)) + 1
		*(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents)).Fcontents + uintptr(v5)*8)) = ts_subtree_from_mut(tls, error_repeat)
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+48)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+1033, int32(1414), uintptr(unsafe.Pointer(&__func__22)))
			}
		}
		error_repeat = ts_subtree_new_node(tls, libc.Uint16FromInt32(libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1)), (*StackSliceArray)(unsafe.Pointer(bp+48)).Fcontents, uint32(0), (*TSParser)(unsafe.Pointer(self)).Flanguage)
	}
	ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, ts_subtree_from_mut(tls, error_repeat), libc.BoolUint8(0 != 0), uint16(0))
	if ts_subtree_has_external_tokens(tls, lookahead) != 0 {
		ts_stack_set_last_external_token(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, ts_subtree_last_external_token(tls, lookahead))
	}
	has_error = libc.BoolUint8(1 != 0)
	i3 = uint32(0)
	for {
		if !(i3 < ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)) {
			break
		}
		status = ts_parser__version_status(tls, self, i3)
		if !(status.Fis_in_error != 0) {
			has_error = libc.BoolUint8(0 != 0)
			break
		}
		goto _10
	_10:
		;
		i3 = i3 + 1
	}
	(*TSParser)(unsafe.Pointer(self)).Fhas_error = has_error
}

var __func__22 = [19]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'r', 'e', 'c', 'o', 'v', 'e', 'r'}

func ts_parser__handle_error(tls *libc.TLS, self uintptr, version StackVersion, _lookahead Subtree) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*Subtree)(unsafe.Pointer(bp)) = _lookahead
	var did_insert_missing_token, did_merge uint8
	var i, v3 uint32
	var lookahead_bytes, previous_version_count, version_count uint32_t
	var missing_symbol TSSymbol
	var missing_tree Subtree
	var padding, position Length
	var state, state_after_missing_symbol TSStateId
	var v, version_with_missing_tree StackVersion
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = did_insert_missing_token, did_merge, i, lookahead_bytes, missing_symbol, missing_tree, padding, position, previous_version_count, state, state_after_missing_symbol, v, version_count, version_with_missing_tree, v3
	previous_version_count = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	ts_parser__do_all_potential_reductions(tls, self, version, uint16(0))
	version_count = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	position = ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	did_insert_missing_token = libc.BoolUint8(0 != 0)
	v = version
	for {
		if !(v < version_count) {
			break
		}
		if !(did_insert_missing_token != 0) {
			state = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, v)
			missing_symbol = uint16(1)
			for {
				if !(libc.Int32FromUint16(missing_symbol) < libc.Int32FromUint16(uint16((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Ftoken_count))) {
					break
				}
				state_after_missing_symbol = ts_language_next_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, state, missing_symbol)
				if libc.Int32FromUint16(state_after_missing_symbol) == 0 || libc.Int32FromUint16(state_after_missing_symbol) == libc.Int32FromUint16(state) {
					goto _2
				}
				if ts_language_has_reduce_action(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, state_after_missing_symbol, ts_subtree_leaf_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))) != 0 {
					ts_lexer_reset(tls, self, position)
					ts_lexer_mark_end(tls, self)
					padding = length_sub(tls, (*TSParser)(unsafe.Pointer(self)).Flexer.Ftoken_end_position, position)
					lookahead_bytes = ts_subtree_total_bytes(tls, *(*Subtree)(unsafe.Pointer(bp))) + ts_subtree_lookahead_bytes(tls, *(*Subtree)(unsafe.Pointer(bp)))
					version_with_missing_tree = ts_stack_copy_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, v)
					missing_tree = ts_subtree_new_missing_leaf(tls, self+1224, missing_symbol, padding, lookahead_bytes, (*TSParser)(unsafe.Pointer(self)).Flanguage)
					ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version_with_missing_tree, missing_tree, libc.BoolUint8(0 != 0), state_after_missing_symbol)
					if ts_parser__do_all_potential_reductions(tls, self, version_with_missing_tree, ts_subtree_leaf_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))) != 0 {
						if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
							libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2416, libc.VaList(bp+16, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, missing_symbol), libc.Int32FromUint16(ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version_with_missing_tree))))
							ts_parser__log(tls, self)
						}
						did_insert_missing_token = libc.BoolUint8(1 != 0)
						break
					}
				}
				goto _2
			_2:
				;
				missing_symbol = missing_symbol + 1
			}
		}
		ts_stack_push(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, v, Subtree{}, libc.BoolUint8(0 != 0), uint16(0))
		if v == version {
			v3 = previous_version_count
		} else {
			v3 = v + uint32(1)
		}
		v = v3
		goto _1
	_1:
	}
	i = previous_version_count
	for {
		if !(i < version_count) {
			break
		}
		did_merge = ts_stack_merge(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, previous_version_count)
		_ = libc.Uint64FromInt64(4)
		{
			if !(did_merge != 0) {
				libc.X__assert_fail(tls, __ccgo_ts+2457, __ccgo_ts+1033, int32(1518), uintptr(unsafe.Pointer(&__func__23)))
			}
		}
		goto _4
	_4:
		;
		i = i + 1
	}
	ts_stack_record_summary(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, MAX_SUMMARY_DEPTH)
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
		ts_parser__breakdown_lookahead(tls, self, bp, uint16(0), self+1368)
	}
	ts_parser__recover(tls, self, version, *(*Subtree)(unsafe.Pointer(bp)))
	if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
		ts_stack_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
		libc.Xfputs(tls, __ccgo_ts+1187, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
	}
}

var __func__23 = [24]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'h', 'a', 'n', 'd', 'l', 'e', '_', 'e', 'r', 'r', 'o', 'r'}

func ts_parser__check_progress(tls *libc.TLS, self uintptr, lookahead uintptr, position uintptr, operations uint32) (r uint8) {
	*(*uint32)(unsafe.Pointer(self + 1412)) += operations
	if (*TSParser)(unsafe.Pointer(self)).Foperation_count >= OP_COUNT_PER_PARSER_CALLBACK_CHECK {
		(*TSParser)(unsafe.Pointer(self)).Foperation_count = uint32(0)
	}
	if position != libc.UintptrFromInt32(0) {
		(*TSParser)(unsafe.Pointer(self)).Fparse_state.Fcurrent_byte_offset = *(*uint32_t)(unsafe.Pointer(position))
		(*TSParser)(unsafe.Pointer(self)).Fparse_state.Fhas_error = (*TSParser)(unsafe.Pointer(self)).Fhas_error
	}
	if (*TSParser)(unsafe.Pointer(self)).Foperation_count == uint32(0) && ((*TSParser)(unsafe.Pointer(self)).Fparse_options.Fprogress_callback != 0 && (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSParser)(unsafe.Pointer(self)).Fparse_options.Fprogress_callback})))(tls, self+1456) != 0) {
		if lookahead != 0 && *(*uintptr)(unsafe.Pointer(lookahead)) != 0 {
			ts_subtree_release(tls, self+1224, *(*Subtree)(unsafe.Pointer(lookahead)))
		}
		return libc.BoolUint8(0 != 0)
	}
	return libc.BoolUint8(1 != 0)
}

func ts_parser__advance(tls *libc.TLS, self uintptr, version StackVersion, allow_node_reuse uint8) (r uint8) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var did_reduce, did_reuse, end_of_non_terminal_extra, is_fragile, needs_lex uint8
	var i uint32_t
	var last_external_token Subtree
	var last_reduction_version, reduction_version StackVersion
	var next_state TSStateId
	var _ /* action at bp+32 */ TSParseAction
	var _ /* lookahead at bp+8 */ Subtree
	var _ /* mutable_lookahead at bp+40 */ MutableSubtree
	var _ /* position at bp+4 */ uint32_t
	var _ /* state at bp+0 */ TSStateId
	var _ /* table_entry at bp+16 */ TableEntry
	_, _, _, _, _, _, _, _, _, _ = did_reduce, did_reuse, end_of_non_terminal_extra, i, is_fragile, last_external_token, last_reduction_version, needs_lex, next_state, reduction_version
	*(*TSStateId)(unsafe.Pointer(bp)) = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	*(*uint32_t)(unsafe.Pointer(bp + 4)) = ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version).Fbytes
	last_external_token = ts_stack_last_external_token(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
	did_reuse = libc.BoolUint8(1 != 0)
	*(*Subtree)(unsafe.Pointer(bp + 8)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp + 8)) = Subtree{}
	*(*TableEntry)(unsafe.Pointer(bp + 16)) = TableEntry{}
	if allow_node_reuse != 0 {
		*(*Subtree)(unsafe.Pointer(bp + 8)) = ts_parser__reuse_node(tls, self, version, bp, *(*uint32_t)(unsafe.Pointer(bp + 4)), last_external_token, bp+16)
	}
	if !(*(*uintptr)(unsafe.Pointer(bp + 8)) != 0) {
		did_reuse = libc.BoolUint8(0 != 0)
		*(*Subtree)(unsafe.Pointer(bp + 8)) = ts_parser__get_cached_token(tls, self, *(*TSStateId)(unsafe.Pointer(bp)), uint64(*(*uint32_t)(unsafe.Pointer(bp + 4))), last_external_token, bp+16)
	}
	needs_lex = libc.BoolUint8(!(*(*uintptr)(unsafe.Pointer(bp + 8)) != 0))
	for {
		if needs_lex != 0 {
			needs_lex = libc.BoolUint8(0 != 0)
			*(*Subtree)(unsafe.Pointer(bp + 8)) = ts_parser__lex(tls, self, version, *(*TSStateId)(unsafe.Pointer(bp)))
			if (*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error != 0 {
				return libc.BoolUint8(0 != 0)
			}
			if *(*uintptr)(unsafe.Pointer(bp + 8)) != 0 {
				ts_parser__set_cached_token(tls, self, *(*uint32_t)(unsafe.Pointer(bp + 4)), last_external_token, *(*Subtree)(unsafe.Pointer(bp + 8)))
				ts_language_table_entry(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSStateId)(unsafe.Pointer(bp)), ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8))), bp+16)
			} else {
				ts_language_table_entry(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSStateId)(unsafe.Pointer(bp)), uint16(0), bp+16)
			}
		}
		if !(ts_parser__check_progress(tls, self, bp+8, bp+4, uint32(1)) != 0) {
			return libc.BoolUint8(0 != 0)
		}
		did_reduce = libc.BoolUint8(0 != 0)
		last_reduction_version = libc.Uint32FromInt32(-libc.Int32FromInt32(1))
		i = uint32(0)
		for {
			if !(i < (*(*TableEntry)(unsafe.Pointer(bp + 16))).Faction_count) {
				break
			}
			*(*TSParseAction)(unsafe.Pointer(bp + 32)) = TSParseAction{}
			*(*struct {
				Freduce [0]struct {
					Ftype1              uint8
					Fchild_count        uint8
					Fsymbol             uint16
					Fdynamic_precedence int16
					Fproduction_id      uint16
				}
				Ftype1 [0]uint8
				Fshift struct {
					Ftype1      uint8
					Fstate      uint16
					Fextra      uint8
					Frepetition uint8
				}
				F__ccgo_pad3 [2]byte
			})(unsafe.Pointer(bp + 32)) = *(*TSParseAction)(unsafe.Pointer((*(*TableEntry)(unsafe.Pointer(bp + 16))).Factions + uintptr(i)*8))
			switch libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(bp + 32))))) {
			case int32(TSParseActionTypeShift):
				if (*(*TSParseAction)(unsafe.Pointer(bp + 32))).Fshift.Frepetition != 0 {
					break
				}
				if (*(*TSParseAction)(unsafe.Pointer(bp + 32))).Fshift.Fextra != 0 {
					next_state = *(*TSStateId)(unsafe.Pointer(bp))
					if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
						libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2467, 0)
						ts_parser__log(tls, self)
					}
				} else {
					next_state = (*(*TSParseAction)(unsafe.Pointer(bp + 32))).Fshift.Fstate
					if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
						libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2479, libc.VaList(bp+56, libc.Int32FromUint16(next_state)))
						ts_parser__log(tls, self)
					}
				}
				if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) > uint32(0) {
					ts_parser__breakdown_lookahead(tls, self, bp+8, *(*TSStateId)(unsafe.Pointer(bp)), self+1368)
					next_state = ts_language_next_state(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSStateId)(unsafe.Pointer(bp)), ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8))))
				}
				ts_parser__shift(tls, self, version, next_state, *(*Subtree)(unsafe.Pointer(bp + 8)), (*(*TSParseAction)(unsafe.Pointer(bp + 32))).Fshift.Fextra)
				if did_reuse != 0 {
					reusable_node_advance(tls, self+1368)
				}
				return libc.BoolUint8(1 != 0)
			case int32(TSParseActionTypeReduce):
				is_fragile = libc.BoolUint8((*(*TableEntry)(unsafe.Pointer(bp + 16))).Faction_count > uint32(1))
				end_of_non_terminal_extra = libc.BoolUint8(*(*uintptr)(unsafe.Pointer(bp + 8)) == libc.UintptrFromInt32(0))
				if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
					libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2494, libc.VaList(bp+56, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*(*struct {
						Ftype1              uint8_t
						Fchild_count        uint8_t
						Fsymbol             TSSymbol
						Fdynamic_precedence int16_t
						Fproduction_id      uint16_t
					})(unsafe.Pointer(bp + 32))).Fsymbol), libc.Int32FromUint8((*(*struct {
						Ftype1              uint8_t
						Fchild_count        uint8_t
						Fsymbol             TSSymbol
						Fdynamic_precedence int16_t
						Fproduction_id      uint16_t
					})(unsafe.Pointer(bp + 32))).Fchild_count)))
					ts_parser__log(tls, self)
				}
				reduction_version = ts_parser__reduce(tls, self, version, (*(*struct {
					Ftype1              uint8_t
					Fchild_count        uint8_t
					Fsymbol             TSSymbol
					Fdynamic_precedence int16_t
					Fproduction_id      uint16_t
				})(unsafe.Pointer(bp + 32))).Fsymbol, uint32((*(*struct {
					Ftype1              uint8_t
					Fchild_count        uint8_t
					Fsymbol             TSSymbol
					Fdynamic_precedence int16_t
					Fproduction_id      uint16_t
				})(unsafe.Pointer(bp + 32))).Fchild_count), int32((*(*struct {
					Ftype1              uint8_t
					Fchild_count        uint8_t
					Fsymbol             TSSymbol
					Fdynamic_precedence int16_t
					Fproduction_id      uint16_t
				})(unsafe.Pointer(bp + 32))).Fdynamic_precedence), (*(*struct {
					Ftype1              uint8_t
					Fchild_count        uint8_t
					Fsymbol             TSSymbol
					Fdynamic_precedence int16_t
					Fproduction_id      uint16_t
				})(unsafe.Pointer(bp + 32))).Fproduction_id, is_fragile, end_of_non_terminal_extra)
				did_reduce = libc.BoolUint8(1 != 0)
				if reduction_version != libc.Uint32FromInt32(-libc.Int32FromInt32(1)) {
					last_reduction_version = reduction_version
				}
			case int32(TSParseActionTypeAccept):
				if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
					libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2524, 0)
					ts_parser__log(tls, self)
				}
				ts_parser__accept(tls, self, version, *(*Subtree)(unsafe.Pointer(bp + 8)))
				return libc.BoolUint8(1 != 0)
			case int32(TSParseActionTypeRecover):
				if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) > uint32(0) {
					ts_parser__breakdown_lookahead(tls, self, bp+8, uint16(0), self+1368)
				}
				ts_parser__recover(tls, self, version, *(*Subtree)(unsafe.Pointer(bp + 8)))
				if did_reuse != 0 {
					reusable_node_advance(tls, self+1368)
				}
				return libc.BoolUint8(1 != 0)
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		if last_reduction_version != libc.Uint32FromInt32(-libc.Int32FromInt32(1)) {
			ts_stack_renumber_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, last_reduction_version, version)
			if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				ts_stack_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
				libc.Xfputs(tls, __ccgo_ts+1187, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
			}
			*(*TSStateId)(unsafe.Pointer(bp)) = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
			if !(*(*uintptr)(unsafe.Pointer(bp + 8)) != 0) {
				needs_lex = libc.BoolUint8(1 != 0)
			} else {
				ts_language_table_entry(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSStateId)(unsafe.Pointer(bp)), ts_subtree_leaf_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8))), bp+16)
			}
			goto _1
		}
		if did_reduce != 0 {
			if *(*uintptr)(unsafe.Pointer(bp + 8)) != 0 {
				ts_subtree_release(tls, self+1224, *(*Subtree)(unsafe.Pointer(bp + 8)))
			}
			ts_stack_halt(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
			return libc.BoolUint8(1 != 0)
		}
		if ts_subtree_is_keyword(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 && libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) != libc.Int32FromUint16((*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fkeyword_capture_token) && !(ts_language_is_reserved_word(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSStateId)(unsafe.Pointer(bp)), ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) != 0) {
			ts_language_table_entry(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, *(*TSStateId)(unsafe.Pointer(bp)), (*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fkeyword_capture_token, bp+16)
			if (*(*TableEntry)(unsafe.Pointer(bp + 16))).Faction_count > uint32(0) {
				if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
					libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2531, libc.VaList(bp+56, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))), ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fkeyword_capture_token)))
					ts_parser__log(tls, self)
				}
				*(*MutableSubtree)(unsafe.Pointer(bp + 40)) = MutableSubtree{}
				*(*struct {
					Fptr  [0]uintptr
					Fdata SubtreeInlineData
				})(unsafe.Pointer(bp + 40)) = ts_subtree_make_mut(tls, self+1224, *(*Subtree)(unsafe.Pointer(bp + 8)))
				ts_subtree_set_symbol(tls, bp+40, (*TSLanguage)(unsafe.Pointer((*TSParser)(unsafe.Pointer(self)).Flanguage)).Fkeyword_capture_token, (*TSParser)(unsafe.Pointer(self)).Flanguage)
				*(*Subtree)(unsafe.Pointer(bp + 8)) = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 40)))
				goto _1
			}
		}
		if ts_parser__breakdown_top_of_stack(tls, self, version) != 0 {
			*(*TSStateId)(unsafe.Pointer(bp)) = ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)
			ts_subtree_release(tls, self+1224, *(*Subtree)(unsafe.Pointer(bp + 8)))
			needs_lex = libc.BoolUint8(1 != 0)
			goto _1
		}
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2572, libc.VaList(bp+56, ts_language_symbol_name(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage, ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8))))))
			ts_parser__log(tls, self)
		}
		ts_stack_pause(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version, *(*Subtree)(unsafe.Pointer(bp + 8)))
		return libc.BoolUint8(1 != 0)
		goto _1
	_1:
	}
	return r
}

func ts_parser__condense_stack(tls *libc.TLS, self uintptr) (r uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var has_unpaused_version, made_changes uint8
	var i, i1, j, n StackVersion
	var lookahead Subtree
	var min_error_cost uint32
	var status_i, status_j ErrorStatus
	_, _, _, _, _, _, _, _, _, _ = has_unpaused_version, i, i1, j, lookahead, made_changes, min_error_cost, n, status_i, status_j
	made_changes = libc.BoolUint8(0 != 0)
	min_error_cost = libc.Uint32FromInt32(0x7fffffff)*libc.Uint32FromUint32(2) + libc.Uint32FromUint32(1)
	i = uint32(0)
	for {
		if !(i < ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)) {
			break
		}
		if ts_stack_is_halted(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i) != 0 {
			ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i)
			i = i - 1
			goto _1
		}
		status_i = ts_parser__version_status(tls, self, i)
		if !(status_i.Fis_in_error != 0) && status_i.Fcost < min_error_cost {
			min_error_cost = status_i.Fcost
		}
		j = uint32(0)
		for {
			if !(j < i) {
				break
			}
			status_j = ts_parser__version_status(tls, self, j)
			switch ts_parser__compare_versions(tls, self, status_j, status_i) {
			case int32(ErrorComparisonTakeLeft):
				made_changes = libc.BoolUint8(1 != 0)
				ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i)
				i = i - 1
				j = i
			case int32(ErrorComparisonPreferLeft):
				fallthrough
			case int32(ErrorComparisonNone):
				if ts_stack_merge(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, j, i) != 0 {
					made_changes = libc.BoolUint8(1 != 0)
					i = i - 1
					j = i
				}
			case int32(ErrorComparisonPreferRight):
				made_changes = libc.BoolUint8(1 != 0)
				if ts_stack_merge(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, j, i) != 0 {
					i = i - 1
					j = i
				} else {
					ts_stack_swap_versions(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i, j)
				}
			case int32(ErrorComparisonTakeRight):
				made_changes = libc.BoolUint8(1 != 0)
				ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, j)
				i = i - 1
				j = j - 1
				break
			}
			goto _2
		_2:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	for ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack) > MAX_VERSION_COUNT {
		ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, MAX_VERSION_COUNT)
		made_changes = libc.BoolUint8(1 != 0)
	}
	if ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack) > uint32(0) {
		has_unpaused_version = libc.BoolUint8(0 != 0)
		i1 = uint32(0)
		n = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
		for {
			if !(i1 < n) {
				break
			}
			if ts_stack_is_paused(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i1) != 0 {
				if !(has_unpaused_version != 0) && (*TSParser)(unsafe.Pointer(self)).Faccept_count < MAX_VERSION_COUNT {
					if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
						libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2598, libc.VaList(bp+8, i1))
						ts_parser__log(tls, self)
					}
					min_error_cost = ts_stack_error_cost(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i1)
					lookahead = ts_stack_resume(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i1)
					ts_parser__handle_error(tls, self, i1, lookahead)
					has_unpaused_version = libc.BoolUint8(1 != 0)
				} else {
					ts_stack_remove_version(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, i1)
					made_changes = libc.BoolUint8(1 != 0)
					i1 = i1 - 1
					n = n - 1
				}
			} else {
				has_unpaused_version = libc.BoolUint8(1 != 0)
			}
			goto _3
		_3:
			;
			i1 = i1 + 1
		}
	}
	if made_changes != 0 {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2616, 0)
			ts_parser__log(tls, self)
		}
		if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			ts_stack_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
			libc.Xfputs(tls, __ccgo_ts+1187, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
		}
	}
	return min_error_cost
}

func ts_parser__balance_subtree(tls *libc.TLS, self uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var child1, child2 Subtree
	var i, n, v6 uint32
	var i1, v1 uint32_t
	var operations uint8_t
	var repeat_delta int64
	var v2, v3 uintptr
	var _ /* child at bp+16 */ Subtree
	var _ /* finished_tree at bp+0 */ Subtree
	var _ /* tree at bp+8 */ MutableSubtree
	_, _, _, _, _, _, _, _, _, _, _ = child1, child2, i, i1, n, operations, repeat_delta, v1, v2, v3, v6
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = (*TSParser)(unsafe.Pointer(self)).Ffinished_tree
	if !((*TSParser)(unsafe.Pointer(self)).Fcanceled_balancing != 0) {
		(*MutableSubtreeArray)(unsafe.Pointer(self + 1224 + 16)).Fsize = uint32(0)
		if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) && libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp))) == uint32(1) {
			(*MutableSubtreeArray)(unsafe.Pointer(self + 1224 + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fsize, self+1224+16+12, uint32(1), libc.Uint64FromInt64(8))
			v2 = self + 1224 + 16 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(bp)))
		}
	}
	for (*TSParser)(unsafe.Pointer(self)).Ftree_pool.Ftree_stack.Fsize > uint32(0) {
		if !(ts_parser__check_progress(tls, self, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), uint32(1)) != 0) {
			return libc.BoolUint8(0 != 0)
		}
		*(*MutableSubtree)(unsafe.Pointer(bp + 8)) = MutableSubtree{}
		_ = libc.Uint64FromInt64(4)
		{
			if !((*TSParser)(unsafe.Pointer(self)).Ftree_pool.Ftree_stack.Fsize-libc.Uint32FromInt32(1) < (*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2625, __ccgo_ts+1033, int32(1885), uintptr(unsafe.Pointer(&__func__24)))
			}
		}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 8)) = *(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fcontents + uintptr((*TSParser)(unsafe.Pointer(self)).Ftree_pool.Ftree_stack.Fsize-uint32(1))*8))
		if libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).F__ccgo19_48.F__ccgo0_0.Frepeat_depth) > 0 {
			if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 {
				v2 = libc.UintptrFromInt32(0)
			} else {
				v2 = *(*uintptr)(unsafe.Pointer(bp + 8)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count)*8
			}
			child1 = *(*Subtree)(unsafe.Pointer(v2))
			if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 {
				v3 = libc.UintptrFromInt32(0)
			} else {
				v3 = *(*uintptr)(unsafe.Pointer(bp + 8)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count)*8
			}
			child2 = *(*Subtree)(unsafe.Pointer(v3 + uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count-uint32(1))*8))
			repeat_delta = libc.Int64FromUint32(ts_subtree_repeat_depth(tls, child1)) - libc.Int64FromUint32(ts_subtree_repeat_depth(tls, child2))
			if repeat_delta > 0 {
				n = libc.Uint32FromInt64(repeat_delta)
				i = n / uint32(2)
				for {
					if !(i > uint32(0)) {
						break
					}
					ts_subtree_compress(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 8)), i, (*TSParser)(unsafe.Pointer(self)).Flanguage, self+1224+16)
					n = n - i
					if i>>int32(4) > uint32(0) {
						v6 = i >> int32(4)
					} else {
						v6 = uint32(1)
					}
					operations = uint8(v6)
					if !(ts_parser__check_progress(tls, self, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), uint32(operations)) != 0) {
						return libc.BoolUint8(0 != 0)
					}
					goto _5
				_5:
					;
					i = i / uint32(2)
				}
			}
		}
		v2 = self + 1224 + 16 + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		_ = *(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fcontents + uintptr(v1)*8))
		i1 = uint32(0)
		for {
			if !(i1 < (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count) {
				break
			}
			*(*Subtree)(unsafe.Pointer(bp + 16)) = Subtree{}
			if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 {
				v2 = libc.UintptrFromInt32(0)
			} else {
				v2 = *(*uintptr)(unsafe.Pointer(bp + 8)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count)*8
			}
			*(*struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			})(unsafe.Pointer(bp + 16)) = *(*Subtree)(unsafe.Pointer(v2 + uintptr(i1)*8))
			if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 16))) > uint32(0) && libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp + 16))) == uint32(1) {
				(*MutableSubtreeArray)(unsafe.Pointer(self + 1224 + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fsize, self+1224+16+12, uint32(1), libc.Uint64FromInt64(8))
				v2 = self + 1224 + 16 + 8
				v1 = *(*uint32_t)(unsafe.Pointer(v2))
				*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
				*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(self+1224+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(bp + 16)))
			}
			goto _9
		_9:
			;
			i1 = i1 + 1
		}
	}
	return libc.BoolUint8(1 != 0)
}

var __func__24 = [27]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', '_', 'b', 'a', 'l', 'a', 'n', 'c', 'e', '_', 's', 'u', 'b', 't', 'r', 'e', 'e'}

func ts_parser_has_outstanding_parse(tls *libc.TLS, self uintptr) (r uint8) {
	return libc.BoolUint8((*TSParser)(unsafe.Pointer(self)).Fcanceled_balancing != 0 || (*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload != 0 || libc.Int32FromUint16(ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, uint32(0))) != int32(1) || ts_stack_node_count_since_error(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, uint32(0)) != uint32(0))
}

func ts_parser_new(tls *libc.TLS) (r uintptr) {
	var self uintptr
	_ = self
	self = (*(*func(*libc.TLS, size_t, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_calloc})))(tls, uint64(1), uint64(1480))
	ts_lexer_init(tls, self)
	(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fsize = uint32(0)
	(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fcapacity = uint32(0)
	(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fcontents = libc.UintptrFromInt32(0)
	(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fcontents = _array__reserve(tls, (*ReduceActionSet)(unsafe.Pointer(self+1272)).Fcontents, self+1272+12, libc.Uint64FromInt64(16), uint32(4))
	(*TSParser)(unsafe.Pointer(self)).Ftree_pool = ts_subtree_pool_new(tls, uint32(32))
	(*TSParser)(unsafe.Pointer(self)).Fstack = ts_stack_new(tls, self+1224)
	(*TSParser)(unsafe.Pointer(self)).Ffinished_tree = Subtree{}
	(*TSParser)(unsafe.Pointer(self)).Freusable_node = reusable_node_new(tls)
	(*TSParser)(unsafe.Pointer(self)).Fdot_graph_file = libc.UintptrFromInt32(0)
	(*TSParser)(unsafe.Pointer(self)).Flanguage = libc.UintptrFromInt32(0)
	(*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error = libc.BoolUint8(0 != 0)
	(*TSParser)(unsafe.Pointer(self)).Fhas_error = libc.BoolUint8(0 != 0)
	(*TSParser)(unsafe.Pointer(self)).Fcanceled_balancing = libc.BoolUint8(0 != 0)
	(*TSParser)(unsafe.Pointer(self)).Fexternal_scanner_payload = libc.UintptrFromInt32(0)
	(*TSParser)(unsafe.Pointer(self)).Foperation_count = uint32(0)
	(*TSParser)(unsafe.Pointer(self)).Fold_tree = Subtree{}
	(*TSParser)(unsafe.Pointer(self)).Fincluded_range_differences = TSRangeArray{}
	(*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index = uint32(0)
	ts_parser__set_cached_token(tls, self, uint32(0), Subtree{}, Subtree{})
	return self
}

func ts_parser_delete(tls *libc.TLS, self uintptr) {
	if !(self != 0) {
		return
	}
	ts_parser_set_language(tls, self, libc.UintptrFromInt32(0))
	ts_stack_delete(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	if (*TSParser)(unsafe.Pointer(self)).Freduce_actions.Fcontents != 0 {
		if (*ReduceActionSet)(unsafe.Pointer(self+1272)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*ReduceActionSet)(unsafe.Pointer(self+1272)).Fcontents)
		}
		(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fcontents = libc.UintptrFromInt32(0)
		(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fsize = uint32(0)
		(*ReduceActionSet)(unsafe.Pointer(self + 1272)).Fcapacity = uint32(0)
	}
	if (*TSParser)(unsafe.Pointer(self)).Fincluded_range_differences.Fcontents != 0 {
		if (*TSRangeArray)(unsafe.Pointer(self+1424)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*TSRangeArray)(unsafe.Pointer(self+1424)).Fcontents)
		}
		(*TSRangeArray)(unsafe.Pointer(self + 1424)).Fcontents = libc.UintptrFromInt32(0)
		(*TSRangeArray)(unsafe.Pointer(self + 1424)).Fsize = uint32(0)
		(*TSRangeArray)(unsafe.Pointer(self + 1424)).Fcapacity = uint32(0)
	}
	if *(*uintptr)(unsafe.Pointer(self + 1416)) != 0 {
		ts_subtree_release(tls, self+1224, (*TSParser)(unsafe.Pointer(self)).Fold_tree)
		(*TSParser)(unsafe.Pointer(self)).Fold_tree = Subtree{}
	}
	ts_wasm_store_delete(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store)
	ts_lexer_delete(tls, self)
	ts_parser__set_cached_token(tls, self, uint32(0), Subtree{}, Subtree{})
	ts_subtree_pool_delete(tls, self+1224)
	reusable_node_delete(tls, self+1368)
	if (*SubtreeArray)(unsafe.Pointer(self+1296)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*SubtreeArray)(unsafe.Pointer(self+1296)).Fcontents)
	}
	(*SubtreeArray)(unsafe.Pointer(self + 1296)).Fcontents = libc.UintptrFromInt32(0)
	(*SubtreeArray)(unsafe.Pointer(self + 1296)).Fsize = uint32(0)
	(*SubtreeArray)(unsafe.Pointer(self + 1296)).Fcapacity = uint32(0)
	if (*SubtreeArray)(unsafe.Pointer(self+1312)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*SubtreeArray)(unsafe.Pointer(self+1312)).Fcontents)
	}
	(*SubtreeArray)(unsafe.Pointer(self + 1312)).Fcontents = libc.UintptrFromInt32(0)
	(*SubtreeArray)(unsafe.Pointer(self + 1312)).Fsize = uint32(0)
	(*SubtreeArray)(unsafe.Pointer(self + 1312)).Fcapacity = uint32(0)
	if (*SubtreeArray)(unsafe.Pointer(self+1328)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*SubtreeArray)(unsafe.Pointer(self+1328)).Fcontents)
	}
	(*SubtreeArray)(unsafe.Pointer(self + 1328)).Fcontents = libc.UintptrFromInt32(0)
	(*SubtreeArray)(unsafe.Pointer(self + 1328)).Fsize = uint32(0)
	(*SubtreeArray)(unsafe.Pointer(self + 1328)).Fcapacity = uint32(0)
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, self)
}

func ts_parser_language(tls *libc.TLS, self uintptr) (r uintptr) {
	return (*TSParser)(unsafe.Pointer(self)).Flanguage
}

func ts_parser_set_language(tls *libc.TLS, self uintptr, language uintptr) (r uint8) {
	ts_parser_reset(tls, self)
	ts_language_delete(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage)
	(*TSParser)(unsafe.Pointer(self)).Flanguage = libc.UintptrFromInt32(0)
	if language != 0 {
		if (*TSLanguage)(unsafe.Pointer(language)).Fabi_version > uint32(15) || (*TSLanguage)(unsafe.Pointer(language)).Fabi_version < uint32(13) {
			return libc.BoolUint8(0 != 0)
		}
		if ts_language_is_wasm(tls, language) != 0 {
			if !((*TSParser)(unsafe.Pointer(self)).Fwasm_store != 0) || !(ts_wasm_store_start(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store, self, language) != 0) {
				return libc.BoolUint8(0 != 0)
			}
		}
	}
	(*TSParser)(unsafe.Pointer(self)).Flanguage = ts_language_copy(tls, language)
	return libc.BoolUint8(1 != 0)
}

func ts_parser_logger(tls *libc.TLS, self uintptr) (r TSLogger) {
	return (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger
}

func ts_parser_set_logger(tls *libc.TLS, self uintptr, logger TSLogger) {
	(*TSParser)(unsafe.Pointer(self)).Flexer.Flogger = logger
}

func ts_parser_print_dot_graphs(tls *libc.TLS, self uintptr, fd int32) {
	if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
		libc.Xfclose(tls, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
	}
	if fd >= 0 {
		(*TSParser)(unsafe.Pointer(self)).Fdot_graph_file = libc.Xfdopen(tls, fd, __ccgo_ts+2711)
	} else {
		(*TSParser)(unsafe.Pointer(self)).Fdot_graph_file = libc.UintptrFromInt32(0)
	}
}

func ts_parser_set_included_ranges(tls *libc.TLS, self uintptr, ranges uintptr, count uint32_t) (r uint8) {
	return ts_lexer_set_included_ranges(tls, self, ranges, count)
}

func ts_parser_included_ranges(tls *libc.TLS, self uintptr, count uintptr) (r uintptr) {
	return ts_lexer_included_ranges(tls, self, count)
}

func ts_parser_reset(tls *libc.TLS, self uintptr) {
	ts_parser__external_scanner_destroy(tls, self)
	if (*TSParser)(unsafe.Pointer(self)).Fwasm_store != 0 {
		ts_wasm_store_reset(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store)
	}
	if *(*uintptr)(unsafe.Pointer(self + 1416)) != 0 {
		ts_subtree_release(tls, self+1224, (*TSParser)(unsafe.Pointer(self)).Fold_tree)
		(*TSParser)(unsafe.Pointer(self)).Fold_tree = Subtree{}
	}
	reusable_node_clear(tls, self+1368)
	ts_lexer_reset(tls, self, length_zero(tls))
	ts_stack_clear(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
	ts_parser__set_cached_token(tls, self, uint32(0), Subtree{}, Subtree{})
	if *(*uintptr)(unsafe.Pointer(self + 1288)) != 0 {
		ts_subtree_release(tls, self+1224, (*TSParser)(unsafe.Pointer(self)).Ffinished_tree)
		(*TSParser)(unsafe.Pointer(self)).Ffinished_tree = Subtree{}
	}
	(*TSParser)(unsafe.Pointer(self)).Faccept_count = uint32(0)
	(*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error = libc.BoolUint8(0 != 0)
	(*TSParser)(unsafe.Pointer(self)).Fhas_error = libc.BoolUint8(0 != 0)
	(*TSParser)(unsafe.Pointer(self)).Fcanceled_balancing = libc.BoolUint8(0 != 0)
	(*TSParser)(unsafe.Pointer(self)).Fparse_options = TSParseOptions{}
	(*TSParser)(unsafe.Pointer(self)).Fparse_state = TSParseState{}
}

func ts_parser_parse(tls *libc.TLS, self uintptr, old_tree uintptr, input TSInput) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var allow_node_reuse uint8
	var i, min_error_cost uint32
	var last_position, position, version_count uint32_t
	var range1, range11, result uintptr
	var version StackVersion
	_, _, _, _, _, _, _, _, _, _ = allow_node_reuse, i, last_position, min_error_cost, position, range1, range11, result, version, version_count
	result = libc.UintptrFromInt32(0)
	if !((*TSParser)(unsafe.Pointer(self)).Flanguage != 0) || !(input.Fread != 0) {
		return libc.UintptrFromInt32(0)
	}
	if ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		if !((*TSParser)(unsafe.Pointer(self)).Fwasm_store != 0) {
			return libc.UintptrFromInt32(0)
		}
		ts_wasm_store_start(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store, self, (*TSParser)(unsafe.Pointer(self)).Flanguage)
	}
	ts_lexer_set_input(tls, self, input)
	(*TSRangeArray)(unsafe.Pointer(self + 1424)).Fsize = uint32(0)
	(*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index = uint32(0)
	(*TSParser)(unsafe.Pointer(self)).Foperation_count = uint32(0)
	if ts_parser_has_outstanding_parse(tls, self) != 0 {
		if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
			libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2713, 0)
			ts_parser__log(tls, self)
		}
		if (*TSParser)(unsafe.Pointer(self)).Fcanceled_balancing != 0 {
			goto balance
		}
	} else {
		ts_parser__external_scanner_create(tls, self)
		if (*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error != 0 {
			goto exit
		}
		if old_tree != 0 {
			ts_subtree_retain(tls, (*TSTree)(unsafe.Pointer(old_tree)).Froot)
			(*TSParser)(unsafe.Pointer(self)).Fold_tree = (*TSTree)(unsafe.Pointer(old_tree)).Froot
			ts_range_array_get_changed_ranges(tls, (*TSTree)(unsafe.Pointer(old_tree)).Fincluded_ranges, (*TSTree)(unsafe.Pointer(old_tree)).Fincluded_range_count, (*TSParser)(unsafe.Pointer(self)).Flexer.Fincluded_ranges, (*TSParser)(unsafe.Pointer(self)).Flexer.Fincluded_range_count, self+1424)
			reusable_node_reset(tls, self+1368, (*TSTree)(unsafe.Pointer(old_tree)).Froot)
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2728, 0)
				ts_parser__log(tls, self)
			}
			if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				ts_subtree_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fold_tree, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
				libc.Xfputs(tls, __ccgo_ts+2745, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
			}
			i = uint32(0)
			for {
				if !(i < (*TSParser)(unsafe.Pointer(self)).Fincluded_range_differences.Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(i < (*TSRangeArray)(unsafe.Pointer(self+1424)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+2747, __ccgo_ts+1033, int32(2112), uintptr(unsafe.Pointer(&__func__25)))
					}
				}
				range1 = (*TSRangeArray)(unsafe.Pointer(self+1424)).Fcontents + uintptr(i)*24
				if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
					libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2805, libc.VaList(bp+8, (*TSRange)(unsafe.Pointer(range1)).Fstart_byte, (*TSRange)(unsafe.Pointer(range1)).Fend_byte))
					ts_parser__log(tls, self)
				}
				goto _1
			_1:
				;
				i = i + 1
			}
		} else {
			reusable_node_clear(tls, self+1368)
			if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
				libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2838, 0)
				ts_parser__log(tls, self)
			}
		}
	}
	position = uint32(0)
	last_position = uint32(0)
	version_count = uint32(0)
	for cond := true; cond; cond = version_count != uint32(0) {
		version = uint32(0)
		for {
			version_count = ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
			if !(version < version_count) {
				break
			}
			allow_node_reuse = libc.BoolUint8(version_count == uint32(1))
			for ts_stack_is_active(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version) != 0 {
				if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
					libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+2848, libc.VaList(bp+8, version, ts_stack_version_count(tls, (*TSParser)(unsafe.Pointer(self)).Fstack), libc.Int32FromUint16(ts_stack_state(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version)), ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version).Fextent.Frow, ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version).Fextent.Fcolumn))
					ts_parser__log(tls, self)
				}
				if !(ts_parser__advance(tls, self, version, allow_node_reuse) != 0) {
					if (*TSParser)(unsafe.Pointer(self)).Fhas_scanner_error != 0 {
						goto exit
					}
					return libc.UintptrFromInt32(0)
				}
				if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
					ts_stack_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
					libc.Xfputs(tls, __ccgo_ts+1187, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
				}
				position = ts_stack_position(tls, (*TSParser)(unsafe.Pointer(self)).Fstack, version).Fbytes
				if position > last_position || version > uint32(0) && position == last_position {
					last_position = position
					break
				}
			}
			goto _2
		_2:
			;
			version = version + 1
		}
		min_error_cost = ts_parser__condense_stack(tls, self)
		if *(*uintptr)(unsafe.Pointer(self + 1288)) != 0 && ts_subtree_error_cost(tls, (*TSParser)(unsafe.Pointer(self)).Ffinished_tree) < min_error_cost {
			ts_stack_clear(tls, (*TSParser)(unsafe.Pointer(self)).Fstack)
			break
		}
		for (*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index < (*TSParser)(unsafe.Pointer(self)).Fincluded_range_differences.Fsize {
			_ = libc.Uint64FromInt64(4)
			{
				if !((*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index < (*TSRangeArray)(unsafe.Pointer(self+1424)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+2911, __ccgo_ts+1033, int32(2169), uintptr(unsafe.Pointer(&__func__25)))
				}
			}
			range11 = (*TSRangeArray)(unsafe.Pointer(self+1424)).Fcontents + uintptr((*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index)*24
			if (*TSRange)(unsafe.Pointer(range11)).Fend_byte <= position {
				(*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index = (*TSParser)(unsafe.Pointer(self)).Fincluded_range_difference_index + 1
			} else {
				break
			}
		}
	}
	goto balance
balance:
	;
	_ = libc.Uint64FromInt64(4)
	{
		if !(*(*uintptr)(unsafe.Pointer(self + 1288)) != 0) {
			libc.X__assert_fail(tls, __ccgo_ts+3005, __ccgo_ts+1033, int32(2179), uintptr(unsafe.Pointer(&__func__25)))
		}
	}
	if !(ts_parser__balance_subtree(tls, self) != 0) {
		(*TSParser)(unsafe.Pointer(self)).Fcanceled_balancing = libc.BoolUint8(1 != 0)
		return uintptr(0)
	}
	(*TSParser)(unsafe.Pointer(self)).Fcanceled_balancing = libc.BoolUint8(0 != 0)
	if (*TSParser)(unsafe.Pointer(self)).Flexer.Flogger.Flog != 0 || (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
		libc.X__builtin_snprintf(tls, self+192, uint64(1024), __ccgo_ts+3029, 0)
		ts_parser__log(tls, self)
	}
	if (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file != 0 {
		ts_subtree_print_dot_graph(tls, (*TSParser)(unsafe.Pointer(self)).Ffinished_tree, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
		libc.Xfputs(tls, __ccgo_ts+2745, (*TSParser)(unsafe.Pointer(self)).Fdot_graph_file)
	}
	result = ts_tree_new(tls, (*TSParser)(unsafe.Pointer(self)).Ffinished_tree, (*TSParser)(unsafe.Pointer(self)).Flanguage, (*TSParser)(unsafe.Pointer(self)).Flexer.Fincluded_ranges, (*TSParser)(unsafe.Pointer(self)).Flexer.Fincluded_range_count)
	(*TSParser)(unsafe.Pointer(self)).Ffinished_tree = Subtree{}
	goto exit
exit:
	;
	ts_parser_reset(tls, self)
	return result
}

var __func__25 = [16]int8{'t', 's', '_', 'p', 'a', 'r', 's', 'e', 'r', '_', 'p', 'a', 'r', 's', 'e'}

func ts_parser_parse_with_options(tls *libc.TLS, self uintptr, old_tree uintptr, input TSInput, parse_options TSParseOptions) (r uintptr) {
	var result uintptr
	_ = result
	(*TSParser)(unsafe.Pointer(self)).Fparse_options = parse_options
	(*TSParser)(unsafe.Pointer(self)).Fparse_state.Fpayload = parse_options.Fpayload
	result = ts_parser_parse(tls, self, old_tree, input)
	(*TSParser)(unsafe.Pointer(self)).Fparse_options = TSParseOptions{}
	return result
}

func ts_parser_parse_string(tls *libc.TLS, self uintptr, old_tree uintptr, string1 uintptr, length uint32_t) (r uintptr) {
	return ts_parser_parse_string_encoding(tls, self, old_tree, string1, length, int32(TSInputEncodingUTF8))
}

func ts_parser_parse_string_encoding(tls *libc.TLS, self uintptr, old_tree uintptr, string1 uintptr, length uint32_t, encoding TSInputEncoding1) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* input at bp+0 */ TSStringInput
	*(*TSStringInput)(unsafe.Pointer(bp)) = TSStringInput{
		Fstring1: string1,
		Flength:  length,
	}
	return ts_parser_parse(tls, self, old_tree, TSInput{
		Fpayload:  bp,
		Fread:     __ccgo_fp(ts_string_input_read),
		Fencoding: encoding,
	})
}

func ts_parser_set_wasm_store(tls *libc.TLS, self uintptr, store uintptr) {
	var copy1 uintptr
	_ = copy1
	if (*TSParser)(unsafe.Pointer(self)).Flanguage != 0 && ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		copy1 = ts_language_copy(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage)
		ts_parser_set_language(tls, self, copy1)
		ts_language_delete(tls, copy1)
	}
	ts_wasm_store_delete(tls, (*TSParser)(unsafe.Pointer(self)).Fwasm_store)
	(*TSParser)(unsafe.Pointer(self)).Fwasm_store = store
}

func ts_parser_take_wasm_store(tls *libc.TLS, self uintptr) (r uintptr) {
	var result uintptr
	_ = result
	if (*TSParser)(unsafe.Pointer(self)).Flanguage != 0 && ts_language_is_wasm(tls, (*TSParser)(unsafe.Pointer(self)).Flanguage) != 0 {
		ts_parser_set_language(tls, self, libc.UintptrFromInt32(0))
	}
	result = (*TSParser)(unsafe.Pointer(self)).Fwasm_store
	(*TSParser)(unsafe.Pointer(self)).Fwasm_store = libc.UintptrFromInt32(0)
	return result
}

func ts_point_edit(tls *libc.TLS, point uintptr, byte1 uintptr, edit uintptr) {
	var start_byte uint32_t
	var start_point TSPoint
	_, _ = start_byte, start_point
	start_byte = *(*uint32_t)(unsafe.Pointer(byte1))
	start_point = *(*TSPoint)(unsafe.Pointer(point))
	if start_byte >= (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_byte {
		start_byte = (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_byte + (start_byte - (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_byte)
		start_point = point_add(tls, (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_point, point_sub(tls, start_point, (*TSInputEdit)(unsafe.Pointer(edit)).Fold_end_point))
	} else {
		if start_byte > (*TSInputEdit)(unsafe.Pointer(edit)).Fstart_byte {
			start_byte = (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_byte
			start_point = (*TSInputEdit)(unsafe.Pointer(edit)).Fnew_end_point
		}
	}
	*(*TSPoint)(unsafe.Pointer(point)) = start_point
	*(*uint32_t)(unsafe.Pointer(byte1)) = start_byte
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

type Stream = struct {
	Finput     uintptr
	Fstart     uintptr
	Fend       uintptr
	Fnext      int32_t
	Fnext_size uint8_t
}

type QueryStep = struct {
	Fsymbol                TSSymbol
	Fsupertype_symbol      TSSymbol
	Ffield                 TSFieldId
	Fcapture_ids           [3]uint16_t
	Fdepth                 uint16_t
	Falternative_index     uint16_t
	Fnegated_field_list_id uint16_t
	F__ccgo18              uint8
	F__ccgo19              uint8
}

type Slice = struct {
	Foffset uint32_t
	Flength uint32_t
}

type SymbolTable = struct {
	Fcharacters struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fslices struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
}

type CaptureQuantifiers = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type PatternEntry = struct {
	Fstep_index    uint16_t
	Fpattern_index uint16_t
	Fis_rooted     uint8
}

type QueryPattern = struct {
	Fsteps           Slice
	Fpredicate_steps Slice
	Fstart_byte      uint32_t
	Fend_byte        uint32_t
	Fis_non_local    uint8
}

type StepOffset = struct {
	Fbyte_offset uint32_t
	Fstep_index  uint16_t
}

type QueryState = struct {
	Fid              uint32_t
	Fcapture_list_id uint32_t
	Fstart_depth     uint16_t
	Fstep_index      uint16_t
	Fpattern_index   uint16_t
	F__ccgo14        uint16
}

type CaptureList = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type CaptureListPool = struct {
	Flist struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fempty_list              CaptureList
	Fmax_capture_list_count  uint32_t
	Ffree_capture_list_count uint32_t
}

type AnalysisStateEntry = struct {
	Fparse_state   TSStateId
	Fparent_symbol TSSymbol
	Fchild_index   uint16_t
	F__ccgo6       uint16
}

type AnalysisState = struct {
	Fstack       [8]AnalysisStateEntry
	Fdepth       uint16_t
	Fstep_index  uint16_t
	Froot_symbol TSSymbol
}

type AnalysisStateSet = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type QueryAnalysis = struct {
	Fstates             AnalysisStateSet
	Fnext_states        AnalysisStateSet
	Fdeeper_states      AnalysisStateSet
	Fstate_pool         AnalysisStateSet
	Ffinal_step_indices struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Ffinished_parent_symbols struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fdid_abort uint8
}

type AnalysisSubgraphNode = struct {
	Fstate         TSStateId
	Fproduction_id uint16_t
	F__ccgo4       uint8
}

type AnalysisSubgraph = struct {
	Fsymbol       TSSymbol
	Fstart_states struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	Fnodes struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
}

type AnalysisSubgraphArray = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type StatePredecessorMap = struct {
	Fcontents uintptr
}

var PARENT_DONE = -int32(1)
var PATTERN_DONE_MARKER = libc.Uint16FromInt32(libc.Int32FromInt32(65535))
var NONE = libc.Uint16FromInt32(libc.Int32FromInt32(65535))
var WILDCARD_SYMBOL = uint16(0)
var OP_COUNT_PER_QUERY_CALLBACK_CHECK = uint32(100)

func stream_advance(tls *libc.TLS, self uintptr) (r uint8) {
	var size uint32_t
	_ = size
	*(*uintptr)(unsafe.Pointer(self)) += uintptr((*Stream)(unsafe.Pointer(self)).Fnext_size)
	if (*Stream)(unsafe.Pointer(self)).Finput < (*Stream)(unsafe.Pointer(self)).Fend {
		size = ts_decode_utf8(tls, (*Stream)(unsafe.Pointer(self)).Finput, libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(self)).Fend)-int64((*Stream)(unsafe.Pointer(self)).Finput)), self+24)
		if size > uint32(0) {
			(*Stream)(unsafe.Pointer(self)).Fnext_size = uint8(size)
			return libc.BoolUint8(1 != 0)
		}
	} else {
		(*Stream)(unsafe.Pointer(self)).Fnext_size = uint8(0)
		(*Stream)(unsafe.Pointer(self)).Fnext = int32('\000')
	}
	return libc.BoolUint8(0 != 0)
}

func stream_reset(tls *libc.TLS, self uintptr, input uintptr) {
	(*Stream)(unsafe.Pointer(self)).Finput = input
	(*Stream)(unsafe.Pointer(self)).Fnext_size = uint8(0)
	stream_advance(tls, self)
}

func stream_new(tls *libc.TLS, string1 uintptr, length uint32_t) (r Stream) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* self at bp+0 */ Stream
	*(*Stream)(unsafe.Pointer(bp)) = Stream{
		Finput: string1,
		Fstart: string1,
		Fend:   string1 + uintptr(length),
	}
	stream_advance(tls, bp)
	return *(*Stream)(unsafe.Pointer(bp))
}

func stream_skip_whitespace(tls *libc.TLS, self uintptr) {
	for {
		if libc.Xiswspace(tls, libc.Uint32FromInt32((*Stream)(unsafe.Pointer(self)).Fnext)) != 0 {
			stream_advance(tls, self)
		} else {
			if (*Stream)(unsafe.Pointer(self)).Fnext == int32(';') {
				stream_advance(tls, self)
				for (*Stream)(unsafe.Pointer(self)).Fnext != 0 && (*Stream)(unsafe.Pointer(self)).Fnext != int32('\n') {
					if !(stream_advance(tls, self) != 0) {
						break
					}
				}
			} else {
				break
			}
		}
		goto _1
	_1:
	}
}

func stream_is_ident_start(tls *libc.TLS, self uintptr) (r uint8) {
	return libc.BoolUint8(libc.Xiswalnum(tls, libc.Uint32FromInt32((*Stream)(unsafe.Pointer(self)).Fnext)) != 0 || (*Stream)(unsafe.Pointer(self)).Fnext == int32('_') || (*Stream)(unsafe.Pointer(self)).Fnext == int32('-'))
}

func stream_scan_identifier(tls *libc.TLS, stream uintptr) {
	for cond := true; cond; cond = libc.Xiswalnum(tls, libc.Uint32FromInt32((*Stream)(unsafe.Pointer(stream)).Fnext)) != 0 || (*Stream)(unsafe.Pointer(stream)).Fnext == int32('_') || (*Stream)(unsafe.Pointer(stream)).Fnext == int32('-') || (*Stream)(unsafe.Pointer(stream)).Fnext == int32('.') {
		stream_advance(tls, stream)
	}
}

func stream_offset(tls *libc.TLS, self uintptr) (r uint32_t) {
	return libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(self)).Finput) - int64((*Stream)(unsafe.Pointer(self)).Fstart))
}

func capture_list_pool_new(tls *libc.TLS) (r CaptureListPool) {
	return CaptureListPool{
		Fmax_capture_list_count: libc.Uint32FromUint32(4294967295),
	}
}

func capture_list_pool_reset(tls *libc.TLS, self uintptr) {
	var i uint16_t
	_ = i
	i = uint16(0)
	for {
		if !(libc.Int32FromUint16(i) < libc.Int32FromUint16(uint16((*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize))) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(i) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(433), uintptr(unsafe.Pointer(&__func__26)))
			}
		}
		(*CaptureList)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(i)*16)).Fsize = libc.Uint32FromUint32(4294967295)
		goto _1
	_1:
		;
		i = i + 1
	}
	(*CaptureListPool)(unsafe.Pointer(self)).Ffree_capture_list_count = (*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize
}

var __func__26 = [24]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'l', 'i', 's', 't', '_', 'p', 'o', 'o', 'l', '_', 'r', 'e', 's', 'e', 't'}

func capture_list_pool_delete(tls *libc.TLS, self uintptr) {
	var i uint16_t
	_ = i
	i = uint16(0)
	for {
		if !(libc.Int32FromUint16(i) < libc.Int32FromUint16(uint16((*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize))) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(i) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(440), uintptr(unsafe.Pointer(&__func__27)))
			}
		}
		if (*CaptureList)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents+uintptr(i)*16)).Fcontents != 0 {
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32(i) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(440), uintptr(unsafe.Pointer(&__func__27)))
				}
			}
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*CaptureList)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fcontents+uintptr(i)*16)).Fcontents)
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(i) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(440), uintptr(unsafe.Pointer(&__func__27)))
			}
		}
		(*CaptureList)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(i)*16)).Fcontents = libc.UintptrFromInt32(0)
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(i) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(440), uintptr(unsafe.Pointer(&__func__27)))
			}
		}
		(*CaptureList)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(i)*16)).Fsize = uint32(0)
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(i) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(440), uintptr(unsafe.Pointer(&__func__27)))
			}
		}
		(*CaptureList)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(i)*16)).Fcapacity = uint32(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcapacity = uint32(0)
}

var __func__27 = [25]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'l', 'i', 's', 't', '_', 'p', 'o', 'o', 'l', '_', 'd', 'e', 'l', 'e', 't', 'e'}

func capture_list_pool_get(tls *libc.TLS, self uintptr, id uint16_t) (r uintptr) {
	if uint32(id) >= (*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize {
		return self + 16
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(id) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3112, __ccgo_ts+3070, int32(447), uintptr(unsafe.Pointer(&__func__28)))
		}
	}
	return (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(id)*16
}

var __func__28 = [22]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'l', 'i', 's', 't', '_', 'p', 'o', 'o', 'l', '_', 'g', 'e', 't'}

func capture_list_pool_get_mut(tls *libc.TLS, self uintptr, id uint16_t) (r uintptr) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(id) < (*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3149, __ccgo_ts+3070, int32(451), uintptr(unsafe.Pointer(&__func__29)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(id) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3112, __ccgo_ts+3070, int32(452), uintptr(unsafe.Pointer(&__func__29)))
		}
	}
	return (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(id)*16
}

var __func__29 = [26]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'l', 'i', 's', 't', '_', 'p', 'o', 'o', 'l', '_', 'g', 'e', 't', '_', 'm', 'u', 't'}

func capture_list_pool_is_empty(tls *libc.TLS, self uintptr) (r uint8) {
	return libc.BoolUint8((*CaptureListPool)(unsafe.Pointer(self)).Ffree_capture_list_count == uint32(0) && (*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize >= (*CaptureListPool)(unsafe.Pointer(self)).Fmax_capture_list_count)
}

func capture_list_pool_acquire(tls *libc.TLS, self uintptr) (r uint16_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i uint16_t
	var i1, v2 uint32_t
	var v3 uintptr
	var _ /* list at bp+0 */ CaptureList
	_, _, _, _ = i, i1, v2, v3
	if (*CaptureListPool)(unsafe.Pointer(self)).Ffree_capture_list_count > uint32(0) {
		i = uint16(0)
		for {
			if !(libc.Int32FromUint16(i) < libc.Int32FromUint16(uint16((*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize))) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32(i) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(465), uintptr(unsafe.Pointer(&__func__30)))
				}
			}
			if (*CaptureList)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fcontents+uintptr(i)*16)).Fsize == uint32(4294967295) {
				_ = libc.Uint64FromInt64(4)
				{
					if !(uint32(i) < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+3034, __ccgo_ts+3070, int32(466), uintptr(unsafe.Pointer(&__func__30)))
					}
				}
				(*CaptureList)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self)).Fcontents + uintptr(i)*16)).Fsize = uint32(0)
				(*CaptureListPool)(unsafe.Pointer(self)).Ffree_capture_list_count = (*CaptureListPool)(unsafe.Pointer(self)).Ffree_capture_list_count - 1
				return i
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	i1 = (*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize
	if i1 >= (*CaptureListPool)(unsafe.Pointer(self)).Fmax_capture_list_count {
		return NONE
	}
	(*CaptureList)(unsafe.Pointer(bp)).Fsize = uint32(0)
	(*CaptureList)(unsafe.Pointer(bp)).Fcapacity = uint32(0)
	(*CaptureList)(unsafe.Pointer(bp)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(16))
	v3 = self + 8
	v2 = *(*uint32_t)(unsafe.Pointer(v3))
	*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
	*(*CaptureList)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v2)*16)) = *(*CaptureList)(unsafe.Pointer(bp))
	return uint16(i1)
}

var __func__30 = [26]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'l', 'i', 's', 't', '_', 'p', 'o', 'o', 'l', '_', 'a', 'c', 'q', 'u', 'i', 'r', 'e'}

func capture_list_pool_release(tls *libc.TLS, self uintptr, id uint16_t) {
	if uint32(id) >= (*CaptureListPool)(unsafe.Pointer(self)).Flist.Fsize {
		return
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(id) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3112, __ccgo_ts+3070, int32(487), uintptr(unsafe.Pointer(&__func__31)))
		}
	}
	(*CaptureList)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(id)*16)).Fsize = libc.Uint32FromUint32(4294967295)
	(*CaptureListPool)(unsafe.Pointer(self)).Ffree_capture_list_count = (*CaptureListPool)(unsafe.Pointer(self)).Ffree_capture_list_count + 1
}

var __func__31 = [26]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'l', 'i', 's', 't', '_', 'p', 'o', 'o', 'l', '_', 'r', 'e', 'l', 'e', 'a', 's', 'e'}

func quantifier_mul(tls *libc.TLS, left TSQuantifier1, right TSQuantifier1) (r TSQuantifier1) {
	switch left {
	case int32(TSQuantifierZero):
		return int32(TSQuantifierZero)
	case int32(TSQuantifierZeroOrOne):
		switch right {
		case int32(TSQuantifierZero):
			return int32(TSQuantifierZero)
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierOne):
			return int32(TSQuantifierZeroOrOne)
		case int32(TSQuantifierZeroOrMore):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierZeroOrMore)
		}
	case int32(TSQuantifierZeroOrMore):
		switch right {
		case int32(TSQuantifierZero):
			return int32(TSQuantifierZero)
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierZeroOrMore):
			fallthrough
		case int32(TSQuantifierOne):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierZeroOrMore)
		}
	case int32(TSQuantifierOne):
		return right
	case int32(TSQuantifierOneOrMore):
		switch right {
		case int32(TSQuantifierZero):
			return int32(TSQuantifierZero)
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierZeroOrMore):
			return int32(TSQuantifierZeroOrMore)
		case int32(TSQuantifierOne):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierOneOrMore)
		}
		break
	}
	return int32(TSQuantifierZero)
}

func quantifier_join(tls *libc.TLS, left TSQuantifier1, right TSQuantifier1) (r TSQuantifier1) {
	switch left {
	case int32(TSQuantifierZero):
		switch right {
		case int32(TSQuantifierZero):
			return int32(TSQuantifierZero)
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierOne):
			return int32(TSQuantifierZeroOrOne)
		case int32(TSQuantifierZeroOrMore):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierZeroOrMore)
		}
	case int32(TSQuantifierZeroOrOne):
		switch right {
		case int32(TSQuantifierZero):
			fallthrough
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierOne):
			return int32(TSQuantifierZeroOrOne)
		case int32(TSQuantifierZeroOrMore):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierZeroOrMore)
			break
		}
	case int32(TSQuantifierZeroOrMore):
		return int32(TSQuantifierZeroOrMore)
	case int32(TSQuantifierOne):
		switch right {
		case int32(TSQuantifierZero):
			fallthrough
		case int32(TSQuantifierZeroOrOne):
			return int32(TSQuantifierZeroOrOne)
		case int32(TSQuantifierZeroOrMore):
			return int32(TSQuantifierZeroOrMore)
		case int32(TSQuantifierOne):
			return int32(TSQuantifierOne)
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierOneOrMore)
		}
	case int32(TSQuantifierOneOrMore):
		switch right {
		case int32(TSQuantifierZero):
			fallthrough
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierZeroOrMore):
			return int32(TSQuantifierZeroOrMore)
		case int32(TSQuantifierOne):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierOneOrMore)
		}
		break
	}
	return int32(TSQuantifierZero)
}

func quantifier_add(tls *libc.TLS, left TSQuantifier1, right TSQuantifier1) (r TSQuantifier1) {
	switch left {
	case int32(TSQuantifierZero):
		return right
	case int32(TSQuantifierZeroOrOne):
		switch right {
		case int32(TSQuantifierZero):
			return int32(TSQuantifierZeroOrOne)
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierZeroOrMore):
			return int32(TSQuantifierZeroOrMore)
		case int32(TSQuantifierOne):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierOneOrMore)
		}
	case int32(TSQuantifierZeroOrMore):
		switch right {
		case int32(TSQuantifierZero):
			return int32(TSQuantifierZeroOrMore)
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierZeroOrMore):
			return int32(TSQuantifierZeroOrMore)
		case int32(TSQuantifierOne):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierOneOrMore)
		}
	case int32(TSQuantifierOne):
		switch right {
		case int32(TSQuantifierZero):
			return int32(TSQuantifierOne)
		case int32(TSQuantifierZeroOrOne):
			fallthrough
		case int32(TSQuantifierZeroOrMore):
			fallthrough
		case int32(TSQuantifierOne):
			fallthrough
		case int32(TSQuantifierOneOrMore):
			return int32(TSQuantifierOneOrMore)
		}
	case int32(TSQuantifierOneOrMore):
		return int32(TSQuantifierOneOrMore)
	}
	return int32(TSQuantifierZero)
}

func capture_quantifiers_new(tls *libc.TLS) (r CaptureQuantifiers) {
	return CaptureQuantifiers{}
}

func capture_quantifiers_delete(tls *libc.TLS, self uintptr) {
	if (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents)
	}
	(*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize = uint32(0)
	(*CaptureQuantifiers)(unsafe.Pointer(self)).Fcapacity = uint32(0)
}

func capture_quantifiers_clear(tls *libc.TLS, self uintptr) {
	(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize = uint32(0)
}

func capture_quantifiers_replace(tls *libc.TLS, self uintptr, quantifiers uintptr) {
	(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize = uint32(0)
	(*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents = _array__splice(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents, self+8, self+12, libc.Uint64FromInt64(1), (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize, uint32(0), (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize, (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fcontents)
}

func capture_quantifier_for_id(tls *libc.TLS, self uintptr, id uint16_t) (r TSQuantifier1) {
	var v1 int32
	_ = v1
	if (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize <= uint32(id) {
		v1 = int32(TSQuantifierZero)
	} else {
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(id) < (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3170, __ccgo_ts+3070, int32(687), uintptr(unsafe.Pointer(&__func__32)))
			}
		}
		v1 = libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer((*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents + uintptr(id))))
	}
	return v1
}

var __func__32 = [26]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'q', 'u', 'a', 'n', 't', 'i', 'f', 'i', 'e', 'r', '_', 'f', 'o', 'r', '_', 'i', 'd'}

func capture_quantifiers_add_for_id(tls *libc.TLS, self uintptr, id uint16_t, quantifier TSQuantifier1) {
	var own_quantifier uintptr
	_ = own_quantifier
	if (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize <= uint32(id) {
		if libc.Uint32FromInt32(libc.Int32FromUint16(id)+int32(1))-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize == uint32(0) {
			goto _1
		}
		(*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize, self+12, libc.Uint32FromInt32(libc.Int32FromUint16(id)+int32(1))-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize, libc.Uint64FromInt64(1))
		libc.Xmemset(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents+uintptr((*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize), 0, uint64(libc.Uint32FromInt32(libc.Int32FromUint16(id)+libc.Int32FromInt32(1))-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize)*libc.Uint64FromInt64(1))
		*(*uint32_t)(unsafe.Pointer(self + 8)) += libc.Uint32FromInt32(libc.Int32FromUint16(id)+int32(1)) - (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize
		goto _2
	_2:
		;
		goto _1
	_1: /**/
		//
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(id) < (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3170, __ccgo_ts+3070, int32(699), uintptr(unsafe.Pointer(&__func__33)))
		}
	}
	own_quantifier = (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents + uintptr(id)
	*(*uint8_t)(unsafe.Pointer(own_quantifier)) = libc.Uint8FromInt32(quantifier_add(tls, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(own_quantifier))), quantifier))
}

var __func__33 = [31]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'q', 'u', 'a', 'n', 't', 'i', 'f', 'i', 'e', 'r', 's', '_', 'a', 'd', 'd', '_', 'f', 'o', 'r', '_', 'i', 'd'}

func capture_quantifiers_add_all(tls *libc.TLS, self uintptr, quantifiers uintptr) {
	var id uint16_t
	var own_quantifier, quantifier uintptr
	_, _, _ = id, own_quantifier, quantifier
	if (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize < (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize {
		if (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize == uint32(0) {
			goto _1
		}
		(*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize, self+12, (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize, libc.Uint64FromInt64(1))
		libc.Xmemset(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents+uintptr((*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize), 0, uint64((*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize)*libc.Uint64FromInt64(1))
		*(*uint32_t)(unsafe.Pointer(self + 8)) += (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize - (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize
		goto _2
	_2:
		;
		goto _1
	_1: /**/
		//
	}
	id = uint16(0)
	for {
		if !(libc.Int32FromUint16(id) < libc.Int32FromUint16(uint16((*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize))) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(id) < (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3200, __ccgo_ts+3070, int32(712), uintptr(unsafe.Pointer(&__func__34)))
			}
		}
		quantifier = (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fcontents + uintptr(id)
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(id) < (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3170, __ccgo_ts+3070, int32(713), uintptr(unsafe.Pointer(&__func__34)))
			}
		}
		own_quantifier = (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents + uintptr(id)
		*(*uint8_t)(unsafe.Pointer(own_quantifier)) = libc.Uint8FromInt32(quantifier_add(tls, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(own_quantifier))), libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(quantifier)))))
		goto _3
	_3:
		;
		id = id + 1
	}
}

var __func__34 = [28]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'q', 'u', 'a', 'n', 't', 'i', 'f', 'i', 'e', 'r', 's', '_', 'a', 'd', 'd', '_', 'a', 'l', 'l'}

func capture_quantifiers_mul(tls *libc.TLS, self uintptr, quantifier TSQuantifier1) {
	var id uint16_t
	var own_quantifier uintptr
	_, _ = id, own_quantifier
	id = uint16(0)
	for {
		if !(libc.Int32FromUint16(id) < libc.Int32FromUint16(uint16((*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize))) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(id) < (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3170, __ccgo_ts+3070, int32(724), uintptr(unsafe.Pointer(&__func__35)))
			}
		}
		own_quantifier = (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents + uintptr(id)
		*(*uint8_t)(unsafe.Pointer(own_quantifier)) = libc.Uint8FromInt32(quantifier_mul(tls, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(own_quantifier))), quantifier))
		goto _1
	_1:
		;
		id = id + 1
	}
}

var __func__35 = [24]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'q', 'u', 'a', 'n', 't', 'i', 'f', 'i', 'e', 'r', 's', '_', 'm', 'u', 'l'}

func capture_quantifiers_join_all(tls *libc.TLS, self uintptr, quantifiers uintptr) {
	var id, id1 uint32_t
	var own_quantifier, own_quantifier1, quantifier uintptr
	_, _, _, _, _ = id, id1, own_quantifier, own_quantifier1, quantifier
	if (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize < (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize {
		if (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize == uint32(0) {
			goto _1
		}
		(*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize, self+12, (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize, libc.Uint64FromInt64(1))
		libc.Xmemset(tls, (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents+uintptr((*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize), 0, uint64((*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize-(*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize)*libc.Uint64FromInt64(1))
		*(*uint32_t)(unsafe.Pointer(self + 8)) += (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize - (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize
		goto _2
	_2:
		;
		goto _1
	_1: /**/
		//
	}
	id = uint32(0)
	for {
		if !(id < (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(id < (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3200, __ccgo_ts+3070, int32(738), uintptr(unsafe.Pointer(&__func__36)))
			}
		}
		quantifier = (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fcontents + uintptr(id)
		_ = libc.Uint64FromInt64(4)
		{
			if !(id < (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3170, __ccgo_ts+3070, int32(739), uintptr(unsafe.Pointer(&__func__36)))
			}
		}
		own_quantifier = (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents + uintptr(id)
		*(*uint8_t)(unsafe.Pointer(own_quantifier)) = libc.Uint8FromInt32(quantifier_join(tls, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(own_quantifier))), libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(quantifier)))))
		goto _3
	_3:
		;
		id = id + 1
	}
	id1 = (*CaptureQuantifiers)(unsafe.Pointer(quantifiers)).Fsize
	for {
		if !(id1 < (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(id1 < (*CaptureQuantifiers)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3170, __ccgo_ts+3070, int32(743), uintptr(unsafe.Pointer(&__func__36)))
			}
		}
		own_quantifier1 = (*CaptureQuantifiers)(unsafe.Pointer(self)).Fcontents + uintptr(id1)
		*(*uint8_t)(unsafe.Pointer(own_quantifier1)) = libc.Uint8FromInt32(quantifier_join(tls, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(own_quantifier1))), int32(TSQuantifierZero)))
		goto _4
	_4:
		;
		id1 = id1 + 1
	}
}

var __func__36 = [29]int8{'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'q', 'u', 'a', 'n', 't', 'i', 'f', 'i', 'e', 'r', 's', '_', 'j', 'o', 'i', 'n', '_', 'a', 'l', 'l'}

func symbol_table_new(tls *libc.TLS) (r SymbolTable) {
	return SymbolTable{}
}

func symbol_table_delete(tls *libc.TLS, self uintptr) {
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcapacity = uint32(0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+16)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+16)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 16)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 16)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 16)).Fcapacity = uint32(0)
}

func symbol_table_id_for_name(tls *libc.TLS, self uintptr, name uintptr, length uint32_t) (r int32) {
	var i uint32
	var slice Slice
	var v2 bool
	_, _, _ = i, slice, v2
	i = uint32(0)
	for {
		if !(i < (*SymbolTable)(unsafe.Pointer(self)).Fslices.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+16)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3237, __ccgo_ts+3070, int32(770), uintptr(unsafe.Pointer(&__func__37)))
			}
		}
		slice = *(*Slice)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+16)).Fcontents + uintptr(i)*8))
		if v2 = slice.Flength == length; v2 {
			_ = libc.Uint64FromInt64(4)
			{
				if !(slice.Foffset < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+3275, __ccgo_ts+3070, int32(773), uintptr(unsafe.Pointer(&__func__37)))
				}
			}
		}
		if v2 && !(libc.Xstrncmp(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents+uintptr(slice.Foffset), name, uint64(length)) != 0) {
			return libc.Int32FromUint32(i)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return -int32(1)
}

var __func__37 = [25]int8{'s', 'y', 'm', 'b', 'o', 'l', '_', 't', 'a', 'b', 'l', 'e', '_', 'i', 'd', '_', 'f', 'o', 'r', '_', 'n', 'a', 'm', 'e'}

func symbol_table_name_for_id(tls *libc.TLS, self uintptr, id uint16_t, length uintptr) (r uintptr) {
	var slice Slice
	_ = slice
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(id) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+16)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3328, __ccgo_ts+3070, int32(784), uintptr(unsafe.Pointer(&__func__38)))
		}
	}
	slice = *(*Slice)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+16)).Fcontents + uintptr(id)*8))
	*(*uint32_t)(unsafe.Pointer(length)) = slice.Flength
	_ = libc.Uint64FromInt64(4)
	{
		if !(slice.Foffset < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3275, __ccgo_ts+3070, int32(786), uintptr(unsafe.Pointer(&__func__38)))
		}
	}
	return (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(slice.Foffset)
}

var __func__38 = [25]int8{'s', 'y', 'm', 'b', 'o', 'l', '_', 't', 'a', 'b', 'l', 'e', '_', 'n', 'a', 'm', 'e', '_', 'f', 'o', 'r', '_', 'i', 'd'}

func symbol_table_insert_name(tls *libc.TLS, self uintptr, name uintptr, length uint32_t) (r uint16_t) {
	var id int32
	var slice Slice
	var v3 uint32_t
	var v4 uintptr
	_, _, _, _ = id, slice, v3, v4
	id = symbol_table_id_for_name(tls, self, name, length)
	if id >= 0 {
		return libc.Uint16FromInt32(id)
	}
	slice = Slice{
		Foffset: (*SymbolTable)(unsafe.Pointer(self)).Fcharacters.Fsize,
		Flength: length,
	}
	if length+uint32(1) == uint32(0) {
		goto _1
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize, self+12, length+uint32(1), libc.Uint64FromInt64(1))
	libc.Xmemset(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize), 0, uint64(length+libc.Uint32FromInt32(1))*libc.Uint64FromInt64(1))
	*(*uint32_t)(unsafe.Pointer(self + 8)) += length + uint32(1)
	goto _2
_2:
	;
	goto _1
_1: /**/
	; //
	_ = libc.Uint64FromInt64(4)
	{
		if !(slice.Foffset < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3275, __ccgo_ts+3070, int32(801), uintptr(unsafe.Pointer(&__func__39)))
		}
	}
	libc.Xmemcpy(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents+uintptr(slice.Foffset), name, uint64(length))
	_ = libc.Uint64FromInt64(4)
	{
		if !((*SymbolTable)(unsafe.Pointer(self)).Fcharacters.Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3367, __ccgo_ts+3070, int32(802), uintptr(unsafe.Pointer(&__func__39)))
		}
	}
	*(*int8)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr((*SymbolTable)(unsafe.Pointer(self)).Fcharacters.Fsize-uint32(1)))) = 0
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 16)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+16)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+16)).Fsize, self+16+12, uint32(1), libc.Uint64FromInt64(8))
	v4 = self + 16 + 8
	v3 = *(*uint32_t)(unsafe.Pointer(v4))
	*(*uint32_t)(unsafe.Pointer(v4)) = *(*uint32_t)(unsafe.Pointer(v4)) + 1
	*(*Slice)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+16)).Fcontents + uintptr(v3)*8)) = slice
	return uint16((*SymbolTable)(unsafe.Pointer(self)).Fslices.Fsize - uint32(1))
}

var __func__39 = [25]int8{'s', 'y', 'm', 'b', 'o', 'l', '_', 't', 'a', 'b', 'l', 'e', '_', 'i', 'n', 's', 'e', 'r', 't', '_', 'n', 'a', 'm', 'e'}

func query_step__new(tls *libc.TLS, symbol TSSymbol, depth uint16_t, is_immediate uint8) (r QueryStep) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i uint32
	var _ /* step at bp+0 */ QueryStep
	_ = i
	*(*QueryStep)(unsafe.Pointer(bp)) = QueryStep{
		Fsymbol:            symbol,
		Fdepth:             depth,
		Falternative_index: NONE,
		F__ccgo18:          uint8(0)&0x1<<0 | is_immediate&0x1<<1 | uint8(0)&0x1<<2 | uint8(0)&0x1<<3 | uint8(0)&0x1<<4 | uint8(0)&0x1<<5 | uint8(0)&0x1<<6 | uint8(0)&0x1<<7,
	}
	i = uint32(0)
	for {
		if !(i < uint32(3)) {
			break
		}
		*(*uint16_t)(unsafe.Pointer(bp + 6 + uintptr(i)*2)) = NONE
		goto _1
	_1:
		;
		i = i + 1
	}
	return *(*QueryStep)(unsafe.Pointer(bp))
}

func query_step__add_capture(tls *libc.TLS, self uintptr, capture_id uint16_t) {
	var i uint32
	_ = i
	i = uint32(0)
	for {
		if !(i < uint32(3)) {
			break
		}
		if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i)*2))) == libc.Int32FromUint16(NONE) {
			*(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i)*2)) = capture_id
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func query_step__remove_capture(tls *libc.TLS, self uintptr, capture_id uint16_t) {
	var i uint32
	_ = i
	i = uint32(0)
	for {
		if !(i < uint32(3)) {
			break
		}
		if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i)*2))) == libc.Int32FromUint16(capture_id) {
			*(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i)*2)) = NONE
			for i+uint32(1) < uint32(3) {
				if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i+uint32(1))*2))) == libc.Int32FromUint16(NONE) {
					break
				}
				*(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i)*2)) = *(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i+uint32(1))*2))
				*(*uint16_t)(unsafe.Pointer(self + 6 + uintptr(i+uint32(1))*2)) = NONE
				i = i + 1
			}
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func state_predecessor_map_new(tls *libc.TLS, language uintptr) (r StatePredecessorMap) {
	return StatePredecessorMap{
		Fcontents: (*(*func(*libc.TLS, size_t, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_calloc})))(tls, uint64((*TSLanguage)(unsafe.Pointer(language)).Fstate_count)*libc.Uint64FromInt32(libc.Int32FromInt32(256)+libc.Int32FromInt32(1)), uint64(2)),
	}
}

func state_predecessor_map_delete(tls *libc.TLS, self uintptr) {
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*StatePredecessorMap)(unsafe.Pointer(self)).Fcontents)
}

func state_predecessor_map_add(tls *libc.TLS, self uintptr, state TSStateId, predecessor TSStateId) {
	var count uintptr
	var index size_t
	_, _ = count, index
	index = uint64(state) * libc.Uint64FromInt32(libc.Int32FromInt32(256)+libc.Int32FromInt32(1))
	count = (*StatePredecessorMap)(unsafe.Pointer(self)).Fcontents + uintptr(index)*2
	if libc.Int32FromUint16(*(*TSStateId)(unsafe.Pointer(count))) == 0 || libc.Int32FromUint16(*(*TSStateId)(unsafe.Pointer(count))) < int32(256) && libc.Int32FromUint16(*(*TSStateId)(unsafe.Pointer((*StatePredecessorMap)(unsafe.Pointer(self)).Fcontents + uintptr(index+uint64(*(*TSStateId)(unsafe.Pointer(count))))*2))) != libc.Int32FromUint16(predecessor) {
		*(*TSStateId)(unsafe.Pointer(count)) = *(*TSStateId)(unsafe.Pointer(count)) + 1
		*(*TSStateId)(unsafe.Pointer((*StatePredecessorMap)(unsafe.Pointer(self)).Fcontents + uintptr(index+uint64(*(*TSStateId)(unsafe.Pointer(count))))*2)) = predecessor
	}
}

func state_predecessor_map_get(tls *libc.TLS, self uintptr, state TSStateId, count uintptr) (r uintptr) {
	var index size_t
	_ = index
	index = uint64(state) * libc.Uint64FromInt32(libc.Int32FromInt32(256)+libc.Int32FromInt32(1))
	*(*uint32)(unsafe.Pointer(count)) = uint32(*(*TSStateId)(unsafe.Pointer((*StatePredecessorMap)(unsafe.Pointer(self)).Fcontents + uintptr(index)*2)))
	return (*StatePredecessorMap)(unsafe.Pointer(self)).Fcontents + uintptr(index+uint64(1))*2
}

func analysis_state__recursion_depth(tls *libc.TLS, self uintptr) (r uint32) {
	var i, j, result uint32
	var symbol TSSymbol
	_, _, _, _ = i, j, result, symbol
	result = uint32(0)
	i = uint32(0)
	for {
		if !(i < uint32((*AnalysisState)(unsafe.Pointer(self)).Fdepth)) {
			break
		}
		symbol = (*(*AnalysisStateEntry)(unsafe.Pointer(self + uintptr(i)*8))).Fparent_symbol
		j = uint32(0)
		for {
			if !(j < i) {
				break
			}
			if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(self + uintptr(j)*8))).Fparent_symbol) == libc.Int32FromUint16(symbol) {
				result = result + 1
				break
			}
			goto _2
		_2:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return result
}

func analysis_state__compare(tls *libc.TLS, self uintptr, other uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i uint32
	var _ /* s1 at bp+0 */ AnalysisStateEntry
	var _ /* s2 at bp+8 */ AnalysisStateEntry
	_ = i
	if libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fdepth) < libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(other)))).Fdepth) {
		return int32(1)
	}
	i = uint32(0)
	for {
		if !(i < uint32((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fdepth)) {
			break
		}
		if i >= uint32((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(other)))).Fdepth) {
			return -int32(1)
		}
		*(*AnalysisStateEntry)(unsafe.Pointer(bp)) = *(*AnalysisStateEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)) + uintptr(i)*8))
		*(*AnalysisStateEntry)(unsafe.Pointer(bp + 8)) = *(*AnalysisStateEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(other)) + uintptr(i)*8))
		if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp))).Fchild_index) < libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp + 8))).Fchild_index) {
			return -int32(1)
		}
		if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp))).Fchild_index) > libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp + 8))).Fchild_index) {
			return int32(1)
		}
		if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp))).Fparent_symbol) < libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp + 8))).Fparent_symbol) {
			return -int32(1)
		}
		if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp))).Fparent_symbol) > libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp + 8))).Fparent_symbol) {
			return int32(1)
		}
		if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp))).Fparse_state) < libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp + 8))).Fparse_state) {
			return -int32(1)
		}
		if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp))).Fparse_state) > libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(bp + 8))).Fparse_state) {
			return int32(1)
		}
		if int32(*(*uint16)(unsafe.Pointer(bp + 6))&0x7fff>>0) < int32(*(*uint16)(unsafe.Pointer(bp + 8 + 6))&0x7fff>>0) {
			return -int32(1)
		}
		if int32(*(*uint16)(unsafe.Pointer(bp + 6))&0x7fff>>0) > int32(*(*uint16)(unsafe.Pointer(bp + 8 + 6))&0x7fff>>0) {
			return int32(1)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fstep_index) < libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(other)))).Fstep_index) {
		return -int32(1)
	}
	if libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fstep_index) > libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(other)))).Fstep_index) {
		return int32(1)
	}
	return 0
}

func analysis_state__top(tls *libc.TLS, self uintptr) (r uintptr) {
	if libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(self)).Fdepth) == 0 {
		return self
	}
	return self + uintptr(libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(self)).Fdepth)-int32(1))*8
}

func analysis_state__has_supertype(tls *libc.TLS, self uintptr, symbol TSSymbol) (r uint8) {
	var i uint32
	_ = i
	i = uint32(0)
	for {
		if !(i < uint32((*AnalysisState)(unsafe.Pointer(self)).Fdepth)) {
			break
		}
		if libc.Int32FromUint16((*(*AnalysisStateEntry)(unsafe.Pointer(self + uintptr(i)*8))).Fparent_symbol) == libc.Int32FromUint16(symbol) {
			return libc.BoolUint8(1 != 0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(0 != 0)
}

func analysis_state_pool__clone_or_reuse(tls *libc.TLS, self uintptr, borrowed_item uintptr) (r uintptr) {
	var new_item, v2 uintptr
	var v1 uint32_t
	_, _, _ = new_item, v1, v2
	if (*AnalysisStateSet)(unsafe.Pointer(self)).Fsize != 0 {
		v2 = self + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		new_item = *(*uintptr)(unsafe.Pointer((*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents + uintptr(v1)*8))
	} else {
		new_item = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(70))
	}
	*(*AnalysisState)(unsafe.Pointer(new_item)) = *(*AnalysisState)(unsafe.Pointer(borrowed_item))
	return new_item
}

func analysis_state_set__insert_sorted(tls *libc.TLS, self uintptr, pool uintptr, _borrowed_item uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*uintptr)(unsafe.Pointer(bp)) = _borrowed_item
	var comparison int32
	var half_size, mid_index, size uint32_t
	var _ /* exists at bp+12 */ uint32
	var _ /* index at bp+8 */ uint32
	var _ /* new_item at bp+16 */ uintptr
	_, _, _, _ = comparison, half_size, mid_index, size
	*(*uint32)(unsafe.Pointer(bp + 8)) = uint32(0)
	*(*uint32)(unsafe.Pointer(bp + 12)) = uint32(0)
	size = (*AnalysisStateSet)(unsafe.Pointer(self)).Fsize - *(*uint32)(unsafe.Pointer(bp + 8))
	if size == uint32(0) {
		goto _1
	}
	for size > uint32(1) {
		half_size = size / uint32(2)
		mid_index = *(*uint32)(unsafe.Pointer(bp + 8)) + half_size
		comparison = analysis_state__compare(tls, (*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents+uintptr(mid_index)*8, bp)
		if comparison <= 0 {
			*(*uint32)(unsafe.Pointer(bp + 8)) = mid_index
		}
		size = size - half_size
	}
	comparison = analysis_state__compare(tls, (*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents+uintptr(*(*uint32)(unsafe.Pointer(bp + 8)))*8, bp)
	if comparison == 0 {
		*(*uint32)(unsafe.Pointer(bp + 12)) = uint32(1)
	} else {
		if comparison < 0 {
			*(*uint32)(unsafe.Pointer(bp + 8)) += uint32(1)
		}
	}
	goto _2
_2:
	;
	goto _1
_1: /**/
	; //
	if !(*(*uint32)(unsafe.Pointer(bp + 12)) != 0) {
		*(*uintptr)(unsafe.Pointer(bp + 16)) = analysis_state_pool__clone_or_reuse(tls, pool, *(*uintptr)(unsafe.Pointer(bp)))
		(*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents = _array__splice(tls, (*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents, self+8, self+12, libc.Uint64FromInt64(8), *(*uint32)(unsafe.Pointer(bp + 8)), uint32(0), uint32(1), bp+16)
	}
}

func analysis_state_set__push(tls *libc.TLS, self uintptr, pool uintptr, borrowed_item uintptr) {
	var new_item, v2 uintptr
	var v1 uint32_t
	_, _, _ = new_item, v1, v2
	new_item = analysis_state_pool__clone_or_reuse(tls, pool, borrowed_item)
	(*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents, (*AnalysisStateSet)(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(8))
	v2 = self + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*uintptr)(unsafe.Pointer((*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents + uintptr(v1)*8)) = new_item
}

func analysis_state_set__clear(tls *libc.TLS, self uintptr, pool uintptr) {
	(*AnalysisStateSet)(unsafe.Pointer(pool)).Fcontents = _array__splice(tls, (*AnalysisStateSet)(unsafe.Pointer(pool)).Fcontents, pool+8, pool+12, libc.Uint64FromInt64(8), (*AnalysisStateSet)(unsafe.Pointer(pool)).Fsize, uint32(0), (*AnalysisStateSet)(unsafe.Pointer(self)).Fsize, (*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents)
	(*AnalysisStateSet)(unsafe.Pointer(self)).Fsize = uint32(0)
}

func analysis_state_set__delete(tls *libc.TLS, self uintptr) {
	var i uint32
	_ = i
	i = uint32(0)
	for {
		if !(i < (*AnalysisStateSet)(unsafe.Pointer(self)).Fsize) {
			break
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, *(*uintptr)(unsafe.Pointer((*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	if (*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents)
	}
	(*AnalysisStateSet)(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*AnalysisStateSet)(unsafe.Pointer(self)).Fsize = uint32(0)
	(*AnalysisStateSet)(unsafe.Pointer(self)).Fcapacity = uint32(0)
}

func query_analysis__new(tls *libc.TLS) (r QueryAnalysis) {
	return QueryAnalysis{}
}

func query_analysis__delete(tls *libc.TLS, self uintptr) {
	analysis_state_set__delete(tls, self)
	analysis_state_set__delete(tls, self+16)
	analysis_state_set__delete(tls, self+32)
	analysis_state_set__delete(tls, self+48)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+64)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+64)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 64)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 64)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 64)).Fcapacity = uint32(0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+80)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 80)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 80)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 80)).Fcapacity = uint32(0)
}

func analysis_subgraph_node__compare(tls *libc.TLS, self uintptr, other uintptr) (r int32) {
	if libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(self)).Fstate) < libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(other)).Fstate) {
		return -int32(1)
	}
	if libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(self)).Fstate) > libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(other)).Fstate) {
		return int32(1)
	}
	if int32(*(*uint8)(unsafe.Pointer(self + 4))&0x7f>>0) < int32(*(*uint8)(unsafe.Pointer(other + 4))&0x7f>>0) {
		return -int32(1)
	}
	if int32(*(*uint8)(unsafe.Pointer(self + 4))&0x7f>>0) > int32(*(*uint8)(unsafe.Pointer(other + 4))&0x7f>>0) {
		return int32(1)
	}
	if int32(*(*uint8)(unsafe.Pointer(self + 4))&0x80>>7) < int32(*(*uint8)(unsafe.Pointer(other + 4))&0x80>>7) {
		return -int32(1)
	}
	if int32(*(*uint8)(unsafe.Pointer(self + 4))&0x80>>7) > int32(*(*uint8)(unsafe.Pointer(other + 4))&0x80>>7) {
		return int32(1)
	}
	if libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(self)).Fproduction_id) < libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(other)).Fproduction_id) {
		return -int32(1)
	}
	if libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(self)).Fproduction_id) > libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(other)).Fproduction_id) {
		return int32(1)
	}
	return 0
}

func ts_query__pattern_map_search(tls *libc.TLS, self uintptr, needle TSSymbol, result uintptr) (r uint8) {
	var base_index, half_size, mid_index, size uint32_t
	var mid_symbol, symbol TSSymbol
	_, _, _, _, _, _ = base_index, half_size, mid_index, mid_symbol, size, symbol
	base_index = uint32((*TSQuery)(unsafe.Pointer(self)).Fwildcard_root_pattern_count)
	size = (*TSQuery)(unsafe.Pointer(self)).Fpattern_map.Fsize - base_index
	if size == uint32(0) {
		*(*uint32_t)(unsafe.Pointer(result)) = base_index
		return libc.BoolUint8(0 != 0)
	}
	for size > uint32(1) {
		half_size = size / uint32(2)
		mid_index = base_index + half_size
		_ = libc.Uint64FromInt64(4)
		{
			_ = libc.Uint64FromInt64(4)
			{
				if !(mid_index < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+96)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+3433, __ccgo_ts+3070, int32(1106), uintptr(unsafe.Pointer(&__func__40)))
				}
			}
			if !(uint32((*PatternEntry)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fcontents+uintptr(mid_index)*6)).Fstep_index) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3484, __ccgo_ts+3070, int32(1105), uintptr(unsafe.Pointer(&__func__40)))
			}
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(mid_index < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3433, __ccgo_ts+3070, int32(1106), uintptr(unsafe.Pointer(&__func__40)))
			}
		}
		mid_symbol = (*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents+uintptr(mid_index)*6)).Fstep_index)*20)).Fsymbol
		if libc.Int32FromUint16(needle) > libc.Int32FromUint16(mid_symbol) {
			base_index = mid_index
		}
		size = size - half_size
	}
	_ = libc.Uint64FromInt64(4)
	{
		_ = libc.Uint64FromInt64(4)
		{
			if !(base_index < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3896, __ccgo_ts+3070, int32(1113), uintptr(unsafe.Pointer(&__func__40)))
			}
		}
		if !(uint32((*PatternEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents+uintptr(base_index)*6)).Fstep_index) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3948, __ccgo_ts+3070, int32(1112), uintptr(unsafe.Pointer(&__func__40)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(base_index < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+3896, __ccgo_ts+3070, int32(1113), uintptr(unsafe.Pointer(&__func__40)))
		}
	}
	symbol = (*QueryStep)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+96)).Fcontents+uintptr(base_index)*6)).Fstep_index)*20)).Fsymbol
	if libc.Int32FromUint16(needle) > libc.Int32FromUint16(symbol) {
		base_index = base_index + 1
		if base_index < (*TSQuery)(unsafe.Pointer(self)).Fpattern_map.Fsize {
			_ = libc.Uint64FromInt64(4)
			{
				_ = libc.Uint64FromInt64(4)
				{
					if !(base_index < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+96)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+3896, __ccgo_ts+3070, int32(1120), uintptr(unsafe.Pointer(&__func__40)))
					}
				}
				if !(uint32((*PatternEntry)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+96)).Fcontents+uintptr(base_index)*6)).Fstep_index) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+4364, __ccgo_ts+3070, int32(1119), uintptr(unsafe.Pointer(&__func__40)))
				}
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(base_index < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+96)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+3896, __ccgo_ts+3070, int32(1120), uintptr(unsafe.Pointer(&__func__40)))
				}
			}
			symbol = (*QueryStep)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fcontents+uintptr(base_index)*6)).Fstep_index)*20)).Fsymbol
		}
	}
	*(*uint32_t)(unsafe.Pointer(result)) = base_index
	return libc.BoolUint8(libc.Int32FromUint16(needle) == libc.Int32FromUint16(symbol))
}

var __func__40 = [29]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', '_', 'p', 'a', 't', 't', 'e', 'r', 'n', '_', 'm', 'a', 'p', '_', 's', 'e', 'a', 'r', 'c', 'h'}

func ts_query__pattern_map_insert(tls *libc.TLS, self uintptr, symbol TSSymbol, _new_entry PatternEntry) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*PatternEntry)(unsafe.Pointer(bp)) = _new_entry
	var entry uintptr
	var _ /* index at bp+8 */ uint32_t
	_ = entry
	ts_query__pattern_map_search(tls, self, symbol, bp+8)
	for *(*uint32_t)(unsafe.Pointer(bp + 8)) < (*TSQuery)(unsafe.Pointer(self)).Fpattern_map.Fsize {
		_ = libc.Uint64FromInt64(4)
		{
			if !(*(*uint32_t)(unsafe.Pointer(bp + 8)) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+4780, __ccgo_ts+3070, int32(1144), uintptr(unsafe.Pointer(&__func__41)))
			}
		}
		entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 8)))*6
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32((*PatternEntry)(unsafe.Pointer(entry)).Fstep_index) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+4827, __ccgo_ts+3070, int32(1146), uintptr(unsafe.Pointer(&__func__41)))
			}
		}
		if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents+uintptr((*PatternEntry)(unsafe.Pointer(entry)).Fstep_index)*20)).Fsymbol) == libc.Int32FromUint16(symbol) && libc.Int32FromUint16((*PatternEntry)(unsafe.Pointer(entry)).Fpattern_index) < libc.Int32FromUint16((*(*PatternEntry)(unsafe.Pointer(bp))).Fpattern_index) {
			*(*uint32_t)(unsafe.Pointer(bp + 8)) = *(*uint32_t)(unsafe.Pointer(bp + 8)) + 1
		} else {
			break
		}
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 96)).Fcontents = _array__splice(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+96)).Fcontents, self+96+8, self+96+12, libc.Uint64FromInt64(6), *(*uint32_t)(unsafe.Pointer(bp + 8)), uint32(0), uint32(1), bp)
}

var __func__41 = [29]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', '_', 'p', 'a', 't', 't', 'e', 'r', 'n', '_', 'm', 'a', 'p', '_', 'i', 'n', 's', 'e', 'r', 't'}

func ts_query__perform_analysis(tls *libc.TLS, self uintptr, subgraphs uintptr, analysis uintptr) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var _states, _states1 AnalysisStateSet
	var action, next_state_top, next_step, node, step, subgraph uintptr
	var alias, parent_symbol, sym, visible_symbol TSSymbol
	var child_index, iteration, j, prev_final_step_count, recursion_depth_limit uint32
	var comparison, comparison1, comparison2, comparison3, comparison4, v7, v8 int32
	var did_finish_pattern, does_match uint8
	var field_id, parent_field_id TSFieldId
	var half_size, half_size1, half_size2, half_size3, mid_index, mid_index1, mid_index2, mid_index3, size, size1, size2, size3 uint32_t
	var parse_state TSStateId
	var _ /* _exists at bp+180 */ uint32
	var _ /* _exists at bp+188 */ uint32
	var _ /* _index at bp+176 */ uint32
	var _ /* _index at bp+184 */ uint32
	var _ /* exists at bp+12 */ uint32
	var _ /* field_map at bp+88 */ uintptr
	var _ /* field_map_end at bp+96 */ uintptr
	var _ /* lookahead_iterator at bp+16 */ LookaheadIterator
	var _ /* next_state at bp+104 */ AnalysisState
	var _ /* node_index at bp+80 */ uint32
	var _ /* state at bp+0 */ uintptr
	var _ /* subgraph_index at bp+8 */ uint32
	var _ /* successor at bp+72 */ AnalysisSubgraphNode
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _states, _states1, action, alias, child_index, comparison, comparison1, comparison2, comparison3, comparison4, did_finish_pattern, does_match, field_id, half_size, half_size1, half_size2, half_size3, iteration, j, mid_index, mid_index1, mid_index2, mid_index3, next_state_top, next_step, node, parent_field_id, parent_symbol, parse_state, prev_final_step_count, recursion_depth_limit, size, size1, size2, size3, step, subgraph, sym, visible_symbol, v7, v8
	recursion_depth_limit = uint32(0)
	prev_final_step_count = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(analysis + 64)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(analysis + 80)).Fsize = uint32(0)
	iteration = uint32(0)
	for {
		if iteration == uint32(256) {
			(*QueryAnalysis)(unsafe.Pointer(analysis)).Fdid_abort = libc.BoolUint8(1 != 0)
			break
		}
		if (*QueryAnalysis)(unsafe.Pointer(analysis)).Fstates.Fsize == uint32(0) {
			if (*QueryAnalysis)(unsafe.Pointer(analysis)).Fdeeper_states.Fsize > uint32(0) && (*QueryAnalysis)(unsafe.Pointer(analysis)).Ffinal_step_indices.Fsize > prev_final_step_count {
				prev_final_step_count = (*QueryAnalysis)(unsafe.Pointer(analysis)).Ffinal_step_indices.Fsize
				recursion_depth_limit = recursion_depth_limit + 1
				_states = (*QueryAnalysis)(unsafe.Pointer(analysis)).Fstates
				(*QueryAnalysis)(unsafe.Pointer(analysis)).Fstates = (*QueryAnalysis)(unsafe.Pointer(analysis)).Fdeeper_states
				(*QueryAnalysis)(unsafe.Pointer(analysis)).Fdeeper_states = _states
				goto _1
			}
			break
		}
		analysis_state_set__clear(tls, analysis+16, analysis+48)
		j = uint32(0)
		for {
			if !(j < (*QueryAnalysis)(unsafe.Pointer(analysis)).Fstates.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j < (*AnalysisStateSet)(unsafe.Pointer(analysis)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+4880, __ccgo_ts+3070, int32(1226), uintptr(unsafe.Pointer(&__func__42)))
				}
			}
			*(*uintptr)(unsafe.Pointer(bp)) = *(*uintptr)(unsafe.Pointer((*AnalysisStateSet)(unsafe.Pointer(analysis)).Fcontents + uintptr(j)*8))
			if (*QueryAnalysis)(unsafe.Pointer(analysis)).Fnext_states.Fsize > uint32(0) {
				_ = libc.Uint64FromInt64(4)
				{
					if !((*AnalysisStateSet)(unsafe.Pointer(analysis+16)).Fsize-libc.Uint32FromInt32(1) < (*AnalysisStateSet)(unsafe.Pointer(analysis+16)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+4922, __ccgo_ts+3070, int32(1236), uintptr(unsafe.Pointer(&__func__42)))
					}
				}
				comparison = analysis_state__compare(tls, bp, (*AnalysisStateSet)(unsafe.Pointer(analysis+16)).Fcontents+uintptr((*AnalysisStateSet)(unsafe.Pointer(analysis+16)).Fsize-uint32(1))*8)
				if comparison == 0 {
					analysis_state_set__insert_sorted(tls, analysis+16, analysis+48, *(*uintptr)(unsafe.Pointer(bp)))
					goto _2
				} else {
					if comparison > 0 {
						for j < (*QueryAnalysis)(unsafe.Pointer(analysis)).Fstates.Fsize {
							_ = libc.Uint64FromInt64(4)
							{
								if !(j < (*AnalysisStateSet)(unsafe.Pointer(analysis)).Fsize) {
									libc.X__assert_fail(tls, __ccgo_ts+4880, __ccgo_ts+3070, int32(1249), uintptr(unsafe.Pointer(&__func__42)))
								}
							}
							analysis_state_set__push(tls, analysis+16, analysis+48, *(*uintptr)(unsafe.Pointer((*AnalysisStateSet)(unsafe.Pointer(analysis)).Fcontents + uintptr(j)*8)))
							j = j + 1
						}
						break
					}
				}
			}
			parse_state = (*AnalysisStateEntry)(unsafe.Pointer(analysis_state__top(tls, *(*uintptr)(unsafe.Pointer(bp))))).Fparse_state
			parent_symbol = (*AnalysisStateEntry)(unsafe.Pointer(analysis_state__top(tls, *(*uintptr)(unsafe.Pointer(bp))))).Fparent_symbol
			parent_field_id = libc.Uint16FromInt32(int32(*(*uint16)(unsafe.Pointer(analysis_state__top(tls, *(*uintptr)(unsafe.Pointer(bp))) + 6)) & 0x7fff >> 0))
			child_index = uint32((*AnalysisStateEntry)(unsafe.Pointer(analysis_state__top(tls, *(*uintptr)(unsafe.Pointer(bp))))).Fchild_index)
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fstep_index) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5002, __ccgo_ts+3070, int32(1261), uintptr(unsafe.Pointer(&__func__42)))
				}
			}
			step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fstep_index)*20
			*(*uint32)(unsafe.Pointer(bp + 8)) = uint32(0)
			*(*uint32)(unsafe.Pointer(bp + 12)) = uint32(0)
			size = (*AnalysisSubgraphArray)(unsafe.Pointer(subgraphs)).Fsize - *(*uint32)(unsafe.Pointer(bp + 8))
			if size == uint32(0) {
				goto _3
			}
			for size > uint32(1) {
				half_size = size / uint32(2)
				mid_index = *(*uint32)(unsafe.Pointer(bp + 8)) + half_size
				comparison1 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(subgraphs)).Fcontents + uintptr(mid_index)*40))) - libc.Int32FromUint16(parent_symbol)
				if comparison1 <= 0 {
					*(*uint32)(unsafe.Pointer(bp + 8)) = mid_index
				}
				size = size - half_size
			}
			comparison1 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(subgraphs)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 8)))*40))) - libc.Int32FromUint16(parent_symbol)
			if comparison1 == 0 {
				*(*uint32)(unsafe.Pointer(bp + 12)) = uint32(1)
			} else {
				if comparison1 < 0 {
					*(*uint32)(unsafe.Pointer(bp + 8)) += uint32(1)
				}
			}
			goto _4
		_4:
			;
			goto _3
		_3: /**/
			; //
			if !(*(*uint32)(unsafe.Pointer(bp + 12)) != 0) {
				goto _2
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(*(*uint32)(unsafe.Pointer(bp + 8)) < (*AnalysisSubgraphArray)(unsafe.Pointer(subgraphs)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5055, __ccgo_ts+3070, int32(1266), uintptr(unsafe.Pointer(&__func__42)))
				}
			}
			subgraph = (*AnalysisSubgraphArray)(unsafe.Pointer(subgraphs)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 8)))*40
			*(*LookaheadIterator)(unsafe.Pointer(bp + 16)) = ts_language_lookaheads(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, parse_state)
			for ts_lookahead_iterator__next(tls, bp+16) != 0 {
				sym = (*(*LookaheadIterator)(unsafe.Pointer(bp + 16))).Fsymbol
				*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 72)) = AnalysisSubgraphNode{
					Fstate:   parse_state,
					F__ccgo4: uint8(child_index) & 0x7f << 0,
				}
				if (*(*LookaheadIterator)(unsafe.Pointer(bp + 16))).Faction_count != 0 {
					action = (*(*LookaheadIterator)(unsafe.Pointer(bp + 16))).Factions + uintptr(libc.Int32FromUint16((*(*LookaheadIterator)(unsafe.Pointer(bp + 16))).Faction_count)-int32(1))*8
					if libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(action))) == int32(TSParseActionTypeShift) {
						if !((*TSParseAction)(unsafe.Pointer(action)).Fshift.Fextra != 0) {
							(*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 72))).Fstate = (*TSParseAction)(unsafe.Pointer(action)).Fshift.Fstate
							libc.PostIncBitFieldPtr8Uint8(bp+72+4, 1, 7, 0, 0x7f)
						}
					} else {
						continue
					}
				} else {
					if libc.Int32FromUint16((*(*LookaheadIterator)(unsafe.Pointer(bp + 16))).Fnext_state) != 0 {
						(*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 72))).Fstate = (*(*LookaheadIterator)(unsafe.Pointer(bp + 16))).Fnext_state
						libc.PostIncBitFieldPtr8Uint8(bp+72+4, 1, 7, 0, 0x7f)
					} else {
						continue
					}
				}
				*(*uint32)(unsafe.Pointer(bp + 80)) = uint32(0)
				*(*uint32)(unsafe.Pointer(bp + 12)) = uint32(0)
				size1 = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(subgraph+24)).Fsize - *(*uint32)(unsafe.Pointer(bp + 80))
				if size1 == uint32(0) {
					goto _5
				}
				for size1 > uint32(1) {
					half_size1 = size1 / uint32(2)
					mid_index1 = *(*uint32)(unsafe.Pointer(bp + 80)) + half_size1
					comparison2 = analysis_subgraph_node__compare(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(subgraph+24)).Fcontents+uintptr(mid_index1)*6, bp+72)
					if comparison2 <= 0 {
						*(*uint32)(unsafe.Pointer(bp + 80)) = mid_index1
					}
					size1 = size1 - half_size1
				}
				comparison2 = analysis_subgraph_node__compare(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(subgraph+24)).Fcontents+uintptr(*(*uint32)(unsafe.Pointer(bp + 80)))*6, bp+72)
				if comparison2 == 0 {
					*(*uint32)(unsafe.Pointer(bp + 12)) = uint32(1)
				} else {
					if comparison2 < 0 {
						*(*uint32)(unsafe.Pointer(bp + 80)) += uint32(1)
					}
				}
				goto _6
			_6:
				;
				goto _5
			_5: /**/
				; //
				for *(*uint32)(unsafe.Pointer(bp + 80)) < (*AnalysisSubgraph)(unsafe.Pointer(subgraph)).Fnodes.Fsize {
					_ = libc.Uint64FromInt64(4)
					{
						if !(*(*uint32)(unsafe.Pointer(bp + 80)) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(subgraph+24)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+5102, __ccgo_ts+3070, int32(1302), uintptr(unsafe.Pointer(&__func__42)))
						}
					}
					node = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(subgraph+24)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 80)))*6
					*(*uint32)(unsafe.Pointer(bp + 80)) = *(*uint32)(unsafe.Pointer(bp + 80)) + 1
					if libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer(node)).Fstate) != libc.Int32FromUint16((*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 72))).Fstate) || int32(*(*uint8)(unsafe.Pointer(node + 4))&0x7f>>0) != int32(*(*uint8)(unsafe.Pointer(bp + 72 + 4))&0x7f>>0) {
						break
					}
					alias = ts_language_alias_at(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, uint32((*AnalysisSubgraphNode)(unsafe.Pointer(node)).Fproduction_id), child_index)
					if alias != 0 {
						v7 = libc.Int32FromUint16(alias)
					} else {
						if (*(*TSSymbolMetadata)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Fsymbol_metadata + uintptr(sym)*3))).Fvisible != 0 {
							v8 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Fpublic_symbol_map + uintptr(sym)*2)))
						} else {
							v8 = 0
						}
						v7 = v8
					}
					visible_symbol = libc.Uint16FromInt32(v7)
					field_id = parent_field_id
					if !(field_id != 0) {
						ts_language_field_map(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, uint32((*AnalysisSubgraphNode)(unsafe.Pointer(node)).Fproduction_id), bp+88, bp+96)
						for {
							if !(*(*uintptr)(unsafe.Pointer(bp + 88)) != *(*uintptr)(unsafe.Pointer(bp + 96))) {
								break
							}
							if !((*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Finherited != 0) && uint32((*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Fchild_index) == child_index {
								field_id = (*TSFieldMapEntry)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Ffield_id
								break
							}
							goto _9
						_9:
							;
							*(*uintptr)(unsafe.Pointer(bp + 88)) += 4
						}
					}
					*(*AnalysisState)(unsafe.Pointer(bp + 104)) = *(*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))))
					next_state_top = analysis_state__top(tls, bp+104)
					(*AnalysisStateEntry)(unsafe.Pointer(next_state_top)).Fchild_index = libc.Uint16FromInt32(int32(*(*uint8)(unsafe.Pointer(bp + 72 + 4)) & 0x7f >> 0))
					(*AnalysisStateEntry)(unsafe.Pointer(next_state_top)).Fparse_state = (*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 72))).Fstate
					if int32(*(*uint8)(unsafe.Pointer(node + 4))&0x80>>7) != 0 {
						libc.SetBitFieldPtr16Uint8(next_state_top+6, libc.BoolUint8(1 != 0), 15, 0x8000)
					}
					does_match = libc.BoolUint8(0 != 0)
					if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) {
						does_match = libc.BoolUint8(1 != 0)
					} else {
						if visible_symbol != 0 {
							does_match = libc.BoolUint8(1 != 0)
							if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) == libc.Int32FromUint16(WILDCARD_SYMBOL) {
								if int32(*(*uint8)(unsafe.Pointer(step + 18))&0x1>>0) != 0 && !((*(*TSSymbolMetadata)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Fsymbol_metadata + uintptr(visible_symbol)*3))).Fnamed != 0) {
									does_match = libc.BoolUint8(0 != 0)
								}
							} else {
								if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) != libc.Int32FromUint16(visible_symbol) {
									does_match = libc.BoolUint8(0 != 0)
								}
							}
							if (*QueryStep)(unsafe.Pointer(step)).Ffield != 0 && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Ffield) != libc.Int32FromUint16(field_id) {
								does_match = libc.BoolUint8(0 != 0)
							}
							if (*QueryStep)(unsafe.Pointer(step)).Fsupertype_symbol != 0 && !(analysis_state__has_supertype(tls, *(*uintptr)(unsafe.Pointer(bp)), (*QueryStep)(unsafe.Pointer(step)).Fsupertype_symbol) != 0) {
								does_match = libc.BoolUint8(0 != 0)
							}
						} else {
							if uint32(sym) >= (*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Ftoken_count {
								if !(int32(uint8(*(*uint16)(unsafe.Pointer(next_state_top + 6))&0x8000>>15)) != 0) {
									if libc.Int32FromUint16((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fdepth)+int32(1) >= int32(8) {
										(*QueryAnalysis)(unsafe.Pointer(analysis)).Fdid_abort = libc.BoolUint8(1 != 0)
										continue
									}
									(*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fdepth = (*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fdepth + 1
									next_state_top = analysis_state__top(tls, bp+104)
								}
								*(*AnalysisStateEntry)(unsafe.Pointer(next_state_top)) = AnalysisStateEntry{
									Fparse_state:   parse_state,
									Fparent_symbol: sym,
									F__ccgo6:       field_id&0x7fff<<0 | libc.Uint16FromInt32(0)&0x1<<15,
								}
								if analysis_state__recursion_depth(tls, bp+104) > recursion_depth_limit {
									analysis_state_set__insert_sorted(tls, analysis+32, analysis+48, bp+104)
									continue
								}
							}
						}
					}
					for libc.Int32FromUint16((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fdepth) > 0 && int32(uint8(*(*uint16)(unsafe.Pointer(next_state_top + 6))&0x8000>>15)) != 0 {
						(*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fdepth = (*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fdepth - 1
						next_state_top = analysis_state__top(tls, bp+104)
					}
					next_step = step
					if does_match != 0 {
						for {
							(*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index = (*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index + 1
							_ = libc.Uint64FromInt64(4)
							{
								if !(uint32((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index) < (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+80)).Fsize) {
									libc.X__assert_fail(tls, __ccgo_ts+5152, __ccgo_ts+3070, int32(1409), uintptr(unsafe.Pointer(&__func__42)))
								}
							}
							next_step = (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+80)).Fcontents + uintptr((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index)*20
							if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) || libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth) <= libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth) {
								break
							}
							goto _10
						_10:
						}
					} else {
						if libc.Int32FromUint16((*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 72))).Fstate) == libc.Int32FromUint16(parse_state) {
							continue
						}
					}
					for {
						if int32(*(*uint8)(unsafe.Pointer(next_step + 18))&0x8>>3) != 0 {
							(*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index = (*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index + 1
							next_step += 20
							goto _11
						}
						if !(int32(*(*uint8)(unsafe.Pointer(next_step + 18))&0x10>>4) != 0) {
							_ = libc.Uint64FromInt64(4)
							{
								if !(uint32((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index) < (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+80)).Fsize) {
									libc.X__assert_fail(tls, __ccgo_ts+5152, __ccgo_ts+3070, int32(1433), uintptr(unsafe.Pointer(&__func__42)))
								}
							}
							did_finish_pattern = libc.BoolUint8(libc.Int32FromUint16((*QueryStep)(unsafe.Pointer((*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+80)).Fcontents+uintptr((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index)*20)).Fdepth) != libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth))
							if did_finish_pattern != 0 {
								*(*uint32)(unsafe.Pointer(bp + 176)) = uint32(0)
								*(*uint32)(unsafe.Pointer(bp + 180)) = uint32(0)
								size2 = (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(analysis+80)).Fsize - *(*uint32)(unsafe.Pointer(bp + 176))
								if size2 == uint32(0) {
									goto _12
								}
								for size2 > uint32(1) {
									half_size2 = size2 / uint32(2)
									mid_index2 = *(*uint32)(unsafe.Pointer(bp + 176)) + half_size2
									comparison3 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(analysis+80)).Fcontents + uintptr(mid_index2)*2))) - libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Froot_symbol)
									if comparison3 <= 0 {
										*(*uint32)(unsafe.Pointer(bp + 176)) = mid_index2
									}
									size2 = size2 - half_size2
								}
								comparison3 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(analysis+80)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 176)))*2))) - libc.Int32FromUint16((*AnalysisState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Froot_symbol)
								if comparison3 == 0 {
									*(*uint32)(unsafe.Pointer(bp + 180)) = uint32(1)
								} else {
									if comparison3 < 0 {
										*(*uint32)(unsafe.Pointer(bp + 176)) += uint32(1)
									}
								}
								goto _13
							_13:
								;
								goto _12
							_12: /**/
								; //
								if !(*(*uint32)(unsafe.Pointer(bp + 180)) != 0) {
									(*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(analysis + 80)).Fcontents = _array__splice(tls, (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(analysis+80)).Fcontents, analysis+80+8, analysis+80+12, libc.Uint64FromInt64(2), *(*uint32)(unsafe.Pointer(bp + 176)), uint32(0), uint32(1), *(*uintptr)(unsafe.Pointer(bp))+68)
								}
							} else {
								if libc.Int32FromUint16((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fdepth) == 0 {
									*(*uint32)(unsafe.Pointer(bp + 184)) = uint32(0)
									*(*uint32)(unsafe.Pointer(bp + 188)) = uint32(0)
									size3 = (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(analysis+64)).Fsize - *(*uint32)(unsafe.Pointer(bp + 184))
									if size3 == uint32(0) {
										goto _14
									}
									for size3 > uint32(1) {
										half_size3 = size3 / uint32(2)
										mid_index3 = *(*uint32)(unsafe.Pointer(bp + 184)) + half_size3
										comparison4 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(analysis+64)).Fcontents + uintptr(mid_index3)*2))) - libc.Int32FromUint16((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index)
										if comparison4 <= 0 {
											*(*uint32)(unsafe.Pointer(bp + 184)) = mid_index3
										}
										size3 = size3 - half_size3
									}
									comparison4 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(analysis+64)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 184)))*2))) - libc.Int32FromUint16((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index)
									if comparison4 == 0 {
										*(*uint32)(unsafe.Pointer(bp + 188)) = uint32(1)
									} else {
										if comparison4 < 0 {
											*(*uint32)(unsafe.Pointer(bp + 184)) += uint32(1)
										}
									}
									goto _15
								_15:
									;
									goto _14
								_14: /**/
									; //
									if !(*(*uint32)(unsafe.Pointer(bp + 188)) != 0) {
										(*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(analysis + 64)).Fcontents = _array__splice(tls, (*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(analysis+64)).Fcontents, analysis+64+8, analysis+64+12, libc.Uint64FromInt64(2), *(*uint32)(unsafe.Pointer(bp + 184)), uint32(0), uint32(1), bp+104+66)
									}
								} else {
									analysis_state_set__insert_sorted(tls, analysis+16, analysis+48, bp+104)
								}
							}
						}
						if does_match != 0 && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Falternative_index) != libc.Int32FromUint16(NONE) && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Falternative_index) > libc.Int32FromUint16((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index) {
							(*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index = (*QueryStep)(unsafe.Pointer(next_step)).Falternative_index
							_ = libc.Uint64FromInt64(4)
							{
								if !(uint32((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index) < (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+80)).Fsize) {
									libc.X__assert_fail(tls, __ccgo_ts+5152, __ccgo_ts+3070, int32(1453), uintptr(unsafe.Pointer(&__func__42)))
								}
							}
							next_step = (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+80)).Fcontents + uintptr((*(*AnalysisState)(unsafe.Pointer(bp + 104))).Fstep_index)*20
						} else {
							break
						}
						goto _11
					_11:
					}
				}
			}
			goto _2
		_2:
			;
			j = j + 1
		}
		_states1 = (*QueryAnalysis)(unsafe.Pointer(analysis)).Fstates
		(*QueryAnalysis)(unsafe.Pointer(analysis)).Fstates = (*QueryAnalysis)(unsafe.Pointer(analysis)).Fnext_states
		(*QueryAnalysis)(unsafe.Pointer(analysis)).Fnext_states = _states1
		goto _1
	_1:
		;
		iteration = iteration + 1
	}
}

var __func__42 = [27]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', '_', 'p', 'e', 'r', 'f', 'o', 'r', 'm', '_', 'a', 'n', 'a', 'l', 'y', 's', 'i', 's'}

func ts_query__analyze_patterns(tls *libc.TLS, self uintptr, error_offset uintptr) (r uint8) {
	bp := tls.Alloc(624)
	defer tls.Free(624)
	var action, child_step, next_step, pattern, pattern1, pattern_entry, predecessors, prev_step, step, step1, step2, step3, step4, step5, step6, step_offset, subgraph2, subgraph3, subgraph4, subgraph5, subgraph6, subtypes, symbol, symbol1, v3 uintptr
	var all_patterns_are_valid, done, has_children, is_valid_subtype, is_wildcard, parent_pattern_guaranteed uint8
	var capture_id, impossible_step_index, parent_depth, parent_step_index1, pattern_entry_index uint16_t
	var comparison, comparison1, comparison10, comparison2, comparison3, comparison4, comparison5, comparison6, comparison7, comparison8, comparison9 int32
	var end, end1, first_child_step_index, i, i1, i2, i3, i4, i5, i6, i7, i9, j, j1, j10, j2, j4, j5, j7, j8, j9, k1, k3, offset_idx, start, start1 uint32
	var final_step_index, half_size, half_size1, half_size10, half_size2, half_size3, half_size4, half_size5, half_size6, half_size7, half_size8, half_size9, i8, k, k2, mid_index, mid_index1, mid_index10, mid_index2, mid_index3, mid_index4, mid_index5, mid_index6, mid_index7, mid_index8, mid_index9, parent_step_index, size, size1, size10, size2, size3, size4, size5, size6, size7, size8, size9, v2 uint32_t
	var metadata TSSymbolMetadata
	var next_state, parse_state, parse_state1, state TSStateId
	var parent_symbol, parent_symbol1, sym TSSymbol
	var v22 bool
	var _ /* _exists at bp+244 */ uint32
	var _ /* _exists at bp+292 */ uint32
	var _ /* _exists at bp+600 */ uint32
	var _ /* _exists at bp+620 */ uint32
	var _ /* _index at bp+240 */ uint32
	var _ /* _index at bp+288 */ uint32
	var _ /* _index at bp+596 */ uint32
	var _ /* _index at bp+616 */ uint32
	var _ /* aliases at bp+368 */ uintptr
	var _ /* aliases at bp+384 */ uintptr
	var _ /* aliases_end at bp+376 */ uintptr
	var _ /* aliases_end at bp+392 */ uintptr
	var _ /* analysis at bp+448 */ QueryAnalysis
	var _ /* child_exists at bp+564 */ uint32_t
	var _ /* exists at bp+308 */ uint32
	var _ /* exists at bp+444 */ uint32
	var _ /* exists at bp+556 */ uint32
	var _ /* exists at bp+608 */ uint32
	var _ /* impossible_exists at bp+572 */ uint32_t
	var _ /* index at bp+440 */ uint32
	var _ /* index at bp+604 */ uint32
	var _ /* j at bp+560 */ uint32_t
	var _ /* j at bp+568 */ uint32_t
	var _ /* lookahead_iterator at bp+312 */ LookaheadIterator
	var _ /* next_nodes at bp+400 */ struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	var _ /* node at bp+420 */ AnalysisSubgraphNode
	var _ /* non_rooted_pattern_start_steps at bp+144 */ struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	var _ /* parent_step_indices at bp+160 */ struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	var _ /* predecessor_count at bp+428 */ uint32
	var _ /* predecessor_map at bp+296 */ StatePredecessorMap
	var _ /* predecessor_node at bp+432 */ AnalysisSubgraphNode
	var _ /* predicate_capture_ids at bp+576 */ struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	var _ /* subgraph at bp+200 */ AnalysisSubgraph
	var _ /* subgraph at bp+248 */ AnalysisSubgraph
	var _ /* subgraph_index at bp+304 */ uint32
	var _ /* subgraph_index at bp+552 */ uint32
	var _ /* subgraphs at bp+184 */ AnalysisSubgraphArray
	var _ /* subtype_length at bp+176 */ uint32_t
	var _ /* symbol at bp+612 */ TSSymbol
	var _ /* value_id at bp+592 */ uint16_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = action, all_patterns_are_valid, capture_id, child_step, comparison, comparison1, comparison10, comparison2, comparison3, comparison4, comparison5, comparison6, comparison7, comparison8, comparison9, done, end, end1, final_step_index, first_child_step_index, half_size, half_size1, half_size10, half_size2, half_size3, half_size4, half_size5, half_size6, half_size7, half_size8, half_size9, has_children, i, i1, i2, i3, i4, i5, i6, i7, i8, i9, impossible_step_index, is_valid_subtype, is_wildcard, j, j1, j10, j2, j4, j5, j7, j8, j9, k, k1, k2, k3, metadata, mid_index, mid_index1, mid_index10, mid_index2, mid_index3, mid_index4, mid_index5, mid_index6, mid_index7, mid_index8, mid_index9, next_state, next_step, offset_idx, parent_depth, parent_pattern_guaranteed, parent_step_index, parent_step_index1, parent_symbol, parent_symbol1, parse_state, parse_state1, pattern, pattern1, pattern_entry, pattern_entry_index, predecessors, prev_step, size, size1, size10, size2, size3, size4, size5, size6, size7, size8, size9, start, start1, state, step, step1, step2, step3, step4, step5, step6, step_offset, subgraph2, subgraph3, subgraph4, subgraph5, subgraph6, subtypes, sym, symbol, symbol1, v2, v22, v3
	*(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 144)) = struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}{}
	i = uint32(0)
	for {
		if !(i < (*TSQuery)(unsafe.Pointer(self)).Fpattern_map.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5209, __ccgo_ts+3070, int32(1471), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		pattern = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents + uintptr(i)*6
		if !((*PatternEntry)(unsafe.Pointer(pattern)).Fis_rooted != 0) {
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32((*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5252, __ccgo_ts+3070, int32(1473), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index)*20
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) != libc.Int32FromUint16(WILDCARD_SYMBOL) {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp + 144)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+144)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+144)).Fsize, bp+144+12, uint32(1), libc.Uint64FromInt64(2))
				v3 = bp + 144 + 8
				v2 = *(*uint32_t)(unsafe.Pointer(v3))
				*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
				*(*uint16_t)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+144)).Fcontents + uintptr(v2)*2)) = uint16(i)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 160)) = struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}{}
	all_patterns_are_valid = libc.BoolUint8(1 != 0)
	i1 = uint32(0)
	for {
		if !(i1 < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i1 < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5307, __ccgo_ts+3070, int32(1486), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		step1 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(i1)*20
		if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step1)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) {
			libc.SetBitFieldPtr8Uint8(step1+19, libc.BoolUint8(1 != 0), 0, 0x1)
			libc.SetBitFieldPtr8Uint8(step1+18, libc.BoolUint8(1 != 0), 7, 0x80)
			goto _4
		}
		has_children = libc.BoolUint8(0 != 0)
		is_wildcard = libc.BoolUint8(libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step1)).Fsymbol) == libc.Int32FromUint16(WILDCARD_SYMBOL))
		libc.SetBitFieldPtr8Uint8(step1+18, libc.BoolUint8(libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(step1 + 6))) != libc.Int32FromUint16(NONE)), 6, 0x40)
		j = i1 + uint32(1)
		for {
			if !(j < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5344, __ccgo_ts+3070, int32(1497), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			next_step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(j)*20
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) || libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth) <= libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step1)).Fdepth) {
				break
			}
			if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(next_step + 6))) != libc.Int32FromUint16(NONE) {
				libc.SetBitFieldPtr8Uint8(step1+18, libc.BoolUint8(1 != 0), 6, 0x40)
			}
			if !(is_wildcard != 0) {
				libc.SetBitFieldPtr8Uint8(next_step+18, libc.BoolUint8(1 != 0), 7, 0x80)
				libc.SetBitFieldPtr8Uint8(next_step+19, libc.BoolUint8(1 != 0), 0, 0x1)
			}
			has_children = libc.BoolUint8(1 != 0)
			goto _5
		_5:
			;
			j = j + 1
		}
		if has_children != 0 {
			if !(is_wildcard != 0) {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp + 160)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+160)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+160)).Fsize, bp+160+12, uint32(1), libc.Uint64FromInt64(4))
				v3 = bp + 160 + 8
				v2 = *(*uint32_t)(unsafe.Pointer(v3))
				*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
				*(*uint32_t)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+160)).Fcontents + uintptr(v2)*4)) = i1
			} else {
				if (*QueryStep)(unsafe.Pointer(step1)).Fsupertype_symbol != 0 && (*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Fabi_version >= uint32(15) {
					subtypes = ts_language_subtypes(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*QueryStep)(unsafe.Pointer(step1)).Fsupertype_symbol, bp+176)
					j1 = i1 + uint32(1)
					for {
						if !(j1 < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize) {
							break
						}
						_ = libc.Uint64FromInt64(4)
						{
							if !(j1 < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+80)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+5344, __ccgo_ts+3070, int32(1525), uintptr(unsafe.Pointer(&__func__43)))
							}
						}
						child_step = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+80)).Fcontents + uintptr(j1)*20
						if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(child_step)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) || libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(child_step)).Fdepth) <= libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step1)).Fdepth) {
							break
						}
						if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(child_step)).Fdepth) == libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step1)).Fdepth)+int32(1) && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(child_step)).Fsymbol) != libc.Int32FromUint16(WILDCARD_SYMBOL) {
							is_valid_subtype = libc.BoolUint8(0 != 0)
							k = uint32(0)
							for {
								if !(k < *(*uint32_t)(unsafe.Pointer(bp + 176))) {
									break
								}
								if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(child_step)).Fsymbol) == libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(subtypes + uintptr(k)*2))) {
									is_valid_subtype = libc.BoolUint8(1 != 0)
									break
								}
								goto _9
							_9:
								;
								k = k + 1
							}
							if !(is_valid_subtype != 0) {
								offset_idx = uint32(0)
								for {
									if !(offset_idx < (*TSQuery)(unsafe.Pointer(self)).Fstep_offsets.Fsize) {
										break
									}
									_ = libc.Uint64FromInt64(4)
									{
										if !(offset_idx < (*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(self+144)).Fsize) {
											libc.X__assert_fail(tls, __ccgo_ts+5381, __ccgo_ts+3070, int32(1540), uintptr(unsafe.Pointer(&__func__43)))
										}
									}
									step_offset = (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(self+144)).Fcontents + uintptr(offset_idx)*8
									if uint32((*StepOffset)(unsafe.Pointer(step_offset)).Fstep_index) >= j1 {
										*(*uint32)(unsafe.Pointer(error_offset)) = (*StepOffset)(unsafe.Pointer(step_offset)).Fbyte_offset
										all_patterns_are_valid = libc.BoolUint8(0 != 0)
										goto supertype_cleanup
									}
									goto _10
								_10:
									;
									offset_idx = offset_idx + 1
								}
							}
						}
						goto _8
					_8:
						;
						j1 = j1 + 1
					}
				}
			}
		}
		goto _4
	_4:
		;
		i1 = i1 + 1
	}
	*(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184)) = AnalysisSubgraphArray{}
	i2 = uint32(0)
	for {
		if !(i2 < (*(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp + 160))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i2 < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+160)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5434, __ccgo_ts+3070, int32(1564), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		parent_step_index = *(*uint32_t)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+160)).Fcontents + uintptr(i2)*4))
		_ = libc.Uint64FromInt64(4)
		{
			if !(parent_step_index < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5479, __ccgo_ts+3070, int32(1565), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		parent_symbol = (*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(parent_step_index)*20)).Fsymbol
		*(*AnalysisSubgraph)(unsafe.Pointer(bp + 200)) = AnalysisSubgraph{
			Fsymbol: parent_symbol,
		}
		*(*uint32)(unsafe.Pointer(bp + 240)) = uint32(0)
		*(*uint32)(unsafe.Pointer(bp + 244)) = uint32(0)
		size = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize - *(*uint32)(unsafe.Pointer(bp + 240))
		if size == uint32(0) {
			goto _12
		}
		for size > uint32(1) {
			half_size = size / uint32(2)
			mid_index = *(*uint32)(unsafe.Pointer(bp + 240)) + half_size
			comparison = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(mid_index)*40))) - libc.Int32FromUint16((*(*AnalysisSubgraph)(unsafe.Pointer(bp + 200))).Fsymbol)
			if comparison <= 0 {
				*(*uint32)(unsafe.Pointer(bp + 240)) = mid_index
			}
			size = size - half_size
		}
		comparison = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 240)))*40))) - libc.Int32FromUint16((*(*AnalysisSubgraph)(unsafe.Pointer(bp + 200))).Fsymbol)
		if comparison == 0 {
			*(*uint32)(unsafe.Pointer(bp + 244)) = uint32(1)
		} else {
			if comparison < 0 {
				*(*uint32)(unsafe.Pointer(bp + 240)) += uint32(1)
			}
		}
		goto _13
	_13:
		;
		goto _12
	_12: /**/
		; //
		if !(*(*uint32)(unsafe.Pointer(bp + 244)) != 0) {
			(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184)).Fcontents = _array__splice(tls, (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents, bp+184+8, bp+184+12, libc.Uint64FromInt64(40), *(*uint32)(unsafe.Pointer(bp + 240)), uint32(0), uint32(1), bp+200)
		}
		goto _11
	_11:
		;
		i2 = i2 + 1
	}
	sym = uint16((*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Ftoken_count)
	for {
		if !(libc.Int32FromUint16(sym) < libc.Int32FromUint16(uint16((*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Fsymbol_count))) {
			break
		}
		if !(ts_language_symbol_metadata(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, sym).Fvisible != 0) {
			*(*AnalysisSubgraph)(unsafe.Pointer(bp + 248)) = AnalysisSubgraph{
				Fsymbol: sym,
			}
			*(*uint32)(unsafe.Pointer(bp + 288)) = uint32(0)
			*(*uint32)(unsafe.Pointer(bp + 292)) = uint32(0)
			size1 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize - *(*uint32)(unsafe.Pointer(bp + 288))
			if size1 == uint32(0) {
				goto _15
			}
			for size1 > uint32(1) {
				half_size1 = size1 / uint32(2)
				mid_index1 = *(*uint32)(unsafe.Pointer(bp + 288)) + half_size1
				comparison1 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(mid_index1)*40))) - libc.Int32FromUint16((*(*AnalysisSubgraph)(unsafe.Pointer(bp + 248))).Fsymbol)
				if comparison1 <= 0 {
					*(*uint32)(unsafe.Pointer(bp + 288)) = mid_index1
				}
				size1 = size1 - half_size1
			}
			comparison1 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 288)))*40))) - libc.Int32FromUint16((*(*AnalysisSubgraph)(unsafe.Pointer(bp + 248))).Fsymbol)
			if comparison1 == 0 {
				*(*uint32)(unsafe.Pointer(bp + 292)) = uint32(1)
			} else {
				if comparison1 < 0 {
					*(*uint32)(unsafe.Pointer(bp + 288)) += uint32(1)
				}
			}
			goto _16
		_16:
			;
			goto _15
		_15: /**/
			; //
			if !(*(*uint32)(unsafe.Pointer(bp + 292)) != 0) {
				(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184)).Fcontents = _array__splice(tls, (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents, bp+184+8, bp+184+12, libc.Uint64FromInt64(40), *(*uint32)(unsafe.Pointer(bp + 288)), uint32(0), uint32(1), bp+248)
			}
		}
		goto _14
	_14:
		;
		sym = sym + 1
	}
	*(*StatePredecessorMap)(unsafe.Pointer(bp + 296)) = state_predecessor_map_new(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage)
	state = uint16(1)
	for {
		if !(libc.Int32FromUint16(state) < libc.Int32FromUint16(uint16((*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Fstate_count))) {
			break
		}
		*(*LookaheadIterator)(unsafe.Pointer(bp + 312)) = ts_language_lookaheads(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, state)
		for ts_lookahead_iterator__next(tls, bp+312) != 0 {
			if (*(*LookaheadIterator)(unsafe.Pointer(bp + 312))).Faction_count != 0 {
				i3 = uint32(0)
				for {
					if !(i3 < uint32((*(*LookaheadIterator)(unsafe.Pointer(bp + 312))).Faction_count)) {
						break
					}
					action = (*(*LookaheadIterator)(unsafe.Pointer(bp + 312))).Factions + uintptr(i3)*8
					if libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(action))) == int32(TSParseActionTypeReduce) {
						ts_language_aliases_for_symbol(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*(*struct {
							Ftype1              uint8_t
							Fchild_count        uint8_t
							Fsymbol             TSSymbol
							Fdynamic_precedence int16_t
							Fproduction_id      uint16_t
						})(unsafe.Pointer(action))).Fsymbol, bp+368, bp+376)
						symbol = *(*uintptr)(unsafe.Pointer(bp + 368))
						for {
							if !(symbol < *(*uintptr)(unsafe.Pointer(bp + 376))) {
								break
							}
							*(*uint32)(unsafe.Pointer(bp + 304)) = uint32(0)
							*(*uint32)(unsafe.Pointer(bp + 308)) = uint32(0)
							size2 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize - *(*uint32)(unsafe.Pointer(bp + 304))
							if size2 == uint32(0) {
								goto _20
							}
							for size2 > uint32(1) {
								half_size2 = size2 / uint32(2)
								mid_index2 = *(*uint32)(unsafe.Pointer(bp + 304)) + half_size2
								comparison2 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(mid_index2)*40))) - libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(symbol)))
								if comparison2 <= 0 {
									*(*uint32)(unsafe.Pointer(bp + 304)) = mid_index2
								}
								size2 = size2 - half_size2
							}
							comparison2 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 304)))*40))) - libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(symbol)))
							if comparison2 == 0 {
								*(*uint32)(unsafe.Pointer(bp + 308)) = uint32(1)
							} else {
								if comparison2 < 0 {
									*(*uint32)(unsafe.Pointer(bp + 304)) += uint32(1)
								}
							}
							goto _21
						_21:
							;
							goto _20
						_20: /**/
							; //
							if *(*uint32)(unsafe.Pointer(bp + 308)) != 0 {
								_ = libc.Uint64FromInt64(4)
								{
									if !(*(*uint32)(unsafe.Pointer(bp + 304)) < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
										libc.X__assert_fail(tls, __ccgo_ts+5532, __ccgo_ts+3070, int32(1607), uintptr(unsafe.Pointer(&__func__43)))
									}
								}
								subgraph2 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 304)))*40
								if v22 = (*AnalysisSubgraph)(unsafe.Pointer(subgraph2)).Fnodes.Fsize == uint32(0); !v22 {
									_ = libc.Uint64FromInt64(4)
									{
										if !((*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(subgraph2+24)).Fsize-libc.Uint32FromInt32(1) < (*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(subgraph2+24)).Fsize) {
											libc.X__assert_fail(tls, __ccgo_ts+5580, __ccgo_ts+3070, int32(1608), uintptr(unsafe.Pointer(&__func__43)))
										}
									}
								}
								if v22 || libc.Int32FromUint16((*AnalysisSubgraphNode)(unsafe.Pointer((*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(subgraph2+24)).Fcontents+uintptr((*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(subgraph2+24)).Fsize-uint32(1))*6)).Fstate) != libc.Int32FromUint16(state) {
									(*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph2 + 24)).Fcontents = _array__grow(tls, (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph2+24)).Fcontents, (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph2+24)).Fsize, subgraph2+24+12, uint32(1), libc.Uint64FromInt64(6))
									v3 = subgraph2 + 24 + 8
									v2 = *(*uint32_t)(unsafe.Pointer(v3))
									*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
									*(*AnalysisSubgraphNode)(unsafe.Pointer((*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph2+24)).Fcontents + uintptr(v2)*6)) = AnalysisSubgraphNode{
										Fstate: state,
										Fproduction_id: (*(*struct {
											Ftype1              uint8_t
											Fchild_count        uint8_t
											Fsymbol             TSSymbol
											Fdynamic_precedence int16_t
											Fproduction_id      uint16_t
										})(unsafe.Pointer(action))).Fproduction_id,
										F__ccgo4: (*(*struct {
											Ftype1              uint8_t
											Fchild_count        uint8_t
											Fsymbol             TSSymbol
											Fdynamic_precedence int16_t
											Fproduction_id      uint16_t
										})(unsafe.Pointer(action))).Fchild_count&0x7f<<0 | libc.Uint8FromInt32(1)&0x1<<7,
									}
								}
							}
							goto _19
						_19:
							;
							symbol += 2
						}
					} else {
						if libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(action))) == int32(TSParseActionTypeShift) && !((*TSParseAction)(unsafe.Pointer(action)).Fshift.Fextra != 0) {
							next_state = (*TSParseAction)(unsafe.Pointer(action)).Fshift.Fstate
							state_predecessor_map_add(tls, bp+296, next_state, state)
						}
					}
					goto _18
				_18:
					;
					i3 = i3 + 1
				}
			} else {
				if libc.Int32FromUint16((*(*LookaheadIterator)(unsafe.Pointer(bp + 312))).Fnext_state) != 0 {
					if libc.Int32FromUint16((*(*LookaheadIterator)(unsafe.Pointer(bp + 312))).Fnext_state) != libc.Int32FromUint16(state) {
						state_predecessor_map_add(tls, bp+296, (*(*LookaheadIterator)(unsafe.Pointer(bp + 312))).Fnext_state, state)
					}
					if ts_language_state_is_primary(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, state) != 0 {
						ts_language_aliases_for_symbol(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*(*LookaheadIterator)(unsafe.Pointer(bp + 312))).Fsymbol, bp+384, bp+392)
						symbol1 = *(*uintptr)(unsafe.Pointer(bp + 384))
						for {
							if !(symbol1 < *(*uintptr)(unsafe.Pointer(bp + 392))) {
								break
							}
							*(*uint32)(unsafe.Pointer(bp + 304)) = uint32(0)
							*(*uint32)(unsafe.Pointer(bp + 308)) = uint32(0)
							size3 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize - *(*uint32)(unsafe.Pointer(bp + 304))
							if size3 == uint32(0) {
								goto _26
							}
							for size3 > uint32(1) {
								half_size3 = size3 / uint32(2)
								mid_index3 = *(*uint32)(unsafe.Pointer(bp + 304)) + half_size3
								comparison3 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(mid_index3)*40))) - libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(symbol1)))
								if comparison3 <= 0 {
									*(*uint32)(unsafe.Pointer(bp + 304)) = mid_index3
								}
								size3 = size3 - half_size3
							}
							comparison3 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 304)))*40))) - libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(symbol1)))
							if comparison3 == 0 {
								*(*uint32)(unsafe.Pointer(bp + 308)) = uint32(1)
							} else {
								if comparison3 < 0 {
									*(*uint32)(unsafe.Pointer(bp + 304)) += uint32(1)
								}
							}
							goto _27
						_27:
							;
							goto _26
						_26: /**/
							; //
							if *(*uint32)(unsafe.Pointer(bp + 308)) != 0 {
								_ = libc.Uint64FromInt64(4)
								{
									if !(*(*uint32)(unsafe.Pointer(bp + 304)) < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
										libc.X__assert_fail(tls, __ccgo_ts+5532, __ccgo_ts+3070, int32(1644), uintptr(unsafe.Pointer(&__func__43)))
									}
								}
								subgraph3 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 304)))*40
								if v22 = (*AnalysisSubgraph)(unsafe.Pointer(subgraph3)).Fstart_states.Fsize == uint32(0); !v22 {
									_ = libc.Uint64FromInt64(4)
									{
										if !((*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(subgraph3+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(subgraph3+8)).Fsize) {
											libc.X__assert_fail(tls, __ccgo_ts+5648, __ccgo_ts+3070, int32(1647), uintptr(unsafe.Pointer(&__func__43)))
										}
									}
								}
								if v22 || libc.Int32FromUint16(*(*TSStateId)(unsafe.Pointer((*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(subgraph3+8)).Fcontents + uintptr((*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(subgraph3+8)).Fsize-uint32(1))*2))) != libc.Int32FromUint16(state) {
									(*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph3 + 8)).Fcontents = _array__grow(tls, (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph3+8)).Fcontents, (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph3+8)).Fsize, subgraph3+8+12, uint32(1), libc.Uint64FromInt64(2))
									v3 = subgraph3 + 8 + 8
									v2 = *(*uint32_t)(unsafe.Pointer(v3))
									*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
									*(*TSStateId)(unsafe.Pointer((*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(subgraph3+8)).Fcontents + uintptr(v2)*2)) = state
								}
							}
							goto _25
						_25:
							;
							symbol1 += 2
						}
					}
				}
			}
		}
		goto _17
	_17:
		;
		state = state + 1
	}
	*(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 400)) = struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}{}
	i4 = uint32(0)
	for {
		if !(i4 < (*(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i4 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1661), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		subgraph4 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(i4)*40
		if (*AnalysisSubgraph)(unsafe.Pointer(subgraph4)).Fnodes.Fsize == uint32(0) {
			if (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(subgraph4+8)).Fcontents != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(subgraph4+8)).Fcontents)
			}
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(subgraph4 + 8)).Fcontents = libc.UintptrFromInt32(0)
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(subgraph4 + 8)).Fsize = uint32(0)
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(subgraph4 + 8)).Fcapacity = uint32(0)
			_array__erase(tls, (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents, bp+184+8, libc.Uint64FromInt64(40), i4)
			i4 = i4 - 1
			goto _31
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp + 400)).Fcontents = _array__assign(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+400)).Fcontents, bp+400+8, bp+400+12, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(subgraph4+24)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(subgraph4+24)).Fsize, libc.Uint64FromInt64(6))
		for (*(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp + 400))).Fsize > uint32(0) {
			v3 = bp + 400 + 8
			*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) - 1
			v2 = *(*uint32_t)(unsafe.Pointer(v3))
			*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 420)) = *(*AnalysisSubgraphNode)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+400)).Fcontents + uintptr(v2)*6))
			if int32(*(*uint8)(unsafe.Pointer(bp + 420 + 4))&0x7f>>0) > int32(1) {
				predecessors = state_predecessor_map_get(tls, bp+296, (*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 420))).Fstate, bp+428)
				j2 = uint32(0)
				for {
					if !(j2 < *(*uint32)(unsafe.Pointer(bp + 428))) {
						break
					}
					*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 432)) = AnalysisSubgraphNode{
						Fstate:         *(*TSStateId)(unsafe.Pointer(predecessors + uintptr(j2)*2)),
						Fproduction_id: (*(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 420))).Fproduction_id,
						F__ccgo4:       libc.Uint8FromInt32(int32(*(*uint8)(unsafe.Pointer(bp + 420 + 4))&0x7f>>0)-int32(1))&0x7f<<0 | uint8(0)&0x1<<7,
					}
					*(*uint32)(unsafe.Pointer(bp + 440)) = uint32(0)
					*(*uint32)(unsafe.Pointer(bp + 444)) = uint32(0)
					size4 = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(subgraph4+24)).Fsize - *(*uint32)(unsafe.Pointer(bp + 440))
					if size4 == uint32(0) {
						goto _35
					}
					for size4 > uint32(1) {
						half_size4 = size4 / uint32(2)
						mid_index4 = *(*uint32)(unsafe.Pointer(bp + 440)) + half_size4
						comparison4 = analysis_subgraph_node__compare(tls, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(subgraph4+24)).Fcontents+uintptr(mid_index4)*6, bp+432)
						if comparison4 <= 0 {
							*(*uint32)(unsafe.Pointer(bp + 440)) = mid_index4
						}
						size4 = size4 - half_size4
					}
					comparison4 = analysis_subgraph_node__compare(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(subgraph4+24)).Fcontents+uintptr(*(*uint32)(unsafe.Pointer(bp + 440)))*6, bp+432)
					if comparison4 == 0 {
						*(*uint32)(unsafe.Pointer(bp + 444)) = uint32(1)
					} else {
						if comparison4 < 0 {
							*(*uint32)(unsafe.Pointer(bp + 440)) += uint32(1)
						}
					}
					goto _36
				_36:
					;
					goto _35
				_35: /**/
					; //
					if !(*(*uint32)(unsafe.Pointer(bp + 444)) != 0) {
						(*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(subgraph4 + 24)).Fcontents = _array__splice(tls, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(subgraph4+24)).Fcontents, subgraph4+24+8, subgraph4+24+12, libc.Uint64FromInt64(6), *(*uint32)(unsafe.Pointer(bp + 440)), uint32(0), uint32(1), bp+432)
						(*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(bp + 400)).Fcontents = _array__grow(tls, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(bp+400)).Fcontents, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(bp+400)).Fsize, bp+400+12, uint32(1), libc.Uint64FromInt64(6))
						v3 = bp + 400 + 8
						v2 = *(*uint32_t)(unsafe.Pointer(v3))
						*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
						*(*AnalysisSubgraphNode)(unsafe.Pointer((*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(bp+400)).Fcontents + uintptr(v2)*6)) = *(*AnalysisSubgraphNode)(unsafe.Pointer(bp + 432))
					}
					goto _34
				_34:
					;
					j2 = j2 + 1
				}
			}
		}
		goto _31
	_31:
		;
		i4 = i4 + 1
	}
	*(*QueryAnalysis)(unsafe.Pointer(bp + 448)) = query_analysis__new(tls)
	i5 = uint32(0)
	for {
		if !(i5 < (*(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp + 160))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i5 < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+160)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5434, __ccgo_ts+3070, int32(1725), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		parent_step_index1 = uint16(*(*uint32_t)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+160)).Fcontents + uintptr(i5)*4)))
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(parent_step_index1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5479, __ccgo_ts+3070, int32(1726), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		parent_depth = (*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(parent_step_index1)*20)).Fdepth
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(parent_step_index1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5479, __ccgo_ts+3070, int32(1727), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		parent_symbol1 = (*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(parent_step_index1)*20)).Fsymbol
		if libc.Int32FromUint16(parent_symbol1) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) {
			goto _39
		}
		*(*uint32)(unsafe.Pointer(bp + 552)) = uint32(0)
		*(*uint32)(unsafe.Pointer(bp + 556)) = uint32(0)
		size5 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize - *(*uint32)(unsafe.Pointer(bp + 552))
		if size5 == uint32(0) {
			goto _40
		}
		for size5 > uint32(1) {
			half_size5 = size5 / uint32(2)
			mid_index5 = *(*uint32)(unsafe.Pointer(bp + 552)) + half_size5
			comparison5 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(mid_index5)*40))) - libc.Int32FromUint16(parent_symbol1)
			if comparison5 <= 0 {
				*(*uint32)(unsafe.Pointer(bp + 552)) = mid_index5
			}
			size5 = size5 - half_size5
		}
		comparison5 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 552)))*40))) - libc.Int32FromUint16(parent_symbol1)
		if comparison5 == 0 {
			*(*uint32)(unsafe.Pointer(bp + 556)) = uint32(1)
		} else {
			if comparison5 < 0 {
				*(*uint32)(unsafe.Pointer(bp + 552)) += uint32(1)
			}
		}
		goto _41
	_41:
		;
		goto _40
	_40: /**/
		; //
		if !(*(*uint32)(unsafe.Pointer(bp + 556)) != 0) {
			first_child_step_index = libc.Uint32FromInt32(libc.Int32FromUint16(parent_step_index1) + int32(1))
			*(*uint32_t)(unsafe.Pointer(bp + 560)) = uint32(0)
			*(*uint32_t)(unsafe.Pointer(bp + 564)) = uint32(0)
			size6 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fsize - *(*uint32_t)(unsafe.Pointer(bp + 560))
			if size6 == uint32(0) {
				goto _42
			}
			for size6 > uint32(1) {
				half_size6 = size6 / uint32(2)
				mid_index6 = *(*uint32_t)(unsafe.Pointer(bp + 560)) + half_size6
				comparison6 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+144)).Fcontents + uintptr(mid_index6)*8 + 4))) - libc.Int32FromUint32(first_child_step_index)
				if comparison6 <= 0 {
					*(*uint32_t)(unsafe.Pointer(bp + 560)) = mid_index6
				}
				size6 = size6 - half_size6
			}
			comparison6 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 560)))*8 + 4))) - libc.Int32FromUint32(first_child_step_index)
			if comparison6 == 0 {
				*(*uint32_t)(unsafe.Pointer(bp + 564)) = uint32(1)
			} else {
				if comparison6 < 0 {
					*(*uint32_t)(unsafe.Pointer(bp + 560)) += uint32(1)
				}
			}
			goto _43
		_43:
			;
			goto _42
		_42: /**/
			; //
			_ = libc.Uint64FromInt64(4)
			{
				if !(*(*uint32_t)(unsafe.Pointer(bp + 564)) != 0) {
					libc.X__assert_fail(tls, __ccgo_ts+5765, __ccgo_ts+3070, int32(1738), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(*(*uint32_t)(unsafe.Pointer(bp + 560)) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+144)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5778, __ccgo_ts+3070, int32(1739), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			*(*uint32)(unsafe.Pointer(error_offset)) = (*StepOffset)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 560)))*8)).Fbyte_offset
			all_patterns_are_valid = libc.BoolUint8(0 != 0)
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(*(*uint32)(unsafe.Pointer(bp + 552)) < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5532, __ccgo_ts+3070, int32(1746), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		subgraph5 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 552)))*40
		analysis_state_set__clear(tls, bp+448, bp+448+48)
		analysis_state_set__clear(tls, bp+448+32, bp+448+48)
		j4 = uint32(0)
		for {
			if !(j4 < (*AnalysisSubgraph)(unsafe.Pointer(subgraph5)).Fstart_states.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j4 < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(subgraph5+8)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5822, __ccgo_ts+3070, int32(1750), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			parse_state = *(*TSStateId)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(subgraph5+8)).Fcontents + uintptr(j4)*2))
			*(*AnalysisState)(unsafe.Pointer(bp)) = AnalysisState{
				Fstack: [8]AnalysisStateEntry{
					0: {
						Fparse_state:   parse_state,
						Fparent_symbol: parent_symbol1,
						F__ccgo6:       libc.Uint16FromInt32(0)&0x7fff<<0 | libc.Uint16FromInt32(0)&0x1<<15,
					},
				},
				Fdepth:       uint16(1),
				Fstep_index:  libc.Uint16FromInt32(libc.Int32FromUint16(parent_step_index1) + int32(1)),
				Froot_symbol: parent_symbol1,
			}
			analysis_state_set__push(tls, bp+448, bp+448+48, bp)
			goto _44
		_44:
			;
			j4 = j4 + 1
		}
		(*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Fdid_abort = libc.BoolUint8(0 != 0)
		ts_query__perform_analysis(tls, self, bp+184, bp+448)
		if (*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Fdid_abort != 0 {
			j5 = libc.Uint32FromInt32(libc.Int32FromUint16(parent_step_index1) + int32(1))
			for {
				if !(j5 < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(j5 < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+5344, __ccgo_ts+3070, int32(1781), uintptr(unsafe.Pointer(&__func__43)))
					}
				}
				step2 = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fcontents + uintptr(j5)*20
				if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step2)).Fdepth) <= libc.Int32FromUint16(parent_depth) || libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step2)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) {
					break
				}
				if !(int32(*(*uint8)(unsafe.Pointer(step2 + 18))&0x10>>4) != 0) {
					libc.SetBitFieldPtr8Uint8(step2+19, libc.BoolUint8(0 != 0), 0, 0x1)
					libc.SetBitFieldPtr8Uint8(step2+18, libc.BoolUint8(0 != 0), 7, 0x80)
				}
				goto _45
			_45:
				;
				j5 = j5 + 1
			}
			goto _39
		}
		if (*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Ffinished_parent_symbols.Fsize == uint32(0) {
			if (*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Ffinal_step_indices.Fsize > uint32(0) {
				_ = libc.Uint64FromInt64(4)
				{
					if !((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp+448+64)).Fsize-libc.Uint32FromInt32(1) < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp+448+64)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+5870, __ccgo_ts+3070, int32(1799), uintptr(unsafe.Pointer(&__func__43)))
					}
				}
				impossible_step_index = *(*uint16_t)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+448+64)).Fcontents + uintptr((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+448+64)).Fsize-uint32(1))*2))
			} else {
				impossible_step_index = parent_step_index1
			}
			*(*uint32_t)(unsafe.Pointer(bp + 568)) = uint32(0)
			*(*uint32_t)(unsafe.Pointer(bp + 572)) = uint32(0)
			size7 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fsize - *(*uint32_t)(unsafe.Pointer(bp + 568))
			if size7 == uint32(0) {
				goto _46
			}
			for size7 > uint32(1) {
				half_size7 = size7 / uint32(2)
				mid_index7 = *(*uint32_t)(unsafe.Pointer(bp + 568)) + half_size7
				comparison7 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+144)).Fcontents + uintptr(mid_index7)*8 + 4))) - libc.Int32FromUint16(impossible_step_index)
				if comparison7 <= 0 {
					*(*uint32_t)(unsafe.Pointer(bp + 568)) = mid_index7
				}
				size7 = size7 - half_size7
			}
			comparison7 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 568)))*8 + 4))) - libc.Int32FromUint16(impossible_step_index)
			if comparison7 == 0 {
				*(*uint32_t)(unsafe.Pointer(bp + 572)) = uint32(1)
			} else {
				if comparison7 < 0 {
					*(*uint32_t)(unsafe.Pointer(bp + 568)) += uint32(1)
				}
			}
			goto _47
		_47:
			;
			goto _46
		_46: /**/
			; //
			if *(*uint32_t)(unsafe.Pointer(bp + 568)) >= (*TSQuery)(unsafe.Pointer(self)).Fstep_offsets.Fsize {
				*(*uint32_t)(unsafe.Pointer(bp + 568)) = (*TSQuery)(unsafe.Pointer(self)).Fstep_offsets.Fsize - uint32(1)
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(*(*uint32_t)(unsafe.Pointer(bp + 568)) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+144)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5778, __ccgo_ts+3070, int32(1807), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			*(*uint32)(unsafe.Pointer(error_offset)) = (*StepOffset)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 568)))*8)).Fbyte_offset
			all_patterns_are_valid = libc.BoolUint8(0 != 0)
			break
		}
		j7 = uint32(0)
		for {
			if !(j7 < (*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Ffinal_step_indices.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j7 < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+448+64)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5962, __ccgo_ts+3070, int32(1815), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			final_step_index = uint32(*(*uint16_t)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+448+64)).Fcontents + uintptr(j7)*2)))
			_ = libc.Uint64FromInt64(4)
			{
				if !(final_step_index < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6015, __ccgo_ts+3070, int32(1816), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			step3 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(final_step_index)*20
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Fdepth) != libc.Int32FromUint16(PATTERN_DONE_MARKER) && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Fdepth) > libc.Int32FromUint16(parent_depth) && !(int32(*(*uint8)(unsafe.Pointer(step3 + 18))&0x10>>4) != 0) {
				libc.SetBitFieldPtr8Uint8(step3+19, libc.BoolUint8(0 != 0), 0, 0x1)
				libc.SetBitFieldPtr8Uint8(step3+18, libc.BoolUint8(0 != 0), 7, 0x80)
			}
			goto _48
		_48:
			;
			j7 = j7 + 1
		}
		goto _39
	_39:
		;
		i5 = i5 + 1
	}
	*(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 576)) = struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}{}
	i6 = uint32(0)
	for {
		if !(i6 < (*TSQuery)(unsafe.Pointer(self)).Fpatterns.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i6 < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+128)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+6067, __ccgo_ts+3070, int32(1831), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		pattern1 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fcontents + uintptr(i6)*28
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp + 576)).Fsize = uint32(0)
		start = (*QueryPattern)(unsafe.Pointer(pattern1)).Fpredicate_steps.Foffset
		end = start + (*QueryPattern)(unsafe.Pointer(pattern1)).Fpredicate_steps.Flength
		j8 = start
		for {
			if !(j8 < end) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j8 < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+112)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6107, __ccgo_ts+3070, int32(1840), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			step4 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+112)).Fcontents + uintptr(j8)*8
			if (*TSQueryPredicateStep)(unsafe.Pointer(step4)).Ftype1 == int32(TSQueryPredicateStepTypeCapture) {
				*(*uint16_t)(unsafe.Pointer(bp + 592)) = uint16((*TSQueryPredicateStep)(unsafe.Pointer(step4)).Fvalue_id)
				*(*uint32)(unsafe.Pointer(bp + 596)) = uint32(0)
				*(*uint32)(unsafe.Pointer(bp + 600)) = uint32(0)
				size8 = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+576)).Fsize - *(*uint32)(unsafe.Pointer(bp + 596))
				if size8 == uint32(0) {
					goto _51
				}
				for size8 > uint32(1) {
					half_size8 = size8 / uint32(2)
					mid_index8 = *(*uint32)(unsafe.Pointer(bp + 596)) + half_size8
					comparison8 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp+576)).Fcontents + uintptr(mid_index8)*2))) - libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp + 592)))
					if comparison8 <= 0 {
						*(*uint32)(unsafe.Pointer(bp + 596)) = mid_index8
					}
					size8 = size8 - half_size8
				}
				comparison8 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+576)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 596)))*2))) - libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(bp + 592)))
				if comparison8 == 0 {
					*(*uint32)(unsafe.Pointer(bp + 600)) = uint32(1)
				} else {
					if comparison8 < 0 {
						*(*uint32)(unsafe.Pointer(bp + 596)) += uint32(1)
					}
				}
				goto _52
			_52:
				;
				goto _51
			_51: /**/
				; //
				if !(*(*uint32)(unsafe.Pointer(bp + 600)) != 0) {
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp + 576)).Fcontents = _array__splice(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp+576)).Fcontents, bp+576+8, bp+576+12, libc.Uint64FromInt64(2), *(*uint32)(unsafe.Pointer(bp + 596)), uint32(0), uint32(1), bp+592)
				}
			}
			goto _50
		_50:
			;
			j8 = j8 + 1
		}
		start1 = (*QueryPattern)(unsafe.Pointer(pattern1)).Fsteps.Foffset
		end1 = start1 + (*QueryPattern)(unsafe.Pointer(pattern1)).Fsteps.Flength
		j9 = start1
		for {
			if !(j9 < end1) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j9 < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5344, __ccgo_ts+3070, int32(1853), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			step5 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(j9)*20
			k1 = uint32(0)
			for {
				if !(k1 < uint32(3)) {
					break
				}
				capture_id = *(*uint16_t)(unsafe.Pointer(step5 + 6 + uintptr(k1)*2))
				if libc.Int32FromUint16(capture_id) == libc.Int32FromUint16(NONE) {
					break
				}
				*(*uint32)(unsafe.Pointer(bp + 604)) = uint32(0)
				*(*uint32)(unsafe.Pointer(bp + 608)) = uint32(0)
				size9 = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+576)).Fsize - *(*uint32)(unsafe.Pointer(bp + 604))
				if size9 == uint32(0) {
					goto _55
				}
				for size9 > uint32(1) {
					half_size9 = size9 / uint32(2)
					mid_index9 = *(*uint32)(unsafe.Pointer(bp + 604)) + half_size9
					comparison9 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp+576)).Fcontents + uintptr(mid_index9)*2))) - libc.Int32FromUint16(capture_id)
					if comparison9 <= 0 {
						*(*uint32)(unsafe.Pointer(bp + 604)) = mid_index9
					}
					size9 = size9 - half_size9
				}
				comparison9 = libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+576)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 604)))*2))) - libc.Int32FromUint16(capture_id)
				if comparison9 == 0 {
					*(*uint32)(unsafe.Pointer(bp + 608)) = uint32(1)
				} else {
					if comparison9 < 0 {
						*(*uint32)(unsafe.Pointer(bp + 604)) += uint32(1)
					}
				}
				goto _56
			_56:
				;
				goto _55
			_55: /**/
				; //
				if *(*uint32)(unsafe.Pointer(bp + 608)) != 0 {
					libc.SetBitFieldPtr8Uint8(step5+18, libc.BoolUint8(0 != 0), 7, 0x80)
					break
				}
				goto _54
			_54:
				;
				k1 = k1 + 1
			}
			goto _53
		_53:
			;
			j9 = j9 + 1
		}
		goto _49
	_49:
		;
		i6 = i6 + 1
	}
	done = libc.BoolUint8((*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize == uint32(0))
	for !(done != 0) {
		done = libc.BoolUint8(1 != 0)
		i7 = (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize - uint32(1)
		for {
			if !(i7 > uint32(0)) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i7 < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5307, __ccgo_ts+3070, int32(1873), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			step6 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(i7)*20
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step6)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) {
				goto _57
			}
			parent_pattern_guaranteed = libc.BoolUint8(0 != 0)
			for {
				if int32(*(*uint8)(unsafe.Pointer(step6 + 18))&0x80>>7) != 0 {
					parent_pattern_guaranteed = libc.BoolUint8(1 != 0)
					break
				}
				if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step6)).Falternative_index) == libc.Int32FromUint16(NONE) || uint32((*QueryStep)(unsafe.Pointer(step6)).Falternative_index) < i7 {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(uint32((*QueryStep)(unsafe.Pointer(step6)).Falternative_index) < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+6154, __ccgo_ts+3070, int32(1886), uintptr(unsafe.Pointer(&__func__43)))
					}
				}
				step6 = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fcontents + uintptr((*QueryStep)(unsafe.Pointer(step6)).Falternative_index)*20
				goto _58
			_58:
			}
			if !(parent_pattern_guaranteed != 0) {
				_ = libc.Uint64FromInt64(4)
				{
					if !(i7-libc.Uint32FromInt32(1) < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+6213, __ccgo_ts+3070, int32(1891), uintptr(unsafe.Pointer(&__func__43)))
					}
				}
				prev_step = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fcontents + uintptr(i7-uint32(1))*20
				if !(int32(*(*uint8)(unsafe.Pointer(prev_step + 18))&0x10>>4) != 0) && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(prev_step)).Fdepth) != libc.Int32FromUint16(PATTERN_DONE_MARKER) && int32(*(*uint8)(unsafe.Pointer(prev_step + 18))&0x80>>7) != 0 {
					libc.SetBitFieldPtr8Uint8(prev_step+18, libc.BoolUint8(0 != 0), 7, 0x80)
					done = libc.BoolUint8(0 != 0)
				}
			}
			goto _57
		_57:
			;
			i7 = i7 - 1
		}
	}
	(*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Fdid_abort = libc.BoolUint8(0 != 0)
	i8 = uint32(0)
	for {
		if !(i8 < (*(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp + 144))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i8 < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+144)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+6254, __ccgo_ts+3070, int32(1931), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		pattern_entry_index = *(*uint16_t)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+144)).Fcontents + uintptr(i8)*2))
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(pattern_entry_index) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+6310, __ccgo_ts+3070, int32(1932), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		pattern_entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents + uintptr(pattern_entry_index)*6
		analysis_state_set__clear(tls, bp+448, bp+448+48)
		analysis_state_set__clear(tls, bp+448+32, bp+448+48)
		j10 = uint32(0)
		for {
			if !(j10 < (*(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184))).Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(j10 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6371, __ccgo_ts+3070, int32(1937), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			subgraph6 = (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(j10)*40
			metadata = ts_language_symbol_metadata(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*AnalysisSubgraph)(unsafe.Pointer(subgraph6)).Fsymbol)
			if metadata.Fvisible != 0 || metadata.Fnamed != 0 {
				goto _60
			}
			k2 = uint32(0)
			for {
				if !(k2 < (*AnalysisSubgraph)(unsafe.Pointer(subgraph6)).Fstart_states.Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(k2 < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(subgraph6+8)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+6406, __ccgo_ts+3070, int32(1942), uintptr(unsafe.Pointer(&__func__43)))
					}
				}
				parse_state1 = *(*TSStateId)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(subgraph6+8)).Fcontents + uintptr(k2)*2))
				*(*AnalysisState)(unsafe.Pointer(bp + 70)) = AnalysisState{
					Fstack: [8]AnalysisStateEntry{
						0: {
							Fparse_state:   parse_state1,
							Fparent_symbol: (*AnalysisSubgraph)(unsafe.Pointer(subgraph6)).Fsymbol,
							F__ccgo6:       libc.Uint16FromInt32(0)&0x7fff<<0 | libc.Uint16FromInt32(0)&0x1<<15,
						},
					},
					Fdepth:       uint16(1),
					Fstep_index:  (*PatternEntry)(unsafe.Pointer(pattern_entry)).Fstep_index,
					Froot_symbol: (*AnalysisSubgraph)(unsafe.Pointer(subgraph6)).Fsymbol,
				}
				analysis_state_set__push(tls, bp+448, bp+448+48, bp+70)
				goto _61
			_61:
				;
				k2 = k2 + 1
			}
			goto _60
		_60:
			;
			j10 = j10 + 1
		}
		ts_query__perform_analysis(tls, self, bp+184, bp+448)
		if (*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Ffinished_parent_symbols.Fsize > uint32(0) {
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32((*PatternEntry)(unsafe.Pointer(pattern_entry)).Fpattern_index) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+128)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6454, __ccgo_ts+3070, int32(1971), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			(*QueryPattern)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+128)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer(pattern_entry)).Fpattern_index)*28)).Fis_non_local = libc.BoolUint8(1 != 0)
		}
		k3 = uint32(0)
		for {
			if !(k3 < (*(*QueryAnalysis)(unsafe.Pointer(bp + 448))).Ffinished_parent_symbols.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(k3 < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp+448+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6521, __ccgo_ts+3070, int32(1975), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			*(*TSSymbol)(unsafe.Pointer(bp + 612)) = *(*TSSymbol)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+448+80)).Fcontents + uintptr(k3)*2))
			*(*uint32)(unsafe.Pointer(bp + 616)) = uint32(0)
			*(*uint32)(unsafe.Pointer(bp + 620)) = uint32(0)
			size10 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+192)).Fsize - *(*uint32)(unsafe.Pointer(bp + 616))
			if size10 == uint32(0) {
				goto _63
			}
			for size10 > uint32(1) {
				half_size10 = size10 / uint32(2)
				mid_index10 = *(*uint32)(unsafe.Pointer(bp + 616)) + half_size10
				comparison10 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+192)).Fcontents + uintptr(mid_index10)*2))) - libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(bp + 612)))
				if comparison10 <= 0 {
					*(*uint32)(unsafe.Pointer(bp + 616)) = mid_index10
				}
				size10 = size10 - half_size10
			}
			comparison10 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+192)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 616)))*2))) - libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(bp + 612)))
			if comparison10 == 0 {
				*(*uint32)(unsafe.Pointer(bp + 620)) = uint32(1)
			} else {
				if comparison10 < 0 {
					*(*uint32)(unsafe.Pointer(bp + 616)) += uint32(1)
				}
			}
			goto _64
		_64:
			;
			goto _63
		_63: /**/
			; //
			if !(*(*uint32)(unsafe.Pointer(bp + 620)) != 0) {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 192)).Fcontents = _array__splice(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+192)).Fcontents, self+192+8, self+192+12, libc.Uint64FromInt64(2), *(*uint32)(unsafe.Pointer(bp + 616)), uint32(0), uint32(1), bp+612)
			}
			goto _62
		_62:
			;
			k3 = k3 + 1
		}
		goto _59
	_59:
		;
		i8 = i8 + 1
	}
	i9 = uint32(0)
	for {
		if !(i9 < (*(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184))).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1994), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents+uintptr(i9)*40+8)).Fcontents != 0 {
			_ = libc.Uint64FromInt64(4)
			{
				if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1994), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents+uintptr(i9)*40+8)).Fcontents)
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1994), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(i9)*40 + 8)).Fcontents = libc.UintptrFromInt32(0)
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1994), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(i9)*40 + 8)).Fsize = uint32(0)
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1994), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(i9)*40 + 8)).Fcapacity = uint32(0)
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1995), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents+uintptr(i9)*40+24)).Fcontents != 0 {
			_ = libc.Uint64FromInt64(4)
			{
				if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1995), uintptr(unsafe.Pointer(&__func__43)))
				}
			}
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents+uintptr(i9)*40+24)).Fcontents)
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1995), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(i9)*40 + 24)).Fcontents = libc.UintptrFromInt32(0)
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1995), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(i9)*40 + 24)).Fsize = uint32(0)
		_ = libc.Uint64FromInt64(4)
		{
			if !(i9 < (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5730, __ccgo_ts+3070, int32(1995), uintptr(unsafe.Pointer(&__func__43)))
			}
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents + uintptr(i9)*40 + 24)).Fcapacity = uint32(0)
		goto _65
	_65:
		;
		i9 = i9 + 1
	}
	if (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*AnalysisSubgraphArray)(unsafe.Pointer(bp+184)).Fcontents)
	}
	(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184)).Fcontents = libc.UintptrFromInt32(0)
	(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184)).Fsize = uint32(0)
	(*AnalysisSubgraphArray)(unsafe.Pointer(bp + 184)).Fcapacity = uint32(0)
	query_analysis__delete(tls, bp+448)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+400)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+400)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 400)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 400)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 400)).Fcapacity = uint32(0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+576)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+576)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 576)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 576)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 576)).Fcapacity = uint32(0)
	state_predecessor_map_delete(tls, bp+296)
	goto supertype_cleanup
supertype_cleanup:
	;
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+144)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+144)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 144)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 144)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 144)).Fcapacity = uint32(0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+160)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+160)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 160)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 160)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 160)).Fcapacity = uint32(0)
	return all_patterns_are_valid
}

var __func__43 = [27]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', '_', 'a', 'n', 'a', 'l', 'y', 'z', 'e', '_', 'p', 'a', 't', 't', 'e', 'r', 'n', 's'}

func ts_query__add_negated_fields(tls *libc.TLS, self uintptr, step_index uint16_t, field_ids uintptr, field_count uint16_t) {
	var existing_field_id TSFieldId
	var failed_match uint8
	var i, match_count, start_i uint32
	var step, v3 uintptr
	var v2 uint32_t
	_, _, _, _, _, _, _, _ = existing_field_id, failed_match, i, match_count, start_i, step, v2, v3
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(step_index) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(2016), uintptr(unsafe.Pointer(&__func__44)))
		}
	}
	step = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index)*20
	failed_match = libc.BoolUint8(0 != 0)
	match_count = uint32(0)
	start_i = uint32(0)
	i = uint32(0)
	for {
		if !(i < (*TSQuery)(unsafe.Pointer(self)).Fnegated_fields.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+160)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+6625, __ccgo_ts+3070, int32(2024), uintptr(unsafe.Pointer(&__func__44)))
			}
		}
		existing_field_id = *(*TSFieldId)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+160)).Fcontents + uintptr(i)*2))
		if libc.Int32FromUint16(existing_field_id) == 0 {
			if match_count == uint32(field_count) {
				(*QueryStep)(unsafe.Pointer(step)).Fnegated_field_list_id = uint16(start_i)
				return
			} else {
				start_i = i + uint32(1)
				match_count = uint32(0)
				failed_match = libc.BoolUint8(0 != 0)
			}
		} else {
			if match_count < uint32(field_count) && libc.Int32FromUint16(existing_field_id) == libc.Int32FromUint16(*(*TSFieldId)(unsafe.Pointer(field_ids + uintptr(match_count)*2))) && !(failed_match != 0) {
				match_count = match_count + 1
			} else {
				match_count = uint32(0)
				failed_match = libc.BoolUint8(1 != 0)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	(*QueryStep)(unsafe.Pointer(step)).Fnegated_field_list_id = uint16((*TSQuery)(unsafe.Pointer(self)).Fnegated_fields.Fsize)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 160)).Fcontents = _array__splice(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fcontents, self+160+8, self+160+12, libc.Uint64FromInt64(2), (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fsize, uint32(0), uint32(field_count), field_ids)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 160)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fsize, self+160+12, uint32(1), libc.Uint64FromInt64(2))
	v3 = self + 160 + 8
	v2 = *(*uint32_t)(unsafe.Pointer(v3))
	*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
	*(*TSFieldId)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fcontents + uintptr(v2)*2)) = libc.Uint16FromInt32(libc.Int32FromInt32(0))
}

var __func__44 = [29]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', '_', 'a', 'd', 'd', '_', 'n', 'e', 'g', 'a', 't', 'e', 'd', '_', 'f', 'i', 'e', 'l', 'd', 's'}

func ts_query__parse_string_literal(tls *libc.TLS, self uintptr, stream uintptr) (r TSQueryError1) {
	var is_escaped uint8
	var prev_position, string_start, v12 uintptr
	var v11 uint32_t
	_, _, _, _, _ = is_escaped, prev_position, string_start, v11, v12
	string_start = (*Stream)(unsafe.Pointer(stream)).Finput
	if (*Stream)(unsafe.Pointer(stream)).Fnext != int32('"') {
		return int32(TSQueryErrorSyntax)
	}
	stream_advance(tls, stream)
	prev_position = (*Stream)(unsafe.Pointer(stream)).Finput
	is_escaped = libc.BoolUint8(0 != 0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 176)).Fsize = uint32(0)
	for {
		if is_escaped != 0 {
			is_escaped = libc.BoolUint8(0 != 0)
			switch (*Stream)(unsafe.Pointer(stream)).Fnext {
			case int32('n'):
				goto _2
			case int32('r'):
				goto _3
			case int32('t'):
				goto _4
			case int32('0'):
				goto _5
			default:
				goto _6
			}
			goto _7
		_2:
			;
		_10:
			;
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 176)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fsize, self+176+12, uint32(1), libc.Uint64FromInt64(1))
			v12 = self + 176 + 8
			v11 = *(*uint32_t)(unsafe.Pointer(v12))
			*(*uint32_t)(unsafe.Pointer(v12)) = *(*uint32_t)(unsafe.Pointer(v12)) + 1
			*(*int8)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents + uintptr(v11))) = int8(libc.Int32FromUint8('\n'))
			goto _9
		_9:
			;
			if 0 != 0 {
				goto _10
			}
			goto _8
		_8:
			;
			goto _7
		_3:
			;
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 176)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fsize, self+176+12, uint32(1), libc.Uint64FromInt64(1))
			v12 = self + 176 + 8
			v11 = *(*uint32_t)(unsafe.Pointer(v12))
			*(*uint32_t)(unsafe.Pointer(v12)) = *(*uint32_t)(unsafe.Pointer(v12)) + 1
			*(*int8)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents + uintptr(v11))) = int8(libc.Int32FromUint8('\r'))
			goto _7
		_4:
			;
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 176)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fsize, self+176+12, uint32(1), libc.Uint64FromInt64(1))
			v12 = self + 176 + 8
			v11 = *(*uint32_t)(unsafe.Pointer(v12))
			*(*uint32_t)(unsafe.Pointer(v12)) = *(*uint32_t)(unsafe.Pointer(v12)) + 1
			*(*int8)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents + uintptr(v11))) = int8(libc.Int32FromUint8('\t'))
			goto _7
		_5:
			;
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 176)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fsize, self+176+12, uint32(1), libc.Uint64FromInt64(1))
			v12 = self + 176 + 8
			v11 = *(*uint32_t)(unsafe.Pointer(v12))
			*(*uint32_t)(unsafe.Pointer(v12)) = *(*uint32_t)(unsafe.Pointer(v12)) + 1
			*(*int8)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents + uintptr(v11))) = int8(libc.Int32FromUint8('\000'))
			goto _7
		_6:
			;
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 176)).Fcontents = _array__splice(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents, self+176+8, self+176+12, libc.Uint64FromInt64(1), (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fsize, uint32(0), uint32((*Stream)(unsafe.Pointer(stream)).Fnext_size), (*Stream)(unsafe.Pointer(stream)).Finput)
			goto _7
		_7:
			;
			prev_position = (*Stream)(unsafe.Pointer(stream)).Finput + uintptr((*Stream)(unsafe.Pointer(stream)).Fnext_size)
		} else {
			if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('\\') {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 176)).Fcontents = _array__splice(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+176)).Fcontents, self+176+8, self+176+12, libc.Uint64FromInt64(1), (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+176)).Fsize, uint32(0), libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput)-int64(prev_position)), prev_position)
				prev_position = (*Stream)(unsafe.Pointer(stream)).Finput + uintptr(1)
				is_escaped = libc.BoolUint8(1 != 0)
			} else {
				if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('"') {
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 176)).Fcontents = _array__splice(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+176)).Fcontents, self+176+8, self+176+12, libc.Uint64FromInt64(1), (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+176)).Fsize, uint32(0), libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput)-int64(prev_position)), prev_position)
					stream_advance(tls, stream)
					return int32(TSQueryErrorNone)
				} else {
					if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('\n') {
						stream_reset(tls, stream, string_start)
						return int32(TSQueryErrorSyntax)
					}
				}
			}
		}
		if !(stream_advance(tls, stream) != 0) {
			stream_reset(tls, stream, string_start)
			return int32(TSQueryErrorSyntax)
		}
		goto _1
	_1:
	}
	return r
}

func ts_query__parse_predicate(tls *libc.TLS, self uintptr, stream uintptr) (r TSQueryError1) {
	var capture_id int32
	var capture_length, length, symbol_length, v1 uint32_t
	var capture_name, predicate_name, symbol_start, v2 uintptr
	var e TSQueryError1
	var id, query_id, query_id1 uint16_t
	_, _, _, _, _, _, _, _, _, _, _, _, _ = capture_id, capture_length, capture_name, e, id, length, predicate_name, query_id, query_id1, symbol_length, symbol_start, v1, v2
	if !(stream_is_ident_start(tls, stream) != 0) {
		return int32(TSQueryErrorSyntax)
	}
	predicate_name = (*Stream)(unsafe.Pointer(stream)).Finput
	stream_scan_identifier(tls, stream)
	if (*Stream)(unsafe.Pointer(stream)).Fnext != int32('?') && (*Stream)(unsafe.Pointer(stream)).Fnext != int32('!') {
		return int32(TSQueryErrorSyntax)
	}
	stream_advance(tls, stream)
	length = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(predicate_name))
	id = symbol_table_insert_name(tls, self+32, predicate_name, length)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 112)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+112)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+112)).Fsize, self+112+12, uint32(1), libc.Uint64FromInt64(8))
	v2 = self + 112 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*TSQueryPredicateStep)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+112)).Fcontents + uintptr(v1)*8)) = TSQueryPredicateStep{
		Ftype1:    int32(TSQueryPredicateStepTypeString),
		Fvalue_id: uint32(id),
	}
	stream_skip_whitespace(tls, stream)
	for {
		if (*Stream)(unsafe.Pointer(stream)).Fnext == int32(')') {
			stream_advance(tls, stream)
			stream_skip_whitespace(tls, stream)
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 112)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+112)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+112)).Fsize, self+112+12, uint32(1), libc.Uint64FromInt64(8))
			v2 = self + 112 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*TSQueryPredicateStep)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+112)).Fcontents + uintptr(v1)*8)) = TSQueryPredicateStep{}
			break
		} else {
			if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('@') {
				stream_advance(tls, stream)
				if !(stream_is_ident_start(tls, stream) != 0) {
					return int32(TSQueryErrorSyntax)
				}
				capture_name = (*Stream)(unsafe.Pointer(stream)).Finput
				stream_scan_identifier(tls, stream)
				capture_length = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(capture_name))
				capture_id = symbol_table_id_for_name(tls, self, capture_name, capture_length)
				if capture_id == -int32(1) {
					stream_reset(tls, stream, capture_name)
					return int32(TSQueryErrorCapture)
				}
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 112)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+112)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+112)).Fsize, self+112+12, uint32(1), libc.Uint64FromInt64(8))
				v2 = self + 112 + 8
				v1 = *(*uint32_t)(unsafe.Pointer(v2))
				*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
				*(*TSQueryPredicateStep)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+112)).Fcontents + uintptr(v1)*8)) = TSQueryPredicateStep{
					Ftype1:    int32(TSQueryPredicateStepTypeCapture),
					Fvalue_id: libc.Uint32FromInt32(capture_id),
				}
			} else {
				if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('"') {
					e = ts_query__parse_string_literal(tls, self, stream)
					if e != 0 {
						return e
					}
					query_id = symbol_table_insert_name(tls, self+32, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fcontents, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fsize)
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 112)).Fcontents = _array__grow(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+112)).Fcontents, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+112)).Fsize, self+112+12, uint32(1), libc.Uint64FromInt64(8))
					v2 = self + 112 + 8
					v1 = *(*uint32_t)(unsafe.Pointer(v2))
					*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
					*(*TSQueryPredicateStep)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+112)).Fcontents + uintptr(v1)*8)) = TSQueryPredicateStep{
						Ftype1:    int32(TSQueryPredicateStepTypeString),
						Fvalue_id: uint32(query_id),
					}
				} else {
					if stream_is_ident_start(tls, stream) != 0 {
						symbol_start = (*Stream)(unsafe.Pointer(stream)).Finput
						stream_scan_identifier(tls, stream)
						symbol_length = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(symbol_start))
						query_id1 = symbol_table_insert_name(tls, self+32, symbol_start, symbol_length)
						(*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self + 112)).Fcontents = _array__grow(tls, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+112)).Fcontents, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+112)).Fsize, self+112+12, uint32(1), libc.Uint64FromInt64(8))
						v2 = self + 112 + 8
						v1 = *(*uint32_t)(unsafe.Pointer(v2))
						*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
						*(*TSQueryPredicateStep)(unsafe.Pointer((*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+112)).Fcontents + uintptr(v1)*8)) = TSQueryPredicateStep{
							Ftype1:    int32(TSQueryPredicateStepTypeString),
							Fvalue_id: uint32(query_id1),
						}
					} else {
						return int32(TSQueryErrorSyntax)
					}
				}
			}
		}
		stream_skip_whitespace(tls, stream)
		goto _3
	_3:
	}
	return 0
}

func ts_query__parse_pattern(tls *libc.TLS, self uintptr, stream uintptr, depth uint32_t, is_immediate uint8, capture_quantifiers uintptr) (r TSQueryError1) {
	bp := tls.Alloc(128)
	defer tls.Free(128)
	var alternative_step, capture_name, end_step, field_name, field_name1, last_child_step, missing_node_name, node_name, start_step, step, step1, step2, step3, string_start, string_start1, subtype_node_name, subtypes, v3 uintptr
	var capture_id, last_child_step_index, negated_field_count, step_index1 uint16_t
	var child_is_immediate, child_is_immediate1, is_missing, subtype_is_valid uint8
	var e, e1, e2, e3, e4, e5, e6 TSQueryError1
	var field_id, field_id1 TSFieldId
	var i uint32
	var i1, length, length1, length2, length3, length4, missing_node_length, next_step_index, start_index, starting_step_index, step_index, step_index2, step_index3, v2 uint32_t
	var quantifier TSQuantifier1
	var symbol, symbol1 TSSymbol
	var v1 bool
	var _ /* branch_capture_quantifiers at bp+16 */ CaptureQuantifiers
	var _ /* branch_step_indices at bp+0 */ struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	var _ /* child_capture_quantifiers at bp+32 */ CaptureQuantifiers
	var _ /* child_capture_quantifiers at bp+72 */ CaptureQuantifiers
	var _ /* field_capture_quantifiers at bp+88 */ CaptureQuantifiers
	var _ /* negated_field_ids at bp+52 */ [8]TSFieldId
	var _ /* repeat_step at bp+104 */ QueryStep
	var _ /* subtype_length at bp+48 */ uint32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alternative_step, capture_id, capture_name, child_is_immediate, child_is_immediate1, e, e1, e2, e3, e4, e5, e6, end_step, field_id, field_id1, field_name, field_name1, i, i1, is_missing, last_child_step, last_child_step_index, length, length1, length2, length3, length4, missing_node_length, missing_node_name, negated_field_count, next_step_index, node_name, quantifier, start_index, start_step, starting_step_index, step, step1, step2, step3, step_index, step_index1, step_index2, step_index3, string_start, string_start1, subtype_is_valid, subtype_node_name, subtypes, symbol, symbol1, v1, v2, v3
	if (*Stream)(unsafe.Pointer(stream)).Fnext == 0 {
		return int32(TSQueryErrorSyntax)
	}
	if (*Stream)(unsafe.Pointer(stream)).Fnext == int32(')') || (*Stream)(unsafe.Pointer(stream)).Fnext == int32(']') {
		return PARENT_DONE
	}
	starting_step_index = (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize
	if v1 = (*TSQuery)(unsafe.Pointer(self)).Fstep_offsets.Fsize == uint32(0); !v1 {
		_ = libc.Uint64FromInt64(4)
		{
			if !((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+6671, __ccgo_ts+3070, int32(2244), uintptr(unsafe.Pointer(&__func__45)))
			}
		}
	}
	if v1 || uint32((*StepOffset)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+144)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+144)).Fsize-uint32(1))*8)).Fstep_index) != starting_step_index {
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 144)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+144)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+144)).Fsize, self+144+12, uint32(1), libc.Uint64FromInt64(8))
		v3 = self + 144 + 8
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
		*(*StepOffset)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+144)).Fcontents + uintptr(v2)*8)) = StepOffset{
			Fbyte_offset: stream_offset(tls, stream),
			Fstep_index:  uint16(starting_step_index),
		}
	}
	if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('[') {
		stream_advance(tls, stream)
		stream_skip_whitespace(tls, stream)
		*(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp)) = struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		}{}
		*(*CaptureQuantifiers)(unsafe.Pointer(bp + 16)) = capture_quantifiers_new(tls)
		for {
			start_index = (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize
			e = ts_query__parse_pattern(tls, self, stream, depth, is_immediate, bp+16)
			if e == PARENT_DONE {
				if (*Stream)(unsafe.Pointer(stream)).Fnext == int32(']') && (*(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp))).Fsize > uint32(0) {
					stream_advance(tls, stream)
					break
				}
				e = int32(TSQueryErrorSyntax)
			}
			if e != 0 {
				capture_quantifiers_delete(tls, bp+16)
				if (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp)).Fcontents != 0 {
					(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp)).Fcontents)
				}
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp)).Fcontents = libc.UintptrFromInt32(0)
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp)).Fsize = uint32(0)
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp)).Fcapacity = uint32(0)
				return e
			}
			if start_index == starting_step_index {
				capture_quantifiers_replace(tls, capture_quantifiers, bp+16)
			} else {
				capture_quantifiers_join_all(tls, capture_quantifiers, bp+16)
			}
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fsize, bp+12, uint32(1), libc.Uint64FromInt64(4))
			v3 = bp + 8
			v2 = *(*uint32_t)(unsafe.Pointer(v3))
			*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
			*(*uint32_t)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents + uintptr(v2)*4)) = start_index
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
			v3 = self + 80 + 8
			v2 = *(*uint32_t)(unsafe.Pointer(v3))
			*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
			*(*QueryStep)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(v2)*20)) = query_step__new(tls, uint16(0), uint16(depth), libc.BoolUint8(0 != 0))
			capture_quantifiers_clear(tls, bp+16)
			goto _4
		_4:
		}
		v3 = self + 80 + 8
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) - 1
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		_ = *(*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(v2)*20))
		i = uint32(0)
		for {
			if !(i < (*(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp))).Fsize-uint32(1)) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6745, __ccgo_ts+3070, int32(2298), uintptr(unsafe.Pointer(&__func__45)))
				}
			}
			step_index = *(*uint32_t)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents + uintptr(i)*4))
			_ = libc.Uint64FromInt64(4)
			{
				if !(i+libc.Uint32FromInt32(1) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6790, __ccgo_ts+3070, int32(2299), uintptr(unsafe.Pointer(&__func__45)))
				}
			}
			next_step_index = *(*uint32_t)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents + uintptr(i+uint32(1))*4))
			_ = libc.Uint64FromInt64(4)
			{
				if !(step_index < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(2300), uintptr(unsafe.Pointer(&__func__45)))
				}
			}
			start_step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index)*20
			_ = libc.Uint64FromInt64(4)
			{
				if !(next_step_index-libc.Uint32FromInt32(1) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6839, __ccgo_ts+3070, int32(2301), uintptr(unsafe.Pointer(&__func__45)))
				}
			}
			end_step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(next_step_index-uint32(1))*20
			(*QueryStep)(unsafe.Pointer(start_step)).Falternative_index = uint16(next_step_index)
			(*QueryStep)(unsafe.Pointer(end_step)).Falternative_index = uint16((*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize)
			libc.SetBitFieldPtr8Uint8(end_step+18, libc.BoolUint8(1 != 0), 4, 0x10)
			goto _11
		_11:
			;
			i = i + 1
		}
		capture_quantifiers_delete(tls, bp+16)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp)).Fcapacity = uint32(0)
	} else {
		if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('(') {
			stream_advance(tls, stream)
			stream_skip_whitespace(tls, stream)
			if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('(') || (*Stream)(unsafe.Pointer(stream)).Fnext == int32('"') || (*Stream)(unsafe.Pointer(stream)).Fnext == int32('[') {
				child_is_immediate = is_immediate
				*(*CaptureQuantifiers)(unsafe.Pointer(bp + 32)) = capture_quantifiers_new(tls)
				for {
					if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('.') {
						child_is_immediate = libc.BoolUint8(1 != 0)
						stream_advance(tls, stream)
						stream_skip_whitespace(tls, stream)
					}
					e1 = ts_query__parse_pattern(tls, self, stream, depth, child_is_immediate, bp+32)
					if e1 == PARENT_DONE {
						if (*Stream)(unsafe.Pointer(stream)).Fnext == int32(')') {
							stream_advance(tls, stream)
							break
						}
						e1 = int32(TSQueryErrorSyntax)
					}
					if e1 != 0 {
						capture_quantifiers_delete(tls, bp+32)
						return e1
					}
					capture_quantifiers_add_all(tls, capture_quantifiers, bp+32)
					capture_quantifiers_clear(tls, bp+32)
					child_is_immediate = libc.BoolUint8(0 != 0)
					goto _12
				_12:
				}
				capture_quantifiers_delete(tls, bp+32)
			} else {
				if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('.') || (*Stream)(unsafe.Pointer(stream)).Fnext == int32('#') {
					stream_advance(tls, stream)
					return ts_query__parse_predicate(tls, self, stream)
				} else {
					is_missing = libc.BoolUint8(0 != 0)
					node_name = (*Stream)(unsafe.Pointer(stream)).Finput
					if stream_is_ident_start(tls, stream) != 0 {
						stream_scan_identifier(tls, stream)
						length = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(node_name))
						if length == uint32(1) && int32(*(*int8)(unsafe.Pointer(node_name))) == int32('_') {
							symbol = WILDCARD_SYMBOL
						} else {
							if !(libc.Xstrncmp(tls, node_name, __ccgo_ts+6894, uint64(length)) != 0) {
								is_missing = libc.BoolUint8(1 != 0)
								stream_skip_whitespace(tls, stream)
								if stream_is_ident_start(tls, stream) != 0 {
									missing_node_name = (*Stream)(unsafe.Pointer(stream)).Finput
									stream_scan_identifier(tls, stream)
									missing_node_length = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(missing_node_name))
									symbol = ts_language_symbol_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, missing_node_name, missing_node_length, libc.BoolUint8(1 != 0))
									if !(symbol != 0) {
										stream_reset(tls, stream, missing_node_name)
										return int32(TSQueryErrorNodeType)
									}
								} else {
									if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('"') {
										string_start = (*Stream)(unsafe.Pointer(stream)).Finput
										e2 = ts_query__parse_string_literal(tls, self, stream)
										if e2 != 0 {
											return e2
										}
										symbol = ts_language_symbol_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fcontents, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fsize, libc.BoolUint8(0 != 0))
										if !(symbol != 0) {
											stream_reset(tls, stream, string_start+uintptr(1))
											return int32(TSQueryErrorNodeType)
										}
									} else {
										if (*Stream)(unsafe.Pointer(stream)).Fnext == int32(')') {
											symbol = WILDCARD_SYMBOL
										} else {
											stream_reset(tls, stream, (*Stream)(unsafe.Pointer(stream)).Finput)
											return int32(TSQueryErrorSyntax)
										}
									}
								}
							} else {
								symbol = ts_language_symbol_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, node_name, length, libc.BoolUint8(1 != 0))
								if !(symbol != 0) {
									stream_reset(tls, stream, node_name)
									return int32(TSQueryErrorNodeType)
								}
							}
						}
					} else {
						return int32(TSQueryErrorSyntax)
					}
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fcontents, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
					v3 = self + 80 + 8
					v2 = *(*uint32_t)(unsafe.Pointer(v3))
					*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
					*(*QueryStep)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fcontents + uintptr(v2)*20)) = query_step__new(tls, symbol, uint16(depth), is_immediate)
					_ = libc.Uint64FromInt64(4)
					{
						if !((*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+80)).Fsize-libc.Uint32FromInt32(1) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+80)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+6902, __ccgo_ts+3070, int32(2441), uintptr(unsafe.Pointer(&__func__45)))
						}
					}
					step = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fcontents + uintptr((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize-uint32(1))*20
					if ts_language_symbol_metadata(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, symbol).Fsupertype != 0 {
						(*QueryStep)(unsafe.Pointer(step)).Fsupertype_symbol = (*QueryStep)(unsafe.Pointer(step)).Fsymbol
						(*QueryStep)(unsafe.Pointer(step)).Fsymbol = WILDCARD_SYMBOL
					}
					if is_missing != 0 {
						libc.SetBitFieldPtr8Uint8(step+19, libc.BoolUint8(1 != 0), 1, 0x2)
					}
					if libc.Int32FromUint16(symbol) == libc.Int32FromUint16(WILDCARD_SYMBOL) {
						libc.SetBitFieldPtr8Uint8(step+18, libc.BoolUint8(1 != 0), 0, 0x1)
					}
					if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('/') {
						if !((*QueryStep)(unsafe.Pointer(step)).Fsupertype_symbol != 0) {
							stream_reset(tls, stream, node_name-uintptr(1))
							return int32(TSQueryErrorStructure)
						}
						stream_advance(tls, stream)
						subtype_node_name = (*Stream)(unsafe.Pointer(stream)).Finput
						if stream_is_ident_start(tls, stream) != 0 {
							stream_scan_identifier(tls, stream)
							length1 = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(subtype_node_name))
							(*QueryStep)(unsafe.Pointer(step)).Fsymbol = ts_language_symbol_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, subtype_node_name, length1, libc.BoolUint8(1 != 0))
						} else {
							if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('"') {
								e3 = ts_query__parse_string_literal(tls, self, stream)
								if e3 != 0 {
									return e3
								}
								(*QueryStep)(unsafe.Pointer(step)).Fsymbol = ts_language_symbol_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fcontents, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fsize, libc.BoolUint8(0 != 0))
							} else {
								return int32(TSQueryErrorSyntax)
							}
						}
						if !((*QueryStep)(unsafe.Pointer(step)).Fsymbol != 0) {
							stream_reset(tls, stream, subtype_node_name)
							return int32(TSQueryErrorNodeType)
						}
						if (*TSLanguage)(unsafe.Pointer((*TSQuery)(unsafe.Pointer(self)).Flanguage)).Fabi_version >= uint32(15) {
							subtypes = ts_language_subtypes(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*QueryStep)(unsafe.Pointer(step)).Fsupertype_symbol, bp+48)
							subtype_is_valid = libc.BoolUint8(0 != 0)
							i1 = uint32(0)
							for {
								if !(i1 < *(*uint32_t)(unsafe.Pointer(bp + 48))) {
									break
								}
								if libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(subtypes + uintptr(i1)*2))) == libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) {
									subtype_is_valid = libc.BoolUint8(1 != 0)
									break
								}
								goto _15
							_15:
								;
								i1 = i1 + 1
							}
							if !(subtype_is_valid != 0) {
								stream_reset(tls, stream, node_name-uintptr(1))
								return int32(TSQueryErrorStructure)
							}
						}
					}
					stream_skip_whitespace(tls, stream)
					child_is_immediate1 = libc.BoolUint8(0 != 0)
					last_child_step_index = uint16(0)
					negated_field_count = uint16(0)
					*(*CaptureQuantifiers)(unsafe.Pointer(bp + 72)) = capture_quantifiers_new(tls)
					for {
						if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('!') {
							stream_advance(tls, stream)
							stream_skip_whitespace(tls, stream)
							if !(stream_is_ident_start(tls, stream) != 0) {
								capture_quantifiers_delete(tls, bp+72)
								return int32(TSQueryErrorSyntax)
							}
							field_name = (*Stream)(unsafe.Pointer(stream)).Finput
							stream_scan_identifier(tls, stream)
							length2 = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(field_name))
							stream_skip_whitespace(tls, stream)
							field_id = ts_language_field_id_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, field_name, length2)
							if !(field_id != 0) {
								(*Stream)(unsafe.Pointer(stream)).Finput = field_name
								capture_quantifiers_delete(tls, bp+72)
								return int32(TSQueryErrorField)
							}
							if libc.Int32FromUint16(negated_field_count) < int32(8) {
								(*(*[8]TSFieldId)(unsafe.Pointer(bp + 52)))[negated_field_count] = field_id
								negated_field_count = negated_field_count + 1
							}
							goto _16
						}
						if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('.') {
							child_is_immediate1 = libc.BoolUint8(1 != 0)
							stream_advance(tls, stream)
							stream_skip_whitespace(tls, stream)
						}
						step_index1 = uint16((*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize)
						e4 = ts_query__parse_pattern(tls, self, stream, depth+uint32(1), child_is_immediate1, bp+72)
						if uint32(step_index1) == (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize {
							step_index1 = step_index1 - 1
						}
						if e4 == PARENT_DONE {
							if (*Stream)(unsafe.Pointer(stream)).Fnext == int32(')') {
								if child_is_immediate1 != 0 {
									if libc.Int32FromUint16(last_child_step_index) == 0 {
										capture_quantifiers_delete(tls, bp+72)
										return int32(TSQueryErrorSyntax)
									}
									_ = libc.Uint64FromInt64(4)
									{
										if !(uint32(last_child_step_index) < (*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(self+80)).Fsize) {
											libc.X__assert_fail(tls, __ccgo_ts+6962, __ccgo_ts+3070, int32(2585), uintptr(unsafe.Pointer(&__func__45)))
										}
									}
									last_child_step = (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(self+80)).Fcontents + uintptr(last_child_step_index)*20
									libc.SetBitFieldPtr8Uint8(last_child_step+18, libc.BoolUint8(1 != 0), 2, 0x4)
									if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(last_child_step)).Falternative_index) != libc.Int32FromUint16(NONE) && uint32((*QueryStep)(unsafe.Pointer(last_child_step)).Falternative_index) < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize {
										_ = libc.Uint64FromInt64(4)
										{
											if !(uint32((*QueryStep)(unsafe.Pointer(last_child_step)).Falternative_index) < (*struct {
												Fcontents uintptr
												Fsize     uint32_t
												Fcapacity uint32_t
											})(unsafe.Pointer(self+80)).Fsize) {
												libc.X__assert_fail(tls, __ccgo_ts+7019, __ccgo_ts+3070, int32(2591), uintptr(unsafe.Pointer(&__func__45)))
											}
										}
										alternative_step = (*struct {
											Fcontents uintptr
											Fsize     uint32_t
											Fcapacity uint32_t
										})(unsafe.Pointer(self+80)).Fcontents + uintptr((*QueryStep)(unsafe.Pointer(last_child_step)).Falternative_index)*20
										libc.SetBitFieldPtr8Uint8(alternative_step+18, libc.BoolUint8(1 != 0), 2, 0x4)
										for libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(alternative_step)).Falternative_index) != libc.Int32FromUint16(NONE) && uint32((*QueryStep)(unsafe.Pointer(alternative_step)).Falternative_index) < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize {
											_ = libc.Uint64FromInt64(4)
											{
												if !(uint32((*QueryStep)(unsafe.Pointer(alternative_step)).Falternative_index) < (*struct {
													Fcontents uintptr
													Fsize     uint32_t
													Fcapacity uint32_t
												})(unsafe.Pointer(self+80)).Fsize) {
													libc.X__assert_fail(tls, __ccgo_ts+7089, __ccgo_ts+3070, int32(2597), uintptr(unsafe.Pointer(&__func__45)))
												}
											}
											alternative_step = (*struct {
												Fcontents uintptr
												Fsize     uint32_t
												Fcapacity uint32_t
											})(unsafe.Pointer(self+80)).Fcontents + uintptr((*QueryStep)(unsafe.Pointer(alternative_step)).Falternative_index)*20
											libc.SetBitFieldPtr8Uint8(alternative_step+18, libc.BoolUint8(1 != 0), 2, 0x4)
										}
									}
								}
								if negated_field_count != 0 {
									ts_query__add_negated_fields(tls, self, uint16(starting_step_index), bp+52, negated_field_count)
								}
								stream_advance(tls, stream)
								break
							}
							e4 = int32(TSQueryErrorSyntax)
						}
						if e4 != 0 {
							capture_quantifiers_delete(tls, bp+72)
							return e4
						}
						capture_quantifiers_add_all(tls, capture_quantifiers, bp+72)
						last_child_step_index = step_index1
						child_is_immediate1 = libc.BoolUint8(0 != 0)
						capture_quantifiers_clear(tls, bp+72)
						goto _16
					_16:
					}
					capture_quantifiers_delete(tls, bp+72)
				}
			}
		} else {
			if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('_') {
				stream_advance(tls, stream)
				stream_skip_whitespace(tls, stream)
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
				v3 = self + 80 + 8
				v2 = *(*uint32_t)(unsafe.Pointer(v3))
				*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
				*(*QueryStep)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fcontents + uintptr(v2)*20)) = query_step__new(tls, WILDCARD_SYMBOL, uint16(depth), is_immediate)
			} else {
				if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('"') {
					string_start1 = (*Stream)(unsafe.Pointer(stream)).Finput
					e5 = ts_query__parse_string_literal(tls, self, stream)
					if e5 != 0 {
						return e5
					}
					symbol1 = ts_language_symbol_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fcontents, (*TSQuery)(unsafe.Pointer(self)).Fstring_buffer.Fsize, libc.BoolUint8(0 != 0))
					if !(symbol1 != 0) {
						stream_reset(tls, stream, string_start1+uintptr(1))
						return int32(TSQueryErrorNodeType)
					}
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fcontents, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
					v3 = self + 80 + 8
					v2 = *(*uint32_t)(unsafe.Pointer(v3))
					*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
					*(*QueryStep)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fcontents + uintptr(v2)*20)) = query_step__new(tls, symbol1, uint16(depth), is_immediate)
				} else {
					if stream_is_ident_start(tls, stream) != 0 {
						field_name1 = (*Stream)(unsafe.Pointer(stream)).Finput
						stream_scan_identifier(tls, stream)
						length3 = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(field_name1))
						stream_skip_whitespace(tls, stream)
						if (*Stream)(unsafe.Pointer(stream)).Fnext != int32(':') {
							stream_reset(tls, stream, field_name1)
							return int32(TSQueryErrorSyntax)
						}
						stream_advance(tls, stream)
						stream_skip_whitespace(tls, stream)
						*(*CaptureQuantifiers)(unsafe.Pointer(bp + 88)) = capture_quantifiers_new(tls)
						e6 = ts_query__parse_pattern(tls, self, stream, depth, is_immediate, bp+88)
						if e6 != 0 {
							capture_quantifiers_delete(tls, bp+88)
							if e6 == PARENT_DONE {
								e6 = int32(TSQueryErrorSyntax)
							}
							return e6
						}
						field_id1 = ts_language_field_id_for_name(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage, field_name1, length3)
						if !(field_id1 != 0) {
							(*Stream)(unsafe.Pointer(stream)).Finput = field_name1
							return int32(TSQueryErrorField)
						}
						step_index2 = starting_step_index
						_ = libc.Uint64FromInt64(4)
						{
							if !(step_index2 < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+80)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(2703), uintptr(unsafe.Pointer(&__func__45)))
							}
						}
						step1 = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index2)*20
						for {
							(*QueryStep)(unsafe.Pointer(step1)).Ffield = field_id1
							if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step1)).Falternative_index) != libc.Int32FromUint16(NONE) && uint32((*QueryStep)(unsafe.Pointer(step1)).Falternative_index) > step_index2 && uint32((*QueryStep)(unsafe.Pointer(step1)).Falternative_index) < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize {
								step_index2 = uint32((*QueryStep)(unsafe.Pointer(step1)).Falternative_index)
								_ = libc.Uint64FromInt64(4)
								{
									if !(step_index2 < (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(self+80)).Fsize) {
										libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(2712), uintptr(unsafe.Pointer(&__func__45)))
									}
								}
								step1 = (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index2)*20
							} else {
								break
							}
							goto _21
						_21:
						}
						capture_quantifiers_add_all(tls, capture_quantifiers, bp+88)
						capture_quantifiers_delete(tls, bp+88)
					} else {
						return int32(TSQueryErrorSyntax)
					}
				}
			}
		}
	}
	stream_skip_whitespace(tls, stream)
	quantifier = int32(TSQuantifierOne)
	for {
		if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('+') {
			quantifier = quantifier_join(tls, int32(TSQuantifierOneOrMore), quantifier)
			stream_advance(tls, stream)
			stream_skip_whitespace(tls, stream)
		} else {
			if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('*') {
				quantifier = quantifier_join(tls, int32(TSQuantifierZeroOrMore), quantifier)
				stream_advance(tls, stream)
				stream_skip_whitespace(tls, stream)
			} else {
				if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('?') {
					quantifier = quantifier_join(tls, int32(TSQuantifierZeroOrOne), quantifier)
					stream_advance(tls, stream)
					stream_skip_whitespace(tls, stream)
				} else {
					if (*Stream)(unsafe.Pointer(stream)).Fnext == int32('@') {
						stream_advance(tls, stream)
						if !(stream_is_ident_start(tls, stream) != 0) {
							return int32(TSQueryErrorSyntax)
						}
						capture_name = (*Stream)(unsafe.Pointer(stream)).Finput
						stream_scan_identifier(tls, stream)
						length4 = libc.Uint32FromInt64(int64((*Stream)(unsafe.Pointer(stream)).Finput) - int64(capture_name))
						stream_skip_whitespace(tls, stream)
						capture_id = symbol_table_insert_name(tls, self, capture_name, length4)
						capture_quantifiers_add_for_id(tls, capture_quantifiers, capture_id, int32(TSQuantifierOne))
						step_index3 = starting_step_index
						for {
							_ = libc.Uint64FromInt64(4)
							{
								if !(step_index3 < (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+80)).Fsize) {
									libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(2776), uintptr(unsafe.Pointer(&__func__45)))
								}
							}
							step2 = (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index3)*20
							query_step__add_capture(tls, step2, capture_id)
							if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step2)).Falternative_index) != libc.Int32FromUint16(NONE) && uint32((*QueryStep)(unsafe.Pointer(step2)).Falternative_index) > step_index3 && uint32((*QueryStep)(unsafe.Pointer(step2)).Falternative_index) < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize {
								step_index3 = uint32((*QueryStep)(unsafe.Pointer(step2)).Falternative_index)
							} else {
								break
							}
							goto _23
						_23:
						}
					} else {
						break
					}
				}
			}
		}
		goto _22
	_22:
	}
	switch quantifier {
	case int32(TSQuantifierOneOrMore):
		*(*QueryStep)(unsafe.Pointer(bp + 104)) = query_step__new(tls, WILDCARD_SYMBOL, uint16(depth), libc.BoolUint8(0 != 0))
		(*(*QueryStep)(unsafe.Pointer(bp + 104))).Falternative_index = uint16(starting_step_index)
		libc.SetBitFieldPtr8Uint8(bp+104+18, libc.BoolUint8(1 != 0), 3, 0x8)
		libc.SetBitFieldPtr8Uint8(bp+104+18, libc.BoolUint8(1 != 0), 5, 0x20)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
		v3 = self + 80 + 8
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
		*(*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(v2)*20)) = *(*QueryStep)(unsafe.Pointer(bp + 104))
	case int32(TSQuantifierZeroOrMore):
		*(*QueryStep)(unsafe.Pointer(bp + 104)) = query_step__new(tls, WILDCARD_SYMBOL, uint16(depth), libc.BoolUint8(0 != 0))
		(*(*QueryStep)(unsafe.Pointer(bp + 104))).Falternative_index = uint16(starting_step_index)
		libc.SetBitFieldPtr8Uint8(bp+104+18, libc.BoolUint8(1 != 0), 3, 0x8)
		libc.SetBitFieldPtr8Uint8(bp+104+18, libc.BoolUint8(1 != 0), 5, 0x20)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
		v3 = self + 80 + 8
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
		*(*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(v2)*20)) = *(*QueryStep)(unsafe.Pointer(bp + 104))
		_ = libc.Uint64FromInt64(4)
		{
			if !(starting_step_index < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7160, __ccgo_ts+3070, int32(2816), uintptr(unsafe.Pointer(&__func__45)))
			}
		}
		step3 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(starting_step_index)*20
		for libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Falternative_index) != libc.Int32FromUint16(NONE) && uint32((*QueryStep)(unsafe.Pointer(step3)).Falternative_index) < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize-uint32(1) {
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32((*QueryStep)(unsafe.Pointer(step3)).Falternative_index) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6154, __ccgo_ts+3070, int32(2818), uintptr(unsafe.Pointer(&__func__45)))
				}
			}
			step3 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr((*QueryStep)(unsafe.Pointer(step3)).Falternative_index)*20
		}
		(*QueryStep)(unsafe.Pointer(step3)).Falternative_index = uint16((*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize)
	case int32(TSQuantifierZeroOrOne):
		_ = libc.Uint64FromInt64(4)
		{
			if !(starting_step_index < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7160, __ccgo_ts+3070, int32(2823), uintptr(unsafe.Pointer(&__func__45)))
			}
		}
		step3 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(starting_step_index)*20
		for libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Falternative_index) != libc.Int32FromUint16(NONE) && uint32((*QueryStep)(unsafe.Pointer(step3)).Falternative_index) < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize {
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32((*QueryStep)(unsafe.Pointer(step3)).Falternative_index) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+6154, __ccgo_ts+3070, int32(2825), uintptr(unsafe.Pointer(&__func__45)))
				}
			}
			step3 = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr((*QueryStep)(unsafe.Pointer(step3)).Falternative_index)*20
		}
		(*QueryStep)(unsafe.Pointer(step3)).Falternative_index = uint16((*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize)
	default:
		break
	}
	capture_quantifiers_mul(tls, capture_quantifiers, quantifier)
	return 0
}

var __func__45 = [24]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', '_', 'p', 'a', 'r', 's', 'e', '_', 'p', 'a', 't', 't', 'e', 'r', 'n'}

func ts_query_new(tls *libc.TLS, language uintptr, source uintptr, source_len uint32_t, error_offset uintptr, error_type uintptr) (r uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var child_step, pattern, s, second_step, self, step, target, v2 uintptr
	var copy1 QueryStep
	var copy_idx, i, pat_end, pat_start, pattern_index, start_depth, start_predicate_step_index, start_step_index, step_index, target_idx, v1 uint32_t
	var is_rooted uint8
	var target_alt_index, target_depth, wildcard_root_alternative_index uint16_t
	var _ /* capture_quantifiers at bp+32 */ CaptureQuantifiers
	var _ /* redirect at bp+48 */ QueryStep
	var _ /* stream at bp+0 */ Stream
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = child_step, copy1, copy_idx, i, is_rooted, pat_end, pat_start, pattern, pattern_index, s, second_step, self, start_depth, start_predicate_step_index, start_step_index, step, step_index, target, target_alt_index, target_depth, target_idx, wildcard_root_alternative_index, v1, v2
	if !(language != 0) || (*TSLanguage)(unsafe.Pointer(language)).Fabi_version > uint32(15) || (*TSLanguage)(unsafe.Pointer(language)).Fabi_version < uint32(13) {
		*(*TSQueryError1)(unsafe.Pointer(error_type)) = int32(TSQueryErrorLanguage)
		return libc.UintptrFromInt32(0)
	}
	self = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(224))
	*(*TSQuery)(unsafe.Pointer(self)) = TSQuery{
		Fcaptures:         symbol_table_new(tls),
		Fpredicate_values: symbol_table_new(tls),
		Flanguage:         ts_language_copy(tls, language),
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 160)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fsize, self+160+12, uint32(1), libc.Uint64FromInt64(2))
	v2 = self + 160 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*TSFieldId)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+160)).Fcontents + uintptr(v1)*2)) = libc.Uint16FromInt32(libc.Int32FromInt32(0))
	*(*Stream)(unsafe.Pointer(bp)) = stream_new(tls, source, source_len)
	stream_skip_whitespace(tls, bp)
	for (*(*Stream)(unsafe.Pointer(bp))).Finput < (*(*Stream)(unsafe.Pointer(bp))).Fend {
		pattern_index = (*TSQuery)(unsafe.Pointer(self)).Fpatterns.Fsize
		start_step_index = (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize
		start_predicate_step_index = (*TSQuery)(unsafe.Pointer(self)).Fpredicate_steps.Fsize
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 128)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fsize, self+128+12, uint32(1), libc.Uint64FromInt64(28))
		v2 = self + 128 + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*QueryPattern)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fcontents + uintptr(v1)*28)) = QueryPattern{
			Fsteps: Slice{
				Foffset: start_step_index,
			},
			Fpredicate_steps: Slice{
				Foffset: start_predicate_step_index,
			},
			Fstart_byte: stream_offset(tls, bp),
		}
		*(*CaptureQuantifiers)(unsafe.Pointer(bp + 32)) = capture_quantifiers_new(tls)
		*(*TSQueryError1)(unsafe.Pointer(error_type)) = ts_query__parse_pattern(tls, self, bp, uint32(0), libc.BoolUint8(0 != 0), bp+32)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
		v2 = self + 80 + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*QueryStep)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(v1)*20)) = query_step__new(tls, uint16(0), PATTERN_DONE_MARKER, libc.BoolUint8(0 != 0))
		_ = libc.Uint64FromInt64(4)
		{
			if !((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+128)).Fsize-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+128)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7215, __ccgo_ts+3070, int32(2890), uintptr(unsafe.Pointer(&__func__46)))
			}
		}
		pattern = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fcontents + uintptr((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fsize-uint32(1))*28
		(*QueryPattern)(unsafe.Pointer(pattern)).Fsteps.Flength = (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize - start_step_index
		(*QueryPattern)(unsafe.Pointer(pattern)).Fpredicate_steps.Flength = (*TSQuery)(unsafe.Pointer(self)).Fpredicate_steps.Fsize - start_predicate_step_index
		(*QueryPattern)(unsafe.Pointer(pattern)).Fend_byte = stream_offset(tls, bp)
		if *(*TSQueryError1)(unsafe.Pointer(error_type)) != 0 {
			if *(*TSQueryError1)(unsafe.Pointer(error_type)) == PARENT_DONE {
				*(*TSQueryError1)(unsafe.Pointer(error_type)) = int32(TSQueryErrorSyntax)
			}
			*(*uint32_t)(unsafe.Pointer(error_offset)) = stream_offset(tls, bp)
			capture_quantifiers_delete(tls, bp+32)
			ts_query_delete(tls, self)
			return libc.UintptrFromInt32(0)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 64)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+64)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+64)).Fsize, self+64+12, uint32(1), libc.Uint64FromInt64(16))
		v2 = self + 64 + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*CaptureQuantifiers)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+64)).Fcontents + uintptr(v1)*16)) = *(*CaptureQuantifiers)(unsafe.Pointer(bp + 32))
		wildcard_root_alternative_index = NONE
		for {
			_ = libc.Uint64FromInt64(4)
			{
				if !(start_step_index < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+7281, __ccgo_ts+3070, int32(2911), uintptr(unsafe.Pointer(&__func__46)))
				}
			}
			step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(start_step_index)*20
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) == libc.Int32FromUint16(WILDCARD_SYMBOL) && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth) == 0 && !((*QueryStep)(unsafe.Pointer(step)).Ffield != 0) {
				_ = libc.Uint64FromInt64(4)
				{
					if !(start_step_index+libc.Uint32FromInt32(1) < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+7333, __ccgo_ts+3070, int32(2918), uintptr(unsafe.Pointer(&__func__46)))
					}
				}
				second_step = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fcontents + uintptr(start_step_index+uint32(1))*20
				if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(second_step)).Fsymbol) != libc.Int32FromUint16(WILDCARD_SYMBOL) && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(second_step)).Fdepth) == int32(1) && !(int32(*(*uint8)(unsafe.Pointer(second_step + 18))&0x2>>1) != 0) {
					wildcard_root_alternative_index = (*QueryStep)(unsafe.Pointer(step)).Falternative_index
					start_step_index = start_step_index + uint32(1)
					step = second_step
				}
			}
			start_depth = uint32((*QueryStep)(unsafe.Pointer(step)).Fdepth)
			is_rooted = libc.BoolUint8(start_depth == uint32(0))
			step_index = start_step_index + uint32(1)
			for {
				if !(step_index < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(step_index < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+80)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(2933), uintptr(unsafe.Pointer(&__func__46)))
					}
				}
				child_step = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index)*20
				if int32(*(*uint8)(unsafe.Pointer(child_step + 18))&0x10>>4) != 0 {
					break
				}
				if uint32((*QueryStep)(unsafe.Pointer(child_step)).Fdepth) == start_depth {
					is_rooted = libc.BoolUint8(0 != 0)
					break
				}
				goto _10
			_10:
				;
				step_index = step_index + 1
			}
			ts_query__pattern_map_insert(tls, self, (*QueryStep)(unsafe.Pointer(step)).Fsymbol, PatternEntry{
				Fstep_index:    uint16(start_step_index),
				Fpattern_index: uint16(pattern_index),
				Fis_rooted:     is_rooted,
			})
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) == libc.Int32FromUint16(WILDCARD_SYMBOL) {
				(*TSQuery)(unsafe.Pointer(self)).Fwildcard_root_pattern_count = (*TSQuery)(unsafe.Pointer(self)).Fwildcard_root_pattern_count + 1
			}
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Falternative_index) != libc.Int32FromUint16(NONE) {
				start_step_index = uint32((*QueryStep)(unsafe.Pointer(step)).Falternative_index)
			} else {
				if libc.Int32FromUint16(wildcard_root_alternative_index) != libc.Int32FromUint16(NONE) {
					start_step_index = uint32(wildcard_root_alternative_index)
					wildcard_root_alternative_index = NONE
				} else {
					break
				}
			}
			goto _9
		_9:
		}
		pat_start = (*QueryPattern)(unsafe.Pointer(pattern)).Fsteps.Foffset
		pat_end = pat_start + (*QueryPattern)(unsafe.Pointer(pattern)).Fsteps.Flength - uint32(1)
		i = pat_start
		for {
			if !(i < pat_end) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5307, __ccgo_ts+3070, int32(2975), uintptr(unsafe.Pointer(&__func__46)))
				}
			}
			s = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(i)*20
			if !(int32(*(*uint8)(unsafe.Pointer(s + 18))&0x8>>3) != 0) || libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(s)).Falternative_index) == libc.Int32FromUint16(NONE) || uint32((*QueryStep)(unsafe.Pointer(s)).Falternative_index) >= i {
				goto _11
			}
			target_idx = uint32((*QueryStep)(unsafe.Pointer(s)).Falternative_index)
			_ = libc.Uint64FromInt64(4)
			{
				if !(target_idx < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+7389, __ccgo_ts+3070, int32(2981), uintptr(unsafe.Pointer(&__func__46)))
				}
			}
			target = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(target_idx)*20
			target_alt_index = (*QueryStep)(unsafe.Pointer(target)).Falternative_index
			if libc.Int32FromUint16(target_alt_index) == libc.Int32FromUint16(NONE) || uint32(target_alt_index) <= target_idx || uint32(target_alt_index) >= pat_end {
				goto _11
			}
			copy_idx = (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize
			copy1 = *(*QueryStep)(unsafe.Pointer(target))
			copy1.Falternative_index = NONE
			target_depth = (*QueryStep)(unsafe.Pointer(target)).Fdepth
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
			v2 = self + 80 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*QueryStep)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(v1)*20)) = copy1
			*(*QueryStep)(unsafe.Pointer(bp + 48)) = query_step__new(tls, uint16(0), target_depth, libc.BoolUint8(0 != 0))
			libc.SetBitFieldPtr8Uint8(bp+48+18, libc.BoolUint8(1 != 0), 4, 0x10)
			(*(*QueryStep)(unsafe.Pointer(bp + 48))).Falternative_index = uint16(target_idx + uint32(1))
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 80)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize, self+80+12, uint32(1), libc.Uint64FromInt64(20))
			v2 = self + 80 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*QueryStep)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(v1)*20)) = *(*QueryStep)(unsafe.Pointer(bp + 48))
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5307, __ccgo_ts+3070, int32(3004), uintptr(unsafe.Pointer(&__func__46)))
				}
			}
			s = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(i)*20
			(*QueryStep)(unsafe.Pointer(s)).Falternative_index = uint16(copy_idx)
			goto _11
		_11:
			;
			i = i + 1
		}
	}
	if !(ts_query__analyze_patterns(tls, self, error_offset) != 0) {
		*(*TSQueryError1)(unsafe.Pointer(error_type)) = int32(TSQueryErrorStructure)
		ts_query_delete(tls, self)
		return libc.UintptrFromInt32(0)
	}
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+176)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+176)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 176)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 176)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 176)).Fcapacity = uint32(0)
	return self
}

var __func__46 = [13]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'n', 'e', 'w'}

func ts_query_delete(tls *libc.TLS, self uintptr) {
	var capture_quantifiers uintptr
	var index uint32_t
	_, _ = capture_quantifiers, index
	if self != 0 {
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 80)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 80)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 80)).Fcapacity = uint32(0)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 96)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 96)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 96)).Fcapacity = uint32(0)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+112)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+112)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 112)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 112)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 112)).Fcapacity = uint32(0)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+128)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 128)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 128)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 128)).Fcapacity = uint32(0)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+144)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 144)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 144)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 144)).Fcapacity = uint32(0)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+176)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+176)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 176)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 176)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 176)).Fcapacity = uint32(0)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+160)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+160)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 160)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 160)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 160)).Fcapacity = uint32(0)
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+192)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+192)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 192)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 192)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 192)).Fcapacity = uint32(0)
		ts_language_delete(tls, (*TSQuery)(unsafe.Pointer(self)).Flanguage)
		symbol_table_delete(tls, self)
		symbol_table_delete(tls, self+32)
		index = uint32(0)
		for {
			if !(index < (*TSQuery)(unsafe.Pointer(self)).Fcapture_quantifiers.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(index < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+64)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+7435, __ccgo_ts+3070, int32(3034), uintptr(unsafe.Pointer(&__func__47)))
				}
			}
			capture_quantifiers = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+64)).Fcontents + uintptr(index)*16
			capture_quantifiers_delete(tls, capture_quantifiers)
			goto _1
		_1:
			;
			index = index + 1
		}
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+64)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+64)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 64)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 64)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 64)).Fcapacity = uint32(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, self)
	}
}

var __func__47 = [16]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'd', 'e', 'l', 'e', 't', 'e'}

func ts_query_pattern_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSQuery)(unsafe.Pointer(self)).Fpatterns.Fsize
}

func ts_query_capture_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSQuery)(unsafe.Pointer(self)).Fcaptures.Fslices.Fsize
}

func ts_query_string_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSQuery)(unsafe.Pointer(self)).Fpredicate_values.Fslices.Fsize
}

func ts_query_capture_name_for_id(tls *libc.TLS, self uintptr, index uint32_t, length uintptr) (r uintptr) {
	return symbol_table_name_for_id(tls, self, uint16(index), length)
}

func ts_query_capture_quantifier_for_id(tls *libc.TLS, self uintptr, pattern_index uint32_t, capture_index uint32_t) (r TSQuantifier1) {
	var capture_quantifiers uintptr
	_ = capture_quantifiers
	_ = libc.Uint64FromInt64(4)
	{
		if !(pattern_index < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+64)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+7490, __ccgo_ts+3070, int32(3067), uintptr(unsafe.Pointer(&__func__48)))
		}
	}
	capture_quantifiers = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+64)).Fcontents + uintptr(pattern_index)*16
	return capture_quantifier_for_id(tls, capture_quantifiers, uint16(capture_index))
}

var __func__48 = [35]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'a', 'p', 't', 'u', 'r', 'e', '_', 'q', 'u', 'a', 'n', 't', 'i', 'f', 'i', 'e', 'r', '_', 'f', 'o', 'r', '_', 'i', 'd'}

func ts_query_string_value_for_id(tls *libc.TLS, self uintptr, index uint32_t, length uintptr) (r uintptr) {
	return symbol_table_name_for_id(tls, self+32, uint16(index), length)
}

func ts_query_predicates_for_pattern(tls *libc.TLS, self uintptr, pattern_index uint32_t, step_count uintptr) (r uintptr) {
	var slice Slice
	_ = slice
	_ = libc.Uint64FromInt64(4)
	{
		if !(pattern_index < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+7553, __ccgo_ts+3070, int32(3084), uintptr(unsafe.Pointer(&__func__49)))
		}
	}
	slice = (*QueryPattern)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+128)).Fcontents + uintptr(pattern_index)*28)).Fpredicate_steps
	*(*uint32_t)(unsafe.Pointer(step_count)) = slice.Flength
	if slice.Flength == uint32(0) {
		return libc.UintptrFromInt32(0)
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(slice.Foffset < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+112)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+7605, __ccgo_ts+3070, int32(3087), uintptr(unsafe.Pointer(&__func__49)))
		}
	}
	return (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+112)).Fcontents + uintptr(slice.Foffset)*8
}

var __func__49 = [32]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'p', 'r', 'e', 'd', 'i', 'c', 'a', 't', 'e', 's', '_', 'f', 'o', 'r', '_', 'p', 'a', 't', 't', 'e', 'r', 'n'}

func ts_query_start_byte_for_pattern(tls *libc.TLS, self uintptr, pattern_index uint32_t) (r uint32_t) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(pattern_index < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+7553, __ccgo_ts+3070, int32(3094), uintptr(unsafe.Pointer(&__func__50)))
		}
	}
	return (*QueryPattern)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+128)).Fcontents + uintptr(pattern_index)*28)).Fstart_byte
}

var __func__50 = [32]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 's', 't', 'a', 'r', 't', '_', 'b', 'y', 't', 'e', '_', 'f', 'o', 'r', '_', 'p', 'a', 't', 't', 'e', 'r', 'n'}

func ts_query_end_byte_for_pattern(tls *libc.TLS, self uintptr, pattern_index uint32_t) (r uint32_t) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(pattern_index < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+7553, __ccgo_ts+3070, int32(3101), uintptr(unsafe.Pointer(&__func__51)))
		}
	}
	return (*QueryPattern)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+128)).Fcontents + uintptr(pattern_index)*28)).Fend_byte
}

var __func__51 = [30]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'e', 'n', 'd', '_', 'b', 'y', 't', 'e', '_', 'f', 'o', 'r', '_', 'p', 'a', 't', 't', 'e', 'r', 'n'}

func ts_query_is_pattern_rooted(tls *libc.TLS, self uintptr, pattern_index uint32_t) (r uint8) {
	var entry uintptr
	var i uint32
	_, _ = entry, i
	i = uint32(0)
	for {
		if !(i < (*TSQuery)(unsafe.Pointer(self)).Fpattern_map.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5209, __ccgo_ts+3070, int32(3109), uintptr(unsafe.Pointer(&__func__52)))
			}
		}
		entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents + uintptr(i)*6
		if uint32((*PatternEntry)(unsafe.Pointer(entry)).Fpattern_index) == pattern_index {
			if !((*PatternEntry)(unsafe.Pointer(entry)).Fis_rooted != 0) {
				return libc.BoolUint8(0 != 0)
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return libc.BoolUint8(1 != 0)
}

var __func__52 = [27]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'i', 's', '_', 'p', 'a', 't', 't', 'e', 'r', 'n', '_', 'r', 'o', 'o', 't', 'e', 'd'}

func ts_query_is_pattern_non_local(tls *libc.TLS, self uintptr, pattern_index uint32_t) (r uint8) {
	if pattern_index < (*TSQuery)(unsafe.Pointer(self)).Fpatterns.Fsize {
		_ = libc.Uint64FromInt64(4)
		{
			if !(pattern_index < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+128)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7553, __ccgo_ts+3070, int32(3122), uintptr(unsafe.Pointer(&__func__53)))
			}
		}
		return (*QueryPattern)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+128)).Fcontents + uintptr(pattern_index)*28)).Fis_non_local
	} else {
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var __func__53 = [30]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'i', 's', '_', 'p', 'a', 't', 't', 'e', 'r', 'n', '_', 'n', 'o', 'n', '_', 'l', 'o', 'c', 'a', 'l'}

func ts_query_is_pattern_guaranteed_at_step(tls *libc.TLS, self uintptr, byte_offset uint32_t) (r uint8) {
	var i uint32
	var step_index uint32_t
	var step_offset uintptr
	_, _, _ = i, step_index, step_offset
	step_index = libc.Uint32FromUint32(4294967295)
	i = uint32(0)
	for {
		if !(i < (*TSQuery)(unsafe.Pointer(self)).Fstep_offsets.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+144)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7663, __ccgo_ts+3070, int32(3134), uintptr(unsafe.Pointer(&__func__54)))
			}
		}
		step_offset = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+144)).Fcontents + uintptr(i)*8
		if (*StepOffset)(unsafe.Pointer(step_offset)).Fbyte_offset > byte_offset {
			break
		}
		step_index = uint32((*StepOffset)(unsafe.Pointer(step_offset)).Fstep_index)
		goto _1
	_1:
		;
		i = i + 1
	}
	if step_index < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize {
		_ = libc.Uint64FromInt64(4)
		{
			if !(step_index < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(3139), uintptr(unsafe.Pointer(&__func__54)))
			}
		}
		return libc.Uint8FromInt32(libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index)*20 + 18))&0x80>>7) != 0))
	} else {
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var __func__54 = [39]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'i', 's', '_', 'p', 'a', 't', 't', 'e', 'r', 'n', '_', 'g', 'u', 'a', 'r', 'a', 'n', 't', 'e', 'e', 'd', '_', 'a', 't', '_', 's', 't', 'e', 'p'}

func ts_query__step_is_fallible(tls *libc.TLS, self uintptr, step_index uint16_t) (r uint8) {
	var next_step, step uintptr
	_, _ = next_step, step
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(step_index)+libc.Uint32FromInt32(1) < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+7707, __ccgo_ts+3070, int32(3149), uintptr(unsafe.Pointer(&__func__55)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32(step_index) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+6579, __ccgo_ts+3070, int32(3150), uintptr(unsafe.Pointer(&__func__55)))
		}
	}
	step = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+80)).Fcontents + uintptr(step_index)*20
	_ = libc.Uint64FromInt64(4)
	{
		if !(libc.Uint32FromInt32(libc.Int32FromUint16(step_index)+libc.Int32FromInt32(1)) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+80)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+7751, __ccgo_ts+3070, int32(3151), uintptr(unsafe.Pointer(&__func__55)))
		}
	}
	next_step = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+80)).Fcontents + uintptr(libc.Int32FromUint16(step_index)+int32(1))*20
	return libc.BoolUint8(libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth) != libc.Int32FromUint16(PATTERN_DONE_MARKER) && libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth) > libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth) && (!(int32(*(*uint8)(unsafe.Pointer(next_step + 19))&0x1>>0) != 0) || libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) == libc.Int32FromUint16(WILDCARD_SYMBOL)))
}

var __func__55 = [27]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', '_', 's', 't', 'e', 'p', '_', 'i', 's', '_', 'f', 'a', 'l', 'l', 'i', 'b', 'l', 'e'}

func ts_query_disable_capture(tls *libc.TLS, self uintptr, name uintptr, length uint32_t) {
	var i uint32
	var id int32
	var step uintptr
	_, _, _ = i, id, step
	id = symbol_table_id_for_name(tls, self, name, length)
	if id != -int32(1) {
		i = uint32(0)
		for {
			if !(i < (*TSQuery)(unsafe.Pointer(self)).Fsteps.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+5307, __ccgo_ts+3070, int32(3169), uintptr(unsafe.Pointer(&__func__56)))
				}
			}
			step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+80)).Fcontents + uintptr(i)*20
			query_step__remove_capture(tls, step, libc.Uint16FromInt32(id))
			goto _1
		_1:
			;
			i = i + 1
		}
	}
}

var __func__56 = [25]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'd', 'i', 's', 'a', 'b', 'l', 'e', '_', 'c', 'a', 'p', 't', 'u', 'r', 'e'}

func ts_query_disable_pattern(tls *libc.TLS, self uintptr, pattern_index uint32_t) {
	var i uint32
	var pattern uintptr
	_, _ = i, pattern
	i = uint32(0)
	for {
		if !(i < (*TSQuery)(unsafe.Pointer(self)).Fpattern_map.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+5209, __ccgo_ts+3070, int32(3182), uintptr(unsafe.Pointer(&__func__57)))
			}
		}
		pattern = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+96)).Fcontents + uintptr(i)*6
		if uint32((*PatternEntry)(unsafe.Pointer(pattern)).Fpattern_index) == pattern_index {
			_array__erase(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+96)).Fcontents, self+96+8, libc.Uint64FromInt64(6), i)
			i = i - 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

var __func__57 = [25]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'd', 'i', 's', 'a', 'b', 'l', 'e', '_', 'p', 'a', 't', 't', 'e', 'r', 'n'}

func ts_query_cursor_new(tls *libc.TLS) (r uintptr) {
	var self uintptr
	_ = self
	self = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(208))
	*(*TSQueryCursor)(unsafe.Pointer(self)) = TSQueryCursor{
		Fcapture_list_pool: capture_list_pool_new(tls),
		Fmax_start_depth:   libc.Uint32FromUint32(4294967295),
		Fincluded_range: TSRange{
			Fend_point: TSPoint{
				Frow:    libc.Uint32FromUint32(4294967295),
				Fcolumn: libc.Uint32FromUint32(4294967295),
			},
			Fend_byte: libc.Uint32FromUint32(4294967295),
		},
		Fcontaining_range: TSRange{
			Fend_point: TSPoint{
				Frow:    libc.Uint32FromUint32(4294967295),
				Fcolumn: libc.Uint32FromUint32(4294967295),
			},
			Fend_byte: libc.Uint32FromUint32(4294967295),
		},
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 40)).Fcontents = _array__reserve(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+40)).Fcontents, self+40+12, libc.Uint64FromInt64(16), uint32(8))
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 56)).Fcontents = _array__reserve(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+56)).Fcontents, self+56+12, libc.Uint64FromInt64(16), uint32(8))
	return self
}

func ts_query_cursor_delete(tls *libc.TLS, self uintptr) {
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+40)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+40)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 40)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 40)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 40)).Fcapacity = uint32(0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+56)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+56)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 56)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 56)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 56)).Fcapacity = uint32(0)
	ts_tree_cursor_delete(tls, self+8)
	capture_list_pool_delete(tls, self+72)
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, self)
}

func ts_query_cursor_did_exceed_match_limit(tls *libc.TLS, self uintptr) (r uint8) {
	return (*TSQueryCursor)(unsafe.Pointer(self)).Fdid_exceed_match_limit
}

func ts_query_cursor_match_limit(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*TSQueryCursor)(unsafe.Pointer(self)).Fcapture_list_pool.Fmax_capture_list_count
}

func ts_query_cursor_set_match_limit(tls *libc.TLS, self uintptr, limit uint32_t) {
	(*TSQueryCursor)(unsafe.Pointer(self)).Fcapture_list_pool.Fmax_capture_list_count = limit
}

func ts_query_cursor_exec(tls *libc.TLS, self uintptr, query uintptr, node TSNode) {
	var i uint32
	var step uintptr
	_, _ = i, step
	if query != 0 {
		i = uint32(0)
		for {
			if !(i < (*TSQuery)(unsafe.Pointer(query)).Fsteps.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(query+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+7801, __ccgo_ts+3070, int32(3257), uintptr(unsafe.Pointer(&__func__58)))
				}
			}
			step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(query+80)).Fcontents + uintptr(i)*20
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) {
			} else {
				if int32(*(*uint8)(unsafe.Pointer(step + 18))&0x10>>4) != 0 {
				} else {
					if int32(*(*uint8)(unsafe.Pointer(step + 18))&0x8>>3) != 0 {
					} else {
						if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fsymbol) != libc.Int32FromUint16(WILDCARD_SYMBOL) {
						} else {
						}
					}
				}
			}
			if (*QueryStep)(unsafe.Pointer(step)).Ffield != 0 {
			}
			if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Falternative_index) != libc.Int32FromUint16(NONE) {
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 40)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 56)).Fsize = uint32(0)
	ts_tree_cursor_reset(tls, self+8, node)
	capture_list_pool_reset(tls, self+72)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node = libc.BoolUint8(1 != 0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fnext_state_id = uint32(0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fdepth = uint32(0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fascending = libc.BoolUint8(0 != 0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fhalted = libc.BoolUint8(0 != 0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fquery = query
	(*TSQueryCursor)(unsafe.Pointer(self)).Fdid_exceed_match_limit = libc.BoolUint8(0 != 0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Foperation_count = uint32(0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fquery_options = libc.UintptrFromInt32(0)
	(*TSQueryCursor)(unsafe.Pointer(self)).Fquery_state = TSQueryCursorState{}
}

var __func__58 = [21]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'e', 'x', 'e', 'c'}

func ts_query_cursor_exec_with_options(tls *libc.TLS, self uintptr, query uintptr, node TSNode, query_options uintptr) {
	ts_query_cursor_exec(tls, self, query, node)
	if query_options != 0 {
		(*TSQueryCursor)(unsafe.Pointer(self)).Fquery_options = query_options
		(*TSQueryCursor)(unsafe.Pointer(self)).Fquery_state = TSQueryCursorState{
			Fpayload: (*TSQueryCursorOptions)(unsafe.Pointer(query_options)).Fpayload,
		}
	}
}

func ts_query_cursor_set_byte_range(tls *libc.TLS, self uintptr, start_byte uint32_t, end_byte uint32_t) (r uint8) {
	if end_byte == uint32(0) {
		end_byte = libc.Uint32FromUint32(4294967295)
	}
	if start_byte > end_byte {
		return libc.BoolUint8(0 != 0)
	}
	(*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fstart_byte = start_byte
	(*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fend_byte = end_byte
	return libc.BoolUint8(1 != 0)
}

func ts_query_cursor_set_point_range(tls *libc.TLS, self uintptr, start_point TSPoint, end_point TSPoint) (r uint8) {
	if end_point.Frow == uint32(0) && end_point.Fcolumn == uint32(0) {
		end_point = TSPoint{
			Frow:    libc.Uint32FromUint32(4294967295),
			Fcolumn: libc.Uint32FromUint32(4294967295),
		}
	}
	if point_gt(tls, start_point, end_point) != 0 {
		return libc.BoolUint8(0 != 0)
	}
	(*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fstart_point = start_point
	(*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fend_point = end_point
	return libc.BoolUint8(1 != 0)
}

func ts_query_cursor_set_containing_byte_range(tls *libc.TLS, self uintptr, start_byte uint32_t, end_byte uint32_t) (r uint8) {
	if end_byte == uint32(0) {
		end_byte = libc.Uint32FromUint32(4294967295)
	}
	if start_byte > end_byte {
		return libc.BoolUint8(0 != 0)
	}
	(*TSQueryCursor)(unsafe.Pointer(self)).Fcontaining_range.Fstart_byte = start_byte
	(*TSQueryCursor)(unsafe.Pointer(self)).Fcontaining_range.Fend_byte = end_byte
	return libc.BoolUint8(1 != 0)
}

func ts_query_cursor_set_containing_point_range(tls *libc.TLS, self uintptr, start_point TSPoint, end_point TSPoint) (r uint8) {
	if end_point.Frow == uint32(0) && end_point.Fcolumn == uint32(0) {
		end_point = TSPoint{
			Frow:    libc.Uint32FromUint32(4294967295),
			Fcolumn: libc.Uint32FromUint32(4294967295),
		}
	}
	if point_gt(tls, start_point, end_point) != 0 {
		return libc.BoolUint8(0 != 0)
	}
	(*TSQueryCursor)(unsafe.Pointer(self)).Fcontaining_range.Fstart_point = start_point
	(*TSQueryCursor)(unsafe.Pointer(self)).Fcontaining_range.Fend_point = end_point
	return libc.BoolUint8(1 != 0)
}

func ts_query_cursor__first_in_progress_capture(tls *libc.TLS, self uintptr, state_index uintptr, byte_offset uintptr, pattern_index uintptr, is_definite uintptr) (r uint8) {
	var captures, state, step uintptr
	var i uint32
	var node TSNode
	var node_start_byte uint32_t
	var result uint8
	_, _, _, _, _, _, _ = captures, i, node, node_start_byte, result, state, step
	result = libc.BoolUint8(0 != 0)
	*(*uint32_t)(unsafe.Pointer(state_index)) = libc.Uint32FromUint32(4294967295)
	*(*uint32_t)(unsafe.Pointer(byte_offset)) = libc.Uint32FromUint32(4294967295)
	*(*uint32_t)(unsafe.Pointer(pattern_index)) = libc.Uint32FromUint32(4294967295)
	i = uint32(0)
	for {
		if !(i < (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+40)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7839, __ccgo_ts+3070, int32(3389), uintptr(unsafe.Pointer(&__func__59)))
			}
		}
		state = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+40)).Fcontents + uintptr(i)*16
		if int32(uint8(*(*uint16)(unsafe.Pointer(state + 14))&0x4000>>14)) != 0 {
			goto _1
		}
		captures = capture_list_pool_get(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
		if libc.Uint32FromInt32(int32(*(*uint16)(unsafe.Pointer(state + 14))&0xfff>>0)) >= (*CaptureList)(unsafe.Pointer(captures)).Fsize {
			goto _1
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(int32(*(*uint16)(unsafe.Pointer(state + 14))&0xfff>>0)) < (*CaptureList)(unsafe.Pointer(captures)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7877, __ccgo_ts+3070, int32(3400), uintptr(unsafe.Pointer(&__func__59)))
			}
		}
		node = (*TSQueryCapture)(unsafe.Pointer((*CaptureList)(unsafe.Pointer(captures)).Fcontents + uintptr(int32(*(*uint16)(unsafe.Pointer(state + 14))&0xfff>>0))*40)).Fnode
		if ts_node_end_byte(tls, node) <= (*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fstart_byte || point_lte(tls, ts_node_end_point(tls, node), (*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fstart_point) != 0 {
			libc.PostIncBitFieldPtr16Uint16(state+14, 1, 12, 0, 0xfff)
			i = i - 1
			goto _1
		}
		node_start_byte = ts_node_start_byte(tls, node)
		if !(result != 0) || node_start_byte < *(*uint32_t)(unsafe.Pointer(byte_offset)) || node_start_byte == *(*uint32_t)(unsafe.Pointer(byte_offset)) && uint32((*QueryState)(unsafe.Pointer(state)).Fpattern_index) < *(*uint32_t)(unsafe.Pointer(pattern_index)) {
			_ = libc.Uint64FromInt64(4)
			{
				if !(uint32((*QueryState)(unsafe.Pointer(state)).Fstep_index) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+7938, __ccgo_ts+3070, int32(3416), uintptr(unsafe.Pointer(&__func__59)))
				}
			}
			step = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*QueryState)(unsafe.Pointer(state)).Fstep_index)*20
			if is_definite != 0 {
				*(*uint8)(unsafe.Pointer(is_definite)) = libc.BoolUint8(int32(*(*uint8)(unsafe.Pointer(step + 18))&0x80>>7) != 0 && !(int32(*(*uint8)(unsafe.Pointer(step + 18))&0x2>>1) != 0))
			} else {
				if int32(*(*uint8)(unsafe.Pointer(step + 18))&0x80>>7) != 0 {
					goto _1
				}
			}
			result = libc.BoolUint8(1 != 0)
			*(*uint32_t)(unsafe.Pointer(state_index)) = i
			*(*uint32_t)(unsafe.Pointer(byte_offset)) = node_start_byte
			*(*uint32_t)(unsafe.Pointer(pattern_index)) = uint32((*QueryState)(unsafe.Pointer(state)).Fpattern_index)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return result
}

var __func__59 = [43]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', '_', 'f', 'i', 'r', 's', 't', '_', 'i', 'n', '_', 'p', 'r', 'o', 'g', 'r', 'e', 's', 's', '_', 'c', 'a', 'p', 't', 'u', 'r', 'e'}

func ts_query_cursor__compare_nodes(tls *libc.TLS, left TSNode, right TSNode) (r int32) {
	var left_node_count, left_start, right_node_count, right_start uint32_t
	_, _, _, _ = left_node_count, left_start, right_node_count, right_start
	if left.Fid != right.Fid {
		left_start = ts_node_start_byte(tls, left)
		right_start = ts_node_start_byte(tls, right)
		if left_start < right_start {
			return -int32(1)
		}
		if left_start > right_start {
			return int32(1)
		}
		left_node_count = ts_node_end_byte(tls, left)
		right_node_count = ts_node_end_byte(tls, right)
		if left_node_count > right_node_count {
			return -int32(1)
		}
		if left_node_count < right_node_count {
			return int32(1)
		}
	}
	return 0
}

func ts_query_cursor__compare_captures(tls *libc.TLS, self uintptr, left_state uintptr, right_state uintptr, left_contains_right uintptr, right_contains_left uintptr) {
	var i, j uint32
	var left, left_captures, right, right_captures uintptr
	_, _, _, _, _, _ = i, j, left, left_captures, right, right_captures
	left_captures = capture_list_pool_get(tls, self+72, uint16((*QueryState)(unsafe.Pointer(left_state)).Fcapture_list_id))
	right_captures = capture_list_pool_get(tls, self+72, uint16((*QueryState)(unsafe.Pointer(right_state)).Fcapture_list_id))
	*(*uint8)(unsafe.Pointer(left_contains_right)) = libc.BoolUint8(1 != 0)
	*(*uint8)(unsafe.Pointer(right_contains_left)) = libc.BoolUint8(1 != 0)
	i = uint32(0)
	j = uint32(0)
	for {
		if i < (*CaptureList)(unsafe.Pointer(left_captures)).Fsize {
			if j < (*CaptureList)(unsafe.Pointer(right_captures)).Fsize {
				_ = libc.Uint64FromInt64(4)
				{
					if !(i < (*CaptureList)(unsafe.Pointer(left_captures)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+7998, __ccgo_ts+3070, int32(3472), uintptr(unsafe.Pointer(&__func__60)))
					}
				}
				left = (*CaptureList)(unsafe.Pointer(left_captures)).Fcontents + uintptr(i)*40
				_ = libc.Uint64FromInt64(4)
				{
					if !(j < (*CaptureList)(unsafe.Pointer(right_captures)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+8036, __ccgo_ts+3070, int32(3473), uintptr(unsafe.Pointer(&__func__60)))
					}
				}
				right = (*CaptureList)(unsafe.Pointer(right_captures)).Fcontents + uintptr(j)*40
				if (*TSQueryCapture)(unsafe.Pointer(left)).Fnode.Fid == (*TSQueryCapture)(unsafe.Pointer(right)).Fnode.Fid && (*TSQueryCapture)(unsafe.Pointer(left)).Findex == (*TSQueryCapture)(unsafe.Pointer(right)).Findex {
					i = i + 1
					j = j + 1
				} else {
					switch ts_query_cursor__compare_nodes(tls, (*TSQueryCapture)(unsafe.Pointer(left)).Fnode, (*TSQueryCapture)(unsafe.Pointer(right)).Fnode) {
					case -int32(1):
						*(*uint8)(unsafe.Pointer(right_contains_left)) = libc.BoolUint8(0 != 0)
						i = i + 1
					case int32(1):
						*(*uint8)(unsafe.Pointer(left_contains_right)) = libc.BoolUint8(0 != 0)
						j = j + 1
					default:
						*(*uint8)(unsafe.Pointer(right_contains_left)) = libc.BoolUint8(0 != 0)
						*(*uint8)(unsafe.Pointer(left_contains_right)) = libc.BoolUint8(0 != 0)
						i = i + 1
						j = j + 1
						break
					}
				}
			} else {
				*(*uint8)(unsafe.Pointer(right_contains_left)) = libc.BoolUint8(0 != 0)
				break
			}
		} else {
			if j < (*CaptureList)(unsafe.Pointer(right_captures)).Fsize {
				*(*uint8)(unsafe.Pointer(left_contains_right)) = libc.BoolUint8(0 != 0)
			}
			break
		}
		goto _1
	_1:
	}
}

var __func__60 = [34]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', '_', 'c', 'o', 'm', 'p', 'a', 'r', 'e', '_', 'c', 'a', 'p', 't', 'u', 'r', 'e', 's'}

func ts_query_cursor__add_state(tls *libc.TLS, self uintptr, pattern uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var index, start_depth uint32_t
	var prev_state, step uintptr
	_, _, _, _ = index, prev_state, start_depth, step
	_ = libc.Uint64FromInt64(4)
	{
		if !(uint32((*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8075, __ccgo_ts+3070, int32(3512), uintptr(unsafe.Pointer(&__func__61)))
		}
	}
	step = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index)*20
	start_depth = (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth - uint32((*QueryStep)(unsafe.Pointer(step)).Fdepth)
	index = (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize
	for index > uint32(0) {
		_ = libc.Uint64FromInt64(4)
		{
			if !(index-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+40)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+8137, __ccgo_ts+3070, int32(3536), uintptr(unsafe.Pointer(&__func__61)))
			}
		}
		prev_state = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+40)).Fcontents + uintptr(index-uint32(1))*16
		if uint32((*QueryState)(unsafe.Pointer(prev_state)).Fstart_depth) < start_depth {
			break
		}
		if uint32((*QueryState)(unsafe.Pointer(prev_state)).Fstart_depth) == start_depth {
			if libc.Int32FromUint16((*QueryState)(unsafe.Pointer(prev_state)).Fpattern_index) == libc.Int32FromUint16((*PatternEntry)(unsafe.Pointer(pattern)).Fpattern_index) && libc.Int32FromUint16((*QueryState)(unsafe.Pointer(prev_state)).Fstep_index) == libc.Int32FromUint16((*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index) {
				return
			}
			if libc.Int32FromUint16((*QueryState)(unsafe.Pointer(prev_state)).Fpattern_index) <= libc.Int32FromUint16((*PatternEntry)(unsafe.Pointer(pattern)).Fpattern_index) {
				break
			}
		}
		index = index - 1
	}
	*(*QueryState)(unsafe.Pointer(bp)) = QueryState{
		Fid:              libc.Uint32FromUint32(4294967295),
		Fcapture_list_id: uint32(NONE),
		Fstart_depth:     uint16(start_depth),
		Fstep_index:      (*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index,
		Fpattern_index:   (*PatternEntry)(unsafe.Pointer(pattern)).Fpattern_index,
		F__ccgo14:        libc.Uint16FromInt32(0)&0xfff<<0 | libc.Uint16FromInt32(1)&0x1<<12 | libc.Uint16FromInt32(0)&0x1<<13 | libc.Uint16FromInt32(0)&0x1<<14 | libc.BoolUint16(libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth) == libc.Int32FromInt32(1))&0x1<<15,
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 40)).Fcontents = _array__splice(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+40)).Fcontents, self+40+8, self+40+12, libc.Uint64FromInt64(16), index, uint32(0), uint32(1), bp)
}

var __func__61 = [27]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', '_', 'a', 'd', 'd', '_', 's', 't', 'a', 't', 'e'}

func ts_query_cursor__prepare_to_capture(tls *libc.TLS, self uintptr, state uintptr, state_index_to_preserve uint32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var list, other_state uintptr
	var _ /* byte_offset at bp+4 */ uint32_t
	var _ /* pattern_index at bp+8 */ uint32_t
	var _ /* state_index at bp+0 */ uint32_t
	_, _ = list, other_state
	if (*QueryState)(unsafe.Pointer(state)).Fcapture_list_id == uint32(NONE) {
		(*QueryState)(unsafe.Pointer(state)).Fcapture_list_id = uint32(capture_list_pool_acquire(tls, self+72))
		if (*QueryState)(unsafe.Pointer(state)).Fcapture_list_id == uint32(NONE) {
			(*TSQueryCursor)(unsafe.Pointer(self)).Fdid_exceed_match_limit = libc.BoolUint8(1 != 0)
			if ts_query_cursor__first_in_progress_capture(tls, self, bp, bp+4, bp+8, libc.UintptrFromInt32(0)) != 0 && *(*uint32_t)(unsafe.Pointer(bp)) != state_index_to_preserve {
				_ = libc.Uint64FromInt64(4)
				{
					if !(*(*uint32_t)(unsafe.Pointer(bp)) < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+40)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+8183, __ccgo_ts+3070, int32(3600), uintptr(unsafe.Pointer(&__func__62)))
					}
				}
				other_state = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+40)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp)))*16
				(*QueryState)(unsafe.Pointer(state)).Fcapture_list_id = (*QueryState)(unsafe.Pointer(other_state)).Fcapture_list_id
				(*QueryState)(unsafe.Pointer(other_state)).Fcapture_list_id = uint32(NONE)
				libc.SetBitFieldPtr16Uint8(other_state+14, libc.BoolUint8(1 != 0), 14, 0x4000)
				list = capture_list_pool_get_mut(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
				(*CaptureList)(unsafe.Pointer(list)).Fsize = uint32(0)
				return list
			} else {
				return libc.UintptrFromInt32(0)
			}
		}
	}
	return capture_list_pool_get_mut(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
}

var __func__62 = [36]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', '_', 'p', 'r', 'e', 'p', 'a', 'r', 'e', '_', 't', 'o', '_', 'c', 'a', 'p', 't', 'u', 'r', 'e'}

func ts_query_cursor__capture(tls *libc.TLS, self uintptr, state uintptr, step uintptr, node TSNode) {
	var capture_id uint16_t
	var capture_list, v3 uintptr
	var j uint32
	var v2 uint32_t
	_, _, _, _, _ = capture_id, capture_list, j, v2, v3
	if int32(uint8(*(*uint16)(unsafe.Pointer(state + 14))&0x4000>>14)) != 0 {
		return
	}
	capture_list = ts_query_cursor__prepare_to_capture(tls, self, state, uint32(4294967295))
	if !(capture_list != 0) {
		libc.SetBitFieldPtr16Uint8(state+14, libc.BoolUint8(1 != 0), 14, 0x4000)
		return
	}
	j = uint32(0)
	for {
		if !(j < uint32(3)) {
			break
		}
		capture_id = *(*uint16_t)(unsafe.Pointer(step + 6 + uintptr(j)*2))
		if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(step + 6 + uintptr(j)*2))) == libc.Int32FromUint16(NONE) {
			break
		}
		(*CaptureList)(unsafe.Pointer(capture_list)).Fcontents = _array__grow(tls, (*CaptureList)(unsafe.Pointer(capture_list)).Fcontents, (*CaptureList)(unsafe.Pointer(capture_list)).Fsize, capture_list+12, uint32(1), libc.Uint64FromInt64(40))
		v3 = capture_list + 8
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
		*(*TSQueryCapture)(unsafe.Pointer((*CaptureList)(unsafe.Pointer(capture_list)).Fcontents + uintptr(v2)*40)) = TSQueryCapture{
			Fnode:  node,
			Findex: uint32(capture_id),
		}
		goto _1
	_1:
		;
		j = j + 1
	}
}

func ts_query_cursor__copy_state(tls *libc.TLS, self uintptr, state_ref uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var new_captures, old_captures, state uintptr
	var state_index uint32_t
	var _ /* copy at bp+0 */ QueryState
	_, _, _, _ = new_captures, old_captures, state, state_index
	state = *(*uintptr)(unsafe.Pointer(state_ref))
	state_index = libc.Uint32FromInt64((int64(state) - int64((*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fcontents)) / 16)
	*(*QueryState)(unsafe.Pointer(bp)) = *(*QueryState)(unsafe.Pointer(state))
	(*(*QueryState)(unsafe.Pointer(bp))).Fcapture_list_id = uint32(NONE)
	if (*QueryState)(unsafe.Pointer(state)).Fcapture_list_id != uint32(NONE) {
		new_captures = ts_query_cursor__prepare_to_capture(tls, self, bp, state_index)
		if !(new_captures != 0) {
			return libc.UintptrFromInt32(0)
		}
		old_captures = capture_list_pool_get(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
		(*CaptureList)(unsafe.Pointer(new_captures)).Fcontents = _array__splice(tls, (*CaptureList)(unsafe.Pointer(new_captures)).Fcontents, new_captures+8, new_captures+12, libc.Uint64FromInt64(40), (*CaptureList)(unsafe.Pointer(new_captures)).Fsize, uint32(0), (*CaptureList)(unsafe.Pointer(old_captures)).Fsize, (*CaptureList)(unsafe.Pointer(old_captures)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 40)).Fcontents = _array__splice(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+40)).Fcontents, self+40+8, self+40+12, libc.Uint64FromInt64(16), state_index+uint32(1), uint32(0), uint32(1), bp)
	_ = libc.Uint64FromInt64(4)
	{
		if !(state_index < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+40)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8183, __ccgo_ts+3070, int32(3670), uintptr(unsafe.Pointer(&__func__63)))
		}
	}
	*(*uintptr)(unsafe.Pointer(state_ref)) = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+40)).Fcontents + uintptr(state_index)*16
	_ = libc.Uint64FromInt64(4)
	{
		if !(state_index+libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+40)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8231, __ccgo_ts+3070, int32(3671), uintptr(unsafe.Pointer(&__func__63)))
		}
	}
	return (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+40)).Fcontents + uintptr(state_index+uint32(1))*16
}

var __func__63 = [28]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', '_', 'c', 'o', 'p', 'y', '_', 's', 't', 'a', 't', 'e'}

func ts_query_cursor__should_descend(tls *libc.TLS, self uintptr, node_intersects_range uint8) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var comparison int32
	var half_size, mid_index, size uint32_t
	var i uint32
	var next_step, state uintptr
	var subtree Subtree
	var _ /* exists at bp+0 */ uint8
	var _ /* index at bp+4 */ uint32_t
	_, _, _, _, _, _, _, _ = comparison, half_size, i, mid_index, next_step, size, state, subtree
	if node_intersects_range != 0 && (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth < (*TSQueryCursor)(unsafe.Pointer(self)).Fmax_start_depth {
		return libc.BoolUint8(1 != 0)
	}
	i = uint32(0)
	for {
		if !(i < (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+40)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7839, __ccgo_ts+3070, int32(3686), uintptr(unsafe.Pointer(&__func__64)))
			}
		}
		state = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+40)).Fcontents + uintptr(i)*16
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32((*QueryState)(unsafe.Pointer(state)).Fstep_index) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7938, __ccgo_ts+3070, int32(3687), uintptr(unsafe.Pointer(&__func__64)))
			}
		}
		next_step = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*QueryState)(unsafe.Pointer(state)).Fstep_index)*20
		if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth) != libc.Int32FromUint16(PATTERN_DONE_MARKER) && libc.Uint32FromInt32(libc.Int32FromUint16((*QueryState)(unsafe.Pointer(state)).Fstart_depth)+libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step)).Fdepth)) > (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth {
			return libc.BoolUint8(1 != 0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth >= (*TSQueryCursor)(unsafe.Pointer(self)).Fmax_start_depth {
		return libc.BoolUint8(0 != 0)
	}
	if !((*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node != 0) {
		subtree = ts_tree_cursor_current_subtree(tls, self+8)
		if ts_subtree_is_repetition(tls, subtree) != 0 {
			*(*uint32_t)(unsafe.Pointer(bp + 4)) = uint32(0)
			*(*uint8)(unsafe.Pointer(bp)) = libc.BoolUint8(0 != 0)
			size = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+192)).Fsize - *(*uint32_t)(unsafe.Pointer(bp + 4))
			if size == uint32(0) {
				goto _2
			}
			for size > uint32(1) {
				half_size = size / uint32(2)
				mid_index = *(*uint32_t)(unsafe.Pointer(bp + 4)) + half_size
				comparison = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+192)).Fcontents + uintptr(mid_index)*2))) - libc.Int32FromUint16(ts_subtree_symbol(tls, subtree))
				if comparison <= 0 {
					*(*uint32_t)(unsafe.Pointer(bp + 4)) = mid_index
				}
				size = size - half_size
			}
			comparison = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+192)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 4)))*2))) - libc.Int32FromUint16(ts_subtree_symbol(tls, subtree))
			if comparison == 0 {
				*(*uint8)(unsafe.Pointer(bp)) = libc.BoolUint8(1 != 0)
			} else {
				if comparison < 0 {
					*(*uint32_t)(unsafe.Pointer(bp + 4)) += uint32(1)
				}
			}
			goto _3
		_3:
			;
			goto _2
		_2: /**/
			; //
			return *(*uint8)(unsafe.Pointer(bp))
		}
		return libc.BoolUint8(1 != 0)
	}
	return libc.BoolUint8(0 != 0)
}

var __func__64 = [32]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', '_', 's', 'h', 'o', 'u', 'l', 'd', '_', 'd', 'e', 's', 'c', 'e', 'n', 'd'}

func range_intersects(tls *libc.TLS, a uintptr, b uintptr) (r uint8) {
	var is_empty uint8
	_ = is_empty
	is_empty = libc.BoolUint8((*TSRange)(unsafe.Pointer(a)).Fstart_byte == (*TSRange)(unsafe.Pointer(a)).Fend_byte)
	return libc.BoolUint8(((*TSRange)(unsafe.Pointer(a)).Fend_byte > (*TSRange)(unsafe.Pointer(b)).Fstart_byte || is_empty != 0 && (*TSRange)(unsafe.Pointer(a)).Fend_byte == (*TSRange)(unsafe.Pointer(b)).Fstart_byte) && (point_gt(tls, (*TSRange)(unsafe.Pointer(a)).Fend_point, (*TSRange)(unsafe.Pointer(b)).Fstart_point) != 0 || is_empty != 0 && point_eq(tls, (*TSRange)(unsafe.Pointer(a)).Fend_point, (*TSRange)(unsafe.Pointer(b)).Fstart_point) != 0) && (*TSRange)(unsafe.Pointer(a)).Fstart_byte < (*TSRange)(unsafe.Pointer(b)).Fend_byte && point_lt(tls, (*TSRange)(unsafe.Pointer(a)).Fstart_point, (*TSRange)(unsafe.Pointer(b)).Fend_point) != 0)
}

func range_within(tls *libc.TLS, a uintptr, b uintptr) (r uint8) {
	return libc.BoolUint8((*TSRange)(unsafe.Pointer(a)).Fstart_byte >= (*TSRange)(unsafe.Pointer(b)).Fstart_byte && point_gte(tls, (*TSRange)(unsafe.Pointer(a)).Fstart_point, (*TSRange)(unsafe.Pointer(b)).Fstart_point) != 0 && (*TSRange)(unsafe.Pointer(a)).Fend_byte <= (*TSRange)(unsafe.Pointer(b)).Fend_byte && point_lte(tls, (*TSRange)(unsafe.Pointer(a)).Fend_point, (*TSRange)(unsafe.Pointer(b)).Fend_point) != 0)
}

func ts_query_cursor__advance(tls *libc.TLS, self uintptr, stop_on_definite_step uint8) (r uint8) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var child_step, copy1, negated_field_ids, next_step, next_step1, other_state, pattern, pattern1, skipped_wildcard_step, state1, state3, step, step1, step2, step3, v3 uintptr
	var copy_count, end_index, i, i1, j, j1, k, k1, k2, n, v4 uint32
	var deleted_count, start_depth, start_depth1, v2 uint32_t
	var did_match, did_remove, has_supertype, is_missing, is_named, later_sibling_can_match, node_does_match, node_intersects_containing_range, node_intersects_range, node_is_error, node_within_containing_range, parent_intersects_range, parent_is_error uint8
	var negated_field_id TSFieldId
	var node, parent, parent_node TSNode
	var state QueryState
	var symbol TSSymbol
	var v11 int32
	var v9 bool
	var _ /* can_have_later_siblings_with_this_field at bp+50 */ uint8
	var _ /* child_state at bp+88 */ uintptr
	var _ /* field_id at bp+52 */ TSFieldId
	var _ /* has_later_named_siblings at bp+49 */ uint8
	var _ /* has_later_siblings at bp+48 */ uint8
	var _ /* i at bp+76 */ uint32
	var _ /* left_contains_right at bp+96 */ uint8
	var _ /* node_range at bp+24 */ TSRange
	var _ /* right_contains_left at bp+97 */ uint8
	var _ /* state at bp+80 */ uintptr
	var _ /* supertype_count at bp+72 */ uint32
	var _ /* supertypes at bp+54 */ [8]TSSymbol
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = child_step, copy1, copy_count, deleted_count, did_match, did_remove, end_index, has_supertype, i, i1, is_missing, is_named, j, j1, k, k1, k2, later_sibling_can_match, n, negated_field_id, negated_field_ids, next_step, next_step1, node, node_does_match, node_intersects_containing_range, node_intersects_range, node_is_error, node_within_containing_range, other_state, parent, parent_intersects_range, parent_is_error, parent_node, pattern, pattern1, skipped_wildcard_step, start_depth, start_depth1, state, state1, state3, step, step1, step2, step3, symbol, v11, v2, v3, v4, v9
	did_match = libc.BoolUint8(0 != 0)
	for {
		if (*TSQueryCursor)(unsafe.Pointer(self)).Fhalted != 0 {
			for (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize > uint32(0) {
				v3 = self + 40 + 8
				*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) - 1
				v2 = *(*uint32_t)(unsafe.Pointer(v3))
				state = *(*QueryState)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+40)).Fcontents + uintptr(v2)*16))
				capture_list_pool_release(tls, self+72, uint16(state.Fcapture_list_id))
			}
		}
		v3 = self + 200
		*(*uint32)(unsafe.Pointer(v3)) = *(*uint32)(unsafe.Pointer(v3)) + 1
		v4 = *(*uint32)(unsafe.Pointer(v3))
		if v4 == OP_COUNT_PER_QUERY_CALLBACK_CHECK {
			(*TSQueryCursor)(unsafe.Pointer(self)).Foperation_count = uint32(0)
		}
		if (*TSQueryCursor)(unsafe.Pointer(self)).Fquery_options != 0 && (*TSQueryCursorOptions)(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery_options)).Fprogress_callback != 0 {
			(*TSQueryCursor)(unsafe.Pointer(self)).Fquery_state.Fcurrent_byte_offset = ts_node_start_byte(tls, ts_tree_cursor_current_node(tls, self+8))
		}
		if did_match != 0 || (*TSQueryCursor)(unsafe.Pointer(self)).Fhalted != 0 || (*TSQueryCursor)(unsafe.Pointer(self)).Foperation_count == uint32(0) && ((*TSQueryCursor)(unsafe.Pointer(self)).Fquery_options != 0 && (*TSQueryCursorOptions)(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery_options)).Fprogress_callback != 0 && (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSQueryCursorOptions)(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery_options)).Fprogress_callback})))(tls, self+184) != 0) {
			return did_match
		}
		if (*TSQueryCursor)(unsafe.Pointer(self)).Fascending != 0 {
			if (*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node != 0 {
				deleted_count = uint32(0)
				i = uint32(0)
				n = (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize
				for {
					if !(i < n) {
						break
					}
					_ = libc.Uint64FromInt64(4)
					{
						if !(i < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+40)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+7839, __ccgo_ts+3070, int32(3805), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					state1 = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+40)).Fcontents + uintptr(i)*16
					_ = libc.Uint64FromInt64(4)
					{
						if !(uint32((*QueryState)(unsafe.Pointer(state1)).Fstep_index) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+7938, __ccgo_ts+3070, int32(3806), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					step = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*QueryState)(unsafe.Pointer(state1)).Fstep_index)*20
					if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) && (uint32((*QueryState)(unsafe.Pointer(state1)).Fstart_depth) > (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth || (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth == uint32(0)) {
						(*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self + 56)).Fcontents = _array__grow(tls, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+56)).Fcontents, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+56)).Fsize, self+56+12, uint32(1), libc.Uint64FromInt64(16))
						v3 = self + 56 + 8
						v2 = *(*uint32_t)(unsafe.Pointer(v3))
						*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
						*(*QueryState)(unsafe.Pointer((*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+56)).Fcontents + uintptr(v2)*16)) = *(*QueryState)(unsafe.Pointer(state1))
						did_match = libc.BoolUint8(1 != 0)
						deleted_count = deleted_count + 1
					} else {
						if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step)).Fdepth) != libc.Int32FromUint16(PATTERN_DONE_MARKER) && uint32((*QueryState)(unsafe.Pointer(state1)).Fstart_depth)+uint32((*QueryStep)(unsafe.Pointer(step)).Fdepth) > (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth {
							capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state1)).Fcapture_list_id))
							deleted_count = deleted_count + 1
						} else {
							if deleted_count > uint32(0) {
								_ = libc.Uint64FromInt64(4)
								{
									if !(i-deleted_count < (*struct {
										Fcontents uintptr
										Fsize     uint32_t
										Fcapacity uint32_t
									})(unsafe.Pointer(self+40)).Fsize) {
										libc.X__assert_fail(tls, __ccgo_ts+8283, __ccgo_ts+3070, int32(3839), uintptr(unsafe.Pointer(&__func__65)))
									}
								}
								*(*QueryState)(unsafe.Pointer((*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+40)).Fcontents + uintptr(i-deleted_count)*16)) = *(*QueryState)(unsafe.Pointer(state1))
							}
						}
					}
					goto _6
				_6:
					;
					i = i + 1
				}
				(*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize -= deleted_count
			}
			switch ts_tree_cursor_goto_next_sibling_internal(tls, self+8) {
			case int32(TreeCursorStepVisible):
				if !((*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node != 0) {
					(*TSQueryCursor)(unsafe.Pointer(self)).Fdepth = (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth + 1
					(*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node = libc.BoolUint8(1 != 0)
				}
				(*TSQueryCursor)(unsafe.Pointer(self)).Fascending = libc.BoolUint8(0 != 0)
			case int32(TreeCursorStepHidden):
				if (*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node != 0 {
					(*TSQueryCursor)(unsafe.Pointer(self)).Fdepth = (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth - 1
					(*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node = libc.BoolUint8(0 != 0)
				}
				(*TSQueryCursor)(unsafe.Pointer(self)).Fascending = libc.BoolUint8(0 != 0)
			default:
				if ts_tree_cursor_goto_parent(tls, self+8) != 0 {
					(*TSQueryCursor)(unsafe.Pointer(self)).Fdepth = (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth - 1
				} else {
					(*TSQueryCursor)(unsafe.Pointer(self)).Fhalted = libc.BoolUint8(1 != 0)
				}
			}
		} else {
			node = ts_tree_cursor_current_node(tls, self+8)
			parent_node = ts_tree_cursor_parent_node(tls, self+8)
			if v9 = ts_node_is_null(tls, parent_node) != 0; !v9 {
				*(*TSRange)(unsafe.Pointer(bp)) = TSRange{
					Fstart_point: ts_node_start_point(tls, parent_node),
					Fend_point:   ts_node_end_point(tls, parent_node),
					Fstart_byte:  ts_node_start_byte(tls, parent_node),
					Fend_byte:    ts_node_end_byte(tls, parent_node),
				}
			}
			parent_intersects_range = libc.BoolUint8(v9 || range_intersects(tls, bp, self+120) != 0)
			*(*TSRange)(unsafe.Pointer(bp + 24)) = TSRange{
				Fstart_point: ts_node_start_point(tls, node),
				Fend_point:   ts_node_end_point(tls, node),
				Fstart_byte:  ts_node_start_byte(tls, node),
				Fend_byte:    ts_node_end_byte(tls, node),
			}
			node_intersects_range = libc.BoolUint8(parent_intersects_range != 0 && range_intersects(tls, bp+24, self+120) != 0)
			node_intersects_containing_range = range_intersects(tls, bp+24, self+144)
			node_within_containing_range = range_within(tls, bp+24, self+144)
			if node_within_containing_range != 0 && (*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node != 0 {
				symbol = ts_node_symbol(tls, node)
				is_named = ts_node_is_named(tls, node)
				is_missing = ts_node_is_missing(tls, node)
				*(*TSFieldId)(unsafe.Pointer(bp + 52)) = uint16(0)
				*(*[8]TSSymbol)(unsafe.Pointer(bp + 54)) = [8]TSSymbol{}
				*(*uint32)(unsafe.Pointer(bp + 72)) = uint32(8)
				ts_tree_cursor_current_status(tls, self+8, bp+52, bp+48, bp+49, bp+50, bp+54, bp+72)
				node_is_error = libc.BoolUint8(libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))))
				parent_is_error = libc.BoolUint8(!(ts_node_is_null(tls, parent_node) != 0) && libc.Int32FromUint16(ts_node_symbol(tls, parent_node)) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))))
				if !(node_is_error != 0) {
					i1 = uint32(0)
					for {
						if !(i1 < uint32((*TSQuery)(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery)).Fwildcard_root_pattern_count)) {
							break
						}
						_ = libc.Uint64FromInt64(4)
						{
							if !(i1 < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+96)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8337, __ccgo_ts+3070, int32(3934), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						pattern = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+96)).Fcontents + uintptr(i1)*6
						_ = libc.Uint64FromInt64(4)
						{
							if !(uint32((*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index) < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8075, __ccgo_ts+3070, int32(3938), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						step1 = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer(pattern)).Fstep_index)*20
						start_depth = (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth - uint32((*QueryStep)(unsafe.Pointer(step1)).Fdepth)
						if (*PatternEntry)(unsafe.Pointer(pattern)).Fis_rooted != 0 {
							v11 = libc.Int32FromUint8(node_intersects_range)
						} else {
							v11 = libc.BoolInt32(parent_intersects_range != 0 && !(parent_is_error != 0))
						}
						if v11 != 0 && (!((*QueryStep)(unsafe.Pointer(step1)).Ffield != 0) || libc.Int32FromUint16(*(*TSFieldId)(unsafe.Pointer(bp + 52))) == libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step1)).Ffield)) && (!((*QueryStep)(unsafe.Pointer(step1)).Fsupertype_symbol != 0) || *(*uint32)(unsafe.Pointer(bp + 72)) > uint32(0)) && start_depth <= (*TSQueryCursor)(unsafe.Pointer(self)).Fmax_start_depth {
							ts_query_cursor__add_state(tls, self, pattern)
						}
						goto _10
					_10:
						;
						i1 = i1 + 1
					}
				}
				if ts_query__pattern_map_search(tls, (*TSQueryCursor)(unsafe.Pointer(self)).Fquery, symbol, bp+76) != 0 {
					_ = libc.Uint64FromInt64(4)
					{
						if !(*(*uint32)(unsafe.Pointer(bp + 76)) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+96)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8337, __ccgo_ts+3070, int32(3956), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					pattern1 = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+96)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 76)))*6
					_ = libc.Uint64FromInt64(4)
					{
						if !(uint32((*PatternEntry)(unsafe.Pointer(pattern1)).Fstep_index) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8075, __ccgo_ts+3070, int32(3958), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					step2 = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer(pattern1)).Fstep_index)*20
					start_depth1 = (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth - uint32((*QueryStep)(unsafe.Pointer(step2)).Fdepth)
					for cond := true; cond; cond = libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step2)).Fsymbol) == libc.Int32FromUint16(symbol) {
						if (*PatternEntry)(unsafe.Pointer(pattern1)).Fis_rooted != 0 {
							v11 = libc.Int32FromUint8(node_intersects_range)
						} else {
							v11 = libc.BoolInt32(parent_intersects_range != 0 && !(parent_is_error != 0))
						}
						if v11 != 0 && (!((*QueryStep)(unsafe.Pointer(step2)).Ffield != 0) || libc.Int32FromUint16(*(*TSFieldId)(unsafe.Pointer(bp + 52))) == libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step2)).Ffield)) && start_depth1 <= (*TSQueryCursor)(unsafe.Pointer(self)).Fmax_start_depth {
							ts_query_cursor__add_state(tls, self, pattern1)
						}
						*(*uint32)(unsafe.Pointer(bp + 76)) = *(*uint32)(unsafe.Pointer(bp + 76)) + 1
						if *(*uint32)(unsafe.Pointer(bp + 76)) == (*TSQuery)(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery)).Fpattern_map.Fsize {
							break
						}
						_ = libc.Uint64FromInt64(4)
						{
							if !(*(*uint32)(unsafe.Pointer(bp + 76)) < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+96)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8337, __ccgo_ts+3070, int32(3976), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						pattern1 = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+96)).Fcontents + uintptr(*(*uint32)(unsafe.Pointer(bp + 76)))*6
						_ = libc.Uint64FromInt64(4)
						{
							if !(uint32((*PatternEntry)(unsafe.Pointer(pattern1)).Fstep_index) < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8075, __ccgo_ts+3070, int32(3977), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						step2 = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*PatternEntry)(unsafe.Pointer(pattern1)).Fstep_index)*20
					}
				}
				j = uint32(0)
				copy_count = uint32(0)
				for {
					if !(j < (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize) {
						break
					}
					_ = libc.Uint64FromInt64(4)
					{
						if !(j < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+40)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8387, __ccgo_ts+3070, int32(3983), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					*(*uintptr)(unsafe.Pointer(bp + 80)) = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+40)).Fcontents + uintptr(j)*16
					_ = libc.Uint64FromInt64(4)
					{
						if !(uint32((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstep_index) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+7938, __ccgo_ts+3070, int32(3984), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					step3 = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstep_index)*20
					libc.SetBitFieldPtr16Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+14, libc.BoolUint8(0 != 0), 13, 0x2000)
					copy_count = uint32(0)
					if uint32((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstart_depth)+uint32((*QueryStep)(unsafe.Pointer(step3)).Fdepth) != (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth {
						goto _13
					}
					node_does_match = libc.BoolUint8(0 != 0)
					if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Fsymbol) == libc.Int32FromUint16(WILDCARD_SYMBOL) {
						if int32(*(*uint8)(unsafe.Pointer(step3 + 19))&0x2>>1) != 0 {
							node_does_match = is_missing
						} else {
							node_does_match = libc.BoolUint8(!(node_is_error != 0) && (is_named != 0 || !(int32(*(*uint8)(unsafe.Pointer(step3 + 18))&0x1>>0) != 0)))
						}
					} else {
						node_does_match = libc.BoolUint8(libc.Int32FromUint16(symbol) == libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Fsymbol) && (!(int32(*(*uint8)(unsafe.Pointer(step3 + 19))&0x2>>1) != 0) || is_missing != 0))
					}
					later_sibling_can_match = *(*uint8)(unsafe.Pointer(bp + 48))
					if int32(*(*uint8)(unsafe.Pointer(step3 + 18))&0x2>>1) != 0 && is_named != 0 || int32(uint8(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)) + 14))&0x1000>>12)) != 0 {
						later_sibling_can_match = libc.BoolUint8(0 != 0)
					}
					if int32(*(*uint8)(unsafe.Pointer(step3 + 18))&0x4>>2) != 0 && *(*uint8)(unsafe.Pointer(bp + 49)) != 0 {
						node_does_match = libc.BoolUint8(0 != 0)
					}
					if (*QueryStep)(unsafe.Pointer(step3)).Fsupertype_symbol != 0 {
						has_supertype = libc.BoolUint8(0 != 0)
						k = uint32(0)
						for {
							if !(k < *(*uint32)(unsafe.Pointer(bp + 72))) {
								break
							}
							if libc.Int32FromUint16((*(*[8]TSSymbol)(unsafe.Pointer(bp + 54)))[k]) == libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Fsupertype_symbol) {
								has_supertype = libc.BoolUint8(1 != 0)
								break
							}
							goto _14
						_14:
							;
							k = k + 1
						}
						if !(has_supertype != 0) {
							node_does_match = libc.BoolUint8(0 != 0)
						}
					}
					if (*QueryStep)(unsafe.Pointer(step3)).Ffield != 0 {
						if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Ffield) == libc.Int32FromUint16(*(*TSFieldId)(unsafe.Pointer(bp + 52))) {
							if !(*(*uint8)(unsafe.Pointer(bp + 50)) != 0) {
								later_sibling_can_match = libc.BoolUint8(0 != 0)
							}
						} else {
							node_does_match = libc.BoolUint8(0 != 0)
						}
					}
					if (*QueryStep)(unsafe.Pointer(step3)).Fnegated_field_list_id != 0 {
						_ = libc.Uint64FromInt64(4)
						{
							if !(uint32((*QueryStep)(unsafe.Pointer(step3)).Fnegated_field_list_id) < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+160)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8425, __ccgo_ts+3070, int32(4033), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						negated_field_ids = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+160)).Fcontents + uintptr((*QueryStep)(unsafe.Pointer(step3)).Fnegated_field_list_id)*2
						for {
							negated_field_id = *(*TSFieldId)(unsafe.Pointer(negated_field_ids))
							if negated_field_id != 0 {
								negated_field_ids += 2
								if ts_node_child_by_field_id(tls, node, negated_field_id).Fid != 0 {
									node_does_match = libc.BoolUint8(0 != 0)
									break
								}
							} else {
								break
							}
							goto _15
						_15:
						}
					}
					if !(node_does_match != 0) {
						if !(later_sibling_can_match != 0) {
							capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fcapture_list_id))
							_array__erase(tls, (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), j)
							j = j - 1
						}
						goto _13
					}
					if later_sibling_can_match != 0 && (int32(*(*uint8)(unsafe.Pointer(step3 + 18))&0x40>>6) != 0 || ts_query__step_is_fallible(tls, (*TSQueryCursor)(unsafe.Pointer(self)).Fquery, (*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstep_index) != 0) {
						if ts_query_cursor__copy_state(tls, self, bp+80) != 0 {
							copy_count = copy_count + 1
						}
					}
					if int32(uint8(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)) + 14))&0x8000>>15)) != 0 {
						parent = ts_tree_cursor_parent_node(tls, self+8)
						if ts_node_is_null(tls, parent) != 0 {
							libc.SetBitFieldPtr16Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+14, libc.BoolUint8(1 != 0), 14, 0x4000)
						} else {
							libc.SetBitFieldPtr16Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+14, libc.BoolUint8(0 != 0), 15, 0x8000)
							skipped_wildcard_step = step3
							for cond := true; cond; cond = int32(*(*uint8)(unsafe.Pointer(skipped_wildcard_step + 18))&0x10>>4) != 0 || int32(*(*uint8)(unsafe.Pointer(skipped_wildcard_step + 18))&0x8>>3) != 0 || libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(skipped_wildcard_step)).Fdepth) > 0 {
								skipped_wildcard_step -= 20
							}
							if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(skipped_wildcard_step + 6))) != libc.Int32FromUint16(NONE) {
								ts_query_cursor__capture(tls, self, *(*uintptr)(unsafe.Pointer(bp + 80)), skipped_wildcard_step, parent)
							}
						}
					}
					if libc.Int32FromUint16(*(*uint16_t)(unsafe.Pointer(step3 + 6))) != libc.Int32FromUint16(NONE) {
						ts_query_cursor__capture(tls, self, *(*uintptr)(unsafe.Pointer(bp + 80)), step3, node)
					}
					if int32(uint8(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)) + 14))&0x4000>>14)) != 0 {
						_array__erase(tls, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), j)
						j = j - 1
						goto _13
					}
					(*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstep_index = (*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstep_index + 1
					_ = libc.Uint64FromInt64(4)
					{
						if !(uint32((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstep_index) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+7938, __ccgo_ts+3070, int32(4134), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					next_step = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fstep_index)*20
					if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(step3)).Fsymbol) == libc.Int32FromUint16(WILDCARD_SYMBOL) && !(int32(*(*uint8)(unsafe.Pointer(step3 + 18))&0x1>>0) != 0) && int32(*(*uint8)(unsafe.Pointer(next_step + 18))&0x2>>1) != 0 {
						libc.SetBitFieldPtr16Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+14, libc.BoolUint8(1 != 0), 12, 0x1000)
					} else {
						libc.SetBitFieldPtr16Uint8(*(*uintptr)(unsafe.Pointer(bp + 80))+14, libc.BoolUint8(0 != 0), 12, 0x1000)
					}
					if stop_on_definite_step != 0 && int32(*(*uint8)(unsafe.Pointer(next_step + 18))&0x80>>7) != 0 {
						did_match = libc.BoolUint8(1 != 0)
					}
					end_index = j + uint32(1)
					k1 = j
					for {
						if !(k1 < end_index) {
							break
						}
						_ = libc.Uint64FromInt64(4)
						{
							if !(k1 < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+40)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8504, __ccgo_ts+3070, int32(4157), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						*(*uintptr)(unsafe.Pointer(bp + 88)) = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+40)).Fcontents + uintptr(k1)*16
						_ = libc.Uint64FromInt64(4)
						{
							if !(uint32((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Fstep_index) < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8542, __ccgo_ts+3070, int32(4158), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						child_step = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Fstep_index)*20
						if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(child_step)).Falternative_index) != libc.Int32FromUint16(NONE) {
							if int32(*(*uint8)(unsafe.Pointer(child_step + 18))&0x10>>4) != 0 {
								(*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Fstep_index = (*QueryStep)(unsafe.Pointer(child_step)).Falternative_index
								k1 = k1 - 1
								goto _16
							}
							if int32(*(*uint8)(unsafe.Pointer(child_step + 18))&0x8>>3) != 0 {
								(*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Fstep_index = (*QueryState)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 88)))).Fstep_index + 1
								k1 = k1 - 1
							}
							copy1 = ts_query_cursor__copy_state(tls, self, bp+88)
							if copy1 != 0 {
								end_index = end_index + 1
								copy_count = copy_count + 1
								(*QueryState)(unsafe.Pointer(copy1)).Fstep_index = (*QueryStep)(unsafe.Pointer(child_step)).Falternative_index
								if int32(*(*uint8)(unsafe.Pointer(child_step + 18))&0x20>>5) != 0 {
									libc.SetBitFieldPtr16Uint8(copy1+14, libc.BoolUint8(1 != 0), 12, 0x1000)
								}
							}
						}
						goto _16
					_16:
						;
						k1 = k1 + 1
					}
					goto _13
				_13:
					;
					j = j + (uint32(1) + copy_count)
				}
				j1 = uint32(0)
				for {
					if !(j1 < (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize) {
						break
					}
					_ = libc.Uint64FromInt64(4)
					{
						if !(j1 < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+40)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8387, __ccgo_ts+3070, int32(4199), uintptr(unsafe.Pointer(&__func__65)))
						}
					}
					state3 = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+40)).Fcontents + uintptr(j1)*16
					if int32(uint8(*(*uint16)(unsafe.Pointer(state3 + 14))&0x4000>>14)) != 0 {
						_array__erase(tls, (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), j1)
						j1 = j1 - 1
						goto _17
					}
					did_remove = libc.BoolUint8(0 != 0)
					k2 = j1 + uint32(1)
					for {
						if !(k2 < (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize) {
							break
						}
						_ = libc.Uint64FromInt64(4)
						{
							if !(k2 < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer(self+40)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+8504, __ccgo_ts+3070, int32(4211), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						other_state = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+40)).Fcontents + uintptr(k2)*16
						if libc.Int32FromUint16((*QueryState)(unsafe.Pointer(other_state)).Fstart_depth) != libc.Int32FromUint16((*QueryState)(unsafe.Pointer(state3)).Fstart_depth) || libc.Int32FromUint16((*QueryState)(unsafe.Pointer(other_state)).Fpattern_index) != libc.Int32FromUint16((*QueryState)(unsafe.Pointer(state3)).Fpattern_index) {
							break
						}
						ts_query_cursor__compare_captures(tls, self, state3, other_state, bp+96, bp+97)
						if *(*uint8)(unsafe.Pointer(bp + 96)) != 0 {
							if libc.Int32FromUint16((*QueryState)(unsafe.Pointer(state3)).Fstep_index) == libc.Int32FromUint16((*QueryState)(unsafe.Pointer(other_state)).Fstep_index) {
								capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(other_state)).Fcapture_list_id))
								_array__erase(tls, (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), k2)
								k2 = k2 - 1
								goto _18
							}
							libc.SetBitFieldPtr16Uint8(other_state+14, libc.BoolUint8(1 != 0), 13, 0x2000)
						}
						if *(*uint8)(unsafe.Pointer(bp + 97)) != 0 {
							if libc.Int32FromUint16((*QueryState)(unsafe.Pointer(state3)).Fstep_index) == libc.Int32FromUint16((*QueryState)(unsafe.Pointer(other_state)).Fstep_index) {
								capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state3)).Fcapture_list_id))
								_array__erase(tls, (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), j1)
								j1 = j1 - 1
								did_remove = libc.BoolUint8(1 != 0)
								break
							}
							libc.SetBitFieldPtr16Uint8(state3+14, libc.BoolUint8(1 != 0), 13, 0x2000)
						}
						goto _18
					_18:
						;
						k2 = k2 + 1
					}
					if !(did_remove != 0) {
						_ = libc.Uint64FromInt64(4)
						{
							if !(uint32((*QueryState)(unsafe.Pointer(state3)).Fstep_index) < (*struct {
								Fcontents uintptr
								Fsize     uint32_t
								Fcapacity uint32_t
							})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fsize) {
								libc.X__assert_fail(tls, __ccgo_ts+7938, __ccgo_ts+3070, int32(4271), uintptr(unsafe.Pointer(&__func__65)))
							}
						}
						next_step1 = (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer((*TSQueryCursor)(unsafe.Pointer(self)).Fquery+80)).Fcontents + uintptr((*QueryState)(unsafe.Pointer(state3)).Fstep_index)*20
						if libc.Int32FromUint16((*QueryStep)(unsafe.Pointer(next_step1)).Fdepth) == libc.Int32FromUint16(PATTERN_DONE_MARKER) {
							if int32(uint8(*(*uint16)(unsafe.Pointer(state3 + 14))&0x2000>>13)) != 0 {
							} else {
								(*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self + 56)).Fcontents = _array__grow(tls, (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+56)).Fcontents, (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+56)).Fsize, self+56+12, uint32(1), libc.Uint64FromInt64(16))
								v3 = self + 56 + 8
								v2 = *(*uint32_t)(unsafe.Pointer(v3))
								*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
								*(*QueryState)(unsafe.Pointer((*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+56)).Fcontents + uintptr(v2)*16)) = *(*QueryState)(unsafe.Pointer(state3))
								_array__erase(tls, (*struct {
									Fcontents uintptr
									Fsize     uint32_t
									Fcapacity uint32_t
								})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), libc.Uint32FromInt64((int64(state3)-int64((*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fcontents))/16))
								did_match = libc.BoolUint8(1 != 0)
								j1 = j1 - 1
							}
						}
					}
					goto _17
				_17:
					;
					j1 = j1 + 1
				}
			}
			if node_intersects_containing_range != 0 && ts_query_cursor__should_descend(tls, self, node_intersects_range) != 0 {
				switch ts_tree_cursor_goto_first_child_internal(tls, self+8) {
				case int32(TreeCursorStepVisible):
					(*TSQueryCursor)(unsafe.Pointer(self)).Fdepth = (*TSQueryCursor)(unsafe.Pointer(self)).Fdepth + 1
					(*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node = libc.BoolUint8(1 != 0)
					goto _1
				case int32(TreeCursorStepHidden):
					(*TSQueryCursor)(unsafe.Pointer(self)).Fon_visible_node = libc.BoolUint8(0 != 0)
					goto _1
				default:
					break
				}
			}
			(*TSQueryCursor)(unsafe.Pointer(self)).Fascending = libc.BoolUint8(1 != 0)
		}
		goto _1
	_1:
	}
	return r
}

var __func__65 = [25]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', '_', 'a', 'd', 'v', 'a', 'n', 'c', 'e'}

func ts_query_cursor_next_match(tls *libc.TLS, self uintptr, match uintptr) (r uint8) {
	var captures, state, v2 uintptr
	var v1 uint32_t
	_, _, _, _ = captures, state, v1, v2
	if (*TSQueryCursor)(unsafe.Pointer(self)).Ffinished_states.Fsize == uint32(0) {
		if !(ts_query_cursor__advance(tls, self, libc.BoolUint8(0 != 0)) != 0) {
			return libc.BoolUint8(0 != 0)
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+56)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8608, __ccgo_ts+3070, int32(4316), uintptr(unsafe.Pointer(&__func__66)))
		}
	}
	state = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 56)).Fcontents
	if (*QueryState)(unsafe.Pointer(state)).Fid == uint32(4294967295) {
		v2 = self + 168
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		(*QueryState)(unsafe.Pointer(state)).Fid = v1
	}
	(*TSQueryMatch)(unsafe.Pointer(match)).Fid = (*QueryState)(unsafe.Pointer(state)).Fid
	(*TSQueryMatch)(unsafe.Pointer(match)).Fpattern_index = (*QueryState)(unsafe.Pointer(state)).Fpattern_index
	captures = capture_list_pool_get(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
	(*TSQueryMatch)(unsafe.Pointer(match)).Fcaptures = (*CaptureList)(unsafe.Pointer(captures)).Fcontents
	(*TSQueryMatch)(unsafe.Pointer(match)).Fcapture_count = uint16((*CaptureList)(unsafe.Pointer(captures)).Fsize)
	capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
	_array__erase(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+56)).Fcontents, self+56+8, libc.Uint64FromInt64(16), uint32(0))
	return libc.BoolUint8(1 != 0)
}

var __func__66 = [27]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'n', 'e', 'x', 't', '_', 'm', 'a', 't', 'c', 'h'}

func ts_query_cursor_remove_match(tls *libc.TLS, self uintptr, match_id uint32_t) {
	var i, i1 uint32
	var state, state1 uintptr
	_, _, _, _ = i, i1, state, state1
	i = uint32(0)
	for {
		if !(i < (*TSQueryCursor)(unsafe.Pointer(self)).Ffinished_states.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+56)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+8655, __ccgo_ts+3070, int32(4336), uintptr(unsafe.Pointer(&__func__67)))
			}
		}
		state = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+56)).Fcontents + uintptr(i)*16
		if (*QueryState)(unsafe.Pointer(state)).Fid == match_id {
			capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
			_array__erase(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+56)).Fcontents, self+56+8, libc.Uint64FromInt64(16), i)
			return
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	i1 = uint32(0)
	for {
		if !(i1 < (*TSQueryCursor)(unsafe.Pointer(self)).Fstates.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i1 < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+40)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+7839, __ccgo_ts+3070, int32(4350), uintptr(unsafe.Pointer(&__func__67)))
			}
		}
		state1 = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+40)).Fcontents + uintptr(i1)*16
		if (*QueryState)(unsafe.Pointer(state1)).Fid == match_id {
			capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state1)).Fcapture_list_id))
			_array__erase(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), i1)
			return
		}
		goto _2
	_2:
		;
		i1 = i1 + 1
	}
}

var __func__67 = [29]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'r', 'e', 'm', 'o', 'v', 'e', '_', 'm', 'a', 't', 'c', 'h'}

func ts_query_cursor_next_capture(tls *libc.TLS, self uintptr, match uintptr, capture_index uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var captures, captures1, first_finished_state, state, state1, v4 uintptr
	var first_finished_capture_byte, first_finished_pattern_index, node_start_byte, v3 uint32_t
	var found_unfinished_state, node_follows_range, node_outside_of_range, node_precedes_range uint8
	var i uint32
	var node TSNode
	var _ /* first_unfinished_capture_byte at bp+0 */ uint32_t
	var _ /* first_unfinished_pattern_index at bp+4 */ uint32_t
	var _ /* first_unfinished_state_index at bp+8 */ uint32_t
	var _ /* first_unfinished_state_is_definite at bp+12 */ uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = captures, captures1, first_finished_capture_byte, first_finished_pattern_index, first_finished_state, found_unfinished_state, i, node, node_follows_range, node_outside_of_range, node_precedes_range, node_start_byte, state, state1, v3, v4
	for {
		*(*uint8)(unsafe.Pointer(bp + 12)) = libc.BoolUint8(0 != 0)
		found_unfinished_state = ts_query_cursor__first_in_progress_capture(tls, self, bp+8, bp, bp+4, bp+12)
		first_finished_state = libc.UintptrFromInt32(0)
		first_finished_capture_byte = *(*uint32_t)(unsafe.Pointer(bp))
		first_finished_pattern_index = *(*uint32_t)(unsafe.Pointer(bp + 4))
		i = uint32(0)
		for {
			if !(i < (*TSQueryCursor)(unsafe.Pointer(self)).Ffinished_states.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+56)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+8655, __ccgo_ts+3070, int32(4390), uintptr(unsafe.Pointer(&__func__68)))
				}
			}
			state = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+56)).Fcontents + uintptr(i)*16
			captures = capture_list_pool_get(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
			if libc.Uint32FromInt32(int32(*(*uint16)(unsafe.Pointer(state + 14))&0xfff>>0)) >= (*CaptureList)(unsafe.Pointer(captures)).Fsize {
				capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state)).Fcapture_list_id))
				_array__erase(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+56)).Fcontents, self+56+8, libc.Uint64FromInt64(16), i)
				goto _2
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(libc.Uint32FromInt32(int32(*(*uint16)(unsafe.Pointer(state + 14))&0xfff>>0)) < (*CaptureList)(unsafe.Pointer(captures)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+7877, __ccgo_ts+3070, int32(4406), uintptr(unsafe.Pointer(&__func__68)))
				}
			}
			node = (*TSQueryCapture)(unsafe.Pointer((*CaptureList)(unsafe.Pointer(captures)).Fcontents + uintptr(int32(*(*uint16)(unsafe.Pointer(state + 14))&0xfff>>0))*40)).Fnode
			node_precedes_range = libc.BoolUint8(ts_node_end_byte(tls, node) <= (*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fstart_byte || point_lte(tls, ts_node_end_point(tls, node), (*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fstart_point) != 0)
			node_follows_range = libc.BoolUint8(ts_node_start_byte(tls, node) >= (*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fend_byte || point_gte(tls, ts_node_start_point(tls, node), (*TSQueryCursor)(unsafe.Pointer(self)).Fincluded_range.Fend_point) != 0)
			node_outside_of_range = libc.BoolUint8(node_precedes_range != 0 || node_follows_range != 0)
			if node_outside_of_range != 0 {
				libc.PostIncBitFieldPtr16Uint16(state+14, 1, 12, 0, 0xfff)
				goto _2
			}
			node_start_byte = ts_node_start_byte(tls, node)
			if node_start_byte < first_finished_capture_byte || node_start_byte == first_finished_capture_byte && uint32((*QueryState)(unsafe.Pointer(state)).Fpattern_index) < first_finished_pattern_index {
				first_finished_state = state
				first_finished_capture_byte = node_start_byte
				first_finished_pattern_index = uint32((*QueryState)(unsafe.Pointer(state)).Fpattern_index)
			}
			i = i + 1
			goto _2
		_2:
		}
		if first_finished_state != 0 {
			state1 = first_finished_state
		} else {
			if *(*uint8)(unsafe.Pointer(bp + 12)) != 0 {
				_ = libc.Uint64FromInt64(4)
				{
					if !(*(*uint32_t)(unsafe.Pointer(bp + 8)) < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+40)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+8702, __ccgo_ts+3070, int32(4446), uintptr(unsafe.Pointer(&__func__68)))
					}
				}
				state1 = (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+40)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 8)))*16
			} else {
				state1 = libc.UintptrFromInt32(0)
			}
		}
		if state1 != 0 {
			if (*QueryState)(unsafe.Pointer(state1)).Fid == uint32(4294967295) {
				v4 = self + 168
				v3 = *(*uint32_t)(unsafe.Pointer(v4))
				*(*uint32_t)(unsafe.Pointer(v4)) = *(*uint32_t)(unsafe.Pointer(v4)) + 1
				(*QueryState)(unsafe.Pointer(state1)).Fid = v3
			}
			(*TSQueryMatch)(unsafe.Pointer(match)).Fid = (*QueryState)(unsafe.Pointer(state1)).Fid
			(*TSQueryMatch)(unsafe.Pointer(match)).Fpattern_index = (*QueryState)(unsafe.Pointer(state1)).Fpattern_index
			captures1 = capture_list_pool_get(tls, self+72, uint16((*QueryState)(unsafe.Pointer(state1)).Fcapture_list_id))
			(*TSQueryMatch)(unsafe.Pointer(match)).Fcaptures = (*CaptureList)(unsafe.Pointer(captures1)).Fcontents
			(*TSQueryMatch)(unsafe.Pointer(match)).Fcapture_count = uint16((*CaptureList)(unsafe.Pointer(captures1)).Fsize)
			*(*uint32_t)(unsafe.Pointer(capture_index)) = libc.Uint32FromInt32(int32(*(*uint16)(unsafe.Pointer(state1 + 14)) & 0xfff >> 0))
			libc.PostIncBitFieldPtr16Uint16(state1+14, 1, 12, 0, 0xfff)
			return libc.BoolUint8(1 != 0)
		}
		if capture_list_pool_is_empty(tls, self+72) != 0 && found_unfinished_state != 0 {
			_ = libc.Uint64FromInt64(4)
			{
				if !(*(*uint32_t)(unsafe.Pointer(bp + 8)) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+40)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+8702, __ccgo_ts+3070, int32(4475), uintptr(unsafe.Pointer(&__func__68)))
				}
			}
			capture_list_pool_release(tls, self+72, uint16((*QueryState)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+40)).Fcontents+uintptr(*(*uint32_t)(unsafe.Pointer(bp + 8)))*16)).Fcapture_list_id))
			_array__erase(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+40)).Fcontents, self+40+8, libc.Uint64FromInt64(16), *(*uint32_t)(unsafe.Pointer(bp + 8)))
		}
		if !(ts_query_cursor__advance(tls, self, libc.BoolUint8(1 != 0)) != 0) && (*TSQueryCursor)(unsafe.Pointer(self)).Ffinished_states.Fsize == uint32(0) {
			return libc.BoolUint8(0 != 0)
		}
		goto _1
	_1:
	}
	return r
}

var __func__68 = [29]int8{'t', 's', '_', 'q', 'u', 'e', 'r', 'y', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'n', 'e', 'x', 't', '_', 'c', 'a', 'p', 't', 'u', 'r', 'e'}

func ts_query_cursor_set_max_start_depth(tls *libc.TLS, self uintptr, max_start_depth uint32_t) {
	(*TSQueryCursor)(unsafe.Pointer(self)).Fmax_start_depth = max_start_depth
}

type StackNode = struct {
	Fstate              TSStateId
	Fposition           Length
	Flinks              [8]StackLink
	Flink_count         uint16
	Fref_count          uint32_t
	Ferror_cost         uint32
	Fnode_count         uint32
	Fdynamic_precedence int32
}

type StackLink = struct {
	Fnode       uintptr
	Fsubtree    Subtree
	Fis_pending uint8
}

type StackIterator = struct {
	Fnode          uintptr
	Fsubtrees      SubtreeArray
	Fsubtree_count uint32_t
	Fis_pending    uint8
}

type StackNodeArray = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type StackStatus = int32

const StackStatusActive = 0
const StackStatusPaused = 1
const StackStatusHalted = 2

type StackHead = struct {
	Fnode                     uintptr
	Fsummary                  uintptr
	Fnode_count_at_last_error uint32
	Flast_external_token      Subtree
	Flookahead_when_paused    Subtree
	Fstatus                   StackStatus
}

type StackAction = uint32

const StackActionNone = 0
const StackActionStop = 1
const StackActionPop = 2

type StackCallback = uintptr

func stack_node_retain(tls *libc.TLS, self uintptr) {
	if !(self != 0) {
		return
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !((*StackNode)(unsafe.Pointer(self)).Fref_count > libc.Uint32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+8767, __ccgo_ts+8787, int32(85), uintptr(unsafe.Pointer(&__func__69)))
		}
	}
	(*StackNode)(unsafe.Pointer(self)).Fref_count = (*StackNode)(unsafe.Pointer(self)).Fref_count + 1
	_ = libc.Uint64FromInt64(4)
	{
		if !((*StackNode)(unsafe.Pointer(self)).Fref_count != libc.Uint32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+8829, __ccgo_ts+8787, int32(87), uintptr(unsafe.Pointer(&__func__69)))
		}
	}
}

var __func__69 = [18]int8{'s', 't', 'a', 'c', 'k', '_', 'n', 'o', 'd', 'e', '_', 'r', 'e', 't', 'a', 'i', 'n'}

func stack_node_release(tls *libc.TLS, self uintptr, pool uintptr, subtree_pool uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var first_predecessor, v3 uintptr
	var i uint32
	var v2 uint32_t
	var _ /* link at bp+0 */ StackLink
	var _ /* link at bp+24 */ StackLink
	_, _, _, _ = first_predecessor, i, v2, v3
	goto recur
recur:
	;
	_ = libc.Uint64FromInt64(4)
	{
		if !((*StackNode)(unsafe.Pointer(self)).Fref_count != libc.Uint32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+8829, __ccgo_ts+8787, int32(96), uintptr(unsafe.Pointer(&__func__70)))
		}
	}
	(*StackNode)(unsafe.Pointer(self)).Fref_count = (*StackNode)(unsafe.Pointer(self)).Fref_count - 1
	if (*StackNode)(unsafe.Pointer(self)).Fref_count > uint32(0) {
		return
	}
	first_predecessor = libc.UintptrFromInt32(0)
	if libc.Int32FromUint16((*StackNode)(unsafe.Pointer(self)).Flink_count) > 0 {
		i = libc.Uint32FromInt32(libc.Int32FromUint16((*StackNode)(unsafe.Pointer(self)).Flink_count) - int32(1))
		for {
			if !(i > uint32(0)) {
				break
			}
			*(*StackLink)(unsafe.Pointer(bp)) = StackLink{}
			*(*struct {
				Fnode    uintptr
				Fsubtree struct {
					Fptr  [0]uintptr
					Fdata SubtreeInlineData
				}
				Fis_pending uint8
			})(unsafe.Pointer(bp)) = *(*StackLink)(unsafe.Pointer(self + 16 + uintptr(i)*24))
			if *(*uintptr)(unsafe.Pointer(bp + 8)) != 0 {
				ts_subtree_release(tls, subtree_pool, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree)
			}
			stack_node_release(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fnode, pool, subtree_pool)
			goto _1
		_1:
			;
			i = i - 1
		}
		*(*StackLink)(unsafe.Pointer(bp + 24)) = StackLink{}
		*(*struct {
			Fnode    uintptr
			Fsubtree struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Fis_pending uint8
		})(unsafe.Pointer(bp + 24)) = *(*StackLink)(unsafe.Pointer(self + 16))
		if *(*uintptr)(unsafe.Pointer(bp + 24 + 8)) != 0 {
			ts_subtree_release(tls, subtree_pool, (*(*StackLink)(unsafe.Pointer(bp + 24))).Fsubtree)
		}
		first_predecessor = (*(*StackLink)(unsafe.Pointer(self + 16))).Fnode
	}
	if (*StackNodeArray)(unsafe.Pointer(pool)).Fsize < uint32(50) {
		(*StackNodeArray)(unsafe.Pointer(pool)).Fcontents = _array__grow(tls, (*StackNodeArray)(unsafe.Pointer(pool)).Fcontents, (*StackNodeArray)(unsafe.Pointer(pool)).Fsize, pool+12, uint32(1), libc.Uint64FromInt64(8))
		v3 = pool + 8
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
		*(*uintptr)(unsafe.Pointer((*StackNodeArray)(unsafe.Pointer(pool)).Fcontents + uintptr(v2)*8)) = self
	} else {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, self)
	}
	if first_predecessor != 0 {
		self = first_predecessor
		goto recur
	}
}

var __func__70 = [19]int8{'s', 't', 'a', 'c', 'k', '_', 'n', 'o', 'd', 'e', '_', 'r', 'e', 'l', 'e', 'a', 's', 'e'}

func stack__subtree_node_count(tls *libc.TLS, subtree Subtree) (r uint32_t) {
	var count uint32_t
	_ = count
	count = ts_subtree_visible_descendant_count(tls, subtree)
	if ts_subtree_visible(tls, subtree) != 0 {
		count = count + 1
	}
	if libc.Int32FromUint16(ts_subtree_symbol(tls, subtree)) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
		count = count + 1
	}
	return count
}

func stack_node_new(tls *libc.TLS, previous_node uintptr, _subtree Subtree, is_pending uint8, state TSStateId, pool uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _subtree
	var node, v1, v3 uintptr
	var v2 uint32_t
	_, _, _, _ = node, v1, v2, v3
	if (*StackNodeArray)(unsafe.Pointer(pool)).Fsize > uint32(0) {
		v3 = pool + 8
		*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) - 1
		v2 = *(*uint32_t)(unsafe.Pointer(v3))
		v1 = *(*uintptr)(unsafe.Pointer((*StackNodeArray)(unsafe.Pointer(pool)).Fcontents + uintptr(v2)*8))
	} else {
		v1 = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(232))
	}
	node = v1
	*(*StackNode)(unsafe.Pointer(node)) = StackNode{
		Fstate:     state,
		Fref_count: uint32(1),
	}
	if previous_node != 0 {
		(*StackNode)(unsafe.Pointer(node)).Flink_count = uint16(1)
		*(*StackLink)(unsafe.Pointer(node + 16)) = StackLink{
			Fnode:       previous_node,
			Fsubtree:    *(*Subtree)(unsafe.Pointer(bp)),
			Fis_pending: is_pending,
		}
		(*StackNode)(unsafe.Pointer(node)).Fposition = (*StackNode)(unsafe.Pointer(previous_node)).Fposition
		(*StackNode)(unsafe.Pointer(node)).Ferror_cost = (*StackNode)(unsafe.Pointer(previous_node)).Ferror_cost
		(*StackNode)(unsafe.Pointer(node)).Fdynamic_precedence = (*StackNode)(unsafe.Pointer(previous_node)).Fdynamic_precedence
		(*StackNode)(unsafe.Pointer(node)).Fnode_count = (*StackNode)(unsafe.Pointer(previous_node)).Fnode_count
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			*(*uint32)(unsafe.Pointer(node + 216)) += ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp)))
			(*StackNode)(unsafe.Pointer(node)).Fposition = length_add(tls, (*StackNode)(unsafe.Pointer(node)).Fposition, ts_subtree_total_size(tls, *(*Subtree)(unsafe.Pointer(bp))))
			*(*uint32)(unsafe.Pointer(node + 220)) += stack__subtree_node_count(tls, *(*Subtree)(unsafe.Pointer(bp)))
			*(*int32)(unsafe.Pointer(node + 224)) += ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp)))
		}
	} else {
		(*StackNode)(unsafe.Pointer(node)).Fposition = length_zero(tls)
		(*StackNode)(unsafe.Pointer(node)).Ferror_cost = uint32(0)
	}
	return node
}

func stack__subtree_is_equivalent(tls *libc.TLS, _left Subtree, _right Subtree) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _left
	*(*Subtree)(unsafe.Pointer(bp + 8)) = _right
	if *(*uintptr)(unsafe.Pointer(bp)) == *(*uintptr)(unsafe.Pointer(bp + 8)) {
		return libc.BoolUint8(1 != 0)
	}
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) || !(*(*uintptr)(unsafe.Pointer(bp + 8)) != 0) {
		return libc.BoolUint8(0 != 0)
	}
	if libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))) != libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) {
		return libc.BoolUint8(0 != 0)
	}
	if ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) && ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) > uint32(0) {
		return libc.BoolUint8(1 != 0)
	}
	return libc.BoolUint8(ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(bp))).Fbytes == ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(bp + 8))).Fbytes && ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(bp))).Fbytes == ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(bp + 8))).Fbytes && ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) == ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) && libc.Int32FromUint8(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp)))) == libc.Int32FromUint8(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) && ts_subtree_external_scanner_state_eq(tls, *(*Subtree)(unsafe.Pointer(bp)), *(*Subtree)(unsafe.Pointer(bp + 8))) != 0)
}

func stack_node_add_link(tls *libc.TLS, self uintptr, _link StackLink, subtree_pool uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*StackLink)(unsafe.Pointer(bp)) = _link
	var dynamic_precedence int32_t
	var dynamic_precedence1, i, j int32
	var existing_link, v4 uintptr
	var node_count uint32
	var v3 uint16
	_, _, _, _, _, _, _, _ = dynamic_precedence, dynamic_precedence1, existing_link, i, j, node_count, v3, v4
	if (*(*StackLink)(unsafe.Pointer(bp))).Fnode == self {
		return
	}
	i = 0
	for {
		if !(i < libc.Int32FromUint16((*StackNode)(unsafe.Pointer(self)).Flink_count)) {
			break
		}
		existing_link = self + 16 + uintptr(i)*24
		if stack__subtree_is_equivalent(tls, (*StackLink)(unsafe.Pointer(existing_link)).Fsubtree, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree) != 0 {
			if (*StackLink)(unsafe.Pointer(existing_link)).Fnode == (*(*StackLink)(unsafe.Pointer(bp))).Fnode {
				if ts_subtree_dynamic_precedence(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree) > ts_subtree_dynamic_precedence(tls, (*StackLink)(unsafe.Pointer(existing_link)).Fsubtree) {
					ts_subtree_retain(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree)
					ts_subtree_release(tls, subtree_pool, (*StackLink)(unsafe.Pointer(existing_link)).Fsubtree)
					(*StackLink)(unsafe.Pointer(existing_link)).Fsubtree = (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree
					(*StackNode)(unsafe.Pointer(self)).Fdynamic_precedence = (*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Fdynamic_precedence + ts_subtree_dynamic_precedence(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree)
				}
				return
			}
			if libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*StackLink)(unsafe.Pointer(existing_link)).Fnode)).Fstate) == libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Fstate) && (*StackNode)(unsafe.Pointer((*StackLink)(unsafe.Pointer(existing_link)).Fnode)).Fposition.Fbytes == (*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Fposition.Fbytes && (*StackNode)(unsafe.Pointer((*StackLink)(unsafe.Pointer(existing_link)).Fnode)).Ferror_cost == (*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Ferror_cost {
				j = 0
				for {
					if !(j < libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Flink_count)) {
						break
					}
					stack_node_add_link(tls, (*StackLink)(unsafe.Pointer(existing_link)).Fnode, *(*StackLink)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode + 16 + uintptr(j)*24)), subtree_pool)
					goto _2
				_2:
					;
					j = j + 1
				}
				dynamic_precedence = (*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Fdynamic_precedence
				if *(*uintptr)(unsafe.Pointer(bp + 8)) != 0 {
					dynamic_precedence = dynamic_precedence + ts_subtree_dynamic_precedence(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree)
				}
				if dynamic_precedence > (*StackNode)(unsafe.Pointer(self)).Fdynamic_precedence {
					(*StackNode)(unsafe.Pointer(self)).Fdynamic_precedence = dynamic_precedence
				}
				return
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if libc.Int32FromUint16((*StackNode)(unsafe.Pointer(self)).Flink_count) == int32(8) {
		return
	}
	stack_node_retain(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fnode)
	node_count = (*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Fnode_count
	dynamic_precedence1 = (*StackNode)(unsafe.Pointer((*(*StackLink)(unsafe.Pointer(bp))).Fnode)).Fdynamic_precedence
	v4 = self + 208
	v3 = *(*uint16)(unsafe.Pointer(v4))
	*(*uint16)(unsafe.Pointer(v4)) = *(*uint16)(unsafe.Pointer(v4)) + 1
	*(*StackLink)(unsafe.Pointer(self + 16 + uintptr(v3)*24)) = *(*StackLink)(unsafe.Pointer(bp))
	if *(*uintptr)(unsafe.Pointer(bp + 8)) != 0 {
		ts_subtree_retain(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree)
		node_count = node_count + stack__subtree_node_count(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree)
		dynamic_precedence1 = dynamic_precedence1 + ts_subtree_dynamic_precedence(tls, (*(*StackLink)(unsafe.Pointer(bp))).Fsubtree)
	}
	if node_count > (*StackNode)(unsafe.Pointer(self)).Fnode_count {
		(*StackNode)(unsafe.Pointer(self)).Fnode_count = node_count
	}
	if dynamic_precedence1 > (*StackNode)(unsafe.Pointer(self)).Fdynamic_precedence {
		(*StackNode)(unsafe.Pointer(self)).Fdynamic_precedence = dynamic_precedence1
	}
}

func stack_head_delete(tls *libc.TLS, self uintptr, pool uintptr, subtree_pool uintptr) {
	if (*StackHead)(unsafe.Pointer(self)).Fnode != 0 {
		if *(*uintptr)(unsafe.Pointer(self + 24)) != 0 {
			ts_subtree_release(tls, subtree_pool, (*StackHead)(unsafe.Pointer(self)).Flast_external_token)
		}
		if *(*uintptr)(unsafe.Pointer(self + 32)) != 0 {
			ts_subtree_release(tls, subtree_pool, (*StackHead)(unsafe.Pointer(self)).Flookahead_when_paused)
		}
		if (*StackHead)(unsafe.Pointer(self)).Fsummary != 0 {
			if (*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(self)).Fsummary)).Fcontents != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(self)).Fsummary)).Fcontents)
			}
			(*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(self)).Fsummary)).Fcontents = libc.UintptrFromInt32(0)
			(*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(self)).Fsummary)).Fsize = uint32(0)
			(*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(self)).Fsummary)).Fcapacity = uint32(0)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*StackHead)(unsafe.Pointer(self)).Fsummary)
		}
		stack_node_release(tls, (*StackHead)(unsafe.Pointer(self)).Fnode, pool, subtree_pool)
	}
}

func ts_stack__add_version(tls *libc.TLS, self uintptr, original_version StackVersion, node uintptr) (r StackVersion) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var v1 uint32_t
	var v2 uintptr
	var _ /* head at bp+0 */ StackHead
	_, _ = v1, v2
	_ = libc.Uint64FromInt64(4)
	{
		if !(original_version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8850, __ccgo_ts+8787, int32(293), uintptr(unsafe.Pointer(&__func__71)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(original_version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8850, __ccgo_ts+8787, int32(294), uintptr(unsafe.Pointer(&__func__71)))
		}
	}
	*(*StackHead)(unsafe.Pointer(bp)) = StackHead{}
	*(*uintptr)(unsafe.Pointer(bp)) = node
	*(*uint32)(unsafe.Pointer(bp + 16)) = (*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(original_version)*48)).Fnode_count_at_last_error
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp + 24)) = (*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(original_version)*48)).Flast_external_token
	*(*StackStatus)(unsafe.Pointer(bp + 40)) = int32(StackStatusActive)
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp + 32)) = Subtree{}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(48))
	v2 = self + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*48)) = *(*StackHead)(unsafe.Pointer(bp))
	stack_node_retain(tls, node)
	if *(*uintptr)(unsafe.Pointer(bp + 24)) != 0 {
		ts_subtree_retain(tls, (*(*StackHead)(unsafe.Pointer(bp))).Flast_external_token)
	}
	return (*Stack)(unsafe.Pointer(self)).Fheads.Fsize - libc.Uint32FromInt32(1)
}

var __func__71 = [22]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', '_', 'a', 'd', 'd', '_', 'v', 'e', 'r', 's', 'i', 'o', 'n'}

func ts_stack__add_slice(tls *libc.TLS, self uintptr, original_version StackVersion, node uintptr, subtrees uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, v2 uint32_t
	var slice1 StackSlice
	var version, version1 StackVersion
	var v3 uintptr
	var _ /* slice at bp+0 */ StackSlice
	_, _, _, _, _, _ = i, slice1, version, version1, v2, v3
	i = (*Stack)(unsafe.Pointer(self)).Fslices.Fsize - uint32(1)
	for {
		if !(i+uint32(1) > uint32(0)) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*StackSliceArray)(unsafe.Pointer(self+16)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+3237, __ccgo_ts+8787, int32(311), uintptr(unsafe.Pointer(&__func__72)))
			}
		}
		version = (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(self+16)).Fcontents + uintptr(i)*24)).Fversion
		_ = libc.Uint64FromInt64(4)
		{
			if !(version < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(312), uintptr(unsafe.Pointer(&__func__72)))
			}
		}
		if (*StackHead)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents+uintptr(version)*48)).Fnode == node {
			*(*StackSlice)(unsafe.Pointer(bp)) = StackSlice{
				Fsubtrees: *(*SubtreeArray)(unsafe.Pointer(subtrees)),
				Fversion:  version,
			}
			(*StackSliceArray)(unsafe.Pointer(self + 16)).Fcontents = _array__splice(tls, (*StackSliceArray)(unsafe.Pointer(self+16)).Fcontents, self+16+8, self+16+12, libc.Uint64FromInt64(24), i+uint32(1), uint32(0), uint32(1), bp)
			return
		}
		goto _1
	_1:
		;
		i = i - 1
	}
	version1 = ts_stack__add_version(tls, self, original_version, node)
	slice1 = StackSlice{
		Fsubtrees: *(*SubtreeArray)(unsafe.Pointer(subtrees)),
		Fversion:  version1,
	}
	(*StackSliceArray)(unsafe.Pointer(self + 16)).Fcontents = _array__grow(tls, (*StackSliceArray)(unsafe.Pointer(self+16)).Fcontents, (*StackSliceArray)(unsafe.Pointer(self+16)).Fsize, self+16+12, uint32(1), libc.Uint64FromInt64(24))
	v3 = self + 16 + 8
	v2 = *(*uint32_t)(unsafe.Pointer(v3))
	*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
	*(*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(self+16)).Fcontents + uintptr(v2)*24)) = slice1
}

var __func__72 = [20]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', '_', 'a', 'd', 'd', '_', 's', 'l', 'i', 'c', 'e'}

func stack__iter(tls *libc.TLS, self uintptr, version StackVersion, __ccgo_fp_callback StackCallback, payload uintptr, goal_subtree_count int32) (r StackSliceArray) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var action StackAction
	var current_iterator StackIterator
	var head, iterator, next_iterator, node, v2 uintptr
	var i, j, size, v1 uint32_t
	var include_subtrees, should_pop, should_stop uint8
	var _ /* link at bp+48 */ StackLink
	var _ /* new_iterator at bp+0 */ StackIterator
	var _ /* subtrees at bp+32 */ SubtreeArray
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = action, current_iterator, head, i, include_subtrees, iterator, j, next_iterator, node, should_pop, should_stop, size, v1, v2
	(*StackSliceArray)(unsafe.Pointer(self + 16)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fsize = uint32(0)
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(334), uintptr(unsafe.Pointer(&__func__73)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	*(*StackIterator)(unsafe.Pointer(bp)) = StackIterator{
		Fnode:       (*StackHead)(unsafe.Pointer(head)).Fnode,
		Fis_pending: libc.BoolUint8(1 != 0),
	}
	include_subtrees = libc.BoolUint8(0 != 0)
	if goal_subtree_count >= 0 {
		include_subtrees = libc.BoolUint8(1 != 0)
		(*SubtreeArray)(unsafe.Pointer(bp + 8)).Fcontents = _array__reserve(tls, (*SubtreeArray)(unsafe.Pointer(bp+8)).Fcontents, bp+8+12, libc.Uint64FromInt64(8), uint32(uint64(uint32(ts_subtree_alloc_size(tls, libc.Uint32FromInt32(goal_subtree_count))))/uint64(8)))
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+32)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+32)).Fsize, self+32+12, uint32(1), libc.Uint64FromInt64(32))
	v2 = self + 32 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*StackIterator)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+32)).Fcontents + uintptr(v1)*32)) = *(*StackIterator)(unsafe.Pointer(bp))
	for (*Stack)(unsafe.Pointer(self)).Fiterators.Fsize > uint32(0) {
		i = uint32(0)
		size = (*Stack)(unsafe.Pointer(self)).Fiterators.Fsize
		for {
			if !(i < size) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+32)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+8945, __ccgo_ts+8787, int32(352), uintptr(unsafe.Pointer(&__func__73)))
				}
			}
			iterator = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+32)).Fcontents + uintptr(i)*32
			node = (*StackIterator)(unsafe.Pointer(iterator)).Fnode
			action = (*(*func(*libc.TLS, uintptr, uintptr) StackAction)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_callback})))(tls, payload, iterator)
			should_pop = uint8(libc.BoolUint32(action&uint32(StackActionPop) != 0))
			should_stop = libc.BoolUint8(action&uint32(StackActionStop) != 0 || libc.Int32FromUint16((*StackNode)(unsafe.Pointer(node)).Flink_count) == 0)
			if should_pop != 0 {
				*(*SubtreeArray)(unsafe.Pointer(bp + 32)) = (*StackIterator)(unsafe.Pointer(iterator)).Fsubtrees
				if !(should_stop != 0) {
					ts_subtree_array_copy(tls, *(*SubtreeArray)(unsafe.Pointer(bp + 32)), bp+32)
				}
				ts_subtree_array_reverse(tls, bp+32)
				ts_stack__add_slice(tls, self, version, node, bp+32)
			}
			if should_stop != 0 {
				if !(should_pop != 0) {
					ts_subtree_array_delete(tls, (*Stack)(unsafe.Pointer(self)).Fsubtree_pool, iterator+8)
				}
				_array__erase(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+32)).Fcontents, self+32+8, libc.Uint64FromInt64(32), i)
				i = i - 1
				size = size - 1
				goto _3
			}
			j = uint32(1)
			for {
				if !(j <= uint32((*StackNode)(unsafe.Pointer(node)).Flink_count)) {
					break
				}
				if j == uint32((*StackNode)(unsafe.Pointer(node)).Flink_count) {
					*(*StackLink)(unsafe.Pointer(bp + 48)) = *(*StackLink)(unsafe.Pointer(node + 16))
					_ = libc.Uint64FromInt64(4)
					{
						if !(i < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+32)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8945, __ccgo_ts+8787, int32(387), uintptr(unsafe.Pointer(&__func__73)))
						}
					}
					next_iterator = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents + uintptr(i)*32
				} else {
					if (*Stack)(unsafe.Pointer(self)).Fiterators.Fsize >= uint32(64) {
						goto _4
					}
					*(*StackLink)(unsafe.Pointer(bp + 48)) = *(*StackLink)(unsafe.Pointer(node + 16 + uintptr(j)*24))
					_ = libc.Uint64FromInt64(4)
					{
						if !(i < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+32)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8945, __ccgo_ts+8787, int32(391), uintptr(unsafe.Pointer(&__func__73)))
						}
					}
					current_iterator = *(*StackIterator)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents + uintptr(i)*32))
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 32)).Fcontents = _array__grow(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fsize, self+32+12, uint32(1), libc.Uint64FromInt64(32))
					v2 = self + 32 + 8
					v1 = *(*uint32_t)(unsafe.Pointer(v2))
					*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
					*(*StackIterator)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents + uintptr(v1)*32)) = current_iterator
					_ = libc.Uint64FromInt64(4)
					{
						if !((*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+32)).Fsize-libc.Uint32FromInt32(1) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+32)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8986, __ccgo_ts+8787, int32(393), uintptr(unsafe.Pointer(&__func__73)))
						}
					}
					next_iterator = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents + uintptr((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fsize-uint32(1))*32
					ts_subtree_array_copy(tls, (*StackIterator)(unsafe.Pointer(next_iterator)).Fsubtrees, next_iterator+8)
				}
				(*StackIterator)(unsafe.Pointer(next_iterator)).Fnode = (*(*StackLink)(unsafe.Pointer(bp + 48))).Fnode
				if *(*uintptr)(unsafe.Pointer(bp + 48 + 8)) != 0 {
					if include_subtrees != 0 {
						(*SubtreeArray)(unsafe.Pointer(next_iterator + 8)).Fcontents = _array__grow(tls, (*SubtreeArray)(unsafe.Pointer(next_iterator+8)).Fcontents, (*SubtreeArray)(unsafe.Pointer(next_iterator+8)).Fsize, next_iterator+8+12, uint32(1), libc.Uint64FromInt64(8))
						v2 = next_iterator + 8 + 8
						v1 = *(*uint32_t)(unsafe.Pointer(v2))
						*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
						*(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(next_iterator+8)).Fcontents + uintptr(v1)*8)) = (*(*StackLink)(unsafe.Pointer(bp + 48))).Fsubtree
						ts_subtree_retain(tls, (*(*StackLink)(unsafe.Pointer(bp + 48))).Fsubtree)
					}
					if !(ts_subtree_extra(tls, (*(*StackLink)(unsafe.Pointer(bp + 48))).Fsubtree) != 0) {
						(*StackIterator)(unsafe.Pointer(next_iterator)).Fsubtree_count = (*StackIterator)(unsafe.Pointer(next_iterator)).Fsubtree_count + 1
						if !((*(*StackLink)(unsafe.Pointer(bp + 48))).Fis_pending != 0) {
							(*StackIterator)(unsafe.Pointer(next_iterator)).Fis_pending = libc.BoolUint8(0 != 0)
						}
					}
				} else {
					(*StackIterator)(unsafe.Pointer(next_iterator)).Fsubtree_count = (*StackIterator)(unsafe.Pointer(next_iterator)).Fsubtree_count + 1
					(*StackIterator)(unsafe.Pointer(next_iterator)).Fis_pending = libc.BoolUint8(0 != 0)
				}
				goto _4
			_4:
				;
				j = j + 1
			}
			goto _3
		_3:
			;
			i = i + 1
		}
	}
	return (*Stack)(unsafe.Pointer(self)).Fslices
}

var __func__73 = [12]int8{'s', 't', 'a', 'c', 'k', '_', '_', 'i', 't', 'e', 'r'}

func ts_stack_new(tls *libc.TLS, subtree_pool uintptr) (r uintptr) {
	var self uintptr
	_ = self
	self = (*(*func(*libc.TLS, size_t, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_calloc})))(tls, uint64(1), uint64(80))
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcapacity = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*StackSliceArray)(unsafe.Pointer(self + 16)).Fsize = uint32(0)
	(*StackSliceArray)(unsafe.Pointer(self + 16)).Fcapacity = uint32(0)
	(*StackSliceArray)(unsafe.Pointer(self + 16)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fcapacity = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fcontents = libc.UintptrFromInt32(0)
	(*StackNodeArray)(unsafe.Pointer(self + 48)).Fsize = uint32(0)
	(*StackNodeArray)(unsafe.Pointer(self + 48)).Fcapacity = uint32(0)
	(*StackNodeArray)(unsafe.Pointer(self + 48)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__reserve(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, self+12, libc.Uint64FromInt64(48), uint32(4))
	(*StackSliceArray)(unsafe.Pointer(self + 16)).Fcontents = _array__reserve(tls, (*StackSliceArray)(unsafe.Pointer(self+16)).Fcontents, self+16+12, libc.Uint64FromInt64(24), uint32(4))
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fcontents = _array__reserve(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+32)).Fcontents, self+32+12, libc.Uint64FromInt64(32), uint32(4))
	(*StackNodeArray)(unsafe.Pointer(self + 48)).Fcontents = _array__reserve(tls, (*StackNodeArray)(unsafe.Pointer(self+48)).Fcontents, self+48+12, libc.Uint64FromInt64(8), uint32(50))
	(*Stack)(unsafe.Pointer(self)).Fsubtree_pool = subtree_pool
	(*Stack)(unsafe.Pointer(self)).Fbase_node = stack_node_new(tls, libc.UintptrFromInt32(0), Subtree{}, libc.BoolUint8(0 != 0), uint16(1), self+48)
	ts_stack_clear(tls, self)
	return self
}

func ts_stack_delete(tls *libc.TLS, self uintptr) {
	var i, i1 uint32_t
	_, _ = i, i1
	if (*Stack)(unsafe.Pointer(self)).Fslices.Fcontents != 0 {
		if (*StackSliceArray)(unsafe.Pointer(self+16)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*StackSliceArray)(unsafe.Pointer(self+16)).Fcontents)
		}
		(*StackSliceArray)(unsafe.Pointer(self + 16)).Fcontents = libc.UintptrFromInt32(0)
		(*StackSliceArray)(unsafe.Pointer(self + 16)).Fsize = uint32(0)
		(*StackSliceArray)(unsafe.Pointer(self + 16)).Fcapacity = uint32(0)
	}
	if (*Stack)(unsafe.Pointer(self)).Fiterators.Fcontents != 0 {
		if (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+32)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+32)).Fcontents)
		}
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 32)).Fcontents = libc.UintptrFromInt32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 32)).Fsize = uint32(0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 32)).Fcapacity = uint32(0)
	}
	stack_node_release(tls, (*Stack)(unsafe.Pointer(self)).Fbase_node, self+48, (*Stack)(unsafe.Pointer(self)).Fsubtree_pool)
	i = uint32(0)
	for {
		if !(i < (*Stack)(unsafe.Pointer(self)).Fheads.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+9054, __ccgo_ts+8787, int32(447), uintptr(unsafe.Pointer(&__func__74)))
			}
		}
		stack_head_delete(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents+uintptr(i)*48, self+48, (*Stack)(unsafe.Pointer(self)).Fsubtree_pool)
		goto _1
	_1:
		;
		i = i + 1
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	if (*Stack)(unsafe.Pointer(self)).Fnode_pool.Fcontents != 0 {
		i1 = uint32(0)
		for {
			if !(i1 < (*Stack)(unsafe.Pointer(self)).Fnode_pool.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i1 < (*StackNodeArray)(unsafe.Pointer(self+48)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+9091, __ccgo_ts+8787, int32(452), uintptr(unsafe.Pointer(&__func__74)))
				}
			}
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, *(*uintptr)(unsafe.Pointer((*StackNodeArray)(unsafe.Pointer(self+48)).Fcontents + uintptr(i1)*8)))
			goto _2
		_2:
			;
			i1 = i1 + 1
		}
		if (*StackNodeArray)(unsafe.Pointer(self+48)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*StackNodeArray)(unsafe.Pointer(self+48)).Fcontents)
		}
		(*StackNodeArray)(unsafe.Pointer(self + 48)).Fcontents = libc.UintptrFromInt32(0)
		(*StackNodeArray)(unsafe.Pointer(self + 48)).Fsize = uint32(0)
		(*StackNodeArray)(unsafe.Pointer(self + 48)).Fcapacity = uint32(0)
	}
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcapacity = uint32(0)
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, self)
}

var __func__74 = [16]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'd', 'e', 'l', 'e', 't', 'e'}

func ts_stack_version_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	return (*Stack)(unsafe.Pointer(self)).Fheads.Fsize
}

func ts_stack_halted_version_count(tls *libc.TLS, self uintptr) (r uint32_t) {
	var count, i uint32_t
	var head uintptr
	_, _, _ = count, head, i
	count = uint32(0)
	i = uint32(0)
	for {
		if !(i < (*Stack)(unsafe.Pointer(self)).Fheads.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+9054, __ccgo_ts+8787, int32(466), uintptr(unsafe.Pointer(&__func__75)))
			}
		}
		head = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(i)*48
		if (*StackHead)(unsafe.Pointer(head)).Fstatus == int32(StackStatusHalted) {
			count = count + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return count
}

var __func__75 = [30]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'h', 'a', 'l', 't', 'e', 'd', '_', 'v', 'e', 'r', 's', 'i', 'o', 'n', '_', 'c', 'o', 'u', 'n', 't'}

func ts_stack_state(tls *libc.TLS, self uintptr, version StackVersion) (r TSStateId) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(475), uintptr(unsafe.Pointer(&__func__76)))
		}
	}
	return (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48)).Fnode)).Fstate
}

var __func__76 = [15]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 's', 't', 'a', 't', 'e'}

func ts_stack_position(tls *libc.TLS, self uintptr, version StackVersion) (r Length) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(479), uintptr(unsafe.Pointer(&__func__77)))
		}
	}
	return (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48)).Fnode)).Fposition
}

var __func__77 = [18]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'p', 'o', 's', 'i', 't', 'i', 'o', 'n'}

func ts_stack_last_external_token(tls *libc.TLS, self uintptr, version StackVersion) (r Subtree) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(483), uintptr(unsafe.Pointer(&__func__78)))
		}
	}
	return (*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48)).Flast_external_token
}

var __func__78 = [29]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'l', 'a', 's', 't', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 't', 'o', 'k', 'e', 'n'}

func ts_stack_set_last_external_token(tls *libc.TLS, self uintptr, version StackVersion, _token Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _token
	var head uintptr
	_ = head
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(487), uintptr(unsafe.Pointer(&__func__79)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer(bp)))
	}
	if *(*uintptr)(unsafe.Pointer(head + 24)) != 0 {
		ts_subtree_release(tls, (*Stack)(unsafe.Pointer(self)).Fsubtree_pool, (*StackHead)(unsafe.Pointer(head)).Flast_external_token)
	}
	(*StackHead)(unsafe.Pointer(head)).Flast_external_token = *(*Subtree)(unsafe.Pointer(bp))
}

var __func__79 = [33]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 's', 'e', 't', '_', 'l', 'a', 's', 't', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 't', 'o', 'k', 'e', 'n'}

func ts_stack_error_cost(tls *libc.TLS, self uintptr, version StackVersion) (r uint32) {
	var head uintptr
	var result uint32
	_, _ = head, result
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(494), uintptr(unsafe.Pointer(&__func__80)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	result = (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fnode)).Ferror_cost
	if (*StackHead)(unsafe.Pointer(head)).Fstatus == int32(StackStatusPaused) || libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fnode)).Fstate) == 0 && !(*(*uintptr)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fnode + 16 + 8)) != 0) {
		result = result + uint32(500)
	}
	return result
}

var __func__80 = [20]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'e', 'r', 'r', 'o', 'r', '_', 'c', 'o', 's', 't'}

func ts_stack_node_count_since_error(tls *libc.TLS, self uintptr, version StackVersion) (r uint32) {
	var head uintptr
	_ = head
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(505), uintptr(unsafe.Pointer(&__func__81)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	if (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fnode)).Fnode_count < (*StackHead)(unsafe.Pointer(head)).Fnode_count_at_last_error {
		(*StackHead)(unsafe.Pointer(head)).Fnode_count_at_last_error = (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fnode)).Fnode_count
	}
	return (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fnode)).Fnode_count - (*StackHead)(unsafe.Pointer(head)).Fnode_count_at_last_error
}

var __func__81 = [32]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'n', 'o', 'd', 'e', '_', 'c', 'o', 'u', 'n', 't', '_', 's', 'i', 'n', 'c', 'e', '_', 'e', 'r', 'r', 'o', 'r'}

func ts_stack_push(tls *libc.TLS, self uintptr, version StackVersion, _subtree Subtree, pending uint8, state TSStateId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _subtree
	var head, new_node uintptr
	_, _ = head, new_node
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(519), uintptr(unsafe.Pointer(&__func__82)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	new_node = stack_node_new(tls, (*StackHead)(unsafe.Pointer(head)).Fnode, *(*Subtree)(unsafe.Pointer(bp)), pending, state, self+48)
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) {
		(*StackHead)(unsafe.Pointer(head)).Fnode_count_at_last_error = (*StackNode)(unsafe.Pointer(new_node)).Fnode_count
	}
	(*StackHead)(unsafe.Pointer(head)).Fnode = new_node
}

var __func__82 = [14]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'p', 'u', 's', 'h'}

func pop_count_callback(tls *libc.TLS, payload uintptr, iterator uintptr) (r StackAction) {
	var goal_subtree_count uintptr
	_ = goal_subtree_count
	goal_subtree_count = payload
	if (*StackIterator)(unsafe.Pointer(iterator)).Fsubtree_count == *(*uint32)(unsafe.Pointer(goal_subtree_count)) {
		return libc.Uint32FromInt32(int32(StackActionPop) | int32(StackActionStop))
	} else {
		return uint32(StackActionNone)
	}
	return r
}

func ts_stack_pop_count(tls *libc.TLS, self uintptr, version StackVersion, _count uint32_t) (r StackSliceArray) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*uint32_t)(unsafe.Pointer(bp)) = _count
	return stack__iter(tls, self, version, __ccgo_fp(pop_count_callback), bp, libc.Int32FromUint32(*(*uint32_t)(unsafe.Pointer(bp))))
}

func pop_pending_callback(tls *libc.TLS, payload uintptr, iterator uintptr) (r StackAction) {
	_ = payload
	if (*StackIterator)(unsafe.Pointer(iterator)).Fsubtree_count >= uint32(1) {
		if (*StackIterator)(unsafe.Pointer(iterator)).Fis_pending != 0 {
			return libc.Uint32FromInt32(int32(StackActionPop) | int32(StackActionStop))
		} else {
			return uint32(StackActionStop)
		}
	} else {
		return uint32(StackActionNone)
	}
	return r
}

func ts_stack_pop_pending(tls *libc.TLS, self uintptr, version StackVersion) (r StackSliceArray) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* pop at bp+0 */ StackSliceArray
	*(*StackSliceArray)(unsafe.Pointer(bp)) = stack__iter(tls, self, version, __ccgo_fp(pop_pending_callback), libc.UintptrFromInt32(0), 0)
	if (*(*StackSliceArray)(unsafe.Pointer(bp))).Fsize > uint32(0) {
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+8787, int32(555), uintptr(unsafe.Pointer(&__func__83)))
			}
		}
		ts_stack_renumber_version(tls, self, (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents)).Fversion, version)
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+8787, int32(556), uintptr(unsafe.Pointer(&__func__83)))
			}
		}
		(*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp)).Fcontents)).Fversion = version
	}
	return *(*StackSliceArray)(unsafe.Pointer(bp))
}

var __func__83 = [21]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'p', 'o', 'p', '_', 'p', 'e', 'n', 'd', 'i', 'n', 'g'}

func pop_error_callback(tls *libc.TLS, payload uintptr, iterator uintptr) (r StackAction) {
	var found_error uintptr
	var v1 bool
	_, _ = found_error, v1
	if (*StackIterator)(unsafe.Pointer(iterator)).Fsubtrees.Fsize > uint32(0) {
		found_error = payload
		if v1 = !(*(*uint8)(unsafe.Pointer(found_error)) != 0); v1 {
			_ = libc.Uint64FromInt64(4)
			{
				if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*SubtreeArray)(unsafe.Pointer(iterator+8)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+9132, __ccgo_ts+8787, int32(564), uintptr(unsafe.Pointer(&__func__84)))
				}
			}
		}
		if v1 && ts_subtree_is_error(tls, *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(iterator + 8)).Fcontents))) != 0 {
			*(*uint8)(unsafe.Pointer(found_error)) = libc.BoolUint8(1 != 0)
			return libc.Uint32FromInt32(int32(StackActionPop) | int32(StackActionStop))
		} else {
			return uint32(StackActionStop)
		}
	} else {
		return uint32(StackActionNone)
	}
	return r
}

var __func__84 = [19]int8{'p', 'o', 'p', '_', 'e', 'r', 'r', 'o', 'r', '_', 'c', 'a', 'l', 'l', 'b', 'a', 'c', 'k'}

func ts_stack_pop_error(tls *libc.TLS, self uintptr, version StackVersion) (r SubtreeArray) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i uint32
	var node uintptr
	var _ /* found_error at bp+0 */ uint8
	var _ /* pop at bp+8 */ StackSliceArray
	_, _ = i, node
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(576), uintptr(unsafe.Pointer(&__func__85)))
		}
	}
	node = (*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48)).Fnode
	i = uint32(0)
	for {
		if !(i < uint32((*StackNode)(unsafe.Pointer(node)).Flink_count)) {
			break
		}
		if *(*uintptr)(unsafe.Pointer(node + 16 + uintptr(i)*24 + 8)) != 0 && ts_subtree_is_error(tls, (*(*StackLink)(unsafe.Pointer(node + 16 + uintptr(i)*24))).Fsubtree) != 0 {
			*(*uint8)(unsafe.Pointer(bp)) = libc.BoolUint8(0 != 0)
			*(*StackSliceArray)(unsafe.Pointer(bp + 8)) = stack__iter(tls, self, version, __ccgo_fp(pop_error_callback), bp, int32(1))
			if (*(*StackSliceArray)(unsafe.Pointer(bp + 8))).Fsize > uint32(0) {
				_ = libc.Uint64FromInt64(4)
				{
					if !((*(*StackSliceArray)(unsafe.Pointer(bp + 8))).Fsize == libc.Uint32FromInt32(1)) {
						libc.X__assert_fail(tls, __ccgo_ts+9176, __ccgo_ts+8787, int32(582), uintptr(unsafe.Pointer(&__func__85)))
					}
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+8)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+8787, int32(583), uintptr(unsafe.Pointer(&__func__85)))
					}
				}
				ts_stack_renumber_version(tls, self, (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp+8)).Fcontents)).Fversion, version)
				_ = libc.Uint64FromInt64(4)
				{
					if !(libc.Uint32FromInt32(libc.Int32FromInt32(0)) < (*StackSliceArray)(unsafe.Pointer(bp+8)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+2152, __ccgo_ts+8787, int32(584), uintptr(unsafe.Pointer(&__func__85)))
					}
				}
				return (*StackSlice)(unsafe.Pointer((*StackSliceArray)(unsafe.Pointer(bp + 8)).Fcontents)).Fsubtrees
			}
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return SubtreeArray{}
}

var __func__85 = [19]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'p', 'o', 'p', '_', 'e', 'r', 'r', 'o', 'r'}

func pop_all_callback(tls *libc.TLS, payload uintptr, iterator uintptr) (r StackAction) {
	var v1 int32
	_ = v1
	_ = payload
	if libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*StackIterator)(unsafe.Pointer(iterator)).Fnode)).Flink_count) == 0 {
		v1 = int32(StackActionPop)
	} else {
		v1 = int32(StackActionNone)
	}
	return libc.Uint32FromInt32(v1)
}

func ts_stack_pop_all(tls *libc.TLS, self uintptr, version StackVersion) (r StackSliceArray) {
	return stack__iter(tls, self, version, __ccgo_fp(pop_all_callback), libc.UintptrFromInt32(0), 0)
}

type SummarizeStackSession = struct {
	Fsummary   uintptr
	Fmax_depth uint32
}

func summarize_stack_callback(tls *libc.TLS, payload uintptr, iterator uintptr) (r StackAction) {
	var depth, i uint32
	var entry StackSummaryEntry
	var session, v3 uintptr
	var state TSStateId
	var v2 uint32_t
	_, _, _, _, _, _, _ = depth, entry, i, session, state, v2, v3
	session = payload
	state = (*StackNode)(unsafe.Pointer((*StackIterator)(unsafe.Pointer(iterator)).Fnode)).Fstate
	depth = (*StackIterator)(unsafe.Pointer(iterator)).Fsubtree_count
	if depth > (*SummarizeStackSession)(unsafe.Pointer(session)).Fmax_depth {
		return uint32(StackActionStop)
	}
	i = (*StackSummary)(unsafe.Pointer((*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary)).Fsize - uint32(1)
	for {
		if !(i+uint32(1) > uint32(0)) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*StackSummary)(unsafe.Pointer((*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+9190, __ccgo_ts+8787, int32(612), uintptr(unsafe.Pointer(&__func__86)))
			}
		}
		entry = *(*StackSummaryEntry)(unsafe.Pointer((*StackSummary)(unsafe.Pointer((*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary)).Fcontents + uintptr(i)*20))
		if entry.Fdepth < depth {
			break
		}
		if entry.Fdepth == depth && libc.Int32FromUint16(entry.Fstate) == libc.Int32FromUint16(state) {
			return uint32(StackActionNone)
		}
		goto _1
	_1:
		;
		i = i - 1
	}
	(*StackSummary)(unsafe.Pointer((*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary)).Fcontents = _array__grow(tls, (*StackSummary)(unsafe.Pointer((*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary)).Fcontents, (*StackSummary)(unsafe.Pointer((*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary)).Fsize, (*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary+12, uint32(1), libc.Uint64FromInt64(20))
	v3 = (*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary + 8
	v2 = *(*uint32_t)(unsafe.Pointer(v3))
	*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
	*(*StackSummaryEntry)(unsafe.Pointer((*StackSummary)(unsafe.Pointer((*SummarizeStackSession)(unsafe.Pointer(session)).Fsummary)).Fcontents + uintptr(v2)*20)) = StackSummaryEntry{
		Fposition: (*StackNode)(unsafe.Pointer((*StackIterator)(unsafe.Pointer(iterator)).Fnode)).Fposition,
		Fdepth:    depth,
		Fstate:    state,
	}
	return uint32(StackActionNone)
}

var __func__86 = [25]int8{'s', 'u', 'm', 'm', 'a', 'r', 'i', 'z', 'e', '_', 's', 't', 'a', 'c', 'k', '_', 'c', 'a', 'l', 'l', 'b', 'a', 'c', 'k'}

func ts_stack_record_summary(tls *libc.TLS, self uintptr, version StackVersion, max_depth uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var head uintptr
	var _ /* session at bp+0 */ SummarizeStackSession
	_ = head
	*(*SummarizeStackSession)(unsafe.Pointer(bp)) = SummarizeStackSession{
		Fsummary:   (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(16)),
		Fmax_depth: max_depth,
	}
	(*StackSummary)(unsafe.Pointer((*(*SummarizeStackSession)(unsafe.Pointer(bp))).Fsummary)).Fsize = uint32(0)
	(*StackSummary)(unsafe.Pointer((*(*SummarizeStackSession)(unsafe.Pointer(bp))).Fsummary)).Fcapacity = uint32(0)
	(*StackSummary)(unsafe.Pointer((*(*SummarizeStackSession)(unsafe.Pointer(bp))).Fsummary)).Fcontents = libc.UintptrFromInt32(0)
	stack__iter(tls, self, version, __ccgo_fp(summarize_stack_callback), bp, -int32(1))
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(631), uintptr(unsafe.Pointer(&__func__87)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	if (*StackHead)(unsafe.Pointer(head)).Fsummary != 0 {
		if (*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fcontents)
		}
		(*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fcontents = libc.UintptrFromInt32(0)
		(*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fsize = uint32(0)
		(*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fcapacity = uint32(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*StackHead)(unsafe.Pointer(head)).Fsummary)
	}
	(*StackHead)(unsafe.Pointer(head)).Fsummary = (*(*SummarizeStackSession)(unsafe.Pointer(bp))).Fsummary
}

var __func__87 = [24]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'r', 'e', 'c', 'o', 'r', 'd', '_', 's', 'u', 'm', 'm', 'a', 'r', 'y'}

func ts_stack_get_summary(tls *libc.TLS, self uintptr, version StackVersion) (r uintptr) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(640), uintptr(unsafe.Pointer(&__func__88)))
		}
	}
	return (*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48)).Fsummary
}

var __func__88 = [21]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'g', 'e', 't', '_', 's', 'u', 'm', 'm', 'a', 'r', 'y'}

func ts_stack_dynamic_precedence(tls *libc.TLS, self uintptr, version StackVersion) (r int32) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(644), uintptr(unsafe.Pointer(&__func__89)))
		}
	}
	return (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48)).Fnode)).Fdynamic_precedence
}

var __func__89 = [28]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'd', 'y', 'n', 'a', 'm', 'i', 'c', '_', 'p', 'r', 'e', 'c', 'e', 'd', 'e', 'n', 'c', 'e'}

func ts_stack_has_advanced_since_error(tls *libc.TLS, self uintptr, version StackVersion) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var head, node uintptr
	var _ /* subtree at bp+0 */ Subtree
	_, _ = head, node
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(648), uintptr(unsafe.Pointer(&__func__90)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	node = (*StackHead)(unsafe.Pointer(head)).Fnode
	if (*StackNode)(unsafe.Pointer(node)).Ferror_cost == uint32(0) {
		return libc.BoolUint8(1 != 0)
	}
	for node != 0 {
		if libc.Int32FromUint16((*StackNode)(unsafe.Pointer(node)).Flink_count) > 0 {
			*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
			*(*struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			})(unsafe.Pointer(bp)) = (*(*StackLink)(unsafe.Pointer(node + 16))).Fsubtree
			if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
				if ts_subtree_total_bytes(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
					return libc.BoolUint8(1 != 0)
				} else {
					if (*StackNode)(unsafe.Pointer(node)).Fnode_count > (*StackHead)(unsafe.Pointer(head)).Fnode_count_at_last_error && ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp))) == uint32(0) {
						node = (*(*StackLink)(unsafe.Pointer(node + 16))).Fnode
						continue
					}
				}
			}
		}
		break
	}
	return libc.BoolUint8(0 != 0)
}

var __func__90 = [34]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'h', 'a', 's', '_', 'a', 'd', 'v', 'a', 'n', 'c', 'e', 'd', '_', 's', 'i', 'n', 'c', 'e', '_', 'e', 'r', 'r', 'o', 'r'}

func ts_stack_remove_version(tls *libc.TLS, self uintptr, version StackVersion) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(672), uintptr(unsafe.Pointer(&__func__91)))
		}
	}
	stack_head_delete(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents+uintptr(version)*48, self+48, (*Stack)(unsafe.Pointer(self)).Fsubtree_pool)
	_array__erase(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, self+8, libc.Uint64FromInt64(48), version)
}

var __func__91 = [24]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'r', 'e', 'm', 'o', 'v', 'e', '_', 'v', 'e', 'r', 's', 'i', 'o', 'n'}

func ts_stack_renumber_version(tls *libc.TLS, self uintptr, v1 StackVersion, v2 StackVersion) {
	var source_head, target_head uintptr
	_, _ = source_head, target_head
	if v1 == v2 {
		return
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(v2 < v1) {
			libc.X__assert_fail(tls, __ccgo_ts+9231, __ccgo_ts+8787, int32(678), uintptr(unsafe.Pointer(&__func__92)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(v1 < (*Stack)(unsafe.Pointer(self)).Fheads.Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9239, __ccgo_ts+8787, int32(679), uintptr(unsafe.Pointer(&__func__92)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(v1 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9271, __ccgo_ts+8787, int32(680), uintptr(unsafe.Pointer(&__func__92)))
		}
	}
	source_head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*48
	_ = libc.Uint64FromInt64(4)
	{
		if !(v2 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9309, __ccgo_ts+8787, int32(681), uintptr(unsafe.Pointer(&__func__92)))
		}
	}
	target_head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v2)*48
	if (*StackHead)(unsafe.Pointer(target_head)).Fsummary != 0 && !((*StackHead)(unsafe.Pointer(source_head)).Fsummary != 0) {
		(*StackHead)(unsafe.Pointer(source_head)).Fsummary = (*StackHead)(unsafe.Pointer(target_head)).Fsummary
		(*StackHead)(unsafe.Pointer(target_head)).Fsummary = libc.UintptrFromInt32(0)
	}
	stack_head_delete(tls, target_head, self+48, (*Stack)(unsafe.Pointer(self)).Fsubtree_pool)
	*(*StackHead)(unsafe.Pointer(target_head)) = *(*StackHead)(unsafe.Pointer(source_head))
	_array__erase(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, self+8, libc.Uint64FromInt64(48), v1)
}

var __func__92 = [26]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'r', 'e', 'n', 'u', 'm', 'b', 'e', 'r', '_', 'v', 'e', 'r', 's', 'i', 'o', 'n'}

func ts_stack_swap_versions(tls *libc.TLS, self uintptr, v1 StackVersion, v2 StackVersion) {
	var temporary_head StackHead
	_ = temporary_head
	_ = libc.Uint64FromInt64(4)
	{
		if !(v1 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9271, __ccgo_ts+8787, int32(692), uintptr(unsafe.Pointer(&__func__93)))
		}
	}
	temporary_head = *(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*48))
	_ = libc.Uint64FromInt64(4)
	{
		if !(v1 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9271, __ccgo_ts+8787, int32(693), uintptr(unsafe.Pointer(&__func__93)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(v2 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9309, __ccgo_ts+8787, int32(693), uintptr(unsafe.Pointer(&__func__93)))
		}
	}
	*(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*48)) = *(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v2)*48))
	_ = libc.Uint64FromInt64(4)
	{
		if !(v2 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9309, __ccgo_ts+8787, int32(694), uintptr(unsafe.Pointer(&__func__93)))
		}
	}
	*(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v2)*48)) = temporary_head
}

var __func__93 = [23]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 's', 'w', 'a', 'p', '_', 'v', 'e', 'r', 's', 'i', 'o', 'n', 's'}

func ts_stack_copy_version(tls *libc.TLS, self uintptr, version StackVersion) (r StackVersion) {
	var head, v2 uintptr
	var version_head StackHead
	var v1 uint32_t
	_, _, _, _ = head, version_head, v1, v2
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*Stack)(unsafe.Pointer(self)).Fheads.Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9347, __ccgo_ts+8787, int32(698), uintptr(unsafe.Pointer(&__func__94)))
		}
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(699), uintptr(unsafe.Pointer(&__func__94)))
		}
	}
	version_head = *(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48))
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(48))
	v2 = self + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v1)*48)) = version_head
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9374, __ccgo_ts+8787, int32(701), uintptr(unsafe.Pointer(&__func__94)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize-uint32(1))*48
	stack_node_retain(tls, (*StackHead)(unsafe.Pointer(head)).Fnode)
	if *(*uintptr)(unsafe.Pointer(head + 24)) != 0 {
		ts_subtree_retain(tls, (*StackHead)(unsafe.Pointer(head)).Flast_external_token)
	}
	(*StackHead)(unsafe.Pointer(head)).Fsummary = libc.UintptrFromInt32(0)
	return (*Stack)(unsafe.Pointer(self)).Fheads.Fsize - uint32(1)
}

var __func__94 = [22]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'c', 'o', 'p', 'y', '_', 'v', 'e', 'r', 's', 'i', 'o', 'n'}

func ts_stack_merge(tls *libc.TLS, self uintptr, version1 StackVersion, version2 StackVersion) (r uint8) {
	var head1, head2 uintptr
	var i uint32_t
	_, _, _ = head1, head2, i
	if !(ts_stack_can_merge(tls, self, version1, version2) != 0) {
		return libc.BoolUint8(0 != 0)
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(version1 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9434, __ccgo_ts+8787, int32(710), uintptr(unsafe.Pointer(&__func__95)))
		}
	}
	head1 = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version1)*48
	_ = libc.Uint64FromInt64(4)
	{
		if !(version2 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9478, __ccgo_ts+8787, int32(711), uintptr(unsafe.Pointer(&__func__95)))
		}
	}
	head2 = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version2)*48
	i = uint32(0)
	for {
		if !(i < uint32((*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head2)).Fnode)).Flink_count)) {
			break
		}
		stack_node_add_link(tls, (*StackHead)(unsafe.Pointer(head1)).Fnode, *(*StackLink)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head2)).Fnode + 16 + uintptr(i)*24)), (*Stack)(unsafe.Pointer(self)).Fsubtree_pool)
		goto _1
	_1:
		;
		i = i + 1
	}
	if libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head1)).Fnode)).Fstate) == 0 {
		(*StackHead)(unsafe.Pointer(head1)).Fnode_count_at_last_error = (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head1)).Fnode)).Fnode_count
	}
	ts_stack_remove_version(tls, self, version2)
	return libc.BoolUint8(1 != 0)
}

var __func__95 = [15]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'm', 'e', 'r', 'g', 'e'}

func ts_stack_can_merge(tls *libc.TLS, self uintptr, version1 StackVersion, version2 StackVersion) (r uint8) {
	var head1, head2 uintptr
	_, _ = head1, head2
	_ = libc.Uint64FromInt64(4)
	{
		if !(version1 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9434, __ccgo_ts+8787, int32(723), uintptr(unsafe.Pointer(&__func__96)))
		}
	}
	head1 = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version1)*48
	_ = libc.Uint64FromInt64(4)
	{
		if !(version2 < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+9478, __ccgo_ts+8787, int32(724), uintptr(unsafe.Pointer(&__func__96)))
		}
	}
	head2 = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version2)*48
	return libc.BoolUint8((*StackHead)(unsafe.Pointer(head1)).Fstatus == int32(StackStatusActive) && (*StackHead)(unsafe.Pointer(head2)).Fstatus == int32(StackStatusActive) && libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head1)).Fnode)).Fstate) == libc.Int32FromUint16((*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head2)).Fnode)).Fstate) && (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head1)).Fnode)).Fposition.Fbytes == (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head2)).Fnode)).Fposition.Fbytes && (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head1)).Fnode)).Ferror_cost == (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head2)).Fnode)).Ferror_cost && ts_subtree_external_scanner_state_eq(tls, (*StackHead)(unsafe.Pointer(head1)).Flast_external_token, (*StackHead)(unsafe.Pointer(head2)).Flast_external_token) != 0)
}

var __func__96 = [19]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'c', 'a', 'n', '_', 'm', 'e', 'r', 'g', 'e'}

func ts_stack_halt(tls *libc.TLS, self uintptr, version StackVersion) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(735), uintptr(unsafe.Pointer(&__func__97)))
		}
	}
	(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48)).Fstatus = int32(StackStatusHalted)
}

var __func__97 = [14]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'h', 'a', 'l', 't'}

func ts_stack_pause(tls *libc.TLS, self uintptr, version StackVersion, lookahead Subtree) {
	var head uintptr
	_ = head
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(739), uintptr(unsafe.Pointer(&__func__98)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	(*StackHead)(unsafe.Pointer(head)).Fstatus = int32(StackStatusPaused)
	(*StackHead)(unsafe.Pointer(head)).Flookahead_when_paused = lookahead
	(*StackHead)(unsafe.Pointer(head)).Fnode_count_at_last_error = (*StackNode)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fnode)).Fnode_count
}

var __func__98 = [15]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'p', 'a', 'u', 's', 'e'}

func ts_stack_is_active(tls *libc.TLS, self uintptr, version StackVersion) (r uint8) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(746), uintptr(unsafe.Pointer(&__func__99)))
		}
	}
	return libc.BoolUint8((*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents+uintptr(version)*48)).Fstatus == int32(StackStatusActive))
}

var __func__99 = [19]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'i', 's', '_', 'a', 'c', 't', 'i', 'v', 'e'}

func ts_stack_is_halted(tls *libc.TLS, self uintptr, version StackVersion) (r uint8) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(750), uintptr(unsafe.Pointer(&__func__100)))
		}
	}
	return libc.BoolUint8((*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents+uintptr(version)*48)).Fstatus == int32(StackStatusHalted))
}

var __func__100 = [19]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'i', 's', '_', 'h', 'a', 'l', 't', 'e', 'd'}

func ts_stack_is_paused(tls *libc.TLS, self uintptr, version StackVersion) (r uint8) {
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(754), uintptr(unsafe.Pointer(&__func__101)))
		}
	}
	return libc.BoolUint8((*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents+uintptr(version)*48)).Fstatus == int32(StackStatusPaused))
}

var __func__101 = [19]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'i', 's', '_', 'p', 'a', 'u', 's', 'e', 'd'}

func ts_stack_resume(tls *libc.TLS, self uintptr, version StackVersion) (r Subtree) {
	var head uintptr
	var result Subtree
	_, _ = head, result
	_ = libc.Uint64FromInt64(4)
	{
		if !(version < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+8902, __ccgo_ts+8787, int32(758), uintptr(unsafe.Pointer(&__func__102)))
		}
	}
	head = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(version)*48
	_ = libc.Uint64FromInt64(4)
	{
		if !((*StackHead)(unsafe.Pointer(head)).Fstatus == int32(StackStatusPaused)) {
			libc.X__assert_fail(tls, __ccgo_ts+9522, __ccgo_ts+8787, int32(759), uintptr(unsafe.Pointer(&__func__102)))
		}
	}
	result = (*StackHead)(unsafe.Pointer(head)).Flookahead_when_paused
	(*StackHead)(unsafe.Pointer(head)).Fstatus = int32(StackStatusActive)
	(*StackHead)(unsafe.Pointer(head)).Flookahead_when_paused = Subtree{}
	return result
}

var __func__102 = [16]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'r', 'e', 's', 'u', 'm', 'e'}

func ts_stack_clear(tls *libc.TLS, self uintptr) {
	var i, v2 uint32_t
	var v3 uintptr
	_, _, _ = i, v2, v3
	stack_node_retain(tls, (*Stack)(unsafe.Pointer(self)).Fbase_node)
	i = uint32(0)
	for {
		if !(i < (*Stack)(unsafe.Pointer(self)).Fheads.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+9054, __ccgo_ts+8787, int32(769), uintptr(unsafe.Pointer(&__func__103)))
			}
		}
		stack_head_delete(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents+uintptr(i)*48, self+48, (*Stack)(unsafe.Pointer(self)).Fsubtree_pool)
		goto _1
	_1:
		;
		i = i + 1
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(48))
	v3 = self + 8
	v2 = *(*uint32_t)(unsafe.Pointer(v3))
	*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
	*(*StackHead)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self)).Fcontents + uintptr(v2)*48)) = StackHead{
		Fnode:                  (*Stack)(unsafe.Pointer(self)).Fbase_node,
		Flast_external_token:   Subtree{},
		Flookahead_when_paused: Subtree{},
	}
}

var __func__103 = [15]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'c', 'l', 'e', 'a', 'r'}

func ts_stack_print_dot_graph(tls *libc.TLS, self uintptr, language uintptr, f uintptr) (r uint8) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var all_iterators_done, quoted uint8
	var data, head, next_iterator, node, state, v5 uintptr
	var i, i1, j, j1, j2, v4 uint32_t
	var iterator StackIterator
	var j3 int32
	var _ /* link at bp+16 */ StackLink
	var _ /* visited_nodes at bp+0 */ struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = all_iterators_done, data, head, i, i1, iterator, j, j1, j2, j3, next_iterator, node, quoted, state, v4, v5
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fcontents = _array__reserve(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+32)).Fcontents, self+32+12, libc.Uint64FromInt64(32), uint32(32))
	if !(f != 0) {
		f = libc.Xstderr
	}
	libc.Xfprintf(tls, f, __ccgo_ts+9556, 0)
	libc.Xfprintf(tls, f, __ccgo_ts+9573, 0)
	libc.Xfprintf(tls, f, __ccgo_ts+9588, 0)
	*(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp)) = struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}{}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 32)).Fsize = uint32(0)
	i = uint32(0)
	for {
		if !(i < (*Stack)(unsafe.Pointer(self)).Fheads.Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+9054, __ccgo_ts+8787, int32(792), uintptr(unsafe.Pointer(&__func__104)))
			}
		}
		head = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self)).Fcontents + uintptr(i)*48
		if (*StackHead)(unsafe.Pointer(head)).Fstatus == int32(StackStatusHalted) {
			goto _1
		}
		libc.Xfprintf(tls, f, __ccgo_ts+9611, libc.VaList(bp+48, i))
		libc.Xfprintf(tls, f, __ccgo_ts+9648, libc.VaList(bp+48, i, (*StackHead)(unsafe.Pointer(head)).Fnode))
		if (*StackHead)(unsafe.Pointer(head)).Fstatus == int32(StackStatusPaused) {
			libc.Xfprintf(tls, f, __ccgo_ts+9674, 0)
		}
		libc.Xfprintf(tls, f, __ccgo_ts+9685, libc.VaList(bp+48, i, ts_stack_node_count_since_error(tls, self, i), ts_stack_error_cost(tls, self, i)))
		if (*StackHead)(unsafe.Pointer(head)).Fsummary != 0 {
			libc.Xfprintf(tls, f, __ccgo_ts+9769, 0)
			j = uint32(0)
			for {
				if !(j < (*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(j < (*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+9779, __ccgo_ts+8787, int32(810), uintptr(unsafe.Pointer(&__func__104)))
					}
				}
				libc.Xfprintf(tls, f, __ccgo_ts+9817, libc.VaList(bp+48, libc.Int32FromUint16((*StackSummaryEntry)(unsafe.Pointer((*StackSummary)(unsafe.Pointer((*StackHead)(unsafe.Pointer(head)).Fsummary)).Fcontents+uintptr(j)*20)).Fstate)))
				goto _2
			_2:
				;
				j = j + 1
			}
		}
		if *(*uintptr)(unsafe.Pointer(head + 24)) != 0 {
			state = *(*uintptr)(unsafe.Pointer(head + 24)) + 48
			data = ts_external_scanner_state_data(tls, state)
			libc.Xfprintf(tls, f, __ccgo_ts+9821, 0)
			j1 = uint32(0)
			for {
				if !(j1 < (*ExternalScannerState)(unsafe.Pointer(state)).Flength) {
					break
				}
				libc.Xfprintf(tls, f, __ccgo_ts+9846, libc.VaList(bp+48, int32(*(*int8)(unsafe.Pointer(data + uintptr(j1))))))
				goto _3
			_3:
				;
				j1 = j1 + 1
			}
		}
		libc.Xfprintf(tls, f, __ccgo_ts+9851, 0)
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 32)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+32)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+32)).Fsize, self+32+12, uint32(1), libc.Uint64FromInt64(32))
		v5 = self + 32 + 8
		v4 = *(*uint32_t)(unsafe.Pointer(v5))
		*(*uint32_t)(unsafe.Pointer(v5)) = *(*uint32_t)(unsafe.Pointer(v5)) + 1
		*(*StackIterator)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+32)).Fcontents + uintptr(v4)*32)) = StackIterator{
			Fnode: (*StackHead)(unsafe.Pointer(head)).Fnode,
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	all_iterators_done = libc.BoolUint8(0 != 0)
	for !(all_iterators_done != 0) {
		all_iterators_done = libc.BoolUint8(1 != 0)
		i1 = uint32(0)
		for {
			if !(i1 < (*Stack)(unsafe.Pointer(self)).Fiterators.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i1 < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+32)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+8945, __ccgo_ts+8787, int32(831), uintptr(unsafe.Pointer(&__func__104)))
				}
			}
			iterator = *(*StackIterator)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+32)).Fcontents + uintptr(i1)*32))
			node = iterator.Fnode
			j2 = uint32(0)
			for {
				if !(j2 < (*(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp))).Fsize) {
					break
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(j2 < (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(bp)).Fsize) {
						libc.X__assert_fail(tls, __ccgo_ts+9855, __ccgo_ts+8787, int32(835), uintptr(unsafe.Pointer(&__func__104)))
					}
				}
				if *(*uintptr)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(bp)).Fcontents + uintptr(j2)*8)) == node {
					node = libc.UintptrFromInt32(0)
					break
				}
				goto _7
			_7:
				;
				j2 = j2 + 1
			}
			if !(node != 0) {
				goto _6
			}
			all_iterators_done = libc.BoolUint8(0 != 0)
			libc.Xfprintf(tls, f, __ccgo_ts+9894, libc.VaList(bp+48, node))
			if libc.Int32FromUint16((*StackNode)(unsafe.Pointer(node)).Fstate) == 0 {
				libc.Xfprintf(tls, f, __ccgo_ts+9904, 0)
			} else {
				if libc.Int32FromUint16((*StackNode)(unsafe.Pointer(node)).Flink_count) == int32(1) && *(*uintptr)(unsafe.Pointer(node + 16 + 8)) != 0 && ts_subtree_extra(tls, (*(*StackLink)(unsafe.Pointer(node + 16))).Fsubtree) != 0 {
					libc.Xfprintf(tls, f, __ccgo_ts+9914, 0)
				} else {
					libc.Xfprintf(tls, f, __ccgo_ts+9944, libc.VaList(bp+48, libc.Int32FromUint16((*StackNode)(unsafe.Pointer(node)).Fstate)))
				}
			}
			libc.Xfprintf(tls, f, __ccgo_ts+9955, libc.VaList(bp+48, (*StackNode)(unsafe.Pointer(node)).Fposition.Fextent.Frow+uint32(1), (*StackNode)(unsafe.Pointer(node)).Fposition.Fextent.Fcolumn, (*StackNode)(unsafe.Pointer(node)).Fnode_count, (*StackNode)(unsafe.Pointer(node)).Ferror_cost, (*StackNode)(unsafe.Pointer(node)).Fdynamic_precedence))
			j3 = 0
			for {
				if !(j3 < libc.Int32FromUint16((*StackNode)(unsafe.Pointer(node)).Flink_count)) {
					break
				}
				*(*StackLink)(unsafe.Pointer(bp + 16)) = StackLink{}
				*(*struct {
					Fnode    uintptr
					Fsubtree struct {
						Fptr  [0]uintptr
						Fdata SubtreeInlineData
					}
					Fis_pending uint8
				})(unsafe.Pointer(bp + 16)) = *(*StackLink)(unsafe.Pointer(node + 16 + uintptr(j3)*24))
				libc.Xfprintf(tls, f, __ccgo_ts+10037, libc.VaList(bp+48, node, (*(*StackLink)(unsafe.Pointer(bp + 16))).Fnode))
				if (*(*StackLink)(unsafe.Pointer(bp + 16))).Fis_pending != 0 {
					libc.Xfprintf(tls, f, __ccgo_ts+10058, 0)
				}
				if *(*uintptr)(unsafe.Pointer(bp + 16 + 8)) != 0 && ts_subtree_extra(tls, (*(*StackLink)(unsafe.Pointer(bp + 16))).Fsubtree) != 0 {
					libc.Xfprintf(tls, f, __ccgo_ts+10072, 0)
				}
				if !(*(*uintptr)(unsafe.Pointer(bp + 16 + 8)) != 0) {
					libc.Xfprintf(tls, f, __ccgo_ts+10088, 0)
				} else {
					libc.Xfprintf(tls, f, __ccgo_ts+10098, 0)
					quoted = libc.BoolUint8(ts_subtree_visible(tls, (*(*StackLink)(unsafe.Pointer(bp + 16))).Fsubtree) != 0 && !(ts_subtree_named(tls, (*(*StackLink)(unsafe.Pointer(bp + 16))).Fsubtree) != 0))
					if quoted != 0 {
						libc.Xfprintf(tls, f, __ccgo_ts+10106, 0)
					}
					ts_language_write_symbol_as_dot_string(tls, language, f, ts_subtree_symbol(tls, (*(*StackLink)(unsafe.Pointer(bp + 16))).Fsubtree))
					if quoted != 0 {
						libc.Xfprintf(tls, f, __ccgo_ts+10106, 0)
					}
					libc.Xfprintf(tls, f, __ccgo_ts+10108, 0)
					libc.Xfprintf(tls, f, __ccgo_ts+10110, libc.VaList(bp+48, ts_subtree_error_cost(tls, (*(*StackLink)(unsafe.Pointer(bp + 16))).Fsubtree), ts_subtree_dynamic_precedence(tls, (*(*StackLink)(unsafe.Pointer(bp + 16))).Fsubtree)))
				}
				libc.Xfprintf(tls, f, __ccgo_ts+10163, 0)
				if j3 == 0 {
					_ = libc.Uint64FromInt64(4)
					{
						if !(i1 < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+32)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8945, __ccgo_ts+8787, int32(894), uintptr(unsafe.Pointer(&__func__104)))
						}
					}
					next_iterator = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents + uintptr(i1)*32
				} else {
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 32)).Fcontents = _array__grow(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fsize, self+32+12, uint32(1), libc.Uint64FromInt64(32))
					v5 = self + 32 + 8
					v4 = *(*uint32_t)(unsafe.Pointer(v5))
					*(*uint32_t)(unsafe.Pointer(v5)) = *(*uint32_t)(unsafe.Pointer(v5)) + 1
					*(*StackIterator)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents + uintptr(v4)*32)) = iterator
					_ = libc.Uint64FromInt64(4)
					{
						if !((*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+32)).Fsize-libc.Uint32FromInt32(1) < (*struct {
							Fcontents uintptr
							Fsize     uint32_t
							Fcapacity uint32_t
						})(unsafe.Pointer(self+32)).Fsize) {
							libc.X__assert_fail(tls, __ccgo_ts+8986, __ccgo_ts+8787, int32(897), uintptr(unsafe.Pointer(&__func__104)))
						}
					}
					next_iterator = (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fcontents + uintptr((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+32)).Fsize-uint32(1))*32
				}
				(*StackIterator)(unsafe.Pointer(next_iterator)).Fnode = (*(*StackLink)(unsafe.Pointer(bp + 16))).Fnode
				goto _8
			_8:
				;
				j3 = j3 + 1
			}
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fsize, bp+12, uint32(1), libc.Uint64FromInt64(8))
			v5 = bp + 8
			v4 = *(*uint32_t)(unsafe.Pointer(v5))
			*(*uint32_t)(unsafe.Pointer(v5)) = *(*uint32_t)(unsafe.Pointer(v5)) + 1
			*(*uintptr)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp)).Fcontents + uintptr(v4)*8)) = node
			goto _6
		_6:
			;
			i1 = i1 + 1
		}
	}
	libc.Xfprintf(tls, f, __ccgo_ts+10167, 0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp)).Fcapacity = uint32(0)
	return libc.BoolUint8(1 != 0)
}

var __func__104 = [25]int8{'t', 's', '_', 's', 't', 'a', 'c', 'k', '_', 'p', 'r', 'i', 'n', 't', '_', 'd', 'o', 't', '_', 'g', 'r', 'a', 'p', 'h'}

const _ISupper = 256
const _ISlower = 512
const _ISalpha = 1024
const _ISdigit = 2048
const _ISxdigit = 4096
const _ISspace = 8192
const _ISprint = 16384
const _ISgraph = 32768
const _ISblank = 1
const _IScntrl = 2
const _ISpunct = 4
const _ISalnum = 8

func atomic_load(tls *libc.TLS, p uintptr) (r size_t) {
	return libc.AtomicLoadNUint64(p, libc.Int32FromInt32(0))
}

func atomic_inc(tls *libc.TLS, p uintptr) (r uint32_t) {
	return __atomic_add_fetch(tls, p, uint32(1), int32(5))
}

func atomic_dec(tls *libc.TLS, p uintptr) (r uint32_t) {
	return __atomic_sub_fetch(tls, p, uint32(1), int32(5))
}

type Edit = struct {
	Fstart   Length
	Fold_end Length
	Fnew_end Length
}

func ts_external_scanner_state_init(tls *libc.TLS, self uintptr, data uintptr, length uint32) {
	(*ExternalScannerState)(unsafe.Pointer(self)).Flength = length
	if uint64(length) > uint64(24) {
		(*ExternalScannerState)(unsafe.Pointer(self)).F__ccgo0_0.Flong_data = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(length))
		libc.Xmemcpy(tls, (*ExternalScannerState)(unsafe.Pointer(self)).F__ccgo0_0.Flong_data, data, uint64(length))
	} else {
		libc.Xmemcpy(tls, self, data, uint64(length))
	}
}

func ts_external_scanner_state_copy(tls *libc.TLS, self uintptr) (r ExternalScannerState) {
	var result ExternalScannerState
	_ = result
	result = *(*ExternalScannerState)(unsafe.Pointer(self))
	if uint64((*ExternalScannerState)(unsafe.Pointer(self)).Flength) > uint64(24) {
		result.F__ccgo0_0.Flong_data = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64((*ExternalScannerState)(unsafe.Pointer(self)).Flength))
		libc.Xmemcpy(tls, result.F__ccgo0_0.Flong_data, (*ExternalScannerState)(unsafe.Pointer(self)).F__ccgo0_0.Flong_data, uint64((*ExternalScannerState)(unsafe.Pointer(self)).Flength))
	}
	return result
}

func ts_external_scanner_state_delete(tls *libc.TLS, self uintptr) {
	if uint64((*ExternalScannerState)(unsafe.Pointer(self)).Flength) > uint64(24) {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*ExternalScannerState)(unsafe.Pointer(self)).F__ccgo0_0.Flong_data)
	}
}

func ts_external_scanner_state_data(tls *libc.TLS, self uintptr) (r uintptr) {
	if uint64((*ExternalScannerState)(unsafe.Pointer(self)).Flength) > uint64(24) {
		return (*ExternalScannerState)(unsafe.Pointer(self)).F__ccgo0_0.Flong_data
	} else {
		return self
	}
	return r
}

func ts_external_scanner_state_eq(tls *libc.TLS, self uintptr, buffer uintptr, length uint32) (r uint8) {
	return libc.BoolUint8((*ExternalScannerState)(unsafe.Pointer(self)).Flength == length && libc.Xmemcmp(tls, ts_external_scanner_state_data(tls, self), buffer, uint64(length)) == 0)
}

func ts_subtree_array_copy(tls *libc.TLS, self SubtreeArray, dest uintptr) {
	var i uint32_t
	_ = i
	(*SubtreeArray)(unsafe.Pointer(dest)).Fsize = self.Fsize
	(*SubtreeArray)(unsafe.Pointer(dest)).Fcapacity = self.Fcapacity
	(*SubtreeArray)(unsafe.Pointer(dest)).Fcontents = self.Fcontents
	if self.Fcapacity > uint32(0) {
		(*SubtreeArray)(unsafe.Pointer(dest)).Fcontents = (*(*func(*libc.TLS, size_t, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_calloc})))(tls, uint64(self.Fcapacity), uint64(8))
		libc.Xmemcpy(tls, (*SubtreeArray)(unsafe.Pointer(dest)).Fcontents, self.Fcontents, uint64(self.Fsize)*uint64(8))
		i = uint32(0)
		for {
			if !(i < self.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*SubtreeArray)(unsafe.Pointer(dest)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+10170, __ccgo_ts+10199, int32(76), uintptr(unsafe.Pointer(&__func__105)))
				}
			}
			ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(dest)).Fcontents + uintptr(i)*8)))
			goto _1
		_1:
			;
			i = i + 1
		}
	}
}

var __func__105 = [22]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'a', 'r', 'r', 'a', 'y', '_', 'c', 'o', 'p', 'y'}

func ts_subtree_array_clear(tls *libc.TLS, pool uintptr, self uintptr) {
	var i uint32_t
	_ = i
	i = uint32(0)
	for {
		if !(i < (*SubtreeArray)(unsafe.Pointer(self)).Fsize) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*SubtreeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+382, __ccgo_ts+10199, int32(83), uintptr(unsafe.Pointer(&__func__106)))
			}
		}
		ts_subtree_release(tls, pool, *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(i)*8)))
		goto _1
	_1:
		;
		i = i + 1
	}
	(*SubtreeArray)(unsafe.Pointer(self)).Fsize = uint32(0)
}

var __func__106 = [23]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'a', 'r', 'r', 'a', 'y', '_', 'c', 'l', 'e', 'a', 'r'}

func ts_subtree_array_delete(tls *libc.TLS, pool uintptr, self uintptr) {
	ts_subtree_array_clear(tls, pool, self)
	if (*SubtreeArray)(unsafe.Pointer(self)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*SubtreeArray)(unsafe.Pointer(self)).Fcontents)
	}
	(*SubtreeArray)(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
	(*SubtreeArray)(unsafe.Pointer(self)).Fsize = uint32(0)
	(*SubtreeArray)(unsafe.Pointer(self)).Fcapacity = uint32(0)
}

func ts_subtree_array_remove_trailing_extras(tls *libc.TLS, self uintptr, destination uintptr) {
	var last Subtree
	var v1 uint32_t
	var v2 uintptr
	_, _, _ = last, v1, v2
	(*SubtreeArray)(unsafe.Pointer(destination)).Fsize = uint32(0)
	for (*SubtreeArray)(unsafe.Pointer(self)).Fsize > uint32(0) {
		_ = libc.Uint64FromInt64(4)
		{
			if !((*SubtreeArray)(unsafe.Pointer(self)).Fsize-libc.Uint32FromInt32(1) < (*SubtreeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+10243, __ccgo_ts+10199, int32(99), uintptr(unsafe.Pointer(&__func__107)))
			}
		}
		last = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr((*SubtreeArray)(unsafe.Pointer(self)).Fsize-uint32(1))*8))
		if ts_subtree_extra(tls, last) != 0 {
			(*SubtreeArray)(unsafe.Pointer(self)).Fsize = (*SubtreeArray)(unsafe.Pointer(self)).Fsize - 1
			(*SubtreeArray)(unsafe.Pointer(destination)).Fcontents = _array__grow(tls, (*SubtreeArray)(unsafe.Pointer(destination)).Fcontents, (*SubtreeArray)(unsafe.Pointer(destination)).Fsize, destination+12, uint32(1), libc.Uint64FromInt64(8))
			v2 = destination + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(destination)).Fcontents + uintptr(v1)*8)) = last
		} else {
			break
		}
	}
	ts_subtree_array_reverse(tls, destination)
}

var __func__107 = [40]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'a', 'r', 'r', 'a', 'y', '_', 'r', 'e', 'm', 'o', 'v', 'e', '_', 't', 'r', 'a', 'i', 'l', 'i', 'n', 'g', '_', 'e', 'x', 't', 'r', 'a', 's'}

func ts_subtree_array_reverse(tls *libc.TLS, self uintptr) {
	var i, limit uint32_t
	var reverse_index size_t
	var swap Subtree
	_, _, _, _ = i, limit, reverse_index, swap
	i = uint32(0)
	limit = (*SubtreeArray)(unsafe.Pointer(self)).Fsize / uint32(2)
	for {
		if !(i < limit) {
			break
		}
		reverse_index = uint64((*SubtreeArray)(unsafe.Pointer(self)).Fsize - uint32(1) - i)
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*SubtreeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+382, __ccgo_ts+10199, int32(113), uintptr(unsafe.Pointer(&__func__108)))
			}
		}
		swap = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(i)*8))
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*SubtreeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+382, __ccgo_ts+10199, int32(114), uintptr(unsafe.Pointer(&__func__108)))
			}
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(reverse_index) < (*SubtreeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+10285, __ccgo_ts+10199, int32(114), uintptr(unsafe.Pointer(&__func__108)))
			}
		}
		*(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(i)*8)) = *(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(reverse_index)*8))
		_ = libc.Uint64FromInt64(4)
		{
			if !(uint32(reverse_index) < (*SubtreeArray)(unsafe.Pointer(self)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+10285, __ccgo_ts+10199, int32(115), uintptr(unsafe.Pointer(&__func__108)))
			}
		}
		*(*Subtree)(unsafe.Pointer((*SubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(reverse_index)*8)) = swap
		goto _1
	_1:
		;
		i = i + 1
	}
}

var __func__108 = [25]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'a', 'r', 'r', 'a', 'y', '_', 'r', 'e', 'v', 'e', 'r', 's', 'e'}

func ts_subtree_pool_new(tls *libc.TLS, capacity uint32_t) (r SubtreePool) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* self at bp+0 */ SubtreePool
	*(*SubtreePool)(unsafe.Pointer(bp)) = SubtreePool{}
	(*MutableSubtreeArray)(unsafe.Pointer(bp)).Fcontents = _array__reserve(tls, (*MutableSubtreeArray)(unsafe.Pointer(bp)).Fcontents, bp+12, libc.Uint64FromInt64(8), capacity)
	return *(*SubtreePool)(unsafe.Pointer(bp))
}

func ts_subtree_pool_delete(tls *libc.TLS, self uintptr) {
	var i uint32
	_ = i
	if (*SubtreePool)(unsafe.Pointer(self)).Ffree_trees.Fcontents != 0 {
		i = uint32(0)
		for {
			if !(i < (*SubtreePool)(unsafe.Pointer(self)).Ffree_trees.Fsize) {
				break
			}
			_ = libc.Uint64FromInt64(4)
			{
				if !(i < (*MutableSubtreeArray)(unsafe.Pointer(self)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+10326, __ccgo_ts+10199, int32(130), uintptr(unsafe.Pointer(&__func__109)))
				}
			}
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, *(*uintptr)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(i)*8)))
			goto _1
		_1:
			;
			i = i + 1
		}
		if (*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents)
		}
		(*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents = libc.UintptrFromInt32(0)
		(*MutableSubtreeArray)(unsafe.Pointer(self)).Fsize = uint32(0)
		(*MutableSubtreeArray)(unsafe.Pointer(self)).Fcapacity = uint32(0)
	}
	if (*SubtreePool)(unsafe.Pointer(self)).Ftree_stack.Fcontents != 0 {
		if (*MutableSubtreeArray)(unsafe.Pointer(self+16)).Fcontents != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*MutableSubtreeArray)(unsafe.Pointer(self+16)).Fcontents)
		}
		(*MutableSubtreeArray)(unsafe.Pointer(self + 16)).Fcontents = libc.UintptrFromInt32(0)
		(*MutableSubtreeArray)(unsafe.Pointer(self + 16)).Fsize = uint32(0)
		(*MutableSubtreeArray)(unsafe.Pointer(self + 16)).Fcapacity = uint32(0)
	}
}

var __func__109 = [23]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'p', 'o', 'o', 'l', '_', 'd', 'e', 'l', 'e', 't', 'e'}

func ts_subtree_pool_allocate(tls *libc.TLS, self uintptr) (r uintptr) {
	var v1 uint32_t
	var v2 uintptr
	_, _ = v1, v2
	if (*SubtreePool)(unsafe.Pointer(self)).Ffree_trees.Fsize > uint32(0) {
		v2 = self + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		return *(*uintptr)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(v1)*8))
	} else {
		return (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(80))
	}
	return r
}

func ts_subtree_pool_free(tls *libc.TLS, self uintptr, tree uintptr) {
	var v1 uint32_t
	var v2 uintptr
	_, _ = v1, v2
	if (*SubtreePool)(unsafe.Pointer(self)).Ffree_trees.Fcapacity > uint32(0) && (*SubtreePool)(unsafe.Pointer(self)).Ffree_trees.Fsize+uint32(1) <= uint32(32) {
		(*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(self)).Fsize, self+12, uint32(1), libc.Uint64FromInt64(8))
		v2 = self + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(self)).Fcontents + uintptr(v1)*8)) = *(*MutableSubtree)(unsafe.Pointer(&struct{ f uintptr }{f: tree}))
	} else {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, tree)
	}
}

func ts_subtree_can_inline(tls *libc.TLS, padding Length, size Length, lookahead_bytes uint32_t) (r uint8) {
	return libc.BoolUint8(padding.Fbytes < libc.Uint32FromInt32(libc.Int32FromInt32(255)) && padding.Fextent.Frow < uint32(16) && padding.Fextent.Fcolumn < libc.Uint32FromInt32(libc.Int32FromInt32(255)) && size.Fbytes < libc.Uint32FromInt32(libc.Int32FromInt32(255)) && size.Fextent.Frow == uint32(0) && size.Fextent.Fcolumn < libc.Uint32FromInt32(libc.Int32FromInt32(255)) && lookahead_bytes < uint32(16))
}

func ts_subtree_new_leaf(tls *libc.TLS, pool uintptr, symbol TSSymbol, padding Length, size Length, lookahead_bytes uint32_t, parse_state TSStateId, has_external_tokens uint8, depends_on_column uint8, is_keyword uint8, language uintptr) (r Subtree) {
	var data uintptr
	var extra, is_inline uint8
	var metadata TSSymbolMetadata
	_, _, _, _ = data, extra, is_inline, metadata
	metadata = ts_language_symbol_metadata(tls, language, symbol)
	extra = libc.BoolUint8(libc.Int32FromUint16(symbol) == 0)
	is_inline = libc.BoolUint8(libc.Int32FromUint16(symbol) <= int32(255) && !(has_external_tokens != 0) && ts_subtree_can_inline(tls, padding, size, lookahead_bytes) != 0)
	if is_inline != 0 {
		return *(*Subtree)(unsafe.Pointer(&SubtreeInlineData{
			F__ccgo0:         libc.Uint8FromInt32(1)&0x1<<0 | metadata.Fvisible&0x1<<1 | metadata.Fnamed&0x1<<2 | extra&0x1<<3 | libc.Uint8FromInt32(0)&0x1<<4 | libc.Uint8FromInt32(0)&0x1<<5 | is_keyword&0x1<<6,
			Fsymbol:          uint8(symbol),
			Fparse_state:     parse_state,
			Fpadding_columns: uint8(padding.Fextent.Fcolumn),
			F__ccgo5:         uint8(padding.Fextent.Frow)&0xf<<0 | uint8(lookahead_bytes)&0xf<<4,
			Fpadding_bytes:   uint8(padding.Fbytes),
			Fsize_bytes:      uint8(size.Fbytes),
		}))
	} else {
		data = ts_subtree_pool_allocate(tls, pool)
		*(*SubtreeHeapData)(unsafe.Pointer(data)) = SubtreeHeapData{
			Fref_count:       uint32(1),
			Fpadding:         padding,
			Fsize:            size,
			Flookahead_bytes: lookahead_bytes,
			Fsymbol:          symbol,
			Fparse_state:     parse_state,
			F__ccgo44:        metadata.Fvisible&0x1<<0 | metadata.Fnamed&0x1<<1 | extra&0x1<<2 | libc.Uint8FromInt32(0)&0x1<<3 | libc.Uint8FromInt32(0)&0x1<<4 | libc.Uint8FromInt32(0)&0x1<<5 | has_external_tokens&0x1<<6 | libc.Uint8FromInt32(0)&0x1<<7,
			F__ccgo45:        depends_on_column&0x1<<0 | libc.Uint8FromInt32(0)&0x1<<1 | is_keyword&0x1<<2,
		}
		return *(*Subtree)(unsafe.Pointer(&struct{ f uintptr }{f: data}))
	}
	return r
}

func ts_subtree_set_symbol(tls *libc.TLS, self uintptr, symbol TSSymbol, language uintptr) {
	var metadata TSSymbolMetadata
	_ = metadata
	metadata = ts_language_symbol_metadata(tls, language, symbol)
	if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Int32FromUint16(symbol) < libc.Int32FromInt32(255)) {
				libc.X__assert_fail(tls, __ccgo_ts+10368, __ccgo_ts+10199, int32(233), uintptr(unsafe.Pointer(&__func__110)))
			}
		}
		(*MutableSubtree)(unsafe.Pointer(self)).Fdata.Fsymbol = uint8(symbol)
		libc.SetBitFieldPtr8Uint8(self+0, metadata.Fnamed, 2, 0x4)
		libc.SetBitFieldPtr8Uint8(self+0, metadata.Fvisible, 1, 0x2)
	} else {
		(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fsymbol = symbol
		libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(self))+44, metadata.Fnamed, 1, 0x2)
		libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(self))+44, metadata.Fvisible, 0, 0x1)
	}
}

var __func__110 = [22]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 's', 'e', 't', '_', 's', 'y', 'm', 'b', 'o', 'l'}

func ts_subtree_new_error(tls *libc.TLS, pool uintptr, lookahead_char int32_t, padding Length, size Length, bytes_scanned uint32_t, parse_state TSStateId, language uintptr) (r Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var data uintptr
	var _ /* result at bp+0 */ Subtree
	_ = data
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = ts_subtree_new_leaf(tls, pool, libc.Uint16FromInt32(-libc.Int32FromInt32(1)), padding, size, bytes_scanned, parse_state, libc.BoolUint8(0 != 0), libc.BoolUint8(0 != 0), libc.BoolUint8(0 != 0), language)
	data = *(*uintptr)(unsafe.Pointer(bp))
	libc.SetBitFieldPtr8Uint8(data+44, libc.BoolUint8(1 != 0), 3, 0x8)
	libc.SetBitFieldPtr8Uint8(data+44, libc.BoolUint8(1 != 0), 4, 0x10)
	*(*int32_t)(unsafe.Add(unsafe.Pointer(data), 48)) = lookahead_char
	return *(*Subtree)(unsafe.Pointer(bp))
}

func ts_subtree_clone(tls *libc.TLS, _self Subtree) (r MutableSubtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var alloc_size size_t
	var i uint32_t
	var new_children, old_children, result, v1 uintptr
	_, _, _, _, _, _ = alloc_size, i, new_children, old_children, result, v1
	alloc_size = ts_subtree_alloc_size(tls, (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)
	new_children = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, alloc_size)
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
	}
	old_children = v1
	libc.Xmemcpy(tls, new_children, old_children, alloc_size)
	result = new_children + uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
	if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count > uint32(0) {
		i = uint32(0)
		for {
			if !(i < (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count) {
				break
			}
			ts_subtree_retain(tls, *(*Subtree)(unsafe.Pointer(new_children + uintptr(i)*8)))
			goto _2
		_2:
			;
			i = i + 1
		}
	} else {
		if int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x40>>6) != 0 {
			*(*ExternalScannerState)(unsafe.Add(unsafe.Pointer(result), 48)) = ts_external_scanner_state_copy(tls, *(*uintptr)(unsafe.Pointer(bp))+48)
		}
	}
	libc.AtomicStorePUint32(result, libc.Uint32FromInt32(1))
	return *(*MutableSubtree)(unsafe.Pointer(&struct{ f uintptr }{f: result}))
}

func ts_subtree_make_mut(tls *libc.TLS, pool uintptr, _self Subtree) (r MutableSubtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var result MutableSubtree
	_ = result
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		return *(*MutableSubtree)(unsafe.Pointer(&struct{ f SubtreeInlineData }{f: *(*SubtreeInlineData)(unsafe.Pointer(bp))}))
	}
	if libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp))) == uint32(1) {
		return ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(bp)))
	}
	result = ts_subtree_clone(tls, *(*Subtree)(unsafe.Pointer(bp)))
	ts_subtree_release(tls, pool, *(*Subtree)(unsafe.Pointer(bp)))
	return result
}

func ts_subtree_compress(tls *libc.TLS, self MutableSubtree, count uint32, language uintptr, stack uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var grandchild1 MutableSubtree
	var i, initial_stack_size uint32
	var symbol TSSymbol
	var v8 uint32_t
	var v2, v3 uintptr
	var _ /* child at bp+24 */ MutableSubtree
	var _ /* child at bp+8 */ MutableSubtree
	var _ /* grandchild at bp+16 */ MutableSubtree
	var _ /* tree at bp+0 */ MutableSubtree
	_, _, _, _, _, _, _ = grandchild1, i, initial_stack_size, symbol, v2, v3, v8
	initial_stack_size = (*MutableSubtreeArray)(unsafe.Pointer(stack)).Fsize
	*(*MutableSubtree)(unsafe.Pointer(bp)) = MutableSubtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = self
	symbol = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol
	i = uint32(0)
	for {
		if !(i < count) {
			break
		}
		if libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp))) > uint32(1) || (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count < uint32(2) {
			break
		}
		*(*MutableSubtree)(unsafe.Pointer(bp + 8)) = MutableSubtree{}
		if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
			v2 = libc.UintptrFromInt32(0)
		} else {
			v2 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
		}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 8)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(v2)))
		if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 || (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count < uint32(2) || libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp + 8))) > uint32(1) || libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fsymbol) != libc.Int32FromUint16(symbol) {
			break
		}
		*(*MutableSubtree)(unsafe.Pointer(bp + 16)) = MutableSubtree{}
		if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 {
			v2 = libc.UintptrFromInt32(0)
		} else {
			v2 = *(*uintptr)(unsafe.Pointer(bp + 8)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count)*8
		}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 16)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(v2)))
		if int32(*(*uint8)(unsafe.Pointer(bp + 16 + 0))&0x1>>0) != 0 || (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 16)))).Fchild_count < uint32(2) || libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp + 16))) > uint32(1) || libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 16)))).Fsymbol) != libc.Int32FromUint16(symbol) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
			v2 = libc.UintptrFromInt32(0)
		} else {
			v2 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
		}
		*(*Subtree)(unsafe.Pointer(v2)) = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 16)))
		if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 {
			v2 = libc.UintptrFromInt32(0)
		} else {
			v2 = *(*uintptr)(unsafe.Pointer(bp + 8)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count)*8
		}
		if int32(*(*uint8)(unsafe.Pointer(bp + 16 + 0))&0x1>>0) != 0 {
			v3 = libc.UintptrFromInt32(0)
		} else {
			v3 = *(*uintptr)(unsafe.Pointer(bp + 16)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 16)))).Fchild_count)*8
		}
		*(*Subtree)(unsafe.Pointer(v2)) = *(*Subtree)(unsafe.Pointer(v3 + uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 16)))).Fchild_count-uint32(1))*8))
		if int32(*(*uint8)(unsafe.Pointer(bp + 16 + 0))&0x1>>0) != 0 {
			v2 = libc.UintptrFromInt32(0)
		} else {
			v2 = *(*uintptr)(unsafe.Pointer(bp + 16)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 16)))).Fchild_count)*8
		}
		*(*Subtree)(unsafe.Pointer(v2 + uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 16)))).Fchild_count-uint32(1))*8)) = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 8)))
		(*MutableSubtreeArray)(unsafe.Pointer(stack)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(stack)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(stack)).Fsize, stack+12, uint32(1), libc.Uint64FromInt64(8))
		v2 = stack + 8
		v8 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(stack)).Fcontents + uintptr(v8)*8)) = *(*MutableSubtree)(unsafe.Pointer(bp))
		*(*MutableSubtree)(unsafe.Pointer(bp)) = *(*MutableSubtree)(unsafe.Pointer(bp + 16))
		goto _1
	_1:
		;
		i = i + 1
	}
	for (*MutableSubtreeArray)(unsafe.Pointer(stack)).Fsize > initial_stack_size {
		v2 = stack + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v8 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*MutableSubtree)(unsafe.Pointer(bp)) = *(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(stack)).Fcontents + uintptr(v8)*8))
		*(*MutableSubtree)(unsafe.Pointer(bp + 24)) = MutableSubtree{}
		if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
			v2 = libc.UintptrFromInt32(0)
		} else {
			v2 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
		}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 24)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(v2)))
		if int32(*(*uint8)(unsafe.Pointer(bp + 24 + 0))&0x1>>0) != 0 {
			v3 = libc.UintptrFromInt32(0)
		} else {
			v3 = *(*uintptr)(unsafe.Pointer(bp + 24)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24)))).Fchild_count)*8
		}
		grandchild1 = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(v3 + uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24)))).Fchild_count-uint32(1))*8)))
		ts_subtree_summarize_children(tls, grandchild1, language)
		ts_subtree_summarize_children(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 24)), language)
		ts_subtree_summarize_children(tls, *(*MutableSubtree)(unsafe.Pointer(bp)), language)
	}
}

func ts_subtree_summarize_children(tls *libc.TLS, _self MutableSubtree, language uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*MutableSubtree)(unsafe.Pointer(bp)) = _self
	var alias_sequence, children, v1 uintptr
	var child_lookahead_end_byte, grandchild_count, i, lookahead_end_byte, structural_index uint32_t
	var first_child, last_child Subtree
	var _ /* child at bp+8 */ Subtree
	_, _, _, _, _, _, _, _, _, _ = alias_sequence, child_lookahead_end_byte, children, first_child, grandchild_count, i, last_child, lookahead_end_byte, structural_index, v1
	_ = libc.Uint64FromInt64(4)
	{
		if !!(int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0) {
			libc.X__assert_fail(tls, __ccgo_ts+10383, __ccgo_ts+10199, int32(343), uintptr(unsafe.Pointer(&__func__111)))
		}
	}
	(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count = uint32(0)
	(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count = uint32(0)
	(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Ferror_cost = uint32(0)
	(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Frepeat_depth = uint16(0)
	(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_descendant_count = uint32(0)
	libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.BoolUint8(0 != 0), 6, 0x40)
	libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+45, libc.BoolUint8(0 != 0), 0, 0x1)
	libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.BoolUint8(0 != 0), 7, 0x80)
	(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fdynamic_precedence = 0
	structural_index = uint32(0)
	alias_sequence = ts_language_alias_sequence(tls, language, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id))
	lookahead_end_byte = uint32(0)
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
	}
	children = v1
	i = uint32(0)
	for {
		if !(i < (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count) {
			break
		}
		*(*Subtree)(unsafe.Pointer(bp + 8)) = Subtree{}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 8)) = *(*Subtree)(unsafe.Pointer(children + uintptr(i)*8))
		if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize.Fextent.Frow == uint32(0) && ts_subtree_depends_on_column(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+45, libc.BoolUint8(1 != 0), 0, 0x1)
		}
		if ts_subtree_has_external_scanner_state_change(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.BoolUint8(1 != 0), 7, 0x80)
		}
		if i == uint32(0) {
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fpadding = ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize = ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
		} else {
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize = length_add(tls, (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize, ts_subtree_total_size(tls, *(*Subtree)(unsafe.Pointer(bp + 8))))
		}
		child_lookahead_end_byte = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fpadding.Fbytes + (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize.Fbytes + ts_subtree_lookahead_bytes(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
		if child_lookahead_end_byte > lookahead_end_byte {
			lookahead_end_byte = child_lookahead_end_byte
		}
		if libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) != libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
			*(*uint32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 32)) += ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
		}
		grandchild_count = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
		if libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) || libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
			if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0) && !(ts_subtree_is_error(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 && grandchild_count == uint32(0)) {
				if ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
					*(*uint32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 32)) += uint32(100)
				} else {
					if grandchild_count > uint32(0) {
						*(*uint32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 32)) += uint32(100) * (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count
					}
				}
			}
		}
		*(*int32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 60)) += ts_subtree_dynamic_precedence(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
		*(*uint32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 56)) += ts_subtree_visible_descendant_count(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
		if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0) && libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) != 0 && alias_sequence != 0 && libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(alias_sequence + uintptr(structural_index)*2))) != 0 {
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_descendant_count = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_descendant_count + 1
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count + 1
			if ts_language_symbol_metadata(tls, language, *(*TSSymbol)(unsafe.Pointer(alias_sequence + uintptr(structural_index)*2))).Fnamed != 0 {
				(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count + 1
			}
		} else {
			if ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
				(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_descendant_count = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_descendant_count + 1
				(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count + 1
				if ts_subtree_named(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
					(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count + 1
				}
			} else {
				if grandchild_count > uint32(0) {
					*(*uint32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 48)) += (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).F__ccgo19_48.F__ccgo0_0.Fvisible_child_count
					*(*uint32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 52)) += (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count
				}
			}
		}
		if ts_subtree_has_external_tokens(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.BoolUint8(1 != 0), 6, 0x40)
		}
		if ts_subtree_is_error(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0 {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.AssignBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.BoolUint8(1 != 0), 1, 4, 0x10), 3, 0x8)
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fparse_state = libc.Uint16FromInt32(libc.Int32FromInt32(0x7fff)*libc.Int32FromInt32(2) + libc.Int32FromInt32(1))
		}
		if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) != 0) {
			structural_index = structural_index + 1
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Flookahead_bytes = lookahead_end_byte - (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize.Fbytes - (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fpadding.Fbytes
	if libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) || libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1) {
		*(*uint32_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 32)) += uint32(500) + uint32(1)*(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize.Fbytes + uint32(30)*(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize.Fextent.Frow
	}
	if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count > uint32(0) {
		first_child = *(*Subtree)(unsafe.Pointer(children))
		last_child = *(*Subtree)(unsafe.Pointer(children + uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count-uint32(1))*8))
		(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Ffirst_leaf.Fsymbol = ts_subtree_leaf_symbol(tls, first_child)
		(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Ffirst_leaf.Fparse_state = ts_subtree_leaf_parse_state(tls, first_child)
		if ts_subtree_fragile_left(tls, first_child) != 0 {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.BoolUint8(1 != 0), 3, 0x8)
		}
		if ts_subtree_fragile_right(tls, last_child) != 0 {
			libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, libc.BoolUint8(1 != 0), 4, 0x10)
		}
		if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count >= uint32(2) && !(int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x1>>0) != 0) && !(int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x2>>1) != 0) && libc.Int32FromUint16(ts_subtree_symbol(tls, first_child)) == libc.Int32FromUint16((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsymbol) {
			if ts_subtree_repeat_depth(tls, first_child) > ts_subtree_repeat_depth(tls, last_child) {
				(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Frepeat_depth = uint16(ts_subtree_repeat_depth(tls, first_child) + uint32(1))
			} else {
				(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Frepeat_depth = uint16(ts_subtree_repeat_depth(tls, last_child) + uint32(1))
			}
		}
	}
}

var __func__111 = [30]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 's', 'u', 'm', 'm', 'a', 'r', 'i', 'z', 'e', '_', 'c', 'h', 'i', 'l', 'd', 'r', 'e', 'n'}

func ts_subtree_new_node(tls *libc.TLS, symbol TSSymbol, children uintptr, production_id uint32, language uintptr) (r MutableSubtree) {
	var data uintptr
	var fragile uint8
	var metadata TSSymbolMetadata
	var new_byte_size size_t
	var result MutableSubtree
	_, _, _, _, _ = data, fragile, metadata, new_byte_size, result
	metadata = ts_language_symbol_metadata(tls, language, symbol)
	fragile = libc.BoolUint8(libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1))) || libc.Int32FromUint16(symbol) == libc.Int32FromUint16(libc.Uint16FromInt32(-libc.Int32FromInt32(1)))-libc.Int32FromInt32(1))
	new_byte_size = ts_subtree_alloc_size(tls, (*SubtreeArray)(unsafe.Pointer(children)).Fsize)
	if uint64((*SubtreeArray)(unsafe.Pointer(children)).Fcapacity)*uint64(8) < new_byte_size {
		(*SubtreeArray)(unsafe.Pointer(children)).Fcontents = (*(*func(*libc.TLS, uintptr, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_realloc})))(tls, (*SubtreeArray)(unsafe.Pointer(children)).Fcontents, new_byte_size)
		(*SubtreeArray)(unsafe.Pointer(children)).Fcapacity = uint32(new_byte_size / libc.Uint64FromInt64(8))
	}
	data = (*SubtreeArray)(unsafe.Pointer(children)).Fcontents + uintptr((*SubtreeArray)(unsafe.Pointer(children)).Fsize)*8
	*(*SubtreeHeapData)(unsafe.Pointer(data)) = SubtreeHeapData{
		Fref_count:   uint32(1),
		Fchild_count: (*SubtreeArray)(unsafe.Pointer(children)).Fsize,
		Fsymbol:      symbol,
		F__ccgo44:    metadata.Fvisible&0x1<<0 | metadata.Fnamed&0x1<<1 | fragile&0x1<<3 | fragile&0x1<<4 | libc.Uint8FromInt32(0)&0x1<<5 | libc.Uint8FromInt32(0)&0x1<<7,
		F__ccgo45:    libc.Uint8FromInt32(0) & 0x1 << 2,
		F__ccgo19_48: *(*struct {
			Fexternal_scanner_state [0]ExternalScannerState
			Flookahead_char         [0]int32_t
			F__ccgo0_0              struct {
				Fvisible_child_count      uint32_t
				Fnamed_child_count        uint32_t
				Fvisible_descendant_count uint32_t
				Fdynamic_precedence       int32_t
				Frepeat_depth             uint16_t
				Fproduction_id            uint16_t
				Ffirst_leaf               struct {
					Fsymbol      TSSymbol
					Fparse_state TSStateId
				}
			}
			F__ccgo_pad3 [8]byte
		})(unsafe.Pointer(&struct {
			f struct {
				Fvisible_child_count      uint32_t
				Fnamed_child_count        uint32_t
				Fvisible_descendant_count uint32_t
				Fdynamic_precedence       int32_t
				Frepeat_depth             uint16_t
				Fproduction_id            uint16_t
				Ffirst_leaf               struct {
					Fsymbol      TSSymbol
					Fparse_state TSStateId
				}
			}
			_ [8]byte
		}{f: struct {
			Fvisible_child_count      uint32_t
			Fnamed_child_count        uint32_t
			Fvisible_descendant_count uint32_t
			Fdynamic_precedence       int32_t
			Frepeat_depth             uint16_t
			Fproduction_id            uint16_t
			Ffirst_leaf               struct {
				Fsymbol      TSSymbol
				Fparse_state TSStateId
			}
		}{
			Fproduction_id: uint16(production_id),
		}})),
	}
	result = *(*MutableSubtree)(unsafe.Pointer(&struct{ f uintptr }{f: data}))
	ts_subtree_summarize_children(tls, result, language)
	return result
}

func ts_subtree_new_error_node(tls *libc.TLS, children uintptr, extra uint8, language uintptr) (r Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* result at bp+0 */ MutableSubtree
	*(*MutableSubtree)(unsafe.Pointer(bp)) = MutableSubtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = ts_subtree_new_node(tls, libc.Uint16FromInt32(-libc.Int32FromInt32(1)), children, uint32(0), language)
	libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+44, extra, 2, 0x4)
	return ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp)))
}

func ts_subtree_new_missing_leaf(tls *libc.TLS, pool uintptr, symbol TSSymbol, padding Length, lookahead_bytes uint32_t, language uintptr) (r Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* result at bp+0 */ Subtree
	*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
	*(*struct {
		Fptr  [0]uintptr
		Fdata SubtreeInlineData
	})(unsafe.Pointer(bp)) = ts_subtree_new_leaf(tls, pool, symbol, padding, length_zero(tls), lookahead_bytes, uint16(0), libc.BoolUint8(0 != 0), libc.BoolUint8(0 != 0), libc.BoolUint8(0 != 0), language)
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		libc.SetBitFieldPtr8Uint8(bp+0, libc.BoolUint8(1 != 0), 5, 0x20)
	} else {
		libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(bp))+45, libc.BoolUint8(1 != 0), 1, 0x2)
	}
	return *(*Subtree)(unsafe.Pointer(bp))
}

func ts_subtree_retain(tls *libc.TLS, _self Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		return
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !(libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp))) > libc.Uint32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+10404, __ccgo_ts+10199, int32(560), uintptr(unsafe.Pointer(&__func__112)))
		}
	}
	atomic_inc(tls, *(*uintptr)(unsafe.Pointer(bp)))
	_ = libc.Uint64FromInt64(4)
	{
		if !(libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp))) != libc.Uint32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+10428, __ccgo_ts+10199, int32(562), uintptr(unsafe.Pointer(&__func__112)))
		}
	}
}

var __func__112 = [18]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'r', 'e', 't', 'a', 'i', 'n'}

func ts_subtree_release(tls *libc.TLS, pool uintptr, _self Subtree) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var children, v2, v4 uintptr
	var i, v1 uint32_t
	var _ /* child at bp+16 */ Subtree
	var _ /* tree at bp+8 */ MutableSubtree
	_, _, _, _, _ = children, i, v1, v2, v4
	if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
		return
	}
	(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fsize = uint32(0)
	_ = libc.Uint64FromInt64(4)
	{
		if !(libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp))) > libc.Uint32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+10404, __ccgo_ts+10199, int32(569), uintptr(unsafe.Pointer(&__func__113)))
		}
	}
	if atomic_dec(tls, *(*uintptr)(unsafe.Pointer(bp))) == uint32(0) {
		(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fsize, pool+16+12, uint32(1), libc.Uint64FromInt64(8))
		v2 = pool + 16 + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(bp)))
	}
	for (*SubtreePool)(unsafe.Pointer(pool)).Ftree_stack.Fsize > uint32(0) {
		*(*MutableSubtree)(unsafe.Pointer(bp + 8)) = MutableSubtree{}
		v2 = pool + 16 + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 8)) = *(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8))
		if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count > uint32(0) {
			if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 {
				v4 = libc.UintptrFromInt32(0)
			} else {
				v4 = *(*uintptr)(unsafe.Pointer(bp + 8)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count)*8
			}
			children = v4
			i = uint32(0)
			for {
				if !(i < (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count) {
					break
				}
				*(*Subtree)(unsafe.Pointer(bp + 16)) = Subtree{}
				*(*struct {
					Fptr  [0]uintptr
					Fdata SubtreeInlineData
				})(unsafe.Pointer(bp + 16)) = *(*Subtree)(unsafe.Pointer(children + uintptr(i)*8))
				if int32(*(*uint8)(unsafe.Pointer(bp + 16 + 0))&0x1>>0) != 0 {
					goto _6
				}
				_ = libc.Uint64FromInt64(4)
				{
					if !(libc.AtomicLoadPUint32(*(*uintptr)(unsafe.Pointer(bp + 16))) > libc.Uint32FromInt32(0)) {
						libc.X__assert_fail(tls, __ccgo_ts+10453, __ccgo_ts+10199, int32(581), uintptr(unsafe.Pointer(&__func__113)))
					}
				}
				if atomic_dec(tls, *(*uintptr)(unsafe.Pointer(bp + 16))) == uint32(0) {
					(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fsize, pool+16+12, uint32(1), libc.Uint64FromInt64(8))
					v2 = pool + 16 + 8
					v1 = *(*uint32_t)(unsafe.Pointer(v2))
					*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
					*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(bp + 16)))
				}
				goto _6
			_6:
				;
				i = i + 1
			}
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, children)
		} else {
			if int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)) + 44))&0x40>>6) != 0 {
				ts_external_scanner_state_delete(tls, *(*uintptr)(unsafe.Pointer(bp + 8))+48)
			}
			ts_subtree_pool_free(tls, pool, *(*uintptr)(unsafe.Pointer(bp + 8)))
		}
	}
}

var __func__113 = [19]int8{'t', 's', '_', 's', 'u', 'b', 't', 'r', 'e', 'e', '_', 'r', 'e', 'l', 'e', 'a', 's', 'e'}

func ts_subtree_compare(tls *libc.TLS, _left Subtree, _right Subtree, pool uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _left
	*(*Subtree)(unsafe.Pointer(bp + 8)) = _right
	var i, v1 uint32_t
	var left_child, right_child Subtree
	var result int32
	var v2, v4 uintptr
	_, _, _, _, _, _, _ = i, left_child, result, right_child, v1, v2, v4
	(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fsize, pool+16+12, uint32(1), libc.Uint64FromInt64(8))
	v2 = pool + 16 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(bp)))
	(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fsize, pool+16+12, uint32(1), libc.Uint64FromInt64(8))
	v2 = pool + 16 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))
	for (*SubtreePool)(unsafe.Pointer(pool)).Ftree_stack.Fsize > uint32(0) {
		v2 = pool + 16 + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*Subtree)(unsafe.Pointer(bp + 8)) = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)))
		v2 = pool + 16 + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*Subtree)(unsafe.Pointer(bp)) = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)))
		result = 0
		if libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))) < libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) {
			result = -int32(1)
		} else {
			if libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp + 8)))) < libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp)))) {
				result = int32(1)
			} else {
				if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) < ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) {
					result = -int32(1)
				} else {
					if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp + 8))) < ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) {
						result = int32(1)
					}
				}
			}
		}
		if result != 0 {
			(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fsize = uint32(0)
			return result
		}
		i = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp)))
		for {
			if !(i > uint32(0)) {
				break
			}
			if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
				v2 = libc.UintptrFromInt32(0)
			} else {
				v2 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
			}
			left_child = *(*Subtree)(unsafe.Pointer(v2 + uintptr(i-uint32(1))*8))
			if int32(*(*uint8)(unsafe.Pointer(bp + 8 + 0))&0x1>>0) != 0 {
				v4 = libc.UintptrFromInt32(0)
			} else {
				v4 = *(*uintptr)(unsafe.Pointer(bp + 8)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8)))).Fchild_count)*8
			}
			right_child = *(*Subtree)(unsafe.Pointer(v4 + uintptr(i-uint32(1))*8))
			(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fsize, pool+16+12, uint32(1), libc.Uint64FromInt64(8))
			v2 = pool + 16 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, left_child)
			(*MutableSubtreeArray)(unsafe.Pointer(pool + 16)).Fcontents = _array__grow(tls, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents, (*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fsize, pool+16+12, uint32(1), libc.Uint64FromInt64(8))
			v2 = pool + 16 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*MutableSubtree)(unsafe.Pointer((*MutableSubtreeArray)(unsafe.Pointer(pool+16)).Fcontents + uintptr(v1)*8)) = ts_subtree_to_mut_unsafe(tls, right_child)
			goto _9
		_9:
			;
			i = i - 1
		}
	}
	return 0
}

func ts_subtree_set_has_changes(tls *libc.TLS, self uintptr) {
	if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
		libc.SetBitFieldPtr8Uint8(self+0, libc.BoolUint8(1 != 0), 4, 0x10)
	} else {
		libc.SetBitFieldPtr8Uint8(*(*uintptr)(unsafe.Pointer(self))+44, libc.BoolUint8(1 != 0), 5, 0x20)
	}
}

func ts_subtree_edit(tls *libc.TLS, _self Subtree, input_edit uintptr, pool uintptr) (r Subtree) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	var child, data, v2 uintptr
	var child_edit, edit Edit
	var child_left, child_right, child_size, padding, size, total_size Length
	var column_shifted, is_noop, is_pure_insertion, parent_depends_on_column uint8
	var end_byte, i, lookahead_bytes, n, v1 uint32_t
	var entry struct {
		Ftree uintptr
		Fedit Edit
	}
	var _ /* result at bp+24 */ MutableSubtree
	var _ /* stack at bp+8 */ struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = child, child_edit, child_left, child_right, child_size, column_shifted, data, edit, end_byte, entry, i, is_noop, is_pure_insertion, lookahead_bytes, n, padding, parent_depends_on_column, size, total_size, v1, v2
	*(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)) = struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	}{}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+8)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+8)).Fsize, bp+8+12, uint32(1), libc.Uint64FromInt64(48))
	v2 = bp + 8 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*struct {
		Ftree uintptr
		Fedit Edit
	})(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+8)).Fcontents + uintptr(v1)*48)) = struct {
		Ftree uintptr
		Fedit Edit
	}{
		Ftree: bp,
		Fedit: Edit{
			Fstart: Length{
				Fbytes:  (*TSInputEdit)(unsafe.Pointer(input_edit)).Fstart_byte,
				Fextent: (*TSInputEdit)(unsafe.Pointer(input_edit)).Fstart_point,
			},
			Fold_end: Length{
				Fbytes:  (*TSInputEdit)(unsafe.Pointer(input_edit)).Fold_end_byte,
				Fextent: (*TSInputEdit)(unsafe.Pointer(input_edit)).Fold_end_point,
			},
			Fnew_end: Length{
				Fbytes:  (*TSInputEdit)(unsafe.Pointer(input_edit)).Fnew_end_byte,
				Fextent: (*TSInputEdit)(unsafe.Pointer(input_edit)).Fnew_end_point,
			},
		},
	}
	for (*(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8))).Fsize != 0 {
		v2 = bp + 8 + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		entry = *(*struct {
			Ftree uintptr
			Fedit Edit
		})(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+8)).Fcontents + uintptr(v1)*48))
		edit = entry.Fedit
		is_noop = libc.BoolUint8(edit.Fold_end.Fbytes == edit.Fstart.Fbytes && edit.Fnew_end.Fbytes == edit.Fstart.Fbytes)
		is_pure_insertion = libc.BoolUint8(edit.Fold_end.Fbytes == edit.Fstart.Fbytes)
		parent_depends_on_column = ts_subtree_depends_on_column(tls, *(*Subtree)(unsafe.Pointer(entry.Ftree)))
		column_shifted = libc.BoolUint8(edit.Fnew_end.Fextent.Fcolumn != edit.Fold_end.Fextent.Fcolumn)
		size = ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(entry.Ftree)))
		padding = ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(entry.Ftree)))
		total_size = length_add(tls, padding, size)
		lookahead_bytes = ts_subtree_lookahead_bytes(tls, *(*Subtree)(unsafe.Pointer(entry.Ftree)))
		end_byte = total_size.Fbytes + lookahead_bytes
		if edit.Fstart.Fbytes > end_byte || is_noop != 0 && edit.Fstart.Fbytes == end_byte {
			continue
		}
		if edit.Fold_end.Fbytes <= padding.Fbytes {
			padding = length_add(tls, edit.Fnew_end, length_sub(tls, padding, edit.Fold_end))
		} else {
			if edit.Fstart.Fbytes < padding.Fbytes {
				size = length_saturating_sub(tls, size, length_sub(tls, edit.Fold_end, padding))
				padding = edit.Fnew_end
			} else {
				if edit.Fstart.Fbytes < total_size.Fbytes || edit.Fstart.Fbytes == total_size.Fbytes && is_pure_insertion != 0 {
					size = length_add(tls, length_sub(tls, edit.Fnew_end, padding), length_saturating_sub(tls, total_size, edit.Fold_end))
				}
			}
		}
		*(*MutableSubtree)(unsafe.Pointer(bp + 24)) = MutableSubtree{}
		*(*struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		})(unsafe.Pointer(bp + 24)) = ts_subtree_make_mut(tls, pool, *(*Subtree)(unsafe.Pointer(entry.Ftree)))
		if int32(*(*uint8)(unsafe.Pointer(bp + 24 + 0))&0x1>>0) != 0 {
			if ts_subtree_can_inline(tls, padding, size, lookahead_bytes) != 0 {
				(*(*SubtreeInlineData)(unsafe.Pointer(bp + 24))).Fpadding_bytes = uint8(padding.Fbytes)
				libc.SetBitFieldPtr8Uint8(bp+24+5, uint8(padding.Fextent.Frow), 0, 0xf)
				(*(*SubtreeInlineData)(unsafe.Pointer(bp + 24))).Fpadding_columns = uint8(padding.Fextent.Fcolumn)
				(*(*SubtreeInlineData)(unsafe.Pointer(bp + 24))).Fsize_bytes = uint8(size.Fbytes)
			} else {
				data = ts_subtree_pool_allocate(tls, pool)
				libc.AtomicStorePUint32(data, libc.Uint32FromInt32(1))
				(*SubtreeHeapData)(unsafe.Pointer(data)).Fpadding = padding
				(*SubtreeHeapData)(unsafe.Pointer(data)).Fsize = size
				(*SubtreeHeapData)(unsafe.Pointer(data)).Flookahead_bytes = lookahead_bytes
				(*SubtreeHeapData)(unsafe.Pointer(data)).Ferror_cost = uint32(0)
				(*SubtreeHeapData)(unsafe.Pointer(data)).Fchild_count = uint32(0)
				(*SubtreeHeapData)(unsafe.Pointer(data)).Fsymbol = uint16((*(*SubtreeInlineData)(unsafe.Pointer(bp + 24))).Fsymbol)
				(*SubtreeHeapData)(unsafe.Pointer(data)).Fparse_state = (*(*SubtreeInlineData)(unsafe.Pointer(bp + 24))).Fparse_state
				libc.SetBitFieldPtr8Uint8(data+44, libc.Uint8FromInt32(libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(bp + 24 + 0))&0x2>>1) != 0)), 0, 0x1)
				libc.SetBitFieldPtr8Uint8(data+44, libc.Uint8FromInt32(libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(bp + 24 + 0))&0x4>>2) != 0)), 1, 0x2)
				libc.SetBitFieldPtr8Uint8(data+44, libc.Uint8FromInt32(libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(bp + 24 + 0))&0x8>>3) != 0)), 2, 0x4)
				libc.SetBitFieldPtr8Uint8(data+44, libc.BoolUint8(0 != 0), 3, 0x8)
				libc.SetBitFieldPtr8Uint8(data+44, libc.BoolUint8(0 != 0), 4, 0x10)
				libc.SetBitFieldPtr8Uint8(data+44, libc.BoolUint8(0 != 0), 5, 0x20)
				libc.SetBitFieldPtr8Uint8(data+44, libc.BoolUint8(0 != 0), 6, 0x40)
				libc.SetBitFieldPtr8Uint8(data+45, libc.BoolUint8(0 != 0), 0, 0x1)
				libc.SetBitFieldPtr8Uint8(data+45, libc.Uint8FromInt32(libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(bp + 24 + 0))&0x20>>5) != 0)), 1, 0x2)
				libc.SetBitFieldPtr8Uint8(data+45, libc.Uint8FromInt32(libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(bp + 24 + 0))&0x40>>6) != 0)), 2, 0x4)
				*(*uintptr)(unsafe.Pointer(bp + 24)) = data
			}
		} else {
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24)))).Fpadding = padding
			(*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24)))).Fsize = size
		}
		ts_subtree_set_has_changes(tls, bp+24)
		*(*Subtree)(unsafe.Pointer(entry.Ftree)) = ts_subtree_from_mut(tls, *(*MutableSubtree)(unsafe.Pointer(bp + 24)))
		child_right = length_zero(tls)
		i = uint32(0)
		n = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(entry.Ftree)))
		for {
			if !(i < n) {
				break
			}
			if int32(*(*uint8)(unsafe.Pointer(entry.Ftree + 0))&0x1>>0) != 0 {
				v2 = libc.UintptrFromInt32(0)
			} else {
				v2 = *(*uintptr)(unsafe.Pointer(entry.Ftree)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(entry.Ftree)))).Fchild_count)*8
			}
			child = v2 + uintptr(i)*8
			child_size = ts_subtree_total_size(tls, *(*Subtree)(unsafe.Pointer(child)))
			child_left = child_right
			child_right = length_add(tls, child_left, child_size)
			if child_right.Fbytes+ts_subtree_lookahead_bytes(tls, *(*Subtree)(unsafe.Pointer(child))) < edit.Fstart.Fbytes {
				goto _5
			}
			if (child_left.Fbytes > edit.Fold_end.Fbytes || child_left.Fbytes == edit.Fold_end.Fbytes && child_size.Fbytes > uint32(0) && i > uint32(0)) && (!(parent_depends_on_column != 0) || child_left.Fextent.Frow > padding.Fextent.Frow) && (!(ts_subtree_depends_on_column(tls, *(*Subtree)(unsafe.Pointer(child))) != 0) || !(column_shifted != 0) || child_left.Fextent.Frow > edit.Fold_end.Fextent.Frow) {
				break
			}
			child_edit = Edit{
				Fstart:   length_saturating_sub(tls, edit.Fstart, child_left),
				Fold_end: length_saturating_sub(tls, edit.Fold_end, child_left),
				Fnew_end: length_saturating_sub(tls, edit.Fnew_end, child_left),
			}
			if child_right.Fbytes > edit.Fstart.Fbytes || child_right.Fbytes == edit.Fstart.Fbytes && is_pure_insertion != 0 {
				edit.Fnew_end = edit.Fstart
			} else {
				child_edit.Fold_end = child_edit.Fstart
				child_edit.Fnew_end = child_edit.Fstart
			}
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp + 8)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+8)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+8)).Fsize, bp+8+12, uint32(1), libc.Uint64FromInt64(48))
			v2 = bp + 8 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*struct {
				Ftree uintptr
				Fedit Edit
			})(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(bp+8)).Fcontents + uintptr(v1)*48)) = struct {
				Ftree uintptr
				Fedit Edit
			}{
				Ftree: child,
				Fedit: child_edit,
			}
			goto _5
		_5:
			;
			i = i + 1
		}
	}
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+8)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+8)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)).Fcapacity = uint32(0)
	return *(*Subtree)(unsafe.Pointer(bp))
}

func ts_subtree_last_external_token(tls *libc.TLS, _tree Subtree) (r Subtree) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _tree
	var child Subtree
	var i uint32_t
	var v2 uintptr
	_, _, _ = child, i, v2
	if !(ts_subtree_has_external_tokens(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0) {
		return Subtree{}
	}
	for (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count > uint32(0) {
		i = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count - uint32(1)
		for {
			if !(i+uint32(1) > uint32(0)) {
				break
			}
			if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
				v2 = libc.UintptrFromInt32(0)
			} else {
				v2 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
			}
			child = *(*Subtree)(unsafe.Pointer(v2 + uintptr(i)*8))
			if ts_subtree_has_external_tokens(tls, child) != 0 {
				*(*Subtree)(unsafe.Pointer(bp)) = child
				break
			}
			goto _1
		_1:
			;
			i = i - 1
		}
	}
	return *(*Subtree)(unsafe.Pointer(bp))
}

func ts_subtree__write_char_to_string(tls *libc.TLS, str uintptr, n size_t, chr int32_t) (r size_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	if chr == -int32(1) {
		return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, str, n, __ccgo_ts+10478, 0))
	} else {
		if chr == int32('\000') {
			return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, str, n, __ccgo_ts+10486, 0))
		} else {
			if chr == int32('\n') {
				return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, str, n, __ccgo_ts+10491, 0))
			} else {
				if chr == int32('\t') {
					return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, str, n, __ccgo_ts+10496, 0))
				} else {
					if chr == int32('\r') {
						return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, str, n, __ccgo_ts+10501, 0))
					} else {
						if 0 < chr && chr < int32(128) && libc.Int32FromUint16(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(chr)*2)))&libc.Int32FromUint16(uint16(_ISprint)) != 0 {
							return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, str, n, __ccgo_ts+10506, libc.VaList(bp+8, chr)))
						} else {
							return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, str, n, __ccgo_ts+10511, libc.VaList(bp+8, chr)))
						}
					}
				}
			}
		}
	}
	return r
}

var ROOT_FIELD = __ccgo_ts + 10514

func ts_subtree__write_to_string(tls *libc.TLS, _self Subtree, _string uintptr, limit size_t, language uintptr, include_all uint8, alias_symbol TSSymbol, alias_is_named uint8, field_name uintptr) (r size_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	*(*uintptr)(unsafe.Pointer(bp + 8)) = _string
	var alias_sequence, child_field_name, map1, symbol_name, symbol_name1, writer, v1 uintptr
	var child Subtree
	var i, structural_child_index uint32_t
	var is_root, is_visible, subtree_alias_is_named uint8
	var subtree_alias_symbol, symbol, symbol1 TSSymbol
	var v2, v4 int32
	var v3 bool
	var _ /* cursor at bp+16 */ uintptr
	var _ /* field_map at bp+24 */ uintptr
	var _ /* field_map_end at bp+32 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alias_sequence, child, child_field_name, i, is_root, is_visible, map1, structural_child_index, subtree_alias_is_named, subtree_alias_symbol, symbol, symbol1, symbol_name, symbol_name1, writer, v1, v2, v3, v4
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) {
		return libc.Uint64FromInt32(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(bp + 8)), limit, __ccgo_ts+10523, 0))
	}
	*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 8))
	if limit > uint64(1) {
		v1 = bp + 16
	} else {
		v1 = bp + 8
	}
	writer = v1
	is_root = libc.BoolUint8(field_name == ROOT_FIELD)
	if v3 = include_all != 0 || ts_subtree_missing(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0; !v3 {
		if alias_symbol != 0 {
			v2 = libc.Int32FromUint8(alias_is_named)
		} else {
			v2 = libc.BoolInt32(ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 && ts_subtree_named(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0)
		}
	}
	is_visible = libc.BoolUint8(v3 || v2 != 0)
	if is_visible != 0 {
		if !(is_root != 0) {
			*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10530, 0))
			if field_name != 0 {
				*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10532, libc.VaList(bp+48, field_name)))
			}
		}
		if ts_subtree_is_error(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 && ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) == uint32(0) && (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fsize.Fbytes > uint32(0) {
			*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10537, 0))
			*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(ts_subtree__write_char_to_string(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, *(*int32_t)(unsafe.Add(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))), 48))))
		} else {
			if alias_symbol != 0 {
				v2 = libc.Int32FromUint16(alias_symbol)
			} else {
				v2 = libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))
			}
			symbol = libc.Uint16FromInt32(v2)
			symbol_name = ts_language_symbol_name(tls, language, symbol)
			if ts_subtree_missing(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
				*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10550, 0))
				if alias_is_named != 0 || ts_subtree_named(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
					*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10560, libc.VaList(bp+48, symbol_name)))
				} else {
					*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10563, libc.VaList(bp+48, symbol_name)))
				}
			} else {
				*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10568, libc.VaList(bp+48, symbol_name)))
			}
		}
	} else {
		if is_root != 0 {
			if alias_symbol != 0 {
				v2 = libc.Int32FromUint16(alias_symbol)
			} else {
				v2 = libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))
			}
			symbol1 = libc.Uint16FromInt32(v2)
			symbol_name1 = ts_language_symbol_name(tls, language, symbol1)
			if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
				*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10568, libc.VaList(bp+48, symbol_name1)))
			} else {
				if ts_subtree_named(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
					*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10572, libc.VaList(bp+48, symbol_name1)))
				} else {
					*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10577, libc.VaList(bp+48, symbol_name1)))
				}
			}
		}
	}
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0 {
		alias_sequence = ts_language_alias_sequence(tls, language, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id))
		ts_language_field_map(tls, language, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), bp+24, bp+32)
		structural_child_index = uint32(0)
		i = uint32(0)
		for {
			if !(i < (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count) {
				break
			}
			if int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0 {
				v1 = libc.UintptrFromInt32(0)
			} else {
				v1 = *(*uintptr)(unsafe.Pointer(bp)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count)*8
			}
			child = *(*Subtree)(unsafe.Pointer(v1 + uintptr(i)*8))
			if ts_subtree_extra(tls, child) != 0 {
				*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(ts_subtree__write_to_string(tls, child, *(*uintptr)(unsafe.Pointer(writer)), limit, language, include_all, uint16(0), libc.BoolUint8(0 != 0), libc.UintptrFromInt32(0)))
			} else {
				if alias_sequence != 0 {
					v2 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(alias_sequence + uintptr(structural_child_index)*2)))
				} else {
					v2 = 0
				}
				subtree_alias_symbol = libc.Uint16FromInt32(v2)
				if subtree_alias_symbol != 0 {
					v4 = libc.Int32FromUint8(ts_language_symbol_metadata(tls, language, subtree_alias_symbol).Fnamed)
				} else {
					v4 = 0
				}
				subtree_alias_is_named = libc.Uint8FromInt32(libc.BoolInt32(v4 != 0))
				if is_visible != 0 {
					v1 = libc.UintptrFromInt32(0)
				} else {
					v1 = field_name
				}
				child_field_name = v1
				map1 = *(*uintptr)(unsafe.Pointer(bp + 24))
				for {
					if !(map1 < *(*uintptr)(unsafe.Pointer(bp + 32))) {
						break
					}
					if !((*TSFieldMapEntry)(unsafe.Pointer(map1)).Finherited != 0) && uint32((*TSFieldMapEntry)(unsafe.Pointer(map1)).Fchild_index) == structural_child_index {
						child_field_name = *(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(language)).Ffield_names + uintptr((*TSFieldMapEntry)(unsafe.Pointer(map1)).Ffield_id)*8))
						break
					}
					goto _11
				_11:
					;
					map1 += 4
				}
				*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(ts_subtree__write_to_string(tls, child, *(*uintptr)(unsafe.Pointer(writer)), limit, language, include_all, subtree_alias_symbol, subtree_alias_is_named, child_field_name))
				structural_child_index = structural_child_index + 1
			}
			goto _6
		_6:
			;
			i = i + 1
		}
	}
	if is_visible != 0 {
		*(*uintptr)(unsafe.Pointer(bp + 16)) = *(*uintptr)(unsafe.Pointer(bp + 16)) + uintptr(libc.X__builtin_snprintf(tls, *(*uintptr)(unsafe.Pointer(writer)), limit, __ccgo_ts+10584, 0))
	}
	return libc.Uint64FromInt64(int64(*(*uintptr)(unsafe.Pointer(bp + 16))) - int64(*(*uintptr)(unsafe.Pointer(bp + 8))))
}

func ts_subtree_string(tls *libc.TLS, self Subtree, alias_symbol TSSymbol, alias_is_named uint8, language uintptr, include_all uint8) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var result uintptr
	var size size_t
	var _ /* scratch_string at bp+0 */ [1]int8
	_, _ = result, size
	size = ts_subtree__write_to_string(tls, self, bp, uint64(1), language, include_all, alias_symbol, alias_is_named, ROOT_FIELD) + uint64(1)
	result = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, size*uint64(1))
	ts_subtree__write_to_string(tls, self, result, size, language, include_all, alias_symbol, alias_is_named, ROOT_FIELD)
	return result
}

func ts_subtree__print_dot_graph(tls *libc.TLS, self uintptr, start_offset uint32_t, language uintptr, alias_symbol TSSymbol, f uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var child, v3 uintptr
	var child_info_offset, child_start_offset, end_offset, i, n uint32_t
	var subtree_alias_symbol, subtree_symbol, symbol TSSymbol
	var v1 int32
	_, _, _, _, _, _, _, _, _, _, _ = child, child_info_offset, child_start_offset, end_offset, i, n, subtree_alias_symbol, subtree_symbol, symbol, v1, v3
	subtree_symbol = ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(self)))
	if alias_symbol != 0 {
		v1 = libc.Int32FromUint16(alias_symbol)
	} else {
		v1 = libc.Int32FromUint16(subtree_symbol)
	}
	symbol = libc.Uint16FromInt32(v1)
	end_offset = start_offset + ts_subtree_total_bytes(tls, *(*Subtree)(unsafe.Pointer(self)))
	libc.Xfprintf(tls, f, __ccgo_ts+10586, libc.VaList(bp+8, self))
	ts_language_write_symbol_as_dot_string(tls, language, f, symbol)
	libc.Xfprintf(tls, f, __ccgo_ts+10108, 0)
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(self))) == uint32(0) {
		libc.Xfprintf(tls, f, __ccgo_ts+10603, 0)
	}
	if ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(self))) != 0 {
		libc.Xfprintf(tls, f, __ccgo_ts+10621, 0)
	}
	if ts_subtree_has_changes(tls, *(*Subtree)(unsafe.Pointer(self))) != 0 {
		libc.Xfprintf(tls, f, __ccgo_ts+10638, 0)
	}
	libc.Xfprintf(tls, f, __ccgo_ts+10664, libc.VaList(bp+8, start_offset, end_offset, libc.Int32FromUint16(ts_subtree_parse_state(tls, *(*Subtree)(unsafe.Pointer(self)))), ts_subtree_error_cost(tls, *(*Subtree)(unsafe.Pointer(self))), libc.Int32FromUint8(ts_subtree_has_changes(tls, *(*Subtree)(unsafe.Pointer(self)))), libc.Int32FromUint8(ts_subtree_depends_on_column(tls, *(*Subtree)(unsafe.Pointer(self)))), ts_subtree_visible_descendant_count(tls, *(*Subtree)(unsafe.Pointer(self))), ts_subtree_repeat_depth(tls, *(*Subtree)(unsafe.Pointer(self))), ts_subtree_lookahead_bytes(tls, *(*Subtree)(unsafe.Pointer(self)))))
	if ts_subtree_is_error(tls, *(*Subtree)(unsafe.Pointer(self))) != 0 && ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(self))) == uint32(0) && *(*int32_t)(unsafe.Add(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self))), 48)) != 0 {
		libc.Xfprintf(tls, f, __ccgo_ts+10811, libc.VaList(bp+8, *(*int32_t)(unsafe.Add(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self))), 48))))
	}
	libc.Xfprintf(tls, f, __ccgo_ts+9851, 0)
	child_start_offset = start_offset
	child_info_offset = libc.Uint32FromInt32(libc.Int32FromUint16((*TSLanguage)(unsafe.Pointer(language)).Fmax_alias_sequence_length) * libc.Int32FromUint16(ts_subtree_production_id(tls, *(*Subtree)(unsafe.Pointer(self)))))
	i = uint32(0)
	n = ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer(self)))
	for {
		if !(i < n) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
			v3 = libc.UintptrFromInt32(0)
		} else {
			v3 = *(*uintptr)(unsafe.Pointer(self)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count)*8
		}
		child = v3 + uintptr(i)*8
		subtree_alias_symbol = uint16(0)
		if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(child))) != 0) && child_info_offset != 0 {
			subtree_alias_symbol = *(*TSSymbol)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer(language)).Falias_sequences + uintptr(child_info_offset)*2))
			child_info_offset = child_info_offset + 1
		}
		ts_subtree__print_dot_graph(tls, child, child_start_offset, language, subtree_alias_symbol, f)
		libc.Xfprintf(tls, f, __ccgo_ts+10828, libc.VaList(bp+8, self, child, i))
		child_start_offset = child_start_offset + ts_subtree_total_bytes(tls, *(*Subtree)(unsafe.Pointer(child)))
		goto _2
	_2:
		;
		i = i + 1
	}
}

func ts_subtree_print_dot_graph(tls *libc.TLS, _self Subtree, language uintptr, f uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	libc.Xfprintf(tls, f, __ccgo_ts+10861, 0)
	libc.Xfprintf(tls, f, __ccgo_ts+9588, 0)
	ts_subtree__print_dot_graph(tls, bp, uint32(0), language, uint16(0), f)
	libc.Xfprintf(tls, f, __ccgo_ts+10167, 0)
}

func ts_subtree_external_scanner_state(tls *libc.TLS, _self Subtree) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*Subtree)(unsafe.Pointer(bp)) = _self
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 && !(int32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0) != 0) && int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 44))&0x40>>6) != 0 && (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fchild_count == uint32(0) {
		return *(*uintptr)(unsafe.Pointer(bp)) + 48
	} else {
		return uintptr(unsafe.Pointer(&empty_state))
	}
	return r
}

var empty_state = ExternalScannerState{}

func ts_subtree_external_scanner_state_eq(tls *libc.TLS, self Subtree, other Subtree) (r uint8) {
	var state_other, state_self uintptr
	_, _ = state_other, state_self
	state_self = ts_subtree_external_scanner_state(tls, self)
	state_other = ts_subtree_external_scanner_state(tls, other)
	return ts_external_scanner_state_eq(tls, state_self, ts_external_scanner_state_data(tls, state_other), (*ExternalScannerState)(unsafe.Pointer(state_other)).Flength)
}

type CursorChildIterator = struct {
	Fparent                 Subtree
	Ftree                   uintptr
	Fposition               Length
	Fchild_index            uint32_t
	Fstructural_child_index uint32_t
	Fdescendant_index       uint32_t
	Falias_sequence         uintptr
}

func ts_tree_cursor_is_entry_visible(tls *libc.TLS, self uintptr, index uint32_t) (r uint8) {
	var entry, parent_entry uintptr
	_, _ = entry, parent_entry
	_ = libc.Uint64FromInt64(4)
	{
		if !(index < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+10877, __ccgo_ts+10918, int32(19), uintptr(unsafe.Pointer(&__func__114)))
		}
	}
	entry = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr(index)*32
	if index == uint32(0) || ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))) != 0 {
		return libc.BoolUint8(1 != 0)
	} else {
		if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))) != 0) {
			_ = libc.Uint64FromInt64(4)
			{
				if !(index-libc.Uint32FromInt32(1) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+10966, __ccgo_ts+10918, int32(23), uintptr(unsafe.Pointer(&__func__114)))
				}
			}
			parent_entry = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents + uintptr(index-uint32(1))*32
			return uint8(libc.BoolUint16(ts_language_alias_at(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), (*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index) != 0))
		} else {
			return libc.BoolUint8(0 != 0)
		}
	}
	return r
}

var __func__114 = [32]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'i', 's', '_', 'e', 'n', 't', 'r', 'y', '_', 'v', 'i', 's', 'i', 'b', 'l', 'e'}

func ts_tree_cursor_iterate_children(tls *libc.TLS, self uintptr) (r CursorChildIterator) {
	var alias_sequence, last_entry uintptr
	var descendant_index uint32_t
	_, _, _ = alias_sequence, descendant_index, last_entry
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+10918, int32(35), uintptr(unsafe.Pointer(&__func__115)))
		}
	}
	last_entry = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32
	if ts_subtree_child_count(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fsubtree))) == uint32(0) {
		return CursorChildIterator{
			Fparent:   Subtree{},
			Ftree:     (*TreeCursor)(unsafe.Pointer(self)).Ftree,
			Fposition: length_zero(tls),
		}
	}
	alias_sequence = ts_language_alias_sequence(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fsubtree)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id))
	descendant_index = (*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fdescendant_index
	if ts_tree_cursor_is_entry_visible(tls, self, (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize-uint32(1)) != 0 {
		descendant_index = descendant_index + uint32(1)
	}
	return CursorChildIterator{
		Fparent:           *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fsubtree)),
		Ftree:             (*TreeCursor)(unsafe.Pointer(self)).Ftree,
		Fposition:         (*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fposition,
		Fdescendant_index: descendant_index,
		Falias_sequence:   alias_sequence,
	}
}

var __func__115 = [32]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'i', 't', 'e', 'r', 'a', 't', 'e', '_', 'c', 'h', 'i', 'l', 'd', 'r', 'e', 'n'}

func ts_tree_cursor_child_iterator_next(tls *libc.TLS, self uintptr, result uintptr, visible uintptr) (r uint8) {
	var child, v1 uintptr
	var extra uint8
	var next_child Subtree
	_, _, _, _ = child, extra, next_child, v1
	if !(*(*uintptr)(unsafe.Pointer(self)) != 0) || (*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index == (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count {
		return libc.BoolUint8(0 != 0)
	}
	if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = *(*uintptr)(unsafe.Pointer(self)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count)*8
	}
	child = v1 + uintptr((*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index)*8
	*(*TreeCursorEntry)(unsafe.Pointer(result)) = TreeCursorEntry{
		Fsubtree:                child,
		Fposition:               (*CursorChildIterator)(unsafe.Pointer(self)).Fposition,
		Fchild_index:            (*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index,
		Fstructural_child_index: (*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index,
		Fdescendant_index:       (*CursorChildIterator)(unsafe.Pointer(self)).Fdescendant_index,
	}
	*(*uint8)(unsafe.Pointer(visible)) = ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer(child)))
	extra = ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(child)))
	if !(extra != 0) {
		if (*CursorChildIterator)(unsafe.Pointer(self)).Falias_sequence != 0 {
			v1 = visible
			*(*uint8)(unsafe.Pointer(v1)) = uint8(int32(*(*uint8)(unsafe.Pointer(v1))) | libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*CursorChildIterator)(unsafe.Pointer(self)).Falias_sequence + uintptr((*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index)*2))))
		}
		(*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index = (*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index + 1
	}
	*(*uint32_t)(unsafe.Pointer(self + 36)) += ts_subtree_visible_descendant_count(tls, *(*Subtree)(unsafe.Pointer(child)))
	if *(*uint8)(unsafe.Pointer(visible)) != 0 {
		*(*uint32_t)(unsafe.Pointer(self + 36)) += uint32(1)
	}
	(*CursorChildIterator)(unsafe.Pointer(self)).Fposition = length_add(tls, (*CursorChildIterator)(unsafe.Pointer(self)).Fposition, ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(child))))
	(*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index = (*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index + 1
	if (*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index < (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count {
		if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
			v1 = libc.UintptrFromInt32(0)
		} else {
			v1 = *(*uintptr)(unsafe.Pointer(self)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count)*8
		}
		next_child = *(*Subtree)(unsafe.Pointer(v1 + uintptr((*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index)*8))
		(*CursorChildIterator)(unsafe.Pointer(self)).Fposition = length_add(tls, (*CursorChildIterator)(unsafe.Pointer(self)).Fposition, ts_subtree_padding(tls, next_child))
	}
	return libc.BoolUint8(1 != 0)
}

func length_backtrack(tls *libc.TLS, a Length, b Length) (r Length) {
	var result Length
	_ = result
	if length_is_undefined(tls, a) != 0 || b.Fextent.Frow != uint32(0) {
		return LENGTH_UNDEFINED
	}
	result.Fbytes = a.Fbytes - b.Fbytes
	result.Fextent.Frow = a.Fextent.Frow
	result.Fextent.Fcolumn = a.Fextent.Fcolumn - b.Fextent.Fcolumn
	return result
}

func ts_tree_cursor_child_iterator_previous(tls *libc.TLS, self uintptr, result uintptr, visible uintptr) (r uint8) {
	var child, v1 uintptr
	var extra uint8
	var previous_child Subtree
	var size Length
	_, _, _, _, _ = child, extra, previous_child, size, v1
	if !(*(*uintptr)(unsafe.Pointer(self)) != 0) || int32(libc.Int8FromUint32((*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index)) == -int32(1) {
		return libc.BoolUint8(0 != 0)
	}
	if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = *(*uintptr)(unsafe.Pointer(self)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count)*8
	}
	child = v1 + uintptr((*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index)*8
	*(*TreeCursorEntry)(unsafe.Pointer(result)) = TreeCursorEntry{
		Fsubtree:                child,
		Fposition:               (*CursorChildIterator)(unsafe.Pointer(self)).Fposition,
		Fchild_index:            (*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index,
		Fstructural_child_index: (*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index,
	}
	*(*uint8)(unsafe.Pointer(visible)) = ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer(child)))
	extra = ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(child)))
	(*CursorChildIterator)(unsafe.Pointer(self)).Fposition = length_backtrack(tls, (*CursorChildIterator)(unsafe.Pointer(self)).Fposition, ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(child))))
	(*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index = (*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index - 1
	if !(extra != 0) && (*CursorChildIterator)(unsafe.Pointer(self)).Falias_sequence != 0 {
		v1 = visible
		*(*uint8)(unsafe.Pointer(v1)) = uint8(int32(*(*uint8)(unsafe.Pointer(v1))) | libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer((*CursorChildIterator)(unsafe.Pointer(self)).Falias_sequence + uintptr((*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index)*2))))
		if (*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index > uint32(0) {
			(*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index = (*CursorChildIterator)(unsafe.Pointer(self)).Fstructural_child_index - 1
		}
	}
	if (*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index < (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count {
		if int32(*(*uint8)(unsafe.Pointer(self + 0))&0x1>>0) != 0 {
			v1 = libc.UintptrFromInt32(0)
		} else {
			v1 = *(*uintptr)(unsafe.Pointer(self)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(self)))).Fchild_count)*8
		}
		previous_child = *(*Subtree)(unsafe.Pointer(v1 + uintptr((*CursorChildIterator)(unsafe.Pointer(self)).Fchild_index)*8))
		size = ts_subtree_size(tls, previous_child)
		(*CursorChildIterator)(unsafe.Pointer(self)).Fposition = length_backtrack(tls, (*CursorChildIterator)(unsafe.Pointer(self)).Fposition, size)
	}
	return libc.BoolUint8(1 != 0)
}

func ts_tree_cursor_new(tls *libc.TLS, node TSNode) (r TSTreeCursor) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* self at bp+0 */ TSTreeCursor
	*(*TSTreeCursor)(unsafe.Pointer(bp)) = TSTreeCursor{}
	ts_tree_cursor_init(tls, bp, node)
	return *(*TSTreeCursor)(unsafe.Pointer(bp))
}

func ts_tree_cursor_reset(tls *libc.TLS, _self uintptr, node TSNode) {
	ts_tree_cursor_init(tls, _self, node)
}

func ts_tree_cursor_init(tls *libc.TLS, self uintptr, _node TSNode) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TSNode)(unsafe.Pointer(bp)) = _node
	var v1 uint32_t
	var v2 uintptr
	_, _ = v1, v2
	(*TreeCursor)(unsafe.Pointer(self)).Ftree = (*(*TSNode)(unsafe.Pointer(bp))).Ftree
	(*TreeCursor)(unsafe.Pointer(self)).Froot_alias_symbol = uint16(*(*uint32_t)(unsafe.Pointer(bp + 3*4)))
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
	v2 = self + 8 + 8
	v1 = *(*uint32_t)(unsafe.Pointer(v2))
	*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
	*(*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = TreeCursorEntry{
		Fsubtree: (*(*TSNode)(unsafe.Pointer(bp))).Fid,
		Fposition: Length{
			Fbytes:  ts_node_start_byte(tls, *(*TSNode)(unsafe.Pointer(bp))),
			Fextent: ts_node_start_point(tls, *(*TSNode)(unsafe.Pointer(bp))),
		},
	}
}

func ts_tree_cursor_delete(tls *libc.TLS, _self uintptr) {
	var self uintptr
	_ = self
	self = _self
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 8)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self + 8)).Fcapacity = uint32(0)
}

func ts_tree_cursor_goto_first_child_internal(tls *libc.TLS, _self uintptr) (r TreeCursorStep) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var self, v2 uintptr
	var v1 uint32_t
	var _ /* entry at bp+8 */ TreeCursorEntry
	var _ /* iterator at bp+40 */ CursorChildIterator
	var _ /* visible at bp+0 */ uint8
	_, _, _ = self, v1, v2
	self = _self
	*(*CursorChildIterator)(unsafe.Pointer(bp + 40)) = CursorChildIterator{}
	*(*struct {
		Fparent struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		}
		Ftree     uintptr
		Fposition struct {
			Fbytes  uint32
			Fextent TSPoint
		}
		Fchild_index            uint32
		Fstructural_child_index uint32
		Fdescendant_index       uint32
		Falias_sequence         uintptr
	})(unsafe.Pointer(bp + 40)) = ts_tree_cursor_iterate_children(tls, self)
	for ts_tree_cursor_child_iterator_next(tls, bp+40, bp+8, bp) != 0 {
		if *(*uint8)(unsafe.Pointer(bp)) != 0 {
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
			v2 = self + 8 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*TreeCursorEntry)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = *(*TreeCursorEntry)(unsafe.Pointer(bp + 8))
			return int32(TreeCursorStepVisible)
		}
		if ts_subtree_visible_child_count(tls, *(*Subtree)(unsafe.Pointer((*(*TreeCursorEntry)(unsafe.Pointer(bp + 8))).Fsubtree))) > uint32(0) {
			(*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents, (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
			v2 = self + 8 + 8
			v1 = *(*uint32_t)(unsafe.Pointer(v2))
			*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
			*(*TreeCursorEntry)(unsafe.Pointer((*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = *(*TreeCursorEntry)(unsafe.Pointer(bp + 8))
			return int32(TreeCursorStepHidden)
		}
	}
	return int32(TreeCursorStepNone)
}

func ts_tree_cursor_goto_first_child(tls *libc.TLS, self uintptr) (r uint8) {
	for {
		switch ts_tree_cursor_goto_first_child_internal(tls, self) {
		case int32(TreeCursorStepHidden):
			goto _1
		case int32(TreeCursorStepVisible):
			return libc.BoolUint8(1 != 0)
		default:
			return libc.BoolUint8(0 != 0)
		}
		goto _1
	_1:
	}
	return r
}

func ts_tree_cursor_goto_last_child_internal(tls *libc.TLS, _self uintptr) (r TreeCursorStep) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var last_entry TreeCursorEntry
	var last_step TreeCursorStep
	var self, v2 uintptr
	var v1 uint32_t
	var _ /* entry at bp+8 */ TreeCursorEntry
	var _ /* iterator at bp+40 */ CursorChildIterator
	var _ /* visible at bp+0 */ uint8
	_, _, _, _, _ = last_entry, last_step, self, v1, v2
	self = _self
	*(*CursorChildIterator)(unsafe.Pointer(bp + 40)) = CursorChildIterator{}
	*(*struct {
		Fparent struct {
			Fptr  [0]uintptr
			Fdata SubtreeInlineData
		}
		Ftree     uintptr
		Fposition struct {
			Fbytes  uint32
			Fextent TSPoint
		}
		Fchild_index            uint32
		Fstructural_child_index uint32
		Fdescendant_index       uint32
		Falias_sequence         uintptr
	})(unsafe.Pointer(bp + 40)) = ts_tree_cursor_iterate_children(tls, self)
	if !(*(*uintptr)(unsafe.Pointer(bp + 40)) != 0) || (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)))).Fchild_count == uint32(0) {
		return int32(TreeCursorStepNone)
	}
	last_entry = TreeCursorEntry{}
	last_step = int32(TreeCursorStepNone)
	for ts_tree_cursor_child_iterator_next(tls, bp+40, bp+8, bp) != 0 {
		if *(*uint8)(unsafe.Pointer(bp)) != 0 {
			last_entry = *(*TreeCursorEntry)(unsafe.Pointer(bp + 8))
			last_step = int32(TreeCursorStepVisible)
		} else {
			if ts_subtree_visible_child_count(tls, *(*Subtree)(unsafe.Pointer((*(*TreeCursorEntry)(unsafe.Pointer(bp + 8))).Fsubtree))) > uint32(0) {
				last_entry = *(*TreeCursorEntry)(unsafe.Pointer(bp + 8))
				last_step = int32(TreeCursorStepHidden)
			}
		}
	}
	if last_entry.Fsubtree != 0 {
		(*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
		v2 = self + 8 + 8
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
		*(*TreeCursorEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = last_entry
		return last_step
	}
	return int32(TreeCursorStepNone)
}

func ts_tree_cursor_goto_last_child(tls *libc.TLS, self uintptr) (r uint8) {
	for {
		switch ts_tree_cursor_goto_last_child_internal(tls, self) {
		case int32(TreeCursorStepHidden):
			goto _1
		case int32(TreeCursorStepVisible):
			return libc.BoolUint8(1 != 0)
		default:
			return libc.BoolUint8(0 != 0)
		}
		goto _1
	_1:
	}
	return r
}

func ts_tree_cursor_goto_first_child_for_byte_and_point(tls *libc.TLS, _self uintptr, goal_byte uint32_t, goal_point TSPoint) (r int64_t) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var at_goal, did_descend uint8
	var entry_end Length
	var initial_size, visible_child_count, visible_child_index, v1 uint32_t
	var self, v2 uintptr
	var _ /* entry at bp+8 */ TreeCursorEntry
	var _ /* iterator at bp+40 */ CursorChildIterator
	var _ /* visible at bp+0 */ uint8
	_, _, _, _, _, _, _, _, _ = at_goal, did_descend, entry_end, initial_size, self, visible_child_count, visible_child_index, v1, v2
	self = _self
	initial_size = (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize
	visible_child_index = uint32(0)
	for cond := true; cond; cond = did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		*(*CursorChildIterator)(unsafe.Pointer(bp + 40)) = CursorChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Fdescendant_index       uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 40)) = ts_tree_cursor_iterate_children(tls, self)
		for ts_tree_cursor_child_iterator_next(tls, bp+40, bp+8, bp) != 0 {
			entry_end = length_add(tls, (*(*TreeCursorEntry)(unsafe.Pointer(bp + 8))).Fposition, ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer((*(*TreeCursorEntry)(unsafe.Pointer(bp + 8))).Fsubtree))))
			at_goal = libc.BoolUint8(entry_end.Fbytes > goal_byte && point_gt(tls, entry_end.Fextent, goal_point) != 0)
			visible_child_count = ts_subtree_visible_child_count(tls, *(*Subtree)(unsafe.Pointer((*(*TreeCursorEntry)(unsafe.Pointer(bp + 8))).Fsubtree)))
			if at_goal != 0 {
				if *(*uint8)(unsafe.Pointer(bp)) != 0 {
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+8)).Fcontents, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
					v2 = self + 8 + 8
					v1 = *(*uint32_t)(unsafe.Pointer(v2))
					*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
					*(*TreeCursorEntry)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = *(*TreeCursorEntry)(unsafe.Pointer(bp + 8))
					return libc.Int64FromUint32(visible_child_index)
				}
				if visible_child_count > uint32(0) {
					(*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+8)).Fcontents, (*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
					v2 = self + 8 + 8
					v1 = *(*uint32_t)(unsafe.Pointer(v2))
					*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
					*(*TreeCursorEntry)(unsafe.Pointer((*struct {
						Fcontents uintptr
						Fsize     uint32_t
						Fcapacity uint32_t
					})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = *(*TreeCursorEntry)(unsafe.Pointer(bp + 8))
					did_descend = libc.BoolUint8(1 != 0)
					break
				}
			} else {
				if *(*uint8)(unsafe.Pointer(bp)) != 0 {
					visible_child_index = visible_child_index + 1
				} else {
					visible_child_index = visible_child_index + visible_child_count
				}
			}
		}
	}
	(*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize = initial_size
	return int64(-int32(1))
}

func ts_tree_cursor_goto_first_child_for_byte(tls *libc.TLS, self uintptr, goal_byte uint32_t) (r int64_t) {
	return ts_tree_cursor_goto_first_child_for_byte_and_point(tls, self, goal_byte, TSPoint{})
}

func ts_tree_cursor_goto_first_child_for_point(tls *libc.TLS, self uintptr, goal_point TSPoint) (r int64_t) {
	return ts_tree_cursor_goto_first_child_for_byte_and_point(tls, self, uint32(0), goal_point)
}

type __ccgo_fp__Xts_tree_cursor_goto_sibling_internal_1 = func(*libc.TLS, uintptr, uintptr, uintptr) uint8

func ts_tree_cursor_goto_sibling_internal(tls *libc.TLS, _self uintptr, __ccgo_fp_advance uintptr) (r TreeCursorStep) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var initial_size, v1 uint32_t
	var self, v2 uintptr
	var _ /* entry at bp+0 */ TreeCursorEntry
	var _ /* iterator at bp+32 */ CursorChildIterator
	var _ /* visible at bp+80 */ uint8
	_, _, _, _ = initial_size, self, v1, v2
	self = _self
	initial_size = (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize
	for (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize > uint32(1) {
		v2 = self + 8 + 8
		*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) - 1
		v1 = *(*uint32_t)(unsafe.Pointer(v2))
		*(*TreeCursorEntry)(unsafe.Pointer(bp)) = *(*TreeCursorEntry)(unsafe.Pointer((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32))
		*(*CursorChildIterator)(unsafe.Pointer(bp + 32)) = CursorChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Fdescendant_index       uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 32)) = ts_tree_cursor_iterate_children(tls, self)
		(*(*CursorChildIterator)(unsafe.Pointer(bp + 32))).Fchild_index = (*(*TreeCursorEntry)(unsafe.Pointer(bp))).Fchild_index
		(*(*CursorChildIterator)(unsafe.Pointer(bp + 32))).Fstructural_child_index = (*(*TreeCursorEntry)(unsafe.Pointer(bp))).Fstructural_child_index
		(*(*CursorChildIterator)(unsafe.Pointer(bp + 32))).Fposition = (*(*TreeCursorEntry)(unsafe.Pointer(bp))).Fposition
		(*(*CursorChildIterator)(unsafe.Pointer(bp + 32))).Fdescendant_index = (*(*TreeCursorEntry)(unsafe.Pointer(bp))).Fdescendant_index
		*(*uint8)(unsafe.Pointer(bp + 80)) = libc.BoolUint8(0 != 0)
		(*(*func(*libc.TLS, uintptr, uintptr, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_advance})))(tls, bp+32, bp, bp+80)
		if *(*uint8)(unsafe.Pointer(bp + 80)) != 0 && (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize+uint32(1) < initial_size {
			break
		}
		for (*(*func(*libc.TLS, uintptr, uintptr, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_advance})))(tls, bp+32, bp, bp+80) != 0 {
			if *(*uint8)(unsafe.Pointer(bp + 80)) != 0 {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
				v2 = self + 8 + 8
				v1 = *(*uint32_t)(unsafe.Pointer(v2))
				*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
				*(*TreeCursorEntry)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = *(*TreeCursorEntry)(unsafe.Pointer(bp))
				return int32(TreeCursorStepVisible)
			}
			if ts_subtree_visible_child_count(tls, *(*Subtree)(unsafe.Pointer((*(*TreeCursorEntry)(unsafe.Pointer(bp))).Fsubtree))) != 0 {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
				v2 = self + 8 + 8
				v1 = *(*uint32_t)(unsafe.Pointer(v2))
				*(*uint32_t)(unsafe.Pointer(v2)) = *(*uint32_t)(unsafe.Pointer(v2)) + 1
				*(*TreeCursorEntry)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents + uintptr(v1)*32)) = *(*TreeCursorEntry)(unsafe.Pointer(bp))
				return int32(TreeCursorStepHidden)
			}
		}
	}
	(*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize = initial_size
	return int32(TreeCursorStepNone)
}

func ts_tree_cursor_goto_next_sibling_internal(tls *libc.TLS, _self uintptr) (r TreeCursorStep) {
	return ts_tree_cursor_goto_sibling_internal(tls, _self, __ccgo_fp(ts_tree_cursor_child_iterator_next))
}

func ts_tree_cursor_goto_next_sibling(tls *libc.TLS, self uintptr) (r uint8) {
	switch ts_tree_cursor_goto_next_sibling_internal(tls, self) {
	case int32(TreeCursorStepHidden):
		ts_tree_cursor_goto_first_child(tls, self)
		return libc.BoolUint8(1 != 0)
	case int32(TreeCursorStepVisible):
		return libc.BoolUint8(1 != 0)
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

func ts_tree_cursor_goto_previous_sibling_internal(tls *libc.TLS, _self uintptr) (r TreeCursorStep) {
	var child_index, i uint32_t
	var children, parent, self, v1 uintptr
	var position Length
	var step TreeCursorStep
	_, _, _, _, _, _, _, _ = child_index, children, i, parent, position, self, step, v1
	self = _self
	step = ts_tree_cursor_goto_sibling_internal(tls, _self, __ccgo_fp(ts_tree_cursor_child_iterator_previous))
	if step == int32(TreeCursorStepNone) {
		return step
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+10918, int32(373), uintptr(unsafe.Pointer(&__func__116)))
		}
	}
	if !(length_is_undefined(tls, (*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents+uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32)).Fposition) != 0) {
		return step
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !((*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize-libc.Uint32FromInt32(2) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+11011, __ccgo_ts+10918, int32(377), uintptr(unsafe.Pointer(&__func__116)))
		}
	}
	parent = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize-uint32(2))*32
	position = (*TreeCursorEntry)(unsafe.Pointer(parent)).Fposition
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+10918, int32(379), uintptr(unsafe.Pointer(&__func__116)))
		}
	}
	child_index = (*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32)).Fchild_index
	if int32(*(*uint8)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent)).Fsubtree + 0))&0x1>>0) != 0 {
		v1 = libc.UintptrFromInt32(0)
	} else {
		v1 = *(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent)).Fsubtree)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent)).Fsubtree)))).Fchild_count)*8
	}
	children = v1
	if child_index > uint32(0) {
		position = length_add(tls, position, ts_subtree_size(tls, *(*Subtree)(unsafe.Pointer(children))))
		i = uint32(1)
		for {
			if !(i < child_index) {
				break
			}
			position = length_add(tls, position, ts_subtree_total_size(tls, *(*Subtree)(unsafe.Pointer(children + uintptr(i)*8))))
			goto _2
		_2:
			;
			i = i + 1
		}
		position = length_add(tls, position, ts_subtree_padding(tls, *(*Subtree)(unsafe.Pointer(children + uintptr(child_index)*8))))
	}
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+10918, int32(391), uintptr(unsafe.Pointer(&__func__116)))
		}
	}
	(*TreeCursorEntry)(unsafe.Pointer((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32)).Fposition = position
	return step
}

var __func__116 = [46]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'g', 'o', 't', 'o', '_', 'p', 'r', 'e', 'v', 'i', 'o', 'u', 's', '_', 's', 'i', 'b', 'l', 'i', 'n', 'g', '_', 'i', 'n', 't', 'e', 'r', 'n', 'a', 'l'}

func ts_tree_cursor_goto_previous_sibling(tls *libc.TLS, self uintptr) (r uint8) {
	switch ts_tree_cursor_goto_previous_sibling_internal(tls, self) {
	case int32(TreeCursorStepHidden):
		ts_tree_cursor_goto_last_child(tls, self)
		return libc.BoolUint8(1 != 0)
	case int32(TreeCursorStepVisible):
		return libc.BoolUint8(1 != 0)
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

func ts_tree_cursor_goto_parent(tls *libc.TLS, _self uintptr) (r uint8) {
	var i uint32
	var self uintptr
	_, _ = i, self
	self = _self
	i = (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize - uint32(2)
	for {
		if !(i+uint32(1) > uint32(0)) {
			break
		}
		if ts_tree_cursor_is_entry_visible(tls, self, i) != 0 {
			(*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize = i + uint32(1)
			return libc.BoolUint8(1 != 0)
		}
		goto _1
	_1:
		;
		i = i - 1
	}
	return libc.BoolUint8(0 != 0)
}

func ts_tree_cursor_goto_descendant(tls *libc.TLS, _self uintptr, goal_descendant_index uint32_t) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var did_descend uint8
	var entry, self, v4 uintptr
	var i, next_descendant_index, v3 uint32_t
	var v2 int32
	var _ /* entry at bp+8 */ TreeCursorEntry
	var _ /* iterator at bp+40 */ CursorChildIterator
	var _ /* visible at bp+0 */ uint8
	_, _, _, _, _, _, _, _ = did_descend, entry, i, next_descendant_index, self, v2, v3, v4
	self = _self
	for {
		i = (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize - uint32(1)
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+11067, __ccgo_ts+10918, int32(428), uintptr(unsafe.Pointer(&__func__117)))
			}
		}
		entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(i)*32
		if ts_tree_cursor_is_entry_visible(tls, self, i) != 0 {
			v2 = int32(1)
		} else {
			v2 = 0
		}
		next_descendant_index = (*TreeCursorEntry)(unsafe.Pointer(entry)).Fdescendant_index + libc.Uint32FromInt32(v2) + ts_subtree_visible_descendant_count(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree)))
		if (*TreeCursorEntry)(unsafe.Pointer(entry)).Fdescendant_index <= goal_descendant_index && next_descendant_index > goal_descendant_index {
			break
		} else {
			if (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize <= uint32(1) {
				return
			} else {
				(*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize = (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize - 1
			}
		}
		goto _1
	_1:
	}
	did_descend = libc.BoolUint8(1 != 0)
	for cond := true; cond; cond = did_descend != 0 {
		did_descend = libc.BoolUint8(0 != 0)
		*(*CursorChildIterator)(unsafe.Pointer(bp + 40)) = CursorChildIterator{}
		*(*struct {
			Fparent struct {
				Fptr  [0]uintptr
				Fdata SubtreeInlineData
			}
			Ftree     uintptr
			Fposition struct {
				Fbytes  uint32
				Fextent TSPoint
			}
			Fchild_index            uint32
			Fstructural_child_index uint32
			Fdescendant_index       uint32
			Falias_sequence         uintptr
		})(unsafe.Pointer(bp + 40)) = ts_tree_cursor_iterate_children(tls, self)
		if (*(*CursorChildIterator)(unsafe.Pointer(bp + 40))).Fdescendant_index > goal_descendant_index {
			return
		}
		for ts_tree_cursor_child_iterator_next(tls, bp+40, bp+8, bp) != 0 {
			if (*(*CursorChildIterator)(unsafe.Pointer(bp + 40))).Fdescendant_index > goal_descendant_index {
				(*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self + 8)).Fcontents = _array__grow(tls, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents, (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fsize, self+8+12, uint32(1), libc.Uint64FromInt64(32))
				v4 = self + 8 + 8
				v3 = *(*uint32_t)(unsafe.Pointer(v4))
				*(*uint32_t)(unsafe.Pointer(v4)) = *(*uint32_t)(unsafe.Pointer(v4)) + 1
				*(*TreeCursorEntry)(unsafe.Pointer((*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fcontents + uintptr(v3)*32)) = *(*TreeCursorEntry)(unsafe.Pointer(bp + 8))
				if *(*uint8)(unsafe.Pointer(bp)) != 0 && (*(*TreeCursorEntry)(unsafe.Pointer(bp + 8))).Fdescendant_index == goal_descendant_index {
					return
				} else {
					did_descend = libc.BoolUint8(1 != 0)
					break
				}
			}
		}
	}
}

var __func__117 = [31]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'g', 'o', 't', 'o', '_', 'd', 'e', 's', 'c', 'e', 'n', 'd', 'a', 'n', 't'}

func ts_tree_cursor_current_descendant_index(tls *libc.TLS, _self uintptr) (r uint32_t) {
	var last_entry, self uintptr
	_, _ = last_entry, self
	self = _self
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+10918, int32(472), uintptr(unsafe.Pointer(&__func__118)))
		}
	}
	last_entry = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32
	return (*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fdescendant_index
}

var __func__118 = [40]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'c', 'u', 'r', 'r', 'e', 'n', 't', '_', 'd', 'e', 's', 'c', 'e', 'n', 'd', 'a', 'n', 't', '_', 'i', 'n', 'd', 'e', 'x'}

func ts_tree_cursor_current_node(tls *libc.TLS, _self uintptr) (r TSNode) {
	var alias_symbol TSSymbol
	var is_extra uint8
	var last_entry, parent_entry, self uintptr
	var v1 int32
	_, _, _, _, _, _ = alias_symbol, is_extra, last_entry, parent_entry, self, v1
	self = _self
	_ = libc.Uint64FromInt64(4)
	{
		if !((*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize-libc.Uint32FromInt32(1) < (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fsize) {
			libc.X__assert_fail(tls, __ccgo_ts+165, __ccgo_ts+10918, int32(478), uintptr(unsafe.Pointer(&__func__119)))
		}
	}
	last_entry = (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fcontents + uintptr((*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(self+8)).Fsize-uint32(1))*32
	is_extra = ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fsubtree)))
	if is_extra != 0 {
		v1 = 0
	} else {
		v1 = libc.Int32FromUint16((*TreeCursor)(unsafe.Pointer(self)).Froot_alias_symbol)
	}
	alias_symbol = libc.Uint16FromInt32(v1)
	if (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize > uint32(1) && !(is_extra != 0) {
		_ = libc.Uint64FromInt64(4)
		{
			if !((*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize-libc.Uint32FromInt32(2) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+11011, __ccgo_ts+10918, int32(482), uintptr(unsafe.Pointer(&__func__119)))
			}
		}
		parent_entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr((*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize-uint32(2))*32
		alias_symbol = ts_language_alias_at(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), (*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fstructural_child_index)
	}
	return ts_node_new(tls, (*TreeCursor)(unsafe.Pointer(self)).Ftree, (*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fsubtree, (*TreeCursorEntry)(unsafe.Pointer(last_entry)).Fposition, alias_symbol)
}

var __func__119 = [28]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'c', 'u', 'r', 'r', 'e', 'n', 't', '_', 'n', 'o', 'd', 'e'}

func ts_tree_cursor_current_status(tls *libc.TLS, _self uintptr, field_id uintptr, has_later_siblings uintptr, has_later_named_siblings uintptr, can_have_later_siblings_with_this_field uintptr, supertypes uintptr, supertype_count uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var alias_sequence, entry, map1, map11, parent_entry, self, v4 uintptr
	var entry_metadata, sibling_metadata TSSymbolMetadata
	var entry_symbol TSSymbol
	var i, j, max_supertypes, sibling_count, structural_child_index uint32
	var v2 int32
	var _ /* field_map at bp+8 */ uintptr
	var _ /* field_map_end at bp+16 */ uintptr
	var _ /* sibling at bp+0 */ Subtree
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alias_sequence, entry, entry_metadata, entry_symbol, i, j, map1, map11, max_supertypes, parent_entry, self, sibling_count, sibling_metadata, structural_child_index, v2, v4
	self = _self
	max_supertypes = *(*uint32)(unsafe.Pointer(supertype_count))
	*(*TSFieldId)(unsafe.Pointer(field_id)) = uint16(0)
	*(*uint32)(unsafe.Pointer(supertype_count)) = uint32(0)
	*(*uint8)(unsafe.Pointer(has_later_siblings)) = libc.BoolUint8(0 != 0)
	*(*uint8)(unsafe.Pointer(has_later_named_siblings)) = libc.BoolUint8(0 != 0)
	*(*uint8)(unsafe.Pointer(can_have_later_siblings_with_this_field)) = libc.BoolUint8(0 != 0)
	i = (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize - uint32(1)
	for {
		if !(i > uint32(0)) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+11067, __ccgo_ts+10918, int32(519), uintptr(unsafe.Pointer(&__func__120)))
			}
		}
		entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(i)*32
		_ = libc.Uint64FromInt64(4)
		{
			if !(i-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+11104, __ccgo_ts+10918, int32(520), uintptr(unsafe.Pointer(&__func__120)))
			}
		}
		parent_entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(i-uint32(1))*32
		alias_sequence = ts_language_alias_sequence(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id))
		if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))) != 0) && alias_sequence != 0 && *(*TSSymbol)(unsafe.Pointer(alias_sequence + uintptr((*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index)*2)) != 0 {
			v2 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(alias_sequence + uintptr((*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index)*2)))
		} else {
			v2 = libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))))
		}
		entry_symbol = libc.Uint16FromInt32(v2)
		entry_metadata = ts_language_symbol_metadata(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, entry_symbol)
		if i != (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize-uint32(1) && entry_metadata.Fvisible != 0 {
			break
		}
		if entry_metadata.Fsupertype != 0 && *(*uint32)(unsafe.Pointer(supertype_count)) < max_supertypes {
			*(*TSSymbol)(unsafe.Pointer(supertypes + uintptr(*(*uint32)(unsafe.Pointer(supertype_count)))*2)) = entry_symbol
			*(*uint32)(unsafe.Pointer(supertype_count)) = *(*uint32)(unsafe.Pointer(supertype_count)) + 1
		}
		if !(*(*uint8)(unsafe.Pointer(has_later_siblings)) != 0) {
			sibling_count = (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).Fchild_count
			structural_child_index = (*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index
			if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))) != 0) {
				structural_child_index = structural_child_index + 1
			}
			j = (*TreeCursorEntry)(unsafe.Pointer(entry)).Fchild_index + uint32(1)
			for {
				if !(j < sibling_count) {
					break
				}
				*(*Subtree)(unsafe.Pointer(bp)) = Subtree{}
				if int32(*(*uint8)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree + 0))&0x1>>0) != 0 {
					v4 = libc.UintptrFromInt32(0)
				} else {
					v4 = *(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)) - uintptr((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).Fchild_count)*8
				}
				*(*struct {
					Fptr  [0]uintptr
					Fdata SubtreeInlineData
				})(unsafe.Pointer(bp)) = *(*Subtree)(unsafe.Pointer(v4 + uintptr(j)*8))
				if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0) && alias_sequence != 0 && *(*TSSymbol)(unsafe.Pointer(alias_sequence + uintptr(structural_child_index)*2)) != 0 {
					v2 = libc.Int32FromUint16(*(*TSSymbol)(unsafe.Pointer(alias_sequence + uintptr(structural_child_index)*2)))
				} else {
					v2 = libc.Int32FromUint16(ts_subtree_symbol(tls, *(*Subtree)(unsafe.Pointer(bp))))
				}
				sibling_metadata = ts_language_symbol_metadata(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, libc.Uint16FromInt32(v2))
				if sibling_metadata.Fvisible != 0 {
					*(*uint8)(unsafe.Pointer(has_later_siblings)) = libc.BoolUint8(1 != 0)
					if *(*uint8)(unsafe.Pointer(has_later_named_siblings)) != 0 {
						break
					}
					if sibling_metadata.Fnamed != 0 {
						*(*uint8)(unsafe.Pointer(has_later_named_siblings)) = libc.BoolUint8(1 != 0)
						break
					}
				} else {
					if ts_subtree_visible_child_count(tls, *(*Subtree)(unsafe.Pointer(bp))) > uint32(0) {
						*(*uint8)(unsafe.Pointer(has_later_siblings)) = libc.BoolUint8(1 != 0)
						if *(*uint8)(unsafe.Pointer(has_later_named_siblings)) != 0 {
							break
						}
						if (*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).F__ccgo19_48.F__ccgo0_0.Fnamed_child_count > uint32(0) {
							*(*uint8)(unsafe.Pointer(has_later_named_siblings)) = libc.BoolUint8(1 != 0)
							break
						}
					}
				}
				if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer(bp))) != 0) {
					structural_child_index = structural_child_index + 1
				}
				goto _3
			_3:
				;
				j = j + 1
			}
		}
		if !(ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))) != 0) {
			ts_language_field_map(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), bp+8, bp+16)
			if !(*(*TSFieldId)(unsafe.Pointer(field_id)) != 0) {
				map1 = *(*uintptr)(unsafe.Pointer(bp + 8))
				for {
					if !(map1 < *(*uintptr)(unsafe.Pointer(bp + 16))) {
						break
					}
					if !((*TSFieldMapEntry)(unsafe.Pointer(map1)).Finherited != 0) && uint32((*TSFieldMapEntry)(unsafe.Pointer(map1)).Fchild_index) == (*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index {
						*(*TSFieldId)(unsafe.Pointer(field_id)) = (*TSFieldMapEntry)(unsafe.Pointer(map1)).Ffield_id
						break
					}
					goto _6
				_6:
					;
					map1 += 4
				}
			}
			if *(*TSFieldId)(unsafe.Pointer(field_id)) != 0 {
				map11 = *(*uintptr)(unsafe.Pointer(bp + 8))
				for {
					if !(map11 < *(*uintptr)(unsafe.Pointer(bp + 16))) {
						break
					}
					if libc.Int32FromUint16((*TSFieldMapEntry)(unsafe.Pointer(map11)).Ffield_id) == libc.Int32FromUint16(*(*TSFieldId)(unsafe.Pointer(field_id))) && uint32((*TSFieldMapEntry)(unsafe.Pointer(map11)).Fchild_index) > (*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index {
						*(*uint8)(unsafe.Pointer(can_have_later_siblings_with_this_field)) = libc.BoolUint8(1 != 0)
						break
					}
					goto _7
				_7:
					;
					map11 += 4
				}
			}
		}
		goto _1
	_1:
		;
		i = i - 1
	}
}

var __func__120 = [30]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'c', 'u', 'r', 'r', 'e', 'n', 't', '_', 's', 't', 'a', 't', 'u', 's'}

func ts_tree_cursor_current_depth(tls *libc.TLS, _self uintptr) (r uint32_t) {
	var depth uint32_t
	var i uint32
	var self uintptr
	_, _, _ = depth, i, self
	self = _self
	depth = uint32(0)
	i = uint32(1)
	for {
		if !(i < (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize) {
			break
		}
		if ts_tree_cursor_is_entry_visible(tls, self, i) != 0 {
			depth = depth + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return depth
}

func ts_tree_cursor_parent_node(tls *libc.TLS, _self uintptr) (r TSNode) {
	var alias_symbol TSSymbol
	var entry, parent_entry, self uintptr
	var i int32
	var is_visible uint8
	_, _, _, _, _, _ = alias_symbol, entry, i, is_visible, parent_entry, self
	self = _self
	i = libc.Int32FromUint32((*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize) - int32(2)
	for {
		if !(i >= 0) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(libc.Uint32FromInt32(i) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+11067, __ccgo_ts+10918, int32(633), uintptr(unsafe.Pointer(&__func__121)))
			}
		}
		entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(i)*32
		is_visible = libc.BoolUint8(1 != 0)
		alias_symbol = uint16(0)
		if i > 0 {
			_ = libc.Uint64FromInt64(4)
			{
				if !(libc.Uint32FromInt32(i-libc.Int32FromInt32(1)) < (*struct {
					Fcontents uintptr
					Fsize     uint32_t
					Fcapacity uint32_t
				})(unsafe.Pointer(self+8)).Fsize) {
					libc.X__assert_fail(tls, __ccgo_ts+11104, __ccgo_ts+10918, int32(637), uintptr(unsafe.Pointer(&__func__121)))
				}
			}
			parent_entry = (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fcontents + uintptr(i-int32(1))*32
			alias_symbol = ts_language_alias_at(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), (*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index)
			is_visible = libc.BoolUint8(libc.Int32FromUint16(alias_symbol) != 0 || ts_subtree_visible(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))) != 0)
		}
		if is_visible != 0 {
			return ts_node_new(tls, (*TreeCursor)(unsafe.Pointer(self)).Ftree, (*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree, (*TreeCursorEntry)(unsafe.Pointer(entry)).Fposition, alias_symbol)
		}
		goto _1
	_1:
		;
		i = i - 1
	}
	return ts_node_new(tls, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), length_zero(tls), uint16(0))
}

var __func__121 = [27]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'p', 'a', 'r', 'e', 'n', 't', '_', 'n', 'o', 'd', 'e'}

func ts_tree_cursor_current_field_id(tls *libc.TLS, _self uintptr) (r TSFieldId) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var entry, map1, parent_entry, self uintptr
	var i uint32
	var _ /* field_map at bp+0 */ uintptr
	var _ /* field_map_end at bp+8 */ uintptr
	_, _, _, _, _ = entry, i, map1, parent_entry, self
	self = _self
	i = (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize - uint32(1)
	for {
		if !(i > uint32(0)) {
			break
		}
		_ = libc.Uint64FromInt64(4)
		{
			if !(i < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+11067, __ccgo_ts+10918, int32(662), uintptr(unsafe.Pointer(&__func__122)))
			}
		}
		entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(i)*32
		_ = libc.Uint64FromInt64(4)
		{
			if !(i-libc.Uint32FromInt32(1) < (*struct {
				Fcontents uintptr
				Fsize     uint32_t
				Fcapacity uint32_t
			})(unsafe.Pointer(self+8)).Fsize) {
				libc.X__assert_fail(tls, __ccgo_ts+11104, __ccgo_ts+10918, int32(663), uintptr(unsafe.Pointer(&__func__122)))
			}
		}
		parent_entry = (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(self+8)).Fcontents + uintptr(i-uint32(1))*32
		if i != (*TreeCursor)(unsafe.Pointer(self)).Fstack.Fsize-uint32(1) && ts_tree_cursor_is_entry_visible(tls, self, i) != 0 {
			break
		}
		if ts_subtree_extra(tls, *(*Subtree)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(entry)).Fsubtree))) != 0 {
			break
		}
		ts_language_field_map(tls, (*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage, uint32((*SubtreeHeapData)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer((*TreeCursorEntry)(unsafe.Pointer(parent_entry)).Fsubtree)))).F__ccgo19_48.F__ccgo0_0.Fproduction_id), bp, bp+8)
		map1 = *(*uintptr)(unsafe.Pointer(bp))
		for {
			if !(map1 < *(*uintptr)(unsafe.Pointer(bp + 8))) {
				break
			}
			if !((*TSFieldMapEntry)(unsafe.Pointer(map1)).Finherited != 0) && uint32((*TSFieldMapEntry)(unsafe.Pointer(map1)).Fchild_index) == (*TreeCursorEntry)(unsafe.Pointer(entry)).Fstructural_child_index {
				return (*TSFieldMapEntry)(unsafe.Pointer(map1)).Ffield_id
			}
			goto _2
		_2:
			;
			map1 += 4
		}
		goto _1
	_1:
		;
		i = i - 1
	}
	return uint16(0)
}

var __func__122 = [32]int8{'t', 's', '_', 't', 'r', 'e', 'e', '_', 'c', 'u', 'r', 's', 'o', 'r', '_', 'c', 'u', 'r', 'r', 'e', 'n', 't', '_', 'f', 'i', 'e', 'l', 'd', '_', 'i', 'd'}

func ts_tree_cursor_current_field_name(tls *libc.TLS, _self uintptr) (r uintptr) {
	var id TSFieldId
	var self uintptr
	_, _ = id, self
	id = ts_tree_cursor_current_field_id(tls, _self)
	if id != 0 {
		self = _self
		return *(*uintptr)(unsafe.Pointer((*TSLanguage)(unsafe.Pointer((*TSTree)(unsafe.Pointer((*TreeCursor)(unsafe.Pointer(self)).Ftree)).Flanguage)).Ffield_names + uintptr(id)*8))
	} else {
		return libc.UintptrFromInt32(0)
	}
	return r
}

func ts_tree_cursor_copy(tls *libc.TLS, _cursor uintptr) (r TSTreeCursor) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var copy1, cursor uintptr
	var _ /* res at bp+0 */ TSTreeCursor
	_, _ = copy1, cursor
	cursor = _cursor
	*(*TSTreeCursor)(unsafe.Pointer(bp)) = TSTreeCursor{}
	copy1 = bp
	(*TreeCursor)(unsafe.Pointer(copy1)).Ftree = (*TreeCursor)(unsafe.Pointer(cursor)).Ftree
	(*TreeCursor)(unsafe.Pointer(copy1)).Froot_alias_symbol = (*TreeCursor)(unsafe.Pointer(cursor)).Froot_alias_symbol
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1 + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1 + 8)).Fcapacity = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1 + 8)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1 + 8)).Fcontents = _array__splice(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1+8)).Fcontents, copy1+8+8, copy1+8+12, libc.Uint64FromInt64(32), (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1+8)).Fsize, uint32(0), (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor+8)).Fsize, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor+8)).Fcontents)
	return *(*TSTreeCursor)(unsafe.Pointer(bp))
}

func ts_tree_cursor_reset_to(tls *libc.TLS, _dst uintptr, _src uintptr) {
	var copy1, cursor uintptr
	_, _ = copy1, cursor
	cursor = _src
	copy1 = _dst
	(*TreeCursor)(unsafe.Pointer(copy1)).Ftree = (*TreeCursor)(unsafe.Pointer(cursor)).Ftree
	(*TreeCursor)(unsafe.Pointer(copy1)).Froot_alias_symbol = (*TreeCursor)(unsafe.Pointer(cursor)).Froot_alias_symbol
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1 + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1 + 8)).Fcontents = _array__splice(tls, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1+8)).Fcontents, copy1+8+8, copy1+8+12, libc.Uint64FromInt64(32), (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(copy1+8)).Fsize, uint32(0), (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor+8)).Fsize, (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(cursor+8)).Fcontents)
}

func ts_tree_new(tls *libc.TLS, root Subtree, language uintptr, included_ranges uintptr, included_range_count uint32) (r uintptr) {
	var result uintptr
	_ = result
	result = (*(*func(*libc.TLS, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_malloc})))(tls, uint64(32))
	(*TSTree)(unsafe.Pointer(result)).Froot = root
	(*TSTree)(unsafe.Pointer(result)).Flanguage = ts_language_copy(tls, language)
	(*TSTree)(unsafe.Pointer(result)).Fincluded_ranges = (*(*func(*libc.TLS, size_t, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_calloc})))(tls, uint64(included_range_count), uint64(24))
	libc.Xmemcpy(tls, (*TSTree)(unsafe.Pointer(result)).Fincluded_ranges, included_ranges, uint64(included_range_count)*uint64(24))
	(*TSTree)(unsafe.Pointer(result)).Fincluded_range_count = included_range_count
	return result
}

func ts_tree_copy(tls *libc.TLS, self uintptr) (r uintptr) {
	ts_subtree_retain(tls, (*TSTree)(unsafe.Pointer(self)).Froot)
	return ts_tree_new(tls, (*TSTree)(unsafe.Pointer(self)).Froot, (*TSTree)(unsafe.Pointer(self)).Flanguage, (*TSTree)(unsafe.Pointer(self)).Fincluded_ranges, (*TSTree)(unsafe.Pointer(self)).Fincluded_range_count)
}

func ts_tree_delete(tls *libc.TLS, self uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* pool at bp+0 */ SubtreePool
	if !(self != 0) {
		return
	}
	*(*SubtreePool)(unsafe.Pointer(bp)) = ts_subtree_pool_new(tls, uint32(0))
	ts_subtree_release(tls, bp, (*TSTree)(unsafe.Pointer(self)).Froot)
	ts_subtree_pool_delete(tls, bp)
	ts_language_delete(tls, (*TSTree)(unsafe.Pointer(self)).Flanguage)
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*TSTree)(unsafe.Pointer(self)).Fincluded_ranges)
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, self)
}

func ts_tree_root_node(tls *libc.TLS, self uintptr) (r TSNode) {
	return ts_node_new(tls, self, self, ts_subtree_padding(tls, (*TSTree)(unsafe.Pointer(self)).Froot), uint16(0))
}

func ts_tree_root_node_with_offset(tls *libc.TLS, self uintptr, offset_bytes uint32_t, offset_extent TSPoint) (r TSNode) {
	var offset Length
	_ = offset
	offset = Length{
		Fbytes:  offset_bytes,
		Fextent: offset_extent,
	}
	return ts_node_new(tls, self, self, length_add(tls, offset, ts_subtree_padding(tls, (*TSTree)(unsafe.Pointer(self)).Froot)), uint16(0))
}

func ts_tree_language(tls *libc.TLS, self uintptr) (r uintptr) {
	return (*TSTree)(unsafe.Pointer(self)).Flanguage
}

func ts_tree_edit(tls *libc.TLS, self uintptr, edit uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i uint32
	var _ /* pool at bp+0 */ SubtreePool
	_ = i
	i = uint32(0)
	for {
		if !(i < (*TSTree)(unsafe.Pointer(self)).Fincluded_range_count) {
			break
		}
		ts_range_edit(tls, (*TSTree)(unsafe.Pointer(self)).Fincluded_ranges+uintptr(i)*24, edit)
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*SubtreePool)(unsafe.Pointer(bp)) = ts_subtree_pool_new(tls, uint32(0))
	(*TSTree)(unsafe.Pointer(self)).Froot = ts_subtree_edit(tls, (*TSTree)(unsafe.Pointer(self)).Froot, edit, bp)
	ts_subtree_pool_delete(tls, bp)
}

func ts_tree_included_ranges(tls *libc.TLS, self uintptr, length uintptr) (r uintptr) {
	var ranges uintptr
	_ = ranges
	*(*uint32_t)(unsafe.Pointer(length)) = (*TSTree)(unsafe.Pointer(self)).Fincluded_range_count
	ranges = (*(*func(*libc.TLS, size_t, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{ts_current_calloc})))(tls, uint64((*TSTree)(unsafe.Pointer(self)).Fincluded_range_count), uint64(24))
	libc.Xmemcpy(tls, ranges, (*TSTree)(unsafe.Pointer(self)).Fincluded_ranges, uint64((*TSTree)(unsafe.Pointer(self)).Fincluded_range_count)*uint64(24))
	return ranges
}

func ts_tree_get_changed_ranges(tls *libc.TLS, old_tree uintptr, new_tree uintptr, length uintptr) (r uintptr) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var _ /* cursor1 at bp+0 */ TreeCursor
	var _ /* cursor2 at bp+32 */ TreeCursor
	var _ /* included_range_differences at bp+64 */ TSRangeArray
	var _ /* result at bp+80 */ uintptr
	*(*TreeCursor)(unsafe.Pointer(bp)) = TreeCursor{}
	*(*TreeCursor)(unsafe.Pointer(bp + 32)) = TreeCursor{}
	ts_tree_cursor_init(tls, bp, ts_tree_root_node(tls, old_tree))
	ts_tree_cursor_init(tls, bp+32, ts_tree_root_node(tls, new_tree))
	*(*TSRangeArray)(unsafe.Pointer(bp + 64)) = TSRangeArray{}
	ts_range_array_get_changed_ranges(tls, (*TSTree)(unsafe.Pointer(old_tree)).Fincluded_ranges, (*TSTree)(unsafe.Pointer(old_tree)).Fincluded_range_count, (*TSTree)(unsafe.Pointer(new_tree)).Fincluded_ranges, (*TSTree)(unsafe.Pointer(new_tree)).Fincluded_range_count, bp+64)
	*(*uint32_t)(unsafe.Pointer(length)) = ts_subtree_get_changed_ranges(tls, old_tree, new_tree, bp, bp+32, (*TSTree)(unsafe.Pointer(old_tree)).Flanguage, bp+64, bp+80)
	if (*TSRangeArray)(unsafe.Pointer(bp+64)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*TSRangeArray)(unsafe.Pointer(bp+64)).Fcontents)
	}
	(*TSRangeArray)(unsafe.Pointer(bp + 64)).Fcontents = libc.UintptrFromInt32(0)
	(*TSRangeArray)(unsafe.Pointer(bp + 64)).Fsize = uint32(0)
	(*TSRangeArray)(unsafe.Pointer(bp + 64)).Fcapacity = uint32(0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+8)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+8)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 8)).Fcapacity = uint32(0)
	if (*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp+32+8)).Fcontents != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{ts_current_free})))(tls, (*struct {
			Fcontents uintptr
			Fsize     uint32_t
			Fcapacity uint32_t
		})(unsafe.Pointer(bp+32+8)).Fcontents)
	}
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 32 + 8)).Fcontents = libc.UintptrFromInt32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 32 + 8)).Fsize = uint32(0)
	(*struct {
		Fcontents uintptr
		Fsize     uint32_t
		Fcapacity uint32_t
	})(unsafe.Pointer(bp + 32 + 8)).Fcapacity = uint32(0)
	return *(*uintptr)(unsafe.Pointer(bp + 80))
}

type useconds_t = uint32

type socklen_t = uint32

const _PC_LINK_MAX = 0
const _PC_MAX_CANON = 1
const _PC_MAX_INPUT = 2
const _PC_NAME_MAX = 3
const _PC_PATH_MAX = 4
const _PC_PIPE_BUF = 5
const _PC_CHOWN_RESTRICTED = 6
const _PC_NO_TRUNC = 7
const _PC_VDISABLE = 8
const _PC_SYNC_IO = 9
const _PC_ASYNC_IO = 10
const _PC_PRIO_IO = 11
const _PC_SOCK_MAXBUF = 12
const _PC_FILESIZEBITS = 13
const _PC_REC_INCR_XFER_SIZE = 14
const _PC_REC_MAX_XFER_SIZE = 15
const _PC_REC_MIN_XFER_SIZE = 16
const _PC_REC_XFER_ALIGN = 17
const _PC_ALLOC_SIZE_MIN = 18
const _PC_SYMLINK_MAX = 19
const _PC_2_SYMLINKS = 20
const _SC_ARG_MAX = 0
const _SC_CHILD_MAX = 1
const _SC_CLK_TCK = 2
const _SC_NGROUPS_MAX = 3
const _SC_OPEN_MAX = 4
const _SC_STREAM_MAX = 5
const _SC_TZNAME_MAX = 6
const _SC_JOB_CONTROL = 7
const _SC_SAVED_IDS = 8
const _SC_REALTIME_SIGNALS = 9
const _SC_PRIORITY_SCHEDULING = 10
const _SC_TIMERS = 11
const _SC_ASYNCHRONOUS_IO = 12
const _SC_PRIORITIZED_IO = 13
const _SC_SYNCHRONIZED_IO = 14
const _SC_FSYNC = 15
const _SC_MAPPED_FILES = 16
const _SC_MEMLOCK = 17
const _SC_MEMLOCK_RANGE = 18
const _SC_MEMORY_PROTECTION = 19
const _SC_MESSAGE_PASSING = 20
const _SC_SEMAPHORES = 21
const _SC_SHARED_MEMORY_OBJECTS = 22
const _SC_AIO_LISTIO_MAX = 23
const _SC_AIO_MAX = 24
const _SC_AIO_PRIO_DELTA_MAX = 25
const _SC_DELAYTIMER_MAX = 26
const _SC_MQ_OPEN_MAX = 27
const _SC_MQ_PRIO_MAX = 28
const _SC_VERSION = 29
const _SC_PAGESIZE = 30
const _SC_RTSIG_MAX = 31
const _SC_SEM_NSEMS_MAX = 32
const _SC_SEM_VALUE_MAX = 33
const _SC_SIGQUEUE_MAX = 34
const _SC_TIMER_MAX = 35
const _SC_BC_BASE_MAX = 36
const _SC_BC_DIM_MAX = 37
const _SC_BC_SCALE_MAX = 38
const _SC_BC_STRING_MAX = 39
const _SC_COLL_WEIGHTS_MAX = 40
const _SC_EQUIV_CLASS_MAX = 41
const _SC_EXPR_NEST_MAX = 42
const _SC_LINE_MAX = 43
const _SC_RE_DUP_MAX = 44
const _SC_CHARCLASS_NAME_MAX = 45
const _SC_2_VERSION = 46
const _SC_2_C_BIND = 47
const _SC_2_C_DEV = 48
const _SC_2_FORT_DEV = 49
const _SC_2_FORT_RUN = 50
const _SC_2_SW_DEV = 51
const _SC_2_LOCALEDEF = 52
const _SC_PII = 53
const _SC_PII_XTI = 54
const _SC_PII_SOCKET = 55
const _SC_PII_INTERNET = 56
const _SC_PII_OSI = 57
const _SC_POLL = 58
const _SC_SELECT = 59
const _SC_UIO_MAXIOV = 60
const _SC_IOV_MAX = 60
const _SC_PII_INTERNET_STREAM = 61
const _SC_PII_INTERNET_DGRAM = 62
const _SC_PII_OSI_COTS = 63
const _SC_PII_OSI_CLTS = 64
const _SC_PII_OSI_M = 65
const _SC_T_IOV_MAX = 66
const _SC_THREADS = 67
const _SC_THREAD_SAFE_FUNCTIONS = 68
const _SC_GETGR_R_SIZE_MAX = 69
const _SC_GETPW_R_SIZE_MAX = 70
const _SC_LOGIN_NAME_MAX = 71
const _SC_TTY_NAME_MAX = 72
const _SC_THREAD_DESTRUCTOR_ITERATIONS = 73
const _SC_THREAD_KEYS_MAX = 74
const _SC_THREAD_STACK_MIN = 75
const _SC_THREAD_THREADS_MAX = 76
const _SC_THREAD_ATTR_STACKADDR = 77
const _SC_THREAD_ATTR_STACKSIZE = 78
const _SC_THREAD_PRIORITY_SCHEDULING = 79
const _SC_THREAD_PRIO_INHERIT = 80
const _SC_THREAD_PRIO_PROTECT = 81
const _SC_THREAD_PROCESS_SHARED = 82
const _SC_NPROCESSORS_CONF = 83
const _SC_NPROCESSORS_ONLN = 84
const _SC_PHYS_PAGES = 85
const _SC_AVPHYS_PAGES = 86
const _SC_ATEXIT_MAX = 87
const _SC_PASS_MAX = 88
const _SC_XOPEN_VERSION = 89
const _SC_XOPEN_XCU_VERSION = 90
const _SC_XOPEN_UNIX = 91
const _SC_XOPEN_CRYPT = 92
const _SC_XOPEN_ENH_I18N = 93
const _SC_XOPEN_SHM = 94
const _SC_2_CHAR_TERM = 95
const _SC_2_C_VERSION = 96
const _SC_2_UPE = 97
const _SC_XOPEN_XPG2 = 98
const _SC_XOPEN_XPG3 = 99
const _SC_XOPEN_XPG4 = 100
const _SC_CHAR_BIT = 101
const _SC_CHAR_MAX = 102
const _SC_CHAR_MIN = 103
const _SC_INT_MAX = 104
const _SC_INT_MIN = 105
const _SC_LONG_BIT = 106
const _SC_WORD_BIT = 107
const _SC_MB_LEN_MAX = 108
const _SC_NZERO = 109
const _SC_SSIZE_MAX = 110
const _SC_SCHAR_MAX = 111
const _SC_SCHAR_MIN = 112
const _SC_SHRT_MAX = 113
const _SC_SHRT_MIN = 114
const _SC_UCHAR_MAX = 115
const _SC_UINT_MAX = 116
const _SC_ULONG_MAX = 117
const _SC_USHRT_MAX = 118
const _SC_NL_ARGMAX = 119
const _SC_NL_LANGMAX = 120
const _SC_NL_MSGMAX = 121
const _SC_NL_NMAX = 122
const _SC_NL_SETMAX = 123
const _SC_NL_TEXTMAX = 124
const _SC_XBS5_ILP32_OFF32 = 125
const _SC_XBS5_ILP32_OFFBIG = 126
const _SC_XBS5_LP64_OFF64 = 127
const _SC_XBS5_LPBIG_OFFBIG = 128
const _SC_XOPEN_LEGACY = 129
const _SC_XOPEN_REALTIME = 130
const _SC_XOPEN_REALTIME_THREADS = 131
const _SC_ADVISORY_INFO = 132
const _SC_BARRIERS = 133
const _SC_BASE = 134
const _SC_C_LANG_SUPPORT = 135
const _SC_C_LANG_SUPPORT_R = 136
const _SC_CLOCK_SELECTION = 137
const _SC_CPUTIME = 138
const _SC_THREAD_CPUTIME = 139
const _SC_DEVICE_IO = 140
const _SC_DEVICE_SPECIFIC = 141
const _SC_DEVICE_SPECIFIC_R = 142
const _SC_FD_MGMT = 143
const _SC_FIFO = 144
const _SC_PIPE = 145
const _SC_FILE_ATTRIBUTES = 146
const _SC_FILE_LOCKING = 147
const _SC_FILE_SYSTEM = 148
const _SC_MONOTONIC_CLOCK = 149
const _SC_MULTI_PROCESS = 150
const _SC_SINGLE_PROCESS = 151
const _SC_NETWORKING = 152
const _SC_READER_WRITER_LOCKS = 153
const _SC_SPIN_LOCKS = 154
const _SC_REGEXP = 155
const _SC_REGEX_VERSION = 156
const _SC_SHELL = 157
const _SC_SIGNALS = 158
const _SC_SPAWN = 159
const _SC_SPORADIC_SERVER = 160
const _SC_THREAD_SPORADIC_SERVER = 161
const _SC_SYSTEM_DATABASE = 162
const _SC_SYSTEM_DATABASE_R = 163
const _SC_TIMEOUTS = 164
const _SC_TYPED_MEMORY_OBJECTS = 165
const _SC_USER_GROUPS = 166
const _SC_USER_GROUPS_R = 167
const _SC_2_PBS = 168
const _SC_2_PBS_ACCOUNTING = 169
const _SC_2_PBS_LOCATE = 170
const _SC_2_PBS_MESSAGE = 171
const _SC_2_PBS_TRACK = 172
const _SC_SYMLOOP_MAX = 173
const _SC_STREAMS = 174
const _SC_2_PBS_CHECKPOINT = 175
const _SC_V6_ILP32_OFF32 = 176
const _SC_V6_ILP32_OFFBIG = 177
const _SC_V6_LP64_OFF64 = 178
const _SC_V6_LPBIG_OFFBIG = 179
const _SC_HOST_NAME_MAX = 180
const _SC_TRACE = 181
const _SC_TRACE_EVENT_FILTER = 182
const _SC_TRACE_INHERIT = 183
const _SC_TRACE_LOG = 184
const _SC_LEVEL1_ICACHE_SIZE = 185
const _SC_LEVEL1_ICACHE_ASSOC = 186
const _SC_LEVEL1_ICACHE_LINESIZE = 187
const _SC_LEVEL1_DCACHE_SIZE = 188
const _SC_LEVEL1_DCACHE_ASSOC = 189
const _SC_LEVEL1_DCACHE_LINESIZE = 190
const _SC_LEVEL2_CACHE_SIZE = 191
const _SC_LEVEL2_CACHE_ASSOC = 192
const _SC_LEVEL2_CACHE_LINESIZE = 193
const _SC_LEVEL3_CACHE_SIZE = 194
const _SC_LEVEL3_CACHE_ASSOC = 195
const _SC_LEVEL3_CACHE_LINESIZE = 196
const _SC_LEVEL4_CACHE_SIZE = 197
const _SC_LEVEL4_CACHE_ASSOC = 198
const _SC_LEVEL4_CACHE_LINESIZE = 199
const _SC_IPV6 = 235
const _SC_RAW_SOCKETS = 236
const _SC_V7_ILP32_OFF32 = 237
const _SC_V7_ILP32_OFFBIG = 238
const _SC_V7_LP64_OFF64 = 239
const _SC_V7_LPBIG_OFFBIG = 240
const _SC_SS_REPL_MAX = 241
const _SC_TRACE_EVENT_NAME_MAX = 242
const _SC_TRACE_NAME_MAX = 243
const _SC_TRACE_SYS_MAX = 244
const _SC_TRACE_USER_EVENT_MAX = 245
const _SC_XOPEN_STREAMS = 246
const _SC_THREAD_ROBUST_PRIO_INHERIT = 247
const _SC_THREAD_ROBUST_PRIO_PROTECT = 248
const _SC_MINSIGSTKSZ = 249
const _SC_SIGSTKSZ = 250
const _CS_PATH = 0
const _CS_V6_WIDTH_RESTRICTED_ENVS = 1
const _CS_GNU_LIBC_VERSION = 2
const _CS_GNU_LIBPTHREAD_VERSION = 3
const _CS_V5_WIDTH_RESTRICTED_ENVS = 4
const _CS_V7_WIDTH_RESTRICTED_ENVS = 5
const _CS_LFS_CFLAGS = 1000
const _CS_LFS_LDFLAGS = 1001
const _CS_LFS_LIBS = 1002
const _CS_LFS_LINTFLAGS = 1003
const _CS_LFS64_CFLAGS = 1004
const _CS_LFS64_LDFLAGS = 1005
const _CS_LFS64_LIBS = 1006
const _CS_LFS64_LINTFLAGS = 1007
const _CS_XBS5_ILP32_OFF32_CFLAGS = 1100
const _CS_XBS5_ILP32_OFF32_LDFLAGS = 1101
const _CS_XBS5_ILP32_OFF32_LIBS = 1102
const _CS_XBS5_ILP32_OFF32_LINTFLAGS = 1103
const _CS_XBS5_ILP32_OFFBIG_CFLAGS = 1104
const _CS_XBS5_ILP32_OFFBIG_LDFLAGS = 1105
const _CS_XBS5_ILP32_OFFBIG_LIBS = 1106
const _CS_XBS5_ILP32_OFFBIG_LINTFLAGS = 1107
const _CS_XBS5_LP64_OFF64_CFLAGS = 1108
const _CS_XBS5_LP64_OFF64_LDFLAGS = 1109
const _CS_XBS5_LP64_OFF64_LIBS = 1110
const _CS_XBS5_LP64_OFF64_LINTFLAGS = 1111
const _CS_XBS5_LPBIG_OFFBIG_CFLAGS = 1112
const _CS_XBS5_LPBIG_OFFBIG_LDFLAGS = 1113
const _CS_XBS5_LPBIG_OFFBIG_LIBS = 1114
const _CS_XBS5_LPBIG_OFFBIG_LINTFLAGS = 1115
const _CS_POSIX_V6_ILP32_OFF32_CFLAGS = 1116
const _CS_POSIX_V6_ILP32_OFF32_LDFLAGS = 1117
const _CS_POSIX_V6_ILP32_OFF32_LIBS = 1118
const _CS_POSIX_V6_ILP32_OFF32_LINTFLAGS = 1119
const _CS_POSIX_V6_ILP32_OFFBIG_CFLAGS = 1120
const _CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS = 1121
const _CS_POSIX_V6_ILP32_OFFBIG_LIBS = 1122
const _CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS = 1123
const _CS_POSIX_V6_LP64_OFF64_CFLAGS = 1124
const _CS_POSIX_V6_LP64_OFF64_LDFLAGS = 1125
const _CS_POSIX_V6_LP64_OFF64_LIBS = 1126
const _CS_POSIX_V6_LP64_OFF64_LINTFLAGS = 1127
const _CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS = 1128
const _CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS = 1129
const _CS_POSIX_V6_LPBIG_OFFBIG_LIBS = 1130
const _CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS = 1131
const _CS_POSIX_V7_ILP32_OFF32_CFLAGS = 1132
const _CS_POSIX_V7_ILP32_OFF32_LDFLAGS = 1133
const _CS_POSIX_V7_ILP32_OFF32_LIBS = 1134
const _CS_POSIX_V7_ILP32_OFF32_LINTFLAGS = 1135
const _CS_POSIX_V7_ILP32_OFFBIG_CFLAGS = 1136
const _CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS = 1137
const _CS_POSIX_V7_ILP32_OFFBIG_LIBS = 1138
const _CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS = 1139
const _CS_POSIX_V7_LP64_OFF64_CFLAGS = 1140
const _CS_POSIX_V7_LP64_OFF64_LDFLAGS = 1141
const _CS_POSIX_V7_LP64_OFF64_LIBS = 1142
const _CS_POSIX_V7_LP64_OFF64_LINTFLAGS = 1143
const _CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS = 1144
const _CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS = 1145
const _CS_POSIX_V7_LPBIG_OFFBIG_LIBS = 1146
const _CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS = 1147
const _CS_V6_ENV = 1148
const _CS_V7_ENV = 1149

func _ts_dup(tls *libc.TLS, file_descriptor int32) (r int32) {
	return libc.Xdup(tls, file_descriptor)
}

func ts_tree_print_dot_graph(tls *libc.TLS, self uintptr, file_descriptor int32) {
	var file uintptr
	_ = file
	file = libc.Xfdopen(tls, _ts_dup(tls, file_descriptor), __ccgo_ts+2711)
	ts_subtree_print_dot_graph(tls, (*TSTree)(unsafe.Pointer(self)).Froot, (*TSTree)(unsafe.Pointer(self)).Flanguage, file)
	libc.Xfclose(tls, file)
}

func ts_wasm_store_delete(tls *libc.TLS, self uintptr) {
	_ = self
}

func ts_wasm_store_start(tls *libc.TLS, self uintptr, lexer uintptr, language uintptr) (r uint8) {
	_ = self
	_ = lexer
	_ = language
	return libc.BoolUint8(0 != 0)
}

func ts_wasm_store_reset(tls *libc.TLS, self uintptr) {
	_ = self
}

func ts_wasm_store_call_lex_main(tls *libc.TLS, self uintptr, state TSStateId) (r uint8) {
	_ = self
	_ = state
	return libc.BoolUint8(0 != 0)
}

func ts_wasm_store_call_lex_keyword(tls *libc.TLS, self uintptr, state TSStateId) (r uint8) {
	_ = self
	_ = state
	return libc.BoolUint8(0 != 0)
}

func ts_wasm_store_call_scanner_create(tls *libc.TLS, self uintptr) (r uint32_t) {
	_ = self
	return uint32(0)
}

func ts_wasm_store_call_scanner_destroy(tls *libc.TLS, self uintptr, scanner_address uint32_t) {
	_ = self
	_ = scanner_address
}

func ts_wasm_store_call_scanner_scan(tls *libc.TLS, self uintptr, scanner_address uint32_t, valid_tokens_ix uint32_t) (r uint8) {
	_ = self
	_ = scanner_address
	_ = valid_tokens_ix
	return libc.BoolUint8(0 != 0)
}

func ts_wasm_store_call_scanner_serialize(tls *libc.TLS, self uintptr, scanner_address uint32_t, buffer uintptr) (r uint32_t) {
	_ = self
	_ = scanner_address
	_ = buffer
	return uint32(0)
}

func ts_wasm_store_call_scanner_deserialize(tls *libc.TLS, self uintptr, scanner_address uint32_t, buffer uintptr, length uint32) {
	_ = self
	_ = scanner_address
	_ = buffer
	_ = length
}

func ts_wasm_store_has_error(tls *libc.TLS, self uintptr) (r uint8) {
	_ = self
	return libc.BoolUint8(0 != 0)
}

func ts_language_is_wasm(tls *libc.TLS, self uintptr) (r uint8) {
	_ = self
	return libc.BoolUint8(0 != 0)
}

func ts_wasm_language_retain(tls *libc.TLS, self uintptr) {
	_ = self
}

func ts_wasm_language_release(tls *libc.TLS, self uintptr) {
	_ = self
}

func __atomic_add_fetch(tls *libc.TLS, p uintptr, v uint32_t, m int32) (r uint32_t) {
	*(*uint32_t)(unsafe.Pointer(p)) += v
	return libc.AtomicLoadPUint32(p)
}

func __atomic_sub_fetch(tls *libc.TLS, p uintptr, v uint32_t, m int32) (r uint32_t) {
	*(*uint32_t)(unsafe.Pointer(p)) -= v
	return libc.AtomicLoadPUint32(p)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var ts_current_calloc = uintptr(0)

var ts_current_free = uintptr(0)

var ts_current_malloc = uintptr(0)

var ts_current_realloc = uintptr(0)

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "tree-sitter failed to allocate %zu bytes\x00tree-sitter failed to reallocate %zu bytes\x00index < *size\x00third-party/tree-sitter/lib/src/./././././array.h\x00old_end <= *size\x00(uint32_t)((&self->stack)->size - 1) < (&self->stack)->size\x00third-party/tree-sitter/lib/src/./././tree_cursor.h\x00\\n\x00\\t\x00(uint32_t)((self)->size - 1) < (self)->size\x00third-party/tree-sitter/lib/src/./get_changed_ranges.c\x00(uint32_t)(i) < (self)->size\x00(uint32_t)((&self->cursor.stack)->size - 1) < (&self->cursor.stack)->size\x00(uint32_t)(self->cursor.stack.size - 2) < (&self->cursor.stack)->size\x00(uint32_t)(i) < (&self->cursor.stack)->size\x00(uint32_t)(i - 1) < (&self->cursor.stack)->size\x00(uint32_t)(included_range_difference_index) < (included_range_differences)->size\x00symbol < self->token_count\x00third-party/tree-sitter/lib/src/./language.c\x00ERROR\x00_ERROR\x00 000000000000\x1000\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1e\x0f\x0f\x0f\x00\x00\x00\x00\x00skip character:'%c'\x00skip character:%d\x00consume character:'%c'\x00consume character:%d\x00third-party/tree-sitter/lib/src/././reusable_node.h\x00\x00graph {\nlabel=\"\x00\"\n}\n\n\x00(uint32_t)(i) < (&pop)->size\x00third-party/tree-sitter/lib/src/./parser.c\x00(uint32_t)(0) < (&slice.subtrees)->size\x00(uint32_t)(j) < (&slice.subtrees)->size\x00breakdown_top_of_stack tree:%s\x00\n\n\x00state_mismatch sym:%s\x00length <= 1024\x00no_lookahead_after_non_terminal_extra\x00lex_external state:%d, row:%u, column:%u\x00ignore_empty_external_token symbol:%s\x00lex_internal state:%d, row:%u, column:%u\x00skip_unrecognized_character\x00lexed_lookahead sym:\x00, size:%u\x00before_reusable_node symbol:%s\x00past_reusable_node symbol:%s\x00reusable_node_has_different_external_scanner_state symbol:%s\x00has_changes\x00is_error\x00is_missing\x00is_fragile\x00contains_different_included_range\x00cant_reuse_node_%s tree:%s\x00cant_reuse_node symbol:%s, first_leaf_symbol:%s\x00reuse_node symbol:%s\x00select_smaller_error symbol:%s, over_symbol:%s\x00select_higher_precedence symbol:%s, prec:%d, over_symbol:%s, other_prec:%d\x00select_earlier symbol:%s, over_symbol:%s\x00select_existing symbol:%s, over_symbol:%s\x00aborting reduce with too many versions\x00(uint32_t)(i + 1) < (&pop)->size\x00(uint32_t)(j) < (&self->trailing_extras)->size\x00ts_subtree_is_eof(lookahead)\x00(uint32_t)(j) < (&trees)->size\x00!tree.data.is_inline\x00root.ptr\x00(uint32_t)(0) < (&pop)->size\x00(uint32_t)(j) < (&self->reduce_actions)->size\x00error_trees.size == 1\x00(uint32_t)(0) < (&error_trees)->size\x00(uint32_t)(i) < (summary)->size\x00recover_to_previous state:%u, depth:%u\x00removed paused version:%u\x00recover_eof\x00skip_token symbol:%s\x00recover_with_missing symbol:%s, state:%u\x00did_merge\x00shift_extra\x00shift state:%u\x00reduce sym:%s, child_count:%u\x00accept\x00switch from_keyword:%s, to_word_token:%s\x00detect_error lookahead:%s\x00resume version:%u\x00condense\x00(uint32_t)(self->tree_pool.tree_stack.size - 1) < (&self->tree_pool.tree_stack)->size\x00a\x00resume_parsing\x00parse_after_edit\x00\n\x00(uint32_t)(i) < (&self->included_range_differences)->size\x00different_included_range %u - %u\x00new_parse\x00process version:%u, version_count:%u, state:%d, row:%u, col:%u\x00(uint32_t)(self->included_range_difference_index) < (&self->included_range_differences)->size\x00self->finished_tree.ptr\x00done\x00(uint32_t)(i) < (&self->list)->size\x00third-party/tree-sitter/lib/src/./query.c\x00(uint32_t)(id) < (&self->list)->size\x00id < self->list.size\x00(uint32_t)(id) < (self)->size\x00(uint32_t)(id) < (quantifiers)->size\x00(uint32_t)(i) < (&self->slices)->size\x00(uint32_t)(slice.offset) < (&self->characters)->size\x00(uint32_t)(id) < (&self->slices)->size\x00(uint32_t)(self->characters.size - 1) < (&self->characters)->size\x00(uint32_t)(mid_index) < (&self->pattern_map)->size\x00(uint32_t)((((void) sizeof (((uint32_t)(mid_index) < (&self->pattern_map)->size) ? 1 : 0), __extension__ ({ if ((uint32_t)(mid_index) < (&self->pattern_map)->size) ; else __assert_fail (\"(uint32_t)(mid_index) < (&self->pattern_map)->size\", \"third-party/tree-sitter/lib/src/./query.c\", 1106, __extension__ __PRETTY_FUNCTION__); })), &(&self->pattern_map)->contents[mid_index])->step_index) < (&self->steps)->size\x00(uint32_t)(base_index) < (&self->pattern_map)->size\x00(uint32_t)((((void) sizeof (((uint32_t)(base_index) < (&self->pattern_map)->size) ? 1 : 0), __extension__ ({ if ((uint32_t)(base_index) < (&self->pattern_map)->size) ; else __assert_fail (\"(uint32_t)(base_index) < (&self->pattern_map)->size\", \"third-party/tree-sitter/lib/src/./query.c\", 1113, __extension__ __PRETTY_FUNCTION__); })), &(&self->pattern_map)->contents[base_index])->step_index) < (&self->steps)->size\x00(uint32_t)((((void) sizeof (((uint32_t)(base_index) < (&self->pattern_map)->size) ? 1 : 0), __extension__ ({ if ((uint32_t)(base_index) < (&self->pattern_map)->size) ; else __assert_fail (\"(uint32_t)(base_index) < (&self->pattern_map)->size\", \"third-party/tree-sitter/lib/src/./query.c\", 1120, __extension__ __PRETTY_FUNCTION__); })), &(&self->pattern_map)->contents[base_index])->step_index) < (&self->steps)->size\x00(uint32_t)(index) < (&self->pattern_map)->size\x00(uint32_t)(entry->step_index) < (&self->steps)->size\x00(uint32_t)(j) < (&analysis->states)->size\x00(uint32_t)((&analysis->next_states)->size - 1) < (&analysis->next_states)->size\x00(uint32_t)(state->step_index) < (&self->steps)->size\x00(uint32_t)(subgraph_index) < (subgraphs)->size\x00(uint32_t)(node_index) < (&subgraph->nodes)->size\x00(uint32_t)(next_state.step_index) < (&self->steps)->size\x00(uint32_t)(i) < (&self->pattern_map)->size\x00(uint32_t)(pattern->step_index) < (&self->steps)->size\x00(uint32_t)(i) < (&self->steps)->size\x00(uint32_t)(j) < (&self->steps)->size\x00(uint32_t)(offset_idx) < (&self->step_offsets)->size\x00(uint32_t)(i) < (&parent_step_indices)->size\x00(uint32_t)(parent_step_index) < (&self->steps)->size\x00(uint32_t)(subgraph_index) < (&subgraphs)->size\x00(uint32_t)((&subgraph->nodes)->size - 1) < (&subgraph->nodes)->size\x00(uint32_t)((&subgraph->start_states)->size - 1) < (&subgraph->start_states)->size\x00(uint32_t)(i) < (&subgraphs)->size\x00child_exists\x00(uint32_t)(j) < (&self->step_offsets)->size\x00(uint32_t)(j) < (&subgraph->start_states)->size\x00(uint32_t)((&analysis.final_step_indices)->size - 1) < (&analysis.final_step_indices)->size\x00(uint32_t)(j) < (&analysis.final_step_indices)->size\x00(uint32_t)(final_step_index) < (&self->steps)->size\x00(uint32_t)(i) < (&self->patterns)->size\x00(uint32_t)(j) < (&self->predicate_steps)->size\x00(uint32_t)(step->alternative_index) < (&self->steps)->size\x00(uint32_t)(i - 1) < (&self->steps)->size\x00(uint32_t)(i) < (&non_rooted_pattern_start_steps)->size\x00(uint32_t)(pattern_entry_index) < (&self->pattern_map)->size\x00(uint32_t)(j) < (&subgraphs)->size\x00(uint32_t)(k) < (&subgraph->start_states)->size\x00(uint32_t)(pattern_entry->pattern_index) < (&self->patterns)->size\x00(uint32_t)(k) < (&analysis.finished_parent_symbols)->size\x00(uint32_t)(step_index) < (&self->steps)->size\x00(uint32_t)(i) < (&self->negated_fields)->size\x00(uint32_t)((&self->step_offsets)->size - 1) < (&self->step_offsets)->size\x00(uint32_t)(i) < (&branch_step_indices)->size\x00(uint32_t)(i + 1) < (&branch_step_indices)->size\x00(uint32_t)(next_step_index - 1) < (&self->steps)->size\x00MISSING\x00(uint32_t)((&self->steps)->size - 1) < (&self->steps)->size\x00(uint32_t)(last_child_step_index) < (&self->steps)->size\x00(uint32_t)(last_child_step->alternative_index) < (&self->steps)->size\x00(uint32_t)(alternative_step->alternative_index) < (&self->steps)->size\x00(uint32_t)(starting_step_index) < (&self->steps)->size\x00(uint32_t)((&self->patterns)->size - 1) < (&self->patterns)->size\x00(uint32_t)(start_step_index) < (&self->steps)->size\x00(uint32_t)(start_step_index + 1) < (&self->steps)->size\x00(uint32_t)(target_idx) < (&self->steps)->size\x00(uint32_t)(index) < (&self->capture_quantifiers)->size\x00(uint32_t)(pattern_index) < (&self->capture_quantifiers)->size\x00(uint32_t)(pattern_index) < (&self->patterns)->size\x00(uint32_t)(slice.offset) < (&self->predicate_steps)->size\x00(uint32_t)(i) < (&self->step_offsets)->size\x00(uint32_t)step_index + 1 < self->steps.size\x00(uint32_t)(step_index + 1) < (&self->steps)->size\x00(uint32_t)(i) < (&query->steps)->size\x00(uint32_t)(i) < (&self->states)->size\x00(uint32_t)(state->consumed_capture_count) < (captures)->size\x00(uint32_t)(state->step_index) < (&self->query->steps)->size\x00(uint32_t)(i) < (left_captures)->size\x00(uint32_t)(j) < (right_captures)->size\x00(uint32_t)(pattern->step_index) < (&self->query->steps)->size\x00(uint32_t)(index - 1) < (&self->states)->size\x00(uint32_t)(state_index) < (&self->states)->size\x00(uint32_t)(state_index + 1) < (&self->states)->size\x00(uint32_t)(i - deleted_count) < (&self->states)->size\x00(uint32_t)(i) < (&self->query->pattern_map)->size\x00(uint32_t)(j) < (&self->states)->size\x00(uint32_t)(step->negated_field_list_id) < (&self->query->negated_fields)->size\x00(uint32_t)(k) < (&self->states)->size\x00(uint32_t)(child_state->step_index) < (&self->query->steps)->size\x00(uint32_t)(0) < (&self->finished_states)->size\x00(uint32_t)(i) < (&self->finished_states)->size\x00(uint32_t)(first_unfinished_state_index) < (&self->states)->size\x00self->ref_count > 0\x00third-party/tree-sitter/lib/src/./stack.c\x00self->ref_count != 0\x00(uint32_t)(original_version) < (&self->heads)->size\x00(uint32_t)(version) < (&self->heads)->size\x00(uint32_t)(i) < (&self->iterators)->size\x00(uint32_t)((&self->iterators)->size - 1) < (&self->iterators)->size\x00(uint32_t)(i) < (&self->heads)->size\x00(uint32_t)(i) < (&self->node_pool)->size\x00(uint32_t)(0) < (&iterator->subtrees)->size\x00pop.size == 1\x00(uint32_t)(i) < (session->summary)->size\x00v2 < v1\x00(uint32_t)v1 < self->heads.size\x00(uint32_t)(v1) < (&self->heads)->size\x00(uint32_t)(v2) < (&self->heads)->size\x00version < self->heads.size\x00(uint32_t)((&self->heads)->size - 1) < (&self->heads)->size\x00(uint32_t)(version1) < (&self->heads)->size\x00(uint32_t)(version2) < (&self->heads)->size\x00head->status == StackStatusPaused\x00digraph stack {\n\x00rankdir=\"RL\";\n\x00edge [arrowhead=none]\n\x00node_head_%u [shape=none, label=\"\"]\n\x00node_head_%u -> node_%p [\x00color=red \x00label=%u, fontcolor=blue, weight=10000, labeltooltip=\"node_count: %u\nerror_cost: %u\x00\nsummary:\x00(uint32_t)(j) < (head->summary)->size\x00 %u\x00\nexternal_scanner_state:\x00 %2X\x00\"]\n\x00(uint32_t)(j) < (&visited_nodes)->size\x00node_%p [\x00label=\"?\"\x00shape=point margin=0 label=\"\"\x00label=\"%d\"\x00 tooltip=\"position: %u,%u\nnode_count:%u\nerror_cost: %u\ndynamic_precedence: %d\"];\n\x00node_%p -> node_%p [\x00style=dashed \x00fontcolor=gray \x00color=red\x00label=\"\x00'\x00\"\x00labeltooltip=\"error_cost: %u\ndynamic_precedence: %d\"\x00];\n\x00}\n\x00(uint32_t)(i) < (dest)->size\x00third-party/tree-sitter/lib/src/./subtree.c\x00(uint32_t)(self->size - 1) < (self)->size\x00(uint32_t)(reverse_index) < (self)->size\x00(uint32_t)(i) < (&self->free_trees)->size\x00symbol < (255)\x00!self.data.is_inline\x00self.ptr->ref_count > 0\x00self.ptr->ref_count != 0\x00child.ptr->ref_count > 0\x00INVALID\x00'\\0'\x00'\\n'\x00'\\t'\x00'\\r'\x00'%c'\x00%d\x00__ROOT__\x00(NULL)\x00 \x00%s: \x00(UNEXPECTED \x00(MISSING \x00%s\x00\"%s\"\x00(%s\x00(%s)\x00(\"%s\")\x00)\x00tree_%p [label=\"\x00, shape=plaintext\x00, fontcolor=gray\x00, color=green, penwidth=2\x00, tooltip=\"range: %u - %u\nstate: %d\nerror-cost: %u\nhas-changes: %u\ndepends-on-column: %u\ndescendant-count: %u\nrepeat-depth: %u\nlookahead-bytes: %u\x00\ncharacter: '%c'\x00tree_%p -> tree_%p [tooltip=%u]\n\x00digraph tree {\n\x00(uint32_t)(index) < (&self->stack)->size\x00third-party/tree-sitter/lib/src/./tree_cursor.c\x00(uint32_t)(index - 1) < (&self->stack)->size\x00(uint32_t)(self->stack.size - 2) < (&self->stack)->size\x00(uint32_t)(i) < (&self->stack)->size\x00(uint32_t)(i - 1) < (&self->stack)->size\x00"
