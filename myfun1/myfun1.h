#ifndef __MYFUN1_H
#define __MYFUN1_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>

#pragma pack(1)
typedef struct {
    uint32_t a;
    uint8_t b;
    uint16_t c;
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
void SerialStructS11Ptr(uint8_t *out, uint32_t *len, void *p);

#ifdef __cplusplus
}
#endif

#endif