import { NgModule } from '@angular/core';
import { SharedModule } from '../shared/shared.module';
import { FormComponent } from './form/form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { NgForOf, NgSwitch, NgSwitchCase } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatGridListModule } from '@angular/material/grid-list';

@NgModule({
  declarations: [
    FormComponent
  ],
  exports: [
    FormComponent
  ],
  imports: [
    SharedModule,
    ReactiveFormsModule,
    NgForOf,
    NgSwitch,
    NgSwitchCase,
    MatInputModule,
    MatSelectModule,
    MatGridListModule
  ]
})
export class ComponentsModule {
}
