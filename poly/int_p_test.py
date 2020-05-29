"""Tests for module int_p."""
import unittest

from poly import int_p


class TestIntP(unittest.TestCase):
    """Tests the int_p.IntP class."""

    def test_add(self):
        """Tests the add method."""
        self.assertEqual(
            repr(int_p.IntP(
                (2, (1,)),
            ) + int_p.IntP(
                (2, (1,)),
            )),
            '4x'
        )
        self.assertEqual(
            repr(int_p.IntP(
                (1, (0, 1)),
                (1, (1, 0)),
                (2, (1, 1)),
            ) + int_p.IntP(
                (2, (1, 1)),
            )),
            '4xy + x + y'
        )


if __name__ == '__main__':
    unittest.main()
