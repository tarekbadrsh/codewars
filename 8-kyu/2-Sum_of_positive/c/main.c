#include <stdio.h>

int main()
{
  const int values[] = {1, 2, 3, 4, 5, 6};
  int res = positive_sum(values, sizeof(values) / sizeof(values[0]));

  printf("result : %d\n", res);
  return 0;
}


// How to build.
// gcc -c positive_sum.c -o positive_sum.out
// gcc -c main.c -o main.out
// gcc positive_sum.out main.out
// ./a.out

