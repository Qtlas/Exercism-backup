#include "binary.h"
#include <string.h>
#include <math.h>

int convert(const char *input){
    int len = strlen(input);
    int result = 0;
    for(int i=len-1; i>=0;i--){
        if(input[i] == '1') {
            result += pow(2, len-i-1);
        } else if (input[i] != '0') return INVALID;
    }
    return result;
}