import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class IssueService {
  url = 'http://localhost:4200'
  constructor(private http: HttpClient) { }

  getIssues(){
    return this.http.get(`${this.url}/issues`);
  }
}
