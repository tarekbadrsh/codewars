// https://github.com/libcheck/check
#include <stdlib.h>
#include <stdint.h>
#include <check.h>
#include "remove_char.h"

START_TEST(test_remove_char)
{
    char *input = "12345";
    char *actual = remove_char("", input);
    char *expected = "234";

    ck_assert_str_eq(actual, expected);
}
END_TEST

struct TestCase
{
    char input[50];
    char expected[50];
};

START_TEST(test_remove_char_2)
{
    char *strings[] = {"eloquent", "country", "person", "place", "ok"};
    char *expected[] = {"loquen", "ountr", "erso", "lac", ""};

    struct TestCase testcases[7] = {
        {"", ""},
        {"eloquent", "loquen"},
        {"country", "ountr"},
        {"person", "erso"},
        {"place", "lac"},
        {"ok", ""},
        {"test", "es"}};
    for (int i = 0; i < 7; i++)
    {
        char *actual = remove_char("", testcases[i].input);
        ck_assert_str_eq(actual, testcases[i].expected);
    }
}
END_TEST

Suite *remove_char_suite(void)
{
    Suite *s;
    TCase *tc_core;

    s = suite_create("remove_char");

    /* Core test case */
    tc_core = tcase_create("Core");

    tcase_add_test(tc_core, test_remove_char);
    tcase_add_test(tc_core, test_remove_char_2);

    suite_add_tcase(s, tc_core);

    return (s);
}

int main(void)
{
    int number_failed;
    Suite *s;
    SRunner *sr;

    s = remove_char_suite();
    sr = srunner_create(s);

    srunner_run_all(sr, CK_NORMAL);
    number_failed = srunner_ntests_failed(sr);
    srunner_free(sr);
    return (number_failed == 0) ? EXIT_SUCCESS : EXIT_FAILURE;
}

// How to build.
// gcc -c remove_char.c -o remove_char.out
// gcc -c remove_char_test.c -o remove_char_test.out
// gcc remove_char.out remove_char_test.out -o test.out `pkg-config --cflags --libs check`
// ./test.out
