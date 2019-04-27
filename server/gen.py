#gen.py
import tensorflow as tf
import scipy.misc as misc
import numpy as np

from .owngen import generator

inputs_raw = tf.placeholder(
    tf.float32, shape=[1, None, None, 3], name='inputs_raw')
with tf.variable_scope('generator'):
    gen_output = generator(inputs_raw)

outputs = (gen_output + 1) / 2
converted_outputs = tf.image.convert_image_dtype(
    outputs, dtype=tf.uint8, saturate=True)

outputStr = tf.map_fn(tf.image.encode_png, converted_outputs,
                      dtype=tf.string, name='output_pngs')

try:
    session = tf.Session()
except Exception as e:
    print(e)
    exit()

saver = tf.train.Saver(tf.get_collection(
    tf.GraphKeys.GLOBAL_VARIABLES, scope='generator'))
saver.restore(session, "SRGAN_pre-trained/model-200000")


def inference(f):
    img = misc.imread(f, mode="RGB").astype(np.float32)
    img = img / np.max(img)
    with tf.device("/cpu:0"):
        res = session.run(outputStr, feed_dict={inputs_raw: [img]})
    return res[0]
