//
// Created by lining on 3/31/22.
//

#ifndef _EXAMPLE_H
#define _EXAMPLE_H

#include <cstdint>

typedef struct {
   uint8_t a;
   uint16_t b;
   uint32_t c;
}S1;

typedef struct {
    uint8_t len;
    S1 s1;
}S11;

int SetS11(S11 *s11);

#endif //_EXAMPLE_H
