// Code generated for linux/amd64 by 'ccgo /tmp/grammar-gen-1085035175/preprocessed.c -o /tmp/grammar-gen-1085035175/grammar.go', DO NOT EDIT.

//go:build linux && amd64

package grammar_lua

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
	F__pad5          size_t
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
	F__pad5          size_t
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
	F__pad5          size_t
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

type __locale_struct = struct {
	F__locales       [13]uintptr
	F__ctype_b       uintptr
	F__ctype_tolower uintptr
	F__ctype_toupper uintptr
	F__names         [13]uintptr
}

type __locale_t = uintptr

type locale_t = uintptr

type TokenType = int32

const BLOCK_COMMENT_START = 0
const BLOCK_COMMENT_CONTENT = 1
const BLOCK_COMMENT_END = 2
const BLOCK_STRING_START = 3
const BLOCK_STRING_CONTENT = 4
const BLOCK_STRING_END = 5

func consume(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(0 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(1 != 0))
}

func consume_char(tls *libc.TLS, c int8, lexer uintptr) (r uint8) {
	if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32(c) {
		return libc.BoolUint8(0 != 0)
	}
	consume(tls, lexer)
	return libc.BoolUint8(1 != 0)
}

func consume_and_count_char(tls *libc.TLS, c int8, lexer uintptr) (r uint8_t) {
	var count uint8_t
	_ = count
	count = uint8(0)
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(c) {
		count = count + 1
		consume(tls, lexer)
	}
	return count
}

func skip_whitespaces(tls *libc.TLS, lexer uintptr) {
	for libc.Xiswspace(tls, libc.Uint32FromInt32((*TSLexer)(unsafe.Pointer(lexer)).Flookahead)) != 0 {
		skip(tls, lexer)
	}
}

type Scanner = struct {
	Fending_char int8
	Flevel_count uint8_t
}

func reset_state(tls *libc.TLS, scanner uintptr) {
	(*Scanner)(unsafe.Pointer(scanner)).Fending_char = 0
	(*Scanner)(unsafe.Pointer(scanner)).Flevel_count = uint8(0)
}

func tree_sitter_lua_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = libc.Xcalloc(tls, uint64(1), uint64(2))
	return scanner
}

func tree_sitter_lua_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	libc.Xfree(tls, scanner)
}

func tree_sitter_lua_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	*(*int8)(unsafe.Pointer(buffer)) = (*Scanner)(unsafe.Pointer(scanner)).Fending_char
	*(*int8)(unsafe.Pointer(buffer + 1)) = libc.Int8FromUint8((*Scanner)(unsafe.Pointer(scanner)).Flevel_count)
	return uint32(2)
}

func tree_sitter_lua_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	if length == uint32(0) {
		return
	}
	(*Scanner)(unsafe.Pointer(scanner)).Fending_char = *(*int8)(unsafe.Pointer(buffer))
	if length == uint32(1) {
		return
	}
	(*Scanner)(unsafe.Pointer(scanner)).Flevel_count = libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(buffer + 1)))
}

func scan_block_start(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	var level uint8_t
	_ = level
	if consume_char(tls, int8('['), lexer) != 0 {
		level = consume_and_count_char(tls, int8('='), lexer)
		if consume_char(tls, int8('['), lexer) != 0 {
			(*Scanner)(unsafe.Pointer(scanner)).Flevel_count = level
			return libc.BoolUint8(1 != 0)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_block_end(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	var level uint8_t
	_ = level
	if consume_char(tls, int8(']'), lexer) != 0 {
		level = consume_and_count_char(tls, int8('='), lexer)
		if libc.Int32FromUint8((*Scanner)(unsafe.Pointer(scanner)).Flevel_count) == libc.Int32FromUint8(level) && consume_char(tls, int8(']'), lexer) != 0 {
			return libc.BoolUint8(1 != 0)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_block_content(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32(']') {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
			if scan_block_end(tls, scanner, lexer) != 0 {
				return libc.BoolUint8(1 != 0)
			}
		} else {
			consume(tls, lexer)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_comment_start(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	if consume_char(tls, int8('-'), lexer) != 0 && consume_char(tls, int8('-'), lexer) != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if scan_block_start(tls, scanner, lexer) != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_COMMENT_START)
			return libc.BoolUint8(1 != 0)
		}
	}
	return libc.BoolUint8(0 != 0)
}

func scan_comment_content(tls *libc.TLS, scanner uintptr, lexer uintptr) (r uint8) {
	if int32((*Scanner)(unsafe.Pointer(scanner)).Fending_char) == 0 {
		if scan_block_content(tls, scanner, lexer) != 0 {
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_COMMENT_CONTENT)
			return libc.BoolUint8(1 != 0)
		}
		return libc.BoolUint8(0 != 0)
	}
	for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32((*Scanner)(unsafe.Pointer(scanner)).Fending_char) {
			reset_state(tls, scanner)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_COMMENT_CONTENT)
			return libc.BoolUint8(1 != 0)
		}
		consume(tls, lexer)
	}
	return libc.BoolUint8(0 != 0)
}

func tree_sitter_lua_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_STRING_END))) != 0 && scan_block_end(tls, scanner, lexer) != 0 {
		reset_state(tls, scanner)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_STRING_END)
		return libc.BoolUint8(1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_STRING_CONTENT))) != 0 && scan_block_content(tls, scanner, lexer) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_STRING_CONTENT)
		return libc.BoolUint8(1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_COMMENT_END))) != 0 && int32((*Scanner)(unsafe.Pointer(scanner)).Fending_char) == 0 && scan_block_end(tls, scanner, lexer) != 0 {
		reset_state(tls, scanner)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_COMMENT_END)
		return libc.BoolUint8(1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_COMMENT_CONTENT))) != 0 && scan_comment_content(tls, scanner, lexer) != 0 {
		return libc.BoolUint8(1 != 0)
	}
	skip_whitespaces(tls, lexer)
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_STRING_START))) != 0 && scan_block_start(tls, scanner, lexer) != 0 {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BLOCK_STRING_START)
		return libc.BoolUint8(1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BLOCK_COMMENT_START))) != 0 {
		if scan_comment_start(tls, scanner, lexer) != 0 {
			return libc.BoolUint8(1 != 0)
		}
	}
	return libc.BoolUint8(0 != 0)
}

type ts_symbol_identifiers = int32

const sym_identifier = 1
const sym_hash_bang_line = 2
const anon_sym_return = 3
const anon_sym_SEMI = 4
const anon_sym_EQ = 5
const anon_sym_COMMA = 6
const anon_sym_COLON_COLON = 7
const sym_break_statement = 8
const anon_sym_goto = 9
const anon_sym_do = 10
const anon_sym_end = 11
const anon_sym_while = 12
const anon_sym_repeat = 13
const anon_sym_until = 14
const anon_sym_if = 15
const anon_sym_then = 16
const anon_sym_elseif = 17
const anon_sym_else = 18
const anon_sym_for = 19
const anon_sym_in = 20
const anon_sym_function = 21
const anon_sym_local = 22
const anon_sym_global = 23
const anon_sym_DOT = 24
const anon_sym_COLON = 25
const anon_sym_STAR = 26
const anon_sym_LT = 27
const anon_sym_GT = 28
const sym_nil = 29
const sym_false = 30
const sym_true = 31
const sym_number = 32
const anon_sym_DQUOTE = 33
const anon_sym_SQUOTE = 34
const aux_sym__doublequote_string_content_token1 = 35
const aux_sym__singlequote_string_content_token1 = 36
const sym_escape_sequence = 37
const sym_vararg_expression = 38
const anon_sym_LPAREN = 39
const anon_sym_RPAREN = 40
const anon_sym_LBRACK = 41
const anon_sym_RBRACK = 42
const anon_sym_LBRACE = 43
const anon_sym_RBRACE = 44
const anon_sym_or = 45
const anon_sym_and = 46
const anon_sym_LT_EQ = 47
const anon_sym_EQ_EQ = 48
const anon_sym_TILDE_EQ = 49
const anon_sym_GT_EQ = 50
const anon_sym_PIPE = 51
const anon_sym_TILDE = 52
const anon_sym_AMP = 53
const anon_sym_LT_LT = 54
const anon_sym_GT_GT = 55
const anon_sym_PLUS = 56
const anon_sym_DASH = 57
const anon_sym_SLASH = 58
const anon_sym_SLASH_SLASH = 59
const anon_sym_PERCENT = 60
const anon_sym_DOT_DOT = 61
const anon_sym_CARET = 62
const anon_sym_not = 63
const anon_sym_POUND = 64
const anon_sym_DASH_DASH = 65
const aux_sym_comment_token1 = 66
const sym__block_comment_start = 67
const sym__block_comment_content = 68
const sym__block_comment_end = 69
const sym__block_string_start = 70
const sym__block_string_content = 71
const sym__block_string_end = 72
const sym_chunk = 73
const sym__block = 74
const sym_statement = 75
const sym_return_statement = 76
const sym_empty_statement = 77
const sym_assignment_statement = 78
const sym__variable_assignment_varlist = 79
const sym__variable_assignment_explist = 80
const sym_label_statement = 81
const sym_goto_statement = 82
const sym_do_statement = 83
const sym_while_statement = 84
const sym_repeat_statement = 85
const sym_if_statement = 86
const sym_elseif_statement = 87
const sym_else_statement = 88
const sym_for_statement = 89
const sym_for_generic_clause = 90
const sym_for_numeric_clause = 91
const sym__name_list = 92
const sym_declaration = 93
const sym_function_declaration = 94
const sym__local_function_declaration = 95
const sym__global_function_declaration = 96
const sym__function_name = 97
const sym__function_name_prefix_expression = 98
const sym__function_name_dot_index_expression = 99
const sym__function_name_method_index_expression = 100
const sym_variable_declaration = 101
const sym__global_variable_declaration = 102
const sym__variable_assignment = 103
const sym__att_name_list = 104
const sym__global_implicit_variable_declaration = 105
const sym__attrib = 106
const sym__expression_list = 107
const sym_expression = 108
const sym_string = 109
const sym__quote_string = 110
const aux_sym__doublequote_string_content = 111
const aux_sym__singlequote_string_content = 112
const sym__block_string = 113
const sym_function_definition = 114
const sym__function_body = 115
const sym_parameters = 116
const sym__parameter_list = 117
const sym__vararg_parameter = 118
const sym__prefix_expression = 119
const sym_variable = 120
const sym_bracket_index_expression = 121
const sym_dot_index_expression = 122
const sym_function_call = 123
const sym_method_index_expression = 124
const sym_arguments = 125
const sym_parenthesized_expression = 126
const sym_table_constructor = 127
const sym__field_list = 128
const sym__field_sep = 129
const sym_field = 130
const sym_binary_expression = 131
const sym_unary_expression = 132
const sym_comment = 133
const sym__contextual_keyword = 134
const aux_sym_chunk_repeat1 = 135
const aux_sym__variable_assignment_varlist_repeat1 = 136
const aux_sym__variable_assignment_explist_repeat1 = 137
const aux_sym_if_statement_repeat1 = 138
const aux_sym__name_list_repeat1 = 139
const aux_sym__att_name_list_repeat1 = 140
const aux_sym__expression_list_repeat1 = 141
const aux_sym__field_list_repeat1 = 142

var ts_symbol_names = [143]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 15,
	3:   __ccgo_ts + 30,
	4:   __ccgo_ts + 37,
	5:   __ccgo_ts + 39,
	6:   __ccgo_ts + 41,
	7:   __ccgo_ts + 43,
	8:   __ccgo_ts + 46,
	9:   __ccgo_ts + 62,
	10:  __ccgo_ts + 67,
	11:  __ccgo_ts,
	12:  __ccgo_ts + 70,
	13:  __ccgo_ts + 76,
	14:  __ccgo_ts + 83,
	15:  __ccgo_ts + 89,
	16:  __ccgo_ts + 92,
	17:  __ccgo_ts + 97,
	18:  __ccgo_ts + 104,
	19:  __ccgo_ts + 109,
	20:  __ccgo_ts + 113,
	21:  __ccgo_ts + 116,
	22:  __ccgo_ts + 125,
	23:  __ccgo_ts + 131,
	24:  __ccgo_ts + 138,
	25:  __ccgo_ts + 140,
	26:  __ccgo_ts + 142,
	27:  __ccgo_ts + 144,
	28:  __ccgo_ts + 146,
	29:  __ccgo_ts + 148,
	30:  __ccgo_ts + 152,
	31:  __ccgo_ts + 158,
	32:  __ccgo_ts + 163,
	33:  __ccgo_ts + 170,
	34:  __ccgo_ts + 172,
	35:  __ccgo_ts + 174,
	36:  __ccgo_ts + 209,
	37:  __ccgo_ts + 244,
	38:  __ccgo_ts + 260,
	39:  __ccgo_ts + 278,
	40:  __ccgo_ts + 280,
	41:  __ccgo_ts + 282,
	42:  __ccgo_ts + 284,
	43:  __ccgo_ts + 286,
	44:  __ccgo_ts + 288,
	45:  __ccgo_ts + 290,
	46:  __ccgo_ts + 293,
	47:  __ccgo_ts + 297,
	48:  __ccgo_ts + 300,
	49:  __ccgo_ts + 303,
	50:  __ccgo_ts + 306,
	51:  __ccgo_ts + 309,
	52:  __ccgo_ts + 311,
	53:  __ccgo_ts + 313,
	54:  __ccgo_ts + 315,
	55:  __ccgo_ts + 318,
	56:  __ccgo_ts + 321,
	57:  __ccgo_ts + 323,
	58:  __ccgo_ts + 325,
	59:  __ccgo_ts + 327,
	60:  __ccgo_ts + 330,
	61:  __ccgo_ts + 332,
	62:  __ccgo_ts + 335,
	63:  __ccgo_ts + 337,
	64:  __ccgo_ts + 341,
	65:  __ccgo_ts + 343,
	66:  __ccgo_ts + 346,
	67:  __ccgo_ts + 362,
	68:  __ccgo_ts + 346,
	69:  __ccgo_ts + 365,
	70:  __ccgo_ts + 362,
	71:  __ccgo_ts + 368,
	72:  __ccgo_ts + 365,
	73:  __ccgo_ts + 383,
	74:  __ccgo_ts + 389,
	75:  __ccgo_ts + 395,
	76:  __ccgo_ts + 405,
	77:  __ccgo_ts + 422,
	78:  __ccgo_ts + 438,
	79:  __ccgo_ts + 459,
	80:  __ccgo_ts + 473,
	81:  __ccgo_ts + 489,
	82:  __ccgo_ts + 505,
	83:  __ccgo_ts + 520,
	84:  __ccgo_ts + 533,
	85:  __ccgo_ts + 549,
	86:  __ccgo_ts + 566,
	87:  __ccgo_ts + 579,
	88:  __ccgo_ts + 596,
	89:  __ccgo_ts + 611,
	90:  __ccgo_ts + 625,
	91:  __ccgo_ts + 644,
	92:  __ccgo_ts + 459,
	93:  __ccgo_ts + 663,
	94:  __ccgo_ts + 675,
	95:  __ccgo_ts + 675,
	96:  __ccgo_ts + 675,
	97:  __ccgo_ts + 696,
	98:  __ccgo_ts + 711,
	99:  __ccgo_ts + 744,
	100: __ccgo_ts + 765,
	101: __ccgo_ts + 789,
	102: __ccgo_ts + 789,
	103: __ccgo_ts + 438,
	104: __ccgo_ts + 459,
	105: __ccgo_ts + 810,
	106: __ccgo_ts + 840,
	107: __ccgo_ts + 473,
	108: __ccgo_ts + 850,
	109: __ccgo_ts + 861,
	110: __ccgo_ts + 868,
	111: __ccgo_ts + 882,
	112: __ccgo_ts + 910,
	113: __ccgo_ts + 938,
	114: __ccgo_ts + 952,
	115: __ccgo_ts + 972,
	116: __ccgo_ts + 987,
	117: __ccgo_ts + 998,
	118: __ccgo_ts + 1014,
	119: __ccgo_ts + 1032,
	120: __ccgo_ts + 1051,
	121: __ccgo_ts + 1060,
	122: __ccgo_ts + 744,
	123: __ccgo_ts + 1085,
	124: __ccgo_ts + 765,
	125: __ccgo_ts + 1099,
	126: __ccgo_ts + 1109,
	127: __ccgo_ts + 1134,
	128: __ccgo_ts + 1152,
	129: __ccgo_ts + 1164,
	130: __ccgo_ts + 1175,
	131: __ccgo_ts + 1181,
	132: __ccgo_ts + 1199,
	133: __ccgo_ts + 1216,
	134: __ccgo_ts + 1224,
	135: __ccgo_ts + 1244,
	136: __ccgo_ts + 1258,
	137: __ccgo_ts + 1295,
	138: __ccgo_ts + 1332,
	139: __ccgo_ts + 1353,
	140: __ccgo_ts + 1372,
	141: __ccgo_ts + 1395,
	142: __ccgo_ts + 1420,
}

var ts_symbol_map = [143]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(sym_hash_bang_line),
	3:   uint16(anon_sym_return),
	4:   uint16(anon_sym_SEMI),
	5:   uint16(anon_sym_EQ),
	6:   uint16(anon_sym_COMMA),
	7:   uint16(anon_sym_COLON_COLON),
	8:   uint16(sym_break_statement),
	9:   uint16(anon_sym_goto),
	10:  uint16(anon_sym_do),
	11:  uint16(anon_sym_end),
	12:  uint16(anon_sym_while),
	13:  uint16(anon_sym_repeat),
	14:  uint16(anon_sym_until),
	15:  uint16(anon_sym_if),
	16:  uint16(anon_sym_then),
	17:  uint16(anon_sym_elseif),
	18:  uint16(anon_sym_else),
	19:  uint16(anon_sym_for),
	20:  uint16(anon_sym_in),
	21:  uint16(anon_sym_function),
	22:  uint16(anon_sym_local),
	23:  uint16(anon_sym_global),
	24:  uint16(anon_sym_DOT),
	25:  uint16(anon_sym_COLON),
	26:  uint16(anon_sym_STAR),
	27:  uint16(anon_sym_LT),
	28:  uint16(anon_sym_GT),
	29:  uint16(sym_nil),
	30:  uint16(sym_false),
	31:  uint16(sym_true),
	32:  uint16(sym_number),
	33:  uint16(anon_sym_DQUOTE),
	34:  uint16(anon_sym_SQUOTE),
	35:  uint16(aux_sym__doublequote_string_content_token1),
	36:  uint16(aux_sym__singlequote_string_content_token1),
	37:  uint16(sym_escape_sequence),
	38:  uint16(sym_vararg_expression),
	39:  uint16(anon_sym_LPAREN),
	40:  uint16(anon_sym_RPAREN),
	41:  uint16(anon_sym_LBRACK),
	42:  uint16(anon_sym_RBRACK),
	43:  uint16(anon_sym_LBRACE),
	44:  uint16(anon_sym_RBRACE),
	45:  uint16(anon_sym_or),
	46:  uint16(anon_sym_and),
	47:  uint16(anon_sym_LT_EQ),
	48:  uint16(anon_sym_EQ_EQ),
	49:  uint16(anon_sym_TILDE_EQ),
	50:  uint16(anon_sym_GT_EQ),
	51:  uint16(anon_sym_PIPE),
	52:  uint16(anon_sym_TILDE),
	53:  uint16(anon_sym_AMP),
	54:  uint16(anon_sym_LT_LT),
	55:  uint16(anon_sym_GT_GT),
	56:  uint16(anon_sym_PLUS),
	57:  uint16(anon_sym_DASH),
	58:  uint16(anon_sym_SLASH),
	59:  uint16(anon_sym_SLASH_SLASH),
	60:  uint16(anon_sym_PERCENT),
	61:  uint16(anon_sym_DOT_DOT),
	62:  uint16(anon_sym_CARET),
	63:  uint16(anon_sym_not),
	64:  uint16(anon_sym_POUND),
	65:  uint16(anon_sym_DASH_DASH),
	66:  uint16(sym__block_comment_content),
	67:  uint16(sym__block_comment_start),
	68:  uint16(sym__block_comment_content),
	69:  uint16(sym__block_comment_end),
	70:  uint16(sym__block_comment_start),
	71:  uint16(sym__block_string_content),
	72:  uint16(sym__block_comment_end),
	73:  uint16(sym_chunk),
	74:  uint16(sym__block),
	75:  uint16(sym_statement),
	76:  uint16(sym_return_statement),
	77:  uint16(sym_empty_statement),
	78:  uint16(sym_assignment_statement),
	79:  uint16(sym__variable_assignment_varlist),
	80:  uint16(sym__variable_assignment_explist),
	81:  uint16(sym_label_statement),
	82:  uint16(sym_goto_statement),
	83:  uint16(sym_do_statement),
	84:  uint16(sym_while_statement),
	85:  uint16(sym_repeat_statement),
	86:  uint16(sym_if_statement),
	87:  uint16(sym_elseif_statement),
	88:  uint16(sym_else_statement),
	89:  uint16(sym_for_statement),
	90:  uint16(sym_for_generic_clause),
	91:  uint16(sym_for_numeric_clause),
	92:  uint16(sym__variable_assignment_varlist),
	93:  uint16(sym_declaration),
	94:  uint16(sym_function_declaration),
	95:  uint16(sym_function_declaration),
	96:  uint16(sym_function_declaration),
	97:  uint16(sym__function_name),
	98:  uint16(sym__function_name_prefix_expression),
	99:  uint16(sym_dot_index_expression),
	100: uint16(sym_method_index_expression),
	101: uint16(sym_variable_declaration),
	102: uint16(sym_variable_declaration),
	103: uint16(sym_assignment_statement),
	104: uint16(sym__variable_assignment_varlist),
	105: uint16(sym__global_implicit_variable_declaration),
	106: uint16(sym__attrib),
	107: uint16(sym__variable_assignment_explist),
	108: uint16(sym_expression),
	109: uint16(sym_string),
	110: uint16(sym__quote_string),
	111: uint16(aux_sym__doublequote_string_content),
	112: uint16(aux_sym__singlequote_string_content),
	113: uint16(sym__block_string),
	114: uint16(sym_function_definition),
	115: uint16(sym__function_body),
	116: uint16(sym_parameters),
	117: uint16(sym__parameter_list),
	118: uint16(sym__vararg_parameter),
	119: uint16(sym__prefix_expression),
	120: uint16(sym_variable),
	121: uint16(sym_bracket_index_expression),
	122: uint16(sym_dot_index_expression),
	123: uint16(sym_function_call),
	124: uint16(sym_method_index_expression),
	125: uint16(sym_arguments),
	126: uint16(sym_parenthesized_expression),
	127: uint16(sym_table_constructor),
	128: uint16(sym__field_list),
	129: uint16(sym__field_sep),
	130: uint16(sym_field),
	131: uint16(sym_binary_expression),
	132: uint16(sym_unary_expression),
	133: uint16(sym_comment),
	134: uint16(sym__contextual_keyword),
	135: uint16(aux_sym_chunk_repeat1),
	136: uint16(aux_sym__variable_assignment_varlist_repeat1),
	137: uint16(aux_sym__variable_assignment_explist_repeat1),
	138: uint16(aux_sym_if_statement_repeat1),
	139: uint16(aux_sym__name_list_repeat1),
	140: uint16(aux_sym__att_name_list_repeat1),
	141: uint16(aux_sym__expression_list_repeat1),
	142: uint16(aux_sym__field_list_repeat1),
}

var ts_symbol_metadata = [143]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
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
		Fnamed:   libc.BoolUint8(1 != 0),
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
	},
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
		Fnamed:   libc.BoolUint8(1 != 0),
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
	},
	34: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	35: {},
	36: {},
	37: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	38: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	41: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	42: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	43: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	44: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	45: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	48: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	52: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	58: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	59: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	60: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	61: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	64: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	65: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	66: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	67: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	68: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	69: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	70: {
		Fvisible: libc.BoolUint8(1 != 0),
	},
	71: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	72: {
		Fvisible: libc.BoolUint8(1 != 0),
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
		Fnamed:     libc.BoolUint8(1 != 0),
		Fsupertype: libc.BoolUint8(1 != 0),
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
		Fnamed:     libc.BoolUint8(1 != 0),
		Fsupertype: libc.BoolUint8(1 != 0),
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
		Fnamed: libc.BoolUint8(1 != 0),
	},
	98: {
		Fnamed: libc.BoolUint8(1 != 0),
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
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	107: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	108: {
		Fnamed:     libc.BoolUint8(1 != 0),
		Fsupertype: libc.BoolUint8(1 != 0),
	},
	109: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	110: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	111: {},
	112: {},
	113: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	114: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	115: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	116: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	117: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	118: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	119: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	120: {
		Fnamed:     libc.BoolUint8(1 != 0),
		Fsupertype: libc.BoolUint8(1 != 0),
	},
	121: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	122: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	123: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	124: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	125: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	126: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	127: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	128: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	129: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	130: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	131: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	132: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	133: {
		Fvisible: libc.BoolUint8(1 != 0),
		Fnamed:   libc.BoolUint8(1 != 0),
	},
	134: {
		Fnamed: libc.BoolUint8(1 != 0),
	},
	135: {},
	136: {},
	137: {},
	138: {},
	139: {},
	140: {},
	141: {},
	142: {},
}

type ts_field_identifiers = int32

const field_alternative = 1
const field_arguments = 2
const field_attribute = 3
const field_body = 4
const field_clause = 5
const field_condition = 6
const field_consequence = 7
const field_content = 8
const field_end = 9
const field_field = 10
const field_global_declaration = 11
const field_left = 12
const field_local_declaration = 13
const field_method = 14
const field_name = 15
const field_operand = 16
const field_operator = 17
const field_parameters = 18
const field_right = 19
const field_start = 20
const field_step = 21
const field_table = 22
const field_value = 23

var ts_field_names = [24]uintptr{
	0:  libc.UintptrFromInt32(0),
	1:  __ccgo_ts + 1440,
	2:  __ccgo_ts + 1099,
	3:  __ccgo_ts + 840,
	4:  __ccgo_ts + 1452,
	5:  __ccgo_ts + 1457,
	6:  __ccgo_ts + 1464,
	7:  __ccgo_ts + 1474,
	8:  __ccgo_ts + 1486,
	9:  __ccgo_ts,
	10: __ccgo_ts + 1175,
	11: __ccgo_ts + 1494,
	12: __ccgo_ts + 1513,
	13: __ccgo_ts + 1518,
	14: __ccgo_ts + 1536,
	15: __ccgo_ts + 1543,
	16: __ccgo_ts + 1548,
	17: __ccgo_ts + 1556,
	18: __ccgo_ts + 987,
	19: __ccgo_ts + 1565,
	20: __ccgo_ts + 1571,
	21: __ccgo_ts + 1577,
	22: __ccgo_ts + 1582,
	23: __ccgo_ts + 1588,
}

var ts_field_map_slices = [70]TSMapSlice{
	1: {
		Flength: uint16(2),
	},
	2: {
		Findex:  uint16(2),
		Flength: uint16(4),
	},
	3: {
		Findex:  uint16(6),
		Flength: uint16(4),
	},
	4: {
		Findex:  uint16(10),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(11),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(12),
		Flength: uint16(2),
	},
	7: {
		Findex:  uint16(14),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(15),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(17),
		Flength: uint16(3),
	},
	10: {
		Findex:  uint16(20),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(22),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(24),
		Flength: uint16(1),
	},
	13: {
		Findex:  uint16(25),
		Flength: uint16(2),
	},
	14: {
		Findex:  uint16(27),
		Flength: uint16(2),
	},
	15: {
		Findex:  uint16(29),
		Flength: uint16(2),
	},
	16: {
		Findex:  uint16(31),
		Flength: uint16(3),
	},
	17: {
		Findex:  uint16(34),
		Flength: uint16(2),
	},
	18: {
		Findex:  uint16(36),
		Flength: uint16(2),
	},
	19: {
		Findex:  uint16(38),
		Flength: uint16(1),
	},
	20: {
		Findex:  uint16(39),
		Flength: uint16(2),
	},
	21: {
		Findex:  uint16(41),
		Flength: uint16(1),
	},
	22: {
		Findex:  uint16(42),
		Flength: uint16(1),
	},
	23: {
		Findex:  uint16(43),
		Flength: uint16(3),
	},
	24: {
		Findex:  uint16(46),
		Flength: uint16(2),
	},
	25: {
		Findex:  uint16(48),
		Flength: uint16(3),
	},
	26: {
		Findex:  uint16(51),
		Flength: uint16(2),
	},
	27: {
		Findex:  uint16(53),
		Flength: uint16(1),
	},
	28: {
		Findex:  uint16(54),
		Flength: uint16(3),
	},
	29: {
		Findex:  uint16(57),
		Flength: uint16(2),
	},
	30: {
		Findex:  uint16(59),
		Flength: uint16(2),
	},
	31: {
		Findex:  uint16(61),
		Flength: uint16(1),
	},
	32: {
		Findex:  uint16(62),
		Flength: uint16(2),
	},
	33: {
		Findex:  uint16(64),
		Flength: uint16(1),
	},
	34: {
		Findex:  uint16(65),
		Flength: uint16(1),
	},
	35: {
		Findex:  uint16(31),
		Flength: uint16(3),
	},
	36: {
		Findex:  uint16(66),
		Flength: uint16(3),
	},
	37: {
		Findex:  uint16(69),
		Flength: uint16(1),
	},
	38: {
		Findex:  uint16(70),
		Flength: uint16(2),
	},
	39: {
		Findex:  uint16(72),
		Flength: uint16(1),
	},
	40: {
		Findex:  uint16(73),
		Flength: uint16(1),
	},
	41: {
		Findex:  uint16(74),
		Flength: uint16(3),
	},
	42: {
		Findex:  uint16(77),
		Flength: uint16(4),
	},
	43: {
		Findex:  uint16(81),
		Flength: uint16(4),
	},
	44: {
		Findex:  uint16(85),
		Flength: uint16(4),
	},
	45: {
		Findex:  uint16(89),
		Flength: uint16(3),
	},
	46: {
		Findex:  uint16(92),
		Flength: uint16(4),
	},
	47: {
		Findex:  uint16(96),
		Flength: uint16(2),
	},
	48: {
		Findex:  uint16(98),
		Flength: uint16(1),
	},
	49: {
		Findex:  uint16(99),
		Flength: uint16(2),
	},
	50: {
		Findex:  uint16(101),
		Flength: uint16(2),
	},
	51: {
		Findex:  uint16(103),
		Flength: uint16(2),
	},
	52: {
		Findex:  uint16(105),
		Flength: uint16(2),
	},
	53: {
		Findex:  uint16(107),
		Flength: uint16(2),
	},
	54: {
		Findex:  uint16(109),
		Flength: uint16(2),
	},
	55: {
		Findex:  uint16(111),
		Flength: uint16(2),
	},
	56: {
		Findex:  uint16(113),
		Flength: uint16(2),
	},
	57: {
		Findex:  uint16(115),
		Flength: uint16(2),
	},
	58: {
		Findex:  uint16(117),
		Flength: uint16(5),
	},
	59: {
		Findex:  uint16(122),
		Flength: uint16(1),
	},
	60: {
		Findex:  uint16(123),
		Flength: uint16(2),
	},
	61: {
		Findex:  uint16(125),
		Flength: uint16(2),
	},
	62: {
		Findex:  uint16(127),
		Flength: uint16(3),
	},
	63: {
		Findex:  uint16(130),
		Flength: uint16(3),
	},
	64: {
		Findex:  uint16(133),
		Flength: uint16(3),
	},
	65: {
		Findex:  uint16(136),
		Flength: uint16(4),
	},
	66: {
		Findex:  uint16(140),
		Flength: uint16(3),
	},
	67: {
		Findex:  uint16(143),
		Flength: uint16(3),
	},
	68: {
		Findex:  uint16(146),
		Flength: uint16(4),
	},
	69: {
		Findex:  uint16(150),
		Flength: uint16(5),
	},
}

var ts_field_map_entries = [155]TSFieldMapEntry{
	0: {
		Ffield_id:  uint16(field_global_declaration),
		Finherited: libc.BoolUint8(1 != 0),
	},
	1: {
		Ffield_id:  uint16(field_local_declaration),
		Finherited: libc.BoolUint8(1 != 0),
	},
	2: {
		Ffield_id:  uint16(field_body),
		Finherited: libc.BoolUint8(1 != 0),
	},
	3: {
		Ffield_id: uint16(field_local_declaration),
	},
	4: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(1 != 0),
	},
	5: {
		Ffield_id:  uint16(field_parameters),
		Finherited: libc.BoolUint8(1 != 0),
	},
	6: {
		Ffield_id:  uint16(field_body),
		Finherited: libc.BoolUint8(1 != 0),
	},
	7: {
		Ffield_id: uint16(field_global_declaration),
	},
	8: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(1 != 0),
	},
	9: {
		Ffield_id:  uint16(field_parameters),
		Finherited: libc.BoolUint8(1 != 0),
	},
	10: {
		Ffield_id: uint16(field_local_declaration),
	},
	11: {
		Ffield_id: uint16(field_global_declaration),
	},
	12: {
		Ffield_id:  uint16(field_attribute),
		Finherited: libc.BoolUint8(1 != 0),
	},
	13: {
		Ffield_id: uint16(field_global_declaration),
	},
	14: {
		Ffield_id: uint16(field_name),
	},
	15: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(1),
	},
	16: {
		Ffield_id: uint16(field_start),
	},
	17: {
		Ffield_id:  uint16(field_content),
		Finherited: libc.BoolUint8(1 != 0),
	},
	18: {
		Ffield_id:  uint16(field_end),
		Finherited: libc.BoolUint8(1 != 0),
	},
	19: {
		Ffield_id:  uint16(field_start),
		Finherited: libc.BoolUint8(1 != 0),
	},
	20: {
		Ffield_id:  uint16(field_field),
		Finherited: libc.BoolUint8(1 != 0),
	},
	21: {
		Ffield_id:  uint16(field_table),
		Finherited: libc.BoolUint8(1 != 0),
	},
	22: {
		Ffield_id:  uint16(field_method),
		Finherited: libc.BoolUint8(1 != 0),
	},
	23: {
		Ffield_id:  uint16(field_table),
		Finherited: libc.BoolUint8(1 != 0),
	},
	24: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	25: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	26: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	27: {
		Ffield_id:    uint16(field_arguments),
		Fchild_index: uint8(1),
	},
	28: {
		Ffield_id: uint16(field_name),
	},
	29: {
		Ffield_id: uint16(field_name),
	},
	30: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	31: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(1),
	},
	32: {
		Ffield_id:    uint16(field_end),
		Fchild_index: uint8(2),
	},
	33: {
		Ffield_id: uint16(field_start),
	},
	34: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	35: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	36: {
		Ffield_id:    uint16(field_end),
		Fchild_index: uint8(1),
	},
	37: {
		Ffield_id: uint16(field_start),
	},
	38: {
		Ffield_id: uint16(field_value),
	},
	39: {
		Ffield_id:    uint16(field_operand),
		Fchild_index: uint8(1),
	},
	40: {
		Ffield_id: uint16(field_operator),
	},
	41: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
	},
	42: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(2),
	},
	43: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	44: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	45: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	46: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(1),
	},
	47: {
		Ffield_id: uint16(field_name),
	},
	48: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	49: {
		Ffield_id: uint16(field_name),
	},
	50: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	51: {
		Ffield_id: uint16(field_attribute),
	},
	52: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	53: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(1),
	},
	54: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(1 != 0),
	},
	55: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	56: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	57: {
		Ffield_id:    uint16(field_field),
		Fchild_index: uint8(2),
	},
	58: {
		Ffield_id: uint16(field_table),
	},
	59: {
		Ffield_id:    uint16(field_method),
		Fchild_index: uint8(2),
	},
	60: {
		Ffield_id: uint16(field_table),
	},
	61: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	62: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(1 != 0),
	},
	63: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	64: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(1 != 0),
	},
	65: {
		Ffield_id: uint16(field_parameters),
	},
	66: {
		Ffield_id: uint16(field_left),
	},
	67: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	68: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	69: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	70: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
	},
	71: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(3),
	},
	72: {
		Ffield_id: uint16(field_alternative),
	},
	73: {
		Ffield_id:    uint16(field_clause),
		Fchild_index: uint8(1),
	},
	74: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	75: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	76: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	77: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(1),
	},
	78: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	79: {
		Ffield_id: uint16(field_name),
	},
	80: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	81: {
		Ffield_id:  uint16(field_attribute),
		Finherited: libc.BoolUint8(1 != 0),
	},
	82: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	83: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(1 != 0),
	},
	84: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	85: {
		Ffield_id:  uint16(field_attribute),
		Finherited: libc.BoolUint8(1 != 0),
	},
	86: {
		Ffield_id:  uint16(field_name),
		Finherited: libc.BoolUint8(1 != 0),
	},
	87: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	88: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	89: {
		Ffield_id: uint16(field_attribute),
	},
	90: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(2),
	},
	91: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	92: {
		Ffield_id: uint16(field_attribute),
	},
	93: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	94: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	95: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	96: {
		Ffield_id: uint16(field_value),
	},
	97: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	98: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	99: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(1),
	},
	100: {
		Ffield_id: uint16(field_parameters),
	},
	101: {
		Ffield_id: uint16(field_name),
	},
	102: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(2),
	},
	103: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(3),
	},
	104: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	105: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	106: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(3),
	},
	107: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(3),
	},
	108: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	109: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	110: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	111: {
		Ffield_id:  uint16(field_alternative),
		Finherited: libc.BoolUint8(1 != 0),
	},
	112: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	113: {
		Ffield_id:    uint16(field_body),
		Fchild_index: uint8(3),
	},
	114: {
		Ffield_id:    uint16(field_clause),
		Fchild_index: uint8(1),
	},
	115: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(2),
	},
	116: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	117: {
		Ffield_id: uint16(field_attribute),
	},
	118: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(2),
	},
	119: {
		Ffield_id:    uint16(field_attribute),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	120: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	121: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	122: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
	},
	123: {
		Ffield_id:  uint16(field_value),
		Finherited: libc.BoolUint8(1 != 0),
	},
	124: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	125: {
		Ffield_id: uint16(field_name),
	},
	126: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	127: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
	},
	128: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	129: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(3),
	},
	130: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	131: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	132: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(3),
	},
	133: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	134: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
	},
	135: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	136: {
		Ffield_id:    uint16(field_end),
		Fchild_index: uint8(4),
	},
	137: {
		Ffield_id: uint16(field_name),
	},
	138: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	139: {
		Ffield_id:    uint16(field_start),
		Fchild_index: uint8(2),
	},
	140: {
		Ffield_id: uint16(field_name),
	},
	141: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	142: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	143: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	144: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(3),
	},
	145: {
		Ffield_id:    uint16(field_value),
		Fchild_index: uint8(4),
	},
	146: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(1 != 0),
	},
	147: {
		Ffield_id:    uint16(field_alternative),
		Fchild_index: uint8(5),
	},
	148: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	149: {
		Ffield_id:    uint16(field_consequence),
		Fchild_index: uint8(3),
	},
	150: {
		Ffield_id:    uint16(field_end),
		Fchild_index: uint8(4),
	},
	151: {
		Ffield_id: uint16(field_name),
	},
	152: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	153: {
		Ffield_id:    uint16(field_start),
		Fchild_index: uint8(2),
	},
	154: {
		Ffield_id:    uint16(field_step),
		Fchild_index: uint8(6),
	},
}

var ts_alias_sequences = [70][7]TSSymbol{
	0: {},
	35: {
		1: uint16(sym__block_string_content),
	},
}

var ts_non_terminal_alias_map = [9]uint16_t{
	0: uint16(aux_sym__doublequote_string_content),
	1: uint16(2),
	2: uint16(aux_sym__doublequote_string_content),
	3: uint16(sym__block_string_content),
	4: uint16(aux_sym__singlequote_string_content),
	5: uint16(2),
	6: uint16(aux_sym__singlequote_string_content),
	7: uint16(sym__block_string_content),
}

var ts_primary_state_ids = [282]TSStateId{
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
	24:  uint16(18),
	25:  uint16(25),
	26:  uint16(26),
	27:  uint16(27),
	28:  uint16(28),
	29:  uint16(29),
	30:  uint16(19),
	31:  uint16(6),
	32:  uint16(32),
	33:  uint16(33),
	34:  uint16(34),
	35:  uint16(35),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(17),
	40:  uint16(40),
	41:  uint16(14),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(15),
	45:  uint16(16),
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
	62:  uint16(61),
	63:  uint16(60),
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
	134: uint16(134),
	135: uint16(135),
	136: uint16(136),
	137: uint16(137),
	138: uint16(138),
	139: uint16(139),
	140: uint16(140),
	141: uint16(141),
	142: uint16(142),
	143: uint16(143),
	144: uint16(144),
	145: uint16(145),
	146: uint16(146),
	147: uint16(147),
	148: uint16(148),
	149: uint16(149),
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
	163: uint16(163),
	164: uint16(164),
	165: uint16(165),
	166: uint16(166),
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
	178: uint16(178),
	179: uint16(179),
	180: uint16(180),
	181: uint16(181),
	182: uint16(182),
	183: uint16(183),
	184: uint16(184),
	185: uint16(185),
	186: uint16(186),
	187: uint16(187),
	188: uint16(188),
	189: uint16(189),
	190: uint16(190),
	191: uint16(191),
	192: uint16(192),
	193: uint16(193),
	194: uint16(194),
	195: uint16(195),
	196: uint16(196),
	197: uint16(193),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(201),
	202: uint16(202),
	203: uint16(203),
	204: uint16(202),
	205: uint16(203),
	206: uint16(196),
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
	218: uint16(218),
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
	234: uint16(234),
	235: uint16(235),
	236: uint16(236),
	237: uint16(237),
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
	273: uint16(273),
	274: uint16(232),
	275: uint16(275),
	276: uint16(276),
	277: uint16(242),
	278: uint16(275),
	279: uint16(279),
	280: uint16(280),
	281: uint16(281),
}

var ts_supertype_symbols = [4]TSSymbol{
	0: uint16(sym_declaration),
	1: uint16(sym_expression),
	2: uint16(sym_statement),
	3: uint16(sym_variable),
}

var ts_supertype_map_slices = [121]TSMapSlice{
	75: {
		Findex:  uint16(19),
		Flength: uint16(12),
	},
	93: {
		Flength: uint16(6),
	},
	108: {
		Findex:  uint16(6),
		Flength: uint16(13),
	},
	120: {
		Findex:  uint16(31),
		Flength: uint16(4),
	},
}

var ts_supertype_map_entries = [35]TSSymbol{
	0:  uint16(sym__global_function_declaration),
	1:  uint16(sym__global_implicit_variable_declaration),
	2:  uint16(sym__global_variable_declaration),
	3:  uint16(sym__local_function_declaration),
	4:  uint16(sym_function_declaration),
	5:  uint16(sym_variable_declaration),
	6:  uint16(sym_binary_expression),
	7:  uint16(sym_false),
	8:  uint16(sym_function_call),
	9:  uint16(sym_function_definition),
	10: uint16(sym_nil),
	11: uint16(sym_number),
	12: uint16(sym_parenthesized_expression),
	13: uint16(sym_string),
	14: uint16(sym_table_constructor),
	15: uint16(sym_true),
	16: uint16(sym_unary_expression),
	17: uint16(sym_vararg_expression),
	18: uint16(sym_variable),
	19: uint16(sym_assignment_statement),
	20: uint16(sym_break_statement),
	21: uint16(sym_declaration),
	22: uint16(sym_do_statement),
	23: uint16(sym_empty_statement),
	24: uint16(sym_for_statement),
	25: uint16(sym_function_call),
	26: uint16(sym_goto_statement),
	27: uint16(sym_if_statement),
	28: uint16(sym_label_statement),
	29: uint16(sym_repeat_statement),
	30: uint16(sym_while_statement),
	31: uint16(anon_sym_global),
	32: uint16(sym_bracket_index_expression),
	33: uint16(sym_dot_index_expression),
	34: uint16(sym_identifier),
}

var sym_escape_sequence_character_set_1 = [12]TSCharacterRange{
	0: {
		Fstart: int32('\n'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32('"'),
		Fend:   int32('"'),
	},
	2: {
		Fstart: int32('\''),
		Fend:   int32('\''),
	},
	3: {
		Fstart: int32('0'),
		Fend:   int32('9'),
	},
	4: {
		Fstart: int32('\\'),
		Fend:   int32('\\'),
	},
	5: {
		Fstart: int32('a'),
		Fend:   int32('b'),
	},
	6: {
		Fstart: int32('f'),
		Fend:   int32('f'),
	},
	7: {
		Fstart: int32('n'),
		Fend:   int32('n'),
	},
	8: {
		Fstart: int32('r'),
		Fend:   int32('r'),
	},
	9: {
		Fstart: int32('t'),
		Fend:   int32('v'),
	},
	10: {
		Fstart: int32('x'),
		Fend:   int32('x'),
	},
	11: {
		Fstart: int32('z'),
		Fend:   int32('z'),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4, i5, i6 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, i5, i6, lookahead, result, skip
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
			state = uint16(21)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(108)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(18)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(36)
			goto next_state
		}
		if lookahead > int32(' ') && (lookahead < int32('{') || int32(0x9f) < lookahead) {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('"') {
			state = uint16(42)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(6)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(44)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\'') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(6)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(47)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('.') {
			state = uint16(4)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('.') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('.') {
			state = uint16(15)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('u') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('z') {
			state = uint16(51)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(53)
			goto next_state
		}
		if set_contains(tls, uintptr(unsafe.Pointer(&sym_escape_sequence_character_set_1)), uint32(12), lookahead) != 0 {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('{') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('}') {
			state = uint16(50)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(13)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(13):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(14):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(15):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(16):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(17):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(18):
		if eof != 0 {
			state = uint16(21)
			goto next_state
		}
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(104)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(18)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(36)
			goto next_state
		}
		if lookahead > int32(' ') && (lookahead < int32('[') || int32('^') < lookahead) && (lookahead < int32('{') || int32(0x9f) < lookahead) {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(19):
		if eof != 0 {
			state = uint16(21)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(19)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(36)
			goto next_state
		}
		if lookahead > int32(' ') && (lookahead < int32('%') || int32('>') < lookahead) && (lookahead < int32('[') || int32('^') < lookahead) && (lookahead < int32('{') || int32(0x9f) < lookahead) {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(20):
		if eof != 0 {
			state = uint16(21)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(100)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead > int32(' ') && (lookahead < int32('%') || int32('>') < lookahead) && (lookahead < int32('[') || int32('^') < lookahead) && (lookahead < int32('{') || int32(0x9f) < lookahead) {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(0)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_hash_bang_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(23):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(24):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(28):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(76)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(29):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('<') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('=') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(34):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token4[i4]) == lookahead {
				state = map_token4[i4+uint32(1)]
				goto next_state
			}
			goto _5
		_5:
			;
			i4 = i4 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token5[i5]) == lookahead {
				state = map_token5[i5+uint32(1)]
				goto next_state
			}
			goto _6
		_6:
			;
			i5 = i5 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i6]) == lookahead {
				state = map_token6[i6+uint32(1)]
				goto next_state
			}
			goto _7
		_7:
			;
			i6 = i6 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(38):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('L') || lookahead == int32('l') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('U') || lookahead == int32('u') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('0') || lookahead == int32('1') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('P') || lookahead == int32('p') {
			state = uint16(9)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('I') || lookahead == int32('i') {
			state = uint16(34)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__doublequote_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(45)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(44)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__doublequote_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(46)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__doublequote_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__singlequote_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(48)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(47)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(48):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__singlequote_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(49)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__singlequote_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(50):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(53):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_vararg_expression)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('/') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CARET)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead > int32(' ') && lookahead != int32('"') && lookahead != int32('#') && (lookahead < int32('%') || int32('/') < lookahead) && (lookahead < int32(':') || int32('>') < lookahead) && (lookahead < int32('[') || int32('^') < lookahead) && (lookahead < int32('{') || int32(0x9f) < lookahead) {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(0x0b) || lookahead == int32('\f') || lookahead == int32(' ') {
			state = uint16(82)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(81)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\r') {
			state = uint16(84)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var map_token = [54]uint16_t{
	0:  uint16('"'),
	1:  uint16(42),
	2:  uint16('#'),
	3:  uint16(78),
	4:  uint16('%'),
	5:  uint16(75),
	6:  uint16('&'),
	7:  uint16(68),
	8:  uint16('\''),
	9:  uint16(43),
	10: uint16('('),
	11: uint16(55),
	12: uint16(')'),
	13: uint16(56),
	14: uint16('*'),
	15: uint16(30),
	16: uint16('+'),
	17: uint16(71),
	18: uint16(','),
	19: uint16(25),
	20: uint16('-'),
	21: uint16(72),
	22: uint16('.'),
	23: uint16(28),
	24: uint16('/'),
	25: uint16(73),
	26: uint16('0'),
	27: uint16(35),
	28: uint16(':'),
	29: uint16(29),
	30: uint16(';'),
	31: uint16(23),
	32: uint16('<'),
	33: uint16(31),
	34: uint16('='),
	35: uint16(24),
	36: uint16('>'),
	37: uint16(33),
	38: uint16('['),
	39: uint16(57),
	40: uint16('\\'),
	41: uint16(6),
	42: uint16(']'),
	43: uint16(58),
	44: uint16('^'),
	45: uint16(77),
	46: uint16('{'),
	47: uint16(59),
	48: uint16('|'),
	49: uint16(65),
	50: uint16('}'),
	51: uint16(60),
	52: uint16('~'),
	53: uint16(67),
}

var map_token1 = [52]uint16_t{
	0:  uint16('"'),
	1:  uint16(42),
	2:  uint16('#'),
	3:  uint16(78),
	4:  uint16('%'),
	5:  uint16(75),
	6:  uint16('&'),
	7:  uint16(68),
	8:  uint16('\''),
	9:  uint16(43),
	10: uint16('('),
	11: uint16(55),
	12: uint16(')'),
	13: uint16(56),
	14: uint16('*'),
	15: uint16(30),
	16: uint16('+'),
	17: uint16(71),
	18: uint16(','),
	19: uint16(25),
	20: uint16('-'),
	21: uint16(72),
	22: uint16('.'),
	23: uint16(28),
	24: uint16('/'),
	25: uint16(73),
	26: uint16('0'),
	27: uint16(35),
	28: uint16(':'),
	29: uint16(29),
	30: uint16(';'),
	31: uint16(23),
	32: uint16('<'),
	33: uint16(31),
	34: uint16('='),
	35: uint16(24),
	36: uint16('>'),
	37: uint16(33),
	38: uint16('['),
	39: uint16(57),
	40: uint16(']'),
	41: uint16(58),
	42: uint16('^'),
	43: uint16(77),
	44: uint16('{'),
	45: uint16(59),
	46: uint16('|'),
	47: uint16(65),
	48: uint16('}'),
	49: uint16(60),
	50: uint16('~'),
	51: uint16(67),
}

var map_token2 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(42),
	2:  uint16('#'),
	3:  uint16(78),
	4:  uint16('\''),
	5:  uint16(43),
	6:  uint16('('),
	7:  uint16(55),
	8:  uint16(')'),
	9:  uint16(56),
	10: uint16('-'),
	11: uint16(72),
	12: uint16('.'),
	13: uint16(3),
	14: uint16('0'),
	15: uint16(35),
	16: uint16(';'),
	17: uint16(23),
	18: uint16('>'),
	19: uint16(32),
	20: uint16('['),
	21: uint16(57),
	22: uint16('{'),
	23: uint16(59),
	24: uint16('}'),
	25: uint16(60),
	26: uint16('~'),
	27: uint16(66),
}

var map_token3 = [50]uint16_t{
	0:  uint16('"'),
	1:  uint16(42),
	2:  uint16('#'),
	3:  uint16(22),
	4:  uint16('%'),
	5:  uint16(75),
	6:  uint16('&'),
	7:  uint16(68),
	8:  uint16('\''),
	9:  uint16(43),
	10: uint16('('),
	11: uint16(55),
	12: uint16(')'),
	13: uint16(56),
	14: uint16('*'),
	15: uint16(30),
	16: uint16('+'),
	17: uint16(71),
	18: uint16(','),
	19: uint16(25),
	20: uint16('-'),
	21: uint16(72),
	22: uint16('.'),
	23: uint16(27),
	24: uint16('/'),
	25: uint16(73),
	26: uint16(':'),
	27: uint16(29),
	28: uint16(';'),
	29: uint16(23),
	30: uint16('<'),
	31: uint16(31),
	32: uint16('='),
	33: uint16(24),
	34: uint16('>'),
	35: uint16(33),
	36: uint16('['),
	37: uint16(57),
	38: uint16(']'),
	39: uint16(58),
	40: uint16('^'),
	41: uint16(77),
	42: uint16('{'),
	43: uint16(59),
	44: uint16('|'),
	45: uint16(65),
	46: uint16('}'),
	47: uint16(60),
	48: uint16('~'),
	49: uint16(67),
}

var map_token4 = [26]uint16_t{
	0:  uint16('.'),
	1:  uint16(38),
	2:  uint16('B'),
	3:  uint16(12),
	4:  uint16('b'),
	5:  uint16(12),
	6:  uint16('E'),
	7:  uint16(9),
	8:  uint16('e'),
	9:  uint16(9),
	10: uint16('I'),
	11: uint16(34),
	12: uint16('i'),
	13: uint16(34),
	14: uint16('L'),
	15: uint16(10),
	16: uint16('l'),
	17: uint16(10),
	18: uint16('U'),
	19: uint16(11),
	20: uint16('u'),
	21: uint16(11),
	22: uint16('X'),
	23: uint16(5),
	24: uint16('x'),
	25: uint16(5),
}

var map_token5 = [18]uint16_t{
	0:  uint16('.'),
	1:  uint16(38),
	2:  uint16('E'),
	3:  uint16(9),
	4:  uint16('e'),
	5:  uint16(9),
	6:  uint16('I'),
	7:  uint16(34),
	8:  uint16('i'),
	9:  uint16(34),
	10: uint16('L'),
	11: uint16(10),
	12: uint16('l'),
	13: uint16(10),
	14: uint16('U'),
	15: uint16(11),
	16: uint16('u'),
	17: uint16(11),
}

var map_token6 = [18]uint16_t{
	0:  uint16('.'),
	1:  uint16(40),
	2:  uint16('I'),
	3:  uint16(34),
	4:  uint16('i'),
	5:  uint16(34),
	6:  uint16('L'),
	7:  uint16(10),
	8:  uint16('l'),
	9:  uint16(10),
	10: uint16('P'),
	11: uint16(9),
	12: uint16('p'),
	13: uint16(9),
	14: uint16('U'),
	15: uint16(11),
	16: uint16('u'),
	17: uint16(11),
}

func ts_lex_keywords(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
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
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token7[i]) == lookahead {
				state = map_token7[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('n') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('r') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('o') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('l') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('a') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('l') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('f') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('o') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('i') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('r') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('e') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('h') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('n') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('h') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('d') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(17):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_do)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(18):
		if lookahead == int32('s') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('d') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('l') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('r') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('n') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('o') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('t') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(26):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_in)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		if lookahead == int32('c') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('l') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('t') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_or)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(31):
		if lookahead == int32('p') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('e') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('u') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('t') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('i') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(36):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_and)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		if lookahead == int32('a') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('e') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(39):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		if lookahead == int32('s') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_for)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(42):
		if lookahead == int32('c') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('b') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('o') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('a') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_nil)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_not)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		if lookahead == int32('e') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('u') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('n') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('e') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('i') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('l') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('k') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('e') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('t') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('a') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_goto)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		if lookahead == int32('l') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('a') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('r') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_then)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		if lookahead == int32('l') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('e') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_break_statement)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		if lookahead == int32('f') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		if lookahead == int32('i') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('l') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_local)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		if lookahead == int32('t') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('n') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_until)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_while)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_elseif)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		if lookahead == int32('o') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_global)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_repeat)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_return)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		if lookahead == int32('n') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_function)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(0 != 0)
	}
	return r
}

var map_token7 = [28]uint16_t{
	0:  uint16('a'),
	1:  uint16(1),
	2:  uint16('b'),
	3:  uint16(2),
	4:  uint16('d'),
	5:  uint16(3),
	6:  uint16('e'),
	7:  uint16(4),
	8:  uint16('f'),
	9:  uint16(5),
	10: uint16('g'),
	11: uint16(6),
	12: uint16('i'),
	13: uint16(7),
	14: uint16('l'),
	15: uint16(8),
	16: uint16('n'),
	17: uint16(9),
	18: uint16('o'),
	19: uint16(10),
	20: uint16('r'),
	21: uint16(11),
	22: uint16('t'),
	23: uint16(12),
	24: uint16('u'),
	25: uint16(13),
	26: uint16('w'),
	27: uint16(14),
}

var ts_lex_modes = [282]TSLexerMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	3: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	4: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	5: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	6: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	7: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	8: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	9: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	10: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	11: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	12: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	13: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	14: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	15: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	16: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	17: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	18: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	19: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	20: {
		Fexternal_lex_state: uint16(2),
	},
	21: {
		Fexternal_lex_state: uint16(2),
	},
	22: {
		Fexternal_lex_state: uint16(2),
	},
	23: {
		Fexternal_lex_state: uint16(2),
	},
	24: {
		Fexternal_lex_state: uint16(2),
	},
	25: {
		Fexternal_lex_state: uint16(2),
	},
	26: {
		Fexternal_lex_state: uint16(2),
	},
	27: {
		Fexternal_lex_state: uint16(2),
	},
	28: {
		Fexternal_lex_state: uint16(2),
	},
	29: {
		Fexternal_lex_state: uint16(2),
	},
	30: {
		Fexternal_lex_state: uint16(2),
	},
	31: {
		Fexternal_lex_state: uint16(2),
	},
	32: {
		Fexternal_lex_state: uint16(2),
	},
	33: {
		Fexternal_lex_state: uint16(2),
	},
	34: {
		Fexternal_lex_state: uint16(2),
	},
	35: {
		Fexternal_lex_state: uint16(2),
	},
	36: {
		Fexternal_lex_state: uint16(2),
	},
	37: {
		Fexternal_lex_state: uint16(2),
	},
	38: {
		Fexternal_lex_state: uint16(2),
	},
	39: {
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Fexternal_lex_state: uint16(2),
	},
	49: {
		Fexternal_lex_state: uint16(2),
	},
	50: {
		Fexternal_lex_state: uint16(2),
	},
	51: {
		Fexternal_lex_state: uint16(2),
	},
	52: {
		Fexternal_lex_state: uint16(2),
	},
	53: {
		Fexternal_lex_state: uint16(2),
	},
	54: {
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Fexternal_lex_state: uint16(2),
	},
	56: {
		Fexternal_lex_state: uint16(2),
	},
	57: {
		Fexternal_lex_state: uint16(2),
	},
	58: {
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Fexternal_lex_state: uint16(2),
	},
	60: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	61: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	62: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	63: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	64: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	65: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	66: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	67: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	68: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	69: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	70: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	71: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	72: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	73: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	74: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	75: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	76: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	77: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	78: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	79: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	80: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	81: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	82: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	83: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	84: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	85: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	86: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	87: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	88: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	89: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	90: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	91: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	92: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	93: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	94: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	95: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	96: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	97: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	98: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	99: {
		Fexternal_lex_state: uint16(2),
	},
	100: {
		Fexternal_lex_state: uint16(2),
	},
	101: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	102: {
		Fexternal_lex_state: uint16(2),
	},
	103: {
		Fexternal_lex_state: uint16(2),
	},
	104: {
		Fexternal_lex_state: uint16(2),
	},
	105: {
		Fexternal_lex_state: uint16(2),
	},
	106: {
		Fexternal_lex_state: uint16(2),
	},
	107: {
		Fexternal_lex_state: uint16(2),
	},
	108: {
		Fexternal_lex_state: uint16(2),
	},
	109: {
		Fexternal_lex_state: uint16(2),
	},
	110: {
		Fexternal_lex_state: uint16(2),
	},
	111: {
		Fexternal_lex_state: uint16(2),
	},
	112: {
		Fexternal_lex_state: uint16(2),
	},
	113: {
		Fexternal_lex_state: uint16(2),
	},
	114: {
		Fexternal_lex_state: uint16(2),
	},
	115: {
		Fexternal_lex_state: uint16(2),
	},
	116: {
		Fexternal_lex_state: uint16(2),
	},
	117: {
		Fexternal_lex_state: uint16(2),
	},
	118: {
		Fexternal_lex_state: uint16(2),
	},
	119: {
		Fexternal_lex_state: uint16(2),
	},
	120: {
		Fexternal_lex_state: uint16(2),
	},
	121: {
		Fexternal_lex_state: uint16(2),
	},
	122: {
		Fexternal_lex_state: uint16(2),
	},
	123: {
		Fexternal_lex_state: uint16(2),
	},
	124: {
		Fexternal_lex_state: uint16(2),
	},
	125: {
		Fexternal_lex_state: uint16(2),
	},
	126: {
		Fexternal_lex_state: uint16(2),
	},
	127: {
		Fexternal_lex_state: uint16(2),
	},
	128: {
		Fexternal_lex_state: uint16(2),
	},
	129: {
		Fexternal_lex_state: uint16(2),
	},
	130: {
		Fexternal_lex_state: uint16(2),
	},
	131: {
		Fexternal_lex_state: uint16(2),
	},
	132: {
		Fexternal_lex_state: uint16(2),
	},
	133: {
		Fexternal_lex_state: uint16(2),
	},
	134: {
		Fexternal_lex_state: uint16(2),
	},
	135: {
		Fexternal_lex_state: uint16(2),
	},
	136: {
		Fexternal_lex_state: uint16(2),
	},
	137: {
		Fexternal_lex_state: uint16(2),
	},
	138: {
		Fexternal_lex_state: uint16(2),
	},
	139: {
		Fexternal_lex_state: uint16(2),
	},
	140: {
		Fexternal_lex_state: uint16(2),
	},
	141: {
		Fexternal_lex_state: uint16(2),
	},
	142: {
		Fexternal_lex_state: uint16(2),
	},
	143: {
		Fexternal_lex_state: uint16(2),
	},
	144: {
		Fexternal_lex_state: uint16(2),
	},
	145: {
		Fexternal_lex_state: uint16(2),
	},
	146: {
		Fexternal_lex_state: uint16(2),
	},
	147: {
		Fexternal_lex_state: uint16(2),
	},
	148: {
		Fexternal_lex_state: uint16(2),
	},
	149: {
		Fexternal_lex_state: uint16(2),
	},
	150: {
		Fexternal_lex_state: uint16(2),
	},
	151: {
		Fexternal_lex_state: uint16(2),
	},
	152: {
		Fexternal_lex_state: uint16(2),
	},
	153: {
		Fexternal_lex_state: uint16(2),
	},
	154: {
		Fexternal_lex_state: uint16(2),
	},
	155: {
		Fexternal_lex_state: uint16(2),
	},
	156: {
		Fexternal_lex_state: uint16(2),
	},
	157: {
		Fexternal_lex_state: uint16(2),
	},
	158: {
		Fexternal_lex_state: uint16(2),
	},
	159: {
		Fexternal_lex_state: uint16(2),
	},
	160: {
		Fexternal_lex_state: uint16(2),
	},
	161: {
		Fexternal_lex_state: uint16(2),
	},
	162: {
		Fexternal_lex_state: uint16(2),
	},
	163: {
		Fexternal_lex_state: uint16(2),
	},
	164: {
		Fexternal_lex_state: uint16(2),
	},
	165: {
		Fexternal_lex_state: uint16(2),
	},
	166: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(3),
	},
	167: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	168: {
		Fexternal_lex_state: uint16(2),
	},
	169: {
		Fexternal_lex_state: uint16(2),
	},
	170: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	171: {
		Fexternal_lex_state: uint16(2),
	},
	172: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	173: {
		Fexternal_lex_state: uint16(2),
	},
	174: {
		Fexternal_lex_state: uint16(3),
	},
	175: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	176: {
		Fexternal_lex_state: uint16(2),
	},
	177: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(3),
	},
	178: {
		Fexternal_lex_state: uint16(2),
	},
	179: {
		Fexternal_lex_state: uint16(2),
	},
	180: {
		Fexternal_lex_state: uint16(2),
	},
	181: {
		Fexternal_lex_state: uint16(2),
	},
	182: {
		Fexternal_lex_state: uint16(2),
	},
	183: {
		Fexternal_lex_state: uint16(2),
	},
	184: {
		Fexternal_lex_state: uint16(2),
	},
	185: {
		Fexternal_lex_state: uint16(2),
	},
	186: {
		Fexternal_lex_state: uint16(2),
	},
	187: {
		Fexternal_lex_state: uint16(3),
	},
	188: {
		Fexternal_lex_state: uint16(2),
	},
	189: {
		Fexternal_lex_state: uint16(2),
	},
	190: {
		Fexternal_lex_state: uint16(2),
	},
	191: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	192: {
		Fexternal_lex_state: uint16(2),
	},
	193: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	194: {
		Fexternal_lex_state: uint16(2),
	},
	195: {
		Fexternal_lex_state: uint16(2),
	},
	196: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	197: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	198: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	199: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	200: {
		Fexternal_lex_state: uint16(2),
	},
	201: {
		Fexternal_lex_state: uint16(2),
	},
	202: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	203: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	204: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	205: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	206: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	207: {
		Fexternal_lex_state: uint16(2),
	},
	208: {
		Flex_state:          uint16(2),
		Fexternal_lex_state: uint16(2),
	},
	209: {
		Fexternal_lex_state: uint16(2),
	},
	210: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(2),
	},
	211: {
		Fexternal_lex_state: uint16(2),
	},
	212: {
		Fexternal_lex_state: uint16(2),
	},
	213: {
		Fexternal_lex_state: uint16(2),
	},
	214: {
		Flex_state:          uint16(19),
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
		Fexternal_lex_state: uint16(2),
	},
	219: {
		Fexternal_lex_state: uint16(2),
	},
	220: {
		Fexternal_lex_state: uint16(2),
	},
	221: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	222: {
		Fexternal_lex_state: uint16(2),
	},
	223: {
		Fexternal_lex_state: uint16(2),
	},
	224: {
		Fexternal_lex_state: uint16(2),
	},
	225: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(2),
	},
	226: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(2),
	},
	227: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(2),
	},
	228: {
		Flex_state:          uint16(20),
		Fexternal_lex_state: uint16(2),
	},
	229: {
		Fexternal_lex_state: uint16(2),
	},
	230: {
		Fexternal_lex_state: uint16(2),
	},
	231: {
		Fexternal_lex_state: uint16(2),
	},
	232: {
		Fexternal_lex_state: uint16(4),
	},
	233: {
		Fexternal_lex_state: uint16(2),
	},
	234: {
		Fexternal_lex_state: uint16(2),
	},
	235: {
		Fexternal_lex_state: uint16(2),
	},
	236: {
		Fexternal_lex_state: uint16(2),
	},
	237: {
		Flex_state:          uint16(19),
		Fexternal_lex_state: uint16(2),
	},
	238: {
		Fexternal_lex_state: uint16(2),
	},
	239: {
		Fexternal_lex_state: uint16(2),
	},
	240: {
		Fexternal_lex_state: uint16(2),
	},
	241: {
		Fexternal_lex_state: uint16(2),
	},
	242: {
		Fexternal_lex_state: uint16(2),
	},
	243: {
		Fexternal_lex_state: uint16(2),
	},
	244: {
		Fexternal_lex_state: uint16(2),
	},
	245: {
		Fexternal_lex_state: uint16(2),
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
		Fexternal_lex_state: uint16(2),
	},
	251: {
		Fexternal_lex_state: uint16(2),
	},
	252: {
		Fexternal_lex_state: uint16(5),
	},
	253: {
		Fexternal_lex_state: uint16(6),
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
		Fexternal_lex_state: uint16(2),
	},
	264: {
		Flex_state:          uint16(82),
		Fexternal_lex_state: uint16(2),
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
		Fexternal_lex_state: uint16(2),
	},
	270: {
		Fexternal_lex_state: uint16(2),
	},
	271: {
		Fexternal_lex_state: uint16(2),
	},
	272: {
		Fexternal_lex_state: uint16(2),
	},
	273: {
		Fexternal_lex_state: uint16(2),
	},
	274: {
		Fexternal_lex_state: uint16(4),
	},
	275: {
		Fexternal_lex_state: uint16(7),
	},
	276: {
		Fexternal_lex_state: uint16(2),
	},
	277: {
		Fexternal_lex_state: uint16(2),
	},
	278: {
		Fexternal_lex_state: uint16(7),
	},
	279: {
		Fexternal_lex_state: uint16(2),
	},
	280: {
		Flex_state: libc.Uint16FromInt32(-libc.Int32FromInt32(1)),
	},
	281: {
		Flex_state: libc.Uint16FromInt32(-libc.Int32FromInt32(1)),
	},
}

var ts_parse_table = [2][143]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
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
		31: uint16(1),
		32: uint16(1),
		33: uint16(1),
		34: uint16(1),
		37: uint16(1),
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
		52: uint16(1),
		53: uint16(1),
		54: uint16(1),
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		58: uint16(1),
		59: uint16(1),
		60: uint16(1),
		61: uint16(1),
		62: uint16(1),
		63: uint16(1),
		64: uint16(1),
		65: uint16(3),
		67: uint16(5),
		68: uint16(1),
		69: uint16(1),
		70: uint16(1),
		71: uint16(1),
		72: uint16(1),
	},
	1: {
		0:   uint16(7),
		1:   uint16(9),
		2:   uint16(11),
		3:   uint16(13),
		4:   uint16(15),
		7:   uint16(17),
		8:   uint16(19),
		9:   uint16(21),
		10:  uint16(23),
		12:  uint16(25),
		13:  uint16(27),
		15:  uint16(29),
		19:  uint16(31),
		21:  uint16(33),
		22:  uint16(35),
		23:  uint16(37),
		39:  uint16(39),
		65:  uint16(3),
		67:  uint16(5),
		73:  uint16(263),
		75:  uint16(161),
		76:  uint16(261),
		77:  uint16(160),
		78:  uint16(160),
		79:  uint16(255),
		81:  uint16(160),
		82:  uint16(160),
		83:  uint16(160),
		84:  uint16(160),
		85:  uint16(160),
		86:  uint16(160),
		89:  uint16(160),
		93:  uint16(132),
		94:  uint16(134),
		95:  uint16(137),
		96:  uint16(139),
		101: uint16(131),
		102: uint16(142),
		105: uint16(143),
		119: uint16(170),
		120: uint16(172),
		121: uint16(2),
		122: uint16(2),
		123: uint16(101),
		124: uint16(174),
		126: uint16(177),
		133: uint16(1),
		134: uint16(2),
		135: uint16(55),
	},
}

var ts_small_parse_table = [12930]uint16_t{
	0:     uint16(5),
	1:     uint16(3),
	2:     uint16(1),
	3:     uint16(anon_sym_DASH_DASH),
	4:     uint16(5),
	5:     uint16(1),
	6:     uint16(sym__block_comment_start),
	7:     uint16(2),
	8:     uint16(1),
	9:     uint16(sym_comment),
	10:    uint16(41),
	11:    uint16(27),
	12:    uint16(sym__block_string_start),
	14:    uint16(anon_sym_SEMI),
	15:    uint16(anon_sym_COMMA),
	16:    uint16(anon_sym_COLON_COLON),
	17:    uint16(anon_sym_STAR),
	18:    uint16(anon_sym_DQUOTE),
	19:    uint16(anon_sym_SQUOTE),
	20:    uint16(anon_sym_LPAREN),
	21:    uint16(anon_sym_RPAREN),
	22:    uint16(anon_sym_LBRACK),
	23:    uint16(anon_sym_RBRACK),
	24:    uint16(anon_sym_LBRACE),
	25:    uint16(anon_sym_RBRACE),
	26:    uint16(anon_sym_LT_EQ),
	27:    uint16(anon_sym_EQ_EQ),
	28:    uint16(anon_sym_TILDE_EQ),
	29:    uint16(anon_sym_GT_EQ),
	30:    uint16(anon_sym_PIPE),
	31:    uint16(anon_sym_AMP),
	32:    uint16(anon_sym_LT_LT),
	33:    uint16(anon_sym_GT_GT),
	34:    uint16(anon_sym_PLUS),
	35:    uint16(anon_sym_SLASH_SLASH),
	36:    uint16(anon_sym_PERCENT),
	37:    uint16(anon_sym_DOT_DOT),
	38:    uint16(anon_sym_CARET),
	39:    uint16(43),
	40:    uint16(27),
	41:    uint16(anon_sym_return),
	42:    uint16(anon_sym_EQ),
	43:    uint16(sym_break_statement),
	44:    uint16(anon_sym_goto),
	45:    uint16(anon_sym_do),
	46:    uint16(anon_sym_end),
	47:    uint16(anon_sym_while),
	48:    uint16(anon_sym_repeat),
	49:    uint16(anon_sym_until),
	50:    uint16(anon_sym_if),
	51:    uint16(anon_sym_then),
	52:    uint16(anon_sym_elseif),
	53:    uint16(anon_sym_else),
	54:    uint16(anon_sym_for),
	55:    uint16(anon_sym_function),
	56:    uint16(anon_sym_local),
	57:    uint16(anon_sym_global),
	58:    uint16(anon_sym_DOT),
	59:    uint16(anon_sym_COLON),
	60:    uint16(anon_sym_LT),
	61:    uint16(anon_sym_GT),
	62:    uint16(anon_sym_or),
	63:    uint16(anon_sym_and),
	64:    uint16(anon_sym_TILDE),
	65:    uint16(anon_sym_DASH),
	66:    uint16(anon_sym_SLASH),
	67:    uint16(sym_identifier),
	68:    uint16(5),
	69:    uint16(3),
	70:    uint16(1),
	71:    uint16(anon_sym_DASH_DASH),
	72:    uint16(5),
	73:    uint16(1),
	74:    uint16(sym__block_comment_start),
	75:    uint16(3),
	76:    uint16(1),
	77:    uint16(sym_comment),
	78:    uint16(45),
	79:    uint16(27),
	80:    uint16(sym__block_string_start),
	82:    uint16(anon_sym_SEMI),
	83:    uint16(anon_sym_COMMA),
	84:    uint16(anon_sym_COLON_COLON),
	85:    uint16(anon_sym_STAR),
	86:    uint16(anon_sym_DQUOTE),
	87:    uint16(anon_sym_SQUOTE),
	88:    uint16(anon_sym_LPAREN),
	89:    uint16(anon_sym_RPAREN),
	90:    uint16(anon_sym_LBRACK),
	91:    uint16(anon_sym_RBRACK),
	92:    uint16(anon_sym_LBRACE),
	93:    uint16(anon_sym_RBRACE),
	94:    uint16(anon_sym_LT_EQ),
	95:    uint16(anon_sym_EQ_EQ),
	96:    uint16(anon_sym_TILDE_EQ),
	97:    uint16(anon_sym_GT_EQ),
	98:    uint16(anon_sym_PIPE),
	99:    uint16(anon_sym_AMP),
	100:   uint16(anon_sym_LT_LT),
	101:   uint16(anon_sym_GT_GT),
	102:   uint16(anon_sym_PLUS),
	103:   uint16(anon_sym_SLASH_SLASH),
	104:   uint16(anon_sym_PERCENT),
	105:   uint16(anon_sym_DOT_DOT),
	106:   uint16(anon_sym_CARET),
	107:   uint16(47),
	108:   uint16(27),
	109:   uint16(anon_sym_return),
	110:   uint16(anon_sym_EQ),
	111:   uint16(sym_break_statement),
	112:   uint16(anon_sym_goto),
	113:   uint16(anon_sym_do),
	114:   uint16(anon_sym_end),
	115:   uint16(anon_sym_while),
	116:   uint16(anon_sym_repeat),
	117:   uint16(anon_sym_until),
	118:   uint16(anon_sym_if),
	119:   uint16(anon_sym_then),
	120:   uint16(anon_sym_elseif),
	121:   uint16(anon_sym_else),
	122:   uint16(anon_sym_for),
	123:   uint16(anon_sym_function),
	124:   uint16(anon_sym_local),
	125:   uint16(anon_sym_global),
	126:   uint16(anon_sym_DOT),
	127:   uint16(anon_sym_COLON),
	128:   uint16(anon_sym_LT),
	129:   uint16(anon_sym_GT),
	130:   uint16(anon_sym_or),
	131:   uint16(anon_sym_and),
	132:   uint16(anon_sym_TILDE),
	133:   uint16(anon_sym_DASH),
	134:   uint16(anon_sym_SLASH),
	135:   uint16(sym_identifier),
	136:   uint16(5),
	137:   uint16(3),
	138:   uint16(1),
	139:   uint16(anon_sym_DASH_DASH),
	140:   uint16(5),
	141:   uint16(1),
	142:   uint16(sym__block_comment_start),
	143:   uint16(4),
	144:   uint16(1),
	145:   uint16(sym_comment),
	146:   uint16(49),
	147:   uint16(27),
	148:   uint16(sym__block_string_start),
	150:   uint16(anon_sym_SEMI),
	151:   uint16(anon_sym_COMMA),
	152:   uint16(anon_sym_COLON_COLON),
	153:   uint16(anon_sym_STAR),
	154:   uint16(anon_sym_DQUOTE),
	155:   uint16(anon_sym_SQUOTE),
	156:   uint16(anon_sym_LPAREN),
	157:   uint16(anon_sym_RPAREN),
	158:   uint16(anon_sym_LBRACK),
	159:   uint16(anon_sym_RBRACK),
	160:   uint16(anon_sym_LBRACE),
	161:   uint16(anon_sym_RBRACE),
	162:   uint16(anon_sym_LT_EQ),
	163:   uint16(anon_sym_EQ_EQ),
	164:   uint16(anon_sym_TILDE_EQ),
	165:   uint16(anon_sym_GT_EQ),
	166:   uint16(anon_sym_PIPE),
	167:   uint16(anon_sym_AMP),
	168:   uint16(anon_sym_LT_LT),
	169:   uint16(anon_sym_GT_GT),
	170:   uint16(anon_sym_PLUS),
	171:   uint16(anon_sym_SLASH_SLASH),
	172:   uint16(anon_sym_PERCENT),
	173:   uint16(anon_sym_DOT_DOT),
	174:   uint16(anon_sym_CARET),
	175:   uint16(51),
	176:   uint16(27),
	177:   uint16(anon_sym_return),
	178:   uint16(anon_sym_EQ),
	179:   uint16(sym_break_statement),
	180:   uint16(anon_sym_goto),
	181:   uint16(anon_sym_do),
	182:   uint16(anon_sym_end),
	183:   uint16(anon_sym_while),
	184:   uint16(anon_sym_repeat),
	185:   uint16(anon_sym_until),
	186:   uint16(anon_sym_if),
	187:   uint16(anon_sym_then),
	188:   uint16(anon_sym_elseif),
	189:   uint16(anon_sym_else),
	190:   uint16(anon_sym_for),
	191:   uint16(anon_sym_function),
	192:   uint16(anon_sym_local),
	193:   uint16(anon_sym_global),
	194:   uint16(anon_sym_DOT),
	195:   uint16(anon_sym_COLON),
	196:   uint16(anon_sym_LT),
	197:   uint16(anon_sym_GT),
	198:   uint16(anon_sym_or),
	199:   uint16(anon_sym_and),
	200:   uint16(anon_sym_TILDE),
	201:   uint16(anon_sym_DASH),
	202:   uint16(anon_sym_SLASH),
	203:   uint16(sym_identifier),
	204:   uint16(5),
	205:   uint16(3),
	206:   uint16(1),
	207:   uint16(anon_sym_DASH_DASH),
	208:   uint16(5),
	209:   uint16(1),
	210:   uint16(sym__block_comment_start),
	211:   uint16(5),
	212:   uint16(1),
	213:   uint16(sym_comment),
	214:   uint16(53),
	215:   uint16(27),
	216:   uint16(sym__block_string_start),
	218:   uint16(anon_sym_SEMI),
	219:   uint16(anon_sym_COMMA),
	220:   uint16(anon_sym_COLON_COLON),
	221:   uint16(anon_sym_STAR),
	222:   uint16(anon_sym_DQUOTE),
	223:   uint16(anon_sym_SQUOTE),
	224:   uint16(anon_sym_LPAREN),
	225:   uint16(anon_sym_RPAREN),
	226:   uint16(anon_sym_LBRACK),
	227:   uint16(anon_sym_RBRACK),
	228:   uint16(anon_sym_LBRACE),
	229:   uint16(anon_sym_RBRACE),
	230:   uint16(anon_sym_LT_EQ),
	231:   uint16(anon_sym_EQ_EQ),
	232:   uint16(anon_sym_TILDE_EQ),
	233:   uint16(anon_sym_GT_EQ),
	234:   uint16(anon_sym_PIPE),
	235:   uint16(anon_sym_AMP),
	236:   uint16(anon_sym_LT_LT),
	237:   uint16(anon_sym_GT_GT),
	238:   uint16(anon_sym_PLUS),
	239:   uint16(anon_sym_SLASH_SLASH),
	240:   uint16(anon_sym_PERCENT),
	241:   uint16(anon_sym_DOT_DOT),
	242:   uint16(anon_sym_CARET),
	243:   uint16(55),
	244:   uint16(27),
	245:   uint16(anon_sym_return),
	246:   uint16(anon_sym_EQ),
	247:   uint16(sym_break_statement),
	248:   uint16(anon_sym_goto),
	249:   uint16(anon_sym_do),
	250:   uint16(anon_sym_end),
	251:   uint16(anon_sym_while),
	252:   uint16(anon_sym_repeat),
	253:   uint16(anon_sym_until),
	254:   uint16(anon_sym_if),
	255:   uint16(anon_sym_then),
	256:   uint16(anon_sym_elseif),
	257:   uint16(anon_sym_else),
	258:   uint16(anon_sym_for),
	259:   uint16(anon_sym_function),
	260:   uint16(anon_sym_local),
	261:   uint16(anon_sym_global),
	262:   uint16(anon_sym_DOT),
	263:   uint16(anon_sym_COLON),
	264:   uint16(anon_sym_LT),
	265:   uint16(anon_sym_GT),
	266:   uint16(anon_sym_or),
	267:   uint16(anon_sym_and),
	268:   uint16(anon_sym_TILDE),
	269:   uint16(anon_sym_DASH),
	270:   uint16(anon_sym_SLASH),
	271:   uint16(sym_identifier),
	272:   uint16(5),
	273:   uint16(3),
	274:   uint16(1),
	275:   uint16(anon_sym_DASH_DASH),
	276:   uint16(5),
	277:   uint16(1),
	278:   uint16(sym__block_comment_start),
	279:   uint16(6),
	280:   uint16(1),
	281:   uint16(sym_comment),
	282:   uint16(59),
	283:   uint16(26),
	284:   uint16(anon_sym_return),
	285:   uint16(sym_break_statement),
	286:   uint16(anon_sym_goto),
	287:   uint16(anon_sym_do),
	288:   uint16(anon_sym_end),
	289:   uint16(anon_sym_while),
	290:   uint16(anon_sym_repeat),
	291:   uint16(anon_sym_until),
	292:   uint16(anon_sym_if),
	293:   uint16(anon_sym_then),
	294:   uint16(anon_sym_elseif),
	295:   uint16(anon_sym_else),
	296:   uint16(anon_sym_for),
	297:   uint16(anon_sym_function),
	298:   uint16(anon_sym_local),
	299:   uint16(anon_sym_global),
	300:   uint16(anon_sym_DOT),
	301:   uint16(anon_sym_COLON),
	302:   uint16(anon_sym_LT),
	303:   uint16(anon_sym_GT),
	304:   uint16(anon_sym_or),
	305:   uint16(anon_sym_and),
	306:   uint16(anon_sym_TILDE),
	307:   uint16(anon_sym_DASH),
	308:   uint16(anon_sym_SLASH),
	309:   uint16(sym_identifier),
	310:   uint16(57),
	311:   uint16(27),
	312:   uint16(sym__block_string_start),
	314:   uint16(anon_sym_SEMI),
	315:   uint16(anon_sym_COMMA),
	316:   uint16(anon_sym_COLON_COLON),
	317:   uint16(anon_sym_STAR),
	318:   uint16(anon_sym_DQUOTE),
	319:   uint16(anon_sym_SQUOTE),
	320:   uint16(anon_sym_LPAREN),
	321:   uint16(anon_sym_RPAREN),
	322:   uint16(anon_sym_LBRACK),
	323:   uint16(anon_sym_RBRACK),
	324:   uint16(anon_sym_LBRACE),
	325:   uint16(anon_sym_RBRACE),
	326:   uint16(anon_sym_LT_EQ),
	327:   uint16(anon_sym_EQ_EQ),
	328:   uint16(anon_sym_TILDE_EQ),
	329:   uint16(anon_sym_GT_EQ),
	330:   uint16(anon_sym_PIPE),
	331:   uint16(anon_sym_AMP),
	332:   uint16(anon_sym_LT_LT),
	333:   uint16(anon_sym_GT_GT),
	334:   uint16(anon_sym_PLUS),
	335:   uint16(anon_sym_SLASH_SLASH),
	336:   uint16(anon_sym_PERCENT),
	337:   uint16(anon_sym_DOT_DOT),
	338:   uint16(anon_sym_CARET),
	339:   uint16(7),
	340:   uint16(3),
	341:   uint16(1),
	342:   uint16(anon_sym_DASH_DASH),
	343:   uint16(5),
	344:   uint16(1),
	345:   uint16(sym__block_comment_start),
	346:   uint16(7),
	347:   uint16(1),
	348:   uint16(sym_comment),
	349:   uint16(65),
	350:   uint16(2),
	351:   uint16(anon_sym_DOT),
	352:   uint16(anon_sym_COLON),
	353:   uint16(67),
	354:   uint16(6),
	355:   uint16(sym__block_string_start),
	356:   uint16(anon_sym_DQUOTE),
	357:   uint16(anon_sym_SQUOTE),
	358:   uint16(anon_sym_LPAREN),
	359:   uint16(anon_sym_LBRACK),
	360:   uint16(anon_sym_LBRACE),
	361:   uint16(61),
	362:   uint16(21),
	364:   uint16(anon_sym_SEMI),
	365:   uint16(anon_sym_COMMA),
	366:   uint16(anon_sym_COLON_COLON),
	367:   uint16(anon_sym_STAR),
	368:   uint16(anon_sym_RPAREN),
	369:   uint16(anon_sym_RBRACK),
	370:   uint16(anon_sym_RBRACE),
	371:   uint16(anon_sym_LT_EQ),
	372:   uint16(anon_sym_EQ_EQ),
	373:   uint16(anon_sym_TILDE_EQ),
	374:   uint16(anon_sym_GT_EQ),
	375:   uint16(anon_sym_PIPE),
	376:   uint16(anon_sym_AMP),
	377:   uint16(anon_sym_LT_LT),
	378:   uint16(anon_sym_GT_GT),
	379:   uint16(anon_sym_PLUS),
	380:   uint16(anon_sym_SLASH_SLASH),
	381:   uint16(anon_sym_PERCENT),
	382:   uint16(anon_sym_DOT_DOT),
	383:   uint16(anon_sym_CARET),
	384:   uint16(63),
	385:   uint16(24),
	386:   uint16(anon_sym_return),
	387:   uint16(sym_break_statement),
	388:   uint16(anon_sym_goto),
	389:   uint16(anon_sym_do),
	390:   uint16(anon_sym_end),
	391:   uint16(anon_sym_while),
	392:   uint16(anon_sym_repeat),
	393:   uint16(anon_sym_until),
	394:   uint16(anon_sym_if),
	395:   uint16(anon_sym_then),
	396:   uint16(anon_sym_elseif),
	397:   uint16(anon_sym_else),
	398:   uint16(anon_sym_for),
	399:   uint16(anon_sym_function),
	400:   uint16(anon_sym_local),
	401:   uint16(anon_sym_global),
	402:   uint16(anon_sym_LT),
	403:   uint16(anon_sym_GT),
	404:   uint16(anon_sym_or),
	405:   uint16(anon_sym_and),
	406:   uint16(anon_sym_TILDE),
	407:   uint16(anon_sym_DASH),
	408:   uint16(anon_sym_SLASH),
	409:   uint16(sym_identifier),
	410:   uint16(5),
	411:   uint16(3),
	412:   uint16(1),
	413:   uint16(anon_sym_DASH_DASH),
	414:   uint16(5),
	415:   uint16(1),
	416:   uint16(sym__block_comment_start),
	417:   uint16(8),
	418:   uint16(1),
	419:   uint16(sym_comment),
	420:   uint16(71),
	421:   uint16(26),
	422:   uint16(anon_sym_return),
	423:   uint16(sym_break_statement),
	424:   uint16(anon_sym_goto),
	425:   uint16(anon_sym_do),
	426:   uint16(anon_sym_end),
	427:   uint16(anon_sym_while),
	428:   uint16(anon_sym_repeat),
	429:   uint16(anon_sym_until),
	430:   uint16(anon_sym_if),
	431:   uint16(anon_sym_then),
	432:   uint16(anon_sym_elseif),
	433:   uint16(anon_sym_else),
	434:   uint16(anon_sym_for),
	435:   uint16(anon_sym_function),
	436:   uint16(anon_sym_local),
	437:   uint16(anon_sym_global),
	438:   uint16(anon_sym_DOT),
	439:   uint16(anon_sym_COLON),
	440:   uint16(anon_sym_LT),
	441:   uint16(anon_sym_GT),
	442:   uint16(anon_sym_or),
	443:   uint16(anon_sym_and),
	444:   uint16(anon_sym_TILDE),
	445:   uint16(anon_sym_DASH),
	446:   uint16(anon_sym_SLASH),
	447:   uint16(sym_identifier),
	448:   uint16(69),
	449:   uint16(27),
	450:   uint16(sym__block_string_start),
	452:   uint16(anon_sym_SEMI),
	453:   uint16(anon_sym_COMMA),
	454:   uint16(anon_sym_COLON_COLON),
	455:   uint16(anon_sym_STAR),
	456:   uint16(anon_sym_DQUOTE),
	457:   uint16(anon_sym_SQUOTE),
	458:   uint16(anon_sym_LPAREN),
	459:   uint16(anon_sym_RPAREN),
	460:   uint16(anon_sym_LBRACK),
	461:   uint16(anon_sym_RBRACK),
	462:   uint16(anon_sym_LBRACE),
	463:   uint16(anon_sym_RBRACE),
	464:   uint16(anon_sym_LT_EQ),
	465:   uint16(anon_sym_EQ_EQ),
	466:   uint16(anon_sym_TILDE_EQ),
	467:   uint16(anon_sym_GT_EQ),
	468:   uint16(anon_sym_PIPE),
	469:   uint16(anon_sym_AMP),
	470:   uint16(anon_sym_LT_LT),
	471:   uint16(anon_sym_GT_GT),
	472:   uint16(anon_sym_PLUS),
	473:   uint16(anon_sym_SLASH_SLASH),
	474:   uint16(anon_sym_PERCENT),
	475:   uint16(anon_sym_DOT_DOT),
	476:   uint16(anon_sym_CARET),
	477:   uint16(5),
	478:   uint16(3),
	479:   uint16(1),
	480:   uint16(anon_sym_DASH_DASH),
	481:   uint16(5),
	482:   uint16(1),
	483:   uint16(sym__block_comment_start),
	484:   uint16(9),
	485:   uint16(1),
	486:   uint16(sym_comment),
	487:   uint16(75),
	488:   uint16(26),
	489:   uint16(anon_sym_return),
	490:   uint16(sym_break_statement),
	491:   uint16(anon_sym_goto),
	492:   uint16(anon_sym_do),
	493:   uint16(anon_sym_end),
	494:   uint16(anon_sym_while),
	495:   uint16(anon_sym_repeat),
	496:   uint16(anon_sym_until),
	497:   uint16(anon_sym_if),
	498:   uint16(anon_sym_then),
	499:   uint16(anon_sym_elseif),
	500:   uint16(anon_sym_else),
	501:   uint16(anon_sym_for),
	502:   uint16(anon_sym_function),
	503:   uint16(anon_sym_local),
	504:   uint16(anon_sym_global),
	505:   uint16(anon_sym_DOT),
	506:   uint16(anon_sym_COLON),
	507:   uint16(anon_sym_LT),
	508:   uint16(anon_sym_GT),
	509:   uint16(anon_sym_or),
	510:   uint16(anon_sym_and),
	511:   uint16(anon_sym_TILDE),
	512:   uint16(anon_sym_DASH),
	513:   uint16(anon_sym_SLASH),
	514:   uint16(sym_identifier),
	515:   uint16(73),
	516:   uint16(27),
	517:   uint16(sym__block_string_start),
	519:   uint16(anon_sym_SEMI),
	520:   uint16(anon_sym_COMMA),
	521:   uint16(anon_sym_COLON_COLON),
	522:   uint16(anon_sym_STAR),
	523:   uint16(anon_sym_DQUOTE),
	524:   uint16(anon_sym_SQUOTE),
	525:   uint16(anon_sym_LPAREN),
	526:   uint16(anon_sym_RPAREN),
	527:   uint16(anon_sym_LBRACK),
	528:   uint16(anon_sym_RBRACK),
	529:   uint16(anon_sym_LBRACE),
	530:   uint16(anon_sym_RBRACE),
	531:   uint16(anon_sym_LT_EQ),
	532:   uint16(anon_sym_EQ_EQ),
	533:   uint16(anon_sym_TILDE_EQ),
	534:   uint16(anon_sym_GT_EQ),
	535:   uint16(anon_sym_PIPE),
	536:   uint16(anon_sym_AMP),
	537:   uint16(anon_sym_LT_LT),
	538:   uint16(anon_sym_GT_GT),
	539:   uint16(anon_sym_PLUS),
	540:   uint16(anon_sym_SLASH_SLASH),
	541:   uint16(anon_sym_PERCENT),
	542:   uint16(anon_sym_DOT_DOT),
	543:   uint16(anon_sym_CARET),
	544:   uint16(5),
	545:   uint16(3),
	546:   uint16(1),
	547:   uint16(anon_sym_DASH_DASH),
	548:   uint16(5),
	549:   uint16(1),
	550:   uint16(sym__block_comment_start),
	551:   uint16(10),
	552:   uint16(1),
	553:   uint16(sym_comment),
	554:   uint16(79),
	555:   uint16(26),
	556:   uint16(anon_sym_return),
	557:   uint16(sym_break_statement),
	558:   uint16(anon_sym_goto),
	559:   uint16(anon_sym_do),
	560:   uint16(anon_sym_end),
	561:   uint16(anon_sym_while),
	562:   uint16(anon_sym_repeat),
	563:   uint16(anon_sym_until),
	564:   uint16(anon_sym_if),
	565:   uint16(anon_sym_then),
	566:   uint16(anon_sym_elseif),
	567:   uint16(anon_sym_else),
	568:   uint16(anon_sym_for),
	569:   uint16(anon_sym_function),
	570:   uint16(anon_sym_local),
	571:   uint16(anon_sym_global),
	572:   uint16(anon_sym_DOT),
	573:   uint16(anon_sym_COLON),
	574:   uint16(anon_sym_LT),
	575:   uint16(anon_sym_GT),
	576:   uint16(anon_sym_or),
	577:   uint16(anon_sym_and),
	578:   uint16(anon_sym_TILDE),
	579:   uint16(anon_sym_DASH),
	580:   uint16(anon_sym_SLASH),
	581:   uint16(sym_identifier),
	582:   uint16(77),
	583:   uint16(27),
	584:   uint16(sym__block_string_start),
	586:   uint16(anon_sym_SEMI),
	587:   uint16(anon_sym_COMMA),
	588:   uint16(anon_sym_COLON_COLON),
	589:   uint16(anon_sym_STAR),
	590:   uint16(anon_sym_DQUOTE),
	591:   uint16(anon_sym_SQUOTE),
	592:   uint16(anon_sym_LPAREN),
	593:   uint16(anon_sym_RPAREN),
	594:   uint16(anon_sym_LBRACK),
	595:   uint16(anon_sym_RBRACK),
	596:   uint16(anon_sym_LBRACE),
	597:   uint16(anon_sym_RBRACE),
	598:   uint16(anon_sym_LT_EQ),
	599:   uint16(anon_sym_EQ_EQ),
	600:   uint16(anon_sym_TILDE_EQ),
	601:   uint16(anon_sym_GT_EQ),
	602:   uint16(anon_sym_PIPE),
	603:   uint16(anon_sym_AMP),
	604:   uint16(anon_sym_LT_LT),
	605:   uint16(anon_sym_GT_GT),
	606:   uint16(anon_sym_PLUS),
	607:   uint16(anon_sym_SLASH_SLASH),
	608:   uint16(anon_sym_PERCENT),
	609:   uint16(anon_sym_DOT_DOT),
	610:   uint16(anon_sym_CARET),
	611:   uint16(5),
	612:   uint16(3),
	613:   uint16(1),
	614:   uint16(anon_sym_DASH_DASH),
	615:   uint16(5),
	616:   uint16(1),
	617:   uint16(sym__block_comment_start),
	618:   uint16(11),
	619:   uint16(1),
	620:   uint16(sym_comment),
	621:   uint16(83),
	622:   uint16(26),
	623:   uint16(anon_sym_return),
	624:   uint16(sym_break_statement),
	625:   uint16(anon_sym_goto),
	626:   uint16(anon_sym_do),
	627:   uint16(anon_sym_end),
	628:   uint16(anon_sym_while),
	629:   uint16(anon_sym_repeat),
	630:   uint16(anon_sym_until),
	631:   uint16(anon_sym_if),
	632:   uint16(anon_sym_then),
	633:   uint16(anon_sym_elseif),
	634:   uint16(anon_sym_else),
	635:   uint16(anon_sym_for),
	636:   uint16(anon_sym_function),
	637:   uint16(anon_sym_local),
	638:   uint16(anon_sym_global),
	639:   uint16(anon_sym_DOT),
	640:   uint16(anon_sym_COLON),
	641:   uint16(anon_sym_LT),
	642:   uint16(anon_sym_GT),
	643:   uint16(anon_sym_or),
	644:   uint16(anon_sym_and),
	645:   uint16(anon_sym_TILDE),
	646:   uint16(anon_sym_DASH),
	647:   uint16(anon_sym_SLASH),
	648:   uint16(sym_identifier),
	649:   uint16(81),
	650:   uint16(27),
	651:   uint16(sym__block_string_start),
	653:   uint16(anon_sym_SEMI),
	654:   uint16(anon_sym_COMMA),
	655:   uint16(anon_sym_COLON_COLON),
	656:   uint16(anon_sym_STAR),
	657:   uint16(anon_sym_DQUOTE),
	658:   uint16(anon_sym_SQUOTE),
	659:   uint16(anon_sym_LPAREN),
	660:   uint16(anon_sym_RPAREN),
	661:   uint16(anon_sym_LBRACK),
	662:   uint16(anon_sym_RBRACK),
	663:   uint16(anon_sym_LBRACE),
	664:   uint16(anon_sym_RBRACE),
	665:   uint16(anon_sym_LT_EQ),
	666:   uint16(anon_sym_EQ_EQ),
	667:   uint16(anon_sym_TILDE_EQ),
	668:   uint16(anon_sym_GT_EQ),
	669:   uint16(anon_sym_PIPE),
	670:   uint16(anon_sym_AMP),
	671:   uint16(anon_sym_LT_LT),
	672:   uint16(anon_sym_GT_GT),
	673:   uint16(anon_sym_PLUS),
	674:   uint16(anon_sym_SLASH_SLASH),
	675:   uint16(anon_sym_PERCENT),
	676:   uint16(anon_sym_DOT_DOT),
	677:   uint16(anon_sym_CARET),
	678:   uint16(5),
	679:   uint16(3),
	680:   uint16(1),
	681:   uint16(anon_sym_DASH_DASH),
	682:   uint16(5),
	683:   uint16(1),
	684:   uint16(sym__block_comment_start),
	685:   uint16(12),
	686:   uint16(1),
	687:   uint16(sym_comment),
	688:   uint16(87),
	689:   uint16(26),
	690:   uint16(anon_sym_return),
	691:   uint16(sym_break_statement),
	692:   uint16(anon_sym_goto),
	693:   uint16(anon_sym_do),
	694:   uint16(anon_sym_end),
	695:   uint16(anon_sym_while),
	696:   uint16(anon_sym_repeat),
	697:   uint16(anon_sym_until),
	698:   uint16(anon_sym_if),
	699:   uint16(anon_sym_then),
	700:   uint16(anon_sym_elseif),
	701:   uint16(anon_sym_else),
	702:   uint16(anon_sym_for),
	703:   uint16(anon_sym_function),
	704:   uint16(anon_sym_local),
	705:   uint16(anon_sym_global),
	706:   uint16(anon_sym_DOT),
	707:   uint16(anon_sym_COLON),
	708:   uint16(anon_sym_LT),
	709:   uint16(anon_sym_GT),
	710:   uint16(anon_sym_or),
	711:   uint16(anon_sym_and),
	712:   uint16(anon_sym_TILDE),
	713:   uint16(anon_sym_DASH),
	714:   uint16(anon_sym_SLASH),
	715:   uint16(sym_identifier),
	716:   uint16(85),
	717:   uint16(27),
	718:   uint16(sym__block_string_start),
	720:   uint16(anon_sym_SEMI),
	721:   uint16(anon_sym_COMMA),
	722:   uint16(anon_sym_COLON_COLON),
	723:   uint16(anon_sym_STAR),
	724:   uint16(anon_sym_DQUOTE),
	725:   uint16(anon_sym_SQUOTE),
	726:   uint16(anon_sym_LPAREN),
	727:   uint16(anon_sym_RPAREN),
	728:   uint16(anon_sym_LBRACK),
	729:   uint16(anon_sym_RBRACK),
	730:   uint16(anon_sym_LBRACE),
	731:   uint16(anon_sym_RBRACE),
	732:   uint16(anon_sym_LT_EQ),
	733:   uint16(anon_sym_EQ_EQ),
	734:   uint16(anon_sym_TILDE_EQ),
	735:   uint16(anon_sym_GT_EQ),
	736:   uint16(anon_sym_PIPE),
	737:   uint16(anon_sym_AMP),
	738:   uint16(anon_sym_LT_LT),
	739:   uint16(anon_sym_GT_GT),
	740:   uint16(anon_sym_PLUS),
	741:   uint16(anon_sym_SLASH_SLASH),
	742:   uint16(anon_sym_PERCENT),
	743:   uint16(anon_sym_DOT_DOT),
	744:   uint16(anon_sym_CARET),
	745:   uint16(5),
	746:   uint16(3),
	747:   uint16(1),
	748:   uint16(anon_sym_DASH_DASH),
	749:   uint16(5),
	750:   uint16(1),
	751:   uint16(sym__block_comment_start),
	752:   uint16(13),
	753:   uint16(1),
	754:   uint16(sym_comment),
	755:   uint16(91),
	756:   uint16(26),
	757:   uint16(anon_sym_return),
	758:   uint16(sym_break_statement),
	759:   uint16(anon_sym_goto),
	760:   uint16(anon_sym_do),
	761:   uint16(anon_sym_end),
	762:   uint16(anon_sym_while),
	763:   uint16(anon_sym_repeat),
	764:   uint16(anon_sym_until),
	765:   uint16(anon_sym_if),
	766:   uint16(anon_sym_then),
	767:   uint16(anon_sym_elseif),
	768:   uint16(anon_sym_else),
	769:   uint16(anon_sym_for),
	770:   uint16(anon_sym_function),
	771:   uint16(anon_sym_local),
	772:   uint16(anon_sym_global),
	773:   uint16(anon_sym_DOT),
	774:   uint16(anon_sym_COLON),
	775:   uint16(anon_sym_LT),
	776:   uint16(anon_sym_GT),
	777:   uint16(anon_sym_or),
	778:   uint16(anon_sym_and),
	779:   uint16(anon_sym_TILDE),
	780:   uint16(anon_sym_DASH),
	781:   uint16(anon_sym_SLASH),
	782:   uint16(sym_identifier),
	783:   uint16(89),
	784:   uint16(27),
	785:   uint16(sym__block_string_start),
	787:   uint16(anon_sym_SEMI),
	788:   uint16(anon_sym_COMMA),
	789:   uint16(anon_sym_COLON_COLON),
	790:   uint16(anon_sym_STAR),
	791:   uint16(anon_sym_DQUOTE),
	792:   uint16(anon_sym_SQUOTE),
	793:   uint16(anon_sym_LPAREN),
	794:   uint16(anon_sym_RPAREN),
	795:   uint16(anon_sym_LBRACK),
	796:   uint16(anon_sym_RBRACK),
	797:   uint16(anon_sym_LBRACE),
	798:   uint16(anon_sym_RBRACE),
	799:   uint16(anon_sym_LT_EQ),
	800:   uint16(anon_sym_EQ_EQ),
	801:   uint16(anon_sym_TILDE_EQ),
	802:   uint16(anon_sym_GT_EQ),
	803:   uint16(anon_sym_PIPE),
	804:   uint16(anon_sym_AMP),
	805:   uint16(anon_sym_LT_LT),
	806:   uint16(anon_sym_GT_GT),
	807:   uint16(anon_sym_PLUS),
	808:   uint16(anon_sym_SLASH_SLASH),
	809:   uint16(anon_sym_PERCENT),
	810:   uint16(anon_sym_DOT_DOT),
	811:   uint16(anon_sym_CARET),
	812:   uint16(5),
	813:   uint16(3),
	814:   uint16(1),
	815:   uint16(anon_sym_DASH_DASH),
	816:   uint16(5),
	817:   uint16(1),
	818:   uint16(sym__block_comment_start),
	819:   uint16(14),
	820:   uint16(1),
	821:   uint16(sym_comment),
	822:   uint16(95),
	823:   uint16(26),
	824:   uint16(anon_sym_return),
	825:   uint16(sym_break_statement),
	826:   uint16(anon_sym_goto),
	827:   uint16(anon_sym_do),
	828:   uint16(anon_sym_end),
	829:   uint16(anon_sym_while),
	830:   uint16(anon_sym_repeat),
	831:   uint16(anon_sym_until),
	832:   uint16(anon_sym_if),
	833:   uint16(anon_sym_then),
	834:   uint16(anon_sym_elseif),
	835:   uint16(anon_sym_else),
	836:   uint16(anon_sym_for),
	837:   uint16(anon_sym_function),
	838:   uint16(anon_sym_local),
	839:   uint16(anon_sym_global),
	840:   uint16(anon_sym_DOT),
	841:   uint16(anon_sym_COLON),
	842:   uint16(anon_sym_LT),
	843:   uint16(anon_sym_GT),
	844:   uint16(anon_sym_or),
	845:   uint16(anon_sym_and),
	846:   uint16(anon_sym_TILDE),
	847:   uint16(anon_sym_DASH),
	848:   uint16(anon_sym_SLASH),
	849:   uint16(sym_identifier),
	850:   uint16(93),
	851:   uint16(27),
	852:   uint16(sym__block_string_start),
	854:   uint16(anon_sym_SEMI),
	855:   uint16(anon_sym_COMMA),
	856:   uint16(anon_sym_COLON_COLON),
	857:   uint16(anon_sym_STAR),
	858:   uint16(anon_sym_DQUOTE),
	859:   uint16(anon_sym_SQUOTE),
	860:   uint16(anon_sym_LPAREN),
	861:   uint16(anon_sym_RPAREN),
	862:   uint16(anon_sym_LBRACK),
	863:   uint16(anon_sym_RBRACK),
	864:   uint16(anon_sym_LBRACE),
	865:   uint16(anon_sym_RBRACE),
	866:   uint16(anon_sym_LT_EQ),
	867:   uint16(anon_sym_EQ_EQ),
	868:   uint16(anon_sym_TILDE_EQ),
	869:   uint16(anon_sym_GT_EQ),
	870:   uint16(anon_sym_PIPE),
	871:   uint16(anon_sym_AMP),
	872:   uint16(anon_sym_LT_LT),
	873:   uint16(anon_sym_GT_GT),
	874:   uint16(anon_sym_PLUS),
	875:   uint16(anon_sym_SLASH_SLASH),
	876:   uint16(anon_sym_PERCENT),
	877:   uint16(anon_sym_DOT_DOT),
	878:   uint16(anon_sym_CARET),
	879:   uint16(5),
	880:   uint16(3),
	881:   uint16(1),
	882:   uint16(anon_sym_DASH_DASH),
	883:   uint16(5),
	884:   uint16(1),
	885:   uint16(sym__block_comment_start),
	886:   uint16(15),
	887:   uint16(1),
	888:   uint16(sym_comment),
	889:   uint16(99),
	890:   uint16(26),
	891:   uint16(anon_sym_return),
	892:   uint16(sym_break_statement),
	893:   uint16(anon_sym_goto),
	894:   uint16(anon_sym_do),
	895:   uint16(anon_sym_end),
	896:   uint16(anon_sym_while),
	897:   uint16(anon_sym_repeat),
	898:   uint16(anon_sym_until),
	899:   uint16(anon_sym_if),
	900:   uint16(anon_sym_then),
	901:   uint16(anon_sym_elseif),
	902:   uint16(anon_sym_else),
	903:   uint16(anon_sym_for),
	904:   uint16(anon_sym_function),
	905:   uint16(anon_sym_local),
	906:   uint16(anon_sym_global),
	907:   uint16(anon_sym_DOT),
	908:   uint16(anon_sym_COLON),
	909:   uint16(anon_sym_LT),
	910:   uint16(anon_sym_GT),
	911:   uint16(anon_sym_or),
	912:   uint16(anon_sym_and),
	913:   uint16(anon_sym_TILDE),
	914:   uint16(anon_sym_DASH),
	915:   uint16(anon_sym_SLASH),
	916:   uint16(sym_identifier),
	917:   uint16(97),
	918:   uint16(27),
	919:   uint16(sym__block_string_start),
	921:   uint16(anon_sym_SEMI),
	922:   uint16(anon_sym_COMMA),
	923:   uint16(anon_sym_COLON_COLON),
	924:   uint16(anon_sym_STAR),
	925:   uint16(anon_sym_DQUOTE),
	926:   uint16(anon_sym_SQUOTE),
	927:   uint16(anon_sym_LPAREN),
	928:   uint16(anon_sym_RPAREN),
	929:   uint16(anon_sym_LBRACK),
	930:   uint16(anon_sym_RBRACK),
	931:   uint16(anon_sym_LBRACE),
	932:   uint16(anon_sym_RBRACE),
	933:   uint16(anon_sym_LT_EQ),
	934:   uint16(anon_sym_EQ_EQ),
	935:   uint16(anon_sym_TILDE_EQ),
	936:   uint16(anon_sym_GT_EQ),
	937:   uint16(anon_sym_PIPE),
	938:   uint16(anon_sym_AMP),
	939:   uint16(anon_sym_LT_LT),
	940:   uint16(anon_sym_GT_GT),
	941:   uint16(anon_sym_PLUS),
	942:   uint16(anon_sym_SLASH_SLASH),
	943:   uint16(anon_sym_PERCENT),
	944:   uint16(anon_sym_DOT_DOT),
	945:   uint16(anon_sym_CARET),
	946:   uint16(5),
	947:   uint16(3),
	948:   uint16(1),
	949:   uint16(anon_sym_DASH_DASH),
	950:   uint16(5),
	951:   uint16(1),
	952:   uint16(sym__block_comment_start),
	953:   uint16(16),
	954:   uint16(1),
	955:   uint16(sym_comment),
	956:   uint16(95),
	957:   uint16(26),
	958:   uint16(anon_sym_return),
	959:   uint16(sym_break_statement),
	960:   uint16(anon_sym_goto),
	961:   uint16(anon_sym_do),
	962:   uint16(anon_sym_end),
	963:   uint16(anon_sym_while),
	964:   uint16(anon_sym_repeat),
	965:   uint16(anon_sym_until),
	966:   uint16(anon_sym_if),
	967:   uint16(anon_sym_then),
	968:   uint16(anon_sym_elseif),
	969:   uint16(anon_sym_else),
	970:   uint16(anon_sym_for),
	971:   uint16(anon_sym_function),
	972:   uint16(anon_sym_local),
	973:   uint16(anon_sym_global),
	974:   uint16(anon_sym_DOT),
	975:   uint16(anon_sym_COLON),
	976:   uint16(anon_sym_LT),
	977:   uint16(anon_sym_GT),
	978:   uint16(anon_sym_or),
	979:   uint16(anon_sym_and),
	980:   uint16(anon_sym_TILDE),
	981:   uint16(anon_sym_DASH),
	982:   uint16(anon_sym_SLASH),
	983:   uint16(sym_identifier),
	984:   uint16(93),
	985:   uint16(27),
	986:   uint16(sym__block_string_start),
	988:   uint16(anon_sym_SEMI),
	989:   uint16(anon_sym_COMMA),
	990:   uint16(anon_sym_COLON_COLON),
	991:   uint16(anon_sym_STAR),
	992:   uint16(anon_sym_DQUOTE),
	993:   uint16(anon_sym_SQUOTE),
	994:   uint16(anon_sym_LPAREN),
	995:   uint16(anon_sym_RPAREN),
	996:   uint16(anon_sym_LBRACK),
	997:   uint16(anon_sym_RBRACK),
	998:   uint16(anon_sym_LBRACE),
	999:   uint16(anon_sym_RBRACE),
	1000:  uint16(anon_sym_LT_EQ),
	1001:  uint16(anon_sym_EQ_EQ),
	1002:  uint16(anon_sym_TILDE_EQ),
	1003:  uint16(anon_sym_GT_EQ),
	1004:  uint16(anon_sym_PIPE),
	1005:  uint16(anon_sym_AMP),
	1006:  uint16(anon_sym_LT_LT),
	1007:  uint16(anon_sym_GT_GT),
	1008:  uint16(anon_sym_PLUS),
	1009:  uint16(anon_sym_SLASH_SLASH),
	1010:  uint16(anon_sym_PERCENT),
	1011:  uint16(anon_sym_DOT_DOT),
	1012:  uint16(anon_sym_CARET),
	1013:  uint16(5),
	1014:  uint16(3),
	1015:  uint16(1),
	1016:  uint16(anon_sym_DASH_DASH),
	1017:  uint16(5),
	1018:  uint16(1),
	1019:  uint16(sym__block_comment_start),
	1020:  uint16(17),
	1021:  uint16(1),
	1022:  uint16(sym_comment),
	1023:  uint16(103),
	1024:  uint16(26),
	1025:  uint16(anon_sym_return),
	1026:  uint16(sym_break_statement),
	1027:  uint16(anon_sym_goto),
	1028:  uint16(anon_sym_do),
	1029:  uint16(anon_sym_end),
	1030:  uint16(anon_sym_while),
	1031:  uint16(anon_sym_repeat),
	1032:  uint16(anon_sym_until),
	1033:  uint16(anon_sym_if),
	1034:  uint16(anon_sym_then),
	1035:  uint16(anon_sym_elseif),
	1036:  uint16(anon_sym_else),
	1037:  uint16(anon_sym_for),
	1038:  uint16(anon_sym_function),
	1039:  uint16(anon_sym_local),
	1040:  uint16(anon_sym_global),
	1041:  uint16(anon_sym_DOT),
	1042:  uint16(anon_sym_COLON),
	1043:  uint16(anon_sym_LT),
	1044:  uint16(anon_sym_GT),
	1045:  uint16(anon_sym_or),
	1046:  uint16(anon_sym_and),
	1047:  uint16(anon_sym_TILDE),
	1048:  uint16(anon_sym_DASH),
	1049:  uint16(anon_sym_SLASH),
	1050:  uint16(sym_identifier),
	1051:  uint16(101),
	1052:  uint16(27),
	1053:  uint16(sym__block_string_start),
	1055:  uint16(anon_sym_SEMI),
	1056:  uint16(anon_sym_COMMA),
	1057:  uint16(anon_sym_COLON_COLON),
	1058:  uint16(anon_sym_STAR),
	1059:  uint16(anon_sym_DQUOTE),
	1060:  uint16(anon_sym_SQUOTE),
	1061:  uint16(anon_sym_LPAREN),
	1062:  uint16(anon_sym_RPAREN),
	1063:  uint16(anon_sym_LBRACK),
	1064:  uint16(anon_sym_RBRACK),
	1065:  uint16(anon_sym_LBRACE),
	1066:  uint16(anon_sym_RBRACE),
	1067:  uint16(anon_sym_LT_EQ),
	1068:  uint16(anon_sym_EQ_EQ),
	1069:  uint16(anon_sym_TILDE_EQ),
	1070:  uint16(anon_sym_GT_EQ),
	1071:  uint16(anon_sym_PIPE),
	1072:  uint16(anon_sym_AMP),
	1073:  uint16(anon_sym_LT_LT),
	1074:  uint16(anon_sym_GT_GT),
	1075:  uint16(anon_sym_PLUS),
	1076:  uint16(anon_sym_SLASH_SLASH),
	1077:  uint16(anon_sym_PERCENT),
	1078:  uint16(anon_sym_DOT_DOT),
	1079:  uint16(anon_sym_CARET),
	1080:  uint16(5),
	1081:  uint16(3),
	1082:  uint16(1),
	1083:  uint16(anon_sym_DASH_DASH),
	1084:  uint16(5),
	1085:  uint16(1),
	1086:  uint16(sym__block_comment_start),
	1087:  uint16(18),
	1088:  uint16(1),
	1089:  uint16(sym_comment),
	1090:  uint16(107),
	1091:  uint16(26),
	1092:  uint16(anon_sym_return),
	1093:  uint16(sym_break_statement),
	1094:  uint16(anon_sym_goto),
	1095:  uint16(anon_sym_do),
	1096:  uint16(anon_sym_end),
	1097:  uint16(anon_sym_while),
	1098:  uint16(anon_sym_repeat),
	1099:  uint16(anon_sym_until),
	1100:  uint16(anon_sym_if),
	1101:  uint16(anon_sym_then),
	1102:  uint16(anon_sym_elseif),
	1103:  uint16(anon_sym_else),
	1104:  uint16(anon_sym_for),
	1105:  uint16(anon_sym_function),
	1106:  uint16(anon_sym_local),
	1107:  uint16(anon_sym_global),
	1108:  uint16(anon_sym_DOT),
	1109:  uint16(anon_sym_COLON),
	1110:  uint16(anon_sym_LT),
	1111:  uint16(anon_sym_GT),
	1112:  uint16(anon_sym_or),
	1113:  uint16(anon_sym_and),
	1114:  uint16(anon_sym_TILDE),
	1115:  uint16(anon_sym_DASH),
	1116:  uint16(anon_sym_SLASH),
	1117:  uint16(sym_identifier),
	1118:  uint16(105),
	1119:  uint16(27),
	1120:  uint16(sym__block_string_start),
	1122:  uint16(anon_sym_SEMI),
	1123:  uint16(anon_sym_COMMA),
	1124:  uint16(anon_sym_COLON_COLON),
	1125:  uint16(anon_sym_STAR),
	1126:  uint16(anon_sym_DQUOTE),
	1127:  uint16(anon_sym_SQUOTE),
	1128:  uint16(anon_sym_LPAREN),
	1129:  uint16(anon_sym_RPAREN),
	1130:  uint16(anon_sym_LBRACK),
	1131:  uint16(anon_sym_RBRACK),
	1132:  uint16(anon_sym_LBRACE),
	1133:  uint16(anon_sym_RBRACE),
	1134:  uint16(anon_sym_LT_EQ),
	1135:  uint16(anon_sym_EQ_EQ),
	1136:  uint16(anon_sym_TILDE_EQ),
	1137:  uint16(anon_sym_GT_EQ),
	1138:  uint16(anon_sym_PIPE),
	1139:  uint16(anon_sym_AMP),
	1140:  uint16(anon_sym_LT_LT),
	1141:  uint16(anon_sym_GT_GT),
	1142:  uint16(anon_sym_PLUS),
	1143:  uint16(anon_sym_SLASH_SLASH),
	1144:  uint16(anon_sym_PERCENT),
	1145:  uint16(anon_sym_DOT_DOT),
	1146:  uint16(anon_sym_CARET),
	1147:  uint16(5),
	1148:  uint16(3),
	1149:  uint16(1),
	1150:  uint16(anon_sym_DASH_DASH),
	1151:  uint16(5),
	1152:  uint16(1),
	1153:  uint16(sym__block_comment_start),
	1154:  uint16(19),
	1155:  uint16(1),
	1156:  uint16(sym_comment),
	1157:  uint16(111),
	1158:  uint16(26),
	1159:  uint16(anon_sym_return),
	1160:  uint16(sym_break_statement),
	1161:  uint16(anon_sym_goto),
	1162:  uint16(anon_sym_do),
	1163:  uint16(anon_sym_end),
	1164:  uint16(anon_sym_while),
	1165:  uint16(anon_sym_repeat),
	1166:  uint16(anon_sym_until),
	1167:  uint16(anon_sym_if),
	1168:  uint16(anon_sym_then),
	1169:  uint16(anon_sym_elseif),
	1170:  uint16(anon_sym_else),
	1171:  uint16(anon_sym_for),
	1172:  uint16(anon_sym_function),
	1173:  uint16(anon_sym_local),
	1174:  uint16(anon_sym_global),
	1175:  uint16(anon_sym_DOT),
	1176:  uint16(anon_sym_COLON),
	1177:  uint16(anon_sym_LT),
	1178:  uint16(anon_sym_GT),
	1179:  uint16(anon_sym_or),
	1180:  uint16(anon_sym_and),
	1181:  uint16(anon_sym_TILDE),
	1182:  uint16(anon_sym_DASH),
	1183:  uint16(anon_sym_SLASH),
	1184:  uint16(sym_identifier),
	1185:  uint16(109),
	1186:  uint16(27),
	1187:  uint16(sym__block_string_start),
	1189:  uint16(anon_sym_SEMI),
	1190:  uint16(anon_sym_COMMA),
	1191:  uint16(anon_sym_COLON_COLON),
	1192:  uint16(anon_sym_STAR),
	1193:  uint16(anon_sym_DQUOTE),
	1194:  uint16(anon_sym_SQUOTE),
	1195:  uint16(anon_sym_LPAREN),
	1196:  uint16(anon_sym_RPAREN),
	1197:  uint16(anon_sym_LBRACK),
	1198:  uint16(anon_sym_RBRACK),
	1199:  uint16(anon_sym_LBRACE),
	1200:  uint16(anon_sym_RBRACE),
	1201:  uint16(anon_sym_LT_EQ),
	1202:  uint16(anon_sym_EQ_EQ),
	1203:  uint16(anon_sym_TILDE_EQ),
	1204:  uint16(anon_sym_GT_EQ),
	1205:  uint16(anon_sym_PIPE),
	1206:  uint16(anon_sym_AMP),
	1207:  uint16(anon_sym_LT_LT),
	1208:  uint16(anon_sym_GT_GT),
	1209:  uint16(anon_sym_PLUS),
	1210:  uint16(anon_sym_SLASH_SLASH),
	1211:  uint16(anon_sym_PERCENT),
	1212:  uint16(anon_sym_DOT_DOT),
	1213:  uint16(anon_sym_CARET),
	1214:  uint16(43),
	1215:  uint16(3),
	1216:  uint16(1),
	1217:  uint16(anon_sym_DASH_DASH),
	1218:  uint16(5),
	1219:  uint16(1),
	1220:  uint16(sym__block_comment_start),
	1221:  uint16(9),
	1222:  uint16(1),
	1223:  uint16(sym_identifier),
	1224:  uint16(15),
	1225:  uint16(1),
	1226:  uint16(anon_sym_SEMI),
	1227:  uint16(17),
	1228:  uint16(1),
	1229:  uint16(anon_sym_COLON_COLON),
	1230:  uint16(19),
	1231:  uint16(1),
	1232:  uint16(sym_break_statement),
	1233:  uint16(21),
	1234:  uint16(1),
	1235:  uint16(anon_sym_goto),
	1236:  uint16(23),
	1237:  uint16(1),
	1238:  uint16(anon_sym_do),
	1239:  uint16(25),
	1240:  uint16(1),
	1241:  uint16(anon_sym_while),
	1242:  uint16(27),
	1243:  uint16(1),
	1244:  uint16(anon_sym_repeat),
	1245:  uint16(29),
	1246:  uint16(1),
	1247:  uint16(anon_sym_if),
	1248:  uint16(31),
	1249:  uint16(1),
	1250:  uint16(anon_sym_for),
	1251:  uint16(33),
	1252:  uint16(1),
	1253:  uint16(anon_sym_function),
	1254:  uint16(35),
	1255:  uint16(1),
	1256:  uint16(anon_sym_local),
	1257:  uint16(37),
	1258:  uint16(1),
	1259:  uint16(anon_sym_global),
	1260:  uint16(39),
	1261:  uint16(1),
	1262:  uint16(anon_sym_LPAREN),
	1263:  uint16(113),
	1264:  uint16(1),
	1265:  uint16(anon_sym_return),
	1266:  uint16(115),
	1267:  uint16(1),
	1268:  uint16(anon_sym_end),
	1269:  uint16(117),
	1270:  uint16(1),
	1271:  uint16(anon_sym_elseif),
	1272:  uint16(119),
	1273:  uint16(1),
	1274:  uint16(anon_sym_else),
	1275:  uint16(20),
	1276:  uint16(1),
	1277:  uint16(sym_comment),
	1278:  uint16(21),
	1279:  uint16(1),
	1280:  uint16(aux_sym_chunk_repeat1),
	1281:  uint16(101),
	1282:  uint16(1),
	1283:  uint16(sym_function_call),
	1284:  uint16(131),
	1285:  uint16(1),
	1286:  uint16(sym_variable_declaration),
	1287:  uint16(132),
	1288:  uint16(1),
	1289:  uint16(sym_declaration),
	1290:  uint16(134),
	1291:  uint16(1),
	1292:  uint16(sym_function_declaration),
	1293:  uint16(137),
	1294:  uint16(1),
	1295:  uint16(sym__local_function_declaration),
	1296:  uint16(139),
	1297:  uint16(1),
	1298:  uint16(sym__global_function_declaration),
	1299:  uint16(142),
	1300:  uint16(1),
	1301:  uint16(sym__global_variable_declaration),
	1302:  uint16(143),
	1303:  uint16(1),
	1304:  uint16(sym__global_implicit_variable_declaration),
	1305:  uint16(161),
	1306:  uint16(1),
	1307:  uint16(sym_statement),
	1308:  uint16(170),
	1309:  uint16(1),
	1310:  uint16(sym__prefix_expression),
	1311:  uint16(172),
	1312:  uint16(1),
	1313:  uint16(sym_variable),
	1314:  uint16(174),
	1315:  uint16(1),
	1316:  uint16(sym_method_index_expression),
	1317:  uint16(177),
	1318:  uint16(1),
	1319:  uint16(sym_parenthesized_expression),
	1320:  uint16(179),
	1321:  uint16(1),
	1322:  uint16(aux_sym_if_statement_repeat1),
	1323:  uint16(180),
	1324:  uint16(1),
	1325:  uint16(sym__block),
	1326:  uint16(194),
	1327:  uint16(1),
	1328:  uint16(sym_return_statement),
	1329:  uint16(213),
	1330:  uint16(1),
	1331:  uint16(sym_elseif_statement),
	1332:  uint16(233),
	1333:  uint16(1),
	1334:  uint16(sym_else_statement),
	1335:  uint16(255),
	1336:  uint16(1),
	1337:  uint16(sym__variable_assignment_varlist),
	1338:  uint16(2),
	1339:  uint16(3),
	1340:  uint16(sym_bracket_index_expression),
	1341:  uint16(sym_dot_index_expression),
	1342:  uint16(sym__contextual_keyword),
	1343:  uint16(160),
	1344:  uint16(9),
	1345:  uint16(sym_empty_statement),
	1346:  uint16(sym_assignment_statement),
	1347:  uint16(sym_label_statement),
	1348:  uint16(sym_goto_statement),
	1349:  uint16(sym_do_statement),
	1350:  uint16(sym_while_statement),
	1351:  uint16(sym_repeat_statement),
	1352:  uint16(sym_if_statement),
	1353:  uint16(sym_for_statement),
	1354:  uint16(37),
	1355:  uint16(3),
	1356:  uint16(1),
	1357:  uint16(anon_sym_DASH_DASH),
	1358:  uint16(5),
	1359:  uint16(1),
	1360:  uint16(sym__block_comment_start),
	1361:  uint16(9),
	1362:  uint16(1),
	1363:  uint16(sym_identifier),
	1364:  uint16(15),
	1365:  uint16(1),
	1366:  uint16(anon_sym_SEMI),
	1367:  uint16(17),
	1368:  uint16(1),
	1369:  uint16(anon_sym_COLON_COLON),
	1370:  uint16(19),
	1371:  uint16(1),
	1372:  uint16(sym_break_statement),
	1373:  uint16(21),
	1374:  uint16(1),
	1375:  uint16(anon_sym_goto),
	1376:  uint16(23),
	1377:  uint16(1),
	1378:  uint16(anon_sym_do),
	1379:  uint16(25),
	1380:  uint16(1),
	1381:  uint16(anon_sym_while),
	1382:  uint16(27),
	1383:  uint16(1),
	1384:  uint16(anon_sym_repeat),
	1385:  uint16(29),
	1386:  uint16(1),
	1387:  uint16(anon_sym_if),
	1388:  uint16(31),
	1389:  uint16(1),
	1390:  uint16(anon_sym_for),
	1391:  uint16(33),
	1392:  uint16(1),
	1393:  uint16(anon_sym_function),
	1394:  uint16(35),
	1395:  uint16(1),
	1396:  uint16(anon_sym_local),
	1397:  uint16(37),
	1398:  uint16(1),
	1399:  uint16(anon_sym_global),
	1400:  uint16(39),
	1401:  uint16(1),
	1402:  uint16(anon_sym_LPAREN),
	1403:  uint16(113),
	1404:  uint16(1),
	1405:  uint16(anon_sym_return),
	1406:  uint16(21),
	1407:  uint16(1),
	1408:  uint16(sym_comment),
	1409:  uint16(22),
	1410:  uint16(1),
	1411:  uint16(aux_sym_chunk_repeat1),
	1412:  uint16(101),
	1413:  uint16(1),
	1414:  uint16(sym_function_call),
	1415:  uint16(131),
	1416:  uint16(1),
	1417:  uint16(sym_variable_declaration),
	1418:  uint16(132),
	1419:  uint16(1),
	1420:  uint16(sym_declaration),
	1421:  uint16(134),
	1422:  uint16(1),
	1423:  uint16(sym_function_declaration),
	1424:  uint16(137),
	1425:  uint16(1),
	1426:  uint16(sym__local_function_declaration),
	1427:  uint16(139),
	1428:  uint16(1),
	1429:  uint16(sym__global_function_declaration),
	1430:  uint16(142),
	1431:  uint16(1),
	1432:  uint16(sym__global_variable_declaration),
	1433:  uint16(143),
	1434:  uint16(1),
	1435:  uint16(sym__global_implicit_variable_declaration),
	1436:  uint16(161),
	1437:  uint16(1),
	1438:  uint16(sym_statement),
	1439:  uint16(170),
	1440:  uint16(1),
	1441:  uint16(sym__prefix_expression),
	1442:  uint16(172),
	1443:  uint16(1),
	1444:  uint16(sym_variable),
	1445:  uint16(174),
	1446:  uint16(1),
	1447:  uint16(sym_method_index_expression),
	1448:  uint16(177),
	1449:  uint16(1),
	1450:  uint16(sym_parenthesized_expression),
	1451:  uint16(192),
	1452:  uint16(1),
	1453:  uint16(sym_return_statement),
	1454:  uint16(255),
	1455:  uint16(1),
	1456:  uint16(sym__variable_assignment_varlist),
	1457:  uint16(2),
	1458:  uint16(3),
	1459:  uint16(sym_bracket_index_expression),
	1460:  uint16(sym_dot_index_expression),
	1461:  uint16(sym__contextual_keyword),
	1462:  uint16(121),
	1463:  uint16(4),
	1464:  uint16(anon_sym_end),
	1465:  uint16(anon_sym_until),
	1466:  uint16(anon_sym_elseif),
	1467:  uint16(anon_sym_else),
	1468:  uint16(160),
	1469:  uint16(9),
	1470:  uint16(sym_empty_statement),
	1471:  uint16(sym_assignment_statement),
	1472:  uint16(sym_label_statement),
	1473:  uint16(sym_goto_statement),
	1474:  uint16(sym_do_statement),
	1475:  uint16(sym_while_statement),
	1476:  uint16(sym_repeat_statement),
	1477:  uint16(sym_if_statement),
	1478:  uint16(sym_for_statement),
	1479:  uint16(35),
	1480:  uint16(3),
	1481:  uint16(1),
	1482:  uint16(anon_sym_DASH_DASH),
	1483:  uint16(5),
	1484:  uint16(1),
	1485:  uint16(sym__block_comment_start),
	1486:  uint16(123),
	1487:  uint16(1),
	1489:  uint16(125),
	1490:  uint16(1),
	1491:  uint16(sym_identifier),
	1492:  uint16(130),
	1493:  uint16(1),
	1494:  uint16(anon_sym_SEMI),
	1495:  uint16(133),
	1496:  uint16(1),
	1497:  uint16(anon_sym_COLON_COLON),
	1498:  uint16(136),
	1499:  uint16(1),
	1500:  uint16(sym_break_statement),
	1501:  uint16(139),
	1502:  uint16(1),
	1503:  uint16(anon_sym_goto),
	1504:  uint16(142),
	1505:  uint16(1),
	1506:  uint16(anon_sym_do),
	1507:  uint16(145),
	1508:  uint16(1),
	1509:  uint16(anon_sym_while),
	1510:  uint16(148),
	1511:  uint16(1),
	1512:  uint16(anon_sym_repeat),
	1513:  uint16(151),
	1514:  uint16(1),
	1515:  uint16(anon_sym_if),
	1516:  uint16(154),
	1517:  uint16(1),
	1518:  uint16(anon_sym_for),
	1519:  uint16(157),
	1520:  uint16(1),
	1521:  uint16(anon_sym_function),
	1522:  uint16(160),
	1523:  uint16(1),
	1524:  uint16(anon_sym_local),
	1525:  uint16(163),
	1526:  uint16(1),
	1527:  uint16(anon_sym_global),
	1528:  uint16(166),
	1529:  uint16(1),
	1530:  uint16(anon_sym_LPAREN),
	1531:  uint16(101),
	1532:  uint16(1),
	1533:  uint16(sym_function_call),
	1534:  uint16(131),
	1535:  uint16(1),
	1536:  uint16(sym_variable_declaration),
	1537:  uint16(132),
	1538:  uint16(1),
	1539:  uint16(sym_declaration),
	1540:  uint16(134),
	1541:  uint16(1),
	1542:  uint16(sym_function_declaration),
	1543:  uint16(137),
	1544:  uint16(1),
	1545:  uint16(sym__local_function_declaration),
	1546:  uint16(139),
	1547:  uint16(1),
	1548:  uint16(sym__global_function_declaration),
	1549:  uint16(142),
	1550:  uint16(1),
	1551:  uint16(sym__global_variable_declaration),
	1552:  uint16(143),
	1553:  uint16(1),
	1554:  uint16(sym__global_implicit_variable_declaration),
	1555:  uint16(161),
	1556:  uint16(1),
	1557:  uint16(sym_statement),
	1558:  uint16(170),
	1559:  uint16(1),
	1560:  uint16(sym__prefix_expression),
	1561:  uint16(172),
	1562:  uint16(1),
	1563:  uint16(sym_variable),
	1564:  uint16(174),
	1565:  uint16(1),
	1566:  uint16(sym_method_index_expression),
	1567:  uint16(177),
	1568:  uint16(1),
	1569:  uint16(sym_parenthesized_expression),
	1570:  uint16(255),
	1571:  uint16(1),
	1572:  uint16(sym__variable_assignment_varlist),
	1573:  uint16(22),
	1574:  uint16(2),
	1575:  uint16(sym_comment),
	1576:  uint16(aux_sym_chunk_repeat1),
	1577:  uint16(2),
	1578:  uint16(3),
	1579:  uint16(sym_bracket_index_expression),
	1580:  uint16(sym_dot_index_expression),
	1581:  uint16(sym__contextual_keyword),
	1582:  uint16(128),
	1583:  uint16(5),
	1584:  uint16(anon_sym_return),
	1585:  uint16(anon_sym_end),
	1586:  uint16(anon_sym_until),
	1587:  uint16(anon_sym_elseif),
	1588:  uint16(anon_sym_else),
	1589:  uint16(160),
	1590:  uint16(9),
	1591:  uint16(sym_empty_statement),
	1592:  uint16(sym_assignment_statement),
	1593:  uint16(sym_label_statement),
	1594:  uint16(sym_goto_statement),
	1595:  uint16(sym_do_statement),
	1596:  uint16(sym_while_statement),
	1597:  uint16(sym_repeat_statement),
	1598:  uint16(sym_if_statement),
	1599:  uint16(sym_for_statement),
	1600:  uint16(38),
	1601:  uint16(3),
	1602:  uint16(1),
	1603:  uint16(anon_sym_DASH_DASH),
	1604:  uint16(5),
	1605:  uint16(1),
	1606:  uint16(sym__block_comment_start),
	1607:  uint16(9),
	1608:  uint16(1),
	1609:  uint16(sym_identifier),
	1610:  uint16(15),
	1611:  uint16(1),
	1612:  uint16(anon_sym_SEMI),
	1613:  uint16(17),
	1614:  uint16(1),
	1615:  uint16(anon_sym_COLON_COLON),
	1616:  uint16(19),
	1617:  uint16(1),
	1618:  uint16(sym_break_statement),
	1619:  uint16(21),
	1620:  uint16(1),
	1621:  uint16(anon_sym_goto),
	1622:  uint16(23),
	1623:  uint16(1),
	1624:  uint16(anon_sym_do),
	1625:  uint16(25),
	1626:  uint16(1),
	1627:  uint16(anon_sym_while),
	1628:  uint16(27),
	1629:  uint16(1),
	1630:  uint16(anon_sym_repeat),
	1631:  uint16(29),
	1632:  uint16(1),
	1633:  uint16(anon_sym_if),
	1634:  uint16(31),
	1635:  uint16(1),
	1636:  uint16(anon_sym_for),
	1637:  uint16(33),
	1638:  uint16(1),
	1639:  uint16(anon_sym_function),
	1640:  uint16(35),
	1641:  uint16(1),
	1642:  uint16(anon_sym_local),
	1643:  uint16(37),
	1644:  uint16(1),
	1645:  uint16(anon_sym_global),
	1646:  uint16(39),
	1647:  uint16(1),
	1648:  uint16(anon_sym_LPAREN),
	1649:  uint16(113),
	1650:  uint16(1),
	1651:  uint16(anon_sym_return),
	1652:  uint16(21),
	1653:  uint16(1),
	1654:  uint16(aux_sym_chunk_repeat1),
	1655:  uint16(23),
	1656:  uint16(1),
	1657:  uint16(sym_comment),
	1658:  uint16(101),
	1659:  uint16(1),
	1660:  uint16(sym_function_call),
	1661:  uint16(131),
	1662:  uint16(1),
	1663:  uint16(sym_variable_declaration),
	1664:  uint16(132),
	1665:  uint16(1),
	1666:  uint16(sym_declaration),
	1667:  uint16(134),
	1668:  uint16(1),
	1669:  uint16(sym_function_declaration),
	1670:  uint16(137),
	1671:  uint16(1),
	1672:  uint16(sym__local_function_declaration),
	1673:  uint16(139),
	1674:  uint16(1),
	1675:  uint16(sym__global_function_declaration),
	1676:  uint16(142),
	1677:  uint16(1),
	1678:  uint16(sym__global_variable_declaration),
	1679:  uint16(143),
	1680:  uint16(1),
	1681:  uint16(sym__global_implicit_variable_declaration),
	1682:  uint16(161),
	1683:  uint16(1),
	1684:  uint16(sym_statement),
	1685:  uint16(170),
	1686:  uint16(1),
	1687:  uint16(sym__prefix_expression),
	1688:  uint16(172),
	1689:  uint16(1),
	1690:  uint16(sym_variable),
	1691:  uint16(174),
	1692:  uint16(1),
	1693:  uint16(sym_method_index_expression),
	1694:  uint16(177),
	1695:  uint16(1),
	1696:  uint16(sym_parenthesized_expression),
	1697:  uint16(194),
	1698:  uint16(1),
	1699:  uint16(sym_return_statement),
	1700:  uint16(224),
	1701:  uint16(1),
	1702:  uint16(sym__block),
	1703:  uint16(255),
	1704:  uint16(1),
	1705:  uint16(sym__variable_assignment_varlist),
	1706:  uint16(169),
	1707:  uint16(3),
	1708:  uint16(anon_sym_end),
	1709:  uint16(anon_sym_elseif),
	1710:  uint16(anon_sym_else),
	1711:  uint16(2),
	1712:  uint16(3),
	1713:  uint16(sym_bracket_index_expression),
	1714:  uint16(sym_dot_index_expression),
	1715:  uint16(sym__contextual_keyword),
	1716:  uint16(160),
	1717:  uint16(9),
	1718:  uint16(sym_empty_statement),
	1719:  uint16(sym_assignment_statement),
	1720:  uint16(sym_label_statement),
	1721:  uint16(sym_goto_statement),
	1722:  uint16(sym_do_statement),
	1723:  uint16(sym_while_statement),
	1724:  uint16(sym_repeat_statement),
	1725:  uint16(sym_if_statement),
	1726:  uint16(sym_for_statement),
	1727:  uint16(5),
	1728:  uint16(3),
	1729:  uint16(1),
	1730:  uint16(anon_sym_DASH_DASH),
	1731:  uint16(5),
	1732:  uint16(1),
	1733:  uint16(sym__block_comment_start),
	1734:  uint16(24),
	1735:  uint16(1),
	1736:  uint16(sym_comment),
	1737:  uint16(105),
	1738:  uint16(22),
	1740:  uint16(anon_sym_SEMI),
	1741:  uint16(anon_sym_COMMA),
	1742:  uint16(anon_sym_COLON_COLON),
	1743:  uint16(anon_sym_STAR),
	1744:  uint16(anon_sym_LPAREN),
	1745:  uint16(anon_sym_RPAREN),
	1746:  uint16(anon_sym_RBRACK),
	1747:  uint16(anon_sym_RBRACE),
	1748:  uint16(anon_sym_LT_EQ),
	1749:  uint16(anon_sym_EQ_EQ),
	1750:  uint16(anon_sym_TILDE_EQ),
	1751:  uint16(anon_sym_GT_EQ),
	1752:  uint16(anon_sym_PIPE),
	1753:  uint16(anon_sym_AMP),
	1754:  uint16(anon_sym_LT_LT),
	1755:  uint16(anon_sym_GT_GT),
	1756:  uint16(anon_sym_PLUS),
	1757:  uint16(anon_sym_SLASH_SLASH),
	1758:  uint16(anon_sym_PERCENT),
	1759:  uint16(anon_sym_DOT_DOT),
	1760:  uint16(anon_sym_CARET),
	1761:  uint16(107),
	1762:  uint16(24),
	1763:  uint16(anon_sym_return),
	1764:  uint16(sym_break_statement),
	1765:  uint16(anon_sym_goto),
	1766:  uint16(anon_sym_do),
	1767:  uint16(anon_sym_end),
	1768:  uint16(anon_sym_while),
	1769:  uint16(anon_sym_repeat),
	1770:  uint16(anon_sym_until),
	1771:  uint16(anon_sym_if),
	1772:  uint16(anon_sym_then),
	1773:  uint16(anon_sym_elseif),
	1774:  uint16(anon_sym_else),
	1775:  uint16(anon_sym_for),
	1776:  uint16(anon_sym_function),
	1777:  uint16(anon_sym_local),
	1778:  uint16(anon_sym_global),
	1779:  uint16(anon_sym_LT),
	1780:  uint16(anon_sym_GT),
	1781:  uint16(anon_sym_or),
	1782:  uint16(anon_sym_and),
	1783:  uint16(anon_sym_TILDE),
	1784:  uint16(anon_sym_DASH),
	1785:  uint16(anon_sym_SLASH),
	1786:  uint16(sym_identifier),
	1787:  uint16(13),
	1788:  uint16(3),
	1789:  uint16(1),
	1790:  uint16(anon_sym_DASH_DASH),
	1791:  uint16(5),
	1792:  uint16(1),
	1793:  uint16(sym__block_comment_start),
	1794:  uint16(177),
	1795:  uint16(1),
	1796:  uint16(anon_sym_AMP),
	1797:  uint16(181),
	1798:  uint16(1),
	1799:  uint16(anon_sym_PLUS),
	1800:  uint16(183),
	1801:  uint16(1),
	1802:  uint16(anon_sym_DASH),
	1803:  uint16(185),
	1804:  uint16(1),
	1805:  uint16(anon_sym_SLASH),
	1806:  uint16(187),
	1807:  uint16(1),
	1808:  uint16(anon_sym_DOT_DOT),
	1809:  uint16(189),
	1810:  uint16(1),
	1811:  uint16(anon_sym_CARET),
	1812:  uint16(25),
	1813:  uint16(1),
	1814:  uint16(sym_comment),
	1815:  uint16(179),
	1816:  uint16(2),
	1817:  uint16(anon_sym_LT_LT),
	1818:  uint16(anon_sym_GT_GT),
	1819:  uint16(175),
	1820:  uint16(3),
	1821:  uint16(anon_sym_STAR),
	1822:  uint16(anon_sym_SLASH_SLASH),
	1823:  uint16(anon_sym_PERCENT),
	1824:  uint16(171),
	1825:  uint16(13),
	1827:  uint16(anon_sym_SEMI),
	1828:  uint16(anon_sym_COMMA),
	1829:  uint16(anon_sym_COLON_COLON),
	1830:  uint16(anon_sym_LPAREN),
	1831:  uint16(anon_sym_RPAREN),
	1832:  uint16(anon_sym_RBRACK),
	1833:  uint16(anon_sym_RBRACE),
	1834:  uint16(anon_sym_LT_EQ),
	1835:  uint16(anon_sym_EQ_EQ),
	1836:  uint16(anon_sym_TILDE_EQ),
	1837:  uint16(anon_sym_GT_EQ),
	1838:  uint16(anon_sym_PIPE),
	1839:  uint16(173),
	1840:  uint16(22),
	1841:  uint16(anon_sym_return),
	1842:  uint16(sym_break_statement),
	1843:  uint16(anon_sym_goto),
	1844:  uint16(anon_sym_do),
	1845:  uint16(anon_sym_end),
	1846:  uint16(anon_sym_while),
	1847:  uint16(anon_sym_repeat),
	1848:  uint16(anon_sym_until),
	1849:  uint16(anon_sym_if),
	1850:  uint16(anon_sym_then),
	1851:  uint16(anon_sym_elseif),
	1852:  uint16(anon_sym_else),
	1853:  uint16(anon_sym_for),
	1854:  uint16(anon_sym_function),
	1855:  uint16(anon_sym_local),
	1856:  uint16(anon_sym_global),
	1857:  uint16(anon_sym_LT),
	1858:  uint16(anon_sym_GT),
	1859:  uint16(anon_sym_or),
	1860:  uint16(anon_sym_and),
	1861:  uint16(anon_sym_TILDE),
	1862:  uint16(sym_identifier),
	1863:  uint16(17),
	1864:  uint16(3),
	1865:  uint16(1),
	1866:  uint16(anon_sym_DASH_DASH),
	1867:  uint16(5),
	1868:  uint16(1),
	1869:  uint16(sym__block_comment_start),
	1870:  uint16(177),
	1871:  uint16(1),
	1872:  uint16(anon_sym_AMP),
	1873:  uint16(181),
	1874:  uint16(1),
	1875:  uint16(anon_sym_PLUS),
	1876:  uint16(183),
	1877:  uint16(1),
	1878:  uint16(anon_sym_DASH),
	1879:  uint16(185),
	1880:  uint16(1),
	1881:  uint16(anon_sym_SLASH),
	1882:  uint16(187),
	1883:  uint16(1),
	1884:  uint16(anon_sym_DOT_DOT),
	1885:  uint16(189),
	1886:  uint16(1),
	1887:  uint16(anon_sym_CARET),
	1888:  uint16(195),
	1889:  uint16(1),
	1890:  uint16(anon_sym_PIPE),
	1891:  uint16(197),
	1892:  uint16(1),
	1893:  uint16(anon_sym_TILDE),
	1894:  uint16(26),
	1895:  uint16(1),
	1896:  uint16(sym_comment),
	1897:  uint16(179),
	1898:  uint16(2),
	1899:  uint16(anon_sym_LT_LT),
	1900:  uint16(anon_sym_GT_GT),
	1901:  uint16(191),
	1902:  uint16(2),
	1903:  uint16(anon_sym_LT),
	1904:  uint16(anon_sym_GT),
	1905:  uint16(175),
	1906:  uint16(3),
	1907:  uint16(anon_sym_STAR),
	1908:  uint16(anon_sym_SLASH_SLASH),
	1909:  uint16(anon_sym_PERCENT),
	1910:  uint16(193),
	1911:  uint16(4),
	1912:  uint16(anon_sym_LT_EQ),
	1913:  uint16(anon_sym_EQ_EQ),
	1914:  uint16(anon_sym_TILDE_EQ),
	1915:  uint16(anon_sym_GT_EQ),
	1916:  uint16(171),
	1917:  uint16(8),
	1919:  uint16(anon_sym_SEMI),
	1920:  uint16(anon_sym_COMMA),
	1921:  uint16(anon_sym_COLON_COLON),
	1922:  uint16(anon_sym_LPAREN),
	1923:  uint16(anon_sym_RPAREN),
	1924:  uint16(anon_sym_RBRACK),
	1925:  uint16(anon_sym_RBRACE),
	1926:  uint16(173),
	1927:  uint16(19),
	1928:  uint16(anon_sym_return),
	1929:  uint16(sym_break_statement),
	1930:  uint16(anon_sym_goto),
	1931:  uint16(anon_sym_do),
	1932:  uint16(anon_sym_end),
	1933:  uint16(anon_sym_while),
	1934:  uint16(anon_sym_repeat),
	1935:  uint16(anon_sym_until),
	1936:  uint16(anon_sym_if),
	1937:  uint16(anon_sym_then),
	1938:  uint16(anon_sym_elseif),
	1939:  uint16(anon_sym_else),
	1940:  uint16(anon_sym_for),
	1941:  uint16(anon_sym_function),
	1942:  uint16(anon_sym_local),
	1943:  uint16(anon_sym_global),
	1944:  uint16(anon_sym_or),
	1945:  uint16(anon_sym_and),
	1946:  uint16(sym_identifier),
	1947:  uint16(18),
	1948:  uint16(3),
	1949:  uint16(1),
	1950:  uint16(anon_sym_DASH_DASH),
	1951:  uint16(5),
	1952:  uint16(1),
	1953:  uint16(sym__block_comment_start),
	1954:  uint16(177),
	1955:  uint16(1),
	1956:  uint16(anon_sym_AMP),
	1957:  uint16(181),
	1958:  uint16(1),
	1959:  uint16(anon_sym_PLUS),
	1960:  uint16(183),
	1961:  uint16(1),
	1962:  uint16(anon_sym_DASH),
	1963:  uint16(185),
	1964:  uint16(1),
	1965:  uint16(anon_sym_SLASH),
	1966:  uint16(187),
	1967:  uint16(1),
	1968:  uint16(anon_sym_DOT_DOT),
	1969:  uint16(189),
	1970:  uint16(1),
	1971:  uint16(anon_sym_CARET),
	1972:  uint16(195),
	1973:  uint16(1),
	1974:  uint16(anon_sym_PIPE),
	1975:  uint16(197),
	1976:  uint16(1),
	1977:  uint16(anon_sym_TILDE),
	1978:  uint16(199),
	1979:  uint16(1),
	1980:  uint16(anon_sym_and),
	1981:  uint16(27),
	1982:  uint16(1),
	1983:  uint16(sym_comment),
	1984:  uint16(179),
	1985:  uint16(2),
	1986:  uint16(anon_sym_LT_LT),
	1987:  uint16(anon_sym_GT_GT),
	1988:  uint16(191),
	1989:  uint16(2),
	1990:  uint16(anon_sym_LT),
	1991:  uint16(anon_sym_GT),
	1992:  uint16(175),
	1993:  uint16(3),
	1994:  uint16(anon_sym_STAR),
	1995:  uint16(anon_sym_SLASH_SLASH),
	1996:  uint16(anon_sym_PERCENT),
	1997:  uint16(193),
	1998:  uint16(4),
	1999:  uint16(anon_sym_LT_EQ),
	2000:  uint16(anon_sym_EQ_EQ),
	2001:  uint16(anon_sym_TILDE_EQ),
	2002:  uint16(anon_sym_GT_EQ),
	2003:  uint16(171),
	2004:  uint16(8),
	2006:  uint16(anon_sym_SEMI),
	2007:  uint16(anon_sym_COMMA),
	2008:  uint16(anon_sym_COLON_COLON),
	2009:  uint16(anon_sym_LPAREN),
	2010:  uint16(anon_sym_RPAREN),
	2011:  uint16(anon_sym_RBRACK),
	2012:  uint16(anon_sym_RBRACE),
	2013:  uint16(173),
	2014:  uint16(18),
	2015:  uint16(anon_sym_return),
	2016:  uint16(sym_break_statement),
	2017:  uint16(anon_sym_goto),
	2018:  uint16(anon_sym_do),
	2019:  uint16(anon_sym_end),
	2020:  uint16(anon_sym_while),
	2021:  uint16(anon_sym_repeat),
	2022:  uint16(anon_sym_until),
	2023:  uint16(anon_sym_if),
	2024:  uint16(anon_sym_then),
	2025:  uint16(anon_sym_elseif),
	2026:  uint16(anon_sym_else),
	2027:  uint16(anon_sym_for),
	2028:  uint16(anon_sym_function),
	2029:  uint16(anon_sym_local),
	2030:  uint16(anon_sym_global),
	2031:  uint16(anon_sym_or),
	2032:  uint16(sym_identifier),
	2033:  uint16(6),
	2034:  uint16(3),
	2035:  uint16(1),
	2036:  uint16(anon_sym_DASH_DASH),
	2037:  uint16(5),
	2038:  uint16(1),
	2039:  uint16(sym__block_comment_start),
	2040:  uint16(189),
	2041:  uint16(1),
	2042:  uint16(anon_sym_CARET),
	2043:  uint16(28),
	2044:  uint16(1),
	2045:  uint16(sym_comment),
	2046:  uint16(171),
	2047:  uint16(21),
	2049:  uint16(anon_sym_SEMI),
	2050:  uint16(anon_sym_COMMA),
	2051:  uint16(anon_sym_COLON_COLON),
	2052:  uint16(anon_sym_STAR),
	2053:  uint16(anon_sym_LPAREN),
	2054:  uint16(anon_sym_RPAREN),
	2055:  uint16(anon_sym_RBRACK),
	2056:  uint16(anon_sym_RBRACE),
	2057:  uint16(anon_sym_LT_EQ),
	2058:  uint16(anon_sym_EQ_EQ),
	2059:  uint16(anon_sym_TILDE_EQ),
	2060:  uint16(anon_sym_GT_EQ),
	2061:  uint16(anon_sym_PIPE),
	2062:  uint16(anon_sym_AMP),
	2063:  uint16(anon_sym_LT_LT),
	2064:  uint16(anon_sym_GT_GT),
	2065:  uint16(anon_sym_PLUS),
	2066:  uint16(anon_sym_SLASH_SLASH),
	2067:  uint16(anon_sym_PERCENT),
	2068:  uint16(anon_sym_DOT_DOT),
	2069:  uint16(173),
	2070:  uint16(24),
	2071:  uint16(anon_sym_return),
	2072:  uint16(sym_break_statement),
	2073:  uint16(anon_sym_goto),
	2074:  uint16(anon_sym_do),
	2075:  uint16(anon_sym_end),
	2076:  uint16(anon_sym_while),
	2077:  uint16(anon_sym_repeat),
	2078:  uint16(anon_sym_until),
	2079:  uint16(anon_sym_if),
	2080:  uint16(anon_sym_then),
	2081:  uint16(anon_sym_elseif),
	2082:  uint16(anon_sym_else),
	2083:  uint16(anon_sym_for),
	2084:  uint16(anon_sym_function),
	2085:  uint16(anon_sym_local),
	2086:  uint16(anon_sym_global),
	2087:  uint16(anon_sym_LT),
	2088:  uint16(anon_sym_GT),
	2089:  uint16(anon_sym_or),
	2090:  uint16(anon_sym_and),
	2091:  uint16(anon_sym_TILDE),
	2092:  uint16(anon_sym_DASH),
	2093:  uint16(anon_sym_SLASH),
	2094:  uint16(sym_identifier),
	2095:  uint16(5),
	2096:  uint16(3),
	2097:  uint16(1),
	2098:  uint16(anon_sym_DASH_DASH),
	2099:  uint16(5),
	2100:  uint16(1),
	2101:  uint16(sym__block_comment_start),
	2102:  uint16(29),
	2103:  uint16(1),
	2104:  uint16(sym_comment),
	2105:  uint16(201),
	2106:  uint16(22),
	2108:  uint16(anon_sym_SEMI),
	2109:  uint16(anon_sym_COMMA),
	2110:  uint16(anon_sym_COLON_COLON),
	2111:  uint16(anon_sym_STAR),
	2112:  uint16(anon_sym_LPAREN),
	2113:  uint16(anon_sym_RPAREN),
	2114:  uint16(anon_sym_RBRACK),
	2115:  uint16(anon_sym_RBRACE),
	2116:  uint16(anon_sym_LT_EQ),
	2117:  uint16(anon_sym_EQ_EQ),
	2118:  uint16(anon_sym_TILDE_EQ),
	2119:  uint16(anon_sym_GT_EQ),
	2120:  uint16(anon_sym_PIPE),
	2121:  uint16(anon_sym_AMP),
	2122:  uint16(anon_sym_LT_LT),
	2123:  uint16(anon_sym_GT_GT),
	2124:  uint16(anon_sym_PLUS),
	2125:  uint16(anon_sym_SLASH_SLASH),
	2126:  uint16(anon_sym_PERCENT),
	2127:  uint16(anon_sym_DOT_DOT),
	2128:  uint16(anon_sym_CARET),
	2129:  uint16(203),
	2130:  uint16(24),
	2131:  uint16(anon_sym_return),
	2132:  uint16(sym_break_statement),
	2133:  uint16(anon_sym_goto),
	2134:  uint16(anon_sym_do),
	2135:  uint16(anon_sym_end),
	2136:  uint16(anon_sym_while),
	2137:  uint16(anon_sym_repeat),
	2138:  uint16(anon_sym_until),
	2139:  uint16(anon_sym_if),
	2140:  uint16(anon_sym_then),
	2141:  uint16(anon_sym_elseif),
	2142:  uint16(anon_sym_else),
	2143:  uint16(anon_sym_for),
	2144:  uint16(anon_sym_function),
	2145:  uint16(anon_sym_local),
	2146:  uint16(anon_sym_global),
	2147:  uint16(anon_sym_LT),
	2148:  uint16(anon_sym_GT),
	2149:  uint16(anon_sym_or),
	2150:  uint16(anon_sym_and),
	2151:  uint16(anon_sym_TILDE),
	2152:  uint16(anon_sym_DASH),
	2153:  uint16(anon_sym_SLASH),
	2154:  uint16(sym_identifier),
	2155:  uint16(5),
	2156:  uint16(3),
	2157:  uint16(1),
	2158:  uint16(anon_sym_DASH_DASH),
	2159:  uint16(5),
	2160:  uint16(1),
	2161:  uint16(sym__block_comment_start),
	2162:  uint16(30),
	2163:  uint16(1),
	2164:  uint16(sym_comment),
	2165:  uint16(109),
	2166:  uint16(22),
	2168:  uint16(anon_sym_SEMI),
	2169:  uint16(anon_sym_COMMA),
	2170:  uint16(anon_sym_COLON_COLON),
	2171:  uint16(anon_sym_STAR),
	2172:  uint16(anon_sym_LPAREN),
	2173:  uint16(anon_sym_RPAREN),
	2174:  uint16(anon_sym_RBRACK),
	2175:  uint16(anon_sym_RBRACE),
	2176:  uint16(anon_sym_LT_EQ),
	2177:  uint16(anon_sym_EQ_EQ),
	2178:  uint16(anon_sym_TILDE_EQ),
	2179:  uint16(anon_sym_GT_EQ),
	2180:  uint16(anon_sym_PIPE),
	2181:  uint16(anon_sym_AMP),
	2182:  uint16(anon_sym_LT_LT),
	2183:  uint16(anon_sym_GT_GT),
	2184:  uint16(anon_sym_PLUS),
	2185:  uint16(anon_sym_SLASH_SLASH),
	2186:  uint16(anon_sym_PERCENT),
	2187:  uint16(anon_sym_DOT_DOT),
	2188:  uint16(anon_sym_CARET),
	2189:  uint16(111),
	2190:  uint16(24),
	2191:  uint16(anon_sym_return),
	2192:  uint16(sym_break_statement),
	2193:  uint16(anon_sym_goto),
	2194:  uint16(anon_sym_do),
	2195:  uint16(anon_sym_end),
	2196:  uint16(anon_sym_while),
	2197:  uint16(anon_sym_repeat),
	2198:  uint16(anon_sym_until),
	2199:  uint16(anon_sym_if),
	2200:  uint16(anon_sym_then),
	2201:  uint16(anon_sym_elseif),
	2202:  uint16(anon_sym_else),
	2203:  uint16(anon_sym_for),
	2204:  uint16(anon_sym_function),
	2205:  uint16(anon_sym_local),
	2206:  uint16(anon_sym_global),
	2207:  uint16(anon_sym_LT),
	2208:  uint16(anon_sym_GT),
	2209:  uint16(anon_sym_or),
	2210:  uint16(anon_sym_and),
	2211:  uint16(anon_sym_TILDE),
	2212:  uint16(anon_sym_DASH),
	2213:  uint16(anon_sym_SLASH),
	2214:  uint16(sym_identifier),
	2215:  uint16(5),
	2216:  uint16(3),
	2217:  uint16(1),
	2218:  uint16(anon_sym_DASH_DASH),
	2219:  uint16(5),
	2220:  uint16(1),
	2221:  uint16(sym__block_comment_start),
	2222:  uint16(31),
	2223:  uint16(1),
	2224:  uint16(sym_comment),
	2225:  uint16(57),
	2226:  uint16(22),
	2228:  uint16(anon_sym_SEMI),
	2229:  uint16(anon_sym_COMMA),
	2230:  uint16(anon_sym_COLON_COLON),
	2231:  uint16(anon_sym_STAR),
	2232:  uint16(anon_sym_LPAREN),
	2233:  uint16(anon_sym_RPAREN),
	2234:  uint16(anon_sym_RBRACK),
	2235:  uint16(anon_sym_RBRACE),
	2236:  uint16(anon_sym_LT_EQ),
	2237:  uint16(anon_sym_EQ_EQ),
	2238:  uint16(anon_sym_TILDE_EQ),
	2239:  uint16(anon_sym_GT_EQ),
	2240:  uint16(anon_sym_PIPE),
	2241:  uint16(anon_sym_AMP),
	2242:  uint16(anon_sym_LT_LT),
	2243:  uint16(anon_sym_GT_GT),
	2244:  uint16(anon_sym_PLUS),
	2245:  uint16(anon_sym_SLASH_SLASH),
	2246:  uint16(anon_sym_PERCENT),
	2247:  uint16(anon_sym_DOT_DOT),
	2248:  uint16(anon_sym_CARET),
	2249:  uint16(59),
	2250:  uint16(24),
	2251:  uint16(anon_sym_return),
	2252:  uint16(sym_break_statement),
	2253:  uint16(anon_sym_goto),
	2254:  uint16(anon_sym_do),
	2255:  uint16(anon_sym_end),
	2256:  uint16(anon_sym_while),
	2257:  uint16(anon_sym_repeat),
	2258:  uint16(anon_sym_until),
	2259:  uint16(anon_sym_if),
	2260:  uint16(anon_sym_then),
	2261:  uint16(anon_sym_elseif),
	2262:  uint16(anon_sym_else),
	2263:  uint16(anon_sym_for),
	2264:  uint16(anon_sym_function),
	2265:  uint16(anon_sym_local),
	2266:  uint16(anon_sym_global),
	2267:  uint16(anon_sym_LT),
	2268:  uint16(anon_sym_GT),
	2269:  uint16(anon_sym_or),
	2270:  uint16(anon_sym_and),
	2271:  uint16(anon_sym_TILDE),
	2272:  uint16(anon_sym_DASH),
	2273:  uint16(anon_sym_SLASH),
	2274:  uint16(sym_identifier),
	2275:  uint16(15),
	2276:  uint16(3),
	2277:  uint16(1),
	2278:  uint16(anon_sym_DASH_DASH),
	2279:  uint16(5),
	2280:  uint16(1),
	2281:  uint16(sym__block_comment_start),
	2282:  uint16(177),
	2283:  uint16(1),
	2284:  uint16(anon_sym_AMP),
	2285:  uint16(181),
	2286:  uint16(1),
	2287:  uint16(anon_sym_PLUS),
	2288:  uint16(183),
	2289:  uint16(1),
	2290:  uint16(anon_sym_DASH),
	2291:  uint16(185),
	2292:  uint16(1),
	2293:  uint16(anon_sym_SLASH),
	2294:  uint16(187),
	2295:  uint16(1),
	2296:  uint16(anon_sym_DOT_DOT),
	2297:  uint16(189),
	2298:  uint16(1),
	2299:  uint16(anon_sym_CARET),
	2300:  uint16(195),
	2301:  uint16(1),
	2302:  uint16(anon_sym_PIPE),
	2303:  uint16(197),
	2304:  uint16(1),
	2305:  uint16(anon_sym_TILDE),
	2306:  uint16(32),
	2307:  uint16(1),
	2308:  uint16(sym_comment),
	2309:  uint16(179),
	2310:  uint16(2),
	2311:  uint16(anon_sym_LT_LT),
	2312:  uint16(anon_sym_GT_GT),
	2313:  uint16(175),
	2314:  uint16(3),
	2315:  uint16(anon_sym_STAR),
	2316:  uint16(anon_sym_SLASH_SLASH),
	2317:  uint16(anon_sym_PERCENT),
	2318:  uint16(171),
	2319:  uint16(12),
	2321:  uint16(anon_sym_SEMI),
	2322:  uint16(anon_sym_COMMA),
	2323:  uint16(anon_sym_COLON_COLON),
	2324:  uint16(anon_sym_LPAREN),
	2325:  uint16(anon_sym_RPAREN),
	2326:  uint16(anon_sym_RBRACK),
	2327:  uint16(anon_sym_RBRACE),
	2328:  uint16(anon_sym_LT_EQ),
	2329:  uint16(anon_sym_EQ_EQ),
	2330:  uint16(anon_sym_TILDE_EQ),
	2331:  uint16(anon_sym_GT_EQ),
	2332:  uint16(173),
	2333:  uint16(21),
	2334:  uint16(anon_sym_return),
	2335:  uint16(sym_break_statement),
	2336:  uint16(anon_sym_goto),
	2337:  uint16(anon_sym_do),
	2338:  uint16(anon_sym_end),
	2339:  uint16(anon_sym_while),
	2340:  uint16(anon_sym_repeat),
	2341:  uint16(anon_sym_until),
	2342:  uint16(anon_sym_if),
	2343:  uint16(anon_sym_then),
	2344:  uint16(anon_sym_elseif),
	2345:  uint16(anon_sym_else),
	2346:  uint16(anon_sym_for),
	2347:  uint16(anon_sym_function),
	2348:  uint16(anon_sym_local),
	2349:  uint16(anon_sym_global),
	2350:  uint16(anon_sym_LT),
	2351:  uint16(anon_sym_GT),
	2352:  uint16(anon_sym_or),
	2353:  uint16(anon_sym_and),
	2354:  uint16(sym_identifier),
	2355:  uint16(14),
	2356:  uint16(3),
	2357:  uint16(1),
	2358:  uint16(anon_sym_DASH_DASH),
	2359:  uint16(5),
	2360:  uint16(1),
	2361:  uint16(sym__block_comment_start),
	2362:  uint16(177),
	2363:  uint16(1),
	2364:  uint16(anon_sym_AMP),
	2365:  uint16(181),
	2366:  uint16(1),
	2367:  uint16(anon_sym_PLUS),
	2368:  uint16(183),
	2369:  uint16(1),
	2370:  uint16(anon_sym_DASH),
	2371:  uint16(185),
	2372:  uint16(1),
	2373:  uint16(anon_sym_SLASH),
	2374:  uint16(187),
	2375:  uint16(1),
	2376:  uint16(anon_sym_DOT_DOT),
	2377:  uint16(189),
	2378:  uint16(1),
	2379:  uint16(anon_sym_CARET),
	2380:  uint16(197),
	2381:  uint16(1),
	2382:  uint16(anon_sym_TILDE),
	2383:  uint16(33),
	2384:  uint16(1),
	2385:  uint16(sym_comment),
	2386:  uint16(179),
	2387:  uint16(2),
	2388:  uint16(anon_sym_LT_LT),
	2389:  uint16(anon_sym_GT_GT),
	2390:  uint16(175),
	2391:  uint16(3),
	2392:  uint16(anon_sym_STAR),
	2393:  uint16(anon_sym_SLASH_SLASH),
	2394:  uint16(anon_sym_PERCENT),
	2395:  uint16(171),
	2396:  uint16(13),
	2398:  uint16(anon_sym_SEMI),
	2399:  uint16(anon_sym_COMMA),
	2400:  uint16(anon_sym_COLON_COLON),
	2401:  uint16(anon_sym_LPAREN),
	2402:  uint16(anon_sym_RPAREN),
	2403:  uint16(anon_sym_RBRACK),
	2404:  uint16(anon_sym_RBRACE),
	2405:  uint16(anon_sym_LT_EQ),
	2406:  uint16(anon_sym_EQ_EQ),
	2407:  uint16(anon_sym_TILDE_EQ),
	2408:  uint16(anon_sym_GT_EQ),
	2409:  uint16(anon_sym_PIPE),
	2410:  uint16(173),
	2411:  uint16(21),
	2412:  uint16(anon_sym_return),
	2413:  uint16(sym_break_statement),
	2414:  uint16(anon_sym_goto),
	2415:  uint16(anon_sym_do),
	2416:  uint16(anon_sym_end),
	2417:  uint16(anon_sym_while),
	2418:  uint16(anon_sym_repeat),
	2419:  uint16(anon_sym_until),
	2420:  uint16(anon_sym_if),
	2421:  uint16(anon_sym_then),
	2422:  uint16(anon_sym_elseif),
	2423:  uint16(anon_sym_else),
	2424:  uint16(anon_sym_for),
	2425:  uint16(anon_sym_function),
	2426:  uint16(anon_sym_local),
	2427:  uint16(anon_sym_global),
	2428:  uint16(anon_sym_LT),
	2429:  uint16(anon_sym_GT),
	2430:  uint16(anon_sym_or),
	2431:  uint16(anon_sym_and),
	2432:  uint16(sym_identifier),
	2433:  uint16(12),
	2434:  uint16(3),
	2435:  uint16(1),
	2436:  uint16(anon_sym_DASH_DASH),
	2437:  uint16(5),
	2438:  uint16(1),
	2439:  uint16(sym__block_comment_start),
	2440:  uint16(181),
	2441:  uint16(1),
	2442:  uint16(anon_sym_PLUS),
	2443:  uint16(183),
	2444:  uint16(1),
	2445:  uint16(anon_sym_DASH),
	2446:  uint16(185),
	2447:  uint16(1),
	2448:  uint16(anon_sym_SLASH),
	2449:  uint16(187),
	2450:  uint16(1),
	2451:  uint16(anon_sym_DOT_DOT),
	2452:  uint16(189),
	2453:  uint16(1),
	2454:  uint16(anon_sym_CARET),
	2455:  uint16(34),
	2456:  uint16(1),
	2457:  uint16(sym_comment),
	2458:  uint16(179),
	2459:  uint16(2),
	2460:  uint16(anon_sym_LT_LT),
	2461:  uint16(anon_sym_GT_GT),
	2462:  uint16(175),
	2463:  uint16(3),
	2464:  uint16(anon_sym_STAR),
	2465:  uint16(anon_sym_SLASH_SLASH),
	2466:  uint16(anon_sym_PERCENT),
	2467:  uint16(171),
	2468:  uint16(14),
	2470:  uint16(anon_sym_SEMI),
	2471:  uint16(anon_sym_COMMA),
	2472:  uint16(anon_sym_COLON_COLON),
	2473:  uint16(anon_sym_LPAREN),
	2474:  uint16(anon_sym_RPAREN),
	2475:  uint16(anon_sym_RBRACK),
	2476:  uint16(anon_sym_RBRACE),
	2477:  uint16(anon_sym_LT_EQ),
	2478:  uint16(anon_sym_EQ_EQ),
	2479:  uint16(anon_sym_TILDE_EQ),
	2480:  uint16(anon_sym_GT_EQ),
	2481:  uint16(anon_sym_PIPE),
	2482:  uint16(anon_sym_AMP),
	2483:  uint16(173),
	2484:  uint16(22),
	2485:  uint16(anon_sym_return),
	2486:  uint16(sym_break_statement),
	2487:  uint16(anon_sym_goto),
	2488:  uint16(anon_sym_do),
	2489:  uint16(anon_sym_end),
	2490:  uint16(anon_sym_while),
	2491:  uint16(anon_sym_repeat),
	2492:  uint16(anon_sym_until),
	2493:  uint16(anon_sym_if),
	2494:  uint16(anon_sym_then),
	2495:  uint16(anon_sym_elseif),
	2496:  uint16(anon_sym_else),
	2497:  uint16(anon_sym_for),
	2498:  uint16(anon_sym_function),
	2499:  uint16(anon_sym_local),
	2500:  uint16(anon_sym_global),
	2501:  uint16(anon_sym_LT),
	2502:  uint16(anon_sym_GT),
	2503:  uint16(anon_sym_or),
	2504:  uint16(anon_sym_and),
	2505:  uint16(anon_sym_TILDE),
	2506:  uint16(sym_identifier),
	2507:  uint16(11),
	2508:  uint16(3),
	2509:  uint16(1),
	2510:  uint16(anon_sym_DASH_DASH),
	2511:  uint16(5),
	2512:  uint16(1),
	2513:  uint16(sym__block_comment_start),
	2514:  uint16(181),
	2515:  uint16(1),
	2516:  uint16(anon_sym_PLUS),
	2517:  uint16(183),
	2518:  uint16(1),
	2519:  uint16(anon_sym_DASH),
	2520:  uint16(185),
	2521:  uint16(1),
	2522:  uint16(anon_sym_SLASH),
	2523:  uint16(187),
	2524:  uint16(1),
	2525:  uint16(anon_sym_DOT_DOT),
	2526:  uint16(189),
	2527:  uint16(1),
	2528:  uint16(anon_sym_CARET),
	2529:  uint16(35),
	2530:  uint16(1),
	2531:  uint16(sym_comment),
	2532:  uint16(175),
	2533:  uint16(3),
	2534:  uint16(anon_sym_STAR),
	2535:  uint16(anon_sym_SLASH_SLASH),
	2536:  uint16(anon_sym_PERCENT),
	2537:  uint16(171),
	2538:  uint16(16),
	2540:  uint16(anon_sym_SEMI),
	2541:  uint16(anon_sym_COMMA),
	2542:  uint16(anon_sym_COLON_COLON),
	2543:  uint16(anon_sym_LPAREN),
	2544:  uint16(anon_sym_RPAREN),
	2545:  uint16(anon_sym_RBRACK),
	2546:  uint16(anon_sym_RBRACE),
	2547:  uint16(anon_sym_LT_EQ),
	2548:  uint16(anon_sym_EQ_EQ),
	2549:  uint16(anon_sym_TILDE_EQ),
	2550:  uint16(anon_sym_GT_EQ),
	2551:  uint16(anon_sym_PIPE),
	2552:  uint16(anon_sym_AMP),
	2553:  uint16(anon_sym_LT_LT),
	2554:  uint16(anon_sym_GT_GT),
	2555:  uint16(173),
	2556:  uint16(22),
	2557:  uint16(anon_sym_return),
	2558:  uint16(sym_break_statement),
	2559:  uint16(anon_sym_goto),
	2560:  uint16(anon_sym_do),
	2561:  uint16(anon_sym_end),
	2562:  uint16(anon_sym_while),
	2563:  uint16(anon_sym_repeat),
	2564:  uint16(anon_sym_until),
	2565:  uint16(anon_sym_if),
	2566:  uint16(anon_sym_then),
	2567:  uint16(anon_sym_elseif),
	2568:  uint16(anon_sym_else),
	2569:  uint16(anon_sym_for),
	2570:  uint16(anon_sym_function),
	2571:  uint16(anon_sym_local),
	2572:  uint16(anon_sym_global),
	2573:  uint16(anon_sym_LT),
	2574:  uint16(anon_sym_GT),
	2575:  uint16(anon_sym_or),
	2576:  uint16(anon_sym_and),
	2577:  uint16(anon_sym_TILDE),
	2578:  uint16(sym_identifier),
	2579:  uint16(8),
	2580:  uint16(3),
	2581:  uint16(1),
	2582:  uint16(anon_sym_DASH_DASH),
	2583:  uint16(5),
	2584:  uint16(1),
	2585:  uint16(sym__block_comment_start),
	2586:  uint16(185),
	2587:  uint16(1),
	2588:  uint16(anon_sym_SLASH),
	2589:  uint16(189),
	2590:  uint16(1),
	2591:  uint16(anon_sym_CARET),
	2592:  uint16(36),
	2593:  uint16(1),
	2594:  uint16(sym_comment),
	2595:  uint16(175),
	2596:  uint16(3),
	2597:  uint16(anon_sym_STAR),
	2598:  uint16(anon_sym_SLASH_SLASH),
	2599:  uint16(anon_sym_PERCENT),
	2600:  uint16(171),
	2601:  uint16(18),
	2603:  uint16(anon_sym_SEMI),
	2604:  uint16(anon_sym_COMMA),
	2605:  uint16(anon_sym_COLON_COLON),
	2606:  uint16(anon_sym_LPAREN),
	2607:  uint16(anon_sym_RPAREN),
	2608:  uint16(anon_sym_RBRACK),
	2609:  uint16(anon_sym_RBRACE),
	2610:  uint16(anon_sym_LT_EQ),
	2611:  uint16(anon_sym_EQ_EQ),
	2612:  uint16(anon_sym_TILDE_EQ),
	2613:  uint16(anon_sym_GT_EQ),
	2614:  uint16(anon_sym_PIPE),
	2615:  uint16(anon_sym_AMP),
	2616:  uint16(anon_sym_LT_LT),
	2617:  uint16(anon_sym_GT_GT),
	2618:  uint16(anon_sym_PLUS),
	2619:  uint16(anon_sym_DOT_DOT),
	2620:  uint16(173),
	2621:  uint16(23),
	2622:  uint16(anon_sym_return),
	2623:  uint16(sym_break_statement),
	2624:  uint16(anon_sym_goto),
	2625:  uint16(anon_sym_do),
	2626:  uint16(anon_sym_end),
	2627:  uint16(anon_sym_while),
	2628:  uint16(anon_sym_repeat),
	2629:  uint16(anon_sym_until),
	2630:  uint16(anon_sym_if),
	2631:  uint16(anon_sym_then),
	2632:  uint16(anon_sym_elseif),
	2633:  uint16(anon_sym_else),
	2634:  uint16(anon_sym_for),
	2635:  uint16(anon_sym_function),
	2636:  uint16(anon_sym_local),
	2637:  uint16(anon_sym_global),
	2638:  uint16(anon_sym_LT),
	2639:  uint16(anon_sym_GT),
	2640:  uint16(anon_sym_or),
	2641:  uint16(anon_sym_and),
	2642:  uint16(anon_sym_TILDE),
	2643:  uint16(anon_sym_DASH),
	2644:  uint16(sym_identifier),
	2645:  uint16(11),
	2646:  uint16(3),
	2647:  uint16(1),
	2648:  uint16(anon_sym_DASH_DASH),
	2649:  uint16(5),
	2650:  uint16(1),
	2651:  uint16(sym__block_comment_start),
	2652:  uint16(181),
	2653:  uint16(1),
	2654:  uint16(anon_sym_PLUS),
	2655:  uint16(183),
	2656:  uint16(1),
	2657:  uint16(anon_sym_DASH),
	2658:  uint16(185),
	2659:  uint16(1),
	2660:  uint16(anon_sym_SLASH),
	2661:  uint16(187),
	2662:  uint16(1),
	2663:  uint16(anon_sym_DOT_DOT),
	2664:  uint16(189),
	2665:  uint16(1),
	2666:  uint16(anon_sym_CARET),
	2667:  uint16(37),
	2668:  uint16(1),
	2669:  uint16(sym_comment),
	2670:  uint16(175),
	2671:  uint16(3),
	2672:  uint16(anon_sym_STAR),
	2673:  uint16(anon_sym_SLASH_SLASH),
	2674:  uint16(anon_sym_PERCENT),
	2675:  uint16(171),
	2676:  uint16(16),
	2678:  uint16(anon_sym_SEMI),
	2679:  uint16(anon_sym_COMMA),
	2680:  uint16(anon_sym_COLON_COLON),
	2681:  uint16(anon_sym_LPAREN),
	2682:  uint16(anon_sym_RPAREN),
	2683:  uint16(anon_sym_RBRACK),
	2684:  uint16(anon_sym_RBRACE),
	2685:  uint16(anon_sym_LT_EQ),
	2686:  uint16(anon_sym_EQ_EQ),
	2687:  uint16(anon_sym_TILDE_EQ),
	2688:  uint16(anon_sym_GT_EQ),
	2689:  uint16(anon_sym_PIPE),
	2690:  uint16(anon_sym_AMP),
	2691:  uint16(anon_sym_LT_LT),
	2692:  uint16(anon_sym_GT_GT),
	2693:  uint16(173),
	2694:  uint16(22),
	2695:  uint16(anon_sym_return),
	2696:  uint16(sym_break_statement),
	2697:  uint16(anon_sym_goto),
	2698:  uint16(anon_sym_do),
	2699:  uint16(anon_sym_end),
	2700:  uint16(anon_sym_while),
	2701:  uint16(anon_sym_repeat),
	2702:  uint16(anon_sym_until),
	2703:  uint16(anon_sym_if),
	2704:  uint16(anon_sym_then),
	2705:  uint16(anon_sym_elseif),
	2706:  uint16(anon_sym_else),
	2707:  uint16(anon_sym_for),
	2708:  uint16(anon_sym_function),
	2709:  uint16(anon_sym_local),
	2710:  uint16(anon_sym_global),
	2711:  uint16(anon_sym_LT),
	2712:  uint16(anon_sym_GT),
	2713:  uint16(anon_sym_or),
	2714:  uint16(anon_sym_and),
	2715:  uint16(anon_sym_TILDE),
	2716:  uint16(sym_identifier),
	2717:  uint16(6),
	2718:  uint16(3),
	2719:  uint16(1),
	2720:  uint16(anon_sym_DASH_DASH),
	2721:  uint16(5),
	2722:  uint16(1),
	2723:  uint16(sym__block_comment_start),
	2724:  uint16(189),
	2725:  uint16(1),
	2726:  uint16(anon_sym_CARET),
	2727:  uint16(38),
	2728:  uint16(1),
	2729:  uint16(sym_comment),
	2730:  uint16(171),
	2731:  uint16(21),
	2733:  uint16(anon_sym_SEMI),
	2734:  uint16(anon_sym_COMMA),
	2735:  uint16(anon_sym_COLON_COLON),
	2736:  uint16(anon_sym_STAR),
	2737:  uint16(anon_sym_LPAREN),
	2738:  uint16(anon_sym_RPAREN),
	2739:  uint16(anon_sym_RBRACK),
	2740:  uint16(anon_sym_RBRACE),
	2741:  uint16(anon_sym_LT_EQ),
	2742:  uint16(anon_sym_EQ_EQ),
	2743:  uint16(anon_sym_TILDE_EQ),
	2744:  uint16(anon_sym_GT_EQ),
	2745:  uint16(anon_sym_PIPE),
	2746:  uint16(anon_sym_AMP),
	2747:  uint16(anon_sym_LT_LT),
	2748:  uint16(anon_sym_GT_GT),
	2749:  uint16(anon_sym_PLUS),
	2750:  uint16(anon_sym_SLASH_SLASH),
	2751:  uint16(anon_sym_PERCENT),
	2752:  uint16(anon_sym_DOT_DOT),
	2753:  uint16(173),
	2754:  uint16(24),
	2755:  uint16(anon_sym_return),
	2756:  uint16(sym_break_statement),
	2757:  uint16(anon_sym_goto),
	2758:  uint16(anon_sym_do),
	2759:  uint16(anon_sym_end),
	2760:  uint16(anon_sym_while),
	2761:  uint16(anon_sym_repeat),
	2762:  uint16(anon_sym_until),
	2763:  uint16(anon_sym_if),
	2764:  uint16(anon_sym_then),
	2765:  uint16(anon_sym_elseif),
	2766:  uint16(anon_sym_else),
	2767:  uint16(anon_sym_for),
	2768:  uint16(anon_sym_function),
	2769:  uint16(anon_sym_local),
	2770:  uint16(anon_sym_global),
	2771:  uint16(anon_sym_LT),
	2772:  uint16(anon_sym_GT),
	2773:  uint16(anon_sym_or),
	2774:  uint16(anon_sym_and),
	2775:  uint16(anon_sym_TILDE),
	2776:  uint16(anon_sym_DASH),
	2777:  uint16(anon_sym_SLASH),
	2778:  uint16(sym_identifier),
	2779:  uint16(5),
	2780:  uint16(3),
	2781:  uint16(1),
	2782:  uint16(anon_sym_DASH_DASH),
	2783:  uint16(5),
	2784:  uint16(1),
	2785:  uint16(sym__block_comment_start),
	2786:  uint16(39),
	2787:  uint16(1),
	2788:  uint16(sym_comment),
	2789:  uint16(101),
	2790:  uint16(22),
	2792:  uint16(anon_sym_SEMI),
	2793:  uint16(anon_sym_COMMA),
	2794:  uint16(anon_sym_COLON_COLON),
	2795:  uint16(anon_sym_STAR),
	2796:  uint16(anon_sym_LPAREN),
	2797:  uint16(anon_sym_RPAREN),
	2798:  uint16(anon_sym_RBRACK),
	2799:  uint16(anon_sym_RBRACE),
	2800:  uint16(anon_sym_LT_EQ),
	2801:  uint16(anon_sym_EQ_EQ),
	2802:  uint16(anon_sym_TILDE_EQ),
	2803:  uint16(anon_sym_GT_EQ),
	2804:  uint16(anon_sym_PIPE),
	2805:  uint16(anon_sym_AMP),
	2806:  uint16(anon_sym_LT_LT),
	2807:  uint16(anon_sym_GT_GT),
	2808:  uint16(anon_sym_PLUS),
	2809:  uint16(anon_sym_SLASH_SLASH),
	2810:  uint16(anon_sym_PERCENT),
	2811:  uint16(anon_sym_DOT_DOT),
	2812:  uint16(anon_sym_CARET),
	2813:  uint16(103),
	2814:  uint16(24),
	2815:  uint16(anon_sym_return),
	2816:  uint16(sym_break_statement),
	2817:  uint16(anon_sym_goto),
	2818:  uint16(anon_sym_do),
	2819:  uint16(anon_sym_end),
	2820:  uint16(anon_sym_while),
	2821:  uint16(anon_sym_repeat),
	2822:  uint16(anon_sym_until),
	2823:  uint16(anon_sym_if),
	2824:  uint16(anon_sym_then),
	2825:  uint16(anon_sym_elseif),
	2826:  uint16(anon_sym_else),
	2827:  uint16(anon_sym_for),
	2828:  uint16(anon_sym_function),
	2829:  uint16(anon_sym_local),
	2830:  uint16(anon_sym_global),
	2831:  uint16(anon_sym_LT),
	2832:  uint16(anon_sym_GT),
	2833:  uint16(anon_sym_or),
	2834:  uint16(anon_sym_and),
	2835:  uint16(anon_sym_TILDE),
	2836:  uint16(anon_sym_DASH),
	2837:  uint16(anon_sym_SLASH),
	2838:  uint16(sym_identifier),
	2839:  uint16(5),
	2840:  uint16(3),
	2841:  uint16(1),
	2842:  uint16(anon_sym_DASH_DASH),
	2843:  uint16(5),
	2844:  uint16(1),
	2845:  uint16(sym__block_comment_start),
	2846:  uint16(40),
	2847:  uint16(1),
	2848:  uint16(sym_comment),
	2849:  uint16(61),
	2850:  uint16(22),
	2852:  uint16(anon_sym_SEMI),
	2853:  uint16(anon_sym_COMMA),
	2854:  uint16(anon_sym_COLON_COLON),
	2855:  uint16(anon_sym_STAR),
	2856:  uint16(anon_sym_LPAREN),
	2857:  uint16(anon_sym_RPAREN),
	2858:  uint16(anon_sym_RBRACK),
	2859:  uint16(anon_sym_RBRACE),
	2860:  uint16(anon_sym_LT_EQ),
	2861:  uint16(anon_sym_EQ_EQ),
	2862:  uint16(anon_sym_TILDE_EQ),
	2863:  uint16(anon_sym_GT_EQ),
	2864:  uint16(anon_sym_PIPE),
	2865:  uint16(anon_sym_AMP),
	2866:  uint16(anon_sym_LT_LT),
	2867:  uint16(anon_sym_GT_GT),
	2868:  uint16(anon_sym_PLUS),
	2869:  uint16(anon_sym_SLASH_SLASH),
	2870:  uint16(anon_sym_PERCENT),
	2871:  uint16(anon_sym_DOT_DOT),
	2872:  uint16(anon_sym_CARET),
	2873:  uint16(63),
	2874:  uint16(24),
	2875:  uint16(anon_sym_return),
	2876:  uint16(sym_break_statement),
	2877:  uint16(anon_sym_goto),
	2878:  uint16(anon_sym_do),
	2879:  uint16(anon_sym_end),
	2880:  uint16(anon_sym_while),
	2881:  uint16(anon_sym_repeat),
	2882:  uint16(anon_sym_until),
	2883:  uint16(anon_sym_if),
	2884:  uint16(anon_sym_then),
	2885:  uint16(anon_sym_elseif),
	2886:  uint16(anon_sym_else),
	2887:  uint16(anon_sym_for),
	2888:  uint16(anon_sym_function),
	2889:  uint16(anon_sym_local),
	2890:  uint16(anon_sym_global),
	2891:  uint16(anon_sym_LT),
	2892:  uint16(anon_sym_GT),
	2893:  uint16(anon_sym_or),
	2894:  uint16(anon_sym_and),
	2895:  uint16(anon_sym_TILDE),
	2896:  uint16(anon_sym_DASH),
	2897:  uint16(anon_sym_SLASH),
	2898:  uint16(sym_identifier),
	2899:  uint16(5),
	2900:  uint16(3),
	2901:  uint16(1),
	2902:  uint16(anon_sym_DASH_DASH),
	2903:  uint16(5),
	2904:  uint16(1),
	2905:  uint16(sym__block_comment_start),
	2906:  uint16(41),
	2907:  uint16(1),
	2908:  uint16(sym_comment),
	2909:  uint16(93),
	2910:  uint16(22),
	2912:  uint16(anon_sym_SEMI),
	2913:  uint16(anon_sym_COMMA),
	2914:  uint16(anon_sym_COLON_COLON),
	2915:  uint16(anon_sym_STAR),
	2916:  uint16(anon_sym_LPAREN),
	2917:  uint16(anon_sym_RPAREN),
	2918:  uint16(anon_sym_RBRACK),
	2919:  uint16(anon_sym_RBRACE),
	2920:  uint16(anon_sym_LT_EQ),
	2921:  uint16(anon_sym_EQ_EQ),
	2922:  uint16(anon_sym_TILDE_EQ),
	2923:  uint16(anon_sym_GT_EQ),
	2924:  uint16(anon_sym_PIPE),
	2925:  uint16(anon_sym_AMP),
	2926:  uint16(anon_sym_LT_LT),
	2927:  uint16(anon_sym_GT_GT),
	2928:  uint16(anon_sym_PLUS),
	2929:  uint16(anon_sym_SLASH_SLASH),
	2930:  uint16(anon_sym_PERCENT),
	2931:  uint16(anon_sym_DOT_DOT),
	2932:  uint16(anon_sym_CARET),
	2933:  uint16(95),
	2934:  uint16(24),
	2935:  uint16(anon_sym_return),
	2936:  uint16(sym_break_statement),
	2937:  uint16(anon_sym_goto),
	2938:  uint16(anon_sym_do),
	2939:  uint16(anon_sym_end),
	2940:  uint16(anon_sym_while),
	2941:  uint16(anon_sym_repeat),
	2942:  uint16(anon_sym_until),
	2943:  uint16(anon_sym_if),
	2944:  uint16(anon_sym_then),
	2945:  uint16(anon_sym_elseif),
	2946:  uint16(anon_sym_else),
	2947:  uint16(anon_sym_for),
	2948:  uint16(anon_sym_function),
	2949:  uint16(anon_sym_local),
	2950:  uint16(anon_sym_global),
	2951:  uint16(anon_sym_LT),
	2952:  uint16(anon_sym_GT),
	2953:  uint16(anon_sym_or),
	2954:  uint16(anon_sym_and),
	2955:  uint16(anon_sym_TILDE),
	2956:  uint16(anon_sym_DASH),
	2957:  uint16(anon_sym_SLASH),
	2958:  uint16(sym_identifier),
	2959:  uint16(5),
	2960:  uint16(3),
	2961:  uint16(1),
	2962:  uint16(anon_sym_DASH_DASH),
	2963:  uint16(5),
	2964:  uint16(1),
	2965:  uint16(sym__block_comment_start),
	2966:  uint16(42),
	2967:  uint16(1),
	2968:  uint16(sym_comment),
	2969:  uint16(205),
	2970:  uint16(22),
	2972:  uint16(anon_sym_SEMI),
	2973:  uint16(anon_sym_COMMA),
	2974:  uint16(anon_sym_COLON_COLON),
	2975:  uint16(anon_sym_STAR),
	2976:  uint16(anon_sym_LPAREN),
	2977:  uint16(anon_sym_RPAREN),
	2978:  uint16(anon_sym_RBRACK),
	2979:  uint16(anon_sym_RBRACE),
	2980:  uint16(anon_sym_LT_EQ),
	2981:  uint16(anon_sym_EQ_EQ),
	2982:  uint16(anon_sym_TILDE_EQ),
	2983:  uint16(anon_sym_GT_EQ),
	2984:  uint16(anon_sym_PIPE),
	2985:  uint16(anon_sym_AMP),
	2986:  uint16(anon_sym_LT_LT),
	2987:  uint16(anon_sym_GT_GT),
	2988:  uint16(anon_sym_PLUS),
	2989:  uint16(anon_sym_SLASH_SLASH),
	2990:  uint16(anon_sym_PERCENT),
	2991:  uint16(anon_sym_DOT_DOT),
	2992:  uint16(anon_sym_CARET),
	2993:  uint16(207),
	2994:  uint16(24),
	2995:  uint16(anon_sym_return),
	2996:  uint16(sym_break_statement),
	2997:  uint16(anon_sym_goto),
	2998:  uint16(anon_sym_do),
	2999:  uint16(anon_sym_end),
	3000:  uint16(anon_sym_while),
	3001:  uint16(anon_sym_repeat),
	3002:  uint16(anon_sym_until),
	3003:  uint16(anon_sym_if),
	3004:  uint16(anon_sym_then),
	3005:  uint16(anon_sym_elseif),
	3006:  uint16(anon_sym_else),
	3007:  uint16(anon_sym_for),
	3008:  uint16(anon_sym_function),
	3009:  uint16(anon_sym_local),
	3010:  uint16(anon_sym_global),
	3011:  uint16(anon_sym_LT),
	3012:  uint16(anon_sym_GT),
	3013:  uint16(anon_sym_or),
	3014:  uint16(anon_sym_and),
	3015:  uint16(anon_sym_TILDE),
	3016:  uint16(anon_sym_DASH),
	3017:  uint16(anon_sym_SLASH),
	3018:  uint16(sym_identifier),
	3019:  uint16(5),
	3020:  uint16(3),
	3021:  uint16(1),
	3022:  uint16(anon_sym_DASH_DASH),
	3023:  uint16(5),
	3024:  uint16(1),
	3025:  uint16(sym__block_comment_start),
	3026:  uint16(43),
	3027:  uint16(1),
	3028:  uint16(sym_comment),
	3029:  uint16(209),
	3030:  uint16(22),
	3032:  uint16(anon_sym_SEMI),
	3033:  uint16(anon_sym_COMMA),
	3034:  uint16(anon_sym_COLON_COLON),
	3035:  uint16(anon_sym_STAR),
	3036:  uint16(anon_sym_LPAREN),
	3037:  uint16(anon_sym_RPAREN),
	3038:  uint16(anon_sym_RBRACK),
	3039:  uint16(anon_sym_RBRACE),
	3040:  uint16(anon_sym_LT_EQ),
	3041:  uint16(anon_sym_EQ_EQ),
	3042:  uint16(anon_sym_TILDE_EQ),
	3043:  uint16(anon_sym_GT_EQ),
	3044:  uint16(anon_sym_PIPE),
	3045:  uint16(anon_sym_AMP),
	3046:  uint16(anon_sym_LT_LT),
	3047:  uint16(anon_sym_GT_GT),
	3048:  uint16(anon_sym_PLUS),
	3049:  uint16(anon_sym_SLASH_SLASH),
	3050:  uint16(anon_sym_PERCENT),
	3051:  uint16(anon_sym_DOT_DOT),
	3052:  uint16(anon_sym_CARET),
	3053:  uint16(211),
	3054:  uint16(24),
	3055:  uint16(anon_sym_return),
	3056:  uint16(sym_break_statement),
	3057:  uint16(anon_sym_goto),
	3058:  uint16(anon_sym_do),
	3059:  uint16(anon_sym_end),
	3060:  uint16(anon_sym_while),
	3061:  uint16(anon_sym_repeat),
	3062:  uint16(anon_sym_until),
	3063:  uint16(anon_sym_if),
	3064:  uint16(anon_sym_then),
	3065:  uint16(anon_sym_elseif),
	3066:  uint16(anon_sym_else),
	3067:  uint16(anon_sym_for),
	3068:  uint16(anon_sym_function),
	3069:  uint16(anon_sym_local),
	3070:  uint16(anon_sym_global),
	3071:  uint16(anon_sym_LT),
	3072:  uint16(anon_sym_GT),
	3073:  uint16(anon_sym_or),
	3074:  uint16(anon_sym_and),
	3075:  uint16(anon_sym_TILDE),
	3076:  uint16(anon_sym_DASH),
	3077:  uint16(anon_sym_SLASH),
	3078:  uint16(sym_identifier),
	3079:  uint16(5),
	3080:  uint16(3),
	3081:  uint16(1),
	3082:  uint16(anon_sym_DASH_DASH),
	3083:  uint16(5),
	3084:  uint16(1),
	3085:  uint16(sym__block_comment_start),
	3086:  uint16(44),
	3087:  uint16(1),
	3088:  uint16(sym_comment),
	3089:  uint16(97),
	3090:  uint16(22),
	3092:  uint16(anon_sym_SEMI),
	3093:  uint16(anon_sym_COMMA),
	3094:  uint16(anon_sym_COLON_COLON),
	3095:  uint16(anon_sym_STAR),
	3096:  uint16(anon_sym_LPAREN),
	3097:  uint16(anon_sym_RPAREN),
	3098:  uint16(anon_sym_RBRACK),
	3099:  uint16(anon_sym_RBRACE),
	3100:  uint16(anon_sym_LT_EQ),
	3101:  uint16(anon_sym_EQ_EQ),
	3102:  uint16(anon_sym_TILDE_EQ),
	3103:  uint16(anon_sym_GT_EQ),
	3104:  uint16(anon_sym_PIPE),
	3105:  uint16(anon_sym_AMP),
	3106:  uint16(anon_sym_LT_LT),
	3107:  uint16(anon_sym_GT_GT),
	3108:  uint16(anon_sym_PLUS),
	3109:  uint16(anon_sym_SLASH_SLASH),
	3110:  uint16(anon_sym_PERCENT),
	3111:  uint16(anon_sym_DOT_DOT),
	3112:  uint16(anon_sym_CARET),
	3113:  uint16(99),
	3114:  uint16(24),
	3115:  uint16(anon_sym_return),
	3116:  uint16(sym_break_statement),
	3117:  uint16(anon_sym_goto),
	3118:  uint16(anon_sym_do),
	3119:  uint16(anon_sym_end),
	3120:  uint16(anon_sym_while),
	3121:  uint16(anon_sym_repeat),
	3122:  uint16(anon_sym_until),
	3123:  uint16(anon_sym_if),
	3124:  uint16(anon_sym_then),
	3125:  uint16(anon_sym_elseif),
	3126:  uint16(anon_sym_else),
	3127:  uint16(anon_sym_for),
	3128:  uint16(anon_sym_function),
	3129:  uint16(anon_sym_local),
	3130:  uint16(anon_sym_global),
	3131:  uint16(anon_sym_LT),
	3132:  uint16(anon_sym_GT),
	3133:  uint16(anon_sym_or),
	3134:  uint16(anon_sym_and),
	3135:  uint16(anon_sym_TILDE),
	3136:  uint16(anon_sym_DASH),
	3137:  uint16(anon_sym_SLASH),
	3138:  uint16(sym_identifier),
	3139:  uint16(5),
	3140:  uint16(3),
	3141:  uint16(1),
	3142:  uint16(anon_sym_DASH_DASH),
	3143:  uint16(5),
	3144:  uint16(1),
	3145:  uint16(sym__block_comment_start),
	3146:  uint16(45),
	3147:  uint16(1),
	3148:  uint16(sym_comment),
	3149:  uint16(93),
	3150:  uint16(22),
	3152:  uint16(anon_sym_SEMI),
	3153:  uint16(anon_sym_COMMA),
	3154:  uint16(anon_sym_COLON_COLON),
	3155:  uint16(anon_sym_STAR),
	3156:  uint16(anon_sym_LPAREN),
	3157:  uint16(anon_sym_RPAREN),
	3158:  uint16(anon_sym_RBRACK),
	3159:  uint16(anon_sym_RBRACE),
	3160:  uint16(anon_sym_LT_EQ),
	3161:  uint16(anon_sym_EQ_EQ),
	3162:  uint16(anon_sym_TILDE_EQ),
	3163:  uint16(anon_sym_GT_EQ),
	3164:  uint16(anon_sym_PIPE),
	3165:  uint16(anon_sym_AMP),
	3166:  uint16(anon_sym_LT_LT),
	3167:  uint16(anon_sym_GT_GT),
	3168:  uint16(anon_sym_PLUS),
	3169:  uint16(anon_sym_SLASH_SLASH),
	3170:  uint16(anon_sym_PERCENT),
	3171:  uint16(anon_sym_DOT_DOT),
	3172:  uint16(anon_sym_CARET),
	3173:  uint16(95),
	3174:  uint16(24),
	3175:  uint16(anon_sym_return),
	3176:  uint16(sym_break_statement),
	3177:  uint16(anon_sym_goto),
	3178:  uint16(anon_sym_do),
	3179:  uint16(anon_sym_end),
	3180:  uint16(anon_sym_while),
	3181:  uint16(anon_sym_repeat),
	3182:  uint16(anon_sym_until),
	3183:  uint16(anon_sym_if),
	3184:  uint16(anon_sym_then),
	3185:  uint16(anon_sym_elseif),
	3186:  uint16(anon_sym_else),
	3187:  uint16(anon_sym_for),
	3188:  uint16(anon_sym_function),
	3189:  uint16(anon_sym_local),
	3190:  uint16(anon_sym_global),
	3191:  uint16(anon_sym_LT),
	3192:  uint16(anon_sym_GT),
	3193:  uint16(anon_sym_or),
	3194:  uint16(anon_sym_and),
	3195:  uint16(anon_sym_TILDE),
	3196:  uint16(anon_sym_DASH),
	3197:  uint16(anon_sym_SLASH),
	3198:  uint16(sym_identifier),
	3199:  uint16(6),
	3200:  uint16(3),
	3201:  uint16(1),
	3202:  uint16(anon_sym_DASH_DASH),
	3203:  uint16(5),
	3204:  uint16(1),
	3205:  uint16(sym__block_comment_start),
	3206:  uint16(189),
	3207:  uint16(1),
	3208:  uint16(anon_sym_CARET),
	3209:  uint16(46),
	3210:  uint16(1),
	3211:  uint16(sym_comment),
	3212:  uint16(213),
	3213:  uint16(21),
	3215:  uint16(anon_sym_SEMI),
	3216:  uint16(anon_sym_COMMA),
	3217:  uint16(anon_sym_COLON_COLON),
	3218:  uint16(anon_sym_STAR),
	3219:  uint16(anon_sym_LPAREN),
	3220:  uint16(anon_sym_RPAREN),
	3221:  uint16(anon_sym_RBRACK),
	3222:  uint16(anon_sym_RBRACE),
	3223:  uint16(anon_sym_LT_EQ),
	3224:  uint16(anon_sym_EQ_EQ),
	3225:  uint16(anon_sym_TILDE_EQ),
	3226:  uint16(anon_sym_GT_EQ),
	3227:  uint16(anon_sym_PIPE),
	3228:  uint16(anon_sym_AMP),
	3229:  uint16(anon_sym_LT_LT),
	3230:  uint16(anon_sym_GT_GT),
	3231:  uint16(anon_sym_PLUS),
	3232:  uint16(anon_sym_SLASH_SLASH),
	3233:  uint16(anon_sym_PERCENT),
	3234:  uint16(anon_sym_DOT_DOT),
	3235:  uint16(215),
	3236:  uint16(24),
	3237:  uint16(anon_sym_return),
	3238:  uint16(sym_break_statement),
	3239:  uint16(anon_sym_goto),
	3240:  uint16(anon_sym_do),
	3241:  uint16(anon_sym_end),
	3242:  uint16(anon_sym_while),
	3243:  uint16(anon_sym_repeat),
	3244:  uint16(anon_sym_until),
	3245:  uint16(anon_sym_if),
	3246:  uint16(anon_sym_then),
	3247:  uint16(anon_sym_elseif),
	3248:  uint16(anon_sym_else),
	3249:  uint16(anon_sym_for),
	3250:  uint16(anon_sym_function),
	3251:  uint16(anon_sym_local),
	3252:  uint16(anon_sym_global),
	3253:  uint16(anon_sym_LT),
	3254:  uint16(anon_sym_GT),
	3255:  uint16(anon_sym_or),
	3256:  uint16(anon_sym_and),
	3257:  uint16(anon_sym_TILDE),
	3258:  uint16(anon_sym_DASH),
	3259:  uint16(anon_sym_SLASH),
	3260:  uint16(sym_identifier),
	3261:  uint16(38),
	3262:  uint16(3),
	3263:  uint16(1),
	3264:  uint16(anon_sym_DASH_DASH),
	3265:  uint16(5),
	3266:  uint16(1),
	3267:  uint16(sym__block_comment_start),
	3268:  uint16(9),
	3269:  uint16(1),
	3270:  uint16(sym_identifier),
	3271:  uint16(15),
	3272:  uint16(1),
	3273:  uint16(anon_sym_SEMI),
	3274:  uint16(17),
	3275:  uint16(1),
	3276:  uint16(anon_sym_COLON_COLON),
	3277:  uint16(19),
	3278:  uint16(1),
	3279:  uint16(sym_break_statement),
	3280:  uint16(21),
	3281:  uint16(1),
	3282:  uint16(anon_sym_goto),
	3283:  uint16(23),
	3284:  uint16(1),
	3285:  uint16(anon_sym_do),
	3286:  uint16(25),
	3287:  uint16(1),
	3288:  uint16(anon_sym_while),
	3289:  uint16(27),
	3290:  uint16(1),
	3291:  uint16(anon_sym_repeat),
	3292:  uint16(29),
	3293:  uint16(1),
	3294:  uint16(anon_sym_if),
	3295:  uint16(31),
	3296:  uint16(1),
	3297:  uint16(anon_sym_for),
	3298:  uint16(33),
	3299:  uint16(1),
	3300:  uint16(anon_sym_function),
	3301:  uint16(35),
	3302:  uint16(1),
	3303:  uint16(anon_sym_local),
	3304:  uint16(37),
	3305:  uint16(1),
	3306:  uint16(anon_sym_global),
	3307:  uint16(39),
	3308:  uint16(1),
	3309:  uint16(anon_sym_LPAREN),
	3310:  uint16(113),
	3311:  uint16(1),
	3312:  uint16(anon_sym_return),
	3313:  uint16(217),
	3314:  uint16(1),
	3315:  uint16(anon_sym_end),
	3316:  uint16(21),
	3317:  uint16(1),
	3318:  uint16(aux_sym_chunk_repeat1),
	3319:  uint16(47),
	3320:  uint16(1),
	3321:  uint16(sym_comment),
	3322:  uint16(101),
	3323:  uint16(1),
	3324:  uint16(sym_function_call),
	3325:  uint16(131),
	3326:  uint16(1),
	3327:  uint16(sym_variable_declaration),
	3328:  uint16(132),
	3329:  uint16(1),
	3330:  uint16(sym_declaration),
	3331:  uint16(134),
	3332:  uint16(1),
	3333:  uint16(sym_function_declaration),
	3334:  uint16(137),
	3335:  uint16(1),
	3336:  uint16(sym__local_function_declaration),
	3337:  uint16(139),
	3338:  uint16(1),
	3339:  uint16(sym__global_function_declaration),
	3340:  uint16(142),
	3341:  uint16(1),
	3342:  uint16(sym__global_variable_declaration),
	3343:  uint16(143),
	3344:  uint16(1),
	3345:  uint16(sym__global_implicit_variable_declaration),
	3346:  uint16(161),
	3347:  uint16(1),
	3348:  uint16(sym_statement),
	3349:  uint16(170),
	3350:  uint16(1),
	3351:  uint16(sym__prefix_expression),
	3352:  uint16(172),
	3353:  uint16(1),
	3354:  uint16(sym_variable),
	3355:  uint16(174),
	3356:  uint16(1),
	3357:  uint16(sym_method_index_expression),
	3358:  uint16(177),
	3359:  uint16(1),
	3360:  uint16(sym_parenthesized_expression),
	3361:  uint16(194),
	3362:  uint16(1),
	3363:  uint16(sym_return_statement),
	3364:  uint16(255),
	3365:  uint16(1),
	3366:  uint16(sym__variable_assignment_varlist),
	3367:  uint16(269),
	3368:  uint16(1),
	3369:  uint16(sym__block),
	3370:  uint16(2),
	3371:  uint16(3),
	3372:  uint16(sym_bracket_index_expression),
	3373:  uint16(sym_dot_index_expression),
	3374:  uint16(sym__contextual_keyword),
	3375:  uint16(160),
	3376:  uint16(9),
	3377:  uint16(sym_empty_statement),
	3378:  uint16(sym_assignment_statement),
	3379:  uint16(sym_label_statement),
	3380:  uint16(sym_goto_statement),
	3381:  uint16(sym_do_statement),
	3382:  uint16(sym_while_statement),
	3383:  uint16(sym_repeat_statement),
	3384:  uint16(sym_if_statement),
	3385:  uint16(sym_for_statement),
	3386:  uint16(38),
	3387:  uint16(3),
	3388:  uint16(1),
	3389:  uint16(anon_sym_DASH_DASH),
	3390:  uint16(5),
	3391:  uint16(1),
	3392:  uint16(sym__block_comment_start),
	3393:  uint16(9),
	3394:  uint16(1),
	3395:  uint16(sym_identifier),
	3396:  uint16(15),
	3397:  uint16(1),
	3398:  uint16(anon_sym_SEMI),
	3399:  uint16(17),
	3400:  uint16(1),
	3401:  uint16(anon_sym_COLON_COLON),
	3402:  uint16(19),
	3403:  uint16(1),
	3404:  uint16(sym_break_statement),
	3405:  uint16(21),
	3406:  uint16(1),
	3407:  uint16(anon_sym_goto),
	3408:  uint16(23),
	3409:  uint16(1),
	3410:  uint16(anon_sym_do),
	3411:  uint16(25),
	3412:  uint16(1),
	3413:  uint16(anon_sym_while),
	3414:  uint16(27),
	3415:  uint16(1),
	3416:  uint16(anon_sym_repeat),
	3417:  uint16(29),
	3418:  uint16(1),
	3419:  uint16(anon_sym_if),
	3420:  uint16(31),
	3421:  uint16(1),
	3422:  uint16(anon_sym_for),
	3423:  uint16(33),
	3424:  uint16(1),
	3425:  uint16(anon_sym_function),
	3426:  uint16(35),
	3427:  uint16(1),
	3428:  uint16(anon_sym_local),
	3429:  uint16(37),
	3430:  uint16(1),
	3431:  uint16(anon_sym_global),
	3432:  uint16(39),
	3433:  uint16(1),
	3434:  uint16(anon_sym_LPAREN),
	3435:  uint16(113),
	3436:  uint16(1),
	3437:  uint16(anon_sym_return),
	3438:  uint16(219),
	3439:  uint16(1),
	3440:  uint16(anon_sym_until),
	3441:  uint16(21),
	3442:  uint16(1),
	3443:  uint16(aux_sym_chunk_repeat1),
	3444:  uint16(48),
	3445:  uint16(1),
	3446:  uint16(sym_comment),
	3447:  uint16(101),
	3448:  uint16(1),
	3449:  uint16(sym_function_call),
	3450:  uint16(131),
	3451:  uint16(1),
	3452:  uint16(sym_variable_declaration),
	3453:  uint16(132),
	3454:  uint16(1),
	3455:  uint16(sym_declaration),
	3456:  uint16(134),
	3457:  uint16(1),
	3458:  uint16(sym_function_declaration),
	3459:  uint16(137),
	3460:  uint16(1),
	3461:  uint16(sym__local_function_declaration),
	3462:  uint16(139),
	3463:  uint16(1),
	3464:  uint16(sym__global_function_declaration),
	3465:  uint16(142),
	3466:  uint16(1),
	3467:  uint16(sym__global_variable_declaration),
	3468:  uint16(143),
	3469:  uint16(1),
	3470:  uint16(sym__global_implicit_variable_declaration),
	3471:  uint16(161),
	3472:  uint16(1),
	3473:  uint16(sym_statement),
	3474:  uint16(170),
	3475:  uint16(1),
	3476:  uint16(sym__prefix_expression),
	3477:  uint16(172),
	3478:  uint16(1),
	3479:  uint16(sym_variable),
	3480:  uint16(174),
	3481:  uint16(1),
	3482:  uint16(sym_method_index_expression),
	3483:  uint16(177),
	3484:  uint16(1),
	3485:  uint16(sym_parenthesized_expression),
	3486:  uint16(194),
	3487:  uint16(1),
	3488:  uint16(sym_return_statement),
	3489:  uint16(241),
	3490:  uint16(1),
	3491:  uint16(sym__block),
	3492:  uint16(255),
	3493:  uint16(1),
	3494:  uint16(sym__variable_assignment_varlist),
	3495:  uint16(2),
	3496:  uint16(3),
	3497:  uint16(sym_bracket_index_expression),
	3498:  uint16(sym_dot_index_expression),
	3499:  uint16(sym__contextual_keyword),
	3500:  uint16(160),
	3501:  uint16(9),
	3502:  uint16(sym_empty_statement),
	3503:  uint16(sym_assignment_statement),
	3504:  uint16(sym_label_statement),
	3505:  uint16(sym_goto_statement),
	3506:  uint16(sym_do_statement),
	3507:  uint16(sym_while_statement),
	3508:  uint16(sym_repeat_statement),
	3509:  uint16(sym_if_statement),
	3510:  uint16(sym_for_statement),
	3511:  uint16(38),
	3512:  uint16(3),
	3513:  uint16(1),
	3514:  uint16(anon_sym_DASH_DASH),
	3515:  uint16(5),
	3516:  uint16(1),
	3517:  uint16(sym__block_comment_start),
	3518:  uint16(9),
	3519:  uint16(1),
	3520:  uint16(sym_identifier),
	3521:  uint16(15),
	3522:  uint16(1),
	3523:  uint16(anon_sym_SEMI),
	3524:  uint16(17),
	3525:  uint16(1),
	3526:  uint16(anon_sym_COLON_COLON),
	3527:  uint16(19),
	3528:  uint16(1),
	3529:  uint16(sym_break_statement),
	3530:  uint16(21),
	3531:  uint16(1),
	3532:  uint16(anon_sym_goto),
	3533:  uint16(23),
	3534:  uint16(1),
	3535:  uint16(anon_sym_do),
	3536:  uint16(25),
	3537:  uint16(1),
	3538:  uint16(anon_sym_while),
	3539:  uint16(27),
	3540:  uint16(1),
	3541:  uint16(anon_sym_repeat),
	3542:  uint16(29),
	3543:  uint16(1),
	3544:  uint16(anon_sym_if),
	3545:  uint16(31),
	3546:  uint16(1),
	3547:  uint16(anon_sym_for),
	3548:  uint16(33),
	3549:  uint16(1),
	3550:  uint16(anon_sym_function),
	3551:  uint16(35),
	3552:  uint16(1),
	3553:  uint16(anon_sym_local),
	3554:  uint16(37),
	3555:  uint16(1),
	3556:  uint16(anon_sym_global),
	3557:  uint16(39),
	3558:  uint16(1),
	3559:  uint16(anon_sym_LPAREN),
	3560:  uint16(113),
	3561:  uint16(1),
	3562:  uint16(anon_sym_return),
	3563:  uint16(221),
	3564:  uint16(1),
	3565:  uint16(anon_sym_end),
	3566:  uint16(21),
	3567:  uint16(1),
	3568:  uint16(aux_sym_chunk_repeat1),
	3569:  uint16(49),
	3570:  uint16(1),
	3571:  uint16(sym_comment),
	3572:  uint16(101),
	3573:  uint16(1),
	3574:  uint16(sym_function_call),
	3575:  uint16(131),
	3576:  uint16(1),
	3577:  uint16(sym_variable_declaration),
	3578:  uint16(132),
	3579:  uint16(1),
	3580:  uint16(sym_declaration),
	3581:  uint16(134),
	3582:  uint16(1),
	3583:  uint16(sym_function_declaration),
	3584:  uint16(137),
	3585:  uint16(1),
	3586:  uint16(sym__local_function_declaration),
	3587:  uint16(139),
	3588:  uint16(1),
	3589:  uint16(sym__global_function_declaration),
	3590:  uint16(142),
	3591:  uint16(1),
	3592:  uint16(sym__global_variable_declaration),
	3593:  uint16(143),
	3594:  uint16(1),
	3595:  uint16(sym__global_implicit_variable_declaration),
	3596:  uint16(161),
	3597:  uint16(1),
	3598:  uint16(sym_statement),
	3599:  uint16(170),
	3600:  uint16(1),
	3601:  uint16(sym__prefix_expression),
	3602:  uint16(172),
	3603:  uint16(1),
	3604:  uint16(sym_variable),
	3605:  uint16(174),
	3606:  uint16(1),
	3607:  uint16(sym_method_index_expression),
	3608:  uint16(177),
	3609:  uint16(1),
	3610:  uint16(sym_parenthesized_expression),
	3611:  uint16(194),
	3612:  uint16(1),
	3613:  uint16(sym_return_statement),
	3614:  uint16(234),
	3615:  uint16(1),
	3616:  uint16(sym__block),
	3617:  uint16(255),
	3618:  uint16(1),
	3619:  uint16(sym__variable_assignment_varlist),
	3620:  uint16(2),
	3621:  uint16(3),
	3622:  uint16(sym_bracket_index_expression),
	3623:  uint16(sym_dot_index_expression),
	3624:  uint16(sym__contextual_keyword),
	3625:  uint16(160),
	3626:  uint16(9),
	3627:  uint16(sym_empty_statement),
	3628:  uint16(sym_assignment_statement),
	3629:  uint16(sym_label_statement),
	3630:  uint16(sym_goto_statement),
	3631:  uint16(sym_do_statement),
	3632:  uint16(sym_while_statement),
	3633:  uint16(sym_repeat_statement),
	3634:  uint16(sym_if_statement),
	3635:  uint16(sym_for_statement),
	3636:  uint16(38),
	3637:  uint16(3),
	3638:  uint16(1),
	3639:  uint16(anon_sym_DASH_DASH),
	3640:  uint16(5),
	3641:  uint16(1),
	3642:  uint16(sym__block_comment_start),
	3643:  uint16(9),
	3644:  uint16(1),
	3645:  uint16(sym_identifier),
	3646:  uint16(15),
	3647:  uint16(1),
	3648:  uint16(anon_sym_SEMI),
	3649:  uint16(17),
	3650:  uint16(1),
	3651:  uint16(anon_sym_COLON_COLON),
	3652:  uint16(19),
	3653:  uint16(1),
	3654:  uint16(sym_break_statement),
	3655:  uint16(21),
	3656:  uint16(1),
	3657:  uint16(anon_sym_goto),
	3658:  uint16(23),
	3659:  uint16(1),
	3660:  uint16(anon_sym_do),
	3661:  uint16(25),
	3662:  uint16(1),
	3663:  uint16(anon_sym_while),
	3664:  uint16(27),
	3665:  uint16(1),
	3666:  uint16(anon_sym_repeat),
	3667:  uint16(29),
	3668:  uint16(1),
	3669:  uint16(anon_sym_if),
	3670:  uint16(31),
	3671:  uint16(1),
	3672:  uint16(anon_sym_for),
	3673:  uint16(33),
	3674:  uint16(1),
	3675:  uint16(anon_sym_function),
	3676:  uint16(35),
	3677:  uint16(1),
	3678:  uint16(anon_sym_local),
	3679:  uint16(37),
	3680:  uint16(1),
	3681:  uint16(anon_sym_global),
	3682:  uint16(39),
	3683:  uint16(1),
	3684:  uint16(anon_sym_LPAREN),
	3685:  uint16(113),
	3686:  uint16(1),
	3687:  uint16(anon_sym_return),
	3688:  uint16(223),
	3689:  uint16(1),
	3690:  uint16(anon_sym_end),
	3691:  uint16(21),
	3692:  uint16(1),
	3693:  uint16(aux_sym_chunk_repeat1),
	3694:  uint16(50),
	3695:  uint16(1),
	3696:  uint16(sym_comment),
	3697:  uint16(101),
	3698:  uint16(1),
	3699:  uint16(sym_function_call),
	3700:  uint16(131),
	3701:  uint16(1),
	3702:  uint16(sym_variable_declaration),
	3703:  uint16(132),
	3704:  uint16(1),
	3705:  uint16(sym_declaration),
	3706:  uint16(134),
	3707:  uint16(1),
	3708:  uint16(sym_function_declaration),
	3709:  uint16(137),
	3710:  uint16(1),
	3711:  uint16(sym__local_function_declaration),
	3712:  uint16(139),
	3713:  uint16(1),
	3714:  uint16(sym__global_function_declaration),
	3715:  uint16(142),
	3716:  uint16(1),
	3717:  uint16(sym__global_variable_declaration),
	3718:  uint16(143),
	3719:  uint16(1),
	3720:  uint16(sym__global_implicit_variable_declaration),
	3721:  uint16(161),
	3722:  uint16(1),
	3723:  uint16(sym_statement),
	3724:  uint16(170),
	3725:  uint16(1),
	3726:  uint16(sym__prefix_expression),
	3727:  uint16(172),
	3728:  uint16(1),
	3729:  uint16(sym_variable),
	3730:  uint16(174),
	3731:  uint16(1),
	3732:  uint16(sym_method_index_expression),
	3733:  uint16(177),
	3734:  uint16(1),
	3735:  uint16(sym_parenthesized_expression),
	3736:  uint16(194),
	3737:  uint16(1),
	3738:  uint16(sym_return_statement),
	3739:  uint16(255),
	3740:  uint16(1),
	3741:  uint16(sym__variable_assignment_varlist),
	3742:  uint16(279),
	3743:  uint16(1),
	3744:  uint16(sym__block),
	3745:  uint16(2),
	3746:  uint16(3),
	3747:  uint16(sym_bracket_index_expression),
	3748:  uint16(sym_dot_index_expression),
	3749:  uint16(sym__contextual_keyword),
	3750:  uint16(160),
	3751:  uint16(9),
	3752:  uint16(sym_empty_statement),
	3753:  uint16(sym_assignment_statement),
	3754:  uint16(sym_label_statement),
	3755:  uint16(sym_goto_statement),
	3756:  uint16(sym_do_statement),
	3757:  uint16(sym_while_statement),
	3758:  uint16(sym_repeat_statement),
	3759:  uint16(sym_if_statement),
	3760:  uint16(sym_for_statement),
	3761:  uint16(38),
	3762:  uint16(3),
	3763:  uint16(1),
	3764:  uint16(anon_sym_DASH_DASH),
	3765:  uint16(5),
	3766:  uint16(1),
	3767:  uint16(sym__block_comment_start),
	3768:  uint16(9),
	3769:  uint16(1),
	3770:  uint16(sym_identifier),
	3771:  uint16(15),
	3772:  uint16(1),
	3773:  uint16(anon_sym_SEMI),
	3774:  uint16(17),
	3775:  uint16(1),
	3776:  uint16(anon_sym_COLON_COLON),
	3777:  uint16(19),
	3778:  uint16(1),
	3779:  uint16(sym_break_statement),
	3780:  uint16(21),
	3781:  uint16(1),
	3782:  uint16(anon_sym_goto),
	3783:  uint16(23),
	3784:  uint16(1),
	3785:  uint16(anon_sym_do),
	3786:  uint16(25),
	3787:  uint16(1),
	3788:  uint16(anon_sym_while),
	3789:  uint16(27),
	3790:  uint16(1),
	3791:  uint16(anon_sym_repeat),
	3792:  uint16(29),
	3793:  uint16(1),
	3794:  uint16(anon_sym_if),
	3795:  uint16(31),
	3796:  uint16(1),
	3797:  uint16(anon_sym_for),
	3798:  uint16(33),
	3799:  uint16(1),
	3800:  uint16(anon_sym_function),
	3801:  uint16(35),
	3802:  uint16(1),
	3803:  uint16(anon_sym_local),
	3804:  uint16(37),
	3805:  uint16(1),
	3806:  uint16(anon_sym_global),
	3807:  uint16(39),
	3808:  uint16(1),
	3809:  uint16(anon_sym_LPAREN),
	3810:  uint16(113),
	3811:  uint16(1),
	3812:  uint16(anon_sym_return),
	3813:  uint16(225),
	3814:  uint16(1),
	3815:  uint16(anon_sym_end),
	3816:  uint16(21),
	3817:  uint16(1),
	3818:  uint16(aux_sym_chunk_repeat1),
	3819:  uint16(51),
	3820:  uint16(1),
	3821:  uint16(sym_comment),
	3822:  uint16(101),
	3823:  uint16(1),
	3824:  uint16(sym_function_call),
	3825:  uint16(131),
	3826:  uint16(1),
	3827:  uint16(sym_variable_declaration),
	3828:  uint16(132),
	3829:  uint16(1),
	3830:  uint16(sym_declaration),
	3831:  uint16(134),
	3832:  uint16(1),
	3833:  uint16(sym_function_declaration),
	3834:  uint16(137),
	3835:  uint16(1),
	3836:  uint16(sym__local_function_declaration),
	3837:  uint16(139),
	3838:  uint16(1),
	3839:  uint16(sym__global_function_declaration),
	3840:  uint16(142),
	3841:  uint16(1),
	3842:  uint16(sym__global_variable_declaration),
	3843:  uint16(143),
	3844:  uint16(1),
	3845:  uint16(sym__global_implicit_variable_declaration),
	3846:  uint16(161),
	3847:  uint16(1),
	3848:  uint16(sym_statement),
	3849:  uint16(170),
	3850:  uint16(1),
	3851:  uint16(sym__prefix_expression),
	3852:  uint16(172),
	3853:  uint16(1),
	3854:  uint16(sym_variable),
	3855:  uint16(174),
	3856:  uint16(1),
	3857:  uint16(sym_method_index_expression),
	3858:  uint16(177),
	3859:  uint16(1),
	3860:  uint16(sym_parenthesized_expression),
	3861:  uint16(194),
	3862:  uint16(1),
	3863:  uint16(sym_return_statement),
	3864:  uint16(255),
	3865:  uint16(1),
	3866:  uint16(sym__variable_assignment_varlist),
	3867:  uint16(270),
	3868:  uint16(1),
	3869:  uint16(sym__block),
	3870:  uint16(2),
	3871:  uint16(3),
	3872:  uint16(sym_bracket_index_expression),
	3873:  uint16(sym_dot_index_expression),
	3874:  uint16(sym__contextual_keyword),
	3875:  uint16(160),
	3876:  uint16(9),
	3877:  uint16(sym_empty_statement),
	3878:  uint16(sym_assignment_statement),
	3879:  uint16(sym_label_statement),
	3880:  uint16(sym_goto_statement),
	3881:  uint16(sym_do_statement),
	3882:  uint16(sym_while_statement),
	3883:  uint16(sym_repeat_statement),
	3884:  uint16(sym_if_statement),
	3885:  uint16(sym_for_statement),
	3886:  uint16(38),
	3887:  uint16(3),
	3888:  uint16(1),
	3889:  uint16(anon_sym_DASH_DASH),
	3890:  uint16(5),
	3891:  uint16(1),
	3892:  uint16(sym__block_comment_start),
	3893:  uint16(9),
	3894:  uint16(1),
	3895:  uint16(sym_identifier),
	3896:  uint16(15),
	3897:  uint16(1),
	3898:  uint16(anon_sym_SEMI),
	3899:  uint16(17),
	3900:  uint16(1),
	3901:  uint16(anon_sym_COLON_COLON),
	3902:  uint16(19),
	3903:  uint16(1),
	3904:  uint16(sym_break_statement),
	3905:  uint16(21),
	3906:  uint16(1),
	3907:  uint16(anon_sym_goto),
	3908:  uint16(23),
	3909:  uint16(1),
	3910:  uint16(anon_sym_do),
	3911:  uint16(25),
	3912:  uint16(1),
	3913:  uint16(anon_sym_while),
	3914:  uint16(27),
	3915:  uint16(1),
	3916:  uint16(anon_sym_repeat),
	3917:  uint16(29),
	3918:  uint16(1),
	3919:  uint16(anon_sym_if),
	3920:  uint16(31),
	3921:  uint16(1),
	3922:  uint16(anon_sym_for),
	3923:  uint16(33),
	3924:  uint16(1),
	3925:  uint16(anon_sym_function),
	3926:  uint16(35),
	3927:  uint16(1),
	3928:  uint16(anon_sym_local),
	3929:  uint16(37),
	3930:  uint16(1),
	3931:  uint16(anon_sym_global),
	3932:  uint16(39),
	3933:  uint16(1),
	3934:  uint16(anon_sym_LPAREN),
	3935:  uint16(113),
	3936:  uint16(1),
	3937:  uint16(anon_sym_return),
	3938:  uint16(227),
	3939:  uint16(1),
	3940:  uint16(anon_sym_end),
	3941:  uint16(21),
	3942:  uint16(1),
	3943:  uint16(aux_sym_chunk_repeat1),
	3944:  uint16(52),
	3945:  uint16(1),
	3946:  uint16(sym_comment),
	3947:  uint16(101),
	3948:  uint16(1),
	3949:  uint16(sym_function_call),
	3950:  uint16(131),
	3951:  uint16(1),
	3952:  uint16(sym_variable_declaration),
	3953:  uint16(132),
	3954:  uint16(1),
	3955:  uint16(sym_declaration),
	3956:  uint16(134),
	3957:  uint16(1),
	3958:  uint16(sym_function_declaration),
	3959:  uint16(137),
	3960:  uint16(1),
	3961:  uint16(sym__local_function_declaration),
	3962:  uint16(139),
	3963:  uint16(1),
	3964:  uint16(sym__global_function_declaration),
	3965:  uint16(142),
	3966:  uint16(1),
	3967:  uint16(sym__global_variable_declaration),
	3968:  uint16(143),
	3969:  uint16(1),
	3970:  uint16(sym__global_implicit_variable_declaration),
	3971:  uint16(161),
	3972:  uint16(1),
	3973:  uint16(sym_statement),
	3974:  uint16(170),
	3975:  uint16(1),
	3976:  uint16(sym__prefix_expression),
	3977:  uint16(172),
	3978:  uint16(1),
	3979:  uint16(sym_variable),
	3980:  uint16(174),
	3981:  uint16(1),
	3982:  uint16(sym_method_index_expression),
	3983:  uint16(177),
	3984:  uint16(1),
	3985:  uint16(sym_parenthesized_expression),
	3986:  uint16(194),
	3987:  uint16(1),
	3988:  uint16(sym_return_statement),
	3989:  uint16(255),
	3990:  uint16(1),
	3991:  uint16(sym__variable_assignment_varlist),
	3992:  uint16(271),
	3993:  uint16(1),
	3994:  uint16(sym__block),
	3995:  uint16(2),
	3996:  uint16(3),
	3997:  uint16(sym_bracket_index_expression),
	3998:  uint16(sym_dot_index_expression),
	3999:  uint16(sym__contextual_keyword),
	4000:  uint16(160),
	4001:  uint16(9),
	4002:  uint16(sym_empty_statement),
	4003:  uint16(sym_assignment_statement),
	4004:  uint16(sym_label_statement),
	4005:  uint16(sym_goto_statement),
	4006:  uint16(sym_do_statement),
	4007:  uint16(sym_while_statement),
	4008:  uint16(sym_repeat_statement),
	4009:  uint16(sym_if_statement),
	4010:  uint16(sym_for_statement),
	4011:  uint16(37),
	4012:  uint16(3),
	4013:  uint16(1),
	4014:  uint16(anon_sym_DASH_DASH),
	4015:  uint16(5),
	4016:  uint16(1),
	4017:  uint16(sym__block_comment_start),
	4018:  uint16(9),
	4019:  uint16(1),
	4020:  uint16(sym_identifier),
	4021:  uint16(13),
	4022:  uint16(1),
	4023:  uint16(anon_sym_return),
	4024:  uint16(15),
	4025:  uint16(1),
	4026:  uint16(anon_sym_SEMI),
	4027:  uint16(17),
	4028:  uint16(1),
	4029:  uint16(anon_sym_COLON_COLON),
	4030:  uint16(19),
	4031:  uint16(1),
	4032:  uint16(sym_break_statement),
	4033:  uint16(21),
	4034:  uint16(1),
	4035:  uint16(anon_sym_goto),
	4036:  uint16(23),
	4037:  uint16(1),
	4038:  uint16(anon_sym_do),
	4039:  uint16(25),
	4040:  uint16(1),
	4041:  uint16(anon_sym_while),
	4042:  uint16(27),
	4043:  uint16(1),
	4044:  uint16(anon_sym_repeat),
	4045:  uint16(29),
	4046:  uint16(1),
	4047:  uint16(anon_sym_if),
	4048:  uint16(31),
	4049:  uint16(1),
	4050:  uint16(anon_sym_for),
	4051:  uint16(33),
	4052:  uint16(1),
	4053:  uint16(anon_sym_function),
	4054:  uint16(35),
	4055:  uint16(1),
	4056:  uint16(anon_sym_local),
	4057:  uint16(37),
	4058:  uint16(1),
	4059:  uint16(anon_sym_global),
	4060:  uint16(39),
	4061:  uint16(1),
	4062:  uint16(anon_sym_LPAREN),
	4063:  uint16(229),
	4064:  uint16(1),
	4066:  uint16(53),
	4067:  uint16(1),
	4068:  uint16(sym_comment),
	4069:  uint16(54),
	4070:  uint16(1),
	4071:  uint16(aux_sym_chunk_repeat1),
	4072:  uint16(101),
	4073:  uint16(1),
	4074:  uint16(sym_function_call),
	4075:  uint16(131),
	4076:  uint16(1),
	4077:  uint16(sym_variable_declaration),
	4078:  uint16(132),
	4079:  uint16(1),
	4080:  uint16(sym_declaration),
	4081:  uint16(134),
	4082:  uint16(1),
	4083:  uint16(sym_function_declaration),
	4084:  uint16(137),
	4085:  uint16(1),
	4086:  uint16(sym__local_function_declaration),
	4087:  uint16(139),
	4088:  uint16(1),
	4089:  uint16(sym__global_function_declaration),
	4090:  uint16(142),
	4091:  uint16(1),
	4092:  uint16(sym__global_variable_declaration),
	4093:  uint16(143),
	4094:  uint16(1),
	4095:  uint16(sym__global_implicit_variable_declaration),
	4096:  uint16(161),
	4097:  uint16(1),
	4098:  uint16(sym_statement),
	4099:  uint16(170),
	4100:  uint16(1),
	4101:  uint16(sym__prefix_expression),
	4102:  uint16(172),
	4103:  uint16(1),
	4104:  uint16(sym_variable),
	4105:  uint16(174),
	4106:  uint16(1),
	4107:  uint16(sym_method_index_expression),
	4108:  uint16(177),
	4109:  uint16(1),
	4110:  uint16(sym_parenthesized_expression),
	4111:  uint16(255),
	4112:  uint16(1),
	4113:  uint16(sym__variable_assignment_varlist),
	4114:  uint16(266),
	4115:  uint16(1),
	4116:  uint16(sym_return_statement),
	4117:  uint16(2),
	4118:  uint16(3),
	4119:  uint16(sym_bracket_index_expression),
	4120:  uint16(sym_dot_index_expression),
	4121:  uint16(sym__contextual_keyword),
	4122:  uint16(160),
	4123:  uint16(9),
	4124:  uint16(sym_empty_statement),
	4125:  uint16(sym_assignment_statement),
	4126:  uint16(sym_label_statement),
	4127:  uint16(sym_goto_statement),
	4128:  uint16(sym_do_statement),
	4129:  uint16(sym_while_statement),
	4130:  uint16(sym_repeat_statement),
	4131:  uint16(sym_if_statement),
	4132:  uint16(sym_for_statement),
	4133:  uint16(37),
	4134:  uint16(3),
	4135:  uint16(1),
	4136:  uint16(anon_sym_DASH_DASH),
	4137:  uint16(5),
	4138:  uint16(1),
	4139:  uint16(sym__block_comment_start),
	4140:  uint16(9),
	4141:  uint16(1),
	4142:  uint16(sym_identifier),
	4143:  uint16(13),
	4144:  uint16(1),
	4145:  uint16(anon_sym_return),
	4146:  uint16(15),
	4147:  uint16(1),
	4148:  uint16(anon_sym_SEMI),
	4149:  uint16(17),
	4150:  uint16(1),
	4151:  uint16(anon_sym_COLON_COLON),
	4152:  uint16(19),
	4153:  uint16(1),
	4154:  uint16(sym_break_statement),
	4155:  uint16(21),
	4156:  uint16(1),
	4157:  uint16(anon_sym_goto),
	4158:  uint16(23),
	4159:  uint16(1),
	4160:  uint16(anon_sym_do),
	4161:  uint16(25),
	4162:  uint16(1),
	4163:  uint16(anon_sym_while),
	4164:  uint16(27),
	4165:  uint16(1),
	4166:  uint16(anon_sym_repeat),
	4167:  uint16(29),
	4168:  uint16(1),
	4169:  uint16(anon_sym_if),
	4170:  uint16(31),
	4171:  uint16(1),
	4172:  uint16(anon_sym_for),
	4173:  uint16(33),
	4174:  uint16(1),
	4175:  uint16(anon_sym_function),
	4176:  uint16(35),
	4177:  uint16(1),
	4178:  uint16(anon_sym_local),
	4179:  uint16(37),
	4180:  uint16(1),
	4181:  uint16(anon_sym_global),
	4182:  uint16(39),
	4183:  uint16(1),
	4184:  uint16(anon_sym_LPAREN),
	4185:  uint16(231),
	4186:  uint16(1),
	4188:  uint16(22),
	4189:  uint16(1),
	4190:  uint16(aux_sym_chunk_repeat1),
	4191:  uint16(54),
	4192:  uint16(1),
	4193:  uint16(sym_comment),
	4194:  uint16(101),
	4195:  uint16(1),
	4196:  uint16(sym_function_call),
	4197:  uint16(131),
	4198:  uint16(1),
	4199:  uint16(sym_variable_declaration),
	4200:  uint16(132),
	4201:  uint16(1),
	4202:  uint16(sym_declaration),
	4203:  uint16(134),
	4204:  uint16(1),
	4205:  uint16(sym_function_declaration),
	4206:  uint16(137),
	4207:  uint16(1),
	4208:  uint16(sym__local_function_declaration),
	4209:  uint16(139),
	4210:  uint16(1),
	4211:  uint16(sym__global_function_declaration),
	4212:  uint16(142),
	4213:  uint16(1),
	4214:  uint16(sym__global_variable_declaration),
	4215:  uint16(143),
	4216:  uint16(1),
	4217:  uint16(sym__global_implicit_variable_declaration),
	4218:  uint16(161),
	4219:  uint16(1),
	4220:  uint16(sym_statement),
	4221:  uint16(170),
	4222:  uint16(1),
	4223:  uint16(sym__prefix_expression),
	4224:  uint16(172),
	4225:  uint16(1),
	4226:  uint16(sym_variable),
	4227:  uint16(174),
	4228:  uint16(1),
	4229:  uint16(sym_method_index_expression),
	4230:  uint16(177),
	4231:  uint16(1),
	4232:  uint16(sym_parenthesized_expression),
	4233:  uint16(245),
	4234:  uint16(1),
	4235:  uint16(sym_return_statement),
	4236:  uint16(255),
	4237:  uint16(1),
	4238:  uint16(sym__variable_assignment_varlist),
	4239:  uint16(2),
	4240:  uint16(3),
	4241:  uint16(sym_bracket_index_expression),
	4242:  uint16(sym_dot_index_expression),
	4243:  uint16(sym__contextual_keyword),
	4244:  uint16(160),
	4245:  uint16(9),
	4246:  uint16(sym_empty_statement),
	4247:  uint16(sym_assignment_statement),
	4248:  uint16(sym_label_statement),
	4249:  uint16(sym_goto_statement),
	4250:  uint16(sym_do_statement),
	4251:  uint16(sym_while_statement),
	4252:  uint16(sym_repeat_statement),
	4253:  uint16(sym_if_statement),
	4254:  uint16(sym_for_statement),
	4255:  uint16(37),
	4256:  uint16(3),
	4257:  uint16(1),
	4258:  uint16(anon_sym_DASH_DASH),
	4259:  uint16(5),
	4260:  uint16(1),
	4261:  uint16(sym__block_comment_start),
	4262:  uint16(9),
	4263:  uint16(1),
	4264:  uint16(sym_identifier),
	4265:  uint16(13),
	4266:  uint16(1),
	4267:  uint16(anon_sym_return),
	4268:  uint16(15),
	4269:  uint16(1),
	4270:  uint16(anon_sym_SEMI),
	4271:  uint16(17),
	4272:  uint16(1),
	4273:  uint16(anon_sym_COLON_COLON),
	4274:  uint16(19),
	4275:  uint16(1),
	4276:  uint16(sym_break_statement),
	4277:  uint16(21),
	4278:  uint16(1),
	4279:  uint16(anon_sym_goto),
	4280:  uint16(23),
	4281:  uint16(1),
	4282:  uint16(anon_sym_do),
	4283:  uint16(25),
	4284:  uint16(1),
	4285:  uint16(anon_sym_while),
	4286:  uint16(27),
	4287:  uint16(1),
	4288:  uint16(anon_sym_repeat),
	4289:  uint16(29),
	4290:  uint16(1),
	4291:  uint16(anon_sym_if),
	4292:  uint16(31),
	4293:  uint16(1),
	4294:  uint16(anon_sym_for),
	4295:  uint16(33),
	4296:  uint16(1),
	4297:  uint16(anon_sym_function),
	4298:  uint16(35),
	4299:  uint16(1),
	4300:  uint16(anon_sym_local),
	4301:  uint16(37),
	4302:  uint16(1),
	4303:  uint16(anon_sym_global),
	4304:  uint16(39),
	4305:  uint16(1),
	4306:  uint16(anon_sym_LPAREN),
	4307:  uint16(229),
	4308:  uint16(1),
	4310:  uint16(22),
	4311:  uint16(1),
	4312:  uint16(aux_sym_chunk_repeat1),
	4313:  uint16(55),
	4314:  uint16(1),
	4315:  uint16(sym_comment),
	4316:  uint16(101),
	4317:  uint16(1),
	4318:  uint16(sym_function_call),
	4319:  uint16(131),
	4320:  uint16(1),
	4321:  uint16(sym_variable_declaration),
	4322:  uint16(132),
	4323:  uint16(1),
	4324:  uint16(sym_declaration),
	4325:  uint16(134),
	4326:  uint16(1),
	4327:  uint16(sym_function_declaration),
	4328:  uint16(137),
	4329:  uint16(1),
	4330:  uint16(sym__local_function_declaration),
	4331:  uint16(139),
	4332:  uint16(1),
	4333:  uint16(sym__global_function_declaration),
	4334:  uint16(142),
	4335:  uint16(1),
	4336:  uint16(sym__global_variable_declaration),
	4337:  uint16(143),
	4338:  uint16(1),
	4339:  uint16(sym__global_implicit_variable_declaration),
	4340:  uint16(161),
	4341:  uint16(1),
	4342:  uint16(sym_statement),
	4343:  uint16(170),
	4344:  uint16(1),
	4345:  uint16(sym__prefix_expression),
	4346:  uint16(172),
	4347:  uint16(1),
	4348:  uint16(sym_variable),
	4349:  uint16(174),
	4350:  uint16(1),
	4351:  uint16(sym_method_index_expression),
	4352:  uint16(177),
	4353:  uint16(1),
	4354:  uint16(sym_parenthesized_expression),
	4355:  uint16(255),
	4356:  uint16(1),
	4357:  uint16(sym__variable_assignment_varlist),
	4358:  uint16(266),
	4359:  uint16(1),
	4360:  uint16(sym_return_statement),
	4361:  uint16(2),
	4362:  uint16(3),
	4363:  uint16(sym_bracket_index_expression),
	4364:  uint16(sym_dot_index_expression),
	4365:  uint16(sym__contextual_keyword),
	4366:  uint16(160),
	4367:  uint16(9),
	4368:  uint16(sym_empty_statement),
	4369:  uint16(sym_assignment_statement),
	4370:  uint16(sym_label_statement),
	4371:  uint16(sym_goto_statement),
	4372:  uint16(sym_do_statement),
	4373:  uint16(sym_while_statement),
	4374:  uint16(sym_repeat_statement),
	4375:  uint16(sym_if_statement),
	4376:  uint16(sym_for_statement),
	4377:  uint16(21),
	4378:  uint16(3),
	4379:  uint16(1),
	4380:  uint16(anon_sym_DASH_DASH),
	4381:  uint16(5),
	4382:  uint16(1),
	4383:  uint16(sym__block_comment_start),
	4384:  uint16(177),
	4385:  uint16(1),
	4386:  uint16(anon_sym_AMP),
	4387:  uint16(181),
	4388:  uint16(1),
	4389:  uint16(anon_sym_PLUS),
	4390:  uint16(183),
	4391:  uint16(1),
	4392:  uint16(anon_sym_DASH),
	4393:  uint16(185),
	4394:  uint16(1),
	4395:  uint16(anon_sym_SLASH),
	4396:  uint16(187),
	4397:  uint16(1),
	4398:  uint16(anon_sym_DOT_DOT),
	4399:  uint16(189),
	4400:  uint16(1),
	4401:  uint16(anon_sym_CARET),
	4402:  uint16(195),
	4403:  uint16(1),
	4404:  uint16(anon_sym_PIPE),
	4405:  uint16(197),
	4406:  uint16(1),
	4407:  uint16(anon_sym_TILDE),
	4408:  uint16(199),
	4409:  uint16(1),
	4410:  uint16(anon_sym_and),
	4411:  uint16(237),
	4412:  uint16(1),
	4413:  uint16(anon_sym_COMMA),
	4414:  uint16(239),
	4415:  uint16(1),
	4416:  uint16(anon_sym_or),
	4417:  uint16(56),
	4418:  uint16(1),
	4419:  uint16(sym_comment),
	4420:  uint16(128),
	4421:  uint16(1),
	4422:  uint16(aux_sym__variable_assignment_explist_repeat1),
	4423:  uint16(179),
	4424:  uint16(2),
	4425:  uint16(anon_sym_LT_LT),
	4426:  uint16(anon_sym_GT_GT),
	4427:  uint16(191),
	4428:  uint16(2),
	4429:  uint16(anon_sym_LT),
	4430:  uint16(anon_sym_GT),
	4431:  uint16(175),
	4432:  uint16(3),
	4433:  uint16(anon_sym_STAR),
	4434:  uint16(anon_sym_SLASH_SLASH),
	4435:  uint16(anon_sym_PERCENT),
	4436:  uint16(193),
	4437:  uint16(4),
	4438:  uint16(anon_sym_LT_EQ),
	4439:  uint16(anon_sym_EQ_EQ),
	4440:  uint16(anon_sym_TILDE_EQ),
	4441:  uint16(anon_sym_GT_EQ),
	4442:  uint16(233),
	4443:  uint16(4),
	4445:  uint16(anon_sym_SEMI),
	4446:  uint16(anon_sym_COLON_COLON),
	4447:  uint16(anon_sym_LPAREN),
	4448:  uint16(235),
	4449:  uint16(16),
	4450:  uint16(anon_sym_return),
	4451:  uint16(sym_break_statement),
	4452:  uint16(anon_sym_goto),
	4453:  uint16(anon_sym_do),
	4454:  uint16(anon_sym_end),
	4455:  uint16(anon_sym_while),
	4456:  uint16(anon_sym_repeat),
	4457:  uint16(anon_sym_until),
	4458:  uint16(anon_sym_if),
	4459:  uint16(anon_sym_elseif),
	4460:  uint16(anon_sym_else),
	4461:  uint16(anon_sym_for),
	4462:  uint16(anon_sym_function),
	4463:  uint16(anon_sym_local),
	4464:  uint16(anon_sym_global),
	4465:  uint16(sym_identifier),
	4466:  uint16(19),
	4467:  uint16(3),
	4468:  uint16(1),
	4469:  uint16(anon_sym_DASH_DASH),
	4470:  uint16(5),
	4471:  uint16(1),
	4472:  uint16(sym__block_comment_start),
	4473:  uint16(177),
	4474:  uint16(1),
	4475:  uint16(anon_sym_AMP),
	4476:  uint16(181),
	4477:  uint16(1),
	4478:  uint16(anon_sym_PLUS),
	4479:  uint16(183),
	4480:  uint16(1),
	4481:  uint16(anon_sym_DASH),
	4482:  uint16(185),
	4483:  uint16(1),
	4484:  uint16(anon_sym_SLASH),
	4485:  uint16(187),
	4486:  uint16(1),
	4487:  uint16(anon_sym_DOT_DOT),
	4488:  uint16(189),
	4489:  uint16(1),
	4490:  uint16(anon_sym_CARET),
	4491:  uint16(195),
	4492:  uint16(1),
	4493:  uint16(anon_sym_PIPE),
	4494:  uint16(197),
	4495:  uint16(1),
	4496:  uint16(anon_sym_TILDE),
	4497:  uint16(199),
	4498:  uint16(1),
	4499:  uint16(anon_sym_and),
	4500:  uint16(239),
	4501:  uint16(1),
	4502:  uint16(anon_sym_or),
	4503:  uint16(57),
	4504:  uint16(1),
	4505:  uint16(sym_comment),
	4506:  uint16(179),
	4507:  uint16(2),
	4508:  uint16(anon_sym_LT_LT),
	4509:  uint16(anon_sym_GT_GT),
	4510:  uint16(191),
	4511:  uint16(2),
	4512:  uint16(anon_sym_LT),
	4513:  uint16(anon_sym_GT),
	4514:  uint16(175),
	4515:  uint16(3),
	4516:  uint16(anon_sym_STAR),
	4517:  uint16(anon_sym_SLASH_SLASH),
	4518:  uint16(anon_sym_PERCENT),
	4519:  uint16(193),
	4520:  uint16(4),
	4521:  uint16(anon_sym_LT_EQ),
	4522:  uint16(anon_sym_EQ_EQ),
	4523:  uint16(anon_sym_TILDE_EQ),
	4524:  uint16(anon_sym_GT_EQ),
	4525:  uint16(241),
	4526:  uint16(5),
	4528:  uint16(anon_sym_SEMI),
	4529:  uint16(anon_sym_COMMA),
	4530:  uint16(anon_sym_COLON_COLON),
	4531:  uint16(anon_sym_LPAREN),
	4532:  uint16(243),
	4533:  uint16(16),
	4534:  uint16(anon_sym_return),
	4535:  uint16(sym_break_statement),
	4536:  uint16(anon_sym_goto),
	4537:  uint16(anon_sym_do),
	4538:  uint16(anon_sym_end),
	4539:  uint16(anon_sym_while),
	4540:  uint16(anon_sym_repeat),
	4541:  uint16(anon_sym_until),
	4542:  uint16(anon_sym_if),
	4543:  uint16(anon_sym_elseif),
	4544:  uint16(anon_sym_else),
	4545:  uint16(anon_sym_for),
	4546:  uint16(anon_sym_function),
	4547:  uint16(anon_sym_local),
	4548:  uint16(anon_sym_global),
	4549:  uint16(sym_identifier),
	4550:  uint16(19),
	4551:  uint16(3),
	4552:  uint16(1),
	4553:  uint16(anon_sym_DASH_DASH),
	4554:  uint16(5),
	4555:  uint16(1),
	4556:  uint16(sym__block_comment_start),
	4557:  uint16(177),
	4558:  uint16(1),
	4559:  uint16(anon_sym_AMP),
	4560:  uint16(181),
	4561:  uint16(1),
	4562:  uint16(anon_sym_PLUS),
	4563:  uint16(183),
	4564:  uint16(1),
	4565:  uint16(anon_sym_DASH),
	4566:  uint16(185),
	4567:  uint16(1),
	4568:  uint16(anon_sym_SLASH),
	4569:  uint16(187),
	4570:  uint16(1),
	4571:  uint16(anon_sym_DOT_DOT),
	4572:  uint16(189),
	4573:  uint16(1),
	4574:  uint16(anon_sym_CARET),
	4575:  uint16(195),
	4576:  uint16(1),
	4577:  uint16(anon_sym_PIPE),
	4578:  uint16(197),
	4579:  uint16(1),
	4580:  uint16(anon_sym_TILDE),
	4581:  uint16(199),
	4582:  uint16(1),
	4583:  uint16(anon_sym_and),
	4584:  uint16(239),
	4585:  uint16(1),
	4586:  uint16(anon_sym_or),
	4587:  uint16(58),
	4588:  uint16(1),
	4589:  uint16(sym_comment),
	4590:  uint16(179),
	4591:  uint16(2),
	4592:  uint16(anon_sym_LT_LT),
	4593:  uint16(anon_sym_GT_GT),
	4594:  uint16(191),
	4595:  uint16(2),
	4596:  uint16(anon_sym_LT),
	4597:  uint16(anon_sym_GT),
	4598:  uint16(175),
	4599:  uint16(3),
	4600:  uint16(anon_sym_STAR),
	4601:  uint16(anon_sym_SLASH_SLASH),
	4602:  uint16(anon_sym_PERCENT),
	4603:  uint16(193),
	4604:  uint16(4),
	4605:  uint16(anon_sym_LT_EQ),
	4606:  uint16(anon_sym_EQ_EQ),
	4607:  uint16(anon_sym_TILDE_EQ),
	4608:  uint16(anon_sym_GT_EQ),
	4609:  uint16(245),
	4610:  uint16(4),
	4612:  uint16(anon_sym_SEMI),
	4613:  uint16(anon_sym_COLON_COLON),
	4614:  uint16(anon_sym_LPAREN),
	4615:  uint16(247),
	4616:  uint16(16),
	4617:  uint16(anon_sym_return),
	4618:  uint16(sym_break_statement),
	4619:  uint16(anon_sym_goto),
	4620:  uint16(anon_sym_do),
	4621:  uint16(anon_sym_end),
	4622:  uint16(anon_sym_while),
	4623:  uint16(anon_sym_repeat),
	4624:  uint16(anon_sym_until),
	4625:  uint16(anon_sym_if),
	4626:  uint16(anon_sym_elseif),
	4627:  uint16(anon_sym_else),
	4628:  uint16(anon_sym_for),
	4629:  uint16(anon_sym_function),
	4630:  uint16(anon_sym_local),
	4631:  uint16(anon_sym_global),
	4632:  uint16(sym_identifier),
	4633:  uint16(19),
	4634:  uint16(3),
	4635:  uint16(1),
	4636:  uint16(anon_sym_DASH_DASH),
	4637:  uint16(5),
	4638:  uint16(1),
	4639:  uint16(sym__block_comment_start),
	4640:  uint16(177),
	4641:  uint16(1),
	4642:  uint16(anon_sym_AMP),
	4643:  uint16(181),
	4644:  uint16(1),
	4645:  uint16(anon_sym_PLUS),
	4646:  uint16(183),
	4647:  uint16(1),
	4648:  uint16(anon_sym_DASH),
	4649:  uint16(185),
	4650:  uint16(1),
	4651:  uint16(anon_sym_SLASH),
	4652:  uint16(187),
	4653:  uint16(1),
	4654:  uint16(anon_sym_DOT_DOT),
	4655:  uint16(189),
	4656:  uint16(1),
	4657:  uint16(anon_sym_CARET),
	4658:  uint16(195),
	4659:  uint16(1),
	4660:  uint16(anon_sym_PIPE),
	4661:  uint16(197),
	4662:  uint16(1),
	4663:  uint16(anon_sym_TILDE),
	4664:  uint16(199),
	4665:  uint16(1),
	4666:  uint16(anon_sym_and),
	4667:  uint16(239),
	4668:  uint16(1),
	4669:  uint16(anon_sym_or),
	4670:  uint16(59),
	4671:  uint16(1),
	4672:  uint16(sym_comment),
	4673:  uint16(179),
	4674:  uint16(2),
	4675:  uint16(anon_sym_LT_LT),
	4676:  uint16(anon_sym_GT_GT),
	4677:  uint16(191),
	4678:  uint16(2),
	4679:  uint16(anon_sym_LT),
	4680:  uint16(anon_sym_GT),
	4681:  uint16(175),
	4682:  uint16(3),
	4683:  uint16(anon_sym_STAR),
	4684:  uint16(anon_sym_SLASH_SLASH),
	4685:  uint16(anon_sym_PERCENT),
	4686:  uint16(193),
	4687:  uint16(4),
	4688:  uint16(anon_sym_LT_EQ),
	4689:  uint16(anon_sym_EQ_EQ),
	4690:  uint16(anon_sym_TILDE_EQ),
	4691:  uint16(anon_sym_GT_EQ),
	4692:  uint16(249),
	4693:  uint16(4),
	4695:  uint16(anon_sym_SEMI),
	4696:  uint16(anon_sym_COLON_COLON),
	4697:  uint16(anon_sym_LPAREN),
	4698:  uint16(251),
	4699:  uint16(16),
	4700:  uint16(anon_sym_return),
	4701:  uint16(sym_break_statement),
	4702:  uint16(anon_sym_goto),
	4703:  uint16(anon_sym_do),
	4704:  uint16(anon_sym_end),
	4705:  uint16(anon_sym_while),
	4706:  uint16(anon_sym_repeat),
	4707:  uint16(anon_sym_until),
	4708:  uint16(anon_sym_if),
	4709:  uint16(anon_sym_elseif),
	4710:  uint16(anon_sym_else),
	4711:  uint16(anon_sym_for),
	4712:  uint16(anon_sym_function),
	4713:  uint16(anon_sym_local),
	4714:  uint16(anon_sym_global),
	4715:  uint16(sym_identifier),
	4716:  uint16(26),
	4717:  uint16(3),
	4718:  uint16(1),
	4719:  uint16(anon_sym_DASH_DASH),
	4720:  uint16(5),
	4721:  uint16(1),
	4722:  uint16(sym__block_comment_start),
	4723:  uint16(9),
	4724:  uint16(1),
	4725:  uint16(sym_identifier),
	4726:  uint16(39),
	4727:  uint16(1),
	4728:  uint16(anon_sym_LPAREN),
	4729:  uint16(253),
	4730:  uint16(1),
	4731:  uint16(anon_sym_SEMI),
	4732:  uint16(257),
	4733:  uint16(1),
	4734:  uint16(anon_sym_function),
	4735:  uint16(259),
	4736:  uint16(1),
	4737:  uint16(anon_sym_global),
	4738:  uint16(265),
	4739:  uint16(1),
	4740:  uint16(anon_sym_DQUOTE),
	4741:  uint16(267),
	4742:  uint16(1),
	4743:  uint16(anon_sym_SQUOTE),
	4744:  uint16(269),
	4745:  uint16(1),
	4746:  uint16(anon_sym_LBRACE),
	4747:  uint16(275),
	4748:  uint16(1),
	4749:  uint16(sym__block_string_start),
	4750:  uint16(41),
	4751:  uint16(1),
	4752:  uint16(sym__quote_string),
	4753:  uint16(45),
	4754:  uint16(1),
	4755:  uint16(sym__block_string),
	4756:  uint16(60),
	4757:  uint16(1),
	4758:  uint16(sym_comment),
	4759:  uint16(100),
	4760:  uint16(1),
	4761:  uint16(sym_expression),
	4762:  uint16(170),
	4763:  uint16(1),
	4764:  uint16(sym__prefix_expression),
	4765:  uint16(174),
	4766:  uint16(1),
	4767:  uint16(sym_method_index_expression),
	4768:  uint16(182),
	4769:  uint16(1),
	4770:  uint16(sym__expression_list),
	4771:  uint16(263),
	4772:  uint16(2),
	4773:  uint16(sym_number),
	4774:  uint16(sym_vararg_expression),
	4775:  uint16(271),
	4776:  uint16(2),
	4777:  uint16(anon_sym_TILDE),
	4778:  uint16(anon_sym_POUND),
	4779:  uint16(273),
	4780:  uint16(2),
	4781:  uint16(anon_sym_DASH),
	4782:  uint16(anon_sym_not),
	4783:  uint16(261),
	4784:  uint16(3),
	4785:  uint16(sym_nil),
	4786:  uint16(sym_false),
	4787:  uint16(sym_true),
	4788:  uint16(2),
	4789:  uint16(3),
	4790:  uint16(sym_bracket_index_expression),
	4791:  uint16(sym_dot_index_expression),
	4792:  uint16(sym__contextual_keyword),
	4793:  uint16(7),
	4794:  uint16(3),
	4795:  uint16(sym_variable),
	4796:  uint16(sym_function_call),
	4797:  uint16(sym_parenthesized_expression),
	4798:  uint16(255),
	4799:  uint16(4),
	4800:  uint16(anon_sym_end),
	4801:  uint16(anon_sym_until),
	4802:  uint16(anon_sym_elseif),
	4803:  uint16(anon_sym_else),
	4804:  uint16(40),
	4805:  uint16(5),
	4806:  uint16(sym_string),
	4807:  uint16(sym_function_definition),
	4808:  uint16(sym_table_constructor),
	4809:  uint16(sym_binary_expression),
	4810:  uint16(sym_unary_expression),
	4811:  uint16(28),
	4812:  uint16(3),
	4813:  uint16(1),
	4814:  uint16(anon_sym_DASH_DASH),
	4815:  uint16(5),
	4816:  uint16(1),
	4817:  uint16(sym__block_comment_start),
	4818:  uint16(39),
	4819:  uint16(1),
	4820:  uint16(anon_sym_LPAREN),
	4821:  uint16(257),
	4822:  uint16(1),
	4823:  uint16(anon_sym_function),
	4824:  uint16(259),
	4825:  uint16(1),
	4826:  uint16(anon_sym_global),
	4827:  uint16(265),
	4828:  uint16(1),
	4829:  uint16(anon_sym_DQUOTE),
	4830:  uint16(267),
	4831:  uint16(1),
	4832:  uint16(anon_sym_SQUOTE),
	4833:  uint16(269),
	4834:  uint16(1),
	4835:  uint16(anon_sym_LBRACE),
	4836:  uint16(275),
	4837:  uint16(1),
	4838:  uint16(sym__block_string_start),
	4839:  uint16(277),
	4840:  uint16(1),
	4841:  uint16(sym_identifier),
	4842:  uint16(279),
	4843:  uint16(1),
	4844:  uint16(anon_sym_LBRACK),
	4845:  uint16(281),
	4846:  uint16(1),
	4847:  uint16(anon_sym_RBRACE),
	4848:  uint16(41),
	4849:  uint16(1),
	4850:  uint16(sym__quote_string),
	4851:  uint16(45),
	4852:  uint16(1),
	4853:  uint16(sym__block_string),
	4854:  uint16(61),
	4855:  uint16(1),
	4856:  uint16(sym_comment),
	4857:  uint16(76),
	4858:  uint16(1),
	4859:  uint16(sym__contextual_keyword),
	4860:  uint16(105),
	4861:  uint16(1),
	4862:  uint16(sym_expression),
	4863:  uint16(170),
	4864:  uint16(1),
	4865:  uint16(sym__prefix_expression),
	4866:  uint16(174),
	4867:  uint16(1),
	4868:  uint16(sym_method_index_expression),
	4869:  uint16(188),
	4870:  uint16(1),
	4871:  uint16(sym_field),
	4872:  uint16(242),
	4873:  uint16(1),
	4874:  uint16(sym__field_list),
	4875:  uint16(263),
	4876:  uint16(2),
	4877:  uint16(sym_number),
	4878:  uint16(sym_vararg_expression),
	4879:  uint16(271),
	4880:  uint16(2),
	4881:  uint16(anon_sym_TILDE),
	4882:  uint16(anon_sym_POUND),
	4883:  uint16(273),
	4884:  uint16(2),
	4885:  uint16(anon_sym_DASH),
	4886:  uint16(anon_sym_not),
	4887:  uint16(2),
	4888:  uint16(2),
	4889:  uint16(sym_bracket_index_expression),
	4890:  uint16(sym_dot_index_expression),
	4891:  uint16(261),
	4892:  uint16(3),
	4893:  uint16(sym_nil),
	4894:  uint16(sym_false),
	4895:  uint16(sym_true),
	4896:  uint16(7),
	4897:  uint16(3),
	4898:  uint16(sym_variable),
	4899:  uint16(sym_function_call),
	4900:  uint16(sym_parenthesized_expression),
	4901:  uint16(40),
	4902:  uint16(5),
	4903:  uint16(sym_string),
	4904:  uint16(sym_function_definition),
	4905:  uint16(sym_table_constructor),
	4906:  uint16(sym_binary_expression),
	4907:  uint16(sym_unary_expression),
	4908:  uint16(28),
	4909:  uint16(3),
	4910:  uint16(1),
	4911:  uint16(anon_sym_DASH_DASH),
	4912:  uint16(5),
	4913:  uint16(1),
	4914:  uint16(sym__block_comment_start),
	4915:  uint16(39),
	4916:  uint16(1),
	4917:  uint16(anon_sym_LPAREN),
	4918:  uint16(257),
	4919:  uint16(1),
	4920:  uint16(anon_sym_function),
	4921:  uint16(259),
	4922:  uint16(1),
	4923:  uint16(anon_sym_global),
	4924:  uint16(265),
	4925:  uint16(1),
	4926:  uint16(anon_sym_DQUOTE),
	4927:  uint16(267),
	4928:  uint16(1),
	4929:  uint16(anon_sym_SQUOTE),
	4930:  uint16(269),
	4931:  uint16(1),
	4932:  uint16(anon_sym_LBRACE),
	4933:  uint16(275),
	4934:  uint16(1),
	4935:  uint16(sym__block_string_start),
	4936:  uint16(277),
	4937:  uint16(1),
	4938:  uint16(sym_identifier),
	4939:  uint16(279),
	4940:  uint16(1),
	4941:  uint16(anon_sym_LBRACK),
	4942:  uint16(283),
	4943:  uint16(1),
	4944:  uint16(anon_sym_RBRACE),
	4945:  uint16(41),
	4946:  uint16(1),
	4947:  uint16(sym__quote_string),
	4948:  uint16(45),
	4949:  uint16(1),
	4950:  uint16(sym__block_string),
	4951:  uint16(62),
	4952:  uint16(1),
	4953:  uint16(sym_comment),
	4954:  uint16(76),
	4955:  uint16(1),
	4956:  uint16(sym__contextual_keyword),
	4957:  uint16(105),
	4958:  uint16(1),
	4959:  uint16(sym_expression),
	4960:  uint16(170),
	4961:  uint16(1),
	4962:  uint16(sym__prefix_expression),
	4963:  uint16(174),
	4964:  uint16(1),
	4965:  uint16(sym_method_index_expression),
	4966:  uint16(188),
	4967:  uint16(1),
	4968:  uint16(sym_field),
	4969:  uint16(277),
	4970:  uint16(1),
	4971:  uint16(sym__field_list),
	4972:  uint16(263),
	4973:  uint16(2),
	4974:  uint16(sym_number),
	4975:  uint16(sym_vararg_expression),
	4976:  uint16(271),
	4977:  uint16(2),
	4978:  uint16(anon_sym_TILDE),
	4979:  uint16(anon_sym_POUND),
	4980:  uint16(273),
	4981:  uint16(2),
	4982:  uint16(anon_sym_DASH),
	4983:  uint16(anon_sym_not),
	4984:  uint16(2),
	4985:  uint16(2),
	4986:  uint16(sym_bracket_index_expression),
	4987:  uint16(sym_dot_index_expression),
	4988:  uint16(261),
	4989:  uint16(3),
	4990:  uint16(sym_nil),
	4991:  uint16(sym_false),
	4992:  uint16(sym_true),
	4993:  uint16(7),
	4994:  uint16(3),
	4995:  uint16(sym_variable),
	4996:  uint16(sym_function_call),
	4997:  uint16(sym_parenthesized_expression),
	4998:  uint16(40),
	4999:  uint16(5),
	5000:  uint16(sym_string),
	5001:  uint16(sym_function_definition),
	5002:  uint16(sym_table_constructor),
	5003:  uint16(sym_binary_expression),
	5004:  uint16(sym_unary_expression),
	5005:  uint16(26),
	5006:  uint16(3),
	5007:  uint16(1),
	5008:  uint16(anon_sym_DASH_DASH),
	5009:  uint16(5),
	5010:  uint16(1),
	5011:  uint16(sym__block_comment_start),
	5012:  uint16(9),
	5013:  uint16(1),
	5014:  uint16(sym_identifier),
	5015:  uint16(39),
	5016:  uint16(1),
	5017:  uint16(anon_sym_LPAREN),
	5018:  uint16(253),
	5019:  uint16(1),
	5020:  uint16(anon_sym_SEMI),
	5021:  uint16(257),
	5022:  uint16(1),
	5023:  uint16(anon_sym_function),
	5024:  uint16(259),
	5025:  uint16(1),
	5026:  uint16(anon_sym_global),
	5027:  uint16(265),
	5028:  uint16(1),
	5029:  uint16(anon_sym_DQUOTE),
	5030:  uint16(267),
	5031:  uint16(1),
	5032:  uint16(anon_sym_SQUOTE),
	5033:  uint16(269),
	5034:  uint16(1),
	5035:  uint16(anon_sym_LBRACE),
	5036:  uint16(275),
	5037:  uint16(1),
	5038:  uint16(sym__block_string_start),
	5039:  uint16(285),
	5040:  uint16(1),
	5042:  uint16(41),
	5043:  uint16(1),
	5044:  uint16(sym__quote_string),
	5045:  uint16(45),
	5046:  uint16(1),
	5047:  uint16(sym__block_string),
	5048:  uint16(63),
	5049:  uint16(1),
	5050:  uint16(sym_comment),
	5051:  uint16(100),
	5052:  uint16(1),
	5053:  uint16(sym_expression),
	5054:  uint16(170),
	5055:  uint16(1),
	5056:  uint16(sym__prefix_expression),
	5057:  uint16(174),
	5058:  uint16(1),
	5059:  uint16(sym_method_index_expression),
	5060:  uint16(182),
	5061:  uint16(1),
	5062:  uint16(sym__expression_list),
	5063:  uint16(263),
	5064:  uint16(2),
	5065:  uint16(sym_number),
	5066:  uint16(sym_vararg_expression),
	5067:  uint16(271),
	5068:  uint16(2),
	5069:  uint16(anon_sym_TILDE),
	5070:  uint16(anon_sym_POUND),
	5071:  uint16(273),
	5072:  uint16(2),
	5073:  uint16(anon_sym_DASH),
	5074:  uint16(anon_sym_not),
	5075:  uint16(261),
	5076:  uint16(3),
	5077:  uint16(sym_nil),
	5078:  uint16(sym_false),
	5079:  uint16(sym_true),
	5080:  uint16(2),
	5081:  uint16(3),
	5082:  uint16(sym_bracket_index_expression),
	5083:  uint16(sym_dot_index_expression),
	5084:  uint16(sym__contextual_keyword),
	5085:  uint16(7),
	5086:  uint16(3),
	5087:  uint16(sym_variable),
	5088:  uint16(sym_function_call),
	5089:  uint16(sym_parenthesized_expression),
	5090:  uint16(40),
	5091:  uint16(5),
	5092:  uint16(sym_string),
	5093:  uint16(sym_function_definition),
	5094:  uint16(sym_table_constructor),
	5095:  uint16(sym_binary_expression),
	5096:  uint16(sym_unary_expression),
	5097:  uint16(27),
	5098:  uint16(3),
	5099:  uint16(1),
	5100:  uint16(anon_sym_DASH_DASH),
	5101:  uint16(5),
	5102:  uint16(1),
	5103:  uint16(sym__block_comment_start),
	5104:  uint16(39),
	5105:  uint16(1),
	5106:  uint16(anon_sym_LPAREN),
	5107:  uint16(257),
	5108:  uint16(1),
	5109:  uint16(anon_sym_function),
	5110:  uint16(259),
	5111:  uint16(1),
	5112:  uint16(anon_sym_global),
	5113:  uint16(265),
	5114:  uint16(1),
	5115:  uint16(anon_sym_DQUOTE),
	5116:  uint16(267),
	5117:  uint16(1),
	5118:  uint16(anon_sym_SQUOTE),
	5119:  uint16(269),
	5120:  uint16(1),
	5121:  uint16(anon_sym_LBRACE),
	5122:  uint16(275),
	5123:  uint16(1),
	5124:  uint16(sym__block_string_start),
	5125:  uint16(277),
	5126:  uint16(1),
	5127:  uint16(sym_identifier),
	5128:  uint16(279),
	5129:  uint16(1),
	5130:  uint16(anon_sym_LBRACK),
	5131:  uint16(287),
	5132:  uint16(1),
	5133:  uint16(anon_sym_RBRACE),
	5134:  uint16(41),
	5135:  uint16(1),
	5136:  uint16(sym__quote_string),
	5137:  uint16(45),
	5138:  uint16(1),
	5139:  uint16(sym__block_string),
	5140:  uint16(64),
	5141:  uint16(1),
	5142:  uint16(sym_comment),
	5143:  uint16(76),
	5144:  uint16(1),
	5145:  uint16(sym__contextual_keyword),
	5146:  uint16(105),
	5147:  uint16(1),
	5148:  uint16(sym_expression),
	5149:  uint16(170),
	5150:  uint16(1),
	5151:  uint16(sym__prefix_expression),
	5152:  uint16(174),
	5153:  uint16(1),
	5154:  uint16(sym_method_index_expression),
	5155:  uint16(211),
	5156:  uint16(1),
	5157:  uint16(sym_field),
	5158:  uint16(263),
	5159:  uint16(2),
	5160:  uint16(sym_number),
	5161:  uint16(sym_vararg_expression),
	5162:  uint16(271),
	5163:  uint16(2),
	5164:  uint16(anon_sym_TILDE),
	5165:  uint16(anon_sym_POUND),
	5166:  uint16(273),
	5167:  uint16(2),
	5168:  uint16(anon_sym_DASH),
	5169:  uint16(anon_sym_not),
	5170:  uint16(2),
	5171:  uint16(2),
	5172:  uint16(sym_bracket_index_expression),
	5173:  uint16(sym_dot_index_expression),
	5174:  uint16(261),
	5175:  uint16(3),
	5176:  uint16(sym_nil),
	5177:  uint16(sym_false),
	5178:  uint16(sym_true),
	5179:  uint16(7),
	5180:  uint16(3),
	5181:  uint16(sym_variable),
	5182:  uint16(sym_function_call),
	5183:  uint16(sym_parenthesized_expression),
	5184:  uint16(40),
	5185:  uint16(5),
	5186:  uint16(sym_string),
	5187:  uint16(sym_function_definition),
	5188:  uint16(sym_table_constructor),
	5189:  uint16(sym_binary_expression),
	5190:  uint16(sym_unary_expression),
	5191:  uint16(27),
	5192:  uint16(3),
	5193:  uint16(1),
	5194:  uint16(anon_sym_DASH_DASH),
	5195:  uint16(5),
	5196:  uint16(1),
	5197:  uint16(sym__block_comment_start),
	5198:  uint16(39),
	5199:  uint16(1),
	5200:  uint16(anon_sym_LPAREN),
	5201:  uint16(257),
	5202:  uint16(1),
	5203:  uint16(anon_sym_function),
	5204:  uint16(259),
	5205:  uint16(1),
	5206:  uint16(anon_sym_global),
	5207:  uint16(265),
	5208:  uint16(1),
	5209:  uint16(anon_sym_DQUOTE),
	5210:  uint16(267),
	5211:  uint16(1),
	5212:  uint16(anon_sym_SQUOTE),
	5213:  uint16(269),
	5214:  uint16(1),
	5215:  uint16(anon_sym_LBRACE),
	5216:  uint16(275),
	5217:  uint16(1),
	5218:  uint16(sym__block_string_start),
	5219:  uint16(277),
	5220:  uint16(1),
	5221:  uint16(sym_identifier),
	5222:  uint16(279),
	5223:  uint16(1),
	5224:  uint16(anon_sym_LBRACK),
	5225:  uint16(289),
	5226:  uint16(1),
	5227:  uint16(anon_sym_RBRACE),
	5228:  uint16(41),
	5229:  uint16(1),
	5230:  uint16(sym__quote_string),
	5231:  uint16(45),
	5232:  uint16(1),
	5233:  uint16(sym__block_string),
	5234:  uint16(65),
	5235:  uint16(1),
	5236:  uint16(sym_comment),
	5237:  uint16(76),
	5238:  uint16(1),
	5239:  uint16(sym__contextual_keyword),
	5240:  uint16(105),
	5241:  uint16(1),
	5242:  uint16(sym_expression),
	5243:  uint16(170),
	5244:  uint16(1),
	5245:  uint16(sym__prefix_expression),
	5246:  uint16(174),
	5247:  uint16(1),
	5248:  uint16(sym_method_index_expression),
	5249:  uint16(211),
	5250:  uint16(1),
	5251:  uint16(sym_field),
	5252:  uint16(263),
	5253:  uint16(2),
	5254:  uint16(sym_number),
	5255:  uint16(sym_vararg_expression),
	5256:  uint16(271),
	5257:  uint16(2),
	5258:  uint16(anon_sym_TILDE),
	5259:  uint16(anon_sym_POUND),
	5260:  uint16(273),
	5261:  uint16(2),
	5262:  uint16(anon_sym_DASH),
	5263:  uint16(anon_sym_not),
	5264:  uint16(2),
	5265:  uint16(2),
	5266:  uint16(sym_bracket_index_expression),
	5267:  uint16(sym_dot_index_expression),
	5268:  uint16(261),
	5269:  uint16(3),
	5270:  uint16(sym_nil),
	5271:  uint16(sym_false),
	5272:  uint16(sym_true),
	5273:  uint16(7),
	5274:  uint16(3),
	5275:  uint16(sym_variable),
	5276:  uint16(sym_function_call),
	5277:  uint16(sym_parenthesized_expression),
	5278:  uint16(40),
	5279:  uint16(5),
	5280:  uint16(sym_string),
	5281:  uint16(sym_function_definition),
	5282:  uint16(sym_table_constructor),
	5283:  uint16(sym_binary_expression),
	5284:  uint16(sym_unary_expression),
	5285:  uint16(26),
	5286:  uint16(3),
	5287:  uint16(1),
	5288:  uint16(anon_sym_DASH_DASH),
	5289:  uint16(5),
	5290:  uint16(1),
	5291:  uint16(sym__block_comment_start),
	5292:  uint16(39),
	5293:  uint16(1),
	5294:  uint16(anon_sym_LPAREN),
	5295:  uint16(257),
	5296:  uint16(1),
	5297:  uint16(anon_sym_function),
	5298:  uint16(259),
	5299:  uint16(1),
	5300:  uint16(anon_sym_global),
	5301:  uint16(265),
	5302:  uint16(1),
	5303:  uint16(anon_sym_DQUOTE),
	5304:  uint16(267),
	5305:  uint16(1),
	5306:  uint16(anon_sym_SQUOTE),
	5307:  uint16(269),
	5308:  uint16(1),
	5309:  uint16(anon_sym_LBRACE),
	5310:  uint16(275),
	5311:  uint16(1),
	5312:  uint16(sym__block_string_start),
	5313:  uint16(277),
	5314:  uint16(1),
	5315:  uint16(sym_identifier),
	5316:  uint16(279),
	5317:  uint16(1),
	5318:  uint16(anon_sym_LBRACK),
	5319:  uint16(41),
	5320:  uint16(1),
	5321:  uint16(sym__quote_string),
	5322:  uint16(45),
	5323:  uint16(1),
	5324:  uint16(sym__block_string),
	5325:  uint16(66),
	5326:  uint16(1),
	5327:  uint16(sym_comment),
	5328:  uint16(76),
	5329:  uint16(1),
	5330:  uint16(sym__contextual_keyword),
	5331:  uint16(105),
	5332:  uint16(1),
	5333:  uint16(sym_expression),
	5334:  uint16(170),
	5335:  uint16(1),
	5336:  uint16(sym__prefix_expression),
	5337:  uint16(174),
	5338:  uint16(1),
	5339:  uint16(sym_method_index_expression),
	5340:  uint16(211),
	5341:  uint16(1),
	5342:  uint16(sym_field),
	5343:  uint16(263),
	5344:  uint16(2),
	5345:  uint16(sym_number),
	5346:  uint16(sym_vararg_expression),
	5347:  uint16(271),
	5348:  uint16(2),
	5349:  uint16(anon_sym_TILDE),
	5350:  uint16(anon_sym_POUND),
	5351:  uint16(273),
	5352:  uint16(2),
	5353:  uint16(anon_sym_DASH),
	5354:  uint16(anon_sym_not),
	5355:  uint16(2),
	5356:  uint16(2),
	5357:  uint16(sym_bracket_index_expression),
	5358:  uint16(sym_dot_index_expression),
	5359:  uint16(261),
	5360:  uint16(3),
	5361:  uint16(sym_nil),
	5362:  uint16(sym_false),
	5363:  uint16(sym_true),
	5364:  uint16(7),
	5365:  uint16(3),
	5366:  uint16(sym_variable),
	5367:  uint16(sym_function_call),
	5368:  uint16(sym_parenthesized_expression),
	5369:  uint16(40),
	5370:  uint16(5),
	5371:  uint16(sym_string),
	5372:  uint16(sym_function_definition),
	5373:  uint16(sym_table_constructor),
	5374:  uint16(sym_binary_expression),
	5375:  uint16(sym_unary_expression),
	5376:  uint16(24),
	5377:  uint16(3),
	5378:  uint16(1),
	5379:  uint16(anon_sym_DASH_DASH),
	5380:  uint16(5),
	5381:  uint16(1),
	5382:  uint16(sym__block_comment_start),
	5383:  uint16(9),
	5384:  uint16(1),
	5385:  uint16(sym_identifier),
	5386:  uint16(39),
	5387:  uint16(1),
	5388:  uint16(anon_sym_LPAREN),
	5389:  uint16(257),
	5390:  uint16(1),
	5391:  uint16(anon_sym_function),
	5392:  uint16(259),
	5393:  uint16(1),
	5394:  uint16(anon_sym_global),
	5395:  uint16(265),
	5396:  uint16(1),
	5397:  uint16(anon_sym_DQUOTE),
	5398:  uint16(267),
	5399:  uint16(1),
	5400:  uint16(anon_sym_SQUOTE),
	5401:  uint16(269),
	5402:  uint16(1),
	5403:  uint16(anon_sym_LBRACE),
	5404:  uint16(275),
	5405:  uint16(1),
	5406:  uint16(sym__block_string_start),
	5407:  uint16(41),
	5408:  uint16(1),
	5409:  uint16(sym__quote_string),
	5410:  uint16(45),
	5411:  uint16(1),
	5412:  uint16(sym__block_string),
	5413:  uint16(56),
	5414:  uint16(1),
	5415:  uint16(sym_expression),
	5416:  uint16(67),
	5417:  uint16(1),
	5418:  uint16(sym_comment),
	5419:  uint16(136),
	5420:  uint16(1),
	5421:  uint16(sym__variable_assignment_explist),
	5422:  uint16(170),
	5423:  uint16(1),
	5424:  uint16(sym__prefix_expression),
	5425:  uint16(174),
	5426:  uint16(1),
	5427:  uint16(sym_method_index_expression),
	5428:  uint16(263),
	5429:  uint16(2),
	5430:  uint16(sym_number),
	5431:  uint16(sym_vararg_expression),
	5432:  uint16(271),
	5433:  uint16(2),
	5434:  uint16(anon_sym_TILDE),
	5435:  uint16(anon_sym_POUND),
	5436:  uint16(273),
	5437:  uint16(2),
	5438:  uint16(anon_sym_DASH),
	5439:  uint16(anon_sym_not),
	5440:  uint16(261),
	5441:  uint16(3),
	5442:  uint16(sym_nil),
	5443:  uint16(sym_false),
	5444:  uint16(sym_true),
	5445:  uint16(2),
	5446:  uint16(3),
	5447:  uint16(sym_bracket_index_expression),
	5448:  uint16(sym_dot_index_expression),
	5449:  uint16(sym__contextual_keyword),
	5450:  uint16(7),
	5451:  uint16(3),
	5452:  uint16(sym_variable),
	5453:  uint16(sym_function_call),
	5454:  uint16(sym_parenthesized_expression),
	5455:  uint16(40),
	5456:  uint16(5),
	5457:  uint16(sym_string),
	5458:  uint16(sym_function_definition),
	5459:  uint16(sym_table_constructor),
	5460:  uint16(sym_binary_expression),
	5461:  uint16(sym_unary_expression),
	5462:  uint16(24),
	5463:  uint16(3),
	5464:  uint16(1),
	5465:  uint16(anon_sym_DASH_DASH),
	5466:  uint16(5),
	5467:  uint16(1),
	5468:  uint16(sym__block_comment_start),
	5469:  uint16(9),
	5470:  uint16(1),
	5471:  uint16(sym_identifier),
	5472:  uint16(39),
	5473:  uint16(1),
	5474:  uint16(anon_sym_LPAREN),
	5475:  uint16(257),
	5476:  uint16(1),
	5477:  uint16(anon_sym_function),
	5478:  uint16(259),
	5479:  uint16(1),
	5480:  uint16(anon_sym_global),
	5481:  uint16(265),
	5482:  uint16(1),
	5483:  uint16(anon_sym_DQUOTE),
	5484:  uint16(267),
	5485:  uint16(1),
	5486:  uint16(anon_sym_SQUOTE),
	5487:  uint16(269),
	5488:  uint16(1),
	5489:  uint16(anon_sym_LBRACE),
	5490:  uint16(275),
	5491:  uint16(1),
	5492:  uint16(sym__block_string_start),
	5493:  uint16(41),
	5494:  uint16(1),
	5495:  uint16(sym__quote_string),
	5496:  uint16(45),
	5497:  uint16(1),
	5498:  uint16(sym__block_string),
	5499:  uint16(56),
	5500:  uint16(1),
	5501:  uint16(sym_expression),
	5502:  uint16(68),
	5503:  uint16(1),
	5504:  uint16(sym_comment),
	5505:  uint16(149),
	5506:  uint16(1),
	5507:  uint16(sym__variable_assignment_explist),
	5508:  uint16(170),
	5509:  uint16(1),
	5510:  uint16(sym__prefix_expression),
	5511:  uint16(174),
	5512:  uint16(1),
	5513:  uint16(sym_method_index_expression),
	5514:  uint16(263),
	5515:  uint16(2),
	5516:  uint16(sym_number),
	5517:  uint16(sym_vararg_expression),
	5518:  uint16(271),
	5519:  uint16(2),
	5520:  uint16(anon_sym_TILDE),
	5521:  uint16(anon_sym_POUND),
	5522:  uint16(273),
	5523:  uint16(2),
	5524:  uint16(anon_sym_DASH),
	5525:  uint16(anon_sym_not),
	5526:  uint16(261),
	5527:  uint16(3),
	5528:  uint16(sym_nil),
	5529:  uint16(sym_false),
	5530:  uint16(sym_true),
	5531:  uint16(2),
	5532:  uint16(3),
	5533:  uint16(sym_bracket_index_expression),
	5534:  uint16(sym_dot_index_expression),
	5535:  uint16(sym__contextual_keyword),
	5536:  uint16(7),
	5537:  uint16(3),
	5538:  uint16(sym_variable),
	5539:  uint16(sym_function_call),
	5540:  uint16(sym_parenthesized_expression),
	5541:  uint16(40),
	5542:  uint16(5),
	5543:  uint16(sym_string),
	5544:  uint16(sym_function_definition),
	5545:  uint16(sym_table_constructor),
	5546:  uint16(sym_binary_expression),
	5547:  uint16(sym_unary_expression),
	5548:  uint16(24),
	5549:  uint16(3),
	5550:  uint16(1),
	5551:  uint16(anon_sym_DASH_DASH),
	5552:  uint16(5),
	5553:  uint16(1),
	5554:  uint16(sym__block_comment_start),
	5555:  uint16(9),
	5556:  uint16(1),
	5557:  uint16(sym_identifier),
	5558:  uint16(39),
	5559:  uint16(1),
	5560:  uint16(anon_sym_LPAREN),
	5561:  uint16(257),
	5562:  uint16(1),
	5563:  uint16(anon_sym_function),
	5564:  uint16(259),
	5565:  uint16(1),
	5566:  uint16(anon_sym_global),
	5567:  uint16(265),
	5568:  uint16(1),
	5569:  uint16(anon_sym_DQUOTE),
	5570:  uint16(267),
	5571:  uint16(1),
	5572:  uint16(anon_sym_SQUOTE),
	5573:  uint16(269),
	5574:  uint16(1),
	5575:  uint16(anon_sym_LBRACE),
	5576:  uint16(275),
	5577:  uint16(1),
	5578:  uint16(sym__block_string_start),
	5579:  uint16(291),
	5580:  uint16(1),
	5581:  uint16(anon_sym_RPAREN),
	5582:  uint16(41),
	5583:  uint16(1),
	5584:  uint16(sym__quote_string),
	5585:  uint16(45),
	5586:  uint16(1),
	5587:  uint16(sym__block_string),
	5588:  uint16(69),
	5589:  uint16(1),
	5590:  uint16(sym_comment),
	5591:  uint16(104),
	5592:  uint16(1),
	5593:  uint16(sym_expression),
	5594:  uint16(170),
	5595:  uint16(1),
	5596:  uint16(sym__prefix_expression),
	5597:  uint16(174),
	5598:  uint16(1),
	5599:  uint16(sym_method_index_expression),
	5600:  uint16(263),
	5601:  uint16(2),
	5602:  uint16(sym_number),
	5603:  uint16(sym_vararg_expression),
	5604:  uint16(271),
	5605:  uint16(2),
	5606:  uint16(anon_sym_TILDE),
	5607:  uint16(anon_sym_POUND),
	5608:  uint16(273),
	5609:  uint16(2),
	5610:  uint16(anon_sym_DASH),
	5611:  uint16(anon_sym_not),
	5612:  uint16(261),
	5613:  uint16(3),
	5614:  uint16(sym_nil),
	5615:  uint16(sym_false),
	5616:  uint16(sym_true),
	5617:  uint16(2),
	5618:  uint16(3),
	5619:  uint16(sym_bracket_index_expression),
	5620:  uint16(sym_dot_index_expression),
	5621:  uint16(sym__contextual_keyword),
	5622:  uint16(7),
	5623:  uint16(3),
	5624:  uint16(sym_variable),
	5625:  uint16(sym_function_call),
	5626:  uint16(sym_parenthesized_expression),
	5627:  uint16(40),
	5628:  uint16(5),
	5629:  uint16(sym_string),
	5630:  uint16(sym_function_definition),
	5631:  uint16(sym_table_constructor),
	5632:  uint16(sym_binary_expression),
	5633:  uint16(sym_unary_expression),
	5634:  uint16(24),
	5635:  uint16(3),
	5636:  uint16(1),
	5637:  uint16(anon_sym_DASH_DASH),
	5638:  uint16(5),
	5639:  uint16(1),
	5640:  uint16(sym__block_comment_start),
	5641:  uint16(9),
	5642:  uint16(1),
	5643:  uint16(sym_identifier),
	5644:  uint16(39),
	5645:  uint16(1),
	5646:  uint16(anon_sym_LPAREN),
	5647:  uint16(257),
	5648:  uint16(1),
	5649:  uint16(anon_sym_function),
	5650:  uint16(259),
	5651:  uint16(1),
	5652:  uint16(anon_sym_global),
	5653:  uint16(265),
	5654:  uint16(1),
	5655:  uint16(anon_sym_DQUOTE),
	5656:  uint16(267),
	5657:  uint16(1),
	5658:  uint16(anon_sym_SQUOTE),
	5659:  uint16(269),
	5660:  uint16(1),
	5661:  uint16(anon_sym_LBRACE),
	5662:  uint16(275),
	5663:  uint16(1),
	5664:  uint16(sym__block_string_start),
	5665:  uint16(41),
	5666:  uint16(1),
	5667:  uint16(sym__quote_string),
	5668:  uint16(45),
	5669:  uint16(1),
	5670:  uint16(sym__block_string),
	5671:  uint16(70),
	5672:  uint16(1),
	5673:  uint16(sym_comment),
	5674:  uint16(100),
	5675:  uint16(1),
	5676:  uint16(sym_expression),
	5677:  uint16(170),
	5678:  uint16(1),
	5679:  uint16(sym__prefix_expression),
	5680:  uint16(174),
	5681:  uint16(1),
	5682:  uint16(sym_method_index_expression),
	5683:  uint16(231),
	5684:  uint16(1),
	5685:  uint16(sym__expression_list),
	5686:  uint16(263),
	5687:  uint16(2),
	5688:  uint16(sym_number),
	5689:  uint16(sym_vararg_expression),
	5690:  uint16(271),
	5691:  uint16(2),
	5692:  uint16(anon_sym_TILDE),
	5693:  uint16(anon_sym_POUND),
	5694:  uint16(273),
	5695:  uint16(2),
	5696:  uint16(anon_sym_DASH),
	5697:  uint16(anon_sym_not),
	5698:  uint16(261),
	5699:  uint16(3),
	5700:  uint16(sym_nil),
	5701:  uint16(sym_false),
	5702:  uint16(sym_true),
	5703:  uint16(2),
	5704:  uint16(3),
	5705:  uint16(sym_bracket_index_expression),
	5706:  uint16(sym_dot_index_expression),
	5707:  uint16(sym__contextual_keyword),
	5708:  uint16(7),
	5709:  uint16(3),
	5710:  uint16(sym_variable),
	5711:  uint16(sym_function_call),
	5712:  uint16(sym_parenthesized_expression),
	5713:  uint16(40),
	5714:  uint16(5),
	5715:  uint16(sym_string),
	5716:  uint16(sym_function_definition),
	5717:  uint16(sym_table_constructor),
	5718:  uint16(sym_binary_expression),
	5719:  uint16(sym_unary_expression),
	5720:  uint16(23),
	5721:  uint16(3),
	5722:  uint16(1),
	5723:  uint16(anon_sym_DASH_DASH),
	5724:  uint16(5),
	5725:  uint16(1),
	5726:  uint16(sym__block_comment_start),
	5727:  uint16(9),
	5728:  uint16(1),
	5729:  uint16(sym_identifier),
	5730:  uint16(39),
	5731:  uint16(1),
	5732:  uint16(anon_sym_LPAREN),
	5733:  uint16(257),
	5734:  uint16(1),
	5735:  uint16(anon_sym_function),
	5736:  uint16(259),
	5737:  uint16(1),
	5738:  uint16(anon_sym_global),
	5739:  uint16(265),
	5740:  uint16(1),
	5741:  uint16(anon_sym_DQUOTE),
	5742:  uint16(267),
	5743:  uint16(1),
	5744:  uint16(anon_sym_SQUOTE),
	5745:  uint16(269),
	5746:  uint16(1),
	5747:  uint16(anon_sym_LBRACE),
	5748:  uint16(275),
	5749:  uint16(1),
	5750:  uint16(sym__block_string_start),
	5751:  uint16(41),
	5752:  uint16(1),
	5753:  uint16(sym__quote_string),
	5754:  uint16(45),
	5755:  uint16(1),
	5756:  uint16(sym__block_string),
	5757:  uint16(71),
	5758:  uint16(1),
	5759:  uint16(sym_comment),
	5760:  uint16(99),
	5761:  uint16(1),
	5762:  uint16(sym_expression),
	5763:  uint16(170),
	5764:  uint16(1),
	5765:  uint16(sym__prefix_expression),
	5766:  uint16(174),
	5767:  uint16(1),
	5768:  uint16(sym_method_index_expression),
	5769:  uint16(263),
	5770:  uint16(2),
	5771:  uint16(sym_number),
	5772:  uint16(sym_vararg_expression),
	5773:  uint16(271),
	5774:  uint16(2),
	5775:  uint16(anon_sym_TILDE),
	5776:  uint16(anon_sym_POUND),
	5777:  uint16(273),
	5778:  uint16(2),
	5779:  uint16(anon_sym_DASH),
	5780:  uint16(anon_sym_not),
	5781:  uint16(261),
	5782:  uint16(3),
	5783:  uint16(sym_nil),
	5784:  uint16(sym_false),
	5785:  uint16(sym_true),
	5786:  uint16(2),
	5787:  uint16(3),
	5788:  uint16(sym_bracket_index_expression),
	5789:  uint16(sym_dot_index_expression),
	5790:  uint16(sym__contextual_keyword),
	5791:  uint16(7),
	5792:  uint16(3),
	5793:  uint16(sym_variable),
	5794:  uint16(sym_function_call),
	5795:  uint16(sym_parenthesized_expression),
	5796:  uint16(40),
	5797:  uint16(5),
	5798:  uint16(sym_string),
	5799:  uint16(sym_function_definition),
	5800:  uint16(sym_table_constructor),
	5801:  uint16(sym_binary_expression),
	5802:  uint16(sym_unary_expression),
	5803:  uint16(23),
	5804:  uint16(3),
	5805:  uint16(1),
	5806:  uint16(anon_sym_DASH_DASH),
	5807:  uint16(5),
	5808:  uint16(1),
	5809:  uint16(sym__block_comment_start),
	5810:  uint16(9),
	5811:  uint16(1),
	5812:  uint16(sym_identifier),
	5813:  uint16(39),
	5814:  uint16(1),
	5815:  uint16(anon_sym_LPAREN),
	5816:  uint16(257),
	5817:  uint16(1),
	5818:  uint16(anon_sym_function),
	5819:  uint16(259),
	5820:  uint16(1),
	5821:  uint16(anon_sym_global),
	5822:  uint16(265),
	5823:  uint16(1),
	5824:  uint16(anon_sym_DQUOTE),
	5825:  uint16(267),
	5826:  uint16(1),
	5827:  uint16(anon_sym_SQUOTE),
	5828:  uint16(269),
	5829:  uint16(1),
	5830:  uint16(anon_sym_LBRACE),
	5831:  uint16(275),
	5832:  uint16(1),
	5833:  uint16(sym__block_string_start),
	5834:  uint16(27),
	5835:  uint16(1),
	5836:  uint16(sym_expression),
	5837:  uint16(41),
	5838:  uint16(1),
	5839:  uint16(sym__quote_string),
	5840:  uint16(45),
	5841:  uint16(1),
	5842:  uint16(sym__block_string),
	5843:  uint16(72),
	5844:  uint16(1),
	5845:  uint16(sym_comment),
	5846:  uint16(170),
	5847:  uint16(1),
	5848:  uint16(sym__prefix_expression),
	5849:  uint16(174),
	5850:  uint16(1),
	5851:  uint16(sym_method_index_expression),
	5852:  uint16(263),
	5853:  uint16(2),
	5854:  uint16(sym_number),
	5855:  uint16(sym_vararg_expression),
	5856:  uint16(271),
	5857:  uint16(2),
	5858:  uint16(anon_sym_TILDE),
	5859:  uint16(anon_sym_POUND),
	5860:  uint16(273),
	5861:  uint16(2),
	5862:  uint16(anon_sym_DASH),
	5863:  uint16(anon_sym_not),
	5864:  uint16(261),
	5865:  uint16(3),
	5866:  uint16(sym_nil),
	5867:  uint16(sym_false),
	5868:  uint16(sym_true),
	5869:  uint16(2),
	5870:  uint16(3),
	5871:  uint16(sym_bracket_index_expression),
	5872:  uint16(sym_dot_index_expression),
	5873:  uint16(sym__contextual_keyword),
	5874:  uint16(7),
	5875:  uint16(3),
	5876:  uint16(sym_variable),
	5877:  uint16(sym_function_call),
	5878:  uint16(sym_parenthesized_expression),
	5879:  uint16(40),
	5880:  uint16(5),
	5881:  uint16(sym_string),
	5882:  uint16(sym_function_definition),
	5883:  uint16(sym_table_constructor),
	5884:  uint16(sym_binary_expression),
	5885:  uint16(sym_unary_expression),
	5886:  uint16(23),
	5887:  uint16(3),
	5888:  uint16(1),
	5889:  uint16(anon_sym_DASH_DASH),
	5890:  uint16(5),
	5891:  uint16(1),
	5892:  uint16(sym__block_comment_start),
	5893:  uint16(9),
	5894:  uint16(1),
	5895:  uint16(sym_identifier),
	5896:  uint16(39),
	5897:  uint16(1),
	5898:  uint16(anon_sym_LPAREN),
	5899:  uint16(257),
	5900:  uint16(1),
	5901:  uint16(anon_sym_function),
	5902:  uint16(259),
	5903:  uint16(1),
	5904:  uint16(anon_sym_global),
	5905:  uint16(265),
	5906:  uint16(1),
	5907:  uint16(anon_sym_DQUOTE),
	5908:  uint16(267),
	5909:  uint16(1),
	5910:  uint16(anon_sym_SQUOTE),
	5911:  uint16(269),
	5912:  uint16(1),
	5913:  uint16(anon_sym_LBRACE),
	5914:  uint16(275),
	5915:  uint16(1),
	5916:  uint16(sym__block_string_start),
	5917:  uint16(41),
	5918:  uint16(1),
	5919:  uint16(sym__quote_string),
	5920:  uint16(45),
	5921:  uint16(1),
	5922:  uint16(sym__block_string),
	5923:  uint16(73),
	5924:  uint16(1),
	5925:  uint16(sym_comment),
	5926:  uint16(127),
	5927:  uint16(1),
	5928:  uint16(sym_expression),
	5929:  uint16(170),
	5930:  uint16(1),
	5931:  uint16(sym__prefix_expression),
	5932:  uint16(174),
	5933:  uint16(1),
	5934:  uint16(sym_method_index_expression),
	5935:  uint16(263),
	5936:  uint16(2),
	5937:  uint16(sym_number),
	5938:  uint16(sym_vararg_expression),
	5939:  uint16(271),
	5940:  uint16(2),
	5941:  uint16(anon_sym_TILDE),
	5942:  uint16(anon_sym_POUND),
	5943:  uint16(273),
	5944:  uint16(2),
	5945:  uint16(anon_sym_DASH),
	5946:  uint16(anon_sym_not),
	5947:  uint16(261),
	5948:  uint16(3),
	5949:  uint16(sym_nil),
	5950:  uint16(sym_false),
	5951:  uint16(sym_true),
	5952:  uint16(2),
	5953:  uint16(3),
	5954:  uint16(sym_bracket_index_expression),
	5955:  uint16(sym_dot_index_expression),
	5956:  uint16(sym__contextual_keyword),
	5957:  uint16(7),
	5958:  uint16(3),
	5959:  uint16(sym_variable),
	5960:  uint16(sym_function_call),
	5961:  uint16(sym_parenthesized_expression),
	5962:  uint16(40),
	5963:  uint16(5),
	5964:  uint16(sym_string),
	5965:  uint16(sym_function_definition),
	5966:  uint16(sym_table_constructor),
	5967:  uint16(sym_binary_expression),
	5968:  uint16(sym_unary_expression),
	5969:  uint16(23),
	5970:  uint16(3),
	5971:  uint16(1),
	5972:  uint16(anon_sym_DASH_DASH),
	5973:  uint16(5),
	5974:  uint16(1),
	5975:  uint16(sym__block_comment_start),
	5976:  uint16(9),
	5977:  uint16(1),
	5978:  uint16(sym_identifier),
	5979:  uint16(39),
	5980:  uint16(1),
	5981:  uint16(anon_sym_LPAREN),
	5982:  uint16(257),
	5983:  uint16(1),
	5984:  uint16(anon_sym_function),
	5985:  uint16(259),
	5986:  uint16(1),
	5987:  uint16(anon_sym_global),
	5988:  uint16(265),
	5989:  uint16(1),
	5990:  uint16(anon_sym_DQUOTE),
	5991:  uint16(267),
	5992:  uint16(1),
	5993:  uint16(anon_sym_SQUOTE),
	5994:  uint16(269),
	5995:  uint16(1),
	5996:  uint16(anon_sym_LBRACE),
	5997:  uint16(275),
	5998:  uint16(1),
	5999:  uint16(sym__block_string_start),
	6000:  uint16(41),
	6001:  uint16(1),
	6002:  uint16(sym__quote_string),
	6003:  uint16(45),
	6004:  uint16(1),
	6005:  uint16(sym__block_string),
	6006:  uint16(74),
	6007:  uint16(1),
	6008:  uint16(sym_comment),
	6009:  uint16(118),
	6010:  uint16(1),
	6011:  uint16(sym_expression),
	6012:  uint16(170),
	6013:  uint16(1),
	6014:  uint16(sym__prefix_expression),
	6015:  uint16(174),
	6016:  uint16(1),
	6017:  uint16(sym_method_index_expression),
	6018:  uint16(263),
	6019:  uint16(2),
	6020:  uint16(sym_number),
	6021:  uint16(sym_vararg_expression),
	6022:  uint16(271),
	6023:  uint16(2),
	6024:  uint16(anon_sym_TILDE),
	6025:  uint16(anon_sym_POUND),
	6026:  uint16(273),
	6027:  uint16(2),
	6028:  uint16(anon_sym_DASH),
	6029:  uint16(anon_sym_not),
	6030:  uint16(261),
	6031:  uint16(3),
	6032:  uint16(sym_nil),
	6033:  uint16(sym_false),
	6034:  uint16(sym_true),
	6035:  uint16(2),
	6036:  uint16(3),
	6037:  uint16(sym_bracket_index_expression),
	6038:  uint16(sym_dot_index_expression),
	6039:  uint16(sym__contextual_keyword),
	6040:  uint16(7),
	6041:  uint16(3),
	6042:  uint16(sym_variable),
	6043:  uint16(sym_function_call),
	6044:  uint16(sym_parenthesized_expression),
	6045:  uint16(40),
	6046:  uint16(5),
	6047:  uint16(sym_string),
	6048:  uint16(sym_function_definition),
	6049:  uint16(sym_table_constructor),
	6050:  uint16(sym_binary_expression),
	6051:  uint16(sym_unary_expression),
	6052:  uint16(23),
	6053:  uint16(3),
	6054:  uint16(1),
	6055:  uint16(anon_sym_DASH_DASH),
	6056:  uint16(5),
	6057:  uint16(1),
	6058:  uint16(sym__block_comment_start),
	6059:  uint16(9),
	6060:  uint16(1),
	6061:  uint16(sym_identifier),
	6062:  uint16(39),
	6063:  uint16(1),
	6064:  uint16(anon_sym_LPAREN),
	6065:  uint16(257),
	6066:  uint16(1),
	6067:  uint16(anon_sym_function),
	6068:  uint16(259),
	6069:  uint16(1),
	6070:  uint16(anon_sym_global),
	6071:  uint16(265),
	6072:  uint16(1),
	6073:  uint16(anon_sym_DQUOTE),
	6074:  uint16(267),
	6075:  uint16(1),
	6076:  uint16(anon_sym_SQUOTE),
	6077:  uint16(269),
	6078:  uint16(1),
	6079:  uint16(anon_sym_LBRACE),
	6080:  uint16(275),
	6081:  uint16(1),
	6082:  uint16(sym__block_string_start),
	6083:  uint16(41),
	6084:  uint16(1),
	6085:  uint16(sym__quote_string),
	6086:  uint16(45),
	6087:  uint16(1),
	6088:  uint16(sym__block_string),
	6089:  uint16(75),
	6090:  uint16(1),
	6091:  uint16(sym_comment),
	6092:  uint16(120),
	6093:  uint16(1),
	6094:  uint16(sym_expression),
	6095:  uint16(170),
	6096:  uint16(1),
	6097:  uint16(sym__prefix_expression),
	6098:  uint16(174),
	6099:  uint16(1),
	6100:  uint16(sym_method_index_expression),
	6101:  uint16(263),
	6102:  uint16(2),
	6103:  uint16(sym_number),
	6104:  uint16(sym_vararg_expression),
	6105:  uint16(271),
	6106:  uint16(2),
	6107:  uint16(anon_sym_TILDE),
	6108:  uint16(anon_sym_POUND),
	6109:  uint16(273),
	6110:  uint16(2),
	6111:  uint16(anon_sym_DASH),
	6112:  uint16(anon_sym_not),
	6113:  uint16(261),
	6114:  uint16(3),
	6115:  uint16(sym_nil),
	6116:  uint16(sym_false),
	6117:  uint16(sym_true),
	6118:  uint16(2),
	6119:  uint16(3),
	6120:  uint16(sym_bracket_index_expression),
	6121:  uint16(sym_dot_index_expression),
	6122:  uint16(sym__contextual_keyword),
	6123:  uint16(7),
	6124:  uint16(3),
	6125:  uint16(sym_variable),
	6126:  uint16(sym_function_call),
	6127:  uint16(sym_parenthesized_expression),
	6128:  uint16(40),
	6129:  uint16(5),
	6130:  uint16(sym_string),
	6131:  uint16(sym_function_definition),
	6132:  uint16(sym_table_constructor),
	6133:  uint16(sym_binary_expression),
	6134:  uint16(sym_unary_expression),
	6135:  uint16(6),
	6136:  uint16(3),
	6137:  uint16(1),
	6138:  uint16(anon_sym_DASH_DASH),
	6139:  uint16(5),
	6140:  uint16(1),
	6141:  uint16(sym__block_comment_start),
	6142:  uint16(293),
	6143:  uint16(1),
	6144:  uint16(anon_sym_EQ),
	6145:  uint16(76),
	6146:  uint16(1),
	6147:  uint16(sym_comment),
	6148:  uint16(43),
	6149:  uint16(6),
	6150:  uint16(anon_sym_DOT),
	6151:  uint16(anon_sym_LT),
	6152:  uint16(anon_sym_GT),
	6153:  uint16(anon_sym_TILDE),
	6154:  uint16(anon_sym_DASH),
	6155:  uint16(anon_sym_SLASH),
	6156:  uint16(41),
	6157:  uint16(26),
	6158:  uint16(sym__block_string_start),
	6159:  uint16(anon_sym_SEMI),
	6160:  uint16(anon_sym_COMMA),
	6161:  uint16(anon_sym_COLON),
	6162:  uint16(anon_sym_STAR),
	6163:  uint16(anon_sym_DQUOTE),
	6164:  uint16(anon_sym_SQUOTE),
	6165:  uint16(anon_sym_LPAREN),
	6166:  uint16(anon_sym_LBRACK),
	6167:  uint16(anon_sym_LBRACE),
	6168:  uint16(anon_sym_RBRACE),
	6169:  uint16(anon_sym_or),
	6170:  uint16(anon_sym_and),
	6171:  uint16(anon_sym_LT_EQ),
	6172:  uint16(anon_sym_EQ_EQ),
	6173:  uint16(anon_sym_TILDE_EQ),
	6174:  uint16(anon_sym_GT_EQ),
	6175:  uint16(anon_sym_PIPE),
	6176:  uint16(anon_sym_AMP),
	6177:  uint16(anon_sym_LT_LT),
	6178:  uint16(anon_sym_GT_GT),
	6179:  uint16(anon_sym_PLUS),
	6180:  uint16(anon_sym_SLASH_SLASH),
	6181:  uint16(anon_sym_PERCENT),
	6182:  uint16(anon_sym_DOT_DOT),
	6183:  uint16(anon_sym_CARET),
	6184:  uint16(23),
	6185:  uint16(3),
	6186:  uint16(1),
	6187:  uint16(anon_sym_DASH_DASH),
	6188:  uint16(5),
	6189:  uint16(1),
	6190:  uint16(sym__block_comment_start),
	6191:  uint16(9),
	6192:  uint16(1),
	6193:  uint16(sym_identifier),
	6194:  uint16(39),
	6195:  uint16(1),
	6196:  uint16(anon_sym_LPAREN),
	6197:  uint16(257),
	6198:  uint16(1),
	6199:  uint16(anon_sym_function),
	6200:  uint16(259),
	6201:  uint16(1),
	6202:  uint16(anon_sym_global),
	6203:  uint16(265),
	6204:  uint16(1),
	6205:  uint16(anon_sym_DQUOTE),
	6206:  uint16(267),
	6207:  uint16(1),
	6208:  uint16(anon_sym_SQUOTE),
	6209:  uint16(269),
	6210:  uint16(1),
	6211:  uint16(anon_sym_LBRACE),
	6212:  uint16(275),
	6213:  uint16(1),
	6214:  uint16(sym__block_string_start),
	6215:  uint16(41),
	6216:  uint16(1),
	6217:  uint16(sym__quote_string),
	6218:  uint16(45),
	6219:  uint16(1),
	6220:  uint16(sym__block_string),
	6221:  uint16(77),
	6222:  uint16(1),
	6223:  uint16(sym_comment),
	6224:  uint16(125),
	6225:  uint16(1),
	6226:  uint16(sym_expression),
	6227:  uint16(170),
	6228:  uint16(1),
	6229:  uint16(sym__prefix_expression),
	6230:  uint16(174),
	6231:  uint16(1),
	6232:  uint16(sym_method_index_expression),
	6233:  uint16(263),
	6234:  uint16(2),
	6235:  uint16(sym_number),
	6236:  uint16(sym_vararg_expression),
	6237:  uint16(271),
	6238:  uint16(2),
	6239:  uint16(anon_sym_TILDE),
	6240:  uint16(anon_sym_POUND),
	6241:  uint16(273),
	6242:  uint16(2),
	6243:  uint16(anon_sym_DASH),
	6244:  uint16(anon_sym_not),
	6245:  uint16(261),
	6246:  uint16(3),
	6247:  uint16(sym_nil),
	6248:  uint16(sym_false),
	6249:  uint16(sym_true),
	6250:  uint16(2),
	6251:  uint16(3),
	6252:  uint16(sym_bracket_index_expression),
	6253:  uint16(sym_dot_index_expression),
	6254:  uint16(sym__contextual_keyword),
	6255:  uint16(7),
	6256:  uint16(3),
	6257:  uint16(sym_variable),
	6258:  uint16(sym_function_call),
	6259:  uint16(sym_parenthesized_expression),
	6260:  uint16(40),
	6261:  uint16(5),
	6262:  uint16(sym_string),
	6263:  uint16(sym_function_definition),
	6264:  uint16(sym_table_constructor),
	6265:  uint16(sym_binary_expression),
	6266:  uint16(sym_unary_expression),
	6267:  uint16(23),
	6268:  uint16(3),
	6269:  uint16(1),
	6270:  uint16(anon_sym_DASH_DASH),
	6271:  uint16(5),
	6272:  uint16(1),
	6273:  uint16(sym__block_comment_start),
	6274:  uint16(9),
	6275:  uint16(1),
	6276:  uint16(sym_identifier),
	6277:  uint16(39),
	6278:  uint16(1),
	6279:  uint16(anon_sym_LPAREN),
	6280:  uint16(257),
	6281:  uint16(1),
	6282:  uint16(anon_sym_function),
	6283:  uint16(259),
	6284:  uint16(1),
	6285:  uint16(anon_sym_global),
	6286:  uint16(265),
	6287:  uint16(1),
	6288:  uint16(anon_sym_DQUOTE),
	6289:  uint16(267),
	6290:  uint16(1),
	6291:  uint16(anon_sym_SQUOTE),
	6292:  uint16(269),
	6293:  uint16(1),
	6294:  uint16(anon_sym_LBRACE),
	6295:  uint16(275),
	6296:  uint16(1),
	6297:  uint16(sym__block_string_start),
	6298:  uint16(41),
	6299:  uint16(1),
	6300:  uint16(sym__quote_string),
	6301:  uint16(45),
	6302:  uint16(1),
	6303:  uint16(sym__block_string),
	6304:  uint16(78),
	6305:  uint16(1),
	6306:  uint16(sym_comment),
	6307:  uint16(124),
	6308:  uint16(1),
	6309:  uint16(sym_expression),
	6310:  uint16(170),
	6311:  uint16(1),
	6312:  uint16(sym__prefix_expression),
	6313:  uint16(174),
	6314:  uint16(1),
	6315:  uint16(sym_method_index_expression),
	6316:  uint16(263),
	6317:  uint16(2),
	6318:  uint16(sym_number),
	6319:  uint16(sym_vararg_expression),
	6320:  uint16(271),
	6321:  uint16(2),
	6322:  uint16(anon_sym_TILDE),
	6323:  uint16(anon_sym_POUND),
	6324:  uint16(273),
	6325:  uint16(2),
	6326:  uint16(anon_sym_DASH),
	6327:  uint16(anon_sym_not),
	6328:  uint16(261),
	6329:  uint16(3),
	6330:  uint16(sym_nil),
	6331:  uint16(sym_false),
	6332:  uint16(sym_true),
	6333:  uint16(2),
	6334:  uint16(3),
	6335:  uint16(sym_bracket_index_expression),
	6336:  uint16(sym_dot_index_expression),
	6337:  uint16(sym__contextual_keyword),
	6338:  uint16(7),
	6339:  uint16(3),
	6340:  uint16(sym_variable),
	6341:  uint16(sym_function_call),
	6342:  uint16(sym_parenthesized_expression),
	6343:  uint16(40),
	6344:  uint16(5),
	6345:  uint16(sym_string),
	6346:  uint16(sym_function_definition),
	6347:  uint16(sym_table_constructor),
	6348:  uint16(sym_binary_expression),
	6349:  uint16(sym_unary_expression),
	6350:  uint16(23),
	6351:  uint16(3),
	6352:  uint16(1),
	6353:  uint16(anon_sym_DASH_DASH),
	6354:  uint16(5),
	6355:  uint16(1),
	6356:  uint16(sym__block_comment_start),
	6357:  uint16(9),
	6358:  uint16(1),
	6359:  uint16(sym_identifier),
	6360:  uint16(39),
	6361:  uint16(1),
	6362:  uint16(anon_sym_LPAREN),
	6363:  uint16(257),
	6364:  uint16(1),
	6365:  uint16(anon_sym_function),
	6366:  uint16(259),
	6367:  uint16(1),
	6368:  uint16(anon_sym_global),
	6369:  uint16(265),
	6370:  uint16(1),
	6371:  uint16(anon_sym_DQUOTE),
	6372:  uint16(267),
	6373:  uint16(1),
	6374:  uint16(anon_sym_SQUOTE),
	6375:  uint16(269),
	6376:  uint16(1),
	6377:  uint16(anon_sym_LBRACE),
	6378:  uint16(275),
	6379:  uint16(1),
	6380:  uint16(sym__block_string_start),
	6381:  uint16(41),
	6382:  uint16(1),
	6383:  uint16(sym__quote_string),
	6384:  uint16(45),
	6385:  uint16(1),
	6386:  uint16(sym__block_string),
	6387:  uint16(58),
	6388:  uint16(1),
	6389:  uint16(sym_expression),
	6390:  uint16(79),
	6391:  uint16(1),
	6392:  uint16(sym_comment),
	6393:  uint16(170),
	6394:  uint16(1),
	6395:  uint16(sym__prefix_expression),
	6396:  uint16(174),
	6397:  uint16(1),
	6398:  uint16(sym_method_index_expression),
	6399:  uint16(263),
	6400:  uint16(2),
	6401:  uint16(sym_number),
	6402:  uint16(sym_vararg_expression),
	6403:  uint16(271),
	6404:  uint16(2),
	6405:  uint16(anon_sym_TILDE),
	6406:  uint16(anon_sym_POUND),
	6407:  uint16(273),
	6408:  uint16(2),
	6409:  uint16(anon_sym_DASH),
	6410:  uint16(anon_sym_not),
	6411:  uint16(261),
	6412:  uint16(3),
	6413:  uint16(sym_nil),
	6414:  uint16(sym_false),
	6415:  uint16(sym_true),
	6416:  uint16(2),
	6417:  uint16(3),
	6418:  uint16(sym_bracket_index_expression),
	6419:  uint16(sym_dot_index_expression),
	6420:  uint16(sym__contextual_keyword),
	6421:  uint16(7),
	6422:  uint16(3),
	6423:  uint16(sym_variable),
	6424:  uint16(sym_function_call),
	6425:  uint16(sym_parenthesized_expression),
	6426:  uint16(40),
	6427:  uint16(5),
	6428:  uint16(sym_string),
	6429:  uint16(sym_function_definition),
	6430:  uint16(sym_table_constructor),
	6431:  uint16(sym_binary_expression),
	6432:  uint16(sym_unary_expression),
	6433:  uint16(23),
	6434:  uint16(3),
	6435:  uint16(1),
	6436:  uint16(anon_sym_DASH_DASH),
	6437:  uint16(5),
	6438:  uint16(1),
	6439:  uint16(sym__block_comment_start),
	6440:  uint16(9),
	6441:  uint16(1),
	6442:  uint16(sym_identifier),
	6443:  uint16(39),
	6444:  uint16(1),
	6445:  uint16(anon_sym_LPAREN),
	6446:  uint16(257),
	6447:  uint16(1),
	6448:  uint16(anon_sym_function),
	6449:  uint16(259),
	6450:  uint16(1),
	6451:  uint16(anon_sym_global),
	6452:  uint16(265),
	6453:  uint16(1),
	6454:  uint16(anon_sym_DQUOTE),
	6455:  uint16(267),
	6456:  uint16(1),
	6457:  uint16(anon_sym_SQUOTE),
	6458:  uint16(269),
	6459:  uint16(1),
	6460:  uint16(anon_sym_LBRACE),
	6461:  uint16(275),
	6462:  uint16(1),
	6463:  uint16(sym__block_string_start),
	6464:  uint16(41),
	6465:  uint16(1),
	6466:  uint16(sym__quote_string),
	6467:  uint16(45),
	6468:  uint16(1),
	6469:  uint16(sym__block_string),
	6470:  uint16(80),
	6471:  uint16(1),
	6472:  uint16(sym_comment),
	6473:  uint16(119),
	6474:  uint16(1),
	6475:  uint16(sym_expression),
	6476:  uint16(170),
	6477:  uint16(1),
	6478:  uint16(sym__prefix_expression),
	6479:  uint16(174),
	6480:  uint16(1),
	6481:  uint16(sym_method_index_expression),
	6482:  uint16(263),
	6483:  uint16(2),
	6484:  uint16(sym_number),
	6485:  uint16(sym_vararg_expression),
	6486:  uint16(271),
	6487:  uint16(2),
	6488:  uint16(anon_sym_TILDE),
	6489:  uint16(anon_sym_POUND),
	6490:  uint16(273),
	6491:  uint16(2),
	6492:  uint16(anon_sym_DASH),
	6493:  uint16(anon_sym_not),
	6494:  uint16(261),
	6495:  uint16(3),
	6496:  uint16(sym_nil),
	6497:  uint16(sym_false),
	6498:  uint16(sym_true),
	6499:  uint16(2),
	6500:  uint16(3),
	6501:  uint16(sym_bracket_index_expression),
	6502:  uint16(sym_dot_index_expression),
	6503:  uint16(sym__contextual_keyword),
	6504:  uint16(7),
	6505:  uint16(3),
	6506:  uint16(sym_variable),
	6507:  uint16(sym_function_call),
	6508:  uint16(sym_parenthesized_expression),
	6509:  uint16(40),
	6510:  uint16(5),
	6511:  uint16(sym_string),
	6512:  uint16(sym_function_definition),
	6513:  uint16(sym_table_constructor),
	6514:  uint16(sym_binary_expression),
	6515:  uint16(sym_unary_expression),
	6516:  uint16(23),
	6517:  uint16(3),
	6518:  uint16(1),
	6519:  uint16(anon_sym_DASH_DASH),
	6520:  uint16(5),
	6521:  uint16(1),
	6522:  uint16(sym__block_comment_start),
	6523:  uint16(9),
	6524:  uint16(1),
	6525:  uint16(sym_identifier),
	6526:  uint16(39),
	6527:  uint16(1),
	6528:  uint16(anon_sym_LPAREN),
	6529:  uint16(257),
	6530:  uint16(1),
	6531:  uint16(anon_sym_function),
	6532:  uint16(259),
	6533:  uint16(1),
	6534:  uint16(anon_sym_global),
	6535:  uint16(265),
	6536:  uint16(1),
	6537:  uint16(anon_sym_DQUOTE),
	6538:  uint16(267),
	6539:  uint16(1),
	6540:  uint16(anon_sym_SQUOTE),
	6541:  uint16(269),
	6542:  uint16(1),
	6543:  uint16(anon_sym_LBRACE),
	6544:  uint16(275),
	6545:  uint16(1),
	6546:  uint16(sym__block_string_start),
	6547:  uint16(28),
	6548:  uint16(1),
	6549:  uint16(sym_expression),
	6550:  uint16(41),
	6551:  uint16(1),
	6552:  uint16(sym__quote_string),
	6553:  uint16(45),
	6554:  uint16(1),
	6555:  uint16(sym__block_string),
	6556:  uint16(81),
	6557:  uint16(1),
	6558:  uint16(sym_comment),
	6559:  uint16(170),
	6560:  uint16(1),
	6561:  uint16(sym__prefix_expression),
	6562:  uint16(174),
	6563:  uint16(1),
	6564:  uint16(sym_method_index_expression),
	6565:  uint16(263),
	6566:  uint16(2),
	6567:  uint16(sym_number),
	6568:  uint16(sym_vararg_expression),
	6569:  uint16(271),
	6570:  uint16(2),
	6571:  uint16(anon_sym_TILDE),
	6572:  uint16(anon_sym_POUND),
	6573:  uint16(273),
	6574:  uint16(2),
	6575:  uint16(anon_sym_DASH),
	6576:  uint16(anon_sym_not),
	6577:  uint16(261),
	6578:  uint16(3),
	6579:  uint16(sym_nil),
	6580:  uint16(sym_false),
	6581:  uint16(sym_true),
	6582:  uint16(2),
	6583:  uint16(3),
	6584:  uint16(sym_bracket_index_expression),
	6585:  uint16(sym_dot_index_expression),
	6586:  uint16(sym__contextual_keyword),
	6587:  uint16(7),
	6588:  uint16(3),
	6589:  uint16(sym_variable),
	6590:  uint16(sym_function_call),
	6591:  uint16(sym_parenthesized_expression),
	6592:  uint16(40),
	6593:  uint16(5),
	6594:  uint16(sym_string),
	6595:  uint16(sym_function_definition),
	6596:  uint16(sym_table_constructor),
	6597:  uint16(sym_binary_expression),
	6598:  uint16(sym_unary_expression),
	6599:  uint16(23),
	6600:  uint16(3),
	6601:  uint16(1),
	6602:  uint16(anon_sym_DASH_DASH),
	6603:  uint16(5),
	6604:  uint16(1),
	6605:  uint16(sym__block_comment_start),
	6606:  uint16(9),
	6607:  uint16(1),
	6608:  uint16(sym_identifier),
	6609:  uint16(39),
	6610:  uint16(1),
	6611:  uint16(anon_sym_LPAREN),
	6612:  uint16(257),
	6613:  uint16(1),
	6614:  uint16(anon_sym_function),
	6615:  uint16(259),
	6616:  uint16(1),
	6617:  uint16(anon_sym_global),
	6618:  uint16(265),
	6619:  uint16(1),
	6620:  uint16(anon_sym_DQUOTE),
	6621:  uint16(267),
	6622:  uint16(1),
	6623:  uint16(anon_sym_SQUOTE),
	6624:  uint16(269),
	6625:  uint16(1),
	6626:  uint16(anon_sym_LBRACE),
	6627:  uint16(275),
	6628:  uint16(1),
	6629:  uint16(sym__block_string_start),
	6630:  uint16(41),
	6631:  uint16(1),
	6632:  uint16(sym__quote_string),
	6633:  uint16(45),
	6634:  uint16(1),
	6635:  uint16(sym__block_string),
	6636:  uint16(57),
	6637:  uint16(1),
	6638:  uint16(sym_expression),
	6639:  uint16(82),
	6640:  uint16(1),
	6641:  uint16(sym_comment),
	6642:  uint16(170),
	6643:  uint16(1),
	6644:  uint16(sym__prefix_expression),
	6645:  uint16(174),
	6646:  uint16(1),
	6647:  uint16(sym_method_index_expression),
	6648:  uint16(263),
	6649:  uint16(2),
	6650:  uint16(sym_number),
	6651:  uint16(sym_vararg_expression),
	6652:  uint16(271),
	6653:  uint16(2),
	6654:  uint16(anon_sym_TILDE),
	6655:  uint16(anon_sym_POUND),
	6656:  uint16(273),
	6657:  uint16(2),
	6658:  uint16(anon_sym_DASH),
	6659:  uint16(anon_sym_not),
	6660:  uint16(261),
	6661:  uint16(3),
	6662:  uint16(sym_nil),
	6663:  uint16(sym_false),
	6664:  uint16(sym_true),
	6665:  uint16(2),
	6666:  uint16(3),
	6667:  uint16(sym_bracket_index_expression),
	6668:  uint16(sym_dot_index_expression),
	6669:  uint16(sym__contextual_keyword),
	6670:  uint16(7),
	6671:  uint16(3),
	6672:  uint16(sym_variable),
	6673:  uint16(sym_function_call),
	6674:  uint16(sym_parenthesized_expression),
	6675:  uint16(40),
	6676:  uint16(5),
	6677:  uint16(sym_string),
	6678:  uint16(sym_function_definition),
	6679:  uint16(sym_table_constructor),
	6680:  uint16(sym_binary_expression),
	6681:  uint16(sym_unary_expression),
	6682:  uint16(23),
	6683:  uint16(3),
	6684:  uint16(1),
	6685:  uint16(anon_sym_DASH_DASH),
	6686:  uint16(5),
	6687:  uint16(1),
	6688:  uint16(sym__block_comment_start),
	6689:  uint16(9),
	6690:  uint16(1),
	6691:  uint16(sym_identifier),
	6692:  uint16(39),
	6693:  uint16(1),
	6694:  uint16(anon_sym_LPAREN),
	6695:  uint16(257),
	6696:  uint16(1),
	6697:  uint16(anon_sym_function),
	6698:  uint16(259),
	6699:  uint16(1),
	6700:  uint16(anon_sym_global),
	6701:  uint16(265),
	6702:  uint16(1),
	6703:  uint16(anon_sym_DQUOTE),
	6704:  uint16(267),
	6705:  uint16(1),
	6706:  uint16(anon_sym_SQUOTE),
	6707:  uint16(269),
	6708:  uint16(1),
	6709:  uint16(anon_sym_LBRACE),
	6710:  uint16(275),
	6711:  uint16(1),
	6712:  uint16(sym__block_string_start),
	6713:  uint16(32),
	6714:  uint16(1),
	6715:  uint16(sym_expression),
	6716:  uint16(41),
	6717:  uint16(1),
	6718:  uint16(sym__quote_string),
	6719:  uint16(45),
	6720:  uint16(1),
	6721:  uint16(sym__block_string),
	6722:  uint16(83),
	6723:  uint16(1),
	6724:  uint16(sym_comment),
	6725:  uint16(170),
	6726:  uint16(1),
	6727:  uint16(sym__prefix_expression),
	6728:  uint16(174),
	6729:  uint16(1),
	6730:  uint16(sym_method_index_expression),
	6731:  uint16(263),
	6732:  uint16(2),
	6733:  uint16(sym_number),
	6734:  uint16(sym_vararg_expression),
	6735:  uint16(271),
	6736:  uint16(2),
	6737:  uint16(anon_sym_TILDE),
	6738:  uint16(anon_sym_POUND),
	6739:  uint16(273),
	6740:  uint16(2),
	6741:  uint16(anon_sym_DASH),
	6742:  uint16(anon_sym_not),
	6743:  uint16(261),
	6744:  uint16(3),
	6745:  uint16(sym_nil),
	6746:  uint16(sym_false),
	6747:  uint16(sym_true),
	6748:  uint16(2),
	6749:  uint16(3),
	6750:  uint16(sym_bracket_index_expression),
	6751:  uint16(sym_dot_index_expression),
	6752:  uint16(sym__contextual_keyword),
	6753:  uint16(7),
	6754:  uint16(3),
	6755:  uint16(sym_variable),
	6756:  uint16(sym_function_call),
	6757:  uint16(sym_parenthesized_expression),
	6758:  uint16(40),
	6759:  uint16(5),
	6760:  uint16(sym_string),
	6761:  uint16(sym_function_definition),
	6762:  uint16(sym_table_constructor),
	6763:  uint16(sym_binary_expression),
	6764:  uint16(sym_unary_expression),
	6765:  uint16(23),
	6766:  uint16(3),
	6767:  uint16(1),
	6768:  uint16(anon_sym_DASH_DASH),
	6769:  uint16(5),
	6770:  uint16(1),
	6771:  uint16(sym__block_comment_start),
	6772:  uint16(9),
	6773:  uint16(1),
	6774:  uint16(sym_identifier),
	6775:  uint16(39),
	6776:  uint16(1),
	6777:  uint16(anon_sym_LPAREN),
	6778:  uint16(257),
	6779:  uint16(1),
	6780:  uint16(anon_sym_function),
	6781:  uint16(259),
	6782:  uint16(1),
	6783:  uint16(anon_sym_global),
	6784:  uint16(265),
	6785:  uint16(1),
	6786:  uint16(anon_sym_DQUOTE),
	6787:  uint16(267),
	6788:  uint16(1),
	6789:  uint16(anon_sym_SQUOTE),
	6790:  uint16(269),
	6791:  uint16(1),
	6792:  uint16(anon_sym_LBRACE),
	6793:  uint16(275),
	6794:  uint16(1),
	6795:  uint16(sym__block_string_start),
	6796:  uint16(26),
	6797:  uint16(1),
	6798:  uint16(sym_expression),
	6799:  uint16(41),
	6800:  uint16(1),
	6801:  uint16(sym__quote_string),
	6802:  uint16(45),
	6803:  uint16(1),
	6804:  uint16(sym__block_string),
	6805:  uint16(84),
	6806:  uint16(1),
	6807:  uint16(sym_comment),
	6808:  uint16(170),
	6809:  uint16(1),
	6810:  uint16(sym__prefix_expression),
	6811:  uint16(174),
	6812:  uint16(1),
	6813:  uint16(sym_method_index_expression),
	6814:  uint16(263),
	6815:  uint16(2),
	6816:  uint16(sym_number),
	6817:  uint16(sym_vararg_expression),
	6818:  uint16(271),
	6819:  uint16(2),
	6820:  uint16(anon_sym_TILDE),
	6821:  uint16(anon_sym_POUND),
	6822:  uint16(273),
	6823:  uint16(2),
	6824:  uint16(anon_sym_DASH),
	6825:  uint16(anon_sym_not),
	6826:  uint16(261),
	6827:  uint16(3),
	6828:  uint16(sym_nil),
	6829:  uint16(sym_false),
	6830:  uint16(sym_true),
	6831:  uint16(2),
	6832:  uint16(3),
	6833:  uint16(sym_bracket_index_expression),
	6834:  uint16(sym_dot_index_expression),
	6835:  uint16(sym__contextual_keyword),
	6836:  uint16(7),
	6837:  uint16(3),
	6838:  uint16(sym_variable),
	6839:  uint16(sym_function_call),
	6840:  uint16(sym_parenthesized_expression),
	6841:  uint16(40),
	6842:  uint16(5),
	6843:  uint16(sym_string),
	6844:  uint16(sym_function_definition),
	6845:  uint16(sym_table_constructor),
	6846:  uint16(sym_binary_expression),
	6847:  uint16(sym_unary_expression),
	6848:  uint16(23),
	6849:  uint16(3),
	6850:  uint16(1),
	6851:  uint16(anon_sym_DASH_DASH),
	6852:  uint16(5),
	6853:  uint16(1),
	6854:  uint16(sym__block_comment_start),
	6855:  uint16(9),
	6856:  uint16(1),
	6857:  uint16(sym_identifier),
	6858:  uint16(39),
	6859:  uint16(1),
	6860:  uint16(anon_sym_LPAREN),
	6861:  uint16(257),
	6862:  uint16(1),
	6863:  uint16(anon_sym_function),
	6864:  uint16(259),
	6865:  uint16(1),
	6866:  uint16(anon_sym_global),
	6867:  uint16(265),
	6868:  uint16(1),
	6869:  uint16(anon_sym_DQUOTE),
	6870:  uint16(267),
	6871:  uint16(1),
	6872:  uint16(anon_sym_SQUOTE),
	6873:  uint16(269),
	6874:  uint16(1),
	6875:  uint16(anon_sym_LBRACE),
	6876:  uint16(275),
	6877:  uint16(1),
	6878:  uint16(sym__block_string_start),
	6879:  uint16(41),
	6880:  uint16(1),
	6881:  uint16(sym__quote_string),
	6882:  uint16(45),
	6883:  uint16(1),
	6884:  uint16(sym__block_string),
	6885:  uint16(85),
	6886:  uint16(1),
	6887:  uint16(sym_comment),
	6888:  uint16(123),
	6889:  uint16(1),
	6890:  uint16(sym_expression),
	6891:  uint16(170),
	6892:  uint16(1),
	6893:  uint16(sym__prefix_expression),
	6894:  uint16(174),
	6895:  uint16(1),
	6896:  uint16(sym_method_index_expression),
	6897:  uint16(263),
	6898:  uint16(2),
	6899:  uint16(sym_number),
	6900:  uint16(sym_vararg_expression),
	6901:  uint16(271),
	6902:  uint16(2),
	6903:  uint16(anon_sym_TILDE),
	6904:  uint16(anon_sym_POUND),
	6905:  uint16(273),
	6906:  uint16(2),
	6907:  uint16(anon_sym_DASH),
	6908:  uint16(anon_sym_not),
	6909:  uint16(261),
	6910:  uint16(3),
	6911:  uint16(sym_nil),
	6912:  uint16(sym_false),
	6913:  uint16(sym_true),
	6914:  uint16(2),
	6915:  uint16(3),
	6916:  uint16(sym_bracket_index_expression),
	6917:  uint16(sym_dot_index_expression),
	6918:  uint16(sym__contextual_keyword),
	6919:  uint16(7),
	6920:  uint16(3),
	6921:  uint16(sym_variable),
	6922:  uint16(sym_function_call),
	6923:  uint16(sym_parenthesized_expression),
	6924:  uint16(40),
	6925:  uint16(5),
	6926:  uint16(sym_string),
	6927:  uint16(sym_function_definition),
	6928:  uint16(sym_table_constructor),
	6929:  uint16(sym_binary_expression),
	6930:  uint16(sym_unary_expression),
	6931:  uint16(23),
	6932:  uint16(3),
	6933:  uint16(1),
	6934:  uint16(anon_sym_DASH_DASH),
	6935:  uint16(5),
	6936:  uint16(1),
	6937:  uint16(sym__block_comment_start),
	6938:  uint16(9),
	6939:  uint16(1),
	6940:  uint16(sym_identifier),
	6941:  uint16(39),
	6942:  uint16(1),
	6943:  uint16(anon_sym_LPAREN),
	6944:  uint16(257),
	6945:  uint16(1),
	6946:  uint16(anon_sym_function),
	6947:  uint16(259),
	6948:  uint16(1),
	6949:  uint16(anon_sym_global),
	6950:  uint16(265),
	6951:  uint16(1),
	6952:  uint16(anon_sym_DQUOTE),
	6953:  uint16(267),
	6954:  uint16(1),
	6955:  uint16(anon_sym_SQUOTE),
	6956:  uint16(269),
	6957:  uint16(1),
	6958:  uint16(anon_sym_LBRACE),
	6959:  uint16(275),
	6960:  uint16(1),
	6961:  uint16(sym__block_string_start),
	6962:  uint16(41),
	6963:  uint16(1),
	6964:  uint16(sym__quote_string),
	6965:  uint16(45),
	6966:  uint16(1),
	6967:  uint16(sym__block_string),
	6968:  uint16(86),
	6969:  uint16(1),
	6970:  uint16(sym_comment),
	6971:  uint16(115),
	6972:  uint16(1),
	6973:  uint16(sym_expression),
	6974:  uint16(170),
	6975:  uint16(1),
	6976:  uint16(sym__prefix_expression),
	6977:  uint16(174),
	6978:  uint16(1),
	6979:  uint16(sym_method_index_expression),
	6980:  uint16(263),
	6981:  uint16(2),
	6982:  uint16(sym_number),
	6983:  uint16(sym_vararg_expression),
	6984:  uint16(271),
	6985:  uint16(2),
	6986:  uint16(anon_sym_TILDE),
	6987:  uint16(anon_sym_POUND),
	6988:  uint16(273),
	6989:  uint16(2),
	6990:  uint16(anon_sym_DASH),
	6991:  uint16(anon_sym_not),
	6992:  uint16(261),
	6993:  uint16(3),
	6994:  uint16(sym_nil),
	6995:  uint16(sym_false),
	6996:  uint16(sym_true),
	6997:  uint16(2),
	6998:  uint16(3),
	6999:  uint16(sym_bracket_index_expression),
	7000:  uint16(sym_dot_index_expression),
	7001:  uint16(sym__contextual_keyword),
	7002:  uint16(7),
	7003:  uint16(3),
	7004:  uint16(sym_variable),
	7005:  uint16(sym_function_call),
	7006:  uint16(sym_parenthesized_expression),
	7007:  uint16(40),
	7008:  uint16(5),
	7009:  uint16(sym_string),
	7010:  uint16(sym_function_definition),
	7011:  uint16(sym_table_constructor),
	7012:  uint16(sym_binary_expression),
	7013:  uint16(sym_unary_expression),
	7014:  uint16(23),
	7015:  uint16(3),
	7016:  uint16(1),
	7017:  uint16(anon_sym_DASH_DASH),
	7018:  uint16(5),
	7019:  uint16(1),
	7020:  uint16(sym__block_comment_start),
	7021:  uint16(9),
	7022:  uint16(1),
	7023:  uint16(sym_identifier),
	7024:  uint16(39),
	7025:  uint16(1),
	7026:  uint16(anon_sym_LPAREN),
	7027:  uint16(257),
	7028:  uint16(1),
	7029:  uint16(anon_sym_function),
	7030:  uint16(259),
	7031:  uint16(1),
	7032:  uint16(anon_sym_global),
	7033:  uint16(265),
	7034:  uint16(1),
	7035:  uint16(anon_sym_DQUOTE),
	7036:  uint16(267),
	7037:  uint16(1),
	7038:  uint16(anon_sym_SQUOTE),
	7039:  uint16(269),
	7040:  uint16(1),
	7041:  uint16(anon_sym_LBRACE),
	7042:  uint16(275),
	7043:  uint16(1),
	7044:  uint16(sym__block_string_start),
	7045:  uint16(33),
	7046:  uint16(1),
	7047:  uint16(sym_expression),
	7048:  uint16(41),
	7049:  uint16(1),
	7050:  uint16(sym__quote_string),
	7051:  uint16(45),
	7052:  uint16(1),
	7053:  uint16(sym__block_string),
	7054:  uint16(87),
	7055:  uint16(1),
	7056:  uint16(sym_comment),
	7057:  uint16(170),
	7058:  uint16(1),
	7059:  uint16(sym__prefix_expression),
	7060:  uint16(174),
	7061:  uint16(1),
	7062:  uint16(sym_method_index_expression),
	7063:  uint16(263),
	7064:  uint16(2),
	7065:  uint16(sym_number),
	7066:  uint16(sym_vararg_expression),
	7067:  uint16(271),
	7068:  uint16(2),
	7069:  uint16(anon_sym_TILDE),
	7070:  uint16(anon_sym_POUND),
	7071:  uint16(273),
	7072:  uint16(2),
	7073:  uint16(anon_sym_DASH),
	7074:  uint16(anon_sym_not),
	7075:  uint16(261),
	7076:  uint16(3),
	7077:  uint16(sym_nil),
	7078:  uint16(sym_false),
	7079:  uint16(sym_true),
	7080:  uint16(2),
	7081:  uint16(3),
	7082:  uint16(sym_bracket_index_expression),
	7083:  uint16(sym_dot_index_expression),
	7084:  uint16(sym__contextual_keyword),
	7085:  uint16(7),
	7086:  uint16(3),
	7087:  uint16(sym_variable),
	7088:  uint16(sym_function_call),
	7089:  uint16(sym_parenthesized_expression),
	7090:  uint16(40),
	7091:  uint16(5),
	7092:  uint16(sym_string),
	7093:  uint16(sym_function_definition),
	7094:  uint16(sym_table_constructor),
	7095:  uint16(sym_binary_expression),
	7096:  uint16(sym_unary_expression),
	7097:  uint16(23),
	7098:  uint16(3),
	7099:  uint16(1),
	7100:  uint16(anon_sym_DASH_DASH),
	7101:  uint16(5),
	7102:  uint16(1),
	7103:  uint16(sym__block_comment_start),
	7104:  uint16(9),
	7105:  uint16(1),
	7106:  uint16(sym_identifier),
	7107:  uint16(39),
	7108:  uint16(1),
	7109:  uint16(anon_sym_LPAREN),
	7110:  uint16(257),
	7111:  uint16(1),
	7112:  uint16(anon_sym_function),
	7113:  uint16(259),
	7114:  uint16(1),
	7115:  uint16(anon_sym_global),
	7116:  uint16(265),
	7117:  uint16(1),
	7118:  uint16(anon_sym_DQUOTE),
	7119:  uint16(267),
	7120:  uint16(1),
	7121:  uint16(anon_sym_SQUOTE),
	7122:  uint16(269),
	7123:  uint16(1),
	7124:  uint16(anon_sym_LBRACE),
	7125:  uint16(275),
	7126:  uint16(1),
	7127:  uint16(sym__block_string_start),
	7128:  uint16(25),
	7129:  uint16(1),
	7130:  uint16(sym_expression),
	7131:  uint16(41),
	7132:  uint16(1),
	7133:  uint16(sym__quote_string),
	7134:  uint16(45),
	7135:  uint16(1),
	7136:  uint16(sym__block_string),
	7137:  uint16(88),
	7138:  uint16(1),
	7139:  uint16(sym_comment),
	7140:  uint16(170),
	7141:  uint16(1),
	7142:  uint16(sym__prefix_expression),
	7143:  uint16(174),
	7144:  uint16(1),
	7145:  uint16(sym_method_index_expression),
	7146:  uint16(263),
	7147:  uint16(2),
	7148:  uint16(sym_number),
	7149:  uint16(sym_vararg_expression),
	7150:  uint16(271),
	7151:  uint16(2),
	7152:  uint16(anon_sym_TILDE),
	7153:  uint16(anon_sym_POUND),
	7154:  uint16(273),
	7155:  uint16(2),
	7156:  uint16(anon_sym_DASH),
	7157:  uint16(anon_sym_not),
	7158:  uint16(261),
	7159:  uint16(3),
	7160:  uint16(sym_nil),
	7161:  uint16(sym_false),
	7162:  uint16(sym_true),
	7163:  uint16(2),
	7164:  uint16(3),
	7165:  uint16(sym_bracket_index_expression),
	7166:  uint16(sym_dot_index_expression),
	7167:  uint16(sym__contextual_keyword),
	7168:  uint16(7),
	7169:  uint16(3),
	7170:  uint16(sym_variable),
	7171:  uint16(sym_function_call),
	7172:  uint16(sym_parenthesized_expression),
	7173:  uint16(40),
	7174:  uint16(5),
	7175:  uint16(sym_string),
	7176:  uint16(sym_function_definition),
	7177:  uint16(sym_table_constructor),
	7178:  uint16(sym_binary_expression),
	7179:  uint16(sym_unary_expression),
	7180:  uint16(23),
	7181:  uint16(3),
	7182:  uint16(1),
	7183:  uint16(anon_sym_DASH_DASH),
	7184:  uint16(5),
	7185:  uint16(1),
	7186:  uint16(sym__block_comment_start),
	7187:  uint16(9),
	7188:  uint16(1),
	7189:  uint16(sym_identifier),
	7190:  uint16(39),
	7191:  uint16(1),
	7192:  uint16(anon_sym_LPAREN),
	7193:  uint16(257),
	7194:  uint16(1),
	7195:  uint16(anon_sym_function),
	7196:  uint16(259),
	7197:  uint16(1),
	7198:  uint16(anon_sym_global),
	7199:  uint16(265),
	7200:  uint16(1),
	7201:  uint16(anon_sym_DQUOTE),
	7202:  uint16(267),
	7203:  uint16(1),
	7204:  uint16(anon_sym_SQUOTE),
	7205:  uint16(269),
	7206:  uint16(1),
	7207:  uint16(anon_sym_LBRACE),
	7208:  uint16(275),
	7209:  uint16(1),
	7210:  uint16(sym__block_string_start),
	7211:  uint16(41),
	7212:  uint16(1),
	7213:  uint16(sym__quote_string),
	7214:  uint16(45),
	7215:  uint16(1),
	7216:  uint16(sym__block_string),
	7217:  uint16(89),
	7218:  uint16(1),
	7219:  uint16(sym_comment),
	7220:  uint16(108),
	7221:  uint16(1),
	7222:  uint16(sym_expression),
	7223:  uint16(170),
	7224:  uint16(1),
	7225:  uint16(sym__prefix_expression),
	7226:  uint16(174),
	7227:  uint16(1),
	7228:  uint16(sym_method_index_expression),
	7229:  uint16(263),
	7230:  uint16(2),
	7231:  uint16(sym_number),
	7232:  uint16(sym_vararg_expression),
	7233:  uint16(271),
	7234:  uint16(2),
	7235:  uint16(anon_sym_TILDE),
	7236:  uint16(anon_sym_POUND),
	7237:  uint16(273),
	7238:  uint16(2),
	7239:  uint16(anon_sym_DASH),
	7240:  uint16(anon_sym_not),
	7241:  uint16(261),
	7242:  uint16(3),
	7243:  uint16(sym_nil),
	7244:  uint16(sym_false),
	7245:  uint16(sym_true),
	7246:  uint16(2),
	7247:  uint16(3),
	7248:  uint16(sym_bracket_index_expression),
	7249:  uint16(sym_dot_index_expression),
	7250:  uint16(sym__contextual_keyword),
	7251:  uint16(7),
	7252:  uint16(3),
	7253:  uint16(sym_variable),
	7254:  uint16(sym_function_call),
	7255:  uint16(sym_parenthesized_expression),
	7256:  uint16(40),
	7257:  uint16(5),
	7258:  uint16(sym_string),
	7259:  uint16(sym_function_definition),
	7260:  uint16(sym_table_constructor),
	7261:  uint16(sym_binary_expression),
	7262:  uint16(sym_unary_expression),
	7263:  uint16(23),
	7264:  uint16(3),
	7265:  uint16(1),
	7266:  uint16(anon_sym_DASH_DASH),
	7267:  uint16(5),
	7268:  uint16(1),
	7269:  uint16(sym__block_comment_start),
	7270:  uint16(9),
	7271:  uint16(1),
	7272:  uint16(sym_identifier),
	7273:  uint16(39),
	7274:  uint16(1),
	7275:  uint16(anon_sym_LPAREN),
	7276:  uint16(257),
	7277:  uint16(1),
	7278:  uint16(anon_sym_function),
	7279:  uint16(259),
	7280:  uint16(1),
	7281:  uint16(anon_sym_global),
	7282:  uint16(265),
	7283:  uint16(1),
	7284:  uint16(anon_sym_DQUOTE),
	7285:  uint16(267),
	7286:  uint16(1),
	7287:  uint16(anon_sym_SQUOTE),
	7288:  uint16(269),
	7289:  uint16(1),
	7290:  uint16(anon_sym_LBRACE),
	7291:  uint16(275),
	7292:  uint16(1),
	7293:  uint16(sym__block_string_start),
	7294:  uint16(38),
	7295:  uint16(1),
	7296:  uint16(sym_expression),
	7297:  uint16(41),
	7298:  uint16(1),
	7299:  uint16(sym__quote_string),
	7300:  uint16(45),
	7301:  uint16(1),
	7302:  uint16(sym__block_string),
	7303:  uint16(90),
	7304:  uint16(1),
	7305:  uint16(sym_comment),
	7306:  uint16(170),
	7307:  uint16(1),
	7308:  uint16(sym__prefix_expression),
	7309:  uint16(174),
	7310:  uint16(1),
	7311:  uint16(sym_method_index_expression),
	7312:  uint16(263),
	7313:  uint16(2),
	7314:  uint16(sym_number),
	7315:  uint16(sym_vararg_expression),
	7316:  uint16(271),
	7317:  uint16(2),
	7318:  uint16(anon_sym_TILDE),
	7319:  uint16(anon_sym_POUND),
	7320:  uint16(273),
	7321:  uint16(2),
	7322:  uint16(anon_sym_DASH),
	7323:  uint16(anon_sym_not),
	7324:  uint16(261),
	7325:  uint16(3),
	7326:  uint16(sym_nil),
	7327:  uint16(sym_false),
	7328:  uint16(sym_true),
	7329:  uint16(2),
	7330:  uint16(3),
	7331:  uint16(sym_bracket_index_expression),
	7332:  uint16(sym_dot_index_expression),
	7333:  uint16(sym__contextual_keyword),
	7334:  uint16(7),
	7335:  uint16(3),
	7336:  uint16(sym_variable),
	7337:  uint16(sym_function_call),
	7338:  uint16(sym_parenthesized_expression),
	7339:  uint16(40),
	7340:  uint16(5),
	7341:  uint16(sym_string),
	7342:  uint16(sym_function_definition),
	7343:  uint16(sym_table_constructor),
	7344:  uint16(sym_binary_expression),
	7345:  uint16(sym_unary_expression),
	7346:  uint16(23),
	7347:  uint16(3),
	7348:  uint16(1),
	7349:  uint16(anon_sym_DASH_DASH),
	7350:  uint16(5),
	7351:  uint16(1),
	7352:  uint16(sym__block_comment_start),
	7353:  uint16(9),
	7354:  uint16(1),
	7355:  uint16(sym_identifier),
	7356:  uint16(39),
	7357:  uint16(1),
	7358:  uint16(anon_sym_LPAREN),
	7359:  uint16(257),
	7360:  uint16(1),
	7361:  uint16(anon_sym_function),
	7362:  uint16(259),
	7363:  uint16(1),
	7364:  uint16(anon_sym_global),
	7365:  uint16(265),
	7366:  uint16(1),
	7367:  uint16(anon_sym_DQUOTE),
	7368:  uint16(267),
	7369:  uint16(1),
	7370:  uint16(anon_sym_SQUOTE),
	7371:  uint16(269),
	7372:  uint16(1),
	7373:  uint16(anon_sym_LBRACE),
	7374:  uint16(275),
	7375:  uint16(1),
	7376:  uint16(sym__block_string_start),
	7377:  uint16(34),
	7378:  uint16(1),
	7379:  uint16(sym_expression),
	7380:  uint16(41),
	7381:  uint16(1),
	7382:  uint16(sym__quote_string),
	7383:  uint16(45),
	7384:  uint16(1),
	7385:  uint16(sym__block_string),
	7386:  uint16(91),
	7387:  uint16(1),
	7388:  uint16(sym_comment),
	7389:  uint16(170),
	7390:  uint16(1),
	7391:  uint16(sym__prefix_expression),
	7392:  uint16(174),
	7393:  uint16(1),
	7394:  uint16(sym_method_index_expression),
	7395:  uint16(263),
	7396:  uint16(2),
	7397:  uint16(sym_number),
	7398:  uint16(sym_vararg_expression),
	7399:  uint16(271),
	7400:  uint16(2),
	7401:  uint16(anon_sym_TILDE),
	7402:  uint16(anon_sym_POUND),
	7403:  uint16(273),
	7404:  uint16(2),
	7405:  uint16(anon_sym_DASH),
	7406:  uint16(anon_sym_not),
	7407:  uint16(261),
	7408:  uint16(3),
	7409:  uint16(sym_nil),
	7410:  uint16(sym_false),
	7411:  uint16(sym_true),
	7412:  uint16(2),
	7413:  uint16(3),
	7414:  uint16(sym_bracket_index_expression),
	7415:  uint16(sym_dot_index_expression),
	7416:  uint16(sym__contextual_keyword),
	7417:  uint16(7),
	7418:  uint16(3),
	7419:  uint16(sym_variable),
	7420:  uint16(sym_function_call),
	7421:  uint16(sym_parenthesized_expression),
	7422:  uint16(40),
	7423:  uint16(5),
	7424:  uint16(sym_string),
	7425:  uint16(sym_function_definition),
	7426:  uint16(sym_table_constructor),
	7427:  uint16(sym_binary_expression),
	7428:  uint16(sym_unary_expression),
	7429:  uint16(23),
	7430:  uint16(3),
	7431:  uint16(1),
	7432:  uint16(anon_sym_DASH_DASH),
	7433:  uint16(5),
	7434:  uint16(1),
	7435:  uint16(sym__block_comment_start),
	7436:  uint16(9),
	7437:  uint16(1),
	7438:  uint16(sym_identifier),
	7439:  uint16(39),
	7440:  uint16(1),
	7441:  uint16(anon_sym_LPAREN),
	7442:  uint16(257),
	7443:  uint16(1),
	7444:  uint16(anon_sym_function),
	7445:  uint16(259),
	7446:  uint16(1),
	7447:  uint16(anon_sym_global),
	7448:  uint16(265),
	7449:  uint16(1),
	7450:  uint16(anon_sym_DQUOTE),
	7451:  uint16(267),
	7452:  uint16(1),
	7453:  uint16(anon_sym_SQUOTE),
	7454:  uint16(269),
	7455:  uint16(1),
	7456:  uint16(anon_sym_LBRACE),
	7457:  uint16(275),
	7458:  uint16(1),
	7459:  uint16(sym__block_string_start),
	7460:  uint16(41),
	7461:  uint16(1),
	7462:  uint16(sym__quote_string),
	7463:  uint16(45),
	7464:  uint16(1),
	7465:  uint16(sym__block_string),
	7466:  uint16(92),
	7467:  uint16(1),
	7468:  uint16(sym_comment),
	7469:  uint16(126),
	7470:  uint16(1),
	7471:  uint16(sym_expression),
	7472:  uint16(170),
	7473:  uint16(1),
	7474:  uint16(sym__prefix_expression),
	7475:  uint16(174),
	7476:  uint16(1),
	7477:  uint16(sym_method_index_expression),
	7478:  uint16(263),
	7479:  uint16(2),
	7480:  uint16(sym_number),
	7481:  uint16(sym_vararg_expression),
	7482:  uint16(271),
	7483:  uint16(2),
	7484:  uint16(anon_sym_TILDE),
	7485:  uint16(anon_sym_POUND),
	7486:  uint16(273),
	7487:  uint16(2),
	7488:  uint16(anon_sym_DASH),
	7489:  uint16(anon_sym_not),
	7490:  uint16(261),
	7491:  uint16(3),
	7492:  uint16(sym_nil),
	7493:  uint16(sym_false),
	7494:  uint16(sym_true),
	7495:  uint16(2),
	7496:  uint16(3),
	7497:  uint16(sym_bracket_index_expression),
	7498:  uint16(sym_dot_index_expression),
	7499:  uint16(sym__contextual_keyword),
	7500:  uint16(7),
	7501:  uint16(3),
	7502:  uint16(sym_variable),
	7503:  uint16(sym_function_call),
	7504:  uint16(sym_parenthesized_expression),
	7505:  uint16(40),
	7506:  uint16(5),
	7507:  uint16(sym_string),
	7508:  uint16(sym_function_definition),
	7509:  uint16(sym_table_constructor),
	7510:  uint16(sym_binary_expression),
	7511:  uint16(sym_unary_expression),
	7512:  uint16(23),
	7513:  uint16(3),
	7514:  uint16(1),
	7515:  uint16(anon_sym_DASH_DASH),
	7516:  uint16(5),
	7517:  uint16(1),
	7518:  uint16(sym__block_comment_start),
	7519:  uint16(9),
	7520:  uint16(1),
	7521:  uint16(sym_identifier),
	7522:  uint16(39),
	7523:  uint16(1),
	7524:  uint16(anon_sym_LPAREN),
	7525:  uint16(257),
	7526:  uint16(1),
	7527:  uint16(anon_sym_function),
	7528:  uint16(259),
	7529:  uint16(1),
	7530:  uint16(anon_sym_global),
	7531:  uint16(265),
	7532:  uint16(1),
	7533:  uint16(anon_sym_DQUOTE),
	7534:  uint16(267),
	7535:  uint16(1),
	7536:  uint16(anon_sym_SQUOTE),
	7537:  uint16(269),
	7538:  uint16(1),
	7539:  uint16(anon_sym_LBRACE),
	7540:  uint16(275),
	7541:  uint16(1),
	7542:  uint16(sym__block_string_start),
	7543:  uint16(41),
	7544:  uint16(1),
	7545:  uint16(sym__quote_string),
	7546:  uint16(45),
	7547:  uint16(1),
	7548:  uint16(sym__block_string),
	7549:  uint16(93),
	7550:  uint16(1),
	7551:  uint16(sym_comment),
	7552:  uint16(106),
	7553:  uint16(1),
	7554:  uint16(sym_expression),
	7555:  uint16(170),
	7556:  uint16(1),
	7557:  uint16(sym__prefix_expression),
	7558:  uint16(174),
	7559:  uint16(1),
	7560:  uint16(sym_method_index_expression),
	7561:  uint16(263),
	7562:  uint16(2),
	7563:  uint16(sym_number),
	7564:  uint16(sym_vararg_expression),
	7565:  uint16(271),
	7566:  uint16(2),
	7567:  uint16(anon_sym_TILDE),
	7568:  uint16(anon_sym_POUND),
	7569:  uint16(273),
	7570:  uint16(2),
	7571:  uint16(anon_sym_DASH),
	7572:  uint16(anon_sym_not),
	7573:  uint16(261),
	7574:  uint16(3),
	7575:  uint16(sym_nil),
	7576:  uint16(sym_false),
	7577:  uint16(sym_true),
	7578:  uint16(2),
	7579:  uint16(3),
	7580:  uint16(sym_bracket_index_expression),
	7581:  uint16(sym_dot_index_expression),
	7582:  uint16(sym__contextual_keyword),
	7583:  uint16(7),
	7584:  uint16(3),
	7585:  uint16(sym_variable),
	7586:  uint16(sym_function_call),
	7587:  uint16(sym_parenthesized_expression),
	7588:  uint16(40),
	7589:  uint16(5),
	7590:  uint16(sym_string),
	7591:  uint16(sym_function_definition),
	7592:  uint16(sym_table_constructor),
	7593:  uint16(sym_binary_expression),
	7594:  uint16(sym_unary_expression),
	7595:  uint16(23),
	7596:  uint16(3),
	7597:  uint16(1),
	7598:  uint16(anon_sym_DASH_DASH),
	7599:  uint16(5),
	7600:  uint16(1),
	7601:  uint16(sym__block_comment_start),
	7602:  uint16(9),
	7603:  uint16(1),
	7604:  uint16(sym_identifier),
	7605:  uint16(39),
	7606:  uint16(1),
	7607:  uint16(anon_sym_LPAREN),
	7608:  uint16(257),
	7609:  uint16(1),
	7610:  uint16(anon_sym_function),
	7611:  uint16(259),
	7612:  uint16(1),
	7613:  uint16(anon_sym_global),
	7614:  uint16(265),
	7615:  uint16(1),
	7616:  uint16(anon_sym_DQUOTE),
	7617:  uint16(267),
	7618:  uint16(1),
	7619:  uint16(anon_sym_SQUOTE),
	7620:  uint16(269),
	7621:  uint16(1),
	7622:  uint16(anon_sym_LBRACE),
	7623:  uint16(275),
	7624:  uint16(1),
	7625:  uint16(sym__block_string_start),
	7626:  uint16(35),
	7627:  uint16(1),
	7628:  uint16(sym_expression),
	7629:  uint16(41),
	7630:  uint16(1),
	7631:  uint16(sym__quote_string),
	7632:  uint16(45),
	7633:  uint16(1),
	7634:  uint16(sym__block_string),
	7635:  uint16(94),
	7636:  uint16(1),
	7637:  uint16(sym_comment),
	7638:  uint16(170),
	7639:  uint16(1),
	7640:  uint16(sym__prefix_expression),
	7641:  uint16(174),
	7642:  uint16(1),
	7643:  uint16(sym_method_index_expression),
	7644:  uint16(263),
	7645:  uint16(2),
	7646:  uint16(sym_number),
	7647:  uint16(sym_vararg_expression),
	7648:  uint16(271),
	7649:  uint16(2),
	7650:  uint16(anon_sym_TILDE),
	7651:  uint16(anon_sym_POUND),
	7652:  uint16(273),
	7653:  uint16(2),
	7654:  uint16(anon_sym_DASH),
	7655:  uint16(anon_sym_not),
	7656:  uint16(261),
	7657:  uint16(3),
	7658:  uint16(sym_nil),
	7659:  uint16(sym_false),
	7660:  uint16(sym_true),
	7661:  uint16(2),
	7662:  uint16(3),
	7663:  uint16(sym_bracket_index_expression),
	7664:  uint16(sym_dot_index_expression),
	7665:  uint16(sym__contextual_keyword),
	7666:  uint16(7),
	7667:  uint16(3),
	7668:  uint16(sym_variable),
	7669:  uint16(sym_function_call),
	7670:  uint16(sym_parenthesized_expression),
	7671:  uint16(40),
	7672:  uint16(5),
	7673:  uint16(sym_string),
	7674:  uint16(sym_function_definition),
	7675:  uint16(sym_table_constructor),
	7676:  uint16(sym_binary_expression),
	7677:  uint16(sym_unary_expression),
	7678:  uint16(23),
	7679:  uint16(3),
	7680:  uint16(1),
	7681:  uint16(anon_sym_DASH_DASH),
	7682:  uint16(5),
	7683:  uint16(1),
	7684:  uint16(sym__block_comment_start),
	7685:  uint16(9),
	7686:  uint16(1),
	7687:  uint16(sym_identifier),
	7688:  uint16(39),
	7689:  uint16(1),
	7690:  uint16(anon_sym_LPAREN),
	7691:  uint16(257),
	7692:  uint16(1),
	7693:  uint16(anon_sym_function),
	7694:  uint16(259),
	7695:  uint16(1),
	7696:  uint16(anon_sym_global),
	7697:  uint16(265),
	7698:  uint16(1),
	7699:  uint16(anon_sym_DQUOTE),
	7700:  uint16(267),
	7701:  uint16(1),
	7702:  uint16(anon_sym_SQUOTE),
	7703:  uint16(269),
	7704:  uint16(1),
	7705:  uint16(anon_sym_LBRACE),
	7706:  uint16(275),
	7707:  uint16(1),
	7708:  uint16(sym__block_string_start),
	7709:  uint16(36),
	7710:  uint16(1),
	7711:  uint16(sym_expression),
	7712:  uint16(41),
	7713:  uint16(1),
	7714:  uint16(sym__quote_string),
	7715:  uint16(45),
	7716:  uint16(1),
	7717:  uint16(sym__block_string),
	7718:  uint16(95),
	7719:  uint16(1),
	7720:  uint16(sym_comment),
	7721:  uint16(170),
	7722:  uint16(1),
	7723:  uint16(sym__prefix_expression),
	7724:  uint16(174),
	7725:  uint16(1),
	7726:  uint16(sym_method_index_expression),
	7727:  uint16(263),
	7728:  uint16(2),
	7729:  uint16(sym_number),
	7730:  uint16(sym_vararg_expression),
	7731:  uint16(271),
	7732:  uint16(2),
	7733:  uint16(anon_sym_TILDE),
	7734:  uint16(anon_sym_POUND),
	7735:  uint16(273),
	7736:  uint16(2),
	7737:  uint16(anon_sym_DASH),
	7738:  uint16(anon_sym_not),
	7739:  uint16(261),
	7740:  uint16(3),
	7741:  uint16(sym_nil),
	7742:  uint16(sym_false),
	7743:  uint16(sym_true),
	7744:  uint16(2),
	7745:  uint16(3),
	7746:  uint16(sym_bracket_index_expression),
	7747:  uint16(sym_dot_index_expression),
	7748:  uint16(sym__contextual_keyword),
	7749:  uint16(7),
	7750:  uint16(3),
	7751:  uint16(sym_variable),
	7752:  uint16(sym_function_call),
	7753:  uint16(sym_parenthesized_expression),
	7754:  uint16(40),
	7755:  uint16(5),
	7756:  uint16(sym_string),
	7757:  uint16(sym_function_definition),
	7758:  uint16(sym_table_constructor),
	7759:  uint16(sym_binary_expression),
	7760:  uint16(sym_unary_expression),
	7761:  uint16(23),
	7762:  uint16(3),
	7763:  uint16(1),
	7764:  uint16(anon_sym_DASH_DASH),
	7765:  uint16(5),
	7766:  uint16(1),
	7767:  uint16(sym__block_comment_start),
	7768:  uint16(9),
	7769:  uint16(1),
	7770:  uint16(sym_identifier),
	7771:  uint16(39),
	7772:  uint16(1),
	7773:  uint16(anon_sym_LPAREN),
	7774:  uint16(257),
	7775:  uint16(1),
	7776:  uint16(anon_sym_function),
	7777:  uint16(259),
	7778:  uint16(1),
	7779:  uint16(anon_sym_global),
	7780:  uint16(265),
	7781:  uint16(1),
	7782:  uint16(anon_sym_DQUOTE),
	7783:  uint16(267),
	7784:  uint16(1),
	7785:  uint16(anon_sym_SQUOTE),
	7786:  uint16(269),
	7787:  uint16(1),
	7788:  uint16(anon_sym_LBRACE),
	7789:  uint16(275),
	7790:  uint16(1),
	7791:  uint16(sym__block_string_start),
	7792:  uint16(37),
	7793:  uint16(1),
	7794:  uint16(sym_expression),
	7795:  uint16(41),
	7796:  uint16(1),
	7797:  uint16(sym__quote_string),
	7798:  uint16(45),
	7799:  uint16(1),
	7800:  uint16(sym__block_string),
	7801:  uint16(96),
	7802:  uint16(1),
	7803:  uint16(sym_comment),
	7804:  uint16(170),
	7805:  uint16(1),
	7806:  uint16(sym__prefix_expression),
	7807:  uint16(174),
	7808:  uint16(1),
	7809:  uint16(sym_method_index_expression),
	7810:  uint16(263),
	7811:  uint16(2),
	7812:  uint16(sym_number),
	7813:  uint16(sym_vararg_expression),
	7814:  uint16(271),
	7815:  uint16(2),
	7816:  uint16(anon_sym_TILDE),
	7817:  uint16(anon_sym_POUND),
	7818:  uint16(273),
	7819:  uint16(2),
	7820:  uint16(anon_sym_DASH),
	7821:  uint16(anon_sym_not),
	7822:  uint16(261),
	7823:  uint16(3),
	7824:  uint16(sym_nil),
	7825:  uint16(sym_false),
	7826:  uint16(sym_true),
	7827:  uint16(2),
	7828:  uint16(3),
	7829:  uint16(sym_bracket_index_expression),
	7830:  uint16(sym_dot_index_expression),
	7831:  uint16(sym__contextual_keyword),
	7832:  uint16(7),
	7833:  uint16(3),
	7834:  uint16(sym_variable),
	7835:  uint16(sym_function_call),
	7836:  uint16(sym_parenthesized_expression),
	7837:  uint16(40),
	7838:  uint16(5),
	7839:  uint16(sym_string),
	7840:  uint16(sym_function_definition),
	7841:  uint16(sym_table_constructor),
	7842:  uint16(sym_binary_expression),
	7843:  uint16(sym_unary_expression),
	7844:  uint16(23),
	7845:  uint16(3),
	7846:  uint16(1),
	7847:  uint16(anon_sym_DASH_DASH),
	7848:  uint16(5),
	7849:  uint16(1),
	7850:  uint16(sym__block_comment_start),
	7851:  uint16(9),
	7852:  uint16(1),
	7853:  uint16(sym_identifier),
	7854:  uint16(39),
	7855:  uint16(1),
	7856:  uint16(anon_sym_LPAREN),
	7857:  uint16(257),
	7858:  uint16(1),
	7859:  uint16(anon_sym_function),
	7860:  uint16(259),
	7861:  uint16(1),
	7862:  uint16(anon_sym_global),
	7863:  uint16(265),
	7864:  uint16(1),
	7865:  uint16(anon_sym_DQUOTE),
	7866:  uint16(267),
	7867:  uint16(1),
	7868:  uint16(anon_sym_SQUOTE),
	7869:  uint16(269),
	7870:  uint16(1),
	7871:  uint16(anon_sym_LBRACE),
	7872:  uint16(275),
	7873:  uint16(1),
	7874:  uint16(sym__block_string_start),
	7875:  uint16(41),
	7876:  uint16(1),
	7877:  uint16(sym__quote_string),
	7878:  uint16(45),
	7879:  uint16(1),
	7880:  uint16(sym__block_string),
	7881:  uint16(46),
	7882:  uint16(1),
	7883:  uint16(sym_expression),
	7884:  uint16(97),
	7885:  uint16(1),
	7886:  uint16(sym_comment),
	7887:  uint16(170),
	7888:  uint16(1),
	7889:  uint16(sym__prefix_expression),
	7890:  uint16(174),
	7891:  uint16(1),
	7892:  uint16(sym_method_index_expression),
	7893:  uint16(263),
	7894:  uint16(2),
	7895:  uint16(sym_number),
	7896:  uint16(sym_vararg_expression),
	7897:  uint16(271),
	7898:  uint16(2),
	7899:  uint16(anon_sym_TILDE),
	7900:  uint16(anon_sym_POUND),
	7901:  uint16(273),
	7902:  uint16(2),
	7903:  uint16(anon_sym_DASH),
	7904:  uint16(anon_sym_not),
	7905:  uint16(261),
	7906:  uint16(3),
	7907:  uint16(sym_nil),
	7908:  uint16(sym_false),
	7909:  uint16(sym_true),
	7910:  uint16(2),
	7911:  uint16(3),
	7912:  uint16(sym_bracket_index_expression),
	7913:  uint16(sym_dot_index_expression),
	7914:  uint16(sym__contextual_keyword),
	7915:  uint16(7),
	7916:  uint16(3),
	7917:  uint16(sym_variable),
	7918:  uint16(sym_function_call),
	7919:  uint16(sym_parenthesized_expression),
	7920:  uint16(40),
	7921:  uint16(5),
	7922:  uint16(sym_string),
	7923:  uint16(sym_function_definition),
	7924:  uint16(sym_table_constructor),
	7925:  uint16(sym_binary_expression),
	7926:  uint16(sym_unary_expression),
	7927:  uint16(23),
	7928:  uint16(3),
	7929:  uint16(1),
	7930:  uint16(anon_sym_DASH_DASH),
	7931:  uint16(5),
	7932:  uint16(1),
	7933:  uint16(sym__block_comment_start),
	7934:  uint16(9),
	7935:  uint16(1),
	7936:  uint16(sym_identifier),
	7937:  uint16(39),
	7938:  uint16(1),
	7939:  uint16(anon_sym_LPAREN),
	7940:  uint16(257),
	7941:  uint16(1),
	7942:  uint16(anon_sym_function),
	7943:  uint16(259),
	7944:  uint16(1),
	7945:  uint16(anon_sym_global),
	7946:  uint16(265),
	7947:  uint16(1),
	7948:  uint16(anon_sym_DQUOTE),
	7949:  uint16(267),
	7950:  uint16(1),
	7951:  uint16(anon_sym_SQUOTE),
	7952:  uint16(269),
	7953:  uint16(1),
	7954:  uint16(anon_sym_LBRACE),
	7955:  uint16(275),
	7956:  uint16(1),
	7957:  uint16(sym__block_string_start),
	7958:  uint16(41),
	7959:  uint16(1),
	7960:  uint16(sym__quote_string),
	7961:  uint16(45),
	7962:  uint16(1),
	7963:  uint16(sym__block_string),
	7964:  uint16(59),
	7965:  uint16(1),
	7966:  uint16(sym_expression),
	7967:  uint16(98),
	7968:  uint16(1),
	7969:  uint16(sym_comment),
	7970:  uint16(170),
	7971:  uint16(1),
	7972:  uint16(sym__prefix_expression),
	7973:  uint16(174),
	7974:  uint16(1),
	7975:  uint16(sym_method_index_expression),
	7976:  uint16(263),
	7977:  uint16(2),
	7978:  uint16(sym_number),
	7979:  uint16(sym_vararg_expression),
	7980:  uint16(271),
	7981:  uint16(2),
	7982:  uint16(anon_sym_TILDE),
	7983:  uint16(anon_sym_POUND),
	7984:  uint16(273),
	7985:  uint16(2),
	7986:  uint16(anon_sym_DASH),
	7987:  uint16(anon_sym_not),
	7988:  uint16(261),
	7989:  uint16(3),
	7990:  uint16(sym_nil),
	7991:  uint16(sym_false),
	7992:  uint16(sym_true),
	7993:  uint16(2),
	7994:  uint16(3),
	7995:  uint16(sym_bracket_index_expression),
	7996:  uint16(sym_dot_index_expression),
	7997:  uint16(sym__contextual_keyword),
	7998:  uint16(7),
	7999:  uint16(3),
	8000:  uint16(sym_variable),
	8001:  uint16(sym_function_call),
	8002:  uint16(sym_parenthesized_expression),
	8003:  uint16(40),
	8004:  uint16(5),
	8005:  uint16(sym_string),
	8006:  uint16(sym_function_definition),
	8007:  uint16(sym_table_constructor),
	8008:  uint16(sym_binary_expression),
	8009:  uint16(sym_unary_expression),
	8010:  uint16(19),
	8011:  uint16(3),
	8012:  uint16(1),
	8013:  uint16(anon_sym_DASH_DASH),
	8014:  uint16(5),
	8015:  uint16(1),
	8016:  uint16(sym__block_comment_start),
	8017:  uint16(177),
	8018:  uint16(1),
	8019:  uint16(anon_sym_AMP),
	8020:  uint16(181),
	8021:  uint16(1),
	8022:  uint16(anon_sym_PLUS),
	8023:  uint16(183),
	8024:  uint16(1),
	8025:  uint16(anon_sym_DASH),
	8026:  uint16(185),
	8027:  uint16(1),
	8028:  uint16(anon_sym_SLASH),
	8029:  uint16(187),
	8030:  uint16(1),
	8031:  uint16(anon_sym_DOT_DOT),
	8032:  uint16(189),
	8033:  uint16(1),
	8034:  uint16(anon_sym_CARET),
	8035:  uint16(195),
	8036:  uint16(1),
	8037:  uint16(anon_sym_PIPE),
	8038:  uint16(197),
	8039:  uint16(1),
	8040:  uint16(anon_sym_TILDE),
	8041:  uint16(297),
	8042:  uint16(1),
	8043:  uint16(anon_sym_else),
	8044:  uint16(299),
	8045:  uint16(1),
	8046:  uint16(anon_sym_or),
	8047:  uint16(301),
	8048:  uint16(1),
	8049:  uint16(anon_sym_and),
	8050:  uint16(99),
	8051:  uint16(1),
	8052:  uint16(sym_comment),
	8053:  uint16(179),
	8054:  uint16(2),
	8055:  uint16(anon_sym_LT_LT),
	8056:  uint16(anon_sym_GT_GT),
	8057:  uint16(191),
	8058:  uint16(2),
	8059:  uint16(anon_sym_LT),
	8060:  uint16(anon_sym_GT),
	8061:  uint16(175),
	8062:  uint16(3),
	8063:  uint16(anon_sym_STAR),
	8064:  uint16(anon_sym_SLASH_SLASH),
	8065:  uint16(anon_sym_PERCENT),
	8066:  uint16(193),
	8067:  uint16(4),
	8068:  uint16(anon_sym_LT_EQ),
	8069:  uint16(anon_sym_EQ_EQ),
	8070:  uint16(anon_sym_TILDE_EQ),
	8071:  uint16(anon_sym_GT_EQ),
	8072:  uint16(295),
	8073:  uint16(8),
	8075:  uint16(anon_sym_SEMI),
	8076:  uint16(anon_sym_COMMA),
	8077:  uint16(anon_sym_do),
	8078:  uint16(anon_sym_end),
	8079:  uint16(anon_sym_until),
	8080:  uint16(anon_sym_elseif),
	8081:  uint16(anon_sym_RPAREN),
	8082:  uint16(21),
	8083:  uint16(3),
	8084:  uint16(1),
	8085:  uint16(anon_sym_DASH_DASH),
	8086:  uint16(5),
	8087:  uint16(1),
	8088:  uint16(sym__block_comment_start),
	8089:  uint16(177),
	8090:  uint16(1),
	8091:  uint16(anon_sym_AMP),
	8092:  uint16(181),
	8093:  uint16(1),
	8094:  uint16(anon_sym_PLUS),
	8095:  uint16(183),
	8096:  uint16(1),
	8097:  uint16(anon_sym_DASH),
	8098:  uint16(185),
	8099:  uint16(1),
	8100:  uint16(anon_sym_SLASH),
	8101:  uint16(187),
	8102:  uint16(1),
	8103:  uint16(anon_sym_DOT_DOT),
	8104:  uint16(189),
	8105:  uint16(1),
	8106:  uint16(anon_sym_CARET),
	8107:  uint16(195),
	8108:  uint16(1),
	8109:  uint16(anon_sym_PIPE),
	8110:  uint16(197),
	8111:  uint16(1),
	8112:  uint16(anon_sym_TILDE),
	8113:  uint16(299),
	8114:  uint16(1),
	8115:  uint16(anon_sym_or),
	8116:  uint16(301),
	8117:  uint16(1),
	8118:  uint16(anon_sym_and),
	8119:  uint16(305),
	8120:  uint16(1),
	8121:  uint16(anon_sym_COMMA),
	8122:  uint16(307),
	8123:  uint16(1),
	8124:  uint16(anon_sym_else),
	8125:  uint16(100),
	8126:  uint16(1),
	8127:  uint16(sym_comment),
	8128:  uint16(176),
	8129:  uint16(1),
	8130:  uint16(aux_sym__expression_list_repeat1),
	8131:  uint16(179),
	8132:  uint16(2),
	8133:  uint16(anon_sym_LT_LT),
	8134:  uint16(anon_sym_GT_GT),
	8135:  uint16(191),
	8136:  uint16(2),
	8137:  uint16(anon_sym_LT),
	8138:  uint16(anon_sym_GT),
	8139:  uint16(175),
	8140:  uint16(3),
	8141:  uint16(anon_sym_STAR),
	8142:  uint16(anon_sym_SLASH_SLASH),
	8143:  uint16(anon_sym_PERCENT),
	8144:  uint16(193),
	8145:  uint16(4),
	8146:  uint16(anon_sym_LT_EQ),
	8147:  uint16(anon_sym_EQ_EQ),
	8148:  uint16(anon_sym_TILDE_EQ),
	8149:  uint16(anon_sym_GT_EQ),
	8150:  uint16(303),
	8151:  uint16(6),
	8153:  uint16(anon_sym_SEMI),
	8154:  uint16(anon_sym_do),
	8155:  uint16(anon_sym_end),
	8156:  uint16(anon_sym_until),
	8157:  uint16(anon_sym_elseif),
	8158:  uint16(7),
	8159:  uint16(3),
	8160:  uint16(1),
	8161:  uint16(anon_sym_DASH_DASH),
	8162:  uint16(5),
	8163:  uint16(1),
	8164:  uint16(sym__block_comment_start),
	8165:  uint16(65),
	8166:  uint16(1),
	8167:  uint16(anon_sym_COLON),
	8168:  uint16(101),
	8169:  uint16(1),
	8170:  uint16(sym_comment),
	8171:  uint16(309),
	8172:  uint16(3),
	8174:  uint16(anon_sym_SEMI),
	8175:  uint16(anon_sym_COLON_COLON),
	8176:  uint16(67),
	8177:  uint16(7),
	8178:  uint16(sym__block_string_start),
	8179:  uint16(anon_sym_DOT),
	8180:  uint16(anon_sym_DQUOTE),
	8181:  uint16(anon_sym_SQUOTE),
	8182:  uint16(anon_sym_LPAREN),
	8183:  uint16(anon_sym_LBRACK),
	8184:  uint16(anon_sym_LBRACE),
	8185:  uint16(311),
	8186:  uint16(16),
	8187:  uint16(anon_sym_return),
	8188:  uint16(sym_break_statement),
	8189:  uint16(anon_sym_goto),
	8190:  uint16(anon_sym_do),
	8191:  uint16(anon_sym_end),
	8192:  uint16(anon_sym_while),
	8193:  uint16(anon_sym_repeat),
	8194:  uint16(anon_sym_until),
	8195:  uint16(anon_sym_if),
	8196:  uint16(anon_sym_elseif),
	8197:  uint16(anon_sym_else),
	8198:  uint16(anon_sym_for),
	8199:  uint16(anon_sym_function),
	8200:  uint16(anon_sym_local),
	8201:  uint16(anon_sym_global),
	8202:  uint16(sym_identifier),
	8203:  uint16(9),
	8204:  uint16(3),
	8205:  uint16(1),
	8206:  uint16(anon_sym_DASH_DASH),
	8207:  uint16(5),
	8208:  uint16(1),
	8209:  uint16(sym__block_comment_start),
	8210:  uint16(317),
	8211:  uint16(1),
	8212:  uint16(anon_sym_COMMA),
	8213:  uint16(319),
	8214:  uint16(1),
	8215:  uint16(anon_sym_LT),
	8216:  uint16(102),
	8217:  uint16(1),
	8218:  uint16(sym_comment),
	8219:  uint16(113),
	8220:  uint16(1),
	8221:  uint16(aux_sym__att_name_list_repeat1),
	8222:  uint16(116),
	8223:  uint16(1),
	8224:  uint16(sym__attrib),
	8225:  uint16(313),
	8226:  uint16(5),
	8228:  uint16(anon_sym_SEMI),
	8229:  uint16(anon_sym_EQ),
	8230:  uint16(anon_sym_COLON_COLON),
	8231:  uint16(anon_sym_LPAREN),
	8232:  uint16(315),
	8233:  uint16(16),
	8234:  uint16(anon_sym_return),
	8235:  uint16(sym_break_statement),
	8236:  uint16(anon_sym_goto),
	8237:  uint16(anon_sym_do),
	8238:  uint16(anon_sym_end),
	8239:  uint16(anon_sym_while),
	8240:  uint16(anon_sym_repeat),
	8241:  uint16(anon_sym_until),
	8242:  uint16(anon_sym_if),
	8243:  uint16(anon_sym_elseif),
	8244:  uint16(anon_sym_else),
	8245:  uint16(anon_sym_for),
	8246:  uint16(anon_sym_function),
	8247:  uint16(anon_sym_local),
	8248:  uint16(anon_sym_global),
	8249:  uint16(sym_identifier),
	8250:  uint16(9),
	8251:  uint16(3),
	8252:  uint16(1),
	8253:  uint16(anon_sym_DASH_DASH),
	8254:  uint16(5),
	8255:  uint16(1),
	8256:  uint16(sym__block_comment_start),
	8257:  uint16(317),
	8258:  uint16(1),
	8259:  uint16(anon_sym_COMMA),
	8260:  uint16(319),
	8261:  uint16(1),
	8262:  uint16(anon_sym_LT),
	8263:  uint16(103),
	8264:  uint16(1),
	8265:  uint16(sym_comment),
	8266:  uint16(110),
	8267:  uint16(1),
	8268:  uint16(sym__attrib),
	8269:  uint16(111),
	8270:  uint16(1),
	8271:  uint16(aux_sym__att_name_list_repeat1),
	8272:  uint16(321),
	8273:  uint16(5),
	8275:  uint16(anon_sym_SEMI),
	8276:  uint16(anon_sym_EQ),
	8277:  uint16(anon_sym_COLON_COLON),
	8278:  uint16(anon_sym_LPAREN),
	8279:  uint16(323),
	8280:  uint16(16),
	8281:  uint16(anon_sym_return),
	8282:  uint16(sym_break_statement),
	8283:  uint16(anon_sym_goto),
	8284:  uint16(anon_sym_do),
	8285:  uint16(anon_sym_end),
	8286:  uint16(anon_sym_while),
	8287:  uint16(anon_sym_repeat),
	8288:  uint16(anon_sym_until),
	8289:  uint16(anon_sym_if),
	8290:  uint16(anon_sym_elseif),
	8291:  uint16(anon_sym_else),
	8292:  uint16(anon_sym_for),
	8293:  uint16(anon_sym_function),
	8294:  uint16(anon_sym_local),
	8295:  uint16(anon_sym_global),
	8296:  uint16(sym_identifier),
	8297:  uint16(20),
	8298:  uint16(3),
	8299:  uint16(1),
	8300:  uint16(anon_sym_DASH_DASH),
	8301:  uint16(5),
	8302:  uint16(1),
	8303:  uint16(sym__block_comment_start),
	8304:  uint16(177),
	8305:  uint16(1),
	8306:  uint16(anon_sym_AMP),
	8307:  uint16(181),
	8308:  uint16(1),
	8309:  uint16(anon_sym_PLUS),
	8310:  uint16(183),
	8311:  uint16(1),
	8312:  uint16(anon_sym_DASH),
	8313:  uint16(185),
	8314:  uint16(1),
	8315:  uint16(anon_sym_SLASH),
	8316:  uint16(187),
	8317:  uint16(1),
	8318:  uint16(anon_sym_DOT_DOT),
	8319:  uint16(189),
	8320:  uint16(1),
	8321:  uint16(anon_sym_CARET),
	8322:  uint16(195),
	8323:  uint16(1),
	8324:  uint16(anon_sym_PIPE),
	8325:  uint16(197),
	8326:  uint16(1),
	8327:  uint16(anon_sym_TILDE),
	8328:  uint16(299),
	8329:  uint16(1),
	8330:  uint16(anon_sym_or),
	8331:  uint16(301),
	8332:  uint16(1),
	8333:  uint16(anon_sym_and),
	8334:  uint16(305),
	8335:  uint16(1),
	8336:  uint16(anon_sym_COMMA),
	8337:  uint16(325),
	8338:  uint16(1),
	8339:  uint16(anon_sym_RPAREN),
	8340:  uint16(104),
	8341:  uint16(1),
	8342:  uint16(sym_comment),
	8343:  uint16(212),
	8344:  uint16(1),
	8345:  uint16(aux_sym__expression_list_repeat1),
	8346:  uint16(179),
	8347:  uint16(2),
	8348:  uint16(anon_sym_LT_LT),
	8349:  uint16(anon_sym_GT_GT),
	8350:  uint16(191),
	8351:  uint16(2),
	8352:  uint16(anon_sym_LT),
	8353:  uint16(anon_sym_GT),
	8354:  uint16(175),
	8355:  uint16(3),
	8356:  uint16(anon_sym_STAR),
	8357:  uint16(anon_sym_SLASH_SLASH),
	8358:  uint16(anon_sym_PERCENT),
	8359:  uint16(193),
	8360:  uint16(4),
	8361:  uint16(anon_sym_LT_EQ),
	8362:  uint16(anon_sym_EQ_EQ),
	8363:  uint16(anon_sym_TILDE_EQ),
	8364:  uint16(anon_sym_GT_EQ),
	8365:  uint16(18),
	8366:  uint16(3),
	8367:  uint16(1),
	8368:  uint16(anon_sym_DASH_DASH),
	8369:  uint16(5),
	8370:  uint16(1),
	8371:  uint16(sym__block_comment_start),
	8372:  uint16(177),
	8373:  uint16(1),
	8374:  uint16(anon_sym_AMP),
	8375:  uint16(181),
	8376:  uint16(1),
	8377:  uint16(anon_sym_PLUS),
	8378:  uint16(183),
	8379:  uint16(1),
	8380:  uint16(anon_sym_DASH),
	8381:  uint16(185),
	8382:  uint16(1),
	8383:  uint16(anon_sym_SLASH),
	8384:  uint16(187),
	8385:  uint16(1),
	8386:  uint16(anon_sym_DOT_DOT),
	8387:  uint16(189),
	8388:  uint16(1),
	8389:  uint16(anon_sym_CARET),
	8390:  uint16(195),
	8391:  uint16(1),
	8392:  uint16(anon_sym_PIPE),
	8393:  uint16(197),
	8394:  uint16(1),
	8395:  uint16(anon_sym_TILDE),
	8396:  uint16(299),
	8397:  uint16(1),
	8398:  uint16(anon_sym_or),
	8399:  uint16(301),
	8400:  uint16(1),
	8401:  uint16(anon_sym_and),
	8402:  uint16(105),
	8403:  uint16(1),
	8404:  uint16(sym_comment),
	8405:  uint16(179),
	8406:  uint16(2),
	8407:  uint16(anon_sym_LT_LT),
	8408:  uint16(anon_sym_GT_GT),
	8409:  uint16(191),
	8410:  uint16(2),
	8411:  uint16(anon_sym_LT),
	8412:  uint16(anon_sym_GT),
	8413:  uint16(175),
	8414:  uint16(3),
	8415:  uint16(anon_sym_STAR),
	8416:  uint16(anon_sym_SLASH_SLASH),
	8417:  uint16(anon_sym_PERCENT),
	8418:  uint16(327),
	8419:  uint16(3),
	8420:  uint16(anon_sym_SEMI),
	8421:  uint16(anon_sym_COMMA),
	8422:  uint16(anon_sym_RBRACE),
	8423:  uint16(193),
	8424:  uint16(4),
	8425:  uint16(anon_sym_LT_EQ),
	8426:  uint16(anon_sym_EQ_EQ),
	8427:  uint16(anon_sym_TILDE_EQ),
	8428:  uint16(anon_sym_GT_EQ),
	8429:  uint16(18),
	8430:  uint16(3),
	8431:  uint16(1),
	8432:  uint16(anon_sym_DASH_DASH),
	8433:  uint16(5),
	8434:  uint16(1),
	8435:  uint16(sym__block_comment_start),
	8436:  uint16(177),
	8437:  uint16(1),
	8438:  uint16(anon_sym_AMP),
	8439:  uint16(181),
	8440:  uint16(1),
	8441:  uint16(anon_sym_PLUS),
	8442:  uint16(183),
	8443:  uint16(1),
	8444:  uint16(anon_sym_DASH),
	8445:  uint16(185),
	8446:  uint16(1),
	8447:  uint16(anon_sym_SLASH),
	8448:  uint16(187),
	8449:  uint16(1),
	8450:  uint16(anon_sym_DOT_DOT),
	8451:  uint16(189),
	8452:  uint16(1),
	8453:  uint16(anon_sym_CARET),
	8454:  uint16(195),
	8455:  uint16(1),
	8456:  uint16(anon_sym_PIPE),
	8457:  uint16(197),
	8458:  uint16(1),
	8459:  uint16(anon_sym_TILDE),
	8460:  uint16(299),
	8461:  uint16(1),
	8462:  uint16(anon_sym_or),
	8463:  uint16(301),
	8464:  uint16(1),
	8465:  uint16(anon_sym_and),
	8466:  uint16(106),
	8467:  uint16(1),
	8468:  uint16(sym_comment),
	8469:  uint16(179),
	8470:  uint16(2),
	8471:  uint16(anon_sym_LT_LT),
	8472:  uint16(anon_sym_GT_GT),
	8473:  uint16(191),
	8474:  uint16(2),
	8475:  uint16(anon_sym_LT),
	8476:  uint16(anon_sym_GT),
	8477:  uint16(175),
	8478:  uint16(3),
	8479:  uint16(anon_sym_STAR),
	8480:  uint16(anon_sym_SLASH_SLASH),
	8481:  uint16(anon_sym_PERCENT),
	8482:  uint16(329),
	8483:  uint16(3),
	8484:  uint16(anon_sym_SEMI),
	8485:  uint16(anon_sym_COMMA),
	8486:  uint16(anon_sym_RBRACE),
	8487:  uint16(193),
	8488:  uint16(4),
	8489:  uint16(anon_sym_LT_EQ),
	8490:  uint16(anon_sym_EQ_EQ),
	8491:  uint16(anon_sym_TILDE_EQ),
	8492:  uint16(anon_sym_GT_EQ),
	8493:  uint16(7),
	8494:  uint16(3),
	8495:  uint16(1),
	8496:  uint16(anon_sym_DASH_DASH),
	8497:  uint16(5),
	8498:  uint16(1),
	8499:  uint16(sym__block_comment_start),
	8500:  uint16(319),
	8501:  uint16(1),
	8502:  uint16(anon_sym_LT),
	8503:  uint16(107),
	8504:  uint16(1),
	8505:  uint16(sym_comment),
	8506:  uint16(122),
	8507:  uint16(1),
	8508:  uint16(sym__attrib),
	8509:  uint16(331),
	8510:  uint16(6),
	8512:  uint16(anon_sym_SEMI),
	8513:  uint16(anon_sym_EQ),
	8514:  uint16(anon_sym_COMMA),
	8515:  uint16(anon_sym_COLON_COLON),
	8516:  uint16(anon_sym_LPAREN),
	8517:  uint16(333),
	8518:  uint16(16),
	8519:  uint16(anon_sym_return),
	8520:  uint16(sym_break_statement),
	8521:  uint16(anon_sym_goto),
	8522:  uint16(anon_sym_do),
	8523:  uint16(anon_sym_end),
	8524:  uint16(anon_sym_while),
	8525:  uint16(anon_sym_repeat),
	8526:  uint16(anon_sym_until),
	8527:  uint16(anon_sym_if),
	8528:  uint16(anon_sym_elseif),
	8529:  uint16(anon_sym_else),
	8530:  uint16(anon_sym_for),
	8531:  uint16(anon_sym_function),
	8532:  uint16(anon_sym_local),
	8533:  uint16(anon_sym_global),
	8534:  uint16(sym_identifier),
	8535:  uint16(18),
	8536:  uint16(3),
	8537:  uint16(1),
	8538:  uint16(anon_sym_DASH_DASH),
	8539:  uint16(5),
	8540:  uint16(1),
	8541:  uint16(sym__block_comment_start),
	8542:  uint16(177),
	8543:  uint16(1),
	8544:  uint16(anon_sym_AMP),
	8545:  uint16(181),
	8546:  uint16(1),
	8547:  uint16(anon_sym_PLUS),
	8548:  uint16(183),
	8549:  uint16(1),
	8550:  uint16(anon_sym_DASH),
	8551:  uint16(185),
	8552:  uint16(1),
	8553:  uint16(anon_sym_SLASH),
	8554:  uint16(187),
	8555:  uint16(1),
	8556:  uint16(anon_sym_DOT_DOT),
	8557:  uint16(189),
	8558:  uint16(1),
	8559:  uint16(anon_sym_CARET),
	8560:  uint16(195),
	8561:  uint16(1),
	8562:  uint16(anon_sym_PIPE),
	8563:  uint16(197),
	8564:  uint16(1),
	8565:  uint16(anon_sym_TILDE),
	8566:  uint16(299),
	8567:  uint16(1),
	8568:  uint16(anon_sym_or),
	8569:  uint16(301),
	8570:  uint16(1),
	8571:  uint16(anon_sym_and),
	8572:  uint16(108),
	8573:  uint16(1),
	8574:  uint16(sym_comment),
	8575:  uint16(179),
	8576:  uint16(2),
	8577:  uint16(anon_sym_LT_LT),
	8578:  uint16(anon_sym_GT_GT),
	8579:  uint16(191),
	8580:  uint16(2),
	8581:  uint16(anon_sym_LT),
	8582:  uint16(anon_sym_GT),
	8583:  uint16(175),
	8584:  uint16(3),
	8585:  uint16(anon_sym_STAR),
	8586:  uint16(anon_sym_SLASH_SLASH),
	8587:  uint16(anon_sym_PERCENT),
	8588:  uint16(335),
	8589:  uint16(3),
	8590:  uint16(anon_sym_SEMI),
	8591:  uint16(anon_sym_COMMA),
	8592:  uint16(anon_sym_RBRACE),
	8593:  uint16(193),
	8594:  uint16(4),
	8595:  uint16(anon_sym_LT_EQ),
	8596:  uint16(anon_sym_EQ_EQ),
	8597:  uint16(anon_sym_TILDE_EQ),
	8598:  uint16(anon_sym_GT_EQ),
	8599:  uint16(7),
	8600:  uint16(3),
	8601:  uint16(1),
	8602:  uint16(anon_sym_DASH_DASH),
	8603:  uint16(5),
	8604:  uint16(1),
	8605:  uint16(sym__block_comment_start),
	8606:  uint16(317),
	8607:  uint16(1),
	8608:  uint16(anon_sym_COMMA),
	8609:  uint16(109),
	8610:  uint16(1),
	8611:  uint16(sym_comment),
	8612:  uint16(117),
	8613:  uint16(1),
	8614:  uint16(aux_sym__att_name_list_repeat1),
	8615:  uint16(337),
	8616:  uint16(5),
	8618:  uint16(anon_sym_SEMI),
	8619:  uint16(anon_sym_EQ),
	8620:  uint16(anon_sym_COLON_COLON),
	8621:  uint16(anon_sym_LPAREN),
	8622:  uint16(339),
	8623:  uint16(16),
	8624:  uint16(anon_sym_return),
	8625:  uint16(sym_break_statement),
	8626:  uint16(anon_sym_goto),
	8627:  uint16(anon_sym_do),
	8628:  uint16(anon_sym_end),
	8629:  uint16(anon_sym_while),
	8630:  uint16(anon_sym_repeat),
	8631:  uint16(anon_sym_until),
	8632:  uint16(anon_sym_if),
	8633:  uint16(anon_sym_elseif),
	8634:  uint16(anon_sym_else),
	8635:  uint16(anon_sym_for),
	8636:  uint16(anon_sym_function),
	8637:  uint16(anon_sym_local),
	8638:  uint16(anon_sym_global),
	8639:  uint16(sym_identifier),
	8640:  uint16(7),
	8641:  uint16(3),
	8642:  uint16(1),
	8643:  uint16(anon_sym_DASH_DASH),
	8644:  uint16(5),
	8645:  uint16(1),
	8646:  uint16(sym__block_comment_start),
	8647:  uint16(317),
	8648:  uint16(1),
	8649:  uint16(anon_sym_COMMA),
	8650:  uint16(110),
	8651:  uint16(1),
	8652:  uint16(sym_comment),
	8653:  uint16(112),
	8654:  uint16(1),
	8655:  uint16(aux_sym__att_name_list_repeat1),
	8656:  uint16(341),
	8657:  uint16(5),
	8659:  uint16(anon_sym_SEMI),
	8660:  uint16(anon_sym_EQ),
	8661:  uint16(anon_sym_COLON_COLON),
	8662:  uint16(anon_sym_LPAREN),
	8663:  uint16(343),
	8664:  uint16(16),
	8665:  uint16(anon_sym_return),
	8666:  uint16(sym_break_statement),
	8667:  uint16(anon_sym_goto),
	8668:  uint16(anon_sym_do),
	8669:  uint16(anon_sym_end),
	8670:  uint16(anon_sym_while),
	8671:  uint16(anon_sym_repeat),
	8672:  uint16(anon_sym_until),
	8673:  uint16(anon_sym_if),
	8674:  uint16(anon_sym_elseif),
	8675:  uint16(anon_sym_else),
	8676:  uint16(anon_sym_for),
	8677:  uint16(anon_sym_function),
	8678:  uint16(anon_sym_local),
	8679:  uint16(anon_sym_global),
	8680:  uint16(sym_identifier),
	8681:  uint16(7),
	8682:  uint16(3),
	8683:  uint16(1),
	8684:  uint16(anon_sym_DASH_DASH),
	8685:  uint16(5),
	8686:  uint16(1),
	8687:  uint16(sym__block_comment_start),
	8688:  uint16(317),
	8689:  uint16(1),
	8690:  uint16(anon_sym_COMMA),
	8691:  uint16(111),
	8692:  uint16(1),
	8693:  uint16(sym_comment),
	8694:  uint16(117),
	8695:  uint16(1),
	8696:  uint16(aux_sym__att_name_list_repeat1),
	8697:  uint16(345),
	8698:  uint16(5),
	8700:  uint16(anon_sym_SEMI),
	8701:  uint16(anon_sym_EQ),
	8702:  uint16(anon_sym_COLON_COLON),
	8703:  uint16(anon_sym_LPAREN),
	8704:  uint16(347),
	8705:  uint16(16),
	8706:  uint16(anon_sym_return),
	8707:  uint16(sym_break_statement),
	8708:  uint16(anon_sym_goto),
	8709:  uint16(anon_sym_do),
	8710:  uint16(anon_sym_end),
	8711:  uint16(anon_sym_while),
	8712:  uint16(anon_sym_repeat),
	8713:  uint16(anon_sym_until),
	8714:  uint16(anon_sym_if),
	8715:  uint16(anon_sym_elseif),
	8716:  uint16(anon_sym_else),
	8717:  uint16(anon_sym_for),
	8718:  uint16(anon_sym_function),
	8719:  uint16(anon_sym_local),
	8720:  uint16(anon_sym_global),
	8721:  uint16(sym_identifier),
	8722:  uint16(7),
	8723:  uint16(3),
	8724:  uint16(1),
	8725:  uint16(anon_sym_DASH_DASH),
	8726:  uint16(5),
	8727:  uint16(1),
	8728:  uint16(sym__block_comment_start),
	8729:  uint16(317),
	8730:  uint16(1),
	8731:  uint16(anon_sym_COMMA),
	8732:  uint16(112),
	8733:  uint16(1),
	8734:  uint16(sym_comment),
	8735:  uint16(117),
	8736:  uint16(1),
	8737:  uint16(aux_sym__att_name_list_repeat1),
	8738:  uint16(349),
	8739:  uint16(5),
	8741:  uint16(anon_sym_SEMI),
	8742:  uint16(anon_sym_EQ),
	8743:  uint16(anon_sym_COLON_COLON),
	8744:  uint16(anon_sym_LPAREN),
	8745:  uint16(351),
	8746:  uint16(16),
	8747:  uint16(anon_sym_return),
	8748:  uint16(sym_break_statement),
	8749:  uint16(anon_sym_goto),
	8750:  uint16(anon_sym_do),
	8751:  uint16(anon_sym_end),
	8752:  uint16(anon_sym_while),
	8753:  uint16(anon_sym_repeat),
	8754:  uint16(anon_sym_until),
	8755:  uint16(anon_sym_if),
	8756:  uint16(anon_sym_elseif),
	8757:  uint16(anon_sym_else),
	8758:  uint16(anon_sym_for),
	8759:  uint16(anon_sym_function),
	8760:  uint16(anon_sym_local),
	8761:  uint16(anon_sym_global),
	8762:  uint16(sym_identifier),
	8763:  uint16(7),
	8764:  uint16(3),
	8765:  uint16(1),
	8766:  uint16(anon_sym_DASH_DASH),
	8767:  uint16(5),
	8768:  uint16(1),
	8769:  uint16(sym__block_comment_start),
	8770:  uint16(317),
	8771:  uint16(1),
	8772:  uint16(anon_sym_COMMA),
	8773:  uint16(113),
	8774:  uint16(1),
	8775:  uint16(sym_comment),
	8776:  uint16(117),
	8777:  uint16(1),
	8778:  uint16(aux_sym__att_name_list_repeat1),
	8779:  uint16(353),
	8780:  uint16(5),
	8782:  uint16(anon_sym_SEMI),
	8783:  uint16(anon_sym_EQ),
	8784:  uint16(anon_sym_COLON_COLON),
	8785:  uint16(anon_sym_LPAREN),
	8786:  uint16(355),
	8787:  uint16(16),
	8788:  uint16(anon_sym_return),
	8789:  uint16(sym_break_statement),
	8790:  uint16(anon_sym_goto),
	8791:  uint16(anon_sym_do),
	8792:  uint16(anon_sym_end),
	8793:  uint16(anon_sym_while),
	8794:  uint16(anon_sym_repeat),
	8795:  uint16(anon_sym_until),
	8796:  uint16(anon_sym_if),
	8797:  uint16(anon_sym_elseif),
	8798:  uint16(anon_sym_else),
	8799:  uint16(anon_sym_for),
	8800:  uint16(anon_sym_function),
	8801:  uint16(anon_sym_local),
	8802:  uint16(anon_sym_global),
	8803:  uint16(sym_identifier),
	8804:  uint16(5),
	8805:  uint16(3),
	8806:  uint16(1),
	8807:  uint16(anon_sym_DASH_DASH),
	8808:  uint16(5),
	8809:  uint16(1),
	8810:  uint16(sym__block_comment_start),
	8811:  uint16(114),
	8812:  uint16(1),
	8813:  uint16(sym_comment),
	8814:  uint16(357),
	8815:  uint16(7),
	8817:  uint16(anon_sym_SEMI),
	8818:  uint16(anon_sym_EQ),
	8819:  uint16(anon_sym_COMMA),
	8820:  uint16(anon_sym_COLON_COLON),
	8821:  uint16(anon_sym_STAR),
	8822:  uint16(anon_sym_LPAREN),
	8823:  uint16(359),
	8824:  uint16(16),
	8825:  uint16(anon_sym_return),
	8826:  uint16(sym_break_statement),
	8827:  uint16(anon_sym_goto),
	8828:  uint16(anon_sym_do),
	8829:  uint16(anon_sym_end),
	8830:  uint16(anon_sym_while),
	8831:  uint16(anon_sym_repeat),
	8832:  uint16(anon_sym_until),
	8833:  uint16(anon_sym_if),
	8834:  uint16(anon_sym_elseif),
	8835:  uint16(anon_sym_else),
	8836:  uint16(anon_sym_for),
	8837:  uint16(anon_sym_function),
	8838:  uint16(anon_sym_local),
	8839:  uint16(anon_sym_global),
	8840:  uint16(sym_identifier),
	8841:  uint16(19),
	8842:  uint16(3),
	8843:  uint16(1),
	8844:  uint16(anon_sym_DASH_DASH),
	8845:  uint16(5),
	8846:  uint16(1),
	8847:  uint16(sym__block_comment_start),
	8848:  uint16(177),
	8849:  uint16(1),
	8850:  uint16(anon_sym_AMP),
	8851:  uint16(181),
	8852:  uint16(1),
	8853:  uint16(anon_sym_PLUS),
	8854:  uint16(183),
	8855:  uint16(1),
	8856:  uint16(anon_sym_DASH),
	8857:  uint16(185),
	8858:  uint16(1),
	8859:  uint16(anon_sym_SLASH),
	8860:  uint16(187),
	8861:  uint16(1),
	8862:  uint16(anon_sym_DOT_DOT),
	8863:  uint16(189),
	8864:  uint16(1),
	8865:  uint16(anon_sym_CARET),
	8866:  uint16(195),
	8867:  uint16(1),
	8868:  uint16(anon_sym_PIPE),
	8869:  uint16(197),
	8870:  uint16(1),
	8871:  uint16(anon_sym_TILDE),
	8872:  uint16(299),
	8873:  uint16(1),
	8874:  uint16(anon_sym_or),
	8875:  uint16(301),
	8876:  uint16(1),
	8877:  uint16(anon_sym_and),
	8878:  uint16(361),
	8879:  uint16(1),
	8880:  uint16(anon_sym_COMMA),
	8881:  uint16(363),
	8882:  uint16(1),
	8883:  uint16(anon_sym_do),
	8884:  uint16(115),
	8885:  uint16(1),
	8886:  uint16(sym_comment),
	8887:  uint16(179),
	8888:  uint16(2),
	8889:  uint16(anon_sym_LT_LT),
	8890:  uint16(anon_sym_GT_GT),
	8891:  uint16(191),
	8892:  uint16(2),
	8893:  uint16(anon_sym_LT),
	8894:  uint16(anon_sym_GT),
	8895:  uint16(175),
	8896:  uint16(3),
	8897:  uint16(anon_sym_STAR),
	8898:  uint16(anon_sym_SLASH_SLASH),
	8899:  uint16(anon_sym_PERCENT),
	8900:  uint16(193),
	8901:  uint16(4),
	8902:  uint16(anon_sym_LT_EQ),
	8903:  uint16(anon_sym_EQ_EQ),
	8904:  uint16(anon_sym_TILDE_EQ),
	8905:  uint16(anon_sym_GT_EQ),
	8906:  uint16(7),
	8907:  uint16(3),
	8908:  uint16(1),
	8909:  uint16(anon_sym_DASH_DASH),
	8910:  uint16(5),
	8911:  uint16(1),
	8912:  uint16(sym__block_comment_start),
	8913:  uint16(317),
	8914:  uint16(1),
	8915:  uint16(anon_sym_COMMA),
	8916:  uint16(109),
	8917:  uint16(1),
	8918:  uint16(aux_sym__att_name_list_repeat1),
	8919:  uint16(116),
	8920:  uint16(1),
	8921:  uint16(sym_comment),
	8922:  uint16(365),
	8923:  uint16(5),
	8925:  uint16(anon_sym_SEMI),
	8926:  uint16(anon_sym_EQ),
	8927:  uint16(anon_sym_COLON_COLON),
	8928:  uint16(anon_sym_LPAREN),
	8929:  uint16(367),
	8930:  uint16(16),
	8931:  uint16(anon_sym_return),
	8932:  uint16(sym_break_statement),
	8933:  uint16(anon_sym_goto),
	8934:  uint16(anon_sym_do),
	8935:  uint16(anon_sym_end),
	8936:  uint16(anon_sym_while),
	8937:  uint16(anon_sym_repeat),
	8938:  uint16(anon_sym_until),
	8939:  uint16(anon_sym_if),
	8940:  uint16(anon_sym_elseif),
	8941:  uint16(anon_sym_else),
	8942:  uint16(anon_sym_for),
	8943:  uint16(anon_sym_function),
	8944:  uint16(anon_sym_local),
	8945:  uint16(anon_sym_global),
	8946:  uint16(sym_identifier),
	8947:  uint16(6),
	8948:  uint16(3),
	8949:  uint16(1),
	8950:  uint16(anon_sym_DASH_DASH),
	8951:  uint16(5),
	8952:  uint16(1),
	8953:  uint16(sym__block_comment_start),
	8954:  uint16(373),
	8955:  uint16(1),
	8956:  uint16(anon_sym_COMMA),
	8957:  uint16(117),
	8958:  uint16(2),
	8959:  uint16(sym_comment),
	8960:  uint16(aux_sym__att_name_list_repeat1),
	8961:  uint16(369),
	8962:  uint16(5),
	8964:  uint16(anon_sym_SEMI),
	8965:  uint16(anon_sym_EQ),
	8966:  uint16(anon_sym_COLON_COLON),
	8967:  uint16(anon_sym_LPAREN),
	8968:  uint16(371),
	8969:  uint16(16),
	8970:  uint16(anon_sym_return),
	8971:  uint16(sym_break_statement),
	8972:  uint16(anon_sym_goto),
	8973:  uint16(anon_sym_do),
	8974:  uint16(anon_sym_end),
	8975:  uint16(anon_sym_while),
	8976:  uint16(anon_sym_repeat),
	8977:  uint16(anon_sym_until),
	8978:  uint16(anon_sym_if),
	8979:  uint16(anon_sym_elseif),
	8980:  uint16(anon_sym_else),
	8981:  uint16(anon_sym_for),
	8982:  uint16(anon_sym_function),
	8983:  uint16(anon_sym_local),
	8984:  uint16(anon_sym_global),
	8985:  uint16(sym_identifier),
	8986:  uint16(18),
	8987:  uint16(3),
	8988:  uint16(1),
	8989:  uint16(anon_sym_DASH_DASH),
	8990:  uint16(5),
	8991:  uint16(1),
	8992:  uint16(sym__block_comment_start),
	8993:  uint16(177),
	8994:  uint16(1),
	8995:  uint16(anon_sym_AMP),
	8996:  uint16(181),
	8997:  uint16(1),
	8998:  uint16(anon_sym_PLUS),
	8999:  uint16(183),
	9000:  uint16(1),
	9001:  uint16(anon_sym_DASH),
	9002:  uint16(185),
	9003:  uint16(1),
	9004:  uint16(anon_sym_SLASH),
	9005:  uint16(187),
	9006:  uint16(1),
	9007:  uint16(anon_sym_DOT_DOT),
	9008:  uint16(189),
	9009:  uint16(1),
	9010:  uint16(anon_sym_CARET),
	9011:  uint16(195),
	9012:  uint16(1),
	9013:  uint16(anon_sym_PIPE),
	9014:  uint16(197),
	9015:  uint16(1),
	9016:  uint16(anon_sym_TILDE),
	9017:  uint16(299),
	9018:  uint16(1),
	9019:  uint16(anon_sym_or),
	9020:  uint16(301),
	9021:  uint16(1),
	9022:  uint16(anon_sym_and),
	9023:  uint16(376),
	9024:  uint16(1),
	9025:  uint16(anon_sym_RBRACK),
	9026:  uint16(118),
	9027:  uint16(1),
	9028:  uint16(sym_comment),
	9029:  uint16(179),
	9030:  uint16(2),
	9031:  uint16(anon_sym_LT_LT),
	9032:  uint16(anon_sym_GT_GT),
	9033:  uint16(191),
	9034:  uint16(2),
	9035:  uint16(anon_sym_LT),
	9036:  uint16(anon_sym_GT),
	9037:  uint16(175),
	9038:  uint16(3),
	9039:  uint16(anon_sym_STAR),
	9040:  uint16(anon_sym_SLASH_SLASH),
	9041:  uint16(anon_sym_PERCENT),
	9042:  uint16(193),
	9043:  uint16(4),
	9044:  uint16(anon_sym_LT_EQ),
	9045:  uint16(anon_sym_EQ_EQ),
	9046:  uint16(anon_sym_TILDE_EQ),
	9047:  uint16(anon_sym_GT_EQ),
	9048:  uint16(18),
	9049:  uint16(3),
	9050:  uint16(1),
	9051:  uint16(anon_sym_DASH_DASH),
	9052:  uint16(5),
	9053:  uint16(1),
	9054:  uint16(sym__block_comment_start),
	9055:  uint16(177),
	9056:  uint16(1),
	9057:  uint16(anon_sym_AMP),
	9058:  uint16(181),
	9059:  uint16(1),
	9060:  uint16(anon_sym_PLUS),
	9061:  uint16(183),
	9062:  uint16(1),
	9063:  uint16(anon_sym_DASH),
	9064:  uint16(185),
	9065:  uint16(1),
	9066:  uint16(anon_sym_SLASH),
	9067:  uint16(187),
	9068:  uint16(1),
	9069:  uint16(anon_sym_DOT_DOT),
	9070:  uint16(189),
	9071:  uint16(1),
	9072:  uint16(anon_sym_CARET),
	9073:  uint16(195),
	9074:  uint16(1),
	9075:  uint16(anon_sym_PIPE),
	9076:  uint16(197),
	9077:  uint16(1),
	9078:  uint16(anon_sym_TILDE),
	9079:  uint16(299),
	9080:  uint16(1),
	9081:  uint16(anon_sym_or),
	9082:  uint16(301),
	9083:  uint16(1),
	9084:  uint16(anon_sym_and),
	9085:  uint16(378),
	9086:  uint16(1),
	9087:  uint16(anon_sym_then),
	9088:  uint16(119),
	9089:  uint16(1),
	9090:  uint16(sym_comment),
	9091:  uint16(179),
	9092:  uint16(2),
	9093:  uint16(anon_sym_LT_LT),
	9094:  uint16(anon_sym_GT_GT),
	9095:  uint16(191),
	9096:  uint16(2),
	9097:  uint16(anon_sym_LT),
	9098:  uint16(anon_sym_GT),
	9099:  uint16(175),
	9100:  uint16(3),
	9101:  uint16(anon_sym_STAR),
	9102:  uint16(anon_sym_SLASH_SLASH),
	9103:  uint16(anon_sym_PERCENT),
	9104:  uint16(193),
	9105:  uint16(4),
	9106:  uint16(anon_sym_LT_EQ),
	9107:  uint16(anon_sym_EQ_EQ),
	9108:  uint16(anon_sym_TILDE_EQ),
	9109:  uint16(anon_sym_GT_EQ),
	9110:  uint16(18),
	9111:  uint16(3),
	9112:  uint16(1),
	9113:  uint16(anon_sym_DASH_DASH),
	9114:  uint16(5),
	9115:  uint16(1),
	9116:  uint16(sym__block_comment_start),
	9117:  uint16(177),
	9118:  uint16(1),
	9119:  uint16(anon_sym_AMP),
	9120:  uint16(181),
	9121:  uint16(1),
	9122:  uint16(anon_sym_PLUS),
	9123:  uint16(183),
	9124:  uint16(1),
	9125:  uint16(anon_sym_DASH),
	9126:  uint16(185),
	9127:  uint16(1),
	9128:  uint16(anon_sym_SLASH),
	9129:  uint16(187),
	9130:  uint16(1),
	9131:  uint16(anon_sym_DOT_DOT),
	9132:  uint16(189),
	9133:  uint16(1),
	9134:  uint16(anon_sym_CARET),
	9135:  uint16(195),
	9136:  uint16(1),
	9137:  uint16(anon_sym_PIPE),
	9138:  uint16(197),
	9139:  uint16(1),
	9140:  uint16(anon_sym_TILDE),
	9141:  uint16(299),
	9142:  uint16(1),
	9143:  uint16(anon_sym_or),
	9144:  uint16(301),
	9145:  uint16(1),
	9146:  uint16(anon_sym_and),
	9147:  uint16(380),
	9148:  uint16(1),
	9149:  uint16(anon_sym_RPAREN),
	9150:  uint16(120),
	9151:  uint16(1),
	9152:  uint16(sym_comment),
	9153:  uint16(179),
	9154:  uint16(2),
	9155:  uint16(anon_sym_LT_LT),
	9156:  uint16(anon_sym_GT_GT),
	9157:  uint16(191),
	9158:  uint16(2),
	9159:  uint16(anon_sym_LT),
	9160:  uint16(anon_sym_GT),
	9161:  uint16(175),
	9162:  uint16(3),
	9163:  uint16(anon_sym_STAR),
	9164:  uint16(anon_sym_SLASH_SLASH),
	9165:  uint16(anon_sym_PERCENT),
	9166:  uint16(193),
	9167:  uint16(4),
	9168:  uint16(anon_sym_LT_EQ),
	9169:  uint16(anon_sym_EQ_EQ),
	9170:  uint16(anon_sym_TILDE_EQ),
	9171:  uint16(anon_sym_GT_EQ),
	9172:  uint16(6),
	9173:  uint16(3),
	9174:  uint16(1),
	9175:  uint16(anon_sym_DASH_DASH),
	9176:  uint16(5),
	9177:  uint16(1),
	9178:  uint16(sym__block_comment_start),
	9179:  uint16(386),
	9180:  uint16(1),
	9181:  uint16(anon_sym_COMMA),
	9182:  uint16(121),
	9183:  uint16(2),
	9184:  uint16(sym_comment),
	9185:  uint16(aux_sym__variable_assignment_explist_repeat1),
	9186:  uint16(382),
	9187:  uint16(4),
	9189:  uint16(anon_sym_SEMI),
	9190:  uint16(anon_sym_COLON_COLON),
	9191:  uint16(anon_sym_LPAREN),
	9192:  uint16(384),
	9193:  uint16(16),
	9194:  uint16(anon_sym_return),
	9195:  uint16(sym_break_statement),
	9196:  uint16(anon_sym_goto),
	9197:  uint16(anon_sym_do),
	9198:  uint16(anon_sym_end),
	9199:  uint16(anon_sym_while),
	9200:  uint16(anon_sym_repeat),
	9201:  uint16(anon_sym_until),
	9202:  uint16(anon_sym_if),
	9203:  uint16(anon_sym_elseif),
	9204:  uint16(anon_sym_else),
	9205:  uint16(anon_sym_for),
	9206:  uint16(anon_sym_function),
	9207:  uint16(anon_sym_local),
	9208:  uint16(anon_sym_global),
	9209:  uint16(sym_identifier),
	9210:  uint16(5),
	9211:  uint16(3),
	9212:  uint16(1),
	9213:  uint16(anon_sym_DASH_DASH),
	9214:  uint16(5),
	9215:  uint16(1),
	9216:  uint16(sym__block_comment_start),
	9217:  uint16(122),
	9218:  uint16(1),
	9219:  uint16(sym_comment),
	9220:  uint16(389),
	9221:  uint16(6),
	9223:  uint16(anon_sym_SEMI),
	9224:  uint16(anon_sym_EQ),
	9225:  uint16(anon_sym_COMMA),
	9226:  uint16(anon_sym_COLON_COLON),
	9227:  uint16(anon_sym_LPAREN),
	9228:  uint16(391),
	9229:  uint16(16),
	9230:  uint16(anon_sym_return),
	9231:  uint16(sym_break_statement),
	9232:  uint16(anon_sym_goto),
	9233:  uint16(anon_sym_do),
	9234:  uint16(anon_sym_end),
	9235:  uint16(anon_sym_while),
	9236:  uint16(anon_sym_repeat),
	9237:  uint16(anon_sym_until),
	9238:  uint16(anon_sym_if),
	9239:  uint16(anon_sym_elseif),
	9240:  uint16(anon_sym_else),
	9241:  uint16(anon_sym_for),
	9242:  uint16(anon_sym_function),
	9243:  uint16(anon_sym_local),
	9244:  uint16(anon_sym_global),
	9245:  uint16(sym_identifier),
	9246:  uint16(18),
	9247:  uint16(3),
	9248:  uint16(1),
	9249:  uint16(anon_sym_DASH_DASH),
	9250:  uint16(5),
	9251:  uint16(1),
	9252:  uint16(sym__block_comment_start),
	9253:  uint16(177),
	9254:  uint16(1),
	9255:  uint16(anon_sym_AMP),
	9256:  uint16(181),
	9257:  uint16(1),
	9258:  uint16(anon_sym_PLUS),
	9259:  uint16(183),
	9260:  uint16(1),
	9261:  uint16(anon_sym_DASH),
	9262:  uint16(185),
	9263:  uint16(1),
	9264:  uint16(anon_sym_SLASH),
	9265:  uint16(187),
	9266:  uint16(1),
	9267:  uint16(anon_sym_DOT_DOT),
	9268:  uint16(189),
	9269:  uint16(1),
	9270:  uint16(anon_sym_CARET),
	9271:  uint16(195),
	9272:  uint16(1),
	9273:  uint16(anon_sym_PIPE),
	9274:  uint16(197),
	9275:  uint16(1),
	9276:  uint16(anon_sym_TILDE),
	9277:  uint16(299),
	9278:  uint16(1),
	9279:  uint16(anon_sym_or),
	9280:  uint16(301),
	9281:  uint16(1),
	9282:  uint16(anon_sym_and),
	9283:  uint16(393),
	9284:  uint16(1),
	9285:  uint16(anon_sym_RBRACK),
	9286:  uint16(123),
	9287:  uint16(1),
	9288:  uint16(sym_comment),
	9289:  uint16(179),
	9290:  uint16(2),
	9291:  uint16(anon_sym_LT_LT),
	9292:  uint16(anon_sym_GT_GT),
	9293:  uint16(191),
	9294:  uint16(2),
	9295:  uint16(anon_sym_LT),
	9296:  uint16(anon_sym_GT),
	9297:  uint16(175),
	9298:  uint16(3),
	9299:  uint16(anon_sym_STAR),
	9300:  uint16(anon_sym_SLASH_SLASH),
	9301:  uint16(anon_sym_PERCENT),
	9302:  uint16(193),
	9303:  uint16(4),
	9304:  uint16(anon_sym_LT_EQ),
	9305:  uint16(anon_sym_EQ_EQ),
	9306:  uint16(anon_sym_TILDE_EQ),
	9307:  uint16(anon_sym_GT_EQ),
	9308:  uint16(18),
	9309:  uint16(3),
	9310:  uint16(1),
	9311:  uint16(anon_sym_DASH_DASH),
	9312:  uint16(5),
	9313:  uint16(1),
	9314:  uint16(sym__block_comment_start),
	9315:  uint16(177),
	9316:  uint16(1),
	9317:  uint16(anon_sym_AMP),
	9318:  uint16(181),
	9319:  uint16(1),
	9320:  uint16(anon_sym_PLUS),
	9321:  uint16(183),
	9322:  uint16(1),
	9323:  uint16(anon_sym_DASH),
	9324:  uint16(185),
	9325:  uint16(1),
	9326:  uint16(anon_sym_SLASH),
	9327:  uint16(187),
	9328:  uint16(1),
	9329:  uint16(anon_sym_DOT_DOT),
	9330:  uint16(189),
	9331:  uint16(1),
	9332:  uint16(anon_sym_CARET),
	9333:  uint16(195),
	9334:  uint16(1),
	9335:  uint16(anon_sym_PIPE),
	9336:  uint16(197),
	9337:  uint16(1),
	9338:  uint16(anon_sym_TILDE),
	9339:  uint16(299),
	9340:  uint16(1),
	9341:  uint16(anon_sym_or),
	9342:  uint16(301),
	9343:  uint16(1),
	9344:  uint16(anon_sym_and),
	9345:  uint16(395),
	9346:  uint16(1),
	9347:  uint16(anon_sym_COMMA),
	9348:  uint16(124),
	9349:  uint16(1),
	9350:  uint16(sym_comment),
	9351:  uint16(179),
	9352:  uint16(2),
	9353:  uint16(anon_sym_LT_LT),
	9354:  uint16(anon_sym_GT_GT),
	9355:  uint16(191),
	9356:  uint16(2),
	9357:  uint16(anon_sym_LT),
	9358:  uint16(anon_sym_GT),
	9359:  uint16(175),
	9360:  uint16(3),
	9361:  uint16(anon_sym_STAR),
	9362:  uint16(anon_sym_SLASH_SLASH),
	9363:  uint16(anon_sym_PERCENT),
	9364:  uint16(193),
	9365:  uint16(4),
	9366:  uint16(anon_sym_LT_EQ),
	9367:  uint16(anon_sym_EQ_EQ),
	9368:  uint16(anon_sym_TILDE_EQ),
	9369:  uint16(anon_sym_GT_EQ),
	9370:  uint16(18),
	9371:  uint16(3),
	9372:  uint16(1),
	9373:  uint16(anon_sym_DASH_DASH),
	9374:  uint16(5),
	9375:  uint16(1),
	9376:  uint16(sym__block_comment_start),
	9377:  uint16(177),
	9378:  uint16(1),
	9379:  uint16(anon_sym_AMP),
	9380:  uint16(181),
	9381:  uint16(1),
	9382:  uint16(anon_sym_PLUS),
	9383:  uint16(183),
	9384:  uint16(1),
	9385:  uint16(anon_sym_DASH),
	9386:  uint16(185),
	9387:  uint16(1),
	9388:  uint16(anon_sym_SLASH),
	9389:  uint16(187),
	9390:  uint16(1),
	9391:  uint16(anon_sym_DOT_DOT),
	9392:  uint16(189),
	9393:  uint16(1),
	9394:  uint16(anon_sym_CARET),
	9395:  uint16(195),
	9396:  uint16(1),
	9397:  uint16(anon_sym_PIPE),
	9398:  uint16(197),
	9399:  uint16(1),
	9400:  uint16(anon_sym_TILDE),
	9401:  uint16(299),
	9402:  uint16(1),
	9403:  uint16(anon_sym_or),
	9404:  uint16(301),
	9405:  uint16(1),
	9406:  uint16(anon_sym_and),
	9407:  uint16(397),
	9408:  uint16(1),
	9409:  uint16(anon_sym_do),
	9410:  uint16(125),
	9411:  uint16(1),
	9412:  uint16(sym_comment),
	9413:  uint16(179),
	9414:  uint16(2),
	9415:  uint16(anon_sym_LT_LT),
	9416:  uint16(anon_sym_GT_GT),
	9417:  uint16(191),
	9418:  uint16(2),
	9419:  uint16(anon_sym_LT),
	9420:  uint16(anon_sym_GT),
	9421:  uint16(175),
	9422:  uint16(3),
	9423:  uint16(anon_sym_STAR),
	9424:  uint16(anon_sym_SLASH_SLASH),
	9425:  uint16(anon_sym_PERCENT),
	9426:  uint16(193),
	9427:  uint16(4),
	9428:  uint16(anon_sym_LT_EQ),
	9429:  uint16(anon_sym_EQ_EQ),
	9430:  uint16(anon_sym_TILDE_EQ),
	9431:  uint16(anon_sym_GT_EQ),
	9432:  uint16(18),
	9433:  uint16(3),
	9434:  uint16(1),
	9435:  uint16(anon_sym_DASH_DASH),
	9436:  uint16(5),
	9437:  uint16(1),
	9438:  uint16(sym__block_comment_start),
	9439:  uint16(177),
	9440:  uint16(1),
	9441:  uint16(anon_sym_AMP),
	9442:  uint16(181),
	9443:  uint16(1),
	9444:  uint16(anon_sym_PLUS),
	9445:  uint16(183),
	9446:  uint16(1),
	9447:  uint16(anon_sym_DASH),
	9448:  uint16(185),
	9449:  uint16(1),
	9450:  uint16(anon_sym_SLASH),
	9451:  uint16(187),
	9452:  uint16(1),
	9453:  uint16(anon_sym_DOT_DOT),
	9454:  uint16(189),
	9455:  uint16(1),
	9456:  uint16(anon_sym_CARET),
	9457:  uint16(195),
	9458:  uint16(1),
	9459:  uint16(anon_sym_PIPE),
	9460:  uint16(197),
	9461:  uint16(1),
	9462:  uint16(anon_sym_TILDE),
	9463:  uint16(299),
	9464:  uint16(1),
	9465:  uint16(anon_sym_or),
	9466:  uint16(301),
	9467:  uint16(1),
	9468:  uint16(anon_sym_and),
	9469:  uint16(399),
	9470:  uint16(1),
	9471:  uint16(anon_sym_do),
	9472:  uint16(126),
	9473:  uint16(1),
	9474:  uint16(sym_comment),
	9475:  uint16(179),
	9476:  uint16(2),
	9477:  uint16(anon_sym_LT_LT),
	9478:  uint16(anon_sym_GT_GT),
	9479:  uint16(191),
	9480:  uint16(2),
	9481:  uint16(anon_sym_LT),
	9482:  uint16(anon_sym_GT),
	9483:  uint16(175),
	9484:  uint16(3),
	9485:  uint16(anon_sym_STAR),
	9486:  uint16(anon_sym_SLASH_SLASH),
	9487:  uint16(anon_sym_PERCENT),
	9488:  uint16(193),
	9489:  uint16(4),
	9490:  uint16(anon_sym_LT_EQ),
	9491:  uint16(anon_sym_EQ_EQ),
	9492:  uint16(anon_sym_TILDE_EQ),
	9493:  uint16(anon_sym_GT_EQ),
	9494:  uint16(18),
	9495:  uint16(3),
	9496:  uint16(1),
	9497:  uint16(anon_sym_DASH_DASH),
	9498:  uint16(5),
	9499:  uint16(1),
	9500:  uint16(sym__block_comment_start),
	9501:  uint16(177),
	9502:  uint16(1),
	9503:  uint16(anon_sym_AMP),
	9504:  uint16(181),
	9505:  uint16(1),
	9506:  uint16(anon_sym_PLUS),
	9507:  uint16(183),
	9508:  uint16(1),
	9509:  uint16(anon_sym_DASH),
	9510:  uint16(185),
	9511:  uint16(1),
	9512:  uint16(anon_sym_SLASH),
	9513:  uint16(187),
	9514:  uint16(1),
	9515:  uint16(anon_sym_DOT_DOT),
	9516:  uint16(189),
	9517:  uint16(1),
	9518:  uint16(anon_sym_CARET),
	9519:  uint16(195),
	9520:  uint16(1),
	9521:  uint16(anon_sym_PIPE),
	9522:  uint16(197),
	9523:  uint16(1),
	9524:  uint16(anon_sym_TILDE),
	9525:  uint16(299),
	9526:  uint16(1),
	9527:  uint16(anon_sym_or),
	9528:  uint16(301),
	9529:  uint16(1),
	9530:  uint16(anon_sym_and),
	9531:  uint16(401),
	9532:  uint16(1),
	9533:  uint16(anon_sym_then),
	9534:  uint16(127),
	9535:  uint16(1),
	9536:  uint16(sym_comment),
	9537:  uint16(179),
	9538:  uint16(2),
	9539:  uint16(anon_sym_LT_LT),
	9540:  uint16(anon_sym_GT_GT),
	9541:  uint16(191),
	9542:  uint16(2),
	9543:  uint16(anon_sym_LT),
	9544:  uint16(anon_sym_GT),
	9545:  uint16(175),
	9546:  uint16(3),
	9547:  uint16(anon_sym_STAR),
	9548:  uint16(anon_sym_SLASH_SLASH),
	9549:  uint16(anon_sym_PERCENT),
	9550:  uint16(193),
	9551:  uint16(4),
	9552:  uint16(anon_sym_LT_EQ),
	9553:  uint16(anon_sym_EQ_EQ),
	9554:  uint16(anon_sym_TILDE_EQ),
	9555:  uint16(anon_sym_GT_EQ),
	9556:  uint16(7),
	9557:  uint16(3),
	9558:  uint16(1),
	9559:  uint16(anon_sym_DASH_DASH),
	9560:  uint16(5),
	9561:  uint16(1),
	9562:  uint16(sym__block_comment_start),
	9563:  uint16(237),
	9564:  uint16(1),
	9565:  uint16(anon_sym_COMMA),
	9566:  uint16(121),
	9567:  uint16(1),
	9568:  uint16(aux_sym__variable_assignment_explist_repeat1),
	9569:  uint16(128),
	9570:  uint16(1),
	9571:  uint16(sym_comment),
	9572:  uint16(403),
	9573:  uint16(4),
	9575:  uint16(anon_sym_SEMI),
	9576:  uint16(anon_sym_COLON_COLON),
	9577:  uint16(anon_sym_LPAREN),
	9578:  uint16(405),
	9579:  uint16(16),
	9580:  uint16(anon_sym_return),
	9581:  uint16(sym_break_statement),
	9582:  uint16(anon_sym_goto),
	9583:  uint16(anon_sym_do),
	9584:  uint16(anon_sym_end),
	9585:  uint16(anon_sym_while),
	9586:  uint16(anon_sym_repeat),
	9587:  uint16(anon_sym_until),
	9588:  uint16(anon_sym_if),
	9589:  uint16(anon_sym_elseif),
	9590:  uint16(anon_sym_else),
	9591:  uint16(anon_sym_for),
	9592:  uint16(anon_sym_function),
	9593:  uint16(anon_sym_local),
	9594:  uint16(anon_sym_global),
	9595:  uint16(sym_identifier),
	9596:  uint16(6),
	9597:  uint16(3),
	9598:  uint16(1),
	9599:  uint16(anon_sym_DASH_DASH),
	9600:  uint16(5),
	9601:  uint16(1),
	9602:  uint16(sym__block_comment_start),
	9603:  uint16(411),
	9604:  uint16(1),
	9605:  uint16(anon_sym_EQ),
	9606:  uint16(129),
	9607:  uint16(1),
	9608:  uint16(sym_comment),
	9609:  uint16(407),
	9610:  uint16(4),
	9612:  uint16(anon_sym_SEMI),
	9613:  uint16(anon_sym_COLON_COLON),
	9614:  uint16(anon_sym_LPAREN),
	9615:  uint16(409),
	9616:  uint16(16),
	9617:  uint16(anon_sym_return),
	9618:  uint16(sym_break_statement),
	9619:  uint16(anon_sym_goto),
	9620:  uint16(anon_sym_do),
	9621:  uint16(anon_sym_end),
	9622:  uint16(anon_sym_while),
	9623:  uint16(anon_sym_repeat),
	9624:  uint16(anon_sym_until),
	9625:  uint16(anon_sym_if),
	9626:  uint16(anon_sym_elseif),
	9627:  uint16(anon_sym_else),
	9628:  uint16(anon_sym_for),
	9629:  uint16(anon_sym_function),
	9630:  uint16(anon_sym_local),
	9631:  uint16(anon_sym_global),
	9632:  uint16(sym_identifier),
	9633:  uint16(6),
	9634:  uint16(3),
	9635:  uint16(1),
	9636:  uint16(anon_sym_DASH_DASH),
	9637:  uint16(5),
	9638:  uint16(1),
	9639:  uint16(sym__block_comment_start),
	9640:  uint16(411),
	9641:  uint16(1),
	9642:  uint16(anon_sym_EQ),
	9643:  uint16(130),
	9644:  uint16(1),
	9645:  uint16(sym_comment),
	9646:  uint16(413),
	9647:  uint16(4),
	9649:  uint16(anon_sym_SEMI),
	9650:  uint16(anon_sym_COLON_COLON),
	9651:  uint16(anon_sym_LPAREN),
	9652:  uint16(415),
	9653:  uint16(16),
	9654:  uint16(anon_sym_return),
	9655:  uint16(sym_break_statement),
	9656:  uint16(anon_sym_goto),
	9657:  uint16(anon_sym_do),
	9658:  uint16(anon_sym_end),
	9659:  uint16(anon_sym_while),
	9660:  uint16(anon_sym_repeat),
	9661:  uint16(anon_sym_until),
	9662:  uint16(anon_sym_if),
	9663:  uint16(anon_sym_elseif),
	9664:  uint16(anon_sym_else),
	9665:  uint16(anon_sym_for),
	9666:  uint16(anon_sym_function),
	9667:  uint16(anon_sym_local),
	9668:  uint16(anon_sym_global),
	9669:  uint16(sym_identifier),
	9670:  uint16(5),
	9671:  uint16(3),
	9672:  uint16(1),
	9673:  uint16(anon_sym_DASH_DASH),
	9674:  uint16(5),
	9675:  uint16(1),
	9676:  uint16(sym__block_comment_start),
	9677:  uint16(131),
	9678:  uint16(1),
	9679:  uint16(sym_comment),
	9680:  uint16(417),
	9681:  uint16(4),
	9683:  uint16(anon_sym_SEMI),
	9684:  uint16(anon_sym_COLON_COLON),
	9685:  uint16(anon_sym_LPAREN),
	9686:  uint16(419),
	9687:  uint16(16),
	9688:  uint16(anon_sym_return),
	9689:  uint16(sym_break_statement),
	9690:  uint16(anon_sym_goto),
	9691:  uint16(anon_sym_do),
	9692:  uint16(anon_sym_end),
	9693:  uint16(anon_sym_while),
	9694:  uint16(anon_sym_repeat),
	9695:  uint16(anon_sym_until),
	9696:  uint16(anon_sym_if),
	9697:  uint16(anon_sym_elseif),
	9698:  uint16(anon_sym_else),
	9699:  uint16(anon_sym_for),
	9700:  uint16(anon_sym_function),
	9701:  uint16(anon_sym_local),
	9702:  uint16(anon_sym_global),
	9703:  uint16(sym_identifier),
	9704:  uint16(5),
	9705:  uint16(3),
	9706:  uint16(1),
	9707:  uint16(anon_sym_DASH_DASH),
	9708:  uint16(5),
	9709:  uint16(1),
	9710:  uint16(sym__block_comment_start),
	9711:  uint16(132),
	9712:  uint16(1),
	9713:  uint16(sym_comment),
	9714:  uint16(421),
	9715:  uint16(4),
	9717:  uint16(anon_sym_SEMI),
	9718:  uint16(anon_sym_COLON_COLON),
	9719:  uint16(anon_sym_LPAREN),
	9720:  uint16(423),
	9721:  uint16(16),
	9722:  uint16(anon_sym_return),
	9723:  uint16(sym_break_statement),
	9724:  uint16(anon_sym_goto),
	9725:  uint16(anon_sym_do),
	9726:  uint16(anon_sym_end),
	9727:  uint16(anon_sym_while),
	9728:  uint16(anon_sym_repeat),
	9729:  uint16(anon_sym_until),
	9730:  uint16(anon_sym_if),
	9731:  uint16(anon_sym_elseif),
	9732:  uint16(anon_sym_else),
	9733:  uint16(anon_sym_for),
	9734:  uint16(anon_sym_function),
	9735:  uint16(anon_sym_local),
	9736:  uint16(anon_sym_global),
	9737:  uint16(sym_identifier),
	9738:  uint16(5),
	9739:  uint16(3),
	9740:  uint16(1),
	9741:  uint16(anon_sym_DASH_DASH),
	9742:  uint16(5),
	9743:  uint16(1),
	9744:  uint16(sym__block_comment_start),
	9745:  uint16(133),
	9746:  uint16(1),
	9747:  uint16(sym_comment),
	9748:  uint16(425),
	9749:  uint16(4),
	9751:  uint16(anon_sym_SEMI),
	9752:  uint16(anon_sym_COLON_COLON),
	9753:  uint16(anon_sym_LPAREN),
	9754:  uint16(427),
	9755:  uint16(16),
	9756:  uint16(anon_sym_return),
	9757:  uint16(sym_break_statement),
	9758:  uint16(anon_sym_goto),
	9759:  uint16(anon_sym_do),
	9760:  uint16(anon_sym_end),
	9761:  uint16(anon_sym_while),
	9762:  uint16(anon_sym_repeat),
	9763:  uint16(anon_sym_until),
	9764:  uint16(anon_sym_if),
	9765:  uint16(anon_sym_elseif),
	9766:  uint16(anon_sym_else),
	9767:  uint16(anon_sym_for),
	9768:  uint16(anon_sym_function),
	9769:  uint16(anon_sym_local),
	9770:  uint16(anon_sym_global),
	9771:  uint16(sym_identifier),
	9772:  uint16(5),
	9773:  uint16(3),
	9774:  uint16(1),
	9775:  uint16(anon_sym_DASH_DASH),
	9776:  uint16(5),
	9777:  uint16(1),
	9778:  uint16(sym__block_comment_start),
	9779:  uint16(134),
	9780:  uint16(1),
	9781:  uint16(sym_comment),
	9782:  uint16(429),
	9783:  uint16(4),
	9785:  uint16(anon_sym_SEMI),
	9786:  uint16(anon_sym_COLON_COLON),
	9787:  uint16(anon_sym_LPAREN),
	9788:  uint16(431),
	9789:  uint16(16),
	9790:  uint16(anon_sym_return),
	9791:  uint16(sym_break_statement),
	9792:  uint16(anon_sym_goto),
	9793:  uint16(anon_sym_do),
	9794:  uint16(anon_sym_end),
	9795:  uint16(anon_sym_while),
	9796:  uint16(anon_sym_repeat),
	9797:  uint16(anon_sym_until),
	9798:  uint16(anon_sym_if),
	9799:  uint16(anon_sym_elseif),
	9800:  uint16(anon_sym_else),
	9801:  uint16(anon_sym_for),
	9802:  uint16(anon_sym_function),
	9803:  uint16(anon_sym_local),
	9804:  uint16(anon_sym_global),
	9805:  uint16(sym_identifier),
	9806:  uint16(5),
	9807:  uint16(3),
	9808:  uint16(1),
	9809:  uint16(anon_sym_DASH_DASH),
	9810:  uint16(5),
	9811:  uint16(1),
	9812:  uint16(sym__block_comment_start),
	9813:  uint16(135),
	9814:  uint16(1),
	9815:  uint16(sym_comment),
	9816:  uint16(433),
	9817:  uint16(4),
	9819:  uint16(anon_sym_SEMI),
	9820:  uint16(anon_sym_COLON_COLON),
	9821:  uint16(anon_sym_LPAREN),
	9822:  uint16(435),
	9823:  uint16(16),
	9824:  uint16(anon_sym_return),
	9825:  uint16(sym_break_statement),
	9826:  uint16(anon_sym_goto),
	9827:  uint16(anon_sym_do),
	9828:  uint16(anon_sym_end),
	9829:  uint16(anon_sym_while),
	9830:  uint16(anon_sym_repeat),
	9831:  uint16(anon_sym_until),
	9832:  uint16(anon_sym_if),
	9833:  uint16(anon_sym_elseif),
	9834:  uint16(anon_sym_else),
	9835:  uint16(anon_sym_for),
	9836:  uint16(anon_sym_function),
	9837:  uint16(anon_sym_local),
	9838:  uint16(anon_sym_global),
	9839:  uint16(sym_identifier),
	9840:  uint16(5),
	9841:  uint16(3),
	9842:  uint16(1),
	9843:  uint16(anon_sym_DASH_DASH),
	9844:  uint16(5),
	9845:  uint16(1),
	9846:  uint16(sym__block_comment_start),
	9847:  uint16(136),
	9848:  uint16(1),
	9849:  uint16(sym_comment),
	9850:  uint16(437),
	9851:  uint16(4),
	9853:  uint16(anon_sym_SEMI),
	9854:  uint16(anon_sym_COLON_COLON),
	9855:  uint16(anon_sym_LPAREN),
	9856:  uint16(439),
	9857:  uint16(16),
	9858:  uint16(anon_sym_return),
	9859:  uint16(sym_break_statement),
	9860:  uint16(anon_sym_goto),
	9861:  uint16(anon_sym_do),
	9862:  uint16(anon_sym_end),
	9863:  uint16(anon_sym_while),
	9864:  uint16(anon_sym_repeat),
	9865:  uint16(anon_sym_until),
	9866:  uint16(anon_sym_if),
	9867:  uint16(anon_sym_elseif),
	9868:  uint16(anon_sym_else),
	9869:  uint16(anon_sym_for),
	9870:  uint16(anon_sym_function),
	9871:  uint16(anon_sym_local),
	9872:  uint16(anon_sym_global),
	9873:  uint16(sym_identifier),
	9874:  uint16(5),
	9875:  uint16(3),
	9876:  uint16(1),
	9877:  uint16(anon_sym_DASH_DASH),
	9878:  uint16(5),
	9879:  uint16(1),
	9880:  uint16(sym__block_comment_start),
	9881:  uint16(137),
	9882:  uint16(1),
	9883:  uint16(sym_comment),
	9884:  uint16(441),
	9885:  uint16(4),
	9887:  uint16(anon_sym_SEMI),
	9888:  uint16(anon_sym_COLON_COLON),
	9889:  uint16(anon_sym_LPAREN),
	9890:  uint16(443),
	9891:  uint16(16),
	9892:  uint16(anon_sym_return),
	9893:  uint16(sym_break_statement),
	9894:  uint16(anon_sym_goto),
	9895:  uint16(anon_sym_do),
	9896:  uint16(anon_sym_end),
	9897:  uint16(anon_sym_while),
	9898:  uint16(anon_sym_repeat),
	9899:  uint16(anon_sym_until),
	9900:  uint16(anon_sym_if),
	9901:  uint16(anon_sym_elseif),
	9902:  uint16(anon_sym_else),
	9903:  uint16(anon_sym_for),
	9904:  uint16(anon_sym_function),
	9905:  uint16(anon_sym_local),
	9906:  uint16(anon_sym_global),
	9907:  uint16(sym_identifier),
	9908:  uint16(5),
	9909:  uint16(3),
	9910:  uint16(1),
	9911:  uint16(anon_sym_DASH_DASH),
	9912:  uint16(5),
	9913:  uint16(1),
	9914:  uint16(sym__block_comment_start),
	9915:  uint16(138),
	9916:  uint16(1),
	9917:  uint16(sym_comment),
	9918:  uint16(445),
	9919:  uint16(4),
	9921:  uint16(anon_sym_SEMI),
	9922:  uint16(anon_sym_COLON_COLON),
	9923:  uint16(anon_sym_LPAREN),
	9924:  uint16(447),
	9925:  uint16(16),
	9926:  uint16(anon_sym_return),
	9927:  uint16(sym_break_statement),
	9928:  uint16(anon_sym_goto),
	9929:  uint16(anon_sym_do),
	9930:  uint16(anon_sym_end),
	9931:  uint16(anon_sym_while),
	9932:  uint16(anon_sym_repeat),
	9933:  uint16(anon_sym_until),
	9934:  uint16(anon_sym_if),
	9935:  uint16(anon_sym_elseif),
	9936:  uint16(anon_sym_else),
	9937:  uint16(anon_sym_for),
	9938:  uint16(anon_sym_function),
	9939:  uint16(anon_sym_local),
	9940:  uint16(anon_sym_global),
	9941:  uint16(sym_identifier),
	9942:  uint16(5),
	9943:  uint16(3),
	9944:  uint16(1),
	9945:  uint16(anon_sym_DASH_DASH),
	9946:  uint16(5),
	9947:  uint16(1),
	9948:  uint16(sym__block_comment_start),
	9949:  uint16(139),
	9950:  uint16(1),
	9951:  uint16(sym_comment),
	9952:  uint16(449),
	9953:  uint16(4),
	9955:  uint16(anon_sym_SEMI),
	9956:  uint16(anon_sym_COLON_COLON),
	9957:  uint16(anon_sym_LPAREN),
	9958:  uint16(451),
	9959:  uint16(16),
	9960:  uint16(anon_sym_return),
	9961:  uint16(sym_break_statement),
	9962:  uint16(anon_sym_goto),
	9963:  uint16(anon_sym_do),
	9964:  uint16(anon_sym_end),
	9965:  uint16(anon_sym_while),
	9966:  uint16(anon_sym_repeat),
	9967:  uint16(anon_sym_until),
	9968:  uint16(anon_sym_if),
	9969:  uint16(anon_sym_elseif),
	9970:  uint16(anon_sym_else),
	9971:  uint16(anon_sym_for),
	9972:  uint16(anon_sym_function),
	9973:  uint16(anon_sym_local),
	9974:  uint16(anon_sym_global),
	9975:  uint16(sym_identifier),
	9976:  uint16(5),
	9977:  uint16(3),
	9978:  uint16(1),
	9979:  uint16(anon_sym_DASH_DASH),
	9980:  uint16(5),
	9981:  uint16(1),
	9982:  uint16(sym__block_comment_start),
	9983:  uint16(140),
	9984:  uint16(1),
	9985:  uint16(sym_comment),
	9986:  uint16(453),
	9987:  uint16(4),
	9989:  uint16(anon_sym_SEMI),
	9990:  uint16(anon_sym_COLON_COLON),
	9991:  uint16(anon_sym_LPAREN),
	9992:  uint16(455),
	9993:  uint16(16),
	9994:  uint16(anon_sym_return),
	9995:  uint16(sym_break_statement),
	9996:  uint16(anon_sym_goto),
	9997:  uint16(anon_sym_do),
	9998:  uint16(anon_sym_end),
	9999:  uint16(anon_sym_while),
	10000: uint16(anon_sym_repeat),
	10001: uint16(anon_sym_until),
	10002: uint16(anon_sym_if),
	10003: uint16(anon_sym_elseif),
	10004: uint16(anon_sym_else),
	10005: uint16(anon_sym_for),
	10006: uint16(anon_sym_function),
	10007: uint16(anon_sym_local),
	10008: uint16(anon_sym_global),
	10009: uint16(sym_identifier),
	10010: uint16(5),
	10011: uint16(3),
	10012: uint16(1),
	10013: uint16(anon_sym_DASH_DASH),
	10014: uint16(5),
	10015: uint16(1),
	10016: uint16(sym__block_comment_start),
	10017: uint16(141),
	10018: uint16(1),
	10019: uint16(sym_comment),
	10020: uint16(457),
	10021: uint16(4),
	10023: uint16(anon_sym_SEMI),
	10024: uint16(anon_sym_COLON_COLON),
	10025: uint16(anon_sym_LPAREN),
	10026: uint16(459),
	10027: uint16(16),
	10028: uint16(anon_sym_return),
	10029: uint16(sym_break_statement),
	10030: uint16(anon_sym_goto),
	10031: uint16(anon_sym_do),
	10032: uint16(anon_sym_end),
	10033: uint16(anon_sym_while),
	10034: uint16(anon_sym_repeat),
	10035: uint16(anon_sym_until),
	10036: uint16(anon_sym_if),
	10037: uint16(anon_sym_elseif),
	10038: uint16(anon_sym_else),
	10039: uint16(anon_sym_for),
	10040: uint16(anon_sym_function),
	10041: uint16(anon_sym_local),
	10042: uint16(anon_sym_global),
	10043: uint16(sym_identifier),
	10044: uint16(5),
	10045: uint16(3),
	10046: uint16(1),
	10047: uint16(anon_sym_DASH_DASH),
	10048: uint16(5),
	10049: uint16(1),
	10050: uint16(sym__block_comment_start),
	10051: uint16(142),
	10052: uint16(1),
	10053: uint16(sym_comment),
	10054: uint16(461),
	10055: uint16(4),
	10057: uint16(anon_sym_SEMI),
	10058: uint16(anon_sym_COLON_COLON),
	10059: uint16(anon_sym_LPAREN),
	10060: uint16(463),
	10061: uint16(16),
	10062: uint16(anon_sym_return),
	10063: uint16(sym_break_statement),
	10064: uint16(anon_sym_goto),
	10065: uint16(anon_sym_do),
	10066: uint16(anon_sym_end),
	10067: uint16(anon_sym_while),
	10068: uint16(anon_sym_repeat),
	10069: uint16(anon_sym_until),
	10070: uint16(anon_sym_if),
	10071: uint16(anon_sym_elseif),
	10072: uint16(anon_sym_else),
	10073: uint16(anon_sym_for),
	10074: uint16(anon_sym_function),
	10075: uint16(anon_sym_local),
	10076: uint16(anon_sym_global),
	10077: uint16(sym_identifier),
	10078: uint16(5),
	10079: uint16(3),
	10080: uint16(1),
	10081: uint16(anon_sym_DASH_DASH),
	10082: uint16(5),
	10083: uint16(1),
	10084: uint16(sym__block_comment_start),
	10085: uint16(143),
	10086: uint16(1),
	10087: uint16(sym_comment),
	10088: uint16(465),
	10089: uint16(4),
	10091: uint16(anon_sym_SEMI),
	10092: uint16(anon_sym_COLON_COLON),
	10093: uint16(anon_sym_LPAREN),
	10094: uint16(467),
	10095: uint16(16),
	10096: uint16(anon_sym_return),
	10097: uint16(sym_break_statement),
	10098: uint16(anon_sym_goto),
	10099: uint16(anon_sym_do),
	10100: uint16(anon_sym_end),
	10101: uint16(anon_sym_while),
	10102: uint16(anon_sym_repeat),
	10103: uint16(anon_sym_until),
	10104: uint16(anon_sym_if),
	10105: uint16(anon_sym_elseif),
	10106: uint16(anon_sym_else),
	10107: uint16(anon_sym_for),
	10108: uint16(anon_sym_function),
	10109: uint16(anon_sym_local),
	10110: uint16(anon_sym_global),
	10111: uint16(sym_identifier),
	10112: uint16(5),
	10113: uint16(3),
	10114: uint16(1),
	10115: uint16(anon_sym_DASH_DASH),
	10116: uint16(5),
	10117: uint16(1),
	10118: uint16(sym__block_comment_start),
	10119: uint16(144),
	10120: uint16(1),
	10121: uint16(sym_comment),
	10122: uint16(469),
	10123: uint16(4),
	10125: uint16(anon_sym_SEMI),
	10126: uint16(anon_sym_COLON_COLON),
	10127: uint16(anon_sym_LPAREN),
	10128: uint16(471),
	10129: uint16(16),
	10130: uint16(anon_sym_return),
	10131: uint16(sym_break_statement),
	10132: uint16(anon_sym_goto),
	10133: uint16(anon_sym_do),
	10134: uint16(anon_sym_end),
	10135: uint16(anon_sym_while),
	10136: uint16(anon_sym_repeat),
	10137: uint16(anon_sym_until),
	10138: uint16(anon_sym_if),
	10139: uint16(anon_sym_elseif),
	10140: uint16(anon_sym_else),
	10141: uint16(anon_sym_for),
	10142: uint16(anon_sym_function),
	10143: uint16(anon_sym_local),
	10144: uint16(anon_sym_global),
	10145: uint16(sym_identifier),
	10146: uint16(5),
	10147: uint16(3),
	10148: uint16(1),
	10149: uint16(anon_sym_DASH_DASH),
	10150: uint16(5),
	10151: uint16(1),
	10152: uint16(sym__block_comment_start),
	10153: uint16(145),
	10154: uint16(1),
	10155: uint16(sym_comment),
	10156: uint16(473),
	10157: uint16(4),
	10159: uint16(anon_sym_SEMI),
	10160: uint16(anon_sym_COLON_COLON),
	10161: uint16(anon_sym_LPAREN),
	10162: uint16(475),
	10163: uint16(16),
	10164: uint16(anon_sym_return),
	10165: uint16(sym_break_statement),
	10166: uint16(anon_sym_goto),
	10167: uint16(anon_sym_do),
	10168: uint16(anon_sym_end),
	10169: uint16(anon_sym_while),
	10170: uint16(anon_sym_repeat),
	10171: uint16(anon_sym_until),
	10172: uint16(anon_sym_if),
	10173: uint16(anon_sym_elseif),
	10174: uint16(anon_sym_else),
	10175: uint16(anon_sym_for),
	10176: uint16(anon_sym_function),
	10177: uint16(anon_sym_local),
	10178: uint16(anon_sym_global),
	10179: uint16(sym_identifier),
	10180: uint16(5),
	10181: uint16(3),
	10182: uint16(1),
	10183: uint16(anon_sym_DASH_DASH),
	10184: uint16(5),
	10185: uint16(1),
	10186: uint16(sym__block_comment_start),
	10187: uint16(146),
	10188: uint16(1),
	10189: uint16(sym_comment),
	10190: uint16(477),
	10191: uint16(4),
	10193: uint16(anon_sym_SEMI),
	10194: uint16(anon_sym_COLON_COLON),
	10195: uint16(anon_sym_LPAREN),
	10196: uint16(479),
	10197: uint16(16),
	10198: uint16(anon_sym_return),
	10199: uint16(sym_break_statement),
	10200: uint16(anon_sym_goto),
	10201: uint16(anon_sym_do),
	10202: uint16(anon_sym_end),
	10203: uint16(anon_sym_while),
	10204: uint16(anon_sym_repeat),
	10205: uint16(anon_sym_until),
	10206: uint16(anon_sym_if),
	10207: uint16(anon_sym_elseif),
	10208: uint16(anon_sym_else),
	10209: uint16(anon_sym_for),
	10210: uint16(anon_sym_function),
	10211: uint16(anon_sym_local),
	10212: uint16(anon_sym_global),
	10213: uint16(sym_identifier),
	10214: uint16(5),
	10215: uint16(3),
	10216: uint16(1),
	10217: uint16(anon_sym_DASH_DASH),
	10218: uint16(5),
	10219: uint16(1),
	10220: uint16(sym__block_comment_start),
	10221: uint16(147),
	10222: uint16(1),
	10223: uint16(sym_comment),
	10224: uint16(481),
	10225: uint16(4),
	10227: uint16(anon_sym_SEMI),
	10228: uint16(anon_sym_COLON_COLON),
	10229: uint16(anon_sym_LPAREN),
	10230: uint16(483),
	10231: uint16(16),
	10232: uint16(anon_sym_return),
	10233: uint16(sym_break_statement),
	10234: uint16(anon_sym_goto),
	10235: uint16(anon_sym_do),
	10236: uint16(anon_sym_end),
	10237: uint16(anon_sym_while),
	10238: uint16(anon_sym_repeat),
	10239: uint16(anon_sym_until),
	10240: uint16(anon_sym_if),
	10241: uint16(anon_sym_elseif),
	10242: uint16(anon_sym_else),
	10243: uint16(anon_sym_for),
	10244: uint16(anon_sym_function),
	10245: uint16(anon_sym_local),
	10246: uint16(anon_sym_global),
	10247: uint16(sym_identifier),
	10248: uint16(5),
	10249: uint16(3),
	10250: uint16(1),
	10251: uint16(anon_sym_DASH_DASH),
	10252: uint16(5),
	10253: uint16(1),
	10254: uint16(sym__block_comment_start),
	10255: uint16(148),
	10256: uint16(1),
	10257: uint16(sym_comment),
	10258: uint16(485),
	10259: uint16(4),
	10261: uint16(anon_sym_SEMI),
	10262: uint16(anon_sym_COLON_COLON),
	10263: uint16(anon_sym_LPAREN),
	10264: uint16(487),
	10265: uint16(16),
	10266: uint16(anon_sym_return),
	10267: uint16(sym_break_statement),
	10268: uint16(anon_sym_goto),
	10269: uint16(anon_sym_do),
	10270: uint16(anon_sym_end),
	10271: uint16(anon_sym_while),
	10272: uint16(anon_sym_repeat),
	10273: uint16(anon_sym_until),
	10274: uint16(anon_sym_if),
	10275: uint16(anon_sym_elseif),
	10276: uint16(anon_sym_else),
	10277: uint16(anon_sym_for),
	10278: uint16(anon_sym_function),
	10279: uint16(anon_sym_local),
	10280: uint16(anon_sym_global),
	10281: uint16(sym_identifier),
	10282: uint16(5),
	10283: uint16(3),
	10284: uint16(1),
	10285: uint16(anon_sym_DASH_DASH),
	10286: uint16(5),
	10287: uint16(1),
	10288: uint16(sym__block_comment_start),
	10289: uint16(149),
	10290: uint16(1),
	10291: uint16(sym_comment),
	10292: uint16(489),
	10293: uint16(4),
	10295: uint16(anon_sym_SEMI),
	10296: uint16(anon_sym_COLON_COLON),
	10297: uint16(anon_sym_LPAREN),
	10298: uint16(491),
	10299: uint16(16),
	10300: uint16(anon_sym_return),
	10301: uint16(sym_break_statement),
	10302: uint16(anon_sym_goto),
	10303: uint16(anon_sym_do),
	10304: uint16(anon_sym_end),
	10305: uint16(anon_sym_while),
	10306: uint16(anon_sym_repeat),
	10307: uint16(anon_sym_until),
	10308: uint16(anon_sym_if),
	10309: uint16(anon_sym_elseif),
	10310: uint16(anon_sym_else),
	10311: uint16(anon_sym_for),
	10312: uint16(anon_sym_function),
	10313: uint16(anon_sym_local),
	10314: uint16(anon_sym_global),
	10315: uint16(sym_identifier),
	10316: uint16(5),
	10317: uint16(3),
	10318: uint16(1),
	10319: uint16(anon_sym_DASH_DASH),
	10320: uint16(5),
	10321: uint16(1),
	10322: uint16(sym__block_comment_start),
	10323: uint16(150),
	10324: uint16(1),
	10325: uint16(sym_comment),
	10326: uint16(493),
	10327: uint16(4),
	10329: uint16(anon_sym_SEMI),
	10330: uint16(anon_sym_COLON_COLON),
	10331: uint16(anon_sym_LPAREN),
	10332: uint16(495),
	10333: uint16(16),
	10334: uint16(anon_sym_return),
	10335: uint16(sym_break_statement),
	10336: uint16(anon_sym_goto),
	10337: uint16(anon_sym_do),
	10338: uint16(anon_sym_end),
	10339: uint16(anon_sym_while),
	10340: uint16(anon_sym_repeat),
	10341: uint16(anon_sym_until),
	10342: uint16(anon_sym_if),
	10343: uint16(anon_sym_elseif),
	10344: uint16(anon_sym_else),
	10345: uint16(anon_sym_for),
	10346: uint16(anon_sym_function),
	10347: uint16(anon_sym_local),
	10348: uint16(anon_sym_global),
	10349: uint16(sym_identifier),
	10350: uint16(5),
	10351: uint16(3),
	10352: uint16(1),
	10353: uint16(anon_sym_DASH_DASH),
	10354: uint16(5),
	10355: uint16(1),
	10356: uint16(sym__block_comment_start),
	10357: uint16(151),
	10358: uint16(1),
	10359: uint16(sym_comment),
	10360: uint16(497),
	10361: uint16(4),
	10363: uint16(anon_sym_SEMI),
	10364: uint16(anon_sym_COLON_COLON),
	10365: uint16(anon_sym_LPAREN),
	10366: uint16(499),
	10367: uint16(16),
	10368: uint16(anon_sym_return),
	10369: uint16(sym_break_statement),
	10370: uint16(anon_sym_goto),
	10371: uint16(anon_sym_do),
	10372: uint16(anon_sym_end),
	10373: uint16(anon_sym_while),
	10374: uint16(anon_sym_repeat),
	10375: uint16(anon_sym_until),
	10376: uint16(anon_sym_if),
	10377: uint16(anon_sym_elseif),
	10378: uint16(anon_sym_else),
	10379: uint16(anon_sym_for),
	10380: uint16(anon_sym_function),
	10381: uint16(anon_sym_local),
	10382: uint16(anon_sym_global),
	10383: uint16(sym_identifier),
	10384: uint16(5),
	10385: uint16(3),
	10386: uint16(1),
	10387: uint16(anon_sym_DASH_DASH),
	10388: uint16(5),
	10389: uint16(1),
	10390: uint16(sym__block_comment_start),
	10391: uint16(152),
	10392: uint16(1),
	10393: uint16(sym_comment),
	10394: uint16(501),
	10395: uint16(4),
	10397: uint16(anon_sym_SEMI),
	10398: uint16(anon_sym_COLON_COLON),
	10399: uint16(anon_sym_LPAREN),
	10400: uint16(503),
	10401: uint16(16),
	10402: uint16(anon_sym_return),
	10403: uint16(sym_break_statement),
	10404: uint16(anon_sym_goto),
	10405: uint16(anon_sym_do),
	10406: uint16(anon_sym_end),
	10407: uint16(anon_sym_while),
	10408: uint16(anon_sym_repeat),
	10409: uint16(anon_sym_until),
	10410: uint16(anon_sym_if),
	10411: uint16(anon_sym_elseif),
	10412: uint16(anon_sym_else),
	10413: uint16(anon_sym_for),
	10414: uint16(anon_sym_function),
	10415: uint16(anon_sym_local),
	10416: uint16(anon_sym_global),
	10417: uint16(sym_identifier),
	10418: uint16(5),
	10419: uint16(3),
	10420: uint16(1),
	10421: uint16(anon_sym_DASH_DASH),
	10422: uint16(5),
	10423: uint16(1),
	10424: uint16(sym__block_comment_start),
	10425: uint16(153),
	10426: uint16(1),
	10427: uint16(sym_comment),
	10428: uint16(505),
	10429: uint16(4),
	10431: uint16(anon_sym_SEMI),
	10432: uint16(anon_sym_COLON_COLON),
	10433: uint16(anon_sym_LPAREN),
	10434: uint16(507),
	10435: uint16(16),
	10436: uint16(anon_sym_return),
	10437: uint16(sym_break_statement),
	10438: uint16(anon_sym_goto),
	10439: uint16(anon_sym_do),
	10440: uint16(anon_sym_end),
	10441: uint16(anon_sym_while),
	10442: uint16(anon_sym_repeat),
	10443: uint16(anon_sym_until),
	10444: uint16(anon_sym_if),
	10445: uint16(anon_sym_elseif),
	10446: uint16(anon_sym_else),
	10447: uint16(anon_sym_for),
	10448: uint16(anon_sym_function),
	10449: uint16(anon_sym_local),
	10450: uint16(anon_sym_global),
	10451: uint16(sym_identifier),
	10452: uint16(5),
	10453: uint16(3),
	10454: uint16(1),
	10455: uint16(anon_sym_DASH_DASH),
	10456: uint16(5),
	10457: uint16(1),
	10458: uint16(sym__block_comment_start),
	10459: uint16(154),
	10460: uint16(1),
	10461: uint16(sym_comment),
	10462: uint16(509),
	10463: uint16(4),
	10465: uint16(anon_sym_SEMI),
	10466: uint16(anon_sym_COLON_COLON),
	10467: uint16(anon_sym_LPAREN),
	10468: uint16(511),
	10469: uint16(16),
	10470: uint16(anon_sym_return),
	10471: uint16(sym_break_statement),
	10472: uint16(anon_sym_goto),
	10473: uint16(anon_sym_do),
	10474: uint16(anon_sym_end),
	10475: uint16(anon_sym_while),
	10476: uint16(anon_sym_repeat),
	10477: uint16(anon_sym_until),
	10478: uint16(anon_sym_if),
	10479: uint16(anon_sym_elseif),
	10480: uint16(anon_sym_else),
	10481: uint16(anon_sym_for),
	10482: uint16(anon_sym_function),
	10483: uint16(anon_sym_local),
	10484: uint16(anon_sym_global),
	10485: uint16(sym_identifier),
	10486: uint16(5),
	10487: uint16(3),
	10488: uint16(1),
	10489: uint16(anon_sym_DASH_DASH),
	10490: uint16(5),
	10491: uint16(1),
	10492: uint16(sym__block_comment_start),
	10493: uint16(155),
	10494: uint16(1),
	10495: uint16(sym_comment),
	10496: uint16(513),
	10497: uint16(4),
	10499: uint16(anon_sym_SEMI),
	10500: uint16(anon_sym_COLON_COLON),
	10501: uint16(anon_sym_LPAREN),
	10502: uint16(515),
	10503: uint16(16),
	10504: uint16(anon_sym_return),
	10505: uint16(sym_break_statement),
	10506: uint16(anon_sym_goto),
	10507: uint16(anon_sym_do),
	10508: uint16(anon_sym_end),
	10509: uint16(anon_sym_while),
	10510: uint16(anon_sym_repeat),
	10511: uint16(anon_sym_until),
	10512: uint16(anon_sym_if),
	10513: uint16(anon_sym_elseif),
	10514: uint16(anon_sym_else),
	10515: uint16(anon_sym_for),
	10516: uint16(anon_sym_function),
	10517: uint16(anon_sym_local),
	10518: uint16(anon_sym_global),
	10519: uint16(sym_identifier),
	10520: uint16(5),
	10521: uint16(3),
	10522: uint16(1),
	10523: uint16(anon_sym_DASH_DASH),
	10524: uint16(5),
	10525: uint16(1),
	10526: uint16(sym__block_comment_start),
	10527: uint16(156),
	10528: uint16(1),
	10529: uint16(sym_comment),
	10530: uint16(517),
	10531: uint16(4),
	10533: uint16(anon_sym_SEMI),
	10534: uint16(anon_sym_COLON_COLON),
	10535: uint16(anon_sym_LPAREN),
	10536: uint16(519),
	10537: uint16(16),
	10538: uint16(anon_sym_return),
	10539: uint16(sym_break_statement),
	10540: uint16(anon_sym_goto),
	10541: uint16(anon_sym_do),
	10542: uint16(anon_sym_end),
	10543: uint16(anon_sym_while),
	10544: uint16(anon_sym_repeat),
	10545: uint16(anon_sym_until),
	10546: uint16(anon_sym_if),
	10547: uint16(anon_sym_elseif),
	10548: uint16(anon_sym_else),
	10549: uint16(anon_sym_for),
	10550: uint16(anon_sym_function),
	10551: uint16(anon_sym_local),
	10552: uint16(anon_sym_global),
	10553: uint16(sym_identifier),
	10554: uint16(5),
	10555: uint16(3),
	10556: uint16(1),
	10557: uint16(anon_sym_DASH_DASH),
	10558: uint16(5),
	10559: uint16(1),
	10560: uint16(sym__block_comment_start),
	10561: uint16(157),
	10562: uint16(1),
	10563: uint16(sym_comment),
	10564: uint16(521),
	10565: uint16(4),
	10567: uint16(anon_sym_SEMI),
	10568: uint16(anon_sym_COLON_COLON),
	10569: uint16(anon_sym_LPAREN),
	10570: uint16(523),
	10571: uint16(16),
	10572: uint16(anon_sym_return),
	10573: uint16(sym_break_statement),
	10574: uint16(anon_sym_goto),
	10575: uint16(anon_sym_do),
	10576: uint16(anon_sym_end),
	10577: uint16(anon_sym_while),
	10578: uint16(anon_sym_repeat),
	10579: uint16(anon_sym_until),
	10580: uint16(anon_sym_if),
	10581: uint16(anon_sym_elseif),
	10582: uint16(anon_sym_else),
	10583: uint16(anon_sym_for),
	10584: uint16(anon_sym_function),
	10585: uint16(anon_sym_local),
	10586: uint16(anon_sym_global),
	10587: uint16(sym_identifier),
	10588: uint16(5),
	10589: uint16(3),
	10590: uint16(1),
	10591: uint16(anon_sym_DASH_DASH),
	10592: uint16(5),
	10593: uint16(1),
	10594: uint16(sym__block_comment_start),
	10595: uint16(158),
	10596: uint16(1),
	10597: uint16(sym_comment),
	10598: uint16(525),
	10599: uint16(4),
	10601: uint16(anon_sym_SEMI),
	10602: uint16(anon_sym_COLON_COLON),
	10603: uint16(anon_sym_LPAREN),
	10604: uint16(527),
	10605: uint16(16),
	10606: uint16(anon_sym_return),
	10607: uint16(sym_break_statement),
	10608: uint16(anon_sym_goto),
	10609: uint16(anon_sym_do),
	10610: uint16(anon_sym_end),
	10611: uint16(anon_sym_while),
	10612: uint16(anon_sym_repeat),
	10613: uint16(anon_sym_until),
	10614: uint16(anon_sym_if),
	10615: uint16(anon_sym_elseif),
	10616: uint16(anon_sym_else),
	10617: uint16(anon_sym_for),
	10618: uint16(anon_sym_function),
	10619: uint16(anon_sym_local),
	10620: uint16(anon_sym_global),
	10621: uint16(sym_identifier),
	10622: uint16(5),
	10623: uint16(3),
	10624: uint16(1),
	10625: uint16(anon_sym_DASH_DASH),
	10626: uint16(5),
	10627: uint16(1),
	10628: uint16(sym__block_comment_start),
	10629: uint16(159),
	10630: uint16(1),
	10631: uint16(sym_comment),
	10632: uint16(529),
	10633: uint16(4),
	10635: uint16(anon_sym_SEMI),
	10636: uint16(anon_sym_COLON_COLON),
	10637: uint16(anon_sym_LPAREN),
	10638: uint16(531),
	10639: uint16(16),
	10640: uint16(anon_sym_return),
	10641: uint16(sym_break_statement),
	10642: uint16(anon_sym_goto),
	10643: uint16(anon_sym_do),
	10644: uint16(anon_sym_end),
	10645: uint16(anon_sym_while),
	10646: uint16(anon_sym_repeat),
	10647: uint16(anon_sym_until),
	10648: uint16(anon_sym_if),
	10649: uint16(anon_sym_elseif),
	10650: uint16(anon_sym_else),
	10651: uint16(anon_sym_for),
	10652: uint16(anon_sym_function),
	10653: uint16(anon_sym_local),
	10654: uint16(anon_sym_global),
	10655: uint16(sym_identifier),
	10656: uint16(5),
	10657: uint16(3),
	10658: uint16(1),
	10659: uint16(anon_sym_DASH_DASH),
	10660: uint16(5),
	10661: uint16(1),
	10662: uint16(sym__block_comment_start),
	10663: uint16(160),
	10664: uint16(1),
	10665: uint16(sym_comment),
	10666: uint16(309),
	10667: uint16(4),
	10669: uint16(anon_sym_SEMI),
	10670: uint16(anon_sym_COLON_COLON),
	10671: uint16(anon_sym_LPAREN),
	10672: uint16(311),
	10673: uint16(16),
	10674: uint16(anon_sym_return),
	10675: uint16(sym_break_statement),
	10676: uint16(anon_sym_goto),
	10677: uint16(anon_sym_do),
	10678: uint16(anon_sym_end),
	10679: uint16(anon_sym_while),
	10680: uint16(anon_sym_repeat),
	10681: uint16(anon_sym_until),
	10682: uint16(anon_sym_if),
	10683: uint16(anon_sym_elseif),
	10684: uint16(anon_sym_else),
	10685: uint16(anon_sym_for),
	10686: uint16(anon_sym_function),
	10687: uint16(anon_sym_local),
	10688: uint16(anon_sym_global),
	10689: uint16(sym_identifier),
	10690: uint16(5),
	10691: uint16(3),
	10692: uint16(1),
	10693: uint16(anon_sym_DASH_DASH),
	10694: uint16(5),
	10695: uint16(1),
	10696: uint16(sym__block_comment_start),
	10697: uint16(161),
	10698: uint16(1),
	10699: uint16(sym_comment),
	10700: uint16(533),
	10701: uint16(4),
	10703: uint16(anon_sym_SEMI),
	10704: uint16(anon_sym_COLON_COLON),
	10705: uint16(anon_sym_LPAREN),
	10706: uint16(535),
	10707: uint16(16),
	10708: uint16(anon_sym_return),
	10709: uint16(sym_break_statement),
	10710: uint16(anon_sym_goto),
	10711: uint16(anon_sym_do),
	10712: uint16(anon_sym_end),
	10713: uint16(anon_sym_while),
	10714: uint16(anon_sym_repeat),
	10715: uint16(anon_sym_until),
	10716: uint16(anon_sym_if),
	10717: uint16(anon_sym_elseif),
	10718: uint16(anon_sym_else),
	10719: uint16(anon_sym_for),
	10720: uint16(anon_sym_function),
	10721: uint16(anon_sym_local),
	10722: uint16(anon_sym_global),
	10723: uint16(sym_identifier),
	10724: uint16(5),
	10725: uint16(3),
	10726: uint16(1),
	10727: uint16(anon_sym_DASH_DASH),
	10728: uint16(5),
	10729: uint16(1),
	10730: uint16(sym__block_comment_start),
	10731: uint16(162),
	10732: uint16(1),
	10733: uint16(sym_comment),
	10734: uint16(537),
	10735: uint16(4),
	10737: uint16(anon_sym_SEMI),
	10738: uint16(anon_sym_COLON_COLON),
	10739: uint16(anon_sym_LPAREN),
	10740: uint16(539),
	10741: uint16(16),
	10742: uint16(anon_sym_return),
	10743: uint16(sym_break_statement),
	10744: uint16(anon_sym_goto),
	10745: uint16(anon_sym_do),
	10746: uint16(anon_sym_end),
	10747: uint16(anon_sym_while),
	10748: uint16(anon_sym_repeat),
	10749: uint16(anon_sym_until),
	10750: uint16(anon_sym_if),
	10751: uint16(anon_sym_elseif),
	10752: uint16(anon_sym_else),
	10753: uint16(anon_sym_for),
	10754: uint16(anon_sym_function),
	10755: uint16(anon_sym_local),
	10756: uint16(anon_sym_global),
	10757: uint16(sym_identifier),
	10758: uint16(5),
	10759: uint16(3),
	10760: uint16(1),
	10761: uint16(anon_sym_DASH_DASH),
	10762: uint16(5),
	10763: uint16(1),
	10764: uint16(sym__block_comment_start),
	10765: uint16(163),
	10766: uint16(1),
	10767: uint16(sym_comment),
	10768: uint16(541),
	10769: uint16(4),
	10771: uint16(anon_sym_SEMI),
	10772: uint16(anon_sym_COLON_COLON),
	10773: uint16(anon_sym_LPAREN),
	10774: uint16(543),
	10775: uint16(16),
	10776: uint16(anon_sym_return),
	10777: uint16(sym_break_statement),
	10778: uint16(anon_sym_goto),
	10779: uint16(anon_sym_do),
	10780: uint16(anon_sym_end),
	10781: uint16(anon_sym_while),
	10782: uint16(anon_sym_repeat),
	10783: uint16(anon_sym_until),
	10784: uint16(anon_sym_if),
	10785: uint16(anon_sym_elseif),
	10786: uint16(anon_sym_else),
	10787: uint16(anon_sym_for),
	10788: uint16(anon_sym_function),
	10789: uint16(anon_sym_local),
	10790: uint16(anon_sym_global),
	10791: uint16(sym_identifier),
	10792: uint16(5),
	10793: uint16(3),
	10794: uint16(1),
	10795: uint16(anon_sym_DASH_DASH),
	10796: uint16(5),
	10797: uint16(1),
	10798: uint16(sym__block_comment_start),
	10799: uint16(164),
	10800: uint16(1),
	10801: uint16(sym_comment),
	10802: uint16(545),
	10803: uint16(4),
	10805: uint16(anon_sym_SEMI),
	10806: uint16(anon_sym_COLON_COLON),
	10807: uint16(anon_sym_LPAREN),
	10808: uint16(547),
	10809: uint16(16),
	10810: uint16(anon_sym_return),
	10811: uint16(sym_break_statement),
	10812: uint16(anon_sym_goto),
	10813: uint16(anon_sym_do),
	10814: uint16(anon_sym_end),
	10815: uint16(anon_sym_while),
	10816: uint16(anon_sym_repeat),
	10817: uint16(anon_sym_until),
	10818: uint16(anon_sym_if),
	10819: uint16(anon_sym_elseif),
	10820: uint16(anon_sym_else),
	10821: uint16(anon_sym_for),
	10822: uint16(anon_sym_function),
	10823: uint16(anon_sym_local),
	10824: uint16(anon_sym_global),
	10825: uint16(sym_identifier),
	10826: uint16(5),
	10827: uint16(3),
	10828: uint16(1),
	10829: uint16(anon_sym_DASH_DASH),
	10830: uint16(5),
	10831: uint16(1),
	10832: uint16(sym__block_comment_start),
	10833: uint16(165),
	10834: uint16(1),
	10835: uint16(sym_comment),
	10836: uint16(549),
	10837: uint16(4),
	10839: uint16(anon_sym_SEMI),
	10840: uint16(anon_sym_COLON_COLON),
	10841: uint16(anon_sym_LPAREN),
	10842: uint16(551),
	10843: uint16(16),
	10844: uint16(anon_sym_return),
	10845: uint16(sym_break_statement),
	10846: uint16(anon_sym_goto),
	10847: uint16(anon_sym_do),
	10848: uint16(anon_sym_end),
	10849: uint16(anon_sym_while),
	10850: uint16(anon_sym_repeat),
	10851: uint16(anon_sym_until),
	10852: uint16(anon_sym_if),
	10853: uint16(anon_sym_elseif),
	10854: uint16(anon_sym_else),
	10855: uint16(anon_sym_for),
	10856: uint16(anon_sym_function),
	10857: uint16(anon_sym_local),
	10858: uint16(anon_sym_global),
	10859: uint16(sym_identifier),
	10860: uint16(5),
	10861: uint16(3),
	10862: uint16(1),
	10863: uint16(anon_sym_DASH_DASH),
	10864: uint16(5),
	10865: uint16(1),
	10866: uint16(sym__block_comment_start),
	10867: uint16(166),
	10868: uint16(1),
	10869: uint16(sym_comment),
	10870: uint16(553),
	10871: uint16(8),
	10872: uint16(anon_sym_function),
	10873: uint16(anon_sym_global),
	10874: uint16(sym_nil),
	10875: uint16(sym_false),
	10876: uint16(sym_true),
	10877: uint16(anon_sym_DASH),
	10878: uint16(anon_sym_not),
	10879: uint16(sym_identifier),
	10880: uint16(555),
	10881: uint16(11),
	10882: uint16(sym__block_string_start),
	10883: uint16(sym_number),
	10884: uint16(anon_sym_DQUOTE),
	10885: uint16(anon_sym_SQUOTE),
	10886: uint16(sym_vararg_expression),
	10887: uint16(anon_sym_LPAREN),
	10888: uint16(anon_sym_LBRACK),
	10889: uint16(anon_sym_LBRACE),
	10890: uint16(anon_sym_RBRACE),
	10891: uint16(anon_sym_TILDE),
	10892: uint16(anon_sym_POUND),
	10893: uint16(11),
	10894: uint16(3),
	10895: uint16(1),
	10896: uint16(anon_sym_DASH_DASH),
	10897: uint16(5),
	10898: uint16(1),
	10899: uint16(sym__block_comment_start),
	10900: uint16(319),
	10901: uint16(1),
	10902: uint16(anon_sym_LT),
	10903: uint16(557),
	10904: uint16(1),
	10905: uint16(sym_identifier),
	10906: uint16(559),
	10907: uint16(1),
	10908: uint16(anon_sym_function),
	10909: uint16(561),
	10910: uint16(1),
	10911: uint16(anon_sym_STAR),
	10912: uint16(129),
	10913: uint16(1),
	10914: uint16(sym__att_name_list),
	10915: uint16(141),
	10916: uint16(1),
	10917: uint16(sym__variable_assignment),
	10918: uint16(167),
	10919: uint16(1),
	10920: uint16(sym_comment),
	10921: uint16(230),
	10922: uint16(1),
	10923: uint16(sym__attrib),
	10924: uint16(45),
	10925: uint16(10),
	10926: uint16(sym__block_string_start),
	10927: uint16(anon_sym_EQ),
	10928: uint16(anon_sym_COMMA),
	10929: uint16(anon_sym_DOT),
	10930: uint16(anon_sym_COLON),
	10931: uint16(anon_sym_DQUOTE),
	10932: uint16(anon_sym_SQUOTE),
	10933: uint16(anon_sym_LPAREN),
	10934: uint16(anon_sym_LBRACK),
	10935: uint16(anon_sym_LBRACE),
	10936: uint16(5),
	10937: uint16(3),
	10938: uint16(1),
	10939: uint16(anon_sym_DASH_DASH),
	10940: uint16(5),
	10941: uint16(1),
	10942: uint16(sym__block_comment_start),
	10943: uint16(168),
	10944: uint16(1),
	10945: uint16(sym_comment),
	10946: uint16(565),
	10947: uint16(3),
	10948: uint16(anon_sym_SEMI),
	10949: uint16(anon_sym_COLON_COLON),
	10950: uint16(anon_sym_LPAREN),
	10951: uint16(563),
	10952: uint16(13),
	10953: uint16(anon_sym_return),
	10954: uint16(sym_break_statement),
	10955: uint16(anon_sym_goto),
	10956: uint16(anon_sym_do),
	10957: uint16(anon_sym_end),
	10958: uint16(anon_sym_while),
	10959: uint16(anon_sym_repeat),
	10960: uint16(anon_sym_if),
	10961: uint16(anon_sym_for),
	10962: uint16(anon_sym_function),
	10963: uint16(anon_sym_local),
	10964: uint16(anon_sym_global),
	10965: uint16(sym_identifier),
	10966: uint16(5),
	10967: uint16(3),
	10968: uint16(1),
	10969: uint16(anon_sym_DASH_DASH),
	10970: uint16(5),
	10971: uint16(1),
	10972: uint16(sym__block_comment_start),
	10973: uint16(169),
	10974: uint16(1),
	10975: uint16(sym_comment),
	10976: uint16(569),
	10977: uint16(3),
	10978: uint16(anon_sym_SEMI),
	10979: uint16(anon_sym_COLON_COLON),
	10980: uint16(anon_sym_LPAREN),
	10981: uint16(567),
	10982: uint16(13),
	10983: uint16(anon_sym_return),
	10984: uint16(sym_break_statement),
	10985: uint16(anon_sym_goto),
	10986: uint16(anon_sym_do),
	10987: uint16(anon_sym_end),
	10988: uint16(anon_sym_while),
	10989: uint16(anon_sym_repeat),
	10990: uint16(anon_sym_if),
	10991: uint16(anon_sym_for),
	10992: uint16(anon_sym_function),
	10993: uint16(anon_sym_local),
	10994: uint16(anon_sym_global),
	10995: uint16(sym_identifier),
	10996: uint16(15),
	10997: uint16(3),
	10998: uint16(1),
	10999: uint16(anon_sym_DASH_DASH),
	11000: uint16(5),
	11001: uint16(1),
	11002: uint16(sym__block_comment_start),
	11003: uint16(571),
	11004: uint16(1),
	11005: uint16(anon_sym_DOT),
	11006: uint16(573),
	11007: uint16(1),
	11008: uint16(anon_sym_COLON),
	11009: uint16(575),
	11010: uint16(1),
	11011: uint16(anon_sym_DQUOTE),
	11012: uint16(577),
	11013: uint16(1),
	11014: uint16(anon_sym_SQUOTE),
	11015: uint16(579),
	11016: uint16(1),
	11017: uint16(anon_sym_LPAREN),
	11018: uint16(581),
	11019: uint16(1),
	11020: uint16(anon_sym_LBRACK),
	11021: uint16(583),
	11022: uint16(1),
	11023: uint16(anon_sym_LBRACE),
	11024: uint16(585),
	11025: uint16(1),
	11026: uint16(sym__block_string_start),
	11027: uint16(9),
	11028: uint16(1),
	11029: uint16(sym_arguments),
	11030: uint16(14),
	11031: uint16(1),
	11032: uint16(sym__quote_string),
	11033: uint16(16),
	11034: uint16(1),
	11035: uint16(sym__block_string),
	11036: uint16(170),
	11037: uint16(1),
	11038: uint16(sym_comment),
	11039: uint16(8),
	11040: uint16(2),
	11041: uint16(sym_string),
	11042: uint16(sym_table_constructor),
	11043: uint16(11),
	11044: uint16(3),
	11045: uint16(1),
	11046: uint16(anon_sym_DASH_DASH),
	11047: uint16(5),
	11048: uint16(1),
	11049: uint16(sym__block_comment_start),
	11050: uint16(9),
	11051: uint16(1),
	11052: uint16(sym_identifier),
	11053: uint16(39),
	11054: uint16(1),
	11055: uint16(anon_sym_LPAREN),
	11056: uint16(259),
	11057: uint16(1),
	11058: uint16(anon_sym_global),
	11059: uint16(170),
	11060: uint16(1),
	11061: uint16(sym__prefix_expression),
	11062: uint16(171),
	11063: uint16(1),
	11064: uint16(sym_comment),
	11065: uint16(174),
	11066: uint16(1),
	11067: uint16(sym_method_index_expression),
	11068: uint16(175),
	11069: uint16(1),
	11070: uint16(sym_variable),
	11071: uint16(177),
	11072: uint16(2),
	11073: uint16(sym_function_call),
	11074: uint16(sym_parenthesized_expression),
	11075: uint16(2),
	11076: uint16(3),
	11077: uint16(sym_bracket_index_expression),
	11078: uint16(sym_dot_index_expression),
	11079: uint16(sym__contextual_keyword),
	11080: uint16(7),
	11081: uint16(3),
	11082: uint16(1),
	11083: uint16(anon_sym_DASH_DASH),
	11084: uint16(5),
	11085: uint16(1),
	11086: uint16(sym__block_comment_start),
	11087: uint16(587),
	11088: uint16(1),
	11089: uint16(anon_sym_EQ),
	11090: uint16(589),
	11091: uint16(1),
	11092: uint16(anon_sym_COMMA),
	11093: uint16(172),
	11094: uint16(1),
	11095: uint16(sym_comment),
	11096: uint16(217),
	11097: uint16(1),
	11098: uint16(aux_sym__variable_assignment_varlist_repeat1),
	11099: uint16(67),
	11100: uint16(8),
	11101: uint16(sym__block_string_start),
	11102: uint16(anon_sym_DOT),
	11103: uint16(anon_sym_COLON),
	11104: uint16(anon_sym_DQUOTE),
	11105: uint16(anon_sym_SQUOTE),
	11106: uint16(anon_sym_LPAREN),
	11107: uint16(anon_sym_LBRACK),
	11108: uint16(anon_sym_LBRACE),
	11109: uint16(6),
	11110: uint16(3),
	11111: uint16(1),
	11112: uint16(anon_sym_DASH_DASH),
	11113: uint16(5),
	11114: uint16(1),
	11115: uint16(sym__block_comment_start),
	11116: uint16(297),
	11117: uint16(1),
	11118: uint16(anon_sym_else),
	11119: uint16(591),
	11120: uint16(1),
	11121: uint16(anon_sym_COMMA),
	11122: uint16(173),
	11123: uint16(2),
	11124: uint16(sym_comment),
	11125: uint16(aux_sym__expression_list_repeat1),
	11126: uint16(295),
	11127: uint16(7),
	11129: uint16(anon_sym_SEMI),
	11130: uint16(anon_sym_do),
	11131: uint16(anon_sym_end),
	11132: uint16(anon_sym_until),
	11133: uint16(anon_sym_elseif),
	11134: uint16(anon_sym_RPAREN),
	11135: uint16(12),
	11136: uint16(3),
	11137: uint16(1),
	11138: uint16(anon_sym_DASH_DASH),
	11139: uint16(5),
	11140: uint16(1),
	11141: uint16(sym__block_comment_start),
	11142: uint16(575),
	11143: uint16(1),
	11144: uint16(anon_sym_DQUOTE),
	11145: uint16(577),
	11146: uint16(1),
	11147: uint16(anon_sym_SQUOTE),
	11148: uint16(579),
	11149: uint16(1),
	11150: uint16(anon_sym_LPAREN),
	11151: uint16(583),
	11152: uint16(1),
	11153: uint16(anon_sym_LBRACE),
	11154: uint16(585),
	11155: uint16(1),
	11156: uint16(sym__block_string_start),
	11157: uint16(9),
	11158: uint16(1),
	11159: uint16(sym_arguments),
	11160: uint16(14),
	11161: uint16(1),
	11162: uint16(sym__quote_string),
	11163: uint16(16),
	11164: uint16(1),
	11165: uint16(sym__block_string),
	11166: uint16(174),
	11167: uint16(1),
	11168: uint16(sym_comment),
	11169: uint16(8),
	11170: uint16(2),
	11171: uint16(sym_string),
	11172: uint16(sym_table_constructor),
	11173: uint16(5),
	11174: uint16(3),
	11175: uint16(1),
	11176: uint16(anon_sym_DASH_DASH),
	11177: uint16(5),
	11178: uint16(1),
	11179: uint16(sym__block_comment_start),
	11180: uint16(175),
	11181: uint16(1),
	11182: uint16(sym_comment),
	11183: uint16(594),
	11184: uint16(2),
	11185: uint16(anon_sym_EQ),
	11186: uint16(anon_sym_COMMA),
	11187: uint16(67),
	11188: uint16(8),
	11189: uint16(sym__block_string_start),
	11190: uint16(anon_sym_DOT),
	11191: uint16(anon_sym_COLON),
	11192: uint16(anon_sym_DQUOTE),
	11193: uint16(anon_sym_SQUOTE),
	11194: uint16(anon_sym_LPAREN),
	11195: uint16(anon_sym_LBRACK),
	11196: uint16(anon_sym_LBRACE),
	11197: uint16(7),
	11198: uint16(3),
	11199: uint16(1),
	11200: uint16(anon_sym_DASH_DASH),
	11201: uint16(5),
	11202: uint16(1),
	11203: uint16(sym__block_comment_start),
	11204: uint16(305),
	11205: uint16(1),
	11206: uint16(anon_sym_COMMA),
	11207: uint16(598),
	11208: uint16(1),
	11209: uint16(anon_sym_else),
	11210: uint16(173),
	11211: uint16(1),
	11212: uint16(aux_sym__expression_list_repeat1),
	11213: uint16(176),
	11214: uint16(1),
	11215: uint16(sym_comment),
	11216: uint16(596),
	11217: uint16(6),
	11219: uint16(anon_sym_SEMI),
	11220: uint16(anon_sym_do),
	11221: uint16(anon_sym_end),
	11222: uint16(anon_sym_until),
	11223: uint16(anon_sym_elseif),
	11224: uint16(4),
	11225: uint16(3),
	11226: uint16(1),
	11227: uint16(anon_sym_DASH_DASH),
	11228: uint16(5),
	11229: uint16(1),
	11230: uint16(sym__block_comment_start),
	11231: uint16(177),
	11232: uint16(1),
	11233: uint16(sym_comment),
	11234: uint16(67),
	11235: uint16(8),
	11236: uint16(sym__block_string_start),
	11237: uint16(anon_sym_DOT),
	11238: uint16(anon_sym_COLON),
	11239: uint16(anon_sym_DQUOTE),
	11240: uint16(anon_sym_SQUOTE),
	11241: uint16(anon_sym_LPAREN),
	11242: uint16(anon_sym_LBRACK),
	11243: uint16(anon_sym_LBRACE),
	11244: uint16(9),
	11245: uint16(3),
	11246: uint16(1),
	11247: uint16(anon_sym_DASH_DASH),
	11248: uint16(5),
	11249: uint16(1),
	11250: uint16(sym__block_comment_start),
	11251: uint16(119),
	11252: uint16(1),
	11253: uint16(anon_sym_else),
	11254: uint16(600),
	11255: uint16(1),
	11256: uint16(anon_sym_end),
	11257: uint16(602),
	11258: uint16(1),
	11259: uint16(anon_sym_elseif),
	11260: uint16(178),
	11261: uint16(1),
	11262: uint16(sym_comment),
	11263: uint16(186),
	11264: uint16(1),
	11265: uint16(aux_sym_if_statement_repeat1),
	11266: uint16(213),
	11267: uint16(1),
	11268: uint16(sym_elseif_statement),
	11269: uint16(254),
	11270: uint16(1),
	11271: uint16(sym_else_statement),
	11272: uint16(9),
	11273: uint16(3),
	11274: uint16(1),
	11275: uint16(anon_sym_DASH_DASH),
	11276: uint16(5),
	11277: uint16(1),
	11278: uint16(sym__block_comment_start),
	11279: uint16(119),
	11280: uint16(1),
	11281: uint16(anon_sym_else),
	11282: uint16(602),
	11283: uint16(1),
	11284: uint16(anon_sym_elseif),
	11285: uint16(604),
	11286: uint16(1),
	11287: uint16(anon_sym_end),
	11288: uint16(179),
	11289: uint16(1),
	11290: uint16(sym_comment),
	11291: uint16(186),
	11292: uint16(1),
	11293: uint16(aux_sym_if_statement_repeat1),
	11294: uint16(213),
	11295: uint16(1),
	11296: uint16(sym_elseif_statement),
	11297: uint16(240),
	11298: uint16(1),
	11299: uint16(sym_else_statement),
	11300: uint16(9),
	11301: uint16(3),
	11302: uint16(1),
	11303: uint16(anon_sym_DASH_DASH),
	11304: uint16(5),
	11305: uint16(1),
	11306: uint16(sym__block_comment_start),
	11307: uint16(119),
	11308: uint16(1),
	11309: uint16(anon_sym_else),
	11310: uint16(602),
	11311: uint16(1),
	11312: uint16(anon_sym_elseif),
	11313: uint16(606),
	11314: uint16(1),
	11315: uint16(anon_sym_end),
	11316: uint16(178),
	11317: uint16(1),
	11318: uint16(aux_sym_if_statement_repeat1),
	11319: uint16(180),
	11320: uint16(1),
	11321: uint16(sym_comment),
	11322: uint16(213),
	11323: uint16(1),
	11324: uint16(sym_elseif_statement),
	11325: uint16(236),
	11326: uint16(1),
	11327: uint16(sym_else_statement),
	11328: uint16(9),
	11329: uint16(3),
	11330: uint16(1),
	11331: uint16(anon_sym_DASH_DASH),
	11332: uint16(5),
	11333: uint16(1),
	11334: uint16(sym__block_comment_start),
	11335: uint16(319),
	11336: uint16(1),
	11337: uint16(anon_sym_LT),
	11338: uint16(557),
	11339: uint16(1),
	11340: uint16(sym_identifier),
	11341: uint16(608),
	11342: uint16(1),
	11343: uint16(anon_sym_function),
	11344: uint16(130),
	11345: uint16(1),
	11346: uint16(sym__att_name_list),
	11347: uint16(138),
	11348: uint16(1),
	11349: uint16(sym__variable_assignment),
	11350: uint16(181),
	11351: uint16(1),
	11352: uint16(sym_comment),
	11353: uint16(239),
	11354: uint16(1),
	11355: uint16(sym__attrib),
	11356: uint16(6),
	11357: uint16(3),
	11358: uint16(1),
	11359: uint16(anon_sym_DASH_DASH),
	11360: uint16(5),
	11361: uint16(1),
	11362: uint16(sym__block_comment_start),
	11363: uint16(612),
	11364: uint16(1),
	11365: uint16(anon_sym_SEMI),
	11366: uint16(614),
	11367: uint16(1),
	11368: uint16(anon_sym_else),
	11369: uint16(182),
	11370: uint16(1),
	11371: uint16(sym_comment),
	11372: uint16(610),
	11373: uint16(4),
	11375: uint16(anon_sym_end),
	11376: uint16(anon_sym_until),
	11377: uint16(anon_sym_elseif),
	11378: uint16(6),
	11379: uint16(3),
	11380: uint16(1),
	11381: uint16(anon_sym_DASH_DASH),
	11382: uint16(5),
	11383: uint16(1),
	11384: uint16(sym__block_comment_start),
	11385: uint16(619),
	11386: uint16(1),
	11387: uint16(anon_sym_RBRACE),
	11388: uint16(66),
	11389: uint16(1),
	11390: uint16(sym__field_sep),
	11391: uint16(616),
	11392: uint16(2),
	11393: uint16(anon_sym_SEMI),
	11394: uint16(anon_sym_COMMA),
	11395: uint16(183),
	11396: uint16(2),
	11397: uint16(sym_comment),
	11398: uint16(aux_sym__field_list_repeat1),
	11399: uint16(7),
	11400: uint16(3),
	11401: uint16(1),
	11402: uint16(anon_sym_DASH_DASH),
	11403: uint16(5),
	11404: uint16(1),
	11405: uint16(sym__block_comment_start),
	11406: uint16(287),
	11407: uint16(1),
	11408: uint16(anon_sym_RBRACE),
	11409: uint16(65),
	11410: uint16(1),
	11411: uint16(sym__field_sep),
	11412: uint16(183),
	11413: uint16(1),
	11414: uint16(aux_sym__field_list_repeat1),
	11415: uint16(184),
	11416: uint16(1),
	11417: uint16(sym_comment),
	11418: uint16(621),
	11419: uint16(2),
	11420: uint16(anon_sym_SEMI),
	11421: uint16(anon_sym_COMMA),
	11422: uint16(5),
	11423: uint16(3),
	11424: uint16(1),
	11425: uint16(anon_sym_DASH_DASH),
	11426: uint16(5),
	11427: uint16(1),
	11428: uint16(sym__block_comment_start),
	11429: uint16(614),
	11430: uint16(1),
	11431: uint16(anon_sym_else),
	11432: uint16(185),
	11433: uint16(1),
	11434: uint16(sym_comment),
	11435: uint16(610),
	11436: uint16(4),
	11438: uint16(anon_sym_end),
	11439: uint16(anon_sym_until),
	11440: uint16(anon_sym_elseif),
	11441: uint16(7),
	11442: uint16(3),
	11443: uint16(1),
	11444: uint16(anon_sym_DASH_DASH),
	11445: uint16(5),
	11446: uint16(1),
	11447: uint16(sym__block_comment_start),
	11448: uint16(623),
	11449: uint16(1),
	11450: uint16(anon_sym_end),
	11451: uint16(625),
	11452: uint16(1),
	11453: uint16(anon_sym_elseif),
	11454: uint16(628),
	11455: uint16(1),
	11456: uint16(anon_sym_else),
	11457: uint16(213),
	11458: uint16(1),
	11459: uint16(sym_elseif_statement),
	11460: uint16(186),
	11461: uint16(2),
	11462: uint16(sym_comment),
	11463: uint16(aux_sym_if_statement_repeat1),
	11464: uint16(4),
	11465: uint16(3),
	11466: uint16(1),
	11467: uint16(anon_sym_DASH_DASH),
	11468: uint16(5),
	11469: uint16(1),
	11470: uint16(sym__block_comment_start),
	11471: uint16(187),
	11472: uint16(1),
	11473: uint16(sym_comment),
	11474: uint16(630),
	11475: uint16(5),
	11476: uint16(sym__block_string_start),
	11477: uint16(anon_sym_DQUOTE),
	11478: uint16(anon_sym_SQUOTE),
	11479: uint16(anon_sym_LPAREN),
	11480: uint16(anon_sym_LBRACE),
	11481: uint16(7),
	11482: uint16(3),
	11483: uint16(1),
	11484: uint16(anon_sym_DASH_DASH),
	11485: uint16(5),
	11486: uint16(1),
	11487: uint16(sym__block_comment_start),
	11488: uint16(632),
	11489: uint16(1),
	11490: uint16(anon_sym_RBRACE),
	11491: uint16(64),
	11492: uint16(1),
	11493: uint16(sym__field_sep),
	11494: uint16(184),
	11495: uint16(1),
	11496: uint16(aux_sym__field_list_repeat1),
	11497: uint16(188),
	11498: uint16(1),
	11499: uint16(sym_comment),
	11500: uint16(621),
	11501: uint16(2),
	11502: uint16(anon_sym_SEMI),
	11503: uint16(anon_sym_COMMA),
	11504: uint16(5),
	11505: uint16(3),
	11506: uint16(1),
	11507: uint16(anon_sym_DASH_DASH),
	11508: uint16(5),
	11509: uint16(1),
	11510: uint16(sym__block_comment_start),
	11511: uint16(636),
	11512: uint16(1),
	11513: uint16(anon_sym_else),
	11514: uint16(189),
	11515: uint16(1),
	11516: uint16(sym_comment),
	11517: uint16(634),
	11518: uint16(4),
	11520: uint16(anon_sym_end),
	11521: uint16(anon_sym_until),
	11522: uint16(anon_sym_elseif),
	11523: uint16(8),
	11524: uint16(3),
	11525: uint16(1),
	11526: uint16(anon_sym_DASH_DASH),
	11527: uint16(5),
	11528: uint16(1),
	11529: uint16(sym__block_comment_start),
	11530: uint16(638),
	11531: uint16(1),
	11532: uint16(sym_identifier),
	11533: uint16(190),
	11534: uint16(1),
	11535: uint16(sym_comment),
	11536: uint16(222),
	11537: uint16(1),
	11538: uint16(sym__function_name),
	11539: uint16(226),
	11540: uint16(1),
	11541: uint16(sym__function_name_dot_index_expression),
	11542: uint16(227),
	11543: uint16(1),
	11544: uint16(sym__function_name_prefix_expression),
	11545: uint16(250),
	11546: uint16(1),
	11547: uint16(sym__function_name_method_index_expression),
	11548: uint16(8),
	11549: uint16(3),
	11550: uint16(1),
	11551: uint16(anon_sym_DASH_DASH),
	11552: uint16(5),
	11553: uint16(1),
	11554: uint16(sym__block_comment_start),
	11555: uint16(640),
	11556: uint16(1),
	11557: uint16(sym_identifier),
	11558: uint16(642),
	11559: uint16(1),
	11560: uint16(sym_vararg_expression),
	11561: uint16(644),
	11562: uint16(1),
	11563: uint16(anon_sym_RPAREN),
	11564: uint16(191),
	11565: uint16(1),
	11566: uint16(sym_comment),
	11567: uint16(246),
	11568: uint16(1),
	11569: uint16(sym__parameter_list),
	11570: uint16(273),
	11571: uint16(1),
	11572: uint16(sym__vararg_parameter),
	11573: uint16(5),
	11574: uint16(3),
	11575: uint16(1),
	11576: uint16(anon_sym_DASH_DASH),
	11577: uint16(5),
	11578: uint16(1),
	11579: uint16(sym__block_comment_start),
	11580: uint16(648),
	11581: uint16(1),
	11582: uint16(anon_sym_else),
	11583: uint16(192),
	11584: uint16(1),
	11585: uint16(sym_comment),
	11586: uint16(646),
	11587: uint16(3),
	11588: uint16(anon_sym_end),
	11589: uint16(anon_sym_until),
	11590: uint16(anon_sym_elseif),
	11591: uint16(6),
	11592: uint16(5),
	11593: uint16(1),
	11594: uint16(sym__block_comment_start),
	11595: uint16(650),
	11596: uint16(1),
	11597: uint16(anon_sym_SQUOTE),
	11598: uint16(654),
	11599: uint16(1),
	11600: uint16(anon_sym_DASH_DASH),
	11601: uint16(193),
	11602: uint16(1),
	11603: uint16(sym_comment),
	11604: uint16(199),
	11605: uint16(1),
	11606: uint16(aux_sym__singlequote_string_content),
	11607: uint16(652),
	11608: uint16(2),
	11609: uint16(aux_sym__singlequote_string_content_token1),
	11610: uint16(sym_escape_sequence),
	11611: uint16(5),
	11612: uint16(3),
	11613: uint16(1),
	11614: uint16(anon_sym_DASH_DASH),
	11615: uint16(5),
	11616: uint16(1),
	11617: uint16(sym__block_comment_start),
	11618: uint16(121),
	11619: uint16(1),
	11620: uint16(anon_sym_else),
	11621: uint16(194),
	11622: uint16(1),
	11623: uint16(sym_comment),
	11624: uint16(656),
	11625: uint16(3),
	11626: uint16(anon_sym_end),
	11627: uint16(anon_sym_until),
	11628: uint16(anon_sym_elseif),
	11629: uint16(5),
	11630: uint16(3),
	11631: uint16(1),
	11632: uint16(anon_sym_DASH_DASH),
	11633: uint16(5),
	11634: uint16(1),
	11635: uint16(sym__block_comment_start),
	11636: uint16(658),
	11637: uint16(1),
	11638: uint16(anon_sym_COMMA),
	11639: uint16(661),
	11640: uint16(2),
	11641: uint16(anon_sym_in),
	11642: uint16(anon_sym_RPAREN),
	11643: uint16(195),
	11644: uint16(2),
	11645: uint16(sym_comment),
	11646: uint16(aux_sym__name_list_repeat1),
	11647: uint16(6),
	11648: uint16(5),
	11649: uint16(1),
	11650: uint16(sym__block_comment_start),
	11651: uint16(654),
	11652: uint16(1),
	11653: uint16(anon_sym_DASH_DASH),
	11654: uint16(663),
	11655: uint16(1),
	11656: uint16(anon_sym_DQUOTE),
	11657: uint16(196),
	11658: uint16(1),
	11659: uint16(sym_comment),
	11660: uint16(198),
	11661: uint16(1),
	11662: uint16(aux_sym__doublequote_string_content),
	11663: uint16(665),
	11664: uint16(2),
	11665: uint16(aux_sym__doublequote_string_content_token1),
	11666: uint16(sym_escape_sequence),
	11667: uint16(6),
	11668: uint16(5),
	11669: uint16(1),
	11670: uint16(sym__block_comment_start),
	11671: uint16(654),
	11672: uint16(1),
	11673: uint16(anon_sym_DASH_DASH),
	11674: uint16(663),
	11675: uint16(1),
	11676: uint16(anon_sym_SQUOTE),
	11677: uint16(197),
	11678: uint16(1),
	11679: uint16(sym_comment),
	11680: uint16(199),
	11681: uint16(1),
	11682: uint16(aux_sym__singlequote_string_content),
	11683: uint16(652),
	11684: uint16(2),
	11685: uint16(aux_sym__singlequote_string_content_token1),
	11686: uint16(sym_escape_sequence),
	11687: uint16(5),
	11688: uint16(5),
	11689: uint16(1),
	11690: uint16(sym__block_comment_start),
	11691: uint16(654),
	11692: uint16(1),
	11693: uint16(anon_sym_DASH_DASH),
	11694: uint16(667),
	11695: uint16(1),
	11696: uint16(anon_sym_DQUOTE),
	11697: uint16(669),
	11698: uint16(2),
	11699: uint16(aux_sym__doublequote_string_content_token1),
	11700: uint16(sym_escape_sequence),
	11701: uint16(198),
	11702: uint16(2),
	11703: uint16(aux_sym__doublequote_string_content),
	11704: uint16(sym_comment),
	11705: uint16(5),
	11706: uint16(5),
	11707: uint16(1),
	11708: uint16(sym__block_comment_start),
	11709: uint16(654),
	11710: uint16(1),
	11711: uint16(anon_sym_DASH_DASH),
	11712: uint16(672),
	11713: uint16(1),
	11714: uint16(anon_sym_SQUOTE),
	11715: uint16(674),
	11716: uint16(2),
	11717: uint16(aux_sym__singlequote_string_content_token1),
	11718: uint16(sym_escape_sequence),
	11719: uint16(199),
	11720: uint16(2),
	11721: uint16(aux_sym__singlequote_string_content),
	11722: uint16(sym_comment),
	11723: uint16(7),
	11724: uint16(3),
	11725: uint16(1),
	11726: uint16(anon_sym_DASH_DASH),
	11727: uint16(5),
	11728: uint16(1),
	11729: uint16(sym__block_comment_start),
	11730: uint16(677),
	11731: uint16(1),
	11732: uint16(anon_sym_EQ),
	11733: uint16(679),
	11734: uint16(1),
	11735: uint16(anon_sym_COMMA),
	11736: uint16(681),
	11737: uint16(1),
	11738: uint16(anon_sym_in),
	11739: uint16(200),
	11740: uint16(1),
	11741: uint16(sym_comment),
	11742: uint16(218),
	11743: uint16(1),
	11744: uint16(aux_sym__name_list_repeat1),
	11745: uint16(6),
	11746: uint16(3),
	11747: uint16(1),
	11748: uint16(anon_sym_DASH_DASH),
	11749: uint16(5),
	11750: uint16(1),
	11751: uint16(sym__block_comment_start),
	11752: uint16(683),
	11753: uint16(1),
	11754: uint16(sym_identifier),
	11755: uint16(201),
	11756: uint16(1),
	11757: uint16(sym_comment),
	11758: uint16(260),
	11759: uint16(1),
	11760: uint16(sym__name_list),
	11761: uint16(247),
	11762: uint16(2),
	11763: uint16(sym_for_generic_clause),
	11764: uint16(sym_for_numeric_clause),
	11765: uint16(6),
	11766: uint16(5),
	11767: uint16(1),
	11768: uint16(sym__block_comment_start),
	11769: uint16(654),
	11770: uint16(1),
	11771: uint16(anon_sym_DASH_DASH),
	11772: uint16(685),
	11773: uint16(1),
	11774: uint16(anon_sym_DQUOTE),
	11775: uint16(196),
	11776: uint16(1),
	11777: uint16(aux_sym__doublequote_string_content),
	11778: uint16(202),
	11779: uint16(1),
	11780: uint16(sym_comment),
	11781: uint16(665),
	11782: uint16(2),
	11783: uint16(aux_sym__doublequote_string_content_token1),
	11784: uint16(sym_escape_sequence),
	11785: uint16(6),
	11786: uint16(5),
	11787: uint16(1),
	11788: uint16(sym__block_comment_start),
	11789: uint16(654),
	11790: uint16(1),
	11791: uint16(anon_sym_DASH_DASH),
	11792: uint16(685),
	11793: uint16(1),
	11794: uint16(anon_sym_SQUOTE),
	11795: uint16(197),
	11796: uint16(1),
	11797: uint16(aux_sym__singlequote_string_content),
	11798: uint16(203),
	11799: uint16(1),
	11800: uint16(sym_comment),
	11801: uint16(652),
	11802: uint16(2),
	11803: uint16(aux_sym__singlequote_string_content_token1),
	11804: uint16(sym_escape_sequence),
	11805: uint16(6),
	11806: uint16(5),
	11807: uint16(1),
	11808: uint16(sym__block_comment_start),
	11809: uint16(654),
	11810: uint16(1),
	11811: uint16(anon_sym_DASH_DASH),
	11812: uint16(687),
	11813: uint16(1),
	11814: uint16(anon_sym_DQUOTE),
	11815: uint16(204),
	11816: uint16(1),
	11817: uint16(sym_comment),
	11818: uint16(206),
	11819: uint16(1),
	11820: uint16(aux_sym__doublequote_string_content),
	11821: uint16(665),
	11822: uint16(2),
	11823: uint16(aux_sym__doublequote_string_content_token1),
	11824: uint16(sym_escape_sequence),
	11825: uint16(6),
	11826: uint16(5),
	11827: uint16(1),
	11828: uint16(sym__block_comment_start),
	11829: uint16(654),
	11830: uint16(1),
	11831: uint16(anon_sym_DASH_DASH),
	11832: uint16(687),
	11833: uint16(1),
	11834: uint16(anon_sym_SQUOTE),
	11835: uint16(193),
	11836: uint16(1),
	11837: uint16(aux_sym__singlequote_string_content),
	11838: uint16(205),
	11839: uint16(1),
	11840: uint16(sym_comment),
	11841: uint16(652),
	11842: uint16(2),
	11843: uint16(aux_sym__singlequote_string_content_token1),
	11844: uint16(sym_escape_sequence),
	11845: uint16(6),
	11846: uint16(5),
	11847: uint16(1),
	11848: uint16(sym__block_comment_start),
	11849: uint16(650),
	11850: uint16(1),
	11851: uint16(anon_sym_DQUOTE),
	11852: uint16(654),
	11853: uint16(1),
	11854: uint16(anon_sym_DASH_DASH),
	11855: uint16(198),
	11856: uint16(1),
	11857: uint16(aux_sym__doublequote_string_content),
	11858: uint16(206),
	11859: uint16(1),
	11860: uint16(sym_comment),
	11861: uint16(665),
	11862: uint16(2),
	11863: uint16(aux_sym__doublequote_string_content_token1),
	11864: uint16(sym_escape_sequence),
	11865: uint16(5),
	11866: uint16(3),
	11867: uint16(1),
	11868: uint16(anon_sym_DASH_DASH),
	11869: uint16(5),
	11870: uint16(1),
	11871: uint16(sym__block_comment_start),
	11872: uint16(689),
	11873: uint16(1),
	11874: uint16(anon_sym_EQ),
	11875: uint16(691),
	11876: uint16(1),
	11877: uint16(anon_sym_COMMA),
	11878: uint16(207),
	11879: uint16(2),
	11880: uint16(sym_comment),
	11881: uint16(aux_sym__variable_assignment_varlist_repeat1),
	11882: uint16(5),
	11883: uint16(5),
	11884: uint16(1),
	11885: uint16(sym__block_comment_start),
	11886: uint16(654),
	11887: uint16(1),
	11888: uint16(anon_sym_DASH_DASH),
	11889: uint16(694),
	11890: uint16(1),
	11891: uint16(anon_sym_SQUOTE),
	11892: uint16(208),
	11893: uint16(1),
	11894: uint16(sym_comment),
	11895: uint16(696),
	11896: uint16(2),
	11897: uint16(aux_sym__singlequote_string_content_token1),
	11898: uint16(sym_escape_sequence),
	11899: uint16(6),
	11900: uint16(3),
	11901: uint16(1),
	11902: uint16(anon_sym_DASH_DASH),
	11903: uint16(5),
	11904: uint16(1),
	11905: uint16(sym__block_comment_start),
	11906: uint16(698),
	11907: uint16(1),
	11908: uint16(anon_sym_COMMA),
	11909: uint16(700),
	11910: uint16(1),
	11911: uint16(anon_sym_RPAREN),
	11912: uint16(209),
	11913: uint16(1),
	11914: uint16(sym_comment),
	11915: uint16(215),
	11916: uint16(1),
	11917: uint16(aux_sym__name_list_repeat1),
	11918: uint16(5),
	11919: uint16(5),
	11920: uint16(1),
	11921: uint16(sym__block_comment_start),
	11922: uint16(654),
	11923: uint16(1),
	11924: uint16(anon_sym_DASH_DASH),
	11925: uint16(702),
	11926: uint16(1),
	11927: uint16(anon_sym_DQUOTE),
	11928: uint16(210),
	11929: uint16(1),
	11930: uint16(sym_comment),
	11931: uint16(704),
	11932: uint16(2),
	11933: uint16(aux_sym__doublequote_string_content_token1),
	11934: uint16(sym_escape_sequence),
	11935: uint16(4),
	11936: uint16(3),
	11937: uint16(1),
	11938: uint16(anon_sym_DASH_DASH),
	11939: uint16(5),
	11940: uint16(1),
	11941: uint16(sym__block_comment_start),
	11942: uint16(211),
	11943: uint16(1),
	11944: uint16(sym_comment),
	11945: uint16(619),
	11946: uint16(3),
	11947: uint16(anon_sym_SEMI),
	11948: uint16(anon_sym_COMMA),
	11949: uint16(anon_sym_RBRACE),
	11950: uint16(6),
	11951: uint16(3),
	11952: uint16(1),
	11953: uint16(anon_sym_DASH_DASH),
	11954: uint16(5),
	11955: uint16(1),
	11956: uint16(sym__block_comment_start),
	11957: uint16(305),
	11958: uint16(1),
	11959: uint16(anon_sym_COMMA),
	11960: uint16(706),
	11961: uint16(1),
	11962: uint16(anon_sym_RPAREN),
	11963: uint16(173),
	11964: uint16(1),
	11965: uint16(aux_sym__expression_list_repeat1),
	11966: uint16(212),
	11967: uint16(1),
	11968: uint16(sym_comment),
	11969: uint16(5),
	11970: uint16(3),
	11971: uint16(1),
	11972: uint16(anon_sym_DASH_DASH),
	11973: uint16(5),
	11974: uint16(1),
	11975: uint16(sym__block_comment_start),
	11976: uint16(710),
	11977: uint16(1),
	11978: uint16(anon_sym_else),
	11979: uint16(213),
	11980: uint16(1),
	11981: uint16(sym_comment),
	11982: uint16(708),
	11983: uint16(2),
	11984: uint16(anon_sym_end),
	11985: uint16(anon_sym_elseif),
	11986: uint16(6),
	11987: uint16(3),
	11988: uint16(1),
	11989: uint16(anon_sym_DASH_DASH),
	11990: uint16(5),
	11991: uint16(1),
	11992: uint16(sym__block_comment_start),
	11993: uint16(642),
	11994: uint16(1),
	11995: uint16(sym_vararg_expression),
	11996: uint16(712),
	11997: uint16(1),
	11998: uint16(sym_identifier),
	11999: uint16(214),
	12000: uint16(1),
	12001: uint16(sym_comment),
	12002: uint16(248),
	12003: uint16(1),
	12004: uint16(sym__vararg_parameter),
	12005: uint16(6),
	12006: uint16(3),
	12007: uint16(1),
	12008: uint16(anon_sym_DASH_DASH),
	12009: uint16(5),
	12010: uint16(1),
	12011: uint16(sym__block_comment_start),
	12012: uint16(714),
	12013: uint16(1),
	12014: uint16(anon_sym_COMMA),
	12015: uint16(716),
	12016: uint16(1),
	12017: uint16(anon_sym_RPAREN),
	12018: uint16(195),
	12019: uint16(1),
	12020: uint16(aux_sym__name_list_repeat1),
	12021: uint16(215),
	12022: uint16(1),
	12023: uint16(sym_comment),
	12024: uint16(6),
	12025: uint16(3),
	12026: uint16(1),
	12027: uint16(anon_sym_DASH_DASH),
	12028: uint16(5),
	12029: uint16(1),
	12030: uint16(sym__block_comment_start),
	12031: uint16(718),
	12032: uint16(1),
	12033: uint16(anon_sym_LPAREN),
	12034: uint16(50),
	12035: uint16(1),
	12036: uint16(sym_parameters),
	12037: uint16(147),
	12038: uint16(1),
	12039: uint16(sym__function_body),
	12040: uint16(216),
	12041: uint16(1),
	12042: uint16(sym_comment),
	12043: uint16(6),
	12044: uint16(3),
	12045: uint16(1),
	12046: uint16(anon_sym_DASH_DASH),
	12047: uint16(5),
	12048: uint16(1),
	12049: uint16(sym__block_comment_start),
	12050: uint16(589),
	12051: uint16(1),
	12052: uint16(anon_sym_COMMA),
	12053: uint16(720),
	12054: uint16(1),
	12055: uint16(anon_sym_EQ),
	12056: uint16(207),
	12057: uint16(1),
	12058: uint16(aux_sym__variable_assignment_varlist_repeat1),
	12059: uint16(217),
	12060: uint16(1),
	12061: uint16(sym_comment),
	12062: uint16(6),
	12063: uint16(3),
	12064: uint16(1),
	12065: uint16(anon_sym_DASH_DASH),
	12066: uint16(5),
	12067: uint16(1),
	12068: uint16(sym__block_comment_start),
	12069: uint16(679),
	12070: uint16(1),
	12071: uint16(anon_sym_COMMA),
	12072: uint16(722),
	12073: uint16(1),
	12074: uint16(anon_sym_in),
	12075: uint16(195),
	12076: uint16(1),
	12077: uint16(aux_sym__name_list_repeat1),
	12078: uint16(218),
	12079: uint16(1),
	12080: uint16(sym_comment),
	12081: uint16(6),
	12082: uint16(3),
	12083: uint16(1),
	12084: uint16(anon_sym_DASH_DASH),
	12085: uint16(5),
	12086: uint16(1),
	12087: uint16(sym__block_comment_start),
	12088: uint16(718),
	12089: uint16(1),
	12090: uint16(anon_sym_LPAREN),
	12091: uint16(43),
	12092: uint16(1),
	12093: uint16(sym__function_body),
	12094: uint16(50),
	12095: uint16(1),
	12096: uint16(sym_parameters),
	12097: uint16(219),
	12098: uint16(1),
	12099: uint16(sym_comment),
	12100: uint16(4),
	12101: uint16(3),
	12102: uint16(1),
	12103: uint16(anon_sym_DASH_DASH),
	12104: uint16(5),
	12105: uint16(1),
	12106: uint16(sym__block_comment_start),
	12107: uint16(220),
	12108: uint16(1),
	12109: uint16(sym_comment),
	12110: uint16(724),
	12111: uint16(3),
	12112: uint16(anon_sym_COMMA),
	12113: uint16(anon_sym_in),
	12114: uint16(anon_sym_RPAREN),
	12115: uint16(6),
	12116: uint16(3),
	12117: uint16(1),
	12118: uint16(anon_sym_DASH_DASH),
	12119: uint16(5),
	12120: uint16(1),
	12121: uint16(sym__block_comment_start),
	12122: uint16(642),
	12123: uint16(1),
	12124: uint16(sym_vararg_expression),
	12125: uint16(712),
	12126: uint16(1),
	12127: uint16(sym_identifier),
	12128: uint16(221),
	12129: uint16(1),
	12130: uint16(sym_comment),
	12131: uint16(257),
	12132: uint16(1),
	12133: uint16(sym__vararg_parameter),
	12134: uint16(6),
	12135: uint16(3),
	12136: uint16(1),
	12137: uint16(anon_sym_DASH_DASH),
	12138: uint16(5),
	12139: uint16(1),
	12140: uint16(sym__block_comment_start),
	12141: uint16(718),
	12142: uint16(1),
	12143: uint16(anon_sym_LPAREN),
	12144: uint16(50),
	12145: uint16(1),
	12146: uint16(sym_parameters),
	12147: uint16(133),
	12148: uint16(1),
	12149: uint16(sym__function_body),
	12150: uint16(222),
	12151: uint16(1),
	12152: uint16(sym_comment),
	12153: uint16(6),
	12154: uint16(3),
	12155: uint16(1),
	12156: uint16(anon_sym_DASH_DASH),
	12157: uint16(5),
	12158: uint16(1),
	12159: uint16(sym__block_comment_start),
	12160: uint16(718),
	12161: uint16(1),
	12162: uint16(anon_sym_LPAREN),
	12163: uint16(50),
	12164: uint16(1),
	12165: uint16(sym_parameters),
	12166: uint16(150),
	12167: uint16(1),
	12168: uint16(sym__function_body),
	12169: uint16(223),
	12170: uint16(1),
	12171: uint16(sym_comment),
	12172: uint16(5),
	12173: uint16(3),
	12174: uint16(1),
	12175: uint16(anon_sym_DASH_DASH),
	12176: uint16(5),
	12177: uint16(1),
	12178: uint16(sym__block_comment_start),
	12179: uint16(728),
	12180: uint16(1),
	12181: uint16(anon_sym_else),
	12182: uint16(224),
	12183: uint16(1),
	12184: uint16(sym_comment),
	12185: uint16(726),
	12186: uint16(2),
	12187: uint16(anon_sym_end),
	12188: uint16(anon_sym_elseif),
	12189: uint16(4),
	12190: uint16(3),
	12191: uint16(1),
	12192: uint16(anon_sym_DASH_DASH),
	12193: uint16(5),
	12194: uint16(1),
	12195: uint16(sym__block_comment_start),
	12196: uint16(225),
	12197: uint16(1),
	12198: uint16(sym_comment),
	12199: uint16(730),
	12200: uint16(3),
	12201: uint16(anon_sym_DOT),
	12202: uint16(anon_sym_COLON),
	12203: uint16(anon_sym_LPAREN),
	12204: uint16(4),
	12205: uint16(3),
	12206: uint16(1),
	12207: uint16(anon_sym_DASH_DASH),
	12208: uint16(5),
	12209: uint16(1),
	12210: uint16(sym__block_comment_start),
	12211: uint16(226),
	12212: uint16(1),
	12213: uint16(sym_comment),
	12214: uint16(732),
	12215: uint16(3),
	12216: uint16(anon_sym_DOT),
	12217: uint16(anon_sym_COLON),
	12218: uint16(anon_sym_LPAREN),
	12219: uint16(6),
	12220: uint16(3),
	12221: uint16(1),
	12222: uint16(anon_sym_DASH_DASH),
	12223: uint16(5),
	12224: uint16(1),
	12225: uint16(sym__block_comment_start),
	12226: uint16(734),
	12227: uint16(1),
	12228: uint16(anon_sym_DOT),
	12229: uint16(736),
	12230: uint16(1),
	12231: uint16(anon_sym_COLON),
	12232: uint16(738),
	12233: uint16(1),
	12234: uint16(anon_sym_LPAREN),
	12235: uint16(227),
	12236: uint16(1),
	12237: uint16(sym_comment),
	12238: uint16(4),
	12239: uint16(3),
	12240: uint16(1),
	12241: uint16(anon_sym_DASH_DASH),
	12242: uint16(5),
	12243: uint16(1),
	12244: uint16(sym__block_comment_start),
	12245: uint16(228),
	12246: uint16(1),
	12247: uint16(sym_comment),
	12248: uint16(740),
	12249: uint16(3),
	12250: uint16(anon_sym_DOT),
	12251: uint16(anon_sym_COLON),
	12252: uint16(anon_sym_LPAREN),
	12253: uint16(5),
	12254: uint16(3),
	12255: uint16(1),
	12256: uint16(anon_sym_DASH_DASH),
	12257: uint16(5),
	12258: uint16(1),
	12259: uint16(sym__block_comment_start),
	12260: uint16(742),
	12261: uint16(1),
	12262: uint16(sym_identifier),
	12263: uint16(744),
	12264: uint16(1),
	12265: uint16(anon_sym_RPAREN),
	12266: uint16(229),
	12267: uint16(1),
	12268: uint16(sym_comment),
	12269: uint16(5),
	12270: uint16(3),
	12271: uint16(1),
	12272: uint16(anon_sym_DASH_DASH),
	12273: uint16(5),
	12274: uint16(1),
	12275: uint16(sym__block_comment_start),
	12276: uint16(746),
	12277: uint16(1),
	12278: uint16(sym_identifier),
	12279: uint16(748),
	12280: uint16(1),
	12281: uint16(anon_sym_STAR),
	12282: uint16(230),
	12283: uint16(1),
	12284: uint16(sym_comment),
	12285: uint16(4),
	12286: uint16(3),
	12287: uint16(1),
	12288: uint16(anon_sym_DASH_DASH),
	12289: uint16(5),
	12290: uint16(1),
	12291: uint16(sym__block_comment_start),
	12292: uint16(750),
	12293: uint16(1),
	12294: uint16(anon_sym_do),
	12295: uint16(231),
	12296: uint16(1),
	12297: uint16(sym_comment),
	12298: uint16(4),
	12299: uint16(3),
	12300: uint16(1),
	12301: uint16(anon_sym_DASH_DASH),
	12302: uint16(5),
	12303: uint16(1),
	12304: uint16(sym__block_comment_start),
	12305: uint16(752),
	12306: uint16(1),
	12307: uint16(sym__block_string_end),
	12308: uint16(232),
	12309: uint16(1),
	12310: uint16(sym_comment),
	12311: uint16(4),
	12312: uint16(3),
	12313: uint16(1),
	12314: uint16(anon_sym_DASH_DASH),
	12315: uint16(5),
	12316: uint16(1),
	12317: uint16(sym__block_comment_start),
	12318: uint16(754),
	12319: uint16(1),
	12320: uint16(anon_sym_end),
	12321: uint16(233),
	12322: uint16(1),
	12323: uint16(sym_comment),
	12324: uint16(4),
	12325: uint16(3),
	12326: uint16(1),
	12327: uint16(anon_sym_DASH_DASH),
	12328: uint16(5),
	12329: uint16(1),
	12330: uint16(sym__block_comment_start),
	12331: uint16(756),
	12332: uint16(1),
	12333: uint16(anon_sym_end),
	12334: uint16(234),
	12335: uint16(1),
	12336: uint16(sym_comment),
	12337: uint16(4),
	12338: uint16(3),
	12339: uint16(1),
	12340: uint16(anon_sym_DASH_DASH),
	12341: uint16(5),
	12342: uint16(1),
	12343: uint16(sym__block_comment_start),
	12344: uint16(758),
	12345: uint16(1),
	12346: uint16(sym_identifier),
	12347: uint16(235),
	12348: uint16(1),
	12349: uint16(sym_comment),
	12350: uint16(4),
	12351: uint16(3),
	12352: uint16(1),
	12353: uint16(anon_sym_DASH_DASH),
	12354: uint16(5),
	12355: uint16(1),
	12356: uint16(sym__block_comment_start),
	12357: uint16(760),
	12358: uint16(1),
	12359: uint16(anon_sym_end),
	12360: uint16(236),
	12361: uint16(1),
	12362: uint16(sym_comment),
	12363: uint16(4),
	12364: uint16(3),
	12365: uint16(1),
	12366: uint16(anon_sym_DASH_DASH),
	12367: uint16(5),
	12368: uint16(1),
	12369: uint16(sym__block_comment_start),
	12370: uint16(762),
	12371: uint16(1),
	12372: uint16(anon_sym_GT),
	12373: uint16(237),
	12374: uint16(1),
	12375: uint16(sym_comment),
	12376: uint16(4),
	12377: uint16(3),
	12378: uint16(1),
	12379: uint16(anon_sym_DASH_DASH),
	12380: uint16(5),
	12381: uint16(1),
	12382: uint16(sym__block_comment_start),
	12383: uint16(764),
	12384: uint16(1),
	12385: uint16(sym_identifier),
	12386: uint16(238),
	12387: uint16(1),
	12388: uint16(sym_comment),
	12389: uint16(4),
	12390: uint16(3),
	12391: uint16(1),
	12392: uint16(anon_sym_DASH_DASH),
	12393: uint16(5),
	12394: uint16(1),
	12395: uint16(sym__block_comment_start),
	12396: uint16(746),
	12397: uint16(1),
	12398: uint16(sym_identifier),
	12399: uint16(239),
	12400: uint16(1),
	12401: uint16(sym_comment),
	12402: uint16(4),
	12403: uint16(3),
	12404: uint16(1),
	12405: uint16(anon_sym_DASH_DASH),
	12406: uint16(5),
	12407: uint16(1),
	12408: uint16(sym__block_comment_start),
	12409: uint16(766),
	12410: uint16(1),
	12411: uint16(anon_sym_end),
	12412: uint16(240),
	12413: uint16(1),
	12414: uint16(sym_comment),
	12415: uint16(4),
	12416: uint16(3),
	12417: uint16(1),
	12418: uint16(anon_sym_DASH_DASH),
	12419: uint16(5),
	12420: uint16(1),
	12421: uint16(sym__block_comment_start),
	12422: uint16(768),
	12423: uint16(1),
	12424: uint16(anon_sym_until),
	12425: uint16(241),
	12426: uint16(1),
	12427: uint16(sym_comment),
	12428: uint16(4),
	12429: uint16(3),
	12430: uint16(1),
	12431: uint16(anon_sym_DASH_DASH),
	12432: uint16(5),
	12433: uint16(1),
	12434: uint16(sym__block_comment_start),
	12435: uint16(770),
	12436: uint16(1),
	12437: uint16(anon_sym_RBRACE),
	12438: uint16(242),
	12439: uint16(1),
	12440: uint16(sym_comment),
	12441: uint16(4),
	12442: uint16(3),
	12443: uint16(1),
	12444: uint16(anon_sym_DASH_DASH),
	12445: uint16(5),
	12446: uint16(1),
	12447: uint16(sym__block_comment_start),
	12448: uint16(772),
	12449: uint16(1),
	12450: uint16(sym_identifier),
	12451: uint16(243),
	12452: uint16(1),
	12453: uint16(sym_comment),
	12454: uint16(4),
	12455: uint16(3),
	12456: uint16(1),
	12457: uint16(anon_sym_DASH_DASH),
	12458: uint16(5),
	12459: uint16(1),
	12460: uint16(sym__block_comment_start),
	12461: uint16(774),
	12462: uint16(1),
	12463: uint16(sym_identifier),
	12464: uint16(244),
	12465: uint16(1),
	12466: uint16(sym_comment),
	12467: uint16(4),
	12468: uint16(3),
	12469: uint16(1),
	12470: uint16(anon_sym_DASH_DASH),
	12471: uint16(5),
	12472: uint16(1),
	12473: uint16(sym__block_comment_start),
	12474: uint16(776),
	12475: uint16(1),
	12477: uint16(245),
	12478: uint16(1),
	12479: uint16(sym_comment),
	12480: uint16(4),
	12481: uint16(3),
	12482: uint16(1),
	12483: uint16(anon_sym_DASH_DASH),
	12484: uint16(5),
	12485: uint16(1),
	12486: uint16(sym__block_comment_start),
	12487: uint16(778),
	12488: uint16(1),
	12489: uint16(anon_sym_RPAREN),
	12490: uint16(246),
	12491: uint16(1),
	12492: uint16(sym_comment),
	12493: uint16(4),
	12494: uint16(3),
	12495: uint16(1),
	12496: uint16(anon_sym_DASH_DASH),
	12497: uint16(5),
	12498: uint16(1),
	12499: uint16(sym__block_comment_start),
	12500: uint16(780),
	12501: uint16(1),
	12502: uint16(anon_sym_do),
	12503: uint16(247),
	12504: uint16(1),
	12505: uint16(sym_comment),
	12506: uint16(4),
	12507: uint16(3),
	12508: uint16(1),
	12509: uint16(anon_sym_DASH_DASH),
	12510: uint16(5),
	12511: uint16(1),
	12512: uint16(sym__block_comment_start),
	12513: uint16(782),
	12514: uint16(1),
	12515: uint16(anon_sym_RPAREN),
	12516: uint16(248),
	12517: uint16(1),
	12518: uint16(sym_comment),
	12519: uint16(4),
	12520: uint16(3),
	12521: uint16(1),
	12522: uint16(anon_sym_DASH_DASH),
	12523: uint16(5),
	12524: uint16(1),
	12525: uint16(sym__block_comment_start),
	12526: uint16(784),
	12527: uint16(1),
	12528: uint16(anon_sym_RPAREN),
	12529: uint16(249),
	12530: uint16(1),
	12531: uint16(sym_comment),
	12532: uint16(4),
	12533: uint16(3),
	12534: uint16(1),
	12535: uint16(anon_sym_DASH_DASH),
	12536: uint16(5),
	12537: uint16(1),
	12538: uint16(sym__block_comment_start),
	12539: uint16(786),
	12540: uint16(1),
	12541: uint16(anon_sym_LPAREN),
	12542: uint16(250),
	12543: uint16(1),
	12544: uint16(sym_comment),
	12545: uint16(4),
	12546: uint16(3),
	12547: uint16(1),
	12548: uint16(anon_sym_DASH_DASH),
	12549: uint16(5),
	12550: uint16(1),
	12551: uint16(sym__block_comment_start),
	12552: uint16(788),
	12553: uint16(1),
	12554: uint16(sym_identifier),
	12555: uint16(251),
	12556: uint16(1),
	12557: uint16(sym_comment),
	12558: uint16(4),
	12559: uint16(3),
	12560: uint16(1),
	12561: uint16(anon_sym_DASH_DASH),
	12562: uint16(5),
	12563: uint16(1),
	12564: uint16(sym__block_comment_start),
	12565: uint16(790),
	12566: uint16(1),
	12567: uint16(sym__block_comment_content),
	12568: uint16(252),
	12569: uint16(1),
	12570: uint16(sym_comment),
	12571: uint16(4),
	12572: uint16(3),
	12573: uint16(1),
	12574: uint16(anon_sym_DASH_DASH),
	12575: uint16(5),
	12576: uint16(1),
	12577: uint16(sym__block_comment_start),
	12578: uint16(792),
	12579: uint16(1),
	12580: uint16(sym__block_comment_end),
	12581: uint16(253),
	12582: uint16(1),
	12583: uint16(sym_comment),
	12584: uint16(4),
	12585: uint16(3),
	12586: uint16(1),
	12587: uint16(anon_sym_DASH_DASH),
	12588: uint16(5),
	12589: uint16(1),
	12590: uint16(sym__block_comment_start),
	12591: uint16(794),
	12592: uint16(1),
	12593: uint16(anon_sym_end),
	12594: uint16(254),
	12595: uint16(1),
	12596: uint16(sym_comment),
	12597: uint16(4),
	12598: uint16(3),
	12599: uint16(1),
	12600: uint16(anon_sym_DASH_DASH),
	12601: uint16(5),
	12602: uint16(1),
	12603: uint16(sym__block_comment_start),
	12604: uint16(796),
	12605: uint16(1),
	12606: uint16(anon_sym_EQ),
	12607: uint16(255),
	12608: uint16(1),
	12609: uint16(sym_comment),
	12610: uint16(4),
	12611: uint16(3),
	12612: uint16(1),
	12613: uint16(anon_sym_DASH_DASH),
	12614: uint16(5),
	12615: uint16(1),
	12616: uint16(sym__block_comment_start),
	12617: uint16(798),
	12618: uint16(1),
	12619: uint16(sym_identifier),
	12620: uint16(256),
	12621: uint16(1),
	12622: uint16(sym_comment),
	12623: uint16(4),
	12624: uint16(3),
	12625: uint16(1),
	12626: uint16(anon_sym_DASH_DASH),
	12627: uint16(5),
	12628: uint16(1),
	12629: uint16(sym__block_comment_start),
	12630: uint16(800),
	12631: uint16(1),
	12632: uint16(anon_sym_RPAREN),
	12633: uint16(257),
	12634: uint16(1),
	12635: uint16(sym_comment),
	12636: uint16(4),
	12637: uint16(3),
	12638: uint16(1),
	12639: uint16(anon_sym_DASH_DASH),
	12640: uint16(5),
	12641: uint16(1),
	12642: uint16(sym__block_comment_start),
	12643: uint16(802),
	12644: uint16(1),
	12645: uint16(sym_identifier),
	12646: uint16(258),
	12647: uint16(1),
	12648: uint16(sym_comment),
	12649: uint16(4),
	12650: uint16(3),
	12651: uint16(1),
	12652: uint16(anon_sym_DASH_DASH),
	12653: uint16(5),
	12654: uint16(1),
	12655: uint16(sym__block_comment_start),
	12656: uint16(804),
	12657: uint16(1),
	12658: uint16(sym_identifier),
	12659: uint16(259),
	12660: uint16(1),
	12661: uint16(sym_comment),
	12662: uint16(4),
	12663: uint16(3),
	12664: uint16(1),
	12665: uint16(anon_sym_DASH_DASH),
	12666: uint16(5),
	12667: uint16(1),
	12668: uint16(sym__block_comment_start),
	12669: uint16(806),
	12670: uint16(1),
	12671: uint16(anon_sym_in),
	12672: uint16(260),
	12673: uint16(1),
	12674: uint16(sym_comment),
	12675: uint16(4),
	12676: uint16(3),
	12677: uint16(1),
	12678: uint16(anon_sym_DASH_DASH),
	12679: uint16(5),
	12680: uint16(1),
	12681: uint16(sym__block_comment_start),
	12682: uint16(229),
	12683: uint16(1),
	12685: uint16(261),
	12686: uint16(1),
	12687: uint16(sym_comment),
	12688: uint16(4),
	12689: uint16(3),
	12690: uint16(1),
	12691: uint16(anon_sym_DASH_DASH),
	12692: uint16(5),
	12693: uint16(1),
	12694: uint16(sym__block_comment_start),
	12695: uint16(808),
	12696: uint16(1),
	12697: uint16(sym_identifier),
	12698: uint16(262),
	12699: uint16(1),
	12700: uint16(sym_comment),
	12701: uint16(4),
	12702: uint16(3),
	12703: uint16(1),
	12704: uint16(anon_sym_DASH_DASH),
	12705: uint16(5),
	12706: uint16(1),
	12707: uint16(sym__block_comment_start),
	12708: uint16(810),
	12709: uint16(1),
	12711: uint16(263),
	12712: uint16(1),
	12713: uint16(sym_comment),
	12714: uint16(4),
	12715: uint16(5),
	12716: uint16(1),
	12717: uint16(sym__block_comment_start),
	12718: uint16(654),
	12719: uint16(1),
	12720: uint16(anon_sym_DASH_DASH),
	12721: uint16(812),
	12722: uint16(1),
	12723: uint16(aux_sym_comment_token1),
	12724: uint16(264),
	12725: uint16(1),
	12726: uint16(sym_comment),
	12727: uint16(4),
	12728: uint16(3),
	12729: uint16(1),
	12730: uint16(anon_sym_DASH_DASH),
	12731: uint16(5),
	12732: uint16(1),
	12733: uint16(sym__block_comment_start),
	12734: uint16(814),
	12735: uint16(1),
	12736: uint16(anon_sym_COLON_COLON),
	12737: uint16(265),
	12738: uint16(1),
	12739: uint16(sym_comment),
	12740: uint16(4),
	12741: uint16(3),
	12742: uint16(1),
	12743: uint16(anon_sym_DASH_DASH),
	12744: uint16(5),
	12745: uint16(1),
	12746: uint16(sym__block_comment_start),
	12747: uint16(231),
	12748: uint16(1),
	12750: uint16(266),
	12751: uint16(1),
	12752: uint16(sym_comment),
	12753: uint16(4),
	12754: uint16(3),
	12755: uint16(1),
	12756: uint16(anon_sym_DASH_DASH),
	12757: uint16(5),
	12758: uint16(1),
	12759: uint16(sym__block_comment_start),
	12760: uint16(816),
	12761: uint16(1),
	12762: uint16(anon_sym_EQ),
	12763: uint16(267),
	12764: uint16(1),
	12765: uint16(sym_comment),
	12766: uint16(4),
	12767: uint16(3),
	12768: uint16(1),
	12769: uint16(anon_sym_DASH_DASH),
	12770: uint16(5),
	12771: uint16(1),
	12772: uint16(sym__block_comment_start),
	12773: uint16(712),
	12774: uint16(1),
	12775: uint16(sym_identifier),
	12776: uint16(268),
	12777: uint16(1),
	12778: uint16(sym_comment),
	12779: uint16(4),
	12780: uint16(3),
	12781: uint16(1),
	12782: uint16(anon_sym_DASH_DASH),
	12783: uint16(5),
	12784: uint16(1),
	12785: uint16(sym__block_comment_start),
	12786: uint16(818),
	12787: uint16(1),
	12788: uint16(anon_sym_end),
	12789: uint16(269),
	12790: uint16(1),
	12791: uint16(sym_comment),
	12792: uint16(4),
	12793: uint16(3),
	12794: uint16(1),
	12795: uint16(anon_sym_DASH_DASH),
	12796: uint16(5),
	12797: uint16(1),
	12798: uint16(sym__block_comment_start),
	12799: uint16(820),
	12800: uint16(1),
	12801: uint16(anon_sym_end),
	12802: uint16(270),
	12803: uint16(1),
	12804: uint16(sym_comment),
	12805: uint16(4),
	12806: uint16(3),
	12807: uint16(1),
	12808: uint16(anon_sym_DASH_DASH),
	12809: uint16(5),
	12810: uint16(1),
	12811: uint16(sym__block_comment_start),
	12812: uint16(822),
	12813: uint16(1),
	12814: uint16(anon_sym_end),
	12815: uint16(271),
	12816: uint16(1),
	12817: uint16(sym_comment),
	12818: uint16(4),
	12819: uint16(3),
	12820: uint16(1),
	12821: uint16(anon_sym_DASH_DASH),
	12822: uint16(5),
	12823: uint16(1),
	12824: uint16(sym__block_comment_start),
	12825: uint16(824),
	12826: uint16(1),
	12827: uint16(sym_identifier),
	12828: uint16(272),
	12829: uint16(1),
	12830: uint16(sym_comment),
	12831: uint16(4),
	12832: uint16(3),
	12833: uint16(1),
	12834: uint16(anon_sym_DASH_DASH),
	12835: uint16(5),
	12836: uint16(1),
	12837: uint16(sym__block_comment_start),
	12838: uint16(826),
	12839: uint16(1),
	12840: uint16(anon_sym_RPAREN),
	12841: uint16(273),
	12842: uint16(1),
	12843: uint16(sym_comment),
	12844: uint16(4),
	12845: uint16(3),
	12846: uint16(1),
	12847: uint16(anon_sym_DASH_DASH),
	12848: uint16(5),
	12849: uint16(1),
	12850: uint16(sym__block_comment_start),
	12851: uint16(828),
	12852: uint16(1),
	12853: uint16(sym__block_string_end),
	12854: uint16(274),
	12855: uint16(1),
	12856: uint16(sym_comment),
	12857: uint16(4),
	12858: uint16(3),
	12859: uint16(1),
	12860: uint16(anon_sym_DASH_DASH),
	12861: uint16(5),
	12862: uint16(1),
	12863: uint16(sym__block_comment_start),
	12864: uint16(830),
	12865: uint16(1),
	12866: uint16(sym__block_string_content),
	12867: uint16(275),
	12868: uint16(1),
	12869: uint16(sym_comment),
	12870: uint16(4),
	12871: uint16(3),
	12872: uint16(1),
	12873: uint16(anon_sym_DASH_DASH),
	12874: uint16(5),
	12875: uint16(1),
	12876: uint16(sym__block_comment_start),
	12877: uint16(832),
	12878: uint16(1),
	12879: uint16(anon_sym_LPAREN),
	12880: uint16(276),
	12881: uint16(1),
	12882: uint16(sym_comment),
	12883: uint16(4),
	12884: uint16(3),
	12885: uint16(1),
	12886: uint16(anon_sym_DASH_DASH),
	12887: uint16(5),
	12888: uint16(1),
	12889: uint16(sym__block_comment_start),
	12890: uint16(834),
	12891: uint16(1),
	12892: uint16(anon_sym_RBRACE),
	12893: uint16(277),
	12894: uint16(1),
	12895: uint16(sym_comment),
	12896: uint16(4),
	12897: uint16(3),
	12898: uint16(1),
	12899: uint16(anon_sym_DASH_DASH),
	12900: uint16(5),
	12901: uint16(1),
	12902: uint16(sym__block_comment_start),
	12903: uint16(836),
	12904: uint16(1),
	12905: uint16(sym__block_string_content),
	12906: uint16(278),
	12907: uint16(1),
	12908: uint16(sym_comment),
	12909: uint16(4),
	12910: uint16(3),
	12911: uint16(1),
	12912: uint16(anon_sym_DASH_DASH),
	12913: uint16(5),
	12914: uint16(1),
	12915: uint16(sym__block_comment_start),
	12916: uint16(838),
	12917: uint16(1),
	12918: uint16(anon_sym_end),
	12919: uint16(279),
	12920: uint16(1),
	12921: uint16(sym_comment),
	12922: uint16(1),
	12923: uint16(840),
	12924: uint16(1),
	12926: uint16(1),
	12927: uint16(842),
	12928: uint16(1),
}

var ts_small_parse_table_map = [280]uint32_t{
	1:   uint32(68),
	2:   uint32(136),
	3:   uint32(204),
	4:   uint32(272),
	5:   uint32(339),
	6:   uint32(410),
	7:   uint32(477),
	8:   uint32(544),
	9:   uint32(611),
	10:  uint32(678),
	11:  uint32(745),
	12:  uint32(812),
	13:  uint32(879),
	14:  uint32(946),
	15:  uint32(1013),
	16:  uint32(1080),
	17:  uint32(1147),
	18:  uint32(1214),
	19:  uint32(1354),
	20:  uint32(1479),
	21:  uint32(1600),
	22:  uint32(1727),
	23:  uint32(1787),
	24:  uint32(1863),
	25:  uint32(1947),
	26:  uint32(2033),
	27:  uint32(2095),
	28:  uint32(2155),
	29:  uint32(2215),
	30:  uint32(2275),
	31:  uint32(2355),
	32:  uint32(2433),
	33:  uint32(2507),
	34:  uint32(2579),
	35:  uint32(2645),
	36:  uint32(2717),
	37:  uint32(2779),
	38:  uint32(2839),
	39:  uint32(2899),
	40:  uint32(2959),
	41:  uint32(3019),
	42:  uint32(3079),
	43:  uint32(3139),
	44:  uint32(3199),
	45:  uint32(3261),
	46:  uint32(3386),
	47:  uint32(3511),
	48:  uint32(3636),
	49:  uint32(3761),
	50:  uint32(3886),
	51:  uint32(4011),
	52:  uint32(4133),
	53:  uint32(4255),
	54:  uint32(4377),
	55:  uint32(4466),
	56:  uint32(4550),
	57:  uint32(4633),
	58:  uint32(4716),
	59:  uint32(4811),
	60:  uint32(4908),
	61:  uint32(5005),
	62:  uint32(5097),
	63:  uint32(5191),
	64:  uint32(5285),
	65:  uint32(5376),
	66:  uint32(5462),
	67:  uint32(5548),
	68:  uint32(5634),
	69:  uint32(5720),
	70:  uint32(5803),
	71:  uint32(5886),
	72:  uint32(5969),
	73:  uint32(6052),
	74:  uint32(6135),
	75:  uint32(6184),
	76:  uint32(6267),
	77:  uint32(6350),
	78:  uint32(6433),
	79:  uint32(6516),
	80:  uint32(6599),
	81:  uint32(6682),
	82:  uint32(6765),
	83:  uint32(6848),
	84:  uint32(6931),
	85:  uint32(7014),
	86:  uint32(7097),
	87:  uint32(7180),
	88:  uint32(7263),
	89:  uint32(7346),
	90:  uint32(7429),
	91:  uint32(7512),
	92:  uint32(7595),
	93:  uint32(7678),
	94:  uint32(7761),
	95:  uint32(7844),
	96:  uint32(7927),
	97:  uint32(8010),
	98:  uint32(8082),
	99:  uint32(8158),
	100: uint32(8203),
	101: uint32(8250),
	102: uint32(8297),
	103: uint32(8365),
	104: uint32(8429),
	105: uint32(8493),
	106: uint32(8535),
	107: uint32(8599),
	108: uint32(8640),
	109: uint32(8681),
	110: uint32(8722),
	111: uint32(8763),
	112: uint32(8804),
	113: uint32(8841),
	114: uint32(8906),
	115: uint32(8947),
	116: uint32(8986),
	117: uint32(9048),
	118: uint32(9110),
	119: uint32(9172),
	120: uint32(9210),
	121: uint32(9246),
	122: uint32(9308),
	123: uint32(9370),
	124: uint32(9432),
	125: uint32(9494),
	126: uint32(9556),
	127: uint32(9596),
	128: uint32(9633),
	129: uint32(9670),
	130: uint32(9704),
	131: uint32(9738),
	132: uint32(9772),
	133: uint32(9806),
	134: uint32(9840),
	135: uint32(9874),
	136: uint32(9908),
	137: uint32(9942),
	138: uint32(9976),
	139: uint32(10010),
	140: uint32(10044),
	141: uint32(10078),
	142: uint32(10112),
	143: uint32(10146),
	144: uint32(10180),
	145: uint32(10214),
	146: uint32(10248),
	147: uint32(10282),
	148: uint32(10316),
	149: uint32(10350),
	150: uint32(10384),
	151: uint32(10418),
	152: uint32(10452),
	153: uint32(10486),
	154: uint32(10520),
	155: uint32(10554),
	156: uint32(10588),
	157: uint32(10622),
	158: uint32(10656),
	159: uint32(10690),
	160: uint32(10724),
	161: uint32(10758),
	162: uint32(10792),
	163: uint32(10826),
	164: uint32(10860),
	165: uint32(10893),
	166: uint32(10936),
	167: uint32(10966),
	168: uint32(10996),
	169: uint32(11043),
	170: uint32(11080),
	171: uint32(11109),
	172: uint32(11135),
	173: uint32(11173),
	174: uint32(11197),
	175: uint32(11224),
	176: uint32(11244),
	177: uint32(11272),
	178: uint32(11300),
	179: uint32(11328),
	180: uint32(11356),
	181: uint32(11378),
	182: uint32(11399),
	183: uint32(11422),
	184: uint32(11441),
	185: uint32(11464),
	186: uint32(11481),
	187: uint32(11504),
	188: uint32(11523),
	189: uint32(11548),
	190: uint32(11573),
	191: uint32(11591),
	192: uint32(11611),
	193: uint32(11629),
	194: uint32(11647),
	195: uint32(11667),
	196: uint32(11687),
	197: uint32(11705),
	198: uint32(11723),
	199: uint32(11745),
	200: uint32(11765),
	201: uint32(11785),
	202: uint32(11805),
	203: uint32(11825),
	204: uint32(11845),
	205: uint32(11865),
	206: uint32(11882),
	207: uint32(11899),
	208: uint32(11918),
	209: uint32(11935),
	210: uint32(11950),
	211: uint32(11969),
	212: uint32(11986),
	213: uint32(12005),
	214: uint32(12024),
	215: uint32(12043),
	216: uint32(12062),
	217: uint32(12081),
	218: uint32(12100),
	219: uint32(12115),
	220: uint32(12134),
	221: uint32(12153),
	222: uint32(12172),
	223: uint32(12189),
	224: uint32(12204),
	225: uint32(12219),
	226: uint32(12238),
	227: uint32(12253),
	228: uint32(12269),
	229: uint32(12285),
	230: uint32(12298),
	231: uint32(12311),
	232: uint32(12324),
	233: uint32(12337),
	234: uint32(12350),
	235: uint32(12363),
	236: uint32(12376),
	237: uint32(12389),
	238: uint32(12402),
	239: uint32(12415),
	240: uint32(12428),
	241: uint32(12441),
	242: uint32(12454),
	243: uint32(12467),
	244: uint32(12480),
	245: uint32(12493),
	246: uint32(12506),
	247: uint32(12519),
	248: uint32(12532),
	249: uint32(12545),
	250: uint32(12558),
	251: uint32(12571),
	252: uint32(12584),
	253: uint32(12597),
	254: uint32(12610),
	255: uint32(12623),
	256: uint32(12636),
	257: uint32(12649),
	258: uint32(12662),
	259: uint32(12675),
	260: uint32(12688),
	261: uint32(12701),
	262: uint32(12714),
	263: uint32(12727),
	264: uint32(12740),
	265: uint32(12753),
	266: uint32(12766),
	267: uint32(12779),
	268: uint32(12792),
	269: uint32(12805),
	270: uint32(12818),
	271: uint32(12831),
	272: uint32(12844),
	273: uint32(12857),
	274: uint32(12870),
	275: uint32(12883),
	276: uint32(12896),
	277: uint32(12909),
	278: uint32(12922),
	279: uint32(12926),
}

var ts_parse_actions = [844]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(264)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(252)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_chunk),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(272)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(262)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(73)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(190)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_variable),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_variable),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__contextual_keyword),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__contextual_keyword),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_dot_index_expression),
		Fproduction_id: uint16(29),
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
		Fcount: uint8(1),
	}})),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_dot_index_expression),
		Fproduction_id: uint16(29),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_bracket_index_expression),
		Fproduction_id: uint16(29),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_bracket_index_expression),
		Fproduction_id: uint16(29),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	58: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_table_constructor),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_table_constructor),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	62: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__prefix_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	68: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__prefix_expression),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_arguments),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_arguments),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(14),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_function_call),
		Fproduction_id: uint16(14),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parenthesized_expression),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parenthesized_expression),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_arguments),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_arguments),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_arguments),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	90: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_arguments),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_arguments),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_string),
		Fproduction_id: uint16(9),
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
		Fcount: uint8(1),
	}})),
	96: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_string),
		Fproduction_id: uint16(9),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__quote_string),
		Fproduction_id: uint16(18),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__quote_string),
		Fproduction_id: uint16(18),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_table_constructor),
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
		Fsymbol:      uint16(sym_table_constructor),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__block_string),
		Fproduction_id: uint16(16),
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
		Fcount: uint8(1),
	}})),
	108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__block_string),
		Fproduction_id: uint16(16),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__quote_string),
		Fproduction_id: uint16(35),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__quote_string),
		Fproduction_id: uint16(35),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(145)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__block),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fcount: uint8(2),
	}})),
	126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
	})))),
	127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(148)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(272)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(160)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(262)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	143: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(47)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(77)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(48)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(73)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(201)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(190)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(181)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount: uint8(2),
	}})),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(167)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(75)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_elseif_statement),
		Fproduction_id: uint16(37),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(36),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(36),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		Fcount: uint8(1),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__function_body),
		Fproduction_id: uint16(34),
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
		Fcount: uint8(1),
	}})),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__function_body),
		Fproduction_id: uint16(34),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__function_body),
		Fproduction_id: uint16(49),
	})))),
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
		Fcount: uint8(1),
	}})),
	208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__function_body),
		Fproduction_id: uint16(49),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_function_definition),
		Fproduction_id: uint16(17),
	})))),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_function_definition),
		Fproduction_id: uint16(17),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(20),
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
		Fcount: uint8(1),
	}})),
	216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_unary_expression),
		Fproduction_id: uint16(20),
	})))),
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
		Fcount: uint8(1),
	}})),
	218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(146)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_else_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(144)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(165)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_chunk),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_chunk),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__variable_assignment_explist),
		Fproduction_id: uint16(19),
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
		Fcount: uint8(1),
	}})),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__variable_assignment_explist),
		Fproduction_id: uint16(19),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(72)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_explist_repeat1),
		Fproduction_id: uint16(59),
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
		Fcount: uint8(1),
	}})),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_explist_repeat1),
		Fproduction_id: uint16(59),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	246: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_repeat_statement),
		Fproduction_id: uint16(22),
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
		Fcount: uint8(1),
	}})),
	248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_repeat_statement),
		Fproduction_id: uint16(22),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_repeat_statement),
		Fproduction_id: uint16(38),
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
		Fcount: uint8(1),
	}})),
	252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_repeat_statement),
		Fproduction_id: uint16(38),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_return_statement),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(219)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(202)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(203)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(275)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_return_statement),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__field_list),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__field_list),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fcount: uint8(1),
	}})),
	294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__expression_list_repeat1),
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
		Fcount: uint8(1),
	}})),
	298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__expression_list_repeat1),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	302: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression_list),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression_list),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_statement),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_statement),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(26),
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
		Fcount: uint8(1),
	}})),
	316: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(26),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(238)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(259)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(7),
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
		Fcount: uint8(1),
	}})),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(7),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(19),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(50),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__att_name_list_repeat1),
		Fproduction_id: uint16(31),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__att_name_list_repeat1),
		Fproduction_id: uint16(31),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	336: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(67),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(58),
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
		Fcount: uint8(1),
	}})),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(58),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(24),
	})))),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(24),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(25),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(25),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(42),
	})))),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(42),
	})))),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(46),
	})))),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	356: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(46),
	})))),
	357: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__attrib),
	})))),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	360: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__attrib),
	})))),
	361: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	363: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_for_numeric_clause),
		Fproduction_id: uint16(65),
	})))),
	365: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(45),
	})))),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__att_name_list),
		Fproduction_id: uint16(45),
	})))),
	369: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__att_name_list_repeat1),
		Fproduction_id: uint16(43),
	})))),
	371: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__att_name_list_repeat1),
		Fproduction_id: uint16(43),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__att_name_list_repeat1),
		Fproduction_id: uint16(43),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(238)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(267)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	381: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_explist_repeat1),
		Fproduction_id: uint16(60),
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
		Fcount: uint8(1),
	}})),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_explist_repeat1),
		Fproduction_id: uint16(60),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	387: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_explist_repeat1),
		Fproduction_id: uint16(60),
	})))),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(aux_sym__att_name_list_repeat1),
		Fproduction_id: uint16(57),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(aux_sym__att_name_list_repeat1),
		Fproduction_id: uint16(57),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	400: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_for_numeric_clause),
		Fproduction_id: uint16(69),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__variable_assignment_explist),
		Fproduction_id: uint16(47),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__variable_assignment_explist),
		Fproduction_id: uint16(47),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	408: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__global_variable_declaration),
		Fproduction_id: uint16(13),
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
		Fcount: uint8(1),
	}})),
	410: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__global_variable_declaration),
		Fproduction_id: uint16(13),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(13),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(13),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(4),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_statement),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount: uint8(1),
	}})),
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_statement),
		Fproduction_id: uint16(1),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	426: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(23),
	})))),
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
		Fcount: uint8(1),
	}})),
	428: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_function_declaration),
		Fproduction_id: uint16(23),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym_declaration),
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
		Fcount: uint8(1),
	}})),
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_declaration),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:        uint16(sym__global_implicit_variable_declaration),
		Fproduction_id: uint16(27),
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
		Fsymbol:        uint16(sym__global_implicit_variable_declaration),
		Fproduction_id: uint16(27),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	438: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_statement),
		Fproduction_id: uint16(28),
	})))),
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
		Fcount: uint8(1),
	}})),
	440: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_assignment_statement),
		Fproduction_id: uint16(28),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(2),
	})))),
	443: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(2),
	})))),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(12),
	})))),
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
		Fcount: uint8(1),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_variable_declaration),
		Fproduction_id: uint16(12),
	})))),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(3),
	})))),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(3),
	})))),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__global_implicit_variable_declaration),
	})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__global_implicit_variable_declaration),
	})))),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__global_variable_declaration),
		Fproduction_id: uint16(12),
	})))),
	459: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	460: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__global_variable_declaration),
		Fproduction_id: uint16(12),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	462: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount: uint8(1),
	}})),
	464: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(5),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount: uint8(1),
	}})),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_declaration),
		Fproduction_id: uint16(6),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(37),
	})))),
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
		Fcount: uint8(1),
	}})),
	472: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(37),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	474: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(37),
	})))),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	476: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(37),
	})))),
	477: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	478: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_do_statement),
	})))),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_do_statement),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	482: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__local_function_declaration),
		Fproduction_id: uint16(41),
	})))),
	483: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	484: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__local_function_declaration),
		Fproduction_id: uint16(41),
	})))),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_empty_statement),
	})))),
	487: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	488: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_empty_statement),
	})))),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__variable_assignment),
		Fproduction_id: uint16(44),
	})))),
	491: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	492: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__variable_assignment),
		Fproduction_id: uint16(44),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	494: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__global_function_declaration),
		Fproduction_id: uint16(41),
	})))),
	495: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	496: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__global_function_declaration),
		Fproduction_id: uint16(41),
	})))),
	497: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	498: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(51),
	})))),
	499: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	500: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_while_statement),
		Fproduction_id: uint16(51),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	502: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(52),
	})))),
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
		Fcount: uint8(1),
	}})),
	504: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(52),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	506: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(53),
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
		Fcount: uint8(1),
	}})),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(53),
	})))),
	509: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	510: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(54),
	})))),
	511: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	512: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(54),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	514: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_for_statement),
		Fproduction_id: uint16(56),
	})))),
	515: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	516: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_for_statement),
		Fproduction_id: uint16(56),
	})))),
	517: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	518: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(62),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(62),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(63),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(63),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(64),
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
		Fcount: uint8(1),
	}})),
	528: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(64),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	530: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(68),
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
		Fcount: uint8(1),
	}})),
	532: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_if_statement),
		Fproduction_id: uint16(68),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fcount: uint8(1),
	}})),
	536: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_chunk_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	538: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_label_statement),
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
		Fcount: uint8(1),
	}})),
	540: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_label_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	542: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_do_statement),
		Fproduction_id: uint16(21),
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(sym_do_statement),
		Fproduction_id: uint16(21),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	546: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_goto_statement),
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
		Fcount: uint8(1),
	}})),
	548: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_goto_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	550: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_for_statement),
		Fproduction_id: uint16(40),
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
		Fcount: uint8(1),
	}})),
	552: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_for_statement),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__field_sep),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__field_sep),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
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
		Fcount: uint8(1),
	}})),
	560: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(243)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	562: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_parameters),
		Fproduction_id: uint16(48),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	566: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_parameters),
		Fproduction_id: uint16(48),
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
		Fcount: uint8(1),
	}})),
	568: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_parameters),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	570: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_parameters),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(256)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(258)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	576: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(204)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	578: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	580: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	582: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	584: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	586: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(278)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	588: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__variable_assignment_varlist),
		Fproduction_id: uint16(7),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	590: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	592: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__expression_list_repeat1),
	})))),
	593: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(71)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	595: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_varlist_repeat1),
		Fproduction_id: uint16(31),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	597: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__expression_list),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__expression_list),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	605: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(154)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(251)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	611: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_return_statement),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	613: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(189)),
	}})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_return_statement),
	})))),
	616: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__field_list_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(166)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	620: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__field_list_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	624: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_statement_repeat1),
		Fproduction_id: uint16(55),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	626: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_statement_repeat1),
		Fproduction_id: uint16(55),
	})))),
	627: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(80)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	628: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	629: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_if_statement_repeat1),
		Fproduction_id: uint16(55),
	})))),
	630: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_method_index_expression),
		Fproduction_id: uint16(30),
	})))),
	632: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__field_list),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_return_statement),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_return_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	639: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(225)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(209)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	643: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(169)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	647: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__block),
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
		Fcount: uint8(1),
	}})),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__block),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(208)),
	}})))),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(264)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	657: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__block),
	})))),
	658: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	659: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__name_list_repeat1),
		Fproduction_id: uint16(32),
	})))),
	660: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(268)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	662: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__name_list_repeat1),
		Fproduction_id: uint16(32),
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
		Fcount: uint8(1),
	}})),
	664: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	665: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	666: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
	}})))),
	667: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	668: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__doublequote_string_content),
	})))),
	669: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	670: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__doublequote_string_content),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(210)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__singlequote_string_content),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	675: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__singlequote_string_content),
	})))),
	676: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(208)),
		Frepetition: libc.BoolUint8(1 != 0),
	}})))),
	677: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	678: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	679: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	680: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(268)),
	}})))),
	681: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	682: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__name_list),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	684: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
	}})))),
	685: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	686: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	687: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	688: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	689: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	690: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_varlist_repeat1),
		Fproduction_id: uint16(32),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	692: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__variable_assignment_varlist_repeat1),
		Fproduction_id: uint16(32),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(171)),
		Frepetition: libc.BoolUint8(1 != 0),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__singlequote_string_content),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	697: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__singlequote_string_content),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	701: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__parameter_list),
		Fproduction_id: uint16(7),
	})))),
	702: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	703: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__doublequote_string_content),
	})))),
	704: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym__doublequote_string_content),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	707: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	708: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_if_statement_repeat1),
		Fproduction_id: uint16(39),
	})))),
	710: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	711: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_if_statement_repeat1),
		Fproduction_id: uint16(39),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(221)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	717: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__parameter_list),
		Fproduction_id: uint16(15),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(191)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	721: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__variable_assignment_varlist),
		Fproduction_id: uint16(15),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	723: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__name_list),
		Fproduction_id: uint16(15),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	725: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__name_list_repeat1),
		Fproduction_id: uint16(31),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	727: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_elseif_statement),
		Fproduction_id: uint16(52),
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
		Fcount: uint8(1),
	}})),
	729: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_elseif_statement),
		Fproduction_id: uint16(52),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:      uint16(sym__function_name_prefix_expression),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	733: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__function_name_prefix_expression),
		Fproduction_id: uint16(10),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(235)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(244)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	739: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__function_name),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	741: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__function_name_dot_index_expression),
		Fproduction_id: uint16(29),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	745: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__vararg_parameter),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	749: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fsymbol:        uint16(sym_for_generic_clause),
		Fproduction_id: uint16(33),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(153)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	757: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_else_statement),
		Fproduction_id: uint16(21),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(156)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	765: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	767: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(158)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	769: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	771: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	773: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(223)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	775: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(276)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	777: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_chunk),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	779: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	783: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__parameter_list),
		Fproduction_id: uint16(61),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	785: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__vararg_parameter),
		Fproduction_id: uint16(31),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	787: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__function_name),
		Fproduction_id: uint16(11),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(216)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(253)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	793: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(281)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	795: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	801: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__parameter_list),
		Fproduction_id: uint16(66),
	})))),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	803: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
	}})))),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(237)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	811: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		Freusable: libc.BoolUint8(1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(162)),
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
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	817: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	818: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	819: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(163)),
	}})))),
	820: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	821: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
	}})))),
	822: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	823: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
	}})))),
	824: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	825: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(265)),
	}})))),
	826: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__parameter_list),
		Fproduction_id: uint16(33),
	})))),
	828: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	829: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	830: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	831: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
	}})))),
	832: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	833: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__function_name_method_index_expression),
		Fproduction_id: uint16(30),
	})))),
	834: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	835: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	836: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	837: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(274)),
	}})))),
	838: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	839: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	840: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	841: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_comment),
		Fproduction_id: uint16(8),
	})))),
	842: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(1 != 0),
	}})),
	843: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_comment),
		Fproduction_id: uint16(16),
	})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__block_comment_start = 0
const ts_external_token__block_comment_content = 1
const ts_external_token__block_comment_end = 2
const ts_external_token__block_string_start = 3
const ts_external_token__block_string_content = 4
const ts_external_token__block_string_end = 5

var ts_external_scanner_symbol_map = [6]TSSymbol{
	0: uint16(sym__block_comment_start),
	1: uint16(sym__block_comment_content),
	2: uint16(sym__block_comment_end),
	3: uint16(sym__block_string_start),
	4: uint16(sym__block_string_content),
	5: uint16(sym__block_string_end),
}

var ts_external_scanner_states = [8][6]uint8{
	1: {
		0: libc.BoolUint8(1 != 0),
		1: libc.BoolUint8(1 != 0),
		2: libc.BoolUint8(1 != 0),
		3: libc.BoolUint8(1 != 0),
		4: libc.BoolUint8(1 != 0),
		5: libc.BoolUint8(1 != 0),
	},
	2: {
		0: libc.BoolUint8(1 != 0),
	},
	3: {
		0: libc.BoolUint8(1 != 0),
		3: libc.BoolUint8(1 != 0),
	},
	4: {
		0: libc.BoolUint8(1 != 0),
		5: libc.BoolUint8(1 != 0),
	},
	5: {
		0: libc.BoolUint8(1 != 0),
		1: libc.BoolUint8(1 != 0),
	},
	6: {
		0: libc.BoolUint8(1 != 0),
		2: libc.BoolUint8(1 != 0),
	},
	7: {
		0: libc.BoolUint8(1 != 0),
		4: libc.BoolUint8(1 != 0),
	},
}

func tree_sitter_lua(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(15),
	Fsymbol_count:              uint32(143),
	Ftoken_count:               uint32(73),
	Fexternal_token_count:      uint32(6),
	Fstate_count:               uint32(282),
	Flarge_state_count:         uint32(2),
	Fproduction_id_count:       uint32(70),
	Ffield_count:               uint32(23),
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
	Fkeyword_capture_token:     uint16(sym_identifier),
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
	Fname:                  __ccgo_ts + 1594,
	Fsupertype_count:       uint32(4),
	Fsupertype_symbols:     uintptr(unsafe.Pointer(&ts_supertype_symbols)),
	Fsupertype_map_slices:  uintptr(unsafe.Pointer(&ts_supertype_map_slices)),
	Fsupertype_map_entries: uintptr(unsafe.Pointer(&ts_supertype_map_entries)),
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(4),
		Fpatch_version: uint8(1),
	},
}

func init() {
	p := unsafe.Pointer(&language)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(ts_lex)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_lua_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_lua_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_lua_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_lua_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_lua_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00identifier\x00hash_bang_line\x00return\x00;\x00=\x00,\x00::\x00break_statement\x00goto\x00do\x00while\x00repeat\x00until\x00if\x00then\x00elseif\x00else\x00for\x00in\x00function\x00local\x00global\x00.\x00:\x00*\x00<\x00>\x00nil\x00false\x00true\x00number\x00\"\x00'\x00_doublequote_string_content_token1\x00_singlequote_string_content_token1\x00escape_sequence\x00vararg_expression\x00(\x00)\x00[\x00]\x00{\x00}\x00or\x00and\x00<=\x00==\x00~=\x00>=\x00|\x00~\x00&\x00<<\x00>>\x00+\x00-\x00/\x00//\x00%\x00..\x00^\x00not\x00#\x00--\x00comment_content\x00[[\x00]]\x00string_content\x00chunk\x00block\x00statement\x00return_statement\x00empty_statement\x00assignment_statement\x00variable_list\x00expression_list\x00label_statement\x00goto_statement\x00do_statement\x00while_statement\x00repeat_statement\x00if_statement\x00elseif_statement\x00else_statement\x00for_statement\x00for_generic_clause\x00for_numeric_clause\x00declaration\x00function_declaration\x00_function_name\x00_function_name_prefix_expression\x00dot_index_expression\x00method_index_expression\x00variable_declaration\x00implicit_variable_declaration\x00attribute\x00expression\x00string\x00_quote_string\x00_doublequote_string_content\x00_singlequote_string_content\x00_block_string\x00function_definition\x00_function_body\x00parameters\x00_parameter_list\x00_vararg_parameter\x00_prefix_expression\x00variable\x00bracket_index_expression\x00function_call\x00arguments\x00parenthesized_expression\x00table_constructor\x00_field_list\x00_field_sep\x00field\x00binary_expression\x00unary_expression\x00comment\x00_contextual_keyword\x00chunk_repeat1\x00_variable_assignment_varlist_repeat1\x00_variable_assignment_explist_repeat1\x00if_statement_repeat1\x00_name_list_repeat1\x00_att_name_list_repeat1\x00_expression_list_repeat1\x00_field_list_repeat1\x00alternative\x00body\x00clause\x00condition\x00consequence\x00content\x00global_declaration\x00left\x00local_declaration\x00method\x00name\x00operand\x00operator\x00right\x00start\x00step\x00table\x00value\x00lua\x00"
