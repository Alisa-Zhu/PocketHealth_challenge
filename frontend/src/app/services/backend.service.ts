import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BackendService {
  getUserFavoriteColor(): Observable <string> {
    // This is where you would make an HTTP request to your backend to retrieve the user's favorite color.
    return of("[USER'S FAVORITE COLOR]");}}