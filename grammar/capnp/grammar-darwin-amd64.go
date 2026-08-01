// Code generated for darwin/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -D__attribute__(...)= -D__extension__= -D_Nonnull= -D_Nullable= -D_Null_unspecified= -DAPI_AVAILABLE(...)= -DAPI_UNAVAILABLE(...)= -DAPI_DEPRECATED(...)= -DAPI_DEPRECATED_WITH_REPLACEMENT(...)= -D__API_AVAILABLE(...)= -D__API_UNAVAILABLE(...)= -D__API_DEPRECATED(...)= -D__API_DEPRECATED_WITH_REPLACEMENT(...)= -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-capnp/src -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-capnp -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-capnp/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build darwin && amd64

package grammar_capnp

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 17
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
const EXTERNAL_TOKEN_COUNT = 0
const FIELD_COUNT = 0
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
const LARGE_STATE_COUNT = 2
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
const MAX_ALIAS_SEQUENCE_LENGTH = 13
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
const PRODUCTION_ID_COUNT = 35
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
const STATE_COUNT = 553
const SV_INTERRUPT = "SA_RESTART"
const SV_NOCLDSTOP = "SA_NOCLDSTOP"
const SV_NODEFER = "SA_NODEFER"
const SV_ONSTACK = "SA_ONSTACK"
const SV_RESETHAND = "SA_RESETHAND"
const SV_SIGINFO = "SA_SIGINFO"
const SYMBOL_COUNT = 157
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
const TOKEN_COUNT = 71
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
const WEXITED = 0x00000004
const WNOHANG = 0x00000001
const WNOWAIT = 0x00000020
const WSTOPPED = 0x00000008
const WUNTRACED = 0x00000002
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

const sym_identifier = 1
const sym_unique_id = 2
const anon_sym_SEMI = 3
const anon_sym_using = 4
const anon_sym_EQ = 5
const anon_sym_import = 6
const anon_sym_DOT = 7
const anon_sym_DOLLARimport = 8
const anon_sym_LPAREN = 9
const anon_sym_RPAREN = 10
const anon_sym_namespace = 11
const anon_sym_DOLLAR = 12
const aux_sym_top_level_annotation_body_token1 = 13
const anon_sym_COMMA = 14
const anon_sym_annotation = 15
const anon_sym_COLON = 16
const anon_sym_STAR = 17
const anon_sym_const = 18
const anon_sym_enumerant = 19
const anon_sym_field = 20
const anon_sym_file = 21
const anon_sym_method = 22
const anon_sym_param = 23
const anon_sym_enum = 24
const anon_sym_group = 25
const anon_sym_interface = 26
const anon_sym_struct = 27
const anon_sym_union = 28
const anon_sym_LBRACK = 29
const anon_sym_RBRACK = 30
const anon_sym_LBRACE = 31
const anon_sym_RBRACE = 32
const anon_sym_extends = 33
const anon_sym_DASH_GT = 34
const anon_sym_AnyPointer = 35
const anon_sym_Bool = 36
const anon_sym_Int8 = 37
const anon_sym_Int16 = 38
const anon_sym_Int32 = 39
const anon_sym_Int64 = 40
const anon_sym_UInt8 = 41
const anon_sym_UInt16 = 42
const anon_sym_UInt32 = 43
const anon_sym_UInt64 = 44
const anon_sym_Float32 = 45
const anon_sym_Float64 = 46
const anon_sym_Text = 47
const anon_sym_Data = 48
const anon_sym_Void = 49
const anon_sym_List = 50
const sym_number = 51
const sym__normal_version = 52
const aux_sym__inline_version_token1 = 53
const sym_float = 54
const anon_sym_true = 55
const anon_sym_false = 56
const anon_sym_0x = 57
const aux_sym_data_token1 = 58
const sym_void = 59
const anon_sym_embed = 60
const anon_sym_DQUOTE = 61
const anon_sym_SQUOTE = 62
const anon_sym_BQUOTE = 63
const sym_unescaped_double_string_fragment = 64
const sym_unescaped_single_string_fragment = 65
const sym_unescaped_block_string_fragment = 66
const aux_sym__escape_sequence_token1 = 67
const sym_escape_sequence = 68
const sym__identifier_no_period = 69
const sym_comment = 70
const sym_message = 71
const sym_statement = 72
const sym_unique_id_statement = 73
const sym_using_directive = 74
const sym_replace_using = 75
const sym_import_using = 76
const sym_import = 77
const sym_top_level_annotation = 78
const sym_top_level_annotation_body = 79
const sym_annotation = 80
const sym_annotation_targets = 81
const sym_annotation_target = 82
const sym_annotation_literal = 83
const sym_annotation_array = 84
const sym__annotation_array_def = 85
const sym_definition = 86
const sym_struct = 87
const sym_nested_struct = 88
const sym_enum = 89
const sym_nested_enum = 90
const sym_enum_field = 91
const sym_group = 92
const sym_field = 93
const sym_union = 94
const sym_nested_union = 95
const sym__unnamed_union = 96
const sym__named_union = 97
const sym_union_field = 98
const sym_interface = 99
const sym_method = 100
const sym_method_parameters = 101
const sym_parameters = 102
const sym_parameter = 103
const sym_return_type = 104
const sym_named_return_types = 105
const sym_unnamed_return_type = 106
const sym_named_return_type = 107
const sym_field_type = 108
const sym_primitive_type = 109
const sym_list_type = 110
const sym_custom_type = 111
const sym_const = 112
const sym_const_value = 113
const sym__same_scope_const_value = 114
const sym_field_version = 115
const sym__inline_version = 116
const sym_boolean = 117
const sym_data = 118
const sym_const_list = 119
const sym_struct_shorthand = 120
const sym__internal_const_identifier = 121
const sym_embedded_file = 122
const sym_generics = 123
const sym_implicit_generics = 124
const sym_generic_parameters = 125
const sym_string = 126
const sym_concatenated_string = 127
const sym_block_text = 128
const sym__escape_sequence = 129
const sym__annotation_definition_identifier = 130
const sym__method_identifier = 131
const aux_sym_message_repeat1 = 132
const aux_sym_import_using_repeat1 = 133
const aux_sym_top_level_annotation_body_repeat1 = 134
const aux_sym_annotation_repeat1 = 135
const aux_sym_annotation_targets_repeat1 = 136
const aux_sym__annotation_call_repeat1 = 137
const aux_sym_annotation_array_repeat1 = 138
const aux_sym__annotation_array_def_repeat1 = 139
const aux_sym__annotation_array_def_repeat2 = 140
const aux_sym_struct_repeat1 = 141
const aux_sym_enum_repeat1 = 142
const aux_sym_group_repeat1 = 143
const aux_sym__unnamed_union_repeat1 = 144
const aux_sym_interface_repeat1 = 145
const aux_sym_parameters_repeat1 = 146
const aux_sym_named_return_type_repeat1 = 147
const aux_sym_custom_type_repeat1 = 148
const aux_sym_struct_shorthand_repeat1 = 149
const aux_sym__internal_const_identifier_repeat1 = 150
const aux_sym_generic_parameters_repeat1 = 151
const aux_sym_string_repeat1 = 152
const aux_sym_string_repeat2 = 153
const aux_sym_concatenated_string_repeat1 = 154
const aux_sym_block_text_repeat1 = 155
const aux_sym_block_text_repeat2 = 156
const alias_sym_annotation_identifier = 157
const alias_sym_attribute = 158
const alias_sym_enum_identifier = 159
const alias_sym_enum_member = 160
const alias_sym_extend_type = 161
const alias_sym_field_identifier = 162
const alias_sym_generic_identifier = 163
const alias_sym_implicit_generic_parameters = 164
const alias_sym_import_path = 165
const alias_sym_local_const = 166
const alias_sym_method_identifier = 167
const alias_sym_namespace = 168
const alias_sym_param_identifier = 169
const alias_sym_property = 170
const alias_sym_return_identifier = 171
const alias_sym_type_definition = 172
const alias_sym_type_identifier = 173

var ts_symbol_names = [174]uintptr{
	0:   __ccgo_ts,
	1:   __ccgo_ts + 4,
	2:   __ccgo_ts + 37,
	3:   __ccgo_ts + 47,
	4:   __ccgo_ts + 49,
	5:   __ccgo_ts + 55,
	6:   __ccgo_ts + 57,
	7:   __ccgo_ts + 64,
	8:   __ccgo_ts + 66,
	9:   __ccgo_ts + 74,
	10:  __ccgo_ts + 76,
	11:  __ccgo_ts + 78,
	12:  __ccgo_ts + 88,
	13:  __ccgo_ts + 90,
	14:  __ccgo_ts + 107,
	15:  __ccgo_ts + 109,
	16:  __ccgo_ts + 120,
	17:  __ccgo_ts + 122,
	18:  __ccgo_ts + 124,
	19:  __ccgo_ts + 130,
	20:  __ccgo_ts + 140,
	21:  __ccgo_ts + 146,
	22:  __ccgo_ts + 151,
	23:  __ccgo_ts + 158,
	24:  __ccgo_ts + 164,
	25:  __ccgo_ts + 169,
	26:  __ccgo_ts + 175,
	27:  __ccgo_ts + 185,
	28:  __ccgo_ts + 192,
	29:  __ccgo_ts + 198,
	30:  __ccgo_ts + 200,
	31:  __ccgo_ts + 202,
	32:  __ccgo_ts + 204,
	33:  __ccgo_ts + 206,
	34:  __ccgo_ts + 214,
	35:  __ccgo_ts + 217,
	36:  __ccgo_ts + 228,
	37:  __ccgo_ts + 233,
	38:  __ccgo_ts + 238,
	39:  __ccgo_ts + 244,
	40:  __ccgo_ts + 250,
	41:  __ccgo_ts + 256,
	42:  __ccgo_ts + 262,
	43:  __ccgo_ts + 269,
	44:  __ccgo_ts + 276,
	45:  __ccgo_ts + 283,
	46:  __ccgo_ts + 291,
	47:  __ccgo_ts + 299,
	48:  __ccgo_ts + 304,
	49:  __ccgo_ts + 309,
	50:  __ccgo_ts + 314,
	51:  __ccgo_ts + 319,
	52:  __ccgo_ts + 326,
	53:  __ccgo_ts + 342,
	54:  __ccgo_ts + 355,
	55:  __ccgo_ts + 361,
	56:  __ccgo_ts + 366,
	57:  __ccgo_ts + 372,
	58:  __ccgo_ts + 381,
	59:  __ccgo_ts + 393,
	60:  __ccgo_ts + 398,
	61:  __ccgo_ts + 404,
	62:  __ccgo_ts + 406,
	63:  __ccgo_ts + 408,
	64:  __ccgo_ts + 410,
	65:  __ccgo_ts + 410,
	66:  __ccgo_ts + 410,
	67:  __ccgo_ts + 426,
	68:  __ccgo_ts + 450,
	69:  __ccgo_ts + 466,
	70:  __ccgo_ts + 483,
	71:  __ccgo_ts + 491,
	72:  __ccgo_ts + 499,
	73:  __ccgo_ts + 509,
	74:  __ccgo_ts + 529,
	75:  __ccgo_ts + 545,
	76:  __ccgo_ts + 559,
	77:  __ccgo_ts + 57,
	78:  __ccgo_ts + 572,
	79:  __ccgo_ts + 593,
	80:  __ccgo_ts + 109,
	81:  __ccgo_ts + 619,
	82:  __ccgo_ts + 638,
	83:  __ccgo_ts + 656,
	84:  __ccgo_ts + 675,
	85:  __ccgo_ts + 692,
	86:  __ccgo_ts + 714,
	87:  __ccgo_ts + 185,
	88:  __ccgo_ts + 725,
	89:  __ccgo_ts + 164,
	90:  __ccgo_ts + 739,
	91:  __ccgo_ts + 751,
	92:  __ccgo_ts + 169,
	93:  __ccgo_ts + 140,
	94:  __ccgo_ts + 192,
	95:  __ccgo_ts + 762,
	96:  __ccgo_ts + 775,
	97:  __ccgo_ts + 790,
	98:  __ccgo_ts + 803,
	99:  __ccgo_ts + 175,
	100: __ccgo_ts + 151,
	101: __ccgo_ts + 815,
	102: __ccgo_ts + 833,
	103: __ccgo_ts + 844,
	104: __ccgo_ts + 854,
	105: __ccgo_ts + 866,
	106: __ccgo_ts + 885,
	107: __ccgo_ts + 905,
	108: __ccgo_ts + 923,
	109: __ccgo_ts + 934,
	110: __ccgo_ts + 949,
	111: __ccgo_ts + 959,
	112: __ccgo_ts + 124,
	113: __ccgo_ts + 971,
	114: __ccgo_ts + 983,
	115: __ccgo_ts + 1007,
	116: __ccgo_ts + 1021,
	117: __ccgo_ts + 1037,
	118: __ccgo_ts + 1045,
	119: __ccgo_ts + 1050,
	120: __ccgo_ts + 1061,
	121: __ccgo_ts + 1078,
	122: __ccgo_ts + 1105,
	123: __ccgo_ts + 1119,
	124: __ccgo_ts + 1128,
	125: __ccgo_ts + 1146,
	126: __ccgo_ts + 1165,
	127: __ccgo_ts + 1172,
	128: __ccgo_ts + 1192,
	129: __ccgo_ts + 1203,
	130: __ccgo_ts + 1220,
	131: __ccgo_ts + 1254,
	132: __ccgo_ts + 1273,
	133: __ccgo_ts + 1289,
	134: __ccgo_ts + 1310,
	135: __ccgo_ts + 1344,
	136: __ccgo_ts + 1363,
	137: __ccgo_ts + 1390,
	138: __ccgo_ts + 1415,
	139: __ccgo_ts + 1440,
	140: __ccgo_ts + 1470,
	141: __ccgo_ts + 1500,
	142: __ccgo_ts + 1515,
	143: __ccgo_ts + 1528,
	144: __ccgo_ts + 1542,
	145: __ccgo_ts + 1565,
	146: __ccgo_ts + 1583,
	147: __ccgo_ts + 1602,
	148: __ccgo_ts + 1628,
	149: __ccgo_ts + 1648,
	150: __ccgo_ts + 1673,
	151: __ccgo_ts + 1708,
	152: __ccgo_ts + 1735,
	153: __ccgo_ts + 1750,
	154: __ccgo_ts + 1765,
	155: __ccgo_ts + 1793,
	156: __ccgo_ts + 1812,
	157: __ccgo_ts + 1831,
	158: __ccgo_ts + 1853,
	159: __ccgo_ts + 1863,
	160: __ccgo_ts + 1879,
	161: __ccgo_ts + 1891,
	162: __ccgo_ts + 1903,
	163: __ccgo_ts + 1920,
	164: __ccgo_ts + 1939,
	165: __ccgo_ts + 1967,
	166: __ccgo_ts + 1979,
	167: __ccgo_ts + 1991,
	168: __ccgo_ts + 78,
	169: __ccgo_ts + 2009,
	170: __ccgo_ts + 2026,
	171: __ccgo_ts + 2035,
	172: __ccgo_ts + 2053,
	173: __ccgo_ts + 2069,
}

var ts_symbol_map = [174]TSSymbol{
	1:   uint16(sym_identifier),
	2:   uint16(sym_unique_id),
	3:   uint16(anon_sym_SEMI),
	4:   uint16(anon_sym_using),
	5:   uint16(anon_sym_EQ),
	6:   uint16(anon_sym_import),
	7:   uint16(anon_sym_DOT),
	8:   uint16(anon_sym_DOLLARimport),
	9:   uint16(anon_sym_LPAREN),
	10:  uint16(anon_sym_RPAREN),
	11:  uint16(anon_sym_namespace),
	12:  uint16(anon_sym_DOLLAR),
	13:  uint16(aux_sym_top_level_annotation_body_token1),
	14:  uint16(anon_sym_COMMA),
	15:  uint16(anon_sym_annotation),
	16:  uint16(anon_sym_COLON),
	17:  uint16(anon_sym_STAR),
	18:  uint16(anon_sym_const),
	19:  uint16(anon_sym_enumerant),
	20:  uint16(anon_sym_field),
	21:  uint16(anon_sym_file),
	22:  uint16(anon_sym_method),
	23:  uint16(anon_sym_param),
	24:  uint16(anon_sym_enum),
	25:  uint16(anon_sym_group),
	26:  uint16(anon_sym_interface),
	27:  uint16(anon_sym_struct),
	28:  uint16(anon_sym_union),
	29:  uint16(anon_sym_LBRACK),
	30:  uint16(anon_sym_RBRACK),
	31:  uint16(anon_sym_LBRACE),
	32:  uint16(anon_sym_RBRACE),
	33:  uint16(anon_sym_extends),
	34:  uint16(anon_sym_DASH_GT),
	35:  uint16(anon_sym_AnyPointer),
	36:  uint16(anon_sym_Bool),
	37:  uint16(anon_sym_Int8),
	38:  uint16(anon_sym_Int16),
	39:  uint16(anon_sym_Int32),
	40:  uint16(anon_sym_Int64),
	41:  uint16(anon_sym_UInt8),
	42:  uint16(anon_sym_UInt16),
	43:  uint16(anon_sym_UInt32),
	44:  uint16(anon_sym_UInt64),
	45:  uint16(anon_sym_Float32),
	46:  uint16(anon_sym_Float64),
	47:  uint16(anon_sym_Text),
	48:  uint16(anon_sym_Data),
	49:  uint16(anon_sym_Void),
	50:  uint16(anon_sym_List),
	51:  uint16(sym_number),
	52:  uint16(sym__normal_version),
	53:  uint16(aux_sym__inline_version_token1),
	54:  uint16(sym_float),
	55:  uint16(anon_sym_true),
	56:  uint16(anon_sym_false),
	57:  uint16(anon_sym_0x),
	58:  uint16(aux_sym_data_token1),
	59:  uint16(sym_void),
	60:  uint16(anon_sym_embed),
	61:  uint16(anon_sym_DQUOTE),
	62:  uint16(anon_sym_SQUOTE),
	63:  uint16(anon_sym_BQUOTE),
	64:  uint16(sym_unescaped_double_string_fragment),
	65:  uint16(sym_unescaped_double_string_fragment),
	66:  uint16(sym_unescaped_double_string_fragment),
	67:  uint16(aux_sym__escape_sequence_token1),
	68:  uint16(sym_escape_sequence),
	69:  uint16(sym__identifier_no_period),
	70:  uint16(sym_comment),
	71:  uint16(sym_message),
	72:  uint16(sym_statement),
	73:  uint16(sym_unique_id_statement),
	74:  uint16(sym_using_directive),
	75:  uint16(sym_replace_using),
	76:  uint16(sym_import_using),
	77:  uint16(sym_import),
	78:  uint16(sym_top_level_annotation),
	79:  uint16(sym_top_level_annotation_body),
	80:  uint16(sym_annotation),
	81:  uint16(sym_annotation_targets),
	82:  uint16(sym_annotation_target),
	83:  uint16(sym_annotation_literal),
	84:  uint16(sym_annotation_array),
	85:  uint16(sym__annotation_array_def),
	86:  uint16(sym_definition),
	87:  uint16(sym_struct),
	88:  uint16(sym_nested_struct),
	89:  uint16(sym_enum),
	90:  uint16(sym_nested_enum),
	91:  uint16(sym_enum_field),
	92:  uint16(sym_group),
	93:  uint16(sym_field),
	94:  uint16(sym_union),
	95:  uint16(sym_nested_union),
	96:  uint16(sym__unnamed_union),
	97:  uint16(sym__named_union),
	98:  uint16(sym_union_field),
	99:  uint16(sym_interface),
	100: uint16(sym_method),
	101: uint16(sym_method_parameters),
	102: uint16(sym_parameters),
	103: uint16(sym_parameter),
	104: uint16(sym_return_type),
	105: uint16(sym_named_return_types),
	106: uint16(sym_unnamed_return_type),
	107: uint16(sym_named_return_type),
	108: uint16(sym_field_type),
	109: uint16(sym_primitive_type),
	110: uint16(sym_list_type),
	111: uint16(sym_custom_type),
	112: uint16(sym_const),
	113: uint16(sym_const_value),
	114: uint16(sym__same_scope_const_value),
	115: uint16(sym_field_version),
	116: uint16(sym__inline_version),
	117: uint16(sym_boolean),
	118: uint16(sym_data),
	119: uint16(sym_const_list),
	120: uint16(sym_struct_shorthand),
	121: uint16(sym__internal_const_identifier),
	122: uint16(sym_embedded_file),
	123: uint16(sym_generics),
	124: uint16(sym_implicit_generics),
	125: uint16(sym_generic_parameters),
	126: uint16(sym_string),
	127: uint16(sym_concatenated_string),
	128: uint16(sym_block_text),
	129: uint16(sym__escape_sequence),
	130: uint16(sym__annotation_definition_identifier),
	131: uint16(sym__method_identifier),
	132: uint16(aux_sym_message_repeat1),
	133: uint16(aux_sym_import_using_repeat1),
	134: uint16(aux_sym_top_level_annotation_body_repeat1),
	135: uint16(aux_sym_annotation_repeat1),
	136: uint16(aux_sym_annotation_targets_repeat1),
	137: uint16(aux_sym__annotation_call_repeat1),
	138: uint16(aux_sym_annotation_array_repeat1),
	139: uint16(aux_sym__annotation_array_def_repeat1),
	140: uint16(aux_sym__annotation_array_def_repeat2),
	141: uint16(aux_sym_struct_repeat1),
	142: uint16(aux_sym_enum_repeat1),
	143: uint16(aux_sym_group_repeat1),
	144: uint16(aux_sym__unnamed_union_repeat1),
	145: uint16(aux_sym_interface_repeat1),
	146: uint16(aux_sym_parameters_repeat1),
	147: uint16(aux_sym_named_return_type_repeat1),
	148: uint16(aux_sym_custom_type_repeat1),
	149: uint16(aux_sym_struct_shorthand_repeat1),
	150: uint16(aux_sym__internal_const_identifier_repeat1),
	151: uint16(aux_sym_generic_parameters_repeat1),
	152: uint16(aux_sym_string_repeat1),
	153: uint16(aux_sym_string_repeat2),
	154: uint16(aux_sym_concatenated_string_repeat1),
	155: uint16(aux_sym_block_text_repeat1),
	156: uint16(aux_sym_block_text_repeat2),
	157: uint16(alias_sym_annotation_identifier),
	158: uint16(alias_sym_attribute),
	159: uint16(alias_sym_enum_identifier),
	160: uint16(alias_sym_enum_member),
	161: uint16(alias_sym_extend_type),
	162: uint16(alias_sym_field_identifier),
	163: uint16(alias_sym_generic_identifier),
	164: uint16(alias_sym_implicit_generic_parameters),
	165: uint16(alias_sym_import_path),
	166: uint16(alias_sym_local_const),
	167: uint16(alias_sym_method_identifier),
	168: uint16(alias_sym_namespace),
	169: uint16(alias_sym_param_identifier),
	170: uint16(alias_sym_property),
	171: uint16(alias_sym_return_identifier),
	172: uint16(alias_sym_type_definition),
	173: uint16(alias_sym_type_identifier),
}

var ts_symbol_metadata = [174]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	2: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	52: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	57: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
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
	},
	61: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	63: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	67: {},
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	86: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	97: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	115: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	116: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	120: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	121: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	122: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	126: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	127: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	128: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	129: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	130: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	131: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
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
	145: {},
	146: {},
	147: {},
	148: {},
	149: {},
	150: {},
	151: {},
	152: {},
	153: {},
	154: {},
	155: {},
	156: {},
	157: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	158: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	159: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	160: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	161: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	162: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	163: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	164: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	165: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	166: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	167: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	168: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	169: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	170: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	171: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	172: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	173: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [35][13]TSSymbol{
	0: {},
	1: {
		1: uint16(alias_sym_import_path),
	},
	2: {
		1: uint16(alias_sym_type_identifier),
	},
	3: {
		0: uint16(alias_sym_type_definition),
		2: uint16(alias_sym_type_identifier),
	},
	4: {
		0: uint16(alias_sym_type_identifier),
	},
	5: {
		1: uint16(alias_sym_annotation_identifier),
	},
	6: {
		1: uint16(alias_sym_enum_identifier),
	},
	7: {
		0: uint16(alias_sym_generic_identifier),
	},
	8: {
		0: uint16(alias_sym_method_identifier),
	},
	9: {
		0: uint16(alias_sym_type_identifier),
		3: uint16(alias_sym_import_path),
	},
	10: {
		2: uint16(alias_sym_import_path),
	},
	11: {
		2: uint16(alias_sym_namespace),
	},
	12: {
		1: uint16(alias_sym_attribute),
	},
	13: {
		0: uint16(alias_sym_enum_member),
	},
	14: {
		1: uint16(alias_sym_generic_identifier),
	},
	15: {
		0: uint16(alias_sym_property),
	},
	16: {
		2: uint16(alias_sym_annotation_identifier),
	},
	17: {
		1: uint16(sym__identifier_no_period),
	},
	18: {
		2: uint16(alias_sym_type_identifier),
	},
	19: {
		1: uint16(alias_sym_import_path),
		5: uint16(alias_sym_namespace),
	},
	20: {
		1: uint16(alias_sym_local_const),
	},
	21: {
		1: uint16(alias_sym_implicit_generic_parameters),
	},
	22: {
		3: uint16(alias_sym_type_identifier),
	},
	23: {
		1: uint16(alias_sym_type_identifier),
		4: uint16(alias_sym_extend_type),
	},
	24: {
		0: uint16(alias_sym_field_identifier),
	},
	25: {
		1: uint16(alias_sym_annotation_identifier),
		3: uint16(alias_sym_param_identifier),
	},
	26: {
		1: uint16(alias_sym_type_identifier),
		5: uint16(alias_sym_extend_type),
	},
	27: {
		0: uint16(alias_sym_param_identifier),
	},
	28: {
		0: uint16(alias_sym_field_identifier),
		4: uint16(alias_sym_import_path),
	},
	29: {
		2: uint16(alias_sym_import_path),
		7: uint16(alias_sym_namespace),
	},
	30: {
		1: uint16(alias_sym_annotation_identifier),
		4: uint16(alias_sym_param_identifier),
	},
	31: {
		1: uint16(alias_sym_type_identifier),
		6: uint16(alias_sym_extend_type),
	},
	32: {
		0: uint16(alias_sym_return_identifier),
		2: uint16(alias_sym_type_identifier),
	},
	33: {
		1: uint16(alias_sym_property),
	},
	34: {
		1: uint16(alias_sym_return_identifier),
		3: uint16(alias_sym_type_identifier),
	},
}

var ts_non_terminal_alias_map = [18]uint16_t{
	0:  uint16(sym_field_type),
	1:  uint16(2),
	2:  uint16(sym_field_type),
	3:  uint16(alias_sym_generic_identifier),
	4:  uint16(sym_const_value),
	5:  uint16(2),
	6:  uint16(sym_const_value),
	7:  uint16(alias_sym_local_const),
	8:  uint16(sym_generic_parameters),
	9:  uint16(2),
	10: uint16(sym_generic_parameters),
	11: uint16(alias_sym_implicit_generic_parameters),
	12: uint16(sym_string),
	13: uint16(3),
	14: uint16(sym_string),
	15: uint16(alias_sym_import_path),
	16: uint16(alias_sym_namespace),
}

var ts_primary_state_ids = [553]TSStateId{
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
	35:  uint16(28),
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
	197: uint16(197),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(201),
	202: uint16(202),
	203: uint16(203),
	204: uint16(204),
	205: uint16(205),
	206: uint16(206),
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
	243: uint16(196),
	244: uint16(200),
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
	287: uint16(287),
	288: uint16(288),
	289: uint16(289),
	290: uint16(290),
	291: uint16(291),
	292: uint16(292),
	293: uint16(293),
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
	307: uint16(307),
	308: uint16(308),
	309: uint16(309),
	310: uint16(304),
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
	332: uint16(332),
	333: uint16(333),
	334: uint16(334),
	335: uint16(335),
	336: uint16(336),
	337: uint16(288),
	338: uint16(338),
	339: uint16(339),
	340: uint16(340),
	341: uint16(341),
	342: uint16(342),
	343: uint16(343),
	344: uint16(344),
	345: uint16(345),
	346: uint16(346),
	347: uint16(347),
	348: uint16(348),
	349: uint16(349),
	350: uint16(350),
	351: uint16(351),
	352: uint16(352),
	353: uint16(353),
	354: uint16(354),
	355: uint16(355),
	356: uint16(356),
	357: uint16(357),
	358: uint16(358),
	359: uint16(359),
	360: uint16(360),
	361: uint16(361),
	362: uint16(362),
	363: uint16(363),
	364: uint16(364),
	365: uint16(365),
	366: uint16(366),
	367: uint16(367),
	368: uint16(368),
	369: uint16(369),
	370: uint16(370),
	371: uint16(371),
	372: uint16(372),
	373: uint16(373),
	374: uint16(374),
	375: uint16(375),
	376: uint16(376),
	377: uint16(377),
	378: uint16(378),
	379: uint16(379),
	380: uint16(380),
	381: uint16(381),
	382: uint16(382),
	383: uint16(383),
	384: uint16(384),
	385: uint16(385),
	386: uint16(386),
	387: uint16(387),
	388: uint16(388),
	389: uint16(389),
	390: uint16(390),
	391: uint16(391),
	392: uint16(392),
	393: uint16(393),
	394: uint16(394),
	395: uint16(395),
	396: uint16(396),
	397: uint16(397),
	398: uint16(398),
	399: uint16(399),
	400: uint16(400),
	401: uint16(401),
	402: uint16(402),
	403: uint16(403),
	404: uint16(404),
	405: uint16(290),
	406: uint16(406),
	407: uint16(407),
	408: uint16(408),
	409: uint16(409),
	410: uint16(410),
	411: uint16(411),
	412: uint16(412),
	413: uint16(413),
	414: uint16(414),
	415: uint16(415),
	416: uint16(416),
	417: uint16(417),
	418: uint16(418),
	419: uint16(419),
	420: uint16(420),
	421: uint16(421),
	422: uint16(422),
	423: uint16(423),
	424: uint16(424),
	425: uint16(425),
	426: uint16(426),
	427: uint16(427),
	428: uint16(428),
	429: uint16(429),
	430: uint16(430),
	431: uint16(431),
	432: uint16(432),
	433: uint16(433),
	434: uint16(434),
	435: uint16(435),
	436: uint16(436),
	437: uint16(437),
	438: uint16(438),
	439: uint16(439),
	440: uint16(440),
	441: uint16(441),
	442: uint16(442),
	443: uint16(443),
	444: uint16(444),
	445: uint16(445),
	446: uint16(446),
	447: uint16(447),
	448: uint16(448),
	449: uint16(449),
	450: uint16(450),
	451: uint16(451),
	452: uint16(452),
	453: uint16(453),
	454: uint16(454),
	455: uint16(455),
	456: uint16(456),
	457: uint16(457),
	458: uint16(458),
	459: uint16(459),
	460: uint16(460),
	461: uint16(461),
	462: uint16(462),
	463: uint16(463),
	464: uint16(464),
	465: uint16(465),
	466: uint16(444),
	467: uint16(467),
	468: uint16(468),
	469: uint16(469),
	470: uint16(470),
	471: uint16(471),
	472: uint16(472),
	473: uint16(473),
	474: uint16(474),
	475: uint16(475),
	476: uint16(476),
	477: uint16(477),
	478: uint16(478),
	479: uint16(479),
	480: uint16(480),
	481: uint16(481),
	482: uint16(482),
	483: uint16(483),
	484: uint16(484),
	485: uint16(485),
	486: uint16(486),
	487: uint16(487),
	488: uint16(488),
	489: uint16(489),
	490: uint16(490),
	491: uint16(491),
	492: uint16(492),
	493: uint16(493),
	494: uint16(494),
	495: uint16(495),
	496: uint16(496),
	497: uint16(497),
	498: uint16(498),
	499: uint16(499),
	500: uint16(500),
	501: uint16(501),
	502: uint16(502),
	503: uint16(503),
	504: uint16(504),
	505: uint16(505),
	506: uint16(506),
	507: uint16(507),
	508: uint16(508),
	509: uint16(509),
	510: uint16(510),
	511: uint16(511),
	512: uint16(512),
	513: uint16(513),
	514: uint16(514),
	515: uint16(515),
	516: uint16(516),
	517: uint16(517),
	518: uint16(518),
	519: uint16(519),
	520: uint16(520),
	521: uint16(521),
	522: uint16(522),
	523: uint16(523),
	524: uint16(524),
	525: uint16(525),
	526: uint16(526),
	527: uint16(527),
	528: uint16(528),
	529: uint16(529),
	530: uint16(530),
	531: uint16(531),
	532: uint16(532),
	533: uint16(533),
	534: uint16(534),
	535: uint16(535),
	536: uint16(536),
	537: uint16(537),
	538: uint16(538),
	539: uint16(539),
	540: uint16(540),
	541: uint16(541),
	542: uint16(542),
	543: uint16(543),
	544: uint16(544),
	545: uint16(545),
	546: uint16(546),
	547: uint16(547),
	548: uint16(548),
	549: uint16(549),
	550: uint16(550),
	551: uint16(551),
	552: uint16(519),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i2, i3, i4, i5 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _, _ = eof, i, i1, i2, i3, i4, i5, lookahead, result, skip
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
			state = uint16(55)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(68)/libc.Uint64FromInt64(2)) {
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
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(150)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(139)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(151)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(153)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(72)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(73)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('\n') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('\r') {
			state = uint16(116)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(49)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(117)
			goto next_state
		}
		if lookahead == int32('"') || lookahead == int32('\'') || lookahead == int32('?') || lookahead == int32('\\') || lookahead == int32('a') || lookahead == int32('b') || lookahead == int32('f') || lookahead == int32('n') || lookahead == int32('r') || int32('t') <= lookahead && lookahead <= int32('v') {
			state = uint16(118)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32(' ') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(96)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(13):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32(']') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(131)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(120)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(132)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(134)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(72)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(13)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(14):
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(7)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(72)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(15):
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
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(71)
			goto next_state
		}
		if lookahead == int32('`') {
			state = uint16(105)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(158)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(170)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(173)
			goto next_state
		}
		if lookahead == int32('v') {
			state = uint16(172)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('"') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(106)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(5)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			state = uint16(107)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('"') {
			state = uint16(96)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('#') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(64)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(63)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(57)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(71)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('#') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('#') {
			state = uint16(177)
			goto next_state
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(21):
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			state = uint16(113)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || int32('_') <= lookahead && lookahead <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('#') {
			state = uint16(109)
			goto next_state
		}
		if lookahead == int32('\'') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(5)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			state = uint16(110)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(79)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(30)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(78)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(30)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('0') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('>') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('f') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('m') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('n') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('o') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('p') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('r') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('t') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('{') {
			state = uint16(47)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('}') {
			state = uint16(118)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('+') || lookahead == int32('-') {
			state = uint16(42)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(39):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(40):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(41):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(42):
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(43):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(44):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(45):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(46):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(47):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(48):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(49):
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(50):
		if eof != 0 {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(51):
		if eof != 0 {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(52):
		if eof != 0 {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(53):
		if eof != 0 {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(54):
		if eof != 0 {
			state = uint16(55)
			goto next_state
		}
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
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
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(53)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(73)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(54)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unique_id)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(45)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLARimport)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOLLAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_top_level_annotation_body_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('X') {
			state = uint16(43)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(95)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('X') || lookahead == int32('x') {
			state = uint16(43)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('0') {
			state = uint16(80)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		if int32('1') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('_') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(43)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('_') {
			state = uint16(41)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_number)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__normal_version)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(86)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__inline_version_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('E') || lookahead == int32('e') {
			state = uint16(37)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_float)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_true)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_false)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_0x)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('F') || int32('a') <= lookahead && lookahead <= int32('f') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_data_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_void)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_void)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_void)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_embed)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_embed)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_embed)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_double_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(108)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(106)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('#') && lookahead != int32('\\') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_double_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(106)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			state = uint16(107)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('#') && lookahead != int32('\\') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_double_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('"') && lookahead != int32('\\') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_single_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(111)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(109)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_single_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(109)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			state = uint16(110)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_single_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\'') && lookahead != int32('\\') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_block_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('#') {
			state = uint16(112)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(';') && lookahead != int32('\\') && lookahead != int32('`') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_block_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(112)
			goto next_state
		}
		if lookahead == int32('$') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(115)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			state = uint16(113)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(';') && lookahead != int32('\\') && lookahead != int32(']') && (lookahead < int32('_') || int32('z') < lookahead) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_block_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(114)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(';') && lookahead != int32('\\') && (lookahead < int32('_') || int32('z') < lookahead) {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_unescaped_block_string_fragment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32(';') && lookahead != int32('\\') && lookahead != int32('`') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__escape_sequence_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('7') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(130)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(133)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('b') {
			state = uint16(127)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(97)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('d') {
			state = uint16(100)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(89)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(92)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(124)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(76)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(123)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(136)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(122)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(128)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(76)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(129)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(137)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(126)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(125)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(157)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(149)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(152)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(146)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(99)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(91)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(94)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(143)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(84)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(142)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(155)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(147)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(84)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(156)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(145)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(157):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(157)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(168)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(171)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('b') {
			state = uint16(165)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(98)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(101)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(90)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(93)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(162)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(83)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(161)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(174)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(160)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(166)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(83)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(167)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(175)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(164)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(163)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__identifier_no_period)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('#') {
			state = uint16(177)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(177)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [34]uint16_t{
	0:  uint16('"'),
	1:  uint16(103),
	2:  uint16('#'),
	3:  uint16(177),
	4:  uint16('$'),
	5:  uint16(65),
	6:  uint16('\''),
	7:  uint16(104),
	8:  uint16('('),
	9:  uint16(62),
	10: uint16(')'),
	11: uint16(63),
	12: uint16('*'),
	13: uint16(69),
	14: uint16('+'),
	15: uint16(23),
	16: uint16(','),
	17: uint16(67),
	18: uint16('-'),
	19: uint16(24),
	20: uint16('.'),
	21: uint16(60),
	22: uint16('0'),
	23: uint16(77),
	24: uint16(':'),
	25: uint16(68),
	26: uint16(';'),
	27: uint16(57),
	28: uint16('='),
	29: uint16(58),
	30: uint16('@'),
	31: uint16(40),
	32: uint16('['),
	33: uint16(70),
}

var map_token1 = [32]uint16_t{
	0:  uint16('"'),
	1:  uint16(103),
	2:  uint16('#'),
	3:  uint16(177),
	4:  uint16('$'),
	5:  uint16(64),
	6:  uint16('\''),
	7:  uint16(104),
	8:  uint16('('),
	9:  uint16(62),
	10: uint16(')'),
	11: uint16(63),
	12: uint16('+'),
	13: uint16(23),
	14: uint16(','),
	15: uint16(67),
	16: uint16('-'),
	17: uint16(24),
	18: uint16('.'),
	19: uint16(39),
	20: uint16('0'),
	21: uint16(77),
	22: uint16(':'),
	23: uint16(68),
	24: uint16(';'),
	25: uint16(57),
	26: uint16('='),
	27: uint16(58),
	28: uint16('@'),
	29: uint16(26),
	30: uint16('['),
	31: uint16(70),
}

var map_token2 = [28]uint16_t{
	0:  uint16('"'),
	1:  uint16(103),
	2:  uint16('#'),
	3:  uint16(177),
	4:  uint16('$'),
	5:  uint16(64),
	6:  uint16('\''),
	7:  uint16(104),
	8:  uint16('('),
	9:  uint16(62),
	10: uint16(')'),
	11: uint16(63),
	12: uint16(','),
	13: uint16(67),
	14: uint16('-'),
	15: uint16(27),
	16: uint16('.'),
	17: uint16(59),
	18: uint16(':'),
	19: uint16(68),
	20: uint16(';'),
	21: uint16(57),
	22: uint16('='),
	23: uint16(58),
	24: uint16('@'),
	25: uint16(26),
	26: uint16('['),
	27: uint16(70),
}

var map_token3 = [22]uint16_t{
	0:  uint16('"'),
	1:  uint16(103),
	2:  uint16('#'),
	3:  uint16(177),
	4:  uint16('\''),
	5:  uint16(104),
	6:  uint16('('),
	7:  uint16(62),
	8:  uint16(')'),
	9:  uint16(63),
	10: uint16('+'),
	11: uint16(23),
	12: uint16(','),
	13: uint16(67),
	14: uint16('-'),
	15: uint16(25),
	16: uint16('.'),
	17: uint16(60),
	18: uint16('0'),
	19: uint16(77),
	20: uint16('['),
	21: uint16(70),
}

var map_token4 = [16]uint16_t{
	0:  uint16('#'),
	1:  uint16(112),
	2:  uint16('$'),
	3:  uint16(115),
	4:  uint16(')'),
	5:  uint16(115),
	6:  uint16(','),
	7:  uint16(115),
	8:  uint16(';'),
	9:  uint16(57),
	10: uint16('\\'),
	11: uint16(5),
	12: uint16(']'),
	13: uint16(115),
	14: uint16('`'),
	15: uint16(105),
}

var map_token5 = [26]uint16_t{
	0:  uint16('"'),
	1:  uint16(17),
	2:  uint16('#'),
	3:  uint16(177),
	4:  uint16('$'),
	5:  uint16(65),
	6:  uint16('('),
	7:  uint16(62),
	8:  uint16(')'),
	9:  uint16(63),
	10: uint16('*'),
	11: uint16(69),
	12: uint16(','),
	13: uint16(67),
	14: uint16('.'),
	15: uint16(59),
	16: uint16(':'),
	17: uint16(68),
	18: uint16(';'),
	19: uint16(57),
	20: uint16('='),
	21: uint16(58),
	22: uint16('@'),
	23: uint16(26),
	24: uint16('['),
	25: uint16(70),
}

func ts_lex_keywords(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
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
	switch libc.Int32FromUint16(state) {
	case 0:
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token6[i]) == lookahead {
				state = map_token6[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if lookahead == int32('\\') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(11)
			goto next_state
		}
		if lookahead == int32('c') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('e') {
			state = uint16(13)
			goto next_state
		}
		if lookahead == int32('f') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('g') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(16)
			goto next_state
		}
		if lookahead == int32('m') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('p') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(21)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') || lookahead == int32(0x200b) || lookahead == int32(0x2060) || lookahead == int32(0xfeff) {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('n') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('o') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('a') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('l') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('n') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('i') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('e') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('I') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('o') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead == int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('n') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('o') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('n') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('i') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('r') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('m') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('e') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('a') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('a') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('t') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('n') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('y') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('o') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('t') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('o') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('t') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('s') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('x') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('n') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('i') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('n') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('n') {
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
		if lookahead == int32('t') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('e') {
			state = uint16(59)
			goto next_state
		}
		if lookahead == int32('l') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('o') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('p') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('t') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('t') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('m') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('r') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('r') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('i') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead == int32('i') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead == int32('P') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('l') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('a') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('a') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('1') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(75)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(76)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('t') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('t') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('t') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('d') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('o') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('s') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('m') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('e') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('l') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('e') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('u') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('o') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('e') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('h') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('e') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('a') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('u') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('o') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('n') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('o') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Bool)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Data)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		if lookahead == int32('t') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('6') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('2') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('4') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Int8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_List)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		if lookahead == int32('1') {
			state = uint16(102)
			goto next_state
		}
		if lookahead == int32('3') {
			state = uint16(103)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(104)
			goto next_state
		}
		if lookahead == int32('8') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Void)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		if lookahead == int32('t') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('t') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('n') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('d') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_file)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		if lookahead == int32('p') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('r') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('r') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('o') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('s') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('m') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('c') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('n') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('g') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('i') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('3') {
			state = uint16(121)
			goto next_state
		}
		if lookahead == int32('6') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Int16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Int32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Int64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(102):
		if lookahead == int32('6') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('2') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('4') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UInt8)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		if lookahead == int32('a') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_const)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(108):
		if lookahead == int32('r') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('d') {
			state = uint16(128)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_field)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_group)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(112):
		if lookahead == int32('t') {
			state = uint16(129)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('f') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('d') {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('p') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_param)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(117):
		if lookahead == int32('t') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_union)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_using)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(120):
		if lookahead == int32('n') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('2') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('4') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UInt16)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UInt32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_UInt64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(126):
		if lookahead == int32('t') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('a') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('s') {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_import)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(130):
		if lookahead == int32('a') {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_method)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(132):
		if lookahead == int32('a') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_struct)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(134):
		if lookahead == int32('t') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Float32)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_Float64)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(137):
		if lookahead == int32('i') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead == int32('n') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_extends)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(140):
		if lookahead == int32('c') {
			state = uint16(145)
			goto next_state
		}
		return result
	case int32(141):
		if lookahead == int32('c') {
			state = uint16(146)
			goto next_state
		}
		return result
	case int32(142):
		if lookahead == int32('e') {
			state = uint16(147)
			goto next_state
		}
		return result
	case int32(143):
		if lookahead == int32('o') {
			state = uint16(148)
			goto next_state
		}
		return result
	case int32(144):
		if lookahead == int32('t') {
			state = uint16(149)
			goto next_state
		}
		return result
	case int32(145):
		if lookahead == int32('e') {
			state = uint16(150)
			goto next_state
		}
		return result
	case int32(146):
		if lookahead == int32('e') {
			state = uint16(151)
			goto next_state
		}
		return result
	case int32(147):
		if lookahead == int32('r') {
			state = uint16(152)
			goto next_state
		}
		return result
	case int32(148):
		if lookahead == int32('n') {
			state = uint16(153)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_enumerant)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_interface)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_namespace)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AnyPointer)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_annotation)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token6 = [18]uint16_t{
	0:  uint16('A'),
	1:  uint16(1),
	2:  uint16('B'),
	3:  uint16(2),
	4:  uint16('D'),
	5:  uint16(3),
	6:  uint16('F'),
	7:  uint16(4),
	8:  uint16('I'),
	9:  uint16(5),
	10: uint16('L'),
	11: uint16(6),
	12: uint16('T'),
	13: uint16(7),
	14: uint16('U'),
	15: uint16(8),
	16: uint16('V'),
	17: uint16(9),
}

var ts_lex_modes = [553]TSLexMode{
	0: {},
	1: {
		Flex_state: uint16(54),
	},
	2: {
		Flex_state: uint16(13),
	},
	3: {
		Flex_state: uint16(13),
	},
	4: {
		Flex_state: uint16(13),
	},
	5: {
		Flex_state: uint16(13),
	},
	6: {
		Flex_state: uint16(13),
	},
	7: {
		Flex_state: uint16(13),
	},
	8: {
		Flex_state: uint16(13),
	},
	9: {
		Flex_state: uint16(13),
	},
	10: {
		Flex_state: uint16(13),
	},
	11: {
		Flex_state: uint16(13),
	},
	12: {
		Flex_state: uint16(15),
	},
	13: {
		Flex_state: uint16(13),
	},
	14: {
		Flex_state: uint16(15),
	},
	15: {
		Flex_state: uint16(15),
	},
	16: {
		Flex_state: uint16(15),
	},
	17: {
		Flex_state: uint16(13),
	},
	18: {
		Flex_state: uint16(15),
	},
	19: {
		Flex_state: uint16(15),
	},
	20: {
		Flex_state: uint16(15),
	},
	21: {
		Flex_state: uint16(15),
	},
	22: {
		Flex_state: uint16(15),
	},
	23: {
		Flex_state: uint16(15),
	},
	24: {
		Flex_state: uint16(15),
	},
	25: {
		Flex_state: uint16(15),
	},
	26: {
		Flex_state: uint16(15),
	},
	27: {
		Flex_state: uint16(15),
	},
	28: {
		Flex_state: uint16(15),
	},
	29: {
		Flex_state: uint16(15),
	},
	30: {
		Flex_state: uint16(15),
	},
	31: {
		Flex_state: uint16(15),
	},
	32: {
		Flex_state: uint16(15),
	},
	33: {
		Flex_state: uint16(15),
	},
	34: {
		Flex_state: uint16(15),
	},
	35: {
		Flex_state: uint16(15),
	},
	36: {
		Flex_state: uint16(15),
	},
	37: {
		Flex_state: uint16(15),
	},
	38: {
		Flex_state: uint16(15),
	},
	39: {
		Flex_state: uint16(15),
	},
	40: {
		Flex_state: uint16(15),
	},
	41: {
		Flex_state: uint16(54),
	},
	42: {
		Flex_state: uint16(54),
	},
	43: {
		Flex_state: uint16(54),
	},
	44: {
		Flex_state: uint16(54),
	},
	45: {
		Flex_state: uint16(54),
	},
	46: {
		Flex_state: uint16(54),
	},
	47: {
		Flex_state: uint16(54),
	},
	48: {
		Flex_state: uint16(54),
	},
	49: {
		Flex_state: uint16(54),
	},
	50: {
		Flex_state: uint16(54),
	},
	51: {
		Flex_state: uint16(54),
	},
	52: {
		Flex_state: uint16(54),
	},
	53: {
		Flex_state: uint16(54),
	},
	54: {
		Flex_state: uint16(54),
	},
	55: {
		Flex_state: uint16(54),
	},
	56: {
		Flex_state: uint16(54),
	},
	57: {
		Flex_state: uint16(54),
	},
	58: {
		Flex_state: uint16(54),
	},
	59: {
		Flex_state: uint16(54),
	},
	60: {
		Flex_state: uint16(54),
	},
	61: {
		Flex_state: uint16(54),
	},
	62: {
		Flex_state: uint16(54),
	},
	63: {
		Flex_state: uint16(54),
	},
	64: {
		Flex_state: uint16(54),
	},
	65: {
		Flex_state: uint16(54),
	},
	66: {
		Flex_state: uint16(54),
	},
	67: {
		Flex_state: uint16(54),
	},
	68: {
		Flex_state: uint16(54),
	},
	69: {
		Flex_state: uint16(54),
	},
	70: {
		Flex_state: uint16(54),
	},
	71: {
		Flex_state: uint16(54),
	},
	72: {
		Flex_state: uint16(54),
	},
	73: {
		Flex_state: uint16(54),
	},
	74: {
		Flex_state: uint16(54),
	},
	75: {
		Flex_state: uint16(54),
	},
	76: {
		Flex_state: uint16(54),
	},
	77: {
		Flex_state: uint16(54),
	},
	78: {
		Flex_state: uint16(54),
	},
	79: {
		Flex_state: uint16(54),
	},
	80: {
		Flex_state: uint16(54),
	},
	81: {
		Flex_state: uint16(54),
	},
	82: {
		Flex_state: uint16(54),
	},
	83: {
		Flex_state: uint16(54),
	},
	84: {
		Flex_state: uint16(54),
	},
	85: {
		Flex_state: uint16(54),
	},
	86: {
		Flex_state: uint16(54),
	},
	87: {
		Flex_state: uint16(54),
	},
	88: {
		Flex_state: uint16(54),
	},
	89: {
		Flex_state: uint16(54),
	},
	90: {
		Flex_state: uint16(54),
	},
	91: {
		Flex_state: uint16(54),
	},
	92: {
		Flex_state: uint16(54),
	},
	93: {
		Flex_state: uint16(54),
	},
	94: {
		Flex_state: uint16(54),
	},
	95: {
		Flex_state: uint16(54),
	},
	96: {
		Flex_state: uint16(54),
	},
	97: {
		Flex_state: uint16(54),
	},
	98: {
		Flex_state: uint16(54),
	},
	99: {
		Flex_state: uint16(54),
	},
	100: {
		Flex_state: uint16(54),
	},
	101: {
		Flex_state: uint16(54),
	},
	102: {
		Flex_state: uint16(54),
	},
	103: {
		Flex_state: uint16(54),
	},
	104: {
		Flex_state: uint16(54),
	},
	105: {
		Flex_state: uint16(54),
	},
	106: {
		Flex_state: uint16(54),
	},
	107: {
		Flex_state: uint16(21),
	},
	108: {
		Flex_state: uint16(21),
	},
	109: {
		Flex_state: uint16(21),
	},
	110: {
		Flex_state: uint16(14),
	},
	111: {
		Flex_state: uint16(54),
	},
	112: {
		Flex_state: uint16(54),
	},
	113: {
		Flex_state: uint16(54),
	},
	114: {
		Flex_state: uint16(54),
	},
	115: {
		Flex_state: uint16(54),
	},
	116: {
		Flex_state: uint16(14),
	},
	117: {
		Flex_state: uint16(54),
	},
	118: {
		Flex_state: uint16(54),
	},
	119: {
		Flex_state: uint16(54),
	},
	120: {
		Flex_state: uint16(54),
	},
	121: {
		Flex_state: uint16(54),
	},
	122: {
		Flex_state: uint16(54),
	},
	123: {
		Flex_state: uint16(54),
	},
	124: {
		Flex_state: uint16(54),
	},
	125: {
		Flex_state: uint16(54),
	},
	126: {
		Flex_state: uint16(54),
	},
	127: {
		Flex_state: uint16(54),
	},
	128: {
		Flex_state: uint16(54),
	},
	129: {
		Flex_state: uint16(54),
	},
	130: {
		Flex_state: uint16(54),
	},
	131: {
		Flex_state: uint16(54),
	},
	132: {
		Flex_state: uint16(54),
	},
	133: {
		Flex_state: uint16(54),
	},
	134: {
		Flex_state: uint16(54),
	},
	135: {
		Flex_state: uint16(54),
	},
	136: {
		Flex_state: uint16(54),
	},
	137: {
		Flex_state: uint16(54),
	},
	138: {
		Flex_state: uint16(54),
	},
	139: {
		Flex_state: uint16(54),
	},
	140: {
		Flex_state: uint16(54),
	},
	141: {
		Flex_state: uint16(54),
	},
	142: {
		Flex_state: uint16(14),
	},
	143: {
		Flex_state: uint16(54),
	},
	144: {
		Flex_state: uint16(54),
	},
	145: {
		Flex_state: uint16(54),
	},
	146: {
		Flex_state: uint16(14),
	},
	147: {
		Flex_state: uint16(14),
	},
	148: {
		Flex_state: uint16(54),
	},
	149: {
		Flex_state: uint16(14),
	},
	150: {
		Flex_state: uint16(54),
	},
	151: {
		Flex_state: uint16(14),
	},
	152: {
		Flex_state: uint16(14),
	},
	153: {
		Flex_state: uint16(54),
	},
	154: {
		Flex_state: uint16(14),
	},
	155: {
		Flex_state: uint16(14),
	},
	156: {
		Flex_state: uint16(54),
	},
	157: {
		Flex_state: uint16(54),
	},
	158: {
		Flex_state: uint16(14),
	},
	159: {
		Flex_state: uint16(54),
	},
	160: {
		Flex_state: uint16(54),
	},
	161: {
		Flex_state: uint16(54),
	},
	162: {
		Flex_state: uint16(54),
	},
	163: {
		Flex_state: uint16(54),
	},
	164: {
		Flex_state: uint16(54),
	},
	165: {
		Flex_state: uint16(54),
	},
	166: {
		Flex_state: uint16(54),
	},
	167: {
		Flex_state: uint16(14),
	},
	168: {
		Flex_state: uint16(14),
	},
	169: {
		Flex_state: uint16(54),
	},
	170: {
		Flex_state: uint16(54),
	},
	171: {
		Flex_state: uint16(54),
	},
	172: {
		Flex_state: uint16(14),
	},
	173: {
		Flex_state: uint16(54),
	},
	174: {
		Flex_state: uint16(54),
	},
	175: {
		Flex_state: uint16(54),
	},
	176: {
		Flex_state: uint16(14),
	},
	177: {
		Flex_state: uint16(54),
	},
	178: {
		Flex_state: uint16(54),
	},
	179: {
		Flex_state: uint16(54),
	},
	180: {
		Flex_state: uint16(54),
	},
	181: {
		Flex_state: uint16(54),
	},
	182: {
		Flex_state: uint16(54),
	},
	183: {
		Flex_state: uint16(54),
	},
	184: {
		Flex_state: uint16(54),
	},
	185: {
		Flex_state: uint16(54),
	},
	186: {
		Flex_state: uint16(54),
	},
	187: {
		Flex_state: uint16(54),
	},
	188: {
		Flex_state: uint16(14),
	},
	189: {
		Flex_state: uint16(14),
	},
	190: {
		Flex_state: uint16(14),
	},
	191: {
		Flex_state: uint16(14),
	},
	192: {
		Flex_state: uint16(14),
	},
	193: {
		Flex_state: uint16(14),
	},
	194: {
		Flex_state: uint16(14),
	},
	195: {
		Flex_state: uint16(13),
	},
	196: {
		Flex_state: uint16(18),
	},
	197: {
		Flex_state: uint16(14),
	},
	198: {
		Flex_state: uint16(14),
	},
	199: {
		Flex_state: uint16(54),
	},
	200: {
		Flex_state: uint16(18),
	},
	201: {
		Flex_state: uint16(13),
	},
	202: {
		Flex_state: uint16(14),
	},
	203: {
		Flex_state: uint16(22),
	},
	204: {
		Flex_state: uint16(14),
	},
	205: {
		Flex_state: uint16(13),
	},
	206: {
		Flex_state: uint16(13),
	},
	207: {
		Flex_state: uint16(16),
	},
	208: {
		Flex_state: uint16(14),
	},
	209: {
		Flex_state: uint16(13),
	},
	210: {
		Flex_state: uint16(14),
	},
	211: {
		Flex_state: uint16(14),
	},
	212: {
		Flex_state: uint16(14),
	},
	213: {
		Flex_state: uint16(13),
	},
	214: {
		Flex_state: uint16(14),
	},
	215: {
		Flex_state: uint16(14),
	},
	216: {
		Flex_state: uint16(14),
	},
	217: {
		Flex_state: uint16(13),
	},
	218: {
		Flex_state: uint16(13),
	},
	219: {
		Flex_state: uint16(13),
	},
	220: {
		Flex_state: uint16(16),
	},
	221: {
		Flex_state: uint16(13),
	},
	222: {
		Flex_state: uint16(16),
	},
	223: {
		Flex_state: uint16(22),
	},
	224: {
		Flex_state: uint16(13),
	},
	225: {
		Flex_state: uint16(14),
	},
	226: {
		Flex_state: uint16(13),
	},
	227: {
		Flex_state: uint16(13),
	},
	228: {
		Flex_state: uint16(13),
	},
	229: {
		Flex_state: uint16(13),
	},
	230: {
		Flex_state: uint16(13),
	},
	231: {
		Flex_state: uint16(14),
	},
	232: {
		Flex_state: uint16(13),
	},
	233: {
		Flex_state: uint16(14),
	},
	234: {
		Flex_state: uint16(13),
	},
	235: {
		Flex_state: uint16(13),
	},
	236: {
		Flex_state: uint16(22),
	},
	237: {},
	238: {
		Flex_state: uint16(13),
	},
	239: {
		Flex_state: uint16(54),
	},
	240: {},
	241: {
		Flex_state: uint16(54),
	},
	242: {
		Flex_state: uint16(54),
	},
	243: {
		Flex_state: uint16(54),
	},
	244: {
		Flex_state: uint16(54),
	},
	245: {
		Flex_state: uint16(54),
	},
	246: {
		Flex_state: uint16(54),
	},
	247: {
		Flex_state: uint16(54),
	},
	248: {
		Flex_state: uint16(54),
	},
	249: {
		Flex_state: uint16(54),
	},
	250: {},
	251: {
		Flex_state: uint16(54),
	},
	252: {},
	253: {
		Flex_state: uint16(54),
	},
	254: {
		Flex_state: uint16(54),
	},
	255: {
		Flex_state: uint16(54),
	},
	256: {
		Flex_state: uint16(54),
	},
	257: {
		Flex_state: uint16(54),
	},
	258: {
		Flex_state: uint16(54),
	},
	259: {
		Flex_state: uint16(13),
	},
	260: {
		Flex_state: uint16(54),
	},
	261: {
		Flex_state: uint16(54),
	},
	262: {},
	263: {
		Flex_state: uint16(54),
	},
	264: {
		Flex_state: uint16(54),
	},
	265: {
		Flex_state: uint16(54),
	},
	266: {},
	267: {
		Flex_state: uint16(54),
	},
	268: {
		Flex_state: uint16(13),
	},
	269: {
		Flex_state: uint16(54),
	},
	270: {
		Flex_state: uint16(13),
	},
	271: {},
	272: {
		Flex_state: uint16(13),
	},
	273: {},
	274: {
		Flex_state: uint16(54),
	},
	275: {},
	276: {
		Flex_state: uint16(13),
	},
	277: {
		Flex_state: uint16(54),
	},
	278: {
		Flex_state: uint16(13),
	},
	279: {
		Flex_state: uint16(54),
	},
	280: {
		Flex_state: uint16(54),
	},
	281: {},
	282: {},
	283: {
		Flex_state: uint16(13),
	},
	284: {},
	285: {
		Flex_state: uint16(54),
	},
	286: {},
	287: {
		Flex_state: uint16(13),
	},
	288: {},
	289: {
		Flex_state: uint16(54),
	},
	290: {},
	291: {
		Flex_state: uint16(14),
	},
	292: {},
	293: {
		Flex_state: uint16(13),
	},
	294: {},
	295: {
		Flex_state: uint16(14),
	},
	296: {
		Flex_state: uint16(13),
	},
	297: {
		Flex_state: uint16(54),
	},
	298: {},
	299: {
		Flex_state: uint16(13),
	},
	300: {
		Flex_state: uint16(13),
	},
	301: {
		Flex_state: uint16(54),
	},
	302: {},
	303: {
		Flex_state: uint16(13),
	},
	304: {},
	305: {
		Flex_state: uint16(54),
	},
	306: {
		Flex_state: uint16(13),
	},
	307: {
		Flex_state: uint16(54),
	},
	308: {},
	309: {
		Flex_state: uint16(13),
	},
	310: {},
	311: {},
	312: {
		Flex_state: uint16(13),
	},
	313: {
		Flex_state: uint16(13),
	},
	314: {
		Flex_state: uint16(13),
	},
	315: {},
	316: {},
	317: {},
	318: {},
	319: {
		Flex_state: uint16(13),
	},
	320: {
		Flex_state: uint16(13),
	},
	321: {},
	322: {
		Flex_state: uint16(13),
	},
	323: {},
	324: {
		Flex_state: uint16(13),
	},
	325: {},
	326: {},
	327: {
		Flex_state: uint16(54),
	},
	328: {
		Flex_state: uint16(13),
	},
	329: {
		Flex_state: uint16(54),
	},
	330: {},
	331: {},
	332: {},
	333: {},
	334: {
		Flex_state: uint16(54),
	},
	335: {},
	336: {},
	337: {},
	338: {
		Flex_state: uint16(13),
	},
	339: {},
	340: {},
	341: {
		Flex_state: uint16(54),
	},
	342: {},
	343: {},
	344: {},
	345: {},
	346: {},
	347: {
		Flex_state: uint16(54),
	},
	348: {
		Flex_state: uint16(13),
	},
	349: {
		Flex_state: uint16(13),
	},
	350: {},
	351: {
		Flex_state: uint16(13),
	},
	352: {
		Flex_state: uint16(13),
	},
	353: {
		Flex_state: uint16(13),
	},
	354: {
		Flex_state: uint16(13),
	},
	355: {
		Flex_state: uint16(13),
	},
	356: {},
	357: {
		Flex_state: uint16(13),
	},
	358: {
		Flex_state: uint16(13),
	},
	359: {
		Flex_state: uint16(54),
	},
	360: {},
	361: {
		Flex_state: uint16(13),
	},
	362: {
		Flex_state: uint16(54),
	},
	363: {
		Flex_state: uint16(13),
	},
	364: {
		Flex_state: uint16(13),
	},
	365: {
		Flex_state: uint16(13),
	},
	366: {
		Flex_state: uint16(13),
	},
	367: {},
	368: {
		Flex_state: uint16(54),
	},
	369: {
		Flex_state: uint16(13),
	},
	370: {
		Flex_state: uint16(13),
	},
	371: {
		Flex_state: uint16(13),
	},
	372: {
		Flex_state: uint16(54),
	},
	373: {
		Flex_state: uint16(13),
	},
	374: {
		Flex_state: uint16(13),
	},
	375: {
		Flex_state: uint16(13),
	},
	376: {
		Flex_state: uint16(54),
	},
	377: {
		Flex_state: uint16(13),
	},
	378: {
		Flex_state: uint16(54),
	},
	379: {},
	380: {
		Flex_state: uint16(13),
	},
	381: {
		Flex_state: uint16(13),
	},
	382: {
		Flex_state: uint16(54),
	},
	383: {},
	384: {},
	385: {},
	386: {
		Flex_state: uint16(13),
	},
	387: {
		Flex_state: uint16(54),
	},
	388: {},
	389: {},
	390: {
		Flex_state: uint16(54),
	},
	391: {
		Flex_state: uint16(13),
	},
	392: {},
	393: {
		Flex_state: uint16(13),
	},
	394: {
		Flex_state: uint16(13),
	},
	395: {
		Flex_state: uint16(13),
	},
	396: {},
	397: {
		Flex_state: uint16(13),
	},
	398: {},
	399: {
		Flex_state: uint16(13),
	},
	400: {},
	401: {},
	402: {
		Flex_state: uint16(13),
	},
	403: {
		Flex_state: uint16(13),
	},
	404: {
		Flex_state: uint16(13),
	},
	405: {},
	406: {},
	407: {
		Flex_state: uint16(54),
	},
	408: {},
	409: {},
	410: {
		Flex_state: uint16(54),
	},
	411: {
		Flex_state: uint16(13),
	},
	412: {
		Flex_state: uint16(13),
	},
	413: {
		Flex_state: uint16(13),
	},
	414: {
		Flex_state: uint16(13),
	},
	415: {
		Flex_state: uint16(13),
	},
	416: {
		Flex_state: uint16(13),
	},
	417: {},
	418: {
		Flex_state: uint16(13),
	},
	419: {
		Flex_state: uint16(13),
	},
	420: {
		Flex_state: uint16(13),
	},
	421: {
		Flex_state: uint16(13),
	},
	422: {
		Flex_state: uint16(13),
	},
	423: {
		Flex_state: uint16(13),
	},
	424: {
		Flex_state: uint16(13),
	},
	425: {
		Flex_state: uint16(13),
	},
	426: {
		Flex_state: uint16(13),
	},
	427: {
		Flex_state: uint16(13),
	},
	428: {
		Flex_state: uint16(54),
	},
	429: {},
	430: {
		Flex_state: uint16(13),
	},
	431: {
		Flex_state: uint16(13),
	},
	432: {
		Flex_state: uint16(13),
	},
	433: {},
	434: {
		Flex_state: uint16(54),
	},
	435: {
		Flex_state: uint16(13),
	},
	436: {
		Flex_state: uint16(13),
	},
	437: {
		Flex_state: uint16(13),
	},
	438: {
		Flex_state: uint16(54),
	},
	439: {
		Flex_state: uint16(13),
	},
	440: {},
	441: {
		Flex_state: uint16(13),
	},
	442: {
		Flex_state: uint16(54),
	},
	443: {
		Flex_state: uint16(20),
	},
	444: {
		Flex_state: uint16(18),
	},
	445: {
		Flex_state: uint16(18),
	},
	446: {
		Flex_state: uint16(54),
	},
	447: {},
	448: {},
	449: {},
	450: {},
	451: {
		Flex_state: uint16(54),
	},
	452: {
		Flex_state: uint16(54),
	},
	453: {
		Flex_state: uint16(54),
	},
	454: {},
	455: {
		Flex_state: uint16(54),
	},
	456: {},
	457: {
		Flex_state: uint16(54),
	},
	458: {
		Flex_state: uint16(54),
	},
	459: {
		Flex_state: uint16(54),
	},
	460: {
		Flex_state: uint16(13),
	},
	461: {},
	462: {
		Flex_state: uint16(54),
	},
	463: {
		Flex_state: uint16(13),
	},
	464: {
		Flex_state: uint16(54),
	},
	465: {
		Flex_state: uint16(13),
	},
	466: {
		Flex_state: uint16(18),
	},
	467: {},
	468: {
		Flex_state: uint16(13),
	},
	469: {},
	470: {
		Flex_state: uint16(54),
	},
	471: {},
	472: {},
	473: {
		Flex_state: uint16(18),
	},
	474: {},
	475: {
		Flex_state: uint16(54),
	},
	476: {},
	477: {},
	478: {},
	479: {},
	480: {},
	481: {
		Flex_state: uint16(54),
	},
	482: {},
	483: {},
	484: {},
	485: {},
	486: {},
	487: {},
	488: {},
	489: {
		Flex_state: uint16(54),
	},
	490: {},
	491: {},
	492: {},
	493: {
		Flex_state: uint16(54),
	},
	494: {},
	495: {},
	496: {},
	497: {},
	498: {},
	499: {
		Flex_state: uint16(54),
	},
	500: {},
	501: {},
	502: {},
	503: {},
	504: {
		Flex_state: uint16(54),
	},
	505: {
		Flex_state: uint16(54),
	},
	506: {},
	507: {},
	508: {},
	509: {},
	510: {
		Flex_state: uint16(54),
	},
	511: {},
	512: {
		Flex_state: uint16(54),
	},
	513: {},
	514: {
		Flex_state: uint16(54),
	},
	515: {},
	516: {
		Flex_state: uint16(54),
	},
	517: {},
	518: {
		Flex_state: uint16(54),
	},
	519: {},
	520: {},
	521: {},
	522: {},
	523: {},
	524: {},
	525: {},
	526: {},
	527: {},
	528: {},
	529: {},
	530: {},
	531: {},
	532: {},
	533: {},
	534: {},
	535: {},
	536: {
		Flex_state: uint16(54),
	},
	537: {},
	538: {},
	539: {},
	540: {},
	541: {},
	542: {
		Flex_state: uint16(54),
	},
	543: {
		Flex_state: uint16(54),
	},
	544: {},
	545: {
		Flex_state: uint16(54),
	},
	546: {
		Flex_state: uint16(54),
	},
	547: {
		Flex_state: uint16(54),
	},
	548: {
		Flex_state: uint16(18),
	},
	549: {},
	550: {},
	551: {
		Flex_state: uint16(54),
	},
	552: {},
}

var ts_parse_table = [2][157]uint16_t{
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
		55: uint16(1),
		56: uint16(1),
		57: uint16(1),
		59: uint16(1),
		60: uint16(1),
		61: uint16(1),
		62: uint16(1),
		63: uint16(1),
		70: uint16(3),
	},
	1: {
		0:   uint16(5),
		2:   uint16(7),
		4:   uint16(9),
		8:   uint16(11),
		12:  uint16(13),
		15:  uint16(15),
		18:  uint16(17),
		24:  uint16(19),
		26:  uint16(21),
		27:  uint16(23),
		70:  uint16(3),
		71:  uint16(540),
		72:  uint16(51),
		73:  uint16(161),
		74:  uint16(161),
		77:  uint16(161),
		78:  uint16(161),
		80:  uint16(161),
		86:  uint16(161),
		87:  uint16(160),
		89:  uint16(160),
		99:  uint16(160),
		112: uint16(160),
		132: uint16(51),
	},
}

var ts_small_parse_table = [10821]uint16_t{
	0:     uint16(25),
	1:     uint16(3),
	2:     uint16(1),
	3:     uint16(sym_comment),
	4:     uint16(25),
	5:     uint16(1),
	6:     uint16(sym_identifier),
	7:     uint16(27),
	8:     uint16(1),
	9:     uint16(anon_sym_LPAREN),
	10:    uint16(29),
	11:    uint16(1),
	12:    uint16(anon_sym_RPAREN),
	13:    uint16(31),
	14:    uint16(1),
	15:    uint16(anon_sym_LBRACK),
	16:    uint16(35),
	17:    uint16(1),
	18:    uint16(anon_sym_List),
	19:    uint16(41),
	20:    uint16(1),
	21:    uint16(anon_sym_0x),
	22:    uint16(43),
	23:    uint16(1),
	24:    uint16(anon_sym_embed),
	25:    uint16(45),
	26:    uint16(1),
	27:    uint16(anon_sym_DQUOTE),
	28:    uint16(47),
	29:    uint16(1),
	30:    uint16(anon_sym_SQUOTE),
	31:    uint16(49),
	32:    uint16(1),
	33:    uint16(anon_sym_BQUOTE),
	34:    uint16(51),
	35:    uint16(1),
	36:    uint16(sym__identifier_no_period),
	37:    uint16(147),
	38:    uint16(1),
	39:    uint16(sym_string),
	40:    uint16(190),
	41:    uint16(1),
	42:    uint16(aux_sym_block_text_repeat2),
	43:    uint16(294),
	44:    uint16(1),
	45:    uint16(sym_field_type),
	46:    uint16(388),
	47:    uint16(1),
	48:    uint16(sym_const_value),
	49:    uint16(389),
	50:    uint16(1),
	51:    uint16(sym__annotation_array_def),
	52:    uint16(466),
	53:    uint16(1),
	54:    uint16(aux_sym__internal_const_identifier_repeat1),
	55:    uint16(478),
	56:    uint16(1),
	57:    uint16(sym_generic_parameters),
	58:    uint16(488),
	59:    uint16(1),
	60:    uint16(sym_annotation_array),
	61:    uint16(39),
	62:    uint16(2),
	63:    uint16(anon_sym_true),
	64:    uint16(anon_sym_false),
	65:    uint16(37),
	66:    uint16(3),
	67:    uint16(sym_number),
	68:    uint16(sym_float),
	69:    uint16(sym_void),
	70:    uint16(238),
	71:    uint16(3),
	72:    uint16(sym_primitive_type),
	73:    uint16(sym_list_type),
	74:    uint16(sym_custom_type),
	75:    uint16(210),
	76:    uint16(8),
	77:    uint16(sym_boolean),
	78:    uint16(sym_data),
	79:    uint16(sym_const_list),
	80:    uint16(sym_struct_shorthand),
	81:    uint16(sym__internal_const_identifier),
	82:    uint16(sym_embedded_file),
	83:    uint16(sym_concatenated_string),
	84:    uint16(sym_block_text),
	85:    uint16(33),
	86:    uint16(15),
	87:    uint16(anon_sym_AnyPointer),
	88:    uint16(anon_sym_Bool),
	89:    uint16(anon_sym_Int8),
	90:    uint16(anon_sym_Int16),
	91:    uint16(anon_sym_Int32),
	92:    uint16(anon_sym_Int64),
	93:    uint16(anon_sym_UInt8),
	94:    uint16(anon_sym_UInt16),
	95:    uint16(anon_sym_UInt32),
	96:    uint16(anon_sym_UInt64),
	97:    uint16(anon_sym_Float32),
	98:    uint16(anon_sym_Float64),
	99:    uint16(anon_sym_Text),
	100:   uint16(anon_sym_Data),
	101:   uint16(anon_sym_Void),
	102:   uint16(21),
	103:   uint16(3),
	104:   uint16(1),
	105:   uint16(sym_comment),
	106:   uint16(27),
	107:   uint16(1),
	108:   uint16(anon_sym_LPAREN),
	109:   uint16(31),
	110:   uint16(1),
	111:   uint16(anon_sym_LBRACK),
	112:   uint16(41),
	113:   uint16(1),
	114:   uint16(anon_sym_0x),
	115:   uint16(43),
	116:   uint16(1),
	117:   uint16(anon_sym_embed),
	118:   uint16(45),
	119:   uint16(1),
	120:   uint16(anon_sym_DQUOTE),
	121:   uint16(47),
	122:   uint16(1),
	123:   uint16(anon_sym_SQUOTE),
	124:   uint16(49),
	125:   uint16(1),
	126:   uint16(anon_sym_BQUOTE),
	127:   uint16(51),
	128:   uint16(1),
	129:   uint16(sym__identifier_no_period),
	130:   uint16(53),
	131:   uint16(1),
	132:   uint16(sym_identifier),
	133:   uint16(55),
	134:   uint16(1),
	135:   uint16(anon_sym_RPAREN),
	136:   uint16(147),
	137:   uint16(1),
	138:   uint16(sym_string),
	139:   uint16(190),
	140:   uint16(1),
	141:   uint16(aux_sym_block_text_repeat2),
	142:   uint16(281),
	143:   uint16(1),
	144:   uint16(sym_const_value),
	145:   uint16(343),
	146:   uint16(1),
	147:   uint16(sym_annotation_array),
	148:   uint16(389),
	149:   uint16(1),
	150:   uint16(sym__annotation_array_def),
	151:   uint16(434),
	152:   uint16(1),
	153:   uint16(aux_sym_struct_shorthand_repeat1),
	154:   uint16(466),
	155:   uint16(1),
	156:   uint16(aux_sym__internal_const_identifier_repeat1),
	157:   uint16(39),
	158:   uint16(2),
	159:   uint16(anon_sym_true),
	160:   uint16(anon_sym_false),
	161:   uint16(37),
	162:   uint16(3),
	163:   uint16(sym_number),
	164:   uint16(sym_float),
	165:   uint16(sym_void),
	166:   uint16(210),
	167:   uint16(8),
	168:   uint16(sym_boolean),
	169:   uint16(sym_data),
	170:   uint16(sym_const_list),
	171:   uint16(sym_struct_shorthand),
	172:   uint16(sym__internal_const_identifier),
	173:   uint16(sym_embedded_file),
	174:   uint16(sym_concatenated_string),
	175:   uint16(sym_block_text),
	176:   uint16(20),
	177:   uint16(3),
	178:   uint16(1),
	179:   uint16(sym_comment),
	180:   uint16(27),
	181:   uint16(1),
	182:   uint16(anon_sym_LPAREN),
	183:   uint16(31),
	184:   uint16(1),
	185:   uint16(anon_sym_LBRACK),
	186:   uint16(41),
	187:   uint16(1),
	188:   uint16(anon_sym_0x),
	189:   uint16(43),
	190:   uint16(1),
	191:   uint16(anon_sym_embed),
	192:   uint16(45),
	193:   uint16(1),
	194:   uint16(anon_sym_DQUOTE),
	195:   uint16(47),
	196:   uint16(1),
	197:   uint16(anon_sym_SQUOTE),
	198:   uint16(49),
	199:   uint16(1),
	200:   uint16(anon_sym_BQUOTE),
	201:   uint16(51),
	202:   uint16(1),
	203:   uint16(sym__identifier_no_period),
	204:   uint16(57),
	205:   uint16(1),
	206:   uint16(sym_identifier),
	207:   uint16(59),
	208:   uint16(1),
	209:   uint16(anon_sym_RPAREN),
	210:   uint16(147),
	211:   uint16(1),
	212:   uint16(sym_string),
	213:   uint16(190),
	214:   uint16(1),
	215:   uint16(aux_sym_block_text_repeat2),
	216:   uint16(388),
	217:   uint16(1),
	218:   uint16(sym_const_value),
	219:   uint16(389),
	220:   uint16(1),
	221:   uint16(sym__annotation_array_def),
	222:   uint16(466),
	223:   uint16(1),
	224:   uint16(aux_sym__internal_const_identifier_repeat1),
	225:   uint16(528),
	226:   uint16(1),
	227:   uint16(sym_annotation_array),
	228:   uint16(39),
	229:   uint16(2),
	230:   uint16(anon_sym_true),
	231:   uint16(anon_sym_false),
	232:   uint16(37),
	233:   uint16(3),
	234:   uint16(sym_number),
	235:   uint16(sym_float),
	236:   uint16(sym_void),
	237:   uint16(210),
	238:   uint16(8),
	239:   uint16(sym_boolean),
	240:   uint16(sym_data),
	241:   uint16(sym_const_list),
	242:   uint16(sym_struct_shorthand),
	243:   uint16(sym__internal_const_identifier),
	244:   uint16(sym_embedded_file),
	245:   uint16(sym_concatenated_string),
	246:   uint16(sym_block_text),
	247:   uint16(20),
	248:   uint16(3),
	249:   uint16(1),
	250:   uint16(sym_comment),
	251:   uint16(27),
	252:   uint16(1),
	253:   uint16(anon_sym_LPAREN),
	254:   uint16(31),
	255:   uint16(1),
	256:   uint16(anon_sym_LBRACK),
	257:   uint16(41),
	258:   uint16(1),
	259:   uint16(anon_sym_0x),
	260:   uint16(43),
	261:   uint16(1),
	262:   uint16(anon_sym_embed),
	263:   uint16(45),
	264:   uint16(1),
	265:   uint16(anon_sym_DQUOTE),
	266:   uint16(47),
	267:   uint16(1),
	268:   uint16(anon_sym_SQUOTE),
	269:   uint16(49),
	270:   uint16(1),
	271:   uint16(anon_sym_BQUOTE),
	272:   uint16(51),
	273:   uint16(1),
	274:   uint16(sym__identifier_no_period),
	275:   uint16(57),
	276:   uint16(1),
	277:   uint16(sym_identifier),
	278:   uint16(61),
	279:   uint16(1),
	280:   uint16(anon_sym_RPAREN),
	281:   uint16(147),
	282:   uint16(1),
	283:   uint16(sym_string),
	284:   uint16(190),
	285:   uint16(1),
	286:   uint16(aux_sym_block_text_repeat2),
	287:   uint16(388),
	288:   uint16(1),
	289:   uint16(sym_const_value),
	290:   uint16(389),
	291:   uint16(1),
	292:   uint16(sym__annotation_array_def),
	293:   uint16(466),
	294:   uint16(1),
	295:   uint16(aux_sym__internal_const_identifier_repeat1),
	296:   uint16(508),
	297:   uint16(1),
	298:   uint16(sym_annotation_array),
	299:   uint16(39),
	300:   uint16(2),
	301:   uint16(anon_sym_true),
	302:   uint16(anon_sym_false),
	303:   uint16(37),
	304:   uint16(3),
	305:   uint16(sym_number),
	306:   uint16(sym_float),
	307:   uint16(sym_void),
	308:   uint16(210),
	309:   uint16(8),
	310:   uint16(sym_boolean),
	311:   uint16(sym_data),
	312:   uint16(sym_const_list),
	313:   uint16(sym_struct_shorthand),
	314:   uint16(sym__internal_const_identifier),
	315:   uint16(sym_embedded_file),
	316:   uint16(sym_concatenated_string),
	317:   uint16(sym_block_text),
	318:   uint16(20),
	319:   uint16(3),
	320:   uint16(1),
	321:   uint16(sym_comment),
	322:   uint16(27),
	323:   uint16(1),
	324:   uint16(anon_sym_LPAREN),
	325:   uint16(31),
	326:   uint16(1),
	327:   uint16(anon_sym_LBRACK),
	328:   uint16(41),
	329:   uint16(1),
	330:   uint16(anon_sym_0x),
	331:   uint16(43),
	332:   uint16(1),
	333:   uint16(anon_sym_embed),
	334:   uint16(45),
	335:   uint16(1),
	336:   uint16(anon_sym_DQUOTE),
	337:   uint16(47),
	338:   uint16(1),
	339:   uint16(anon_sym_SQUOTE),
	340:   uint16(49),
	341:   uint16(1),
	342:   uint16(anon_sym_BQUOTE),
	343:   uint16(51),
	344:   uint16(1),
	345:   uint16(sym__identifier_no_period),
	346:   uint16(57),
	347:   uint16(1),
	348:   uint16(sym_identifier),
	349:   uint16(63),
	350:   uint16(1),
	351:   uint16(anon_sym_RPAREN),
	352:   uint16(147),
	353:   uint16(1),
	354:   uint16(sym_string),
	355:   uint16(190),
	356:   uint16(1),
	357:   uint16(aux_sym_block_text_repeat2),
	358:   uint16(388),
	359:   uint16(1),
	360:   uint16(sym_const_value),
	361:   uint16(389),
	362:   uint16(1),
	363:   uint16(sym__annotation_array_def),
	364:   uint16(466),
	365:   uint16(1),
	366:   uint16(aux_sym__internal_const_identifier_repeat1),
	367:   uint16(509),
	368:   uint16(1),
	369:   uint16(sym_annotation_array),
	370:   uint16(39),
	371:   uint16(2),
	372:   uint16(anon_sym_true),
	373:   uint16(anon_sym_false),
	374:   uint16(37),
	375:   uint16(3),
	376:   uint16(sym_number),
	377:   uint16(sym_float),
	378:   uint16(sym_void),
	379:   uint16(210),
	380:   uint16(8),
	381:   uint16(sym_boolean),
	382:   uint16(sym_data),
	383:   uint16(sym_const_list),
	384:   uint16(sym_struct_shorthand),
	385:   uint16(sym__internal_const_identifier),
	386:   uint16(sym_embedded_file),
	387:   uint16(sym_concatenated_string),
	388:   uint16(sym_block_text),
	389:   uint16(20),
	390:   uint16(3),
	391:   uint16(1),
	392:   uint16(sym_comment),
	393:   uint16(27),
	394:   uint16(1),
	395:   uint16(anon_sym_LPAREN),
	396:   uint16(31),
	397:   uint16(1),
	398:   uint16(anon_sym_LBRACK),
	399:   uint16(41),
	400:   uint16(1),
	401:   uint16(anon_sym_0x),
	402:   uint16(43),
	403:   uint16(1),
	404:   uint16(anon_sym_embed),
	405:   uint16(45),
	406:   uint16(1),
	407:   uint16(anon_sym_DQUOTE),
	408:   uint16(47),
	409:   uint16(1),
	410:   uint16(anon_sym_SQUOTE),
	411:   uint16(49),
	412:   uint16(1),
	413:   uint16(anon_sym_BQUOTE),
	414:   uint16(51),
	415:   uint16(1),
	416:   uint16(sym__identifier_no_period),
	417:   uint16(57),
	418:   uint16(1),
	419:   uint16(sym_identifier),
	420:   uint16(65),
	421:   uint16(1),
	422:   uint16(anon_sym_RPAREN),
	423:   uint16(147),
	424:   uint16(1),
	425:   uint16(sym_string),
	426:   uint16(190),
	427:   uint16(1),
	428:   uint16(aux_sym_block_text_repeat2),
	429:   uint16(388),
	430:   uint16(1),
	431:   uint16(sym_const_value),
	432:   uint16(389),
	433:   uint16(1),
	434:   uint16(sym__annotation_array_def),
	435:   uint16(466),
	436:   uint16(1),
	437:   uint16(aux_sym__internal_const_identifier_repeat1),
	438:   uint16(486),
	439:   uint16(1),
	440:   uint16(sym_annotation_array),
	441:   uint16(39),
	442:   uint16(2),
	443:   uint16(anon_sym_true),
	444:   uint16(anon_sym_false),
	445:   uint16(37),
	446:   uint16(3),
	447:   uint16(sym_number),
	448:   uint16(sym_float),
	449:   uint16(sym_void),
	450:   uint16(210),
	451:   uint16(8),
	452:   uint16(sym_boolean),
	453:   uint16(sym_data),
	454:   uint16(sym_const_list),
	455:   uint16(sym_struct_shorthand),
	456:   uint16(sym__internal_const_identifier),
	457:   uint16(sym_embedded_file),
	458:   uint16(sym_concatenated_string),
	459:   uint16(sym_block_text),
	460:   uint16(20),
	461:   uint16(3),
	462:   uint16(1),
	463:   uint16(sym_comment),
	464:   uint16(27),
	465:   uint16(1),
	466:   uint16(anon_sym_LPAREN),
	467:   uint16(31),
	468:   uint16(1),
	469:   uint16(anon_sym_LBRACK),
	470:   uint16(41),
	471:   uint16(1),
	472:   uint16(anon_sym_0x),
	473:   uint16(43),
	474:   uint16(1),
	475:   uint16(anon_sym_embed),
	476:   uint16(45),
	477:   uint16(1),
	478:   uint16(anon_sym_DQUOTE),
	479:   uint16(47),
	480:   uint16(1),
	481:   uint16(anon_sym_SQUOTE),
	482:   uint16(49),
	483:   uint16(1),
	484:   uint16(anon_sym_BQUOTE),
	485:   uint16(51),
	486:   uint16(1),
	487:   uint16(sym__identifier_no_period),
	488:   uint16(57),
	489:   uint16(1),
	490:   uint16(sym_identifier),
	491:   uint16(67),
	492:   uint16(1),
	493:   uint16(anon_sym_RPAREN),
	494:   uint16(147),
	495:   uint16(1),
	496:   uint16(sym_string),
	497:   uint16(190),
	498:   uint16(1),
	499:   uint16(aux_sym_block_text_repeat2),
	500:   uint16(388),
	501:   uint16(1),
	502:   uint16(sym_const_value),
	503:   uint16(389),
	504:   uint16(1),
	505:   uint16(sym__annotation_array_def),
	506:   uint16(466),
	507:   uint16(1),
	508:   uint16(aux_sym__internal_const_identifier_repeat1),
	509:   uint16(511),
	510:   uint16(1),
	511:   uint16(sym_annotation_array),
	512:   uint16(39),
	513:   uint16(2),
	514:   uint16(anon_sym_true),
	515:   uint16(anon_sym_false),
	516:   uint16(37),
	517:   uint16(3),
	518:   uint16(sym_number),
	519:   uint16(sym_float),
	520:   uint16(sym_void),
	521:   uint16(210),
	522:   uint16(8),
	523:   uint16(sym_boolean),
	524:   uint16(sym_data),
	525:   uint16(sym_const_list),
	526:   uint16(sym_struct_shorthand),
	527:   uint16(sym__internal_const_identifier),
	528:   uint16(sym_embedded_file),
	529:   uint16(sym_concatenated_string),
	530:   uint16(sym_block_text),
	531:   uint16(20),
	532:   uint16(3),
	533:   uint16(1),
	534:   uint16(sym_comment),
	535:   uint16(27),
	536:   uint16(1),
	537:   uint16(anon_sym_LPAREN),
	538:   uint16(31),
	539:   uint16(1),
	540:   uint16(anon_sym_LBRACK),
	541:   uint16(41),
	542:   uint16(1),
	543:   uint16(anon_sym_0x),
	544:   uint16(43),
	545:   uint16(1),
	546:   uint16(anon_sym_embed),
	547:   uint16(45),
	548:   uint16(1),
	549:   uint16(anon_sym_DQUOTE),
	550:   uint16(47),
	551:   uint16(1),
	552:   uint16(anon_sym_SQUOTE),
	553:   uint16(49),
	554:   uint16(1),
	555:   uint16(anon_sym_BQUOTE),
	556:   uint16(51),
	557:   uint16(1),
	558:   uint16(sym__identifier_no_period),
	559:   uint16(57),
	560:   uint16(1),
	561:   uint16(sym_identifier),
	562:   uint16(69),
	563:   uint16(1),
	564:   uint16(anon_sym_RPAREN),
	565:   uint16(147),
	566:   uint16(1),
	567:   uint16(sym_string),
	568:   uint16(190),
	569:   uint16(1),
	570:   uint16(aux_sym_block_text_repeat2),
	571:   uint16(388),
	572:   uint16(1),
	573:   uint16(sym_const_value),
	574:   uint16(389),
	575:   uint16(1),
	576:   uint16(sym__annotation_array_def),
	577:   uint16(466),
	578:   uint16(1),
	579:   uint16(aux_sym__internal_const_identifier_repeat1),
	580:   uint16(526),
	581:   uint16(1),
	582:   uint16(sym_annotation_array),
	583:   uint16(39),
	584:   uint16(2),
	585:   uint16(anon_sym_true),
	586:   uint16(anon_sym_false),
	587:   uint16(37),
	588:   uint16(3),
	589:   uint16(sym_number),
	590:   uint16(sym_float),
	591:   uint16(sym_void),
	592:   uint16(210),
	593:   uint16(8),
	594:   uint16(sym_boolean),
	595:   uint16(sym_data),
	596:   uint16(sym_const_list),
	597:   uint16(sym_struct_shorthand),
	598:   uint16(sym__internal_const_identifier),
	599:   uint16(sym_embedded_file),
	600:   uint16(sym_concatenated_string),
	601:   uint16(sym_block_text),
	602:   uint16(19),
	603:   uint16(3),
	604:   uint16(1),
	605:   uint16(sym_comment),
	606:   uint16(27),
	607:   uint16(1),
	608:   uint16(anon_sym_LPAREN),
	609:   uint16(31),
	610:   uint16(1),
	611:   uint16(anon_sym_LBRACK),
	612:   uint16(41),
	613:   uint16(1),
	614:   uint16(anon_sym_0x),
	615:   uint16(43),
	616:   uint16(1),
	617:   uint16(anon_sym_embed),
	618:   uint16(45),
	619:   uint16(1),
	620:   uint16(anon_sym_DQUOTE),
	621:   uint16(47),
	622:   uint16(1),
	623:   uint16(anon_sym_SQUOTE),
	624:   uint16(49),
	625:   uint16(1),
	626:   uint16(anon_sym_BQUOTE),
	627:   uint16(51),
	628:   uint16(1),
	629:   uint16(sym__identifier_no_period),
	630:   uint16(71),
	631:   uint16(1),
	632:   uint16(sym_identifier),
	633:   uint16(73),
	634:   uint16(1),
	635:   uint16(anon_sym_RBRACK),
	636:   uint16(147),
	637:   uint16(1),
	638:   uint16(sym_string),
	639:   uint16(190),
	640:   uint16(1),
	641:   uint16(aux_sym_block_text_repeat2),
	642:   uint16(339),
	643:   uint16(1),
	644:   uint16(sym_annotation_array),
	645:   uint16(466),
	646:   uint16(1),
	647:   uint16(aux_sym__internal_const_identifier_repeat1),
	648:   uint16(39),
	649:   uint16(2),
	650:   uint16(anon_sym_true),
	651:   uint16(anon_sym_false),
	652:   uint16(389),
	653:   uint16(2),
	654:   uint16(sym__annotation_array_def),
	655:   uint16(sym_const_value),
	656:   uint16(37),
	657:   uint16(3),
	658:   uint16(sym_number),
	659:   uint16(sym_float),
	660:   uint16(sym_void),
	661:   uint16(210),
	662:   uint16(8),
	663:   uint16(sym_boolean),
	664:   uint16(sym_data),
	665:   uint16(sym_const_list),
	666:   uint16(sym_struct_shorthand),
	667:   uint16(sym__internal_const_identifier),
	668:   uint16(sym_embedded_file),
	669:   uint16(sym_concatenated_string),
	670:   uint16(sym_block_text),
	671:   uint16(20),
	672:   uint16(3),
	673:   uint16(1),
	674:   uint16(sym_comment),
	675:   uint16(27),
	676:   uint16(1),
	677:   uint16(anon_sym_LPAREN),
	678:   uint16(31),
	679:   uint16(1),
	680:   uint16(anon_sym_LBRACK),
	681:   uint16(41),
	682:   uint16(1),
	683:   uint16(anon_sym_0x),
	684:   uint16(43),
	685:   uint16(1),
	686:   uint16(anon_sym_embed),
	687:   uint16(45),
	688:   uint16(1),
	689:   uint16(anon_sym_DQUOTE),
	690:   uint16(47),
	691:   uint16(1),
	692:   uint16(anon_sym_SQUOTE),
	693:   uint16(49),
	694:   uint16(1),
	695:   uint16(anon_sym_BQUOTE),
	696:   uint16(51),
	697:   uint16(1),
	698:   uint16(sym__identifier_no_period),
	699:   uint16(59),
	700:   uint16(1),
	701:   uint16(anon_sym_RPAREN),
	702:   uint16(75),
	703:   uint16(1),
	704:   uint16(sym_identifier),
	705:   uint16(147),
	706:   uint16(1),
	707:   uint16(sym_string),
	708:   uint16(190),
	709:   uint16(1),
	710:   uint16(aux_sym_block_text_repeat2),
	711:   uint16(388),
	712:   uint16(1),
	713:   uint16(sym_const_value),
	714:   uint16(389),
	715:   uint16(1),
	716:   uint16(sym__annotation_array_def),
	717:   uint16(466),
	718:   uint16(1),
	719:   uint16(aux_sym__internal_const_identifier_repeat1),
	720:   uint16(528),
	721:   uint16(1),
	722:   uint16(sym_annotation_array),
	723:   uint16(39),
	724:   uint16(2),
	725:   uint16(anon_sym_true),
	726:   uint16(anon_sym_false),
	727:   uint16(37),
	728:   uint16(3),
	729:   uint16(sym_number),
	730:   uint16(sym_float),
	731:   uint16(sym_void),
	732:   uint16(210),
	733:   uint16(8),
	734:   uint16(sym_boolean),
	735:   uint16(sym_data),
	736:   uint16(sym_const_list),
	737:   uint16(sym_struct_shorthand),
	738:   uint16(sym__internal_const_identifier),
	739:   uint16(sym_embedded_file),
	740:   uint16(sym_concatenated_string),
	741:   uint16(sym_block_text),
	742:   uint16(17),
	743:   uint16(3),
	744:   uint16(1),
	745:   uint16(sym_comment),
	746:   uint16(41),
	747:   uint16(1),
	748:   uint16(anon_sym_0x),
	749:   uint16(43),
	750:   uint16(1),
	751:   uint16(anon_sym_embed),
	752:   uint16(45),
	753:   uint16(1),
	754:   uint16(anon_sym_DQUOTE),
	755:   uint16(47),
	756:   uint16(1),
	757:   uint16(anon_sym_SQUOTE),
	758:   uint16(49),
	759:   uint16(1),
	760:   uint16(anon_sym_BQUOTE),
	761:   uint16(51),
	762:   uint16(1),
	763:   uint16(sym__identifier_no_period),
	764:   uint16(77),
	765:   uint16(1),
	766:   uint16(anon_sym_LPAREN),
	767:   uint16(81),
	768:   uint16(1),
	769:   uint16(anon_sym_LBRACK),
	770:   uint16(147),
	771:   uint16(1),
	772:   uint16(sym_string),
	773:   uint16(190),
	774:   uint16(1),
	775:   uint16(aux_sym_block_text_repeat2),
	776:   uint16(367),
	777:   uint16(1),
	778:   uint16(sym_const_value),
	779:   uint16(466),
	780:   uint16(1),
	781:   uint16(aux_sym__internal_const_identifier_repeat1),
	782:   uint16(39),
	783:   uint16(2),
	784:   uint16(anon_sym_true),
	785:   uint16(anon_sym_false),
	786:   uint16(37),
	787:   uint16(3),
	788:   uint16(sym_number),
	789:   uint16(sym_float),
	790:   uint16(sym_void),
	791:   uint16(79),
	792:   uint16(3),
	793:   uint16(anon_sym_RPAREN),
	794:   uint16(anon_sym_COMMA),
	795:   uint16(anon_sym_RBRACK),
	796:   uint16(210),
	797:   uint16(8),
	798:   uint16(sym_boolean),
	799:   uint16(sym_data),
	800:   uint16(sym_const_list),
	801:   uint16(sym_struct_shorthand),
	802:   uint16(sym__internal_const_identifier),
	803:   uint16(sym_embedded_file),
	804:   uint16(sym_concatenated_string),
	805:   uint16(sym_block_text),
	806:   uint16(19),
	807:   uint16(3),
	808:   uint16(1),
	809:   uint16(sym_comment),
	810:   uint16(27),
	811:   uint16(1),
	812:   uint16(anon_sym_LPAREN),
	813:   uint16(31),
	814:   uint16(1),
	815:   uint16(anon_sym_LBRACK),
	816:   uint16(41),
	817:   uint16(1),
	818:   uint16(anon_sym_0x),
	819:   uint16(43),
	820:   uint16(1),
	821:   uint16(anon_sym_embed),
	822:   uint16(45),
	823:   uint16(1),
	824:   uint16(anon_sym_DQUOTE),
	825:   uint16(47),
	826:   uint16(1),
	827:   uint16(anon_sym_SQUOTE),
	828:   uint16(49),
	829:   uint16(1),
	830:   uint16(anon_sym_BQUOTE),
	831:   uint16(51),
	832:   uint16(1),
	833:   uint16(sym__identifier_no_period),
	834:   uint16(71),
	835:   uint16(1),
	836:   uint16(sym_identifier),
	837:   uint16(147),
	838:   uint16(1),
	839:   uint16(sym_string),
	840:   uint16(190),
	841:   uint16(1),
	842:   uint16(aux_sym_block_text_repeat2),
	843:   uint16(281),
	844:   uint16(1),
	845:   uint16(sym_const_value),
	846:   uint16(389),
	847:   uint16(1),
	848:   uint16(sym__annotation_array_def),
	849:   uint16(400),
	850:   uint16(1),
	851:   uint16(sym_annotation_array),
	852:   uint16(466),
	853:   uint16(1),
	854:   uint16(aux_sym__internal_const_identifier_repeat1),
	855:   uint16(39),
	856:   uint16(2),
	857:   uint16(anon_sym_true),
	858:   uint16(anon_sym_false),
	859:   uint16(37),
	860:   uint16(3),
	861:   uint16(sym_number),
	862:   uint16(sym_float),
	863:   uint16(sym_void),
	864:   uint16(210),
	865:   uint16(8),
	866:   uint16(sym_boolean),
	867:   uint16(sym_data),
	868:   uint16(sym_const_list),
	869:   uint16(sym_struct_shorthand),
	870:   uint16(sym__internal_const_identifier),
	871:   uint16(sym_embedded_file),
	872:   uint16(sym_concatenated_string),
	873:   uint16(sym_block_text),
	874:   uint16(17),
	875:   uint16(3),
	876:   uint16(1),
	877:   uint16(sym_comment),
	878:   uint16(41),
	879:   uint16(1),
	880:   uint16(anon_sym_0x),
	881:   uint16(43),
	882:   uint16(1),
	883:   uint16(anon_sym_embed),
	884:   uint16(45),
	885:   uint16(1),
	886:   uint16(anon_sym_DQUOTE),
	887:   uint16(47),
	888:   uint16(1),
	889:   uint16(anon_sym_SQUOTE),
	890:   uint16(49),
	891:   uint16(1),
	892:   uint16(anon_sym_BQUOTE),
	893:   uint16(51),
	894:   uint16(1),
	895:   uint16(sym__identifier_no_period),
	896:   uint16(77),
	897:   uint16(1),
	898:   uint16(anon_sym_LPAREN),
	899:   uint16(81),
	900:   uint16(1),
	901:   uint16(anon_sym_LBRACK),
	902:   uint16(147),
	903:   uint16(1),
	904:   uint16(sym_string),
	905:   uint16(190),
	906:   uint16(1),
	907:   uint16(aux_sym_block_text_repeat2),
	908:   uint16(367),
	909:   uint16(1),
	910:   uint16(sym_const_value),
	911:   uint16(466),
	912:   uint16(1),
	913:   uint16(aux_sym__internal_const_identifier_repeat1),
	914:   uint16(39),
	915:   uint16(2),
	916:   uint16(anon_sym_true),
	917:   uint16(anon_sym_false),
	918:   uint16(37),
	919:   uint16(3),
	920:   uint16(sym_number),
	921:   uint16(sym_float),
	922:   uint16(sym_void),
	923:   uint16(83),
	924:   uint16(3),
	925:   uint16(anon_sym_RPAREN),
	926:   uint16(anon_sym_COMMA),
	927:   uint16(anon_sym_RBRACK),
	928:   uint16(210),
	929:   uint16(8),
	930:   uint16(sym_boolean),
	931:   uint16(sym_data),
	932:   uint16(sym_const_list),
	933:   uint16(sym_struct_shorthand),
	934:   uint16(sym__internal_const_identifier),
	935:   uint16(sym_embedded_file),
	936:   uint16(sym_concatenated_string),
	937:   uint16(sym_block_text),
	938:   uint16(18),
	939:   uint16(3),
	940:   uint16(1),
	941:   uint16(sym_comment),
	942:   uint16(41),
	943:   uint16(1),
	944:   uint16(anon_sym_0x),
	945:   uint16(43),
	946:   uint16(1),
	947:   uint16(anon_sym_embed),
	948:   uint16(45),
	949:   uint16(1),
	950:   uint16(anon_sym_DQUOTE),
	951:   uint16(47),
	952:   uint16(1),
	953:   uint16(anon_sym_SQUOTE),
	954:   uint16(49),
	955:   uint16(1),
	956:   uint16(anon_sym_BQUOTE),
	957:   uint16(77),
	958:   uint16(1),
	959:   uint16(anon_sym_LPAREN),
	960:   uint16(81),
	961:   uint16(1),
	962:   uint16(anon_sym_LBRACK),
	963:   uint16(85),
	964:   uint16(1),
	965:   uint16(anon_sym_DOT),
	966:   uint16(87),
	967:   uint16(1),
	968:   uint16(sym__identifier_no_period),
	969:   uint16(147),
	970:   uint16(1),
	971:   uint16(sym_string),
	972:   uint16(190),
	973:   uint16(1),
	974:   uint16(aux_sym_block_text_repeat2),
	975:   uint16(301),
	976:   uint16(1),
	977:   uint16(sym_const_value),
	978:   uint16(368),
	979:   uint16(1),
	980:   uint16(sym__same_scope_const_value),
	981:   uint16(444),
	982:   uint16(1),
	983:   uint16(aux_sym__internal_const_identifier_repeat1),
	984:   uint16(39),
	985:   uint16(2),
	986:   uint16(anon_sym_true),
	987:   uint16(anon_sym_false),
	988:   uint16(37),
	989:   uint16(3),
	990:   uint16(sym_number),
	991:   uint16(sym_float),
	992:   uint16(sym_void),
	993:   uint16(210),
	994:   uint16(8),
	995:   uint16(sym_boolean),
	996:   uint16(sym_data),
	997:   uint16(sym_const_list),
	998:   uint16(sym_struct_shorthand),
	999:   uint16(sym__internal_const_identifier),
	1000:  uint16(sym_embedded_file),
	1001:  uint16(sym_concatenated_string),
	1002:  uint16(sym_block_text),
	1003:  uint16(17),
	1004:  uint16(3),
	1005:  uint16(1),
	1006:  uint16(sym_comment),
	1007:  uint16(41),
	1008:  uint16(1),
	1009:  uint16(anon_sym_0x),
	1010:  uint16(43),
	1011:  uint16(1),
	1012:  uint16(anon_sym_embed),
	1013:  uint16(45),
	1014:  uint16(1),
	1015:  uint16(anon_sym_DQUOTE),
	1016:  uint16(47),
	1017:  uint16(1),
	1018:  uint16(anon_sym_SQUOTE),
	1019:  uint16(49),
	1020:  uint16(1),
	1021:  uint16(anon_sym_BQUOTE),
	1022:  uint16(77),
	1023:  uint16(1),
	1024:  uint16(anon_sym_LPAREN),
	1025:  uint16(81),
	1026:  uint16(1),
	1027:  uint16(anon_sym_LBRACK),
	1028:  uint16(85),
	1029:  uint16(1),
	1030:  uint16(anon_sym_DOT),
	1031:  uint16(87),
	1032:  uint16(1),
	1033:  uint16(sym__identifier_no_period),
	1034:  uint16(147),
	1035:  uint16(1),
	1036:  uint16(sym_string),
	1037:  uint16(190),
	1038:  uint16(1),
	1039:  uint16(aux_sym_block_text_repeat2),
	1040:  uint16(444),
	1041:  uint16(1),
	1042:  uint16(aux_sym__internal_const_identifier_repeat1),
	1043:  uint16(39),
	1044:  uint16(2),
	1045:  uint16(anon_sym_true),
	1046:  uint16(anon_sym_false),
	1047:  uint16(368),
	1048:  uint16(2),
	1049:  uint16(sym_const_value),
	1050:  uint16(sym__same_scope_const_value),
	1051:  uint16(37),
	1052:  uint16(3),
	1053:  uint16(sym_number),
	1054:  uint16(sym_float),
	1055:  uint16(sym_void),
	1056:  uint16(210),
	1057:  uint16(8),
	1058:  uint16(sym_boolean),
	1059:  uint16(sym_data),
	1060:  uint16(sym_const_list),
	1061:  uint16(sym_struct_shorthand),
	1062:  uint16(sym__internal_const_identifier),
	1063:  uint16(sym_embedded_file),
	1064:  uint16(sym_concatenated_string),
	1065:  uint16(sym_block_text),
	1066:  uint16(17),
	1067:  uint16(3),
	1068:  uint16(1),
	1069:  uint16(sym_comment),
	1070:  uint16(41),
	1071:  uint16(1),
	1072:  uint16(anon_sym_0x),
	1073:  uint16(43),
	1074:  uint16(1),
	1075:  uint16(anon_sym_embed),
	1076:  uint16(45),
	1077:  uint16(1),
	1078:  uint16(anon_sym_DQUOTE),
	1079:  uint16(47),
	1080:  uint16(1),
	1081:  uint16(anon_sym_SQUOTE),
	1082:  uint16(49),
	1083:  uint16(1),
	1084:  uint16(anon_sym_BQUOTE),
	1085:  uint16(51),
	1086:  uint16(1),
	1087:  uint16(sym__identifier_no_period),
	1088:  uint16(77),
	1089:  uint16(1),
	1090:  uint16(anon_sym_LPAREN),
	1091:  uint16(81),
	1092:  uint16(1),
	1093:  uint16(anon_sym_LBRACK),
	1094:  uint16(89),
	1095:  uint16(1),
	1096:  uint16(sym_identifier),
	1097:  uint16(91),
	1098:  uint16(1),
	1099:  uint16(anon_sym_RPAREN),
	1100:  uint16(190),
	1101:  uint16(1),
	1102:  uint16(aux_sym_block_text_repeat2),
	1103:  uint16(266),
	1104:  uint16(1),
	1105:  uint16(sym_string),
	1106:  uint16(466),
	1107:  uint16(1),
	1108:  uint16(aux_sym__internal_const_identifier_repeat1),
	1109:  uint16(39),
	1110:  uint16(2),
	1111:  uint16(anon_sym_true),
	1112:  uint16(anon_sym_false),
	1113:  uint16(93),
	1114:  uint16(3),
	1115:  uint16(sym_number),
	1116:  uint16(sym_float),
	1117:  uint16(sym_void),
	1118:  uint16(490),
	1119:  uint16(8),
	1120:  uint16(sym_boolean),
	1121:  uint16(sym_data),
	1122:  uint16(sym_const_list),
	1123:  uint16(sym_struct_shorthand),
	1124:  uint16(sym__internal_const_identifier),
	1125:  uint16(sym_embedded_file),
	1126:  uint16(sym_concatenated_string),
	1127:  uint16(sym_block_text),
	1128:  uint16(17),
	1129:  uint16(3),
	1130:  uint16(1),
	1131:  uint16(sym_comment),
	1132:  uint16(41),
	1133:  uint16(1),
	1134:  uint16(anon_sym_0x),
	1135:  uint16(43),
	1136:  uint16(1),
	1137:  uint16(anon_sym_embed),
	1138:  uint16(45),
	1139:  uint16(1),
	1140:  uint16(anon_sym_DQUOTE),
	1141:  uint16(47),
	1142:  uint16(1),
	1143:  uint16(anon_sym_SQUOTE),
	1144:  uint16(49),
	1145:  uint16(1),
	1146:  uint16(anon_sym_BQUOTE),
	1147:  uint16(51),
	1148:  uint16(1),
	1149:  uint16(sym__identifier_no_period),
	1150:  uint16(77),
	1151:  uint16(1),
	1152:  uint16(anon_sym_LPAREN),
	1153:  uint16(81),
	1154:  uint16(1),
	1155:  uint16(anon_sym_LBRACK),
	1156:  uint16(95),
	1157:  uint16(1),
	1158:  uint16(anon_sym_RBRACK),
	1159:  uint16(147),
	1160:  uint16(1),
	1161:  uint16(sym_string),
	1162:  uint16(190),
	1163:  uint16(1),
	1164:  uint16(aux_sym_block_text_repeat2),
	1165:  uint16(367),
	1166:  uint16(1),
	1167:  uint16(sym_const_value),
	1168:  uint16(466),
	1169:  uint16(1),
	1170:  uint16(aux_sym__internal_const_identifier_repeat1),
	1171:  uint16(39),
	1172:  uint16(2),
	1173:  uint16(anon_sym_true),
	1174:  uint16(anon_sym_false),
	1175:  uint16(37),
	1176:  uint16(3),
	1177:  uint16(sym_number),
	1178:  uint16(sym_float),
	1179:  uint16(sym_void),
	1180:  uint16(210),
	1181:  uint16(8),
	1182:  uint16(sym_boolean),
	1183:  uint16(sym_data),
	1184:  uint16(sym_const_list),
	1185:  uint16(sym_struct_shorthand),
	1186:  uint16(sym__internal_const_identifier),
	1187:  uint16(sym_embedded_file),
	1188:  uint16(sym_concatenated_string),
	1189:  uint16(sym_block_text),
	1190:  uint16(17),
	1191:  uint16(3),
	1192:  uint16(1),
	1193:  uint16(sym_comment),
	1194:  uint16(41),
	1195:  uint16(1),
	1196:  uint16(anon_sym_0x),
	1197:  uint16(43),
	1198:  uint16(1),
	1199:  uint16(anon_sym_embed),
	1200:  uint16(45),
	1201:  uint16(1),
	1202:  uint16(anon_sym_DQUOTE),
	1203:  uint16(47),
	1204:  uint16(1),
	1205:  uint16(anon_sym_SQUOTE),
	1206:  uint16(49),
	1207:  uint16(1),
	1208:  uint16(anon_sym_BQUOTE),
	1209:  uint16(51),
	1210:  uint16(1),
	1211:  uint16(sym__identifier_no_period),
	1212:  uint16(73),
	1213:  uint16(1),
	1214:  uint16(anon_sym_RBRACK),
	1215:  uint16(77),
	1216:  uint16(1),
	1217:  uint16(anon_sym_LPAREN),
	1218:  uint16(81),
	1219:  uint16(1),
	1220:  uint16(anon_sym_LBRACK),
	1221:  uint16(147),
	1222:  uint16(1),
	1223:  uint16(sym_string),
	1224:  uint16(190),
	1225:  uint16(1),
	1226:  uint16(aux_sym_block_text_repeat2),
	1227:  uint16(308),
	1228:  uint16(1),
	1229:  uint16(sym_const_value),
	1230:  uint16(466),
	1231:  uint16(1),
	1232:  uint16(aux_sym__internal_const_identifier_repeat1),
	1233:  uint16(39),
	1234:  uint16(2),
	1235:  uint16(anon_sym_true),
	1236:  uint16(anon_sym_false),
	1237:  uint16(37),
	1238:  uint16(3),
	1239:  uint16(sym_number),
	1240:  uint16(sym_float),
	1241:  uint16(sym_void),
	1242:  uint16(210),
	1243:  uint16(8),
	1244:  uint16(sym_boolean),
	1245:  uint16(sym_data),
	1246:  uint16(sym_const_list),
	1247:  uint16(sym_struct_shorthand),
	1248:  uint16(sym__internal_const_identifier),
	1249:  uint16(sym_embedded_file),
	1250:  uint16(sym_concatenated_string),
	1251:  uint16(sym_block_text),
	1252:  uint16(17),
	1253:  uint16(3),
	1254:  uint16(1),
	1255:  uint16(sym_comment),
	1256:  uint16(41),
	1257:  uint16(1),
	1258:  uint16(anon_sym_0x),
	1259:  uint16(43),
	1260:  uint16(1),
	1261:  uint16(anon_sym_embed),
	1262:  uint16(45),
	1263:  uint16(1),
	1264:  uint16(anon_sym_DQUOTE),
	1265:  uint16(47),
	1266:  uint16(1),
	1267:  uint16(anon_sym_SQUOTE),
	1268:  uint16(49),
	1269:  uint16(1),
	1270:  uint16(anon_sym_BQUOTE),
	1271:  uint16(51),
	1272:  uint16(1),
	1273:  uint16(sym__identifier_no_period),
	1274:  uint16(77),
	1275:  uint16(1),
	1276:  uint16(anon_sym_LPAREN),
	1277:  uint16(81),
	1278:  uint16(1),
	1279:  uint16(anon_sym_LBRACK),
	1280:  uint16(97),
	1281:  uint16(1),
	1282:  uint16(anon_sym_RBRACK),
	1283:  uint16(147),
	1284:  uint16(1),
	1285:  uint16(sym_string),
	1286:  uint16(190),
	1287:  uint16(1),
	1288:  uint16(aux_sym_block_text_repeat2),
	1289:  uint16(367),
	1290:  uint16(1),
	1291:  uint16(sym_const_value),
	1292:  uint16(466),
	1293:  uint16(1),
	1294:  uint16(aux_sym__internal_const_identifier_repeat1),
	1295:  uint16(39),
	1296:  uint16(2),
	1297:  uint16(anon_sym_true),
	1298:  uint16(anon_sym_false),
	1299:  uint16(37),
	1300:  uint16(3),
	1301:  uint16(sym_number),
	1302:  uint16(sym_float),
	1303:  uint16(sym_void),
	1304:  uint16(210),
	1305:  uint16(8),
	1306:  uint16(sym_boolean),
	1307:  uint16(sym_data),
	1308:  uint16(sym_const_list),
	1309:  uint16(sym_struct_shorthand),
	1310:  uint16(sym__internal_const_identifier),
	1311:  uint16(sym_embedded_file),
	1312:  uint16(sym_concatenated_string),
	1313:  uint16(sym_block_text),
	1314:  uint16(16),
	1315:  uint16(3),
	1316:  uint16(1),
	1317:  uint16(sym_comment),
	1318:  uint16(41),
	1319:  uint16(1),
	1320:  uint16(anon_sym_0x),
	1321:  uint16(43),
	1322:  uint16(1),
	1323:  uint16(anon_sym_embed),
	1324:  uint16(45),
	1325:  uint16(1),
	1326:  uint16(anon_sym_DQUOTE),
	1327:  uint16(47),
	1328:  uint16(1),
	1329:  uint16(anon_sym_SQUOTE),
	1330:  uint16(49),
	1331:  uint16(1),
	1332:  uint16(anon_sym_BQUOTE),
	1333:  uint16(51),
	1334:  uint16(1),
	1335:  uint16(sym__identifier_no_period),
	1336:  uint16(77),
	1337:  uint16(1),
	1338:  uint16(anon_sym_LPAREN),
	1339:  uint16(81),
	1340:  uint16(1),
	1341:  uint16(anon_sym_LBRACK),
	1342:  uint16(147),
	1343:  uint16(1),
	1344:  uint16(sym_string),
	1345:  uint16(190),
	1346:  uint16(1),
	1347:  uint16(aux_sym_block_text_repeat2),
	1348:  uint16(466),
	1349:  uint16(1),
	1350:  uint16(aux_sym__internal_const_identifier_repeat1),
	1351:  uint16(502),
	1352:  uint16(1),
	1353:  uint16(sym_const_value),
	1354:  uint16(39),
	1355:  uint16(2),
	1356:  uint16(anon_sym_true),
	1357:  uint16(anon_sym_false),
	1358:  uint16(37),
	1359:  uint16(3),
	1360:  uint16(sym_number),
	1361:  uint16(sym_float),
	1362:  uint16(sym_void),
	1363:  uint16(210),
	1364:  uint16(8),
	1365:  uint16(sym_boolean),
	1366:  uint16(sym_data),
	1367:  uint16(sym_const_list),
	1368:  uint16(sym_struct_shorthand),
	1369:  uint16(sym__internal_const_identifier),
	1370:  uint16(sym_embedded_file),
	1371:  uint16(sym_concatenated_string),
	1372:  uint16(sym_block_text),
	1373:  uint16(16),
	1374:  uint16(3),
	1375:  uint16(1),
	1376:  uint16(sym_comment),
	1377:  uint16(41),
	1378:  uint16(1),
	1379:  uint16(anon_sym_0x),
	1380:  uint16(43),
	1381:  uint16(1),
	1382:  uint16(anon_sym_embed),
	1383:  uint16(45),
	1384:  uint16(1),
	1385:  uint16(anon_sym_DQUOTE),
	1386:  uint16(47),
	1387:  uint16(1),
	1388:  uint16(anon_sym_SQUOTE),
	1389:  uint16(49),
	1390:  uint16(1),
	1391:  uint16(anon_sym_BQUOTE),
	1392:  uint16(51),
	1393:  uint16(1),
	1394:  uint16(sym__identifier_no_period),
	1395:  uint16(77),
	1396:  uint16(1),
	1397:  uint16(anon_sym_LPAREN),
	1398:  uint16(81),
	1399:  uint16(1),
	1400:  uint16(anon_sym_LBRACK),
	1401:  uint16(147),
	1402:  uint16(1),
	1403:  uint16(sym_string),
	1404:  uint16(190),
	1405:  uint16(1),
	1406:  uint16(aux_sym_block_text_repeat2),
	1407:  uint16(350),
	1408:  uint16(1),
	1409:  uint16(sym_const_value),
	1410:  uint16(466),
	1411:  uint16(1),
	1412:  uint16(aux_sym__internal_const_identifier_repeat1),
	1413:  uint16(39),
	1414:  uint16(2),
	1415:  uint16(anon_sym_true),
	1416:  uint16(anon_sym_false),
	1417:  uint16(37),
	1418:  uint16(3),
	1419:  uint16(sym_number),
	1420:  uint16(sym_float),
	1421:  uint16(sym_void),
	1422:  uint16(210),
	1423:  uint16(8),
	1424:  uint16(sym_boolean),
	1425:  uint16(sym_data),
	1426:  uint16(sym_const_list),
	1427:  uint16(sym_struct_shorthand),
	1428:  uint16(sym__internal_const_identifier),
	1429:  uint16(sym_embedded_file),
	1430:  uint16(sym_concatenated_string),
	1431:  uint16(sym_block_text),
	1432:  uint16(16),
	1433:  uint16(3),
	1434:  uint16(1),
	1435:  uint16(sym_comment),
	1436:  uint16(41),
	1437:  uint16(1),
	1438:  uint16(anon_sym_0x),
	1439:  uint16(43),
	1440:  uint16(1),
	1441:  uint16(anon_sym_embed),
	1442:  uint16(45),
	1443:  uint16(1),
	1444:  uint16(anon_sym_DQUOTE),
	1445:  uint16(47),
	1446:  uint16(1),
	1447:  uint16(anon_sym_SQUOTE),
	1448:  uint16(49),
	1449:  uint16(1),
	1450:  uint16(anon_sym_BQUOTE),
	1451:  uint16(77),
	1452:  uint16(1),
	1453:  uint16(anon_sym_LPAREN),
	1454:  uint16(81),
	1455:  uint16(1),
	1456:  uint16(anon_sym_LBRACK),
	1457:  uint16(87),
	1458:  uint16(1),
	1459:  uint16(sym__identifier_no_period),
	1460:  uint16(147),
	1461:  uint16(1),
	1462:  uint16(sym_string),
	1463:  uint16(190),
	1464:  uint16(1),
	1465:  uint16(aux_sym_block_text_repeat2),
	1466:  uint16(329),
	1467:  uint16(1),
	1468:  uint16(sym_const_value),
	1469:  uint16(444),
	1470:  uint16(1),
	1471:  uint16(aux_sym__internal_const_identifier_repeat1),
	1472:  uint16(39),
	1473:  uint16(2),
	1474:  uint16(anon_sym_true),
	1475:  uint16(anon_sym_false),
	1476:  uint16(37),
	1477:  uint16(3),
	1478:  uint16(sym_number),
	1479:  uint16(sym_float),
	1480:  uint16(sym_void),
	1481:  uint16(210),
	1482:  uint16(8),
	1483:  uint16(sym_boolean),
	1484:  uint16(sym_data),
	1485:  uint16(sym_const_list),
	1486:  uint16(sym_struct_shorthand),
	1487:  uint16(sym__internal_const_identifier),
	1488:  uint16(sym_embedded_file),
	1489:  uint16(sym_concatenated_string),
	1490:  uint16(sym_block_text),
	1491:  uint16(16),
	1492:  uint16(3),
	1493:  uint16(1),
	1494:  uint16(sym_comment),
	1495:  uint16(41),
	1496:  uint16(1),
	1497:  uint16(anon_sym_0x),
	1498:  uint16(43),
	1499:  uint16(1),
	1500:  uint16(anon_sym_embed),
	1501:  uint16(45),
	1502:  uint16(1),
	1503:  uint16(anon_sym_DQUOTE),
	1504:  uint16(47),
	1505:  uint16(1),
	1506:  uint16(anon_sym_SQUOTE),
	1507:  uint16(49),
	1508:  uint16(1),
	1509:  uint16(anon_sym_BQUOTE),
	1510:  uint16(51),
	1511:  uint16(1),
	1512:  uint16(sym__identifier_no_period),
	1513:  uint16(77),
	1514:  uint16(1),
	1515:  uint16(anon_sym_LPAREN),
	1516:  uint16(81),
	1517:  uint16(1),
	1518:  uint16(anon_sym_LBRACK),
	1519:  uint16(147),
	1520:  uint16(1),
	1521:  uint16(sym_string),
	1522:  uint16(190),
	1523:  uint16(1),
	1524:  uint16(aux_sym_block_text_repeat2),
	1525:  uint16(456),
	1526:  uint16(1),
	1527:  uint16(sym_const_value),
	1528:  uint16(466),
	1529:  uint16(1),
	1530:  uint16(aux_sym__internal_const_identifier_repeat1),
	1531:  uint16(39),
	1532:  uint16(2),
	1533:  uint16(anon_sym_true),
	1534:  uint16(anon_sym_false),
	1535:  uint16(37),
	1536:  uint16(3),
	1537:  uint16(sym_number),
	1538:  uint16(sym_float),
	1539:  uint16(sym_void),
	1540:  uint16(210),
	1541:  uint16(8),
	1542:  uint16(sym_boolean),
	1543:  uint16(sym_data),
	1544:  uint16(sym_const_list),
	1545:  uint16(sym_struct_shorthand),
	1546:  uint16(sym__internal_const_identifier),
	1547:  uint16(sym_embedded_file),
	1548:  uint16(sym_concatenated_string),
	1549:  uint16(sym_block_text),
	1550:  uint16(16),
	1551:  uint16(3),
	1552:  uint16(1),
	1553:  uint16(sym_comment),
	1554:  uint16(41),
	1555:  uint16(1),
	1556:  uint16(anon_sym_0x),
	1557:  uint16(43),
	1558:  uint16(1),
	1559:  uint16(anon_sym_embed),
	1560:  uint16(45),
	1561:  uint16(1),
	1562:  uint16(anon_sym_DQUOTE),
	1563:  uint16(47),
	1564:  uint16(1),
	1565:  uint16(anon_sym_SQUOTE),
	1566:  uint16(49),
	1567:  uint16(1),
	1568:  uint16(anon_sym_BQUOTE),
	1569:  uint16(51),
	1570:  uint16(1),
	1571:  uint16(sym__identifier_no_period),
	1572:  uint16(77),
	1573:  uint16(1),
	1574:  uint16(anon_sym_LPAREN),
	1575:  uint16(81),
	1576:  uint16(1),
	1577:  uint16(anon_sym_LBRACK),
	1578:  uint16(147),
	1579:  uint16(1),
	1580:  uint16(sym_string),
	1581:  uint16(190),
	1582:  uint16(1),
	1583:  uint16(aux_sym_block_text_repeat2),
	1584:  uint16(466),
	1585:  uint16(1),
	1586:  uint16(aux_sym__internal_const_identifier_repeat1),
	1587:  uint16(472),
	1588:  uint16(1),
	1589:  uint16(sym_const_value),
	1590:  uint16(39),
	1591:  uint16(2),
	1592:  uint16(anon_sym_true),
	1593:  uint16(anon_sym_false),
	1594:  uint16(37),
	1595:  uint16(3),
	1596:  uint16(sym_number),
	1597:  uint16(sym_float),
	1598:  uint16(sym_void),
	1599:  uint16(210),
	1600:  uint16(8),
	1601:  uint16(sym_boolean),
	1602:  uint16(sym_data),
	1603:  uint16(sym_const_list),
	1604:  uint16(sym_struct_shorthand),
	1605:  uint16(sym__internal_const_identifier),
	1606:  uint16(sym_embedded_file),
	1607:  uint16(sym_concatenated_string),
	1608:  uint16(sym_block_text),
	1609:  uint16(16),
	1610:  uint16(3),
	1611:  uint16(1),
	1612:  uint16(sym_comment),
	1613:  uint16(41),
	1614:  uint16(1),
	1615:  uint16(anon_sym_0x),
	1616:  uint16(43),
	1617:  uint16(1),
	1618:  uint16(anon_sym_embed),
	1619:  uint16(45),
	1620:  uint16(1),
	1621:  uint16(anon_sym_DQUOTE),
	1622:  uint16(47),
	1623:  uint16(1),
	1624:  uint16(anon_sym_SQUOTE),
	1625:  uint16(49),
	1626:  uint16(1),
	1627:  uint16(anon_sym_BQUOTE),
	1628:  uint16(51),
	1629:  uint16(1),
	1630:  uint16(sym__identifier_no_period),
	1631:  uint16(77),
	1632:  uint16(1),
	1633:  uint16(anon_sym_LPAREN),
	1634:  uint16(81),
	1635:  uint16(1),
	1636:  uint16(anon_sym_LBRACK),
	1637:  uint16(147),
	1638:  uint16(1),
	1639:  uint16(sym_string),
	1640:  uint16(190),
	1641:  uint16(1),
	1642:  uint16(aux_sym_block_text_repeat2),
	1643:  uint16(466),
	1644:  uint16(1),
	1645:  uint16(aux_sym__internal_const_identifier_repeat1),
	1646:  uint16(523),
	1647:  uint16(1),
	1648:  uint16(sym_const_value),
	1649:  uint16(39),
	1650:  uint16(2),
	1651:  uint16(anon_sym_true),
	1652:  uint16(anon_sym_false),
	1653:  uint16(37),
	1654:  uint16(3),
	1655:  uint16(sym_number),
	1656:  uint16(sym_float),
	1657:  uint16(sym_void),
	1658:  uint16(210),
	1659:  uint16(8),
	1660:  uint16(sym_boolean),
	1661:  uint16(sym_data),
	1662:  uint16(sym_const_list),
	1663:  uint16(sym_struct_shorthand),
	1664:  uint16(sym__internal_const_identifier),
	1665:  uint16(sym_embedded_file),
	1666:  uint16(sym_concatenated_string),
	1667:  uint16(sym_block_text),
	1668:  uint16(16),
	1669:  uint16(3),
	1670:  uint16(1),
	1671:  uint16(sym_comment),
	1672:  uint16(41),
	1673:  uint16(1),
	1674:  uint16(anon_sym_0x),
	1675:  uint16(43),
	1676:  uint16(1),
	1677:  uint16(anon_sym_embed),
	1678:  uint16(45),
	1679:  uint16(1),
	1680:  uint16(anon_sym_DQUOTE),
	1681:  uint16(47),
	1682:  uint16(1),
	1683:  uint16(anon_sym_SQUOTE),
	1684:  uint16(49),
	1685:  uint16(1),
	1686:  uint16(anon_sym_BQUOTE),
	1687:  uint16(51),
	1688:  uint16(1),
	1689:  uint16(sym__identifier_no_period),
	1690:  uint16(77),
	1691:  uint16(1),
	1692:  uint16(anon_sym_LPAREN),
	1693:  uint16(81),
	1694:  uint16(1),
	1695:  uint16(anon_sym_LBRACK),
	1696:  uint16(147),
	1697:  uint16(1),
	1698:  uint16(sym_string),
	1699:  uint16(190),
	1700:  uint16(1),
	1701:  uint16(aux_sym_block_text_repeat2),
	1702:  uint16(332),
	1703:  uint16(1),
	1704:  uint16(sym_const_value),
	1705:  uint16(466),
	1706:  uint16(1),
	1707:  uint16(aux_sym__internal_const_identifier_repeat1),
	1708:  uint16(39),
	1709:  uint16(2),
	1710:  uint16(anon_sym_true),
	1711:  uint16(anon_sym_false),
	1712:  uint16(37),
	1713:  uint16(3),
	1714:  uint16(sym_number),
	1715:  uint16(sym_float),
	1716:  uint16(sym_void),
	1717:  uint16(210),
	1718:  uint16(8),
	1719:  uint16(sym_boolean),
	1720:  uint16(sym_data),
	1721:  uint16(sym_const_list),
	1722:  uint16(sym_struct_shorthand),
	1723:  uint16(sym__internal_const_identifier),
	1724:  uint16(sym_embedded_file),
	1725:  uint16(sym_concatenated_string),
	1726:  uint16(sym_block_text),
	1727:  uint16(16),
	1728:  uint16(3),
	1729:  uint16(1),
	1730:  uint16(sym_comment),
	1731:  uint16(41),
	1732:  uint16(1),
	1733:  uint16(anon_sym_0x),
	1734:  uint16(43),
	1735:  uint16(1),
	1736:  uint16(anon_sym_embed),
	1737:  uint16(45),
	1738:  uint16(1),
	1739:  uint16(anon_sym_DQUOTE),
	1740:  uint16(47),
	1741:  uint16(1),
	1742:  uint16(anon_sym_SQUOTE),
	1743:  uint16(49),
	1744:  uint16(1),
	1745:  uint16(anon_sym_BQUOTE),
	1746:  uint16(51),
	1747:  uint16(1),
	1748:  uint16(sym__identifier_no_period),
	1749:  uint16(77),
	1750:  uint16(1),
	1751:  uint16(anon_sym_LPAREN),
	1752:  uint16(81),
	1753:  uint16(1),
	1754:  uint16(anon_sym_LBRACK),
	1755:  uint16(147),
	1756:  uint16(1),
	1757:  uint16(sym_string),
	1758:  uint16(190),
	1759:  uint16(1),
	1760:  uint16(aux_sym_block_text_repeat2),
	1761:  uint16(310),
	1762:  uint16(1),
	1763:  uint16(sym_const_value),
	1764:  uint16(466),
	1765:  uint16(1),
	1766:  uint16(aux_sym__internal_const_identifier_repeat1),
	1767:  uint16(39),
	1768:  uint16(2),
	1769:  uint16(anon_sym_true),
	1770:  uint16(anon_sym_false),
	1771:  uint16(37),
	1772:  uint16(3),
	1773:  uint16(sym_number),
	1774:  uint16(sym_float),
	1775:  uint16(sym_void),
	1776:  uint16(210),
	1777:  uint16(8),
	1778:  uint16(sym_boolean),
	1779:  uint16(sym_data),
	1780:  uint16(sym_const_list),
	1781:  uint16(sym_struct_shorthand),
	1782:  uint16(sym__internal_const_identifier),
	1783:  uint16(sym_embedded_file),
	1784:  uint16(sym_concatenated_string),
	1785:  uint16(sym_block_text),
	1786:  uint16(16),
	1787:  uint16(3),
	1788:  uint16(1),
	1789:  uint16(sym_comment),
	1790:  uint16(41),
	1791:  uint16(1),
	1792:  uint16(anon_sym_0x),
	1793:  uint16(43),
	1794:  uint16(1),
	1795:  uint16(anon_sym_embed),
	1796:  uint16(45),
	1797:  uint16(1),
	1798:  uint16(anon_sym_DQUOTE),
	1799:  uint16(47),
	1800:  uint16(1),
	1801:  uint16(anon_sym_SQUOTE),
	1802:  uint16(49),
	1803:  uint16(1),
	1804:  uint16(anon_sym_BQUOTE),
	1805:  uint16(51),
	1806:  uint16(1),
	1807:  uint16(sym__identifier_no_period),
	1808:  uint16(77),
	1809:  uint16(1),
	1810:  uint16(anon_sym_LPAREN),
	1811:  uint16(81),
	1812:  uint16(1),
	1813:  uint16(anon_sym_LBRACK),
	1814:  uint16(147),
	1815:  uint16(1),
	1816:  uint16(sym_string),
	1817:  uint16(190),
	1818:  uint16(1),
	1819:  uint16(aux_sym_block_text_repeat2),
	1820:  uint16(406),
	1821:  uint16(1),
	1822:  uint16(sym_const_value),
	1823:  uint16(466),
	1824:  uint16(1),
	1825:  uint16(aux_sym__internal_const_identifier_repeat1),
	1826:  uint16(39),
	1827:  uint16(2),
	1828:  uint16(anon_sym_true),
	1829:  uint16(anon_sym_false),
	1830:  uint16(37),
	1831:  uint16(3),
	1832:  uint16(sym_number),
	1833:  uint16(sym_float),
	1834:  uint16(sym_void),
	1835:  uint16(210),
	1836:  uint16(8),
	1837:  uint16(sym_boolean),
	1838:  uint16(sym_data),
	1839:  uint16(sym_const_list),
	1840:  uint16(sym_struct_shorthand),
	1841:  uint16(sym__internal_const_identifier),
	1842:  uint16(sym_embedded_file),
	1843:  uint16(sym_concatenated_string),
	1844:  uint16(sym_block_text),
	1845:  uint16(16),
	1846:  uint16(3),
	1847:  uint16(1),
	1848:  uint16(sym_comment),
	1849:  uint16(41),
	1850:  uint16(1),
	1851:  uint16(anon_sym_0x),
	1852:  uint16(43),
	1853:  uint16(1),
	1854:  uint16(anon_sym_embed),
	1855:  uint16(45),
	1856:  uint16(1),
	1857:  uint16(anon_sym_DQUOTE),
	1858:  uint16(47),
	1859:  uint16(1),
	1860:  uint16(anon_sym_SQUOTE),
	1861:  uint16(49),
	1862:  uint16(1),
	1863:  uint16(anon_sym_BQUOTE),
	1864:  uint16(51),
	1865:  uint16(1),
	1866:  uint16(sym__identifier_no_period),
	1867:  uint16(77),
	1868:  uint16(1),
	1869:  uint16(anon_sym_LPAREN),
	1870:  uint16(81),
	1871:  uint16(1),
	1872:  uint16(anon_sym_LBRACK),
	1873:  uint16(147),
	1874:  uint16(1),
	1875:  uint16(sym_string),
	1876:  uint16(190),
	1877:  uint16(1),
	1878:  uint16(aux_sym_block_text_repeat2),
	1879:  uint16(436),
	1880:  uint16(1),
	1881:  uint16(sym_const_value),
	1882:  uint16(466),
	1883:  uint16(1),
	1884:  uint16(aux_sym__internal_const_identifier_repeat1),
	1885:  uint16(39),
	1886:  uint16(2),
	1887:  uint16(anon_sym_true),
	1888:  uint16(anon_sym_false),
	1889:  uint16(37),
	1890:  uint16(3),
	1891:  uint16(sym_number),
	1892:  uint16(sym_float),
	1893:  uint16(sym_void),
	1894:  uint16(210),
	1895:  uint16(8),
	1896:  uint16(sym_boolean),
	1897:  uint16(sym_data),
	1898:  uint16(sym_const_list),
	1899:  uint16(sym_struct_shorthand),
	1900:  uint16(sym__internal_const_identifier),
	1901:  uint16(sym_embedded_file),
	1902:  uint16(sym_concatenated_string),
	1903:  uint16(sym_block_text),
	1904:  uint16(16),
	1905:  uint16(3),
	1906:  uint16(1),
	1907:  uint16(sym_comment),
	1908:  uint16(41),
	1909:  uint16(1),
	1910:  uint16(anon_sym_0x),
	1911:  uint16(43),
	1912:  uint16(1),
	1913:  uint16(anon_sym_embed),
	1914:  uint16(45),
	1915:  uint16(1),
	1916:  uint16(anon_sym_DQUOTE),
	1917:  uint16(47),
	1918:  uint16(1),
	1919:  uint16(anon_sym_SQUOTE),
	1920:  uint16(49),
	1921:  uint16(1),
	1922:  uint16(anon_sym_BQUOTE),
	1923:  uint16(51),
	1924:  uint16(1),
	1925:  uint16(sym__identifier_no_period),
	1926:  uint16(77),
	1927:  uint16(1),
	1928:  uint16(anon_sym_LPAREN),
	1929:  uint16(81),
	1930:  uint16(1),
	1931:  uint16(anon_sym_LBRACK),
	1932:  uint16(147),
	1933:  uint16(1),
	1934:  uint16(sym_string),
	1935:  uint16(190),
	1936:  uint16(1),
	1937:  uint16(aux_sym_block_text_repeat2),
	1938:  uint16(318),
	1939:  uint16(1),
	1940:  uint16(sym_const_value),
	1941:  uint16(466),
	1942:  uint16(1),
	1943:  uint16(aux_sym__internal_const_identifier_repeat1),
	1944:  uint16(39),
	1945:  uint16(2),
	1946:  uint16(anon_sym_true),
	1947:  uint16(anon_sym_false),
	1948:  uint16(37),
	1949:  uint16(3),
	1950:  uint16(sym_number),
	1951:  uint16(sym_float),
	1952:  uint16(sym_void),
	1953:  uint16(210),
	1954:  uint16(8),
	1955:  uint16(sym_boolean),
	1956:  uint16(sym_data),
	1957:  uint16(sym_const_list),
	1958:  uint16(sym_struct_shorthand),
	1959:  uint16(sym__internal_const_identifier),
	1960:  uint16(sym_embedded_file),
	1961:  uint16(sym_concatenated_string),
	1962:  uint16(sym_block_text),
	1963:  uint16(16),
	1964:  uint16(3),
	1965:  uint16(1),
	1966:  uint16(sym_comment),
	1967:  uint16(41),
	1968:  uint16(1),
	1969:  uint16(anon_sym_0x),
	1970:  uint16(43),
	1971:  uint16(1),
	1972:  uint16(anon_sym_embed),
	1973:  uint16(45),
	1974:  uint16(1),
	1975:  uint16(anon_sym_DQUOTE),
	1976:  uint16(47),
	1977:  uint16(1),
	1978:  uint16(anon_sym_SQUOTE),
	1979:  uint16(49),
	1980:  uint16(1),
	1981:  uint16(anon_sym_BQUOTE),
	1982:  uint16(51),
	1983:  uint16(1),
	1984:  uint16(sym__identifier_no_period),
	1985:  uint16(77),
	1986:  uint16(1),
	1987:  uint16(anon_sym_LPAREN),
	1988:  uint16(81),
	1989:  uint16(1),
	1990:  uint16(anon_sym_LBRACK),
	1991:  uint16(147),
	1992:  uint16(1),
	1993:  uint16(sym_string),
	1994:  uint16(190),
	1995:  uint16(1),
	1996:  uint16(aux_sym_block_text_repeat2),
	1997:  uint16(349),
	1998:  uint16(1),
	1999:  uint16(sym_const_value),
	2000:  uint16(466),
	2001:  uint16(1),
	2002:  uint16(aux_sym__internal_const_identifier_repeat1),
	2003:  uint16(39),
	2004:  uint16(2),
	2005:  uint16(anon_sym_true),
	2006:  uint16(anon_sym_false),
	2007:  uint16(37),
	2008:  uint16(3),
	2009:  uint16(sym_number),
	2010:  uint16(sym_float),
	2011:  uint16(sym_void),
	2012:  uint16(210),
	2013:  uint16(8),
	2014:  uint16(sym_boolean),
	2015:  uint16(sym_data),
	2016:  uint16(sym_const_list),
	2017:  uint16(sym_struct_shorthand),
	2018:  uint16(sym__internal_const_identifier),
	2019:  uint16(sym_embedded_file),
	2020:  uint16(sym_concatenated_string),
	2021:  uint16(sym_block_text),
	2022:  uint16(16),
	2023:  uint16(3),
	2024:  uint16(1),
	2025:  uint16(sym_comment),
	2026:  uint16(41),
	2027:  uint16(1),
	2028:  uint16(anon_sym_0x),
	2029:  uint16(43),
	2030:  uint16(1),
	2031:  uint16(anon_sym_embed),
	2032:  uint16(45),
	2033:  uint16(1),
	2034:  uint16(anon_sym_DQUOTE),
	2035:  uint16(47),
	2036:  uint16(1),
	2037:  uint16(anon_sym_SQUOTE),
	2038:  uint16(49),
	2039:  uint16(1),
	2040:  uint16(anon_sym_BQUOTE),
	2041:  uint16(51),
	2042:  uint16(1),
	2043:  uint16(sym__identifier_no_period),
	2044:  uint16(77),
	2045:  uint16(1),
	2046:  uint16(anon_sym_LPAREN),
	2047:  uint16(81),
	2048:  uint16(1),
	2049:  uint16(anon_sym_LBRACK),
	2050:  uint16(147),
	2051:  uint16(1),
	2052:  uint16(sym_string),
	2053:  uint16(190),
	2054:  uint16(1),
	2055:  uint16(aux_sym_block_text_repeat2),
	2056:  uint16(447),
	2057:  uint16(1),
	2058:  uint16(sym_const_value),
	2059:  uint16(466),
	2060:  uint16(1),
	2061:  uint16(aux_sym__internal_const_identifier_repeat1),
	2062:  uint16(39),
	2063:  uint16(2),
	2064:  uint16(anon_sym_true),
	2065:  uint16(anon_sym_false),
	2066:  uint16(37),
	2067:  uint16(3),
	2068:  uint16(sym_number),
	2069:  uint16(sym_float),
	2070:  uint16(sym_void),
	2071:  uint16(210),
	2072:  uint16(8),
	2073:  uint16(sym_boolean),
	2074:  uint16(sym_data),
	2075:  uint16(sym_const_list),
	2076:  uint16(sym_struct_shorthand),
	2077:  uint16(sym__internal_const_identifier),
	2078:  uint16(sym_embedded_file),
	2079:  uint16(sym_concatenated_string),
	2080:  uint16(sym_block_text),
	2081:  uint16(16),
	2082:  uint16(3),
	2083:  uint16(1),
	2084:  uint16(sym_comment),
	2085:  uint16(41),
	2086:  uint16(1),
	2087:  uint16(anon_sym_0x),
	2088:  uint16(43),
	2089:  uint16(1),
	2090:  uint16(anon_sym_embed),
	2091:  uint16(45),
	2092:  uint16(1),
	2093:  uint16(anon_sym_DQUOTE),
	2094:  uint16(47),
	2095:  uint16(1),
	2096:  uint16(anon_sym_SQUOTE),
	2097:  uint16(49),
	2098:  uint16(1),
	2099:  uint16(anon_sym_BQUOTE),
	2100:  uint16(51),
	2101:  uint16(1),
	2102:  uint16(sym__identifier_no_period),
	2103:  uint16(77),
	2104:  uint16(1),
	2105:  uint16(anon_sym_LPAREN),
	2106:  uint16(81),
	2107:  uint16(1),
	2108:  uint16(anon_sym_LBRACK),
	2109:  uint16(147),
	2110:  uint16(1),
	2111:  uint16(sym_string),
	2112:  uint16(190),
	2113:  uint16(1),
	2114:  uint16(aux_sym_block_text_repeat2),
	2115:  uint16(367),
	2116:  uint16(1),
	2117:  uint16(sym_const_value),
	2118:  uint16(466),
	2119:  uint16(1),
	2120:  uint16(aux_sym__internal_const_identifier_repeat1),
	2121:  uint16(39),
	2122:  uint16(2),
	2123:  uint16(anon_sym_true),
	2124:  uint16(anon_sym_false),
	2125:  uint16(37),
	2126:  uint16(3),
	2127:  uint16(sym_number),
	2128:  uint16(sym_float),
	2129:  uint16(sym_void),
	2130:  uint16(210),
	2131:  uint16(8),
	2132:  uint16(sym_boolean),
	2133:  uint16(sym_data),
	2134:  uint16(sym_const_list),
	2135:  uint16(sym_struct_shorthand),
	2136:  uint16(sym__internal_const_identifier),
	2137:  uint16(sym_embedded_file),
	2138:  uint16(sym_concatenated_string),
	2139:  uint16(sym_block_text),
	2140:  uint16(16),
	2141:  uint16(3),
	2142:  uint16(1),
	2143:  uint16(sym_comment),
	2144:  uint16(41),
	2145:  uint16(1),
	2146:  uint16(anon_sym_0x),
	2147:  uint16(43),
	2148:  uint16(1),
	2149:  uint16(anon_sym_embed),
	2150:  uint16(45),
	2151:  uint16(1),
	2152:  uint16(anon_sym_DQUOTE),
	2153:  uint16(47),
	2154:  uint16(1),
	2155:  uint16(anon_sym_SQUOTE),
	2156:  uint16(49),
	2157:  uint16(1),
	2158:  uint16(anon_sym_BQUOTE),
	2159:  uint16(51),
	2160:  uint16(1),
	2161:  uint16(sym__identifier_no_period),
	2162:  uint16(77),
	2163:  uint16(1),
	2164:  uint16(anon_sym_LPAREN),
	2165:  uint16(81),
	2166:  uint16(1),
	2167:  uint16(anon_sym_LBRACK),
	2168:  uint16(147),
	2169:  uint16(1),
	2170:  uint16(sym_string),
	2171:  uint16(190),
	2172:  uint16(1),
	2173:  uint16(aux_sym_block_text_repeat2),
	2174:  uint16(304),
	2175:  uint16(1),
	2176:  uint16(sym_const_value),
	2177:  uint16(466),
	2178:  uint16(1),
	2179:  uint16(aux_sym__internal_const_identifier_repeat1),
	2180:  uint16(39),
	2181:  uint16(2),
	2182:  uint16(anon_sym_true),
	2183:  uint16(anon_sym_false),
	2184:  uint16(37),
	2185:  uint16(3),
	2186:  uint16(sym_number),
	2187:  uint16(sym_float),
	2188:  uint16(sym_void),
	2189:  uint16(210),
	2190:  uint16(8),
	2191:  uint16(sym_boolean),
	2192:  uint16(sym_data),
	2193:  uint16(sym_const_list),
	2194:  uint16(sym_struct_shorthand),
	2195:  uint16(sym__internal_const_identifier),
	2196:  uint16(sym_embedded_file),
	2197:  uint16(sym_concatenated_string),
	2198:  uint16(sym_block_text),
	2199:  uint16(16),
	2200:  uint16(3),
	2201:  uint16(1),
	2202:  uint16(sym_comment),
	2203:  uint16(41),
	2204:  uint16(1),
	2205:  uint16(anon_sym_0x),
	2206:  uint16(43),
	2207:  uint16(1),
	2208:  uint16(anon_sym_embed),
	2209:  uint16(45),
	2210:  uint16(1),
	2211:  uint16(anon_sym_DQUOTE),
	2212:  uint16(47),
	2213:  uint16(1),
	2214:  uint16(anon_sym_SQUOTE),
	2215:  uint16(49),
	2216:  uint16(1),
	2217:  uint16(anon_sym_BQUOTE),
	2218:  uint16(51),
	2219:  uint16(1),
	2220:  uint16(sym__identifier_no_period),
	2221:  uint16(77),
	2222:  uint16(1),
	2223:  uint16(anon_sym_LPAREN),
	2224:  uint16(81),
	2225:  uint16(1),
	2226:  uint16(anon_sym_LBRACK),
	2227:  uint16(147),
	2228:  uint16(1),
	2229:  uint16(sym_string),
	2230:  uint16(190),
	2231:  uint16(1),
	2232:  uint16(aux_sym_block_text_repeat2),
	2233:  uint16(346),
	2234:  uint16(1),
	2235:  uint16(sym_const_value),
	2236:  uint16(466),
	2237:  uint16(1),
	2238:  uint16(aux_sym__internal_const_identifier_repeat1),
	2239:  uint16(39),
	2240:  uint16(2),
	2241:  uint16(anon_sym_true),
	2242:  uint16(anon_sym_false),
	2243:  uint16(37),
	2244:  uint16(3),
	2245:  uint16(sym_number),
	2246:  uint16(sym_float),
	2247:  uint16(sym_void),
	2248:  uint16(210),
	2249:  uint16(8),
	2250:  uint16(sym_boolean),
	2251:  uint16(sym_data),
	2252:  uint16(sym_const_list),
	2253:  uint16(sym_struct_shorthand),
	2254:  uint16(sym__internal_const_identifier),
	2255:  uint16(sym_embedded_file),
	2256:  uint16(sym_concatenated_string),
	2257:  uint16(sym_block_text),
	2258:  uint16(16),
	2259:  uint16(3),
	2260:  uint16(1),
	2261:  uint16(sym_comment),
	2262:  uint16(41),
	2263:  uint16(1),
	2264:  uint16(anon_sym_0x),
	2265:  uint16(43),
	2266:  uint16(1),
	2267:  uint16(anon_sym_embed),
	2268:  uint16(45),
	2269:  uint16(1),
	2270:  uint16(anon_sym_DQUOTE),
	2271:  uint16(47),
	2272:  uint16(1),
	2273:  uint16(anon_sym_SQUOTE),
	2274:  uint16(49),
	2275:  uint16(1),
	2276:  uint16(anon_sym_BQUOTE),
	2277:  uint16(51),
	2278:  uint16(1),
	2279:  uint16(sym__identifier_no_period),
	2280:  uint16(77),
	2281:  uint16(1),
	2282:  uint16(anon_sym_LPAREN),
	2283:  uint16(81),
	2284:  uint16(1),
	2285:  uint16(anon_sym_LBRACK),
	2286:  uint16(147),
	2287:  uint16(1),
	2288:  uint16(sym_string),
	2289:  uint16(190),
	2290:  uint16(1),
	2291:  uint16(aux_sym_block_text_repeat2),
	2292:  uint16(408),
	2293:  uint16(1),
	2294:  uint16(sym_const_value),
	2295:  uint16(466),
	2296:  uint16(1),
	2297:  uint16(aux_sym__internal_const_identifier_repeat1),
	2298:  uint16(39),
	2299:  uint16(2),
	2300:  uint16(anon_sym_true),
	2301:  uint16(anon_sym_false),
	2302:  uint16(37),
	2303:  uint16(3),
	2304:  uint16(sym_number),
	2305:  uint16(sym_float),
	2306:  uint16(sym_void),
	2307:  uint16(210),
	2308:  uint16(8),
	2309:  uint16(sym_boolean),
	2310:  uint16(sym_data),
	2311:  uint16(sym_const_list),
	2312:  uint16(sym_struct_shorthand),
	2313:  uint16(sym__internal_const_identifier),
	2314:  uint16(sym_embedded_file),
	2315:  uint16(sym_concatenated_string),
	2316:  uint16(sym_block_text),
	2317:  uint16(16),
	2318:  uint16(3),
	2319:  uint16(1),
	2320:  uint16(sym_comment),
	2321:  uint16(41),
	2322:  uint16(1),
	2323:  uint16(anon_sym_0x),
	2324:  uint16(43),
	2325:  uint16(1),
	2326:  uint16(anon_sym_embed),
	2327:  uint16(45),
	2328:  uint16(1),
	2329:  uint16(anon_sym_DQUOTE),
	2330:  uint16(47),
	2331:  uint16(1),
	2332:  uint16(anon_sym_SQUOTE),
	2333:  uint16(49),
	2334:  uint16(1),
	2335:  uint16(anon_sym_BQUOTE),
	2336:  uint16(51),
	2337:  uint16(1),
	2338:  uint16(sym__identifier_no_period),
	2339:  uint16(77),
	2340:  uint16(1),
	2341:  uint16(anon_sym_LPAREN),
	2342:  uint16(81),
	2343:  uint16(1),
	2344:  uint16(anon_sym_LBRACK),
	2345:  uint16(147),
	2346:  uint16(1),
	2347:  uint16(sym_string),
	2348:  uint16(190),
	2349:  uint16(1),
	2350:  uint16(aux_sym_block_text_repeat2),
	2351:  uint16(448),
	2352:  uint16(1),
	2353:  uint16(sym_const_value),
	2354:  uint16(466),
	2355:  uint16(1),
	2356:  uint16(aux_sym__internal_const_identifier_repeat1),
	2357:  uint16(39),
	2358:  uint16(2),
	2359:  uint16(anon_sym_true),
	2360:  uint16(anon_sym_false),
	2361:  uint16(37),
	2362:  uint16(3),
	2363:  uint16(sym_number),
	2364:  uint16(sym_float),
	2365:  uint16(sym_void),
	2366:  uint16(210),
	2367:  uint16(8),
	2368:  uint16(sym_boolean),
	2369:  uint16(sym_data),
	2370:  uint16(sym_const_list),
	2371:  uint16(sym_struct_shorthand),
	2372:  uint16(sym__internal_const_identifier),
	2373:  uint16(sym_embedded_file),
	2374:  uint16(sym_concatenated_string),
	2375:  uint16(sym_block_text),
	2376:  uint16(16),
	2377:  uint16(3),
	2378:  uint16(1),
	2379:  uint16(sym_comment),
	2380:  uint16(41),
	2381:  uint16(1),
	2382:  uint16(anon_sym_0x),
	2383:  uint16(43),
	2384:  uint16(1),
	2385:  uint16(anon_sym_embed),
	2386:  uint16(45),
	2387:  uint16(1),
	2388:  uint16(anon_sym_DQUOTE),
	2389:  uint16(47),
	2390:  uint16(1),
	2391:  uint16(anon_sym_SQUOTE),
	2392:  uint16(49),
	2393:  uint16(1),
	2394:  uint16(anon_sym_BQUOTE),
	2395:  uint16(51),
	2396:  uint16(1),
	2397:  uint16(sym__identifier_no_period),
	2398:  uint16(77),
	2399:  uint16(1),
	2400:  uint16(anon_sym_LPAREN),
	2401:  uint16(81),
	2402:  uint16(1),
	2403:  uint16(anon_sym_LBRACK),
	2404:  uint16(147),
	2405:  uint16(1),
	2406:  uint16(sym_string),
	2407:  uint16(190),
	2408:  uint16(1),
	2409:  uint16(aux_sym_block_text_repeat2),
	2410:  uint16(449),
	2411:  uint16(1),
	2412:  uint16(sym_const_value),
	2413:  uint16(466),
	2414:  uint16(1),
	2415:  uint16(aux_sym__internal_const_identifier_repeat1),
	2416:  uint16(39),
	2417:  uint16(2),
	2418:  uint16(anon_sym_true),
	2419:  uint16(anon_sym_false),
	2420:  uint16(37),
	2421:  uint16(3),
	2422:  uint16(sym_number),
	2423:  uint16(sym_float),
	2424:  uint16(sym_void),
	2425:  uint16(210),
	2426:  uint16(8),
	2427:  uint16(sym_boolean),
	2428:  uint16(sym_data),
	2429:  uint16(sym_const_list),
	2430:  uint16(sym_struct_shorthand),
	2431:  uint16(sym__internal_const_identifier),
	2432:  uint16(sym_embedded_file),
	2433:  uint16(sym_concatenated_string),
	2434:  uint16(sym_block_text),
	2435:  uint16(16),
	2436:  uint16(3),
	2437:  uint16(1),
	2438:  uint16(sym_comment),
	2439:  uint16(41),
	2440:  uint16(1),
	2441:  uint16(anon_sym_0x),
	2442:  uint16(43),
	2443:  uint16(1),
	2444:  uint16(anon_sym_embed),
	2445:  uint16(45),
	2446:  uint16(1),
	2447:  uint16(anon_sym_DQUOTE),
	2448:  uint16(47),
	2449:  uint16(1),
	2450:  uint16(anon_sym_SQUOTE),
	2451:  uint16(49),
	2452:  uint16(1),
	2453:  uint16(anon_sym_BQUOTE),
	2454:  uint16(51),
	2455:  uint16(1),
	2456:  uint16(sym__identifier_no_period),
	2457:  uint16(77),
	2458:  uint16(1),
	2459:  uint16(anon_sym_LPAREN),
	2460:  uint16(81),
	2461:  uint16(1),
	2462:  uint16(anon_sym_LBRACK),
	2463:  uint16(147),
	2464:  uint16(1),
	2465:  uint16(sym_string),
	2466:  uint16(190),
	2467:  uint16(1),
	2468:  uint16(aux_sym_block_text_repeat2),
	2469:  uint16(393),
	2470:  uint16(1),
	2471:  uint16(sym_const_value),
	2472:  uint16(466),
	2473:  uint16(1),
	2474:  uint16(aux_sym__internal_const_identifier_repeat1),
	2475:  uint16(39),
	2476:  uint16(2),
	2477:  uint16(anon_sym_true),
	2478:  uint16(anon_sym_false),
	2479:  uint16(37),
	2480:  uint16(3),
	2481:  uint16(sym_number),
	2482:  uint16(sym_float),
	2483:  uint16(sym_void),
	2484:  uint16(210),
	2485:  uint16(8),
	2486:  uint16(sym_boolean),
	2487:  uint16(sym_data),
	2488:  uint16(sym_const_list),
	2489:  uint16(sym_struct_shorthand),
	2490:  uint16(sym__internal_const_identifier),
	2491:  uint16(sym_embedded_file),
	2492:  uint16(sym_concatenated_string),
	2493:  uint16(sym_block_text),
	2494:  uint16(15),
	2495:  uint16(3),
	2496:  uint16(1),
	2497:  uint16(sym_comment),
	2498:  uint16(99),
	2499:  uint16(1),
	2500:  uint16(sym_identifier),
	2501:  uint16(101),
	2502:  uint16(1),
	2503:  uint16(anon_sym_using),
	2504:  uint16(103),
	2505:  uint16(1),
	2506:  uint16(anon_sym_annotation),
	2507:  uint16(105),
	2508:  uint16(1),
	2509:  uint16(anon_sym_const),
	2510:  uint16(107),
	2511:  uint16(1),
	2512:  uint16(anon_sym_enum),
	2513:  uint16(109),
	2514:  uint16(1),
	2515:  uint16(anon_sym_interface),
	2516:  uint16(111),
	2517:  uint16(1),
	2518:  uint16(anon_sym_struct),
	2519:  uint16(113),
	2520:  uint16(1),
	2521:  uint16(anon_sym_union),
	2522:  uint16(115),
	2523:  uint16(1),
	2524:  uint16(anon_sym_RBRACE),
	2525:  uint16(177),
	2526:  uint16(1),
	2527:  uint16(sym_struct),
	2528:  uint16(179),
	2529:  uint16(1),
	2530:  uint16(sym_enum),
	2531:  uint16(182),
	2532:  uint16(2),
	2533:  uint16(sym__unnamed_union),
	2534:  uint16(sym__named_union),
	2535:  uint16(43),
	2536:  uint16(3),
	2537:  uint16(sym_using_directive),
	2538:  uint16(sym_field),
	2539:  uint16(aux_sym_struct_repeat1),
	2540:  uint16(165),
	2541:  uint16(7),
	2542:  uint16(sym_annotation),
	2543:  uint16(sym_nested_struct),
	2544:  uint16(sym_nested_enum),
	2545:  uint16(sym_group),
	2546:  uint16(sym_union),
	2547:  uint16(sym_interface),
	2548:  uint16(sym_const),
	2549:  uint16(15),
	2550:  uint16(3),
	2551:  uint16(1),
	2552:  uint16(sym_comment),
	2553:  uint16(99),
	2554:  uint16(1),
	2555:  uint16(sym_identifier),
	2556:  uint16(101),
	2557:  uint16(1),
	2558:  uint16(anon_sym_using),
	2559:  uint16(103),
	2560:  uint16(1),
	2561:  uint16(anon_sym_annotation),
	2562:  uint16(105),
	2563:  uint16(1),
	2564:  uint16(anon_sym_const),
	2565:  uint16(107),
	2566:  uint16(1),
	2567:  uint16(anon_sym_enum),
	2568:  uint16(109),
	2569:  uint16(1),
	2570:  uint16(anon_sym_interface),
	2571:  uint16(111),
	2572:  uint16(1),
	2573:  uint16(anon_sym_struct),
	2574:  uint16(113),
	2575:  uint16(1),
	2576:  uint16(anon_sym_union),
	2577:  uint16(117),
	2578:  uint16(1),
	2579:  uint16(anon_sym_RBRACE),
	2580:  uint16(177),
	2581:  uint16(1),
	2582:  uint16(sym_struct),
	2583:  uint16(179),
	2584:  uint16(1),
	2585:  uint16(sym_enum),
	2586:  uint16(182),
	2587:  uint16(2),
	2588:  uint16(sym__unnamed_union),
	2589:  uint16(sym__named_union),
	2590:  uint16(46),
	2591:  uint16(3),
	2592:  uint16(sym_using_directive),
	2593:  uint16(sym_field),
	2594:  uint16(aux_sym_struct_repeat1),
	2595:  uint16(165),
	2596:  uint16(7),
	2597:  uint16(sym_annotation),
	2598:  uint16(sym_nested_struct),
	2599:  uint16(sym_nested_enum),
	2600:  uint16(sym_group),
	2601:  uint16(sym_union),
	2602:  uint16(sym_interface),
	2603:  uint16(sym_const),
	2604:  uint16(15),
	2605:  uint16(3),
	2606:  uint16(1),
	2607:  uint16(sym_comment),
	2608:  uint16(119),
	2609:  uint16(1),
	2610:  uint16(sym_identifier),
	2611:  uint16(122),
	2612:  uint16(1),
	2613:  uint16(anon_sym_using),
	2614:  uint16(125),
	2615:  uint16(1),
	2616:  uint16(anon_sym_annotation),
	2617:  uint16(128),
	2618:  uint16(1),
	2619:  uint16(anon_sym_const),
	2620:  uint16(131),
	2621:  uint16(1),
	2622:  uint16(anon_sym_enum),
	2623:  uint16(134),
	2624:  uint16(1),
	2625:  uint16(anon_sym_interface),
	2626:  uint16(137),
	2627:  uint16(1),
	2628:  uint16(anon_sym_struct),
	2629:  uint16(140),
	2630:  uint16(1),
	2631:  uint16(anon_sym_union),
	2632:  uint16(143),
	2633:  uint16(1),
	2634:  uint16(anon_sym_RBRACE),
	2635:  uint16(177),
	2636:  uint16(1),
	2637:  uint16(sym_struct),
	2638:  uint16(179),
	2639:  uint16(1),
	2640:  uint16(sym_enum),
	2641:  uint16(182),
	2642:  uint16(2),
	2643:  uint16(sym__unnamed_union),
	2644:  uint16(sym__named_union),
	2645:  uint16(43),
	2646:  uint16(3),
	2647:  uint16(sym_using_directive),
	2648:  uint16(sym_field),
	2649:  uint16(aux_sym_struct_repeat1),
	2650:  uint16(165),
	2651:  uint16(7),
	2652:  uint16(sym_annotation),
	2653:  uint16(sym_nested_struct),
	2654:  uint16(sym_nested_enum),
	2655:  uint16(sym_group),
	2656:  uint16(sym_union),
	2657:  uint16(sym_interface),
	2658:  uint16(sym_const),
	2659:  uint16(15),
	2660:  uint16(3),
	2661:  uint16(1),
	2662:  uint16(sym_comment),
	2663:  uint16(99),
	2664:  uint16(1),
	2665:  uint16(sym_identifier),
	2666:  uint16(101),
	2667:  uint16(1),
	2668:  uint16(anon_sym_using),
	2669:  uint16(103),
	2670:  uint16(1),
	2671:  uint16(anon_sym_annotation),
	2672:  uint16(105),
	2673:  uint16(1),
	2674:  uint16(anon_sym_const),
	2675:  uint16(107),
	2676:  uint16(1),
	2677:  uint16(anon_sym_enum),
	2678:  uint16(109),
	2679:  uint16(1),
	2680:  uint16(anon_sym_interface),
	2681:  uint16(111),
	2682:  uint16(1),
	2683:  uint16(anon_sym_struct),
	2684:  uint16(113),
	2685:  uint16(1),
	2686:  uint16(anon_sym_union),
	2687:  uint16(115),
	2688:  uint16(1),
	2689:  uint16(anon_sym_RBRACE),
	2690:  uint16(177),
	2691:  uint16(1),
	2692:  uint16(sym_struct),
	2693:  uint16(179),
	2694:  uint16(1),
	2695:  uint16(sym_enum),
	2696:  uint16(182),
	2697:  uint16(2),
	2698:  uint16(sym__unnamed_union),
	2699:  uint16(sym__named_union),
	2700:  uint16(47),
	2701:  uint16(3),
	2702:  uint16(sym_using_directive),
	2703:  uint16(sym_field),
	2704:  uint16(aux_sym_struct_repeat1),
	2705:  uint16(165),
	2706:  uint16(7),
	2707:  uint16(sym_annotation),
	2708:  uint16(sym_nested_struct),
	2709:  uint16(sym_nested_enum),
	2710:  uint16(sym_group),
	2711:  uint16(sym_union),
	2712:  uint16(sym_interface),
	2713:  uint16(sym_const),
	2714:  uint16(15),
	2715:  uint16(3),
	2716:  uint16(1),
	2717:  uint16(sym_comment),
	2718:  uint16(99),
	2719:  uint16(1),
	2720:  uint16(sym_identifier),
	2721:  uint16(101),
	2722:  uint16(1),
	2723:  uint16(anon_sym_using),
	2724:  uint16(103),
	2725:  uint16(1),
	2726:  uint16(anon_sym_annotation),
	2727:  uint16(105),
	2728:  uint16(1),
	2729:  uint16(anon_sym_const),
	2730:  uint16(107),
	2731:  uint16(1),
	2732:  uint16(anon_sym_enum),
	2733:  uint16(109),
	2734:  uint16(1),
	2735:  uint16(anon_sym_interface),
	2736:  uint16(111),
	2737:  uint16(1),
	2738:  uint16(anon_sym_struct),
	2739:  uint16(113),
	2740:  uint16(1),
	2741:  uint16(anon_sym_union),
	2742:  uint16(145),
	2743:  uint16(1),
	2744:  uint16(anon_sym_RBRACE),
	2745:  uint16(177),
	2746:  uint16(1),
	2747:  uint16(sym_struct),
	2748:  uint16(179),
	2749:  uint16(1),
	2750:  uint16(sym_enum),
	2751:  uint16(182),
	2752:  uint16(2),
	2753:  uint16(sym__unnamed_union),
	2754:  uint16(sym__named_union),
	2755:  uint16(49),
	2756:  uint16(3),
	2757:  uint16(sym_using_directive),
	2758:  uint16(sym_field),
	2759:  uint16(aux_sym_struct_repeat1),
	2760:  uint16(165),
	2761:  uint16(7),
	2762:  uint16(sym_annotation),
	2763:  uint16(sym_nested_struct),
	2764:  uint16(sym_nested_enum),
	2765:  uint16(sym_group),
	2766:  uint16(sym_union),
	2767:  uint16(sym_interface),
	2768:  uint16(sym_const),
	2769:  uint16(15),
	2770:  uint16(3),
	2771:  uint16(1),
	2772:  uint16(sym_comment),
	2773:  uint16(99),
	2774:  uint16(1),
	2775:  uint16(sym_identifier),
	2776:  uint16(101),
	2777:  uint16(1),
	2778:  uint16(anon_sym_using),
	2779:  uint16(103),
	2780:  uint16(1),
	2781:  uint16(anon_sym_annotation),
	2782:  uint16(105),
	2783:  uint16(1),
	2784:  uint16(anon_sym_const),
	2785:  uint16(107),
	2786:  uint16(1),
	2787:  uint16(anon_sym_enum),
	2788:  uint16(109),
	2789:  uint16(1),
	2790:  uint16(anon_sym_interface),
	2791:  uint16(111),
	2792:  uint16(1),
	2793:  uint16(anon_sym_struct),
	2794:  uint16(113),
	2795:  uint16(1),
	2796:  uint16(anon_sym_union),
	2797:  uint16(147),
	2798:  uint16(1),
	2799:  uint16(anon_sym_RBRACE),
	2800:  uint16(177),
	2801:  uint16(1),
	2802:  uint16(sym_struct),
	2803:  uint16(179),
	2804:  uint16(1),
	2805:  uint16(sym_enum),
	2806:  uint16(182),
	2807:  uint16(2),
	2808:  uint16(sym__unnamed_union),
	2809:  uint16(sym__named_union),
	2810:  uint16(43),
	2811:  uint16(3),
	2812:  uint16(sym_using_directive),
	2813:  uint16(sym_field),
	2814:  uint16(aux_sym_struct_repeat1),
	2815:  uint16(165),
	2816:  uint16(7),
	2817:  uint16(sym_annotation),
	2818:  uint16(sym_nested_struct),
	2819:  uint16(sym_nested_enum),
	2820:  uint16(sym_group),
	2821:  uint16(sym_union),
	2822:  uint16(sym_interface),
	2823:  uint16(sym_const),
	2824:  uint16(15),
	2825:  uint16(3),
	2826:  uint16(1),
	2827:  uint16(sym_comment),
	2828:  uint16(99),
	2829:  uint16(1),
	2830:  uint16(sym_identifier),
	2831:  uint16(101),
	2832:  uint16(1),
	2833:  uint16(anon_sym_using),
	2834:  uint16(103),
	2835:  uint16(1),
	2836:  uint16(anon_sym_annotation),
	2837:  uint16(105),
	2838:  uint16(1),
	2839:  uint16(anon_sym_const),
	2840:  uint16(107),
	2841:  uint16(1),
	2842:  uint16(anon_sym_enum),
	2843:  uint16(109),
	2844:  uint16(1),
	2845:  uint16(anon_sym_interface),
	2846:  uint16(111),
	2847:  uint16(1),
	2848:  uint16(anon_sym_struct),
	2849:  uint16(113),
	2850:  uint16(1),
	2851:  uint16(anon_sym_union),
	2852:  uint16(117),
	2853:  uint16(1),
	2854:  uint16(anon_sym_RBRACE),
	2855:  uint16(177),
	2856:  uint16(1),
	2857:  uint16(sym_struct),
	2858:  uint16(179),
	2859:  uint16(1),
	2860:  uint16(sym_enum),
	2861:  uint16(182),
	2862:  uint16(2),
	2863:  uint16(sym__unnamed_union),
	2864:  uint16(sym__named_union),
	2865:  uint16(43),
	2866:  uint16(3),
	2867:  uint16(sym_using_directive),
	2868:  uint16(sym_field),
	2869:  uint16(aux_sym_struct_repeat1),
	2870:  uint16(165),
	2871:  uint16(7),
	2872:  uint16(sym_annotation),
	2873:  uint16(sym_nested_struct),
	2874:  uint16(sym_nested_enum),
	2875:  uint16(sym_group),
	2876:  uint16(sym_union),
	2877:  uint16(sym_interface),
	2878:  uint16(sym_const),
	2879:  uint16(15),
	2880:  uint16(3),
	2881:  uint16(1),
	2882:  uint16(sym_comment),
	2883:  uint16(99),
	2884:  uint16(1),
	2885:  uint16(sym_identifier),
	2886:  uint16(101),
	2887:  uint16(1),
	2888:  uint16(anon_sym_using),
	2889:  uint16(103),
	2890:  uint16(1),
	2891:  uint16(anon_sym_annotation),
	2892:  uint16(105),
	2893:  uint16(1),
	2894:  uint16(anon_sym_const),
	2895:  uint16(107),
	2896:  uint16(1),
	2897:  uint16(anon_sym_enum),
	2898:  uint16(109),
	2899:  uint16(1),
	2900:  uint16(anon_sym_interface),
	2901:  uint16(111),
	2902:  uint16(1),
	2903:  uint16(anon_sym_struct),
	2904:  uint16(113),
	2905:  uint16(1),
	2906:  uint16(anon_sym_union),
	2907:  uint16(149),
	2908:  uint16(1),
	2909:  uint16(anon_sym_RBRACE),
	2910:  uint16(177),
	2911:  uint16(1),
	2912:  uint16(sym_struct),
	2913:  uint16(179),
	2914:  uint16(1),
	2915:  uint16(sym_enum),
	2916:  uint16(182),
	2917:  uint16(2),
	2918:  uint16(sym__unnamed_union),
	2919:  uint16(sym__named_union),
	2920:  uint16(41),
	2921:  uint16(3),
	2922:  uint16(sym_using_directive),
	2923:  uint16(sym_field),
	2924:  uint16(aux_sym_struct_repeat1),
	2925:  uint16(165),
	2926:  uint16(7),
	2927:  uint16(sym_annotation),
	2928:  uint16(sym_nested_struct),
	2929:  uint16(sym_nested_enum),
	2930:  uint16(sym_group),
	2931:  uint16(sym_union),
	2932:  uint16(sym_interface),
	2933:  uint16(sym_const),
	2934:  uint16(15),
	2935:  uint16(3),
	2936:  uint16(1),
	2937:  uint16(sym_comment),
	2938:  uint16(99),
	2939:  uint16(1),
	2940:  uint16(sym_identifier),
	2941:  uint16(101),
	2942:  uint16(1),
	2943:  uint16(anon_sym_using),
	2944:  uint16(103),
	2945:  uint16(1),
	2946:  uint16(anon_sym_annotation),
	2947:  uint16(105),
	2948:  uint16(1),
	2949:  uint16(anon_sym_const),
	2950:  uint16(107),
	2951:  uint16(1),
	2952:  uint16(anon_sym_enum),
	2953:  uint16(109),
	2954:  uint16(1),
	2955:  uint16(anon_sym_interface),
	2956:  uint16(111),
	2957:  uint16(1),
	2958:  uint16(anon_sym_struct),
	2959:  uint16(113),
	2960:  uint16(1),
	2961:  uint16(anon_sym_union),
	2962:  uint16(149),
	2963:  uint16(1),
	2964:  uint16(anon_sym_RBRACE),
	2965:  uint16(177),
	2966:  uint16(1),
	2967:  uint16(sym_struct),
	2968:  uint16(179),
	2969:  uint16(1),
	2970:  uint16(sym_enum),
	2971:  uint16(182),
	2972:  uint16(2),
	2973:  uint16(sym__unnamed_union),
	2974:  uint16(sym__named_union),
	2975:  uint16(43),
	2976:  uint16(3),
	2977:  uint16(sym_using_directive),
	2978:  uint16(sym_field),
	2979:  uint16(aux_sym_struct_repeat1),
	2980:  uint16(165),
	2981:  uint16(7),
	2982:  uint16(sym_annotation),
	2983:  uint16(sym_nested_struct),
	2984:  uint16(sym_nested_enum),
	2985:  uint16(sym_group),
	2986:  uint16(sym_union),
	2987:  uint16(sym_interface),
	2988:  uint16(sym_const),
	2989:  uint16(8),
	2990:  uint16(3),
	2991:  uint16(1),
	2992:  uint16(sym_comment),
	2993:  uint16(35),
	2994:  uint16(1),
	2995:  uint16(anon_sym_List),
	2996:  uint16(151),
	2997:  uint16(1),
	2998:  uint16(sym_identifier),
	2999:  uint16(153),
	3000:  uint16(1),
	3001:  uint16(anon_sym_import),
	3002:  uint16(155),
	3003:  uint16(1),
	3004:  uint16(anon_sym_union),
	3005:  uint16(293),
	3006:  uint16(1),
	3007:  uint16(sym_field_type),
	3008:  uint16(238),
	3009:  uint16(3),
	3010:  uint16(sym_primitive_type),
	3011:  uint16(sym_list_type),
	3012:  uint16(sym_custom_type),
	3013:  uint16(33),
	3014:  uint16(15),
	3015:  uint16(anon_sym_AnyPointer),
	3016:  uint16(anon_sym_Bool),
	3017:  uint16(anon_sym_Int8),
	3018:  uint16(anon_sym_Int16),
	3019:  uint16(anon_sym_Int32),
	3020:  uint16(anon_sym_Int64),
	3021:  uint16(anon_sym_UInt8),
	3022:  uint16(anon_sym_UInt16),
	3023:  uint16(anon_sym_UInt32),
	3024:  uint16(anon_sym_UInt64),
	3025:  uint16(anon_sym_Float32),
	3026:  uint16(anon_sym_Float64),
	3027:  uint16(anon_sym_Text),
	3028:  uint16(anon_sym_Data),
	3029:  uint16(anon_sym_Void),
	3030:  uint16(14),
	3031:  uint16(3),
	3032:  uint16(1),
	3033:  uint16(sym_comment),
	3034:  uint16(7),
	3035:  uint16(1),
	3036:  uint16(sym_unique_id),
	3037:  uint16(9),
	3038:  uint16(1),
	3039:  uint16(anon_sym_using),
	3040:  uint16(11),
	3041:  uint16(1),
	3042:  uint16(anon_sym_DOLLARimport),
	3043:  uint16(13),
	3044:  uint16(1),
	3045:  uint16(anon_sym_DOLLAR),
	3046:  uint16(15),
	3047:  uint16(1),
	3048:  uint16(anon_sym_annotation),
	3049:  uint16(17),
	3050:  uint16(1),
	3051:  uint16(anon_sym_const),
	3052:  uint16(19),
	3053:  uint16(1),
	3054:  uint16(anon_sym_enum),
	3055:  uint16(21),
	3056:  uint16(1),
	3057:  uint16(anon_sym_interface),
	3058:  uint16(23),
	3059:  uint16(1),
	3060:  uint16(anon_sym_struct),
	3061:  uint16(157),
	3062:  uint16(1),
	3064:  uint16(52),
	3065:  uint16(2),
	3066:  uint16(sym_statement),
	3067:  uint16(aux_sym_message_repeat1),
	3068:  uint16(160),
	3069:  uint16(4),
	3070:  uint16(sym_struct),
	3071:  uint16(sym_enum),
	3072:  uint16(sym_interface),
	3073:  uint16(sym_const),
	3074:  uint16(161),
	3075:  uint16(6),
	3076:  uint16(sym_unique_id_statement),
	3077:  uint16(sym_using_directive),
	3078:  uint16(sym_import),
	3079:  uint16(sym_top_level_annotation),
	3080:  uint16(sym_annotation),
	3081:  uint16(sym_definition),
	3082:  uint16(14),
	3083:  uint16(3),
	3084:  uint16(1),
	3085:  uint16(sym_comment),
	3086:  uint16(159),
	3087:  uint16(1),
	3089:  uint16(161),
	3090:  uint16(1),
	3091:  uint16(sym_unique_id),
	3092:  uint16(164),
	3093:  uint16(1),
	3094:  uint16(anon_sym_using),
	3095:  uint16(167),
	3096:  uint16(1),
	3097:  uint16(anon_sym_DOLLARimport),
	3098:  uint16(170),
	3099:  uint16(1),
	3100:  uint16(anon_sym_DOLLAR),
	3101:  uint16(173),
	3102:  uint16(1),
	3103:  uint16(anon_sym_annotation),
	3104:  uint16(176),
	3105:  uint16(1),
	3106:  uint16(anon_sym_const),
	3107:  uint16(179),
	3108:  uint16(1),
	3109:  uint16(anon_sym_enum),
	3110:  uint16(182),
	3111:  uint16(1),
	3112:  uint16(anon_sym_interface),
	3113:  uint16(185),
	3114:  uint16(1),
	3115:  uint16(anon_sym_struct),
	3116:  uint16(52),
	3117:  uint16(2),
	3118:  uint16(sym_statement),
	3119:  uint16(aux_sym_message_repeat1),
	3120:  uint16(160),
	3121:  uint16(4),
	3122:  uint16(sym_struct),
	3123:  uint16(sym_enum),
	3124:  uint16(sym_interface),
	3125:  uint16(sym_const),
	3126:  uint16(161),
	3127:  uint16(6),
	3128:  uint16(sym_unique_id_statement),
	3129:  uint16(sym_using_directive),
	3130:  uint16(sym_import),
	3131:  uint16(sym_top_level_annotation),
	3132:  uint16(sym_annotation),
	3133:  uint16(sym_definition),
	3134:  uint16(7),
	3135:  uint16(3),
	3136:  uint16(1),
	3137:  uint16(sym_comment),
	3138:  uint16(35),
	3139:  uint16(1),
	3140:  uint16(anon_sym_List),
	3141:  uint16(151),
	3142:  uint16(1),
	3143:  uint16(sym_identifier),
	3144:  uint16(155),
	3145:  uint16(1),
	3146:  uint16(anon_sym_union),
	3147:  uint16(306),
	3148:  uint16(1),
	3149:  uint16(sym_field_type),
	3150:  uint16(238),
	3151:  uint16(3),
	3152:  uint16(sym_primitive_type),
	3153:  uint16(sym_list_type),
	3154:  uint16(sym_custom_type),
	3155:  uint16(33),
	3156:  uint16(15),
	3157:  uint16(anon_sym_AnyPointer),
	3158:  uint16(anon_sym_Bool),
	3159:  uint16(anon_sym_Int8),
	3160:  uint16(anon_sym_Int16),
	3161:  uint16(anon_sym_Int32),
	3162:  uint16(anon_sym_Int64),
	3163:  uint16(anon_sym_UInt8),
	3164:  uint16(anon_sym_UInt16),
	3165:  uint16(anon_sym_UInt32),
	3166:  uint16(anon_sym_UInt64),
	3167:  uint16(anon_sym_Float32),
	3168:  uint16(anon_sym_Float64),
	3169:  uint16(anon_sym_Text),
	3170:  uint16(anon_sym_Data),
	3171:  uint16(anon_sym_Void),
	3172:  uint16(7),
	3173:  uint16(3),
	3174:  uint16(1),
	3175:  uint16(sym_comment),
	3176:  uint16(35),
	3177:  uint16(1),
	3178:  uint16(anon_sym_List),
	3179:  uint16(151),
	3180:  uint16(1),
	3181:  uint16(sym_identifier),
	3182:  uint16(294),
	3183:  uint16(1),
	3184:  uint16(sym_field_type),
	3185:  uint16(537),
	3186:  uint16(1),
	3187:  uint16(sym_generic_parameters),
	3188:  uint16(238),
	3189:  uint16(3),
	3190:  uint16(sym_primitive_type),
	3191:  uint16(sym_list_type),
	3192:  uint16(sym_custom_type),
	3193:  uint16(33),
	3194:  uint16(15),
	3195:  uint16(anon_sym_AnyPointer),
	3196:  uint16(anon_sym_Bool),
	3197:  uint16(anon_sym_Int8),
	3198:  uint16(anon_sym_Int16),
	3199:  uint16(anon_sym_Int32),
	3200:  uint16(anon_sym_Int64),
	3201:  uint16(anon_sym_UInt8),
	3202:  uint16(anon_sym_UInt16),
	3203:  uint16(anon_sym_UInt32),
	3204:  uint16(anon_sym_UInt64),
	3205:  uint16(anon_sym_Float32),
	3206:  uint16(anon_sym_Float64),
	3207:  uint16(anon_sym_Text),
	3208:  uint16(anon_sym_Data),
	3209:  uint16(anon_sym_Void),
	3210:  uint16(7),
	3211:  uint16(3),
	3212:  uint16(1),
	3213:  uint16(sym_comment),
	3214:  uint16(35),
	3215:  uint16(1),
	3216:  uint16(anon_sym_List),
	3217:  uint16(151),
	3218:  uint16(1),
	3219:  uint16(sym_identifier),
	3220:  uint16(294),
	3221:  uint16(1),
	3222:  uint16(sym_field_type),
	3223:  uint16(478),
	3224:  uint16(1),
	3225:  uint16(sym_generic_parameters),
	3226:  uint16(238),
	3227:  uint16(3),
	3228:  uint16(sym_primitive_type),
	3229:  uint16(sym_list_type),
	3230:  uint16(sym_custom_type),
	3231:  uint16(33),
	3232:  uint16(15),
	3233:  uint16(anon_sym_AnyPointer),
	3234:  uint16(anon_sym_Bool),
	3235:  uint16(anon_sym_Int8),
	3236:  uint16(anon_sym_Int16),
	3237:  uint16(anon_sym_Int32),
	3238:  uint16(anon_sym_Int64),
	3239:  uint16(anon_sym_UInt8),
	3240:  uint16(anon_sym_UInt16),
	3241:  uint16(anon_sym_UInt32),
	3242:  uint16(anon_sym_UInt64),
	3243:  uint16(anon_sym_Float32),
	3244:  uint16(anon_sym_Float64),
	3245:  uint16(anon_sym_Text),
	3246:  uint16(anon_sym_Data),
	3247:  uint16(anon_sym_Void),
	3248:  uint16(14),
	3249:  uint16(3),
	3250:  uint16(1),
	3251:  uint16(sym_comment),
	3252:  uint16(99),
	3253:  uint16(1),
	3254:  uint16(sym_identifier),
	3255:  uint16(103),
	3256:  uint16(1),
	3257:  uint16(anon_sym_annotation),
	3258:  uint16(105),
	3259:  uint16(1),
	3260:  uint16(anon_sym_const),
	3261:  uint16(107),
	3262:  uint16(1),
	3263:  uint16(anon_sym_enum),
	3264:  uint16(109),
	3265:  uint16(1),
	3266:  uint16(anon_sym_interface),
	3267:  uint16(111),
	3268:  uint16(1),
	3269:  uint16(anon_sym_struct),
	3270:  uint16(113),
	3271:  uint16(1),
	3272:  uint16(anon_sym_union),
	3273:  uint16(188),
	3274:  uint16(1),
	3275:  uint16(anon_sym_RBRACE),
	3276:  uint16(177),
	3277:  uint16(1),
	3278:  uint16(sym_struct),
	3279:  uint16(179),
	3280:  uint16(1),
	3281:  uint16(sym_enum),
	3282:  uint16(66),
	3283:  uint16(2),
	3284:  uint16(sym_field),
	3285:  uint16(aux_sym_group_repeat1),
	3286:  uint16(182),
	3287:  uint16(2),
	3288:  uint16(sym__unnamed_union),
	3289:  uint16(sym__named_union),
	3290:  uint16(165),
	3291:  uint16(7),
	3292:  uint16(sym_annotation),
	3293:  uint16(sym_nested_struct),
	3294:  uint16(sym_nested_enum),
	3295:  uint16(sym_group),
	3296:  uint16(sym_union),
	3297:  uint16(sym_interface),
	3298:  uint16(sym_const),
	3299:  uint16(6),
	3300:  uint16(3),
	3301:  uint16(1),
	3302:  uint16(sym_comment),
	3303:  uint16(35),
	3304:  uint16(1),
	3305:  uint16(anon_sym_List),
	3306:  uint16(151),
	3307:  uint16(1),
	3308:  uint16(sym_identifier),
	3309:  uint16(482),
	3310:  uint16(1),
	3311:  uint16(sym_field_type),
	3312:  uint16(238),
	3313:  uint16(3),
	3314:  uint16(sym_primitive_type),
	3315:  uint16(sym_list_type),
	3316:  uint16(sym_custom_type),
	3317:  uint16(33),
	3318:  uint16(15),
	3319:  uint16(anon_sym_AnyPointer),
	3320:  uint16(anon_sym_Bool),
	3321:  uint16(anon_sym_Int8),
	3322:  uint16(anon_sym_Int16),
	3323:  uint16(anon_sym_Int32),
	3324:  uint16(anon_sym_Int64),
	3325:  uint16(anon_sym_UInt8),
	3326:  uint16(anon_sym_UInt16),
	3327:  uint16(anon_sym_UInt32),
	3328:  uint16(anon_sym_UInt64),
	3329:  uint16(anon_sym_Float32),
	3330:  uint16(anon_sym_Float64),
	3331:  uint16(anon_sym_Text),
	3332:  uint16(anon_sym_Data),
	3333:  uint16(anon_sym_Void),
	3334:  uint16(14),
	3335:  uint16(3),
	3336:  uint16(1),
	3337:  uint16(sym_comment),
	3338:  uint16(99),
	3339:  uint16(1),
	3340:  uint16(sym_identifier),
	3341:  uint16(103),
	3342:  uint16(1),
	3343:  uint16(anon_sym_annotation),
	3344:  uint16(105),
	3345:  uint16(1),
	3346:  uint16(anon_sym_const),
	3347:  uint16(107),
	3348:  uint16(1),
	3349:  uint16(anon_sym_enum),
	3350:  uint16(109),
	3351:  uint16(1),
	3352:  uint16(anon_sym_interface),
	3353:  uint16(111),
	3354:  uint16(1),
	3355:  uint16(anon_sym_struct),
	3356:  uint16(113),
	3357:  uint16(1),
	3358:  uint16(anon_sym_union),
	3359:  uint16(190),
	3360:  uint16(1),
	3361:  uint16(anon_sym_RBRACE),
	3362:  uint16(177),
	3363:  uint16(1),
	3364:  uint16(sym_struct),
	3365:  uint16(179),
	3366:  uint16(1),
	3367:  uint16(sym_enum),
	3368:  uint16(59),
	3369:  uint16(2),
	3370:  uint16(sym_field),
	3371:  uint16(aux_sym_group_repeat1),
	3372:  uint16(182),
	3373:  uint16(2),
	3374:  uint16(sym__unnamed_union),
	3375:  uint16(sym__named_union),
	3376:  uint16(165),
	3377:  uint16(7),
	3378:  uint16(sym_annotation),
	3379:  uint16(sym_nested_struct),
	3380:  uint16(sym_nested_enum),
	3381:  uint16(sym_group),
	3382:  uint16(sym_union),
	3383:  uint16(sym_interface),
	3384:  uint16(sym_const),
	3385:  uint16(14),
	3386:  uint16(3),
	3387:  uint16(1),
	3388:  uint16(sym_comment),
	3389:  uint16(99),
	3390:  uint16(1),
	3391:  uint16(sym_identifier),
	3392:  uint16(103),
	3393:  uint16(1),
	3394:  uint16(anon_sym_annotation),
	3395:  uint16(105),
	3396:  uint16(1),
	3397:  uint16(anon_sym_const),
	3398:  uint16(107),
	3399:  uint16(1),
	3400:  uint16(anon_sym_enum),
	3401:  uint16(109),
	3402:  uint16(1),
	3403:  uint16(anon_sym_interface),
	3404:  uint16(111),
	3405:  uint16(1),
	3406:  uint16(anon_sym_struct),
	3407:  uint16(113),
	3408:  uint16(1),
	3409:  uint16(anon_sym_union),
	3410:  uint16(188),
	3411:  uint16(1),
	3412:  uint16(anon_sym_RBRACE),
	3413:  uint16(177),
	3414:  uint16(1),
	3415:  uint16(sym_struct),
	3416:  uint16(179),
	3417:  uint16(1),
	3418:  uint16(sym_enum),
	3419:  uint16(62),
	3420:  uint16(2),
	3421:  uint16(sym_field),
	3422:  uint16(aux_sym_group_repeat1),
	3423:  uint16(182),
	3424:  uint16(2),
	3425:  uint16(sym__unnamed_union),
	3426:  uint16(sym__named_union),
	3427:  uint16(165),
	3428:  uint16(7),
	3429:  uint16(sym_annotation),
	3430:  uint16(sym_nested_struct),
	3431:  uint16(sym_nested_enum),
	3432:  uint16(sym_group),
	3433:  uint16(sym_union),
	3434:  uint16(sym_interface),
	3435:  uint16(sym_const),
	3436:  uint16(6),
	3437:  uint16(3),
	3438:  uint16(1),
	3439:  uint16(sym_comment),
	3440:  uint16(35),
	3441:  uint16(1),
	3442:  uint16(anon_sym_List),
	3443:  uint16(151),
	3444:  uint16(1),
	3445:  uint16(sym_identifier),
	3446:  uint16(403),
	3447:  uint16(1),
	3448:  uint16(sym_field_type),
	3449:  uint16(238),
	3450:  uint16(3),
	3451:  uint16(sym_primitive_type),
	3452:  uint16(sym_list_type),
	3453:  uint16(sym_custom_type),
	3454:  uint16(33),
	3455:  uint16(15),
	3456:  uint16(anon_sym_AnyPointer),
	3457:  uint16(anon_sym_Bool),
	3458:  uint16(anon_sym_Int8),
	3459:  uint16(anon_sym_Int16),
	3460:  uint16(anon_sym_Int32),
	3461:  uint16(anon_sym_Int64),
	3462:  uint16(anon_sym_UInt8),
	3463:  uint16(anon_sym_UInt16),
	3464:  uint16(anon_sym_UInt32),
	3465:  uint16(anon_sym_UInt64),
	3466:  uint16(anon_sym_Float32),
	3467:  uint16(anon_sym_Float64),
	3468:  uint16(anon_sym_Text),
	3469:  uint16(anon_sym_Data),
	3470:  uint16(anon_sym_Void),
	3471:  uint16(6),
	3472:  uint16(3),
	3473:  uint16(1),
	3474:  uint16(sym_comment),
	3475:  uint16(35),
	3476:  uint16(1),
	3477:  uint16(anon_sym_List),
	3478:  uint16(151),
	3479:  uint16(1),
	3480:  uint16(sym_identifier),
	3481:  uint16(306),
	3482:  uint16(1),
	3483:  uint16(sym_field_type),
	3484:  uint16(238),
	3485:  uint16(3),
	3486:  uint16(sym_primitive_type),
	3487:  uint16(sym_list_type),
	3488:  uint16(sym_custom_type),
	3489:  uint16(33),
	3490:  uint16(15),
	3491:  uint16(anon_sym_AnyPointer),
	3492:  uint16(anon_sym_Bool),
	3493:  uint16(anon_sym_Int8),
	3494:  uint16(anon_sym_Int16),
	3495:  uint16(anon_sym_Int32),
	3496:  uint16(anon_sym_Int64),
	3497:  uint16(anon_sym_UInt8),
	3498:  uint16(anon_sym_UInt16),
	3499:  uint16(anon_sym_UInt32),
	3500:  uint16(anon_sym_UInt64),
	3501:  uint16(anon_sym_Float32),
	3502:  uint16(anon_sym_Float64),
	3503:  uint16(anon_sym_Text),
	3504:  uint16(anon_sym_Data),
	3505:  uint16(anon_sym_Void),
	3506:  uint16(14),
	3507:  uint16(3),
	3508:  uint16(1),
	3509:  uint16(sym_comment),
	3510:  uint16(192),
	3511:  uint16(1),
	3512:  uint16(sym_identifier),
	3513:  uint16(195),
	3514:  uint16(1),
	3515:  uint16(anon_sym_annotation),
	3516:  uint16(198),
	3517:  uint16(1),
	3518:  uint16(anon_sym_const),
	3519:  uint16(201),
	3520:  uint16(1),
	3521:  uint16(anon_sym_enum),
	3522:  uint16(204),
	3523:  uint16(1),
	3524:  uint16(anon_sym_interface),
	3525:  uint16(207),
	3526:  uint16(1),
	3527:  uint16(anon_sym_struct),
	3528:  uint16(210),
	3529:  uint16(1),
	3530:  uint16(anon_sym_union),
	3531:  uint16(213),
	3532:  uint16(1),
	3533:  uint16(anon_sym_RBRACE),
	3534:  uint16(177),
	3535:  uint16(1),
	3536:  uint16(sym_struct),
	3537:  uint16(179),
	3538:  uint16(1),
	3539:  uint16(sym_enum),
	3540:  uint16(62),
	3541:  uint16(2),
	3542:  uint16(sym_field),
	3543:  uint16(aux_sym_group_repeat1),
	3544:  uint16(182),
	3545:  uint16(2),
	3546:  uint16(sym__unnamed_union),
	3547:  uint16(sym__named_union),
	3548:  uint16(165),
	3549:  uint16(7),
	3550:  uint16(sym_annotation),
	3551:  uint16(sym_nested_struct),
	3552:  uint16(sym_nested_enum),
	3553:  uint16(sym_group),
	3554:  uint16(sym_union),
	3555:  uint16(sym_interface),
	3556:  uint16(sym_const),
	3557:  uint16(6),
	3558:  uint16(3),
	3559:  uint16(1),
	3560:  uint16(sym_comment),
	3561:  uint16(35),
	3562:  uint16(1),
	3563:  uint16(anon_sym_List),
	3564:  uint16(151),
	3565:  uint16(1),
	3566:  uint16(sym_identifier),
	3567:  uint16(268),
	3568:  uint16(1),
	3569:  uint16(sym_field_type),
	3570:  uint16(238),
	3571:  uint16(3),
	3572:  uint16(sym_primitive_type),
	3573:  uint16(sym_list_type),
	3574:  uint16(sym_custom_type),
	3575:  uint16(33),
	3576:  uint16(15),
	3577:  uint16(anon_sym_AnyPointer),
	3578:  uint16(anon_sym_Bool),
	3579:  uint16(anon_sym_Int8),
	3580:  uint16(anon_sym_Int16),
	3581:  uint16(anon_sym_Int32),
	3582:  uint16(anon_sym_Int64),
	3583:  uint16(anon_sym_UInt8),
	3584:  uint16(anon_sym_UInt16),
	3585:  uint16(anon_sym_UInt32),
	3586:  uint16(anon_sym_UInt64),
	3587:  uint16(anon_sym_Float32),
	3588:  uint16(anon_sym_Float64),
	3589:  uint16(anon_sym_Text),
	3590:  uint16(anon_sym_Data),
	3591:  uint16(anon_sym_Void),
	3592:  uint16(6),
	3593:  uint16(3),
	3594:  uint16(1),
	3595:  uint16(sym_comment),
	3596:  uint16(35),
	3597:  uint16(1),
	3598:  uint16(anon_sym_List),
	3599:  uint16(151),
	3600:  uint16(1),
	3601:  uint16(sym_identifier),
	3602:  uint16(439),
	3603:  uint16(1),
	3604:  uint16(sym_field_type),
	3605:  uint16(238),
	3606:  uint16(3),
	3607:  uint16(sym_primitive_type),
	3608:  uint16(sym_list_type),
	3609:  uint16(sym_custom_type),
	3610:  uint16(33),
	3611:  uint16(15),
	3612:  uint16(anon_sym_AnyPointer),
	3613:  uint16(anon_sym_Bool),
	3614:  uint16(anon_sym_Int8),
	3615:  uint16(anon_sym_Int16),
	3616:  uint16(anon_sym_Int32),
	3617:  uint16(anon_sym_Int64),
	3618:  uint16(anon_sym_UInt8),
	3619:  uint16(anon_sym_UInt16),
	3620:  uint16(anon_sym_UInt32),
	3621:  uint16(anon_sym_UInt64),
	3622:  uint16(anon_sym_Float32),
	3623:  uint16(anon_sym_Float64),
	3624:  uint16(anon_sym_Text),
	3625:  uint16(anon_sym_Data),
	3626:  uint16(anon_sym_Void),
	3627:  uint16(6),
	3628:  uint16(3),
	3629:  uint16(1),
	3630:  uint16(sym_comment),
	3631:  uint16(35),
	3632:  uint16(1),
	3633:  uint16(anon_sym_List),
	3634:  uint16(151),
	3635:  uint16(1),
	3636:  uint16(sym_identifier),
	3637:  uint16(383),
	3638:  uint16(1),
	3639:  uint16(sym_field_type),
	3640:  uint16(238),
	3641:  uint16(3),
	3642:  uint16(sym_primitive_type),
	3643:  uint16(sym_list_type),
	3644:  uint16(sym_custom_type),
	3645:  uint16(33),
	3646:  uint16(15),
	3647:  uint16(anon_sym_AnyPointer),
	3648:  uint16(anon_sym_Bool),
	3649:  uint16(anon_sym_Int8),
	3650:  uint16(anon_sym_Int16),
	3651:  uint16(anon_sym_Int32),
	3652:  uint16(anon_sym_Int64),
	3653:  uint16(anon_sym_UInt8),
	3654:  uint16(anon_sym_UInt16),
	3655:  uint16(anon_sym_UInt32),
	3656:  uint16(anon_sym_UInt64),
	3657:  uint16(anon_sym_Float32),
	3658:  uint16(anon_sym_Float64),
	3659:  uint16(anon_sym_Text),
	3660:  uint16(anon_sym_Data),
	3661:  uint16(anon_sym_Void),
	3662:  uint16(14),
	3663:  uint16(3),
	3664:  uint16(1),
	3665:  uint16(sym_comment),
	3666:  uint16(99),
	3667:  uint16(1),
	3668:  uint16(sym_identifier),
	3669:  uint16(103),
	3670:  uint16(1),
	3671:  uint16(anon_sym_annotation),
	3672:  uint16(105),
	3673:  uint16(1),
	3674:  uint16(anon_sym_const),
	3675:  uint16(107),
	3676:  uint16(1),
	3677:  uint16(anon_sym_enum),
	3678:  uint16(109),
	3679:  uint16(1),
	3680:  uint16(anon_sym_interface),
	3681:  uint16(111),
	3682:  uint16(1),
	3683:  uint16(anon_sym_struct),
	3684:  uint16(113),
	3685:  uint16(1),
	3686:  uint16(anon_sym_union),
	3687:  uint16(215),
	3688:  uint16(1),
	3689:  uint16(anon_sym_RBRACE),
	3690:  uint16(177),
	3691:  uint16(1),
	3692:  uint16(sym_struct),
	3693:  uint16(179),
	3694:  uint16(1),
	3695:  uint16(sym_enum),
	3696:  uint16(62),
	3697:  uint16(2),
	3698:  uint16(sym_field),
	3699:  uint16(aux_sym_group_repeat1),
	3700:  uint16(182),
	3701:  uint16(2),
	3702:  uint16(sym__unnamed_union),
	3703:  uint16(sym__named_union),
	3704:  uint16(165),
	3705:  uint16(7),
	3706:  uint16(sym_annotation),
	3707:  uint16(sym_nested_struct),
	3708:  uint16(sym_nested_enum),
	3709:  uint16(sym_group),
	3710:  uint16(sym_union),
	3711:  uint16(sym_interface),
	3712:  uint16(sym_const),
	3713:  uint16(6),
	3714:  uint16(3),
	3715:  uint16(1),
	3716:  uint16(sym_comment),
	3717:  uint16(35),
	3718:  uint16(1),
	3719:  uint16(anon_sym_List),
	3720:  uint16(151),
	3721:  uint16(1),
	3722:  uint16(sym_identifier),
	3723:  uint16(496),
	3724:  uint16(1),
	3725:  uint16(sym_field_type),
	3726:  uint16(238),
	3727:  uint16(3),
	3728:  uint16(sym_primitive_type),
	3729:  uint16(sym_list_type),
	3730:  uint16(sym_custom_type),
	3731:  uint16(33),
	3732:  uint16(15),
	3733:  uint16(anon_sym_AnyPointer),
	3734:  uint16(anon_sym_Bool),
	3735:  uint16(anon_sym_Int8),
	3736:  uint16(anon_sym_Int16),
	3737:  uint16(anon_sym_Int32),
	3738:  uint16(anon_sym_Int64),
	3739:  uint16(anon_sym_UInt8),
	3740:  uint16(anon_sym_UInt16),
	3741:  uint16(anon_sym_UInt32),
	3742:  uint16(anon_sym_UInt64),
	3743:  uint16(anon_sym_Float32),
	3744:  uint16(anon_sym_Float64),
	3745:  uint16(anon_sym_Text),
	3746:  uint16(anon_sym_Data),
	3747:  uint16(anon_sym_Void),
	3748:  uint16(6),
	3749:  uint16(3),
	3750:  uint16(1),
	3751:  uint16(sym_comment),
	3752:  uint16(35),
	3753:  uint16(1),
	3754:  uint16(anon_sym_List),
	3755:  uint16(151),
	3756:  uint16(1),
	3757:  uint16(sym_identifier),
	3758:  uint16(485),
	3759:  uint16(1),
	3760:  uint16(sym_field_type),
	3761:  uint16(238),
	3762:  uint16(3),
	3763:  uint16(sym_primitive_type),
	3764:  uint16(sym_list_type),
	3765:  uint16(sym_custom_type),
	3766:  uint16(33),
	3767:  uint16(15),
	3768:  uint16(anon_sym_AnyPointer),
	3769:  uint16(anon_sym_Bool),
	3770:  uint16(anon_sym_Int8),
	3771:  uint16(anon_sym_Int16),
	3772:  uint16(anon_sym_Int32),
	3773:  uint16(anon_sym_Int64),
	3774:  uint16(anon_sym_UInt8),
	3775:  uint16(anon_sym_UInt16),
	3776:  uint16(anon_sym_UInt32),
	3777:  uint16(anon_sym_UInt64),
	3778:  uint16(anon_sym_Float32),
	3779:  uint16(anon_sym_Float64),
	3780:  uint16(anon_sym_Text),
	3781:  uint16(anon_sym_Data),
	3782:  uint16(anon_sym_Void),
	3783:  uint16(4),
	3784:  uint16(3),
	3785:  uint16(1),
	3786:  uint16(sym_comment),
	3787:  uint16(219),
	3788:  uint16(1),
	3789:  uint16(anon_sym_enum),
	3790:  uint16(471),
	3791:  uint16(1),
	3792:  uint16(sym_annotation_target),
	3793:  uint16(217),
	3794:  uint16(12),
	3795:  uint16(anon_sym_annotation),
	3796:  uint16(anon_sym_STAR),
	3797:  uint16(anon_sym_const),
	3798:  uint16(anon_sym_enumerant),
	3799:  uint16(anon_sym_field),
	3800:  uint16(anon_sym_file),
	3801:  uint16(anon_sym_method),
	3802:  uint16(anon_sym_param),
	3803:  uint16(anon_sym_group),
	3804:  uint16(anon_sym_interface),
	3805:  uint16(anon_sym_struct),
	3806:  uint16(anon_sym_union),
	3807:  uint16(4),
	3808:  uint16(3),
	3809:  uint16(1),
	3810:  uint16(sym_comment),
	3811:  uint16(219),
	3812:  uint16(1),
	3813:  uint16(anon_sym_enum),
	3814:  uint16(396),
	3815:  uint16(1),
	3816:  uint16(sym_annotation_target),
	3817:  uint16(217),
	3818:  uint16(12),
	3819:  uint16(anon_sym_annotation),
	3820:  uint16(anon_sym_STAR),
	3821:  uint16(anon_sym_const),
	3822:  uint16(anon_sym_enumerant),
	3823:  uint16(anon_sym_field),
	3824:  uint16(anon_sym_file),
	3825:  uint16(anon_sym_method),
	3826:  uint16(anon_sym_param),
	3827:  uint16(anon_sym_group),
	3828:  uint16(anon_sym_interface),
	3829:  uint16(anon_sym_struct),
	3830:  uint16(anon_sym_union),
	3831:  uint16(3),
	3832:  uint16(3),
	3833:  uint16(1),
	3834:  uint16(sym_comment),
	3835:  uint16(221),
	3836:  uint16(4),
	3838:  uint16(sym_unique_id),
	3839:  uint16(anon_sym_DOLLARimport),
	3840:  uint16(anon_sym_RBRACE),
	3841:  uint16(223),
	3842:  uint16(9),
	3843:  uint16(anon_sym_using),
	3844:  uint16(anon_sym_DOLLAR),
	3845:  uint16(anon_sym_annotation),
	3846:  uint16(anon_sym_const),
	3847:  uint16(anon_sym_enum),
	3848:  uint16(anon_sym_interface),
	3849:  uint16(anon_sym_struct),
	3850:  uint16(anon_sym_union),
	3851:  uint16(sym_identifier),
	3852:  uint16(3),
	3853:  uint16(3),
	3854:  uint16(1),
	3855:  uint16(sym_comment),
	3856:  uint16(225),
	3857:  uint16(4),
	3859:  uint16(sym_unique_id),
	3860:  uint16(anon_sym_DOLLARimport),
	3861:  uint16(anon_sym_RBRACE),
	3862:  uint16(227),
	3863:  uint16(9),
	3864:  uint16(anon_sym_using),
	3865:  uint16(anon_sym_DOLLAR),
	3866:  uint16(anon_sym_annotation),
	3867:  uint16(anon_sym_const),
	3868:  uint16(anon_sym_enum),
	3869:  uint16(anon_sym_interface),
	3870:  uint16(anon_sym_struct),
	3871:  uint16(anon_sym_union),
	3872:  uint16(sym_identifier),
	3873:  uint16(3),
	3874:  uint16(3),
	3875:  uint16(1),
	3876:  uint16(sym_comment),
	3877:  uint16(229),
	3878:  uint16(4),
	3880:  uint16(sym_unique_id),
	3881:  uint16(anon_sym_DOLLARimport),
	3882:  uint16(anon_sym_RBRACE),
	3883:  uint16(231),
	3884:  uint16(9),
	3885:  uint16(anon_sym_using),
	3886:  uint16(anon_sym_DOLLAR),
	3887:  uint16(anon_sym_annotation),
	3888:  uint16(anon_sym_const),
	3889:  uint16(anon_sym_enum),
	3890:  uint16(anon_sym_interface),
	3891:  uint16(anon_sym_struct),
	3892:  uint16(anon_sym_union),
	3893:  uint16(sym_identifier),
	3894:  uint16(3),
	3895:  uint16(3),
	3896:  uint16(1),
	3897:  uint16(sym_comment),
	3898:  uint16(233),
	3899:  uint16(4),
	3901:  uint16(sym_unique_id),
	3902:  uint16(anon_sym_DOLLARimport),
	3903:  uint16(anon_sym_RBRACE),
	3904:  uint16(235),
	3905:  uint16(9),
	3906:  uint16(anon_sym_using),
	3907:  uint16(anon_sym_DOLLAR),
	3908:  uint16(anon_sym_annotation),
	3909:  uint16(anon_sym_const),
	3910:  uint16(anon_sym_enum),
	3911:  uint16(anon_sym_interface),
	3912:  uint16(anon_sym_struct),
	3913:  uint16(anon_sym_union),
	3914:  uint16(sym_identifier),
	3915:  uint16(3),
	3916:  uint16(3),
	3917:  uint16(1),
	3918:  uint16(sym_comment),
	3919:  uint16(237),
	3920:  uint16(4),
	3922:  uint16(sym_unique_id),
	3923:  uint16(anon_sym_DOLLARimport),
	3924:  uint16(anon_sym_RBRACE),
	3925:  uint16(239),
	3926:  uint16(9),
	3927:  uint16(anon_sym_using),
	3928:  uint16(anon_sym_DOLLAR),
	3929:  uint16(anon_sym_annotation),
	3930:  uint16(anon_sym_const),
	3931:  uint16(anon_sym_enum),
	3932:  uint16(anon_sym_interface),
	3933:  uint16(anon_sym_struct),
	3934:  uint16(anon_sym_union),
	3935:  uint16(sym_identifier),
	3936:  uint16(3),
	3937:  uint16(3),
	3938:  uint16(1),
	3939:  uint16(sym_comment),
	3940:  uint16(241),
	3941:  uint16(4),
	3943:  uint16(sym_unique_id),
	3944:  uint16(anon_sym_DOLLARimport),
	3945:  uint16(anon_sym_RBRACE),
	3946:  uint16(243),
	3947:  uint16(9),
	3948:  uint16(anon_sym_using),
	3949:  uint16(anon_sym_DOLLAR),
	3950:  uint16(anon_sym_annotation),
	3951:  uint16(anon_sym_const),
	3952:  uint16(anon_sym_enum),
	3953:  uint16(anon_sym_interface),
	3954:  uint16(anon_sym_struct),
	3955:  uint16(anon_sym_union),
	3956:  uint16(sym_identifier),
	3957:  uint16(3),
	3958:  uint16(3),
	3959:  uint16(1),
	3960:  uint16(sym_comment),
	3961:  uint16(245),
	3962:  uint16(4),
	3964:  uint16(sym_unique_id),
	3965:  uint16(anon_sym_DOLLARimport),
	3966:  uint16(anon_sym_RBRACE),
	3967:  uint16(247),
	3968:  uint16(9),
	3969:  uint16(anon_sym_using),
	3970:  uint16(anon_sym_DOLLAR),
	3971:  uint16(anon_sym_annotation),
	3972:  uint16(anon_sym_const),
	3973:  uint16(anon_sym_enum),
	3974:  uint16(anon_sym_interface),
	3975:  uint16(anon_sym_struct),
	3976:  uint16(anon_sym_union),
	3977:  uint16(sym_identifier),
	3978:  uint16(3),
	3979:  uint16(3),
	3980:  uint16(1),
	3981:  uint16(sym_comment),
	3982:  uint16(249),
	3983:  uint16(4),
	3985:  uint16(sym_unique_id),
	3986:  uint16(anon_sym_DOLLARimport),
	3987:  uint16(anon_sym_RBRACE),
	3988:  uint16(251),
	3989:  uint16(9),
	3990:  uint16(anon_sym_using),
	3991:  uint16(anon_sym_DOLLAR),
	3992:  uint16(anon_sym_annotation),
	3993:  uint16(anon_sym_const),
	3994:  uint16(anon_sym_enum),
	3995:  uint16(anon_sym_interface),
	3996:  uint16(anon_sym_struct),
	3997:  uint16(anon_sym_union),
	3998:  uint16(sym_identifier),
	3999:  uint16(3),
	4000:  uint16(3),
	4001:  uint16(1),
	4002:  uint16(sym_comment),
	4003:  uint16(253),
	4004:  uint16(4),
	4006:  uint16(sym_unique_id),
	4007:  uint16(anon_sym_DOLLARimport),
	4008:  uint16(anon_sym_RBRACE),
	4009:  uint16(255),
	4010:  uint16(9),
	4011:  uint16(anon_sym_using),
	4012:  uint16(anon_sym_DOLLAR),
	4013:  uint16(anon_sym_annotation),
	4014:  uint16(anon_sym_const),
	4015:  uint16(anon_sym_enum),
	4016:  uint16(anon_sym_interface),
	4017:  uint16(anon_sym_struct),
	4018:  uint16(anon_sym_union),
	4019:  uint16(sym_identifier),
	4020:  uint16(3),
	4021:  uint16(3),
	4022:  uint16(1),
	4023:  uint16(sym_comment),
	4024:  uint16(257),
	4025:  uint16(4),
	4027:  uint16(sym_unique_id),
	4028:  uint16(anon_sym_DOLLARimport),
	4029:  uint16(anon_sym_RBRACE),
	4030:  uint16(259),
	4031:  uint16(9),
	4032:  uint16(anon_sym_using),
	4033:  uint16(anon_sym_DOLLAR),
	4034:  uint16(anon_sym_annotation),
	4035:  uint16(anon_sym_const),
	4036:  uint16(anon_sym_enum),
	4037:  uint16(anon_sym_interface),
	4038:  uint16(anon_sym_struct),
	4039:  uint16(anon_sym_union),
	4040:  uint16(sym_identifier),
	4041:  uint16(3),
	4042:  uint16(3),
	4043:  uint16(1),
	4044:  uint16(sym_comment),
	4045:  uint16(261),
	4046:  uint16(4),
	4048:  uint16(sym_unique_id),
	4049:  uint16(anon_sym_DOLLARimport),
	4050:  uint16(anon_sym_RBRACE),
	4051:  uint16(263),
	4052:  uint16(9),
	4053:  uint16(anon_sym_using),
	4054:  uint16(anon_sym_DOLLAR),
	4055:  uint16(anon_sym_annotation),
	4056:  uint16(anon_sym_const),
	4057:  uint16(anon_sym_enum),
	4058:  uint16(anon_sym_interface),
	4059:  uint16(anon_sym_struct),
	4060:  uint16(anon_sym_union),
	4061:  uint16(sym_identifier),
	4062:  uint16(3),
	4063:  uint16(3),
	4064:  uint16(1),
	4065:  uint16(sym_comment),
	4066:  uint16(265),
	4067:  uint16(4),
	4069:  uint16(sym_unique_id),
	4070:  uint16(anon_sym_DOLLARimport),
	4071:  uint16(anon_sym_RBRACE),
	4072:  uint16(267),
	4073:  uint16(9),
	4074:  uint16(anon_sym_using),
	4075:  uint16(anon_sym_DOLLAR),
	4076:  uint16(anon_sym_annotation),
	4077:  uint16(anon_sym_const),
	4078:  uint16(anon_sym_enum),
	4079:  uint16(anon_sym_interface),
	4080:  uint16(anon_sym_struct),
	4081:  uint16(anon_sym_union),
	4082:  uint16(sym_identifier),
	4083:  uint16(3),
	4084:  uint16(3),
	4085:  uint16(1),
	4086:  uint16(sym_comment),
	4087:  uint16(269),
	4088:  uint16(4),
	4090:  uint16(sym_unique_id),
	4091:  uint16(anon_sym_DOLLARimport),
	4092:  uint16(anon_sym_RBRACE),
	4093:  uint16(271),
	4094:  uint16(9),
	4095:  uint16(anon_sym_using),
	4096:  uint16(anon_sym_DOLLAR),
	4097:  uint16(anon_sym_annotation),
	4098:  uint16(anon_sym_const),
	4099:  uint16(anon_sym_enum),
	4100:  uint16(anon_sym_interface),
	4101:  uint16(anon_sym_struct),
	4102:  uint16(anon_sym_union),
	4103:  uint16(sym_identifier),
	4104:  uint16(3),
	4105:  uint16(3),
	4106:  uint16(1),
	4107:  uint16(sym_comment),
	4108:  uint16(273),
	4109:  uint16(4),
	4111:  uint16(sym_unique_id),
	4112:  uint16(anon_sym_DOLLARimport),
	4113:  uint16(anon_sym_RBRACE),
	4114:  uint16(275),
	4115:  uint16(9),
	4116:  uint16(anon_sym_using),
	4117:  uint16(anon_sym_DOLLAR),
	4118:  uint16(anon_sym_annotation),
	4119:  uint16(anon_sym_const),
	4120:  uint16(anon_sym_enum),
	4121:  uint16(anon_sym_interface),
	4122:  uint16(anon_sym_struct),
	4123:  uint16(anon_sym_union),
	4124:  uint16(sym_identifier),
	4125:  uint16(3),
	4126:  uint16(3),
	4127:  uint16(1),
	4128:  uint16(sym_comment),
	4129:  uint16(277),
	4130:  uint16(4),
	4132:  uint16(sym_unique_id),
	4133:  uint16(anon_sym_DOLLARimport),
	4134:  uint16(anon_sym_RBRACE),
	4135:  uint16(279),
	4136:  uint16(9),
	4137:  uint16(anon_sym_using),
	4138:  uint16(anon_sym_DOLLAR),
	4139:  uint16(anon_sym_annotation),
	4140:  uint16(anon_sym_const),
	4141:  uint16(anon_sym_enum),
	4142:  uint16(anon_sym_interface),
	4143:  uint16(anon_sym_struct),
	4144:  uint16(anon_sym_union),
	4145:  uint16(sym_identifier),
	4146:  uint16(3),
	4147:  uint16(3),
	4148:  uint16(1),
	4149:  uint16(sym_comment),
	4150:  uint16(281),
	4151:  uint16(4),
	4153:  uint16(sym_unique_id),
	4154:  uint16(anon_sym_DOLLARimport),
	4155:  uint16(anon_sym_RBRACE),
	4156:  uint16(283),
	4157:  uint16(9),
	4158:  uint16(anon_sym_using),
	4159:  uint16(anon_sym_DOLLAR),
	4160:  uint16(anon_sym_annotation),
	4161:  uint16(anon_sym_const),
	4162:  uint16(anon_sym_enum),
	4163:  uint16(anon_sym_interface),
	4164:  uint16(anon_sym_struct),
	4165:  uint16(anon_sym_union),
	4166:  uint16(sym_identifier),
	4167:  uint16(3),
	4168:  uint16(3),
	4169:  uint16(1),
	4170:  uint16(sym_comment),
	4171:  uint16(285),
	4172:  uint16(4),
	4174:  uint16(sym_unique_id),
	4175:  uint16(anon_sym_DOLLARimport),
	4176:  uint16(anon_sym_RBRACE),
	4177:  uint16(287),
	4178:  uint16(9),
	4179:  uint16(anon_sym_using),
	4180:  uint16(anon_sym_DOLLAR),
	4181:  uint16(anon_sym_annotation),
	4182:  uint16(anon_sym_const),
	4183:  uint16(anon_sym_enum),
	4184:  uint16(anon_sym_interface),
	4185:  uint16(anon_sym_struct),
	4186:  uint16(anon_sym_union),
	4187:  uint16(sym_identifier),
	4188:  uint16(3),
	4189:  uint16(3),
	4190:  uint16(1),
	4191:  uint16(sym_comment),
	4192:  uint16(289),
	4193:  uint16(4),
	4195:  uint16(sym_unique_id),
	4196:  uint16(anon_sym_DOLLARimport),
	4197:  uint16(anon_sym_RBRACE),
	4198:  uint16(291),
	4199:  uint16(9),
	4200:  uint16(anon_sym_using),
	4201:  uint16(anon_sym_DOLLAR),
	4202:  uint16(anon_sym_annotation),
	4203:  uint16(anon_sym_const),
	4204:  uint16(anon_sym_enum),
	4205:  uint16(anon_sym_interface),
	4206:  uint16(anon_sym_struct),
	4207:  uint16(anon_sym_union),
	4208:  uint16(sym_identifier),
	4209:  uint16(3),
	4210:  uint16(3),
	4211:  uint16(1),
	4212:  uint16(sym_comment),
	4213:  uint16(293),
	4214:  uint16(4),
	4216:  uint16(sym_unique_id),
	4217:  uint16(anon_sym_DOLLARimport),
	4218:  uint16(anon_sym_RBRACE),
	4219:  uint16(295),
	4220:  uint16(9),
	4221:  uint16(anon_sym_using),
	4222:  uint16(anon_sym_DOLLAR),
	4223:  uint16(anon_sym_annotation),
	4224:  uint16(anon_sym_const),
	4225:  uint16(anon_sym_enum),
	4226:  uint16(anon_sym_interface),
	4227:  uint16(anon_sym_struct),
	4228:  uint16(anon_sym_union),
	4229:  uint16(sym_identifier),
	4230:  uint16(3),
	4231:  uint16(3),
	4232:  uint16(1),
	4233:  uint16(sym_comment),
	4234:  uint16(297),
	4235:  uint16(4),
	4237:  uint16(sym_unique_id),
	4238:  uint16(anon_sym_DOLLARimport),
	4239:  uint16(anon_sym_RBRACE),
	4240:  uint16(299),
	4241:  uint16(9),
	4242:  uint16(anon_sym_using),
	4243:  uint16(anon_sym_DOLLAR),
	4244:  uint16(anon_sym_annotation),
	4245:  uint16(anon_sym_const),
	4246:  uint16(anon_sym_enum),
	4247:  uint16(anon_sym_interface),
	4248:  uint16(anon_sym_struct),
	4249:  uint16(anon_sym_union),
	4250:  uint16(sym_identifier),
	4251:  uint16(3),
	4252:  uint16(3),
	4253:  uint16(1),
	4254:  uint16(sym_comment),
	4255:  uint16(301),
	4256:  uint16(4),
	4258:  uint16(sym_unique_id),
	4259:  uint16(anon_sym_DOLLARimport),
	4260:  uint16(anon_sym_RBRACE),
	4261:  uint16(303),
	4262:  uint16(9),
	4263:  uint16(anon_sym_using),
	4264:  uint16(anon_sym_DOLLAR),
	4265:  uint16(anon_sym_annotation),
	4266:  uint16(anon_sym_const),
	4267:  uint16(anon_sym_enum),
	4268:  uint16(anon_sym_interface),
	4269:  uint16(anon_sym_struct),
	4270:  uint16(anon_sym_union),
	4271:  uint16(sym_identifier),
	4272:  uint16(3),
	4273:  uint16(3),
	4274:  uint16(1),
	4275:  uint16(sym_comment),
	4276:  uint16(305),
	4277:  uint16(4),
	4279:  uint16(sym_unique_id),
	4280:  uint16(anon_sym_DOLLARimport),
	4281:  uint16(anon_sym_RBRACE),
	4282:  uint16(307),
	4283:  uint16(9),
	4284:  uint16(anon_sym_using),
	4285:  uint16(anon_sym_DOLLAR),
	4286:  uint16(anon_sym_annotation),
	4287:  uint16(anon_sym_const),
	4288:  uint16(anon_sym_enum),
	4289:  uint16(anon_sym_interface),
	4290:  uint16(anon_sym_struct),
	4291:  uint16(anon_sym_union),
	4292:  uint16(sym_identifier),
	4293:  uint16(3),
	4294:  uint16(3),
	4295:  uint16(1),
	4296:  uint16(sym_comment),
	4297:  uint16(309),
	4298:  uint16(4),
	4300:  uint16(sym_unique_id),
	4301:  uint16(anon_sym_DOLLARimport),
	4302:  uint16(anon_sym_RBRACE),
	4303:  uint16(311),
	4304:  uint16(9),
	4305:  uint16(anon_sym_using),
	4306:  uint16(anon_sym_DOLLAR),
	4307:  uint16(anon_sym_annotation),
	4308:  uint16(anon_sym_const),
	4309:  uint16(anon_sym_enum),
	4310:  uint16(anon_sym_interface),
	4311:  uint16(anon_sym_struct),
	4312:  uint16(anon_sym_union),
	4313:  uint16(sym_identifier),
	4314:  uint16(3),
	4315:  uint16(3),
	4316:  uint16(1),
	4317:  uint16(sym_comment),
	4318:  uint16(313),
	4319:  uint16(4),
	4321:  uint16(sym_unique_id),
	4322:  uint16(anon_sym_DOLLARimport),
	4323:  uint16(anon_sym_RBRACE),
	4324:  uint16(315),
	4325:  uint16(9),
	4326:  uint16(anon_sym_using),
	4327:  uint16(anon_sym_DOLLAR),
	4328:  uint16(anon_sym_annotation),
	4329:  uint16(anon_sym_const),
	4330:  uint16(anon_sym_enum),
	4331:  uint16(anon_sym_interface),
	4332:  uint16(anon_sym_struct),
	4333:  uint16(anon_sym_union),
	4334:  uint16(sym_identifier),
	4335:  uint16(3),
	4336:  uint16(3),
	4337:  uint16(1),
	4338:  uint16(sym_comment),
	4339:  uint16(317),
	4340:  uint16(4),
	4342:  uint16(sym_unique_id),
	4343:  uint16(anon_sym_DOLLARimport),
	4344:  uint16(anon_sym_RBRACE),
	4345:  uint16(319),
	4346:  uint16(9),
	4347:  uint16(anon_sym_using),
	4348:  uint16(anon_sym_DOLLAR),
	4349:  uint16(anon_sym_annotation),
	4350:  uint16(anon_sym_const),
	4351:  uint16(anon_sym_enum),
	4352:  uint16(anon_sym_interface),
	4353:  uint16(anon_sym_struct),
	4354:  uint16(anon_sym_union),
	4355:  uint16(sym_identifier),
	4356:  uint16(3),
	4357:  uint16(3),
	4358:  uint16(1),
	4359:  uint16(sym_comment),
	4360:  uint16(321),
	4361:  uint16(4),
	4363:  uint16(sym_unique_id),
	4364:  uint16(anon_sym_DOLLARimport),
	4365:  uint16(anon_sym_RBRACE),
	4366:  uint16(323),
	4367:  uint16(9),
	4368:  uint16(anon_sym_using),
	4369:  uint16(anon_sym_DOLLAR),
	4370:  uint16(anon_sym_annotation),
	4371:  uint16(anon_sym_const),
	4372:  uint16(anon_sym_enum),
	4373:  uint16(anon_sym_interface),
	4374:  uint16(anon_sym_struct),
	4375:  uint16(anon_sym_union),
	4376:  uint16(sym_identifier),
	4377:  uint16(3),
	4378:  uint16(3),
	4379:  uint16(1),
	4380:  uint16(sym_comment),
	4381:  uint16(325),
	4382:  uint16(4),
	4384:  uint16(sym_unique_id),
	4385:  uint16(anon_sym_DOLLARimport),
	4386:  uint16(anon_sym_RBRACE),
	4387:  uint16(327),
	4388:  uint16(9),
	4389:  uint16(anon_sym_using),
	4390:  uint16(anon_sym_DOLLAR),
	4391:  uint16(anon_sym_annotation),
	4392:  uint16(anon_sym_const),
	4393:  uint16(anon_sym_enum),
	4394:  uint16(anon_sym_interface),
	4395:  uint16(anon_sym_struct),
	4396:  uint16(anon_sym_union),
	4397:  uint16(sym_identifier),
	4398:  uint16(3),
	4399:  uint16(3),
	4400:  uint16(1),
	4401:  uint16(sym_comment),
	4402:  uint16(329),
	4403:  uint16(4),
	4405:  uint16(sym_unique_id),
	4406:  uint16(anon_sym_DOLLARimport),
	4407:  uint16(anon_sym_RBRACE),
	4408:  uint16(331),
	4409:  uint16(9),
	4410:  uint16(anon_sym_using),
	4411:  uint16(anon_sym_DOLLAR),
	4412:  uint16(anon_sym_annotation),
	4413:  uint16(anon_sym_const),
	4414:  uint16(anon_sym_enum),
	4415:  uint16(anon_sym_interface),
	4416:  uint16(anon_sym_struct),
	4417:  uint16(anon_sym_union),
	4418:  uint16(sym_identifier),
	4419:  uint16(3),
	4420:  uint16(3),
	4421:  uint16(1),
	4422:  uint16(sym_comment),
	4423:  uint16(333),
	4424:  uint16(4),
	4426:  uint16(sym_unique_id),
	4427:  uint16(anon_sym_DOLLARimport),
	4428:  uint16(anon_sym_RBRACE),
	4429:  uint16(335),
	4430:  uint16(9),
	4431:  uint16(anon_sym_using),
	4432:  uint16(anon_sym_DOLLAR),
	4433:  uint16(anon_sym_annotation),
	4434:  uint16(anon_sym_const),
	4435:  uint16(anon_sym_enum),
	4436:  uint16(anon_sym_interface),
	4437:  uint16(anon_sym_struct),
	4438:  uint16(anon_sym_union),
	4439:  uint16(sym_identifier),
	4440:  uint16(3),
	4441:  uint16(3),
	4442:  uint16(1),
	4443:  uint16(sym_comment),
	4444:  uint16(337),
	4445:  uint16(4),
	4447:  uint16(sym_unique_id),
	4448:  uint16(anon_sym_DOLLARimport),
	4449:  uint16(anon_sym_RBRACE),
	4450:  uint16(339),
	4451:  uint16(9),
	4452:  uint16(anon_sym_using),
	4453:  uint16(anon_sym_DOLLAR),
	4454:  uint16(anon_sym_annotation),
	4455:  uint16(anon_sym_const),
	4456:  uint16(anon_sym_enum),
	4457:  uint16(anon_sym_interface),
	4458:  uint16(anon_sym_struct),
	4459:  uint16(anon_sym_union),
	4460:  uint16(sym_identifier),
	4461:  uint16(3),
	4462:  uint16(3),
	4463:  uint16(1),
	4464:  uint16(sym_comment),
	4465:  uint16(341),
	4466:  uint16(4),
	4468:  uint16(sym_unique_id),
	4469:  uint16(anon_sym_DOLLARimport),
	4470:  uint16(anon_sym_RBRACE),
	4471:  uint16(343),
	4472:  uint16(9),
	4473:  uint16(anon_sym_using),
	4474:  uint16(anon_sym_DOLLAR),
	4475:  uint16(anon_sym_annotation),
	4476:  uint16(anon_sym_const),
	4477:  uint16(anon_sym_enum),
	4478:  uint16(anon_sym_interface),
	4479:  uint16(anon_sym_struct),
	4480:  uint16(anon_sym_union),
	4481:  uint16(sym_identifier),
	4482:  uint16(3),
	4483:  uint16(3),
	4484:  uint16(1),
	4485:  uint16(sym_comment),
	4486:  uint16(345),
	4487:  uint16(4),
	4489:  uint16(sym_unique_id),
	4490:  uint16(anon_sym_DOLLARimport),
	4491:  uint16(anon_sym_RBRACE),
	4492:  uint16(347),
	4493:  uint16(9),
	4494:  uint16(anon_sym_using),
	4495:  uint16(anon_sym_DOLLAR),
	4496:  uint16(anon_sym_annotation),
	4497:  uint16(anon_sym_const),
	4498:  uint16(anon_sym_enum),
	4499:  uint16(anon_sym_interface),
	4500:  uint16(anon_sym_struct),
	4501:  uint16(anon_sym_union),
	4502:  uint16(sym_identifier),
	4503:  uint16(3),
	4504:  uint16(3),
	4505:  uint16(1),
	4506:  uint16(sym_comment),
	4507:  uint16(349),
	4508:  uint16(4),
	4510:  uint16(sym_unique_id),
	4511:  uint16(anon_sym_DOLLARimport),
	4512:  uint16(anon_sym_RBRACE),
	4513:  uint16(351),
	4514:  uint16(9),
	4515:  uint16(anon_sym_using),
	4516:  uint16(anon_sym_DOLLAR),
	4517:  uint16(anon_sym_annotation),
	4518:  uint16(anon_sym_const),
	4519:  uint16(anon_sym_enum),
	4520:  uint16(anon_sym_interface),
	4521:  uint16(anon_sym_struct),
	4522:  uint16(anon_sym_union),
	4523:  uint16(sym_identifier),
	4524:  uint16(3),
	4525:  uint16(3),
	4526:  uint16(1),
	4527:  uint16(sym_comment),
	4528:  uint16(353),
	4529:  uint16(4),
	4531:  uint16(sym_unique_id),
	4532:  uint16(anon_sym_DOLLARimport),
	4533:  uint16(anon_sym_RBRACE),
	4534:  uint16(355),
	4535:  uint16(9),
	4536:  uint16(anon_sym_using),
	4537:  uint16(anon_sym_DOLLAR),
	4538:  uint16(anon_sym_annotation),
	4539:  uint16(anon_sym_const),
	4540:  uint16(anon_sym_enum),
	4541:  uint16(anon_sym_interface),
	4542:  uint16(anon_sym_struct),
	4543:  uint16(anon_sym_union),
	4544:  uint16(sym_identifier),
	4545:  uint16(3),
	4546:  uint16(3),
	4547:  uint16(1),
	4548:  uint16(sym_comment),
	4549:  uint16(357),
	4550:  uint16(4),
	4552:  uint16(sym_unique_id),
	4553:  uint16(anon_sym_DOLLARimport),
	4554:  uint16(anon_sym_RBRACE),
	4555:  uint16(359),
	4556:  uint16(9),
	4557:  uint16(anon_sym_using),
	4558:  uint16(anon_sym_DOLLAR),
	4559:  uint16(anon_sym_annotation),
	4560:  uint16(anon_sym_const),
	4561:  uint16(anon_sym_enum),
	4562:  uint16(anon_sym_interface),
	4563:  uint16(anon_sym_struct),
	4564:  uint16(anon_sym_union),
	4565:  uint16(sym_identifier),
	4566:  uint16(3),
	4567:  uint16(3),
	4568:  uint16(1),
	4569:  uint16(sym_comment),
	4570:  uint16(361),
	4571:  uint16(4),
	4573:  uint16(sym_unique_id),
	4574:  uint16(anon_sym_DOLLARimport),
	4575:  uint16(anon_sym_RBRACE),
	4576:  uint16(363),
	4577:  uint16(9),
	4578:  uint16(anon_sym_using),
	4579:  uint16(anon_sym_DOLLAR),
	4580:  uint16(anon_sym_annotation),
	4581:  uint16(anon_sym_const),
	4582:  uint16(anon_sym_enum),
	4583:  uint16(anon_sym_interface),
	4584:  uint16(anon_sym_struct),
	4585:  uint16(anon_sym_union),
	4586:  uint16(sym_identifier),
	4587:  uint16(5),
	4588:  uint16(367),
	4589:  uint16(1),
	4590:  uint16(sym_unescaped_block_string_fragment),
	4591:  uint16(371),
	4592:  uint16(1),
	4593:  uint16(sym_comment),
	4594:  uint16(369),
	4595:  uint16(2),
	4596:  uint16(aux_sym__escape_sequence_token1),
	4597:  uint16(sym_escape_sequence),
	4598:  uint16(109),
	4599:  uint16(2),
	4600:  uint16(sym__escape_sequence),
	4601:  uint16(aux_sym_block_text_repeat1),
	4602:  uint16(365),
	4603:  uint16(7),
	4604:  uint16(anon_sym_SEMI),
	4605:  uint16(anon_sym_RPAREN),
	4606:  uint16(anon_sym_DOLLAR),
	4607:  uint16(anon_sym_COMMA),
	4608:  uint16(anon_sym_RBRACK),
	4609:  uint16(anon_sym_BQUOTE),
	4610:  uint16(sym_identifier),
	4611:  uint16(5),
	4612:  uint16(371),
	4613:  uint16(1),
	4614:  uint16(sym_comment),
	4615:  uint16(375),
	4616:  uint16(1),
	4617:  uint16(sym_unescaped_block_string_fragment),
	4618:  uint16(378),
	4619:  uint16(2),
	4620:  uint16(aux_sym__escape_sequence_token1),
	4621:  uint16(sym_escape_sequence),
	4622:  uint16(108),
	4623:  uint16(2),
	4624:  uint16(sym__escape_sequence),
	4625:  uint16(aux_sym_block_text_repeat1),
	4626:  uint16(373),
	4627:  uint16(7),
	4628:  uint16(anon_sym_SEMI),
	4629:  uint16(anon_sym_RPAREN),
	4630:  uint16(anon_sym_DOLLAR),
	4631:  uint16(anon_sym_COMMA),
	4632:  uint16(anon_sym_RBRACK),
	4633:  uint16(anon_sym_BQUOTE),
	4634:  uint16(sym_identifier),
	4635:  uint16(5),
	4636:  uint16(371),
	4637:  uint16(1),
	4638:  uint16(sym_comment),
	4639:  uint16(383),
	4640:  uint16(1),
	4641:  uint16(sym_unescaped_block_string_fragment),
	4642:  uint16(385),
	4643:  uint16(2),
	4644:  uint16(aux_sym__escape_sequence_token1),
	4645:  uint16(sym_escape_sequence),
	4646:  uint16(108),
	4647:  uint16(2),
	4648:  uint16(sym__escape_sequence),
	4649:  uint16(aux_sym_block_text_repeat1),
	4650:  uint16(381),
	4651:  uint16(7),
	4652:  uint16(anon_sym_SEMI),
	4653:  uint16(anon_sym_RPAREN),
	4654:  uint16(anon_sym_DOLLAR),
	4655:  uint16(anon_sym_COMMA),
	4656:  uint16(anon_sym_RBRACK),
	4657:  uint16(anon_sym_BQUOTE),
	4658:  uint16(sym_identifier),
	4659:  uint16(2),
	4660:  uint16(3),
	4661:  uint16(1),
	4662:  uint16(sym_comment),
	4663:  uint16(387),
	4664:  uint16(12),
	4665:  uint16(sym_unique_id),
	4666:  uint16(anon_sym_SEMI),
	4667:  uint16(anon_sym_EQ),
	4668:  uint16(anon_sym_DOT),
	4669:  uint16(anon_sym_LPAREN),
	4670:  uint16(anon_sym_RPAREN),
	4671:  uint16(anon_sym_DOLLAR),
	4672:  uint16(anon_sym_COMMA),
	4673:  uint16(anon_sym_RBRACK),
	4674:  uint16(anon_sym_LBRACE),
	4675:  uint16(anon_sym_extends),
	4676:  uint16(anon_sym_DASH_GT),
	4677:  uint16(8),
	4678:  uint16(3),
	4679:  uint16(1),
	4680:  uint16(sym_comment),
	4681:  uint16(107),
	4682:  uint16(1),
	4683:  uint16(anon_sym_enum),
	4684:  uint16(109),
	4685:  uint16(1),
	4686:  uint16(anon_sym_interface),
	4687:  uint16(111),
	4688:  uint16(1),
	4689:  uint16(anon_sym_struct),
	4690:  uint16(389),
	4691:  uint16(1),
	4692:  uint16(sym_identifier),
	4693:  uint16(391),
	4694:  uint16(1),
	4695:  uint16(anon_sym_RBRACE),
	4696:  uint16(298),
	4697:  uint16(1),
	4698:  uint16(sym__method_identifier),
	4699:  uint16(133),
	4700:  uint16(5),
	4701:  uint16(sym_struct),
	4702:  uint16(sym_enum),
	4703:  uint16(sym_interface),
	4704:  uint16(sym_method),
	4705:  uint16(aux_sym_interface_repeat1),
	4706:  uint16(8),
	4707:  uint16(3),
	4708:  uint16(1),
	4709:  uint16(sym_comment),
	4710:  uint16(107),
	4711:  uint16(1),
	4712:  uint16(anon_sym_enum),
	4713:  uint16(109),
	4714:  uint16(1),
	4715:  uint16(anon_sym_interface),
	4716:  uint16(111),
	4717:  uint16(1),
	4718:  uint16(anon_sym_struct),
	4719:  uint16(389),
	4720:  uint16(1),
	4721:  uint16(sym_identifier),
	4722:  uint16(393),
	4723:  uint16(1),
	4724:  uint16(anon_sym_RBRACE),
	4725:  uint16(298),
	4726:  uint16(1),
	4727:  uint16(sym__method_identifier),
	4728:  uint16(130),
	4729:  uint16(5),
	4730:  uint16(sym_struct),
	4731:  uint16(sym_enum),
	4732:  uint16(sym_interface),
	4733:  uint16(sym_method),
	4734:  uint16(aux_sym_interface_repeat1),
	4735:  uint16(8),
	4736:  uint16(3),
	4737:  uint16(1),
	4738:  uint16(sym_comment),
	4739:  uint16(107),
	4740:  uint16(1),
	4741:  uint16(anon_sym_enum),
	4742:  uint16(109),
	4743:  uint16(1),
	4744:  uint16(anon_sym_interface),
	4745:  uint16(111),
	4746:  uint16(1),
	4747:  uint16(anon_sym_struct),
	4748:  uint16(389),
	4749:  uint16(1),
	4750:  uint16(sym_identifier),
	4751:  uint16(395),
	4752:  uint16(1),
	4753:  uint16(anon_sym_RBRACE),
	4754:  uint16(298),
	4755:  uint16(1),
	4756:  uint16(sym__method_identifier),
	4757:  uint16(136),
	4758:  uint16(5),
	4759:  uint16(sym_struct),
	4760:  uint16(sym_enum),
	4761:  uint16(sym_interface),
	4762:  uint16(sym_method),
	4763:  uint16(aux_sym_interface_repeat1),
	4764:  uint16(8),
	4765:  uint16(3),
	4766:  uint16(1),
	4767:  uint16(sym_comment),
	4768:  uint16(107),
	4769:  uint16(1),
	4770:  uint16(anon_sym_enum),
	4771:  uint16(109),
	4772:  uint16(1),
	4773:  uint16(anon_sym_interface),
	4774:  uint16(111),
	4775:  uint16(1),
	4776:  uint16(anon_sym_struct),
	4777:  uint16(389),
	4778:  uint16(1),
	4779:  uint16(sym_identifier),
	4780:  uint16(397),
	4781:  uint16(1),
	4782:  uint16(anon_sym_RBRACE),
	4783:  uint16(298),
	4784:  uint16(1),
	4785:  uint16(sym__method_identifier),
	4786:  uint16(121),
	4787:  uint16(5),
	4788:  uint16(sym_struct),
	4789:  uint16(sym_enum),
	4790:  uint16(sym_interface),
	4791:  uint16(sym_method),
	4792:  uint16(aux_sym_interface_repeat1),
	4793:  uint16(8),
	4794:  uint16(3),
	4795:  uint16(1),
	4796:  uint16(sym_comment),
	4797:  uint16(107),
	4798:  uint16(1),
	4799:  uint16(anon_sym_enum),
	4800:  uint16(109),
	4801:  uint16(1),
	4802:  uint16(anon_sym_interface),
	4803:  uint16(111),
	4804:  uint16(1),
	4805:  uint16(anon_sym_struct),
	4806:  uint16(389),
	4807:  uint16(1),
	4808:  uint16(sym_identifier),
	4809:  uint16(399),
	4810:  uint16(1),
	4811:  uint16(anon_sym_RBRACE),
	4812:  uint16(298),
	4813:  uint16(1),
	4814:  uint16(sym__method_identifier),
	4815:  uint16(136),
	4816:  uint16(5),
	4817:  uint16(sym_struct),
	4818:  uint16(sym_enum),
	4819:  uint16(sym_interface),
	4820:  uint16(sym_method),
	4821:  uint16(aux_sym_interface_repeat1),
	4822:  uint16(7),
	4823:  uint16(3),
	4824:  uint16(1),
	4825:  uint16(sym_comment),
	4826:  uint16(403),
	4827:  uint16(1),
	4828:  uint16(anon_sym_DOT),
	4829:  uint16(405),
	4830:  uint16(1),
	4831:  uint16(anon_sym_LPAREN),
	4832:  uint16(152),
	4833:  uint16(1),
	4834:  uint16(sym_generics),
	4835:  uint16(155),
	4836:  uint16(1),
	4837:  uint16(aux_sym__annotation_call_repeat1),
	4838:  uint16(224),
	4839:  uint16(1),
	4840:  uint16(sym_annotation_literal),
	4841:  uint16(401),
	4842:  uint16(6),
	4843:  uint16(anon_sym_SEMI),
	4844:  uint16(anon_sym_EQ),
	4845:  uint16(anon_sym_RPAREN),
	4846:  uint16(anon_sym_DOLLAR),
	4847:  uint16(anon_sym_COMMA),
	4848:  uint16(anon_sym_LBRACE),
	4849:  uint16(8),
	4850:  uint16(3),
	4851:  uint16(1),
	4852:  uint16(sym_comment),
	4853:  uint16(107),
	4854:  uint16(1),
	4855:  uint16(anon_sym_enum),
	4856:  uint16(109),
	4857:  uint16(1),
	4858:  uint16(anon_sym_interface),
	4859:  uint16(111),
	4860:  uint16(1),
	4861:  uint16(anon_sym_struct),
	4862:  uint16(389),
	4863:  uint16(1),
	4864:  uint16(sym_identifier),
	4865:  uint16(407),
	4866:  uint16(1),
	4867:  uint16(anon_sym_RBRACE),
	4868:  uint16(298),
	4869:  uint16(1),
	4870:  uint16(sym__method_identifier),
	4871:  uint16(136),
	4872:  uint16(5),
	4873:  uint16(sym_struct),
	4874:  uint16(sym_enum),
	4875:  uint16(sym_interface),
	4876:  uint16(sym_method),
	4877:  uint16(aux_sym_interface_repeat1),
	4878:  uint16(8),
	4879:  uint16(3),
	4880:  uint16(1),
	4881:  uint16(sym_comment),
	4882:  uint16(107),
	4883:  uint16(1),
	4884:  uint16(anon_sym_enum),
	4885:  uint16(109),
	4886:  uint16(1),
	4887:  uint16(anon_sym_interface),
	4888:  uint16(111),
	4889:  uint16(1),
	4890:  uint16(anon_sym_struct),
	4891:  uint16(389),
	4892:  uint16(1),
	4893:  uint16(sym_identifier),
	4894:  uint16(407),
	4895:  uint16(1),
	4896:  uint16(anon_sym_RBRACE),
	4897:  uint16(298),
	4898:  uint16(1),
	4899:  uint16(sym__method_identifier),
	4900:  uint16(119),
	4901:  uint16(5),
	4902:  uint16(sym_struct),
	4903:  uint16(sym_enum),
	4904:  uint16(sym_interface),
	4905:  uint16(sym_method),
	4906:  uint16(aux_sym_interface_repeat1),
	4907:  uint16(8),
	4908:  uint16(3),
	4909:  uint16(1),
	4910:  uint16(sym_comment),
	4911:  uint16(107),
	4912:  uint16(1),
	4913:  uint16(anon_sym_enum),
	4914:  uint16(109),
	4915:  uint16(1),
	4916:  uint16(anon_sym_interface),
	4917:  uint16(111),
	4918:  uint16(1),
	4919:  uint16(anon_sym_struct),
	4920:  uint16(389),
	4921:  uint16(1),
	4922:  uint16(sym_identifier),
	4923:  uint16(409),
	4924:  uint16(1),
	4925:  uint16(anon_sym_RBRACE),
	4926:  uint16(298),
	4927:  uint16(1),
	4928:  uint16(sym__method_identifier),
	4929:  uint16(136),
	4930:  uint16(5),
	4931:  uint16(sym_struct),
	4932:  uint16(sym_enum),
	4933:  uint16(sym_interface),
	4934:  uint16(sym_method),
	4935:  uint16(aux_sym_interface_repeat1),
	4936:  uint16(8),
	4937:  uint16(3),
	4938:  uint16(1),
	4939:  uint16(sym_comment),
	4940:  uint16(107),
	4941:  uint16(1),
	4942:  uint16(anon_sym_enum),
	4943:  uint16(109),
	4944:  uint16(1),
	4945:  uint16(anon_sym_interface),
	4946:  uint16(111),
	4947:  uint16(1),
	4948:  uint16(anon_sym_struct),
	4949:  uint16(389),
	4950:  uint16(1),
	4951:  uint16(sym_identifier),
	4952:  uint16(411),
	4953:  uint16(1),
	4954:  uint16(anon_sym_RBRACE),
	4955:  uint16(298),
	4956:  uint16(1),
	4957:  uint16(sym__method_identifier),
	4958:  uint16(137),
	4959:  uint16(5),
	4960:  uint16(sym_struct),
	4961:  uint16(sym_enum),
	4962:  uint16(sym_interface),
	4963:  uint16(sym_method),
	4964:  uint16(aux_sym_interface_repeat1),
	4965:  uint16(8),
	4966:  uint16(3),
	4967:  uint16(1),
	4968:  uint16(sym_comment),
	4969:  uint16(107),
	4970:  uint16(1),
	4971:  uint16(anon_sym_enum),
	4972:  uint16(109),
	4973:  uint16(1),
	4974:  uint16(anon_sym_interface),
	4975:  uint16(111),
	4976:  uint16(1),
	4977:  uint16(anon_sym_struct),
	4978:  uint16(389),
	4979:  uint16(1),
	4980:  uint16(sym_identifier),
	4981:  uint16(411),
	4982:  uint16(1),
	4983:  uint16(anon_sym_RBRACE),
	4984:  uint16(298),
	4985:  uint16(1),
	4986:  uint16(sym__method_identifier),
	4987:  uint16(136),
	4988:  uint16(5),
	4989:  uint16(sym_struct),
	4990:  uint16(sym_enum),
	4991:  uint16(sym_interface),
	4992:  uint16(sym_method),
	4993:  uint16(aux_sym_interface_repeat1),
	4994:  uint16(8),
	4995:  uint16(3),
	4996:  uint16(1),
	4997:  uint16(sym_comment),
	4998:  uint16(107),
	4999:  uint16(1),
	5000:  uint16(anon_sym_enum),
	5001:  uint16(109),
	5002:  uint16(1),
	5003:  uint16(anon_sym_interface),
	5004:  uint16(111),
	5005:  uint16(1),
	5006:  uint16(anon_sym_struct),
	5007:  uint16(389),
	5008:  uint16(1),
	5009:  uint16(sym_identifier),
	5010:  uint16(399),
	5011:  uint16(1),
	5012:  uint16(anon_sym_RBRACE),
	5013:  uint16(298),
	5014:  uint16(1),
	5015:  uint16(sym__method_identifier),
	5016:  uint16(135),
	5017:  uint16(5),
	5018:  uint16(sym_struct),
	5019:  uint16(sym_enum),
	5020:  uint16(sym_interface),
	5021:  uint16(sym_method),
	5022:  uint16(aux_sym_interface_repeat1),
	5023:  uint16(8),
	5024:  uint16(3),
	5025:  uint16(1),
	5026:  uint16(sym_comment),
	5027:  uint16(107),
	5028:  uint16(1),
	5029:  uint16(anon_sym_enum),
	5030:  uint16(109),
	5031:  uint16(1),
	5032:  uint16(anon_sym_interface),
	5033:  uint16(111),
	5034:  uint16(1),
	5035:  uint16(anon_sym_struct),
	5036:  uint16(389),
	5037:  uint16(1),
	5038:  uint16(sym_identifier),
	5039:  uint16(391),
	5040:  uint16(1),
	5041:  uint16(anon_sym_RBRACE),
	5042:  uint16(298),
	5043:  uint16(1),
	5044:  uint16(sym__method_identifier),
	5045:  uint16(136),
	5046:  uint16(5),
	5047:  uint16(sym_struct),
	5048:  uint16(sym_enum),
	5049:  uint16(sym_interface),
	5050:  uint16(sym_method),
	5051:  uint16(aux_sym_interface_repeat1),
	5052:  uint16(8),
	5053:  uint16(3),
	5054:  uint16(1),
	5055:  uint16(sym_comment),
	5056:  uint16(107),
	5057:  uint16(1),
	5058:  uint16(anon_sym_enum),
	5059:  uint16(109),
	5060:  uint16(1),
	5061:  uint16(anon_sym_interface),
	5062:  uint16(111),
	5063:  uint16(1),
	5064:  uint16(anon_sym_struct),
	5065:  uint16(389),
	5066:  uint16(1),
	5067:  uint16(sym_identifier),
	5068:  uint16(413),
	5069:  uint16(1),
	5070:  uint16(anon_sym_RBRACE),
	5071:  uint16(298),
	5072:  uint16(1),
	5073:  uint16(sym__method_identifier),
	5074:  uint16(128),
	5075:  uint16(5),
	5076:  uint16(sym_struct),
	5077:  uint16(sym_enum),
	5078:  uint16(sym_interface),
	5079:  uint16(sym_method),
	5080:  uint16(aux_sym_interface_repeat1),
	5081:  uint16(8),
	5082:  uint16(3),
	5083:  uint16(1),
	5084:  uint16(sym_comment),
	5085:  uint16(107),
	5086:  uint16(1),
	5087:  uint16(anon_sym_enum),
	5088:  uint16(109),
	5089:  uint16(1),
	5090:  uint16(anon_sym_interface),
	5091:  uint16(111),
	5092:  uint16(1),
	5093:  uint16(anon_sym_struct),
	5094:  uint16(389),
	5095:  uint16(1),
	5096:  uint16(sym_identifier),
	5097:  uint16(415),
	5098:  uint16(1),
	5099:  uint16(anon_sym_RBRACE),
	5100:  uint16(298),
	5101:  uint16(1),
	5102:  uint16(sym__method_identifier),
	5103:  uint16(129),
	5104:  uint16(5),
	5105:  uint16(sym_struct),
	5106:  uint16(sym_enum),
	5107:  uint16(sym_interface),
	5108:  uint16(sym_method),
	5109:  uint16(aux_sym_interface_repeat1),
	5110:  uint16(8),
	5111:  uint16(3),
	5112:  uint16(1),
	5113:  uint16(sym_comment),
	5114:  uint16(107),
	5115:  uint16(1),
	5116:  uint16(anon_sym_enum),
	5117:  uint16(109),
	5118:  uint16(1),
	5119:  uint16(anon_sym_interface),
	5120:  uint16(111),
	5121:  uint16(1),
	5122:  uint16(anon_sym_struct),
	5123:  uint16(389),
	5124:  uint16(1),
	5125:  uint16(sym_identifier),
	5126:  uint16(415),
	5127:  uint16(1),
	5128:  uint16(anon_sym_RBRACE),
	5129:  uint16(298),
	5130:  uint16(1),
	5131:  uint16(sym__method_identifier),
	5132:  uint16(136),
	5133:  uint16(5),
	5134:  uint16(sym_struct),
	5135:  uint16(sym_enum),
	5136:  uint16(sym_interface),
	5137:  uint16(sym_method),
	5138:  uint16(aux_sym_interface_repeat1),
	5139:  uint16(8),
	5140:  uint16(3),
	5141:  uint16(1),
	5142:  uint16(sym_comment),
	5143:  uint16(107),
	5144:  uint16(1),
	5145:  uint16(anon_sym_enum),
	5146:  uint16(109),
	5147:  uint16(1),
	5148:  uint16(anon_sym_interface),
	5149:  uint16(111),
	5150:  uint16(1),
	5151:  uint16(anon_sym_struct),
	5152:  uint16(389),
	5153:  uint16(1),
	5154:  uint16(sym_identifier),
	5155:  uint16(417),
	5156:  uint16(1),
	5157:  uint16(anon_sym_RBRACE),
	5158:  uint16(298),
	5159:  uint16(1),
	5160:  uint16(sym__method_identifier),
	5161:  uint16(115),
	5162:  uint16(5),
	5163:  uint16(sym_struct),
	5164:  uint16(sym_enum),
	5165:  uint16(sym_interface),
	5166:  uint16(sym_method),
	5167:  uint16(aux_sym_interface_repeat1),
	5168:  uint16(8),
	5169:  uint16(3),
	5170:  uint16(1),
	5171:  uint16(sym_comment),
	5172:  uint16(107),
	5173:  uint16(1),
	5174:  uint16(anon_sym_enum),
	5175:  uint16(109),
	5176:  uint16(1),
	5177:  uint16(anon_sym_interface),
	5178:  uint16(111),
	5179:  uint16(1),
	5180:  uint16(anon_sym_struct),
	5181:  uint16(389),
	5182:  uint16(1),
	5183:  uint16(sym_identifier),
	5184:  uint16(419),
	5185:  uint16(1),
	5186:  uint16(anon_sym_RBRACE),
	5187:  uint16(298),
	5188:  uint16(1),
	5189:  uint16(sym__method_identifier),
	5190:  uint16(136),
	5191:  uint16(5),
	5192:  uint16(sym_struct),
	5193:  uint16(sym_enum),
	5194:  uint16(sym_interface),
	5195:  uint16(sym_method),
	5196:  uint16(aux_sym_interface_repeat1),
	5197:  uint16(8),
	5198:  uint16(3),
	5199:  uint16(1),
	5200:  uint16(sym_comment),
	5201:  uint16(107),
	5202:  uint16(1),
	5203:  uint16(anon_sym_enum),
	5204:  uint16(109),
	5205:  uint16(1),
	5206:  uint16(anon_sym_interface),
	5207:  uint16(111),
	5208:  uint16(1),
	5209:  uint16(anon_sym_struct),
	5210:  uint16(389),
	5211:  uint16(1),
	5212:  uint16(sym_identifier),
	5213:  uint16(421),
	5214:  uint16(1),
	5215:  uint16(anon_sym_RBRACE),
	5216:  uint16(298),
	5217:  uint16(1),
	5218:  uint16(sym__method_identifier),
	5219:  uint16(136),
	5220:  uint16(5),
	5221:  uint16(sym_struct),
	5222:  uint16(sym_enum),
	5223:  uint16(sym_interface),
	5224:  uint16(sym_method),
	5225:  uint16(aux_sym_interface_repeat1),
	5226:  uint16(8),
	5227:  uint16(3),
	5228:  uint16(1),
	5229:  uint16(sym_comment),
	5230:  uint16(107),
	5231:  uint16(1),
	5232:  uint16(anon_sym_enum),
	5233:  uint16(109),
	5234:  uint16(1),
	5235:  uint16(anon_sym_interface),
	5236:  uint16(111),
	5237:  uint16(1),
	5238:  uint16(anon_sym_struct),
	5239:  uint16(389),
	5240:  uint16(1),
	5241:  uint16(sym_identifier),
	5242:  uint16(423),
	5243:  uint16(1),
	5244:  uint16(anon_sym_RBRACE),
	5245:  uint16(298),
	5246:  uint16(1),
	5247:  uint16(sym__method_identifier),
	5248:  uint16(136),
	5249:  uint16(5),
	5250:  uint16(sym_struct),
	5251:  uint16(sym_enum),
	5252:  uint16(sym_interface),
	5253:  uint16(sym_method),
	5254:  uint16(aux_sym_interface_repeat1),
	5255:  uint16(8),
	5256:  uint16(3),
	5257:  uint16(1),
	5258:  uint16(sym_comment),
	5259:  uint16(107),
	5260:  uint16(1),
	5261:  uint16(anon_sym_enum),
	5262:  uint16(109),
	5263:  uint16(1),
	5264:  uint16(anon_sym_interface),
	5265:  uint16(111),
	5266:  uint16(1),
	5267:  uint16(anon_sym_struct),
	5268:  uint16(389),
	5269:  uint16(1),
	5270:  uint16(sym_identifier),
	5271:  uint16(423),
	5272:  uint16(1),
	5273:  uint16(anon_sym_RBRACE),
	5274:  uint16(298),
	5275:  uint16(1),
	5276:  uint16(sym__method_identifier),
	5277:  uint16(126),
	5278:  uint16(5),
	5279:  uint16(sym_struct),
	5280:  uint16(sym_enum),
	5281:  uint16(sym_interface),
	5282:  uint16(sym_method),
	5283:  uint16(aux_sym_interface_repeat1),
	5284:  uint16(8),
	5285:  uint16(3),
	5286:  uint16(1),
	5287:  uint16(sym_comment),
	5288:  uint16(107),
	5289:  uint16(1),
	5290:  uint16(anon_sym_enum),
	5291:  uint16(109),
	5292:  uint16(1),
	5293:  uint16(anon_sym_interface),
	5294:  uint16(111),
	5295:  uint16(1),
	5296:  uint16(anon_sym_struct),
	5297:  uint16(389),
	5298:  uint16(1),
	5299:  uint16(sym_identifier),
	5300:  uint16(425),
	5301:  uint16(1),
	5302:  uint16(anon_sym_RBRACE),
	5303:  uint16(298),
	5304:  uint16(1),
	5305:  uint16(sym__method_identifier),
	5306:  uint16(123),
	5307:  uint16(5),
	5308:  uint16(sym_struct),
	5309:  uint16(sym_enum),
	5310:  uint16(sym_interface),
	5311:  uint16(sym_method),
	5312:  uint16(aux_sym_interface_repeat1),
	5313:  uint16(8),
	5314:  uint16(3),
	5315:  uint16(1),
	5316:  uint16(sym_comment),
	5317:  uint16(107),
	5318:  uint16(1),
	5319:  uint16(anon_sym_enum),
	5320:  uint16(109),
	5321:  uint16(1),
	5322:  uint16(anon_sym_interface),
	5323:  uint16(111),
	5324:  uint16(1),
	5325:  uint16(anon_sym_struct),
	5326:  uint16(389),
	5327:  uint16(1),
	5328:  uint16(sym_identifier),
	5329:  uint16(427),
	5330:  uint16(1),
	5331:  uint16(anon_sym_RBRACE),
	5332:  uint16(298),
	5333:  uint16(1),
	5334:  uint16(sym__method_identifier),
	5335:  uint16(136),
	5336:  uint16(5),
	5337:  uint16(sym_struct),
	5338:  uint16(sym_enum),
	5339:  uint16(sym_interface),
	5340:  uint16(sym_method),
	5341:  uint16(aux_sym_interface_repeat1),
	5342:  uint16(8),
	5343:  uint16(3),
	5344:  uint16(1),
	5345:  uint16(sym_comment),
	5346:  uint16(107),
	5347:  uint16(1),
	5348:  uint16(anon_sym_enum),
	5349:  uint16(109),
	5350:  uint16(1),
	5351:  uint16(anon_sym_interface),
	5352:  uint16(111),
	5353:  uint16(1),
	5354:  uint16(anon_sym_struct),
	5355:  uint16(389),
	5356:  uint16(1),
	5357:  uint16(sym_identifier),
	5358:  uint16(427),
	5359:  uint16(1),
	5360:  uint16(anon_sym_RBRACE),
	5361:  uint16(298),
	5362:  uint16(1),
	5363:  uint16(sym__method_identifier),
	5364:  uint16(113),
	5365:  uint16(5),
	5366:  uint16(sym_struct),
	5367:  uint16(sym_enum),
	5368:  uint16(sym_interface),
	5369:  uint16(sym_method),
	5370:  uint16(aux_sym_interface_repeat1),
	5371:  uint16(8),
	5372:  uint16(3),
	5373:  uint16(1),
	5374:  uint16(sym_comment),
	5375:  uint16(107),
	5376:  uint16(1),
	5377:  uint16(anon_sym_enum),
	5378:  uint16(109),
	5379:  uint16(1),
	5380:  uint16(anon_sym_interface),
	5381:  uint16(111),
	5382:  uint16(1),
	5383:  uint16(anon_sym_struct),
	5384:  uint16(389),
	5385:  uint16(1),
	5386:  uint16(sym_identifier),
	5387:  uint16(429),
	5388:  uint16(1),
	5389:  uint16(anon_sym_RBRACE),
	5390:  uint16(298),
	5391:  uint16(1),
	5392:  uint16(sym__method_identifier),
	5393:  uint16(136),
	5394:  uint16(5),
	5395:  uint16(sym_struct),
	5396:  uint16(sym_enum),
	5397:  uint16(sym_interface),
	5398:  uint16(sym_method),
	5399:  uint16(aux_sym_interface_repeat1),
	5400:  uint16(8),
	5401:  uint16(3),
	5402:  uint16(1),
	5403:  uint16(sym_comment),
	5404:  uint16(431),
	5405:  uint16(1),
	5406:  uint16(sym_identifier),
	5407:  uint16(434),
	5408:  uint16(1),
	5409:  uint16(anon_sym_enum),
	5410:  uint16(437),
	5411:  uint16(1),
	5412:  uint16(anon_sym_interface),
	5413:  uint16(440),
	5414:  uint16(1),
	5415:  uint16(anon_sym_struct),
	5416:  uint16(443),
	5417:  uint16(1),
	5418:  uint16(anon_sym_RBRACE),
	5419:  uint16(298),
	5420:  uint16(1),
	5421:  uint16(sym__method_identifier),
	5422:  uint16(136),
	5423:  uint16(5),
	5424:  uint16(sym_struct),
	5425:  uint16(sym_enum),
	5426:  uint16(sym_interface),
	5427:  uint16(sym_method),
	5428:  uint16(aux_sym_interface_repeat1),
	5429:  uint16(8),
	5430:  uint16(3),
	5431:  uint16(1),
	5432:  uint16(sym_comment),
	5433:  uint16(107),
	5434:  uint16(1),
	5435:  uint16(anon_sym_enum),
	5436:  uint16(109),
	5437:  uint16(1),
	5438:  uint16(anon_sym_interface),
	5439:  uint16(111),
	5440:  uint16(1),
	5441:  uint16(anon_sym_struct),
	5442:  uint16(389),
	5443:  uint16(1),
	5444:  uint16(sym_identifier),
	5445:  uint16(413),
	5446:  uint16(1),
	5447:  uint16(anon_sym_RBRACE),
	5448:  uint16(298),
	5449:  uint16(1),
	5450:  uint16(sym__method_identifier),
	5451:  uint16(136),
	5452:  uint16(5),
	5453:  uint16(sym_struct),
	5454:  uint16(sym_enum),
	5455:  uint16(sym_interface),
	5456:  uint16(sym_method),
	5457:  uint16(aux_sym_interface_repeat1),
	5458:  uint16(8),
	5459:  uint16(3),
	5460:  uint16(1),
	5461:  uint16(sym_comment),
	5462:  uint16(107),
	5463:  uint16(1),
	5464:  uint16(anon_sym_enum),
	5465:  uint16(109),
	5466:  uint16(1),
	5467:  uint16(anon_sym_interface),
	5468:  uint16(111),
	5469:  uint16(1),
	5470:  uint16(anon_sym_struct),
	5471:  uint16(389),
	5472:  uint16(1),
	5473:  uint16(sym_identifier),
	5474:  uint16(429),
	5475:  uint16(1),
	5476:  uint16(anon_sym_RBRACE),
	5477:  uint16(298),
	5478:  uint16(1),
	5479:  uint16(sym__method_identifier),
	5480:  uint16(117),
	5481:  uint16(5),
	5482:  uint16(sym_struct),
	5483:  uint16(sym_enum),
	5484:  uint16(sym_interface),
	5485:  uint16(sym_method),
	5486:  uint16(aux_sym_interface_repeat1),
	5487:  uint16(8),
	5488:  uint16(3),
	5489:  uint16(1),
	5490:  uint16(sym_comment),
	5491:  uint16(445),
	5492:  uint16(1),
	5493:  uint16(sym_identifier),
	5494:  uint16(447),
	5495:  uint16(1),
	5496:  uint16(anon_sym_union),
	5497:  uint16(449),
	5498:  uint16(1),
	5499:  uint16(anon_sym_RBRACE),
	5500:  uint16(376),
	5501:  uint16(1),
	5502:  uint16(sym_union),
	5503:  uint16(150),
	5504:  uint16(2),
	5505:  uint16(sym_union_field),
	5506:  uint16(aux_sym__unnamed_union_repeat1),
	5507:  uint16(182),
	5508:  uint16(2),
	5509:  uint16(sym__unnamed_union),
	5510:  uint16(sym__named_union),
	5511:  uint16(378),
	5512:  uint16(2),
	5513:  uint16(sym_group),
	5514:  uint16(sym_nested_union),
	5515:  uint16(3),
	5516:  uint16(3),
	5517:  uint16(1),
	5518:  uint16(sym_comment),
	5519:  uint16(453),
	5520:  uint16(1),
	5521:  uint16(anon_sym_DOLLAR),
	5522:  uint16(451),
	5523:  uint16(9),
	5525:  uint16(sym_unique_id),
	5526:  uint16(anon_sym_using),
	5527:  uint16(anon_sym_DOLLARimport),
	5528:  uint16(anon_sym_annotation),
	5529:  uint16(anon_sym_const),
	5530:  uint16(anon_sym_enum),
	5531:  uint16(anon_sym_interface),
	5532:  uint16(anon_sym_struct),
	5533:  uint16(8),
	5534:  uint16(3),
	5535:  uint16(1),
	5536:  uint16(sym_comment),
	5537:  uint16(445),
	5538:  uint16(1),
	5539:  uint16(sym_identifier),
	5540:  uint16(447),
	5541:  uint16(1),
	5542:  uint16(anon_sym_union),
	5543:  uint16(449),
	5544:  uint16(1),
	5545:  uint16(anon_sym_RBRACE),
	5546:  uint16(376),
	5547:  uint16(1),
	5548:  uint16(sym_union),
	5549:  uint16(153),
	5550:  uint16(2),
	5551:  uint16(sym_union_field),
	5552:  uint16(aux_sym__unnamed_union_repeat1),
	5553:  uint16(182),
	5554:  uint16(2),
	5555:  uint16(sym__unnamed_union),
	5556:  uint16(sym__named_union),
	5557:  uint16(378),
	5558:  uint16(2),
	5559:  uint16(sym_group),
	5560:  uint16(sym_nested_union),
	5561:  uint16(5),
	5562:  uint16(3),
	5563:  uint16(1),
	5564:  uint16(sym_comment),
	5565:  uint16(457),
	5566:  uint16(1),
	5567:  uint16(anon_sym_DQUOTE),
	5568:  uint16(460),
	5569:  uint16(1),
	5570:  uint16(anon_sym_SQUOTE),
	5571:  uint16(142),
	5572:  uint16(2),
	5573:  uint16(sym_string),
	5574:  uint16(aux_sym_concatenated_string_repeat1),
	5575:  uint16(455),
	5576:  uint16(6),
	5577:  uint16(anon_sym_SEMI),
	5578:  uint16(anon_sym_RPAREN),
	5579:  uint16(anon_sym_DOLLAR),
	5580:  uint16(anon_sym_COMMA),
	5581:  uint16(anon_sym_RBRACK),
	5582:  uint16(sym_identifier),
	5583:  uint16(3),
	5584:  uint16(3),
	5585:  uint16(1),
	5586:  uint16(sym_comment),
	5587:  uint16(465),
	5588:  uint16(1),
	5589:  uint16(anon_sym_DOLLAR),
	5590:  uint16(463),
	5591:  uint16(9),
	5593:  uint16(sym_unique_id),
	5594:  uint16(anon_sym_using),
	5595:  uint16(anon_sym_DOLLARimport),
	5596:  uint16(anon_sym_annotation),
	5597:  uint16(anon_sym_const),
	5598:  uint16(anon_sym_enum),
	5599:  uint16(anon_sym_interface),
	5600:  uint16(anon_sym_struct),
	5601:  uint16(3),
	5602:  uint16(3),
	5603:  uint16(1),
	5604:  uint16(sym_comment),
	5605:  uint16(469),
	5606:  uint16(1),
	5607:  uint16(anon_sym_DOLLAR),
	5608:  uint16(467),
	5609:  uint16(9),
	5611:  uint16(sym_unique_id),
	5612:  uint16(anon_sym_using),
	5613:  uint16(anon_sym_DOLLARimport),
	5614:  uint16(anon_sym_annotation),
	5615:  uint16(anon_sym_const),
	5616:  uint16(anon_sym_enum),
	5617:  uint16(anon_sym_interface),
	5618:  uint16(anon_sym_struct),
	5619:  uint16(8),
	5620:  uint16(3),
	5621:  uint16(1),
	5622:  uint16(sym_comment),
	5623:  uint16(445),
	5624:  uint16(1),
	5625:  uint16(sym_identifier),
	5626:  uint16(447),
	5627:  uint16(1),
	5628:  uint16(anon_sym_union),
	5629:  uint16(471),
	5630:  uint16(1),
	5631:  uint16(anon_sym_RBRACE),
	5632:  uint16(376),
	5633:  uint16(1),
	5634:  uint16(sym_union),
	5635:  uint16(157),
	5636:  uint16(2),
	5637:  uint16(sym_union_field),
	5638:  uint16(aux_sym__unnamed_union_repeat1),
	5639:  uint16(182),
	5640:  uint16(2),
	5641:  uint16(sym__unnamed_union),
	5642:  uint16(sym__named_union),
	5643:  uint16(378),
	5644:  uint16(2),
	5645:  uint16(sym_group),
	5646:  uint16(sym_nested_union),
	5647:  uint16(5),
	5648:  uint16(3),
	5649:  uint16(1),
	5650:  uint16(sym_comment),
	5651:  uint16(45),
	5652:  uint16(1),
	5653:  uint16(anon_sym_DQUOTE),
	5654:  uint16(47),
	5655:  uint16(1),
	5656:  uint16(anon_sym_SQUOTE),
	5657:  uint16(142),
	5658:  uint16(2),
	5659:  uint16(sym_string),
	5660:  uint16(aux_sym_concatenated_string_repeat1),
	5661:  uint16(473),
	5662:  uint16(6),
	5663:  uint16(anon_sym_SEMI),
	5664:  uint16(anon_sym_RPAREN),
	5665:  uint16(anon_sym_DOLLAR),
	5666:  uint16(anon_sym_COMMA),
	5667:  uint16(anon_sym_RBRACK),
	5668:  uint16(sym_identifier),
	5669:  uint16(5),
	5670:  uint16(3),
	5671:  uint16(1),
	5672:  uint16(sym_comment),
	5673:  uint16(45),
	5674:  uint16(1),
	5675:  uint16(anon_sym_DQUOTE),
	5676:  uint16(47),
	5677:  uint16(1),
	5678:  uint16(anon_sym_SQUOTE),
	5679:  uint16(146),
	5680:  uint16(2),
	5681:  uint16(sym_string),
	5682:  uint16(aux_sym_concatenated_string_repeat1),
	5683:  uint16(475),
	5684:  uint16(6),
	5685:  uint16(anon_sym_SEMI),
	5686:  uint16(anon_sym_RPAREN),
	5687:  uint16(anon_sym_DOLLAR),
	5688:  uint16(anon_sym_COMMA),
	5689:  uint16(anon_sym_RBRACK),
	5690:  uint16(sym_identifier),
	5691:  uint16(8),
	5692:  uint16(3),
	5693:  uint16(1),
	5694:  uint16(sym_comment),
	5695:  uint16(445),
	5696:  uint16(1),
	5697:  uint16(sym_identifier),
	5698:  uint16(447),
	5699:  uint16(1),
	5700:  uint16(anon_sym_union),
	5701:  uint16(477),
	5702:  uint16(1),
	5703:  uint16(anon_sym_RBRACE),
	5704:  uint16(376),
	5705:  uint16(1),
	5706:  uint16(sym_union),
	5707:  uint16(156),
	5708:  uint16(2),
	5709:  uint16(sym_union_field),
	5710:  uint16(aux_sym__unnamed_union_repeat1),
	5711:  uint16(182),
	5712:  uint16(2),
	5713:  uint16(sym__unnamed_union),
	5714:  uint16(sym__named_union),
	5715:  uint16(378),
	5716:  uint16(2),
	5717:  uint16(sym_group),
	5718:  uint16(sym_nested_union),
	5719:  uint16(6),
	5720:  uint16(3),
	5721:  uint16(1),
	5722:  uint16(sym_comment),
	5723:  uint16(403),
	5724:  uint16(1),
	5725:  uint16(anon_sym_DOT),
	5726:  uint16(481),
	5727:  uint16(1),
	5728:  uint16(anon_sym_LPAREN),
	5729:  uint16(168),
	5730:  uint16(1),
	5731:  uint16(aux_sym__annotation_call_repeat1),
	5732:  uint16(201),
	5733:  uint16(1),
	5734:  uint16(sym_annotation_literal),
	5735:  uint16(479),
	5736:  uint16(6),
	5737:  uint16(anon_sym_SEMI),
	5738:  uint16(anon_sym_EQ),
	5739:  uint16(anon_sym_RPAREN),
	5740:  uint16(anon_sym_DOLLAR),
	5741:  uint16(anon_sym_COMMA),
	5742:  uint16(anon_sym_LBRACE),
	5743:  uint16(8),
	5744:  uint16(3),
	5745:  uint16(1),
	5746:  uint16(sym_comment),
	5747:  uint16(483),
	5748:  uint16(1),
	5749:  uint16(sym_identifier),
	5750:  uint16(486),
	5751:  uint16(1),
	5752:  uint16(anon_sym_union),
	5753:  uint16(489),
	5754:  uint16(1),
	5755:  uint16(anon_sym_RBRACE),
	5756:  uint16(376),
	5757:  uint16(1),
	5758:  uint16(sym_union),
	5759:  uint16(150),
	5760:  uint16(2),
	5761:  uint16(sym_union_field),
	5762:  uint16(aux_sym__unnamed_union_repeat1),
	5763:  uint16(182),
	5764:  uint16(2),
	5765:  uint16(sym__unnamed_union),
	5766:  uint16(sym__named_union),
	5767:  uint16(378),
	5768:  uint16(2),
	5769:  uint16(sym_group),
	5770:  uint16(sym_nested_union),
	5771:  uint16(6),
	5772:  uint16(3),
	5773:  uint16(1),
	5774:  uint16(sym_comment),
	5775:  uint16(403),
	5776:  uint16(1),
	5777:  uint16(anon_sym_DOT),
	5778:  uint16(493),
	5779:  uint16(1),
	5780:  uint16(anon_sym_LPAREN),
	5781:  uint16(168),
	5782:  uint16(1),
	5783:  uint16(aux_sym__annotation_call_repeat1),
	5784:  uint16(235),
	5785:  uint16(1),
	5786:  uint16(sym_annotation_literal),
	5787:  uint16(491),
	5788:  uint16(6),
	5789:  uint16(anon_sym_SEMI),
	5790:  uint16(anon_sym_EQ),
	5791:  uint16(anon_sym_RPAREN),
	5792:  uint16(anon_sym_DOLLAR),
	5793:  uint16(anon_sym_COMMA),
	5794:  uint16(anon_sym_LBRACE),
	5795:  uint16(6),
	5796:  uint16(3),
	5797:  uint16(1),
	5798:  uint16(sym_comment),
	5799:  uint16(403),
	5800:  uint16(1),
	5801:  uint16(anon_sym_DOT),
	5802:  uint16(497),
	5803:  uint16(1),
	5804:  uint16(anon_sym_LPAREN),
	5805:  uint16(151),
	5806:  uint16(1),
	5807:  uint16(aux_sym__annotation_call_repeat1),
	5808:  uint16(228),
	5809:  uint16(1),
	5810:  uint16(sym_annotation_literal),
	5811:  uint16(495),
	5812:  uint16(6),
	5813:  uint16(anon_sym_SEMI),
	5814:  uint16(anon_sym_EQ),
	5815:  uint16(anon_sym_RPAREN),
	5816:  uint16(anon_sym_DOLLAR),
	5817:  uint16(anon_sym_COMMA),
	5818:  uint16(anon_sym_LBRACE),
	5819:  uint16(8),
	5820:  uint16(3),
	5821:  uint16(1),
	5822:  uint16(sym_comment),
	5823:  uint16(445),
	5824:  uint16(1),
	5825:  uint16(sym_identifier),
	5826:  uint16(447),
	5827:  uint16(1),
	5828:  uint16(anon_sym_union),
	5829:  uint16(477),
	5830:  uint16(1),
	5831:  uint16(anon_sym_RBRACE),
	5832:  uint16(376),
	5833:  uint16(1),
	5834:  uint16(sym_union),
	5835:  uint16(150),
	5836:  uint16(2),
	5837:  uint16(sym_union_field),
	5838:  uint16(aux_sym__unnamed_union_repeat1),
	5839:  uint16(182),
	5840:  uint16(2),
	5841:  uint16(sym__unnamed_union),
	5842:  uint16(sym__named_union),
	5843:  uint16(378),
	5844:  uint16(2),
	5845:  uint16(sym_group),
	5846:  uint16(sym_nested_union),
	5847:  uint16(6),
	5848:  uint16(3),
	5849:  uint16(1),
	5850:  uint16(sym_comment),
	5851:  uint16(501),
	5852:  uint16(1),
	5853:  uint16(anon_sym_DOT),
	5854:  uint16(503),
	5855:  uint16(1),
	5856:  uint16(anon_sym_LPAREN),
	5857:  uint16(192),
	5858:  uint16(1),
	5859:  uint16(aux_sym_custom_type_repeat1),
	5860:  uint16(193),
	5861:  uint16(1),
	5862:  uint16(sym_generics),
	5863:  uint16(499),
	5864:  uint16(6),
	5865:  uint16(anon_sym_SEMI),
	5866:  uint16(anon_sym_EQ),
	5867:  uint16(anon_sym_RPAREN),
	5868:  uint16(anon_sym_DOLLAR),
	5869:  uint16(anon_sym_COMMA),
	5870:  uint16(anon_sym_RBRACK),
	5871:  uint16(6),
	5872:  uint16(3),
	5873:  uint16(1),
	5874:  uint16(sym_comment),
	5875:  uint16(403),
	5876:  uint16(1),
	5877:  uint16(anon_sym_DOT),
	5878:  uint16(505),
	5879:  uint16(1),
	5880:  uint16(anon_sym_LPAREN),
	5881:  uint16(168),
	5882:  uint16(1),
	5883:  uint16(aux_sym__annotation_call_repeat1),
	5884:  uint16(228),
	5885:  uint16(1),
	5886:  uint16(sym_annotation_literal),
	5887:  uint16(495),
	5888:  uint16(6),
	5889:  uint16(anon_sym_SEMI),
	5890:  uint16(anon_sym_EQ),
	5891:  uint16(anon_sym_RPAREN),
	5892:  uint16(anon_sym_DOLLAR),
	5893:  uint16(anon_sym_COMMA),
	5894:  uint16(anon_sym_LBRACE),
	5895:  uint16(8),
	5896:  uint16(3),
	5897:  uint16(1),
	5898:  uint16(sym_comment),
	5899:  uint16(445),
	5900:  uint16(1),
	5901:  uint16(sym_identifier),
	5902:  uint16(447),
	5903:  uint16(1),
	5904:  uint16(anon_sym_union),
	5905:  uint16(507),
	5906:  uint16(1),
	5907:  uint16(anon_sym_RBRACE),
	5908:  uint16(376),
	5909:  uint16(1),
	5910:  uint16(sym_union),
	5911:  uint16(150),
	5912:  uint16(2),
	5913:  uint16(sym_union_field),
	5914:  uint16(aux_sym__unnamed_union_repeat1),
	5915:  uint16(182),
	5916:  uint16(2),
	5917:  uint16(sym__unnamed_union),
	5918:  uint16(sym__named_union),
	5919:  uint16(378),
	5920:  uint16(2),
	5921:  uint16(sym_group),
	5922:  uint16(sym_nested_union),
	5923:  uint16(8),
	5924:  uint16(3),
	5925:  uint16(1),
	5926:  uint16(sym_comment),
	5927:  uint16(445),
	5928:  uint16(1),
	5929:  uint16(sym_identifier),
	5930:  uint16(447),
	5931:  uint16(1),
	5932:  uint16(anon_sym_union),
	5933:  uint16(509),
	5934:  uint16(1),
	5935:  uint16(anon_sym_RBRACE),
	5936:  uint16(376),
	5937:  uint16(1),
	5938:  uint16(sym_union),
	5939:  uint16(150),
	5940:  uint16(2),
	5941:  uint16(sym_union_field),
	5942:  uint16(aux_sym__unnamed_union_repeat1),
	5943:  uint16(182),
	5944:  uint16(2),
	5945:  uint16(sym__unnamed_union),
	5946:  uint16(sym__named_union),
	5947:  uint16(378),
	5948:  uint16(2),
	5949:  uint16(sym_group),
	5950:  uint16(sym_nested_union),
	5951:  uint16(6),
	5952:  uint16(3),
	5953:  uint16(1),
	5954:  uint16(sym_comment),
	5955:  uint16(403),
	5956:  uint16(1),
	5957:  uint16(anon_sym_DOT),
	5958:  uint16(513),
	5959:  uint16(1),
	5960:  uint16(anon_sym_LPAREN),
	5961:  uint16(168),
	5962:  uint16(1),
	5963:  uint16(aux_sym__annotation_call_repeat1),
	5964:  uint16(209),
	5965:  uint16(1),
	5966:  uint16(sym_annotation_literal),
	5967:  uint16(511),
	5968:  uint16(6),
	5969:  uint16(anon_sym_SEMI),
	5970:  uint16(anon_sym_EQ),
	5971:  uint16(anon_sym_RPAREN),
	5972:  uint16(anon_sym_DOLLAR),
	5973:  uint16(anon_sym_COMMA),
	5974:  uint16(anon_sym_LBRACE),
	5975:  uint16(3),
	5976:  uint16(3),
	5977:  uint16(1),
	5978:  uint16(sym_comment),
	5979:  uint16(517),
	5980:  uint16(1),
	5981:  uint16(anon_sym_DOLLAR),
	5982:  uint16(515),
	5983:  uint16(9),
	5985:  uint16(sym_unique_id),
	5986:  uint16(anon_sym_using),
	5987:  uint16(anon_sym_DOLLARimport),
	5988:  uint16(anon_sym_annotation),
	5989:  uint16(anon_sym_const),
	5990:  uint16(anon_sym_enum),
	5991:  uint16(anon_sym_interface),
	5992:  uint16(anon_sym_struct),
	5993:  uint16(3),
	5994:  uint16(3),
	5995:  uint16(1),
	5996:  uint16(sym_comment),
	5997:  uint16(521),
	5998:  uint16(1),
	5999:  uint16(anon_sym_DOLLAR),
	6000:  uint16(519),
	6001:  uint16(9),
	6003:  uint16(sym_unique_id),
	6004:  uint16(anon_sym_using),
	6005:  uint16(anon_sym_DOLLARimport),
	6006:  uint16(anon_sym_annotation),
	6007:  uint16(anon_sym_const),
	6008:  uint16(anon_sym_enum),
	6009:  uint16(anon_sym_interface),
	6010:  uint16(anon_sym_struct),
	6011:  uint16(3),
	6012:  uint16(3),
	6013:  uint16(1),
	6014:  uint16(sym_comment),
	6015:  uint16(525),
	6016:  uint16(1),
	6017:  uint16(anon_sym_DOLLAR),
	6018:  uint16(523),
	6019:  uint16(9),
	6021:  uint16(sym_unique_id),
	6022:  uint16(anon_sym_using),
	6023:  uint16(anon_sym_DOLLARimport),
	6024:  uint16(anon_sym_annotation),
	6025:  uint16(anon_sym_const),
	6026:  uint16(anon_sym_enum),
	6027:  uint16(anon_sym_interface),
	6028:  uint16(anon_sym_struct),
	6029:  uint16(3),
	6030:  uint16(3),
	6031:  uint16(1),
	6032:  uint16(sym_comment),
	6033:  uint16(529),
	6034:  uint16(1),
	6035:  uint16(anon_sym_DOLLAR),
	6036:  uint16(527),
	6037:  uint16(9),
	6039:  uint16(sym_unique_id),
	6040:  uint16(anon_sym_using),
	6041:  uint16(anon_sym_DOLLARimport),
	6042:  uint16(anon_sym_annotation),
	6043:  uint16(anon_sym_const),
	6044:  uint16(anon_sym_enum),
	6045:  uint16(anon_sym_interface),
	6046:  uint16(anon_sym_struct),
	6047:  uint16(3),
	6048:  uint16(3),
	6049:  uint16(1),
	6050:  uint16(sym_comment),
	6051:  uint16(533),
	6052:  uint16(1),
	6053:  uint16(anon_sym_DOLLAR),
	6054:  uint16(531),
	6055:  uint16(9),
	6057:  uint16(sym_unique_id),
	6058:  uint16(anon_sym_using),
	6059:  uint16(anon_sym_DOLLARimport),
	6060:  uint16(anon_sym_annotation),
	6061:  uint16(anon_sym_const),
	6062:  uint16(anon_sym_enum),
	6063:  uint16(anon_sym_interface),
	6064:  uint16(anon_sym_struct),
	6065:  uint16(8),
	6066:  uint16(3),
	6067:  uint16(1),
	6068:  uint16(sym_comment),
	6069:  uint16(445),
	6070:  uint16(1),
	6071:  uint16(sym_identifier),
	6072:  uint16(447),
	6073:  uint16(1),
	6074:  uint16(anon_sym_union),
	6075:  uint16(535),
	6076:  uint16(1),
	6077:  uint16(anon_sym_RBRACE),
	6078:  uint16(376),
	6079:  uint16(1),
	6080:  uint16(sym_union),
	6081:  uint16(139),
	6082:  uint16(2),
	6083:  uint16(sym_union_field),
	6084:  uint16(aux_sym__unnamed_union_repeat1),
	6085:  uint16(182),
	6086:  uint16(2),
	6087:  uint16(sym__unnamed_union),
	6088:  uint16(sym__named_union),
	6089:  uint16(378),
	6090:  uint16(2),
	6091:  uint16(sym_group),
	6092:  uint16(sym_nested_union),
	6093:  uint16(3),
	6094:  uint16(3),
	6095:  uint16(1),
	6096:  uint16(sym_comment),
	6097:  uint16(539),
	6098:  uint16(1),
	6099:  uint16(anon_sym_RBRACE),
	6100:  uint16(537),
	6101:  uint16(8),
	6102:  uint16(anon_sym_using),
	6103:  uint16(anon_sym_annotation),
	6104:  uint16(anon_sym_const),
	6105:  uint16(anon_sym_enum),
	6106:  uint16(anon_sym_interface),
	6107:  uint16(anon_sym_struct),
	6108:  uint16(anon_sym_union),
	6109:  uint16(sym_identifier),
	6110:  uint16(3),
	6111:  uint16(3),
	6112:  uint16(1),
	6113:  uint16(sym_comment),
	6114:  uint16(543),
	6115:  uint16(1),
	6116:  uint16(anon_sym_RBRACE),
	6117:  uint16(541),
	6118:  uint16(8),
	6119:  uint16(anon_sym_using),
	6120:  uint16(anon_sym_annotation),
	6121:  uint16(anon_sym_const),
	6122:  uint16(anon_sym_enum),
	6123:  uint16(anon_sym_interface),
	6124:  uint16(anon_sym_struct),
	6125:  uint16(anon_sym_union),
	6126:  uint16(sym_identifier),
	6127:  uint16(4),
	6128:  uint16(3),
	6129:  uint16(1),
	6130:  uint16(sym_comment),
	6131:  uint16(503),
	6132:  uint16(1),
	6133:  uint16(anon_sym_LPAREN),
	6134:  uint16(197),
	6135:  uint16(1),
	6136:  uint16(sym_generics),
	6137:  uint16(545),
	6138:  uint16(7),
	6139:  uint16(anon_sym_SEMI),
	6140:  uint16(anon_sym_EQ),
	6141:  uint16(anon_sym_DOT),
	6142:  uint16(anon_sym_RPAREN),
	6143:  uint16(anon_sym_DOLLAR),
	6144:  uint16(anon_sym_COMMA),
	6145:  uint16(anon_sym_RBRACK),
	6146:  uint16(4),
	6147:  uint16(3),
	6148:  uint16(1),
	6149:  uint16(sym_comment),
	6150:  uint16(549),
	6151:  uint16(1),
	6152:  uint16(anon_sym_DOT),
	6153:  uint16(168),
	6154:  uint16(1),
	6155:  uint16(aux_sym__annotation_call_repeat1),
	6156:  uint16(547),
	6157:  uint16(7),
	6158:  uint16(anon_sym_SEMI),
	6159:  uint16(anon_sym_EQ),
	6160:  uint16(anon_sym_LPAREN),
	6161:  uint16(anon_sym_RPAREN),
	6162:  uint16(anon_sym_DOLLAR),
	6163:  uint16(anon_sym_COMMA),
	6164:  uint16(anon_sym_LBRACE),
	6165:  uint16(3),
	6166:  uint16(3),
	6167:  uint16(1),
	6168:  uint16(sym_comment),
	6169:  uint16(554),
	6170:  uint16(1),
	6171:  uint16(anon_sym_RBRACE),
	6172:  uint16(552),
	6173:  uint16(8),
	6174:  uint16(anon_sym_using),
	6175:  uint16(anon_sym_annotation),
	6176:  uint16(anon_sym_const),
	6177:  uint16(anon_sym_enum),
	6178:  uint16(anon_sym_interface),
	6179:  uint16(anon_sym_struct),
	6180:  uint16(anon_sym_union),
	6181:  uint16(sym_identifier),
	6182:  uint16(3),
	6183:  uint16(3),
	6184:  uint16(1),
	6185:  uint16(sym_comment),
	6186:  uint16(558),
	6187:  uint16(1),
	6188:  uint16(anon_sym_RBRACE),
	6189:  uint16(556),
	6190:  uint16(8),
	6191:  uint16(anon_sym_using),
	6192:  uint16(anon_sym_annotation),
	6193:  uint16(anon_sym_const),
	6194:  uint16(anon_sym_enum),
	6195:  uint16(anon_sym_interface),
	6196:  uint16(anon_sym_struct),
	6197:  uint16(anon_sym_union),
	6198:  uint16(sym_identifier),
	6199:  uint16(3),
	6200:  uint16(3),
	6201:  uint16(1),
	6202:  uint16(sym_comment),
	6203:  uint16(562),
	6204:  uint16(1),
	6205:  uint16(anon_sym_RBRACE),
	6206:  uint16(560),
	6207:  uint16(8),
	6208:  uint16(anon_sym_using),
	6209:  uint16(anon_sym_annotation),
	6210:  uint16(anon_sym_const),
	6211:  uint16(anon_sym_enum),
	6212:  uint16(anon_sym_interface),
	6213:  uint16(anon_sym_struct),
	6214:  uint16(anon_sym_union),
	6215:  uint16(sym_identifier),
	6216:  uint16(2),
	6217:  uint16(3),
	6218:  uint16(1),
	6219:  uint16(sym_comment),
	6220:  uint16(564),
	6221:  uint16(9),
	6222:  uint16(anon_sym_SEMI),
	6223:  uint16(anon_sym_DOT),
	6224:  uint16(anon_sym_RPAREN),
	6225:  uint16(anon_sym_DOLLAR),
	6226:  uint16(anon_sym_COMMA),
	6227:  uint16(anon_sym_RBRACK),
	6228:  uint16(anon_sym_DQUOTE),
	6229:  uint16(anon_sym_SQUOTE),
	6230:  uint16(sym_identifier),
	6231:  uint16(3),
	6232:  uint16(3),
	6233:  uint16(1),
	6234:  uint16(sym_comment),
	6235:  uint16(568),
	6236:  uint16(1),
	6237:  uint16(anon_sym_RBRACE),
	6238:  uint16(566),
	6239:  uint16(8),
	6240:  uint16(anon_sym_using),
	6241:  uint16(anon_sym_annotation),
	6242:  uint16(anon_sym_const),
	6243:  uint16(anon_sym_enum),
	6244:  uint16(anon_sym_interface),
	6245:  uint16(anon_sym_struct),
	6246:  uint16(anon_sym_union),
	6247:  uint16(sym_identifier),
	6248:  uint16(3),
	6249:  uint16(3),
	6250:  uint16(1),
	6251:  uint16(sym_comment),
	6252:  uint16(572),
	6253:  uint16(1),
	6254:  uint16(anon_sym_RBRACE),
	6255:  uint16(570),
	6256:  uint16(8),
	6257:  uint16(anon_sym_using),
	6258:  uint16(anon_sym_annotation),
	6259:  uint16(anon_sym_const),
	6260:  uint16(anon_sym_enum),
	6261:  uint16(anon_sym_interface),
	6262:  uint16(anon_sym_struct),
	6263:  uint16(anon_sym_union),
	6264:  uint16(sym_identifier),
	6265:  uint16(3),
	6266:  uint16(3),
	6267:  uint16(1),
	6268:  uint16(sym_comment),
	6269:  uint16(576),
	6270:  uint16(1),
	6271:  uint16(anon_sym_RBRACE),
	6272:  uint16(574),
	6273:  uint16(8),
	6274:  uint16(anon_sym_using),
	6275:  uint16(anon_sym_annotation),
	6276:  uint16(anon_sym_const),
	6277:  uint16(anon_sym_enum),
	6278:  uint16(anon_sym_interface),
	6279:  uint16(anon_sym_struct),
	6280:  uint16(anon_sym_union),
	6281:  uint16(sym_identifier),
	6282:  uint16(2),
	6283:  uint16(3),
	6284:  uint16(1),
	6285:  uint16(sym_comment),
	6286:  uint16(578),
	6287:  uint16(9),
	6288:  uint16(anon_sym_SEMI),
	6289:  uint16(anon_sym_DOT),
	6290:  uint16(anon_sym_RPAREN),
	6291:  uint16(anon_sym_DOLLAR),
	6292:  uint16(anon_sym_COMMA),
	6293:  uint16(anon_sym_RBRACK),
	6294:  uint16(anon_sym_DQUOTE),
	6295:  uint16(anon_sym_SQUOTE),
	6296:  uint16(sym_identifier),
	6297:  uint16(3),
	6298:  uint16(3),
	6299:  uint16(1),
	6300:  uint16(sym_comment),
	6301:  uint16(582),
	6302:  uint16(1),
	6303:  uint16(anon_sym_RBRACE),
	6304:  uint16(580),
	6305:  uint16(8),
	6306:  uint16(anon_sym_using),
	6307:  uint16(anon_sym_annotation),
	6308:  uint16(anon_sym_const),
	6309:  uint16(anon_sym_enum),
	6310:  uint16(anon_sym_interface),
	6311:  uint16(anon_sym_struct),
	6312:  uint16(anon_sym_union),
	6313:  uint16(sym_identifier),
	6314:  uint16(3),
	6315:  uint16(3),
	6316:  uint16(1),
	6317:  uint16(sym_comment),
	6318:  uint16(586),
	6319:  uint16(1),
	6320:  uint16(anon_sym_RBRACE),
	6321:  uint16(584),
	6322:  uint16(8),
	6323:  uint16(anon_sym_using),
	6324:  uint16(anon_sym_annotation),
	6325:  uint16(anon_sym_const),
	6326:  uint16(anon_sym_enum),
	6327:  uint16(anon_sym_interface),
	6328:  uint16(anon_sym_struct),
	6329:  uint16(anon_sym_union),
	6330:  uint16(sym_identifier),
	6331:  uint16(3),
	6332:  uint16(3),
	6333:  uint16(1),
	6334:  uint16(sym_comment),
	6335:  uint16(590),
	6336:  uint16(1),
	6337:  uint16(anon_sym_RBRACE),
	6338:  uint16(588),
	6339:  uint16(8),
	6340:  uint16(anon_sym_using),
	6341:  uint16(anon_sym_annotation),
	6342:  uint16(anon_sym_const),
	6343:  uint16(anon_sym_enum),
	6344:  uint16(anon_sym_interface),
	6345:  uint16(anon_sym_struct),
	6346:  uint16(anon_sym_union),
	6347:  uint16(sym_identifier),
	6348:  uint16(3),
	6349:  uint16(3),
	6350:  uint16(1),
	6351:  uint16(sym_comment),
	6352:  uint16(594),
	6353:  uint16(1),
	6354:  uint16(anon_sym_RBRACE),
	6355:  uint16(592),
	6356:  uint16(8),
	6357:  uint16(anon_sym_using),
	6358:  uint16(anon_sym_annotation),
	6359:  uint16(anon_sym_const),
	6360:  uint16(anon_sym_enum),
	6361:  uint16(anon_sym_interface),
	6362:  uint16(anon_sym_struct),
	6363:  uint16(anon_sym_union),
	6364:  uint16(sym_identifier),
	6365:  uint16(3),
	6366:  uint16(3),
	6367:  uint16(1),
	6368:  uint16(sym_comment),
	6369:  uint16(598),
	6370:  uint16(1),
	6371:  uint16(anon_sym_RBRACE),
	6372:  uint16(596),
	6373:  uint16(8),
	6374:  uint16(anon_sym_using),
	6375:  uint16(anon_sym_annotation),
	6376:  uint16(anon_sym_const),
	6377:  uint16(anon_sym_enum),
	6378:  uint16(anon_sym_interface),
	6379:  uint16(anon_sym_struct),
	6380:  uint16(anon_sym_union),
	6381:  uint16(sym_identifier),
	6382:  uint16(3),
	6383:  uint16(3),
	6384:  uint16(1),
	6385:  uint16(sym_comment),
	6386:  uint16(602),
	6387:  uint16(1),
	6388:  uint16(anon_sym_RBRACE),
	6389:  uint16(600),
	6390:  uint16(8),
	6391:  uint16(anon_sym_using),
	6392:  uint16(anon_sym_annotation),
	6393:  uint16(anon_sym_const),
	6394:  uint16(anon_sym_enum),
	6395:  uint16(anon_sym_interface),
	6396:  uint16(anon_sym_struct),
	6397:  uint16(anon_sym_union),
	6398:  uint16(sym_identifier),
	6399:  uint16(3),
	6400:  uint16(3),
	6401:  uint16(1),
	6402:  uint16(sym_comment),
	6403:  uint16(606),
	6404:  uint16(1),
	6405:  uint16(anon_sym_RBRACE),
	6406:  uint16(604),
	6407:  uint16(8),
	6408:  uint16(anon_sym_using),
	6409:  uint16(anon_sym_annotation),
	6410:  uint16(anon_sym_const),
	6411:  uint16(anon_sym_enum),
	6412:  uint16(anon_sym_interface),
	6413:  uint16(anon_sym_struct),
	6414:  uint16(anon_sym_union),
	6415:  uint16(sym_identifier),
	6416:  uint16(3),
	6417:  uint16(3),
	6418:  uint16(1),
	6419:  uint16(sym_comment),
	6420:  uint16(610),
	6421:  uint16(1),
	6422:  uint16(anon_sym_RBRACE),
	6423:  uint16(608),
	6424:  uint16(8),
	6425:  uint16(anon_sym_using),
	6426:  uint16(anon_sym_annotation),
	6427:  uint16(anon_sym_const),
	6428:  uint16(anon_sym_enum),
	6429:  uint16(anon_sym_interface),
	6430:  uint16(anon_sym_struct),
	6431:  uint16(anon_sym_union),
	6432:  uint16(sym_identifier),
	6433:  uint16(3),
	6434:  uint16(3),
	6435:  uint16(1),
	6436:  uint16(sym_comment),
	6437:  uint16(614),
	6438:  uint16(1),
	6439:  uint16(anon_sym_RBRACE),
	6440:  uint16(612),
	6441:  uint16(8),
	6442:  uint16(anon_sym_using),
	6443:  uint16(anon_sym_annotation),
	6444:  uint16(anon_sym_const),
	6445:  uint16(anon_sym_enum),
	6446:  uint16(anon_sym_interface),
	6447:  uint16(anon_sym_struct),
	6448:  uint16(anon_sym_union),
	6449:  uint16(sym_identifier),
	6450:  uint16(3),
	6451:  uint16(3),
	6452:  uint16(1),
	6453:  uint16(sym_comment),
	6454:  uint16(618),
	6455:  uint16(1),
	6456:  uint16(anon_sym_RBRACE),
	6457:  uint16(616),
	6458:  uint16(8),
	6459:  uint16(anon_sym_using),
	6460:  uint16(anon_sym_annotation),
	6461:  uint16(anon_sym_const),
	6462:  uint16(anon_sym_enum),
	6463:  uint16(anon_sym_interface),
	6464:  uint16(anon_sym_struct),
	6465:  uint16(anon_sym_union),
	6466:  uint16(sym_identifier),
	6467:  uint16(3),
	6468:  uint16(3),
	6469:  uint16(1),
	6470:  uint16(sym_comment),
	6471:  uint16(622),
	6472:  uint16(1),
	6473:  uint16(anon_sym_RBRACE),
	6474:  uint16(620),
	6475:  uint16(8),
	6476:  uint16(anon_sym_using),
	6477:  uint16(anon_sym_annotation),
	6478:  uint16(anon_sym_const),
	6479:  uint16(anon_sym_enum),
	6480:  uint16(anon_sym_interface),
	6481:  uint16(anon_sym_struct),
	6482:  uint16(anon_sym_union),
	6483:  uint16(sym_identifier),
	6484:  uint16(4),
	6485:  uint16(3),
	6486:  uint16(1),
	6487:  uint16(sym_comment),
	6488:  uint16(626),
	6489:  uint16(1),
	6490:  uint16(anon_sym_DOT),
	6491:  uint16(188),
	6492:  uint16(1),
	6493:  uint16(aux_sym_custom_type_repeat1),
	6494:  uint16(624),
	6495:  uint16(6),
	6496:  uint16(anon_sym_SEMI),
	6497:  uint16(anon_sym_EQ),
	6498:  uint16(anon_sym_RPAREN),
	6499:  uint16(anon_sym_DOLLAR),
	6500:  uint16(anon_sym_COMMA),
	6501:  uint16(anon_sym_RBRACK),
	6502:  uint16(4),
	6503:  uint16(3),
	6504:  uint16(1),
	6505:  uint16(sym_comment),
	6506:  uint16(501),
	6507:  uint16(1),
	6508:  uint16(anon_sym_DOT),
	6509:  uint16(188),
	6510:  uint16(1),
	6511:  uint16(aux_sym_custom_type_repeat1),
	6512:  uint16(629),
	6513:  uint16(6),
	6514:  uint16(anon_sym_SEMI),
	6515:  uint16(anon_sym_EQ),
	6516:  uint16(anon_sym_RPAREN),
	6517:  uint16(anon_sym_DOLLAR),
	6518:  uint16(anon_sym_COMMA),
	6519:  uint16(anon_sym_RBRACK),
	6520:  uint16(4),
	6521:  uint16(3),
	6522:  uint16(1),
	6523:  uint16(sym_comment),
	6524:  uint16(49),
	6525:  uint16(1),
	6526:  uint16(anon_sym_BQUOTE),
	6527:  uint16(194),
	6528:  uint16(1),
	6529:  uint16(aux_sym_block_text_repeat2),
	6530:  uint16(631),
	6531:  uint16(6),
	6532:  uint16(anon_sym_SEMI),
	6533:  uint16(anon_sym_RPAREN),
	6534:  uint16(anon_sym_DOLLAR),
	6535:  uint16(anon_sym_COMMA),
	6536:  uint16(anon_sym_RBRACK),
	6537:  uint16(sym_identifier),
	6538:  uint16(2),
	6539:  uint16(3),
	6540:  uint16(1),
	6541:  uint16(sym_comment),
	6542:  uint16(633),
	6543:  uint16(8),
	6544:  uint16(anon_sym_SEMI),
	6545:  uint16(anon_sym_EQ),
	6546:  uint16(anon_sym_DOT),
	6547:  uint16(anon_sym_LPAREN),
	6548:  uint16(anon_sym_RPAREN),
	6549:  uint16(anon_sym_DOLLAR),
	6550:  uint16(anon_sym_COMMA),
	6551:  uint16(anon_sym_LBRACE),
	6552:  uint16(4),
	6553:  uint16(3),
	6554:  uint16(1),
	6555:  uint16(sym_comment),
	6556:  uint16(501),
	6557:  uint16(1),
	6558:  uint16(anon_sym_DOT),
	6559:  uint16(188),
	6560:  uint16(1),
	6561:  uint16(aux_sym_custom_type_repeat1),
	6562:  uint16(635),
	6563:  uint16(6),
	6564:  uint16(anon_sym_SEMI),
	6565:  uint16(anon_sym_EQ),
	6566:  uint16(anon_sym_RPAREN),
	6567:  uint16(anon_sym_DOLLAR),
	6568:  uint16(anon_sym_COMMA),
	6569:  uint16(anon_sym_RBRACK),
	6570:  uint16(4),
	6571:  uint16(3),
	6572:  uint16(1),
	6573:  uint16(sym_comment),
	6574:  uint16(501),
	6575:  uint16(1),
	6576:  uint16(anon_sym_DOT),
	6577:  uint16(189),
	6578:  uint16(1),
	6579:  uint16(aux_sym_custom_type_repeat1),
	6580:  uint16(635),
	6581:  uint16(6),
	6582:  uint16(anon_sym_SEMI),
	6583:  uint16(anon_sym_EQ),
	6584:  uint16(anon_sym_RPAREN),
	6585:  uint16(anon_sym_DOLLAR),
	6586:  uint16(anon_sym_COMMA),
	6587:  uint16(anon_sym_RBRACK),
	6588:  uint16(4),
	6589:  uint16(3),
	6590:  uint16(1),
	6591:  uint16(sym_comment),
	6592:  uint16(639),
	6593:  uint16(1),
	6594:  uint16(anon_sym_BQUOTE),
	6595:  uint16(194),
	6596:  uint16(1),
	6597:  uint16(aux_sym_block_text_repeat2),
	6598:  uint16(637),
	6599:  uint16(6),
	6600:  uint16(anon_sym_SEMI),
	6601:  uint16(anon_sym_RPAREN),
	6602:  uint16(anon_sym_DOLLAR),
	6603:  uint16(anon_sym_COMMA),
	6604:  uint16(anon_sym_RBRACK),
	6605:  uint16(sym_identifier),
	6606:  uint16(4),
	6607:  uint16(3),
	6608:  uint16(1),
	6609:  uint16(sym_comment),
	6610:  uint16(644),
	6611:  uint16(1),
	6612:  uint16(anon_sym_DOLLAR),
	6613:  uint16(195),
	6614:  uint16(1),
	6615:  uint16(aux_sym_annotation_repeat1),
	6616:  uint16(642),
	6617:  uint16(5),
	6618:  uint16(anon_sym_SEMI),
	6619:  uint16(anon_sym_EQ),
	6620:  uint16(anon_sym_RPAREN),
	6621:  uint16(anon_sym_COMMA),
	6622:  uint16(anon_sym_LBRACE),
	6623:  uint16(4),
	6624:  uint16(3),
	6625:  uint16(1),
	6626:  uint16(sym_comment),
	6627:  uint16(649),
	6628:  uint16(1),
	6629:  uint16(anon_sym_DOT),
	6630:  uint16(651),
	6631:  uint16(1),
	6632:  uint16(sym__identifier_no_period),
	6633:  uint16(647),
	6634:  uint16(5),
	6635:  uint16(anon_sym_SEMI),
	6636:  uint16(anon_sym_RPAREN),
	6637:  uint16(anon_sym_DOLLAR),
	6638:  uint16(anon_sym_COMMA),
	6639:  uint16(anon_sym_RBRACK),
	6640:  uint16(2),
	6641:  uint16(3),
	6642:  uint16(1),
	6643:  uint16(sym_comment),
	6644:  uint16(653),
	6645:  uint16(7),
	6646:  uint16(anon_sym_SEMI),
	6647:  uint16(anon_sym_EQ),
	6648:  uint16(anon_sym_DOT),
	6649:  uint16(anon_sym_RPAREN),
	6650:  uint16(anon_sym_DOLLAR),
	6651:  uint16(anon_sym_COMMA),
	6652:  uint16(anon_sym_RBRACK),
	6653:  uint16(8),
	6654:  uint16(3),
	6655:  uint16(1),
	6656:  uint16(sym_comment),
	6657:  uint16(503),
	6658:  uint16(1),
	6659:  uint16(anon_sym_LPAREN),
	6660:  uint16(655),
	6661:  uint16(1),
	6662:  uint16(sym_unique_id),
	6663:  uint16(657),
	6664:  uint16(1),
	6665:  uint16(anon_sym_DOLLAR),
	6666:  uint16(659),
	6667:  uint16(1),
	6668:  uint16(anon_sym_LBRACE),
	6669:  uint16(661),
	6670:  uint16(1),
	6671:  uint16(anon_sym_extends),
	6672:  uint16(295),
	6673:  uint16(1),
	6674:  uint16(sym_generics),
	6675:  uint16(361),
	6676:  uint16(1),
	6677:  uint16(aux_sym_annotation_repeat1),
	6678:  uint16(7),
	6679:  uint16(3),
	6680:  uint16(1),
	6681:  uint16(sym_comment),
	6682:  uint16(501),
	6683:  uint16(1),
	6684:  uint16(anon_sym_DOT),
	6685:  uint16(503),
	6686:  uint16(1),
	6687:  uint16(anon_sym_LPAREN),
	6688:  uint16(663),
	6689:  uint16(1),
	6690:  uint16(anon_sym_EQ),
	6691:  uint16(192),
	6692:  uint16(1),
	6693:  uint16(aux_sym_custom_type_repeat1),
	6694:  uint16(193),
	6695:  uint16(1),
	6696:  uint16(sym_generics),
	6697:  uint16(499),
	6698:  uint16(2),
	6699:  uint16(anon_sym_RPAREN),
	6700:  uint16(anon_sym_COMMA),
	6701:  uint16(4),
	6702:  uint16(3),
	6703:  uint16(1),
	6704:  uint16(sym_comment),
	6705:  uint16(649),
	6706:  uint16(1),
	6707:  uint16(anon_sym_DOT),
	6708:  uint16(651),
	6709:  uint16(1),
	6710:  uint16(sym__identifier_no_period),
	6711:  uint16(665),
	6712:  uint16(5),
	6713:  uint16(anon_sym_SEMI),
	6714:  uint16(anon_sym_RPAREN),
	6715:  uint16(anon_sym_DOLLAR),
	6716:  uint16(anon_sym_COMMA),
	6717:  uint16(anon_sym_RBRACK),
	6718:  uint16(2),
	6719:  uint16(3),
	6720:  uint16(1),
	6721:  uint16(sym_comment),
	6722:  uint16(667),
	6723:  uint16(6),
	6724:  uint16(anon_sym_SEMI),
	6725:  uint16(anon_sym_EQ),
	6726:  uint16(anon_sym_RPAREN),
	6727:  uint16(anon_sym_DOLLAR),
	6728:  uint16(anon_sym_COMMA),
	6729:  uint16(anon_sym_LBRACE),
	6730:  uint16(2),
	6731:  uint16(3),
	6732:  uint16(1),
	6733:  uint16(sym_comment),
	6734:  uint16(669),
	6735:  uint16(6),
	6736:  uint16(anon_sym_SEMI),
	6737:  uint16(anon_sym_RPAREN),
	6738:  uint16(anon_sym_DOLLAR),
	6739:  uint16(anon_sym_COMMA),
	6740:  uint16(anon_sym_RBRACK),
	6741:  uint16(sym_identifier),
	6742:  uint16(5),
	6743:  uint16(371),
	6744:  uint16(1),
	6745:  uint16(sym_comment),
	6746:  uint16(671),
	6747:  uint16(1),
	6748:  uint16(anon_sym_SQUOTE),
	6749:  uint16(673),
	6750:  uint16(1),
	6751:  uint16(sym_unescaped_single_string_fragment),
	6752:  uint16(675),
	6753:  uint16(2),
	6754:  uint16(aux_sym__escape_sequence_token1),
	6755:  uint16(sym_escape_sequence),
	6756:  uint16(236),
	6757:  uint16(2),
	6758:  uint16(sym__escape_sequence),
	6759:  uint16(aux_sym_string_repeat2),
	6760:  uint16(2),
	6761:  uint16(3),
	6762:  uint16(1),
	6763:  uint16(sym_comment),
	6764:  uint16(677),
	6765:  uint16(6),
	6766:  uint16(anon_sym_SEMI),
	6767:  uint16(anon_sym_RPAREN),
	6768:  uint16(anon_sym_DOLLAR),
	6769:  uint16(anon_sym_COMMA),
	6770:  uint16(anon_sym_RBRACK),
	6771:  uint16(sym_identifier),
	6772:  uint16(2),
	6773:  uint16(3),
	6774:  uint16(1),
	6775:  uint16(sym_comment),
	6776:  uint16(679),
	6777:  uint16(6),
	6778:  uint16(anon_sym_SEMI),
	6779:  uint16(anon_sym_EQ),
	6780:  uint16(anon_sym_RPAREN),
	6781:  uint16(anon_sym_DOLLAR),
	6782:  uint16(anon_sym_COMMA),
	6783:  uint16(anon_sym_LBRACE),
	6784:  uint16(2),
	6785:  uint16(3),
	6786:  uint16(1),
	6787:  uint16(sym_comment),
	6788:  uint16(681),
	6789:  uint16(6),
	6790:  uint16(anon_sym_SEMI),
	6791:  uint16(anon_sym_EQ),
	6792:  uint16(anon_sym_RPAREN),
	6793:  uint16(anon_sym_DOLLAR),
	6794:  uint16(anon_sym_COMMA),
	6795:  uint16(anon_sym_LBRACE),
	6796:  uint16(5),
	6797:  uint16(371),
	6798:  uint16(1),
	6799:  uint16(sym_comment),
	6800:  uint16(671),
	6801:  uint16(1),
	6802:  uint16(anon_sym_DQUOTE),
	6803:  uint16(683),
	6804:  uint16(1),
	6805:  uint16(sym_unescaped_double_string_fragment),
	6806:  uint16(685),
	6807:  uint16(2),
	6808:  uint16(aux_sym__escape_sequence_token1),
	6809:  uint16(sym_escape_sequence),
	6810:  uint16(220),
	6811:  uint16(2),
	6812:  uint16(sym__escape_sequence),
	6813:  uint16(aux_sym_string_repeat1),
	6814:  uint16(2),
	6815:  uint16(3),
	6816:  uint16(1),
	6817:  uint16(sym_comment),
	6818:  uint16(687),
	6819:  uint16(6),
	6820:  uint16(anon_sym_SEMI),
	6821:  uint16(anon_sym_LPAREN),
	6822:  uint16(anon_sym_DOLLAR),
	6823:  uint16(anon_sym_COLON),
	6824:  uint16(anon_sym_LBRACK),
	6825:  uint16(sym_identifier),
	6826:  uint16(2),
	6827:  uint16(3),
	6828:  uint16(1),
	6829:  uint16(sym_comment),
	6830:  uint16(689),
	6831:  uint16(6),
	6832:  uint16(anon_sym_SEMI),
	6833:  uint16(anon_sym_EQ),
	6834:  uint16(anon_sym_RPAREN),
	6835:  uint16(anon_sym_DOLLAR),
	6836:  uint16(anon_sym_COMMA),
	6837:  uint16(anon_sym_LBRACE),
	6838:  uint16(2),
	6839:  uint16(3),
	6840:  uint16(1),
	6841:  uint16(sym_comment),
	6842:  uint16(475),
	6843:  uint16(6),
	6844:  uint16(anon_sym_SEMI),
	6845:  uint16(anon_sym_RPAREN),
	6846:  uint16(anon_sym_DOLLAR),
	6847:  uint16(anon_sym_COMMA),
	6848:  uint16(anon_sym_RBRACK),
	6849:  uint16(sym_identifier),
	6850:  uint16(2),
	6851:  uint16(3),
	6852:  uint16(1),
	6853:  uint16(sym_comment),
	6854:  uint16(691),
	6855:  uint16(6),
	6856:  uint16(anon_sym_SEMI),
	6857:  uint16(anon_sym_RPAREN),
	6858:  uint16(anon_sym_DOLLAR),
	6859:  uint16(anon_sym_COMMA),
	6860:  uint16(anon_sym_RBRACK),
	6861:  uint16(sym_identifier),
	6862:  uint16(2),
	6863:  uint16(3),
	6864:  uint16(1),
	6865:  uint16(sym_comment),
	6866:  uint16(693),
	6867:  uint16(6),
	6868:  uint16(anon_sym_SEMI),
	6869:  uint16(anon_sym_RPAREN),
	6870:  uint16(anon_sym_DOLLAR),
	6871:  uint16(anon_sym_COMMA),
	6872:  uint16(anon_sym_RBRACK),
	6873:  uint16(sym_identifier),
	6874:  uint16(2),
	6875:  uint16(3),
	6876:  uint16(1),
	6877:  uint16(sym_comment),
	6878:  uint16(479),
	6879:  uint16(6),
	6880:  uint16(anon_sym_SEMI),
	6881:  uint16(anon_sym_EQ),
	6882:  uint16(anon_sym_RPAREN),
	6883:  uint16(anon_sym_DOLLAR),
	6884:  uint16(anon_sym_COMMA),
	6885:  uint16(anon_sym_LBRACE),
	6886:  uint16(2),
	6887:  uint16(3),
	6888:  uint16(1),
	6889:  uint16(sym_comment),
	6890:  uint16(695),
	6891:  uint16(6),
	6892:  uint16(anon_sym_SEMI),
	6893:  uint16(anon_sym_RPAREN),
	6894:  uint16(anon_sym_DOLLAR),
	6895:  uint16(anon_sym_COMMA),
	6896:  uint16(anon_sym_RBRACK),
	6897:  uint16(sym_identifier),
	6898:  uint16(2),
	6899:  uint16(3),
	6900:  uint16(1),
	6901:  uint16(sym_comment),
	6902:  uint16(697),
	6903:  uint16(6),
	6904:  uint16(anon_sym_SEMI),
	6905:  uint16(anon_sym_RPAREN),
	6906:  uint16(anon_sym_DOLLAR),
	6907:  uint16(anon_sym_COMMA),
	6908:  uint16(anon_sym_RBRACK),
	6909:  uint16(sym_identifier),
	6910:  uint16(7),
	6911:  uint16(3),
	6912:  uint16(1),
	6913:  uint16(sym_comment),
	6914:  uint16(503),
	6915:  uint16(1),
	6916:  uint16(anon_sym_LPAREN),
	6917:  uint16(657),
	6918:  uint16(1),
	6919:  uint16(anon_sym_DOLLAR),
	6920:  uint16(699),
	6921:  uint16(1),
	6922:  uint16(anon_sym_LBRACE),
	6923:  uint16(701),
	6924:  uint16(1),
	6925:  uint16(anon_sym_extends),
	6926:  uint16(291),
	6927:  uint16(1),
	6928:  uint16(sym_generics),
	6929:  uint16(412),
	6930:  uint16(1),
	6931:  uint16(aux_sym_annotation_repeat1),
	6932:  uint16(2),
	6933:  uint16(3),
	6934:  uint16(1),
	6935:  uint16(sym_comment),
	6936:  uint16(703),
	6937:  uint16(6),
	6938:  uint16(anon_sym_SEMI),
	6939:  uint16(anon_sym_EQ),
	6940:  uint16(anon_sym_RPAREN),
	6941:  uint16(anon_sym_DOLLAR),
	6942:  uint16(anon_sym_COMMA),
	6943:  uint16(anon_sym_LBRACE),
	6944:  uint16(7),
	6945:  uint16(3),
	6946:  uint16(1),
	6947:  uint16(sym_comment),
	6948:  uint16(503),
	6949:  uint16(1),
	6950:  uint16(anon_sym_LPAREN),
	6951:  uint16(657),
	6952:  uint16(1),
	6953:  uint16(anon_sym_DOLLAR),
	6954:  uint16(705),
	6955:  uint16(1),
	6956:  uint16(anon_sym_SEMI),
	6957:  uint16(707),
	6958:  uint16(1),
	6959:  uint16(anon_sym_DASH_GT),
	6960:  uint16(270),
	6961:  uint16(1),
	6962:  uint16(sym_generics),
	6963:  uint16(375),
	6964:  uint16(1),
	6965:  uint16(aux_sym_annotation_repeat1),
	6966:  uint16(2),
	6967:  uint16(3),
	6968:  uint16(1),
	6969:  uint16(sym_comment),
	6970:  uint16(709),
	6971:  uint16(6),
	6972:  uint16(anon_sym_SEMI),
	6973:  uint16(anon_sym_EQ),
	6974:  uint16(anon_sym_RPAREN),
	6975:  uint16(anon_sym_DOLLAR),
	6976:  uint16(anon_sym_COMMA),
	6977:  uint16(anon_sym_RBRACK),
	6978:  uint16(5),
	6979:  uint16(371),
	6980:  uint16(1),
	6981:  uint16(sym_comment),
	6982:  uint16(711),
	6983:  uint16(1),
	6984:  uint16(anon_sym_DQUOTE),
	6985:  uint16(713),
	6986:  uint16(1),
	6987:  uint16(sym_unescaped_double_string_fragment),
	6988:  uint16(716),
	6989:  uint16(2),
	6990:  uint16(aux_sym__escape_sequence_token1),
	6991:  uint16(sym_escape_sequence),
	6992:  uint16(220),
	6993:  uint16(2),
	6994:  uint16(sym__escape_sequence),
	6995:  uint16(aux_sym_string_repeat1),
	6996:  uint16(2),
	6997:  uint16(3),
	6998:  uint16(1),
	6999:  uint16(sym_comment),
	7000:  uint16(719),
	7001:  uint16(6),
	7002:  uint16(anon_sym_SEMI),
	7003:  uint16(anon_sym_EQ),
	7004:  uint16(anon_sym_RPAREN),
	7005:  uint16(anon_sym_DOLLAR),
	7006:  uint16(anon_sym_COMMA),
	7007:  uint16(anon_sym_LBRACE),
	7008:  uint16(5),
	7009:  uint16(371),
	7010:  uint16(1),
	7011:  uint16(sym_comment),
	7012:  uint16(721),
	7013:  uint16(1),
	7014:  uint16(anon_sym_DQUOTE),
	7015:  uint16(723),
	7016:  uint16(1),
	7017:  uint16(sym_unescaped_double_string_fragment),
	7018:  uint16(725),
	7019:  uint16(2),
	7020:  uint16(aux_sym__escape_sequence_token1),
	7021:  uint16(sym_escape_sequence),
	7022:  uint16(207),
	7023:  uint16(2),
	7024:  uint16(sym__escape_sequence),
	7025:  uint16(aux_sym_string_repeat1),
	7026:  uint16(5),
	7027:  uint16(371),
	7028:  uint16(1),
	7029:  uint16(sym_comment),
	7030:  uint16(721),
	7031:  uint16(1),
	7032:  uint16(anon_sym_SQUOTE),
	7033:  uint16(727),
	7034:  uint16(1),
	7035:  uint16(sym_unescaped_single_string_fragment),
	7036:  uint16(729),
	7037:  uint16(2),
	7038:  uint16(aux_sym__escape_sequence_token1),
	7039:  uint16(sym_escape_sequence),
	7040:  uint16(203),
	7041:  uint16(2),
	7042:  uint16(sym__escape_sequence),
	7043:  uint16(aux_sym_string_repeat2),
	7044:  uint16(2),
	7045:  uint16(3),
	7046:  uint16(1),
	7047:  uint16(sym_comment),
	7048:  uint16(495),
	7049:  uint16(6),
	7050:  uint16(anon_sym_SEMI),
	7051:  uint16(anon_sym_EQ),
	7052:  uint16(anon_sym_RPAREN),
	7053:  uint16(anon_sym_DOLLAR),
	7054:  uint16(anon_sym_COMMA),
	7055:  uint16(anon_sym_LBRACE),
	7056:  uint16(2),
	7057:  uint16(3),
	7058:  uint16(1),
	7059:  uint16(sym_comment),
	7060:  uint16(731),
	7061:  uint16(6),
	7062:  uint16(anon_sym_SEMI),
	7063:  uint16(anon_sym_RPAREN),
	7064:  uint16(anon_sym_DOLLAR),
	7065:  uint16(anon_sym_COMMA),
	7066:  uint16(anon_sym_RBRACK),
	7067:  uint16(sym_identifier),
	7068:  uint16(2),
	7069:  uint16(3),
	7070:  uint16(1),
	7071:  uint16(sym_comment),
	7072:  uint16(733),
	7073:  uint16(6),
	7074:  uint16(anon_sym_SEMI),
	7075:  uint16(anon_sym_EQ),
	7076:  uint16(anon_sym_RPAREN),
	7077:  uint16(anon_sym_DOLLAR),
	7078:  uint16(anon_sym_COMMA),
	7079:  uint16(anon_sym_LBRACE),
	7080:  uint16(7),
	7081:  uint16(3),
	7082:  uint16(1),
	7083:  uint16(sym_comment),
	7084:  uint16(503),
	7085:  uint16(1),
	7086:  uint16(anon_sym_LPAREN),
	7087:  uint16(657),
	7088:  uint16(1),
	7089:  uint16(anon_sym_DOLLAR),
	7090:  uint16(735),
	7091:  uint16(1),
	7092:  uint16(sym_unique_id),
	7093:  uint16(737),
	7094:  uint16(1),
	7095:  uint16(anon_sym_LBRACE),
	7096:  uint16(287),
	7097:  uint16(1),
	7098:  uint16(sym_generics),
	7099:  uint16(371),
	7100:  uint16(1),
	7101:  uint16(aux_sym_annotation_repeat1),
	7102:  uint16(2),
	7103:  uint16(3),
	7104:  uint16(1),
	7105:  uint16(sym_comment),
	7106:  uint16(491),
	7107:  uint16(6),
	7108:  uint16(anon_sym_SEMI),
	7109:  uint16(anon_sym_EQ),
	7110:  uint16(anon_sym_RPAREN),
	7111:  uint16(anon_sym_DOLLAR),
	7112:  uint16(anon_sym_COMMA),
	7113:  uint16(anon_sym_LBRACE),
	7114:  uint16(2),
	7115:  uint16(3),
	7116:  uint16(1),
	7117:  uint16(sym_comment),
	7118:  uint16(739),
	7119:  uint16(6),
	7120:  uint16(anon_sym_SEMI),
	7121:  uint16(anon_sym_EQ),
	7122:  uint16(anon_sym_RPAREN),
	7123:  uint16(anon_sym_DOLLAR),
	7124:  uint16(anon_sym_COMMA),
	7125:  uint16(anon_sym_LBRACE),
	7126:  uint16(2),
	7127:  uint16(3),
	7128:  uint16(1),
	7129:  uint16(sym_comment),
	7130:  uint16(741),
	7131:  uint16(6),
	7132:  uint16(anon_sym_SEMI),
	7133:  uint16(anon_sym_EQ),
	7134:  uint16(anon_sym_RPAREN),
	7135:  uint16(anon_sym_DOLLAR),
	7136:  uint16(anon_sym_COMMA),
	7137:  uint16(anon_sym_RBRACK),
	7138:  uint16(2),
	7139:  uint16(3),
	7140:  uint16(1),
	7141:  uint16(sym_comment),
	7142:  uint16(743),
	7143:  uint16(6),
	7144:  uint16(anon_sym_SEMI),
	7145:  uint16(anon_sym_RPAREN),
	7146:  uint16(anon_sym_DOLLAR),
	7147:  uint16(anon_sym_COMMA),
	7148:  uint16(anon_sym_RBRACK),
	7149:  uint16(sym_identifier),
	7150:  uint16(2),
	7151:  uint16(3),
	7152:  uint16(1),
	7153:  uint16(sym_comment),
	7154:  uint16(511),
	7155:  uint16(6),
	7156:  uint16(anon_sym_SEMI),
	7157:  uint16(anon_sym_EQ),
	7158:  uint16(anon_sym_RPAREN),
	7159:  uint16(anon_sym_DOLLAR),
	7160:  uint16(anon_sym_COMMA),
	7161:  uint16(anon_sym_LBRACE),
	7162:  uint16(2),
	7163:  uint16(3),
	7164:  uint16(1),
	7165:  uint16(sym_comment),
	7166:  uint16(745),
	7167:  uint16(6),
	7168:  uint16(anon_sym_SEMI),
	7169:  uint16(anon_sym_RPAREN),
	7170:  uint16(anon_sym_DOLLAR),
	7171:  uint16(anon_sym_COMMA),
	7172:  uint16(anon_sym_RBRACK),
	7173:  uint16(sym_identifier),
	7174:  uint16(7),
	7175:  uint16(3),
	7176:  uint16(1),
	7177:  uint16(sym_comment),
	7178:  uint16(503),
	7179:  uint16(1),
	7180:  uint16(anon_sym_LPAREN),
	7181:  uint16(657),
	7182:  uint16(1),
	7183:  uint16(anon_sym_DOLLAR),
	7184:  uint16(747),
	7185:  uint16(1),
	7186:  uint16(anon_sym_SEMI),
	7187:  uint16(749),
	7188:  uint16(1),
	7189:  uint16(anon_sym_DASH_GT),
	7190:  uint16(299),
	7191:  uint16(1),
	7192:  uint16(sym_generics),
	7193:  uint16(314),
	7194:  uint16(1),
	7195:  uint16(aux_sym_annotation_repeat1),
	7196:  uint16(2),
	7197:  uint16(3),
	7198:  uint16(1),
	7199:  uint16(sym_comment),
	7200:  uint16(751),
	7201:  uint16(6),
	7202:  uint16(anon_sym_SEMI),
	7203:  uint16(anon_sym_EQ),
	7204:  uint16(anon_sym_RPAREN),
	7205:  uint16(anon_sym_DOLLAR),
	7206:  uint16(anon_sym_COMMA),
	7207:  uint16(anon_sym_LBRACE),
	7208:  uint16(5),
	7209:  uint16(371),
	7210:  uint16(1),
	7211:  uint16(sym_comment),
	7212:  uint16(753),
	7213:  uint16(1),
	7214:  uint16(anon_sym_SQUOTE),
	7215:  uint16(755),
	7216:  uint16(1),
	7217:  uint16(sym_unescaped_single_string_fragment),
	7218:  uint16(758),
	7219:  uint16(2),
	7220:  uint16(aux_sym__escape_sequence_token1),
	7221:  uint16(sym_escape_sequence),
	7222:  uint16(236),
	7223:  uint16(2),
	7224:  uint16(sym__escape_sequence),
	7225:  uint16(aux_sym_string_repeat2),
	7226:  uint16(7),
	7227:  uint16(3),
	7228:  uint16(1),
	7229:  uint16(sym_comment),
	7230:  uint16(503),
	7231:  uint16(1),
	7232:  uint16(anon_sym_LPAREN),
	7233:  uint16(761),
	7234:  uint16(1),
	7235:  uint16(anon_sym_EQ),
	7236:  uint16(763),
	7237:  uint16(1),
	7238:  uint16(anon_sym_RPAREN),
	7239:  uint16(765),
	7240:  uint16(1),
	7241:  uint16(anon_sym_COMMA),
	7242:  uint16(284),
	7243:  uint16(1),
	7244:  uint16(sym_generics),
	7245:  uint16(336),
	7246:  uint16(1),
	7247:  uint16(aux_sym_named_return_type_repeat1),
	7248:  uint16(2),
	7249:  uint16(3),
	7250:  uint16(1),
	7251:  uint16(sym_comment),
	7252:  uint16(767),
	7253:  uint16(6),
	7254:  uint16(anon_sym_SEMI),
	7255:  uint16(anon_sym_EQ),
	7256:  uint16(anon_sym_RPAREN),
	7257:  uint16(anon_sym_DOLLAR),
	7258:  uint16(anon_sym_COMMA),
	7259:  uint16(anon_sym_RBRACK),
	7260:  uint16(6),
	7261:  uint16(3),
	7262:  uint16(1),
	7263:  uint16(sym_comment),
	7264:  uint16(769),
	7265:  uint16(1),
	7266:  uint16(sym_identifier),
	7267:  uint16(771),
	7268:  uint16(1),
	7269:  uint16(anon_sym_LPAREN),
	7270:  uint16(272),
	7271:  uint16(1),
	7272:  uint16(sym_unnamed_return_type),
	7273:  uint16(370),
	7274:  uint16(1),
	7275:  uint16(sym_return_type),
	7276:  uint16(465),
	7277:  uint16(1),
	7278:  uint16(sym_named_return_types),
	7279:  uint16(5),
	7280:  uint16(3),
	7281:  uint16(1),
	7282:  uint16(sym_comment),
	7283:  uint16(503),
	7284:  uint16(1),
	7285:  uint16(anon_sym_LPAREN),
	7286:  uint16(773),
	7287:  uint16(1),
	7288:  uint16(anon_sym_EQ),
	7289:  uint16(311),
	7290:  uint16(1),
	7291:  uint16(sym_generics),
	7292:  uint16(775),
	7293:  uint16(2),
	7294:  uint16(anon_sym_RPAREN),
	7295:  uint16(anon_sym_COMMA),
	7296:  uint16(3),
	7297:  uint16(3),
	7298:  uint16(1),
	7299:  uint16(sym_comment),
	7300:  uint16(779),
	7301:  uint16(1),
	7302:  uint16(anon_sym_RBRACE),
	7303:  uint16(777),
	7304:  uint16(4),
	7305:  uint16(anon_sym_enum),
	7306:  uint16(anon_sym_interface),
	7307:  uint16(anon_sym_struct),
	7308:  uint16(sym_identifier),
	7309:  uint16(3),
	7310:  uint16(3),
	7311:  uint16(1),
	7312:  uint16(sym_comment),
	7313:  uint16(783),
	7314:  uint16(1),
	7315:  uint16(anon_sym_RBRACE),
	7316:  uint16(781),
	7317:  uint16(4),
	7318:  uint16(anon_sym_enum),
	7319:  uint16(anon_sym_interface),
	7320:  uint16(anon_sym_struct),
	7321:  uint16(sym_identifier),
	7322:  uint16(4),
	7323:  uint16(3),
	7324:  uint16(1),
	7325:  uint16(sym_comment),
	7326:  uint16(649),
	7327:  uint16(1),
	7328:  uint16(anon_sym_DOT),
	7329:  uint16(785),
	7330:  uint16(1),
	7331:  uint16(sym__identifier_no_period),
	7332:  uint16(647),
	7333:  uint16(3),
	7334:  uint16(anon_sym_RPAREN),
	7335:  uint16(anon_sym_COMMA),
	7336:  uint16(sym_identifier),
	7337:  uint16(4),
	7338:  uint16(3),
	7339:  uint16(1),
	7340:  uint16(sym_comment),
	7341:  uint16(649),
	7342:  uint16(1),
	7343:  uint16(anon_sym_DOT),
	7344:  uint16(785),
	7345:  uint16(1),
	7346:  uint16(sym__identifier_no_period),
	7347:  uint16(665),
	7348:  uint16(3),
	7349:  uint16(anon_sym_RPAREN),
	7350:  uint16(anon_sym_COMMA),
	7351:  uint16(sym_identifier),
	7352:  uint16(3),
	7353:  uint16(3),
	7354:  uint16(1),
	7355:  uint16(sym_comment),
	7356:  uint16(789),
	7357:  uint16(1),
	7358:  uint16(anon_sym_RBRACE),
	7359:  uint16(787),
	7360:  uint16(4),
	7361:  uint16(anon_sym_enum),
	7362:  uint16(anon_sym_interface),
	7363:  uint16(anon_sym_struct),
	7364:  uint16(sym_identifier),
	7365:  uint16(3),
	7366:  uint16(3),
	7367:  uint16(1),
	7368:  uint16(sym_comment),
	7369:  uint16(793),
	7370:  uint16(1),
	7371:  uint16(anon_sym_RBRACE),
	7372:  uint16(791),
	7373:  uint16(4),
	7374:  uint16(anon_sym_enum),
	7375:  uint16(anon_sym_interface),
	7376:  uint16(anon_sym_struct),
	7377:  uint16(sym_identifier),
	7378:  uint16(3),
	7379:  uint16(3),
	7380:  uint16(1),
	7381:  uint16(sym_comment),
	7382:  uint16(797),
	7383:  uint16(1),
	7384:  uint16(anon_sym_RBRACE),
	7385:  uint16(795),
	7386:  uint16(4),
	7387:  uint16(anon_sym_enum),
	7388:  uint16(anon_sym_interface),
	7389:  uint16(anon_sym_struct),
	7390:  uint16(sym_identifier),
	7391:  uint16(3),
	7392:  uint16(3),
	7393:  uint16(1),
	7394:  uint16(sym_comment),
	7395:  uint16(801),
	7396:  uint16(1),
	7397:  uint16(anon_sym_RBRACE),
	7398:  uint16(799),
	7399:  uint16(4),
	7400:  uint16(anon_sym_enum),
	7401:  uint16(anon_sym_interface),
	7402:  uint16(anon_sym_struct),
	7403:  uint16(sym_identifier),
	7404:  uint16(3),
	7405:  uint16(3),
	7406:  uint16(1),
	7407:  uint16(sym_comment),
	7408:  uint16(805),
	7409:  uint16(1),
	7410:  uint16(anon_sym_RBRACE),
	7411:  uint16(803),
	7412:  uint16(4),
	7413:  uint16(anon_sym_enum),
	7414:  uint16(anon_sym_interface),
	7415:  uint16(anon_sym_struct),
	7416:  uint16(sym_identifier),
	7417:  uint16(6),
	7418:  uint16(3),
	7419:  uint16(1),
	7420:  uint16(sym_comment),
	7421:  uint16(807),
	7422:  uint16(1),
	7423:  uint16(anon_sym_LBRACE),
	7424:  uint16(809),
	7425:  uint16(1),
	7426:  uint16(sym__normal_version),
	7427:  uint16(811),
	7428:  uint16(1),
	7429:  uint16(aux_sym__inline_version_token1),
	7430:  uint16(208),
	7431:  uint16(1),
	7432:  uint16(sym__inline_version),
	7433:  uint16(539),
	7434:  uint16(1),
	7435:  uint16(sym_field_version),
	7436:  uint16(6),
	7437:  uint16(3),
	7438:  uint16(1),
	7439:  uint16(sym_comment),
	7440:  uint16(813),
	7441:  uint16(1),
	7442:  uint16(sym_identifier),
	7443:  uint16(815),
	7444:  uint16(1),
	7445:  uint16(anon_sym_LPAREN),
	7446:  uint16(817),
	7447:  uint16(1),
	7448:  uint16(anon_sym_LBRACK),
	7449:  uint16(278),
	7450:  uint16(1),
	7451:  uint16(sym_method_parameters),
	7452:  uint16(382),
	7453:  uint16(1),
	7454:  uint16(sym_implicit_generics),
	7455:  uint16(6),
	7456:  uint16(3),
	7457:  uint16(1),
	7458:  uint16(sym_comment),
	7459:  uint16(809),
	7460:  uint16(1),
	7461:  uint16(sym__normal_version),
	7462:  uint16(811),
	7463:  uint16(1),
	7464:  uint16(aux_sym__inline_version_token1),
	7465:  uint16(819),
	7466:  uint16(1),
	7467:  uint16(anon_sym_COLON),
	7468:  uint16(208),
	7469:  uint16(1),
	7470:  uint16(sym__inline_version),
	7471:  uint16(538),
	7472:  uint16(1),
	7473:  uint16(sym_field_version),
	7474:  uint16(6),
	7475:  uint16(3),
	7476:  uint16(1),
	7477:  uint16(sym_comment),
	7478:  uint16(769),
	7479:  uint16(1),
	7480:  uint16(sym_identifier),
	7481:  uint16(771),
	7482:  uint16(1),
	7483:  uint16(anon_sym_LPAREN),
	7484:  uint16(272),
	7485:  uint16(1),
	7486:  uint16(sym_unnamed_return_type),
	7487:  uint16(309),
	7488:  uint16(1),
	7489:  uint16(sym_return_type),
	7490:  uint16(465),
	7491:  uint16(1),
	7492:  uint16(sym_named_return_types),
	7493:  uint16(3),
	7494:  uint16(3),
	7495:  uint16(1),
	7496:  uint16(sym_comment),
	7497:  uint16(823),
	7498:  uint16(1),
	7499:  uint16(anon_sym_RBRACE),
	7500:  uint16(821),
	7501:  uint16(4),
	7502:  uint16(anon_sym_enum),
	7503:  uint16(anon_sym_interface),
	7504:  uint16(anon_sym_struct),
	7505:  uint16(sym_identifier),
	7506:  uint16(3),
	7507:  uint16(3),
	7508:  uint16(1),
	7509:  uint16(sym_comment),
	7510:  uint16(827),
	7511:  uint16(1),
	7512:  uint16(anon_sym_RBRACE),
	7513:  uint16(825),
	7514:  uint16(4),
	7515:  uint16(anon_sym_enum),
	7516:  uint16(anon_sym_interface),
	7517:  uint16(anon_sym_struct),
	7518:  uint16(sym_identifier),
	7519:  uint16(6),
	7520:  uint16(3),
	7521:  uint16(1),
	7522:  uint16(sym_comment),
	7523:  uint16(769),
	7524:  uint16(1),
	7525:  uint16(sym_identifier),
	7526:  uint16(771),
	7527:  uint16(1),
	7528:  uint16(anon_sym_LPAREN),
	7529:  uint16(272),
	7530:  uint16(1),
	7531:  uint16(sym_unnamed_return_type),
	7532:  uint16(427),
	7533:  uint16(1),
	7534:  uint16(sym_return_type),
	7535:  uint16(465),
	7536:  uint16(1),
	7537:  uint16(sym_named_return_types),
	7538:  uint16(3),
	7539:  uint16(3),
	7540:  uint16(1),
	7541:  uint16(sym_comment),
	7542:  uint16(831),
	7543:  uint16(1),
	7544:  uint16(anon_sym_RBRACE),
	7545:  uint16(829),
	7546:  uint16(4),
	7547:  uint16(anon_sym_enum),
	7548:  uint16(anon_sym_interface),
	7549:  uint16(anon_sym_struct),
	7550:  uint16(sym_identifier),
	7551:  uint16(3),
	7552:  uint16(3),
	7553:  uint16(1),
	7554:  uint16(sym_comment),
	7555:  uint16(835),
	7556:  uint16(1),
	7557:  uint16(anon_sym_RBRACE),
	7558:  uint16(833),
	7559:  uint16(4),
	7560:  uint16(anon_sym_enum),
	7561:  uint16(anon_sym_interface),
	7562:  uint16(anon_sym_struct),
	7563:  uint16(sym_identifier),
	7564:  uint16(5),
	7565:  uint16(3),
	7566:  uint16(1),
	7567:  uint16(sym_comment),
	7568:  uint16(657),
	7569:  uint16(1),
	7570:  uint16(anon_sym_DOLLAR),
	7571:  uint16(837),
	7572:  uint16(1),
	7573:  uint16(anon_sym_EQ),
	7574:  uint16(195),
	7575:  uint16(1),
	7576:  uint16(aux_sym_annotation_repeat1),
	7577:  uint16(839),
	7578:  uint16(2),
	7579:  uint16(anon_sym_RPAREN),
	7580:  uint16(anon_sym_COMMA),
	7581:  uint16(3),
	7582:  uint16(3),
	7583:  uint16(1),
	7584:  uint16(sym_comment),
	7585:  uint16(843),
	7586:  uint16(1),
	7587:  uint16(anon_sym_RBRACE),
	7588:  uint16(841),
	7589:  uint16(4),
	7590:  uint16(anon_sym_enum),
	7591:  uint16(anon_sym_interface),
	7592:  uint16(anon_sym_struct),
	7593:  uint16(sym_identifier),
	7594:  uint16(6),
	7595:  uint16(3),
	7596:  uint16(1),
	7597:  uint16(sym_comment),
	7598:  uint16(769),
	7599:  uint16(1),
	7600:  uint16(sym_identifier),
	7601:  uint16(771),
	7602:  uint16(1),
	7603:  uint16(anon_sym_LPAREN),
	7604:  uint16(272),
	7605:  uint16(1),
	7606:  uint16(sym_unnamed_return_type),
	7607:  uint16(424),
	7608:  uint16(1),
	7609:  uint16(sym_return_type),
	7610:  uint16(465),
	7611:  uint16(1),
	7612:  uint16(sym_named_return_types),
	7613:  uint16(6),
	7614:  uint16(3),
	7615:  uint16(1),
	7616:  uint16(sym_comment),
	7617:  uint16(809),
	7618:  uint16(1),
	7619:  uint16(sym__normal_version),
	7620:  uint16(811),
	7621:  uint16(1),
	7622:  uint16(aux_sym__inline_version_token1),
	7623:  uint16(819),
	7624:  uint16(1),
	7625:  uint16(anon_sym_COLON),
	7626:  uint16(208),
	7627:  uint16(1),
	7628:  uint16(sym__inline_version),
	7629:  uint16(495),
	7630:  uint16(1),
	7631:  uint16(sym_field_version),
	7632:  uint16(3),
	7633:  uint16(3),
	7634:  uint16(1),
	7635:  uint16(sym_comment),
	7636:  uint16(847),
	7637:  uint16(1),
	7638:  uint16(anon_sym_RBRACE),
	7639:  uint16(845),
	7640:  uint16(4),
	7641:  uint16(anon_sym_enum),
	7642:  uint16(anon_sym_interface),
	7643:  uint16(anon_sym_struct),
	7644:  uint16(sym_identifier),
	7645:  uint16(6),
	7646:  uint16(3),
	7647:  uint16(1),
	7648:  uint16(sym_comment),
	7649:  uint16(769),
	7650:  uint16(1),
	7651:  uint16(sym_identifier),
	7652:  uint16(771),
	7653:  uint16(1),
	7654:  uint16(anon_sym_LPAREN),
	7655:  uint16(272),
	7656:  uint16(1),
	7657:  uint16(sym_unnamed_return_type),
	7658:  uint16(364),
	7659:  uint16(1),
	7660:  uint16(sym_return_type),
	7661:  uint16(465),
	7662:  uint16(1),
	7663:  uint16(sym_named_return_types),
	7664:  uint16(6),
	7665:  uint16(3),
	7666:  uint16(1),
	7667:  uint16(sym_comment),
	7668:  uint16(769),
	7669:  uint16(1),
	7670:  uint16(sym_identifier),
	7671:  uint16(771),
	7672:  uint16(1),
	7673:  uint16(anon_sym_LPAREN),
	7674:  uint16(272),
	7675:  uint16(1),
	7676:  uint16(sym_unnamed_return_type),
	7677:  uint16(416),
	7678:  uint16(1),
	7679:  uint16(sym_return_type),
	7680:  uint16(465),
	7681:  uint16(1),
	7682:  uint16(sym_named_return_types),
	7683:  uint16(5),
	7684:  uint16(3),
	7685:  uint16(1),
	7686:  uint16(sym_comment),
	7687:  uint16(45),
	7688:  uint16(1),
	7689:  uint16(anon_sym_DQUOTE),
	7690:  uint16(47),
	7691:  uint16(1),
	7692:  uint16(anon_sym_SQUOTE),
	7693:  uint16(849),
	7694:  uint16(1),
	7695:  uint16(anon_sym_RPAREN),
	7696:  uint16(146),
	7697:  uint16(2),
	7698:  uint16(sym_string),
	7699:  uint16(aux_sym_concatenated_string_repeat1),
	7700:  uint16(3),
	7701:  uint16(3),
	7702:  uint16(1),
	7703:  uint16(sym_comment),
	7704:  uint16(853),
	7705:  uint16(1),
	7706:  uint16(anon_sym_RBRACE),
	7707:  uint16(851),
	7708:  uint16(4),
	7709:  uint16(anon_sym_enum),
	7710:  uint16(anon_sym_interface),
	7711:  uint16(anon_sym_struct),
	7712:  uint16(sym_identifier),
	7713:  uint16(5),
	7714:  uint16(3),
	7715:  uint16(1),
	7716:  uint16(sym_comment),
	7717:  uint16(657),
	7718:  uint16(1),
	7719:  uint16(anon_sym_DOLLAR),
	7720:  uint16(855),
	7721:  uint16(1),
	7722:  uint16(anon_sym_EQ),
	7723:  uint16(259),
	7724:  uint16(1),
	7725:  uint16(aux_sym_annotation_repeat1),
	7726:  uint16(857),
	7727:  uint16(2),
	7728:  uint16(anon_sym_RPAREN),
	7729:  uint16(anon_sym_COMMA),
	7730:  uint16(3),
	7731:  uint16(3),
	7732:  uint16(1),
	7733:  uint16(sym_comment),
	7734:  uint16(861),
	7735:  uint16(1),
	7736:  uint16(anon_sym_RBRACE),
	7737:  uint16(859),
	7738:  uint16(4),
	7739:  uint16(anon_sym_enum),
	7740:  uint16(anon_sym_interface),
	7741:  uint16(anon_sym_struct),
	7742:  uint16(sym_identifier),
	7743:  uint16(5),
	7744:  uint16(3),
	7745:  uint16(1),
	7746:  uint16(sym_comment),
	7747:  uint16(657),
	7748:  uint16(1),
	7749:  uint16(anon_sym_DOLLAR),
	7750:  uint16(863),
	7751:  uint16(1),
	7752:  uint16(anon_sym_SEMI),
	7753:  uint16(865),
	7754:  uint16(1),
	7755:  uint16(anon_sym_DASH_GT),
	7756:  uint16(430),
	7757:  uint16(1),
	7758:  uint16(aux_sym_annotation_repeat1),
	7759:  uint16(4),
	7760:  uint16(3),
	7761:  uint16(1),
	7762:  uint16(sym_comment),
	7763:  uint16(869),
	7764:  uint16(1),
	7765:  uint16(anon_sym_COMMA),
	7766:  uint16(271),
	7767:  uint16(1),
	7768:  uint16(aux_sym__annotation_array_def_repeat2),
	7769:  uint16(867),
	7770:  uint16(2),
	7771:  uint16(anon_sym_RPAREN),
	7772:  uint16(anon_sym_RBRACK),
	7773:  uint16(4),
	7774:  uint16(3),
	7775:  uint16(1),
	7776:  uint16(sym_comment),
	7777:  uint16(503),
	7778:  uint16(1),
	7779:  uint16(anon_sym_LPAREN),
	7780:  uint16(460),
	7781:  uint16(1),
	7782:  uint16(sym_generics),
	7783:  uint16(872),
	7784:  uint16(2),
	7785:  uint16(anon_sym_SEMI),
	7786:  uint16(anon_sym_DOLLAR),
	7787:  uint16(4),
	7788:  uint16(3),
	7789:  uint16(1),
	7790:  uint16(sym_comment),
	7791:  uint16(876),
	7792:  uint16(1),
	7793:  uint16(anon_sym_COMMA),
	7794:  uint16(282),
	7795:  uint16(1),
	7796:  uint16(aux_sym_generic_parameters_repeat1),
	7797:  uint16(874),
	7798:  uint16(2),
	7799:  uint16(anon_sym_RPAREN),
	7800:  uint16(anon_sym_RBRACK),
	7801:  uint16(4),
	7802:  uint16(3),
	7803:  uint16(1),
	7804:  uint16(sym_comment),
	7805:  uint16(878),
	7806:  uint16(1),
	7807:  uint16(sym_identifier),
	7808:  uint16(881),
	7809:  uint16(1),
	7810:  uint16(anon_sym_RBRACE),
	7811:  uint16(274),
	7812:  uint16(2),
	7813:  uint16(sym_enum_field),
	7814:  uint16(aux_sym_enum_repeat1),
	7815:  uint16(4),
	7816:  uint16(3),
	7817:  uint16(1),
	7818:  uint16(sym_comment),
	7819:  uint16(885),
	7820:  uint16(1),
	7821:  uint16(anon_sym_COMMA),
	7822:  uint16(275),
	7823:  uint16(1),
	7824:  uint16(aux_sym__annotation_array_def_repeat1),
	7825:  uint16(883),
	7826:  uint16(2),
	7827:  uint16(anon_sym_RPAREN),
	7828:  uint16(anon_sym_RBRACK),
	7829:  uint16(5),
	7830:  uint16(3),
	7831:  uint16(1),
	7832:  uint16(sym_comment),
	7833:  uint16(657),
	7834:  uint16(1),
	7835:  uint16(anon_sym_DOLLAR),
	7836:  uint16(888),
	7837:  uint16(1),
	7838:  uint16(anon_sym_SEMI),
	7839:  uint16(890),
	7840:  uint16(1),
	7841:  uint16(anon_sym_COLON),
	7842:  uint16(437),
	7843:  uint16(1),
	7844:  uint16(aux_sym_annotation_repeat1),
	7845:  uint16(4),
	7846:  uint16(3),
	7847:  uint16(1),
	7848:  uint16(sym_comment),
	7849:  uint16(892),
	7850:  uint16(1),
	7851:  uint16(sym_identifier),
	7852:  uint16(894),
	7853:  uint16(1),
	7854:  uint16(anon_sym_RBRACE),
	7855:  uint16(274),
	7856:  uint16(2),
	7857:  uint16(sym_enum_field),
	7858:  uint16(aux_sym_enum_repeat1),
	7859:  uint16(5),
	7860:  uint16(3),
	7861:  uint16(1),
	7862:  uint16(sym_comment),
	7863:  uint16(657),
	7864:  uint16(1),
	7865:  uint16(anon_sym_DOLLAR),
	7866:  uint16(896),
	7867:  uint16(1),
	7868:  uint16(anon_sym_SEMI),
	7869:  uint16(898),
	7870:  uint16(1),
	7871:  uint16(anon_sym_DASH_GT),
	7872:  uint16(319),
	7873:  uint16(1),
	7874:  uint16(aux_sym_annotation_repeat1),
	7875:  uint16(5),
	7876:  uint16(3),
	7877:  uint16(1),
	7878:  uint16(sym_comment),
	7879:  uint16(900),
	7880:  uint16(1),
	7881:  uint16(sym_identifier),
	7882:  uint16(902),
	7883:  uint16(1),
	7884:  uint16(anon_sym_RPAREN),
	7885:  uint16(317),
	7886:  uint16(1),
	7887:  uint16(sym_parameter),
	7888:  uint16(541),
	7889:  uint16(1),
	7890:  uint16(sym_parameters),
	7891:  uint16(4),
	7892:  uint16(3),
	7893:  uint16(1),
	7894:  uint16(sym_comment),
	7895:  uint16(904),
	7896:  uint16(1),
	7897:  uint16(sym_identifier),
	7898:  uint16(906),
	7899:  uint16(1),
	7900:  uint16(anon_sym_import),
	7901:  uint16(530),
	7902:  uint16(2),
	7903:  uint16(sym_replace_using),
	7904:  uint16(sym_import_using),
	7905:  uint16(4),
	7906:  uint16(3),
	7907:  uint16(1),
	7908:  uint16(sym_comment),
	7909:  uint16(910),
	7910:  uint16(1),
	7911:  uint16(anon_sym_COMMA),
	7912:  uint16(288),
	7913:  uint16(1),
	7914:  uint16(aux_sym__annotation_array_def_repeat2),
	7915:  uint16(908),
	7916:  uint16(2),
	7917:  uint16(anon_sym_RPAREN),
	7918:  uint16(anon_sym_RBRACK),
	7919:  uint16(4),
	7920:  uint16(3),
	7921:  uint16(1),
	7922:  uint16(sym_comment),
	7923:  uint16(914),
	7924:  uint16(1),
	7925:  uint16(anon_sym_COMMA),
	7926:  uint16(282),
	7927:  uint16(1),
	7928:  uint16(aux_sym_generic_parameters_repeat1),
	7929:  uint16(912),
	7930:  uint16(2),
	7931:  uint16(anon_sym_RPAREN),
	7932:  uint16(anon_sym_RBRACK),
	7933:  uint16(5),
	7934:  uint16(3),
	7935:  uint16(1),
	7936:  uint16(sym_comment),
	7937:  uint16(657),
	7938:  uint16(1),
	7939:  uint16(anon_sym_DOLLAR),
	7940:  uint16(917),
	7941:  uint16(1),
	7942:  uint16(anon_sym_SEMI),
	7943:  uint16(919),
	7944:  uint16(1),
	7945:  uint16(anon_sym_EQ),
	7946:  uint16(195),
	7947:  uint16(1),
	7948:  uint16(aux_sym_annotation_repeat1),
	7949:  uint16(5),
	7950:  uint16(3),
	7951:  uint16(1),
	7952:  uint16(sym_comment),
	7953:  uint16(765),
	7954:  uint16(1),
	7955:  uint16(anon_sym_COMMA),
	7956:  uint16(921),
	7957:  uint16(1),
	7958:  uint16(anon_sym_EQ),
	7959:  uint16(923),
	7960:  uint16(1),
	7961:  uint16(anon_sym_RPAREN),
	7962:  uint16(331),
	7963:  uint16(1),
	7964:  uint16(aux_sym_named_return_type_repeat1),
	7965:  uint16(4),
	7966:  uint16(3),
	7967:  uint16(1),
	7968:  uint16(sym_comment),
	7969:  uint16(892),
	7970:  uint16(1),
	7971:  uint16(sym_identifier),
	7972:  uint16(925),
	7973:  uint16(1),
	7974:  uint16(anon_sym_RBRACE),
	7975:  uint16(277),
	7976:  uint16(2),
	7977:  uint16(sym_enum_field),
	7978:  uint16(aux_sym_enum_repeat1),
	7979:  uint16(5),
	7980:  uint16(3),
	7981:  uint16(1),
	7982:  uint16(sym_comment),
	7983:  uint16(809),
	7984:  uint16(1),
	7985:  uint16(sym__normal_version),
	7986:  uint16(811),
	7987:  uint16(1),
	7988:  uint16(aux_sym__inline_version_token1),
	7989:  uint16(208),
	7990:  uint16(1),
	7991:  uint16(sym__inline_version),
	7992:  uint16(435),
	7993:  uint16(1),
	7994:  uint16(sym_field_version),
	7995:  uint16(5),
	7996:  uint16(3),
	7997:  uint16(1),
	7998:  uint16(sym_comment),
	7999:  uint16(657),
	8000:  uint16(1),
	8001:  uint16(anon_sym_DOLLAR),
	8002:  uint16(927),
	8003:  uint16(1),
	8004:  uint16(sym_unique_id),
	8005:  uint16(929),
	8006:  uint16(1),
	8007:  uint16(anon_sym_LBRACE),
	8008:  uint16(420),
	8009:  uint16(1),
	8010:  uint16(aux_sym_annotation_repeat1),
	8011:  uint16(3),
	8012:  uint16(3),
	8013:  uint16(1),
	8014:  uint16(sym_comment),
	8015:  uint16(271),
	8016:  uint16(1),
	8017:  uint16(aux_sym__annotation_array_def_repeat2),
	8018:  uint16(79),
	8019:  uint16(3),
	8020:  uint16(anon_sym_RPAREN),
	8021:  uint16(anon_sym_COMMA),
	8022:  uint16(anon_sym_RBRACK),
	8023:  uint16(4),
	8024:  uint16(3),
	8025:  uint16(1),
	8026:  uint16(sym_comment),
	8027:  uint16(892),
	8028:  uint16(1),
	8029:  uint16(sym_identifier),
	8030:  uint16(925),
	8031:  uint16(1),
	8032:  uint16(anon_sym_RBRACE),
	8033:  uint16(274),
	8034:  uint16(2),
	8035:  uint16(sym_enum_field),
	8036:  uint16(aux_sym_enum_repeat1),
	8037:  uint16(3),
	8038:  uint16(3),
	8039:  uint16(1),
	8040:  uint16(sym_comment),
	8041:  uint16(275),
	8042:  uint16(1),
	8043:  uint16(aux_sym__annotation_array_def_repeat1),
	8044:  uint16(931),
	8045:  uint16(3),
	8046:  uint16(anon_sym_RPAREN),
	8047:  uint16(anon_sym_COMMA),
	8048:  uint16(anon_sym_RBRACK),
	8049:  uint16(5),
	8050:  uint16(3),
	8051:  uint16(1),
	8052:  uint16(sym_comment),
	8053:  uint16(657),
	8054:  uint16(1),
	8055:  uint16(anon_sym_DOLLAR),
	8056:  uint16(933),
	8057:  uint16(1),
	8058:  uint16(anon_sym_LBRACE),
	8059:  uint16(935),
	8060:  uint16(1),
	8061:  uint16(anon_sym_extends),
	8062:  uint16(432),
	8063:  uint16(1),
	8064:  uint16(aux_sym_annotation_repeat1),
	8065:  uint16(5),
	8066:  uint16(3),
	8067:  uint16(1),
	8068:  uint16(sym_comment),
	8069:  uint16(45),
	8070:  uint16(1),
	8071:  uint16(anon_sym_DQUOTE),
	8072:  uint16(47),
	8073:  uint16(1),
	8074:  uint16(anon_sym_SQUOTE),
	8075:  uint16(937),
	8076:  uint16(1),
	8077:  uint16(anon_sym_LPAREN),
	8078:  uint16(453),
	8079:  uint16(1),
	8080:  uint16(sym_string),
	8081:  uint16(5),
	8082:  uint16(3),
	8083:  uint16(1),
	8084:  uint16(sym_comment),
	8085:  uint16(657),
	8086:  uint16(1),
	8087:  uint16(anon_sym_DOLLAR),
	8088:  uint16(939),
	8089:  uint16(1),
	8090:  uint16(anon_sym_SEMI),
	8091:  uint16(941),
	8092:  uint16(1),
	8093:  uint16(anon_sym_EQ),
	8094:  uint16(394),
	8095:  uint16(1),
	8096:  uint16(aux_sym_annotation_repeat1),
	8097:  uint16(4),
	8098:  uint16(3),
	8099:  uint16(1),
	8100:  uint16(sym_comment),
	8101:  uint16(876),
	8102:  uint16(1),
	8103:  uint16(anon_sym_COMMA),
	8104:  uint16(273),
	8105:  uint16(1),
	8106:  uint16(aux_sym_generic_parameters_repeat1),
	8107:  uint16(943),
	8108:  uint16(2),
	8109:  uint16(anon_sym_RPAREN),
	8110:  uint16(anon_sym_RBRACK),
	8111:  uint16(5),
	8112:  uint16(3),
	8113:  uint16(1),
	8114:  uint16(sym_comment),
	8115:  uint16(657),
	8116:  uint16(1),
	8117:  uint16(anon_sym_DOLLAR),
	8118:  uint16(699),
	8119:  uint16(1),
	8120:  uint16(anon_sym_LBRACE),
	8121:  uint16(701),
	8122:  uint16(1),
	8123:  uint16(anon_sym_extends),
	8124:  uint16(412),
	8125:  uint16(1),
	8126:  uint16(aux_sym_annotation_repeat1),
	8127:  uint16(5),
	8128:  uint16(3),
	8129:  uint16(1),
	8130:  uint16(sym_comment),
	8131:  uint16(657),
	8132:  uint16(1),
	8133:  uint16(anon_sym_DOLLAR),
	8134:  uint16(945),
	8135:  uint16(1),
	8136:  uint16(sym_unique_id),
	8137:  uint16(947),
	8138:  uint16(1),
	8139:  uint16(anon_sym_LBRACE),
	8140:  uint16(357),
	8141:  uint16(1),
	8142:  uint16(aux_sym_annotation_repeat1),
	8143:  uint16(4),
	8144:  uint16(3),
	8145:  uint16(1),
	8146:  uint16(sym_comment),
	8147:  uint16(892),
	8148:  uint16(1),
	8149:  uint16(sym_identifier),
	8150:  uint16(949),
	8151:  uint16(1),
	8152:  uint16(anon_sym_RBRACE),
	8153:  uint16(274),
	8154:  uint16(2),
	8155:  uint16(sym_enum_field),
	8156:  uint16(aux_sym_enum_repeat1),
	8157:  uint16(5),
	8158:  uint16(3),
	8159:  uint16(1),
	8160:  uint16(sym_comment),
	8161:  uint16(809),
	8162:  uint16(1),
	8163:  uint16(sym__normal_version),
	8164:  uint16(811),
	8165:  uint16(1),
	8166:  uint16(aux_sym__inline_version_token1),
	8167:  uint16(208),
	8168:  uint16(1),
	8169:  uint16(sym__inline_version),
	8170:  uint16(251),
	8171:  uint16(1),
	8172:  uint16(sym_field_version),
	8173:  uint16(5),
	8174:  uint16(3),
	8175:  uint16(1),
	8176:  uint16(sym_comment),
	8177:  uint16(657),
	8178:  uint16(1),
	8179:  uint16(anon_sym_DOLLAR),
	8180:  uint16(951),
	8181:  uint16(1),
	8182:  uint16(anon_sym_SEMI),
	8183:  uint16(953),
	8184:  uint16(1),
	8185:  uint16(anon_sym_DASH_GT),
	8186:  uint16(369),
	8187:  uint16(1),
	8188:  uint16(aux_sym_annotation_repeat1),
	8189:  uint16(5),
	8190:  uint16(3),
	8191:  uint16(1),
	8192:  uint16(sym_comment),
	8193:  uint16(657),
	8194:  uint16(1),
	8195:  uint16(anon_sym_DOLLAR),
	8196:  uint16(955),
	8197:  uint16(1),
	8198:  uint16(anon_sym_SEMI),
	8199:  uint16(957),
	8200:  uint16(1),
	8201:  uint16(anon_sym_COLON),
	8202:  uint16(402),
	8203:  uint16(1),
	8204:  uint16(aux_sym_annotation_repeat1),
	8205:  uint16(4),
	8206:  uint16(3),
	8207:  uint16(1),
	8208:  uint16(sym_comment),
	8209:  uint16(959),
	8210:  uint16(1),
	8211:  uint16(sym_identifier),
	8212:  uint16(290),
	8213:  uint16(1),
	8214:  uint16(aux_sym__annotation_array_def_repeat1),
	8215:  uint16(961),
	8216:  uint16(2),
	8217:  uint16(anon_sym_RPAREN),
	8218:  uint16(anon_sym_COMMA),
	8219:  uint16(4),
	8220:  uint16(3),
	8221:  uint16(1),
	8222:  uint16(sym_comment),
	8223:  uint16(965),
	8224:  uint16(1),
	8225:  uint16(anon_sym_COMMA),
	8226:  uint16(302),
	8227:  uint16(1),
	8228:  uint16(aux_sym_annotation_array_repeat1),
	8229:  uint16(963),
	8230:  uint16(2),
	8231:  uint16(anon_sym_RPAREN),
	8232:  uint16(anon_sym_RBRACK),
	8233:  uint16(5),
	8234:  uint16(3),
	8235:  uint16(1),
	8236:  uint16(sym_comment),
	8237:  uint16(657),
	8238:  uint16(1),
	8239:  uint16(anon_sym_DOLLAR),
	8240:  uint16(968),
	8241:  uint16(1),
	8242:  uint16(anon_sym_SEMI),
	8243:  uint16(970),
	8244:  uint16(1),
	8245:  uint16(anon_sym_DASH_GT),
	8246:  uint16(377),
	8247:  uint16(1),
	8248:  uint16(aux_sym_annotation_repeat1),
	8249:  uint16(4),
	8250:  uint16(3),
	8251:  uint16(1),
	8252:  uint16(sym_comment),
	8253:  uint16(972),
	8254:  uint16(1),
	8255:  uint16(anon_sym_COMMA),
	8256:  uint16(290),
	8257:  uint16(1),
	8258:  uint16(aux_sym__annotation_array_def_repeat1),
	8259:  uint16(961),
	8260:  uint16(2),
	8261:  uint16(anon_sym_RPAREN),
	8262:  uint16(anon_sym_RBRACK),
	8263:  uint16(4),
	8264:  uint16(3),
	8265:  uint16(1),
	8266:  uint16(sym_comment),
	8267:  uint16(892),
	8268:  uint16(1),
	8269:  uint16(sym_identifier),
	8270:  uint16(974),
	8271:  uint16(1),
	8272:  uint16(anon_sym_RBRACE),
	8273:  uint16(289),
	8274:  uint16(2),
	8275:  uint16(sym_enum_field),
	8276:  uint16(aux_sym_enum_repeat1),
	8277:  uint16(5),
	8278:  uint16(3),
	8279:  uint16(1),
	8280:  uint16(sym_comment),
	8281:  uint16(657),
	8282:  uint16(1),
	8283:  uint16(anon_sym_DOLLAR),
	8284:  uint16(976),
	8285:  uint16(1),
	8286:  uint16(anon_sym_SEMI),
	8287:  uint16(978),
	8288:  uint16(1),
	8289:  uint16(anon_sym_EQ),
	8290:  uint16(283),
	8291:  uint16(1),
	8292:  uint16(aux_sym_annotation_repeat1),
	8293:  uint16(4),
	8294:  uint16(3),
	8295:  uint16(1),
	8296:  uint16(sym_comment),
	8297:  uint16(892),
	8298:  uint16(1),
	8299:  uint16(sym_identifier),
	8300:  uint16(894),
	8301:  uint16(1),
	8302:  uint16(anon_sym_RBRACE),
	8303:  uint16(297),
	8304:  uint16(2),
	8305:  uint16(sym_enum_field),
	8306:  uint16(aux_sym_enum_repeat1),
	8307:  uint16(4),
	8308:  uint16(3),
	8309:  uint16(1),
	8310:  uint16(sym_comment),
	8311:  uint16(980),
	8312:  uint16(1),
	8313:  uint16(anon_sym_COMMA),
	8314:  uint16(982),
	8315:  uint16(1),
	8316:  uint16(anon_sym_RBRACK),
	8317:  uint16(409),
	8318:  uint16(1),
	8319:  uint16(aux_sym__annotation_array_def_repeat2),
	8320:  uint16(4),
	8321:  uint16(3),
	8322:  uint16(1),
	8323:  uint16(sym_comment),
	8324:  uint16(657),
	8325:  uint16(1),
	8326:  uint16(anon_sym_DOLLAR),
	8327:  uint16(984),
	8328:  uint16(1),
	8329:  uint16(anon_sym_SEMI),
	8330:  uint16(415),
	8331:  uint16(1),
	8332:  uint16(aux_sym_annotation_repeat1),
	8333:  uint16(4),
	8334:  uint16(3),
	8335:  uint16(1),
	8336:  uint16(sym_comment),
	8337:  uint16(961),
	8338:  uint16(1),
	8339:  uint16(anon_sym_RPAREN),
	8340:  uint16(972),
	8341:  uint16(1),
	8342:  uint16(anon_sym_COMMA),
	8343:  uint16(405),
	8344:  uint16(1),
	8345:  uint16(aux_sym__annotation_array_def_repeat1),
	8346:  uint16(3),
	8347:  uint16(3),
	8348:  uint16(1),
	8349:  uint16(sym_comment),
	8350:  uint16(986),
	8351:  uint16(1),
	8352:  uint16(anon_sym_EQ),
	8353:  uint16(988),
	8354:  uint16(2),
	8355:  uint16(anon_sym_RPAREN),
	8356:  uint16(anon_sym_COMMA),
	8357:  uint16(4),
	8358:  uint16(3),
	8359:  uint16(1),
	8360:  uint16(sym_comment),
	8361:  uint16(657),
	8362:  uint16(1),
	8363:  uint16(anon_sym_DOLLAR),
	8364:  uint16(990),
	8365:  uint16(1),
	8366:  uint16(anon_sym_LBRACE),
	8367:  uint16(195),
	8368:  uint16(1),
	8369:  uint16(aux_sym_annotation_repeat1),
	8370:  uint16(4),
	8371:  uint16(3),
	8372:  uint16(1),
	8373:  uint16(sym_comment),
	8374:  uint16(657),
	8375:  uint16(1),
	8376:  uint16(anon_sym_DOLLAR),
	8377:  uint16(990),
	8378:  uint16(1),
	8379:  uint16(anon_sym_LBRACE),
	8380:  uint16(381),
	8381:  uint16(1),
	8382:  uint16(aux_sym_annotation_repeat1),
	8383:  uint16(4),
	8384:  uint16(3),
	8385:  uint16(1),
	8386:  uint16(sym_comment),
	8387:  uint16(657),
	8388:  uint16(1),
	8389:  uint16(anon_sym_DOLLAR),
	8390:  uint16(951),
	8391:  uint16(1),
	8392:  uint16(anon_sym_SEMI),
	8393:  uint16(195),
	8394:  uint16(1),
	8395:  uint16(aux_sym_annotation_repeat1),
	8396:  uint16(4),
	8397:  uint16(3),
	8398:  uint16(1),
	8399:  uint16(sym_comment),
	8400:  uint16(765),
	8401:  uint16(1),
	8402:  uint16(anon_sym_COMMA),
	8403:  uint16(992),
	8404:  uint16(1),
	8405:  uint16(anon_sym_RPAREN),
	8406:  uint16(330),
	8407:  uint16(1),
	8408:  uint16(aux_sym_named_return_type_repeat1),
	8409:  uint16(4),
	8410:  uint16(3),
	8411:  uint16(1),
	8412:  uint16(sym_comment),
	8413:  uint16(45),
	8414:  uint16(1),
	8415:  uint16(anon_sym_DQUOTE),
	8416:  uint16(47),
	8417:  uint16(1),
	8418:  uint16(anon_sym_SQUOTE),
	8419:  uint16(347),
	8420:  uint16(1),
	8421:  uint16(sym_string),
	8422:  uint16(4),
	8423:  uint16(3),
	8424:  uint16(1),
	8425:  uint16(sym_comment),
	8426:  uint16(994),
	8427:  uint16(1),
	8428:  uint16(anon_sym_RPAREN),
	8429:  uint16(996),
	8430:  uint16(1),
	8431:  uint16(anon_sym_COMMA),
	8432:  uint16(360),
	8433:  uint16(1),
	8434:  uint16(aux_sym_parameters_repeat1),
	8435:  uint16(4),
	8436:  uint16(3),
	8437:  uint16(1),
	8438:  uint16(sym_comment),
	8439:  uint16(765),
	8440:  uint16(1),
	8441:  uint16(anon_sym_COMMA),
	8442:  uint16(998),
	8443:  uint16(1),
	8444:  uint16(anon_sym_RPAREN),
	8445:  uint16(315),
	8446:  uint16(1),
	8447:  uint16(aux_sym_named_return_type_repeat1),
	8448:  uint16(4),
	8449:  uint16(3),
	8450:  uint16(1),
	8451:  uint16(sym_comment),
	8452:  uint16(657),
	8453:  uint16(1),
	8454:  uint16(anon_sym_DOLLAR),
	8455:  uint16(968),
	8456:  uint16(1),
	8457:  uint16(anon_sym_SEMI),
	8458:  uint16(195),
	8459:  uint16(1),
	8460:  uint16(aux_sym_annotation_repeat1),
	8461:  uint16(4),
	8462:  uint16(3),
	8463:  uint16(1),
	8464:  uint16(sym_comment),
	8465:  uint16(657),
	8466:  uint16(1),
	8467:  uint16(anon_sym_DOLLAR),
	8468:  uint16(1000),
	8469:  uint16(1),
	8470:  uint16(anon_sym_LBRACE),
	8471:  uint16(195),
	8472:  uint16(1),
	8473:  uint16(aux_sym_annotation_repeat1),
	8474:  uint16(4),
	8475:  uint16(3),
	8476:  uint16(1),
	8477:  uint16(sym_comment),
	8478:  uint16(45),
	8479:  uint16(1),
	8480:  uint16(anon_sym_DQUOTE),
	8481:  uint16(47),
	8482:  uint16(1),
	8483:  uint16(anon_sym_SQUOTE),
	8484:  uint16(533),
	8485:  uint16(1),
	8486:  uint16(sym_string),
	8487:  uint16(4),
	8488:  uint16(3),
	8489:  uint16(1),
	8490:  uint16(sym_comment),
	8491:  uint16(657),
	8492:  uint16(1),
	8493:  uint16(anon_sym_DOLLAR),
	8494:  uint16(1002),
	8495:  uint16(1),
	8496:  uint16(anon_sym_LBRACE),
	8497:  uint16(195),
	8498:  uint16(1),
	8499:  uint16(aux_sym_annotation_repeat1),
	8500:  uint16(4),
	8501:  uint16(3),
	8502:  uint16(1),
	8503:  uint16(sym_comment),
	8504:  uint16(45),
	8505:  uint16(1),
	8506:  uint16(anon_sym_DQUOTE),
	8507:  uint16(47),
	8508:  uint16(1),
	8509:  uint16(anon_sym_SQUOTE),
	8510:  uint16(390),
	8511:  uint16(1),
	8512:  uint16(sym_string),
	8513:  uint16(4),
	8514:  uint16(3),
	8515:  uint16(1),
	8516:  uint16(sym_comment),
	8517:  uint16(657),
	8518:  uint16(1),
	8519:  uint16(anon_sym_DOLLAR),
	8520:  uint16(1002),
	8521:  uint16(1),
	8522:  uint16(anon_sym_LBRACE),
	8523:  uint16(391),
	8524:  uint16(1),
	8525:  uint16(aux_sym_annotation_repeat1),
	8526:  uint16(4),
	8527:  uint16(3),
	8528:  uint16(1),
	8529:  uint16(sym_comment),
	8530:  uint16(503),
	8531:  uint16(1),
	8532:  uint16(anon_sym_LPAREN),
	8533:  uint16(1004),
	8534:  uint16(1),
	8535:  uint16(anon_sym_RPAREN),
	8536:  uint16(506),
	8537:  uint16(1),
	8538:  uint16(sym_generics),
	8539:  uint16(4),
	8540:  uint16(3),
	8541:  uint16(1),
	8542:  uint16(sym_comment),
	8543:  uint16(765),
	8544:  uint16(1),
	8545:  uint16(anon_sym_COMMA),
	8546:  uint16(998),
	8547:  uint16(1),
	8548:  uint16(anon_sym_RPAREN),
	8549:  uint16(330),
	8550:  uint16(1),
	8551:  uint16(aux_sym_named_return_type_repeat1),
	8552:  uint16(3),
	8553:  uint16(3),
	8554:  uint16(1),
	8555:  uint16(sym_comment),
	8556:  uint16(1008),
	8557:  uint16(1),
	8558:  uint16(anon_sym_RBRACE),
	8559:  uint16(1006),
	8560:  uint16(2),
	8561:  uint16(anon_sym_union),
	8562:  uint16(sym_identifier),
	8563:  uint16(4),
	8564:  uint16(3),
	8565:  uint16(1),
	8566:  uint16(sym_comment),
	8567:  uint16(657),
	8568:  uint16(1),
	8569:  uint16(anon_sym_DOLLAR),
	8570:  uint16(1010),
	8571:  uint16(1),
	8572:  uint16(anon_sym_LBRACE),
	8573:  uint16(352),
	8574:  uint16(1),
	8575:  uint16(aux_sym_annotation_repeat1),
	8576:  uint16(2),
	8577:  uint16(3),
	8578:  uint16(1),
	8579:  uint16(sym_comment),
	8580:  uint16(1012),
	8581:  uint16(3),
	8582:  uint16(anon_sym_RPAREN),
	8583:  uint16(anon_sym_COMMA),
	8584:  uint16(sym_identifier),
	8585:  uint16(4),
	8586:  uint16(3),
	8587:  uint16(1),
	8588:  uint16(sym_comment),
	8589:  uint16(1014),
	8590:  uint16(1),
	8591:  uint16(anon_sym_RPAREN),
	8592:  uint16(1016),
	8593:  uint16(1),
	8594:  uint16(anon_sym_COMMA),
	8595:  uint16(330),
	8596:  uint16(1),
	8597:  uint16(aux_sym_named_return_type_repeat1),
	8598:  uint16(4),
	8599:  uint16(3),
	8600:  uint16(1),
	8601:  uint16(sym_comment),
	8602:  uint16(765),
	8603:  uint16(1),
	8604:  uint16(anon_sym_COMMA),
	8605:  uint16(1019),
	8606:  uint16(1),
	8607:  uint16(anon_sym_RPAREN),
	8608:  uint16(330),
	8609:  uint16(1),
	8610:  uint16(aux_sym_named_return_type_repeat1),
	8611:  uint16(4),
	8612:  uint16(3),
	8613:  uint16(1),
	8614:  uint16(sym_comment),
	8615:  uint16(765),
	8616:  uint16(1),
	8617:  uint16(anon_sym_COMMA),
	8618:  uint16(1019),
	8619:  uint16(1),
	8620:  uint16(anon_sym_RPAREN),
	8621:  uint16(326),
	8622:  uint16(1),
	8623:  uint16(aux_sym_named_return_type_repeat1),
	8624:  uint16(4),
	8625:  uint16(3),
	8626:  uint16(1),
	8627:  uint16(sym_comment),
	8628:  uint16(45),
	8629:  uint16(1),
	8630:  uint16(anon_sym_DQUOTE),
	8631:  uint16(47),
	8632:  uint16(1),
	8633:  uint16(anon_sym_SQUOTE),
	8634:  uint16(513),
	8635:  uint16(1),
	8636:  uint16(sym_string),
	8637:  uint16(3),
	8638:  uint16(3),
	8639:  uint16(1),
	8640:  uint16(sym_comment),
	8641:  uint16(1023),
	8642:  uint16(1),
	8643:  uint16(anon_sym_RBRACE),
	8644:  uint16(1021),
	8645:  uint16(2),
	8646:  uint16(anon_sym_union),
	8647:  uint16(sym_identifier),
	8648:  uint16(4),
	8649:  uint16(3),
	8650:  uint16(1),
	8651:  uint16(sym_comment),
	8652:  uint16(1025),
	8653:  uint16(1),
	8654:  uint16(anon_sym_RPAREN),
	8655:  uint16(1027),
	8656:  uint16(1),
	8657:  uint16(anon_sym_COMMA),
	8658:  uint16(335),
	8659:  uint16(1),
	8660:  uint16(aux_sym_top_level_annotation_body_repeat1),
	8661:  uint16(4),
	8662:  uint16(3),
	8663:  uint16(1),
	8664:  uint16(sym_comment),
	8665:  uint16(765),
	8666:  uint16(1),
	8667:  uint16(anon_sym_COMMA),
	8668:  uint16(923),
	8669:  uint16(1),
	8670:  uint16(anon_sym_RPAREN),
	8671:  uint16(330),
	8672:  uint16(1),
	8673:  uint16(aux_sym_named_return_type_repeat1),
	8674:  uint16(4),
	8675:  uint16(3),
	8676:  uint16(1),
	8677:  uint16(sym_comment),
	8678:  uint16(79),
	8679:  uint16(1),
	8680:  uint16(anon_sym_RPAREN),
	8681:  uint16(1030),
	8682:  uint16(1),
	8683:  uint16(anon_sym_COMMA),
	8684:  uint16(271),
	8685:  uint16(1),
	8686:  uint16(aux_sym__annotation_array_def_repeat2),
	8687:  uint16(4),
	8688:  uint16(3),
	8689:  uint16(1),
	8690:  uint16(sym_comment),
	8691:  uint16(657),
	8692:  uint16(1),
	8693:  uint16(anon_sym_DOLLAR),
	8694:  uint16(1032),
	8695:  uint16(1),
	8696:  uint16(anon_sym_SEMI),
	8697:  uint16(195),
	8698:  uint16(1),
	8699:  uint16(aux_sym_annotation_repeat1),
	8700:  uint16(4),
	8701:  uint16(3),
	8702:  uint16(1),
	8703:  uint16(sym_comment),
	8704:  uint16(1034),
	8705:  uint16(1),
	8706:  uint16(anon_sym_COMMA),
	8707:  uint16(1036),
	8708:  uint16(1),
	8709:  uint16(anon_sym_RBRACK),
	8710:  uint16(344),
	8711:  uint16(1),
	8712:  uint16(aux_sym_annotation_array_repeat1),
	8713:  uint16(2),
	8714:  uint16(3),
	8715:  uint16(1),
	8716:  uint16(sym_comment),
	8717:  uint16(1038),
	8718:  uint16(3),
	8719:  uint16(anon_sym_RPAREN),
	8720:  uint16(anon_sym_COMMA),
	8721:  uint16(anon_sym_RBRACK),
	8722:  uint16(4),
	8723:  uint16(3),
	8724:  uint16(1),
	8725:  uint16(sym_comment),
	8726:  uint16(1040),
	8727:  uint16(1),
	8728:  uint16(sym_unique_id),
	8729:  uint16(1042),
	8730:  uint16(1),
	8731:  uint16(anon_sym_LPAREN),
	8732:  uint16(300),
	8733:  uint16(1),
	8734:  uint16(sym_annotation_targets),
	8735:  uint16(4),
	8736:  uint16(3),
	8737:  uint16(1),
	8738:  uint16(sym_comment),
	8739:  uint16(1034),
	8740:  uint16(1),
	8741:  uint16(anon_sym_COMMA),
	8742:  uint16(1044),
	8743:  uint16(1),
	8744:  uint16(anon_sym_RPAREN),
	8745:  uint16(302),
	8746:  uint16(1),
	8747:  uint16(aux_sym_annotation_array_repeat1),
	8748:  uint16(4),
	8749:  uint16(3),
	8750:  uint16(1),
	8751:  uint16(sym_comment),
	8752:  uint16(1034),
	8753:  uint16(1),
	8754:  uint16(anon_sym_COMMA),
	8755:  uint16(1036),
	8756:  uint16(1),
	8757:  uint16(anon_sym_RPAREN),
	8758:  uint16(342),
	8759:  uint16(1),
	8760:  uint16(aux_sym_annotation_array_repeat1),
	8761:  uint16(4),
	8762:  uint16(3),
	8763:  uint16(1),
	8764:  uint16(sym_comment),
	8765:  uint16(1034),
	8766:  uint16(1),
	8767:  uint16(anon_sym_COMMA),
	8768:  uint16(1044),
	8769:  uint16(1),
	8770:  uint16(anon_sym_RBRACK),
	8771:  uint16(302),
	8772:  uint16(1),
	8773:  uint16(aux_sym_annotation_array_repeat1),
	8774:  uint16(3),
	8775:  uint16(3),
	8776:  uint16(1),
	8777:  uint16(sym_comment),
	8778:  uint16(271),
	8779:  uint16(1),
	8780:  uint16(aux_sym__annotation_array_def_repeat2),
	8781:  uint16(79),
	8782:  uint16(2),
	8783:  uint16(anon_sym_COMMA),
	8784:  uint16(anon_sym_RBRACK),
	8785:  uint16(4),
	8786:  uint16(3),
	8787:  uint16(1),
	8788:  uint16(sym_comment),
	8789:  uint16(961),
	8790:  uint16(1),
	8791:  uint16(anon_sym_RPAREN),
	8792:  uint16(972),
	8793:  uint16(1),
	8794:  uint16(anon_sym_COMMA),
	8795:  uint16(405),
	8796:  uint16(1),
	8797:  uint16(aux_sym__annotation_array_def_repeat1),
	8798:  uint16(4),
	8799:  uint16(3),
	8800:  uint16(1),
	8801:  uint16(sym_comment),
	8802:  uint16(1046),
	8803:  uint16(1),
	8804:  uint16(anon_sym_SEMI),
	8805:  uint16(1048),
	8806:  uint16(1),
	8807:  uint16(anon_sym_DOT),
	8808:  uint16(372),
	8809:  uint16(1),
	8810:  uint16(aux_sym_import_using_repeat1),
	8811:  uint16(4),
	8812:  uint16(3),
	8813:  uint16(1),
	8814:  uint16(sym_comment),
	8815:  uint16(657),
	8816:  uint16(1),
	8817:  uint16(anon_sym_DOLLAR),
	8818:  uint16(1050),
	8819:  uint16(1),
	8820:  uint16(anon_sym_SEMI),
	8821:  uint16(195),
	8822:  uint16(1),
	8823:  uint16(aux_sym_annotation_repeat1),
	8824:  uint16(4),
	8825:  uint16(3),
	8826:  uint16(1),
	8827:  uint16(sym_comment),
	8828:  uint16(657),
	8829:  uint16(1),
	8830:  uint16(anon_sym_DOLLAR),
	8831:  uint16(1050),
	8832:  uint16(1),
	8833:  uint16(anon_sym_SEMI),
	8834:  uint16(338),
	8835:  uint16(1),
	8836:  uint16(aux_sym_annotation_repeat1),
	8837:  uint16(2),
	8838:  uint16(3),
	8839:  uint16(1),
	8840:  uint16(sym_comment),
	8841:  uint16(1052),
	8842:  uint16(3),
	8843:  uint16(anon_sym_RPAREN),
	8844:  uint16(anon_sym_COMMA),
	8845:  uint16(anon_sym_RBRACK),
	8846:  uint16(4),
	8847:  uint16(3),
	8848:  uint16(1),
	8849:  uint16(sym_comment),
	8850:  uint16(657),
	8851:  uint16(1),
	8852:  uint16(anon_sym_DOLLAR),
	8853:  uint16(1054),
	8854:  uint16(1),
	8855:  uint16(anon_sym_SEMI),
	8856:  uint16(195),
	8857:  uint16(1),
	8858:  uint16(aux_sym_annotation_repeat1),
	8859:  uint16(4),
	8860:  uint16(3),
	8861:  uint16(1),
	8862:  uint16(sym_comment),
	8863:  uint16(657),
	8864:  uint16(1),
	8865:  uint16(anon_sym_DOLLAR),
	8866:  uint16(1056),
	8867:  uint16(1),
	8868:  uint16(anon_sym_LBRACE),
	8869:  uint16(195),
	8870:  uint16(1),
	8871:  uint16(aux_sym_annotation_repeat1),
	8872:  uint16(4),
	8873:  uint16(3),
	8874:  uint16(1),
	8875:  uint16(sym_comment),
	8876:  uint16(657),
	8877:  uint16(1),
	8878:  uint16(anon_sym_DOLLAR),
	8879:  uint16(1056),
	8880:  uint16(1),
	8881:  uint16(anon_sym_LBRACE),
	8882:  uint16(411),
	8883:  uint16(1),
	8884:  uint16(aux_sym_annotation_repeat1),
	8885:  uint16(4),
	8886:  uint16(3),
	8887:  uint16(1),
	8888:  uint16(sym_comment),
	8889:  uint16(657),
	8890:  uint16(1),
	8891:  uint16(anon_sym_DOLLAR),
	8892:  uint16(1058),
	8893:  uint16(1),
	8894:  uint16(anon_sym_LBRACE),
	8895:  uint16(413),
	8896:  uint16(1),
	8897:  uint16(aux_sym_annotation_repeat1),
	8898:  uint16(4),
	8899:  uint16(3),
	8900:  uint16(1),
	8901:  uint16(sym_comment),
	8902:  uint16(657),
	8903:  uint16(1),
	8904:  uint16(anon_sym_DOLLAR),
	8905:  uint16(1060),
	8906:  uint16(1),
	8907:  uint16(anon_sym_LBRACE),
	8908:  uint16(404),
	8909:  uint16(1),
	8910:  uint16(aux_sym_annotation_repeat1),
	8911:  uint16(4),
	8912:  uint16(3),
	8913:  uint16(1),
	8914:  uint16(sym_comment),
	8915:  uint16(1062),
	8916:  uint16(1),
	8917:  uint16(anon_sym_RPAREN),
	8918:  uint16(1064),
	8919:  uint16(1),
	8920:  uint16(anon_sym_COMMA),
	8921:  uint16(335),
	8922:  uint16(1),
	8923:  uint16(aux_sym_top_level_annotation_body_repeat1),
	8924:  uint16(4),
	8925:  uint16(3),
	8926:  uint16(1),
	8927:  uint16(sym_comment),
	8928:  uint16(657),
	8929:  uint16(1),
	8930:  uint16(anon_sym_DOLLAR),
	8931:  uint16(1060),
	8932:  uint16(1),
	8933:  uint16(anon_sym_LBRACE),
	8934:  uint16(195),
	8935:  uint16(1),
	8936:  uint16(aux_sym_annotation_repeat1),
	8937:  uint16(2),
	8938:  uint16(3),
	8939:  uint16(1),
	8940:  uint16(sym_comment),
	8941:  uint16(1066),
	8942:  uint16(3),
	8943:  uint16(anon_sym_SEMI),
	8944:  uint16(anon_sym_DOLLAR),
	8945:  uint16(anon_sym_DASH_GT),
	8946:  uint16(3),
	8947:  uint16(3),
	8948:  uint16(1),
	8949:  uint16(sym_comment),
	8950:  uint16(1070),
	8951:  uint16(1),
	8952:  uint16(anon_sym_RBRACE),
	8953:  uint16(1068),
	8954:  uint16(2),
	8955:  uint16(anon_sym_union),
	8956:  uint16(sym_identifier),
	8957:  uint16(4),
	8958:  uint16(3),
	8959:  uint16(1),
	8960:  uint16(sym_comment),
	8961:  uint16(996),
	8962:  uint16(1),
	8963:  uint16(anon_sym_COMMA),
	8964:  uint16(1072),
	8965:  uint16(1),
	8966:  uint16(anon_sym_RPAREN),
	8967:  uint16(417),
	8968:  uint16(1),
	8969:  uint16(aux_sym_parameters_repeat1),
	8970:  uint16(4),
	8971:  uint16(3),
	8972:  uint16(1),
	8973:  uint16(sym_comment),
	8974:  uint16(657),
	8975:  uint16(1),
	8976:  uint16(anon_sym_DOLLAR),
	8977:  uint16(699),
	8978:  uint16(1),
	8979:  uint16(anon_sym_LBRACE),
	8980:  uint16(195),
	8981:  uint16(1),
	8982:  uint16(aux_sym_annotation_repeat1),
	8983:  uint16(4),
	8984:  uint16(3),
	8985:  uint16(1),
	8986:  uint16(sym_comment),
	8987:  uint16(1074),
	8988:  uint16(1),
	8989:  uint16(sym_identifier),
	8990:  uint16(1076),
	8991:  uint16(1),
	8992:  uint16(anon_sym_RPAREN),
	8993:  uint16(503),
	8994:  uint16(1),
	8995:  uint16(sym_named_return_type),
	8996:  uint16(2),
	8997:  uint16(3),
	8998:  uint16(1),
	8999:  uint16(sym_comment),
	9000:  uint16(1078),
	9001:  uint16(3),
	9002:  uint16(anon_sym_SEMI),
	9003:  uint16(anon_sym_LPAREN),
	9004:  uint16(anon_sym_DOLLAR),
	9005:  uint16(4),
	9006:  uint16(3),
	9007:  uint16(1),
	9008:  uint16(sym_comment),
	9009:  uint16(657),
	9010:  uint16(1),
	9011:  uint16(anon_sym_DOLLAR),
	9012:  uint16(1080),
	9013:  uint16(1),
	9014:  uint16(anon_sym_SEMI),
	9015:  uint16(422),
	9016:  uint16(1),
	9017:  uint16(aux_sym_annotation_repeat1),
	9018:  uint16(4),
	9019:  uint16(3),
	9020:  uint16(1),
	9021:  uint16(sym_comment),
	9022:  uint16(657),
	9023:  uint16(1),
	9024:  uint16(anon_sym_DOLLAR),
	9025:  uint16(1082),
	9026:  uint16(1),
	9027:  uint16(anon_sym_SEMI),
	9028:  uint16(195),
	9029:  uint16(1),
	9030:  uint16(aux_sym_annotation_repeat1),
	9031:  uint16(4),
	9032:  uint16(3),
	9033:  uint16(1),
	9034:  uint16(sym_comment),
	9035:  uint16(657),
	9036:  uint16(1),
	9037:  uint16(anon_sym_DOLLAR),
	9038:  uint16(929),
	9039:  uint16(1),
	9040:  uint16(anon_sym_LBRACE),
	9041:  uint16(420),
	9042:  uint16(1),
	9043:  uint16(aux_sym_annotation_repeat1),
	9044:  uint16(2),
	9045:  uint16(3),
	9046:  uint16(1),
	9047:  uint16(sym_comment),
	9048:  uint16(867),
	9049:  uint16(3),
	9050:  uint16(anon_sym_RPAREN),
	9051:  uint16(anon_sym_COMMA),
	9052:  uint16(anon_sym_RBRACK),
	9053:  uint16(3),
	9054:  uint16(3),
	9055:  uint16(1),
	9056:  uint16(sym_comment),
	9057:  uint16(1084),
	9058:  uint16(1),
	9059:  uint16(anon_sym_COMMA),
	9060:  uint16(959),
	9061:  uint16(2),
	9062:  uint16(anon_sym_RPAREN),
	9063:  uint16(sym_identifier),
	9064:  uint16(4),
	9065:  uint16(3),
	9066:  uint16(1),
	9067:  uint16(sym_comment),
	9068:  uint16(657),
	9069:  uint16(1),
	9070:  uint16(anon_sym_DOLLAR),
	9071:  uint16(1080),
	9072:  uint16(1),
	9073:  uint16(anon_sym_SEMI),
	9074:  uint16(195),
	9075:  uint16(1),
	9076:  uint16(aux_sym_annotation_repeat1),
	9077:  uint16(4),
	9078:  uint16(3),
	9079:  uint16(1),
	9080:  uint16(sym_comment),
	9081:  uint16(657),
	9082:  uint16(1),
	9083:  uint16(anon_sym_DOLLAR),
	9084:  uint16(1086),
	9085:  uint16(1),
	9086:  uint16(anon_sym_SEMI),
	9087:  uint16(426),
	9088:  uint16(1),
	9089:  uint16(aux_sym_annotation_repeat1),
	9090:  uint16(4),
	9091:  uint16(3),
	9092:  uint16(1),
	9093:  uint16(sym_comment),
	9094:  uint16(657),
	9095:  uint16(1),
	9096:  uint16(anon_sym_DOLLAR),
	9097:  uint16(929),
	9098:  uint16(1),
	9099:  uint16(anon_sym_LBRACE),
	9100:  uint16(195),
	9101:  uint16(1),
	9102:  uint16(aux_sym_annotation_repeat1),
	9103:  uint16(4),
	9104:  uint16(3),
	9105:  uint16(1),
	9106:  uint16(sym_comment),
	9107:  uint16(1048),
	9108:  uint16(1),
	9109:  uint16(anon_sym_DOT),
	9110:  uint16(1088),
	9111:  uint16(1),
	9112:  uint16(anon_sym_SEMI),
	9113:  uint16(428),
	9114:  uint16(1),
	9115:  uint16(aux_sym_import_using_repeat1),
	9116:  uint16(4),
	9117:  uint16(3),
	9118:  uint16(1),
	9119:  uint16(sym_comment),
	9120:  uint16(657),
	9121:  uint16(1),
	9122:  uint16(anon_sym_DOLLAR),
	9123:  uint16(1090),
	9124:  uint16(1),
	9125:  uint16(anon_sym_LBRACE),
	9126:  uint16(322),
	9127:  uint16(1),
	9128:  uint16(aux_sym_annotation_repeat1),
	9129:  uint16(4),
	9130:  uint16(3),
	9131:  uint16(1),
	9132:  uint16(sym_comment),
	9133:  uint16(657),
	9134:  uint16(1),
	9135:  uint16(anon_sym_DOLLAR),
	9136:  uint16(1092),
	9137:  uint16(1),
	9138:  uint16(anon_sym_LBRACE),
	9139:  uint16(320),
	9140:  uint16(1),
	9141:  uint16(aux_sym_annotation_repeat1),
	9142:  uint16(4),
	9143:  uint16(3),
	9144:  uint16(1),
	9145:  uint16(sym_comment),
	9146:  uint16(657),
	9147:  uint16(1),
	9148:  uint16(anon_sym_DOLLAR),
	9149:  uint16(863),
	9150:  uint16(1),
	9151:  uint16(anon_sym_SEMI),
	9152:  uint16(195),
	9153:  uint16(1),
	9154:  uint16(aux_sym_annotation_repeat1),
	9155:  uint16(3),
	9156:  uint16(3),
	9157:  uint16(1),
	9158:  uint16(sym_comment),
	9159:  uint16(1096),
	9160:  uint16(1),
	9161:  uint16(anon_sym_RBRACE),
	9162:  uint16(1094),
	9163:  uint16(2),
	9164:  uint16(anon_sym_union),
	9165:  uint16(sym_identifier),
	9166:  uint16(4),
	9167:  uint16(3),
	9168:  uint16(1),
	9169:  uint16(sym_comment),
	9170:  uint16(657),
	9171:  uint16(1),
	9172:  uint16(anon_sym_DOLLAR),
	9173:  uint16(1086),
	9174:  uint16(1),
	9175:  uint16(anon_sym_SEMI),
	9176:  uint16(195),
	9177:  uint16(1),
	9178:  uint16(aux_sym_annotation_repeat1),
	9179:  uint16(3),
	9180:  uint16(3),
	9181:  uint16(1),
	9182:  uint16(sym_comment),
	9183:  uint16(1100),
	9184:  uint16(1),
	9185:  uint16(anon_sym_RBRACE),
	9186:  uint16(1098),
	9187:  uint16(2),
	9188:  uint16(anon_sym_union),
	9189:  uint16(sym_identifier),
	9190:  uint16(4),
	9191:  uint16(3),
	9192:  uint16(1),
	9193:  uint16(sym_comment),
	9194:  uint16(45),
	9195:  uint16(1),
	9196:  uint16(anon_sym_DQUOTE),
	9197:  uint16(47),
	9198:  uint16(1),
	9199:  uint16(anon_sym_SQUOTE),
	9200:  uint16(517),
	9201:  uint16(1),
	9202:  uint16(sym_string),
	9203:  uint16(4),
	9204:  uint16(3),
	9205:  uint16(1),
	9206:  uint16(sym_comment),
	9207:  uint16(657),
	9208:  uint16(1),
	9209:  uint16(anon_sym_DOLLAR),
	9210:  uint16(1102),
	9211:  uint16(1),
	9212:  uint16(anon_sym_LBRACE),
	9213:  uint16(312),
	9214:  uint16(1),
	9215:  uint16(aux_sym_annotation_repeat1),
	9216:  uint16(4),
	9217:  uint16(3),
	9218:  uint16(1),
	9219:  uint16(sym_comment),
	9220:  uint16(657),
	9221:  uint16(1),
	9222:  uint16(anon_sym_DOLLAR),
	9223:  uint16(1104),
	9224:  uint16(1),
	9225:  uint16(anon_sym_LBRACE),
	9226:  uint16(195),
	9227:  uint16(1),
	9228:  uint16(aux_sym_annotation_repeat1),
	9229:  uint16(4),
	9230:  uint16(3),
	9231:  uint16(1),
	9232:  uint16(sym_comment),
	9233:  uint16(815),
	9234:  uint16(1),
	9235:  uint16(anon_sym_LPAREN),
	9236:  uint16(1106),
	9237:  uint16(1),
	9238:  uint16(sym_identifier),
	9239:  uint16(303),
	9240:  uint16(1),
	9241:  uint16(sym_method_parameters),
	9242:  uint16(2),
	9243:  uint16(3),
	9244:  uint16(1),
	9245:  uint16(sym_comment),
	9246:  uint16(1108),
	9247:  uint16(3),
	9248:  uint16(anon_sym_RPAREN),
	9249:  uint16(anon_sym_COMMA),
	9250:  uint16(anon_sym_RBRACK),
	9251:  uint16(4),
	9252:  uint16(3),
	9253:  uint16(1),
	9254:  uint16(sym_comment),
	9255:  uint16(503),
	9256:  uint16(1),
	9257:  uint16(anon_sym_LPAREN),
	9258:  uint16(1110),
	9259:  uint16(1),
	9260:  uint16(anon_sym_SEMI),
	9261:  uint16(524),
	9262:  uint16(1),
	9263:  uint16(sym_generics),
	9264:  uint16(4),
	9265:  uint16(3),
	9266:  uint16(1),
	9267:  uint16(sym_comment),
	9268:  uint16(503),
	9269:  uint16(1),
	9270:  uint16(anon_sym_LPAREN),
	9271:  uint16(1112),
	9272:  uint16(1),
	9273:  uint16(anon_sym_RPAREN),
	9274:  uint16(535),
	9275:  uint16(1),
	9276:  uint16(sym_generics),
	9277:  uint16(4),
	9278:  uint16(3),
	9279:  uint16(1),
	9280:  uint16(sym_comment),
	9281:  uint16(657),
	9282:  uint16(1),
	9283:  uint16(anon_sym_DOLLAR),
	9284:  uint16(1114),
	9285:  uint16(1),
	9286:  uint16(anon_sym_SEMI),
	9287:  uint16(195),
	9288:  uint16(1),
	9289:  uint16(aux_sym_annotation_repeat1),
	9290:  uint16(4),
	9291:  uint16(3),
	9292:  uint16(1),
	9293:  uint16(sym_comment),
	9294:  uint16(55),
	9295:  uint16(1),
	9296:  uint16(anon_sym_RPAREN),
	9297:  uint16(1116),
	9298:  uint16(1),
	9299:  uint16(sym_identifier),
	9300:  uint16(434),
	9301:  uint16(1),
	9302:  uint16(aux_sym_struct_shorthand_repeat1),
	9303:  uint16(4),
	9304:  uint16(3),
	9305:  uint16(1),
	9306:  uint16(sym_comment),
	9307:  uint16(908),
	9308:  uint16(1),
	9309:  uint16(anon_sym_RPAREN),
	9310:  uint16(910),
	9311:  uint16(1),
	9312:  uint16(anon_sym_COMMA),
	9313:  uint16(337),
	9314:  uint16(1),
	9315:  uint16(aux_sym__annotation_array_def_repeat2),
	9316:  uint16(2),
	9317:  uint16(3),
	9318:  uint16(1),
	9319:  uint16(sym_comment),
	9320:  uint16(1118),
	9321:  uint16(3),
	9322:  uint16(anon_sym_RPAREN),
	9323:  uint16(anon_sym_COMMA),
	9324:  uint16(anon_sym_RBRACK),
	9325:  uint16(4),
	9326:  uint16(3),
	9327:  uint16(1),
	9328:  uint16(sym_comment),
	9329:  uint16(1048),
	9330:  uint16(1),
	9331:  uint16(anon_sym_DOT),
	9332:  uint16(1120),
	9333:  uint16(1),
	9334:  uint16(anon_sym_SEMI),
	9335:  uint16(438),
	9336:  uint16(1),
	9337:  uint16(aux_sym_import_using_repeat1),
	9338:  uint16(4),
	9339:  uint16(3),
	9340:  uint16(1),
	9341:  uint16(sym_comment),
	9342:  uint16(657),
	9343:  uint16(1),
	9344:  uint16(anon_sym_DOLLAR),
	9345:  uint16(1122),
	9346:  uint16(1),
	9347:  uint16(anon_sym_LBRACE),
	9348:  uint16(195),
	9349:  uint16(1),
	9350:  uint16(aux_sym_annotation_repeat1),
	9351:  uint16(4),
	9352:  uint16(3),
	9353:  uint16(1),
	9354:  uint16(sym_comment),
	9355:  uint16(45),
	9356:  uint16(1),
	9357:  uint16(anon_sym_DQUOTE),
	9358:  uint16(47),
	9359:  uint16(1),
	9360:  uint16(anon_sym_SQUOTE),
	9361:  uint16(225),
	9362:  uint16(1),
	9363:  uint16(sym_string),
	9364:  uint16(4),
	9365:  uint16(3),
	9366:  uint16(1),
	9367:  uint16(sym_comment),
	9368:  uint16(657),
	9369:  uint16(1),
	9370:  uint16(anon_sym_DOLLAR),
	9371:  uint16(1124),
	9372:  uint16(1),
	9373:  uint16(anon_sym_SEMI),
	9374:  uint16(348),
	9375:  uint16(1),
	9376:  uint16(aux_sym_annotation_repeat1),
	9377:  uint16(4),
	9378:  uint16(3),
	9379:  uint16(1),
	9380:  uint16(sym_comment),
	9381:  uint16(657),
	9382:  uint16(1),
	9383:  uint16(anon_sym_DOLLAR),
	9384:  uint16(1126),
	9385:  uint16(1),
	9386:  uint16(anon_sym_SEMI),
	9387:  uint16(195),
	9388:  uint16(1),
	9389:  uint16(aux_sym_annotation_repeat1),
	9390:  uint16(4),
	9391:  uint16(3),
	9392:  uint16(1),
	9393:  uint16(sym_comment),
	9394:  uint16(657),
	9395:  uint16(1),
	9396:  uint16(anon_sym_DOLLAR),
	9397:  uint16(1128),
	9398:  uint16(1),
	9399:  uint16(anon_sym_SEMI),
	9400:  uint16(195),
	9401:  uint16(1),
	9402:  uint16(aux_sym_annotation_repeat1),
	9403:  uint16(4),
	9404:  uint16(3),
	9405:  uint16(1),
	9406:  uint16(sym_comment),
	9407:  uint16(1130),
	9408:  uint16(1),
	9409:  uint16(anon_sym_RPAREN),
	9410:  uint16(1132),
	9411:  uint16(1),
	9412:  uint16(anon_sym_COMMA),
	9413:  uint16(440),
	9414:  uint16(1),
	9415:  uint16(aux_sym_annotation_targets_repeat1),
	9416:  uint16(4),
	9417:  uint16(3),
	9418:  uint16(1),
	9419:  uint16(sym_comment),
	9420:  uint16(657),
	9421:  uint16(1),
	9422:  uint16(anon_sym_DOLLAR),
	9423:  uint16(1134),
	9424:  uint16(1),
	9425:  uint16(anon_sym_SEMI),
	9426:  uint16(195),
	9427:  uint16(1),
	9428:  uint16(aux_sym_annotation_repeat1),
	9429:  uint16(4),
	9430:  uint16(3),
	9431:  uint16(1),
	9432:  uint16(sym_comment),
	9433:  uint16(1136),
	9434:  uint16(1),
	9435:  uint16(anon_sym_RPAREN),
	9436:  uint16(1138),
	9437:  uint16(1),
	9438:  uint16(anon_sym_COMMA),
	9439:  uint16(398),
	9440:  uint16(1),
	9441:  uint16(aux_sym_annotation_targets_repeat1),
	9442:  uint16(2),
	9443:  uint16(3),
	9444:  uint16(1),
	9445:  uint16(sym_comment),
	9446:  uint16(1141),
	9447:  uint16(3),
	9448:  uint16(anon_sym_SEMI),
	9449:  uint16(anon_sym_DOLLAR),
	9450:  uint16(anon_sym_COLON),
	9451:  uint16(2),
	9452:  uint16(3),
	9453:  uint16(1),
	9454:  uint16(sym_comment),
	9455:  uint16(963),
	9456:  uint16(3),
	9457:  uint16(anon_sym_RPAREN),
	9458:  uint16(anon_sym_COMMA),
	9459:  uint16(anon_sym_RBRACK),
	9460:  uint16(2),
	9461:  uint16(3),
	9462:  uint16(1),
	9463:  uint16(sym_comment),
	9464:  uint16(1143),
	9465:  uint16(3),
	9466:  uint16(anon_sym_RPAREN),
	9467:  uint16(anon_sym_COMMA),
	9468:  uint16(anon_sym_RBRACK),
	9469:  uint16(4),
	9470:  uint16(3),
	9471:  uint16(1),
	9472:  uint16(sym_comment),
	9473:  uint16(657),
	9474:  uint16(1),
	9475:  uint16(anon_sym_DOLLAR),
	9476:  uint16(888),
	9477:  uint16(1),
	9478:  uint16(anon_sym_SEMI),
	9479:  uint16(195),
	9480:  uint16(1),
	9481:  uint16(aux_sym_annotation_repeat1),
	9482:  uint16(4),
	9483:  uint16(3),
	9484:  uint16(1),
	9485:  uint16(sym_comment),
	9486:  uint16(657),
	9487:  uint16(1),
	9488:  uint16(anon_sym_DOLLAR),
	9489:  uint16(1134),
	9490:  uint16(1),
	9491:  uint16(anon_sym_SEMI),
	9492:  uint16(351),
	9493:  uint16(1),
	9494:  uint16(aux_sym_annotation_repeat1),
	9495:  uint16(4),
	9496:  uint16(3),
	9497:  uint16(1),
	9498:  uint16(sym_comment),
	9499:  uint16(657),
	9500:  uint16(1),
	9501:  uint16(anon_sym_DOLLAR),
	9502:  uint16(1145),
	9503:  uint16(1),
	9504:  uint16(anon_sym_LBRACE),
	9505:  uint16(195),
	9506:  uint16(1),
	9507:  uint16(aux_sym_annotation_repeat1),
	9508:  uint16(4),
	9509:  uint16(3),
	9510:  uint16(1),
	9511:  uint16(sym_comment),
	9512:  uint16(931),
	9513:  uint16(1),
	9514:  uint16(anon_sym_RPAREN),
	9515:  uint16(972),
	9516:  uint16(1),
	9517:  uint16(anon_sym_COMMA),
	9518:  uint16(275),
	9519:  uint16(1),
	9520:  uint16(aux_sym__annotation_array_def_repeat1),
	9521:  uint16(4),
	9522:  uint16(3),
	9523:  uint16(1),
	9524:  uint16(sym_comment),
	9525:  uint16(961),
	9526:  uint16(1),
	9527:  uint16(anon_sym_RPAREN),
	9528:  uint16(972),
	9529:  uint16(1),
	9530:  uint16(anon_sym_COMMA),
	9531:  uint16(405),
	9532:  uint16(1),
	9533:  uint16(aux_sym__annotation_array_def_repeat1),
	9534:  uint16(3),
	9535:  uint16(3),
	9536:  uint16(1),
	9537:  uint16(sym_comment),
	9538:  uint16(1149),
	9539:  uint16(1),
	9540:  uint16(anon_sym_RBRACE),
	9541:  uint16(1147),
	9542:  uint16(2),
	9543:  uint16(anon_sym_union),
	9544:  uint16(sym_identifier),
	9545:  uint16(4),
	9546:  uint16(3),
	9547:  uint16(1),
	9548:  uint16(sym_comment),
	9549:  uint16(1064),
	9550:  uint16(1),
	9551:  uint16(anon_sym_COMMA),
	9552:  uint16(1151),
	9553:  uint16(1),
	9554:  uint16(anon_sym_RPAREN),
	9555:  uint16(356),
	9556:  uint16(1),
	9557:  uint16(aux_sym_top_level_annotation_body_repeat1),
	9558:  uint16(4),
	9559:  uint16(3),
	9560:  uint16(1),
	9561:  uint16(sym_comment),
	9562:  uint16(97),
	9563:  uint16(1),
	9564:  uint16(anon_sym_RBRACK),
	9565:  uint16(1153),
	9566:  uint16(1),
	9567:  uint16(anon_sym_COMMA),
	9568:  uint16(271),
	9569:  uint16(1),
	9570:  uint16(aux_sym__annotation_array_def_repeat2),
	9571:  uint16(4),
	9572:  uint16(3),
	9573:  uint16(1),
	9574:  uint16(sym_comment),
	9575:  uint16(1155),
	9576:  uint16(1),
	9577:  uint16(sym_identifier),
	9578:  uint16(1158),
	9579:  uint16(1),
	9580:  uint16(anon_sym_RPAREN),
	9581:  uint16(410),
	9582:  uint16(1),
	9583:  uint16(aux_sym_struct_shorthand_repeat1),
	9584:  uint16(4),
	9585:  uint16(3),
	9586:  uint16(1),
	9587:  uint16(sym_comment),
	9588:  uint16(657),
	9589:  uint16(1),
	9590:  uint16(anon_sym_DOLLAR),
	9591:  uint16(1160),
	9592:  uint16(1),
	9593:  uint16(anon_sym_LBRACE),
	9594:  uint16(195),
	9595:  uint16(1),
	9596:  uint16(aux_sym_annotation_repeat1),
	9597:  uint16(4),
	9598:  uint16(3),
	9599:  uint16(1),
	9600:  uint16(sym_comment),
	9601:  uint16(657),
	9602:  uint16(1),
	9603:  uint16(anon_sym_DOLLAR),
	9604:  uint16(933),
	9605:  uint16(1),
	9606:  uint16(anon_sym_LBRACE),
	9607:  uint16(195),
	9608:  uint16(1),
	9609:  uint16(aux_sym_annotation_repeat1),
	9610:  uint16(4),
	9611:  uint16(3),
	9612:  uint16(1),
	9613:  uint16(sym_comment),
	9614:  uint16(657),
	9615:  uint16(1),
	9616:  uint16(anon_sym_DOLLAR),
	9617:  uint16(1162),
	9618:  uint16(1),
	9619:  uint16(anon_sym_LBRACE),
	9620:  uint16(195),
	9621:  uint16(1),
	9622:  uint16(aux_sym_annotation_repeat1),
	9623:  uint16(4),
	9624:  uint16(3),
	9625:  uint16(1),
	9626:  uint16(sym_comment),
	9627:  uint16(657),
	9628:  uint16(1),
	9629:  uint16(anon_sym_DOLLAR),
	9630:  uint16(1162),
	9631:  uint16(1),
	9632:  uint16(anon_sym_LBRACE),
	9633:  uint16(425),
	9634:  uint16(1),
	9635:  uint16(aux_sym_annotation_repeat1),
	9636:  uint16(4),
	9637:  uint16(3),
	9638:  uint16(1),
	9639:  uint16(sym_comment),
	9640:  uint16(657),
	9641:  uint16(1),
	9642:  uint16(anon_sym_DOLLAR),
	9643:  uint16(1164),
	9644:  uint16(1),
	9645:  uint16(anon_sym_SEMI),
	9646:  uint16(195),
	9647:  uint16(1),
	9648:  uint16(aux_sym_annotation_repeat1),
	9649:  uint16(4),
	9650:  uint16(3),
	9651:  uint16(1),
	9652:  uint16(sym_comment),
	9653:  uint16(657),
	9654:  uint16(1),
	9655:  uint16(anon_sym_DOLLAR),
	9656:  uint16(1166),
	9657:  uint16(1),
	9658:  uint16(anon_sym_SEMI),
	9659:  uint16(365),
	9660:  uint16(1),
	9661:  uint16(aux_sym_annotation_repeat1),
	9662:  uint16(4),
	9663:  uint16(3),
	9664:  uint16(1),
	9665:  uint16(sym_comment),
	9666:  uint16(1168),
	9667:  uint16(1),
	9668:  uint16(anon_sym_RPAREN),
	9669:  uint16(1170),
	9670:  uint16(1),
	9671:  uint16(anon_sym_COMMA),
	9672:  uint16(417),
	9673:  uint16(1),
	9674:  uint16(aux_sym_parameters_repeat1),
	9675:  uint16(4),
	9676:  uint16(3),
	9677:  uint16(1),
	9678:  uint16(sym_comment),
	9679:  uint16(657),
	9680:  uint16(1),
	9681:  uint16(anon_sym_DOLLAR),
	9682:  uint16(1166),
	9683:  uint16(1),
	9684:  uint16(anon_sym_SEMI),
	9685:  uint16(195),
	9686:  uint16(1),
	9687:  uint16(aux_sym_annotation_repeat1),
	9688:  uint16(4),
	9689:  uint16(3),
	9690:  uint16(1),
	9691:  uint16(sym_comment),
	9692:  uint16(657),
	9693:  uint16(1),
	9694:  uint16(anon_sym_DOLLAR),
	9695:  uint16(1173),
	9696:  uint16(1),
	9697:  uint16(anon_sym_SEMI),
	9698:  uint16(195),
	9699:  uint16(1),
	9700:  uint16(aux_sym_annotation_repeat1),
	9701:  uint16(4),
	9702:  uint16(3),
	9703:  uint16(1),
	9704:  uint16(sym_comment),
	9705:  uint16(657),
	9706:  uint16(1),
	9707:  uint16(anon_sym_DOLLAR),
	9708:  uint16(1175),
	9709:  uint16(1),
	9710:  uint16(anon_sym_LBRACE),
	9711:  uint16(195),
	9712:  uint16(1),
	9713:  uint16(aux_sym_annotation_repeat1),
	9714:  uint16(4),
	9715:  uint16(3),
	9716:  uint16(1),
	9717:  uint16(sym_comment),
	9718:  uint16(657),
	9719:  uint16(1),
	9720:  uint16(anon_sym_DOLLAR),
	9721:  uint16(1177),
	9722:  uint16(1),
	9723:  uint16(anon_sym_LBRACE),
	9724:  uint16(195),
	9725:  uint16(1),
	9726:  uint16(aux_sym_annotation_repeat1),
	9727:  uint16(4),
	9728:  uint16(3),
	9729:  uint16(1),
	9730:  uint16(sym_comment),
	9731:  uint16(657),
	9732:  uint16(1),
	9733:  uint16(anon_sym_DOLLAR),
	9734:  uint16(1179),
	9735:  uint16(1),
	9736:  uint16(anon_sym_SEMI),
	9737:  uint16(195),
	9738:  uint16(1),
	9739:  uint16(aux_sym_annotation_repeat1),
	9740:  uint16(4),
	9741:  uint16(3),
	9742:  uint16(1),
	9743:  uint16(sym_comment),
	9744:  uint16(657),
	9745:  uint16(1),
	9746:  uint16(anon_sym_DOLLAR),
	9747:  uint16(1175),
	9748:  uint16(1),
	9749:  uint16(anon_sym_LBRACE),
	9750:  uint16(421),
	9751:  uint16(1),
	9752:  uint16(aux_sym_annotation_repeat1),
	9753:  uint16(4),
	9754:  uint16(3),
	9755:  uint16(1),
	9756:  uint16(sym_comment),
	9757:  uint16(657),
	9758:  uint16(1),
	9759:  uint16(anon_sym_DOLLAR),
	9760:  uint16(1179),
	9761:  uint16(1),
	9762:  uint16(anon_sym_SEMI),
	9763:  uint16(419),
	9764:  uint16(1),
	9765:  uint16(aux_sym_annotation_repeat1),
	9766:  uint16(4),
	9767:  uint16(3),
	9768:  uint16(1),
	9769:  uint16(sym_comment),
	9770:  uint16(657),
	9771:  uint16(1),
	9772:  uint16(anon_sym_DOLLAR),
	9773:  uint16(1181),
	9774:  uint16(1),
	9775:  uint16(anon_sym_LBRACE),
	9776:  uint16(195),
	9777:  uint16(1),
	9778:  uint16(aux_sym_annotation_repeat1),
	9779:  uint16(4),
	9780:  uint16(3),
	9781:  uint16(1),
	9782:  uint16(sym_comment),
	9783:  uint16(657),
	9784:  uint16(1),
	9785:  uint16(anon_sym_DOLLAR),
	9786:  uint16(984),
	9787:  uint16(1),
	9788:  uint16(anon_sym_SEMI),
	9789:  uint16(195),
	9790:  uint16(1),
	9791:  uint16(aux_sym_annotation_repeat1),
	9792:  uint16(4),
	9793:  uint16(3),
	9794:  uint16(1),
	9795:  uint16(sym_comment),
	9796:  uint16(657),
	9797:  uint16(1),
	9798:  uint16(anon_sym_DOLLAR),
	9799:  uint16(1183),
	9800:  uint16(1),
	9801:  uint16(anon_sym_SEMI),
	9802:  uint16(418),
	9803:  uint16(1),
	9804:  uint16(aux_sym_annotation_repeat1),
	9805:  uint16(4),
	9806:  uint16(3),
	9807:  uint16(1),
	9808:  uint16(sym_comment),
	9809:  uint16(1185),
	9810:  uint16(1),
	9811:  uint16(anon_sym_SEMI),
	9812:  uint16(1187),
	9813:  uint16(1),
	9814:  uint16(anon_sym_DOT),
	9815:  uint16(428),
	9816:  uint16(1),
	9817:  uint16(aux_sym_import_using_repeat1),
	9818:  uint16(4),
	9819:  uint16(3),
	9820:  uint16(1),
	9821:  uint16(sym_comment),
	9822:  uint16(503),
	9823:  uint16(1),
	9824:  uint16(anon_sym_LPAREN),
	9825:  uint16(1190),
	9826:  uint16(1),
	9827:  uint16(anon_sym_RPAREN),
	9828:  uint16(494),
	9829:  uint16(1),
	9830:  uint16(sym_generics),
	9831:  uint16(4),
	9832:  uint16(3),
	9833:  uint16(1),
	9834:  uint16(sym_comment),
	9835:  uint16(657),
	9836:  uint16(1),
	9837:  uint16(anon_sym_DOLLAR),
	9838:  uint16(1183),
	9839:  uint16(1),
	9840:  uint16(anon_sym_SEMI),
	9841:  uint16(195),
	9842:  uint16(1),
	9843:  uint16(aux_sym_annotation_repeat1),
	9844:  uint16(2),
	9845:  uint16(3),
	9846:  uint16(1),
	9847:  uint16(sym_comment),
	9848:  uint16(1192),
	9849:  uint16(3),
	9850:  uint16(anon_sym_SEMI),
	9851:  uint16(anon_sym_DOLLAR),
	9852:  uint16(anon_sym_DASH_GT),
	9853:  uint16(4),
	9854:  uint16(3),
	9855:  uint16(1),
	9856:  uint16(sym_comment),
	9857:  uint16(657),
	9858:  uint16(1),
	9859:  uint16(anon_sym_DOLLAR),
	9860:  uint16(1194),
	9861:  uint16(1),
	9862:  uint16(anon_sym_LBRACE),
	9863:  uint16(195),
	9864:  uint16(1),
	9865:  uint16(aux_sym_annotation_repeat1),
	9866:  uint16(4),
	9867:  uint16(3),
	9868:  uint16(1),
	9869:  uint16(sym_comment),
	9870:  uint16(45),
	9871:  uint16(1),
	9872:  uint16(anon_sym_DQUOTE),
	9873:  uint16(47),
	9874:  uint16(1),
	9875:  uint16(anon_sym_SQUOTE),
	9876:  uint16(479),
	9877:  uint16(1),
	9878:  uint16(sym_string),
	9879:  uint16(4),
	9880:  uint16(3),
	9881:  uint16(1),
	9882:  uint16(sym_comment),
	9883:  uint16(1116),
	9884:  uint16(1),
	9885:  uint16(sym_identifier),
	9886:  uint16(1196),
	9887:  uint16(1),
	9888:  uint16(anon_sym_RPAREN),
	9889:  uint16(410),
	9890:  uint16(1),
	9891:  uint16(aux_sym_struct_shorthand_repeat1),
	9892:  uint16(4),
	9893:  uint16(3),
	9894:  uint16(1),
	9895:  uint16(sym_comment),
	9896:  uint16(657),
	9897:  uint16(1),
	9898:  uint16(anon_sym_DOLLAR),
	9899:  uint16(1198),
	9900:  uint16(1),
	9901:  uint16(anon_sym_SEMI),
	9902:  uint16(386),
	9903:  uint16(1),
	9904:  uint16(aux_sym_annotation_repeat1),
	9905:  uint16(4),
	9906:  uint16(3),
	9907:  uint16(1),
	9908:  uint16(sym_comment),
	9909:  uint16(657),
	9910:  uint16(1),
	9911:  uint16(anon_sym_DOLLAR),
	9912:  uint16(1200),
	9913:  uint16(1),
	9914:  uint16(anon_sym_SEMI),
	9915:  uint16(395),
	9916:  uint16(1),
	9917:  uint16(aux_sym_annotation_repeat1),
	9918:  uint16(4),
	9919:  uint16(3),
	9920:  uint16(1),
	9921:  uint16(sym_comment),
	9922:  uint16(657),
	9923:  uint16(1),
	9924:  uint16(anon_sym_DOLLAR),
	9925:  uint16(1202),
	9926:  uint16(1),
	9927:  uint16(anon_sym_SEMI),
	9928:  uint16(195),
	9929:  uint16(1),
	9930:  uint16(aux_sym_annotation_repeat1),
	9931:  uint16(4),
	9932:  uint16(3),
	9933:  uint16(1),
	9934:  uint16(sym_comment),
	9935:  uint16(1048),
	9936:  uint16(1),
	9937:  uint16(anon_sym_DOT),
	9938:  uint16(1204),
	9939:  uint16(1),
	9940:  uint16(anon_sym_SEMI),
	9941:  uint16(428),
	9942:  uint16(1),
	9943:  uint16(aux_sym_import_using_repeat1),
	9944:  uint16(4),
	9945:  uint16(3),
	9946:  uint16(1),
	9947:  uint16(sym_comment),
	9948:  uint16(657),
	9949:  uint16(1),
	9950:  uint16(anon_sym_DOLLAR),
	9951:  uint16(1202),
	9952:  uint16(1),
	9953:  uint16(anon_sym_SEMI),
	9954:  uint16(397),
	9955:  uint16(1),
	9956:  uint16(aux_sym_annotation_repeat1),
	9957:  uint16(4),
	9958:  uint16(3),
	9959:  uint16(1),
	9960:  uint16(sym_comment),
	9961:  uint16(1132),
	9962:  uint16(1),
	9963:  uint16(anon_sym_COMMA),
	9964:  uint16(1206),
	9965:  uint16(1),
	9966:  uint16(anon_sym_RPAREN),
	9967:  uint16(398),
	9968:  uint16(1),
	9969:  uint16(aux_sym_annotation_targets_repeat1),
	9970:  uint16(2),
	9971:  uint16(3),
	9972:  uint16(1),
	9973:  uint16(sym_comment),
	9974:  uint16(1208),
	9975:  uint16(3),
	9976:  uint16(anon_sym_SEMI),
	9977:  uint16(anon_sym_DOLLAR),
	9978:  uint16(anon_sym_COLON),
	9979:  uint16(3),
	9980:  uint16(3),
	9981:  uint16(1),
	9982:  uint16(sym_comment),
	9983:  uint16(1210),
	9984:  uint16(1),
	9985:  uint16(anon_sym_SEMI),
	9986:  uint16(1212),
	9987:  uint16(1),
	9988:  uint16(anon_sym_DOT),
	9989:  uint16(3),
	9990:  uint16(3),
	9991:  uint16(1),
	9992:  uint16(sym_comment),
	9993:  uint16(1214),
	9994:  uint16(1),
	9995:  uint16(aux_sym_top_level_annotation_body_token1),
	9996:  uint16(521),
	9997:  uint16(1),
	9998:  uint16(sym_top_level_annotation_body),
	9999:  uint16(3),
	10000: uint16(3),
	10001: uint16(1),
	10002: uint16(sym_comment),
	10003: uint16(1216),
	10004: uint16(1),
	10005: uint16(sym__identifier_no_period),
	10006: uint16(445),
	10007: uint16(1),
	10008: uint16(aux_sym__internal_const_identifier_repeat1),
	10009: uint16(3),
	10010: uint16(3),
	10011: uint16(1),
	10012: uint16(sym_comment),
	10013: uint16(1218),
	10014: uint16(1),
	10015: uint16(sym__identifier_no_period),
	10016: uint16(445),
	10017: uint16(1),
	10018: uint16(aux_sym__internal_const_identifier_repeat1),
	10019: uint16(3),
	10020: uint16(3),
	10021: uint16(1),
	10022: uint16(sym_comment),
	10023: uint16(1221),
	10024: uint16(1),
	10025: uint16(sym_identifier),
	10026: uint16(341),
	10027: uint16(1),
	10028: uint16(sym__annotation_definition_identifier),
	10029: uint16(2),
	10030: uint16(3),
	10031: uint16(1),
	10032: uint16(sym_comment),
	10033: uint16(1223),
	10034: uint16(2),
	10035: uint16(anon_sym_RPAREN),
	10036: uint16(anon_sym_COMMA),
	10037: uint16(2),
	10038: uint16(3),
	10039: uint16(1),
	10040: uint16(sym_comment),
	10041: uint16(1225),
	10042: uint16(2),
	10043: uint16(anon_sym_RPAREN),
	10044: uint16(anon_sym_COMMA),
	10045: uint16(2),
	10046: uint16(3),
	10047: uint16(1),
	10048: uint16(sym_comment),
	10049: uint16(1227),
	10050: uint16(2),
	10051: uint16(anon_sym_RPAREN),
	10052: uint16(anon_sym_COMMA),
	10053: uint16(3),
	10054: uint16(3),
	10055: uint16(1),
	10056: uint16(sym_comment),
	10057: uint16(1229),
	10058: uint16(1),
	10059: uint16(anon_sym_SEMI),
	10060: uint16(1231),
	10061: uint16(1),
	10062: uint16(anon_sym_EQ),
	10063: uint16(2),
	10064: uint16(3),
	10065: uint16(1),
	10066: uint16(sym_comment),
	10067: uint16(1233),
	10068: uint16(2),
	10069: uint16(anon_sym_RPAREN),
	10070: uint16(sym_identifier),
	10071: uint16(2),
	10072: uint16(3),
	10073: uint16(1),
	10074: uint16(sym_comment),
	10075: uint16(1235),
	10076: uint16(2),
	10077: uint16(anon_sym_RBRACE),
	10078: uint16(sym_identifier),
	10079: uint16(3),
	10080: uint16(3),
	10081: uint16(1),
	10082: uint16(sym_comment),
	10083: uint16(1237),
	10084: uint16(1),
	10085: uint16(anon_sym_SEMI),
	10086: uint16(1239),
	10087: uint16(1),
	10088: uint16(anon_sym_DOT),
	10089: uint16(3),
	10090: uint16(3),
	10091: uint16(1),
	10092: uint16(sym_comment),
	10093: uint16(1241),
	10094: uint16(1),
	10095: uint16(anon_sym_SEMI),
	10096: uint16(1243),
	10097: uint16(1),
	10098: uint16(anon_sym_LPAREN),
	10099: uint16(3),
	10100: uint16(3),
	10101: uint16(1),
	10102: uint16(sym_comment),
	10103: uint16(1245),
	10104: uint16(1),
	10105: uint16(sym_unique_id),
	10106: uint16(1247),
	10107: uint16(1),
	10108: uint16(anon_sym_COLON),
	10109: uint16(2),
	10110: uint16(3),
	10111: uint16(1),
	10112: uint16(sym_comment),
	10113: uint16(1249),
	10114: uint16(2),
	10115: uint16(anon_sym_RPAREN),
	10116: uint16(anon_sym_COMMA),
	10117: uint16(3),
	10118: uint16(3),
	10119: uint16(1),
	10120: uint16(sym_comment),
	10121: uint16(1251),
	10122: uint16(1),
	10123: uint16(anon_sym_group),
	10124: uint16(1253),
	10125: uint16(1),
	10126: uint16(anon_sym_union),
	10127: uint16(2),
	10128: uint16(3),
	10129: uint16(1),
	10130: uint16(sym_comment),
	10131: uint16(1255),
	10132: uint16(2),
	10133: uint16(anon_sym_SEMI),
	10134: uint16(anon_sym_DOT),
	10135: uint16(3),
	10136: uint16(3),
	10137: uint16(1),
	10138: uint16(sym_comment),
	10139: uint16(1257),
	10140: uint16(1),
	10141: uint16(sym_identifier),
	10142: uint16(1259),
	10143: uint16(1),
	10144: uint16(anon_sym_import),
	10145: uint16(2),
	10146: uint16(3),
	10147: uint16(1),
	10148: uint16(sym_comment),
	10149: uint16(1261),
	10150: uint16(2),
	10151: uint16(anon_sym_SEMI),
	10152: uint16(anon_sym_DOLLAR),
	10153: uint16(3),
	10154: uint16(3),
	10155: uint16(1),
	10156: uint16(sym_comment),
	10157: uint16(1042),
	10158: uint16(1),
	10159: uint16(anon_sym_LPAREN),
	10160: uint16(276),
	10161: uint16(1),
	10162: uint16(sym_annotation_targets),
	10163: uint16(3),
	10164: uint16(3),
	10165: uint16(1),
	10166: uint16(sym_comment),
	10167: uint16(900),
	10168: uint16(1),
	10169: uint16(sym_identifier),
	10170: uint16(469),
	10171: uint16(1),
	10172: uint16(sym_parameter),
	10173: uint16(2),
	10174: uint16(3),
	10175: uint16(1),
	10176: uint16(sym_comment),
	10177: uint16(1263),
	10178: uint16(2),
	10179: uint16(anon_sym_SEMI),
	10180: uint16(anon_sym_DOLLAR),
	10181: uint16(2),
	10182: uint16(3),
	10183: uint16(1),
	10184: uint16(sym_comment),
	10185: uint16(1265),
	10186: uint16(2),
	10187: uint16(anon_sym_LPAREN),
	10188: uint16(sym_identifier),
	10189: uint16(2),
	10190: uint16(3),
	10191: uint16(1),
	10192: uint16(sym_comment),
	10193: uint16(872),
	10194: uint16(2),
	10195: uint16(anon_sym_SEMI),
	10196: uint16(anon_sym_DOLLAR),
	10197: uint16(3),
	10198: uint16(3),
	10199: uint16(1),
	10200: uint16(sym_comment),
	10201: uint16(1267),
	10202: uint16(1),
	10203: uint16(sym__identifier_no_period),
	10204: uint16(445),
	10205: uint16(1),
	10206: uint16(aux_sym__internal_const_identifier_repeat1),
	10207: uint16(3),
	10208: uint16(3),
	10209: uint16(1),
	10210: uint16(sym_comment),
	10211: uint16(1269),
	10212: uint16(1),
	10213: uint16(sym__normal_version),
	10214: uint16(1271),
	10215: uint16(1),
	10216: uint16(aux_sym__inline_version_token1),
	10217: uint16(2),
	10218: uint16(3),
	10219: uint16(1),
	10220: uint16(sym_comment),
	10221: uint16(1273),
	10222: uint16(2),
	10223: uint16(anon_sym_SEMI),
	10224: uint16(anon_sym_DOLLAR),
	10225: uint16(2),
	10226: uint16(3),
	10227: uint16(1),
	10228: uint16(sym_comment),
	10229: uint16(1168),
	10230: uint16(2),
	10231: uint16(anon_sym_RPAREN),
	10232: uint16(anon_sym_COMMA),
	10233: uint16(2),
	10234: uint16(3),
	10235: uint16(1),
	10236: uint16(sym_comment),
	10237: uint16(1275),
	10238: uint16(2),
	10239: uint16(anon_sym_RBRACE),
	10240: uint16(sym_identifier),
	10241: uint16(2),
	10242: uint16(3),
	10243: uint16(1),
	10244: uint16(sym_comment),
	10245: uint16(1136),
	10246: uint16(2),
	10247: uint16(anon_sym_RPAREN),
	10248: uint16(anon_sym_COMMA),
	10249: uint16(2),
	10250: uint16(3),
	10251: uint16(1),
	10252: uint16(sym_comment),
	10253: uint16(1277),
	10254: uint16(2),
	10255: uint16(anon_sym_RPAREN),
	10256: uint16(anon_sym_COMMA),
	10257: uint16(3),
	10258: uint16(3),
	10259: uint16(1),
	10260: uint16(sym_comment),
	10261: uint16(649),
	10262: uint16(1),
	10263: uint16(anon_sym_DOT),
	10264: uint16(651),
	10265: uint16(1),
	10266: uint16(sym__identifier_no_period),
	10267: uint16(2),
	10268: uint16(3),
	10269: uint16(1),
	10270: uint16(sym_comment),
	10271: uint16(1279),
	10272: uint16(2),
	10273: uint16(anon_sym_RPAREN),
	10274: uint16(anon_sym_COMMA),
	10275: uint16(2),
	10276: uint16(3),
	10277: uint16(1),
	10278: uint16(sym_comment),
	10279: uint16(1281),
	10280: uint16(1),
	10281: uint16(sym_identifier),
	10282: uint16(2),
	10283: uint16(3),
	10284: uint16(1),
	10285: uint16(sym_comment),
	10286: uint16(807),
	10287: uint16(1),
	10288: uint16(anon_sym_LBRACE),
	10289: uint16(2),
	10290: uint16(3),
	10291: uint16(1),
	10292: uint16(sym_comment),
	10293: uint16(1283),
	10294: uint16(1),
	10295: uint16(anon_sym_LPAREN),
	10296: uint16(2),
	10297: uint16(3),
	10298: uint16(1),
	10299: uint16(sym_comment),
	10300: uint16(1285),
	10301: uint16(1),
	10302: uint16(anon_sym_RPAREN),
	10303: uint16(2),
	10304: uint16(3),
	10305: uint16(1),
	10306: uint16(sym_comment),
	10307: uint16(1287),
	10308: uint16(1),
	10309: uint16(anon_sym_RPAREN),
	10310: uint16(2),
	10311: uint16(3),
	10312: uint16(1),
	10313: uint16(sym_comment),
	10314: uint16(1289),
	10315: uint16(1),
	10316: uint16(anon_sym_LPAREN),
	10317: uint16(2),
	10318: uint16(3),
	10319: uint16(1),
	10320: uint16(sym_comment),
	10321: uint16(1291),
	10322: uint16(1),
	10323: uint16(sym_identifier),
	10324: uint16(2),
	10325: uint16(3),
	10326: uint16(1),
	10327: uint16(sym_comment),
	10328: uint16(1293),
	10329: uint16(1),
	10330: uint16(anon_sym_EQ),
	10331: uint16(2),
	10332: uint16(3),
	10333: uint16(1),
	10334: uint16(sym_comment),
	10335: uint16(1295),
	10336: uint16(1),
	10337: uint16(anon_sym_LPAREN),
	10338: uint16(2),
	10339: uint16(3),
	10340: uint16(1),
	10341: uint16(sym_comment),
	10342: uint16(1297),
	10343: uint16(1),
	10344: uint16(anon_sym_SEMI),
	10345: uint16(2),
	10346: uint16(3),
	10347: uint16(1),
	10348: uint16(sym_comment),
	10349: uint16(1299),
	10350: uint16(1),
	10351: uint16(anon_sym_RPAREN),
	10352: uint16(2),
	10353: uint16(3),
	10354: uint16(1),
	10355: uint16(sym_comment),
	10356: uint16(61),
	10357: uint16(1),
	10358: uint16(anon_sym_RPAREN),
	10359: uint16(2),
	10360: uint16(3),
	10361: uint16(1),
	10362: uint16(sym_comment),
	10363: uint16(1301),
	10364: uint16(1),
	10365: uint16(anon_sym_EQ),
	10366: uint16(2),
	10367: uint16(3),
	10368: uint16(1),
	10369: uint16(sym_comment),
	10370: uint16(59),
	10371: uint16(1),
	10372: uint16(anon_sym_RPAREN),
	10373: uint16(2),
	10374: uint16(3),
	10375: uint16(1),
	10376: uint16(sym_comment),
	10377: uint16(1303),
	10378: uint16(1),
	10379: uint16(aux_sym_data_token1),
	10380: uint16(2),
	10381: uint16(3),
	10382: uint16(1),
	10383: uint16(sym_comment),
	10384: uint16(1305),
	10385: uint16(1),
	10386: uint16(anon_sym_RPAREN),
	10387: uint16(2),
	10388: uint16(3),
	10389: uint16(1),
	10390: uint16(sym_comment),
	10391: uint16(1307),
	10392: uint16(1),
	10393: uint16(anon_sym_SEMI),
	10394: uint16(2),
	10395: uint16(3),
	10396: uint16(1),
	10397: uint16(sym_comment),
	10398: uint16(1309),
	10399: uint16(1),
	10400: uint16(anon_sym_LPAREN),
	10401: uint16(2),
	10402: uint16(3),
	10403: uint16(1),
	10404: uint16(sym_comment),
	10405: uint16(1311),
	10406: uint16(1),
	10407: uint16(sym_identifier),
	10408: uint16(2),
	10409: uint16(3),
	10410: uint16(1),
	10411: uint16(sym_comment),
	10412: uint16(1313),
	10413: uint16(1),
	10414: uint16(anon_sym_RPAREN),
	10415: uint16(2),
	10416: uint16(3),
	10417: uint16(1),
	10418: uint16(sym_comment),
	10419: uint16(1315),
	10420: uint16(1),
	10421: uint16(anon_sym_COLON),
	10422: uint16(2),
	10423: uint16(3),
	10424: uint16(1),
	10425: uint16(sym_comment),
	10426: uint16(1317),
	10427: uint16(1),
	10428: uint16(anon_sym_EQ),
	10429: uint16(2),
	10430: uint16(3),
	10431: uint16(1),
	10432: uint16(sym_comment),
	10433: uint16(1319),
	10434: uint16(1),
	10435: uint16(anon_sym_SEMI),
	10436: uint16(2),
	10437: uint16(3),
	10438: uint16(1),
	10439: uint16(sym_comment),
	10440: uint16(1321),
	10441: uint16(1),
	10442: uint16(anon_sym_SEMI),
	10443: uint16(2),
	10444: uint16(3),
	10445: uint16(1),
	10446: uint16(sym_comment),
	10447: uint16(1323),
	10448: uint16(1),
	10449: uint16(sym_identifier),
	10450: uint16(2),
	10451: uint16(3),
	10452: uint16(1),
	10453: uint16(sym_comment),
	10454: uint16(1325),
	10455: uint16(1),
	10456: uint16(anon_sym_COLON),
	10457: uint16(2),
	10458: uint16(3),
	10459: uint16(1),
	10460: uint16(sym_comment),
	10461: uint16(1327),
	10462: uint16(1),
	10463: uint16(anon_sym_LPAREN),
	10464: uint16(2),
	10465: uint16(3),
	10466: uint16(1),
	10467: uint16(sym_comment),
	10468: uint16(1329),
	10469: uint16(1),
	10470: uint16(anon_sym_SEMI),
	10471: uint16(2),
	10472: uint16(3),
	10473: uint16(1),
	10474: uint16(sym_comment),
	10475: uint16(1331),
	10476: uint16(1),
	10477: uint16(anon_sym_RPAREN),
	10478: uint16(2),
	10479: uint16(3),
	10480: uint16(1),
	10481: uint16(sym_comment),
	10482: uint16(1333),
	10483: uint16(1),
	10484: uint16(sym_identifier),
	10485: uint16(2),
	10486: uint16(3),
	10487: uint16(1),
	10488: uint16(sym_comment),
	10489: uint16(1335),
	10490: uint16(1),
	10491: uint16(sym_identifier),
	10492: uint16(2),
	10493: uint16(3),
	10494: uint16(1),
	10495: uint16(sym_comment),
	10496: uint16(1337),
	10497: uint16(1),
	10498: uint16(anon_sym_RPAREN),
	10499: uint16(2),
	10500: uint16(3),
	10501: uint16(1),
	10502: uint16(sym_comment),
	10503: uint16(1339),
	10504: uint16(1),
	10505: uint16(anon_sym_COLON),
	10506: uint16(2),
	10507: uint16(3),
	10508: uint16(1),
	10509: uint16(sym_comment),
	10510: uint16(1341),
	10511: uint16(1),
	10512: uint16(anon_sym_RPAREN),
	10513: uint16(2),
	10514: uint16(3),
	10515: uint16(1),
	10516: uint16(sym_comment),
	10517: uint16(1343),
	10518: uint16(1),
	10519: uint16(anon_sym_RPAREN),
	10520: uint16(2),
	10521: uint16(3),
	10522: uint16(1),
	10523: uint16(sym_comment),
	10524: uint16(1345),
	10525: uint16(1),
	10526: uint16(sym_identifier),
	10527: uint16(2),
	10528: uint16(3),
	10529: uint16(1),
	10530: uint16(sym_comment),
	10531: uint16(69),
	10532: uint16(1),
	10533: uint16(anon_sym_RPAREN),
	10534: uint16(2),
	10535: uint16(3),
	10536: uint16(1),
	10537: uint16(sym_comment),
	10538: uint16(1347),
	10539: uint16(1),
	10540: uint16(anon_sym_namespace),
	10541: uint16(2),
	10542: uint16(3),
	10543: uint16(1),
	10544: uint16(sym_comment),
	10545: uint16(1349),
	10546: uint16(1),
	10547: uint16(anon_sym_RPAREN),
	10548: uint16(2),
	10549: uint16(3),
	10550: uint16(1),
	10551: uint16(sym_comment),
	10552: uint16(1351),
	10553: uint16(1),
	10554: uint16(sym_identifier),
	10555: uint16(2),
	10556: uint16(3),
	10557: uint16(1),
	10558: uint16(sym_comment),
	10559: uint16(1353),
	10560: uint16(1),
	10561: uint16(anon_sym_EQ),
	10562: uint16(2),
	10563: uint16(3),
	10564: uint16(1),
	10565: uint16(sym_comment),
	10566: uint16(1355),
	10567: uint16(1),
	10568: uint16(anon_sym_namespace),
	10569: uint16(2),
	10570: uint16(3),
	10571: uint16(1),
	10572: uint16(sym_comment),
	10573: uint16(1357),
	10574: uint16(1),
	10575: uint16(anon_sym_SEMI),
	10576: uint16(2),
	10577: uint16(3),
	10578: uint16(1),
	10579: uint16(sym_comment),
	10580: uint16(1359),
	10581: uint16(1),
	10582: uint16(sym_identifier),
	10583: uint16(2),
	10584: uint16(3),
	10585: uint16(1),
	10586: uint16(sym_comment),
	10587: uint16(1361),
	10588: uint16(1),
	10589: uint16(anon_sym_EQ),
	10590: uint16(2),
	10591: uint16(3),
	10592: uint16(1),
	10593: uint16(sym_comment),
	10594: uint16(1363),
	10595: uint16(1),
	10596: uint16(anon_sym_EQ),
	10597: uint16(2),
	10598: uint16(3),
	10599: uint16(1),
	10600: uint16(sym_comment),
	10601: uint16(1365),
	10602: uint16(1),
	10603: uint16(anon_sym_SEMI),
	10604: uint16(2),
	10605: uint16(3),
	10606: uint16(1),
	10607: uint16(sym_comment),
	10608: uint16(1367),
	10609: uint16(1),
	10610: uint16(anon_sym_SEMI),
	10611: uint16(2),
	10612: uint16(3),
	10613: uint16(1),
	10614: uint16(sym_comment),
	10615: uint16(1369),
	10616: uint16(1),
	10617: uint16(anon_sym_SEMI),
	10618: uint16(2),
	10619: uint16(3),
	10620: uint16(1),
	10621: uint16(sym_comment),
	10622: uint16(1371),
	10623: uint16(1),
	10624: uint16(anon_sym_SEMI),
	10625: uint16(2),
	10626: uint16(3),
	10627: uint16(1),
	10628: uint16(sym_comment),
	10629: uint16(1373),
	10630: uint16(1),
	10631: uint16(anon_sym_SEMI),
	10632: uint16(2),
	10633: uint16(3),
	10634: uint16(1),
	10635: uint16(sym_comment),
	10636: uint16(1375),
	10637: uint16(1),
	10638: uint16(anon_sym_RPAREN),
	10639: uint16(2),
	10640: uint16(3),
	10641: uint16(1),
	10642: uint16(sym_comment),
	10643: uint16(1377),
	10644: uint16(1),
	10645: uint16(anon_sym_EQ),
	10646: uint16(2),
	10647: uint16(3),
	10648: uint16(1),
	10649: uint16(sym_comment),
	10650: uint16(63),
	10651: uint16(1),
	10652: uint16(anon_sym_RPAREN),
	10653: uint16(2),
	10654: uint16(3),
	10655: uint16(1),
	10656: uint16(sym_comment),
	10657: uint16(1379),
	10658: uint16(1),
	10659: uint16(anon_sym_COLON),
	10660: uint16(2),
	10661: uint16(3),
	10662: uint16(1),
	10663: uint16(sym_comment),
	10664: uint16(1381),
	10665: uint16(1),
	10666: uint16(anon_sym_SEMI),
	10667: uint16(2),
	10668: uint16(3),
	10669: uint16(1),
	10670: uint16(sym_comment),
	10671: uint16(1383),
	10672: uint16(1),
	10673: uint16(anon_sym_LPAREN),
	10674: uint16(2),
	10675: uint16(3),
	10676: uint16(1),
	10677: uint16(sym_comment),
	10678: uint16(1385),
	10679: uint16(1),
	10680: uint16(anon_sym_EQ),
	10681: uint16(2),
	10682: uint16(3),
	10683: uint16(1),
	10684: uint16(sym_comment),
	10685: uint16(1387),
	10686: uint16(1),
	10687: uint16(anon_sym_RPAREN),
	10688: uint16(2),
	10689: uint16(3),
	10690: uint16(1),
	10691: uint16(sym_comment),
	10692: uint16(1389),
	10693: uint16(1),
	10694: uint16(anon_sym_EQ),
	10695: uint16(2),
	10696: uint16(3),
	10697: uint16(1),
	10698: uint16(sym_comment),
	10699: uint16(1391),
	10700: uint16(1),
	10701: uint16(anon_sym_RPAREN),
	10702: uint16(2),
	10703: uint16(3),
	10704: uint16(1),
	10705: uint16(sym_comment),
	10706: uint16(1393),
	10707: uint16(1),
	10708: uint16(sym_identifier),
	10709: uint16(2),
	10710: uint16(3),
	10711: uint16(1),
	10712: uint16(sym_comment),
	10713: uint16(1395),
	10714: uint16(1),
	10715: uint16(anon_sym_RBRACK),
	10716: uint16(2),
	10717: uint16(3),
	10718: uint16(1),
	10719: uint16(sym_comment),
	10720: uint16(1397),
	10721: uint16(1),
	10722: uint16(anon_sym_COLON),
	10723: uint16(2),
	10724: uint16(3),
	10725: uint16(1),
	10726: uint16(sym_comment),
	10727: uint16(1399),
	10728: uint16(1),
	10729: uint16(anon_sym_COLON),
	10730: uint16(2),
	10731: uint16(3),
	10732: uint16(1),
	10733: uint16(sym_comment),
	10734: uint16(1401),
	10735: uint16(1),
	10737: uint16(2),
	10738: uint16(3),
	10739: uint16(1),
	10740: uint16(sym_comment),
	10741: uint16(1403),
	10742: uint16(1),
	10743: uint16(anon_sym_RPAREN),
	10744: uint16(2),
	10745: uint16(3),
	10746: uint16(1),
	10747: uint16(sym_comment),
	10748: uint16(1405),
	10749: uint16(1),
	10750: uint16(sym_identifier),
	10751: uint16(2),
	10752: uint16(3),
	10753: uint16(1),
	10754: uint16(sym_comment),
	10755: uint16(1407),
	10756: uint16(1),
	10757: uint16(sym_identifier),
	10758: uint16(2),
	10759: uint16(3),
	10760: uint16(1),
	10761: uint16(sym_comment),
	10762: uint16(1409),
	10763: uint16(1),
	10764: uint16(anon_sym_SEMI),
	10765: uint16(2),
	10766: uint16(3),
	10767: uint16(1),
	10768: uint16(sym_comment),
	10769: uint16(1411),
	10770: uint16(1),
	10771: uint16(sym_identifier),
	10772: uint16(2),
	10773: uint16(3),
	10774: uint16(1),
	10775: uint16(sym_comment),
	10776: uint16(1413),
	10777: uint16(1),
	10778: uint16(sym_identifier),
	10779: uint16(2),
	10780: uint16(3),
	10781: uint16(1),
	10782: uint16(sym_comment),
	10783: uint16(1415),
	10784: uint16(1),
	10785: uint16(sym_identifier),
	10786: uint16(2),
	10787: uint16(3),
	10788: uint16(1),
	10789: uint16(sym_comment),
	10790: uint16(1417),
	10791: uint16(1),
	10792: uint16(sym__identifier_no_period),
	10793: uint16(2),
	10794: uint16(3),
	10795: uint16(1),
	10796: uint16(sym_comment),
	10797: uint16(1419),
	10798: uint16(1),
	10799: uint16(anon_sym_SEMI),
	10800: uint16(2),
	10801: uint16(3),
	10802: uint16(1),
	10803: uint16(sym_comment),
	10804: uint16(1421),
	10805: uint16(1),
	10806: uint16(anon_sym_COLON),
	10807: uint16(2),
	10808: uint16(3),
	10809: uint16(1),
	10810: uint16(sym_comment),
	10811: uint16(1423),
	10812: uint16(1),
	10813: uint16(sym_identifier),
	10814: uint16(2),
	10815: uint16(3),
	10816: uint16(1),
	10817: uint16(sym_comment),
	10818: uint16(1425),
	10819: uint16(1),
	10820: uint16(anon_sym_EQ),
}

var ts_small_parse_table_map = [551]uint32_t{
	1:   uint32(102),
	2:   uint32(176),
	3:   uint32(247),
	4:   uint32(318),
	5:   uint32(389),
	6:   uint32(460),
	7:   uint32(531),
	8:   uint32(602),
	9:   uint32(671),
	10:  uint32(742),
	11:  uint32(806),
	12:  uint32(874),
	13:  uint32(938),
	14:  uint32(1003),
	15:  uint32(1066),
	16:  uint32(1128),
	17:  uint32(1190),
	18:  uint32(1252),
	19:  uint32(1314),
	20:  uint32(1373),
	21:  uint32(1432),
	22:  uint32(1491),
	23:  uint32(1550),
	24:  uint32(1609),
	25:  uint32(1668),
	26:  uint32(1727),
	27:  uint32(1786),
	28:  uint32(1845),
	29:  uint32(1904),
	30:  uint32(1963),
	31:  uint32(2022),
	32:  uint32(2081),
	33:  uint32(2140),
	34:  uint32(2199),
	35:  uint32(2258),
	36:  uint32(2317),
	37:  uint32(2376),
	38:  uint32(2435),
	39:  uint32(2494),
	40:  uint32(2549),
	41:  uint32(2604),
	42:  uint32(2659),
	43:  uint32(2714),
	44:  uint32(2769),
	45:  uint32(2824),
	46:  uint32(2879),
	47:  uint32(2934),
	48:  uint32(2989),
	49:  uint32(3030),
	50:  uint32(3082),
	51:  uint32(3134),
	52:  uint32(3172),
	53:  uint32(3210),
	54:  uint32(3248),
	55:  uint32(3299),
	56:  uint32(3334),
	57:  uint32(3385),
	58:  uint32(3436),
	59:  uint32(3471),
	60:  uint32(3506),
	61:  uint32(3557),
	62:  uint32(3592),
	63:  uint32(3627),
	64:  uint32(3662),
	65:  uint32(3713),
	66:  uint32(3748),
	67:  uint32(3783),
	68:  uint32(3807),
	69:  uint32(3831),
	70:  uint32(3852),
	71:  uint32(3873),
	72:  uint32(3894),
	73:  uint32(3915),
	74:  uint32(3936),
	75:  uint32(3957),
	76:  uint32(3978),
	77:  uint32(3999),
	78:  uint32(4020),
	79:  uint32(4041),
	80:  uint32(4062),
	81:  uint32(4083),
	82:  uint32(4104),
	83:  uint32(4125),
	84:  uint32(4146),
	85:  uint32(4167),
	86:  uint32(4188),
	87:  uint32(4209),
	88:  uint32(4230),
	89:  uint32(4251),
	90:  uint32(4272),
	91:  uint32(4293),
	92:  uint32(4314),
	93:  uint32(4335),
	94:  uint32(4356),
	95:  uint32(4377),
	96:  uint32(4398),
	97:  uint32(4419),
	98:  uint32(4440),
	99:  uint32(4461),
	100: uint32(4482),
	101: uint32(4503),
	102: uint32(4524),
	103: uint32(4545),
	104: uint32(4566),
	105: uint32(4587),
	106: uint32(4611),
	107: uint32(4635),
	108: uint32(4659),
	109: uint32(4677),
	110: uint32(4706),
	111: uint32(4735),
	112: uint32(4764),
	113: uint32(4793),
	114: uint32(4822),
	115: uint32(4849),
	116: uint32(4878),
	117: uint32(4907),
	118: uint32(4936),
	119: uint32(4965),
	120: uint32(4994),
	121: uint32(5023),
	122: uint32(5052),
	123: uint32(5081),
	124: uint32(5110),
	125: uint32(5139),
	126: uint32(5168),
	127: uint32(5197),
	128: uint32(5226),
	129: uint32(5255),
	130: uint32(5284),
	131: uint32(5313),
	132: uint32(5342),
	133: uint32(5371),
	134: uint32(5400),
	135: uint32(5429),
	136: uint32(5458),
	137: uint32(5487),
	138: uint32(5515),
	139: uint32(5533),
	140: uint32(5561),
	141: uint32(5583),
	142: uint32(5601),
	143: uint32(5619),
	144: uint32(5647),
	145: uint32(5669),
	146: uint32(5691),
	147: uint32(5719),
	148: uint32(5743),
	149: uint32(5771),
	150: uint32(5795),
	151: uint32(5819),
	152: uint32(5847),
	153: uint32(5871),
	154: uint32(5895),
	155: uint32(5923),
	156: uint32(5951),
	157: uint32(5975),
	158: uint32(5993),
	159: uint32(6011),
	160: uint32(6029),
	161: uint32(6047),
	162: uint32(6065),
	163: uint32(6093),
	164: uint32(6110),
	165: uint32(6127),
	166: uint32(6146),
	167: uint32(6165),
	168: uint32(6182),
	169: uint32(6199),
	170: uint32(6216),
	171: uint32(6231),
	172: uint32(6248),
	173: uint32(6265),
	174: uint32(6282),
	175: uint32(6297),
	176: uint32(6314),
	177: uint32(6331),
	178: uint32(6348),
	179: uint32(6365),
	180: uint32(6382),
	181: uint32(6399),
	182: uint32(6416),
	183: uint32(6433),
	184: uint32(6450),
	185: uint32(6467),
	186: uint32(6484),
	187: uint32(6502),
	188: uint32(6520),
	189: uint32(6538),
	190: uint32(6552),
	191: uint32(6570),
	192: uint32(6588),
	193: uint32(6606),
	194: uint32(6623),
	195: uint32(6640),
	196: uint32(6653),
	197: uint32(6678),
	198: uint32(6701),
	199: uint32(6718),
	200: uint32(6730),
	201: uint32(6742),
	202: uint32(6760),
	203: uint32(6772),
	204: uint32(6784),
	205: uint32(6796),
	206: uint32(6814),
	207: uint32(6826),
	208: uint32(6838),
	209: uint32(6850),
	210: uint32(6862),
	211: uint32(6874),
	212: uint32(6886),
	213: uint32(6898),
	214: uint32(6910),
	215: uint32(6932),
	216: uint32(6944),
	217: uint32(6966),
	218: uint32(6978),
	219: uint32(6996),
	220: uint32(7008),
	221: uint32(7026),
	222: uint32(7044),
	223: uint32(7056),
	224: uint32(7068),
	225: uint32(7080),
	226: uint32(7102),
	227: uint32(7114),
	228: uint32(7126),
	229: uint32(7138),
	230: uint32(7150),
	231: uint32(7162),
	232: uint32(7174),
	233: uint32(7196),
	234: uint32(7208),
	235: uint32(7226),
	236: uint32(7248),
	237: uint32(7260),
	238: uint32(7279),
	239: uint32(7296),
	240: uint32(7309),
	241: uint32(7322),
	242: uint32(7337),
	243: uint32(7352),
	244: uint32(7365),
	245: uint32(7378),
	246: uint32(7391),
	247: uint32(7404),
	248: uint32(7417),
	249: uint32(7436),
	250: uint32(7455),
	251: uint32(7474),
	252: uint32(7493),
	253: uint32(7506),
	254: uint32(7519),
	255: uint32(7538),
	256: uint32(7551),
	257: uint32(7564),
	258: uint32(7581),
	259: uint32(7594),
	260: uint32(7613),
	261: uint32(7632),
	262: uint32(7645),
	263: uint32(7664),
	264: uint32(7683),
	265: uint32(7700),
	266: uint32(7713),
	267: uint32(7730),
	268: uint32(7743),
	269: uint32(7759),
	270: uint32(7773),
	271: uint32(7787),
	272: uint32(7801),
	273: uint32(7815),
	274: uint32(7829),
	275: uint32(7845),
	276: uint32(7859),
	277: uint32(7875),
	278: uint32(7891),
	279: uint32(7905),
	280: uint32(7919),
	281: uint32(7933),
	282: uint32(7949),
	283: uint32(7965),
	284: uint32(7979),
	285: uint32(7995),
	286: uint32(8011),
	287: uint32(8023),
	288: uint32(8037),
	289: uint32(8049),
	290: uint32(8065),
	291: uint32(8081),
	292: uint32(8097),
	293: uint32(8111),
	294: uint32(8127),
	295: uint32(8143),
	296: uint32(8157),
	297: uint32(8173),
	298: uint32(8189),
	299: uint32(8205),
	300: uint32(8219),
	301: uint32(8233),
	302: uint32(8249),
	303: uint32(8263),
	304: uint32(8277),
	305: uint32(8293),
	306: uint32(8307),
	307: uint32(8320),
	308: uint32(8333),
	309: uint32(8346),
	310: uint32(8357),
	311: uint32(8370),
	312: uint32(8383),
	313: uint32(8396),
	314: uint32(8409),
	315: uint32(8422),
	316: uint32(8435),
	317: uint32(8448),
	318: uint32(8461),
	319: uint32(8474),
	320: uint32(8487),
	321: uint32(8500),
	322: uint32(8513),
	323: uint32(8526),
	324: uint32(8539),
	325: uint32(8552),
	326: uint32(8563),
	327: uint32(8576),
	328: uint32(8585),
	329: uint32(8598),
	330: uint32(8611),
	331: uint32(8624),
	332: uint32(8637),
	333: uint32(8648),
	334: uint32(8661),
	335: uint32(8674),
	336: uint32(8687),
	337: uint32(8700),
	338: uint32(8713),
	339: uint32(8722),
	340: uint32(8735),
	341: uint32(8748),
	342: uint32(8761),
	343: uint32(8774),
	344: uint32(8785),
	345: uint32(8798),
	346: uint32(8811),
	347: uint32(8824),
	348: uint32(8837),
	349: uint32(8846),
	350: uint32(8859),
	351: uint32(8872),
	352: uint32(8885),
	353: uint32(8898),
	354: uint32(8911),
	355: uint32(8924),
	356: uint32(8937),
	357: uint32(8946),
	358: uint32(8957),
	359: uint32(8970),
	360: uint32(8983),
	361: uint32(8996),
	362: uint32(9005),
	363: uint32(9018),
	364: uint32(9031),
	365: uint32(9044),
	366: uint32(9053),
	367: uint32(9064),
	368: uint32(9077),
	369: uint32(9090),
	370: uint32(9103),
	371: uint32(9116),
	372: uint32(9129),
	373: uint32(9142),
	374: uint32(9155),
	375: uint32(9166),
	376: uint32(9179),
	377: uint32(9190),
	378: uint32(9203),
	379: uint32(9216),
	380: uint32(9229),
	381: uint32(9242),
	382: uint32(9251),
	383: uint32(9264),
	384: uint32(9277),
	385: uint32(9290),
	386: uint32(9303),
	387: uint32(9316),
	388: uint32(9325),
	389: uint32(9338),
	390: uint32(9351),
	391: uint32(9364),
	392: uint32(9377),
	393: uint32(9390),
	394: uint32(9403),
	395: uint32(9416),
	396: uint32(9429),
	397: uint32(9442),
	398: uint32(9451),
	399: uint32(9460),
	400: uint32(9469),
	401: uint32(9482),
	402: uint32(9495),
	403: uint32(9508),
	404: uint32(9521),
	405: uint32(9534),
	406: uint32(9545),
	407: uint32(9558),
	408: uint32(9571),
	409: uint32(9584),
	410: uint32(9597),
	411: uint32(9610),
	412: uint32(9623),
	413: uint32(9636),
	414: uint32(9649),
	415: uint32(9662),
	416: uint32(9675),
	417: uint32(9688),
	418: uint32(9701),
	419: uint32(9714),
	420: uint32(9727),
	421: uint32(9740),
	422: uint32(9753),
	423: uint32(9766),
	424: uint32(9779),
	425: uint32(9792),
	426: uint32(9805),
	427: uint32(9818),
	428: uint32(9831),
	429: uint32(9844),
	430: uint32(9853),
	431: uint32(9866),
	432: uint32(9879),
	433: uint32(9892),
	434: uint32(9905),
	435: uint32(9918),
	436: uint32(9931),
	437: uint32(9944),
	438: uint32(9957),
	439: uint32(9970),
	440: uint32(9979),
	441: uint32(9989),
	442: uint32(9999),
	443: uint32(10009),
	444: uint32(10019),
	445: uint32(10029),
	446: uint32(10037),
	447: uint32(10045),
	448: uint32(10053),
	449: uint32(10063),
	450: uint32(10071),
	451: uint32(10079),
	452: uint32(10089),
	453: uint32(10099),
	454: uint32(10109),
	455: uint32(10117),
	456: uint32(10127),
	457: uint32(10135),
	458: uint32(10145),
	459: uint32(10153),
	460: uint32(10163),
	461: uint32(10173),
	462: uint32(10181),
	463: uint32(10189),
	464: uint32(10197),
	465: uint32(10207),
	466: uint32(10217),
	467: uint32(10225),
	468: uint32(10233),
	469: uint32(10241),
	470: uint32(10249),
	471: uint32(10257),
	472: uint32(10267),
	473: uint32(10275),
	474: uint32(10282),
	475: uint32(10289),
	476: uint32(10296),
	477: uint32(10303),
	478: uint32(10310),
	479: uint32(10317),
	480: uint32(10324),
	481: uint32(10331),
	482: uint32(10338),
	483: uint32(10345),
	484: uint32(10352),
	485: uint32(10359),
	486: uint32(10366),
	487: uint32(10373),
	488: uint32(10380),
	489: uint32(10387),
	490: uint32(10394),
	491: uint32(10401),
	492: uint32(10408),
	493: uint32(10415),
	494: uint32(10422),
	495: uint32(10429),
	496: uint32(10436),
	497: uint32(10443),
	498: uint32(10450),
	499: uint32(10457),
	500: uint32(10464),
	501: uint32(10471),
	502: uint32(10478),
	503: uint32(10485),
	504: uint32(10492),
	505: uint32(10499),
	506: uint32(10506),
	507: uint32(10513),
	508: uint32(10520),
	509: uint32(10527),
	510: uint32(10534),
	511: uint32(10541),
	512: uint32(10548),
	513: uint32(10555),
	514: uint32(10562),
	515: uint32(10569),
	516: uint32(10576),
	517: uint32(10583),
	518: uint32(10590),
	519: uint32(10597),
	520: uint32(10604),
	521: uint32(10611),
	522: uint32(10618),
	523: uint32(10625),
	524: uint32(10632),
	525: uint32(10639),
	526: uint32(10646),
	527: uint32(10653),
	528: uint32(10660),
	529: uint32(10667),
	530: uint32(10674),
	531: uint32(10681),
	532: uint32(10688),
	533: uint32(10695),
	534: uint32(10702),
	535: uint32(10709),
	536: uint32(10716),
	537: uint32(10723),
	538: uint32(10730),
	539: uint32(10737),
	540: uint32(10744),
	541: uint32(10751),
	542: uint32(10758),
	543: uint32(10765),
	544: uint32(10772),
	545: uint32(10779),
	546: uint32(10786),
	547: uint32(10793),
	548: uint32(10800),
	549: uint32(10807),
	550: uint32(10814),
}

var ts_parse_actions = [1427]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_message),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(525)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(292)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(443)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(446)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(547)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(546)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(545)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(543)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(199)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(230)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(483)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(489)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(392)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(196)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(515)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(552)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(206)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(209)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(519)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(212)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(520)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(387)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym__annotation_array_def),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__annotation_array_def),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(243)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(487)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(491)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(490)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(233)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(202)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(262)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(446)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(547)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(546)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(545)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(543)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(476)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(262)),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(446)),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(547)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(546)),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(545)),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(543)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(476)),
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
		Fsymbol:      uint16(aux_sym_struct_repeat1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(154)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(323)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(324)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_message),
	})))),
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
		Fsymbol:      uint16(aux_sym_message_repeat1),
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
		Fsymbol:      uint16(aux_sym_message_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(525)),
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
		Fsymbol:      uint16(aux_sym_message_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(292)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(443)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(446)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(547)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(546)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_message_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(545)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_message_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(543)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_group_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(262)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_group_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(446)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_group_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(547)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_group_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(546)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_group_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(545)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_group_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(543)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_group_repeat1),
	})))),
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(476)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_group_repeat1),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(186)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(474)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(474)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_annotation),
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
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_annotation),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Fcount: uint8(1),
	}})),
	232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_annotation),
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
		Fcount: uint8(1),
	}})),
	240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_annotation),
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
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_const),
		Fproduction_id: uint16(17),
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
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_const),
		Fproduction_id: uint16(17),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_using_directive),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_using_directive),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	258: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Fcount: uint8(1),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Fcount: uint8(1),
	}})),
	268: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Fcount: uint8(1),
	}})),
	276: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_const),
		Fproduction_id: uint16(17),
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
		Fcount: uint8(1),
	}})),
	280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_const),
		Fproduction_id: uint16(17),
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
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_annotation),
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
		Fcount: uint8(1),
	}})),
	284: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_annotation),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
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
		Fcount: uint8(1),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	296: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(13),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
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
		Fcount: uint8(1),
	}})),
	300: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(13),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Fcount: uint8(1),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_annotation),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_annotation),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_const),
		Fproduction_id: uint16(17),
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
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_const),
		Fproduction_id: uint16(17),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(12),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Fcount: uint8(1),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(12),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(26),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_struct),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_using_directive),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_using_directive),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(12),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
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
		Fchild_count:   uint8(12),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(31),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(23),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_enum),
		Fproduction_id: uint16(6),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_interface),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_annotation),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_annotation),
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
		Fcount: uint8(1),
	}})),
	366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_block_text_repeat2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
	}})))),
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
		Fcount: uint8(1),
	}})),
	370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_text_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	376: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_text_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(108)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	379: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_text_repeat1),
	})))),
	380: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(108)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_block_text_repeat2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_generics),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(467)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(5),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(551)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fcount: uint8(2),
	}})),
	432: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_interface_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(467)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_interface_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(546)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	438: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_interface_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(545)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	441: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_interface_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(543)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_interface_repeat1),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(252)),
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
		Fcount: uint8(1),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(250)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	452: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym_concatenated_string_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	458: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_concatenated_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(222)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	461: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_concatenated_string_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(223)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(10),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	468: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(29),
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
		Fcount: uint8(1),
	}})),
	470: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(29),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_concatenated_string),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	476: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_const_value),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	478: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	480: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(25),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__unnamed_union_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(252)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	487: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__unnamed_union_repeat1),
	})))),
	488: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(250)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	490: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__unnamed_union_repeat1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	492: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(5),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	494: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	496: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(5),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	500: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_custom_type),
		Fproduction_id: uint16(4),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(499)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	508: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	510: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	512: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(30),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	514: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	516: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unique_id_statement),
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
		Fcount: uint8(1),
	}})),
	518: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_unique_id_statement),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_definition),
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
		Fcount: uint8(1),
	}})),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_definition),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	524: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	528: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_top_level_annotation),
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
		Fcount: uint8(1),
	}})),
	530: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_top_level_annotation),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	532: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(19),
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
		Fcount: uint8(1),
	}})),
	534: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_import),
		Fproduction_id: uint16(19),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(173)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field),
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
		Fcount: uint8(1),
	}})),
	542: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__unnamed_union),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	544: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__unnamed_union),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	546: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_custom_type_repeat1),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__annotation_call_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	550: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__annotation_call_repeat1),
	})))),
	551: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(551)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	552: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	553: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(28),
	})))),
	554: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	555: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(28),
	})))),
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
		Fcount: uint8(1),
	}})),
	557: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
	})))),
	558: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	559: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
	})))),
	560: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	561: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_group),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	563: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_group),
		Fproduction_id: uint16(4),
	})))),
	564: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	565: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	567: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
	}})),
	571: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	573: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__unnamed_union),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	577: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__unnamed_union),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	579: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_nested_struct),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_nested_struct),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_nested_enum),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	591: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_nested_enum),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_group),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_group),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
	}})),
	601: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_union),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	603: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_union),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(28),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	607: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(28),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
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
		Fcount: uint8(1),
	}})),
	613: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	615: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym__named_union),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
	}})),
	617: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_group),
		Fproduction_id: uint16(4),
	})))),
	618: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	619: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_group),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount: uint8(1),
	}})),
	621: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
	})))),
	622: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	623: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_field),
		Fproduction_id: uint16(24),
	})))),
	624: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	625: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_custom_type_repeat1),
	})))),
	626: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	627: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_custom_type_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(499)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_custom_type),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_block_text),
	})))),
	633: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	634: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__annotation_call_repeat1),
		Fproduction_id: uint16(12),
	})))),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	636: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_custom_type),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_text_repeat2),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	640: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_block_text_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_annotation_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	645: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_annotation_repeat1),
	})))),
	646: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(505)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	647: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	648: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__internal_const_identifier),
	})))),
	649: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(548)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	652: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym__internal_const_identifier_repeat1),
		Fproduction_id: uint16(4),
	})))),
	653: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	654: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(aux_sym_custom_type_repeat1),
		Fproduction_id: uint16(2),
	})))),
	655: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(216)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(505)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(127)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	662: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(501)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	666: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__internal_const_identifier),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	668: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(25),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	670: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_const_list),
	})))),
	671: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	672: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	673: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	674: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(236)),
	}})))),
	675: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(236)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	678: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_data),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	680: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(12),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(30),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	682: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(5),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	688: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_version),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	690: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(30),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	692: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_struct_shorthand),
	})))),
	693: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	694: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_const_list),
	})))),
	695: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	696: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_const_list),
	})))),
	697: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	698: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_struct_shorthand),
	})))),
	699: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	700: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
	}})))),
	701: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(480)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(5),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(257)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	708: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_list_type),
	})))),
	711: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	712: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	713: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	714: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(220)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	717: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	718: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(220)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	719: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	720: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(30),
	})))),
	721: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	722: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	723: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	724: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	725: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	726: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	727: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	728: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	729: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	730: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	731: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	732: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_embedded_file),
	})))),
	733: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	734: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(11),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(25),
	})))),
	735: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	736: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(366)),
	}})))),
	737: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	738: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	739: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	740: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(25),
	})))),
	741: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	742: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_primitive_type),
	})))),
	743: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	744: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	745: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	746: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_const_list),
	})))),
	747: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	748: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(255)),
	}})))),
	749: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	750: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	751: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	752: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(aux_sym_annotation_repeat1),
		Fproduction_id: uint16(5),
	})))),
	753: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	754: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	755: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	756: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	757: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(236)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	759: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	760: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(236)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	761: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	762: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	763: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	764: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_named_return_type),
		Fproduction_id: uint16(32),
	})))),
	765: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	766: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(518)),
	}})))),
	767: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	768: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_field_type),
	})))),
	769: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	770: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	771: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	772: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(362)),
	}})))),
	773: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	774: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	775: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	776: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(aux_sym_named_return_type_repeat1),
		Fproduction_id: uint16(34),
	})))),
	777: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	778: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_method),
	})))),
	779: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	780: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_method),
	})))),
	781: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	782: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	783: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	784: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	785: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	786: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym__internal_const_identifier_repeat1),
		Fproduction_id: uint16(4),
	})))),
	787: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	788: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	789: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	790: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	791: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	792: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	793: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	794: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	795: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	796: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	797: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	798: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	799: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	800: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	801: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	802: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	803: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	804: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_method),
	})))),
	805: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	806: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_method),
	})))),
	807: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	808: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	809: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	810: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	811: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	812: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	813: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	814: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	815: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	816: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	817: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	818: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	819: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	820: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(457)),
	}})))),
	821: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	822: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_method),
	})))),
	823: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	824: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_method),
	})))),
	825: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	826: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	827: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	828: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	829: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	830: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	831: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	832: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(22),
	})))),
	833: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	834: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	835: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	836: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	838: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	839: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	840: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(27),
	})))),
	841: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	842: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_method),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	844: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_method),
	})))),
	845: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	846: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	847: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	848: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	849: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	850: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(549)),
	}})))),
	851: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	852: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_method),
	})))),
	853: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	854: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_method),
	})))),
	855: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	856: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	857: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	858: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(27),
	})))),
	859: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	860: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	861: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	862: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_method),
		Fproduction_id: uint16(18),
	})))),
	863: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	864: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(248)),
	}})))),
	865: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	866: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	867: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	868: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__annotation_array_def_repeat2),
	})))),
	869: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	870: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__annotation_array_def_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(34)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	873: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_return_type),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	875: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_generic_parameters),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	879: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_repeat1),
	})))),
	880: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	881: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	882: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_enum_repeat1),
	})))),
	883: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	884: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__annotation_array_def_repeat1),
	})))),
	885: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	886: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__annotation_array_def_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(481)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(286)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(267)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(239)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	901: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(550)),
	}})))),
	902: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	903: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(431)),
	}})))),
	904: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	905: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(450)),
	}})))),
	906: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(316)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__annotation_array_def),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	913: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_generic_parameters_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	915: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_generic_parameters_repeat1),
	})))),
	916: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(65)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	917: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	918: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(359)),
	}})))),
	919: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	920: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	921: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	922: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	923: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	924: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_named_return_type),
		Fproduction_id: uint16(32),
	})))),
	925: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	926: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	927: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	928: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(423)),
	}})))),
	929: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	930: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	931: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	932: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym__annotation_array_def),
		Fproduction_id: uint16(15),
	})))),
	933: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	934: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	935: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	936: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(531)),
	}})))),
	937: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	938: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	939: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	940: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	941: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	942: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	943: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	944: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_generic_parameters),
		Fproduction_id: uint16(7),
	})))),
	945: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	946: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	947: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	948: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	949: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	950: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	951: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	952: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	953: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	954: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	955: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	956: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	957: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	958: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	959: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	960: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(aux_sym_struct_shorthand_repeat1),
		Fproduction_id: uint16(15),
	})))),
	961: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	962: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__annotation_array_def),
		Fproduction_id: uint16(15),
	})))),
	963: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	964: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_annotation_array_repeat1),
	})))),
	965: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	966: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_annotation_array_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(13)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(260)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	971: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	972: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	973: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(481)),
	}})))),
	974: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	977: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(407)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	981: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(214)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(aux_sym_named_return_type_repeat1),
		Fproduction_id: uint16(34),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_return_type),
		Fproduction_id: uint16(32),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_parameters),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(462)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_return_type),
		Fproduction_id: uint16(32),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(354)),
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
		Fcount: uint8(1),
	}})),
	1007: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__same_scope_const_value),
		Fproduction_id: uint16(20),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_named_return_type_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1017: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_named_return_type_repeat1),
	})))),
	1018: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(518)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1019: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1020: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_named_return_type),
		Fproduction_id: uint16(32),
	})))),
	1021: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1022: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
	1023: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1024: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
	1025: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1026: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_top_level_annotation_body_repeat1),
	})))),
	1027: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1028: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_top_level_annotation_body_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(504)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(340)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_annotation_array),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(461)),
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
	1043: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(401)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_import_using),
		Fproduction_id: uint16(1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(475)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(aux_sym__annotation_array_def_repeat1),
		Fproduction_id: uint16(33),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(131)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(522)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(504)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_method_parameters),
	})))),
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
		Fcount: uint8(1),
	}})),
	1069: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
	1072: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1073: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	1074: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1075: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(500)),
	}})))),
	1076: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1077: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(468)),
	}})))),
	1078: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1079: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_unnamed_return_type),
		Fproduction_id: uint16(4),
	})))),
	1080: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1081: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1082: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1083: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(242)),
	}})))),
	1084: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1085: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(451)),
	}})))),
	1086: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1087: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(241)),
	}})))),
	1088: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1089: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_import_using),
		Fproduction_id: uint16(1),
	})))),
	1090: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1091: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1092: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1093: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1094: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1095: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_nested_union),
	})))),
	1096: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1097: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_nested_union),
	})))),
	1098: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1099: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_union_field),
	})))),
	1100: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1101: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_union_field),
	})))),
	1102: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1103: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1104: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1106: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1107: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1108: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1109: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_generic_parameters_repeat1),
		Fproduction_id: uint16(14),
	})))),
	1110: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_replace_using),
		Fproduction_id: uint16(3),
	})))),
	1112: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1113: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1114: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1115: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(452)),
	}})))),
	1116: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(532)),
	}})))),
	1118: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1119: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_annotation_array),
	})))),
	1120: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1122: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1124: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1126: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1127: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1128: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1129: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1130: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(441)),
	}})))),
	1132: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1133: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1134: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1135: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1136: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_annotation_targets_repeat1),
	})))),
	1138: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1139: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_annotation_targets_repeat1),
	})))),
	1140: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1141: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_annotation_targets),
	})))),
	1143: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1144: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_annotation_array),
	})))),
	1145: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1146: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1147: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
	1149: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_union_field),
		Fproduction_id: uint16(24),
	})))),
	1151: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1152: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(498)),
	}})))),
	1153: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1154: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1155: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_struct_shorthand_repeat1),
	})))),
	1157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(532)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1158: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1159: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_struct_shorthand_repeat1),
	})))),
	1160: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1162: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1164: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1165: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(254)),
	}})))),
	1166: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1167: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(246)),
	}})))),
	1168: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
	})))),
	1170: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1171: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
	})))),
	1172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(462)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1173: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1175: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1176: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1177: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1179: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1181: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1183: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1184: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(247)),
	}})))),
	1185: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_import_using_repeat1),
	})))),
	1187: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1188: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_import_using_repeat1),
	})))),
	1189: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(475)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1190: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1191: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(380)),
	}})))),
	1192: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_method_parameters),
	})))),
	1194: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(118)),
	}})))),
	1196: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1197: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1198: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1199: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(470)),
	}})))),
	1200: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1201: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1202: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1203: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1204: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(183)),
	}})))),
	1206: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1207: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(399)),
	}})))),
	1208: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1209: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_annotation_targets),
	})))),
	1210: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1212: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(516)),
	}})))),
	1214: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(454)),
	}})))),
	1216: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1218: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	1219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__internal_const_identifier_repeat1),
	})))),
	1220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(473)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	1221: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1222: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1223: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(aux_sym_top_level_annotation_body_repeat1),
		Fproduction_id: uint16(5),
	})))),
	1225: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(aux_sym_named_return_type_repeat1),
		Fproduction_id: uint16(34),
	})))),
	1227: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(aux_sym_named_return_type_repeat1),
		Fproduction_id: uint16(34),
	})))),
	1229: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1230: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1231: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1232: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(459)),
	}})))),
	1233: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1234: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(aux_sym_struct_shorthand_repeat1),
		Fproduction_id: uint16(15),
	})))),
	1235: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1236: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_enum_field),
		Fproduction_id: uint16(13),
	})))),
	1237: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1238: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1239: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1240: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(512)),
	}})))),
	1241: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1242: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_top_level_annotation_body),
	})))),
	1243: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1244: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1245: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1246: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(507)),
	}})))),
	1247: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1248: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1249: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(27),
	})))),
	1251: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1252: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1253: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1254: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(373)),
	}})))),
	1255: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1256: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_import_using_repeat1),
		Fproduction_id: uint16(2),
	})))),
	1257: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1258: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(384)),
	}})))),
	1259: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(379)),
	}})))),
	1261: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1262: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_return_type),
	})))),
	1263: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_named_return_types),
	})))),
	1265: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_implicit_generics),
		Fproduction_id: uint16(21),
	})))),
	1267: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1268: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1269: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	1270: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__method_identifier),
		Fproduction_id: uint16(8),
	})))),
	1271: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1272: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym__method_identifier),
		Fproduction_id: uint16(8),
	})))),
	1273: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1274: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_named_return_types),
	})))),
	1275: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1276: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_enum_field),
		Fproduction_id: uint16(13),
	})))),
	1277: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1278: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_parameter),
		Fproduction_id: uint16(27),
	})))),
	1279: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1280: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_annotation_target),
	})))),
	1281: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(458)),
	}})))),
	1283: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1284: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1285: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1286: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
	}})))),
	1287: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(497)),
	}})))),
	1289: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1290: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(542)),
	}})))),
	1291: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1292: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(534)),
	}})))),
	1293: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1294: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1295: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1296: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1297: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1298: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1299: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1300: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1301: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1302: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1303: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1304: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1305: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(544)),
	}})))),
	1307: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_top_level_annotation_body),
	})))),
	1309: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(433)),
	}})))),
	1311: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1312: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1313: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1315: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1316: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1317: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1318: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1319: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1321: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_top_level_annotation_body),
		Fproduction_id: uint16(16),
	})))),
	1323: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1325: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(510)),
	}})))),
	1327: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(514)),
	}})))),
	1329: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(334)),
	}})))),
	1331: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1332: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(463)),
	}})))),
	1333: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1334: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(527)),
	}})))),
	1335: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1336: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1337: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(414)),
	}})))),
	1339: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1341: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1342: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1343: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1344: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1345: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1346: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1347: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1348: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(492)),
	}})))),
	1349: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1350: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(442)),
	}})))),
	1351: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(429)),
	}})))),
	1353: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1354: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1355: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1356: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(477)),
	}})))),
	1357: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1358: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_import_using),
		Fproduction_id: uint16(9),
	})))),
	1359: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1360: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(529)),
	}})))),
	1361: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1362: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1363: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1364: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1365: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1366: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1367: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1368: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_top_level_annotation_body),
		Fproduction_id: uint16(16),
	})))),
	1369: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1370: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1371: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1372: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_replace_using),
		Fproduction_id: uint16(3),
	})))),
	1373: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1375: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1376: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1377: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1379: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1380: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(536)),
	}})))),
	1381: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1382: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1383: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1384: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(493)),
	}})))),
	1385: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1386: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1387: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1388: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(484)),
	}})))),
	1389: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1390: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1391: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1392: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1393: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1394: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(240)),
	}})))),
	1395: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1396: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(464)),
	}})))),
	1397: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1398: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1399: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1400: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1401: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1402: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	1403: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(358)),
	}})))),
	1405: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1406: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(385)),
	}})))),
	1407: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1408: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1409: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1410: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_top_level_annotation_body),
	})))),
	1411: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1413: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1415: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1416: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(455)),
	}})))),
	1417: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym__internal_const_identifier_repeat1),
		Fproduction_id: uint16(4),
	})))),
	1419: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_top_level_annotation_body),
		Fproduction_id: uint16(11),
	})))),
	1421: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1423: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1425: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1426: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
}

func tree_sitter_capnp(tls *libc.TLS) (r uintptr) {
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
	Fkeyword_capture_token:     uint16(sym_identifier),
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
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

var __ccgo_ts1 = "end\x00annotation_definition_identifier\x00unique_id\x00;\x00using\x00=\x00import\x00.\x00$import\x00(\x00)\x00namespace\x00$\x00_type_identifier\x00,\x00annotation\x00:\x00*\x00const\x00enumerant\x00field\x00file\x00method\x00param\x00enum\x00group\x00interface\x00struct\x00union\x00[\x00]\x00{\x00}\x00extends\x00->\x00AnyPointer\x00Bool\x00Int8\x00Int16\x00Int32\x00Int64\x00UInt8\x00UInt16\x00UInt32\x00UInt64\x00Float32\x00Float64\x00Text\x00Data\x00Void\x00List\x00number\x00_normal_version\x00inline_field\x00float\x00true\x00false\x00data_hex\x00data_string\x00void\x00embed\x00\"\x00'\x00`\x00string_fragment\x00_escape_sequence_token1\x00escape_sequence\x00const_identifier\x00comment\x00message\x00statement\x00unique_id_statement\x00using_directive\x00replace_using\x00import_using\x00top_level_annotation\x00top_level_annotation_body\x00annotation_targets\x00annotation_target\x00annotation_literal\x00annotation_array\x00_annotation_array_def\x00definition\x00nested_struct\x00nested_enum\x00enum_field\x00nested_union\x00_unnamed_union\x00_named_union\x00union_field\x00method_parameters\x00parameters\x00parameter\x00return_type\x00named_return_types\x00unnamed_return_type\x00named_return_type\x00field_type\x00primitive_type\x00list_type\x00custom_type\x00const_value\x00_same_scope_const_value\x00field_version\x00_inline_version\x00boolean\x00data\x00const_list\x00struct_shorthand\x00_internal_const_identifier\x00embedded_file\x00generics\x00implicit_generics\x00generic_parameters\x00string\x00concatenated_string\x00block_text\x00_escape_sequence\x00_annotation_definition_identifier\x00_method_identifier\x00message_repeat1\x00import_using_repeat1\x00top_level_annotation_body_repeat1\x00annotation_repeat1\x00annotation_targets_repeat1\x00_annotation_call_repeat1\x00annotation_array_repeat1\x00_annotation_array_def_repeat1\x00_annotation_array_def_repeat2\x00struct_repeat1\x00enum_repeat1\x00group_repeat1\x00_unnamed_union_repeat1\x00interface_repeat1\x00parameters_repeat1\x00named_return_type_repeat1\x00custom_type_repeat1\x00struct_shorthand_repeat1\x00_internal_const_identifier_repeat1\x00generic_parameters_repeat1\x00string_repeat1\x00string_repeat2\x00concatenated_string_repeat1\x00block_text_repeat1\x00block_text_repeat2\x00annotation_identifier\x00attribute\x00enum_identifier\x00enum_member\x00extend_type\x00field_identifier\x00generic_identifier\x00implicit_generic_parameters\x00import_path\x00local_const\x00method_identifier\x00param_identifier\x00property\x00return_identifier\x00type_definition\x00type_identifier\x00"
