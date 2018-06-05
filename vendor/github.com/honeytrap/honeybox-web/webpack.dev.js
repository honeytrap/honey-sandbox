const webpack = require('webpack');

const { gitDescribeSync } = require('git-describe');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');
const dotenv = require('dotenv');

const gitInfo = gitDescribeSync();
dotenv.config();

module.exports = merge(common, {
	devtool: 'eval-cheap-module-source-map"',
	plugins: [
		new webpack.DefinePlugin({
			'process.env': {
				NODE_ENV: JSON.stringify(process.env.NODE_ENV || 'development'),
				CLIENT_VERSION: JSON.stringify(gitInfo.raw),
				SERVER_URI: JSON.stringify(process.env.SERVER_URI),
				MOCK_SERVER: process.env.MOCK_SERVER
			}
		}),
		new webpack.HotModuleReplacementPlugin(),
	]
});