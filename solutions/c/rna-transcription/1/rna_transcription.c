#include "rna_transcription.h"
#include <string.h>
#include <stdlib.h>

char *to_rna(const char *dna){
    char *c = malloc(strlen(dna) + 1);
    int len = strlen(dna);
    for(int i=0; i < len; i++){
        if(dna[i] == 'G') c[i] = 'C';
        else if (dna[i] == 'C') c[i] = 'G';
        else if (dna[i] == 'T') c[i] = 'A';
        else if (dna[i] == 'A') c[i] = 'U';
    }
    c[len] = '\0';
    return c;
}