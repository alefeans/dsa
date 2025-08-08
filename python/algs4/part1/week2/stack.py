from typing import Any, Self


class Node:
    def __init__(self, value: Any) -> None:
        self.value = value
        self.next: Node | None = None

    def __str__(self) -> str:
        return f"{self.value} -> {self.next}"


class Stack:
    def __init__(self) -> None:
        self.first: Node | None = None
        self.size = 0
        self.max = 0

    def push(self, value: Any) -> None:
        self.size += 1
        if self.first is None:
            self.first = Node(value)
            return
        old = self.first
        self.first = Node(value)
        self.first.next = old

    def pop(self) -> Any:
        if self.first is None:
            return
        self.size -= 1
        value = self.first.value
        self.first = self.first.next
        return value

    def is_empty(self) -> bool:
        return self.size == 0

    def __len__(self) -> int:
        return self.size

    def __str__(self) -> str:
        if self.first is None:
            return "Empty stack"
        return f"{self.first}"

    def __iter__(self) -> Self:
        return self

    def __next__(self) -> Any:
        if self.is_empty():
            raise StopIteration
        return self.pop()
