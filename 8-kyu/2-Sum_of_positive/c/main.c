#include <stddef.h>

#include <stdio.h>

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

int main()
{
  const int values[] = {1, 2, 3, 4, 5, 6};
  int res = positive_sum(values, sizeof(values) / sizeof(values[0]));

  printf("result : %d\n", res);
  return 0;
}