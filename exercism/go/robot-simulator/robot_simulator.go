package robot

import "fmt"

var N = Dir(0)
var S = Dir(1)
var E = Dir(2)
var W = Dir(3)

func (d Dir) String() string {
	conversion := map[Dir]string{N: "N", S: "S", E: "E", W: "W"}
	if dir, ok := conversion[d]; !ok {
		return "?"
	} else {
		return dir
	}
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
	case S:
		Step1Robot.Y += -1
	case E:
		Step1Robot.X += 1
	case W:
		Step1Robot.X += -1
	}
}

func Left() {
	rotation := map[Dir]Dir{N: W, W: S, S: E, E: N}
	Step1Robot.Dir = rotation[Step1Robot.Dir]
}

func Right() {
	rotation := map[Dir]Dir{N: E, E: S, S: W, W: N}
	Step1Robot.Dir = rotation[Step1Robot.Dir]
}

func (r *Step2Robot) Left() {
	rotation := map[Dir]Dir{N: W, W: S, S: E, E: N}
	r.Dir = rotation[r.Dir]
}

func (r *Step2Robot) Right() {
	rotation := map[Dir]Dir{N: E, E: S, S: W, W: N}
	r.Dir = rotation[r.Dir]
}

func (r *Step2Robot) Advance() {
	switch r.Dir {
	case N:
		r.Northing += 1
	case S:
		r.Northing += -1
	case E:
		r.Easting += 1
	case W:
		r.Easting += -1
	}
}

func (r *Step2Robot) Action(a Action) {
	switch a.cmd {
	case 'R':
		r.Right()
	case 'L':
		r.Left()
	case 'A':
		r.Advance()
	}
}

type Action struct {
	cmd Command
}

func StartRobot(cmd chan Command, action chan Action) {
	// The test program then sends commands to Robot.  When it is done sending
	// commands, it closes the command channel.  Robot must accept commands and
	// inform Room of actions it is attempting.  When it senses the command channel
	// closing, it must shut itself down.
	go func() {
		for c := range cmd {
			_ = c
			action <- Action{
				cmd: c,
			}
		}

		defer close(action) // Is it required? Or it's being done by the test?
	}()
}

func checkMapLimits(extent Rect, robot Step2Robot) bool {
	if robot.Easting < extent.Min.Easting || robot.Easting > extent.Max.Easting {
		return false
	}

	if robot.Northing < extent.Min.Northing || robot.Northing > extent.Max.Northing {
		return false
	}
	return true
}

func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	// The room must interpret the physical
	// consequences of the robot actions.  When it senses the robot shutting down,
	// it sends a final report back to the test program, telling the robot's final
	// position and direction.

	go func() {
		defer close(rep)
		for a := range act {
			newRobot := robot
			newRobot.Action(a)
			if checkMapLimits(extent, newRobot) {
				robot = newRobot
			}
		}

		rep <- robot
	}()

}

// Messages must be sent on the log channel for
// *  A robot without a name  // Robot itself. The robot has to be active? Or we just ignore it? :/ TODO: Check!!
// *  Duplicate robot names // From the room
// *  Robots placed at the same place // From the room?
// *  A robot placed outside of the room // Room
// *  An undefined command in a script // ??? ok
// *  An action from an unknown robot // By name?
// *  A robot attempting to advance into a wall // By room
// *  A robot attempting to advance into another robot // By room

type Action3 struct {
	action    rune
	robotName string
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "robot without a name" // If no one is reading this could block at infinitum
	}

	for _, act := range script {
		switch act {
		case 'L', 'R', 'A':
		default:
			log <- fmt.Sprintf("Unknown command: %v", act)
			continue
		}

		action <- Action3{
			action:    act,
			robotName: name,
		}
	}

}

func duplicateNames(robots []Step3Robot, log chan string) {
	robotNames := map[string]bool{}
	for _, r := range robots {
		if _, ok := robotNames[r.Name]; ok {
			log <- fmt.Sprintf("two or more robots share the same name: %v", r.Name)
		}
		robotNames[r.Name] = true
	}
}

func robotsByName(robots []Step3Robot) map[string]Step3Robot {
	robotNames := map[string]Step3Robot{}
	for _, r := range robots {
		robotNames[r.Name] = r
	}
	return robotNames
}

func checkMapLimits3(extent Rect, robot Step3Robot) bool {
	if robot.Easting < extent.Min.Easting || robot.Easting > extent.Max.Easting {
		return false
	}

	if robot.Northing < extent.Min.Northing || robot.Northing > extent.Max.Northing {
		return false
	}
	return true
}

// Messages must be sent on the log channel for
// *  Robots placed at the same place // From the room?
// *  A robot attempting to advance into another robot // By room

func filterValidRobots(extent Rect, robots []Step3Robot, log chan string) []Step3Robot {
	n := make([]Step3Robot, 0)
	for _, r := range robots {
		if checkMapLimits3(extent, r) {
			n = append(n, r)
		} else {
			log <- fmt.Sprintf("robot with name %v is out of bounds", r.Name)
		}
	}

	return n
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	duplicateNames(robots, log)
	byName := robotsByName(robots)
	robots = filterValidRobots(extent, robots, log)

	for act := range action {
		robot, found := byName[act.robotName]
		if !found {
			log <- fmt.Sprintf("unknown robot %v", act.robotName)
			continue
		}

		copyRobot := robot
		copyRobot.Action(Action{cmd: Command(act.action)})
		if !checkMapLimits3(extent, copyRobot) {
			log <- fmt.Sprintf("robot %v trying to get outbounds", act.robotName)
			continue
		}

		robot = copyRobot
	}
	report <- robots
}
