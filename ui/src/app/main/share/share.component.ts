import { Component } from '@angular/core';
import { FieldDefinitionService } from '../../service/field-definition.service';
import { Definition } from '../../schema/field';

@Component({
  selector: 'app-share',
  templateUrl: './share.component.html',
  styleUrls: ['./share.component.less'],
  providers: [FieldDefinitionService]
})
export class ShareComponent {

  fields?: Definition<any>[];

  constructor(service: FieldDefinitionService) {
    service.getFiledDefinitions('').subscribe(data =>
      this.fields = data
    );
  }
}
