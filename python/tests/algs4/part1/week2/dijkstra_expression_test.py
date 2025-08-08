import pytest

from algs4.part1.week2.dijkstra_expression import Eval


@pytest.mark.parametrize(
    "input,expected",
    [
        ("(3 + 5)", 8),
        ("(4 - 2)", 2),
        ("(6 * 9)", 54),
        ("(9 / 3)", 3),
        ("(1 + (3 - 2))", 2),
        ("(4 + (5 - (3 + 2)))", 4),
        ("(4 + (5 - (3 + (2)))", 4),
    ],
)
def test_if_returns_correct_result_for_arithmetic_expression(input, expected):
    assert Eval().arithmetic_expression(input) == expected
