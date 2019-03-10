// https://github.com/libcheck/check
#include <stdlib.h>
#include <stdint.h>
#include <check.h>
#include "dna_to_rna.h"

#include <stdio.h>

START_TEST(test_dna_to_rna)
{
    char *input = "123T456789xTTT";
    char *actual = dna_to_rna(input);
    char *expected = "123U456789xUUU";

    ck_assert_str_eq(actual, expected);
}
END_TEST

struct TestCase
{
    char input[50];
    char expected[50];
};

START_TEST(test_dna_to_rna_2)
{
    struct TestCase testcases[7] = {
        {"", ""},
        {"abc", "abc"},
        {"abcT", "abcU"},
        {"T", "U"},
        {"TTT", "UUU"},
        {"123T456789xTTT", "123U456789xUUU"},
        {"BTentlTey", "BUentlUey"}};
    for (int i = 0; i < 7; i++)
    {
        char *actual = dna_to_rna(testcases[i].input);
        ck_assert_str_eq(actual, testcases[i].expected);
    }
}
END_TEST

START_TEST(test_dna_to_rna_B)
{
    struct TestCase testcases[7] = {
        {"", ""},
        {"abc", "abc"},
        {"abcT", "abcU"},
        {"T", "U"},
        {"TTT", "UUU"},
        {"123T456789xTTT", "123U456789xUUU"},
        {"BTentlTey", "BUentlUey"}};
    for (int i = 0; i < 7; i++)
    {
        char *actual = dna_to_rna(testcases[i].input);
        ck_assert_str_eq(actual, testcases[i].expected);
    }
}
END_TEST

Suite *dna_to_rna_suite(void)
{
    Suite *s;
    TCase *tc_core;

    s = suite_create("dna_to_rna");

    /* Core test case */
    tc_core = tcase_create("Core");

    tcase_add_test(tc_core, test_dna_to_rna);
    tcase_add_test(tc_core, test_dna_to_rna_2);
    tcase_add_test(tc_core, test_dna_to_rna_B);

    suite_add_tcase(s, tc_core);

    return (s);
}

int main(void)
{
    int number_failed;
    Suite *s;
    SRunner *sr;

    s = dna_to_rna_suite();
    sr = srunner_create(s);

    srunner_run_all(sr, CK_NORMAL);
    number_failed = srunner_ntests_failed(sr);
    srunner_free(sr);
    return (number_failed == 0) ? EXIT_SUCCESS : EXIT_FAILURE;
}

// How to build.
// gcc -c dna_to_rna.c -o dna_to_rna.out
// gcc -c dna_to_rna_test.c -o dna_to_rna_test.out
// gcc dna_to_rna.out dna_to_rna_test.out -o test.out `pkg-config --cflags --libs check`
// ./test.out
