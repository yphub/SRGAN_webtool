# main.py
import tornado.ioloop
import tornado.web
import tensorflow
from tornado.web import StaticFileHandler
from io import BytesIO

from server.handler import handlers
# import server.gen as generator

print("init tornado apps...")
app = tornado.web.Application(
    handlers + [(r"/(.*)", StaticFileHandler,
                 {"path": "webfront/dist", "default_filename": "index.html"})],
    debug=True
)
app.listen(8082)

print("Starting tornado event loop listening at port 8082")
try:
    tornado.ioloop.IOLoop.instance().start()
except KeyboardInterrupt as e:
    print("closing session...")
    tensorflow.Session().close()
