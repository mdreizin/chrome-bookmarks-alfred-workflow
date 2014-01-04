# -*- coding: utf-8 -*-
import getopt
import sys
import packages.alfred as alfred
import packages.workflow as core


def main(argv):
    try:
        opts, args = getopt.getopt(argv, 'v:c:q:', ['vendor', 'command', 'query'])
    except getopt.error:
        sys.exit(2)

    vendor = None
    command = None
    query = None

    for opt, arg in opts:
        if opt in ('-v', '--vendor'):
            vendor = arg.strip()
        elif opt in ('-c', '--command'):
            command = arg.strip()
        elif opt in ('-q', '--query'):
            query = arg.strip()

    if vendor and command:
        workflow = alfred.Workflow()
        provider = core.providers.create(vendor, workflow.settings)

        icon = provider.icon

        if command == 'get.bookmarks':
            results = map(lambda x: workflow.feedback.Item(
                attributes={'uid': workflow.uid(), 'arg': x['url'], 'valid': u'yes'},
                icon=icon,
                title=x['title'],
                subtitle=x['url']
            ), provider.get_bookmarks(query))

            if not results:
                results = workflow.feedback.Item(
                    attributes={'uid': workflow.uid(), 'valid': u'no'},
                    icon=icon,
                    title=u'No bookmarks found',
                    subtitle=u'No bookmarks matching your query were found'
                )

            workflow.feedback.add(results)

            xml = workflow.feedback.to_xml()

            alfred.write(xml)

if __name__ == '__main__':
    main(list(alfred.args()))