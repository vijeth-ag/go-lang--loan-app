import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  loanApiUrl = environment.loanApiUrl

  constructor(
    private httpClient: HttpClient
  ) { }

  async getAllLoans(){
    return await this.httpClient.get(this.loanApiUrl + "/loans")
  }

  async approveLoan(userName : string){
    console.log('at apiser approved loan',userName)
    return await this.httpClient.post(this.loanApiUrl+"/approve", {"userName":userName})
  }

  async rejectLoan(userName : string){
    return await this.httpClient.post(this.loanApiUrl + "/reject", {"userName":userName})
  }
}
