import { NgModule } from '@angular/core';
import { SharedModule } from '../shared/shared.module';
import { FormComponent } from './form/form.component';

@NgModule({
  declarations: [
    FormComponent
  ],
  exports: [
    FormComponent
  ],
  imports: [
    SharedModule
  ]
})
export class ComponentsModule {
}
