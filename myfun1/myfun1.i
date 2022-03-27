
%module myfun1

%{
#include "myfun1.h"
%}

/* Let's just grab the original header file here */
%include "myfun1.h"