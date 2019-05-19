import tensorflow as tf
import numpy as np
from server.owngen import generator

def makeModel():    
    with tf.device("/cpu:0"):
        with tf.Session() as sess:

            inputs_str = tf.placeholder(tf.string, name="inputs_str")

            # string => int[][]
            inputs_intarr = tf.image.decode_image(inputs_str, channels=3)

            # float[][] => float[][]
            inputs_arr = tf.image.convert_image_dtype(inputs_intarr, dtype=tf.float32)

            inputs_shape = tf.shape(inputs_arr)
            inputs_raw = tf.reshape(inputs_arr, [1,inputs_shape[0], inputs_shape[1], 3], name='inputs_raw')

            with tf.variable_scope('generator'):
                op = generator(inputs_raw)
            
            outputs = (op + 1) / 2
            converted_outputs = tf.image.convert_image_dtype(
                outputs, dtype=tf.uint8, saturate=True)

            outputStr = tf.image.encode_png(converted_outputs[0], name="output_pngs")

            saver = tf.train.Saver(tf.get_collection(
            tf.GraphKeys.GLOBAL_VARIABLES, scope='generator'))
            saver.restore(sess, "SRGAN_pre-trained/model-200000")

            saver = tf.train.Saver()
            saver.save(sess, "SRGAN_web/model", global_step=200000)

def testInference():
    with tf.device("/cpu:0"):
        with tf.Session() as sess:
            saver = tf.train.import_meta_graph('SRGAN_web/model-200000.meta')
            saver.restore(sess, "SRGAN_web/model-200000")
            graph = tf.get_default_graph()
            inputs_str = graph.get_tensor_by_name("inputs_str:0")
            output_str = graph.get_tensor_by_name("output_pngs:0")        

            with open("test/img_001.png", "rb") as f:
                cont = f.read()        
            
            res = sess.run(output_str, feed_dict={inputs_str: cont})
            with open("test/img_001.if.png", "wb") as f:
                f.write(res)    

makeModel()
# testInference()