#ifndef __MYFUN1_H
#define __MYFUN1_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>
#include <stdlib.h>
#include <unistd.h>

#pragma pack(1)
typedef struct {
    int a;
    char b;
    short c;
} S11;
#pragma pack()

typedef struct {
    int a;
    char b;
    short c;
}CGO_S11;

void setStructS11(S11 *s);
void setStructCGO_S11(CGO_S11 *s);

void SerialStructS11(uint8_t *out, uint32_t *len, S11 *s);

#ifdef __cplusplus
}
#endif

#endif