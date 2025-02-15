import timeit
import tracemalloc  # a library that measures how much memory the algorithm uses
import matplotlib.pyplot as plt 


def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        for j in range(0, n - i - 1):
            if arr[j] > arr[j+1]:
                arr[j],arr[j+1] = arr[j+1], arr[j]
                

def merge_sort(arr):
    if len(arr) <= 1:
        return arr 
    
    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result


#benchmarking
arr = [i for i in range(1000,0,-1)]

#number is how many time we repeat the function in this case 10 and get the average time
bubble_time = timeit.timeit(lambda: bubble_sort(arr.copy()), number=10)
merge_time = timeit.timeit(lambda: merge_sort(arr.copy()), number=10)


print(f"Bubble Sort Time: {bubble_time:.5f} seconds")
print(f"Merge Sort Time: {merge_time:.5f} seconds")


tracemalloc.start()
bubble_sort(arr.copy())
#result like (0, 8268)   o -> current memory usage in bytes
# 8268 -> peak memory usage in bytes
print("bubble sort memory: ", tracemalloc.get_traced_memory())
tracemalloc.stop()

tracemalloc.start()
merge_sort(arr.copy())
print("Merge sort memory: ",tracemalloc.get_traced_memory())
tracemalloc.stop() 

#analyzing scalability
sizes  = [1000,2000,3000,4000,5000,6000, 7000,8000, 9000,10000]
bubble_times = []
merge_times = []

for size in sizes:
    arr = [i for i in range(size,0 -1)]
    bubble_times.append(timeit.timeit(lambda:bubble_sort(arr.copy()), number = 1))
    merge_times.append(timeit.timeit(lambda: merge_sort(arr.copy()),number =1))

plt.plot(sizes, bubble_times, label="Bubble Sort", marker = "o")
plt.plot(sizes, merge_times, label="Merge Sort", marker="s")
plt.xlabel("Input Size(N):")
plt.ylabel("Time (seconds)")
plt.legend()
plt.show()