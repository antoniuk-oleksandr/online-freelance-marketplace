<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { handleSelectTriggerClick } from '@/common-components/Select/handlers'
  import type { SelectProps } from '@/types/SelectProps'

  type SelectTriggerProps = LayoutProps & {
    isOpen: boolean
    title: string
    SelectMenuProps: SelectProps
    setIsOpen?: (value: boolean) => void
  }

  const { isOpen, setIsOpen, children, title, SelectMenuProps }: SelectTriggerProps = $props()
</script>

<button
  onclick={() => handleSelectTriggerClick(isOpen, title, SelectMenuProps, setIsOpen)}
  class="Select-trigger rounded-lg border bg-light-palette-background-block dark:bg-dark-palette-background-block border-light-palette-divider dark:border-dark-palette-divider"
>
  {@render children()}
</button>
