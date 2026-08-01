// Code generated for darwin/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -D__attribute__(...)= -D__extension__= -D_Nonnull= -D_Nullable= -D_Null_unspecified= -DAPI_AVAILABLE(...)= -DAPI_UNAVAILABLE(...)= -DAPI_DEPRECATED(...)= -DAPI_DEPRECATED_WITH_REPLACEMENT(...)= -D__API_AVAILABLE(...)= -D__API_UNAVAILABLE(...)= -D__API_DEPRECATED(...)= -D__API_DEPRECATED_WITH_REPLACEMENT(...)= -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-kconfig/src -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-kconfig -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build darwin && amd64

package grammar_kconfig

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 0
const BADSIG = "SIG_ERR"
const BIG_ENDIAN = "__DARWIN_BIG_ENDIAN"
const BUS_ADRALN = 1
const BUS_ADRERR = 2
const BUS_NOOP = 0
const BUS_OBJERR = 3
const BYTE_ORDER = "__DARWIN_BYTE_ORDER"
const CLD_CONTINUED = 6
const CLD_DUMPED = 3
const CLD_EXITED = 1
const CLD_KILLED = 2
const CLD_NOOP = 0
const CLD_STOPPED = 5
const CLD_TRAPPED = 4
const CPUMON_MAKE_FATAL = 0x1000
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXTERNAL_TOKEN_COUNT = 1
const FIELD_COUNT = 5
const FOOTPRINT_INTERVAL_RESET = 0x1
const FPE_FLTDIV = 1
const FPE_FLTINV = 5
const FPE_FLTOVF = 2
const FPE_FLTRES = 4
const FPE_FLTSUB = 6
const FPE_FLTUND = 3
const FPE_INTDIV = 7
const FPE_INTOVF = 8
const FPE_NOOP = 0
const FP_CHOP = 3
const FP_PREC_24B = 0
const FP_PREC_53B = 2
const FP_PREC_64B = 3
const FP_RND_DOWN = 1
const FP_RND_NEAR = 0
const FP_RND_UP = 2
const FP_STATE_BYTES = 512
const ILL_BADSTK = 8
const ILL_COPROC = 7
const ILL_ILLADR = 5
const ILL_ILLOPC = 1
const ILL_ILLOPN = 4
const ILL_ILLTRP = 2
const ILL_NOOP = 0
const ILL_PRVOPC = 3
const ILL_PRVREG = 6
const INTMAX_MAX = "__INTMAX_MAX__"
const INTPTR_MAX = "__INTPTR_MAX__"
const INT_FAST16_MAX = "__INT_LEAST16_MAX"
const INT_FAST16_MIN = "__INT_LEAST16_MIN"
const INT_FAST32_MAX = "__INT_LEAST32_MAX"
const INT_FAST32_MIN = "__INT_LEAST32_MIN"
const INT_FAST64_MAX = "__INT_LEAST64_MAX"
const INT_FAST64_MIN = "__INT_LEAST64_MIN"
const INT_FAST8_MAX = "__INT_LEAST8_MAX"
const INT_FAST8_MIN = "__INT_LEAST8_MIN"
const INT_LEAST16_MAX = "__INT_LEAST16_MAX"
const INT_LEAST16_MIN = "__INT_LEAST16_MIN"
const INT_LEAST32_MAX = "__INT_LEAST32_MAX"
const INT_LEAST32_MIN = "__INT_LEAST32_MIN"
const INT_LEAST64_MAX = "__INT_LEAST64_MAX"
const INT_LEAST64_MIN = "__INT_LEAST64_MIN"
const INT_LEAST8_MAX = "__INT_LEAST8_MAX"
const INT_LEAST8_MIN = "__INT_LEAST8_MIN"
const IOPOL_APPLICATION = "IOPOL_STANDARD"
const IOPOL_ATIME_UPDATES_DEFAULT = 0
const IOPOL_ATIME_UPDATES_OFF = 1
const IOPOL_DEFAULT = 0
const IOPOL_IMPORTANT = 1
const IOPOL_MATERIALIZE_DATALESS_FILES_DEFAULT = 0
const IOPOL_MATERIALIZE_DATALESS_FILES_OFF = 1
const IOPOL_MATERIALIZE_DATALESS_FILES_ON = 2
const IOPOL_NORMAL = "IOPOL_IMPORTANT"
const IOPOL_PASSIVE = 2
const IOPOL_SCOPE_DARWIN_BG = 2
const IOPOL_SCOPE_PROCESS = 0
const IOPOL_SCOPE_THREAD = 1
const IOPOL_STANDARD = 5
const IOPOL_THROTTLE = 3
const IOPOL_TYPE_DISK = 0
const IOPOL_TYPE_VFS_ALLOW_LOW_SPACE_WRITES = 9
const IOPOL_TYPE_VFS_ATIME_UPDATES = 2
const IOPOL_TYPE_VFS_DISALLOW_RW_FOR_O_EVTONLY = 10
const IOPOL_TYPE_VFS_IGNORE_CONTENT_PROTECTION = 6
const IOPOL_TYPE_VFS_IGNORE_PERMISSIONS = 7
const IOPOL_TYPE_VFS_MATERIALIZE_DATALESS_FILES = 3
const IOPOL_TYPE_VFS_SKIP_MTIME_UPDATE = 8
const IOPOL_TYPE_VFS_STATFS_NO_DATA_VOLUME = 4
const IOPOL_TYPE_VFS_TRIGGER_RESOLVE = 5
const IOPOL_UTILITY = 4
const IOPOL_VFS_ALLOW_LOW_SPACE_WRITES_OFF = 0
const IOPOL_VFS_ALLOW_LOW_SPACE_WRITES_ON = 1
const IOPOL_VFS_CONTENT_PROTECTION_DEFAULT = 0
const IOPOL_VFS_CONTENT_PROTECTION_IGNORE = 1
const IOPOL_VFS_DISALLOW_RW_FOR_O_EVTONLY_DEFAULT = 0
const IOPOL_VFS_DISALLOW_RW_FOR_O_EVTONLY_ON = 1
const IOPOL_VFS_IGNORE_PERMISSIONS_OFF = 0
const IOPOL_VFS_IGNORE_PERMISSIONS_ON = 1
const IOPOL_VFS_NOCACHE_WRITE_FS_BLKSIZE_DEFAULT = 0
const IOPOL_VFS_NOCACHE_WRITE_FS_BLKSIZE_ON = 1
const IOPOL_VFS_SKIP_MTIME_UPDATE_IGNORE = 2
const IOPOL_VFS_SKIP_MTIME_UPDATE_OFF = 0
const IOPOL_VFS_SKIP_MTIME_UPDATE_ON = 1
const IOPOL_VFS_STATFS_FORCE_NO_DATA_VOLUME = 1
const IOPOL_VFS_STATFS_NO_DATA_VOLUME_DEFAULT = 0
const IOPOL_VFS_TRIGGER_RESOLVE_DEFAULT = 0
const IOPOL_VFS_TRIGGER_RESOLVE_OFF = 1
const LANGUAGE_VERSION = 14
const LARGE_STATE_COUNT = 10
const LITTLE_ENDIAN = "__DARWIN_LITTLE_ENDIAN"
const MAC_OS_VERSION_11_0 = "__MAC_11_0"
const MAC_OS_VERSION_11_1 = "__MAC_11_1"
const MAC_OS_VERSION_11_3 = "__MAC_11_3"
const MAC_OS_VERSION_11_4 = "__MAC_11_4"
const MAC_OS_VERSION_11_5 = "__MAC_11_5"
const MAC_OS_VERSION_11_6 = "__MAC_11_6"
const MAC_OS_VERSION_12_0 = "__MAC_12_0"
const MAC_OS_VERSION_12_1 = "__MAC_12_1"
const MAC_OS_VERSION_12_2 = "__MAC_12_2"
const MAC_OS_VERSION_12_3 = "__MAC_12_3"
const MAC_OS_VERSION_12_4 = "__MAC_12_4"
const MAC_OS_VERSION_12_5 = "__MAC_12_5"
const MAC_OS_VERSION_12_6 = "__MAC_12_6"
const MAC_OS_VERSION_12_7 = "__MAC_12_7"
const MAC_OS_VERSION_13_0 = "__MAC_13_0"
const MAC_OS_VERSION_13_1 = "__MAC_13_1"
const MAC_OS_VERSION_13_2 = "__MAC_13_2"
const MAC_OS_VERSION_13_3 = "__MAC_13_3"
const MAC_OS_VERSION_13_4 = "__MAC_13_4"
const MAC_OS_VERSION_13_5 = "__MAC_13_5"
const MAC_OS_VERSION_13_6 = "__MAC_13_6"
const MAC_OS_VERSION_13_7 = "__MAC_13_7"
const MAC_OS_VERSION_14_0 = "__MAC_14_0"
const MAC_OS_VERSION_14_1 = "__MAC_14_1"
const MAC_OS_VERSION_14_2 = "__MAC_14_2"
const MAC_OS_VERSION_14_3 = "__MAC_14_3"
const MAC_OS_VERSION_14_4 = "__MAC_14_4"
const MAC_OS_VERSION_14_5 = "__MAC_14_5"
const MAC_OS_VERSION_14_6 = "__MAC_14_6"
const MAC_OS_VERSION_14_7 = "__MAC_14_7"
const MAC_OS_VERSION_15_0 = "__MAC_15_0"
const MAC_OS_VERSION_15_1 = "__MAC_15_1"
const MAC_OS_VERSION_15_2 = "__MAC_15_2"
const MAC_OS_VERSION_15_3 = "__MAC_15_3"
const MAC_OS_VERSION_15_4 = "__MAC_15_4"
const MAC_OS_VERSION_15_5 = "__MAC_15_5"
const MAC_OS_X_VERSION_10_0 = "__MAC_10_0"
const MAC_OS_X_VERSION_10_1 = "__MAC_10_1"
const MAC_OS_X_VERSION_10_10 = "__MAC_10_10"
const MAC_OS_X_VERSION_10_10_2 = "__MAC_10_10_2"
const MAC_OS_X_VERSION_10_10_3 = "__MAC_10_10_3"
const MAC_OS_X_VERSION_10_11 = "__MAC_10_11"
const MAC_OS_X_VERSION_10_11_2 = "__MAC_10_11_2"
const MAC_OS_X_VERSION_10_11_3 = "__MAC_10_11_3"
const MAC_OS_X_VERSION_10_11_4 = "__MAC_10_11_4"
const MAC_OS_X_VERSION_10_12 = "__MAC_10_12"
const MAC_OS_X_VERSION_10_12_1 = "__MAC_10_12_1"
const MAC_OS_X_VERSION_10_12_2 = "__MAC_10_12_2"
const MAC_OS_X_VERSION_10_12_4 = "__MAC_10_12_4"
const MAC_OS_X_VERSION_10_13 = "__MAC_10_13"
const MAC_OS_X_VERSION_10_13_1 = "__MAC_10_13_1"
const MAC_OS_X_VERSION_10_13_2 = "__MAC_10_13_2"
const MAC_OS_X_VERSION_10_13_4 = "__MAC_10_13_4"
const MAC_OS_X_VERSION_10_14 = "__MAC_10_14"
const MAC_OS_X_VERSION_10_14_1 = "__MAC_10_14_1"
const MAC_OS_X_VERSION_10_14_4 = "__MAC_10_14_4"
const MAC_OS_X_VERSION_10_14_5 = "__MAC_10_14_5"
const MAC_OS_X_VERSION_10_14_6 = "__MAC_10_14_6"
const MAC_OS_X_VERSION_10_15 = "__MAC_10_15"
const MAC_OS_X_VERSION_10_15_1 = "__MAC_10_15_1"
const MAC_OS_X_VERSION_10_15_4 = "__MAC_10_15_4"
const MAC_OS_X_VERSION_10_16 = "__MAC_10_16"
const MAC_OS_X_VERSION_10_2 = "__MAC_10_2"
const MAC_OS_X_VERSION_10_3 = "__MAC_10_3"
const MAC_OS_X_VERSION_10_4 = "__MAC_10_4"
const MAC_OS_X_VERSION_10_5 = "__MAC_10_5"
const MAC_OS_X_VERSION_10_6 = "__MAC_10_6"
const MAC_OS_X_VERSION_10_7 = "__MAC_10_7"
const MAC_OS_X_VERSION_10_8 = "__MAC_10_8"
const MAC_OS_X_VERSION_10_9 = "__MAC_10_9"
const MAX_ALIAS_SEQUENCE_LENGTH = 5
const MB_CUR_MAX = "__mb_cur_max"
const MINSIGSTKSZ = 32768
const NSIG = "__DARWIN_NSIG"
const NULL = "__DARWIN_NULL"
const PDP_ENDIAN = "__DARWIN_PDP_ENDIAN"
const POLL_ERR = 4
const POLL_HUP = 6
const POLL_IN = 1
const POLL_MSG = 3
const POLL_OUT = 2
const POLL_PRI = 5
const PRIO_DARWIN_BG = 0x1000
const PRIO_DARWIN_NONUI = 0x1001
const PRIO_DARWIN_PROCESS = 4
const PRIO_DARWIN_THREAD = 3
const PRIO_MAX = 20
const PRIO_PGRP = 1
const PRIO_PROCESS = 0
const PRIO_USER = 2
const PRODUCTION_ID_COUNT = 9
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const RAND_MAX = 0x7fffffff
const RLIMIT_AS = 5
const RLIMIT_CORE = 4
const RLIMIT_CPU = 0
const RLIMIT_CPU_USAGE_MONITOR = 0x2
const RLIMIT_DATA = 2
const RLIMIT_FOOTPRINT_INTERVAL = 0x4
const RLIMIT_FSIZE = 1
const RLIMIT_MEMLOCK = 6
const RLIMIT_NOFILE = 8
const RLIMIT_NPROC = 7
const RLIMIT_RSS = "RLIMIT_AS"
const RLIMIT_STACK = 3
const RLIMIT_THREAD_CPULIMITS = 0x3
const RLIMIT_WAKEUPS_MONITOR = 0x1
const RLIM_NLIMITS = 9
const RLIM_SAVED_CUR = "RLIM_INFINITY"
const RLIM_SAVED_MAX = "RLIM_INFINITY"
const RUSAGE_INFO_CURRENT = "RUSAGE_INFO_V6"
const RUSAGE_INFO_V0 = 0
const RUSAGE_INFO_V1 = 1
const RUSAGE_INFO_V2 = 2
const RUSAGE_INFO_V3 = 3
const RUSAGE_INFO_V4 = 4
const RUSAGE_INFO_V5 = 5
const RUSAGE_INFO_V6 = 6
const RUSAGE_SELF = 0
const RU_PROC_RUNS_RESLIDE = 0x00000001
const SA_64REGSET = 0x0200
const SA_NOCLDSTOP = 0x0008
const SA_NOCLDWAIT = 0x0020
const SA_NODEFER = 0x0010
const SA_ONSTACK = 0x0001
const SA_RESETHAND = 0x0004
const SA_RESTART = 0x0002
const SA_SIGINFO = 0x0040
const SA_USERTRAMP = 0x0100
const SEGV_ACCERR = 2
const SEGV_MAPERR = 1
const SEGV_NOOP = 0
const SIGABRT = 6
const SIGALRM = 14
const SIGBUS = 10
const SIGCHLD = 20
const SIGCONT = 19
const SIGEMT = 7
const SIGEV_NONE = 0
const SIGEV_SIGNAL = 1
const SIGEV_THREAD = 3
const SIGFPE = 8
const SIGHUP = 1
const SIGILL = 4
const SIGINFO = 29
const SIGINT = 2
const SIGIO = 23
const SIGIOT = "SIGABRT"
const SIGKILL = 9
const SIGPIPE = 13
const SIGPROF = 27
const SIGQUIT = 3
const SIGSEGV = 11
const SIGSTKSZ = 131072
const SIGSTOP = 17
const SIGSYS = 12
const SIGTERM = 15
const SIGTRAP = 5
const SIGTSTP = 18
const SIGTTIN = 21
const SIGTTOU = 22
const SIGURG = 16
const SIGUSR1 = 30
const SIGUSR2 = 31
const SIGVTALRM = 26
const SIGWINCH = 28
const SIGXCPU = 24
const SIGXFSZ = 25
const SIG_BLOCK = 1
const SIG_SETMASK = 3
const SIG_UNBLOCK = 2
const SIZE_MAX = "__SIZE_MAX__"
const SI_ASYNCIO = 0x10004
const SI_MESGQ = 0x10005
const SI_QUEUE = 0x10002
const SI_TIMER = 0x10003
const SI_USER = 0x10001
const SS_DISABLE = 0x0004
const SS_ONSTACK = 0x0001
const STATE_COUNT = 383
const SV_INTERRUPT = "SA_RESTART"
const SV_NOCLDSTOP = "SA_NOCLDSTOP"
const SV_NODEFER = "SA_NODEFER"
const SV_ONSTACK = "SA_ONSTACK"
const SV_RESETHAND = "SA_RESETHAND"
const SV_SIGINFO = "SA_SIGINFO"
const SYMBOL_COUNT = 103
const TARGET_IPHONE_SIMULATOR = 0
const TARGET_OS_ARROW = 0
const TARGET_OS_BRIDGE = 0
const TARGET_OS_DRIVERKIT = 0
const TARGET_OS_EMBEDDED = 0
const TARGET_OS_IOS = 0
const TARGET_OS_IOSMAC = 0
const TARGET_OS_IPHONE = 0
const TARGET_OS_LINUX = 0
const TARGET_OS_MAC = 1
const TARGET_OS_MACCATALYST = 0
const TARGET_OS_NANO = 0
const TARGET_OS_OSX = 1
const TARGET_OS_SIMULATOR = 0
const TARGET_OS_TV = 0
const TARGET_OS_UIKITFORMAC = 0
const TARGET_OS_UNIX = 0
const TARGET_OS_VISION = 0
const TARGET_OS_WATCH = 0
const TARGET_OS_WIN32 = 0
const TARGET_OS_WINDOWS = 0
const TARGET_OS_XR = 0
const TOKEN_COUNT = 64
const TRAP_BRKPT = 1
const TRAP_TRACE = 2
const TREE_SITTER_SERIALIZATION_BUFFER_SIZE = 1024
const UINTMAX_MAX = "__UINTMAX_MAX__"
const UINTPTR_MAX = "__UINTPTR_MAX__"
const UINT_FAST16_MAX = "__UINT_LEAST16_MAX"
const UINT_FAST32_MAX = "__UINT_LEAST32_MAX"
const UINT_FAST64_MAX = "__UINT_LEAST64_MAX"
const UINT_FAST8_MAX = "__UINT_LEAST8_MAX"
const UINT_LEAST16_MAX = "__UINT_LEAST16_MAX"
const UINT_LEAST32_MAX = "__UINT_LEAST32_MAX"
const UINT_LEAST64_MAX = "__UINT_LEAST64_MAX"
const UINT_LEAST8_MAX = "__UINT_LEAST8_MAX"
const WAIT_MYPGRP = 0
const WAKEMON_DISABLE = 0x02
const WAKEMON_ENABLE = 0x01
const WAKEMON_GET_PARAMS = 0x04
const WAKEMON_MAKE_FATAL = 0x10
const WAKEMON_SET_DEFAULTS = 0x08
const WCHAR_MAX = "__WCHAR_MAX__"
const WCONTINUED = 0x00000010
const WCOREFLAG = 0200
const WEOF = "__DARWIN_WEOF"
const WEXITED = 0x00000004
const WNOHANG = 0x00000001
const WNOWAIT = 0x00000020
const WSTOPPED = 0x00000008
const WUNTRACED = 0x00000002
const _CTYPE_A = 256
const _CTYPE_B = 131072
const _CTYPE_C = 512
const _CTYPE_D = 1024
const _CTYPE_G = 2048
const _CTYPE_I = 524288
const _CTYPE_L = 4096
const _CTYPE_P = 8192
const _CTYPE_Q = 2097152
const _CTYPE_R = 262144
const _CTYPE_S = 16384
const _CTYPE_SW0 = 0x20000000
const _CTYPE_SW1 = 0x40000000
const _CTYPE_SW2 = 0x80000000
const _CTYPE_SW3 = 0xc0000000
const _CTYPE_SWM = 3758096384
const _CTYPE_SWS = 30
const _CTYPE_T = 1048576
const _CTYPE_U = 32768
const _CTYPE_X = 65536
const _DARWIN_FEATURE_64_BIT_INODE = 1
const _DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
const _DARWIN_FEATURE_UNIX_CONFORMANCE = 3
const _FORTIFY_SOURCE = 2
const _I386_SIGNAL_H_ = 1
const _LIBC_COUNT__MB_LEN_MAX = "_LIBC_UNSAFE_INDEXABLE"
const _LIBC_COUNT__PATH_MAX = "_LIBC_UNSAFE_INDEXABLE"
const _LP64 = 1
const _QUAD_HIGHWORD = 1
const _QUAD_LOWWORD = 0
const _RLIMIT_POSIX_FLAG = 0x1000
const _RUNE_MAGIC_A = "RuneMagA"
const _STRUCT_MCONTEXT = "_STRUCT_MCONTEXT64"
const _WSTOPPED = 0177
const _X86_INSTRUCTION_STATE_CACHELINE_SIZE = 64
const __API_TO_BE_DEPRECATED = 100000
const __API_TO_BE_DEPRECATED_DRIVERKIT = 100000
const __API_TO_BE_DEPRECATED_IOS = 100000
const __API_TO_BE_DEPRECATED_IOSAPPLICATIONEXTENSION = 100000
const __API_TO_BE_DEPRECATED_KERNELKIT = 100000
const __API_TO_BE_DEPRECATED_MACCATALYST = 100000
const __API_TO_BE_DEPRECATED_MACCATALYSTAPPLICATIONEXTENSION = 100000
const __API_TO_BE_DEPRECATED_MACOS = 100000
const __API_TO_BE_DEPRECATED_MACOSAPPLICATIONEXTENSION = 100000
const __API_TO_BE_DEPRECATED_TVOS = 100000
const __API_TO_BE_DEPRECATED_TVOSAPPLICATIONEXTENSION = 100000
const __API_TO_BE_DEPRECATED_VISIONOS = 100000
const __API_TO_BE_DEPRECATED_VISIONOSAPPLICATIONEXTENSION = 100000
const __API_TO_BE_DEPRECATED_WATCHOS = 100000
const __API_TO_BE_DEPRECATED_WATCHOSAPPLICATIONEXTENSION = 100000
const __APPLE_CC__ = 6000
const __APPLE__ = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __AVAILABILITY_FILE = "AvailabilityVersions.h"
const __AVAILABILITY_VERSIONS_VERSION_HASH = 93585900
const __AVAILABILITY_VERSIONS_VERSION_STRING = "Local"
const __BIGGEST_ALIGNMENT__ = 16
const __BITINT_MAXWIDTH__ = 8388608
const __BLOCKS__ = 1
const __BOOL_WIDTH__ = 8
const __BRIDGEOS_2_0 = 20000
const __BRIDGEOS_3_0 = 30000
const __BRIDGEOS_3_1 = 30100
const __BRIDGEOS_3_4 = 30400
const __BRIDGEOS_4_0 = 40000
const __BRIDGEOS_4_1 = 40100
const __BRIDGEOS_5_0 = 50000
const __BRIDGEOS_5_1 = 50100
const __BRIDGEOS_5_3 = 50300
const __BRIDGEOS_6_0 = 60000
const __BRIDGEOS_6_2 = 60200
const __BRIDGEOS_6_4 = 60400
const __BRIDGEOS_6_5 = 60500
const __BRIDGEOS_6_6 = 60600
const __BRIDGEOS_7_0 = 70000
const __BRIDGEOS_7_1 = 70100
const __BRIDGEOS_7_2 = 70200
const __BRIDGEOS_7_3 = 70300
const __BRIDGEOS_7_4 = 70400
const __BRIDGEOS_7_6 = 70600
const __BRIDGEOS_8_0 = 80000
const __BRIDGEOS_8_1 = 80100
const __BRIDGEOS_8_2 = 80200
const __BRIDGEOS_8_3 = 80300
const __BRIDGEOS_8_4 = 80400
const __BRIDGEOS_8_5 = 80500
const __BRIDGEOS_8_6 = 80600
const __BRIDGEOS_9_0 = 90000
const __BRIDGEOS_9_1 = 90100
const __BRIDGEOS_9_2 = 90200
const __BRIDGEOS_9_3 = 90300
const __BRIDGEOS_9_4 = 90400
const __BRIDGEOS_9_5 = 90500
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CLANG_ATOMIC_BOOL_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR_LOCK_FREE = 2
const __CLANG_ATOMIC_INT_LOCK_FREE = 2
const __CLANG_ATOMIC_LLONG_LOCK_FREE = 2
const __CLANG_ATOMIC_LONG_LOCK_FREE = 2
const __CLANG_ATOMIC_POINTER_LOCK_FREE = 2
const __CLANG_ATOMIC_SHORT_LOCK_FREE = 2
const __CLANG_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __CONSTANT_CFSTRINGS__ = 1
const __DARWIN_64_BIT_INO_T = 1
const __DARWIN_BIG_ENDIAN = 4321
const __DARWIN_BYTE_ORDER = "__DARWIN_LITTLE_ENDIAN"
const __DARWIN_CTYPE_TOP_inline = "__header_inline"
const __DARWIN_CTYPE_inline = "__header_inline"
const __DARWIN_C_ANSI = 010000
const __DARWIN_C_FULL = 900000
const __DARWIN_C_LEVEL = "__DARWIN_C_FULL"
const __DARWIN_LITTLE_ENDIAN = 1234
const __DARWIN_NON_CANCELABLE = 0
const __DARWIN_NO_LONG_LONG = 0
const __DARWIN_NSIG = 32
const __DARWIN_ONLY_64_BIT_INO_T = 0
const __DARWIN_ONLY_UNIX_CONFORMANCE = 1
const __DARWIN_ONLY_VERS_1050 = 0
const __DARWIN_PDP_ENDIAN = 3412
const __DARWIN_SUF_1050 = "$1050"
const __DARWIN_SUF_64_BIT_INO_T = "$INODE64"
const __DARWIN_SUF_EXTSN = "$DARWIN_EXTSN"
const __DARWIN_UNIX03 = 1
const __DARWIN_VERS_1050 = 1
const __DARWIN_WCHAR_MAX = "__WCHAR_MAX__"
const __DARWIN_WCTYPE_TOP_inline = "__header_inline"
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DENORM_MIN__ = 4.9406564584124654e-324
const __DBL_DIG__ = 15
const __DBL_EPSILON__ = 2.2204460492503131e-16
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DBL_MAX__ = 1.7976931348623157e+308
const __DBL_MIN__ = 2.2250738585072014e-308
const __DBL_NORM_MAX__ = 1.7976931348623157e+308
const __DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
const __DRIVERKIT_19_0 = 190000
const __DRIVERKIT_20_0 = 200000
const __DRIVERKIT_21_0 = 210000
const __DRIVERKIT_22_0 = 220000
const __DRIVERKIT_22_4 = 220400
const __DRIVERKIT_22_5 = 220500
const __DRIVERKIT_22_6 = 220600
const __DRIVERKIT_23_0 = 230000
const __DRIVERKIT_23_1 = 230100
const __DRIVERKIT_23_2 = 230200
const __DRIVERKIT_23_3 = 230300
const __DRIVERKIT_23_4 = 230400
const __DRIVERKIT_23_5 = 230500
const __DRIVERKIT_23_6 = 230600
const __DRIVERKIT_24_0 = 240000
const __DRIVERKIT_24_1 = 240100
const __DRIVERKIT_24_2 = 240200
const __DRIVERKIT_24_3 = 240300
const __DRIVERKIT_24_4 = 240400
const __DRIVERKIT_24_5 = 240500
const __DYNAMIC__ = 1
const __ENABLE_LEGACY_MAC_AVAILABILITY = 1
const __ENVIRONMENT_MAC_OS_X_VERSION_MIN_REQUIRED__ = 150000
const __ENVIRONMENT_OS_VERSION_MIN_REQUIRED__ = 150000
const __FINITE_MATH_ONLY__ = 0
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.9604644775390625e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.765625e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.5504e+4
const __FLT16_MIN__ = 6.103515625e-5
const __FLT16_NORM_MAX__ = 6.5504e+4
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209290e-7
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282347e+38
const __FLT_MIN__ = 1.17549435e-38
const __FLT_NORM_MAX__ = 3.40282347e+38
const __FLT_RADIX__ = 2
const __FPCLASS_NEGINF = 0x0004
const __FPCLASS_NEGNORMAL = 0x0008
const __FPCLASS_NEGSUBNORMAL = 0x0010
const __FPCLASS_NEGZERO = 0x0020
const __FPCLASS_POSINF = 0x0200
const __FPCLASS_POSNORMAL = 0x0100
const __FPCLASS_POSSUBNORMAL = 0x0080
const __FPCLASS_POSZERO = 0x0040
const __FPCLASS_QNAN = 0x0002
const __FPCLASS_SNAN = 0x0001
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
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC__ = 4
const __GXX_ABI_VERSION = 1002
const __INT16_FMTd__ = "hd"
const __INT16_FMTi__ = "hi"
const __INT16_MAX__ = 32767
const __INT16_TYPE__ = "short"
const __INT32_FMTd__ = "d"
const __INT32_FMTi__ = "i"
const __INT32_MAX__ = 2147483647
const __INT32_TYPE__ = "int"
const __INT64_C_SUFFIX__ = "LL"
const __INT64_FMTd__ = "lld"
const __INT64_FMTi__ = "lli"
const __INT64_MAX__ = 9223372036854775807
const __INT8_FMTd__ = "hhd"
const __INT8_FMTi__ = "hhi"
const __INT8_MAX__ = 127
const __INTMAX_C_SUFFIX__ = "L"
const __INTMAX_FMTd__ = "ld"
const __INTMAX_FMTi__ = "li"
const __INTMAX_MAX__ = 9223372036854775807
const __INTMAX_WIDTH__ = 64
const __INTPTR_FMTd__ = "ld"
const __INTPTR_FMTi__ = "li"
const __INTPTR_MAX__ = 9223372036854775807
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_FMTd__ = "hd"
const __INT_FAST16_FMTi__ = "hi"
const __INT_FAST16_MAX__ = 32767
const __INT_FAST16_TYPE__ = "short"
const __INT_FAST16_WIDTH__ = 16
const __INT_FAST32_FMTd__ = "d"
const __INT_FAST32_FMTi__ = "i"
const __INT_FAST32_MAX__ = 2147483647
const __INT_FAST32_TYPE__ = "int"
const __INT_FAST32_WIDTH__ = 32
const __INT_FAST64_FMTd__ = "lld"
const __INT_FAST64_FMTi__ = "lli"
const __INT_FAST64_MAX__ = 9223372036854775807
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_FMTd__ = "hhd"
const __INT_FAST8_FMTi__ = "hhi"
const __INT_FAST8_MAX__ = 127
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_FMTd__ = "hd"
const __INT_LEAST16_FMTi__ = "hi"
const __INT_LEAST16_MAX__ = 32767
const __INT_LEAST16_TYPE__ = "short"
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_FMTd__ = "d"
const __INT_LEAST32_FMTi__ = "i"
const __INT_LEAST32_MAX__ = 2147483647
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_FMTd__ = "lld"
const __INT_LEAST64_FMTi__ = "lli"
const __INT_LEAST64_MAX = "INT64_MAX"
const __INT_LEAST64_MAX__ = 9223372036854775807
const __INT_LEAST64_MIN = "INT64_MIN"
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_FMTd__ = "hhd"
const __INT_LEAST8_FMTi__ = "hhi"
const __INT_LEAST8_MAX__ = 127
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 2147483647
const __INT_WIDTH__ = 32
const __IPHONE_10_0 = 100000
const __IPHONE_10_1 = 100100
const __IPHONE_10_2 = 100200
const __IPHONE_10_3 = 100300
const __IPHONE_11_0 = 110000
const __IPHONE_11_1 = 110100
const __IPHONE_11_2 = 110200
const __IPHONE_11_3 = 110300
const __IPHONE_11_4 = 110400
const __IPHONE_12_0 = 120000
const __IPHONE_12_1 = 120100
const __IPHONE_12_2 = 120200
const __IPHONE_12_3 = 120300
const __IPHONE_12_4 = 120400
const __IPHONE_13_0 = 130000
const __IPHONE_13_1 = 130100
const __IPHONE_13_2 = 130200
const __IPHONE_13_3 = 130300
const __IPHONE_13_4 = 130400
const __IPHONE_13_5 = 130500
const __IPHONE_13_6 = 130600
const __IPHONE_13_7 = 130700
const __IPHONE_14_0 = 140000
const __IPHONE_14_1 = 140100
const __IPHONE_14_2 = 140200
const __IPHONE_14_3 = 140300
const __IPHONE_14_4 = 140400
const __IPHONE_14_5 = 140500
const __IPHONE_14_6 = 140600
const __IPHONE_14_7 = 140700
const __IPHONE_14_8 = 140800
const __IPHONE_15_0 = 150000
const __IPHONE_15_1 = 150100
const __IPHONE_15_2 = 150200
const __IPHONE_15_3 = 150300
const __IPHONE_15_4 = 150400
const __IPHONE_15_5 = 150500
const __IPHONE_15_6 = 150600
const __IPHONE_15_7 = 150700
const __IPHONE_15_8 = 150800
const __IPHONE_16_0 = 160000
const __IPHONE_16_1 = 160100
const __IPHONE_16_2 = 160200
const __IPHONE_16_3 = 160300
const __IPHONE_16_4 = 160400
const __IPHONE_16_5 = 160500
const __IPHONE_16_6 = 160600
const __IPHONE_16_7 = 160700
const __IPHONE_17_0 = 170000
const __IPHONE_17_1 = 170100
const __IPHONE_17_2 = 170200
const __IPHONE_17_3 = 170300
const __IPHONE_17_4 = 170400
const __IPHONE_17_5 = 170500
const __IPHONE_17_6 = 170600
const __IPHONE_17_7 = 170700
const __IPHONE_18_0 = 180000
const __IPHONE_18_1 = 180100
const __IPHONE_18_2 = 180200
const __IPHONE_18_3 = 180300
const __IPHONE_18_4 = 180400
const __IPHONE_18_5 = 180500
const __IPHONE_2_0 = 20000
const __IPHONE_2_1 = 20100
const __IPHONE_2_2 = 20200
const __IPHONE_3_0 = 30000
const __IPHONE_3_1 = 30100
const __IPHONE_3_2 = 30200
const __IPHONE_4_0 = 40000
const __IPHONE_4_1 = 40100
const __IPHONE_4_2 = 40200
const __IPHONE_4_3 = 40300
const __IPHONE_5_0 = 50000
const __IPHONE_5_1 = 50100
const __IPHONE_6_0 = 60000
const __IPHONE_6_1 = 60100
const __IPHONE_7_0 = 70000
const __IPHONE_7_1 = 70100
const __IPHONE_8_0 = 80000
const __IPHONE_8_1 = 80100
const __IPHONE_8_2 = 80200
const __IPHONE_8_3 = 80300
const __IPHONE_8_4 = 80400
const __IPHONE_9_0 = 90000
const __IPHONE_9_1 = 90100
const __IPHONE_9_2 = 90200
const __IPHONE_9_3 = 90300
const __LAHF_SAHF__ = 1
const __LASTBRANCH_MAX = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.9406564584124654e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.2204460492503131e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.7976931348623157e+308
const __LDBL_MIN__ = 2.2250738585072014e-308
const __LDBL_NORM_MAX__ = 1.7976931348623157e+308
const __LITTLE_ENDIAN__ = 1
const __LLONG_WIDTH__ = 64
const __LONG_LONG_MAX__ = 9223372036854775807
const __LONG_MAX__ = 9223372036854775807
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MACH__ = 1
const __MAC_10_0 = 1000
const __MAC_10_1 = 1010
const __MAC_10_10 = 101000
const __MAC_10_10_2 = 101002
const __MAC_10_10_3 = 101003
const __MAC_10_11 = 101100
const __MAC_10_11_2 = 101102
const __MAC_10_11_3 = 101103
const __MAC_10_11_4 = 101104
const __MAC_10_12 = 101200
const __MAC_10_12_1 = 101201
const __MAC_10_12_2 = 101202
const __MAC_10_12_4 = 101204
const __MAC_10_13 = 101300
const __MAC_10_13_1 = 101301
const __MAC_10_13_2 = 101302
const __MAC_10_13_4 = 101304
const __MAC_10_14 = 101400
const __MAC_10_14_1 = 101401
const __MAC_10_14_4 = 101404
const __MAC_10_14_5 = 101405
const __MAC_10_14_6 = 101406
const __MAC_10_15 = 101500
const __MAC_10_15_1 = 101501
const __MAC_10_15_4 = 101504
const __MAC_10_16 = 101600
const __MAC_10_2 = 1020
const __MAC_10_3 = 1030
const __MAC_10_4 = 1040
const __MAC_10_5 = 1050
const __MAC_10_6 = 1060
const __MAC_10_7 = 1070
const __MAC_10_8 = 1080
const __MAC_10_9 = 1090
const __MAC_11_0 = 110000
const __MAC_11_1 = 110100
const __MAC_11_3 = 110300
const __MAC_11_4 = 110400
const __MAC_11_5 = 110500
const __MAC_11_6 = 110600
const __MAC_12_0 = 120000
const __MAC_12_1 = 120100
const __MAC_12_2 = 120200
const __MAC_12_3 = 120300
const __MAC_12_4 = 120400
const __MAC_12_5 = 120500
const __MAC_12_6 = 120600
const __MAC_12_7 = 120700
const __MAC_13_0 = 130000
const __MAC_13_1 = 130100
const __MAC_13_2 = 130200
const __MAC_13_3 = 130300
const __MAC_13_4 = 130400
const __MAC_13_5 = 130500
const __MAC_13_6 = 130600
const __MAC_13_7 = 130700
const __MAC_14_0 = 140000
const __MAC_14_1 = 140100
const __MAC_14_2 = 140200
const __MAC_14_3 = 140300
const __MAC_14_4 = 140400
const __MAC_14_5 = 140500
const __MAC_14_6 = 140600
const __MAC_14_7 = 140700
const __MAC_15_0 = 150000
const __MAC_15_1 = 150100
const __MAC_15_2 = 150200
const __MAC_15_3 = 150300
const __MAC_15_4 = 150400
const __MAC_15_5 = 150500
const __MAC_OS_X_VERSION_MAX_ALLOWED = "__MAC_15_5"
const __MAC_OS_X_VERSION_MIN_REQUIRED = "__ENVIRONMENT_MAC_OS_X_VERSION_MIN_REQUIRED__"
const __MEMORY_SCOPE_DEVICE = 1
const __MEMORY_SCOPE_SINGLE = 4
const __MEMORY_SCOPE_SYSTEM = 0
const __MEMORY_SCOPE_WRKGRP = 2
const __MEMORY_SCOPE_WVFRNT = 3
const __MMX__ = 1
const __NO_INLINE__ = 1
const __NO_MATH_ERRNO__ = 1
const __NO_MATH_INLINES = 1
const __OBJC_BOOL_IS_BOOL = 0
const __OPENCL_MEMORY_SCOPE_ALL_SVM_DEVICES = 3
const __OPENCL_MEMORY_SCOPE_DEVICE = 2
const __OPENCL_MEMORY_SCOPE_SUB_GROUP = 4
const __OPENCL_MEMORY_SCOPE_WORK_GROUP = 1
const __OPENCL_MEMORY_SCOPE_WORK_ITEM = 0
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __POINTER_WIDTH__ = 64
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTHREAD_ATTR_SIZE__ = 56
const __PTHREAD_CONDATTR_SIZE__ = 8
const __PTHREAD_COND_SIZE__ = 40
const __PTHREAD_MUTEXATTR_SIZE__ = 8
const __PTHREAD_MUTEX_SIZE__ = 56
const __PTHREAD_ONCE_SIZE__ = 8
const __PTHREAD_RWLOCKATTR_SIZE__ = 16
const __PTHREAD_RWLOCK_SIZE__ = 192
const __PTHREAD_SIZE__ = 8176
const __PTRDIFF_FMTd__ = "ld"
const __PTRDIFF_FMTi__ = "li"
const __PTRDIFF_MAX__ = 9223372036854775807
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 127
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
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
const __SIZE_FMTX__ = "lX"
const __SIZE_FMTo__ = "lo"
const __SIZE_FMTu__ = "lu"
const __SIZE_FMTx__ = "lx"
const __SIZE_MAX__ = 18446744073709551615
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE3__ = 1
const __SSE4_1__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP__ = 1
const __SSSE3__ = 1
const __STDC_EMBED_EMPTY__ = 2
const __STDC_EMBED_FOUND__ = 1
const __STDC_EMBED_NOT_FOUND__ = 0
const __STDC_HOSTED__ = 1
const __STDC_NO_THREADS__ = 1
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201112
const __STDC_WANT_LIB_EXT1__ = 1
const __STDC__ = 1
const __TVOS_10_0 = 100000
const __TVOS_10_0_1 = 100001
const __TVOS_10_1 = 100100
const __TVOS_10_2 = 100200
const __TVOS_11_0 = 110000
const __TVOS_11_1 = 110100
const __TVOS_11_2 = 110200
const __TVOS_11_3 = 110300
const __TVOS_11_4 = 110400
const __TVOS_12_0 = 120000
const __TVOS_12_1 = 120100
const __TVOS_12_2 = 120200
const __TVOS_12_3 = 120300
const __TVOS_12_4 = 120400
const __TVOS_13_0 = 130000
const __TVOS_13_2 = 130200
const __TVOS_13_3 = 130300
const __TVOS_13_4 = 130400
const __TVOS_14_0 = 140000
const __TVOS_14_1 = 140100
const __TVOS_14_2 = 140200
const __TVOS_14_3 = 140300
const __TVOS_14_5 = 140500
const __TVOS_14_6 = 140600
const __TVOS_14_7 = 140700
const __TVOS_15_0 = 150000
const __TVOS_15_1 = 150100
const __TVOS_15_2 = 150200
const __TVOS_15_3 = 150300
const __TVOS_15_4 = 150400
const __TVOS_15_5 = 150500
const __TVOS_15_6 = 150600
const __TVOS_16_0 = 160000
const __TVOS_16_1 = 160100
const __TVOS_16_2 = 160200
const __TVOS_16_3 = 160300
const __TVOS_16_4 = 160400
const __TVOS_16_5 = 160500
const __TVOS_16_6 = 160600
const __TVOS_17_0 = 170000
const __TVOS_17_1 = 170100
const __TVOS_17_2 = 170200
const __TVOS_17_3 = 170300
const __TVOS_17_4 = 170400
const __TVOS_17_5 = 170500
const __TVOS_17_6 = 170600
const __TVOS_18_0 = 180000
const __TVOS_18_1 = 180100
const __TVOS_18_2 = 180200
const __TVOS_18_3 = 180300
const __TVOS_18_4 = 180400
const __TVOS_18_5 = 180500
const __TVOS_9_0 = 90000
const __TVOS_9_1 = 90100
const __TVOS_9_2 = 90200
const __UINT16_FMTX__ = "hX"
const __UINT16_FMTo__ = "ho"
const __UINT16_FMTu__ = "hu"
const __UINT16_FMTx__ = "hx"
const __UINT16_MAX__ = 65535
const __UINT32_C_SUFFIX__ = "U"
const __UINT32_FMTX__ = "X"
const __UINT32_FMTo__ = "o"
const __UINT32_FMTu__ = "u"
const __UINT32_FMTx__ = "x"
const __UINT32_MAX__ = 4294967295
const __UINT64_C_SUFFIX__ = "ULL"
const __UINT64_FMTX__ = "llX"
const __UINT64_FMTo__ = "llo"
const __UINT64_FMTu__ = "llu"
const __UINT64_FMTx__ = "llx"
const __UINT64_MAX__ = "18446744073709551615U"
const __UINT8_FMTX__ = "hhX"
const __UINT8_FMTo__ = "hho"
const __UINT8_FMTu__ = "hhu"
const __UINT8_FMTx__ = "hhx"
const __UINT8_MAX__ = 255
const __UINTMAX_C_SUFFIX__ = "UL"
const __UINTMAX_FMTX__ = "lX"
const __UINTMAX_FMTo__ = "lo"
const __UINTMAX_FMTu__ = "lu"
const __UINTMAX_FMTx__ = "lx"
const __UINTMAX_MAX__ = 18446744073709551615
const __UINTMAX_WIDTH__ = 64
const __UINTPTR_FMTX__ = "lX"
const __UINTPTR_FMTo__ = "lo"
const __UINTPTR_FMTu__ = "lu"
const __UINTPTR_FMTx__ = "lx"
const __UINTPTR_MAX__ = 18446744073709551615
const __UINTPTR_WIDTH__ = 64
const __UINT_FAST16_FMTX__ = "hX"
const __UINT_FAST16_FMTo__ = "ho"
const __UINT_FAST16_FMTu__ = "hu"
const __UINT_FAST16_FMTx__ = "hx"
const __UINT_FAST16_MAX__ = 65535
const __UINT_FAST32_FMTX__ = "X"
const __UINT_FAST32_FMTo__ = "o"
const __UINT_FAST32_FMTu__ = "u"
const __UINT_FAST32_FMTx__ = "x"
const __UINT_FAST32_MAX__ = 4294967295
const __UINT_FAST64_FMTX__ = "llX"
const __UINT_FAST64_FMTo__ = "llo"
const __UINT_FAST64_FMTu__ = "llu"
const __UINT_FAST64_FMTx__ = "llx"
const __UINT_FAST64_MAX__ = "18446744073709551615U"
const __UINT_FAST8_FMTX__ = "hhX"
const __UINT_FAST8_FMTo__ = "hho"
const __UINT_FAST8_FMTu__ = "hhu"
const __UINT_FAST8_FMTx__ = "hhx"
const __UINT_FAST8_MAX__ = 255
const __UINT_LEAST16_FMTX__ = "hX"
const __UINT_LEAST16_FMTo__ = "ho"
const __UINT_LEAST16_FMTu__ = "hu"
const __UINT_LEAST16_FMTx__ = "hx"
const __UINT_LEAST16_MAX__ = 65535
const __UINT_LEAST32_FMTX__ = "X"
const __UINT_LEAST32_FMTo__ = "o"
const __UINT_LEAST32_FMTu__ = "u"
const __UINT_LEAST32_FMTx__ = "x"
const __UINT_LEAST32_MAX__ = 4294967295
const __UINT_LEAST64_FMTX__ = "llX"
const __UINT_LEAST64_FMTo__ = "llo"
const __UINT_LEAST64_FMTu__ = "llu"
const __UINT_LEAST64_FMTx__ = "llx"
const __UINT_LEAST64_MAX = "UINT64_MAX"
const __UINT_LEAST64_MAX__ = "18446744073709551615U"
const __UINT_LEAST8_FMTX__ = "hhX"
const __UINT_LEAST8_FMTo__ = "hho"
const __UINT_LEAST8_FMTu__ = "hhu"
const __UINT_LEAST8_FMTx__ = "hhx"
const __UINT_LEAST8_MAX__ = 255
const __USER_LABEL_PREFIX__ = "_"
const __VERSION__ = "Apple LLVM 17.0.0 (clang-1700.0.13.5)"
const __VISIONOS_1_0 = 10000
const __VISIONOS_1_1 = 10100
const __VISIONOS_1_2 = 10200
const __VISIONOS_1_3 = 10300
const __VISIONOS_2_0 = 20000
const __VISIONOS_2_1 = 20100
const __VISIONOS_2_2 = 20200
const __VISIONOS_2_3 = 20300
const __VISIONOS_2_4 = 20400
const __VISIONOS_2_5 = 20500
const __WATCHOS_10_0 = 100000
const __WATCHOS_10_1 = 100100
const __WATCHOS_10_2 = 100200
const __WATCHOS_10_3 = 100300
const __WATCHOS_10_4 = 100400
const __WATCHOS_10_5 = 100500
const __WATCHOS_10_6 = 100600
const __WATCHOS_10_7 = 100700
const __WATCHOS_11_0 = 110000
const __WATCHOS_11_1 = 110100
const __WATCHOS_11_2 = 110200
const __WATCHOS_11_3 = 110300
const __WATCHOS_11_4 = 110400
const __WATCHOS_11_5 = 110500
const __WATCHOS_1_0 = 10000
const __WATCHOS_2_0 = 20000
const __WATCHOS_2_1 = 20100
const __WATCHOS_2_2 = 20200
const __WATCHOS_3_0 = 30000
const __WATCHOS_3_1 = 30100
const __WATCHOS_3_1_1 = 30101
const __WATCHOS_3_2 = 30200
const __WATCHOS_4_0 = 40000
const __WATCHOS_4_1 = 40100
const __WATCHOS_4_2 = 40200
const __WATCHOS_4_3 = 40300
const __WATCHOS_5_0 = 50000
const __WATCHOS_5_1 = 50100
const __WATCHOS_5_2 = 50200
const __WATCHOS_5_3 = 50300
const __WATCHOS_6_0 = 60000
const __WATCHOS_6_1 = 60100
const __WATCHOS_6_2 = 60200
const __WATCHOS_7_0 = 70000
const __WATCHOS_7_1 = 70100
const __WATCHOS_7_2 = 70200
const __WATCHOS_7_3 = 70300
const __WATCHOS_7_4 = 70400
const __WATCHOS_7_5 = 70500
const __WATCHOS_7_6 = 70600
const __WATCHOS_8_0 = 80000
const __WATCHOS_8_1 = 80100
const __WATCHOS_8_3 = 80300
const __WATCHOS_8_4 = 80400
const __WATCHOS_8_5 = 80500
const __WATCHOS_8_6 = 80600
const __WATCHOS_8_7 = 80700
const __WATCHOS_8_8 = 80800
const __WATCHOS_9_0 = 90000
const __WATCHOS_9_1 = 90100
const __WATCHOS_9_2 = 90200
const __WATCHOS_9_3 = 90300
const __WATCHOS_9_4 = 90400
const __WATCHOS_9_5 = 90500
const __WATCHOS_9_6 = 90600
const __WCHAR_MAX__ = 2147483647
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 2147483647
const __WINT_TYPE__ = "int"
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __apple_build_version__ = 17000013
const __bool_true_false_are_defined = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 17
const __clang_minor__ = 0
const __clang_patchlevel__ = 0
const __clang_version__ = "17.0.0 (clang-1700.0.13.5)"
const __clang_wide_literal_encoding__ = "UTF-32"
const __code_model_small__ = 1
const __const = "const"
const __core2 = 1
const __core2__ = 1
const __has_bounds_safety_attributes = 0
const __has_ptrcheck = 0
const __has_safe_buffers = 0
const __header_inline = "inline"
const __llvm__ = 1
const __nonnull = "_Nonnull"
const __null_unspecified = "_Null_unspecified"
const __nullable = "_Nullable"
const __pic__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __signed = "signed"
const __tune_core2__ = 1
const __volatile = "volatile"
const __x86_64 = 1
const __x86_64__ = 1
const bool1 = "_Bool"
const chan1 = "chan_token"
const defer1 = "defer_token"
const fallthrough1 = "fallthrough_token"
const false1 = 0
const func1 = "func_token"
const go1 = "go_token"
const import1 = "import_token"
const interface1 = "interface_token"
const map1 = "map_token"
const package1 = "package_token"
const range1 = "range_token"
const ru_first = "ru_ixrss"
const ru_last = "ru_nivcsw"
const select2 = "select_token"
const sv_onstack = "sv_flags"
const true1 = 1
const ts_builtin_sym_end = 0
const type1 = "type_token"
const var1 = "var_token"

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type int64_t = int64

type uint64_t = uint64

type int_least64_t = int64

type uint_least64_t = uint64

type int_fast64_t = int64

type uint_fast64_t = uint64

type int32_t = int32

type uint32_t = uint32

type int_least32_t = int32

type uint_least32_t = uint32

type int_fast32_t = int32

type uint_fast32_t = uint32

type int16_t = int16

type uint16_t = uint16

type int_least16_t = int16

type uint_least16_t = uint16

type int_fast16_t = int16

type uint_fast16_t = uint16

type int8_t = int8

type uint8_t = uint8

type int_least8_t = int8

type uint_least8_t = uint8

type int_fast8_t = int8

type uint_fast8_t = uint8

type intptr_t = int64

type uintptr_t = uint64

type intmax_t = int64

type uintmax_t = uint64

type __int8_t = int8

type __uint8_t = uint8

type __int16_t = int16

type __uint16_t = uint16

type __int32_t = int32

type __uint32_t = uint32

type __int64_t = int64

type __uint64_t = uint64

type __darwin_intptr_t = int64

type __darwin_natural_t = uint32

type __darwin_ct_rune_t = int32

type __mbstate_t = struct {
	F_mbstateL  [0]int64
	F__mbstate8 [128]int8
}

type __darwin_mbstate_t = struct {
	F_mbstateL  [0]int64
	F__mbstate8 [128]int8
}

type __darwin_ptrdiff_t = int64

type __darwin_size_t = uint64

type __darwin_va_list = uintptr

type __darwin_wchar_t = int32

type __darwin_rune_t = int32

type __darwin_wint_t = int32

type __darwin_clock_t = uint64

type __darwin_socklen_t = uint32

type __darwin_ssize_t = int64

type __darwin_time_t = int64

type __darwin_blkcnt_t = int64

type __darwin_blksize_t = int32

type __darwin_dev_t = int32

type __darwin_fsblkcnt_t = uint32

type __darwin_fsfilcnt_t = uint32

type __darwin_gid_t = uint32

type __darwin_id_t = uint32

type __darwin_ino64_t = uint64

type __darwin_ino_t = uint64

type __darwin_mach_port_name_t = uint32

type __darwin_mach_port_t = uint32

type __darwin_mode_t = uint16

type __darwin_off_t = int64

type __darwin_pid_t = int32

type __darwin_sigset_t = uint32

type __darwin_suseconds_t = int32

type __darwin_uid_t = uint32

type __darwin_useconds_t = uint32

type __darwin_uuid_t = [16]uint8

type __darwin_uuid_string_t = [37]int8

type __darwin_pthread_handler_rec = struct {
	F__routine uintptr
	F__arg     uintptr
	F__next    uintptr
}

type _opaque_pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type _opaque_pthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
}

type _opaque_pthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type _opaque_pthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type _opaque_pthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type _opaque_pthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type _opaque_pthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
}

type _opaque_pthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
}

type _opaque_pthread_t = struct {
	F__sig           int64
	F__cleanup_stack uintptr
	F__opaque        [8176]int8
}

type __darwin_pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type __darwin_pthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
}

type __darwin_pthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type __darwin_pthread_key_t = uint64

type __darwin_pthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type __darwin_pthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type __darwin_pthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
}

type __darwin_pthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
}

type __darwin_pthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
}

type __darwin_pthread_t = uintptr

type __darwin_nl_item = int32

type __darwin_wctrans_t = int32

type __darwin_wctype_t = uint32

type idtype_t = int32

const P_ALL = 0
const P_PID = 1
const P_PGID = 2

type pid_t = int32

type id_t = uint32

type sig_atomic_t = int32

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type u_int64_t = uint64

type register_t = int64

type user_addr_t = uint64

type user_size_t = uint64

type user_ssize_t = int64

type user_long_t = int64

type user_ulong_t = uint64

type user_time_t = int64

type user_off_t = int64

type syscall_arg_t = uint64

type __darwin_i386_thread_state = struct {
	F__eax    uint32
	F__ebx    uint32
	F__ecx    uint32
	F__edx    uint32
	F__edi    uint32
	F__esi    uint32
	F__ebp    uint32
	F__esp    uint32
	F__ss     uint32
	F__eflags uint32
	F__eip    uint32
	F__cs     uint32
	F__ds     uint32
	F__es     uint32
	F__fs     uint32
	F__gs     uint32
}

type __darwin_fp_control = struct {
	F__ccgo0 uint16
}

type __darwin_fp_control_t = struct {
	F__ccgo0 uint16
}

type __darwin_fp_status = struct {
	F__ccgo0 uint16
}

type __darwin_fp_status_t = struct {
	F__ccgo0 uint16
}

type __darwin_mmst_reg = struct {
	F__mmst_reg  [10]int8
	F__mmst_rsrv [6]int8
}

type __darwin_xmm_reg = struct {
	F__xmm_reg [16]int8
}

type __darwin_ymm_reg = struct {
	F__ymm_reg [32]int8
}

type __darwin_zmm_reg = struct {
	F__zmm_reg [64]int8
}

type __darwin_opmask_reg = struct {
	F__opmask_reg [8]int8
}

type __darwin_i386_float_state = struct {
	F__fpu_reserved  [2]int32
	F__fpu_fcw       __darwin_fp_control
	F__fpu_fsw       __darwin_fp_status
	F__fpu_ftw       __uint8_t
	F__fpu_rsrv1     __uint8_t
	F__fpu_fop       __uint16_t
	F__fpu_ip        __uint32_t
	F__fpu_cs        __uint16_t
	F__fpu_rsrv2     __uint16_t
	F__fpu_dp        __uint32_t
	F__fpu_ds        __uint16_t
	F__fpu_rsrv3     __uint16_t
	F__fpu_mxcsr     __uint32_t
	F__fpu_mxcsrmask __uint32_t
	F__fpu_stmm0     __darwin_mmst_reg
	F__fpu_stmm1     __darwin_mmst_reg
	F__fpu_stmm2     __darwin_mmst_reg
	F__fpu_stmm3     __darwin_mmst_reg
	F__fpu_stmm4     __darwin_mmst_reg
	F__fpu_stmm5     __darwin_mmst_reg
	F__fpu_stmm6     __darwin_mmst_reg
	F__fpu_stmm7     __darwin_mmst_reg
	F__fpu_xmm0      __darwin_xmm_reg
	F__fpu_xmm1      __darwin_xmm_reg
	F__fpu_xmm2      __darwin_xmm_reg
	F__fpu_xmm3      __darwin_xmm_reg
	F__fpu_xmm4      __darwin_xmm_reg
	F__fpu_xmm5      __darwin_xmm_reg
	F__fpu_xmm6      __darwin_xmm_reg
	F__fpu_xmm7      __darwin_xmm_reg
	F__fpu_rsrv4     [224]int8
	F__fpu_reserved1 int32
}

type __darwin_i386_avx_state = struct {
	F__fpu_reserved  [2]int32
	F__fpu_fcw       __darwin_fp_control
	F__fpu_fsw       __darwin_fp_status
	F__fpu_ftw       __uint8_t
	F__fpu_rsrv1     __uint8_t
	F__fpu_fop       __uint16_t
	F__fpu_ip        __uint32_t
	F__fpu_cs        __uint16_t
	F__fpu_rsrv2     __uint16_t
	F__fpu_dp        __uint32_t
	F__fpu_ds        __uint16_t
	F__fpu_rsrv3     __uint16_t
	F__fpu_mxcsr     __uint32_t
	F__fpu_mxcsrmask __uint32_t
	F__fpu_stmm0     __darwin_mmst_reg
	F__fpu_stmm1     __darwin_mmst_reg
	F__fpu_stmm2     __darwin_mmst_reg
	F__fpu_stmm3     __darwin_mmst_reg
	F__fpu_stmm4     __darwin_mmst_reg
	F__fpu_stmm5     __darwin_mmst_reg
	F__fpu_stmm6     __darwin_mmst_reg
	F__fpu_stmm7     __darwin_mmst_reg
	F__fpu_xmm0      __darwin_xmm_reg
	F__fpu_xmm1      __darwin_xmm_reg
	F__fpu_xmm2      __darwin_xmm_reg
	F__fpu_xmm3      __darwin_xmm_reg
	F__fpu_xmm4      __darwin_xmm_reg
	F__fpu_xmm5      __darwin_xmm_reg
	F__fpu_xmm6      __darwin_xmm_reg
	F__fpu_xmm7      __darwin_xmm_reg
	F__fpu_rsrv4     [224]int8
	F__fpu_reserved1 int32
	F__avx_reserved1 [64]int8
	F__fpu_ymmh0     __darwin_xmm_reg
	F__fpu_ymmh1     __darwin_xmm_reg
	F__fpu_ymmh2     __darwin_xmm_reg
	F__fpu_ymmh3     __darwin_xmm_reg
	F__fpu_ymmh4     __darwin_xmm_reg
	F__fpu_ymmh5     __darwin_xmm_reg
	F__fpu_ymmh6     __darwin_xmm_reg
	F__fpu_ymmh7     __darwin_xmm_reg
}

type __darwin_i386_avx512_state = struct {
	F__fpu_reserved  [2]int32
	F__fpu_fcw       __darwin_fp_control
	F__fpu_fsw       __darwin_fp_status
	F__fpu_ftw       __uint8_t
	F__fpu_rsrv1     __uint8_t
	F__fpu_fop       __uint16_t
	F__fpu_ip        __uint32_t
	F__fpu_cs        __uint16_t
	F__fpu_rsrv2     __uint16_t
	F__fpu_dp        __uint32_t
	F__fpu_ds        __uint16_t
	F__fpu_rsrv3     __uint16_t
	F__fpu_mxcsr     __uint32_t
	F__fpu_mxcsrmask __uint32_t
	F__fpu_stmm0     __darwin_mmst_reg
	F__fpu_stmm1     __darwin_mmst_reg
	F__fpu_stmm2     __darwin_mmst_reg
	F__fpu_stmm3     __darwin_mmst_reg
	F__fpu_stmm4     __darwin_mmst_reg
	F__fpu_stmm5     __darwin_mmst_reg
	F__fpu_stmm6     __darwin_mmst_reg
	F__fpu_stmm7     __darwin_mmst_reg
	F__fpu_xmm0      __darwin_xmm_reg
	F__fpu_xmm1      __darwin_xmm_reg
	F__fpu_xmm2      __darwin_xmm_reg
	F__fpu_xmm3      __darwin_xmm_reg
	F__fpu_xmm4      __darwin_xmm_reg
	F__fpu_xmm5      __darwin_xmm_reg
	F__fpu_xmm6      __darwin_xmm_reg
	F__fpu_xmm7      __darwin_xmm_reg
	F__fpu_rsrv4     [224]int8
	F__fpu_reserved1 int32
	F__avx_reserved1 [64]int8
	F__fpu_ymmh0     __darwin_xmm_reg
	F__fpu_ymmh1     __darwin_xmm_reg
	F__fpu_ymmh2     __darwin_xmm_reg
	F__fpu_ymmh3     __darwin_xmm_reg
	F__fpu_ymmh4     __darwin_xmm_reg
	F__fpu_ymmh5     __darwin_xmm_reg
	F__fpu_ymmh6     __darwin_xmm_reg
	F__fpu_ymmh7     __darwin_xmm_reg
	F__fpu_k0        __darwin_opmask_reg
	F__fpu_k1        __darwin_opmask_reg
	F__fpu_k2        __darwin_opmask_reg
	F__fpu_k3        __darwin_opmask_reg
	F__fpu_k4        __darwin_opmask_reg
	F__fpu_k5        __darwin_opmask_reg
	F__fpu_k6        __darwin_opmask_reg
	F__fpu_k7        __darwin_opmask_reg
	F__fpu_zmmh0     __darwin_ymm_reg
	F__fpu_zmmh1     __darwin_ymm_reg
	F__fpu_zmmh2     __darwin_ymm_reg
	F__fpu_zmmh3     __darwin_ymm_reg
	F__fpu_zmmh4     __darwin_ymm_reg
	F__fpu_zmmh5     __darwin_ymm_reg
	F__fpu_zmmh6     __darwin_ymm_reg
	F__fpu_zmmh7     __darwin_ymm_reg
}

type __darwin_i386_exception_state = struct {
	F__trapno     __uint16_t
	F__cpu        __uint16_t
	F__err        __uint32_t
	F__faultvaddr __uint32_t
}

type __darwin_x86_debug_state32 = struct {
	F__dr0 uint32
	F__dr1 uint32
	F__dr2 uint32
	F__dr3 uint32
	F__dr4 uint32
	F__dr5 uint32
	F__dr6 uint32
	F__dr7 uint32
}

type __x86_instruction_state = struct {
	F__insn_stream_valid_bytes int32
	F__insn_offset             int32
	F__out_of_synch            int32
	F__insn_bytes              [2380]__uint8_t
	F__insn_cacheline          [64]__uint8_t
}

type __last_branch_record = struct {
	F__from_ip __uint64_t
	F__to_ip   __uint64_t
	F__ccgo16  uint32
}

type __last_branch_state = struct {
	F__lbr_count int32
	F__ccgo4     uint32
	F__lbrs      [32]__last_branch_record
}

type __x86_pagein_state = struct {
	F__pagein_error int32
}

type __darwin_x86_thread_state64 = struct {
	F__rax    __uint64_t
	F__rbx    __uint64_t
	F__rcx    __uint64_t
	F__rdx    __uint64_t
	F__rdi    __uint64_t
	F__rsi    __uint64_t
	F__rbp    __uint64_t
	F__rsp    __uint64_t
	F__r8     __uint64_t
	F__r9     __uint64_t
	F__r10    __uint64_t
	F__r11    __uint64_t
	F__r12    __uint64_t
	F__r13    __uint64_t
	F__r14    __uint64_t
	F__r15    __uint64_t
	F__rip    __uint64_t
	F__rflags __uint64_t
	F__cs     __uint64_t
	F__fs     __uint64_t
	F__gs     __uint64_t
}

type __darwin_x86_thread_full_state64 = struct {
	F__ss64   __darwin_x86_thread_state64
	F__ds     __uint64_t
	F__es     __uint64_t
	F__ss     __uint64_t
	F__gsbase __uint64_t
}

type __darwin_x86_float_state64 = struct {
	F__fpu_reserved  [2]int32
	F__fpu_fcw       __darwin_fp_control
	F__fpu_fsw       __darwin_fp_status
	F__fpu_ftw       __uint8_t
	F__fpu_rsrv1     __uint8_t
	F__fpu_fop       __uint16_t
	F__fpu_ip        __uint32_t
	F__fpu_cs        __uint16_t
	F__fpu_rsrv2     __uint16_t
	F__fpu_dp        __uint32_t
	F__fpu_ds        __uint16_t
	F__fpu_rsrv3     __uint16_t
	F__fpu_mxcsr     __uint32_t
	F__fpu_mxcsrmask __uint32_t
	F__fpu_stmm0     __darwin_mmst_reg
	F__fpu_stmm1     __darwin_mmst_reg
	F__fpu_stmm2     __darwin_mmst_reg
	F__fpu_stmm3     __darwin_mmst_reg
	F__fpu_stmm4     __darwin_mmst_reg
	F__fpu_stmm5     __darwin_mmst_reg
	F__fpu_stmm6     __darwin_mmst_reg
	F__fpu_stmm7     __darwin_mmst_reg
	F__fpu_xmm0      __darwin_xmm_reg
	F__fpu_xmm1      __darwin_xmm_reg
	F__fpu_xmm2      __darwin_xmm_reg
	F__fpu_xmm3      __darwin_xmm_reg
	F__fpu_xmm4      __darwin_xmm_reg
	F__fpu_xmm5      __darwin_xmm_reg
	F__fpu_xmm6      __darwin_xmm_reg
	F__fpu_xmm7      __darwin_xmm_reg
	F__fpu_xmm8      __darwin_xmm_reg
	F__fpu_xmm9      __darwin_xmm_reg
	F__fpu_xmm10     __darwin_xmm_reg
	F__fpu_xmm11     __darwin_xmm_reg
	F__fpu_xmm12     __darwin_xmm_reg
	F__fpu_xmm13     __darwin_xmm_reg
	F__fpu_xmm14     __darwin_xmm_reg
	F__fpu_xmm15     __darwin_xmm_reg
	F__fpu_rsrv4     [96]int8
	F__fpu_reserved1 int32
}

type __darwin_x86_avx_state64 = struct {
	F__fpu_reserved  [2]int32
	F__fpu_fcw       __darwin_fp_control
	F__fpu_fsw       __darwin_fp_status
	F__fpu_ftw       __uint8_t
	F__fpu_rsrv1     __uint8_t
	F__fpu_fop       __uint16_t
	F__fpu_ip        __uint32_t
	F__fpu_cs        __uint16_t
	F__fpu_rsrv2     __uint16_t
	F__fpu_dp        __uint32_t
	F__fpu_ds        __uint16_t
	F__fpu_rsrv3     __uint16_t
	F__fpu_mxcsr     __uint32_t
	F__fpu_mxcsrmask __uint32_t
	F__fpu_stmm0     __darwin_mmst_reg
	F__fpu_stmm1     __darwin_mmst_reg
	F__fpu_stmm2     __darwin_mmst_reg
	F__fpu_stmm3     __darwin_mmst_reg
	F__fpu_stmm4     __darwin_mmst_reg
	F__fpu_stmm5     __darwin_mmst_reg
	F__fpu_stmm6     __darwin_mmst_reg
	F__fpu_stmm7     __darwin_mmst_reg
	F__fpu_xmm0      __darwin_xmm_reg
	F__fpu_xmm1      __darwin_xmm_reg
	F__fpu_xmm2      __darwin_xmm_reg
	F__fpu_xmm3      __darwin_xmm_reg
	F__fpu_xmm4      __darwin_xmm_reg
	F__fpu_xmm5      __darwin_xmm_reg
	F__fpu_xmm6      __darwin_xmm_reg
	F__fpu_xmm7      __darwin_xmm_reg
	F__fpu_xmm8      __darwin_xmm_reg
	F__fpu_xmm9      __darwin_xmm_reg
	F__fpu_xmm10     __darwin_xmm_reg
	F__fpu_xmm11     __darwin_xmm_reg
	F__fpu_xmm12     __darwin_xmm_reg
	F__fpu_xmm13     __darwin_xmm_reg
	F__fpu_xmm14     __darwin_xmm_reg
	F__fpu_xmm15     __darwin_xmm_reg
	F__fpu_rsrv4     [96]int8
	F__fpu_reserved1 int32
	F__avx_reserved1 [64]int8
	F__fpu_ymmh0     __darwin_xmm_reg
	F__fpu_ymmh1     __darwin_xmm_reg
	F__fpu_ymmh2     __darwin_xmm_reg
	F__fpu_ymmh3     __darwin_xmm_reg
	F__fpu_ymmh4     __darwin_xmm_reg
	F__fpu_ymmh5     __darwin_xmm_reg
	F__fpu_ymmh6     __darwin_xmm_reg
	F__fpu_ymmh7     __darwin_xmm_reg
	F__fpu_ymmh8     __darwin_xmm_reg
	F__fpu_ymmh9     __darwin_xmm_reg
	F__fpu_ymmh10    __darwin_xmm_reg
	F__fpu_ymmh11    __darwin_xmm_reg
	F__fpu_ymmh12    __darwin_xmm_reg
	F__fpu_ymmh13    __darwin_xmm_reg
	F__fpu_ymmh14    __darwin_xmm_reg
	F__fpu_ymmh15    __darwin_xmm_reg
}

type __darwin_x86_avx512_state64 = struct {
	F__fpu_reserved  [2]int32
	F__fpu_fcw       __darwin_fp_control
	F__fpu_fsw       __darwin_fp_status
	F__fpu_ftw       __uint8_t
	F__fpu_rsrv1     __uint8_t
	F__fpu_fop       __uint16_t
	F__fpu_ip        __uint32_t
	F__fpu_cs        __uint16_t
	F__fpu_rsrv2     __uint16_t
	F__fpu_dp        __uint32_t
	F__fpu_ds        __uint16_t
	F__fpu_rsrv3     __uint16_t
	F__fpu_mxcsr     __uint32_t
	F__fpu_mxcsrmask __uint32_t
	F__fpu_stmm0     __darwin_mmst_reg
	F__fpu_stmm1     __darwin_mmst_reg
	F__fpu_stmm2     __darwin_mmst_reg
	F__fpu_stmm3     __darwin_mmst_reg
	F__fpu_stmm4     __darwin_mmst_reg
	F__fpu_stmm5     __darwin_mmst_reg
	F__fpu_stmm6     __darwin_mmst_reg
	F__fpu_stmm7     __darwin_mmst_reg
	F__fpu_xmm0      __darwin_xmm_reg
	F__fpu_xmm1      __darwin_xmm_reg
	F__fpu_xmm2      __darwin_xmm_reg
	F__fpu_xmm3      __darwin_xmm_reg
	F__fpu_xmm4      __darwin_xmm_reg
	F__fpu_xmm5      __darwin_xmm_reg
	F__fpu_xmm6      __darwin_xmm_reg
	F__fpu_xmm7      __darwin_xmm_reg
	F__fpu_xmm8      __darwin_xmm_reg
	F__fpu_xmm9      __darwin_xmm_reg
	F__fpu_xmm10     __darwin_xmm_reg
	F__fpu_xmm11     __darwin_xmm_reg
	F__fpu_xmm12     __darwin_xmm_reg
	F__fpu_xmm13     __darwin_xmm_reg
	F__fpu_xmm14     __darwin_xmm_reg
	F__fpu_xmm15     __darwin_xmm_reg
	F__fpu_rsrv4     [96]int8
	F__fpu_reserved1 int32
	F__avx_reserved1 [64]int8
	F__fpu_ymmh0     __darwin_xmm_reg
	F__fpu_ymmh1     __darwin_xmm_reg
	F__fpu_ymmh2     __darwin_xmm_reg
	F__fpu_ymmh3     __darwin_xmm_reg
	F__fpu_ymmh4     __darwin_xmm_reg
	F__fpu_ymmh5     __darwin_xmm_reg
	F__fpu_ymmh6     __darwin_xmm_reg
	F__fpu_ymmh7     __darwin_xmm_reg
	F__fpu_ymmh8     __darwin_xmm_reg
	F__fpu_ymmh9     __darwin_xmm_reg
	F__fpu_ymmh10    __darwin_xmm_reg
	F__fpu_ymmh11    __darwin_xmm_reg
	F__fpu_ymmh12    __darwin_xmm_reg
	F__fpu_ymmh13    __darwin_xmm_reg
	F__fpu_ymmh14    __darwin_xmm_reg
	F__fpu_ymmh15    __darwin_xmm_reg
	F__fpu_k0        __darwin_opmask_reg
	F__fpu_k1        __darwin_opmask_reg
	F__fpu_k2        __darwin_opmask_reg
	F__fpu_k3        __darwin_opmask_reg
	F__fpu_k4        __darwin_opmask_reg
	F__fpu_k5        __darwin_opmask_reg
	F__fpu_k6        __darwin_opmask_reg
	F__fpu_k7        __darwin_opmask_reg
	F__fpu_zmmh0     __darwin_ymm_reg
	F__fpu_zmmh1     __darwin_ymm_reg
	F__fpu_zmmh2     __darwin_ymm_reg
	F__fpu_zmmh3     __darwin_ymm_reg
	F__fpu_zmmh4     __darwin_ymm_reg
	F__fpu_zmmh5     __darwin_ymm_reg
	F__fpu_zmmh6     __darwin_ymm_reg
	F__fpu_zmmh7     __darwin_ymm_reg
	F__fpu_zmmh8     __darwin_ymm_reg
	F__fpu_zmmh9     __darwin_ymm_reg
	F__fpu_zmmh10    __darwin_ymm_reg
	F__fpu_zmmh11    __darwin_ymm_reg
	F__fpu_zmmh12    __darwin_ymm_reg
	F__fpu_zmmh13    __darwin_ymm_reg
	F__fpu_zmmh14    __darwin_ymm_reg
	F__fpu_zmmh15    __darwin_ymm_reg
	F__fpu_zmm16     __darwin_zmm_reg
	F__fpu_zmm17     __darwin_zmm_reg
	F__fpu_zmm18     __darwin_zmm_reg
	F__fpu_zmm19     __darwin_zmm_reg
	F__fpu_zmm20     __darwin_zmm_reg
	F__fpu_zmm21     __darwin_zmm_reg
	F__fpu_zmm22     __darwin_zmm_reg
	F__fpu_zmm23     __darwin_zmm_reg
	F__fpu_zmm24     __darwin_zmm_reg
	F__fpu_zmm25     __darwin_zmm_reg
	F__fpu_zmm26     __darwin_zmm_reg
	F__fpu_zmm27     __darwin_zmm_reg
	F__fpu_zmm28     __darwin_zmm_reg
	F__fpu_zmm29     __darwin_zmm_reg
	F__fpu_zmm30     __darwin_zmm_reg
	F__fpu_zmm31     __darwin_zmm_reg
}

type __darwin_x86_exception_state64 = struct {
	F__trapno     __uint16_t
	F__cpu        __uint16_t
	F__err        __uint32_t
	F__faultvaddr __uint64_t
}

type __darwin_x86_debug_state64 = struct {
	F__dr0 __uint64_t
	F__dr1 __uint64_t
	F__dr2 __uint64_t
	F__dr3 __uint64_t
	F__dr4 __uint64_t
	F__dr5 __uint64_t
	F__dr6 __uint64_t
	F__dr7 __uint64_t
}

type __darwin_x86_cpmu_state64 = struct {
	F__ctrs [16]__uint64_t
}

type __darwin_mcontext32 = struct {
	F__es __darwin_i386_exception_state
	F__ss __darwin_i386_thread_state
	F__fs __darwin_i386_float_state
}

type __darwin_mcontext_avx32 = struct {
	F__es __darwin_i386_exception_state
	F__ss __darwin_i386_thread_state
	F__fs __darwin_i386_avx_state
}

type __darwin_mcontext_avx512_32 = struct {
	F__es __darwin_i386_exception_state
	F__ss __darwin_i386_thread_state
	F__fs __darwin_i386_avx512_state
}

type __darwin_mcontext64 = struct {
	F__es __darwin_x86_exception_state64
	F__ss __darwin_x86_thread_state64
	F__fs __darwin_x86_float_state64
}

type __darwin_mcontext64_full = struct {
	F__es __darwin_x86_exception_state64
	F__ss __darwin_x86_thread_full_state64
	F__fs __darwin_x86_float_state64
}

type __darwin_mcontext_avx64 = struct {
	F__es __darwin_x86_exception_state64
	F__ss __darwin_x86_thread_state64
	F__fs __darwin_x86_avx_state64
}

type __darwin_mcontext_avx64_full = struct {
	F__es __darwin_x86_exception_state64
	F__ss __darwin_x86_thread_full_state64
	F__fs __darwin_x86_avx_state64
}

type __darwin_mcontext_avx512_64 = struct {
	F__es __darwin_x86_exception_state64
	F__ss __darwin_x86_thread_state64
	F__fs __darwin_x86_avx512_state64
}

type __darwin_mcontext_avx512_64_full = struct {
	F__es __darwin_x86_exception_state64
	F__ss __darwin_x86_thread_full_state64
	F__fs __darwin_x86_avx512_state64
}

type mcontext_t = uintptr

type pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
}

type __darwin_sigaltstack = struct {
	Fss_sp    uintptr
	Fss_size  __darwin_size_t
	Fss_flags int32
}

type stack_t = struct {
	Fss_sp    uintptr
	Fss_size  __darwin_size_t
	Fss_flags int32
}

type __darwin_ucontext = struct {
	Fuc_onstack  int32
	Fuc_sigmask  __darwin_sigset_t
	Fuc_stack    __darwin_sigaltstack
	Fuc_link     uintptr
	Fuc_mcsize   __darwin_size_t
	Fuc_mcontext uintptr
}

type ucontext_t = struct {
	Fuc_onstack  int32
	Fuc_sigmask  __darwin_sigset_t
	Fuc_stack    __darwin_sigaltstack
	Fuc_link     uintptr
	Fuc_mcsize   __darwin_size_t
	Fuc_mcontext uintptr
}

type sigset_t = uint32

type size_t = uint64

type uid_t = uint32

type sigval = struct {
	Fsival_ptr   [0]uintptr
	Fsival_int   int32
	F__ccgo_pad2 [4]byte
}

type sigevent = struct {
	Fsigev_notify            int32
	Fsigev_signo             int32
	Fsigev_value             sigval
	Fsigev_notify_function   uintptr
	Fsigev_notify_attributes uintptr
}

type siginfo_t = struct {
	Fsi_signo  int32
	Fsi_errno  int32
	Fsi_code   int32
	Fsi_pid    pid_t
	Fsi_uid    uid_t
	Fsi_status int32
	Fsi_addr   uintptr
	Fsi_value  sigval
	Fsi_band   int64
	F__pad     [7]uint64
}

type __siginfo = siginfo_t

type __sigaction_u = struct {
	F__sa_sigaction [0]uintptr
	F__sa_handler   uintptr
}

type __sigaction = struct {
	F__sigaction_u __sigaction_u
	Fsa_tramp      uintptr
	Fsa_mask       sigset_t
	Fsa_flags      int32
}

type sigaction1 = struct {
	F__sigaction_u __sigaction_u
	Fsa_mask       sigset_t
	Fsa_flags      int32
}

type sig_t = uintptr

type sigvec = struct {
	Fsv_handler uintptr
	Fsv_mask    int32
	Fsv_flags   int32
}

type sigstack = struct {
	Fss_sp      uintptr
	Fss_onstack int32
}

type timeval = struct {
	Ftv_sec  __darwin_time_t
	Ftv_usec __darwin_suseconds_t
}

type rlim_t = uint64

type rusage = struct {
	Fru_utime    timeval
	Fru_stime    timeval
	Fru_maxrss   int64
	Fru_ixrss    int64
	Fru_idrss    int64
	Fru_isrss    int64
	Fru_minflt   int64
	Fru_majflt   int64
	Fru_nswap    int64
	Fru_inblock  int64
	Fru_oublock  int64
	Fru_msgsnd   int64
	Fru_msgrcv   int64
	Fru_nsignals int64
	Fru_nvcsw    int64
	Fru_nivcsw   int64
}

type rusage_info_t = uintptr

type rusage_info_v0 = struct {
	Fri_uuid               [16]uint8_t
	Fri_user_time          uint64_t
	Fri_system_time        uint64_t
	Fri_pkg_idle_wkups     uint64_t
	Fri_interrupt_wkups    uint64_t
	Fri_pageins            uint64_t
	Fri_wired_size         uint64_t
	Fri_resident_size      uint64_t
	Fri_phys_footprint     uint64_t
	Fri_proc_start_abstime uint64_t
	Fri_proc_exit_abstime  uint64_t
}

type rusage_info_v1 = struct {
	Fri_uuid                  [16]uint8_t
	Fri_user_time             uint64_t
	Fri_system_time           uint64_t
	Fri_pkg_idle_wkups        uint64_t
	Fri_interrupt_wkups       uint64_t
	Fri_pageins               uint64_t
	Fri_wired_size            uint64_t
	Fri_resident_size         uint64_t
	Fri_phys_footprint        uint64_t
	Fri_proc_start_abstime    uint64_t
	Fri_proc_exit_abstime     uint64_t
	Fri_child_user_time       uint64_t
	Fri_child_system_time     uint64_t
	Fri_child_pkg_idle_wkups  uint64_t
	Fri_child_interrupt_wkups uint64_t
	Fri_child_pageins         uint64_t
	Fri_child_elapsed_abstime uint64_t
}

type rusage_info_v2 = struct {
	Fri_uuid                  [16]uint8_t
	Fri_user_time             uint64_t
	Fri_system_time           uint64_t
	Fri_pkg_idle_wkups        uint64_t
	Fri_interrupt_wkups       uint64_t
	Fri_pageins               uint64_t
	Fri_wired_size            uint64_t
	Fri_resident_size         uint64_t
	Fri_phys_footprint        uint64_t
	Fri_proc_start_abstime    uint64_t
	Fri_proc_exit_abstime     uint64_t
	Fri_child_user_time       uint64_t
	Fri_child_system_time     uint64_t
	Fri_child_pkg_idle_wkups  uint64_t
	Fri_child_interrupt_wkups uint64_t
	Fri_child_pageins         uint64_t
	Fri_child_elapsed_abstime uint64_t
	Fri_diskio_bytesread      uint64_t
	Fri_diskio_byteswritten   uint64_t
}

type rusage_info_v3 = struct {
	Fri_uuid                          [16]uint8_t
	Fri_user_time                     uint64_t
	Fri_system_time                   uint64_t
	Fri_pkg_idle_wkups                uint64_t
	Fri_interrupt_wkups               uint64_t
	Fri_pageins                       uint64_t
	Fri_wired_size                    uint64_t
	Fri_resident_size                 uint64_t
	Fri_phys_footprint                uint64_t
	Fri_proc_start_abstime            uint64_t
	Fri_proc_exit_abstime             uint64_t
	Fri_child_user_time               uint64_t
	Fri_child_system_time             uint64_t
	Fri_child_pkg_idle_wkups          uint64_t
	Fri_child_interrupt_wkups         uint64_t
	Fri_child_pageins                 uint64_t
	Fri_child_elapsed_abstime         uint64_t
	Fri_diskio_bytesread              uint64_t
	Fri_diskio_byteswritten           uint64_t
	Fri_cpu_time_qos_default          uint64_t
	Fri_cpu_time_qos_maintenance      uint64_t
	Fri_cpu_time_qos_background       uint64_t
	Fri_cpu_time_qos_utility          uint64_t
	Fri_cpu_time_qos_legacy           uint64_t
	Fri_cpu_time_qos_user_initiated   uint64_t
	Fri_cpu_time_qos_user_interactive uint64_t
	Fri_billed_system_time            uint64_t
	Fri_serviced_system_time          uint64_t
}

type rusage_info_v4 = struct {
	Fri_uuid                          [16]uint8_t
	Fri_user_time                     uint64_t
	Fri_system_time                   uint64_t
	Fri_pkg_idle_wkups                uint64_t
	Fri_interrupt_wkups               uint64_t
	Fri_pageins                       uint64_t
	Fri_wired_size                    uint64_t
	Fri_resident_size                 uint64_t
	Fri_phys_footprint                uint64_t
	Fri_proc_start_abstime            uint64_t
	Fri_proc_exit_abstime             uint64_t
	Fri_child_user_time               uint64_t
	Fri_child_system_time             uint64_t
	Fri_child_pkg_idle_wkups          uint64_t
	Fri_child_interrupt_wkups         uint64_t
	Fri_child_pageins                 uint64_t
	Fri_child_elapsed_abstime         uint64_t
	Fri_diskio_bytesread              uint64_t
	Fri_diskio_byteswritten           uint64_t
	Fri_cpu_time_qos_default          uint64_t
	Fri_cpu_time_qos_maintenance      uint64_t
	Fri_cpu_time_qos_background       uint64_t
	Fri_cpu_time_qos_utility          uint64_t
	Fri_cpu_time_qos_legacy           uint64_t
	Fri_cpu_time_qos_user_initiated   uint64_t
	Fri_cpu_time_qos_user_interactive uint64_t
	Fri_billed_system_time            uint64_t
	Fri_serviced_system_time          uint64_t
	Fri_logical_writes                uint64_t
	Fri_lifetime_max_phys_footprint   uint64_t
	Fri_instructions                  uint64_t
	Fri_cycles                        uint64_t
	Fri_billed_energy                 uint64_t
	Fri_serviced_energy               uint64_t
	Fri_interval_max_phys_footprint   uint64_t
	Fri_runnable_time                 uint64_t
}

type rusage_info_v5 = struct {
	Fri_uuid                          [16]uint8_t
	Fri_user_time                     uint64_t
	Fri_system_time                   uint64_t
	Fri_pkg_idle_wkups                uint64_t
	Fri_interrupt_wkups               uint64_t
	Fri_pageins                       uint64_t
	Fri_wired_size                    uint64_t
	Fri_resident_size                 uint64_t
	Fri_phys_footprint                uint64_t
	Fri_proc_start_abstime            uint64_t
	Fri_proc_exit_abstime             uint64_t
	Fri_child_user_time               uint64_t
	Fri_child_system_time             uint64_t
	Fri_child_pkg_idle_wkups          uint64_t
	Fri_child_interrupt_wkups         uint64_t
	Fri_child_pageins                 uint64_t
	Fri_child_elapsed_abstime         uint64_t
	Fri_diskio_bytesread              uint64_t
	Fri_diskio_byteswritten           uint64_t
	Fri_cpu_time_qos_default          uint64_t
	Fri_cpu_time_qos_maintenance      uint64_t
	Fri_cpu_time_qos_background       uint64_t
	Fri_cpu_time_qos_utility          uint64_t
	Fri_cpu_time_qos_legacy           uint64_t
	Fri_cpu_time_qos_user_initiated   uint64_t
	Fri_cpu_time_qos_user_interactive uint64_t
	Fri_billed_system_time            uint64_t
	Fri_serviced_system_time          uint64_t
	Fri_logical_writes                uint64_t
	Fri_lifetime_max_phys_footprint   uint64_t
	Fri_instructions                  uint64_t
	Fri_cycles                        uint64_t
	Fri_billed_energy                 uint64_t
	Fri_serviced_energy               uint64_t
	Fri_interval_max_phys_footprint   uint64_t
	Fri_runnable_time                 uint64_t
	Fri_flags                         uint64_t
}

type rusage_info_v6 = struct {
	Fri_uuid                          [16]uint8_t
	Fri_user_time                     uint64_t
	Fri_system_time                   uint64_t
	Fri_pkg_idle_wkups                uint64_t
	Fri_interrupt_wkups               uint64_t
	Fri_pageins                       uint64_t
	Fri_wired_size                    uint64_t
	Fri_resident_size                 uint64_t
	Fri_phys_footprint                uint64_t
	Fri_proc_start_abstime            uint64_t
	Fri_proc_exit_abstime             uint64_t
	Fri_child_user_time               uint64_t
	Fri_child_system_time             uint64_t
	Fri_child_pkg_idle_wkups          uint64_t
	Fri_child_interrupt_wkups         uint64_t
	Fri_child_pageins                 uint64_t
	Fri_child_elapsed_abstime         uint64_t
	Fri_diskio_bytesread              uint64_t
	Fri_diskio_byteswritten           uint64_t
	Fri_cpu_time_qos_default          uint64_t
	Fri_cpu_time_qos_maintenance      uint64_t
	Fri_cpu_time_qos_background       uint64_t
	Fri_cpu_time_qos_utility          uint64_t
	Fri_cpu_time_qos_legacy           uint64_t
	Fri_cpu_time_qos_user_initiated   uint64_t
	Fri_cpu_time_qos_user_interactive uint64_t
	Fri_billed_system_time            uint64_t
	Fri_serviced_system_time          uint64_t
	Fri_logical_writes                uint64_t
	Fri_lifetime_max_phys_footprint   uint64_t
	Fri_instructions                  uint64_t
	Fri_cycles                        uint64_t
	Fri_billed_energy                 uint64_t
	Fri_serviced_energy               uint64_t
	Fri_interval_max_phys_footprint   uint64_t
	Fri_runnable_time                 uint64_t
	Fri_flags                         uint64_t
	Fri_user_ptime                    uint64_t
	Fri_system_ptime                  uint64_t
	Fri_pinstructions                 uint64_t
	Fri_pcycles                       uint64_t
	Fri_energy_nj                     uint64_t
	Fri_penergy_nj                    uint64_t
	Fri_secure_time_in_system         uint64_t
	Fri_secure_ptime_in_system        uint64_t
	Fri_neural_footprint              uint64_t
	Fri_lifetime_max_neural_footprint uint64_t
	Fri_interval_max_neural_footprint uint64_t
	Fri_reserved                      [9]uint64_t
}

type rusage_info_current = struct {
	Fri_uuid                          [16]uint8_t
	Fri_user_time                     uint64_t
	Fri_system_time                   uint64_t
	Fri_pkg_idle_wkups                uint64_t
	Fri_interrupt_wkups               uint64_t
	Fri_pageins                       uint64_t
	Fri_wired_size                    uint64_t
	Fri_resident_size                 uint64_t
	Fri_phys_footprint                uint64_t
	Fri_proc_start_abstime            uint64_t
	Fri_proc_exit_abstime             uint64_t
	Fri_child_user_time               uint64_t
	Fri_child_system_time             uint64_t
	Fri_child_pkg_idle_wkups          uint64_t
	Fri_child_interrupt_wkups         uint64_t
	Fri_child_pageins                 uint64_t
	Fri_child_elapsed_abstime         uint64_t
	Fri_diskio_bytesread              uint64_t
	Fri_diskio_byteswritten           uint64_t
	Fri_cpu_time_qos_default          uint64_t
	Fri_cpu_time_qos_maintenance      uint64_t
	Fri_cpu_time_qos_background       uint64_t
	Fri_cpu_time_qos_utility          uint64_t
	Fri_cpu_time_qos_legacy           uint64_t
	Fri_cpu_time_qos_user_initiated   uint64_t
	Fri_cpu_time_qos_user_interactive uint64_t
	Fri_billed_system_time            uint64_t
	Fri_serviced_system_time          uint64_t
	Fri_logical_writes                uint64_t
	Fri_lifetime_max_phys_footprint   uint64_t
	Fri_instructions                  uint64_t
	Fri_cycles                        uint64_t
	Fri_billed_energy                 uint64_t
	Fri_serviced_energy               uint64_t
	Fri_interval_max_phys_footprint   uint64_t
	Fri_runnable_time                 uint64_t
	Fri_flags                         uint64_t
	Fri_user_ptime                    uint64_t
	Fri_system_ptime                  uint64_t
	Fri_pinstructions                 uint64_t
	Fri_pcycles                       uint64_t
	Fri_energy_nj                     uint64_t
	Fri_penergy_nj                    uint64_t
	Fri_secure_time_in_system         uint64_t
	Fri_secure_ptime_in_system        uint64_t
	Fri_neural_footprint              uint64_t
	Fri_lifetime_max_neural_footprint uint64_t
	Fri_interval_max_neural_footprint uint64_t
	Fri_reserved                      [9]uint64_t
}

type rlimit = struct {
	Frlim_cur rlim_t
	Frlim_max rlim_t
}

type proc_rlimit_control_wakeupmon = struct {
	Fwm_flags uint32_t
	Fwm_rate  int32_t
}

type wait = struct {
	Fw_T [0]struct {
		F__ccgo0 uint32
	}
	Fw_S [0]struct {
		F__ccgo0 uint32
	}
	Fw_status int32
}

type ct_rune_t = int32

type rune_t = int32

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

type malloc_type_id_t = uint64

type dev_t = int32

type mode_t = uint16

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

type wctrans_t = int32

type wint_t = int32

type wctype_t = uint32

type _RuneEntry = struct {
	F__min   __darwin_rune_t
	F__max   __darwin_rune_t
	F__map   __darwin_rune_t
	F__types uintptr
}

type _RuneRange = struct {
	F__nranges int32
	F__ranges  uintptr
}

type _RuneCharClass = struct {
	F__name [14]int8
	F__mask __uint32_t
}

type _RuneLocale = struct {
	F__magic        [8]int8
	F__encoding     [32]int8
	F__sgetrune     uintptr
	F__sputrune     uintptr
	F__invalid_rune __darwin_rune_t
	F__runetype     [256]__uint32_t
	F__maplower     [256]__darwin_rune_t
	F__mapupper     [256]__darwin_rune_t
	F__runetype_ext _RuneRange
	F__maplower_ext _RuneRange
	F__mapupper_ext _RuneRange
	F__variable     uintptr
	F__variable_len int32
	F__ncharclasses int32
	F__charclasses  uintptr
}

type TokenType = int32

const HELP_TEXT = 0

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

func tree_sitter_kconfig_external_scanner_create(tls *libc.TLS) (r uintptr) {
	return libc.UintptrFromInt32(0)
}

func tree_sitter_kconfig_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
}

func tree_sitter_kconfig_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	return uint32(0)
}

func tree_sitter_kconfig_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
}

func tree_sitter_kconfig_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var next_col, start_col uint32_t
	var v1, v5, v7, v8 int32
	var v3 __darwin_ct_rune_t
	var v4 uint64
	_, _, _, _, _, _, _, _ = next_col, start_col, v1, v3, v4, v5, v7, v8
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(HELP_TEXT))) != 0 {
		start_col = uint32(0)
		for {
			v3 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			v4 = uint64(0x00004000)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _9
		_9:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _6
		_6:
			v1 = v5
			goto _2
		_2:
			if !(v1 != 0) {
				break
			}
			switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
			case int32(' '):
				start_col = start_col + 1
			case int32('\t'):
				start_col = start_col + uint32(8)
				// Align col to next tab stop, ignore up to 7 leading spaces
				start_col = start_col - start_col%uint32(8)
			default:
				break
			}
			skip(tls, lexer)
		}
		goto scan_line
	scan_line:
		;
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
			advance(tls, lexer)
		}
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HELP_TEXT)
			return libc.BoolUint8(true1 != 0)
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		next_col = uint32(0)
		for {
			v3 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			v4 = uint64(0x00004000)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _18
		_18:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _15
		_15:
			v1 = v5
			goto _11
		_11:
			if !(v1 != 0) {
				break
			}
			switch (*TSLexer)(unsafe.Pointer(lexer)).Flookahead {
			case int32(' '):
				next_col = next_col + 1
			case int32('\t'):
				next_col = next_col + uint32(8)
				// Align col to next tab stop, ignore up to 7 leading spaces
				next_col = next_col - next_col%uint32(8)
			default:
				break
			}
			advance(tls, lexer)
		}
		if next_col >= start_col {
			goto scan_line
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(HELP_TEXT)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

type ts_symbol_identifiers = int32

const sym_symbol = 1
const anon_sym_mainmenu = 2
const anon_sym_config = 3
const anon_sym_configdefault = 4
const anon_sym_menuconfig = 5
const anon_sym_choice = 6
const anon_sym_endchoice = 7
const anon_sym_comment = 8
const anon_sym_menu = 9
const anon_sym_endmenu = 10
const anon_sym_if = 11
const anon_sym_endif = 12
const anon_sym_source = 13
const anon_sym_rsource = 14
const anon_sym_osource = 15
const anon_sym_orsource = 16
const anon_sym_EQ = 17
const anon_sym_COLON_EQ = 18
const anon_sym_PLUS_EQ = 19
const anon_sym_QMARK_EQ = 20
const anon_sym_COMMA = 21
const aux_sym_variable_token1 = 22
const anon_sym_bool = 23
const anon_sym_tristate = 24
const anon_sym_int = 25
const anon_sym_hex = 26
const anon_sym_string = 27
const aux_sym_type_definition_token1 = 28
const anon_sym_prompt = 29
const anon_sym_default = 30
const anon_sym_def_bool = 31
const anon_sym_def_tristate = 32
const anon_sym_def_int = 33
const anon_sym_def_hex = 34
const anon_sym_def_string = 35
const anon_sym_dependson = 36
const anon_sym_select = 37
const anon_sym_imply = 38
const anon_sym_visibleif = 39
const anon_sym_range = 40
const anon_sym_help = 41
const sym_optional = 42
const sym_modules = 43
const anon_sym_BANG = 44
const anon_sym_PIPE_PIPE = 45
const anon_sym_AMP_AMP = 46
const anon_sym_BANG_EQ = 47
const anon_sym_LT = 48
const anon_sym_GT = 49
const anon_sym_LT_EQ = 50
const anon_sym_GT_EQ = 51
const anon_sym_LPAREN = 52
const anon_sym_RPAREN = 53
const anon_sym_DOLLAR_LPAREN = 54
const sym_macro_content = 55
const anon_sym_DQUOTE = 56
const aux_sym_string_token1 = 57
const aux_sym_string_token2 = 58
const anon_sym_SQUOTE = 59
const aux_sym_string_token3 = 60
const sym_text = 61
const sym_comment = 62
const sym__help_text = 63
const sym_configuration = 64
const sym__entry = 65
const sym_mainmenu = 66
const sym_config = 67
const sym_configdefault = 68
const sym_menuconfig = 69
const sym_choice = 70
const sym_comment_entry = 71
const sym_menu = 72
const sym_if = 73
const sym_source = 74
const sym_variable = 75
const sym__config_option = 76
const sym_type_definition = 77
const sym_input_prompt = 78
const sym_default_value = 79
const sym_type_definition_default = 80
const sym_dependencies = 81
const sym_reverse_dependencies = 82
const sym_weak_reverse_dependencies = 83
const sym_limiting_menu_display = 84
const sym_numerical_ranges = 85
const sym_help_text = 86
const sym_conditional_clause = 87
const sym_expression = 88
const sym_unary_expression = 89
const sym_binary_expression = 90
const sym_parenthesized_expression = 91
const sym_macro_variable = 92
const sym_string = 93
const sym_name = 94
const aux_sym_configuration_repeat1 = 95
const aux_sym_config_repeat1 = 96
const aux_sym_configdefault_repeat1 = 97
const aux_sym_variable_repeat1 = 98
const aux_sym_macro_variable_repeat1 = 99
const aux_sym_string_repeat1 = 100
const aux_sym_string_repeat2 = 101
const aux_sym_name_repeat1 = 102

var ts_symbol_names = [103]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 11,
	3:   __ccgo_ts + 20,
	4:   __ccgo_ts + 27,
	5:   __ccgo_ts + 41,
	6:   __ccgo_ts + 52,
	7:   __ccgo_ts + 59,
	8:   __ccgo_ts + 69,
	9:   __ccgo_ts + 77,
	10:  __ccgo_ts + 82,
	11:  __ccgo_ts + 90,
	12:  __ccgo_ts + 93,
	13:  __ccgo_ts + 99,
	14:  __ccgo_ts + 106,
	15:  __ccgo_ts + 114,
	16:  __ccgo_ts + 122,
	17:  __ccgo_ts + 131,
	18:  __ccgo_ts + 133,
	19:  __ccgo_ts + 136,
	20:  __ccgo_ts + 139,
	21:  __ccgo_ts + 142,
	22:  __ccgo_ts + 144,
	23:  __ccgo_ts + 160,
	24:  __ccgo_ts + 165,
	25:  __ccgo_ts + 174,
	26:  __ccgo_ts + 178,
	27:  __ccgo_ts + 182,
	28:  __ccgo_ts + 189,
	29:  __ccgo_ts + 212,
	30:  __ccgo_ts + 219,
	31:  __ccgo_ts + 227,
	32:  __ccgo_ts + 236,
	33:  __ccgo_ts + 249,
	34:  __ccgo_ts + 257,
	35:  __ccgo_ts + 265,
	36:  __ccgo_ts + 276,
	37:  __ccgo_ts + 287,
	38:  __ccgo_ts + 294,
	39:  __ccgo_ts + 300,
	40:  __ccgo_ts + 311,
	41:  __ccgo_ts + 317,
	42:  __ccgo_ts + 322,
	43:  __ccgo_ts + 331,
	44:  __ccgo_ts + 339,
	45:  __ccgo_ts + 341,
	46:  __ccgo_ts + 344,
	47:  __ccgo_ts + 347,
	48:  __ccgo_ts + 350,
	49:  __ccgo_ts + 352,
	50:  __ccgo_ts + 354,
	51:  __ccgo_ts + 357,
	52:  __ccgo_ts + 360,
	53:  __ccgo_ts + 362,
	54:  __ccgo_ts + 364,
	55:  __ccgo_ts + 367,
	56:  __ccgo_ts + 381,
	57:  __ccgo_ts + 383,
	58:  __ccgo_ts + 398,
	59:  __ccgo_ts + 412,
	60:  __ccgo_ts + 383,
	61:  __ccgo_ts + 414,
	62:  __ccgo_ts + 69,
	63:  __ccgo_ts + 414,
	64:  __ccgo_ts + 419,
	65:  __ccgo_ts + 433,
	66:  __ccgo_ts + 11,
	67:  __ccgo_ts + 20,
	68:  __ccgo_ts + 27,
	69:  __ccgo_ts + 41,
	70:  __ccgo_ts + 52,
	71:  __ccgo_ts + 440,
	72:  __ccgo_ts + 77,
	73:  __ccgo_ts + 90,
	74:  __ccgo_ts + 99,
	75:  __ccgo_ts + 454,
	76:  __ccgo_ts + 463,
	77:  __ccgo_ts + 478,
	78:  __ccgo_ts + 494,
	79:  __ccgo_ts + 507,
	80:  __ccgo_ts + 521,
	81:  __ccgo_ts + 545,
	82:  __ccgo_ts + 558,
	83:  __ccgo_ts + 579,
	84:  __ccgo_ts + 605,
	85:  __ccgo_ts + 627,
	86:  __ccgo_ts + 644,
	87:  __ccgo_ts + 654,
	88:  __ccgo_ts + 673,
	89:  __ccgo_ts + 684,
	90:  __ccgo_ts + 701,
	91:  __ccgo_ts + 719,
	92:  __ccgo_ts + 744,
	93:  __ccgo_ts + 182,
	94:  __ccgo_ts + 759,
	95:  __ccgo_ts + 764,
	96:  __ccgo_ts + 786,
	97:  __ccgo_ts + 801,
	98:  __ccgo_ts + 823,
	99:  __ccgo_ts + 840,
	100: __ccgo_ts + 863,
	101: __ccgo_ts + 878,
	102: __ccgo_ts + 893,
}

var ts_symbol_map = [103]TSSymbol{
	1:   uint16(sym_symbol),
	2:   uint16(anon_sym_mainmenu),
	3:   uint16(anon_sym_config),
	4:   uint16(anon_sym_configdefault),
	5:   uint16(anon_sym_menuconfig),
	6:   uint16(anon_sym_choice),
	7:   uint16(anon_sym_endchoice),
	8:   uint16(anon_sym_comment),
	9:   uint16(anon_sym_menu),
	10:  uint16(anon_sym_endmenu),
	11:  uint16(anon_sym_if),
	12:  uint16(anon_sym_endif),
	13:  uint16(anon_sym_source),
	14:  uint16(anon_sym_rsource),
	15:  uint16(anon_sym_osource),
	16:  uint16(anon_sym_orsource),
	17:  uint16(anon_sym_EQ),
	18:  uint16(anon_sym_COLON_EQ),
	19:  uint16(anon_sym_PLUS_EQ),
	20:  uint16(anon_sym_QMARK_EQ),
	21:  uint16(anon_sym_COMMA),
	22:  uint16(aux_sym_variable_token1),
	23:  uint16(anon_sym_bool),
	24:  uint16(anon_sym_tristate),
	25:  uint16(anon_sym_int),
	26:  uint16(anon_sym_hex),
	27:  uint16(anon_sym_string),
	28:  uint16(aux_sym_type_definition_token1),
	29:  uint16(anon_sym_prompt),
	30:  uint16(anon_sym_default),
	31:  uint16(anon_sym_def_bool),
	32:  uint16(anon_sym_def_tristate),
	33:  uint16(anon_sym_def_int),
	34:  uint16(anon_sym_def_hex),
	35:  uint16(anon_sym_def_string),
	36:  uint16(anon_sym_dependson),
	37:  uint16(anon_sym_select),
	38:  uint16(anon_sym_imply),
	39:  uint16(anon_sym_visibleif),
	40:  uint16(anon_sym_range),
	41:  uint16(anon_sym_help),
	42:  uint16(sym_optional),
	43:  uint16(sym_modules),
	44:  uint16(anon_sym_BANG),
	45:  uint16(anon_sym_PIPE_PIPE),
	46:  uint16(anon_sym_AMP_AMP),
	47:  uint16(anon_sym_BANG_EQ),
	48:  uint16(anon_sym_LT),
	49:  uint16(anon_sym_GT),
	50:  uint16(anon_sym_LT_EQ),
	51:  uint16(anon_sym_GT_EQ),
	52:  uint16(anon_sym_LPAREN),
	53:  uint16(anon_sym_RPAREN),
	54:  uint16(anon_sym_DOLLAR_LPAREN),
	55:  uint16(sym_macro_content),
	56:  uint16(anon_sym_DQUOTE),
	57:  uint16(aux_sym_string_token1),
	58:  uint16(aux_sym_string_token2),
	59:  uint16(anon_sym_SQUOTE),
	60:  uint16(aux_sym_string_token1),
	61:  uint16(sym_text),
	62:  uint16(sym_comment),
	63:  uint16(sym_text),
	64:  uint16(sym_configuration),
	65:  uint16(sym__entry),
	66:  uint16(sym_mainmenu),
	67:  uint16(sym_config),
	68:  uint16(sym_configdefault),
	69:  uint16(sym_menuconfig),
	70:  uint16(sym_choice),
	71:  uint16(sym_comment_entry),
	72:  uint16(sym_menu),
	73:  uint16(sym_if),
	74:  uint16(sym_source),
	75:  uint16(sym_variable),
	76:  uint16(sym__config_option),
	77:  uint16(sym_type_definition),
	78:  uint16(sym_input_prompt),
	79:  uint16(sym_default_value),
	80:  uint16(sym_type_definition_default),
	81:  uint16(sym_dependencies),
	82:  uint16(sym_reverse_dependencies),
	83:  uint16(sym_weak_reverse_dependencies),
	84:  uint16(sym_limiting_menu_display),
	85:  uint16(sym_numerical_ranges),
	86:  uint16(sym_help_text),
	87:  uint16(sym_conditional_clause),
	88:  uint16(sym_expression),
	89:  uint16(sym_unary_expression),
	90:  uint16(sym_binary_expression),
	91:  uint16(sym_parenthesized_expression),
	92:  uint16(sym_macro_variable),
	93:  uint16(sym_string),
	94:  uint16(sym_name),
	95:  uint16(aux_sym_configuration_repeat1),
	96:  uint16(aux_sym_config_repeat1),
	97:  uint16(aux_sym_configdefault_repeat1),
	98:  uint16(aux_sym_variable_repeat1),
	99:  uint16(aux_sym_macro_variable_repeat1),
	100: uint16(aux_sym_string_repeat1),
	101: uint16(aux_sym_string_repeat2),
	102: uint16(aux_sym_name_repeat1),
}

var ts_symbol_metadata = [103]TSSymbolMetadata{
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
	22: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	43: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	52: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	54: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	55: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	58: {},
	59: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
	95:  {},
	96:  {},
	97:  {},
	98:  {},
	99:  {},
	100: {},
	101: {},
	102: {},
}

type ts_field_identifiers = int32

const field_condition = 1
const field_left = 2
const field_name = 3
const field_operator = 4
const field_right = 5

var ts_field_names = [6]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 906,
	2: __ccgo_ts + 916,
	3: __ccgo_ts + 759,
	4: __ccgo_ts + 921,
	5: __ccgo_ts + 930,
}

var ts_field_map_slices = [9]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
	6: {
		Findex:  uint16(4),
		Flength: uint16(3),
	},
	7: {
		Findex:  uint16(7),
		Flength: uint16(2),
	},
	8: {
		Findex:  uint16(9),
		Flength: uint16(2),
	},
}

var ts_field_map_entries = [11]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_condition),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id: uint16(field_left),
	},
	3: {
		Ffield_id: uint16(field_right),
	},
	4: {
		Ffield_id: uint16(field_left),
	},
	5: {
		Ffield_id:    uint16(field_operator),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
	},
	7: {
		Ffield_id: uint16(field_left),
	},
	8: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	9: {
		Ffield_id:  uint16(field_right),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	10: {
		Ffield_id:    uint16(field_right),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [9][5]TSSymbol{
	0: {},
	2: {
		0: uint16(sym_text),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(sym_string),
	1: uint16(2),
	2: uint16(sym_string),
	3: uint16(sym_text),
}

var ts_primary_state_ids = [383]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(2),
	7:   uint16(3),
	8:   uint16(4),
	9:   uint16(5),
	10:  uint16(10),
	11:  uint16(11),
	12:  uint16(12),
	13:  uint16(13),
	14:  uint16(14),
	15:  uint16(15),
	16:  uint16(11),
	17:  uint16(17),
	18:  uint16(10),
	19:  uint16(19),
	20:  uint16(20),
	21:  uint16(21),
	22:  uint16(22),
	23:  uint16(23),
	24:  uint16(12),
	25:  uint16(14),
	26:  uint16(15),
	27:  uint16(17),
	28:  uint16(13),
	29:  uint16(22),
	30:  uint16(19),
	31:  uint16(21),
	32:  uint16(23),
	33:  uint16(20),
	34:  uint16(34),
	35:  uint16(35),
	36:  uint16(36),
	37:  uint16(37),
	38:  uint16(38),
	39:  uint16(39),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(38),
	43:  uint16(41),
	44:  uint16(37),
	45:  uint16(34),
	46:  uint16(40),
	47:  uint16(35),
	48:  uint16(36),
	49:  uint16(39),
	50:  uint16(50),
	51:  uint16(50),
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
	69:  uint16(67),
	70:  uint16(61),
	71:  uint16(57),
	72:  uint16(65),
	73:  uint16(66),
	74:  uint16(53),
	75:  uint16(64),
	76:  uint16(58),
	77:  uint16(62),
	78:  uint16(55),
	79:  uint16(56),
	80:  uint16(59),
	81:  uint16(52),
	82:  uint16(68),
	83:  uint16(54),
	84:  uint16(63),
	85:  uint16(85),
	86:  uint16(85),
	87:  uint16(60),
	88:  uint16(88),
	89:  uint16(88),
	90:  uint16(90),
	91:  uint16(90),
	92:  uint16(92),
	93:  uint16(92),
	94:  uint16(11),
	95:  uint16(10),
	96:  uint16(19),
	97:  uint16(22),
	98:  uint16(98),
	99:  uint16(20),
	100: uint16(23),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(101),
	106: uint16(98),
	107: uint16(107),
	108: uint16(102),
	109: uint16(109),
	110: uint16(103),
	111: uint16(109),
	112: uint16(107),
	113: uint16(21),
	114: uint16(35),
	115: uint16(34),
	116: uint16(38),
	117: uint16(37),
	118: uint16(40),
	119: uint16(36),
	120: uint16(39),
	121: uint16(121),
	122: uint16(122),
	123: uint16(11),
	124: uint16(10),
	125: uint16(125),
	126: uint16(125),
	127: uint16(122),
	128: uint16(121),
	129: uint16(67),
	130: uint16(11),
	131: uint16(57),
	132: uint16(132),
	133: uint16(10),
	134: uint16(132),
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
	151: uint16(19),
	152: uint16(22),
	153: uint16(37),
	154: uint16(34),
	155: uint16(38),
	156: uint16(35),
	157: uint16(21),
	158: uint16(40),
	159: uint16(36),
	160: uint16(39),
	161: uint16(23),
	162: uint16(20),
	163: uint16(57),
	164: uint16(67),
	165: uint16(165),
	166: uint16(166),
	167: uint16(167),
	168: uint16(168),
	169: uint16(169),
	170: uint16(142),
	171: uint16(171),
	172: uint16(19),
	173: uint16(173),
	174: uint16(174),
	175: uint16(175),
	176: uint16(168),
	177: uint16(169),
	178: uint16(171),
	179: uint16(173),
	180: uint16(166),
	181: uint16(140),
	182: uint16(182),
	183: uint16(139),
	184: uint16(168),
	185: uint16(169),
	186: uint16(171),
	187: uint16(173),
	188: uint16(141),
	189: uint16(168),
	190: uint16(143),
	191: uint16(169),
	192: uint16(144),
	193: uint16(145),
	194: uint16(168),
	195: uint16(169),
	196: uint16(171),
	197: uint16(173),
	198: uint16(22),
	199: uint16(171),
	200: uint16(173),
	201: uint16(201),
	202: uint16(22),
	203: uint16(23),
	204: uint16(168),
	205: uint16(171),
	206: uint16(173),
	207: uint16(19),
	208: uint16(20),
	209: uint16(137),
	210: uint16(146),
	211: uint16(138),
	212: uint16(147),
	213: uint16(148),
	214: uint16(182),
	215: uint16(175),
	216: uint16(216),
	217: uint16(136),
	218: uint16(150),
	219: uint16(175),
	220: uint16(175),
	221: uint16(175),
	222: uint16(175),
	223: uint16(174),
	224: uint16(201),
	225: uint16(167),
	226: uint16(165),
	227: uint16(21),
	228: uint16(166),
	229: uint16(201),
	230: uint16(166),
	231: uint16(201),
	232: uint16(166),
	233: uint16(169),
	234: uint16(21),
	235: uint16(235),
	236: uint16(235),
	237: uint16(237),
	238: uint16(237),
	239: uint16(235),
	240: uint16(235),
	241: uint16(36),
	242: uint16(39),
	243: uint16(40),
	244: uint16(34),
	245: uint16(38),
	246: uint16(37),
	247: uint16(35),
	248: uint16(248),
	249: uint16(40),
	250: uint16(248),
	251: uint16(248),
	252: uint16(248),
	253: uint16(39),
	254: uint16(36),
	255: uint16(255),
	256: uint16(255),
	257: uint16(257),
	258: uint16(248),
	259: uint16(259),
	260: uint16(260),
	261: uint16(261),
	262: uint16(262),
	263: uint16(259),
	264: uint16(262),
	265: uint16(260),
	266: uint16(266),
	267: uint16(267),
	268: uint16(259),
	269: uint16(269),
	270: uint16(262),
	271: uint16(259),
	272: uint16(262),
	273: uint16(259),
	274: uint16(262),
	275: uint16(259),
	276: uint16(262),
	277: uint16(259),
	278: uint16(262),
	279: uint16(259),
	280: uint16(262),
	281: uint16(281),
	282: uint16(266),
	283: uint16(267),
	284: uint16(269),
	285: uint16(281),
	286: uint16(286),
	287: uint16(287),
	288: uint16(288),
	289: uint16(289),
	290: uint16(288),
	291: uint16(291),
	292: uint16(287),
	293: uint16(291),
	294: uint16(289),
	295: uint16(288),
	296: uint16(291),
	297: uint16(291),
	298: uint16(287),
	299: uint16(289),
	300: uint16(288),
	301: uint16(291),
	302: uint16(287),
	303: uint16(291),
	304: uint16(289),
	305: uint16(288),
	306: uint16(287),
	307: uint16(287),
	308: uint16(289),
	309: uint16(289),
	310: uint16(288),
	311: uint16(288),
	312: uint16(289),
	313: uint16(287),
	314: uint16(314),
	315: uint16(291),
	316: uint16(23),
	317: uint16(19),
	318: uint16(20),
	319: uint16(319),
	320: uint16(22),
	321: uint16(321),
	322: uint16(321),
	323: uint16(23),
	324: uint16(20),
	325: uint16(23),
	326: uint16(20),
	327: uint16(327),
	328: uint16(328),
	329: uint16(329),
	330: uint16(330),
	331: uint16(331),
	332: uint16(332),
	333: uint16(333),
	334: uint16(329),
	335: uint16(335),
	336: uint16(333),
	337: uint16(337),
	338: uint16(332),
	339: uint16(327),
	340: uint16(330),
	341: uint16(341),
	342: uint16(328),
	343: uint16(331),
	344: uint16(344),
	345: uint16(345),
	346: uint16(345),
	347: uint16(335),
	348: uint16(330),
	349: uint16(349),
	350: uint16(349),
	351: uint16(345),
	352: uint16(341),
	353: uint16(344),
	354: uint16(337),
	355: uint16(56),
	356: uint16(66),
	357: uint16(357),
	358: uint16(358),
	359: uint16(359),
	360: uint16(360),
	361: uint16(359),
	362: uint16(362),
	363: uint16(363),
	364: uint16(364),
	365: uint16(365),
	366: uint16(358),
	367: uint16(367),
	368: uint16(363),
	369: uint16(357),
	370: uint16(370),
	371: uint16(364),
	372: uint16(372),
	373: uint16(358),
	374: uint16(372),
	375: uint16(365),
	376: uint16(376),
	377: uint16(364),
	378: uint16(376),
	379: uint16(358),
	380: uint16(362),
	381: uint16(360),
	382: uint16(370),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, lookahead, result, skip
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
	switch libc.Int32FromUint16(state) {
	case 0:
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(26)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(3):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(26)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(4):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\f') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(110)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(7):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(26)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('\n') {
			state = uint16(85)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(84)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\n') {
			state = uint16(86)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(84)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('"') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(56)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(64)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(69)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('"') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(8)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(80)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('#') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(92)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(9)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(91)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('&') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('(') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('(') {
			state = uint16(57)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('(') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(73)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32(')') {
			state = uint16(73)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('=') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('=') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('=') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('=') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('f') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('i') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('n') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('o') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('|') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead != 0 && lookahead != int32('(') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(28):
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(29):
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(30):
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(31):
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(32):
		if eof != 0 {
			state = uint16(33)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('|') {
			state = uint16(26)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(34):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(35):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(36):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_variable_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_variable_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_type_definition_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_dependson)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_visibleif)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('=') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(83)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(94)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(114)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') {
			state = uint16(116)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(62)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') {
			state = uint16(116)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(62)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(65)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && (lookahead < int32('\'') || int32(')') < lookahead) {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(69)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && (lookahead < int32('\'') || int32(')') < lookahead) {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('"') || int32('\'') <= lookahead && lookahead <= int32(')') {
			state = uint16(117)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('"') || int32('\'') <= lookahead && lookahead <= int32(')') {
			state = uint16(117)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(68)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(61)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(67)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(64)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(69)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('"') || int32('$') < lookahead) && (lookahead < int32('\'') || int32(')') < lookahead) {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') {
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && (lookahead < int32('\'') || int32(')') < lookahead) {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(71)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && (lookahead < int32('\'') || int32(')') < lookahead) {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') {
			state = uint16(17)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_macro_content)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(74)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(79)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\\') {
			state = uint16(117)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(82)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(79)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(81)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(8)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(80)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('"') || int32('$') < lookahead) {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(83)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(58)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(83)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(83)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(8)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\\') {
			state = uint16(9)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('\'') || lookahead == int32('\\') {
			state = uint16(117)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(93)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(90)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(92)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(9)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(91)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(59)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(94)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_token3)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('$') {
			state = uint16(94)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(' ') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(' ') {
			state = uint16(23)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(107)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(96)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(101)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(98)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(100)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_symbol)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(110)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\r') {
			state = uint16(113)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(60)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(70)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(66)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(73)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(117)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(66)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(117)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [32]uint16_t{
	0:  uint16('!'),
	1:  uint16(45),
	2:  uint16('"'),
	3:  uint16(76),
	4:  uint16('#'),
	5:  uint16(117),
	6:  uint16('$'),
	7:  uint16(14),
	8:  uint16('&'),
	9:  uint16(13),
	10: uint16('\''),
	11: uint16(87),
	12: uint16('('),
	13: uint16(54),
	14: uint16(')'),
	15: uint16(56),
	16: uint16('+'),
	17: uint16(19),
	18: uint16(','),
	19: uint16(38),
	20: uint16('-'),
	21: uint16(109),
	22: uint16(':'),
	23: uint16(20),
	24: uint16('<'),
	25: uint16(50),
	26: uint16('='),
	27: uint16(34),
	28: uint16('>'),
	29: uint16(51),
	30: uint16('?'),
	31: uint16(21),
}

var map_token1 = [28]uint16_t{
	0:  uint16('\n'),
	1:  uint16(40),
	2:  uint16('\r'),
	3:  uint16(3),
	4:  uint16('!'),
	5:  uint16(45),
	6:  uint16('"'),
	7:  uint16(76),
	8:  uint16('#'),
	9:  uint16(117),
	10: uint16('$'),
	11: uint16(14),
	12: uint16('&'),
	13: uint16(13),
	14: uint16('\''),
	15: uint16(87),
	16: uint16('('),
	17: uint16(54),
	18: uint16(','),
	19: uint16(38),
	20: uint16('-'),
	21: uint16(109),
	22: uint16('<'),
	23: uint16(50),
	24: uint16('='),
	25: uint16(34),
	26: uint16('>'),
	27: uint16(51),
}

var map_token2 = [22]uint16_t{
	0:  uint16('\n'),
	1:  uint16(41),
	2:  uint16('\r'),
	3:  uint16(4),
	4:  uint16('!'),
	5:  uint16(46),
	6:  uint16('"'),
	7:  uint16(77),
	8:  uint16('#'),
	9:  uint16(113),
	10: uint16('$'),
	11: uint16(112),
	12: uint16('\''),
	13: uint16(88),
	14: uint16('('),
	15: uint16(55),
	16: uint16(','),
	17: uint16(39),
	18: uint16('-'),
	19: uint16(110),
	20: uint16('\\'),
	21: uint16(111),
}

var map_token3 = [22]uint16_t{
	0:  uint16('\n'),
	1:  uint16(42),
	2:  uint16('!'),
	3:  uint16(18),
	4:  uint16('"'),
	5:  uint16(76),
	6:  uint16('#'),
	7:  uint16(117),
	8:  uint16('$'),
	9:  uint16(14),
	10: uint16('&'),
	11: uint16(13),
	12: uint16('\''),
	13: uint16(87),
	14: uint16('-'),
	15: uint16(109),
	16: uint16('<'),
	17: uint16(50),
	18: uint16('='),
	19: uint16(34),
	20: uint16('>'),
	21: uint16(51),
}

var map_token4 = [24]uint16_t{
	0:  uint16('!'),
	1:  uint16(45),
	2:  uint16('"'),
	3:  uint16(76),
	4:  uint16('#'),
	5:  uint16(117),
	6:  uint16('$'),
	7:  uint16(14),
	8:  uint16('&'),
	9:  uint16(13),
	10: uint16('\''),
	11: uint16(87),
	12: uint16('('),
	13: uint16(54),
	14: uint16(')'),
	15: uint16(56),
	16: uint16('-'),
	17: uint16(109),
	18: uint16('<'),
	19: uint16(50),
	20: uint16('='),
	21: uint16(34),
	22: uint16('>'),
	23: uint16(51),
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
	switch libc.Int32FromUint16(state) {
	case 0:
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(3)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(13)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('o') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('h') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('e') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('n') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('e') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('f') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('a') {
			state = uint16(24)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('p') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32('s') {
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
		if lookahead == int32('a') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('e') {
			state = uint16(33)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('r') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('o') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('o') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('m') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('f') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('d') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('l') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_if)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(22):
		if lookahead == int32('p') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('t') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('i') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('n') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('d') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('t') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('s') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('o') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('o') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('n') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('o') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('l') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('u') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('r') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('i') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('l') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('i') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('m') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('f') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('_') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('c') {
			state = uint16(66)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('p') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_hex)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		if lookahead == int32('l') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_int)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		if lookahead == int32('n') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('u') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('u') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('i') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('o') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('u') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('m') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('g') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('u') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('e') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('r') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('i') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('s') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		if lookahead == int32('c') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('e') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('i') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('b') {
			state = uint16(87)
			goto next_state
		}
		if lookahead == int32('h') {
			state = uint16(88)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(89)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(90)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('u') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('h') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('f') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('e') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_help)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		if lookahead == int32('y') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('m') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_menu)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('l') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('o') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('u') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('r') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('p') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('e') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('r') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('c') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('c') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('n') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('t') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('e') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('n') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('g') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('o') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('e') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('n') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('t') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('r') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('l') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('o') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_endif)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		if lookahead == int32('n') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_imply)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		if lookahead == int32('e') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('o') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('e') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('n') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('r') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('c') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('t') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_range)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		if lookahead == int32('c') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('t') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('e') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('g') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('a') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_choice)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(111):
		if lookahead == int32('t') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_config)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('o') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('x') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('t') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('r') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('i') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('t') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('i') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('u') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('n') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('n') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('s') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead == int32('a') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('c') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('e') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_prompt)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(128):
		if lookahead == int32('e') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_select)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_source)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		if lookahead == int32('t') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		if lookahead == int32('e') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('l') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_def_hex)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_def_int)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(138):
		if lookahead == int32('i') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(139):
		if lookahead == int32('s') {
			state = uint16(154)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_default)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(141):
		if lookahead == int32('c') {
			state = uint16(155)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_endmenu)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(143):
		if lookahead == int32('u') {
			state = uint16(156)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('f') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_modules)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(146):
		if lookahead == int32('l') {
			state = uint16(158)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead == int32('e') {
			state = uint16(159)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_osource)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rsource)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(150):
		if lookahead == int32('e') {
			state = uint16(160)
			goto next_state
		}
		return result
	case int32(151):
		if lookahead == int32('f') {
			state = uint16(161)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_def_bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(153):
		if lookahead == int32('n') {
			state = uint16(162)
			goto next_state
		}
		return result
	case int32(154):
		if lookahead == int32('t') {
			state = uint16(163)
			goto next_state
		}
		return result
	case int32(155):
		if lookahead == int32('e') {
			state = uint16(164)
			goto next_state
		}
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_mainmenu)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(157):
		if lookahead == int32('i') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_optional)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_orsource)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_tristate)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(161):
		if lookahead == int32('a') {
			state = uint16(166)
			goto next_state
		}
		return result
	case int32(162):
		if lookahead == int32('g') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(163):
		if lookahead == int32('a') {
			state = uint16(168)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_endchoice)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(165):
		if lookahead == int32('g') {
			state = uint16(169)
			goto next_state
		}
		return result
	case int32(166):
		if lookahead == int32('u') {
			state = uint16(170)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_def_string)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		if lookahead == int32('t') {
			state = uint16(171)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_menuconfig)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(170):
		if lookahead == int32('l') {
			state = uint16(172)
			goto next_state
		}
		return result
	case int32(171):
		if lookahead == int32('e') {
			state = uint16(173)
			goto next_state
		}
		return result
	case int32(172):
		if lookahead == int32('t') {
			state = uint16(174)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_def_tristate)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_configdefault)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var ts_lex_modes = [383]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(32),
	},
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
	65: {},
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
	77: {},
	78: {},
	79: {},
	80: {},
	81: {},
	82: {},
	83: {},
	84: {},
	85: {
		Flex_state: uint16(32),
	},
	86: {
		Flex_state: uint16(32),
	},
	87: {},
	88: {},
	89: {},
	90: {},
	91: {},
	92: {},
	93: {},
	94: {
		Flex_state: uint16(32),
	},
	95: {
		Flex_state: uint16(32),
	},
	96: {
		Flex_state: uint16(32),
	},
	97: {
		Flex_state: uint16(32),
	},
	98: {
		Flex_state: uint16(32),
	},
	99: {
		Flex_state: uint16(32),
	},
	100: {
		Flex_state: uint16(32),
	},
	101: {
		Flex_state: uint16(32),
	},
	102: {
		Flex_state: uint16(32),
	},
	103: {
		Flex_state: uint16(32),
	},
	104: {
		Flex_state: uint16(32),
	},
	105: {
		Flex_state: uint16(32),
	},
	106: {
		Flex_state: uint16(32),
	},
	107: {
		Flex_state: uint16(32),
	},
	108: {
		Flex_state: uint16(32),
	},
	109: {
		Flex_state: uint16(32),
	},
	110: {
		Flex_state: uint16(32),
	},
	111: {
		Flex_state: uint16(32),
	},
	112: {
		Flex_state: uint16(32),
	},
	113: {
		Flex_state: uint16(32),
	},
	114: {
		Flex_state: uint16(32),
	},
	115: {
		Flex_state: uint16(32),
	},
	116: {
		Flex_state: uint16(32),
	},
	117: {
		Flex_state: uint16(32),
	},
	118: {
		Flex_state: uint16(32),
	},
	119: {
		Flex_state: uint16(32),
	},
	120: {
		Flex_state: uint16(32),
	},
	121: {
		Flex_state: uint16(32),
	},
	122: {
		Flex_state: uint16(32),
	},
	123: {
		Flex_state: uint16(3),
	},
	124: {
		Flex_state: uint16(3),
	},
	125: {
		Flex_state: uint16(4),
	},
	126: {
		Flex_state: uint16(4),
	},
	127: {
		Flex_state: uint16(32),
	},
	128: {
		Flex_state: uint16(32),
	},
	129: {
		Flex_state: uint16(32),
	},
	130: {
		Flex_state: uint16(7),
	},
	131: {
		Flex_state: uint16(32),
	},
	132: {
		Flex_state: uint16(3),
	},
	133: {
		Flex_state: uint16(7),
	},
	134: {
		Flex_state: uint16(3),
	},
	135: {
		Flex_state: uint16(3),
	},
	136: {
		Flex_state: uint16(32),
	},
	137: {
		Flex_state: uint16(32),
	},
	138: {
		Flex_state: uint16(32),
	},
	139: {
		Flex_state: uint16(32),
	},
	140: {
		Flex_state: uint16(32),
	},
	141: {
		Flex_state: uint16(32),
	},
	142: {
		Flex_state: uint16(32),
	},
	143: {
		Flex_state: uint16(32),
	},
	144: {
		Flex_state: uint16(32),
	},
	145: {
		Flex_state: uint16(32),
	},
	146: {
		Flex_state: uint16(32),
	},
	147: {
		Flex_state: uint16(32),
	},
	148: {
		Flex_state: uint16(32),
	},
	149: {
		Flex_state: uint16(3),
	},
	150: {
		Flex_state: uint16(32),
	},
	151: {
		Flex_state: uint16(3),
	},
	152: {
		Flex_state: uint16(3),
	},
	153: {
		Flex_state: uint16(3),
	},
	154: {
		Flex_state: uint16(3),
	},
	155: {
		Flex_state: uint16(3),
	},
	156: {
		Flex_state: uint16(3),
	},
	157: {
		Flex_state: uint16(3),
	},
	158: {
		Flex_state: uint16(3),
	},
	159: {
		Flex_state: uint16(3),
	},
	160: {
		Flex_state: uint16(3),
	},
	161: {
		Flex_state: uint16(3),
	},
	162: {
		Flex_state: uint16(3),
	},
	163: {
		Flex_state: uint16(32),
	},
	164: {
		Flex_state: uint16(32),
	},
	165: {
		Flex_state: uint16(32),
	},
	166: {
		Flex_state: uint16(32),
	},
	167: {
		Flex_state: uint16(32),
	},
	168: {
		Flex_state: uint16(32),
	},
	169: {
		Flex_state: uint16(32),
	},
	170: {
		Flex_state: uint16(32),
	},
	171: {
		Flex_state: uint16(32),
	},
	172: {
		Flex_state: uint16(32),
	},
	173: {
		Flex_state: uint16(32),
	},
	174: {
		Flex_state: uint16(32),
	},
	175: {
		Flex_state: uint16(32),
	},
	176: {
		Flex_state: uint16(32),
	},
	177: {
		Flex_state: uint16(32),
	},
	178: {
		Flex_state: uint16(32),
	},
	179: {
		Flex_state: uint16(32),
	},
	180: {
		Flex_state: uint16(32),
	},
	181: {
		Flex_state: uint16(32),
	},
	182: {
		Flex_state: uint16(32),
	},
	183: {
		Flex_state: uint16(32),
	},
	184: {
		Flex_state: uint16(32),
	},
	185: {
		Flex_state: uint16(32),
	},
	186: {
		Flex_state: uint16(32),
	},
	187: {
		Flex_state: uint16(32),
	},
	188: {
		Flex_state: uint16(32),
	},
	189: {
		Flex_state: uint16(32),
	},
	190: {
		Flex_state: uint16(32),
	},
	191: {
		Flex_state: uint16(32),
	},
	192: {
		Flex_state: uint16(32),
	},
	193: {
		Flex_state: uint16(32),
	},
	194: {
		Flex_state: uint16(32),
	},
	195: {
		Flex_state: uint16(32),
	},
	196: {
		Flex_state: uint16(32),
	},
	197: {
		Flex_state: uint16(32),
	},
	198: {
		Flex_state: uint16(32),
	},
	199: {
		Flex_state: uint16(32),
	},
	200: {
		Flex_state: uint16(32),
	},
	201: {
		Flex_state: uint16(32),
	},
	202: {
		Flex_state: uint16(7),
	},
	203: {
		Flex_state: uint16(7),
	},
	204: {
		Flex_state: uint16(32),
	},
	205: {
		Flex_state: uint16(32),
	},
	206: {
		Flex_state: uint16(32),
	},
	207: {
		Flex_state: uint16(7),
	},
	208: {
		Flex_state: uint16(7),
	},
	209: {
		Flex_state: uint16(32),
	},
	210: {
		Flex_state: uint16(32),
	},
	211: {
		Flex_state: uint16(32),
	},
	212: {
		Flex_state: uint16(32),
	},
	213: {
		Flex_state: uint16(32),
	},
	214: {
		Flex_state: uint16(32),
	},
	215: {
		Flex_state: uint16(32),
	},
	216: {
		Flex_state: uint16(32),
	},
	217: {
		Flex_state: uint16(32),
	},
	218: {
		Flex_state: uint16(32),
	},
	219: {
		Flex_state: uint16(32),
	},
	220: {
		Flex_state: uint16(32),
	},
	221: {
		Flex_state: uint16(32),
	},
	222: {
		Flex_state: uint16(32),
	},
	223: {
		Flex_state: uint16(32),
	},
	224: {
		Flex_state: uint16(32),
	},
	225: {
		Flex_state: uint16(32),
	},
	226: {
		Flex_state: uint16(32),
	},
	227: {
		Flex_state: uint16(7),
	},
	228: {
		Flex_state: uint16(32),
	},
	229: {
		Flex_state: uint16(32),
	},
	230: {
		Flex_state: uint16(32),
	},
	231: {
		Flex_state: uint16(32),
	},
	232: {
		Flex_state: uint16(32),
	},
	233: {
		Flex_state: uint16(32),
	},
	234: {
		Flex_state: uint16(32),
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
	238: {
		Flex_state: uint16(7),
	},
	239: {
		Flex_state: uint16(7),
	},
	240: {
		Flex_state: uint16(7),
	},
	241: {
		Flex_state: uint16(7),
	},
	242: {
		Flex_state: uint16(7),
	},
	243: {
		Flex_state: uint16(7),
	},
	244: {
		Flex_state: uint16(7),
	},
	245: {
		Flex_state: uint16(7),
	},
	246: {
		Flex_state: uint16(7),
	},
	247: {
		Flex_state: uint16(7),
	},
	248: {},
	249: {},
	250: {},
	251: {},
	252: {},
	253: {},
	254: {},
	255: {
		Flex_state: uint16(7),
	},
	256: {
		Flex_state: uint16(7),
	},
	257: {
		Flex_state: uint16(7),
	},
	258: {},
	259: {
		Flex_state: uint16(10),
	},
	260: {
		Flex_state: uint16(7),
	},
	261: {
		Flex_state: uint16(10),
	},
	262: {
		Flex_state: uint16(10),
	},
	263: {
		Flex_state: uint16(10),
	},
	264: {
		Flex_state: uint16(10),
	},
	265: {
		Flex_state: uint16(7),
	},
	266: {
		Flex_state: uint16(32),
	},
	267: {
		Flex_state: uint16(32),
	},
	268: {
		Flex_state: uint16(10),
	},
	269: {
		Flex_state: uint16(32),
	},
	270: {
		Flex_state: uint16(10),
	},
	271: {
		Flex_state: uint16(10),
	},
	272: {
		Flex_state: uint16(10),
	},
	273: {
		Flex_state: uint16(10),
	},
	274: {
		Flex_state: uint16(10),
	},
	275: {
		Flex_state: uint16(10),
	},
	276: {
		Flex_state: uint16(10),
	},
	277: {
		Flex_state: uint16(10),
	},
	278: {
		Flex_state: uint16(10),
	},
	279: {
		Flex_state: uint16(10),
	},
	280: {
		Flex_state: uint16(10),
	},
	281: {
		Flex_state: uint16(32),
	},
	282: {
		Flex_state: uint16(32),
	},
	283: {
		Flex_state: uint16(32),
	},
	284: {
		Flex_state: uint16(32),
	},
	285: {
		Flex_state: uint16(32),
	},
	286: {
		Flex_state: uint16(11),
	},
	287: {
		Flex_state: uint16(12),
	},
	288: {
		Flex_state: uint16(12),
	},
	289: {
		Flex_state: uint16(11),
	},
	290: {
		Flex_state: uint16(12),
	},
	291: {
		Flex_state: uint16(11),
	},
	292: {
		Flex_state: uint16(12),
	},
	293: {
		Flex_state: uint16(11),
	},
	294: {
		Flex_state: uint16(11),
	},
	295: {
		Flex_state: uint16(12),
	},
	296: {
		Flex_state: uint16(11),
	},
	297: {
		Flex_state: uint16(11),
	},
	298: {
		Flex_state: uint16(12),
	},
	299: {
		Flex_state: uint16(11),
	},
	300: {
		Flex_state: uint16(12),
	},
	301: {
		Flex_state: uint16(11),
	},
	302: {
		Flex_state: uint16(12),
	},
	303: {
		Flex_state: uint16(11),
	},
	304: {
		Flex_state: uint16(11),
	},
	305: {
		Flex_state: uint16(12),
	},
	306: {
		Flex_state: uint16(12),
	},
	307: {
		Flex_state: uint16(12),
	},
	308: {
		Flex_state: uint16(11),
	},
	309: {
		Flex_state: uint16(11),
	},
	310: {
		Flex_state: uint16(12),
	},
	311: {
		Flex_state: uint16(12),
	},
	312: {
		Flex_state: uint16(11),
	},
	313: {
		Flex_state: uint16(12),
	},
	314: {
		Flex_state: uint16(12),
	},
	315: {
		Flex_state: uint16(11),
	},
	316: {
		Flex_state: uint16(10),
	},
	317: {
		Flex_state: uint16(10),
	},
	318: {
		Flex_state: uint16(10),
	},
	319: {
		Flex_state: uint16(10),
	},
	320: {
		Flex_state: uint16(10),
	},
	321: {},
	322: {},
	323: {
		Flex_state: uint16(11),
	},
	324: {
		Flex_state: uint16(12),
	},
	325: {
		Flex_state: uint16(12),
	},
	326: {
		Flex_state: uint16(11),
	},
	327: {
		Flex_state: uint16(7),
	},
	328: {
		Flex_state: uint16(32),
	},
	329: {
		Flex_state: uint16(7),
	},
	330: {
		Flex_state: uint16(7),
	},
	331: {
		Flex_state: uint16(32),
	},
	332: {},
	333: {},
	334: {
		Flex_state: uint16(7),
	},
	335: {
		Flex_state: uint16(7),
	},
	336: {},
	337: {
		Flex_state: uint16(7),
	},
	338: {},
	339: {
		Flex_state: uint16(7),
	},
	340: {
		Flex_state: uint16(7),
	},
	341: {},
	342: {
		Flex_state: uint16(32),
	},
	343: {
		Flex_state: uint16(32),
	},
	344: {},
	345: {},
	346: {},
	347: {
		Flex_state: uint16(7),
	},
	348: {
		Flex_state: uint16(7),
	},
	349: {
		Flex_state: uint16(32),
	},
	350: {
		Flex_state: uint16(32),
	},
	351: {},
	352: {},
	353: {},
	354: {
		Flex_state: uint16(7),
	},
	355: {
		Flex_state: uint16(7),
	},
	356: {
		Flex_state: uint16(7),
	},
	357: {
		Flex_state: uint16(7),
	},
	358: {
		Flex_state: uint16(7),
	},
	359: {
		Flex_state: uint16(7),
	},
	360: {
		Fexternal_lex_state: uint16(1),
	},
	361: {
		Flex_state: uint16(7),
	},
	362: {
		Flex_state: uint16(7),
	},
	363: {
		Flex_state: uint16(3),
	},
	364: {
		Flex_state: uint16(7),
	},
	365: {
		Flex_state: uint16(7),
	},
	366: {
		Flex_state: uint16(7),
	},
	367: {},
	368: {
		Flex_state: uint16(3),
	},
	369: {
		Flex_state: uint16(7),
	},
	370: {
		Flex_state: uint16(7),
	},
	371: {
		Flex_state: uint16(7),
	},
	372: {
		Flex_state: uint16(32),
	},
	373: {
		Flex_state: uint16(7),
	},
	374: {
		Flex_state: uint16(32),
	},
	375: {
		Flex_state: uint16(7),
	},
	376: {
		Flex_state: uint16(7),
	},
	377: {
		Flex_state: uint16(7),
	},
	378: {
		Flex_state: uint16(7),
	},
	379: {
		Flex_state: uint16(7),
	},
	380: {
		Flex_state: uint16(7),
	},
	381: {
		Fexternal_lex_state: uint16(1),
	},
	382: {
		Flex_state: uint16(7),
	},
}

var ts_parse_table = [10][103]uint16_t{
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
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(1),
		27: uint16(1),
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
		52: uint16(1),
		53: uint16(1),
		54: uint16(1),
		56: uint16(1),
		59: uint16(1),
		62: uint16(3),
		63: uint16(1),
	},
	1: {
		0:  uint16(5),
		1:  uint16(7),
		2:  uint16(9),
		3:  uint16(11),
		4:  uint16(13),
		5:  uint16(15),
		6:  uint16(17),
		8:  uint16(19),
		9:  uint16(21),
		11: uint16(23),
		13: uint16(25),
		14: uint16(25),
		15: uint16(25),
		16: uint16(25),
		62: uint16(3),
		64: uint16(367),
		65: uint16(104),
		66: uint16(104),
		67: uint16(104),
		68: uint16(104),
		69: uint16(104),
		70: uint16(104),
		71: uint16(104),
		72: uint16(104),
		73: uint16(104),
		74: uint16(104),
		75: uint16(104),
		95: uint16(104),
	},
	2: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		8:  uint16(39),
		9:  uint16(41),
		10: uint16(43),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(69),
		43: uint16(69),
		62: uint16(3),
		65: uint16(102),
		66: uint16(102),
		67: uint16(102),
		68: uint16(102),
		69: uint16(102),
		70: uint16(102),
		71: uint16(102),
		72: uint16(102),
		73: uint16(102),
		74: uint16(102),
		75: uint16(102),
		76: uint16(5),
		77: uint16(5),
		78: uint16(5),
		79: uint16(5),
		80: uint16(5),
		81: uint16(5),
		82: uint16(5),
		83: uint16(5),
		84: uint16(5),
		85: uint16(5),
		86: uint16(5),
		95: uint16(102),
		96: uint16(5),
	},
	3: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		7:  uint16(71),
		8:  uint16(39),
		9:  uint16(41),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(73),
		43: uint16(73),
		62: uint16(3),
		65: uint16(112),
		66: uint16(112),
		67: uint16(112),
		68: uint16(112),
		69: uint16(112),
		70: uint16(112),
		71: uint16(112),
		72: uint16(112),
		73: uint16(112),
		74: uint16(112),
		75: uint16(112),
		76: uint16(12),
		77: uint16(12),
		78: uint16(12),
		79: uint16(12),
		80: uint16(12),
		81: uint16(12),
		82: uint16(12),
		83: uint16(12),
		84: uint16(12),
		85: uint16(12),
		86: uint16(12),
		95: uint16(112),
		96: uint16(12),
	},
	4: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		7:  uint16(75),
		8:  uint16(39),
		9:  uint16(41),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(73),
		43: uint16(73),
		62: uint16(3),
		65: uint16(109),
		66: uint16(109),
		67: uint16(109),
		68: uint16(109),
		69: uint16(109),
		70: uint16(109),
		71: uint16(109),
		72: uint16(109),
		73: uint16(109),
		74: uint16(109),
		75: uint16(109),
		76: uint16(12),
		77: uint16(12),
		78: uint16(12),
		79: uint16(12),
		80: uint16(12),
		81: uint16(12),
		82: uint16(12),
		83: uint16(12),
		84: uint16(12),
		85: uint16(12),
		86: uint16(12),
		95: uint16(109),
		96: uint16(12),
	},
	5: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		8:  uint16(39),
		9:  uint16(41),
		10: uint16(77),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(73),
		43: uint16(73),
		62: uint16(3),
		65: uint16(105),
		66: uint16(105),
		67: uint16(105),
		68: uint16(105),
		69: uint16(105),
		70: uint16(105),
		71: uint16(105),
		72: uint16(105),
		73: uint16(105),
		74: uint16(105),
		75: uint16(105),
		76: uint16(12),
		77: uint16(12),
		78: uint16(12),
		79: uint16(12),
		80: uint16(12),
		81: uint16(12),
		82: uint16(12),
		83: uint16(12),
		84: uint16(12),
		85: uint16(12),
		86: uint16(12),
		95: uint16(105),
		96: uint16(12),
	},
	6: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		8:  uint16(39),
		9:  uint16(41),
		10: uint16(79),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(81),
		43: uint16(81),
		62: uint16(3),
		65: uint16(108),
		66: uint16(108),
		67: uint16(108),
		68: uint16(108),
		69: uint16(108),
		70: uint16(108),
		71: uint16(108),
		72: uint16(108),
		73: uint16(108),
		74: uint16(108),
		75: uint16(108),
		76: uint16(9),
		77: uint16(9),
		78: uint16(9),
		79: uint16(9),
		80: uint16(9),
		81: uint16(9),
		82: uint16(9),
		83: uint16(9),
		84: uint16(9),
		85: uint16(9),
		86: uint16(9),
		95: uint16(108),
		96: uint16(9),
	},
	7: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		7:  uint16(83),
		8:  uint16(39),
		9:  uint16(41),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(73),
		43: uint16(73),
		62: uint16(3),
		65: uint16(107),
		66: uint16(107),
		67: uint16(107),
		68: uint16(107),
		69: uint16(107),
		70: uint16(107),
		71: uint16(107),
		72: uint16(107),
		73: uint16(107),
		74: uint16(107),
		75: uint16(107),
		76: uint16(12),
		77: uint16(12),
		78: uint16(12),
		79: uint16(12),
		80: uint16(12),
		81: uint16(12),
		82: uint16(12),
		83: uint16(12),
		84: uint16(12),
		85: uint16(12),
		86: uint16(12),
		95: uint16(107),
		96: uint16(12),
	},
	8: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		7:  uint16(85),
		8:  uint16(39),
		9:  uint16(41),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(73),
		43: uint16(73),
		62: uint16(3),
		65: uint16(111),
		66: uint16(111),
		67: uint16(111),
		68: uint16(111),
		69: uint16(111),
		70: uint16(111),
		71: uint16(111),
		72: uint16(111),
		73: uint16(111),
		74: uint16(111),
		75: uint16(111),
		76: uint16(12),
		77: uint16(12),
		78: uint16(12),
		79: uint16(12),
		80: uint16(12),
		81: uint16(12),
		82: uint16(12),
		83: uint16(12),
		84: uint16(12),
		85: uint16(12),
		86: uint16(12),
		95: uint16(111),
		96: uint16(12),
	},
	9: {
		1:  uint16(27),
		2:  uint16(29),
		3:  uint16(31),
		4:  uint16(33),
		5:  uint16(35),
		6:  uint16(37),
		8:  uint16(39),
		9:  uint16(41),
		10: uint16(87),
		11: uint16(45),
		13: uint16(47),
		14: uint16(47),
		15: uint16(47),
		16: uint16(47),
		23: uint16(49),
		24: uint16(49),
		25: uint16(49),
		26: uint16(49),
		27: uint16(49),
		29: uint16(51),
		30: uint16(53),
		31: uint16(55),
		32: uint16(55),
		33: uint16(55),
		34: uint16(55),
		35: uint16(55),
		36: uint16(57),
		37: uint16(59),
		38: uint16(61),
		39: uint16(63),
		40: uint16(65),
		41: uint16(67),
		42: uint16(73),
		43: uint16(73),
		62: uint16(3),
		65: uint16(101),
		66: uint16(101),
		67: uint16(101),
		68: uint16(101),
		69: uint16(101),
		70: uint16(101),
		71: uint16(101),
		72: uint16(101),
		73: uint16(101),
		74: uint16(101),
		75: uint16(101),
		76: uint16(12),
		77: uint16(12),
		78: uint16(12),
		79: uint16(12),
		80: uint16(12),
		81: uint16(12),
		82: uint16(12),
		83: uint16(12),
		84: uint16(12),
		85: uint16(12),
		86: uint16(12),
		95: uint16(101),
		96: uint16(12),
	},
}

var ts_small_parse_table = [11811]uint16_t{
	0:     uint16(8),
	1:     uint16(3),
	2:     uint16(1),
	3:     uint16(sym_comment),
	4:     uint16(89),
	5:     uint16(1),
	6:     uint16(sym_symbol),
	7:     uint16(95),
	8:     uint16(1),
	9:     uint16(anon_sym_DOLLAR_LPAREN),
	10:    uint16(97),
	11:    uint16(1),
	12:    uint16(anon_sym_DQUOTE),
	13:    uint16(99),
	14:    uint16(1),
	15:    uint16(anon_sym_SQUOTE),
	16:    uint16(11),
	17:    uint16(3),
	18:    uint16(sym_macro_variable),
	19:    uint16(sym_string),
	20:    uint16(aux_sym_name_repeat1),
	21:    uint16(93),
	22:    uint16(8),
	23:    uint16(anon_sym_EQ),
	24:    uint16(anon_sym_dependson),
	25:    uint16(anon_sym_visibleif),
	26:    uint16(anon_sym_PIPE_PIPE),
	27:    uint16(anon_sym_AMP_AMP),
	28:    uint16(anon_sym_BANG_EQ),
	29:    uint16(anon_sym_LT_EQ),
	30:    uint16(anon_sym_GT_EQ),
	31:    uint16(91),
	32:    uint16(35),
	33:    uint16(anon_sym_mainmenu),
	34:    uint16(anon_sym_config),
	35:    uint16(anon_sym_configdefault),
	36:    uint16(anon_sym_menuconfig),
	37:    uint16(anon_sym_choice),
	38:    uint16(anon_sym_endchoice),
	39:    uint16(anon_sym_comment),
	40:    uint16(anon_sym_menu),
	41:    uint16(anon_sym_endmenu),
	42:    uint16(anon_sym_if),
	43:    uint16(anon_sym_endif),
	44:    uint16(anon_sym_source),
	45:    uint16(anon_sym_rsource),
	46:    uint16(anon_sym_osource),
	47:    uint16(anon_sym_orsource),
	48:    uint16(anon_sym_bool),
	49:    uint16(anon_sym_tristate),
	50:    uint16(anon_sym_int),
	51:    uint16(anon_sym_hex),
	52:    uint16(anon_sym_string),
	53:    uint16(anon_sym_prompt),
	54:    uint16(anon_sym_default),
	55:    uint16(anon_sym_def_bool),
	56:    uint16(anon_sym_def_tristate),
	57:    uint16(anon_sym_def_int),
	58:    uint16(anon_sym_def_hex),
	59:    uint16(anon_sym_def_string),
	60:    uint16(anon_sym_select),
	61:    uint16(anon_sym_imply),
	62:    uint16(anon_sym_range),
	63:    uint16(anon_sym_help),
	64:    uint16(sym_optional),
	65:    uint16(sym_modules),
	66:    uint16(anon_sym_LT),
	67:    uint16(anon_sym_GT),
	68:    uint16(8),
	69:    uint16(3),
	70:    uint16(1),
	71:    uint16(sym_comment),
	72:    uint16(101),
	73:    uint16(1),
	74:    uint16(sym_symbol),
	75:    uint16(108),
	76:    uint16(1),
	77:    uint16(anon_sym_DOLLAR_LPAREN),
	78:    uint16(111),
	79:    uint16(1),
	80:    uint16(anon_sym_DQUOTE),
	81:    uint16(114),
	82:    uint16(1),
	83:    uint16(anon_sym_SQUOTE),
	84:    uint16(11),
	85:    uint16(3),
	86:    uint16(sym_macro_variable),
	87:    uint16(sym_string),
	88:    uint16(aux_sym_name_repeat1),
	89:    uint16(106),
	90:    uint16(8),
	91:    uint16(anon_sym_EQ),
	92:    uint16(anon_sym_dependson),
	93:    uint16(anon_sym_visibleif),
	94:    uint16(anon_sym_PIPE_PIPE),
	95:    uint16(anon_sym_AMP_AMP),
	96:    uint16(anon_sym_BANG_EQ),
	97:    uint16(anon_sym_LT_EQ),
	98:    uint16(anon_sym_GT_EQ),
	99:    uint16(104),
	100:   uint16(35),
	101:   uint16(anon_sym_mainmenu),
	102:   uint16(anon_sym_config),
	103:   uint16(anon_sym_configdefault),
	104:   uint16(anon_sym_menuconfig),
	105:   uint16(anon_sym_choice),
	106:   uint16(anon_sym_endchoice),
	107:   uint16(anon_sym_comment),
	108:   uint16(anon_sym_menu),
	109:   uint16(anon_sym_endmenu),
	110:   uint16(anon_sym_if),
	111:   uint16(anon_sym_endif),
	112:   uint16(anon_sym_source),
	113:   uint16(anon_sym_rsource),
	114:   uint16(anon_sym_osource),
	115:   uint16(anon_sym_orsource),
	116:   uint16(anon_sym_bool),
	117:   uint16(anon_sym_tristate),
	118:   uint16(anon_sym_int),
	119:   uint16(anon_sym_hex),
	120:   uint16(anon_sym_string),
	121:   uint16(anon_sym_prompt),
	122:   uint16(anon_sym_default),
	123:   uint16(anon_sym_def_bool),
	124:   uint16(anon_sym_def_tristate),
	125:   uint16(anon_sym_def_int),
	126:   uint16(anon_sym_def_hex),
	127:   uint16(anon_sym_def_string),
	128:   uint16(anon_sym_select),
	129:   uint16(anon_sym_imply),
	130:   uint16(anon_sym_range),
	131:   uint16(anon_sym_help),
	132:   uint16(sym_optional),
	133:   uint16(sym_modules),
	134:   uint16(anon_sym_LT),
	135:   uint16(anon_sym_GT),
	136:   uint16(14),
	137:   uint16(3),
	138:   uint16(1),
	139:   uint16(sym_comment),
	140:   uint16(122),
	141:   uint16(1),
	142:   uint16(anon_sym_prompt),
	143:   uint16(125),
	144:   uint16(1),
	145:   uint16(anon_sym_default),
	146:   uint16(131),
	147:   uint16(1),
	148:   uint16(anon_sym_dependson),
	149:   uint16(134),
	150:   uint16(1),
	151:   uint16(anon_sym_select),
	152:   uint16(137),
	153:   uint16(1),
	154:   uint16(anon_sym_imply),
	155:   uint16(140),
	156:   uint16(1),
	157:   uint16(anon_sym_visibleif),
	158:   uint16(143),
	159:   uint16(1),
	160:   uint16(anon_sym_range),
	161:   uint16(146),
	162:   uint16(1),
	163:   uint16(anon_sym_help),
	164:   uint16(149),
	165:   uint16(2),
	166:   uint16(sym_optional),
	167:   uint16(sym_modules),
	168:   uint16(119),
	169:   uint16(5),
	170:   uint16(anon_sym_bool),
	171:   uint16(anon_sym_tristate),
	172:   uint16(anon_sym_int),
	173:   uint16(anon_sym_hex),
	174:   uint16(anon_sym_string),
	175:   uint16(128),
	176:   uint16(5),
	177:   uint16(anon_sym_def_bool),
	178:   uint16(anon_sym_def_tristate),
	179:   uint16(anon_sym_def_int),
	180:   uint16(anon_sym_def_hex),
	181:   uint16(anon_sym_def_string),
	182:   uint16(12),
	183:   uint16(12),
	184:   uint16(sym__config_option),
	185:   uint16(sym_type_definition),
	186:   uint16(sym_input_prompt),
	187:   uint16(sym_default_value),
	188:   uint16(sym_type_definition_default),
	189:   uint16(sym_dependencies),
	190:   uint16(sym_reverse_dependencies),
	191:   uint16(sym_weak_reverse_dependencies),
	192:   uint16(sym_limiting_menu_display),
	193:   uint16(sym_numerical_ranges),
	194:   uint16(sym_help_text),
	195:   uint16(aux_sym_config_repeat1),
	196:   uint16(117),
	197:   uint16(16),
	198:   uint16(anon_sym_mainmenu),
	199:   uint16(anon_sym_config),
	200:   uint16(anon_sym_configdefault),
	201:   uint16(anon_sym_menuconfig),
	202:   uint16(anon_sym_choice),
	203:   uint16(anon_sym_endchoice),
	204:   uint16(anon_sym_comment),
	205:   uint16(anon_sym_menu),
	206:   uint16(anon_sym_endmenu),
	207:   uint16(anon_sym_if),
	208:   uint16(anon_sym_endif),
	209:   uint16(anon_sym_source),
	210:   uint16(anon_sym_rsource),
	211:   uint16(anon_sym_osource),
	212:   uint16(anon_sym_orsource),
	213:   uint16(sym_symbol),
	214:   uint16(14),
	215:   uint16(3),
	216:   uint16(1),
	217:   uint16(sym_comment),
	218:   uint16(51),
	219:   uint16(1),
	220:   uint16(anon_sym_prompt),
	221:   uint16(53),
	222:   uint16(1),
	223:   uint16(anon_sym_default),
	224:   uint16(57),
	225:   uint16(1),
	226:   uint16(anon_sym_dependson),
	227:   uint16(59),
	228:   uint16(1),
	229:   uint16(anon_sym_select),
	230:   uint16(61),
	231:   uint16(1),
	232:   uint16(anon_sym_imply),
	233:   uint16(63),
	234:   uint16(1),
	235:   uint16(anon_sym_visibleif),
	236:   uint16(65),
	237:   uint16(1),
	238:   uint16(anon_sym_range),
	239:   uint16(67),
	240:   uint16(1),
	241:   uint16(anon_sym_help),
	242:   uint16(154),
	243:   uint16(2),
	244:   uint16(sym_optional),
	245:   uint16(sym_modules),
	246:   uint16(49),
	247:   uint16(5),
	248:   uint16(anon_sym_bool),
	249:   uint16(anon_sym_tristate),
	250:   uint16(anon_sym_int),
	251:   uint16(anon_sym_hex),
	252:   uint16(anon_sym_string),
	253:   uint16(55),
	254:   uint16(5),
	255:   uint16(anon_sym_def_bool),
	256:   uint16(anon_sym_def_tristate),
	257:   uint16(anon_sym_def_int),
	258:   uint16(anon_sym_def_hex),
	259:   uint16(anon_sym_def_string),
	260:   uint16(17),
	261:   uint16(12),
	262:   uint16(sym__config_option),
	263:   uint16(sym_type_definition),
	264:   uint16(sym_input_prompt),
	265:   uint16(sym_default_value),
	266:   uint16(sym_type_definition_default),
	267:   uint16(sym_dependencies),
	268:   uint16(sym_reverse_dependencies),
	269:   uint16(sym_weak_reverse_dependencies),
	270:   uint16(sym_limiting_menu_display),
	271:   uint16(sym_numerical_ranges),
	272:   uint16(sym_help_text),
	273:   uint16(aux_sym_config_repeat1),
	274:   uint16(152),
	275:   uint16(16),
	276:   uint16(anon_sym_mainmenu),
	277:   uint16(anon_sym_config),
	278:   uint16(anon_sym_configdefault),
	279:   uint16(anon_sym_menuconfig),
	280:   uint16(anon_sym_choice),
	281:   uint16(anon_sym_endchoice),
	282:   uint16(anon_sym_comment),
	283:   uint16(anon_sym_menu),
	284:   uint16(anon_sym_endmenu),
	285:   uint16(anon_sym_if),
	286:   uint16(anon_sym_endif),
	287:   uint16(anon_sym_source),
	288:   uint16(anon_sym_rsource),
	289:   uint16(anon_sym_osource),
	290:   uint16(anon_sym_orsource),
	291:   uint16(sym_symbol),
	292:   uint16(14),
	293:   uint16(3),
	294:   uint16(1),
	295:   uint16(sym_comment),
	296:   uint16(51),
	297:   uint16(1),
	298:   uint16(anon_sym_prompt),
	299:   uint16(53),
	300:   uint16(1),
	301:   uint16(anon_sym_default),
	302:   uint16(57),
	303:   uint16(1),
	304:   uint16(anon_sym_dependson),
	305:   uint16(59),
	306:   uint16(1),
	307:   uint16(anon_sym_select),
	308:   uint16(61),
	309:   uint16(1),
	310:   uint16(anon_sym_imply),
	311:   uint16(63),
	312:   uint16(1),
	313:   uint16(anon_sym_visibleif),
	314:   uint16(65),
	315:   uint16(1),
	316:   uint16(anon_sym_range),
	317:   uint16(67),
	318:   uint16(1),
	319:   uint16(anon_sym_help),
	320:   uint16(73),
	321:   uint16(2),
	322:   uint16(sym_optional),
	323:   uint16(sym_modules),
	324:   uint16(49),
	325:   uint16(5),
	326:   uint16(anon_sym_bool),
	327:   uint16(anon_sym_tristate),
	328:   uint16(anon_sym_int),
	329:   uint16(anon_sym_hex),
	330:   uint16(anon_sym_string),
	331:   uint16(55),
	332:   uint16(5),
	333:   uint16(anon_sym_def_bool),
	334:   uint16(anon_sym_def_tristate),
	335:   uint16(anon_sym_def_int),
	336:   uint16(anon_sym_def_hex),
	337:   uint16(anon_sym_def_string),
	338:   uint16(12),
	339:   uint16(12),
	340:   uint16(sym__config_option),
	341:   uint16(sym_type_definition),
	342:   uint16(sym_input_prompt),
	343:   uint16(sym_default_value),
	344:   uint16(sym_type_definition_default),
	345:   uint16(sym_dependencies),
	346:   uint16(sym_reverse_dependencies),
	347:   uint16(sym_weak_reverse_dependencies),
	348:   uint16(sym_limiting_menu_display),
	349:   uint16(sym_numerical_ranges),
	350:   uint16(sym_help_text),
	351:   uint16(aux_sym_config_repeat1),
	352:   uint16(156),
	353:   uint16(16),
	354:   uint16(anon_sym_mainmenu),
	355:   uint16(anon_sym_config),
	356:   uint16(anon_sym_configdefault),
	357:   uint16(anon_sym_menuconfig),
	358:   uint16(anon_sym_choice),
	359:   uint16(anon_sym_endchoice),
	360:   uint16(anon_sym_comment),
	361:   uint16(anon_sym_menu),
	362:   uint16(anon_sym_endmenu),
	363:   uint16(anon_sym_if),
	364:   uint16(anon_sym_endif),
	365:   uint16(anon_sym_source),
	366:   uint16(anon_sym_rsource),
	367:   uint16(anon_sym_osource),
	368:   uint16(anon_sym_orsource),
	369:   uint16(sym_symbol),
	370:   uint16(14),
	371:   uint16(3),
	372:   uint16(1),
	373:   uint16(sym_comment),
	374:   uint16(51),
	375:   uint16(1),
	376:   uint16(anon_sym_prompt),
	377:   uint16(53),
	378:   uint16(1),
	379:   uint16(anon_sym_default),
	380:   uint16(57),
	381:   uint16(1),
	382:   uint16(anon_sym_dependson),
	383:   uint16(59),
	384:   uint16(1),
	385:   uint16(anon_sym_select),
	386:   uint16(61),
	387:   uint16(1),
	388:   uint16(anon_sym_imply),
	389:   uint16(63),
	390:   uint16(1),
	391:   uint16(anon_sym_visibleif),
	392:   uint16(65),
	393:   uint16(1),
	394:   uint16(anon_sym_range),
	395:   uint16(67),
	396:   uint16(1),
	397:   uint16(anon_sym_help),
	398:   uint16(73),
	399:   uint16(2),
	400:   uint16(sym_optional),
	401:   uint16(sym_modules),
	402:   uint16(49),
	403:   uint16(5),
	404:   uint16(anon_sym_bool),
	405:   uint16(anon_sym_tristate),
	406:   uint16(anon_sym_int),
	407:   uint16(anon_sym_hex),
	408:   uint16(anon_sym_string),
	409:   uint16(55),
	410:   uint16(5),
	411:   uint16(anon_sym_def_bool),
	412:   uint16(anon_sym_def_tristate),
	413:   uint16(anon_sym_def_int),
	414:   uint16(anon_sym_def_hex),
	415:   uint16(anon_sym_def_string),
	416:   uint16(12),
	417:   uint16(12),
	418:   uint16(sym__config_option),
	419:   uint16(sym_type_definition),
	420:   uint16(sym_input_prompt),
	421:   uint16(sym_default_value),
	422:   uint16(sym_type_definition_default),
	423:   uint16(sym_dependencies),
	424:   uint16(sym_reverse_dependencies),
	425:   uint16(sym_weak_reverse_dependencies),
	426:   uint16(sym_limiting_menu_display),
	427:   uint16(sym_numerical_ranges),
	428:   uint16(sym_help_text),
	429:   uint16(aux_sym_config_repeat1),
	430:   uint16(158),
	431:   uint16(16),
	432:   uint16(anon_sym_mainmenu),
	433:   uint16(anon_sym_config),
	434:   uint16(anon_sym_configdefault),
	435:   uint16(anon_sym_menuconfig),
	436:   uint16(anon_sym_choice),
	437:   uint16(anon_sym_endchoice),
	438:   uint16(anon_sym_comment),
	439:   uint16(anon_sym_menu),
	440:   uint16(anon_sym_endmenu),
	441:   uint16(anon_sym_if),
	442:   uint16(anon_sym_endif),
	443:   uint16(anon_sym_source),
	444:   uint16(anon_sym_rsource),
	445:   uint16(anon_sym_osource),
	446:   uint16(anon_sym_orsource),
	447:   uint16(sym_symbol),
	448:   uint16(8),
	449:   uint16(3),
	450:   uint16(1),
	451:   uint16(sym_comment),
	452:   uint16(160),
	453:   uint16(1),
	454:   uint16(sym_symbol),
	455:   uint16(163),
	456:   uint16(1),
	457:   uint16(anon_sym_DOLLAR_LPAREN),
	458:   uint16(166),
	459:   uint16(1),
	460:   uint16(anon_sym_DQUOTE),
	461:   uint16(169),
	462:   uint16(1),
	463:   uint16(anon_sym_SQUOTE),
	464:   uint16(16),
	465:   uint16(3),
	466:   uint16(sym_macro_variable),
	467:   uint16(sym_string),
	468:   uint16(aux_sym_name_repeat1),
	469:   uint16(106),
	470:   uint16(9),
	472:   uint16(anon_sym_EQ),
	473:   uint16(anon_sym_dependson),
	474:   uint16(anon_sym_visibleif),
	475:   uint16(anon_sym_PIPE_PIPE),
	476:   uint16(anon_sym_AMP_AMP),
	477:   uint16(anon_sym_BANG_EQ),
	478:   uint16(anon_sym_LT_EQ),
	479:   uint16(anon_sym_GT_EQ),
	480:   uint16(104),
	481:   uint16(32),
	482:   uint16(anon_sym_mainmenu),
	483:   uint16(anon_sym_config),
	484:   uint16(anon_sym_configdefault),
	485:   uint16(anon_sym_menuconfig),
	486:   uint16(anon_sym_choice),
	487:   uint16(anon_sym_comment),
	488:   uint16(anon_sym_menu),
	489:   uint16(anon_sym_if),
	490:   uint16(anon_sym_source),
	491:   uint16(anon_sym_rsource),
	492:   uint16(anon_sym_osource),
	493:   uint16(anon_sym_orsource),
	494:   uint16(anon_sym_bool),
	495:   uint16(anon_sym_tristate),
	496:   uint16(anon_sym_int),
	497:   uint16(anon_sym_hex),
	498:   uint16(anon_sym_string),
	499:   uint16(anon_sym_prompt),
	500:   uint16(anon_sym_default),
	501:   uint16(anon_sym_def_bool),
	502:   uint16(anon_sym_def_tristate),
	503:   uint16(anon_sym_def_int),
	504:   uint16(anon_sym_def_hex),
	505:   uint16(anon_sym_def_string),
	506:   uint16(anon_sym_select),
	507:   uint16(anon_sym_imply),
	508:   uint16(anon_sym_range),
	509:   uint16(anon_sym_help),
	510:   uint16(sym_optional),
	511:   uint16(sym_modules),
	512:   uint16(anon_sym_LT),
	513:   uint16(anon_sym_GT),
	514:   uint16(14),
	515:   uint16(3),
	516:   uint16(1),
	517:   uint16(sym_comment),
	518:   uint16(51),
	519:   uint16(1),
	520:   uint16(anon_sym_prompt),
	521:   uint16(53),
	522:   uint16(1),
	523:   uint16(anon_sym_default),
	524:   uint16(57),
	525:   uint16(1),
	526:   uint16(anon_sym_dependson),
	527:   uint16(59),
	528:   uint16(1),
	529:   uint16(anon_sym_select),
	530:   uint16(61),
	531:   uint16(1),
	532:   uint16(anon_sym_imply),
	533:   uint16(63),
	534:   uint16(1),
	535:   uint16(anon_sym_visibleif),
	536:   uint16(65),
	537:   uint16(1),
	538:   uint16(anon_sym_range),
	539:   uint16(67),
	540:   uint16(1),
	541:   uint16(anon_sym_help),
	542:   uint16(73),
	543:   uint16(2),
	544:   uint16(sym_optional),
	545:   uint16(sym_modules),
	546:   uint16(49),
	547:   uint16(5),
	548:   uint16(anon_sym_bool),
	549:   uint16(anon_sym_tristate),
	550:   uint16(anon_sym_int),
	551:   uint16(anon_sym_hex),
	552:   uint16(anon_sym_string),
	553:   uint16(55),
	554:   uint16(5),
	555:   uint16(anon_sym_def_bool),
	556:   uint16(anon_sym_def_tristate),
	557:   uint16(anon_sym_def_int),
	558:   uint16(anon_sym_def_hex),
	559:   uint16(anon_sym_def_string),
	560:   uint16(12),
	561:   uint16(12),
	562:   uint16(sym__config_option),
	563:   uint16(sym_type_definition),
	564:   uint16(sym_input_prompt),
	565:   uint16(sym_default_value),
	566:   uint16(sym_type_definition_default),
	567:   uint16(sym_dependencies),
	568:   uint16(sym_reverse_dependencies),
	569:   uint16(sym_weak_reverse_dependencies),
	570:   uint16(sym_limiting_menu_display),
	571:   uint16(sym_numerical_ranges),
	572:   uint16(sym_help_text),
	573:   uint16(aux_sym_config_repeat1),
	574:   uint16(172),
	575:   uint16(16),
	576:   uint16(anon_sym_mainmenu),
	577:   uint16(anon_sym_config),
	578:   uint16(anon_sym_configdefault),
	579:   uint16(anon_sym_menuconfig),
	580:   uint16(anon_sym_choice),
	581:   uint16(anon_sym_endchoice),
	582:   uint16(anon_sym_comment),
	583:   uint16(anon_sym_menu),
	584:   uint16(anon_sym_endmenu),
	585:   uint16(anon_sym_if),
	586:   uint16(anon_sym_endif),
	587:   uint16(anon_sym_source),
	588:   uint16(anon_sym_rsource),
	589:   uint16(anon_sym_osource),
	590:   uint16(anon_sym_orsource),
	591:   uint16(sym_symbol),
	592:   uint16(8),
	593:   uint16(3),
	594:   uint16(1),
	595:   uint16(sym_comment),
	596:   uint16(174),
	597:   uint16(1),
	598:   uint16(sym_symbol),
	599:   uint16(176),
	600:   uint16(1),
	601:   uint16(anon_sym_DOLLAR_LPAREN),
	602:   uint16(178),
	603:   uint16(1),
	604:   uint16(anon_sym_DQUOTE),
	605:   uint16(180),
	606:   uint16(1),
	607:   uint16(anon_sym_SQUOTE),
	608:   uint16(16),
	609:   uint16(3),
	610:   uint16(sym_macro_variable),
	611:   uint16(sym_string),
	612:   uint16(aux_sym_name_repeat1),
	613:   uint16(93),
	614:   uint16(9),
	616:   uint16(anon_sym_EQ),
	617:   uint16(anon_sym_dependson),
	618:   uint16(anon_sym_visibleif),
	619:   uint16(anon_sym_PIPE_PIPE),
	620:   uint16(anon_sym_AMP_AMP),
	621:   uint16(anon_sym_BANG_EQ),
	622:   uint16(anon_sym_LT_EQ),
	623:   uint16(anon_sym_GT_EQ),
	624:   uint16(91),
	625:   uint16(32),
	626:   uint16(anon_sym_mainmenu),
	627:   uint16(anon_sym_config),
	628:   uint16(anon_sym_configdefault),
	629:   uint16(anon_sym_menuconfig),
	630:   uint16(anon_sym_choice),
	631:   uint16(anon_sym_comment),
	632:   uint16(anon_sym_menu),
	633:   uint16(anon_sym_if),
	634:   uint16(anon_sym_source),
	635:   uint16(anon_sym_rsource),
	636:   uint16(anon_sym_osource),
	637:   uint16(anon_sym_orsource),
	638:   uint16(anon_sym_bool),
	639:   uint16(anon_sym_tristate),
	640:   uint16(anon_sym_int),
	641:   uint16(anon_sym_hex),
	642:   uint16(anon_sym_string),
	643:   uint16(anon_sym_prompt),
	644:   uint16(anon_sym_default),
	645:   uint16(anon_sym_def_bool),
	646:   uint16(anon_sym_def_tristate),
	647:   uint16(anon_sym_def_int),
	648:   uint16(anon_sym_def_hex),
	649:   uint16(anon_sym_def_string),
	650:   uint16(anon_sym_select),
	651:   uint16(anon_sym_imply),
	652:   uint16(anon_sym_range),
	653:   uint16(anon_sym_help),
	654:   uint16(sym_optional),
	655:   uint16(sym_modules),
	656:   uint16(anon_sym_LT),
	657:   uint16(anon_sym_GT),
	658:   uint16(3),
	659:   uint16(3),
	660:   uint16(1),
	661:   uint16(sym_comment),
	662:   uint16(184),
	663:   uint16(11),
	664:   uint16(anon_sym_EQ),
	665:   uint16(anon_sym_dependson),
	666:   uint16(anon_sym_visibleif),
	667:   uint16(anon_sym_PIPE_PIPE),
	668:   uint16(anon_sym_AMP_AMP),
	669:   uint16(anon_sym_BANG_EQ),
	670:   uint16(anon_sym_LT_EQ),
	671:   uint16(anon_sym_GT_EQ),
	672:   uint16(anon_sym_DOLLAR_LPAREN),
	673:   uint16(anon_sym_DQUOTE),
	674:   uint16(anon_sym_SQUOTE),
	675:   uint16(182),
	676:   uint16(36),
	677:   uint16(anon_sym_mainmenu),
	678:   uint16(anon_sym_config),
	679:   uint16(anon_sym_configdefault),
	680:   uint16(anon_sym_menuconfig),
	681:   uint16(anon_sym_choice),
	682:   uint16(anon_sym_endchoice),
	683:   uint16(anon_sym_comment),
	684:   uint16(anon_sym_menu),
	685:   uint16(anon_sym_endmenu),
	686:   uint16(anon_sym_if),
	687:   uint16(anon_sym_endif),
	688:   uint16(anon_sym_source),
	689:   uint16(anon_sym_rsource),
	690:   uint16(anon_sym_osource),
	691:   uint16(anon_sym_orsource),
	692:   uint16(anon_sym_bool),
	693:   uint16(anon_sym_tristate),
	694:   uint16(anon_sym_int),
	695:   uint16(anon_sym_hex),
	696:   uint16(anon_sym_string),
	697:   uint16(anon_sym_prompt),
	698:   uint16(anon_sym_default),
	699:   uint16(anon_sym_def_bool),
	700:   uint16(anon_sym_def_tristate),
	701:   uint16(anon_sym_def_int),
	702:   uint16(anon_sym_def_hex),
	703:   uint16(anon_sym_def_string),
	704:   uint16(anon_sym_select),
	705:   uint16(anon_sym_imply),
	706:   uint16(anon_sym_range),
	707:   uint16(anon_sym_help),
	708:   uint16(sym_optional),
	709:   uint16(sym_modules),
	710:   uint16(anon_sym_LT),
	711:   uint16(anon_sym_GT),
	712:   uint16(sym_symbol),
	713:   uint16(3),
	714:   uint16(3),
	715:   uint16(1),
	716:   uint16(sym_comment),
	717:   uint16(188),
	718:   uint16(11),
	719:   uint16(anon_sym_EQ),
	720:   uint16(anon_sym_dependson),
	721:   uint16(anon_sym_visibleif),
	722:   uint16(anon_sym_PIPE_PIPE),
	723:   uint16(anon_sym_AMP_AMP),
	724:   uint16(anon_sym_BANG_EQ),
	725:   uint16(anon_sym_LT_EQ),
	726:   uint16(anon_sym_GT_EQ),
	727:   uint16(anon_sym_DOLLAR_LPAREN),
	728:   uint16(anon_sym_DQUOTE),
	729:   uint16(anon_sym_SQUOTE),
	730:   uint16(186),
	731:   uint16(36),
	732:   uint16(anon_sym_mainmenu),
	733:   uint16(anon_sym_config),
	734:   uint16(anon_sym_configdefault),
	735:   uint16(anon_sym_menuconfig),
	736:   uint16(anon_sym_choice),
	737:   uint16(anon_sym_endchoice),
	738:   uint16(anon_sym_comment),
	739:   uint16(anon_sym_menu),
	740:   uint16(anon_sym_endmenu),
	741:   uint16(anon_sym_if),
	742:   uint16(anon_sym_endif),
	743:   uint16(anon_sym_source),
	744:   uint16(anon_sym_rsource),
	745:   uint16(anon_sym_osource),
	746:   uint16(anon_sym_orsource),
	747:   uint16(anon_sym_bool),
	748:   uint16(anon_sym_tristate),
	749:   uint16(anon_sym_int),
	750:   uint16(anon_sym_hex),
	751:   uint16(anon_sym_string),
	752:   uint16(anon_sym_prompt),
	753:   uint16(anon_sym_default),
	754:   uint16(anon_sym_def_bool),
	755:   uint16(anon_sym_def_tristate),
	756:   uint16(anon_sym_def_int),
	757:   uint16(anon_sym_def_hex),
	758:   uint16(anon_sym_def_string),
	759:   uint16(anon_sym_select),
	760:   uint16(anon_sym_imply),
	761:   uint16(anon_sym_range),
	762:   uint16(anon_sym_help),
	763:   uint16(sym_optional),
	764:   uint16(sym_modules),
	765:   uint16(anon_sym_LT),
	766:   uint16(anon_sym_GT),
	767:   uint16(sym_symbol),
	768:   uint16(4),
	769:   uint16(3),
	770:   uint16(1),
	771:   uint16(sym_comment),
	772:   uint16(196),
	773:   uint16(3),
	774:   uint16(anon_sym_DOLLAR_LPAREN),
	775:   uint16(anon_sym_DQUOTE),
	776:   uint16(anon_sym_SQUOTE),
	777:   uint16(193),
	778:   uint16(8),
	779:   uint16(anon_sym_EQ),
	780:   uint16(anon_sym_dependson),
	781:   uint16(anon_sym_visibleif),
	782:   uint16(anon_sym_PIPE_PIPE),
	783:   uint16(anon_sym_AMP_AMP),
	784:   uint16(anon_sym_BANG_EQ),
	785:   uint16(anon_sym_LT_EQ),
	786:   uint16(anon_sym_GT_EQ),
	787:   uint16(190),
	788:   uint16(36),
	789:   uint16(anon_sym_mainmenu),
	790:   uint16(anon_sym_config),
	791:   uint16(anon_sym_configdefault),
	792:   uint16(anon_sym_menuconfig),
	793:   uint16(anon_sym_choice),
	794:   uint16(anon_sym_endchoice),
	795:   uint16(anon_sym_comment),
	796:   uint16(anon_sym_menu),
	797:   uint16(anon_sym_endmenu),
	798:   uint16(anon_sym_if),
	799:   uint16(anon_sym_endif),
	800:   uint16(anon_sym_source),
	801:   uint16(anon_sym_rsource),
	802:   uint16(anon_sym_osource),
	803:   uint16(anon_sym_orsource),
	804:   uint16(anon_sym_bool),
	805:   uint16(anon_sym_tristate),
	806:   uint16(anon_sym_int),
	807:   uint16(anon_sym_hex),
	808:   uint16(anon_sym_string),
	809:   uint16(anon_sym_prompt),
	810:   uint16(anon_sym_default),
	811:   uint16(anon_sym_def_bool),
	812:   uint16(anon_sym_def_tristate),
	813:   uint16(anon_sym_def_int),
	814:   uint16(anon_sym_def_hex),
	815:   uint16(anon_sym_def_string),
	816:   uint16(anon_sym_select),
	817:   uint16(anon_sym_imply),
	818:   uint16(anon_sym_range),
	819:   uint16(anon_sym_help),
	820:   uint16(sym_optional),
	821:   uint16(sym_modules),
	822:   uint16(anon_sym_LT),
	823:   uint16(anon_sym_GT),
	824:   uint16(sym_symbol),
	825:   uint16(3),
	826:   uint16(3),
	827:   uint16(1),
	828:   uint16(sym_comment),
	829:   uint16(200),
	830:   uint16(11),
	831:   uint16(anon_sym_EQ),
	832:   uint16(anon_sym_dependson),
	833:   uint16(anon_sym_visibleif),
	834:   uint16(anon_sym_PIPE_PIPE),
	835:   uint16(anon_sym_AMP_AMP),
	836:   uint16(anon_sym_BANG_EQ),
	837:   uint16(anon_sym_LT_EQ),
	838:   uint16(anon_sym_GT_EQ),
	839:   uint16(anon_sym_DOLLAR_LPAREN),
	840:   uint16(anon_sym_DQUOTE),
	841:   uint16(anon_sym_SQUOTE),
	842:   uint16(198),
	843:   uint16(36),
	844:   uint16(anon_sym_mainmenu),
	845:   uint16(anon_sym_config),
	846:   uint16(anon_sym_configdefault),
	847:   uint16(anon_sym_menuconfig),
	848:   uint16(anon_sym_choice),
	849:   uint16(anon_sym_endchoice),
	850:   uint16(anon_sym_comment),
	851:   uint16(anon_sym_menu),
	852:   uint16(anon_sym_endmenu),
	853:   uint16(anon_sym_if),
	854:   uint16(anon_sym_endif),
	855:   uint16(anon_sym_source),
	856:   uint16(anon_sym_rsource),
	857:   uint16(anon_sym_osource),
	858:   uint16(anon_sym_orsource),
	859:   uint16(anon_sym_bool),
	860:   uint16(anon_sym_tristate),
	861:   uint16(anon_sym_int),
	862:   uint16(anon_sym_hex),
	863:   uint16(anon_sym_string),
	864:   uint16(anon_sym_prompt),
	865:   uint16(anon_sym_default),
	866:   uint16(anon_sym_def_bool),
	867:   uint16(anon_sym_def_tristate),
	868:   uint16(anon_sym_def_int),
	869:   uint16(anon_sym_def_hex),
	870:   uint16(anon_sym_def_string),
	871:   uint16(anon_sym_select),
	872:   uint16(anon_sym_imply),
	873:   uint16(anon_sym_range),
	874:   uint16(anon_sym_help),
	875:   uint16(sym_optional),
	876:   uint16(sym_modules),
	877:   uint16(anon_sym_LT),
	878:   uint16(anon_sym_GT),
	879:   uint16(sym_symbol),
	880:   uint16(3),
	881:   uint16(3),
	882:   uint16(1),
	883:   uint16(sym_comment),
	884:   uint16(204),
	885:   uint16(11),
	886:   uint16(anon_sym_EQ),
	887:   uint16(anon_sym_dependson),
	888:   uint16(anon_sym_visibleif),
	889:   uint16(anon_sym_PIPE_PIPE),
	890:   uint16(anon_sym_AMP_AMP),
	891:   uint16(anon_sym_BANG_EQ),
	892:   uint16(anon_sym_LT_EQ),
	893:   uint16(anon_sym_GT_EQ),
	894:   uint16(anon_sym_DOLLAR_LPAREN),
	895:   uint16(anon_sym_DQUOTE),
	896:   uint16(anon_sym_SQUOTE),
	897:   uint16(202),
	898:   uint16(36),
	899:   uint16(anon_sym_mainmenu),
	900:   uint16(anon_sym_config),
	901:   uint16(anon_sym_configdefault),
	902:   uint16(anon_sym_menuconfig),
	903:   uint16(anon_sym_choice),
	904:   uint16(anon_sym_endchoice),
	905:   uint16(anon_sym_comment),
	906:   uint16(anon_sym_menu),
	907:   uint16(anon_sym_endmenu),
	908:   uint16(anon_sym_if),
	909:   uint16(anon_sym_endif),
	910:   uint16(anon_sym_source),
	911:   uint16(anon_sym_rsource),
	912:   uint16(anon_sym_osource),
	913:   uint16(anon_sym_orsource),
	914:   uint16(anon_sym_bool),
	915:   uint16(anon_sym_tristate),
	916:   uint16(anon_sym_int),
	917:   uint16(anon_sym_hex),
	918:   uint16(anon_sym_string),
	919:   uint16(anon_sym_prompt),
	920:   uint16(anon_sym_default),
	921:   uint16(anon_sym_def_bool),
	922:   uint16(anon_sym_def_tristate),
	923:   uint16(anon_sym_def_int),
	924:   uint16(anon_sym_def_hex),
	925:   uint16(anon_sym_def_string),
	926:   uint16(anon_sym_select),
	927:   uint16(anon_sym_imply),
	928:   uint16(anon_sym_range),
	929:   uint16(anon_sym_help),
	930:   uint16(sym_optional),
	931:   uint16(sym_modules),
	932:   uint16(anon_sym_LT),
	933:   uint16(anon_sym_GT),
	934:   uint16(sym_symbol),
	935:   uint16(15),
	936:   uint16(3),
	937:   uint16(1),
	938:   uint16(sym_comment),
	939:   uint16(206),
	940:   uint16(1),
	942:   uint16(211),
	943:   uint16(1),
	944:   uint16(anon_sym_prompt),
	945:   uint16(214),
	946:   uint16(1),
	947:   uint16(anon_sym_default),
	948:   uint16(220),
	949:   uint16(1),
	950:   uint16(anon_sym_dependson),
	951:   uint16(223),
	952:   uint16(1),
	953:   uint16(anon_sym_select),
	954:   uint16(226),
	955:   uint16(1),
	956:   uint16(anon_sym_imply),
	957:   uint16(229),
	958:   uint16(1),
	959:   uint16(anon_sym_visibleif),
	960:   uint16(232),
	961:   uint16(1),
	962:   uint16(anon_sym_range),
	963:   uint16(235),
	964:   uint16(1),
	965:   uint16(anon_sym_help),
	966:   uint16(238),
	967:   uint16(2),
	968:   uint16(sym_optional),
	969:   uint16(sym_modules),
	970:   uint16(208),
	971:   uint16(5),
	972:   uint16(anon_sym_bool),
	973:   uint16(anon_sym_tristate),
	974:   uint16(anon_sym_int),
	975:   uint16(anon_sym_hex),
	976:   uint16(anon_sym_string),
	977:   uint16(217),
	978:   uint16(5),
	979:   uint16(anon_sym_def_bool),
	980:   uint16(anon_sym_def_tristate),
	981:   uint16(anon_sym_def_int),
	982:   uint16(anon_sym_def_hex),
	983:   uint16(anon_sym_def_string),
	984:   uint16(24),
	985:   uint16(12),
	986:   uint16(sym__config_option),
	987:   uint16(sym_type_definition),
	988:   uint16(sym_input_prompt),
	989:   uint16(sym_default_value),
	990:   uint16(sym_type_definition_default),
	991:   uint16(sym_dependencies),
	992:   uint16(sym_reverse_dependencies),
	993:   uint16(sym_weak_reverse_dependencies),
	994:   uint16(sym_limiting_menu_display),
	995:   uint16(sym_numerical_ranges),
	996:   uint16(sym_help_text),
	997:   uint16(aux_sym_config_repeat1),
	998:   uint16(117),
	999:   uint16(13),
	1000:  uint16(anon_sym_mainmenu),
	1001:  uint16(anon_sym_config),
	1002:  uint16(anon_sym_configdefault),
	1003:  uint16(anon_sym_menuconfig),
	1004:  uint16(anon_sym_choice),
	1005:  uint16(anon_sym_comment),
	1006:  uint16(anon_sym_menu),
	1007:  uint16(anon_sym_if),
	1008:  uint16(anon_sym_source),
	1009:  uint16(anon_sym_rsource),
	1010:  uint16(anon_sym_osource),
	1011:  uint16(anon_sym_orsource),
	1012:  uint16(sym_symbol),
	1013:  uint16(15),
	1014:  uint16(3),
	1015:  uint16(1),
	1016:  uint16(sym_comment),
	1017:  uint16(241),
	1018:  uint16(1),
	1020:  uint16(245),
	1021:  uint16(1),
	1022:  uint16(anon_sym_prompt),
	1023:  uint16(247),
	1024:  uint16(1),
	1025:  uint16(anon_sym_default),
	1026:  uint16(251),
	1027:  uint16(1),
	1028:  uint16(anon_sym_dependson),
	1029:  uint16(253),
	1030:  uint16(1),
	1031:  uint16(anon_sym_select),
	1032:  uint16(255),
	1033:  uint16(1),
	1034:  uint16(anon_sym_imply),
	1035:  uint16(257),
	1036:  uint16(1),
	1037:  uint16(anon_sym_visibleif),
	1038:  uint16(259),
	1039:  uint16(1),
	1040:  uint16(anon_sym_range),
	1041:  uint16(261),
	1042:  uint16(1),
	1043:  uint16(anon_sym_help),
	1044:  uint16(263),
	1045:  uint16(2),
	1046:  uint16(sym_optional),
	1047:  uint16(sym_modules),
	1048:  uint16(243),
	1049:  uint16(5),
	1050:  uint16(anon_sym_bool),
	1051:  uint16(anon_sym_tristate),
	1052:  uint16(anon_sym_int),
	1053:  uint16(anon_sym_hex),
	1054:  uint16(anon_sym_string),
	1055:  uint16(249),
	1056:  uint16(5),
	1057:  uint16(anon_sym_def_bool),
	1058:  uint16(anon_sym_def_tristate),
	1059:  uint16(anon_sym_def_int),
	1060:  uint16(anon_sym_def_hex),
	1061:  uint16(anon_sym_def_string),
	1062:  uint16(24),
	1063:  uint16(12),
	1064:  uint16(sym__config_option),
	1065:  uint16(sym_type_definition),
	1066:  uint16(sym_input_prompt),
	1067:  uint16(sym_default_value),
	1068:  uint16(sym_type_definition_default),
	1069:  uint16(sym_dependencies),
	1070:  uint16(sym_reverse_dependencies),
	1071:  uint16(sym_weak_reverse_dependencies),
	1072:  uint16(sym_limiting_menu_display),
	1073:  uint16(sym_numerical_ranges),
	1074:  uint16(sym_help_text),
	1075:  uint16(aux_sym_config_repeat1),
	1076:  uint16(156),
	1077:  uint16(13),
	1078:  uint16(anon_sym_mainmenu),
	1079:  uint16(anon_sym_config),
	1080:  uint16(anon_sym_configdefault),
	1081:  uint16(anon_sym_menuconfig),
	1082:  uint16(anon_sym_choice),
	1083:  uint16(anon_sym_comment),
	1084:  uint16(anon_sym_menu),
	1085:  uint16(anon_sym_if),
	1086:  uint16(anon_sym_source),
	1087:  uint16(anon_sym_rsource),
	1088:  uint16(anon_sym_osource),
	1089:  uint16(anon_sym_orsource),
	1090:  uint16(sym_symbol),
	1091:  uint16(15),
	1092:  uint16(3),
	1093:  uint16(1),
	1094:  uint16(sym_comment),
	1095:  uint16(245),
	1096:  uint16(1),
	1097:  uint16(anon_sym_prompt),
	1098:  uint16(247),
	1099:  uint16(1),
	1100:  uint16(anon_sym_default),
	1101:  uint16(251),
	1102:  uint16(1),
	1103:  uint16(anon_sym_dependson),
	1104:  uint16(253),
	1105:  uint16(1),
	1106:  uint16(anon_sym_select),
	1107:  uint16(255),
	1108:  uint16(1),
	1109:  uint16(anon_sym_imply),
	1110:  uint16(257),
	1111:  uint16(1),
	1112:  uint16(anon_sym_visibleif),
	1113:  uint16(259),
	1114:  uint16(1),
	1115:  uint16(anon_sym_range),
	1116:  uint16(261),
	1117:  uint16(1),
	1118:  uint16(anon_sym_help),
	1119:  uint16(265),
	1120:  uint16(1),
	1122:  uint16(263),
	1123:  uint16(2),
	1124:  uint16(sym_optional),
	1125:  uint16(sym_modules),
	1126:  uint16(243),
	1127:  uint16(5),
	1128:  uint16(anon_sym_bool),
	1129:  uint16(anon_sym_tristate),
	1130:  uint16(anon_sym_int),
	1131:  uint16(anon_sym_hex),
	1132:  uint16(anon_sym_string),
	1133:  uint16(249),
	1134:  uint16(5),
	1135:  uint16(anon_sym_def_bool),
	1136:  uint16(anon_sym_def_tristate),
	1137:  uint16(anon_sym_def_int),
	1138:  uint16(anon_sym_def_hex),
	1139:  uint16(anon_sym_def_string),
	1140:  uint16(24),
	1141:  uint16(12),
	1142:  uint16(sym__config_option),
	1143:  uint16(sym_type_definition),
	1144:  uint16(sym_input_prompt),
	1145:  uint16(sym_default_value),
	1146:  uint16(sym_type_definition_default),
	1147:  uint16(sym_dependencies),
	1148:  uint16(sym_reverse_dependencies),
	1149:  uint16(sym_weak_reverse_dependencies),
	1150:  uint16(sym_limiting_menu_display),
	1151:  uint16(sym_numerical_ranges),
	1152:  uint16(sym_help_text),
	1153:  uint16(aux_sym_config_repeat1),
	1154:  uint16(158),
	1155:  uint16(13),
	1156:  uint16(anon_sym_mainmenu),
	1157:  uint16(anon_sym_config),
	1158:  uint16(anon_sym_configdefault),
	1159:  uint16(anon_sym_menuconfig),
	1160:  uint16(anon_sym_choice),
	1161:  uint16(anon_sym_comment),
	1162:  uint16(anon_sym_menu),
	1163:  uint16(anon_sym_if),
	1164:  uint16(anon_sym_source),
	1165:  uint16(anon_sym_rsource),
	1166:  uint16(anon_sym_osource),
	1167:  uint16(anon_sym_orsource),
	1168:  uint16(sym_symbol),
	1169:  uint16(15),
	1170:  uint16(3),
	1171:  uint16(1),
	1172:  uint16(sym_comment),
	1173:  uint16(245),
	1174:  uint16(1),
	1175:  uint16(anon_sym_prompt),
	1176:  uint16(247),
	1177:  uint16(1),
	1178:  uint16(anon_sym_default),
	1179:  uint16(251),
	1180:  uint16(1),
	1181:  uint16(anon_sym_dependson),
	1182:  uint16(253),
	1183:  uint16(1),
	1184:  uint16(anon_sym_select),
	1185:  uint16(255),
	1186:  uint16(1),
	1187:  uint16(anon_sym_imply),
	1188:  uint16(257),
	1189:  uint16(1),
	1190:  uint16(anon_sym_visibleif),
	1191:  uint16(259),
	1192:  uint16(1),
	1193:  uint16(anon_sym_range),
	1194:  uint16(261),
	1195:  uint16(1),
	1196:  uint16(anon_sym_help),
	1197:  uint16(267),
	1198:  uint16(1),
	1200:  uint16(263),
	1201:  uint16(2),
	1202:  uint16(sym_optional),
	1203:  uint16(sym_modules),
	1204:  uint16(243),
	1205:  uint16(5),
	1206:  uint16(anon_sym_bool),
	1207:  uint16(anon_sym_tristate),
	1208:  uint16(anon_sym_int),
	1209:  uint16(anon_sym_hex),
	1210:  uint16(anon_sym_string),
	1211:  uint16(249),
	1212:  uint16(5),
	1213:  uint16(anon_sym_def_bool),
	1214:  uint16(anon_sym_def_tristate),
	1215:  uint16(anon_sym_def_int),
	1216:  uint16(anon_sym_def_hex),
	1217:  uint16(anon_sym_def_string),
	1218:  uint16(24),
	1219:  uint16(12),
	1220:  uint16(sym__config_option),
	1221:  uint16(sym_type_definition),
	1222:  uint16(sym_input_prompt),
	1223:  uint16(sym_default_value),
	1224:  uint16(sym_type_definition_default),
	1225:  uint16(sym_dependencies),
	1226:  uint16(sym_reverse_dependencies),
	1227:  uint16(sym_weak_reverse_dependencies),
	1228:  uint16(sym_limiting_menu_display),
	1229:  uint16(sym_numerical_ranges),
	1230:  uint16(sym_help_text),
	1231:  uint16(aux_sym_config_repeat1),
	1232:  uint16(172),
	1233:  uint16(13),
	1234:  uint16(anon_sym_mainmenu),
	1235:  uint16(anon_sym_config),
	1236:  uint16(anon_sym_configdefault),
	1237:  uint16(anon_sym_menuconfig),
	1238:  uint16(anon_sym_choice),
	1239:  uint16(anon_sym_comment),
	1240:  uint16(anon_sym_menu),
	1241:  uint16(anon_sym_if),
	1242:  uint16(anon_sym_source),
	1243:  uint16(anon_sym_rsource),
	1244:  uint16(anon_sym_osource),
	1245:  uint16(anon_sym_orsource),
	1246:  uint16(sym_symbol),
	1247:  uint16(15),
	1248:  uint16(3),
	1249:  uint16(1),
	1250:  uint16(sym_comment),
	1251:  uint16(245),
	1252:  uint16(1),
	1253:  uint16(anon_sym_prompt),
	1254:  uint16(247),
	1255:  uint16(1),
	1256:  uint16(anon_sym_default),
	1257:  uint16(251),
	1258:  uint16(1),
	1259:  uint16(anon_sym_dependson),
	1260:  uint16(253),
	1261:  uint16(1),
	1262:  uint16(anon_sym_select),
	1263:  uint16(255),
	1264:  uint16(1),
	1265:  uint16(anon_sym_imply),
	1266:  uint16(257),
	1267:  uint16(1),
	1268:  uint16(anon_sym_visibleif),
	1269:  uint16(259),
	1270:  uint16(1),
	1271:  uint16(anon_sym_range),
	1272:  uint16(261),
	1273:  uint16(1),
	1274:  uint16(anon_sym_help),
	1275:  uint16(269),
	1276:  uint16(1),
	1278:  uint16(271),
	1279:  uint16(2),
	1280:  uint16(sym_optional),
	1281:  uint16(sym_modules),
	1282:  uint16(243),
	1283:  uint16(5),
	1284:  uint16(anon_sym_bool),
	1285:  uint16(anon_sym_tristate),
	1286:  uint16(anon_sym_int),
	1287:  uint16(anon_sym_hex),
	1288:  uint16(anon_sym_string),
	1289:  uint16(249),
	1290:  uint16(5),
	1291:  uint16(anon_sym_def_bool),
	1292:  uint16(anon_sym_def_tristate),
	1293:  uint16(anon_sym_def_int),
	1294:  uint16(anon_sym_def_hex),
	1295:  uint16(anon_sym_def_string),
	1296:  uint16(27),
	1297:  uint16(12),
	1298:  uint16(sym__config_option),
	1299:  uint16(sym_type_definition),
	1300:  uint16(sym_input_prompt),
	1301:  uint16(sym_default_value),
	1302:  uint16(sym_type_definition_default),
	1303:  uint16(sym_dependencies),
	1304:  uint16(sym_reverse_dependencies),
	1305:  uint16(sym_weak_reverse_dependencies),
	1306:  uint16(sym_limiting_menu_display),
	1307:  uint16(sym_numerical_ranges),
	1308:  uint16(sym_help_text),
	1309:  uint16(aux_sym_config_repeat1),
	1310:  uint16(152),
	1311:  uint16(13),
	1312:  uint16(anon_sym_mainmenu),
	1313:  uint16(anon_sym_config),
	1314:  uint16(anon_sym_configdefault),
	1315:  uint16(anon_sym_menuconfig),
	1316:  uint16(anon_sym_choice),
	1317:  uint16(anon_sym_comment),
	1318:  uint16(anon_sym_menu),
	1319:  uint16(anon_sym_if),
	1320:  uint16(anon_sym_source),
	1321:  uint16(anon_sym_rsource),
	1322:  uint16(anon_sym_osource),
	1323:  uint16(anon_sym_orsource),
	1324:  uint16(sym_symbol),
	1325:  uint16(3),
	1326:  uint16(3),
	1327:  uint16(1),
	1328:  uint16(sym_comment),
	1329:  uint16(200),
	1330:  uint16(12),
	1332:  uint16(anon_sym_EQ),
	1333:  uint16(anon_sym_dependson),
	1334:  uint16(anon_sym_visibleif),
	1335:  uint16(anon_sym_PIPE_PIPE),
	1336:  uint16(anon_sym_AMP_AMP),
	1337:  uint16(anon_sym_BANG_EQ),
	1338:  uint16(anon_sym_LT_EQ),
	1339:  uint16(anon_sym_GT_EQ),
	1340:  uint16(anon_sym_DOLLAR_LPAREN),
	1341:  uint16(anon_sym_DQUOTE),
	1342:  uint16(anon_sym_SQUOTE),
	1343:  uint16(198),
	1344:  uint16(33),
	1345:  uint16(anon_sym_mainmenu),
	1346:  uint16(anon_sym_config),
	1347:  uint16(anon_sym_configdefault),
	1348:  uint16(anon_sym_menuconfig),
	1349:  uint16(anon_sym_choice),
	1350:  uint16(anon_sym_comment),
	1351:  uint16(anon_sym_menu),
	1352:  uint16(anon_sym_if),
	1353:  uint16(anon_sym_source),
	1354:  uint16(anon_sym_rsource),
	1355:  uint16(anon_sym_osource),
	1356:  uint16(anon_sym_orsource),
	1357:  uint16(anon_sym_bool),
	1358:  uint16(anon_sym_tristate),
	1359:  uint16(anon_sym_int),
	1360:  uint16(anon_sym_hex),
	1361:  uint16(anon_sym_string),
	1362:  uint16(anon_sym_prompt),
	1363:  uint16(anon_sym_default),
	1364:  uint16(anon_sym_def_bool),
	1365:  uint16(anon_sym_def_tristate),
	1366:  uint16(anon_sym_def_int),
	1367:  uint16(anon_sym_def_hex),
	1368:  uint16(anon_sym_def_string),
	1369:  uint16(anon_sym_select),
	1370:  uint16(anon_sym_imply),
	1371:  uint16(anon_sym_range),
	1372:  uint16(anon_sym_help),
	1373:  uint16(sym_optional),
	1374:  uint16(sym_modules),
	1375:  uint16(anon_sym_LT),
	1376:  uint16(anon_sym_GT),
	1377:  uint16(sym_symbol),
	1378:  uint16(3),
	1379:  uint16(3),
	1380:  uint16(1),
	1381:  uint16(sym_comment),
	1382:  uint16(184),
	1383:  uint16(12),
	1385:  uint16(anon_sym_EQ),
	1386:  uint16(anon_sym_dependson),
	1387:  uint16(anon_sym_visibleif),
	1388:  uint16(anon_sym_PIPE_PIPE),
	1389:  uint16(anon_sym_AMP_AMP),
	1390:  uint16(anon_sym_BANG_EQ),
	1391:  uint16(anon_sym_LT_EQ),
	1392:  uint16(anon_sym_GT_EQ),
	1393:  uint16(anon_sym_DOLLAR_LPAREN),
	1394:  uint16(anon_sym_DQUOTE),
	1395:  uint16(anon_sym_SQUOTE),
	1396:  uint16(182),
	1397:  uint16(33),
	1398:  uint16(anon_sym_mainmenu),
	1399:  uint16(anon_sym_config),
	1400:  uint16(anon_sym_configdefault),
	1401:  uint16(anon_sym_menuconfig),
	1402:  uint16(anon_sym_choice),
	1403:  uint16(anon_sym_comment),
	1404:  uint16(anon_sym_menu),
	1405:  uint16(anon_sym_if),
	1406:  uint16(anon_sym_source),
	1407:  uint16(anon_sym_rsource),
	1408:  uint16(anon_sym_osource),
	1409:  uint16(anon_sym_orsource),
	1410:  uint16(anon_sym_bool),
	1411:  uint16(anon_sym_tristate),
	1412:  uint16(anon_sym_int),
	1413:  uint16(anon_sym_hex),
	1414:  uint16(anon_sym_string),
	1415:  uint16(anon_sym_prompt),
	1416:  uint16(anon_sym_default),
	1417:  uint16(anon_sym_def_bool),
	1418:  uint16(anon_sym_def_tristate),
	1419:  uint16(anon_sym_def_int),
	1420:  uint16(anon_sym_def_hex),
	1421:  uint16(anon_sym_def_string),
	1422:  uint16(anon_sym_select),
	1423:  uint16(anon_sym_imply),
	1424:  uint16(anon_sym_range),
	1425:  uint16(anon_sym_help),
	1426:  uint16(sym_optional),
	1427:  uint16(sym_modules),
	1428:  uint16(anon_sym_LT),
	1429:  uint16(anon_sym_GT),
	1430:  uint16(sym_symbol),
	1431:  uint16(4),
	1432:  uint16(3),
	1433:  uint16(1),
	1434:  uint16(sym_comment),
	1435:  uint16(196),
	1436:  uint16(3),
	1437:  uint16(anon_sym_DOLLAR_LPAREN),
	1438:  uint16(anon_sym_DQUOTE),
	1439:  uint16(anon_sym_SQUOTE),
	1440:  uint16(193),
	1441:  uint16(9),
	1443:  uint16(anon_sym_EQ),
	1444:  uint16(anon_sym_dependson),
	1445:  uint16(anon_sym_visibleif),
	1446:  uint16(anon_sym_PIPE_PIPE),
	1447:  uint16(anon_sym_AMP_AMP),
	1448:  uint16(anon_sym_BANG_EQ),
	1449:  uint16(anon_sym_LT_EQ),
	1450:  uint16(anon_sym_GT_EQ),
	1451:  uint16(190),
	1452:  uint16(33),
	1453:  uint16(anon_sym_mainmenu),
	1454:  uint16(anon_sym_config),
	1455:  uint16(anon_sym_configdefault),
	1456:  uint16(anon_sym_menuconfig),
	1457:  uint16(anon_sym_choice),
	1458:  uint16(anon_sym_comment),
	1459:  uint16(anon_sym_menu),
	1460:  uint16(anon_sym_if),
	1461:  uint16(anon_sym_source),
	1462:  uint16(anon_sym_rsource),
	1463:  uint16(anon_sym_osource),
	1464:  uint16(anon_sym_orsource),
	1465:  uint16(anon_sym_bool),
	1466:  uint16(anon_sym_tristate),
	1467:  uint16(anon_sym_int),
	1468:  uint16(anon_sym_hex),
	1469:  uint16(anon_sym_string),
	1470:  uint16(anon_sym_prompt),
	1471:  uint16(anon_sym_default),
	1472:  uint16(anon_sym_def_bool),
	1473:  uint16(anon_sym_def_tristate),
	1474:  uint16(anon_sym_def_int),
	1475:  uint16(anon_sym_def_hex),
	1476:  uint16(anon_sym_def_string),
	1477:  uint16(anon_sym_select),
	1478:  uint16(anon_sym_imply),
	1479:  uint16(anon_sym_range),
	1480:  uint16(anon_sym_help),
	1481:  uint16(sym_optional),
	1482:  uint16(sym_modules),
	1483:  uint16(anon_sym_LT),
	1484:  uint16(anon_sym_GT),
	1485:  uint16(sym_symbol),
	1486:  uint16(3),
	1487:  uint16(3),
	1488:  uint16(1),
	1489:  uint16(sym_comment),
	1490:  uint16(204),
	1491:  uint16(12),
	1493:  uint16(anon_sym_EQ),
	1494:  uint16(anon_sym_dependson),
	1495:  uint16(anon_sym_visibleif),
	1496:  uint16(anon_sym_PIPE_PIPE),
	1497:  uint16(anon_sym_AMP_AMP),
	1498:  uint16(anon_sym_BANG_EQ),
	1499:  uint16(anon_sym_LT_EQ),
	1500:  uint16(anon_sym_GT_EQ),
	1501:  uint16(anon_sym_DOLLAR_LPAREN),
	1502:  uint16(anon_sym_DQUOTE),
	1503:  uint16(anon_sym_SQUOTE),
	1504:  uint16(202),
	1505:  uint16(33),
	1506:  uint16(anon_sym_mainmenu),
	1507:  uint16(anon_sym_config),
	1508:  uint16(anon_sym_configdefault),
	1509:  uint16(anon_sym_menuconfig),
	1510:  uint16(anon_sym_choice),
	1511:  uint16(anon_sym_comment),
	1512:  uint16(anon_sym_menu),
	1513:  uint16(anon_sym_if),
	1514:  uint16(anon_sym_source),
	1515:  uint16(anon_sym_rsource),
	1516:  uint16(anon_sym_osource),
	1517:  uint16(anon_sym_orsource),
	1518:  uint16(anon_sym_bool),
	1519:  uint16(anon_sym_tristate),
	1520:  uint16(anon_sym_int),
	1521:  uint16(anon_sym_hex),
	1522:  uint16(anon_sym_string),
	1523:  uint16(anon_sym_prompt),
	1524:  uint16(anon_sym_default),
	1525:  uint16(anon_sym_def_bool),
	1526:  uint16(anon_sym_def_tristate),
	1527:  uint16(anon_sym_def_int),
	1528:  uint16(anon_sym_def_hex),
	1529:  uint16(anon_sym_def_string),
	1530:  uint16(anon_sym_select),
	1531:  uint16(anon_sym_imply),
	1532:  uint16(anon_sym_range),
	1533:  uint16(anon_sym_help),
	1534:  uint16(sym_optional),
	1535:  uint16(sym_modules),
	1536:  uint16(anon_sym_LT),
	1537:  uint16(anon_sym_GT),
	1538:  uint16(sym_symbol),
	1539:  uint16(3),
	1540:  uint16(3),
	1541:  uint16(1),
	1542:  uint16(sym_comment),
	1543:  uint16(188),
	1544:  uint16(12),
	1546:  uint16(anon_sym_EQ),
	1547:  uint16(anon_sym_dependson),
	1548:  uint16(anon_sym_visibleif),
	1549:  uint16(anon_sym_PIPE_PIPE),
	1550:  uint16(anon_sym_AMP_AMP),
	1551:  uint16(anon_sym_BANG_EQ),
	1552:  uint16(anon_sym_LT_EQ),
	1553:  uint16(anon_sym_GT_EQ),
	1554:  uint16(anon_sym_DOLLAR_LPAREN),
	1555:  uint16(anon_sym_DQUOTE),
	1556:  uint16(anon_sym_SQUOTE),
	1557:  uint16(186),
	1558:  uint16(33),
	1559:  uint16(anon_sym_mainmenu),
	1560:  uint16(anon_sym_config),
	1561:  uint16(anon_sym_configdefault),
	1562:  uint16(anon_sym_menuconfig),
	1563:  uint16(anon_sym_choice),
	1564:  uint16(anon_sym_comment),
	1565:  uint16(anon_sym_menu),
	1566:  uint16(anon_sym_if),
	1567:  uint16(anon_sym_source),
	1568:  uint16(anon_sym_rsource),
	1569:  uint16(anon_sym_osource),
	1570:  uint16(anon_sym_orsource),
	1571:  uint16(anon_sym_bool),
	1572:  uint16(anon_sym_tristate),
	1573:  uint16(anon_sym_int),
	1574:  uint16(anon_sym_hex),
	1575:  uint16(anon_sym_string),
	1576:  uint16(anon_sym_prompt),
	1577:  uint16(anon_sym_default),
	1578:  uint16(anon_sym_def_bool),
	1579:  uint16(anon_sym_def_tristate),
	1580:  uint16(anon_sym_def_int),
	1581:  uint16(anon_sym_def_hex),
	1582:  uint16(anon_sym_def_string),
	1583:  uint16(anon_sym_select),
	1584:  uint16(anon_sym_imply),
	1585:  uint16(anon_sym_range),
	1586:  uint16(anon_sym_help),
	1587:  uint16(sym_optional),
	1588:  uint16(sym_modules),
	1589:  uint16(anon_sym_LT),
	1590:  uint16(anon_sym_GT),
	1591:  uint16(sym_symbol),
	1592:  uint16(3),
	1593:  uint16(3),
	1594:  uint16(1),
	1595:  uint16(sym_comment),
	1596:  uint16(275),
	1597:  uint16(8),
	1598:  uint16(anon_sym_EQ),
	1599:  uint16(anon_sym_dependson),
	1600:  uint16(anon_sym_visibleif),
	1601:  uint16(anon_sym_PIPE_PIPE),
	1602:  uint16(anon_sym_AMP_AMP),
	1603:  uint16(anon_sym_BANG_EQ),
	1604:  uint16(anon_sym_LT_EQ),
	1605:  uint16(anon_sym_GT_EQ),
	1606:  uint16(273),
	1607:  uint16(36),
	1608:  uint16(anon_sym_mainmenu),
	1609:  uint16(anon_sym_config),
	1610:  uint16(anon_sym_configdefault),
	1611:  uint16(anon_sym_menuconfig),
	1612:  uint16(anon_sym_choice),
	1613:  uint16(anon_sym_endchoice),
	1614:  uint16(anon_sym_comment),
	1615:  uint16(anon_sym_menu),
	1616:  uint16(anon_sym_endmenu),
	1617:  uint16(anon_sym_if),
	1618:  uint16(anon_sym_endif),
	1619:  uint16(anon_sym_source),
	1620:  uint16(anon_sym_rsource),
	1621:  uint16(anon_sym_osource),
	1622:  uint16(anon_sym_orsource),
	1623:  uint16(anon_sym_bool),
	1624:  uint16(anon_sym_tristate),
	1625:  uint16(anon_sym_int),
	1626:  uint16(anon_sym_hex),
	1627:  uint16(anon_sym_string),
	1628:  uint16(anon_sym_prompt),
	1629:  uint16(anon_sym_default),
	1630:  uint16(anon_sym_def_bool),
	1631:  uint16(anon_sym_def_tristate),
	1632:  uint16(anon_sym_def_int),
	1633:  uint16(anon_sym_def_hex),
	1634:  uint16(anon_sym_def_string),
	1635:  uint16(anon_sym_select),
	1636:  uint16(anon_sym_imply),
	1637:  uint16(anon_sym_range),
	1638:  uint16(anon_sym_help),
	1639:  uint16(sym_optional),
	1640:  uint16(sym_modules),
	1641:  uint16(anon_sym_LT),
	1642:  uint16(anon_sym_GT),
	1643:  uint16(sym_symbol),
	1644:  uint16(3),
	1645:  uint16(3),
	1646:  uint16(1),
	1647:  uint16(sym_comment),
	1648:  uint16(279),
	1649:  uint16(8),
	1650:  uint16(anon_sym_EQ),
	1651:  uint16(anon_sym_dependson),
	1652:  uint16(anon_sym_visibleif),
	1653:  uint16(anon_sym_PIPE_PIPE),
	1654:  uint16(anon_sym_AMP_AMP),
	1655:  uint16(anon_sym_BANG_EQ),
	1656:  uint16(anon_sym_LT_EQ),
	1657:  uint16(anon_sym_GT_EQ),
	1658:  uint16(277),
	1659:  uint16(36),
	1660:  uint16(anon_sym_mainmenu),
	1661:  uint16(anon_sym_config),
	1662:  uint16(anon_sym_configdefault),
	1663:  uint16(anon_sym_menuconfig),
	1664:  uint16(anon_sym_choice),
	1665:  uint16(anon_sym_endchoice),
	1666:  uint16(anon_sym_comment),
	1667:  uint16(anon_sym_menu),
	1668:  uint16(anon_sym_endmenu),
	1669:  uint16(anon_sym_if),
	1670:  uint16(anon_sym_endif),
	1671:  uint16(anon_sym_source),
	1672:  uint16(anon_sym_rsource),
	1673:  uint16(anon_sym_osource),
	1674:  uint16(anon_sym_orsource),
	1675:  uint16(anon_sym_bool),
	1676:  uint16(anon_sym_tristate),
	1677:  uint16(anon_sym_int),
	1678:  uint16(anon_sym_hex),
	1679:  uint16(anon_sym_string),
	1680:  uint16(anon_sym_prompt),
	1681:  uint16(anon_sym_default),
	1682:  uint16(anon_sym_def_bool),
	1683:  uint16(anon_sym_def_tristate),
	1684:  uint16(anon_sym_def_int),
	1685:  uint16(anon_sym_def_hex),
	1686:  uint16(anon_sym_def_string),
	1687:  uint16(anon_sym_select),
	1688:  uint16(anon_sym_imply),
	1689:  uint16(anon_sym_range),
	1690:  uint16(anon_sym_help),
	1691:  uint16(sym_optional),
	1692:  uint16(sym_modules),
	1693:  uint16(anon_sym_LT),
	1694:  uint16(anon_sym_GT),
	1695:  uint16(sym_symbol),
	1696:  uint16(6),
	1697:  uint16(3),
	1698:  uint16(1),
	1699:  uint16(sym_comment),
	1700:  uint16(281),
	1701:  uint16(1),
	1702:  uint16(anon_sym_EQ),
	1703:  uint16(285),
	1704:  uint16(2),
	1705:  uint16(anon_sym_LT),
	1706:  uint16(anon_sym_GT),
	1707:  uint16(283),
	1708:  uint16(3),
	1709:  uint16(anon_sym_BANG_EQ),
	1710:  uint16(anon_sym_LT_EQ),
	1711:  uint16(anon_sym_GT_EQ),
	1712:  uint16(279),
	1713:  uint16(4),
	1714:  uint16(anon_sym_dependson),
	1715:  uint16(anon_sym_visibleif),
	1716:  uint16(anon_sym_PIPE_PIPE),
	1717:  uint16(anon_sym_AMP_AMP),
	1718:  uint16(277),
	1719:  uint16(34),
	1720:  uint16(anon_sym_mainmenu),
	1721:  uint16(anon_sym_config),
	1722:  uint16(anon_sym_configdefault),
	1723:  uint16(anon_sym_menuconfig),
	1724:  uint16(anon_sym_choice),
	1725:  uint16(anon_sym_endchoice),
	1726:  uint16(anon_sym_comment),
	1727:  uint16(anon_sym_menu),
	1728:  uint16(anon_sym_endmenu),
	1729:  uint16(anon_sym_if),
	1730:  uint16(anon_sym_endif),
	1731:  uint16(anon_sym_source),
	1732:  uint16(anon_sym_rsource),
	1733:  uint16(anon_sym_osource),
	1734:  uint16(anon_sym_orsource),
	1735:  uint16(anon_sym_bool),
	1736:  uint16(anon_sym_tristate),
	1737:  uint16(anon_sym_int),
	1738:  uint16(anon_sym_hex),
	1739:  uint16(anon_sym_string),
	1740:  uint16(anon_sym_prompt),
	1741:  uint16(anon_sym_default),
	1742:  uint16(anon_sym_def_bool),
	1743:  uint16(anon_sym_def_tristate),
	1744:  uint16(anon_sym_def_int),
	1745:  uint16(anon_sym_def_hex),
	1746:  uint16(anon_sym_def_string),
	1747:  uint16(anon_sym_select),
	1748:  uint16(anon_sym_imply),
	1749:  uint16(anon_sym_range),
	1750:  uint16(anon_sym_help),
	1751:  uint16(sym_optional),
	1752:  uint16(sym_modules),
	1753:  uint16(sym_symbol),
	1754:  uint16(3),
	1755:  uint16(3),
	1756:  uint16(1),
	1757:  uint16(sym_comment),
	1758:  uint16(289),
	1759:  uint16(8),
	1760:  uint16(anon_sym_EQ),
	1761:  uint16(anon_sym_dependson),
	1762:  uint16(anon_sym_visibleif),
	1763:  uint16(anon_sym_PIPE_PIPE),
	1764:  uint16(anon_sym_AMP_AMP),
	1765:  uint16(anon_sym_BANG_EQ),
	1766:  uint16(anon_sym_LT_EQ),
	1767:  uint16(anon_sym_GT_EQ),
	1768:  uint16(287),
	1769:  uint16(36),
	1770:  uint16(anon_sym_mainmenu),
	1771:  uint16(anon_sym_config),
	1772:  uint16(anon_sym_configdefault),
	1773:  uint16(anon_sym_menuconfig),
	1774:  uint16(anon_sym_choice),
	1775:  uint16(anon_sym_endchoice),
	1776:  uint16(anon_sym_comment),
	1777:  uint16(anon_sym_menu),
	1778:  uint16(anon_sym_endmenu),
	1779:  uint16(anon_sym_if),
	1780:  uint16(anon_sym_endif),
	1781:  uint16(anon_sym_source),
	1782:  uint16(anon_sym_rsource),
	1783:  uint16(anon_sym_osource),
	1784:  uint16(anon_sym_orsource),
	1785:  uint16(anon_sym_bool),
	1786:  uint16(anon_sym_tristate),
	1787:  uint16(anon_sym_int),
	1788:  uint16(anon_sym_hex),
	1789:  uint16(anon_sym_string),
	1790:  uint16(anon_sym_prompt),
	1791:  uint16(anon_sym_default),
	1792:  uint16(anon_sym_def_bool),
	1793:  uint16(anon_sym_def_tristate),
	1794:  uint16(anon_sym_def_int),
	1795:  uint16(anon_sym_def_hex),
	1796:  uint16(anon_sym_def_string),
	1797:  uint16(anon_sym_select),
	1798:  uint16(anon_sym_imply),
	1799:  uint16(anon_sym_range),
	1800:  uint16(anon_sym_help),
	1801:  uint16(sym_optional),
	1802:  uint16(sym_modules),
	1803:  uint16(anon_sym_LT),
	1804:  uint16(anon_sym_GT),
	1805:  uint16(sym_symbol),
	1806:  uint16(3),
	1807:  uint16(3),
	1808:  uint16(1),
	1809:  uint16(sym_comment),
	1810:  uint16(293),
	1811:  uint16(8),
	1812:  uint16(anon_sym_EQ),
	1813:  uint16(anon_sym_dependson),
	1814:  uint16(anon_sym_visibleif),
	1815:  uint16(anon_sym_PIPE_PIPE),
	1816:  uint16(anon_sym_AMP_AMP),
	1817:  uint16(anon_sym_BANG_EQ),
	1818:  uint16(anon_sym_LT_EQ),
	1819:  uint16(anon_sym_GT_EQ),
	1820:  uint16(291),
	1821:  uint16(36),
	1822:  uint16(anon_sym_mainmenu),
	1823:  uint16(anon_sym_config),
	1824:  uint16(anon_sym_configdefault),
	1825:  uint16(anon_sym_menuconfig),
	1826:  uint16(anon_sym_choice),
	1827:  uint16(anon_sym_endchoice),
	1828:  uint16(anon_sym_comment),
	1829:  uint16(anon_sym_menu),
	1830:  uint16(anon_sym_endmenu),
	1831:  uint16(anon_sym_if),
	1832:  uint16(anon_sym_endif),
	1833:  uint16(anon_sym_source),
	1834:  uint16(anon_sym_rsource),
	1835:  uint16(anon_sym_osource),
	1836:  uint16(anon_sym_orsource),
	1837:  uint16(anon_sym_bool),
	1838:  uint16(anon_sym_tristate),
	1839:  uint16(anon_sym_int),
	1840:  uint16(anon_sym_hex),
	1841:  uint16(anon_sym_string),
	1842:  uint16(anon_sym_prompt),
	1843:  uint16(anon_sym_default),
	1844:  uint16(anon_sym_def_bool),
	1845:  uint16(anon_sym_def_tristate),
	1846:  uint16(anon_sym_def_int),
	1847:  uint16(anon_sym_def_hex),
	1848:  uint16(anon_sym_def_string),
	1849:  uint16(anon_sym_select),
	1850:  uint16(anon_sym_imply),
	1851:  uint16(anon_sym_range),
	1852:  uint16(anon_sym_help),
	1853:  uint16(sym_optional),
	1854:  uint16(sym_modules),
	1855:  uint16(anon_sym_LT),
	1856:  uint16(anon_sym_GT),
	1857:  uint16(sym_symbol),
	1858:  uint16(4),
	1859:  uint16(3),
	1860:  uint16(1),
	1861:  uint16(sym_comment),
	1862:  uint16(281),
	1863:  uint16(1),
	1864:  uint16(anon_sym_EQ),
	1865:  uint16(279),
	1866:  uint16(7),
	1867:  uint16(anon_sym_dependson),
	1868:  uint16(anon_sym_visibleif),
	1869:  uint16(anon_sym_PIPE_PIPE),
	1870:  uint16(anon_sym_AMP_AMP),
	1871:  uint16(anon_sym_BANG_EQ),
	1872:  uint16(anon_sym_LT_EQ),
	1873:  uint16(anon_sym_GT_EQ),
	1874:  uint16(277),
	1875:  uint16(36),
	1876:  uint16(anon_sym_mainmenu),
	1877:  uint16(anon_sym_config),
	1878:  uint16(anon_sym_configdefault),
	1879:  uint16(anon_sym_menuconfig),
	1880:  uint16(anon_sym_choice),
	1881:  uint16(anon_sym_endchoice),
	1882:  uint16(anon_sym_comment),
	1883:  uint16(anon_sym_menu),
	1884:  uint16(anon_sym_endmenu),
	1885:  uint16(anon_sym_if),
	1886:  uint16(anon_sym_endif),
	1887:  uint16(anon_sym_source),
	1888:  uint16(anon_sym_rsource),
	1889:  uint16(anon_sym_osource),
	1890:  uint16(anon_sym_orsource),
	1891:  uint16(anon_sym_bool),
	1892:  uint16(anon_sym_tristate),
	1893:  uint16(anon_sym_int),
	1894:  uint16(anon_sym_hex),
	1895:  uint16(anon_sym_string),
	1896:  uint16(anon_sym_prompt),
	1897:  uint16(anon_sym_default),
	1898:  uint16(anon_sym_def_bool),
	1899:  uint16(anon_sym_def_tristate),
	1900:  uint16(anon_sym_def_int),
	1901:  uint16(anon_sym_def_hex),
	1902:  uint16(anon_sym_def_string),
	1903:  uint16(anon_sym_select),
	1904:  uint16(anon_sym_imply),
	1905:  uint16(anon_sym_range),
	1906:  uint16(anon_sym_help),
	1907:  uint16(sym_optional),
	1908:  uint16(sym_modules),
	1909:  uint16(anon_sym_LT),
	1910:  uint16(anon_sym_GT),
	1911:  uint16(sym_symbol),
	1912:  uint16(7),
	1913:  uint16(3),
	1914:  uint16(1),
	1915:  uint16(sym_comment),
	1916:  uint16(281),
	1917:  uint16(1),
	1918:  uint16(anon_sym_EQ),
	1919:  uint16(295),
	1920:  uint16(1),
	1921:  uint16(anon_sym_AMP_AMP),
	1922:  uint16(285),
	1923:  uint16(2),
	1924:  uint16(anon_sym_LT),
	1925:  uint16(anon_sym_GT),
	1926:  uint16(279),
	1927:  uint16(3),
	1928:  uint16(anon_sym_dependson),
	1929:  uint16(anon_sym_visibleif),
	1930:  uint16(anon_sym_PIPE_PIPE),
	1931:  uint16(283),
	1932:  uint16(3),
	1933:  uint16(anon_sym_BANG_EQ),
	1934:  uint16(anon_sym_LT_EQ),
	1935:  uint16(anon_sym_GT_EQ),
	1936:  uint16(277),
	1937:  uint16(34),
	1938:  uint16(anon_sym_mainmenu),
	1939:  uint16(anon_sym_config),
	1940:  uint16(anon_sym_configdefault),
	1941:  uint16(anon_sym_menuconfig),
	1942:  uint16(anon_sym_choice),
	1943:  uint16(anon_sym_endchoice),
	1944:  uint16(anon_sym_comment),
	1945:  uint16(anon_sym_menu),
	1946:  uint16(anon_sym_endmenu),
	1947:  uint16(anon_sym_if),
	1948:  uint16(anon_sym_endif),
	1949:  uint16(anon_sym_source),
	1950:  uint16(anon_sym_rsource),
	1951:  uint16(anon_sym_osource),
	1952:  uint16(anon_sym_orsource),
	1953:  uint16(anon_sym_bool),
	1954:  uint16(anon_sym_tristate),
	1955:  uint16(anon_sym_int),
	1956:  uint16(anon_sym_hex),
	1957:  uint16(anon_sym_string),
	1958:  uint16(anon_sym_prompt),
	1959:  uint16(anon_sym_default),
	1960:  uint16(anon_sym_def_bool),
	1961:  uint16(anon_sym_def_tristate),
	1962:  uint16(anon_sym_def_int),
	1963:  uint16(anon_sym_def_hex),
	1964:  uint16(anon_sym_def_string),
	1965:  uint16(anon_sym_select),
	1966:  uint16(anon_sym_imply),
	1967:  uint16(anon_sym_range),
	1968:  uint16(anon_sym_help),
	1969:  uint16(sym_optional),
	1970:  uint16(sym_modules),
	1971:  uint16(sym_symbol),
	1972:  uint16(8),
	1973:  uint16(3),
	1974:  uint16(1),
	1975:  uint16(sym_comment),
	1976:  uint16(281),
	1977:  uint16(1),
	1978:  uint16(anon_sym_EQ),
	1979:  uint16(295),
	1980:  uint16(1),
	1981:  uint16(anon_sym_AMP_AMP),
	1982:  uint16(301),
	1983:  uint16(1),
	1984:  uint16(anon_sym_PIPE_PIPE),
	1985:  uint16(285),
	1986:  uint16(2),
	1987:  uint16(anon_sym_LT),
	1988:  uint16(anon_sym_GT),
	1989:  uint16(299),
	1990:  uint16(2),
	1991:  uint16(anon_sym_dependson),
	1992:  uint16(anon_sym_visibleif),
	1993:  uint16(283),
	1994:  uint16(3),
	1995:  uint16(anon_sym_BANG_EQ),
	1996:  uint16(anon_sym_LT_EQ),
	1997:  uint16(anon_sym_GT_EQ),
	1998:  uint16(297),
	1999:  uint16(34),
	2000:  uint16(anon_sym_mainmenu),
	2001:  uint16(anon_sym_config),
	2002:  uint16(anon_sym_configdefault),
	2003:  uint16(anon_sym_menuconfig),
	2004:  uint16(anon_sym_choice),
	2005:  uint16(anon_sym_endchoice),
	2006:  uint16(anon_sym_comment),
	2007:  uint16(anon_sym_menu),
	2008:  uint16(anon_sym_endmenu),
	2009:  uint16(anon_sym_if),
	2010:  uint16(anon_sym_endif),
	2011:  uint16(anon_sym_source),
	2012:  uint16(anon_sym_rsource),
	2013:  uint16(anon_sym_osource),
	2014:  uint16(anon_sym_orsource),
	2015:  uint16(anon_sym_bool),
	2016:  uint16(anon_sym_tristate),
	2017:  uint16(anon_sym_int),
	2018:  uint16(anon_sym_hex),
	2019:  uint16(anon_sym_string),
	2020:  uint16(anon_sym_prompt),
	2021:  uint16(anon_sym_default),
	2022:  uint16(anon_sym_def_bool),
	2023:  uint16(anon_sym_def_tristate),
	2024:  uint16(anon_sym_def_int),
	2025:  uint16(anon_sym_def_hex),
	2026:  uint16(anon_sym_def_string),
	2027:  uint16(anon_sym_select),
	2028:  uint16(anon_sym_imply),
	2029:  uint16(anon_sym_range),
	2030:  uint16(anon_sym_help),
	2031:  uint16(sym_optional),
	2032:  uint16(sym_modules),
	2033:  uint16(sym_symbol),
	2034:  uint16(3),
	2035:  uint16(3),
	2036:  uint16(1),
	2037:  uint16(sym_comment),
	2038:  uint16(293),
	2039:  uint16(9),
	2041:  uint16(anon_sym_EQ),
	2042:  uint16(anon_sym_dependson),
	2043:  uint16(anon_sym_visibleif),
	2044:  uint16(anon_sym_PIPE_PIPE),
	2045:  uint16(anon_sym_AMP_AMP),
	2046:  uint16(anon_sym_BANG_EQ),
	2047:  uint16(anon_sym_LT_EQ),
	2048:  uint16(anon_sym_GT_EQ),
	2049:  uint16(291),
	2050:  uint16(33),
	2051:  uint16(anon_sym_mainmenu),
	2052:  uint16(anon_sym_config),
	2053:  uint16(anon_sym_configdefault),
	2054:  uint16(anon_sym_menuconfig),
	2055:  uint16(anon_sym_choice),
	2056:  uint16(anon_sym_comment),
	2057:  uint16(anon_sym_menu),
	2058:  uint16(anon_sym_if),
	2059:  uint16(anon_sym_source),
	2060:  uint16(anon_sym_rsource),
	2061:  uint16(anon_sym_osource),
	2062:  uint16(anon_sym_orsource),
	2063:  uint16(anon_sym_bool),
	2064:  uint16(anon_sym_tristate),
	2065:  uint16(anon_sym_int),
	2066:  uint16(anon_sym_hex),
	2067:  uint16(anon_sym_string),
	2068:  uint16(anon_sym_prompt),
	2069:  uint16(anon_sym_default),
	2070:  uint16(anon_sym_def_bool),
	2071:  uint16(anon_sym_def_tristate),
	2072:  uint16(anon_sym_def_int),
	2073:  uint16(anon_sym_def_hex),
	2074:  uint16(anon_sym_def_string),
	2075:  uint16(anon_sym_select),
	2076:  uint16(anon_sym_imply),
	2077:  uint16(anon_sym_range),
	2078:  uint16(anon_sym_help),
	2079:  uint16(sym_optional),
	2080:  uint16(sym_modules),
	2081:  uint16(anon_sym_LT),
	2082:  uint16(anon_sym_GT),
	2083:  uint16(sym_symbol),
	2084:  uint16(8),
	2085:  uint16(3),
	2086:  uint16(1),
	2087:  uint16(sym_comment),
	2088:  uint16(303),
	2089:  uint16(1),
	2090:  uint16(anon_sym_EQ),
	2091:  uint16(305),
	2092:  uint16(1),
	2093:  uint16(anon_sym_PIPE_PIPE),
	2094:  uint16(307),
	2095:  uint16(1),
	2096:  uint16(anon_sym_AMP_AMP),
	2097:  uint16(311),
	2098:  uint16(2),
	2099:  uint16(anon_sym_LT),
	2100:  uint16(anon_sym_GT),
	2101:  uint16(299),
	2102:  uint16(3),
	2104:  uint16(anon_sym_dependson),
	2105:  uint16(anon_sym_visibleif),
	2106:  uint16(309),
	2107:  uint16(3),
	2108:  uint16(anon_sym_BANG_EQ),
	2109:  uint16(anon_sym_LT_EQ),
	2110:  uint16(anon_sym_GT_EQ),
	2111:  uint16(297),
	2112:  uint16(31),
	2113:  uint16(anon_sym_mainmenu),
	2114:  uint16(anon_sym_config),
	2115:  uint16(anon_sym_configdefault),
	2116:  uint16(anon_sym_menuconfig),
	2117:  uint16(anon_sym_choice),
	2118:  uint16(anon_sym_comment),
	2119:  uint16(anon_sym_menu),
	2120:  uint16(anon_sym_if),
	2121:  uint16(anon_sym_source),
	2122:  uint16(anon_sym_rsource),
	2123:  uint16(anon_sym_osource),
	2124:  uint16(anon_sym_orsource),
	2125:  uint16(anon_sym_bool),
	2126:  uint16(anon_sym_tristate),
	2127:  uint16(anon_sym_int),
	2128:  uint16(anon_sym_hex),
	2129:  uint16(anon_sym_string),
	2130:  uint16(anon_sym_prompt),
	2131:  uint16(anon_sym_default),
	2132:  uint16(anon_sym_def_bool),
	2133:  uint16(anon_sym_def_tristate),
	2134:  uint16(anon_sym_def_int),
	2135:  uint16(anon_sym_def_hex),
	2136:  uint16(anon_sym_def_string),
	2137:  uint16(anon_sym_select),
	2138:  uint16(anon_sym_imply),
	2139:  uint16(anon_sym_range),
	2140:  uint16(anon_sym_help),
	2141:  uint16(sym_optional),
	2142:  uint16(sym_modules),
	2143:  uint16(sym_symbol),
	2144:  uint16(3),
	2145:  uint16(3),
	2146:  uint16(1),
	2147:  uint16(sym_comment),
	2148:  uint16(289),
	2149:  uint16(9),
	2151:  uint16(anon_sym_EQ),
	2152:  uint16(anon_sym_dependson),
	2153:  uint16(anon_sym_visibleif),
	2154:  uint16(anon_sym_PIPE_PIPE),
	2155:  uint16(anon_sym_AMP_AMP),
	2156:  uint16(anon_sym_BANG_EQ),
	2157:  uint16(anon_sym_LT_EQ),
	2158:  uint16(anon_sym_GT_EQ),
	2159:  uint16(287),
	2160:  uint16(33),
	2161:  uint16(anon_sym_mainmenu),
	2162:  uint16(anon_sym_config),
	2163:  uint16(anon_sym_configdefault),
	2164:  uint16(anon_sym_menuconfig),
	2165:  uint16(anon_sym_choice),
	2166:  uint16(anon_sym_comment),
	2167:  uint16(anon_sym_menu),
	2168:  uint16(anon_sym_if),
	2169:  uint16(anon_sym_source),
	2170:  uint16(anon_sym_rsource),
	2171:  uint16(anon_sym_osource),
	2172:  uint16(anon_sym_orsource),
	2173:  uint16(anon_sym_bool),
	2174:  uint16(anon_sym_tristate),
	2175:  uint16(anon_sym_int),
	2176:  uint16(anon_sym_hex),
	2177:  uint16(anon_sym_string),
	2178:  uint16(anon_sym_prompt),
	2179:  uint16(anon_sym_default),
	2180:  uint16(anon_sym_def_bool),
	2181:  uint16(anon_sym_def_tristate),
	2182:  uint16(anon_sym_def_int),
	2183:  uint16(anon_sym_def_hex),
	2184:  uint16(anon_sym_def_string),
	2185:  uint16(anon_sym_select),
	2186:  uint16(anon_sym_imply),
	2187:  uint16(anon_sym_range),
	2188:  uint16(anon_sym_help),
	2189:  uint16(sym_optional),
	2190:  uint16(sym_modules),
	2191:  uint16(anon_sym_LT),
	2192:  uint16(anon_sym_GT),
	2193:  uint16(sym_symbol),
	2194:  uint16(3),
	2195:  uint16(3),
	2196:  uint16(1),
	2197:  uint16(sym_comment),
	2198:  uint16(275),
	2199:  uint16(9),
	2201:  uint16(anon_sym_EQ),
	2202:  uint16(anon_sym_dependson),
	2203:  uint16(anon_sym_visibleif),
	2204:  uint16(anon_sym_PIPE_PIPE),
	2205:  uint16(anon_sym_AMP_AMP),
	2206:  uint16(anon_sym_BANG_EQ),
	2207:  uint16(anon_sym_LT_EQ),
	2208:  uint16(anon_sym_GT_EQ),
	2209:  uint16(273),
	2210:  uint16(33),
	2211:  uint16(anon_sym_mainmenu),
	2212:  uint16(anon_sym_config),
	2213:  uint16(anon_sym_configdefault),
	2214:  uint16(anon_sym_menuconfig),
	2215:  uint16(anon_sym_choice),
	2216:  uint16(anon_sym_comment),
	2217:  uint16(anon_sym_menu),
	2218:  uint16(anon_sym_if),
	2219:  uint16(anon_sym_source),
	2220:  uint16(anon_sym_rsource),
	2221:  uint16(anon_sym_osource),
	2222:  uint16(anon_sym_orsource),
	2223:  uint16(anon_sym_bool),
	2224:  uint16(anon_sym_tristate),
	2225:  uint16(anon_sym_int),
	2226:  uint16(anon_sym_hex),
	2227:  uint16(anon_sym_string),
	2228:  uint16(anon_sym_prompt),
	2229:  uint16(anon_sym_default),
	2230:  uint16(anon_sym_def_bool),
	2231:  uint16(anon_sym_def_tristate),
	2232:  uint16(anon_sym_def_int),
	2233:  uint16(anon_sym_def_hex),
	2234:  uint16(anon_sym_def_string),
	2235:  uint16(anon_sym_select),
	2236:  uint16(anon_sym_imply),
	2237:  uint16(anon_sym_range),
	2238:  uint16(anon_sym_help),
	2239:  uint16(sym_optional),
	2240:  uint16(sym_modules),
	2241:  uint16(anon_sym_LT),
	2242:  uint16(anon_sym_GT),
	2243:  uint16(sym_symbol),
	2244:  uint16(7),
	2245:  uint16(3),
	2246:  uint16(1),
	2247:  uint16(sym_comment),
	2248:  uint16(303),
	2249:  uint16(1),
	2250:  uint16(anon_sym_EQ),
	2251:  uint16(307),
	2252:  uint16(1),
	2253:  uint16(anon_sym_AMP_AMP),
	2254:  uint16(311),
	2255:  uint16(2),
	2256:  uint16(anon_sym_LT),
	2257:  uint16(anon_sym_GT),
	2258:  uint16(309),
	2259:  uint16(3),
	2260:  uint16(anon_sym_BANG_EQ),
	2261:  uint16(anon_sym_LT_EQ),
	2262:  uint16(anon_sym_GT_EQ),
	2263:  uint16(279),
	2264:  uint16(4),
	2266:  uint16(anon_sym_dependson),
	2267:  uint16(anon_sym_visibleif),
	2268:  uint16(anon_sym_PIPE_PIPE),
	2269:  uint16(277),
	2270:  uint16(31),
	2271:  uint16(anon_sym_mainmenu),
	2272:  uint16(anon_sym_config),
	2273:  uint16(anon_sym_configdefault),
	2274:  uint16(anon_sym_menuconfig),
	2275:  uint16(anon_sym_choice),
	2276:  uint16(anon_sym_comment),
	2277:  uint16(anon_sym_menu),
	2278:  uint16(anon_sym_if),
	2279:  uint16(anon_sym_source),
	2280:  uint16(anon_sym_rsource),
	2281:  uint16(anon_sym_osource),
	2282:  uint16(anon_sym_orsource),
	2283:  uint16(anon_sym_bool),
	2284:  uint16(anon_sym_tristate),
	2285:  uint16(anon_sym_int),
	2286:  uint16(anon_sym_hex),
	2287:  uint16(anon_sym_string),
	2288:  uint16(anon_sym_prompt),
	2289:  uint16(anon_sym_default),
	2290:  uint16(anon_sym_def_bool),
	2291:  uint16(anon_sym_def_tristate),
	2292:  uint16(anon_sym_def_int),
	2293:  uint16(anon_sym_def_hex),
	2294:  uint16(anon_sym_def_string),
	2295:  uint16(anon_sym_select),
	2296:  uint16(anon_sym_imply),
	2297:  uint16(anon_sym_range),
	2298:  uint16(anon_sym_help),
	2299:  uint16(sym_optional),
	2300:  uint16(sym_modules),
	2301:  uint16(sym_symbol),
	2302:  uint16(3),
	2303:  uint16(3),
	2304:  uint16(1),
	2305:  uint16(sym_comment),
	2306:  uint16(279),
	2307:  uint16(9),
	2309:  uint16(anon_sym_EQ),
	2310:  uint16(anon_sym_dependson),
	2311:  uint16(anon_sym_visibleif),
	2312:  uint16(anon_sym_PIPE_PIPE),
	2313:  uint16(anon_sym_AMP_AMP),
	2314:  uint16(anon_sym_BANG_EQ),
	2315:  uint16(anon_sym_LT_EQ),
	2316:  uint16(anon_sym_GT_EQ),
	2317:  uint16(277),
	2318:  uint16(33),
	2319:  uint16(anon_sym_mainmenu),
	2320:  uint16(anon_sym_config),
	2321:  uint16(anon_sym_configdefault),
	2322:  uint16(anon_sym_menuconfig),
	2323:  uint16(anon_sym_choice),
	2324:  uint16(anon_sym_comment),
	2325:  uint16(anon_sym_menu),
	2326:  uint16(anon_sym_if),
	2327:  uint16(anon_sym_source),
	2328:  uint16(anon_sym_rsource),
	2329:  uint16(anon_sym_osource),
	2330:  uint16(anon_sym_orsource),
	2331:  uint16(anon_sym_bool),
	2332:  uint16(anon_sym_tristate),
	2333:  uint16(anon_sym_int),
	2334:  uint16(anon_sym_hex),
	2335:  uint16(anon_sym_string),
	2336:  uint16(anon_sym_prompt),
	2337:  uint16(anon_sym_default),
	2338:  uint16(anon_sym_def_bool),
	2339:  uint16(anon_sym_def_tristate),
	2340:  uint16(anon_sym_def_int),
	2341:  uint16(anon_sym_def_hex),
	2342:  uint16(anon_sym_def_string),
	2343:  uint16(anon_sym_select),
	2344:  uint16(anon_sym_imply),
	2345:  uint16(anon_sym_range),
	2346:  uint16(anon_sym_help),
	2347:  uint16(sym_optional),
	2348:  uint16(sym_modules),
	2349:  uint16(anon_sym_LT),
	2350:  uint16(anon_sym_GT),
	2351:  uint16(sym_symbol),
	2352:  uint16(6),
	2353:  uint16(3),
	2354:  uint16(1),
	2355:  uint16(sym_comment),
	2356:  uint16(303),
	2357:  uint16(1),
	2358:  uint16(anon_sym_EQ),
	2359:  uint16(311),
	2360:  uint16(2),
	2361:  uint16(anon_sym_LT),
	2362:  uint16(anon_sym_GT),
	2363:  uint16(309),
	2364:  uint16(3),
	2365:  uint16(anon_sym_BANG_EQ),
	2366:  uint16(anon_sym_LT_EQ),
	2367:  uint16(anon_sym_GT_EQ),
	2368:  uint16(279),
	2369:  uint16(5),
	2371:  uint16(anon_sym_dependson),
	2372:  uint16(anon_sym_visibleif),
	2373:  uint16(anon_sym_PIPE_PIPE),
	2374:  uint16(anon_sym_AMP_AMP),
	2375:  uint16(277),
	2376:  uint16(31),
	2377:  uint16(anon_sym_mainmenu),
	2378:  uint16(anon_sym_config),
	2379:  uint16(anon_sym_configdefault),
	2380:  uint16(anon_sym_menuconfig),
	2381:  uint16(anon_sym_choice),
	2382:  uint16(anon_sym_comment),
	2383:  uint16(anon_sym_menu),
	2384:  uint16(anon_sym_if),
	2385:  uint16(anon_sym_source),
	2386:  uint16(anon_sym_rsource),
	2387:  uint16(anon_sym_osource),
	2388:  uint16(anon_sym_orsource),
	2389:  uint16(anon_sym_bool),
	2390:  uint16(anon_sym_tristate),
	2391:  uint16(anon_sym_int),
	2392:  uint16(anon_sym_hex),
	2393:  uint16(anon_sym_string),
	2394:  uint16(anon_sym_prompt),
	2395:  uint16(anon_sym_default),
	2396:  uint16(anon_sym_def_bool),
	2397:  uint16(anon_sym_def_tristate),
	2398:  uint16(anon_sym_def_int),
	2399:  uint16(anon_sym_def_hex),
	2400:  uint16(anon_sym_def_string),
	2401:  uint16(anon_sym_select),
	2402:  uint16(anon_sym_imply),
	2403:  uint16(anon_sym_range),
	2404:  uint16(anon_sym_help),
	2405:  uint16(sym_optional),
	2406:  uint16(sym_modules),
	2407:  uint16(sym_symbol),
	2408:  uint16(4),
	2409:  uint16(3),
	2410:  uint16(1),
	2411:  uint16(sym_comment),
	2412:  uint16(303),
	2413:  uint16(1),
	2414:  uint16(anon_sym_EQ),
	2415:  uint16(279),
	2416:  uint16(8),
	2418:  uint16(anon_sym_dependson),
	2419:  uint16(anon_sym_visibleif),
	2420:  uint16(anon_sym_PIPE_PIPE),
	2421:  uint16(anon_sym_AMP_AMP),
	2422:  uint16(anon_sym_BANG_EQ),
	2423:  uint16(anon_sym_LT_EQ),
	2424:  uint16(anon_sym_GT_EQ),
	2425:  uint16(277),
	2426:  uint16(33),
	2427:  uint16(anon_sym_mainmenu),
	2428:  uint16(anon_sym_config),
	2429:  uint16(anon_sym_configdefault),
	2430:  uint16(anon_sym_menuconfig),
	2431:  uint16(anon_sym_choice),
	2432:  uint16(anon_sym_comment),
	2433:  uint16(anon_sym_menu),
	2434:  uint16(anon_sym_if),
	2435:  uint16(anon_sym_source),
	2436:  uint16(anon_sym_rsource),
	2437:  uint16(anon_sym_osource),
	2438:  uint16(anon_sym_orsource),
	2439:  uint16(anon_sym_bool),
	2440:  uint16(anon_sym_tristate),
	2441:  uint16(anon_sym_int),
	2442:  uint16(anon_sym_hex),
	2443:  uint16(anon_sym_string),
	2444:  uint16(anon_sym_prompt),
	2445:  uint16(anon_sym_default),
	2446:  uint16(anon_sym_def_bool),
	2447:  uint16(anon_sym_def_tristate),
	2448:  uint16(anon_sym_def_int),
	2449:  uint16(anon_sym_def_hex),
	2450:  uint16(anon_sym_def_string),
	2451:  uint16(anon_sym_select),
	2452:  uint16(anon_sym_imply),
	2453:  uint16(anon_sym_range),
	2454:  uint16(anon_sym_help),
	2455:  uint16(sym_optional),
	2456:  uint16(sym_modules),
	2457:  uint16(anon_sym_LT),
	2458:  uint16(anon_sym_GT),
	2459:  uint16(sym_symbol),
	2460:  uint16(19),
	2461:  uint16(3),
	2462:  uint16(1),
	2463:  uint16(sym_comment),
	2464:  uint16(51),
	2465:  uint16(1),
	2466:  uint16(anon_sym_prompt),
	2467:  uint16(53),
	2468:  uint16(1),
	2469:  uint16(anon_sym_default),
	2470:  uint16(57),
	2471:  uint16(1),
	2472:  uint16(anon_sym_dependson),
	2473:  uint16(59),
	2474:  uint16(1),
	2475:  uint16(anon_sym_select),
	2476:  uint16(61),
	2477:  uint16(1),
	2478:  uint16(anon_sym_imply),
	2479:  uint16(63),
	2480:  uint16(1),
	2481:  uint16(anon_sym_visibleif),
	2482:  uint16(65),
	2483:  uint16(1),
	2484:  uint16(anon_sym_range),
	2485:  uint16(67),
	2486:  uint16(1),
	2487:  uint16(anon_sym_help),
	2488:  uint16(95),
	2489:  uint16(1),
	2490:  uint16(anon_sym_DOLLAR_LPAREN),
	2491:  uint16(97),
	2492:  uint16(1),
	2493:  uint16(anon_sym_DQUOTE),
	2494:  uint16(99),
	2495:  uint16(1),
	2496:  uint16(anon_sym_SQUOTE),
	2497:  uint16(313),
	2498:  uint16(1),
	2499:  uint16(sym_symbol),
	2500:  uint16(89),
	2501:  uint16(1),
	2502:  uint16(sym_name),
	2503:  uint16(315),
	2504:  uint16(2),
	2505:  uint16(sym_optional),
	2506:  uint16(sym_modules),
	2507:  uint16(10),
	2508:  uint16(3),
	2509:  uint16(sym_macro_variable),
	2510:  uint16(sym_string),
	2511:  uint16(aux_sym_name_repeat1),
	2512:  uint16(49),
	2513:  uint16(5),
	2514:  uint16(anon_sym_bool),
	2515:  uint16(anon_sym_tristate),
	2516:  uint16(anon_sym_int),
	2517:  uint16(anon_sym_hex),
	2518:  uint16(anon_sym_string),
	2519:  uint16(55),
	2520:  uint16(5),
	2521:  uint16(anon_sym_def_bool),
	2522:  uint16(anon_sym_def_tristate),
	2523:  uint16(anon_sym_def_int),
	2524:  uint16(anon_sym_def_hex),
	2525:  uint16(anon_sym_def_string),
	2526:  uint16(3),
	2527:  uint16(12),
	2528:  uint16(sym__config_option),
	2529:  uint16(sym_type_definition),
	2530:  uint16(sym_input_prompt),
	2531:  uint16(sym_default_value),
	2532:  uint16(sym_type_definition_default),
	2533:  uint16(sym_dependencies),
	2534:  uint16(sym_reverse_dependencies),
	2535:  uint16(sym_weak_reverse_dependencies),
	2536:  uint16(sym_limiting_menu_display),
	2537:  uint16(sym_numerical_ranges),
	2538:  uint16(sym_help_text),
	2539:  uint16(aux_sym_config_repeat1),
	2540:  uint16(19),
	2541:  uint16(3),
	2542:  uint16(1),
	2543:  uint16(sym_comment),
	2544:  uint16(51),
	2545:  uint16(1),
	2546:  uint16(anon_sym_prompt),
	2547:  uint16(53),
	2548:  uint16(1),
	2549:  uint16(anon_sym_default),
	2550:  uint16(57),
	2551:  uint16(1),
	2552:  uint16(anon_sym_dependson),
	2553:  uint16(59),
	2554:  uint16(1),
	2555:  uint16(anon_sym_select),
	2556:  uint16(61),
	2557:  uint16(1),
	2558:  uint16(anon_sym_imply),
	2559:  uint16(63),
	2560:  uint16(1),
	2561:  uint16(anon_sym_visibleif),
	2562:  uint16(65),
	2563:  uint16(1),
	2564:  uint16(anon_sym_range),
	2565:  uint16(67),
	2566:  uint16(1),
	2567:  uint16(anon_sym_help),
	2568:  uint16(95),
	2569:  uint16(1),
	2570:  uint16(anon_sym_DOLLAR_LPAREN),
	2571:  uint16(97),
	2572:  uint16(1),
	2573:  uint16(anon_sym_DQUOTE),
	2574:  uint16(99),
	2575:  uint16(1),
	2576:  uint16(anon_sym_SQUOTE),
	2577:  uint16(313),
	2578:  uint16(1),
	2579:  uint16(sym_symbol),
	2580:  uint16(88),
	2581:  uint16(1),
	2582:  uint16(sym_name),
	2583:  uint16(317),
	2584:  uint16(2),
	2585:  uint16(sym_optional),
	2586:  uint16(sym_modules),
	2587:  uint16(10),
	2588:  uint16(3),
	2589:  uint16(sym_macro_variable),
	2590:  uint16(sym_string),
	2591:  uint16(aux_sym_name_repeat1),
	2592:  uint16(49),
	2593:  uint16(5),
	2594:  uint16(anon_sym_bool),
	2595:  uint16(anon_sym_tristate),
	2596:  uint16(anon_sym_int),
	2597:  uint16(anon_sym_hex),
	2598:  uint16(anon_sym_string),
	2599:  uint16(55),
	2600:  uint16(5),
	2601:  uint16(anon_sym_def_bool),
	2602:  uint16(anon_sym_def_tristate),
	2603:  uint16(anon_sym_def_int),
	2604:  uint16(anon_sym_def_hex),
	2605:  uint16(anon_sym_def_string),
	2606:  uint16(7),
	2607:  uint16(12),
	2608:  uint16(sym__config_option),
	2609:  uint16(sym_type_definition),
	2610:  uint16(sym_input_prompt),
	2611:  uint16(sym_default_value),
	2612:  uint16(sym_type_definition_default),
	2613:  uint16(sym_dependencies),
	2614:  uint16(sym_reverse_dependencies),
	2615:  uint16(sym_weak_reverse_dependencies),
	2616:  uint16(sym_limiting_menu_display),
	2617:  uint16(sym_numerical_ranges),
	2618:  uint16(sym_help_text),
	2619:  uint16(aux_sym_config_repeat1),
	2620:  uint16(3),
	2621:  uint16(3),
	2622:  uint16(1),
	2623:  uint16(sym_comment),
	2624:  uint16(321),
	2625:  uint16(2),
	2626:  uint16(anon_sym_dependson),
	2627:  uint16(anon_sym_visibleif),
	2628:  uint16(319),
	2629:  uint16(34),
	2630:  uint16(anon_sym_mainmenu),
	2631:  uint16(anon_sym_config),
	2632:  uint16(anon_sym_configdefault),
	2633:  uint16(anon_sym_menuconfig),
	2634:  uint16(anon_sym_choice),
	2635:  uint16(anon_sym_endchoice),
	2636:  uint16(anon_sym_comment),
	2637:  uint16(anon_sym_menu),
	2638:  uint16(anon_sym_endmenu),
	2639:  uint16(anon_sym_if),
	2640:  uint16(anon_sym_endif),
	2641:  uint16(anon_sym_source),
	2642:  uint16(anon_sym_rsource),
	2643:  uint16(anon_sym_osource),
	2644:  uint16(anon_sym_orsource),
	2645:  uint16(anon_sym_bool),
	2646:  uint16(anon_sym_tristate),
	2647:  uint16(anon_sym_int),
	2648:  uint16(anon_sym_hex),
	2649:  uint16(anon_sym_string),
	2650:  uint16(anon_sym_prompt),
	2651:  uint16(anon_sym_default),
	2652:  uint16(anon_sym_def_bool),
	2653:  uint16(anon_sym_def_tristate),
	2654:  uint16(anon_sym_def_int),
	2655:  uint16(anon_sym_def_hex),
	2656:  uint16(anon_sym_def_string),
	2657:  uint16(anon_sym_select),
	2658:  uint16(anon_sym_imply),
	2659:  uint16(anon_sym_range),
	2660:  uint16(anon_sym_help),
	2661:  uint16(sym_optional),
	2662:  uint16(sym_modules),
	2663:  uint16(sym_symbol),
	2664:  uint16(3),
	2665:  uint16(3),
	2666:  uint16(1),
	2667:  uint16(sym_comment),
	2668:  uint16(325),
	2669:  uint16(2),
	2670:  uint16(anon_sym_dependson),
	2671:  uint16(anon_sym_visibleif),
	2672:  uint16(323),
	2673:  uint16(34),
	2674:  uint16(anon_sym_mainmenu),
	2675:  uint16(anon_sym_config),
	2676:  uint16(anon_sym_configdefault),
	2677:  uint16(anon_sym_menuconfig),
	2678:  uint16(anon_sym_choice),
	2679:  uint16(anon_sym_endchoice),
	2680:  uint16(anon_sym_comment),
	2681:  uint16(anon_sym_menu),
	2682:  uint16(anon_sym_endmenu),
	2683:  uint16(anon_sym_if),
	2684:  uint16(anon_sym_endif),
	2685:  uint16(anon_sym_source),
	2686:  uint16(anon_sym_rsource),
	2687:  uint16(anon_sym_osource),
	2688:  uint16(anon_sym_orsource),
	2689:  uint16(anon_sym_bool),
	2690:  uint16(anon_sym_tristate),
	2691:  uint16(anon_sym_int),
	2692:  uint16(anon_sym_hex),
	2693:  uint16(anon_sym_string),
	2694:  uint16(anon_sym_prompt),
	2695:  uint16(anon_sym_default),
	2696:  uint16(anon_sym_def_bool),
	2697:  uint16(anon_sym_def_tristate),
	2698:  uint16(anon_sym_def_int),
	2699:  uint16(anon_sym_def_hex),
	2700:  uint16(anon_sym_def_string),
	2701:  uint16(anon_sym_select),
	2702:  uint16(anon_sym_imply),
	2703:  uint16(anon_sym_range),
	2704:  uint16(anon_sym_help),
	2705:  uint16(sym_optional),
	2706:  uint16(sym_modules),
	2707:  uint16(sym_symbol),
	2708:  uint16(3),
	2709:  uint16(3),
	2710:  uint16(1),
	2711:  uint16(sym_comment),
	2712:  uint16(329),
	2713:  uint16(2),
	2714:  uint16(anon_sym_dependson),
	2715:  uint16(anon_sym_visibleif),
	2716:  uint16(327),
	2717:  uint16(34),
	2718:  uint16(anon_sym_mainmenu),
	2719:  uint16(anon_sym_config),
	2720:  uint16(anon_sym_configdefault),
	2721:  uint16(anon_sym_menuconfig),
	2722:  uint16(anon_sym_choice),
	2723:  uint16(anon_sym_endchoice),
	2724:  uint16(anon_sym_comment),
	2725:  uint16(anon_sym_menu),
	2726:  uint16(anon_sym_endmenu),
	2727:  uint16(anon_sym_if),
	2728:  uint16(anon_sym_endif),
	2729:  uint16(anon_sym_source),
	2730:  uint16(anon_sym_rsource),
	2731:  uint16(anon_sym_osource),
	2732:  uint16(anon_sym_orsource),
	2733:  uint16(anon_sym_bool),
	2734:  uint16(anon_sym_tristate),
	2735:  uint16(anon_sym_int),
	2736:  uint16(anon_sym_hex),
	2737:  uint16(anon_sym_string),
	2738:  uint16(anon_sym_prompt),
	2739:  uint16(anon_sym_default),
	2740:  uint16(anon_sym_def_bool),
	2741:  uint16(anon_sym_def_tristate),
	2742:  uint16(anon_sym_def_int),
	2743:  uint16(anon_sym_def_hex),
	2744:  uint16(anon_sym_def_string),
	2745:  uint16(anon_sym_select),
	2746:  uint16(anon_sym_imply),
	2747:  uint16(anon_sym_range),
	2748:  uint16(anon_sym_help),
	2749:  uint16(sym_optional),
	2750:  uint16(sym_modules),
	2751:  uint16(sym_symbol),
	2752:  uint16(3),
	2753:  uint16(3),
	2754:  uint16(1),
	2755:  uint16(sym_comment),
	2756:  uint16(333),
	2757:  uint16(2),
	2758:  uint16(anon_sym_dependson),
	2759:  uint16(anon_sym_visibleif),
	2760:  uint16(331),
	2761:  uint16(34),
	2762:  uint16(anon_sym_mainmenu),
	2763:  uint16(anon_sym_config),
	2764:  uint16(anon_sym_configdefault),
	2765:  uint16(anon_sym_menuconfig),
	2766:  uint16(anon_sym_choice),
	2767:  uint16(anon_sym_endchoice),
	2768:  uint16(anon_sym_comment),
	2769:  uint16(anon_sym_menu),
	2770:  uint16(anon_sym_endmenu),
	2771:  uint16(anon_sym_if),
	2772:  uint16(anon_sym_endif),
	2773:  uint16(anon_sym_source),
	2774:  uint16(anon_sym_rsource),
	2775:  uint16(anon_sym_osource),
	2776:  uint16(anon_sym_orsource),
	2777:  uint16(anon_sym_bool),
	2778:  uint16(anon_sym_tristate),
	2779:  uint16(anon_sym_int),
	2780:  uint16(anon_sym_hex),
	2781:  uint16(anon_sym_string),
	2782:  uint16(anon_sym_prompt),
	2783:  uint16(anon_sym_default),
	2784:  uint16(anon_sym_def_bool),
	2785:  uint16(anon_sym_def_tristate),
	2786:  uint16(anon_sym_def_int),
	2787:  uint16(anon_sym_def_hex),
	2788:  uint16(anon_sym_def_string),
	2789:  uint16(anon_sym_select),
	2790:  uint16(anon_sym_imply),
	2791:  uint16(anon_sym_range),
	2792:  uint16(anon_sym_help),
	2793:  uint16(sym_optional),
	2794:  uint16(sym_modules),
	2795:  uint16(sym_symbol),
	2796:  uint16(3),
	2797:  uint16(3),
	2798:  uint16(1),
	2799:  uint16(sym_comment),
	2800:  uint16(337),
	2801:  uint16(2),
	2802:  uint16(anon_sym_dependson),
	2803:  uint16(anon_sym_visibleif),
	2804:  uint16(335),
	2805:  uint16(34),
	2806:  uint16(anon_sym_mainmenu),
	2807:  uint16(anon_sym_config),
	2808:  uint16(anon_sym_configdefault),
	2809:  uint16(anon_sym_menuconfig),
	2810:  uint16(anon_sym_choice),
	2811:  uint16(anon_sym_endchoice),
	2812:  uint16(anon_sym_comment),
	2813:  uint16(anon_sym_menu),
	2814:  uint16(anon_sym_endmenu),
	2815:  uint16(anon_sym_if),
	2816:  uint16(anon_sym_endif),
	2817:  uint16(anon_sym_source),
	2818:  uint16(anon_sym_rsource),
	2819:  uint16(anon_sym_osource),
	2820:  uint16(anon_sym_orsource),
	2821:  uint16(anon_sym_bool),
	2822:  uint16(anon_sym_tristate),
	2823:  uint16(anon_sym_int),
	2824:  uint16(anon_sym_hex),
	2825:  uint16(anon_sym_string),
	2826:  uint16(anon_sym_prompt),
	2827:  uint16(anon_sym_default),
	2828:  uint16(anon_sym_def_bool),
	2829:  uint16(anon_sym_def_tristate),
	2830:  uint16(anon_sym_def_int),
	2831:  uint16(anon_sym_def_hex),
	2832:  uint16(anon_sym_def_string),
	2833:  uint16(anon_sym_select),
	2834:  uint16(anon_sym_imply),
	2835:  uint16(anon_sym_range),
	2836:  uint16(anon_sym_help),
	2837:  uint16(sym_optional),
	2838:  uint16(sym_modules),
	2839:  uint16(sym_symbol),
	2840:  uint16(3),
	2841:  uint16(3),
	2842:  uint16(1),
	2843:  uint16(sym_comment),
	2844:  uint16(341),
	2845:  uint16(2),
	2846:  uint16(anon_sym_dependson),
	2847:  uint16(anon_sym_visibleif),
	2848:  uint16(339),
	2849:  uint16(34),
	2850:  uint16(anon_sym_mainmenu),
	2851:  uint16(anon_sym_config),
	2852:  uint16(anon_sym_configdefault),
	2853:  uint16(anon_sym_menuconfig),
	2854:  uint16(anon_sym_choice),
	2855:  uint16(anon_sym_endchoice),
	2856:  uint16(anon_sym_comment),
	2857:  uint16(anon_sym_menu),
	2858:  uint16(anon_sym_endmenu),
	2859:  uint16(anon_sym_if),
	2860:  uint16(anon_sym_endif),
	2861:  uint16(anon_sym_source),
	2862:  uint16(anon_sym_rsource),
	2863:  uint16(anon_sym_osource),
	2864:  uint16(anon_sym_orsource),
	2865:  uint16(anon_sym_bool),
	2866:  uint16(anon_sym_tristate),
	2867:  uint16(anon_sym_int),
	2868:  uint16(anon_sym_hex),
	2869:  uint16(anon_sym_string),
	2870:  uint16(anon_sym_prompt),
	2871:  uint16(anon_sym_default),
	2872:  uint16(anon_sym_def_bool),
	2873:  uint16(anon_sym_def_tristate),
	2874:  uint16(anon_sym_def_int),
	2875:  uint16(anon_sym_def_hex),
	2876:  uint16(anon_sym_def_string),
	2877:  uint16(anon_sym_select),
	2878:  uint16(anon_sym_imply),
	2879:  uint16(anon_sym_range),
	2880:  uint16(anon_sym_help),
	2881:  uint16(sym_optional),
	2882:  uint16(sym_modules),
	2883:  uint16(sym_symbol),
	2884:  uint16(3),
	2885:  uint16(3),
	2886:  uint16(1),
	2887:  uint16(sym_comment),
	2888:  uint16(345),
	2889:  uint16(2),
	2890:  uint16(anon_sym_dependson),
	2891:  uint16(anon_sym_visibleif),
	2892:  uint16(343),
	2893:  uint16(34),
	2894:  uint16(anon_sym_mainmenu),
	2895:  uint16(anon_sym_config),
	2896:  uint16(anon_sym_configdefault),
	2897:  uint16(anon_sym_menuconfig),
	2898:  uint16(anon_sym_choice),
	2899:  uint16(anon_sym_endchoice),
	2900:  uint16(anon_sym_comment),
	2901:  uint16(anon_sym_menu),
	2902:  uint16(anon_sym_endmenu),
	2903:  uint16(anon_sym_if),
	2904:  uint16(anon_sym_endif),
	2905:  uint16(anon_sym_source),
	2906:  uint16(anon_sym_rsource),
	2907:  uint16(anon_sym_osource),
	2908:  uint16(anon_sym_orsource),
	2909:  uint16(anon_sym_bool),
	2910:  uint16(anon_sym_tristate),
	2911:  uint16(anon_sym_int),
	2912:  uint16(anon_sym_hex),
	2913:  uint16(anon_sym_string),
	2914:  uint16(anon_sym_prompt),
	2915:  uint16(anon_sym_default),
	2916:  uint16(anon_sym_def_bool),
	2917:  uint16(anon_sym_def_tristate),
	2918:  uint16(anon_sym_def_int),
	2919:  uint16(anon_sym_def_hex),
	2920:  uint16(anon_sym_def_string),
	2921:  uint16(anon_sym_select),
	2922:  uint16(anon_sym_imply),
	2923:  uint16(anon_sym_range),
	2924:  uint16(anon_sym_help),
	2925:  uint16(sym_optional),
	2926:  uint16(sym_modules),
	2927:  uint16(sym_symbol),
	2928:  uint16(3),
	2929:  uint16(3),
	2930:  uint16(1),
	2931:  uint16(sym_comment),
	2932:  uint16(349),
	2933:  uint16(2),
	2934:  uint16(anon_sym_dependson),
	2935:  uint16(anon_sym_visibleif),
	2936:  uint16(347),
	2937:  uint16(34),
	2938:  uint16(anon_sym_mainmenu),
	2939:  uint16(anon_sym_config),
	2940:  uint16(anon_sym_configdefault),
	2941:  uint16(anon_sym_menuconfig),
	2942:  uint16(anon_sym_choice),
	2943:  uint16(anon_sym_endchoice),
	2944:  uint16(anon_sym_comment),
	2945:  uint16(anon_sym_menu),
	2946:  uint16(anon_sym_endmenu),
	2947:  uint16(anon_sym_if),
	2948:  uint16(anon_sym_endif),
	2949:  uint16(anon_sym_source),
	2950:  uint16(anon_sym_rsource),
	2951:  uint16(anon_sym_osource),
	2952:  uint16(anon_sym_orsource),
	2953:  uint16(anon_sym_bool),
	2954:  uint16(anon_sym_tristate),
	2955:  uint16(anon_sym_int),
	2956:  uint16(anon_sym_hex),
	2957:  uint16(anon_sym_string),
	2958:  uint16(anon_sym_prompt),
	2959:  uint16(anon_sym_default),
	2960:  uint16(anon_sym_def_bool),
	2961:  uint16(anon_sym_def_tristate),
	2962:  uint16(anon_sym_def_int),
	2963:  uint16(anon_sym_def_hex),
	2964:  uint16(anon_sym_def_string),
	2965:  uint16(anon_sym_select),
	2966:  uint16(anon_sym_imply),
	2967:  uint16(anon_sym_range),
	2968:  uint16(anon_sym_help),
	2969:  uint16(sym_optional),
	2970:  uint16(sym_modules),
	2971:  uint16(sym_symbol),
	2972:  uint16(3),
	2973:  uint16(3),
	2974:  uint16(1),
	2975:  uint16(sym_comment),
	2976:  uint16(353),
	2977:  uint16(2),
	2978:  uint16(anon_sym_dependson),
	2979:  uint16(anon_sym_visibleif),
	2980:  uint16(351),
	2981:  uint16(34),
	2982:  uint16(anon_sym_mainmenu),
	2983:  uint16(anon_sym_config),
	2984:  uint16(anon_sym_configdefault),
	2985:  uint16(anon_sym_menuconfig),
	2986:  uint16(anon_sym_choice),
	2987:  uint16(anon_sym_endchoice),
	2988:  uint16(anon_sym_comment),
	2989:  uint16(anon_sym_menu),
	2990:  uint16(anon_sym_endmenu),
	2991:  uint16(anon_sym_if),
	2992:  uint16(anon_sym_endif),
	2993:  uint16(anon_sym_source),
	2994:  uint16(anon_sym_rsource),
	2995:  uint16(anon_sym_osource),
	2996:  uint16(anon_sym_orsource),
	2997:  uint16(anon_sym_bool),
	2998:  uint16(anon_sym_tristate),
	2999:  uint16(anon_sym_int),
	3000:  uint16(anon_sym_hex),
	3001:  uint16(anon_sym_string),
	3002:  uint16(anon_sym_prompt),
	3003:  uint16(anon_sym_default),
	3004:  uint16(anon_sym_def_bool),
	3005:  uint16(anon_sym_def_tristate),
	3006:  uint16(anon_sym_def_int),
	3007:  uint16(anon_sym_def_hex),
	3008:  uint16(anon_sym_def_string),
	3009:  uint16(anon_sym_select),
	3010:  uint16(anon_sym_imply),
	3011:  uint16(anon_sym_range),
	3012:  uint16(anon_sym_help),
	3013:  uint16(sym_optional),
	3014:  uint16(sym_modules),
	3015:  uint16(sym_symbol),
	3016:  uint16(3),
	3017:  uint16(3),
	3018:  uint16(1),
	3019:  uint16(sym_comment),
	3020:  uint16(357),
	3021:  uint16(2),
	3022:  uint16(anon_sym_dependson),
	3023:  uint16(anon_sym_visibleif),
	3024:  uint16(355),
	3025:  uint16(34),
	3026:  uint16(anon_sym_mainmenu),
	3027:  uint16(anon_sym_config),
	3028:  uint16(anon_sym_configdefault),
	3029:  uint16(anon_sym_menuconfig),
	3030:  uint16(anon_sym_choice),
	3031:  uint16(anon_sym_endchoice),
	3032:  uint16(anon_sym_comment),
	3033:  uint16(anon_sym_menu),
	3034:  uint16(anon_sym_endmenu),
	3035:  uint16(anon_sym_if),
	3036:  uint16(anon_sym_endif),
	3037:  uint16(anon_sym_source),
	3038:  uint16(anon_sym_rsource),
	3039:  uint16(anon_sym_osource),
	3040:  uint16(anon_sym_orsource),
	3041:  uint16(anon_sym_bool),
	3042:  uint16(anon_sym_tristate),
	3043:  uint16(anon_sym_int),
	3044:  uint16(anon_sym_hex),
	3045:  uint16(anon_sym_string),
	3046:  uint16(anon_sym_prompt),
	3047:  uint16(anon_sym_default),
	3048:  uint16(anon_sym_def_bool),
	3049:  uint16(anon_sym_def_tristate),
	3050:  uint16(anon_sym_def_int),
	3051:  uint16(anon_sym_def_hex),
	3052:  uint16(anon_sym_def_string),
	3053:  uint16(anon_sym_select),
	3054:  uint16(anon_sym_imply),
	3055:  uint16(anon_sym_range),
	3056:  uint16(anon_sym_help),
	3057:  uint16(sym_optional),
	3058:  uint16(sym_modules),
	3059:  uint16(sym_symbol),
	3060:  uint16(3),
	3061:  uint16(3),
	3062:  uint16(1),
	3063:  uint16(sym_comment),
	3064:  uint16(361),
	3065:  uint16(2),
	3066:  uint16(anon_sym_dependson),
	3067:  uint16(anon_sym_visibleif),
	3068:  uint16(359),
	3069:  uint16(34),
	3070:  uint16(anon_sym_mainmenu),
	3071:  uint16(anon_sym_config),
	3072:  uint16(anon_sym_configdefault),
	3073:  uint16(anon_sym_menuconfig),
	3074:  uint16(anon_sym_choice),
	3075:  uint16(anon_sym_endchoice),
	3076:  uint16(anon_sym_comment),
	3077:  uint16(anon_sym_menu),
	3078:  uint16(anon_sym_endmenu),
	3079:  uint16(anon_sym_if),
	3080:  uint16(anon_sym_endif),
	3081:  uint16(anon_sym_source),
	3082:  uint16(anon_sym_rsource),
	3083:  uint16(anon_sym_osource),
	3084:  uint16(anon_sym_orsource),
	3085:  uint16(anon_sym_bool),
	3086:  uint16(anon_sym_tristate),
	3087:  uint16(anon_sym_int),
	3088:  uint16(anon_sym_hex),
	3089:  uint16(anon_sym_string),
	3090:  uint16(anon_sym_prompt),
	3091:  uint16(anon_sym_default),
	3092:  uint16(anon_sym_def_bool),
	3093:  uint16(anon_sym_def_tristate),
	3094:  uint16(anon_sym_def_int),
	3095:  uint16(anon_sym_def_hex),
	3096:  uint16(anon_sym_def_string),
	3097:  uint16(anon_sym_select),
	3098:  uint16(anon_sym_imply),
	3099:  uint16(anon_sym_range),
	3100:  uint16(anon_sym_help),
	3101:  uint16(sym_optional),
	3102:  uint16(sym_modules),
	3103:  uint16(sym_symbol),
	3104:  uint16(3),
	3105:  uint16(3),
	3106:  uint16(1),
	3107:  uint16(sym_comment),
	3108:  uint16(365),
	3109:  uint16(2),
	3110:  uint16(anon_sym_dependson),
	3111:  uint16(anon_sym_visibleif),
	3112:  uint16(363),
	3113:  uint16(34),
	3114:  uint16(anon_sym_mainmenu),
	3115:  uint16(anon_sym_config),
	3116:  uint16(anon_sym_configdefault),
	3117:  uint16(anon_sym_menuconfig),
	3118:  uint16(anon_sym_choice),
	3119:  uint16(anon_sym_endchoice),
	3120:  uint16(anon_sym_comment),
	3121:  uint16(anon_sym_menu),
	3122:  uint16(anon_sym_endmenu),
	3123:  uint16(anon_sym_if),
	3124:  uint16(anon_sym_endif),
	3125:  uint16(anon_sym_source),
	3126:  uint16(anon_sym_rsource),
	3127:  uint16(anon_sym_osource),
	3128:  uint16(anon_sym_orsource),
	3129:  uint16(anon_sym_bool),
	3130:  uint16(anon_sym_tristate),
	3131:  uint16(anon_sym_int),
	3132:  uint16(anon_sym_hex),
	3133:  uint16(anon_sym_string),
	3134:  uint16(anon_sym_prompt),
	3135:  uint16(anon_sym_default),
	3136:  uint16(anon_sym_def_bool),
	3137:  uint16(anon_sym_def_tristate),
	3138:  uint16(anon_sym_def_int),
	3139:  uint16(anon_sym_def_hex),
	3140:  uint16(anon_sym_def_string),
	3141:  uint16(anon_sym_select),
	3142:  uint16(anon_sym_imply),
	3143:  uint16(anon_sym_range),
	3144:  uint16(anon_sym_help),
	3145:  uint16(sym_optional),
	3146:  uint16(sym_modules),
	3147:  uint16(sym_symbol),
	3148:  uint16(3),
	3149:  uint16(3),
	3150:  uint16(1),
	3151:  uint16(sym_comment),
	3152:  uint16(369),
	3153:  uint16(2),
	3154:  uint16(anon_sym_dependson),
	3155:  uint16(anon_sym_visibleif),
	3156:  uint16(367),
	3157:  uint16(34),
	3158:  uint16(anon_sym_mainmenu),
	3159:  uint16(anon_sym_config),
	3160:  uint16(anon_sym_configdefault),
	3161:  uint16(anon_sym_menuconfig),
	3162:  uint16(anon_sym_choice),
	3163:  uint16(anon_sym_endchoice),
	3164:  uint16(anon_sym_comment),
	3165:  uint16(anon_sym_menu),
	3166:  uint16(anon_sym_endmenu),
	3167:  uint16(anon_sym_if),
	3168:  uint16(anon_sym_endif),
	3169:  uint16(anon_sym_source),
	3170:  uint16(anon_sym_rsource),
	3171:  uint16(anon_sym_osource),
	3172:  uint16(anon_sym_orsource),
	3173:  uint16(anon_sym_bool),
	3174:  uint16(anon_sym_tristate),
	3175:  uint16(anon_sym_int),
	3176:  uint16(anon_sym_hex),
	3177:  uint16(anon_sym_string),
	3178:  uint16(anon_sym_prompt),
	3179:  uint16(anon_sym_default),
	3180:  uint16(anon_sym_def_bool),
	3181:  uint16(anon_sym_def_tristate),
	3182:  uint16(anon_sym_def_int),
	3183:  uint16(anon_sym_def_hex),
	3184:  uint16(anon_sym_def_string),
	3185:  uint16(anon_sym_select),
	3186:  uint16(anon_sym_imply),
	3187:  uint16(anon_sym_range),
	3188:  uint16(anon_sym_help),
	3189:  uint16(sym_optional),
	3190:  uint16(sym_modules),
	3191:  uint16(sym_symbol),
	3192:  uint16(3),
	3193:  uint16(3),
	3194:  uint16(1),
	3195:  uint16(sym_comment),
	3196:  uint16(373),
	3197:  uint16(2),
	3198:  uint16(anon_sym_dependson),
	3199:  uint16(anon_sym_visibleif),
	3200:  uint16(371),
	3201:  uint16(34),
	3202:  uint16(anon_sym_mainmenu),
	3203:  uint16(anon_sym_config),
	3204:  uint16(anon_sym_configdefault),
	3205:  uint16(anon_sym_menuconfig),
	3206:  uint16(anon_sym_choice),
	3207:  uint16(anon_sym_endchoice),
	3208:  uint16(anon_sym_comment),
	3209:  uint16(anon_sym_menu),
	3210:  uint16(anon_sym_endmenu),
	3211:  uint16(anon_sym_if),
	3212:  uint16(anon_sym_endif),
	3213:  uint16(anon_sym_source),
	3214:  uint16(anon_sym_rsource),
	3215:  uint16(anon_sym_osource),
	3216:  uint16(anon_sym_orsource),
	3217:  uint16(anon_sym_bool),
	3218:  uint16(anon_sym_tristate),
	3219:  uint16(anon_sym_int),
	3220:  uint16(anon_sym_hex),
	3221:  uint16(anon_sym_string),
	3222:  uint16(anon_sym_prompt),
	3223:  uint16(anon_sym_default),
	3224:  uint16(anon_sym_def_bool),
	3225:  uint16(anon_sym_def_tristate),
	3226:  uint16(anon_sym_def_int),
	3227:  uint16(anon_sym_def_hex),
	3228:  uint16(anon_sym_def_string),
	3229:  uint16(anon_sym_select),
	3230:  uint16(anon_sym_imply),
	3231:  uint16(anon_sym_range),
	3232:  uint16(anon_sym_help),
	3233:  uint16(sym_optional),
	3234:  uint16(sym_modules),
	3235:  uint16(sym_symbol),
	3236:  uint16(3),
	3237:  uint16(3),
	3238:  uint16(1),
	3239:  uint16(sym_comment),
	3240:  uint16(377),
	3241:  uint16(2),
	3242:  uint16(anon_sym_dependson),
	3243:  uint16(anon_sym_visibleif),
	3244:  uint16(375),
	3245:  uint16(34),
	3246:  uint16(anon_sym_mainmenu),
	3247:  uint16(anon_sym_config),
	3248:  uint16(anon_sym_configdefault),
	3249:  uint16(anon_sym_menuconfig),
	3250:  uint16(anon_sym_choice),
	3251:  uint16(anon_sym_endchoice),
	3252:  uint16(anon_sym_comment),
	3253:  uint16(anon_sym_menu),
	3254:  uint16(anon_sym_endmenu),
	3255:  uint16(anon_sym_if),
	3256:  uint16(anon_sym_endif),
	3257:  uint16(anon_sym_source),
	3258:  uint16(anon_sym_rsource),
	3259:  uint16(anon_sym_osource),
	3260:  uint16(anon_sym_orsource),
	3261:  uint16(anon_sym_bool),
	3262:  uint16(anon_sym_tristate),
	3263:  uint16(anon_sym_int),
	3264:  uint16(anon_sym_hex),
	3265:  uint16(anon_sym_string),
	3266:  uint16(anon_sym_prompt),
	3267:  uint16(anon_sym_default),
	3268:  uint16(anon_sym_def_bool),
	3269:  uint16(anon_sym_def_tristate),
	3270:  uint16(anon_sym_def_int),
	3271:  uint16(anon_sym_def_hex),
	3272:  uint16(anon_sym_def_string),
	3273:  uint16(anon_sym_select),
	3274:  uint16(anon_sym_imply),
	3275:  uint16(anon_sym_range),
	3276:  uint16(anon_sym_help),
	3277:  uint16(sym_optional),
	3278:  uint16(sym_modules),
	3279:  uint16(sym_symbol),
	3280:  uint16(3),
	3281:  uint16(3),
	3282:  uint16(1),
	3283:  uint16(sym_comment),
	3284:  uint16(381),
	3285:  uint16(2),
	3286:  uint16(anon_sym_dependson),
	3287:  uint16(anon_sym_visibleif),
	3288:  uint16(379),
	3289:  uint16(34),
	3290:  uint16(anon_sym_mainmenu),
	3291:  uint16(anon_sym_config),
	3292:  uint16(anon_sym_configdefault),
	3293:  uint16(anon_sym_menuconfig),
	3294:  uint16(anon_sym_choice),
	3295:  uint16(anon_sym_endchoice),
	3296:  uint16(anon_sym_comment),
	3297:  uint16(anon_sym_menu),
	3298:  uint16(anon_sym_endmenu),
	3299:  uint16(anon_sym_if),
	3300:  uint16(anon_sym_endif),
	3301:  uint16(anon_sym_source),
	3302:  uint16(anon_sym_rsource),
	3303:  uint16(anon_sym_osource),
	3304:  uint16(anon_sym_orsource),
	3305:  uint16(anon_sym_bool),
	3306:  uint16(anon_sym_tristate),
	3307:  uint16(anon_sym_int),
	3308:  uint16(anon_sym_hex),
	3309:  uint16(anon_sym_string),
	3310:  uint16(anon_sym_prompt),
	3311:  uint16(anon_sym_default),
	3312:  uint16(anon_sym_def_bool),
	3313:  uint16(anon_sym_def_tristate),
	3314:  uint16(anon_sym_def_int),
	3315:  uint16(anon_sym_def_hex),
	3316:  uint16(anon_sym_def_string),
	3317:  uint16(anon_sym_select),
	3318:  uint16(anon_sym_imply),
	3319:  uint16(anon_sym_range),
	3320:  uint16(anon_sym_help),
	3321:  uint16(sym_optional),
	3322:  uint16(sym_modules),
	3323:  uint16(sym_symbol),
	3324:  uint16(3),
	3325:  uint16(3),
	3326:  uint16(1),
	3327:  uint16(sym_comment),
	3328:  uint16(385),
	3329:  uint16(2),
	3330:  uint16(anon_sym_dependson),
	3331:  uint16(anon_sym_visibleif),
	3332:  uint16(383),
	3333:  uint16(34),
	3334:  uint16(anon_sym_mainmenu),
	3335:  uint16(anon_sym_config),
	3336:  uint16(anon_sym_configdefault),
	3337:  uint16(anon_sym_menuconfig),
	3338:  uint16(anon_sym_choice),
	3339:  uint16(anon_sym_endchoice),
	3340:  uint16(anon_sym_comment),
	3341:  uint16(anon_sym_menu),
	3342:  uint16(anon_sym_endmenu),
	3343:  uint16(anon_sym_if),
	3344:  uint16(anon_sym_endif),
	3345:  uint16(anon_sym_source),
	3346:  uint16(anon_sym_rsource),
	3347:  uint16(anon_sym_osource),
	3348:  uint16(anon_sym_orsource),
	3349:  uint16(anon_sym_bool),
	3350:  uint16(anon_sym_tristate),
	3351:  uint16(anon_sym_int),
	3352:  uint16(anon_sym_hex),
	3353:  uint16(anon_sym_string),
	3354:  uint16(anon_sym_prompt),
	3355:  uint16(anon_sym_default),
	3356:  uint16(anon_sym_def_bool),
	3357:  uint16(anon_sym_def_tristate),
	3358:  uint16(anon_sym_def_int),
	3359:  uint16(anon_sym_def_hex),
	3360:  uint16(anon_sym_def_string),
	3361:  uint16(anon_sym_select),
	3362:  uint16(anon_sym_imply),
	3363:  uint16(anon_sym_range),
	3364:  uint16(anon_sym_help),
	3365:  uint16(sym_optional),
	3366:  uint16(sym_modules),
	3367:  uint16(sym_symbol),
	3368:  uint16(3),
	3369:  uint16(3),
	3370:  uint16(1),
	3371:  uint16(sym_comment),
	3372:  uint16(381),
	3373:  uint16(3),
	3375:  uint16(anon_sym_dependson),
	3376:  uint16(anon_sym_visibleif),
	3377:  uint16(379),
	3378:  uint16(31),
	3379:  uint16(anon_sym_mainmenu),
	3380:  uint16(anon_sym_config),
	3381:  uint16(anon_sym_configdefault),
	3382:  uint16(anon_sym_menuconfig),
	3383:  uint16(anon_sym_choice),
	3384:  uint16(anon_sym_comment),
	3385:  uint16(anon_sym_menu),
	3386:  uint16(anon_sym_if),
	3387:  uint16(anon_sym_source),
	3388:  uint16(anon_sym_rsource),
	3389:  uint16(anon_sym_osource),
	3390:  uint16(anon_sym_orsource),
	3391:  uint16(anon_sym_bool),
	3392:  uint16(anon_sym_tristate),
	3393:  uint16(anon_sym_int),
	3394:  uint16(anon_sym_hex),
	3395:  uint16(anon_sym_string),
	3396:  uint16(anon_sym_prompt),
	3397:  uint16(anon_sym_default),
	3398:  uint16(anon_sym_def_bool),
	3399:  uint16(anon_sym_def_tristate),
	3400:  uint16(anon_sym_def_int),
	3401:  uint16(anon_sym_def_hex),
	3402:  uint16(anon_sym_def_string),
	3403:  uint16(anon_sym_select),
	3404:  uint16(anon_sym_imply),
	3405:  uint16(anon_sym_range),
	3406:  uint16(anon_sym_help),
	3407:  uint16(sym_optional),
	3408:  uint16(sym_modules),
	3409:  uint16(sym_symbol),
	3410:  uint16(3),
	3411:  uint16(3),
	3412:  uint16(1),
	3413:  uint16(sym_comment),
	3414:  uint16(357),
	3415:  uint16(3),
	3417:  uint16(anon_sym_dependson),
	3418:  uint16(anon_sym_visibleif),
	3419:  uint16(355),
	3420:  uint16(31),
	3421:  uint16(anon_sym_mainmenu),
	3422:  uint16(anon_sym_config),
	3423:  uint16(anon_sym_configdefault),
	3424:  uint16(anon_sym_menuconfig),
	3425:  uint16(anon_sym_choice),
	3426:  uint16(anon_sym_comment),
	3427:  uint16(anon_sym_menu),
	3428:  uint16(anon_sym_if),
	3429:  uint16(anon_sym_source),
	3430:  uint16(anon_sym_rsource),
	3431:  uint16(anon_sym_osource),
	3432:  uint16(anon_sym_orsource),
	3433:  uint16(anon_sym_bool),
	3434:  uint16(anon_sym_tristate),
	3435:  uint16(anon_sym_int),
	3436:  uint16(anon_sym_hex),
	3437:  uint16(anon_sym_string),
	3438:  uint16(anon_sym_prompt),
	3439:  uint16(anon_sym_default),
	3440:  uint16(anon_sym_def_bool),
	3441:  uint16(anon_sym_def_tristate),
	3442:  uint16(anon_sym_def_int),
	3443:  uint16(anon_sym_def_hex),
	3444:  uint16(anon_sym_def_string),
	3445:  uint16(anon_sym_select),
	3446:  uint16(anon_sym_imply),
	3447:  uint16(anon_sym_range),
	3448:  uint16(anon_sym_help),
	3449:  uint16(sym_optional),
	3450:  uint16(sym_modules),
	3451:  uint16(sym_symbol),
	3452:  uint16(3),
	3453:  uint16(3),
	3454:  uint16(1),
	3455:  uint16(sym_comment),
	3456:  uint16(341),
	3457:  uint16(3),
	3459:  uint16(anon_sym_dependson),
	3460:  uint16(anon_sym_visibleif),
	3461:  uint16(339),
	3462:  uint16(31),
	3463:  uint16(anon_sym_mainmenu),
	3464:  uint16(anon_sym_config),
	3465:  uint16(anon_sym_configdefault),
	3466:  uint16(anon_sym_menuconfig),
	3467:  uint16(anon_sym_choice),
	3468:  uint16(anon_sym_comment),
	3469:  uint16(anon_sym_menu),
	3470:  uint16(anon_sym_if),
	3471:  uint16(anon_sym_source),
	3472:  uint16(anon_sym_rsource),
	3473:  uint16(anon_sym_osource),
	3474:  uint16(anon_sym_orsource),
	3475:  uint16(anon_sym_bool),
	3476:  uint16(anon_sym_tristate),
	3477:  uint16(anon_sym_int),
	3478:  uint16(anon_sym_hex),
	3479:  uint16(anon_sym_string),
	3480:  uint16(anon_sym_prompt),
	3481:  uint16(anon_sym_default),
	3482:  uint16(anon_sym_def_bool),
	3483:  uint16(anon_sym_def_tristate),
	3484:  uint16(anon_sym_def_int),
	3485:  uint16(anon_sym_def_hex),
	3486:  uint16(anon_sym_def_string),
	3487:  uint16(anon_sym_select),
	3488:  uint16(anon_sym_imply),
	3489:  uint16(anon_sym_range),
	3490:  uint16(anon_sym_help),
	3491:  uint16(sym_optional),
	3492:  uint16(sym_modules),
	3493:  uint16(sym_symbol),
	3494:  uint16(3),
	3495:  uint16(3),
	3496:  uint16(1),
	3497:  uint16(sym_comment),
	3498:  uint16(373),
	3499:  uint16(3),
	3501:  uint16(anon_sym_dependson),
	3502:  uint16(anon_sym_visibleif),
	3503:  uint16(371),
	3504:  uint16(31),
	3505:  uint16(anon_sym_mainmenu),
	3506:  uint16(anon_sym_config),
	3507:  uint16(anon_sym_configdefault),
	3508:  uint16(anon_sym_menuconfig),
	3509:  uint16(anon_sym_choice),
	3510:  uint16(anon_sym_comment),
	3511:  uint16(anon_sym_menu),
	3512:  uint16(anon_sym_if),
	3513:  uint16(anon_sym_source),
	3514:  uint16(anon_sym_rsource),
	3515:  uint16(anon_sym_osource),
	3516:  uint16(anon_sym_orsource),
	3517:  uint16(anon_sym_bool),
	3518:  uint16(anon_sym_tristate),
	3519:  uint16(anon_sym_int),
	3520:  uint16(anon_sym_hex),
	3521:  uint16(anon_sym_string),
	3522:  uint16(anon_sym_prompt),
	3523:  uint16(anon_sym_default),
	3524:  uint16(anon_sym_def_bool),
	3525:  uint16(anon_sym_def_tristate),
	3526:  uint16(anon_sym_def_int),
	3527:  uint16(anon_sym_def_hex),
	3528:  uint16(anon_sym_def_string),
	3529:  uint16(anon_sym_select),
	3530:  uint16(anon_sym_imply),
	3531:  uint16(anon_sym_range),
	3532:  uint16(anon_sym_help),
	3533:  uint16(sym_optional),
	3534:  uint16(sym_modules),
	3535:  uint16(sym_symbol),
	3536:  uint16(3),
	3537:  uint16(3),
	3538:  uint16(1),
	3539:  uint16(sym_comment),
	3540:  uint16(377),
	3541:  uint16(3),
	3543:  uint16(anon_sym_dependson),
	3544:  uint16(anon_sym_visibleif),
	3545:  uint16(375),
	3546:  uint16(31),
	3547:  uint16(anon_sym_mainmenu),
	3548:  uint16(anon_sym_config),
	3549:  uint16(anon_sym_configdefault),
	3550:  uint16(anon_sym_menuconfig),
	3551:  uint16(anon_sym_choice),
	3552:  uint16(anon_sym_comment),
	3553:  uint16(anon_sym_menu),
	3554:  uint16(anon_sym_if),
	3555:  uint16(anon_sym_source),
	3556:  uint16(anon_sym_rsource),
	3557:  uint16(anon_sym_osource),
	3558:  uint16(anon_sym_orsource),
	3559:  uint16(anon_sym_bool),
	3560:  uint16(anon_sym_tristate),
	3561:  uint16(anon_sym_int),
	3562:  uint16(anon_sym_hex),
	3563:  uint16(anon_sym_string),
	3564:  uint16(anon_sym_prompt),
	3565:  uint16(anon_sym_default),
	3566:  uint16(anon_sym_def_bool),
	3567:  uint16(anon_sym_def_tristate),
	3568:  uint16(anon_sym_def_int),
	3569:  uint16(anon_sym_def_hex),
	3570:  uint16(anon_sym_def_string),
	3571:  uint16(anon_sym_select),
	3572:  uint16(anon_sym_imply),
	3573:  uint16(anon_sym_range),
	3574:  uint16(anon_sym_help),
	3575:  uint16(sym_optional),
	3576:  uint16(sym_modules),
	3577:  uint16(sym_symbol),
	3578:  uint16(3),
	3579:  uint16(3),
	3580:  uint16(1),
	3581:  uint16(sym_comment),
	3582:  uint16(325),
	3583:  uint16(3),
	3585:  uint16(anon_sym_dependson),
	3586:  uint16(anon_sym_visibleif),
	3587:  uint16(323),
	3588:  uint16(31),
	3589:  uint16(anon_sym_mainmenu),
	3590:  uint16(anon_sym_config),
	3591:  uint16(anon_sym_configdefault),
	3592:  uint16(anon_sym_menuconfig),
	3593:  uint16(anon_sym_choice),
	3594:  uint16(anon_sym_comment),
	3595:  uint16(anon_sym_menu),
	3596:  uint16(anon_sym_if),
	3597:  uint16(anon_sym_source),
	3598:  uint16(anon_sym_rsource),
	3599:  uint16(anon_sym_osource),
	3600:  uint16(anon_sym_orsource),
	3601:  uint16(anon_sym_bool),
	3602:  uint16(anon_sym_tristate),
	3603:  uint16(anon_sym_int),
	3604:  uint16(anon_sym_hex),
	3605:  uint16(anon_sym_string),
	3606:  uint16(anon_sym_prompt),
	3607:  uint16(anon_sym_default),
	3608:  uint16(anon_sym_def_bool),
	3609:  uint16(anon_sym_def_tristate),
	3610:  uint16(anon_sym_def_int),
	3611:  uint16(anon_sym_def_hex),
	3612:  uint16(anon_sym_def_string),
	3613:  uint16(anon_sym_select),
	3614:  uint16(anon_sym_imply),
	3615:  uint16(anon_sym_range),
	3616:  uint16(anon_sym_help),
	3617:  uint16(sym_optional),
	3618:  uint16(sym_modules),
	3619:  uint16(sym_symbol),
	3620:  uint16(3),
	3621:  uint16(3),
	3622:  uint16(1),
	3623:  uint16(sym_comment),
	3624:  uint16(369),
	3625:  uint16(3),
	3627:  uint16(anon_sym_dependson),
	3628:  uint16(anon_sym_visibleif),
	3629:  uint16(367),
	3630:  uint16(31),
	3631:  uint16(anon_sym_mainmenu),
	3632:  uint16(anon_sym_config),
	3633:  uint16(anon_sym_configdefault),
	3634:  uint16(anon_sym_menuconfig),
	3635:  uint16(anon_sym_choice),
	3636:  uint16(anon_sym_comment),
	3637:  uint16(anon_sym_menu),
	3638:  uint16(anon_sym_if),
	3639:  uint16(anon_sym_source),
	3640:  uint16(anon_sym_rsource),
	3641:  uint16(anon_sym_osource),
	3642:  uint16(anon_sym_orsource),
	3643:  uint16(anon_sym_bool),
	3644:  uint16(anon_sym_tristate),
	3645:  uint16(anon_sym_int),
	3646:  uint16(anon_sym_hex),
	3647:  uint16(anon_sym_string),
	3648:  uint16(anon_sym_prompt),
	3649:  uint16(anon_sym_default),
	3650:  uint16(anon_sym_def_bool),
	3651:  uint16(anon_sym_def_tristate),
	3652:  uint16(anon_sym_def_int),
	3653:  uint16(anon_sym_def_hex),
	3654:  uint16(anon_sym_def_string),
	3655:  uint16(anon_sym_select),
	3656:  uint16(anon_sym_imply),
	3657:  uint16(anon_sym_range),
	3658:  uint16(anon_sym_help),
	3659:  uint16(sym_optional),
	3660:  uint16(sym_modules),
	3661:  uint16(sym_symbol),
	3662:  uint16(3),
	3663:  uint16(3),
	3664:  uint16(1),
	3665:  uint16(sym_comment),
	3666:  uint16(345),
	3667:  uint16(3),
	3669:  uint16(anon_sym_dependson),
	3670:  uint16(anon_sym_visibleif),
	3671:  uint16(343),
	3672:  uint16(31),
	3673:  uint16(anon_sym_mainmenu),
	3674:  uint16(anon_sym_config),
	3675:  uint16(anon_sym_configdefault),
	3676:  uint16(anon_sym_menuconfig),
	3677:  uint16(anon_sym_choice),
	3678:  uint16(anon_sym_comment),
	3679:  uint16(anon_sym_menu),
	3680:  uint16(anon_sym_if),
	3681:  uint16(anon_sym_source),
	3682:  uint16(anon_sym_rsource),
	3683:  uint16(anon_sym_osource),
	3684:  uint16(anon_sym_orsource),
	3685:  uint16(anon_sym_bool),
	3686:  uint16(anon_sym_tristate),
	3687:  uint16(anon_sym_int),
	3688:  uint16(anon_sym_hex),
	3689:  uint16(anon_sym_string),
	3690:  uint16(anon_sym_prompt),
	3691:  uint16(anon_sym_default),
	3692:  uint16(anon_sym_def_bool),
	3693:  uint16(anon_sym_def_tristate),
	3694:  uint16(anon_sym_def_int),
	3695:  uint16(anon_sym_def_hex),
	3696:  uint16(anon_sym_def_string),
	3697:  uint16(anon_sym_select),
	3698:  uint16(anon_sym_imply),
	3699:  uint16(anon_sym_range),
	3700:  uint16(anon_sym_help),
	3701:  uint16(sym_optional),
	3702:  uint16(sym_modules),
	3703:  uint16(sym_symbol),
	3704:  uint16(3),
	3705:  uint16(3),
	3706:  uint16(1),
	3707:  uint16(sym_comment),
	3708:  uint16(361),
	3709:  uint16(3),
	3711:  uint16(anon_sym_dependson),
	3712:  uint16(anon_sym_visibleif),
	3713:  uint16(359),
	3714:  uint16(31),
	3715:  uint16(anon_sym_mainmenu),
	3716:  uint16(anon_sym_config),
	3717:  uint16(anon_sym_configdefault),
	3718:  uint16(anon_sym_menuconfig),
	3719:  uint16(anon_sym_choice),
	3720:  uint16(anon_sym_comment),
	3721:  uint16(anon_sym_menu),
	3722:  uint16(anon_sym_if),
	3723:  uint16(anon_sym_source),
	3724:  uint16(anon_sym_rsource),
	3725:  uint16(anon_sym_osource),
	3726:  uint16(anon_sym_orsource),
	3727:  uint16(anon_sym_bool),
	3728:  uint16(anon_sym_tristate),
	3729:  uint16(anon_sym_int),
	3730:  uint16(anon_sym_hex),
	3731:  uint16(anon_sym_string),
	3732:  uint16(anon_sym_prompt),
	3733:  uint16(anon_sym_default),
	3734:  uint16(anon_sym_def_bool),
	3735:  uint16(anon_sym_def_tristate),
	3736:  uint16(anon_sym_def_int),
	3737:  uint16(anon_sym_def_hex),
	3738:  uint16(anon_sym_def_string),
	3739:  uint16(anon_sym_select),
	3740:  uint16(anon_sym_imply),
	3741:  uint16(anon_sym_range),
	3742:  uint16(anon_sym_help),
	3743:  uint16(sym_optional),
	3744:  uint16(sym_modules),
	3745:  uint16(sym_symbol),
	3746:  uint16(3),
	3747:  uint16(3),
	3748:  uint16(1),
	3749:  uint16(sym_comment),
	3750:  uint16(333),
	3751:  uint16(3),
	3753:  uint16(anon_sym_dependson),
	3754:  uint16(anon_sym_visibleif),
	3755:  uint16(331),
	3756:  uint16(31),
	3757:  uint16(anon_sym_mainmenu),
	3758:  uint16(anon_sym_config),
	3759:  uint16(anon_sym_configdefault),
	3760:  uint16(anon_sym_menuconfig),
	3761:  uint16(anon_sym_choice),
	3762:  uint16(anon_sym_comment),
	3763:  uint16(anon_sym_menu),
	3764:  uint16(anon_sym_if),
	3765:  uint16(anon_sym_source),
	3766:  uint16(anon_sym_rsource),
	3767:  uint16(anon_sym_osource),
	3768:  uint16(anon_sym_orsource),
	3769:  uint16(anon_sym_bool),
	3770:  uint16(anon_sym_tristate),
	3771:  uint16(anon_sym_int),
	3772:  uint16(anon_sym_hex),
	3773:  uint16(anon_sym_string),
	3774:  uint16(anon_sym_prompt),
	3775:  uint16(anon_sym_default),
	3776:  uint16(anon_sym_def_bool),
	3777:  uint16(anon_sym_def_tristate),
	3778:  uint16(anon_sym_def_int),
	3779:  uint16(anon_sym_def_hex),
	3780:  uint16(anon_sym_def_string),
	3781:  uint16(anon_sym_select),
	3782:  uint16(anon_sym_imply),
	3783:  uint16(anon_sym_range),
	3784:  uint16(anon_sym_help),
	3785:  uint16(sym_optional),
	3786:  uint16(sym_modules),
	3787:  uint16(sym_symbol),
	3788:  uint16(3),
	3789:  uint16(3),
	3790:  uint16(1),
	3791:  uint16(sym_comment),
	3792:  uint16(337),
	3793:  uint16(3),
	3795:  uint16(anon_sym_dependson),
	3796:  uint16(anon_sym_visibleif),
	3797:  uint16(335),
	3798:  uint16(31),
	3799:  uint16(anon_sym_mainmenu),
	3800:  uint16(anon_sym_config),
	3801:  uint16(anon_sym_configdefault),
	3802:  uint16(anon_sym_menuconfig),
	3803:  uint16(anon_sym_choice),
	3804:  uint16(anon_sym_comment),
	3805:  uint16(anon_sym_menu),
	3806:  uint16(anon_sym_if),
	3807:  uint16(anon_sym_source),
	3808:  uint16(anon_sym_rsource),
	3809:  uint16(anon_sym_osource),
	3810:  uint16(anon_sym_orsource),
	3811:  uint16(anon_sym_bool),
	3812:  uint16(anon_sym_tristate),
	3813:  uint16(anon_sym_int),
	3814:  uint16(anon_sym_hex),
	3815:  uint16(anon_sym_string),
	3816:  uint16(anon_sym_prompt),
	3817:  uint16(anon_sym_default),
	3818:  uint16(anon_sym_def_bool),
	3819:  uint16(anon_sym_def_tristate),
	3820:  uint16(anon_sym_def_int),
	3821:  uint16(anon_sym_def_hex),
	3822:  uint16(anon_sym_def_string),
	3823:  uint16(anon_sym_select),
	3824:  uint16(anon_sym_imply),
	3825:  uint16(anon_sym_range),
	3826:  uint16(anon_sym_help),
	3827:  uint16(sym_optional),
	3828:  uint16(sym_modules),
	3829:  uint16(sym_symbol),
	3830:  uint16(3),
	3831:  uint16(3),
	3832:  uint16(1),
	3833:  uint16(sym_comment),
	3834:  uint16(349),
	3835:  uint16(3),
	3837:  uint16(anon_sym_dependson),
	3838:  uint16(anon_sym_visibleif),
	3839:  uint16(347),
	3840:  uint16(31),
	3841:  uint16(anon_sym_mainmenu),
	3842:  uint16(anon_sym_config),
	3843:  uint16(anon_sym_configdefault),
	3844:  uint16(anon_sym_menuconfig),
	3845:  uint16(anon_sym_choice),
	3846:  uint16(anon_sym_comment),
	3847:  uint16(anon_sym_menu),
	3848:  uint16(anon_sym_if),
	3849:  uint16(anon_sym_source),
	3850:  uint16(anon_sym_rsource),
	3851:  uint16(anon_sym_osource),
	3852:  uint16(anon_sym_orsource),
	3853:  uint16(anon_sym_bool),
	3854:  uint16(anon_sym_tristate),
	3855:  uint16(anon_sym_int),
	3856:  uint16(anon_sym_hex),
	3857:  uint16(anon_sym_string),
	3858:  uint16(anon_sym_prompt),
	3859:  uint16(anon_sym_default),
	3860:  uint16(anon_sym_def_bool),
	3861:  uint16(anon_sym_def_tristate),
	3862:  uint16(anon_sym_def_int),
	3863:  uint16(anon_sym_def_hex),
	3864:  uint16(anon_sym_def_string),
	3865:  uint16(anon_sym_select),
	3866:  uint16(anon_sym_imply),
	3867:  uint16(anon_sym_range),
	3868:  uint16(anon_sym_help),
	3869:  uint16(sym_optional),
	3870:  uint16(sym_modules),
	3871:  uint16(sym_symbol),
	3872:  uint16(3),
	3873:  uint16(3),
	3874:  uint16(1),
	3875:  uint16(sym_comment),
	3876:  uint16(321),
	3877:  uint16(3),
	3879:  uint16(anon_sym_dependson),
	3880:  uint16(anon_sym_visibleif),
	3881:  uint16(319),
	3882:  uint16(31),
	3883:  uint16(anon_sym_mainmenu),
	3884:  uint16(anon_sym_config),
	3885:  uint16(anon_sym_configdefault),
	3886:  uint16(anon_sym_menuconfig),
	3887:  uint16(anon_sym_choice),
	3888:  uint16(anon_sym_comment),
	3889:  uint16(anon_sym_menu),
	3890:  uint16(anon_sym_if),
	3891:  uint16(anon_sym_source),
	3892:  uint16(anon_sym_rsource),
	3893:  uint16(anon_sym_osource),
	3894:  uint16(anon_sym_orsource),
	3895:  uint16(anon_sym_bool),
	3896:  uint16(anon_sym_tristate),
	3897:  uint16(anon_sym_int),
	3898:  uint16(anon_sym_hex),
	3899:  uint16(anon_sym_string),
	3900:  uint16(anon_sym_prompt),
	3901:  uint16(anon_sym_default),
	3902:  uint16(anon_sym_def_bool),
	3903:  uint16(anon_sym_def_tristate),
	3904:  uint16(anon_sym_def_int),
	3905:  uint16(anon_sym_def_hex),
	3906:  uint16(anon_sym_def_string),
	3907:  uint16(anon_sym_select),
	3908:  uint16(anon_sym_imply),
	3909:  uint16(anon_sym_range),
	3910:  uint16(anon_sym_help),
	3911:  uint16(sym_optional),
	3912:  uint16(sym_modules),
	3913:  uint16(sym_symbol),
	3914:  uint16(3),
	3915:  uint16(3),
	3916:  uint16(1),
	3917:  uint16(sym_comment),
	3918:  uint16(385),
	3919:  uint16(3),
	3921:  uint16(anon_sym_dependson),
	3922:  uint16(anon_sym_visibleif),
	3923:  uint16(383),
	3924:  uint16(31),
	3925:  uint16(anon_sym_mainmenu),
	3926:  uint16(anon_sym_config),
	3927:  uint16(anon_sym_configdefault),
	3928:  uint16(anon_sym_menuconfig),
	3929:  uint16(anon_sym_choice),
	3930:  uint16(anon_sym_comment),
	3931:  uint16(anon_sym_menu),
	3932:  uint16(anon_sym_if),
	3933:  uint16(anon_sym_source),
	3934:  uint16(anon_sym_rsource),
	3935:  uint16(anon_sym_osource),
	3936:  uint16(anon_sym_orsource),
	3937:  uint16(anon_sym_bool),
	3938:  uint16(anon_sym_tristate),
	3939:  uint16(anon_sym_int),
	3940:  uint16(anon_sym_hex),
	3941:  uint16(anon_sym_string),
	3942:  uint16(anon_sym_prompt),
	3943:  uint16(anon_sym_default),
	3944:  uint16(anon_sym_def_bool),
	3945:  uint16(anon_sym_def_tristate),
	3946:  uint16(anon_sym_def_int),
	3947:  uint16(anon_sym_def_hex),
	3948:  uint16(anon_sym_def_string),
	3949:  uint16(anon_sym_select),
	3950:  uint16(anon_sym_imply),
	3951:  uint16(anon_sym_range),
	3952:  uint16(anon_sym_help),
	3953:  uint16(sym_optional),
	3954:  uint16(sym_modules),
	3955:  uint16(sym_symbol),
	3956:  uint16(3),
	3957:  uint16(3),
	3958:  uint16(1),
	3959:  uint16(sym_comment),
	3960:  uint16(329),
	3961:  uint16(3),
	3963:  uint16(anon_sym_dependson),
	3964:  uint16(anon_sym_visibleif),
	3965:  uint16(327),
	3966:  uint16(31),
	3967:  uint16(anon_sym_mainmenu),
	3968:  uint16(anon_sym_config),
	3969:  uint16(anon_sym_configdefault),
	3970:  uint16(anon_sym_menuconfig),
	3971:  uint16(anon_sym_choice),
	3972:  uint16(anon_sym_comment),
	3973:  uint16(anon_sym_menu),
	3974:  uint16(anon_sym_if),
	3975:  uint16(anon_sym_source),
	3976:  uint16(anon_sym_rsource),
	3977:  uint16(anon_sym_osource),
	3978:  uint16(anon_sym_orsource),
	3979:  uint16(anon_sym_bool),
	3980:  uint16(anon_sym_tristate),
	3981:  uint16(anon_sym_int),
	3982:  uint16(anon_sym_hex),
	3983:  uint16(anon_sym_string),
	3984:  uint16(anon_sym_prompt),
	3985:  uint16(anon_sym_default),
	3986:  uint16(anon_sym_def_bool),
	3987:  uint16(anon_sym_def_tristate),
	3988:  uint16(anon_sym_def_int),
	3989:  uint16(anon_sym_def_hex),
	3990:  uint16(anon_sym_def_string),
	3991:  uint16(anon_sym_select),
	3992:  uint16(anon_sym_imply),
	3993:  uint16(anon_sym_range),
	3994:  uint16(anon_sym_help),
	3995:  uint16(sym_optional),
	3996:  uint16(sym_modules),
	3997:  uint16(sym_symbol),
	3998:  uint16(3),
	3999:  uint16(3),
	4000:  uint16(1),
	4001:  uint16(sym_comment),
	4002:  uint16(365),
	4003:  uint16(3),
	4005:  uint16(anon_sym_dependson),
	4006:  uint16(anon_sym_visibleif),
	4007:  uint16(363),
	4008:  uint16(31),
	4009:  uint16(anon_sym_mainmenu),
	4010:  uint16(anon_sym_config),
	4011:  uint16(anon_sym_configdefault),
	4012:  uint16(anon_sym_menuconfig),
	4013:  uint16(anon_sym_choice),
	4014:  uint16(anon_sym_comment),
	4015:  uint16(anon_sym_menu),
	4016:  uint16(anon_sym_if),
	4017:  uint16(anon_sym_source),
	4018:  uint16(anon_sym_rsource),
	4019:  uint16(anon_sym_osource),
	4020:  uint16(anon_sym_orsource),
	4021:  uint16(anon_sym_bool),
	4022:  uint16(anon_sym_tristate),
	4023:  uint16(anon_sym_int),
	4024:  uint16(anon_sym_hex),
	4025:  uint16(anon_sym_string),
	4026:  uint16(anon_sym_prompt),
	4027:  uint16(anon_sym_default),
	4028:  uint16(anon_sym_def_bool),
	4029:  uint16(anon_sym_def_tristate),
	4030:  uint16(anon_sym_def_int),
	4031:  uint16(anon_sym_def_hex),
	4032:  uint16(anon_sym_def_string),
	4033:  uint16(anon_sym_select),
	4034:  uint16(anon_sym_imply),
	4035:  uint16(anon_sym_range),
	4036:  uint16(anon_sym_help),
	4037:  uint16(sym_optional),
	4038:  uint16(sym_modules),
	4039:  uint16(sym_symbol),
	4040:  uint16(18),
	4041:  uint16(3),
	4042:  uint16(1),
	4043:  uint16(sym_comment),
	4044:  uint16(27),
	4045:  uint16(1),
	4046:  uint16(sym_symbol),
	4047:  uint16(29),
	4048:  uint16(1),
	4049:  uint16(anon_sym_mainmenu),
	4050:  uint16(31),
	4051:  uint16(1),
	4052:  uint16(anon_sym_config),
	4053:  uint16(33),
	4054:  uint16(1),
	4055:  uint16(anon_sym_configdefault),
	4056:  uint16(35),
	4057:  uint16(1),
	4058:  uint16(anon_sym_menuconfig),
	4059:  uint16(37),
	4060:  uint16(1),
	4061:  uint16(anon_sym_choice),
	4062:  uint16(39),
	4063:  uint16(1),
	4064:  uint16(anon_sym_comment),
	4065:  uint16(41),
	4066:  uint16(1),
	4067:  uint16(anon_sym_menu),
	4068:  uint16(45),
	4069:  uint16(1),
	4070:  uint16(anon_sym_if),
	4071:  uint16(387),
	4072:  uint16(1),
	4073:  uint16(anon_sym_endif),
	4074:  uint16(389),
	4075:  uint16(1),
	4076:  uint16(anon_sym_EQ),
	4077:  uint16(391),
	4078:  uint16(1),
	4079:  uint16(anon_sym_PIPE_PIPE),
	4080:  uint16(393),
	4081:  uint16(1),
	4082:  uint16(anon_sym_AMP_AMP),
	4083:  uint16(397),
	4084:  uint16(2),
	4085:  uint16(anon_sym_LT),
	4086:  uint16(anon_sym_GT),
	4087:  uint16(395),
	4088:  uint16(3),
	4089:  uint16(anon_sym_BANG_EQ),
	4090:  uint16(anon_sym_LT_EQ),
	4091:  uint16(anon_sym_GT_EQ),
	4092:  uint16(47),
	4093:  uint16(4),
	4094:  uint16(anon_sym_source),
	4095:  uint16(anon_sym_rsource),
	4096:  uint16(anon_sym_osource),
	4097:  uint16(anon_sym_orsource),
	4098:  uint16(110),
	4099:  uint16(12),
	4100:  uint16(sym__entry),
	4101:  uint16(sym_mainmenu),
	4102:  uint16(sym_config),
	4103:  uint16(sym_configdefault),
	4104:  uint16(sym_menuconfig),
	4105:  uint16(sym_choice),
	4106:  uint16(sym_comment_entry),
	4107:  uint16(sym_menu),
	4108:  uint16(sym_if),
	4109:  uint16(sym_source),
	4110:  uint16(sym_variable),
	4111:  uint16(aux_sym_configuration_repeat1),
	4112:  uint16(18),
	4113:  uint16(3),
	4114:  uint16(1),
	4115:  uint16(sym_comment),
	4116:  uint16(27),
	4117:  uint16(1),
	4118:  uint16(sym_symbol),
	4119:  uint16(29),
	4120:  uint16(1),
	4121:  uint16(anon_sym_mainmenu),
	4122:  uint16(31),
	4123:  uint16(1),
	4124:  uint16(anon_sym_config),
	4125:  uint16(33),
	4126:  uint16(1),
	4127:  uint16(anon_sym_configdefault),
	4128:  uint16(35),
	4129:  uint16(1),
	4130:  uint16(anon_sym_menuconfig),
	4131:  uint16(37),
	4132:  uint16(1),
	4133:  uint16(anon_sym_choice),
	4134:  uint16(39),
	4135:  uint16(1),
	4136:  uint16(anon_sym_comment),
	4137:  uint16(41),
	4138:  uint16(1),
	4139:  uint16(anon_sym_menu),
	4140:  uint16(45),
	4141:  uint16(1),
	4142:  uint16(anon_sym_if),
	4143:  uint16(389),
	4144:  uint16(1),
	4145:  uint16(anon_sym_EQ),
	4146:  uint16(391),
	4147:  uint16(1),
	4148:  uint16(anon_sym_PIPE_PIPE),
	4149:  uint16(393),
	4150:  uint16(1),
	4151:  uint16(anon_sym_AMP_AMP),
	4152:  uint16(399),
	4153:  uint16(1),
	4154:  uint16(anon_sym_endif),
	4155:  uint16(397),
	4156:  uint16(2),
	4157:  uint16(anon_sym_LT),
	4158:  uint16(anon_sym_GT),
	4159:  uint16(395),
	4160:  uint16(3),
	4161:  uint16(anon_sym_BANG_EQ),
	4162:  uint16(anon_sym_LT_EQ),
	4163:  uint16(anon_sym_GT_EQ),
	4164:  uint16(47),
	4165:  uint16(4),
	4166:  uint16(anon_sym_source),
	4167:  uint16(anon_sym_rsource),
	4168:  uint16(anon_sym_osource),
	4169:  uint16(anon_sym_orsource),
	4170:  uint16(103),
	4171:  uint16(12),
	4172:  uint16(sym__entry),
	4173:  uint16(sym_mainmenu),
	4174:  uint16(sym_config),
	4175:  uint16(sym_configdefault),
	4176:  uint16(sym_menuconfig),
	4177:  uint16(sym_choice),
	4178:  uint16(sym_comment_entry),
	4179:  uint16(sym_menu),
	4180:  uint16(sym_if),
	4181:  uint16(sym_source),
	4182:  uint16(sym_variable),
	4183:  uint16(aux_sym_configuration_repeat1),
	4184:  uint16(3),
	4185:  uint16(3),
	4186:  uint16(1),
	4187:  uint16(sym_comment),
	4188:  uint16(353),
	4189:  uint16(3),
	4191:  uint16(anon_sym_dependson),
	4192:  uint16(anon_sym_visibleif),
	4193:  uint16(351),
	4194:  uint16(31),
	4195:  uint16(anon_sym_mainmenu),
	4196:  uint16(anon_sym_config),
	4197:  uint16(anon_sym_configdefault),
	4198:  uint16(anon_sym_menuconfig),
	4199:  uint16(anon_sym_choice),
	4200:  uint16(anon_sym_comment),
	4201:  uint16(anon_sym_menu),
	4202:  uint16(anon_sym_if),
	4203:  uint16(anon_sym_source),
	4204:  uint16(anon_sym_rsource),
	4205:  uint16(anon_sym_osource),
	4206:  uint16(anon_sym_orsource),
	4207:  uint16(anon_sym_bool),
	4208:  uint16(anon_sym_tristate),
	4209:  uint16(anon_sym_int),
	4210:  uint16(anon_sym_hex),
	4211:  uint16(anon_sym_string),
	4212:  uint16(anon_sym_prompt),
	4213:  uint16(anon_sym_default),
	4214:  uint16(anon_sym_def_bool),
	4215:  uint16(anon_sym_def_tristate),
	4216:  uint16(anon_sym_def_int),
	4217:  uint16(anon_sym_def_hex),
	4218:  uint16(anon_sym_def_string),
	4219:  uint16(anon_sym_select),
	4220:  uint16(anon_sym_imply),
	4221:  uint16(anon_sym_range),
	4222:  uint16(anon_sym_help),
	4223:  uint16(sym_optional),
	4224:  uint16(sym_modules),
	4225:  uint16(sym_symbol),
	4226:  uint16(13),
	4227:  uint16(3),
	4228:  uint16(1),
	4229:  uint16(sym_comment),
	4230:  uint16(57),
	4231:  uint16(1),
	4232:  uint16(anon_sym_dependson),
	4233:  uint16(63),
	4234:  uint16(1),
	4235:  uint16(anon_sym_visibleif),
	4236:  uint16(403),
	4237:  uint16(1),
	4238:  uint16(anon_sym_prompt),
	4239:  uint16(405),
	4240:  uint16(1),
	4241:  uint16(anon_sym_default),
	4242:  uint16(409),
	4243:  uint16(1),
	4244:  uint16(anon_sym_select),
	4245:  uint16(411),
	4246:  uint16(1),
	4247:  uint16(anon_sym_imply),
	4248:  uint16(413),
	4249:  uint16(1),
	4250:  uint16(anon_sym_range),
	4251:  uint16(415),
	4252:  uint16(1),
	4253:  uint16(anon_sym_help),
	4254:  uint16(417),
	4255:  uint16(2),
	4256:  uint16(sym_optional),
	4257:  uint16(sym_modules),
	4258:  uint16(401),
	4259:  uint16(5),
	4260:  uint16(anon_sym_bool),
	4261:  uint16(anon_sym_tristate),
	4262:  uint16(anon_sym_int),
	4263:  uint16(anon_sym_hex),
	4264:  uint16(anon_sym_string),
	4265:  uint16(407),
	4266:  uint16(5),
	4267:  uint16(anon_sym_def_bool),
	4268:  uint16(anon_sym_def_tristate),
	4269:  uint16(anon_sym_def_int),
	4270:  uint16(anon_sym_def_hex),
	4271:  uint16(anon_sym_def_string),
	4272:  uint16(8),
	4273:  uint16(12),
	4274:  uint16(sym__config_option),
	4275:  uint16(sym_type_definition),
	4276:  uint16(sym_input_prompt),
	4277:  uint16(sym_default_value),
	4278:  uint16(sym_type_definition_default),
	4279:  uint16(sym_dependencies),
	4280:  uint16(sym_reverse_dependencies),
	4281:  uint16(sym_weak_reverse_dependencies),
	4282:  uint16(sym_limiting_menu_display),
	4283:  uint16(sym_numerical_ranges),
	4284:  uint16(sym_help_text),
	4285:  uint16(aux_sym_config_repeat1),
	4286:  uint16(13),
	4287:  uint16(3),
	4288:  uint16(1),
	4289:  uint16(sym_comment),
	4290:  uint16(57),
	4291:  uint16(1),
	4292:  uint16(anon_sym_dependson),
	4293:  uint16(63),
	4294:  uint16(1),
	4295:  uint16(anon_sym_visibleif),
	4296:  uint16(403),
	4297:  uint16(1),
	4298:  uint16(anon_sym_prompt),
	4299:  uint16(405),
	4300:  uint16(1),
	4301:  uint16(anon_sym_default),
	4302:  uint16(409),
	4303:  uint16(1),
	4304:  uint16(anon_sym_select),
	4305:  uint16(411),
	4306:  uint16(1),
	4307:  uint16(anon_sym_imply),
	4308:  uint16(413),
	4309:  uint16(1),
	4310:  uint16(anon_sym_range),
	4311:  uint16(415),
	4312:  uint16(1),
	4313:  uint16(anon_sym_help),
	4314:  uint16(419),
	4315:  uint16(2),
	4316:  uint16(sym_optional),
	4317:  uint16(sym_modules),
	4318:  uint16(401),
	4319:  uint16(5),
	4320:  uint16(anon_sym_bool),
	4321:  uint16(anon_sym_tristate),
	4322:  uint16(anon_sym_int),
	4323:  uint16(anon_sym_hex),
	4324:  uint16(anon_sym_string),
	4325:  uint16(407),
	4326:  uint16(5),
	4327:  uint16(anon_sym_def_bool),
	4328:  uint16(anon_sym_def_tristate),
	4329:  uint16(anon_sym_def_int),
	4330:  uint16(anon_sym_def_hex),
	4331:  uint16(anon_sym_def_string),
	4332:  uint16(4),
	4333:  uint16(12),
	4334:  uint16(sym__config_option),
	4335:  uint16(sym_type_definition),
	4336:  uint16(sym_input_prompt),
	4337:  uint16(sym_default_value),
	4338:  uint16(sym_type_definition_default),
	4339:  uint16(sym_dependencies),
	4340:  uint16(sym_reverse_dependencies),
	4341:  uint16(sym_weak_reverse_dependencies),
	4342:  uint16(sym_limiting_menu_display),
	4343:  uint16(sym_numerical_ranges),
	4344:  uint16(sym_help_text),
	4345:  uint16(aux_sym_config_repeat1),
	4346:  uint16(13),
	4347:  uint16(3),
	4348:  uint16(1),
	4349:  uint16(sym_comment),
	4350:  uint16(57),
	4351:  uint16(1),
	4352:  uint16(anon_sym_dependson),
	4353:  uint16(63),
	4354:  uint16(1),
	4355:  uint16(anon_sym_visibleif),
	4356:  uint16(403),
	4357:  uint16(1),
	4358:  uint16(anon_sym_prompt),
	4359:  uint16(405),
	4360:  uint16(1),
	4361:  uint16(anon_sym_default),
	4362:  uint16(409),
	4363:  uint16(1),
	4364:  uint16(anon_sym_select),
	4365:  uint16(411),
	4366:  uint16(1),
	4367:  uint16(anon_sym_imply),
	4368:  uint16(413),
	4369:  uint16(1),
	4370:  uint16(anon_sym_range),
	4371:  uint16(415),
	4372:  uint16(1),
	4373:  uint16(anon_sym_help),
	4374:  uint16(421),
	4375:  uint16(2),
	4376:  uint16(sym_optional),
	4377:  uint16(sym_modules),
	4378:  uint16(401),
	4379:  uint16(5),
	4380:  uint16(anon_sym_bool),
	4381:  uint16(anon_sym_tristate),
	4382:  uint16(anon_sym_int),
	4383:  uint16(anon_sym_hex),
	4384:  uint16(anon_sym_string),
	4385:  uint16(407),
	4386:  uint16(5),
	4387:  uint16(anon_sym_def_bool),
	4388:  uint16(anon_sym_def_tristate),
	4389:  uint16(anon_sym_def_int),
	4390:  uint16(anon_sym_def_hex),
	4391:  uint16(anon_sym_def_string),
	4392:  uint16(14),
	4393:  uint16(12),
	4394:  uint16(sym__config_option),
	4395:  uint16(sym_type_definition),
	4396:  uint16(sym_input_prompt),
	4397:  uint16(sym_default_value),
	4398:  uint16(sym_type_definition_default),
	4399:  uint16(sym_dependencies),
	4400:  uint16(sym_reverse_dependencies),
	4401:  uint16(sym_weak_reverse_dependencies),
	4402:  uint16(sym_limiting_menu_display),
	4403:  uint16(sym_numerical_ranges),
	4404:  uint16(sym_help_text),
	4405:  uint16(aux_sym_config_repeat1),
	4406:  uint16(13),
	4407:  uint16(3),
	4408:  uint16(1),
	4409:  uint16(sym_comment),
	4410:  uint16(251),
	4411:  uint16(1),
	4412:  uint16(anon_sym_dependson),
	4413:  uint16(257),
	4414:  uint16(1),
	4415:  uint16(anon_sym_visibleif),
	4416:  uint16(425),
	4417:  uint16(1),
	4418:  uint16(anon_sym_prompt),
	4419:  uint16(427),
	4420:  uint16(1),
	4421:  uint16(anon_sym_default),
	4422:  uint16(431),
	4423:  uint16(1),
	4424:  uint16(anon_sym_select),
	4425:  uint16(433),
	4426:  uint16(1),
	4427:  uint16(anon_sym_imply),
	4428:  uint16(435),
	4429:  uint16(1),
	4430:  uint16(anon_sym_range),
	4431:  uint16(437),
	4432:  uint16(1),
	4433:  uint16(anon_sym_help),
	4434:  uint16(439),
	4435:  uint16(2),
	4436:  uint16(sym_optional),
	4437:  uint16(sym_modules),
	4438:  uint16(423),
	4439:  uint16(5),
	4440:  uint16(anon_sym_bool),
	4441:  uint16(anon_sym_tristate),
	4442:  uint16(anon_sym_int),
	4443:  uint16(anon_sym_hex),
	4444:  uint16(anon_sym_string),
	4445:  uint16(429),
	4446:  uint16(5),
	4447:  uint16(anon_sym_def_bool),
	4448:  uint16(anon_sym_def_tristate),
	4449:  uint16(anon_sym_def_int),
	4450:  uint16(anon_sym_def_hex),
	4451:  uint16(anon_sym_def_string),
	4452:  uint16(25),
	4453:  uint16(12),
	4454:  uint16(sym__config_option),
	4455:  uint16(sym_type_definition),
	4456:  uint16(sym_input_prompt),
	4457:  uint16(sym_default_value),
	4458:  uint16(sym_type_definition_default),
	4459:  uint16(sym_dependencies),
	4460:  uint16(sym_reverse_dependencies),
	4461:  uint16(sym_weak_reverse_dependencies),
	4462:  uint16(sym_limiting_menu_display),
	4463:  uint16(sym_numerical_ranges),
	4464:  uint16(sym_help_text),
	4465:  uint16(aux_sym_config_repeat1),
	4466:  uint16(13),
	4467:  uint16(3),
	4468:  uint16(1),
	4469:  uint16(sym_comment),
	4470:  uint16(57),
	4471:  uint16(1),
	4472:  uint16(anon_sym_dependson),
	4473:  uint16(63),
	4474:  uint16(1),
	4475:  uint16(anon_sym_visibleif),
	4476:  uint16(403),
	4477:  uint16(1),
	4478:  uint16(anon_sym_prompt),
	4479:  uint16(405),
	4480:  uint16(1),
	4481:  uint16(anon_sym_default),
	4482:  uint16(409),
	4483:  uint16(1),
	4484:  uint16(anon_sym_select),
	4485:  uint16(411),
	4486:  uint16(1),
	4487:  uint16(anon_sym_imply),
	4488:  uint16(413),
	4489:  uint16(1),
	4490:  uint16(anon_sym_range),
	4491:  uint16(415),
	4492:  uint16(1),
	4493:  uint16(anon_sym_help),
	4494:  uint16(441),
	4495:  uint16(2),
	4496:  uint16(sym_optional),
	4497:  uint16(sym_modules),
	4498:  uint16(401),
	4499:  uint16(5),
	4500:  uint16(anon_sym_bool),
	4501:  uint16(anon_sym_tristate),
	4502:  uint16(anon_sym_int),
	4503:  uint16(anon_sym_hex),
	4504:  uint16(anon_sym_string),
	4505:  uint16(407),
	4506:  uint16(5),
	4507:  uint16(anon_sym_def_bool),
	4508:  uint16(anon_sym_def_tristate),
	4509:  uint16(anon_sym_def_int),
	4510:  uint16(anon_sym_def_hex),
	4511:  uint16(anon_sym_def_string),
	4512:  uint16(15),
	4513:  uint16(12),
	4514:  uint16(sym__config_option),
	4515:  uint16(sym_type_definition),
	4516:  uint16(sym_input_prompt),
	4517:  uint16(sym_default_value),
	4518:  uint16(sym_type_definition_default),
	4519:  uint16(sym_dependencies),
	4520:  uint16(sym_reverse_dependencies),
	4521:  uint16(sym_weak_reverse_dependencies),
	4522:  uint16(sym_limiting_menu_display),
	4523:  uint16(sym_numerical_ranges),
	4524:  uint16(sym_help_text),
	4525:  uint16(aux_sym_config_repeat1),
	4526:  uint16(13),
	4527:  uint16(3),
	4528:  uint16(1),
	4529:  uint16(sym_comment),
	4530:  uint16(251),
	4531:  uint16(1),
	4532:  uint16(anon_sym_dependson),
	4533:  uint16(257),
	4534:  uint16(1),
	4535:  uint16(anon_sym_visibleif),
	4536:  uint16(425),
	4537:  uint16(1),
	4538:  uint16(anon_sym_prompt),
	4539:  uint16(427),
	4540:  uint16(1),
	4541:  uint16(anon_sym_default),
	4542:  uint16(431),
	4543:  uint16(1),
	4544:  uint16(anon_sym_select),
	4545:  uint16(433),
	4546:  uint16(1),
	4547:  uint16(anon_sym_imply),
	4548:  uint16(435),
	4549:  uint16(1),
	4550:  uint16(anon_sym_range),
	4551:  uint16(437),
	4552:  uint16(1),
	4553:  uint16(anon_sym_help),
	4554:  uint16(443),
	4555:  uint16(2),
	4556:  uint16(sym_optional),
	4557:  uint16(sym_modules),
	4558:  uint16(423),
	4559:  uint16(5),
	4560:  uint16(anon_sym_bool),
	4561:  uint16(anon_sym_tristate),
	4562:  uint16(anon_sym_int),
	4563:  uint16(anon_sym_hex),
	4564:  uint16(anon_sym_string),
	4565:  uint16(429),
	4566:  uint16(5),
	4567:  uint16(anon_sym_def_bool),
	4568:  uint16(anon_sym_def_tristate),
	4569:  uint16(anon_sym_def_int),
	4570:  uint16(anon_sym_def_hex),
	4571:  uint16(anon_sym_def_string),
	4572:  uint16(26),
	4573:  uint16(12),
	4574:  uint16(sym__config_option),
	4575:  uint16(sym_type_definition),
	4576:  uint16(sym_input_prompt),
	4577:  uint16(sym_default_value),
	4578:  uint16(sym_type_definition_default),
	4579:  uint16(sym_dependencies),
	4580:  uint16(sym_reverse_dependencies),
	4581:  uint16(sym_weak_reverse_dependencies),
	4582:  uint16(sym_limiting_menu_display),
	4583:  uint16(sym_numerical_ranges),
	4584:  uint16(sym_help_text),
	4585:  uint16(aux_sym_config_repeat1),
	4586:  uint16(8),
	4587:  uint16(3),
	4588:  uint16(1),
	4589:  uint16(sym_comment),
	4590:  uint16(445),
	4591:  uint16(1),
	4592:  uint16(sym_symbol),
	4593:  uint16(448),
	4594:  uint16(1),
	4595:  uint16(anon_sym_DOLLAR_LPAREN),
	4596:  uint16(451),
	4597:  uint16(1),
	4598:  uint16(anon_sym_DQUOTE),
	4599:  uint16(454),
	4600:  uint16(1),
	4601:  uint16(anon_sym_SQUOTE),
	4602:  uint16(94),
	4603:  uint16(3),
	4604:  uint16(sym_macro_variable),
	4605:  uint16(sym_string),
	4606:  uint16(aux_sym_name_repeat1),
	4607:  uint16(106),
	4608:  uint16(7),
	4609:  uint16(anon_sym_EQ),
	4610:  uint16(anon_sym_PIPE_PIPE),
	4611:  uint16(anon_sym_AMP_AMP),
	4612:  uint16(anon_sym_BANG_EQ),
	4613:  uint16(anon_sym_LT_EQ),
	4614:  uint16(anon_sym_GT_EQ),
	4615:  uint16(anon_sym_RPAREN),
	4616:  uint16(104),
	4617:  uint16(16),
	4618:  uint16(anon_sym_mainmenu),
	4619:  uint16(anon_sym_config),
	4620:  uint16(anon_sym_configdefault),
	4621:  uint16(anon_sym_menuconfig),
	4622:  uint16(anon_sym_choice),
	4623:  uint16(anon_sym_comment),
	4624:  uint16(anon_sym_menu),
	4625:  uint16(anon_sym_if),
	4626:  uint16(anon_sym_endif),
	4627:  uint16(anon_sym_source),
	4628:  uint16(anon_sym_rsource),
	4629:  uint16(anon_sym_osource),
	4630:  uint16(anon_sym_orsource),
	4631:  uint16(anon_sym_default),
	4632:  uint16(anon_sym_LT),
	4633:  uint16(anon_sym_GT),
	4634:  uint16(8),
	4635:  uint16(3),
	4636:  uint16(1),
	4637:  uint16(sym_comment),
	4638:  uint16(457),
	4639:  uint16(1),
	4640:  uint16(sym_symbol),
	4641:  uint16(459),
	4642:  uint16(1),
	4643:  uint16(anon_sym_DOLLAR_LPAREN),
	4644:  uint16(461),
	4645:  uint16(1),
	4646:  uint16(anon_sym_DQUOTE),
	4647:  uint16(463),
	4648:  uint16(1),
	4649:  uint16(anon_sym_SQUOTE),
	4650:  uint16(94),
	4651:  uint16(3),
	4652:  uint16(sym_macro_variable),
	4653:  uint16(sym_string),
	4654:  uint16(aux_sym_name_repeat1),
	4655:  uint16(93),
	4656:  uint16(7),
	4657:  uint16(anon_sym_EQ),
	4658:  uint16(anon_sym_PIPE_PIPE),
	4659:  uint16(anon_sym_AMP_AMP),
	4660:  uint16(anon_sym_BANG_EQ),
	4661:  uint16(anon_sym_LT_EQ),
	4662:  uint16(anon_sym_GT_EQ),
	4663:  uint16(anon_sym_RPAREN),
	4664:  uint16(91),
	4665:  uint16(16),
	4666:  uint16(anon_sym_mainmenu),
	4667:  uint16(anon_sym_config),
	4668:  uint16(anon_sym_configdefault),
	4669:  uint16(anon_sym_menuconfig),
	4670:  uint16(anon_sym_choice),
	4671:  uint16(anon_sym_comment),
	4672:  uint16(anon_sym_menu),
	4673:  uint16(anon_sym_if),
	4674:  uint16(anon_sym_endif),
	4675:  uint16(anon_sym_source),
	4676:  uint16(anon_sym_rsource),
	4677:  uint16(anon_sym_osource),
	4678:  uint16(anon_sym_orsource),
	4679:  uint16(anon_sym_default),
	4680:  uint16(anon_sym_LT),
	4681:  uint16(anon_sym_GT),
	4682:  uint16(3),
	4683:  uint16(3),
	4684:  uint16(1),
	4685:  uint16(sym_comment),
	4686:  uint16(184),
	4687:  uint16(10),
	4688:  uint16(anon_sym_EQ),
	4689:  uint16(anon_sym_PIPE_PIPE),
	4690:  uint16(anon_sym_AMP_AMP),
	4691:  uint16(anon_sym_BANG_EQ),
	4692:  uint16(anon_sym_LT_EQ),
	4693:  uint16(anon_sym_GT_EQ),
	4694:  uint16(anon_sym_RPAREN),
	4695:  uint16(anon_sym_DOLLAR_LPAREN),
	4696:  uint16(anon_sym_DQUOTE),
	4697:  uint16(anon_sym_SQUOTE),
	4698:  uint16(182),
	4699:  uint16(19),
	4700:  uint16(anon_sym_mainmenu),
	4701:  uint16(anon_sym_config),
	4702:  uint16(anon_sym_configdefault),
	4703:  uint16(anon_sym_menuconfig),
	4704:  uint16(anon_sym_choice),
	4705:  uint16(anon_sym_endchoice),
	4706:  uint16(anon_sym_comment),
	4707:  uint16(anon_sym_menu),
	4708:  uint16(anon_sym_endmenu),
	4709:  uint16(anon_sym_if),
	4710:  uint16(anon_sym_endif),
	4711:  uint16(anon_sym_source),
	4712:  uint16(anon_sym_rsource),
	4713:  uint16(anon_sym_osource),
	4714:  uint16(anon_sym_orsource),
	4715:  uint16(anon_sym_default),
	4716:  uint16(anon_sym_LT),
	4717:  uint16(anon_sym_GT),
	4718:  uint16(sym_symbol),
	4719:  uint16(3),
	4720:  uint16(3),
	4721:  uint16(1),
	4722:  uint16(sym_comment),
	4723:  uint16(200),
	4724:  uint16(10),
	4725:  uint16(anon_sym_EQ),
	4726:  uint16(anon_sym_PIPE_PIPE),
	4727:  uint16(anon_sym_AMP_AMP),
	4728:  uint16(anon_sym_BANG_EQ),
	4729:  uint16(anon_sym_LT_EQ),
	4730:  uint16(anon_sym_GT_EQ),
	4731:  uint16(anon_sym_RPAREN),
	4732:  uint16(anon_sym_DOLLAR_LPAREN),
	4733:  uint16(anon_sym_DQUOTE),
	4734:  uint16(anon_sym_SQUOTE),
	4735:  uint16(198),
	4736:  uint16(19),
	4737:  uint16(anon_sym_mainmenu),
	4738:  uint16(anon_sym_config),
	4739:  uint16(anon_sym_configdefault),
	4740:  uint16(anon_sym_menuconfig),
	4741:  uint16(anon_sym_choice),
	4742:  uint16(anon_sym_endchoice),
	4743:  uint16(anon_sym_comment),
	4744:  uint16(anon_sym_menu),
	4745:  uint16(anon_sym_endmenu),
	4746:  uint16(anon_sym_if),
	4747:  uint16(anon_sym_endif),
	4748:  uint16(anon_sym_source),
	4749:  uint16(anon_sym_rsource),
	4750:  uint16(anon_sym_osource),
	4751:  uint16(anon_sym_orsource),
	4752:  uint16(anon_sym_default),
	4753:  uint16(anon_sym_LT),
	4754:  uint16(anon_sym_GT),
	4755:  uint16(sym_symbol),
	4756:  uint16(13),
	4757:  uint16(3),
	4758:  uint16(1),
	4759:  uint16(sym_comment),
	4760:  uint16(465),
	4761:  uint16(1),
	4762:  uint16(sym_symbol),
	4763:  uint16(468),
	4764:  uint16(1),
	4765:  uint16(anon_sym_mainmenu),
	4766:  uint16(471),
	4767:  uint16(1),
	4768:  uint16(anon_sym_config),
	4769:  uint16(474),
	4770:  uint16(1),
	4771:  uint16(anon_sym_configdefault),
	4772:  uint16(477),
	4773:  uint16(1),
	4774:  uint16(anon_sym_menuconfig),
	4775:  uint16(480),
	4776:  uint16(1),
	4777:  uint16(anon_sym_choice),
	4778:  uint16(485),
	4779:  uint16(1),
	4780:  uint16(anon_sym_comment),
	4781:  uint16(488),
	4782:  uint16(1),
	4783:  uint16(anon_sym_menu),
	4784:  uint16(491),
	4785:  uint16(1),
	4786:  uint16(anon_sym_if),
	4787:  uint16(483),
	4788:  uint16(3),
	4789:  uint16(anon_sym_endchoice),
	4790:  uint16(anon_sym_endmenu),
	4791:  uint16(anon_sym_endif),
	4792:  uint16(494),
	4793:  uint16(4),
	4794:  uint16(anon_sym_source),
	4795:  uint16(anon_sym_rsource),
	4796:  uint16(anon_sym_osource),
	4797:  uint16(anon_sym_orsource),
	4798:  uint16(98),
	4799:  uint16(12),
	4800:  uint16(sym__entry),
	4801:  uint16(sym_mainmenu),
	4802:  uint16(sym_config),
	4803:  uint16(sym_configdefault),
	4804:  uint16(sym_menuconfig),
	4805:  uint16(sym_choice),
	4806:  uint16(sym_comment_entry),
	4807:  uint16(sym_menu),
	4808:  uint16(sym_if),
	4809:  uint16(sym_source),
	4810:  uint16(sym_variable),
	4811:  uint16(aux_sym_configuration_repeat1),
	4812:  uint16(3),
	4813:  uint16(3),
	4814:  uint16(1),
	4815:  uint16(sym_comment),
	4816:  uint16(188),
	4817:  uint16(10),
	4818:  uint16(anon_sym_EQ),
	4819:  uint16(anon_sym_PIPE_PIPE),
	4820:  uint16(anon_sym_AMP_AMP),
	4821:  uint16(anon_sym_BANG_EQ),
	4822:  uint16(anon_sym_LT_EQ),
	4823:  uint16(anon_sym_GT_EQ),
	4824:  uint16(anon_sym_RPAREN),
	4825:  uint16(anon_sym_DOLLAR_LPAREN),
	4826:  uint16(anon_sym_DQUOTE),
	4827:  uint16(anon_sym_SQUOTE),
	4828:  uint16(186),
	4829:  uint16(17),
	4830:  uint16(anon_sym_mainmenu),
	4831:  uint16(anon_sym_config),
	4832:  uint16(anon_sym_configdefault),
	4833:  uint16(anon_sym_menuconfig),
	4834:  uint16(anon_sym_choice),
	4835:  uint16(anon_sym_comment),
	4836:  uint16(anon_sym_menu),
	4837:  uint16(anon_sym_if),
	4838:  uint16(anon_sym_endif),
	4839:  uint16(anon_sym_source),
	4840:  uint16(anon_sym_rsource),
	4841:  uint16(anon_sym_osource),
	4842:  uint16(anon_sym_orsource),
	4843:  uint16(anon_sym_default),
	4844:  uint16(anon_sym_LT),
	4845:  uint16(anon_sym_GT),
	4846:  uint16(sym_symbol),
	4847:  uint16(3),
	4848:  uint16(3),
	4849:  uint16(1),
	4850:  uint16(sym_comment),
	4851:  uint16(204),
	4852:  uint16(10),
	4853:  uint16(anon_sym_EQ),
	4854:  uint16(anon_sym_PIPE_PIPE),
	4855:  uint16(anon_sym_AMP_AMP),
	4856:  uint16(anon_sym_BANG_EQ),
	4857:  uint16(anon_sym_LT_EQ),
	4858:  uint16(anon_sym_GT_EQ),
	4859:  uint16(anon_sym_RPAREN),
	4860:  uint16(anon_sym_DOLLAR_LPAREN),
	4861:  uint16(anon_sym_DQUOTE),
	4862:  uint16(anon_sym_SQUOTE),
	4863:  uint16(202),
	4864:  uint16(17),
	4865:  uint16(anon_sym_mainmenu),
	4866:  uint16(anon_sym_config),
	4867:  uint16(anon_sym_configdefault),
	4868:  uint16(anon_sym_menuconfig),
	4869:  uint16(anon_sym_choice),
	4870:  uint16(anon_sym_comment),
	4871:  uint16(anon_sym_menu),
	4872:  uint16(anon_sym_if),
	4873:  uint16(anon_sym_endif),
	4874:  uint16(anon_sym_source),
	4875:  uint16(anon_sym_rsource),
	4876:  uint16(anon_sym_osource),
	4877:  uint16(anon_sym_orsource),
	4878:  uint16(anon_sym_default),
	4879:  uint16(anon_sym_LT),
	4880:  uint16(anon_sym_GT),
	4881:  uint16(sym_symbol),
	4882:  uint16(13),
	4883:  uint16(3),
	4884:  uint16(1),
	4885:  uint16(sym_comment),
	4886:  uint16(27),
	4887:  uint16(1),
	4888:  uint16(sym_symbol),
	4889:  uint16(29),
	4890:  uint16(1),
	4891:  uint16(anon_sym_mainmenu),
	4892:  uint16(31),
	4893:  uint16(1),
	4894:  uint16(anon_sym_config),
	4895:  uint16(33),
	4896:  uint16(1),
	4897:  uint16(anon_sym_configdefault),
	4898:  uint16(35),
	4899:  uint16(1),
	4900:  uint16(anon_sym_menuconfig),
	4901:  uint16(37),
	4902:  uint16(1),
	4903:  uint16(anon_sym_choice),
	4904:  uint16(39),
	4905:  uint16(1),
	4906:  uint16(anon_sym_comment),
	4907:  uint16(41),
	4908:  uint16(1),
	4909:  uint16(anon_sym_menu),
	4910:  uint16(45),
	4911:  uint16(1),
	4912:  uint16(anon_sym_if),
	4913:  uint16(497),
	4914:  uint16(1),
	4915:  uint16(anon_sym_endmenu),
	4916:  uint16(47),
	4917:  uint16(4),
	4918:  uint16(anon_sym_source),
	4919:  uint16(anon_sym_rsource),
	4920:  uint16(anon_sym_osource),
	4921:  uint16(anon_sym_orsource),
	4922:  uint16(98),
	4923:  uint16(12),
	4924:  uint16(sym__entry),
	4925:  uint16(sym_mainmenu),
	4926:  uint16(sym_config),
	4927:  uint16(sym_configdefault),
	4928:  uint16(sym_menuconfig),
	4929:  uint16(sym_choice),
	4930:  uint16(sym_comment_entry),
	4931:  uint16(sym_menu),
	4932:  uint16(sym_if),
	4933:  uint16(sym_source),
	4934:  uint16(sym_variable),
	4935:  uint16(aux_sym_configuration_repeat1),
	4936:  uint16(13),
	4937:  uint16(3),
	4938:  uint16(1),
	4939:  uint16(sym_comment),
	4940:  uint16(27),
	4941:  uint16(1),
	4942:  uint16(sym_symbol),
	4943:  uint16(29),
	4944:  uint16(1),
	4945:  uint16(anon_sym_mainmenu),
	4946:  uint16(31),
	4947:  uint16(1),
	4948:  uint16(anon_sym_config),
	4949:  uint16(33),
	4950:  uint16(1),
	4951:  uint16(anon_sym_configdefault),
	4952:  uint16(35),
	4953:  uint16(1),
	4954:  uint16(anon_sym_menuconfig),
	4955:  uint16(37),
	4956:  uint16(1),
	4957:  uint16(anon_sym_choice),
	4958:  uint16(39),
	4959:  uint16(1),
	4960:  uint16(anon_sym_comment),
	4961:  uint16(41),
	4962:  uint16(1),
	4963:  uint16(anon_sym_menu),
	4964:  uint16(45),
	4965:  uint16(1),
	4966:  uint16(anon_sym_if),
	4967:  uint16(77),
	4968:  uint16(1),
	4969:  uint16(anon_sym_endmenu),
	4970:  uint16(47),
	4971:  uint16(4),
	4972:  uint16(anon_sym_source),
	4973:  uint16(anon_sym_rsource),
	4974:  uint16(anon_sym_osource),
	4975:  uint16(anon_sym_orsource),
	4976:  uint16(98),
	4977:  uint16(12),
	4978:  uint16(sym__entry),
	4979:  uint16(sym_mainmenu),
	4980:  uint16(sym_config),
	4981:  uint16(sym_configdefault),
	4982:  uint16(sym_menuconfig),
	4983:  uint16(sym_choice),
	4984:  uint16(sym_comment_entry),
	4985:  uint16(sym_menu),
	4986:  uint16(sym_if),
	4987:  uint16(sym_source),
	4988:  uint16(sym_variable),
	4989:  uint16(aux_sym_configuration_repeat1),
	4990:  uint16(13),
	4991:  uint16(3),
	4992:  uint16(1),
	4993:  uint16(sym_comment),
	4994:  uint16(27),
	4995:  uint16(1),
	4996:  uint16(sym_symbol),
	4997:  uint16(29),
	4998:  uint16(1),
	4999:  uint16(anon_sym_mainmenu),
	5000:  uint16(31),
	5001:  uint16(1),
	5002:  uint16(anon_sym_config),
	5003:  uint16(33),
	5004:  uint16(1),
	5005:  uint16(anon_sym_configdefault),
	5006:  uint16(35),
	5007:  uint16(1),
	5008:  uint16(anon_sym_menuconfig),
	5009:  uint16(37),
	5010:  uint16(1),
	5011:  uint16(anon_sym_choice),
	5012:  uint16(39),
	5013:  uint16(1),
	5014:  uint16(anon_sym_comment),
	5015:  uint16(41),
	5016:  uint16(1),
	5017:  uint16(anon_sym_menu),
	5018:  uint16(45),
	5019:  uint16(1),
	5020:  uint16(anon_sym_if),
	5021:  uint16(499),
	5022:  uint16(1),
	5023:  uint16(anon_sym_endif),
	5024:  uint16(47),
	5025:  uint16(4),
	5026:  uint16(anon_sym_source),
	5027:  uint16(anon_sym_rsource),
	5028:  uint16(anon_sym_osource),
	5029:  uint16(anon_sym_orsource),
	5030:  uint16(98),
	5031:  uint16(12),
	5032:  uint16(sym__entry),
	5033:  uint16(sym_mainmenu),
	5034:  uint16(sym_config),
	5035:  uint16(sym_configdefault),
	5036:  uint16(sym_menuconfig),
	5037:  uint16(sym_choice),
	5038:  uint16(sym_comment_entry),
	5039:  uint16(sym_menu),
	5040:  uint16(sym_if),
	5041:  uint16(sym_source),
	5042:  uint16(sym_variable),
	5043:  uint16(aux_sym_configuration_repeat1),
	5044:  uint16(13),
	5045:  uint16(3),
	5046:  uint16(1),
	5047:  uint16(sym_comment),
	5048:  uint16(7),
	5049:  uint16(1),
	5050:  uint16(sym_symbol),
	5051:  uint16(9),
	5052:  uint16(1),
	5053:  uint16(anon_sym_mainmenu),
	5054:  uint16(11),
	5055:  uint16(1),
	5056:  uint16(anon_sym_config),
	5057:  uint16(13),
	5058:  uint16(1),
	5059:  uint16(anon_sym_configdefault),
	5060:  uint16(15),
	5061:  uint16(1),
	5062:  uint16(anon_sym_menuconfig),
	5063:  uint16(17),
	5064:  uint16(1),
	5065:  uint16(anon_sym_choice),
	5066:  uint16(19),
	5067:  uint16(1),
	5068:  uint16(anon_sym_comment),
	5069:  uint16(21),
	5070:  uint16(1),
	5071:  uint16(anon_sym_menu),
	5072:  uint16(23),
	5073:  uint16(1),
	5074:  uint16(anon_sym_if),
	5075:  uint16(501),
	5076:  uint16(1),
	5078:  uint16(25),
	5079:  uint16(4),
	5080:  uint16(anon_sym_source),
	5081:  uint16(anon_sym_rsource),
	5082:  uint16(anon_sym_osource),
	5083:  uint16(anon_sym_orsource),
	5084:  uint16(106),
	5085:  uint16(12),
	5086:  uint16(sym__entry),
	5087:  uint16(sym_mainmenu),
	5088:  uint16(sym_config),
	5089:  uint16(sym_configdefault),
	5090:  uint16(sym_menuconfig),
	5091:  uint16(sym_choice),
	5092:  uint16(sym_comment_entry),
	5093:  uint16(sym_menu),
	5094:  uint16(sym_if),
	5095:  uint16(sym_source),
	5096:  uint16(sym_variable),
	5097:  uint16(aux_sym_configuration_repeat1),
	5098:  uint16(13),
	5099:  uint16(3),
	5100:  uint16(1),
	5101:  uint16(sym_comment),
	5102:  uint16(27),
	5103:  uint16(1),
	5104:  uint16(sym_symbol),
	5105:  uint16(29),
	5106:  uint16(1),
	5107:  uint16(anon_sym_mainmenu),
	5108:  uint16(31),
	5109:  uint16(1),
	5110:  uint16(anon_sym_config),
	5111:  uint16(33),
	5112:  uint16(1),
	5113:  uint16(anon_sym_configdefault),
	5114:  uint16(35),
	5115:  uint16(1),
	5116:  uint16(anon_sym_menuconfig),
	5117:  uint16(37),
	5118:  uint16(1),
	5119:  uint16(anon_sym_choice),
	5120:  uint16(39),
	5121:  uint16(1),
	5122:  uint16(anon_sym_comment),
	5123:  uint16(41),
	5124:  uint16(1),
	5125:  uint16(anon_sym_menu),
	5126:  uint16(45),
	5127:  uint16(1),
	5128:  uint16(anon_sym_if),
	5129:  uint16(503),
	5130:  uint16(1),
	5131:  uint16(anon_sym_endmenu),
	5132:  uint16(47),
	5133:  uint16(4),
	5134:  uint16(anon_sym_source),
	5135:  uint16(anon_sym_rsource),
	5136:  uint16(anon_sym_osource),
	5137:  uint16(anon_sym_orsource),
	5138:  uint16(98),
	5139:  uint16(12),
	5140:  uint16(sym__entry),
	5141:  uint16(sym_mainmenu),
	5142:  uint16(sym_config),
	5143:  uint16(sym_configdefault),
	5144:  uint16(sym_menuconfig),
	5145:  uint16(sym_choice),
	5146:  uint16(sym_comment_entry),
	5147:  uint16(sym_menu),
	5148:  uint16(sym_if),
	5149:  uint16(sym_source),
	5150:  uint16(sym_variable),
	5151:  uint16(aux_sym_configuration_repeat1),
	5152:  uint16(13),
	5153:  uint16(3),
	5154:  uint16(1),
	5155:  uint16(sym_comment),
	5156:  uint16(505),
	5157:  uint16(1),
	5159:  uint16(507),
	5160:  uint16(1),
	5161:  uint16(sym_symbol),
	5162:  uint16(510),
	5163:  uint16(1),
	5164:  uint16(anon_sym_mainmenu),
	5165:  uint16(513),
	5166:  uint16(1),
	5167:  uint16(anon_sym_config),
	5168:  uint16(516),
	5169:  uint16(1),
	5170:  uint16(anon_sym_configdefault),
	5171:  uint16(519),
	5172:  uint16(1),
	5173:  uint16(anon_sym_menuconfig),
	5174:  uint16(522),
	5175:  uint16(1),
	5176:  uint16(anon_sym_choice),
	5177:  uint16(525),
	5178:  uint16(1),
	5179:  uint16(anon_sym_comment),
	5180:  uint16(528),
	5181:  uint16(1),
	5182:  uint16(anon_sym_menu),
	5183:  uint16(531),
	5184:  uint16(1),
	5185:  uint16(anon_sym_if),
	5186:  uint16(534),
	5187:  uint16(4),
	5188:  uint16(anon_sym_source),
	5189:  uint16(anon_sym_rsource),
	5190:  uint16(anon_sym_osource),
	5191:  uint16(anon_sym_orsource),
	5192:  uint16(106),
	5193:  uint16(12),
	5194:  uint16(sym__entry),
	5195:  uint16(sym_mainmenu),
	5196:  uint16(sym_config),
	5197:  uint16(sym_configdefault),
	5198:  uint16(sym_menuconfig),
	5199:  uint16(sym_choice),
	5200:  uint16(sym_comment_entry),
	5201:  uint16(sym_menu),
	5202:  uint16(sym_if),
	5203:  uint16(sym_source),
	5204:  uint16(sym_variable),
	5205:  uint16(aux_sym_configuration_repeat1),
	5206:  uint16(13),
	5207:  uint16(3),
	5208:  uint16(1),
	5209:  uint16(sym_comment),
	5210:  uint16(27),
	5211:  uint16(1),
	5212:  uint16(sym_symbol),
	5213:  uint16(29),
	5214:  uint16(1),
	5215:  uint16(anon_sym_mainmenu),
	5216:  uint16(31),
	5217:  uint16(1),
	5218:  uint16(anon_sym_config),
	5219:  uint16(33),
	5220:  uint16(1),
	5221:  uint16(anon_sym_configdefault),
	5222:  uint16(35),
	5223:  uint16(1),
	5224:  uint16(anon_sym_menuconfig),
	5225:  uint16(37),
	5226:  uint16(1),
	5227:  uint16(anon_sym_choice),
	5228:  uint16(39),
	5229:  uint16(1),
	5230:  uint16(anon_sym_comment),
	5231:  uint16(41),
	5232:  uint16(1),
	5233:  uint16(anon_sym_menu),
	5234:  uint16(45),
	5235:  uint16(1),
	5236:  uint16(anon_sym_if),
	5237:  uint16(537),
	5238:  uint16(1),
	5239:  uint16(anon_sym_endchoice),
	5240:  uint16(47),
	5241:  uint16(4),
	5242:  uint16(anon_sym_source),
	5243:  uint16(anon_sym_rsource),
	5244:  uint16(anon_sym_osource),
	5245:  uint16(anon_sym_orsource),
	5246:  uint16(98),
	5247:  uint16(12),
	5248:  uint16(sym__entry),
	5249:  uint16(sym_mainmenu),
	5250:  uint16(sym_config),
	5251:  uint16(sym_configdefault),
	5252:  uint16(sym_menuconfig),
	5253:  uint16(sym_choice),
	5254:  uint16(sym_comment_entry),
	5255:  uint16(sym_menu),
	5256:  uint16(sym_if),
	5257:  uint16(sym_source),
	5258:  uint16(sym_variable),
	5259:  uint16(aux_sym_configuration_repeat1),
	5260:  uint16(13),
	5261:  uint16(3),
	5262:  uint16(1),
	5263:  uint16(sym_comment),
	5264:  uint16(27),
	5265:  uint16(1),
	5266:  uint16(sym_symbol),
	5267:  uint16(29),
	5268:  uint16(1),
	5269:  uint16(anon_sym_mainmenu),
	5270:  uint16(31),
	5271:  uint16(1),
	5272:  uint16(anon_sym_config),
	5273:  uint16(33),
	5274:  uint16(1),
	5275:  uint16(anon_sym_configdefault),
	5276:  uint16(35),
	5277:  uint16(1),
	5278:  uint16(anon_sym_menuconfig),
	5279:  uint16(37),
	5280:  uint16(1),
	5281:  uint16(anon_sym_choice),
	5282:  uint16(39),
	5283:  uint16(1),
	5284:  uint16(anon_sym_comment),
	5285:  uint16(41),
	5286:  uint16(1),
	5287:  uint16(anon_sym_menu),
	5288:  uint16(45),
	5289:  uint16(1),
	5290:  uint16(anon_sym_if),
	5291:  uint16(87),
	5292:  uint16(1),
	5293:  uint16(anon_sym_endmenu),
	5294:  uint16(47),
	5295:  uint16(4),
	5296:  uint16(anon_sym_source),
	5297:  uint16(anon_sym_rsource),
	5298:  uint16(anon_sym_osource),
	5299:  uint16(anon_sym_orsource),
	5300:  uint16(98),
	5301:  uint16(12),
	5302:  uint16(sym__entry),
	5303:  uint16(sym_mainmenu),
	5304:  uint16(sym_config),
	5305:  uint16(sym_configdefault),
	5306:  uint16(sym_menuconfig),
	5307:  uint16(sym_choice),
	5308:  uint16(sym_comment_entry),
	5309:  uint16(sym_menu),
	5310:  uint16(sym_if),
	5311:  uint16(sym_source),
	5312:  uint16(sym_variable),
	5313:  uint16(aux_sym_configuration_repeat1),
	5314:  uint16(13),
	5315:  uint16(3),
	5316:  uint16(1),
	5317:  uint16(sym_comment),
	5318:  uint16(27),
	5319:  uint16(1),
	5320:  uint16(sym_symbol),
	5321:  uint16(29),
	5322:  uint16(1),
	5323:  uint16(anon_sym_mainmenu),
	5324:  uint16(31),
	5325:  uint16(1),
	5326:  uint16(anon_sym_config),
	5327:  uint16(33),
	5328:  uint16(1),
	5329:  uint16(anon_sym_configdefault),
	5330:  uint16(35),
	5331:  uint16(1),
	5332:  uint16(anon_sym_menuconfig),
	5333:  uint16(37),
	5334:  uint16(1),
	5335:  uint16(anon_sym_choice),
	5336:  uint16(39),
	5337:  uint16(1),
	5338:  uint16(anon_sym_comment),
	5339:  uint16(41),
	5340:  uint16(1),
	5341:  uint16(anon_sym_menu),
	5342:  uint16(45),
	5343:  uint16(1),
	5344:  uint16(anon_sym_if),
	5345:  uint16(539),
	5346:  uint16(1),
	5347:  uint16(anon_sym_endchoice),
	5348:  uint16(47),
	5349:  uint16(4),
	5350:  uint16(anon_sym_source),
	5351:  uint16(anon_sym_rsource),
	5352:  uint16(anon_sym_osource),
	5353:  uint16(anon_sym_orsource),
	5354:  uint16(98),
	5355:  uint16(12),
	5356:  uint16(sym__entry),
	5357:  uint16(sym_mainmenu),
	5358:  uint16(sym_config),
	5359:  uint16(sym_configdefault),
	5360:  uint16(sym_menuconfig),
	5361:  uint16(sym_choice),
	5362:  uint16(sym_comment_entry),
	5363:  uint16(sym_menu),
	5364:  uint16(sym_if),
	5365:  uint16(sym_source),
	5366:  uint16(sym_variable),
	5367:  uint16(aux_sym_configuration_repeat1),
	5368:  uint16(13),
	5369:  uint16(3),
	5370:  uint16(1),
	5371:  uint16(sym_comment),
	5372:  uint16(27),
	5373:  uint16(1),
	5374:  uint16(sym_symbol),
	5375:  uint16(29),
	5376:  uint16(1),
	5377:  uint16(anon_sym_mainmenu),
	5378:  uint16(31),
	5379:  uint16(1),
	5380:  uint16(anon_sym_config),
	5381:  uint16(33),
	5382:  uint16(1),
	5383:  uint16(anon_sym_configdefault),
	5384:  uint16(35),
	5385:  uint16(1),
	5386:  uint16(anon_sym_menuconfig),
	5387:  uint16(37),
	5388:  uint16(1),
	5389:  uint16(anon_sym_choice),
	5390:  uint16(39),
	5391:  uint16(1),
	5392:  uint16(anon_sym_comment),
	5393:  uint16(41),
	5394:  uint16(1),
	5395:  uint16(anon_sym_menu),
	5396:  uint16(45),
	5397:  uint16(1),
	5398:  uint16(anon_sym_if),
	5399:  uint16(541),
	5400:  uint16(1),
	5401:  uint16(anon_sym_endif),
	5402:  uint16(47),
	5403:  uint16(4),
	5404:  uint16(anon_sym_source),
	5405:  uint16(anon_sym_rsource),
	5406:  uint16(anon_sym_osource),
	5407:  uint16(anon_sym_orsource),
	5408:  uint16(98),
	5409:  uint16(12),
	5410:  uint16(sym__entry),
	5411:  uint16(sym_mainmenu),
	5412:  uint16(sym_config),
	5413:  uint16(sym_configdefault),
	5414:  uint16(sym_menuconfig),
	5415:  uint16(sym_choice),
	5416:  uint16(sym_comment_entry),
	5417:  uint16(sym_menu),
	5418:  uint16(sym_if),
	5419:  uint16(sym_source),
	5420:  uint16(sym_variable),
	5421:  uint16(aux_sym_configuration_repeat1),
	5422:  uint16(13),
	5423:  uint16(3),
	5424:  uint16(1),
	5425:  uint16(sym_comment),
	5426:  uint16(27),
	5427:  uint16(1),
	5428:  uint16(sym_symbol),
	5429:  uint16(29),
	5430:  uint16(1),
	5431:  uint16(anon_sym_mainmenu),
	5432:  uint16(31),
	5433:  uint16(1),
	5434:  uint16(anon_sym_config),
	5435:  uint16(33),
	5436:  uint16(1),
	5437:  uint16(anon_sym_configdefault),
	5438:  uint16(35),
	5439:  uint16(1),
	5440:  uint16(anon_sym_menuconfig),
	5441:  uint16(37),
	5442:  uint16(1),
	5443:  uint16(anon_sym_choice),
	5444:  uint16(39),
	5445:  uint16(1),
	5446:  uint16(anon_sym_comment),
	5447:  uint16(41),
	5448:  uint16(1),
	5449:  uint16(anon_sym_menu),
	5450:  uint16(45),
	5451:  uint16(1),
	5452:  uint16(anon_sym_if),
	5453:  uint16(543),
	5454:  uint16(1),
	5455:  uint16(anon_sym_endchoice),
	5456:  uint16(47),
	5457:  uint16(4),
	5458:  uint16(anon_sym_source),
	5459:  uint16(anon_sym_rsource),
	5460:  uint16(anon_sym_osource),
	5461:  uint16(anon_sym_orsource),
	5462:  uint16(98),
	5463:  uint16(12),
	5464:  uint16(sym__entry),
	5465:  uint16(sym_mainmenu),
	5466:  uint16(sym_config),
	5467:  uint16(sym_configdefault),
	5468:  uint16(sym_menuconfig),
	5469:  uint16(sym_choice),
	5470:  uint16(sym_comment_entry),
	5471:  uint16(sym_menu),
	5472:  uint16(sym_if),
	5473:  uint16(sym_source),
	5474:  uint16(sym_variable),
	5475:  uint16(aux_sym_configuration_repeat1),
	5476:  uint16(13),
	5477:  uint16(3),
	5478:  uint16(1),
	5479:  uint16(sym_comment),
	5480:  uint16(27),
	5481:  uint16(1),
	5482:  uint16(sym_symbol),
	5483:  uint16(29),
	5484:  uint16(1),
	5485:  uint16(anon_sym_mainmenu),
	5486:  uint16(31),
	5487:  uint16(1),
	5488:  uint16(anon_sym_config),
	5489:  uint16(33),
	5490:  uint16(1),
	5491:  uint16(anon_sym_configdefault),
	5492:  uint16(35),
	5493:  uint16(1),
	5494:  uint16(anon_sym_menuconfig),
	5495:  uint16(37),
	5496:  uint16(1),
	5497:  uint16(anon_sym_choice),
	5498:  uint16(39),
	5499:  uint16(1),
	5500:  uint16(anon_sym_comment),
	5501:  uint16(41),
	5502:  uint16(1),
	5503:  uint16(anon_sym_menu),
	5504:  uint16(45),
	5505:  uint16(1),
	5506:  uint16(anon_sym_if),
	5507:  uint16(545),
	5508:  uint16(1),
	5509:  uint16(anon_sym_endchoice),
	5510:  uint16(47),
	5511:  uint16(4),
	5512:  uint16(anon_sym_source),
	5513:  uint16(anon_sym_rsource),
	5514:  uint16(anon_sym_osource),
	5515:  uint16(anon_sym_orsource),
	5516:  uint16(98),
	5517:  uint16(12),
	5518:  uint16(sym__entry),
	5519:  uint16(sym_mainmenu),
	5520:  uint16(sym_config),
	5521:  uint16(sym_configdefault),
	5522:  uint16(sym_menuconfig),
	5523:  uint16(sym_choice),
	5524:  uint16(sym_comment_entry),
	5525:  uint16(sym_menu),
	5526:  uint16(sym_if),
	5527:  uint16(sym_source),
	5528:  uint16(sym_variable),
	5529:  uint16(aux_sym_configuration_repeat1),
	5530:  uint16(4),
	5531:  uint16(3),
	5532:  uint16(1),
	5533:  uint16(sym_comment),
	5534:  uint16(196),
	5535:  uint16(3),
	5536:  uint16(anon_sym_DOLLAR_LPAREN),
	5537:  uint16(anon_sym_DQUOTE),
	5538:  uint16(anon_sym_SQUOTE),
	5539:  uint16(193),
	5540:  uint16(6),
	5541:  uint16(anon_sym_EQ),
	5542:  uint16(anon_sym_PIPE_PIPE),
	5543:  uint16(anon_sym_AMP_AMP),
	5544:  uint16(anon_sym_BANG_EQ),
	5545:  uint16(anon_sym_LT_EQ),
	5546:  uint16(anon_sym_GT_EQ),
	5547:  uint16(190),
	5548:  uint16(16),
	5549:  uint16(anon_sym_mainmenu),
	5550:  uint16(anon_sym_config),
	5551:  uint16(anon_sym_configdefault),
	5552:  uint16(anon_sym_menuconfig),
	5553:  uint16(anon_sym_choice),
	5554:  uint16(anon_sym_comment),
	5555:  uint16(anon_sym_menu),
	5556:  uint16(anon_sym_if),
	5557:  uint16(anon_sym_endif),
	5558:  uint16(anon_sym_source),
	5559:  uint16(anon_sym_rsource),
	5560:  uint16(anon_sym_osource),
	5561:  uint16(anon_sym_orsource),
	5562:  uint16(anon_sym_LT),
	5563:  uint16(anon_sym_GT),
	5564:  uint16(sym_symbol),
	5565:  uint16(3),
	5566:  uint16(3),
	5567:  uint16(1),
	5568:  uint16(sym_comment),
	5569:  uint16(279),
	5570:  uint16(7),
	5571:  uint16(anon_sym_EQ),
	5572:  uint16(anon_sym_PIPE_PIPE),
	5573:  uint16(anon_sym_AMP_AMP),
	5574:  uint16(anon_sym_BANG_EQ),
	5575:  uint16(anon_sym_LT_EQ),
	5576:  uint16(anon_sym_GT_EQ),
	5577:  uint16(anon_sym_RPAREN),
	5578:  uint16(277),
	5579:  uint16(16),
	5580:  uint16(anon_sym_mainmenu),
	5581:  uint16(anon_sym_config),
	5582:  uint16(anon_sym_configdefault),
	5583:  uint16(anon_sym_menuconfig),
	5584:  uint16(anon_sym_choice),
	5585:  uint16(anon_sym_comment),
	5586:  uint16(anon_sym_menu),
	5587:  uint16(anon_sym_if),
	5588:  uint16(anon_sym_endif),
	5589:  uint16(anon_sym_source),
	5590:  uint16(anon_sym_rsource),
	5591:  uint16(anon_sym_osource),
	5592:  uint16(anon_sym_orsource),
	5593:  uint16(anon_sym_LT),
	5594:  uint16(anon_sym_GT),
	5595:  uint16(sym_symbol),
	5596:  uint16(3),
	5597:  uint16(3),
	5598:  uint16(1),
	5599:  uint16(sym_comment),
	5600:  uint16(275),
	5601:  uint16(7),
	5602:  uint16(anon_sym_EQ),
	5603:  uint16(anon_sym_PIPE_PIPE),
	5604:  uint16(anon_sym_AMP_AMP),
	5605:  uint16(anon_sym_BANG_EQ),
	5606:  uint16(anon_sym_LT_EQ),
	5607:  uint16(anon_sym_GT_EQ),
	5608:  uint16(anon_sym_RPAREN),
	5609:  uint16(273),
	5610:  uint16(16),
	5611:  uint16(anon_sym_mainmenu),
	5612:  uint16(anon_sym_config),
	5613:  uint16(anon_sym_configdefault),
	5614:  uint16(anon_sym_menuconfig),
	5615:  uint16(anon_sym_choice),
	5616:  uint16(anon_sym_comment),
	5617:  uint16(anon_sym_menu),
	5618:  uint16(anon_sym_if),
	5619:  uint16(anon_sym_endif),
	5620:  uint16(anon_sym_source),
	5621:  uint16(anon_sym_rsource),
	5622:  uint16(anon_sym_osource),
	5623:  uint16(anon_sym_orsource),
	5624:  uint16(anon_sym_LT),
	5625:  uint16(anon_sym_GT),
	5626:  uint16(sym_symbol),
	5627:  uint16(3),
	5628:  uint16(3),
	5629:  uint16(1),
	5630:  uint16(sym_comment),
	5631:  uint16(293),
	5632:  uint16(7),
	5633:  uint16(anon_sym_EQ),
	5634:  uint16(anon_sym_PIPE_PIPE),
	5635:  uint16(anon_sym_AMP_AMP),
	5636:  uint16(anon_sym_BANG_EQ),
	5637:  uint16(anon_sym_LT_EQ),
	5638:  uint16(anon_sym_GT_EQ),
	5639:  uint16(anon_sym_RPAREN),
	5640:  uint16(291),
	5641:  uint16(16),
	5642:  uint16(anon_sym_mainmenu),
	5643:  uint16(anon_sym_config),
	5644:  uint16(anon_sym_configdefault),
	5645:  uint16(anon_sym_menuconfig),
	5646:  uint16(anon_sym_choice),
	5647:  uint16(anon_sym_comment),
	5648:  uint16(anon_sym_menu),
	5649:  uint16(anon_sym_if),
	5650:  uint16(anon_sym_endif),
	5651:  uint16(anon_sym_source),
	5652:  uint16(anon_sym_rsource),
	5653:  uint16(anon_sym_osource),
	5654:  uint16(anon_sym_orsource),
	5655:  uint16(anon_sym_LT),
	5656:  uint16(anon_sym_GT),
	5657:  uint16(sym_symbol),
	5658:  uint16(3),
	5659:  uint16(3),
	5660:  uint16(1),
	5661:  uint16(sym_comment),
	5662:  uint16(289),
	5663:  uint16(7),
	5664:  uint16(anon_sym_EQ),
	5665:  uint16(anon_sym_PIPE_PIPE),
	5666:  uint16(anon_sym_AMP_AMP),
	5667:  uint16(anon_sym_BANG_EQ),
	5668:  uint16(anon_sym_LT_EQ),
	5669:  uint16(anon_sym_GT_EQ),
	5670:  uint16(anon_sym_RPAREN),
	5671:  uint16(287),
	5672:  uint16(16),
	5673:  uint16(anon_sym_mainmenu),
	5674:  uint16(anon_sym_config),
	5675:  uint16(anon_sym_configdefault),
	5676:  uint16(anon_sym_menuconfig),
	5677:  uint16(anon_sym_choice),
	5678:  uint16(anon_sym_comment),
	5679:  uint16(anon_sym_menu),
	5680:  uint16(anon_sym_if),
	5681:  uint16(anon_sym_endif),
	5682:  uint16(anon_sym_source),
	5683:  uint16(anon_sym_rsource),
	5684:  uint16(anon_sym_osource),
	5685:  uint16(anon_sym_orsource),
	5686:  uint16(anon_sym_LT),
	5687:  uint16(anon_sym_GT),
	5688:  uint16(sym_symbol),
	5689:  uint16(7),
	5690:  uint16(3),
	5691:  uint16(1),
	5692:  uint16(sym_comment),
	5693:  uint16(279),
	5694:  uint16(1),
	5695:  uint16(anon_sym_PIPE_PIPE),
	5696:  uint16(389),
	5697:  uint16(1),
	5698:  uint16(anon_sym_EQ),
	5699:  uint16(393),
	5700:  uint16(1),
	5701:  uint16(anon_sym_AMP_AMP),
	5702:  uint16(397),
	5703:  uint16(2),
	5704:  uint16(anon_sym_LT),
	5705:  uint16(anon_sym_GT),
	5706:  uint16(395),
	5707:  uint16(3),
	5708:  uint16(anon_sym_BANG_EQ),
	5709:  uint16(anon_sym_LT_EQ),
	5710:  uint16(anon_sym_GT_EQ),
	5711:  uint16(277),
	5712:  uint16(14),
	5713:  uint16(anon_sym_mainmenu),
	5714:  uint16(anon_sym_config),
	5715:  uint16(anon_sym_configdefault),
	5716:  uint16(anon_sym_menuconfig),
	5717:  uint16(anon_sym_choice),
	5718:  uint16(anon_sym_comment),
	5719:  uint16(anon_sym_menu),
	5720:  uint16(anon_sym_if),
	5721:  uint16(anon_sym_endif),
	5722:  uint16(anon_sym_source),
	5723:  uint16(anon_sym_rsource),
	5724:  uint16(anon_sym_osource),
	5725:  uint16(anon_sym_orsource),
	5726:  uint16(sym_symbol),
	5727:  uint16(6),
	5728:  uint16(3),
	5729:  uint16(1),
	5730:  uint16(sym_comment),
	5731:  uint16(389),
	5732:  uint16(1),
	5733:  uint16(anon_sym_EQ),
	5734:  uint16(279),
	5735:  uint16(2),
	5736:  uint16(anon_sym_PIPE_PIPE),
	5737:  uint16(anon_sym_AMP_AMP),
	5738:  uint16(397),
	5739:  uint16(2),
	5740:  uint16(anon_sym_LT),
	5741:  uint16(anon_sym_GT),
	5742:  uint16(395),
	5743:  uint16(3),
	5744:  uint16(anon_sym_BANG_EQ),
	5745:  uint16(anon_sym_LT_EQ),
	5746:  uint16(anon_sym_GT_EQ),
	5747:  uint16(277),
	5748:  uint16(14),
	5749:  uint16(anon_sym_mainmenu),
	5750:  uint16(anon_sym_config),
	5751:  uint16(anon_sym_configdefault),
	5752:  uint16(anon_sym_menuconfig),
	5753:  uint16(anon_sym_choice),
	5754:  uint16(anon_sym_comment),
	5755:  uint16(anon_sym_menu),
	5756:  uint16(anon_sym_if),
	5757:  uint16(anon_sym_endif),
	5758:  uint16(anon_sym_source),
	5759:  uint16(anon_sym_rsource),
	5760:  uint16(anon_sym_osource),
	5761:  uint16(anon_sym_orsource),
	5762:  uint16(sym_symbol),
	5763:  uint16(4),
	5764:  uint16(3),
	5765:  uint16(1),
	5766:  uint16(sym_comment),
	5767:  uint16(389),
	5768:  uint16(1),
	5769:  uint16(anon_sym_EQ),
	5770:  uint16(279),
	5771:  uint16(5),
	5772:  uint16(anon_sym_PIPE_PIPE),
	5773:  uint16(anon_sym_AMP_AMP),
	5774:  uint16(anon_sym_BANG_EQ),
	5775:  uint16(anon_sym_LT_EQ),
	5776:  uint16(anon_sym_GT_EQ),
	5777:  uint16(277),
	5778:  uint16(16),
	5779:  uint16(anon_sym_mainmenu),
	5780:  uint16(anon_sym_config),
	5781:  uint16(anon_sym_configdefault),
	5782:  uint16(anon_sym_menuconfig),
	5783:  uint16(anon_sym_choice),
	5784:  uint16(anon_sym_comment),
	5785:  uint16(anon_sym_menu),
	5786:  uint16(anon_sym_if),
	5787:  uint16(anon_sym_endif),
	5788:  uint16(anon_sym_source),
	5789:  uint16(anon_sym_rsource),
	5790:  uint16(anon_sym_osource),
	5791:  uint16(anon_sym_orsource),
	5792:  uint16(anon_sym_LT),
	5793:  uint16(anon_sym_GT),
	5794:  uint16(sym_symbol),
	5795:  uint16(4),
	5796:  uint16(3),
	5797:  uint16(1),
	5798:  uint16(sym_comment),
	5799:  uint16(549),
	5800:  uint16(1),
	5801:  uint16(anon_sym_default),
	5802:  uint16(122),
	5803:  uint16(2),
	5804:  uint16(sym_default_value),
	5805:  uint16(aux_sym_configdefault_repeat1),
	5806:  uint16(547),
	5807:  uint16(16),
	5808:  uint16(anon_sym_mainmenu),
	5809:  uint16(anon_sym_config),
	5810:  uint16(anon_sym_configdefault),
	5811:  uint16(anon_sym_menuconfig),
	5812:  uint16(anon_sym_choice),
	5813:  uint16(anon_sym_endchoice),
	5814:  uint16(anon_sym_comment),
	5815:  uint16(anon_sym_menu),
	5816:  uint16(anon_sym_endmenu),
	5817:  uint16(anon_sym_if),
	5818:  uint16(anon_sym_endif),
	5819:  uint16(anon_sym_source),
	5820:  uint16(anon_sym_rsource),
	5821:  uint16(anon_sym_osource),
	5822:  uint16(anon_sym_orsource),
	5823:  uint16(sym_symbol),
	5824:  uint16(4),
	5825:  uint16(3),
	5826:  uint16(1),
	5827:  uint16(sym_comment),
	5828:  uint16(553),
	5829:  uint16(1),
	5830:  uint16(anon_sym_default),
	5831:  uint16(122),
	5832:  uint16(2),
	5833:  uint16(sym_default_value),
	5834:  uint16(aux_sym_configdefault_repeat1),
	5835:  uint16(551),
	5836:  uint16(16),
	5837:  uint16(anon_sym_mainmenu),
	5838:  uint16(anon_sym_config),
	5839:  uint16(anon_sym_configdefault),
	5840:  uint16(anon_sym_menuconfig),
	5841:  uint16(anon_sym_choice),
	5842:  uint16(anon_sym_endchoice),
	5843:  uint16(anon_sym_comment),
	5844:  uint16(anon_sym_menu),
	5845:  uint16(anon_sym_endmenu),
	5846:  uint16(anon_sym_if),
	5847:  uint16(anon_sym_endif),
	5848:  uint16(anon_sym_source),
	5849:  uint16(anon_sym_rsource),
	5850:  uint16(anon_sym_osource),
	5851:  uint16(anon_sym_orsource),
	5852:  uint16(sym_symbol),
	5853:  uint16(8),
	5854:  uint16(106),
	5855:  uint16(1),
	5856:  uint16(aux_sym_variable_token1),
	5857:  uint16(556),
	5858:  uint16(1),
	5859:  uint16(sym_symbol),
	5860:  uint16(559),
	5861:  uint16(1),
	5862:  uint16(anon_sym_DOLLAR_LPAREN),
	5863:  uint16(562),
	5864:  uint16(1),
	5865:  uint16(anon_sym_DQUOTE),
	5866:  uint16(565),
	5867:  uint16(1),
	5868:  uint16(anon_sym_SQUOTE),
	5869:  uint16(568),
	5870:  uint16(1),
	5871:  uint16(sym_comment),
	5872:  uint16(123),
	5873:  uint16(3),
	5874:  uint16(sym_macro_variable),
	5875:  uint16(sym_string),
	5876:  uint16(aux_sym_name_repeat1),
	5877:  uint16(104),
	5878:  uint16(11),
	5879:  uint16(anon_sym_EQ),
	5880:  uint16(anon_sym_COMMA),
	5881:  uint16(anon_sym_BANG),
	5882:  uint16(anon_sym_PIPE_PIPE),
	5883:  uint16(anon_sym_AMP_AMP),
	5884:  uint16(anon_sym_BANG_EQ),
	5885:  uint16(anon_sym_LT),
	5886:  uint16(anon_sym_GT),
	5887:  uint16(anon_sym_LT_EQ),
	5888:  uint16(anon_sym_GT_EQ),
	5889:  uint16(anon_sym_LPAREN),
	5890:  uint16(8),
	5891:  uint16(93),
	5892:  uint16(1),
	5893:  uint16(aux_sym_variable_token1),
	5894:  uint16(568),
	5895:  uint16(1),
	5896:  uint16(sym_comment),
	5897:  uint16(570),
	5898:  uint16(1),
	5899:  uint16(sym_symbol),
	5900:  uint16(572),
	5901:  uint16(1),
	5902:  uint16(anon_sym_DOLLAR_LPAREN),
	5903:  uint16(574),
	5904:  uint16(1),
	5905:  uint16(anon_sym_DQUOTE),
	5906:  uint16(576),
	5907:  uint16(1),
	5908:  uint16(anon_sym_SQUOTE),
	5909:  uint16(123),
	5910:  uint16(3),
	5911:  uint16(sym_macro_variable),
	5912:  uint16(sym_string),
	5913:  uint16(aux_sym_name_repeat1),
	5914:  uint16(91),
	5915:  uint16(11),
	5916:  uint16(anon_sym_EQ),
	5917:  uint16(anon_sym_COMMA),
	5918:  uint16(anon_sym_BANG),
	5919:  uint16(anon_sym_PIPE_PIPE),
	5920:  uint16(anon_sym_AMP_AMP),
	5921:  uint16(anon_sym_BANG_EQ),
	5922:  uint16(anon_sym_LT),
	5923:  uint16(anon_sym_GT),
	5924:  uint16(anon_sym_LT_EQ),
	5925:  uint16(anon_sym_GT_EQ),
	5926:  uint16(anon_sym_LPAREN),
	5927:  uint16(15),
	5928:  uint16(568),
	5929:  uint16(1),
	5930:  uint16(sym_comment),
	5931:  uint16(572),
	5932:  uint16(1),
	5933:  uint16(anon_sym_DOLLAR_LPAREN),
	5934:  uint16(574),
	5935:  uint16(1),
	5936:  uint16(anon_sym_DQUOTE),
	5937:  uint16(576),
	5938:  uint16(1),
	5939:  uint16(anon_sym_SQUOTE),
	5940:  uint16(578),
	5941:  uint16(1),
	5942:  uint16(sym_symbol),
	5943:  uint16(580),
	5944:  uint16(1),
	5945:  uint16(anon_sym_COMMA),
	5946:  uint16(582),
	5947:  uint16(1),
	5948:  uint16(aux_sym_variable_token1),
	5949:  uint16(584),
	5950:  uint16(1),
	5951:  uint16(anon_sym_BANG),
	5952:  uint16(586),
	5953:  uint16(1),
	5954:  uint16(anon_sym_LPAREN),
	5955:  uint16(588),
	5956:  uint16(1),
	5957:  uint16(sym_text),
	5958:  uint16(124),
	5959:  uint16(1),
	5960:  uint16(aux_sym_name_repeat1),
	5961:  uint16(134),
	5962:  uint16(1),
	5963:  uint16(aux_sym_variable_repeat1),
	5964:  uint16(149),
	5965:  uint16(1),
	5966:  uint16(sym_expression),
	5967:  uint16(157),
	5968:  uint16(2),
	5969:  uint16(sym_macro_variable),
	5970:  uint16(sym_string),
	5971:  uint16(153),
	5972:  uint16(4),
	5973:  uint16(sym_unary_expression),
	5974:  uint16(sym_binary_expression),
	5975:  uint16(sym_parenthesized_expression),
	5976:  uint16(sym_name),
	5977:  uint16(15),
	5978:  uint16(568),
	5979:  uint16(1),
	5980:  uint16(sym_comment),
	5981:  uint16(572),
	5982:  uint16(1),
	5983:  uint16(anon_sym_DOLLAR_LPAREN),
	5984:  uint16(574),
	5985:  uint16(1),
	5986:  uint16(anon_sym_DQUOTE),
	5987:  uint16(576),
	5988:  uint16(1),
	5989:  uint16(anon_sym_SQUOTE),
	5990:  uint16(578),
	5991:  uint16(1),
	5992:  uint16(sym_symbol),
	5993:  uint16(584),
	5994:  uint16(1),
	5995:  uint16(anon_sym_BANG),
	5996:  uint16(586),
	5997:  uint16(1),
	5998:  uint16(anon_sym_LPAREN),
	5999:  uint16(590),
	6000:  uint16(1),
	6001:  uint16(anon_sym_COMMA),
	6002:  uint16(592),
	6003:  uint16(1),
	6004:  uint16(aux_sym_variable_token1),
	6005:  uint16(594),
	6006:  uint16(1),
	6007:  uint16(sym_text),
	6008:  uint16(124),
	6009:  uint16(1),
	6010:  uint16(aux_sym_name_repeat1),
	6011:  uint16(132),
	6012:  uint16(1),
	6013:  uint16(aux_sym_variable_repeat1),
	6014:  uint16(149),
	6015:  uint16(1),
	6016:  uint16(sym_expression),
	6017:  uint16(157),
	6018:  uint16(2),
	6019:  uint16(sym_macro_variable),
	6020:  uint16(sym_string),
	6021:  uint16(153),
	6022:  uint16(4),
	6023:  uint16(sym_unary_expression),
	6024:  uint16(sym_binary_expression),
	6025:  uint16(sym_parenthesized_expression),
	6026:  uint16(sym_name),
	6027:  uint16(5),
	6028:  uint16(3),
	6029:  uint16(1),
	6030:  uint16(sym_comment),
	6031:  uint16(596),
	6032:  uint16(1),
	6034:  uint16(598),
	6035:  uint16(1),
	6036:  uint16(anon_sym_default),
	6037:  uint16(127),
	6038:  uint16(2),
	6039:  uint16(sym_default_value),
	6040:  uint16(aux_sym_configdefault_repeat1),
	6041:  uint16(551),
	6042:  uint16(13),
	6043:  uint16(anon_sym_mainmenu),
	6044:  uint16(anon_sym_config),
	6045:  uint16(anon_sym_configdefault),
	6046:  uint16(anon_sym_menuconfig),
	6047:  uint16(anon_sym_choice),
	6048:  uint16(anon_sym_comment),
	6049:  uint16(anon_sym_menu),
	6050:  uint16(anon_sym_if),
	6051:  uint16(anon_sym_source),
	6052:  uint16(anon_sym_rsource),
	6053:  uint16(anon_sym_osource),
	6054:  uint16(anon_sym_orsource),
	6055:  uint16(sym_symbol),
	6056:  uint16(5),
	6057:  uint16(3),
	6058:  uint16(1),
	6059:  uint16(sym_comment),
	6060:  uint16(601),
	6061:  uint16(1),
	6063:  uint16(603),
	6064:  uint16(1),
	6065:  uint16(anon_sym_default),
	6066:  uint16(127),
	6067:  uint16(2),
	6068:  uint16(sym_default_value),
	6069:  uint16(aux_sym_configdefault_repeat1),
	6070:  uint16(547),
	6071:  uint16(13),
	6072:  uint16(anon_sym_mainmenu),
	6073:  uint16(anon_sym_config),
	6074:  uint16(anon_sym_configdefault),
	6075:  uint16(anon_sym_menuconfig),
	6076:  uint16(anon_sym_choice),
	6077:  uint16(anon_sym_comment),
	6078:  uint16(anon_sym_menu),
	6079:  uint16(anon_sym_if),
	6080:  uint16(anon_sym_source),
	6081:  uint16(anon_sym_rsource),
	6082:  uint16(anon_sym_osource),
	6083:  uint16(anon_sym_orsource),
	6084:  uint16(sym_symbol),
	6085:  uint16(2),
	6086:  uint16(3),
	6087:  uint16(1),
	6088:  uint16(sym_comment),
	6089:  uint16(379),
	6090:  uint16(17),
	6091:  uint16(anon_sym_mainmenu),
	6092:  uint16(anon_sym_config),
	6093:  uint16(anon_sym_configdefault),
	6094:  uint16(anon_sym_menuconfig),
	6095:  uint16(anon_sym_choice),
	6096:  uint16(anon_sym_endchoice),
	6097:  uint16(anon_sym_comment),
	6098:  uint16(anon_sym_menu),
	6099:  uint16(anon_sym_endmenu),
	6100:  uint16(anon_sym_if),
	6101:  uint16(anon_sym_endif),
	6102:  uint16(anon_sym_source),
	6103:  uint16(anon_sym_rsource),
	6104:  uint16(anon_sym_osource),
	6105:  uint16(anon_sym_orsource),
	6106:  uint16(anon_sym_default),
	6107:  uint16(sym_symbol),
	6108:  uint16(8),
	6109:  uint16(106),
	6110:  uint16(1),
	6111:  uint16(aux_sym_type_definition_token1),
	6112:  uint16(568),
	6113:  uint16(1),
	6114:  uint16(sym_comment),
	6115:  uint16(605),
	6116:  uint16(1),
	6117:  uint16(sym_symbol),
	6118:  uint16(608),
	6119:  uint16(1),
	6120:  uint16(anon_sym_DOLLAR_LPAREN),
	6121:  uint16(611),
	6122:  uint16(1),
	6123:  uint16(anon_sym_DQUOTE),
	6124:  uint16(614),
	6125:  uint16(1),
	6126:  uint16(anon_sym_SQUOTE),
	6127:  uint16(130),
	6128:  uint16(3),
	6129:  uint16(sym_macro_variable),
	6130:  uint16(sym_string),
	6131:  uint16(aux_sym_name_repeat1),
	6132:  uint16(104),
	6133:  uint16(9),
	6134:  uint16(anon_sym_if),
	6135:  uint16(anon_sym_EQ),
	6136:  uint16(anon_sym_PIPE_PIPE),
	6137:  uint16(anon_sym_AMP_AMP),
	6138:  uint16(anon_sym_BANG_EQ),
	6139:  uint16(anon_sym_LT),
	6140:  uint16(anon_sym_GT),
	6141:  uint16(anon_sym_LT_EQ),
	6142:  uint16(anon_sym_GT_EQ),
	6143:  uint16(2),
	6144:  uint16(3),
	6145:  uint16(1),
	6146:  uint16(sym_comment),
	6147:  uint16(339),
	6148:  uint16(17),
	6149:  uint16(anon_sym_mainmenu),
	6150:  uint16(anon_sym_config),
	6151:  uint16(anon_sym_configdefault),
	6152:  uint16(anon_sym_menuconfig),
	6153:  uint16(anon_sym_choice),
	6154:  uint16(anon_sym_endchoice),
	6155:  uint16(anon_sym_comment),
	6156:  uint16(anon_sym_menu),
	6157:  uint16(anon_sym_endmenu),
	6158:  uint16(anon_sym_if),
	6159:  uint16(anon_sym_endif),
	6160:  uint16(anon_sym_source),
	6161:  uint16(anon_sym_rsource),
	6162:  uint16(anon_sym_osource),
	6163:  uint16(anon_sym_orsource),
	6164:  uint16(anon_sym_default),
	6165:  uint16(sym_symbol),
	6166:  uint16(14),
	6167:  uint16(568),
	6168:  uint16(1),
	6169:  uint16(sym_comment),
	6170:  uint16(572),
	6171:  uint16(1),
	6172:  uint16(anon_sym_DOLLAR_LPAREN),
	6173:  uint16(574),
	6174:  uint16(1),
	6175:  uint16(anon_sym_DQUOTE),
	6176:  uint16(576),
	6177:  uint16(1),
	6178:  uint16(anon_sym_SQUOTE),
	6179:  uint16(578),
	6180:  uint16(1),
	6181:  uint16(sym_symbol),
	6182:  uint16(584),
	6183:  uint16(1),
	6184:  uint16(anon_sym_BANG),
	6185:  uint16(586),
	6186:  uint16(1),
	6187:  uint16(anon_sym_LPAREN),
	6188:  uint16(617),
	6189:  uint16(1),
	6190:  uint16(anon_sym_COMMA),
	6191:  uint16(619),
	6192:  uint16(1),
	6193:  uint16(aux_sym_variable_token1),
	6194:  uint16(124),
	6195:  uint16(1),
	6196:  uint16(aux_sym_name_repeat1),
	6197:  uint16(135),
	6198:  uint16(1),
	6199:  uint16(aux_sym_variable_repeat1),
	6200:  uint16(149),
	6201:  uint16(1),
	6202:  uint16(sym_expression),
	6203:  uint16(157),
	6204:  uint16(2),
	6205:  uint16(sym_macro_variable),
	6206:  uint16(sym_string),
	6207:  uint16(153),
	6208:  uint16(4),
	6209:  uint16(sym_unary_expression),
	6210:  uint16(sym_binary_expression),
	6211:  uint16(sym_parenthesized_expression),
	6212:  uint16(sym_name),
	6213:  uint16(8),
	6214:  uint16(93),
	6215:  uint16(1),
	6216:  uint16(aux_sym_type_definition_token1),
	6217:  uint16(568),
	6218:  uint16(1),
	6219:  uint16(sym_comment),
	6220:  uint16(621),
	6221:  uint16(1),
	6222:  uint16(sym_symbol),
	6223:  uint16(623),
	6224:  uint16(1),
	6225:  uint16(anon_sym_DOLLAR_LPAREN),
	6226:  uint16(625),
	6227:  uint16(1),
	6228:  uint16(anon_sym_DQUOTE),
	6229:  uint16(627),
	6230:  uint16(1),
	6231:  uint16(anon_sym_SQUOTE),
	6232:  uint16(130),
	6233:  uint16(3),
	6234:  uint16(sym_macro_variable),
	6235:  uint16(sym_string),
	6236:  uint16(aux_sym_name_repeat1),
	6237:  uint16(91),
	6238:  uint16(9),
	6239:  uint16(anon_sym_if),
	6240:  uint16(anon_sym_EQ),
	6241:  uint16(anon_sym_PIPE_PIPE),
	6242:  uint16(anon_sym_AMP_AMP),
	6243:  uint16(anon_sym_BANG_EQ),
	6244:  uint16(anon_sym_LT),
	6245:  uint16(anon_sym_GT),
	6246:  uint16(anon_sym_LT_EQ),
	6247:  uint16(anon_sym_GT_EQ),
	6248:  uint16(14),
	6249:  uint16(568),
	6250:  uint16(1),
	6251:  uint16(sym_comment),
	6252:  uint16(572),
	6253:  uint16(1),
	6254:  uint16(anon_sym_DOLLAR_LPAREN),
	6255:  uint16(574),
	6256:  uint16(1),
	6257:  uint16(anon_sym_DQUOTE),
	6258:  uint16(576),
	6259:  uint16(1),
	6260:  uint16(anon_sym_SQUOTE),
	6261:  uint16(578),
	6262:  uint16(1),
	6263:  uint16(sym_symbol),
	6264:  uint16(584),
	6265:  uint16(1),
	6266:  uint16(anon_sym_BANG),
	6267:  uint16(586),
	6268:  uint16(1),
	6269:  uint16(anon_sym_LPAREN),
	6270:  uint16(617),
	6271:  uint16(1),
	6272:  uint16(anon_sym_COMMA),
	6273:  uint16(629),
	6274:  uint16(1),
	6275:  uint16(aux_sym_variable_token1),
	6276:  uint16(124),
	6277:  uint16(1),
	6278:  uint16(aux_sym_name_repeat1),
	6279:  uint16(135),
	6280:  uint16(1),
	6281:  uint16(aux_sym_variable_repeat1),
	6282:  uint16(149),
	6283:  uint16(1),
	6284:  uint16(sym_expression),
	6285:  uint16(157),
	6286:  uint16(2),
	6287:  uint16(sym_macro_variable),
	6288:  uint16(sym_string),
	6289:  uint16(153),
	6290:  uint16(4),
	6291:  uint16(sym_unary_expression),
	6292:  uint16(sym_binary_expression),
	6293:  uint16(sym_parenthesized_expression),
	6294:  uint16(sym_name),
	6295:  uint16(14),
	6296:  uint16(568),
	6297:  uint16(1),
	6298:  uint16(sym_comment),
	6299:  uint16(631),
	6300:  uint16(1),
	6301:  uint16(sym_symbol),
	6302:  uint16(634),
	6303:  uint16(1),
	6304:  uint16(anon_sym_COMMA),
	6305:  uint16(637),
	6306:  uint16(1),
	6307:  uint16(aux_sym_variable_token1),
	6308:  uint16(639),
	6309:  uint16(1),
	6310:  uint16(anon_sym_BANG),
	6311:  uint16(642),
	6312:  uint16(1),
	6313:  uint16(anon_sym_LPAREN),
	6314:  uint16(645),
	6315:  uint16(1),
	6316:  uint16(anon_sym_DOLLAR_LPAREN),
	6317:  uint16(648),
	6318:  uint16(1),
	6319:  uint16(anon_sym_DQUOTE),
	6320:  uint16(651),
	6321:  uint16(1),
	6322:  uint16(anon_sym_SQUOTE),
	6323:  uint16(124),
	6324:  uint16(1),
	6325:  uint16(aux_sym_name_repeat1),
	6326:  uint16(135),
	6327:  uint16(1),
	6328:  uint16(aux_sym_variable_repeat1),
	6329:  uint16(149),
	6330:  uint16(1),
	6331:  uint16(sym_expression),
	6332:  uint16(157),
	6333:  uint16(2),
	6334:  uint16(sym_macro_variable),
	6335:  uint16(sym_string),
	6336:  uint16(153),
	6337:  uint16(4),
	6338:  uint16(sym_unary_expression),
	6339:  uint16(sym_binary_expression),
	6340:  uint16(sym_parenthesized_expression),
	6341:  uint16(sym_name),
	6342:  uint16(2),
	6343:  uint16(3),
	6344:  uint16(1),
	6345:  uint16(sym_comment),
	6346:  uint16(654),
	6347:  uint16(16),
	6348:  uint16(anon_sym_mainmenu),
	6349:  uint16(anon_sym_config),
	6350:  uint16(anon_sym_configdefault),
	6351:  uint16(anon_sym_menuconfig),
	6352:  uint16(anon_sym_choice),
	6353:  uint16(anon_sym_endchoice),
	6354:  uint16(anon_sym_comment),
	6355:  uint16(anon_sym_menu),
	6356:  uint16(anon_sym_endmenu),
	6357:  uint16(anon_sym_if),
	6358:  uint16(anon_sym_endif),
	6359:  uint16(anon_sym_source),
	6360:  uint16(anon_sym_rsource),
	6361:  uint16(anon_sym_osource),
	6362:  uint16(anon_sym_orsource),
	6363:  uint16(sym_symbol),
	6364:  uint16(2),
	6365:  uint16(3),
	6366:  uint16(1),
	6367:  uint16(sym_comment),
	6368:  uint16(656),
	6369:  uint16(16),
	6370:  uint16(anon_sym_mainmenu),
	6371:  uint16(anon_sym_config),
	6372:  uint16(anon_sym_configdefault),
	6373:  uint16(anon_sym_menuconfig),
	6374:  uint16(anon_sym_choice),
	6375:  uint16(anon_sym_endchoice),
	6376:  uint16(anon_sym_comment),
	6377:  uint16(anon_sym_menu),
	6378:  uint16(anon_sym_endmenu),
	6379:  uint16(anon_sym_if),
	6380:  uint16(anon_sym_endif),
	6381:  uint16(anon_sym_source),
	6382:  uint16(anon_sym_rsource),
	6383:  uint16(anon_sym_osource),
	6384:  uint16(anon_sym_orsource),
	6385:  uint16(sym_symbol),
	6386:  uint16(2),
	6387:  uint16(3),
	6388:  uint16(1),
	6389:  uint16(sym_comment),
	6390:  uint16(658),
	6391:  uint16(16),
	6392:  uint16(anon_sym_mainmenu),
	6393:  uint16(anon_sym_config),
	6394:  uint16(anon_sym_configdefault),
	6395:  uint16(anon_sym_menuconfig),
	6396:  uint16(anon_sym_choice),
	6397:  uint16(anon_sym_endchoice),
	6398:  uint16(anon_sym_comment),
	6399:  uint16(anon_sym_menu),
	6400:  uint16(anon_sym_endmenu),
	6401:  uint16(anon_sym_if),
	6402:  uint16(anon_sym_endif),
	6403:  uint16(anon_sym_source),
	6404:  uint16(anon_sym_rsource),
	6405:  uint16(anon_sym_osource),
	6406:  uint16(anon_sym_orsource),
	6407:  uint16(sym_symbol),
	6408:  uint16(2),
	6409:  uint16(3),
	6410:  uint16(1),
	6411:  uint16(sym_comment),
	6412:  uint16(660),
	6413:  uint16(16),
	6414:  uint16(anon_sym_mainmenu),
	6415:  uint16(anon_sym_config),
	6416:  uint16(anon_sym_configdefault),
	6417:  uint16(anon_sym_menuconfig),
	6418:  uint16(anon_sym_choice),
	6419:  uint16(anon_sym_endchoice),
	6420:  uint16(anon_sym_comment),
	6421:  uint16(anon_sym_menu),
	6422:  uint16(anon_sym_endmenu),
	6423:  uint16(anon_sym_if),
	6424:  uint16(anon_sym_endif),
	6425:  uint16(anon_sym_source),
	6426:  uint16(anon_sym_rsource),
	6427:  uint16(anon_sym_osource),
	6428:  uint16(anon_sym_orsource),
	6429:  uint16(sym_symbol),
	6430:  uint16(2),
	6431:  uint16(3),
	6432:  uint16(1),
	6433:  uint16(sym_comment),
	6434:  uint16(662),
	6435:  uint16(16),
	6436:  uint16(anon_sym_mainmenu),
	6437:  uint16(anon_sym_config),
	6438:  uint16(anon_sym_configdefault),
	6439:  uint16(anon_sym_menuconfig),
	6440:  uint16(anon_sym_choice),
	6441:  uint16(anon_sym_endchoice),
	6442:  uint16(anon_sym_comment),
	6443:  uint16(anon_sym_menu),
	6444:  uint16(anon_sym_endmenu),
	6445:  uint16(anon_sym_if),
	6446:  uint16(anon_sym_endif),
	6447:  uint16(anon_sym_source),
	6448:  uint16(anon_sym_rsource),
	6449:  uint16(anon_sym_osource),
	6450:  uint16(anon_sym_orsource),
	6451:  uint16(sym_symbol),
	6452:  uint16(2),
	6453:  uint16(3),
	6454:  uint16(1),
	6455:  uint16(sym_comment),
	6456:  uint16(664),
	6457:  uint16(16),
	6458:  uint16(anon_sym_mainmenu),
	6459:  uint16(anon_sym_config),
	6460:  uint16(anon_sym_configdefault),
	6461:  uint16(anon_sym_menuconfig),
	6462:  uint16(anon_sym_choice),
	6463:  uint16(anon_sym_endchoice),
	6464:  uint16(anon_sym_comment),
	6465:  uint16(anon_sym_menu),
	6466:  uint16(anon_sym_endmenu),
	6467:  uint16(anon_sym_if),
	6468:  uint16(anon_sym_endif),
	6469:  uint16(anon_sym_source),
	6470:  uint16(anon_sym_rsource),
	6471:  uint16(anon_sym_osource),
	6472:  uint16(anon_sym_orsource),
	6473:  uint16(sym_symbol),
	6474:  uint16(2),
	6475:  uint16(3),
	6476:  uint16(1),
	6477:  uint16(sym_comment),
	6478:  uint16(666),
	6479:  uint16(16),
	6480:  uint16(anon_sym_mainmenu),
	6481:  uint16(anon_sym_config),
	6482:  uint16(anon_sym_configdefault),
	6483:  uint16(anon_sym_menuconfig),
	6484:  uint16(anon_sym_choice),
	6485:  uint16(anon_sym_endchoice),
	6486:  uint16(anon_sym_comment),
	6487:  uint16(anon_sym_menu),
	6488:  uint16(anon_sym_endmenu),
	6489:  uint16(anon_sym_if),
	6490:  uint16(anon_sym_endif),
	6491:  uint16(anon_sym_source),
	6492:  uint16(anon_sym_rsource),
	6493:  uint16(anon_sym_osource),
	6494:  uint16(anon_sym_orsource),
	6495:  uint16(sym_symbol),
	6496:  uint16(2),
	6497:  uint16(3),
	6498:  uint16(1),
	6499:  uint16(sym_comment),
	6500:  uint16(668),
	6501:  uint16(16),
	6502:  uint16(anon_sym_mainmenu),
	6503:  uint16(anon_sym_config),
	6504:  uint16(anon_sym_configdefault),
	6505:  uint16(anon_sym_menuconfig),
	6506:  uint16(anon_sym_choice),
	6507:  uint16(anon_sym_endchoice),
	6508:  uint16(anon_sym_comment),
	6509:  uint16(anon_sym_menu),
	6510:  uint16(anon_sym_endmenu),
	6511:  uint16(anon_sym_if),
	6512:  uint16(anon_sym_endif),
	6513:  uint16(anon_sym_source),
	6514:  uint16(anon_sym_rsource),
	6515:  uint16(anon_sym_osource),
	6516:  uint16(anon_sym_orsource),
	6517:  uint16(sym_symbol),
	6518:  uint16(2),
	6519:  uint16(3),
	6520:  uint16(1),
	6521:  uint16(sym_comment),
	6522:  uint16(670),
	6523:  uint16(16),
	6524:  uint16(anon_sym_mainmenu),
	6525:  uint16(anon_sym_config),
	6526:  uint16(anon_sym_configdefault),
	6527:  uint16(anon_sym_menuconfig),
	6528:  uint16(anon_sym_choice),
	6529:  uint16(anon_sym_endchoice),
	6530:  uint16(anon_sym_comment),
	6531:  uint16(anon_sym_menu),
	6532:  uint16(anon_sym_endmenu),
	6533:  uint16(anon_sym_if),
	6534:  uint16(anon_sym_endif),
	6535:  uint16(anon_sym_source),
	6536:  uint16(anon_sym_rsource),
	6537:  uint16(anon_sym_osource),
	6538:  uint16(anon_sym_orsource),
	6539:  uint16(sym_symbol),
	6540:  uint16(2),
	6541:  uint16(3),
	6542:  uint16(1),
	6543:  uint16(sym_comment),
	6544:  uint16(672),
	6545:  uint16(16),
	6546:  uint16(anon_sym_mainmenu),
	6547:  uint16(anon_sym_config),
	6548:  uint16(anon_sym_configdefault),
	6549:  uint16(anon_sym_menuconfig),
	6550:  uint16(anon_sym_choice),
	6551:  uint16(anon_sym_endchoice),
	6552:  uint16(anon_sym_comment),
	6553:  uint16(anon_sym_menu),
	6554:  uint16(anon_sym_endmenu),
	6555:  uint16(anon_sym_if),
	6556:  uint16(anon_sym_endif),
	6557:  uint16(anon_sym_source),
	6558:  uint16(anon_sym_rsource),
	6559:  uint16(anon_sym_osource),
	6560:  uint16(anon_sym_orsource),
	6561:  uint16(sym_symbol),
	6562:  uint16(2),
	6563:  uint16(3),
	6564:  uint16(1),
	6565:  uint16(sym_comment),
	6566:  uint16(674),
	6567:  uint16(16),
	6568:  uint16(anon_sym_mainmenu),
	6569:  uint16(anon_sym_config),
	6570:  uint16(anon_sym_configdefault),
	6571:  uint16(anon_sym_menuconfig),
	6572:  uint16(anon_sym_choice),
	6573:  uint16(anon_sym_endchoice),
	6574:  uint16(anon_sym_comment),
	6575:  uint16(anon_sym_menu),
	6576:  uint16(anon_sym_endmenu),
	6577:  uint16(anon_sym_if),
	6578:  uint16(anon_sym_endif),
	6579:  uint16(anon_sym_source),
	6580:  uint16(anon_sym_rsource),
	6581:  uint16(anon_sym_osource),
	6582:  uint16(anon_sym_orsource),
	6583:  uint16(sym_symbol),
	6584:  uint16(2),
	6585:  uint16(3),
	6586:  uint16(1),
	6587:  uint16(sym_comment),
	6588:  uint16(676),
	6589:  uint16(16),
	6590:  uint16(anon_sym_mainmenu),
	6591:  uint16(anon_sym_config),
	6592:  uint16(anon_sym_configdefault),
	6593:  uint16(anon_sym_menuconfig),
	6594:  uint16(anon_sym_choice),
	6595:  uint16(anon_sym_endchoice),
	6596:  uint16(anon_sym_comment),
	6597:  uint16(anon_sym_menu),
	6598:  uint16(anon_sym_endmenu),
	6599:  uint16(anon_sym_if),
	6600:  uint16(anon_sym_endif),
	6601:  uint16(anon_sym_source),
	6602:  uint16(anon_sym_rsource),
	6603:  uint16(anon_sym_osource),
	6604:  uint16(anon_sym_orsource),
	6605:  uint16(sym_symbol),
	6606:  uint16(2),
	6607:  uint16(3),
	6608:  uint16(1),
	6609:  uint16(sym_comment),
	6610:  uint16(678),
	6611:  uint16(16),
	6612:  uint16(anon_sym_mainmenu),
	6613:  uint16(anon_sym_config),
	6614:  uint16(anon_sym_configdefault),
	6615:  uint16(anon_sym_menuconfig),
	6616:  uint16(anon_sym_choice),
	6617:  uint16(anon_sym_endchoice),
	6618:  uint16(anon_sym_comment),
	6619:  uint16(anon_sym_menu),
	6620:  uint16(anon_sym_endmenu),
	6621:  uint16(anon_sym_if),
	6622:  uint16(anon_sym_endif),
	6623:  uint16(anon_sym_source),
	6624:  uint16(anon_sym_rsource),
	6625:  uint16(anon_sym_osource),
	6626:  uint16(anon_sym_orsource),
	6627:  uint16(sym_symbol),
	6628:  uint16(7),
	6629:  uint16(568),
	6630:  uint16(1),
	6631:  uint16(sym_comment),
	6632:  uint16(682),
	6633:  uint16(1),
	6634:  uint16(anon_sym_EQ),
	6635:  uint16(684),
	6636:  uint16(1),
	6637:  uint16(aux_sym_variable_token1),
	6638:  uint16(686),
	6639:  uint16(1),
	6640:  uint16(anon_sym_PIPE_PIPE),
	6641:  uint16(688),
	6642:  uint16(1),
	6643:  uint16(anon_sym_AMP_AMP),
	6644:  uint16(690),
	6645:  uint16(5),
	6646:  uint16(anon_sym_BANG_EQ),
	6647:  uint16(anon_sym_LT),
	6648:  uint16(anon_sym_GT),
	6649:  uint16(anon_sym_LT_EQ),
	6650:  uint16(anon_sym_GT_EQ),
	6651:  uint16(680),
	6652:  uint16(7),
	6653:  uint16(anon_sym_COMMA),
	6654:  uint16(anon_sym_BANG),
	6655:  uint16(anon_sym_LPAREN),
	6656:  uint16(anon_sym_DOLLAR_LPAREN),
	6657:  uint16(anon_sym_DQUOTE),
	6658:  uint16(anon_sym_SQUOTE),
	6659:  uint16(sym_symbol),
	6660:  uint16(2),
	6661:  uint16(3),
	6662:  uint16(1),
	6663:  uint16(sym_comment),
	6664:  uint16(692),
	6665:  uint16(16),
	6666:  uint16(anon_sym_mainmenu),
	6667:  uint16(anon_sym_config),
	6668:  uint16(anon_sym_configdefault),
	6669:  uint16(anon_sym_menuconfig),
	6670:  uint16(anon_sym_choice),
	6671:  uint16(anon_sym_endchoice),
	6672:  uint16(anon_sym_comment),
	6673:  uint16(anon_sym_menu),
	6674:  uint16(anon_sym_endmenu),
	6675:  uint16(anon_sym_if),
	6676:  uint16(anon_sym_endif),
	6677:  uint16(anon_sym_source),
	6678:  uint16(anon_sym_rsource),
	6679:  uint16(anon_sym_osource),
	6680:  uint16(anon_sym_orsource),
	6681:  uint16(sym_symbol),
	6682:  uint16(3),
	6683:  uint16(184),
	6684:  uint16(1),
	6685:  uint16(aux_sym_variable_token1),
	6686:  uint16(568),
	6687:  uint16(1),
	6688:  uint16(sym_comment),
	6689:  uint16(182),
	6690:  uint16(15),
	6691:  uint16(anon_sym_EQ),
	6692:  uint16(anon_sym_COMMA),
	6693:  uint16(anon_sym_BANG),
	6694:  uint16(anon_sym_PIPE_PIPE),
	6695:  uint16(anon_sym_AMP_AMP),
	6696:  uint16(anon_sym_BANG_EQ),
	6697:  uint16(anon_sym_LT),
	6698:  uint16(anon_sym_GT),
	6699:  uint16(anon_sym_LT_EQ),
	6700:  uint16(anon_sym_GT_EQ),
	6701:  uint16(anon_sym_LPAREN),
	6702:  uint16(anon_sym_DOLLAR_LPAREN),
	6703:  uint16(anon_sym_DQUOTE),
	6704:  uint16(anon_sym_SQUOTE),
	6705:  uint16(sym_symbol),
	6706:  uint16(3),
	6707:  uint16(200),
	6708:  uint16(1),
	6709:  uint16(aux_sym_variable_token1),
	6710:  uint16(568),
	6711:  uint16(1),
	6712:  uint16(sym_comment),
	6713:  uint16(198),
	6714:  uint16(15),
	6715:  uint16(anon_sym_EQ),
	6716:  uint16(anon_sym_COMMA),
	6717:  uint16(anon_sym_BANG),
	6718:  uint16(anon_sym_PIPE_PIPE),
	6719:  uint16(anon_sym_AMP_AMP),
	6720:  uint16(anon_sym_BANG_EQ),
	6721:  uint16(anon_sym_LT),
	6722:  uint16(anon_sym_GT),
	6723:  uint16(anon_sym_LT_EQ),
	6724:  uint16(anon_sym_GT_EQ),
	6725:  uint16(anon_sym_LPAREN),
	6726:  uint16(anon_sym_DOLLAR_LPAREN),
	6727:  uint16(anon_sym_DQUOTE),
	6728:  uint16(anon_sym_SQUOTE),
	6729:  uint16(sym_symbol),
	6730:  uint16(3),
	6731:  uint16(289),
	6732:  uint16(1),
	6733:  uint16(aux_sym_variable_token1),
	6734:  uint16(568),
	6735:  uint16(1),
	6736:  uint16(sym_comment),
	6737:  uint16(287),
	6738:  uint16(15),
	6739:  uint16(anon_sym_EQ),
	6740:  uint16(anon_sym_COMMA),
	6741:  uint16(anon_sym_BANG),
	6742:  uint16(anon_sym_PIPE_PIPE),
	6743:  uint16(anon_sym_AMP_AMP),
	6744:  uint16(anon_sym_BANG_EQ),
	6745:  uint16(anon_sym_LT),
	6746:  uint16(anon_sym_GT),
	6747:  uint16(anon_sym_LT_EQ),
	6748:  uint16(anon_sym_GT_EQ),
	6749:  uint16(anon_sym_LPAREN),
	6750:  uint16(anon_sym_DOLLAR_LPAREN),
	6751:  uint16(anon_sym_DQUOTE),
	6752:  uint16(anon_sym_SQUOTE),
	6753:  uint16(sym_symbol),
	6754:  uint16(3),
	6755:  uint16(275),
	6756:  uint16(1),
	6757:  uint16(aux_sym_variable_token1),
	6758:  uint16(568),
	6759:  uint16(1),
	6760:  uint16(sym_comment),
	6761:  uint16(273),
	6762:  uint16(15),
	6763:  uint16(anon_sym_EQ),
	6764:  uint16(anon_sym_COMMA),
	6765:  uint16(anon_sym_BANG),
	6766:  uint16(anon_sym_PIPE_PIPE),
	6767:  uint16(anon_sym_AMP_AMP),
	6768:  uint16(anon_sym_BANG_EQ),
	6769:  uint16(anon_sym_LT),
	6770:  uint16(anon_sym_GT),
	6771:  uint16(anon_sym_LT_EQ),
	6772:  uint16(anon_sym_GT_EQ),
	6773:  uint16(anon_sym_LPAREN),
	6774:  uint16(anon_sym_DOLLAR_LPAREN),
	6775:  uint16(anon_sym_DQUOTE),
	6776:  uint16(anon_sym_SQUOTE),
	6777:  uint16(sym_symbol),
	6778:  uint16(3),
	6779:  uint16(293),
	6780:  uint16(1),
	6781:  uint16(aux_sym_variable_token1),
	6782:  uint16(568),
	6783:  uint16(1),
	6784:  uint16(sym_comment),
	6785:  uint16(291),
	6786:  uint16(15),
	6787:  uint16(anon_sym_EQ),
	6788:  uint16(anon_sym_COMMA),
	6789:  uint16(anon_sym_BANG),
	6790:  uint16(anon_sym_PIPE_PIPE),
	6791:  uint16(anon_sym_AMP_AMP),
	6792:  uint16(anon_sym_BANG_EQ),
	6793:  uint16(anon_sym_LT),
	6794:  uint16(anon_sym_GT),
	6795:  uint16(anon_sym_LT_EQ),
	6796:  uint16(anon_sym_GT_EQ),
	6797:  uint16(anon_sym_LPAREN),
	6798:  uint16(anon_sym_DOLLAR_LPAREN),
	6799:  uint16(anon_sym_DQUOTE),
	6800:  uint16(anon_sym_SQUOTE),
	6801:  uint16(sym_symbol),
	6802:  uint16(3),
	6803:  uint16(279),
	6804:  uint16(1),
	6805:  uint16(aux_sym_variable_token1),
	6806:  uint16(568),
	6807:  uint16(1),
	6808:  uint16(sym_comment),
	6809:  uint16(277),
	6810:  uint16(15),
	6811:  uint16(anon_sym_EQ),
	6812:  uint16(anon_sym_COMMA),
	6813:  uint16(anon_sym_BANG),
	6814:  uint16(anon_sym_PIPE_PIPE),
	6815:  uint16(anon_sym_AMP_AMP),
	6816:  uint16(anon_sym_BANG_EQ),
	6817:  uint16(anon_sym_LT),
	6818:  uint16(anon_sym_GT),
	6819:  uint16(anon_sym_LT_EQ),
	6820:  uint16(anon_sym_GT_EQ),
	6821:  uint16(anon_sym_LPAREN),
	6822:  uint16(anon_sym_DOLLAR_LPAREN),
	6823:  uint16(anon_sym_DQUOTE),
	6824:  uint16(anon_sym_SQUOTE),
	6825:  uint16(sym_symbol),
	6826:  uint16(3),
	6827:  uint16(193),
	6828:  uint16(1),
	6829:  uint16(aux_sym_variable_token1),
	6830:  uint16(568),
	6831:  uint16(1),
	6832:  uint16(sym_comment),
	6833:  uint16(190),
	6834:  uint16(15),
	6835:  uint16(anon_sym_EQ),
	6836:  uint16(anon_sym_COMMA),
	6837:  uint16(anon_sym_BANG),
	6838:  uint16(anon_sym_PIPE_PIPE),
	6839:  uint16(anon_sym_AMP_AMP),
	6840:  uint16(anon_sym_BANG_EQ),
	6841:  uint16(anon_sym_LT),
	6842:  uint16(anon_sym_GT),
	6843:  uint16(anon_sym_LT_EQ),
	6844:  uint16(anon_sym_GT_EQ),
	6845:  uint16(anon_sym_LPAREN),
	6846:  uint16(anon_sym_DOLLAR_LPAREN),
	6847:  uint16(anon_sym_DQUOTE),
	6848:  uint16(anon_sym_SQUOTE),
	6849:  uint16(sym_symbol),
	6850:  uint16(6),
	6851:  uint16(279),
	6852:  uint16(1),
	6853:  uint16(aux_sym_variable_token1),
	6854:  uint16(568),
	6855:  uint16(1),
	6856:  uint16(sym_comment),
	6857:  uint16(682),
	6858:  uint16(1),
	6859:  uint16(anon_sym_EQ),
	6860:  uint16(688),
	6861:  uint16(1),
	6862:  uint16(anon_sym_AMP_AMP),
	6863:  uint16(690),
	6864:  uint16(5),
	6865:  uint16(anon_sym_BANG_EQ),
	6866:  uint16(anon_sym_LT),
	6867:  uint16(anon_sym_GT),
	6868:  uint16(anon_sym_LT_EQ),
	6869:  uint16(anon_sym_GT_EQ),
	6870:  uint16(277),
	6871:  uint16(8),
	6872:  uint16(anon_sym_COMMA),
	6873:  uint16(anon_sym_BANG),
	6874:  uint16(anon_sym_PIPE_PIPE),
	6875:  uint16(anon_sym_LPAREN),
	6876:  uint16(anon_sym_DOLLAR_LPAREN),
	6877:  uint16(anon_sym_DQUOTE),
	6878:  uint16(anon_sym_SQUOTE),
	6879:  uint16(sym_symbol),
	6880:  uint16(5),
	6881:  uint16(279),
	6882:  uint16(1),
	6883:  uint16(aux_sym_variable_token1),
	6884:  uint16(568),
	6885:  uint16(1),
	6886:  uint16(sym_comment),
	6887:  uint16(682),
	6888:  uint16(1),
	6889:  uint16(anon_sym_EQ),
	6890:  uint16(690),
	6891:  uint16(5),
	6892:  uint16(anon_sym_BANG_EQ),
	6893:  uint16(anon_sym_LT),
	6894:  uint16(anon_sym_GT),
	6895:  uint16(anon_sym_LT_EQ),
	6896:  uint16(anon_sym_GT_EQ),
	6897:  uint16(277),
	6898:  uint16(9),
	6899:  uint16(anon_sym_COMMA),
	6900:  uint16(anon_sym_BANG),
	6901:  uint16(anon_sym_PIPE_PIPE),
	6902:  uint16(anon_sym_AMP_AMP),
	6903:  uint16(anon_sym_LPAREN),
	6904:  uint16(anon_sym_DOLLAR_LPAREN),
	6905:  uint16(anon_sym_DQUOTE),
	6906:  uint16(anon_sym_SQUOTE),
	6907:  uint16(sym_symbol),
	6908:  uint16(4),
	6909:  uint16(279),
	6910:  uint16(1),
	6911:  uint16(aux_sym_variable_token1),
	6912:  uint16(568),
	6913:  uint16(1),
	6914:  uint16(sym_comment),
	6915:  uint16(682),
	6916:  uint16(1),
	6917:  uint16(anon_sym_EQ),
	6918:  uint16(277),
	6919:  uint16(14),
	6920:  uint16(anon_sym_COMMA),
	6921:  uint16(anon_sym_BANG),
	6922:  uint16(anon_sym_PIPE_PIPE),
	6923:  uint16(anon_sym_AMP_AMP),
	6924:  uint16(anon_sym_BANG_EQ),
	6925:  uint16(anon_sym_LT),
	6926:  uint16(anon_sym_GT),
	6927:  uint16(anon_sym_LT_EQ),
	6928:  uint16(anon_sym_GT_EQ),
	6929:  uint16(anon_sym_LPAREN),
	6930:  uint16(anon_sym_DOLLAR_LPAREN),
	6931:  uint16(anon_sym_DQUOTE),
	6932:  uint16(anon_sym_SQUOTE),
	6933:  uint16(sym_symbol),
	6934:  uint16(3),
	6935:  uint16(204),
	6936:  uint16(1),
	6937:  uint16(aux_sym_variable_token1),
	6938:  uint16(568),
	6939:  uint16(1),
	6940:  uint16(sym_comment),
	6941:  uint16(202),
	6942:  uint16(15),
	6943:  uint16(anon_sym_EQ),
	6944:  uint16(anon_sym_COMMA),
	6945:  uint16(anon_sym_BANG),
	6946:  uint16(anon_sym_PIPE_PIPE),
	6947:  uint16(anon_sym_AMP_AMP),
	6948:  uint16(anon_sym_BANG_EQ),
	6949:  uint16(anon_sym_LT),
	6950:  uint16(anon_sym_GT),
	6951:  uint16(anon_sym_LT_EQ),
	6952:  uint16(anon_sym_GT_EQ),
	6953:  uint16(anon_sym_LPAREN),
	6954:  uint16(anon_sym_DOLLAR_LPAREN),
	6955:  uint16(anon_sym_DQUOTE),
	6956:  uint16(anon_sym_SQUOTE),
	6957:  uint16(sym_symbol),
	6958:  uint16(3),
	6959:  uint16(188),
	6960:  uint16(1),
	6961:  uint16(aux_sym_variable_token1),
	6962:  uint16(568),
	6963:  uint16(1),
	6964:  uint16(sym_comment),
	6965:  uint16(186),
	6966:  uint16(15),
	6967:  uint16(anon_sym_EQ),
	6968:  uint16(anon_sym_COMMA),
	6969:  uint16(anon_sym_BANG),
	6970:  uint16(anon_sym_PIPE_PIPE),
	6971:  uint16(anon_sym_AMP_AMP),
	6972:  uint16(anon_sym_BANG_EQ),
	6973:  uint16(anon_sym_LT),
	6974:  uint16(anon_sym_GT),
	6975:  uint16(anon_sym_LT_EQ),
	6976:  uint16(anon_sym_GT_EQ),
	6977:  uint16(anon_sym_LPAREN),
	6978:  uint16(anon_sym_DOLLAR_LPAREN),
	6979:  uint16(anon_sym_DQUOTE),
	6980:  uint16(anon_sym_SQUOTE),
	6981:  uint16(sym_symbol),
	6982:  uint16(3),
	6983:  uint16(3),
	6984:  uint16(1),
	6985:  uint16(sym_comment),
	6986:  uint16(341),
	6987:  uint16(1),
	6989:  uint16(339),
	6990:  uint16(14),
	6991:  uint16(anon_sym_mainmenu),
	6992:  uint16(anon_sym_config),
	6993:  uint16(anon_sym_configdefault),
	6994:  uint16(anon_sym_menuconfig),
	6995:  uint16(anon_sym_choice),
	6996:  uint16(anon_sym_comment),
	6997:  uint16(anon_sym_menu),
	6998:  uint16(anon_sym_if),
	6999:  uint16(anon_sym_source),
	7000:  uint16(anon_sym_rsource),
	7001:  uint16(anon_sym_osource),
	7002:  uint16(anon_sym_orsource),
	7003:  uint16(anon_sym_default),
	7004:  uint16(sym_symbol),
	7005:  uint16(3),
	7006:  uint16(3),
	7007:  uint16(1),
	7008:  uint16(sym_comment),
	7009:  uint16(381),
	7010:  uint16(1),
	7012:  uint16(379),
	7013:  uint16(14),
	7014:  uint16(anon_sym_mainmenu),
	7015:  uint16(anon_sym_config),
	7016:  uint16(anon_sym_configdefault),
	7017:  uint16(anon_sym_menuconfig),
	7018:  uint16(anon_sym_choice),
	7019:  uint16(anon_sym_comment),
	7020:  uint16(anon_sym_menu),
	7021:  uint16(anon_sym_if),
	7022:  uint16(anon_sym_source),
	7023:  uint16(anon_sym_rsource),
	7024:  uint16(anon_sym_osource),
	7025:  uint16(anon_sym_orsource),
	7026:  uint16(anon_sym_default),
	7027:  uint16(sym_symbol),
	7028:  uint16(11),
	7029:  uint16(3),
	7030:  uint16(1),
	7031:  uint16(sym_comment),
	7032:  uint16(694),
	7033:  uint16(1),
	7034:  uint16(sym_symbol),
	7035:  uint16(696),
	7036:  uint16(1),
	7037:  uint16(anon_sym_BANG),
	7038:  uint16(698),
	7039:  uint16(1),
	7040:  uint16(anon_sym_LPAREN),
	7041:  uint16(700),
	7042:  uint16(1),
	7043:  uint16(anon_sym_DOLLAR_LPAREN),
	7044:  uint16(702),
	7045:  uint16(1),
	7046:  uint16(anon_sym_DQUOTE),
	7047:  uint16(704),
	7048:  uint16(1),
	7049:  uint16(anon_sym_SQUOTE),
	7050:  uint16(133),
	7051:  uint16(1),
	7052:  uint16(aux_sym_name_repeat1),
	7053:  uint16(255),
	7054:  uint16(1),
	7055:  uint16(sym_expression),
	7056:  uint16(227),
	7057:  uint16(2),
	7058:  uint16(sym_macro_variable),
	7059:  uint16(sym_string),
	7060:  uint16(246),
	7061:  uint16(4),
	7062:  uint16(sym_unary_expression),
	7063:  uint16(sym_binary_expression),
	7064:  uint16(sym_parenthesized_expression),
	7065:  uint16(sym_name),
	7066:  uint16(11),
	7067:  uint16(3),
	7068:  uint16(1),
	7069:  uint16(sym_comment),
	7070:  uint16(459),
	7071:  uint16(1),
	7072:  uint16(anon_sym_DOLLAR_LPAREN),
	7073:  uint16(461),
	7074:  uint16(1),
	7075:  uint16(anon_sym_DQUOTE),
	7076:  uint16(463),
	7077:  uint16(1),
	7078:  uint16(anon_sym_SQUOTE),
	7079:  uint16(706),
	7080:  uint16(1),
	7081:  uint16(sym_symbol),
	7082:  uint16(708),
	7083:  uint16(1),
	7084:  uint16(anon_sym_BANG),
	7085:  uint16(710),
	7086:  uint16(1),
	7087:  uint16(anon_sym_LPAREN),
	7088:  uint16(95),
	7089:  uint16(1),
	7090:  uint16(aux_sym_name_repeat1),
	7091:  uint16(252),
	7092:  uint16(1),
	7093:  uint16(sym_expression),
	7094:  uint16(234),
	7095:  uint16(2),
	7096:  uint16(sym_macro_variable),
	7097:  uint16(sym_string),
	7098:  uint16(117),
	7099:  uint16(4),
	7100:  uint16(sym_unary_expression),
	7101:  uint16(sym_binary_expression),
	7102:  uint16(sym_parenthesized_expression),
	7103:  uint16(sym_name),
	7104:  uint16(11),
	7105:  uint16(3),
	7106:  uint16(1),
	7107:  uint16(sym_comment),
	7108:  uint16(694),
	7109:  uint16(1),
	7110:  uint16(sym_symbol),
	7111:  uint16(696),
	7112:  uint16(1),
	7113:  uint16(anon_sym_BANG),
	7114:  uint16(698),
	7115:  uint16(1),
	7116:  uint16(anon_sym_LPAREN),
	7117:  uint16(700),
	7118:  uint16(1),
	7119:  uint16(anon_sym_DOLLAR_LPAREN),
	7120:  uint16(702),
	7121:  uint16(1),
	7122:  uint16(anon_sym_DQUOTE),
	7123:  uint16(704),
	7124:  uint16(1),
	7125:  uint16(anon_sym_SQUOTE),
	7126:  uint16(133),
	7127:  uint16(1),
	7128:  uint16(aux_sym_name_repeat1),
	7129:  uint16(237),
	7130:  uint16(1),
	7131:  uint16(sym_expression),
	7132:  uint16(227),
	7133:  uint16(2),
	7134:  uint16(sym_macro_variable),
	7135:  uint16(sym_string),
	7136:  uint16(246),
	7137:  uint16(4),
	7138:  uint16(sym_unary_expression),
	7139:  uint16(sym_binary_expression),
	7140:  uint16(sym_parenthesized_expression),
	7141:  uint16(sym_name),
	7142:  uint16(11),
	7143:  uint16(3),
	7144:  uint16(1),
	7145:  uint16(sym_comment),
	7146:  uint16(694),
	7147:  uint16(1),
	7148:  uint16(sym_symbol),
	7149:  uint16(696),
	7150:  uint16(1),
	7151:  uint16(anon_sym_BANG),
	7152:  uint16(698),
	7153:  uint16(1),
	7154:  uint16(anon_sym_LPAREN),
	7155:  uint16(700),
	7156:  uint16(1),
	7157:  uint16(anon_sym_DOLLAR_LPAREN),
	7158:  uint16(702),
	7159:  uint16(1),
	7160:  uint16(anon_sym_DQUOTE),
	7161:  uint16(704),
	7162:  uint16(1),
	7163:  uint16(anon_sym_SQUOTE),
	7164:  uint16(133),
	7165:  uint16(1),
	7166:  uint16(aux_sym_name_repeat1),
	7167:  uint16(247),
	7168:  uint16(1),
	7169:  uint16(sym_expression),
	7170:  uint16(227),
	7171:  uint16(2),
	7172:  uint16(sym_macro_variable),
	7173:  uint16(sym_string),
	7174:  uint16(246),
	7175:  uint16(4),
	7176:  uint16(sym_unary_expression),
	7177:  uint16(sym_binary_expression),
	7178:  uint16(sym_parenthesized_expression),
	7179:  uint16(sym_name),
	7180:  uint16(11),
	7181:  uint16(3),
	7182:  uint16(1),
	7183:  uint16(sym_comment),
	7184:  uint16(694),
	7185:  uint16(1),
	7186:  uint16(sym_symbol),
	7187:  uint16(696),
	7188:  uint16(1),
	7189:  uint16(anon_sym_BANG),
	7190:  uint16(698),
	7191:  uint16(1),
	7192:  uint16(anon_sym_LPAREN),
	7193:  uint16(700),
	7194:  uint16(1),
	7195:  uint16(anon_sym_DOLLAR_LPAREN),
	7196:  uint16(702),
	7197:  uint16(1),
	7198:  uint16(anon_sym_DQUOTE),
	7199:  uint16(704),
	7200:  uint16(1),
	7201:  uint16(anon_sym_SQUOTE),
	7202:  uint16(133),
	7203:  uint16(1),
	7204:  uint16(aux_sym_name_repeat1),
	7205:  uint16(243),
	7206:  uint16(1),
	7207:  uint16(sym_expression),
	7208:  uint16(227),
	7209:  uint16(2),
	7210:  uint16(sym_macro_variable),
	7211:  uint16(sym_string),
	7212:  uint16(246),
	7213:  uint16(4),
	7214:  uint16(sym_unary_expression),
	7215:  uint16(sym_binary_expression),
	7216:  uint16(sym_parenthesized_expression),
	7217:  uint16(sym_name),
	7218:  uint16(3),
	7219:  uint16(3),
	7220:  uint16(1),
	7221:  uint16(sym_comment),
	7222:  uint16(712),
	7223:  uint16(1),
	7225:  uint16(666),
	7226:  uint16(13),
	7227:  uint16(anon_sym_mainmenu),
	7228:  uint16(anon_sym_config),
	7229:  uint16(anon_sym_configdefault),
	7230:  uint16(anon_sym_menuconfig),
	7231:  uint16(anon_sym_choice),
	7232:  uint16(anon_sym_comment),
	7233:  uint16(anon_sym_menu),
	7234:  uint16(anon_sym_if),
	7235:  uint16(anon_sym_source),
	7236:  uint16(anon_sym_rsource),
	7237:  uint16(anon_sym_osource),
	7238:  uint16(anon_sym_orsource),
	7239:  uint16(sym_symbol),
	7240:  uint16(11),
	7241:  uint16(3),
	7242:  uint16(1),
	7243:  uint16(sym_comment),
	7244:  uint16(694),
	7245:  uint16(1),
	7246:  uint16(sym_symbol),
	7247:  uint16(696),
	7248:  uint16(1),
	7249:  uint16(anon_sym_BANG),
	7250:  uint16(698),
	7251:  uint16(1),
	7252:  uint16(anon_sym_LPAREN),
	7253:  uint16(700),
	7254:  uint16(1),
	7255:  uint16(anon_sym_DOLLAR_LPAREN),
	7256:  uint16(702),
	7257:  uint16(1),
	7258:  uint16(anon_sym_DQUOTE),
	7259:  uint16(704),
	7260:  uint16(1),
	7261:  uint16(anon_sym_SQUOTE),
	7262:  uint16(133),
	7263:  uint16(1),
	7264:  uint16(aux_sym_name_repeat1),
	7265:  uint16(241),
	7266:  uint16(1),
	7267:  uint16(sym_expression),
	7268:  uint16(227),
	7269:  uint16(2),
	7270:  uint16(sym_macro_variable),
	7271:  uint16(sym_string),
	7272:  uint16(246),
	7273:  uint16(4),
	7274:  uint16(sym_unary_expression),
	7275:  uint16(sym_binary_expression),
	7276:  uint16(sym_parenthesized_expression),
	7277:  uint16(sym_name),
	7278:  uint16(3),
	7279:  uint16(3),
	7280:  uint16(1),
	7281:  uint16(sym_comment),
	7282:  uint16(184),
	7283:  uint16(1),
	7285:  uint16(182),
	7286:  uint16(13),
	7287:  uint16(anon_sym_mainmenu),
	7288:  uint16(anon_sym_config),
	7289:  uint16(anon_sym_configdefault),
	7290:  uint16(anon_sym_menuconfig),
	7291:  uint16(anon_sym_choice),
	7292:  uint16(anon_sym_comment),
	7293:  uint16(anon_sym_menu),
	7294:  uint16(anon_sym_if),
	7295:  uint16(anon_sym_source),
	7296:  uint16(anon_sym_rsource),
	7297:  uint16(anon_sym_osource),
	7298:  uint16(anon_sym_orsource),
	7299:  uint16(sym_symbol),
	7300:  uint16(11),
	7301:  uint16(3),
	7302:  uint16(1),
	7303:  uint16(sym_comment),
	7304:  uint16(694),
	7305:  uint16(1),
	7306:  uint16(sym_symbol),
	7307:  uint16(696),
	7308:  uint16(1),
	7309:  uint16(anon_sym_BANG),
	7310:  uint16(698),
	7311:  uint16(1),
	7312:  uint16(anon_sym_LPAREN),
	7313:  uint16(700),
	7314:  uint16(1),
	7315:  uint16(anon_sym_DOLLAR_LPAREN),
	7316:  uint16(702),
	7317:  uint16(1),
	7318:  uint16(anon_sym_DQUOTE),
	7319:  uint16(704),
	7320:  uint16(1),
	7321:  uint16(anon_sym_SQUOTE),
	7322:  uint16(133),
	7323:  uint16(1),
	7324:  uint16(aux_sym_name_repeat1),
	7325:  uint16(242),
	7326:  uint16(1),
	7327:  uint16(sym_expression),
	7328:  uint16(227),
	7329:  uint16(2),
	7330:  uint16(sym_macro_variable),
	7331:  uint16(sym_string),
	7332:  uint16(246),
	7333:  uint16(4),
	7334:  uint16(sym_unary_expression),
	7335:  uint16(sym_binary_expression),
	7336:  uint16(sym_parenthesized_expression),
	7337:  uint16(sym_name),
	7338:  uint16(11),
	7339:  uint16(3),
	7340:  uint16(1),
	7341:  uint16(sym_comment),
	7342:  uint16(459),
	7343:  uint16(1),
	7344:  uint16(anon_sym_DOLLAR_LPAREN),
	7345:  uint16(461),
	7346:  uint16(1),
	7347:  uint16(anon_sym_DQUOTE),
	7348:  uint16(463),
	7349:  uint16(1),
	7350:  uint16(anon_sym_SQUOTE),
	7351:  uint16(710),
	7352:  uint16(1),
	7353:  uint16(anon_sym_LPAREN),
	7354:  uint16(714),
	7355:  uint16(1),
	7356:  uint16(sym_symbol),
	7357:  uint16(716),
	7358:  uint16(1),
	7359:  uint16(anon_sym_BANG),
	7360:  uint16(86),
	7361:  uint16(1),
	7362:  uint16(sym_expression),
	7363:  uint16(95),
	7364:  uint16(1),
	7365:  uint16(aux_sym_name_repeat1),
	7366:  uint16(113),
	7367:  uint16(2),
	7368:  uint16(sym_macro_variable),
	7369:  uint16(sym_string),
	7370:  uint16(117),
	7371:  uint16(4),
	7372:  uint16(sym_unary_expression),
	7373:  uint16(sym_binary_expression),
	7374:  uint16(sym_parenthesized_expression),
	7375:  uint16(sym_name),
	7376:  uint16(11),
	7377:  uint16(3),
	7378:  uint16(1),
	7379:  uint16(sym_comment),
	7380:  uint16(459),
	7381:  uint16(1),
	7382:  uint16(anon_sym_DOLLAR_LPAREN),
	7383:  uint16(461),
	7384:  uint16(1),
	7385:  uint16(anon_sym_DQUOTE),
	7386:  uint16(463),
	7387:  uint16(1),
	7388:  uint16(anon_sym_SQUOTE),
	7389:  uint16(710),
	7390:  uint16(1),
	7391:  uint16(anon_sym_LPAREN),
	7392:  uint16(714),
	7393:  uint16(1),
	7394:  uint16(sym_symbol),
	7395:  uint16(716),
	7396:  uint16(1),
	7397:  uint16(anon_sym_BANG),
	7398:  uint16(95),
	7399:  uint16(1),
	7400:  uint16(aux_sym_name_repeat1),
	7401:  uint16(115),
	7402:  uint16(1),
	7403:  uint16(sym_expression),
	7404:  uint16(113),
	7405:  uint16(2),
	7406:  uint16(sym_macro_variable),
	7407:  uint16(sym_string),
	7408:  uint16(117),
	7409:  uint16(4),
	7410:  uint16(sym_unary_expression),
	7411:  uint16(sym_binary_expression),
	7412:  uint16(sym_parenthesized_expression),
	7413:  uint16(sym_name),
	7414:  uint16(11),
	7415:  uint16(3),
	7416:  uint16(1),
	7417:  uint16(sym_comment),
	7418:  uint16(95),
	7419:  uint16(1),
	7420:  uint16(anon_sym_DOLLAR_LPAREN),
	7421:  uint16(97),
	7422:  uint16(1),
	7423:  uint16(anon_sym_DQUOTE),
	7424:  uint16(99),
	7425:  uint16(1),
	7426:  uint16(anon_sym_SQUOTE),
	7427:  uint16(718),
	7428:  uint16(1),
	7429:  uint16(sym_symbol),
	7430:  uint16(720),
	7431:  uint16(1),
	7432:  uint16(anon_sym_BANG),
	7433:  uint16(722),
	7434:  uint16(1),
	7435:  uint16(anon_sym_LPAREN),
	7436:  uint16(10),
	7437:  uint16(1),
	7438:  uint16(aux_sym_name_repeat1),
	7439:  uint16(35),
	7440:  uint16(1),
	7441:  uint16(sym_expression),
	7442:  uint16(21),
	7443:  uint16(2),
	7444:  uint16(sym_macro_variable),
	7445:  uint16(sym_string),
	7446:  uint16(37),
	7447:  uint16(4),
	7448:  uint16(sym_unary_expression),
	7449:  uint16(sym_binary_expression),
	7450:  uint16(sym_parenthesized_expression),
	7451:  uint16(sym_name),
	7452:  uint16(11),
	7453:  uint16(3),
	7454:  uint16(1),
	7455:  uint16(sym_comment),
	7456:  uint16(95),
	7457:  uint16(1),
	7458:  uint16(anon_sym_DOLLAR_LPAREN),
	7459:  uint16(97),
	7460:  uint16(1),
	7461:  uint16(anon_sym_DQUOTE),
	7462:  uint16(99),
	7463:  uint16(1),
	7464:  uint16(anon_sym_SQUOTE),
	7465:  uint16(718),
	7466:  uint16(1),
	7467:  uint16(sym_symbol),
	7468:  uint16(720),
	7469:  uint16(1),
	7470:  uint16(anon_sym_BANG),
	7471:  uint16(722),
	7472:  uint16(1),
	7473:  uint16(anon_sym_LPAREN),
	7474:  uint16(10),
	7475:  uint16(1),
	7476:  uint16(aux_sym_name_repeat1),
	7477:  uint16(40),
	7478:  uint16(1),
	7479:  uint16(sym_expression),
	7480:  uint16(21),
	7481:  uint16(2),
	7482:  uint16(sym_macro_variable),
	7483:  uint16(sym_string),
	7484:  uint16(37),
	7485:  uint16(4),
	7486:  uint16(sym_unary_expression),
	7487:  uint16(sym_binary_expression),
	7488:  uint16(sym_parenthesized_expression),
	7489:  uint16(sym_name),
	7490:  uint16(11),
	7491:  uint16(3),
	7492:  uint16(1),
	7493:  uint16(sym_comment),
	7494:  uint16(95),
	7495:  uint16(1),
	7496:  uint16(anon_sym_DOLLAR_LPAREN),
	7497:  uint16(97),
	7498:  uint16(1),
	7499:  uint16(anon_sym_DQUOTE),
	7500:  uint16(99),
	7501:  uint16(1),
	7502:  uint16(anon_sym_SQUOTE),
	7503:  uint16(718),
	7504:  uint16(1),
	7505:  uint16(sym_symbol),
	7506:  uint16(720),
	7507:  uint16(1),
	7508:  uint16(anon_sym_BANG),
	7509:  uint16(722),
	7510:  uint16(1),
	7511:  uint16(anon_sym_LPAREN),
	7512:  uint16(10),
	7513:  uint16(1),
	7514:  uint16(aux_sym_name_repeat1),
	7515:  uint16(36),
	7516:  uint16(1),
	7517:  uint16(sym_expression),
	7518:  uint16(21),
	7519:  uint16(2),
	7520:  uint16(sym_macro_variable),
	7521:  uint16(sym_string),
	7522:  uint16(37),
	7523:  uint16(4),
	7524:  uint16(sym_unary_expression),
	7525:  uint16(sym_binary_expression),
	7526:  uint16(sym_parenthesized_expression),
	7527:  uint16(sym_name),
	7528:  uint16(11),
	7529:  uint16(3),
	7530:  uint16(1),
	7531:  uint16(sym_comment),
	7532:  uint16(95),
	7533:  uint16(1),
	7534:  uint16(anon_sym_DOLLAR_LPAREN),
	7535:  uint16(97),
	7536:  uint16(1),
	7537:  uint16(anon_sym_DQUOTE),
	7538:  uint16(99),
	7539:  uint16(1),
	7540:  uint16(anon_sym_SQUOTE),
	7541:  uint16(718),
	7542:  uint16(1),
	7543:  uint16(sym_symbol),
	7544:  uint16(720),
	7545:  uint16(1),
	7546:  uint16(anon_sym_BANG),
	7547:  uint16(722),
	7548:  uint16(1),
	7549:  uint16(anon_sym_LPAREN),
	7550:  uint16(10),
	7551:  uint16(1),
	7552:  uint16(aux_sym_name_repeat1),
	7553:  uint16(39),
	7554:  uint16(1),
	7555:  uint16(sym_expression),
	7556:  uint16(21),
	7557:  uint16(2),
	7558:  uint16(sym_macro_variable),
	7559:  uint16(sym_string),
	7560:  uint16(37),
	7561:  uint16(4),
	7562:  uint16(sym_unary_expression),
	7563:  uint16(sym_binary_expression),
	7564:  uint16(sym_parenthesized_expression),
	7565:  uint16(sym_name),
	7566:  uint16(11),
	7567:  uint16(3),
	7568:  uint16(1),
	7569:  uint16(sym_comment),
	7570:  uint16(459),
	7571:  uint16(1),
	7572:  uint16(anon_sym_DOLLAR_LPAREN),
	7573:  uint16(461),
	7574:  uint16(1),
	7575:  uint16(anon_sym_DQUOTE),
	7576:  uint16(463),
	7577:  uint16(1),
	7578:  uint16(anon_sym_SQUOTE),
	7579:  uint16(706),
	7580:  uint16(1),
	7581:  uint16(sym_symbol),
	7582:  uint16(708),
	7583:  uint16(1),
	7584:  uint16(anon_sym_BANG),
	7585:  uint16(710),
	7586:  uint16(1),
	7587:  uint16(anon_sym_LPAREN),
	7588:  uint16(95),
	7589:  uint16(1),
	7590:  uint16(aux_sym_name_repeat1),
	7591:  uint16(250),
	7592:  uint16(1),
	7593:  uint16(sym_expression),
	7594:  uint16(234),
	7595:  uint16(2),
	7596:  uint16(sym_macro_variable),
	7597:  uint16(sym_string),
	7598:  uint16(117),
	7599:  uint16(4),
	7600:  uint16(sym_unary_expression),
	7601:  uint16(sym_binary_expression),
	7602:  uint16(sym_parenthesized_expression),
	7603:  uint16(sym_name),
	7604:  uint16(3),
	7605:  uint16(3),
	7606:  uint16(1),
	7607:  uint16(sym_comment),
	7608:  uint16(724),
	7609:  uint16(1),
	7611:  uint16(662),
	7612:  uint16(13),
	7613:  uint16(anon_sym_mainmenu),
	7614:  uint16(anon_sym_config),
	7615:  uint16(anon_sym_configdefault),
	7616:  uint16(anon_sym_menuconfig),
	7617:  uint16(anon_sym_choice),
	7618:  uint16(anon_sym_comment),
	7619:  uint16(anon_sym_menu),
	7620:  uint16(anon_sym_if),
	7621:  uint16(anon_sym_source),
	7622:  uint16(anon_sym_rsource),
	7623:  uint16(anon_sym_osource),
	7624:  uint16(anon_sym_orsource),
	7625:  uint16(sym_symbol),
	7626:  uint16(11),
	7627:  uint16(3),
	7628:  uint16(1),
	7629:  uint16(sym_comment),
	7630:  uint16(95),
	7631:  uint16(1),
	7632:  uint16(anon_sym_DOLLAR_LPAREN),
	7633:  uint16(97),
	7634:  uint16(1),
	7635:  uint16(anon_sym_DQUOTE),
	7636:  uint16(99),
	7637:  uint16(1),
	7638:  uint16(anon_sym_SQUOTE),
	7639:  uint16(718),
	7640:  uint16(1),
	7641:  uint16(sym_symbol),
	7642:  uint16(720),
	7643:  uint16(1),
	7644:  uint16(anon_sym_BANG),
	7645:  uint16(722),
	7646:  uint16(1),
	7647:  uint16(anon_sym_LPAREN),
	7648:  uint16(10),
	7649:  uint16(1),
	7650:  uint16(aux_sym_name_repeat1),
	7651:  uint16(41),
	7652:  uint16(1),
	7653:  uint16(sym_expression),
	7654:  uint16(21),
	7655:  uint16(2),
	7656:  uint16(sym_macro_variable),
	7657:  uint16(sym_string),
	7658:  uint16(37),
	7659:  uint16(4),
	7660:  uint16(sym_unary_expression),
	7661:  uint16(sym_binary_expression),
	7662:  uint16(sym_parenthesized_expression),
	7663:  uint16(sym_name),
	7664:  uint16(3),
	7665:  uint16(3),
	7666:  uint16(1),
	7667:  uint16(sym_comment),
	7668:  uint16(726),
	7669:  uint16(1),
	7671:  uint16(660),
	7672:  uint16(13),
	7673:  uint16(anon_sym_mainmenu),
	7674:  uint16(anon_sym_config),
	7675:  uint16(anon_sym_configdefault),
	7676:  uint16(anon_sym_menuconfig),
	7677:  uint16(anon_sym_choice),
	7678:  uint16(anon_sym_comment),
	7679:  uint16(anon_sym_menu),
	7680:  uint16(anon_sym_if),
	7681:  uint16(anon_sym_source),
	7682:  uint16(anon_sym_rsource),
	7683:  uint16(anon_sym_osource),
	7684:  uint16(anon_sym_orsource),
	7685:  uint16(sym_symbol),
	7686:  uint16(11),
	7687:  uint16(3),
	7688:  uint16(1),
	7689:  uint16(sym_comment),
	7690:  uint16(459),
	7691:  uint16(1),
	7692:  uint16(anon_sym_DOLLAR_LPAREN),
	7693:  uint16(461),
	7694:  uint16(1),
	7695:  uint16(anon_sym_DQUOTE),
	7696:  uint16(463),
	7697:  uint16(1),
	7698:  uint16(anon_sym_SQUOTE),
	7699:  uint16(706),
	7700:  uint16(1),
	7701:  uint16(sym_symbol),
	7702:  uint16(708),
	7703:  uint16(1),
	7704:  uint16(anon_sym_BANG),
	7705:  uint16(710),
	7706:  uint16(1),
	7707:  uint16(anon_sym_LPAREN),
	7708:  uint16(95),
	7709:  uint16(1),
	7710:  uint16(aux_sym_name_repeat1),
	7711:  uint16(114),
	7712:  uint16(1),
	7713:  uint16(sym_expression),
	7714:  uint16(234),
	7715:  uint16(2),
	7716:  uint16(sym_macro_variable),
	7717:  uint16(sym_string),
	7718:  uint16(117),
	7719:  uint16(4),
	7720:  uint16(sym_unary_expression),
	7721:  uint16(sym_binary_expression),
	7722:  uint16(sym_parenthesized_expression),
	7723:  uint16(sym_name),
	7724:  uint16(11),
	7725:  uint16(3),
	7726:  uint16(1),
	7727:  uint16(sym_comment),
	7728:  uint16(459),
	7729:  uint16(1),
	7730:  uint16(anon_sym_DOLLAR_LPAREN),
	7731:  uint16(461),
	7732:  uint16(1),
	7733:  uint16(anon_sym_DQUOTE),
	7734:  uint16(463),
	7735:  uint16(1),
	7736:  uint16(anon_sym_SQUOTE),
	7737:  uint16(706),
	7738:  uint16(1),
	7739:  uint16(sym_symbol),
	7740:  uint16(708),
	7741:  uint16(1),
	7742:  uint16(anon_sym_BANG),
	7743:  uint16(710),
	7744:  uint16(1),
	7745:  uint16(anon_sym_LPAREN),
	7746:  uint16(95),
	7747:  uint16(1),
	7748:  uint16(aux_sym_name_repeat1),
	7749:  uint16(249),
	7750:  uint16(1),
	7751:  uint16(sym_expression),
	7752:  uint16(234),
	7753:  uint16(2),
	7754:  uint16(sym_macro_variable),
	7755:  uint16(sym_string),
	7756:  uint16(117),
	7757:  uint16(4),
	7758:  uint16(sym_unary_expression),
	7759:  uint16(sym_binary_expression),
	7760:  uint16(sym_parenthesized_expression),
	7761:  uint16(sym_name),
	7762:  uint16(11),
	7763:  uint16(3),
	7764:  uint16(1),
	7765:  uint16(sym_comment),
	7766:  uint16(459),
	7767:  uint16(1),
	7768:  uint16(anon_sym_DOLLAR_LPAREN),
	7769:  uint16(461),
	7770:  uint16(1),
	7771:  uint16(anon_sym_DQUOTE),
	7772:  uint16(463),
	7773:  uint16(1),
	7774:  uint16(anon_sym_SQUOTE),
	7775:  uint16(706),
	7776:  uint16(1),
	7777:  uint16(sym_symbol),
	7778:  uint16(708),
	7779:  uint16(1),
	7780:  uint16(anon_sym_BANG),
	7781:  uint16(710),
	7782:  uint16(1),
	7783:  uint16(anon_sym_LPAREN),
	7784:  uint16(95),
	7785:  uint16(1),
	7786:  uint16(aux_sym_name_repeat1),
	7787:  uint16(254),
	7788:  uint16(1),
	7789:  uint16(sym_expression),
	7790:  uint16(234),
	7791:  uint16(2),
	7792:  uint16(sym_macro_variable),
	7793:  uint16(sym_string),
	7794:  uint16(117),
	7795:  uint16(4),
	7796:  uint16(sym_unary_expression),
	7797:  uint16(sym_binary_expression),
	7798:  uint16(sym_parenthesized_expression),
	7799:  uint16(sym_name),
	7800:  uint16(11),
	7801:  uint16(3),
	7802:  uint16(1),
	7803:  uint16(sym_comment),
	7804:  uint16(459),
	7805:  uint16(1),
	7806:  uint16(anon_sym_DOLLAR_LPAREN),
	7807:  uint16(461),
	7808:  uint16(1),
	7809:  uint16(anon_sym_DQUOTE),
	7810:  uint16(463),
	7811:  uint16(1),
	7812:  uint16(anon_sym_SQUOTE),
	7813:  uint16(706),
	7814:  uint16(1),
	7815:  uint16(sym_symbol),
	7816:  uint16(708),
	7817:  uint16(1),
	7818:  uint16(anon_sym_BANG),
	7819:  uint16(710),
	7820:  uint16(1),
	7821:  uint16(anon_sym_LPAREN),
	7822:  uint16(95),
	7823:  uint16(1),
	7824:  uint16(aux_sym_name_repeat1),
	7825:  uint16(253),
	7826:  uint16(1),
	7827:  uint16(sym_expression),
	7828:  uint16(234),
	7829:  uint16(2),
	7830:  uint16(sym_macro_variable),
	7831:  uint16(sym_string),
	7832:  uint16(117),
	7833:  uint16(4),
	7834:  uint16(sym_unary_expression),
	7835:  uint16(sym_binary_expression),
	7836:  uint16(sym_parenthesized_expression),
	7837:  uint16(sym_name),
	7838:  uint16(3),
	7839:  uint16(3),
	7840:  uint16(1),
	7841:  uint16(sym_comment),
	7842:  uint16(728),
	7843:  uint16(1),
	7845:  uint16(664),
	7846:  uint16(13),
	7847:  uint16(anon_sym_mainmenu),
	7848:  uint16(anon_sym_config),
	7849:  uint16(anon_sym_configdefault),
	7850:  uint16(anon_sym_menuconfig),
	7851:  uint16(anon_sym_choice),
	7852:  uint16(anon_sym_comment),
	7853:  uint16(anon_sym_menu),
	7854:  uint16(anon_sym_if),
	7855:  uint16(anon_sym_source),
	7856:  uint16(anon_sym_rsource),
	7857:  uint16(anon_sym_osource),
	7858:  uint16(anon_sym_orsource),
	7859:  uint16(sym_symbol),
	7860:  uint16(11),
	7861:  uint16(3),
	7862:  uint16(1),
	7863:  uint16(sym_comment),
	7864:  uint16(459),
	7865:  uint16(1),
	7866:  uint16(anon_sym_DOLLAR_LPAREN),
	7867:  uint16(461),
	7868:  uint16(1),
	7869:  uint16(anon_sym_DQUOTE),
	7870:  uint16(463),
	7871:  uint16(1),
	7872:  uint16(anon_sym_SQUOTE),
	7873:  uint16(710),
	7874:  uint16(1),
	7875:  uint16(anon_sym_LPAREN),
	7876:  uint16(714),
	7877:  uint16(1),
	7878:  uint16(sym_symbol),
	7879:  uint16(716),
	7880:  uint16(1),
	7881:  uint16(anon_sym_BANG),
	7882:  uint16(95),
	7883:  uint16(1),
	7884:  uint16(aux_sym_name_repeat1),
	7885:  uint16(114),
	7886:  uint16(1),
	7887:  uint16(sym_expression),
	7888:  uint16(113),
	7889:  uint16(2),
	7890:  uint16(sym_macro_variable),
	7891:  uint16(sym_string),
	7892:  uint16(117),
	7893:  uint16(4),
	7894:  uint16(sym_unary_expression),
	7895:  uint16(sym_binary_expression),
	7896:  uint16(sym_parenthesized_expression),
	7897:  uint16(sym_name),
	7898:  uint16(3),
	7899:  uint16(3),
	7900:  uint16(1),
	7901:  uint16(sym_comment),
	7902:  uint16(730),
	7903:  uint16(1),
	7905:  uint16(668),
	7906:  uint16(13),
	7907:  uint16(anon_sym_mainmenu),
	7908:  uint16(anon_sym_config),
	7909:  uint16(anon_sym_configdefault),
	7910:  uint16(anon_sym_menuconfig),
	7911:  uint16(anon_sym_choice),
	7912:  uint16(anon_sym_comment),
	7913:  uint16(anon_sym_menu),
	7914:  uint16(anon_sym_if),
	7915:  uint16(anon_sym_source),
	7916:  uint16(anon_sym_rsource),
	7917:  uint16(anon_sym_osource),
	7918:  uint16(anon_sym_orsource),
	7919:  uint16(sym_symbol),
	7920:  uint16(11),
	7921:  uint16(3),
	7922:  uint16(1),
	7923:  uint16(sym_comment),
	7924:  uint16(459),
	7925:  uint16(1),
	7926:  uint16(anon_sym_DOLLAR_LPAREN),
	7927:  uint16(461),
	7928:  uint16(1),
	7929:  uint16(anon_sym_DQUOTE),
	7930:  uint16(463),
	7931:  uint16(1),
	7932:  uint16(anon_sym_SQUOTE),
	7933:  uint16(710),
	7934:  uint16(1),
	7935:  uint16(anon_sym_LPAREN),
	7936:  uint16(714),
	7937:  uint16(1),
	7938:  uint16(sym_symbol),
	7939:  uint16(716),
	7940:  uint16(1),
	7941:  uint16(anon_sym_BANG),
	7942:  uint16(95),
	7943:  uint16(1),
	7944:  uint16(aux_sym_name_repeat1),
	7945:  uint16(118),
	7946:  uint16(1),
	7947:  uint16(sym_expression),
	7948:  uint16(113),
	7949:  uint16(2),
	7950:  uint16(sym_macro_variable),
	7951:  uint16(sym_string),
	7952:  uint16(117),
	7953:  uint16(4),
	7954:  uint16(sym_unary_expression),
	7955:  uint16(sym_binary_expression),
	7956:  uint16(sym_parenthesized_expression),
	7957:  uint16(sym_name),
	7958:  uint16(3),
	7959:  uint16(3),
	7960:  uint16(1),
	7961:  uint16(sym_comment),
	7962:  uint16(732),
	7963:  uint16(1),
	7965:  uint16(670),
	7966:  uint16(13),
	7967:  uint16(anon_sym_mainmenu),
	7968:  uint16(anon_sym_config),
	7969:  uint16(anon_sym_configdefault),
	7970:  uint16(anon_sym_menuconfig),
	7971:  uint16(anon_sym_choice),
	7972:  uint16(anon_sym_comment),
	7973:  uint16(anon_sym_menu),
	7974:  uint16(anon_sym_if),
	7975:  uint16(anon_sym_source),
	7976:  uint16(anon_sym_rsource),
	7977:  uint16(anon_sym_osource),
	7978:  uint16(anon_sym_orsource),
	7979:  uint16(sym_symbol),
	7980:  uint16(3),
	7981:  uint16(3),
	7982:  uint16(1),
	7983:  uint16(sym_comment),
	7984:  uint16(734),
	7985:  uint16(1),
	7987:  uint16(672),
	7988:  uint16(13),
	7989:  uint16(anon_sym_mainmenu),
	7990:  uint16(anon_sym_config),
	7991:  uint16(anon_sym_configdefault),
	7992:  uint16(anon_sym_menuconfig),
	7993:  uint16(anon_sym_choice),
	7994:  uint16(anon_sym_comment),
	7995:  uint16(anon_sym_menu),
	7996:  uint16(anon_sym_if),
	7997:  uint16(anon_sym_source),
	7998:  uint16(anon_sym_rsource),
	7999:  uint16(anon_sym_osource),
	8000:  uint16(anon_sym_orsource),
	8001:  uint16(sym_symbol),
	8002:  uint16(11),
	8003:  uint16(3),
	8004:  uint16(1),
	8005:  uint16(sym_comment),
	8006:  uint16(736),
	8007:  uint16(1),
	8008:  uint16(sym_symbol),
	8009:  uint16(738),
	8010:  uint16(1),
	8011:  uint16(anon_sym_BANG),
	8012:  uint16(740),
	8013:  uint16(1),
	8014:  uint16(anon_sym_LPAREN),
	8015:  uint16(742),
	8016:  uint16(1),
	8017:  uint16(anon_sym_DOLLAR_LPAREN),
	8018:  uint16(744),
	8019:  uint16(1),
	8020:  uint16(anon_sym_DQUOTE),
	8021:  uint16(746),
	8022:  uint16(1),
	8023:  uint16(anon_sym_SQUOTE),
	8024:  uint16(124),
	8025:  uint16(1),
	8026:  uint16(aux_sym_name_repeat1),
	8027:  uint16(156),
	8028:  uint16(1),
	8029:  uint16(sym_expression),
	8030:  uint16(157),
	8031:  uint16(2),
	8032:  uint16(sym_macro_variable),
	8033:  uint16(sym_string),
	8034:  uint16(153),
	8035:  uint16(4),
	8036:  uint16(sym_unary_expression),
	8037:  uint16(sym_binary_expression),
	8038:  uint16(sym_parenthesized_expression),
	8039:  uint16(sym_name),
	8040:  uint16(11),
	8041:  uint16(3),
	8042:  uint16(1),
	8043:  uint16(sym_comment),
	8044:  uint16(736),
	8045:  uint16(1),
	8046:  uint16(sym_symbol),
	8047:  uint16(738),
	8048:  uint16(1),
	8049:  uint16(anon_sym_BANG),
	8050:  uint16(740),
	8051:  uint16(1),
	8052:  uint16(anon_sym_LPAREN),
	8053:  uint16(742),
	8054:  uint16(1),
	8055:  uint16(anon_sym_DOLLAR_LPAREN),
	8056:  uint16(744),
	8057:  uint16(1),
	8058:  uint16(anon_sym_DQUOTE),
	8059:  uint16(746),
	8060:  uint16(1),
	8061:  uint16(anon_sym_SQUOTE),
	8062:  uint16(124),
	8063:  uint16(1),
	8064:  uint16(aux_sym_name_repeat1),
	8065:  uint16(158),
	8066:  uint16(1),
	8067:  uint16(sym_expression),
	8068:  uint16(157),
	8069:  uint16(2),
	8070:  uint16(sym_macro_variable),
	8071:  uint16(sym_string),
	8072:  uint16(153),
	8073:  uint16(4),
	8074:  uint16(sym_unary_expression),
	8075:  uint16(sym_binary_expression),
	8076:  uint16(sym_parenthesized_expression),
	8077:  uint16(sym_name),
	8078:  uint16(11),
	8079:  uint16(3),
	8080:  uint16(1),
	8081:  uint16(sym_comment),
	8082:  uint16(736),
	8083:  uint16(1),
	8084:  uint16(sym_symbol),
	8085:  uint16(738),
	8086:  uint16(1),
	8087:  uint16(anon_sym_BANG),
	8088:  uint16(740),
	8089:  uint16(1),
	8090:  uint16(anon_sym_LPAREN),
	8091:  uint16(742),
	8092:  uint16(1),
	8093:  uint16(anon_sym_DOLLAR_LPAREN),
	8094:  uint16(744),
	8095:  uint16(1),
	8096:  uint16(anon_sym_DQUOTE),
	8097:  uint16(746),
	8098:  uint16(1),
	8099:  uint16(anon_sym_SQUOTE),
	8100:  uint16(124),
	8101:  uint16(1),
	8102:  uint16(aux_sym_name_repeat1),
	8103:  uint16(159),
	8104:  uint16(1),
	8105:  uint16(sym_expression),
	8106:  uint16(157),
	8107:  uint16(2),
	8108:  uint16(sym_macro_variable),
	8109:  uint16(sym_string),
	8110:  uint16(153),
	8111:  uint16(4),
	8112:  uint16(sym_unary_expression),
	8113:  uint16(sym_binary_expression),
	8114:  uint16(sym_parenthesized_expression),
	8115:  uint16(sym_name),
	8116:  uint16(11),
	8117:  uint16(3),
	8118:  uint16(1),
	8119:  uint16(sym_comment),
	8120:  uint16(736),
	8121:  uint16(1),
	8122:  uint16(sym_symbol),
	8123:  uint16(738),
	8124:  uint16(1),
	8125:  uint16(anon_sym_BANG),
	8126:  uint16(740),
	8127:  uint16(1),
	8128:  uint16(anon_sym_LPAREN),
	8129:  uint16(742),
	8130:  uint16(1),
	8131:  uint16(anon_sym_DOLLAR_LPAREN),
	8132:  uint16(744),
	8133:  uint16(1),
	8134:  uint16(anon_sym_DQUOTE),
	8135:  uint16(746),
	8136:  uint16(1),
	8137:  uint16(anon_sym_SQUOTE),
	8138:  uint16(124),
	8139:  uint16(1),
	8140:  uint16(aux_sym_name_repeat1),
	8141:  uint16(160),
	8142:  uint16(1),
	8143:  uint16(sym_expression),
	8144:  uint16(157),
	8145:  uint16(2),
	8146:  uint16(sym_macro_variable),
	8147:  uint16(sym_string),
	8148:  uint16(153),
	8149:  uint16(4),
	8150:  uint16(sym_unary_expression),
	8151:  uint16(sym_binary_expression),
	8152:  uint16(sym_parenthesized_expression),
	8153:  uint16(sym_name),
	8154:  uint16(3),
	8155:  uint16(3),
	8156:  uint16(1),
	8157:  uint16(sym_comment),
	8158:  uint16(200),
	8159:  uint16(1),
	8161:  uint16(198),
	8162:  uint16(13),
	8163:  uint16(anon_sym_mainmenu),
	8164:  uint16(anon_sym_config),
	8165:  uint16(anon_sym_configdefault),
	8166:  uint16(anon_sym_menuconfig),
	8167:  uint16(anon_sym_choice),
	8168:  uint16(anon_sym_comment),
	8169:  uint16(anon_sym_menu),
	8170:  uint16(anon_sym_if),
	8171:  uint16(anon_sym_source),
	8172:  uint16(anon_sym_rsource),
	8173:  uint16(anon_sym_osource),
	8174:  uint16(anon_sym_orsource),
	8175:  uint16(sym_symbol),
	8176:  uint16(11),
	8177:  uint16(3),
	8178:  uint16(1),
	8179:  uint16(sym_comment),
	8180:  uint16(459),
	8181:  uint16(1),
	8182:  uint16(anon_sym_DOLLAR_LPAREN),
	8183:  uint16(461),
	8184:  uint16(1),
	8185:  uint16(anon_sym_DQUOTE),
	8186:  uint16(463),
	8187:  uint16(1),
	8188:  uint16(anon_sym_SQUOTE),
	8189:  uint16(710),
	8190:  uint16(1),
	8191:  uint16(anon_sym_LPAREN),
	8192:  uint16(714),
	8193:  uint16(1),
	8194:  uint16(sym_symbol),
	8195:  uint16(716),
	8196:  uint16(1),
	8197:  uint16(anon_sym_BANG),
	8198:  uint16(95),
	8199:  uint16(1),
	8200:  uint16(aux_sym_name_repeat1),
	8201:  uint16(119),
	8202:  uint16(1),
	8203:  uint16(sym_expression),
	8204:  uint16(113),
	8205:  uint16(2),
	8206:  uint16(sym_macro_variable),
	8207:  uint16(sym_string),
	8208:  uint16(117),
	8209:  uint16(4),
	8210:  uint16(sym_unary_expression),
	8211:  uint16(sym_binary_expression),
	8212:  uint16(sym_parenthesized_expression),
	8213:  uint16(sym_name),
	8214:  uint16(11),
	8215:  uint16(3),
	8216:  uint16(1),
	8217:  uint16(sym_comment),
	8218:  uint16(459),
	8219:  uint16(1),
	8220:  uint16(anon_sym_DOLLAR_LPAREN),
	8221:  uint16(461),
	8222:  uint16(1),
	8223:  uint16(anon_sym_DQUOTE),
	8224:  uint16(463),
	8225:  uint16(1),
	8226:  uint16(anon_sym_SQUOTE),
	8227:  uint16(710),
	8228:  uint16(1),
	8229:  uint16(anon_sym_LPAREN),
	8230:  uint16(714),
	8231:  uint16(1),
	8232:  uint16(sym_symbol),
	8233:  uint16(716),
	8234:  uint16(1),
	8235:  uint16(anon_sym_BANG),
	8236:  uint16(95),
	8237:  uint16(1),
	8238:  uint16(aux_sym_name_repeat1),
	8239:  uint16(120),
	8240:  uint16(1),
	8241:  uint16(sym_expression),
	8242:  uint16(113),
	8243:  uint16(2),
	8244:  uint16(sym_macro_variable),
	8245:  uint16(sym_string),
	8246:  uint16(117),
	8247:  uint16(4),
	8248:  uint16(sym_unary_expression),
	8249:  uint16(sym_binary_expression),
	8250:  uint16(sym_parenthesized_expression),
	8251:  uint16(sym_name),
	8252:  uint16(11),
	8253:  uint16(3),
	8254:  uint16(1),
	8255:  uint16(sym_comment),
	8256:  uint16(694),
	8257:  uint16(1),
	8258:  uint16(sym_symbol),
	8259:  uint16(696),
	8260:  uint16(1),
	8261:  uint16(anon_sym_BANG),
	8262:  uint16(698),
	8263:  uint16(1),
	8264:  uint16(anon_sym_LPAREN),
	8265:  uint16(700),
	8266:  uint16(1),
	8267:  uint16(anon_sym_DOLLAR_LPAREN),
	8268:  uint16(702),
	8269:  uint16(1),
	8270:  uint16(anon_sym_DQUOTE),
	8271:  uint16(704),
	8272:  uint16(1),
	8273:  uint16(anon_sym_SQUOTE),
	8274:  uint16(133),
	8275:  uint16(1),
	8276:  uint16(aux_sym_name_repeat1),
	8277:  uint16(239),
	8278:  uint16(1),
	8279:  uint16(sym_expression),
	8280:  uint16(227),
	8281:  uint16(2),
	8282:  uint16(sym_macro_variable),
	8283:  uint16(sym_string),
	8284:  uint16(246),
	8285:  uint16(4),
	8286:  uint16(sym_unary_expression),
	8287:  uint16(sym_binary_expression),
	8288:  uint16(sym_parenthesized_expression),
	8289:  uint16(sym_name),
	8290:  uint16(3),
	8291:  uint16(200),
	8292:  uint16(1),
	8293:  uint16(aux_sym_type_definition_token1),
	8294:  uint16(568),
	8295:  uint16(1),
	8296:  uint16(sym_comment),
	8297:  uint16(198),
	8298:  uint16(13),
	8299:  uint16(anon_sym_if),
	8300:  uint16(anon_sym_EQ),
	8301:  uint16(anon_sym_PIPE_PIPE),
	8302:  uint16(anon_sym_AMP_AMP),
	8303:  uint16(anon_sym_BANG_EQ),
	8304:  uint16(anon_sym_LT),
	8305:  uint16(anon_sym_GT),
	8306:  uint16(anon_sym_LT_EQ),
	8307:  uint16(anon_sym_GT_EQ),
	8308:  uint16(anon_sym_DOLLAR_LPAREN),
	8309:  uint16(anon_sym_DQUOTE),
	8310:  uint16(anon_sym_SQUOTE),
	8311:  uint16(sym_symbol),
	8312:  uint16(3),
	8313:  uint16(204),
	8314:  uint16(1),
	8315:  uint16(aux_sym_type_definition_token1),
	8316:  uint16(568),
	8317:  uint16(1),
	8318:  uint16(sym_comment),
	8319:  uint16(202),
	8320:  uint16(13),
	8321:  uint16(anon_sym_if),
	8322:  uint16(anon_sym_EQ),
	8323:  uint16(anon_sym_PIPE_PIPE),
	8324:  uint16(anon_sym_AMP_AMP),
	8325:  uint16(anon_sym_BANG_EQ),
	8326:  uint16(anon_sym_LT),
	8327:  uint16(anon_sym_GT),
	8328:  uint16(anon_sym_LT_EQ),
	8329:  uint16(anon_sym_GT_EQ),
	8330:  uint16(anon_sym_DOLLAR_LPAREN),
	8331:  uint16(anon_sym_DQUOTE),
	8332:  uint16(anon_sym_SQUOTE),
	8333:  uint16(sym_symbol),
	8334:  uint16(11),
	8335:  uint16(3),
	8336:  uint16(1),
	8337:  uint16(sym_comment),
	8338:  uint16(176),
	8339:  uint16(1),
	8340:  uint16(anon_sym_DOLLAR_LPAREN),
	8341:  uint16(178),
	8342:  uint16(1),
	8343:  uint16(anon_sym_DQUOTE),
	8344:  uint16(180),
	8345:  uint16(1),
	8346:  uint16(anon_sym_SQUOTE),
	8347:  uint16(748),
	8348:  uint16(1),
	8349:  uint16(sym_symbol),
	8350:  uint16(750),
	8351:  uint16(1),
	8352:  uint16(anon_sym_BANG),
	8353:  uint16(752),
	8354:  uint16(1),
	8355:  uint16(anon_sym_LPAREN),
	8356:  uint16(18),
	8357:  uint16(1),
	8358:  uint16(aux_sym_name_repeat1),
	8359:  uint16(47),
	8360:  uint16(1),
	8361:  uint16(sym_expression),
	8362:  uint16(31),
	8363:  uint16(2),
	8364:  uint16(sym_macro_variable),
	8365:  uint16(sym_string),
	8366:  uint16(44),
	8367:  uint16(4),
	8368:  uint16(sym_unary_expression),
	8369:  uint16(sym_binary_expression),
	8370:  uint16(sym_parenthesized_expression),
	8371:  uint16(sym_name),
	8372:  uint16(11),
	8373:  uint16(3),
	8374:  uint16(1),
	8375:  uint16(sym_comment),
	8376:  uint16(176),
	8377:  uint16(1),
	8378:  uint16(anon_sym_DOLLAR_LPAREN),
	8379:  uint16(178),
	8380:  uint16(1),
	8381:  uint16(anon_sym_DQUOTE),
	8382:  uint16(180),
	8383:  uint16(1),
	8384:  uint16(anon_sym_SQUOTE),
	8385:  uint16(748),
	8386:  uint16(1),
	8387:  uint16(sym_symbol),
	8388:  uint16(750),
	8389:  uint16(1),
	8390:  uint16(anon_sym_BANG),
	8391:  uint16(752),
	8392:  uint16(1),
	8393:  uint16(anon_sym_LPAREN),
	8394:  uint16(18),
	8395:  uint16(1),
	8396:  uint16(aux_sym_name_repeat1),
	8397:  uint16(48),
	8398:  uint16(1),
	8399:  uint16(sym_expression),
	8400:  uint16(31),
	8401:  uint16(2),
	8402:  uint16(sym_macro_variable),
	8403:  uint16(sym_string),
	8404:  uint16(44),
	8405:  uint16(4),
	8406:  uint16(sym_unary_expression),
	8407:  uint16(sym_binary_expression),
	8408:  uint16(sym_parenthesized_expression),
	8409:  uint16(sym_name),
	8410:  uint16(11),
	8411:  uint16(3),
	8412:  uint16(1),
	8413:  uint16(sym_comment),
	8414:  uint16(176),
	8415:  uint16(1),
	8416:  uint16(anon_sym_DOLLAR_LPAREN),
	8417:  uint16(178),
	8418:  uint16(1),
	8419:  uint16(anon_sym_DQUOTE),
	8420:  uint16(180),
	8421:  uint16(1),
	8422:  uint16(anon_sym_SQUOTE),
	8423:  uint16(748),
	8424:  uint16(1),
	8425:  uint16(sym_symbol),
	8426:  uint16(750),
	8427:  uint16(1),
	8428:  uint16(anon_sym_BANG),
	8429:  uint16(752),
	8430:  uint16(1),
	8431:  uint16(anon_sym_LPAREN),
	8432:  uint16(18),
	8433:  uint16(1),
	8434:  uint16(aux_sym_name_repeat1),
	8435:  uint16(49),
	8436:  uint16(1),
	8437:  uint16(sym_expression),
	8438:  uint16(31),
	8439:  uint16(2),
	8440:  uint16(sym_macro_variable),
	8441:  uint16(sym_string),
	8442:  uint16(44),
	8443:  uint16(4),
	8444:  uint16(sym_unary_expression),
	8445:  uint16(sym_binary_expression),
	8446:  uint16(sym_parenthesized_expression),
	8447:  uint16(sym_name),
	8448:  uint16(3),
	8449:  uint16(184),
	8450:  uint16(1),
	8451:  uint16(aux_sym_type_definition_token1),
	8452:  uint16(568),
	8453:  uint16(1),
	8454:  uint16(sym_comment),
	8455:  uint16(182),
	8456:  uint16(13),
	8457:  uint16(anon_sym_if),
	8458:  uint16(anon_sym_EQ),
	8459:  uint16(anon_sym_PIPE_PIPE),
	8460:  uint16(anon_sym_AMP_AMP),
	8461:  uint16(anon_sym_BANG_EQ),
	8462:  uint16(anon_sym_LT),
	8463:  uint16(anon_sym_GT),
	8464:  uint16(anon_sym_LT_EQ),
	8465:  uint16(anon_sym_GT_EQ),
	8466:  uint16(anon_sym_DOLLAR_LPAREN),
	8467:  uint16(anon_sym_DQUOTE),
	8468:  uint16(anon_sym_SQUOTE),
	8469:  uint16(sym_symbol),
	8470:  uint16(3),
	8471:  uint16(188),
	8472:  uint16(1),
	8473:  uint16(aux_sym_type_definition_token1),
	8474:  uint16(568),
	8475:  uint16(1),
	8476:  uint16(sym_comment),
	8477:  uint16(186),
	8478:  uint16(13),
	8479:  uint16(anon_sym_if),
	8480:  uint16(anon_sym_EQ),
	8481:  uint16(anon_sym_PIPE_PIPE),
	8482:  uint16(anon_sym_AMP_AMP),
	8483:  uint16(anon_sym_BANG_EQ),
	8484:  uint16(anon_sym_LT),
	8485:  uint16(anon_sym_GT),
	8486:  uint16(anon_sym_LT_EQ),
	8487:  uint16(anon_sym_GT_EQ),
	8488:  uint16(anon_sym_DOLLAR_LPAREN),
	8489:  uint16(anon_sym_DQUOTE),
	8490:  uint16(anon_sym_SQUOTE),
	8491:  uint16(sym_symbol),
	8492:  uint16(3),
	8493:  uint16(3),
	8494:  uint16(1),
	8495:  uint16(sym_comment),
	8496:  uint16(754),
	8497:  uint16(1),
	8499:  uint16(656),
	8500:  uint16(13),
	8501:  uint16(anon_sym_mainmenu),
	8502:  uint16(anon_sym_config),
	8503:  uint16(anon_sym_configdefault),
	8504:  uint16(anon_sym_menuconfig),
	8505:  uint16(anon_sym_choice),
	8506:  uint16(anon_sym_comment),
	8507:  uint16(anon_sym_menu),
	8508:  uint16(anon_sym_if),
	8509:  uint16(anon_sym_source),
	8510:  uint16(anon_sym_rsource),
	8511:  uint16(anon_sym_osource),
	8512:  uint16(anon_sym_orsource),
	8513:  uint16(sym_symbol),
	8514:  uint16(3),
	8515:  uint16(3),
	8516:  uint16(1),
	8517:  uint16(sym_comment),
	8518:  uint16(756),
	8519:  uint16(1),
	8521:  uint16(674),
	8522:  uint16(13),
	8523:  uint16(anon_sym_mainmenu),
	8524:  uint16(anon_sym_config),
	8525:  uint16(anon_sym_configdefault),
	8526:  uint16(anon_sym_menuconfig),
	8527:  uint16(anon_sym_choice),
	8528:  uint16(anon_sym_comment),
	8529:  uint16(anon_sym_menu),
	8530:  uint16(anon_sym_if),
	8531:  uint16(anon_sym_source),
	8532:  uint16(anon_sym_rsource),
	8533:  uint16(anon_sym_osource),
	8534:  uint16(anon_sym_orsource),
	8535:  uint16(sym_symbol),
	8536:  uint16(3),
	8537:  uint16(3),
	8538:  uint16(1),
	8539:  uint16(sym_comment),
	8540:  uint16(758),
	8541:  uint16(1),
	8543:  uint16(658),
	8544:  uint16(13),
	8545:  uint16(anon_sym_mainmenu),
	8546:  uint16(anon_sym_config),
	8547:  uint16(anon_sym_configdefault),
	8548:  uint16(anon_sym_menuconfig),
	8549:  uint16(anon_sym_choice),
	8550:  uint16(anon_sym_comment),
	8551:  uint16(anon_sym_menu),
	8552:  uint16(anon_sym_if),
	8553:  uint16(anon_sym_source),
	8554:  uint16(anon_sym_rsource),
	8555:  uint16(anon_sym_osource),
	8556:  uint16(anon_sym_orsource),
	8557:  uint16(sym_symbol),
	8558:  uint16(3),
	8559:  uint16(3),
	8560:  uint16(1),
	8561:  uint16(sym_comment),
	8562:  uint16(760),
	8563:  uint16(1),
	8565:  uint16(676),
	8566:  uint16(13),
	8567:  uint16(anon_sym_mainmenu),
	8568:  uint16(anon_sym_config),
	8569:  uint16(anon_sym_configdefault),
	8570:  uint16(anon_sym_menuconfig),
	8571:  uint16(anon_sym_choice),
	8572:  uint16(anon_sym_comment),
	8573:  uint16(anon_sym_menu),
	8574:  uint16(anon_sym_if),
	8575:  uint16(anon_sym_source),
	8576:  uint16(anon_sym_rsource),
	8577:  uint16(anon_sym_osource),
	8578:  uint16(anon_sym_orsource),
	8579:  uint16(sym_symbol),
	8580:  uint16(3),
	8581:  uint16(3),
	8582:  uint16(1),
	8583:  uint16(sym_comment),
	8584:  uint16(762),
	8585:  uint16(1),
	8587:  uint16(678),
	8588:  uint16(13),
	8589:  uint16(anon_sym_mainmenu),
	8590:  uint16(anon_sym_config),
	8591:  uint16(anon_sym_configdefault),
	8592:  uint16(anon_sym_menuconfig),
	8593:  uint16(anon_sym_choice),
	8594:  uint16(anon_sym_comment),
	8595:  uint16(anon_sym_menu),
	8596:  uint16(anon_sym_if),
	8597:  uint16(anon_sym_source),
	8598:  uint16(anon_sym_rsource),
	8599:  uint16(anon_sym_osource),
	8600:  uint16(anon_sym_orsource),
	8601:  uint16(sym_symbol),
	8602:  uint16(11),
	8603:  uint16(3),
	8604:  uint16(1),
	8605:  uint16(sym_comment),
	8606:  uint16(176),
	8607:  uint16(1),
	8608:  uint16(anon_sym_DOLLAR_LPAREN),
	8609:  uint16(178),
	8610:  uint16(1),
	8611:  uint16(anon_sym_DQUOTE),
	8612:  uint16(180),
	8613:  uint16(1),
	8614:  uint16(anon_sym_SQUOTE),
	8615:  uint16(748),
	8616:  uint16(1),
	8617:  uint16(sym_symbol),
	8618:  uint16(750),
	8619:  uint16(1),
	8620:  uint16(anon_sym_BANG),
	8621:  uint16(752),
	8622:  uint16(1),
	8623:  uint16(anon_sym_LPAREN),
	8624:  uint16(18),
	8625:  uint16(1),
	8626:  uint16(aux_sym_name_repeat1),
	8627:  uint16(43),
	8628:  uint16(1),
	8629:  uint16(sym_expression),
	8630:  uint16(31),
	8631:  uint16(2),
	8632:  uint16(sym_macro_variable),
	8633:  uint16(sym_string),
	8634:  uint16(44),
	8635:  uint16(4),
	8636:  uint16(sym_unary_expression),
	8637:  uint16(sym_binary_expression),
	8638:  uint16(sym_parenthesized_expression),
	8639:  uint16(sym_name),
	8640:  uint16(11),
	8641:  uint16(3),
	8642:  uint16(1),
	8643:  uint16(sym_comment),
	8644:  uint16(694),
	8645:  uint16(1),
	8646:  uint16(sym_symbol),
	8647:  uint16(696),
	8648:  uint16(1),
	8649:  uint16(anon_sym_BANG),
	8650:  uint16(698),
	8651:  uint16(1),
	8652:  uint16(anon_sym_LPAREN),
	8653:  uint16(700),
	8654:  uint16(1),
	8655:  uint16(anon_sym_DOLLAR_LPAREN),
	8656:  uint16(702),
	8657:  uint16(1),
	8658:  uint16(anon_sym_DQUOTE),
	8659:  uint16(704),
	8660:  uint16(1),
	8661:  uint16(anon_sym_SQUOTE),
	8662:  uint16(133),
	8663:  uint16(1),
	8664:  uint16(aux_sym_name_repeat1),
	8665:  uint16(244),
	8666:  uint16(1),
	8667:  uint16(sym_expression),
	8668:  uint16(227),
	8669:  uint16(2),
	8670:  uint16(sym_macro_variable),
	8671:  uint16(sym_string),
	8672:  uint16(246),
	8673:  uint16(4),
	8674:  uint16(sym_unary_expression),
	8675:  uint16(sym_binary_expression),
	8676:  uint16(sym_parenthesized_expression),
	8677:  uint16(sym_name),
	8678:  uint16(11),
	8679:  uint16(3),
	8680:  uint16(1),
	8681:  uint16(sym_comment),
	8682:  uint16(694),
	8683:  uint16(1),
	8684:  uint16(sym_symbol),
	8685:  uint16(696),
	8686:  uint16(1),
	8687:  uint16(anon_sym_BANG),
	8688:  uint16(698),
	8689:  uint16(1),
	8690:  uint16(anon_sym_LPAREN),
	8691:  uint16(700),
	8692:  uint16(1),
	8693:  uint16(anon_sym_DOLLAR_LPAREN),
	8694:  uint16(702),
	8695:  uint16(1),
	8696:  uint16(anon_sym_DQUOTE),
	8697:  uint16(704),
	8698:  uint16(1),
	8699:  uint16(anon_sym_SQUOTE),
	8700:  uint16(133),
	8701:  uint16(1),
	8702:  uint16(aux_sym_name_repeat1),
	8703:  uint16(257),
	8704:  uint16(1),
	8705:  uint16(sym_expression),
	8706:  uint16(227),
	8707:  uint16(2),
	8708:  uint16(sym_macro_variable),
	8709:  uint16(sym_string),
	8710:  uint16(246),
	8711:  uint16(4),
	8712:  uint16(sym_unary_expression),
	8713:  uint16(sym_binary_expression),
	8714:  uint16(sym_parenthesized_expression),
	8715:  uint16(sym_name),
	8716:  uint16(3),
	8717:  uint16(3),
	8718:  uint16(1),
	8719:  uint16(sym_comment),
	8720:  uint16(764),
	8721:  uint16(1),
	8723:  uint16(654),
	8724:  uint16(13),
	8725:  uint16(anon_sym_mainmenu),
	8726:  uint16(anon_sym_config),
	8727:  uint16(anon_sym_configdefault),
	8728:  uint16(anon_sym_menuconfig),
	8729:  uint16(anon_sym_choice),
	8730:  uint16(anon_sym_comment),
	8731:  uint16(anon_sym_menu),
	8732:  uint16(anon_sym_if),
	8733:  uint16(anon_sym_source),
	8734:  uint16(anon_sym_rsource),
	8735:  uint16(anon_sym_osource),
	8736:  uint16(anon_sym_orsource),
	8737:  uint16(sym_symbol),
	8738:  uint16(3),
	8739:  uint16(3),
	8740:  uint16(1),
	8741:  uint16(sym_comment),
	8742:  uint16(766),
	8743:  uint16(1),
	8745:  uint16(692),
	8746:  uint16(13),
	8747:  uint16(anon_sym_mainmenu),
	8748:  uint16(anon_sym_config),
	8749:  uint16(anon_sym_configdefault),
	8750:  uint16(anon_sym_menuconfig),
	8751:  uint16(anon_sym_choice),
	8752:  uint16(anon_sym_comment),
	8753:  uint16(anon_sym_menu),
	8754:  uint16(anon_sym_if),
	8755:  uint16(anon_sym_source),
	8756:  uint16(anon_sym_rsource),
	8757:  uint16(anon_sym_osource),
	8758:  uint16(anon_sym_orsource),
	8759:  uint16(sym_symbol),
	8760:  uint16(11),
	8761:  uint16(3),
	8762:  uint16(1),
	8763:  uint16(sym_comment),
	8764:  uint16(95),
	8765:  uint16(1),
	8766:  uint16(anon_sym_DOLLAR_LPAREN),
	8767:  uint16(97),
	8768:  uint16(1),
	8769:  uint16(anon_sym_DQUOTE),
	8770:  uint16(99),
	8771:  uint16(1),
	8772:  uint16(anon_sym_SQUOTE),
	8773:  uint16(718),
	8774:  uint16(1),
	8775:  uint16(sym_symbol),
	8776:  uint16(720),
	8777:  uint16(1),
	8778:  uint16(anon_sym_BANG),
	8779:  uint16(722),
	8780:  uint16(1),
	8781:  uint16(anon_sym_LPAREN),
	8782:  uint16(10),
	8783:  uint16(1),
	8784:  uint16(aux_sym_name_repeat1),
	8785:  uint16(34),
	8786:  uint16(1),
	8787:  uint16(sym_expression),
	8788:  uint16(21),
	8789:  uint16(2),
	8790:  uint16(sym_macro_variable),
	8791:  uint16(sym_string),
	8792:  uint16(37),
	8793:  uint16(4),
	8794:  uint16(sym_unary_expression),
	8795:  uint16(sym_binary_expression),
	8796:  uint16(sym_parenthesized_expression),
	8797:  uint16(sym_name),
	8798:  uint16(11),
	8799:  uint16(3),
	8800:  uint16(1),
	8801:  uint16(sym_comment),
	8802:  uint16(459),
	8803:  uint16(1),
	8804:  uint16(anon_sym_DOLLAR_LPAREN),
	8805:  uint16(461),
	8806:  uint16(1),
	8807:  uint16(anon_sym_DQUOTE),
	8808:  uint16(463),
	8809:  uint16(1),
	8810:  uint16(anon_sym_SQUOTE),
	8811:  uint16(706),
	8812:  uint16(1),
	8813:  uint16(sym_symbol),
	8814:  uint16(708),
	8815:  uint16(1),
	8816:  uint16(anon_sym_BANG),
	8817:  uint16(710),
	8818:  uint16(1),
	8819:  uint16(anon_sym_LPAREN),
	8820:  uint16(95),
	8821:  uint16(1),
	8822:  uint16(aux_sym_name_repeat1),
	8823:  uint16(115),
	8824:  uint16(1),
	8825:  uint16(sym_expression),
	8826:  uint16(234),
	8827:  uint16(2),
	8828:  uint16(sym_macro_variable),
	8829:  uint16(sym_string),
	8830:  uint16(117),
	8831:  uint16(4),
	8832:  uint16(sym_unary_expression),
	8833:  uint16(sym_binary_expression),
	8834:  uint16(sym_parenthesized_expression),
	8835:  uint16(sym_name),
	8836:  uint16(11),
	8837:  uint16(3),
	8838:  uint16(1),
	8839:  uint16(sym_comment),
	8840:  uint16(736),
	8841:  uint16(1),
	8842:  uint16(sym_symbol),
	8843:  uint16(738),
	8844:  uint16(1),
	8845:  uint16(anon_sym_BANG),
	8846:  uint16(740),
	8847:  uint16(1),
	8848:  uint16(anon_sym_LPAREN),
	8849:  uint16(742),
	8850:  uint16(1),
	8851:  uint16(anon_sym_DOLLAR_LPAREN),
	8852:  uint16(744),
	8853:  uint16(1),
	8854:  uint16(anon_sym_DQUOTE),
	8855:  uint16(746),
	8856:  uint16(1),
	8857:  uint16(anon_sym_SQUOTE),
	8858:  uint16(124),
	8859:  uint16(1),
	8860:  uint16(aux_sym_name_repeat1),
	8861:  uint16(154),
	8862:  uint16(1),
	8863:  uint16(sym_expression),
	8864:  uint16(157),
	8865:  uint16(2),
	8866:  uint16(sym_macro_variable),
	8867:  uint16(sym_string),
	8868:  uint16(153),
	8869:  uint16(4),
	8870:  uint16(sym_unary_expression),
	8871:  uint16(sym_binary_expression),
	8872:  uint16(sym_parenthesized_expression),
	8873:  uint16(sym_name),
	8874:  uint16(11),
	8875:  uint16(3),
	8876:  uint16(1),
	8877:  uint16(sym_comment),
	8878:  uint16(176),
	8879:  uint16(1),
	8880:  uint16(anon_sym_DOLLAR_LPAREN),
	8881:  uint16(178),
	8882:  uint16(1),
	8883:  uint16(anon_sym_DQUOTE),
	8884:  uint16(180),
	8885:  uint16(1),
	8886:  uint16(anon_sym_SQUOTE),
	8887:  uint16(748),
	8888:  uint16(1),
	8889:  uint16(sym_symbol),
	8890:  uint16(750),
	8891:  uint16(1),
	8892:  uint16(anon_sym_BANG),
	8893:  uint16(752),
	8894:  uint16(1),
	8895:  uint16(anon_sym_LPAREN),
	8896:  uint16(18),
	8897:  uint16(1),
	8898:  uint16(aux_sym_name_repeat1),
	8899:  uint16(45),
	8900:  uint16(1),
	8901:  uint16(sym_expression),
	8902:  uint16(31),
	8903:  uint16(2),
	8904:  uint16(sym_macro_variable),
	8905:  uint16(sym_string),
	8906:  uint16(44),
	8907:  uint16(4),
	8908:  uint16(sym_unary_expression),
	8909:  uint16(sym_binary_expression),
	8910:  uint16(sym_parenthesized_expression),
	8911:  uint16(sym_name),
	8912:  uint16(11),
	8913:  uint16(3),
	8914:  uint16(1),
	8915:  uint16(sym_comment),
	8916:  uint16(459),
	8917:  uint16(1),
	8918:  uint16(anon_sym_DOLLAR_LPAREN),
	8919:  uint16(461),
	8920:  uint16(1),
	8921:  uint16(anon_sym_DQUOTE),
	8922:  uint16(463),
	8923:  uint16(1),
	8924:  uint16(anon_sym_SQUOTE),
	8925:  uint16(710),
	8926:  uint16(1),
	8927:  uint16(anon_sym_LPAREN),
	8928:  uint16(714),
	8929:  uint16(1),
	8930:  uint16(sym_symbol),
	8931:  uint16(716),
	8932:  uint16(1),
	8933:  uint16(anon_sym_BANG),
	8934:  uint16(85),
	8935:  uint16(1),
	8936:  uint16(sym_expression),
	8937:  uint16(95),
	8938:  uint16(1),
	8939:  uint16(aux_sym_name_repeat1),
	8940:  uint16(113),
	8941:  uint16(2),
	8942:  uint16(sym_macro_variable),
	8943:  uint16(sym_string),
	8944:  uint16(117),
	8945:  uint16(4),
	8946:  uint16(sym_unary_expression),
	8947:  uint16(sym_binary_expression),
	8948:  uint16(sym_parenthesized_expression),
	8949:  uint16(sym_name),
	8950:  uint16(11),
	8951:  uint16(3),
	8952:  uint16(1),
	8953:  uint16(sym_comment),
	8954:  uint16(694),
	8955:  uint16(1),
	8956:  uint16(sym_symbol),
	8957:  uint16(696),
	8958:  uint16(1),
	8959:  uint16(anon_sym_BANG),
	8960:  uint16(698),
	8961:  uint16(1),
	8962:  uint16(anon_sym_LPAREN),
	8963:  uint16(700),
	8964:  uint16(1),
	8965:  uint16(anon_sym_DOLLAR_LPAREN),
	8966:  uint16(702),
	8967:  uint16(1),
	8968:  uint16(anon_sym_DQUOTE),
	8969:  uint16(704),
	8970:  uint16(1),
	8971:  uint16(anon_sym_SQUOTE),
	8972:  uint16(133),
	8973:  uint16(1),
	8974:  uint16(aux_sym_name_repeat1),
	8975:  uint16(240),
	8976:  uint16(1),
	8977:  uint16(sym_expression),
	8978:  uint16(227),
	8979:  uint16(2),
	8980:  uint16(sym_macro_variable),
	8981:  uint16(sym_string),
	8982:  uint16(246),
	8983:  uint16(4),
	8984:  uint16(sym_unary_expression),
	8985:  uint16(sym_binary_expression),
	8986:  uint16(sym_parenthesized_expression),
	8987:  uint16(sym_name),
	8988:  uint16(11),
	8989:  uint16(3),
	8990:  uint16(1),
	8991:  uint16(sym_comment),
	8992:  uint16(694),
	8993:  uint16(1),
	8994:  uint16(sym_symbol),
	8995:  uint16(696),
	8996:  uint16(1),
	8997:  uint16(anon_sym_BANG),
	8998:  uint16(698),
	8999:  uint16(1),
	9000:  uint16(anon_sym_LPAREN),
	9001:  uint16(700),
	9002:  uint16(1),
	9003:  uint16(anon_sym_DOLLAR_LPAREN),
	9004:  uint16(702),
	9005:  uint16(1),
	9006:  uint16(anon_sym_DQUOTE),
	9007:  uint16(704),
	9008:  uint16(1),
	9009:  uint16(anon_sym_SQUOTE),
	9010:  uint16(133),
	9011:  uint16(1),
	9012:  uint16(aux_sym_name_repeat1),
	9013:  uint16(238),
	9014:  uint16(1),
	9015:  uint16(sym_expression),
	9016:  uint16(227),
	9017:  uint16(2),
	9018:  uint16(sym_macro_variable),
	9019:  uint16(sym_string),
	9020:  uint16(246),
	9021:  uint16(4),
	9022:  uint16(sym_unary_expression),
	9023:  uint16(sym_binary_expression),
	9024:  uint16(sym_parenthesized_expression),
	9025:  uint16(sym_name),
	9026:  uint16(11),
	9027:  uint16(3),
	9028:  uint16(1),
	9029:  uint16(sym_comment),
	9030:  uint16(694),
	9031:  uint16(1),
	9032:  uint16(sym_symbol),
	9033:  uint16(696),
	9034:  uint16(1),
	9035:  uint16(anon_sym_BANG),
	9036:  uint16(698),
	9037:  uint16(1),
	9038:  uint16(anon_sym_LPAREN),
	9039:  uint16(700),
	9040:  uint16(1),
	9041:  uint16(anon_sym_DOLLAR_LPAREN),
	9042:  uint16(702),
	9043:  uint16(1),
	9044:  uint16(anon_sym_DQUOTE),
	9045:  uint16(704),
	9046:  uint16(1),
	9047:  uint16(anon_sym_SQUOTE),
	9048:  uint16(133),
	9049:  uint16(1),
	9050:  uint16(aux_sym_name_repeat1),
	9051:  uint16(256),
	9052:  uint16(1),
	9053:  uint16(sym_expression),
	9054:  uint16(227),
	9055:  uint16(2),
	9056:  uint16(sym_macro_variable),
	9057:  uint16(sym_string),
	9058:  uint16(246),
	9059:  uint16(4),
	9060:  uint16(sym_unary_expression),
	9061:  uint16(sym_binary_expression),
	9062:  uint16(sym_parenthesized_expression),
	9063:  uint16(sym_name),
	9064:  uint16(4),
	9065:  uint16(193),
	9066:  uint16(1),
	9067:  uint16(aux_sym_type_definition_token1),
	9068:  uint16(568),
	9069:  uint16(1),
	9070:  uint16(sym_comment),
	9071:  uint16(768),
	9072:  uint16(4),
	9073:  uint16(anon_sym_DOLLAR_LPAREN),
	9074:  uint16(anon_sym_DQUOTE),
	9075:  uint16(anon_sym_SQUOTE),
	9076:  uint16(sym_symbol),
	9077:  uint16(190),
	9078:  uint16(9),
	9079:  uint16(anon_sym_if),
	9080:  uint16(anon_sym_EQ),
	9081:  uint16(anon_sym_PIPE_PIPE),
	9082:  uint16(anon_sym_AMP_AMP),
	9083:  uint16(anon_sym_BANG_EQ),
	9084:  uint16(anon_sym_LT),
	9085:  uint16(anon_sym_GT),
	9086:  uint16(anon_sym_LT_EQ),
	9087:  uint16(anon_sym_GT_EQ),
	9088:  uint16(11),
	9089:  uint16(3),
	9090:  uint16(1),
	9091:  uint16(sym_comment),
	9092:  uint16(459),
	9093:  uint16(1),
	9094:  uint16(anon_sym_DOLLAR_LPAREN),
	9095:  uint16(461),
	9096:  uint16(1),
	9097:  uint16(anon_sym_DQUOTE),
	9098:  uint16(463),
	9099:  uint16(1),
	9100:  uint16(anon_sym_SQUOTE),
	9101:  uint16(706),
	9102:  uint16(1),
	9103:  uint16(sym_symbol),
	9104:  uint16(708),
	9105:  uint16(1),
	9106:  uint16(anon_sym_BANG),
	9107:  uint16(710),
	9108:  uint16(1),
	9109:  uint16(anon_sym_LPAREN),
	9110:  uint16(95),
	9111:  uint16(1),
	9112:  uint16(aux_sym_name_repeat1),
	9113:  uint16(258),
	9114:  uint16(1),
	9115:  uint16(sym_expression),
	9116:  uint16(234),
	9117:  uint16(2),
	9118:  uint16(sym_macro_variable),
	9119:  uint16(sym_string),
	9120:  uint16(117),
	9121:  uint16(4),
	9122:  uint16(sym_unary_expression),
	9123:  uint16(sym_binary_expression),
	9124:  uint16(sym_parenthesized_expression),
	9125:  uint16(sym_name),
	9126:  uint16(11),
	9127:  uint16(3),
	9128:  uint16(1),
	9129:  uint16(sym_comment),
	9130:  uint16(694),
	9131:  uint16(1),
	9132:  uint16(sym_symbol),
	9133:  uint16(696),
	9134:  uint16(1),
	9135:  uint16(anon_sym_BANG),
	9136:  uint16(698),
	9137:  uint16(1),
	9138:  uint16(anon_sym_LPAREN),
	9139:  uint16(700),
	9140:  uint16(1),
	9141:  uint16(anon_sym_DOLLAR_LPAREN),
	9142:  uint16(702),
	9143:  uint16(1),
	9144:  uint16(anon_sym_DQUOTE),
	9145:  uint16(704),
	9146:  uint16(1),
	9147:  uint16(anon_sym_SQUOTE),
	9148:  uint16(133),
	9149:  uint16(1),
	9150:  uint16(aux_sym_name_repeat1),
	9151:  uint16(235),
	9152:  uint16(1),
	9153:  uint16(sym_expression),
	9154:  uint16(227),
	9155:  uint16(2),
	9156:  uint16(sym_macro_variable),
	9157:  uint16(sym_string),
	9158:  uint16(246),
	9159:  uint16(4),
	9160:  uint16(sym_unary_expression),
	9161:  uint16(sym_binary_expression),
	9162:  uint16(sym_parenthesized_expression),
	9163:  uint16(sym_name),
	9164:  uint16(11),
	9165:  uint16(3),
	9166:  uint16(1),
	9167:  uint16(sym_comment),
	9168:  uint16(459),
	9169:  uint16(1),
	9170:  uint16(anon_sym_DOLLAR_LPAREN),
	9171:  uint16(461),
	9172:  uint16(1),
	9173:  uint16(anon_sym_DQUOTE),
	9174:  uint16(463),
	9175:  uint16(1),
	9176:  uint16(anon_sym_SQUOTE),
	9177:  uint16(706),
	9178:  uint16(1),
	9179:  uint16(sym_symbol),
	9180:  uint16(708),
	9181:  uint16(1),
	9182:  uint16(anon_sym_BANG),
	9183:  uint16(710),
	9184:  uint16(1),
	9185:  uint16(anon_sym_LPAREN),
	9186:  uint16(95),
	9187:  uint16(1),
	9188:  uint16(aux_sym_name_repeat1),
	9189:  uint16(248),
	9190:  uint16(1),
	9191:  uint16(sym_expression),
	9192:  uint16(234),
	9193:  uint16(2),
	9194:  uint16(sym_macro_variable),
	9195:  uint16(sym_string),
	9196:  uint16(117),
	9197:  uint16(4),
	9198:  uint16(sym_unary_expression),
	9199:  uint16(sym_binary_expression),
	9200:  uint16(sym_parenthesized_expression),
	9201:  uint16(sym_name),
	9202:  uint16(11),
	9203:  uint16(3),
	9204:  uint16(1),
	9205:  uint16(sym_comment),
	9206:  uint16(694),
	9207:  uint16(1),
	9208:  uint16(sym_symbol),
	9209:  uint16(696),
	9210:  uint16(1),
	9211:  uint16(anon_sym_BANG),
	9212:  uint16(698),
	9213:  uint16(1),
	9214:  uint16(anon_sym_LPAREN),
	9215:  uint16(700),
	9216:  uint16(1),
	9217:  uint16(anon_sym_DOLLAR_LPAREN),
	9218:  uint16(702),
	9219:  uint16(1),
	9220:  uint16(anon_sym_DQUOTE),
	9221:  uint16(704),
	9222:  uint16(1),
	9223:  uint16(anon_sym_SQUOTE),
	9224:  uint16(133),
	9225:  uint16(1),
	9226:  uint16(aux_sym_name_repeat1),
	9227:  uint16(236),
	9228:  uint16(1),
	9229:  uint16(sym_expression),
	9230:  uint16(227),
	9231:  uint16(2),
	9232:  uint16(sym_macro_variable),
	9233:  uint16(sym_string),
	9234:  uint16(246),
	9235:  uint16(4),
	9236:  uint16(sym_unary_expression),
	9237:  uint16(sym_binary_expression),
	9238:  uint16(sym_parenthesized_expression),
	9239:  uint16(sym_name),
	9240:  uint16(11),
	9241:  uint16(3),
	9242:  uint16(1),
	9243:  uint16(sym_comment),
	9244:  uint16(459),
	9245:  uint16(1),
	9246:  uint16(anon_sym_DOLLAR_LPAREN),
	9247:  uint16(461),
	9248:  uint16(1),
	9249:  uint16(anon_sym_DQUOTE),
	9250:  uint16(463),
	9251:  uint16(1),
	9252:  uint16(anon_sym_SQUOTE),
	9253:  uint16(706),
	9254:  uint16(1),
	9255:  uint16(sym_symbol),
	9256:  uint16(708),
	9257:  uint16(1),
	9258:  uint16(anon_sym_BANG),
	9259:  uint16(710),
	9260:  uint16(1),
	9261:  uint16(anon_sym_LPAREN),
	9262:  uint16(95),
	9263:  uint16(1),
	9264:  uint16(aux_sym_name_repeat1),
	9265:  uint16(251),
	9266:  uint16(1),
	9267:  uint16(sym_expression),
	9268:  uint16(234),
	9269:  uint16(2),
	9270:  uint16(sym_macro_variable),
	9271:  uint16(sym_string),
	9272:  uint16(117),
	9273:  uint16(4),
	9274:  uint16(sym_unary_expression),
	9275:  uint16(sym_binary_expression),
	9276:  uint16(sym_parenthesized_expression),
	9277:  uint16(sym_name),
	9278:  uint16(11),
	9279:  uint16(3),
	9280:  uint16(1),
	9281:  uint16(sym_comment),
	9282:  uint16(176),
	9283:  uint16(1),
	9284:  uint16(anon_sym_DOLLAR_LPAREN),
	9285:  uint16(178),
	9286:  uint16(1),
	9287:  uint16(anon_sym_DQUOTE),
	9288:  uint16(180),
	9289:  uint16(1),
	9290:  uint16(anon_sym_SQUOTE),
	9291:  uint16(748),
	9292:  uint16(1),
	9293:  uint16(sym_symbol),
	9294:  uint16(750),
	9295:  uint16(1),
	9296:  uint16(anon_sym_BANG),
	9297:  uint16(752),
	9298:  uint16(1),
	9299:  uint16(anon_sym_LPAREN),
	9300:  uint16(18),
	9301:  uint16(1),
	9302:  uint16(aux_sym_name_repeat1),
	9303:  uint16(46),
	9304:  uint16(1),
	9305:  uint16(sym_expression),
	9306:  uint16(31),
	9307:  uint16(2),
	9308:  uint16(sym_macro_variable),
	9309:  uint16(sym_string),
	9310:  uint16(44),
	9311:  uint16(4),
	9312:  uint16(sym_unary_expression),
	9313:  uint16(sym_binary_expression),
	9314:  uint16(sym_parenthesized_expression),
	9315:  uint16(sym_name),
	9316:  uint16(4),
	9317:  uint16(3),
	9318:  uint16(1),
	9319:  uint16(sym_comment),
	9320:  uint16(190),
	9321:  uint16(2),
	9322:  uint16(anon_sym_LT),
	9323:  uint16(anon_sym_GT),
	9324:  uint16(196),
	9325:  uint16(4),
	9326:  uint16(anon_sym_DOLLAR_LPAREN),
	9327:  uint16(anon_sym_DQUOTE),
	9328:  uint16(anon_sym_SQUOTE),
	9329:  uint16(sym_symbol),
	9330:  uint16(193),
	9331:  uint16(7),
	9332:  uint16(anon_sym_EQ),
	9333:  uint16(anon_sym_PIPE_PIPE),
	9334:  uint16(anon_sym_AMP_AMP),
	9335:  uint16(anon_sym_BANG_EQ),
	9336:  uint16(anon_sym_LT_EQ),
	9337:  uint16(anon_sym_GT_EQ),
	9338:  uint16(anon_sym_RPAREN),
	9339:  uint16(8),
	9340:  uint16(568),
	9341:  uint16(1),
	9342:  uint16(sym_comment),
	9343:  uint16(770),
	9344:  uint16(1),
	9345:  uint16(anon_sym_if),
	9346:  uint16(772),
	9347:  uint16(1),
	9348:  uint16(anon_sym_EQ),
	9349:  uint16(774),
	9350:  uint16(1),
	9351:  uint16(aux_sym_type_definition_token1),
	9352:  uint16(776),
	9353:  uint16(1),
	9354:  uint16(anon_sym_PIPE_PIPE),
	9355:  uint16(778),
	9356:  uint16(1),
	9357:  uint16(anon_sym_AMP_AMP),
	9358:  uint16(379),
	9359:  uint16(1),
	9360:  uint16(sym_conditional_clause),
	9361:  uint16(780),
	9362:  uint16(5),
	9363:  uint16(anon_sym_BANG_EQ),
	9364:  uint16(anon_sym_LT),
	9365:  uint16(anon_sym_GT),
	9366:  uint16(anon_sym_LT_EQ),
	9367:  uint16(anon_sym_GT_EQ),
	9368:  uint16(8),
	9369:  uint16(568),
	9370:  uint16(1),
	9371:  uint16(sym_comment),
	9372:  uint16(770),
	9373:  uint16(1),
	9374:  uint16(anon_sym_if),
	9375:  uint16(772),
	9376:  uint16(1),
	9377:  uint16(anon_sym_EQ),
	9378:  uint16(776),
	9379:  uint16(1),
	9380:  uint16(anon_sym_PIPE_PIPE),
	9381:  uint16(778),
	9382:  uint16(1),
	9383:  uint16(anon_sym_AMP_AMP),
	9384:  uint16(782),
	9385:  uint16(1),
	9386:  uint16(aux_sym_type_definition_token1),
	9387:  uint16(366),
	9388:  uint16(1),
	9389:  uint16(sym_conditional_clause),
	9390:  uint16(780),
	9391:  uint16(5),
	9392:  uint16(anon_sym_BANG_EQ),
	9393:  uint16(anon_sym_LT),
	9394:  uint16(anon_sym_GT),
	9395:  uint16(anon_sym_LT_EQ),
	9396:  uint16(anon_sym_GT_EQ),
	9397:  uint16(8),
	9398:  uint16(568),
	9399:  uint16(1),
	9400:  uint16(sym_comment),
	9401:  uint16(770),
	9402:  uint16(1),
	9403:  uint16(anon_sym_if),
	9404:  uint16(772),
	9405:  uint16(1),
	9406:  uint16(anon_sym_EQ),
	9407:  uint16(776),
	9408:  uint16(1),
	9409:  uint16(anon_sym_PIPE_PIPE),
	9410:  uint16(778),
	9411:  uint16(1),
	9412:  uint16(anon_sym_AMP_AMP),
	9413:  uint16(784),
	9414:  uint16(1),
	9415:  uint16(aux_sym_type_definition_token1),
	9416:  uint16(365),
	9417:  uint16(1),
	9418:  uint16(sym_conditional_clause),
	9419:  uint16(780),
	9420:  uint16(5),
	9421:  uint16(anon_sym_BANG_EQ),
	9422:  uint16(anon_sym_LT),
	9423:  uint16(anon_sym_GT),
	9424:  uint16(anon_sym_LT_EQ),
	9425:  uint16(anon_sym_GT_EQ),
	9426:  uint16(8),
	9427:  uint16(568),
	9428:  uint16(1),
	9429:  uint16(sym_comment),
	9430:  uint16(770),
	9431:  uint16(1),
	9432:  uint16(anon_sym_if),
	9433:  uint16(772),
	9434:  uint16(1),
	9435:  uint16(anon_sym_EQ),
	9436:  uint16(776),
	9437:  uint16(1),
	9438:  uint16(anon_sym_PIPE_PIPE),
	9439:  uint16(778),
	9440:  uint16(1),
	9441:  uint16(anon_sym_AMP_AMP),
	9442:  uint16(786),
	9443:  uint16(1),
	9444:  uint16(aux_sym_type_definition_token1),
	9445:  uint16(375),
	9446:  uint16(1),
	9447:  uint16(sym_conditional_clause),
	9448:  uint16(780),
	9449:  uint16(5),
	9450:  uint16(anon_sym_BANG_EQ),
	9451:  uint16(anon_sym_LT),
	9452:  uint16(anon_sym_GT),
	9453:  uint16(anon_sym_LT_EQ),
	9454:  uint16(anon_sym_GT_EQ),
	9455:  uint16(8),
	9456:  uint16(568),
	9457:  uint16(1),
	9458:  uint16(sym_comment),
	9459:  uint16(770),
	9460:  uint16(1),
	9461:  uint16(anon_sym_if),
	9462:  uint16(772),
	9463:  uint16(1),
	9464:  uint16(anon_sym_EQ),
	9465:  uint16(776),
	9466:  uint16(1),
	9467:  uint16(anon_sym_PIPE_PIPE),
	9468:  uint16(778),
	9469:  uint16(1),
	9470:  uint16(anon_sym_AMP_AMP),
	9471:  uint16(788),
	9472:  uint16(1),
	9473:  uint16(aux_sym_type_definition_token1),
	9474:  uint16(358),
	9475:  uint16(1),
	9476:  uint16(sym_conditional_clause),
	9477:  uint16(780),
	9478:  uint16(5),
	9479:  uint16(anon_sym_BANG_EQ),
	9480:  uint16(anon_sym_LT),
	9481:  uint16(anon_sym_GT),
	9482:  uint16(anon_sym_LT_EQ),
	9483:  uint16(anon_sym_GT_EQ),
	9484:  uint16(8),
	9485:  uint16(568),
	9486:  uint16(1),
	9487:  uint16(sym_comment),
	9488:  uint16(770),
	9489:  uint16(1),
	9490:  uint16(anon_sym_if),
	9491:  uint16(772),
	9492:  uint16(1),
	9493:  uint16(anon_sym_EQ),
	9494:  uint16(776),
	9495:  uint16(1),
	9496:  uint16(anon_sym_PIPE_PIPE),
	9497:  uint16(778),
	9498:  uint16(1),
	9499:  uint16(anon_sym_AMP_AMP),
	9500:  uint16(790),
	9501:  uint16(1),
	9502:  uint16(aux_sym_type_definition_token1),
	9503:  uint16(373),
	9504:  uint16(1),
	9505:  uint16(sym_conditional_clause),
	9506:  uint16(780),
	9507:  uint16(5),
	9508:  uint16(anon_sym_BANG_EQ),
	9509:  uint16(anon_sym_LT),
	9510:  uint16(anon_sym_GT),
	9511:  uint16(anon_sym_LT_EQ),
	9512:  uint16(anon_sym_GT_EQ),
	9513:  uint16(5),
	9514:  uint16(279),
	9515:  uint16(1),
	9516:  uint16(aux_sym_type_definition_token1),
	9517:  uint16(568),
	9518:  uint16(1),
	9519:  uint16(sym_comment),
	9520:  uint16(772),
	9521:  uint16(1),
	9522:  uint16(anon_sym_EQ),
	9523:  uint16(277),
	9524:  uint16(3),
	9525:  uint16(anon_sym_if),
	9526:  uint16(anon_sym_PIPE_PIPE),
	9527:  uint16(anon_sym_AMP_AMP),
	9528:  uint16(780),
	9529:  uint16(5),
	9530:  uint16(anon_sym_BANG_EQ),
	9531:  uint16(anon_sym_LT),
	9532:  uint16(anon_sym_GT),
	9533:  uint16(anon_sym_LT_EQ),
	9534:  uint16(anon_sym_GT_EQ),
	9535:  uint16(4),
	9536:  uint16(279),
	9537:  uint16(1),
	9538:  uint16(aux_sym_type_definition_token1),
	9539:  uint16(568),
	9540:  uint16(1),
	9541:  uint16(sym_comment),
	9542:  uint16(772),
	9543:  uint16(1),
	9544:  uint16(anon_sym_EQ),
	9545:  uint16(277),
	9546:  uint16(8),
	9547:  uint16(anon_sym_if),
	9548:  uint16(anon_sym_PIPE_PIPE),
	9549:  uint16(anon_sym_AMP_AMP),
	9550:  uint16(anon_sym_BANG_EQ),
	9551:  uint16(anon_sym_LT),
	9552:  uint16(anon_sym_GT),
	9553:  uint16(anon_sym_LT_EQ),
	9554:  uint16(anon_sym_GT_EQ),
	9555:  uint16(6),
	9556:  uint16(279),
	9557:  uint16(1),
	9558:  uint16(aux_sym_type_definition_token1),
	9559:  uint16(568),
	9560:  uint16(1),
	9561:  uint16(sym_comment),
	9562:  uint16(772),
	9563:  uint16(1),
	9564:  uint16(anon_sym_EQ),
	9565:  uint16(778),
	9566:  uint16(1),
	9567:  uint16(anon_sym_AMP_AMP),
	9568:  uint16(277),
	9569:  uint16(2),
	9570:  uint16(anon_sym_if),
	9571:  uint16(anon_sym_PIPE_PIPE),
	9572:  uint16(780),
	9573:  uint16(5),
	9574:  uint16(anon_sym_BANG_EQ),
	9575:  uint16(anon_sym_LT),
	9576:  uint16(anon_sym_GT),
	9577:  uint16(anon_sym_LT_EQ),
	9578:  uint16(anon_sym_GT_EQ),
	9579:  uint16(3),
	9580:  uint16(275),
	9581:  uint16(1),
	9582:  uint16(aux_sym_type_definition_token1),
	9583:  uint16(568),
	9584:  uint16(1),
	9585:  uint16(sym_comment),
	9586:  uint16(273),
	9587:  uint16(9),
	9588:  uint16(anon_sym_if),
	9589:  uint16(anon_sym_EQ),
	9590:  uint16(anon_sym_PIPE_PIPE),
	9591:  uint16(anon_sym_AMP_AMP),
	9592:  uint16(anon_sym_BANG_EQ),
	9593:  uint16(anon_sym_LT),
	9594:  uint16(anon_sym_GT),
	9595:  uint16(anon_sym_LT_EQ),
	9596:  uint16(anon_sym_GT_EQ),
	9597:  uint16(3),
	9598:  uint16(293),
	9599:  uint16(1),
	9600:  uint16(aux_sym_type_definition_token1),
	9601:  uint16(568),
	9602:  uint16(1),
	9603:  uint16(sym_comment),
	9604:  uint16(291),
	9605:  uint16(9),
	9606:  uint16(anon_sym_if),
	9607:  uint16(anon_sym_EQ),
	9608:  uint16(anon_sym_PIPE_PIPE),
	9609:  uint16(anon_sym_AMP_AMP),
	9610:  uint16(anon_sym_BANG_EQ),
	9611:  uint16(anon_sym_LT),
	9612:  uint16(anon_sym_GT),
	9613:  uint16(anon_sym_LT_EQ),
	9614:  uint16(anon_sym_GT_EQ),
	9615:  uint16(3),
	9616:  uint16(289),
	9617:  uint16(1),
	9618:  uint16(aux_sym_type_definition_token1),
	9619:  uint16(568),
	9620:  uint16(1),
	9621:  uint16(sym_comment),
	9622:  uint16(287),
	9623:  uint16(9),
	9624:  uint16(anon_sym_if),
	9625:  uint16(anon_sym_EQ),
	9626:  uint16(anon_sym_PIPE_PIPE),
	9627:  uint16(anon_sym_AMP_AMP),
	9628:  uint16(anon_sym_BANG_EQ),
	9629:  uint16(anon_sym_LT),
	9630:  uint16(anon_sym_GT),
	9631:  uint16(anon_sym_LT_EQ),
	9632:  uint16(anon_sym_GT_EQ),
	9633:  uint16(3),
	9634:  uint16(279),
	9635:  uint16(1),
	9636:  uint16(aux_sym_type_definition_token1),
	9637:  uint16(568),
	9638:  uint16(1),
	9639:  uint16(sym_comment),
	9640:  uint16(277),
	9641:  uint16(9),
	9642:  uint16(anon_sym_if),
	9643:  uint16(anon_sym_EQ),
	9644:  uint16(anon_sym_PIPE_PIPE),
	9645:  uint16(anon_sym_AMP_AMP),
	9646:  uint16(anon_sym_BANG_EQ),
	9647:  uint16(anon_sym_LT),
	9648:  uint16(anon_sym_GT),
	9649:  uint16(anon_sym_LT_EQ),
	9650:  uint16(anon_sym_GT_EQ),
	9651:  uint16(7),
	9652:  uint16(3),
	9653:  uint16(1),
	9654:  uint16(sym_comment),
	9655:  uint16(792),
	9656:  uint16(1),
	9657:  uint16(anon_sym_EQ),
	9658:  uint16(794),
	9659:  uint16(1),
	9660:  uint16(anon_sym_PIPE_PIPE),
	9661:  uint16(796),
	9662:  uint16(1),
	9663:  uint16(anon_sym_AMP_AMP),
	9664:  uint16(802),
	9665:  uint16(1),
	9666:  uint16(anon_sym_RPAREN),
	9667:  uint16(800),
	9668:  uint16(2),
	9669:  uint16(anon_sym_LT),
	9670:  uint16(anon_sym_GT),
	9671:  uint16(798),
	9672:  uint16(3),
	9673:  uint16(anon_sym_BANG_EQ),
	9674:  uint16(anon_sym_LT_EQ),
	9675:  uint16(anon_sym_GT_EQ),
	9676:  uint16(6),
	9677:  uint16(3),
	9678:  uint16(1),
	9679:  uint16(sym_comment),
	9680:  uint16(792),
	9681:  uint16(1),
	9682:  uint16(anon_sym_EQ),
	9683:  uint16(796),
	9684:  uint16(1),
	9685:  uint16(anon_sym_AMP_AMP),
	9686:  uint16(279),
	9687:  uint16(2),
	9688:  uint16(anon_sym_PIPE_PIPE),
	9689:  uint16(anon_sym_RPAREN),
	9690:  uint16(800),
	9691:  uint16(2),
	9692:  uint16(anon_sym_LT),
	9693:  uint16(anon_sym_GT),
	9694:  uint16(798),
	9695:  uint16(3),
	9696:  uint16(anon_sym_BANG_EQ),
	9697:  uint16(anon_sym_LT_EQ),
	9698:  uint16(anon_sym_GT_EQ),
	9699:  uint16(7),
	9700:  uint16(3),
	9701:  uint16(1),
	9702:  uint16(sym_comment),
	9703:  uint16(792),
	9704:  uint16(1),
	9705:  uint16(anon_sym_EQ),
	9706:  uint16(794),
	9707:  uint16(1),
	9708:  uint16(anon_sym_PIPE_PIPE),
	9709:  uint16(796),
	9710:  uint16(1),
	9711:  uint16(anon_sym_AMP_AMP),
	9712:  uint16(804),
	9713:  uint16(1),
	9714:  uint16(anon_sym_RPAREN),
	9715:  uint16(800),
	9716:  uint16(2),
	9717:  uint16(anon_sym_LT),
	9718:  uint16(anon_sym_GT),
	9719:  uint16(798),
	9720:  uint16(3),
	9721:  uint16(anon_sym_BANG_EQ),
	9722:  uint16(anon_sym_LT_EQ),
	9723:  uint16(anon_sym_GT_EQ),
	9724:  uint16(7),
	9725:  uint16(3),
	9726:  uint16(1),
	9727:  uint16(sym_comment),
	9728:  uint16(792),
	9729:  uint16(1),
	9730:  uint16(anon_sym_EQ),
	9731:  uint16(794),
	9732:  uint16(1),
	9733:  uint16(anon_sym_PIPE_PIPE),
	9734:  uint16(796),
	9735:  uint16(1),
	9736:  uint16(anon_sym_AMP_AMP),
	9737:  uint16(806),
	9738:  uint16(1),
	9739:  uint16(anon_sym_RPAREN),
	9740:  uint16(800),
	9741:  uint16(2),
	9742:  uint16(anon_sym_LT),
	9743:  uint16(anon_sym_GT),
	9744:  uint16(798),
	9745:  uint16(3),
	9746:  uint16(anon_sym_BANG_EQ),
	9747:  uint16(anon_sym_LT_EQ),
	9748:  uint16(anon_sym_GT_EQ),
	9749:  uint16(7),
	9750:  uint16(3),
	9751:  uint16(1),
	9752:  uint16(sym_comment),
	9753:  uint16(792),
	9754:  uint16(1),
	9755:  uint16(anon_sym_EQ),
	9756:  uint16(794),
	9757:  uint16(1),
	9758:  uint16(anon_sym_PIPE_PIPE),
	9759:  uint16(796),
	9760:  uint16(1),
	9761:  uint16(anon_sym_AMP_AMP),
	9762:  uint16(808),
	9763:  uint16(1),
	9764:  uint16(anon_sym_RPAREN),
	9765:  uint16(800),
	9766:  uint16(2),
	9767:  uint16(anon_sym_LT),
	9768:  uint16(anon_sym_GT),
	9769:  uint16(798),
	9770:  uint16(3),
	9771:  uint16(anon_sym_BANG_EQ),
	9772:  uint16(anon_sym_LT_EQ),
	9773:  uint16(anon_sym_GT_EQ),
	9774:  uint16(4),
	9775:  uint16(3),
	9776:  uint16(1),
	9777:  uint16(sym_comment),
	9778:  uint16(792),
	9779:  uint16(1),
	9780:  uint16(anon_sym_EQ),
	9781:  uint16(277),
	9782:  uint16(2),
	9783:  uint16(anon_sym_LT),
	9784:  uint16(anon_sym_GT),
	9785:  uint16(279),
	9786:  uint16(6),
	9787:  uint16(anon_sym_PIPE_PIPE),
	9788:  uint16(anon_sym_AMP_AMP),
	9789:  uint16(anon_sym_BANG_EQ),
	9790:  uint16(anon_sym_LT_EQ),
	9791:  uint16(anon_sym_GT_EQ),
	9792:  uint16(anon_sym_RPAREN),
	9793:  uint16(5),
	9794:  uint16(3),
	9795:  uint16(1),
	9796:  uint16(sym_comment),
	9797:  uint16(792),
	9798:  uint16(1),
	9799:  uint16(anon_sym_EQ),
	9800:  uint16(800),
	9801:  uint16(2),
	9802:  uint16(anon_sym_LT),
	9803:  uint16(anon_sym_GT),
	9804:  uint16(279),
	9805:  uint16(3),
	9806:  uint16(anon_sym_PIPE_PIPE),
	9807:  uint16(anon_sym_AMP_AMP),
	9808:  uint16(anon_sym_RPAREN),
	9809:  uint16(798),
	9810:  uint16(3),
	9811:  uint16(anon_sym_BANG_EQ),
	9812:  uint16(anon_sym_LT_EQ),
	9813:  uint16(anon_sym_GT_EQ),
	9814:  uint16(6),
	9815:  uint16(568),
	9816:  uint16(1),
	9817:  uint16(sym_comment),
	9818:  uint16(772),
	9819:  uint16(1),
	9820:  uint16(anon_sym_EQ),
	9821:  uint16(776),
	9822:  uint16(1),
	9823:  uint16(anon_sym_PIPE_PIPE),
	9824:  uint16(778),
	9825:  uint16(1),
	9826:  uint16(anon_sym_AMP_AMP),
	9827:  uint16(810),
	9828:  uint16(1),
	9829:  uint16(aux_sym_type_definition_token1),
	9830:  uint16(780),
	9831:  uint16(5),
	9832:  uint16(anon_sym_BANG_EQ),
	9833:  uint16(anon_sym_LT),
	9834:  uint16(anon_sym_GT),
	9835:  uint16(anon_sym_LT_EQ),
	9836:  uint16(anon_sym_GT_EQ),
	9837:  uint16(6),
	9838:  uint16(568),
	9839:  uint16(1),
	9840:  uint16(sym_comment),
	9841:  uint16(772),
	9842:  uint16(1),
	9843:  uint16(anon_sym_EQ),
	9844:  uint16(776),
	9845:  uint16(1),
	9846:  uint16(anon_sym_PIPE_PIPE),
	9847:  uint16(778),
	9848:  uint16(1),
	9849:  uint16(anon_sym_AMP_AMP),
	9850:  uint16(812),
	9851:  uint16(1),
	9852:  uint16(aux_sym_type_definition_token1),
	9853:  uint16(780),
	9854:  uint16(5),
	9855:  uint16(anon_sym_BANG_EQ),
	9856:  uint16(anon_sym_LT),
	9857:  uint16(anon_sym_GT),
	9858:  uint16(anon_sym_LT_EQ),
	9859:  uint16(anon_sym_GT_EQ),
	9860:  uint16(6),
	9861:  uint16(568),
	9862:  uint16(1),
	9863:  uint16(sym_comment),
	9864:  uint16(772),
	9865:  uint16(1),
	9866:  uint16(anon_sym_EQ),
	9867:  uint16(776),
	9868:  uint16(1),
	9869:  uint16(anon_sym_PIPE_PIPE),
	9870:  uint16(778),
	9871:  uint16(1),
	9872:  uint16(anon_sym_AMP_AMP),
	9873:  uint16(814),
	9874:  uint16(1),
	9875:  uint16(aux_sym_type_definition_token1),
	9876:  uint16(780),
	9877:  uint16(5),
	9878:  uint16(anon_sym_BANG_EQ),
	9879:  uint16(anon_sym_LT),
	9880:  uint16(anon_sym_GT),
	9881:  uint16(anon_sym_LT_EQ),
	9882:  uint16(anon_sym_GT_EQ),
	9883:  uint16(7),
	9884:  uint16(3),
	9885:  uint16(1),
	9886:  uint16(sym_comment),
	9887:  uint16(792),
	9888:  uint16(1),
	9889:  uint16(anon_sym_EQ),
	9890:  uint16(794),
	9891:  uint16(1),
	9892:  uint16(anon_sym_PIPE_PIPE),
	9893:  uint16(796),
	9894:  uint16(1),
	9895:  uint16(anon_sym_AMP_AMP),
	9896:  uint16(816),
	9897:  uint16(1),
	9898:  uint16(anon_sym_RPAREN),
	9899:  uint16(800),
	9900:  uint16(2),
	9901:  uint16(anon_sym_LT),
	9902:  uint16(anon_sym_GT),
	9903:  uint16(798),
	9904:  uint16(3),
	9905:  uint16(anon_sym_BANG_EQ),
	9906:  uint16(anon_sym_LT_EQ),
	9907:  uint16(anon_sym_GT_EQ),
	9908:  uint16(8),
	9909:  uint16(568),
	9910:  uint16(1),
	9911:  uint16(sym_comment),
	9912:  uint16(818),
	9913:  uint16(1),
	9914:  uint16(anon_sym_RPAREN),
	9915:  uint16(820),
	9916:  uint16(1),
	9917:  uint16(anon_sym_DOLLAR_LPAREN),
	9918:  uint16(822),
	9919:  uint16(1),
	9920:  uint16(sym_macro_content),
	9921:  uint16(824),
	9922:  uint16(1),
	9923:  uint16(anon_sym_DQUOTE),
	9924:  uint16(826),
	9925:  uint16(1),
	9926:  uint16(anon_sym_SQUOTE),
	9927:  uint16(319),
	9928:  uint16(1),
	9929:  uint16(sym_string),
	9930:  uint16(261),
	9931:  uint16(2),
	9932:  uint16(sym_macro_variable),
	9933:  uint16(aux_sym_macro_variable_repeat1),
	9934:  uint16(8),
	9935:  uint16(568),
	9936:  uint16(1),
	9937:  uint16(sym_comment),
	9938:  uint16(625),
	9939:  uint16(1),
	9940:  uint16(anon_sym_DQUOTE),
	9941:  uint16(627),
	9942:  uint16(1),
	9943:  uint16(anon_sym_SQUOTE),
	9944:  uint16(770),
	9945:  uint16(1),
	9946:  uint16(anon_sym_if),
	9947:  uint16(828),
	9948:  uint16(1),
	9949:  uint16(aux_sym_type_definition_token1),
	9950:  uint16(830),
	9951:  uint16(1),
	9952:  uint16(anon_sym_prompt),
	9953:  uint16(369),
	9954:  uint16(1),
	9955:  uint16(sym_conditional_clause),
	9956:  uint16(337),
	9957:  uint16(2),
	9958:  uint16(sym_input_prompt),
	9959:  uint16(sym_string),
	9960:  uint16(8),
	9961:  uint16(568),
	9962:  uint16(1),
	9963:  uint16(sym_comment),
	9964:  uint16(832),
	9965:  uint16(1),
	9966:  uint16(anon_sym_RPAREN),
	9967:  uint16(834),
	9968:  uint16(1),
	9969:  uint16(anon_sym_DOLLAR_LPAREN),
	9970:  uint16(837),
	9971:  uint16(1),
	9972:  uint16(sym_macro_content),
	9973:  uint16(840),
	9974:  uint16(1),
	9975:  uint16(anon_sym_DQUOTE),
	9976:  uint16(843),
	9977:  uint16(1),
	9978:  uint16(anon_sym_SQUOTE),
	9979:  uint16(319),
	9980:  uint16(1),
	9981:  uint16(sym_string),
	9982:  uint16(261),
	9983:  uint16(2),
	9984:  uint16(sym_macro_variable),
	9985:  uint16(aux_sym_macro_variable_repeat1),
	9986:  uint16(8),
	9987:  uint16(568),
	9988:  uint16(1),
	9989:  uint16(sym_comment),
	9990:  uint16(820),
	9991:  uint16(1),
	9992:  uint16(anon_sym_DOLLAR_LPAREN),
	9993:  uint16(824),
	9994:  uint16(1),
	9995:  uint16(anon_sym_DQUOTE),
	9996:  uint16(826),
	9997:  uint16(1),
	9998:  uint16(anon_sym_SQUOTE),
	9999:  uint16(846),
	10000: uint16(1),
	10001: uint16(anon_sym_RPAREN),
	10002: uint16(848),
	10003: uint16(1),
	10004: uint16(sym_macro_content),
	10005: uint16(319),
	10006: uint16(1),
	10007: uint16(sym_string),
	10008: uint16(268),
	10009: uint16(2),
	10010: uint16(sym_macro_variable),
	10011: uint16(aux_sym_macro_variable_repeat1),
	10012: uint16(8),
	10013: uint16(568),
	10014: uint16(1),
	10015: uint16(sym_comment),
	10016: uint16(820),
	10017: uint16(1),
	10018: uint16(anon_sym_DOLLAR_LPAREN),
	10019: uint16(822),
	10020: uint16(1),
	10021: uint16(sym_macro_content),
	10022: uint16(824),
	10023: uint16(1),
	10024: uint16(anon_sym_DQUOTE),
	10025: uint16(826),
	10026: uint16(1),
	10027: uint16(anon_sym_SQUOTE),
	10028: uint16(850),
	10029: uint16(1),
	10030: uint16(anon_sym_RPAREN),
	10031: uint16(319),
	10032: uint16(1),
	10033: uint16(sym_string),
	10034: uint16(261),
	10035: uint16(2),
	10036: uint16(sym_macro_variable),
	10037: uint16(aux_sym_macro_variable_repeat1),
	10038: uint16(8),
	10039: uint16(568),
	10040: uint16(1),
	10041: uint16(sym_comment),
	10042: uint16(820),
	10043: uint16(1),
	10044: uint16(anon_sym_DOLLAR_LPAREN),
	10045: uint16(824),
	10046: uint16(1),
	10047: uint16(anon_sym_DQUOTE),
	10048: uint16(826),
	10049: uint16(1),
	10050: uint16(anon_sym_SQUOTE),
	10051: uint16(852),
	10052: uint16(1),
	10053: uint16(anon_sym_RPAREN),
	10054: uint16(854),
	10055: uint16(1),
	10056: uint16(sym_macro_content),
	10057: uint16(319),
	10058: uint16(1),
	10059: uint16(sym_string),
	10060: uint16(263),
	10061: uint16(2),
	10062: uint16(sym_macro_variable),
	10063: uint16(aux_sym_macro_variable_repeat1),
	10064: uint16(8),
	10065: uint16(568),
	10066: uint16(1),
	10067: uint16(sym_comment),
	10068: uint16(625),
	10069: uint16(1),
	10070: uint16(anon_sym_DQUOTE),
	10071: uint16(627),
	10072: uint16(1),
	10073: uint16(anon_sym_SQUOTE),
	10074: uint16(770),
	10075: uint16(1),
	10076: uint16(anon_sym_if),
	10077: uint16(830),
	10078: uint16(1),
	10079: uint16(anon_sym_prompt),
	10080: uint16(856),
	10081: uint16(1),
	10082: uint16(aux_sym_type_definition_token1),
	10083: uint16(357),
	10084: uint16(1),
	10085: uint16(sym_conditional_clause),
	10086: uint16(354),
	10087: uint16(2),
	10088: uint16(sym_input_prompt),
	10089: uint16(sym_string),
	10090: uint16(7),
	10091: uint16(3),
	10092: uint16(1),
	10093: uint16(sym_comment),
	10094: uint16(459),
	10095: uint16(1),
	10096: uint16(anon_sym_DOLLAR_LPAREN),
	10097: uint16(461),
	10098: uint16(1),
	10099: uint16(anon_sym_DQUOTE),
	10100: uint16(463),
	10101: uint16(1),
	10102: uint16(anon_sym_SQUOTE),
	10103: uint16(858),
	10104: uint16(1),
	10105: uint16(sym_symbol),
	10106: uint16(343),
	10107: uint16(1),
	10108: uint16(sym_name),
	10109: uint16(95),
	10110: uint16(3),
	10111: uint16(sym_macro_variable),
	10112: uint16(sym_string),
	10113: uint16(aux_sym_name_repeat1),
	10114: uint16(7),
	10115: uint16(3),
	10116: uint16(1),
	10117: uint16(sym_comment),
	10118: uint16(95),
	10119: uint16(1),
	10120: uint16(anon_sym_DOLLAR_LPAREN),
	10121: uint16(97),
	10122: uint16(1),
	10123: uint16(anon_sym_DQUOTE),
	10124: uint16(99),
	10125: uint16(1),
	10126: uint16(anon_sym_SQUOTE),
	10127: uint16(860),
	10128: uint16(1),
	10129: uint16(sym_symbol),
	10130: uint16(93),
	10131: uint16(1),
	10132: uint16(sym_name),
	10133: uint16(10),
	10134: uint16(3),
	10135: uint16(sym_macro_variable),
	10136: uint16(sym_string),
	10137: uint16(aux_sym_name_repeat1),
	10138: uint16(8),
	10139: uint16(568),
	10140: uint16(1),
	10141: uint16(sym_comment),
	10142: uint16(820),
	10143: uint16(1),
	10144: uint16(anon_sym_DOLLAR_LPAREN),
	10145: uint16(822),
	10146: uint16(1),
	10147: uint16(sym_macro_content),
	10148: uint16(824),
	10149: uint16(1),
	10150: uint16(anon_sym_DQUOTE),
	10151: uint16(826),
	10152: uint16(1),
	10153: uint16(anon_sym_SQUOTE),
	10154: uint16(862),
	10155: uint16(1),
	10156: uint16(anon_sym_RPAREN),
	10157: uint16(319),
	10158: uint16(1),
	10159: uint16(sym_string),
	10160: uint16(261),
	10161: uint16(2),
	10162: uint16(sym_macro_variable),
	10163: uint16(aux_sym_macro_variable_repeat1),
	10164: uint16(7),
	10165: uint16(3),
	10166: uint16(1),
	10167: uint16(sym_comment),
	10168: uint16(700),
	10169: uint16(1),
	10170: uint16(anon_sym_DOLLAR_LPAREN),
	10171: uint16(702),
	10172: uint16(1),
	10173: uint16(anon_sym_DQUOTE),
	10174: uint16(704),
	10175: uint16(1),
	10176: uint16(anon_sym_SQUOTE),
	10177: uint16(864),
	10178: uint16(1),
	10179: uint16(sym_symbol),
	10180: uint16(329),
	10181: uint16(1),
	10182: uint16(sym_name),
	10183: uint16(133),
	10184: uint16(3),
	10185: uint16(sym_macro_variable),
	10186: uint16(sym_string),
	10187: uint16(aux_sym_name_repeat1),
	10188: uint16(8),
	10189: uint16(568),
	10190: uint16(1),
	10191: uint16(sym_comment),
	10192: uint16(820),
	10193: uint16(1),
	10194: uint16(anon_sym_DOLLAR_LPAREN),
	10195: uint16(824),
	10196: uint16(1),
	10197: uint16(anon_sym_DQUOTE),
	10198: uint16(826),
	10199: uint16(1),
	10200: uint16(anon_sym_SQUOTE),
	10201: uint16(866),
	10202: uint16(1),
	10203: uint16(anon_sym_RPAREN),
	10204: uint16(868),
	10205: uint16(1),
	10206: uint16(sym_macro_content),
	10207: uint16(319),
	10208: uint16(1),
	10209: uint16(sym_string),
	10210: uint16(271),
	10211: uint16(2),
	10212: uint16(sym_macro_variable),
	10213: uint16(aux_sym_macro_variable_repeat1),
	10214: uint16(8),
	10215: uint16(568),
	10216: uint16(1),
	10217: uint16(sym_comment),
	10218: uint16(820),
	10219: uint16(1),
	10220: uint16(anon_sym_DOLLAR_LPAREN),
	10221: uint16(822),
	10222: uint16(1),
	10223: uint16(sym_macro_content),
	10224: uint16(824),
	10225: uint16(1),
	10226: uint16(anon_sym_DQUOTE),
	10227: uint16(826),
	10228: uint16(1),
	10229: uint16(anon_sym_SQUOTE),
	10230: uint16(870),
	10231: uint16(1),
	10232: uint16(anon_sym_RPAREN),
	10233: uint16(319),
	10234: uint16(1),
	10235: uint16(sym_string),
	10236: uint16(261),
	10237: uint16(2),
	10238: uint16(sym_macro_variable),
	10239: uint16(aux_sym_macro_variable_repeat1),
	10240: uint16(8),
	10241: uint16(568),
	10242: uint16(1),
	10243: uint16(sym_comment),
	10244: uint16(820),
	10245: uint16(1),
	10246: uint16(anon_sym_DOLLAR_LPAREN),
	10247: uint16(824),
	10248: uint16(1),
	10249: uint16(anon_sym_DQUOTE),
	10250: uint16(826),
	10251: uint16(1),
	10252: uint16(anon_sym_SQUOTE),
	10253: uint16(872),
	10254: uint16(1),
	10255: uint16(anon_sym_RPAREN),
	10256: uint16(874),
	10257: uint16(1),
	10258: uint16(sym_macro_content),
	10259: uint16(319),
	10260: uint16(1),
	10261: uint16(sym_string),
	10262: uint16(273),
	10263: uint16(2),
	10264: uint16(sym_macro_variable),
	10265: uint16(aux_sym_macro_variable_repeat1),
	10266: uint16(8),
	10267: uint16(568),
	10268: uint16(1),
	10269: uint16(sym_comment),
	10270: uint16(820),
	10271: uint16(1),
	10272: uint16(anon_sym_DOLLAR_LPAREN),
	10273: uint16(822),
	10274: uint16(1),
	10275: uint16(sym_macro_content),
	10276: uint16(824),
	10277: uint16(1),
	10278: uint16(anon_sym_DQUOTE),
	10279: uint16(826),
	10280: uint16(1),
	10281: uint16(anon_sym_SQUOTE),
	10282: uint16(876),
	10283: uint16(1),
	10284: uint16(anon_sym_RPAREN),
	10285: uint16(319),
	10286: uint16(1),
	10287: uint16(sym_string),
	10288: uint16(261),
	10289: uint16(2),
	10290: uint16(sym_macro_variable),
	10291: uint16(aux_sym_macro_variable_repeat1),
	10292: uint16(8),
	10293: uint16(568),
	10294: uint16(1),
	10295: uint16(sym_comment),
	10296: uint16(820),
	10297: uint16(1),
	10298: uint16(anon_sym_DOLLAR_LPAREN),
	10299: uint16(824),
	10300: uint16(1),
	10301: uint16(anon_sym_DQUOTE),
	10302: uint16(826),
	10303: uint16(1),
	10304: uint16(anon_sym_SQUOTE),
	10305: uint16(878),
	10306: uint16(1),
	10307: uint16(anon_sym_RPAREN),
	10308: uint16(880),
	10309: uint16(1),
	10310: uint16(sym_macro_content),
	10311: uint16(319),
	10312: uint16(1),
	10313: uint16(sym_string),
	10314: uint16(275),
	10315: uint16(2),
	10316: uint16(sym_macro_variable),
	10317: uint16(aux_sym_macro_variable_repeat1),
	10318: uint16(8),
	10319: uint16(568),
	10320: uint16(1),
	10321: uint16(sym_comment),
	10322: uint16(820),
	10323: uint16(1),
	10324: uint16(anon_sym_DOLLAR_LPAREN),
	10325: uint16(822),
	10326: uint16(1),
	10327: uint16(sym_macro_content),
	10328: uint16(824),
	10329: uint16(1),
	10330: uint16(anon_sym_DQUOTE),
	10331: uint16(826),
	10332: uint16(1),
	10333: uint16(anon_sym_SQUOTE),
	10334: uint16(882),
	10335: uint16(1),
	10336: uint16(anon_sym_RPAREN),
	10337: uint16(319),
	10338: uint16(1),
	10339: uint16(sym_string),
	10340: uint16(261),
	10341: uint16(2),
	10342: uint16(sym_macro_variable),
	10343: uint16(aux_sym_macro_variable_repeat1),
	10344: uint16(8),
	10345: uint16(568),
	10346: uint16(1),
	10347: uint16(sym_comment),
	10348: uint16(820),
	10349: uint16(1),
	10350: uint16(anon_sym_DOLLAR_LPAREN),
	10351: uint16(824),
	10352: uint16(1),
	10353: uint16(anon_sym_DQUOTE),
	10354: uint16(826),
	10355: uint16(1),
	10356: uint16(anon_sym_SQUOTE),
	10357: uint16(884),
	10358: uint16(1),
	10359: uint16(anon_sym_RPAREN),
	10360: uint16(886),
	10361: uint16(1),
	10362: uint16(sym_macro_content),
	10363: uint16(319),
	10364: uint16(1),
	10365: uint16(sym_string),
	10366: uint16(277),
	10367: uint16(2),
	10368: uint16(sym_macro_variable),
	10369: uint16(aux_sym_macro_variable_repeat1),
	10370: uint16(8),
	10371: uint16(568),
	10372: uint16(1),
	10373: uint16(sym_comment),
	10374: uint16(820),
	10375: uint16(1),
	10376: uint16(anon_sym_DOLLAR_LPAREN),
	10377: uint16(822),
	10378: uint16(1),
	10379: uint16(sym_macro_content),
	10380: uint16(824),
	10381: uint16(1),
	10382: uint16(anon_sym_DQUOTE),
	10383: uint16(826),
	10384: uint16(1),
	10385: uint16(anon_sym_SQUOTE),
	10386: uint16(888),
	10387: uint16(1),
	10388: uint16(anon_sym_RPAREN),
	10389: uint16(319),
	10390: uint16(1),
	10391: uint16(sym_string),
	10392: uint16(261),
	10393: uint16(2),
	10394: uint16(sym_macro_variable),
	10395: uint16(aux_sym_macro_variable_repeat1),
	10396: uint16(8),
	10397: uint16(568),
	10398: uint16(1),
	10399: uint16(sym_comment),
	10400: uint16(820),
	10401: uint16(1),
	10402: uint16(anon_sym_DOLLAR_LPAREN),
	10403: uint16(824),
	10404: uint16(1),
	10405: uint16(anon_sym_DQUOTE),
	10406: uint16(826),
	10407: uint16(1),
	10408: uint16(anon_sym_SQUOTE),
	10409: uint16(890),
	10410: uint16(1),
	10411: uint16(anon_sym_RPAREN),
	10412: uint16(892),
	10413: uint16(1),
	10414: uint16(sym_macro_content),
	10415: uint16(319),
	10416: uint16(1),
	10417: uint16(sym_string),
	10418: uint16(279),
	10419: uint16(2),
	10420: uint16(sym_macro_variable),
	10421: uint16(aux_sym_macro_variable_repeat1),
	10422: uint16(8),
	10423: uint16(568),
	10424: uint16(1),
	10425: uint16(sym_comment),
	10426: uint16(820),
	10427: uint16(1),
	10428: uint16(anon_sym_DOLLAR_LPAREN),
	10429: uint16(822),
	10430: uint16(1),
	10431: uint16(sym_macro_content),
	10432: uint16(824),
	10433: uint16(1),
	10434: uint16(anon_sym_DQUOTE),
	10435: uint16(826),
	10436: uint16(1),
	10437: uint16(anon_sym_SQUOTE),
	10438: uint16(894),
	10439: uint16(1),
	10440: uint16(anon_sym_RPAREN),
	10441: uint16(319),
	10442: uint16(1),
	10443: uint16(sym_string),
	10444: uint16(261),
	10445: uint16(2),
	10446: uint16(sym_macro_variable),
	10447: uint16(aux_sym_macro_variable_repeat1),
	10448: uint16(8),
	10449: uint16(568),
	10450: uint16(1),
	10451: uint16(sym_comment),
	10452: uint16(820),
	10453: uint16(1),
	10454: uint16(anon_sym_DOLLAR_LPAREN),
	10455: uint16(824),
	10456: uint16(1),
	10457: uint16(anon_sym_DQUOTE),
	10458: uint16(826),
	10459: uint16(1),
	10460: uint16(anon_sym_SQUOTE),
	10461: uint16(896),
	10462: uint16(1),
	10463: uint16(anon_sym_RPAREN),
	10464: uint16(898),
	10465: uint16(1),
	10466: uint16(sym_macro_content),
	10467: uint16(319),
	10468: uint16(1),
	10469: uint16(sym_string),
	10470: uint16(259),
	10471: uint16(2),
	10472: uint16(sym_macro_variable),
	10473: uint16(aux_sym_macro_variable_repeat1),
	10474: uint16(7),
	10475: uint16(3),
	10476: uint16(1),
	10477: uint16(sym_comment),
	10478: uint16(95),
	10479: uint16(1),
	10480: uint16(anon_sym_DOLLAR_LPAREN),
	10481: uint16(97),
	10482: uint16(1),
	10483: uint16(anon_sym_DQUOTE),
	10484: uint16(99),
	10485: uint16(1),
	10486: uint16(anon_sym_SQUOTE),
	10487: uint16(860),
	10488: uint16(1),
	10489: uint16(sym_symbol),
	10490: uint16(90),
	10491: uint16(1),
	10492: uint16(sym_name),
	10493: uint16(10),
	10494: uint16(3),
	10495: uint16(sym_macro_variable),
	10496: uint16(sym_string),
	10497: uint16(aux_sym_name_repeat1),
	10498: uint16(7),
	10499: uint16(3),
	10500: uint16(1),
	10501: uint16(sym_comment),
	10502: uint16(459),
	10503: uint16(1),
	10504: uint16(anon_sym_DOLLAR_LPAREN),
	10505: uint16(461),
	10506: uint16(1),
	10507: uint16(anon_sym_DQUOTE),
	10508: uint16(463),
	10509: uint16(1),
	10510: uint16(anon_sym_SQUOTE),
	10511: uint16(858),
	10512: uint16(1),
	10513: uint16(sym_symbol),
	10514: uint16(331),
	10515: uint16(1),
	10516: uint16(sym_name),
	10517: uint16(95),
	10518: uint16(3),
	10519: uint16(sym_macro_variable),
	10520: uint16(sym_string),
	10521: uint16(aux_sym_name_repeat1),
	10522: uint16(7),
	10523: uint16(3),
	10524: uint16(1),
	10525: uint16(sym_comment),
	10526: uint16(95),
	10527: uint16(1),
	10528: uint16(anon_sym_DOLLAR_LPAREN),
	10529: uint16(97),
	10530: uint16(1),
	10531: uint16(anon_sym_DQUOTE),
	10532: uint16(99),
	10533: uint16(1),
	10534: uint16(anon_sym_SQUOTE),
	10535: uint16(860),
	10536: uint16(1),
	10537: uint16(sym_symbol),
	10538: uint16(92),
	10539: uint16(1),
	10540: uint16(sym_name),
	10541: uint16(10),
	10542: uint16(3),
	10543: uint16(sym_macro_variable),
	10544: uint16(sym_string),
	10545: uint16(aux_sym_name_repeat1),
	10546: uint16(7),
	10547: uint16(3),
	10548: uint16(1),
	10549: uint16(sym_comment),
	10550: uint16(700),
	10551: uint16(1),
	10552: uint16(anon_sym_DOLLAR_LPAREN),
	10553: uint16(702),
	10554: uint16(1),
	10555: uint16(anon_sym_DQUOTE),
	10556: uint16(704),
	10557: uint16(1),
	10558: uint16(anon_sym_SQUOTE),
	10559: uint16(864),
	10560: uint16(1),
	10561: uint16(sym_symbol),
	10562: uint16(334),
	10563: uint16(1),
	10564: uint16(sym_name),
	10565: uint16(133),
	10566: uint16(3),
	10567: uint16(sym_macro_variable),
	10568: uint16(sym_string),
	10569: uint16(aux_sym_name_repeat1),
	10570: uint16(7),
	10571: uint16(3),
	10572: uint16(1),
	10573: uint16(sym_comment),
	10574: uint16(95),
	10575: uint16(1),
	10576: uint16(anon_sym_DOLLAR_LPAREN),
	10577: uint16(97),
	10578: uint16(1),
	10579: uint16(anon_sym_DQUOTE),
	10580: uint16(99),
	10581: uint16(1),
	10582: uint16(anon_sym_SQUOTE),
	10583: uint16(860),
	10584: uint16(1),
	10585: uint16(sym_symbol),
	10586: uint16(91),
	10587: uint16(1),
	10588: uint16(sym_name),
	10589: uint16(10),
	10590: uint16(3),
	10591: uint16(sym_macro_variable),
	10592: uint16(sym_string),
	10593: uint16(aux_sym_name_repeat1),
	10594: uint16(5),
	10595: uint16(568),
	10596: uint16(1),
	10597: uint16(sym_comment),
	10598: uint16(900),
	10599: uint16(1),
	10600: uint16(anon_sym_DOLLAR_LPAREN),
	10601: uint16(903),
	10602: uint16(1),
	10603: uint16(anon_sym_DQUOTE),
	10604: uint16(905),
	10605: uint16(2),
	10606: uint16(aux_sym_string_token1),
	10607: uint16(aux_sym_string_token2),
	10608: uint16(286),
	10609: uint16(2),
	10610: uint16(sym_macro_variable),
	10611: uint16(aux_sym_string_repeat1),
	10612: uint16(5),
	10613: uint16(568),
	10614: uint16(1),
	10615: uint16(sym_comment),
	10616: uint16(908),
	10617: uint16(1),
	10618: uint16(anon_sym_DOLLAR_LPAREN),
	10619: uint16(912),
	10620: uint16(1),
	10621: uint16(anon_sym_SQUOTE),
	10622: uint16(910),
	10623: uint16(2),
	10624: uint16(aux_sym_string_token2),
	10625: uint16(aux_sym_string_token3),
	10626: uint16(290),
	10627: uint16(2),
	10628: uint16(sym_macro_variable),
	10629: uint16(aux_sym_string_repeat2),
	10630: uint16(5),
	10631: uint16(568),
	10632: uint16(1),
	10633: uint16(sym_comment),
	10634: uint16(908),
	10635: uint16(1),
	10636: uint16(anon_sym_DOLLAR_LPAREN),
	10637: uint16(916),
	10638: uint16(1),
	10639: uint16(anon_sym_SQUOTE),
	10640: uint16(914),
	10641: uint16(2),
	10642: uint16(aux_sym_string_token2),
	10643: uint16(aux_sym_string_token3),
	10644: uint16(314),
	10645: uint16(2),
	10646: uint16(sym_macro_variable),
	10647: uint16(aux_sym_string_repeat2),
	10648: uint16(5),
	10649: uint16(568),
	10650: uint16(1),
	10651: uint16(sym_comment),
	10652: uint16(918),
	10653: uint16(1),
	10654: uint16(anon_sym_DOLLAR_LPAREN),
	10655: uint16(920),
	10656: uint16(1),
	10657: uint16(anon_sym_DQUOTE),
	10658: uint16(922),
	10659: uint16(2),
	10660: uint16(aux_sym_string_token1),
	10661: uint16(aux_sym_string_token2),
	10662: uint16(286),
	10663: uint16(2),
	10664: uint16(sym_macro_variable),
	10665: uint16(aux_sym_string_repeat1),
	10666: uint16(5),
	10667: uint16(568),
	10668: uint16(1),
	10669: uint16(sym_comment),
	10670: uint16(908),
	10671: uint16(1),
	10672: uint16(anon_sym_DOLLAR_LPAREN),
	10673: uint16(920),
	10674: uint16(1),
	10675: uint16(anon_sym_SQUOTE),
	10676: uint16(914),
	10677: uint16(2),
	10678: uint16(aux_sym_string_token2),
	10679: uint16(aux_sym_string_token3),
	10680: uint16(314),
	10681: uint16(2),
	10682: uint16(sym_macro_variable),
	10683: uint16(aux_sym_string_repeat2),
	10684: uint16(5),
	10685: uint16(568),
	10686: uint16(1),
	10687: uint16(sym_comment),
	10688: uint16(918),
	10689: uint16(1),
	10690: uint16(anon_sym_DOLLAR_LPAREN),
	10691: uint16(924),
	10692: uint16(1),
	10693: uint16(anon_sym_DQUOTE),
	10694: uint16(926),
	10695: uint16(2),
	10696: uint16(aux_sym_string_token1),
	10697: uint16(aux_sym_string_token2),
	10698: uint16(294),
	10699: uint16(2),
	10700: uint16(sym_macro_variable),
	10701: uint16(aux_sym_string_repeat1),
	10702: uint16(5),
	10703: uint16(568),
	10704: uint16(1),
	10705: uint16(sym_comment),
	10706: uint16(908),
	10707: uint16(1),
	10708: uint16(anon_sym_DOLLAR_LPAREN),
	10709: uint16(924),
	10710: uint16(1),
	10711: uint16(anon_sym_SQUOTE),
	10712: uint16(928),
	10713: uint16(2),
	10714: uint16(aux_sym_string_token2),
	10715: uint16(aux_sym_string_token3),
	10716: uint16(295),
	10717: uint16(2),
	10718: uint16(sym_macro_variable),
	10719: uint16(aux_sym_string_repeat2),
	10720: uint16(5),
	10721: uint16(568),
	10722: uint16(1),
	10723: uint16(sym_comment),
	10724: uint16(918),
	10725: uint16(1),
	10726: uint16(anon_sym_DOLLAR_LPAREN),
	10727: uint16(930),
	10728: uint16(1),
	10729: uint16(anon_sym_DQUOTE),
	10730: uint16(932),
	10731: uint16(2),
	10732: uint16(aux_sym_string_token1),
	10733: uint16(aux_sym_string_token2),
	10734: uint16(312),
	10735: uint16(2),
	10736: uint16(sym_macro_variable),
	10737: uint16(aux_sym_string_repeat1),
	10738: uint16(5),
	10739: uint16(568),
	10740: uint16(1),
	10741: uint16(sym_comment),
	10742: uint16(918),
	10743: uint16(1),
	10744: uint16(anon_sym_DOLLAR_LPAREN),
	10745: uint16(934),
	10746: uint16(1),
	10747: uint16(anon_sym_DQUOTE),
	10748: uint16(922),
	10749: uint16(2),
	10750: uint16(aux_sym_string_token1),
	10751: uint16(aux_sym_string_token2),
	10752: uint16(286),
	10753: uint16(2),
	10754: uint16(sym_macro_variable),
	10755: uint16(aux_sym_string_repeat1),
	10756: uint16(5),
	10757: uint16(568),
	10758: uint16(1),
	10759: uint16(sym_comment),
	10760: uint16(908),
	10761: uint16(1),
	10762: uint16(anon_sym_DOLLAR_LPAREN),
	10763: uint16(934),
	10764: uint16(1),
	10765: uint16(anon_sym_SQUOTE),
	10766: uint16(914),
	10767: uint16(2),
	10768: uint16(aux_sym_string_token2),
	10769: uint16(aux_sym_string_token3),
	10770: uint16(314),
	10771: uint16(2),
	10772: uint16(sym_macro_variable),
	10773: uint16(aux_sym_string_repeat2),
	10774: uint16(5),
	10775: uint16(568),
	10776: uint16(1),
	10777: uint16(sym_comment),
	10778: uint16(918),
	10779: uint16(1),
	10780: uint16(anon_sym_DOLLAR_LPAREN),
	10781: uint16(936),
	10782: uint16(1),
	10783: uint16(anon_sym_DQUOTE),
	10784: uint16(938),
	10785: uint16(2),
	10786: uint16(aux_sym_string_token1),
	10787: uint16(aux_sym_string_token2),
	10788: uint16(308),
	10789: uint16(2),
	10790: uint16(sym_macro_variable),
	10791: uint16(aux_sym_string_repeat1),
	10792: uint16(5),
	10793: uint16(568),
	10794: uint16(1),
	10795: uint16(sym_comment),
	10796: uint16(918),
	10797: uint16(1),
	10798: uint16(anon_sym_DOLLAR_LPAREN),
	10799: uint16(940),
	10800: uint16(1),
	10801: uint16(anon_sym_DQUOTE),
	10802: uint16(942),
	10803: uint16(2),
	10804: uint16(aux_sym_string_token1),
	10805: uint16(aux_sym_string_token2),
	10806: uint16(299),
	10807: uint16(2),
	10808: uint16(sym_macro_variable),
	10809: uint16(aux_sym_string_repeat1),
	10810: uint16(5),
	10811: uint16(568),
	10812: uint16(1),
	10813: uint16(sym_comment),
	10814: uint16(908),
	10815: uint16(1),
	10816: uint16(anon_sym_DOLLAR_LPAREN),
	10817: uint16(940),
	10818: uint16(1),
	10819: uint16(anon_sym_SQUOTE),
	10820: uint16(944),
	10821: uint16(2),
	10822: uint16(aux_sym_string_token2),
	10823: uint16(aux_sym_string_token3),
	10824: uint16(300),
	10825: uint16(2),
	10826: uint16(sym_macro_variable),
	10827: uint16(aux_sym_string_repeat2),
	10828: uint16(5),
	10829: uint16(568),
	10830: uint16(1),
	10831: uint16(sym_comment),
	10832: uint16(918),
	10833: uint16(1),
	10834: uint16(anon_sym_DOLLAR_LPAREN),
	10835: uint16(946),
	10836: uint16(1),
	10837: uint16(anon_sym_DQUOTE),
	10838: uint16(922),
	10839: uint16(2),
	10840: uint16(aux_sym_string_token1),
	10841: uint16(aux_sym_string_token2),
	10842: uint16(286),
	10843: uint16(2),
	10844: uint16(sym_macro_variable),
	10845: uint16(aux_sym_string_repeat1),
	10846: uint16(5),
	10847: uint16(568),
	10848: uint16(1),
	10849: uint16(sym_comment),
	10850: uint16(908),
	10851: uint16(1),
	10852: uint16(anon_sym_DOLLAR_LPAREN),
	10853: uint16(946),
	10854: uint16(1),
	10855: uint16(anon_sym_SQUOTE),
	10856: uint16(914),
	10857: uint16(2),
	10858: uint16(aux_sym_string_token2),
	10859: uint16(aux_sym_string_token3),
	10860: uint16(314),
	10861: uint16(2),
	10862: uint16(sym_macro_variable),
	10863: uint16(aux_sym_string_repeat2),
	10864: uint16(5),
	10865: uint16(568),
	10866: uint16(1),
	10867: uint16(sym_comment),
	10868: uint16(918),
	10869: uint16(1),
	10870: uint16(anon_sym_DOLLAR_LPAREN),
	10871: uint16(948),
	10872: uint16(1),
	10873: uint16(anon_sym_DQUOTE),
	10874: uint16(950),
	10875: uint16(2),
	10876: uint16(aux_sym_string_token1),
	10877: uint16(aux_sym_string_token2),
	10878: uint16(304),
	10879: uint16(2),
	10880: uint16(sym_macro_variable),
	10881: uint16(aux_sym_string_repeat1),
	10882: uint16(5),
	10883: uint16(568),
	10884: uint16(1),
	10885: uint16(sym_comment),
	10886: uint16(908),
	10887: uint16(1),
	10888: uint16(anon_sym_DOLLAR_LPAREN),
	10889: uint16(948),
	10890: uint16(1),
	10891: uint16(anon_sym_SQUOTE),
	10892: uint16(952),
	10893: uint16(2),
	10894: uint16(aux_sym_string_token2),
	10895: uint16(aux_sym_string_token3),
	10896: uint16(305),
	10897: uint16(2),
	10898: uint16(sym_macro_variable),
	10899: uint16(aux_sym_string_repeat2),
	10900: uint16(5),
	10901: uint16(568),
	10902: uint16(1),
	10903: uint16(sym_comment),
	10904: uint16(918),
	10905: uint16(1),
	10906: uint16(anon_sym_DOLLAR_LPAREN),
	10907: uint16(954),
	10908: uint16(1),
	10909: uint16(anon_sym_DQUOTE),
	10910: uint16(956),
	10911: uint16(2),
	10912: uint16(aux_sym_string_token1),
	10913: uint16(aux_sym_string_token2),
	10914: uint16(309),
	10915: uint16(2),
	10916: uint16(sym_macro_variable),
	10917: uint16(aux_sym_string_repeat1),
	10918: uint16(5),
	10919: uint16(568),
	10920: uint16(1),
	10921: uint16(sym_comment),
	10922: uint16(918),
	10923: uint16(1),
	10924: uint16(anon_sym_DOLLAR_LPAREN),
	10925: uint16(958),
	10926: uint16(1),
	10927: uint16(anon_sym_DQUOTE),
	10928: uint16(922),
	10929: uint16(2),
	10930: uint16(aux_sym_string_token1),
	10931: uint16(aux_sym_string_token2),
	10932: uint16(286),
	10933: uint16(2),
	10934: uint16(sym_macro_variable),
	10935: uint16(aux_sym_string_repeat1),
	10936: uint16(5),
	10937: uint16(568),
	10938: uint16(1),
	10939: uint16(sym_comment),
	10940: uint16(908),
	10941: uint16(1),
	10942: uint16(anon_sym_DOLLAR_LPAREN),
	10943: uint16(958),
	10944: uint16(1),
	10945: uint16(anon_sym_SQUOTE),
	10946: uint16(914),
	10947: uint16(2),
	10948: uint16(aux_sym_string_token2),
	10949: uint16(aux_sym_string_token3),
	10950: uint16(314),
	10951: uint16(2),
	10952: uint16(sym_macro_variable),
	10953: uint16(aux_sym_string_repeat2),
	10954: uint16(5),
	10955: uint16(568),
	10956: uint16(1),
	10957: uint16(sym_comment),
	10958: uint16(908),
	10959: uint16(1),
	10960: uint16(anon_sym_DOLLAR_LPAREN),
	10961: uint16(954),
	10962: uint16(1),
	10963: uint16(anon_sym_SQUOTE),
	10964: uint16(960),
	10965: uint16(2),
	10966: uint16(aux_sym_string_token2),
	10967: uint16(aux_sym_string_token3),
	10968: uint16(310),
	10969: uint16(2),
	10970: uint16(sym_macro_variable),
	10971: uint16(aux_sym_string_repeat2),
	10972: uint16(5),
	10973: uint16(568),
	10974: uint16(1),
	10975: uint16(sym_comment),
	10976: uint16(908),
	10977: uint16(1),
	10978: uint16(anon_sym_DOLLAR_LPAREN),
	10979: uint16(930),
	10980: uint16(1),
	10981: uint16(anon_sym_SQUOTE),
	10982: uint16(962),
	10983: uint16(2),
	10984: uint16(aux_sym_string_token2),
	10985: uint16(aux_sym_string_token3),
	10986: uint16(288),
	10987: uint16(2),
	10988: uint16(sym_macro_variable),
	10989: uint16(aux_sym_string_repeat2),
	10990: uint16(5),
	10991: uint16(568),
	10992: uint16(1),
	10993: uint16(sym_comment),
	10994: uint16(918),
	10995: uint16(1),
	10996: uint16(anon_sym_DOLLAR_LPAREN),
	10997: uint16(964),
	10998: uint16(1),
	10999: uint16(anon_sym_DQUOTE),
	11000: uint16(922),
	11001: uint16(2),
	11002: uint16(aux_sym_string_token1),
	11003: uint16(aux_sym_string_token2),
	11004: uint16(286),
	11005: uint16(2),
	11006: uint16(sym_macro_variable),
	11007: uint16(aux_sym_string_repeat1),
	11008: uint16(5),
	11009: uint16(568),
	11010: uint16(1),
	11011: uint16(sym_comment),
	11012: uint16(918),
	11013: uint16(1),
	11014: uint16(anon_sym_DOLLAR_LPAREN),
	11015: uint16(966),
	11016: uint16(1),
	11017: uint16(anon_sym_DQUOTE),
	11018: uint16(922),
	11019: uint16(2),
	11020: uint16(aux_sym_string_token1),
	11021: uint16(aux_sym_string_token2),
	11022: uint16(286),
	11023: uint16(2),
	11024: uint16(sym_macro_variable),
	11025: uint16(aux_sym_string_repeat1),
	11026: uint16(5),
	11027: uint16(568),
	11028: uint16(1),
	11029: uint16(sym_comment),
	11030: uint16(908),
	11031: uint16(1),
	11032: uint16(anon_sym_DOLLAR_LPAREN),
	11033: uint16(966),
	11034: uint16(1),
	11035: uint16(anon_sym_SQUOTE),
	11036: uint16(914),
	11037: uint16(2),
	11038: uint16(aux_sym_string_token2),
	11039: uint16(aux_sym_string_token3),
	11040: uint16(314),
	11041: uint16(2),
	11042: uint16(sym_macro_variable),
	11043: uint16(aux_sym_string_repeat2),
	11044: uint16(5),
	11045: uint16(568),
	11046: uint16(1),
	11047: uint16(sym_comment),
	11048: uint16(908),
	11049: uint16(1),
	11050: uint16(anon_sym_DOLLAR_LPAREN),
	11051: uint16(964),
	11052: uint16(1),
	11053: uint16(anon_sym_SQUOTE),
	11054: uint16(914),
	11055: uint16(2),
	11056: uint16(aux_sym_string_token2),
	11057: uint16(aux_sym_string_token3),
	11058: uint16(314),
	11059: uint16(2),
	11060: uint16(sym_macro_variable),
	11061: uint16(aux_sym_string_repeat2),
	11062: uint16(5),
	11063: uint16(568),
	11064: uint16(1),
	11065: uint16(sym_comment),
	11066: uint16(916),
	11067: uint16(1),
	11068: uint16(anon_sym_DQUOTE),
	11069: uint16(918),
	11070: uint16(1),
	11071: uint16(anon_sym_DOLLAR_LPAREN),
	11072: uint16(922),
	11073: uint16(2),
	11074: uint16(aux_sym_string_token1),
	11075: uint16(aux_sym_string_token2),
	11076: uint16(286),
	11077: uint16(2),
	11078: uint16(sym_macro_variable),
	11079: uint16(aux_sym_string_repeat1),
	11080: uint16(5),
	11081: uint16(568),
	11082: uint16(1),
	11083: uint16(sym_comment),
	11084: uint16(908),
	11085: uint16(1),
	11086: uint16(anon_sym_DOLLAR_LPAREN),
	11087: uint16(936),
	11088: uint16(1),
	11089: uint16(anon_sym_SQUOTE),
	11090: uint16(968),
	11091: uint16(2),
	11092: uint16(aux_sym_string_token2),
	11093: uint16(aux_sym_string_token3),
	11094: uint16(311),
	11095: uint16(2),
	11096: uint16(sym_macro_variable),
	11097: uint16(aux_sym_string_repeat2),
	11098: uint16(5),
	11099: uint16(568),
	11100: uint16(1),
	11101: uint16(sym_comment),
	11102: uint16(970),
	11103: uint16(1),
	11104: uint16(anon_sym_DOLLAR_LPAREN),
	11105: uint16(976),
	11106: uint16(1),
	11107: uint16(anon_sym_SQUOTE),
	11108: uint16(973),
	11109: uint16(2),
	11110: uint16(aux_sym_string_token2),
	11111: uint16(aux_sym_string_token3),
	11112: uint16(314),
	11113: uint16(2),
	11114: uint16(sym_macro_variable),
	11115: uint16(aux_sym_string_repeat2),
	11116: uint16(5),
	11117: uint16(568),
	11118: uint16(1),
	11119: uint16(sym_comment),
	11120: uint16(912),
	11121: uint16(1),
	11122: uint16(anon_sym_DQUOTE),
	11123: uint16(918),
	11124: uint16(1),
	11125: uint16(anon_sym_DOLLAR_LPAREN),
	11126: uint16(978),
	11127: uint16(2),
	11128: uint16(aux_sym_string_token1),
	11129: uint16(aux_sym_string_token2),
	11130: uint16(289),
	11131: uint16(2),
	11132: uint16(sym_macro_variable),
	11133: uint16(aux_sym_string_repeat1),
	11134: uint16(2),
	11135: uint16(568),
	11136: uint16(1),
	11137: uint16(sym_comment),
	11138: uint16(202),
	11139: uint16(5),
	11140: uint16(anon_sym_RPAREN),
	11141: uint16(anon_sym_DOLLAR_LPAREN),
	11142: uint16(sym_macro_content),
	11143: uint16(anon_sym_DQUOTE),
	11144: uint16(anon_sym_SQUOTE),
	11145: uint16(2),
	11146: uint16(568),
	11147: uint16(1),
	11148: uint16(sym_comment),
	11149: uint16(182),
	11150: uint16(5),
	11151: uint16(anon_sym_RPAREN),
	11152: uint16(anon_sym_DOLLAR_LPAREN),
	11153: uint16(sym_macro_content),
	11154: uint16(anon_sym_DQUOTE),
	11155: uint16(anon_sym_SQUOTE),
	11156: uint16(2),
	11157: uint16(568),
	11158: uint16(1),
	11159: uint16(sym_comment),
	11160: uint16(186),
	11161: uint16(5),
	11162: uint16(anon_sym_RPAREN),
	11163: uint16(anon_sym_DOLLAR_LPAREN),
	11164: uint16(sym_macro_content),
	11165: uint16(anon_sym_DQUOTE),
	11166: uint16(anon_sym_SQUOTE),
	11167: uint16(2),
	11168: uint16(568),
	11169: uint16(1),
	11170: uint16(sym_comment),
	11171: uint16(980),
	11172: uint16(5),
	11173: uint16(anon_sym_RPAREN),
	11174: uint16(anon_sym_DOLLAR_LPAREN),
	11175: uint16(sym_macro_content),
	11176: uint16(anon_sym_DQUOTE),
	11177: uint16(anon_sym_SQUOTE),
	11178: uint16(2),
	11179: uint16(568),
	11180: uint16(1),
	11181: uint16(sym_comment),
	11182: uint16(198),
	11183: uint16(5),
	11184: uint16(anon_sym_RPAREN),
	11185: uint16(anon_sym_DOLLAR_LPAREN),
	11186: uint16(sym_macro_content),
	11187: uint16(anon_sym_DQUOTE),
	11188: uint16(anon_sym_SQUOTE),
	11189: uint16(2),
	11190: uint16(3),
	11191: uint16(1),
	11192: uint16(sym_comment),
	11193: uint16(982),
	11194: uint16(4),
	11195: uint16(anon_sym_EQ),
	11196: uint16(anon_sym_COLON_EQ),
	11197: uint16(anon_sym_PLUS_EQ),
	11198: uint16(anon_sym_QMARK_EQ),
	11199: uint16(2),
	11200: uint16(3),
	11201: uint16(1),
	11202: uint16(sym_comment),
	11203: uint16(984),
	11204: uint16(4),
	11205: uint16(anon_sym_EQ),
	11206: uint16(anon_sym_COLON_EQ),
	11207: uint16(anon_sym_PLUS_EQ),
	11208: uint16(anon_sym_QMARK_EQ),
	11209: uint16(2),
	11210: uint16(568),
	11211: uint16(1),
	11212: uint16(sym_comment),
	11213: uint16(202),
	11214: uint16(4),
	11215: uint16(anon_sym_DOLLAR_LPAREN),
	11216: uint16(anon_sym_DQUOTE),
	11217: uint16(aux_sym_string_token1),
	11218: uint16(aux_sym_string_token2),
	11219: uint16(2),
	11220: uint16(568),
	11221: uint16(1),
	11222: uint16(sym_comment),
	11223: uint16(186),
	11224: uint16(4),
	11225: uint16(anon_sym_DOLLAR_LPAREN),
	11226: uint16(aux_sym_string_token2),
	11227: uint16(anon_sym_SQUOTE),
	11228: uint16(aux_sym_string_token3),
	11229: uint16(2),
	11230: uint16(568),
	11231: uint16(1),
	11232: uint16(sym_comment),
	11233: uint16(202),
	11234: uint16(4),
	11235: uint16(anon_sym_DOLLAR_LPAREN),
	11236: uint16(aux_sym_string_token2),
	11237: uint16(anon_sym_SQUOTE),
	11238: uint16(aux_sym_string_token3),
	11239: uint16(2),
	11240: uint16(568),
	11241: uint16(1),
	11242: uint16(sym_comment),
	11243: uint16(186),
	11244: uint16(4),
	11245: uint16(anon_sym_DOLLAR_LPAREN),
	11246: uint16(anon_sym_DQUOTE),
	11247: uint16(aux_sym_string_token1),
	11248: uint16(aux_sym_string_token2),
	11249: uint16(4),
	11250: uint16(568),
	11251: uint16(1),
	11252: uint16(sym_comment),
	11253: uint16(770),
	11254: uint16(1),
	11255: uint16(anon_sym_if),
	11256: uint16(986),
	11257: uint16(1),
	11258: uint16(aux_sym_type_definition_token1),
	11259: uint16(361),
	11260: uint16(1),
	11261: uint16(sym_conditional_clause),
	11262: uint16(4),
	11263: uint16(3),
	11264: uint16(1),
	11265: uint16(sym_comment),
	11266: uint16(459),
	11267: uint16(1),
	11268: uint16(anon_sym_DOLLAR_LPAREN),
	11269: uint16(988),
	11270: uint16(1),
	11271: uint16(sym_symbol),
	11272: uint16(350),
	11273: uint16(1),
	11274: uint16(sym_macro_variable),
	11275: uint16(4),
	11276: uint16(568),
	11277: uint16(1),
	11278: uint16(sym_comment),
	11279: uint16(770),
	11280: uint16(1),
	11281: uint16(anon_sym_if),
	11282: uint16(990),
	11283: uint16(1),
	11284: uint16(aux_sym_type_definition_token1),
	11285: uint16(378),
	11286: uint16(1),
	11287: uint16(sym_conditional_clause),
	11288: uint16(4),
	11289: uint16(568),
	11290: uint16(1),
	11291: uint16(sym_comment),
	11292: uint16(770),
	11293: uint16(1),
	11294: uint16(anon_sym_if),
	11295: uint16(992),
	11296: uint16(1),
	11297: uint16(aux_sym_type_definition_token1),
	11298: uint16(371),
	11299: uint16(1),
	11300: uint16(sym_conditional_clause),
	11301: uint16(3),
	11302: uint16(3),
	11303: uint16(1),
	11304: uint16(sym_comment),
	11305: uint16(994),
	11306: uint16(1),
	11307: uint16(anon_sym_default),
	11308: uint16(121),
	11309: uint16(2),
	11310: uint16(sym_default_value),
	11311: uint16(aux_sym_configdefault_repeat1),
	11312: uint16(4),
	11313: uint16(3),
	11314: uint16(1),
	11315: uint16(sym_comment),
	11316: uint16(178),
	11317: uint16(1),
	11318: uint16(anon_sym_DQUOTE),
	11319: uint16(180),
	11320: uint16(1),
	11321: uint16(anon_sym_SQUOTE),
	11322: uint16(28),
	11323: uint16(1),
	11324: uint16(sym_string),
	11325: uint16(4),
	11326: uint16(3),
	11327: uint16(1),
	11328: uint16(sym_comment),
	11329: uint16(461),
	11330: uint16(1),
	11331: uint16(anon_sym_DQUOTE),
	11332: uint16(463),
	11333: uint16(1),
	11334: uint16(anon_sym_SQUOTE),
	11335: uint16(137),
	11336: uint16(1),
	11337: uint16(sym_string),
	11338: uint16(4),
	11339: uint16(568),
	11340: uint16(1),
	11341: uint16(sym_comment),
	11342: uint16(770),
	11343: uint16(1),
	11344: uint16(anon_sym_if),
	11345: uint16(996),
	11346: uint16(1),
	11347: uint16(aux_sym_type_definition_token1),
	11348: uint16(376),
	11349: uint16(1),
	11350: uint16(sym_conditional_clause),
	11351: uint16(4),
	11352: uint16(568),
	11353: uint16(1),
	11354: uint16(sym_comment),
	11355: uint16(770),
	11356: uint16(1),
	11357: uint16(anon_sym_if),
	11358: uint16(998),
	11359: uint16(1),
	11360: uint16(aux_sym_type_definition_token1),
	11361: uint16(380),
	11362: uint16(1),
	11363: uint16(sym_conditional_clause),
	11364: uint16(4),
	11365: uint16(3),
	11366: uint16(1),
	11367: uint16(sym_comment),
	11368: uint16(1000),
	11369: uint16(1),
	11370: uint16(anon_sym_DQUOTE),
	11371: uint16(1002),
	11372: uint16(1),
	11373: uint16(anon_sym_SQUOTE),
	11374: uint16(209),
	11375: uint16(1),
	11376: uint16(sym_string),
	11377: uint16(4),
	11378: uint16(568),
	11379: uint16(1),
	11380: uint16(sym_comment),
	11381: uint16(770),
	11382: uint16(1),
	11383: uint16(anon_sym_if),
	11384: uint16(1004),
	11385: uint16(1),
	11386: uint16(aux_sym_type_definition_token1),
	11387: uint16(382),
	11388: uint16(1),
	11389: uint16(sym_conditional_clause),
	11390: uint16(4),
	11391: uint16(3),
	11392: uint16(1),
	11393: uint16(sym_comment),
	11394: uint16(97),
	11395: uint16(1),
	11396: uint16(anon_sym_DQUOTE),
	11397: uint16(99),
	11398: uint16(1),
	11399: uint16(anon_sym_SQUOTE),
	11400: uint16(13),
	11401: uint16(1),
	11402: uint16(sym_string),
	11403: uint16(4),
	11404: uint16(568),
	11405: uint16(1),
	11406: uint16(sym_comment),
	11407: uint16(770),
	11408: uint16(1),
	11409: uint16(anon_sym_if),
	11410: uint16(1006),
	11411: uint16(1),
	11412: uint16(aux_sym_type_definition_token1),
	11413: uint16(359),
	11414: uint16(1),
	11415: uint16(sym_conditional_clause),
	11416: uint16(4),
	11417: uint16(568),
	11418: uint16(1),
	11419: uint16(sym_comment),
	11420: uint16(770),
	11421: uint16(1),
	11422: uint16(anon_sym_if),
	11423: uint16(1008),
	11424: uint16(1),
	11425: uint16(aux_sym_type_definition_token1),
	11426: uint16(364),
	11427: uint16(1),
	11428: uint16(sym_conditional_clause),
	11429: uint16(4),
	11430: uint16(3),
	11431: uint16(1),
	11432: uint16(sym_comment),
	11433: uint16(461),
	11434: uint16(1),
	11435: uint16(anon_sym_DQUOTE),
	11436: uint16(463),
	11437: uint16(1),
	11438: uint16(anon_sym_SQUOTE),
	11439: uint16(139),
	11440: uint16(1),
	11441: uint16(sym_string),
	11442: uint16(4),
	11443: uint16(3),
	11444: uint16(1),
	11445: uint16(sym_comment),
	11446: uint16(459),
	11447: uint16(1),
	11448: uint16(anon_sym_DOLLAR_LPAREN),
	11449: uint16(1010),
	11450: uint16(1),
	11451: uint16(sym_symbol),
	11452: uint16(349),
	11453: uint16(1),
	11454: uint16(sym_macro_variable),
	11455: uint16(3),
	11456: uint16(3),
	11457: uint16(1),
	11458: uint16(sym_comment),
	11459: uint16(1012),
	11460: uint16(1),
	11461: uint16(anon_sym_default),
	11462: uint16(128),
	11463: uint16(2),
	11464: uint16(sym_default_value),
	11465: uint16(aux_sym_configdefault_repeat1),
	11466: uint16(4),
	11467: uint16(3),
	11468: uint16(1),
	11469: uint16(sym_comment),
	11470: uint16(97),
	11471: uint16(1),
	11472: uint16(anon_sym_DQUOTE),
	11473: uint16(99),
	11474: uint16(1),
	11475: uint16(anon_sym_SQUOTE),
	11476: uint16(6),
	11477: uint16(1),
	11478: uint16(sym_string),
	11479: uint16(4),
	11480: uint16(3),
	11481: uint16(1),
	11482: uint16(sym_comment),
	11483: uint16(702),
	11484: uint16(1),
	11485: uint16(anon_sym_DQUOTE),
	11486: uint16(704),
	11487: uint16(1),
	11488: uint16(anon_sym_SQUOTE),
	11489: uint16(340),
	11490: uint16(1),
	11491: uint16(sym_string),
	11492: uint16(4),
	11493: uint16(3),
	11494: uint16(1),
	11495: uint16(sym_comment),
	11496: uint16(702),
	11497: uint16(1),
	11498: uint16(anon_sym_DQUOTE),
	11499: uint16(704),
	11500: uint16(1),
	11501: uint16(anon_sym_SQUOTE),
	11502: uint16(330),
	11503: uint16(1),
	11504: uint16(sym_string),
	11505: uint16(4),
	11506: uint16(568),
	11507: uint16(1),
	11508: uint16(sym_comment),
	11509: uint16(770),
	11510: uint16(1),
	11511: uint16(anon_sym_if),
	11512: uint16(1014),
	11513: uint16(1),
	11514: uint16(aux_sym_type_definition_token1),
	11515: uint16(362),
	11516: uint16(1),
	11517: uint16(sym_conditional_clause),
	11518: uint16(4),
	11519: uint16(568),
	11520: uint16(1),
	11521: uint16(sym_comment),
	11522: uint16(770),
	11523: uint16(1),
	11524: uint16(anon_sym_if),
	11525: uint16(1016),
	11526: uint16(1),
	11527: uint16(aux_sym_type_definition_token1),
	11528: uint16(377),
	11529: uint16(1),
	11530: uint16(sym_conditional_clause),
	11531: uint16(4),
	11532: uint16(3),
	11533: uint16(1),
	11534: uint16(sym_comment),
	11535: uint16(700),
	11536: uint16(1),
	11537: uint16(anon_sym_DOLLAR_LPAREN),
	11538: uint16(1018),
	11539: uint16(1),
	11540: uint16(sym_symbol),
	11541: uint16(327),
	11542: uint16(1),
	11543: uint16(sym_macro_variable),
	11544: uint16(4),
	11545: uint16(3),
	11546: uint16(1),
	11547: uint16(sym_comment),
	11548: uint16(700),
	11549: uint16(1),
	11550: uint16(anon_sym_DOLLAR_LPAREN),
	11551: uint16(1020),
	11552: uint16(1),
	11553: uint16(sym_symbol),
	11554: uint16(339),
	11555: uint16(1),
	11556: uint16(sym_macro_variable),
	11557: uint16(4),
	11558: uint16(3),
	11559: uint16(1),
	11560: uint16(sym_comment),
	11561: uint16(702),
	11562: uint16(1),
	11563: uint16(anon_sym_DQUOTE),
	11564: uint16(704),
	11565: uint16(1),
	11566: uint16(anon_sym_SQUOTE),
	11567: uint16(348),
	11568: uint16(1),
	11569: uint16(sym_string),
	11570: uint16(4),
	11571: uint16(3),
	11572: uint16(1),
	11573: uint16(sym_comment),
	11574: uint16(1000),
	11575: uint16(1),
	11576: uint16(anon_sym_DQUOTE),
	11577: uint16(1002),
	11578: uint16(1),
	11579: uint16(anon_sym_SQUOTE),
	11580: uint16(183),
	11581: uint16(1),
	11582: uint16(sym_string),
	11583: uint16(4),
	11584: uint16(3),
	11585: uint16(1),
	11586: uint16(sym_comment),
	11587: uint16(97),
	11588: uint16(1),
	11589: uint16(anon_sym_DQUOTE),
	11590: uint16(99),
	11591: uint16(1),
	11592: uint16(anon_sym_SQUOTE),
	11593: uint16(2),
	11594: uint16(1),
	11595: uint16(sym_string),
	11596: uint16(4),
	11597: uint16(568),
	11598: uint16(1),
	11599: uint16(sym_comment),
	11600: uint16(770),
	11601: uint16(1),
	11602: uint16(anon_sym_if),
	11603: uint16(1022),
	11604: uint16(1),
	11605: uint16(aux_sym_type_definition_token1),
	11606: uint16(370),
	11607: uint16(1),
	11608: uint16(sym_conditional_clause),
	11609: uint16(3),
	11610: uint16(335),
	11611: uint16(1),
	11612: uint16(anon_sym_if),
	11613: uint16(337),
	11614: uint16(1),
	11615: uint16(aux_sym_type_definition_token1),
	11616: uint16(568),
	11617: uint16(1),
	11618: uint16(sym_comment),
	11619: uint16(3),
	11620: uint16(375),
	11621: uint16(1),
	11622: uint16(anon_sym_if),
	11623: uint16(377),
	11624: uint16(1),
	11625: uint16(aux_sym_type_definition_token1),
	11626: uint16(568),
	11627: uint16(1),
	11628: uint16(sym_comment),
	11629: uint16(2),
	11630: uint16(568),
	11631: uint16(1),
	11632: uint16(sym_comment),
	11633: uint16(1022),
	11634: uint16(1),
	11635: uint16(aux_sym_type_definition_token1),
	11636: uint16(2),
	11637: uint16(568),
	11638: uint16(1),
	11639: uint16(sym_comment),
	11640: uint16(1024),
	11641: uint16(1),
	11642: uint16(aux_sym_type_definition_token1),
	11643: uint16(2),
	11644: uint16(568),
	11645: uint16(1),
	11646: uint16(sym_comment),
	11647: uint16(1026),
	11648: uint16(1),
	11649: uint16(aux_sym_type_definition_token1),
	11650: uint16(2),
	11651: uint16(3),
	11652: uint16(1),
	11653: uint16(sym_comment),
	11654: uint16(1028),
	11655: uint16(1),
	11656: uint16(sym__help_text),
	11657: uint16(2),
	11658: uint16(568),
	11659: uint16(1),
	11660: uint16(sym_comment),
	11661: uint16(1030),
	11662: uint16(1),
	11663: uint16(aux_sym_type_definition_token1),
	11664: uint16(2),
	11665: uint16(568),
	11666: uint16(1),
	11667: uint16(sym_comment),
	11668: uint16(1032),
	11669: uint16(1),
	11670: uint16(aux_sym_type_definition_token1),
	11671: uint16(2),
	11672: uint16(568),
	11673: uint16(1),
	11674: uint16(sym_comment),
	11675: uint16(1034),
	11676: uint16(1),
	11677: uint16(aux_sym_variable_token1),
	11678: uint16(2),
	11679: uint16(568),
	11680: uint16(1),
	11681: uint16(sym_comment),
	11682: uint16(1036),
	11683: uint16(1),
	11684: uint16(aux_sym_type_definition_token1),
	11685: uint16(2),
	11686: uint16(568),
	11687: uint16(1),
	11688: uint16(sym_comment),
	11689: uint16(1038),
	11690: uint16(1),
	11691: uint16(aux_sym_type_definition_token1),
	11692: uint16(2),
	11693: uint16(568),
	11694: uint16(1),
	11695: uint16(sym_comment),
	11696: uint16(1040),
	11697: uint16(1),
	11698: uint16(aux_sym_type_definition_token1),
	11699: uint16(2),
	11700: uint16(3),
	11701: uint16(1),
	11702: uint16(sym_comment),
	11703: uint16(1042),
	11704: uint16(1),
	11706: uint16(2),
	11707: uint16(568),
	11708: uint16(1),
	11709: uint16(sym_comment),
	11710: uint16(1044),
	11711: uint16(1),
	11712: uint16(aux_sym_variable_token1),
	11713: uint16(2),
	11714: uint16(568),
	11715: uint16(1),
	11716: uint16(sym_comment),
	11717: uint16(1004),
	11718: uint16(1),
	11719: uint16(aux_sym_type_definition_token1),
	11720: uint16(2),
	11721: uint16(568),
	11722: uint16(1),
	11723: uint16(sym_comment),
	11724: uint16(1046),
	11725: uint16(1),
	11726: uint16(aux_sym_type_definition_token1),
	11727: uint16(2),
	11728: uint16(568),
	11729: uint16(1),
	11730: uint16(sym_comment),
	11731: uint16(1048),
	11732: uint16(1),
	11733: uint16(aux_sym_type_definition_token1),
	11734: uint16(2),
	11735: uint16(3),
	11736: uint16(1),
	11737: uint16(sym_comment),
	11738: uint16(1050),
	11739: uint16(1),
	11740: uint16(sym_symbol),
	11741: uint16(2),
	11742: uint16(568),
	11743: uint16(1),
	11744: uint16(sym_comment),
	11745: uint16(1052),
	11746: uint16(1),
	11747: uint16(aux_sym_type_definition_token1),
	11748: uint16(2),
	11749: uint16(3),
	11750: uint16(1),
	11751: uint16(sym_comment),
	11752: uint16(1054),
	11753: uint16(1),
	11754: uint16(sym_symbol),
	11755: uint16(2),
	11756: uint16(568),
	11757: uint16(1),
	11758: uint16(sym_comment),
	11759: uint16(1056),
	11760: uint16(1),
	11761: uint16(aux_sym_type_definition_token1),
	11762: uint16(2),
	11763: uint16(568),
	11764: uint16(1),
	11765: uint16(sym_comment),
	11766: uint16(1058),
	11767: uint16(1),
	11768: uint16(aux_sym_type_definition_token1),
	11769: uint16(2),
	11770: uint16(568),
	11771: uint16(1),
	11772: uint16(sym_comment),
	11773: uint16(1060),
	11774: uint16(1),
	11775: uint16(aux_sym_type_definition_token1),
	11776: uint16(2),
	11777: uint16(568),
	11778: uint16(1),
	11779: uint16(sym_comment),
	11780: uint16(1062),
	11781: uint16(1),
	11782: uint16(aux_sym_type_definition_token1),
	11783: uint16(2),
	11784: uint16(568),
	11785: uint16(1),
	11786: uint16(sym_comment),
	11787: uint16(1064),
	11788: uint16(1),
	11789: uint16(aux_sym_type_definition_token1),
	11790: uint16(2),
	11791: uint16(568),
	11792: uint16(1),
	11793: uint16(sym_comment),
	11794: uint16(1066),
	11795: uint16(1),
	11796: uint16(aux_sym_type_definition_token1),
	11797: uint16(2),
	11798: uint16(3),
	11799: uint16(1),
	11800: uint16(sym_comment),
	11801: uint16(1068),
	11802: uint16(1),
	11803: uint16(sym__help_text),
	11804: uint16(2),
	11805: uint16(568),
	11806: uint16(1),
	11807: uint16(sym_comment),
	11808: uint16(1070),
	11809: uint16(1),
	11810: uint16(aux_sym_type_definition_token1),
}

var ts_small_parse_table_map = [373]uint32_t{
	1:   uint32(68),
	2:   uint32(136),
	3:   uint32(214),
	4:   uint32(292),
	5:   uint32(370),
	6:   uint32(448),
	7:   uint32(514),
	8:   uint32(592),
	9:   uint32(658),
	10:  uint32(713),
	11:  uint32(768),
	12:  uint32(825),
	13:  uint32(880),
	14:  uint32(935),
	15:  uint32(1013),
	16:  uint32(1091),
	17:  uint32(1169),
	18:  uint32(1247),
	19:  uint32(1325),
	20:  uint32(1378),
	21:  uint32(1431),
	22:  uint32(1486),
	23:  uint32(1539),
	24:  uint32(1592),
	25:  uint32(1644),
	26:  uint32(1696),
	27:  uint32(1754),
	28:  uint32(1806),
	29:  uint32(1858),
	30:  uint32(1912),
	31:  uint32(1972),
	32:  uint32(2034),
	33:  uint32(2084),
	34:  uint32(2144),
	35:  uint32(2194),
	36:  uint32(2244),
	37:  uint32(2302),
	38:  uint32(2352),
	39:  uint32(2408),
	40:  uint32(2460),
	41:  uint32(2540),
	42:  uint32(2620),
	43:  uint32(2664),
	44:  uint32(2708),
	45:  uint32(2752),
	46:  uint32(2796),
	47:  uint32(2840),
	48:  uint32(2884),
	49:  uint32(2928),
	50:  uint32(2972),
	51:  uint32(3016),
	52:  uint32(3060),
	53:  uint32(3104),
	54:  uint32(3148),
	55:  uint32(3192),
	56:  uint32(3236),
	57:  uint32(3280),
	58:  uint32(3324),
	59:  uint32(3368),
	60:  uint32(3410),
	61:  uint32(3452),
	62:  uint32(3494),
	63:  uint32(3536),
	64:  uint32(3578),
	65:  uint32(3620),
	66:  uint32(3662),
	67:  uint32(3704),
	68:  uint32(3746),
	69:  uint32(3788),
	70:  uint32(3830),
	71:  uint32(3872),
	72:  uint32(3914),
	73:  uint32(3956),
	74:  uint32(3998),
	75:  uint32(4040),
	76:  uint32(4112),
	77:  uint32(4184),
	78:  uint32(4226),
	79:  uint32(4286),
	80:  uint32(4346),
	81:  uint32(4406),
	82:  uint32(4466),
	83:  uint32(4526),
	84:  uint32(4586),
	85:  uint32(4634),
	86:  uint32(4682),
	87:  uint32(4719),
	88:  uint32(4756),
	89:  uint32(4812),
	90:  uint32(4847),
	91:  uint32(4882),
	92:  uint32(4936),
	93:  uint32(4990),
	94:  uint32(5044),
	95:  uint32(5098),
	96:  uint32(5152),
	97:  uint32(5206),
	98:  uint32(5260),
	99:  uint32(5314),
	100: uint32(5368),
	101: uint32(5422),
	102: uint32(5476),
	103: uint32(5530),
	104: uint32(5565),
	105: uint32(5596),
	106: uint32(5627),
	107: uint32(5658),
	108: uint32(5689),
	109: uint32(5727),
	110: uint32(5763),
	111: uint32(5795),
	112: uint32(5824),
	113: uint32(5853),
	114: uint32(5890),
	115: uint32(5927),
	116: uint32(5977),
	117: uint32(6027),
	118: uint32(6056),
	119: uint32(6085),
	120: uint32(6108),
	121: uint32(6143),
	122: uint32(6166),
	123: uint32(6213),
	124: uint32(6248),
	125: uint32(6295),
	126: uint32(6342),
	127: uint32(6364),
	128: uint32(6386),
	129: uint32(6408),
	130: uint32(6430),
	131: uint32(6452),
	132: uint32(6474),
	133: uint32(6496),
	134: uint32(6518),
	135: uint32(6540),
	136: uint32(6562),
	137: uint32(6584),
	138: uint32(6606),
	139: uint32(6628),
	140: uint32(6660),
	141: uint32(6682),
	142: uint32(6706),
	143: uint32(6730),
	144: uint32(6754),
	145: uint32(6778),
	146: uint32(6802),
	147: uint32(6826),
	148: uint32(6850),
	149: uint32(6880),
	150: uint32(6908),
	151: uint32(6934),
	152: uint32(6958),
	153: uint32(6982),
	154: uint32(7005),
	155: uint32(7028),
	156: uint32(7066),
	157: uint32(7104),
	158: uint32(7142),
	159: uint32(7180),
	160: uint32(7218),
	161: uint32(7240),
	162: uint32(7278),
	163: uint32(7300),
	164: uint32(7338),
	165: uint32(7376),
	166: uint32(7414),
	167: uint32(7452),
	168: uint32(7490),
	169: uint32(7528),
	170: uint32(7566),
	171: uint32(7604),
	172: uint32(7626),
	173: uint32(7664),
	174: uint32(7686),
	175: uint32(7724),
	176: uint32(7762),
	177: uint32(7800),
	178: uint32(7838),
	179: uint32(7860),
	180: uint32(7898),
	181: uint32(7920),
	182: uint32(7958),
	183: uint32(7980),
	184: uint32(8002),
	185: uint32(8040),
	186: uint32(8078),
	187: uint32(8116),
	188: uint32(8154),
	189: uint32(8176),
	190: uint32(8214),
	191: uint32(8252),
	192: uint32(8290),
	193: uint32(8312),
	194: uint32(8334),
	195: uint32(8372),
	196: uint32(8410),
	197: uint32(8448),
	198: uint32(8470),
	199: uint32(8492),
	200: uint32(8514),
	201: uint32(8536),
	202: uint32(8558),
	203: uint32(8580),
	204: uint32(8602),
	205: uint32(8640),
	206: uint32(8678),
	207: uint32(8716),
	208: uint32(8738),
	209: uint32(8760),
	210: uint32(8798),
	211: uint32(8836),
	212: uint32(8874),
	213: uint32(8912),
	214: uint32(8950),
	215: uint32(8988),
	216: uint32(9026),
	217: uint32(9064),
	218: uint32(9088),
	219: uint32(9126),
	220: uint32(9164),
	221: uint32(9202),
	222: uint32(9240),
	223: uint32(9278),
	224: uint32(9316),
	225: uint32(9339),
	226: uint32(9368),
	227: uint32(9397),
	228: uint32(9426),
	229: uint32(9455),
	230: uint32(9484),
	231: uint32(9513),
	232: uint32(9535),
	233: uint32(9555),
	234: uint32(9579),
	235: uint32(9597),
	236: uint32(9615),
	237: uint32(9633),
	238: uint32(9651),
	239: uint32(9676),
	240: uint32(9699),
	241: uint32(9724),
	242: uint32(9749),
	243: uint32(9774),
	244: uint32(9793),
	245: uint32(9814),
	246: uint32(9837),
	247: uint32(9860),
	248: uint32(9883),
	249: uint32(9908),
	250: uint32(9934),
	251: uint32(9960),
	252: uint32(9986),
	253: uint32(10012),
	254: uint32(10038),
	255: uint32(10064),
	256: uint32(10090),
	257: uint32(10114),
	258: uint32(10138),
	259: uint32(10164),
	260: uint32(10188),
	261: uint32(10214),
	262: uint32(10240),
	263: uint32(10266),
	264: uint32(10292),
	265: uint32(10318),
	266: uint32(10344),
	267: uint32(10370),
	268: uint32(10396),
	269: uint32(10422),
	270: uint32(10448),
	271: uint32(10474),
	272: uint32(10498),
	273: uint32(10522),
	274: uint32(10546),
	275: uint32(10570),
	276: uint32(10594),
	277: uint32(10612),
	278: uint32(10630),
	279: uint32(10648),
	280: uint32(10666),
	281: uint32(10684),
	282: uint32(10702),
	283: uint32(10720),
	284: uint32(10738),
	285: uint32(10756),
	286: uint32(10774),
	287: uint32(10792),
	288: uint32(10810),
	289: uint32(10828),
	290: uint32(10846),
	291: uint32(10864),
	292: uint32(10882),
	293: uint32(10900),
	294: uint32(10918),
	295: uint32(10936),
	296: uint32(10954),
	297: uint32(10972),
	298: uint32(10990),
	299: uint32(11008),
	300: uint32(11026),
	301: uint32(11044),
	302: uint32(11062),
	303: uint32(11080),
	304: uint32(11098),
	305: uint32(11116),
	306: uint32(11134),
	307: uint32(11145),
	308: uint32(11156),
	309: uint32(11167),
	310: uint32(11178),
	311: uint32(11189),
	312: uint32(11199),
	313: uint32(11209),
	314: uint32(11219),
	315: uint32(11229),
	316: uint32(11239),
	317: uint32(11249),
	318: uint32(11262),
	319: uint32(11275),
	320: uint32(11288),
	321: uint32(11301),
	322: uint32(11312),
	323: uint32(11325),
	324: uint32(11338),
	325: uint32(11351),
	326: uint32(11364),
	327: uint32(11377),
	328: uint32(11390),
	329: uint32(11403),
	330: uint32(11416),
	331: uint32(11429),
	332: uint32(11442),
	333: uint32(11455),
	334: uint32(11466),
	335: uint32(11479),
	336: uint32(11492),
	337: uint32(11505),
	338: uint32(11518),
	339: uint32(11531),
	340: uint32(11544),
	341: uint32(11557),
	342: uint32(11570),
	343: uint32(11583),
	344: uint32(11596),
	345: uint32(11609),
	346: uint32(11619),
	347: uint32(11629),
	348: uint32(11636),
	349: uint32(11643),
	350: uint32(11650),
	351: uint32(11657),
	352: uint32(11664),
	353: uint32(11671),
	354: uint32(11678),
	355: uint32(11685),
	356: uint32(11692),
	357: uint32(11699),
	358: uint32(11706),
	359: uint32(11713),
	360: uint32(11720),
	361: uint32(11727),
	362: uint32(11734),
	363: uint32(11741),
	364: uint32(11748),
	365: uint32(11755),
	366: uint32(11762),
	367: uint32(11769),
	368: uint32(11776),
	369: uint32(11783),
	370: uint32(11790),
	371: uint32(11797),
	372: uint32(11804),
}

var ts_parse_actions = [1072]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(322)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(336)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(266)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(267)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(332)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(353)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(174)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(352)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(321)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(333)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(281)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(282)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(283)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(338)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(344)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(223)),
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
		Fcount: uint8(1),
	}})),
	48: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(341)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(260)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(345)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(165)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(269)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(374)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(182)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(342)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(381)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(211)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(193)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(140)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(138)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(143)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(145)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:         uint8(TSParseActionTypeReduce),
		Fchild_count:        uint8(1),
		Fsymbol:             uint16(sym_name),
		Fdynamic_precedence: int16(-int32(1)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	94: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:         uint8(TSParseActionTypeReduce),
		Fchild_count:        uint8(1),
		Fsymbol:             uint16(sym_name),
		Fdynamic_precedence: int16(-int32(1)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(293)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(307)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(11)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(264)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
	113: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(293)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(307)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(260)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(345)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(201)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(167)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
	133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(165)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(269)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(374)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	141: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(182)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(342)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	147: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(381)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
	151: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(12)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
	}})),
	153: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_comment_entry),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_config),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_menuconfig),
		Fproduction_id: uint16(1),
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
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(16)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(280)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(315)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(287)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_comment_entry),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(315)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(287)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_macro_variable),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_macro_variable),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(sym_expression),
	})))),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_string),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_macro_variable),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_macro_variable),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(265)),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(346)),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(225)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(226)),
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
		Fcount: uint8(2),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(284)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(372)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(214)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(328)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(360)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	239: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_config_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(24)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_config),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(346)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(225)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(226)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(284)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(372)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(328)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(360)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_menuconfig),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_comment_entry),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_comment_entry),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unary_expression),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unary_expression),
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
		Fcount: uint8(1),
	}})),
	278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_binary_expression),
		Fproduction_id: uint16(6),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(176)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression),
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
		Fsymbol:      uint16(sym_expression),
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
		Fcount: uint8(1),
	}})),
	292: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_parenthesized_expression),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
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
		Fsymbol:      uint16(sym_limiting_menu_display),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_limiting_menu_display),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(233)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(206)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(206)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fcount: uint8(1),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_weak_reverse_dependencies),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_weak_reverse_dependencies),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type_definition_default),
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
		Fsymbol:      uint16(sym_type_definition_default),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_numerical_ranges),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_numerical_ranges),
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
		Fcount: uint8(1),
	}})),
	332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_definition),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_definition),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_prompt),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_input_prompt),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_default_value),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_default_value),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_reverse_dependencies),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_reverse_dependencies),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_definition_default),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_type_definition_default),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_help_text),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_help_text),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_reverse_dependencies),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_reverse_dependencies),
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
		Fsymbol:      uint16(sym_weak_reverse_dependencies),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_weak_reverse_dependencies),
	})))),
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
		Fcount: uint8(1),
	}})),
	364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_type_definition),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_type_definition),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_dependencies),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_dependencies),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type_definition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_type_definition),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_prompt),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_input_prompt),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_default_value),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_default_value),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_numerical_ranges),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_numerical_ranges),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(199)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(200)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(188)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(260)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(345)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	410: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(269)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(374)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(342)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(381)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(265)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(346)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(284)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(372)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(328)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(360)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	444: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	446: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(94)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
	450: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(303)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	455: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
	456: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(306)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
	}})),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(262)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(303)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(306)),
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
		Fcount: uint8(2),
	}})),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(321)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	469: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(333)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	472: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(281)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	475: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	476: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(282)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(283)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	482: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(51)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fcount: uint8(2),
	}})),
	486: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(338)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(344)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	492: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(223)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	495: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	496: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(341)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
	}})),
	498: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(218)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(322)),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(336)),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		Fsymbol:      uint16(aux_sym_configuration_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(266)),
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
		Fcount: uint8(2),
	}})),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	521: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(267)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	523: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(50)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	526: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	527: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(332)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	528: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	529: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(353)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	532: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	533: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(174)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	535: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(352)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	538: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(217)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	542: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(192)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_configdefault),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_configdefault_repeat1),
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
		Fcount: uint8(2),
	}})),
	554: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_configdefault_repeat1),
	})))),
	555: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(231)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	557: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(123)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	560: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
	561: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(278)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	562: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	563: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(301)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	566: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(302)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	569: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(123)),
	}})))),
	572: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	574: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	575: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(301)),
	}})))),
	576: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(302)),
	}})))),
	578: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		Fcount: uint8(1),
	}})),
	581: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(142)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(221)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(368)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
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
		Fcount: uint8(1),
	}})),
	595: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(363)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_configdefault_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_configdefault_repeat1),
	})))),
	600: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(229)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	601: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	602: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_configdefault),
		Fproduction_id: uint16(1),
	})))),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	604: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	605: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	606: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(130)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	609: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
	610: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(276)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	611: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	612: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_name_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(297)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(298)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	620: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(213)),
	}})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
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
		Fcount: uint8(1),
	}})),
	624: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(297)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	628: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(298)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	632: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		Fcount: uint8(2),
	}})),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
	})))),
	636: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	637: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	638: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
	})))),
	639: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	640: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(221)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
	})))),
	644: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(232)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	646: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
	})))),
	647: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(278)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
	})))),
	650: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(301)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	651: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	652: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(8),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(302)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_choice),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount: uint8(1),
	}})),
	657: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_mainmenu),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	659: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_choice),
	})))),
	660: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	661: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_source),
	})))),
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
		Fcount: uint8(1),
	}})),
	663: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_menu),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	665: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if),
		Fproduction_id: uint16(3),
	})))),
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
		Fcount: uint8(1),
	}})),
	667: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
	}})),
	669: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_choice),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	671: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_choice),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_menu),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	675: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if),
		Fproduction_id: uint16(3),
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
		Fcount: uint8(1),
	}})),
	677: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
	}})),
	679: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(7),
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
		Fcount: uint8(1),
	}})),
	681: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(5),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(194)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_variable_repeat1),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount: uint8(1),
	}})),
	687: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	689: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(196)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	691: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
	}})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_menu),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(227)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(215)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	701: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	703: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(297)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	705: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(298)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(234)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	709: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(4),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(175)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(230)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_menu),
		Fproduction_id: uint16(1),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	727: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_source),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_if),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_choice),
		Fproduction_id: uint16(1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_choice),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	735: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_menu),
		Fproduction_id: uint16(1),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(278)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(301)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(302)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(222)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_mainmenu),
		Fproduction_id: uint16(1),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_if),
		Fproduction_id: uint16(3),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	759: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_choice),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	761: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_variable),
		Fproduction_id: uint16(7),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_choice),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_menu),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	769: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_name_repeat1),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(216)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Fcount: uint8(1),
	}})),
	777: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(173)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(67)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(185)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(116)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_conditional_clause),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(245)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(274)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(261)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(291)),
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
		Fcount: uint8(1),
	}})),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(292)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(351)),
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
		Fcount: uint8(1),
	}})),
	833: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_macro_variable_repeat1),
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
		Fcount: uint8(2),
	}})),
	835: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_macro_variable_repeat1),
	})))),
	836: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(274)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	837: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	838: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_macro_variable_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(261)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	841: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_macro_variable_repeat1),
	})))),
	842: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(291)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	843: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	844: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_macro_variable_repeat1),
	})))),
	845: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(292)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	846: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	847: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	848: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	849: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	850: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	851: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	852: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	853: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	854: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	855: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(263)),
	}})))),
	856: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	857: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	858: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	859: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	860: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	861: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	862: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	863: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	864: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	865: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
	}})))),
	866: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	867: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(323)),
	}})))),
	868: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	869: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(271)),
	}})))),
	870: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	871: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(326)),
	}})))),
	872: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	873: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(325)),
	}})))),
	874: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	875: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(273)),
	}})))),
	876: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	877: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(324)),
	}})))),
	878: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	879: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(316)),
	}})))),
	880: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	881: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	882: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	883: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(318)),
	}})))),
	884: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	885: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	886: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	887: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(277)),
	}})))),
	888: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	889: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	890: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	891: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(161)),
	}})))),
	892: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	893: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(279)),
	}})))),
	894: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	895: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	896: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	897: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	898: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	899: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	900: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	901: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	902: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(270)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	903: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	904: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	905: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	906: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	907: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(286)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	908: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	910: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	911: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(290)),
	}})))),
	912: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	913: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	914: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	915: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(314)),
	}})))),
	916: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	917: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	918: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	919: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(270)),
	}})))),
	920: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	921: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	922: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	923: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(286)),
	}})))),
	924: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	925: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(320)),
	}})))),
	926: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	927: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(294)),
	}})))),
	928: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	929: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(295)),
	}})))),
	930: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	931: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	932: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	933: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(312)),
	}})))),
	934: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	935: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(317)),
	}})))),
	936: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	937: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(198)),
	}})))),
	938: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	939: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(308)),
	}})))),
	940: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	941: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	942: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	943: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(299)),
	}})))),
	944: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	945: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(300)),
	}})))),
	946: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	947: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
	}})))),
	948: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	949: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	950: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	951: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(304)),
	}})))),
	952: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	953: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(305)),
	}})))),
	954: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	955: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	956: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	957: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(309)),
	}})))),
	958: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	959: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	960: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	961: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(310)),
	}})))),
	962: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	963: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(288)),
	}})))),
	964: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	965: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(172)),
	}})))),
	966: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	967: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	968: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	969: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(311)),
	}})))),
	970: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	971: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat2),
	})))),
	972: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	973: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	974: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat2),
	})))),
	975: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(314)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	976: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	977: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_repeat2),
	})))),
	978: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	979: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(289)),
	}})))),
	980: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	981: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_macro_variable_repeat1),
		Fproduction_id: uint16(2),
	})))),
	982: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	983: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	984: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	985: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
	}})))),
	986: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	987: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	988: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	989: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(350)),
	}})))),
	990: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	991: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	992: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	993: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	994: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	995: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
	}})))),
	996: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	997: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	998: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	999: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1000: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1001: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(296)),
	}})))),
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1003: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(313)),
	}})))),
	1004: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1005: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1006: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1007: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1008: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1009: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1010: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1011: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(349)),
	}})))),
	1012: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1013: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1014: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1015: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1016: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1017: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(356)),
	}})))),
	1018: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1019: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(327)),
	}})))),
	1020: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1021: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(339)),
	}})))),
	1022: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1023: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1024: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1025: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
	}})))),
	1026: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1027: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1028: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1029: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1030: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1031: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1032: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1033: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1034: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1035: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(212)),
	}})))),
	1036: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1037: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1038: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1039: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1040: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1041: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1042: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1043: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	1044: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1045: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(147)),
	}})))),
	1046: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1047: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1048: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1049: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1050: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1051: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(335)),
	}})))),
	1052: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1053: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1054: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1055: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(347)),
	}})))),
	1056: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1057: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1058: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1059: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1060: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1061: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(355)),
	}})))),
	1062: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1063: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1064: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1065: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1066: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1067: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1068: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1069: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1070: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1071: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token__help_text = 0

var ts_external_scanner_symbol_map = [1]TSSymbol{
	0: uint16(sym__help_text),
}

var ts_external_scanner_states = [2][1]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_kconfig(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Ftoken_count:               uint32(TOKEN_COUNT),
	Fexternal_token_count:      uint32(EXTERNAL_TOKEN_COUNT),
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
	Fkeyword_capture_token:     uint16(sym_symbol),
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
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(ts_lex_keywords)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_kconfig_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_kconfig_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_kconfig_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_kconfig_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_kconfig_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "end\x00symbol\x00mainmenu\x00config\x00configdefault\x00menuconfig\x00choice\x00endchoice\x00comment\x00menu\x00endmenu\x00if\x00endif\x00source\x00rsource\x00osource\x00orsource\x00=\x00:=\x00+=\x00?=\x00,\x00variable_token1\x00bool\x00tristate\x00int\x00hex\x00string\x00type_definition_token1\x00prompt\x00default\x00def_bool\x00def_tristate\x00def_int\x00def_hex\x00def_string\x00depends on\x00select\x00imply\x00visible if\x00range\x00help\x00optional\x00modules\x00!\x00||\x00&&\x00!=\x00<\x00>\x00<=\x00>=\x00(\x00)\x00$(\x00macro_content\x00\"\x00string_content\x00string_token2\x00'\x00text\x00configuration\x00_entry\x00comment_entry\x00variable\x00_config_option\x00type_definition\x00input_prompt\x00default_value\x00type_definition_default\x00dependencies\x00reverse_dependencies\x00weak_reverse_dependencies\x00limiting_menu_display\x00numerical_ranges\x00help_text\x00conditional_clause\x00expression\x00unary_expression\x00binary_expression\x00parenthesized_expression\x00macro_variable\x00name\x00configuration_repeat1\x00config_repeat1\x00configdefault_repeat1\x00variable_repeat1\x00macro_variable_repeat1\x00string_repeat1\x00string_repeat2\x00name_repeat1\x00condition\x00left\x00operator\x00right\x00"
