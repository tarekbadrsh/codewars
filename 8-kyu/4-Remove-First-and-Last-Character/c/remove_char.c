#include <string.h>
#include <stdlib.h>

#include "remove_char.h"

char *remove_char(char *dst, const char *src)
{
    int ln = strlen(src) - 2;
    if (ln <= 0)
    {
        return "";
    }
    dst = malloc((size_t)ln);
    for (int i = 0; i < ln; i++)
    {
        dst[i] = src[i + 1];
    }
    dst[ln] = '\0';
    return dst;
}

// work only with Clang compiler
char *remove_char_B(char *dst, const char *src)
{
    strcpy(dst, src + 1);
    dst[strlen(dst) - 1] = '\0';
    return dst;
}