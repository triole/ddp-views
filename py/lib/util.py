import pprint


def read_file(filename):
    fileObject = open(filename, 'r')
    data = fileObject.read()
    return data


def pp(data):
    pp = pprint.PrettyPrinter(indent=4)
    pprint.PrettyPrinter(width=41, compact=True)
    pp.pprint(data)
