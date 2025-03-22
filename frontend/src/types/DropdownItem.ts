export type DropdownItem = {
  title: string,
  clickAction: () => void,
  groupId?: number
  color?: string,
  closeDropdown?: boolean,
  badge?: string,
  icon?: string,
  customColor?: string,
  iconColor?: string,
  dividerAfter?: boolean,
  dividerBefore?: boolean,
}
