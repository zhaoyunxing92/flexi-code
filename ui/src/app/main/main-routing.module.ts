import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { ShareComponent } from './share/share.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full',
  },
  {
    title: 'Login Flexi Code',
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'share',
    component: ShareComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class MainRoutingModule {
}
