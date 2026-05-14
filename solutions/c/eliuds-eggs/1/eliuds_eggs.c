#include "eliuds_eggs.h"


int egg_count(int n) {
    if (n==0) return 0;
    return n%2 + egg_count(n/2);
}