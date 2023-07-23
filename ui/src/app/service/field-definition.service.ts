import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

import { Text, Textarea, Dropdown, Definition } from '../schema/field';
import { ValueScope } from '../schema/field/definition';

@Injectable({
  providedIn: 'root'
})
export class FieldDefinitionService {

  constructor() {
  }

  /**
   * 获取字段定义
   * @param table 表id
   */
  getFiledDefinitions(table: string): Observable<Definition<any>[]> {

    /**
     * 全部字段
     */
    const fields: Definition<string>[] = [

      new Dropdown({
        id: 'brave',
        name: '爱好',
        size: 2,
        placeholder: '请选择爱好',
        options: [
          {label: '吃饭', value: '1'},
          {label: '编码', value: '2'}
        ],
        order: 2
      }),

      new Text({
        id: 'name',
        name: '姓名',
        required: true,
        order: 1,
        scope: {min: 3, max: 10}
      }),

      new Text({
        id: 'number',
        name: '工号',
        size: 2,
        required: true,
        order: 1,
        scope: {min: 0, max: 25}
      }),

      new Textarea({
        id: 'desc',
        name: '描述',
        order: 3
      })
    ];
    return of(fields.sort((a, b) => a.order - b.order));
  }
}
