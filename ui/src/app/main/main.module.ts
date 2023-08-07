import { NgModule } from '@angular/core';

import { MainRoutingModule } from './main-routing.module';
import { CoreModule } from '../core/core.module';
import { AppComponent } from './app/app.component';
import { LoginComponent } from './login/login.component';
import { ShareComponent } from './share/share.component';
import { ReactiveFormsModule } from '@angular/forms';
import { NgForOf, NgIf, NgSwitch, NgSwitchCase } from '@angular/common';
import { ComponentsModule } from '../components/components.module';
import {MatCardModule} from "@angular/material/card";
import {MatButtonModule} from "@angular/material/button";
import {MatGridListModule} from "@angular/material/grid-list";
import {MatInputModule} from "@angular/material/input";

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
    ComponentsModule,
    MatCardModule,
    MatButtonModule,
    MatGridListModule,
    MatInputModule
  ]
})
export class MainModule { }
