import { NgModule } from '@angular/core';

import { WorkbenchRoutingModule } from './workbench-routing.module';
import { SharedModule } from '../shared/shared.module';

@NgModule({
  declarations: [],
  imports: [
    SharedModule,
    WorkbenchRoutingModule
  ]
})
export class WorkbenchModule { }
