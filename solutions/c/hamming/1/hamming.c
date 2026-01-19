#include "hamming.h"
#include <string.h>

int compute(const char *lhs, const char *rhs){
    int l = strlen(lhs), r = strlen(rhs);
    int dist = 0;
    if (l!=r) return -1;
    for(int i=0; i < l; i++){
        if (lhs[i] != rhs[i]) dist++;
    }
    return dist;
}