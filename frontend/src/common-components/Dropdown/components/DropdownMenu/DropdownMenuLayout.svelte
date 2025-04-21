<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { flyFade } from '@/utils/utils'

  type DropdownMenuLayoutProps = LayoutProps & {
    menuRef: HTMLDivElement | undefined
    positionX?: 'left' | 'right'
    positionY?: 'top' | 'bottom'
    menuWidth?: string,
    menuStyles?: string
  }

  let {
    children,
    menuWidth,
    menuRef = $bindable(),
    positionX = 'right',
    positionY = 'bottom',
    menuStyles,
  }: DropdownMenuLayoutProps = $props()
</script>

<div
  class="p-6 lg:p-0 !z-[60] fixed lg:static top-0 lg:top-auto right-0 lg:right-auto size-full lg:size-fit grid lg:block place-items-center"
>
  <div
    bind:this={menuRef}
    transition:flyFade={{ y: 10, duration: 300 }}
    class="{menuStyles} w-[90%] xl:w-full sm:w-[70%] lg:w-[40%] flex max-h-96 overflow-auto flex-col shadow-xl gap-y-2 py-1 lg:p-1 lg:mt-2 rounded-lg lg:rounded-lg right-0 whitespace-nowrap bg-light-palette-background-block dark:bg-dark-palette-background-block lg:absolute h-fit 
    lg:{menuWidth ? menuWidth : 'min-w-52'} 
    {positionX === 'left' ? 'lg:left-0' : 'lg:right-0'}
    {positionY === 'bottom' ? 'lg:top-full' : 'lg:bottom-full'}"
  >
    {@render children()}
  </div>
</div>
