// https://github.com/libcheck/check
#include <stdlib.h>
#include <stdint.h>
#include <check.h>
#include "positive_sum.h"

START_TEST(test_positive_sum)
{
    const int values[] = {1, 2, 3, 4, 5};
    int res = positive_sum(values, sizeof(values) / sizeof(values[0]));

    ck_assert_int_eq(positive_sum(values, sizeof(values) / sizeof(values[0])), 15);
}
END_TEST

START_TEST(test_positive_sum_2)
{
    const int values[] = {1, 2, 3, 4, 5 , 6};
    int res = positive_sum(values, sizeof(values) / sizeof(values[0]));

    ck_assert_int_eq(positive_sum(values, sizeof(values) / sizeof(values[0])), 21);
}
END_TEST

Suite *positive_sum_suite(void)
{
    Suite *s;
    TCase *tc_core;

    s = suite_create("positive_sum");

    /* Core test case */
    tc_core = tcase_create("Core");

    tcase_add_test(tc_core, test_positive_sum);
    tcase_add_test(tc_core, test_positive_sum_2);

    suite_add_tcase(s, tc_core);

    return (s);
}



int main(void)
{
    int number_failed;
    Suite *s;
    SRunner *sr;

    s = positive_sum_suite();
    sr = srunner_create(s);

    srunner_run_all(sr, CK_NORMAL);
    number_failed = srunner_ntests_failed(sr);
    srunner_free(sr);
    return (number_failed == 0) ? EXIT_SUCCESS : EXIT_FAILURE;
}


// How to build.
// gcc -c positive_sum.c -o positive_sum.out
// gcc -c positive_sum_test.c -o positive_sum_test.out
// gcc positive_sum.out positive_sum_test.out -o test.out `pkg-config --cflags --libs check`

