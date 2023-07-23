import { Injectable } from '@angular/core';
import { Definition } from '../schema/field';
import { FormControl, FormGroup, ValidatorFn, Validators } from '@angular/forms';

@Injectable()
export class FieldControlService {

  /**
   * 字段定义转换为表单
   *
   * @param fields 字段定义
   */
  toFormGroup(fields: Definition<string>[]): FormGroup {
    const group: any = {};
    fields.forEach((def: Definition<string>): void => {
      let opts: Validators[] = [] || null;
      if (def.required) {
        opts.push(Validators.required)
      }
      if (def.scope) {
        let min = def.scope?.min;
        if (min) {
          opts.push(Validators.minLength(min))
        }
        let max = def.scope?.max;
        if (max) {
          opts.push(Validators.minLength(max))
        }
      }
      // @ts-ignore
      group[ def.id ] = new FormControl(def.value || '', opts);
    });
    return new FormGroup(group);
  }
}
