def insertion_sort(arr):
    for i in range(len(arr)):
        j = i
        while j > 0:
            if arr[j-1] > arr[j]:
                current = arr[j]
                arr[j] = arr[j-1]
                arr[j-1] = current
            j = j - 1
    return arr
 

