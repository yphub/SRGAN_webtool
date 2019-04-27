module.exports = {
    devServer: {
        proxy: {
            '/inference': {
                target: 'http://localhost:8081',
                changeOrigin: true
            }
        }
    }
}