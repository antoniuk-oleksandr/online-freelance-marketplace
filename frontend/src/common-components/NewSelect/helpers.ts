import type { DropdownItem } from "@/types/DropdownItem"
import type { NewSelectItem } from "@/types/NewSelectItem"

export const makeSelectValues = (
  initialSelectValues: NewSelectItem[],
  selectedIndexes: number[],
  setSelectedIndexes: (value: number[]) => void,
): DropdownItem[] => initialSelectValues.map((item, index) => {
  const includes = selectedIndexes.includes(index);

  return {
    ...item,
    icon: includes ? 'akar-icons:radio-fill' : 'akar-icons:radio',
    iconColor: includes ? 'cyan-500' : '',
    clickAction: () => updateSelectIndex(index, selectedIndexes, setSelectedIndexes, item.groupId),
  }
})

const updateSelectIndex = (
  index: number,
  selectedIndexes: number[],
  setSelectedIndexes: (value: number[]) => void,
  groupId?: number,
) => {
  if (!groupId) groupId = 0

  selectedIndexes[groupId] = index
  setSelectedIndexes([...selectedIndexes])
}

export const setInitialSelectIndex = (
  initialSelectIndex?: number,
  multipleInitialSelectIndexes?: number[]
): number[] => {
  if (multipleInitialSelectIndexes !== undefined) return multipleInitialSelectIndexes
  if (initialSelectIndex !== undefined) return [initialSelectIndex]
  return [0]
}

export const performOnSelectChange = (
  items: DropdownItem[],
  selectedIndexes: number[],
  onSelectChange: (value: string | string[]) => void
) => {
  if (selectedIndexes.length === 1) {
    onSelectChange(items[selectedIndexes[0]].title)
  } else {
    const values = items.reduce((acc, item, index) => {
      if (selectedIndexes.includes(index)) acc.push(item.title)
      return acc
    }, [] as string[])

    onSelectChange(values)
  }
}
