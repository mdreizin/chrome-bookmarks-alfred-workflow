'use strict';

var path = require('path'),
    del = require('del'),
    gulp = require('gulp'),
    foreach = require('gulp-foreach'),
    zip = require('gulp-zip'),
    runSequence = require('run-sequence');

var src = './';

gulp.task('clean', function(callback) {
    del([path.join(src, '**/*.alfredworkflow')], callback);
});

gulp.task('compress', function() {
    return gulp.src([path.join(src, '**/src'), '!node_modules/**/*.*'])
        .pipe(foreach(function(stream, file) {
            var dirname = path.dirname(file.path),
                basename = path.basename(dirname);

            return gulp.src(path.join(file.path, '**/*.*'))
                    .pipe(zip(basename + '.alfredworkflow'))
                    .pipe(gulp.dest(path.join(dirname, 'zip')));
        }));
});

gulp.task('build', [], function(callback) {
    runSequence('clean', 'compress', callback);
});

gulp.task('default', function(callback) {
    runSequence('build', callback);
});
