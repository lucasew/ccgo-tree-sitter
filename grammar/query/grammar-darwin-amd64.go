// Code generated for darwin/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -D__attribute__(...)= -D__extension__= -D_Nonnull= -D_Nullable= -D_Null_unspecified= -DAPI_AVAILABLE(...)= -DAPI_UNAVAILABLE(...)= -DAPI_DEPRECATED(...)= -DAPI_DEPRECATED_WITH_REPLACEMENT(...)= -D__API_AVAILABLE(...)= -D__API_UNAVAILABLE(...)= -D__API_DEPRECATED(...)= -D__API_DEPRECATED_WITH_REPLACEMENT(...)= -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-query/src -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-query -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-query/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build darwin && amd64

package grammar_query

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
const EXTERNAL_TOKEN_COUNT = 0
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
const LANGUAGE_VERSION = 15
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
const MAX_ALIAS_SEQUENCE_LENGTH = 10
const MAX_RESERVED_WORD_SET_SIZE = 0
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
const PRODUCTION_ID_COUNT = 22
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
const STATE_COUNT = 111
const SUPERTYPE_COUNT = 1
const SV_INTERRUPT = "SA_RESTART"
const SV_NOCLDSTOP = "SA_NOCLDSTOP"
const SV_NODEFER = "SA_NODEFER"
const SV_ONSTACK = "SA_ONSTACK"
const SV_RESETHAND = "SA_RESETHAND"
const SV_SIGINFO = "SA_SIGINFO"
const SYMBOL_COUNT = 51
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
const TOKEN_COUNT = 24
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

const anon_sym_DOT = 1
const sym_escape_sequence = 2
const anon_sym_STAR = 3
const anon_sym_PLUS = 4
const anon_sym_QMARK = 5
const sym_identifier = 6
const aux_sym__immediate_identifier_token1 = 7
const anon_sym__ = 8
const anon_sym_AT = 9
const anon_sym_DQUOTE = 10
const anon_sym_DQUOTE2 = 11
const aux_sym_string_content_token1 = 12
const sym_comment = 13
const anon_sym_LBRACK = 14
const anon_sym_RBRACK = 15
const anon_sym_LPAREN = 16
const anon_sym_RPAREN = 17
const anon_sym_MISSING = 18
const anon_sym_SLASH = 19
const anon_sym_COLON = 20
const anon_sym_BANG = 21
const anon_sym_POUND = 22
const sym_predicate_type = 23
const sym_program = 24
const sym_definition = 25
const sym__group_expression = 26
const sym__named_node_expression = 27
const sym_quantifier = 28
const sym__immediate_identifier = 29
const sym__node_identifier = 30
const sym_capture = 31
const sym_string = 32
const sym__immediate_string = 33
const sym_string_content = 34
const sym_parameters = 35
const sym_list = 36
const sym_grouping = 37
const sym_missing_node = 38
const sym_anonymous_node = 39
const sym_named_node = 40
const sym__field_name = 41
const sym_field_definition = 42
const sym_negated_field = 43
const sym_predicate = 44
const aux_sym_program_repeat1 = 45
const aux_sym_string_content_repeat1 = 46
const aux_sym_parameters_repeat1 = 47
const aux_sym_list_repeat1 = 48
const aux_sym_grouping_repeat1 = 49
const aux_sym_named_node_repeat1 = 50

var ts_symbol_names = [51]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 6,
	3:  __ccgo_ts + 22,
	4:  __ccgo_ts + 24,
	5:  __ccgo_ts + 26,
	6:  __ccgo_ts + 28,
	7:  __ccgo_ts + 28,
	8:  __ccgo_ts + 39,
	9:  __ccgo_ts + 41,
	10: __ccgo_ts + 43,
	11: __ccgo_ts + 43,
	12: __ccgo_ts + 45,
	13: __ccgo_ts + 67,
	14: __ccgo_ts + 75,
	15: __ccgo_ts + 77,
	16: __ccgo_ts + 79,
	17: __ccgo_ts + 81,
	18: __ccgo_ts + 83,
	19: __ccgo_ts + 91,
	20: __ccgo_ts + 93,
	21: __ccgo_ts + 95,
	22: __ccgo_ts + 97,
	23: __ccgo_ts + 99,
	24: __ccgo_ts + 114,
	25: __ccgo_ts + 122,
	26: __ccgo_ts + 133,
	27: __ccgo_ts + 151,
	28: __ccgo_ts + 174,
	29: __ccgo_ts + 185,
	30: __ccgo_ts + 207,
	31: __ccgo_ts + 224,
	32: __ccgo_ts + 232,
	33: __ccgo_ts + 232,
	34: __ccgo_ts + 239,
	35: __ccgo_ts + 254,
	36: __ccgo_ts + 265,
	37: __ccgo_ts + 270,
	38: __ccgo_ts + 279,
	39: __ccgo_ts + 292,
	40: __ccgo_ts + 307,
	41: __ccgo_ts + 318,
	42: __ccgo_ts + 330,
	43: __ccgo_ts + 347,
	44: __ccgo_ts + 361,
	45: __ccgo_ts + 371,
	46: __ccgo_ts + 387,
	47: __ccgo_ts + 410,
	48: __ccgo_ts + 429,
	49: __ccgo_ts + 442,
	50: __ccgo_ts + 459,
}

var ts_symbol_map = [51]TSSymbol{
	1:  uint16(anon_sym_DOT),
	2:  uint16(sym_escape_sequence),
	3:  uint16(anon_sym_STAR),
	4:  uint16(anon_sym_PLUS),
	5:  uint16(anon_sym_QMARK),
	6:  uint16(sym_identifier),
	7:  uint16(sym_identifier),
	8:  uint16(anon_sym__),
	9:  uint16(anon_sym_AT),
	10: uint16(anon_sym_DQUOTE),
	11: uint16(anon_sym_DQUOTE),
	12: uint16(aux_sym_string_content_token1),
	13: uint16(sym_comment),
	14: uint16(anon_sym_LBRACK),
	15: uint16(anon_sym_RBRACK),
	16: uint16(anon_sym_LPAREN),
	17: uint16(anon_sym_RPAREN),
	18: uint16(anon_sym_MISSING),
	19: uint16(anon_sym_SLASH),
	20: uint16(anon_sym_COLON),
	21: uint16(anon_sym_BANG),
	22: uint16(anon_sym_POUND),
	23: uint16(sym_predicate_type),
	24: uint16(sym_program),
	25: uint16(sym_definition),
	26: uint16(sym__group_expression),
	27: uint16(sym__named_node_expression),
	28: uint16(sym_quantifier),
	29: uint16(sym__immediate_identifier),
	30: uint16(sym__node_identifier),
	31: uint16(sym_capture),
	32: uint16(sym_string),
	33: uint16(sym_string),
	34: uint16(sym_string_content),
	35: uint16(sym_parameters),
	36: uint16(sym_list),
	37: uint16(sym_grouping),
	38: uint16(sym_missing_node),
	39: uint16(sym_anonymous_node),
	40: uint16(sym_named_node),
	41: uint16(sym__field_name),
	42: uint16(sym_field_definition),
	43: uint16(sym_negated_field),
	44: uint16(sym_predicate),
	45: uint16(aux_sym_program_repeat1),
	46: uint16(aux_sym_string_content_repeat1),
	47: uint16(aux_sym_parameters_repeat1),
	48: uint16(aux_sym_list_repeat1),
	49: uint16(aux_sym_grouping_repeat1),
	50: uint16(aux_sym_named_node_repeat1),
}

var ts_symbol_metadata = [51]TSSymbolMetadata{
	0: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	1: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	12: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	24: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	25: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	26: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	27: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	29: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	31: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	32: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	33: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	34: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	38: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	41: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	45: {},
	46: {},
	47: {},
	48: {},
	49: {},
	50: {},
}

type ts_field_identifiers = int32

const field_name = 1
const field_parameters = 2
const field_quantifier = 3
const field_supertype = 4
const field_type = 5

var ts_field_names = [6]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 478,
	2: __ccgo_ts + 254,
	3: __ccgo_ts + 174,
	4: __ccgo_ts + 483,
	5: __ccgo_ts + 493,
}

var ts_field_map_slices = [22]TSMapSlice{
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
		Flength: uint16(2),
	},
	6: {
		Findex:  uint16(7),
		Flength: uint16(1),
	},
	7: {
		Findex:  uint16(8),
		Flength: uint16(1),
	},
	8: {
		Findex:  uint16(9),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(11),
		Flength: uint16(3),
	},
	10: {
		Findex:  uint16(14),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(16),
		Flength: uint16(2),
	},
	12: {
		Findex:  uint16(18),
		Flength: uint16(2),
	},
	13: {
		Findex:  uint16(20),
		Flength: uint16(4),
	},
	14: {
		Findex:  uint16(24),
		Flength: uint16(3),
	},
	15: {
		Findex:  uint16(27),
		Flength: uint16(2),
	},
	16: {
		Findex:  uint16(29),
		Flength: uint16(3),
	},
	17: {
		Findex:  uint16(32),
		Flength: uint16(2),
	},
	18: {
		Findex:  uint16(34),
		Flength: uint16(3),
	},
	19: {
		Findex:  uint16(37),
		Flength: uint16(2),
	},
	20: {
		Findex:  uint16(39),
		Flength: uint16(3),
	},
	21: {
		Findex:  uint16(42),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [45]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_name),
	},
	1: {
		Ffield_id: uint16(field_quantifier),
	},
	2: {
		Ffield_id: uint16(field_name),
	},
	3: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	5: {
		Ffield_id:  uint16(field_quantifier),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	6: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	7: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	8: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	9: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	10: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	12: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	13: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(3),
	},
	14: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	15: {
		Ffield_id:    uint16(field_supertype),
		Fchild_index: uint8(1),
	},
	16: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	17: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	18: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	19: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(4),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	21: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(2),
	},
	22: {
		Ffield_id:    uint16(field_parameters),
		Fchild_index: uint8(4),
	},
	23: {
		Ffield_id:    uint16(field_type),
		Fchild_index: uint8(3),
	},
	24: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	25: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	26: {
		Ffield_id:    uint16(field_supertype),
		Fchild_index: uint8(1),
	},
	27: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	28: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(5),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	29: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	30: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	31: {
		Ffield_id:    uint16(field_supertype),
		Fchild_index: uint8(1),
	},
	32: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	33: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(6),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	34: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	35: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(7),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	36: {
		Ffield_id:    uint16(field_supertype),
		Fchild_index: uint8(1),
	},
	37: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(1),
	},
	38: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(7),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	39: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	40: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(8),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	41: {
		Ffield_id:    uint16(field_supertype),
		Fchild_index: uint8(1),
	},
	42: {
		Ffield_id:    uint16(field_name),
		Fchild_index: uint8(3),
	},
	43: {
		Ffield_id:    uint16(field_quantifier),
		Fchild_index: uint8(9),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	44: {
		Ffield_id:    uint16(field_supertype),
		Fchild_index: uint8(1),
	},
}

var ts_alias_sequences = [22][10]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [111]TSStateId{
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
	59:  uint16(57),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(62),
	63:  uint16(63),
	64:  uint16(64),
	65:  uint16(65),
	66:  uint16(38),
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
	90:  uint16(88),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(93),
	94:  uint16(94),
	95:  uint16(60),
	96:  uint16(96),
	97:  uint16(61),
	98:  uint16(62),
	99:  uint16(99),
	100: uint16(100),
	101: uint16(100),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(106),
	107: uint16(107),
	108: uint16(108),
	109: uint16(109),
	110: uint16(104),
}

var ts_supertype_symbols = [1]TSSymbol{
	0: uint16(sym_definition),
}

var ts_supertype_map_slices = [26]TSMapSlice{
	25: {
		Flength: uint16(7),
	},
}

