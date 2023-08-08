import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TokenService {
  readonly key: string = 'token';

  set token(token: string) {
    localStorage.setItem(this.key, token)
  }

  get token(): string {
    return localStorage.getItem(this.key) || ''
  }

  get reject(): boolean {
    return this.token.length === 0;
  }
}
