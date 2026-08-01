// Code generated for darwin/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -D__attribute__(...)= -D__extension__= -D_Nonnull= -D_Nullable= -D_Null_unspecified= -DAPI_AVAILABLE(...)= -DAPI_UNAVAILABLE(...)= -DAPI_DEPRECATED(...)= -DAPI_DEPRECATED_WITH_REPLACEMENT(...)= -D__API_AVAILABLE(...)= -D__API_UNAVAILABLE(...)= -D__API_DEPRECATED(...)= -D__API_DEPRECATED_WITH_REPLACEMENT(...)= -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-heex/src -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-heex -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-heex/src/parser.c -o grammar.go', DO NOT EDIT.

//go:build darwin && amd64

package grammar_heex

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 3
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
const MAX_ALIAS_SEQUENCE_LENGTH = 4
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
const PRODUCTION_ID_COUNT = 5
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
const STATE_COUNT = 116
const SUPERTYPE_COUNT = 0
const SV_INTERRUPT = "SA_RESTART"
const SV_NOCLDSTOP = "SA_NOCLDSTOP"
const SV_NODEFER = "SA_NODEFER"
const SV_ONSTACK = "SA_ONSTACK"
const SV_RESETHAND = "SA_RESETHAND"
const SV_SIGINFO = "SA_SIGINFO"
const SYMBOL_COUNT = 87
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
const TOKEN_COUNT = 51
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

const anon_sym_LT_BANG = 1
const anon_sym_DOCTYPE = 2
const anon_sym_html = 3
const anon_sym_GT = 4
const anon_sym_LT = 5
const anon_sym_LT_SLASH = 6
const anon_sym_SLASH_GT = 7
const anon_sym_LT_COLON = 8
const anon_sym_LT_SLASH_COLON = 9
const anon_sym_LBRACE = 10
const anon_sym_RBRACE = 11
const aux_sym__expression_value_token1 = 12
const anon_sym_EQ = 13
const anon_sym_COLONlet = 14
const anon_sym_COLONfor = 15
const anon_sym_COLONkey = 16
const anon_sym_COLONstream = 17
const anon_sym_COLONif = 18
const sym_attribute_value = 19
const anon_sym_SQUOTE = 20
const aux_sym_quoted_attribute_value_token1 = 21
const anon_sym_DQUOTE = 22
const aux_sym_quoted_attribute_value_token2 = 23
const anon_sym_LT_PERCENT = 24
const anon_sym_LT_PERCENT_EQ = 25
const anon_sym_LT_PERCENT_PERCENT = 26
const anon_sym_LT_PERCENT_PERCENT_EQ = 27
const anon_sym_PERCENT_GT = 28
const anon_sym_LT_BANG_DASH_DASH = 29
const aux_sym__html_comment_token1 = 30
const anon_sym_DASH_DASH_GT = 31
const anon_sym_LT_PERCENT_BANG_DASH_DASH = 32
const anon_sym_DASH_DASH_PERCENT_GT = 33
const anon_sym_LT_PERCENT_POUND = 34
const anon_sym_do = 35
const anon_sym_catch = 36
const anon_sym_rescue = 37
const anon_sym_after = 38
const anon_sym_else = 39
const anon_sym_DASH_GT = 40
const anon_sym_POUND = 41
const aux_sym_ending_expression_value_token1 = 42
const anon_sym_DOT = 43
const sym_module = 44
const sym_function = 45
const sym__code = 46
const sym_tag_name = 47
const sym_attribute_name = 48
const sym_text = 49
const sym_entity = 50
const sym_fragment = 51
const sym__node = 52
const sym_doctype = 53
const sym_tag = 54
const sym_component = 55
const sym_slot = 56
const sym_start_tag = 57
const sym_end_tag = 58
const sym_self_closing_tag = 59
const sym_start_component = 60
const sym_end_component = 61
const sym_self_closing_component = 62
const sym_start_slot = 63
const sym_end_slot = 64
const sym_self_closing_slot = 65
const sym_expression = 66
const sym__expression_value = 67
const sym_special_attribute = 68
const sym_special_attribute_name = 69
const sym_attribute = 70
const sym_quoted_attribute_value = 71
const sym_directive = 72
const sym_comment = 73
const sym__html_comment = 74
const sym__bang_comment = 75
const sym__hash_comment = 76
const sym_expression_value = 77
const sym_partial_expression_value = 78
const sym_ending_expression_value = 79
const sym_component_name = 80
const aux_sym_fragment_repeat1 = 81
const aux_sym_component_repeat1 = 82
const aux_sym_start_tag_repeat1 = 83
const aux_sym_expression_repeat1 = 84
const aux_sym__html_comment_repeat1 = 85
const aux_sym__hash_comment_repeat1 = 86
const anon_alias_sym_left = 87
const anon_alias_sym_right = 88
const alias_sym_slot_name = 89

var ts_symbol_names = [90]uintptr{
	0:  __ccgo_ts,
	1:  __ccgo_ts + 4,
	2:  __ccgo_ts + 7,
	3:  __ccgo_ts + 15,
	4:  __ccgo_ts + 20,
	5:  __ccgo_ts + 22,
	6:  __ccgo_ts + 24,
	7:  __ccgo_ts + 27,
	8:  __ccgo_ts + 30,
	9:  __ccgo_ts + 33,
	10: __ccgo_ts + 37,
	11: __ccgo_ts + 39,
	12: __ccgo_ts + 41,
	13: __ccgo_ts + 66,
	14: __ccgo_ts + 68,
	15: __ccgo_ts + 73,
	16: __ccgo_ts + 78,
	17: __ccgo_ts + 83,
	18: __ccgo_ts + 91,
	19: __ccgo_ts + 95,
	20: __ccgo_ts + 111,
	21: __ccgo_ts + 95,
	22: __ccgo_ts + 113,
	23: __ccgo_ts + 95,
	24: __ccgo_ts + 115,
	25: __ccgo_ts + 118,
	26: __ccgo_ts + 122,
	27: __ccgo_ts + 126,
	28: __ccgo_ts + 131,
	29: __ccgo_ts + 134,
	30: __ccgo_ts + 139,
	31: __ccgo_ts + 160,
	32: __ccgo_ts + 164,
	33: __ccgo_ts + 170,
	34: __ccgo_ts + 175,
	35: __ccgo_ts + 179,
	36: __ccgo_ts + 182,
	37: __ccgo_ts + 188,
	38: __ccgo_ts + 195,
	39: __ccgo_ts + 201,
	40: __ccgo_ts + 206,
	41: __ccgo_ts + 209,
	42: __ccgo_ts + 211,
	43: __ccgo_ts + 242,
	44: __ccgo_ts + 244,
	45: __ccgo_ts + 251,
	46: __ccgo_ts + 260,
	47: __ccgo_ts + 266,
	48: __ccgo_ts + 275,
	49: __ccgo_ts + 290,
	50: __ccgo_ts + 295,
	51: __ccgo_ts + 302,
	52: __ccgo_ts + 311,
	53: __ccgo_ts + 317,
	54: __ccgo_ts + 325,
	55: __ccgo_ts + 329,
	56: __ccgo_ts + 339,
	57: __ccgo_ts + 344,
	58: __ccgo_ts + 354,
	59: __ccgo_ts + 362,
	60: __ccgo_ts + 379,
	61: __ccgo_ts + 395,
	62: __ccgo_ts + 409,
	63: __ccgo_ts + 432,
	64: __ccgo_ts + 443,
	65: __ccgo_ts + 452,
	66: __ccgo_ts + 470,
	67: __ccgo_ts + 481,
	68: __ccgo_ts + 499,
	69: __ccgo_ts + 517,
	70: __ccgo_ts + 540,
	71: __ccgo_ts + 550,
	72: __ccgo_ts + 573,
	73: __ccgo_ts + 583,
	74: __ccgo_ts + 591,
	75: __ccgo_ts + 605,
	76: __ccgo_ts + 619,
	77: __ccgo_ts + 633,
	78: __ccgo_ts + 650,
	79: __ccgo_ts + 675,
	80: __ccgo_ts + 699,
	81: __ccgo_ts + 714,
	82: __ccgo_ts + 731,
	83: __ccgo_ts + 749,
	84: __ccgo_ts + 767,
	85: __ccgo_ts + 786,
	86: __ccgo_ts + 808,
	87: __ccgo_ts + 830,
	88: __ccgo_ts + 835,
	89: __ccgo_ts + 841,
}

var ts_symbol_map = [90]TSSymbol{
	1:  uint16(anon_sym_LT_BANG),
	2:  uint16(anon_sym_DOCTYPE),
	3:  uint16(anon_sym_html),
	4:  uint16(anon_sym_GT),
	5:  uint16(anon_sym_LT),
	6:  uint16(anon_sym_LT_SLASH),
	7:  uint16(anon_sym_SLASH_GT),
	8:  uint16(anon_sym_LT_COLON),
	9:  uint16(anon_sym_LT_SLASH_COLON),
	10: uint16(anon_sym_LBRACE),
	11: uint16(anon_sym_RBRACE),
	12: uint16(aux_sym__expression_value_token1),
	13: uint16(anon_sym_EQ),
	14: uint16(anon_sym_COLONlet),
	15: uint16(anon_sym_COLONfor),
	16: uint16(anon_sym_COLONkey),
	17: uint16(anon_sym_COLONstream),
	18: uint16(anon_sym_COLONif),
	19: uint16(sym_attribute_value),
	20: uint16(anon_sym_SQUOTE),
	21: uint16(sym_attribute_value),
	22: uint16(anon_sym_DQUOTE),
	23: uint16(sym_attribute_value),
	24: uint16(anon_sym_LT_PERCENT),
	25: uint16(anon_sym_LT_PERCENT_EQ),
	26: uint16(anon_sym_LT_PERCENT_PERCENT),
	27: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	28: uint16(anon_sym_PERCENT_GT),
	29: uint16(anon_sym_LT_BANG_DASH_DASH),
	30: uint16(aux_sym__html_comment_token1),
	31: uint16(anon_sym_DASH_DASH_GT),
	32: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	33: uint16(anon_sym_DASH_DASH_PERCENT_GT),
	34: uint16(anon_sym_LT_PERCENT_POUND),
	35: uint16(anon_sym_do),
	36: uint16(anon_sym_catch),
	37: uint16(anon_sym_rescue),
	38: uint16(anon_sym_after),
	39: uint16(anon_sym_else),
	40: uint16(anon_sym_DASH_GT),
	41: uint16(anon_sym_POUND),
	42: uint16(aux_sym_ending_expression_value_token1),
	43: uint16(anon_sym_DOT),
	44: uint16(sym_module),
	45: uint16(sym_function),
	46: uint16(sym__code),
	47: uint16(sym_tag_name),
	48: uint16(sym_attribute_name),
	49: uint16(sym_text),
	50: uint16(sym_entity),
	51: uint16(sym_fragment),
	52: uint16(sym__node),
	53: uint16(sym_doctype),
	54: uint16(sym_tag),
	55: uint16(sym_component),
	56: uint16(sym_slot),
	57: uint16(sym_start_tag),
	58: uint16(sym_end_tag),
	59: uint16(sym_self_closing_tag),
	60: uint16(sym_start_component),
	61: uint16(sym_end_component),
	62: uint16(sym_self_closing_component),
	63: uint16(sym_start_slot),
	64: uint16(sym_end_slot),
	65: uint16(sym_self_closing_slot),
	66: uint16(sym_expression),
	67: uint16(sym__expression_value),
	68: uint16(sym_special_attribute),
	69: uint16(sym_special_attribute_name),
	70: uint16(sym_attribute),
	71: uint16(sym_quoted_attribute_value),
	72: uint16(sym_directive),
	73: uint16(sym_comment),
	74: uint16(sym__html_comment),
	75: uint16(sym__bang_comment),
	76: uint16(sym__hash_comment),
	77: uint16(sym_expression_value),
	78: uint16(sym_partial_expression_value),
	79: uint16(sym_ending_expression_value),
	80: uint16(sym_component_name),
	81: uint16(aux_sym_fragment_repeat1),
	82: uint16(aux_sym_component_repeat1),
	83: uint16(aux_sym_start_tag_repeat1),
	84: uint16(aux_sym_expression_repeat1),
	85: uint16(aux_sym__html_comment_repeat1),
	86: uint16(aux_sym__hash_comment_repeat1),
	87: uint16(anon_alias_sym_left),
	88: uint16(anon_alias_sym_right),
	89: uint16(alias_sym_slot_name),
}