var ts_supertype_map_entries = [7]TSSymbol{
	0: uint16(sym_anonymous_node),
	1: uint16(sym_field_definition),
	2: uint16(sym_grouping),
	3: uint16(sym_list),
	4: uint16(sym_missing_node),
	5: uint16(sym_named_node),
	6: uint16(sym_predicate),
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
			state = uint16(11)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(72)/libc.Uint64FromInt64(2)) {
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
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('-') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('"') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(7)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(29)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('"') {
			state = uint16(28)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(31)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(3):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
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
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(4):
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
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
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('"') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32(';') {
			state = uint16(31)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(26)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(5)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32(';') {
			state = uint16(31)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(8):
		if eof != 0 {
			state = uint16(11)
			goto next_state
		}
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(64)/libc.Uint64FromInt64(2)) {
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
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(9):
		if eof != 0 {
			state = uint16(11)
			goto next_state
		}
		i4 = uint32(0)
		for {
			if !(uint64(i4) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if int32('-') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(10):
		if eof != 0 {
			state = uint16(11)
			goto next_state
		}
		i5 = uint32(0)
		for {
			if !(uint64(i5) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
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
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(10)
			goto next_state
		}
		if lookahead == int32('-') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(11):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(12):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(13):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_escape_sequence)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(14):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(15):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(16):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(17):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('G') {
			state = uint16(36)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(18):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('I') {
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(19):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('I') {
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(20):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('N') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(21):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('S') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(22):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('S') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(23):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_identifier)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(24):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__immediate_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(25):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym__)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(28):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(29):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(';') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(29)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) && lookahead != int32('"') && lookahead != int32('\\') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(30):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_string_content_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('"') && lookahead != int32('\\') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_comment)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32(0x17f) && lookahead != int32(0x212a) {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(32):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(33):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
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
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_MISSING)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') || lookahead == int32('.') || int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(37):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(38):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(39):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(40):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(41):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_predicate_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [36]uint16_t{
	0:  uint16('!'),
	1:  uint16(39),
	2:  uint16('"'),
	3:  uint16(28),
	4:  uint16('#'),
	5:  uint16(40),
	6:  uint16('('),
	7:  uint16(34),
	8:  uint16(')'),
	9:  uint16(35),
	10: uint16('*'),
	11: uint16(14),
	12: uint16('+'),
	13: uint16(15),
	14: uint16('.'),
	15: uint16(12),
	16: uint16('/'),
	17: uint16(37),
	18: uint16(':'),
	19: uint16(38),
	20: uint16(';'),
	21: uint16(31),
	22: uint16('?'),
	23: uint16(16),
	24: uint16('@'),
	25: uint16(26),
	26: uint16('M'),
	27: uint16(18),
	28: uint16('['),
	29: uint16(32),
	30: uint16('\\'),
	31: uint16(7),
	32: uint16(']'),
	33: uint16(33),
	34: uint16('_'),
	35: uint16(25),
}

var map_token1 = [20]uint16_t{
	0:  uint16('"'),
	1:  uint16(27),
	2:  uint16('#'),
	3:  uint16(40),
	4:  uint16('('),
	5:  uint16(34),
	6:  uint16('.'),
	7:  uint16(12),
	8:  uint16(';'),
	9:  uint16(31),
	10: uint16('M'),
	11: uint16(18),
	12: uint16('['),
	13: uint16(32),
	14: uint16('_'),
	15: uint16(25),
	16: uint16('!'),
	17: uint16(41),
	18: uint16('?'),
	19: uint16(41),
}

var map_token2 = [16]uint16_t{
	0:  uint16('"'),
	1:  uint16(27),
	2:  uint16('#'),
	3:  uint16(40),
	4:  uint16('('),
	5:  uint16(34),
	6:  uint16('.'),
	7:  uint16(12),
	8:  uint16(';'),
	9:  uint16(31),
	10: uint16('M'),
	11: uint16(18),
	12: uint16('['),
	13: uint16(32),
	14: uint16('_'),
	15: uint16(25),
}

var map_token3 = [32]uint16_t{
	0:  uint16('!'),
	1:  uint16(39),
	2:  uint16('"'),
	3:  uint16(27),
	4:  uint16('#'),
	5:  uint16(40),
	6:  uint16('('),
	7:  uint16(34),
	8:  uint16(')'),
	9:  uint16(35),
	10: uint16('*'),
	11: uint16(14),
	12: uint16('+'),
	13: uint16(15),
	14: uint16('.'),
	15: uint16(12),
	16: uint16(':'),
	17: uint16(38),
	18: uint16(';'),
	19: uint16(31),
	20: uint16('?'),
	21: uint16(16),
	22: uint16('@'),
	23: uint16(26),
	24: uint16('M'),
	25: uint16(18),
	26: uint16('['),
	27: uint16(32),
	28: uint16(']'),
	29: uint16(33),
	30: uint16('_'),
	31: uint16(25),
}

var map_token4 = [30]uint16_t{
	0:  uint16('!'),
	1:  uint16(39),
	2:  uint16('"'),
	3:  uint16(27),
	4:  uint16('('),
	5:  uint16(34),
	6:  uint16(')'),
	7:  uint16(35),
	8:  uint16('*'),
	9:  uint16(14),
	10: uint16('+'),
	11: uint16(15),
	12: uint16('.'),
	13: uint16(12),
	14: uint16('/'),
	15: uint16(37),
	16: uint16(':'),
	17: uint16(38),
	18: uint16(';'),
	19: uint16(31),
	20: uint16('?'),
	21: uint16(16),
	22: uint16('@'),
	23: uint16(26),
	24: uint16('['),
	25: uint16(32),
	26: uint16(']'),
	27: uint16(33),
	28: uint16('_'),
	29: uint16(25),
}

var map_token5 = [28]uint16_t{
	0:  uint16('!'),
	1:  uint16(39),
	2:  uint16('"'),
	3:  uint16(27),
	4:  uint16('('),
	5:  uint16(34),
	6:  uint16(')'),
	7:  uint16(35),
	8:  uint16('*'),
	9:  uint16(14),
	10: uint16('+'),
	11: uint16(15),
	12: uint16('.'),
	13: uint16(12),
	14: uint16(':'),
	15: uint16(38),
	16: uint16(';'),
	17: uint16(31),
	18: uint16('?'),
	19: uint16(16),
	20: uint16('@'),
	21: uint16(26),
	22: uint16('['),
	23: uint16(32),
	24: uint16(']'),
	25: uint16(33),
	26: uint16('_'),
	27: uint16(25),
}

var ts_lex_modes = [111]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(9),
	},
	2: {
		Flex_state: uint16(3),
	},
	3: {
		Flex_state: uint16(9),
	},
	4: {
		Flex_state: uint16(9),
	},
	5: {
		Flex_state: uint16(9),
	},
	6: {
		Flex_state: uint16(9),
	},
	7: {
		Flex_state: uint16(9),
	},
	8: {
		Flex_state: uint16(9),
	},
	9: {
		Flex_state: uint16(9),
	},
	10: {
		Flex_state: uint16(9),
	},
	11: {
		Flex_state: uint16(9),
	},
	12: {
		Flex_state: uint16(9),
	},
	13: {
		Flex_state: uint16(9),
	},
	14: {
		Flex_state: uint16(9),
	},
	15: {
		Flex_state: uint16(9),
	},
	16: {
		Flex_state: uint16(9),
	},
	17: {
		Flex_state: uint16(9),
	},
	18: {
		Flex_state: uint16(9),
	},
	19: {
		Flex_state: uint16(9),
	},
	20: {
		Flex_state: uint16(9),
	},
	21: {
		Flex_state: uint16(9),
	},
	22: {
		Flex_state: uint16(9),
	},
	23: {
		Flex_state: uint16(9),
	},
	24: {
		Flex_state: uint16(9),
	},
	25: {
		Flex_state: uint16(9),
	},
	26: {
		Flex_state: uint16(9),
	},
	27: {
		Flex_state: uint16(9),
	},
	28: {
		Flex_state: uint16(9),
	},
	29: {
		Flex_state: uint16(9),
	},
	30: {
		Flex_state: uint16(9),
	},
	31: {
		Flex_state: uint16(9),
	},
	32: {
		Flex_state: uint16(9),
	},
	33: {
		Flex_state: uint16(9),
	},
	34: {
		Flex_state: uint16(9),
	},
	35: {
		Flex_state: uint16(9),
	},
	36: {
		Flex_state: uint16(9),
	},
	37: {
		Flex_state: uint16(9),
	},
	38: {
		Flex_state: uint16(9),
	},
	39: {
		Flex_state: uint16(9),
	},
	40: {
		Flex_state: uint16(9),
	},
	41: {
		Flex_state: uint16(9),
	},
	42: {
		Flex_state: uint16(9),
	},
	43: {
		Flex_state: uint16(9),
	},
	44: {
		Flex_state: uint16(9),
	},
	45: {
		Flex_state: uint16(9),
	},
	46: {
		Flex_state: uint16(9),
	},
	47: {
		Flex_state: uint16(9),
	},
	48: {
		Flex_state: uint16(9),
	},
	49: {
		Flex_state: uint16(9),
	},
	50: {
		Flex_state: uint16(9),
	},
	51: {
		Flex_state: uint16(9),
	},
	52: {
		Flex_state: uint16(9),
	},
	53: {
		Flex_state: uint16(9),
	},
	54: {
		Flex_state: uint16(9),
	},
	55: {
		Flex_state: uint16(9),
	},
	56: {
		Flex_state: uint16(9),
	},
	57: {
		Flex_state: uint16(9),
	},
	58: {
		Flex_state: uint16(9),
	},
	59: {
		Flex_state: uint16(9),
	},
	60: {
		Flex_state: uint16(9),
	},
	61: {
		Flex_state: uint16(9),
	},
	62: {
		Flex_state: uint16(9),
	},
	63: {
		Flex_state: uint16(9),
	},
	64: {
		Flex_state: uint16(9),
	},
	65: {
		Flex_state: uint16(9),
	},
	66: {
		Flex_state: uint16(9),
	},
	67: {
		Flex_state: uint16(9),
	},
	68: {
		Flex_state: uint16(9),
	},
	69: {
		Flex_state: uint16(9),
	},
	70: {
		Flex_state: uint16(9),
	},
	71: {
		Flex_state: uint16(9),
	},
	72: {
		Flex_state: uint16(9),
	},
	73: {
		Flex_state: uint16(9),
	},
	74: {
		Flex_state: uint16(5),
	},
	75: {
		Flex_state: uint16(9),
	},
	76: {
		Flex_state: uint16(9),
	},
	77: {
		Flex_state: uint16(9),
	},
	78: {
		Flex_state: uint16(9),
	},
	79: {
		Flex_state: uint16(9),
	},
	80: {
		Flex_state: uint16(9),
	},
	81: {
		Flex_state: uint16(9),
	},
	82: {
		Flex_state: uint16(9),
	},
	83: {
		Flex_state: uint16(9),
	},
	84: {
		Flex_state: uint16(5),
	},
	85: {
		Flex_state: uint16(9),
	},
	86: {
		Flex_state: uint16(9),
	},
	87: {
		Flex_state: uint16(5),
	},
	88: {
		Flex_state: uint16(1),
	},
	89: {
		Flex_state: uint16(9),
	},
	90: {
		Flex_state: uint16(1),
	},
	91: {
		Flex_state: uint16(1),
	},
	92: {
		Flex_state: uint16(1),
	},
	93: {
		Flex_state: uint16(1),
	},
	94: {
		Flex_state: uint16(5),
	},
	95: {
		Flex_state: uint16(5),
	},
	96: {
		Flex_state: uint16(2),
	},
	97: {
		Flex_state: uint16(5),
	},
	98: {
		Flex_state: uint16(5),
	},
	99: {
		Flex_state: uint16(2),
	},
	100: {
		Flex_state: uint16(2),
	},
	101: {
		Flex_state: uint16(2),
	},
	102: {
		Flex_state: uint16(5),
	},
	103: {},
	104: {
		Flex_state: uint16(2),
	},
	105: {},
	106: {},
	107: {
		Flex_state: uint16(3),
	},
	108: {
		Flex_state: uint16(2),
	},
	109: {},
	110: {
		Flex_state: uint16(2),
	},
}

var ts_parse_table = [2][51]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
		4:  uint16(1),
		5:  uint16(1),
		6:  uint16(1),
		8:  uint16(1),
		9:  uint16(1),
		10: uint16(1),
		11: uint16(1),
		13: uint16(3),
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
	},
	1: {
		0:  uint16(5),
		6:  uint16(7),
		8:  uint16(9),
		10: uint16(11),
		13: uint16(3),
		14: uint16(13),
		16: uint16(15),
		24: uint16(105),
		25: uint16(33),
		32: uint16(38),
		36: uint16(68),
		37: uint16(68),
		38: uint16(68),
		39: uint16(68),
		40: uint16(68),
		41: uint16(59),
		42: uint16(68),
		44: uint16(68),
		45: uint16(33),
	},
}

