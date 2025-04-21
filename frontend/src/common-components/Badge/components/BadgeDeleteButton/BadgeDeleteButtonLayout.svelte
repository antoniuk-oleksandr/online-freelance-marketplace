<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { fade } from 'svelte/transition'

  type BadgeDeleteButtonLayoutProps = LayoutProps & {
    deleteFunc: (index: number) => void
    index: number
    isMobile: boolean
  }

  const { children, deleteFunc, index }: BadgeDeleteButtonLayoutProps = $props()
</script>

<button
  type="button"
  onclick={() => deleteFunc(index)}
  class="opacity-100 md:opacity-0 group-hover:opacity-100 right-0 top-0 absolute -mr-[14px] md:-mr-[13px] -mt-[13px] md:-mt-[13px] bg-light-palette-background-block dark:bg-dark-palette-background-block rounded-full p-1.5 text-red-500 hover:bg-red-100 dark:hover:bg-red-500/10 hover:text-red-600 transition-all ease-in-out duration-200 cursor-pointer"
  transition:fade={{ duration: 200 }}
>
  {@render children()}
</button>
