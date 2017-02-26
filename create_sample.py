#!/usr/bin/env python


import random
import time
import sys

start = time.time()
with open('numbers.txt', 'w') as f:
    while True:
        try:
            f.write('%d\n' % random.randrange(0, 1000000000000000))
        except KeyboardInterrupt:
            end = time.time()
            print('Stopped sample creation after %d seconds' % (end-start))
            f.close()
            sys.exit(0)
