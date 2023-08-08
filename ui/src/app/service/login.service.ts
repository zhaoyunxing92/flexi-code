import { Injectable } from '@angular/core';
import { TokenService } from './token.service';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private tokenService: TokenService) {
  }

  login(email: string, password: string): void {
    console.log(email, password)
    this.tokenService.token = 'flexi'
  }
}
