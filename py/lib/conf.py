import re
from os import listdir
from os.path import join as pj
from os.path import realpath as rp

from toml import load


class Conf():
    def __init__(self):
        self.scriptname = rp(__file__)
        self.libdir = '/'.join(self.scriptname.split('/')[:-1])
        self.pydir = '/'.join(self.scriptname.split('/')[:-2])
        self.basedir = '/'.join(self.scriptname.split('/')[:-3])
        self.confdir = pj(self.basedir, 'conf')
        self.conf = load(pj(self.confdir, 'conf.toml'))
        self.meta_template = load(pj(self.confdir, 'meta_template.toml'))
        self.content_files = self.list_content_files()

    def list_content_files(self):
        path = pj(self.basedir, 'content')
        arr = []
        for fil in listdir(path):
            arr.append(pj(path, fil))
        arr.sort()
        return arr

    def filename_to_uri_split(self, filename):
        return re.search(r'[^/]+.$', filename).group(0).split('.')[0]

    def filename_to_uri_prefix(self, filename):
        return self.conf['uri_prefix'] + self.filename_to_uri_split(filename)

    def filename_to_uri(self, filename):
        return self.filename_to_uri_prefix(filename) + '/views/ddp'

    def filename_to_title(self, filename):
        t = self.filename_to_uri_split(filename)
        t = t.replace('_', ' ').replace('ddp', 'DDP')
        return t
