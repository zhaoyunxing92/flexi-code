/**
 * 字段定义
 */
export class Definition<T> {

  /**
   * 字段的唯一标识
   */
  id: string;

  /**
   * 字段值
   */
  value: T | undefined;

  /**
   *字段名称
   */
  name: string;

  /**
   * 别名
   */
  alias: string;

  /**
   * 图标
   */
  icon?: string;

  /**
   * 提示语
   */
  placeholder: string;

  /**
   * 是否必填
   */
  required: boolean;

  /**
   * 是否禁止
   */
  disabled: boolean;

  /**
   * 字段取值范围
   */
  scope?: ValueScope;

  /**
   * 字段顺序
   */
  order: number;

  /**
   * 字段类型
   */
  type: string;

  /**
   * 格式化value 值
   */
  format: string;

  /**
   * 选项值
   */
  options: { label: string, value: string }[];

  /**
   * 字段大小 1,2,3,4 四种规格
   */
  size: number;

  constructor(options: {
    id?: string;
    value?: T;
    alias?: string;
    icon?: string;
    name?: string;
    required?: boolean;
    disabled?: boolean;
    order?: number;
    size?: number;
    controlType?: string;
    type?: string;
    format?: string;
    placeholder?: string;
    scope?: ValueScope;
    options?: { label: string, value: string }[];
  } = {}) {
    this.value = options.value;
    this.id = options.id || '';
    this.icon = options.icon || '';
    this.alias = options.alias || this.id;
    this.name = options.name || '';
    this.placeholder = options.placeholder || `请输入${this.name}`
    this.required = !!options.required;
    this.disabled = !!options.disabled;
    this.order = options.order === undefined ? 1 : options.order;
    this.type = options.type || '';
    this.format = options.format || '';
    this.scope = options.scope || {min: 0, max: 50}
    this.options = options.options || [];
    this.size = options.size || 4;
  }
}

export class ValueScope {

  /**
   * 最小值
   */
  min?: number;

  /**
   * 最大值
   */
  max?: number;

  constructor(options: { min: number, max: number }) {
    this.min = options.min;
    this.max = options.max;
  }
}


