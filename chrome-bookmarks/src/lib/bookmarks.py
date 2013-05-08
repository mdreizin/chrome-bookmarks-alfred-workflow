# -*- coding: utf-8 -*-
import os
import re
import json

__BOOKMARKS_PATH = os.path.expanduser('~/Library/Application Support/Google/Chrome/Default/Bookmarks')
__BOOKMARKS_CHROMIUM_PATH = os.path.expanduser('~/Library/Application Support/Chromium/Default/Bookmarks')


def __load(is_chromium):
    path = __BOOKMARKS_CHROMIUM_PATH if is_chromium else __BOOKMARKS_PATH

    try:
        with open(path, 'r') as io:
            data = json.load(io)
    except:
        data = None

    return data


def __inspect(data, chain, predicate):
    if data:
        if type(data) == dict:
            if 'type' in data:
                if data['type'] == 'folder':
                    __inspect(data['children'], chain, predicate)
                elif data['type'] == 'url':
                    if predicate(data):
                        chain.append({'title': data['name'], 'url': data['url']})
            else:
                for value in data.itervalues():
                    __inspect(value, chain, predicate)
        elif type(data) == list:
            for value in data:
                __inspect(value, chain, predicate)
    return


def __find(predicate, is_chromium):
    items = []

    data = __load(is_chromium)

    if data:
        __inspect(data['roots'], items, predicate)

    return items


def find(query, is_chromium):
    r = re.compile(re.escape(query), re.UNICODE | re.IGNORECASE)

    return __find((lambda x: bool(r.search(x['name'])) or bool(r.search(x['url']))), is_chromium)