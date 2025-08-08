from typing import List


def merge(left: List[int], right: List[int]) -> List[int]:
    result = []
    left_idx, right_idx = 0, 0

    while left_idx < len(left) and right_idx < len(right):
        if left[left_idx] < right[right_idx]:
            result.append(left[left_idx])
            left_idx += 1
        else:
            result.append(right[right_idx])
            right_idx += 1

    result.extend(left[left_idx:])
    result.extend(right[right_idx:])
    return result


def merge_sort(arr: List[int]) -> List[int]:
    if len(arr) < 2:
        return arr

    mid = len(arr) // 2
    left, right = merge_sort(arr[:mid]), merge_sort(arr[mid:])
    return merge(left, right)


def not_in_place_quick_sort(arr):
    if len(arr) < 2:
        return arr

    pivot = arr.pop()
    lower = [x for x in arr if x < pivot]
    greater = [x for x in arr if x > pivot]
    return not_in_place_quick_sort(lower) + [pivot] + not_in_place_quick_sort(greater)
