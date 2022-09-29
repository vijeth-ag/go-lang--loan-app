import { Component, OnInit } from '@angular/core';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-loan',
  templateUrl: './loan.component.html',
  styleUrls: ['./loan.component.css']
})
export class LoanComponent implements OnInit {
  
  loanApplications: any;

  resp  = [
    {"ID":"632f125c21bd1b0f5af1cafa","Username":"ram2@gmail.com","FirstName":"Ram","Lastname":"g","MonthlySalary":1020,"LoanAmount":111,"Status":{"StatusName":"applied","Reason":"none"},"LoanAppliedDate":"1664029276"},
    
    {"ID":"632f125c21bd1b0f5af1cafa","Username":"vijeth2@gmail.com","FirstName":"Vijeth","Lastname":"ag","MonthlySalary":100,"LoanAmount":11,"Status":{"StatusName":"applied","Reason":"none"},"LoanAppliedDate":"1664029276"}
    ]

  constructor(
    private apiService: ApiService
  ) { }

  ngOnInit(): void {
    console.log('loan')
    this.getAllLoans()
  }

  async getAllLoans(){
    (await this.apiService.getAllLoans()).subscribe((data) => {
      // this.loanApplications = data;
      this.loanApplications = this.resp
      console.log('loanApplications:',this.loanApplications)
    })
  }

  async approve(loan: any){
    (await this.apiService.approveLoan(loan["Username"])).subscribe((data) => {
      console.log('loan approved',data)
    })
  }

  async reject(loan: any){
    (await this.apiService.rejectLoan(loan["Username"])).subscribe((data) => {
      console.log('loan rejected',data)
    })
  }  
}
