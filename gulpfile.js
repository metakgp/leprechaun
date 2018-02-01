var gulp          = require("gulp");
var fs            = require("fs");
var include       = require("gulp-include");
var replace       = require('gulp-replace');

gulp.task("html", function() {
	console.log("-- gulp is running task 'scripts'");

	var commit = "";

	const exec = require('child_process').exec;
	const child = exec('git rev-list HEAD | head -n1',
	(error, stdout, stderr) => {
		if (!error) {
			commit = stdout;
		}

		var head_str = "";
		if (commit.length > 0) {
			head_str = " (" + commit.substr(0, 6) + ")";
		}

		try {
			fs.mkdirSync("dist");
		} catch(e) {
			// pass
		}

		gulp.src("html/*.html")
		.pipe(include())
		.pipe(replace("[git_head]", head_str))
		.on('error', console.log)
		.pipe(gulp.dest("dist/html"));

		gulp.src("templates/*.html")
		.pipe(include())
		.on('error', console.log)
		.pipe(gulp.dest("dist/templates"));

	});
});

gulp.task("default", ["html"]);
