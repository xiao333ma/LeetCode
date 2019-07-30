package main

import (
	"strings"
)

func main() {

	emails :=  []string{"t12+12+11@leetcode.com","test.e.mail+bob.cathy@leetcode.com",
	"testemail+david@lee.tcode.com"}
	 numUniqueEmails(emails)

}

/*
Every email consists of a local name and a domain name, separated by the @ sign.

For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.

Besides lowercase letters, these emails may contain '.'s or '+'s.

If you add periods ('.') between some characters in the local name part of an email address,
mail sent there will be forwarded to the same address without dots in the local name.
For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.
(Note that this rule does not apply for domain names.)

If you add a plus ('+') in the local name, everything after the first plus sign will be ignored.
This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.
(Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time.

Given a list of emails, we send one email to each address in the list.  How many different addresses actually receive mails?



Example 1:

Input: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
Output: 2
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails


Note:

1 <= emails[i].length <= 100
1 <= emails.length <= 100
Each emails[i] contains exactly one '@' character.

解题思路：

正常切割拼接字符串
存储在 map 中保证唯一性
统计 map 个数
时间复杂度 O(N)
*/

func numUniqueEmails(emails []string) int {

	result := make(map[string]int, 10)
	for _, email := range emails {
		emailStrings := strings.Split(email, "@")
		prefix := strings.Split(emailStrings[0], "+")[0]
		prefixString := strings.Replace(prefix, ".", "", -1)
		trueEmail := prefixString + emailStrings[1]
		result[trueEmail] = 1
	}
	return len(result)
}