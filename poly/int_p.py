"""Module int_p implements a polynomial wint integer terms and coefficients."""
from typing import Iterable, List, Tuple

from poly.int_t import IntT


# TODO: Subclass typing.List[IntT] when pylint supports this.
# See https://github.com/PyCQA/pylint/issues/3129.
class IntP(list):
    """A polynomial with int terms and coefficients.

    This type implements a sparse representation, i.e. only non-zero terms are
    stored.
    """

    def __init__(self, *args: Tuple[int, Iterable[int]]) -> None:
        """Initialise with a list of constants and terms."""
        for const, ind in args:
            self.append(IntT(const, ind))

    def __add__(self, other: 'IntP') -> 'IntP':
        """Add two polynomials together."""
        if not isinstance(other, type(self)):
            return NotImplemented
        ret = self.__class__()
        ret.extend(self)
        ret.extend(other)
        return ret._compact()

    def __repr__(self) -> str:
        terms: List[str] = []

        for term in self:
            if not term.const:
                # Exclude "+ 0" terms.
                continue
            if term.const > 0:
                terms += "+"
            else:
                terms += "-"
                term.const *= -1
            terms.append(repr(term))

        return ' '.join(terms[1:])

    def _compact(self):
        """Merges terms with the same indeterminates."""
        self.sort(reverse=True)

        ret = self.__class__()
        for term in self:
            size = len(ret)
            if not size:
                ret.append(term)
                continue
            if ret[size-1].ind_eq(term):
                ret[size-1].const += term.const
                continue
            print('XXX: {} > {}'.format(ret, term))
            ret.append(term)
        return ret
