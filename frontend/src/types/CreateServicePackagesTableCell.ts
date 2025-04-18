export type CreateServicePackagesTableCell = {
  type: 'text',
  text?: string,
  styles?: string,
} | {
  type: 'textarea',
  name: string,
  rightIcon?: string,
  styles?: string,
  placeholder?: string,
} | {
  type: 'input',
  name: string,
  inputType?: 'text' | 'number',
  rightBadge?: string,
  styles?: string,
  placeholder?: string,
}