var ts_small_parse_table = [3011]uint16_t{
	0:    uint16(14),
	1:    uint16(3),
	2:    uint16(1),
	3:    uint16(sym_comment),
	4:    uint16(11),
	5:    uint16(1),
	6:    uint16(anon_sym_DQUOTE),
	7:    uint16(13),
	8:    uint16(1),
	9:    uint16(anon_sym_LBRACK),
	10:   uint16(15),
	11:   uint16(1),
	12:   uint16(anon_sym_LPAREN),
	13:   uint16(19),
	14:   uint16(1),
	15:   uint16(sym_identifier),
	16:   uint16(21),
	17:   uint16(1),
	18:   uint16(anon_sym__),
	19:   uint16(23),
	20:   uint16(1),
	21:   uint16(anon_sym_MISSING),
	22:   uint16(3),
	23:   uint16(1),
	24:   uint16(sym__node_identifier),
	25:   uint16(18),
	26:   uint16(1),
	27:   uint16(aux_sym_grouping_repeat1),
	28:   uint16(57),
	29:   uint16(1),
	30:   uint16(sym__field_name),
	31:   uint16(66),
	32:   uint16(1),
	33:   uint16(sym_string),
	34:   uint16(17),
	35:   uint16(2),
	36:   uint16(anon_sym_DOT),
	37:   uint16(anon_sym_POUND),
	38:   uint16(86),
	39:   uint16(2),
	40:   uint16(sym_definition),
	41:   uint16(sym__group_expression),
	42:   uint16(68),
	43:   uint16(7),
	44:   uint16(sym_list),
	45:   uint16(sym_grouping),
	46:   uint16(sym_missing_node),
	47:   uint16(sym_anonymous_node),
	48:   uint16(sym_named_node),
	49:   uint16(sym_field_definition),
	50:   uint16(sym_predicate),
	51:   uint16(14),
	52:   uint16(3),
	53:   uint16(1),
	54:   uint16(sym_comment),
	55:   uint16(7),
	56:   uint16(1),
	57:   uint16(sym_identifier),
	58:   uint16(9),
	59:   uint16(1),
	60:   uint16(anon_sym__),
	61:   uint16(11),
	62:   uint16(1),
	63:   uint16(anon_sym_DQUOTE),
	64:   uint16(13),
	65:   uint16(1),
	66:   uint16(anon_sym_LBRACK),
	67:   uint16(15),
	68:   uint16(1),
	69:   uint16(anon_sym_LPAREN),
	70:   uint16(25),
	71:   uint16(1),
	72:   uint16(anon_sym_DOT),
	73:   uint16(27),
	74:   uint16(1),
	75:   uint16(anon_sym_RPAREN),
	76:   uint16(29),
	77:   uint16(1),
	78:   uint16(anon_sym_BANG),
	79:   uint16(5),
	80:   uint16(1),
	81:   uint16(aux_sym_named_node_repeat1),
	82:   uint16(38),
	83:   uint16(1),
	84:   uint16(sym_string),
	85:   uint16(59),
	86:   uint16(1),
	87:   uint16(sym__field_name),
	88:   uint16(80),
	89:   uint16(3),
	90:   uint16(sym_definition),
	91:   uint16(sym__named_node_expression),
	92:   uint16(sym_negated_field),
	93:   uint16(68),
	94:   uint16(7),
	95:   uint16(sym_list),
	96:   uint16(sym_grouping),
	97:   uint16(sym_missing_node),
	98:   uint16(sym_anonymous_node),
	99:   uint16(sym_named_node),
	100:  uint16(sym_field_definition),
	101:  uint16(sym_predicate),
	102:  uint16(14),
	103:  uint16(3),
	104:  uint16(1),
	105:  uint16(sym_comment),
	106:  uint16(7),
	107:  uint16(1),
	108:  uint16(sym_identifier),
	109:  uint16(9),
	110:  uint16(1),
	111:  uint16(anon_sym__),
	112:  uint16(11),
	113:  uint16(1),
	114:  uint16(anon_sym_DQUOTE),
	115:  uint16(13),
	116:  uint16(1),
	117:  uint16(anon_sym_LBRACK),
	118:  uint16(15),
	119:  uint16(1),
	120:  uint16(anon_sym_LPAREN),
	121:  uint16(29),
	122:  uint16(1),
	123:  uint16(anon_sym_BANG),
	124:  uint16(31),
	125:  uint16(1),
	126:  uint16(anon_sym_DOT),
	127:  uint16(33),
	128:  uint16(1),
	129:  uint16(anon_sym_RPAREN),
	130:  uint16(7),
	131:  uint16(1),
	132:  uint16(aux_sym_named_node_repeat1),
	133:  uint16(38),
	134:  uint16(1),
	135:  uint16(sym_string),
	136:  uint16(59),
	137:  uint16(1),
	138:  uint16(sym__field_name),
	139:  uint16(72),
	140:  uint16(3),
	141:  uint16(sym_definition),
	142:  uint16(sym__named_node_expression),
	143:  uint16(sym_negated_field),
	144:  uint16(68),
	145:  uint16(7),
	146:  uint16(sym_list),
	147:  uint16(sym_grouping),
	148:  uint16(sym_missing_node),
	149:  uint16(sym_anonymous_node),
	150:  uint16(sym_named_node),
	151:  uint16(sym_field_definition),
	152:  uint16(sym_predicate),
	153:  uint16(13),
	154:  uint16(3),
	155:  uint16(1),
	156:  uint16(sym_comment),
	157:  uint16(7),
	158:  uint16(1),
	159:  uint16(sym_identifier),
	160:  uint16(9),
	161:  uint16(1),
	162:  uint16(anon_sym__),
	163:  uint16(11),
	164:  uint16(1),
	165:  uint16(anon_sym_DQUOTE),
	166:  uint16(13),
	167:  uint16(1),
	168:  uint16(anon_sym_LBRACK),
	169:  uint16(15),
	170:  uint16(1),
	171:  uint16(anon_sym_LPAREN),
	172:  uint16(29),
	173:  uint16(1),
	174:  uint16(anon_sym_BANG),
	175:  uint16(35),
	176:  uint16(1),
	177:  uint16(anon_sym_RPAREN),
	178:  uint16(8),
	179:  uint16(1),
	180:  uint16(aux_sym_named_node_repeat1),
	181:  uint16(38),
	182:  uint16(1),
	183:  uint16(sym_string),
	184:  uint16(59),
	185:  uint16(1),
	186:  uint16(sym__field_name),
	187:  uint16(82),
	188:  uint16(3),
	189:  uint16(sym_definition),
	190:  uint16(sym__named_node_expression),
	191:  uint16(sym_negated_field),
	192:  uint16(68),
	193:  uint16(7),
	194:  uint16(sym_list),
	195:  uint16(sym_grouping),
	196:  uint16(sym_missing_node),
	197:  uint16(sym_anonymous_node),
	198:  uint16(sym_named_node),
	199:  uint16(sym_field_definition),
	200:  uint16(sym_predicate),
	201:  uint16(13),
	202:  uint16(3),
	203:  uint16(1),
	204:  uint16(sym_comment),
	205:  uint16(7),
	206:  uint16(1),
	207:  uint16(sym_identifier),
	208:  uint16(9),
	209:  uint16(1),
	210:  uint16(anon_sym__),
	211:  uint16(11),
	212:  uint16(1),
	213:  uint16(anon_sym_DQUOTE),
	214:  uint16(13),
	215:  uint16(1),
	216:  uint16(anon_sym_LBRACK),
	217:  uint16(15),
	218:  uint16(1),
	219:  uint16(anon_sym_LPAREN),
	220:  uint16(29),
	221:  uint16(1),
	222:  uint16(anon_sym_BANG),
	223:  uint16(37),
	224:  uint16(1),
	225:  uint16(anon_sym_RPAREN),
	226:  uint16(8),
	227:  uint16(1),
	228:  uint16(aux_sym_named_node_repeat1),
	229:  uint16(38),
	230:  uint16(1),
	231:  uint16(sym_string),
	232:  uint16(59),
	233:  uint16(1),
	234:  uint16(sym__field_name),
	235:  uint16(73),
	236:  uint16(3),
	237:  uint16(sym_definition),
	238:  uint16(sym__named_node_expression),
	239:  uint16(sym_negated_field),
	240:  uint16(68),
	241:  uint16(7),
	242:  uint16(sym_list),
	243:  uint16(sym_grouping),
	244:  uint16(sym_missing_node),
	245:  uint16(sym_anonymous_node),
	246:  uint16(sym_named_node),
	247:  uint16(sym_field_definition),
	248:  uint16(sym_predicate),
	249:  uint16(13),
	250:  uint16(3),
	251:  uint16(1),
	252:  uint16(sym_comment),
	253:  uint16(7),
	254:  uint16(1),
	255:  uint16(sym_identifier),
	256:  uint16(9),
	257:  uint16(1),
	258:  uint16(anon_sym__),
	259:  uint16(11),
	260:  uint16(1),
	261:  uint16(anon_sym_DQUOTE),
	262:  uint16(13),
	263:  uint16(1),
	264:  uint16(anon_sym_LBRACK),
	265:  uint16(15),
	266:  uint16(1),
	267:  uint16(anon_sym_LPAREN),
	268:  uint16(29),
	269:  uint16(1),
	270:  uint16(anon_sym_BANG),
	271:  uint16(39),
	272:  uint16(1),
	273:  uint16(anon_sym_RPAREN),
	274:  uint16(8),
	275:  uint16(1),
	276:  uint16(aux_sym_named_node_repeat1),
	277:  uint16(38),
	278:  uint16(1),
	279:  uint16(sym_string),
	280:  uint16(59),
	281:  uint16(1),
	282:  uint16(sym__field_name),
	283:  uint16(78),
	284:  uint16(3),
	285:  uint16(sym_definition),
	286:  uint16(sym__named_node_expression),
	287:  uint16(sym_negated_field),
	288:  uint16(68),
	289:  uint16(7),
	290:  uint16(sym_list),
	291:  uint16(sym_grouping),
	292:  uint16(sym_missing_node),
	293:  uint16(sym_anonymous_node),
	294:  uint16(sym_named_node),
	295:  uint16(sym_field_definition),
	296:  uint16(sym_predicate),
	297:  uint16(13),
	298:  uint16(3),
	299:  uint16(1),
	300:  uint16(sym_comment),
	301:  uint16(41),
	302:  uint16(1),
	303:  uint16(sym_identifier),
	304:  uint16(44),
	305:  uint16(1),
	306:  uint16(anon_sym__),
	307:  uint16(47),
	308:  uint16(1),
	309:  uint16(anon_sym_DQUOTE),
	310:  uint16(50),
	311:  uint16(1),
	312:  uint16(anon_sym_LBRACK),
	313:  uint16(53),
	314:  uint16(1),
	315:  uint16(anon_sym_LPAREN),
	316:  uint16(56),
	317:  uint16(1),
	318:  uint16(anon_sym_RPAREN),
	319:  uint16(58),
	320:  uint16(1),
	321:  uint16(anon_sym_BANG),
	322:  uint16(8),
	323:  uint16(1),
	324:  uint16(aux_sym_named_node_repeat1),
	325:  uint16(38),
	326:  uint16(1),
	327:  uint16(sym_string),
	328:  uint16(59),
	329:  uint16(1),
	330:  uint16(sym__field_name),
	331:  uint16(76),
	332:  uint16(3),
	333:  uint16(sym_definition),
	334:  uint16(sym__named_node_expression),
	335:  uint16(sym_negated_field),
	336:  uint16(68),
	337:  uint16(7),
	338:  uint16(sym_list),
	339:  uint16(sym_grouping),
	340:  uint16(sym_missing_node),
	341:  uint16(sym_anonymous_node),
	342:  uint16(sym_named_node),
	343:  uint16(sym_field_definition),
	344:  uint16(sym_predicate),
	345:  uint16(13),
	346:  uint16(3),
	347:  uint16(1),
	348:  uint16(sym_comment),
	349:  uint16(7),
	350:  uint16(1),
	351:  uint16(sym_identifier),
	352:  uint16(9),
	353:  uint16(1),
	354:  uint16(anon_sym__),
	355:  uint16(11),
	356:  uint16(1),
	357:  uint16(anon_sym_DQUOTE),
	358:  uint16(13),
	359:  uint16(1),
	360:  uint16(anon_sym_LBRACK),
	361:  uint16(15),
	362:  uint16(1),
	363:  uint16(anon_sym_LPAREN),
	364:  uint16(29),
	365:  uint16(1),
	366:  uint16(anon_sym_BANG),
	367:  uint16(61),
	368:  uint16(1),
	369:  uint16(anon_sym_RPAREN),
	370:  uint16(8),
	371:  uint16(1),
	372:  uint16(aux_sym_named_node_repeat1),
	373:  uint16(38),
	374:  uint16(1),
	375:  uint16(sym_string),
	376:  uint16(59),
	377:  uint16(1),
	378:  uint16(sym__field_name),
	379:  uint16(81),
	380:  uint16(3),
	381:  uint16(sym_definition),
	382:  uint16(sym__named_node_expression),
	383:  uint16(sym_negated_field),
	384:  uint16(68),
	385:  uint16(7),
	386:  uint16(sym_list),
	387:  uint16(sym_grouping),
	388:  uint16(sym_missing_node),
	389:  uint16(sym_anonymous_node),
	390:  uint16(sym_named_node),
	391:  uint16(sym_field_definition),
	392:  uint16(sym_predicate),
	393:  uint16(12),
	394:  uint16(3),
	395:  uint16(1),
	396:  uint16(sym_comment),
	397:  uint16(7),
	398:  uint16(1),
	399:  uint16(sym_identifier),
	400:  uint16(9),
	401:  uint16(1),
	402:  uint16(anon_sym__),
	403:  uint16(11),
	404:  uint16(1),
	405:  uint16(anon_sym_DQUOTE),
	406:  uint16(13),
	407:  uint16(1),
	408:  uint16(anon_sym_LBRACK),
	409:  uint16(15),
	410:  uint16(1),
	411:  uint16(anon_sym_LPAREN),
	412:  uint16(29),
	413:  uint16(1),
	414:  uint16(anon_sym_BANG),
	415:  uint16(6),
	416:  uint16(1),
	417:  uint16(aux_sym_named_node_repeat1),
	418:  uint16(38),
	419:  uint16(1),
	420:  uint16(sym_string),
	421:  uint16(59),
	422:  uint16(1),
	423:  uint16(sym__field_name),
	424:  uint16(82),
	425:  uint16(3),
	426:  uint16(sym_definition),
	427:  uint16(sym__named_node_expression),
	428:  uint16(sym_negated_field),
	429:  uint16(68),
	430:  uint16(7),
	431:  uint16(sym_list),
	432:  uint16(sym_grouping),
	433:  uint16(sym_missing_node),
	434:  uint16(sym_anonymous_node),
	435:  uint16(sym_named_node),
	436:  uint16(sym_field_definition),
	437:  uint16(sym_predicate),
	438:  uint16(12),
	439:  uint16(3),
	440:  uint16(1),
	441:  uint16(sym_comment),
	442:  uint16(7),
	443:  uint16(1),
	444:  uint16(sym_identifier),
	445:  uint16(9),
	446:  uint16(1),
	447:  uint16(anon_sym__),
	448:  uint16(11),
	449:  uint16(1),
	450:  uint16(anon_sym_DQUOTE),
	451:  uint16(13),
	452:  uint16(1),
	453:  uint16(anon_sym_LBRACK),
	454:  uint16(15),
	455:  uint16(1),
	456:  uint16(anon_sym_LPAREN),
	457:  uint16(29),
	458:  uint16(1),
	459:  uint16(anon_sym_BANG),
	460:  uint16(37),
	461:  uint16(1),
	462:  uint16(anon_sym_RPAREN),
	463:  uint16(38),
	464:  uint16(1),
	465:  uint16(sym_string),
	466:  uint16(59),
	467:  uint16(1),
	468:  uint16(sym__field_name),
	469:  uint16(75),
	470:  uint16(3),
	471:  uint16(sym_definition),
	472:  uint16(sym__named_node_expression),
	473:  uint16(sym_negated_field),
	474:  uint16(68),
	475:  uint16(7),
	476:  uint16(sym_list),
	477:  uint16(sym_grouping),
	478:  uint16(sym_missing_node),
	479:  uint16(sym_anonymous_node),
	480:  uint16(sym_named_node),
	481:  uint16(sym_field_definition),
	482:  uint16(sym_predicate),
	483:  uint16(12),
	484:  uint16(3),
	485:  uint16(1),
	486:  uint16(sym_comment),
	487:  uint16(7),
	488:  uint16(1),
	489:  uint16(sym_identifier),
	490:  uint16(9),
	491:  uint16(1),
	492:  uint16(anon_sym__),
	493:  uint16(11),
	494:  uint16(1),
	495:  uint16(anon_sym_DQUOTE),
	496:  uint16(13),
	497:  uint16(1),
	498:  uint16(anon_sym_LBRACK),
	499:  uint16(15),
	500:  uint16(1),
	501:  uint16(anon_sym_LPAREN),
	502:  uint16(29),
	503:  uint16(1),
	504:  uint16(anon_sym_BANG),
	505:  uint16(63),
	506:  uint16(1),
	507:  uint16(anon_sym_RPAREN),
	508:  uint16(38),
	509:  uint16(1),
	510:  uint16(sym_string),
	511:  uint16(59),
	512:  uint16(1),
	513:  uint16(sym__field_name),
	514:  uint16(75),
	515:  uint16(3),
	516:  uint16(sym_definition),
	517:  uint16(sym__named_node_expression),
	518:  uint16(sym_negated_field),
	519:  uint16(68),
	520:  uint16(7),
	521:  uint16(sym_list),
	522:  uint16(sym_grouping),
	523:  uint16(sym_missing_node),
	524:  uint16(sym_anonymous_node),
	525:  uint16(sym_named_node),
	526:  uint16(sym_field_definition),
	527:  uint16(sym_predicate),
	528:  uint16(12),
	529:  uint16(3),
	530:  uint16(1),
	531:  uint16(sym_comment),
	532:  uint16(7),
	533:  uint16(1),
	534:  uint16(sym_identifier),
	535:  uint16(9),
	536:  uint16(1),
	537:  uint16(anon_sym__),
	538:  uint16(11),
	539:  uint16(1),
	540:  uint16(anon_sym_DQUOTE),
	541:  uint16(13),
	542:  uint16(1),
	543:  uint16(anon_sym_LBRACK),
	544:  uint16(15),
	545:  uint16(1),
	546:  uint16(anon_sym_LPAREN),
	547:  uint16(29),
	548:  uint16(1),
	549:  uint16(anon_sym_BANG),
	550:  uint16(9),
	551:  uint16(1),
	552:  uint16(aux_sym_named_node_repeat1),
	553:  uint16(38),
	554:  uint16(1),
	555:  uint16(sym_string),
	556:  uint16(59),
	557:  uint16(1),
	558:  uint16(sym__field_name),
	559:  uint16(78),
	560:  uint16(3),
	561:  uint16(sym_definition),
	562:  uint16(sym__named_node_expression),
	563:  uint16(sym_negated_field),
	564:  uint16(68),
	565:  uint16(7),
	566:  uint16(sym_list),
	567:  uint16(sym_grouping),
	568:  uint16(sym_missing_node),
	569:  uint16(sym_anonymous_node),
	570:  uint16(sym_named_node),
	571:  uint16(sym_field_definition),
	572:  uint16(sym_predicate),
	573:  uint16(12),
	574:  uint16(3),
	575:  uint16(1),
	576:  uint16(sym_comment),
	577:  uint16(7),
	578:  uint16(1),
	579:  uint16(sym_identifier),
	580:  uint16(9),
	581:  uint16(1),
	582:  uint16(anon_sym__),
	583:  uint16(11),
	584:  uint16(1),
	585:  uint16(anon_sym_DQUOTE),
	586:  uint16(13),
	587:  uint16(1),
	588:  uint16(anon_sym_LBRACK),
	589:  uint16(15),
	590:  uint16(1),
	591:  uint16(anon_sym_LPAREN),
	592:  uint16(29),
	593:  uint16(1),
	594:  uint16(anon_sym_BANG),
	595:  uint16(61),
	596:  uint16(1),
	597:  uint16(anon_sym_RPAREN),
	598:  uint16(38),
	599:  uint16(1),
	600:  uint16(sym_string),
	601:  uint16(59),
	602:  uint16(1),
	603:  uint16(sym__field_name),
	604:  uint16(75),
	605:  uint16(3),
	606:  uint16(sym_definition),
	607:  uint16(sym__named_node_expression),
	608:  uint16(sym_negated_field),
	609:  uint16(68),
	610:  uint16(7),
	611:  uint16(sym_list),
	612:  uint16(sym_grouping),
	613:  uint16(sym_missing_node),
	614:  uint16(sym_anonymous_node),
	615:  uint16(sym_named_node),
	616:  uint16(sym_field_definition),
	617:  uint16(sym_predicate),
	618:  uint16(12),
	619:  uint16(3),
	620:  uint16(1),
	621:  uint16(sym_comment),
	622:  uint16(7),
	623:  uint16(1),
	624:  uint16(sym_identifier),
	625:  uint16(9),
	626:  uint16(1),
	627:  uint16(anon_sym__),
	628:  uint16(11),
	629:  uint16(1),
	630:  uint16(anon_sym_DQUOTE),
	631:  uint16(13),
	632:  uint16(1),
	633:  uint16(anon_sym_LBRACK),
	634:  uint16(15),
	635:  uint16(1),
	636:  uint16(anon_sym_LPAREN),
	637:  uint16(29),
	638:  uint16(1),
	639:  uint16(anon_sym_BANG),
	640:  uint16(65),
	641:  uint16(1),
	642:  uint16(anon_sym_RPAREN),
	643:  uint16(38),
	644:  uint16(1),
	645:  uint16(sym_string),
	646:  uint16(59),
	647:  uint16(1),
	648:  uint16(sym__field_name),
	649:  uint16(75),
	650:  uint16(3),
	651:  uint16(sym_definition),
	652:  uint16(sym__named_node_expression),
	653:  uint16(sym_negated_field),
	654:  uint16(68),
	655:  uint16(7),
	656:  uint16(sym_list),
	657:  uint16(sym_grouping),
	658:  uint16(sym_missing_node),
	659:  uint16(sym_anonymous_node),
	660:  uint16(sym_named_node),
	661:  uint16(sym_field_definition),
	662:  uint16(sym_predicate),
	663:  uint16(12),
	664:  uint16(3),
	665:  uint16(1),
	666:  uint16(sym_comment),
	667:  uint16(7),
	668:  uint16(1),
	669:  uint16(sym_identifier),
	670:  uint16(9),
	671:  uint16(1),
	672:  uint16(anon_sym__),
	673:  uint16(11),
	674:  uint16(1),
	675:  uint16(anon_sym_DQUOTE),
	676:  uint16(13),
	677:  uint16(1),
	678:  uint16(anon_sym_LBRACK),
	679:  uint16(15),
	680:  uint16(1),
	681:  uint16(anon_sym_LPAREN),
	682:  uint16(29),
	683:  uint16(1),
	684:  uint16(anon_sym_BANG),
	685:  uint16(67),
	686:  uint16(1),
	687:  uint16(anon_sym_RPAREN),
	688:  uint16(38),
	689:  uint16(1),
	690:  uint16(sym_string),
	691:  uint16(59),
	692:  uint16(1),
	693:  uint16(sym__field_name),
	694:  uint16(75),
	695:  uint16(3),
	696:  uint16(sym_definition),
	697:  uint16(sym__named_node_expression),
	698:  uint16(sym_negated_field),
	699:  uint16(68),
	700:  uint16(7),
	701:  uint16(sym_list),
	702:  uint16(sym_grouping),
	703:  uint16(sym_missing_node),
	704:  uint16(sym_anonymous_node),
	705:  uint16(sym_named_node),
	706:  uint16(sym_field_definition),
	707:  uint16(sym_predicate),
	708:  uint16(12),
	709:  uint16(3),
	710:  uint16(1),
	711:  uint16(sym_comment),
	712:  uint16(7),
	713:  uint16(1),
	714:  uint16(sym_identifier),
	715:  uint16(9),
	716:  uint16(1),
	717:  uint16(anon_sym__),
	718:  uint16(11),
	719:  uint16(1),
	720:  uint16(anon_sym_DQUOTE),
	721:  uint16(13),
	722:  uint16(1),
	723:  uint16(anon_sym_LBRACK),
	724:  uint16(15),
	725:  uint16(1),
	726:  uint16(anon_sym_LPAREN),
	727:  uint16(29),
	728:  uint16(1),
	729:  uint16(anon_sym_BANG),
	730:  uint16(69),
	731:  uint16(1),
	732:  uint16(anon_sym_RPAREN),
	733:  uint16(38),
	734:  uint16(1),
	735:  uint16(sym_string),
	736:  uint16(59),
	737:  uint16(1),
	738:  uint16(sym__field_name),
	739:  uint16(75),
	740:  uint16(3),
	741:  uint16(sym_definition),
	742:  uint16(sym__named_node_expression),
	743:  uint16(sym_negated_field),
	744:  uint16(68),
	745:  uint16(7),
	746:  uint16(sym_list),
	747:  uint16(sym_grouping),
	748:  uint16(sym_missing_node),
	749:  uint16(sym_anonymous_node),
	750:  uint16(sym_named_node),
	751:  uint16(sym_field_definition),
	752:  uint16(sym_predicate),
	753:  uint16(12),
	754:  uint16(3),
	755:  uint16(1),
	756:  uint16(sym_comment),
	757:  uint16(7),
	758:  uint16(1),
	759:  uint16(sym_identifier),
	760:  uint16(11),
	761:  uint16(1),
	762:  uint16(anon_sym_DQUOTE),
	763:  uint16(13),
	764:  uint16(1),
	765:  uint16(anon_sym_LBRACK),
	766:  uint16(15),
	767:  uint16(1),
	768:  uint16(anon_sym_LPAREN),
	769:  uint16(71),
	770:  uint16(1),
	771:  uint16(anon_sym__),
	772:  uint16(73),
	773:  uint16(1),
	774:  uint16(anon_sym_RPAREN),
	775:  uint16(19),
	776:  uint16(1),
	777:  uint16(aux_sym_grouping_repeat1),
	778:  uint16(57),
	779:  uint16(1),
	780:  uint16(sym__field_name),
	781:  uint16(66),
	782:  uint16(1),
	783:  uint16(sym_string),
	784:  uint16(86),
	785:  uint16(2),
	786:  uint16(sym_definition),
	787:  uint16(sym__group_expression),
	788:  uint16(68),
	789:  uint16(7),
	790:  uint16(sym_list),
	791:  uint16(sym_grouping),
	792:  uint16(sym_missing_node),
	793:  uint16(sym_anonymous_node),
	794:  uint16(sym_named_node),
	795:  uint16(sym_field_definition),
	796:  uint16(sym_predicate),
	797:  uint16(12),
	798:  uint16(3),
	799:  uint16(1),
	800:  uint16(sym_comment),
	801:  uint16(75),
	802:  uint16(1),
	803:  uint16(sym_identifier),
	804:  uint16(78),
	805:  uint16(1),
	806:  uint16(anon_sym__),
	807:  uint16(81),
	808:  uint16(1),
	809:  uint16(anon_sym_DQUOTE),
	810:  uint16(84),
	811:  uint16(1),
	812:  uint16(anon_sym_LBRACK),
	813:  uint16(87),
	814:  uint16(1),
	815:  uint16(anon_sym_LPAREN),
	816:  uint16(90),
	817:  uint16(1),
	818:  uint16(anon_sym_RPAREN),
	819:  uint16(19),
	820:  uint16(1),
	821:  uint16(aux_sym_grouping_repeat1),
	822:  uint16(57),
	823:  uint16(1),
	824:  uint16(sym__field_name),
	825:  uint16(66),
	826:  uint16(1),
	827:  uint16(sym_string),
	828:  uint16(86),
	829:  uint16(2),
	830:  uint16(sym_definition),
	831:  uint16(sym__group_expression),
	832:  uint16(68),
	833:  uint16(7),
	834:  uint16(sym_list),
	835:  uint16(sym_grouping),
	836:  uint16(sym_missing_node),
	837:  uint16(sym_anonymous_node),
	838:  uint16(sym_named_node),
	839:  uint16(sym_field_definition),
	840:  uint16(sym_predicate),
	841:  uint16(11),
	842:  uint16(3),
	843:  uint16(1),
	844:  uint16(sym_comment),
	845:  uint16(94),
	846:  uint16(1),
	847:  uint16(sym_identifier),
	848:  uint16(97),
	849:  uint16(1),
	850:  uint16(anon_sym__),
	851:  uint16(100),
	852:  uint16(1),
	853:  uint16(anon_sym_DQUOTE),
	854:  uint16(103),
	855:  uint16(1),
	856:  uint16(anon_sym_LBRACK),
	857:  uint16(106),
	858:  uint16(1),
	859:  uint16(anon_sym_LPAREN),
	860:  uint16(38),
	861:  uint16(1),
	862:  uint16(sym_string),
	863:  uint16(59),
	864:  uint16(1),
	865:  uint16(sym__field_name),
	866:  uint16(92),
	867:  uint16(2),
	869:  uint16(anon_sym_RBRACK),
	870:  uint16(20),
	871:  uint16(2),
	872:  uint16(sym_definition),
	873:  uint16(aux_sym_program_repeat1),
	874:  uint16(68),
	875:  uint16(7),
	876:  uint16(sym_list),
	877:  uint16(sym_grouping),
	878:  uint16(sym_missing_node),
	879:  uint16(sym_anonymous_node),
	880:  uint16(sym_named_node),
	881:  uint16(sym_field_definition),
	882:  uint16(sym_predicate),
	883:  uint16(11),
	884:  uint16(3),
	885:  uint16(1),
	886:  uint16(sym_comment),
	887:  uint16(7),
	888:  uint16(1),
	889:  uint16(sym_identifier),
	890:  uint16(9),
	891:  uint16(1),
	892:  uint16(anon_sym__),
	893:  uint16(11),
	894:  uint16(1),
	895:  uint16(anon_sym_DQUOTE),
	896:  uint16(13),
	897:  uint16(1),
	898:  uint16(anon_sym_LBRACK),
	899:  uint16(15),
	900:  uint16(1),
	901:  uint16(anon_sym_LPAREN),
	902:  uint16(29),
	903:  uint16(1),
	904:  uint16(anon_sym_BANG),
	905:  uint16(38),
	906:  uint16(1),
	907:  uint16(sym_string),
	908:  uint16(59),
	909:  uint16(1),
	910:  uint16(sym__field_name),
	911:  uint16(75),
	912:  uint16(3),
	913:  uint16(sym_definition),
	914:  uint16(sym__named_node_expression),
	915:  uint16(sym_negated_field),
	916:  uint16(68),
	917:  uint16(7),
	918:  uint16(sym_list),
	919:  uint16(sym_grouping),
	920:  uint16(sym_missing_node),
	921:  uint16(sym_anonymous_node),
	922:  uint16(sym_named_node),
	923:  uint16(sym_field_definition),
	924:  uint16(sym_predicate),
	925:  uint16(7),
	926:  uint16(3),
	927:  uint16(1),
	928:  uint16(sym_comment),
	929:  uint16(115),
	930:  uint16(1),
	931:  uint16(anon_sym_AT),
	932:  uint16(65),
	933:  uint16(1),
	934:  uint16(sym_quantifier),
	935:  uint16(113),
	936:  uint16(2),
	937:  uint16(sym_identifier),
	938:  uint16(anon_sym__),
	939:  uint16(24),
	940:  uint16(2),
	941:  uint16(sym_capture),
	942:  uint16(aux_sym_list_repeat1),
	943:  uint16(111),
	944:  uint16(3),
	945:  uint16(anon_sym_STAR),
	946:  uint16(anon_sym_PLUS),
	947:  uint16(anon_sym_QMARK),
	948:  uint16(109),
	949:  uint16(8),
	951:  uint16(anon_sym_DOT),
	952:  uint16(anon_sym_DQUOTE),
	953:  uint16(anon_sym_LBRACK),
	954:  uint16(anon_sym_RBRACK),
	955:  uint16(anon_sym_LPAREN),
	956:  uint16(anon_sym_RPAREN),
	957:  uint16(anon_sym_BANG),
	958:  uint16(7),
	959:  uint16(3),
	960:  uint16(1),
	961:  uint16(sym_comment),
	962:  uint16(115),
	963:  uint16(1),
	964:  uint16(anon_sym_AT),
	965:  uint16(65),
	966:  uint16(1),
	967:  uint16(sym_quantifier),
	968:  uint16(119),
	969:  uint16(2),
	970:  uint16(sym_identifier),
	971:  uint16(anon_sym__),
	972:  uint16(24),
	973:  uint16(2),
	974:  uint16(sym_capture),
	975:  uint16(aux_sym_list_repeat1),
	976:  uint16(111),
	977:  uint16(3),
	978:  uint16(anon_sym_STAR),
	979:  uint16(anon_sym_PLUS),
	980:  uint16(anon_sym_QMARK),
	981:  uint16(117),
	982:  uint16(8),
	984:  uint16(anon_sym_DOT),
	985:  uint16(anon_sym_DQUOTE),
	986:  uint16(anon_sym_LBRACK),
	987:  uint16(anon_sym_RBRACK),
	988:  uint16(anon_sym_LPAREN),
	989:  uint16(anon_sym_RPAREN),
	990:  uint16(anon_sym_BANG),
	991:  uint16(7),
	992:  uint16(3),
	993:  uint16(1),
	994:  uint16(sym_comment),
	995:  uint16(128),
	996:  uint16(1),
	997:  uint16(anon_sym_AT),
	998:  uint16(65),
	999:  uint16(1),
	1000: uint16(sym_quantifier),
	1001: uint16(126),
	1002: uint16(2),
	1003: uint16(sym_identifier),
	1004: uint16(anon_sym__),
	1005: uint16(24),
	1006: uint16(2),
	1007: uint16(sym_capture),
	1008: uint16(aux_sym_list_repeat1),
	1009: uint16(123),
	1010: uint16(3),
	1011: uint16(anon_sym_STAR),
	1012: uint16(anon_sym_PLUS),
	1013: uint16(anon_sym_QMARK),
	1014: uint16(121),
	1015: uint16(8),
	1017: uint16(anon_sym_DOT),
	1018: uint16(anon_sym_DQUOTE),
	1019: uint16(anon_sym_LBRACK),
	1020: uint16(anon_sym_RBRACK),
	1021: uint16(anon_sym_LPAREN),
	1022: uint16(anon_sym_RPAREN),
	1023: uint16(anon_sym_BANG),
	1024: uint16(7),
	1025: uint16(3),
	1026: uint16(1),
	1027: uint16(sym_comment),
	1028: uint16(115),
	1029: uint16(1),
	1030: uint16(anon_sym_AT),
	1031: uint16(65),
	1032: uint16(1),
	1033: uint16(sym_quantifier),
	1034: uint16(133),
	1035: uint16(2),
	1036: uint16(sym_identifier),
	1037: uint16(anon_sym__),
	1038: uint16(30),
	1039: uint16(2),
	1040: uint16(sym_capture),
	1041: uint16(aux_sym_list_repeat1),
	1042: uint16(111),
	1043: uint16(3),
	1044: uint16(anon_sym_STAR),
	1045: uint16(anon_sym_PLUS),
	1046: uint16(anon_sym_QMARK),
	1047: uint16(131),
	1048: uint16(8),
	1050: uint16(anon_sym_DOT),
	1051: uint16(anon_sym_DQUOTE),
	1052: uint16(anon_sym_LBRACK),
	1053: uint16(anon_sym_RBRACK),
	1054: uint16(anon_sym_LPAREN),
	1055: uint16(anon_sym_RPAREN),
	1056: uint16(anon_sym_BANG),
	1057: uint16(7),
	1058: uint16(3),
	1059: uint16(1),
	1060: uint16(sym_comment),
	1061: uint16(115),
	1062: uint16(1),
	1063: uint16(anon_sym_AT),
	1064: uint16(65),
	1065: uint16(1),
	1066: uint16(sym_quantifier),
	1067: uint16(137),
	1068: uint16(2),
	1069: uint16(sym_identifier),
	1070: uint16(anon_sym__),
	1071: uint16(32),
	1072: uint16(2),
	1073: uint16(sym_capture),
	1074: uint16(aux_sym_list_repeat1),
	1075: uint16(111),
	1076: uint16(3),
	1077: uint16(anon_sym_STAR),
	1078: uint16(anon_sym_PLUS),
	1079: uint16(anon_sym_QMARK),
	1080: uint16(135),
	1081: uint16(8),
	1083: uint16(anon_sym_DOT),
	1084: uint16(anon_sym_DQUOTE),
	1085: uint16(anon_sym_LBRACK),
	1086: uint16(anon_sym_RBRACK),
	1087: uint16(anon_sym_LPAREN),
	1088: uint16(anon_sym_RPAREN),
	1089: uint16(anon_sym_BANG),
	1090: uint16(11),
	1091: uint16(3),
	1092: uint16(1),
	1093: uint16(sym_comment),
	1094: uint16(7),
	1095: uint16(1),
	1096: uint16(sym_identifier),
	1097: uint16(11),
	1098: uint16(1),
	1099: uint16(anon_sym_DQUOTE),
	1100: uint16(13),
	1101: uint16(1),
	1102: uint16(anon_sym_LBRACK),
	1103: uint16(15),
	1104: uint16(1),
	1105: uint16(anon_sym_LPAREN),
	1106: uint16(71),
	1107: uint16(1),
	1108: uint16(anon_sym__),
	1109: uint16(90),
	1110: uint16(1),
	1111: uint16(anon_sym_RPAREN),
	1112: uint16(57),
	1113: uint16(1),
	1114: uint16(sym__field_name),
	1115: uint16(66),
	1116: uint16(1),
	1117: uint16(sym_string),
	1118: uint16(85),
	1119: uint16(2),
	1120: uint16(sym_definition),
	1121: uint16(sym__group_expression),
	1122: uint16(68),
	1123: uint16(7),
	1124: uint16(sym_list),
	1125: uint16(sym_grouping),
	1126: uint16(sym_missing_node),
	1127: uint16(sym_anonymous_node),
	1128: uint16(sym_named_node),
	1129: uint16(sym_field_definition),
	1130: uint16(sym_predicate),
	1131: uint16(7),
	1132: uint16(3),
	1133: uint16(1),
	1134: uint16(sym_comment),
	1135: uint16(115),
	1136: uint16(1),
	1137: uint16(anon_sym_AT),
	1138: uint16(65),
	1139: uint16(1),
	1140: uint16(sym_quantifier),
	1141: uint16(141),
	1142: uint16(2),
	1143: uint16(sym_identifier),
	1144: uint16(anon_sym__),
	1145: uint16(22),
	1146: uint16(2),
	1147: uint16(sym_capture),
	1148: uint16(aux_sym_list_repeat1),
	1149: uint16(111),
	1150: uint16(3),
	1151: uint16(anon_sym_STAR),
	1152: uint16(anon_sym_PLUS),
	1153: uint16(anon_sym_QMARK),
	1154: uint16(139),
	1155: uint16(8),
	1157: uint16(anon_sym_DOT),
	1158: uint16(anon_sym_DQUOTE),
	1159: uint16(anon_sym_LBRACK),
	1160: uint16(anon_sym_RBRACK),
	1161: uint16(anon_sym_LPAREN),
	1162: uint16(anon_sym_RPAREN),
	1163: uint16(anon_sym_BANG),
	1164: uint16(7),
	1165: uint16(3),
	1166: uint16(1),
	1167: uint16(sym_comment),
	1168: uint16(115),
	1169: uint16(1),
	1170: uint16(anon_sym_AT),
	1171: uint16(65),
	1172: uint16(1),
	1173: uint16(sym_quantifier),
	1174: uint16(145),
	1175: uint16(2),
	1176: uint16(sym_identifier),
	1177: uint16(anon_sym__),
	1178: uint16(35),
	1179: uint16(2),
	1180: uint16(sym_capture),
	1181: uint16(aux_sym_list_repeat1),
	1182: uint16(111),
	1183: uint16(3),
	1184: uint16(anon_sym_STAR),
	1185: uint16(anon_sym_PLUS),
	1186: uint16(anon_sym_QMARK),
	1187: uint16(143),
	1188: uint16(8),
	1190: uint16(anon_sym_DOT),
	1191: uint16(anon_sym_DQUOTE),
	1192: uint16(anon_sym_LBRACK),
	1193: uint16(anon_sym_RBRACK),
	1194: uint16(anon_sym_LPAREN),
	1195: uint16(anon_sym_RPAREN),
	1196: uint16(anon_sym_BANG),
	1197: uint16(7),
	1198: uint16(3),
	1199: uint16(1),
	1200: uint16(sym_comment),
	1201: uint16(115),
	1202: uint16(1),
	1203: uint16(anon_sym_AT),
	1204: uint16(65),
	1205: uint16(1),
	1206: uint16(sym_quantifier),
	1207: uint16(149),
	1208: uint16(2),
	1209: uint16(sym_identifier),
	1210: uint16(anon_sym__),
	1211: uint16(24),
	1212: uint16(2),
	1213: uint16(sym_capture),
	1214: uint16(aux_sym_list_repeat1),
	1215: uint16(111),
	1216: uint16(3),
	1217: uint16(anon_sym_STAR),
	1218: uint16(anon_sym_PLUS),
	1219: uint16(anon_sym_QMARK),
	1220: uint16(147),
	1221: uint16(8),
	1223: uint16(anon_sym_DOT),
	1224: uint16(anon_sym_DQUOTE),
	1225: uint16(anon_sym_LBRACK),
	1226: uint16(anon_sym_RBRACK),
	1227: uint16(anon_sym_LPAREN),
	1228: uint16(anon_sym_RPAREN),
	1229: uint16(anon_sym_BANG),
	1230: uint16(7),
	1231: uint16(3),
	1232: uint16(1),
	1233: uint16(sym_comment),
	1234: uint16(115),
	1235: uint16(1),
	1236: uint16(anon_sym_AT),
	1237: uint16(65),
	1238: uint16(1),
	1239: uint16(sym_quantifier),
	1240: uint16(153),
	1241: uint16(2),
	1242: uint16(sym_identifier),
	1243: uint16(anon_sym__),
	1244: uint16(55),
	1245: uint16(2),
	1246: uint16(sym_capture),
	1247: uint16(aux_sym_list_repeat1),
	1248: uint16(111),
	1249: uint16(3),
	1250: uint16(anon_sym_STAR),
	1251: uint16(anon_sym_PLUS),
	1252: uint16(anon_sym_QMARK),
	1253: uint16(151),
	1254: uint16(8),
	1256: uint16(anon_sym_DOT),
	1257: uint16(anon_sym_DQUOTE),
	1258: uint16(anon_sym_LBRACK),
	1259: uint16(anon_sym_RBRACK),
	1260: uint16(anon_sym_LPAREN),
	1261: uint16(anon_sym_RPAREN),
	1262: uint16(anon_sym_BANG),
	1263: uint16(7),
	1264: uint16(3),
	1265: uint16(1),
	1266: uint16(sym_comment),
	1267: uint16(115),
	1268: uint16(1),
	1269: uint16(anon_sym_AT),
	1270: uint16(65),
	1271: uint16(1),
	1272: uint16(sym_quantifier),
	1273: uint16(157),
	1274: uint16(2),
	1275: uint16(sym_identifier),
	1276: uint16(anon_sym__),
	1277: uint16(24),
	1278: uint16(2),
	1279: uint16(sym_capture),
	1280: uint16(aux_sym_list_repeat1),
	1281: uint16(111),
	1282: uint16(3),
	1283: uint16(anon_sym_STAR),
	1284: uint16(anon_sym_PLUS),
	1285: uint16(anon_sym_QMARK),
	1286: uint16(155),
	1287: uint16(8),
	1289: uint16(anon_sym_DOT),
	1290: uint16(anon_sym_DQUOTE),
	1291: uint16(anon_sym_LBRACK),
	1292: uint16(anon_sym_RBRACK),
	1293: uint16(anon_sym_LPAREN),
	1294: uint16(anon_sym_RPAREN),
	1295: uint16(anon_sym_BANG),
	1296: uint16(11),
	1297: uint16(3),
	1298: uint16(1),
	1299: uint16(sym_comment),
	1300: uint16(7),
	1301: uint16(1),
	1302: uint16(sym_identifier),
	1303: uint16(9),
	1304: uint16(1),
	1305: uint16(anon_sym__),
	1306: uint16(11),
	1307: uint16(1),
	1308: uint16(anon_sym_DQUOTE),
	1309: uint16(13),
	1310: uint16(1),
	1311: uint16(anon_sym_LBRACK),
	1312: uint16(15),
	1313: uint16(1),
	1314: uint16(anon_sym_LPAREN),
	1315: uint16(159),
	1316: uint16(1),
	1318: uint16(38),
	1319: uint16(1),
	1320: uint16(sym_string),
	1321: uint16(59),
	1322: uint16(1),
	1323: uint16(sym__field_name),
	1324: uint16(20),
	1325: uint16(2),
	1326: uint16(sym_definition),
	1327: uint16(aux_sym_program_repeat1),
	1328: uint16(68),
	1329: uint16(7),
	1330: uint16(sym_list),
	1331: uint16(sym_grouping),
	1332: uint16(sym_missing_node),
	1333: uint16(sym_anonymous_node),
	1334: uint16(sym_named_node),
	1335: uint16(sym_field_definition),
	1336: uint16(sym_predicate),
	1337: uint16(7),
	1338: uint16(3),
	1339: uint16(1),
	1340: uint16(sym_comment),
	1341: uint16(115),
	1342: uint16(1),
	1343: uint16(anon_sym_AT),
	1344: uint16(65),
	1345: uint16(1),
	1346: uint16(sym_quantifier),
	1347: uint16(163),
	1348: uint16(2),
	1349: uint16(sym_identifier),
	1350: uint16(anon_sym__),
	1351: uint16(40),
	1352: uint16(2),
	1353: uint16(sym_capture),
	1354: uint16(aux_sym_list_repeat1),
	1355: uint16(111),
	1356: uint16(3),
	1357: uint16(anon_sym_STAR),
	1358: uint16(anon_sym_PLUS),
	1359: uint16(anon_sym_QMARK),
	1360: uint16(161),
	1361: uint16(8),
	1363: uint16(anon_sym_DOT),
	1364: uint16(anon_sym_DQUOTE),
	1365: uint16(anon_sym_LBRACK),
	1366: uint16(anon_sym_RBRACK),
	1367: uint16(anon_sym_LPAREN),
	1368: uint16(anon_sym_RPAREN),
	1369: uint16(anon_sym_BANG),
	1370: uint16(7),
	1371: uint16(3),
	1372: uint16(1),
	1373: uint16(sym_comment),
	1374: uint16(115),
	1375: uint16(1),
	1376: uint16(anon_sym_AT),
	1377: uint16(65),
	1378: uint16(1),
	1379: uint16(sym_quantifier),
	1380: uint16(167),
	1381: uint16(2),
	1382: uint16(sym_identifier),
	1383: uint16(anon_sym__),
	1384: uint16(24),
	1385: uint16(2),
	1386: uint16(sym_capture),
	1387: uint16(aux_sym_list_repeat1),
	1388: uint16(111),
	1389: uint16(3),
	1390: uint16(anon_sym_STAR),
	1391: uint16(anon_sym_PLUS),
	1392: uint16(anon_sym_QMARK),
	1393: uint16(165),
	1394: uint16(8),
	1396: uint16(anon_sym_DOT),
	1397: uint16(anon_sym_DQUOTE),
	1398: uint16(anon_sym_LBRACK),
	1399: uint16(anon_sym_RBRACK),
	1400: uint16(anon_sym_LPAREN),
	1401: uint16(anon_sym_RPAREN),
	1402: uint16(anon_sym_BANG),
	1403: uint16(7),
	1404: uint16(3),
	1405: uint16(1),
	1406: uint16(sym_comment),
	1407: uint16(115),
	1408: uint16(1),
	1409: uint16(anon_sym_AT),
	1410: uint16(65),
	1411: uint16(1),
	1412: uint16(sym_quantifier),
	1413: uint16(171),
	1414: uint16(2),
	1415: uint16(sym_identifier),
	1416: uint16(anon_sym__),
	1417: uint16(41),
	1418: uint16(2),
	1419: uint16(sym_capture),
	1420: uint16(aux_sym_list_repeat1),
	1421: uint16(111),
	1422: uint16(3),
	1423: uint16(anon_sym_STAR),
	1424: uint16(anon_sym_PLUS),
	1425: uint16(anon_sym_QMARK),
	1426: uint16(169),
	1427: uint16(8),
	1429: uint16(anon_sym_DOT),
	1430: uint16(anon_sym_DQUOTE),
	1431: uint16(anon_sym_LBRACK),
	1432: uint16(anon_sym_RBRACK),
	1433: uint16(anon_sym_LPAREN),
	1434: uint16(anon_sym_RPAREN),
	1435: uint16(anon_sym_BANG),
	1436: uint16(7),
	1437: uint16(3),
	1438: uint16(1),
	1439: uint16(sym_comment),
	1440: uint16(115),
	1441: uint16(1),
	1442: uint16(anon_sym_AT),
	1443: uint16(65),
	1444: uint16(1),
	1445: uint16(sym_quantifier),
	1446: uint16(175),
	1447: uint16(2),
	1448: uint16(sym_identifier),
	1449: uint16(anon_sym__),
	1450: uint16(24),
	1451: uint16(2),
	1452: uint16(sym_capture),
	1453: uint16(aux_sym_list_repeat1),
	1454: uint16(111),
	1455: uint16(3),
	1456: uint16(anon_sym_STAR),
	1457: uint16(anon_sym_PLUS),
	1458: uint16(anon_sym_QMARK),
	1459: uint16(173),
	1460: uint16(8),
	1462: uint16(anon_sym_DOT),
	1463: uint16(anon_sym_DQUOTE),
	1464: uint16(anon_sym_LBRACK),
	1465: uint16(anon_sym_RBRACK),
	1466: uint16(anon_sym_LPAREN),
	1467: uint16(anon_sym_RPAREN),
	1468: uint16(anon_sym_BANG),
	1469: uint16(7),
	1470: uint16(3),
	1471: uint16(1),
	1472: uint16(sym_comment),
	1473: uint16(115),
	1474: uint16(1),
	1475: uint16(anon_sym_AT),
	1476: uint16(65),
	1477: uint16(1),
	1478: uint16(sym_quantifier),
	1479: uint16(179),
	1480: uint16(2),
	1481: uint16(sym_identifier),
	1482: uint16(anon_sym__),
	1483: uint16(37),
	1484: uint16(2),
	1485: uint16(sym_capture),
	1486: uint16(aux_sym_list_repeat1),
	1487: uint16(111),
	1488: uint16(3),
	1489: uint16(anon_sym_STAR),
	1490: uint16(anon_sym_PLUS),
	1491: uint16(anon_sym_QMARK),
	1492: uint16(177),
	1493: uint16(8),
	1495: uint16(anon_sym_DOT),
	1496: uint16(anon_sym_DQUOTE),
	1497: uint16(anon_sym_LBRACK),
	1498: uint16(anon_sym_RBRACK),
	1499: uint16(anon_sym_LPAREN),
	1500: uint16(anon_sym_RPAREN),
	1501: uint16(anon_sym_BANG),
	1502: uint16(7),
	1503: uint16(3),
	1504: uint16(1),
	1505: uint16(sym_comment),
	1506: uint16(115),
	1507: uint16(1),
	1508: uint16(anon_sym_AT),
	1509: uint16(65),
	1510: uint16(1),
	1511: uint16(sym_quantifier),
	1512: uint16(183),
	1513: uint16(2),
	1514: uint16(sym_identifier),
	1515: uint16(anon_sym__),
	1516: uint16(44),
	1517: uint16(2),
	1518: uint16(sym_capture),
	1519: uint16(aux_sym_list_repeat1),
	1520: uint16(111),
	1521: uint16(3),
	1522: uint16(anon_sym_STAR),
	1523: uint16(anon_sym_PLUS),
	1524: uint16(anon_sym_QMARK),
	1525: uint16(181),
	1526: uint16(8),
	1528: uint16(anon_sym_DOT),
	1529: uint16(anon_sym_DQUOTE),
	1530: uint16(anon_sym_LBRACK),
	1531: uint16(anon_sym_RBRACK),
	1532: uint16(anon_sym_LPAREN),
	1533: uint16(anon_sym_RPAREN),
	1534: uint16(anon_sym_BANG),
	1535: uint16(7),
	1536: uint16(3),
	1537: uint16(1),
	1538: uint16(sym_comment),
	1539: uint16(115),
	1540: uint16(1),
	1541: uint16(anon_sym_AT),
	1542: uint16(65),
	1543: uint16(1),
	1544: uint16(sym_quantifier),
	1545: uint16(187),
	1546: uint16(2),
	1547: uint16(sym_identifier),
	1548: uint16(anon_sym__),
	1549: uint16(24),
	1550: uint16(2),
	1551: uint16(sym_capture),
	1552: uint16(aux_sym_list_repeat1),
	1553: uint16(111),
	1554: uint16(3),
	1555: uint16(anon_sym_STAR),
	1556: uint16(anon_sym_PLUS),
	1557: uint16(anon_sym_QMARK),
	1558: uint16(185),
	1559: uint16(8),
	1561: uint16(anon_sym_DOT),
	1562: uint16(anon_sym_DQUOTE),
	1563: uint16(anon_sym_LBRACK),
	1564: uint16(anon_sym_RBRACK),
	1565: uint16(anon_sym_LPAREN),
	1566: uint16(anon_sym_RPAREN),
	1567: uint16(anon_sym_BANG),
	1568: uint16(7),
	1569: uint16(3),
	1570: uint16(1),
	1571: uint16(sym_comment),
	1572: uint16(115),
	1573: uint16(1),
	1574: uint16(anon_sym_AT),
	1575: uint16(65),
	1576: uint16(1),
	1577: uint16(sym_quantifier),
	1578: uint16(191),
	1579: uint16(2),
	1580: uint16(sym_identifier),
	1581: uint16(anon_sym__),
	1582: uint16(24),
	1583: uint16(2),
	1584: uint16(sym_capture),
	1585: uint16(aux_sym_list_repeat1),
	1586: uint16(111),
	1587: uint16(3),
	1588: uint16(anon_sym_STAR),
	1589: uint16(anon_sym_PLUS),
	1590: uint16(anon_sym_QMARK),
	1591: uint16(189),
	1592: uint16(8),
	1594: uint16(anon_sym_DOT),
	1595: uint16(anon_sym_DQUOTE),
	1596: uint16(anon_sym_LBRACK),
	1597: uint16(anon_sym_RBRACK),
	1598: uint16(anon_sym_LPAREN),
	1599: uint16(anon_sym_RPAREN),
	1600: uint16(anon_sym_BANG),
	1601: uint16(7),
	1602: uint16(3),
	1603: uint16(1),
	1604: uint16(sym_comment),
	1605: uint16(115),
	1606: uint16(1),
	1607: uint16(anon_sym_AT),
	1608: uint16(65),
	1609: uint16(1),
	1610: uint16(sym_quantifier),
	1611: uint16(195),
	1612: uint16(2),
	1613: uint16(sym_identifier),
	1614: uint16(anon_sym__),
	1615: uint16(47),
	1616: uint16(2),
	1617: uint16(sym_capture),
	1618: uint16(aux_sym_list_repeat1),
	1619: uint16(111),
	1620: uint16(3),
	1621: uint16(anon_sym_STAR),
	1622: uint16(anon_sym_PLUS),
	1623: uint16(anon_sym_QMARK),
	1624: uint16(193),
	1625: uint16(8),
	1627: uint16(anon_sym_DOT),
	1628: uint16(anon_sym_DQUOTE),
	1629: uint16(anon_sym_LBRACK),
	1630: uint16(anon_sym_RBRACK),
	1631: uint16(anon_sym_LPAREN),
	1632: uint16(anon_sym_RPAREN),
	1633: uint16(anon_sym_BANG),
	1634: uint16(7),
	1635: uint16(3),
	1636: uint16(1),
	1637: uint16(sym_comment),
	1638: uint16(115),
	1639: uint16(1),
	1640: uint16(anon_sym_AT),
	1641: uint16(65),
	1642: uint16(1),
	1643: uint16(sym_quantifier),
	1644: uint16(199),
	1645: uint16(2),
	1646: uint16(sym_identifier),
	1647: uint16(anon_sym__),
	1648: uint16(48),
	1649: uint16(2),
	1650: uint16(sym_capture),
	1651: uint16(aux_sym_list_repeat1),
	1652: uint16(111),
	1653: uint16(3),
	1654: uint16(anon_sym_STAR),
	1655: uint16(anon_sym_PLUS),
	1656: uint16(anon_sym_QMARK),
	1657: uint16(197),
	1658: uint16(8),
	1660: uint16(anon_sym_DOT),
	1661: uint16(anon_sym_DQUOTE),
	1662: uint16(anon_sym_LBRACK),
	1663: uint16(anon_sym_RBRACK),
	1664: uint16(anon_sym_LPAREN),
	1665: uint16(anon_sym_RPAREN),
	1666: uint16(anon_sym_BANG),
	1667: uint16(7),
	1668: uint16(3),
	1669: uint16(1),
	1670: uint16(sym_comment),
	1671: uint16(115),
	1672: uint16(1),
	1673: uint16(anon_sym_AT),
	1674: uint16(65),
	1675: uint16(1),
	1676: uint16(sym_quantifier),
	1677: uint16(203),
	1678: uint16(2),
	1679: uint16(sym_identifier),
	1680: uint16(anon_sym__),
	1681: uint16(24),
	1682: uint16(2),
	1683: uint16(sym_capture),
	1684: uint16(aux_sym_list_repeat1),
	1685: uint16(111),
	1686: uint16(3),
	1687: uint16(anon_sym_STAR),
	1688: uint16(anon_sym_PLUS),
	1689: uint16(anon_sym_QMARK),
	1690: uint16(201),
	1691: uint16(8),
	1693: uint16(anon_sym_DOT),
	1694: uint16(anon_sym_DQUOTE),
	1695: uint16(anon_sym_LBRACK),
	1696: uint16(anon_sym_RBRACK),
	1697: uint16(anon_sym_LPAREN),
	1698: uint16(anon_sym_RPAREN),
	1699: uint16(anon_sym_BANG),
	1700: uint16(11),
	1701: uint16(3),
	1702: uint16(1),
	1703: uint16(sym_comment),
	1704: uint16(7),
	1705: uint16(1),
	1706: uint16(sym_identifier),
	1707: uint16(9),
	1708: uint16(1),
	1709: uint16(anon_sym__),
	1710: uint16(11),
	1711: uint16(1),
	1712: uint16(anon_sym_DQUOTE),
	1713: uint16(13),
	1714: uint16(1),
	1715: uint16(anon_sym_LBRACK),
	1716: uint16(15),
	1717: uint16(1),
	1718: uint16(anon_sym_LPAREN),
	1719: uint16(205),
	1720: uint16(1),
	1721: uint16(anon_sym_RBRACK),
	1722: uint16(38),
	1723: uint16(1),
	1724: uint16(sym_string),
	1725: uint16(59),
	1726: uint16(1),
	1727: uint16(sym__field_name),
	1728: uint16(20),
	1729: uint16(2),
	1730: uint16(sym_definition),
	1731: uint16(aux_sym_program_repeat1),
	1732: uint16(68),
	1733: uint16(7),
	1734: uint16(sym_list),
	1735: uint16(sym_grouping),
	1736: uint16(sym_missing_node),
	1737: uint16(sym_anonymous_node),
	1738: uint16(sym_named_node),
	1739: uint16(sym_field_definition),
	1740: uint16(sym_predicate),
	1741: uint16(7),
	1742: uint16(3),
	1743: uint16(1),
	1744: uint16(sym_comment),
	1745: uint16(115),
	1746: uint16(1),
	1747: uint16(anon_sym_AT),
	1748: uint16(65),
	1749: uint16(1),
	1750: uint16(sym_quantifier),
	1751: uint16(209),
	1752: uint16(2),
	1753: uint16(sym_identifier),
	1754: uint16(anon_sym__),
	1755: uint16(51),
	1756: uint16(2),
	1757: uint16(sym_capture),
	1758: uint16(aux_sym_list_repeat1),
	1759: uint16(111),
	1760: uint16(3),
	1761: uint16(anon_sym_STAR),
	1762: uint16(anon_sym_PLUS),
	1763: uint16(anon_sym_QMARK),
	1764: uint16(207),
	1765: uint16(8),
	1767: uint16(anon_sym_DOT),
	1768: uint16(anon_sym_DQUOTE),
	1769: uint16(anon_sym_LBRACK),
	1770: uint16(anon_sym_RBRACK),
	1771: uint16(anon_sym_LPAREN),
	1772: uint16(anon_sym_RPAREN),
	1773: uint16(anon_sym_BANG),
	1774: uint16(7),
	1775: uint16(3),
	1776: uint16(1),
	1777: uint16(sym_comment),
	1778: uint16(115),
	1779: uint16(1),
	1780: uint16(anon_sym_AT),
	1781: uint16(65),
	1782: uint16(1),
	1783: uint16(sym_quantifier),
	1784: uint16(213),
	1785: uint16(2),
	1786: uint16(sym_identifier),
	1787: uint16(anon_sym__),
	1788: uint16(24),
	1789: uint16(2),
	1790: uint16(sym_capture),
	1791: uint16(aux_sym_list_repeat1),
	1792: uint16(111),
	1793: uint16(3),
	1794: uint16(anon_sym_STAR),
	1795: uint16(anon_sym_PLUS),
	1796: uint16(anon_sym_QMARK),
	1797: uint16(211),
	1798: uint16(8),
	1800: uint16(anon_sym_DOT),
	1801: uint16(anon_sym_DQUOTE),
	1802: uint16(anon_sym_LBRACK),
	1803: uint16(anon_sym_RBRACK),
	1804: uint16(anon_sym_LPAREN),
	1805: uint16(anon_sym_RPAREN),
	1806: uint16(anon_sym_BANG),
	1807: uint16(7),
	1808: uint16(3),
	1809: uint16(1),
	1810: uint16(sym_comment),
	1811: uint16(115),
	1812: uint16(1),
	1813: uint16(anon_sym_AT),
	1814: uint16(65),
	1815: uint16(1),
	1816: uint16(sym_quantifier),
	1817: uint16(217),
	1818: uint16(2),
	1819: uint16(sym_identifier),
	1820: uint16(anon_sym__),
	1821: uint16(24),
	1822: uint16(2),
	1823: uint16(sym_capture),
	1824: uint16(aux_sym_list_repeat1),
	1825: uint16(111),
	1826: uint16(3),
	1827: uint16(anon_sym_STAR),
	1828: uint16(anon_sym_PLUS),
	1829: uint16(anon_sym_QMARK),
	1830: uint16(215),
	1831: uint16(8),
	1833: uint16(anon_sym_DOT),
	1834: uint16(anon_sym_DQUOTE),
	1835: uint16(anon_sym_LBRACK),
	1836: uint16(anon_sym_RBRACK),
	1837: uint16(anon_sym_LPAREN),
	1838: uint16(anon_sym_RPAREN),
	1839: uint16(anon_sym_BANG),
	1840: uint16(7),
	1841: uint16(3),
	1842: uint16(1),
	1843: uint16(sym_comment),
	1844: uint16(115),
	1845: uint16(1),
	1846: uint16(anon_sym_AT),
	1847: uint16(65),
	1848: uint16(1),
	1849: uint16(sym_quantifier),
	1850: uint16(221),
	1851: uint16(2),
	1852: uint16(sym_identifier),
	1853: uint16(anon_sym__),
	1854: uint16(52),
	1855: uint16(2),
	1856: uint16(sym_capture),
	1857: uint16(aux_sym_list_repeat1),
	1858: uint16(111),
	1859: uint16(3),
	1860: uint16(anon_sym_STAR),
	1861: uint16(anon_sym_PLUS),
	1862: uint16(anon_sym_QMARK),
	1863: uint16(219),
	1864: uint16(8),
	1866: uint16(anon_sym_DOT),
	1867: uint16(anon_sym_DQUOTE),
	1868: uint16(anon_sym_LBRACK),
	1869: uint16(anon_sym_RBRACK),
	1870: uint16(anon_sym_LPAREN),
	1871: uint16(anon_sym_RPAREN),
	1872: uint16(anon_sym_BANG),
	1873: uint16(7),
	1874: uint16(3),
	1875: uint16(1),
	1876: uint16(sym_comment),
	1877: uint16(115),
	1878: uint16(1),
	1879: uint16(anon_sym_AT),
	1880: uint16(65),
	1881: uint16(1),
	1882: uint16(sym_quantifier),
	1883: uint16(225),
	1884: uint16(2),
	1885: uint16(sym_identifier),
	1886: uint16(anon_sym__),
	1887: uint16(53),
	1888: uint16(2),
	1889: uint16(sym_capture),
	1890: uint16(aux_sym_list_repeat1),
	1891: uint16(111),
	1892: uint16(3),
	1893: uint16(anon_sym_STAR),
	1894: uint16(anon_sym_PLUS),
	1895: uint16(anon_sym_QMARK),
	1896: uint16(223),
	1897: uint16(8),
	1899: uint16(anon_sym_DOT),
	1900: uint16(anon_sym_DQUOTE),
	1901: uint16(anon_sym_LBRACK),
	1902: uint16(anon_sym_RBRACK),
	1903: uint16(anon_sym_LPAREN),
	1904: uint16(anon_sym_RPAREN),
	1905: uint16(anon_sym_BANG),
	1906: uint16(7),
	1907: uint16(3),
	1908: uint16(1),
	1909: uint16(sym_comment),
	1910: uint16(115),
	1911: uint16(1),
	1912: uint16(anon_sym_AT),
	1913: uint16(65),
	1914: uint16(1),
	1915: uint16(sym_quantifier),
	1916: uint16(229),
	1917: uint16(2),
	1918: uint16(sym_identifier),
	1919: uint16(anon_sym__),
	1920: uint16(24),
	1921: uint16(2),
	1922: uint16(sym_capture),
	1923: uint16(aux_sym_list_repeat1),
	1924: uint16(111),
	1925: uint16(3),
	1926: uint16(anon_sym_STAR),
	1927: uint16(anon_sym_PLUS),
	1928: uint16(anon_sym_QMARK),
	1929: uint16(227),
	1930: uint16(8),
	1932: uint16(anon_sym_DOT),
	1933: uint16(anon_sym_DQUOTE),
	1934: uint16(anon_sym_LBRACK),
	1935: uint16(anon_sym_RBRACK),
	1936: uint16(anon_sym_LPAREN),
	1937: uint16(anon_sym_RPAREN),
	1938: uint16(anon_sym_BANG),
	1939: uint16(7),
	1940: uint16(3),
	1941: uint16(1),
	1942: uint16(sym_comment),
	1943: uint16(115),
	1944: uint16(1),
	1945: uint16(anon_sym_AT),
	1946: uint16(65),
	1947: uint16(1),
	1948: uint16(sym_quantifier),
	1949: uint16(233),
	1950: uint16(2),
	1951: uint16(sym_identifier),
	1952: uint16(anon_sym__),
	1953: uint16(24),
	1954: uint16(2),
	1955: uint16(sym_capture),
	1956: uint16(aux_sym_list_repeat1),
	1957: uint16(111),
	1958: uint16(3),
	1959: uint16(anon_sym_STAR),
	1960: uint16(anon_sym_PLUS),
	1961: uint16(anon_sym_QMARK),
	1962: uint16(231),
	1963: uint16(8),
	1965: uint16(anon_sym_DOT),
	1966: uint16(anon_sym_DQUOTE),
	1967: uint16(anon_sym_LBRACK),
	1968: uint16(anon_sym_RBRACK),
	1969: uint16(anon_sym_LPAREN),
	1970: uint16(anon_sym_RPAREN),
	1971: uint16(anon_sym_BANG),
	1972: uint16(7),
	1973: uint16(3),
	1974: uint16(1),
	1975: uint16(sym_comment),
	1976: uint16(115),
	1977: uint16(1),
	1978: uint16(anon_sym_AT),
	1979: uint16(65),
	1980: uint16(1),
	1981: uint16(sym_quantifier),
	1982: uint16(237),
	1983: uint16(2),
	1984: uint16(sym_identifier),
	1985: uint16(anon_sym__),
	1986: uint16(24),
	1987: uint16(2),
	1988: uint16(sym_capture),
	1989: uint16(aux_sym_list_repeat1),
	1990: uint16(111),
	1991: uint16(3),
	1992: uint16(anon_sym_STAR),
	1993: uint16(anon_sym_PLUS),
	1994: uint16(anon_sym_QMARK),
	1995: uint16(235),
	1996: uint16(8),
	1998: uint16(anon_sym_DOT),
	1999: uint16(anon_sym_DQUOTE),
	2000: uint16(anon_sym_LBRACK),
	2001: uint16(anon_sym_RBRACK),
	2002: uint16(anon_sym_LPAREN),
	2003: uint16(anon_sym_RPAREN),
	2004: uint16(anon_sym_BANG),
	2005: uint16(7),
	2006: uint16(3),
	2007: uint16(1),
	2008: uint16(sym_comment),
	2009: uint16(115),
	2010: uint16(1),
	2011: uint16(anon_sym_AT),
	2012: uint16(65),
	2013: uint16(1),
	2014: uint16(sym_quantifier),
	2015: uint16(241),
	2016: uint16(2),
	2017: uint16(sym_identifier),
	2018: uint16(anon_sym__),
	2019: uint16(23),
	2020: uint16(2),
	2021: uint16(sym_capture),
	2022: uint16(aux_sym_list_repeat1),
	2023: uint16(111),
	2024: uint16(3),
	2025: uint16(anon_sym_STAR),
	2026: uint16(anon_sym_PLUS),
	2027: uint16(anon_sym_QMARK),
	2028: uint16(239),
	2029: uint16(8),
	2031: uint16(anon_sym_DOT),
	2032: uint16(anon_sym_DQUOTE),
	2033: uint16(anon_sym_LBRACK),
	2034: uint16(anon_sym_RBRACK),
	2035: uint16(anon_sym_LPAREN),
	2036: uint16(anon_sym_RPAREN),
	2037: uint16(anon_sym_BANG),
	2038: uint16(7),
	2039: uint16(3),
	2040: uint16(1),
	2041: uint16(sym_comment),
	2042: uint16(115),
	2043: uint16(1),
	2044: uint16(anon_sym_AT),
	2045: uint16(65),
	2046: uint16(1),
	2047: uint16(sym_quantifier),
	2048: uint16(245),
	2049: uint16(2),
	2050: uint16(sym_identifier),
	2051: uint16(anon_sym__),
	2052: uint16(24),
	2053: uint16(2),
	2054: uint16(sym_capture),
	2055: uint16(aux_sym_list_repeat1),
	2056: uint16(111),
	2057: uint16(3),
	2058: uint16(anon_sym_STAR),
	2059: uint16(anon_sym_PLUS),
	2060: uint16(anon_sym_QMARK),
	2061: uint16(243),
	2062: uint16(8),
	2064: uint16(anon_sym_DOT),
	2065: uint16(anon_sym_DQUOTE),
	2066: uint16(anon_sym_LBRACK),
	2067: uint16(anon_sym_RBRACK),
	2068: uint16(anon_sym_LPAREN),
	2069: uint16(anon_sym_RPAREN),
	2070: uint16(anon_sym_BANG),
	2071: uint16(10),
	2072: uint16(3),
	2073: uint16(1),
	2074: uint16(sym_comment),
	2075: uint16(7),
	2076: uint16(1),
	2077: uint16(sym_identifier),
	2078: uint16(9),
	2079: uint16(1),
	2080: uint16(anon_sym__),
	2081: uint16(11),
	2082: uint16(1),
	2083: uint16(anon_sym_DQUOTE),
	2084: uint16(13),
	2085: uint16(1),
	2086: uint16(anon_sym_LBRACK),
	2087: uint16(15),
	2088: uint16(1),
	2089: uint16(anon_sym_LPAREN),
	2090: uint16(38),
	2091: uint16(1),
	2092: uint16(sym_string),
	2093: uint16(59),
	2094: uint16(1),
	2095: uint16(sym__field_name),
	2096: uint16(45),
	2097: uint16(2),
	2098: uint16(sym_definition),
	2099: uint16(aux_sym_program_repeat1),
	2100: uint16(68),
	2101: uint16(7),
	2102: uint16(sym_list),
	2103: uint16(sym_grouping),
	2104: uint16(sym_missing_node),
	2105: uint16(sym_anonymous_node),
	2106: uint16(sym_named_node),
	2107: uint16(sym_field_definition),
	2108: uint16(sym_predicate),
	2109: uint16(10),
	2110: uint16(3),
	2111: uint16(1),
	2112: uint16(sym_comment),
	2113: uint16(7),
	2114: uint16(1),
	2115: uint16(sym_identifier),
	2116: uint16(11),
	2117: uint16(1),
	2118: uint16(anon_sym_DQUOTE),
	2119: uint16(13),
	2120: uint16(1),
	2121: uint16(anon_sym_LBRACK),
	2122: uint16(15),
	2123: uint16(1),
	2124: uint16(anon_sym_LPAREN),
	2125: uint16(71),
	2126: uint16(1),
	2127: uint16(anon_sym__),
	2128: uint16(57),
	2129: uint16(1),
	2130: uint16(sym__field_name),
	2131: uint16(66),
	2132: uint16(1),
	2133: uint16(sym_string),
	2134: uint16(67),
	2135: uint16(1),
	2136: uint16(sym_definition),
	2137: uint16(68),
	2138: uint16(7),
	2139: uint16(sym_list),
	2140: uint16(sym_grouping),
	2141: uint16(sym_missing_node),
	2142: uint16(sym_anonymous_node),
	2143: uint16(sym_named_node),
	2144: uint16(sym_field_definition),
	2145: uint16(sym_predicate),
	2146: uint16(7),
	2147: uint16(3),
	2148: uint16(1),
	2149: uint16(sym_comment),
	2150: uint16(115),
	2151: uint16(1),
	2152: uint16(anon_sym_AT),
	2153: uint16(65),
	2154: uint16(1),
	2155: uint16(sym_quantifier),
	2156: uint16(249),
	2157: uint16(2),
	2158: uint16(sym_identifier),
	2159: uint16(anon_sym__),
	2160: uint16(37),
	2161: uint16(2),
	2162: uint16(sym_capture),
	2163: uint16(aux_sym_list_repeat1),
	2164: uint16(111),
	2165: uint16(3),
	2166: uint16(anon_sym_STAR),
	2167: uint16(anon_sym_PLUS),
	2168: uint16(anon_sym_QMARK),
	2169: uint16(247),
	2170: uint16(6),
	2171: uint16(anon_sym_DOT),
	2172: uint16(anon_sym_DQUOTE),
	2173: uint16(anon_sym_LBRACK),
	2174: uint16(anon_sym_LPAREN),
	2175: uint16(anon_sym_RPAREN),
	2176: uint16(anon_sym_BANG),
	2177: uint16(10),
	2178: uint16(3),
	2179: uint16(1),
	2180: uint16(sym_comment),
	2181: uint16(7),
	2182: uint16(1),
	2183: uint16(sym_identifier),
	2184: uint16(9),
	2185: uint16(1),
	2186: uint16(anon_sym__),
	2187: uint16(11),
	2188: uint16(1),
	2189: uint16(anon_sym_DQUOTE),
	2190: uint16(13),
	2191: uint16(1),
	2192: uint16(anon_sym_LBRACK),
	2193: uint16(15),
	2194: uint16(1),
	2195: uint16(anon_sym_LPAREN),
	2196: uint16(38),
	2197: uint16(1),
	2198: uint16(sym_string),
	2199: uint16(59),
	2200: uint16(1),
	2201: uint16(sym__field_name),
	2202: uint16(67),
	2203: uint16(1),
	2204: uint16(sym_definition),
	2205: uint16(68),
	2206: uint16(7),
	2207: uint16(sym_list),
	2208: uint16(sym_grouping),
	2209: uint16(sym_missing_node),
	2210: uint16(sym_anonymous_node),
	2211: uint16(sym_named_node),
	2212: uint16(sym_field_definition),
	2213: uint16(sym_predicate),
	2214: uint16(3),
	2215: uint16(3),
	2216: uint16(1),
	2217: uint16(sym_comment),
	2218: uint16(253),
	2219: uint16(2),
	2220: uint16(sym_identifier),
	2221: uint16(anon_sym__),
	2222: uint16(251),
	2223: uint16(12),
	2225: uint16(anon_sym_DOT),
	2226: uint16(anon_sym_STAR),
	2227: uint16(anon_sym_PLUS),
	2228: uint16(anon_sym_QMARK),
	2229: uint16(anon_sym_AT),
	2230: uint16(anon_sym_DQUOTE),
	2231: uint16(anon_sym_LBRACK),
	2232: uint16(anon_sym_RBRACK),
	2233: uint16(anon_sym_LPAREN),
	2234: uint16(anon_sym_RPAREN),
	2235: uint16(anon_sym_BANG),
	2236: uint16(3),
	2237: uint16(3),
	2238: uint16(1),
	2239: uint16(sym_comment),
	2240: uint16(257),
	2241: uint16(2),
	2242: uint16(sym_identifier),
	2243: uint16(anon_sym__),
	2244: uint16(255),
	2245: uint16(12),
	2247: uint16(anon_sym_DOT),
	2248: uint16(anon_sym_STAR),
	2249: uint16(anon_sym_PLUS),
	2250: uint16(anon_sym_QMARK),
	2251: uint16(anon_sym_AT),
	2252: uint16(anon_sym_DQUOTE),
	2253: uint16(anon_sym_LBRACK),
	2254: uint16(anon_sym_RBRACK),
	2255: uint16(anon_sym_LPAREN),
	2256: uint16(anon_sym_RPAREN),
	2257: uint16(anon_sym_BANG),
	2258: uint16(3),
	2259: uint16(3),
	2260: uint16(1),
	2261: uint16(sym_comment),
	2262: uint16(261),
	2263: uint16(2),
	2264: uint16(sym_identifier),
	2265: uint16(anon_sym__),
	2266: uint16(259),
	2267: uint16(12),
	2269: uint16(anon_sym_DOT),
	2270: uint16(anon_sym_STAR),
	2271: uint16(anon_sym_PLUS),
	2272: uint16(anon_sym_QMARK),
	2273: uint16(anon_sym_AT),
	2274: uint16(anon_sym_DQUOTE),
	2275: uint16(anon_sym_LBRACK),
	2276: uint16(anon_sym_RBRACK),
	2277: uint16(anon_sym_LPAREN),
	2278: uint16(anon_sym_RPAREN),
	2279: uint16(anon_sym_BANG),
	2280: uint16(3),
	2281: uint16(3),
	2282: uint16(1),
	2283: uint16(sym_comment),
	2284: uint16(265),
	2285: uint16(2),
	2286: uint16(sym_identifier),
	2287: uint16(anon_sym__),
	2288: uint16(263),
	2289: uint16(12),
	2291: uint16(anon_sym_DOT),
	2292: uint16(anon_sym_STAR),
	2293: uint16(anon_sym_PLUS),
	2294: uint16(anon_sym_QMARK),
	2295: uint16(anon_sym_AT),
	2296: uint16(anon_sym_DQUOTE),
	2297: uint16(anon_sym_LBRACK),
	2298: uint16(anon_sym_RBRACK),
	2299: uint16(anon_sym_LPAREN),
	2300: uint16(anon_sym_RPAREN),
	2301: uint16(anon_sym_BANG),
	2302: uint16(7),
	2303: uint16(3),
	2304: uint16(1),
	2305: uint16(sym_comment),
	2306: uint16(115),
	2307: uint16(1),
	2308: uint16(anon_sym_AT),
	2309: uint16(65),
	2310: uint16(1),
	2311: uint16(sym_quantifier),
	2312: uint16(175),
	2313: uint16(2),
	2314: uint16(sym_identifier),
	2315: uint16(anon_sym__),
	2316: uint16(24),
	2317: uint16(2),
	2318: uint16(sym_capture),
	2319: uint16(aux_sym_list_repeat1),
	2320: uint16(111),
	2321: uint16(3),
	2322: uint16(anon_sym_STAR),
	2323: uint16(anon_sym_PLUS),
	2324: uint16(anon_sym_QMARK),
	2325: uint16(173),
	2326: uint16(5),
	2327: uint16(anon_sym_DOT),
	2328: uint16(anon_sym_DQUOTE),
	2329: uint16(anon_sym_LBRACK),
	2330: uint16(anon_sym_LPAREN),
	2331: uint16(anon_sym_RPAREN),
	2332: uint16(3),
	2333: uint16(3),
	2334: uint16(1),
	2335: uint16(sym_comment),
	2336: uint16(269),
	2337: uint16(2),
	2338: uint16(sym_identifier),
	2339: uint16(anon_sym__),
	2340: uint16(267),
	2341: uint16(12),
	2343: uint16(anon_sym_DOT),
	2344: uint16(anon_sym_STAR),
	2345: uint16(anon_sym_PLUS),
	2346: uint16(anon_sym_QMARK),
	2347: uint16(anon_sym_AT),
	2348: uint16(anon_sym_DQUOTE),
	2349: uint16(anon_sym_LBRACK),
	2350: uint16(anon_sym_RBRACK),
	2351: uint16(anon_sym_LPAREN),
	2352: uint16(anon_sym_RPAREN),
	2353: uint16(anon_sym_BANG),
	2354: uint16(7),
	2355: uint16(3),
	2356: uint16(1),
	2357: uint16(sym_comment),
	2358: uint16(115),
	2359: uint16(1),
	2360: uint16(anon_sym_AT),
	2361: uint16(65),
	2362: uint16(1),
	2363: uint16(sym_quantifier),
	2364: uint16(179),
	2365: uint16(2),
	2366: uint16(sym_identifier),
	2367: uint16(anon_sym__),
	2368: uint16(64),
	2369: uint16(2),
	2370: uint16(sym_capture),
	2371: uint16(aux_sym_list_repeat1),
	2372: uint16(111),
	2373: uint16(3),
	2374: uint16(anon_sym_STAR),
	2375: uint16(anon_sym_PLUS),
	2376: uint16(anon_sym_QMARK),
	2377: uint16(177),
	2378: uint16(5),
	2379: uint16(anon_sym_DOT),
	2380: uint16(anon_sym_DQUOTE),
	2381: uint16(anon_sym_LBRACK),
	2382: uint16(anon_sym_LPAREN),
	2383: uint16(anon_sym_RPAREN),
	2384: uint16(3),
	2385: uint16(3),
	2386: uint16(1),
	2387: uint16(sym_comment),
	2388: uint16(273),
	2389: uint16(2),
	2390: uint16(sym_identifier),
	2391: uint16(anon_sym__),
	2392: uint16(271),
	2393: uint16(8),
	2395: uint16(anon_sym_DOT),
	2396: uint16(anon_sym_DQUOTE),
	2397: uint16(anon_sym_LBRACK),
	2398: uint16(anon_sym_RBRACK),
	2399: uint16(anon_sym_LPAREN),
	2400: uint16(anon_sym_RPAREN),
	2401: uint16(anon_sym_BANG),
	2402: uint16(3),
	2403: uint16(3),
	2404: uint16(1),
	2405: uint16(sym_comment),
	2406: uint16(277),
	2407: uint16(2),
	2408: uint16(sym_identifier),
	2409: uint16(anon_sym__),
	2410: uint16(275),
	2411: uint16(8),
	2413: uint16(anon_sym_DOT),
	2414: uint16(anon_sym_DQUOTE),
	2415: uint16(anon_sym_LBRACK),
	2416: uint16(anon_sym_RBRACK),
	2417: uint16(anon_sym_LPAREN),
	2418: uint16(anon_sym_RPAREN),
	2419: uint16(anon_sym_BANG),
	2420: uint16(3),
	2421: uint16(3),
	2422: uint16(1),
	2423: uint16(sym_comment),
	2424: uint16(281),
	2425: uint16(2),
	2426: uint16(sym_identifier),
	2427: uint16(anon_sym__),
	2428: uint16(279),
	2429: uint16(8),
	2431: uint16(anon_sym_DOT),
	2432: uint16(anon_sym_DQUOTE),
	2433: uint16(anon_sym_LBRACK),
	2434: uint16(anon_sym_RBRACK),
	2435: uint16(anon_sym_LPAREN),
	2436: uint16(anon_sym_RPAREN),
	2437: uint16(anon_sym_BANG),
	2438: uint16(5),
	2439: uint16(3),
	2440: uint16(1),
	2441: uint16(sym_comment),
	2442: uint16(283),
	2443: uint16(1),
	2444: uint16(anon_sym_SLASH),
	2445: uint16(285),
	2446: uint16(1),
	2447: uint16(anon_sym_COLON),
	2448: uint16(249),
	2449: uint16(2),
	2450: uint16(sym_identifier),
	2451: uint16(anon_sym__),
	2452: uint16(247),
	2453: uint16(6),
	2454: uint16(anon_sym_DOT),
	2455: uint16(anon_sym_DQUOTE),
	2456: uint16(anon_sym_LBRACK),
	2457: uint16(anon_sym_LPAREN),
	2458: uint16(anon_sym_RPAREN),
	2459: uint16(anon_sym_BANG),
	2460: uint16(3),
	2461: uint16(3),
	2462: uint16(1),
	2463: uint16(sym_comment),
	2464: uint16(289),
	2465: uint16(2),
	2466: uint16(sym_identifier),
	2467: uint16(anon_sym__),
	2468: uint16(287),
	2469: uint16(8),
	2471: uint16(anon_sym_DOT),
	2472: uint16(anon_sym_DQUOTE),
	2473: uint16(anon_sym_LBRACK),
	2474: uint16(anon_sym_RBRACK),
	2475: uint16(anon_sym_LPAREN),
	2476: uint16(anon_sym_RPAREN),
	2477: uint16(anon_sym_BANG),
	2478: uint16(4),
	2479: uint16(3),
	2480: uint16(1),
	2481: uint16(sym_comment),
	2482: uint16(291),
	2483: uint16(1),
	2484: uint16(anon_sym_DOT),
	2485: uint16(293),
	2486: uint16(2),
	2487: uint16(sym_identifier),
	2488: uint16(anon_sym__),
	2489: uint16(295),
	2490: uint16(5),
	2491: uint16(anon_sym_DQUOTE),
	2492: uint16(anon_sym_LBRACK),
	2493: uint16(anon_sym_LPAREN),
	2494: uint16(anon_sym_RPAREN),
	2495: uint16(anon_sym_BANG),
	2496: uint16(4),
	2497: uint16(3),
	2498: uint16(1),
	2499: uint16(sym_comment),
	2500: uint16(297),
	2501: uint16(1),
	2502: uint16(anon_sym_DOT),
	2503: uint16(293),
	2504: uint16(2),
	2505: uint16(sym_identifier),
	2506: uint16(anon_sym__),
	2507: uint16(295),
	2508: uint16(5),
	2509: uint16(anon_sym_DQUOTE),
	2510: uint16(anon_sym_LBRACK),
	2511: uint16(anon_sym_LPAREN),
	2512: uint16(anon_sym_RPAREN),
	2513: uint16(anon_sym_BANG),
	2514: uint16(7),
	2515: uint16(3),
	2516: uint16(1),
	2517: uint16(sym_comment),
	2518: uint16(299),
	2519: uint16(1),
	2520: uint16(sym_identifier),
	2521: uint16(301),
	2522: uint16(1),
	2523: uint16(anon_sym_AT),
	2524: uint16(303),
	2525: uint16(1),
	2526: uint16(anon_sym_DQUOTE),
	2527: uint16(305),
	2528: uint16(1),
	2529: uint16(anon_sym_RPAREN),
	2530: uint16(106),
	2531: uint16(1),
	2532: uint16(sym_parameters),
	2533: uint16(87),
	2534: uint16(3),
	2535: uint16(sym_capture),
	2536: uint16(sym_string),
	2537: uint16(aux_sym_parameters_repeat1),
	2538: uint16(3),
	2539: uint16(3),
	2540: uint16(1),
	2541: uint16(sym_comment),
	2542: uint16(309),
	2543: uint16(2),
	2544: uint16(sym_identifier),
	2545: uint16(anon_sym__),
	2546: uint16(307),
	2547: uint16(6),
	2548: uint16(anon_sym_DOT),
	2549: uint16(anon_sym_DQUOTE),
	2550: uint16(anon_sym_LBRACK),
	2551: uint16(anon_sym_LPAREN),
	2552: uint16(anon_sym_RPAREN),
	2553: uint16(anon_sym_BANG),
	2554: uint16(4),
	2555: uint16(3),
	2556: uint16(1),
	2557: uint16(sym_comment),
	2558: uint16(311),
	2559: uint16(1),
	2560: uint16(anon_sym_DOT),
	2561: uint16(293),
	2562: uint16(2),
	2563: uint16(sym_identifier),
	2564: uint16(anon_sym__),
	2565: uint16(295),
	2566: uint16(5),
	2567: uint16(anon_sym_DQUOTE),
	2568: uint16(anon_sym_LBRACK),
	2569: uint16(anon_sym_LPAREN),
	2570: uint16(anon_sym_RPAREN),
	2571: uint16(anon_sym_BANG),
	2572: uint16(3),
	2573: uint16(3),
	2574: uint16(1),
	2575: uint16(sym_comment),
	2576: uint16(315),
	2577: uint16(2),
	2578: uint16(sym_identifier),
	2579: uint16(anon_sym__),
	2580: uint16(313),
	2581: uint16(6),
	2582: uint16(anon_sym_DOT),
	2583: uint16(anon_sym_DQUOTE),
	2584: uint16(anon_sym_LBRACK),
	2585: uint16(anon_sym_LPAREN),
	2586: uint16(anon_sym_RPAREN),
	2587: uint16(anon_sym_BANG),
	2588: uint16(4),
	2589: uint16(3),
	2590: uint16(1),
	2591: uint16(sym_comment),
	2592: uint16(317),
	2593: uint16(1),
	2594: uint16(anon_sym_DOT),
	2595: uint16(293),
	2596: uint16(2),
	2597: uint16(sym_identifier),
	2598: uint16(anon_sym__),
	2599: uint16(295),
	2600: uint16(5),
	2601: uint16(anon_sym_DQUOTE),
	2602: uint16(anon_sym_LBRACK),
	2603: uint16(anon_sym_LPAREN),
	2604: uint16(anon_sym_RPAREN),
	2605: uint16(anon_sym_BANG),
	2606: uint16(3),
	2607: uint16(3),
	2608: uint16(1),
	2609: uint16(sym_comment),
	2610: uint16(321),
	2611: uint16(2),
	2612: uint16(sym_identifier),
	2613: uint16(anon_sym__),
	2614: uint16(319),
	2615: uint16(6),
	2616: uint16(anon_sym_DOT),
	2617: uint16(anon_sym_DQUOTE),
	2618: uint16(anon_sym_LBRACK),
	2619: uint16(anon_sym_LPAREN),
	2620: uint16(anon_sym_RPAREN),
	2621: uint16(anon_sym_BANG),
	2622: uint16(4),
	2623: uint16(3),
	2624: uint16(1),
	2625: uint16(sym_comment),
	2626: uint16(323),
	2627: uint16(1),
	2628: uint16(anon_sym_DOT),
	2629: uint16(293),
	2630: uint16(2),
	2631: uint16(sym_identifier),
	2632: uint16(anon_sym__),
	2633: uint16(295),
	2634: uint16(5),
	2635: uint16(anon_sym_DQUOTE),
	2636: uint16(anon_sym_LBRACK),
	2637: uint16(anon_sym_LPAREN),
	2638: uint16(anon_sym_RPAREN),
	2639: uint16(anon_sym_BANG),
	2640: uint16(4),
	2641: uint16(3),
	2642: uint16(1),
	2643: uint16(sym_comment),
	2644: uint16(325),
	2645: uint16(1),
	2646: uint16(anon_sym_DOT),
	2647: uint16(293),
	2648: uint16(2),
	2649: uint16(sym_identifier),
	2650: uint16(anon_sym__),
	2651: uint16(295),
	2652: uint16(5),
	2653: uint16(anon_sym_DQUOTE),
	2654: uint16(anon_sym_LBRACK),
	2655: uint16(anon_sym_LPAREN),
	2656: uint16(anon_sym_RPAREN),
	2657: uint16(anon_sym_BANG),
	2658: uint16(4),
	2659: uint16(3),
	2660: uint16(1),
	2661: uint16(sym_comment),
	2662: uint16(327),
	2663: uint16(1),
	2664: uint16(anon_sym_DOT),
	2665: uint16(293),
	2666: uint16(2),
	2667: uint16(sym_identifier),
	2668: uint16(anon_sym__),
	2669: uint16(295),
	2670: uint16(5),
	2671: uint16(anon_sym_DQUOTE),
	2672: uint16(anon_sym_LBRACK),
	2673: uint16(anon_sym_LPAREN),
	2674: uint16(anon_sym_RPAREN),
	2675: uint16(anon_sym_BANG),
	2676: uint16(3),
	2677: uint16(3),
	2678: uint16(1),
	2679: uint16(sym_comment),
	2680: uint16(331),
	2681: uint16(2),
	2682: uint16(sym_identifier),
	2683: uint16(anon_sym__),
	2684: uint16(329),
	2685: uint16(6),
	2686: uint16(anon_sym_DOT),
	2687: uint16(anon_sym_DQUOTE),
	2688: uint16(anon_sym_LBRACK),
	2689: uint16(anon_sym_LPAREN),
	2690: uint16(anon_sym_RPAREN),
	2691: uint16(anon_sym_BANG),
	2692: uint16(6),
	2693: uint16(3),
	2694: uint16(1),
	2695: uint16(sym_comment),
	2696: uint16(333),
	2697: uint16(1),
	2698: uint16(sym_identifier),
	2699: uint16(336),
	2700: uint16(1),
	2701: uint16(anon_sym_AT),
	2702: uint16(339),
	2703: uint16(1),
	2704: uint16(anon_sym_DQUOTE),
	2705: uint16(342),
	2706: uint16(1),
	2707: uint16(anon_sym_RPAREN),
	2708: uint16(84),
	2709: uint16(3),
	2710: uint16(sym_capture),
	2711: uint16(sym_string),
	2712: uint16(aux_sym_parameters_repeat1),
	2713: uint16(3),
	2714: uint16(3),
	2715: uint16(1),
	2716: uint16(sym_comment),
	2717: uint16(346),
	2718: uint16(2),
	2719: uint16(sym_identifier),
	2720: uint16(anon_sym__),
	2721: uint16(344),
	2722: uint16(5),
	2723: uint16(anon_sym_DOT),
	2724: uint16(anon_sym_DQUOTE),
	2725: uint16(anon_sym_LBRACK),
	2726: uint16(anon_sym_LPAREN),
	2727: uint16(anon_sym_RPAREN),
	2728: uint16(4),
	2729: uint16(3),
	2730: uint16(1),
	2731: uint16(sym_comment),
	2732: uint16(348),
	2733: uint16(1),
	2734: uint16(anon_sym_DOT),
	2735: uint16(350),
	2736: uint16(2),
	2737: uint16(sym_identifier),
	2738: uint16(anon_sym__),
	2739: uint16(352),
	2740: uint16(4),
	2741: uint16(anon_sym_DQUOTE),
	2742: uint16(anon_sym_LBRACK),
	2743: uint16(anon_sym_LPAREN),
	2744: uint16(anon_sym_RPAREN),
	2745: uint16(6),
	2746: uint16(3),
	2747: uint16(1),
	2748: uint16(sym_comment),
	2749: uint16(301),
	2750: uint16(1),
	2751: uint16(anon_sym_AT),
	2752: uint16(303),
	2753: uint16(1),
	2754: uint16(anon_sym_DQUOTE),
	2755: uint16(354),
	2756: uint16(1),
	2757: uint16(sym_identifier),
	2758: uint16(356),
	2759: uint16(1),
	2760: uint16(anon_sym_RPAREN),
	2761: uint16(84),
	2762: uint16(3),
	2763: uint16(sym_capture),
	2764: uint16(sym_string),
	2765: uint16(aux_sym_parameters_repeat1),
	2766: uint16(5),
	2767: uint16(360),
	2768: uint16(1),
	2769: uint16(anon_sym_DQUOTE2),
	2770: uint16(362),
	2771: uint16(1),
	2772: uint16(sym_comment),
	2773: uint16(92),
	2774: uint16(1),
	2775: uint16(aux_sym_string_content_repeat1),
	2776: uint16(104),
	2777: uint16(1),
	2778: uint16(sym_string_content),
	2779: uint16(358),
	2780: uint16(2),
	2781: uint16(sym_escape_sequence),
	2782: uint16(aux_sym_string_content_token1),
	2783: uint16(3),
	2784: uint16(3),
	2785: uint16(1),
	2786: uint16(sym_comment),
	2787: uint16(364),
	2788: uint16(2),
	2789: uint16(sym_identifier),
	2790: uint16(anon_sym__),
	2791: uint16(366),
	2792: uint16(3),
	2793: uint16(anon_sym_DQUOTE),
	2794: uint16(anon_sym_LBRACK),
	2795: uint16(anon_sym_LPAREN),
	2796: uint16(5),
	2797: uint16(362),
	2798: uint16(1),
	2799: uint16(sym_comment),
	2800: uint16(368),
	2801: uint16(1),
	2802: uint16(anon_sym_DQUOTE2),
	2803: uint16(92),
	2804: uint16(1),
	2805: uint16(aux_sym_string_content_repeat1),
	2806: uint16(110),
	2807: uint16(1),
	2808: uint16(sym_string_content),
	2809: uint16(358),
	2810: uint16(2),
	2811: uint16(sym_escape_sequence),
	2812: uint16(aux_sym_string_content_token1),
	2813: uint16(5),
	2814: uint16(362),
	2815: uint16(1),
	2816: uint16(sym_comment),
	2817: uint16(370),
	2818: uint16(1),
	2819: uint16(anon_sym_DQUOTE2),
	2820: uint16(92),
	2821: uint16(1),
	2822: uint16(aux_sym_string_content_repeat1),
	2823: uint16(108),
	2824: uint16(1),
	2825: uint16(sym_string_content),
	2826: uint16(358),
	2827: uint16(2),
	2828: uint16(sym_escape_sequence),
	2829: uint16(aux_sym_string_content_token1),
	2830: uint16(4),
	2831: uint16(362),
	2832: uint16(1),
	2833: uint16(sym_comment),
	2834: uint16(374),
	2835: uint16(1),
	2836: uint16(anon_sym_DQUOTE2),
	2837: uint16(93),
	2838: uint16(1),
	2839: uint16(aux_sym_string_content_repeat1),
	2840: uint16(372),
	2841: uint16(2),
	2842: uint16(sym_escape_sequence),
	2843: uint16(aux_sym_string_content_token1),
	2844: uint16(4),
	2845: uint16(362),
	2846: uint16(1),
	2847: uint16(sym_comment),
	2848: uint16(379),
	2849: uint16(1),
	2850: uint16(anon_sym_DQUOTE2),
	2851: uint16(93),
	2852: uint16(1),
	2853: uint16(aux_sym_string_content_repeat1),
	2854: uint16(376),
	2855: uint16(2),
	2856: uint16(sym_escape_sequence),
	2857: uint16(aux_sym_string_content_token1),
	2858: uint16(5),
	2859: uint16(3),
	2860: uint16(1),
	2861: uint16(sym_comment),
	2862: uint16(11),
	2863: uint16(1),
	2864: uint16(anon_sym_DQUOTE),
	2865: uint16(381),
	2866: uint16(1),
	2867: uint16(sym_identifier),
	2868: uint16(383),
	2869: uint16(1),
	2870: uint16(anon_sym_RPAREN),
	2871: uint16(109),
	2872: uint16(1),
	2873: uint16(sym_string),
	2874: uint16(2),
	2875: uint16(3),
	2876: uint16(1),
	2877: uint16(sym_comment),
	2878: uint16(251),
	2879: uint16(4),
	2880: uint16(sym_identifier),
	2881: uint16(anon_sym_AT),
	2882: uint16(anon_sym_DQUOTE),
	2883: uint16(anon_sym_RPAREN),
	2884: uint16(4),
	2885: uint16(3),
	2886: uint16(1),
	2887: uint16(sym_comment),
	2888: uint16(385),
	2889: uint16(1),
	2890: uint16(aux_sym__immediate_identifier_token1),
	2891: uint16(387),
	2892: uint16(1),
	2893: uint16(anon_sym_DQUOTE2),
	2894: uint16(4),
	2895: uint16(2),
	2896: uint16(sym__immediate_identifier),
	2897: uint16(sym__immediate_string),
	2898: uint16(2),
	2899: uint16(3),
	2900: uint16(1),
	2901: uint16(sym_comment),
	2902: uint16(255),
	2903: uint16(4),
	2904: uint16(sym_identifier),
	2905: uint16(anon_sym_AT),
	2906: uint16(anon_sym_DQUOTE),
	2907: uint16(anon_sym_RPAREN),
	2908: uint16(2),
	2909: uint16(3),
	2910: uint16(1),
	2911: uint16(sym_comment),
	2912: uint16(259),
	2913: uint16(4),
	2914: uint16(sym_identifier),
	2915: uint16(anon_sym_AT),
	2916: uint16(anon_sym_DQUOTE),
	2917: uint16(anon_sym_RPAREN),
	2918: uint16(3),
	2919: uint16(3),
	2920: uint16(1),
	2921: uint16(sym_comment),
	2922: uint16(389),
	2923: uint16(1),
	2924: uint16(aux_sym__immediate_identifier_token1),
	2925: uint16(107),
	2926: uint16(1),
	2927: uint16(sym__immediate_identifier),
	2928: uint16(3),
	2929: uint16(3),
	2930: uint16(1),
	2931: uint16(sym_comment),
	2932: uint16(391),
	2933: uint16(1),
	2934: uint16(aux_sym__immediate_identifier_token1),
	2935: uint16(61),
	2936: uint16(1),
	2937: uint16(sym__immediate_identifier),
	2938: uint16(3),
	2939: uint16(3),
	2940: uint16(1),
	2941: uint16(sym_comment),
	2942: uint16(393),
	2943: uint16(1),
	2944: uint16(aux_sym__immediate_identifier_token1),
	2945: uint16(97),
	2946: uint16(1),
	2947: uint16(sym__immediate_identifier),
	2948: uint16(2),
	2949: uint16(3),
	2950: uint16(1),
	2951: uint16(sym_comment),
	2952: uint16(395),
	2953: uint16(1),
	2954: uint16(sym_identifier),
	2955: uint16(2),
	2956: uint16(3),
	2957: uint16(1),
	2958: uint16(sym_comment),
	2959: uint16(285),
	2960: uint16(1),
	2961: uint16(anon_sym_COLON),
	2962: uint16(2),
	2963: uint16(3),
	2964: uint16(1),
	2965: uint16(sym_comment),
	2966: uint16(397),
	2967: uint16(1),
	2968: uint16(anon_sym_DQUOTE2),
	2969: uint16(2),
	2970: uint16(3),
	2971: uint16(1),
	2972: uint16(sym_comment),
	2973: uint16(399),
	2974: uint16(1),
	2976: uint16(2),
	2977: uint16(3),
	2978: uint16(1),
	2979: uint16(sym_comment),
	2980: uint16(401),
	2981: uint16(1),
	2982: uint16(anon_sym_RPAREN),
	2983: uint16(2),
	2984: uint16(3),
	2985: uint16(1),
	2986: uint16(sym_comment),
	2987: uint16(403),
	2988: uint16(1),
	2989: uint16(sym_predicate_type),
	2990: uint16(2),
	2991: uint16(3),
	2992: uint16(1),
	2993: uint16(sym_comment),
	2994: uint16(405),
	2995: uint16(1),
	2996: uint16(anon_sym_DQUOTE2),
	2997: uint16(2),
	2998: uint16(3),
	2999: uint16(1),
	3000: uint16(sym_comment),
	3001: uint16(407),
	3002: uint16(1),
	3003: uint16(anon_sym_RPAREN),
	3004: uint16(2),
	3005: uint16(3),
	3006: uint16(1),
	3007: uint16(sym_comment),
	3008: uint16(409),
	3009: uint16(1),
	3010: uint16(anon_sym_DQUOTE2),
}

