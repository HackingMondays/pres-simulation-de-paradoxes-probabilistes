var MAX_DAYS = 365;

function hasDuplicates(data) {
    var cache = {};
    for (var i = 0; i < data.length; i++) {
        if (cache[data[i]] === true) {
            return true;
        }
        cache[data[i]] = true;
    }
    return false;
}

function getRandomArray(size) {
    var data = [];
    for (var i=0; i<size; i++) {
        data.push(Math.floor(Math.random()*MAX_DAYS));
    }
    return data;
}

var success = 0;
var total = 100000;
var count = 22;

//noprotect
for (var i=0; i<total; i++) {
    var samples = getRandomArray(count);
    success += hasDuplicates(samples) ? 1 : 0;
}

console.log("Population size: " + count);
console.log("Iterations     : " + total);
console.log("Same birthday  : " + success);
console.log("Rate           : " + ((success / total)*100).toFixed(2) + "%");

