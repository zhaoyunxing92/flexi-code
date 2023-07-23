import { Text, Textarea } from './textbox';

describe('text', () => {
  it('should create an instance', () => {
    expect(new Text()).toBeTruthy();
  });
  it('should create an instance', () => {
    expect(new Textarea()).toBeTruthy();
  });
});
