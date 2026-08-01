// Code generated for darwin/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -D__attribute__(...)= -D__extension__= -D_Nonnull= -D_Nullable= -D_Null_unspecified= -DAPI_AVAILABLE(...)= -DAPI_UNAVAILABLE(...)= -DAPI_DEPRECATED(...)= -DAPI_DEPRECATED_WITH_REPLACEMENT(...)= -D__API_AVAILABLE(...)= -D__API_UNAVAILABLE(...)= -D__API_DEPRECATED(...)= -D__API_DEPRECATED_WITH_REPLACEMENT(...)= -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-doxygen/src -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-doxygen -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build darwin && amd64

package grammar_doxygen

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const ALIAS_COUNT = 1
const BADSIG = "SIG_ERR"
const BIG_ENDIAN = "__DARWIN_BIG_ENDIAN"
const BUFSIZ = 1024
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
const EXTERNAL_TOKEN_COUNT = 5
const FIELD_COUNT = 1
const FILENAME_MAX = 1024
const FOOTPRINT_INTERVAL_RESET = 0x1
const FOPEN_MAX = 20
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
const L_ctermid = 1024
const L_tmpnam = 1024
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
const MAX_ALIAS_SEQUENCE_LENGTH = 7
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
const PRODUCTION_ID_COUNT = 12
const PTRDIFF_MAX = "__PTRDIFF_MAX__"
const P_tmpdir = "/var/tmp/"
const RAND_MAX = 0x7fffffff
const RENAME_EXCL = 0x00000004
const RENAME_NOFOLLOW_ANY = 0x00000010
const RENAME_RESERVED1 = 0x00000008
const RENAME_SECLUDE = 0x00000001
const RENAME_SWAP = 0x00000002
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
const SEEK_CUR = 1
const SEEK_DATA = 4
const SEEK_END = 2
const SEEK_HOLE = 3
const SEEK_SET = 0
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
const STATE_COUNT = 206
const SV_INTERRUPT = "SA_RESTART"
const SV_NOCLDSTOP = "SA_NOCLDSTOP"
const SV_NODEFER = "SA_NODEFER"
const SV_ONSTACK = "SA_ONSTACK"
const SV_RESETHAND = "SA_RESETHAND"
const SV_SIGINFO = "SA_SIGINFO"
const SYMBOL_COUNT = 69
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
const TMP_MAX = 308915776
const TOKEN_COUNT = 47
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
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
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
const __SALC = 0x4000
const __SAPP = 0x0100
const __SCHAR_MAX__ = 127
const __SEG_FS = 1
const __SEG_GS = 1
const __SEOF = 0x0020
const __SERR = 0x0040
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIGN = 0x8000
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
const __SLBF = 0x0001
const __SMBF = 0x0080
const __SMOD = 0x2000
const __SNBF = 0x0002
const __SNPT = 0x0800
const __SOFF = 0x1000
const __SOPT = 0x0400
const __SRD = 0x0004
const __SRW = 0x0010
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE3__ = 1
const __SSE4_1__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP__ = 1
const __SSSE3__ = 1
const __SSTR = 0x0200
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
const __SWR = 0x0008
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
const stderr1 = "__stderrp"
const stdin1 = "__stdinp"
const stdout1 = "__stdoutp"
const sv_onstack = "sv_flags"
const true1 = 1
const ts_builtin_sym_end = 0
const type1 = "type_token"
const var1 = "var_token"

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

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

type size_t = uint64

type ct_rune_t = int32

type rune_t = int32

type wchar_t = int32

type wint_t = int32

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

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type u_int64_t = uint64

type register_t = int64

type intptr_t = int64

type uintptr_t = uint64

type user_addr_t = uint64

type user_size_t = uint64

type user_ssize_t = int64

type user_long_t = int64

type user_ulong_t = uint64

type user_time_t = int64

type user_off_t = int64

type syscall_arg_t = uint64

type va_list = uintptr

type fpos_t = int64

type __sbuf = struct {
	F_base uintptr
	F_size int32
}

type FILE = struct {
	F_p       uintptr
	F_r       int32
	F_w       int32
	F_flags   int16
	F_file    int16
	F_bf      __sbuf
	F_lbfsize int32
	F_cookie  uintptr
	F_close   uintptr
	F_read    uintptr
	F_seek    uintptr
	F_write   uintptr
	F_ub      __sbuf
	F_extra   uintptr
	F_ur      int32
	F_ubuf    [3]uint8
	F_nbuf    [1]uint8
	F_lb      __sbuf
	F_blksize int32
	F_offset  fpos_t
}

type __sFILE = FILE

type off_t = int64

type ssize_t = int64

type wctrans_t = int32

type wctype_t = uint32

type uint64_t = uint64

type int_least64_t = int64

type uint_least64_t = uint64

type int_fast64_t = int64

type uint_fast64_t = uint64

type uint32_t = uint32

type int_least32_t = int32

type uint_least32_t = uint32

type int_fast32_t = int32

type uint_fast32_t = uint32

type uint16_t = uint16

type int_least16_t = int16

type uint_least16_t = uint16

type int_fast16_t = int16

type uint_fast16_t = uint16

type uint8_t = uint8

type int_least8_t = int8

type uint_least8_t = uint8

type int_fast8_t = int8

type uint_fast8_t = uint8

type intmax_t = int64

type uintmax_t = uint64

type idtype_t = int32

const P_ALL = 0
const P_PID = 1
const P_PGID = 2

type pid_t = int32

type id_t = uint32

type sig_atomic_t = int32

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

type TokenType = int32

const BRIEF_TEXT = 0
const CODE_BLOCK_START = 1
const CODE_BLOCK_LANGUAGE = 2
const CODE_BLOCK_CONTENT = 3
const CODE_BLOCK_END = 4

type Scanner = struct {
	Fcodeblock_delimiter_length uint32_t
	Fcodeblock_start_column     uint32_t
}

func advance(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(false1 != 0))
}

func skip(tls *libc.TLS, lexer uintptr) {
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fadvance})))(tls, lexer, libc.BoolUint8(true1 != 0))
}

func tree_sitter_doxygen_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	if (*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_start_column > uint32(255) || (*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length > uint32(255) {
		return uint32(0)
	}
	*(*int8)(unsafe.Pointer(buffer)) = libc.Int8FromUint32((*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length)
	*(*int8)(unsafe.Pointer(buffer + 1)) = libc.Int8FromUint32((*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_start_column)
	return uint32(2)
}

func tree_sitter_doxygen_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var scanner uintptr
	_ = scanner
	scanner = payload
	if length == uint32(2) {
		(*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length = libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(buffer)))
		(*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_start_column = libc.Uint32FromInt8(*(*int8)(unsafe.Pointer(buffer + 1)))
	} else {
		if length != uint32(0) && length != uint32(2) {
			libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts, libc.VaList(bp+8, length))
			libc.Xabort(tls)
		}
	}
}

func tree_sitter_doxygen_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer uintptr, valid_symbols uintptr) (r uint8) {
	var advanced_once uint8
	var col_count, col_count1, column_start, i uint32_t
	var remainder, scanner uintptr
	var v1, v10, v14, v16, v17, v5, v7, v8 int32
	var v12, v3 __darwin_ct_rune_t
	var v13, v4 uint64
	var v19 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = advanced_once, col_count, col_count1, column_start, i, remainder, scanner, v1, v10, v12, v13, v14, v16, v17, v19, v3, v4, v5, v7, v8
	scanner = payload
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(BRIEF_TEXT))) != 0 && !(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CODE_BLOCK_LANGUAGE))) != 0) {
		column_start = uint32(0)
		advanced_once = libc.BoolUint8(false1 != 0)
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
			;
			if !((v1 != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*')) && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0)) {
				break
			}
			skip(tls, lexer)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') || (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return libc.BoolUint8(false1 != 0)
		}
		column_start = (*(*func(*libc.TLS, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fget_column})))(tls, lexer)
		goto content
	content:
		;
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\\') {
			advanced_once = libc.BoolUint8(true1 != 0)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*') {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				advance(tls, lexer)
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') {
					(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BRIEF_TEXT)
					return advanced_once
				}
			} else {
				advance(tls, lexer)
			}
		}
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return libc.BoolUint8(false1 != 0)
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		advance(tls, lexer)
		// go past space, / and * to check next text's column
		for {
			if v19 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n') && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0); v19 {
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
			}
			if !(v19 && (v1 != 0 || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('/') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('*'))) {
				break
			}
			advance(tls, lexer)
		}
		if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) && (*(*func(*libc.TLS, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fget_column})))(tls, lexer) == column_start {
			goto content
		} else {
			if advanced_once != 0 {
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(BRIEF_TEXT)
				return libc.BoolUint8(true1 != 0)
			}
		}
		return libc.BoolUint8(false1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CODE_BLOCK_START))) != 0 {
		for {
			v3 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			v4 = uint64(0x00004000)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _28
		_28:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _25
		_25:
			v1 = v5
			goto _21
		_21:
			;
			if !(v1 != 0 && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0)) {
				break
			}
			skip(tls, lexer)
		}
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return libc.BoolUint8(false1 != 0)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
			(*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_start_column = (*(*func(*libc.TLS, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fget_column})))(tls, lexer)
			advance(tls, lexer)
			(*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length = uint32(1)
			for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
				advance(tls, lexer)
				(*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length = (*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length + 1
			}
			v3 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			v4 = uint64(0x00000100)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _37
		_37:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _34
		_34:
			v1 = v5
			goto _30
		_30:
			if v1 != 0 {
				(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(CODE_BLOCK_START)
				return libc.BoolUint8(true1 != 0)
			}
		}
		return libc.BoolUint8(false1 != 0)
	}
	if v19 = *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CODE_BLOCK_LANGUAGE))) != 0; v19 {
		v3 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
		v4 = libc.Uint64FromInt64(libc.Int64FromInt64(0x00000100) | libc.Int64FromInt64(0x00000400))
		v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
		goto _46
	_46:
		if v8 != 0 {
			v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
		} else {
			v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
		}
		v5 = v7
		goto _43
	_43:
		v1 = v5
		goto _39
	_39:
	}
	if v19 && v1 != 0 {
		for {
			v12 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			v13 = libc.Uint64FromInt64(libc.Int64FromInt64(0x00000100) | libc.Int64FromInt64(0x00000400))
			v17 = libc.BoolInt32(v12 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _56
		_56:
			if v17 != 0 {
				v16 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v12)*4)))&v13 != 0))
			} else {
				v16 = libc.BoolInt32(!!(libc.X__maskrune(tls, v12, v13) != 0))
			}
			v14 = v16
			goto _53
		_53:
			v10 = v14
			goto _49
		_49:
			if !(v10 != 0) {
				break
			}
			advance(tls, lexer)
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		for {
			v3 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			v4 = uint64(0x00004000)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _65
		_65:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _62
		_62:
			v1 = v5
			goto _58
		_58:
			;
			if !(v1 != 0 && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('\n')) {
				break
			}
			advance(tls, lexer)
		}
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(CODE_BLOCK_LANGUAGE)
		return libc.BoolUint8((*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') || (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('}'))
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CODE_BLOCK_CONTENT))) != 0 {
		// optional language
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('{') {
			return libc.BoolUint8(false1 != 0)
		}
		// skip ws and newline before block
		for {
			v3 = (*TSLexer)(unsafe.Pointer(lexer)).Flookahead
			v4 = uint64(0x00004000)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _74
		_74:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _71
		_71:
			v1 = v5
			goto _67
		_67:
			if !(v1 != 0) {
				break
			}
			skip(tls, lexer)
			if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('\n') {
				break
			}
		}
		for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('`') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('@') && !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) {
			advance(tls, lexer)
		}
		if (*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0 {
			return libc.BoolUint8(false1 != 0)
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') && (*(*func(*libc.TLS, uintptr) uint32_t)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fget_column})))(tls, lexer) == (*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_start_column {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
			advance(tls, lexer)
			col_count = uint32(1)
			for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
				advance(tls, lexer)
				col_count = col_count + 1
			}
			if col_count == (*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length {
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(CODE_BLOCK_CONTENT)
				return libc.BoolUint8(true1 != 0)
			}
		}
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('@') {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
			advance(tls, lexer)
			remainder = __ccgo_ts + 73
			i = uint32(0)
			for {
				if !(i < uint32(7)) {
					break
				}
				if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32(*(*int8)(unsafe.Pointer(remainder + uintptr(i)))) {
					return libc.BoolUint8(false1 != 0)
				}
				advance(tls, lexer)
				goto _75
			_75:
				;
				i = i + 1
			}
			(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(CODE_BLOCK_CONTENT)
			return libc.BoolUint8(true1 != 0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CODE_BLOCK_END))) != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
			advance(tls, lexer)
			col_count1 = uint32(1)
			for (*TSLexer)(unsafe.Pointer(lexer)).Flookahead == int32('`') {
				advance(tls, lexer)
				col_count1 = col_count1 + 1
			}
			if col_count1 == (*Scanner)(unsafe.Pointer(scanner)).Fcodeblock_delimiter_length {
				(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(CODE_BLOCK_END)
				return libc.BoolUint8(true1 != 0)
			}
		}
		return libc.BoolUint8(false1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_doxygen_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = libc.Xcalloc(tls, uint64(1), uint64(8))
	return scanner
}

func tree_sitter_doxygen_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var scanner uintptr
	_ = scanner
	scanner = payload
	libc.Xfree(tls, scanner)
}

type ts_symbol_identifiers = int32

const anon_sym_ATbrief = 1
const anon_sym_BSLASHbrief = 2
const aux_sym_brief_header_token1 = 3
const anon_sym_COMMA = 4
const aux_sym_tag_token1 = 5
const aux_sym_tag_token2 = 6
const sym_tag_name_with_argument = 7
const sym_tag_name_with_multiple_arguments = 8
const sym_tag_name_with_types = 9
const sym_tag_name_with_self_types = 10
const sym_tag_name_with_type = 11
const sym_tag_name = 12
const aux_sym_identifier_token1 = 13
const anon_sym_COLON_COLON = 14
const anon_sym_TILDE = 15
const anon_sym_LPAREN = 16
const anon_sym_RPAREN = 17
const anon_sym_LBRACK = 18
const anon_sym_in = 19
const anon_sym_out = 20
const anon_sym_inout = 21
const anon_sym_RBRACK = 22
const anon_sym_BSLASHa = 23
const anon_sym_BSLASHc = 24
const anon_sym_LTa = 25
const aux_sym_link_token1 = 26
const anon_sym_GT = 27
const aux_sym_link_token2 = 28
const anon_sym_LT_SLASHa_GT = 29
const sym_function_link = 30
const anon_sym_ATcode = 31
const anon_sym_LBRACE = 32
const anon_sym_DOT = 33
const anon_sym_RBRACE = 34
const anon_sym_ATendcode = 35
const sym__text = 36
const sym__singleline_begin = 37
const sym__multiline_begin = 38
const anon_sym_SLASH = 39
const anon_sym_STAR_SLASH = 40
const sym__text_line = 41
const sym_brief_text = 42
const sym_code_block_start = 43
const sym_code_block_language = 44
const sym_code_block_content = 45
const sym_code_block_end = 46
const sym_document = 47
const sym_brief_header = 48
const sym_brief_description = 49
const sym_description = 50
const sym_tag = 51
const sym__expression = 52
const sym_identifier = 53
const sym_qualified_identifier = 54
const sym_function = 55
const sym_storageclass = 56
const sym_emphasis = 57
const sym_code_word = 58
const sym_link = 59
const sym_code_block = 60
const sym__multiline_end = 61
const aux_sym_document_repeat1 = 62
const aux_sym_document_repeat2 = 63
const aux_sym_brief_description_repeat1 = 64
const aux_sym_description_repeat1 = 65
const aux_sym_tag_repeat1 = 66
const aux_sym_tag_repeat2 = 67
const aux_sym_qualified_identifier_repeat1 = 68
const alias_sym_code = 69

var ts_symbol_names = [70]uintptr{
	0:  __ccgo_ts + 81,
	1:  __ccgo_ts + 85,
	2:  __ccgo_ts + 85,
	3:  __ccgo_ts + 94,
	4:  __ccgo_ts + 112,
	5:  __ccgo_ts + 114,
	6:  __ccgo_ts + 125,
	7:  __ccgo_ts + 85,
	8:  __ccgo_ts + 85,
	9:  __ccgo_ts + 85,
	10: __ccgo_ts + 85,
	11: __ccgo_ts + 85,
	12: __ccgo_ts + 85,
	13: __ccgo_ts + 130,
	14: __ccgo_ts + 148,
	15: __ccgo_ts + 151,
	16: __ccgo_ts + 153,
	17: __ccgo_ts + 155,
	18: __ccgo_ts + 157,
	19: __ccgo_ts + 159,
	20: __ccgo_ts + 162,
	21: __ccgo_ts + 166,
	22: __ccgo_ts + 172,
	23: __ccgo_ts + 174,
	24: __ccgo_ts + 177,
	25: __ccgo_ts + 180,
	26: __ccgo_ts + 183,
	27: __ccgo_ts + 195,
	28: __ccgo_ts + 197,
	29: __ccgo_ts + 202,
	30: __ccgo_ts + 207,
	31: __ccgo_ts + 221,
	32: __ccgo_ts + 227,
	33: __ccgo_ts + 229,
	34: __ccgo_ts + 231,
	35: __ccgo_ts + 233,
	36: __ccgo_ts + 242,
	37: __ccgo_ts + 248,
	38: __ccgo_ts + 266,
	39: __ccgo_ts + 283,
	40: __ccgo_ts + 285,
	41: __ccgo_ts + 288,
	42: __ccgo_ts + 299,
	43: __ccgo_ts + 310,
	44: __ccgo_ts + 327,
	45: __ccgo_ts + 347,
	46: __ccgo_ts + 366,
	47: __ccgo_ts + 381,
	48: __ccgo_ts + 390,
	49: __ccgo_ts + 94,
	50: __ccgo_ts + 403,
	51: __ccgo_ts + 415,
	52: __ccgo_ts + 419,
	53: __ccgo_ts + 431,
	54: __ccgo_ts + 442,
	55: __ccgo_ts + 463,
	56: __ccgo_ts + 472,
	57: __ccgo_ts + 485,
	58: __ccgo_ts + 494,
	59: __ccgo_ts + 504,
	60: __ccgo_ts + 509,
	61: __ccgo_ts + 520,
	62: __ccgo_ts + 535,
	63: __ccgo_ts + 552,
	64: __ccgo_ts + 569,
	65: __ccgo_ts + 595,
	66: __ccgo_ts + 615,
	67: __ccgo_ts + 627,
	68: __ccgo_ts + 639,
	69: __ccgo_ts + 668,
}

var ts_symbol_map = [70]TSSymbol{
	1:  uint16(sym_tag_name),
	2:  uint16(sym_tag_name),
	3:  uint16(sym_brief_description),
	4:  uint16(anon_sym_COMMA),
	5:  uint16(aux_sym_tag_token1),
	6:  uint16(aux_sym_tag_token2),
	7:  uint16(sym_tag_name),
	8:  uint16(sym_tag_name),
	9:  uint16(sym_tag_name),
	10: uint16(sym_tag_name),
	11: uint16(sym_tag_name),
	12: uint16(sym_tag_name),
	13: uint16(aux_sym_identifier_token1),
	14: uint16(anon_sym_COLON_COLON),
	15: uint16(anon_sym_TILDE),
	16: uint16(anon_sym_LPAREN),
	17: uint16(anon_sym_RPAREN),
	18: uint16(anon_sym_LBRACK),
	19: uint16(anon_sym_in),
	20: uint16(anon_sym_out),
	21: uint16(anon_sym_inout),
	22: uint16(anon_sym_RBRACK),
	23: uint16(anon_sym_BSLASHa),
	24: uint16(anon_sym_BSLASHc),
	25: uint16(anon_sym_LTa),
	26: uint16(aux_sym_link_token1),
	27: uint16(anon_sym_GT),
	28: uint16(aux_sym_link_token2),
	29: uint16(anon_sym_LT_SLASHa_GT),
	30: uint16(sym_function_link),
	31: uint16(anon_sym_ATcode),
	32: uint16(anon_sym_LBRACE),
	33: uint16(anon_sym_DOT),
	34: uint16(anon_sym_RBRACE),
	35: uint16(anon_sym_ATendcode),
	36: uint16(sym__text),
	37: uint16(sym__singleline_begin),
	38: uint16(sym__multiline_begin),
	39: uint16(anon_sym_SLASH),
	40: uint16(anon_sym_STAR_SLASH),
	41: uint16(sym__text_line),
	42: uint16(sym_brief_text),
	43: uint16(sym_code_block_start),
	44: uint16(sym_code_block_language),
	45: uint16(sym_code_block_content),
	46: uint16(sym_code_block_end),
	47: uint16(sym_document),
	48: uint16(sym_brief_header),
	49: uint16(sym_brief_description),
	50: uint16(sym_description),
	51: uint16(sym_tag),
	52: uint16(sym__expression),
	53: uint16(sym_identifier),
	54: uint16(sym_qualified_identifier),
	55: uint16(sym_function),
	56: uint16(sym_storageclass),
	57: uint16(sym_emphasis),
	58: uint16(sym_code_word),
	59: uint16(sym_link),
	60: uint16(sym_code_block),
	61: uint16(sym__multiline_end),
	62: uint16(aux_sym_document_repeat1),
	63: uint16(aux_sym_document_repeat2),
	64: uint16(aux_sym_brief_description_repeat1),
	65: uint16(aux_sym_description_repeat1),
	66: uint16(aux_sym_tag_repeat1),
	67: uint16(aux_sym_tag_repeat2),
	68: uint16(aux_sym_qualified_identifier_repeat1),
	69: uint16(alias_sym_code),
}

var ts_symbol_metadata = [70]TSSymbolMetadata{
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	4: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	5: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	9: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	10: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	12: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	13: {},
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
	26: {},
	27: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	28: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	29: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	30: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	37: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	38: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	39: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	40: {
		Fvisible: libc.BoolUint8(true1 != 0),
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
	45: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	62: {},
	63: {},
	64: {},
	65: {},
	66: {},
	67: {},
	68: {},
	69: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
}

type ts_field_identifiers = int32

const field_function = 1

var ts_field_names = [2]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 463,
}

var ts_field_map_slices = [12]TSFieldMapSlice{
	1: {
		Flength: uint16(1),
	},
	5: {
		Findex:  uint16(1),
		Flength: uint16(2),
	},
	6: {
		Findex:  uint16(3),
		Flength: uint16(2),
	},
	7: {
		Findex:  uint16(5),
		Flength: uint16(2),
	},
	8: {
		Findex:  uint16(7),
		Flength: uint16(2),
	},
	9: {
		Findex:  uint16(9),
		Flength: uint16(3),
	},
	10: {
		Findex:  uint16(12),
		Flength: uint16(2),
	},
	11: {
		Findex:  uint16(14),
		Flength: uint16(3),
	},
}

var ts_field_map_entries = [17]TSFieldMapEntry{
	0: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
	},
	1: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(2),
	},
	3: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
	},
	4: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	5: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
	},
	6: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(3),
	},
	7: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
	},
	8: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	9: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
	},
	10: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(2),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	11: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(3),
	},
	12: {
		Ffield_id:  uint16(field_function),
		Finherited: libc.BoolUint8(true1 != 0),
	},
	13: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	14: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(1),
	},
	15: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(3),
		Finherited:   libc.BoolUint8(true1 != 0),
	},
	16: {
		Ffield_id:    uint16(field_function),
		Fchild_index: uint8(4),
	},
}

var ts_alias_sequences = [12][7]TSSymbol{
	0: {},
	2: {
		1: uint16(aux_sym_tag_token2),
	},
	3: {
		1: uint16(aux_sym_link_token2),
	},
	4: {
		1: uint16(alias_sym_code),
	},
}

var ts_non_terminal_alias_map = [5]uint16_t{
	0: uint16(sym_identifier),
	1: uint16(2),
	2: uint16(sym_identifier),
	3: uint16(aux_sym_tag_token2),
}

