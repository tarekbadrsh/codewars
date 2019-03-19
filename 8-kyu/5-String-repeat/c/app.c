#include <string.h>
#include <stdlib.h>

#include <stdio.h>

#include "app.h"

char *String_repeat(size_t count, const char *src)
{
    int ln = strlen(src);
    int result_ln = ln * count;
    char *result = malloc(result_ln);

    int currantChar = 0;
    for (int x = 0; x < count; x++)
    {
        for (int i = 0; i < ln; i++)
        {
            result[currantChar] = src[i];
            currantChar++;
        }
    }
    result[result_ln] = '\0';
    return result;
}

char *String_repeat_B(size_t count, const char *src)
{
    char *result = calloc(((count * strlen(src)) + 1), sizeof(char));
    while (count--)
        strcat(result, src);
    return result;
}

char *String_repeat_C(size_t count, const char *src)
{
    int length = strlen(src);
    char *dest = malloc(count * length * sizeof(char));
    for (int i = 0; i < count; i++)
    {
        strcpy(dest + i * length, src);
    }
    return dest;
}