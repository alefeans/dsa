from typing import Any, Self


class Node:
    def __init__(self, value: Any) -> None:
        self.value = value
        self.next: Node | None = None

    def __str__(self) -> str:
        return f"{self.value} -> {self.next}"


class Queue:
    def __init__(self) -> None:
        self.first: Node | Any = None
        self.last: Node | Any = None
        self.size = 0

    def is_empty(self) -> bool:
        return self.size == 0

    def enqueue(self, value: Any) -> None:
        old_last = self.last
        self.last = Node(value)
        if self.is_empty():
            self.first = self.last
        else:
            old_last.next = self.last
        self.size += 1

    def dequeue(self) -> Any:
        if self.is_empty():
            return None
        self.size -= 1
        value = self.first.value
        self.first = self.first.next
        return value

    def __len__(self) -> int:
        return self.size

    def __str__(self) -> str:
        if self.is_empty():
            return "Empty queue"
        return f"{self.first}"

    def __iter__(self) -> Self:
        return self

    def __next__(self) -> Any:
        if self.is_empty():
            raise StopIteration
        return self.dequeue()
