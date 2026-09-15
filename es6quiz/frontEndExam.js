// Front End Exam JavaScript solutions.
// YouTube clone task is skipped.

// Question 4: outputs

// Q4.1
var p = new Promise((resolve, reject) => {
  reject(Error("The Fails!"));
});

p.catch((error) => console.log(error)); // Output: Error: The Fails!
p.catch((error) => console.log(error.message)); // Output: The Fails!
p.catch((error) => console.log(error.message)); // Output: The Fails!

// Q4.2
function Person(firstName, lastName) {
  this.firstName = firstName;
  this.lastName = lastName;
}

const member = new Person("Lydia", "Hallie");

Person.getFullName = function () {
  return `${this.firstName} ${this.lastName}`;
};

// console.log(member.getFullName());
// Output: TypeError: member.getFullName is not a function

// Q4.3
const SumBy = (num1) => (num2) => num1 + num2;
const sumByTwo = SumBy(2);
const sumByThree = SumBy(3);

console.log(sumByTwo(4)); // Output: 6
console.log(sumByThree(5)); // Output: 8

// Q4.4
class Chameleon {
  static colorChange(newColor) {
    this.newColor = newColor;
    return this.newColor;
  }

  constructor(newColor) {
    this.newColor = newColor;
  }
}

const freddie = new Chameleon("Purple");
// console.log(freddie.colorChange("orange"));
// Output: TypeError: freddie.colorChange is not a function

// Q4.5
const sampleAge = 20;
const accessAllowed = sampleAge >= 18 ? true : false;
console.log(typeof accessAllowed); // Output: boolean

function greeting() {
  return "Welcome All";
}

console.log(typeof greeting()); // Output: string

// Q4.6
setTimeout(function () {
  setTimeout(function () {
    console.log(2); // Output order: 4
    setTimeout(function () {
      console.log(3); // Output order: 5
    }, 0);
  }, 1000);

  setTimeout(function () {
    console.log(4); // Output order: 3
  });

  console.log(1); // Output order: 2
}, 2000);

console.log(0); // Output order: 1

// Full output order: 0, 1, 4, 2, 3

// Q4.7
function counter() {
  var i = 0;
  return ++i;
}

// console.log(i);
// Output: ReferenceError: i is not defined

// Q4.8
let obj = {
  msg: "hello world",
  x: 10,
};

var x = "msg";

console.log(obj[x]); // Output: hello world
console.log(obj["x"]); // Output: 10

// Q4.9
const euros = [29.76, 41.85, 46.5];

const doubled = euros.reduce((total, amount) => {
  total.push(amount * 2);
  return total;
}, []);

console.log(doubled); // Output: [59.52, 83.7, 93]

// Q4.10
const names = ["Batman", "Catwoman", "Joker", "Bane"];
const fromIndex = 1;
const removeCount = 2;
const newNames = [
  ...names.slice(0, fromIndex),
  ...names.slice(fromIndex + removeCount),
];

console.log(newNames); // Output: ["Batman", "Bane"]

// Question 5.1
function recursiveLength(str) {
  str = String(str);

  if (str === "") {
    return 0;
  }

  return 1 + recursiveLength(str.slice(1));
}

console.log(recursiveLength("hello")); // Output: 5

// Question 5.2
function printMultiplicationTable(number = 12) {
  for (let i = 1; i <= 12; i++) {
    console.log(`${number} * ${i} = ${number * i}`);
  }
}

printMultiplicationTable(12);

// Question 5.3
function elementsOnOddPositions(list) {
  return list.filter((item, index) => index % 2 === 0);
}

console.log(elementsOnOddPositions(["a", "b", "c", "d", "e"]));
// Output: ["a", "c", "e"]

// Question 5.4
function isPrime(number) {
  if (number <= 1 || !Number.isInteger(number)) {
    return false;
  }

  for (let i = 2; i <= Math.sqrt(number); i++) {
    if (number % i === 0) {
      return false;
    }
  }

  return true;
}

console.log(isPrime(17)); // Output: true
console.log(isPrime(18)); // Output: false

