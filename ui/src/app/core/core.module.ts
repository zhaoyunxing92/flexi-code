import { NgModule } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WorkbenchModule } from '../workbench/workbench.module';
import { SharedModule } from '../shared/shared.module';
import { ComponentsModule } from '../components/components.module';

@NgModule({
  imports: [
    RouterOutlet,
  ],
  exports: [
    RouterOutlet,
    SharedModule,
    ComponentsModule,
    WorkbenchModule,
  ]
})
export class CoreModule {
}
