package demo

// "As we known every char(acter) has unique underlying a number, so we can tackle this problem like: there are a list of numbers each showing twice except one. Using bit operation is easy to find the unique number

// const cppImplementaion=`
// class Solution {
// public:
//     char findTheDifference(string s, string t) {
//         	size_t bits=0;
// 	for (size_t i=0;i<s.size();i++){
// 		bits^=s[i];
// 	}
// 	for (size_t i=0;i<t.size();i++){
// 		bits^=t[i];
// 	}
// 	return char(bits);
//
//     }
// };
// `

func findTheDifference(s, t string) byte {
	var bits int
	for i := 0; i < len(s); i++ {
		bits ^= int(s[i])
	}
	for i := 0; i < len(t); i++ {
		bits ^= int(t[i])
	}
	return byte(bits)
}