// Question 5.5: background generator logic
function updateBackground(colorOne, colorTwo) {
  const gradient = `linear-gradient(135deg, ${colorOne}, ${colorTwo})`;

  if (typeof document !== "undefined") {
    document.body.style.background = gradient;
    document.getElementById("css-output").textContent = `background: ${gradient};`;
  }

  return gradient;
}

if (typeof document !== "undefined") {
  const colorOne = document.getElementById("color-one");
  const colorTwo = document.getElementById("color-two");

  if (colorOne && colorTwo) {
    const handleColorChange = () => updateBackground(colorOne.value, colorTwo.value);

    colorOne.addEventListener("input", handleColorChange);
    colorTwo.addEventListener("input", handleColorChange);
    handleColorChange();
  }
}

// Question 5.6
function countVowels(str) {
  const matches = String(str).match(/[aeiou]/gi);
  return matches ? matches.length : 0;
}

console.log(countVowels("hello world")); // Output: 3

// Question 5.7
class Animal {
  setValue(name, age) {
    this.name = name;
    this.age = age;
  }
}

class Zebra extends Animal {
  info() {
    return `${this.name} is ${this.age} years old and comes from Africa.`;
  }
}

class Dolphin extends Animal {
  info() {
    return `${this.name} is ${this.age} years old and lives in the ocean.`;
  }
}

const zebra = new Zebra();
zebra.setValue("Marty", 6);
console.log(zebra.info());

const dolphin = new Dolphin();
dolphin.setValue("Flipper", 8);
console.log(dolphin.info());

// Question 5.9
function MyObject(name, message) {
  this.name = String(name);
  this.message = String(message);
}

MyObject.prototype.getName = function () {
  return this.name;
};

MyObject.prototype.getMessage = function () {
  return this.message;
};

class BetterObject {
  constructor(name, message) {
    this.name = String(name);
    this.message = String(message);
  }

  getName() {
    return this.name;
  }

  getMessage() {
    return this.message;
  }
}

// Question 5.10
class Shape {
  constructor(color = "red", filled = true) {
    this.color = color;
    this.filled = filled;
  }

  getColor() {
    return this.color;
  }

  setColor(color) {
    this.color = color;
  }

  isFilled() {
    return this.filled;
  }

  setFilled(filled) {
    this.filled = filled;
  }

  toString() {
    return `Shape[color=${this.color},filled=${this.filled}]`;
  }
}

class Circle extends Shape {
  constructor(radius = 1.0, color = "red", filled = true) {
    super(color, filled);
    this.radius = radius;
  }

  getRadius() {
    return this.radius;
  }

  setRadius(radius) {
    this.radius = radius;
  }

  getArea() {
    return Math.PI * this.radius * this.radius;
  }

  getPerimeter() {
    return 2 * Math.PI * this.radius;
  }

  toString() {
    return `Circle[${super.toString()},radius=${this.radius}]`;
  }
}

class Rectangle extends Shape {
  constructor(width = 1.0, length = 1.0, color = "red", filled = true) {
    super(color, filled);
    this.width = width;
    this.length = length;
  }

  getWidth() {
    return this.width;
  }

  setWidth(width) {
    this.width = width;
  }

  getLength() {
    return this.length;
  }

  setLength(length) {
    this.length = length;
  }

  getArea() {
    return this.width * this.length;
  }

  getPerimeter() {
    return 2 * (this.width + this.length);
  }

  toString() {
    return `Rectangle[${super.toString()},width=${this.width},length=${this.length}]`;
  }
}

class Square extends Rectangle {
  constructor(side = 1.0, color = "red", filled = true) {
    super(side, side, color, filled);
  }

  getSide() {
    return this.width;
  }

  setSide(side) {
    this.width = side;
    this.length = side;
  }

  setWidth(side) {
    this.setSide(side);
  }

  setLength(side) {
    this.setSide(side);
  }

  toString() {
    return `Square[${super.toString()}]`;
  }
}

const circle = new Circle(2, "blue", false);
console.log(circle.toString());
console.log(circle.getArea());

const square = new Square(4, "green", true);
console.log(square.toString());
console.log(square.getArea());
