# owngen.py
import tensorflow as tf
import tensorflow.contrib.slim as slim


def conv2(batch_input, kernel=3, output_channel=64, stride=1, use_bias=True, scope='conv'):
    # kernel: An integer specifying the width and height of the 2D convolution window
    with tf.variable_scope(scope):
        if use_bias:
            return slim.conv2d(batch_input, output_channel, [kernel, kernel], stride, 'SAME', data_format='NHWC',
                               activation_fn=None, weights_initializer=tf.contrib.layers.xavier_initializer())
        else:
            return slim.conv2d(batch_input, output_channel, [kernel, kernel], stride, 'SAME', data_format='NHWC',
                               activation_fn=None, weights_initializer=tf.contrib.layers.xavier_initializer(),
                               biases_initializer=None)


def batchnorm(inputs):
    return slim.batch_norm(inputs, decay=0.9, epsilon=0.001, updates_collections=tf.GraphKeys.UPDATE_OPS,
                           scale=False, fused=True, is_training=False)


def prelu_tf(inputs):
    with tf.variable_scope('Prelu'):
        alphas = tf.get_variable('alpha', inputs.get_shape(
        )[-1], initializer=tf.zeros_initializer(), dtype=tf.float32)
    pos = tf.nn.relu(inputs)
    neg = alphas * (inputs - abs(inputs)) * 0.5

    return pos + neg


def pixelShuffler(inputs, scale=2):
    size = tf.shape(inputs)
    batch_size = size[0]
    h = size[1]
    w = size[2]
    c = inputs.get_shape().as_list()[-1]

    # Get the target channel size
    channel_target = c // (scale * scale)
    channel_factor = c // channel_target

    shape_1 = [batch_size, h, w, channel_factor //
               scale, channel_factor // scale]
    shape_2 = [batch_size, h * scale, w * scale, 1]

    # Reshape and transpose for periodic shuffling for each channel
    input_split = tf.split(inputs, channel_target, axis=3)
    output = tf.concat([phaseShift(x, scale, shape_1, shape_2)
                        for x in input_split], axis=3)

    return output

def phaseShift(inputs, scale, shape_1, shape_2):
    # Tackle the condition when the batch is None
    X = tf.reshape(inputs, shape_1)
    X = tf.transpose(X, [0, 1, 3, 2, 4])

    return tf.reshape(X, shape_2)

def generator(gen_inputs):

    # The Bx residual blocks
    def residual_block(inputs, output_channel, stride, scope):
        with tf.variable_scope(scope):
            net = conv2(inputs, 3, output_channel, stride,
                        use_bias=False, scope='conv_1')
            net = batchnorm(net)
            net = prelu_tf(net)
            net = conv2(net, 3, output_channel, stride,
                        use_bias=False, scope='conv_2')
            net = batchnorm(net)
            net = net + inputs

        return net

    with tf.variable_scope('generator_unit', reuse=False):
        # The input layer
        with tf.variable_scope('input_stage'):
            net = conv2(gen_inputs, 9, 64, 1, scope='conv')
            net = prelu_tf(net)

        stage1_output = net

        # The residual block parts
        for i in range(1, 17, 1):
            name_scope = 'resblock_%d' % (i)
            net = residual_block(net, 64, 1, name_scope)

        with tf.variable_scope('resblock_output'):
            net = conv2(net, 3, 64, 1, use_bias=False, scope='conv')
            net = batchnorm(net)

        net = net + stage1_output

        with tf.variable_scope('subpixelconv_stage1'):
            net = conv2(net, 3, 256, 1, scope='conv')
            net = pixelShuffler(net)
            net = prelu_tf(net)

        with tf.variable_scope('subpixelconv_stage2'):
            net = conv2(net, 3, 256, 1, scope='conv')
            net = pixelShuffler(net)
            net = prelu_tf(net)

        with tf.variable_scope('output_stage'):
            net = conv2(net, 9, 3, 1, scope='conv')

    return net
