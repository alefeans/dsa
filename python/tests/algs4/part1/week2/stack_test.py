import pytest

from algs4.part1.week2.stack import Stack


@pytest.mark.parametrize(
    "input, expected",
    [
        ([], []),
        ([0], [0]),
        ([0, 1, 2], [2, 1, 0]),
        ([2, 1, 0], [0, 1, 2]),
    ],
)
def test_if_returns_items_in_correct_lifo_order(input, expected):
    s = Stack()
    for i in input:
        s.push(i)

    for e in expected:
        assert e == s.pop()
