<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type OrderByIdChatTabMessageLayoutProps = LayoutProps & {
    wasSentByPartner: boolean
  }

  const { children, wasSentByPartner }: OrderByIdChatTabMessageLayoutProps = $props()
</script>

<div
  class="{wasSentByPartner ? 'rounded-tl-none bg-light-palette-action-hover dark:bg-dark-palette-action-hover' : 'rounded-br-none ring-1 ring-light-palette-divider dark:ring-dark-palette-divider'} items-end flex justify-between rounded-xl max-w-[55%] p-3"
>
  {@render children()}
</div>
