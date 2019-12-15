package main

import "testing"

func TestGreeting(t *testing.T){
	got:= greeting("Code.education Rocks!")
	want:= "<b>Code.education Rocks!</b>"

	if got != want{
		t.Errorf("greeting('Code.education Rocks!') got: %v want: %v", got, want)
	}
}