function multiply(a, b) {
    return a * b
}


function yell(sentence){
    return sentence.toUpperCase();
}

function longerThan(arr1, arr2) {
           return arr1.length > arr2.length
}


console.log(multiply(10, 3))
console.log(yell("Hello"))
arr1 = [2,3, 4, 5, 9]
arr2 = [2,3,4,5]
console.log(longerThan(arr1, arr2))