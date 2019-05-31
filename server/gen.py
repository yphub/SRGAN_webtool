# gen.py
import tensorflow as tf
import numpy as np

from .owngen import generator

try:
    sess = tf.Session()
except Exception as e:
    print(e)
    exit()

saver = tf.train.import_meta_graph('SRGAN_web/model-200000.meta')
saver.restore(sess, "SRGAN_web/model-200000")
graph = tf.get_default_graph()
inputs_str = graph.get_tensor_by_name("inputs_str:0")
output_str = graph.get_tensor_by_name("output_pngs:0")

def inference(st):
    with tf.device("/cpu:0"):
        res = sess.run(output_str, feed_dict={inputs_str: st})
    return res
