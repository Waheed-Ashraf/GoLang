"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class Employee {
    id;
    firstName;
    lastName;
    salary;
    constructor(id, firstName, lastName, salary) {
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.salary = salary;
    }
    getId() {
        return this.id;
    }
    getFirstName() {
        return this.firstName;
    }
    getLastName() {
        return this.lastName;
    }
    getSalary() {
        return this.salary;
    }
    setSalary(salary) {
        this.salary = salary;
    }
    getName() {
        return `${this.firstName} ${this.lastName}`;
    }
    getAnnualSalary() {
        return this.salary * 12;
    }
    raiseSalary(percent) {
        this.salary += this.salary * (percent / 100);
        return this.salary;
    }
    toString() {
        return `Employee[id=${this.id}, name=${this.firstName} ${this.lastName}, salary=${this.salary}]`;
    }
}
const e1 = new Employee(8, "Peter", "Tan", 2500);
console.log(e1.toString());
e1.setSalary(999);
console.log(e1.toString());
console.log("id is: " + e1.getId());
console.log("firstname is: " + e1.getFirstName());
console.log("lastname is: " + e1.getLastName());
console.log("salary is: " + e1.getSalary());
console.log("name is: " + e1.getName());
console.log("annual salary is: " + e1.getAnnualSalary());
console.log(e1.raiseSalary(10));
console.log(e1.toString());
//# sourceMappingURL=task1-4.js.map