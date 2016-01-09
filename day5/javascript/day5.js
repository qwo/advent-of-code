// Read the file and print its contents.
var fs = require('fs');
var filename = "../day5.txt";

fs.readFile(filename, 'utf8', function(err, data) {
  if (err) {
		throw err;
	 }
	 var count = 0;


	 data.split("\n").forEach(function(word){
		 if (matchRule2(word)) {
			 count++;
		 }
	 });
	 console.log(count);
});

function matchRule2(word) {
 var rule1 = word.match(/(..).*\1/);
 var rule2 = word.match(/(.).\1/);
 console.log(rule1 && rule2)
 if (rule1 && rule2) {
	 return true;
 }
}