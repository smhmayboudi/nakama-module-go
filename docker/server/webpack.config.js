const webpack = require('webpack');
const webpackMerge = require('webpack-merge');
const path = require('path');

const directory = path.resolve(__dirname);

/**
 * @type {import("webpack").WebpackOptionsNormalized}
 */
const common = {
  entry: {
    app: 'src/index.ts',
  },
  mode: 'development',
  module: {
    rules: [
      {
        exclude: /(node_modules)/,
        test: /\.jsx?$/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-env']
          },
        },
      },
      {
        exclude: /(node_modules)/,
        test: /\.tsx?$/,
        use: {
          loader: 'ts-loader',
        },
      },
    ],
  },
  optimization: {
    minimize: false,
  },
  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'dist'),
    sourceMapFilename: '[file].map',
  },
  resolve: {
    extensions: [
      '.json',
      '.js',
      '.jsx',
      '.ts',
      '.tsx',
    ],
    modules: [
      'node_modules',
      path.resolve(directory),
    ],
  },
  target: 'web',
};

/**
 * @type {import("webpack").WebpackOptionsNormalized}
 */
module.exports = webpackMerge.merge(common, {
  devtool: 'inline-source-map',
  plugins: [
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify('development'),
    }),
  ],
  resolve: {
    fallback: {
      "assert": false,
      "crypto": false,
      "net": false,
      "readline": false,
      "tls": false,
      "util": false,
      "zlib": false,
    },
  },
});