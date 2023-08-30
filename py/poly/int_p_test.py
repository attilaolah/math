"""Tests for module int_p."""
import unittest

from poly import int_p


class TestIntP(unittest.TestCase):
    """Tests the int_p.IntP class."""

    def test_add(self):
        """Tests the add method."""
        self.assertEqual(
            repr(int_p.IntP(
                (2, (1,)),  # 2x
            ) + int_p.IntP(
                (2, (1,)),  # 2x
            )),
            '4x'
        )
        self.assertEqual(
            repr(int_p.IntP(
                (1, (0, 1)),  # y
                (1, (1, 0)),  # x
                (2, (1, 1)),  # 2xy
            ) + int_p.IntP(
                (2, (1, 1)),  # 2xy
            )),
            '4xy + x + y'
        )


if __name__ == '__main__':
    unittest.main()