var ts_primary_state_ids = [206]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(3),
	4:   uint16(4),
	5:   uint16(5),
	6:   uint16(6),
	7:   uint16(7),
	8:   uint16(3),
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
	19:  uint16(6),
	20:  uint16(9),
	21:  uint16(10),
	22:  uint16(17),
	23:  uint16(7),
	24:  uint16(24),
	25:  uint16(25),
	26:  uint16(11),
	27:  uint16(27),
	28:  uint16(24),
	29:  uint16(29),
	30:  uint16(27),
	31:  uint16(31),
	32:  uint16(32),
	33:  uint16(27),
	34:  uint16(25),
	35:  uint16(35),
	36:  uint16(13),
	37:  uint16(12),
	38:  uint16(14),
	39:  uint16(32),
	40:  uint16(40),
	41:  uint16(41),
	42:  uint16(42),
	43:  uint16(43),
	44:  uint16(27),
	45:  uint16(45),
	46:  uint16(46),
	47:  uint16(47),
	48:  uint16(48),
	49:  uint16(49),
	50:  uint16(16),
	51:  uint16(51),
	52:  uint16(52),
	53:  uint16(53),
	54:  uint16(54),
	55:  uint16(55),
	56:  uint16(56),
	57:  uint16(57),
	58:  uint16(27),
	59:  uint16(15),
	60:  uint16(60),
	61:  uint16(61),
	62:  uint16(42),
	63:  uint16(40),
	64:  uint16(17),
	65:  uint16(65),
	66:  uint16(66),
	67:  uint16(24),
	68:  uint16(27),
	69:  uint16(27),
	70:  uint16(17),
	71:  uint16(71),
	72:  uint16(25),
	73:  uint16(25),
	74:  uint16(24),
	75:  uint16(75),
	76:  uint16(32),
	77:  uint16(27),
	78:  uint16(27),
	79:  uint16(29),
	80:  uint16(31),
	81:  uint16(35),
	82:  uint16(42),
	83:  uint16(27),
	84:  uint16(32),
	85:  uint16(43),
	86:  uint16(40),
	87:  uint16(41),
	88:  uint16(48),
	89:  uint16(60),
	90:  uint16(56),
	91:  uint16(53),
	92:  uint16(49),
	93:  uint16(27),
	94:  uint16(40),
	95:  uint16(47),
	96:  uint16(57),
	97:  uint16(61),
	98:  uint16(42),
	99:  uint16(55),
	100: uint16(100),
	101: uint16(101),
	102: uint16(65),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(106),
	107: uint16(66),
	108: uint16(108),
	109: uint16(109),
	110: uint16(27),
	111: uint16(111),
	112: uint16(112),
	113: uint16(113),
	114: uint16(114),
	115: uint16(115),
	116: uint16(27),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(120),
	121: uint16(75),
	122: uint16(122),
	123: uint16(123),
	124: uint16(124),
	125: uint16(125),
	126: uint16(113),
	127: uint16(118),
	128: uint16(111),
	129: uint16(115),
	130: uint16(130),
	131: uint16(101),
	132: uint16(100),
	133: uint16(105),
	134: uint16(103),
	135: uint16(130),
	136: uint16(114),
	137: uint16(117),
	138: uint16(106),
	139: uint16(109),
	140: uint16(140),
	141: uint16(141),
	142: uint16(140),
	143: uint16(141),
	144: uint16(144),
	145: uint16(144),
	146: uint16(146),
	147: uint16(147),
	148: uint16(147),
	149: uint16(147),
	150: uint16(146),
	151: uint16(151),
	152: uint16(151),
	153: uint16(153),
	154: uint16(153),
	155: uint16(153),
	156: uint16(156),
	157: uint16(157),
	158: uint16(57),
	159: uint16(157),
	160: uint16(153),
	161: uint16(157),
	162: uint16(157),
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
	177: uint16(166),
	178: uint16(165),
	179: uint16(179),
	180: uint16(180),
	181: uint16(164),
	182: uint16(182),
	183: uint16(183),
	184: uint16(179),
	185: uint16(185),
	186: uint16(186),
	187: uint16(174),
	188: uint16(164),
	189: uint16(183),
	190: uint16(183),
	191: uint16(186),
	192: uint16(192),
	193: uint16(183),
	194: uint16(186),
	195: uint16(195),
	196: uint16(169),
	197: uint16(192),
	198: uint16(198),
	199: uint16(169),
	200: uint16(169),
	201: uint16(201),
	202: uint16(176),
	203: uint16(203),
	204: uint16(186),
	205: uint16(182),
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip uint8
	var i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i2, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i3, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i4, i40, i41, i42, i43, i5, i6, i7, i8, i9 uint32_t
	var lookahead int32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i2, i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i3, i30, i31, i32, i33, i34, i35, i36, i37, i38, i39, i4, i40, i41, i42, i43, i5, i6, i7, i8, i9, lookahead, result, skip
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
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(315)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(393)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(356)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(263)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(314)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(316)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(277)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(1):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(393)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(356)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(263)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(277)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(393)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(356)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(263)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(277)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(1)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(392)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(327)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(328)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(399)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(319)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(392)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(327)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(50)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(328)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(4)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(399)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(3)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(319)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(329)
			goto next_state
		}
		return result
	case int32(5):
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
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
		if lookahead != 0 {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(310)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(175)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(6)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(9)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(8)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(12)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(11)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(15)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(14)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(404)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(409)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(262)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(17)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(404)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(409)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(17)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(16)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(262)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(19)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(18)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(276)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('@') || lookahead == int32('\\') {
			state = uint16(133)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(20)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(276)
			goto next_state
		}
		if lookahead == int32('@') || lookahead == int32('\\') {
			state = uint16(133)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(21)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(126)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(22)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(22)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(97)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(126)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(277)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(24)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(277)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(25)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(24)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(404)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(409)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(26)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(269)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(27)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(404)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(409)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(27)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(26)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(269)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(343)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(30)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(29)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(336)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(31)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(391)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(52)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(345)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(32)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(31)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('(') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(42)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(378)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('*') {
			state = uint16(380)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('\n') || lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(38)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('/') {
			state = uint16(394)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('/') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32(':') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32(':') {
			state = uint16(132)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('>') {
			state = uint16(301)
			goto next_state
		}
		return result
	case int32(44):
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
		return result
	case int32(45):
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
		return result
	case int32(46):
		if lookahead == int32('a') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(47):
		if lookahead == int32('a') {
			state = uint16(185)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('a') {
			state = uint16(118)
			goto next_state
		}
		return result
	case int32(49):
		i3 = uint32(0)
		for {
			if !(uint64(i3) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
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
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(50):
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
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(51):
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
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(52):
		i6 = uint32(0)
		for {
			if !(uint64(i6) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
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
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('a') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(54):
		i7 = uint32(0)
		for {
			if !(uint64(i7) < libc.Uint64FromInt64(60)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token7[i7]) == lookahead {
				state = map_token7[i7+uint32(1)]
				goto next_state
			}
			goto _8
		_8:
			;
			i7 = i7 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('g') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(55):
		i8 = uint32(0)
		for {
			if !(uint64(i8) < libc.Uint64FromInt64(56)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token8[i8]) == lookahead {
				state = map_token8[i8+uint32(1)]
				goto next_state
			}
			goto _9
		_9:
			;
			i8 = i8 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('a') {
			state = uint16(115)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('a') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('a') {
			state = uint16(111)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('a') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('a') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('c') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(62):
		if lookahead == int32('c') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(63):
		if lookahead == int32('c') {
			state = uint16(74)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('d') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(65):
		if lookahead == int32('d') {
			state = uint16(124)
			goto next_state
		}
		return result
	case int32(66):
		if lookahead == int32('d') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(67):
		if lookahead == int32('d') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(68):
		if lookahead == int32('d') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('d') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('d') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('e') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('e') {
			state = uint16(312)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('e') {
			state = uint16(317)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('e') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('e') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('e') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('e') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('e') {
			state = uint16(119)
			goto next_state
		}
		return result
	case int32(79):
		if lookahead == int32('e') {
			state = uint16(112)
			goto next_state
		}
		return result
	case int32(80):
		if lookahead == int32('e') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(81):
		if lookahead == int32('e') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(82):
		if lookahead == int32('e') {
			state = uint16(113)
			goto next_state
		}
		return result
	case int32(83):
		if lookahead == int32('f') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('f') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('f') {
			state = uint16(165)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('f') {
			state = uint16(167)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('g') {
			state = uint16(114)
			goto next_state
		}
		return result
	case int32(88):
		if lookahead == int32('i') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(89):
		if lookahead == int32('i') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(90):
		if lookahead == int32('i') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(91):
		if lookahead == int32('l') {
			state = uint16(48)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('l') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('m') {
			state = uint16(183)
			goto next_state
		}
		return result
	case int32(94):
		if lookahead == int32('m') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(95):
		if lookahead == int32('n') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('n') {
			state = uint16(181)
			goto next_state
		}
		return result
	case int32(97):
		if lookahead == int32('n') {
			state = uint16(283)
			goto next_state
		}
		return result
	case int32(98):
		if lookahead == int32('n') {
			state = uint16(67)
			goto next_state
		}
		if lookahead == int32('x') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(99):
		if lookahead == int32('o') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('o') {
			state = uint16(125)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('o') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('o') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(103):
		if lookahead == int32('o') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(104):
		if lookahead == int32('o') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(105):
		if lookahead == int32('p') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(106):
		if lookahead == int32('p') {
			state = uint16(123)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('p') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('p') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('p') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(110):
		if lookahead == int32('r') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(111):
		if lookahead == int32('r') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(112):
		if lookahead == int32('r') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(113):
		if lookahead == int32('r') {
			state = uint16(122)
			goto next_state
		}
		return result
	case int32(114):
		if lookahead == int32('r') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(115):
		if lookahead == int32('r') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(116):
		if lookahead == int32('r') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(117):
		if lookahead == int32('s') {
			state = uint16(189)
			goto next_state
		}
		return result
	case int32(118):
		if lookahead == int32('s') {
			state = uint16(117)
			goto next_state
		}
		return result
	case int32(119):
		if lookahead == int32('s') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(120):
		if lookahead == int32('t') {
			state = uint16(284)
			goto next_state
		}
		return result
	case int32(121):
		if lookahead == int32('t') {
			state = uint16(286)
			goto next_state
		}
		return result
	case int32(122):
		if lookahead == int32('t') {
			state = uint16(130)
			goto next_state
		}
		return result
	case int32(123):
		if lookahead == int32('t') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(124):
		if lookahead == int32('t') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(125):
		if lookahead == int32('u') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(126):
		if lookahead == int32('u') {
			state = uint16(120)
			goto next_state
		}
		return result
	case int32(127):
		if lookahead == int32('u') {
			state = uint16(121)
			goto next_state
		}
		return result
	case int32(128):
		if lookahead == int32('v') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(129):
		if lookahead == int32('x') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(130):
		if lookahead == int32('y') {
			state = uint16(187)
			goto next_state
		}
		return result
	case int32(131):
		if lookahead == int32('y') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(132):
		if lookahead == int32('~') {
			state = uint16(137)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(133):
		if lookahead == int32('~') {
			state = uint16(135)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(191)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(134):
		if lookahead == int32('~') {
			state = uint16(136)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(135):
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(191)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(136):
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(137):
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(138):
		if lookahead != 0 && lookahead != int32('.') && lookahead != int32('<') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(139):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(39)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(393)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(356)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(263)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(277)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(140):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(2)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(393)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(344)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(356)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(44)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(45)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('i') {
			state = uint16(263)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(266)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(277)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(140)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(139)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(346)
			goto next_state
		}
		return result
	case int32(141):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(388)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(137)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(141)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(143)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(142):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(315)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(388)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('{') {
			state = uint16(314)
			goto next_state
		}
		if lookahead == int32('}') {
			state = uint16(316)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(137)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(143)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(143):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(141)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(280)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(388)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(40)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(297)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32(']') {
			state = uint16(288)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(137)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(143)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(144):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(373)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(374)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(138)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(145)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(365)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(145):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(144)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(373)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(49)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(54)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(374)
			goto next_state
		}
		if lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(138)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(145)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(365)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') {
			state = uint16(375)
			goto next_state
		}
		return result
	case int32(146):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(146)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(146)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(278)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(146)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(147)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(272)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(147):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(146)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(278)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(147)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(272)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(148):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(148)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(149)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(149):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(149)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(150):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(148)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(358)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(149)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(151):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(151)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(151)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(151)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(152)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(152):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(151)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(152)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(153):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(151)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(358)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(152)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(154):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(154)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(155)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(155):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(155)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(156):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(154)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(279)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(358)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(155)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(157):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(137)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(157)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(158)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(158):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(157)
			goto next_state
		}
		if lookahead == int32(',') {
			state = uint16(176)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(41)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(137)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(158)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(159):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(159)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(160)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(160):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(160)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(161):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(159)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(358)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(160)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(348)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(162):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32('*') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(162)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(163)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(272)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(163):
		if eof != 0 {
			state = uint16(164)
			goto next_state
		}
		if lookahead == int32('\n') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(162)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(357)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(359)
			goto next_state
		}
		if lookahead == int32('@') {
			state = uint16(51)
			goto next_state
		}
		if lookahead == int32('[') {
			state = uint16(281)
			goto next_state
		}
		if lookahead == int32('\\') {
			state = uint16(55)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(361)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			skip = libc.BoolUint8(true1 != 0)
			state = uint16(163)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(272)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('*') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(362)
			goto next_state
		}
		return result
	case int32(164):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(165):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATbrief)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(166):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATbrief)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(167):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASHbrief)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(168):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASHbrief)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(169):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_brief_header_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(170):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_brief_header_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(401)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(377)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(402)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(171):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_brief_header_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(378)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(376)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(172):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_brief_header_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(377)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(402)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(401)
			goto next_state
		}
		return result
	case int32(173):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_brief_header_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(378)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(174):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_brief_header_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(402)
			goto next_state
		}
		return result
	case int32(175):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_brief_header_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(176):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(177):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_tag_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(178)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(177)
			goto next_state
		}
		return result
	case int32(178):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_tag_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(178)
			goto next_state
		}
		return result
	case int32(179):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_tag_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\t') || int32(0x0b) <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(179)
			goto next_state
		}
		if lookahead != 0 && (lookahead < int32('\t') || int32('\r') < lookahead) {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(180):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_tag_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(180)
			goto next_state
		}
		return result
	case int32(181):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_argument)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(182):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_argument)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(183):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_multiple_arguments)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(184):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_multiple_arguments)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(185):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_types)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(186):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_types)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(187):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_self_types)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(188):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_self_types)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(189):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(190):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name_with_type)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(191):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(192):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(230)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(193):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(248)
			goto next_state
		}
		if lookahead == int32('r') {
			state = uint16(237)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(194):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(186)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(195):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(251)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(196):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(229)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(197):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(243)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(198):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(202)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(199):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('a') {
			state = uint16(201)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('b') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(200):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(211)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(201):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('c') {
			state = uint16(208)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(202):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(188)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(203):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(255)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(204):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(203)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(205):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(209)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(206):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(210)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(207):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(219)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(208):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(182)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(209):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(313)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(210):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(218)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(211):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(239)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(212):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(206)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(213):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(220)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(214):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(244)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(215):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(252)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(216):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(221)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(217):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('e') {
			state = uint16(247)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(218):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(188)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(219):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(222)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(220):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(166)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(221):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('f') {
			state = uint16(168)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(222):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('g') {
			state = uint16(246)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(223):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(235)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(224):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(213)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(225):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('i') {
			state = uint16(216)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(226):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(195)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(205)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(227):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(195)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(228):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(236)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(229):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(184)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(230):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('m') {
			state = uint16(215)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(231):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(188)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(232):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('n') {
			state = uint16(182)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(233):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(222)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(234):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(256)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(235):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(232)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(236):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(198)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(237):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(242)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(238):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(190)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(239):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(253)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(240):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(212)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(241):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(199)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(242):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('p') {
			state = uint16(217)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(243):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(188)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(244):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(228)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(245):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(224)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(246):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(234)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(247):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(254)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(248):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(196)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(249):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('r') {
			state = uint16(225)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(250):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(190)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(251):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(250)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(252):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('s') {
			state = uint16(241)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(253):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(223)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(254):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(259)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(255):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('t') {
			state = uint16(233)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(256):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('u') {
			state = uint16(238)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(257):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('v') {
			state = uint16(214)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(258):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('x') {
			state = uint16(200)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(259):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(188)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(260):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('y') {
			state = uint16(240)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(261):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_tag_name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(262):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(178)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(402)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(405)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(177)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(262)
			goto next_state
		}
		return result
	case int32(263):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if lookahead == int32('n') {
			state = uint16(282)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(264):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(285)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(265):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if lookahead == int32('t') {
			state = uint16(287)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(266):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(264)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(267):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if lookahead == int32('u') {
			state = uint16(265)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(268):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(269):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(402)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(405)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(269)
			goto next_state
		}
		return result
	case int32(270):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(42)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(178)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(270)
			goto next_state
		}
		return result
	case int32(271):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(35)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(42)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(271)
			goto next_state
		}
		return result
	case int32(272):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(376)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(349)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(272)
			goto next_state
		}
		return result
	case int32(273):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_identifier_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(273)
			goto next_state
		}
		return result
	case int32(274):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('~') {
			state = uint16(341)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(275):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COLON_COLON)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('~') {
			state = uint16(353)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(276):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(277):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(335)
			goto next_state
		}
		return result
	case int32(278):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_TILDE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(347)
			goto next_state
		}
		return result
	case int32(279):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(280):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(281):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(282):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_in)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if lookahead == int32('o') {
			state = uint16(267)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(283):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_in)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('o') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(284):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_out)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(285):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_out)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(286):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_inout)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(287):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_inout)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(268)
			goto next_state
		}
		return result
	case int32(288):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(289):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASHa)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(290):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASHa)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('d') {
			state = uint16(203)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(291):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASHc)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(292):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_BSLASHc)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('l') {
			state = uint16(195)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(293):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LTa)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(294):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_link_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(294)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(294)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(294)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(295)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(295):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_link_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(294)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(295)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(296):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_link_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('>') {
			state = uint16(296)
			goto next_state
		}
		return result
	case int32(297):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(298):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_link_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(298)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(298)
			goto next_state
		}
		if lookahead == int32('\t') || lookahead == int32(' ') {
			state = uint16(298)
			goto next_state
		}
		if int32(0x0b) <= lookahead && lookahead <= int32('\r') {
			state = uint16(299)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(299):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_link_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(298)
			goto next_state
		}
		if int32('\t') <= lookahead && lookahead <= int32('\r') || lookahead == int32(' ') {
			state = uint16(299)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(300):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_link_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(300)
			goto next_state
		}
		return result
	case int32(301):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASHa_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(302):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(303):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i9 = uint32(0)
		for {
			if !(uint64(i9) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token9[i9]) == lookahead {
				state = map_token9[i9+uint32(1)]
				goto next_state
			}
			goto _10
		_10:
			;
			i9 = i9 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(303)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(304):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i10 = uint32(0)
		for {
			if !(uint64(i10) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token10[i10]) == lookahead {
				state = map_token10[i10+uint32(1)]
				goto next_state
			}
			goto _11
		_11:
			;
			i10 = i10 + uint32(2)
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(305):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(334)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(398)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(397)
			goto next_state
		}
		return result
	case int32(306):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(398)
			goto next_state
		}
		return result
	case int32(307):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i11 = uint32(0)
		for {
			if !(uint64(i11) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token11[i11]) == lookahead {
				state = map_token11[i11+uint32(1)]
				goto next_state
			}
			goto _12
		_12:
			;
			i11 = i11 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(307)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(308):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i12 = uint32(0)
		for {
			if !(uint64(i12) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token12[i12]) == lookahead {
				state = map_token12[i12+uint32(1)]
				goto next_state
			}
			goto _13
		_13:
			;
			i12 = i12 + uint32(2)
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(309):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(380)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('\n') || lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(38)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(310):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('<') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(311):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_function_link)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		return result
	case int32(312):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATcode)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(313):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATcode)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(261)
			goto next_state
		}
		return result
	case int32(314):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(315):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(316):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(317):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATendcode)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(318):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i13 = uint32(0)
		for {
			if !(uint64(i13) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token13[i13]) == lookahead {
				state = map_token13[i13+uint32(1)]
				goto next_state
			}
			goto _14
		_14:
			;
			i13 = i13 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(318)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(319):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i14 = uint32(0)
		for {
			if !(uint64(i14) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token14[i14]) == lookahead {
				state = map_token14[i14+uint32(1)]
				goto next_state
			}
			goto _15
		_15:
			;
			i14 = i14 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(318)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(320):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i15 = uint32(0)
		for {
			if !(uint64(i15) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token15[i15]) == lookahead {
				state = map_token15[i15+uint32(1)]
				goto next_state
			}
			goto _16
		_16:
			;
			i15 = i15 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(321):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i16 = uint32(0)
		for {
			if !(uint64(i16) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token16[i16]) == lookahead {
				state = map_token16[i16+uint32(1)]
				goto next_state
			}
			goto _17
		_17:
			;
			i16 = i16 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(303)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(322):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i17 = uint32(0)
		for {
			if !(uint64(i17) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token17[i17]) == lookahead {
				state = map_token17[i17+uint32(1)]
				goto next_state
			}
			goto _18
		_18:
			;
			i17 = i17 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(318)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(323):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i18 = uint32(0)
		for {
			if !(uint64(i18) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token18[i18]) == lookahead {
				state = map_token18[i18+uint32(1)]
				goto next_state
			}
			goto _19
		_19:
			;
			i18 = i18 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(318)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(324):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i19 = uint32(0)
		for {
			if !(uint64(i19) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token19[i19]) == lookahead {
				state = map_token19[i19+uint32(1)]
				goto next_state
			}
			goto _20
		_20:
			;
			i19 = i19 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(303)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(325):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i20 = uint32(0)
		for {
			if !(uint64(i20) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token20[i20]) == lookahead {
				state = map_token20[i20+uint32(1)]
				goto next_state
			}
			goto _21
		_21:
			;
			i20 = i20 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(326):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i21 = uint32(0)
		for {
			if !(uint64(i21) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token21[i21]) == lookahead {
				state = map_token21[i21+uint32(1)]
				goto next_state
			}
			goto _22
		_22:
			;
			i21 = i21 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(327):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i22 = uint32(0)
		for {
			if !(uint64(i22) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token22[i22]) == lookahead {
				state = map_token22[i22+uint32(1)]
				goto next_state
			}
			goto _23
		_23:
			;
			i22 = i22 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(328):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i23 = uint32(0)
		for {
			if !(uint64(i23) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token23[i23]) == lookahead {
				state = map_token23[i23+uint32(1)]
				goto next_state
			}
			goto _24
		_24:
			;
			i23 = i23 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(318)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(329):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i24 = uint32(0)
		for {
			if !(uint64(i24) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token24[i24]) == lookahead {
				state = map_token24[i24+uint32(1)]
				goto next_state
			}
			goto _25
		_25:
			;
			i24 = i24 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(330):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i25 = uint32(0)
		for {
			if !(uint64(i25) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token25[i25]) == lookahead {
				state = map_token25[i25+uint32(1)]
				goto next_state
			}
			goto _26
		_26:
			;
			i25 = i25 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(330)
			goto next_state
		}
		return result
	case int32(331):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i26 = uint32(0)
		for {
			if !(uint64(i26) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token26[i26]) == lookahead {
				state = map_token26[i26+uint32(1)]
				goto next_state
			}
			goto _27
		_27:
			;
			i26 = i26 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(331)
			goto next_state
		}
		return result
	case int32(332):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i27 = uint32(0)
		for {
			if !(uint64(i27) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token27[i27]) == lookahead {
				state = map_token27[i27+uint32(1)]
				goto next_state
			}
			goto _28
		_28:
			;
			i27 = i27 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(333):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i28 = uint32(0)
		for {
			if !(uint64(i28) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token28[i28]) == lookahead {
				state = map_token28[i28+uint32(1)]
				goto next_state
			}
			goto _29
		_29:
			;
			i28 = i28 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(5)
			goto next_state
		}
		return result
	case int32(334):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i29 = uint32(0)
		for {
			if !(uint64(i29) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token29[i29]) == lookahead {
				state = map_token29[i29+uint32(1)]
				goto next_state
			}
			goto _30
		_30:
			;
			i29 = i29 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(397)
			goto next_state
		}
		return result
	case int32(335):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(342)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(335)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(336):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(346)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(337)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(335)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(337):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(342)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(339)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(338):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(342)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(341)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(339):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(342)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(340)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(335)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(340):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(342)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(335)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(341):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(342)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(342):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(342)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(343):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(346)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(274)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(344):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(346)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(338)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(345):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(346)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(335)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(346):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(346)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(342)
			goto next_state
		}
		return result
	case int32(347):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(376)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(349)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(347)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(348):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('(') {
			state = uint16(376)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(349)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(347)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(349):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(352)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(350):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32('>') {
			state = uint16(301)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(351):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(353)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(352):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if lookahead == int32('~') {
			state = uint16(354)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(347)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') && lookahead != int32('~') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(353):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(354):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(347)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(355):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(355)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(356):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(360)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(357):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(351)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(358):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(275)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(359):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(293)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(360):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(362)
			goto next_state
		}
		if lookahead == int32('a') {
			state = uint16(350)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(361):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(362)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(347)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && (lookahead < int32('a') || int32('{') < lookahead) && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(362):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(37)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(362)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('*') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(355)
			goto next_state
		}
		return result
	case int32(363):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(401)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(377)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(363)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(402)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(363)
			goto next_state
		}
		return result
	case int32(364):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i30 = uint32(0)
		for {
			if !(uint64(i30) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token30[i30]) == lookahead {
				state = map_token30[i30+uint32(1)]
				goto next_state
			}
			goto _31
		_31:
			;
			i30 = i30 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(364)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(365):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i31 = uint32(0)
		for {
			if !(uint64(i31) < libc.Uint64FromInt64(44)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token31[i31]) == lookahead {
				state = map_token31[i31+uint32(1)]
				goto next_state
			}
			goto _32
		_32:
			;
			i31 = i31 + uint32(2)
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(364)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(366):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i32 = uint32(0)
		for {
			if !(uint64(i32) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token32[i32]) == lookahead {
				state = map_token32[i32+uint32(1)]
				goto next_state
			}
			goto _33
		_33:
			;
			i32 = i32 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(367):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i33 = uint32(0)
		for {
			if !(uint64(i33) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token33[i33]) == lookahead {
				state = map_token33[i33+uint32(1)]
				goto next_state
			}
			goto _34
		_34:
			;
			i33 = i33 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(307)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(368):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i34 = uint32(0)
		for {
			if !(uint64(i34) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token34[i34]) == lookahead {
				state = map_token34[i34+uint32(1)]
				goto next_state
			}
			goto _35
		_35:
			;
			i34 = i34 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(364)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(369):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i35 = uint32(0)
		for {
			if !(uint64(i35) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token35[i35]) == lookahead {
				state = map_token35[i35+uint32(1)]
				goto next_state
			}
			goto _36
		_36:
			;
			i35 = i35 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(364)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(370):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i36 = uint32(0)
		for {
			if !(uint64(i36) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token36[i36]) == lookahead {
				state = map_token36[i36+uint32(1)]
				goto next_state
			}
			goto _37
		_37:
			;
			i36 = i36 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(307)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(371):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i37 = uint32(0)
		for {
			if !(uint64(i37) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token37[i37]) == lookahead {
				state = map_token37[i37+uint32(1)]
				goto next_state
			}
			goto _38
		_38:
			;
			i37 = i37 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(372):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i38 = uint32(0)
		for {
			if !(uint64(i38) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token38[i38]) == lookahead {
				state = map_token38[i38+uint32(1)]
				goto next_state
			}
			goto _39
		_39:
			;
			i38 = i38 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(373):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i39 = uint32(0)
		for {
			if !(uint64(i39) < libc.Uint64FromInt64(40)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token39[i39]) == lookahead {
				state = map_token39[i39+uint32(1)]
				goto next_state
			}
			goto _40
		_40:
			;
			i39 = i39 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(374):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i40 = uint32(0)
		for {
			if !(uint64(i40) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token40[i40]) == lookahead {
				state = map_token40[i40+uint32(1)]
				goto next_state
			}
			goto _41
		_41:
			;
			i40 = i40 + uint32(2)
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(364)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(375):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i41 = uint32(0)
		for {
			if !(uint64(i41) < libc.Uint64FromInt64(36)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token41[i41]) == lookahead {
				state = map_token41[i41+uint32(1)]
				goto next_state
			}
			goto _42
		_42:
			;
			i41 = i41 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(371)
			goto next_state
		}
		return result
	case int32(376):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(34)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(378)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(376)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(376)
			goto next_state
		}
		return result
	case int32(377):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(377)
			goto next_state
		}
		if lookahead == int32('/') || lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(402)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(401)
			goto next_state
		}
		return result
	case int32(378):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(378)
			goto next_state
		}
		if lookahead == int32('/') || lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(35)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(379):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(403)
			goto next_state
		}
		return result
	case int32(380):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i42 = uint32(0)
		for {
			if !(uint64(i42) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token42[i42]) == lookahead {
				state = map_token42[i42+uint32(1)]
				goto next_state
			}
			goto _43
		_43:
			;
			i42 = i42 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(381):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(381)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') && lookahead != int32('/') && lookahead != int32('\\') && lookahead != int32('{') && lookahead != int32('}') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(382):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__singleline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(383):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__singleline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(385)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(389)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(387)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(384):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__singleline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(385)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(382)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32('/') {
			state = uint16(389)
			goto next_state
		}
		return result
	case int32(385):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__singleline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('<') {
			state = uint16(382)
			goto next_state
		}
		return result
	case int32(386):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__multiline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(387):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__multiline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(390)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(389)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(387)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(388):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__multiline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(390)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(389)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(384)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(389):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__multiline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(390)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(386)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32('/') {
			state = uint16(389)
			goto next_state
		}
		return result
	case int32(390):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__multiline_begin)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('<') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(391):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(392):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('!') {
			state = uint16(397)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(326)
			goto next_state
		}
		if lookahead == int32('*') || lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(398)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('.') && lookahead != int32('/') && lookahead != int32('<') {
			state = uint16(325)
			goto next_state
		}
		return result
	case int32(393):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('!') {
			state = uint16(390)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(389)
			goto next_state
		}
		if lookahead == int32('/') {
			state = uint16(383)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(386)
			goto next_state
		}
		return result
	case int32(394):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(395):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		i43 = uint32(0)
		for {
			if !(uint64(i43) < libc.Uint64FromInt64(32)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token43[i43]) == lookahead {
				state = map_token43[i43+uint32(1)]
				goto next_state
			}
			goto _44
		_44:
			;
			i43 = i43 + uint32(2)
		}
		if lookahead != 0 {
			state = uint16(395)
			goto next_state
		}
		return result
	case int32(396):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32(')') {
			state = uint16(306)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(174)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(402)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(396)
			goto next_state
		}
		return result
	case int32(397):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(334)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(403)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(398)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(397)
			goto next_state
		}
		return result
	case int32(398):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('.') {
			state = uint16(169)
			goto next_state
		}
		if lookahead == int32('<') {
			state = uint16(410)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(398)
			goto next_state
		}
		return result
	case int32(399):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('\n') {
			state = uint16(38)
			goto next_state
		}
		if lookahead == int32('.') || lookahead == int32('<') {
			state = uint16(410)
			goto next_state
		}
		if lookahead != 0 {
			state = uint16(398)
			goto next_state
		}
		return result
	case int32(400):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('(') {
			state = uint16(402)
			goto next_state
		}
		if lookahead == int32(':') {
			state = uint16(405)
			goto next_state
		}
		if int32('0') <= lookahead && lookahead <= int32('9') || int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(400)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(401):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead == int32('*') {
			state = uint16(377)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(402)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(401)
			goto next_state
		}
		return result
	case int32(402):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(')') {
			state = uint16(302)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(402)
			goto next_state
		}
		return result
	case int32(403):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('*') {
			state = uint16(379)
			goto next_state
		}
		if lookahead == int32('\\') || lookahead == int32('{') || lookahead == int32('}') {
			state = uint16(410)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(403)
			goto next_state
		}
		return result
	case int32(404):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(407)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(405):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32(':') {
			state = uint16(406)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(406):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('~') {
			state = uint16(409)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(400)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(407):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('~') {
			state = uint16(408)
			goto next_state
		}
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(408):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(311)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(409):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('A') <= lookahead && lookahead <= int32('Z') || lookahead == int32('_') || int32('a') <= lookahead && lookahead <= int32('z') {
			state = uint16(400)
			goto next_state
		}
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	case int32(410):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__text_line)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead != 0 && lookahead != int32('\n') {
			state = uint16(410)
			goto next_state
		}
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [16]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16(')'),
	3:  uint16(309),
	4:  uint16('*'),
	5:  uint16(333),
	6:  uint16('.'),
	7:  uint16(173),
	8:  uint16('<'),
	9:  uint16(34),
	10: uint16('\\'),
	11: uint16(6),
	12: uint16('{'),
	13: uint16(6),
	14: uint16('}'),
	15: uint16(6),
}

var map_token1 = [24]uint16_t{
	0:  uint16('a'),
	1:  uint16(66),
	2:  uint16('b'),
	3:  uint16(110),
	4:  uint16('c'),
	5:  uint16(91),
	6:  uint16('d'),
	7:  uint16(71),
	8:  uint16('e'),
	9:  uint16(98),
	10: uint16('f'),
	11: uint16(95),
	12: uint16('n'),
	13: uint16(46),
	14: uint16('o'),
	15: uint16(128),
	16: uint16('p'),
	17: uint16(56),
	18: uint16('s'),
	19: uint16(47),
	20: uint16('t'),
	21: uint16(131),
	22: uint16('v'),
	23: uint16(58),
}

var map_token2 = [24]uint16_t{
	0:  uint16('a'),
	1:  uint16(289),
	2:  uint16('b'),
	3:  uint16(116),
	4:  uint16('c'),
	5:  uint16(291),
	6:  uint16('d'),
	7:  uint16(71),
	8:  uint16('e'),
	9:  uint16(129),
	10: uint16('f'),
	11: uint16(95),
	12: uint16('n'),
	13: uint16(46),
	14: uint16('o'),
	15: uint16(128),
	16: uint16('p'),
	17: uint16(56),
	18: uint16('s'),
	19: uint16(47),
	20: uint16('t'),
	21: uint16(131),
	22: uint16('v'),
	23: uint16(58),
}

var map_token3 = [30]uint16_t{
	0:  uint16('a'),
	1:  uint16(204),
	2:  uint16('b'),
	3:  uint16(245),
	4:  uint16('c'),
	5:  uint16(227),
	6:  uint16('d'),
	7:  uint16(207),
	8:  uint16('e'),
	9:  uint16(258),
	10: uint16('f'),
	11: uint16(231),
	12: uint16('n'),
	13: uint16(192),
	14: uint16('o'),
	15: uint16(257),
	16: uint16('p'),
	17: uint16(193),
	18: uint16('s'),
	19: uint16(194),
	20: uint16('t'),
	21: uint16(260),
	22: uint16('v'),
	23: uint16(197),
	24: uint16('~'),
	25: uint16(135),
	26: uint16('{'),
	27: uint16(191),
	28: uint16('}'),
	29: uint16(191),
}

var map_token4 = [30]uint16_t{
	0:  uint16('a'),
	1:  uint16(204),
	2:  uint16('b'),
	3:  uint16(245),
	4:  uint16('c'),
	5:  uint16(226),
	6:  uint16('d'),
	7:  uint16(207),
	8:  uint16('e'),
	9:  uint16(258),
	10: uint16('f'),
	11: uint16(231),
	12: uint16('n'),
	13: uint16(192),
	14: uint16('o'),
	15: uint16(257),
	16: uint16('p'),
	17: uint16(193),
	18: uint16('s'),
	19: uint16(194),
	20: uint16('t'),
	21: uint16(260),
	22: uint16('v'),
	23: uint16(197),
	24: uint16('~'),
	25: uint16(135),
	26: uint16('{'),
	27: uint16(191),
	28: uint16('}'),
	29: uint16(191),
}

var map_token5 = [28]uint16_t{
	0:  uint16('a'),
	1:  uint16(204),
	2:  uint16('c'),
	3:  uint16(227),
	4:  uint16('d'),
	5:  uint16(207),
	6:  uint16('e'),
	7:  uint16(258),
	8:  uint16('f'),
	9:  uint16(231),
	10: uint16('n'),
	11: uint16(192),
	12: uint16('o'),
	13: uint16(257),
	14: uint16('p'),
	15: uint16(193),
	16: uint16('s'),
	17: uint16(194),
	18: uint16('t'),
	19: uint16(260),
	20: uint16('v'),
	21: uint16(197),
	22: uint16('~'),
	23: uint16(135),
	24: uint16('{'),
	25: uint16(191),
	26: uint16('}'),
	27: uint16(191),
}

var map_token6 = [28]uint16_t{
	0:  uint16('a'),
	1:  uint16(204),
	2:  uint16('c'),
	3:  uint16(226),
	4:  uint16('d'),
	5:  uint16(207),
	6:  uint16('e'),
	7:  uint16(258),
	8:  uint16('f'),
	9:  uint16(231),
	10: uint16('n'),
	11: uint16(192),
	12: uint16('o'),
	13: uint16(257),
	14: uint16('p'),
	15: uint16(193),
	16: uint16('s'),
	17: uint16(194),
	18: uint16('t'),
	19: uint16(260),
	20: uint16('v'),
	21: uint16(197),
	22: uint16('~'),
	23: uint16(135),
	24: uint16('{'),
	25: uint16(191),
	26: uint16('}'),
	27: uint16(191),
}

var map_token7 = [30]uint16_t{
	0:  uint16('a'),
	1:  uint16(290),
	2:  uint16('b'),
	3:  uint16(249),
	4:  uint16('c'),
	5:  uint16(292),
	6:  uint16('d'),
	7:  uint16(207),
	8:  uint16('e'),
	9:  uint16(258),
	10: uint16('f'),
	11: uint16(231),
	12: uint16('n'),
	13: uint16(192),
	14: uint16('o'),
	15: uint16(257),
	16: uint16('p'),
	17: uint16(193),
	18: uint16('s'),
	19: uint16(194),
	20: uint16('t'),
	21: uint16(260),
	22: uint16('v'),
	23: uint16(197),
	24: uint16('~'),
	25: uint16(135),
	26: uint16('{'),
	27: uint16(191),
	28: uint16('}'),
	29: uint16(191),
}

var map_token8 = [28]uint16_t{
	0:  uint16('a'),
	1:  uint16(290),
	2:  uint16('c'),
	3:  uint16(292),
	4:  uint16('d'),
	5:  uint16(207),
	6:  uint16('e'),
	7:  uint16(258),
	8:  uint16('f'),
	9:  uint16(231),
	10: uint16('n'),
	11: uint16(192),
	12: uint16('o'),
	13: uint16(257),
	14: uint16('p'),
	15: uint16(193),
	16: uint16('s'),
	17: uint16(194),
	18: uint16('t'),
	19: uint16(260),
	20: uint16('v'),
	21: uint16(197),
	22: uint16('~'),
	23: uint16(135),
	24: uint16('{'),
	25: uint16(191),
	26: uint16('}'),
	27: uint16(191),
}

var map_token9 = [16]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16('\\'),
	11: uint16(398),
	12: uint16('{'),
	13: uint16(398),
	14: uint16('}'),
	15: uint16(398),
}

var map_token10 = [16]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16('\\'),
	11: uint16(398),
	12: uint16('{'),
	13: uint16(398),
	14: uint16('}'),
	15: uint16(398),
}

var map_token11 = [16]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16('\n'),
	9:  uint16(38),
	10: uint16('\\'),
	11: uint16(38),
	12: uint16('{'),
	13: uint16(38),
	14: uint16('}'),
	15: uint16(38),
}

var map_token12 = [16]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16('\n'),
	9:  uint16(38),
	10: uint16('\\'),
	11: uint16(38),
	12: uint16('{'),
	13: uint16(38),
	14: uint16('}'),
	15: uint16(38),
}

var map_token13 = [22]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('('),
	5:  uint16(330),
	6:  uint16('*'),
	7:  uint16(334),
	8:  uint16('.'),
	9:  uint16(169),
	10: uint16('/'),
	11: uint16(325),
	12: uint16(':'),
	13: uint16(320),
	14: uint16('<'),
	15: uint16(342),
	16: uint16('\\'),
	17: uint16(398),
	18: uint16('{'),
	19: uint16(398),
	20: uint16('}'),
	21: uint16(398),
}

var map_token14 = [22]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('('),
	5:  uint16(330),
	6:  uint16('/'),
	7:  uint16(326),
	8:  uint16(':'),
	9:  uint16(320),
	10: uint16('.'),
	11: uint16(342),
	12: uint16('<'),
	13: uint16(342),
	14: uint16('*'),
	15: uint16(398),
	16: uint16('\\'),
	17: uint16(398),
	18: uint16('{'),
	19: uint16(398),
	20: uint16('}'),
	21: uint16(398),
}

var map_token15 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16(':'),
	11: uint16(322),
	12: uint16('<'),
	13: uint16(342),
	14: uint16('\\'),
	15: uint16(398),
	16: uint16('{'),
	17: uint16(398),
	18: uint16('}'),
	19: uint16(398),
}

var map_token16 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16('<'),
	11: uint16(342),
	12: uint16('~'),
	13: uint16(324),
	14: uint16('\\'),
	15: uint16(398),
	16: uint16('{'),
	17: uint16(398),
	18: uint16('}'),
	19: uint16(398),
}

var map_token17 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16('<'),
	11: uint16(342),
	12: uint16('~'),
	13: uint16(323),
	14: uint16('\\'),
	15: uint16(398),
	16: uint16('{'),
	17: uint16(398),
	18: uint16('}'),
	19: uint16(398),
}

var map_token18 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16('<'),
	11: uint16(342),
	12: uint16('\\'),
	13: uint16(398),
	14: uint16('{'),
	15: uint16(398),
	16: uint16('}'),
	17: uint16(398),
}

var map_token19 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16('<'),
	11: uint16(342),
	12: uint16('\\'),
	13: uint16(398),
	14: uint16('{'),
	15: uint16(398),
	16: uint16('}'),
	17: uint16(398),
}

var map_token20 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('*'),
	5:  uint16(334),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(325),
	10: uint16('<'),
	11: uint16(342),
	12: uint16('\\'),
	13: uint16(398),
	14: uint16('{'),
	15: uint16(398),
	16: uint16('}'),
	17: uint16(398),
}

var map_token21 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(326),
	8:  uint16('<'),
	9:  uint16(342),
	10: uint16('*'),
	11: uint16(398),
	12: uint16('\\'),
	13: uint16(398),
	14: uint16('{'),
	15: uint16(398),
	16: uint16('}'),
	17: uint16(398),
}

var map_token22 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('/'),
	5:  uint16(326),
	6:  uint16(':'),
	7:  uint16(321),
	8:  uint16('.'),
	9:  uint16(342),
	10: uint16('<'),
	11: uint16(342),
	12: uint16('*'),
	13: uint16(398),
	14: uint16('\\'),
	15: uint16(398),
	16: uint16('{'),
	17: uint16(398),
	18: uint16('}'),
	19: uint16(398),
}

var map_token23 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('/'),
	5:  uint16(326),
	6:  uint16('.'),
	7:  uint16(342),
	8:  uint16('<'),
	9:  uint16(342),
	10: uint16('*'),
	11: uint16(398),
	12: uint16('\\'),
	13: uint16(398),
	14: uint16('{'),
	15: uint16(398),
	16: uint16('}'),
	17: uint16(398),
}

var map_token24 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(397),
	4:  uint16('/'),
	5:  uint16(326),
	6:  uint16('.'),
	7:  uint16(342),
	8:  uint16('<'),
	9:  uint16(342),
	10: uint16('*'),
	11: uint16(398),
	12: uint16('\\'),
	13: uint16(398),
	14: uint16('{'),
	15: uint16(398),
	16: uint16('}'),
	17: uint16(398),
}

var map_token25 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(395),
	4:  uint16(')'),
	5:  uint16(304),
	6:  uint16('*'),
	7:  uint16(332),
	8:  uint16('.'),
	9:  uint16(170),
	10: uint16('/'),
	11: uint16(330),
	12: uint16('<'),
	13: uint16(363),
	14: uint16('\\'),
	15: uint16(396),
	16: uint16('{'),
	17: uint16(396),
	18: uint16('}'),
	19: uint16(396),
}

var map_token26 = [20]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('!'),
	3:  uint16(5),
	4:  uint16(')'),
	5:  uint16(308),
	6:  uint16('*'),
	7:  uint16(333),
	8:  uint16('.'),
	9:  uint16(171),
	10: uint16('/'),
	11: uint16(331),
	12: uint16('<'),
	13: uint16(376),
	14: uint16('\\'),
	15: uint16(6),
	16: uint16('{'),
	17: uint16(6),
	18: uint16('}'),
	19: uint16(6),
}

var map_token27 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16(')'),
	3:  uint16(305),
	4:  uint16('*'),
	5:  uint16(332),
	6:  uint16('.'),
	7:  uint16(172),
	8:  uint16('<'),
	9:  uint16(401),
	10: uint16('/'),
	11: uint16(396),
	12: uint16('\\'),
	13: uint16(396),
	14: uint16('{'),
	15: uint16(396),
	16: uint16('}'),
	17: uint16(396),
}

var map_token28 = [18]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16(')'),
	3:  uint16(309),
	4:  uint16('*'),
	5:  uint16(333),
	6:  uint16('.'),
	7:  uint16(173),
	8:  uint16('<'),
	9:  uint16(34),
	10: uint16('/'),
	11: uint16(6),
	12: uint16('\\'),
	13: uint16(6),
	14: uint16('{'),
	15: uint16(6),
	16: uint16('}'),
	17: uint16(6),
}

var map_token29 = [16]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16('*'),
	3:  uint16(334),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('<'),
	7:  uint16(403),
	8:  uint16('/'),
	9:  uint16(398),
	10: uint16('\\'),
	11: uint16(398),
	12: uint16('{'),
	13: uint16(398),
	14: uint16('}'),
	15: uint16(398),
}

var map_token30 = [22]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('('),
	3:  uint16(331),
	4:  uint16('*'),
	5:  uint16(380),
	6:  uint16('.'),
	7:  uint16(169),
	8:  uint16('/'),
	9:  uint16(371),
	10: uint16(':'),
	11: uint16(366),
	12: uint16('<'),
	13: uint16(355),
	14: uint16('\n'),
	15: uint16(38),
	16: uint16('\\'),
	17: uint16(38),
	18: uint16('{'),
	19: uint16(38),
	20: uint16('}'),
	21: uint16(38),
}

