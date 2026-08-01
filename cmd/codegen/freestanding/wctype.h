#ifndef TS_FS_WCTYPE_H
#define TS_FS_WCTYPE_H

#include "stddef.h"

typedef unsigned int wint_t;
typedef unsigned long wctype_t;

int iswalnum(wint_t wc);
int iswalpha(wint_t wc);
int iswdigit(wint_t wc);
int iswspace(wint_t wc);
int iswxdigit(wint_t wc);
int iswprint(wint_t wc);
int iswupper(wint_t wc);
int iswlower(wint_t wc);
wint_t towupper(wint_t wc);
wint_t towlower(wint_t wc);

#endif
