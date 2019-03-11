data = [1, 4, 5, 6, 5, 7, 4]

function partition (data, begin,  end) {
    let pivot = data[end]
    let i = begin - 1
    let j = begin 
    
    for (; j < end; ++j) {
        if (data[j] < pivot) {
            i++
            tmp = data[j]
            data[j] = data[i]
            data[i] = tmp
        }
    }
    i++
    tmp = data[j]
    data[j] = data[i]
    data[i] = tmp
    return i
}

function quickSort(data, begin, end) {
    if (begin < end) {
        i = partition(data, begin, end) 
        quickSort(data, begin, i - 1)
        quickSort(data, i + 1, end)
    }
}


console.log(data)

// i = partition(data, 0, data.length - 1)

// console.log(i, data)
quickSort(data, 0, data.length - 1)
console.log(data)

data = [1, 4, 4, 3, 3, 7, 4]
console.log(data)
quickSort(data, 0, data.length - 1)
console.log(data)

data = [1, 3, 4, 5, 7]
console.log(data)
quickSort(data, 0, data.length - 1)
console.log(data)

data = [1, 3, 4, 5, 7].reverse()
console.log(data)
quickSort(data, 0, data.length - 1)
console.log(data)
