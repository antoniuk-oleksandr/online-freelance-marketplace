<script lang="ts">
  import SelectMenu from '@/common-components/Select/components/SelectMenu/SelectMenu.svelte'
  import type { SelectProps } from '@/types/SelectProps.ts'
  import { resolveSelectAnimation } from '@/common-components/Select/helper'
  import SelectLayout from '@/common-components/Select/SelectLayout.svelte'
  import SelectTrigger from '@/common-components/Select/components/SelectTrigger/SelectTrigger.svelte'

  const props: SelectProps = $props()

  let isOpen: boolean = $state(false)
  const setIsOpen = (value: boolean) => (isOpen = value)

  let showExitAnimation: boolean = $state(false)
  const setShowExitAnimation = (value: boolean) => (showExitAnimation = value)

  let timeout: number | undefined
  const setTimeoutValue = (value: number | undefined) => (timeout = value)

  $effect(() => resolveSelectAnimation(timeout, setTimeoutValue, isOpen, setShowExitAnimation))
</script>

<SelectLayout>
  <SelectTrigger SelectMenuProps={props} title={props.title} {isOpen} {setIsOpen}>
    {@render props.children()}
  </SelectTrigger>
  <SelectMenu {showExitAnimation} {isOpen} {setIsOpen} {...props} />
</SelectLayout>
