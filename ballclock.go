package ballclock

import (
	"encoding/json"
)

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

	for i := 1; i <= cap(ball_clock.Main); i += 1 {
		ball_clock.Main.addBall(Ball(i), &(ball_clock.Main))
	}
	return
}

func CountDaysTillReset(num_of_balls int) (days int) {
	bc := NewBallClock(num_of_balls)
	initial_state := bc.ToJSON()
	days = 0
	for {
		bc.CycleBalls(60 * 24)

		days += 1
		if initial_state == bc.ToJSON() {
			break
		}
	}

	return days
}

func GetStateAfterCycles(num_of_balls, iterations int) (json_output string) {
	bc := NewBallClock(num_of_balls)
	bc.CycleBalls(iterations)
	return bc.ToJSON()
}

func (bc *BallClock) CycleBalls(number_of_balls int) {
	for i := 0; i < number_of_balls; i += 1 {
		next_ball := (*bc).Main.getBallFromFront()

		next_ball, add_to_five_min_tray := (*bc).Min.addBall(next_ball, &((*bc).Main))

		if add_to_five_min_tray {
			next_ball, add_to_hour_tray := (*bc).FiveMin.addBall(next_ball, &((*bc).Main))
			if add_to_hour_tray {
				next_ball, add_to_main_tray := (*bc).Hour.addBall(next_ball, &((*bc).Main))
				if add_to_main_tray {
					(*bc).Main.addBall(next_ball, &((*bc).Main))
				}
			}
		}
	}
}

func (bt *BallTrack) addBall(new_ball Ball, main_reservoir *BallTrack) (next_ball Ball, had_extra_ball bool) {
	if cap(*bt) == len(*bt) {
		// cycle through all the balls in the main tray and put in Main
		for len(*bt) > 0 {
			ball_for_reservoir := bt.getBall()
			(*main_reservoir).addBall(ball_for_reservoir, main_reservoir)
		}
		return new_ball, true

	} else {
		*bt = append(*bt, new_ball)
	}
	return 0, false
}

func (bt *BallTrack) getBall() (ball Ball) {
	ball, *bt = (*bt)[len(*bt)-1], (*bt)[:len(*bt)-1]
	return
}

func (bt *BallTrack) getBallFromFront() (ball Ball) {
	ball = (*bt)[0]
	copy((*bt)[0:], (*bt)[1:])
	(*bt) = (*bt)[:len(*bt)-1]
	return
}

func (bc *BallClock) ToJSON() (json_string string) {
	json_bytes, _ := json.Marshal(bc)
	json_string = string(json_bytes)
	return json_string
}