var map_token31 = [22]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('('),
	3:  uint16(331),
	4:  uint16('/'),
	5:  uint16(372),
	6:  uint16(':'),
	7:  uint16(366),
	8:  uint16('.'),
	9:  uint16(355),
	10: uint16('<'),
	11: uint16(355),
	12: uint16('\n'),
	13: uint16(38),
	14: uint16('*'),
	15: uint16(38),
	16: uint16('\\'),
	17: uint16(38),
	18: uint16('{'),
	19: uint16(38),
	20: uint16('}'),
	21: uint16(38),
}

var map_token32 = [20]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16(':'),
	9:  uint16(368),
	10: uint16('<'),
	11: uint16(355),
	12: uint16('\n'),
	13: uint16(38),
	14: uint16('\\'),
	15: uint16(38),
	16: uint16('{'),
	17: uint16(38),
	18: uint16('}'),
	19: uint16(38),
}

var map_token33 = [20]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16('<'),
	9:  uint16(355),
	10: uint16('~'),
	11: uint16(370),
	12: uint16('\n'),
	13: uint16(38),
	14: uint16('\\'),
	15: uint16(38),
	16: uint16('{'),
	17: uint16(38),
	18: uint16('}'),
	19: uint16(38),
}

var map_token34 = [20]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16('<'),
	9:  uint16(355),
	10: uint16('~'),
	11: uint16(369),
	12: uint16('\n'),
	13: uint16(38),
	14: uint16('\\'),
	15: uint16(38),
	16: uint16('{'),
	17: uint16(38),
	18: uint16('}'),
	19: uint16(38),
}

var map_token35 = [18]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16('<'),
	9:  uint16(355),
	10: uint16('\n'),
	11: uint16(38),
	12: uint16('\\'),
	13: uint16(38),
	14: uint16('{'),
	15: uint16(38),
	16: uint16('}'),
	17: uint16(38),
}

var map_token36 = [18]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16('<'),
	9:  uint16(355),
	10: uint16('\n'),
	11: uint16(38),
	12: uint16('\\'),
	13: uint16(38),
	14: uint16('{'),
	15: uint16(38),
	16: uint16('}'),
	17: uint16(38),
}

var map_token37 = [18]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('*'),
	3:  uint16(380),
	4:  uint16('.'),
	5:  uint16(169),
	6:  uint16('/'),
	7:  uint16(371),
	8:  uint16('<'),
	9:  uint16(355),
	10: uint16('\n'),
	11: uint16(38),
	12: uint16('\\'),
	13: uint16(38),
	14: uint16('{'),
	15: uint16(38),
	16: uint16('}'),
	17: uint16(38),
}

var map_token38 = [18]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('.'),
	3:  uint16(169),
	4:  uint16('/'),
	5:  uint16(372),
	6:  uint16('<'),
	7:  uint16(355),
	8:  uint16('\n'),
	9:  uint16(38),
	10: uint16('*'),
	11: uint16(38),
	12: uint16('\\'),
	13: uint16(38),
	14: uint16('{'),
	15: uint16(38),
	16: uint16('}'),
	17: uint16(38),
}

var map_token39 = [20]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('/'),
	3:  uint16(372),
	4:  uint16(':'),
	5:  uint16(367),
	6:  uint16('.'),
	7:  uint16(355),
	8:  uint16('<'),
	9:  uint16(355),
	10: uint16('\n'),
	11: uint16(38),
	12: uint16('*'),
	13: uint16(38),
	14: uint16('\\'),
	15: uint16(38),
	16: uint16('{'),
	17: uint16(38),
	18: uint16('}'),
	19: uint16(38),
}

var map_token40 = [18]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('/'),
	3:  uint16(372),
	4:  uint16('.'),
	5:  uint16(355),
	6:  uint16('<'),
	7:  uint16(355),
	8:  uint16('\n'),
	9:  uint16(38),
	10: uint16('*'),
	11: uint16(38),
	12: uint16('\\'),
	13: uint16(38),
	14: uint16('{'),
	15: uint16(38),
	16: uint16('}'),
	17: uint16(38),
}

var map_token41 = [18]uint16_t{
	0:  uint16('!'),
	1:  uint16(36),
	2:  uint16('/'),
	3:  uint16(372),
	4:  uint16('.'),
	5:  uint16(355),
	6:  uint16('<'),
	7:  uint16(355),
	8:  uint16('\n'),
	9:  uint16(38),
	10: uint16('*'),
	11: uint16(38),
	12: uint16('\\'),
	13: uint16(38),
	14: uint16('{'),
	15: uint16(38),
	16: uint16('}'),
	17: uint16(38),
}

var map_token42 = [16]uint16_t{
	0:  uint16('*'),
	1:  uint16(380),
	2:  uint16('.'),
	3:  uint16(169),
	4:  uint16('<'),
	5:  uint16(37),
	6:  uint16('\n'),
	7:  uint16(38),
	8:  uint16('/'),
	9:  uint16(38),
	10: uint16('\\'),
	11: uint16(38),
	12: uint16('{'),
	13: uint16(38),
	14: uint16('}'),
	15: uint16(38),
}

var map_token43 = [16]uint16_t{
	0:  uint16('\n'),
	1:  uint16(38),
	2:  uint16(')'),
	3:  uint16(305),
	4:  uint16('*'),
	5:  uint16(332),
	6:  uint16('.'),
	7:  uint16(172),
	8:  uint16('<'),
	9:  uint16(401),
	10: uint16('\\'),
	11: uint16(396),
	12: uint16('{'),
	13: uint16(396),
	14: uint16('}'),
	15: uint16(396),
}

var ts_lex_modes = [206]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state: uint16(142),
	},
	2: {
		Flex_state:          uint16(3),
		Fexternal_lex_state: uint16(2),
	},
	3: {
		Flex_state:          uint16(24),
		Fexternal_lex_state: uint16(2),
	},
	4: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	5: {
		Flex_state: uint16(145),
	},
	6: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state: uint16(147),
	},
	9: {
		Flex_state:          uint16(31),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	16: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	17: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(2),
	},
	18: {
		Flex_state: uint16(160),
	},
	19: {
		Flex_state: uint16(149),
	},
	20: {
		Flex_state: uint16(163),
	},
	21: {
		Flex_state: uint16(149),
	},
	22: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(2),
	},
	23: {
		Flex_state: uint16(149),
	},
	24: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	25: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	26: {
		Flex_state: uint16(149),
	},
	27: {
		Flex_state:          uint16(10),
		Fexternal_lex_state: uint16(2),
	},
	28: {
		Flex_state:          uint16(28),
		Fexternal_lex_state: uint16(2),
	},
	29: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	30: {
		Flex_state:          uint16(13),
		Fexternal_lex_state: uint16(2),
	},
	31: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(3),
	},
	32: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	33: {
		Flex_state:          uint16(7),
		Fexternal_lex_state: uint16(2),
	},
	34: {
		Flex_state:          uint16(28),
		Fexternal_lex_state: uint16(2),
	},
	35: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(3),
	},
	36: {
		Flex_state: uint16(160),
	},
	37: {
		Flex_state: uint16(160),
	},
	38: {
		Flex_state: uint16(160),
	},
	39: {
		Flex_state:          uint16(28),
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	41: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state:          uint16(8),
		Fexternal_lex_state: uint16(2),
	},
	44: {
		Flex_state:          uint16(28),
		Fexternal_lex_state: uint16(2),
	},
	45: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	46: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	47: {
		Flex_state:          uint16(26),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Flex_state:          uint16(26),
		Fexternal_lex_state: uint16(2),
	},
	49: {
		Flex_state:          uint16(26),
		Fexternal_lex_state: uint16(2),
	},
	50: {
		Flex_state: uint16(160),
	},
	51: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	52: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	53: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	54: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	56: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	57: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	58: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Flex_state: uint16(160),
	},
	60: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	61: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	62: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	63: {
		Flex_state:          uint16(29),
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Flex_state: uint16(153),
	},
	65: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(2),
	},
	66: {
		Flex_state:          uint16(26),
		Fexternal_lex_state: uint16(2),
	},
	67: {
		Flex_state: uint16(150),
	},
	68: {
		Flex_state:          uint16(16),
		Fexternal_lex_state: uint16(2),
	},
	69: {
		Flex_state: uint16(153),
	},
	70: {
		Flex_state: uint16(156),
	},
	71: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	72: {
		Flex_state: uint16(150),
	},
	73: {
		Flex_state: uint16(161),
	},
	74: {
		Flex_state: uint16(161),
	},
	75: {
		Flex_state:          uint16(26),
		Fexternal_lex_state: uint16(2),
	},
	76: {
		Flex_state: uint16(150),
	},
	77: {
		Flex_state: uint16(156),
	},
	78: {
		Flex_state: uint16(150),
	},
	79: {
		Flex_state: uint16(149),
	},
	80: {
		Flex_state:          uint16(160),
		Fexternal_lex_state: uint16(4),
	},
	81: {
		Flex_state:          uint16(160),
		Fexternal_lex_state: uint16(4),
	},
	82: {
		Flex_state: uint16(149),
	},
	83: {
		Flex_state: uint16(161),
	},
	84: {
		Flex_state: uint16(161),
	},
	85: {
		Flex_state: uint16(149),
	},
	86: {
		Flex_state: uint16(149),
	},
	87: {
		Flex_state: uint16(142),
	},
	88: {
		Flex_state: uint16(158),
	},
	89: {
		Flex_state: uint16(160),
	},
	90: {
		Flex_state: uint16(160),
	},
	91: {
		Flex_state: uint16(160),
	},
	92: {
		Flex_state: uint16(158),
	},
	93: {
		Flex_state: uint16(160),
	},
	94: {
		Flex_state: uint16(160),
	},
	95: {
		Flex_state: uint16(158),
	},
	96: {
		Flex_state: uint16(160),
	},
	97: {
		Flex_state: uint16(160),
	},
	98: {
		Flex_state: uint16(160),
	},
	99: {
		Flex_state: uint16(160),
	},
	100: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	101: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	102: {
		Flex_state: uint16(142),
	},
	103: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	104: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	105: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	106: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	107: {
		Flex_state: uint16(158),
	},
	108: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	109: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	110: {
		Flex_state: uint16(142),
	},
	111: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	112: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	113: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	114: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	115: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	116: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	117: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	118: {
		Flex_state:          uint16(18),
		Fexternal_lex_state: uint16(2),
	},
	119: {
		Flex_state: uint16(142),
	},
	120: {
		Flex_state: uint16(142),
	},
	121: {
		Flex_state: uint16(158),
	},
	122: {
		Flex_state: uint16(142),
	},
	123: {
		Flex_state: uint16(142),
	},
	124: {
		Flex_state: uint16(142),
	},
	125: {
		Flex_state: uint16(142),
	},
	126: {
		Flex_state: uint16(142),
	},
	127: {
		Flex_state: uint16(142),
	},
	128: {
		Flex_state: uint16(142),
	},
	129: {
		Flex_state: uint16(142),
	},
	130: {
		Flex_state: uint16(21),
	},
	131: {
		Flex_state: uint16(142),
	},
	132: {
		Flex_state: uint16(142),
	},
	133: {
		Flex_state: uint16(142),
	},
	134: {
		Flex_state: uint16(142),
	},
	135: {
		Flex_state: uint16(21),
	},
	136: {
		Flex_state: uint16(142),
	},
	137: {
		Flex_state: uint16(142),
	},
	138: {
		Flex_state: uint16(142),
	},
	139: {
		Flex_state: uint16(142),
	},
	140: {
		Flex_state: uint16(21),
	},
	141: {
		Flex_state: uint16(21),
	},
	142: {
		Flex_state: uint16(21),
	},
	143: {
		Flex_state: uint16(21),
	},
	144: {
		Flex_state:          uint16(21),
		Fexternal_lex_state: uint16(4),
	},
	145: {
		Flex_state:          uint16(21),
		Fexternal_lex_state: uint16(4),
	},
	146: {
		Flex_state: uint16(24),
	},
	147: {
		Flex_state: uint16(23),
	},
	148: {
		Flex_state: uint16(23),
	},
	149: {
		Flex_state: uint16(23),
	},
	150: {
		Flex_state: uint16(24),
	},
	151: {
		Flex_state: uint16(24),
	},
	152: {
		Flex_state: uint16(24),
	},
	153: {
		Flex_state: uint16(21),
	},
	154: {
		Flex_state: uint16(21),
	},
	155: {
		Flex_state: uint16(21),
	},
	156: {
		Flex_state:          uint16(142),
		Fexternal_lex_state: uint16(5),
	},
	157: {
		Flex_state: uint16(21),
	},
	158: {
		Flex_state: uint16(21),
	},
	159: {
		Flex_state: uint16(21),
	},
	160: {
		Flex_state: uint16(21),
	},
	161: {
		Flex_state: uint16(21),
	},
	162: {
		Flex_state: uint16(21),
	},
	163: {
		Fexternal_lex_state: uint16(5),
	},
	164: {
		Flex_state: uint16(142),
	},
	165: {
		Flex_state: uint16(21),
	},
	166: {
		Flex_state: uint16(179),
	},
	167: {
		Flex_state: uint16(142),
	},
	168: {},
	169: {},
	170: {},
	171: {
		Flex_state: uint16(142),
	},
	172: {},
	173: {},
	174: {
		Flex_state: uint16(142),
	},
	175: {},
	176: {
		Flex_state: uint16(142),
	},
	177: {
		Flex_state: uint16(179),
	},
	178: {
		Flex_state: uint16(21),
	},
	179: {
		Flex_state: uint16(21),
	},
	180: {},
	181: {
		Flex_state: uint16(142),
	},
	182: {
		Flex_state: uint16(295),
	},
	183: {
		Flex_state: uint16(142),
	},
	184: {
		Flex_state: uint16(21),
	},
	185: {
		Fexternal_lex_state: uint16(5),
	},
	186: {
		Flex_state: uint16(142),
	},
	187: {
		Flex_state: uint16(142),
	},
	188: {
		Flex_state: uint16(142),
	},
	189: {
		Flex_state: uint16(142),
	},
	190: {
		Flex_state: uint16(142),
	},
	191: {
		Flex_state: uint16(142),
	},
	192: {
		Flex_state: uint16(299),
	},
	193: {
		Flex_state: uint16(142),
	},
	194: {
		Flex_state: uint16(142),
	},
	195: {
		Fexternal_lex_state: uint16(6),
	},
	196: {},
	197: {
		Flex_state: uint16(299),
	},
	198: {
		Fexternal_lex_state: uint16(7),
	},
	199: {},
	200: {},
	201: {
		Fexternal_lex_state: uint16(6),
	},
	202: {
		Flex_state: uint16(142),
	},
	203: {},
	204: {
		Flex_state: uint16(142),
	},
	205: {
		Flex_state: uint16(295),
	},
}

var ts_parse_table = [2][69]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		4:  uint16(1),
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
		19: uint16(1),
		20: uint16(1),
		21: uint16(1),
		22: uint16(1),
		23: uint16(1),
		24: uint16(1),
		25: uint16(1),
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
	},
	1: {
		37: uint16(3),
		38: uint16(5),
		47: uint16(180),
	},
}

