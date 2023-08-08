import { NgModule } from '@angular/core';
import { CoreModule } from '../core/core.module';
import { AppComponent } from './app/app.component';
import { LoginComponent } from './login/login.component';
import { ShareComponent } from './share/share.component';

@NgModule({
  declarations: [AppComponent, LoginComponent, ShareComponent],
  bootstrap: [AppComponent],
  imports: [
    CoreModule
  ]
})
export class MainModule { }
