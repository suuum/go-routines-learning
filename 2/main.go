// Hint 1: time.Ticker can be used to cancel function
// Hint 2: to calculate time-diff for Advanced lvl use:
//  start := time.Now()
//	// your work
//	t := time.Now()
//	elapsed := t.Sub(start) // 1s or whatever time has passed

package main

import "time"

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool  // can be used for 2nd level task. Premium users won't have 10 seconds limit.
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false if process had to be killed
func HandleRequest(process func(), u *User) bool {
	// TODO: you need to modify only this function and implement logic that will return false for 2 levels of tasks.
	ticker := time.NewTicker(10 * time.Second)
	start := time.Now()
	process()
	t := time.Now()
	elapsed := t.Sub(start) // 1s or whatever time has passed
	u.TimeUsed += int64(elapsed / time.Second)

	for {
		select {
		case <-ticker.C:
			if u.IsPremium == false && u.TimeUsed >= 10 {
				return false
			}

		default:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
