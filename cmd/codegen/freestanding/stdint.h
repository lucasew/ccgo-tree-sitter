/* Freestanding stubs for grammar ccgo — no host libc. */
#ifndef TS_FS_STDINT_H
#define TS_FS_STDINT_H

typedef signed char int8_t;
typedef short int16_t;
typedef int int32_t;
typedef long long int64_t;
typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned int uint32_t;
typedef unsigned long long uint64_t;

typedef int64_t intmax_t;
typedef uint64_t uintmax_t;

#if defined(__SIZEOF_POINTER__) && __SIZEOF_POINTER__ == 4
typedef int32_t intptr_t;
typedef uint32_t uintptr_t;
#else
typedef int64_t intptr_t;
typedef uint64_t uintptr_t;
#endif

#define INT8_MAX 0x7f
#define INT16_MAX 0x7fff
#define INT32_MAX 0x7fffffff
#define INT64_MAX 0x7fffffffffffffffLL
#define UINT8_MAX 0xff
#define UINT16_MAX 0xffff
#define UINT32_MAX 0xffffffffu
#define UINT64_MAX 0xffffffffffffffffull

#endif
