def merge(left, right):
    i=0
    j=0
    k = 0
    
    result = []
    while (i < len(left)) and (j < len(right)):
        if left[i] < right[j]:
            result.insert(k, left[i])
            i = i + 1
        else:
            result.insert(k, right[j])
            j = j +1

        k = k + 1
    
    while i < len(left):
        result.insert(k, left[i])
        i = i + 1
        k = k +1

    while j < len(right):
        result.insert(k, right[j])
        j = j + 1
        k = k +1

    return result
    
def merge_sort(arr):
    if len(arr) <= 1:
        return arr
    
    mid = len(arr)//2
    left = arr[:mid]
    right = arr[mid:]

    left = merge_sort(left)
    right = merge_sort(right)
    return merge(left, right)




    