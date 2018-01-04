var gulp          = require("gulp");
var fs            = require("fs");
var include       = require("gulp-include");

gulp.task("html", function() {
    console.log("-- gulp is running task 'scripts'");

    try {
        fs.mkdirSync("dist");
    } catch(e) {
        // pass
    }

    gulp.src("html/*.html")
    .pipe(include())
    .on('error', console.log)
    .pipe(gulp.dest("dist/html"));

    gulp.src("templates/*.html")
    .pipe(include())
    .on('error', console.log)
    .pipe(gulp.dest("dist/templates"));
});

gulp.task("default", ["html"]);
