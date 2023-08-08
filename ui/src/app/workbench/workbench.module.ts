import { NgModule } from '@angular/core';

import { WorkbenchRoutingModule } from './workbench-routing.module';
import { SharedModule } from '../shared/shared.module';
import { MainComponent } from './main/main.component';
import { RouterOutlet } from '@angular/router';

@NgModule({
  declarations: [
    MainComponent
  ],
  imports: [
    RouterOutlet,
    SharedModule,
    WorkbenchRoutingModule
  ]
})
export class WorkbenchModule { }
