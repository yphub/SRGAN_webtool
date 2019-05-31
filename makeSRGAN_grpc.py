# this file must be called after make SRGAN_web model
import tensorflow as tf

def buildSignatureDef(inputs, outputs, methodName=tf.saved_model.signature_constants.PREDICT_METHOD_NAME):
    return tf.saved_model.build_signature_def(
        inputs={
            k: tf.saved_model.utils.build_tensor_info(inputs[k]) for k in inputs
        },
        outputs={
            k: tf.saved_model.utils.build_tensor_info(outputs[k]) for k in outputs
        },
        method_name=methodName
    )

with tf.Session() as sess:
    saver = tf.train.import_meta_graph('SRGAN_web/model-200000.meta')
    saver.restore(sess, "SRGAN_web/model-200000")
    graph = tf.get_default_graph()
    inputs_str = graph.get_tensor_by_name("inputs_str:0")
    output_str = graph.get_tensor_by_name("output_pngs:0")

    signature_map = {
        "inference": buildSignatureDef(
            inputs={
                "inputs": inputs_str
            },
            outputs={
                "outputs": output_str
            }
        )
    }

    builder = tf.saved_model.builder.SavedModelBuilder("SRGAN_grpc")
    builder.add_meta_graph_and_variables(
        sess,
        tags=[tf.saved_model.tag_constants.SERVING],
        signature_def_map=signature_map    
    )
    builder.save()

print("over")