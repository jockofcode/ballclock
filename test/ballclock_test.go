package testing

import (
	"github.com/jockofcode/ballclock"
	"testing"
)

func TestAddOneBall(t *testing.T) {
	{
		bc := ballclock.NewBallClock(27)
		bc.CycleBalls(1)
		json_string := bc.ToJSON()
		expected_json := "{\"Min\":[1],\"FiveMin\":[],\"Hour\":[],\"Main\":[2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27]}"
		if json_string != expected_json {
			t.Error("\nExpected:", expected_json, "\nReceived:", json_string)
		}
	}
}

func TestAddFifthBall(t *testing.T) {
	/* Test Adding Fifth Ball */
	{
		bc := ballclock.NewBallClock(27)
		bc.CycleBalls(5)
		json_string := bc.ToJSON()
		expected_json := "{\"Min\":[],\"FiveMin\":[5],\"Hour\":[],\"Main\":[6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,4,3,2,1]}"
		if json_string != expected_json {
			t.Error("\nExpected:", expected_json, "\nReceived:", json_string)
		}
	}
	/* End Test Adding Fifth Ball */
}

func TestAddSixtiethBall(t *testing.T) {
	/* Test Adding Sixtieth Ball */
	{
		bc := ballclock.NewBallClock(27)
		bc.CycleBalls(60)
		json_string := bc.ToJSON()
		expected_json := "{\"Min\":[],\"FiveMin\":[],\"Hour\":[24],\"Main\":[16,17,18,4,3,21,22,9,8,7,26,14,13,12,11,1,27,23,19,6,2,25,20,15,10,5]}"
		if json_string != expected_json {
			t.Error("\nExpected:", expected_json, "\nReceived:", json_string)
		}
	}
	/* End Test Adding Sixtieth Ball */
}

func TestAddSevenHundredTwentiethBall(t *testing.T) {
	/* Test Adding Seven Hundred Twentieth Ball */
	{
		bc := ballclock.NewBallClock(27)
		bc.CycleBalls(720)
		json_string := bc.ToJSON()
		expected_json := "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[3,26,13,2,21,16,8,14,6,18,5,12,10,11,9,4,23,19,25,1,27,22,17,7,15,24,20]}"
		if json_string != expected_json {
			t.Error("\nExpected:", expected_json, "\nReceived:", json_string)
		}
	}
	/* End Test Adding Seven Hundred Twentieth  Ball */
}
