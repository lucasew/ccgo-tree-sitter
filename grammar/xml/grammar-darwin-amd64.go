// Code generated for darwin/amd64 by 'ccgo -ignore-unsupported-alignment -ignore-unsupported-atomic-sizes -ignore-vector-functions -std=gnu11 -O0 -D__attribute__(...)= -D__extension__= -D_Nonnull= -D_Nullable= -D_Null_unspecified= -DAPI_AVAILABLE(...)= -DAPI_UNAVAILABLE(...)= -DAPI_DEPRECATED(...)= -DAPI_DEPRECATED_WITH_REPLACEMENT(...)= -D__API_AVAILABLE(...)= -D__API_UNAVAILABLE(...)= -D__API_DEPRECATED(...)= -D__API_DEPRECATED_WITH_REPLACEMENT(...)= -Dfunc=func_token -Dinterface=interface_token -Dselect=select_token -Dchan=chan_token -Dgo=go_token -Dmap=map_token -Dpackage=package_token -Dtype=type_token -Dvar=var_token -Dimport=import_token -Ddefer=defer_token -Dfallthrough=fallthrough_token -Drange=range_token -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-xml/xml/src -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-xml/xml -I /Users/runner/work/ccgo-tree-sitter/ccgo-tree-sitter/third-party/tree-sitter-xml -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/include -I /Users/runner/.cache/workspaced/sources/github/7349d53bdd4ab31f79a3eec0385f924117fa6d06f018e00aadd823c0dcd51ac2/lib/src combined.c -o grammar.go', DO NOT EDIT.

//go:build darwin && amd64

package grammar_xml

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
const EXTERNAL_TOKEN_COUNT = 11
const FIELD_COUNT = 2
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
const MAX_ALIAS_SEQUENCE_LENGTH = 12
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
const STATE_COUNT = 495
const SUPERTYPE_COUNT = 0
const SV_INTERRUPT = "SA_RESTART"
const SV_NOCLDSTOP = "SA_NOCLDSTOP"
const SV_NODEFER = "SA_NODEFER"
const SV_ONSTACK = "SA_ONSTACK"
const SV_RESETHAND = "SA_RESETHAND"
const SV_SIGINFO = "SA_SIGINFO"
const SYMBOL_COUNT = 143
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
const TOKEN_COUNT = 74
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
const __ASSERT_FILE_NAME = "__FILE__"
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
const static_assert = "_Static_assert"
const stderr1 = "__stderrp"
const stdin1 = "__stdinp"
const stdout1 = "__stdoutp"
const sv_onstack = "sv_flags"
const true1 = 1
const ts_builtin_sym_end = 0
const ts_calloc = "calloc"
const ts_free = "free"
const ts_malloc = "malloc"
const ts_realloc = "realloc"
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

const PI_TARGET = 0
const PI_CONTENT = 1
const COMMENT = 2
const CHAR_DATA = 3
const CDATA = 4
const XML_MODEL = 5
const XML_STYLESHEET = 6
const START_TAG_NAME = 7
const END_TAG_NAME = 8
const ERRONEOUS_END_NAME = 9
const SELF_CLOSING_TAG_DELIMITER = 10

func scan_pi_target(tls *libc.TLS, lexer2 uintptr, valid_symbols uintptr) (r uint8) {
	var advanced_once, found_x_first, last_char_hyphen, v2 uint8
	var j uint32
	var v1 wchar_t
	var v10, v11, v4, v8 int32
	var v13, v26 uintptr
	var v6 __darwin_ct_rune_t
	var v7 uint64
	var v46 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = advanced_once, found_x_first, j, last_char_hyphen, v1, v10, v11, v13, v2, v26, v4, v46, v6, v7, v8
	advanced_once = libc.BoolUint8(false1 != 0)
	found_x_first = libc.BoolUint8(false1 != 0)
	v1 = (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead
	v6 = v1
	v7 = uint64(0x00000100)
	v11 = libc.BoolInt32(v6 & ^libc.Int32FromInt32(0x7F) == 0)
	goto _12
_12:
	if v11 != 0 {
		v10 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v6)*4)))&v7 != 0))
	} else {
		v10 = libc.BoolInt32(!!(libc.X__maskrune(tls, v6, v7) != 0))
	}
	v8 = v10
	goto _9
_9:
	v4 = v8
	goto _5
_5:
	;
	v2 = libc.BoolUint8(v4 != 0 || v1 == int32('_') || v1 == int32(':'))
	goto _3
_3:
	if v2 != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead == int32('x') || (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead == int32('X') {
			found_x_first = libc.BoolUint8(true1 != 0)
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer2)).Fmark_end})))(tls, lexer2)
		}
		advanced_once = libc.BoolUint8(true1 != 0)
		v13 = lexer2
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Fadvance})))(tls, v13, libc.BoolUint8(false1 != 0))
	}
	if advanced_once != 0 {
		for {
			v1 = (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead
			v6 = v1
			v7 = libc.Uint64FromInt64(libc.Int64FromInt64(0x00000100) | libc.Int64FromInt64(0x00000400))
			v11 = libc.BoolInt32(v6 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _25
		_25:
			if v11 != 0 {
				v10 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v6)*4)))&v7 != 0))
			} else {
				v10 = libc.BoolInt32(!!(libc.X__maskrune(tls, v6, v7) != 0))
			}
			v8 = v10
			goto _22
		_22:
			v4 = v8
			goto _18
		_18:
			;
			v2 = libc.BoolUint8(v4 != 0 || v1 == int32('_') || v1 == int32(':') || v1 == int32('.') || v1 == int32('-') || v1 == int32(0xB7))
			goto _16
		_16:
			if !(v2 != 0) {
				break
			}
			if found_x_first != 0 && ((*TSLexer)(unsafe.Pointer(lexer2)).Flookahead == int32('m') || (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead == int32('M')) {
				v13 = lexer2
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Fadvance})))(tls, v13, libc.BoolUint8(false1 != 0))
				if (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead == int32('l') || (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead == int32('L') {
					v13 = lexer2
					(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Fadvance})))(tls, v13, libc.BoolUint8(false1 != 0))
					v1 = (*TSLexer)(unsafe.Pointer(lexer2)).Flookahead
					v6 = v1
					v7 = libc.Uint64FromInt64(libc.Int64FromInt64(0x00000100) | libc.Int64FromInt64(0x00000400))
					v11 = libc.BoolInt32(v6 & ^libc.Int32FromInt32(0x7F) == 0)
					goto _39
				_39:
					if v11 != 0 {
						v10 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v6)*4)))&v7 != 0))
					} else {
						v10 = libc.BoolInt32(!!(libc.X__maskrune(tls, v6, v7) != 0))
					}
					v8 = v10
					goto _36
				_36:
					v4 = v8
					goto _32
				_32:
					;
					v2 = libc.BoolUint8(v4 != 0 || v1 == int32('_') || v1 == int32(':') || v1 == int32('.') || v1 == int32('-') || v1 == int32(0xB7))
					goto _30
				_30:
					if v2 != 0 {
						found_x_first = libc.BoolUint8(false1 != 0)
						last_char_hyphen = libc.BoolUint8((*TSLexer)(unsafe.Pointer(lexer2)).Flookahead == int32('-'))
						v13 = lexer2
						(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Fadvance})))(tls, v13, libc.BoolUint8(false1 != 0))
						if last_char_hyphen != 0 {
							if v46 = *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(XML_MODEL))) != 0; v46 {
								v13 = lexer2
								j = uint32(0)
								for {
									if !(j < uint32(5)) {
										break
									}
									if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Feof})))(tls, v13) != 0) && (*TSLexer)(unsafe.Pointer(v13)).Flookahead == int32(*(*int8)(unsafe.Pointer(__ccgo_ts + uintptr(j)))) {
										v26 = v13
										(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v26)).Fadvance})))(tls, v26, libc.BoolUint8(false1 != 0))
									} else {
										v2 = libc.BoolUint8(false1 != 0)
										goto _45
									}
									goto _42
								_42:
									;
									j = j + 1
								}
								v2 = libc.BoolUint8(true1 != 0)
								goto _45
							_45:
							}
							if v46 && v2 != 0 {
								return libc.BoolUint8(false1 != 0)
							}
							if v46 = *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(XML_STYLESHEET))) != 0; v46 {
								v13 = lexer2
								j = uint32(0)
								for {
									if !(j < uint32(10)) {
										break
									}
									if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Feof})))(tls, v13) != 0) && (*TSLexer)(unsafe.Pointer(v13)).Flookahead == int32(*(*int8)(unsafe.Pointer(__ccgo_ts + 6 + uintptr(j)))) {
										v26 = v13
										(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v26)).Fadvance})))(tls, v26, libc.BoolUint8(false1 != 0))
									} else {
										v2 = libc.BoolUint8(false1 != 0)
										goto _51
									}
									goto _48
								_48:
									;
									j = j + 1
								}
								v2 = libc.BoolUint8(true1 != 0)
								goto _51
							_51:
							}
							if v46 && v2 != 0 {
								return libc.BoolUint8(false1 != 0)
							}
						}
					} else {
						return libc.BoolUint8(false1 != 0)
					}
				}
			}
			found_x_first = libc.BoolUint8(false1 != 0)
			v13 = lexer2
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Fadvance})))(tls, v13, libc.BoolUint8(false1 != 0))
		}
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer2)).Fmark_end})))(tls, lexer2)
		(*TSLexer)(unsafe.Pointer(lexer2)).Fresult_symbol = uint16(PI_TARGET)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan_pi_content(tls *libc.TLS, lexer1 uintptr) (r uint8) {
	var v1 uintptr
	_ = v1
	for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead != int32('\n') && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead != int32('?') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	}
	if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead != int32('?') {
		return libc.BoolUint8(false1 != 0)
	}
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
	v1 = lexer1
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('>') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		for (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32(' ') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		}
		if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('\n') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		} else {
			return libc.BoolUint8(false1 != 0)
		}
		(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(PI_CONTENT)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

func scan_comment(tls *libc.TLS, lexer1 uintptr) (r uint8) {
	var v1 uintptr
	_ = v1
	if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	} else {
		return libc.BoolUint8(false1 != 0)
	}
	if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	} else {
		return libc.BoolUint8(false1 != 0)
	}
	for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) {
		if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
			if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('-') {
				v1 = lexer1
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
				break
			}
		} else {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		}
	}
	if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('>') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
		(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(COMMENT)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

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

type rsize_t = uint64

type errno_t = int32

type Array = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

/// A binary search routine, based on Rust's `std::slice::binary_search_by`.
/// This is not what you're looking for, see `array_search_sorted_with` or `array_search_sorted_by`.

/// Helper macro for the `_sorted_by` routines below. This takes the left (existing)
/// parameter by reference in order to work with the generic sorting function above.

/*
 * Copyright (c) 2000, 2007, 2010, 2023 Apple Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*-
 * Copyright (c) 1990, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)string.h	8.1 (Berkeley) 6/2/93
 */

/*
 * Copyright (c) 2023 Apple Inc. All rights reserved.
 *
 * @APPLE_LICENSE_HEADER_START@
 *
 * This file contains Original Code and/or Modifications of Original Code
 * as defined in and that are subject to the Apple Public Source License
 * Version 2.0 (the 'License'). You may not use this file except in
 * compliance with the License. Please obtain a copy of the License at
 * http://www.opensource.apple.com/apsl/ and read it before using this
 * file.
 *
 * The Original Code and all software distributed under the License are
 * distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
 * EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
 * INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
 * Please see the License for the specific language governing rights and
 * limitations under the License.
 *
 * @APPLE_LICENSE_HEADER_END@
 */
/*-
 * Copyright (c) 1990, 1993
 *	The Regents of the University of California.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 * 3. All advertising materials mentioning features or use of this software
 *    must display the following acknowledgement:
 *	This product includes software developed by the University of
 *	California, Berkeley and its contributors.
 * 4. Neither the name of the University nor the names of its contributors
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 *	@(#)string.h	8.1 (Berkeley) 6/2/93
 */

type String = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

type Vector = struct {
	Fcontents uintptr
	Fsize     uint32_t
	Fcapacity uint32_t
}

func string_eq(tls *libc.TLS, a uintptr, b uintptr) (r uint8) {
	if (*String)(unsafe.Pointer(a)).Fsize != (*String)(unsafe.Pointer(b)).Fsize {
		return libc.BoolUint8(false1 != 0)
	}
	return libc.BoolUint8(libc.Xmemcmp(tls, (*String)(unsafe.Pointer(a)).Fcontents, (*String)(unsafe.Pointer(b)).Fcontents, uint64((*String)(unsafe.Pointer(a)).Fsize)) == 0)
}

func scan_tag_name(tls *libc.TLS, lexer1 uintptr) (r String) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var new_capacity1, new_size, v16, v17 uint32_t
	var v1 wchar_t
	var v10, v11, v4, v8 int32
	var v13, v14, v18 uintptr
	var v15 size_t
	var v2 uint8
	var v6 __darwin_ct_rune_t
	var v7 uint64
	var _ /* tag_name at bp+0 */ String
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = new_capacity1, new_size, v1, v10, v11, v13, v14, v15, v16, v17, v18, v2, v4, v6, v7, v8
	*(*String)(unsafe.Pointer(bp)) = String{}
	v1 = (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead
	v6 = v1
	v7 = uint64(0x00000100)
	v11 = libc.BoolInt32(v6 & ^libc.Int32FromInt32(0x7F) == 0)
	goto _12
_12:
	if v11 != 0 {
		v10 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v6)*4)))&v7 != 0))
	} else {
		v10 = libc.BoolInt32(!!(libc.X__maskrune(tls, v6, v7) != 0))
	}
	v8 = v10
	goto _9
_9:
	v4 = v8
	goto _5
_5:
	;
	v2 = libc.BoolUint8(v4 != 0 || v1 == int32('_') || v1 == int32(':'))
	goto _3
_3:
	if v2 != 0 {
		v13 = bp
		new_size = (*Array)(unsafe.Pointer(v13)).Fsize + uint32(1)
		if new_size > (*Array)(unsafe.Pointer(v13)).Fcapacity {
			new_capacity1 = (*Array)(unsafe.Pointer(v13)).Fcapacity * uint32(2)
			if new_capacity1 < libc.Uint32FromInt32(8) {
				new_capacity1 = uint32(8)
			}
			if new_capacity1 < new_size {
				new_capacity1 = new_size
			}
			v14 = v13
			v15 = libc.Uint64FromInt64(1)
			v16 = new_capacity1
			if v16 > (*Array)(unsafe.Pointer(v14)).Fcapacity {
				if (*Array)(unsafe.Pointer(v14)).Fcontents != 0 {
					(*Array)(unsafe.Pointer(v14)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v14)).Fcontents, uint64(v16)*v15)
				} else {
					(*Array)(unsafe.Pointer(v14)).Fcontents = libc.Xmalloc(tls, uint64(v16)*v15)
				}
				(*Array)(unsafe.Pointer(v14)).Fcapacity = v16
			}
		}
		v18 = bp + 8
		v17 = *(*uint32_t)(unsafe.Pointer(v18))
		*(*uint32_t)(unsafe.Pointer(v18)) = *(*uint32_t)(unsafe.Pointer(v18)) + 1
		*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(bp)).Fcontents + uintptr(v17))) = int8((*TSLexer)(unsafe.Pointer(lexer1)).Flookahead)
		v13 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Fadvance})))(tls, v13, libc.BoolUint8(false1 != 0))
	}
	for {
		v1 = (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead
		v6 = v1
		v7 = libc.Uint64FromInt64(libc.Int64FromInt64(0x00000100) | libc.Int64FromInt64(0x00000400))
		v11 = libc.BoolInt32(v6 & ^libc.Int32FromInt32(0x7F) == 0)
		goto _31
	_31:
		if v11 != 0 {
			v10 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v6)*4)))&v7 != 0))
		} else {
			v10 = libc.BoolInt32(!!(libc.X__maskrune(tls, v6, v7) != 0))
		}
		v8 = v10
		goto _28
	_28:
		v4 = v8
		goto _24
	_24:
		;
		v2 = libc.BoolUint8(v4 != 0 || v1 == int32('_') || v1 == int32(':') || v1 == int32('.') || v1 == int32('-') || v1 == int32(0xB7))
		goto _22
	_22:
		if !(v2 != 0) {
			break
		}
		v13 = bp
		new_size = (*Array)(unsafe.Pointer(v13)).Fsize + uint32(1)
		if new_size > (*Array)(unsafe.Pointer(v13)).Fcapacity {
			new_capacity1 = (*Array)(unsafe.Pointer(v13)).Fcapacity * uint32(2)
			if new_capacity1 < libc.Uint32FromInt32(8) {
				new_capacity1 = uint32(8)
			}
			if new_capacity1 < new_size {
				new_capacity1 = new_size
			}
			v14 = v13
			v15 = libc.Uint64FromInt64(1)
			v16 = new_capacity1
			if v16 > (*Array)(unsafe.Pointer(v14)).Fcapacity {
				if (*Array)(unsafe.Pointer(v14)).Fcontents != 0 {
					(*Array)(unsafe.Pointer(v14)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v14)).Fcontents, uint64(v16)*v15)
				} else {
					(*Array)(unsafe.Pointer(v14)).Fcontents = libc.Xmalloc(tls, uint64(v16)*v15)
				}
				(*Array)(unsafe.Pointer(v14)).Fcapacity = v16
			}
		}
		v18 = bp + 8
		v17 = *(*uint32_t)(unsafe.Pointer(v18))
		*(*uint32_t)(unsafe.Pointer(v18)) = *(*uint32_t)(unsafe.Pointer(v18)) + 1
		*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(bp)).Fcontents + uintptr(v17))) = int8((*TSLexer)(unsafe.Pointer(lexer1)).Flookahead)
		v13 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v13)).Fadvance})))(tls, v13, libc.BoolUint8(false1 != 0))
	}
	return *(*String)(unsafe.Pointer(bp))
}

func scan_start_tag_name(tls *libc.TLS, tags uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var new_capacity1, new_size, v5, v6 uint32_t
	var v1, v2, v3 uintptr
	var v4 size_t
	var _ /* tag_name at bp+0 */ String
	_, _, _, _, _, _, _, _ = new_capacity1, new_size, v1, v2, v3, v4, v5, v6
	*(*String)(unsafe.Pointer(bp)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp))).Fsize == uint32(0) {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(START_TAG_NAME)
	v1 = tags
	new_size = (*Array)(unsafe.Pointer(v1)).Fsize + uint32(1)
	if new_size > (*Array)(unsafe.Pointer(v1)).Fcapacity {
		new_capacity1 = (*Array)(unsafe.Pointer(v1)).Fcapacity * uint32(2)
		if new_capacity1 < libc.Uint32FromInt32(8) {
			new_capacity1 = uint32(8)
		}
		if new_capacity1 < new_size {
			new_capacity1 = new_size
		}
		v2 = v1
		v4 = libc.Uint64FromInt64(16)
		v5 = new_capacity1
		if v5 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
			if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
				(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v5)*v4)
			} else {
				(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v5)*v4)
			}
			(*Array)(unsafe.Pointer(v2)).Fcapacity = v5
		}
	}
	v3 = tags + 8
	v6 = *(*uint32_t)(unsafe.Pointer(v3))
	*(*uint32_t)(unsafe.Pointer(v3)) = *(*uint32_t)(unsafe.Pointer(v3)) + 1
	*(*String)(unsafe.Pointer((*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(v6)*16)) = *(*String)(unsafe.Pointer(bp))
	return libc.BoolUint8(true1 != 0)
}

func scan_end_tag_name(tls *libc.TLS, tags uintptr, lexer uintptr) (r uint8) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1, v4 uintptr
	var v2 bool
	var v3 uint32_t
	var _ /* last_tag at bp+16 */ String
	var _ /* tag_name at bp+0 */ String
	_, _, _, _ = v1, v2, v3, v4
	*(*String)(unsafe.Pointer(bp)) = scan_tag_name(tls, lexer)
	if (*(*String)(unsafe.Pointer(bp))).Fsize == uint32(0) {
		v1 = bp
		if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
			(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
		}
		return libc.BoolUint8(false1 != 0)
	}
	if v2 = (*Vector)(unsafe.Pointer(tags)).Fsize > uint32(0); v2 {
		if libc.BoolInt64(!((*Vector)(unsafe.Pointer(tags)).Fsize-libc.Uint32FromInt32(1) < (*Vector)(unsafe.Pointer(tags)).Fsize)) != 0 {
			libc.X__assert_rtn(tls, uintptr(unsafe.Pointer(&__func__1)), __ccgo_ts+17, int32(50), __ccgo_ts+28)
		}
	}
	if v2 && string_eq(tls, (*Vector)(unsafe.Pointer(tags)).Fcontents+uintptr((*Vector)(unsafe.Pointer(tags)).Fsize-uint32(1))*16, bp) != 0 {
		v1 = tags + 8
		*(*uint32_t)(unsafe.Pointer(v1)) = *(*uint32_t)(unsafe.Pointer(v1)) - 1
		v3 = *(*uint32_t)(unsafe.Pointer(v1))
		*(*String)(unsafe.Pointer(bp + 16)) = *(*String)(unsafe.Pointer((*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(v3)*16))
		v4 = bp + 16
		if (*Array)(unsafe.Pointer(v4)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v4)).Fcontents)
			(*Array)(unsafe.Pointer(v4)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v4)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v4)).Fcapacity = uint32(0)
		}
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(END_TAG_NAME)
	} else {
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ERRONEOUS_END_NAME)
	}
	v1 = bp
	if (*Array)(unsafe.Pointer(v1)).Fcontents != 0 {
		libc.Xfree(tls, (*Array)(unsafe.Pointer(v1)).Fcontents)
		(*Array)(unsafe.Pointer(v1)).Fcontents = libc.UintptrFromInt32(0)
		(*Array)(unsafe.Pointer(v1)).Fsize = uint32(0)
		(*Array)(unsafe.Pointer(v1)).Fcapacity = uint32(0)
	}
	return libc.BoolUint8(libc.Int32FromUint16((*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol) == int32(END_TAG_NAME))
}

var __func__1 = [18]int8{'s', 'c', 'a', 'n', '_', 'e', 'n', 'd', '_', 't', 'a', 'g', '_', 'n', 'a', 'm', 'e'}

func scan_self_closing_tag_delimiter(tls *libc.TLS, tags uintptr, lexer1 uintptr) (r uint8) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1, v2 uintptr
	var v3 uint32_t
	var _ /* last_tag at bp+0 */ String
	_, _, _ = v1, v2, v3
	v1 = lexer1
	(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	if !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) && (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('>') {
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	} else {
		return libc.BoolUint8(false1 != 0)
	}
	if (*Vector)(unsafe.Pointer(tags)).Fsize > uint32(0) {
		v1 = tags + 8
		*(*uint32_t)(unsafe.Pointer(v1)) = *(*uint32_t)(unsafe.Pointer(v1)) - 1
		v3 = *(*uint32_t)(unsafe.Pointer(v1))
		*(*String)(unsafe.Pointer(bp)) = *(*String)(unsafe.Pointer((*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(v3)*16))
		v2 = bp
		if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v2)).Fcontents)
			(*Array)(unsafe.Pointer(v2)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v2)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v2)).Fcapacity = uint32(0)
		}
		(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(SELF_CLOSING_TAG_DELIMITER)
	}
	return libc.BoolUint8(true1 != 0)
}

// C documentation
//
//	/// Check if the lexer is in error recovery mode
func in_error_recovery(tls *libc.TLS, valid_symbols uintptr) (r uint8) {
	return libc.BoolUint8(*(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_TARGET))) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_CONTENT))) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(COMMENT))) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CHAR_DATA))) != 0 && *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CDATA))) != 0)
}

// C documentation
//
//	/// Check if the lexer is in a char data node
func in_char_data(tls *libc.TLS, lexer uintptr) (r uint8) {
	return libc.BoolUint8(!((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Feof})))(tls, lexer) != 0) && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('<') && (*TSLexer)(unsafe.Pointer(lexer)).Flookahead != int32('&'))
}

// C documentation
//
//	/// Scan for a CharData node
func scan_char_data(tls *libc.TLS, lexer1 uintptr) (r uint8) {
	var advanced_once uint8
	var v1 uintptr
	_, _ = advanced_once, v1
	advanced_once = libc.BoolUint8(false1 != 0)
	for in_char_data(tls, lexer1) != 0 {
		if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32(']') {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
			if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32(']') {
				v1 = lexer1
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
				if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('>') {
					v1 = lexer1
					(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
					if advanced_once != 0 {
						(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(CHAR_DATA)
						return libc.BoolUint8(false1 != 0)
					}
				}
			}
		}
		advanced_once = libc.BoolUint8(true1 != 0)
		if in_char_data(tls, lexer1) != 0 {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		}
	}
	if advanced_once != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
		(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(CHAR_DATA)
		return libc.BoolUint8(true1 != 0)
	}
	return libc.BoolUint8(false1 != 0)
}

// C documentation
//
//	/// Scan for a CData node
func scan_cdata(tls *libc.TLS, lexer1 uintptr) (r uint8) {
	var advanced_once uint8
	var v1 uintptr
	_, _ = advanced_once, v1
	advanced_once = libc.BoolUint8(false1 != 0)
	for !((*(*func(*libc.TLS, uintptr) uint8)(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Feof})))(tls, lexer1) != 0) {
		if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32(']') {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
			if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32(']') {
				v1 = lexer1
				(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
				if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('>') && advanced_once != 0 {
					(*TSLexer)(unsafe.Pointer(lexer1)).Fresult_symbol = uint16(CDATA)
					return libc.BoolUint8(true1 != 0)
				}
			}
		}
		advanced_once = libc.BoolUint8(true1 != 0)
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_xml_external_scanner_scan(tls *libc.TLS, payload uintptr, lexer1 uintptr, valid_symbols uintptr) (r uint8) {
	var tags, v1 uintptr
	_, _ = tags, v1
	tags = payload
	if in_error_recovery(tls, valid_symbols) != 0 {
		return libc.BoolUint8(false1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_TARGET))) != 0 {
		return scan_pi_target(tls, lexer1, valid_symbols)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(PI_CONTENT))) != 0 {
		return scan_pi_content(tls, lexer1)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CHAR_DATA))) != 0 && scan_char_data(tls, lexer1) != 0 {
		return libc.BoolUint8(true1 != 0)
	}
	if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(CDATA))) != 0 && scan_cdata(tls, lexer1) != 0 {
		return libc.BoolUint8(true1 != 0)
	}
	switch (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead {
	case int32('<'):
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer1)).Fmark_end})))(tls, lexer1)
		v1 = lexer1
		(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
		if (*TSLexer)(unsafe.Pointer(lexer1)).Flookahead == int32('!') {
			v1 = lexer1
			(*(*func(*libc.TLS, uintptr, uint8))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(v1)).Fadvance})))(tls, v1, libc.BoolUint8(false1 != 0))
			return scan_comment(tls, lexer1)
		}
	case int32('/'):
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(SELF_CLOSING_TAG_DELIMITER))) != 0 {
			return scan_self_closing_tag_delimiter(tls, tags, lexer1)
		}
	case int32('\000'):
	default:
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(START_TAG_NAME))) != 0 {
			return scan_start_tag_name(tls, tags, lexer1)
		}
		if *(*uint8)(unsafe.Pointer(valid_symbols + uintptr(END_TAG_NAME))) != 0 {
			return scan_end_tag_name(tls, tags, lexer1)
		}
	}
	return libc.BoolUint8(false1 != 0)
}

func tree_sitter_xml_external_scanner_create(tls *libc.TLS) (r uintptr) {
	var tags uintptr
	_ = tags
	tags = libc.Xcalloc(tls, uint64(1), uint64(16))
	if tags == libc.UintptrFromInt32(0) {
		libc.Xabort(tls)
	}
	(*Vector)(unsafe.Pointer(tags)).Fsize = uint32(0)
	(*Vector)(unsafe.Pointer(tags)).Fcapacity = uint32(0)
	(*Vector)(unsafe.Pointer(tags)).Fcontents = libc.UintptrFromInt32(0)
	return tags
}

func tree_sitter_xml_external_scanner_destroy(tls *libc.TLS, payload uintptr) {
	var i uint32_t
	var tags, v2 uintptr
	_, _, _ = i, tags, v2
	tags = payload
	i = uint32(0)
	for {
		if !(i < (*Vector)(unsafe.Pointer(tags)).Fsize) {
			break
		}
		if libc.BoolInt64(!(i < (*Vector)(unsafe.Pointer(tags)).Fsize)) != 0 {
			libc.X__assert_rtn(tls, uintptr(unsafe.Pointer(&__func__11)), __ccgo_ts+17, int32(200), __ccgo_ts+72)
		}
		v2 = (*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(i)*16
		if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v2)).Fcontents)
			(*Array)(unsafe.Pointer(v2)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v2)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v2)).Fcapacity = uint32(0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	v2 = tags
	if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
		libc.Xfree(tls, (*Array)(unsafe.Pointer(v2)).Fcontents)
		(*Array)(unsafe.Pointer(v2)).Fcontents = libc.UintptrFromInt32(0)
		(*Array)(unsafe.Pointer(v2)).Fsize = uint32(0)
		(*Array)(unsafe.Pointer(v2)).Fcapacity = uint32(0)
	}
	libc.Xfree(tls, tags)
}

var __func__11 = [41]int8{'t', 'r', 'e', 'e', '_', 's', 'i', 't', 't', 'e', 'r', '_', 'x', 'm', 'l', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 's', 'c', 'a', 'n', 'n', 'e', 'r', '_', 'd', 'e', 's', 't', 'r', 'o', 'y'}

func tree_sitter_xml_external_scanner_serialize(tls *libc.TLS, payload uintptr, buffer uintptr) (r uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var name_length, size, v3 uint32_t
	var tag, tags, v4 uintptr
	var v1 uint32
	var _ /* serialized_tag_count at bp+4 */ uint32_t
	var _ /* tag_count at bp+0 */ uint32_t
	_, _, _, _, _, _, _ = name_length, size, tag, tags, v1, v3, v4
	tags = payload
	if (*Vector)(unsafe.Pointer(tags)).Fsize > uint32(65535) {
		v1 = uint32(65535)
	} else {
		v1 = (*Vector)(unsafe.Pointer(tags)).Fsize
	}
	*(*uint32_t)(unsafe.Pointer(bp)) = v1
	*(*uint32_t)(unsafe.Pointer(bp + 4)) = uint32(0)
	size = uint32(4)
	libc.X__builtin___memcpy_chk(tls, buffer+uintptr(size), bp, uint64(size), ^__predefined_size_t(0))
	size = uint32(uint64(size) + libc.Uint64FromInt64(4))
	for {
		if !(*(*uint32_t)(unsafe.Pointer(bp + 4)) < *(*uint32_t)(unsafe.Pointer(bp))) {
			break
		}
		if libc.BoolInt64(!(*(*uint32_t)(unsafe.Pointer(bp + 4)) < (*Vector)(unsafe.Pointer(tags)).Fsize)) != 0 {
			libc.X__assert_rtn(tls, uintptr(unsafe.Pointer(&__func__12)), __ccgo_ts+17, int32(215), __ccgo_ts+101)
		}
		tag = (*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(*(*uint32_t)(unsafe.Pointer(bp + 4)))*16
		name_length = (*String)(unsafe.Pointer(tag)).Fsize
		if name_length > uint32(255) {
			name_length = uint32(255)
		}
		if size+uint32(2)+name_length >= uint32(TREE_SITTER_SERIALIZATION_BUFFER_SIZE) {
			break
		}
		v3 = size
		size = size + 1
		*(*int8)(unsafe.Pointer(buffer + uintptr(v3))) = libc.Int8FromUint32(name_length)
		if name_length > uint32(0) {
			libc.X__builtin___memcpy_chk(tls, buffer+uintptr(size), (*String)(unsafe.Pointer(tag)).Fcontents, uint64(name_length), ^__predefined_size_t(0))
		}
		v4 = tag
		if (*Array)(unsafe.Pointer(v4)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v4)).Fcontents)
			(*Array)(unsafe.Pointer(v4)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v4)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v4)).Fcapacity = uint32(0)
		}
		size = size + name_length
		goto _2
	_2:
		;
		*(*uint32_t)(unsafe.Pointer(bp + 4)) = *(*uint32_t)(unsafe.Pointer(bp + 4)) + 1
	}
	libc.X__builtin___memcpy_chk(tls, buffer, bp+4, uint64(4), ^__predefined_size_t(0))
	return size
}

var __func__12 = [43]int8{'t', 'r', 'e', 'e', '_', 's', 'i', 't', 't', 'e', 'r', '_', 'x', 'm', 'l', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 's', 'c', 'a', 'n', 'n', 'e', 'r', '_', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e'}

func tree_sitter_xml_external_scanner_deserialize(tls *libc.TLS, payload uintptr, buffer uintptr, length uint32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i uint32
	var iter, new_capacity1, new_size, size, v5, v7 uint32_t
	var tag1 String
	var tags, v2, v3, v8 uintptr
	var v4 size_t
	var _ /* serialized_tag_count at bp+4 */ uint32_t
	var _ /* tag at bp+8 */ String
	var _ /* tag_count at bp+0 */ uint32_t
	_, _, _, _, _, _, _, _, _, _, _, _, _ = i, iter, new_capacity1, new_size, size, tag1, tags, v2, v3, v4, v5, v7, v8
	tags = payload
	i = uint32(0)
	for {
		if !(i < (*Vector)(unsafe.Pointer(tags)).Fsize) {
			break
		}
		if libc.BoolInt64(!(i < (*Vector)(unsafe.Pointer(tags)).Fsize)) != 0 {
			libc.X__assert_rtn(tls, uintptr(unsafe.Pointer(&__func__3)), __ccgo_ts+17, int32(239), __ccgo_ts+72)
		}
		v2 = (*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(i)*16
		if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
			libc.Xfree(tls, (*Array)(unsafe.Pointer(v2)).Fcontents)
			(*Array)(unsafe.Pointer(v2)).Fcontents = libc.UintptrFromInt32(0)
			(*Array)(unsafe.Pointer(v2)).Fsize = uint32(0)
			(*Array)(unsafe.Pointer(v2)).Fcapacity = uint32(0)
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	(*Vector)(unsafe.Pointer(tags)).Fsize = uint32(0)
	if length == uint32(0) {
		return
	}
	size = uint32(0)
	*(*uint32_t)(unsafe.Pointer(bp)) = uint32(0)
	*(*uint32_t)(unsafe.Pointer(bp + 4)) = uint32(0)
	libc.X__builtin___memcpy_chk(tls, bp+4, buffer+uintptr(size), uint64(4), ^__predefined_size_t(0))
	size = uint32(uint64(size) + libc.Uint64FromInt64(4))
	libc.X__builtin___memcpy_chk(tls, bp, buffer+uintptr(size), uint64(4), ^__predefined_size_t(0))
	size = uint32(uint64(size) + libc.Uint64FromInt64(4))
	if *(*uint32_t)(unsafe.Pointer(bp)) == uint32(0) {
		return
	}
	v2 = tags
	v4 = libc.Uint64FromInt64(16)
	v5 = *(*uint32_t)(unsafe.Pointer(bp))
	if v5 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
		if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
			(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v5)*v4)
		} else {
			(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v5)*v4)
		}
		(*Array)(unsafe.Pointer(v2)).Fcapacity = v5
	}
	iter = uint32(0)
	for {
		if !(iter < *(*uint32_t)(unsafe.Pointer(bp + 4))) {
			break
		}
		*(*String)(unsafe.Pointer(bp + 8)) = String{}
		v5 = size
		size = size + 1
		(*(*String)(unsafe.Pointer(bp + 8))).Fsize = uint32(libc.Uint8FromInt8(*(*int8)(unsafe.Pointer(buffer + uintptr(v5)))))
		if (*(*String)(unsafe.Pointer(bp + 8))).Fsize > uint32(0) {
			v2 = bp + 8
			v4 = libc.Uint64FromInt64(1)
			v5 = (*(*String)(unsafe.Pointer(bp + 8))).Fsize + uint32(1)
			if v5 > (*Array)(unsafe.Pointer(v2)).Fcapacity {
				if (*Array)(unsafe.Pointer(v2)).Fcontents != 0 {
					(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v2)).Fcontents, uint64(v5)*v4)
				} else {
					(*Array)(unsafe.Pointer(v2)).Fcontents = libc.Xmalloc(tls, uint64(v5)*v4)
				}
				(*Array)(unsafe.Pointer(v2)).Fcapacity = v5
			}
			libc.X__builtin___memcpy_chk(tls, (*(*String)(unsafe.Pointer(bp + 8))).Fcontents, buffer+uintptr(size), uint64((*(*String)(unsafe.Pointer(bp + 8))).Fsize), ^__predefined_size_t(0))
			size = size + (*(*String)(unsafe.Pointer(bp + 8))).Fsize
		}
		v2 = tags
		new_size = (*Array)(unsafe.Pointer(v2)).Fsize + uint32(1)
		if new_size > (*Array)(unsafe.Pointer(v2)).Fcapacity {
			new_capacity1 = (*Array)(unsafe.Pointer(v2)).Fcapacity * uint32(2)
			if new_capacity1 < libc.Uint32FromInt32(8) {
				new_capacity1 = uint32(8)
			}
			if new_capacity1 < new_size {
				new_capacity1 = new_size
			}
			v3 = v2
			v4 = libc.Uint64FromInt64(16)
			v5 = new_capacity1
			if v5 > (*Array)(unsafe.Pointer(v3)).Fcapacity {
				if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
					(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v3)).Fcontents, uint64(v5)*v4)
				} else {
					(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xmalloc(tls, uint64(v5)*v4)
				}
				(*Array)(unsafe.Pointer(v3)).Fcapacity = v5
			}
		}
		v8 = tags + 8
		v7 = *(*uint32_t)(unsafe.Pointer(v8))
		*(*uint32_t)(unsafe.Pointer(v8)) = *(*uint32_t)(unsafe.Pointer(v8)) + 1
		*(*String)(unsafe.Pointer((*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(v7)*16)) = *(*String)(unsafe.Pointer(bp + 8))
		goto _6
	_6:
		;
		iter = iter + 1
	}
	// add zero tags if we didn't read enough, this is because the
	// buffer had no more room but we held more tags.
	for {
		if !(iter < *(*uint32_t)(unsafe.Pointer(bp))) {
			break
		}
		tag1 = String{}
		v2 = tags
		new_size = (*Array)(unsafe.Pointer(v2)).Fsize + uint32(1)
		if new_size > (*Array)(unsafe.Pointer(v2)).Fcapacity {
			new_capacity1 = (*Array)(unsafe.Pointer(v2)).Fcapacity * uint32(2)
			if new_capacity1 < libc.Uint32FromInt32(8) {
				new_capacity1 = uint32(8)
			}
			if new_capacity1 < new_size {
				new_capacity1 = new_size
			}
			v3 = v2
			v4 = libc.Uint64FromInt64(16)
			v5 = new_capacity1
			if v5 > (*Array)(unsafe.Pointer(v3)).Fcapacity {
				if (*Array)(unsafe.Pointer(v3)).Fcontents != 0 {
					(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xrealloc(tls, (*Array)(unsafe.Pointer(v3)).Fcontents, uint64(v5)*v4)
				} else {
					(*Array)(unsafe.Pointer(v3)).Fcontents = libc.Xmalloc(tls, uint64(v5)*v4)
				}
				(*Array)(unsafe.Pointer(v3)).Fcapacity = v5
			}
		}
		v8 = tags + 8
		v7 = *(*uint32_t)(unsafe.Pointer(v8))
		*(*uint32_t)(unsafe.Pointer(v8)) = *(*uint32_t)(unsafe.Pointer(v8)) + 1
		*(*String)(unsafe.Pointer((*Vector)(unsafe.Pointer(tags)).Fcontents + uintptr(v7)*16)) = tag1
		goto _17
	_17:
		;
		iter = iter + 1
	}
}

var __func__3 = [45]int8{'t', 'r', 'e', 'e', '_', 's', 'i', 't', 't', 'e', 'r', '_', 'x', 'm', 'l', '_', 'e', 'x', 't', 'e', 'r', 'n', 'a', 'l', '_', 's', 'c', 'a', 'n', 'n', 'e', 'r', '_', 'd', 'e', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e'}

/* Automatically @generated by tree-sitter v0.25.4 (683a5044a8f496d8e8b57c06ac4f1a722c62205f) */

type ts_symbol_identifiers = int32

const sym_Name = 1
const anon_sym_LT_QMARK = 2
const anon_sym_xml = 3
const anon_sym_QMARK_GT = 4
const anon_sym_standalone = 5
const anon_sym_SQUOTE = 6
const anon_sym_yes = 7
const anon_sym_no = 8
const anon_sym_DQUOTE = 9
const anon_sym_LT_BANG = 10
const anon_sym_DOCTYPE = 11
const anon_sym_LBRACK = 12
const anon_sym_RBRACK = 13
const anon_sym_GT = 14
const anon_sym_LT = 15
const anon_sym_SLASH_GT = 16
const anon_sym_LT_SLASH = 17
const anon_sym_RBRACK_RBRACK_GT = 18
const anon_sym_LT_BANG_LBRACK = 19
const anon_sym_CDATA = 20
const anon_sym_xml_DASHstylesheet = 21
const anon_sym_xml_DASHmodel = 22
const aux_sym_PseudoAttValue_token1 = 23
const aux_sym_PseudoAttValue_token2 = 24
const anon_sym_ELEMENT = 25
const anon_sym_EMPTY = 26
const anon_sym_ANY = 27
const anon_sym_LPAREN = 28
const anon_sym_POUNDPCDATA = 29
const anon_sym_PIPE = 30
const anon_sym_RPAREN = 31
const anon_sym_STAR = 32
const anon_sym_QMARK = 33
const anon_sym_PLUS = 34
const anon_sym_COMMA = 35
const anon_sym_ATTLIST = 36
const sym_TokenizedType = 37
const anon_sym_NOTATION = 38
const anon_sym_POUNDREQUIRED = 39
const anon_sym_POUNDIMPLIED = 40
const anon_sym_POUNDFIXED = 41
const anon_sym_ENTITY = 42
const anon_sym_PERCENT = 43
const aux_sym_EntityValue_token1 = 44
const aux_sym_EntityValue_token2 = 45
const anon_sym_NDATA = 46
const anon_sym_SEMI = 47
const sym__S = 48
const sym_Nmtoken = 49
const anon_sym_AMP = 50
const anon_sym_AMP_POUND = 51
const aux_sym_CharRef_token1 = 52
const anon_sym_AMP_POUNDx = 53
const aux_sym_CharRef_token2 = 54
const anon_sym_SYSTEM = 55
const anon_sym_PUBLIC = 56
const aux_sym_SystemLiteral_token1 = 57
const aux_sym_SystemLiteral_token2 = 58
const aux_sym_PubidLiteral_token1 = 59
const aux_sym_PubidLiteral_token2 = 60
const anon_sym_version = 61
const sym_VersionNum = 62
const anon_sym_encoding = 63
const sym_EncName = 64
const anon_sym_EQ = 65
const sym_PITarget = 66
const sym__pi_content = 67
const sym_Comment = 68
const sym_CharData = 69
const sym_CData = 70
const sym__start_tag_name = 71
const sym__end_tag_name = 72
const sym__erroneous_end_name = 73
const sym_document = 74
const sym_prolog = 75
const sym__Misc = 76
const sym_XMLDecl = 77
const sym__SDDecl = 78
const sym_doctypedecl = 79
const aux_sym__intSubset = 80
const sym_element = 81
const sym_EmptyElemTag = 82
const sym_Attribute = 83
const sym_STag = 84
const sym_ETag = 85
const sym_content = 86
const sym_CDSect = 87
const sym_CDStart = 88
const sym_StyleSheetPI = 89
const sym_XmlModelPI = 90
const sym_PseudoAtt = 91
const sym_PseudoAttValue = 92
const sym__markupdecl = 93
const sym__DeclSep = 94
const sym_elementdecl = 95
const sym_contentspec = 96
const sym_Mixed = 97
const sym_children = 98
const sym__cp = 99
const sym__choice = 100
const sym_AttlistDecl = 101
const sym_AttDef = 102
const sym__AttType = 103
const sym_StringType = 104
const sym__EnumeratedType = 105
const sym_NotationType = 106
const sym_Enumeration = 107
const sym_DefaultDecl = 108
const sym__EntityDecl = 109
const sym_GEDecl = 110
const sym_PEDecl = 111
const sym_EntityValue = 112
const sym_NDataDecl = 113
const sym_NotationDecl = 114
const sym_PEReference = 115
const sym__Reference = 116
const sym_EntityRef = 117
const sym_CharRef = 118
const sym_AttValue = 119
const sym_ExternalID = 120
const sym_PublicID = 121
const sym_SystemLiteral = 122
const sym_PubidLiteral = 123
const sym__VersionInfo = 124
const sym__EncodingDecl = 125
const sym_PI = 126
const sym__Eq = 127
const aux_sym_document_repeat1 = 128
const aux_sym_EmptyElemTag_repeat1 = 129
const aux_sym_content_repeat1 = 130
const aux_sym_StyleSheetPI_repeat1 = 131
const aux_sym_PseudoAttValue_repeat1 = 132
const aux_sym_PseudoAttValue_repeat2 = 133
const aux_sym_Mixed_repeat1 = 134
const aux_sym_Mixed_repeat2 = 135
const aux_sym__choice_repeat1 = 136
const aux_sym__choice_repeat2 = 137
const aux_sym_AttlistDecl_repeat1 = 138
const aux_sym_NotationType_repeat1 = 139
const aux_sym_Enumeration_repeat1 = 140
const aux_sym_EntityValue_repeat1 = 141
const aux_sym_EntityValue_repeat2 = 142

var ts_symbol_names = [143]uintptr{
	0:   __ccgo_ts + 149,
	1:   __ccgo_ts + 153,
	2:   __ccgo_ts + 158,
	3:   __ccgo_ts + 161,
	4:   __ccgo_ts + 165,
	5:   __ccgo_ts + 168,
	6:   __ccgo_ts + 179,
	7:   __ccgo_ts + 181,
	8:   __ccgo_ts + 185,
	9:   __ccgo_ts + 188,
	10:  __ccgo_ts + 190,
	11:  __ccgo_ts + 193,
	12:  __ccgo_ts + 201,
	13:  __ccgo_ts + 203,
	14:  __ccgo_ts + 205,
	15:  __ccgo_ts + 207,
	16:  __ccgo_ts + 209,
	17:  __ccgo_ts + 212,
	18:  __ccgo_ts + 215,
	19:  __ccgo_ts + 219,
	20:  __ccgo_ts + 223,
	21:  __ccgo_ts + 229,
	22:  __ccgo_ts + 244,
	23:  __ccgo_ts + 254,
	24:  __ccgo_ts + 276,
	25:  __ccgo_ts + 298,
	26:  __ccgo_ts + 306,
	27:  __ccgo_ts + 312,
	28:  __ccgo_ts + 316,
	29:  __ccgo_ts + 318,
	30:  __ccgo_ts + 326,
	31:  __ccgo_ts + 328,
	32:  __ccgo_ts + 330,
	33:  __ccgo_ts + 332,
	34:  __ccgo_ts + 334,
	35:  __ccgo_ts + 336,
	36:  __ccgo_ts + 338,
	37:  __ccgo_ts + 346,
	38:  __ccgo_ts + 360,
	39:  __ccgo_ts + 369,
	40:  __ccgo_ts + 379,
	41:  __ccgo_ts + 388,
	42:  __ccgo_ts + 395,
	43:  __ccgo_ts + 402,
	44:  __ccgo_ts + 404,
	45:  __ccgo_ts + 423,
	46:  __ccgo_ts + 442,
	47:  __ccgo_ts + 448,
	48:  __ccgo_ts + 450,
	49:  __ccgo_ts + 453,
	50:  __ccgo_ts + 461,
	51:  __ccgo_ts + 463,
	52:  __ccgo_ts + 466,
	53:  __ccgo_ts + 481,
	54:  __ccgo_ts + 485,
	55:  __ccgo_ts + 500,
	56:  __ccgo_ts + 507,
	57:  __ccgo_ts + 514,
	58:  __ccgo_ts + 514,
	59:  __ccgo_ts + 518,
	60:  __ccgo_ts + 538,
	61:  __ccgo_ts + 558,
	62:  __ccgo_ts + 566,
	63:  __ccgo_ts + 577,
	64:  __ccgo_ts + 586,
	65:  __ccgo_ts + 594,
	66:  __ccgo_ts + 596,
	67:  __ccgo_ts + 605,
	68:  __ccgo_ts + 617,
	69:  __ccgo_ts + 625,
	70:  __ccgo_ts + 634,
	71:  __ccgo_ts + 153,
	72:  __ccgo_ts + 153,
	73:  __ccgo_ts + 640,
	74:  __ccgo_ts + 660,
	75:  __ccgo_ts + 669,
	76:  __ccgo_ts + 676,
	77:  __ccgo_ts + 682,
	78:  __ccgo_ts + 690,
	79:  __ccgo_ts + 698,
	80:  __ccgo_ts + 710,
	81:  __ccgo_ts + 721,
	82:  __ccgo_ts + 729,
	83:  __ccgo_ts + 742,
	84:  __ccgo_ts + 752,
	85:  __ccgo_ts + 757,
	86:  __ccgo_ts + 762,
	87:  __ccgo_ts + 770,
	88:  __ccgo_ts + 777,
	89:  __ccgo_ts + 785,
	90:  __ccgo_ts + 798,
	91:  __ccgo_ts + 809,
	92:  __ccgo_ts + 819,
	93:  __ccgo_ts + 834,
	94:  __ccgo_ts + 846,
	95:  __ccgo_ts + 855,
	96:  __ccgo_ts + 867,
	97:  __ccgo_ts + 879,
	98:  __ccgo_ts + 885,
	99:  __ccgo_ts + 894,
	100: __ccgo_ts + 898,
	101: __ccgo_ts + 906,
	102: __ccgo_ts + 918,
	103: __ccgo_ts + 925,
	104: __ccgo_ts + 934,
	105: __ccgo_ts + 945,
	106: __ccgo_ts + 961,
	107: __ccgo_ts + 974,
	108: __ccgo_ts + 986,
	109: __ccgo_ts + 998,
	110: __ccgo_ts + 1010,
	111: __ccgo_ts + 1017,
	112: __ccgo_ts + 1024,
	113: __ccgo_ts + 1036,
	114: __ccgo_ts + 1046,
	115: __ccgo_ts + 1059,
	116: __ccgo_ts + 1071,
	117: __ccgo_ts + 1082,
	118: __ccgo_ts + 1092,
	119: __ccgo_ts + 1100,
	120: __ccgo_ts + 1109,
	121: __ccgo_ts + 1120,
	122: __ccgo_ts + 1129,
	123: __ccgo_ts + 1143,
	124: __ccgo_ts + 1156,
	125: __ccgo_ts + 1169,
	126: __ccgo_ts + 1183,
	127: __ccgo_ts + 1186,
	128: __ccgo_ts + 1190,
	129: __ccgo_ts + 1207,
	130: __ccgo_ts + 1228,
	131: __ccgo_ts + 1244,
	132: __ccgo_ts + 1265,
	133: __ccgo_ts + 1288,
	134: __ccgo_ts + 1311,
	135: __ccgo_ts + 1325,
	136: __ccgo_ts + 1339,
	137: __ccgo_ts + 1355,
	138: __ccgo_ts + 1371,
	139: __ccgo_ts + 1391,
	140: __ccgo_ts + 1412,
	141: __ccgo_ts + 1432,
	142: __ccgo_ts + 1452,
}

var ts_symbol_map = [143]TSSymbol{
	1:   uint16(sym_Name),
	2:   uint16(anon_sym_LT_QMARK),
	3:   uint16(anon_sym_xml),
	4:   uint16(anon_sym_QMARK_GT),
	5:   uint16(anon_sym_standalone),
	6:   uint16(anon_sym_SQUOTE),
	7:   uint16(anon_sym_yes),
	8:   uint16(anon_sym_no),
	9:   uint16(anon_sym_DQUOTE),
	10:  uint16(anon_sym_LT_BANG),
	11:  uint16(anon_sym_DOCTYPE),
	12:  uint16(anon_sym_LBRACK),
	13:  uint16(anon_sym_RBRACK),
	14:  uint16(anon_sym_GT),
	15:  uint16(anon_sym_LT),
	16:  uint16(anon_sym_SLASH_GT),
	17:  uint16(anon_sym_LT_SLASH),
	18:  uint16(anon_sym_RBRACK_RBRACK_GT),
	19:  uint16(anon_sym_LT_BANG_LBRACK),
	20:  uint16(anon_sym_CDATA),
	21:  uint16(anon_sym_xml_DASHstylesheet),
	22:  uint16(anon_sym_xml_DASHmodel),
	23:  uint16(aux_sym_PseudoAttValue_token1),
	24:  uint16(aux_sym_PseudoAttValue_token2),
	25:  uint16(anon_sym_ELEMENT),
	26:  uint16(anon_sym_EMPTY),
	27:  uint16(anon_sym_ANY),
	28:  uint16(anon_sym_LPAREN),
	29:  uint16(anon_sym_POUNDPCDATA),
	30:  uint16(anon_sym_PIPE),
	31:  uint16(anon_sym_RPAREN),
	32:  uint16(anon_sym_STAR),
	33:  uint16(anon_sym_QMARK),
	34:  uint16(anon_sym_PLUS),
	35:  uint16(anon_sym_COMMA),
	36:  uint16(anon_sym_ATTLIST),
	37:  uint16(sym_TokenizedType),
	38:  uint16(anon_sym_NOTATION),
	39:  uint16(anon_sym_POUNDREQUIRED),
	40:  uint16(anon_sym_POUNDIMPLIED),
	41:  uint16(anon_sym_POUNDFIXED),
	42:  uint16(anon_sym_ENTITY),
	43:  uint16(anon_sym_PERCENT),
	44:  uint16(aux_sym_EntityValue_token1),
	45:  uint16(aux_sym_EntityValue_token2),
	46:  uint16(anon_sym_NDATA),
	47:  uint16(anon_sym_SEMI),
	48:  uint16(sym__S),
	49:  uint16(sym_Nmtoken),
	50:  uint16(anon_sym_AMP),
	51:  uint16(anon_sym_AMP_POUND),
	52:  uint16(aux_sym_CharRef_token1),
	53:  uint16(anon_sym_AMP_POUNDx),
	54:  uint16(aux_sym_CharRef_token2),
	55:  uint16(anon_sym_SYSTEM),
	56:  uint16(anon_sym_PUBLIC),
	57:  uint16(aux_sym_SystemLiteral_token1),
	58:  uint16(aux_sym_SystemLiteral_token1),
	59:  uint16(aux_sym_PubidLiteral_token1),
	60:  uint16(aux_sym_PubidLiteral_token2),
	61:  uint16(anon_sym_version),
	62:  uint16(sym_VersionNum),
	63:  uint16(anon_sym_encoding),
	64:  uint16(sym_EncName),
	65:  uint16(anon_sym_EQ),
	66:  uint16(sym_PITarget),
	67:  uint16(sym__pi_content),
	68:  uint16(sym_Comment),
	69:  uint16(sym_CharData),
	70:  uint16(sym_CData),
	71:  uint16(sym_Name),
	72:  uint16(sym_Name),
	73:  uint16(sym__erroneous_end_name),
	74:  uint16(sym_document),
	75:  uint16(sym_prolog),
	76:  uint16(sym__Misc),
	77:  uint16(sym_XMLDecl),
	78:  uint16(sym__SDDecl),
	79:  uint16(sym_doctypedecl),
	80:  uint16(aux_sym__intSubset),
	81:  uint16(sym_element),
	82:  uint16(sym_EmptyElemTag),
	83:  uint16(sym_Attribute),
	84:  uint16(sym_STag),
	85:  uint16(sym_ETag),
	86:  uint16(sym_content),
	87:  uint16(sym_CDSect),
	88:  uint16(sym_CDStart),
	89:  uint16(sym_StyleSheetPI),
	90:  uint16(sym_XmlModelPI),
	91:  uint16(sym_PseudoAtt),
	92:  uint16(sym_PseudoAttValue),
	93:  uint16(sym__markupdecl),
	94:  uint16(sym__DeclSep),
	95:  uint16(sym_elementdecl),
	96:  uint16(sym_contentspec),
	97:  uint16(sym_Mixed),
	98:  uint16(sym_children),
	99:  uint16(sym__cp),
	100: uint16(sym__choice),
	101: uint16(sym_AttlistDecl),
	102: uint16(sym_AttDef),
	103: uint16(sym__AttType),
	104: uint16(sym_StringType),
	105: uint16(sym__EnumeratedType),
	106: uint16(sym_NotationType),
	107: uint16(sym_Enumeration),
	108: uint16(sym_DefaultDecl),
	109: uint16(sym__EntityDecl),
	110: uint16(sym_GEDecl),
	111: uint16(sym_PEDecl),
	112: uint16(sym_EntityValue),
	113: uint16(sym_NDataDecl),
	114: uint16(sym_NotationDecl),
	115: uint16(sym_PEReference),
	116: uint16(sym__Reference),
	117: uint16(sym_EntityRef),
	118: uint16(sym_CharRef),
	119: uint16(sym_AttValue),
	120: uint16(sym_ExternalID),
	121: uint16(sym_PublicID),
	122: uint16(sym_SystemLiteral),
	123: uint16(sym_PubidLiteral),
	124: uint16(sym__VersionInfo),
	125: uint16(sym__EncodingDecl),
	126: uint16(sym_PI),
	127: uint16(sym__Eq),
	128: uint16(aux_sym_document_repeat1),
	129: uint16(aux_sym_EmptyElemTag_repeat1),
	130: uint16(aux_sym_content_repeat1),
	131: uint16(aux_sym_StyleSheetPI_repeat1),
	132: uint16(aux_sym_PseudoAttValue_repeat1),
	133: uint16(aux_sym_PseudoAttValue_repeat2),
	134: uint16(aux_sym_Mixed_repeat1),
	135: uint16(aux_sym_Mixed_repeat2),
	136: uint16(aux_sym__choice_repeat1),
	137: uint16(aux_sym__choice_repeat2),
	138: uint16(aux_sym_AttlistDecl_repeat1),
	139: uint16(aux_sym_NotationType_repeat1),
	140: uint16(aux_sym_Enumeration_repeat1),
	141: uint16(aux_sym_EntityValue_repeat1),
	142: uint16(aux_sym_EntityValue_repeat2),
}

var ts_symbol_metadata = [143]TSSymbolMetadata{
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
	22: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	23: {},
	24: {},
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
		Fnamed:   libc.BoolUint8(true1 != 0),
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
	44: {},
	45: {},
	46: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	47: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	48: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	49: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	50: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	51: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	52: {},
	53: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	54: {},
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
	59: {},
	60: {},
	61: {
		Fvisible: libc.BoolUint8(true1 != 0),
	},
	62: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	79: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	80: {},
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	94: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	95: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	96: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	97: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	98: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	99: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	100: {
		Fnamed: libc.BoolUint8(true1 != 0),
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
	},
	104: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	105: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	115: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	116: {
		Fnamed:     libc.BoolUint8(true1 != 0),
		Fsupertype: libc.BoolUint8(true1 != 0),
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
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
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
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	125: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	126: {
		Fvisible: libc.BoolUint8(true1 != 0),
		Fnamed:   libc.BoolUint8(true1 != 0),
	},
	127: {
		Fnamed: libc.BoolUint8(true1 != 0),
	},
	128: {},
	129: {},
	130: {},
	131: {},
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
}

type ts_field_identifiers = int32

const field_content = 1
const field_root = 2

var ts_field_names = [3]uintptr{
	0: libc.UintptrFromInt32(0),
	1: __ccgo_ts + 762,
	2: __ccgo_ts + 1472,
}

var ts_field_map_slices = [5]TSMapSlice{
	1: {
		Flength: uint16(1),
	},
	2: {
		Findex:  uint16(1),
		Flength: uint16(1),
	},
	3: {
		Findex:  uint16(2),
		Flength: uint16(1),
	},
	4: {
		Findex:  uint16(3),
		Flength: uint16(1),
	},
}

var ts_field_map_entries = [4]TSFieldMapEntry{
	0: {
		Ffield_id: uint16(field_root),
	},
	1: {
		Ffield_id:    uint16(field_root),
		Fchild_index: uint8(1),
	},
	2: {
		Ffield_id:    uint16(field_root),
		Fchild_index: uint8(2),
	},
	3: {
		Ffield_id:    uint16(field_content),
		Fchild_index: uint8(1),
	},
}

var ts_alias_sequences = [5][12]TSSymbol{}

var ts_non_terminal_alias_map = [1]uint16_t{}

var ts_primary_state_ids = [495]TSStateId{
	1:   uint16(1),
	2:   uint16(2),
	3:   uint16(2),
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
	72:  uint16(72),
	73:  uint16(73),
	74:  uint16(44),
	75:  uint16(75),
	76:  uint16(76),
	77:  uint16(77),
	78:  uint16(78),
	79:  uint16(79),
	80:  uint16(80),
	81:  uint16(81),
	82:  uint16(81),
	83:  uint16(60),
	84:  uint16(43),
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
	100: uint16(80),
	101: uint16(101),
	102: uint16(102),
	103: uint16(103),
	104: uint16(104),
	105: uint16(105),
	106: uint16(106),
	107: uint16(72),
	108: uint16(108),
	109: uint16(109),
	110: uint16(110),
	111: uint16(77),
	112: uint16(112),
	113: uint16(113),
	114: uint16(114),
	115: uint16(115),
	116: uint16(116),
	117: uint16(117),
	118: uint16(118),
	119: uint16(119),
	120: uint16(34),
	121: uint16(56),
	122: uint16(55),
	123: uint16(56),
	124: uint16(34),
	125: uint16(80),
	126: uint16(55),
	127: uint16(127),
	128: uint16(128),
	129: uint16(129),
	130: uint16(130),
	131: uint16(131),
	132: uint16(132),
	133: uint16(133),
	134: uint16(34),
	135: uint16(135),
	136: uint16(136),
	137: uint16(137),
	138: uint16(138),
	139: uint16(139),
	140: uint16(140),
	141: uint16(141),
	142: uint16(80),
	143: uint16(143),
	144: uint16(55),
	145: uint16(56),
	146: uint16(146),
	147: uint16(147),
	148: uint16(80),
	149: uint16(55),
	150: uint16(56),
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
	168: uint16(75),
	169: uint16(169),
	170: uint16(170),
	171: uint16(171),
	172: uint16(172),
	173: uint16(173),
	174: uint16(174),
	175: uint16(73),
	176: uint16(176),
	177: uint16(76),
	178: uint16(178),
	179: uint16(179),
	180: uint16(180),
	181: uint16(181),
	182: uint16(182),
	183: uint16(183),
	184: uint16(184),
	185: uint16(185),
	186: uint16(186),
	187: uint16(79),
	188: uint16(188),
	189: uint16(70),
	190: uint16(190),
	191: uint16(191),
	192: uint16(192),
	193: uint16(193),
	194: uint16(35),
	195: uint16(195),
	196: uint16(196),
	197: uint16(197),
	198: uint16(198),
	199: uint16(199),
	200: uint16(200),
	201: uint16(201),
	202: uint16(71),
	203: uint16(203),
	204: uint16(204),
	205: uint16(205),
	206: uint16(206),
	207: uint16(207),
	208: uint16(208),
	209: uint16(209),
	210: uint16(210),
	211: uint16(211),
	212: uint16(178),
	213: uint16(184),
	214: uint16(185),
	215: uint16(169),
	216: uint16(78),
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
	310: uint16(310),
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
	337: uint16(337),
	338: uint16(338),
	339: uint16(339),
	340: uint16(340),
	341: uint16(341),
	342: uint16(342),
	343: uint16(343),
	344: uint16(344),
	345: uint16(345),
	346: uint16(346),
	347: uint16(240),
	348: uint16(270),
	349: uint16(349),
	350: uint16(350),
	351: uint16(72),
	352: uint16(352),
	353: uint16(353),
	354: uint16(77),
	355: uint16(355),
	356: uint16(356),
	357: uint16(357),
	358: uint16(358),
	359: uint16(359),
	360: uint16(360),
	361: uint16(361),
	362: uint16(362),
	363: uint16(363),
	364: uint16(334),
	365: uint16(365),
	366: uint16(350),
	367: uint16(367),
	368: uint16(368),
	369: uint16(277),
	370: uint16(370),
	371: uint16(371),
	372: uint16(372),
	373: uint16(373),
	374: uint16(374),
	375: uint16(334),
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
	405: uint16(405),
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
	454: uint16(408),
	455: uint16(410),
	456: uint16(439),
	457: uint16(457),
	458: uint16(445),
	459: uint16(459),
	460: uint16(460),
	461: uint16(461),
	462: uint16(462),
	463: uint16(398),
	464: uint16(464),
	465: uint16(408),
	466: uint16(410),
	467: uint16(439),
	468: uint16(398),
	469: uint16(408),
	470: uint16(410),
	471: uint16(398),
	472: uint16(408),
	473: uint16(410),
	474: uint16(426),
	475: uint16(441),
	476: uint16(442),
	477: uint16(462),
	478: uint16(437),
	479: uint16(378),
	480: uint16(446),
	481: uint16(481),
	482: uint16(462),
	483: uint16(437),
	484: uint16(378),
	485: uint16(446),
	486: uint16(481),
	487: uint16(462),
	488: uint16(437),
	489: uint16(378),
	490: uint16(481),
	491: uint16(462),
	492: uint16(437),
	493: uint16(378),
	494: uint16(481),
}

var aux_sym_PubidLiteral_token1_character_set_1 = [9]TSCharacterRange{
	0: {
		Fstart: int32('\n'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32('\r'),
		Fend:   int32('\r'),
	},
	2: {
		Fstart: int32(' '),
		Fend:   int32('!'),
	},
	3: {
		Fstart: int32('#'),
		Fend:   int32('%'),
	},
	4: {
		Fstart: int32('\''),
		Fend:   int32(';'),
	},
	5: {
		Fstart: int32('='),
		Fend:   int32('='),
	},
	6: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	7: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	8: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
}

var aux_sym_PubidLiteral_token2_character_set_1 = [9]TSCharacterRange{
	0: {
		Fstart: int32('\n'),
		Fend:   int32('\n'),
	},
	1: {
		Fstart: int32('\r'),
		Fend:   int32('\r'),
	},
	2: {
		Fstart: int32(' '),
		Fend:   int32('!'),
	},
	3: {
		Fstart: int32('#'),
		Fend:   int32('%'),
	},
	4: {
		Fstart: int32('('),
		Fend:   int32(';'),
	},
	5: {
		Fstart: int32('='),
		Fend:   int32('='),
	},
	6: {
		Fstart: int32('?'),
		Fend:   int32('Z'),
	},
	7: {
		Fstart: int32('_'),
		Fend:   int32('_'),
	},
	8: {
		Fstart: int32('a'),
		Fend:   int32('z'),
	},
}

func ts_lex(tls *libc.TLS, lexer uintptr, state TSStateId) (r uint8) {
	var eof, result, skip, v6 uint8
	var half_size, i, i1, i2, index, mid_index, size uint32_t
	var lookahead1, v5 int32_t
	var range_token, range_token1, v4 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = eof, half_size, i, i1, i2, index, lookahead1, mid_index, range_token, range_token1, result, size, skip, v4, v5, v6
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
			state = uint16(42)
			goto next_state
		}
		i = uint32(0)
		for {
			if !(uint64(i) < libc.Uint64FromInt64(128)/libc.Uint64FromInt64(2)) {
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
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(70)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(67)
			goto next_state
		}
		if int32('G') <= lookahead1 && lookahead1 <= int32('Z') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(68)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(1):
		i1 = uint32(0)
		for {
			if !(uint64(i1) < libc.Uint64FromInt64(76)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token1[i1]) == lookahead1 {
				state = map_token1[i1+uint32(1)]
				goto next_state
			}
			goto _2
		_2:
			;
			i1 = i1 + uint32(2)
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(134)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead1 == int32('"') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('%') {
			state = uint16(91)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(132)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('<') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead1 == int32('"') {
			state = uint16(46)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(132)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('<') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead1 == int32('%') {
			state = uint16(91)
			goto next_state
		}
		if lookahead1 == int32('&') {
			state = uint16(132)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(45)
			goto next_state
		}
		if lookahead1 == int32('<') {
			state = uint16(55)
			goto next_state
		}
		if lookahead1 != 0 {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead1 == int32('&') {
			state = uint16(132)
			goto next_state
		}
		if lookahead1 == int32('\'') {
			state = uint16(45)
			goto next_state
		}
		if lookahead1 != 0 && lookahead1 != int32('<') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead1 == int32('.') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead1 == int32('>') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead1 == int32('>') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead1 == int32('>') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead1 == int32('?') {
			state = uint16(9)
			goto next_state
		}
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(95)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead1 == int32('A') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead1 == int32('A') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead1 == int32('C') {
			state = uint16(14)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead1 == int32('D') {
			state = uint16(11)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead1 == int32('D') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead1 == int32('D') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead1 == int32('D') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead1 == int32('E') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead1 == int32('E') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead1 == int32('E') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead1 == int32('E') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead1 == int32('F') {
			state = uint16(23)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead1 == int32('I') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead1 == int32('I') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead1 == int32('I') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(26):
		if lookahead1 == int32('L') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(27):
		if lookahead1 == int32('M') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead1 == int32('P') {
			state = uint16(13)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead1 == int32('P') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead1 == int32('Q') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(31):
		if lookahead1 == int32('R') {
			state = uint16(21)
			goto next_state
		}
		return result
	case int32(32):
		if lookahead1 == int32('T') {
			state = uint16(12)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead1 == int32('U') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead1 == int32('X') {
			state = uint16(19)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead1 == int32('[') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead1 == int32(']') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(95)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(38):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(39):
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(40):
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(41):
		if eof != 0 {
			state = uint16(42)
			goto next_state
		}
		i2 = uint32(0)
		for {
			if !(uint64(i2) < libc.Uint64FromInt64(104)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token2[i2]) == lookahead1 {
				state = map_token2[i2+uint32(1)]
				goto next_state
			}
			goto _3
		_3:
			;
			i2 = i2 + uint32(2)
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(42):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(ts_builtin_sym_end)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(43):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(44):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DQUOTE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(48):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('[') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(49):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(50):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(51):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(']') {
			state = uint16(8)
			goto next_state
		}
		return result
	case int32(52):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(53):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(48)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(57)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(54):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(47)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(55):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('!') {
			state = uint16(35)
			goto next_state
		}
		if lookahead1 == int32('/') {
			state = uint16(57)
			goto next_state
		}
		if lookahead1 == int32('?') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(56):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SLASH_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(57):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_SLASH)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(58):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RBRACK_RBRACK_GT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(59):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LT_BANG_LBRACK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(60):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(61):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('.') {
			state = uint16(128)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(129)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(130)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('>') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(63):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('D') {
			state = uint16(82)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(64):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(23)
			goto next_state
		}
		if lookahead1 == int32('I') {
			state = uint16(27)
			goto next_state
		}
		if lookahead1 == int32('P') {
			state = uint16(13)
			goto next_state
		}
		if lookahead1 == int32('R') {
			state = uint16(18)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('M') {
			state = uint16(120)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(66):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(119)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(125)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(125)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(68):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(69):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(70):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(129)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(130)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(71):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(72):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(73):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PseudoAttValue_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(74):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_LPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(75):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDPCDATA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(76):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PIPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(77):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_RPAREN)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(78):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_STAR)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_QMARK)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PLUS)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_COMMA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('R') {
			state = uint16(97)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(83):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('R') {
			state = uint16(100)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(84):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(86)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(85):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(86):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(87):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_TokenizedType)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDREQUIRED)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDIMPLIED)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_POUNDFIXED)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PERCENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(92):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(93):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_EntityValue_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SEMI)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym__S)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('\t') || lookahead1 == int32('\n') || lookahead1 == int32('\r') || lookahead1 == int32(' ') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(96):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('D') {
			state = uint16(83)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(103)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(112)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(99):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(117)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(100):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(104)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(101):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(118)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(102):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('E') {
			state = uint16(114)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(104):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('F') {
			state = uint16(85)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(121)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(106):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(99)
			goto next_state
		}
		if lookahead1 == int32('Y') {
			state = uint16(86)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(107):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(124)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(108):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('I') {
			state = uint16(101)
			goto next_state
		}
		if lookahead1 == int32('Y') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(109):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('K') {
			state = uint16(98)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('K') {
			state = uint16(102)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(111):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('M') {
			state = uint16(123)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(112):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(84)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(113):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(122)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(114):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('N') {
			state = uint16(85)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(115):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(109)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(116):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('O') {
			state = uint16(110)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(117):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(86)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(118):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('S') {
			state = uint16(87)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(119):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(105)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(120):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(115)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(121):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(106)
			goto next_state
		}
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(122):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(107)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(123):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(116)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(124):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('T') {
			state = uint16(108)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(125):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(125)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(126):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32(':') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(126)
			goto next_state
		}
		return result
	case int32(127):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Name)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(127)
			goto next_state
		}
		return result
	case int32(128):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(128)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(129):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(129)
			goto next_state
		}
		if int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(130)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(130):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(130)
			goto next_state
		}
		if lookahead1 == int32('-') || lookahead1 == int32('.') || lookahead1 == int32(':') || int32('G') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('g') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(131):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_Nmtoken)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32(':') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') || lookahead1 == int32(0xb7) {
			state = uint16(131)
			goto next_state
		}
		return result
	case int32(132):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('#') {
			state = uint16(133)
			goto next_state
		}
		return result
	case int32(133):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_POUND)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('x') {
			state = uint16(135)
			goto next_state
		}
		return result
	case int32(134):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_CharRef_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(134)
			goto next_state
		}
		return result
	case int32(135):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_AMP_POUNDx)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(136):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_CharRef_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('F') || int32('a') <= lookahead1 && lookahead1 <= int32('f') {
			state = uint16(136)
			goto next_state
		}
		return result
	case int32(137):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_SystemLiteral_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('"') {
			state = uint16(137)
			goto next_state
		}
		return result
	case int32(138):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_SystemLiteral_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 != 0 && lookahead1 != int32('\'') {
			state = uint16(138)
			goto next_state
		}
		return result
	case int32(139):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PubidLiteral_token1)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v4 = uintptr(unsafe.Pointer(&aux_sym_PubidLiteral_token1_character_set_1))
		v5 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v4 + uintptr(mid_index)*8
			if v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v6 = libc.BoolUint8(true1 != 0)
				goto _7
			} else {
				if v5 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v4 + uintptr(index)*8
		v6 = libc.BoolUint8(v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _7
	_7:
		if v6 != 0 {
			state = uint16(139)
			goto next_state
		}
		return result
	case int32(140):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(aux_sym_PubidLiteral_token2)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		v4 = uintptr(unsafe.Pointer(&aux_sym_PubidLiteral_token2_character_set_1))
		v5 = lookahead1
		index = uint32(0)
		size = uint32(9) - index
		for size > libc.Uint32FromInt32(1) {
			half_size = size / uint32(2)
			mid_index = index + half_size
			range_token = v4 + uintptr(mid_index)*8
			if v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
				v6 = libc.BoolUint8(true1 != 0)
				goto _11
			} else {
				if v5 > (*TSCharacterRange)(unsafe.Pointer(range_token)).Fend {
					index = mid_index
				}
			}
			size = size - half_size
		}
		range_token1 = v4 + uintptr(index)*8
		v6 = libc.BoolUint8(v5 >= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fstart && v5 <= (*TSCharacterRange)(unsafe.Pointer(range_token1)).Fend)
		goto _11
	_11:
		if v6 != 0 {
			state = uint16(140)
			goto next_state
		}
		return result
	case int32(141):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_VersionNum)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if int32('0') <= lookahead1 && lookahead1 <= int32('9') {
			state = uint16(141)
			goto next_state
		}
		return result
	case int32(142):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(sym_EncName)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead1 == int32('-') || lookahead1 == int32('.') || int32('0') <= lookahead1 && lookahead1 <= int32('9') || int32('A') <= lookahead1 && lookahead1 <= int32('Z') || lookahead1 == int32('_') || int32('a') <= lookahead1 && lookahead1 <= int32('z') {
			state = uint16(142)
			goto next_state
		}
		return result
	case int32(143):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EQ)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token = [64]uint16_t{
	0:  uint16('"'),
	1:  uint16(46),
	2:  uint16('#'),
	3:  uint16(64),
	4:  uint16('%'),
	5:  uint16(91),
	6:  uint16('&'),
	7:  uint16(132),
	8:  uint16('\''),
	9:  uint16(45),
	10: uint16('('),
	11: uint16(74),
	12: uint16(')'),
	13: uint16(77),
	14: uint16('*'),
	15: uint16(78),
	16: uint16('+'),
	17: uint16(80),
	18: uint16(','),
	19: uint16(81),
	20: uint16('/'),
	21: uint16(62),
	22: uint16('1'),
	23: uint16(61),
	24: uint16(';'),
	25: uint16(94),
	26: uint16('<'),
	27: uint16(53),
	28: uint16('='),
	29: uint16(143),
	30: uint16('>'),
	31: uint16(52),
	32: uint16('?'),
	33: uint16(79),
	34: uint16('E'),
	35: uint16(66),
	36: uint16('I'),
	37: uint16(63),
	38: uint16('N'),
	39: uint16(65),
	40: uint16('['),
	41: uint16(49),
	42: uint16(']'),
	43: uint16(51),
	44: uint16('_'),
	45: uint16(72),
	46: uint16('|'),
	47: uint16(76),
	48: uint16('\t'),
	49: uint16(69),
	50: uint16('\n'),
	51: uint16(69),
	52: uint16('\r'),
	53: uint16(69),
	54: uint16(' '),
	55: uint16(69),
	56: uint16('-'),
	57: uint16(71),
	58: uint16('.'),
	59: uint16(71),
	60: uint16(':'),
	61: uint16(71),
	62: uint16(0xb7),
	63: uint16(71),
}

var map_token1 = [38]uint16_t{
	0:  uint16('"'),
	1:  uint16(46),
	2:  uint16('#'),
	3:  uint16(28),
	4:  uint16('%'),
	5:  uint16(91),
	6:  uint16('\''),
	7:  uint16(45),
	8:  uint16('('),
	9:  uint16(74),
	10: uint16(')'),
	11: uint16(77),
	12: uint16('*'),
	13: uint16(78),
	14: uint16('+'),
	15: uint16(80),
	16: uint16(','),
	17: uint16(81),
	18: uint16('/'),
	19: uint16(7),
	20: uint16('>'),
	21: uint16(52),
	22: uint16('?'),
	23: uint16(79),
	24: uint16('['),
	25: uint16(49),
	26: uint16(']'),
	27: uint16(36),
	28: uint16('|'),
	29: uint16(76),
	30: uint16('\t'),
	31: uint16(95),
	32: uint16('\n'),
	33: uint16(95),
	34: uint16('\r'),
	35: uint16(95),
	36: uint16(' '),
	37: uint16(95),
}

var map_token2 = [52]uint16_t{
	0:  uint16('"'),
	1:  uint16(46),
	2:  uint16('#'),
	3:  uint16(22),
	4:  uint16('%'),
	5:  uint16(91),
	6:  uint16('\''),
	7:  uint16(45),
	8:  uint16('('),
	9:  uint16(74),
	10: uint16(')'),
	11: uint16(77),
	12: uint16('*'),
	13: uint16(78),
	14: uint16('+'),
	15: uint16(80),
	16: uint16(','),
	17: uint16(81),
	18: uint16('/'),
	19: uint16(7),
	20: uint16('1'),
	21: uint16(6),
	22: uint16(';'),
	23: uint16(94),
	24: uint16('<'),
	25: uint16(54),
	26: uint16('='),
	27: uint16(143),
	28: uint16('>'),
	29: uint16(52),
	30: uint16('?'),
	31: uint16(79),
	32: uint16('E'),
	33: uint16(113),
	34: uint16('I'),
	35: uint16(96),
	36: uint16('N'),
	37: uint16(111),
	38: uint16('['),
	39: uint16(49),
	40: uint16(']'),
	41: uint16(50),
	42: uint16('|'),
	43: uint16(76),
	44: uint16('\t'),
	45: uint16(95),
	46: uint16('\n'),
	47: uint16(95),
	48: uint16('\r'),
	49: uint16(95),
	50: uint16(' '),
	51: uint16(95),
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
			if !(uint64(i) < libc.Uint64FromInt64(52)/libc.Uint64FromInt64(2)) {
				break
			}
			if libc.Int32FromUint16(map_token3[i]) == lookahead {
				state = map_token3[i+uint32(1)]
				goto next_state
			}
			goto _1
		_1:
			;
			i = i + uint32(2)
		}
		return result
	case int32(1):
		if lookahead == int32('N') {
			state = uint16(14)
			goto next_state
		}
		if lookahead == int32('T') {
			state = uint16(15)
			goto next_state
		}
		return result
	case int32(2):
		if lookahead == int32('D') {
			state = uint16(16)
			goto next_state
		}
		return result
	case int32(3):
		if lookahead == int32('O') {
			state = uint16(17)
			goto next_state
		}
		return result
	case int32(4):
		if lookahead == int32('L') {
			state = uint16(18)
			goto next_state
		}
		if lookahead == int32('M') {
			state = uint16(19)
			goto next_state
		}
		if lookahead == int32('N') {
			state = uint16(20)
			goto next_state
		}
		return result
	case int32(5):
		if lookahead == int32('D') {
			state = uint16(21)
			goto next_state
		}
		if lookahead == int32('O') {
			state = uint16(22)
			goto next_state
		}
		return result
	case int32(6):
		if lookahead == int32('U') {
			state = uint16(23)
			goto next_state
		}
		return result
	case int32(7):
		if lookahead == int32('Y') {
			state = uint16(24)
			goto next_state
		}
		return result
	case int32(8):
		if lookahead == int32('n') {
			state = uint16(25)
			goto next_state
		}
		return result
	case int32(9):
		if lookahead == int32('o') {
			state = uint16(26)
			goto next_state
		}
		return result
	case int32(10):
		if lookahead == int32('t') {
			state = uint16(27)
			goto next_state
		}
		return result
	case int32(11):
		if lookahead == int32('e') {
			state = uint16(28)
			goto next_state
		}
		return result
	case int32(12):
		if lookahead == int32('m') {
			state = uint16(29)
			goto next_state
		}
		return result
	case int32(13):
		if lookahead == int32('e') {
			state = uint16(30)
			goto next_state
		}
		return result
	case int32(14):
		if lookahead == int32('Y') {
			state = uint16(31)
			goto next_state
		}
		return result
	case int32(15):
		if lookahead == int32('T') {
			state = uint16(32)
			goto next_state
		}
		return result
	case int32(16):
		if lookahead == int32('A') {
			state = uint16(33)
			goto next_state
		}
		return result
	case int32(17):
		if lookahead == int32('C') {
			state = uint16(34)
			goto next_state
		}
		return result
	case int32(18):
		if lookahead == int32('E') {
			state = uint16(35)
			goto next_state
		}
		return result
	case int32(19):
		if lookahead == int32('P') {
			state = uint16(36)
			goto next_state
		}
		return result
	case int32(20):
		if lookahead == int32('T') {
			state = uint16(37)
			goto next_state
		}
		return result
	case int32(21):
		if lookahead == int32('A') {
			state = uint16(38)
			goto next_state
		}
		return result
	case int32(22):
		if lookahead == int32('T') {
			state = uint16(39)
			goto next_state
		}
		return result
	case int32(23):
		if lookahead == int32('B') {
			state = uint16(40)
			goto next_state
		}
		return result
	case int32(24):
		if lookahead == int32('S') {
			state = uint16(41)
			goto next_state
		}
		return result
	case int32(25):
		if lookahead == int32('c') {
			state = uint16(42)
			goto next_state
		}
		return result
	case int32(26):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_no)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(27):
		if lookahead == int32('a') {
			state = uint16(43)
			goto next_state
		}
		return result
	case int32(28):
		if lookahead == int32('r') {
			state = uint16(44)
			goto next_state
		}
		return result
	case int32(29):
		if lookahead == int32('l') {
			state = uint16(45)
			goto next_state
		}
		return result
	case int32(30):
		if lookahead == int32('s') {
			state = uint16(46)
			goto next_state
		}
		return result
	case int32(31):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ANY)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(32):
		if lookahead == int32('L') {
			state = uint16(47)
			goto next_state
		}
		return result
	case int32(33):
		if lookahead == int32('T') {
			state = uint16(48)
			goto next_state
		}
		return result
	case int32(34):
		if lookahead == int32('T') {
			state = uint16(49)
			goto next_state
		}
		return result
	case int32(35):
		if lookahead == int32('M') {
			state = uint16(50)
			goto next_state
		}
		return result
	case int32(36):
		if lookahead == int32('T') {
			state = uint16(51)
			goto next_state
		}
		return result
	case int32(37):
		if lookahead == int32('I') {
			state = uint16(52)
			goto next_state
		}
		return result
	case int32(38):
		if lookahead == int32('T') {
			state = uint16(53)
			goto next_state
		}
		return result
	case int32(39):
		if lookahead == int32('A') {
			state = uint16(54)
			goto next_state
		}
		return result
	case int32(40):
		if lookahead == int32('L') {
			state = uint16(55)
			goto next_state
		}
		return result
	case int32(41):
		if lookahead == int32('T') {
			state = uint16(56)
			goto next_state
		}
		return result
	case int32(42):
		if lookahead == int32('o') {
			state = uint16(57)
			goto next_state
		}
		return result
	case int32(43):
		if lookahead == int32('n') {
			state = uint16(58)
			goto next_state
		}
		return result
	case int32(44):
		if lookahead == int32('s') {
			state = uint16(59)
			goto next_state
		}
		return result
	case int32(45):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_xml)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		if lookahead == int32('-') {
			state = uint16(60)
			goto next_state
		}
		return result
	case int32(46):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_yes)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(47):
		if lookahead == int32('I') {
			state = uint16(61)
			goto next_state
		}
		return result
	case int32(48):
		if lookahead == int32('A') {
			state = uint16(62)
			goto next_state
		}
		return result
	case int32(49):
		if lookahead == int32('Y') {
			state = uint16(63)
			goto next_state
		}
		return result
	case int32(50):
		if lookahead == int32('E') {
			state = uint16(64)
			goto next_state
		}
		return result
	case int32(51):
		if lookahead == int32('Y') {
			state = uint16(65)
			goto next_state
		}
		return result
	case int32(52):
		if lookahead == int32('T') {
			state = uint16(66)
			goto next_state
		}
		return result
	case int32(53):
		if lookahead == int32('A') {
			state = uint16(67)
			goto next_state
		}
		return result
	case int32(54):
		if lookahead == int32('T') {
			state = uint16(68)
			goto next_state
		}
		return result
	case int32(55):
		if lookahead == int32('I') {
			state = uint16(69)
			goto next_state
		}
		return result
	case int32(56):
		if lookahead == int32('E') {
			state = uint16(70)
			goto next_state
		}
		return result
	case int32(57):
		if lookahead == int32('d') {
			state = uint16(71)
			goto next_state
		}
		return result
	case int32(58):
		if lookahead == int32('d') {
			state = uint16(72)
			goto next_state
		}
		return result
	case int32(59):
		if lookahead == int32('i') {
			state = uint16(73)
			goto next_state
		}
		return result
	case int32(60):
		if lookahead == int32('m') {
			state = uint16(74)
			goto next_state
		}
		if lookahead == int32('s') {
			state = uint16(75)
			goto next_state
		}
		return result
	case int32(61):
		if lookahead == int32('S') {
			state = uint16(76)
			goto next_state
		}
		return result
	case int32(62):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_CDATA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(63):
		if lookahead == int32('P') {
			state = uint16(77)
			goto next_state
		}
		return result
	case int32(64):
		if lookahead == int32('N') {
			state = uint16(78)
			goto next_state
		}
		return result
	case int32(65):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_EMPTY)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(66):
		if lookahead == int32('Y') {
			state = uint16(79)
			goto next_state
		}
		return result
	case int32(67):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NDATA)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(68):
		if lookahead == int32('I') {
			state = uint16(80)
			goto next_state
		}
		return result
	case int32(69):
		if lookahead == int32('C') {
			state = uint16(81)
			goto next_state
		}
		return result
	case int32(70):
		if lookahead == int32('M') {
			state = uint16(82)
			goto next_state
		}
		return result
	case int32(71):
		if lookahead == int32('i') {
			state = uint16(83)
			goto next_state
		}
		return result
	case int32(72):
		if lookahead == int32('a') {
			state = uint16(84)
			goto next_state
		}
		return result
	case int32(73):
		if lookahead == int32('o') {
			state = uint16(85)
			goto next_state
		}
		return result
	case int32(74):
		if lookahead == int32('o') {
			state = uint16(86)
			goto next_state
		}
		return result
	case int32(75):
		if lookahead == int32('t') {
			state = uint16(87)
			goto next_state
		}
		return result
	case int32(76):
		if lookahead == int32('T') {
			state = uint16(88)
			goto next_state
		}
		return result
	case int32(77):
		if lookahead == int32('E') {
			state = uint16(89)
			goto next_state
		}
		return result
	case int32(78):
		if lookahead == int32('T') {
			state = uint16(90)
			goto next_state
		}
		return result
	case int32(79):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ENTITY)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(80):
		if lookahead == int32('O') {
			state = uint16(91)
			goto next_state
		}
		return result
	case int32(81):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_PUBLIC)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(82):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_SYSTEM)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(83):
		if lookahead == int32('n') {
			state = uint16(92)
			goto next_state
		}
		return result
	case int32(84):
		if lookahead == int32('l') {
			state = uint16(93)
			goto next_state
		}
		return result
	case int32(85):
		if lookahead == int32('n') {
			state = uint16(94)
			goto next_state
		}
		return result
	case int32(86):
		if lookahead == int32('d') {
			state = uint16(95)
			goto next_state
		}
		return result
	case int32(87):
		if lookahead == int32('y') {
			state = uint16(96)
			goto next_state
		}
		return result
	case int32(88):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ATTLIST)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(89):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_DOCTYPE)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(90):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_ELEMENT)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(91):
		if lookahead == int32('N') {
			state = uint16(97)
			goto next_state
		}
		return result
	case int32(92):
		if lookahead == int32('g') {
			state = uint16(98)
			goto next_state
		}
		return result
	case int32(93):
		if lookahead == int32('o') {
			state = uint16(99)
			goto next_state
		}
		return result
	case int32(94):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_version)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(95):
		if lookahead == int32('e') {
			state = uint16(100)
			goto next_state
		}
		return result
	case int32(96):
		if lookahead == int32('l') {
			state = uint16(101)
			goto next_state
		}
		return result
	case int32(97):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_NOTATION)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(98):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_encoding)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(99):
		if lookahead == int32('n') {
			state = uint16(102)
			goto next_state
		}
		return result
	case int32(100):
		if lookahead == int32('l') {
			state = uint16(103)
			goto next_state
		}
		return result
	case int32(101):
		if lookahead == int32('e') {
			state = uint16(104)
			goto next_state
		}
		return result
	case int32(102):
		if lookahead == int32('e') {
			state = uint16(105)
			goto next_state
		}
		return result
	case int32(103):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_xml_DASHmodel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(104):
		if lookahead == int32('s') {
			state = uint16(106)
			goto next_state
		}
		return result
	case int32(105):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_standalone)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	case int32(106):
		if lookahead == int32('h') {
			state = uint16(107)
			goto next_state
		}
		return result
	case int32(107):
		if lookahead == int32('e') {
			state = uint16(108)
			goto next_state
		}
		return result
	case int32(108):
		if lookahead == int32('e') {
			state = uint16(109)
			goto next_state
		}
		return result
	case int32(109):
		if lookahead == int32('t') {
			state = uint16(110)
			goto next_state
		}
		return result
	case int32(110):
		result = libc.BoolUint8(true1 != 0)
		(*TSLexer)(unsafe.Pointer(lexer)).Fresult_symbol = uint16(anon_sym_xml_DASHstylesheet)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*TSLexer)(unsafe.Pointer(lexer)).Fmark_end})))(tls, lexer)
		return result
	default:
		return libc.BoolUint8(false1 != 0)
	}
	return r
}

var map_token3 = [26]uint16_t{
	0:  uint16('A'),
	1:  uint16(1),
	2:  uint16('C'),
	3:  uint16(2),
	4:  uint16('D'),
	5:  uint16(3),
	6:  uint16('E'),
	7:  uint16(4),
	8:  uint16('N'),
	9:  uint16(5),
	10: uint16('P'),
	11: uint16(6),
	12: uint16('S'),
	13: uint16(7),
	14: uint16('e'),
	15: uint16(8),
	16: uint16('n'),
	17: uint16(9),
	18: uint16('s'),
	19: uint16(10),
	20: uint16('v'),
	21: uint16(11),
	22: uint16('x'),
	23: uint16(12),
	24: uint16('y'),
	25: uint16(13),
}

var ts_lex_modes = [495]TSLexMode{
	0: {
		Fexternal_lex_state: uint16(1),
	},
	1: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	2: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	3: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	4: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	5: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	6: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	7: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	8: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	9: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	10: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	11: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	12: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	13: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	14: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	15: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	16: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	17: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	18: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	19: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	20: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	21: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	22: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	23: {
		Flex_state: uint16(4),
	},
	24: {
		Flex_state: uint16(2),
	},
	25: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	26: {
		Flex_state: uint16(41),
	},
	27: {
		Flex_state: uint16(41),
	},
	28: {
		Flex_state: uint16(4),
	},
	29: {
		Flex_state: uint16(2),
	},
	30: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	31: {
		Flex_state: uint16(41),
	},
	32: {
		Flex_state: uint16(4),
	},
	33: {
		Flex_state: uint16(2),
	},
	34: {
		Flex_state: uint16(1),
	},
	35: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	36: {
		Flex_state: uint16(5),
	},
	37: {
		Flex_state: uint16(3),
	},
	38: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	39: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	40: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	41: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	42: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	43: {
		Flex_state: uint16(5),
	},
	44: {
		Flex_state: uint16(3),
	},
	45: {
		Flex_state: uint16(5),
	},
	46: {
		Flex_state: uint16(3),
	},
	47: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	48: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	49: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	50: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	51: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	52: {
		Flex_state: uint16(1),
	},
	53: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	54: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	55: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	56: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	57: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	58: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	59: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	60: {
		Flex_state: uint16(3),
	},
	61: {
		Flex_state: uint16(3),
	},
	62: {
		Flex_state: uint16(41),
	},
	63: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	64: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	65: {
		Flex_state: uint16(41),
	},
	66: {
		Flex_state: uint16(41),
	},
	67: {
		Flex_state: uint16(41),
	},
	68: {
		Flex_state: uint16(41),
	},
	69: {
		Flex_state: uint16(41),
	},
	70: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	71: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	72: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	73: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	74: {
		Flex_state: uint16(3),
	},
	75: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	76: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	77: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	78: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	79: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	80: {
		Flex_state:          uint16(4),
		Fexternal_lex_state: uint16(3),
	},
	81: {
		Flex_state: uint16(5),
	},
	82: {
		Flex_state: uint16(5),
	},
	83: {
		Flex_state: uint16(3),
	},
	84: {
		Flex_state: uint16(5),
	},
	85: {
		Flex_state: uint16(5),
	},
	86: {
		Flex_state: uint16(41),
	},
	87: {
		Flex_state: uint16(41),
	},
	88: {
		Flex_state: uint16(41),
	},
	89: {
		Flex_state: uint16(41),
	},
	90: {
		Flex_state: uint16(41),
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
		Flex_state: uint16(41),
	},
	95: {
		Flex_state: uint16(1),
	},
	96: {
		Flex_state: uint16(41),
	},
	97: {
		Flex_state: uint16(41),
	},
	98: {
		Flex_state: uint16(1),
	},
	99: {
		Flex_state: uint16(41),
	},
	100: {
		Flex_state: uint16(4),
	},
	101: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	102: {
		Flex_state: uint16(1),
	},
	103: {
		Flex_state: uint16(1),
	},
	104: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	105: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	106: {
		Flex_state: uint16(1),
	},
	107: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	108: {
		Flex_state: uint16(1),
	},
	109: {
		Flex_state: uint16(1),
	},
	110: {
		Flex_state: uint16(41),
	},
	111: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	112: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	113: {
		Flex_state: uint16(1),
	},
	114: {
		Flex_state: uint16(1),
	},
	115: {
		Flex_state: uint16(1),
	},
	116: {
		Flex_state: uint16(1),
	},
	117: {
		Flex_state: uint16(1),
	},
	118: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	119: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	120: {
		Flex_state: uint16(4),
	},
	121: {
		Flex_state: uint16(2),
	},
	122: {
		Flex_state: uint16(4),
	},
	123: {
		Flex_state: uint16(4),
	},
	124: {
		Flex_state: uint16(2),
	},
	125: {
		Flex_state: uint16(2),
	},
	126: {
		Flex_state: uint16(2),
	},
	127: {
		Flex_state: uint16(41),
	},
	128: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	129: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	130: {
		Flex_state: uint16(41),
	},
	131: {
		Flex_state: uint16(41),
	},
	132: {
		Flex_state: uint16(41),
	},
	133: {
		Flex_state: uint16(41),
	},
	134: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	135: {
		Flex_state: uint16(41),
	},
	136: {
		Flex_state: uint16(41),
	},
	137: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	138: {
		Flex_state: uint16(41),
	},
	139: {
		Flex_state: uint16(41),
	},
	140: {
		Flex_state: uint16(41),
	},
	141: {
		Flex_state: uint16(1),
	},
	142: {
		Flex_state: uint16(5),
	},
	143: {
		Flex_state: uint16(41),
	},
	144: {
		Flex_state: uint16(5),
	},
	145: {
		Flex_state: uint16(5),
	},
	146: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	147: {
		Flex_state: uint16(1),
	},
	148: {
		Flex_state: uint16(3),
	},
	149: {
		Flex_state: uint16(3),
	},
	150: {
		Flex_state: uint16(3),
	},
	151: {
		Flex_state: uint16(41),
	},
	152: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	153: {
		Flex_state: uint16(41),
	},
	154: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	155: {
		Flex_state: uint16(41),
	},
	156: {
		Flex_state: uint16(41),
	},
	157: {
		Flex_state: uint16(41),
	},
	158: {
		Flex_state: uint16(41),
	},
	159: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	160: {
		Flex_state: uint16(1),
	},
	161: {
		Flex_state: uint16(41),
	},
	162: {
		Flex_state: uint16(1),
	},
	163: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	164: {
		Flex_state: uint16(1),
	},
	165: {
		Flex_state: uint16(41),
	},
	166: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	167: {
		Flex_state: uint16(41),
	},
	168: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	169: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	170: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	171: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	172: {
		Flex_state: uint16(1),
	},
	173: {},
	174: {
		Flex_state: uint16(41),
	},
	175: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	176: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	177: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	178: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	179: {
		Flex_state: uint16(41),
	},
	180: {
		Flex_state: uint16(1),
	},
	181: {
		Flex_state: uint16(41),
	},
	182: {
		Flex_state: uint16(41),
	},
	183: {
		Flex_state: uint16(41),
	},
	184: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	185: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	186: {},
	187: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	188: {
		Flex_state: uint16(41),
	},
	189: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	190: {
		Flex_state: uint16(41),
	},
	191: {
		Flex_state: uint16(41),
	},
	192: {
		Flex_state: uint16(41),
	},
	193: {
		Flex_state: uint16(41),
	},
	194: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	195: {
		Flex_state: uint16(41),
	},
	196: {
		Flex_state: uint16(41),
	},
	197: {
		Flex_state: uint16(1),
	},
	198: {
		Flex_state: uint16(41),
	},
	199: {
		Flex_state: uint16(41),
	},
	200: {
		Flex_state: uint16(41),
	},
	201: {
		Flex_state: uint16(41),
	},
	202: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	203: {
		Flex_state: uint16(1),
	},
	204: {
		Flex_state: uint16(1),
	},
	205: {
		Flex_state: uint16(1),
	},
	206: {
		Flex_state: uint16(1),
	},
	207: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	208: {
		Flex_state: uint16(10),
	},
	209: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	210: {
		Flex_state: uint16(41),
	},
	211: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	212: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	213: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	214: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	215: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(5),
	},
	216: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	217: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(2),
	},
	218: {
		Flex_state: uint16(41),
	},
	219: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	220: {
		Flex_state: uint16(10),
	},
	221: {
		Flex_state: uint16(41),
	},
	222: {
		Flex_state: uint16(1),
	},
	223: {
		Flex_state: uint16(41),
	},
	224: {
		Flex_state: uint16(10),
	},
	225: {
		Flex_state: uint16(10),
	},
	226: {
		Flex_state: uint16(41),
	},
	227: {
		Flex_state: uint16(10),
	},
	228: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(4),
	},
	229: {
		Flex_state: uint16(10),
	},
	230: {
		Flex_state: uint16(41),
	},
	231: {
		Flex_state: uint16(41),
	},
	232: {
		Flex_state: uint16(41),
	},
	233: {
		Flex_state: uint16(41),
	},
	234: {
		Flex_state: uint16(1),
	},
	235: {
		Flex_state: uint16(1),
	},
	236: {
		Flex_state: uint16(41),
	},
	237: {
		Flex_state: uint16(41),
	},
	238: {
		Flex_state: uint16(41),
	},
	239: {
		Flex_state: uint16(41),
	},
	240: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	241: {
		Flex_state: uint16(41),
	},
	242: {
		Flex_state: uint16(1),
	},
	243: {
		Flex_state: uint16(41),
	},
	244: {
		Flex_state: uint16(41),
	},
	245: {
		Flex_state: uint16(41),
	},
	246: {
		Flex_state: uint16(41),
	},
	247: {
		Flex_state: uint16(10),
	},
	248: {
		Flex_state: uint16(10),
	},
	249: {
		Flex_state: uint16(41),
	},
	250: {
		Flex_state: uint16(1),
	},
	251: {
		Flex_state: uint16(10),
	},
	252: {
		Flex_state: uint16(41),
	},
	253: {
		Flex_state: uint16(41),
	},
	254: {
		Flex_state: uint16(41),
	},
	255: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	256: {
		Flex_state: uint16(41),
	},
	257: {
		Flex_state: uint16(41),
	},
	258: {
		Flex_state: uint16(1),
	},
	259: {
		Flex_state: uint16(1),
	},
	260: {
		Flex_state: uint16(41),
	},
	261: {
		Flex_state: uint16(41),
	},
	262: {
		Flex_state: uint16(1),
	},
	263: {
		Flex_state: uint16(41),
	},
	264: {
		Flex_state: uint16(1),
	},
	265: {
		Flex_state: uint16(10),
	},
	266: {
		Flex_state: uint16(41),
	},
	267: {
		Flex_state: uint16(41),
	},
	268: {
		Flex_state: uint16(10),
	},
	269: {
		Flex_state: uint16(10),
	},
	270: {
		Flex_state:          uint16(41),
		Fexternal_lex_state: uint16(5),
	},
	271: {
		Flex_state: uint16(41),
	},
	272: {
		Flex_state: uint16(41),
	},
	273: {
		Flex_state: uint16(41),
	},
	274: {
		Flex_state: uint16(41),
	},
	275: {
		Flex_state: uint16(41),
	},
	276: {
		Flex_state: uint16(10),
	},
	277: {
		Flex_state: uint16(41),
	},
	278: {
		Flex_state: uint16(41),
	},
	279: {
		Flex_state: uint16(1),
	},
	280: {
		Flex_state: uint16(41),
	},
	281: {
		Flex_state: uint16(41),
	},
	282: {
		Flex_state: uint16(41),
	},
	283: {
		Flex_state: uint16(41),
	},
	284: {
		Flex_state: uint16(41),
	},
	285: {
		Flex_state: uint16(37),
	},
	286: {
		Flex_state: uint16(41),
	},
	287: {
		Flex_state: uint16(41),
	},
	288: {
		Flex_state: uint16(41),
	},
	289: {
		Flex_state: uint16(41),
	},
	290: {
		Flex_state: uint16(41),
	},
	291: {
		Flex_state: uint16(41),
	},
	292: {
		Flex_state: uint16(41),
	},
	293: {
		Flex_state: uint16(41),
	},
	294: {
		Flex_state: uint16(1),
	},
	295: {
		Flex_state: uint16(41),
	},
	296: {
		Flex_state: uint16(41),
	},
	297: {
		Flex_state: uint16(41),
	},
	298: {
		Flex_state: uint16(41),
	},
	299: {
		Flex_state: uint16(41),
	},
	300: {
		Flex_state: uint16(10),
	},
	301: {
		Flex_state: uint16(41),
	},
	302: {
		Flex_state: uint16(41),
	},
	303: {
		Flex_state: uint16(41),
	},
	304: {
		Flex_state: uint16(41),
	},
	305: {
		Flex_state: uint16(37),
	},
	306: {
		Flex_state: uint16(41),
	},
	307: {
		Flex_state: uint16(41),
	},
	308: {
		Flex_state: uint16(41),
	},
	309: {
		Flex_state: uint16(37),
	},
	310: {
		Flex_state: uint16(41),
	},
	311: {
		Flex_state: uint16(10),
	},
	312: {
		Flex_state: uint16(41),
	},
	313: {
		Flex_state: uint16(41),
	},
	314: {
		Flex_state: uint16(41),
	},
	315: {
		Flex_state: uint16(41),
	},
	316: {
		Flex_state: uint16(41),
	},
	317: {
		Flex_state: uint16(41),
	},
	318: {
		Flex_state: uint16(41),
	},
	319: {
		Flex_state: uint16(41),
	},
	320: {
		Flex_state: uint16(41),
	},
	321: {
		Flex_state: uint16(41),
	},
	322: {
		Flex_state: uint16(41),
	},
	323: {
		Flex_state: uint16(41),
	},
	324: {
		Flex_state: uint16(41),
	},
	325: {
		Flex_state: uint16(41),
	},
	326: {
		Flex_state: uint16(41),
	},
	327: {
		Flex_state: uint16(41),
	},
	328: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(6),
	},
	329: {
		Flex_state: uint16(41),
	},
	330: {
		Flex_state: uint16(41),
	},
	331: {
		Flex_state: uint16(1),
	},
	332: {
		Flex_state: uint16(41),
	},
	333: {
		Flex_state: uint16(41),
	},
	334: {
		Flex_state: uint16(10),
	},
	335: {
		Flex_state: uint16(41),
	},
	336: {
		Flex_state: uint16(41),
	},
	337: {
		Flex_state: uint16(10),
	},
	338: {
		Flex_state: uint16(10),
	},
	339: {
		Flex_state: uint16(41),
	},
	340: {
		Flex_state: uint16(41),
	},
	341: {
		Flex_state: uint16(1),
	},
	342: {
		Flex_state: uint16(41),
	},
	343: {
		Flex_state: uint16(41),
	},
	344: {
		Flex_state: uint16(41),
	},
	345: {
		Flex_state: uint16(41),
	},
	346: {
		Flex_state: uint16(41),
	},
	347: {
		Flex_state: uint16(41),
	},
	348: {
		Flex_state: uint16(41),
	},
	349: {
		Flex_state: uint16(41),
	},
	350: {},
	351: {
		Flex_state: uint16(41),
	},
	352: {
		Flex_state: uint16(41),
	},
	353: {
		Flex_state: uint16(10),
	},
	354: {
		Flex_state: uint16(41),
	},
	355: {
		Flex_state: uint16(41),
	},
	356: {
		Flex_state: uint16(10),
	},
	357: {
		Flex_state: uint16(10),
	},
	358: {
		Flex_state: uint16(41),
	},
	359: {
		Flex_state: uint16(1),
	},
	360: {
		Flex_state:          uint16(1),
		Fexternal_lex_state: uint16(6),
	},
	361: {
		Flex_state: uint16(10),
	},
	362: {
		Flex_state: uint16(41),
	},
	363: {
		Flex_state: uint16(41),
	},
	364: {
		Flex_state: uint16(10),
	},
	365: {
		Flex_state: uint16(41),
	},
	366: {},
	367: {
		Flex_state: uint16(41),
	},
	368: {
		Flex_state: uint16(10),
	},
	369: {
		Flex_state: uint16(41),
	},
	370: {
		Flex_state: uint16(41),
	},
	371: {
		Flex_state: uint16(41),
	},
	372: {
		Flex_state: uint16(41),
	},
	373: {
		Flex_state: uint16(1),
	},
	374: {
		Flex_state: uint16(41),
	},
	375: {
		Flex_state: uint16(10),
	},
	376: {
		Flex_state: uint16(41),
	},
	377: {
		Flex_state: uint16(41),
	},
	378: {
		Flex_state: uint16(39),
	},
	379: {
		Flex_state: uint16(41),
	},
	380: {
		Flex_state: uint16(37),
	},
	381: {
		Flex_state: uint16(41),
	},
	382: {
		Flex_state: uint16(41),
	},
	383: {
		Flex_state: uint16(41),
	},
	384: {
		Flex_state: uint16(41),
	},
	385: {
		Flex_state: uint16(41),
	},
	386: {
		Flex_state: uint16(41),
	},
	387: {
		Flex_state: uint16(41),
	},
	388: {
		Flex_state: uint16(41),
	},
	389: {},
	390: {
		Flex_state: uint16(41),
	},
	391: {
		Flex_state: uint16(41),
	},
	392: {
		Flex_state: uint16(41),
	},
	393: {
		Flex_state: uint16(37),
	},
	394: {
		Flex_state: uint16(41),
	},
	395: {
		Flex_state: uint16(40),
	},
	396: {
		Flex_state: uint16(41),
	},
	397: {
		Flex_state: uint16(41),
	},
	398: {
		Flex_state: uint16(41),
	},
	399: {
		Flex_state: uint16(40),
	},
	400: {
		Flex_state: uint16(138),
	},
	401: {
		Flex_state: uint16(137),
	},
	402: {
		Flex_state: uint16(41),
	},
	403: {
		Flex_state: uint16(140),
	},
	404: {
		Flex_state: uint16(139),
	},
	405: {
		Flex_state: uint16(41),
	},
	406: {
		Flex_state: uint16(41),
	},
	407: {
		Flex_state: uint16(41),
	},
	408: {
		Flex_state: uint16(41),
	},
	409: {
		Flex_state: uint16(41),
	},
	410: {
		Flex_state: uint16(41),
	},
	411: {
		Flex_state: uint16(41),
	},
	412: {
		Flex_state: uint16(41),
	},
	413: {
		Flex_state: uint16(41),
	},
	414: {
		Flex_state: uint16(41),
	},
	415: {
		Flex_state: uint16(37),
	},
	416: {
		Flex_state: uint16(41),
	},
	417: {
		Flex_state: uint16(41),
	},
	418: {
		Flex_state: uint16(41),
	},
	419: {
		Flex_state: uint16(41),
	},
	420: {
		Flex_state: uint16(41),
	},
	421: {
		Flex_state: uint16(41),
	},
	422: {
		Flex_state: uint16(41),
	},
	423: {
		Flex_state: uint16(10),
	},
	424: {
		Flex_state: uint16(41),
	},
	425: {
		Flex_state: uint16(41),
	},
	426: {
		Fexternal_lex_state: uint16(7),
	},
	427: {
		Flex_state: uint16(41),
	},
	428: {
		Flex_state: uint16(41),
	},
	429: {
		Flex_state: uint16(41),
	},
	430: {
		Flex_state: uint16(41),
	},
	431: {
		Flex_state: uint16(41),
	},
	432: {
		Flex_state: uint16(41),
	},
	433: {
		Flex_state: uint16(41),
	},
	434: {
		Flex_state: uint16(10),
	},
	435: {
		Flex_state: uint16(1),
	},
	436: {
		Flex_state: uint16(1),
	},
	437: {
		Flex_state: uint16(1),
	},
	438: {
		Flex_state: uint16(41),
	},
	439: {
		Flex_state: uint16(10),
	},
	440: {
		Flex_state: uint16(1),
	},
	441: {
		Fexternal_lex_state: uint16(8),
	},
	442: {
		Fexternal_lex_state: uint16(9),
	},
	443: {
		Flex_state: uint16(1),
	},
	444: {
		Flex_state: uint16(1),
	},
	445: {
		Flex_state: uint16(41),
	},
	446: {
		Fexternal_lex_state: uint16(10),
	},
	447: {
		Flex_state: uint16(41),
	},
	448: {
		Flex_state: uint16(41),
	},
	449: {
		Flex_state: uint16(41),
	},
	450: {
		Flex_state: uint16(41),
	},
	451: {
		Flex_state: uint16(1),
	},
	452: {
		Flex_state: uint16(41),
	},
	453: {
		Flex_state: uint16(41),
	},
	454: {
		Flex_state: uint16(41),
	},
	455: {
		Flex_state: uint16(41),
	},
	456: {
		Flex_state: uint16(10),
	},
	457: {
		Flex_state: uint16(41),
	},
	458: {
		Flex_state: uint16(41),
	},
	459: {
		Flex_state: uint16(41),
	},
	460: {
		Flex_state: uint16(41),
	},
	461: {
		Flex_state: uint16(41),
	},
	462: {
		Flex_state: uint16(1),
	},
	463: {
		Flex_state: uint16(41),
	},
	464: {
		Flex_state: uint16(41),
	},
	465: {
		Flex_state: uint16(41),
	},
	466: {
		Flex_state: uint16(41),
	},
	467: {
		Flex_state: uint16(10),
	},
	468: {
		Flex_state: uint16(41),
	},
	469: {
		Flex_state: uint16(41),
	},
	470: {
		Flex_state: uint16(41),
	},
	471: {
		Flex_state: uint16(41),
	},
	472: {
		Flex_state: uint16(41),
	},
	473: {
		Flex_state: uint16(41),
	},
	474: {
		Fexternal_lex_state: uint16(7),
	},
	475: {
		Fexternal_lex_state: uint16(8),
	},
	476: {
		Fexternal_lex_state: uint16(9),
	},
	477: {
		Flex_state: uint16(1),
	},
	478: {
		Flex_state: uint16(1),
	},
	479: {
		Flex_state: uint16(39),
	},
	480: {
		Fexternal_lex_state: uint16(10),
	},
	481: {
		Flex_state: uint16(1),
	},
	482: {
		Flex_state: uint16(1),
	},
	483: {
		Flex_state: uint16(1),
	},
	484: {
		Flex_state: uint16(39),
	},
	485: {
		Fexternal_lex_state: uint16(10),
	},
	486: {
		Flex_state: uint16(1),
	},
	487: {
		Flex_state: uint16(1),
	},
	488: {
		Flex_state: uint16(1),
	},
	489: {
		Flex_state: uint16(39),
	},
	490: {
		Flex_state: uint16(1),
	},
	491: {
		Flex_state: uint16(1),
	},
	492: {
		Flex_state: uint16(1),
	},
	493: {
		Flex_state: uint16(39),
	},
	494: {
		Flex_state: uint16(1),
	},
}

var ts_parse_table = [2][143]uint16_t{
	0: {
		0:  uint16(1),
		1:  uint16(1),
		2:  uint16(1),
		3:  uint16(1),
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
		61: uint16(1),
		62: uint16(1),
		63: uint16(1),
		64: uint16(1),
		65: uint16(1),
		66: uint16(1),
		67: uint16(1),
		68: uint16(1),
		69: uint16(1),
		70: uint16(1),
		71: uint16(1),
		72: uint16(1),
		73: uint16(1),
	},
	1: {
		2:   uint16(3),
		10:  uint16(5),
		15:  uint16(7),
		48:  uint16(9),
		68:  uint16(11),
		74:  uint16(389),
		75:  uint16(173),
		76:  uint16(20),
		77:  uint16(22),
		79:  uint16(58),
		81:  uint16(38),
		82:  uint16(189),
		84:  uint16(3),
		89:  uint16(20),
		90:  uint16(20),
		126: uint16(20),
		128: uint16(20),
	},
}

var ts_small_parse_table = [5178]uint16_t{
	0:    uint16(15),
	1:    uint16(13),
	2:    uint16(1),
	3:    uint16(anon_sym_LT_QMARK),
	4:    uint16(15),
	5:    uint16(1),
	6:    uint16(anon_sym_LT),
	7:    uint16(17),
	8:    uint16(1),
	9:    uint16(anon_sym_LT_SLASH),
	10:   uint16(19),
	11:   uint16(1),
	12:   uint16(anon_sym_LT_BANG_LBRACK),
	13:   uint16(21),
	14:   uint16(1),
	15:   uint16(anon_sym_AMP),
	16:   uint16(23),
	17:   uint16(1),
	18:   uint16(anon_sym_AMP_POUND),
	19:   uint16(25),
	20:   uint16(1),
	21:   uint16(anon_sym_AMP_POUNDx),
	22:   uint16(2),
	23:   uint16(1),
	24:   uint16(sym_STag),
	25:   uint16(70),
	26:   uint16(1),
	27:   uint16(sym_EmptyElemTag),
	28:   uint16(71),
	29:   uint16(1),
	30:   uint16(sym_ETag),
	31:   uint16(328),
	32:   uint16(1),
	33:   uint16(sym_CDStart),
	34:   uint16(366),
	35:   uint16(1),
	36:   uint16(sym_content),
	37:   uint16(27),
	38:   uint16(2),
	39:   uint16(sym_Comment),
	40:   uint16(sym_CharData),
	41:   uint16(80),
	42:   uint16(2),
	43:   uint16(sym_EntityRef),
	44:   uint16(sym_CharRef),
	45:   uint16(4),
	46:   uint16(5),
	47:   uint16(sym_element),
	48:   uint16(sym_CDSect),
	49:   uint16(sym__Reference),
	50:   uint16(sym_PI),
	51:   uint16(aux_sym_content_repeat1),
	52:   uint16(15),
	53:   uint16(13),
	54:   uint16(1),
	55:   uint16(anon_sym_LT_QMARK),
	56:   uint16(15),
	57:   uint16(1),
	58:   uint16(anon_sym_LT),
	59:   uint16(19),
	60:   uint16(1),
	61:   uint16(anon_sym_LT_BANG_LBRACK),
	62:   uint16(21),
	63:   uint16(1),
	64:   uint16(anon_sym_AMP),
	65:   uint16(23),
	66:   uint16(1),
	67:   uint16(anon_sym_AMP_POUND),
	68:   uint16(25),
	69:   uint16(1),
	70:   uint16(anon_sym_AMP_POUNDx),
	71:   uint16(29),
	72:   uint16(1),
	73:   uint16(anon_sym_LT_SLASH),
	74:   uint16(2),
	75:   uint16(1),
	76:   uint16(sym_STag),
	77:   uint16(70),
	78:   uint16(1),
	79:   uint16(sym_EmptyElemTag),
	80:   uint16(202),
	81:   uint16(1),
	82:   uint16(sym_ETag),
	83:   uint16(328),
	84:   uint16(1),
	85:   uint16(sym_CDStart),
	86:   uint16(350),
	87:   uint16(1),
	88:   uint16(sym_content),
	89:   uint16(27),
	90:   uint16(2),
	91:   uint16(sym_Comment),
	92:   uint16(sym_CharData),
	93:   uint16(80),
	94:   uint16(2),
	95:   uint16(sym_EntityRef),
	96:   uint16(sym_CharRef),
	97:   uint16(4),
	98:   uint16(5),
	99:   uint16(sym_element),
	100:  uint16(sym_CDSect),
	101:  uint16(sym__Reference),
	102:  uint16(sym_PI),
	103:  uint16(aux_sym_content_repeat1),
	104:  uint16(13),
	105:  uint16(13),
	106:  uint16(1),
	107:  uint16(anon_sym_LT_QMARK),
	108:  uint16(15),
	109:  uint16(1),
	110:  uint16(anon_sym_LT),
	111:  uint16(19),
	112:  uint16(1),
	113:  uint16(anon_sym_LT_BANG_LBRACK),
	114:  uint16(21),
	115:  uint16(1),
	116:  uint16(anon_sym_AMP),
	117:  uint16(23),
	118:  uint16(1),
	119:  uint16(anon_sym_AMP_POUND),
	120:  uint16(25),
	121:  uint16(1),
	122:  uint16(anon_sym_AMP_POUNDx),
	123:  uint16(31),
	124:  uint16(1),
	125:  uint16(anon_sym_LT_SLASH),
	126:  uint16(2),
	127:  uint16(1),
	128:  uint16(sym_STag),
	129:  uint16(70),
	130:  uint16(1),
	131:  uint16(sym_EmptyElemTag),
	132:  uint16(328),
	133:  uint16(1),
	134:  uint16(sym_CDStart),
	135:  uint16(33),
	136:  uint16(2),
	137:  uint16(sym_Comment),
	138:  uint16(sym_CharData),
	139:  uint16(80),
	140:  uint16(2),
	141:  uint16(sym_EntityRef),
	142:  uint16(sym_CharRef),
	143:  uint16(5),
	144:  uint16(5),
	145:  uint16(sym_element),
	146:  uint16(sym_CDSect),
	147:  uint16(sym__Reference),
	148:  uint16(sym_PI),
	149:  uint16(aux_sym_content_repeat1),
	150:  uint16(13),
	151:  uint16(35),
	152:  uint16(1),
	153:  uint16(anon_sym_LT_QMARK),
	154:  uint16(38),
	155:  uint16(1),
	156:  uint16(anon_sym_LT),
	157:  uint16(41),
	158:  uint16(1),
	159:  uint16(anon_sym_LT_SLASH),
	160:  uint16(43),
	161:  uint16(1),
	162:  uint16(anon_sym_LT_BANG_LBRACK),
	163:  uint16(46),
	164:  uint16(1),
	165:  uint16(anon_sym_AMP),
	166:  uint16(49),
	167:  uint16(1),
	168:  uint16(anon_sym_AMP_POUND),
	169:  uint16(52),
	170:  uint16(1),
	171:  uint16(anon_sym_AMP_POUNDx),
	172:  uint16(2),
	173:  uint16(1),
	174:  uint16(sym_STag),
	175:  uint16(70),
	176:  uint16(1),
	177:  uint16(sym_EmptyElemTag),
	178:  uint16(328),
	179:  uint16(1),
	180:  uint16(sym_CDStart),
	181:  uint16(55),
	182:  uint16(2),
	183:  uint16(sym_Comment),
	184:  uint16(sym_CharData),
	185:  uint16(80),
	186:  uint16(2),
	187:  uint16(sym_EntityRef),
	188:  uint16(sym_CharRef),
	189:  uint16(5),
	190:  uint16(5),
	191:  uint16(sym_element),
	192:  uint16(sym_CDSect),
	193:  uint16(sym__Reference),
	194:  uint16(sym_PI),
	195:  uint16(aux_sym_content_repeat1),
	196:  uint16(11),
	197:  uint16(3),
	198:  uint16(1),
	199:  uint16(anon_sym_LT_QMARK),
	200:  uint16(5),
	201:  uint16(1),
	202:  uint16(anon_sym_LT_BANG),
	203:  uint16(7),
	204:  uint16(1),
	205:  uint16(anon_sym_LT),
	206:  uint16(3),
	207:  uint16(1),
	208:  uint16(sym_STag),
	209:  uint16(22),
	210:  uint16(1),
	211:  uint16(sym_XMLDecl),
	212:  uint16(39),
	213:  uint16(1),
	214:  uint16(sym_element),
	215:  uint16(58),
	216:  uint16(1),
	217:  uint16(sym_doctypedecl),
	218:  uint16(186),
	219:  uint16(1),
	220:  uint16(sym_prolog),
	221:  uint16(189),
	222:  uint16(1),
	223:  uint16(sym_EmptyElemTag),
	224:  uint16(11),
	225:  uint16(2),
	226:  uint16(sym_Comment),
	227:  uint16(sym__S),
	228:  uint16(20),
	229:  uint16(5),
	230:  uint16(sym__Misc),
	231:  uint16(sym_StyleSheetPI),
	232:  uint16(sym_XmlModelPI),
	233:  uint16(sym_PI),
	234:  uint16(aux_sym_document_repeat1),
	235:  uint16(9),
	236:  uint16(58),
	237:  uint16(1),
	238:  uint16(anon_sym_LT_QMARK),
	239:  uint16(60),
	240:  uint16(1),
	241:  uint16(anon_sym_LT_BANG),
	242:  uint16(62),
	243:  uint16(1),
	244:  uint16(anon_sym_RBRACK),
	245:  uint16(64),
	246:  uint16(1),
	247:  uint16(sym__S),
	248:  uint16(66),
	249:  uint16(1),
	250:  uint16(sym_Comment),
	251:  uint16(9),
	252:  uint16(1),
	253:  uint16(aux_sym__intSubset),
	254:  uint16(167),
	255:  uint16(1),
	256:  uint16(sym__markupdecl),
	257:  uint16(316),
	258:  uint16(2),
	259:  uint16(sym_GEDecl),
	260:  uint16(sym_PEDecl),
	261:  uint16(314),
	262:  uint16(5),
	263:  uint16(sym_elementdecl),
	264:  uint16(sym_AttlistDecl),
	265:  uint16(sym__EntityDecl),
	266:  uint16(sym_NotationDecl),
	267:  uint16(sym_PI),
	268:  uint16(9),
	269:  uint16(58),
	270:  uint16(1),
	271:  uint16(anon_sym_LT_QMARK),
	272:  uint16(60),
	273:  uint16(1),
	274:  uint16(anon_sym_LT_BANG),
	275:  uint16(66),
	276:  uint16(1),
	277:  uint16(sym_Comment),
	278:  uint16(68),
	279:  uint16(1),
	280:  uint16(anon_sym_RBRACK),
	281:  uint16(70),
	282:  uint16(1),
	283:  uint16(sym__S),
	284:  uint16(14),
	285:  uint16(1),
	286:  uint16(aux_sym__intSubset),
	287:  uint16(167),
	288:  uint16(1),
	289:  uint16(sym__markupdecl),
	290:  uint16(316),
	291:  uint16(2),
	292:  uint16(sym_GEDecl),
	293:  uint16(sym_PEDecl),
	294:  uint16(314),
	295:  uint16(5),
	296:  uint16(sym_elementdecl),
	297:  uint16(sym_AttlistDecl),
	298:  uint16(sym__EntityDecl),
	299:  uint16(sym_NotationDecl),
	300:  uint16(sym_PI),
	301:  uint16(9),
	302:  uint16(58),
	303:  uint16(1),
	304:  uint16(anon_sym_LT_QMARK),
	305:  uint16(60),
	306:  uint16(1),
	307:  uint16(anon_sym_LT_BANG),
	308:  uint16(66),
	309:  uint16(1),
	310:  uint16(sym_Comment),
	311:  uint16(68),
	312:  uint16(1),
	313:  uint16(anon_sym_RBRACK),
	314:  uint16(72),
	315:  uint16(1),
	316:  uint16(sym__S),
	317:  uint16(13),
	318:  uint16(1),
	319:  uint16(aux_sym__intSubset),
	320:  uint16(167),
	321:  uint16(1),
	322:  uint16(sym__markupdecl),
	323:  uint16(316),
	324:  uint16(2),
	325:  uint16(sym_GEDecl),
	326:  uint16(sym_PEDecl),
	327:  uint16(314),
	328:  uint16(5),
	329:  uint16(sym_elementdecl),
	330:  uint16(sym_AttlistDecl),
	331:  uint16(sym__EntityDecl),
	332:  uint16(sym_NotationDecl),
	333:  uint16(sym_PI),
	334:  uint16(9),
	335:  uint16(58),
	336:  uint16(1),
	337:  uint16(anon_sym_LT_QMARK),
	338:  uint16(60),
	339:  uint16(1),
	340:  uint16(anon_sym_LT_BANG),
	341:  uint16(66),
	342:  uint16(1),
	343:  uint16(sym_Comment),
	344:  uint16(72),
	345:  uint16(1),
	346:  uint16(sym__S),
	347:  uint16(74),
	348:  uint16(1),
	349:  uint16(anon_sym_RBRACK),
	350:  uint16(13),
	351:  uint16(1),
	352:  uint16(aux_sym__intSubset),
	353:  uint16(167),
	354:  uint16(1),
	355:  uint16(sym__markupdecl),
	356:  uint16(316),
	357:  uint16(2),
	358:  uint16(sym_GEDecl),
	359:  uint16(sym_PEDecl),
	360:  uint16(314),
	361:  uint16(5),
	362:  uint16(sym_elementdecl),
	363:  uint16(sym_AttlistDecl),
	364:  uint16(sym__EntityDecl),
	365:  uint16(sym_NotationDecl),
	366:  uint16(sym_PI),
	367:  uint16(9),
	368:  uint16(58),
	369:  uint16(1),
	370:  uint16(anon_sym_LT_QMARK),
	371:  uint16(60),
	372:  uint16(1),
	373:  uint16(anon_sym_LT_BANG),
	374:  uint16(66),
	375:  uint16(1),
	376:  uint16(sym_Comment),
	377:  uint16(74),
	378:  uint16(1),
	379:  uint16(anon_sym_RBRACK),
	380:  uint16(76),
	381:  uint16(1),
	382:  uint16(sym__S),
	383:  uint16(12),
	384:  uint16(1),
	385:  uint16(aux_sym__intSubset),
	386:  uint16(167),
	387:  uint16(1),
	388:  uint16(sym__markupdecl),
	389:  uint16(316),
	390:  uint16(2),
	391:  uint16(sym_GEDecl),
	392:  uint16(sym_PEDecl),
	393:  uint16(314),
	394:  uint16(5),
	395:  uint16(sym_elementdecl),
	396:  uint16(sym_AttlistDecl),
	397:  uint16(sym__EntityDecl),
	398:  uint16(sym_NotationDecl),
	399:  uint16(sym_PI),
	400:  uint16(9),
	401:  uint16(58),
	402:  uint16(1),
	403:  uint16(anon_sym_LT_QMARK),
	404:  uint16(60),
	405:  uint16(1),
	406:  uint16(anon_sym_LT_BANG),
	407:  uint16(62),
	408:  uint16(1),
	409:  uint16(anon_sym_RBRACK),
	410:  uint16(66),
	411:  uint16(1),
	412:  uint16(sym_Comment),
	413:  uint16(72),
	414:  uint16(1),
	415:  uint16(sym__S),
	416:  uint16(13),
	417:  uint16(1),
	418:  uint16(aux_sym__intSubset),
	419:  uint16(167),
	420:  uint16(1),
	421:  uint16(sym__markupdecl),
	422:  uint16(316),
	423:  uint16(2),
	424:  uint16(sym_GEDecl),
	425:  uint16(sym_PEDecl),
	426:  uint16(314),
	427:  uint16(5),
	428:  uint16(sym_elementdecl),
	429:  uint16(sym_AttlistDecl),
	430:  uint16(sym__EntityDecl),
	431:  uint16(sym_NotationDecl),
	432:  uint16(sym_PI),
	433:  uint16(9),
	434:  uint16(78),
	435:  uint16(1),
	436:  uint16(anon_sym_LT_QMARK),
	437:  uint16(81),
	438:  uint16(1),
	439:  uint16(anon_sym_LT_BANG),
	440:  uint16(84),
	441:  uint16(1),
	442:  uint16(anon_sym_RBRACK),
	443:  uint16(86),
	444:  uint16(1),
	445:  uint16(sym__S),
	446:  uint16(89),
	447:  uint16(1),
	448:  uint16(sym_Comment),
	449:  uint16(13),
	450:  uint16(1),
	451:  uint16(aux_sym__intSubset),
	452:  uint16(167),
	453:  uint16(1),
	454:  uint16(sym__markupdecl),
	455:  uint16(316),
	456:  uint16(2),
	457:  uint16(sym_GEDecl),
	458:  uint16(sym_PEDecl),
	459:  uint16(314),
	460:  uint16(5),
	461:  uint16(sym_elementdecl),
	462:  uint16(sym_AttlistDecl),
	463:  uint16(sym__EntityDecl),
	464:  uint16(sym_NotationDecl),
	465:  uint16(sym_PI),
	466:  uint16(9),
	467:  uint16(58),
	468:  uint16(1),
	469:  uint16(anon_sym_LT_QMARK),
	470:  uint16(60),
	471:  uint16(1),
	472:  uint16(anon_sym_LT_BANG),
	473:  uint16(66),
	474:  uint16(1),
	475:  uint16(sym_Comment),
	476:  uint16(72),
	477:  uint16(1),
	478:  uint16(sym__S),
	479:  uint16(92),
	480:  uint16(1),
	481:  uint16(anon_sym_RBRACK),
	482:  uint16(13),
	483:  uint16(1),
	484:  uint16(aux_sym__intSubset),
	485:  uint16(167),
	486:  uint16(1),
	487:  uint16(sym__markupdecl),
	488:  uint16(316),
	489:  uint16(2),
	490:  uint16(sym_GEDecl),
	491:  uint16(sym_PEDecl),
	492:  uint16(314),
	493:  uint16(5),
	494:  uint16(sym_elementdecl),
	495:  uint16(sym_AttlistDecl),
	496:  uint16(sym__EntityDecl),
	497:  uint16(sym_NotationDecl),
	498:  uint16(sym_PI),
	499:  uint16(9),
	500:  uint16(58),
	501:  uint16(1),
	502:  uint16(anon_sym_LT_QMARK),
	503:  uint16(60),
	504:  uint16(1),
	505:  uint16(anon_sym_LT_BANG),
	506:  uint16(66),
	507:  uint16(1),
	508:  uint16(sym_Comment),
	509:  uint16(94),
	510:  uint16(1),
	511:  uint16(anon_sym_RBRACK),
	512:  uint16(96),
	513:  uint16(1),
	514:  uint16(sym__S),
	515:  uint16(10),
	516:  uint16(1),
	517:  uint16(aux_sym__intSubset),
	518:  uint16(167),
	519:  uint16(1),
	520:  uint16(sym__markupdecl),
	521:  uint16(316),
	522:  uint16(2),
	523:  uint16(sym_GEDecl),
	524:  uint16(sym_PEDecl),
	525:  uint16(314),
	526:  uint16(5),
	527:  uint16(sym_elementdecl),
	528:  uint16(sym_AttlistDecl),
	529:  uint16(sym__EntityDecl),
	530:  uint16(sym_NotationDecl),
	531:  uint16(sym_PI),
	532:  uint16(7),
	533:  uint16(58),
	534:  uint16(1),
	535:  uint16(anon_sym_LT_QMARK),
	536:  uint16(60),
	537:  uint16(1),
	538:  uint16(anon_sym_LT_BANG),
	539:  uint16(62),
	540:  uint16(1),
	541:  uint16(anon_sym_RBRACK),
	542:  uint16(66),
	543:  uint16(1),
	544:  uint16(sym_Comment),
	545:  uint16(188),
	546:  uint16(1),
	547:  uint16(sym__markupdecl),
	548:  uint16(316),
	549:  uint16(2),
	550:  uint16(sym_GEDecl),
	551:  uint16(sym_PEDecl),
	552:  uint16(314),
	553:  uint16(5),
	554:  uint16(sym_elementdecl),
	555:  uint16(sym_AttlistDecl),
	556:  uint16(sym__EntityDecl),
	557:  uint16(sym_NotationDecl),
	558:  uint16(sym_PI),
	559:  uint16(7),
	560:  uint16(58),
	561:  uint16(1),
	562:  uint16(anon_sym_LT_QMARK),
	563:  uint16(60),
	564:  uint16(1),
	565:  uint16(anon_sym_LT_BANG),
	566:  uint16(66),
	567:  uint16(1),
	568:  uint16(sym_Comment),
	569:  uint16(92),
	570:  uint16(1),
	571:  uint16(anon_sym_RBRACK),
	572:  uint16(188),
	573:  uint16(1),
	574:  uint16(sym__markupdecl),
	575:  uint16(316),
	576:  uint16(2),
	577:  uint16(sym_GEDecl),
	578:  uint16(sym_PEDecl),
	579:  uint16(314),
	580:  uint16(5),
	581:  uint16(sym_elementdecl),
	582:  uint16(sym_AttlistDecl),
	583:  uint16(sym__EntityDecl),
	584:  uint16(sym_NotationDecl),
	585:  uint16(sym_PI),
	586:  uint16(7),
	587:  uint16(58),
	588:  uint16(1),
	589:  uint16(anon_sym_LT_QMARK),
	590:  uint16(60),
	591:  uint16(1),
	592:  uint16(anon_sym_LT_BANG),
	593:  uint16(66),
	594:  uint16(1),
	595:  uint16(sym_Comment),
	596:  uint16(74),
	597:  uint16(1),
	598:  uint16(anon_sym_RBRACK),
	599:  uint16(188),
	600:  uint16(1),
	601:  uint16(sym__markupdecl),
	602:  uint16(316),
	603:  uint16(2),
	604:  uint16(sym_GEDecl),
	605:  uint16(sym_PEDecl),
	606:  uint16(314),
	607:  uint16(5),
	608:  uint16(sym_elementdecl),
	609:  uint16(sym_AttlistDecl),
	610:  uint16(sym__EntityDecl),
	611:  uint16(sym_NotationDecl),
	612:  uint16(sym_PI),
	613:  uint16(7),
	614:  uint16(58),
	615:  uint16(1),
	616:  uint16(anon_sym_LT_QMARK),
	617:  uint16(60),
	618:  uint16(1),
	619:  uint16(anon_sym_LT_BANG),
	620:  uint16(66),
	621:  uint16(1),
	622:  uint16(sym_Comment),
	623:  uint16(68),
	624:  uint16(1),
	625:  uint16(anon_sym_RBRACK),
	626:  uint16(188),
	627:  uint16(1),
	628:  uint16(sym__markupdecl),
	629:  uint16(316),
	630:  uint16(2),
	631:  uint16(sym_GEDecl),
	632:  uint16(sym_PEDecl),
	633:  uint16(314),
	634:  uint16(5),
	635:  uint16(sym_elementdecl),
	636:  uint16(sym_AttlistDecl),
	637:  uint16(sym__EntityDecl),
	638:  uint16(sym_NotationDecl),
	639:  uint16(sym_PI),
	640:  uint16(6),
	641:  uint16(5),
	642:  uint16(1),
	643:  uint16(anon_sym_LT_BANG),
	644:  uint16(98),
	645:  uint16(1),
	646:  uint16(anon_sym_LT_QMARK),
	647:  uint16(100),
	648:  uint16(1),
	649:  uint16(anon_sym_LT),
	650:  uint16(50),
	651:  uint16(1),
	652:  uint16(sym_doctypedecl),
	653:  uint16(102),
	654:  uint16(2),
	655:  uint16(sym_Comment),
	656:  uint16(sym__S),
	657:  uint16(25),
	658:  uint16(5),
	659:  uint16(sym__Misc),
	660:  uint16(sym_StyleSheetPI),
	661:  uint16(sym_XmlModelPI),
	662:  uint16(sym_PI),
	663:  uint16(aux_sym_document_repeat1),
	664:  uint16(6),
	665:  uint16(58),
	666:  uint16(1),
	667:  uint16(anon_sym_LT_QMARK),
	668:  uint16(60),
	669:  uint16(1),
	670:  uint16(anon_sym_LT_BANG),
	671:  uint16(66),
	672:  uint16(1),
	673:  uint16(sym_Comment),
	674:  uint16(188),
	675:  uint16(1),
	676:  uint16(sym__markupdecl),
	677:  uint16(316),
	678:  uint16(2),
	679:  uint16(sym_GEDecl),
	680:  uint16(sym_PEDecl),
	681:  uint16(314),
	682:  uint16(5),
	683:  uint16(sym_elementdecl),
	684:  uint16(sym_AttlistDecl),
	685:  uint16(sym__EntityDecl),
	686:  uint16(sym_NotationDecl),
	687:  uint16(sym_PI),
	688:  uint16(6),
	689:  uint16(5),
	690:  uint16(1),
	691:  uint16(anon_sym_LT_BANG),
	692:  uint16(98),
	693:  uint16(1),
	694:  uint16(anon_sym_LT_QMARK),
	695:  uint16(100),
	696:  uint16(1),
	697:  uint16(anon_sym_LT),
	698:  uint16(50),
	699:  uint16(1),
	700:  uint16(sym_doctypedecl),
	701:  uint16(104),
	702:  uint16(2),
	703:  uint16(sym_Comment),
	704:  uint16(sym__S),
	705:  uint16(30),
	706:  uint16(5),
	707:  uint16(sym__Misc),
	708:  uint16(sym_StyleSheetPI),
	709:  uint16(sym_XmlModelPI),
	710:  uint16(sym_PI),
	711:  uint16(aux_sym_document_repeat1),
	712:  uint16(8),
	713:  uint16(106),
	714:  uint16(1),
	715:  uint16(anon_sym_SQUOTE),
	716:  uint16(108),
	717:  uint16(1),
	718:  uint16(anon_sym_PERCENT),
	719:  uint16(110),
	720:  uint16(1),
	721:  uint16(aux_sym_EntityValue_token2),
	722:  uint16(112),
	723:  uint16(1),
	724:  uint16(anon_sym_AMP),
	725:  uint16(114),
	726:  uint16(1),
	727:  uint16(anon_sym_AMP_POUND),
	728:  uint16(116),
	729:  uint16(1),
	730:  uint16(anon_sym_AMP_POUNDx),
	731:  uint16(100),
	732:  uint16(2),
	733:  uint16(sym_EntityRef),
	734:  uint16(sym_CharRef),
	735:  uint16(28),
	736:  uint16(3),
	737:  uint16(sym_PEReference),
	738:  uint16(sym__Reference),
	739:  uint16(aux_sym_EntityValue_repeat2),
	740:  uint16(8),
	741:  uint16(106),
	742:  uint16(1),
	743:  uint16(anon_sym_DQUOTE),
	744:  uint16(118),
	745:  uint16(1),
	746:  uint16(anon_sym_PERCENT),
	747:  uint16(120),
	748:  uint16(1),
	749:  uint16(aux_sym_EntityValue_token1),
	750:  uint16(122),
	751:  uint16(1),
	752:  uint16(anon_sym_AMP),
	753:  uint16(124),
	754:  uint16(1),
	755:  uint16(anon_sym_AMP_POUND),
	756:  uint16(126),
	757:  uint16(1),
	758:  uint16(anon_sym_AMP_POUNDx),
	759:  uint16(125),
	760:  uint16(2),
	761:  uint16(sym_EntityRef),
	762:  uint16(sym_CharRef),
	763:  uint16(29),
	764:  uint16(3),
	765:  uint16(sym_PEReference),
	766:  uint16(sym__Reference),
	767:  uint16(aux_sym_EntityValue_repeat1),
	768:  uint16(5),
	769:  uint16(130),
	770:  uint16(1),
	771:  uint16(anon_sym_LT_QMARK),
	772:  uint16(133),
	773:  uint16(1),
	774:  uint16(anon_sym_LT),
	775:  uint16(128),
	776:  uint16(2),
	778:  uint16(anon_sym_LT_BANG),
	779:  uint16(135),
	780:  uint16(2),
	781:  uint16(sym_Comment),
	782:  uint16(sym__S),
	783:  uint16(25),
	784:  uint16(5),
	785:  uint16(sym__Misc),
	786:  uint16(sym_StyleSheetPI),
	787:  uint16(sym_XmlModelPI),
	788:  uint16(sym_PI),
	789:  uint16(aux_sym_document_repeat1),
	790:  uint16(5),
	791:  uint16(96),
	792:  uint16(1),
	793:  uint16(aux_sym_Mixed_repeat1),
	794:  uint16(174),
	795:  uint16(1),
	796:  uint16(aux_sym_Mixed_repeat2),
	797:  uint16(246),
	798:  uint16(1),
	799:  uint16(sym_PEReference),
	800:  uint16(140),
	801:  uint16(3),
	802:  uint16(anon_sym_STAR),
	803:  uint16(anon_sym_QMARK),
	804:  uint16(anon_sym_PLUS),
	805:  uint16(138),
	806:  uint16(5),
	807:  uint16(anon_sym_PIPE),
	808:  uint16(anon_sym_RPAREN),
	809:  uint16(anon_sym_COMMA),
	810:  uint16(anon_sym_PERCENT),
	811:  uint16(sym__S),
	812:  uint16(8),
	813:  uint16(142),
	814:  uint16(1),
	815:  uint16(anon_sym_CDATA),
	816:  uint16(144),
	817:  uint16(1),
	818:  uint16(anon_sym_LPAREN),
	819:  uint16(146),
	820:  uint16(1),
	821:  uint16(sym_TokenizedType),
	822:  uint16(148),
	823:  uint16(1),
	824:  uint16(anon_sym_NOTATION),
	825:  uint16(150),
	826:  uint16(1),
	827:  uint16(anon_sym_PERCENT),
	828:  uint16(322),
	829:  uint16(1),
	830:  uint16(sym__AttType),
	831:  uint16(376),
	832:  uint16(2),
	833:  uint16(sym_NotationType),
	834:  uint16(sym_Enumeration),
	835:  uint16(318),
	836:  uint16(3),
	837:  uint16(sym_StringType),
	838:  uint16(sym__EnumeratedType),
	839:  uint16(sym_PEReference),
	840:  uint16(8),
	841:  uint16(108),
	842:  uint16(1),
	843:  uint16(anon_sym_PERCENT),
	844:  uint16(112),
	845:  uint16(1),
	846:  uint16(anon_sym_AMP),
	847:  uint16(114),
	848:  uint16(1),
	849:  uint16(anon_sym_AMP_POUND),
	850:  uint16(116),
	851:  uint16(1),
	852:  uint16(anon_sym_AMP_POUNDx),
	853:  uint16(152),
	854:  uint16(1),
	855:  uint16(anon_sym_SQUOTE),
	856:  uint16(154),
	857:  uint16(1),
	858:  uint16(aux_sym_EntityValue_token2),
	859:  uint16(100),
	860:  uint16(2),
	861:  uint16(sym_EntityRef),
	862:  uint16(sym_CharRef),
	863:  uint16(32),
	864:  uint16(3),
	865:  uint16(sym_PEReference),
	866:  uint16(sym__Reference),
	867:  uint16(aux_sym_EntityValue_repeat2),
	868:  uint16(8),
	869:  uint16(118),
	870:  uint16(1),
	871:  uint16(anon_sym_PERCENT),
	872:  uint16(122),
	873:  uint16(1),
	874:  uint16(anon_sym_AMP),
	875:  uint16(124),
	876:  uint16(1),
	877:  uint16(anon_sym_AMP_POUND),
	878:  uint16(126),
	879:  uint16(1),
	880:  uint16(anon_sym_AMP_POUNDx),
	881:  uint16(152),
	882:  uint16(1),
	883:  uint16(anon_sym_DQUOTE),
	884:  uint16(156),
	885:  uint16(1),
	886:  uint16(aux_sym_EntityValue_token1),
	887:  uint16(125),
	888:  uint16(2),
	889:  uint16(sym_EntityRef),
	890:  uint16(sym_CharRef),
	891:  uint16(33),
	892:  uint16(3),
	893:  uint16(sym_PEReference),
	894:  uint16(sym__Reference),
	895:  uint16(aux_sym_EntityValue_repeat1),
	896:  uint16(6),
	897:  uint16(5),
	898:  uint16(1),
	899:  uint16(anon_sym_LT_BANG),
	900:  uint16(98),
	901:  uint16(1),
	902:  uint16(anon_sym_LT_QMARK),
	903:  uint16(158),
	904:  uint16(1),
	905:  uint16(anon_sym_LT),
	906:  uint16(48),
	907:  uint16(1),
	908:  uint16(sym_doctypedecl),
	909:  uint16(102),
	910:  uint16(2),
	911:  uint16(sym_Comment),
	912:  uint16(sym__S),
	913:  uint16(25),
	914:  uint16(5),
	915:  uint16(sym__Misc),
	916:  uint16(sym_StyleSheetPI),
	917:  uint16(sym_XmlModelPI),
	918:  uint16(sym_PI),
	919:  uint16(aux_sym_document_repeat1),
	920:  uint16(5),
	921:  uint16(99),
	922:  uint16(1),
	923:  uint16(aux_sym_Mixed_repeat1),
	924:  uint16(181),
	925:  uint16(1),
	926:  uint16(aux_sym_Mixed_repeat2),
	927:  uint16(246),
	928:  uint16(1),
	929:  uint16(sym_PEReference),
	930:  uint16(140),
	931:  uint16(3),
	932:  uint16(anon_sym_STAR),
	933:  uint16(anon_sym_QMARK),
	934:  uint16(anon_sym_PLUS),
	935:  uint16(138),
	936:  uint16(5),
	937:  uint16(anon_sym_PIPE),
	938:  uint16(anon_sym_RPAREN),
	939:  uint16(anon_sym_COMMA),
	940:  uint16(anon_sym_PERCENT),
	941:  uint16(sym__S),
	942:  uint16(8),
	943:  uint16(160),
	944:  uint16(1),
	945:  uint16(anon_sym_SQUOTE),
	946:  uint16(162),
	947:  uint16(1),
	948:  uint16(anon_sym_PERCENT),
	949:  uint16(165),
	950:  uint16(1),
	951:  uint16(aux_sym_EntityValue_token2),
	952:  uint16(168),
	953:  uint16(1),
	954:  uint16(anon_sym_AMP),
	955:  uint16(171),
	956:  uint16(1),
	957:  uint16(anon_sym_AMP_POUND),
	958:  uint16(174),
	959:  uint16(1),
	960:  uint16(anon_sym_AMP_POUNDx),
	961:  uint16(100),
	962:  uint16(2),
	963:  uint16(sym_EntityRef),
	964:  uint16(sym_CharRef),
	965:  uint16(32),
	966:  uint16(3),
	967:  uint16(sym_PEReference),
	968:  uint16(sym__Reference),
	969:  uint16(aux_sym_EntityValue_repeat2),
	970:  uint16(8),
	971:  uint16(177),
	972:  uint16(1),
	973:  uint16(anon_sym_DQUOTE),
	974:  uint16(179),
	975:  uint16(1),
	976:  uint16(anon_sym_PERCENT),
	977:  uint16(182),
	978:  uint16(1),
	979:  uint16(aux_sym_EntityValue_token1),
	980:  uint16(185),
	981:  uint16(1),
	982:  uint16(anon_sym_AMP),
	983:  uint16(188),
	984:  uint16(1),
	985:  uint16(anon_sym_AMP_POUND),
	986:  uint16(191),
	987:  uint16(1),
	988:  uint16(anon_sym_AMP_POUNDx),
	989:  uint16(125),
	990:  uint16(2),
	991:  uint16(sym_EntityRef),
	992:  uint16(sym_CharRef),
	993:  uint16(33),
	994:  uint16(3),
	995:  uint16(sym_PEReference),
	996:  uint16(sym__Reference),
	997:  uint16(aux_sym_EntityValue_repeat1),
	998:  uint16(1),
	999:  uint16(194),
	1000: uint16(10),
	1001: uint16(anon_sym_GT),
	1002: uint16(anon_sym_PIPE),
	1003: uint16(anon_sym_RPAREN),
	1004: uint16(anon_sym_STAR),
	1005: uint16(anon_sym_QMARK),
	1006: uint16(anon_sym_PLUS),
	1007: uint16(anon_sym_COMMA),
	1008: uint16(anon_sym_PERCENT),
	1009: uint16(sym__S),
	1010: uint16(sym_Name),
	1011: uint16(2),
	1012: uint16(198),
	1013: uint16(3),
	1014: uint16(anon_sym_LT),
	1015: uint16(anon_sym_AMP),
	1016: uint16(anon_sym_AMP_POUND),
	1017: uint16(196),
	1018: uint16(6),
	1019: uint16(sym_Comment),
	1020: uint16(sym_CharData),
	1021: uint16(anon_sym_LT_QMARK),
	1022: uint16(anon_sym_LT_SLASH),
	1023: uint16(anon_sym_LT_BANG_LBRACK),
	1024: uint16(anon_sym_AMP_POUNDx),
	1025: uint16(7),
	1026: uint16(200),
	1027: uint16(1),
	1028: uint16(anon_sym_SQUOTE),
	1029: uint16(202),
	1030: uint16(1),
	1031: uint16(aux_sym_PseudoAttValue_token2),
	1032: uint16(204),
	1033: uint16(1),
	1034: uint16(anon_sym_AMP),
	1035: uint16(206),
	1036: uint16(1),
	1037: uint16(anon_sym_AMP_POUND),
	1038: uint16(208),
	1039: uint16(1),
	1040: uint16(anon_sym_AMP_POUNDx),
	1041: uint16(45),
	1042: uint16(2),
	1043: uint16(sym__Reference),
	1044: uint16(aux_sym_PseudoAttValue_repeat2),
	1045: uint16(142),
	1046: uint16(2),
	1047: uint16(sym_EntityRef),
	1048: uint16(sym_CharRef),
	1049: uint16(7),
	1050: uint16(200),
	1051: uint16(1),
	1052: uint16(anon_sym_DQUOTE),
	1053: uint16(210),
	1054: uint16(1),
	1055: uint16(aux_sym_PseudoAttValue_token1),
	1056: uint16(212),
	1057: uint16(1),
	1058: uint16(anon_sym_AMP),
	1059: uint16(214),
	1060: uint16(1),
	1061: uint16(anon_sym_AMP_POUND),
	1062: uint16(216),
	1063: uint16(1),
	1064: uint16(anon_sym_AMP_POUNDx),
	1065: uint16(46),
	1066: uint16(2),
	1067: uint16(sym__Reference),
	1068: uint16(aux_sym_PseudoAttValue_repeat1),
	1069: uint16(148),
	1070: uint16(2),
	1071: uint16(sym_EntityRef),
	1072: uint16(sym_CharRef),
	1073: uint16(4),
	1074: uint16(98),
	1075: uint16(1),
	1076: uint16(anon_sym_LT_QMARK),
	1077: uint16(218),
	1078: uint16(1),
	1080: uint16(220),
	1081: uint16(2),
	1082: uint16(sym_Comment),
	1083: uint16(sym__S),
	1084: uint16(64),
	1085: uint16(5),
	1086: uint16(sym__Misc),
	1087: uint16(sym_StyleSheetPI),
	1088: uint16(sym_XmlModelPI),
	1089: uint16(sym_PI),
	1090: uint16(aux_sym_document_repeat1),
	1091: uint16(4),
	1092: uint16(98),
	1093: uint16(1),
	1094: uint16(anon_sym_LT_QMARK),
	1095: uint16(222),
	1096: uint16(1),
	1098: uint16(224),
	1099: uint16(2),
	1100: uint16(sym_Comment),
	1101: uint16(sym__S),
	1102: uint16(42),
	1103: uint16(5),
	1104: uint16(sym__Misc),
	1105: uint16(sym_StyleSheetPI),
	1106: uint16(sym_XmlModelPI),
	1107: uint16(sym_PI),
	1108: uint16(aux_sym_document_repeat1),
	1109: uint16(2),
	1110: uint16(228),
	1111: uint16(3),
	1112: uint16(anon_sym_LT),
	1113: uint16(anon_sym_AMP),
	1114: uint16(anon_sym_AMP_POUND),
	1115: uint16(226),
	1116: uint16(6),
	1117: uint16(sym_Comment),
	1118: uint16(sym_CharData),
	1119: uint16(anon_sym_LT_QMARK),
	1120: uint16(anon_sym_LT_SLASH),
	1121: uint16(anon_sym_LT_BANG_LBRACK),
	1122: uint16(anon_sym_AMP_POUNDx),
	1123: uint16(4),
	1124: uint16(98),
	1125: uint16(1),
	1126: uint16(anon_sym_LT_QMARK),
	1127: uint16(230),
	1128: uint16(1),
	1130: uint16(232),
	1131: uint16(2),
	1132: uint16(sym_Comment),
	1133: uint16(sym__S),
	1134: uint16(53),
	1135: uint16(5),
	1136: uint16(sym__Misc),
	1137: uint16(sym_StyleSheetPI),
	1138: uint16(sym_XmlModelPI),
	1139: uint16(sym_PI),
	1140: uint16(aux_sym_document_repeat1),
	1141: uint16(4),
	1142: uint16(98),
	1143: uint16(1),
	1144: uint16(anon_sym_LT_QMARK),
	1145: uint16(234),
	1146: uint16(1),
	1148: uint16(102),
	1149: uint16(2),
	1150: uint16(sym_Comment),
	1151: uint16(sym__S),
	1152: uint16(25),
	1153: uint16(5),
	1154: uint16(sym__Misc),
	1155: uint16(sym_StyleSheetPI),
	1156: uint16(sym_XmlModelPI),
	1157: uint16(sym_PI),
	1158: uint16(aux_sym_document_repeat1),
	1159: uint16(7),
	1160: uint16(202),
	1161: uint16(1),
	1162: uint16(aux_sym_PseudoAttValue_token2),
	1163: uint16(204),
	1164: uint16(1),
	1165: uint16(anon_sym_AMP),
	1166: uint16(206),
	1167: uint16(1),
	1168: uint16(anon_sym_AMP_POUND),
	1169: uint16(208),
	1170: uint16(1),
	1171: uint16(anon_sym_AMP_POUNDx),
	1172: uint16(236),
	1173: uint16(1),
	1174: uint16(anon_sym_SQUOTE),
	1175: uint16(45),
	1176: uint16(2),
	1177: uint16(sym__Reference),
	1178: uint16(aux_sym_PseudoAttValue_repeat2),
	1179: uint16(142),
	1180: uint16(2),
	1181: uint16(sym_EntityRef),
	1182: uint16(sym_CharRef),
	1183: uint16(7),
	1184: uint16(210),
	1185: uint16(1),
	1186: uint16(aux_sym_PseudoAttValue_token1),
	1187: uint16(212),
	1188: uint16(1),
	1189: uint16(anon_sym_AMP),
	1190: uint16(214),
	1191: uint16(1),
	1192: uint16(anon_sym_AMP_POUND),
	1193: uint16(216),
	1194: uint16(1),
	1195: uint16(anon_sym_AMP_POUNDx),
	1196: uint16(236),
	1197: uint16(1),
	1198: uint16(anon_sym_DQUOTE),
	1199: uint16(46),
	1200: uint16(2),
	1201: uint16(sym__Reference),
	1202: uint16(aux_sym_PseudoAttValue_repeat1),
	1203: uint16(148),
	1204: uint16(2),
	1205: uint16(sym_EntityRef),
	1206: uint16(sym_CharRef),
	1207: uint16(7),
	1208: uint16(238),
	1209: uint16(1),
	1210: uint16(anon_sym_SQUOTE),
	1211: uint16(240),
	1212: uint16(1),
	1213: uint16(aux_sym_PseudoAttValue_token2),
	1214: uint16(243),
	1215: uint16(1),
	1216: uint16(anon_sym_AMP),
	1217: uint16(246),
	1218: uint16(1),
	1219: uint16(anon_sym_AMP_POUND),
	1220: uint16(249),
	1221: uint16(1),
	1222: uint16(anon_sym_AMP_POUNDx),
	1223: uint16(45),
	1224: uint16(2),
	1225: uint16(sym__Reference),
	1226: uint16(aux_sym_PseudoAttValue_repeat2),
	1227: uint16(142),
	1228: uint16(2),
	1229: uint16(sym_EntityRef),
	1230: uint16(sym_CharRef),
	1231: uint16(7),
	1232: uint16(252),
	1233: uint16(1),
	1234: uint16(anon_sym_DQUOTE),
	1235: uint16(254),
	1236: uint16(1),
	1237: uint16(aux_sym_PseudoAttValue_token1),
	1238: uint16(257),
	1239: uint16(1),
	1240: uint16(anon_sym_AMP),
	1241: uint16(260),
	1242: uint16(1),
	1243: uint16(anon_sym_AMP_POUND),
	1244: uint16(263),
	1245: uint16(1),
	1246: uint16(anon_sym_AMP_POUNDx),
	1247: uint16(46),
	1248: uint16(2),
	1249: uint16(sym__Reference),
	1250: uint16(aux_sym_PseudoAttValue_repeat1),
	1251: uint16(148),
	1252: uint16(2),
	1253: uint16(sym_EntityRef),
	1254: uint16(sym_CharRef),
	1255: uint16(4),
	1256: uint16(98),
	1257: uint16(1),
	1258: uint16(anon_sym_LT_QMARK),
	1259: uint16(266),
	1260: uint16(1),
	1261: uint16(anon_sym_LT),
	1262: uint16(102),
	1263: uint16(2),
	1264: uint16(sym_Comment),
	1265: uint16(sym__S),
	1266: uint16(25),
	1267: uint16(5),
	1268: uint16(sym__Misc),
	1269: uint16(sym_StyleSheetPI),
	1270: uint16(sym_XmlModelPI),
	1271: uint16(sym_PI),
	1272: uint16(aux_sym_document_repeat1),
	1273: uint16(4),
	1274: uint16(98),
	1275: uint16(1),
	1276: uint16(anon_sym_LT_QMARK),
	1277: uint16(266),
	1278: uint16(1),
	1279: uint16(anon_sym_LT),
	1280: uint16(268),
	1281: uint16(2),
	1282: uint16(sym_Comment),
	1283: uint16(sym__S),
	1284: uint16(54),
	1285: uint16(5),
	1286: uint16(sym__Misc),
	1287: uint16(sym_StyleSheetPI),
	1288: uint16(sym_XmlModelPI),
	1289: uint16(sym_PI),
	1290: uint16(aux_sym_document_repeat1),
	1291: uint16(2),
	1292: uint16(272),
	1293: uint16(3),
	1294: uint16(anon_sym_LT),
	1295: uint16(anon_sym_AMP),
	1296: uint16(anon_sym_AMP_POUND),
	1297: uint16(270),
	1298: uint16(6),
	1299: uint16(sym_Comment),
	1300: uint16(sym_CharData),
	1301: uint16(anon_sym_LT_QMARK),
	1302: uint16(anon_sym_LT_SLASH),
	1303: uint16(anon_sym_LT_BANG_LBRACK),
	1304: uint16(anon_sym_AMP_POUNDx),
	1305: uint16(4),
	1306: uint16(98),
	1307: uint16(1),
	1308: uint16(anon_sym_LT_QMARK),
	1309: uint16(158),
	1310: uint16(1),
	1311: uint16(anon_sym_LT),
	1312: uint16(274),
	1313: uint16(2),
	1314: uint16(sym_Comment),
	1315: uint16(sym__S),
	1316: uint16(47),
	1317: uint16(5),
	1318: uint16(sym__Misc),
	1319: uint16(sym_StyleSheetPI),
	1320: uint16(sym_XmlModelPI),
	1321: uint16(sym_PI),
	1322: uint16(aux_sym_document_repeat1),
	1323: uint16(2),
	1324: uint16(278),
	1325: uint16(3),
	1326: uint16(anon_sym_LT),
	1327: uint16(anon_sym_AMP),
	1328: uint16(anon_sym_AMP_POUND),
	1329: uint16(276),
	1330: uint16(6),
	1331: uint16(sym_Comment),
	1332: uint16(sym_CharData),
	1333: uint16(anon_sym_LT_QMARK),
	1334: uint16(anon_sym_LT_SLASH),
	1335: uint16(anon_sym_LT_BANG_LBRACK),
	1336: uint16(anon_sym_AMP_POUNDx),
	1337: uint16(6),
	1338: uint16(150),
	1339: uint16(1),
	1340: uint16(anon_sym_PERCENT),
	1341: uint16(282),
	1342: uint16(1),
	1343: uint16(anon_sym_LPAREN),
	1344: uint16(157),
	1345: uint16(1),
	1346: uint16(sym__choice),
	1347: uint16(370),
	1348: uint16(1),
	1349: uint16(sym_contentspec),
	1350: uint16(280),
	1351: uint16(2),
	1352: uint16(anon_sym_EMPTY),
	1353: uint16(anon_sym_ANY),
	1354: uint16(358),
	1355: uint16(3),
	1356: uint16(sym_Mixed),
	1357: uint16(sym_children),
	1358: uint16(sym_PEReference),
	1359: uint16(4),
	1360: uint16(98),
	1361: uint16(1),
	1362: uint16(anon_sym_LT_QMARK),
	1363: uint16(284),
	1364: uint16(1),
	1366: uint16(102),
	1367: uint16(2),
	1368: uint16(sym_Comment),
	1369: uint16(sym__S),
	1370: uint16(25),
	1371: uint16(5),
	1372: uint16(sym__Misc),
	1373: uint16(sym_StyleSheetPI),
	1374: uint16(sym_XmlModelPI),
	1375: uint16(sym_PI),
	1376: uint16(aux_sym_document_repeat1),
	1377: uint16(4),
	1378: uint16(98),
	1379: uint16(1),
	1380: uint16(anon_sym_LT_QMARK),
	1381: uint16(286),
	1382: uint16(1),
	1383: uint16(anon_sym_LT),
	1384: uint16(102),
	1385: uint16(2),
	1386: uint16(sym_Comment),
	1387: uint16(sym__S),
	1388: uint16(25),
	1389: uint16(5),
	1390: uint16(sym__Misc),
	1391: uint16(sym_StyleSheetPI),
	1392: uint16(sym_XmlModelPI),
	1393: uint16(sym_PI),
	1394: uint16(aux_sym_document_repeat1),
	1395: uint16(2),
	1396: uint16(290),
	1397: uint16(3),
	1398: uint16(anon_sym_LT),
	1399: uint16(anon_sym_AMP),
	1400: uint16(anon_sym_AMP_POUND),
	1401: uint16(288),
	1402: uint16(6),
	1403: uint16(sym_Comment),
	1404: uint16(sym_CharData),
	1405: uint16(anon_sym_LT_QMARK),
	1406: uint16(anon_sym_LT_SLASH),
	1407: uint16(anon_sym_LT_BANG_LBRACK),
	1408: uint16(anon_sym_AMP_POUNDx),
	1409: uint16(2),
	1410: uint16(294),
	1411: uint16(3),
	1412: uint16(anon_sym_LT),
	1413: uint16(anon_sym_AMP),
	1414: uint16(anon_sym_AMP_POUND),
	1415: uint16(292),
	1416: uint16(6),
	1417: uint16(sym_Comment),
	1418: uint16(sym_CharData),
	1419: uint16(anon_sym_LT_QMARK),
	1420: uint16(anon_sym_LT_SLASH),
	1421: uint16(anon_sym_LT_BANG_LBRACK),
	1422: uint16(anon_sym_AMP_POUNDx),
	1423: uint16(2),
	1424: uint16(298),
	1425: uint16(3),
	1426: uint16(anon_sym_LT),
	1427: uint16(anon_sym_AMP),
	1428: uint16(anon_sym_AMP_POUND),
	1429: uint16(296),
	1430: uint16(6),
	1431: uint16(sym_Comment),
	1432: uint16(sym_CharData),
	1433: uint16(anon_sym_LT_QMARK),
	1434: uint16(anon_sym_LT_SLASH),
	1435: uint16(anon_sym_LT_BANG_LBRACK),
	1436: uint16(anon_sym_AMP_POUNDx),
	1437: uint16(4),
	1438: uint16(98),
	1439: uint16(1),
	1440: uint16(anon_sym_LT_QMARK),
	1441: uint16(100),
	1442: uint16(1),
	1443: uint16(anon_sym_LT),
	1444: uint16(300),
	1445: uint16(2),
	1446: uint16(sym_Comment),
	1447: uint16(sym__S),
	1448: uint16(63),
	1449: uint16(5),
	1450: uint16(sym__Misc),
	1451: uint16(sym_StyleSheetPI),
	1452: uint16(sym_XmlModelPI),
	1453: uint16(sym_PI),
	1454: uint16(aux_sym_document_repeat1),
	1455: uint16(2),
	1456: uint16(304),
	1457: uint16(3),
	1458: uint16(anon_sym_LT),
	1459: uint16(anon_sym_AMP),
	1460: uint16(anon_sym_AMP_POUND),
	1461: uint16(302),
	1462: uint16(6),
	1463: uint16(sym_Comment),
	1464: uint16(sym_CharData),
	1465: uint16(anon_sym_LT_QMARK),
	1466: uint16(anon_sym_LT_SLASH),
	1467: uint16(anon_sym_LT_BANG_LBRACK),
	1468: uint16(anon_sym_AMP_POUNDx),
	1469: uint16(7),
	1470: uint16(212),
	1471: uint16(1),
	1472: uint16(anon_sym_AMP),
	1473: uint16(214),
	1474: uint16(1),
	1475: uint16(anon_sym_AMP_POUND),
	1476: uint16(216),
	1477: uint16(1),
	1478: uint16(anon_sym_AMP_POUNDx),
	1479: uint16(306),
	1480: uint16(1),
	1481: uint16(anon_sym_DQUOTE),
	1482: uint16(308),
	1483: uint16(1),
	1484: uint16(aux_sym_PseudoAttValue_token1),
	1485: uint16(44),
	1486: uint16(2),
	1487: uint16(sym__Reference),
	1488: uint16(aux_sym_PseudoAttValue_repeat1),
	1489: uint16(148),
	1490: uint16(2),
	1491: uint16(sym_EntityRef),
	1492: uint16(sym_CharRef),
	1493: uint16(7),
	1494: uint16(212),
	1495: uint16(1),
	1496: uint16(anon_sym_AMP),
	1497: uint16(214),
	1498: uint16(1),
	1499: uint16(anon_sym_AMP_POUND),
	1500: uint16(216),
	1501: uint16(1),
	1502: uint16(anon_sym_AMP_POUNDx),
	1503: uint16(310),
	1504: uint16(1),
	1505: uint16(anon_sym_DQUOTE),
	1506: uint16(312),
	1507: uint16(1),
	1508: uint16(aux_sym_PseudoAttValue_token1),
	1509: uint16(37),
	1510: uint16(2),
	1511: uint16(sym__Reference),
	1512: uint16(aux_sym_PseudoAttValue_repeat1),
	1513: uint16(148),
	1514: uint16(2),
	1515: uint16(sym_EntityRef),
	1516: uint16(sym_CharRef),
	1517: uint16(1),
	1518: uint16(314),
	1519: uint16(9),
	1520: uint16(anon_sym_GT),
	1521: uint16(anon_sym_PIPE),
	1522: uint16(anon_sym_RPAREN),
	1523: uint16(anon_sym_STAR),
	1524: uint16(anon_sym_QMARK),
	1525: uint16(anon_sym_PLUS),
	1526: uint16(anon_sym_COMMA),
	1527: uint16(anon_sym_PERCENT),
	1528: uint16(sym__S),
	1529: uint16(4),
	1530: uint16(98),
	1531: uint16(1),
	1532: uint16(anon_sym_LT_QMARK),
	1533: uint16(158),
	1534: uint16(1),
	1535: uint16(anon_sym_LT),
	1536: uint16(102),
	1537: uint16(2),
	1538: uint16(sym_Comment),
	1539: uint16(sym__S),
	1540: uint16(25),
	1541: uint16(5),
	1542: uint16(sym__Misc),
	1543: uint16(sym_StyleSheetPI),
	1544: uint16(sym_XmlModelPI),
	1545: uint16(sym_PI),
	1546: uint16(aux_sym_document_repeat1),
	1547: uint16(4),
	1548: uint16(98),
	1549: uint16(1),
	1550: uint16(anon_sym_LT_QMARK),
	1551: uint16(316),
	1552: uint16(1),
	1554: uint16(102),
	1555: uint16(2),
	1556: uint16(sym_Comment),
	1557: uint16(sym__S),
	1558: uint16(25),
	1559: uint16(5),
	1560: uint16(sym__Misc),
	1561: uint16(sym_StyleSheetPI),
	1562: uint16(sym_XmlModelPI),
	1563: uint16(sym_PI),
	1564: uint16(aux_sym_document_repeat1),
	1565: uint16(1),
	1566: uint16(318),
	1567: uint16(9),
	1568: uint16(anon_sym_GT),
	1569: uint16(anon_sym_PIPE),
	1570: uint16(anon_sym_RPAREN),
	1571: uint16(anon_sym_STAR),
	1572: uint16(anon_sym_QMARK),
	1573: uint16(anon_sym_PLUS),
	1574: uint16(anon_sym_COMMA),
	1575: uint16(anon_sym_PERCENT),
	1576: uint16(sym__S),
	1577: uint16(7),
	1578: uint16(150),
	1579: uint16(1),
	1580: uint16(anon_sym_PERCENT),
	1581: uint16(320),
	1582: uint16(1),
	1583: uint16(anon_sym_SQUOTE),
	1584: uint16(322),
	1585: uint16(1),
	1586: uint16(anon_sym_DQUOTE),
	1587: uint16(326),
	1588: uint16(1),
	1589: uint16(anon_sym_POUNDFIXED),
	1590: uint16(291),
	1591: uint16(1),
	1592: uint16(sym_DefaultDecl),
	1593: uint16(324),
	1594: uint16(2),
	1595: uint16(anon_sym_POUNDREQUIRED),
	1596: uint16(anon_sym_POUNDIMPLIED),
	1597: uint16(290),
	1598: uint16(2),
	1599: uint16(sym_PEReference),
	1600: uint16(sym_AttValue),
	1601: uint16(1),
	1602: uint16(328),
	1603: uint16(9),
	1604: uint16(anon_sym_GT),
	1605: uint16(anon_sym_PIPE),
	1606: uint16(anon_sym_RPAREN),
	1607: uint16(anon_sym_STAR),
	1608: uint16(anon_sym_QMARK),
	1609: uint16(anon_sym_PLUS),
	1610: uint16(anon_sym_COMMA),
	1611: uint16(anon_sym_PERCENT),
	1612: uint16(sym__S),
	1613: uint16(1),
	1614: uint16(330),
	1615: uint16(9),
	1616: uint16(anon_sym_GT),
	1617: uint16(anon_sym_PIPE),
	1618: uint16(anon_sym_RPAREN),
	1619: uint16(anon_sym_STAR),
	1620: uint16(anon_sym_QMARK),
	1621: uint16(anon_sym_PLUS),
	1622: uint16(anon_sym_COMMA),
	1623: uint16(anon_sym_PERCENT),
	1624: uint16(sym__S),
	1625: uint16(1),
	1626: uint16(332),
	1627: uint16(9),
	1628: uint16(anon_sym_GT),
	1629: uint16(anon_sym_PIPE),
	1630: uint16(anon_sym_RPAREN),
	1631: uint16(anon_sym_STAR),
	1632: uint16(anon_sym_QMARK),
	1633: uint16(anon_sym_PLUS),
	1634: uint16(anon_sym_COMMA),
	1635: uint16(anon_sym_PERCENT),
	1636: uint16(sym__S),
	1637: uint16(2),
	1638: uint16(336),
	1639: uint16(3),
	1640: uint16(anon_sym_LT),
	1641: uint16(anon_sym_AMP),
	1642: uint16(anon_sym_AMP_POUND),
	1643: uint16(334),
	1644: uint16(6),
	1645: uint16(sym_Comment),
	1646: uint16(sym_CharData),
	1647: uint16(anon_sym_LT_QMARK),
	1648: uint16(anon_sym_LT_SLASH),
	1649: uint16(anon_sym_LT_BANG_LBRACK),
	1650: uint16(anon_sym_AMP_POUNDx),
	1651: uint16(2),
	1652: uint16(340),
	1653: uint16(3),
	1654: uint16(anon_sym_LT),
	1655: uint16(anon_sym_AMP),
	1656: uint16(anon_sym_AMP_POUND),
	1657: uint16(338),
	1658: uint16(6),
	1659: uint16(sym_Comment),
	1660: uint16(sym_CharData),
	1661: uint16(anon_sym_LT_QMARK),
	1662: uint16(anon_sym_LT_SLASH),
	1663: uint16(anon_sym_LT_BANG_LBRACK),
	1664: uint16(anon_sym_AMP_POUNDx),
	1665: uint16(2),
	1666: uint16(344),
	1667: uint16(3),
	1668: uint16(anon_sym_LT),
	1669: uint16(anon_sym_AMP),
	1670: uint16(anon_sym_AMP_POUND),
	1671: uint16(342),
	1672: uint16(6),
	1673: uint16(sym_Comment),
	1674: uint16(sym_CharData),
	1675: uint16(anon_sym_LT_QMARK),
	1676: uint16(anon_sym_LT_SLASH),
	1677: uint16(anon_sym_LT_BANG_LBRACK),
	1678: uint16(anon_sym_AMP_POUNDx),
	1679: uint16(2),
	1680: uint16(348),
	1681: uint16(3),
	1682: uint16(anon_sym_LT),
	1683: uint16(anon_sym_AMP),
	1684: uint16(anon_sym_AMP_POUND),
	1685: uint16(346),
	1686: uint16(6),
	1687: uint16(sym_Comment),
	1688: uint16(sym_CharData),
	1689: uint16(anon_sym_LT_QMARK),
	1690: uint16(anon_sym_LT_SLASH),
	1691: uint16(anon_sym_LT_BANG_LBRACK),
	1692: uint16(anon_sym_AMP_POUNDx),
	1693: uint16(7),
	1694: uint16(210),
	1695: uint16(1),
	1696: uint16(aux_sym_PseudoAttValue_token1),
	1697: uint16(212),
	1698: uint16(1),
	1699: uint16(anon_sym_AMP),
	1700: uint16(214),
	1701: uint16(1),
	1702: uint16(anon_sym_AMP_POUND),
	1703: uint16(216),
	1704: uint16(1),
	1705: uint16(anon_sym_AMP_POUNDx),
	1706: uint16(350),
	1707: uint16(1),
	1708: uint16(anon_sym_DQUOTE),
	1709: uint16(46),
	1710: uint16(2),
	1711: uint16(sym__Reference),
	1712: uint16(aux_sym_PseudoAttValue_repeat1),
	1713: uint16(148),
	1714: uint16(2),
	1715: uint16(sym_EntityRef),
	1716: uint16(sym_CharRef),
	1717: uint16(2),
	1718: uint16(354),
	1719: uint16(3),
	1720: uint16(anon_sym_LT),
	1721: uint16(anon_sym_AMP),
	1722: uint16(anon_sym_AMP_POUND),
	1723: uint16(352),
	1724: uint16(6),
	1725: uint16(sym_Comment),
	1726: uint16(sym_CharData),
	1727: uint16(anon_sym_LT_QMARK),
	1728: uint16(anon_sym_LT_SLASH),
	1729: uint16(anon_sym_LT_BANG_LBRACK),
	1730: uint16(anon_sym_AMP_POUNDx),
	1731: uint16(2),
	1732: uint16(358),
	1733: uint16(3),
	1734: uint16(anon_sym_LT),
	1735: uint16(anon_sym_AMP),
	1736: uint16(anon_sym_AMP_POUND),
	1737: uint16(356),
	1738: uint16(6),
	1739: uint16(sym_Comment),
	1740: uint16(sym_CharData),
	1741: uint16(anon_sym_LT_QMARK),
	1742: uint16(anon_sym_LT_SLASH),
	1743: uint16(anon_sym_LT_BANG_LBRACK),
	1744: uint16(anon_sym_AMP_POUNDx),
	1745: uint16(2),
	1746: uint16(362),
	1747: uint16(3),
	1748: uint16(anon_sym_LT),
	1749: uint16(anon_sym_AMP),
	1750: uint16(anon_sym_AMP_POUND),
	1751: uint16(360),
	1752: uint16(6),
	1753: uint16(sym_Comment),
	1754: uint16(sym_CharData),
	1755: uint16(anon_sym_LT_QMARK),
	1756: uint16(anon_sym_LT_SLASH),
	1757: uint16(anon_sym_LT_BANG_LBRACK),
	1758: uint16(anon_sym_AMP_POUNDx),
	1759: uint16(2),
	1760: uint16(366),
	1761: uint16(3),
	1762: uint16(anon_sym_LT),
	1763: uint16(anon_sym_AMP),
	1764: uint16(anon_sym_AMP_POUND),
	1765: uint16(364),
	1766: uint16(6),
	1767: uint16(sym_Comment),
	1768: uint16(sym_CharData),
	1769: uint16(anon_sym_LT_QMARK),
	1770: uint16(anon_sym_LT_SLASH),
	1771: uint16(anon_sym_LT_BANG_LBRACK),
	1772: uint16(anon_sym_AMP_POUNDx),
	1773: uint16(2),
	1774: uint16(370),
	1775: uint16(3),
	1776: uint16(anon_sym_LT),
	1777: uint16(anon_sym_AMP),
	1778: uint16(anon_sym_AMP_POUND),
	1779: uint16(368),
	1780: uint16(6),
	1781: uint16(sym_Comment),
	1782: uint16(sym_CharData),
	1783: uint16(anon_sym_LT_QMARK),
	1784: uint16(anon_sym_LT_SLASH),
	1785: uint16(anon_sym_LT_BANG_LBRACK),
	1786: uint16(anon_sym_AMP_POUNDx),
	1787: uint16(2),
	1788: uint16(374),
	1789: uint16(3),
	1790: uint16(anon_sym_LT),
	1791: uint16(anon_sym_AMP),
	1792: uint16(anon_sym_AMP_POUND),
	1793: uint16(372),
	1794: uint16(6),
	1795: uint16(sym_Comment),
	1796: uint16(sym_CharData),
	1797: uint16(anon_sym_LT_QMARK),
	1798: uint16(anon_sym_LT_SLASH),
	1799: uint16(anon_sym_LT_BANG_LBRACK),
	1800: uint16(anon_sym_AMP_POUNDx),
	1801: uint16(7),
	1802: uint16(204),
	1803: uint16(1),
	1804: uint16(anon_sym_AMP),
	1805: uint16(206),
	1806: uint16(1),
	1807: uint16(anon_sym_AMP_POUND),
	1808: uint16(208),
	1809: uint16(1),
	1810: uint16(anon_sym_AMP_POUNDx),
	1811: uint16(306),
	1812: uint16(1),
	1813: uint16(anon_sym_SQUOTE),
	1814: uint16(376),
	1815: uint16(1),
	1816: uint16(aux_sym_PseudoAttValue_token2),
	1817: uint16(43),
	1818: uint16(2),
	1819: uint16(sym__Reference),
	1820: uint16(aux_sym_PseudoAttValue_repeat2),
	1821: uint16(142),
	1822: uint16(2),
	1823: uint16(sym_EntityRef),
	1824: uint16(sym_CharRef),
	1825: uint16(7),
	1826: uint16(204),
	1827: uint16(1),
	1828: uint16(anon_sym_AMP),
	1829: uint16(206),
	1830: uint16(1),
	1831: uint16(anon_sym_AMP_POUND),
	1832: uint16(208),
	1833: uint16(1),
	1834: uint16(anon_sym_AMP_POUNDx),
	1835: uint16(378),
	1836: uint16(1),
	1837: uint16(anon_sym_SQUOTE),
	1838: uint16(380),
	1839: uint16(1),
	1840: uint16(aux_sym_PseudoAttValue_token2),
	1841: uint16(84),
	1842: uint16(2),
	1843: uint16(sym__Reference),
	1844: uint16(aux_sym_PseudoAttValue_repeat2),
	1845: uint16(142),
	1846: uint16(2),
	1847: uint16(sym_EntityRef),
	1848: uint16(sym_CharRef),
	1849: uint16(7),
	1850: uint16(212),
	1851: uint16(1),
	1852: uint16(anon_sym_AMP),
	1853: uint16(214),
	1854: uint16(1),
	1855: uint16(anon_sym_AMP_POUND),
	1856: uint16(216),
	1857: uint16(1),
	1858: uint16(anon_sym_AMP_POUNDx),
	1859: uint16(378),
	1860: uint16(1),
	1861: uint16(anon_sym_DQUOTE),
	1862: uint16(382),
	1863: uint16(1),
	1864: uint16(aux_sym_PseudoAttValue_token1),
	1865: uint16(74),
	1866: uint16(2),
	1867: uint16(sym__Reference),
	1868: uint16(aux_sym_PseudoAttValue_repeat1),
	1869: uint16(148),
	1870: uint16(2),
	1871: uint16(sym_EntityRef),
	1872: uint16(sym_CharRef),
	1873: uint16(7),
	1874: uint16(202),
	1875: uint16(1),
	1876: uint16(aux_sym_PseudoAttValue_token2),
	1877: uint16(204),
	1878: uint16(1),
	1879: uint16(anon_sym_AMP),
	1880: uint16(206),
	1881: uint16(1),
	1882: uint16(anon_sym_AMP_POUND),
	1883: uint16(208),
	1884: uint16(1),
	1885: uint16(anon_sym_AMP_POUNDx),
	1886: uint16(350),
	1887: uint16(1),
	1888: uint16(anon_sym_SQUOTE),
	1889: uint16(45),
	1890: uint16(2),
	1891: uint16(sym__Reference),
	1892: uint16(aux_sym_PseudoAttValue_repeat2),
	1893: uint16(142),
	1894: uint16(2),
	1895: uint16(sym_EntityRef),
	1896: uint16(sym_CharRef),
	1897: uint16(7),
	1898: uint16(204),
	1899: uint16(1),
	1900: uint16(anon_sym_AMP),
	1901: uint16(206),
	1902: uint16(1),
	1903: uint16(anon_sym_AMP_POUND),
	1904: uint16(208),
	1905: uint16(1),
	1906: uint16(anon_sym_AMP_POUNDx),
	1907: uint16(310),
	1908: uint16(1),
	1909: uint16(anon_sym_SQUOTE),
	1910: uint16(384),
	1911: uint16(1),
	1912: uint16(aux_sym_PseudoAttValue_token2),
	1913: uint16(36),
	1914: uint16(2),
	1915: uint16(sym__Reference),
	1916: uint16(aux_sym_PseudoAttValue_repeat2),
	1917: uint16(142),
	1918: uint16(2),
	1919: uint16(sym_EntityRef),
	1920: uint16(sym_CharRef),
	1921: uint16(6),
	1922: uint16(150),
	1923: uint16(1),
	1924: uint16(anon_sym_PERCENT),
	1925: uint16(388),
	1926: uint16(1),
	1927: uint16(anon_sym_RPAREN),
	1928: uint16(390),
	1929: uint16(1),
	1930: uint16(sym__S),
	1931: uint16(110),
	1932: uint16(1),
	1933: uint16(aux_sym__choice_repeat1),
	1934: uint16(386),
	1935: uint16(2),
	1936: uint16(anon_sym_PIPE),
	1937: uint16(anon_sym_COMMA),
	1938: uint16(153),
	1939: uint16(2),
	1940: uint16(sym_PEReference),
	1941: uint16(aux_sym__choice_repeat2),
	1942: uint16(2),
	1943: uint16(140),
	1944: uint16(3),
	1945: uint16(anon_sym_STAR),
	1946: uint16(anon_sym_QMARK),
	1947: uint16(anon_sym_PLUS),
	1948: uint16(138),
	1949: uint16(5),
	1950: uint16(anon_sym_PIPE),
	1951: uint16(anon_sym_RPAREN),
	1952: uint16(anon_sym_COMMA),
	1953: uint16(anon_sym_PERCENT),
	1954: uint16(sym__S),
	1955: uint16(6),
	1956: uint16(150),
	1957: uint16(1),
	1958: uint16(anon_sym_PERCENT),
	1959: uint16(388),
	1960: uint16(1),
	1961: uint16(anon_sym_RPAREN),
	1962: uint16(390),
	1963: uint16(1),
	1964: uint16(sym__S),
	1965: uint16(90),
	1966: uint16(1),
	1967: uint16(aux_sym__choice_repeat1),
	1968: uint16(386),
	1969: uint16(2),
	1970: uint16(anon_sym_PIPE),
	1971: uint16(anon_sym_COMMA),
	1972: uint16(153),
	1973: uint16(2),
	1974: uint16(sym_PEReference),
	1975: uint16(aux_sym__choice_repeat2),
	1976: uint16(6),
	1977: uint16(150),
	1978: uint16(1),
	1979: uint16(anon_sym_PERCENT),
	1980: uint16(392),
	1981: uint16(1),
	1982: uint16(anon_sym_RPAREN),
	1983: uint16(394),
	1984: uint16(1),
	1985: uint16(sym__S),
	1986: uint16(86),
	1987: uint16(1),
	1988: uint16(aux_sym__choice_repeat1),
	1989: uint16(386),
	1990: uint16(2),
	1991: uint16(anon_sym_PIPE),
	1992: uint16(anon_sym_COMMA),
	1993: uint16(138),
	1994: uint16(2),
	1995: uint16(sym_PEReference),
	1996: uint16(aux_sym__choice_repeat2),
	1997: uint16(6),
	1998: uint16(150),
	1999: uint16(1),
	2000: uint16(anon_sym_PERCENT),
	2001: uint16(396),
	2002: uint16(1),
	2003: uint16(anon_sym_RPAREN),
	2004: uint16(398),
	2005: uint16(1),
	2006: uint16(sym__S),
	2007: uint16(110),
	2008: uint16(1),
	2009: uint16(aux_sym__choice_repeat1),
	2010: uint16(386),
	2011: uint16(2),
	2012: uint16(anon_sym_PIPE),
	2013: uint16(anon_sym_COMMA),
	2014: uint16(127),
	2015: uint16(2),
	2016: uint16(sym_PEReference),
	2017: uint16(aux_sym__choice_repeat2),
	2018: uint16(8),
	2019: uint16(150),
	2020: uint16(1),
	2021: uint16(anon_sym_PERCENT),
	2022: uint16(400),
	2023: uint16(1),
	2024: uint16(sym_Name),
	2025: uint16(402),
	2026: uint16(1),
	2027: uint16(anon_sym_LPAREN),
	2028: uint16(404),
	2029: uint16(1),
	2030: uint16(anon_sym_POUNDPCDATA),
	2031: uint16(406),
	2032: uint16(1),
	2033: uint16(sym__S),
	2034: uint16(26),
	2035: uint16(1),
	2036: uint16(sym_PEReference),
	2037: uint16(87),
	2038: uint16(1),
	2039: uint16(sym__choice),
	2040: uint16(89),
	2041: uint16(1),
	2042: uint16(sym__cp),
	2043: uint16(6),
	2044: uint16(150),
	2045: uint16(1),
	2046: uint16(anon_sym_PERCENT),
	2047: uint16(400),
	2048: uint16(1),
	2049: uint16(sym_Name),
	2050: uint16(402),
	2051: uint16(1),
	2052: uint16(anon_sym_LPAREN),
	2053: uint16(408),
	2054: uint16(1),
	2055: uint16(sym__S),
	2056: uint16(89),
	2057: uint16(1),
	2058: uint16(sym__cp),
	2059: uint16(87),
	2060: uint16(2),
	2061: uint16(sym__choice),
	2062: uint16(sym_PEReference),
	2063: uint16(6),
	2064: uint16(150),
	2065: uint16(1),
	2066: uint16(anon_sym_PERCENT),
	2067: uint16(400),
	2068: uint16(1),
	2069: uint16(sym_Name),
	2070: uint16(402),
	2071: uint16(1),
	2072: uint16(anon_sym_LPAREN),
	2073: uint16(410),
	2074: uint16(1),
	2075: uint16(sym__S),
	2076: uint16(133),
	2077: uint16(1),
	2078: uint16(sym__cp),
	2079: uint16(87),
	2080: uint16(2),
	2081: uint16(sym__choice),
	2082: uint16(sym_PEReference),
	2083: uint16(7),
	2084: uint16(150),
	2085: uint16(1),
	2086: uint16(anon_sym_PERCENT),
	2087: uint16(412),
	2088: uint16(1),
	2089: uint16(anon_sym_PIPE),
	2090: uint16(414),
	2091: uint16(1),
	2092: uint16(anon_sym_RPAREN),
	2093: uint16(416),
	2094: uint16(1),
	2095: uint16(sym__S),
	2096: uint16(96),
	2097: uint16(1),
	2098: uint16(aux_sym_Mixed_repeat1),
	2099: uint16(174),
	2100: uint16(1),
	2101: uint16(aux_sym_Mixed_repeat2),
	2102: uint16(246),
	2103: uint16(1),
	2104: uint16(sym_PEReference),
	2105: uint16(7),
	2106: uint16(150),
	2107: uint16(1),
	2108: uint16(anon_sym_PERCENT),
	2109: uint16(400),
	2110: uint16(1),
	2111: uint16(sym_Name),
	2112: uint16(402),
	2113: uint16(1),
	2114: uint16(anon_sym_LPAREN),
	2115: uint16(418),
	2116: uint16(1),
	2117: uint16(anon_sym_POUNDPCDATA),
	2118: uint16(31),
	2119: uint16(1),
	2120: uint16(sym_PEReference),
	2121: uint16(87),
	2122: uint16(1),
	2123: uint16(sym__choice),
	2124: uint16(88),
	2125: uint16(1),
	2126: uint16(sym__cp),
	2127: uint16(7),
	2128: uint16(150),
	2129: uint16(1),
	2130: uint16(anon_sym_PERCENT),
	2131: uint16(412),
	2132: uint16(1),
	2133: uint16(anon_sym_PIPE),
	2134: uint16(420),
	2135: uint16(1),
	2136: uint16(anon_sym_RPAREN),
	2137: uint16(422),
	2138: uint16(1),
	2139: uint16(sym__S),
	2140: uint16(140),
	2141: uint16(1),
	2142: uint16(aux_sym_Mixed_repeat1),
	2143: uint16(182),
	2144: uint16(1),
	2145: uint16(aux_sym_Mixed_repeat2),
	2146: uint16(246),
	2147: uint16(1),
	2148: uint16(sym_PEReference),
	2149: uint16(7),
	2150: uint16(150),
	2151: uint16(1),
	2152: uint16(anon_sym_PERCENT),
	2153: uint16(412),
	2154: uint16(1),
	2155: uint16(anon_sym_PIPE),
	2156: uint16(424),
	2157: uint16(1),
	2158: uint16(anon_sym_RPAREN),
	2159: uint16(426),
	2160: uint16(1),
	2161: uint16(sym__S),
	2162: uint16(99),
	2163: uint16(1),
	2164: uint16(aux_sym_Mixed_repeat1),
	2165: uint16(181),
	2166: uint16(1),
	2167: uint16(aux_sym_Mixed_repeat2),
	2168: uint16(246),
	2169: uint16(1),
	2170: uint16(sym_PEReference),
	2171: uint16(6),
	2172: uint16(150),
	2173: uint16(1),
	2174: uint16(anon_sym_PERCENT),
	2175: uint16(400),
	2176: uint16(1),
	2177: uint16(sym_Name),
	2178: uint16(402),
	2179: uint16(1),
	2180: uint16(anon_sym_LPAREN),
	2181: uint16(428),
	2182: uint16(1),
	2183: uint16(sym__S),
	2184: uint16(155),
	2185: uint16(1),
	2186: uint16(sym__cp),
	2187: uint16(87),
	2188: uint16(2),
	2189: uint16(sym__choice),
	2190: uint16(sym_PEReference),
	2191: uint16(7),
	2192: uint16(150),
	2193: uint16(1),
	2194: uint16(anon_sym_PERCENT),
	2195: uint16(412),
	2196: uint16(1),
	2197: uint16(anon_sym_PIPE),
	2198: uint16(430),
	2199: uint16(1),
	2200: uint16(anon_sym_RPAREN),
	2201: uint16(432),
	2202: uint16(1),
	2203: uint16(sym__S),
	2204: uint16(140),
	2205: uint16(1),
	2206: uint16(aux_sym_Mixed_repeat1),
	2207: uint16(192),
	2208: uint16(1),
	2209: uint16(aux_sym_Mixed_repeat2),
	2210: uint16(246),
	2211: uint16(1),
	2212: uint16(sym_PEReference),
	2213: uint16(2),
	2214: uint16(374),
	2215: uint16(2),
	2216: uint16(anon_sym_AMP),
	2217: uint16(anon_sym_AMP_POUND),
	2218: uint16(372),
	2219: uint16(4),
	2220: uint16(anon_sym_SQUOTE),
	2221: uint16(anon_sym_PERCENT),
	2222: uint16(aux_sym_EntityValue_token2),
	2223: uint16(anon_sym_AMP_POUNDx),
	2224: uint16(2),
	2225: uint16(436),
	2226: uint16(1),
	2227: uint16(anon_sym_LT),
	2228: uint16(434),
	2229: uint16(5),
	2230: uint16(sym_Comment),
	2232: uint16(anon_sym_LT_QMARK),
	2233: uint16(anon_sym_LT_BANG),
	2234: uint16(sym__S),
	2235: uint16(5),
	2236: uint16(150),
	2237: uint16(1),
	2238: uint16(anon_sym_PERCENT),
	2239: uint16(400),
	2240: uint16(1),
	2241: uint16(sym_Name),
	2242: uint16(402),
	2243: uint16(1),
	2244: uint16(anon_sym_LPAREN),
	2245: uint16(88),
	2246: uint16(1),
	2247: uint16(sym__cp),
	2248: uint16(87),
	2249: uint16(2),
	2250: uint16(sym__choice),
	2251: uint16(sym_PEReference),
	2252: uint16(6),
	2253: uint16(438),
	2254: uint16(1),
	2255: uint16(anon_sym_SQUOTE),
	2256: uint16(440),
	2257: uint16(1),
	2258: uint16(anon_sym_DQUOTE),
	2259: uint16(442),
	2260: uint16(1),
	2261: uint16(anon_sym_SYSTEM),
	2262: uint16(444),
	2263: uint16(1),
	2264: uint16(anon_sym_PUBLIC),
	2265: uint16(226),
	2266: uint16(1),
	2267: uint16(sym_ExternalID),
	2268: uint16(304),
	2269: uint16(1),
	2270: uint16(sym_EntityValue),
	2271: uint16(2),
	2272: uint16(448),
	2273: uint16(1),
	2274: uint16(anon_sym_LT),
	2275: uint16(446),
	2276: uint16(5),
	2277: uint16(sym_Comment),
	2279: uint16(anon_sym_LT_QMARK),
	2280: uint16(anon_sym_LT_BANG),
	2281: uint16(sym__S),
	2282: uint16(2),
	2283: uint16(452),
	2284: uint16(1),
	2285: uint16(anon_sym_LT),
	2286: uint16(450),
	2287: uint16(5),
	2288: uint16(sym_Comment),
	2290: uint16(anon_sym_LT_QMARK),
	2291: uint16(anon_sym_LT_BANG),
	2292: uint16(sym__S),
	2293: uint16(5),
	2294: uint16(150),
	2295: uint16(1),
	2296: uint16(anon_sym_PERCENT),
	2297: uint16(442),
	2298: uint16(1),
	2299: uint16(anon_sym_SYSTEM),
	2300: uint16(454),
	2301: uint16(1),
	2302: uint16(anon_sym_PUBLIC),
	2303: uint16(384),
	2304: uint16(1),
	2305: uint16(sym_PEReference),
	2306: uint16(288),
	2307: uint16(2),
	2308: uint16(sym_ExternalID),
	2309: uint16(sym_PublicID),
	2310: uint16(2),
	2311: uint16(344),
	2312: uint16(1),
	2313: uint16(anon_sym_LT),
	2314: uint16(342),
	2315: uint16(5),
	2316: uint16(sym_Comment),
	2318: uint16(anon_sym_LT_QMARK),
	2319: uint16(anon_sym_LT_BANG),
	2320: uint16(sym__S),
	2321: uint16(5),
	2322: uint16(150),
	2323: uint16(1),
	2324: uint16(anon_sym_PERCENT),
	2325: uint16(400),
	2326: uint16(1),
	2327: uint16(sym_Name),
	2328: uint16(402),
	2329: uint16(1),
	2330: uint16(anon_sym_LPAREN),
	2331: uint16(133),
	2332: uint16(1),
	2333: uint16(sym__cp),
	2334: uint16(87),
	2335: uint16(2),
	2336: uint16(sym__choice),
	2337: uint16(sym_PEReference),
	2338: uint16(5),
	2339: uint16(438),
	2340: uint16(1),
	2341: uint16(anon_sym_SQUOTE),
	2342: uint16(440),
	2343: uint16(1),
	2344: uint16(anon_sym_DQUOTE),
	2345: uint16(442),
	2346: uint16(1),
	2347: uint16(anon_sym_SYSTEM),
	2348: uint16(444),
	2349: uint16(1),
	2350: uint16(anon_sym_PUBLIC),
	2351: uint16(345),
	2352: uint16(2),
	2353: uint16(sym_EntityValue),
	2354: uint16(sym_ExternalID),
	2355: uint16(4),
	2356: uint16(461),
	2357: uint16(1),
	2358: uint16(sym__S),
	2359: uint16(110),
	2360: uint16(1),
	2361: uint16(aux_sym__choice_repeat1),
	2362: uint16(456),
	2363: uint16(2),
	2364: uint16(anon_sym_PIPE),
	2365: uint16(anon_sym_COMMA),
	2366: uint16(459),
	2367: uint16(2),
	2368: uint16(anon_sym_RPAREN),
	2369: uint16(anon_sym_PERCENT),
	2370: uint16(2),
	2371: uint16(362),
	2372: uint16(1),
	2373: uint16(anon_sym_LT),
	2374: uint16(360),
	2375: uint16(5),
	2376: uint16(sym_Comment),
	2378: uint16(anon_sym_LT_QMARK),
	2379: uint16(anon_sym_LT_BANG),
	2380: uint16(sym__S),
	2381: uint16(2),
	2382: uint16(466),
	2383: uint16(1),
	2384: uint16(anon_sym_LT),
	2385: uint16(464),
	2386: uint16(5),
	2387: uint16(sym_Comment),
	2389: uint16(anon_sym_LT_QMARK),
	2390: uint16(anon_sym_LT_BANG),
	2391: uint16(sym__S),
	2392: uint16(5),
	2393: uint16(150),
	2394: uint16(1),
	2395: uint16(anon_sym_PERCENT),
	2396: uint16(400),
	2397: uint16(1),
	2398: uint16(sym_Name),
	2399: uint16(402),
	2400: uint16(1),
	2401: uint16(anon_sym_LPAREN),
	2402: uint16(135),
	2403: uint16(1),
	2404: uint16(sym__cp),
	2405: uint16(87),
	2406: uint16(2),
	2407: uint16(sym__choice),
	2408: uint16(sym_PEReference),
	2409: uint16(6),
	2410: uint16(150),
	2411: uint16(1),
	2412: uint16(anon_sym_PERCENT),
	2413: uint16(468),
	2414: uint16(1),
	2415: uint16(sym_Name),
	2416: uint16(470),
	2417: uint16(1),
	2418: uint16(anon_sym_PIPE),
	2419: uint16(472),
	2420: uint16(1),
	2421: uint16(sym__S),
	2422: uint16(116),
	2423: uint16(1),
	2424: uint16(aux_sym_NotationType_repeat1),
	2425: uint16(317),
	2426: uint16(1),
	2427: uint16(sym_PEReference),
	2428: uint16(6),
	2429: uint16(150),
	2430: uint16(1),
	2431: uint16(anon_sym_PERCENT),
	2432: uint16(470),
	2433: uint16(1),
	2434: uint16(anon_sym_PIPE),
	2435: uint16(472),
	2436: uint16(1),
	2437: uint16(sym__S),
	2438: uint16(474),
	2439: uint16(1),
	2440: uint16(sym_Name),
	2441: uint16(117),
	2442: uint16(1),
	2443: uint16(aux_sym_NotationType_repeat1),
	2444: uint16(323),
	2445: uint16(1),
	2446: uint16(sym_PEReference),
	2447: uint16(6),
	2448: uint16(150),
	2449: uint16(1),
	2450: uint16(anon_sym_PERCENT),
	2451: uint16(470),
	2452: uint16(1),
	2453: uint16(anon_sym_PIPE),
	2454: uint16(472),
	2455: uint16(1),
	2456: uint16(sym__S),
	2457: uint16(474),
	2458: uint16(1),
	2459: uint16(sym_Name),
	2460: uint16(141),
	2461: uint16(1),
	2462: uint16(aux_sym_NotationType_repeat1),
	2463: uint16(323),
	2464: uint16(1),
	2465: uint16(sym_PEReference),
	2466: uint16(6),
	2467: uint16(150),
	2468: uint16(1),
	2469: uint16(anon_sym_PERCENT),
	2470: uint16(470),
	2471: uint16(1),
	2472: uint16(anon_sym_PIPE),
	2473: uint16(472),
	2474: uint16(1),
	2475: uint16(sym__S),
	2476: uint16(476),
	2477: uint16(1),
	2478: uint16(sym_Name),
	2479: uint16(141),
	2480: uint16(1),
	2481: uint16(aux_sym_NotationType_repeat1),
	2482: uint16(330),
	2483: uint16(1),
	2484: uint16(sym_PEReference),
	2485: uint16(2),
	2486: uint16(480),
	2487: uint16(1),
	2488: uint16(anon_sym_LT),
	2489: uint16(478),
	2490: uint16(5),
	2491: uint16(sym_Comment),
	2493: uint16(anon_sym_LT_QMARK),
	2494: uint16(anon_sym_LT_BANG),
	2495: uint16(sym__S),
	2496: uint16(2),
	2497: uint16(484),
	2498: uint16(1),
	2499: uint16(anon_sym_LT),
	2500: uint16(482),
	2501: uint16(5),
	2502: uint16(sym_Comment),
	2504: uint16(anon_sym_LT_QMARK),
	2505: uint16(anon_sym_LT_BANG),
	2506: uint16(sym__S),
	2507: uint16(2),
	2508: uint16(486),
	2509: uint16(2),
	2510: uint16(anon_sym_AMP),
	2511: uint16(anon_sym_AMP_POUND),
	2512: uint16(194),
	2513: uint16(4),
	2514: uint16(anon_sym_SQUOTE),
	2515: uint16(anon_sym_PERCENT),
	2516: uint16(aux_sym_EntityValue_token2),
	2517: uint16(anon_sym_AMP_POUNDx),
	2518: uint16(2),
	2519: uint16(294),
	2520: uint16(2),
	2521: uint16(anon_sym_AMP),
	2522: uint16(anon_sym_AMP_POUND),
	2523: uint16(292),
	2524: uint16(4),
	2525: uint16(anon_sym_DQUOTE),
	2526: uint16(anon_sym_PERCENT),
	2527: uint16(aux_sym_EntityValue_token1),
	2528: uint16(anon_sym_AMP_POUNDx),
	2529: uint16(2),
	2530: uint16(290),
	2531: uint16(2),
	2532: uint16(anon_sym_AMP),
	2533: uint16(anon_sym_AMP_POUND),
	2534: uint16(288),
	2535: uint16(4),
	2536: uint16(anon_sym_SQUOTE),
	2537: uint16(anon_sym_PERCENT),
	2538: uint16(aux_sym_EntityValue_token2),
	2539: uint16(anon_sym_AMP_POUNDx),
	2540: uint16(2),
	2541: uint16(294),
	2542: uint16(2),
	2543: uint16(anon_sym_AMP),
	2544: uint16(anon_sym_AMP_POUND),
	2545: uint16(292),
	2546: uint16(4),
	2547: uint16(anon_sym_SQUOTE),
	2548: uint16(anon_sym_PERCENT),
	2549: uint16(aux_sym_EntityValue_token2),
	2550: uint16(anon_sym_AMP_POUNDx),
	2551: uint16(2),
	2552: uint16(486),
	2553: uint16(2),
	2554: uint16(anon_sym_AMP),
	2555: uint16(anon_sym_AMP_POUND),
	2556: uint16(194),
	2557: uint16(4),
	2558: uint16(anon_sym_DQUOTE),
	2559: uint16(anon_sym_PERCENT),
	2560: uint16(aux_sym_EntityValue_token1),
	2561: uint16(anon_sym_AMP_POUNDx),
	2562: uint16(2),
	2563: uint16(374),
	2564: uint16(2),
	2565: uint16(anon_sym_AMP),
	2566: uint16(anon_sym_AMP_POUND),
	2567: uint16(372),
	2568: uint16(4),
	2569: uint16(anon_sym_DQUOTE),
	2570: uint16(anon_sym_PERCENT),
	2571: uint16(aux_sym_EntityValue_token1),
	2572: uint16(anon_sym_AMP_POUNDx),
	2573: uint16(2),
	2574: uint16(290),
	2575: uint16(2),
	2576: uint16(anon_sym_AMP),
	2577: uint16(anon_sym_AMP_POUND),
	2578: uint16(288),
	2579: uint16(4),
	2580: uint16(anon_sym_DQUOTE),
	2581: uint16(anon_sym_PERCENT),
	2582: uint16(aux_sym_EntityValue_token1),
	2583: uint16(anon_sym_AMP_POUNDx),
	2584: uint16(4),
	2585: uint16(150),
	2586: uint16(1),
	2587: uint16(anon_sym_PERCENT),
	2588: uint16(488),
	2589: uint16(1),
	2590: uint16(anon_sym_RPAREN),
	2591: uint16(490),
	2592: uint16(1),
	2593: uint16(sym__S),
	2594: uint16(158),
	2595: uint16(2),
	2596: uint16(sym_PEReference),
	2597: uint16(aux_sym__choice_repeat2),
	2598: uint16(2),
	2599: uint16(494),
	2600: uint16(1),
	2601: uint16(anon_sym_LT),
	2602: uint16(492),
	2603: uint16(4),
	2604: uint16(sym_Comment),
	2605: uint16(anon_sym_LT_QMARK),
	2606: uint16(anon_sym_LT_BANG),
	2607: uint16(sym__S),
	2608: uint16(2),
	2609: uint16(498),
	2610: uint16(1),
	2611: uint16(anon_sym_LT),
	2612: uint16(496),
	2613: uint16(4),
	2614: uint16(sym_Comment),
	2615: uint16(anon_sym_LT_QMARK),
	2616: uint16(anon_sym_LT_BANG),
	2617: uint16(sym__S),
	2618: uint16(5),
	2619: uint16(150),
	2620: uint16(1),
	2621: uint16(anon_sym_PERCENT),
	2622: uint16(500),
	2623: uint16(1),
	2624: uint16(anon_sym_PIPE),
	2625: uint16(502),
	2626: uint16(1),
	2627: uint16(anon_sym_RPAREN),
	2628: uint16(199),
	2629: uint16(1),
	2630: uint16(aux_sym_Mixed_repeat2),
	2631: uint16(246),
	2632: uint16(1),
	2633: uint16(sym_PEReference),
	2634: uint16(4),
	2635: uint16(150),
	2636: uint16(1),
	2637: uint16(anon_sym_PERCENT),
	2638: uint16(488),
	2639: uint16(1),
	2640: uint16(anon_sym_RPAREN),
	2641: uint16(230),
	2642: uint16(1),
	2643: uint16(sym_PEReference),
	2644: uint16(504),
	2645: uint16(2),
	2646: uint16(anon_sym_PIPE),
	2647: uint16(anon_sym_COMMA),
	2648: uint16(1),
	2649: uint16(506),
	2650: uint16(5),
	2651: uint16(anon_sym_PIPE),
	2652: uint16(anon_sym_RPAREN),
	2653: uint16(anon_sym_COMMA),
	2654: uint16(anon_sym_PERCENT),
	2655: uint16(sym__S),
	2656: uint16(1),
	2657: uint16(508),
	2658: uint16(5),
	2659: uint16(anon_sym_PIPE),
	2660: uint16(anon_sym_RPAREN),
	2661: uint16(anon_sym_COMMA),
	2662: uint16(anon_sym_PERCENT),
	2663: uint16(sym__S),
	2664: uint16(1),
	2665: uint16(194),
	2666: uint16(5),
	2667: uint16(sym_Comment),
	2668: uint16(anon_sym_LT_QMARK),
	2669: uint16(anon_sym_LT_BANG),
	2670: uint16(anon_sym_RBRACK),
	2671: uint16(sym__S),
	2672: uint16(1),
	2673: uint16(510),
	2674: uint16(5),
	2675: uint16(anon_sym_PIPE),
	2676: uint16(anon_sym_RPAREN),
	2677: uint16(anon_sym_COMMA),
	2678: uint16(anon_sym_PERCENT),
	2679: uint16(sym__S),
	2680: uint16(4),
	2681: uint16(150),
	2682: uint16(1),
	2683: uint16(anon_sym_PERCENT),
	2684: uint16(388),
	2685: uint16(1),
	2686: uint16(anon_sym_RPAREN),
	2687: uint16(230),
	2688: uint16(1),
	2689: uint16(sym_PEReference),
	2690: uint16(504),
	2691: uint16(2),
	2692: uint16(anon_sym_PIPE),
	2693: uint16(anon_sym_COMMA),
	2694: uint16(2),
	2695: uint16(514),
	2696: uint16(1),
	2697: uint16(anon_sym_LT),
	2698: uint16(512),
	2699: uint16(4),
	2700: uint16(sym_Comment),
	2701: uint16(anon_sym_LT_QMARK),
	2702: uint16(anon_sym_LT_BANG),
	2703: uint16(sym__S),
	2704: uint16(4),
	2705: uint16(150),
	2706: uint16(1),
	2707: uint16(anon_sym_PERCENT),
	2708: uint16(388),
	2709: uint16(1),
	2710: uint16(anon_sym_RPAREN),
	2711: uint16(516),
	2712: uint16(1),
	2713: uint16(sym__S),
	2714: uint16(158),
	2715: uint16(2),
	2716: uint16(sym_PEReference),
	2717: uint16(aux_sym__choice_repeat2),
	2718: uint16(5),
	2719: uint16(150),
	2720: uint16(1),
	2721: uint16(anon_sym_PERCENT),
	2722: uint16(430),
	2723: uint16(1),
	2724: uint16(anon_sym_RPAREN),
	2725: uint16(500),
	2726: uint16(1),
	2727: uint16(anon_sym_PIPE),
	2728: uint16(192),
	2729: uint16(1),
	2730: uint16(aux_sym_Mixed_repeat2),
	2731: uint16(246),
	2732: uint16(1),
	2733: uint16(sym_PEReference),
	2734: uint16(4),
	2735: uint16(518),
	2736: uint16(1),
	2737: uint16(anon_sym_PIPE),
	2738: uint16(523),
	2739: uint16(1),
	2740: uint16(sym__S),
	2741: uint16(140),
	2742: uint16(1),
	2743: uint16(aux_sym_Mixed_repeat1),
	2744: uint16(521),
	2745: uint16(2),
	2746: uint16(anon_sym_RPAREN),
	2747: uint16(anon_sym_PERCENT),
	2748: uint16(4),
	2749: uint16(528),
	2750: uint16(1),
	2751: uint16(anon_sym_PIPE),
	2752: uint16(531),
	2753: uint16(1),
	2754: uint16(sym__S),
	2755: uint16(141),
	2756: uint16(1),
	2757: uint16(aux_sym_NotationType_repeat1),
	2758: uint16(526),
	2759: uint16(2),
	2760: uint16(anon_sym_PERCENT),
	2761: uint16(sym_Name),
	2762: uint16(2),
	2763: uint16(374),
	2764: uint16(2),
	2765: uint16(anon_sym_AMP),
	2766: uint16(anon_sym_AMP_POUND),
	2767: uint16(372),
	2768: uint16(3),
	2769: uint16(anon_sym_SQUOTE),
	2770: uint16(aux_sym_PseudoAttValue_token2),
	2771: uint16(anon_sym_AMP_POUNDx),
	2772: uint16(5),
	2773: uint16(150),
	2774: uint16(1),
	2775: uint16(anon_sym_PERCENT),
	2776: uint16(500),
	2777: uint16(1),
	2778: uint16(anon_sym_PIPE),
	2779: uint16(534),
	2780: uint16(1),
	2781: uint16(anon_sym_RPAREN),
	2782: uint16(193),
	2783: uint16(1),
	2784: uint16(aux_sym_Mixed_repeat2),
	2785: uint16(246),
	2786: uint16(1),
	2787: uint16(sym_PEReference),
	2788: uint16(2),
	2789: uint16(290),
	2790: uint16(2),
	2791: uint16(anon_sym_AMP),
	2792: uint16(anon_sym_AMP_POUND),
	2793: uint16(288),
	2794: uint16(3),
	2795: uint16(anon_sym_SQUOTE),
	2796: uint16(aux_sym_PseudoAttValue_token2),
	2797: uint16(anon_sym_AMP_POUNDx),
	2798: uint16(2),
	2799: uint16(294),
	2800: uint16(2),
	2801: uint16(anon_sym_AMP),
	2802: uint16(anon_sym_AMP_POUND),
	2803: uint16(292),
	2804: uint16(3),
	2805: uint16(anon_sym_SQUOTE),
	2806: uint16(aux_sym_PseudoAttValue_token2),
	2807: uint16(anon_sym_AMP_POUNDx),
	2808: uint16(2),
	2809: uint16(538),
	2810: uint16(1),
	2811: uint16(anon_sym_LT),
	2812: uint16(536),
	2813: uint16(4),
	2814: uint16(sym_Comment),
	2815: uint16(anon_sym_LT_QMARK),
	2816: uint16(anon_sym_LT_BANG),
	2817: uint16(sym__S),
	2818: uint16(5),
	2819: uint16(442),
	2820: uint16(1),
	2821: uint16(anon_sym_SYSTEM),
	2822: uint16(444),
	2823: uint16(1),
	2824: uint16(anon_sym_PUBLIC),
	2825: uint16(540),
	2826: uint16(1),
	2827: uint16(anon_sym_LBRACK),
	2828: uint16(542),
	2829: uint16(1),
	2830: uint16(anon_sym_GT),
	2831: uint16(274),
	2832: uint16(1),
	2833: uint16(sym_ExternalID),
	2834: uint16(2),
	2835: uint16(374),
	2836: uint16(2),
	2837: uint16(anon_sym_AMP),
	2838: uint16(anon_sym_AMP_POUND),
	2839: uint16(372),
	2840: uint16(3),
	2841: uint16(anon_sym_DQUOTE),
	2842: uint16(aux_sym_PseudoAttValue_token1),
	2843: uint16(anon_sym_AMP_POUNDx),
	2844: uint16(2),
	2845: uint16(290),
	2846: uint16(2),
	2847: uint16(anon_sym_AMP),
	2848: uint16(anon_sym_AMP_POUND),
	2849: uint16(288),
	2850: uint16(3),
	2851: uint16(anon_sym_DQUOTE),
	2852: uint16(aux_sym_PseudoAttValue_token1),
	2853: uint16(anon_sym_AMP_POUNDx),
	2854: uint16(2),
	2855: uint16(294),
	2856: uint16(2),
	2857: uint16(anon_sym_AMP),
	2858: uint16(anon_sym_AMP_POUND),
	2859: uint16(292),
	2860: uint16(3),
	2861: uint16(anon_sym_DQUOTE),
	2862: uint16(aux_sym_PseudoAttValue_token1),
	2863: uint16(anon_sym_AMP_POUNDx),
	2864: uint16(4),
	2865: uint16(150),
	2866: uint16(1),
	2867: uint16(anon_sym_PERCENT),
	2868: uint16(396),
	2869: uint16(1),
	2870: uint16(anon_sym_RPAREN),
	2871: uint16(230),
	2872: uint16(1),
	2873: uint16(sym_PEReference),
	2874: uint16(504),
	2875: uint16(2),
	2876: uint16(anon_sym_PIPE),
	2877: uint16(anon_sym_COMMA),
	2878: uint16(1),
	2879: uint16(84),
	2880: uint16(5),
	2881: uint16(sym_Comment),
	2882: uint16(anon_sym_LT_QMARK),
	2883: uint16(anon_sym_LT_BANG),
	2884: uint16(anon_sym_RBRACK),
	2885: uint16(sym__S),
	2886: uint16(4),
	2887: uint16(150),
	2888: uint16(1),
	2889: uint16(anon_sym_PERCENT),
	2890: uint16(396),
	2891: uint16(1),
	2892: uint16(anon_sym_RPAREN),
	2893: uint16(544),
	2894: uint16(1),
	2895: uint16(sym__S),
	2896: uint16(158),
	2897: uint16(2),
	2898: uint16(sym_PEReference),
	2899: uint16(aux_sym__choice_repeat2),
	2900: uint16(1),
	2901: uint16(546),
	2902: uint16(5),
	2903: uint16(sym_Comment),
	2904: uint16(anon_sym_LT_QMARK),
	2905: uint16(anon_sym_LT_BANG),
	2906: uint16(anon_sym_RBRACK),
	2907: uint16(sym__S),
	2908: uint16(1),
	2909: uint16(459),
	2910: uint16(5),
	2911: uint16(anon_sym_PIPE),
	2912: uint16(anon_sym_RPAREN),
	2913: uint16(anon_sym_COMMA),
	2914: uint16(anon_sym_PERCENT),
	2915: uint16(sym__S),
	2916: uint16(5),
	2917: uint16(150),
	2918: uint16(1),
	2919: uint16(anon_sym_PERCENT),
	2920: uint16(424),
	2921: uint16(1),
	2922: uint16(anon_sym_RPAREN),
	2923: uint16(500),
	2924: uint16(1),
	2925: uint16(anon_sym_PIPE),
	2926: uint16(181),
	2927: uint16(1),
	2928: uint16(aux_sym_Mixed_repeat2),
	2929: uint16(246),
	2930: uint16(1),
	2931: uint16(sym_PEReference),
	2932: uint16(2),
	2933: uint16(548),
	2934: uint16(2),
	2935: uint16(anon_sym_GT),
	2936: uint16(sym__S),
	2937: uint16(550),
	2938: uint16(3),
	2939: uint16(anon_sym_STAR),
	2940: uint16(anon_sym_QMARK),
	2941: uint16(anon_sym_PLUS),
	2942: uint16(4),
	2943: uint16(552),
	2944: uint16(1),
	2945: uint16(anon_sym_RPAREN),
	2946: uint16(554),
	2947: uint16(1),
	2948: uint16(anon_sym_PERCENT),
	2949: uint16(557),
	2950: uint16(1),
	2951: uint16(sym__S),
	2952: uint16(158),
	2953: uint16(2),
	2954: uint16(sym_PEReference),
	2955: uint16(aux_sym__choice_repeat2),
	2956: uint16(4),
	2957: uint16(560),
	2958: uint16(1),
	2959: uint16(anon_sym_xml),
	2960: uint16(562),
	2961: uint16(1),
	2962: uint16(anon_sym_xml_DASHstylesheet),
	2963: uint16(564),
	2964: uint16(1),
	2965: uint16(anon_sym_xml_DASHmodel),
	2966: uint16(566),
	2967: uint16(1),
	2968: uint16(sym_PITarget),
	2969: uint16(4),
	2970: uint16(150),
	2971: uint16(1),
	2972: uint16(anon_sym_PERCENT),
	2973: uint16(568),
	2974: uint16(1),
	2975: uint16(sym_Name),
	2976: uint16(570),
	2977: uint16(1),
	2978: uint16(anon_sym_GT),
	2979: uint16(289),
	2980: uint16(1),
	2981: uint16(sym_PEReference),
	2982: uint16(3),
	2983: uint16(570),
	2984: uint16(1),
	2985: uint16(anon_sym_GT),
	2986: uint16(572),
	2987: uint16(1),
	2988: uint16(sym__S),
	2989: uint16(165),
	2990: uint16(2),
	2991: uint16(sym_AttDef),
	2992: uint16(aux_sym_AttlistDecl_repeat1),
	2993: uint16(4),
	2994: uint16(574),
	2995: uint16(1),
	2996: uint16(anon_sym_ELEMENT),
	2997: uint16(576),
	2998: uint16(1),
	2999: uint16(anon_sym_ATTLIST),
	3000: uint16(578),
	3001: uint16(1),
	3002: uint16(anon_sym_NOTATION),
	3003: uint16(580),
	3004: uint16(1),
	3005: uint16(anon_sym_ENTITY),
	3006: uint16(2),
	3007: uint16(584),
	3008: uint16(1),
	3009: uint16(anon_sym_LT),
	3010: uint16(582),
	3011: uint16(3),
	3012: uint16(sym_Comment),
	3013: uint16(anon_sym_LT_QMARK),
	3014: uint16(sym__S),
	3015: uint16(4),
	3016: uint16(150),
	3017: uint16(1),
	3018: uint16(anon_sym_PERCENT),
	3019: uint16(568),
	3020: uint16(1),
	3021: uint16(sym_Name),
	3022: uint16(586),
	3023: uint16(1),
	3024: uint16(anon_sym_GT),
	3025: uint16(289),
	3026: uint16(1),
	3027: uint16(sym_PEReference),
	3028: uint16(3),
	3029: uint16(588),
	3030: uint16(1),
	3031: uint16(anon_sym_GT),
	3032: uint16(590),
	3033: uint16(1),
	3034: uint16(sym__S),
	3035: uint16(165),
	3036: uint16(2),
	3037: uint16(sym_AttDef),
	3038: uint16(aux_sym_AttlistDecl_repeat1),
	3039: uint16(2),
	3040: uint16(595),
	3041: uint16(1),
	3042: uint16(anon_sym_LT),
	3043: uint16(593),
	3044: uint16(3),
	3045: uint16(sym_Comment),
	3046: uint16(anon_sym_LT_QMARK),
	3047: uint16(sym__S),
	3048: uint16(3),
	3049: uint16(597),
	3050: uint16(1),
	3051: uint16(anon_sym_PERCENT),
	3052: uint16(599),
	3053: uint16(1),
	3054: uint16(sym__S),
	3055: uint16(152),
	3056: uint16(2),
	3057: uint16(sym__DeclSep),
	3058: uint16(sym_PEReference),
	3059: uint16(1),
	3060: uint16(352),
	3061: uint16(4),
	3062: uint16(sym_Comment),
	3064: uint16(anon_sym_LT_QMARK),
	3065: uint16(sym__S),
	3066: uint16(4),
	3067: uint16(601),
	3068: uint16(1),
	3069: uint16(sym_Name),
	3070: uint16(603),
	3071: uint16(1),
	3072: uint16(anon_sym_GT),
	3073: uint16(605),
	3074: uint16(1),
	3075: uint16(anon_sym_SLASH_GT),
	3076: uint16(255),
	3077: uint16(1),
	3078: uint16(sym_Attribute),
	3079: uint16(2),
	3080: uint16(609),
	3081: uint16(1),
	3082: uint16(anon_sym_LT),
	3083: uint16(607),
	3084: uint16(3),
	3085: uint16(sym_Comment),
	3086: uint16(anon_sym_LT_QMARK),
	3087: uint16(sym__S),
	3088: uint16(2),
	3089: uint16(613),
	3090: uint16(1),
	3091: uint16(anon_sym_LT),
	3092: uint16(611),
	3093: uint16(3),
	3094: uint16(sym_Comment),
	3095: uint16(anon_sym_LT_QMARK),
	3096: uint16(sym__S),
	3097: uint16(4),
	3098: uint16(150),
	3099: uint16(1),
	3100: uint16(anon_sym_PERCENT),
	3101: uint16(615),
	3102: uint16(1),
	3103: uint16(sym_Name),
	3104: uint16(617),
	3105: uint16(1),
	3106: uint16(sym__S),
	3107: uint16(179),
	3108: uint16(1),
	3109: uint16(sym_PEReference),
	3110: uint16(4),
	3111: uint16(619),
	3112: uint16(1),
	3113: uint16(anon_sym_LT),
	3114: uint16(3),
	3115: uint16(1),
	3116: uint16(sym_STag),
	3117: uint16(39),
	3118: uint16(1),
	3119: uint16(sym_element),
	3120: uint16(189),
	3121: uint16(1),
	3122: uint16(sym_EmptyElemTag),
	3123: uint16(4),
	3124: uint16(150),
	3125: uint16(1),
	3126: uint16(anon_sym_PERCENT),
	3127: uint16(424),
	3128: uint16(1),
	3129: uint16(anon_sym_RPAREN),
	3130: uint16(183),
	3131: uint16(1),
	3132: uint16(aux_sym_Mixed_repeat2),
	3133: uint16(246),
	3134: uint16(1),
	3135: uint16(sym_PEReference),
	3136: uint16(1),
	3137: uint16(346),
	3138: uint16(4),
	3139: uint16(sym_Comment),
	3141: uint16(anon_sym_LT_QMARK),
	3142: uint16(sym__S),
	3143: uint16(3),
	3144: uint16(623),
	3145: uint16(1),
	3146: uint16(sym__S),
	3147: uint16(176),
	3148: uint16(1),
	3149: uint16(aux_sym_EmptyElemTag_repeat1),
	3150: uint16(621),
	3151: uint16(2),
	3152: uint16(anon_sym_GT),
	3153: uint16(anon_sym_SLASH_GT),
	3154: uint16(1),
	3155: uint16(356),
	3156: uint16(4),
	3157: uint16(sym_Comment),
	3159: uint16(anon_sym_LT_QMARK),
	3160: uint16(sym__S),
	3161: uint16(4),
	3162: uint16(626),
	3163: uint16(1),
	3164: uint16(anon_sym_GT),
	3165: uint16(628),
	3166: uint16(1),
	3167: uint16(anon_sym_SLASH_GT),
	3168: uint16(630),
	3169: uint16(1),
	3170: uint16(sym__S),
	3171: uint16(185),
	3172: uint16(1),
	3173: uint16(aux_sym_EmptyElemTag_repeat1),
	3174: uint16(1),
	3175: uint16(521),
	3176: uint16(4),
	3177: uint16(anon_sym_PIPE),
	3178: uint16(anon_sym_RPAREN),
	3179: uint16(anon_sym_PERCENT),
	3180: uint16(sym__S),
	3181: uint16(4),
	3182: uint16(150),
	3183: uint16(1),
	3184: uint16(anon_sym_PERCENT),
	3185: uint16(632),
	3186: uint16(1),
	3187: uint16(sym_Name),
	3188: uint16(634),
	3189: uint16(1),
	3190: uint16(sym__S),
	3191: uint16(191),
	3192: uint16(1),
	3193: uint16(sym_PEReference),
	3194: uint16(4),
	3195: uint16(150),
	3196: uint16(1),
	3197: uint16(anon_sym_PERCENT),
	3198: uint16(534),
	3199: uint16(1),
	3200: uint16(anon_sym_RPAREN),
	3201: uint16(183),
	3202: uint16(1),
	3203: uint16(aux_sym_Mixed_repeat2),
	3204: uint16(246),
	3205: uint16(1),
	3206: uint16(sym_PEReference),
	3207: uint16(4),
	3208: uint16(150),
	3209: uint16(1),
	3210: uint16(anon_sym_PERCENT),
	3211: uint16(430),
	3212: uint16(1),
	3213: uint16(anon_sym_RPAREN),
	3214: uint16(183),
	3215: uint16(1),
	3216: uint16(aux_sym_Mixed_repeat2),
	3217: uint16(246),
	3218: uint16(1),
	3219: uint16(sym_PEReference),
	3220: uint16(4),
	3221: uint16(636),
	3222: uint16(1),
	3223: uint16(anon_sym_RPAREN),
	3224: uint16(638),
	3225: uint16(1),
	3226: uint16(anon_sym_PERCENT),
	3227: uint16(183),
	3228: uint16(1),
	3229: uint16(aux_sym_Mixed_repeat2),
	3230: uint16(246),
	3231: uint16(1),
	3232: uint16(sym_PEReference),
	3233: uint16(4),
	3234: uint16(601),
	3235: uint16(1),
	3236: uint16(sym_Name),
	3237: uint16(641),
	3238: uint16(1),
	3239: uint16(anon_sym_GT),
	3240: uint16(643),
	3241: uint16(1),
	3242: uint16(anon_sym_SLASH_GT),
	3243: uint16(255),
	3244: uint16(1),
	3245: uint16(sym_Attribute),
	3246: uint16(4),
	3247: uint16(641),
	3248: uint16(1),
	3249: uint16(anon_sym_GT),
	3250: uint16(643),
	3251: uint16(1),
	3252: uint16(anon_sym_SLASH_GT),
	3253: uint16(645),
	3254: uint16(1),
	3255: uint16(sym__S),
	3256: uint16(176),
	3257: uint16(1),
	3258: uint16(aux_sym_EmptyElemTag_repeat1),
	3259: uint16(4),
	3260: uint16(619),
	3261: uint16(1),
	3262: uint16(anon_sym_LT),
	3263: uint16(3),
	3264: uint16(1),
	3265: uint16(sym_STag),
	3266: uint16(41),
	3267: uint16(1),
	3268: uint16(sym_element),
	3269: uint16(189),
	3270: uint16(1),
	3271: uint16(sym_EmptyElemTag),
	3272: uint16(1),
	3273: uint16(368),
	3274: uint16(4),
	3275: uint16(sym_Comment),
	3277: uint16(anon_sym_LT_QMARK),
	3278: uint16(sym__S),
	3279: uint16(3),
	3280: uint16(597),
	3281: uint16(1),
	3282: uint16(anon_sym_PERCENT),
	3283: uint16(647),
	3284: uint16(1),
	3285: uint16(sym__S),
	3286: uint16(154),
	3287: uint16(2),
	3288: uint16(sym__DeclSep),
	3289: uint16(sym_PEReference),
	3290: uint16(1),
	3291: uint16(334),
	3292: uint16(4),
	3293: uint16(sym_Comment),
	3295: uint16(anon_sym_LT_QMARK),
	3296: uint16(sym__S),
	3297: uint16(4),
	3298: uint16(649),
	3299: uint16(1),
	3300: uint16(anon_sym_PIPE),
	3301: uint16(651),
	3302: uint16(1),
	3303: uint16(anon_sym_RPAREN),
	3304: uint16(653),
	3305: uint16(1),
	3306: uint16(sym__S),
	3307: uint16(196),
	3308: uint16(1),
	3309: uint16(aux_sym_Enumeration_repeat1),
	3310: uint16(1),
	3311: uint16(655),
	3312: uint16(4),
	3313: uint16(anon_sym_PIPE),
	3314: uint16(anon_sym_RPAREN),
	3315: uint16(anon_sym_PERCENT),
	3316: uint16(sym__S),
	3317: uint16(4),
	3318: uint16(150),
	3319: uint16(1),
	3320: uint16(anon_sym_PERCENT),
	3321: uint16(502),
	3322: uint16(1),
	3323: uint16(anon_sym_RPAREN),
	3324: uint16(183),
	3325: uint16(1),
	3326: uint16(aux_sym_Mixed_repeat2),
	3327: uint16(246),
	3328: uint16(1),
	3329: uint16(sym_PEReference),
	3330: uint16(4),
	3331: uint16(150),
	3332: uint16(1),
	3333: uint16(anon_sym_PERCENT),
	3334: uint16(657),
	3335: uint16(1),
	3336: uint16(anon_sym_RPAREN),
	3337: uint16(183),
	3338: uint16(1),
	3339: uint16(aux_sym_Mixed_repeat2),
	3340: uint16(246),
	3341: uint16(1),
	3342: uint16(sym_PEReference),
	3343: uint16(1),
	3344: uint16(196),
	3345: uint16(4),
	3346: uint16(sym_Comment),
	3348: uint16(anon_sym_LT_QMARK),
	3349: uint16(sym__S),
	3350: uint16(4),
	3351: uint16(649),
	3352: uint16(1),
	3353: uint16(anon_sym_PIPE),
	3354: uint16(659),
	3355: uint16(1),
	3356: uint16(anon_sym_RPAREN),
	3357: uint16(661),
	3358: uint16(1),
	3359: uint16(sym__S),
	3360: uint16(200),
	3361: uint16(1),
	3362: uint16(aux_sym_Enumeration_repeat1),
	3363: uint16(4),
	3364: uint16(649),
	3365: uint16(1),
	3366: uint16(anon_sym_PIPE),
	3367: uint16(659),
	3368: uint16(1),
	3369: uint16(anon_sym_RPAREN),
	3370: uint16(661),
	3371: uint16(1),
	3372: uint16(sym__S),
	3373: uint16(201),
	3374: uint16(1),
	3375: uint16(aux_sym_Enumeration_repeat1),
	3376: uint16(4),
	3377: uint16(150),
	3378: uint16(1),
	3379: uint16(anon_sym_PERCENT),
	3380: uint16(663),
	3381: uint16(1),
	3382: uint16(sym_Name),
	3383: uint16(665),
	3384: uint16(1),
	3385: uint16(sym__S),
	3386: uint16(114),
	3387: uint16(1),
	3388: uint16(sym_PEReference),
	3389: uint16(1),
	3390: uint16(667),
	3391: uint16(4),
	3392: uint16(anon_sym_PIPE),
	3393: uint16(anon_sym_RPAREN),
	3394: uint16(anon_sym_PERCENT),
	3395: uint16(sym__S),
	3396: uint16(4),
	3397: uint16(150),
	3398: uint16(1),
	3399: uint16(anon_sym_PERCENT),
	3400: uint16(669),
	3401: uint16(1),
	3402: uint16(anon_sym_RPAREN),
	3403: uint16(183),
	3404: uint16(1),
	3405: uint16(aux_sym_Mixed_repeat2),
	3406: uint16(246),
	3407: uint16(1),
	3408: uint16(sym_PEReference),
	3409: uint16(4),
	3410: uint16(649),
	3411: uint16(1),
	3412: uint16(anon_sym_PIPE),
	3413: uint16(671),
	3414: uint16(1),
	3415: uint16(anon_sym_RPAREN),
	3416: uint16(673),
	3417: uint16(1),
	3418: uint16(sym__S),
	3419: uint16(201),
	3420: uint16(1),
	3421: uint16(aux_sym_Enumeration_repeat1),
	3422: uint16(4),
	3423: uint16(675),
	3424: uint16(1),
	3425: uint16(anon_sym_PIPE),
	3426: uint16(678),
	3427: uint16(1),
	3428: uint16(anon_sym_RPAREN),
	3429: uint16(680),
	3430: uint16(1),
	3431: uint16(sym__S),
	3432: uint16(201),
	3433: uint16(1),
	3434: uint16(aux_sym_Enumeration_repeat1),
	3435: uint16(1),
	3436: uint16(338),
	3437: uint16(4),
	3438: uint16(sym_Comment),
	3440: uint16(anon_sym_LT_QMARK),
	3441: uint16(sym__S),
	3442: uint16(2),
	3443: uint16(685),
	3444: uint16(1),
	3445: uint16(sym__S),
	3446: uint16(683),
	3447: uint16(3),
	3448: uint16(anon_sym_PIPE),
	3449: uint16(anon_sym_PERCENT),
	3450: uint16(sym_Name),
	3451: uint16(1),
	3452: uint16(526),
	3453: uint16(4),
	3454: uint16(anon_sym_PIPE),
	3455: uint16(anon_sym_PERCENT),
	3456: uint16(sym__S),
	3457: uint16(sym_Name),
	3458: uint16(2),
	3459: uint16(688),
	3460: uint16(1),
	3461: uint16(sym__S),
	3462: uint16(526),
	3463: uint16(3),
	3464: uint16(anon_sym_PIPE),
	3465: uint16(anon_sym_PERCENT),
	3466: uint16(sym_Name),
	3467: uint16(1),
	3468: uint16(691),
	3469: uint16(4),
	3470: uint16(anon_sym_PIPE),
	3471: uint16(anon_sym_PERCENT),
	3472: uint16(sym__S),
	3473: uint16(sym_Name),
	3474: uint16(2),
	3475: uint16(695),
	3476: uint16(1),
	3477: uint16(anon_sym_LT),
	3478: uint16(693),
	3479: uint16(3),
	3480: uint16(sym_Comment),
	3481: uint16(anon_sym_LT_QMARK),
	3482: uint16(sym__S),
	3483: uint16(4),
	3484: uint16(697),
	3485: uint16(1),
	3486: uint16(anon_sym_QMARK_GT),
	3487: uint16(699),
	3488: uint16(1),
	3489: uint16(sym__S),
	3490: uint16(251),
	3491: uint16(1),
	3492: uint16(sym__EncodingDecl),
	3493: uint16(353),
	3494: uint16(1),
	3495: uint16(sym__SDDecl),
	3496: uint16(2),
	3497: uint16(703),
	3498: uint16(1),
	3499: uint16(anon_sym_LT),
	3500: uint16(701),
	3501: uint16(3),
	3502: uint16(sym_Comment),
	3503: uint16(anon_sym_LT_QMARK),
	3504: uint16(sym__S),
	3505: uint16(3),
	3506: uint16(705),
	3507: uint16(1),
	3508: uint16(anon_sym_GT),
	3509: uint16(707),
	3510: uint16(1),
	3511: uint16(sym__S),
	3512: uint16(161),
	3513: uint16(2),
	3514: uint16(sym_AttDef),
	3515: uint16(aux_sym_AttlistDecl_repeat1),
	3516: uint16(2),
	3517: uint16(711),
	3518: uint16(1),
	3519: uint16(anon_sym_LT),
	3520: uint16(709),
	3521: uint16(3),
	3522: uint16(sym_Comment),
	3523: uint16(anon_sym_LT_QMARK),
	3524: uint16(sym__S),
	3525: uint16(4),
	3526: uint16(626),
	3527: uint16(1),
	3528: uint16(anon_sym_GT),
	3529: uint16(713),
	3530: uint16(1),
	3531: uint16(anon_sym_SLASH_GT),
	3532: uint16(715),
	3533: uint16(1),
	3534: uint16(sym__S),
	3535: uint16(214),
	3536: uint16(1),
	3537: uint16(aux_sym_EmptyElemTag_repeat1),
	3538: uint16(4),
	3539: uint16(601),
	3540: uint16(1),
	3541: uint16(sym_Name),
	3542: uint16(641),
	3543: uint16(1),
	3544: uint16(anon_sym_GT),
	3545: uint16(717),
	3546: uint16(1),
	3547: uint16(anon_sym_SLASH_GT),
	3548: uint16(255),
	3549: uint16(1),
	3550: uint16(sym_Attribute),
	3551: uint16(4),
	3552: uint16(641),
	3553: uint16(1),
	3554: uint16(anon_sym_GT),
	3555: uint16(717),
	3556: uint16(1),
	3557: uint16(anon_sym_SLASH_GT),
	3558: uint16(719),
	3559: uint16(1),
	3560: uint16(sym__S),
	3561: uint16(176),
	3562: uint16(1),
	3563: uint16(aux_sym_EmptyElemTag_repeat1),
	3564: uint16(4),
	3565: uint16(601),
	3566: uint16(1),
	3567: uint16(sym_Name),
	3568: uint16(603),
	3569: uint16(1),
	3570: uint16(anon_sym_GT),
	3571: uint16(721),
	3572: uint16(1),
	3573: uint16(anon_sym_SLASH_GT),
	3574: uint16(255),
	3575: uint16(1),
	3576: uint16(sym_Attribute),
	3577: uint16(1),
	3578: uint16(364),
	3579: uint16(4),
	3580: uint16(sym_Comment),
	3582: uint16(anon_sym_LT_QMARK),
	3583: uint16(sym__S),
	3584: uint16(2),
	3585: uint16(725),
	3586: uint16(1),
	3587: uint16(anon_sym_LT),
	3588: uint16(723),
	3589: uint16(3),
	3590: uint16(sym_Comment),
	3591: uint16(anon_sym_LT_QMARK),
	3592: uint16(sym__S),
	3593: uint16(1),
	3594: uint16(727),
	3595: uint16(3),
	3596: uint16(anon_sym_LBRACK),
	3597: uint16(anon_sym_GT),
	3598: uint16(sym__S),
	3599: uint16(1),
	3600: uint16(729),
	3601: uint16(3),
	3602: uint16(anon_sym_GT),
	3603: uint16(anon_sym_SLASH_GT),
	3604: uint16(sym__S),
	3605: uint16(3),
	3606: uint16(731),
	3607: uint16(1),
	3608: uint16(anon_sym_QMARK_GT),
	3609: uint16(733),
	3610: uint16(1),
	3611: uint16(sym__S),
	3612: uint16(248),
	3613: uint16(1),
	3614: uint16(aux_sym_StyleSheetPI_repeat1),
	3615: uint16(3),
	3616: uint16(735),
	3617: uint16(1),
	3618: uint16(sym__S),
	3619: uint16(737),
	3620: uint16(1),
	3621: uint16(anon_sym_EQ),
	3622: uint16(263),
	3623: uint16(1),
	3624: uint16(sym__Eq),
	3625: uint16(3),
	3626: uint16(150),
	3627: uint16(1),
	3628: uint16(anon_sym_PERCENT),
	3629: uint16(632),
	3630: uint16(1),
	3631: uint16(sym_Name),
	3632: uint16(191),
	3633: uint16(1),
	3634: uint16(sym_PEReference),
	3635: uint16(2),
	3636: uint16(741),
	3637: uint16(1),
	3638: uint16(anon_sym_STAR),
	3639: uint16(739),
	3640: uint16(2),
	3641: uint16(anon_sym_GT),
	3642: uint16(sym__S),
	3643: uint16(3),
	3644: uint16(743),
	3645: uint16(1),
	3646: uint16(sym_Name),
	3647: uint16(745),
	3648: uint16(1),
	3649: uint16(anon_sym_QMARK_GT),
	3650: uint16(368),
	3651: uint16(1),
	3652: uint16(sym_PseudoAtt),
	3653: uint16(3),
	3654: uint16(747),
	3655: uint16(1),
	3656: uint16(anon_sym_QMARK_GT),
	3657: uint16(749),
	3658: uint16(1),
	3659: uint16(sym__S),
	3660: uint16(225),
	3661: uint16(1),
	3662: uint16(aux_sym_StyleSheetPI_repeat1),
	3663: uint16(3),
	3664: uint16(752),
	3665: uint16(1),
	3666: uint16(anon_sym_GT),
	3667: uint16(754),
	3668: uint16(1),
	3669: uint16(sym__S),
	3670: uint16(342),
	3671: uint16(1),
	3672: uint16(sym_NDataDecl),
	3673: uint16(3),
	3674: uint16(743),
	3675: uint16(1),
	3676: uint16(sym_Name),
	3677: uint16(756),
	3678: uint16(1),
	3679: uint16(anon_sym_QMARK_GT),
	3680: uint16(368),
	3681: uint16(1),
	3682: uint16(sym_PseudoAtt),
	3683: uint16(3),
	3684: uint16(562),
	3685: uint16(1),
	3686: uint16(anon_sym_xml_DASHstylesheet),
	3687: uint16(564),
	3688: uint16(1),
	3689: uint16(anon_sym_xml_DASHmodel),
	3690: uint16(566),
	3691: uint16(1),
	3692: uint16(sym_PITarget),
	3693: uint16(3),
	3694: uint16(743),
	3695: uint16(1),
	3696: uint16(sym_Name),
	3697: uint16(758),
	3698: uint16(1),
	3699: uint16(anon_sym_QMARK_GT),
	3700: uint16(368),
	3701: uint16(1),
	3702: uint16(sym_PseudoAtt),
	3703: uint16(1),
	3704: uint16(552),
	3705: uint16(3),
	3706: uint16(anon_sym_RPAREN),
	3707: uint16(anon_sym_PERCENT),
	3708: uint16(sym__S),
	3709: uint16(3),
	3710: uint16(150),
	3711: uint16(1),
	3712: uint16(anon_sym_PERCENT),
	3713: uint16(396),
	3714: uint16(1),
	3715: uint16(anon_sym_RPAREN),
	3716: uint16(230),
	3717: uint16(1),
	3718: uint16(sym_PEReference),
	3719: uint16(3),
	3720: uint16(760),
	3721: uint16(1),
	3722: uint16(anon_sym_LBRACK),
	3723: uint16(762),
	3724: uint16(1),
	3725: uint16(anon_sym_GT),
	3726: uint16(764),
	3727: uint16(1),
	3728: uint16(sym__S),
	3729: uint16(3),
	3730: uint16(766),
	3731: uint16(1),
	3732: uint16(anon_sym_SQUOTE),
	3733: uint16(768),
	3734: uint16(1),
	3735: uint16(anon_sym_DQUOTE),
	3736: uint16(267),
	3737: uint16(1),
	3738: uint16(sym_SystemLiteral),
	3739: uint16(3),
	3740: uint16(150),
	3741: uint16(1),
	3742: uint16(anon_sym_PERCENT),
	3743: uint16(770),
	3744: uint16(1),
	3745: uint16(sym_Name),
	3746: uint16(293),
	3747: uint16(1),
	3748: uint16(sym_PEReference),
	3749: uint16(3),
	3750: uint16(150),
	3751: uint16(1),
	3752: uint16(anon_sym_PERCENT),
	3753: uint16(772),
	3754: uint16(1),
	3755: uint16(sym_Name),
	3756: uint16(198),
	3757: uint16(1),
	3758: uint16(sym_PEReference),
	3759: uint16(2),
	3760: uint16(776),
	3761: uint16(1),
	3762: uint16(anon_sym_STAR),
	3763: uint16(774),
	3764: uint16(2),
	3765: uint16(anon_sym_GT),
	3766: uint16(sym__S),
	3767: uint16(3),
	3768: uint16(778),
	3769: uint16(1),
	3770: uint16(anon_sym_SQUOTE),
	3771: uint16(780),
	3772: uint16(1),
	3773: uint16(anon_sym_DQUOTE),
	3774: uint16(406),
	3775: uint16(1),
	3776: uint16(sym_PubidLiteral),
	3777: uint16(2),
	3778: uint16(784),
	3779: uint16(1),
	3780: uint16(sym__S),
	3781: uint16(782),
	3782: uint16(2),
	3783: uint16(anon_sym_SQUOTE),
	3784: uint16(anon_sym_DQUOTE),
	3785: uint16(3),
	3786: uint16(150),
	3787: uint16(1),
	3788: uint16(anon_sym_PERCENT),
	3789: uint16(488),
	3790: uint16(1),
	3791: uint16(anon_sym_RPAREN),
	3792: uint16(230),
	3793: uint16(1),
	3794: uint16(sym_PEReference),
	3795: uint16(1),
	3796: uint16(786),
	3797: uint16(3),
	3798: uint16(anon_sym_GT),
	3799: uint16(anon_sym_SLASH_GT),
	3800: uint16(sym__S),
	3801: uint16(1),
	3802: uint16(788),
	3803: uint16(3),
	3804: uint16(anon_sym_LBRACK),
	3805: uint16(anon_sym_GT),
	3806: uint16(sym__S),
	3807: uint16(3),
	3808: uint16(150),
	3809: uint16(1),
	3810: uint16(anon_sym_PERCENT),
	3811: uint16(568),
	3812: uint16(1),
	3813: uint16(sym_Name),
	3814: uint16(289),
	3815: uint16(1),
	3816: uint16(sym_PEReference),
	3817: uint16(3),
	3818: uint16(778),
	3819: uint16(1),
	3820: uint16(anon_sym_SQUOTE),
	3821: uint16(780),
	3822: uint16(1),
	3823: uint16(anon_sym_DQUOTE),
	3824: uint16(335),
	3825: uint16(1),
	3826: uint16(sym_PubidLiteral),
	3827: uint16(3),
	3828: uint16(778),
	3829: uint16(1),
	3830: uint16(anon_sym_SQUOTE),
	3831: uint16(780),
	3832: uint16(1),
	3833: uint16(anon_sym_DQUOTE),
	3834: uint16(339),
	3835: uint16(1),
	3836: uint16(sym_PubidLiteral),
	3837: uint16(2),
	3838: uint16(792),
	3839: uint16(1),
	3840: uint16(anon_sym_STAR),
	3841: uint16(790),
	3842: uint16(2),
	3843: uint16(anon_sym_GT),
	3844: uint16(sym__S),
	3845: uint16(2),
	3846: uint16(796),
	3847: uint16(1),
	3848: uint16(sym__S),
	3849: uint16(794),
	3850: uint16(2),
	3851: uint16(anon_sym_RPAREN),
	3852: uint16(anon_sym_PERCENT),
	3853: uint16(3),
	3854: uint16(798),
	3855: uint16(1),
	3856: uint16(anon_sym_QMARK_GT),
	3857: uint16(800),
	3858: uint16(1),
	3859: uint16(sym__S),
	3860: uint16(268),
	3861: uint16(1),
	3862: uint16(aux_sym_StyleSheetPI_repeat1),
	3863: uint16(3),
	3864: uint16(756),
	3865: uint16(1),
	3866: uint16(anon_sym_QMARK_GT),
	3867: uint16(802),
	3868: uint16(1),
	3869: uint16(sym__S),
	3870: uint16(225),
	3871: uint16(1),
	3872: uint16(aux_sym_StyleSheetPI_repeat1),
	3873: uint16(1),
	3874: uint16(678),
	3875: uint16(3),
	3876: uint16(anon_sym_PIPE),
	3877: uint16(anon_sym_RPAREN),
	3878: uint16(sym__S),
	3879: uint16(3),
	3880: uint16(150),
	3881: uint16(1),
	3882: uint16(anon_sym_PERCENT),
	3883: uint16(804),
	3884: uint16(1),
	3885: uint16(sym_Name),
	3886: uint16(115),
	3887: uint16(1),
	3888: uint16(sym_PEReference),
	3889: uint16(3),
	3890: uint16(806),
	3891: uint16(1),
	3892: uint16(anon_sym_QMARK_GT),
	3893: uint16(808),
	3894: uint16(1),
	3895: uint16(sym__S),
	3896: uint16(338),
	3897: uint16(1),
	3898: uint16(sym__SDDecl),
	3899: uint16(3),
	3900: uint16(320),
	3901: uint16(1),
	3902: uint16(anon_sym_SQUOTE),
	3903: uint16(322),
	3904: uint16(1),
	3905: uint16(anon_sym_DQUOTE),
	3906: uint16(319),
	3907: uint16(1),
	3908: uint16(sym_AttValue),
	3909: uint16(1),
	3910: uint16(810),
	3911: uint16(3),
	3912: uint16(anon_sym_PIPE),
	3913: uint16(anon_sym_RPAREN),
	3914: uint16(sym__S),
	3915: uint16(3),
	3916: uint16(735),
	3917: uint16(1),
	3918: uint16(sym__S),
	3919: uint16(737),
	3920: uint16(1),
	3921: uint16(anon_sym_EQ),
	3922: uint16(273),
	3923: uint16(1),
	3924: uint16(sym__Eq),
	3925: uint16(1),
	3926: uint16(621),
	3927: uint16(3),
	3928: uint16(anon_sym_GT),
	3929: uint16(anon_sym_SLASH_GT),
	3930: uint16(sym__S),
	3931: uint16(3),
	3932: uint16(735),
	3933: uint16(1),
	3934: uint16(sym__S),
	3935: uint16(737),
	3936: uint16(1),
	3937: uint16(anon_sym_EQ),
	3938: uint16(298),
	3939: uint16(1),
	3940: uint16(sym__Eq),
	3941: uint16(1),
	3942: uint16(812),
	3943: uint16(3),
	3944: uint16(anon_sym_PIPE),
	3945: uint16(anon_sym_RPAREN),
	3946: uint16(sym__S),
	3947: uint16(3),
	3948: uint16(150),
	3949: uint16(1),
	3950: uint16(anon_sym_PERCENT),
	3951: uint16(814),
	3952: uint16(1),
	3953: uint16(sym_Name),
	3954: uint16(419),
	3955: uint16(1),
	3956: uint16(sym_PEReference),
	3957: uint16(3),
	3958: uint16(150),
	3959: uint16(1),
	3960: uint16(anon_sym_PERCENT),
	3961: uint16(816),
	3962: uint16(1),
	3963: uint16(sym_Name),
	3964: uint16(210),
	3965: uint16(1),
	3966: uint16(sym_PEReference),
	3967: uint16(3),
	3968: uint16(735),
	3969: uint16(1),
	3970: uint16(sym__S),
	3971: uint16(737),
	3972: uint16(1),
	3973: uint16(anon_sym_EQ),
	3974: uint16(372),
	3975: uint16(1),
	3976: uint16(sym__Eq),
	3977: uint16(3),
	3978: uint16(735),
	3979: uint16(1),
	3980: uint16(sym__S),
	3981: uint16(737),
	3982: uint16(1),
	3983: uint16(anon_sym_EQ),
	3984: uint16(299),
	3985: uint16(1),
	3986: uint16(sym__Eq),
	3987: uint16(3),
	3988: uint16(150),
	3989: uint16(1),
	3990: uint16(anon_sym_PERCENT),
	3991: uint16(818),
	3992: uint16(1),
	3993: uint16(sym_Name),
	3994: uint16(422),
	3995: uint16(1),
	3996: uint16(sym_PEReference),
	3997: uint16(3),
	3998: uint16(820),
	3999: uint16(1),
	4000: uint16(anon_sym_SQUOTE),
	4001: uint16(822),
	4002: uint16(1),
	4003: uint16(anon_sym_DQUOTE),
	4004: uint16(361),
	4005: uint16(1),
	4006: uint16(sym_PseudoAttValue),
	4007: uint16(3),
	4008: uint16(824),
	4009: uint16(1),
	4010: uint16(sym_Name),
	4011: uint16(826),
	4012: uint16(1),
	4013: uint16(anon_sym_PERCENT),
	4014: uint16(424),
	4015: uint16(1),
	4016: uint16(sym_PEReference),
	4017: uint16(3),
	4018: uint16(743),
	4019: uint16(1),
	4020: uint16(sym_Name),
	4021: uint16(828),
	4022: uint16(1),
	4023: uint16(anon_sym_QMARK_GT),
	4024: uint16(368),
	4025: uint16(1),
	4026: uint16(sym_PseudoAtt),
	4027: uint16(2),
	4028: uint16(832),
	4029: uint16(1),
	4030: uint16(sym__S),
	4031: uint16(830),
	4032: uint16(2),
	4033: uint16(anon_sym_SQUOTE),
	4034: uint16(anon_sym_DQUOTE),
	4035: uint16(1),
	4036: uint16(834),
	4037: uint16(3),
	4038: uint16(anon_sym_LBRACK),
	4039: uint16(anon_sym_GT),
	4040: uint16(sym__S),
	4041: uint16(3),
	4042: uint16(828),
	4043: uint16(1),
	4044: uint16(anon_sym_QMARK_GT),
	4045: uint16(836),
	4046: uint16(1),
	4047: uint16(sym__S),
	4048: uint16(225),
	4049: uint16(1),
	4050: uint16(aux_sym_StyleSheetPI_repeat1),
	4051: uint16(3),
	4052: uint16(806),
	4053: uint16(1),
	4054: uint16(anon_sym_QMARK_GT),
	4055: uint16(838),
	4056: uint16(1),
	4057: uint16(anon_sym_standalone),
	4058: uint16(840),
	4059: uint16(1),
	4060: uint16(anon_sym_encoding),
	4061: uint16(1),
	4062: uint16(842),
	4063: uint16(3),
	4064: uint16(anon_sym_GT),
	4065: uint16(anon_sym_SLASH_GT),
	4066: uint16(sym__S),
	4067: uint16(2),
	4068: uint16(846),
	4069: uint16(1),
	4070: uint16(anon_sym_STAR),
	4071: uint16(844),
	4072: uint16(2),
	4073: uint16(anon_sym_GT),
	4074: uint16(sym__S),
	4075: uint16(3),
	4076: uint16(766),
	4077: uint16(1),
	4078: uint16(anon_sym_SQUOTE),
	4079: uint16(768),
	4080: uint16(1),
	4081: uint16(anon_sym_DQUOTE),
	4082: uint16(218),
	4083: uint16(1),
	4084: uint16(sym_SystemLiteral),
	4085: uint16(3),
	4086: uint16(848),
	4087: uint16(1),
	4088: uint16(anon_sym_SQUOTE),
	4089: uint16(850),
	4090: uint16(1),
	4091: uint16(anon_sym_DQUOTE),
	4092: uint16(219),
	4093: uint16(1),
	4094: uint16(sym_AttValue),
	4095: uint16(3),
	4096: uint16(852),
	4097: uint16(1),
	4098: uint16(anon_sym_LBRACK),
	4099: uint16(854),
	4100: uint16(1),
	4101: uint16(anon_sym_GT),
	4102: uint16(856),
	4103: uint16(1),
	4104: uint16(sym__S),
	4105: uint16(3),
	4106: uint16(150),
	4107: uint16(1),
	4108: uint16(anon_sym_PERCENT),
	4109: uint16(858),
	4110: uint16(1),
	4111: uint16(anon_sym_RPAREN),
	4112: uint16(230),
	4113: uint16(1),
	4114: uint16(sym_PEReference),
	4115: uint16(1),
	4116: uint16(860),
	4117: uint16(2),
	4118: uint16(anon_sym_QMARK_GT),
	4119: uint16(sym__S),
	4120: uint16(2),
	4121: uint16(862),
	4122: uint16(1),
	4123: uint16(anon_sym_GT),
	4124: uint16(864),
	4125: uint16(1),
	4126: uint16(sym__S),
	4127: uint16(2),
	4128: uint16(866),
	4129: uint16(1),
	4130: uint16(anon_sym_LBRACK),
	4131: uint16(868),
	4132: uint16(1),
	4133: uint16(anon_sym_GT),
	4134: uint16(1),
	4135: uint16(870),
	4136: uint16(2),
	4137: uint16(anon_sym_yes),
	4138: uint16(anon_sym_no),
	4139: uint16(1),
	4140: uint16(872),
	4141: uint16(2),
	4142: uint16(anon_sym_PERCENT),
	4143: uint16(sym__S),
	4144: uint16(1),
	4145: uint16(739),
	4146: uint16(2),
	4147: uint16(anon_sym_GT),
	4148: uint16(sym__S),
	4149: uint16(1),
	4150: uint16(504),
	4151: uint16(2),
	4152: uint16(anon_sym_PIPE),
	4153: uint16(anon_sym_COMMA),
	4154: uint16(2),
	4155: uint16(150),
	4156: uint16(1),
	4157: uint16(anon_sym_PERCENT),
	4158: uint16(230),
	4159: uint16(1),
	4160: uint16(sym_PEReference),
	4161: uint16(1),
	4162: uint16(874),
	4163: uint16(2),
	4164: uint16(anon_sym_PERCENT),
	4165: uint16(sym__S),
	4166: uint16(2),
	4167: uint16(876),
	4168: uint16(1),
	4169: uint16(sym__S),
	4170: uint16(878),
	4171: uint16(1),
	4172: uint16(sym_Nmtoken),
	4173: uint16(1),
	4174: uint16(880),
	4175: uint16(2),
	4176: uint16(anon_sym_GT),
	4177: uint16(sym__S),
	4178: uint16(2),
	4179: uint16(659),
	4180: uint16(1),
	4181: uint16(anon_sym_RPAREN),
	4182: uint16(882),
	4183: uint16(1),
	4184: uint16(anon_sym_PIPE),
	4185: uint16(2),
	4186: uint16(884),
	4187: uint16(1),
	4188: uint16(anon_sym_GT),
	4189: uint16(886),
	4190: uint16(1),
	4191: uint16(sym__S),
	4192: uint16(2),
	4193: uint16(588),
	4194: uint16(1),
	4195: uint16(anon_sym_GT),
	4196: uint16(888),
	4197: uint16(1),
	4198: uint16(sym__S),
	4199: uint16(1),
	4200: uint16(891),
	4201: uint16(2),
	4202: uint16(anon_sym_GT),
	4203: uint16(sym__S),
	4204: uint16(1),
	4205: uint16(893),
	4206: uint16(2),
	4207: uint16(anon_sym_GT),
	4208: uint16(sym__S),
	4209: uint16(1),
	4210: uint16(895),
	4211: uint16(2),
	4212: uint16(anon_sym_PERCENT),
	4213: uint16(sym__S),
	4214: uint16(1),
	4215: uint16(897),
	4216: uint16(2),
	4217: uint16(anon_sym_GT),
	4218: uint16(sym__S),
	4219: uint16(1),
	4220: uint16(899),
	4221: uint16(2),
	4222: uint16(anon_sym_yes),
	4223: uint16(anon_sym_no),
	4224: uint16(1),
	4225: uint16(790),
	4226: uint16(2),
	4227: uint16(anon_sym_GT),
	4228: uint16(sym__S),
	4229: uint16(1),
	4230: uint16(636),
	4231: uint16(2),
	4232: uint16(anon_sym_RPAREN),
	4233: uint16(anon_sym_PERCENT),
	4234: uint16(1),
	4235: uint16(901),
	4236: uint16(2),
	4237: uint16(anon_sym_PERCENT),
	4238: uint16(sym__S),
	4239: uint16(2),
	4240: uint16(903),
	4241: uint16(1),
	4242: uint16(anon_sym_SQUOTE),
	4243: uint16(905),
	4244: uint16(1),
	4245: uint16(anon_sym_DQUOTE),
	4246: uint16(2),
	4247: uint16(907),
	4248: uint16(1),
	4249: uint16(anon_sym_SQUOTE),
	4250: uint16(909),
	4251: uint16(1),
	4252: uint16(anon_sym_DQUOTE),
	4253: uint16(1),
	4254: uint16(911),
	4255: uint16(2),
	4256: uint16(anon_sym_QMARK_GT),
	4257: uint16(sym__S),
	4258: uint16(1),
	4259: uint16(913),
	4260: uint16(2),
	4261: uint16(anon_sym_GT),
	4262: uint16(sym__S),
	4263: uint16(2),
	4264: uint16(671),
	4265: uint16(1),
	4266: uint16(anon_sym_RPAREN),
	4267: uint16(882),
	4268: uint16(1),
	4269: uint16(anon_sym_PIPE),
	4270: uint16(1),
	4271: uint16(915),
	4272: uint16(2),
	4273: uint16(anon_sym_PERCENT),
	4274: uint16(sym__S),
	4275: uint16(2),
	4276: uint16(752),
	4277: uint16(1),
	4278: uint16(anon_sym_GT),
	4279: uint16(917),
	4280: uint16(1),
	4281: uint16(sym__S),
	4282: uint16(2),
	4283: uint16(919),
	4284: uint16(1),
	4285: uint16(sym__S),
	4286: uint16(921),
	4287: uint16(1),
	4288: uint16(sym_Nmtoken),
	4289: uint16(1),
	4290: uint16(923),
	4291: uint16(2),
	4292: uint16(anon_sym_GT),
	4293: uint16(sym__S),
	4294: uint16(1),
	4295: uint16(925),
	4296: uint16(2),
	4297: uint16(anon_sym_GT),
	4298: uint16(sym__S),
	4299: uint16(2),
	4300: uint16(854),
	4301: uint16(1),
	4302: uint16(anon_sym_GT),
	4303: uint16(927),
	4304: uint16(1),
	4305: uint16(sym__S),
	4306: uint16(2),
	4307: uint16(929),
	4308: uint16(1),
	4309: uint16(sym__S),
	4310: uint16(931),
	4311: uint16(1),
	4312: uint16(sym_Nmtoken),
	4313: uint16(1),
	4314: uint16(933),
	4315: uint16(2),
	4316: uint16(anon_sym_GT),
	4317: uint16(sym__S),
	4318: uint16(1),
	4319: uint16(935),
	4320: uint16(2),
	4321: uint16(anon_sym_QMARK_GT),
	4322: uint16(sym__S),
	4323: uint16(1),
	4324: uint16(937),
	4325: uint16(2),
	4326: uint16(anon_sym_GT),
	4327: uint16(sym__S),
	4328: uint16(2),
	4329: uint16(882),
	4330: uint16(1),
	4331: uint16(anon_sym_PIPE),
	4332: uint16(939),
	4333: uint16(1),
	4334: uint16(anon_sym_RPAREN),
	4335: uint16(1),
	4336: uint16(941),
	4337: uint16(2),
	4338: uint16(anon_sym_PERCENT),
	4339: uint16(sym__S),
	4340: uint16(2),
	4341: uint16(868),
	4342: uint16(1),
	4343: uint16(anon_sym_GT),
	4344: uint16(943),
	4345: uint16(1),
	4346: uint16(sym__S),
	4347: uint16(1),
	4348: uint16(945),
	4349: uint16(2),
	4350: uint16(anon_sym_PERCENT),
	4351: uint16(sym__S),
	4352: uint16(2),
	4353: uint16(947),
	4354: uint16(1),
	4355: uint16(anon_sym_RPAREN),
	4356: uint16(949),
	4357: uint16(1),
	4358: uint16(sym__S),
	4359: uint16(1),
	4360: uint16(951),
	4361: uint16(2),
	4362: uint16(anon_sym_GT),
	4363: uint16(sym__S),
	4364: uint16(1),
	4365: uint16(953),
	4366: uint16(2),
	4367: uint16(anon_sym_GT),
	4368: uint16(sym__S),
	4369: uint16(1),
	4370: uint16(955),
	4371: uint16(2),
	4372: uint16(anon_sym_GT),
	4373: uint16(sym__S),
	4374: uint16(1),
	4375: uint16(957),
	4376: uint16(2),
	4377: uint16(anon_sym_GT),
	4378: uint16(sym__S),
	4379: uint16(2),
	4380: uint16(959),
	4381: uint16(1),
	4382: uint16(anon_sym_GT),
	4383: uint16(961),
	4384: uint16(1),
	4385: uint16(sym__S),
	4386: uint16(2),
	4387: uint16(963),
	4388: uint16(1),
	4389: uint16(anon_sym_RPAREN),
	4390: uint16(965),
	4391: uint16(1),
	4392: uint16(sym__S),
	4393: uint16(1),
	4394: uint16(830),
	4395: uint16(2),
	4396: uint16(anon_sym_SQUOTE),
	4397: uint16(anon_sym_DQUOTE),
	4398: uint16(1),
	4399: uint16(967),
	4400: uint16(2),
	4401: uint16(anon_sym_PERCENT),
	4402: uint16(sym__S),
	4403: uint16(1),
	4404: uint16(969),
	4405: uint16(2),
	4406: uint16(anon_sym_GT),
	4407: uint16(sym__S),
	4408: uint16(1),
	4409: uint16(971),
	4410: uint16(2),
	4411: uint16(anon_sym_GT),
	4412: uint16(sym__S),
	4413: uint16(2),
	4414: uint16(973),
	4415: uint16(1),
	4416: uint16(anon_sym_RBRACK_RBRACK_GT),
	4417: uint16(975),
	4418: uint16(1),
	4419: uint16(sym_CData),
	4420: uint16(1),
	4421: uint16(977),
	4422: uint16(2),
	4423: uint16(anon_sym_GT),
	4424: uint16(sym__S),
	4425: uint16(2),
	4426: uint16(979),
	4427: uint16(1),
	4428: uint16(anon_sym_RPAREN),
	4429: uint16(981),
	4430: uint16(1),
	4431: uint16(sym__S),
	4432: uint16(2),
	4433: uint16(983),
	4434: uint16(1),
	4435: uint16(anon_sym_GT),
	4436: uint16(985),
	4437: uint16(1),
	4438: uint16(anon_sym_NDATA),
	4439: uint16(1),
	4440: uint16(987),
	4441: uint16(2),
	4442: uint16(anon_sym_GT),
	4443: uint16(sym__S),
	4444: uint16(1),
	4445: uint16(989),
	4446: uint16(2),
	4447: uint16(anon_sym_GT),
	4448: uint16(sym__S),
	4449: uint16(2),
	4450: uint16(991),
	4451: uint16(1),
	4452: uint16(anon_sym_QMARK_GT),
	4453: uint16(993),
	4454: uint16(1),
	4455: uint16(sym__S),
	4456: uint16(2),
	4457: uint16(995),
	4458: uint16(1),
	4459: uint16(anon_sym_GT),
	4460: uint16(997),
	4461: uint16(1),
	4462: uint16(sym__S),
	4463: uint16(2),
	4464: uint16(999),
	4465: uint16(1),
	4466: uint16(sym__S),
	4467: uint16(208),
	4468: uint16(1),
	4469: uint16(sym__VersionInfo),
	4470: uint16(2),
	4471: uint16(838),
	4472: uint16(1),
	4473: uint16(anon_sym_standalone),
	4474: uint16(1001),
	4475: uint16(1),
	4476: uint16(anon_sym_QMARK_GT),
	4477: uint16(2),
	4478: uint16(1001),
	4479: uint16(1),
	4480: uint16(anon_sym_QMARK_GT),
	4481: uint16(1003),
	4482: uint16(1),
	4483: uint16(sym__S),
	4484: uint16(1),
	4485: uint16(995),
	4486: uint16(2),
	4487: uint16(anon_sym_GT),
	4488: uint16(sym__S),
	4489: uint16(1),
	4490: uint16(1005),
	4491: uint16(2),
	4492: uint16(anon_sym_PERCENT),
	4493: uint16(sym__S),
	4494: uint16(2),
	4495: uint16(743),
	4496: uint16(1),
	4497: uint16(sym_Name),
	4498: uint16(368),
	4499: uint16(1),
	4500: uint16(sym_PseudoAtt),
	4501: uint16(2),
	4502: uint16(983),
	4503: uint16(1),
	4504: uint16(anon_sym_GT),
	4505: uint16(1007),
	4506: uint16(1),
	4507: uint16(sym__S),
	4508: uint16(1),
	4509: uint16(1009),
	4510: uint16(2),
	4511: uint16(anon_sym_PERCENT),
	4512: uint16(sym__S),
	4513: uint16(2),
	4514: uint16(1011),
	4515: uint16(1),
	4516: uint16(anon_sym_GT),
	4517: uint16(1013),
	4518: uint16(1),
	4519: uint16(sym__S),
	4520: uint16(2),
	4521: uint16(1015),
	4522: uint16(1),
	4523: uint16(anon_sym_GT),
	4524: uint16(1017),
	4525: uint16(1),
	4526: uint16(sym__S),
	4527: uint16(2),
	4528: uint16(1019),
	4529: uint16(1),
	4530: uint16(anon_sym_GT),
	4531: uint16(1021),
	4532: uint16(1),
	4533: uint16(sym__S),
	4534: uint16(1),
	4535: uint16(786),
	4536: uint16(2),
	4537: uint16(anon_sym_GT),
	4538: uint16(sym__S),
	4539: uint16(1),
	4540: uint16(842),
	4541: uint16(2),
	4542: uint16(anon_sym_GT),
	4543: uint16(sym__S),
	4544: uint16(1),
	4545: uint16(1023),
	4546: uint16(2),
	4547: uint16(anon_sym_PERCENT),
	4548: uint16(sym__S),
	4549: uint16(2),
	4550: uint16(29),
	4551: uint16(1),
	4552: uint16(anon_sym_LT_SLASH),
	4553: uint16(194),
	4554: uint16(1),
	4555: uint16(sym_ETag),
	4556: uint16(1),
	4557: uint16(342),
	4558: uint16(2),
	4559: uint16(anon_sym_PERCENT),
	4560: uint16(sym__S),
	4561: uint16(1),
	4562: uint16(1025),
	4563: uint16(2),
	4564: uint16(anon_sym_GT),
	4565: uint16(sym__S),
	4566: uint16(2),
	4567: uint16(806),
	4568: uint16(1),
	4569: uint16(anon_sym_QMARK_GT),
	4570: uint16(1027),
	4571: uint16(1),
	4572: uint16(sym__S),
	4573: uint16(1),
	4574: uint16(360),
	4575: uint16(2),
	4576: uint16(anon_sym_PERCENT),
	4577: uint16(sym__S),
	4578: uint16(1),
	4579: uint16(774),
	4580: uint16(2),
	4581: uint16(anon_sym_GT),
	4582: uint16(sym__S),
	4583: uint16(1),
	4584: uint16(1029),
	4585: uint16(2),
	4586: uint16(anon_sym_QMARK_GT),
	4587: uint16(sym__S),
	4588: uint16(1),
	4589: uint16(1031),
	4590: uint16(2),
	4591: uint16(anon_sym_QMARK_GT),
	4592: uint16(sym__S),
	4593: uint16(1),
	4594: uint16(1033),
	4595: uint16(2),
	4596: uint16(anon_sym_GT),
	4597: uint16(sym__S),
	4598: uint16(2),
	4599: uint16(1035),
	4600: uint16(1),
	4601: uint16(sym_Name),
	4602: uint16(1037),
	4603: uint16(1),
	4604: uint16(sym__S),
	4605: uint16(1),
	4606: uint16(1039),
	4607: uint16(2),
	4608: uint16(sym_CData),
	4609: uint16(anon_sym_RBRACK_RBRACK_GT),
	4610: uint16(1),
	4611: uint16(1041),
	4612: uint16(2),
	4613: uint16(anon_sym_QMARK_GT),
	4614: uint16(sym__S),
	4615: uint16(1),
	4616: uint16(1043),
	4617: uint16(2),
	4618: uint16(anon_sym_PERCENT),
	4619: uint16(sym__S),
	4620: uint16(1),
	4621: uint16(1045),
	4622: uint16(2),
	4623: uint16(anon_sym_GT),
	4624: uint16(sym__S),
	4625: uint16(2),
	4626: uint16(1047),
	4627: uint16(1),
	4628: uint16(anon_sym_QMARK_GT),
	4629: uint16(1049),
	4630: uint16(1),
	4631: uint16(sym__S),
	4632: uint16(1),
	4633: uint16(1051),
	4634: uint16(2),
	4635: uint16(anon_sym_SQUOTE),
	4636: uint16(anon_sym_DQUOTE),
	4637: uint16(2),
	4638: uint16(17),
	4639: uint16(1),
	4640: uint16(anon_sym_LT_SLASH),
	4641: uint16(35),
	4642: uint16(1),
	4643: uint16(sym_ETag),
	4644: uint16(1),
	4645: uint16(1053),
	4646: uint16(2),
	4647: uint16(anon_sym_PERCENT),
	4648: uint16(sym__S),
	4649: uint16(1),
	4650: uint16(747),
	4651: uint16(2),
	4652: uint16(anon_sym_QMARK_GT),
	4653: uint16(sym__S),
	4654: uint16(2),
	4655: uint16(1055),
	4656: uint16(1),
	4657: uint16(anon_sym_GT),
	4658: uint16(1057),
	4659: uint16(1),
	4660: uint16(sym__S),
	4661: uint16(2),
	4662: uint16(1059),
	4663: uint16(1),
	4664: uint16(anon_sym_GT),
	4665: uint16(1061),
	4666: uint16(1),
	4667: uint16(sym__S),
	4668: uint16(2),
	4669: uint16(1063),
	4670: uint16(1),
	4671: uint16(anon_sym_GT),
	4672: uint16(1065),
	4673: uint16(1),
	4674: uint16(sym__S),
	4675: uint16(2),
	4676: uint16(1067),
	4677: uint16(1),
	4678: uint16(anon_sym_SQUOTE),
	4679: uint16(1069),
	4680: uint16(1),
	4681: uint16(anon_sym_DQUOTE),
	4682: uint16(2),
	4683: uint16(601),
	4684: uint16(1),
	4685: uint16(sym_Name),
	4686: uint16(255),
	4687: uint16(1),
	4688: uint16(sym_Attribute),
	4689: uint16(1),
	4690: uint16(1071),
	4691: uint16(2),
	4692: uint16(anon_sym_PERCENT),
	4693: uint16(sym__S),
	4694: uint16(2),
	4695: uint16(1073),
	4696: uint16(1),
	4697: uint16(anon_sym_QMARK_GT),
	4698: uint16(1075),
	4699: uint16(1),
	4700: uint16(sym__S),
	4701: uint16(1),
	4702: uint16(1077),
	4703: uint16(2),
	4704: uint16(anon_sym_GT),
	4705: uint16(sym__S),
	4706: uint16(1),
	4707: uint16(983),
	4708: uint16(1),
	4709: uint16(anon_sym_GT),
	4710: uint16(1),
	4711: uint16(1079),
	4712: uint16(1),
	4713: uint16(aux_sym_CharRef_token2),
	4714: uint16(1),
	4715: uint16(1081),
	4716: uint16(1),
	4717: uint16(anon_sym_EQ),
	4718: uint16(1),
	4719: uint16(921),
	4720: uint16(1),
	4721: uint16(sym_Nmtoken),
	4722: uint16(1),
	4723: uint16(1083),
	4724: uint16(1),
	4725: uint16(sym__S),
	4726: uint16(1),
	4727: uint16(1085),
	4728: uint16(1),
	4729: uint16(sym__S),
	4730: uint16(1),
	4731: uint16(1087),
	4732: uint16(1),
	4733: uint16(anon_sym_GT),
	4734: uint16(1),
	4735: uint16(1089),
	4736: uint16(1),
	4737: uint16(sym__S),
	4738: uint16(1),
	4739: uint16(1091),
	4740: uint16(1),
	4741: uint16(anon_sym_SQUOTE),
	4742: uint16(1),
	4743: uint16(1093),
	4744: uint16(1),
	4745: uint16(sym__S),
	4746: uint16(1),
	4747: uint16(1091),
	4748: uint16(1),
	4749: uint16(anon_sym_DQUOTE),
	4750: uint16(1),
	4751: uint16(1095),
	4752: uint16(1),
	4753: uint16(anon_sym_STAR),
	4754: uint16(1),
	4755: uint16(1097),
	4756: uint16(1),
	4758: uint16(1),
	4759: uint16(1099),
	4760: uint16(1),
	4761: uint16(sym_VersionNum),
	4762: uint16(1),
	4763: uint16(1011),
	4764: uint16(1),
	4765: uint16(anon_sym_GT),
	4766: uint16(1),
	4767: uint16(741),
	4768: uint16(1),
	4769: uint16(anon_sym_STAR),
	4770: uint16(1),
	4771: uint16(1101),
	4772: uint16(1),
	4773: uint16(sym_Nmtoken),
	4774: uint16(1),
	4775: uint16(882),
	4776: uint16(1),
	4777: uint16(anon_sym_PIPE),
	4778: uint16(1),
	4779: uint16(1103),
	4780: uint16(1),
	4781: uint16(sym_EncName),
	4782: uint16(1),
	4783: uint16(1105),
	4784: uint16(1),
	4785: uint16(anon_sym_GT),
	4786: uint16(1),
	4787: uint16(1107),
	4788: uint16(1),
	4789: uint16(anon_sym_PIPE),
	4790: uint16(1),
	4791: uint16(1109),
	4792: uint16(1),
	4793: uint16(anon_sym_SEMI),
	4794: uint16(1),
	4795: uint16(1111),
	4796: uint16(1),
	4797: uint16(sym_EncName),
	4798: uint16(1),
	4799: uint16(1113),
	4800: uint16(1),
	4801: uint16(aux_sym_SystemLiteral_token2),
	4802: uint16(1),
	4803: uint16(1115),
	4804: uint16(1),
	4805: uint16(aux_sym_SystemLiteral_token1),
	4806: uint16(1),
	4807: uint16(1117),
	4808: uint16(1),
	4809: uint16(sym_VersionNum),
	4810: uint16(1),
	4811: uint16(1119),
	4812: uint16(1),
	4813: uint16(aux_sym_PubidLiteral_token2),
	4814: uint16(1),
	4815: uint16(1121),
	4816: uint16(1),
	4817: uint16(aux_sym_PubidLiteral_token1),
	4818: uint16(1),
	4819: uint16(1123),
	4820: uint16(1),
	4821: uint16(anon_sym_GT),
	4822: uint16(1),
	4823: uint16(997),
	4824: uint16(1),
	4825: uint16(sym__S),
	4826: uint16(1),
	4827: uint16(1125),
	4828: uint16(1),
	4829: uint16(anon_sym_LBRACK),
	4830: uint16(1),
	4831: uint16(1127),
	4832: uint16(1),
	4833: uint16(anon_sym_SEMI),
	4834: uint16(1),
	4835: uint16(963),
	4836: uint16(1),
	4837: uint16(anon_sym_RPAREN),
	4838: uint16(1),
	4839: uint16(1129),
	4840: uint16(1),
	4841: uint16(anon_sym_SEMI),
	4842: uint16(1),
	4843: uint16(1131),
	4844: uint16(1),
	4845: uint16(sym__S),
	4846: uint16(1),
	4847: uint16(979),
	4848: uint16(1),
	4849: uint16(anon_sym_RPAREN),
	4850: uint16(1),
	4851: uint16(1133),
	4852: uint16(1),
	4853: uint16(sym__S),
	4854: uint16(1),
	4855: uint16(1135),
	4856: uint16(1),
	4857: uint16(sym__S),
	4858: uint16(1),
	4859: uint16(1137),
	4860: uint16(1),
	4861: uint16(sym_Nmtoken),
	4862: uint16(1),
	4863: uint16(1139),
	4864: uint16(1),
	4865: uint16(anon_sym_RPAREN),
	4866: uint16(1),
	4867: uint16(1141),
	4868: uint16(1),
	4869: uint16(anon_sym_GT),
	4870: uint16(1),
	4871: uint16(1143),
	4872: uint16(1),
	4873: uint16(anon_sym_LPAREN),
	4874: uint16(1),
	4875: uint16(1145),
	4876: uint16(1),
	4877: uint16(sym__S),
	4878: uint16(1),
	4879: uint16(1147),
	4880: uint16(1),
	4881: uint16(sym__S),
	4882: uint16(1),
	4883: uint16(1149),
	4884: uint16(1),
	4885: uint16(anon_sym_GT),
	4886: uint16(1),
	4887: uint16(1151),
	4888: uint16(1),
	4889: uint16(sym__S),
	4890: uint16(1),
	4891: uint16(1153),
	4892: uint16(1),
	4893: uint16(anon_sym_QMARK_GT),
	4894: uint16(1),
	4895: uint16(1155),
	4896: uint16(1),
	4897: uint16(sym__S),
	4898: uint16(1),
	4899: uint16(868),
	4900: uint16(1),
	4901: uint16(anon_sym_GT),
	4902: uint16(1),
	4903: uint16(1157),
	4904: uint16(1),
	4905: uint16(sym__start_tag_name),
	4906: uint16(1),
	4907: uint16(1063),
	4908: uint16(1),
	4909: uint16(anon_sym_GT),
	4910: uint16(1),
	4911: uint16(776),
	4912: uint16(1),
	4913: uint16(anon_sym_STAR),
	4914: uint16(1),
	4915: uint16(1159),
	4916: uint16(1),
	4917: uint16(anon_sym_SQUOTE),
	4918: uint16(1),
	4919: uint16(500),
	4920: uint16(1),
	4921: uint16(anon_sym_PIPE),
	4922: uint16(1),
	4923: uint16(1159),
	4924: uint16(1),
	4925: uint16(anon_sym_DQUOTE),
	4926: uint16(1),
	4927: uint16(1161),
	4928: uint16(1),
	4929: uint16(anon_sym_SQUOTE),
	4930: uint16(1),
	4931: uint16(1161),
	4932: uint16(1),
	4933: uint16(anon_sym_DQUOTE),
	4934: uint16(1),
	4935: uint16(1001),
	4936: uint16(1),
	4937: uint16(anon_sym_QMARK_GT),
	4938: uint16(1),
	4939: uint16(1163),
	4940: uint16(1),
	4941: uint16(anon_sym_RBRACK_RBRACK_GT),
	4942: uint16(1),
	4943: uint16(1165),
	4944: uint16(1),
	4945: uint16(anon_sym_DOCTYPE),
	4946: uint16(1),
	4947: uint16(1079),
	4948: uint16(1),
	4949: uint16(aux_sym_CharRef_token1),
	4950: uint16(1),
	4951: uint16(1167),
	4952: uint16(1),
	4953: uint16(sym__S),
	4954: uint16(1),
	4955: uint16(1169),
	4956: uint16(1),
	4957: uint16(anon_sym_QMARK_GT),
	4958: uint16(1),
	4959: uint16(1171),
	4960: uint16(1),
	4961: uint16(sym_Name),
	4962: uint16(1),
	4963: uint16(1173),
	4964: uint16(1),
	4965: uint16(sym_PITarget),
	4966: uint16(1),
	4967: uint16(1175),
	4968: uint16(1),
	4969: uint16(sym__end_tag_name),
	4970: uint16(1),
	4971: uint16(1177),
	4972: uint16(1),
	4973: uint16(sym_Name),
	4974: uint16(1),
	4975: uint16(1179),
	4976: uint16(1),
	4977: uint16(anon_sym_CDATA),
	4978: uint16(1),
	4979: uint16(1181),
	4980: uint16(1),
	4981: uint16(anon_sym_GT),
	4982: uint16(1),
	4983: uint16(1183),
	4984: uint16(1),
	4985: uint16(sym__pi_content),
	4986: uint16(1),
	4987: uint16(1185),
	4988: uint16(1),
	4989: uint16(sym__S),
	4990: uint16(1),
	4991: uint16(1187),
	4992: uint16(1),
	4993: uint16(sym__S),
	4994: uint16(1),
	4995: uint16(1189),
	4996: uint16(1),
	4997: uint16(sym__S),
	4998: uint16(1),
	4999: uint16(1019),
	5000: uint16(1),
	5001: uint16(anon_sym_GT),
	5002: uint16(1),
	5003: uint16(1191),
	5004: uint16(1),
	5005: uint16(anon_sym_version),
	5006: uint16(1),
	5007: uint16(1193),
	5008: uint16(1),
	5009: uint16(sym__S),
	5010: uint16(1),
	5011: uint16(1195),
	5012: uint16(1),
	5013: uint16(anon_sym_SQUOTE),
	5014: uint16(1),
	5015: uint16(1197),
	5016: uint16(1),
	5017: uint16(anon_sym_SEMI),
	5018: uint16(1),
	5019: uint16(1199),
	5020: uint16(1),
	5021: uint16(anon_sym_SEMI),
	5022: uint16(1),
	5023: uint16(1201),
	5024: uint16(1),
	5025: uint16(anon_sym_QMARK_GT),
	5026: uint16(1),
	5027: uint16(792),
	5028: uint16(1),
	5029: uint16(anon_sym_STAR),
	5030: uint16(1),
	5031: uint16(1203),
	5032: uint16(1),
	5033: uint16(anon_sym_GT),
	5034: uint16(1),
	5035: uint16(1195),
	5036: uint16(1),
	5037: uint16(anon_sym_DQUOTE),
	5038: uint16(1),
	5039: uint16(1205),
	5040: uint16(1),
	5041: uint16(anon_sym_SQUOTE),
	5042: uint16(1),
	5043: uint16(1205),
	5044: uint16(1),
	5045: uint16(anon_sym_DQUOTE),
	5046: uint16(1),
	5047: uint16(1207),
	5048: uint16(1),
	5049: uint16(sym_Name),
	5050: uint16(1),
	5051: uint16(1209),
	5052: uint16(1),
	5053: uint16(anon_sym_SEMI),
	5054: uint16(1),
	5055: uint16(1211),
	5056: uint16(1),
	5057: uint16(sym__S),
	5058: uint16(1),
	5059: uint16(1213),
	5060: uint16(1),
	5061: uint16(anon_sym_SEMI),
	5062: uint16(1),
	5063: uint16(1215),
	5064: uint16(1),
	5065: uint16(anon_sym_SEMI),
	5066: uint16(1),
	5067: uint16(1217),
	5068: uint16(1),
	5069: uint16(anon_sym_QMARK_GT),
	5070: uint16(1),
	5071: uint16(1219),
	5072: uint16(1),
	5073: uint16(anon_sym_SEMI),
	5074: uint16(1),
	5075: uint16(1221),
	5076: uint16(1),
	5077: uint16(anon_sym_SEMI),
	5078: uint16(1),
	5079: uint16(1223),
	5080: uint16(1),
	5081: uint16(anon_sym_SEMI),
	5082: uint16(1),
	5083: uint16(1225),
	5084: uint16(1),
	5085: uint16(anon_sym_SEMI),
	5086: uint16(1),
	5087: uint16(1227),
	5088: uint16(1),
	5089: uint16(anon_sym_SEMI),
	5090: uint16(1),
	5091: uint16(1229),
	5092: uint16(1),
	5093: uint16(anon_sym_SEMI),
	5094: uint16(1),
	5095: uint16(1231),
	5096: uint16(1),
	5097: uint16(sym__start_tag_name),
	5098: uint16(1),
	5099: uint16(1233),
	5100: uint16(1),
	5101: uint16(sym_PITarget),
	5102: uint16(1),
	5103: uint16(1235),
	5104: uint16(1),
	5105: uint16(sym__end_tag_name),
	5106: uint16(1),
	5107: uint16(1237),
	5108: uint16(1),
	5109: uint16(sym_Name),
	5110: uint16(1),
	5111: uint16(1239),
	5112: uint16(1),
	5113: uint16(aux_sym_CharRef_token1),
	5114: uint16(1),
	5115: uint16(1239),
	5116: uint16(1),
	5117: uint16(aux_sym_CharRef_token2),
	5118: uint16(1),
	5119: uint16(1241),
	5120: uint16(1),
	5121: uint16(sym__pi_content),
	5122: uint16(1),
	5123: uint16(1035),
	5124: uint16(1),
	5125: uint16(sym_Name),
	5126: uint16(1),
	5127: uint16(1243),
	5128: uint16(1),
	5129: uint16(sym_Name),
	5130: uint16(1),
	5131: uint16(1245),
	5132: uint16(1),
	5133: uint16(aux_sym_CharRef_token1),
	5134: uint16(1),
	5135: uint16(1245),
	5136: uint16(1),
	5137: uint16(aux_sym_CharRef_token2),
	5138: uint16(1),
	5139: uint16(1247),
	5140: uint16(1),
	5141: uint16(sym__pi_content),
	5142: uint16(1),
	5143: uint16(1249),
	5144: uint16(1),
	5145: uint16(sym_Name),
	5146: uint16(1),
	5147: uint16(1251),
	5148: uint16(1),
	5149: uint16(sym_Name),
	5150: uint16(1),
	5151: uint16(1253),
	5152: uint16(1),
	5153: uint16(aux_sym_CharRef_token1),
	5154: uint16(1),
	5155: uint16(1253),
	5156: uint16(1),
	5157: uint16(aux_sym_CharRef_token2),
	5158: uint16(1),
	5159: uint16(1255),
	5160: uint16(1),
	5161: uint16(sym_Name),
	5162: uint16(1),
	5163: uint16(1257),
	5164: uint16(1),
	5165: uint16(sym_Name),
	5166: uint16(1),
	5167: uint16(1259),
	5168: uint16(1),
	5169: uint16(aux_sym_CharRef_token1),
	5170: uint16(1),
	5171: uint16(1259),
	5172: uint16(1),
	5173: uint16(aux_sym_CharRef_token2),
	5174: uint16(1),
	5175: uint16(1261),
	5176: uint16(1),
	5177: uint16(sym_Name),
}

var ts_small_parse_table_map = [493]uint32_t{
	1:   uint32(52),
	2:   uint32(104),
	3:   uint32(150),
	4:   uint32(196),
	5:   uint32(235),
	6:   uint32(268),
	7:   uint32(301),
	8:   uint32(334),
	9:   uint32(367),
	10:  uint32(400),
	11:  uint32(433),
	12:  uint32(466),
	13:  uint32(499),
	14:  uint32(532),
	15:  uint32(559),
	16:  uint32(586),
	17:  uint32(613),
	18:  uint32(640),
	19:  uint32(664),
	20:  uint32(688),
	21:  uint32(712),
	22:  uint32(740),
	23:  uint32(768),
	24:  uint32(790),
	25:  uint32(812),
	26:  uint32(840),
	27:  uint32(868),
	28:  uint32(896),
	29:  uint32(920),
	30:  uint32(942),
	31:  uint32(970),
	32:  uint32(998),
	33:  uint32(1011),
	34:  uint32(1025),
	35:  uint32(1049),
	36:  uint32(1073),
	37:  uint32(1091),
	38:  uint32(1109),
	39:  uint32(1123),
	40:  uint32(1141),
	41:  uint32(1159),
	42:  uint32(1183),
	43:  uint32(1207),
	44:  uint32(1231),
	45:  uint32(1255),
	46:  uint32(1273),
	47:  uint32(1291),
	48:  uint32(1305),
	49:  uint32(1323),
	50:  uint32(1337),
	51:  uint32(1359),
	52:  uint32(1377),
	53:  uint32(1395),
	54:  uint32(1409),
	55:  uint32(1423),
	56:  uint32(1437),
	57:  uint32(1455),
	58:  uint32(1469),
	59:  uint32(1493),
	60:  uint32(1517),
	61:  uint32(1529),
	62:  uint32(1547),
	63:  uint32(1565),
	64:  uint32(1577),
	65:  uint32(1601),
	66:  uint32(1613),
	67:  uint32(1625),
	68:  uint32(1637),
	69:  uint32(1651),
	70:  uint32(1665),
	71:  uint32(1679),
	72:  uint32(1693),
	73:  uint32(1717),
	74:  uint32(1731),
	75:  uint32(1745),
	76:  uint32(1759),
	77:  uint32(1773),
	78:  uint32(1787),
	79:  uint32(1801),
	80:  uint32(1825),
	81:  uint32(1849),
	82:  uint32(1873),
	83:  uint32(1897),
	84:  uint32(1921),
	85:  uint32(1942),
	86:  uint32(1955),
	87:  uint32(1976),
	88:  uint32(1997),
	89:  uint32(2018),
	90:  uint32(2043),
	91:  uint32(2063),
	92:  uint32(2083),
	93:  uint32(2105),
	94:  uint32(2127),
	95:  uint32(2149),
	96:  uint32(2171),
	97:  uint32(2191),
	98:  uint32(2213),
	99:  uint32(2224),
	100: uint32(2235),
	101: uint32(2252),
	102: uint32(2271),
	103: uint32(2282),
	104: uint32(2293),
	105: uint32(2310),
	106: uint32(2321),
	107: uint32(2338),
	108: uint32(2355),
	109: uint32(2370),
	110: uint32(2381),
	111: uint32(2392),
	112: uint32(2409),
	113: uint32(2428),
	114: uint32(2447),
	115: uint32(2466),
	116: uint32(2485),
	117: uint32(2496),
	118: uint32(2507),
	119: uint32(2518),
	120: uint32(2529),
	121: uint32(2540),
	122: uint32(2551),
	123: uint32(2562),
	124: uint32(2573),
	125: uint32(2584),
	126: uint32(2598),
	127: uint32(2608),
	128: uint32(2618),
	129: uint32(2634),
	130: uint32(2648),
	131: uint32(2656),
	132: uint32(2664),
	133: uint32(2672),
	134: uint32(2680),
	135: uint32(2694),
	136: uint32(2704),
	137: uint32(2718),
	138: uint32(2734),
	139: uint32(2748),
	140: uint32(2762),
	141: uint32(2772),
	142: uint32(2788),
	143: uint32(2798),
	144: uint32(2808),
	145: uint32(2818),
	146: uint32(2834),
	147: uint32(2844),
	148: uint32(2854),
	149: uint32(2864),
	150: uint32(2878),
	151: uint32(2886),
	152: uint32(2900),
	153: uint32(2908),
	154: uint32(2916),
	155: uint32(2932),
	156: uint32(2942),
	157: uint32(2956),
	158: uint32(2969),
	159: uint32(2982),
	160: uint32(2993),
	161: uint32(3006),
	162: uint32(3015),
	163: uint32(3028),
	164: uint32(3039),
	165: uint32(3048),
	166: uint32(3059),
	167: uint32(3066),
	168: uint32(3079),
	169: uint32(3088),
	170: uint32(3097),
	171: uint32(3110),
	172: uint32(3123),
	173: uint32(3136),
	174: uint32(3143),
	175: uint32(3154),
	176: uint32(3161),
	177: uint32(3174),
	178: uint32(3181),
	179: uint32(3194),
	180: uint32(3207),
	181: uint32(3220),
	182: uint32(3233),
	183: uint32(3246),
	184: uint32(3259),
	185: uint32(3272),
	186: uint32(3279),
	187: uint32(3290),
	188: uint32(3297),
	189: uint32(3310),
	190: uint32(3317),
	191: uint32(3330),
	192: uint32(3343),
	193: uint32(3350),
	194: uint32(3363),
	195: uint32(3376),
	196: uint32(3389),
	197: uint32(3396),
	198: uint32(3409),
	199: uint32(3422),
	200: uint32(3435),
	201: uint32(3442),
	202: uint32(3451),
	203: uint32(3458),
	204: uint32(3467),
	205: uint32(3474),
	206: uint32(3483),
	207: uint32(3496),
	208: uint32(3505),
	209: uint32(3516),
	210: uint32(3525),
	211: uint32(3538),
	212: uint32(3551),
	213: uint32(3564),
	214: uint32(3577),
	215: uint32(3584),
	216: uint32(3593),
	217: uint32(3599),
	218: uint32(3605),
	219: uint32(3615),
	220: uint32(3625),
	221: uint32(3635),
	222: uint32(3643),
	223: uint32(3653),
	224: uint32(3663),
	225: uint32(3673),
	226: uint32(3683),
	227: uint32(3693),
	228: uint32(3703),
	229: uint32(3709),
	230: uint32(3719),
	231: uint32(3729),
	232: uint32(3739),
	233: uint32(3749),
	234: uint32(3759),
	235: uint32(3767),
	236: uint32(3777),
	237: uint32(3785),
	238: uint32(3795),
	239: uint32(3801),
	240: uint32(3807),
	241: uint32(3817),
	242: uint32(3827),
	243: uint32(3837),
	244: uint32(3845),
	245: uint32(3853),
	246: uint32(3863),
	247: uint32(3873),
	248: uint32(3879),
	249: uint32(3889),
	250: uint32(3899),
	251: uint32(3909),
	252: uint32(3915),
	253: uint32(3925),
	254: uint32(3931),
	255: uint32(3941),
	256: uint32(3947),
	257: uint32(3957),
	258: uint32(3967),
	259: uint32(3977),
	260: uint32(3987),
	261: uint32(3997),
	262: uint32(4007),
	263: uint32(4017),
	264: uint32(4027),
	265: uint32(4035),
	266: uint32(4041),
	267: uint32(4051),
	268: uint32(4061),
	269: uint32(4067),
	270: uint32(4075),
	271: uint32(4085),
	272: uint32(4095),
	273: uint32(4105),
	274: uint32(4115),
	275: uint32(4120),
	276: uint32(4127),
	277: uint32(4134),
	278: uint32(4139),
	279: uint32(4144),
	280: uint32(4149),
	281: uint32(4154),
	282: uint32(4161),
	283: uint32(4166),
	284: uint32(4173),
	285: uint32(4178),
	286: uint32(4185),
	287: uint32(4192),
	288: uint32(4199),
	289: uint32(4204),
	290: uint32(4209),
	291: uint32(4214),
	292: uint32(4219),
	293: uint32(4224),
	294: uint32(4229),
	295: uint32(4234),
	296: uint32(4239),
	297: uint32(4246),
	298: uint32(4253),
	299: uint32(4258),
	300: uint32(4263),
	301: uint32(4270),
	302: uint32(4275),
	303: uint32(4282),
	304: uint32(4289),
	305: uint32(4294),
	306: uint32(4299),
	307: uint32(4306),
	308: uint32(4313),
	309: uint32(4318),
	310: uint32(4323),
	311: uint32(4328),
	312: uint32(4335),
	313: uint32(4340),
	314: uint32(4347),
	315: uint32(4352),
	316: uint32(4359),
	317: uint32(4364),
	318: uint32(4369),
	319: uint32(4374),
	320: uint32(4379),
	321: uint32(4386),
	322: uint32(4393),
	323: uint32(4398),
	324: uint32(4403),
	325: uint32(4408),
	326: uint32(4413),
	327: uint32(4420),
	328: uint32(4425),
	329: uint32(4432),
	330: uint32(4439),
	331: uint32(4444),
	332: uint32(4449),
	333: uint32(4456),
	334: uint32(4463),
	335: uint32(4470),
	336: uint32(4477),
	337: uint32(4484),
	338: uint32(4489),
	339: uint32(4494),
	340: uint32(4501),
	341: uint32(4508),
	342: uint32(4513),
	343: uint32(4520),
	344: uint32(4527),
	345: uint32(4534),
	346: uint32(4539),
	347: uint32(4544),
	348: uint32(4549),
	349: uint32(4556),
	350: uint32(4561),
	351: uint32(4566),
	352: uint32(4573),
	353: uint32(4578),
	354: uint32(4583),
	355: uint32(4588),
	356: uint32(4593),
	357: uint32(4598),
	358: uint32(4605),
	359: uint32(4610),
	360: uint32(4615),
	361: uint32(4620),
	362: uint32(4625),
	363: uint32(4632),
	364: uint32(4637),
	365: uint32(4644),
	366: uint32(4649),
	367: uint32(4654),
	368: uint32(4661),
	369: uint32(4668),
	370: uint32(4675),
	371: uint32(4682),
	372: uint32(4689),
	373: uint32(4694),
	374: uint32(4701),
	375: uint32(4706),
	376: uint32(4710),
	377: uint32(4714),
	378: uint32(4718),
	379: uint32(4722),
	380: uint32(4726),
	381: uint32(4730),
	382: uint32(4734),
	383: uint32(4738),
	384: uint32(4742),
	385: uint32(4746),
	386: uint32(4750),
	387: uint32(4754),
	388: uint32(4758),
	389: uint32(4762),
	390: uint32(4766),
	391: uint32(4770),
	392: uint32(4774),
	393: uint32(4778),
	394: uint32(4782),
	395: uint32(4786),
	396: uint32(4790),
	397: uint32(4794),
	398: uint32(4798),
	399: uint32(4802),
	400: uint32(4806),
	401: uint32(4810),
	402: uint32(4814),
	403: uint32(4818),
	404: uint32(4822),
	405: uint32(4826),
	406: uint32(4830),
	407: uint32(4834),
	408: uint32(4838),
	409: uint32(4842),
	410: uint32(4846),
	411: uint32(4850),
	412: uint32(4854),
	413: uint32(4858),
	414: uint32(4862),
	415: uint32(4866),
	416: uint32(4870),
	417: uint32(4874),
	418: uint32(4878),
	419: uint32(4882),
	420: uint32(4886),
	421: uint32(4890),
	422: uint32(4894),
	423: uint32(4898),
	424: uint32(4902),
	425: uint32(4906),
	426: uint32(4910),
	427: uint32(4914),
	428: uint32(4918),
	429: uint32(4922),
	430: uint32(4926),
	431: uint32(4930),
	432: uint32(4934),
	433: uint32(4938),
	434: uint32(4942),
	435: uint32(4946),
	436: uint32(4950),
	437: uint32(4954),
	438: uint32(4958),
	439: uint32(4962),
	440: uint32(4966),
	441: uint32(4970),
	442: uint32(4974),
	443: uint32(4978),
	444: uint32(4982),
	445: uint32(4986),
	446: uint32(4990),
	447: uint32(4994),
	448: uint32(4998),
	449: uint32(5002),
	450: uint32(5006),
	451: uint32(5010),
	452: uint32(5014),
	453: uint32(5018),
	454: uint32(5022),
	455: uint32(5026),
	456: uint32(5030),
	457: uint32(5034),
	458: uint32(5038),
	459: uint32(5042),
	460: uint32(5046),
	461: uint32(5050),
	462: uint32(5054),
	463: uint32(5058),
	464: uint32(5062),
	465: uint32(5066),
	466: uint32(5070),
	467: uint32(5074),
	468: uint32(5078),
	469: uint32(5082),
	470: uint32(5086),
	471: uint32(5090),
	472: uint32(5094),
	473: uint32(5098),
	474: uint32(5102),
	475: uint32(5106),
	476: uint32(5110),
	477: uint32(5114),
	478: uint32(5118),
	479: uint32(5122),
	480: uint32(5126),
	481: uint32(5130),
	482: uint32(5134),
	483: uint32(5138),
	484: uint32(5142),
	485: uint32(5146),
	486: uint32(5150),
	487: uint32(5154),
	488: uint32(5158),
	489: uint32(5162),
	490: uint32(5166),
	491: uint32(5170),
	492: uint32(5174),
}

var ts_parse_actions = [1263]TSParseActionEntry{
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(159)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(436)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(426)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(6)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(20)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(441)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(474)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(476)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(444)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(462)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(437)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(378)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(4)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(442)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_content),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(5)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	36: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
	37: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(441)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	38: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	39: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(474)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_content_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
	45: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(444)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	46: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	47: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(462)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
	51: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(437)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	52: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	53: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(378)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_content_repeat1),
	})))),
	57: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(475)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	61: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	62: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	63: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	64: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(19)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(314)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	69: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(371)),
	}})))),
	70: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(17)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	73: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	74: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(315)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(16)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym__intSubset),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(475)),
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
		Fsymbol:      uint16(aux_sym__intSubset),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(162)),
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
		Fcount:    uint8(1),
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
		Fsymbol:      uint16(aux_sym__intSubset),
	})))),
	86: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
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
	87: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__intSubset),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(21)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fsymbol:      uint16(aux_sym__intSubset),
	})))),
	91: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(308)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(18)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(228)),
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
		Fcount: uint8(1),
	}})),
	101: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_prolog),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	109: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(486)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(488)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	117: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(29)),
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
		Fcount: uint8(1),
	}})),
	123: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(228)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_document_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(25)),
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
		Fsymbol:      uint16(sym__cp),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(132)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(307)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(309)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(452)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(481)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(352)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(32)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fsymbol:      uint16(sym_prolog),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
	})))),
	164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(486)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(32)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	169: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
	})))),
	170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(487)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(488)),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EntityValue_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(489)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(490)),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(33)),
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
		Fcount: uint8(2),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(491)),
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
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(492)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EntityValue_repeat1),
	})))),
	193: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(493)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(sym_PEReference),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_element),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_element),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(45)),
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
		Fcount: uint8(1),
	}})),
	205: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(478)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(479)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	211: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount: uint8(1),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(482)),
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
		Fcount: uint8(1),
	}})),
	215: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	217: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	219: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(1),
		Fsymbol:        uint16(sym_document),
		Fproduction_id: uint16(1),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(64)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	223: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_document),
		Fproduction_id: uint16(2),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(42)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	227: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_STag),
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
		Fcount: uint8(1),
	}})),
	229: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_STag),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	231: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_document),
		Fproduction_id: uint16(3),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	233: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	235: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_document),
		Fproduction_id: uint16(2),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(270)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat2),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	241: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(45)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	244: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat2),
	})))),
	245: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(477)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	247: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(478)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat2),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(479)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	253: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	255: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(46)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(2),
	}})),
	258: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat1),
	})))),
	259: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(482)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount: uint8(2),
	}})),
	261: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(483)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_PseudoAttValue_repeat1),
	})))),
	265: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(484)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_prolog),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	269: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_CDSect),
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
		Fsymbol:      uint16(sym_CDSect),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	277: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_STag),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_STag),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(358)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(91)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	285: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(4),
		Fsymbol:        uint16(sym_document),
		Fproduction_id: uint16(3),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_prolog),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	289: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_EntityRef),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_EntityRef),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	293: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_CharRef),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_CharRef),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_CDSect),
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
		Fcount: uint8(1),
	}})),
	299: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_CDSect),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(63)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	303: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_STag),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_STag),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(240)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(44)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__choice),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(2),
		Fsymbol:        uint16(sym_document),
		Fproduction_id: uint16(1),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	319: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym__choice),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(82)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(83)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(448)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym__choice),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym__choice),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	333: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym__choice),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	335: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
	})))),
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
		Fcount: uint8(1),
	}})),
	337: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_element),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_element),
	})))),
	340: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	341: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_element),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PI),
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
		Fcount: uint8(1),
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
		Fsymbol:      uint16(sym_PI),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(sym_EmptyElemTag),
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
		Fcount: uint8(1),
	}})),
	349: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_EmptyElemTag),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(348)),
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
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_EmptyElemTag),
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
		Fcount: uint8(1),
	}})),
	355: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_EmptyElemTag),
	})))),
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
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_ETag),
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
		Fcount: uint8(1),
	}})),
	359: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_ETag),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_PI),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_PI),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_EmptyElemTag),
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
		Fcount: uint8(1),
	}})),
	367: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_EmptyElemTag),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_ETag),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_ETag),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__Reference),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__Reference),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(43)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(347)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(84)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(74)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	385: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(65)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	391: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	395: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	399: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	403: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	404: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(94)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(95)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	409: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(113)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(172)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(271)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	417: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(97)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(392)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(139)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(223)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(143)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(108)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(428)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(130)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	435: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_XmlModelPI),
	})))),
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
		Fcount: uint8(1),
	}})),
	437: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_XmlModelPI),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(23)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(24)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(447)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(449)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	447: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_XmlModelPI),
	})))),
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
		Fcount: uint8(1),
	}})),
	449: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_XmlModelPI),
	})))),
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
		Fsymbol:      uint16(sym_StyleSheetPI),
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
		Fcount: uint8(1),
	}})),
	453: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_StyleSheetPI),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(381)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	457: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(98)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	462: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(282)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_StyleSheetPI),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_StyleSheetPI),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(317)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(203)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(397)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(323)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(330)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	479: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_StyleSheetPI),
	})))),
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
		Fcount: uint8(1),
	}})),
	481: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_StyleSheetPI),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_XmlModelPI),
	})))),
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
		Fcount: uint8(1),
	}})),
	485: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_XmlModelPI),
	})))),
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
		Fcount: uint8(1),
	}})),
	487: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PEReference),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(68)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(275)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_XMLDecl),
	})))),
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
		Fcount: uint8(1),
	}})),
	495: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_XMLDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_XMLDecl),
	})))),
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
		Fcount: uint8(1),
	}})),
	499: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_XMLDecl),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(180)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(457)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(93)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__cp),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym__choice_repeat1),
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
		Fsymbol:      uint16(sym_XMLDecl),
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
		Fcount: uint8(1),
	}})),
	515: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_XMLDecl),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(231)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	519: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
	})))),
	520: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(172)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	522: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(430)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(203)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(397)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(236)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_XMLDecl),
	})))),
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
		Fcount: uint8(1),
	}})),
	539: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_XMLDecl),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(11)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(171)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(239)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym__intSubset),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_children),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(363)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat2),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	555: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat2),
	})))),
	556: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	558: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym__choice_repeat2),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(283)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount: uint8(1),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(336)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(247)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(220)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(334)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(464)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(374)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(164)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(411)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(413)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(414)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(420)),
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
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_doctypedecl),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_doctypedecl),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	589: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttlistDecl_repeat1),
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
		Fcount:    uint8(2),
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
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttlistDecl_repeat1),
	})))),
	592: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(242)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
	593: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	594: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(11),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
	595: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	596: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(11),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
	597: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	598: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(494)),
	}})))),
	599: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(152)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(59)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	606: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	607: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	608: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(12),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
	609: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	610: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(12),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	612: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
	613: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount: uint8(1),
	}})),
	614: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
	615: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(179)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(222)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(426)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	622: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EmptyElemTag_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	624: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_EmptyElemTag_repeat1),
	})))),
	625: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(373)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(40)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	629: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	631: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(191)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	635: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	637: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Mixed_repeat2),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	639: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Mixed_repeat2),
	})))),
	640: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	641: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	642: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	643: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(168)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(169)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(285)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(245)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(301)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(302)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(114)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym_Mixed_repeat1),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(388)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(312)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(313)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	676: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
	677: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	679: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	681: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
	682: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(394)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	686: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(204)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	689: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
	})))),
	690: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(206)),
		Frepetition: libc.BoolUint8(true1 != 0),
	}})))),
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
		Fsymbol:      uint16(aux_sym_NotationType_repeat1),
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
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_doctypedecl),
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
		Fcount: uint8(1),
	}})),
	696: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_doctypedecl),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(269)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
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
		Fcount: uint8(1),
	}})),
	704: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_doctypedecl),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(280)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(160)),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_doctypedecl),
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
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_doctypedecl),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	714: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	715: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	716: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	717: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(75)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(78)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_doctypedecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_ExternalID),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	730: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_Attribute),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(379)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(238)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_Mixed),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_StyleSheetPI_repeat1),
	})))),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	750: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_StyleSheetPI_repeat1),
	})))),
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
		Fstate:      libc.Uint16FromInt32(libc.Int32FromInt32(341)),
		Frepetition: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(325)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(331)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(104)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(119)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(15)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(209)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(147)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(400)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(401)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(293)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(198)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(295)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(403)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(404)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__Eq),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(324)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_AttValue),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	789: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_SystemLiteral),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(310)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(aux_sym_Mixed_repeat2),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(296)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(112)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(265)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(229)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(115)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(137)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(337)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(aux_sym_Enumeration_repeat1),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(419)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(210)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(422)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(85)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(61)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(424)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(359)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(105)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	831: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym__Eq),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	833: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(365)),
	}})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	835: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_ExternalID),
	})))),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(224)),
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(260)),
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
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	841: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fsymbol:        uint16(sym_AttValue),
		Fproduction_id: uint16(4),
	})))),
	844: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	845: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(281)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(81)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(60)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(7)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(217)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(278)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(69)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym__VersionInfo),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(177)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(445)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(8)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(207)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(453)),
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
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_AttlistDecl),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_AttlistDecl),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(380)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(249)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	881: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_Enumeration),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(305)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(297)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(417)),
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
		Fcount:    uint8(2),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	889: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(aux_sym_AttlistDecl_repeat1),
	})))),
	890: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	891: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	892: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_DefaultDecl),
	})))),
	893: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	894: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_AttDef),
	})))),
	895: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	896: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(10),
		Fsymbol:      uint16(sym_PEDecl),
	})))),
	897: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	898: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_NDataDecl),
	})))),
	899: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	900: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	901: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	902: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_NotationDecl),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	904: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(390)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	906: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(402)),
	}})))),
	907: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	908: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(395)),
	}})))),
	909: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	910: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	911: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	912: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_PseudoAttValue),
	})))),
	913: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	914: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_Enumeration),
	})))),
	915: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	916: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_elementdecl),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(377)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(393)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(253)),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_EntityValue),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_StringType),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(425)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(415)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_PseudoAttValue),
		Fproduction_id: uint16(4),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(5),
		Fsymbol:      uint16(sym_Enumeration),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(321)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__markupdecl),
	})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(391)),
	}})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__EntityDecl),
	})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(327)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(409)),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__AttType),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_DefaultDecl),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_Mixed),
	})))),
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
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_Enumeration),
	})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(4),
		Fsymbol:      uint16(sym_AttDef),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(329)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	966: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(412)),
	}})))),
	967: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	968: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_GEDecl),
	})))),
	969: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	970: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PubidLiteral),
	})))),
	971: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	972: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_NotationType),
	})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	974: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	975: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	976: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(435)),
	}})))),
	977: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	978: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_NotationType),
	})))),
	979: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	980: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	981: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	982: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(416)),
	}})))),
	983: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	984: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(367)),
	}})))),
	985: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	986: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(382)),
	}})))),
	987: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	988: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_NotationType),
	})))),
	989: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	990: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_NotationType),
	})))),
	991: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	992: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	993: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	994: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	995: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	996: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PublicID),
	})))),
	997: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	998: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	999: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1000: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1001: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1002: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(128)),
	}})))),
	1003: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1004: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1005: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1006: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_NotationDecl),
	})))),
	1007: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1008: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(383)),
	}})))),
	1009: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1010: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_PEDecl),
	})))),
	1011: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1012: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1013: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1014: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(427)),
	}})))),
	1015: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1016: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(343)),
	}})))),
	1017: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(421)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1022: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(396)),
	}})))),
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
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(9),
		Fsymbol:      uint16(sym_GEDecl),
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
		Ftype_token:    uint8(TSParseActionTypeReduce),
		Fchild_count:   uint8(3),
		Fsymbol:        uint16(sym_EntityValue),
		Fproduction_id: uint16(4),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1028: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(434)),
	}})))),
	1029: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1030: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym__SDDecl),
	})))),
	1031: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1032: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym__EncodingDecl),
	})))),
	1033: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1034: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym_contentspec),
	})))),
	1035: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1036: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1037: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1038: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1039: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1040: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_CDStart),
	})))),
	1041: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1042: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym_PseudoAtt),
	})))),
	1043: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1044: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(7),
		Fsymbol:      uint16(sym_elementdecl),
	})))),
	1045: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1046: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(2),
		Fsymbol:      uint16(sym_children),
	})))),
	1047: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1048: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1049: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1050: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1051: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1052: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(3),
		Fsymbol:      uint16(sym__Eq),
	})))),
	1053: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1054: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(8),
		Fsymbol:      uint16(sym_GEDecl),
	})))),
	1055: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1056: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1057: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1058: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1059: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1060: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1061: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1062: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(405)),
	}})))),
	1063: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1064: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1065: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1066: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1067: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1068: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1069: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1070: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1071: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1072: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(6),
		Fsymbol:      uint16(sym_AttlistDecl),
	})))),
	1073: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1074: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1075: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1076: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(485)),
	}})))),
	1077: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1078: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		Ftype_token         uint8_t
		Fchild_count        uint8_t
		Fsymbol             TSSymbol
		Fdynamic_precedence int16_t
		Fproduction_id      uint16_t
	}{
		Ftype_token:  uint8(TSParseActionTypeReduce),
		Fchild_count: uint8(1),
		Fsymbol:      uint16(sym__EnumeratedType),
	})))),
	1079: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1080: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(410)),
	}})))),
	1081: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1082: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1083: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1084: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1085: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1086: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1087: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1088: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1089: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1090: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1091: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1092: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1093: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1094: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1095: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1096: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1097: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1098: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f uint8_t
		_ [7]byte
	}{f: uint8(TSParseActionTypeAccept)})),
	1099: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1100: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1101: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1102: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1103: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1104: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(460)),
	}})))),
	1105: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1106: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1107: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1108: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1109: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1110: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1111: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1112: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1113: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1114: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1115: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1116: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1117: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1118: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1119: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1120: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(432)),
	}})))),
	1121: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1122: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1123: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1124: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1125: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1126: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1127: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1128: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1129: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1130: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1131: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1132: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1133: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1134: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1135: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1136: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1137: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1138: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1139: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(333)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(52)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1148: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(106)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(146)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1156: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1157: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1158: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1159: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1160: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1161: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1162: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1163: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1164: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1165: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1166: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(438)),
	}})))),
	1167: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1168: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(440)),
	}})))),
	1169: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1170: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1171: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(232)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(364)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(277)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(386)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(407)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(187)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(439)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1188: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1189: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1190: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1191: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1192: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1193: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1194: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(418)),
	}})))),
	1195: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1196: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1197: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1198: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1199: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1200: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1201: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1202: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1203: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1204: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1205: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1206: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(357)),
	}})))),
	1207: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1208: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(408)),
	}})))),
	1209: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1210: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1211: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1212: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1213: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1214: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1215: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1216: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1217: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
	1218: *(*TSParseActionEntry)(unsafe.Pointer(&*(*TSParseAction)(unsafe.Pointer(&struct {
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
	1219: *(*TSParseActionEntry)(unsafe.Pointer(&struct {
		f struct {
			Fcount    uint8_t
			Freusable uint8
		}
		_ [6]byte
	}{f: struct {
		Fcount    uint8_t
		Freusable uint8
	}{
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
	}})),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(120)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(122)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(121)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(212)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(375)),
	}})))),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(369)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(454)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(455)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(456)),
	}})))),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(465)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(466)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(467)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(469)),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(470)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(471)),
	}})))),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(472)),
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
		Fcount:    uint8(1),
		Freusable: libc.BoolUint8(true1 != 0),
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
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(473)),
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
		f struct {
			Ftype_token uint8_t
			Fstate      TSStateId
			Fextra      uint8
			Frepetition uint8
		}
		_ [2]byte
	}{f: struct {
		Ftype_token uint8_t
		Fstate      TSStateId
		Fextra      uint8
		Frepetition uint8
	}{
		Fstate: libc.Uint16FromInt32(libc.Int32FromInt32(398)),
	}})))),
}

type ts_external_scanner_symbol_identifiers = int32

const ts_external_token_PITarget = 0
const ts_external_token__pi_content = 1
const ts_external_token_Comment = 2
const ts_external_token_CharData = 3
const ts_external_token_CData = 4
const ts_external_token_xml_DASHmodel = 5
const ts_external_token_xml_DASHstylesheet = 6
const ts_external_token__start_tag_name = 7
const ts_external_token__end_tag_name = 8
const ts_external_token__erroneous_end_name = 9
const ts_external_token_SLASH_GT = 10

var ts_external_scanner_symbol_map = [11]TSSymbol{
	0:  uint16(sym_PITarget),
	1:  uint16(sym__pi_content),
	2:  uint16(sym_Comment),
	3:  uint16(sym_CharData),
	4:  uint16(sym_CData),
	5:  uint16(anon_sym_xml_DASHmodel),
	6:  uint16(anon_sym_xml_DASHstylesheet),
	7:  uint16(sym__start_tag_name),
	8:  uint16(sym__end_tag_name),
	9:  uint16(sym__erroneous_end_name),
	10: uint16(anon_sym_SLASH_GT),
}

var ts_external_scanner_states = [11][11]uint8{
	1: {
		0:  libc.BoolUint8(true1 != 0),
		1:  libc.BoolUint8(true1 != 0),
		2:  libc.BoolUint8(true1 != 0),
		3:  libc.BoolUint8(true1 != 0),
		4:  libc.BoolUint8(true1 != 0),
		5:  libc.BoolUint8(true1 != 0),
		6:  libc.BoolUint8(true1 != 0),
		7:  libc.BoolUint8(true1 != 0),
		8:  libc.BoolUint8(true1 != 0),
		9:  libc.BoolUint8(true1 != 0),
		10: libc.BoolUint8(true1 != 0),
	},
	2: {
		2: libc.BoolUint8(true1 != 0),
	},
	3: {
		2: libc.BoolUint8(true1 != 0),
		3: libc.BoolUint8(true1 != 0),
	},
	4: {
		0: libc.BoolUint8(true1 != 0),
		5: libc.BoolUint8(true1 != 0),
		6: libc.BoolUint8(true1 != 0),
	},
	5: {
		10: libc.BoolUint8(true1 != 0),
	},
	6: {
		4: libc.BoolUint8(true1 != 0),
	},
	7: {
		7: libc.BoolUint8(true1 != 0),
	},
	8: {
		0: libc.BoolUint8(true1 != 0),
	},
	9: {
		8: libc.BoolUint8(true1 != 0),
	},
	10: {
		1: libc.BoolUint8(true1 != 0),
	},
}

func tree_sitter_xml(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&language))
}

var language = TSLanguage{
	Fabi_version:               uint32(LANGUAGE_VERSION),
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
	Fkeyword_capture_token:     uint16(sym_Name),
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
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(tree_sitter_xml_external_scanner_create)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(tree_sitter_xml_external_scanner_destroy)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(tree_sitter_xml_external_scanner_scan)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(tree_sitter_xml_external_scanner_serialize)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(tree_sitter_xml_external_scanner_deserialize)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "model\x00stylesheet\x00combined.c\x00(uint32_t)((tags)->size - 1) < (tags)->size\x00(uint32_t)(i) < (tags)->size\x00(uint32_t)(serialized_tag_count) < (tags)->size\x00end\x00Name\x00<?\x00xml\x00?>\x00standalone\x00'\x00yes\x00no\x00\"\x00<!\x00DOCTYPE\x00[\x00]\x00>\x00<\x00/>\x00</\x00]]>\x00<![\x00CDATA\x00xml-stylesheet\x00xml-model\x00PseudoAttValue_token1\x00PseudoAttValue_token2\x00ELEMENT\x00EMPTY\x00ANY\x00(\x00#PCDATA\x00|\x00)\x00*\x00?\x00+\x00,\x00ATTLIST\x00TokenizedType\x00NOTATION\x00#REQUIRED\x00#IMPLIED\x00#FIXED\x00ENTITY\x00%\x00EntityValue_token1\x00EntityValue_token2\x00NDATA\x00;\x00_S\x00Nmtoken\x00&\x00&#\x00CharRef_token1\x00&#x\x00CharRef_token2\x00SYSTEM\x00PUBLIC\x00URI\x00PubidLiteral_token1\x00PubidLiteral_token2\x00version\x00VersionNum\x00encoding\x00EncName\x00=\x00PITarget\x00_pi_content\x00Comment\x00CharData\x00CData\x00_erroneous_end_name\x00document\x00prolog\x00_Misc\x00XMLDecl\x00_SDDecl\x00doctypedecl\x00_intSubset\x00element\x00EmptyElemTag\x00Attribute\x00STag\x00ETag\x00content\x00CDSect\x00CDStart\x00StyleSheetPI\x00XmlModelPI\x00PseudoAtt\x00PseudoAttValue\x00_markupdecl\x00_DeclSep\x00elementdecl\x00contentspec\x00Mixed\x00children\x00_cp\x00_choice\x00AttlistDecl\x00AttDef\x00_AttType\x00StringType\x00_EnumeratedType\x00NotationType\x00Enumeration\x00DefaultDecl\x00_EntityDecl\x00GEDecl\x00PEDecl\x00EntityValue\x00NDataDecl\x00NotationDecl\x00PEReference\x00_Reference\x00EntityRef\x00CharRef\x00AttValue\x00ExternalID\x00PublicID\x00SystemLiteral\x00PubidLiteral\x00_VersionInfo\x00_EncodingDecl\x00PI\x00_Eq\x00document_repeat1\x00EmptyElemTag_repeat1\x00content_repeat1\x00StyleSheetPI_repeat1\x00PseudoAttValue_repeat1\x00PseudoAttValue_repeat2\x00Mixed_repeat1\x00Mixed_repeat2\x00_choice_repeat1\x00_choice_repeat2\x00AttlistDecl_repeat1\x00NotationType_repeat1\x00Enumeration_repeat1\x00EntityValue_repeat1\x00EntityValue_repeat2\x00root\x00"
