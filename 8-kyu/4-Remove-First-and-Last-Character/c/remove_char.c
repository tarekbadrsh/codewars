#include <string.h>
#include <stdlib.h>

#include "remove_char.h"

char *remove_char(char* dst, const char* src) {
    long ln = strlen(src);
    dst = malloc(ln);
    for (long i = 1; i < ln - 1; i++)
    {
        dst[i-1] = src[i];
    }
    dst[ln] = '\0';
    return dst;
}
