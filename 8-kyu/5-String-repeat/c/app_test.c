// https://github.com/libcheck/check
#include <stdlib.h>
#include <stdint.h>
#include <check.h>
#include <stdio.h>
#include "app.h"

START_TEST(test_String_repeat)
{
    char *input = "test";
    int count = 3;
    char *actual = String_repeat(count, input);
    char *expected = "testtesttest";

    ck_assert_str_eq(actual, expected);
}
END_TEST

struct TestCase
{
    int count;
    char input[50];
    char expected[50];
};

START_TEST(test_String_repeat_2)
{

    struct TestCase testcases[] = {
        {4, "a", "aaaa"},
        {3, "hello ", "hello hello hello "},
        {2, "abc", "abcabc"}};

    size_t count = sizeof(testcases) / sizeof(testcases[0]);
    for (int i = 0; i < count; i++)
    {
        char *actual = String_repeat(testcases[i].count, testcases[i].input);
        ck_assert_str_eq(actual, testcases[i].expected);
    }
}
END_TEST

START_TEST(test_String_repeat_2_B)
{

    struct TestCase testcases[] = {
        {4, "a", "aaaa"},
        {3, "hello ", "hello hello hello "},
        {2, "abc", "abcabc"}};

    size_t count = sizeof(testcases) / sizeof(testcases[0]);
    for (int i = 0; i < count; i++)
    {
        char *actual = String_repeat_B(testcases[i].count, testcases[i].input);
        ck_assert_str_eq(actual, testcases[i].expected);
    }
}
END_TEST

START_TEST(test_String_repeat_2_C)
{

    struct TestCase testcases[] = {
        {4, "a", "aaaa"},
        {3, "hello ", "hello hello hello "},
        {2, "abc", "abcabc"}};

    size_t count = sizeof(testcases) / sizeof(testcases[0]);
    for (int i = 0; i < count; i++)
    {
        char *actual = String_repeat_C(testcases[i].count, testcases[i].input);
        ck_assert_str_eq(actual, testcases[i].expected);
    }
}
END_TEST

Suite *String_repeat_suite(void)
{
    Suite *s;
    TCase *tc_core;

    s = suite_create("String_repeat");

    /* Core test case */
    tc_core = tcase_create("Core");

    tcase_add_test(tc_core, test_String_repeat);
    tcase_add_test(tc_core, test_String_repeat_2);
    tcase_add_test(tc_core, test_String_repeat_2_B);
    tcase_add_test(tc_core, test_String_repeat_2_C);

    suite_add_tcase(s, tc_core);

    return (s);
}

int main(void)
{
    int number_failed;
    Suite *s;
    SRunner *sr;

    s = String_repeat_suite();
    sr = srunner_create(s);

    srunner_run_all(sr, CK_NORMAL);
    number_failed = srunner_ntests_failed(sr);
    srunner_free(sr);
    return (number_failed == 0) ? EXIT_SUCCESS : EXIT_FAILURE;
}

// How to build.
// gcc -c String_repeat.c -o String_repeat.out
// gcc -c String_repeat_test.c -o String_repeat_test.out
// gcc String_repeat.out String_repeat_test.out -o test.out `pkg-config --cflags --libs check`
// ./test.out