var ts_small_parse_table = [4187]uint16_t{
	0:    uint16(23),
	1:    uint16(9),
	2:    uint16(1),
	3:    uint16(aux_sym_brief_header_token1),
	4:    uint16(11),
	5:    uint16(1),
	6:    uint16(sym_tag_name_with_argument),
	7:    uint16(13),
	8:    uint16(1),
	9:    uint16(sym_tag_name_with_multiple_arguments),
	10:   uint16(15),
	11:   uint16(1),
	12:   uint16(sym_tag_name_with_types),
	13:   uint16(17),
	14:   uint16(1),
	15:   uint16(sym_tag_name_with_self_types),
	16:   uint16(19),
	17:   uint16(1),
	18:   uint16(sym_tag_name_with_type),
	19:   uint16(21),
	20:   uint16(1),
	21:   uint16(sym_tag_name),
	22:   uint16(23),
	23:   uint16(1),
	24:   uint16(anon_sym_LBRACK),
	25:   uint16(25),
	26:   uint16(1),
	27:   uint16(anon_sym_BSLASHa),
	28:   uint16(27),
	29:   uint16(1),
	30:   uint16(anon_sym_BSLASHc),
	31:   uint16(29),
	32:   uint16(1),
	33:   uint16(anon_sym_LTa),
	34:   uint16(33),
	35:   uint16(1),
	36:   uint16(anon_sym_ATcode),
	37:   uint16(37),
	38:   uint16(1),
	39:   uint16(sym__text_line),
	40:   uint16(39),
	41:   uint16(1),
	42:   uint16(sym_code_block_start),
	43:   uint16(4),
	44:   uint16(1),
	45:   uint16(sym_brief_header),
	46:   uint16(12),
	47:   uint16(1),
	48:   uint16(sym_storageclass),
	49:   uint16(51),
	50:   uint16(1),
	51:   uint16(sym_description),
	52:   uint16(173),
	53:   uint16(1),
	54:   uint16(sym__multiline_end),
	55:   uint16(7),
	56:   uint16(2),
	57:   uint16(anon_sym_ATbrief),
	58:   uint16(anon_sym_BSLASHbrief),
	59:   uint16(31),
	60:   uint16(2),
	61:   uint16(sym_function_link),
	62:   uint16(sym__text),
	63:   uint16(35),
	64:   uint16(2),
	65:   uint16(anon_sym_SLASH),
	66:   uint16(anon_sym_STAR_SLASH),
	67:   uint16(52),
	68:   uint16(3),
	69:   uint16(sym_tag),
	70:   uint16(sym_code_block),
	71:   uint16(aux_sym_document_repeat1),
	72:   uint16(15),
	73:   uint16(4),
	74:   uint16(sym_emphasis),
	75:   uint16(sym_code_word),
	76:   uint16(sym_link),
	77:   uint16(aux_sym_description_repeat1),
	78:   uint16(13),
	79:   uint16(25),
	80:   uint16(1),
	81:   uint16(anon_sym_BSLASHa),
	82:   uint16(27),
	83:   uint16(1),
	84:   uint16(anon_sym_BSLASHc),
	85:   uint16(29),
	86:   uint16(1),
	87:   uint16(anon_sym_LTa),
	88:   uint16(31),
	89:   uint16(1),
	90:   uint16(sym__text),
	91:   uint16(43),
	92:   uint16(1),
	93:   uint16(aux_sym_identifier_token1),
	94:   uint16(45),
	95:   uint16(1),
	96:   uint16(anon_sym_TILDE),
	97:   uint16(49),
	98:   uint16(1),
	99:   uint16(sym_function_link),
	100:  uint16(22),
	101:  uint16(1),
	102:  uint16(sym_identifier),
	103:  uint16(111),
	104:  uint16(1),
	105:  uint16(sym_description),
	106:  uint16(47),
	107:  uint16(3),
	108:  uint16(sym_code_block_start),
	109:  uint16(anon_sym_LBRACK),
	110:  uint16(anon_sym_SLASH),
	111:  uint16(13),
	112:  uint16(3),
	113:  uint16(sym__expression),
	114:  uint16(sym_qualified_identifier),
	115:  uint16(sym_function),
	116:  uint16(15),
	117:  uint16(4),
	118:  uint16(sym_emphasis),
	119:  uint16(sym_code_word),
	120:  uint16(sym_link),
	121:  uint16(aux_sym_description_repeat1),
	122:  uint16(41),
	123:  uint16(9),
	124:  uint16(sym_tag_name_with_argument),
	125:  uint16(sym_tag_name_with_multiple_arguments),
	126:  uint16(sym_tag_name_with_types),
	127:  uint16(sym_tag_name_with_self_types),
	128:  uint16(sym_tag_name_with_type),
	129:  uint16(sym_tag_name),
	130:  uint16(anon_sym_ATcode),
	131:  uint16(anon_sym_STAR_SLASH),
	132:  uint16(sym__text_line),
	133:  uint16(22),
	134:  uint16(11),
	135:  uint16(1),
	136:  uint16(sym_tag_name_with_argument),
	137:  uint16(13),
	138:  uint16(1),
	139:  uint16(sym_tag_name_with_multiple_arguments),
	140:  uint16(15),
	141:  uint16(1),
	142:  uint16(sym_tag_name_with_types),
	143:  uint16(17),
	144:  uint16(1),
	145:  uint16(sym_tag_name_with_self_types),
	146:  uint16(19),
	147:  uint16(1),
	148:  uint16(sym_tag_name_with_type),
	149:  uint16(21),
	150:  uint16(1),
	151:  uint16(sym_tag_name),
	152:  uint16(23),
	153:  uint16(1),
	154:  uint16(anon_sym_LBRACK),
	155:  uint16(25),
	156:  uint16(1),
	157:  uint16(anon_sym_BSLASHa),
	158:  uint16(27),
	159:  uint16(1),
	160:  uint16(anon_sym_BSLASHc),
	161:  uint16(29),
	162:  uint16(1),
	163:  uint16(anon_sym_LTa),
	164:  uint16(31),
	165:  uint16(1),
	166:  uint16(sym__text),
	167:  uint16(33),
	168:  uint16(1),
	169:  uint16(anon_sym_ATcode),
	170:  uint16(39),
	171:  uint16(1),
	172:  uint16(sym_code_block_start),
	173:  uint16(49),
	174:  uint16(1),
	175:  uint16(sym_function_link),
	176:  uint16(51),
	177:  uint16(1),
	178:  uint16(anon_sym_SLASH),
	179:  uint16(53),
	180:  uint16(1),
	181:  uint16(anon_sym_STAR_SLASH),
	182:  uint16(55),
	183:  uint16(1),
	184:  uint16(sym__text_line),
	185:  uint16(12),
	186:  uint16(1),
	187:  uint16(sym_storageclass),
	188:  uint16(54),
	189:  uint16(1),
	190:  uint16(sym_description),
	191:  uint16(170),
	192:  uint16(1),
	193:  uint16(sym__multiline_end),
	194:  uint16(45),
	195:  uint16(3),
	196:  uint16(sym_tag),
	197:  uint16(sym_code_block),
	198:  uint16(aux_sym_document_repeat1),
	199:  uint16(15),
	200:  uint16(4),
	201:  uint16(sym_emphasis),
	202:  uint16(sym_code_word),
	203:  uint16(sym_link),
	204:  uint16(aux_sym_description_repeat1),
	205:  uint16(19),
	206:  uint16(57),
	207:  uint16(1),
	209:  uint16(61),
	210:  uint16(1),
	211:  uint16(aux_sym_brief_header_token1),
	212:  uint16(63),
	213:  uint16(1),
	214:  uint16(sym_tag_name_with_argument),
	215:  uint16(65),
	216:  uint16(1),
	217:  uint16(sym_tag_name_with_multiple_arguments),
	218:  uint16(67),
	219:  uint16(1),
	220:  uint16(sym_tag_name_with_types),
	221:  uint16(69),
	222:  uint16(1),
	223:  uint16(sym_tag_name_with_self_types),
	224:  uint16(71),
	225:  uint16(1),
	226:  uint16(sym_tag_name_with_type),
	227:  uint16(73),
	228:  uint16(1),
	229:  uint16(sym_tag_name),
	230:  uint16(75),
	231:  uint16(1),
	232:  uint16(anon_sym_LBRACK),
	233:  uint16(77),
	234:  uint16(1),
	235:  uint16(anon_sym_BSLASHa),
	236:  uint16(79),
	237:  uint16(1),
	238:  uint16(anon_sym_BSLASHc),
	239:  uint16(81),
	240:  uint16(1),
	241:  uint16(anon_sym_LTa),
	242:  uint16(18),
	243:  uint16(1),
	244:  uint16(sym_brief_header),
	245:  uint16(37),
	246:  uint16(1),
	247:  uint16(sym_storageclass),
	248:  uint16(119),
	249:  uint16(1),
	250:  uint16(sym_description),
	251:  uint16(59),
	252:  uint16(2),
	253:  uint16(anon_sym_ATbrief),
	254:  uint16(anon_sym_BSLASHbrief),
	255:  uint16(83),
	256:  uint16(2),
	257:  uint16(sym_function_link),
	258:  uint16(sym__text),
	259:  uint16(120),
	260:  uint16(2),
	261:  uint16(sym_tag),
	262:  uint16(aux_sym_document_repeat2),
	263:  uint16(59),
	264:  uint16(4),
	265:  uint16(sym_emphasis),
	266:  uint16(sym_code_word),
	267:  uint16(sym_link),
	268:  uint16(aux_sym_description_repeat1),
	269:  uint16(11),
	270:  uint16(25),
	271:  uint16(1),
	272:  uint16(anon_sym_BSLASHa),
	273:  uint16(27),
	274:  uint16(1),
	275:  uint16(anon_sym_BSLASHc),
	276:  uint16(29),
	277:  uint16(1),
	278:  uint16(anon_sym_LTa),
	279:  uint16(31),
	280:  uint16(1),
	281:  uint16(sym__text),
	282:  uint16(49),
	283:  uint16(1),
	284:  uint16(sym_function_link),
	285:  uint16(85),
	286:  uint16(1),
	287:  uint16(anon_sym_COMMA),
	288:  uint16(10),
	289:  uint16(1),
	290:  uint16(aux_sym_tag_repeat1),
	291:  uint16(101),
	292:  uint16(1),
	293:  uint16(sym_description),
	294:  uint16(89),
	295:  uint16(3),
	296:  uint16(sym_code_block_start),
	297:  uint16(anon_sym_LBRACK),
	298:  uint16(anon_sym_SLASH),
	299:  uint16(15),
	300:  uint16(4),
	301:  uint16(sym_emphasis),
	302:  uint16(sym_code_word),
	303:  uint16(sym_link),
	304:  uint16(aux_sym_description_repeat1),
	305:  uint16(87),
	306:  uint16(9),
	307:  uint16(sym_tag_name_with_argument),
	308:  uint16(sym_tag_name_with_multiple_arguments),
	309:  uint16(sym_tag_name_with_types),
	310:  uint16(sym_tag_name_with_self_types),
	311:  uint16(sym_tag_name_with_type),
	312:  uint16(sym_tag_name),
	313:  uint16(anon_sym_ATcode),
	314:  uint16(anon_sym_STAR_SLASH),
	315:  uint16(sym__text_line),
	316:  uint16(11),
	317:  uint16(25),
	318:  uint16(1),
	319:  uint16(anon_sym_BSLASHa),
	320:  uint16(27),
	321:  uint16(1),
	322:  uint16(anon_sym_BSLASHc),
	323:  uint16(29),
	324:  uint16(1),
	325:  uint16(anon_sym_LTa),
	326:  uint16(31),
	327:  uint16(1),
	328:  uint16(sym__text),
	329:  uint16(49),
	330:  uint16(1),
	331:  uint16(sym_function_link),
	332:  uint16(85),
	333:  uint16(1),
	334:  uint16(anon_sym_COMMA),
	335:  uint16(29),
	336:  uint16(1),
	337:  uint16(aux_sym_tag_repeat1),
	338:  uint16(103),
	339:  uint16(1),
	340:  uint16(sym_description),
	341:  uint16(93),
	342:  uint16(3),
	343:  uint16(sym_code_block_start),
	344:  uint16(anon_sym_LBRACK),
	345:  uint16(anon_sym_SLASH),
	346:  uint16(15),
	347:  uint16(4),
	348:  uint16(sym_emphasis),
	349:  uint16(sym_code_word),
	350:  uint16(sym_link),
	351:  uint16(aux_sym_description_repeat1),
	352:  uint16(91),
	353:  uint16(9),
	354:  uint16(sym_tag_name_with_argument),
	355:  uint16(sym_tag_name_with_multiple_arguments),
	356:  uint16(sym_tag_name_with_types),
	357:  uint16(sym_tag_name_with_self_types),
	358:  uint16(sym_tag_name_with_type),
	359:  uint16(sym_tag_name),
	360:  uint16(anon_sym_ATcode),
	361:  uint16(anon_sym_STAR_SLASH),
	362:  uint16(sym__text_line),
	363:  uint16(13),
	364:  uint16(77),
	365:  uint16(1),
	366:  uint16(anon_sym_BSLASHa),
	367:  uint16(79),
	368:  uint16(1),
	369:  uint16(anon_sym_BSLASHc),
	370:  uint16(81),
	371:  uint16(1),
	372:  uint16(anon_sym_LTa),
	373:  uint16(83),
	374:  uint16(1),
	375:  uint16(sym__text),
	376:  uint16(95),
	377:  uint16(1),
	378:  uint16(aux_sym_identifier_token1),
	379:  uint16(97),
	380:  uint16(1),
	381:  uint16(anon_sym_TILDE),
	382:  uint16(99),
	383:  uint16(1),
	384:  uint16(sym_function_link),
	385:  uint16(70),
	386:  uint16(1),
	387:  uint16(sym_identifier),
	388:  uint16(128),
	389:  uint16(1),
	390:  uint16(sym_description),
	391:  uint16(47),
	392:  uint16(2),
	394:  uint16(anon_sym_LBRACK),
	395:  uint16(36),
	396:  uint16(3),
	397:  uint16(sym__expression),
	398:  uint16(sym_qualified_identifier),
	399:  uint16(sym_function),
	400:  uint16(59),
	401:  uint16(4),
	402:  uint16(sym_emphasis),
	403:  uint16(sym_code_word),
	404:  uint16(sym_link),
	405:  uint16(aux_sym_description_repeat1),
	406:  uint16(41),
	407:  uint16(6),
	408:  uint16(sym_tag_name_with_argument),
	409:  uint16(sym_tag_name_with_multiple_arguments),
	410:  uint16(sym_tag_name_with_types),
	411:  uint16(sym_tag_name_with_self_types),
	412:  uint16(sym_tag_name_with_type),
	413:  uint16(sym_tag_name),
	414:  uint16(11),
	415:  uint16(25),
	416:  uint16(1),
	417:  uint16(anon_sym_BSLASHa),
	418:  uint16(27),
	419:  uint16(1),
	420:  uint16(anon_sym_BSLASHc),
	421:  uint16(29),
	422:  uint16(1),
	423:  uint16(anon_sym_LTa),
	424:  uint16(31),
	425:  uint16(1),
	426:  uint16(sym__text),
	427:  uint16(49),
	428:  uint16(1),
	429:  uint16(sym_function_link),
	430:  uint16(101),
	431:  uint16(1),
	432:  uint16(aux_sym_identifier_token1),
	433:  uint16(14),
	434:  uint16(1),
	435:  uint16(sym_identifier),
	436:  uint16(111),
	437:  uint16(1),
	438:  uint16(sym_description),
	439:  uint16(47),
	440:  uint16(3),
	441:  uint16(sym_code_block_start),
	442:  uint16(anon_sym_LBRACK),
	443:  uint16(anon_sym_SLASH),
	444:  uint16(15),
	445:  uint16(4),
	446:  uint16(sym_emphasis),
	447:  uint16(sym_code_word),
	448:  uint16(sym_link),
	449:  uint16(aux_sym_description_repeat1),
	450:  uint16(41),
	451:  uint16(9),
	452:  uint16(sym_tag_name_with_argument),
	453:  uint16(sym_tag_name_with_multiple_arguments),
	454:  uint16(sym_tag_name_with_types),
	455:  uint16(sym_tag_name_with_self_types),
	456:  uint16(sym_tag_name_with_type),
	457:  uint16(sym_tag_name),
	458:  uint16(anon_sym_ATcode),
	459:  uint16(anon_sym_STAR_SLASH),
	460:  uint16(sym__text_line),
	461:  uint16(11),
	462:  uint16(25),
	463:  uint16(1),
	464:  uint16(anon_sym_BSLASHa),
	465:  uint16(27),
	466:  uint16(1),
	467:  uint16(anon_sym_BSLASHc),
	468:  uint16(29),
	469:  uint16(1),
	470:  uint16(anon_sym_LTa),
	471:  uint16(31),
	472:  uint16(1),
	473:  uint16(sym__text),
	474:  uint16(49),
	475:  uint16(1),
	476:  uint16(sym_function_link),
	477:  uint16(85),
	478:  uint16(1),
	479:  uint16(anon_sym_COMMA),
	480:  uint16(29),
	481:  uint16(1),
	482:  uint16(aux_sym_tag_repeat1),
	483:  uint16(118),
	484:  uint16(1),
	485:  uint16(sym_description),
	486:  uint16(105),
	487:  uint16(3),
	488:  uint16(sym_code_block_start),
	489:  uint16(anon_sym_LBRACK),
	490:  uint16(anon_sym_SLASH),
	491:  uint16(15),
	492:  uint16(4),
	493:  uint16(sym_emphasis),
	494:  uint16(sym_code_word),
	495:  uint16(sym_link),
	496:  uint16(aux_sym_description_repeat1),
	497:  uint16(103),
	498:  uint16(9),
	499:  uint16(sym_tag_name_with_argument),
	500:  uint16(sym_tag_name_with_multiple_arguments),
	501:  uint16(sym_tag_name_with_types),
	502:  uint16(sym_tag_name_with_self_types),
	503:  uint16(sym_tag_name_with_type),
	504:  uint16(sym_tag_name),
	505:  uint16(anon_sym_ATcode),
	506:  uint16(anon_sym_STAR_SLASH),
	507:  uint16(sym__text_line),
	508:  uint16(11),
	509:  uint16(25),
	510:  uint16(1),
	511:  uint16(anon_sym_BSLASHa),
	512:  uint16(27),
	513:  uint16(1),
	514:  uint16(anon_sym_BSLASHc),
	515:  uint16(29),
	516:  uint16(1),
	517:  uint16(anon_sym_LTa),
	518:  uint16(31),
	519:  uint16(1),
	520:  uint16(sym__text),
	521:  uint16(49),
	522:  uint16(1),
	523:  uint16(sym_function_link),
	524:  uint16(85),
	525:  uint16(1),
	526:  uint16(anon_sym_COMMA),
	527:  uint16(7),
	528:  uint16(1),
	529:  uint16(aux_sym_tag_repeat1),
	530:  uint16(118),
	531:  uint16(1),
	532:  uint16(sym_description),
	533:  uint16(105),
	534:  uint16(3),
	535:  uint16(sym_code_block_start),
	536:  uint16(anon_sym_LBRACK),
	537:  uint16(anon_sym_SLASH),
	538:  uint16(15),
	539:  uint16(4),
	540:  uint16(sym_emphasis),
	541:  uint16(sym_code_word),
	542:  uint16(sym_link),
	543:  uint16(aux_sym_description_repeat1),
	544:  uint16(103),
	545:  uint16(9),
	546:  uint16(sym_tag_name_with_argument),
	547:  uint16(sym_tag_name_with_multiple_arguments),
	548:  uint16(sym_tag_name_with_types),
	549:  uint16(sym_tag_name_with_self_types),
	550:  uint16(sym_tag_name_with_type),
	551:  uint16(sym_tag_name),
	552:  uint16(anon_sym_ATcode),
	553:  uint16(anon_sym_STAR_SLASH),
	554:  uint16(sym__text_line),
	555:  uint16(9),
	556:  uint16(25),
	557:  uint16(1),
	558:  uint16(anon_sym_BSLASHa),
	559:  uint16(27),
	560:  uint16(1),
	561:  uint16(anon_sym_BSLASHc),
	562:  uint16(29),
	563:  uint16(1),
	564:  uint16(anon_sym_LTa),
	565:  uint16(31),
	566:  uint16(1),
	567:  uint16(sym__text),
	568:  uint16(49),
	569:  uint16(1),
	570:  uint16(sym_function_link),
	571:  uint16(111),
	572:  uint16(1),
	573:  uint16(sym_description),
	574:  uint16(47),
	575:  uint16(3),
	576:  uint16(sym_code_block_start),
	577:  uint16(anon_sym_LBRACK),
	578:  uint16(anon_sym_SLASH),
	579:  uint16(15),
	580:  uint16(4),
	581:  uint16(sym_emphasis),
	582:  uint16(sym_code_word),
	583:  uint16(sym_link),
	584:  uint16(aux_sym_description_repeat1),
	585:  uint16(41),
	586:  uint16(9),
	587:  uint16(sym_tag_name_with_argument),
	588:  uint16(sym_tag_name_with_multiple_arguments),
	589:  uint16(sym_tag_name_with_types),
	590:  uint16(sym_tag_name_with_self_types),
	591:  uint16(sym_tag_name_with_type),
	592:  uint16(sym_tag_name),
	593:  uint16(anon_sym_ATcode),
	594:  uint16(anon_sym_STAR_SLASH),
	595:  uint16(sym__text_line),
	596:  uint16(9),
	597:  uint16(25),
	598:  uint16(1),
	599:  uint16(anon_sym_BSLASHa),
	600:  uint16(27),
	601:  uint16(1),
	602:  uint16(anon_sym_BSLASHc),
	603:  uint16(29),
	604:  uint16(1),
	605:  uint16(anon_sym_LTa),
	606:  uint16(31),
	607:  uint16(1),
	608:  uint16(sym__text),
	609:  uint16(49),
	610:  uint16(1),
	611:  uint16(sym_function_link),
	612:  uint16(101),
	613:  uint16(1),
	614:  uint16(sym_description),
	615:  uint16(89),
	616:  uint16(3),
	617:  uint16(sym_code_block_start),
	618:  uint16(anon_sym_LBRACK),
	619:  uint16(anon_sym_SLASH),
	620:  uint16(15),
	621:  uint16(4),
	622:  uint16(sym_emphasis),
	623:  uint16(sym_code_word),
	624:  uint16(sym_link),
	625:  uint16(aux_sym_description_repeat1),
	626:  uint16(87),
	627:  uint16(9),
	628:  uint16(sym_tag_name_with_argument),
	629:  uint16(sym_tag_name_with_multiple_arguments),
	630:  uint16(sym_tag_name_with_types),
	631:  uint16(sym_tag_name_with_self_types),
	632:  uint16(sym_tag_name_with_type),
	633:  uint16(sym_tag_name),
	634:  uint16(anon_sym_ATcode),
	635:  uint16(anon_sym_STAR_SLASH),
	636:  uint16(sym__text_line),
	637:  uint16(9),
	638:  uint16(25),
	639:  uint16(1),
	640:  uint16(anon_sym_BSLASHa),
	641:  uint16(27),
	642:  uint16(1),
	643:  uint16(anon_sym_BSLASHc),
	644:  uint16(29),
	645:  uint16(1),
	646:  uint16(anon_sym_LTa),
	647:  uint16(31),
	648:  uint16(1),
	649:  uint16(sym__text),
	650:  uint16(49),
	651:  uint16(1),
	652:  uint16(sym_function_link),
	653:  uint16(114),
	654:  uint16(1),
	655:  uint16(sym_description),
	656:  uint16(109),
	657:  uint16(3),
	658:  uint16(sym_code_block_start),
	659:  uint16(anon_sym_LBRACK),
	660:  uint16(anon_sym_SLASH),
	661:  uint16(15),
	662:  uint16(4),
	663:  uint16(sym_emphasis),
	664:  uint16(sym_code_word),
	665:  uint16(sym_link),
	666:  uint16(aux_sym_description_repeat1),
	667:  uint16(107),
	668:  uint16(9),
	669:  uint16(sym_tag_name_with_argument),
	670:  uint16(sym_tag_name_with_multiple_arguments),
	671:  uint16(sym_tag_name_with_types),
	672:  uint16(sym_tag_name_with_self_types),
	673:  uint16(sym_tag_name_with_type),
	674:  uint16(sym_tag_name),
	675:  uint16(anon_sym_ATcode),
	676:  uint16(anon_sym_STAR_SLASH),
	677:  uint16(sym__text_line),
	678:  uint16(8),
	679:  uint16(25),
	680:  uint16(1),
	681:  uint16(anon_sym_BSLASHa),
	682:  uint16(27),
	683:  uint16(1),
	684:  uint16(anon_sym_BSLASHc),
	685:  uint16(29),
	686:  uint16(1),
	687:  uint16(anon_sym_LTa),
	688:  uint16(115),
	689:  uint16(1),
	690:  uint16(sym_function_link),
	691:  uint16(117),
	692:  uint16(1),
	693:  uint16(sym__text),
	694:  uint16(113),
	695:  uint16(3),
	696:  uint16(sym_code_block_start),
	697:  uint16(anon_sym_LBRACK),
	698:  uint16(anon_sym_SLASH),
	699:  uint16(16),
	700:  uint16(4),
	701:  uint16(sym_emphasis),
	702:  uint16(sym_code_word),
	703:  uint16(sym_link),
	704:  uint16(aux_sym_description_repeat1),
	705:  uint16(111),
	706:  uint16(9),
	707:  uint16(sym_tag_name_with_argument),
	708:  uint16(sym_tag_name_with_multiple_arguments),
	709:  uint16(sym_tag_name_with_types),
	710:  uint16(sym_tag_name_with_self_types),
	711:  uint16(sym_tag_name_with_type),
	712:  uint16(sym_tag_name),
	713:  uint16(anon_sym_ATcode),
	714:  uint16(anon_sym_STAR_SLASH),
	715:  uint16(sym__text_line),
	716:  uint16(8),
	717:  uint16(123),
	718:  uint16(1),
	719:  uint16(anon_sym_BSLASHa),
	720:  uint16(126),
	721:  uint16(1),
	722:  uint16(anon_sym_BSLASHc),
	723:  uint16(129),
	724:  uint16(1),
	725:  uint16(anon_sym_LTa),
	726:  uint16(132),
	727:  uint16(1),
	728:  uint16(sym_function_link),
	729:  uint16(135),
	730:  uint16(1),
	731:  uint16(sym__text),
	732:  uint16(121),
	733:  uint16(3),
	734:  uint16(sym_code_block_start),
	735:  uint16(anon_sym_LBRACK),
	736:  uint16(anon_sym_SLASH),
	737:  uint16(16),
	738:  uint16(4),
	739:  uint16(sym_emphasis),
	740:  uint16(sym_code_word),
	741:  uint16(sym_link),
	742:  uint16(aux_sym_description_repeat1),
	743:  uint16(119),
	744:  uint16(9),
	745:  uint16(sym_tag_name_with_argument),
	746:  uint16(sym_tag_name_with_multiple_arguments),
	747:  uint16(sym_tag_name_with_types),
	748:  uint16(sym_tag_name_with_self_types),
	749:  uint16(sym_tag_name_with_type),
	750:  uint16(sym_tag_name),
	751:  uint16(anon_sym_ATcode),
	752:  uint16(anon_sym_STAR_SLASH),
	753:  uint16(sym__text_line),
	754:  uint16(5),
	755:  uint16(142),
	756:  uint16(1),
	757:  uint16(anon_sym_COLON_COLON),
	758:  uint16(144),
	759:  uint16(1),
	760:  uint16(anon_sym_LPAREN),
	761:  uint16(24),
	762:  uint16(1),
	763:  uint16(aux_sym_qualified_identifier_repeat1),
	764:  uint16(138),
	765:  uint16(6),
	766:  uint16(sym_code_block_start),
	767:  uint16(anon_sym_COMMA),
	768:  uint16(anon_sym_LBRACK),
	769:  uint16(anon_sym_LTa),
	770:  uint16(sym_function_link),
	771:  uint16(anon_sym_SLASH),
	772:  uint16(140),
	773:  uint16(12),
	774:  uint16(sym_tag_name_with_argument),
	775:  uint16(sym_tag_name_with_multiple_arguments),
	776:  uint16(sym_tag_name_with_types),
	777:  uint16(sym_tag_name_with_self_types),
	778:  uint16(sym_tag_name_with_type),
	779:  uint16(sym_tag_name),
	780:  uint16(anon_sym_BSLASHa),
	781:  uint16(anon_sym_BSLASHc),
	782:  uint16(anon_sym_ATcode),
	783:  uint16(sym__text),
	784:  uint16(anon_sym_STAR_SLASH),
	785:  uint16(sym__text_line),
	786:  uint16(17),
	787:  uint16(63),
	788:  uint16(1),
	789:  uint16(sym_tag_name_with_argument),
	790:  uint16(65),
	791:  uint16(1),
	792:  uint16(sym_tag_name_with_multiple_arguments),
	793:  uint16(67),
	794:  uint16(1),
	795:  uint16(sym_tag_name_with_types),
	796:  uint16(69),
	797:  uint16(1),
	798:  uint16(sym_tag_name_with_self_types),
	799:  uint16(71),
	800:  uint16(1),
	801:  uint16(sym_tag_name_with_type),
	802:  uint16(73),
	803:  uint16(1),
	804:  uint16(sym_tag_name),
	805:  uint16(75),
	806:  uint16(1),
	807:  uint16(anon_sym_LBRACK),
	808:  uint16(77),
	809:  uint16(1),
	810:  uint16(anon_sym_BSLASHa),
	811:  uint16(79),
	812:  uint16(1),
	813:  uint16(anon_sym_BSLASHc),
	814:  uint16(81),
	815:  uint16(1),
	816:  uint16(anon_sym_LTa),
	817:  uint16(83),
	818:  uint16(1),
	819:  uint16(sym__text),
	820:  uint16(99),
	821:  uint16(1),
	822:  uint16(sym_function_link),
	823:  uint16(146),
	824:  uint16(1),
	826:  uint16(37),
	827:  uint16(1),
	828:  uint16(sym_storageclass),
	829:  uint16(122),
	830:  uint16(1),
	831:  uint16(sym_description),
	832:  uint16(125),
	833:  uint16(2),
	834:  uint16(sym_tag),
	835:  uint16(aux_sym_document_repeat2),
	836:  uint16(59),
	837:  uint16(4),
	838:  uint16(sym_emphasis),
	839:  uint16(sym_code_word),
	840:  uint16(sym_link),
	841:  uint16(aux_sym_description_repeat1),
	842:  uint16(11),
	843:  uint16(77),
	844:  uint16(1),
	845:  uint16(anon_sym_BSLASHa),
	846:  uint16(79),
	847:  uint16(1),
	848:  uint16(anon_sym_BSLASHc),
	849:  uint16(81),
	850:  uint16(1),
	851:  uint16(anon_sym_LTa),
	852:  uint16(83),
	853:  uint16(1),
	854:  uint16(sym__text),
	855:  uint16(99),
	856:  uint16(1),
	857:  uint16(sym_function_link),
	858:  uint16(148),
	859:  uint16(1),
	860:  uint16(anon_sym_COMMA),
	861:  uint16(21),
	862:  uint16(1),
	863:  uint16(aux_sym_tag_repeat1),
	864:  uint16(131),
	865:  uint16(1),
	866:  uint16(sym_description),
	867:  uint16(89),
	868:  uint16(2),
	870:  uint16(anon_sym_LBRACK),
	871:  uint16(59),
	872:  uint16(4),
	873:  uint16(sym_emphasis),
	874:  uint16(sym_code_word),
	875:  uint16(sym_link),
	876:  uint16(aux_sym_description_repeat1),
	877:  uint16(87),
	878:  uint16(6),
	879:  uint16(sym_tag_name_with_argument),
	880:  uint16(sym_tag_name_with_multiple_arguments),
	881:  uint16(sym_tag_name_with_types),
	882:  uint16(sym_tag_name_with_self_types),
	883:  uint16(sym_tag_name_with_type),
	884:  uint16(sym_tag_name),
	885:  uint16(11),
	886:  uint16(77),
	887:  uint16(1),
	888:  uint16(anon_sym_BSLASHa),
	889:  uint16(79),
	890:  uint16(1),
	891:  uint16(anon_sym_BSLASHc),
	892:  uint16(81),
	893:  uint16(1),
	894:  uint16(anon_sym_LTa),
	895:  uint16(83),
	896:  uint16(1),
	897:  uint16(sym__text),
	898:  uint16(99),
	899:  uint16(1),
	900:  uint16(sym_function_link),
	901:  uint16(150),
	902:  uint16(1),
	903:  uint16(aux_sym_identifier_token1),
	904:  uint16(38),
	905:  uint16(1),
	906:  uint16(sym_identifier),
	907:  uint16(128),
	908:  uint16(1),
	909:  uint16(sym_description),
	910:  uint16(47),
	911:  uint16(2),
	913:  uint16(anon_sym_LBRACK),
	914:  uint16(59),
	915:  uint16(4),
	916:  uint16(sym_emphasis),
	917:  uint16(sym_code_word),
	918:  uint16(sym_link),
	919:  uint16(aux_sym_description_repeat1),
	920:  uint16(41),
	921:  uint16(6),
	922:  uint16(sym_tag_name_with_argument),
	923:  uint16(sym_tag_name_with_multiple_arguments),
	924:  uint16(sym_tag_name_with_types),
	925:  uint16(sym_tag_name_with_self_types),
	926:  uint16(sym_tag_name_with_type),
	927:  uint16(sym_tag_name),
	928:  uint16(11),
	929:  uint16(77),
	930:  uint16(1),
	931:  uint16(anon_sym_BSLASHa),
	932:  uint16(79),
	933:  uint16(1),
	934:  uint16(anon_sym_BSLASHc),
	935:  uint16(81),
	936:  uint16(1),
	937:  uint16(anon_sym_LTa),
	938:  uint16(83),
	939:  uint16(1),
	940:  uint16(sym__text),
	941:  uint16(99),
	942:  uint16(1),
	943:  uint16(sym_function_link),
	944:  uint16(148),
	945:  uint16(1),
	946:  uint16(anon_sym_COMMA),
	947:  uint16(79),
	948:  uint16(1),
	949:  uint16(aux_sym_tag_repeat1),
	950:  uint16(127),
	951:  uint16(1),
	952:  uint16(sym_description),
	953:  uint16(105),
	954:  uint16(2),
	956:  uint16(anon_sym_LBRACK),
	957:  uint16(59),
	958:  uint16(4),
	959:  uint16(sym_emphasis),
	960:  uint16(sym_code_word),
	961:  uint16(sym_link),
	962:  uint16(aux_sym_description_repeat1),
	963:  uint16(103),
	964:  uint16(6),
	965:  uint16(sym_tag_name_with_argument),
	966:  uint16(sym_tag_name_with_multiple_arguments),
	967:  uint16(sym_tag_name_with_types),
	968:  uint16(sym_tag_name_with_self_types),
	969:  uint16(sym_tag_name_with_type),
	970:  uint16(sym_tag_name),
	971:  uint16(5),
	972:  uint16(152),
	973:  uint16(1),
	974:  uint16(anon_sym_COLON_COLON),
	975:  uint16(154),
	976:  uint16(1),
	977:  uint16(anon_sym_LPAREN),
	978:  uint16(28),
	979:  uint16(1),
	980:  uint16(aux_sym_qualified_identifier_repeat1),
	981:  uint16(138),
	982:  uint16(5),
	983:  uint16(sym_code_block_start),
	984:  uint16(anon_sym_LBRACK),
	985:  uint16(anon_sym_LTa),
	986:  uint16(sym_function_link),
	987:  uint16(anon_sym_SLASH),
	988:  uint16(140),
	989:  uint16(12),
	990:  uint16(sym_tag_name_with_argument),
	991:  uint16(sym_tag_name_with_multiple_arguments),
	992:  uint16(sym_tag_name_with_types),
	993:  uint16(sym_tag_name_with_self_types),
	994:  uint16(sym_tag_name_with_type),
	995:  uint16(sym_tag_name),
	996:  uint16(anon_sym_BSLASHa),
	997:  uint16(anon_sym_BSLASHc),
	998:  uint16(anon_sym_ATcode),
	999:  uint16(sym__text),
	1000: uint16(anon_sym_STAR_SLASH),
	1001: uint16(sym__text_line),
	1002: uint16(11),
	1003: uint16(77),
	1004: uint16(1),
	1005: uint16(anon_sym_BSLASHa),
	1006: uint16(79),
	1007: uint16(1),
	1008: uint16(anon_sym_BSLASHc),
	1009: uint16(81),
	1010: uint16(1),
	1011: uint16(anon_sym_LTa),
	1012: uint16(83),
	1013: uint16(1),
	1014: uint16(sym__text),
	1015: uint16(99),
	1016: uint16(1),
	1017: uint16(sym_function_link),
	1018: uint16(148),
	1019: uint16(1),
	1020: uint16(anon_sym_COMMA),
	1021: uint16(79),
	1022: uint16(1),
	1023: uint16(aux_sym_tag_repeat1),
	1024: uint16(134),
	1025: uint16(1),
	1026: uint16(sym_description),
	1027: uint16(93),
	1028: uint16(2),
	1030: uint16(anon_sym_LBRACK),
	1031: uint16(59),
	1032: uint16(4),
	1033: uint16(sym_emphasis),
	1034: uint16(sym_code_word),
	1035: uint16(sym_link),
	1036: uint16(aux_sym_description_repeat1),
	1037: uint16(91),
	1038: uint16(6),
	1039: uint16(sym_tag_name_with_argument),
	1040: uint16(sym_tag_name_with_multiple_arguments),
	1041: uint16(sym_tag_name_with_types),
	1042: uint16(sym_tag_name_with_self_types),
	1043: uint16(sym_tag_name_with_type),
	1044: uint16(sym_tag_name),
	1045: uint16(4),
	1046: uint16(142),
	1047: uint16(1),
	1048: uint16(anon_sym_COLON_COLON),
	1049: uint16(25),
	1050: uint16(1),
	1051: uint16(aux_sym_qualified_identifier_repeat1),
	1052: uint16(156),
	1053: uint16(6),
	1054: uint16(sym_code_block_start),
	1055: uint16(anon_sym_COMMA),
	1056: uint16(anon_sym_LBRACK),
	1057: uint16(anon_sym_LTa),
	1058: uint16(sym_function_link),
	1059: uint16(anon_sym_SLASH),
	1060: uint16(158),
	1061: uint16(12),
	1062: uint16(sym_tag_name_with_argument),
	1063: uint16(sym_tag_name_with_multiple_arguments),
	1064: uint16(sym_tag_name_with_types),
	1065: uint16(sym_tag_name_with_self_types),
	1066: uint16(sym_tag_name_with_type),
	1067: uint16(sym_tag_name),
	1068: uint16(anon_sym_BSLASHa),
	1069: uint16(anon_sym_BSLASHc),
	1070: uint16(anon_sym_ATcode),
	1071: uint16(sym__text),
	1072: uint16(anon_sym_STAR_SLASH),
	1073: uint16(sym__text_line),
	1074: uint16(4),
	1075: uint16(164),
	1076: uint16(1),
	1077: uint16(anon_sym_COLON_COLON),
	1078: uint16(25),
	1079: uint16(1),
	1080: uint16(aux_sym_qualified_identifier_repeat1),
	1081: uint16(160),
	1082: uint16(6),
	1083: uint16(sym_code_block_start),
	1084: uint16(anon_sym_COMMA),
	1085: uint16(anon_sym_LBRACK),
	1086: uint16(anon_sym_LTa),
	1087: uint16(sym_function_link),
	1088: uint16(anon_sym_SLASH),
	1089: uint16(162),
	1090: uint16(12),
	1091: uint16(sym_tag_name_with_argument),
	1092: uint16(sym_tag_name_with_multiple_arguments),
	1093: uint16(sym_tag_name_with_types),
	1094: uint16(sym_tag_name_with_self_types),
	1095: uint16(sym_tag_name_with_type),
	1096: uint16(sym_tag_name),
	1097: uint16(anon_sym_BSLASHa),
	1098: uint16(anon_sym_BSLASHc),
	1099: uint16(anon_sym_ATcode),
	1100: uint16(sym__text),
	1101: uint16(anon_sym_STAR_SLASH),
	1102: uint16(sym__text_line),
	1103: uint16(11),
	1104: uint16(77),
	1105: uint16(1),
	1106: uint16(anon_sym_BSLASHa),
	1107: uint16(79),
	1108: uint16(1),
	1109: uint16(anon_sym_BSLASHc),
	1110: uint16(81),
	1111: uint16(1),
	1112: uint16(anon_sym_LTa),
	1113: uint16(83),
	1114: uint16(1),
	1115: uint16(sym__text),
	1116: uint16(99),
	1117: uint16(1),
	1118: uint16(sym_function_link),
	1119: uint16(148),
	1120: uint16(1),
	1121: uint16(anon_sym_COMMA),
	1122: uint16(23),
	1123: uint16(1),
	1124: uint16(aux_sym_tag_repeat1),
	1125: uint16(127),
	1126: uint16(1),
	1127: uint16(sym_description),
	1128: uint16(105),
	1129: uint16(2),
	1131: uint16(anon_sym_LBRACK),
	1132: uint16(59),
	1133: uint16(4),
	1134: uint16(sym_emphasis),
	1135: uint16(sym_code_word),
	1136: uint16(sym_link),
	1137: uint16(aux_sym_description_repeat1),
	1138: uint16(103),
	1139: uint16(6),
	1140: uint16(sym_tag_name_with_argument),
	1141: uint16(sym_tag_name_with_multiple_arguments),
	1142: uint16(sym_tag_name_with_types),
	1143: uint16(sym_tag_name_with_self_types),
	1144: uint16(sym_tag_name_with_type),
	1145: uint16(sym_tag_name),
	1146: uint16(2),
	1147: uint16(167),
	1148: uint16(7),
	1149: uint16(sym_code_block_start),
	1150: uint16(anon_sym_COMMA),
	1151: uint16(anon_sym_LPAREN),
	1152: uint16(anon_sym_LBRACK),
	1153: uint16(anon_sym_LTa),
	1154: uint16(sym_function_link),
	1155: uint16(anon_sym_SLASH),
	1156: uint16(169),
	1157: uint16(13),
	1158: uint16(sym_tag_name_with_argument),
	1159: uint16(sym_tag_name_with_multiple_arguments),
	1160: uint16(sym_tag_name_with_types),
	1161: uint16(sym_tag_name_with_self_types),
	1162: uint16(sym_tag_name_with_type),
	1163: uint16(sym_tag_name),
	1164: uint16(anon_sym_COLON_COLON),
	1165: uint16(anon_sym_BSLASHa),
	1166: uint16(anon_sym_BSLASHc),
	1167: uint16(anon_sym_ATcode),
	1168: uint16(sym__text),
	1169: uint16(anon_sym_STAR_SLASH),
	1170: uint16(sym__text_line),
	1171: uint16(4),
	1172: uint16(152),
	1173: uint16(1),
	1174: uint16(anon_sym_COLON_COLON),
	1175: uint16(34),
	1176: uint16(1),
	1177: uint16(aux_sym_qualified_identifier_repeat1),
	1178: uint16(156),
	1179: uint16(5),
	1180: uint16(sym_code_block_start),
	1181: uint16(anon_sym_LBRACK),
	1182: uint16(anon_sym_LTa),
	1183: uint16(sym_function_link),
	1184: uint16(anon_sym_SLASH),
	1185: uint16(158),
	1186: uint16(12),
	1187: uint16(sym_tag_name_with_argument),
	1188: uint16(sym_tag_name_with_multiple_arguments),
	1189: uint16(sym_tag_name_with_types),
	1190: uint16(sym_tag_name_with_self_types),
	1191: uint16(sym_tag_name_with_type),
	1192: uint16(sym_tag_name),
	1193: uint16(anon_sym_BSLASHa),
	1194: uint16(anon_sym_BSLASHc),
	1195: uint16(anon_sym_ATcode),
	1196: uint16(sym__text),
	1197: uint16(anon_sym_STAR_SLASH),
	1198: uint16(sym__text_line),
	1199: uint16(4),
	1200: uint16(171),
	1201: uint16(1),
	1202: uint16(anon_sym_COMMA),
	1203: uint16(29),
	1204: uint16(1),
	1205: uint16(aux_sym_tag_repeat1),
	1206: uint16(176),
	1207: uint16(5),
	1208: uint16(sym_code_block_start),
	1209: uint16(anon_sym_LBRACK),
	1210: uint16(anon_sym_LTa),
	1211: uint16(sym_function_link),
	1212: uint16(anon_sym_SLASH),
	1213: uint16(174),
	1214: uint16(12),
	1215: uint16(sym_tag_name_with_argument),
	1216: uint16(sym_tag_name_with_multiple_arguments),
	1217: uint16(sym_tag_name_with_types),
	1218: uint16(sym_tag_name_with_self_types),
	1219: uint16(sym_tag_name_with_type),
	1220: uint16(sym_tag_name),
	1221: uint16(anon_sym_BSLASHa),
	1222: uint16(anon_sym_BSLASHc),
	1223: uint16(anon_sym_ATcode),
	1224: uint16(sym__text),
	1225: uint16(anon_sym_STAR_SLASH),
	1226: uint16(sym__text_line),
	1227: uint16(2),
	1228: uint16(167),
	1229: uint16(6),
	1230: uint16(sym_code_block_start),
	1231: uint16(anon_sym_LPAREN),
	1232: uint16(anon_sym_LBRACK),
	1233: uint16(anon_sym_LTa),
	1234: uint16(sym_function_link),
	1235: uint16(anon_sym_SLASH),
	1236: uint16(169),
	1237: uint16(13),
	1238: uint16(sym_tag_name_with_argument),
	1239: uint16(sym_tag_name_with_multiple_arguments),
	1240: uint16(sym_tag_name_with_types),
	1241: uint16(sym_tag_name_with_self_types),
	1242: uint16(sym_tag_name_with_type),
	1243: uint16(sym_tag_name),
	1244: uint16(anon_sym_COLON_COLON),
	1245: uint16(anon_sym_BSLASHa),
	1246: uint16(anon_sym_BSLASHc),
	1247: uint16(anon_sym_ATcode),
	1248: uint16(sym__text),
	1249: uint16(anon_sym_STAR_SLASH),
	1250: uint16(sym__text_line),
	1251: uint16(5),
	1252: uint16(180),
	1253: uint16(1),
	1254: uint16(sym_tag_name),
	1255: uint16(185),
	1256: uint16(1),
	1257: uint16(sym_brief_text),
	1258: uint16(31),
	1259: uint16(1),
	1260: uint16(aux_sym_brief_description_repeat1),
	1261: uint16(183),
	1262: uint16(5),
	1263: uint16(sym_code_block_start),
	1264: uint16(anon_sym_LBRACK),
	1265: uint16(anon_sym_LTa),
	1266: uint16(sym_function_link),
	1267: uint16(anon_sym_SLASH),
	1268: uint16(178),
	1269: uint16(11),
	1270: uint16(sym_tag_name_with_argument),
	1271: uint16(sym_tag_name_with_multiple_arguments),
	1272: uint16(sym_tag_name_with_types),
	1273: uint16(sym_tag_name_with_self_types),
	1274: uint16(sym_tag_name_with_type),
	1275: uint16(anon_sym_BSLASHa),
	1276: uint16(anon_sym_BSLASHc),
	1277: uint16(anon_sym_ATcode),
	1278: uint16(sym__text),
	1279: uint16(anon_sym_STAR_SLASH),
	1280: uint16(sym__text_line),
	1281: uint16(2),
	1282: uint16(160),
	1283: uint16(6),
	1284: uint16(sym_code_block_start),
	1285: uint16(anon_sym_COMMA),
	1286: uint16(anon_sym_LBRACK),
	1287: uint16(anon_sym_LTa),
	1288: uint16(sym_function_link),
	1289: uint16(anon_sym_SLASH),
	1290: uint16(162),
	1291: uint16(13),
	1292: uint16(sym_tag_name_with_argument),
	1293: uint16(sym_tag_name_with_multiple_arguments),
	1294: uint16(sym_tag_name_with_types),
	1295: uint16(sym_tag_name_with_self_types),
	1296: uint16(sym_tag_name_with_type),
	1297: uint16(sym_tag_name),
	1298: uint16(anon_sym_COLON_COLON),
	1299: uint16(anon_sym_BSLASHa),
	1300: uint16(anon_sym_BSLASHc),
	1301: uint16(anon_sym_ATcode),
	1302: uint16(sym__text),
	1303: uint16(anon_sym_STAR_SLASH),
	1304: uint16(sym__text_line),
	1305: uint16(2),
	1306: uint16(167),
	1307: uint16(6),
	1308: uint16(sym_code_block_start),
	1309: uint16(anon_sym_COMMA),
	1310: uint16(anon_sym_LBRACK),
	1311: uint16(anon_sym_LTa),
	1312: uint16(sym_function_link),
	1313: uint16(anon_sym_SLASH),
	1314: uint16(169),
	1315: uint16(13),
	1316: uint16(sym_tag_name_with_argument),
	1317: uint16(sym_tag_name_with_multiple_arguments),
	1318: uint16(sym_tag_name_with_types),
	1319: uint16(sym_tag_name_with_self_types),
	1320: uint16(sym_tag_name_with_type),
	1321: uint16(sym_tag_name),
	1322: uint16(anon_sym_COLON_COLON),
	1323: uint16(anon_sym_BSLASHa),
	1324: uint16(anon_sym_BSLASHc),
	1325: uint16(anon_sym_ATcode),
	1326: uint16(sym__text),
	1327: uint16(anon_sym_STAR_SLASH),
	1328: uint16(sym__text_line),
	1329: uint16(4),
	1330: uint16(188),
	1331: uint16(1),
	1332: uint16(anon_sym_COLON_COLON),
	1333: uint16(34),
	1334: uint16(1),
	1335: uint16(aux_sym_qualified_identifier_repeat1),
	1336: uint16(160),
	1337: uint16(5),
	1338: uint16(sym_code_block_start),
	1339: uint16(anon_sym_LBRACK),
	1340: uint16(anon_sym_LTa),
	1341: uint16(sym_function_link),
	1342: uint16(anon_sym_SLASH),
	1343: uint16(162),
	1344: uint16(12),
	1345: uint16(sym_tag_name_with_argument),
	1346: uint16(sym_tag_name_with_multiple_arguments),
	1347: uint16(sym_tag_name_with_types),
	1348: uint16(sym_tag_name_with_self_types),
	1349: uint16(sym_tag_name_with_type),
	1350: uint16(sym_tag_name),
	1351: uint16(anon_sym_BSLASHa),
	1352: uint16(anon_sym_BSLASHc),
	1353: uint16(anon_sym_ATcode),
	1354: uint16(sym__text),
	1355: uint16(anon_sym_STAR_SLASH),
	1356: uint16(sym__text_line),
	1357: uint16(5),
	1358: uint16(193),
	1359: uint16(1),
	1360: uint16(sym_tag_name),
	1361: uint16(197),
	1362: uint16(1),
	1363: uint16(sym_brief_text),
	1364: uint16(31),
	1365: uint16(1),
	1366: uint16(aux_sym_brief_description_repeat1),
	1367: uint16(195),
	1368: uint16(5),
	1369: uint16(sym_code_block_start),
	1370: uint16(anon_sym_LBRACK),
	1371: uint16(anon_sym_LTa),
	1372: uint16(sym_function_link),
	1373: uint16(anon_sym_SLASH),
	1374: uint16(191),
	1375: uint16(11),
	1376: uint16(sym_tag_name_with_argument),
	1377: uint16(sym_tag_name_with_multiple_arguments),
	1378: uint16(sym_tag_name_with_types),
	1379: uint16(sym_tag_name_with_self_types),
	1380: uint16(sym_tag_name_with_type),
	1381: uint16(anon_sym_BSLASHa),
	1382: uint16(anon_sym_BSLASHc),
	1383: uint16(anon_sym_ATcode),
	1384: uint16(sym__text),
	1385: uint16(anon_sym_STAR_SLASH),
	1386: uint16(sym__text_line),
	1387: uint16(9),
	1388: uint16(77),
	1389: uint16(1),
	1390: uint16(anon_sym_BSLASHa),
	1391: uint16(79),
	1392: uint16(1),
	1393: uint16(anon_sym_BSLASHc),
	1394: uint16(81),
	1395: uint16(1),
	1396: uint16(anon_sym_LTa),
	1397: uint16(83),
	1398: uint16(1),
	1399: uint16(sym__text),
	1400: uint16(99),
	1401: uint16(1),
	1402: uint16(sym_function_link),
	1403: uint16(131),
	1404: uint16(1),
	1405: uint16(sym_description),
	1406: uint16(89),
	1407: uint16(2),
	1409: uint16(anon_sym_LBRACK),
	1410: uint16(59),
	1411: uint16(4),
	1412: uint16(sym_emphasis),
	1413: uint16(sym_code_word),
	1414: uint16(sym_link),
	1415: uint16(aux_sym_description_repeat1),
	1416: uint16(87),
	1417: uint16(6),
	1418: uint16(sym_tag_name_with_argument),
	1419: uint16(sym_tag_name_with_multiple_arguments),
	1420: uint16(sym_tag_name_with_types),
	1421: uint16(sym_tag_name_with_self_types),
	1422: uint16(sym_tag_name_with_type),
	1423: uint16(sym_tag_name),
	1424: uint16(9),
	1425: uint16(77),
	1426: uint16(1),
	1427: uint16(anon_sym_BSLASHa),
	1428: uint16(79),
	1429: uint16(1),
	1430: uint16(anon_sym_BSLASHc),
	1431: uint16(81),
	1432: uint16(1),
	1433: uint16(anon_sym_LTa),
	1434: uint16(83),
	1435: uint16(1),
	1436: uint16(sym__text),
	1437: uint16(99),
	1438: uint16(1),
	1439: uint16(sym_function_link),
	1440: uint16(128),
	1441: uint16(1),
	1442: uint16(sym_description),
	1443: uint16(47),
	1444: uint16(2),
	1446: uint16(anon_sym_LBRACK),
	1447: uint16(59),
	1448: uint16(4),
	1449: uint16(sym_emphasis),
	1450: uint16(sym_code_word),
	1451: uint16(sym_link),
	1452: uint16(aux_sym_description_repeat1),
	1453: uint16(41),
	1454: uint16(6),
	1455: uint16(sym_tag_name_with_argument),
	1456: uint16(sym_tag_name_with_multiple_arguments),
	1457: uint16(sym_tag_name_with_types),
	1458: uint16(sym_tag_name_with_self_types),
	1459: uint16(sym_tag_name_with_type),
	1460: uint16(sym_tag_name),
	1461: uint16(9),
	1462: uint16(77),
	1463: uint16(1),
	1464: uint16(anon_sym_BSLASHa),
	1465: uint16(79),
	1466: uint16(1),
	1467: uint16(anon_sym_BSLASHc),
	1468: uint16(81),
	1469: uint16(1),
	1470: uint16(anon_sym_LTa),
	1471: uint16(83),
	1472: uint16(1),
	1473: uint16(sym__text),
	1474: uint16(99),
	1475: uint16(1),
	1476: uint16(sym_function_link),
	1477: uint16(136),
	1478: uint16(1),
	1479: uint16(sym_description),
	1480: uint16(109),
	1481: uint16(2),
	1483: uint16(anon_sym_LBRACK),
	1484: uint16(59),
	1485: uint16(4),
	1486: uint16(sym_emphasis),
	1487: uint16(sym_code_word),
	1488: uint16(sym_link),
	1489: uint16(aux_sym_description_repeat1),
	1490: uint16(107),
	1491: uint16(6),
	1492: uint16(sym_tag_name_with_argument),
	1493: uint16(sym_tag_name_with_multiple_arguments),
	1494: uint16(sym_tag_name_with_types),
	1495: uint16(sym_tag_name_with_self_types),
	1496: uint16(sym_tag_name_with_type),
	1497: uint16(sym_tag_name),
	1498: uint16(2),
	1499: uint16(160),
	1500: uint16(5),
	1501: uint16(sym_code_block_start),
	1502: uint16(anon_sym_LBRACK),
	1503: uint16(anon_sym_LTa),
	1504: uint16(sym_function_link),
	1505: uint16(anon_sym_SLASH),
	1506: uint16(162),
	1507: uint16(13),
	1508: uint16(sym_tag_name_with_argument),
	1509: uint16(sym_tag_name_with_multiple_arguments),
	1510: uint16(sym_tag_name_with_types),
	1511: uint16(sym_tag_name_with_self_types),
	1512: uint16(sym_tag_name_with_type),
	1513: uint16(sym_tag_name),
	1514: uint16(anon_sym_COLON_COLON),
	1515: uint16(anon_sym_BSLASHa),
	1516: uint16(anon_sym_BSLASHc),
	1517: uint16(anon_sym_ATcode),
	1518: uint16(sym__text),
	1519: uint16(anon_sym_STAR_SLASH),
	1520: uint16(sym__text_line),
	1521: uint16(2),
	1522: uint16(199),
	1523: uint16(6),
	1524: uint16(sym_code_block_start),
	1525: uint16(anon_sym_COMMA),
	1526: uint16(anon_sym_LBRACK),
	1527: uint16(anon_sym_LTa),
	1528: uint16(sym_function_link),
	1529: uint16(anon_sym_SLASH),
	1530: uint16(201),
	1531: uint16(12),
	1532: uint16(sym_tag_name_with_argument),
	1533: uint16(sym_tag_name_with_multiple_arguments),
	1534: uint16(sym_tag_name_with_types),
	1535: uint16(sym_tag_name_with_self_types),
	1536: uint16(sym_tag_name_with_type),
	1537: uint16(sym_tag_name),
	1538: uint16(anon_sym_BSLASHa),
	1539: uint16(anon_sym_BSLASHc),
	1540: uint16(anon_sym_ATcode),
	1541: uint16(sym__text),
	1542: uint16(anon_sym_STAR_SLASH),
	1543: uint16(sym__text_line),
	1544: uint16(8),
	1545: uint16(203),
	1546: uint16(1),
	1547: uint16(anon_sym_COMMA),
	1548: uint16(205),
	1549: uint16(1),
	1550: uint16(aux_sym_tag_token1),
	1551: uint16(209),
	1552: uint16(1),
	1553: uint16(aux_sym_identifier_token1),
	1554: uint16(213),
	1555: uint16(1),
	1556: uint16(sym_function_link),
	1557: uint16(47),
	1558: uint16(1),
	1559: uint16(aux_sym_tag_repeat2),
	1560: uint16(109),
	1561: uint16(1),
	1562: uint16(sym_identifier),
	1563: uint16(211),
	1564: uint16(3),
	1565: uint16(sym_code_block_start),
	1566: uint16(anon_sym_LBRACK),
	1567: uint16(anon_sym_SLASH),
	1568: uint16(207),
	1569: uint16(9),
	1570: uint16(sym_tag_name_with_argument),
	1571: uint16(sym_tag_name_with_multiple_arguments),
	1572: uint16(sym_tag_name_with_types),
	1573: uint16(sym_tag_name_with_self_types),
	1574: uint16(sym_tag_name_with_type),
	1575: uint16(sym_tag_name),
	1576: uint16(anon_sym_ATcode),
	1577: uint16(anon_sym_STAR_SLASH),
	1578: uint16(sym__text_line),
	1579: uint16(2),
	1580: uint16(215),
	1581: uint16(6),
	1582: uint16(sym_code_block_start),
	1583: uint16(anon_sym_COMMA),
	1584: uint16(anon_sym_LBRACK),
	1585: uint16(anon_sym_LTa),
	1586: uint16(sym_function_link),
	1587: uint16(anon_sym_SLASH),
	1588: uint16(217),
	1589: uint16(12),
	1590: uint16(sym_tag_name_with_argument),
	1591: uint16(sym_tag_name_with_multiple_arguments),
	1592: uint16(sym_tag_name_with_types),
	1593: uint16(sym_tag_name_with_self_types),
	1594: uint16(sym_tag_name_with_type),
	1595: uint16(sym_tag_name),
	1596: uint16(anon_sym_BSLASHa),
	1597: uint16(anon_sym_BSLASHc),
	1598: uint16(anon_sym_ATcode),
	1599: uint16(sym__text),
	1600: uint16(anon_sym_STAR_SLASH),
	1601: uint16(sym__text_line),
	1602: uint16(2),
	1603: uint16(176),
	1604: uint16(6),
	1605: uint16(sym_code_block_start),
	1606: uint16(anon_sym_COMMA),
	1607: uint16(anon_sym_LBRACK),
	1608: uint16(anon_sym_LTa),
	1609: uint16(sym_function_link),
	1610: uint16(anon_sym_SLASH),
	1611: uint16(174),
	1612: uint16(12),
	1613: uint16(sym_tag_name_with_argument),
	1614: uint16(sym_tag_name_with_multiple_arguments),
	1615: uint16(sym_tag_name_with_types),
	1616: uint16(sym_tag_name_with_self_types),
	1617: uint16(sym_tag_name_with_type),
	1618: uint16(sym_tag_name),
	1619: uint16(anon_sym_BSLASHa),
	1620: uint16(anon_sym_BSLASHc),
	1621: uint16(anon_sym_ATcode),
	1622: uint16(sym__text),
	1623: uint16(anon_sym_STAR_SLASH),
	1624: uint16(sym__text_line),
	1625: uint16(2),
	1626: uint16(167),
	1627: uint16(5),
	1628: uint16(sym_code_block_start),
	1629: uint16(anon_sym_LBRACK),
	1630: uint16(anon_sym_LTa),
	1631: uint16(sym_function_link),
	1632: uint16(anon_sym_SLASH),
	1633: uint16(169),
	1634: uint16(13),
	1635: uint16(sym_tag_name_with_argument),
	1636: uint16(sym_tag_name_with_multiple_arguments),
	1637: uint16(sym_tag_name_with_types),
	1638: uint16(sym_tag_name_with_self_types),
	1639: uint16(sym_tag_name_with_type),
	1640: uint16(sym_tag_name),
	1641: uint16(anon_sym_COLON_COLON),
	1642: uint16(anon_sym_BSLASHa),
	1643: uint16(anon_sym_BSLASHc),
	1644: uint16(anon_sym_ATcode),
	1645: uint16(sym__text),
	1646: uint16(anon_sym_STAR_SLASH),
	1647: uint16(sym__text_line),
	1648: uint16(15),
	1649: uint16(11),
	1650: uint16(1),
	1651: uint16(sym_tag_name_with_argument),
	1652: uint16(13),
	1653: uint16(1),
	1654: uint16(sym_tag_name_with_multiple_arguments),
	1655: uint16(15),
	1656: uint16(1),
	1657: uint16(sym_tag_name_with_types),
	1658: uint16(17),
	1659: uint16(1),
	1660: uint16(sym_tag_name_with_self_types),
	1661: uint16(19),
	1662: uint16(1),
	1663: uint16(sym_tag_name_with_type),
	1664: uint16(21),
	1665: uint16(1),
	1666: uint16(sym_tag_name),
	1667: uint16(23),
	1668: uint16(1),
	1669: uint16(anon_sym_LBRACK),
	1670: uint16(33),
	1671: uint16(1),
	1672: uint16(anon_sym_ATcode),
	1673: uint16(39),
	1674: uint16(1),
	1675: uint16(sym_code_block_start),
	1676: uint16(219),
	1677: uint16(1),
	1678: uint16(anon_sym_SLASH),
	1679: uint16(221),
	1680: uint16(1),
	1681: uint16(anon_sym_STAR_SLASH),
	1682: uint16(223),
	1683: uint16(1),
	1684: uint16(sym__text_line),
	1685: uint16(12),
	1686: uint16(1),
	1687: uint16(sym_storageclass),
	1688: uint16(203),
	1689: uint16(1),
	1690: uint16(sym__multiline_end),
	1691: uint16(71),
	1692: uint16(3),
	1693: uint16(sym_tag),
	1694: uint16(sym_code_block),
	1695: uint16(aux_sym_document_repeat1),
	1696: uint16(15),
	1697: uint16(11),
	1698: uint16(1),
	1699: uint16(sym_tag_name_with_argument),
	1700: uint16(13),
	1701: uint16(1),
	1702: uint16(sym_tag_name_with_multiple_arguments),
	1703: uint16(15),
	1704: uint16(1),
	1705: uint16(sym_tag_name_with_types),
	1706: uint16(17),
	1707: uint16(1),
	1708: uint16(sym_tag_name_with_self_types),
	1709: uint16(19),
	1710: uint16(1),
	1711: uint16(sym_tag_name_with_type),
	1712: uint16(21),
	1713: uint16(1),
	1714: uint16(sym_tag_name),
	1715: uint16(23),
	1716: uint16(1),
	1717: uint16(anon_sym_LBRACK),
	1718: uint16(33),
	1719: uint16(1),
	1720: uint16(anon_sym_ATcode),
	1721: uint16(39),
	1722: uint16(1),
	1723: uint16(sym_code_block_start),
	1724: uint16(223),
	1725: uint16(1),
	1726: uint16(sym__text_line),
	1727: uint16(225),
	1728: uint16(1),
	1729: uint16(anon_sym_SLASH),
	1730: uint16(227),
	1731: uint16(1),
	1732: uint16(anon_sym_STAR_SLASH),
	1733: uint16(12),
	1734: uint16(1),
	1735: uint16(sym_storageclass),
	1736: uint16(168),
	1737: uint16(1),
	1738: uint16(sym__multiline_end),
	1739: uint16(71),
	1740: uint16(3),
	1741: uint16(sym_tag),
	1742: uint16(sym_code_block),
	1743: uint16(aux_sym_document_repeat1),
	1744: uint16(7),
	1745: uint16(203),
	1746: uint16(1),
	1747: uint16(anon_sym_COMMA),
	1748: uint16(209),
	1749: uint16(1),
	1750: uint16(aux_sym_identifier_token1),
	1751: uint16(233),
	1752: uint16(1),
	1753: uint16(sym_function_link),
	1754: uint16(66),
	1755: uint16(1),
	1756: uint16(aux_sym_tag_repeat2),
	1757: uint16(117),
	1758: uint16(1),
	1759: uint16(sym_identifier),
	1760: uint16(231),
	1761: uint16(3),
	1762: uint16(sym_code_block_start),
	1763: uint16(anon_sym_LBRACK),
	1764: uint16(anon_sym_SLASH),
	1765: uint16(229),
	1766: uint16(9),
	1767: uint16(sym_tag_name_with_argument),
	1768: uint16(sym_tag_name_with_multiple_arguments),
	1769: uint16(sym_tag_name_with_types),
	1770: uint16(sym_tag_name_with_self_types),
	1771: uint16(sym_tag_name_with_type),
	1772: uint16(sym_tag_name),
	1773: uint16(anon_sym_ATcode),
	1774: uint16(anon_sym_STAR_SLASH),
	1775: uint16(sym__text_line),
	1776: uint16(7),
	1777: uint16(203),
	1778: uint16(1),
	1779: uint16(anon_sym_COMMA),
	1780: uint16(209),
	1781: uint16(1),
	1782: uint16(aux_sym_identifier_token1),
	1783: uint16(239),
	1784: uint16(1),
	1785: uint16(sym_function_link),
	1786: uint16(49),
	1787: uint16(1),
	1788: uint16(aux_sym_tag_repeat2),
	1789: uint16(113),
	1790: uint16(1),
	1791: uint16(sym_identifier),
	1792: uint16(237),
	1793: uint16(3),
	1794: uint16(sym_code_block_start),
	1795: uint16(anon_sym_LBRACK),
	1796: uint16(anon_sym_SLASH),
	1797: uint16(235),
	1798: uint16(9),
	1799: uint16(sym_tag_name_with_argument),
	1800: uint16(sym_tag_name_with_multiple_arguments),
	1801: uint16(sym_tag_name_with_types),
	1802: uint16(sym_tag_name_with_self_types),
	1803: uint16(sym_tag_name_with_type),
	1804: uint16(sym_tag_name),
	1805: uint16(anon_sym_ATcode),
	1806: uint16(anon_sym_STAR_SLASH),
	1807: uint16(sym__text_line),
	1808: uint16(7),
	1809: uint16(203),
	1810: uint16(1),
	1811: uint16(anon_sym_COMMA),
	1812: uint16(209),
	1813: uint16(1),
	1814: uint16(aux_sym_identifier_token1),
	1815: uint16(245),
	1816: uint16(1),
	1817: uint16(sym_function_link),
	1818: uint16(66),
	1819: uint16(1),
	1820: uint16(aux_sym_tag_repeat2),
	1821: uint16(100),
	1822: uint16(1),
	1823: uint16(sym_identifier),
	1824: uint16(243),
	1825: uint16(3),
	1826: uint16(sym_code_block_start),
	1827: uint16(anon_sym_LBRACK),
	1828: uint16(anon_sym_SLASH),
	1829: uint16(241),
	1830: uint16(9),
	1831: uint16(sym_tag_name_with_argument),
	1832: uint16(sym_tag_name_with_multiple_arguments),
	1833: uint16(sym_tag_name_with_types),
	1834: uint16(sym_tag_name_with_self_types),
	1835: uint16(sym_tag_name_with_type),
	1836: uint16(sym_tag_name),
	1837: uint16(anon_sym_ATcode),
	1838: uint16(anon_sym_STAR_SLASH),
	1839: uint16(sym__text_line),
	1840: uint16(8),
	1841: uint16(247),
	1842: uint16(1),
	1843: uint16(anon_sym_BSLASHa),
	1844: uint16(250),
	1845: uint16(1),
	1846: uint16(anon_sym_BSLASHc),
	1847: uint16(253),
	1848: uint16(1),
	1849: uint16(anon_sym_LTa),
	1850: uint16(256),
	1851: uint16(1),
	1852: uint16(sym_function_link),
	1853: uint16(259),
	1854: uint16(1),
	1855: uint16(sym__text),
	1856: uint16(121),
	1857: uint16(2),
	1859: uint16(anon_sym_LBRACK),
	1860: uint16(50),
	1861: uint16(4),
	1862: uint16(sym_emphasis),
	1863: uint16(sym_code_word),
	1864: uint16(sym_link),
	1865: uint16(aux_sym_description_repeat1),
	1866: uint16(119),
	1867: uint16(6),
	1868: uint16(sym_tag_name_with_argument),
	1869: uint16(sym_tag_name_with_multiple_arguments),
	1870: uint16(sym_tag_name_with_types),
	1871: uint16(sym_tag_name_with_self_types),
	1872: uint16(sym_tag_name_with_type),
	1873: uint16(sym_tag_name),
	1874: uint16(15),
	1875: uint16(11),
	1876: uint16(1),
	1877: uint16(sym_tag_name_with_argument),
	1878: uint16(13),
	1879: uint16(1),
	1880: uint16(sym_tag_name_with_multiple_arguments),
	1881: uint16(15),
	1882: uint16(1),
	1883: uint16(sym_tag_name_with_types),
	1884: uint16(17),
	1885: uint16(1),
	1886: uint16(sym_tag_name_with_self_types),
	1887: uint16(19),
	1888: uint16(1),
	1889: uint16(sym_tag_name_with_type),
	1890: uint16(21),
	1891: uint16(1),
	1892: uint16(sym_tag_name),
	1893: uint16(23),
	1894: uint16(1),
	1895: uint16(anon_sym_LBRACK),
	1896: uint16(33),
	1897: uint16(1),
	1898: uint16(anon_sym_ATcode),
	1899: uint16(39),
	1900: uint16(1),
	1901: uint16(sym_code_block_start),
	1902: uint16(51),
	1903: uint16(1),
	1904: uint16(anon_sym_SLASH),
	1905: uint16(53),
	1906: uint16(1),
	1907: uint16(anon_sym_STAR_SLASH),
	1908: uint16(55),
	1909: uint16(1),
	1910: uint16(sym__text_line),
	1911: uint16(12),
	1912: uint16(1),
	1913: uint16(sym_storageclass),
	1914: uint16(170),
	1915: uint16(1),
	1916: uint16(sym__multiline_end),
	1917: uint16(45),
	1918: uint16(3),
	1919: uint16(sym_tag),
	1920: uint16(sym_code_block),
	1921: uint16(aux_sym_document_repeat1),
	1922: uint16(15),
	1923: uint16(11),
	1924: uint16(1),
	1925: uint16(sym_tag_name_with_argument),
	1926: uint16(13),
	1927: uint16(1),
	1928: uint16(sym_tag_name_with_multiple_arguments),
	1929: uint16(15),
	1930: uint16(1),
	1931: uint16(sym_tag_name_with_types),
	1932: uint16(17),
	1933: uint16(1),
	1934: uint16(sym_tag_name_with_self_types),
	1935: uint16(19),
	1936: uint16(1),
	1937: uint16(sym_tag_name_with_type),
	1938: uint16(21),
	1939: uint16(1),
	1940: uint16(sym_tag_name),
	1941: uint16(23),
	1942: uint16(1),
	1943: uint16(anon_sym_LBRACK),
	1944: uint16(33),
	1945: uint16(1),
	1946: uint16(anon_sym_ATcode),
	1947: uint16(39),
	1948: uint16(1),
	1949: uint16(sym_code_block_start),
	1950: uint16(51),
	1951: uint16(1),
	1952: uint16(anon_sym_SLASH),
	1953: uint16(53),
	1954: uint16(1),
	1955: uint16(anon_sym_STAR_SLASH),
	1956: uint16(223),
	1957: uint16(1),
	1958: uint16(sym__text_line),
	1959: uint16(12),
	1960: uint16(1),
	1961: uint16(sym_storageclass),
	1962: uint16(170),
	1963: uint16(1),
	1964: uint16(sym__multiline_end),
	1965: uint16(71),
	1966: uint16(3),
	1967: uint16(sym_tag),
	1968: uint16(sym_code_block),
	1969: uint16(aux_sym_document_repeat1),
	1970: uint16(2),
	1971: uint16(264),
	1972: uint16(5),
	1973: uint16(sym_code_block_start),
	1974: uint16(anon_sym_LBRACK),
	1975: uint16(anon_sym_LTa),
	1976: uint16(sym_function_link),
	1977: uint16(anon_sym_SLASH),
	1978: uint16(262),
	1979: uint16(12),
	1980: uint16(sym_tag_name_with_argument),
	1981: uint16(sym_tag_name_with_multiple_arguments),
	1982: uint16(sym_tag_name_with_types),
	1983: uint16(sym_tag_name_with_self_types),
	1984: uint16(sym_tag_name_with_type),
	1985: uint16(sym_tag_name),
	1986: uint16(anon_sym_BSLASHa),
	1987: uint16(anon_sym_BSLASHc),
	1988: uint16(anon_sym_ATcode),
	1989: uint16(sym__text),
	1990: uint16(anon_sym_STAR_SLASH),
	1991: uint16(sym__text_line),
	1992: uint16(15),
	1993: uint16(11),
	1994: uint16(1),
	1995: uint16(sym_tag_name_with_argument),
	1996: uint16(13),
	1997: uint16(1),
	1998: uint16(sym_tag_name_with_multiple_arguments),
	1999: uint16(15),
	2000: uint16(1),
	2001: uint16(sym_tag_name_with_types),
	2002: uint16(17),
	2003: uint16(1),
	2004: uint16(sym_tag_name_with_self_types),
	2005: uint16(19),
	2006: uint16(1),
	2007: uint16(sym_tag_name_with_type),
	2008: uint16(21),
	2009: uint16(1),
	2010: uint16(sym_tag_name),
	2011: uint16(23),
	2012: uint16(1),
	2013: uint16(anon_sym_LBRACK),
	2014: uint16(33),
	2015: uint16(1),
	2016: uint16(anon_sym_ATcode),
	2017: uint16(39),
	2018: uint16(1),
	2019: uint16(sym_code_block_start),
	2020: uint16(219),
	2021: uint16(1),
	2022: uint16(anon_sym_SLASH),
	2023: uint16(221),
	2024: uint16(1),
	2025: uint16(anon_sym_STAR_SLASH),
	2026: uint16(266),
	2027: uint16(1),
	2028: uint16(sym__text_line),
	2029: uint16(12),
	2030: uint16(1),
	2031: uint16(sym_storageclass),
	2032: uint16(203),
	2033: uint16(1),
	2034: uint16(sym__multiline_end),
	2035: uint16(46),
	2036: uint16(3),
	2037: uint16(sym_tag),
	2038: uint16(sym_code_block),
	2039: uint16(aux_sym_document_repeat1),
	2040: uint16(2),
	2041: uint16(270),
	2042: uint16(5),
	2043: uint16(sym_code_block_start),
	2044: uint16(anon_sym_LBRACK),
	2045: uint16(anon_sym_LTa),
	2046: uint16(sym_function_link),
	2047: uint16(anon_sym_SLASH),
	2048: uint16(268),
	2049: uint16(12),
	2050: uint16(sym_tag_name_with_argument),
	2051: uint16(sym_tag_name_with_multiple_arguments),
	2052: uint16(sym_tag_name_with_types),
	2053: uint16(sym_tag_name_with_self_types),
	2054: uint16(sym_tag_name_with_type),
	2055: uint16(sym_tag_name),
	2056: uint16(anon_sym_BSLASHa),
	2057: uint16(anon_sym_BSLASHc),
	2058: uint16(anon_sym_ATcode),
	2059: uint16(sym__text),
	2060: uint16(anon_sym_STAR_SLASH),
	2061: uint16(sym__text_line),
	2062: uint16(2),
	2063: uint16(274),
	2064: uint16(5),
	2065: uint16(sym_code_block_start),
	2066: uint16(anon_sym_LBRACK),
	2067: uint16(anon_sym_LTa),
	2068: uint16(sym_function_link),
	2069: uint16(anon_sym_SLASH),
	2070: uint16(272),
	2071: uint16(12),
	2072: uint16(sym_tag_name_with_argument),
	2073: uint16(sym_tag_name_with_multiple_arguments),
	2074: uint16(sym_tag_name_with_types),
	2075: uint16(sym_tag_name_with_self_types),
	2076: uint16(sym_tag_name_with_type),
	2077: uint16(sym_tag_name),
	2078: uint16(anon_sym_BSLASHa),
	2079: uint16(anon_sym_BSLASHc),
	2080: uint16(anon_sym_ATcode),
	2081: uint16(sym__text),
	2082: uint16(anon_sym_STAR_SLASH),
	2083: uint16(sym__text_line),
	2084: uint16(2),
	2085: uint16(278),
	2086: uint16(5),
	2087: uint16(sym_code_block_start),
	2088: uint16(anon_sym_LBRACK),
	2089: uint16(anon_sym_LTa),
	2090: uint16(sym_function_link),
	2091: uint16(anon_sym_SLASH),
	2092: uint16(276),
	2093: uint16(12),
	2094: uint16(sym_tag_name_with_argument),
	2095: uint16(sym_tag_name_with_multiple_arguments),
	2096: uint16(sym_tag_name_with_types),
	2097: uint16(sym_tag_name_with_self_types),
	2098: uint16(sym_tag_name_with_type),
	2099: uint16(sym_tag_name),
	2100: uint16(anon_sym_BSLASHa),
	2101: uint16(anon_sym_BSLASHc),
	2102: uint16(anon_sym_ATcode),
	2103: uint16(sym__text),
	2104: uint16(anon_sym_STAR_SLASH),
	2105: uint16(sym__text_line),
	2106: uint16(2),
	2107: uint16(167),
	2108: uint16(5),
	2109: uint16(sym_code_block_start),
	2110: uint16(anon_sym_LBRACK),
	2111: uint16(anon_sym_LTa),
	2112: uint16(sym_function_link),
	2113: uint16(anon_sym_SLASH),
	2114: uint16(169),
	2115: uint16(12),
	2116: uint16(sym_tag_name_with_argument),
	2117: uint16(sym_tag_name_with_multiple_arguments),
	2118: uint16(sym_tag_name_with_types),
	2119: uint16(sym_tag_name_with_self_types),
	2120: uint16(sym_tag_name_with_type),
	2121: uint16(sym_tag_name),
	2122: uint16(anon_sym_BSLASHa),
	2123: uint16(anon_sym_BSLASHc),
	2124: uint16(anon_sym_ATcode),
	2125: uint16(sym__text),
	2126: uint16(anon_sym_STAR_SLASH),
	2127: uint16(sym__text_line),
	2128: uint16(8),
	2129: uint16(77),
	2130: uint16(1),
	2131: uint16(anon_sym_BSLASHa),
	2132: uint16(79),
	2133: uint16(1),
	2134: uint16(anon_sym_BSLASHc),
	2135: uint16(81),
	2136: uint16(1),
	2137: uint16(anon_sym_LTa),
	2138: uint16(280),
	2139: uint16(1),
	2140: uint16(sym_function_link),
	2141: uint16(282),
	2142: uint16(1),
	2143: uint16(sym__text),
	2144: uint16(113),
	2145: uint16(2),
	2147: uint16(anon_sym_LBRACK),
	2148: uint16(50),
	2149: uint16(4),
	2150: uint16(sym_emphasis),
	2151: uint16(sym_code_word),
	2152: uint16(sym_link),
	2153: uint16(aux_sym_description_repeat1),
	2154: uint16(111),
	2155: uint16(6),
	2156: uint16(sym_tag_name_with_argument),
	2157: uint16(sym_tag_name_with_multiple_arguments),
	2158: uint16(sym_tag_name_with_types),
	2159: uint16(sym_tag_name_with_self_types),
	2160: uint16(sym_tag_name_with_type),
	2161: uint16(sym_tag_name),
	2162: uint16(2),
	2163: uint16(286),
	2164: uint16(5),
	2165: uint16(sym_code_block_start),
	2166: uint16(anon_sym_LBRACK),
	2167: uint16(anon_sym_LTa),
	2168: uint16(sym_function_link),
	2169: uint16(anon_sym_SLASH),
	2170: uint16(284),
	2171: uint16(12),
	2172: uint16(sym_tag_name_with_argument),
	2173: uint16(sym_tag_name_with_multiple_arguments),
	2174: uint16(sym_tag_name_with_types),
	2175: uint16(sym_tag_name_with_self_types),
	2176: uint16(sym_tag_name_with_type),
	2177: uint16(sym_tag_name),
	2178: uint16(anon_sym_BSLASHa),
	2179: uint16(anon_sym_BSLASHc),
	2180: uint16(anon_sym_ATcode),
	2181: uint16(sym__text),
	2182: uint16(anon_sym_STAR_SLASH),
	2183: uint16(sym__text_line),
	2184: uint16(2),
	2185: uint16(290),
	2186: uint16(5),
	2187: uint16(sym_code_block_start),
	2188: uint16(anon_sym_LBRACK),
	2189: uint16(anon_sym_LTa),
	2190: uint16(sym_function_link),
	2191: uint16(anon_sym_SLASH),
	2192: uint16(288),
	2193: uint16(12),
	2194: uint16(sym_tag_name_with_argument),
	2195: uint16(sym_tag_name_with_multiple_arguments),
	2196: uint16(sym_tag_name_with_types),
	2197: uint16(sym_tag_name_with_self_types),
	2198: uint16(sym_tag_name_with_type),
	2199: uint16(sym_tag_name),
	2200: uint16(anon_sym_BSLASHa),
	2201: uint16(anon_sym_BSLASHc),
	2202: uint16(anon_sym_ATcode),
	2203: uint16(sym__text),
	2204: uint16(anon_sym_STAR_SLASH),
	2205: uint16(sym__text_line),
	2206: uint16(2),
	2207: uint16(215),
	2208: uint16(5),
	2209: uint16(sym_code_block_start),
	2210: uint16(anon_sym_LBRACK),
	2211: uint16(anon_sym_LTa),
	2212: uint16(sym_function_link),
	2213: uint16(anon_sym_SLASH),
	2214: uint16(217),
	2215: uint16(12),
	2216: uint16(sym_tag_name_with_argument),
	2217: uint16(sym_tag_name_with_multiple_arguments),
	2218: uint16(sym_tag_name_with_types),
	2219: uint16(sym_tag_name_with_self_types),
	2220: uint16(sym_tag_name_with_type),
	2221: uint16(sym_tag_name),
	2222: uint16(anon_sym_BSLASHa),
	2223: uint16(anon_sym_BSLASHc),
	2224: uint16(anon_sym_ATcode),
	2225: uint16(sym__text),
	2226: uint16(anon_sym_STAR_SLASH),
	2227: uint16(sym__text_line),
	2228: uint16(2),
	2229: uint16(199),
	2230: uint16(5),
	2231: uint16(sym_code_block_start),
	2232: uint16(anon_sym_LBRACK),
	2233: uint16(anon_sym_LTa),
	2234: uint16(sym_function_link),
	2235: uint16(anon_sym_SLASH),
	2236: uint16(201),
	2237: uint16(12),
	2238: uint16(sym_tag_name_with_argument),
	2239: uint16(sym_tag_name_with_multiple_arguments),
	2240: uint16(sym_tag_name_with_types),
	2241: uint16(sym_tag_name_with_self_types),
	2242: uint16(sym_tag_name_with_type),
	2243: uint16(sym_tag_name),
	2244: uint16(anon_sym_BSLASHa),
	2245: uint16(anon_sym_BSLASHc),
	2246: uint16(anon_sym_ATcode),
	2247: uint16(sym__text),
	2248: uint16(anon_sym_STAR_SLASH),
	2249: uint16(sym__text_line),
	2250: uint16(5),
	2251: uint16(292),
	2252: uint16(1),
	2253: uint16(anon_sym_COLON_COLON),
	2254: uint16(294),
	2255: uint16(1),
	2256: uint16(anon_sym_LPAREN),
	2257: uint16(67),
	2258: uint16(1),
	2259: uint16(aux_sym_qualified_identifier_repeat1),
	2260: uint16(138),
	2261: uint16(5),
	2263: uint16(anon_sym_COMMA),
	2264: uint16(anon_sym_LBRACK),
	2265: uint16(anon_sym_LTa),
	2266: uint16(sym_function_link),
	2267: uint16(140),
	2268: uint16(9),
	2269: uint16(sym_tag_name_with_argument),
	2270: uint16(sym_tag_name_with_multiple_arguments),
	2271: uint16(sym_tag_name_with_types),
	2272: uint16(sym_tag_name_with_self_types),
	2273: uint16(sym_tag_name_with_type),
	2274: uint16(sym_tag_name),
	2275: uint16(anon_sym_BSLASHa),
	2276: uint16(anon_sym_BSLASHc),
	2277: uint16(sym__text),
	2278: uint16(3),
	2279: uint16(298),
	2280: uint16(1),
	2281: uint16(aux_sym_tag_token1),
	2282: uint16(296),
	2283: uint16(5),
	2284: uint16(sym_code_block_start),
	2285: uint16(anon_sym_COMMA),
	2286: uint16(anon_sym_LBRACK),
	2287: uint16(sym_function_link),
	2288: uint16(anon_sym_SLASH),
	2289: uint16(300),
	2290: uint16(10),
	2291: uint16(sym_tag_name_with_argument),
	2292: uint16(sym_tag_name_with_multiple_arguments),
	2293: uint16(sym_tag_name_with_types),
	2294: uint16(sym_tag_name_with_self_types),
	2295: uint16(sym_tag_name_with_type),
	2296: uint16(sym_tag_name),
	2297: uint16(aux_sym_identifier_token1),
	2298: uint16(anon_sym_ATcode),
	2299: uint16(anon_sym_STAR_SLASH),
	2300: uint16(sym__text_line),
	2301: uint16(4),
	2302: uint16(302),
	2303: uint16(1),
	2304: uint16(anon_sym_COMMA),
	2305: uint16(66),
	2306: uint16(1),
	2307: uint16(aux_sym_tag_repeat2),
	2308: uint16(307),
	2309: uint16(4),
	2310: uint16(sym_code_block_start),
	2311: uint16(anon_sym_LBRACK),
	2312: uint16(sym_function_link),
	2313: uint16(anon_sym_SLASH),
	2314: uint16(305),
	2315: uint16(10),
	2316: uint16(sym_tag_name_with_argument),
	2317: uint16(sym_tag_name_with_multiple_arguments),
	2318: uint16(sym_tag_name_with_types),
	2319: uint16(sym_tag_name_with_self_types),
	2320: uint16(sym_tag_name_with_type),
	2321: uint16(sym_tag_name),
	2322: uint16(aux_sym_identifier_token1),
	2323: uint16(anon_sym_ATcode),
	2324: uint16(anon_sym_STAR_SLASH),
	2325: uint16(sym__text_line),
	2326: uint16(4),
	2327: uint16(292),
	2328: uint16(1),
	2329: uint16(anon_sym_COLON_COLON),
	2330: uint16(72),
	2331: uint16(1),
	2332: uint16(aux_sym_qualified_identifier_repeat1),
	2333: uint16(156),
	2334: uint16(5),
	2336: uint16(anon_sym_COMMA),
	2337: uint16(anon_sym_LBRACK),
	2338: uint16(anon_sym_LTa),
	2339: uint16(sym_function_link),
	2340: uint16(158),
	2341: uint16(9),
	2342: uint16(sym_tag_name_with_argument),
	2343: uint16(sym_tag_name_with_multiple_arguments),
	2344: uint16(sym_tag_name_with_types),
	2345: uint16(sym_tag_name_with_self_types),
	2346: uint16(sym_tag_name_with_type),
	2347: uint16(sym_tag_name),
	2348: uint16(anon_sym_BSLASHa),
	2349: uint16(anon_sym_BSLASHc),
	2350: uint16(sym__text),
	2351: uint16(2),
	2352: uint16(167),
	2353: uint16(6),
	2354: uint16(sym_code_block_start),
	2355: uint16(anon_sym_COMMA),
	2356: uint16(aux_sym_tag_token1),
	2357: uint16(anon_sym_LBRACK),
	2358: uint16(sym_function_link),
	2359: uint16(anon_sym_SLASH),
	2360: uint16(169),
	2361: uint16(10),
	2362: uint16(sym_tag_name_with_argument),
	2363: uint16(sym_tag_name_with_multiple_arguments),
	2364: uint16(sym_tag_name_with_types),
	2365: uint16(sym_tag_name_with_self_types),
	2366: uint16(sym_tag_name_with_type),
	2367: uint16(sym_tag_name),
	2368: uint16(aux_sym_identifier_token1),
	2369: uint16(anon_sym_ATcode),
	2370: uint16(anon_sym_STAR_SLASH),
	2371: uint16(sym__text_line),
	2372: uint16(2),
	2373: uint16(167),
	2374: uint16(6),
	2376: uint16(anon_sym_COMMA),
	2377: uint16(anon_sym_LPAREN),
	2378: uint16(anon_sym_LBRACK),
	2379: uint16(anon_sym_LTa),
	2380: uint16(sym_function_link),
	2381: uint16(169),
	2382: uint16(10),
	2383: uint16(sym_tag_name_with_argument),
	2384: uint16(sym_tag_name_with_multiple_arguments),
	2385: uint16(sym_tag_name_with_types),
	2386: uint16(sym_tag_name_with_self_types),
	2387: uint16(sym_tag_name_with_type),
	2388: uint16(sym_tag_name),
	2389: uint16(anon_sym_COLON_COLON),
	2390: uint16(anon_sym_BSLASHa),
	2391: uint16(anon_sym_BSLASHc),
	2392: uint16(sym__text),
	2393: uint16(5),
	2394: uint16(309),
	2395: uint16(1),
	2396: uint16(anon_sym_COLON_COLON),
	2397: uint16(311),
	2398: uint16(1),
	2399: uint16(anon_sym_LPAREN),
	2400: uint16(74),
	2401: uint16(1),
	2402: uint16(aux_sym_qualified_identifier_repeat1),
	2403: uint16(138),
	2404: uint16(4),
	2406: uint16(anon_sym_LBRACK),
	2407: uint16(anon_sym_LTa),
	2408: uint16(sym_function_link),
	2409: uint16(140),
	2410: uint16(9),
	2411: uint16(sym_tag_name_with_argument),
	2412: uint16(sym_tag_name_with_multiple_arguments),
	2413: uint16(sym_tag_name_with_types),
	2414: uint16(sym_tag_name_with_self_types),
	2415: uint16(sym_tag_name_with_type),
	2416: uint16(sym_tag_name),
	2417: uint16(anon_sym_BSLASHa),
	2418: uint16(anon_sym_BSLASHc),
	2419: uint16(sym__text),
	2420: uint16(14),
	2421: uint16(313),
	2422: uint16(1),
	2423: uint16(sym_tag_name_with_argument),
	2424: uint16(316),
	2425: uint16(1),
	2426: uint16(sym_tag_name_with_multiple_arguments),
	2427: uint16(319),
	2428: uint16(1),
	2429: uint16(sym_tag_name_with_types),
	2430: uint16(322),
	2431: uint16(1),
	2432: uint16(sym_tag_name_with_self_types),
	2433: uint16(325),
	2434: uint16(1),
	2435: uint16(sym_tag_name_with_type),
	2436: uint16(328),
	2437: uint16(1),
	2438: uint16(sym_tag_name),
	2439: uint16(331),
	2440: uint16(1),
	2441: uint16(anon_sym_LBRACK),
	2442: uint16(334),
	2443: uint16(1),
	2444: uint16(anon_sym_ATcode),
	2445: uint16(337),
	2446: uint16(1),
	2447: uint16(anon_sym_SLASH),
	2448: uint16(339),
	2449: uint16(1),
	2450: uint16(anon_sym_STAR_SLASH),
	2451: uint16(341),
	2452: uint16(1),
	2453: uint16(sym__text_line),
	2454: uint16(344),
	2455: uint16(1),
	2456: uint16(sym_code_block_start),
	2457: uint16(12),
	2458: uint16(1),
	2459: uint16(sym_storageclass),
	2460: uint16(71),
	2461: uint16(3),
	2462: uint16(sym_tag),
	2463: uint16(sym_code_block),
	2464: uint16(aux_sym_document_repeat1),
	2465: uint16(4),
	2466: uint16(347),
	2467: uint16(1),
	2468: uint16(anon_sym_COLON_COLON),
	2469: uint16(72),
	2470: uint16(1),
	2471: uint16(aux_sym_qualified_identifier_repeat1),
	2472: uint16(160),
	2473: uint16(5),
	2475: uint16(anon_sym_COMMA),
	2476: uint16(anon_sym_LBRACK),
	2477: uint16(anon_sym_LTa),
	2478: uint16(sym_function_link),
	2479: uint16(162),
	2480: uint16(9),
	2481: uint16(sym_tag_name_with_argument),
	2482: uint16(sym_tag_name_with_multiple_arguments),
	2483: uint16(sym_tag_name_with_types),
	2484: uint16(sym_tag_name_with_self_types),
	2485: uint16(sym_tag_name_with_type),
	2486: uint16(sym_tag_name),
	2487: uint16(anon_sym_BSLASHa),
	2488: uint16(anon_sym_BSLASHc),
	2489: uint16(sym__text),
	2490: uint16(4),
	2491: uint16(350),
	2492: uint16(1),
	2493: uint16(anon_sym_COLON_COLON),
	2494: uint16(73),
	2495: uint16(1),
	2496: uint16(aux_sym_qualified_identifier_repeat1),
	2497: uint16(160),
	2498: uint16(4),
	2500: uint16(anon_sym_LBRACK),
	2501: uint16(anon_sym_LTa),
	2502: uint16(sym_function_link),
	2503: uint16(162),
	2504: uint16(9),
	2505: uint16(sym_tag_name_with_argument),
	2506: uint16(sym_tag_name_with_multiple_arguments),
	2507: uint16(sym_tag_name_with_types),
	2508: uint16(sym_tag_name_with_self_types),
	2509: uint16(sym_tag_name_with_type),
	2510: uint16(sym_tag_name),
	2511: uint16(anon_sym_BSLASHa),
	2512: uint16(anon_sym_BSLASHc),
	2513: uint16(sym__text),
	2514: uint16(4),
	2515: uint16(309),
	2516: uint16(1),
	2517: uint16(anon_sym_COLON_COLON),
	2518: uint16(73),
	2519: uint16(1),
	2520: uint16(aux_sym_qualified_identifier_repeat1),
	2521: uint16(156),
	2522: uint16(4),
	2524: uint16(anon_sym_LBRACK),
	2525: uint16(anon_sym_LTa),
	2526: uint16(sym_function_link),
	2527: uint16(158),
	2528: uint16(9),
	2529: uint16(sym_tag_name_with_argument),
	2530: uint16(sym_tag_name_with_multiple_arguments),
	2531: uint16(sym_tag_name_with_types),
	2532: uint16(sym_tag_name_with_self_types),
	2533: uint16(sym_tag_name_with_type),
	2534: uint16(sym_tag_name),
	2535: uint16(anon_sym_BSLASHa),
	2536: uint16(anon_sym_BSLASHc),
	2537: uint16(sym__text),
	2538: uint16(2),
	2539: uint16(353),
	2540: uint16(5),
	2541: uint16(sym_code_block_start),
	2542: uint16(anon_sym_COMMA),
	2543: uint16(anon_sym_LBRACK),
	2544: uint16(sym_function_link),
	2545: uint16(anon_sym_SLASH),
	2546: uint16(355),
	2547: uint16(10),
	2548: uint16(sym_tag_name_with_argument),
	2549: uint16(sym_tag_name_with_multiple_arguments),
	2550: uint16(sym_tag_name_with_types),
	2551: uint16(sym_tag_name_with_self_types),
	2552: uint16(sym_tag_name_with_type),
	2553: uint16(sym_tag_name),
	2554: uint16(aux_sym_identifier_token1),
	2555: uint16(anon_sym_ATcode),
	2556: uint16(anon_sym_STAR_SLASH),
	2557: uint16(sym__text_line),
	2558: uint16(2),
	2559: uint16(160),
	2560: uint16(5),
	2562: uint16(anon_sym_COMMA),
	2563: uint16(anon_sym_LBRACK),
	2564: uint16(anon_sym_LTa),
	2565: uint16(sym_function_link),
	2566: uint16(162),
	2567: uint16(10),
	2568: uint16(sym_tag_name_with_argument),
	2569: uint16(sym_tag_name_with_multiple_arguments),
	2570: uint16(sym_tag_name_with_types),
	2571: uint16(sym_tag_name_with_self_types),
	2572: uint16(sym_tag_name_with_type),
	2573: uint16(sym_tag_name),
	2574: uint16(anon_sym_COLON_COLON),
	2575: uint16(anon_sym_BSLASHa),
	2576: uint16(anon_sym_BSLASHc),
	2577: uint16(sym__text),
	2578: uint16(2),
	2579: uint16(167),
	2580: uint16(5),
	2582: uint16(anon_sym_LPAREN),
	2583: uint16(anon_sym_LBRACK),
	2584: uint16(anon_sym_LTa),
	2585: uint16(sym_function_link),
	2586: uint16(169),
	2587: uint16(10),
	2588: uint16(sym_tag_name_with_argument),
	2589: uint16(sym_tag_name_with_multiple_arguments),
	2590: uint16(sym_tag_name_with_types),
	2591: uint16(sym_tag_name_with_self_types),
	2592: uint16(sym_tag_name_with_type),
	2593: uint16(sym_tag_name),
	2594: uint16(anon_sym_COLON_COLON),
	2595: uint16(anon_sym_BSLASHa),
	2596: uint16(anon_sym_BSLASHc),
	2597: uint16(sym__text),
	2598: uint16(2),
	2599: uint16(167),
	2600: uint16(5),
	2602: uint16(anon_sym_COMMA),
	2603: uint16(anon_sym_LBRACK),
	2604: uint16(anon_sym_LTa),
	2605: uint16(sym_function_link),
	2606: uint16(169),
	2607: uint16(10),
	2608: uint16(sym_tag_name_with_argument),
	2609: uint16(sym_tag_name_with_multiple_arguments),
	2610: uint16(sym_tag_name_with_types),
	2611: uint16(sym_tag_name_with_self_types),
	2612: uint16(sym_tag_name_with_type),
	2613: uint16(sym_tag_name),
	2614: uint16(anon_sym_COLON_COLON),
	2615: uint16(anon_sym_BSLASHa),
	2616: uint16(anon_sym_BSLASHc),
	2617: uint16(sym__text),
	2618: uint16(4),
	2619: uint16(357),
	2620: uint16(1),
	2621: uint16(anon_sym_COMMA),
	2622: uint16(79),
	2623: uint16(1),
	2624: uint16(aux_sym_tag_repeat1),
	2625: uint16(176),
	2626: uint16(4),
	2628: uint16(anon_sym_LBRACK),
	2629: uint16(anon_sym_LTa),
	2630: uint16(sym_function_link),
	2631: uint16(174),
	2632: uint16(9),
	2633: uint16(sym_tag_name_with_argument),
	2634: uint16(sym_tag_name_with_multiple_arguments),
	2635: uint16(sym_tag_name_with_types),
	2636: uint16(sym_tag_name_with_self_types),
	2637: uint16(sym_tag_name_with_type),
	2638: uint16(sym_tag_name),
	2639: uint16(anon_sym_BSLASHa),
	2640: uint16(anon_sym_BSLASHc),
	2641: uint16(sym__text),
	2642: uint16(5),
	2643: uint16(360),
	2644: uint16(1),
	2645: uint16(sym_tag_name),
	2646: uint16(363),
	2647: uint16(1),
	2648: uint16(sym_brief_text),
	2649: uint16(80),
	2650: uint16(1),
	2651: uint16(aux_sym_brief_description_repeat1),
	2652: uint16(183),
	2653: uint16(4),
	2655: uint16(anon_sym_LBRACK),
	2656: uint16(anon_sym_LTa),
	2657: uint16(sym_function_link),
	2658: uint16(178),
	2659: uint16(8),
	2660: uint16(sym_tag_name_with_argument),
	2661: uint16(sym_tag_name_with_multiple_arguments),
	2662: uint16(sym_tag_name_with_types),
	2663: uint16(sym_tag_name_with_self_types),
	2664: uint16(sym_tag_name_with_type),
	2665: uint16(anon_sym_BSLASHa),
	2666: uint16(anon_sym_BSLASHc),
	2667: uint16(sym__text),
	2668: uint16(5),
	2669: uint16(366),
	2670: uint16(1),
	2671: uint16(sym_tag_name),
	2672: uint16(368),
	2673: uint16(1),
	2674: uint16(sym_brief_text),
	2675: uint16(80),
	2676: uint16(1),
	2677: uint16(aux_sym_brief_description_repeat1),
	2678: uint16(195),
	2679: uint16(4),
	2681: uint16(anon_sym_LBRACK),
	2682: uint16(anon_sym_LTa),
	2683: uint16(sym_function_link),
	2684: uint16(191),
	2685: uint16(8),
	2686: uint16(sym_tag_name_with_argument),
	2687: uint16(sym_tag_name_with_multiple_arguments),
	2688: uint16(sym_tag_name_with_types),
	2689: uint16(sym_tag_name_with_self_types),
	2690: uint16(sym_tag_name_with_type),
	2691: uint16(anon_sym_BSLASHa),
	2692: uint16(anon_sym_BSLASHc),
	2693: uint16(sym__text),
	2694: uint16(2),
	2695: uint16(215),
	2696: uint16(5),
	2698: uint16(anon_sym_COMMA),
	2699: uint16(anon_sym_LBRACK),
	2700: uint16(anon_sym_LTa),
	2701: uint16(sym_function_link),
	2702: uint16(217),
	2703: uint16(9),
	2704: uint16(sym_tag_name_with_argument),
	2705: uint16(sym_tag_name_with_multiple_arguments),
	2706: uint16(sym_tag_name_with_types),
	2707: uint16(sym_tag_name_with_self_types),
	2708: uint16(sym_tag_name_with_type),
	2709: uint16(sym_tag_name),
	2710: uint16(anon_sym_BSLASHa),
	2711: uint16(anon_sym_BSLASHc),
	2712: uint16(sym__text),
	2713: uint16(2),
	2714: uint16(167),
	2715: uint16(4),
	2717: uint16(anon_sym_LBRACK),
	2718: uint16(anon_sym_LTa),
	2719: uint16(sym_function_link),
	2720: uint16(169),
	2721: uint16(10),
	2722: uint16(sym_tag_name_with_argument),
	2723: uint16(sym_tag_name_with_multiple_arguments),
	2724: uint16(sym_tag_name_with_types),
	2725: uint16(sym_tag_name_with_self_types),
	2726: uint16(sym_tag_name_with_type),
	2727: uint16(sym_tag_name),
	2728: uint16(anon_sym_COLON_COLON),
	2729: uint16(anon_sym_BSLASHa),
	2730: uint16(anon_sym_BSLASHc),
	2731: uint16(sym__text),
	2732: uint16(2),
	2733: uint16(160),
	2734: uint16(4),
	2736: uint16(anon_sym_LBRACK),
	2737: uint16(anon_sym_LTa),
	2738: uint16(sym_function_link),
	2739: uint16(162),
	2740: uint16(10),
	2741: uint16(sym_tag_name_with_argument),
	2742: uint16(sym_tag_name_with_multiple_arguments),
	2743: uint16(sym_tag_name_with_types),
	2744: uint16(sym_tag_name_with_self_types),
	2745: uint16(sym_tag_name_with_type),
	2746: uint16(sym_tag_name),
	2747: uint16(anon_sym_COLON_COLON),
	2748: uint16(anon_sym_BSLASHa),
	2749: uint16(anon_sym_BSLASHc),
	2750: uint16(sym__text),
	2751: uint16(2),
	2752: uint16(176),
	2753: uint16(5),
	2755: uint16(anon_sym_COMMA),
	2756: uint16(anon_sym_LBRACK),
	2757: uint16(anon_sym_LTa),
	2758: uint16(sym_function_link),
	2759: uint16(174),
	2760: uint16(9),
	2761: uint16(sym_tag_name_with_argument),
	2762: uint16(sym_tag_name_with_multiple_arguments),
	2763: uint16(sym_tag_name_with_types),
	2764: uint16(sym_tag_name_with_self_types),
	2765: uint16(sym_tag_name_with_type),
	2766: uint16(sym_tag_name),
	2767: uint16(anon_sym_BSLASHa),
	2768: uint16(anon_sym_BSLASHc),
	2769: uint16(sym__text),
	2770: uint16(2),
	2771: uint16(199),
	2772: uint16(5),
	2774: uint16(anon_sym_COMMA),
	2775: uint16(anon_sym_LBRACK),
	2776: uint16(anon_sym_LTa),
	2777: uint16(sym_function_link),
	2778: uint16(201),
	2779: uint16(9),
	2780: uint16(sym_tag_name_with_argument),
	2781: uint16(sym_tag_name_with_multiple_arguments),
	2782: uint16(sym_tag_name_with_types),
	2783: uint16(sym_tag_name_with_self_types),
	2784: uint16(sym_tag_name_with_type),
	2785: uint16(sym_tag_name),
	2786: uint16(anon_sym_BSLASHa),
	2787: uint16(anon_sym_BSLASHc),
	2788: uint16(sym__text),
	2789: uint16(8),
	2790: uint16(370),
	2791: uint16(1),
	2792: uint16(anon_sym_COMMA),
	2793: uint16(372),
	2794: uint16(1),
	2795: uint16(aux_sym_tag_token1),
	2796: uint16(374),
	2797: uint16(1),
	2798: uint16(aux_sym_identifier_token1),
	2799: uint16(376),
	2800: uint16(1),
	2801: uint16(sym_function_link),
	2802: uint16(95),
	2803: uint16(1),
	2804: uint16(aux_sym_tag_repeat2),
	2805: uint16(139),
	2806: uint16(1),
	2807: uint16(sym_identifier),
	2808: uint16(211),
	2809: uint16(2),
	2811: uint16(anon_sym_LBRACK),
	2812: uint16(207),
	2813: uint16(6),
	2814: uint16(sym_tag_name_with_argument),
	2815: uint16(sym_tag_name_with_multiple_arguments),
	2816: uint16(sym_tag_name_with_types),
	2817: uint16(sym_tag_name_with_self_types),
	2818: uint16(sym_tag_name_with_type),
	2819: uint16(sym_tag_name),
	2820: uint16(7),
	2821: uint16(370),
	2822: uint16(1),
	2823: uint16(anon_sym_COMMA),
	2824: uint16(374),
	2825: uint16(1),
	2826: uint16(aux_sym_identifier_token1),
	2827: uint16(378),
	2828: uint16(1),
	2829: uint16(sym_function_link),
	2830: uint16(92),
	2831: uint16(1),
	2832: uint16(aux_sym_tag_repeat2),
	2833: uint16(126),
	2834: uint16(1),
	2835: uint16(sym_identifier),
	2836: uint16(237),
	2837: uint16(2),
	2839: uint16(anon_sym_LBRACK),
	2840: uint16(235),
	2841: uint16(6),
	2842: uint16(sym_tag_name_with_argument),
	2843: uint16(sym_tag_name_with_multiple_arguments),
	2844: uint16(sym_tag_name_with_types),
	2845: uint16(sym_tag_name_with_self_types),
	2846: uint16(sym_tag_name_with_type),
	2847: uint16(sym_tag_name),
	2848: uint16(2),
	2849: uint16(286),
	2850: uint16(4),
	2852: uint16(anon_sym_LBRACK),
	2853: uint16(anon_sym_LTa),
	2854: uint16(sym_function_link),
	2855: uint16(284),
	2856: uint16(9),
	2857: uint16(sym_tag_name_with_argument),
	2858: uint16(sym_tag_name_with_multiple_arguments),
	2859: uint16(sym_tag_name_with_types),
	2860: uint16(sym_tag_name_with_self_types),
	2861: uint16(sym_tag_name_with_type),
	2862: uint16(sym_tag_name),
	2863: uint16(anon_sym_BSLASHa),
	2864: uint16(anon_sym_BSLASHc),
	2865: uint16(sym__text),
	2866: uint16(2),
	2867: uint16(274),
	2868: uint16(4),
	2870: uint16(anon_sym_LBRACK),
	2871: uint16(anon_sym_LTa),
	2872: uint16(sym_function_link),
	2873: uint16(272),
	2874: uint16(9),
	2875: uint16(sym_tag_name_with_argument),
	2876: uint16(sym_tag_name_with_multiple_arguments),
	2877: uint16(sym_tag_name_with_types),
	2878: uint16(sym_tag_name_with_self_types),
	2879: uint16(sym_tag_name_with_type),
	2880: uint16(sym_tag_name),
	2881: uint16(anon_sym_BSLASHa),
	2882: uint16(anon_sym_BSLASHc),
	2883: uint16(sym__text),
	2884: uint16(2),
	2885: uint16(264),
	2886: uint16(4),
	2888: uint16(anon_sym_LBRACK),
	2889: uint16(anon_sym_LTa),
	2890: uint16(sym_function_link),
	2891: uint16(262),
	2892: uint16(9),
	2893: uint16(sym_tag_name_with_argument),
	2894: uint16(sym_tag_name_with_multiple_arguments),
	2895: uint16(sym_tag_name_with_types),
	2896: uint16(sym_tag_name_with_self_types),
	2897: uint16(sym_tag_name_with_type),
	2898: uint16(sym_tag_name),
	2899: uint16(anon_sym_BSLASHa),
	2900: uint16(anon_sym_BSLASHc),
	2901: uint16(sym__text),
	2902: uint16(7),
	2903: uint16(370),
	2904: uint16(1),
	2905: uint16(anon_sym_COMMA),
	2906: uint16(374),
	2907: uint16(1),
	2908: uint16(aux_sym_identifier_token1),
	2909: uint16(380),
	2910: uint16(1),
	2911: uint16(sym_function_link),
	2912: uint16(107),
	2913: uint16(1),
	2914: uint16(aux_sym_tag_repeat2),
	2915: uint16(132),
	2916: uint16(1),
	2917: uint16(sym_identifier),
	2918: uint16(243),
	2919: uint16(2),
	2921: uint16(anon_sym_LBRACK),
	2922: uint16(241),
	2923: uint16(6),
	2924: uint16(sym_tag_name_with_argument),
	2925: uint16(sym_tag_name_with_multiple_arguments),
	2926: uint16(sym_tag_name_with_types),
	2927: uint16(sym_tag_name_with_self_types),
	2928: uint16(sym_tag_name_with_type),
	2929: uint16(sym_tag_name),
	2930: uint16(2),
	2931: uint16(167),
	2932: uint16(4),
	2934: uint16(anon_sym_LBRACK),
	2935: uint16(anon_sym_LTa),
	2936: uint16(sym_function_link),
	2937: uint16(169),
	2938: uint16(9),
	2939: uint16(sym_tag_name_with_argument),
	2940: uint16(sym_tag_name_with_multiple_arguments),
	2941: uint16(sym_tag_name_with_types),
	2942: uint16(sym_tag_name_with_self_types),
	2943: uint16(sym_tag_name_with_type),
	2944: uint16(sym_tag_name),
	2945: uint16(anon_sym_BSLASHa),
	2946: uint16(anon_sym_BSLASHc),
	2947: uint16(sym__text),
	2948: uint16(2),
	2949: uint16(199),
	2950: uint16(4),
	2952: uint16(anon_sym_LBRACK),
	2953: uint16(anon_sym_LTa),
	2954: uint16(sym_function_link),
	2955: uint16(201),
	2956: uint16(9),
	2957: uint16(sym_tag_name_with_argument),
	2958: uint16(sym_tag_name_with_multiple_arguments),
	2959: uint16(sym_tag_name_with_types),
	2960: uint16(sym_tag_name_with_self_types),
	2961: uint16(sym_tag_name_with_type),
	2962: uint16(sym_tag_name),
	2963: uint16(anon_sym_BSLASHa),
	2964: uint16(anon_sym_BSLASHc),
	2965: uint16(sym__text),
	2966: uint16(7),
	2967: uint16(370),
	2968: uint16(1),
	2969: uint16(anon_sym_COMMA),
	2970: uint16(374),
	2971: uint16(1),
	2972: uint16(aux_sym_identifier_token1),
	2973: uint16(382),
	2974: uint16(1),
	2975: uint16(sym_function_link),
	2976: uint16(107),
	2977: uint16(1),
	2978: uint16(aux_sym_tag_repeat2),
	2979: uint16(137),
	2980: uint16(1),
	2981: uint16(sym_identifier),
	2982: uint16(231),
	2983: uint16(2),
	2985: uint16(anon_sym_LBRACK),
	2986: uint16(229),
	2987: uint16(6),
	2988: uint16(sym_tag_name_with_argument),
	2989: uint16(sym_tag_name_with_multiple_arguments),
	2990: uint16(sym_tag_name_with_types),
	2991: uint16(sym_tag_name_with_self_types),
	2992: uint16(sym_tag_name_with_type),
	2993: uint16(sym_tag_name),
	2994: uint16(2),
	2995: uint16(278),
	2996: uint16(4),
	2998: uint16(anon_sym_LBRACK),
	2999: uint16(anon_sym_LTa),
	3000: uint16(sym_function_link),
	3001: uint16(276),
	3002: uint16(9),
	3003: uint16(sym_tag_name_with_argument),
	3004: uint16(sym_tag_name_with_multiple_arguments),
	3005: uint16(sym_tag_name_with_types),
	3006: uint16(sym_tag_name_with_self_types),
	3007: uint16(sym_tag_name_with_type),
	3008: uint16(sym_tag_name),
	3009: uint16(anon_sym_BSLASHa),
	3010: uint16(anon_sym_BSLASHc),
	3011: uint16(sym__text),
	3012: uint16(2),
	3013: uint16(290),
	3014: uint16(4),
	3016: uint16(anon_sym_LBRACK),
	3017: uint16(anon_sym_LTa),
	3018: uint16(sym_function_link),
	3019: uint16(288),
	3020: uint16(9),
	3021: uint16(sym_tag_name_with_argument),
	3022: uint16(sym_tag_name_with_multiple_arguments),
	3023: uint16(sym_tag_name_with_types),
	3024: uint16(sym_tag_name_with_self_types),
	3025: uint16(sym_tag_name_with_type),
	3026: uint16(sym_tag_name),
	3027: uint16(anon_sym_BSLASHa),
	3028: uint16(anon_sym_BSLASHc),
	3029: uint16(sym__text),
	3030: uint16(2),
	3031: uint16(215),
	3032: uint16(4),
	3034: uint16(anon_sym_LBRACK),
	3035: uint16(anon_sym_LTa),
	3036: uint16(sym_function_link),
	3037: uint16(217),
	3038: uint16(9),
	3039: uint16(sym_tag_name_with_argument),
	3040: uint16(sym_tag_name_with_multiple_arguments),
	3041: uint16(sym_tag_name_with_types),
	3042: uint16(sym_tag_name_with_self_types),
	3043: uint16(sym_tag_name_with_type),
	3044: uint16(sym_tag_name),
	3045: uint16(anon_sym_BSLASHa),
	3046: uint16(anon_sym_BSLASHc),
	3047: uint16(sym__text),
	3048: uint16(2),
	3049: uint16(270),
	3050: uint16(4),
	3052: uint16(anon_sym_LBRACK),
	3053: uint16(anon_sym_LTa),
	3054: uint16(sym_function_link),
	3055: uint16(268),
	3056: uint16(9),
	3057: uint16(sym_tag_name_with_argument),
	3058: uint16(sym_tag_name_with_multiple_arguments),
	3059: uint16(sym_tag_name_with_types),
	3060: uint16(sym_tag_name_with_self_types),
	3061: uint16(sym_tag_name_with_type),
	3062: uint16(sym_tag_name),
	3063: uint16(anon_sym_BSLASHa),
	3064: uint16(anon_sym_BSLASHc),
	3065: uint16(sym__text),
	3066: uint16(2),
	3067: uint16(386),
	3068: uint16(3),
	3069: uint16(sym_code_block_start),
	3070: uint16(anon_sym_LBRACK),
	3071: uint16(anon_sym_SLASH),
	3072: uint16(384),
	3073: uint16(9),
	3074: uint16(sym_tag_name_with_argument),
	3075: uint16(sym_tag_name_with_multiple_arguments),
	3076: uint16(sym_tag_name_with_types),
	3077: uint16(sym_tag_name_with_self_types),
	3078: uint16(sym_tag_name_with_type),
	3079: uint16(sym_tag_name),
	3080: uint16(anon_sym_ATcode),
	3081: uint16(anon_sym_STAR_SLASH),
	3082: uint16(sym__text_line),
	3083: uint16(2),
	3084: uint16(105),
	3085: uint16(3),
	3086: uint16(sym_code_block_start),
	3087: uint16(anon_sym_LBRACK),
	3088: uint16(anon_sym_SLASH),
	3089: uint16(103),
	3090: uint16(9),
	3091: uint16(sym_tag_name_with_argument),
	3092: uint16(sym_tag_name_with_multiple_arguments),
	3093: uint16(sym_tag_name_with_types),
	3094: uint16(sym_tag_name_with_self_types),
	3095: uint16(sym_tag_name_with_type),
	3096: uint16(sym_tag_name),
	3097: uint16(anon_sym_ATcode),
	3098: uint16(anon_sym_STAR_SLASH),
	3099: uint16(sym__text_line),
	3100: uint16(3),
	3101: uint16(388),
	3102: uint16(1),
	3103: uint16(aux_sym_tag_token1),
	3104: uint16(296),
	3105: uint16(4),
	3107: uint16(anon_sym_COMMA),
	3108: uint16(anon_sym_LBRACK),
	3109: uint16(sym_function_link),
	3110: uint16(300),
	3111: uint16(7),
	3112: uint16(sym_tag_name_with_argument),
	3113: uint16(sym_tag_name_with_multiple_arguments),
	3114: uint16(sym_tag_name_with_types),
	3115: uint16(sym_tag_name_with_self_types),
	3116: uint16(sym_tag_name_with_type),
	3117: uint16(sym_tag_name),
	3118: uint16(aux_sym_identifier_token1),
	3119: uint16(2),
	3120: uint16(392),
	3121: uint16(3),
	3122: uint16(sym_code_block_start),
	3123: uint16(anon_sym_LBRACK),
	3124: uint16(anon_sym_SLASH),
	3125: uint16(390),
	3126: uint16(9),
	3127: uint16(sym_tag_name_with_argument),
	3128: uint16(sym_tag_name_with_multiple_arguments),
	3129: uint16(sym_tag_name_with_types),
	3130: uint16(sym_tag_name_with_self_types),
	3131: uint16(sym_tag_name_with_type),
	3132: uint16(sym_tag_name),
	3133: uint16(anon_sym_ATcode),
	3134: uint16(anon_sym_STAR_SLASH),
	3135: uint16(sym__text_line),
	3136: uint16(2),
	3137: uint16(396),
	3138: uint16(3),
	3139: uint16(sym_code_block_start),
	3140: uint16(anon_sym_LBRACK),
	3141: uint16(anon_sym_SLASH),
	3142: uint16(394),
	3143: uint16(9),
	3144: uint16(sym_tag_name_with_argument),
	3145: uint16(sym_tag_name_with_multiple_arguments),
	3146: uint16(sym_tag_name_with_types),
	3147: uint16(sym_tag_name_with_self_types),
	3148: uint16(sym_tag_name_with_type),
	3149: uint16(sym_tag_name),
	3150: uint16(anon_sym_ATcode),
	3151: uint16(anon_sym_STAR_SLASH),
	3152: uint16(sym__text_line),
	3153: uint16(2),
	3154: uint16(386),
	3155: uint16(3),
	3156: uint16(sym_code_block_start),
	3157: uint16(anon_sym_LBRACK),
	3158: uint16(anon_sym_SLASH),
	3159: uint16(384),
	3160: uint16(9),
	3161: uint16(sym_tag_name_with_argument),
	3162: uint16(sym_tag_name_with_multiple_arguments),
	3163: uint16(sym_tag_name_with_types),
	3164: uint16(sym_tag_name_with_self_types),
	3165: uint16(sym_tag_name_with_type),
	3166: uint16(sym_tag_name),
	3167: uint16(anon_sym_ATcode),
	3168: uint16(anon_sym_STAR_SLASH),
	3169: uint16(sym__text_line),
	3170: uint16(2),
	3171: uint16(400),
	3172: uint16(3),
	3173: uint16(sym_code_block_start),
	3174: uint16(anon_sym_LBRACK),
	3175: uint16(anon_sym_SLASH),
	3176: uint16(398),
	3177: uint16(9),
	3178: uint16(sym_tag_name_with_argument),
	3179: uint16(sym_tag_name_with_multiple_arguments),
	3180: uint16(sym_tag_name_with_types),
	3181: uint16(sym_tag_name_with_self_types),
	3182: uint16(sym_tag_name_with_type),
	3183: uint16(sym_tag_name),
	3184: uint16(anon_sym_ATcode),
	3185: uint16(anon_sym_STAR_SLASH),
	3186: uint16(sym__text_line),
	3187: uint16(4),
	3188: uint16(402),
	3189: uint16(1),
	3190: uint16(anon_sym_COMMA),
	3191: uint16(107),
	3192: uint16(1),
	3193: uint16(aux_sym_tag_repeat2),
	3194: uint16(307),
	3195: uint16(3),
	3197: uint16(anon_sym_LBRACK),
	3198: uint16(sym_function_link),
	3199: uint16(305),
	3200: uint16(7),
	3201: uint16(sym_tag_name_with_argument),
	3202: uint16(sym_tag_name_with_multiple_arguments),
	3203: uint16(sym_tag_name_with_types),
	3204: uint16(sym_tag_name_with_self_types),
	3205: uint16(sym_tag_name_with_type),
	3206: uint16(sym_tag_name),
	3207: uint16(aux_sym_identifier_token1),
	3208: uint16(2),
	3209: uint16(407),
	3210: uint16(3),
	3211: uint16(sym_code_block_start),
	3212: uint16(anon_sym_LBRACK),
	3213: uint16(anon_sym_SLASH),
	3214: uint16(405),
	3215: uint16(9),
	3216: uint16(sym_tag_name_with_argument),
	3217: uint16(sym_tag_name_with_multiple_arguments),
	3218: uint16(sym_tag_name_with_types),
	3219: uint16(sym_tag_name_with_self_types),
	3220: uint16(sym_tag_name_with_type),
	3221: uint16(sym_tag_name),
	3222: uint16(anon_sym_ATcode),
	3223: uint16(anon_sym_STAR_SLASH),
	3224: uint16(sym__text_line),
	3225: uint16(2),
	3226: uint16(411),
	3227: uint16(3),
	3228: uint16(sym_code_block_start),
	3229: uint16(anon_sym_LBRACK),
	3230: uint16(anon_sym_SLASH),
	3231: uint16(409),
	3232: uint16(9),
	3233: uint16(sym_tag_name_with_argument),
	3234: uint16(sym_tag_name_with_multiple_arguments),
	3235: uint16(sym_tag_name_with_types),
	3236: uint16(sym_tag_name_with_self_types),
	3237: uint16(sym_tag_name_with_type),
	3238: uint16(sym_tag_name),
	3239: uint16(anon_sym_ATcode),
	3240: uint16(anon_sym_STAR_SLASH),
	3241: uint16(sym__text_line),
	3242: uint16(2),
	3243: uint16(167),
	3244: uint16(5),
	3246: uint16(anon_sym_COMMA),
	3247: uint16(aux_sym_tag_token1),
	3248: uint16(anon_sym_LBRACK),
	3249: uint16(sym_function_link),
	3250: uint16(169),
	3251: uint16(7),
	3252: uint16(sym_tag_name_with_argument),
	3253: uint16(sym_tag_name_with_multiple_arguments),
	3254: uint16(sym_tag_name_with_types),
	3255: uint16(sym_tag_name_with_self_types),
	3256: uint16(sym_tag_name_with_type),
	3257: uint16(sym_tag_name),
	3258: uint16(aux_sym_identifier_token1),
	3259: uint16(2),
	3260: uint16(89),
	3261: uint16(3),
	3262: uint16(sym_code_block_start),
	3263: uint16(anon_sym_LBRACK),
	3264: uint16(anon_sym_SLASH),
	3265: uint16(87),
	3266: uint16(9),
	3267: uint16(sym_tag_name_with_argument),
	3268: uint16(sym_tag_name_with_multiple_arguments),
	3269: uint16(sym_tag_name_with_types),
	3270: uint16(sym_tag_name_with_self_types),
	3271: uint16(sym_tag_name_with_type),
	3272: uint16(sym_tag_name),
	3273: uint16(anon_sym_ATcode),
	3274: uint16(anon_sym_STAR_SLASH),
	3275: uint16(sym__text_line),
	3276: uint16(2),
	3277: uint16(415),
	3278: uint16(3),
	3279: uint16(sym_code_block_start),
	3280: uint16(anon_sym_LBRACK),
	3281: uint16(anon_sym_SLASH),
	3282: uint16(413),
	3283: uint16(9),
	3284: uint16(sym_tag_name_with_argument),
	3285: uint16(sym_tag_name_with_multiple_arguments),
	3286: uint16(sym_tag_name_with_types),
	3287: uint16(sym_tag_name_with_self_types),
	3288: uint16(sym_tag_name_with_type),
	3289: uint16(sym_tag_name),
	3290: uint16(anon_sym_ATcode),
	3291: uint16(anon_sym_STAR_SLASH),
	3292: uint16(sym__text_line),
	3293: uint16(2),
	3294: uint16(419),
	3295: uint16(3),
	3296: uint16(sym_code_block_start),
	3297: uint16(anon_sym_LBRACK),
	3298: uint16(anon_sym_SLASH),
	3299: uint16(417),
	3300: uint16(9),
	3301: uint16(sym_tag_name_with_argument),
	3302: uint16(sym_tag_name_with_multiple_arguments),
	3303: uint16(sym_tag_name_with_types),
	3304: uint16(sym_tag_name_with_self_types),
	3305: uint16(sym_tag_name_with_type),
	3306: uint16(sym_tag_name),
	3307: uint16(anon_sym_ATcode),
	3308: uint16(anon_sym_STAR_SLASH),
	3309: uint16(sym__text_line),
	3310: uint16(2),
	3311: uint16(423),
	3312: uint16(3),
	3313: uint16(sym_code_block_start),
	3314: uint16(anon_sym_LBRACK),
	3315: uint16(anon_sym_SLASH),
	3316: uint16(421),
	3317: uint16(9),
	3318: uint16(sym_tag_name_with_argument),
	3319: uint16(sym_tag_name_with_multiple_arguments),
	3320: uint16(sym_tag_name_with_types),
	3321: uint16(sym_tag_name_with_self_types),
	3322: uint16(sym_tag_name_with_type),
	3323: uint16(sym_tag_name),
	3324: uint16(anon_sym_ATcode),
	3325: uint16(anon_sym_STAR_SLASH),
	3326: uint16(sym__text_line),
	3327: uint16(2),
	3328: uint16(109),
	3329: uint16(3),
	3330: uint16(sym_code_block_start),
	3331: uint16(anon_sym_LBRACK),
	3332: uint16(anon_sym_SLASH),
	3333: uint16(107),
	3334: uint16(9),
	3335: uint16(sym_tag_name_with_argument),
	3336: uint16(sym_tag_name_with_multiple_arguments),
	3337: uint16(sym_tag_name_with_types),
	3338: uint16(sym_tag_name_with_self_types),
	3339: uint16(sym_tag_name_with_type),
	3340: uint16(sym_tag_name),
	3341: uint16(anon_sym_ATcode),
	3342: uint16(anon_sym_STAR_SLASH),
	3343: uint16(sym__text_line),
	3344: uint16(2),
	3345: uint16(167),
	3346: uint16(3),
	3347: uint16(sym_code_block_start),
	3348: uint16(anon_sym_LBRACK),
	3349: uint16(anon_sym_SLASH),
	3350: uint16(169),
	3351: uint16(9),
	3352: uint16(sym_tag_name_with_argument),
	3353: uint16(sym_tag_name_with_multiple_arguments),
	3354: uint16(sym_tag_name_with_types),
	3355: uint16(sym_tag_name_with_self_types),
	3356: uint16(sym_tag_name_with_type),
	3357: uint16(sym_tag_name),
	3358: uint16(anon_sym_ATcode),
	3359: uint16(anon_sym_STAR_SLASH),
	3360: uint16(sym__text_line),
	3361: uint16(2),
	3362: uint16(400),
	3363: uint16(3),
	3364: uint16(sym_code_block_start),
	3365: uint16(anon_sym_LBRACK),
	3366: uint16(anon_sym_SLASH),
	3367: uint16(398),
	3368: uint16(9),
	3369: uint16(sym_tag_name_with_argument),
	3370: uint16(sym_tag_name_with_multiple_arguments),
	3371: uint16(sym_tag_name_with_types),
	3372: uint16(sym_tag_name_with_self_types),
	3373: uint16(sym_tag_name_with_type),
	3374: uint16(sym_tag_name),
	3375: uint16(anon_sym_ATcode),
	3376: uint16(anon_sym_STAR_SLASH),
	3377: uint16(sym__text_line),
	3378: uint16(2),
	3379: uint16(93),
	3380: uint16(3),
	3381: uint16(sym_code_block_start),
	3382: uint16(anon_sym_LBRACK),
	3383: uint16(anon_sym_SLASH),
	3384: uint16(91),
	3385: uint16(9),
	3386: uint16(sym_tag_name_with_argument),
	3387: uint16(sym_tag_name_with_multiple_arguments),
	3388: uint16(sym_tag_name_with_types),
	3389: uint16(sym_tag_name_with_self_types),
	3390: uint16(sym_tag_name_with_type),
	3391: uint16(sym_tag_name),
	3392: uint16(anon_sym_ATcode),
	3393: uint16(anon_sym_STAR_SLASH),
	3394: uint16(sym__text_line),
	3395: uint16(10),
	3396: uint16(63),
	3397: uint16(1),
	3398: uint16(sym_tag_name_with_argument),
	3399: uint16(65),
	3400: uint16(1),
	3401: uint16(sym_tag_name_with_multiple_arguments),
	3402: uint16(67),
	3403: uint16(1),
	3404: uint16(sym_tag_name_with_types),
	3405: uint16(69),
	3406: uint16(1),
	3407: uint16(sym_tag_name_with_self_types),
	3408: uint16(71),
	3409: uint16(1),
	3410: uint16(sym_tag_name_with_type),
	3411: uint16(73),
	3412: uint16(1),
	3413: uint16(sym_tag_name),
	3414: uint16(75),
	3415: uint16(1),
	3416: uint16(anon_sym_LBRACK),
	3417: uint16(146),
	3418: uint16(1),
	3420: uint16(37),
	3421: uint16(1),
	3422: uint16(sym_storageclass),
	3423: uint16(125),
	3424: uint16(2),
	3425: uint16(sym_tag),
	3426: uint16(aux_sym_document_repeat2),
	3427: uint16(10),
	3428: uint16(63),
	3429: uint16(1),
	3430: uint16(sym_tag_name_with_argument),
	3431: uint16(65),
	3432: uint16(1),
	3433: uint16(sym_tag_name_with_multiple_arguments),
	3434: uint16(67),
	3435: uint16(1),
	3436: uint16(sym_tag_name_with_types),
	3437: uint16(69),
	3438: uint16(1),
	3439: uint16(sym_tag_name_with_self_types),
	3440: uint16(71),
	3441: uint16(1),
	3442: uint16(sym_tag_name_with_type),
	3443: uint16(73),
	3444: uint16(1),
	3445: uint16(sym_tag_name),
	3446: uint16(75),
	3447: uint16(1),
	3448: uint16(anon_sym_LBRACK),
	3449: uint16(146),
	3450: uint16(1),
	3452: uint16(37),
	3453: uint16(1),
	3454: uint16(sym_storageclass),
	3455: uint16(123),
	3456: uint16(2),
	3457: uint16(sym_tag),
	3458: uint16(aux_sym_document_repeat2),
	3459: uint16(2),
	3460: uint16(353),
	3461: uint16(4),
	3463: uint16(anon_sym_COMMA),
	3464: uint16(anon_sym_LBRACK),
	3465: uint16(sym_function_link),
	3466: uint16(355),
	3467: uint16(7),
	3468: uint16(sym_tag_name_with_argument),
	3469: uint16(sym_tag_name_with_multiple_arguments),
	3470: uint16(sym_tag_name_with_types),
	3471: uint16(sym_tag_name_with_self_types),
	3472: uint16(sym_tag_name_with_type),
	3473: uint16(sym_tag_name),
	3474: uint16(aux_sym_identifier_token1),
	3475: uint16(10),
	3476: uint16(63),
	3477: uint16(1),
	3478: uint16(sym_tag_name_with_argument),
	3479: uint16(65),
	3480: uint16(1),
	3481: uint16(sym_tag_name_with_multiple_arguments),
	3482: uint16(67),
	3483: uint16(1),
	3484: uint16(sym_tag_name_with_types),
	3485: uint16(69),
	3486: uint16(1),
	3487: uint16(sym_tag_name_with_self_types),
	3488: uint16(71),
	3489: uint16(1),
	3490: uint16(sym_tag_name_with_type),
	3491: uint16(73),
	3492: uint16(1),
	3493: uint16(sym_tag_name),
	3494: uint16(75),
	3495: uint16(1),
	3496: uint16(anon_sym_LBRACK),
	3497: uint16(425),
	3498: uint16(1),
	3500: uint16(37),
	3501: uint16(1),
	3502: uint16(sym_storageclass),
	3503: uint16(124),
	3504: uint16(2),
	3505: uint16(sym_tag),
	3506: uint16(aux_sym_document_repeat2),
	3507: uint16(10),
	3508: uint16(427),
	3509: uint16(1),
	3511: uint16(429),
	3512: uint16(1),
	3513: uint16(sym_tag_name_with_argument),
	3514: uint16(432),
	3515: uint16(1),
	3516: uint16(sym_tag_name_with_multiple_arguments),
	3517: uint16(435),
	3518: uint16(1),
	3519: uint16(sym_tag_name_with_types),
	3520: uint16(438),
	3521: uint16(1),
	3522: uint16(sym_tag_name_with_self_types),
	3523: uint16(441),
	3524: uint16(1),
	3525: uint16(sym_tag_name_with_type),
	3526: uint16(444),
	3527: uint16(1),
	3528: uint16(sym_tag_name),
	3529: uint16(447),
	3530: uint16(1),
	3531: uint16(anon_sym_LBRACK),
	3532: uint16(37),
	3533: uint16(1),
	3534: uint16(sym_storageclass),
	3535: uint16(123),
	3536: uint16(2),
	3537: uint16(sym_tag),
	3538: uint16(aux_sym_document_repeat2),
	3539: uint16(10),
	3540: uint16(63),
	3541: uint16(1),
	3542: uint16(sym_tag_name_with_argument),
	3543: uint16(65),
	3544: uint16(1),
	3545: uint16(sym_tag_name_with_multiple_arguments),
	3546: uint16(67),
	3547: uint16(1),
	3548: uint16(sym_tag_name_with_types),
	3549: uint16(69),
	3550: uint16(1),
	3551: uint16(sym_tag_name_with_self_types),
	3552: uint16(71),
	3553: uint16(1),
	3554: uint16(sym_tag_name_with_type),
	3555: uint16(73),
	3556: uint16(1),
	3557: uint16(sym_tag_name),
	3558: uint16(75),
	3559: uint16(1),
	3560: uint16(anon_sym_LBRACK),
	3561: uint16(450),
	3562: uint16(1),
	3564: uint16(37),
	3565: uint16(1),
	3566: uint16(sym_storageclass),
	3567: uint16(123),
	3568: uint16(2),
	3569: uint16(sym_tag),
	3570: uint16(aux_sym_document_repeat2),
	3571: uint16(10),
	3572: uint16(63),
	3573: uint16(1),
	3574: uint16(sym_tag_name_with_argument),
	3575: uint16(65),
	3576: uint16(1),
	3577: uint16(sym_tag_name_with_multiple_arguments),
	3578: uint16(67),
	3579: uint16(1),
	3580: uint16(sym_tag_name_with_types),
	3581: uint16(69),
	3582: uint16(1),
	3583: uint16(sym_tag_name_with_self_types),
	3584: uint16(71),
	3585: uint16(1),
	3586: uint16(sym_tag_name_with_type),
	3587: uint16(73),
	3588: uint16(1),
	3589: uint16(sym_tag_name),
	3590: uint16(75),
	3591: uint16(1),
	3592: uint16(anon_sym_LBRACK),
	3593: uint16(425),
	3594: uint16(1),
	3596: uint16(37),
	3597: uint16(1),
	3598: uint16(sym_storageclass),
	3599: uint16(123),
	3600: uint16(2),
	3601: uint16(sym_tag),
	3602: uint16(aux_sym_document_repeat2),
	3603: uint16(2),
	3604: uint16(419),
	3605: uint16(2),
	3607: uint16(anon_sym_LBRACK),
	3608: uint16(417),
	3609: uint16(6),
	3610: uint16(sym_tag_name_with_argument),
	3611: uint16(sym_tag_name_with_multiple_arguments),
	3612: uint16(sym_tag_name_with_types),
	3613: uint16(sym_tag_name_with_self_types),
	3614: uint16(sym_tag_name_with_type),
	3615: uint16(sym_tag_name),
	3616: uint16(2),
	3617: uint16(93),
	3618: uint16(2),
	3620: uint16(anon_sym_LBRACK),
	3621: uint16(91),
	3622: uint16(6),
	3623: uint16(sym_tag_name_with_argument),
	3624: uint16(sym_tag_name_with_multiple_arguments),
	3625: uint16(sym_tag_name_with_types),
	3626: uint16(sym_tag_name_with_self_types),
	3627: uint16(sym_tag_name_with_type),
	3628: uint16(sym_tag_name),
	3629: uint16(2),
	3630: uint16(89),
	3631: uint16(2),
	3633: uint16(anon_sym_LBRACK),
	3634: uint16(87),
	3635: uint16(6),
	3636: uint16(sym_tag_name_with_argument),
	3637: uint16(sym_tag_name_with_multiple_arguments),
	3638: uint16(sym_tag_name_with_types),
	3639: uint16(sym_tag_name_with_self_types),
	3640: uint16(sym_tag_name_with_type),
	3641: uint16(sym_tag_name),
	3642: uint16(2),
	3643: uint16(109),
	3644: uint16(2),
	3646: uint16(anon_sym_LBRACK),
	3647: uint16(107),
	3648: uint16(6),
	3649: uint16(sym_tag_name_with_argument),
	3650: uint16(sym_tag_name_with_multiple_arguments),
	3651: uint16(sym_tag_name_with_types),
	3652: uint16(sym_tag_name_with_self_types),
	3653: uint16(sym_tag_name_with_type),
	3654: uint16(sym_tag_name),
	3655: uint16(6),
	3656: uint16(452),
	3657: uint16(1),
	3658: uint16(aux_sym_identifier_token1),
	3659: uint16(454),
	3660: uint16(1),
	3661: uint16(anon_sym_TILDE),
	3662: uint16(456),
	3663: uint16(1),
	3664: uint16(anon_sym_LBRACK),
	3665: uint16(17),
	3666: uint16(1),
	3667: uint16(sym_identifier),
	3668: uint16(142),
	3669: uint16(1),
	3670: uint16(sym_storageclass),
	3671: uint16(6),
	3672: uint16(3),
	3673: uint16(sym__expression),
	3674: uint16(sym_qualified_identifier),
	3675: uint16(sym_function),
	3676: uint16(2),
	3677: uint16(105),
	3678: uint16(2),
	3680: uint16(anon_sym_LBRACK),
	3681: uint16(103),
	3682: uint16(6),
	3683: uint16(sym_tag_name_with_argument),
	3684: uint16(sym_tag_name_with_multiple_arguments),
	3685: uint16(sym_tag_name_with_types),
	3686: uint16(sym_tag_name_with_self_types),
	3687: uint16(sym_tag_name_with_type),
	3688: uint16(sym_tag_name),
	3689: uint16(2),
	3690: uint16(386),
	3691: uint16(2),
	3693: uint16(anon_sym_LBRACK),
	3694: uint16(384),
	3695: uint16(6),
	3696: uint16(sym_tag_name_with_argument),
	3697: uint16(sym_tag_name_with_multiple_arguments),
	3698: uint16(sym_tag_name_with_types),
	3699: uint16(sym_tag_name_with_self_types),
	3700: uint16(sym_tag_name_with_type),
	3701: uint16(sym_tag_name),
	3702: uint16(2),
	3703: uint16(386),
	3704: uint16(2),
	3706: uint16(anon_sym_LBRACK),
	3707: uint16(384),
	3708: uint16(6),
	3709: uint16(sym_tag_name_with_argument),
	3710: uint16(sym_tag_name_with_multiple_arguments),
	3711: uint16(sym_tag_name_with_types),
	3712: uint16(sym_tag_name_with_self_types),
	3713: uint16(sym_tag_name_with_type),
	3714: uint16(sym_tag_name),
	3715: uint16(2),
	3716: uint16(392),
	3717: uint16(2),
	3719: uint16(anon_sym_LBRACK),
	3720: uint16(390),
	3721: uint16(6),
	3722: uint16(sym_tag_name_with_argument),
	3723: uint16(sym_tag_name_with_multiple_arguments),
	3724: uint16(sym_tag_name_with_types),
	3725: uint16(sym_tag_name_with_self_types),
	3726: uint16(sym_tag_name_with_type),
	3727: uint16(sym_tag_name),
	3728: uint16(6),
	3729: uint16(456),
	3730: uint16(1),
	3731: uint16(anon_sym_LBRACK),
	3732: uint16(458),
	3733: uint16(1),
	3734: uint16(aux_sym_identifier_token1),
	3735: uint16(460),
	3736: uint16(1),
	3737: uint16(anon_sym_TILDE),
	3738: uint16(64),
	3739: uint16(1),
	3740: uint16(sym_identifier),
	3741: uint16(140),
	3742: uint16(1),
	3743: uint16(sym_storageclass),
	3744: uint16(19),
	3745: uint16(3),
	3746: uint16(sym__expression),
	3747: uint16(sym_qualified_identifier),
	3748: uint16(sym_function),
	3749: uint16(2),
	3750: uint16(423),
	3751: uint16(2),
	3753: uint16(anon_sym_LBRACK),
	3754: uint16(421),
	3755: uint16(6),
	3756: uint16(sym_tag_name_with_argument),
	3757: uint16(sym_tag_name_with_multiple_arguments),
	3758: uint16(sym_tag_name_with_types),
	3759: uint16(sym_tag_name_with_self_types),
	3760: uint16(sym_tag_name_with_type),
	3761: uint16(sym_tag_name),
	3762: uint16(2),
	3763: uint16(400),
	3764: uint16(2),
	3766: uint16(anon_sym_LBRACK),
	3767: uint16(398),
	3768: uint16(6),
	3769: uint16(sym_tag_name_with_argument),
	3770: uint16(sym_tag_name_with_multiple_arguments),
	3771: uint16(sym_tag_name_with_types),
	3772: uint16(sym_tag_name_with_self_types),
	3773: uint16(sym_tag_name_with_type),
	3774: uint16(sym_tag_name),
	3775: uint16(2),
	3776: uint16(400),
	3777: uint16(2),
	3779: uint16(anon_sym_LBRACK),
	3780: uint16(398),
	3781: uint16(6),
	3782: uint16(sym_tag_name_with_argument),
	3783: uint16(sym_tag_name_with_multiple_arguments),
	3784: uint16(sym_tag_name_with_types),
	3785: uint16(sym_tag_name_with_self_types),
	3786: uint16(sym_tag_name_with_type),
	3787: uint16(sym_tag_name),
	3788: uint16(2),
	3789: uint16(411),
	3790: uint16(2),
	3792: uint16(anon_sym_LBRACK),
	3793: uint16(409),
	3794: uint16(6),
	3795: uint16(sym_tag_name_with_argument),
	3796: uint16(sym_tag_name_with_multiple_arguments),
	3797: uint16(sym_tag_name_with_types),
	3798: uint16(sym_tag_name_with_self_types),
	3799: uint16(sym_tag_name_with_type),
	3800: uint16(sym_tag_name),
	3801: uint16(4),
	3802: uint16(458),
	3803: uint16(1),
	3804: uint16(aux_sym_identifier_token1),
	3805: uint16(460),
	3806: uint16(1),
	3807: uint16(anon_sym_TILDE),
	3808: uint16(64),
	3809: uint16(1),
	3810: uint16(sym_identifier),
	3811: uint16(26),
	3812: uint16(3),
	3813: uint16(sym__expression),
	3814: uint16(sym_qualified_identifier),
	3815: uint16(sym_function),
	3816: uint16(4),
	3817: uint16(452),
	3818: uint16(1),
	3819: uint16(aux_sym_identifier_token1),
	3820: uint16(454),
	3821: uint16(1),
	3822: uint16(anon_sym_TILDE),
	3823: uint16(17),
	3824: uint16(1),
	3825: uint16(sym_identifier),
	3826: uint16(43),
	3827: uint16(3),
	3828: uint16(sym__expression),
	3829: uint16(sym_qualified_identifier),
	3830: uint16(sym_function),
	3831: uint16(4),
	3832: uint16(452),
	3833: uint16(1),
	3834: uint16(aux_sym_identifier_token1),
	3835: uint16(454),
	3836: uint16(1),
	3837: uint16(anon_sym_TILDE),
	3838: uint16(17),
	3839: uint16(1),
	3840: uint16(sym_identifier),
	3841: uint16(11),
	3842: uint16(3),
	3843: uint16(sym__expression),
	3844: uint16(sym_qualified_identifier),
	3845: uint16(sym_function),
	3846: uint16(4),
	3847: uint16(458),
	3848: uint16(1),
	3849: uint16(aux_sym_identifier_token1),
	3850: uint16(460),
	3851: uint16(1),
	3852: uint16(anon_sym_TILDE),
	3853: uint16(64),
	3854: uint16(1),
	3855: uint16(sym_identifier),
	3856: uint16(85),
	3857: uint16(3),
	3858: uint16(sym__expression),
	3859: uint16(sym_qualified_identifier),
	3860: uint16(sym_function),
	3861: uint16(3),
	3862: uint16(35),
	3863: uint16(1),
	3864: uint16(aux_sym_brief_description_repeat1),
	3865: uint16(56),
	3866: uint16(1),
	3867: uint16(sym_brief_description),
	3868: uint16(462),
	3869: uint16(2),
	3870: uint16(sym_brief_text),
	3871: uint16(sym_tag_name),
	3872: uint16(3),
	3873: uint16(81),
	3874: uint16(1),
	3875: uint16(aux_sym_brief_description_repeat1),
	3876: uint16(90),
	3877: uint16(1),
	3878: uint16(sym_brief_description),
	3879: uint16(464),
	3880: uint16(2),
	3881: uint16(sym_brief_text),
	3882: uint16(sym_tag_name),
	3883: uint16(3),
	3884: uint16(466),
	3885: uint16(1),
	3886: uint16(aux_sym_identifier_token1),
	3887: uint16(468),
	3888: uint16(1),
	3889: uint16(sym_function_link),
	3890: uint16(65),
	3891: uint16(1),
	3892: uint16(sym_identifier),
	3893: uint16(2),
	3894: uint16(470),
	3895: uint16(1),
	3896: uint16(anon_sym_in),
	3897: uint16(472),
	3898: uint16(2),
	3899: uint16(anon_sym_out),
	3900: uint16(anon_sym_inout),
	3901: uint16(2),
	3902: uint16(474),
	3903: uint16(1),
	3904: uint16(anon_sym_in),
	3905: uint16(476),
	3906: uint16(2),
	3907: uint16(anon_sym_out),
	3908: uint16(anon_sym_inout),
	3909: uint16(2),
	3910: uint16(478),
	3911: uint16(1),
	3912: uint16(anon_sym_in),
	3913: uint16(480),
	3914: uint16(2),
	3915: uint16(anon_sym_out),
	3916: uint16(anon_sym_inout),
	3917: uint16(3),
	3918: uint16(374),
	3919: uint16(1),
	3920: uint16(aux_sym_identifier_token1),
	3921: uint16(482),
	3922: uint16(1),
	3923: uint16(sym_function_link),
	3924: uint16(102),
	3925: uint16(1),
	3926: uint16(sym_identifier),
	3927: uint16(3),
	3928: uint16(466),
	3929: uint16(1),
	3930: uint16(aux_sym_identifier_token1),
	3931: uint16(484),
	3932: uint16(1),
	3933: uint16(sym_function_link),
	3934: uint16(41),
	3935: uint16(1),
	3936: uint16(sym_identifier),
	3937: uint16(3),
	3938: uint16(374),
	3939: uint16(1),
	3940: uint16(aux_sym_identifier_token1),
	3941: uint16(486),
	3942: uint16(1),
	3943: uint16(sym_function_link),
	3944: uint16(87),
	3945: uint16(1),
	3946: uint16(sym_identifier),
	3947: uint16(2),
	3948: uint16(488),
	3949: uint16(1),
	3950: uint16(aux_sym_identifier_token1),
	3951: uint16(200),
	3952: uint16(1),
	3953: uint16(sym_identifier),
	3954: uint16(2),
	3955: uint16(488),
	3956: uint16(1),
	3957: uint16(aux_sym_identifier_token1),
	3958: uint16(199),
	3959: uint16(1),
	3960: uint16(sym_identifier),
	3961: uint16(2),
	3962: uint16(488),
	3963: uint16(1),
	3964: uint16(aux_sym_identifier_token1),
	3965: uint16(196),
	3966: uint16(1),
	3967: uint16(sym_identifier),
	3968: uint16(2),
	3969: uint16(490),
	3970: uint16(1),
	3971: uint16(anon_sym_LBRACE),
	3972: uint16(492),
	3973: uint16(1),
	3974: uint16(sym_code_block_content),
	3975: uint16(2),
	3976: uint16(494),
	3977: uint16(1),
	3978: uint16(aux_sym_identifier_token1),
	3979: uint16(32),
	3980: uint16(1),
	3981: uint16(sym_identifier),
	3982: uint16(1),
	3983: uint16(278),
	3984: uint16(2),
	3985: uint16(aux_sym_identifier_token1),
	3986: uint16(anon_sym_TILDE),
	3987: uint16(2),
	3988: uint16(496),
	3989: uint16(1),
	3990: uint16(aux_sym_identifier_token1),
	3991: uint16(39),
	3992: uint16(1),
	3993: uint16(sym_identifier),
	3994: uint16(2),
	3995: uint16(488),
	3996: uint16(1),
	3997: uint16(aux_sym_identifier_token1),
	3998: uint16(169),
	3999: uint16(1),
	4000: uint16(sym_identifier),
	4001: uint16(2),
	4002: uint16(498),
	4003: uint16(1),
	4004: uint16(aux_sym_identifier_token1),
	4005: uint16(76),
	4006: uint16(1),
	4007: uint16(sym_identifier),
	4008: uint16(2),
	4009: uint16(500),
	4010: uint16(1),
	4011: uint16(aux_sym_identifier_token1),
	4012: uint16(84),
	4013: uint16(1),
	4014: uint16(sym_identifier),
	4015: uint16(1),
	4016: uint16(502),
	4017: uint16(1),
	4018: uint16(sym_code_block_content),
	4019: uint16(1),
	4020: uint16(504),
	4021: uint16(1),
	4022: uint16(anon_sym_RBRACK),
	4023: uint16(1),
	4024: uint16(506),
	4025: uint16(1),
	4026: uint16(aux_sym_identifier_token1),
	4027: uint16(1),
	4028: uint16(508),
	4029: uint16(1),
	4030: uint16(aux_sym_tag_token2),
	4031: uint16(1),
	4032: uint16(510),
	4033: uint16(1),
	4034: uint16(anon_sym_RBRACE),
	4035: uint16(1),
	4036: uint16(512),
	4037: uint16(1),
	4039: uint16(1),
	4040: uint16(514),
	4041: uint16(1),
	4042: uint16(anon_sym_LPAREN),
	4043: uint16(1),
	4044: uint16(425),
	4045: uint16(1),
	4047: uint16(1),
	4048: uint16(516),
	4049: uint16(1),
	4050: uint16(anon_sym_DOT),
	4051: uint16(1),
	4052: uint16(518),
	4053: uint16(1),
	4054: uint16(anon_sym_ATendcode),
	4055: uint16(1),
	4056: uint16(146),
	4057: uint16(1),
	4059: uint16(1),
	4060: uint16(520),
	4061: uint16(1),
	4062: uint16(anon_sym_LT_SLASHa_GT),
	4063: uint16(1),
	4064: uint16(522),
	4065: uint16(1),
	4066: uint16(anon_sym_ATendcode),
	4067: uint16(1),
	4068: uint16(524),
	4069: uint16(1),
	4070: uint16(anon_sym_GT),
	4071: uint16(1),
	4072: uint16(526),
	4073: uint16(1),
	4074: uint16(aux_sym_tag_token2),
	4075: uint16(1),
	4076: uint16(528),
	4077: uint16(1),
	4078: uint16(aux_sym_identifier_token1),
	4079: uint16(1),
	4080: uint16(530),
	4081: uint16(1),
	4082: uint16(aux_sym_identifier_token1),
	4083: uint16(1),
	4084: uint16(532),
	4085: uint16(1),
	4087: uint16(1),
	4088: uint16(534),
	4089: uint16(1),
	4090: uint16(anon_sym_RBRACK),
	4091: uint16(1),
	4092: uint16(536),
	4093: uint16(1),
	4094: uint16(aux_sym_link_token1),
	4095: uint16(1),
	4096: uint16(538),
	4097: uint16(1),
	4098: uint16(anon_sym_RPAREN),
	4099: uint16(1),
	4100: uint16(540),
	4101: uint16(1),
	4102: uint16(aux_sym_identifier_token1),
	4103: uint16(1),
	4104: uint16(542),
	4105: uint16(1),
	4106: uint16(sym_code_block_content),
	4107: uint16(1),
	4108: uint16(544),
	4109: uint16(1),
	4110: uint16(anon_sym_RPAREN),
	4111: uint16(1),
	4112: uint16(546),
	4113: uint16(1),
	4114: uint16(anon_sym_LT_SLASHa_GT),
	4115: uint16(1),
	4116: uint16(548),
	4117: uint16(1),
	4118: uint16(anon_sym_RBRACK),
	4119: uint16(1),
	4120: uint16(550),
	4121: uint16(1),
	4122: uint16(anon_sym_RPAREN),
	4123: uint16(1),
	4124: uint16(552),
	4125: uint16(1),
	4126: uint16(anon_sym_RPAREN),
	4127: uint16(1),
	4128: uint16(554),
	4129: uint16(1),
	4130: uint16(anon_sym_RPAREN),
	4131: uint16(1),
	4132: uint16(556),
	4133: uint16(1),
	4134: uint16(aux_sym_link_token2),
	4135: uint16(1),
	4136: uint16(558),
	4137: uint16(1),
	4138: uint16(anon_sym_RPAREN),
	4139: uint16(1),
	4140: uint16(560),
	4141: uint16(1),
	4142: uint16(anon_sym_RPAREN),
	4143: uint16(1),
	4144: uint16(562),
	4145: uint16(1),
	4146: uint16(sym_code_block_language),
	4147: uint16(1),
	4148: uint16(564),
	4149: uint16(1),
	4150: uint16(anon_sym_LPAREN),
	4151: uint16(1),
	4152: uint16(566),
	4153: uint16(1),
	4154: uint16(aux_sym_link_token2),
	4155: uint16(1),
	4156: uint16(568),
	4157: uint16(1),
	4158: uint16(sym_code_block_end),
	4159: uint16(1),
	4160: uint16(570),
	4161: uint16(1),
	4162: uint16(anon_sym_LPAREN),
	4163: uint16(1),
	4164: uint16(572),
	4165: uint16(1),
	4166: uint16(anon_sym_LPAREN),
	4167: uint16(1),
	4168: uint16(574),
	4169: uint16(1),
	4170: uint16(sym_code_block_language),
	4171: uint16(1),
	4172: uint16(576),
	4173: uint16(1),
	4174: uint16(anon_sym_GT),
	4175: uint16(1),
	4176: uint16(450),
	4177: uint16(1),
	4179: uint16(1),
	4180: uint16(578),
	4181: uint16(1),
	4182: uint16(anon_sym_RPAREN),
	4183: uint16(1),
	4184: uint16(580),
	4185: uint16(1),
	4186: uint16(aux_sym_link_token1),
}

