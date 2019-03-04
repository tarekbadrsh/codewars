// https://stackoverflow.com/questions/65820/unit-testing-c-code
// https://libcheck.github.io/check/
//https://github.com/libcheck/check

// #include <criterion/criterion.h>
// #include <stddef.h>

// int positive_sum(const int *values, size_t count);

// Test(CoreTests, ShouldPassAllTheTestsProvided) {
//   {
//     const int values[] = { 1,2,3,4,5 };
//     int received = positive_sum(values, sizeof(values)/sizeof(values[0]));
//     int expected = 15;
//     cr_assert_eq(expected, received, "Expected: %d Received: %d", expected, received);
//   }
  
//   {
//     const int values[] = { 1,-2,3,4,5 };
//     int received = positive_sum(values, sizeof(values)/sizeof(values[0]));
//     int expected = 13;
//     cr_assert_eq(expected, received, "Expected: %d Received: %d", expected, received);
//   }
  
//   {
//     const int values[] = { -1,2,3,4,-5 };
//     int received = positive_sum(values, sizeof(values)/sizeof(values[0]));
//     int expected = 9;
//     cr_assert_eq(expected, received, "Expected: %d Received: %d", expected, received);
//   }
  
//   {
//     const int values[] = { -1,-2,-3,-4,-5 };
//     int received = positive_sum(values, sizeof(values)/sizeof(values[0]));
//     int expected = 0;
//     cr_assert_eq(expected, received, "Expected: %d Received: %d", expected, received);
//   }
// }