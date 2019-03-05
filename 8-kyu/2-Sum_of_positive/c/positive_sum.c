#include <stddef.h>

#include "positive_sum.h"

int positive_sum(const int *values, size_t count)
{
  int result = 0;
  for (size_t i = 0; i < count; ++i)
  {
    if (values[i] > 0)
    {
      result += values[i];
    }
  }
  return result;
}


// How to build.
// gcc -c positive_sum.c -o positive_sum