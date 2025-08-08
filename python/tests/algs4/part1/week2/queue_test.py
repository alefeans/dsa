import pytest

from algs4.part1.week2.queue import Queue


@pytest.mark.parametrize(
    "input, expected",
    [
        ([], []),
        ([0], [0]),
        ([0, 1, 2], [0, 1, 2]),
        ([1, 0, 3], [1, 0, 3]),
    ],
)
def test_if_returns_items_in_correct_fifo_order(input, expected):
    q = Queue()
    for i in input:
        q.enqueue(i)

    for e in expected:
        assert e == q.dequeue()