var ts_small_parse_table_map = [109]uint32_t{
	1:   uint32(51),
	2:   uint32(102),
	3:   uint32(153),
	4:   uint32(201),
	5:   uint32(249),
	6:   uint32(297),
	7:   uint32(345),
	8:   uint32(393),
	9:   uint32(438),
	10:  uint32(483),
	11:  uint32(528),
	12:  uint32(573),
	13:  uint32(618),
	14:  uint32(663),
	15:  uint32(708),
	16:  uint32(753),
	17:  uint32(797),
	18:  uint32(841),
	19:  uint32(883),
	20:  uint32(925),
	21:  uint32(958),
	22:  uint32(991),
	23:  uint32(1024),
	24:  uint32(1057),
	25:  uint32(1090),
	26:  uint32(1131),
	27:  uint32(1164),
	28:  uint32(1197),
	29:  uint32(1230),
	30:  uint32(1263),
	31:  uint32(1296),
	32:  uint32(1337),
	33:  uint32(1370),
	34:  uint32(1403),
	35:  uint32(1436),
	36:  uint32(1469),
	37:  uint32(1502),
	38:  uint32(1535),
	39:  uint32(1568),
	40:  uint32(1601),
	41:  uint32(1634),
	42:  uint32(1667),
	43:  uint32(1700),
	44:  uint32(1741),
	45:  uint32(1774),
	46:  uint32(1807),
	47:  uint32(1840),
	48:  uint32(1873),
	49:  uint32(1906),
	50:  uint32(1939),
	51:  uint32(1972),
	52:  uint32(2005),
	53:  uint32(2038),
	54:  uint32(2071),
	55:  uint32(2109),
	56:  uint32(2146),
	57:  uint32(2177),
	58:  uint32(2214),
	59:  uint32(2236),
	60:  uint32(2258),
	61:  uint32(2280),
	62:  uint32(2302),
	63:  uint32(2332),
	64:  uint32(2354),
	65:  uint32(2384),
	66:  uint32(2402),
	67:  uint32(2420),
	68:  uint32(2438),
	69:  uint32(2460),
	70:  uint32(2478),
	71:  uint32(2496),
	72:  uint32(2514),
	73:  uint32(2538),
	74:  uint32(2554),
	75:  uint32(2572),
	76:  uint32(2588),
	77:  uint32(2606),
	78:  uint32(2622),
	79:  uint32(2640),
	80:  uint32(2658),
	81:  uint32(2676),
	82:  uint32(2692),
	83:  uint32(2713),
	84:  uint32(2728),
	85:  uint32(2745),
	86:  uint32(2766),
	87:  uint32(2783),
	88:  uint32(2796),
	89:  uint32(2813),
	90:  uint32(2830),
	91:  uint32(2844),
	92:  uint32(2858),
	93:  uint32(2874),
	94:  uint32(2884),
	95:  uint32(2898),
	96:  uint32(2908),
	97:  uint32(2918),
	98:  uint32(2928),
	99:  uint32(2938),
	100: uint32(2948),
	101: uint32(2955),
	102: uint32(2962),
	103: uint32(2969),
	104: uint32(2976),
	105: uint32(2983),
	106: uint32(2990),
	107: uint32(2997),
	108: uint32(3004),
}

