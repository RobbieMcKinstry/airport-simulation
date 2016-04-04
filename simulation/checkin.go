package simulation

import (
	"container/heap"
)

// check the state, and then handle each transition case
func (c *Counter) Visit() {
	switch c.State {
	case PrintingBoardingPass:
		c.CheckBags()
	case CheckingBags:
		c.MiscDelays()
	case MiscDelays:
		c.ReadyForSecurity()
	case EmptyQueue:
		c.EmptyQueue()
	default:
		panic("Reached an impossible state while at a counter.")
	}
}

// The customer is about to check their bags.
func (c *Counter) CheckBags() {
	c.State = CheckingBags

	// Calculate the total amount of time from now when the event will occur
	sum := 0.0
	for i := 0; int64(i) < c.current.Bags(); i++ {
		sum += BagCheckGen()
	}

	c.Time += uint64(sum)
	heap.Push(c.A.EventHeap, c) // Add the event to be fired once the bags are checked
}

// The customer is about to experience delays
func (c *Counter) MiscDelays() {

	c.State = MiscDelays

	// Calculate the amount of time from now when the delays will be over
	c.Time += uint64(MiscGen())
	heap.Push(c.A.EventHeap, c) // Add the event to be fired once the deplays are over
}

func (c *Counter) ReadyForSecurity() {

	//////////////////////////////////////////////////////////////////////////////////////////
	// First, check to see if the passenger is first class
	// if they're first class, add them to the first class security queue
	// if the passenger is not first class, then find the shortest queue and add them to it
	// then, pull someone out of line, based on if you're in a first class counter or not
	// if there's no one in the queue, then add an empty queue event at one minute from now
	/////////////////////////////////////////////////////////////////////////////////////////
	c.State = AtSecurity

	// pull someone off the queue
	// TODO reduce the complexity by handling a variable number of queues
	// TODO reduce the complexity by handling a first class queue differently from a coach queue
	if c.current.IsFirstClass() {
		c.A.SecurityFirstClass.Append(c.current)
	} else {
		if c.A.SecurityCoach.Size() > c.A.SecurityCoach2.Size() {
			c.A.SecurityCoach.Append(c.current)
		} else {
			c.A.SecurityCoach2.Append(c.current)
		}
	}

	// TODO finish this
	// if no one's in the queue, then requeue this with an incremented time
	if c.IsFirstClass {

	} else {

	}
}

// TODO implement this, diferently for a first class and for a coach queue
func (c *Counter) EmptyQueue() {
	if c.IsFirstClass {
		c.FirstClassEmpty()
	} else {
		c.CoachEmpty()
	}
}

func (c *Counter) FirstClassEmpty() {
	// TODO figure out a way to make queue time fair and balanced so people don't wait too long
}

func (c *Counter) CoachEmpty() {
	// TODO figure out a way to make queue time fair and balanced so people don't wait too long
}
