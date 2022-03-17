#ifndef __MYFUN1_H
#define __MYFUN1_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

#pragma pack(1)
typedef struct{
    int a;
    char b;
}S11;
#pragma pack()
int setStruct1(S11 *s);


#ifdef __cplusplus
}
#endif

#endif