var ts_symbol_metadata = [90]TSSymbolMetadata{
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
	12: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	20: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	21: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	30: {},
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
	42: {},
	43: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	56: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	75: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
	81: {},
	82: {},
	83: {},
	84: {},
	85: {},
	86: {},
	87: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	88: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	89: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

var ts_alias_sequences = [5][4]TSSymbol{
	0: {},
	1: {
		0: uint16(anon_alias_sym_left),
		1: uint16(anon_alias_sym_right),
	},
	2: {
		1: uint16(sym_expression_value),
	},
	3: {
		0: uint16(anon_alias_sym_left),
		2: uint16(anon_alias_sym_right),
	},
	4: {
		1: uint16(alias_sym_slot_name),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(aux_sym_expression_repeat1),
	1: uint16(2),
	2: uint16(aux_sym_expression_repeat1),
	3: uint16(sym_expression_value),
}

var ts_primary_state_ids = [116]TSStateId{
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
	63:  uint16(21),
	64:  uint16(64),
	65:  uint16(14),
	66:  uint16(66),
	67:  uint16(67),
	68:  uint16(68),
	69:  uint16(69),
	70:  uint16(70),
	71:  uint16(71),
	72:  uint16(69),
	73:  uint16(71),
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
	86:  uint16(76),
	87:  uint16(87),
	88:  uint16(88),
	89:  uint16(89),
	90:  uint16(90),
	91:  uint16(91),
	92:  uint16(92),
	93:  uint16(57),
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
}

var sym_module_character_set_1 = [9]TSCharacterRange{
	0: {
		Fend: int32(0x08),
	},
	1: {
		Fstart: int32(0x0e),
		Fend:   int32(0x1f),
	},
	2: {
		Fstart: int32('#'),
		Fend:   int32('&'),
	},
	3: {
		Fstart: int32('('),
		Fend:   int32(','),
	},
	4: {
		Fstart: int32('.'),
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

var sym_attribute_name_character_set_2 = [9]TSCharacterRange{
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
	var eof, result, skip, v4 uint8
	var half_size, i, index, mid_index, size uint32_t
	var lookahead1, v3 int32_t
	var range_token, range_token1, v2 uintptr
	var v6 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v2, v3, v4, v6
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
	lookahead1 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
	eof = (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer)
	switch libc.Int32FromUint16(state) {
	case 0:
		if eof != 0 {
			state = uint16(54)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token[i]) == lookahead1 {
				state = map_token[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(0)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead1 == int32('"') {
			state = uint16(81)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(78)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(67)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('<') || int32('>') < lookahead1) && lookahead1 != int32('}') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('"') {
			state = uint16(81)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(82)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('#') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(20)
			goto next_state
		}
		if lookahead1 == int32('.') {
			state = uint16(114)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') {
			state = uint16(115)
			goto next_state
		}
		if int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(143)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead1 == int32('#') {
			state = uint16(47)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(216)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('%') {
			state = uint16(117)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(131)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('d') {
			state = uint16(135)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(133)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(130)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('%') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(131)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('d') {
			state = uint16(135)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(134)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(130)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(119)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('%') {
			state = uint16(122)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('%') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('&') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(61)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(67)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('>') && lookahead1 != int32('}') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('\'') {
			state = uint16(78)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(79)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('-') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('-') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('-') {
			state = uint16(91)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(90)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('-') {
			state = uint16(93)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(92)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('-') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('.') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(19)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead1 == int32('>') {
			state = uint16(58)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(67)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') && lookahead1 != int32('\'') && (lookahead1 < int32('<') || int32('>') < lookahead1) && lookahead1 != int32('}') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('/') {
			state = uint16(19)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(34)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(71)
			goto next_state
		}
		if lookahead1 == int32('>') {
			state = uint16(58)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(67)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') && lookahead1 != int32('\'') && (lookahead1 < int32('<') || int32('>') < lookahead1) && lookahead1 != int32('}') {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32(':') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('>') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('>') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('>') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('>') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('C') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('D') {
			state = uint16(26)
			goto next_state
		}
		if lookahead1 == int32('h') {
			state = uint16(44)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('E') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('O') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('P') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('T') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('Y') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('a') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('e') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('e') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('e') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('f') {
			state = uint16(39)
			goto next_state
		}
		if lookahead1 == int32('i') {
			state = uint16(35)
			goto next_state
		}
		if lookahead1 == int32('k') {
			state = uint16(31)
			goto next_state
		}
		if lookahead1 == int32('l') {
			state = uint16(33)
			goto next_state
		}
		if lookahead1 == int32('s') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('f') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32('l') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('m') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead1 == int32('m') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead1 == int32('o') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead1 == int32('r') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead1 == int32('r') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead1 == int32('t') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead1 == int32('t') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead1 == int32('t') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(45):
		if lookahead1 == int32('y') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(46):
		if lookahead1 == int32('{') {
			state = uint16(67)
			goto next_state
		}
		if lookahead1 == int32('}') {
			state = uint16(68)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(69)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead1 == int32('X') || lookahead1 == int32('x') {
			state = uint16(50)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(182)
			goto next_state
		}
		return result
	case int32(48):
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(48)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('&') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('{') && lookahead1 != int32('}') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(49):
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(49)
			goto next_state
		}
		if int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(50):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(51):
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(52):
		if eof != 0 {
			state = uint16(54)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(59)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(67)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(52)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('>') && lookahead1 != int32('}') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(53):
		if eof != 0 {
			state = uint16(54)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(4)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(60)
			goto next_state
		}
		if lookahead1 == int32('{') {
			state = uint16(67)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(53)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('>') && lookahead1 != int32('}') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOCTYPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_html)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(63)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(62)
			goto next_state
		}
		if lookahead1 == int32(':') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__expression_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(69)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('{') && lookahead1 != int32('}') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__expression_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('{') && lookahead1 != int32('}') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLONlet)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLONfor)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLONkey)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLONstream)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLONif)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_value)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _5
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _5
		_5:
		}
		if v6 && v4 != 0 || lookahead1 == int32('/') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(79)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('\'') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\'') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(82)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_quoted_attribute_value_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(15)
			goto next_state
		}
		if lookahead1 == int32('#') {
			state = uint16(98)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(86)
			goto next_state
		}
		if lookahead1 == int32('=') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_PERCENT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_PERCENT_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('=') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_PERCENT_PERCENT_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__html_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(91)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(90)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__html_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__html_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(93)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(92)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__html_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym__html_comment_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('-') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_DASH_PERCENT_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_PERCENT_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_do)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_do)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_catch)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_catch)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rescue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_rescue)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_after)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_after)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_else)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_ending_expression_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('}') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32(')') || lookahead1 == int32(']') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_ending_expression_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(')') || lookahead1 == int32(']') || lookahead1 == int32('}') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_ending_expression_value_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(')') || lookahead1 == int32(']') || lookahead1 == int32('}') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_module)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(51)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_module_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _10
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _10
		_10:
		}
		if v6 && v4 != 0 {
			state = uint16(115)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_module_character_set_1))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _15
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _15
		_15:
		}
		if v6 && v4 != 0 && lookahead1 != int32('.') {
			state = uint16(116)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(117)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(131)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('d') {
			state = uint16(135)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(133)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(130)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 == int32('-') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32('a') {
			state = uint16(131)
			goto next_state
		}
		if lookahead1 == int32('c') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('d') {
			state = uint16(135)
			goto next_state
		}
		if lookahead1 == int32('e') {
			state = uint16(134)
			goto next_state
		}
		if lookahead1 == int32('r') {
			state = uint16(130)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(119)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('%') {
			state = uint16(122)
			goto next_state
		}
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('>') {
			state = uint16(109)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('>') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(139)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('c') {
			state = uint16(132)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('c') {
			state = uint16(141)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('d') {
			state = uint16(113)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(136)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(108)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(104)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(137)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('f') {
			state = uint16(140)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('h') {
			state = uint16(102)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(138)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(126)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(138)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('o') {
			state = uint16(100)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('r') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(125)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(128)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(124)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('u') {
			state = uint16(129)
			goto next_state
		}
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__code)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && (lookahead1 < int32('\t') || int32('\r') < lookahead1) && lookahead1 != int32(' ') && lookahead1 != int32('%') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(143)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _20
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _20
		_20:
		}
		if v6 && v4 != 0 && lookahead1 != int32('!') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(144):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _25
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _25
		_25:
		}
		if v6 && v4 != 0 && lookahead1 != int32('!') {
			state = uint16(144)
			goto next_state
		}
		return result
	case int32(145):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(173)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(216)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _30
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _30
		_30:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(146):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('C') {
			state = uint16(150)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _35
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _35
		_35:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(147):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(56)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _40
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _40
		_40:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(148):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(146)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _45
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _45
		_45:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(149):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('P') {
			state = uint16(147)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _50
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _50
		_50:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(150):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(151)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _55
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _55
		_55:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(151):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('Y') {
			state = uint16(149)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _60
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _60
		_60:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(152):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('a') {
			state = uint16(170)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _65
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _65
		_65:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(153):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('c') {
			state = uint16(161)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _70
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _70
		_70:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(154):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('c') {
			state = uint16(172)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _75
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _75
		_75:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(155):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('d') {
			state = uint16(111)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _80
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _80
		_80:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(156):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(166)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _85
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _85
		_85:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(157):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(107)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _90
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _90
		_90:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(158):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(103)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _95
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _95
		_95:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(159):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('e') {
			state = uint16(167)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _100
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _100
		_100:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(160):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('f') {
			state = uint16(171)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _105
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _105
		_105:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(161):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('h') {
			state = uint16(101)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _110
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _110
		_110:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(162):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(168)
			goto next_state
		}
		if lookahead1 == int32('n') {
			state = uint16(155)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _115
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _115
		_115:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(163):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('l') {
			state = uint16(57)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _120
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _120
		_120:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('m') {
			state = uint16(163)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _125
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _125
		_125:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('o') {
			state = uint16(99)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _130
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _130
		_130:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('r') {
			state = uint16(105)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _135
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _135
		_135:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(154)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _140
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _140
		_140:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('s') {
			state = uint16(157)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _145
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _145
		_145:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(164)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _150
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _150
		_150:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(153)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _155
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _155
		_155:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('t') {
			state = uint16(156)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _160
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _160
		_160:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('u') {
			state = uint16(158)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _165
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _165
		_165:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('X') || lookahead1 == int32('x') {
			state = uint16(174)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(182)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _170
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _170
		_170:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(187)
			goto next_state
		}
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _175
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _175
		_175:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_attribute_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if v6 = !(eof != 0); v6 {
			v2 = uintptr(unsafe.Pointer(&sym_attribute_name_character_set_2))
			v3 = lookahead1
			index = uint32(0)
			size = uint32(9) - index
			for size > libc.Uint32FromInt32(1) {
				half_size = size / uint32(2)
				mid_index = index + half_size
				range_token = v2 + uintptr(mid_index)*8
				if v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					v4 = libc.BoolUint8(true1 != 0)
					goto _180
				} else {
					if v3 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
						index = mid_index
					}
				}
				size = size - half_size
			}
			range_token1 = v2 + uintptr(index)*8
			v4 = libc.BoolUint8(v3 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v3 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
			goto _180
		_180:
		}
		if v6 && v4 != 0 {
			state = uint16(175)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead1 && lookahead1 <= int32('\r') || lookahead1 == int32(' ') {
			state = uint16(48)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('&') && lookahead1 != int32('<') && lookahead1 != int32('>') && lookahead1 != int32('{') && lookahead1 != int32('}') {
			state = uint16(176)
			goto next_state
		}
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(179)
			goto next_state
		}
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(184)
			goto next_state
		}
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(186)
			goto next_state
		}
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(188)
			goto next_state
		}
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(190)
			goto next_state
		}
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(191)
			goto next_state
		}
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(192)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(193)
			goto next_state
		}
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(194)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(195)
			goto next_state
		}
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(196)
			goto next_state
		}
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(197)
			goto next_state
		}
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(198)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(199)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(200)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(201)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(202)
			goto next_state
		}
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(203)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(204)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(205)
			goto next_state
		}
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(206)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(207)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(208)
			goto next_state
		}
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(209)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(210)
			goto next_state
		}
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(211)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(212)
			goto next_state
		}
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(213)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(214)
			goto next_state
		}
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_entity)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(';') {
			state = uint16(177)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(215)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [38]uint16_t{
	0:  uint16('"'),
	1:  uint16(81),
	2:  uint16('#'),
	3:  uint16(110),
	4:  uint16('&'),
	5:  uint16(145),
	6:  uint16('\''),
	7:  uint16(78),
	8:  uint16('.'),
	9:  uint16(114),
	10: uint16('/'),
	11: uint16(19),
	12: uint16(':'),
	13: uint16(34),
	14: uint16('<'),
	15: uint16(59),
	16: uint16('='),
	17: uint16(71),
	18: uint16('>'),
	19: uint16(58),
	20: uint16('D'),
	21: uint16(148),
	22: uint16('a'),
	23: uint16(160),
	24: uint16('c'),
	25: uint16(152),
	26: uint16('d'),
	27: uint16(165),
	28: uint16('e'),
	29: uint16(162),
	30: uint16('h'),
	31: uint16(169),
	32: uint16('r'),
	33: uint16(159),
	34: uint16('{'),
	35: uint16(67),
	36: uint16('}'),
	37: uint16(68),
}

var ts_lex_modes = [116]TSLexerMode{
	0: {},
	1: {
		Flex_state: uint16(53),
	},
	2: {
		Flex_state: uint16(53),
	},
	3: {
		Flex_state: uint16(53),
	},
	4: {
		Flex_state: uint16(53),
	},
	5: {
		Flex_state: uint16(52),
	},
	6: {
		Flex_state: uint16(53),
	},
	7: {
		Flex_state: uint16(53),
	},
	8: {
		Flex_state: uint16(9),
	},
	9: {
		Flex_state: uint16(9),
	},
	10: {
		Flex_state: uint16(53),
	},
	11: {
		Flex_state: uint16(52),
	},
	12: {
		Flex_state: uint16(52),
	},
	13: {
		Flex_state: uint16(52),
	},
	14: {
		Flex_state: uint16(52),
	},
	15: {
		Flex_state: uint16(52),
	},
	16: {
		Flex_state: uint16(52),
	},
	17: {
		Flex_state: uint16(52),
	},
	18: {
		Flex_state: uint16(52),
	},
	19: {
		Flex_state: uint16(52),
	},
	20: {
		Flex_state: uint16(52),
	},
	21: {
		Flex_state: uint16(52),
	},
	22: {
		Flex_state: uint16(52),
	},
	23: {
		Flex_state: uint16(52),
	},
	24: {
		Flex_state: uint16(52),
	},
	25: {
		Flex_state: uint16(52),
	},
	26: {
		Flex_state: uint16(52),
	},
	27: {
		Flex_state: uint16(52),
	},
	28: {
		Flex_state: uint16(52),
	},
	29: {
		Flex_state: uint16(52),
	},
	30: {
		Flex_state: uint16(52),
	},
	31: {
		Flex_state: uint16(52),
	},
	32: {
		Flex_state: uint16(52),
	},
	33: {
		Flex_state: uint16(52),
	},
	34: {
		Flex_state: uint16(53),
	},
	35: {
		Flex_state: uint16(53),
	},
	36: {
		Flex_state: uint16(53),
	},
	37: {
		Flex_state: uint16(53),
	},
	38: {
		Flex_state: uint16(17),
	},
	39: {
		Flex_state: uint16(17),
	},
	40: {
		Flex_state: uint16(17),
	},
	41: {
		Flex_state: uint16(53),
	},
	42: {
		Flex_state: uint16(17),
	},
	43: {
		Flex_state: uint16(17),
	},
	44: {
		Flex_state: uint16(17),
	},
	45: {
		Flex_state: uint16(53),
	},
	46: {
		Flex_state: uint16(17),
	},
	47: {
		Flex_state: uint16(53),
	},
	48: {
		Flex_state: uint16(53),
	},
	49: {
		Flex_state: uint16(53),
	},
	50: {
		Flex_state: uint16(9),
	},
	51: {
		Flex_state: uint16(9),
	},
	52: {
		Flex_state: uint16(53),
	},
	53: {
		Flex_state: uint16(5),
	},
	54: {
		Flex_state: uint16(16),
	},
	55: {
		Flex_state: uint16(17),
	},
	56: {
		Flex_state: uint16(17),
	},
	57: {
		Flex_state: uint16(6),
	},
	58: {
		Flex_state: uint16(17),
	},
	59: {
		Flex_state: uint16(17),
	},
	60: {
		Flex_state: uint16(17),
	},
	61: {
		Flex_state: uint16(17),
	},
	62: {
		Flex_state: uint16(6),
	},
	63: {
		Flex_state: uint16(17),
	},
	64: {
		Flex_state: uint16(17),
	},
	65: {
		Flex_state: uint16(17),
	},
	66: {
		Flex_state: uint16(1),
	},
	67: {
		Flex_state: uint16(46),
	},
	68: {
		Flex_state: uint16(46),
	},
	69: {
		Flex_state: uint16(46),
	},
	70: {
		Flex_state: uint16(46),
	},
	71: {
		Flex_state: uint16(46),
	},
	72: {
		Flex_state: uint16(46),
	},
	73: {
		Flex_state: uint16(46),
	},
	74: {
		Flex_state: uint16(3),
	},
	75: {
		Flex_state: uint16(46),
	},
	76: {
		Flex_state: uint16(13),
	},
	77: {
		Flex_state: uint16(7),
	},
	78: {
		Flex_state: uint16(7),
	},
	79: {
		Flex_state: uint16(13),
	},
	80: {
		Flex_state: uint16(46),
	},
	81: {
		Flex_state: uint16(14),
	},
	82: {
		Flex_state: uint16(13),
	},
	83: {
		Flex_state: uint16(7),
	},
	84: {
		Flex_state: uint16(7),
	},
	85: {
		Flex_state: uint16(46),
	},
	86: {
		Flex_state: uint16(14),
	},
	87: {
		Flex_state: uint16(7),
	},
	88: {
		Flex_state: uint16(7),
	},
	89: {
		Flex_state: uint16(7),
	},
	90: {
		Flex_state: uint16(7),
	},
	91: {
		Flex_state: uint16(14),
	},
	92: {
		Flex_state: uint16(3),
	},
	93: {
		Flex_state: uint16(7),
	},
	94: {
		Flex_state: uint16(3),
	},
	95: {},
	96: {
		Flex_state: uint16(2),
	},
	97: {
		Flex_state: uint16(3),
	},
	98: {
		Flex_state: uint16(10),
	},
	99: {
		Flex_state: uint16(24),
	},
	100: {
		Flex_state: uint16(3),
	},
	101: {},
	102: {},
	103: {
		Flex_state: uint16(3),
	},
	104: {},
	105: {},
	106: {},
	107: {},
	108: {},
	109: {
		Flex_state: uint16(3),
	},
	110: {},
	111: {
		Flex_state: uint16(49),
	},
	112: {
		Flex_state: uint16(24),
	},
	113: {},
	114: {
		Flex_state: uint16(3),
	},
	115: {
		Flex_state: uint16(49),
	},
}

var ts_parse_table = [2][87]uint16_t{
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
		13: uint16(1),
		14: uint16(1),
		15: uint16(1),
		16: uint16(1),
		17: uint16(1),
		18: uint16(1),
		20: uint16(1),
		22: uint16(1),
		24: uint16(1),
		25: uint16(1),
		26: uint16(1),
		27: uint16(1),
		29: uint16(1),
		32: uint16(1),
		34: uint16(1),
		35: uint16(1),
		36: uint16(1),
		37: uint16(1),
		38: uint16(1),
		39: uint16(1),
		41: uint16(1),
		42: uint16(1),
		43: uint16(1),
		48: uint16(1),
		50: uint16(1),
	},
	1: {
		0:  uint16(3),
		1:  uint16(5),
		5:  uint16(7),
		10: uint16(9),
		24: uint16(11),
		25: uint16(13),
		26: uint16(11),
		27: uint16(13),
		29: uint16(15),
		32: uint16(17),
		34: uint16(19),
		49: uint16(21),
		50: uint16(21),
		51: uint16(102),
		52: uint16(10),
		53: uint16(10),
		54: uint16(10),
		55: uint16(10),
		57: uint16(6),
		59: uint16(17),
		60: uint16(2),
		62: uint16(18),
		66: uint16(10),
		72: uint16(10),
		73: uint16(10),
		74: uint16(28),
		75: uint16(28),
		76: uint16(28),
		81: uint16(10),
	},
}

var ts_small_parse_table = [2123]uint16_t{
	0:    uint16(20),
	1:    uint16(5),
	2:    uint16(1),
	3:    uint16(anon_sym_LT_BANG),
	4:    uint16(7),
	5:    uint16(1),
	6:    uint16(anon_sym_LT),
	7:    uint16(9),
	8:    uint16(1),
	9:    uint16(anon_sym_LBRACE),
	10:   uint16(15),
	11:   uint16(1),
	12:   uint16(anon_sym_LT_BANG_DASH_DASH),
	13:   uint16(17),
	14:   uint16(1),
	15:   uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	16:   uint16(19),
	17:   uint16(1),
	18:   uint16(anon_sym_LT_PERCENT_POUND),
	19:   uint16(23),
	20:   uint16(1),
	21:   uint16(anon_sym_LT_SLASH),
	22:   uint16(25),
	23:   uint16(1),
	24:   uint16(anon_sym_LT_COLON),
	25:   uint16(2),
	26:   uint16(1),
	27:   uint16(sym_start_component),
	28:   uint16(6),
	29:   uint16(1),
	30:   uint16(sym_start_tag),
	31:   uint16(8),
	32:   uint16(1),
	33:   uint16(sym_start_slot),
	34:   uint16(16),
	35:   uint16(1),
	36:   uint16(sym_end_component),
	37:   uint16(17),
	38:   uint16(1),
	39:   uint16(sym_self_closing_tag),
	40:   uint16(18),
	41:   uint16(1),
	42:   uint16(sym_self_closing_component),
	43:   uint16(35),
	44:   uint16(1),
	45:   uint16(sym_self_closing_slot),
	46:   uint16(11),
	47:   uint16(2),
	48:   uint16(anon_sym_LT_PERCENT),
	49:   uint16(anon_sym_LT_PERCENT_PERCENT),
	50:   uint16(13),
	51:   uint16(2),
	52:   uint16(anon_sym_LT_PERCENT_EQ),
	53:   uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	54:   uint16(27),
	55:   uint16(2),
	56:   uint16(sym_text),
	57:   uint16(sym_entity),
	58:   uint16(28),
	59:   uint16(3),
	60:   uint16(sym__html_comment),
	61:   uint16(sym__bang_comment),
	62:   uint16(sym__hash_comment),
	63:   uint16(3),
	64:   uint16(9),
	65:   uint16(sym__node),
	66:   uint16(sym_doctype),
	67:   uint16(sym_tag),
	68:   uint16(sym_component),
	69:   uint16(sym_slot),
	70:   uint16(sym_expression),
	71:   uint16(sym_directive),
	72:   uint16(sym_comment),
	73:   uint16(aux_sym_component_repeat1),
	74:   uint16(20),
	75:   uint16(5),
	76:   uint16(1),
	77:   uint16(anon_sym_LT_BANG),
	78:   uint16(7),
	79:   uint16(1),
	80:   uint16(anon_sym_LT),
	81:   uint16(9),
	82:   uint16(1),
	83:   uint16(anon_sym_LBRACE),
	84:   uint16(15),
	85:   uint16(1),
	86:   uint16(anon_sym_LT_BANG_DASH_DASH),
	87:   uint16(17),
	88:   uint16(1),
	89:   uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	90:   uint16(19),
	91:   uint16(1),
	92:   uint16(anon_sym_LT_PERCENT_POUND),
	93:   uint16(23),
	94:   uint16(1),
	95:   uint16(anon_sym_LT_SLASH),
	96:   uint16(25),
	97:   uint16(1),
	98:   uint16(anon_sym_LT_COLON),
	99:   uint16(2),
	100:  uint16(1),
	101:  uint16(sym_start_component),
	102:  uint16(6),
	103:  uint16(1),
	104:  uint16(sym_start_tag),
	105:  uint16(8),
	106:  uint16(1),
	107:  uint16(sym_start_slot),
	108:  uint16(17),
	109:  uint16(1),
	110:  uint16(sym_self_closing_tag),
	111:  uint16(18),
	112:  uint16(1),
	113:  uint16(sym_self_closing_component),
	114:  uint16(27),
	115:  uint16(1),
	116:  uint16(sym_end_component),
	117:  uint16(35),
	118:  uint16(1),
	119:  uint16(sym_self_closing_slot),
	120:  uint16(11),
	121:  uint16(2),
	122:  uint16(anon_sym_LT_PERCENT),
	123:  uint16(anon_sym_LT_PERCENT_PERCENT),
	124:  uint16(13),
	125:  uint16(2),
	126:  uint16(anon_sym_LT_PERCENT_EQ),
	127:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	128:  uint16(29),
	129:  uint16(2),
	130:  uint16(sym_text),
	131:  uint16(sym_entity),
	132:  uint16(28),
	133:  uint16(3),
	134:  uint16(sym__html_comment),
	135:  uint16(sym__bang_comment),
	136:  uint16(sym__hash_comment),
	137:  uint16(4),
	138:  uint16(9),
	139:  uint16(sym__node),
	140:  uint16(sym_doctype),
	141:  uint16(sym_tag),
	142:  uint16(sym_component),
	143:  uint16(sym_slot),
	144:  uint16(sym_expression),
	145:  uint16(sym_directive),
	146:  uint16(sym_comment),
	147:  uint16(aux_sym_component_repeat1),
	148:  uint16(19),
	149:  uint16(31),
	150:  uint16(1),
	151:  uint16(anon_sym_LT_BANG),
	152:  uint16(34),
	153:  uint16(1),
	154:  uint16(anon_sym_LT),
	155:  uint16(37),
	156:  uint16(1),
	157:  uint16(anon_sym_LT_SLASH),
	158:  uint16(39),
	159:  uint16(1),
	160:  uint16(anon_sym_LT_COLON),
	161:  uint16(42),
	162:  uint16(1),
	163:  uint16(anon_sym_LBRACE),
	164:  uint16(51),
	165:  uint16(1),
	166:  uint16(anon_sym_LT_BANG_DASH_DASH),
	167:  uint16(54),
	168:  uint16(1),
	169:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	170:  uint16(57),
	171:  uint16(1),
	172:  uint16(anon_sym_LT_PERCENT_POUND),
	173:  uint16(2),
	174:  uint16(1),
	175:  uint16(sym_start_component),
	176:  uint16(6),
	177:  uint16(1),
	178:  uint16(sym_start_tag),
	179:  uint16(8),
	180:  uint16(1),
	181:  uint16(sym_start_slot),
	182:  uint16(17),
	183:  uint16(1),
	184:  uint16(sym_self_closing_tag),
	185:  uint16(18),
	186:  uint16(1),
	187:  uint16(sym_self_closing_component),
	188:  uint16(35),
	189:  uint16(1),
	190:  uint16(sym_self_closing_slot),
	191:  uint16(45),
	192:  uint16(2),
	193:  uint16(anon_sym_LT_PERCENT),
	194:  uint16(anon_sym_LT_PERCENT_PERCENT),
	195:  uint16(48),
	196:  uint16(2),
	197:  uint16(anon_sym_LT_PERCENT_EQ),
	198:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	199:  uint16(60),
	200:  uint16(2),
	201:  uint16(sym_text),
	202:  uint16(sym_entity),
	203:  uint16(28),
	204:  uint16(3),
	205:  uint16(sym__html_comment),
	206:  uint16(sym__bang_comment),
	207:  uint16(sym__hash_comment),
	208:  uint16(4),
	209:  uint16(9),
	210:  uint16(sym__node),
	211:  uint16(sym_doctype),
	212:  uint16(sym_tag),
	213:  uint16(sym_component),
	214:  uint16(sym_slot),
	215:  uint16(sym_expression),
	216:  uint16(sym_directive),
	217:  uint16(sym_comment),
	218:  uint16(aux_sym_component_repeat1),
	219:  uint16(17),
	220:  uint16(65),
	221:  uint16(1),
	222:  uint16(anon_sym_LT_BANG),
	223:  uint16(68),
	224:  uint16(1),
	225:  uint16(anon_sym_LT),
	226:  uint16(71),
	227:  uint16(1),
	228:  uint16(anon_sym_LT_SLASH),
	229:  uint16(73),
	230:  uint16(1),
	231:  uint16(anon_sym_LBRACE),
	232:  uint16(82),
	233:  uint16(1),
	234:  uint16(anon_sym_LT_BANG_DASH_DASH),
	235:  uint16(85),
	236:  uint16(1),
	237:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	238:  uint16(88),
	239:  uint16(1),
	240:  uint16(anon_sym_LT_PERCENT_POUND),
	241:  uint16(2),
	242:  uint16(1),
	243:  uint16(sym_start_component),
	244:  uint16(6),
	245:  uint16(1),
	246:  uint16(sym_start_tag),
	247:  uint16(17),
	248:  uint16(1),
	249:  uint16(sym_self_closing_tag),
	250:  uint16(18),
	251:  uint16(1),
	252:  uint16(sym_self_closing_component),
	253:  uint16(63),
	254:  uint16(2),
	256:  uint16(anon_sym_LT_SLASH_COLON),
	257:  uint16(76),
	258:  uint16(2),
	259:  uint16(anon_sym_LT_PERCENT),
	260:  uint16(anon_sym_LT_PERCENT_PERCENT),
	261:  uint16(79),
	262:  uint16(2),
	263:  uint16(anon_sym_LT_PERCENT_EQ),
	264:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	265:  uint16(91),
	266:  uint16(2),
	267:  uint16(sym_text),
	268:  uint16(sym_entity),
	269:  uint16(28),
	270:  uint16(3),
	271:  uint16(sym__html_comment),
	272:  uint16(sym__bang_comment),
	273:  uint16(sym__hash_comment),
	274:  uint16(5),
	275:  uint16(8),
	276:  uint16(sym__node),
	277:  uint16(sym_doctype),
	278:  uint16(sym_tag),
	279:  uint16(sym_component),
	280:  uint16(sym_expression),
	281:  uint16(sym_directive),
	282:  uint16(sym_comment),
	283:  uint16(aux_sym_fragment_repeat1),
	284:  uint16(17),
	285:  uint16(5),
	286:  uint16(1),
	287:  uint16(anon_sym_LT_BANG),
	288:  uint16(7),
	289:  uint16(1),
	290:  uint16(anon_sym_LT),
	291:  uint16(9),
	292:  uint16(1),
	293:  uint16(anon_sym_LBRACE),
	294:  uint16(15),
	295:  uint16(1),
	296:  uint16(anon_sym_LT_BANG_DASH_DASH),
	297:  uint16(17),
	298:  uint16(1),
	299:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	300:  uint16(19),
	301:  uint16(1),
	302:  uint16(anon_sym_LT_PERCENT_POUND),
	303:  uint16(94),
	304:  uint16(1),
	305:  uint16(anon_sym_LT_SLASH),
	306:  uint16(2),
	307:  uint16(1),
	308:  uint16(sym_start_component),
	309:  uint16(6),
	310:  uint16(1),
	311:  uint16(sym_start_tag),
	312:  uint16(13),
	313:  uint16(1),
	314:  uint16(sym_end_tag),
	315:  uint16(17),
	316:  uint16(1),
	317:  uint16(sym_self_closing_tag),
	318:  uint16(18),
	319:  uint16(1),
	320:  uint16(sym_self_closing_component),
	321:  uint16(11),
	322:  uint16(2),
	323:  uint16(anon_sym_LT_PERCENT),
	324:  uint16(anon_sym_LT_PERCENT_PERCENT),
	325:  uint16(13),
	326:  uint16(2),
	327:  uint16(anon_sym_LT_PERCENT_EQ),
	328:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	329:  uint16(96),
	330:  uint16(2),
	331:  uint16(sym_text),
	332:  uint16(sym_entity),
	333:  uint16(28),
	334:  uint16(3),
	335:  uint16(sym__html_comment),
	336:  uint16(sym__bang_comment),
	337:  uint16(sym__hash_comment),
	338:  uint16(7),
	339:  uint16(8),
	340:  uint16(sym__node),
	341:  uint16(sym_doctype),
	342:  uint16(sym_tag),
	343:  uint16(sym_component),
	344:  uint16(sym_expression),
	345:  uint16(sym_directive),
	346:  uint16(sym_comment),
	347:  uint16(aux_sym_fragment_repeat1),
	348:  uint16(17),
	349:  uint16(5),
	350:  uint16(1),
	351:  uint16(anon_sym_LT_BANG),
	352:  uint16(7),
	353:  uint16(1),
	354:  uint16(anon_sym_LT),
	355:  uint16(9),
	356:  uint16(1),
	357:  uint16(anon_sym_LBRACE),
	358:  uint16(15),
	359:  uint16(1),
	360:  uint16(anon_sym_LT_BANG_DASH_DASH),
	361:  uint16(17),
	362:  uint16(1),
	363:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	364:  uint16(19),
	365:  uint16(1),
	366:  uint16(anon_sym_LT_PERCENT_POUND),
	367:  uint16(94),
	368:  uint16(1),
	369:  uint16(anon_sym_LT_SLASH),
	370:  uint16(2),
	371:  uint16(1),
	372:  uint16(sym_start_component),
	373:  uint16(6),
	374:  uint16(1),
	375:  uint16(sym_start_tag),
	376:  uint16(17),
	377:  uint16(1),
	378:  uint16(sym_self_closing_tag),
	379:  uint16(18),
	380:  uint16(1),
	381:  uint16(sym_self_closing_component),
	382:  uint16(25),
	383:  uint16(1),
	384:  uint16(sym_end_tag),
	385:  uint16(11),
	386:  uint16(2),
	387:  uint16(anon_sym_LT_PERCENT),
	388:  uint16(anon_sym_LT_PERCENT_PERCENT),
	389:  uint16(13),
	390:  uint16(2),
	391:  uint16(anon_sym_LT_PERCENT_EQ),
	392:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	393:  uint16(98),
	394:  uint16(2),
	395:  uint16(sym_text),
	396:  uint16(sym_entity),
	397:  uint16(28),
	398:  uint16(3),
	399:  uint16(sym__html_comment),
	400:  uint16(sym__bang_comment),
	401:  uint16(sym__hash_comment),
	402:  uint16(5),
	403:  uint16(8),
	404:  uint16(sym__node),
	405:  uint16(sym_doctype),
	406:  uint16(sym_tag),
	407:  uint16(sym_component),
	408:  uint16(sym_expression),
	409:  uint16(sym_directive),
	410:  uint16(sym_comment),
	411:  uint16(aux_sym_fragment_repeat1),
	412:  uint16(17),
	413:  uint16(5),
	414:  uint16(1),
	415:  uint16(anon_sym_LT_BANG),
	416:  uint16(7),
	417:  uint16(1),
	418:  uint16(anon_sym_LT),
	419:  uint16(9),
	420:  uint16(1),
	421:  uint16(anon_sym_LBRACE),
	422:  uint16(15),
	423:  uint16(1),
	424:  uint16(anon_sym_LT_BANG_DASH_DASH),
	425:  uint16(17),
	426:  uint16(1),
	427:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	428:  uint16(19),
	429:  uint16(1),
	430:  uint16(anon_sym_LT_PERCENT_POUND),
	431:  uint16(100),
	432:  uint16(1),
	433:  uint16(anon_sym_LT_SLASH_COLON),
	434:  uint16(2),
	435:  uint16(1),
	436:  uint16(sym_start_component),
	437:  uint16(6),
	438:  uint16(1),
	439:  uint16(sym_start_tag),
	440:  uint16(17),
	441:  uint16(1),
	442:  uint16(sym_self_closing_tag),
	443:  uint16(18),
	444:  uint16(1),
	445:  uint16(sym_self_closing_component),
	446:  uint16(36),
	447:  uint16(1),
	448:  uint16(sym_end_slot),
	449:  uint16(11),
	450:  uint16(2),
	451:  uint16(anon_sym_LT_PERCENT),
	452:  uint16(anon_sym_LT_PERCENT_PERCENT),
	453:  uint16(13),
	454:  uint16(2),
	455:  uint16(anon_sym_LT_PERCENT_EQ),
	456:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	457:  uint16(102),
	458:  uint16(2),
	459:  uint16(sym_text),
	460:  uint16(sym_entity),
	461:  uint16(28),
	462:  uint16(3),
	463:  uint16(sym__html_comment),
	464:  uint16(sym__bang_comment),
	465:  uint16(sym__hash_comment),
	466:  uint16(9),
	467:  uint16(8),
	468:  uint16(sym__node),
	469:  uint16(sym_doctype),
	470:  uint16(sym_tag),
	471:  uint16(sym_component),
	472:  uint16(sym_expression),
	473:  uint16(sym_directive),
	474:  uint16(sym_comment),
	475:  uint16(aux_sym_fragment_repeat1),
	476:  uint16(17),
	477:  uint16(5),
	478:  uint16(1),
	479:  uint16(anon_sym_LT_BANG),
	480:  uint16(7),
	481:  uint16(1),
	482:  uint16(anon_sym_LT),
	483:  uint16(9),
	484:  uint16(1),
	485:  uint16(anon_sym_LBRACE),
	486:  uint16(15),
	487:  uint16(1),
	488:  uint16(anon_sym_LT_BANG_DASH_DASH),
	489:  uint16(17),
	490:  uint16(1),
	491:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	492:  uint16(19),
	493:  uint16(1),
	494:  uint16(anon_sym_LT_PERCENT_POUND),
	495:  uint16(100),
	496:  uint16(1),
	497:  uint16(anon_sym_LT_SLASH_COLON),
	498:  uint16(2),
	499:  uint16(1),
	500:  uint16(sym_start_component),
	501:  uint16(6),
	502:  uint16(1),
	503:  uint16(sym_start_tag),
	504:  uint16(17),
	505:  uint16(1),
	506:  uint16(sym_self_closing_tag),
	507:  uint16(18),
	508:  uint16(1),
	509:  uint16(sym_self_closing_component),
	510:  uint16(45),
	511:  uint16(1),
	512:  uint16(sym_end_slot),
	513:  uint16(11),
	514:  uint16(2),
	515:  uint16(anon_sym_LT_PERCENT),
	516:  uint16(anon_sym_LT_PERCENT_PERCENT),
	517:  uint16(13),
	518:  uint16(2),
	519:  uint16(anon_sym_LT_PERCENT_EQ),
	520:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	521:  uint16(98),
	522:  uint16(2),
	523:  uint16(sym_text),
	524:  uint16(sym_entity),
	525:  uint16(28),
	526:  uint16(3),
	527:  uint16(sym__html_comment),
	528:  uint16(sym__bang_comment),
	529:  uint16(sym__hash_comment),
	530:  uint16(5),
	531:  uint16(8),
	532:  uint16(sym__node),
	533:  uint16(sym_doctype),
	534:  uint16(sym_tag),
	535:  uint16(sym_component),
	536:  uint16(sym_expression),
	537:  uint16(sym_directive),
	538:  uint16(sym_comment),
	539:  uint16(aux_sym_fragment_repeat1),
	540:  uint16(16),
	541:  uint16(5),
	542:  uint16(1),
	543:  uint16(anon_sym_LT_BANG),
	544:  uint16(7),
	545:  uint16(1),
	546:  uint16(anon_sym_LT),
	547:  uint16(9),
	548:  uint16(1),
	549:  uint16(anon_sym_LBRACE),
	550:  uint16(15),
	551:  uint16(1),
	552:  uint16(anon_sym_LT_BANG_DASH_DASH),
	553:  uint16(17),
	554:  uint16(1),
	555:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	556:  uint16(19),
	557:  uint16(1),
	558:  uint16(anon_sym_LT_PERCENT_POUND),
	559:  uint16(104),
	560:  uint16(1),
	562:  uint16(2),
	563:  uint16(1),
	564:  uint16(sym_start_component),
	565:  uint16(6),
	566:  uint16(1),
	567:  uint16(sym_start_tag),
	568:  uint16(17),
	569:  uint16(1),
	570:  uint16(sym_self_closing_tag),
	571:  uint16(18),
	572:  uint16(1),
	573:  uint16(sym_self_closing_component),
	574:  uint16(11),
	575:  uint16(2),
	576:  uint16(anon_sym_LT_PERCENT),
	577:  uint16(anon_sym_LT_PERCENT_PERCENT),
	578:  uint16(13),
	579:  uint16(2),
	580:  uint16(anon_sym_LT_PERCENT_EQ),
	581:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	582:  uint16(98),
	583:  uint16(2),
	584:  uint16(sym_text),
	585:  uint16(sym_entity),
	586:  uint16(28),
	587:  uint16(3),
	588:  uint16(sym__html_comment),
	589:  uint16(sym__bang_comment),
	590:  uint16(sym__hash_comment),
	591:  uint16(5),
	592:  uint16(8),
	593:  uint16(sym__node),
	594:  uint16(sym_doctype),
	595:  uint16(sym_tag),
	596:  uint16(sym_component),
	597:  uint16(sym_expression),
	598:  uint16(sym_directive),
	599:  uint16(sym_comment),
	600:  uint16(aux_sym_fragment_repeat1),
	601:  uint16(2),
	602:  uint16(108),
	603:  uint16(5),
	604:  uint16(anon_sym_LT_BANG),
	605:  uint16(anon_sym_LT),
	606:  uint16(anon_sym_LT_SLASH),
	607:  uint16(anon_sym_LT_PERCENT),
	608:  uint16(anon_sym_LT_PERCENT_PERCENT),
	609:  uint16(106),
	610:  uint16(11),
	612:  uint16(anon_sym_LT_COLON),
	613:  uint16(anon_sym_LT_SLASH_COLON),
	614:  uint16(anon_sym_LBRACE),
	615:  uint16(anon_sym_LT_PERCENT_EQ),
	616:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	617:  uint16(anon_sym_LT_BANG_DASH_DASH),
	618:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	619:  uint16(anon_sym_LT_PERCENT_POUND),
	620:  uint16(sym_text),
	621:  uint16(sym_entity),
	622:  uint16(2),
	623:  uint16(112),
	624:  uint16(5),
	625:  uint16(anon_sym_LT_BANG),
	626:  uint16(anon_sym_LT),
	627:  uint16(anon_sym_LT_SLASH),
	628:  uint16(anon_sym_LT_PERCENT),
	629:  uint16(anon_sym_LT_PERCENT_PERCENT),
	630:  uint16(110),
	631:  uint16(11),
	633:  uint16(anon_sym_LT_COLON),
	634:  uint16(anon_sym_LT_SLASH_COLON),
	635:  uint16(anon_sym_LBRACE),
	636:  uint16(anon_sym_LT_PERCENT_EQ),
	637:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	638:  uint16(anon_sym_LT_BANG_DASH_DASH),
	639:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	640:  uint16(anon_sym_LT_PERCENT_POUND),
	641:  uint16(sym_text),
	642:  uint16(sym_entity),
	643:  uint16(2),
	644:  uint16(116),
	645:  uint16(5),
	646:  uint16(anon_sym_LT_BANG),
	647:  uint16(anon_sym_LT),
	648:  uint16(anon_sym_LT_SLASH),
	649:  uint16(anon_sym_LT_PERCENT),
	650:  uint16(anon_sym_LT_PERCENT_PERCENT),
	651:  uint16(114),
	652:  uint16(11),
	654:  uint16(anon_sym_LT_COLON),
	655:  uint16(anon_sym_LT_SLASH_COLON),
	656:  uint16(anon_sym_LBRACE),
	657:  uint16(anon_sym_LT_PERCENT_EQ),
	658:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	659:  uint16(anon_sym_LT_BANG_DASH_DASH),
	660:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	661:  uint16(anon_sym_LT_PERCENT_POUND),
	662:  uint16(sym_text),
	663:  uint16(sym_entity),
	664:  uint16(2),
	665:  uint16(120),
	666:  uint16(5),
	667:  uint16(anon_sym_LT_BANG),
	668:  uint16(anon_sym_LT),
	669:  uint16(anon_sym_LT_SLASH),
	670:  uint16(anon_sym_LT_PERCENT),
	671:  uint16(anon_sym_LT_PERCENT_PERCENT),
	672:  uint16(118),
	673:  uint16(11),
	675:  uint16(anon_sym_LT_COLON),
	676:  uint16(anon_sym_LT_SLASH_COLON),
	677:  uint16(anon_sym_LBRACE),
	678:  uint16(anon_sym_LT_PERCENT_EQ),
	679:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	680:  uint16(anon_sym_LT_BANG_DASH_DASH),
	681:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	682:  uint16(anon_sym_LT_PERCENT_POUND),
	683:  uint16(sym_text),
	684:  uint16(sym_entity),
	685:  uint16(2),
	686:  uint16(124),
	687:  uint16(5),
	688:  uint16(anon_sym_LT_BANG),
	689:  uint16(anon_sym_LT),
	690:  uint16(anon_sym_LT_SLASH),
	691:  uint16(anon_sym_LT_PERCENT),
	692:  uint16(anon_sym_LT_PERCENT_PERCENT),
	693:  uint16(122),
	694:  uint16(11),
	696:  uint16(anon_sym_LT_COLON),
	697:  uint16(anon_sym_LT_SLASH_COLON),
	698:  uint16(anon_sym_LBRACE),
	699:  uint16(anon_sym_LT_PERCENT_EQ),
	700:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	701:  uint16(anon_sym_LT_BANG_DASH_DASH),
	702:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	703:  uint16(anon_sym_LT_PERCENT_POUND),
	704:  uint16(sym_text),
	705:  uint16(sym_entity),
	706:  uint16(2),
	707:  uint16(128),
	708:  uint16(5),
	709:  uint16(anon_sym_LT_BANG),
	710:  uint16(anon_sym_LT),
	711:  uint16(anon_sym_LT_SLASH),
	712:  uint16(anon_sym_LT_PERCENT),
	713:  uint16(anon_sym_LT_PERCENT_PERCENT),
	714:  uint16(126),
	715:  uint16(11),
	717:  uint16(anon_sym_LT_COLON),
	718:  uint16(anon_sym_LT_SLASH_COLON),
	719:  uint16(anon_sym_LBRACE),
	720:  uint16(anon_sym_LT_PERCENT_EQ),
	721:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	722:  uint16(anon_sym_LT_BANG_DASH_DASH),
	723:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	724:  uint16(anon_sym_LT_PERCENT_POUND),
	725:  uint16(sym_text),
	726:  uint16(sym_entity),
	727:  uint16(2),
	728:  uint16(132),
	729:  uint16(5),
	730:  uint16(anon_sym_LT_BANG),
	731:  uint16(anon_sym_LT),
	732:  uint16(anon_sym_LT_SLASH),
	733:  uint16(anon_sym_LT_PERCENT),
	734:  uint16(anon_sym_LT_PERCENT_PERCENT),
	735:  uint16(130),
	736:  uint16(11),
	738:  uint16(anon_sym_LT_COLON),
	739:  uint16(anon_sym_LT_SLASH_COLON),
	740:  uint16(anon_sym_LBRACE),
	741:  uint16(anon_sym_LT_PERCENT_EQ),
	742:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	743:  uint16(anon_sym_LT_BANG_DASH_DASH),
	744:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	745:  uint16(anon_sym_LT_PERCENT_POUND),
	746:  uint16(sym_text),
	747:  uint16(sym_entity),
	748:  uint16(2),
	749:  uint16(136),
	750:  uint16(5),
	751:  uint16(anon_sym_LT_BANG),
	752:  uint16(anon_sym_LT),
	753:  uint16(anon_sym_LT_SLASH),
	754:  uint16(anon_sym_LT_PERCENT),
	755:  uint16(anon_sym_LT_PERCENT_PERCENT),
	756:  uint16(134),
	757:  uint16(11),
	759:  uint16(anon_sym_LT_COLON),
	760:  uint16(anon_sym_LT_SLASH_COLON),
	761:  uint16(anon_sym_LBRACE),
	762:  uint16(anon_sym_LT_PERCENT_EQ),
	763:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	764:  uint16(anon_sym_LT_BANG_DASH_DASH),
	765:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	766:  uint16(anon_sym_LT_PERCENT_POUND),
	767:  uint16(sym_text),
	768:  uint16(sym_entity),
	769:  uint16(2),
	770:  uint16(140),
	771:  uint16(5),
	772:  uint16(anon_sym_LT_BANG),
	773:  uint16(anon_sym_LT),
	774:  uint16(anon_sym_LT_SLASH),
	775:  uint16(anon_sym_LT_PERCENT),
	776:  uint16(anon_sym_LT_PERCENT_PERCENT),
	777:  uint16(138),
	778:  uint16(11),
	780:  uint16(anon_sym_LT_COLON),
	781:  uint16(anon_sym_LT_SLASH_COLON),
	782:  uint16(anon_sym_LBRACE),
	783:  uint16(anon_sym_LT_PERCENT_EQ),
	784:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	785:  uint16(anon_sym_LT_BANG_DASH_DASH),
	786:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	787:  uint16(anon_sym_LT_PERCENT_POUND),
	788:  uint16(sym_text),
	789:  uint16(sym_entity),
	790:  uint16(2),
	791:  uint16(144),
	792:  uint16(5),
	793:  uint16(anon_sym_LT_BANG),
	794:  uint16(anon_sym_LT),
	795:  uint16(anon_sym_LT_SLASH),
	796:  uint16(anon_sym_LT_PERCENT),
	797:  uint16(anon_sym_LT_PERCENT_PERCENT),
	798:  uint16(142),
	799:  uint16(11),
	801:  uint16(anon_sym_LT_COLON),
	802:  uint16(anon_sym_LT_SLASH_COLON),
	803:  uint16(anon_sym_LBRACE),
	804:  uint16(anon_sym_LT_PERCENT_EQ),
	805:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	806:  uint16(anon_sym_LT_BANG_DASH_DASH),
	807:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	808:  uint16(anon_sym_LT_PERCENT_POUND),
	809:  uint16(sym_text),
	810:  uint16(sym_entity),
	811:  uint16(2),
	812:  uint16(148),
	813:  uint16(5),
	814:  uint16(anon_sym_LT_BANG),
	815:  uint16(anon_sym_LT),
	816:  uint16(anon_sym_LT_SLASH),
	817:  uint16(anon_sym_LT_PERCENT),
	818:  uint16(anon_sym_LT_PERCENT_PERCENT),
	819:  uint16(146),
	820:  uint16(11),
	822:  uint16(anon_sym_LT_COLON),
	823:  uint16(anon_sym_LT_SLASH_COLON),
	824:  uint16(anon_sym_LBRACE),
	825:  uint16(anon_sym_LT_PERCENT_EQ),
	826:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	827:  uint16(anon_sym_LT_BANG_DASH_DASH),
	828:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	829:  uint16(anon_sym_LT_PERCENT_POUND),
	830:  uint16(sym_text),
	831:  uint16(sym_entity),
	832:  uint16(2),
	833:  uint16(152),
	834:  uint16(5),
	835:  uint16(anon_sym_LT_BANG),
	836:  uint16(anon_sym_LT),
	837:  uint16(anon_sym_LT_SLASH),
	838:  uint16(anon_sym_LT_PERCENT),
	839:  uint16(anon_sym_LT_PERCENT_PERCENT),
	840:  uint16(150),
	841:  uint16(11),
	843:  uint16(anon_sym_LT_COLON),
	844:  uint16(anon_sym_LT_SLASH_COLON),
	845:  uint16(anon_sym_LBRACE),
	846:  uint16(anon_sym_LT_PERCENT_EQ),
	847:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	848:  uint16(anon_sym_LT_BANG_DASH_DASH),
	849:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	850:  uint16(anon_sym_LT_PERCENT_POUND),
	851:  uint16(sym_text),
	852:  uint16(sym_entity),
	853:  uint16(2),
	854:  uint16(156),
	855:  uint16(5),
	856:  uint16(anon_sym_LT_BANG),
	857:  uint16(anon_sym_LT),
	858:  uint16(anon_sym_LT_SLASH),
	859:  uint16(anon_sym_LT_PERCENT),
	860:  uint16(anon_sym_LT_PERCENT_PERCENT),
	861:  uint16(154),
	862:  uint16(11),
	864:  uint16(anon_sym_LT_COLON),
	865:  uint16(anon_sym_LT_SLASH_COLON),
	866:  uint16(anon_sym_LBRACE),
	867:  uint16(anon_sym_LT_PERCENT_EQ),
	868:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	869:  uint16(anon_sym_LT_BANG_DASH_DASH),
	870:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	871:  uint16(anon_sym_LT_PERCENT_POUND),
	872:  uint16(sym_text),
	873:  uint16(sym_entity),
	874:  uint16(2),
	875:  uint16(160),
	876:  uint16(5),
	877:  uint16(anon_sym_LT_BANG),
	878:  uint16(anon_sym_LT),
	879:  uint16(anon_sym_LT_SLASH),
	880:  uint16(anon_sym_LT_PERCENT),
	881:  uint16(anon_sym_LT_PERCENT_PERCENT),
	882:  uint16(158),
	883:  uint16(11),
	885:  uint16(anon_sym_LT_COLON),
	886:  uint16(anon_sym_LT_SLASH_COLON),
	887:  uint16(anon_sym_LBRACE),
	888:  uint16(anon_sym_LT_PERCENT_EQ),
	889:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	890:  uint16(anon_sym_LT_BANG_DASH_DASH),
	891:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	892:  uint16(anon_sym_LT_PERCENT_POUND),
	893:  uint16(sym_text),
	894:  uint16(sym_entity),
	895:  uint16(2),
	896:  uint16(164),
	897:  uint16(5),
	898:  uint16(anon_sym_LT_BANG),
	899:  uint16(anon_sym_LT),
	900:  uint16(anon_sym_LT_SLASH),
	901:  uint16(anon_sym_LT_PERCENT),
	902:  uint16(anon_sym_LT_PERCENT_PERCENT),
	903:  uint16(162),
	904:  uint16(11),
	906:  uint16(anon_sym_LT_COLON),
	907:  uint16(anon_sym_LT_SLASH_COLON),
	908:  uint16(anon_sym_LBRACE),
	909:  uint16(anon_sym_LT_PERCENT_EQ),
	910:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	911:  uint16(anon_sym_LT_BANG_DASH_DASH),
	912:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	913:  uint16(anon_sym_LT_PERCENT_POUND),
	914:  uint16(sym_text),
	915:  uint16(sym_entity),
	916:  uint16(2),
	917:  uint16(168),
	918:  uint16(5),
	919:  uint16(anon_sym_LT_BANG),
	920:  uint16(anon_sym_LT),
	921:  uint16(anon_sym_LT_SLASH),
	922:  uint16(anon_sym_LT_PERCENT),
	923:  uint16(anon_sym_LT_PERCENT_PERCENT),
	924:  uint16(166),
	925:  uint16(11),
	927:  uint16(anon_sym_LT_COLON),
	928:  uint16(anon_sym_LT_SLASH_COLON),
	929:  uint16(anon_sym_LBRACE),
	930:  uint16(anon_sym_LT_PERCENT_EQ),
	931:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	932:  uint16(anon_sym_LT_BANG_DASH_DASH),
	933:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	934:  uint16(anon_sym_LT_PERCENT_POUND),
	935:  uint16(sym_text),
	936:  uint16(sym_entity),
	937:  uint16(2),
	938:  uint16(172),
	939:  uint16(5),
	940:  uint16(anon_sym_LT_BANG),
	941:  uint16(anon_sym_LT),
	942:  uint16(anon_sym_LT_SLASH),
	943:  uint16(anon_sym_LT_PERCENT),
	944:  uint16(anon_sym_LT_PERCENT_PERCENT),
	945:  uint16(170),
	946:  uint16(11),
	948:  uint16(anon_sym_LT_COLON),
	949:  uint16(anon_sym_LT_SLASH_COLON),
	950:  uint16(anon_sym_LBRACE),
	951:  uint16(anon_sym_LT_PERCENT_EQ),
	952:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	953:  uint16(anon_sym_LT_BANG_DASH_DASH),
	954:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	955:  uint16(anon_sym_LT_PERCENT_POUND),
	956:  uint16(sym_text),
	957:  uint16(sym_entity),
	958:  uint16(2),
	959:  uint16(176),
	960:  uint16(5),
	961:  uint16(anon_sym_LT_BANG),
	962:  uint16(anon_sym_LT),
	963:  uint16(anon_sym_LT_SLASH),
	964:  uint16(anon_sym_LT_PERCENT),
	965:  uint16(anon_sym_LT_PERCENT_PERCENT),
	966:  uint16(174),
	967:  uint16(11),
	969:  uint16(anon_sym_LT_COLON),
	970:  uint16(anon_sym_LT_SLASH_COLON),
	971:  uint16(anon_sym_LBRACE),
	972:  uint16(anon_sym_LT_PERCENT_EQ),
	973:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	974:  uint16(anon_sym_LT_BANG_DASH_DASH),
	975:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	976:  uint16(anon_sym_LT_PERCENT_POUND),
	977:  uint16(sym_text),
	978:  uint16(sym_entity),
	979:  uint16(2),
	980:  uint16(180),
	981:  uint16(5),
	982:  uint16(anon_sym_LT_BANG),
	983:  uint16(anon_sym_LT),
	984:  uint16(anon_sym_LT_SLASH),
	985:  uint16(anon_sym_LT_PERCENT),
	986:  uint16(anon_sym_LT_PERCENT_PERCENT),
	987:  uint16(178),
	988:  uint16(11),
	990:  uint16(anon_sym_LT_COLON),
	991:  uint16(anon_sym_LT_SLASH_COLON),
	992:  uint16(anon_sym_LBRACE),
	993:  uint16(anon_sym_LT_PERCENT_EQ),
	994:  uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	995:  uint16(anon_sym_LT_BANG_DASH_DASH),
	996:  uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	997:  uint16(anon_sym_LT_PERCENT_POUND),
	998:  uint16(sym_text),
	999:  uint16(sym_entity),
	1000: uint16(2),
	1001: uint16(184),
	1002: uint16(5),
	1003: uint16(anon_sym_LT_BANG),
	1004: uint16(anon_sym_LT),
	1005: uint16(anon_sym_LT_SLASH),
	1006: uint16(anon_sym_LT_PERCENT),
	1007: uint16(anon_sym_LT_PERCENT_PERCENT),
	1008: uint16(182),
	1009: uint16(11),
	1011: uint16(anon_sym_LT_COLON),
	1012: uint16(anon_sym_LT_SLASH_COLON),
	1013: uint16(anon_sym_LBRACE),
	1014: uint16(anon_sym_LT_PERCENT_EQ),
	1015: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1016: uint16(anon_sym_LT_BANG_DASH_DASH),
	1017: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1018: uint16(anon_sym_LT_PERCENT_POUND),
	1019: uint16(sym_text),
	1020: uint16(sym_entity),
	1021: uint16(2),
	1022: uint16(188),
	1023: uint16(5),
	1024: uint16(anon_sym_LT_BANG),
	1025: uint16(anon_sym_LT),
	1026: uint16(anon_sym_LT_SLASH),
	1027: uint16(anon_sym_LT_PERCENT),
	1028: uint16(anon_sym_LT_PERCENT_PERCENT),
	1029: uint16(186),
	1030: uint16(11),
	1032: uint16(anon_sym_LT_COLON),
	1033: uint16(anon_sym_LT_SLASH_COLON),
	1034: uint16(anon_sym_LBRACE),
	1035: uint16(anon_sym_LT_PERCENT_EQ),
	1036: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1037: uint16(anon_sym_LT_BANG_DASH_DASH),
	1038: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1039: uint16(anon_sym_LT_PERCENT_POUND),
	1040: uint16(sym_text),
	1041: uint16(sym_entity),
	1042: uint16(2),
	1043: uint16(192),
	1044: uint16(5),
	1045: uint16(anon_sym_LT_BANG),
	1046: uint16(anon_sym_LT),
	1047: uint16(anon_sym_LT_SLASH),
	1048: uint16(anon_sym_LT_PERCENT),
	1049: uint16(anon_sym_LT_PERCENT_PERCENT),
	1050: uint16(190),
	1051: uint16(11),
	1053: uint16(anon_sym_LT_COLON),
	1054: uint16(anon_sym_LT_SLASH_COLON),
	1055: uint16(anon_sym_LBRACE),
	1056: uint16(anon_sym_LT_PERCENT_EQ),
	1057: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1058: uint16(anon_sym_LT_BANG_DASH_DASH),
	1059: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1060: uint16(anon_sym_LT_PERCENT_POUND),
	1061: uint16(sym_text),
	1062: uint16(sym_entity),
	1063: uint16(2),
	1064: uint16(196),
	1065: uint16(5),
	1066: uint16(anon_sym_LT_BANG),
	1067: uint16(anon_sym_LT),
	1068: uint16(anon_sym_LT_SLASH),
	1069: uint16(anon_sym_LT_PERCENT),
	1070: uint16(anon_sym_LT_PERCENT_PERCENT),
	1071: uint16(194),
	1072: uint16(11),
	1074: uint16(anon_sym_LT_COLON),
	1075: uint16(anon_sym_LT_SLASH_COLON),
	1076: uint16(anon_sym_LBRACE),
	1077: uint16(anon_sym_LT_PERCENT_EQ),
	1078: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1079: uint16(anon_sym_LT_BANG_DASH_DASH),
	1080: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1081: uint16(anon_sym_LT_PERCENT_POUND),
	1082: uint16(sym_text),
	1083: uint16(sym_entity),
	1084: uint16(2),
	1085: uint16(198),
	1086: uint16(4),
	1087: uint16(anon_sym_LT_BANG),
	1088: uint16(anon_sym_LT),
	1089: uint16(anon_sym_LT_PERCENT),
	1090: uint16(anon_sym_LT_PERCENT_PERCENT),
	1091: uint16(200),
	1092: uint16(10),
	1093: uint16(anon_sym_LT_SLASH),
	1094: uint16(anon_sym_LT_COLON),
	1095: uint16(anon_sym_LBRACE),
	1096: uint16(anon_sym_LT_PERCENT_EQ),
	1097: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1098: uint16(anon_sym_LT_BANG_DASH_DASH),
	1099: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1100: uint16(anon_sym_LT_PERCENT_POUND),
	1101: uint16(sym_text),
	1102: uint16(sym_entity),
	1103: uint16(2),
	1104: uint16(202),
	1105: uint16(4),
	1106: uint16(anon_sym_LT_BANG),
	1107: uint16(anon_sym_LT),
	1108: uint16(anon_sym_LT_PERCENT),
	1109: uint16(anon_sym_LT_PERCENT_PERCENT),
	1110: uint16(204),
	1111: uint16(10),
	1112: uint16(anon_sym_LT_SLASH),
	1113: uint16(anon_sym_LT_COLON),
	1114: uint16(anon_sym_LBRACE),
	1115: uint16(anon_sym_LT_PERCENT_EQ),
	1116: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1117: uint16(anon_sym_LT_BANG_DASH_DASH),
	1118: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1119: uint16(anon_sym_LT_PERCENT_POUND),
	1120: uint16(sym_text),
	1121: uint16(sym_entity),
	1122: uint16(2),
	1123: uint16(206),
	1124: uint16(4),
	1125: uint16(anon_sym_LT_BANG),
	1126: uint16(anon_sym_LT),
	1127: uint16(anon_sym_LT_PERCENT),
	1128: uint16(anon_sym_LT_PERCENT_PERCENT),
	1129: uint16(208),
	1130: uint16(10),
	1131: uint16(anon_sym_LT_SLASH),
	1132: uint16(anon_sym_LT_COLON),
	1133: uint16(anon_sym_LBRACE),
	1134: uint16(anon_sym_LT_PERCENT_EQ),
	1135: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1136: uint16(anon_sym_LT_BANG_DASH_DASH),
	1137: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1138: uint16(anon_sym_LT_PERCENT_POUND),
	1139: uint16(sym_text),
	1140: uint16(sym_entity),
	1141: uint16(2),
	1142: uint16(210),
	1143: uint16(4),
	1144: uint16(anon_sym_LT_BANG),
	1145: uint16(anon_sym_LT),
	1146: uint16(anon_sym_LT_PERCENT),
	1147: uint16(anon_sym_LT_PERCENT_PERCENT),
	1148: uint16(212),
	1149: uint16(10),
	1150: uint16(anon_sym_LT_SLASH),
	1151: uint16(anon_sym_LT_COLON),
	1152: uint16(anon_sym_LBRACE),
	1153: uint16(anon_sym_LT_PERCENT_EQ),
	1154: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1155: uint16(anon_sym_LT_BANG_DASH_DASH),
	1156: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1157: uint16(anon_sym_LT_PERCENT_POUND),
	1158: uint16(sym_text),
	1159: uint16(sym_entity),
	1160: uint16(6),
	1161: uint16(216),
	1162: uint16(1),
	1163: uint16(anon_sym_LBRACE),
	1164: uint16(222),
	1165: uint16(1),
	1166: uint16(sym_attribute_name),
	1167: uint16(105),
	1168: uint16(1),
	1169: uint16(sym_special_attribute_name),
	1170: uint16(214),
	1171: uint16(2),
	1172: uint16(anon_sym_GT),
	1173: uint16(anon_sym_SLASH_GT),
	1174: uint16(38),
	1175: uint16(4),
	1176: uint16(sym_expression),
	1177: uint16(sym_special_attribute),
	1178: uint16(sym_attribute),
	1179: uint16(aux_sym_start_tag_repeat1),
	1180: uint16(219),
	1181: uint16(5),
	1182: uint16(anon_sym_COLONlet),
	1183: uint16(anon_sym_COLONfor),
	1184: uint16(anon_sym_COLONkey),
	1185: uint16(anon_sym_COLONstream),
	1186: uint16(anon_sym_COLONif),
	1187: uint16(7),
	1188: uint16(225),
	1189: uint16(1),
	1190: uint16(anon_sym_GT),
	1191: uint16(227),
	1192: uint16(1),
	1193: uint16(anon_sym_SLASH_GT),
	1194: uint16(229),
	1195: uint16(1),
	1196: uint16(anon_sym_LBRACE),
	1197: uint16(233),
	1198: uint16(1),
	1199: uint16(sym_attribute_name),
	1200: uint16(105),
	1201: uint16(1),
	1202: uint16(sym_special_attribute_name),
	1203: uint16(38),
	1204: uint16(4),
	1205: uint16(sym_expression),
	1206: uint16(sym_special_attribute),
	1207: uint16(sym_attribute),
	1208: uint16(aux_sym_start_tag_repeat1),
	1209: uint16(231),
	1210: uint16(5),
	1211: uint16(anon_sym_COLONlet),
	1212: uint16(anon_sym_COLONfor),
	1213: uint16(anon_sym_COLONkey),
	1214: uint16(anon_sym_COLONstream),
	1215: uint16(anon_sym_COLONif),
	1216: uint16(7),
	1217: uint16(229),
	1218: uint16(1),
	1219: uint16(anon_sym_LBRACE),
	1220: uint16(233),
	1221: uint16(1),
	1222: uint16(sym_attribute_name),
	1223: uint16(235),
	1224: uint16(1),
	1225: uint16(anon_sym_GT),
	1226: uint16(237),
	1227: uint16(1),
	1228: uint16(anon_sym_SLASH_GT),
	1229: uint16(105),
	1230: uint16(1),
	1231: uint16(sym_special_attribute_name),
	1232: uint16(39),
	1233: uint16(4),
	1234: uint16(sym_expression),
	1235: uint16(sym_special_attribute),
	1236: uint16(sym_attribute),
	1237: uint16(aux_sym_start_tag_repeat1),
	1238: uint16(231),
	1239: uint16(5),
	1240: uint16(anon_sym_COLONlet),
	1241: uint16(anon_sym_COLONfor),
	1242: uint16(anon_sym_COLONkey),
	1243: uint16(anon_sym_COLONstream),
	1244: uint16(anon_sym_COLONif),
	1245: uint16(2),
	1246: uint16(239),
	1247: uint16(4),
	1248: uint16(anon_sym_LT_BANG),
	1249: uint16(anon_sym_LT),
	1250: uint16(anon_sym_LT_PERCENT),
	1251: uint16(anon_sym_LT_PERCENT_PERCENT),
	1252: uint16(241),
	1253: uint16(10),
	1254: uint16(anon_sym_LT_SLASH),
	1255: uint16(anon_sym_LT_COLON),
	1256: uint16(anon_sym_LBRACE),
	1257: uint16(anon_sym_LT_PERCENT_EQ),
	1258: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1259: uint16(anon_sym_LT_BANG_DASH_DASH),
	1260: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1261: uint16(anon_sym_LT_PERCENT_POUND),
	1262: uint16(sym_text),
	1263: uint16(sym_entity),
	1264: uint16(7),
	1265: uint16(229),
	1266: uint16(1),
	1267: uint16(anon_sym_LBRACE),
	1268: uint16(233),
	1269: uint16(1),
	1270: uint16(sym_attribute_name),
	1271: uint16(243),
	1272: uint16(1),
	1273: uint16(anon_sym_GT),
	1274: uint16(245),
	1275: uint16(1),
	1276: uint16(anon_sym_SLASH_GT),
	1277: uint16(105),
	1278: uint16(1),
	1279: uint16(sym_special_attribute_name),
	1280: uint16(44),
	1281: uint16(4),
	1282: uint16(sym_expression),
	1283: uint16(sym_special_attribute),
	1284: uint16(sym_attribute),
	1285: uint16(aux_sym_start_tag_repeat1),
	1286: uint16(231),
	1287: uint16(5),
	1288: uint16(anon_sym_COLONlet),
	1289: uint16(anon_sym_COLONfor),
	1290: uint16(anon_sym_COLONkey),
	1291: uint16(anon_sym_COLONstream),
	1292: uint16(anon_sym_COLONif),
	1293: uint16(7),
	1294: uint16(229),
	1295: uint16(1),
	1296: uint16(anon_sym_LBRACE),
	1297: uint16(233),
	1298: uint16(1),
	1299: uint16(sym_attribute_name),
	1300: uint16(247),
	1301: uint16(1),
	1302: uint16(anon_sym_GT),
	1303: uint16(249),
	1304: uint16(1),
	1305: uint16(anon_sym_SLASH_GT),
	1306: uint16(105),
	1307: uint16(1),
	1308: uint16(sym_special_attribute_name),
	1309: uint16(46),
	1310: uint16(4),
	1311: uint16(sym_expression),
	1312: uint16(sym_special_attribute),
	1313: uint16(sym_attribute),
	1314: uint16(aux_sym_start_tag_repeat1),
	1315: uint16(231),
	1316: uint16(5),
	1317: uint16(anon_sym_COLONlet),
	1318: uint16(anon_sym_COLONfor),
	1319: uint16(anon_sym_COLONkey),
	1320: uint16(anon_sym_COLONstream),
	1321: uint16(anon_sym_COLONif),
	1322: uint16(7),
	1323: uint16(229),
	1324: uint16(1),
	1325: uint16(anon_sym_LBRACE),
	1326: uint16(233),
	1327: uint16(1),
	1328: uint16(sym_attribute_name),
	1329: uint16(251),
	1330: uint16(1),
	1331: uint16(anon_sym_GT),
	1332: uint16(253),
	1333: uint16(1),
	1334: uint16(anon_sym_SLASH_GT),
	1335: uint16(105),
	1336: uint16(1),
	1337: uint16(sym_special_attribute_name),
	1338: uint16(38),
	1339: uint16(4),
	1340: uint16(sym_expression),
	1341: uint16(sym_special_attribute),
	1342: uint16(sym_attribute),
	1343: uint16(aux_sym_start_tag_repeat1),
	1344: uint16(231),
	1345: uint16(5),
	1346: uint16(anon_sym_COLONlet),
	1347: uint16(anon_sym_COLONfor),
	1348: uint16(anon_sym_COLONkey),
	1349: uint16(anon_sym_COLONstream),
	1350: uint16(anon_sym_COLONif),
	1351: uint16(2),
	1352: uint16(255),
	1353: uint16(4),
	1354: uint16(anon_sym_LT_BANG),
	1355: uint16(anon_sym_LT),
	1356: uint16(anon_sym_LT_PERCENT),
	1357: uint16(anon_sym_LT_PERCENT_PERCENT),
	1358: uint16(257),
	1359: uint16(10),
	1360: uint16(anon_sym_LT_SLASH),
	1361: uint16(anon_sym_LT_COLON),
	1362: uint16(anon_sym_LBRACE),
	1363: uint16(anon_sym_LT_PERCENT_EQ),
	1364: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1365: uint16(anon_sym_LT_BANG_DASH_DASH),
	1366: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1367: uint16(anon_sym_LT_PERCENT_POUND),
	1368: uint16(sym_text),
	1369: uint16(sym_entity),
	1370: uint16(7),
	1371: uint16(229),
	1372: uint16(1),
	1373: uint16(anon_sym_LBRACE),
	1374: uint16(233),
	1375: uint16(1),
	1376: uint16(sym_attribute_name),
	1377: uint16(259),
	1378: uint16(1),
	1379: uint16(anon_sym_GT),
	1380: uint16(261),
	1381: uint16(1),
	1382: uint16(anon_sym_SLASH_GT),
	1383: uint16(105),
	1384: uint16(1),
	1385: uint16(sym_special_attribute_name),
	1386: uint16(38),
	1387: uint16(4),
	1388: uint16(sym_expression),
	1389: uint16(sym_special_attribute),
	1390: uint16(sym_attribute),
	1391: uint16(aux_sym_start_tag_repeat1),
	1392: uint16(231),
	1393: uint16(5),
	1394: uint16(anon_sym_COLONlet),
	1395: uint16(anon_sym_COLONfor),
	1396: uint16(anon_sym_COLONkey),
	1397: uint16(anon_sym_COLONstream),
	1398: uint16(anon_sym_COLONif),
	1399: uint16(2),
	1400: uint16(263),
	1401: uint16(4),
	1402: uint16(anon_sym_LT_BANG),
	1403: uint16(anon_sym_LT),
	1404: uint16(anon_sym_LT_PERCENT),
	1405: uint16(anon_sym_LT_PERCENT_PERCENT),
	1406: uint16(265),
	1407: uint16(10),
	1408: uint16(anon_sym_LT_SLASH),
	1409: uint16(anon_sym_LT_COLON),
	1410: uint16(anon_sym_LBRACE),
	1411: uint16(anon_sym_LT_PERCENT_EQ),
	1412: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1413: uint16(anon_sym_LT_BANG_DASH_DASH),
	1414: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1415: uint16(anon_sym_LT_PERCENT_POUND),
	1416: uint16(sym_text),
	1417: uint16(sym_entity),
	1418: uint16(2),
	1419: uint16(267),
	1420: uint16(4),
	1421: uint16(anon_sym_LT_BANG),
	1422: uint16(anon_sym_LT),
	1423: uint16(anon_sym_LT_PERCENT),
	1424: uint16(anon_sym_LT_PERCENT_PERCENT),
	1425: uint16(269),
	1426: uint16(10),
	1427: uint16(anon_sym_LT_SLASH),
	1428: uint16(anon_sym_LT_COLON),
	1429: uint16(anon_sym_LBRACE),
	1430: uint16(anon_sym_LT_PERCENT_EQ),
	1431: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1432: uint16(anon_sym_LT_BANG_DASH_DASH),
	1433: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1434: uint16(anon_sym_LT_PERCENT_POUND),
	1435: uint16(sym_text),
	1436: uint16(sym_entity),
	1437: uint16(2),
	1438: uint16(271),
	1439: uint16(4),
	1440: uint16(anon_sym_LT_BANG),
	1441: uint16(anon_sym_LT),
	1442: uint16(anon_sym_LT_PERCENT),
	1443: uint16(anon_sym_LT_PERCENT_PERCENT),
	1444: uint16(273),
	1445: uint16(9),
	1446: uint16(anon_sym_LT_SLASH),
	1447: uint16(anon_sym_LBRACE),
	1448: uint16(anon_sym_LT_PERCENT_EQ),
	1449: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1450: uint16(anon_sym_LT_BANG_DASH_DASH),
	1451: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1452: uint16(anon_sym_LT_PERCENT_POUND),
	1453: uint16(sym_text),
	1454: uint16(sym_entity),
	1455: uint16(2),
	1456: uint16(275),
	1457: uint16(4),
	1458: uint16(anon_sym_LT_BANG),
	1459: uint16(anon_sym_LT),
	1460: uint16(anon_sym_LT_PERCENT),
	1461: uint16(anon_sym_LT_PERCENT_PERCENT),
	1462: uint16(277),
	1463: uint16(9),
	1464: uint16(anon_sym_LT_SLASH_COLON),
	1465: uint16(anon_sym_LBRACE),
	1466: uint16(anon_sym_LT_PERCENT_EQ),
	1467: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1468: uint16(anon_sym_LT_BANG_DASH_DASH),
	1469: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1470: uint16(anon_sym_LT_PERCENT_POUND),
	1471: uint16(sym_text),
	1472: uint16(sym_entity),
	1473: uint16(2),
	1474: uint16(279),
	1475: uint16(4),
	1476: uint16(anon_sym_LT_BANG),
	1477: uint16(anon_sym_LT),
	1478: uint16(anon_sym_LT_PERCENT),
	1479: uint16(anon_sym_LT_PERCENT_PERCENT),
	1480: uint16(281),
	1481: uint16(9),
	1482: uint16(anon_sym_LT_SLASH_COLON),
	1483: uint16(anon_sym_LBRACE),
	1484: uint16(anon_sym_LT_PERCENT_EQ),
	1485: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1486: uint16(anon_sym_LT_BANG_DASH_DASH),
	1487: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1488: uint16(anon_sym_LT_PERCENT_POUND),
	1489: uint16(sym_text),
	1490: uint16(sym_entity),
	1491: uint16(2),
	1492: uint16(283),
	1493: uint16(4),
	1494: uint16(anon_sym_LT_BANG),
	1495: uint16(anon_sym_LT),
	1496: uint16(anon_sym_LT_PERCENT),
	1497: uint16(anon_sym_LT_PERCENT_PERCENT),
	1498: uint16(285),
	1499: uint16(9),
	1500: uint16(anon_sym_LT_SLASH),
	1501: uint16(anon_sym_LBRACE),
	1502: uint16(anon_sym_LT_PERCENT_EQ),
	1503: uint16(anon_sym_LT_PERCENT_PERCENT_EQ),
	1504: uint16(anon_sym_LT_BANG_DASH_DASH),
	1505: uint16(anon_sym_LT_PERCENT_BANG_DASH_DASH),
	1506: uint16(anon_sym_LT_PERCENT_POUND),
	1507: uint16(sym_text),
	1508: uint16(sym_entity),
	1509: uint16(5),
	1510: uint16(289),
	1511: uint16(1),
	1512: uint16(aux_sym_ending_expression_value_token1),
	1513: uint16(291),
	1514: uint16(1),
	1515: uint16(sym__code),
	1516: uint16(62),
	1517: uint16(1),
	1518: uint16(aux_sym__hash_comment_repeat1),
	1519: uint16(103),
	1520: uint16(3),
	1521: uint16(sym_expression_value),
	1522: uint16(sym_partial_expression_value),
	1523: uint16(sym_ending_expression_value),
	1524: uint16(287),
	1525: uint16(6),
	1526: uint16(anon_sym_do),
	1527: uint16(anon_sym_catch),
	1528: uint16(anon_sym_rescue),
	1529: uint16(anon_sym_after),
	1530: uint16(anon_sym_else),
	1531: uint16(anon_sym_DASH_GT),
	1532: uint16(3),
	1533: uint16(295),
	1534: uint16(1),
	1535: uint16(anon_sym_DOT),
	1536: uint16(297),
	1537: uint16(1),
	1538: uint16(sym_attribute_name),
	1539: uint16(293),
	1540: uint16(8),
	1541: uint16(anon_sym_GT),
	1542: uint16(anon_sym_SLASH_GT),
	1543: uint16(anon_sym_LBRACE),
	1544: uint16(anon_sym_COLONlet),
	1545: uint16(anon_sym_COLONfor),
	1546: uint16(anon_sym_COLONkey),
	1547: uint16(anon_sym_COLONstream),
	1548: uint16(anon_sym_COLONif),
	1549: uint16(2),
	1550: uint16(301),
	1551: uint16(1),
	1552: uint16(anon_sym_EQ),
	1553: uint16(299),
	1554: uint16(9),
	1555: uint16(anon_sym_GT),
	1556: uint16(anon_sym_SLASH_GT),
	1557: uint16(anon_sym_LBRACE),
	1558: uint16(anon_sym_COLONlet),
	1559: uint16(anon_sym_COLONfor),
	1560: uint16(anon_sym_COLONkey),
	1561: uint16(anon_sym_COLONstream),
	1562: uint16(anon_sym_COLONif),
	1563: uint16(sym_attribute_name),
	1564: uint16(1),
	1565: uint16(303),
	1566: uint16(9),
	1567: uint16(anon_sym_GT),
	1568: uint16(anon_sym_SLASH_GT),
	1569: uint16(anon_sym_LBRACE),
	1570: uint16(anon_sym_COLONlet),
	1571: uint16(anon_sym_COLONfor),
	1572: uint16(anon_sym_COLONkey),
	1573: uint16(anon_sym_COLONstream),
	1574: uint16(anon_sym_COLONif),
	1575: uint16(sym_attribute_name),
	1576: uint16(3),
	1577: uint16(307),
	1578: uint16(1),
	1579: uint16(sym__code),
	1580: uint16(57),
	1581: uint16(1),
	1582: uint16(aux_sym__hash_comment_repeat1),
	1583: uint16(305),
	1584: uint16(7),
	1585: uint16(anon_sym_PERCENT_GT),
	1586: uint16(anon_sym_do),
	1587: uint16(anon_sym_catch),
	1588: uint16(anon_sym_rescue),
	1589: uint16(anon_sym_after),
	1590: uint16(anon_sym_else),
	1591: uint16(anon_sym_DASH_GT),
	1592: uint16(1),
	1593: uint16(310),
	1594: uint16(9),
	1595: uint16(anon_sym_GT),
	1596: uint16(anon_sym_SLASH_GT),
	1597: uint16(anon_sym_LBRACE),
	1598: uint16(anon_sym_COLONlet),
	1599: uint16(anon_sym_COLONfor),
	1600: uint16(anon_sym_COLONkey),
	1601: uint16(anon_sym_COLONstream),
	1602: uint16(anon_sym_COLONif),
	1603: uint16(sym_attribute_name),
	1604: uint16(1),
	1605: uint16(312),
	1606: uint16(9),
	1607: uint16(anon_sym_GT),
	1608: uint16(anon_sym_SLASH_GT),
	1609: uint16(anon_sym_LBRACE),
	1610: uint16(anon_sym_COLONlet),
	1611: uint16(anon_sym_COLONfor),
	1612: uint16(anon_sym_COLONkey),
	1613: uint16(anon_sym_COLONstream),
	1614: uint16(anon_sym_COLONif),
	1615: uint16(sym_attribute_name),
	1616: uint16(1),
	1617: uint16(314),
	1618: uint16(9),
	1619: uint16(anon_sym_GT),
	1620: uint16(anon_sym_SLASH_GT),
	1621: uint16(anon_sym_LBRACE),
	1622: uint16(anon_sym_COLONlet),
	1623: uint16(anon_sym_COLONfor),
	1624: uint16(anon_sym_COLONkey),
	1625: uint16(anon_sym_COLONstream),
	1626: uint16(anon_sym_COLONif),
	1627: uint16(sym_attribute_name),
	1628: uint16(1),
	1629: uint16(316),
	1630: uint16(9),
	1631: uint16(anon_sym_GT),
	1632: uint16(anon_sym_SLASH_GT),
	1633: uint16(anon_sym_LBRACE),
	1634: uint16(anon_sym_COLONlet),
	1635: uint16(anon_sym_COLONfor),
	1636: uint16(anon_sym_COLONkey),
	1637: uint16(anon_sym_COLONstream),
	1638: uint16(anon_sym_COLONif),
	1639: uint16(sym_attribute_name),
	1640: uint16(4),
	1641: uint16(318),
	1642: uint16(1),
	1643: uint16(anon_sym_PERCENT_GT),
	1644: uint16(322),
	1645: uint16(1),
	1646: uint16(sym__code),
	1647: uint16(57),
	1648: uint16(1),
	1649: uint16(aux_sym__hash_comment_repeat1),
	1650: uint16(320),
	1651: uint16(6),
	1652: uint16(anon_sym_do),
	1653: uint16(anon_sym_catch),
	1654: uint16(anon_sym_rescue),
	1655: uint16(anon_sym_after),
	1656: uint16(anon_sym_else),
	1657: uint16(anon_sym_DASH_GT),
	1658: uint16(1),
	1659: uint16(146),
	1660: uint16(9),
	1661: uint16(anon_sym_GT),
	1662: uint16(anon_sym_SLASH_GT),
	1663: uint16(anon_sym_LBRACE),
	1664: uint16(anon_sym_COLONlet),
	1665: uint16(anon_sym_COLONfor),
	1666: uint16(anon_sym_COLONkey),
	1667: uint16(anon_sym_COLONstream),
	1668: uint16(anon_sym_COLONif),
	1669: uint16(sym_attribute_name),
	1670: uint16(1),
	1671: uint16(324),
	1672: uint16(9),
	1673: uint16(anon_sym_GT),
	1674: uint16(anon_sym_SLASH_GT),
	1675: uint16(anon_sym_LBRACE),
	1676: uint16(anon_sym_COLONlet),
	1677: uint16(anon_sym_COLONfor),
	1678: uint16(anon_sym_COLONkey),
	1679: uint16(anon_sym_COLONstream),
	1680: uint16(anon_sym_COLONif),
	1681: uint16(sym_attribute_name),
	1682: uint16(1),
	1683: uint16(118),
	1684: uint16(9),
	1685: uint16(anon_sym_GT),
	1686: uint16(anon_sym_SLASH_GT),
	1687: uint16(anon_sym_LBRACE),
	1688: uint16(anon_sym_COLONlet),
	1689: uint16(anon_sym_COLONfor),
	1690: uint16(anon_sym_COLONkey),
	1691: uint16(anon_sym_COLONstream),
	1692: uint16(anon_sym_COLONif),
	1693: uint16(sym_attribute_name),
	1694: uint16(5),
	1695: uint16(229),
	1696: uint16(1),
	1697: uint16(anon_sym_LBRACE),
	1698: uint16(326),
	1699: uint16(1),
	1700: uint16(sym_attribute_value),
	1701: uint16(328),
	1702: uint16(1),
	1703: uint16(anon_sym_SQUOTE),
	1704: uint16(330),
	1705: uint16(1),
	1706: uint16(anon_sym_DQUOTE),
	1707: uint16(56),
	1708: uint16(2),
	1709: uint16(sym_expression),
	1710: uint16(sym_quoted_attribute_value),
	1711: uint16(5),
	1712: uint16(332),
	1713: uint16(1),
	1714: uint16(anon_sym_LBRACE),
	1715: uint16(334),
	1716: uint16(1),
	1717: uint16(anon_sym_RBRACE),
	1718: uint16(336),
	1719: uint16(1),
	1720: uint16(aux_sym__expression_value_token1),
	1721: uint16(68),
	1722: uint16(1),
	1723: uint16(aux_sym_expression_repeat1),
	1724: uint16(85),
	1725: uint16(1),
	1726: uint16(sym__expression_value),
	1727: uint16(5),
	1728: uint16(338),
	1729: uint16(1),
	1730: uint16(anon_sym_LBRACE),
	1731: uint16(341),
	1732: uint16(1),
	1733: uint16(anon_sym_RBRACE),
	1734: uint16(343),
	1735: uint16(1),
	1736: uint16(aux_sym__expression_value_token1),
	1737: uint16(68),
	1738: uint16(1),
	1739: uint16(aux_sym_expression_repeat1),
	1740: uint16(85),
	1741: uint16(1),
	1742: uint16(sym__expression_value),
	1743: uint16(5),
	1744: uint16(332),
	1745: uint16(1),
	1746: uint16(anon_sym_LBRACE),
	1747: uint16(336),
	1748: uint16(1),
	1749: uint16(aux_sym__expression_value_token1),
	1750: uint16(346),
	1751: uint16(1),
	1752: uint16(anon_sym_RBRACE),
	1753: uint16(71),
	1754: uint16(1),
	1755: uint16(aux_sym_expression_repeat1),
	1756: uint16(85),
	1757: uint16(1),
	1758: uint16(sym__expression_value),
	1759: uint16(5),
	1760: uint16(332),
	1761: uint16(1),
	1762: uint16(anon_sym_LBRACE),
	1763: uint16(336),
	1764: uint16(1),
	1765: uint16(aux_sym__expression_value_token1),
	1766: uint16(348),
	1767: uint16(1),
	1768: uint16(anon_sym_RBRACE),
	1769: uint16(67),
	1770: uint16(1),
	1771: uint16(aux_sym_expression_repeat1),
	1772: uint16(85),
	1773: uint16(1),
	1774: uint16(sym__expression_value),
	1775: uint16(5),
	1776: uint16(332),
	1777: uint16(1),
	1778: uint16(anon_sym_LBRACE),
	1779: uint16(336),
	1780: uint16(1),
	1781: uint16(aux_sym__expression_value_token1),
	1782: uint16(350),
	1783: uint16(1),
	1784: uint16(anon_sym_RBRACE),
	1785: uint16(68),
	1786: uint16(1),
	1787: uint16(aux_sym_expression_repeat1),
	1788: uint16(85),
	1789: uint16(1),
	1790: uint16(sym__expression_value),
	1791: uint16(5),
	1792: uint16(332),
	1793: uint16(1),
	1794: uint16(anon_sym_LBRACE),
	1795: uint16(336),
	1796: uint16(1),
	1797: uint16(aux_sym__expression_value_token1),
	1798: uint16(352),
	1799: uint16(1),
	1800: uint16(anon_sym_RBRACE),
	1801: uint16(73),
	1802: uint16(1),
	1803: uint16(aux_sym_expression_repeat1),
	1804: uint16(85),
	1805: uint16(1),
	1806: uint16(sym__expression_value),
	1807: uint16(5),
	1808: uint16(332),
	1809: uint16(1),
	1810: uint16(anon_sym_LBRACE),
	1811: uint16(336),
	1812: uint16(1),
	1813: uint16(aux_sym__expression_value_token1),
	1814: uint16(354),
	1815: uint16(1),
	1816: uint16(anon_sym_RBRACE),
	1817: uint16(68),
	1818: uint16(1),
	1819: uint16(aux_sym_expression_repeat1),
	1820: uint16(85),
	1821: uint16(1),
	1822: uint16(sym__expression_value),
	1823: uint16(4),
	1824: uint16(356),
	1825: uint16(1),
	1826: uint16(anon_sym_DOT),
	1827: uint16(358),
	1828: uint16(1),
	1829: uint16(sym_module),
	1830: uint16(360),
	1831: uint16(1),
	1832: uint16(sym_tag_name),
	1833: uint16(40),
	1834: uint16(1),
	1835: uint16(sym_component_name),
	1836: uint16(2),
	1837: uint16(364),
	1838: uint16(1),
	1839: uint16(aux_sym__expression_value_token1),
	1840: uint16(362),
	1841: uint16(2),
	1842: uint16(anon_sym_LBRACE),
	1843: uint16(anon_sym_RBRACE),
	1844: uint16(3),
	1845: uint16(366),
	1846: uint16(1),
	1847: uint16(aux_sym__html_comment_token1),
	1848: uint16(369),
	1849: uint16(1),
	1850: uint16(anon_sym_DASH_DASH_PERCENT_GT),
	1851: uint16(76),
	1852: uint16(1),
	1853: uint16(aux_sym__html_comment_repeat1),
	1854: uint16(3),
	1855: uint16(371),
	1856: uint16(1),
	1857: uint16(anon_sym_PERCENT_GT),
	1858: uint16(373),
	1859: uint16(1),
	1860: uint16(sym__code),
	1861: uint16(83),
	1862: uint16(1),
	1863: uint16(aux_sym__hash_comment_repeat1),
	1864: uint16(3),
	1865: uint16(375),
	1866: uint16(1),
	1867: uint16(anon_sym_PERCENT_GT),
	1868: uint16(377),
	1869: uint16(1),
	1870: uint16(sym__code),
	1871: uint16(93),
	1872: uint16(1),
	1873: uint16(aux_sym__hash_comment_repeat1),
	1874: uint16(3),
	1875: uint16(379),
	1876: uint16(1),
	1877: uint16(aux_sym__html_comment_token1),
	1878: uint16(381),
	1879: uint16(1),
	1880: uint16(anon_sym_DASH_DASH_PERCENT_GT),
	1881: uint16(82),
	1882: uint16(1),
	1883: uint16(aux_sym__html_comment_repeat1),
	1884: uint16(2),
	1885: uint16(385),
	1886: uint16(1),
	1887: uint16(aux_sym__expression_value_token1),
	1888: uint16(383),
	1889: uint16(2),
	1890: uint16(anon_sym_LBRACE),
	1891: uint16(anon_sym_RBRACE),
	1892: uint16(3),
	1893: uint16(387),
	1894: uint16(1),
	1895: uint16(aux_sym__html_comment_token1),
	1896: uint16(389),
	1897: uint16(1),
	1898: uint16(anon_sym_DASH_DASH_GT),
	1899: uint16(86),
	1900: uint16(1),
	1901: uint16(aux_sym__html_comment_repeat1),
	1902: uint16(3),
	1903: uint16(391),
	1904: uint16(1),
	1905: uint16(aux_sym__html_comment_token1),
	1906: uint16(393),
	1907: uint16(1),
	1908: uint16(anon_sym_DASH_DASH_PERCENT_GT),
	1909: uint16(76),
	1910: uint16(1),
	1911: uint16(aux_sym__html_comment_repeat1),
	1912: uint16(3),
	1913: uint16(377),
	1914: uint16(1),
	1915: uint16(sym__code),
	1916: uint16(395),
	1917: uint16(1),
	1918: uint16(anon_sym_PERCENT_GT),
	1919: uint16(93),
	1920: uint16(1),
	1921: uint16(aux_sym__hash_comment_repeat1),
	1922: uint16(3),
	1923: uint16(395),
	1924: uint16(1),
	1925: uint16(anon_sym_PERCENT_GT),
	1926: uint16(397),
	1927: uint16(1),
	1928: uint16(sym__code),
	1929: uint16(90),
	1930: uint16(1),
	1931: uint16(aux_sym__hash_comment_repeat1),
	1932: uint16(2),
	1933: uint16(401),
	1934: uint16(1),
	1935: uint16(aux_sym__expression_value_token1),
	1936: uint16(399),
	1937: uint16(2),
	1938: uint16(anon_sym_LBRACE),
	1939: uint16(anon_sym_RBRACE),
	1940: uint16(3),
	1941: uint16(369),
	1942: uint16(1),
	1943: uint16(anon_sym_DASH_DASH_GT),
	1944: uint16(403),
	1945: uint16(1),
	1946: uint16(aux_sym__html_comment_token1),
	1947: uint16(86),
	1948: uint16(1),
	1949: uint16(aux_sym__html_comment_repeat1),
	1950: uint16(3),
	1951: uint16(377),
	1952: uint16(1),
	1953: uint16(sym__code),
	1954: uint16(406),
	1955: uint16(1),
	1956: uint16(anon_sym_PERCENT_GT),
	1957: uint16(93),
	1958: uint16(1),
	1959: uint16(aux_sym__hash_comment_repeat1),
	1960: uint16(3),
	1961: uint16(408),
	1962: uint16(1),
	1963: uint16(anon_sym_PERCENT_GT),
	1964: uint16(410),
	1965: uint16(1),
	1966: uint16(sym__code),
	1967: uint16(78),
	1968: uint16(1),
	1969: uint16(aux_sym__hash_comment_repeat1),
	1970: uint16(3),
	1971: uint16(412),
	1972: uint16(1),
	1973: uint16(anon_sym_PERCENT_GT),
	1974: uint16(414),
	1975: uint16(1),
	1976: uint16(sym__code),
	1977: uint16(87),
	1978: uint16(1),
	1979: uint16(aux_sym__hash_comment_repeat1),
	1980: uint16(3),
	1981: uint16(377),
	1982: uint16(1),
	1983: uint16(sym__code),
	1984: uint16(416),
	1985: uint16(1),
	1986: uint16(anon_sym_PERCENT_GT),
	1987: uint16(93),
	1988: uint16(1),
	1989: uint16(aux_sym__hash_comment_repeat1),
	1990: uint16(3),
	1991: uint16(418),
	1992: uint16(1),
	1993: uint16(aux_sym__html_comment_token1),
	1994: uint16(420),
	1995: uint16(1),
	1996: uint16(anon_sym_DASH_DASH_GT),
	1997: uint16(81),
	1998: uint16(1),
	1999: uint16(aux_sym__html_comment_repeat1),
	2000: uint16(3),
	2001: uint16(356),
	2002: uint16(1),
	2003: uint16(anon_sym_DOT),
	2004: uint16(358),
	2005: uint16(1),
	2006: uint16(sym_module),
	2007: uint16(110),
	2008: uint16(1),
	2009: uint16(sym_component_name),
	2010: uint16(3),
	2011: uint16(305),
	2012: uint16(1),
	2013: uint16(anon_sym_PERCENT_GT),
	2014: uint16(422),
	2015: uint16(1),
	2016: uint16(sym__code),
	2017: uint16(93),
	2018: uint16(1),
	2019: uint16(aux_sym__hash_comment_repeat1),
	2020: uint16(2),
	2021: uint16(425),
	2022: uint16(1),
	2023: uint16(anon_sym_PERCENT_GT),
	2024: uint16(427),
	2025: uint16(1),
	2026: uint16(anon_sym_POUND),
	2027: uint16(2),
	2028: uint16(229),
	2029: uint16(1),
	2030: uint16(anon_sym_LBRACE),
	2031: uint16(61),
	2032: uint16(1),
	2033: uint16(sym_expression),
	2034: uint16(2),
	2035: uint16(429),
	2036: uint16(1),
	2037: uint16(anon_sym_DQUOTE),
	2038: uint16(431),
	2039: uint16(1),
	2040: uint16(aux_sym_quoted_attribute_value_token2),
	2041: uint16(2),
	2042: uint16(433),
	2043: uint16(1),
	2044: uint16(anon_sym_PERCENT_GT),
	2045: uint16(435),
	2046: uint16(1),
	2047: uint16(anon_sym_POUND),
	2048: uint16(2),
	2049: uint16(429),
	2050: uint16(1),
	2051: uint16(anon_sym_SQUOTE),
	2052: uint16(437),
	2053: uint16(1),
	2054: uint16(aux_sym_quoted_attribute_value_token1),
	2055: uint16(1),
	2056: uint16(439),
	2057: uint16(1),
	2058: uint16(anon_sym_html),
	2059: uint16(1),
	2060: uint16(441),
	2061: uint16(1),
	2062: uint16(sym_tag_name),
	2063: uint16(1),
	2064: uint16(443),
	2065: uint16(1),
	2066: uint16(anon_sym_EQ),
	2067: uint16(1),
	2068: uint16(445),
	2069: uint16(1),
	2071: uint16(1),
	2072: uint16(447),
	2073: uint16(1),
	2074: uint16(anon_sym_PERCENT_GT),
	2075: uint16(1),
	2076: uint16(449),
	2077: uint16(1),
	2078: uint16(anon_sym_GT),
	2079: uint16(1),
	2080: uint16(451),
	2081: uint16(1),
	2082: uint16(anon_sym_EQ),
	2083: uint16(1),
	2084: uint16(453),
	2085: uint16(1),
	2086: uint16(anon_sym_GT),
	2087: uint16(1),
	2088: uint16(455),
	2089: uint16(1),
	2090: uint16(anon_sym_SQUOTE),
	2091: uint16(1),
	2092: uint16(455),
	2093: uint16(1),
	2094: uint16(anon_sym_DQUOTE),
	2095: uint16(1),
	2096: uint16(457),
	2097: uint16(1),
	2098: uint16(sym_tag_name),
	2099: uint16(1),
	2100: uint16(459),
	2101: uint16(1),
	2102: uint16(anon_sym_GT),
	2103: uint16(1),
	2104: uint16(461),
	2105: uint16(1),
	2106: uint16(sym_function),
	2107: uint16(1),
	2108: uint16(463),
	2109: uint16(1),
	2110: uint16(anon_sym_DOCTYPE),
	2111: uint16(1),
	2112: uint16(465),
	2113: uint16(1),
	2114: uint16(anon_sym_GT),
	2115: uint16(1),
	2116: uint16(467),
	2117: uint16(1),
	2118: uint16(sym_tag_name),
	2119: uint16(1),
	2120: uint16(469),
	2121: uint16(1),
	2122: uint16(sym_function),
}

var ts_small_parse_table_map = [114]uint32_t{
	1:   uint32(74),
	2:   uint32(148),
	3:   uint32(219),
	4:   uint32(284),
	5:   uint32(348),
	6:   uint32(412),
	7:   uint32(476),
	8:   uint32(540),
	9:   uint32(601),
	10:  uint32(622),
	11:  uint32(643),
	12:  uint32(664),
	13:  uint32(685),
	14:  uint32(706),
	15:  uint32(727),
	16:  uint32(748),
	17:  uint32(769),
	18:  uint32(790),
	19:  uint32(811),
	20:  uint32(832),
	21:  uint32(853),
	22:  uint32(874),
	23:  uint32(895),
	24:  uint32(916),
	25:  uint32(937),
	26:  uint32(958),
	27:  uint32(979),
	28:  uint32(1000),
	29:  uint32(1021),
	30:  uint32(1042),
	31:  uint32(1063),
	32:  uint32(1084),
	33:  uint32(1103),
	34:  uint32(1122),
	35:  uint32(1141),
	36:  uint32(1160),
	37:  uint32(1187),
	38:  uint32(1216),
	39:  uint32(1245),
	40:  uint32(1264),
	41:  uint32(1293),
	42:  uint32(1322),
	43:  uint32(1351),
	44:  uint32(1370),
	45:  uint32(1399),
	46:  uint32(1418),
	47:  uint32(1437),
	48:  uint32(1455),
	49:  uint32(1473),
	50:  uint32(1491),
	51:  uint32(1509),
	52:  uint32(1532),
	53:  uint32(1549),
	54:  uint32(1564),
	55:  uint32(1576),
	56:  uint32(1592),
	57:  uint32(1604),
	58:  uint32(1616),
	59:  uint32(1628),
	60:  uint32(1640),
	61:  uint32(1658),
	62:  uint32(1670),
	63:  uint32(1682),
	64:  uint32(1694),
	65:  uint32(1711),
	66:  uint32(1727),
	67:  uint32(1743),
	68:  uint32(1759),
	69:  uint32(1775),
	70:  uint32(1791),
	71:  uint32(1807),
	72:  uint32(1823),
	73:  uint32(1836),
	74:  uint32(1844),
	75:  uint32(1854),
	76:  uint32(1864),
	77:  uint32(1874),
	78:  uint32(1884),
	79:  uint32(1892),
	80:  uint32(1902),
	81:  uint32(1912),
	82:  uint32(1922),
	83:  uint32(1932),
	84:  uint32(1940),
	85:  uint32(1950),
	86:  uint32(1960),
	87:  uint32(1970),
	88:  uint32(1980),
	89:  uint32(1990),
	90:  uint32(2000),
	91:  uint32(2010),
	92:  uint32(2020),
	93:  uint32(2027),
	94:  uint32(2034),
	95:  uint32(2041),
	96:  uint32(2048),
	97:  uint32(2055),
	98:  uint32(2059),
	99:  uint32(2063),
	100: uint32(2067),
	101: uint32(2071),
	102: uint32(2075),
	103: uint32(2079),
	104: uint32(2083),
	105: uint32(2087),
	106: uint32(2091),
	107: uint32(2095),
	108: uint32(2099),
	109: uint32(2103),
	110: uint32(2107),
	111: uint32(2111),
	112: uint32(2115),
	113: uint32(2119),
}

var ts_parse_actions = [471]TSParseActionEntry{
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token: uint8(TSParseActionTypeReduce),
		Fsymbol:     uint16(sym_fragment),
	})))),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(79)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(89)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(10)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(92)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(100)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fcount: uint8(2),
	}})),
	32: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_component_repeat1),
	})))),
	33: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(112)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	34: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	35: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_component_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_component_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	40: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_component_repeat1),
	})))),
	41: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	42: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	43: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_component_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(69)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_component_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(53)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_component_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(53)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_component_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(91)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_component_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_component_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(89)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_component_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(4)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
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
		Fcount: uint8(2),
	}})),
	66: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
	67: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(112)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(74)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	74: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
	75: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	77: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(53)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
	81: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(53)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	83: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(91)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(79)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(89)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_fragment_repeat1),
	})))),
	93: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(109)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	105: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_fragment),
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
		Fsymbol:      uint16(sym__bang_comment),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__bang_comment),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	111: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__html_comment),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__html_comment),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_tag),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_tag),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(sym_expression),
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
		Fcount: uint8(1),
	}})),
	121: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_expression),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym__hash_comment),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__hash_comment),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_component),
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
		Fsymbol:      uint16(sym_component),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	131: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_tag),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	135: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_component),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_component),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_component),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_self_closing_component),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_expression),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_expression),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__html_comment),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__html_comment),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym__bang_comment),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym__bang_comment),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	159: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__hash_comment),
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
		Fcount: uint8(1),
	}})),
	161: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__hash_comment),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	163: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_tag),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_directive),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_component),
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
		Fcount: uint8(1),
	}})),
	173: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_component),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	175: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_comment),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_comment),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
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
		Fcount: uint8(1),
	}})),
	181: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_doctype),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_self_closing_tag),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_self_closing_component),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_self_closing_component),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_tag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	195: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_component),
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
		Fcount: uint8(1),
	}})),
	197: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_end_component),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_self_closing_slot),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_self_closing_slot),
		Fproduction_id: uint16(4),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_slot),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_slot),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_slot),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_slot),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_component),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	213: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_component),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(72)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	220: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(101)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_start_tag_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(55)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(55)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_component),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_component),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(34)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(51)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(47)),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_slot),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_slot),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_self_closing_slot),
		Fproduction_id: uint16(4),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_self_closing_slot),
		Fproduction_id: uint16(4),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_end_slot),
		Fproduction_id: uint16(4),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_end_slot),
		Fproduction_id: uint16(4),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_start_tag),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_start_slot),
		Fproduction_id: uint16(4),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_start_slot),
		Fproduction_id: uint16(4),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_start_slot),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_start_slot),
		Fproduction_id: uint16(4),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_start_tag),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_component_name),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(111)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_component_name),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_attribute),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(66)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_attribute),
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
		Fcount: uint8(1),
	}})),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__hash_comment_repeat1),
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
		Fcount: uint8(2),
	}})),
	308: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__hash_comment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(57)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_component_name),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_component_name),
	})))),
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
		Fsymbol:      uint16(sym_quoted_attribute_value),
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
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_special_attribute),
	})))),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_expression_value),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(57)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_quoted_attribute_value),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(56)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(70)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Fsymbol:      uint16(aux_sym_expression_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(70)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(aux_sym_expression_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	344: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_expression_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(85)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(14)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(21)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__expression_value),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	365: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym__expression_value),
		Fproduction_id: uint16(1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym__html_comment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(76)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__html_comment_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_partial_expression_value),
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
		Fcount: uint8(1),
	}})),
	374: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_ending_expression_value),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__expression_value),
		Fproduction_id: uint16(3),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym__expression_value),
		Fproduction_id: uint16(3),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(86)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(22)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(76)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Fcount: uint8(1),
	}})),
	396: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_partial_expression_value),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(90)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_expression_repeat1),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_expression_repeat1),
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
		Fcount: uint8(2),
	}})),
	404: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__html_comment_repeat1),
	})))),
	405: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(86)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	406: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	407: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	408: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_ending_expression_value),
	})))),
	410: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	411: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	412: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fcount: uint8(1),
	}})),
	415: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_partial_expression_value),
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
		Fcount: uint8(1),
	}})),
	419: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fcount: uint8(2),
	}})),
	423: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__hash_comment_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(93)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_partial_expression_value),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_partial_expression_value),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(77)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(107)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_special_attribute_name),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	446: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(26)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	454: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	466: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
	}})))),
}

