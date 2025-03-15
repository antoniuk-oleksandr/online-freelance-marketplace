export type DropdownItem = {
  title: string,
  clickAction: () => void
  closeDropdown?: boolean
  badge?: string
  icon?: string
  customColor?: string
  dividerAfter?: boolean
  dividerBefore?: boolean
}
