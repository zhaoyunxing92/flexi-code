import { Component, Input, OnInit } from '@angular/core';
import { Definition } from '../../schema/field';
import { FormGroup } from '@angular/forms';
import { FieldControlService } from '../../service/field-control.service';

@Component({
  selector: 'fc-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.less'],
  providers: [FieldControlService]
})
export class FormComponent implements OnInit {
  @Input()
  fields?: Definition<string>[];

  form!: FormGroup;

  constructor(private fieldService: FieldControlService) {
    this.fields = []
  }

  ngOnInit(): void {
    this.form = this.fieldService.toFormGroup(this.fields as Definition<string>[])
  }

  onSubmit(): void {
    console.log(this.form.value)
  }

}
