import { Component } from '@angular/core';
import { OnInit } from '@angular/core';
import { BackendService } from './backend.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  name = '[USER\'S NAME]';
  email = '[USER\'S EMAIL]';
  favoriteColor = '[USER\'S FAVORITE COLOR]';

  constructor(private backendService: BackendService) { }

  ngOnInit() {
    // Retrieve user's favorite color from the backend
    this.backendService.getUserFavoriteColor().subscribe((color: string) => {
      this.favoriteColor = color;

      // Update background color to user's favorite color
      document.body.style.backgroundColor = this.favoriteColor;
    });
  }}
