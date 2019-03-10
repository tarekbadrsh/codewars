#include <stdio.h>
#include "dna_to_rna.h"

int main()
{
    char *input = "123T456789xTTT";
    printf("%s\n", dna_to_rna(input));

    return 0;
}

// How to build.
// gcc -c dna_to_rna.c -o dna_to_rna.out
// gcc -c main.c -o main.out
// gcc dna_to_rna.out main.out
// ./a.out
