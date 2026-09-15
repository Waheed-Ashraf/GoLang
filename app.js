

// const employees = [
//   { id: 1, name: 'John Doe', position: 'Software Engineer' },
//   { id: 2, name: 'Jane Smith', position: 'Project Manager' },
//   { id: 3, name: 'Sam Johnson', position: 'UX Designer' },
// ];

// for (const employee of employees) {
//   console.log(`ID: ${employee.id}, Name: ${employee.name}, Position: ${employee.position}`);
// }

//================================

// // 1- Check if a number is even or odd
// function checkEvenOrOdd(number) {
//   if (number % 2 === 0) {
//     return `${number} is Even`;
//   }

//   return `${number} is Odd`;
// }

// console.log(checkEvenOrOdd(8));
// console.log(checkEvenOrOdd(7));

// //================================

// // 2- Fizz Buzz Game
// function fizzBuzz(limit) {
//   for (let i = 1; i <= limit; i++) {
//     if (i % 3 === 0 && i % 5 === 0) {
//       console.log('FizzBuzz');
//     } else if (i % 3 === 0) {
//       console.log('Fizz');
//     } else if (i % 5 === 0) {
//       console.log('Buzz');
//     } else {
//       console.log(i);
//     }
//   }
// }

// fizzBuzz(15);

// //================================

// // 3- Reverse String
// function reverseString(text) {
//   return text.split('').reverse().join('');
// }

// console.log(reverseString('welcome'));

// //=================================

// // 4- Compute Circle Area and Circumference  

// function calculateCircleArea(radius) {
//   const area = Math.PI * radius * radius;
//   return area;
// }

// function calculateCircleCircumference(radius) {
//   const circumference = 2 * Math.PI * radius;
//   return circumference;
// }
// const radius = 6;
// console.log(`Area of Circle: ${calculateCircleArea(radius)}`);
// console.log(`circumference of Circle: ${calculateCircleCircumference(radius)}`);

// //=================================
// // check numbers 

//  function checkNumbers(num1, num2) {
//   if (num1 ===50 || num2 ===50) {
//     return true;
//   } else if (num1 + num2 === 50) {
//     return true;
//   } else {
//     return false;
//   }
//  }
//  console.log(checkNumbers(50, 20)); // true
//   console.log(checkNumbers(20, 30));


  //=================================
  //Compute The sum of the numbers from 1 to 10  

  // function sumOfNumbers(limit){
  //   let sum = 0;
  //   for(i=0 ; i<= limit; i++){
  //      sum += i;
  //   }
  //          return sum;

  // }

  // console.log(sumOfNumbers(10));


  // ==========================

  // // Display right angle triangle pattern
  // function printTriangle(rows) {
  //   for (let i = 1; i <= rows; i++) {
  //     console.log('*'.repeat(i));
  //   }
  // }

  // printTriangle(5);

  // ==========================

  // Factorial with memoization
  const factorialCache = {};

  function factorial(num) {
    if (factorialCache[num] !== undefined) {
      console.log(`${num} already exists in cache`);
      return factorialCache[num];
    }

    let result = 1;

    for (let i = num; i >= 1; i--) {
      result *= i;
    }

    factorialCache[num] = result;
    console.log(`${num} calculated and saved in cache`);
    return result;
  }

  console.log(factorial(5)); 
  console.log(factorial(5)); 
  console.log(factorialCache);



  // ============================ 

  // the output of the first code will be 1 , 2 , 3 
  // the second code will be 3 , 3 , 3 