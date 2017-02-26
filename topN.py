#!/usr/bin/env python

import argparse
import os
import time

from pprint import pprint as pp

def check(final, number):

    # Optimization to avoid scan the whole list if the number
    # wont make the cut
    if number < final[-1]:
        return

    # Original algorithm
    for i in final:
        if number > i:
            pos = final.index(i)
            final.insert(pos, number)
            final.pop()
            break


def topN(filename, top_numbers):
    final = [x for x in range(int(top_numbers)-1, -1, -1)]
    with open(filename, 'r') as f:
        for number in f:
            if number != '':
                check(final, int(number.strip('\n')))
    pp(final)

def main(filename, top_numbers):
    start = time.time()
    topN(filename, top_numbers)
    end = time.time()
    stat = os.stat(filename)
    print('To sort the %s highest numbers from  %s with size %s MB it took %s' % (top_numbers, filename, (stat.st_size/1048576), end-start))


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='topN: find the top N largest number from a file')
    parser.add_argument('--filename', help='file to search for top numbers', required=True)
    parser.add_argument('--top-numbers', help='number of top numbers to be retrieved', required=True)
    args = parser.parse_args()

    main(args.filename, args.top_numbers)
