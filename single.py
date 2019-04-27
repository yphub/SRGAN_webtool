# single.py
import tornado.ioloop
import tornado.web

print("init tornado apps -- single.py")


class InferenceHandler(tornado.web.RequestHandler):
    def post(self):
        imgBuffer = self.request.files.get('img')[0]
        self.write(imgBuffer.body)


app = tornado.web.Application(
    handlers=[
        (r"/inference", InferenceHandler)
    ],
    debug=True
)
app.listen(8081)

print("Starting tornado event loop listening at port 8081")
tornado.ioloop.IOLoop.instance().start()