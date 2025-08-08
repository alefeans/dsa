from algs4.part1.week2.stack import Stack


class Eval:
    """Very simple implementation of Dijkstra Arithmetic Expression Evaluator"""

    def __init__(self) -> None:
        self.operators = Stack()
        self.operands = Stack()
        self.operators_map = {
            "+": self._sum,
            "-": self._sub,
            "*": self._mul,
            "/": self._div,
        }

    def arithmetic_expression(self, expression: str) -> int:
        for char in expression:
            if char in ("", " ", "("):
                continue
            elif char in ("+", "-", "*", "/"):
                self.operators.push(self.operators_map[char])
            elif char == ")":
                op = self.operators.pop()
                self.operands.push(op(self.operands.pop(), self.operands.pop()))
            else:
                self.operands.push(int(char))

        result: int = self.operands.pop()
        return result

    @staticmethod
    def _sum(a: int, b: int) -> int:
        return a + b

    @staticmethod
    def _sub(a: int, b: int) -> int:
        return b - a

    @staticmethod
    def _mul(a: int, b: int) -> int:
        return b * a

    @staticmethod
    def _div(a: int, b: int) -> int:
        return b // a
