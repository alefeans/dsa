import pytest

from algs4.part1.week3.sorting import merge_sort, not_in_place_quick_sort


@pytest.mark.parametrize(
    "input, expected",
    [
        ([], []),
        ([0], [0]),
        ([2, 1, 0], [0, 1, 2]),
        ([3, 1, 2, 5, 0, 7, 6, 4], [0, 1, 2, 3, 4, 5, 6, 7]),
    ],
)
def test_if_merge_sort_returns_sorted_list_of_numbers(input, expected):
    assert merge_sort(input) == expected


@pytest.mark.parametrize(
    "input, expected",
    [
        ([], []),
        ([0], [0]),
        ([2, 1, 0], [0, 1, 2]),
        ([3, 1, 2, 5, 0, 7, 6, 4], [0, 1, 2, 3, 4, 5, 6, 7]),
    ],
)
def test_if_not_in_place_quick_sort_returns_sorted_list_of_numbers(input, expected):
    assert not_in_place_quick_sort(input) == expected
