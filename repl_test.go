package main

import "testing"

func TestCleanInput(t *testing.T){
    cases := []struct {
    input string
    expected []string
    }{
    {
        input:"hello world, this is a test",
        expected: []string{
            "hello",
            "world,",
            "this",
            "is",
            "a",
            "test",
        }, 
    },
    {
        input:"HELLO world, THIS is A TEst",
        expected: []string{
            "hello",
            "world,",
            "this",
            "is",
            "a",
            "test",
        }, 
    }, 
}
    for _, cs := range cases {
        actual := cleanInput(cs.input)
        if len(actual) != len(cs.expected){
            t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(cs.expected))
            continue
        }

        for i := range actual{
            actualWord := actual[i]
            expectedWord := cs.expected[i]
        if actualWord != expectedWord{
            t.Errorf("%v does not equal to %v",actualWord, expectedWord)
            continue
        }
    }

    }

} 
