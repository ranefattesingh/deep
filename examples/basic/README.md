# Description

This is a sample unit test to demonstrate the basic usage of deep package.

Here we are comparing two `Student` struct objects foo and bar with each other.

`TestForSuccess` test function demonstrates the scenario when both student objects have same content and our test passes.
`TestForFailure` test function demonstrates the scenario when both student object have different content and our test fails with the following result `[Name: foo != bar Subjects.slice[1].Marks: 60 != 61]`