var ts_small_parse_table_map = [204]uint32_t{
	1:   uint32(78),
	2:   uint32(133),
	3:   uint32(205),
	4:   uint32(269),
	5:   uint32(316),
	6:   uint32(363),
	7:   uint32(414),
	8:   uint32(461),
	9:   uint32(508),
	10:  uint32(555),
	11:  uint32(596),
	12:  uint32(637),
	13:  uint32(678),
	14:  uint32(716),
	15:  uint32(754),
	16:  uint32(786),
	17:  uint32(842),
	18:  uint32(885),
	19:  uint32(928),
	20:  uint32(971),
	21:  uint32(1002),
	22:  uint32(1045),
	23:  uint32(1074),
	24:  uint32(1103),
	25:  uint32(1146),
	26:  uint32(1171),
	27:  uint32(1199),
	28:  uint32(1227),
	29:  uint32(1251),
	30:  uint32(1281),
	31:  uint32(1305),
	32:  uint32(1329),
	33:  uint32(1357),
	34:  uint32(1387),
	35:  uint32(1424),
	36:  uint32(1461),
	37:  uint32(1498),
	38:  uint32(1521),
	39:  uint32(1544),
	40:  uint32(1579),
	41:  uint32(1602),
	42:  uint32(1625),
	43:  uint32(1648),
	44:  uint32(1696),
	45:  uint32(1744),
	46:  uint32(1776),
	47:  uint32(1808),
	48:  uint32(1840),
	49:  uint32(1874),
	50:  uint32(1922),
	51:  uint32(1970),
	52:  uint32(1992),
	53:  uint32(2040),
	54:  uint32(2062),
	55:  uint32(2084),
	56:  uint32(2106),
	57:  uint32(2128),
	58:  uint32(2162),
	59:  uint32(2184),
	60:  uint32(2206),
	61:  uint32(2228),
	62:  uint32(2250),
	63:  uint32(2278),
	64:  uint32(2301),
	65:  uint32(2326),
	66:  uint32(2351),
	67:  uint32(2372),
	68:  uint32(2393),
	69:  uint32(2420),
	70:  uint32(2465),
	71:  uint32(2490),
	72:  uint32(2514),
	73:  uint32(2538),
	74:  uint32(2558),
	75:  uint32(2578),
	76:  uint32(2598),
	77:  uint32(2618),
	78:  uint32(2642),
	79:  uint32(2668),
	80:  uint32(2694),
	81:  uint32(2713),
	82:  uint32(2732),
	83:  uint32(2751),
	84:  uint32(2770),
	85:  uint32(2789),
	86:  uint32(2820),
	87:  uint32(2848),
	88:  uint32(2866),
	89:  uint32(2884),
	90:  uint32(2902),
	91:  uint32(2930),
	92:  uint32(2948),
	93:  uint32(2966),
	94:  uint32(2994),
	95:  uint32(3012),
	96:  uint32(3030),
	97:  uint32(3048),
	98:  uint32(3066),
	99:  uint32(3083),
	100: uint32(3100),
	101: uint32(3119),
	102: uint32(3136),
	103: uint32(3153),
	104: uint32(3170),
	105: uint32(3187),
	106: uint32(3208),
	107: uint32(3225),
	108: uint32(3242),
	109: uint32(3259),
	110: uint32(3276),
	111: uint32(3293),
	112: uint32(3310),
	113: uint32(3327),
	114: uint32(3344),
	115: uint32(3361),
	116: uint32(3378),
	117: uint32(3395),
	118: uint32(3427),
	119: uint32(3459),
	120: uint32(3475),
	121: uint32(3507),
	122: uint32(3539),
	123: uint32(3571),
	124: uint32(3603),
	125: uint32(3616),
	126: uint32(3629),
	127: uint32(3642),
	128: uint32(3655),
	129: uint32(3676),
	130: uint32(3689),
	131: uint32(3702),
	132: uint32(3715),
	133: uint32(3728),
	134: uint32(3749),
	135: uint32(3762),
	136: uint32(3775),
	137: uint32(3788),
	138: uint32(3801),
	139: uint32(3816),
	140: uint32(3831),
	141: uint32(3846),
	142: uint32(3861),
	143: uint32(3872),
	144: uint32(3883),
	145: uint32(3893),
	146: uint32(3901),
	147: uint32(3909),
	148: uint32(3917),
	149: uint32(3927),
	150: uint32(3937),
	151: uint32(3947),
	152: uint32(3954),
	153: uint32(3961),
	154: uint32(3968),
	155: uint32(3975),
	156: uint32(3982),
	157: uint32(3987),
	158: uint32(3994),
	159: uint32(4001),
	160: uint32(4008),
	161: uint32(4015),
	162: uint32(4019),
	163: uint32(4023),
	164: uint32(4027),
	165: uint32(4031),
	166: uint32(4035),
	167: uint32(4039),
	168: uint32(4043),
	169: uint32(4047),
	170: uint32(4051),
	171: uint32(4055),
	172: uint32(4059),
	173: uint32(4063),
	174: uint32(4067),
	175: uint32(4071),
	176: uint32(4075),
	177: uint32(4079),
	178: uint32(4083),
	179: uint32(4087),
	180: uint32(4091),
	181: uint32(4095),
	182: uint32(4099),
	183: uint32(4103),
	184: uint32(4107),
	185: uint32(4111),
	186: uint32(4115),
	187: uint32(4119),
	188: uint32(4123),
	189: uint32(4127),
	190: uint32(4131),
	191: uint32(4135),
	192: uint32(4139),
	193: uint32(4143),
	194: uint32(4147),
	195: uint32(4151),
	196: uint32(4155),
	197: uint32(4159),
	198: uint32(4163),
	199: uint32(4167),
	200: uint32(4171),
	201: uint32(4175),
	202: uint32(4179),
	203: uint32(4183),
}

