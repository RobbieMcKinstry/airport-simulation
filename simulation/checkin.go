package simulation

import (
	"container/heap"
)

func (bpp *BoardingPassPrinted) Visit() {
	sum := 0.0
	for i := 0; int64(i) < c.curr.Bags(); i++ {
		sum += BagCheckGen()
	}

	// Make a new BagsChecked struct
	bagsChecked := &BagsChecked{
		A:    bpp.A,
		Time: bpp.GetTime() + uint64(round(sum)),
		Curr: bpp.Curr,
	}

	// Push this new event onto the heap
	bpp.A.EventHeap.Push(bagsChecked)
}

func (bc *BagsChecked) Visit() {
	var next Visitor

	if bc.Curr.IsFirstClass() {
		next = MiscDelaysFinishedFC{
			A:    bc.A,
			Time: bc.GetTime() + uint64(round(MiscGen())),
			Curr: bc.Curr,
		}
	} else {
		next = MiscDelaysFinishedCoach{
			A:    bc.A,
			Time: bc.GetTime() + uint64(round(MiscGen())),
			Curr: bc.Curr,
		}
	}
	bc.A.EventHeap.Push(next)
}

func (misc *MisDelaysFinishedFC) Visit() {

	// They just finished their delays and are ready to move to security.
	// They're first class, so they need to move to the first class security field.
	misc.A.SecurityFirstClass.Append(misc.Curr)

	// Now, look at see if there's anyone in the first class check in queue
	if GetShortest(misc.A.CheckInFirstClass).Empty() {
		empty := CheckInEmptyFirstClass{
			A:    misc.A,
			Time: misc.Time + uint64(1),
		}
		misc.A.EventHeap.Push(empty) // Add this new event to the heap
		return
	}

	// Else, pull someone off the queue
	bpp := &BoardingPassPrinted{
		A:    misc.A,
		Time: misc.Time + BoardingPassGen(),
		Curr: GetLongest(misc.A.CheckInFirstClass).Pop().(International),
	}
	misc.A.EventHeap.Push(bpp) // Add this new person's check in to the heap
}

func (misc *MiscDelaysFinishedCoach) Visit() {

	// They just finished their delays and are ready to move to security.
	// They're first class, so they need to move to the first class security field.
	misc.A.SecurityCoach.Append(misc.Curr)

	// Check if there's anyone in line in the coach queue
	if GetShortest(misc.A.CheckInCoach).Empty() {
		empty := CheckInEmptyCoach{
			A:    misc.A,
			Time: misc.Time + uint64(1),
		}
		misc.A.EventHeap.Push(empty) // Add this new event to the heap
		return
	}

	// Else, pull someone off the queue
	bpp := &BoardingPassPrinted{
		A:    misc.A,
		Time: misc.Time + BoardingPassGen(),
		Curr: GetLongest(misc.A.CheckInCoach).Pop().(Passenger),
	}
	misc.A.EventHeap.Push(bpp) // Add this new person's check in to the heap
}

func (ci *CheckInEmptyCoach) Visit() {
	// Add 1 to the total time that has been wasted
	ci.A.IdleTime += 1
	ci.A.IdleTimeCoach += 1

	// See if there's someone in line. If not, replicate. Else, pull them off the queue and make a new event for them
	if GetShortest(ci.A.CheckInCoach).Empty() {
		empty := CheckInEmptyCoach{
			A:    ci.A,
			Time: ci.Time + uint64(1),
		}
		ci.A.EventHeap.Push(empty) // Add this new event to the heap
		return
	}

	// Else, pull someone off the queue
	bpp := &BoardingPassPrinted{
		A:    ci.A,
		Time: ci.Time + BoardingPassGen(),
		Curr: GetLongest(ci.A.CheckInCoach).Pop().(Passenger),
	}
	ci.A.EventHeap.Push(bpp) // Add this new person's check in to the heap

}

func (ci *CheckInEmptyFirstClass) Visit() {
	// Add 1 to the total time that has been wasted
	ci.A.IdleTime += 1
	ci.A.IdleTimeFirstClass += 1

	// See if there's someone in line. If not, replicate. Else, pull them off the queue and make a new event for them
	if GetShortest(ci.A.CheckInFirstClass).Empty() {
		empty := CheckInEmptyFirstClass{
			A:    ci.A,
			Time: ci.Time + uint64(1),
		}
		ci.A.EventHeap.Push(empty) // Add this new event to the heap
		return
	}

	// Else, pull someone off the queue
	bpp := &BoardingPassPrinted{
		A:    ci.A,
		Time: ci.Time + BoardingPassGen(),
		Curr: GetLongest(ci.A.CheckInFirstClass).Pop().(International),
	}
	ci.A.EventHeap.Push(bpp) // Add this new person's check in to the heap
}
