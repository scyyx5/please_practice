package greetings_test

import (
    "testing"
    
    // We now need to import the "production" package 
    "github.com/example/module/src/greetings"
)

func TestGreeting(t *testing.T) {
    if greetings.Greeting() == "" {
        panic("Greeting failed to produce a result")
    }
}