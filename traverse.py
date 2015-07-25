"""Traverse N-dimensional space in a bound manner:

* First yield the zero point,
* then yield all points for coordinates <= 1,
* then yield all points for coordinates <= 2,
* …and so on.
"""
import sys


def traverse_2d(cur=0, limit=None):
    """Generate (x, y) pairs.

    This function iterates through points on a plane.

    If limit is provided, all coordinates will be <= limit.

    Yields points in this order:

        (0, 0) | cur = 0

        (0, 1) | cur = 1
        (1, 0)
        (1, 1)

        (0, 2) | cur = 2
        (1, 2)
        (2, 0)
        (2, 1)
        (2, 2)

        (0, 3) | cur = 3
        …
    """
    while limit is None or cur <= limit:
        for x__ in range(cur):
            yield x__, cur
        for y__ in range(cur):
            yield cur, y__
        yield cur, cur
        cur += 1


def traverse_3d(cur=0, limit=None):
    """Generate (x, y, z) triplets.

    This function iterates through points in 3D space.

    If limit is provided, all cordinates will be <= limit.

    Yields points in this order:

        (0, 0, 0) | cur = 0

        (0, 0, 1) | cur = 1
        (0, 1, 0)
        (0, 1, 1)
        (1, 0, 0)
        (1, 0, 1)
        (1, 1, 0)
        (1, 1, 1)

        (0, 0, 2) | cur = 2
        (0, 1, 2)
        (0, 2, 0)
        (0, 2, 1)
        (0, 2, 2)
        (1, 0, 2)
        (1, 1, 2)
        (1, 2, 0)
        (1, 2, 1)
        (1, 2, 2)
        (2, 0, 0)
        (2, 0, 1)
        (2, 1, 0)
        (2, 1, 1)
        (2, 0, 2)
        (2, 1, 2)
        (2, 2, 0)
        (2, 2, 1)
        (2, 2, 2)

        (0, 0, 3) | cur = 3
        …
    """
    while limit is None or cur <= limit:
        for x__ in range(cur):
            for y__, z__ in traverse_2d(cur, cur):
                yield x__, y__, z__
        for y__, z__ in traverse_2d(limit=cur):
            yield cur, y__, z__
        cur += 1


def main():
    """Print (n, z, r) triplets."""
    if len(sys.argv) > 1:
        limit = int(sys.argv[1])
    else:
        limit = None
    for coords in traverse_3d(limit=limit):
        print(*coords)


if __name__ == '__main__':
    main()
