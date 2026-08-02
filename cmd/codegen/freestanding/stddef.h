#ifndef TS_FS_STDDEF_H
#define TS_FS_STDDEF_H

#define NULL ((void *)0)

#if defined(__SIZEOF_POINTER__) && __SIZEOF_POINTER__ == 4
typedef unsigned int size_t;
typedef int ptrdiff_t;
#else
typedef unsigned long long size_t;
typedef long long ptrdiff_t;
#endif

typedef int wchar_t;

#endif
