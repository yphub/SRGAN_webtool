# handler.py
from tornado.web import RequestHandler
from tornado.concurrent import run_on_executor
from concurrent.futures import ThreadPoolExecutor
from .gen import inference
from io import BytesIO


class TestHandler(RequestHandler):
    def get(self):
        self.write("Test Response")


class InferenceHandler(RequestHandler):
    executor = ThreadPoolExecutor(1)

    @run_on_executor
    def post(self):
        imgBuffer = self.request.files.get('img')[0]
        res = inference(imgBuffer.body)
        self.write(res)


handlers = [
    (r"/test", TestHandler),
    (r"/inference", InferenceHandler)
]
