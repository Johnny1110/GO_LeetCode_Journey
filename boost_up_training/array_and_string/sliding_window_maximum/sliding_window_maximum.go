package sliding_window_maximum

type MSWIndexMonoDeque struct {
	indexContainer []int
	nums           []int
}

func NewMSWIndexMonoDeque(nums []int) *MSWIndexMonoDeque {
	return &MSWIndexMonoDeque{
		indexContainer: []int{},
		nums:           nums,
	}
}

// Add a new element into indexContainer, start from right:
// if new element is bigger or equals than right side(last) element, remove it
// if new element is lower than
func maxSlidingWindow(nums []int, k int) []int {

	return nil
}
