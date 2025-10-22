package greetings_test

import (
    "testing"
    
    "github.com/stretchr/testify/assert"
    
    // We now need to import the "production" package 
    "github.com/example/module/src/greetings"
)

func TestGreeting(t *testing.T) {
    assert.NotEqual(t, greetings.Greeting(), "")
}