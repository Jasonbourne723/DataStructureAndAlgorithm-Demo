package DataStructureAndAlgorithm

type DataStore struct {
	array []int
}

func NewDataStore(array []int) *DataStore {
	return &DataStore{array: array}
}

func (d *DataStore) BinarySearch(data int) (index int, ok bool) {
	return d.binarySearch(0, len(d.array)-1, data)
}

func (d *DataStore) binarySearch(low int, high int, searchData int) (index int, ok bool) {
	if low > high {
		return index, false
	}
	mid := low + (high-low)/2
	if searchData == d.array[mid] {
		return mid, true
	}
	if searchData > d.array[mid] {
		return d.binarySearch(mid+1, high, searchData)
	} else {
		return d.binarySearch(low, mid-1, searchData)
	}

}
