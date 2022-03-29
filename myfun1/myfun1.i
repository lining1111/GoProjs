
%module myfun1

%{
#include "myfun1.h"
%}

/* Let's just grab the original header file here */
%include "stdint.i"
%include "myfun1.h"