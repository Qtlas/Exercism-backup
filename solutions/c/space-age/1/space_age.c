#include "space_age.h"

float age(planet_t planet, int64_t seconds) {
    float year = seconds / 60 / 60 / 24 / 365;
    switch(planet) {
        case MERCURY:
            return year / 0.2408467;
        case VENUS:
            return year / 0.61519726;
        case MARS:
            return year / 1.8808158;
        case JUPITER:
            return year / 11.862615;
        case SATURN:
            return year / 29.447498;
        case URANUS:
            return year / 84.016846;
        case NEPTUNE:
            return year / 164.79132;
        case EARTH:
            return year;
    }
    return -1;
}