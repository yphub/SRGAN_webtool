import tensorflow as tf
import scipy.misc as misc
import numpy as np

with tf.Session() as sess:
    saver = tf.train.import_meta_graph('SRGAN_pre-trained/model-200000.meta')
    saver.restore(sess, "SRGAN_pre-trained/model-200000")
    graph = tf.get_default_graph()
    inputs_raw = graph.get_tensor_by_name("inputs_raw:0")
    outputStr = graph.get_tensor_by_name("output_pngs:0")

    img = misc.imread("img_001.png", mode="RGB").astype(np.float32)
    img = img / 255
    res = sess.run(outputStr, feed_dict={inputs_raw: [img]})
    with open("img_001.if.png" , "wb") as f:
        f.write(res)