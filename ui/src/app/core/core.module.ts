import { NgModule } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WorkbenchModule } from '../workbench/workbench.module';

@NgModule({
  declarations: [],
  imports: [
    RouterOutlet,
    WorkbenchModule,
  ]
})
export class CoreModule {
}
