import { NgModule } from '@angular/core';

import { MainRoutingModule } from './main-routing.module';
import { CoreModule } from '../core/core.module';
import { AppComponent } from './app/app.component';
import { LoginComponent } from './login/login.component';
import { ShareComponent } from './share/share.component';
import { ReactiveFormsModule } from '@angular/forms';
import { NgForOf, NgIf, NgSwitch, NgSwitchCase } from '@angular/common';
import { ComponentsModule } from '../components/components.module';

@NgModule({
  declarations: [AppComponent, LoginComponent, ShareComponent],
  bootstrap: [AppComponent],
  imports: [
    CoreModule,
    MainRoutingModule,
    ReactiveFormsModule,
    NgForOf,
    NgSwitch,
    NgSwitchCase,
    NgIf,
    ComponentsModule
  ]
})
export class MainModule { }
