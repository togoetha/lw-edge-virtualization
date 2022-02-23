from osv.modules import api

api.require('golang')
default = api.run(cmdline="/go.so /gorestapi.so")
