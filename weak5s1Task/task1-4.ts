

class Employee {
  private id: number;
  private firstName: string;
  private lastName: string;
  private salary: number;

  constructor(id: number, firstName: string, lastName: string, salary: number) {
    this.id = id;
    this.firstName = firstName;
    this.lastName = lastName;
    this.salary = salary;
  }

  getId(): number {
    return this.id;
  }

  getFirstName(): string {
    return this.firstName;
  }

  getLastName(): string {
    return this.lastName;
  }

  getSalary(): number {
    return this.salary;
  }

  setSalary(salary: number): void {
    this.salary = salary;
  }

  getName(): string {
    return `${this.firstName} ${this.lastName}`;
  }

  getAnnualSalary(): number {
    return this.salary * 12;
  }

  raiseSalary(percent: number): number {
    this.salary += this.salary * (percent / 100);
    return this.salary;
  }

  toString(): string {
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

