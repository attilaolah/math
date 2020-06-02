"""Tests for module int_t."""
import unittest

from poly import int_t


class TestInd(unittest.TestCase):
    """Tests the int_t.Ind class."""

    def test_repr(self):
        """Tests the __repr__ method."""
        for item, res in [
                (int_t.Ind(), '1'),
                (int_t.Ind(0), '1'),
                (int_t.Ind(1), 'x'),
                (int_t.Ind(1, 1), 'xy'),
                (int_t.Ind(1, 0, 1), 'xz'),
                (int_t.Ind(0, 1, 0), 'y'),
                (int_t.Ind(0, 0, 0, 0), '1'),
                (int_t.Ind(0, 0, 0, 1), 'x₃'),
                (int_t.Ind(5), 'x⁵'),
                (int_t.Ind(-6), 'x¯⁶'),
                (int_t.Ind(123, -321, 0), 'x¹²³y¯³²¹'),
                (int_t.Ind(-1, 2, 3, -4, 5, 6), 'x₀¯¹x₁²x₂³x₃¯⁴x₄⁵x₅⁶'),
        ]:
            self.assertEqual(repr(item), res)


class TestIntT(unittest.TestCase):
    """Tests the int_t.IntT class."""

    def test_repr(self):
        """Tests the __repr__ method."""
        for item, res in [
                (int_t.IntT(), '0'),
                (int_t.IntT(0), '0'),
                (int_t.IntT(0, (1, 2, 3)), '0'),
                (int_t.IntT(10, (1, 2, 3)), '10xy²z³'),
                (int_t.IntT(-20, (-1, 0, 1)), '-20x¯¹z'),
                (int_t.IntT(5, ()), '5'),
                (int_t.IntT(-8, (0, 0)), '-8'),
                (int_t.IntT(1, (0, 0, 1)), 'z'),
                (int_t.IntT(1, (0, 1, 1)), 'yz'),
        ]:
            self.assertEqual(repr(item), res)


if __name__ == '__main__':
    unittest.main()