func tree_sitter_heex(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
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
	Fprimary_state_ids:         uintptr(unsafe.Pointer(&ts_primary_state_ids)),
	Fname:                      __ccgo_ts + 851,
	Fmetadata: TSLanguageMetadata{
		Fminor_version: uint8(8),
		Fpatch_version: uint8(1),
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

var __ccgo_ts1 = "end\x00<!\x00DOCTYPE\x00html\x00>\x00<\x00</\x00/>\x00<:\x00</:\x00{\x00}\x00_expression_value_token1\x00=\x00:let\x00:for\x00:key\x00:stream\x00:if\x00attribute_value\x00'\x00\"\x00<%\x00<%=\x00<%%\x00<%%=\x00%>\x00<!--\x00_html_comment_token1\x00-->\x00<%!--\x00--%>\x00<%#\x00do\x00catch\x00rescue\x00after\x00else\x00->\x00#\x00ending_expression_value_token1\x00.\x00module\x00function\x00_code\x00tag_name\x00attribute_name\x00text\x00entity\x00fragment\x00_node\x00doctype\x00tag\x00component\x00slot\x00start_tag\x00end_tag\x00self_closing_tag\x00start_component\x00end_component\x00self_closing_component\x00start_slot\x00end_slot\x00self_closing_slot\x00expression\x00_expression_value\x00special_attribute\x00special_attribute_name\x00attribute\x00quoted_attribute_value\x00directive\x00comment\x00_html_comment\x00_bang_comment\x00_hash_comment\x00expression_value\x00partial_expression_value\x00ending_expression_value\x00component_name\x00fragment_repeat1\x00component_repeat1\x00start_tag_repeat1\x00expression_repeat1\x00_html_comment_repeat1\x00_hash_comment_repeat1\x00left\x00right\x00slot_name\x00heex\x00"
