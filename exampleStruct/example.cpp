//
// Created by lining on 3/31/22.
//

#include "example.h"

int SetS11(S11 *s11) {
    s11->len = sizeof(s11->s1);
    s11->s1.a = sizeof(uint8_t);
    s11->s1.b = sizeof(uint16_t);
    s11->s1.c = sizeof(uint32_t);

    return 0;
}
