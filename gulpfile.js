'use strict';

var path = require('path'),
    del = require('del'),
    gulp = require('gulp'),
    foreach = require('gulp-foreach'),
    zip = require('gulp-zip'),
    runSequence = require('run-sequence'),
    Gitdown = require('gitdown');

var src = './src',
    dist = './dist';

gulp.task('clean', function(callback) {
    del(path.join(dist, '**'), callback);
});

gulp.task('dist', function() {
    return gulp.src(path.join(src, '*'))
        .pipe(foreach(function(stream, file) {
            var basename = path.basename(file.path);

            return gulp.src(path.join(file.path, '**'))
                    .pipe(zip(basename + '.alfredworkflow'))
                    .pipe(gulp.dest(dist));
        }));
});

gulp.task('gitdown:readme', function() {
    return Gitdown.read('.gitdown/README.md').write('README.md');
});

gulp.task('gitdown', function(callback) {
    runSequence('gitdown:readme', callback);
});

gulp.task('docs', ['gitdown']);

gulp.task('build', [], function(callback) {
    runSequence('clean', 'docs', 'dist', callback);
});

gulp.task('default', function(callback) {
    runSequence('build', callback);
});
