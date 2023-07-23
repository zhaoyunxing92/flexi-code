import { Definition, ValueScope } from './definition';

/**
 * 单行文本框
 */
export class Text extends Definition<string> {
  override type: string = 'text'
}

/**
 * 多行文本框
 */
export class Textarea extends Definition<string> {
  override type: string = 'textarea'
}
