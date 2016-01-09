package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/* Test Adding One Ball */
	{
		bc := NewBallClock(27)
		bc.cycle_next_ball()
		json_string := bc.to_json()
		expected_json := "{\"Min\":[1],\"FiveMin\":[],\"Hour\":[],\"Main\":[2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27]}"
		if json_string == expected_json {
			fmt.Println("Success at Test Adding One Ball:")
		} else {
			fmt.Printf("Failed at Test Adding One Ball\n\tResult:  %v\n\tExpected: %v\n\n", json_string, expected_json)
		}
	}
	/* End Test Adding One Ball */

	/* Test Adding Fifth Ball */
	{
		bc := NewBallClock(27)
		for i := 0; i < 5; i += 1 {
			bc.cycle_next_ball()
		}
		json_string := bc.to_json()
		expected_json := "{\"Min\":[],\"FiveMin\":[5],\"Hour\":[],\"Main\":[6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,4,3,2,1]}"
		if json_string == expected_json {
			fmt.Println("Success at Test Adding Fifth Ball:")
		} else {
			fmt.Printf("Failed at Test Adding Fifth Ball\n\tResult:  %v\n\tExpected: %v\n\n", json_string, expected_json)
		}
	}
	/* End Test Adding Fifth Ball */
}

type Ball int

type BallTrack []Ball

type BallClock struct {
	Min     BallTrack
	FiveMin BallTrack
	Hour    BallTrack
	Main    BallTrack
}

func NewBallClock(ball_count int) (ball_clock BallClock) {
	ball_clock.Min = make([]Ball, 0, 4)
	ball_clock.FiveMin = make([]Ball, 0, 11)
	ball_clock.Hour = make([]Ball, 0, 11)
	ball_clock.Main = make([]Ball, 0, ball_count)

	for i := 1; i <= ball_count; i += 1 {
		ball_clock.Main.add_ball(Ball(i), &(ball_clock.Main))
	}
	return
}

func (bc *BallClock) cycle_next_ball() {
	next_ball := (*bc).Main.get_ball_from_front()

	next_ball, add_to_next_tray := (*bc).Min.add_ball(next_ball, &((*bc).Main))

	if add_to_next_tray {
		next_ball, _ /*add_to_next_day*/ = (*bc).FiveMin.add_ball(next_ball, &((*bc).Main))
	}
}

func (bt *BallTrack) add_ball(new_ball Ball, main_reservoir *BallTrack) (next_ball Ball, had_extra_ball bool) {
	if cap(*bt) == len(*bt) {
		// cycle through all the balls in the main tray and put in Main
		for len(*bt) > 0 {
			ball_for_reservoir := bt.get_ball()
			(*main_reservoir).add_ball(ball_for_reservoir, main_reservoir)
		}
		return new_ball, true

	} else {
		*bt = append(*bt, new_ball)
	}
	return 0, false
}

func (bt *BallTrack) get_ball() (ball Ball) {
	ball, *bt = (*bt)[len(*bt)-1], (*bt)[:len(*bt)-1]
	return
}

func (bt *BallTrack) get_ball_from_front() (ball Ball) {
	ball = (*bt)[0]
	copy((*bt)[0:], (*bt)[1:])
	(*bt) = (*bt)[:len(*bt)-1]
	return
}

func (bc *BallClock) to_json() (json_string string) {
	json_bytes, _ := json.Marshal(bc)
	json_string = string(json_bytes)
	return json_string
}