var ts_parse_actions = [411]TSParseActionEntry{
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
		Fsymbol:     uint16(sym_program),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(103)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(38)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(28)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(13)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(36)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(39)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
	})))),
	43: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(103)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	44: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	45: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(38)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
	})))),
	49: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	50: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(56)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	54: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
	})))),
	55: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	56: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(49)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(50)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(54)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(103)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(66)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(88)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(56)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(2)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_program_repeat1),
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
		Fcount: uint8(2),
	}})),
	95: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(103)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	98: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(38)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	101: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(88)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(56)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_program_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(2)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(8),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(8),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(21),
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
		Fcount: uint8(1),
	}})),
	120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(10),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(21),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_list_repeat1),
		Fproduction_id: uint16(5),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_list_repeat1),
		Fproduction_id: uint16(5),
	})))),
	125: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(63)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_list_repeat1),
		Fproduction_id: uint16(5),
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
		Fcount:    uint8(2),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_list_repeat1),
		Fproduction_id: uint16(5),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_list),
	})))),
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
		Fcount: uint8(1),
	}})),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_missing_node),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_missing_node),
	})))),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
	}})),
	142: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
	})))),
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
		Fsymbol:      uint16(sym_grouping),
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
		Fcount: uint8(1),
	}})),
	146: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_grouping),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_list),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount: uint8(1),
	}})),
	150: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_list),
		Fproduction_id: uint16(6),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_missing_node),
		Fproduction_id: uint16(7),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_missing_node),
		Fproduction_id: uint16(7),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_missing_node),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount: uint8(1),
	}})),
	158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_missing_node),
		Fproduction_id: uint16(6),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_program),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount: uint8(1),
	}})),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
	})))),
	165: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_grouping),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount: uint8(1),
	}})),
	168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_grouping),
		Fproduction_id: uint16(6),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	174: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_anonymous_node),
		Fproduction_id: uint16(3),
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
		Fcount: uint8(1),
	}})),
	176: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_anonymous_node),
		Fproduction_id: uint16(3),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	178: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_anonymous_node),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount: uint8(1),
	}})),
	180: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_anonymous_node),
		Fproduction_id: uint16(1),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	182: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	186: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(12),
	})))),
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
		Fcount: uint8(1),
	}})),
	188: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(12),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(14),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(14),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
	})))),
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
		Fcount: uint8(1),
	}})),
	196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(15),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(15),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	210: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(16),
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
		Fcount: uint8(1),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(16),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(17),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(17),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(7),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	224: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Fcount: uint8(1),
	}})),
	226: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	228: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(18),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(18),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(19),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(8),
		Fsymbol:        uint16(sym_named_node),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(20),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(20),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(9),
		Fsymbol:        uint16(sym_named_node),
		Fproduction_id: uint16(10),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_missing_node),
		Fproduction_id: uint16(11),
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(sym_missing_node),
		Fproduction_id: uint16(11),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_identifier),
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
		Fcount: uint8(1),
	}})),
	250: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__node_identifier),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_string),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_string),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	256: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_capture),
		Fproduction_id: uint16(4),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_capture),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	260: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	264: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_quantifier),
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
		Fcount: uint8(1),
	}})),
	266: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_quantifier),
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
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_list_repeat1),
		Fproduction_id: uint16(2),
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
		Fcount: uint8(1),
	}})),
	270: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(aux_sym_list_repeat1),
		Fproduction_id: uint16(2),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_field_definition),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(1),
	}})),
	274: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_field_definition),
		Fproduction_id: uint16(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_definition),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_definition),
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
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_predicate),
		Fproduction_id: uint16(13),
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
		Fcount: uint8(1),
	}})),
	282: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(6),
		Fsymbol:        uint16(sym_predicate),
		Fproduction_id: uint16(13),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	288: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_predicate),
		Fproduction_id: uint16(9),
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
		Fcount: uint8(1),
	}})),
	290: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_predicate),
		Fproduction_id: uint16(9),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_named_node_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(101)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__named_node_expression),
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
		Fcount: uint8(1),
	}})),
	310: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__named_node_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__immediate_string),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__immediate_string),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	320: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_negated_field),
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
		Fcount: uint8(1),
	}})),
	322: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_negated_field),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	324: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	328: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__immediate_string),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__immediate_string),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_parameters_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(84)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	338: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(101)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(90)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	343: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__group_expression),
	})))),
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
		Fcount: uint8(1),
	}})),
	347: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__group_expression),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	353: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_grouping_repeat1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(62)),
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
		Fextra: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__field_name),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__field_name),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	375: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_string_content),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	377: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_content_repeat1),
	})))),
	378: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(93)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	380: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_string_content_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
	400: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
	}})))),
}

