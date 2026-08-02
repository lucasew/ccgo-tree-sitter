#ifndef TS_FS_STDLIB_H
#define TS_FS_STDLIB_H

#include "stddef.h"

void *malloc(size_t size);
void *calloc(size_t nmemb, size_t size);
void *realloc(void *ptr, size_t size);
void free(void *ptr);
void abort(void);
void exit(int status);
int abs(int j);
long labs(long j);
int atoi(const char *nptr);

#define EXIT_SUCCESS 0
#define EXIT_FAILURE 1

#endif