var ts_parse_actions = [582]TSParseActionEntry{
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(2)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(144)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(53)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(3)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(151)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(9)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(12)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(149)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(178)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(205)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(156)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(173)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(195)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(30)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(154)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_tag),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(170)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_document),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(135)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(166)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(37)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(147)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(165)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(184)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(182)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(141)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_tag),
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
		Fsymbol:      uint16(sym_tag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_tag),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(58)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_tag),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(2),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(2),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_description),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(178)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
	})))),
	128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(179)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(205)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
	})))),
	134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
	}})),
	136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	137: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__expression),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(193)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_document),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	149: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	157: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_qualified_identifier),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_qualified_identifier),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_qualified_identifier_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_qualified_identifier_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_qualified_identifier_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(157)),
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
		Fcount:    uint8(1),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_identifier),
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
		Fcount: uint8(1),
	}})),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_identifier),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	172: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tag_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(141)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tag_repeat1),
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
		Fsymbol:      uint16(aux_sym_tag_repeat1),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_brief_description_repeat1),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_brief_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(31)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_brief_description_repeat1),
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
		Fsymbol:      uint16(aux_sym_brief_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_qualified_identifier_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(159)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_brief_description),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(31)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_brief_description),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_function),
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
		Fcount: uint8(1),
	}})),
	202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_function),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(48)),
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_tag),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_function),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_function),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(71)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(1),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(8),
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
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(8),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(165)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	251: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(184)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(182)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_description_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(50)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_description_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(50)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_brief_header),
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
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_brief_header),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(46)),
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
		Fcount: uint8(1),
	}})),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_link),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	271: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_link),
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
		Fcount: uint8(1),
	}})),
	273: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_brief_header),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_brief_header),
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
		Fcount: uint8(1),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_storageclass),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_storageclass),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
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
		Fsymbol:        uint16(sym_code_word),
		Fproduction_id: uint16(4),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_code_word),
		Fproduction_id: uint16(4),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_emphasis),
		Fproduction_id: uint16(3),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:        uint16(sym_emphasis),
		Fproduction_id: uint16(3),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(183)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(146)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	306: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(10),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(10),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(189)),
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
		Fcount: uint8(2),
	}})),
	314: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(3)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	317: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(130)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(151)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	323: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(177)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	326: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(9)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	329: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	330: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(149)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	336: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(156)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(71)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(195)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_qualified_identifier_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(161)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	351: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_qualified_identifier_repeat1),
	})))),
	352: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(162)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(1),
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
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_tag_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(143)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	361: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_brief_description_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(80)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_brief_description_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(80)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(80)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(150)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(88)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(110)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(126)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(133)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	383: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(11),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	387: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(5),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(11),
	})))),
	388: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	389: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	390: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	392: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	393: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
	394: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	395: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_code_block),
	})))),
	396: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	397: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_code_block),
	})))),
	398: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(9),
	})))),
	400: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	401: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(9),
	})))),
	402: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	403: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(aux_sym_tag_repeat2),
		Fproduction_id: uint16(10),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(150)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
	}})),
	406: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(5),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	412: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(5),
	})))),
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
		Fcount: uint8(1),
	}})),
	414: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Fcount: uint8(1),
	}})),
	418: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(7),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	420: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(7),
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
		Fcount: uint8(1),
	}})),
	422: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(2),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	424: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_tag),
		Fproduction_id: uint16(2),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_document),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
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
		Fcount: uint8(2),
	}})),
	430: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
	431: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(8)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	433: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(135)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	436: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(152)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	439: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(166)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	442: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(20)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	445: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(37)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	448: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(147)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	451: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(27)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(153)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(148)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(155)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(35)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(188)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(188)),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(181)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(102)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(41)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(87)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	489: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(172)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(175)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(96)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	507: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(129)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	511: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	513: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_document),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(204)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(201)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
	}})))),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(99)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	523: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	524: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	525: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	526: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	529: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	530: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	531: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	532: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	533: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	535: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	536: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	537: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	538: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	539: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	540: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	541: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	542: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	543: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	544: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	545: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	546: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	547: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	548: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	549: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	550: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(98)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	553: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	557: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	563: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(191)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(194)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(167)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(197)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(202)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_brief_text = 0
