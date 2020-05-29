from typing import Iterable, List, Union


class Ind(List[int]):
    """Ind represents the indeterminates of a single term."""

    def __init__(self, *args: int) -> None:
        """Initialise with a list of indeterminates."""
        super().__init__(args)

    def __mul__(self, other: Union['Ind', int]) -> 'Ind':
        """Returns the product of self and other."""
        if not isinstance(other, type(self)):
            return NotImplemented
        if len(self) < len(other):
            self, other = other, self

        ret = self.copy()
        for i, _ in enumerate(self):
            if i < len(other):
                ret[i] += other[i]

        return ret

    def __eq__(self, other: 'Ind') -> bool:
        """Compares two sets of indeterminates for equality."""
        if not isinstance(other, type(self)):
            return False
        if len(self) != len(other):
            return False
        for a, b in zip(self, other):
            if a != b:
                return False
        return True

    def __repr__(self) -> None:
        """A compact, human-readable representation of the indeterminates."""
        s: List[str]
        simple = 'x', 'y', 'z'

        if not self:
            return "1"
        if len(self) <= len(simple):
            s = simple[:len(self)]
        else:
            s = ['x'+self._sub(i) for i, _ in enumerate(self)]

        ret = ''
        for i, x in enumerate(self):
            if x:
                ret += s[i]
                if x != 1:
                    ret += self._sup(x)
        if not ret:
            return '1'

        return ret

    @classmethod
    def _sub(cls, x: int) -> str:
        """Turn x into a Unicode subscript."""
        if not x:
            return '₀'
        return cls._smap(x, '₀₁₂₃₄₅₆₇₈₉₋')

    @classmethod
    def _sup(cls, x: int) -> str:
        """Turn x into a Unicode superscript."""
        return cls._smap(x, '⁰¹²³⁴⁵⁶⁷⁸⁹¯')

    @staticmethod
    def _smap(x: int, chars: str) -> str:
        """Turn x into a Unicode subscript or superscript string."""
        abs_x, s = abs(x), ''
        while abs_x:
            s = chars[abs_x%10] + s
            abs_x //= 10
        if x < 0:
            s = chars[10] + s
        return s


class IntT(Ind):
    """IntT is a single term containing an int coefficient."""
    c: int

    def __init__(self, c: int = 0, ind: Iterable[int] = ()) -> None:
        """Initialise with a constant and optional terms."""
        super().__init__(*ind)
        self.c = c

    def __mul__(self, other: Union['IntT', int]) -> 'IntT':
        """Returns the product of self and other."""
        if not isinstance(other, (type(self), int)):
            return NotImplemented

        if isinstance(other, int):
            return self.__class__(self.c*other, self)

        return self.__class__(self.c*other.c, super().__mul__(other))

    def __eq__(self, other: 'IntT') -> bool:
        if not isinstance(other, type(self)):
            return False
        return (self.c == other.c) and super().__eq__(other)

    def __gt__(self, other: 'IntT') -> bool:
        """Compares two terms."""
        if not isinstance(other, type(self)):
            raise TypeError(
                "'>' not supported between instances of '{}' and '{}'"
                .format(type(self), type(other)))
        for a, b in zip(self, other):
            if a == b:
                continue
            return a > b
        return False

    def __repr__(self) -> str:
        """A compact, human-readable representation of the term."""
        if not self.c:
            return "0"

        s = super().__repr__()
        if self.c == 1:
            return s
        if s == '1':
            return str(self.c)

        return '{:d}{}'.format(self.c, s)