func tree_sitter_query(tls *libc.TLS) (r uintptr) {
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
	Fname:                      __ccgo_ts + 498,
	Fsupertype_count:           uint32(SUPERTYPE_COUNT),
	Fsupertype_symbols:         uintptr(unsafe.Pointer(&ts_supertype_symbols)),
	Fsupertype_map_slices:      uintptr(unsafe.Pointer(&ts_supertype_map_slices)),
	Fsupertype_map_entries:     uintptr(unsafe.Pointer(&ts_supertype_map_entries)),
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(8),
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

var __ccgo_ts1 = "end\x00.\x00escape_sequence\x00*\x00+\x00?\x00identifier\x00_\x00@\x00\"\x00string_content_token1\x00comment\x00[\x00]\x00(\x00)\x00MISSING\x00/\x00:\x00!\x00#\x00predicate_type\x00program\x00definition\x00_group_expression\x00_named_node_expression\x00quantifier\x00_immediate_identifier\x00_node_identifier\x00capture\x00string\x00string_content\x00parameters\x00list\x00grouping\x00missing_node\x00anonymous_node\x00named_node\x00_field_name\x00field_definition\x00negated_field\x00predicate\x00program_repeat1\x00string_content_repeat1\x00parameters_repeat1\x00list_repeat1\x00grouping_repeat1\x00named_node_repeat1\x00name\x00supertype\x00type\x00query\x00"
