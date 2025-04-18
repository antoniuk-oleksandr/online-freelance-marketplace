<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type OrderByIdChatTabPartnerMessageLayoutProps = LayoutProps & {
    renderAvatar: boolean
  }

  const { children, renderAvatar }: OrderByIdChatTabPartnerMessageLayoutProps = $props()
</script>

<div class="{renderAvatar ? '' : 'ml-15'} flex gap-3">
  {@render children()}
</div>
