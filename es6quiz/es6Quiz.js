// Q1
const numbers1 = [65, 44, 12, 4];
numbers1.forEach(myFunction1);
console.log(numbers1); // Output: [650, 440, 120, 40]

function myFunction1(item, index, arr) {
  arr[index] = item * 10;
}

// Q2
const pokemon = ["squirtle", "charmander", "bulbasaur"];

const pokeLength = pokemon.reduce(function (previous, current) {
  return previous + current.length;
}, 0);

console.log(pokeLength); // Output: 27

// Q3
const numbers3 = [5, 10, 15];

const reducer = numbers3.reduce((accumulator, item) => {
  return accumulator + item;
});

console.log(reducer); // Output: 30

// Q4
const euros = [29.76, 41.85, 46.5];

const doubled = euros.reduce((total, amount) => {
  total.push(amount * 2);
  return total;
}, []);

console.log(doubled); // Output: [59.52, 83.7, 93]

// Q5
const numbers5 = [1, 2, 3, 4, 5];
console.log(numbers5.includes(2)); // Output: true
console.log(numbers5.includes(99)); // Output: false

// Q6
const myBoolean = true;

if (myBoolean) {
  const turtles = ["Leonardo", "donatello", "michaelangelo", "raphael"];

  // turtles = turtles.concat("Shredder");
  // console.log(turtles);
  // Output: TypeError: Assignment to constant variable.
}

// Q7
const names = ["Batman", "Catwoman", "Joker", "Bane"];
const fromIndex = 1;
const removeCount = 2;
const newNames = [
  ...names.slice(0, fromIndex),
  ...names.slice(fromIndex + removeCount),
];

console.log(newNames); // Output: ["Batman", "Bane"]

// Q8
console.log(typeof NaN); // Output: "number"
console.log(typeof String); // Output: "function"
console.log(typeof undefind); // Output: "undefined"
console.log(typeof null); // Output: "object"
console.log(typeof [5, 10, 20]); // Output: "object"

// Q9
function showCoords(event) {
  document.getElementById("demo").innerHTML = `
    <p> X = ${event.clientX}</p>
    <p> Y = ${event.clientY}</p>`;
}

// Output: no console output. When called by a mouse event, it updates #demo
// with the event's clientX and clientY values.

// Q10
const person = {
  profile: {
    name: "",
    age: 0,
  },
};

console.log(person.profile.name || "Anonymous"); // Output: "Anonymous"
console.log(person.profile.age || 18); // Output: 18
console.log(person.profile.name ?? "Anonymous"); // Output: ""
console.log(person.profile.age ?? 18); // Output: 0

// Q11
const colors = ["white", "black", "gray"];
const clone = [...colors];
console.log(clone); // Output: ["white", "black", "gray"]
console.log(colors === clone); // Output: false

// Q12
const numbers12 = [1, 2, 3, 4, 5];
const nums = [];

function isEven(number) {
  return number % 2 === 0;
}

const evenNumber = numbers12.find(isEven);
const evenNum = nums.find(isEven);
console.log(evenNumber); // Output: 2
console.log(evenNum); // Output: undefined

// Q13
// let myFunc = (first, last) => ({ firstName: first, lastName: last }),
//   testFunc = (first, last) => { firstName: first, lastName: last };
//
// console.log(myFunc("john", "doe"));
// console.log(testFunc("john", "doe"));
// Output: SyntaxError: Unexpected token ":"

// Q14
function mul(num1) {
  return function (num2) {
    return function (num3) {
      return num1 * num2 * num3;
    };
  };
}

console.log(mul(1)(5)(10)); // Output: 50

// Q15
const arr = ["john", "jack", "john", "jack"];

const result = arr.reduce((x, y) => {
  if (!x[y]) {
    x[y] = 0;
  }

  x[y]++;
  return x;
}, {});

console.log(result); // Output: { john: 2, jack: 2 }
