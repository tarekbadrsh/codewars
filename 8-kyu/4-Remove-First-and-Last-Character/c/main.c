#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "remove_char.h"

int main()
{
    char *input = "12345";
    printf("%s\n", remove_char("", input));

    return 0;
}

// How to build.
// gcc -c remove_char.c -o remove_char.out
// gcc -c main.c -o main.out
// gcc remove_char.out main.out
// ./a.out

