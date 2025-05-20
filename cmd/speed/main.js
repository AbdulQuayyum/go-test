function measureExecutionTime(func, ...args) {
  func(...args);
  
  const iterations = 5;
  const times = [];
  
  for (let i = 0; i < iterations; i++) {
    const start = performance.now();
    const result = func(...args);
    const end = performance.now();
    times.push(end - start);
  }
  
  const totalTime = times.reduce((sum, time) => sum + time, 0);
  const avgTime = totalTime / iterations;
  const minTime = Math.min(...times);
  const maxTime = Math.max(...times);
  
  return {
    averageMs: avgTime,
    minMs: minTime,
    maxMs: maxTime,
    totalMs: totalTime,
    iterations: iterations,
    result: func(...args)
  };
}

function exhaustiveLoop(n) {
  let sum = 0;
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      sum += i * j;
    }
  }
  return sum;
}

const jsResult = measureExecutionTime(exhaustiveLoop, 1000);
console.log(`JavaScript execution stats:`, jsResult);


// averageMs : 1.8400000095367433
// iterations : 5
// maxMs : 2.5
// minMs : 1.5
// result : 249500250000
// totalMs : 9.200000047683716