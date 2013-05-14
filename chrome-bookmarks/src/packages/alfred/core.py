# -*- coding: utf-8 -*-
import unicodedata
import sys

_UNESCAPE_CHARACTERS = u""" ;()"""


def unescape(query, characters=None):
    for character in (_UNESCAPE_CHARACTERS if not characters else characters):
        query = query.replace('\\%s' % character, character)

    return query


def decode(s):
    return unicodedata.normalize('NFC', s.decode('utf-8'))


def args(characters=None):
    return tuple(unescape(decode(arg), characters) for arg in sys.argv[1:])


def write(xml):
    sys.stdout.write(xml)
    sys.stdout.flush()