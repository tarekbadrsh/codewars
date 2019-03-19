#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "app.h"

int main()
{
    const char *str = "hello ";
    size_t count = 3;
    printf("%s\n", String_repeat(count, str));

    return 0;
}

// How to build.
// gcc -c remove_char.c -o remove_char.out
// gcc -c main.c -o main.out
// gcc remove_char.out main.out
// ./a.out

