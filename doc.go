// package tenten implements "a coding system for entering any location on earth with 10m of accuracy using a 10 character code that includes features to prevent errors in entering the code." - purposed by john graham-cumming
//
// reference:
//	http://blog.jgc.org/2010/06/1010-code.html?
//	http://blog.jgc.org/2006/07/simple-code-for-entering-latitude-and.html
//
//
// benchmark on my i7
//	BenchmarkEncode		10000000	       174 ns/op
//	BenchmarkEncodeSlow	  500000	      4880 ns/op
package tenten
