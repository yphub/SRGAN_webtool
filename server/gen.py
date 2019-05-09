#gen.py
import tensorflow as tf
import numpy as np

# from .owngen import generator
from owngen import generator

inputs_str = tf.placeholder(tf.string, name='inputs_str')

inputs_intraw = tf.image.decode_image(inputs_str, channels=3)
inputs_raw = tf.convert_to_tensor(np.array([tf.image.convert_image_dtype(inputs_intraw, dtype=tf.float32, name='inputs_raw')]))

with tf.variable_scope('generator'):
    gen_output = generator(inputs_raw)

outputs = (gen_output + 1) / 2
converted_outputs = tf.image.convert_image_dtype(
    outputs, dtype=tf.uint8, saturate=True)

outputStr = tf.image.encode_png(converted_outputs[0], name="output_pngs")

try:
    session = tf.Session()
except Exception as e:
    print(e)
    exit()

saver = tf.train.Saver(tf.get_collection(
    tf.GraphKeys.GLOBAL_VARIABLES, scope='generator'))
saver.restore(session, "SRGAN_pre-trained/model-200000")


def inference(st):    
    with tf.device("/cpu:0"):
        res = session.run(outputStr, feed_dict={inputs_str: st})
    return res

def saver():
    session.run(tf.global_variables_initializer())
    saver = tf.train.Saver()
    saver.save(session, "SRGAN_pre-trained/model", global_step=200000)

if __name__ == "__main__":
    saver()
    print(outputStr.name)