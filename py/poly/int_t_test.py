"""Tests for module int_t."""
import abc
import json
import os
import unittest

from poly import int_t


class _TestData(abc.ABC):
    """Base class for loading test data."""

    test_data: list
    test_class = abc.abstractproperty

    def load_test_data(self):
        test_data_path = os.path.join(
            os.getcwd(), 'testdata', 'poly', 'int_t.json')
        with open(test_data_path, 'r') as f:
            return json.load(f)[self.test_class.__name__]


class TestInd(unittest.TestCase, _TestData):
    """Tests the int_t.Ind class."""

    test_class = int_t.Ind

    def setUp(self):
        self.test_data = self.load_test_data()

    def test_json(self):
        """Tests the from_json and to_json methods only."""
        for item in self.test_data:
            json_data = json.dumps(item['data'])
            ind = int_t.Ind.from_json(json_data)
            self.assertEqual(ind.to_json(), json_data)

    def test_json_decode_error(self):
        """Tests the from_json and to_json methods."""
        with self.assertRaises(json.JSONDecodeError):
            int_t.Ind.from_json('')
        with self.assertRaises(ValueError):
            int_t.Ind.from_json('{}')
        with self.assertRaises(ValueError):
            int_t.Ind.from_json('[{}]')

    def test_repr(self):
        """Tests the __repr__ method."""
        for item in self.test_data:
            ind = int_t.Ind.from_json(json.dumps(item['data']))
            self.assertEqual(repr(ind), item['repr'])


class TestIntT(unittest.TestCase, _TestData):
    """Tests the int_t.IntT class."""

    test_class = int_t.IntT

    def setUp(self):
        self.test_data = self.load_test_data()

    def test_repr(self):
        """Tests the __repr__ method."""
        for item in self.test_data:
            ind = int_t.IntT.from_json(json.dumps(item['data']))
            self.assertEqual(repr(ind), item['repr'])


if __name__ == '__main__':
    unittest.main()
