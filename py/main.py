from lib.conf import Conf
from lib.req import Req
from lib.util import read_file

if __name__ == '__main__':
    conf = Conf()
    req = Req(conf)

    c = 100

    for i, fil in enumerate(conf.content_files):
        c = c+1
        meta = conf.meta_template
        meta['id'] = c
        meta['uri_prefix'] = conf.filename_to_uri_prefix(fil)
        meta['uri'] = conf.filename_to_uri(fil)
        meta['title_de'] = conf.filename_to_title(fil)
        meta['title_en'] = conf.filename_to_title(fil)
        meta['title_fr'] = conf.filename_to_title(fil)
        meta['title_it'] = conf.filename_to_title(fil)
        meta['template'] = read_file(fil)

        req.put(meta)
