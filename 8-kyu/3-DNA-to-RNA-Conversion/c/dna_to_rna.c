#include <stdlib.h>
#include <string.h>

#include "dna_to_rna.h"

char *dna_to_rna(const char *dna)
{
    long ln = strlen(dna);
    char *result = malloc(ln + 1);
    for (long i = 0; i < ln; i++)
    {
        result[i] = dna[i];
        if (dna[i] == 'T')
        {
            result[i] = 'U';
        }
    }
    result[ln] = '\0';
    return result;
}

// another solution
char *dna_to_rna_B(const char *dna)
{
    int i;
    i = 0;
    char *rna;
    rna = (char *)malloc(sizeof(char) * (strlen(dna)));
    while (dna[i] != '\0')
    {
        if (dna[i] == 'T')
        {
            rna[i] = 'U';
        }
        else
        {
            rna[i] = dna[i];
        }
        i++;
    }
    rna[i] = '\0';
    return rna;
    free(rna);
}

// How to build.
// gcc -c dna_to_rna.c -o dna_to_rna