const ts_external_token_code_block_start = 1
const ts_external_token_code_block_language = 2
const ts_external_token_code_block_content = 3
const ts_external_token_code_block_end = 4

var ts_external_scanner_symbol_map = [5]TSSymbol{
	0: uint16(sym_brief_text),
	1: uint16(sym_code_block_start),
	2: uint16(sym_code_block_language),
	3: uint16(sym_code_block_content),
	4: uint16(sym_code_block_end),
}

var ts_external_scanner_states = [8][5]uint8{
	1: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
		4: libc.BoolUint8(true1 != 0),
	},
	2: {
		1: libc.BoolUint8(true1 != 0),
	},
	3: {
		0: libc.BoolUint8(true1 != 0),
		1: libc.BoolUint8(true1 != 0),
	},
	4: {
		0: libc.BoolUint8(true1 != 0),
	},
	5: {
		3: libc.BoolUint8(true1 != 0),
	},
	6: {
		2: libc.BoolUint8(true1 != 0),
	},
	7: {
		4: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_doxygen(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fversion:                   uint32(LANGUAGE_VERSION),
	Fsymbol_count:              uint32(SYMBOL_COUNT),
	Falias_count:               uint32(ALIAS_COUNT),
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_doxygen_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_doxygen_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_doxygen_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_doxygen_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_doxygen_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "tree-sitter-doxygen: Invalid buffer length %d! This should never happen\n\x00endcode\x00end\x00tag_name\x00brief_description\x00,\x00tag_token1\x00type\x00identifier_token1\x00::\x00~\x00(\x00)\x00[\x00in\x00out\x00inout\x00]\x00\\a\x00\\c\x00<a\x00link_token1\x00>\x00text\x00</a>\x00function_link\x00@code\x00{\x00.\x00}\x00@endcode\x00_text\x00_singleline_begin\x00_multiline_begin\x00/\x00*/\x00_text_line\x00brief_text\x00code_block_start\x00code_block_language\x00code_block_content\x00code_block_end\x00document\x00brief_header\x00description\x00tag\x00_expression\x00identifier\x00qualified_identifier\x00function\x00storageclass\x00emphasis\x00code_word\x00link\x00code_block\x00_multiline_end\x00document_repeat1\x00document_repeat2\x00brief_description_repeat1\x00description_repeat1\x00tag_repeat1\x00tag_repeat2\x00qualified_identifier_repeat1\x00code\x00"
