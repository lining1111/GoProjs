#ifndef __MYFUN_H
#define __MYFUN_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

#pragma pack(1)
typedef struct{
    int a;
    char b;
}S1;
#pragma pack()

int setStruct(S1 *s1);


#ifdef __cplusplus
}
#endif
#endif