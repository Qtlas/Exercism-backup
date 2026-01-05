#include "darts.h"
#include <math.h>

int score(coordinate_t coord) {
    float x = coord.x;
    float y = coord.y;
    float dist = sqrt(x*x + y*y);
    if (dist > 10) {
        return 0;
    } else if (dist > 5) {
        return 1;
    } else if (dist > 1) {
        return 5;
    } else return 10;
}