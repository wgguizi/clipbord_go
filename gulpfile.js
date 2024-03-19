"use strict";

const {src, dest, series, parallel} = require('gulp')
const zip = require('gulp-zip');               //简单字符替换
const del = require('del');                            //node删除命令
const {exec} = require('child_process');               //node 子任务，一般用于处理shell

var binFile = ""

//清理目录
function clean(cb) {
	return del('dist/**', {force:true});
}

function other(cb) {
	// place code for your default task here
	return src(['server.sh'], {base: './'})
		.pipe(dest('dist/'))
}

function other2(cb) {
	// place code for your default task here
	return src([binFile], {base: './'})
		.pipe(dest('dist/'))
}

function end(cb) {
	// place code for your default task here
	return src('dist/**', {base: './dist'})
		.pipe(zip('dist.zip'))
		.pipe(dest('dist/'))
}

//编译go windows版本
function compileWins(cb) {
	exec("set CGO_ENABLED=1&set GOOS=windows&set GOARCH=amd64&go build main.go", function (err, stdout, stderr) {
		console.log(stdout)
		console.log(stderr)
		cb()
	})
	binFile = "main.exe"
}

//编译go linux版本
function compileLinux(cb) {
	exec("set CGO_ENABLED=0&set GOOS=linux&set GOARCH=amd64&go build main.go", function (err, stdout, stderr) {
		console.log(stdout)
		console.log(stderr)
		cb()
	})
	binFile = "main"
}

//编译go mac版本
function compileMac(cb) {
	exec("set CGO_ENABLED=0&set GOOS=darwin&set GOARCH=amd64&go build main.go", function (err, stdout, stderr) {
		console.log(stdout)
		console.log(stderr)
		cb()
	})
	binFile = "main"
}

exports.compileWins = compileWins
exports.compileLinux = compileLinux
exports.compileMac = compileMac

exports.buildWins = series(clean, parallel(compileWins), parallel(other, other2), end)
exports.buildLinux = series(clean, parallel(compileLinux), parallel(other, other2), end)
exports.buildMac = series(clean, parallel(compileMac), parallel(other, other2), end)