const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const mode = process.env.NODE_ENV || 'development';
module.exports = {
    entry: {
        main: path.resolve(__dirname, 'src/index.js'),
    },
    ignoreWarnings: [
        {
            module: /module2\.js\?[34]/,
        },
        {
            module: /[13]/,
            message: /homepage/,
        },
        /warning from compiler/,
        (warning) => true,
    ],
    output: {
        publicPath: "/",
        clean: true,
        path: path.resolve(__dirname, 'dist'),
        filename: '[name].bundle.js',
        assetModuleFilename: 'assets/[hash][ext][query]',
    },
    resolve: {
        alias: {
            constants: path.resolve(__dirname, "src", "constants"),
            images: path.resolve(__dirname, "src", "images"),
            pages: path.resolve(__dirname, "src", "pages"),
            utils: path.resolve(__dirname, "src", "utils"),
        },
    },
    devServer: {
        proxy: [
            {
                context: ['/api/**'],
                target: 'http://localhost:5001',
                pathRewrite: {'^/api': ''},
                changeOrigin: true,
                secure: false,
            },
        ],
        allowedHosts: 'all',
        webSocketServer: 'ws',
        historyApiFallback: true,
        port: 5000,
        open: true,
        compress: true,
        hot: true,
    },
    plugins: [
        new CopyWebpackPlugin({
            patterns: [
                {
                    from: './public',
                    globOptions: {
                        ignore: [
                            "**/index.html",
                        ],
                    },
                }
            ],
        }),
        new CleanWebpackPlugin(),
        new HtmlWebpackPlugin({
            template: "./public/index.html",
            favicon: "./public/favicon.ico",
        }),
    ],
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: ['babel-loader'],
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/i,
                type: 'asset/resource',
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/i,
                type: 'asset/resource',
            },
            {
                test: /\.css$/,
                use: [
                    'style-loader',
                    'css-loader',
                ],
            },
            {
                test: /\.scss$/,
                include: path.resolve(__dirname, "src"),
                use: [
                    'style-loader',
                    {
                        loader: "css-loader",
                        options: {
                            importLoaders: 1,
                            modules: {
                                localIdentName: "[name]__[local]___[hash:base64:5]"
                            }
                        },
                    },
                    'postcss-loader',
                    'sass-loader'
                ],
            },
        ]
    }
};
