#ifndef TS_FS_STDIO_H
#define TS_FS_STDIO_H

#include "stddef.h"

typedef struct ts_fs_FILE {
	int _unused;
} FILE;

extern FILE *stderr;
extern FILE *stdout;
extern FILE *stdin;

int printf(const char *fmt, ...);
int fprintf(FILE *stream, const char *fmt, ...);
int sprintf(char *str, const char *fmt, ...);
int snprintf(char *str, size_t size, const char *fmt, ...);
int puts(const char *s);
int fputs(const char *s, FILE *stream);
void perror(const char *s);

#endif
