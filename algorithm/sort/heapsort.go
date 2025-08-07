package main

import "fmt"

func main() {
	// 定义一个切片，这里我们模拟 10 个元素
	arr := []int{84, 353, 596, 848, 425, 849, 166, 521, 228, 573}
	fmt.Println("Original data:", arr) // 先打印原始数据
	HeapSort(arr)                      // 调用堆排序
	fmt.Println("New data:  ", arr)    // 后打印排序后的数据
}

// HeapSort 堆排序
func HeapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一个个从堆顶取出元素
	for i := n - 1; i >= 0; i-- {
		// 移动当前根到末尾
		arr[i], arr[0] = arr[0], arr[i]

		// 调用 max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// heapify 将以 i 为根的子树调整为最大堆
func heapify(arr []int, n int, i int) {
	largest := i // 初始化最大为根
	l := 2*i + 1 // 左子节点
	r := 2*i + 2 // 右子节点

	// 如果左子节点大于根
	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	// 如果右子节点大于当前的最大值
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// 如果最大值不是根
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 交换

		// 递归地堆化受影响的子树
		heapify(arr, n, largest)
	}
}

//参考链接：https://www.cnblogs.com/realcp1018/p/17954463
