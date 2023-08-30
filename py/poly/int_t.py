"""Module int_t impolements integer indeterminates for polynomials."""
import json
from typing import Iterable, Union


class Ind(list[int]):
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
            raise TypeError(
                "'==' not supported between instances of '{}' and '{}'"
                .format(type(self), type(other)))
        if len(self) != len(other):
            return False
        for ind_a, ind_b in zip(self, other):
            if ind_a != ind_b:
                return False
        return True

    def __repr__(self) -> None:
        """A compact, human-readable representation of the indeterminates."""
        parts: list[str]
        simple = 'x', 'y', 'z'

        if not self:
            return "1"
        if len(self) <= len(simple):
            parts = simple[:len(self)]
        else:
            parts = ['x'+self._sub(i) for i, _ in enumerate(self)]

        ret = ''
        for i, ind in enumerate(self):
            if ind:
                ret += parts[i]
                if ind != 1:
                    ret += self._sup(ind)
        if not ret:
            return '1'

        return ret

    @classmethod
    def from_json(cls, json_data: Union[str, bytes, bytearray]) -> 'Ind':
        """Decode indeterminates from JSON.

        Arguments:
            json_data: JSON-encoded indeterminates. Must be an array of
                       integers.
        Returns:
            The decoded object.
        """
        return cls.from_data(json.loads(json_data))

    @staticmethod
    def from_data(data: list[int]) -> 'Ind':
        if not isinstance(data, list):
            raise ValueError('Expected a list of integers, got: {}.'
                             .format(type(data)))
        for i, item in enumerate(data, 1):
            if not isinstance(item, int):
                raise ValueError('Expected integer at item {}, got: {}.'
                                 .format(i, type(item)))
        # Explicitly spell out Ind, so subclasses can call super().from_data().
        return Ind(*data)

    def to_json(self) -> str:
        return json.dumps(self)

    @classmethod
    def _sub(cls, ind: int) -> str:
        """Turn ind into a Unicode subscript."""
        if not ind:
            return '₀'
        return cls._smap(ind, '₀₁₂₃₄₅₆₇₈₉₋')

    @classmethod
    def _sup(cls, ind: int) -> str:
        """Turn ind into a Unicode superscript."""
        return cls._smap(ind, '⁰¹²³⁴⁵⁶⁷⁸⁹⁻')

    @staticmethod
    def _smap(ind: int, chars: str) -> str:
        """Turn ind into a Unicode subscript or superscript string."""
        abs_x, ret = abs(ind), ''
        while abs_x:
            ret = chars[abs_x % 10] + ret
            abs_x //= 10
        if ind < 0:
            ret = chars[10] + ret
        return ret


class IntT(Ind):
    """IntT is a single term containing an int coefficient."""
    const: int

    def __init__(self, const: int = 0, ind: Iterable[int] = ()) -> None:
        """Initialise with a constant and optional terms."""
        super().__init__(*ind)
        self.const = const

    def __mul__(self, other: Union['IntT', int]) -> 'IntT':
        """Returns the product of self and other."""
        if not isinstance(other, (type(self), int)):
            return NotImplemented

        if isinstance(other, int):
            return self.__class__(self.const*other, self)

        return self.__class__(self.const*other.const, super().__mul__(other))

    def __eq__(self, other: 'IntT') -> bool:
        """Compares two terms for equality."""
        if not isinstance(other, type(self)):
            raise TypeError(
                "'==' not supported between instances of '{}' and '{}'"
                .format(type(self), type(other)))
        return self.ind_eq(other) and (self.const == other.const)

    def __gt__(self, other: 'IntT') -> bool:
        """Compares two terms."""
        if not isinstance(other, type(self)):
            raise TypeError(
                "'>' not supported between instances of '{}' and '{}'"
                .format(type(self), type(other)))
        for ind_a, ind_b in zip(self, other):
            if ind_a == ind_b:
                continue
            return ind_a > ind_b
        return False

    def __repr__(self) -> str:
        """A compact, human-readable representation of the term."""
        if not self.const:
            return "0"

        ret = super().__repr__()
        if self.const == 1:
            return ret
        if ret == '1':
            return str(self.const)

        return '{:d}{}'.format(self.const, ret)

    @classmethod
    def from_data(cls, data: dict[str, Union[int, list[int]]]) -> 'IntT':
        if not isinstance(data, dict):
            raise ValueError('Expected a dict, got: {}.' .format(type(data)))
        const = data.get('const', 0)
        if not isinstance(const, int):
            raise ValueError('Field "const": expected integer, got: {}'
                             .format(type(const)))
        return cls(const, super().from_data(data.get('ind', [])))

    def ind_eq(self, other: 'IntT') -> bool:
        """Compares only the indeterminates (not the constant) for equality."""
        if not isinstance(other, type(self)):
            return NotImplemented
        return super().__eq__(other)
