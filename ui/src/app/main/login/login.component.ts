import {Component, OnInit} from '@angular/core';
import {AbstractControl, FormControl, FormGroup, Validators,} from "@angular/forms";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less']
})
export class LoginComponent implements OnInit {
  login!: FormGroup;

  ngOnInit(): void {
    this.login = new FormGroup({
      account: new FormControl<string>("", [Validators.required, Validators.email]),
      password: new FormControl<string>("", [Validators.required, Validators.minLength(8)]),
    }, {updateOn: 'blur'})
  }

  onSubmit(): void {
    console.log(this.login.value)
  }

  getErrorMessage(ctl: any): string {
    console.log(ctl)
    if (ctl?.hasError("required")) {
      return "请输入账号"
    }
    console.log(this.account?.errors)
    return ""
  }

  get account() {
    return this.login.get("account")
  }

  get password() {
    return this.login.get<string>("password")
  }
}
