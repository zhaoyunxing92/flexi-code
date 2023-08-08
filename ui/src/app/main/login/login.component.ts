import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, Validators, } from '@angular/forms';
import { FieldControlService } from '../../service/field-control.service';
import { LoginService } from '../../service/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.less']
})
export class LoginComponent implements OnInit {
  login!: FormGroup;

  constructor(private loginService: LoginService) {
  }

  ngOnInit(): void {
    this.login = new FormGroup({
      account: new FormControl<string>('', [Validators.required, Validators.email]),
      password: new FormControl<string>('', [Validators.required, Validators.minLength(8)]),
    }, {updateOn: 'change'})
  }

  onSubmit(): void {
    const {account, password} = this.login.value
    this.loginService.login(account, password)
  }

  getErrorMessage(ctl: AbstractControl<string, string>, name: string): string {
    if (ctl?.hasError('required')) {
      return `请输入${name}`
    } else if (ctl?.hasError('email')) {
      return '邮箱格式错误'
    } else if (ctl?.hasError('minlength')) {
      return `${name}长度至少要8位`
    }
    return '参数验证不通过'
  }

  get account(): AbstractControl<string, string> {
    return <AbstractControl>this.login.get('account')
  }

  get password(): AbstractControl<string, string> {
    return <AbstractControl>this.login.get<string>('password')
  }
}
