#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *dna_to_rna(const char *dna)
{
    char nchar = 'U';
    long ln =  strlen(dna);
    char *result = malloc(ln + 1);
    for (long i = 0; i < ln ; i++)
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

int main()
{
    char *input = "123456789xTTT";
    // char *input = "GGCTTAGCGATTGGATCGACGGCC";
    printf("%s\n", dna_to_rna(input));

    return 0;
}