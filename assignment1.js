// Question 1: Answer the following questions

// 1. What is an Event Loop?
// The Event Loop is the JavaScript runtime mechanism that keeps checking whether
// the call stack is empty, then moves ready asynchronous callbacks from queues
// into the stack so they can run.

// 2. How do you add an element at the beginning and end of an array?
const names = ["Ahmed", "Mona"];

// Beginning:
names.unshift("Ali");
names.splice(0, 0, "Sara");

// End:
names.push("Omar");
names.splice(names.length, 0, "Nour");

console.log("Q1 array methods:", names);

// ========================================================================
// Question 2: What is the output?

// 1.
// var b = 1;
// function outer() {
//   var b = 2;
//   function inner() {
//     b++;
//     var b = 3;
//     console.log(b);
//   }
//   inner();
// }
// outer();
// Output: 3

// 2.
// for (let i = 0; i < 5; i++) {
//   setTimeout(function () {
//     console.log(i);
//   }, i * 1000);
// }
// Output: 0, 1, 2, 3, 4
// Each number prints after its own delay because let creates a new i for each loop iteration.

// 3.
// let arr = ["foo", "bar"];
// arr.length = 0;
// arr.push("baz");
// console.log(arr);
// Output: ["baz"]

// 4.
// function func() {
//   for (let key in arguments) {
//     console.log(arguments[key]);
//   }
// }
// func(1, "Hello", true);
// Output: 1, "Hello", true

// 5.
// let car = {
//   carName: "Bmw",
//   carPrice: 1000000,
// };
// console.log(car instanceof Object);
// console.log(Object.entries(car));
// Output:
// true
// [["carName", "Bmw"], ["carPrice", 1000000]]

// ========================================================================
// Question 3

// 1. Create sumObjectValues() that sums all own numeric property values.
function sumObjectValues(obj) {
  let sum = 0;

  for (const key in obj) {
    if (
      Object.prototype.hasOwnProperty.call(obj, key) &&
      typeof obj[key] === "number" &&
      Number.isFinite(obj[key])
    ) {
      sum += obj[key];
    }
  }

  return sum;
}

const invoice = {
  subtotal: 200,
  tax: 30,
  customer: "John",
  discount: 20,
};

console.log("Q3-1:", sumObjectValues(invoice));

// 2. Show 3 asynchronous blocks of code one after the other in sequence.
function delay(message, milliseconds) {
  return new Promise(function (resolve) {
    setTimeout(function () {
      console.log(message);
      resolve();
    }, milliseconds);
  });
}

async function runSequentialAsyncBlocks() {
  await delay("Q3-2: First async block", 100);
  await delay("Q3-2: Second async block", 100);
  await delay("Q3-2: Third async block", 100);
}

runSequentialAsyncBlocks();

// 3. Get the maximum value from a numbers array along with its index.
function getMaxWithIndex(numbers) {
  if (numbers.length === 0) {
    return null;
  }

  let max = numbers[0];
  let index = 0;

  for (let i = 1; i < numbers.length; i++) {
    if (numbers[i] > max) {
      max = numbers[i];
      index = i;
    }
  }

  return { max, index };
}

console.log("Q3-3:", getMaxWithIndex([10, 50, 30, 90, 40]));

// 4. Return the difference between two valid dates as number of days.
function getDaysDifference(firstDate, secondDate) {
  const oneDay = 24 * 60 * 60 * 1000;
  const start = new Date(firstDate);
  const end = new Date(secondDate);

  return Math.abs(end - start) / oneDay;
}

console.log("Q3-4:", getDaysDifference("2026-08-01", "2026-08-17"));

// 5. Calculator interface for 2 number inputs.
function createCalculator(firstNumber, secondNumber) {
  return {
    sum: function () {
      return firstNumber + secondNumber;
    },
    difference: function () {
      return firstNumber - secondNumber;
    },
    product: function () {
      return firstNumber * secondNumber;
    },
    dividend: function () {
      if (secondNumber === 0) {
        return "Cannot divide by zero";
      }

      return firstNumber / secondNumber;
    },
  };
}

const calculator = createCalculator(20, 5);
console.log("Q3-5 sum:", calculator.sum());
console.log("Q3-5 difference:", calculator.difference());
console.log("Q3-5 product:", calculator.product());
console.log("Q3-5 dividend:", calculator.dividend());

// 6. Return multiple values from a function.
function getNumberStats(numbers) {
  const sum = numbers.reduce(function (total, number) {
    return total + number;
  }, 0);

  return {
    sum,
    count: numbers.length,
    average: sum / numbers.length,
  };
}

console.log("Q3-6:", getNumberStats([10, 20, 30]));

// 7. Reverse an array.
function reverseArray(arr) {
  const reversed = [];

  for (let i = arr.length - 1; i >= 0; i--) {
    reversed.push(arr[i]);
  }

  return reversed;
}

console.log("Q3-7:", reverseArray([1, 2, 3, 4]));

// 8. Convert an object into an array of key-value pairs.
function objectToArray(obj) {
  return Object.entries(obj);
}

console.log("Q3-8:", objectToArray({ a: 1, b: 2 }));

// ========================================================================
// Bonus

// 1. Convert 12-hour time format to 24-hour time format.
function convertTo24Hour(time) {
  const match = time.trim().match(/^(\d{1,2}):(\d{2})(?::(\d{2}))?\s*(AM|PM)$/i);

  if (!match) {
    return "Invalid time format";
  }

  let hours = Number(match[1]);
  const minutes = match[2];
  const seconds = match[3];
  const period = match[4].toUpperCase();

  if (period === "AM" && hours === 12) {
    hours = 0;
  } else if (period === "PM" && hours !== 12) {
    hours += 12;
  }

  const formattedHours = String(hours).padStart(2, "0");
  return seconds
    ? `${formattedHours}:${minutes}:${seconds}`
    : `${formattedHours}:${minutes}`;
}

console.log("Bonus-1:", convertTo24Hour("07:45 PM"));
console.log("Bonus-1:", convertTo24Hour("12:10 AM"));

// 2. Make this syntax possible: var a = add(2)(3); // 5
function add(firstNumber) {
  return function (secondNumber) {
    return firstNumber + secondNumber;
  };
}

const a = add(2)(3);
console.log("Bonus-2:", a);

// 3. Check if the user with the name "John" exists in the array of objects.
const users = [
  { id: 1, name: "Ahmed" },
  { id: 2, name: "John" },
  { id: 3, name: "Mona" },
];

function hasUserNamedJohn(usersList) {
  return usersList.some(function (user) {
    return user.name === "John";
  });
}

console.log("Bonus-3:", hasUserNamedJohn(users));
