const webpack = require('webpack');
const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const extractCSS = new ExtractTextPlugin({
	filename: '[name].fonts.css',
	allChunks: true
});
const extractSCSS = new ExtractTextPlugin({
	filename: '[name].styles.css',
	allChunks: true
});

const BUILD_DIR = path.resolve(__dirname, 'build');
const SRC_DIR = path.resolve(__dirname, 'src');

module.exports = {
	node: {
		fs: 'empty'
	},
	entry: {
		index: [
			'url-search-params-polyfill',
			SRC_DIR + '/index.tsx'
		]
	},
	output: {
		path: BUILD_DIR,
		filename: '[name].[hash].bundle.js'
	},
	devServer: {
		contentBase: BUILD_DIR,
		historyApiFallback: {
			disableDotRule: false
		},
		port: 8080,
		compress: true,
		hot: true,
		open: false
	},
	resolve: {
		extensions: ['.ts', '.tsx', '.js']
	},
	module: {
		rules: [
			{
				test: /\.(tsx|ts|js)$/,
				loader: 'awesome-typescript-loader',
				include: [path.join(__dirname, 'src')]
			},
			{
				test: /\.html$/,
				use: {
					loader: 'html-loader',
					options: {
						minifyJS: false
					}
				}
			},
			{
				test: /\.(scss)$/,
				use: ['css-hot-loader'].concat(extractSCSS.extract({
					fallback: 'style-loader',
					use: [
						{
							loader: 'css-loader',
							options: {
								alias: {
									'../images': '../public/images',
									'../fonts': '../public/fonts',
								},
								includePaths: ['node_modules']
							}
						},
						{
							loader: 'sass-loader',
							options: {
								includePaths: ['node_modules']
							}
						}
					]
				}))
			},

			{
				// Css files in src/app/ are modular (module: true in the css-loader config)
				// This means that the class names are prefixed and suffixed with a hash to make them scoped
				test: /\.(css)$/,
				exclude: [
					path.join(__dirname, 'node_modules'),
				],
				use: [
					{
						loader: 'style-loader',
					},
					{
						loader: 'css-loader',
						options: {
							modules: true,
							localIdentName: '[name]__[local]--[hash:base64:5]'
						}
					},
					{
						loader: 'sass-loader'
					}
				]
			},

			{
				test: /\.css$/,
				exclude: [
					path.join(__dirname, 'src'),
				],
				use: extractCSS.extract({
					fallback: 'style-loader',
					use: 'css-loader'
				})
			},
			{
				test: /\.(png|jpg|jpeg|gif|ico|svg)$/,
				use: [
					{
						loader: 'url-loader',
						// loader: 'file-loader',
						options: {
							name: './images/[name].[hash].[ext]'
						}
					}
				]
			},
			{
				test: /\.(woff(2)?|ttf|eot|svg)(\?v=\d+\.\d+\.\d+)?$/,
				loader: 'file-loader',
				options: {
					name: './fonts/[name].[hash].[ext]'
				}
			}]
	},
	optimization: {
		splitChunks: {
			cacheGroups: {
				commons: { test: /[\\/]node_modules[\\/]/, name: "vendors", chunks: "all" }
			}
		}
	},
	plugins: [
		new webpack.NamedModulesPlugin(),
		extractCSS,
		extractSCSS,
		new HtmlWebpackPlugin({
			inject: true,
			template: './public/index.html'
		}),
		new CopyWebpackPlugin([
			{ from: './public/images', to: 'images' },
			{ from: './public/locales', to: 'locales' }
		], { copyUnmodified: false })
	]
};
