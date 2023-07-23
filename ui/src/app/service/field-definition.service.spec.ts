import { TestBed } from '@angular/core/testing';

import { FieldDefinitionService } from './field-definition.service';

describe('FieldDefinitionService', () => {
  let service: FieldDefinitionService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FieldDefinitionService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
